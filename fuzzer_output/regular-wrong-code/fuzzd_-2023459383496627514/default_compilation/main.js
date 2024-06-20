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
    static fm0(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ibq"), _dafny.Seq.UnicodeFromString("qjdrauoj")),_module.__default.safeDivisionInt(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(163)), function (_0_i0) {
          return new BigNumber(68);
        }))).Elements) {
          let _1_v0 = _compr_0;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(163)), function (_0_i0) {
            return new BigNumber(68);
          }))).contains(_1_v0)) {
            _coll0.push([(_1_v0).plus(new BigNumber(-976)),!(false)]);
          }
        }
        return _coll0;
      }()).length), new BigNumber(120)));
    };
    static fm1(globalState) {
      return new _dafny.CodePoint('h'.codePointAt(0));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)))), _dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))));
    };
    static fm3(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))))).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(false,true))).Keys.Elements) {
          let _2_v0 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(false,true))).contains(_2_v0)) {
            _coll1.push([_2_v0,_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)))]);
          }
        }
        return _coll1;
      }());
    };
    static fm4(p0, globalState) {
      let _source0 = _module.D3.create_DC8(new _dafny.CodePoint('h'.codePointAt(0)), !(false), (_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("ndasii")).length))).dtor_cf1, _dafny.Seq.UnicodeFromString("aqwpe"));
      if (_source0.is_DC8) {
        let _3___mcc_h0 = (_source0).cf17;
        let _4___mcc_h1 = (_source0).cf18;
        let _5___mcc_h2 = (_source0).cf19;
        let _6___mcc_h3 = (_source0).cf20;
        let _7_cf20 = _6___mcc_h3;
        let _8_cf19 = _5___mcc_h2;
        let _9_cf18 = _4___mcc_h1;
        let _10_cf17 = _3___mcc_h0;
        return !(_9_cf18);
      } else if (_source0.is_DC7) {
        let _11___mcc_h4 = (_source0).cf16;
        let _12_cf16 = _11___mcc_h4;
        return (false) === (true);
      } else {
        let _13___mcc_h5 = (_source0).cf21;
        let _14_cf21 = _13___mcc_h5;
        return (true) || (true);
      }
    };
    static fm8(p0, globalState) {
      return _dafny.Seq.of(new BigNumber(247), _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(new BigNumber(-345), new BigNumber(863))).length), new BigNumber((_dafny.Seq.UnicodeFromString("xdtj")).length)), new BigNumber(-684));
    };
    static fm12(globalState) {
      return _dafny.Seq.of(((_dafny.ZERO).minus(new BigNumber(((_module.D6.create_DC17((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(998), new BigNumber(270))) {
    let _15_v0 = _compr_2;
    if (((new BigNumber(998)).isLessThanOrEqualTo(_15_v0)) && ((_15_v0).isLessThan(new BigNumber(270)))) {
      _coll2.push([(_15_v0).minus(new BigNumber(322)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(true)).length))).length)]);
    }
  }
  return _coll2;
}()).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(653)), function (_16_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}), false)).dtor_cf39).length))).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("b")).length)),false)).length)))), !(true) || (!(false)));
    };
    static fm14(p0, globalState) {
      return (_dafny.Set.fromElements(!(!(false)), !(!(true)))).Union(_dafny.Set.fromElements(false));
    };
    static fm15(p0, p1, globalState) {
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.of(true, true)) : (_dafny.Seq.of(false, true))), _dafny.Seq.of(false));
    };
    static fm16(p0, globalState) {
      return _module.D0.create_DC1(true, new BigNumber(9));
    };
    static fm17(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC9(_module.D3.create_DC9(_module.D3.create_DC8(new _dafny.CodePoint('k'.codePointAt(0)), false, true, _dafny.Seq.UnicodeFromString("bvdjclmpx")))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC9(_module.D3.create_DC9(_module.D3.create_DC9(_module.D3.create_DC7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(92)), function (_17_i0) {
  return new _dafny.CodePoint('g'.codePointAt(0));
})))))))).length),(_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(!(true))));
    };
    static fm18(p0, p1, globalState) {
      return (((!(!(true))) ? (_dafny.Map.Empty.slice().updateUnsafe(false,!(false))) : (_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm19(p0, p1, globalState) {
      if (true) {
        return _module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(717),new BigNumber((_dafny.Set.fromElements(false)).length))).length), new BigNumber(59));
      } else {
        return new BigNumber((_dafny.MultiSet.fromElements(false, false, false, false, true)).cardinality());
      }
    };
    static fm20(p0, p1, globalState) {
      return (((true) ? (_module.D1.create_DC4(!(false), true, new BigNumber(747), _dafny.Seq.of(new BigNumber(883), new BigNumber(416)), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)))).length))) : (_module.D1.create_DC4(false, true, new BigNumber(-294), (_module.D1.create_DC4(false, false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length),false)).length), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("mpil")).length), new BigNumber(275), new BigNumber(715), new BigNumber((_dafny.Seq.of(true)).length), (_dafny.ZERO).minus(new BigNumber(-940))), new BigNumber(17))).dtor_cf8, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(306),!(true))).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(230))).length)))).length))))).dtor_cf8;
    };
    static fm21(globalState) {
      let _source1 = _module.D3.create_DC8(new _dafny.CodePoint('t'.codePointAt(0)), true, false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), function (_18_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}));
      if (_source1.is_DC8) {
        let _19___mcc_h0 = (_source1).cf17;
        let _20___mcc_h1 = (_source1).cf18;
        let _21___mcc_h2 = (_source1).cf19;
        let _22___mcc_h3 = (_source1).cf20;
        let _23_cf20 = _22___mcc_h3;
        let _24_cf19 = _21___mcc_h2;
        let _25_cf18 = _20___mcc_h1;
        let _26_cf17 = _19___mcc_h0;
        return _module.D0.create_DC0((_23_cf20)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_23_cf20).length))]);
      } else if (_source1.is_DC7) {
        let _27___mcc_h4 = (_source1).cf16;
        let _28_cf16 = _27___mcc_h4;
        return _module.D0.create_DC0(new _dafny.CodePoint('d'.codePointAt(0)));
      } else {
        let _29___mcc_h5 = (_source1).cf21;
        let _30_cf21 = _29___mcc_h5;
        return _module.D0.create_DC0(new _dafny.CodePoint('e'.codePointAt(0)));
      }
    };
    static fm22(globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("icl"), _dafny.Seq.UnicodeFromString("x")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(522)), function (_31_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(202)), function (_32_i1) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC8(new _dafny.CodePoint('w'.codePointAt(0)), false, ((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(!(!(false)), true, false)).cardinality()))).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-362),true)).length))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ykgprxn"), _dafny.Seq.UnicodeFromString("bkcd")));
    };
    static fm24(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(!(false), true)), _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(!(false)))));
    };
    static fm25(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(501),_dafny.Seq.UnicodeFromString("lifhwcw"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(360),_dafny.Seq.UnicodeFromString("altifx")));
    };
    static fm26(globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), function (_33_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-977)), function (_34_i1) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        });
      }))).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ujto"), _dafny.Seq.UnicodeFromString("pfcmr"), _dafny.Seq.UnicodeFromString("jnkits"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(70)), function (_35_i2) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })))).Difference(((true) ? (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ynytpecsu"))) : (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("x")))));
    };
    static fm27(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC5(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)))).length))).minus(new BigNumber(4)), false, ((false) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber(-758)))).length)) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(715))).cardinality()))).length)))), (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), function (_36_i0) {
  return new BigNumber(764);
})).length), new BigNumber(938), new BigNumber(739), new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality()), _dafny.ZERO)).IsSubsetOf(_dafny.Set.fromElements(new BigNumber(352), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("wfi")).length), new BigNumber(763))).length)))).length), new BigNumber(980), new BigNumber((function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of _dafny.IntegerRange(new BigNumber(12), new BigNumber(21))) {
    let _37_v0 = _compr_3;
    if (((new BigNumber(12)).isLessThanOrEqualTo(_37_v0)) && ((_37_v0).isLessThan(new BigNumber(21)))) {
      _coll3.push([(_37_v0).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(279))).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('c'.codePointAt(0)))).length)]);
    }
  }
  return _coll3;
}()).length))), new BigNumber(-549));
    };
    static fm28(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(465)), function (_38_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(895)))).Merge(((true) ? (_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(855))) : (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(677)))));
    };
    static fm29(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(242)), function (_39_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("jcgiuvj")).length);
      })).length))).Union(_dafny.MultiSet.fromElements(new BigNumber(-791), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("s"))).length), new BigNumber(-678))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-166)), function (_40_i1) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length)));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(((!(!(true))) ? (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(657))).cardinality())) : (new BigNumber(769))));
    };
    static fm33(p0, p1, p2, p3, globalState) {
      let _source2 = _module.D8.create_DC19(function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of _dafny.IntegerRange(new BigNumber(473), new BigNumber(374))) {
    let _41_v0 = _compr_4;
    if (((new BigNumber(473)).isLessThanOrEqualTo(_41_v0)) && ((_41_v0).isLessThan(new BigNumber(374)))) {
      _coll4.push([(_41_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("jwdxy")).length)),false]);
    }
  }
  return _coll4;
}());
      if (_source2.is_DC20) {
        let _42___mcc_h0 = (_source2).cf43;
        let _43___mcc_h1 = (_source2).cf44;
        let _44___mcc_h2 = (_source2).cf45;
        let _45___mcc_h3 = (_source2).cf46;
        let _46___mcc_h4 = (_source2).cf47;
        let _47_cf47 = _46___mcc_h4;
        let _48_cf46 = _45___mcc_h3;
        let _49_cf45 = _44___mcc_h2;
        let _50_cf44 = _43___mcc_h1;
        let _51_cf43 = _42___mcc_h0;
        return _dafny.Seq.of(_47_cf47);
      } else {
        let _52___mcc_h5 = (_source2).cf42;
        let _53_cf42 = _52___mcc_h5;
        return _dafny.Seq.of(false, false);
      }
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("sqdajpdkh")).length),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(311),!(false)))).Merge((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(229), new BigNumber(-103))) {
          let _54_v0 = _compr_5;
          if (((new BigNumber(229)).isLessThanOrEqualTo(_54_v0)) && ((_54_v0).isLessThan(new BigNumber(-103)))) {
            _coll5.push([_module.__default.safeModuloInt(_54_v0, new BigNumber(-349)),true]);
          }
        }
        return _coll5;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))).length),false)));
    };
    static fm35(globalState) {
      if (false) {
        if (true) {
          return _module.D6.create_DC17(new BigNumber(121), _dafny.Seq.UnicodeFromString("h"), (_module.D9.create_DC22(new BigNumber(19), true, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(173)), function (_55_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})).length),false))).dtor_cf50);
        } else {
          return _module.D6.create_DC17(new BigNumber(119), _dafny.Seq.UnicodeFromString("ndg"), false);
        }
      } else {
        return _module.D6.create_DC17(new BigNumber(695), _dafny.Seq.UnicodeFromString("stgojmq"), true);
      }
    };
    static m0(p0, p1, globalState) {
      let r0 = false;
      let r1 = _dafny.MultiSet.Empty;
      let r2 = _dafny.Set.Empty;
      let r3 = _dafny.ZERO;
      let _56_v0;
      _56_v0 = false;
      let _57_i0;
      _57_i0 = _dafny.ZERO;
      L0: {
        while (_56_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_57_i0)) {
              break L0;
            }
            _57_i0 = (_57_i0).plus(_dafny.ONE);
            let _58_v1;
            let _nw0 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
            _58_v1 = _nw0;
            let _59_v2;
            _59_v2 = new BigNumber(849);
            let _60_v3;
            _60_v3 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(p1, _module.__default.safeIndex(_59_v2, new BigNumber((p1).length)), !(_56_v0))).length), _59_v2, _59_v2);
            let _index0 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_58_v1).length));
            (_58_v1)[_index0] = (_59_v2).multipliedBy(new BigNumber((_60_v3).length));
            let _61_v4;
            _61_v4 = _dafny.Seq.UnicodeFromString("bcuys");
            let _index1 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_58_v1).length));
            (_58_v1)[_index1] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_61_v4, _61_v4), _dafny.Seq.Concat(_61_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(829)), function (_62_i1) {
              return new _dafny.CodePoint('v'.codePointAt(0));
            })))).length);
            r3 = _59_v2;
            r0 = _56_v0;
            r0 = _56_v0;
          }
        }
      }
      let _63_v5;
      _63_v5 = _dafny.Seq.UnicodeFromString("btljib");
      let _64_v6;
      _64_v6 = new BigNumber(-243);
      let _65_v7;
      _65_v7 = _dafny.Map.Empty.slice().updateUnsafe(_64_v6,!(_56_v0));
      let _66_v8;
      _66_v8 = _module.D8.create_DC19(_65_v7);
      let _pat_let_tv0 = _63_v5;
      let _pat_let_tv1 = _65_v7;
      _63_v5 = function (_source3) {
        if (_source3.is_DC20) {
          let _67___mcc_h0 = (_source3).cf43;
          let _68___mcc_h1 = (_source3).cf44;
          let _69___mcc_h2 = (_source3).cf45;
          let _70___mcc_h3 = (_source3).cf46;
          let _71___mcc_h4 = (_source3).cf47;
          let _72_cf47 = _71___mcc_h4;
          let _73_cf46 = _70___mcc_h3;
          let _74_cf45 = _69___mcc_h2;
          let _75_cf44 = _68___mcc_h1;
          let _76_cf43 = _67___mcc_h0;
          return _dafny.Seq.UnicodeFromString("iuyjsjqt");
        } else {
          let _77___mcc_h5 = (_source3).cf42;
          let _78_cf42 = _77___mcc_h5;
          return _pat_let_tv0;
        }
      }(function (_pat_let0_0) {
        return function (_79_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_80_dt__update_hcf42_h0) {
              return _module.D8.create_DC19(_80_dt__update_hcf42_h0);
            }(_pat_let1_0);
          }(_pat_let_tv1);
        }(_pat_let0_0);
      }(_66_v8));
      let _81_v9;
      let _nw1 = new _module.C0();
      _nw1.__ctor();
      _81_v9 = _nw1;
      let _82_v10;
      _82_v10 = _dafny.MultiSet.fromElements(_56_v0, _56_v0);
      r1 = ((_dafny.MultiSet.fromElements(!(_56_v0))).Intersect((_dafny.MultiSet.FromArray(p1)).update(_56_v0, _module.__default.abs(_64_v6)))).Union((_82_v10).Union(_dafny.MultiSet.fromElements(!(_56_v0))));
      r3 = (_64_v6).plus((_dafny.ZERO).minus(_64_v6));
      r3 = (_module.__default.fm35(globalState)).dtor_cf38;
      r0 = (_64_v6).isLessThan(_64_v6);
      r1 = _82_v10;
      let _83_v11;
      _83_v11 = _module.D3.create_DC7(_63_v5);
      let _pat_let_tv2 = p0;
      let _pat_let_tv3 = _64_v6;
      let _pat_let_tv4 = p0;
      let _pat_let_tv5 = p1;
      let _pat_let_tv6 = _64_v6;
      let _pat_let_tv7 = p1;
      let _pat_let_tv8 = _56_v0;
      let _pat_let_tv9 = _56_v0;
      let _pat_let_tv10 = _56_v0;
      let _pat_let_tv11 = _56_v0;
      let _pat_let_tv12 = _56_v0;
      r2 = function (_source4) {
        if (_source4.is_DC8) {
          let _84___mcc_h6 = (_source4).cf17;
          let _85___mcc_h7 = (_source4).cf18;
          let _86___mcc_h8 = (_source4).cf19;
          let _87___mcc_h9 = (_source4).cf20;
          let _88_cf20 = _87___mcc_h9;
          let _89_cf19 = _86___mcc_h8;
          let _90_cf18 = _85___mcc_h7;
          let _91_cf17 = _84___mcc_h6;
          return (_dafny.Set.fromElements(!(_90_cf18), _90_cf18, (_pat_let_tv2)[_module.__default.safeIndex(_pat_let_tv3, new BigNumber((_pat_let_tv4).length))])).Union(_dafny.Set.fromElements((_pat_let_tv5)[_module.__default.safeIndex(_pat_let_tv6, new BigNumber((_pat_let_tv7).length))], _90_cf18));
        } else if (_source4.is_DC7) {
          let _92___mcc_h10 = (_source4).cf16;
          let _93_cf16 = _92___mcc_h10;
          return (_dafny.Set.fromElements(_pat_let_tv8)).Difference(_dafny.Set.fromElements(_pat_let_tv9));
        } else {
          let _94___mcc_h11 = (_source4).cf21;
          let _95_cf21 = _94___mcc_h11;
          return (_dafny.Set.fromElements(_pat_let_tv10)).Intersect(_dafny.Set.fromElements(_pat_let_tv11, _pat_let_tv12));
        }
      }(_83_v11);
      let _96_v12;
      _96_v12 = new _dafny.CodePoint('w'.codePointAt(0));
      let _97_v13;
      _97_v13 = _dafny.Map.Empty.slice().updateUnsafe(_64_v6,_64_v6);
      let _98_v14;
      _98_v14 = _dafny.Seq.of(new BigNumber((_97_v13).length), _64_v6);
      let _99_v15;
      _99_v15 = _dafny.MultiSet.fromElements(_64_v6);
      r3 = (((_64_v6).isLessThanOrEqualTo(_64_v6)) ? (_module.__default.fm19(_96_v12, _56_v0, globalState)) : (new BigNumber((_dafny.Seq.Concat(_98_v14, _dafny.Seq.of(new BigNumber((_99_v15).cardinality())))).length)));
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _100_v0;
      _100_v0 = false;
      let _101_v1;
      let _nw2 = Array((new BigNumber(26)).toNumber());
      _nw2[(_dafny.ZERO).toNumber()] = _100_v0;
      _nw2[(_dafny.ONE).toNumber()] = _100_v0;
      _nw2[(new BigNumber(2)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(3)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(4)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(5)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(6)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(7)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(8)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(9)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(10)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(11)).toNumber()] = false;
      _nw2[(new BigNumber(12)).toNumber()] = true;
      _nw2[(new BigNumber(13)).toNumber()] = !(_100_v0);
      _nw2[(new BigNumber(14)).toNumber()] = !(_100_v0);
      _nw2[(new BigNumber(15)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(16)).toNumber()] = false;
      _nw2[(new BigNumber(17)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(18)).toNumber()] = false;
      _nw2[(new BigNumber(19)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(20)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(21)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(22)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(23)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(24)).toNumber()] = _100_v0;
      _nw2[(new BigNumber(25)).toNumber()] = _100_v0;
      _101_v1 = _nw2;
      let _102_v2;
      _102_v2 = new BigNumber(599);
      let _103_v3;
      _103_v3 = _dafny.MultiSet.fromElements(_102_v2);
      let _104_v4;
      _104_v4 = _dafny.Seq.of(_102_v2);
      let _105_globalState;
      let _nw3 = new _module.GlobalState();
      _nw3.__ctor(false, new BigNumber(38), new BigNumber(586), false, _101_v1, _101_v1, _103_v3, _104_v4, true, new BigNumber(249));
      _105_globalState = _nw3;
      let _106_v5;
      _106_v5 = _dafny.Map.Empty.slice().updateUnsafe(true,_module.__default.safeModuloInt(new BigNumber(-550), new BigNumber((_104_v4).length)));
      _106_v5 = _106_v5;
      let _107_v6;
      _107_v6 = _dafny.Seq.of(_100_v0, _100_v0);
      let _108_v7;
      let _109_v8;
      let _110_v9;
      let _111_v10;
      let _out0;
      let _out1;
      let _out2;
      let _out3;
      let _outcollector0 = _module.__default.m0(_107_v6, _107_v6, _105_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _out3 = _outcollector0[3];
      _108_v7 = _out0;
      _109_v8 = _out1;
      _110_v9 = _out2;
      _111_v10 = _out3;
      _111_v10 = _102_v2;
      let _112_v11;
      _112_v11 = _dafny.Set.fromElements(_102_v2, _111_v10, _111_v10, _111_v10, _102_v2);
      _112_v11 = _112_v11;
      let _113_v12;
      let _nw4 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
      _113_v12 = _nw4;
      let _index2 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
      (_113_v12)[_index2] = (_dafny.ZERO).minus((_102_v2).multipliedBy((_dafny.ZERO).minus(new BigNumber((_104_v4).length))));
      let _index3 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length));
      (_101_v1)[_index3] = !(_108_v7) || (_108_v7);
      let _114_v13;
      _114_v13 = _dafny.Seq.UnicodeFromString("fe");
      let _115_v14;
      _115_v14 = _dafny.Map.Empty.slice().updateUnsafe(_114_v13,_102_v2);
      let _index4 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
      let _index5 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length));
      let _rhs0 = !(_112_v11).equals(_112_v11);
      let _rhs1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_114_v13, _module.__default.safeIndex(_102_v2, new BigNumber((_114_v13).length)), new _dafny.CodePoint('b'.codePointAt(0))), _114_v13)).length);
      let _rhs2 = _111_v10;
      let _rhs3 = _100_v0;
      let _rhs4 = ((_108_v7) ? (_115_v14) : (_module.__default.fm0(_105_globalState)));
      let _lhs0 = _113_v12;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
      let _lhs2 = _101_v1;
      let _lhs3 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length));
      _100_v0 = _rhs0;
      _102_v2 = _rhs1;
      _lhs0[_lhs1] = _rhs2;
      _lhs2[_lhs3] = _rhs3;
      _115_v14 = _rhs4;
      if ((_112_v11).IsSubsetOf((_112_v11).Intersect(_112_v11))) {
        let _index6 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length));
        (_101_v1)[_index6] = (((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(496)), ((_116_v12) => function (_117_i0) {
          return (_116_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_116_v12).length))];
        })(_113_v12)))).update(_102_v2, _module.__default.abs(_111_v10))).update((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], _module.__default.abs((_dafny.ZERO).minus(_102_v2)))).IsProperSubsetOf(_103_v3);
        let _index7 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
        let _index8 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
        let _rhs5 = (_module.__default.safeDivisionInt((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], new BigNumber((_114_v13).length))).multipliedBy((_dafny.ZERO).minus((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))]));
        let _rhs6 = _module.__default.safeModuloInt((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], (_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))]);
        let _rhs7 = _114_v13;
        let _lhs4 = _113_v12;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
        let _lhs6 = _113_v12;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
        _lhs4[_lhs5] = _rhs5;
        _lhs6[_lhs7] = _rhs6;
        _114_v13 = _rhs7;
        let _118_v15;
        _118_v15 = _dafny.Map.Empty.slice().updateUnsafe(_108_v7,(_101_v1)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length))]);
        let _119_v16;
        _119_v16 = _dafny.Map.Empty.slice().updateUnsafe((_104_v4)[_module.__default.safeIndex(new BigNumber((_118_v15).length), new BigNumber((_104_v4).length))],_102_v2);
        let _index9 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
        let _rhs8 = _103_v3;
        let _rhs9 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_102_v2).multipliedBy(new BigNumber((_119_v16).length)), new BigNumber(373)));
        let _rhs10 = _114_v13;
        let _rhs11 = new BigNumber(192);
        let _lhs8 = _113_v12;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
        _103_v3 = _rhs8;
        _lhs8[_lhs9] = _rhs9;
        _114_v13 = _rhs10;
        _102_v2 = _rhs11;
        _114_v13 = _114_v13;
        let _120_v18;
        _120_v18 = _module.D0.create_DC0(new _dafny.CodePoint('v'.codePointAt(0)));
        let _121_v19;
        _121_v19 = new _dafny.CodePoint('x'.codePointAt(0));
        let _122_v20;
        _122_v20 = _dafny.Map.Empty.slice().updateUnsafe(_121_v19,_114_v13);
        let _123_v22;
        _123_v22 = _dafny.Set.fromElements(_121_v19);
        let _124_v24;
        _124_v24 = _dafny.Map.Empty.slice().updateUnsafe(_121_v19,_118_v15);
        let _125_v26;
        let _nw5 = Array((new BigNumber(29)).toNumber());
        _nw5[(_dafny.ZERO).toNumber()] = function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_114_v13).Elements) {
            let _126_v17 = _compr_6;
            if (_dafny.Seq.contains(_114_v13, _126_v17)) {
              _coll6.push([_126_v17,_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)))]);
            }
          }
          return _coll6;
        }();
        _nw5[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_105_globalState),_114_v13)).update((_120_v18).dtor_cf0, _module.__default.fm2(_100_v0, (_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], new _dafny.CodePoint('q'.codePointAt(0)), (_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], _105_globalState));
        _nw5[(new BigNumber(2)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(3)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(4)).toNumber()] = _module.__default.fm3(_100_v0, _105_globalState);
        _nw5[(new BigNumber(5)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(6)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(7)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_121_v19,_dafny.Seq.of(_module.__default.fm1(_105_globalState)));
        _nw5[(new BigNumber(9)).toNumber()] = (_122_v20).Merge(_122_v20);
        _nw5[(new BigNumber(10)).toNumber()] = function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of (_123_v22).Elements) {
            let _127_v21 = _compr_7;
            if ((_123_v22).contains(_127_v21)) {
              _coll7.push([_127_v21,_dafny.Seq.of(_121_v19)]);
            }
          }
          return _coll7;
        }();
        _nw5[(new BigNumber(11)).toNumber()] = _module.__default.fm3(_100_v0, _105_globalState);
        _nw5[(new BigNumber(12)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(13)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(14)).toNumber()] = (_122_v20).update(_121_v19, _114_v13);
        _nw5[(new BigNumber(15)).toNumber()] = function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_124_v24).Keys.Elements) {
            let _128_v23 = _compr_8;
            if ((_124_v24).contains(_128_v23)) {
              _coll8.push([_128_v23,_114_v13]);
            }
          }
          return _coll8;
        }();
        _nw5[(new BigNumber(16)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(17)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(18)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(19)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(20)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(21)).toNumber()] = (_122_v20).Merge(_122_v20);
        _nw5[(new BigNumber(22)).toNumber()] = ((_100_v0) ? (_122_v20) : (function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of (_123_v22).Elements) {
            let _129_v25 = _compr_9;
            if ((_123_v22).contains(_129_v25)) {
              _coll9.push([_129_v25,_114_v13]);
            }
          }
          return _coll9;
        }()));
        _nw5[(new BigNumber(23)).toNumber()] = (_122_v20).update(_121_v19, _114_v13);
        _nw5[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_121_v19,_114_v13);
        _nw5[(new BigNumber(25)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_120_v18).dtor_cf0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), ((_130_v19) => function (_131_i1) {
          return _130_v19;
        })(_121_v19)));
        _nw5[(new BigNumber(26)).toNumber()] = _module.__default.fm3((_101_v1)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length))], _105_globalState);
        _nw5[(new BigNumber(27)).toNumber()] = _122_v20;
        _nw5[(new BigNumber(28)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_121_v19,_dafny.Seq.update(_114_v13, _module.__default.safeIndex(_102_v2, new BigNumber((_114_v13).length)), _121_v19));
        _125_v26 = _nw5;
        let _index10 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_125_v26).length));
        (_125_v26)[_index10] = _122_v20;
        let _132_v27;
        _132_v27 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_121_v19,_dafny.Seq.update(_114_v13, _module.__default.safeIndex((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], new BigNumber((_114_v13).length)), _121_v19)), _122_v20, _122_v20, _122_v20, (_122_v20).Merge(_122_v20));
        let _index11 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_125_v26).length));
        (_125_v26)[_index11] = (_132_v27)[_module.__default.safeIndex(_102_v2, new BigNumber((_132_v27).length))];
      } else {
        let _133_v28;
        _133_v28 = _dafny.Map.Empty.slice().updateUnsafe(_101_v1,(_101_v1)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length))]);
        _133_v28 = (_133_v28).update(_101_v1, !(!(_100_v0) || (_108_v7)));
        _109_v8 = _109_v8;
        let _index12 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length));
        (_101_v1)[_index12] = (_109_v8).IsProperSubsetOf(_109_v8);
        let _index13 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length));
        (_101_v1)[_index13] = !(_module.__default.fm4(_112_v11, _105_globalState));
        let _134_v29;
        _134_v29 = _module.D6.create_DC17(_102_v2, _114_v13, _100_v0);
        let _135_v30;
        let _nw6 = new _module.C1();
        _nw6.__ctor((new BigNumber(-225)).isLessThanOrEqualTo((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))]), (_134_v29).dtor_cf40, _109_v8);
        _135_v30 = _nw6;
      }
      let _136_v31;
      let _137_v32;
      let _138_v33;
      let _139_v34;
      let _out4;
      let _out5;
      let _out6;
      let _out7;
      let _outcollector1 = _module.__default.m0(_107_v6, ((_108_v7) ? (_107_v6) : (_107_v6)), _105_globalState);
      _out4 = _outcollector1[0];
      _out5 = _outcollector1[1];
      _out6 = _outcollector1[2];
      _out7 = _outcollector1[3];
      _136_v31 = _out4;
      _137_v32 = _out5;
      _138_v33 = _out6;
      _139_v34 = _out7;
      let _140_v35;
      let _nw7 = Array((new BigNumber(6)).toNumber()).fill([]);
      _140_v35 = _nw7;
      let _index14 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_140_v35).length));
      (_140_v35)[_index14] = _113_v12;
      let _index15 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_140_v35).length));
      (_140_v35)[_index15] = _113_v12;
      let _141_v36;
      _141_v36 = _module.D9.create_DC23(_139_v34);
      let _142_v37;
      _142_v37 = _dafny.Map.Empty.slice().updateUnsafe((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))],_102_v2);
      let _hi0 = (((_142_v37).contains(_111_v10)) ? ((_142_v37).get(_111_v10)) : (_111_v10));
      for (let _143_i2 = (_141_v36).dtor_cf52; _143_i2.isLessThan(_hi0); _143_i2 = _143_i2.plus(_dafny.ONE)) {
        let _144_v38;
        let _nw8 = new _module.C5();
        _nw8.__ctor(_100_v0, _109_v8);
        _144_v38 = _nw8;
        let _145_v39;
        _145_v39 = _dafny.Map.Empty.slice().updateUnsafe(_144_v38,_108_v7);
        let _146_v40;
        let _nw9 = new _module.C2();
        _nw9.__ctor(!((_144_v38).f11), _143_i2, _109_v8);
        _146_v40 = _nw9;
        let _147_v41;
        let _nw10 = Array((new BigNumber(8)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _146_v40;
        _nw10[(_dafny.ONE).toNumber()] = _146_v40;
        _nw10[(new BigNumber(2)).toNumber()] = _146_v40;
        _nw10[(new BigNumber(3)).toNumber()] = _146_v40;
        _nw10[(new BigNumber(4)).toNumber()] = _146_v40;
        _nw10[(new BigNumber(5)).toNumber()] = _146_v40;
        _nw10[(new BigNumber(6)).toNumber()] = _146_v40;
        _nw10[(new BigNumber(7)).toNumber()] = _146_v40;
        _147_v41 = _nw10;
        let _148_v42;
        _148_v42 = _dafny.Map.Empty.slice().updateUnsafe(_147_v41,_108_v7);
        let _index16 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length));
        (_101_v1)[_index16] = !((((_145_v39).contains((_144_v38))) ? ((_145_v39).get((_144_v38))) : ((((_148_v42).contains(_147_v41)) ? ((_148_v42).get(_147_v41)) : (true))))) || ((_144_v38).f11);
        let _149_v43;
        let _150_v44;
        let _151_v45;
        let _152_v46;
        let _out8;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector2 = _module.__default.m0(_dafny.Seq.of((_101_v1)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length))]), _107_v6, _105_globalState);
        _out8 = _outcollector2[0];
        _out9 = _outcollector2[1];
        _out10 = _outcollector2[2];
        _out11 = _outcollector2[3];
        _149_v43 = _out8;
        _150_v44 = _out9;
        _151_v45 = _out10;
        _152_v46 = _out11;
        let _153_v47;
        _153_v47 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_107_v6));
        let _154_v48;
        _154_v48 = new _dafny.CodePoint('p'.codePointAt(0));
        let _155_v49;
        _155_v49 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_144_v38).fm5(_151_v45, new BigNumber((_153_v47).length), _152_v46, _105_globalState), _143_i2, _154_v48, _152_v46, _105_globalState),_149_v43);
        let _156_v50;
        _156_v50 = _dafny.Map.Empty.slice().updateUnsafe(_149_v43,_108_v7);
        _155_v49 = (_155_v49).update(_module.__default.fm2((((_156_v50).contains((_144_v38).f11)) ? ((_156_v50).get((_144_v38).f11)) : (_149_v43)), _102_v2, _154_v48, new BigNumber((_dafny.Seq.UnicodeFromString("ivq")).length), _105_globalState), _136_v31);
        let _157_v51;
        let _nw11 = new _module.C0();
        _nw11.__ctor();
        _157_v51 = _nw11;
        _157_v51 = _157_v51;
      }
      let _158_v52;
      _158_v52 = new _dafny.CodePoint('f'.codePointAt(0));
      let _159_v53;
      _159_v53 = _dafny.Seq.of(_114_v13, _114_v13, _module.__default.fm2(_100_v0, _111_v10, _158_v52, new BigNumber((_142_v37).length), _105_globalState), _114_v13);
      let _160_v54;
      let _nw12 = Array((new BigNumber(24)).toNumber());
      _nw12[(_dafny.ZERO).toNumber()] = (_159_v53)[_module.__default.safeIndex(_111_v10, new BigNumber((_159_v53).length))];
      _nw12[(_dafny.ONE).toNumber()] = _114_v13;
      _nw12[(new BigNumber(2)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(3)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_114_v13, _dafny.Seq.update(_dafny.Seq.update(_114_v13, _module.__default.safeIndex((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], new BigNumber((_114_v13).length)), _158_v52), _module.__default.safeIndex(_111_v10, new BigNumber((_dafny.Seq.update(_114_v13, _module.__default.safeIndex((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], new BigNumber((_114_v13).length)), _158_v52)).length)), (_114_v13)[_module.__default.safeIndex(_102_v2, new BigNumber((_114_v13).length))])), _module.__default.safeIndex((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], new BigNumber((_dafny.Seq.Concat(_114_v13, _dafny.Seq.update(_dafny.Seq.update(_114_v13, _module.__default.safeIndex((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], new BigNumber((_114_v13).length)), _158_v52), _module.__default.safeIndex(_111_v10, new BigNumber((_dafny.Seq.update(_114_v13, _module.__default.safeIndex((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))], new BigNumber((_114_v13).length)), _158_v52)).length)), (_114_v13)[_module.__default.safeIndex(_102_v2, new BigNumber((_114_v13).length))]))).length)), _module.__default.fm1(_105_globalState));
      _nw12[(new BigNumber(5)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(6)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), ((_161_v52) => function (_162_i3) {
        return _161_v52;
      })(_158_v52));
      _nw12[(new BigNumber(8)).toNumber()] = _module.__default.fm2(_108_v7, _111_v10, _158_v52, _139_v34, _105_globalState);
      _nw12[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("muykribt");
      _nw12[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_114_v13, _dafny.Seq.UnicodeFromString("atwbgc"));
      _nw12[(new BigNumber(11)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(12)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(13)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(14)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(15)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_114_v13, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("lx"), _module.__default.safeIndex(_111_v10, new BigNumber((_dafny.Seq.UnicodeFromString("lx")).length)), _158_v52));
      _nw12[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_114_v13, _module.__default.safeIndex(_111_v10, new BigNumber((_114_v13).length)), _158_v52);
      _nw12[(new BigNumber(18)).toNumber()] = _114_v13;
      _nw12[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_114_v13, _114_v13);
      _nw12[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_114_v13, _dafny.Seq.UnicodeFromString("yyix"));
      _nw12[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_114_v13, _dafny.Seq.UnicodeFromString("hkxekp"));
      _nw12[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_114_v13, _module.__default.fm2((_101_v1)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length))], new BigNumber(228), _158_v52, _102_v2, _105_globalState));
      _nw12[(new BigNumber(23)).toNumber()] = ((_100_v0) ? (_114_v13) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-159)), ((_163_v52) => function (_164_i4) {
        return _163_v52;
      })(_158_v52))));
      _160_v54 = _nw12;
      let _index17 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_160_v54).length));
      (_160_v54)[_index17] = _dafny.Seq.UnicodeFromString("hythm");
      let _index18 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_160_v54).length));
      let _index19 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
      let _rhs12 = _dafny.Seq.Concat(_114_v13, _module.__default.fm2(_100_v0, (_module.D0.create_DC1(_100_v0, (_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))])).dtor_cf2, _158_v52, new BigNumber(432), _105_globalState));
      let _rhs13 = (_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))];
      let _lhs10 = _160_v54;
      let _lhs11 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_160_v54).length));
      let _lhs12 = _113_v12;
      let _lhs13 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length));
      _lhs10[_lhs11] = _rhs12;
      _lhs12[_lhs13] = _rhs13;
      let _hi1 = (_dafny.ZERO).minus((_102_v2).multipliedBy(_111_v10));
      for (let _165_i5 = (_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))]; _165_i5.isLessThan(_hi1); _165_i5 = _165_i5.plus(_dafny.ONE)) {
        let _166_v55;
        let _nw13 = Array((new BigNumber(12)).toNumber());
        _166_v55 = _nw13;
        let _167_v56;
        _167_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(366),_166_v55);
        _167_v56 = _167_v56;
        let _168_v57;
        let _nw14 = Array((new BigNumber(28)).toNumber());
        _168_v57 = _nw14;
        let _nw15 = Array((new BigNumber(26)).toNumber());
        _168_v57 = _nw15;
        let _index20 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length));
        (_101_v1)[_index20] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_107_v6, _dafny.Seq.update(_107_v6, _module.__default.safeIndex(new BigNumber(314), new BigNumber((_107_v6).length)), _136_v31)), _107_v6);
        let _169_v58;
        _169_v58 = _dafny.Map.Empty.slice().updateUnsafe(_165_i5,(_101_v1)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_101_v1).length))]);
        _102_v2 = (_111_v10).minus(new BigNumber((_169_v58).length));
      }
      _106_v5 = (_106_v5).update(!_dafny.Seq.contains(_107_v6, _108_v7), (_dafny.ZERO).minus((_113_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_113_v12).length))]));
      _112_v11 = _112_v11;
      let _170_v59;
      let _nw16 = new _module.C3();
      _nw16.__ctor(_136_v31);
      _170_v59 = _nw16;
      let _171_v60;
      _171_v60 = _dafny.Seq.of(_170_v59, _170_v59, _170_v59, _170_v59, _170_v59);
      _171_v60 = _171_v60;
      let _172_v61;
      _172_v61 = _137_v32;
      let _pat_let_tv13 = _113_v12;
      let _pat_let_tv14 = _113_v12;
      _108_v7 = function (_source5) {
        let _173___mcc_h0 = _source5;
        let _174_cf41 = _173___mcc_h0;
        return !(new BigNumber(381)).isEqualTo((_pat_let_tv14)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_pat_let_tv13).length))]);
      }(_172_v61);
      let _175_v62;
      let _176_v63;
      let _177_v64;
      let _178_v65;
      let _out12;
      let _out13;
      let _out14;
      let _out15;
      let _outcollector3 = _module.__default.m0(_107_v6, _107_v6, _105_globalState);
      _out12 = _outcollector3[0];
      _out13 = _outcollector3[1];
      _out14 = _outcollector3[2];
      _out15 = _outcollector3[3];
      _175_v62 = _out12;
      _176_v63 = _out13;
      _177_v64 = _out14;
      _178_v65 = _out15;
      process.stdout.write(_dafny.toString(_100_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v1)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_102_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_v3).equals(_dafny.MultiSet.fromElements(new BigNumber(599)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_104_v4, _dafny.Seq.of(new BigNumber(599)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f4)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f5)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f6).equals(_dafny.MultiSet.fromElements(new BigNumber(599)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_105_globalState).f7, _dafny.Seq.of(new BigNumber(599)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO).updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_107_v6, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_108_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_109_v8).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_110_v9).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_111_v10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_112_v11).equals(_dafny.Set.fromElements(new BigNumber(599)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_113_v12)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_114_v13).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_115_v14).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ibqqjdrauoj"),_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_136_v31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_v32).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v33).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_139_v34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_140_v35)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_141_v36).dtor_cf52));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_142_v37).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,new BigNumber(192)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_158_v52));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_159_v53, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("fe"), _dafny.Seq.UnicodeFromString("fe"), _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))), _dafny.Seq.UnicodeFromString("fe")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_160_v54)[new BigNumber(7)], _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_160_v54)[new BigNumber(8)], _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(15)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(18)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(19)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(20)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(21)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_160_v54)[new BigNumber(22)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_160_v54)[new BigNumber(23)], _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_171_v60).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_172_v61)).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_v62));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_176_v63).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_177_v64).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_178_v65));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf3) {
      let $dt = new D0(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, _dafny.ZERO);
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
    static create_DC4(cf5, cf6, cf7, cf8, cf9) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC5(cf10, cf11, cf12, cf13, cf14) {
      let $dt = new D1(1);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
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
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf5 === other.cf5 && this.cf6 === other.cf6 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf4 === other.cf4;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false, false, _dafny.ZERO, _dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC6(cf15) {
      let $dt = new D2(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
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
    static create_DC8(cf17, cf18, cf19, cf20) {
      let $dt = new D3(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC7(cf16) {
      let $dt = new D3(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf21) {
      let $dt = new D3(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + this.cf20.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + this.cf16.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18 && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(new _dafny.CodePoint('D'.codePointAt(0)), false, false, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC11(cf23, cf24, cf25) {
      let $dt = new D4(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC10(cf22) {
      let $dt = new D4(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf22 === other.cf22;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_module.D0.Default(), [], _dafny.ZERO);
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
    static create_DC13(cf27, cf28, cf29, cf30) {
      let $dt = new D5(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC14(cf31, cf32, cf33, cf34, cf35) {
      let $dt = new D5(1);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC12(cf26) {
      let $dt = new D5(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC15(cf36) {
      let $dt = new D5(3);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28 && this.cf29 === other.cf29 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(_dafny.ZERO, [], [], _dafny.ZERO);
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
    static create_DC17(cf38, cf39, cf40) {
      let $dt = new D6(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC16(cf37) {
      let $dt = new D6(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf38) + ", " + this.cf39.toVerbatimString(true) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39) && this.cf40 === other.cf40;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), false);
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
    static create_DC18(cf41) {
      let $dt = new D7(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
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
    static create_DC20(cf43, cf44, cf45, cf46, cf47) {
      let $dt = new D8(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC19(cf42) {
      let $dt = new D8(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43) && this.cf44 === other.cf44 && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46) && this.cf47 === other.cf47;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC20(new _dafny.CodePoint('D'.codePointAt(0)), false, false, _dafny.Map.Empty, false);
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
    static create_DC22(cf49, cf50, cf51) {
      let $dt = new D9(0);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC23(cf52) {
      let $dt = new D9(1);
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC24(cf53) {
      let $dt = new D9(2);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC21(cf48) {
      let $dt = new D9(3);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get is_DC21() { return this.$tag === 3; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf49, other.cf49) && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC22(_dafny.ZERO, false, _dafny.Map.Empty);
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
    static create_DC25(cf54) {
      let $dt = new D10(0);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC25" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf54 === other.cf54;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
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

  $module.T2 = class T2 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.ZERO;
      this.f5 = [];
      this._f0 = false;
      this._f2 = _dafny.ZERO;
      this._f3 = false;
      this._f4 = [];
      this._f6 = _dafny.MultiSet.Empty;
      this._f7 = _dafny.Seq.of();
      this._f8 = false;
      this._f9 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [_module.T2];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm5(p0, p1, p2, globalState) {
      let _this = this;
      return true;
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, true)).length), new BigNumber(-766)), _dafny.Seq.of(new BigNumber(131)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("igq")).length));
    };
    fm13(p0, globalState) {
      let _this = this;
      return !_dafny.Seq.contains(_dafny.Seq.Concat((_module.D3.create_DC8(new _dafny.CodePoint('h'.codePointAt(0)), true, false, _dafny.Seq.UnicodeFromString("xnohburd"))).dtor_cf20, _dafny.Seq.UnicodeFromString("wami")), new _dafny.CodePoint('c'.codePointAt(0)));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f10 = _dafny.MultiSet.Empty;
      this._f14 = false;
      this._f15 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f14, f15, f10) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f10 = f10;
      return;
    }
    m5(globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Map.Empty;
      let _179_v0;
      _179_v0 = false;
      _179_v0 = _179_v0;
      let _180_v1;
      let _nw17 = new _module.C0();
      _nw17.__ctor();
      _180_v1 = _nw17;
      let _181_v2;
      _181_v2 = _module.D4.create_DC10(_180_v1);
      let _182_v3;
      _182_v3 = new BigNumber(979);
      let _183_v4;
      _183_v4 = new _dafny.CodePoint('a'.codePointAt(0));
      let _rhs14 = (_181_v2).dtor_cf22;
      let _rhs15 = _dafny.areEqual(_module.__default.fm2(_179_v0, _182_v3, _183_v4, _182_v3, globalState), _dafny.Seq.UnicodeFromString("w"));
      _180_v1 = _rhs14;
      _179_v0 = _rhs15;
      let _184_v5;
      _184_v5 = _dafny.Seq.of(_182_v3);
      let _hi2 = (_184_v5)[_module.__default.safeIndex(_182_v3, new BigNumber((_184_v5).length))];
      for (let _185_i0 = _182_v3; _185_i0.isLessThan(_hi2); _185_i0 = _185_i0.plus(_dafny.ONE)) {
        let _186_v6;
        let _nw18 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _186_v6 = _nw18;
        let _187_v7;
        _187_v7 = _dafny.Map.Empty.slice().updateUnsafe(true,_186_v6);
        let _188_v8;
        let _nw19 = Array((new BigNumber(4)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = _186_v6;
        _nw19[(_dafny.ONE).toNumber()] = _186_v6;
        _nw19[(new BigNumber(2)).toNumber()] = (((_187_v7).contains(_179_v0)) ? ((_187_v7).get(_179_v0)) : (_186_v6));
        _nw19[(new BigNumber(3)).toNumber()] = _186_v6;
        _188_v8 = _nw19;
        let _index21 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_188_v8).length));
        (_188_v8)[_index21] = _186_v6;
        let _index22 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_188_v8).length));
        (_188_v8)[_index22] = ((false) ? (_186_v6) : (_186_v6));
        let _189_v9;
        let _init0 = ((_190_v4) => function (_191_i1) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_190_v4), _dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), ((_192_v4) => function (_193_i2) {
            return _192_v4;
          })(_190_v4)));
        })(_183_v4);
        let _nw20 = Array((new BigNumber(10)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw20.length); _i0_0++) {
          _nw20[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _189_v9 = _nw20;
        let _194_v10;
        _194_v10 = _dafny.Seq.of(_183_v4);
        let _195_v11;
        _195_v11 = _dafny.MultiSet.fromElements(_185_i0, _182_v3, new BigNumber((_194_v10).length), _185_i0, new BigNumber((_184_v5).length));
        let _196_v12;
        _196_v12 = _dafny.Seq.of((_194_v10)[_module.__default.safeIndex(new BigNumber((_195_v11).cardinality()), new BigNumber((_194_v10).length))], _183_v4, _183_v4, _183_v4, new _dafny.CodePoint('s'.codePointAt(0)));
        let _index23 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_189_v9).length));
        (_189_v9)[_index23] = _dafny.Seq.Concat(_196_v12, _196_v12);
        let _index24 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_189_v9).length));
        (_189_v9)[_index24] = _196_v12;
        let _index25 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_189_v9).length));
        (_189_v9)[_index25] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-892)), ((_197_v4) => function (_198_i3) {
          return _197_v4;
        })(_183_v4));
        let _199_v13;
        let _nw21 = new _module.C0();
        _nw21.__ctor();
        _199_v13 = _nw21;
      }
      let _200_v14;
      _200_v14 = _dafny.Set.fromElements(_182_v3);
      let _201_v16;
      _201_v16 = _dafny.Seq.of(_200_v14, _dafny.Set.fromElements(new BigNumber((_184_v5).length), _182_v3));
      let _202_v17;
      let _nw22 = Array((new BigNumber(17)).toNumber());
      _nw22[(_dafny.ZERO).toNumber()] = _200_v14;
      _nw22[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(_182_v3, (_180_v1).fm6(_179_v0, _182_v3, globalState))).Union(_200_v14);
      _nw22[(new BigNumber(2)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(3)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(4)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(5)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(6)).toNumber()] = (_200_v14).Union(_200_v14);
      _nw22[(new BigNumber(7)).toNumber()] = function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(35), new BigNumber(104))) {
          let _203_v15 = _compr_10;
          if (((new BigNumber(35)).isLessThanOrEqualTo(_203_v15)) && ((_203_v15).isLessThan(new BigNumber(104)))) {
            _coll10.add(_module.__default.safeModuloInt(_203_v15, ((((_this).f10).contains((_this).f14)) ? (((_this).f10).get((_this).f14)) : (new BigNumber((_dafny.Set.fromElements((_this).f14)).length)))));
          }
        }
        return _coll10;
      }();
      _nw22[(new BigNumber(8)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(9)).toNumber()] = (_200_v14).Union(_200_v14);
      _nw22[(new BigNumber(10)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(11)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(12)).toNumber()] = (_201_v16)[_module.__default.safeIndex((_dafny.ZERO).minus(_182_v3), new BigNumber((_201_v16).length))];
      _nw22[(new BigNumber(13)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(14)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(15)).toNumber()] = _200_v14;
      _nw22[(new BigNumber(16)).toNumber()] = _200_v14;
      _202_v17 = _nw22;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_202_v17).length))) {
        let _204_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_204_i4)) && ((_204_i4).isLessThan(new BigNumber((_202_v17).length))))) {
          (_202_v17)[(_204_i4)] = _200_v14;
        }
      }
      let _205_i5;
      _205_i5 = _dafny.ZERO;
      L1: {
        while (_module.__default.fm4((((_this).f14) ? (_200_v14) : (_200_v14)), globalState)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_205_i5)) {
              break L1;
            }
            _205_i5 = (_205_i5).plus(_dafny.ONE);
            r1 = new BigNumber(618);
            if ((_this).f15) {
              let _206_v18;
              let _nw23 = new _module.C0();
              _nw23.__ctor();
              _206_v18 = _nw23;
              _206_v18 = _206_v18;
              let _207_v19;
              _207_v19 = _dafny.Seq.UnicodeFromString("efljrb");
              let _208_v20;
              _208_v20 = _dafny.Map.Empty.slice().updateUnsafe(_207_v19,true);
              let _209_v22;
              _209_v22 = _dafny.Seq.of(_207_v19);
              let _210_v23;
              _210_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f15);
              let _211_v24;
              _211_v24 = _dafny.Seq.of(_210_v23, _210_v23, _dafny.Map.Empty.slice().updateUnsafe(_179_v0,(_this).f15), _210_v23, _210_v23);
              let _212_v25;
              _212_v25 = _dafny.Seq.of((_211_v24)[_module.__default.safeIndex(_182_v3, new BigNumber((_211_v24).length))]);
              let _213_v26;
              _213_v26 = _dafny.Map.Empty.slice().updateUnsafe(((_211_v24)[_module.__default.safeIndex(_182_v3, new BigNumber((_211_v24).length))]).update(_179_v0, (_this).f14),new BigNumber(433));
              let _214_v28;
              let _nw24 = Array((new BigNumber(27)).toNumber());
              _nw24[(_dafny.ZERO).toNumber()] = new BigNumber(((_208_v20).Merge(function () {
                let _coll11 = new _dafny.Map();
                for (const _compr_11 of (_209_v22).Elements) {
                  let _215_v21 = _compr_11;
                  if (_dafny.Seq.contains(_209_v22, _215_v21)) {
                    _coll11.push([_215_v21,_179_v0]);
                  }
                }
                return _coll11;
              }())).length);
              _nw24[(_dafny.ONE).toNumber()] = _182_v3;
              _nw24[(new BigNumber(2)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(3)).toNumber()] = (_182_v3).multipliedBy(_182_v3);
              _nw24[(new BigNumber(4)).toNumber()] = new BigNumber(((_this).f10).cardinality());
              _nw24[(new BigNumber(5)).toNumber()] = (_206_v18).fm6((_this).f14, new BigNumber((_207_v19).length), globalState);
              _nw24[(new BigNumber(6)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(7)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(8)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(9)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(10)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(11)).toNumber()] = new BigNumber((_207_v19).length);
              _nw24[(new BigNumber(12)).toNumber()] = new BigNumber(350);
              _nw24[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(439)), ((_216_v4) => function (_217_i6) {
                return _216_v4;
              })(_183_v4))).length);
              _nw24[(new BigNumber(14)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(15)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(16)).toNumber()] = (_180_v1).fm6(!((_this).f15), _182_v3, globalState);
              _nw24[(new BigNumber(17)).toNumber()] = new BigNumber((_200_v14).length);
              _nw24[(new BigNumber(18)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_184_v5, _module.__default.safeIndex(_182_v3, new BigNumber((_184_v5).length)), _182_v3)).length), _182_v3);
              _nw24[(new BigNumber(20)).toNumber()] = (_180_v1).fm6((_this).f15, _182_v3, globalState);
              _nw24[(new BigNumber(21)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(22)).toNumber()] = _182_v3;
              _nw24[(new BigNumber(23)).toNumber()] = _module.__default.safeModuloInt(_182_v3, _182_v3);
              _nw24[(new BigNumber(24)).toNumber()] = (_182_v3).plus(_182_v3);
              _nw24[(new BigNumber(25)).toNumber()] = (((_213_v26).contains(_210_v23)) ? ((_213_v26).get(_210_v23)) : ((_dafny.ZERO).minus(new BigNumber((function () {
                let _coll12 = new _dafny.Set();
                for (const _compr_12 of _dafny.IntegerRange(new BigNumber(218), new BigNumber(629))) {
                  let _218_v27 = _compr_12;
                  if (((new BigNumber(218)).isLessThanOrEqualTo(_218_v27)) && ((_218_v27).isLessThan(new BigNumber(629)))) {
                    _coll12.add((_218_v27).multipliedBy(_182_v3));
                  }
                }
                return _coll12;
              }()).length))));
              _nw24[(new BigNumber(26)).toNumber()] = _182_v3;
              _214_v28 = _nw24;
              let _index26 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_214_v28).length));
              (_214_v28)[_index26] = new BigNumber(883);
              let _index27 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_214_v28).length));
              (_214_v28)[_index27] = _182_v3;
              _179_v0 = false;
              let _219_v29;
              let _init1 = function (_220_i7) {
                return (_this).f15;
              };
              let _nw25 = Array((new BigNumber(9)).toNumber());
              for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw25.length); _i0_1++) {
                _nw25[_i0_1] = _init1(new BigNumber(_i0_1));
              }
              _219_v29 = _nw25;
              r0 = _219_v29;
              let _rhs16 = ((((_this).f15) ? ((_214_v28)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_214_v28).length))]) : (new BigNumber((_207_v19).length)))).plus(_182_v3);
              let _rhs17 = (_this).f14;
              let _lhs14 = globalState;
              _lhs14.f1 = _rhs16;
              _179_v0 = _rhs17;
            } else {
              let _221_v30;
              let _nw26 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
              _221_v30 = _nw26;
              let _index28 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_221_v30).length));
              (_221_v30)[_index28] = _module.__default.safeDivisionInt(_182_v3, new BigNumber(((_this).f10).cardinality()));
              let _index29 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_221_v30).length));
              (_221_v30)[_index29] = ((((_this).f10).contains((_182_v3).isEqualTo(_182_v3))) ? (((_this).f10).get((_182_v3).isEqualTo(_182_v3))) : (_182_v3));
              let _222_v31;
              _222_v31 = _dafny.Seq.of((_this).f14, (_this).f14);
              let _rhs18 = _dafny.Seq.Concat(_222_v31, _222_v31);
              let _rhs19 = _179_v0;
              let _rhs20 = new BigNumber(490);
              let _lhs15 = globalState;
              _222_v31 = _rhs18;
              _179_v0 = _rhs19;
              _lhs15.f1 = _rhs20;
              let _223_v32;
              _223_v32 = _module.D1.create_DC5(_182_v3, (_this).f14, new BigNumber(166), !(_179_v0) || (_179_v0), _182_v3);
              let _224_v33;
              _224_v33 = _dafny.Seq.UnicodeFromString("ko");
              _223_v32 = _module.D1.create_DC5((_221_v30)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_221_v30).length))], (_this).f14, _module.__default.safeModuloInt(new BigNumber(791), (_221_v30)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_221_v30).length))]), _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("scqxnyhq"), _224_v33), new BigNumber((_dafny.Seq.Concat(_222_v31, _222_v31)).length));
              let _225_v34;
              _225_v34 = _dafny.MultiSet.fromElements(new BigNumber(330));
              _182_v3 = (((_225_v34).contains((_221_v30)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_221_v30).length))])) ? ((_225_v34).get((_221_v30)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_221_v30).length))])) : ((_221_v30)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_221_v30).length))]));
              let _226_v35;
              let _init2 = function (_227_i8) {
                return (_this).f14;
              };
              let _nw27 = Array((new BigNumber(28)).toNumber());
              for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw27.length); _i0_2++) {
                _nw27[_i0_2] = _init2(new BigNumber(_i0_2));
              }
              _226_v35 = _nw27;
              let _index30 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_226_v35).length));
              (_226_v35)[_index30] = _179_v0;
              let _index31 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_226_v35).length));
              let _rhs21 = _179_v0;
              let _rhs22 = !((_this).f14);
              let _rhs23 = _224_v33;
              let _rhs24 = (_221_v30)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_221_v30).length))];
              let _rhs25 = (_221_v30)[_module.__default.safeIndex(new BigNumber(249), new BigNumber((_221_v30).length))];
              let _lhs16 = _226_v35;
              let _lhs17 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_226_v35).length));
              let _lhs18 = globalState;
              let _lhs19 = globalState;
              _lhs16[_lhs17] = _rhs21;
              _179_v0 = _rhs22;
              _224_v33 = _rhs23;
              _lhs18.f1 = _rhs24;
              _lhs19.f1 = _rhs25;
            }
            let _228_v36;
            _228_v36 = _dafny.MultiSet.FromArray(_184_v5);
            let _source6 = _228_v36;
            let _229___mcc_h0 = _source6;
            let _230_cf15 = _229___mcc_h0;
            let _231_v37;
            let _nw28 = new _module.C0();
            _nw28.__ctor();
            _231_v37 = _nw28;
            let _232_v38;
            let _nw29 = new _module.C0();
            _nw29.__ctor();
            _232_v38 = _nw29;
            let _233_v39;
            let _nw30 = Array((new BigNumber(20)).toNumber());
            _nw30[(_dafny.ZERO).toNumber()] = (_this).f15;
            _nw30[(_dafny.ONE).toNumber()] = (_this).f15;
            _nw30[(new BigNumber(2)).toNumber()] = _179_v0;
            _nw30[(new BigNumber(3)).toNumber()] = _179_v0;
            _nw30[(new BigNumber(4)).toNumber()] = (_this).f15;
            _nw30[(new BigNumber(5)).toNumber()] = (_this).f14;
            _nw30[(new BigNumber(6)).toNumber()] = _module.__default.fm4(_200_v14, globalState);
            _nw30[(new BigNumber(7)).toNumber()] = false;
            _nw30[(new BigNumber(8)).toNumber()] = _179_v0;
            _nw30[(new BigNumber(9)).toNumber()] = _179_v0;
            _nw30[(new BigNumber(10)).toNumber()] = (_this).f15;
            _nw30[(new BigNumber(11)).toNumber()] = (_this).f15;
            _nw30[(new BigNumber(12)).toNumber()] = true;
            _nw30[(new BigNumber(13)).toNumber()] = !(true);
            _nw30[(new BigNumber(14)).toNumber()] = (_this).f14;
            _nw30[(new BigNumber(15)).toNumber()] = (_this).f15;
            _nw30[(new BigNumber(16)).toNumber()] = (_this).f15;
            _nw30[(new BigNumber(17)).toNumber()] = (_this).f15;
            _nw30[(new BigNumber(18)).toNumber()] = _179_v0;
            _nw30[(new BigNumber(19)).toNumber()] = (_this).f15;
            _233_v39 = _nw30;
            let _234_v40;
            _234_v40 = _dafny.MultiSet.fromElements(_233_v39);
            _179_v0 = ((_234_v40).Intersect(_234_v40)).IsDisjointFrom((_234_v40).Difference(_234_v40));
            (globalState).f1 = (_182_v3).plus(_182_v3);
            let _rhs26 = !(_179_v0);
            let _rhs27 = (_this).f14;
            let _rhs28 = !(new BigNumber(-259)).isEqualTo((_180_v1).fm6((_this).f14, _182_v3, globalState));
            let _rhs29 = new BigNumber(2);
            _179_v0 = _rhs26;
            _179_v0 = _rhs27;
            _179_v0 = _rhs28;
            r1 = _rhs29;
          }
        }
      }
      let _hi3 = _182_v3;
      for (let _235_i9 = (_182_v3).multipliedBy(_182_v3); _235_i9.isLessThan(_hi3); _235_i9 = _235_i9.plus(_dafny.ONE)) {
        let _236_v41;
        _236_v41 = _dafny.Map.Empty.slice().updateUnsafe(_182_v3,_235_i9);
        _236_v41 = (_236_v41).update((_dafny.ZERO).minus((_235_i9).minus(new BigNumber(341))), _235_i9);
        let _237_v42;
        _237_v42 = _dafny.Seq.UnicodeFromString("vkio");
        let _238_v43;
        _238_v43 = _dafny.Seq.of((_this).f14);
        let _239_v44;
        let _init3 = function (_240_i10) {
          return (_this).f15;
        };
        let _nw31 = Array((new BigNumber(23)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw31.length); _i0_3++) {
          _nw31[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _239_v44 = _nw31;
        let _241_v45;
        _241_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_module.__default.fm1(globalState), _183_v4, _183_v4)).length),_239_v44);
        let _242_v48;
        let _nw32 = Array((new BigNumber(25)).toNumber());
        _nw32[(_dafny.ZERO).toNumber()] = (_this).f14;
        _nw32[(_dafny.ONE).toNumber()] = (_this).f14;
        _nw32[(new BigNumber(2)).toNumber()] = (_this).f14;
        _nw32[(new BigNumber(3)).toNumber()] = (_this).f15;
        _nw32[(new BigNumber(4)).toNumber()] = ((_this).f10).IsSubsetOf((_this).f10);
        _nw32[(new BigNumber(5)).toNumber()] = _179_v0;
        _nw32[(new BigNumber(6)).toNumber()] = ((_this).f14) || ((_this).f14);
        _nw32[(new BigNumber(7)).toNumber()] = (_this).f15;
        _nw32[(new BigNumber(8)).toNumber()] = _179_v0;
        _nw32[(new BigNumber(9)).toNumber()] = (_235_i9).isLessThanOrEqualTo(_182_v3);
        _nw32[(new BigNumber(10)).toNumber()] = _179_v0;
        _nw32[(new BigNumber(11)).toNumber()] = (_this).f15;
        _nw32[(new BigNumber(12)).toNumber()] = !_dafny.areEqual(_237_v42, _237_v42);
        _nw32[(new BigNumber(13)).toNumber()] = !((_this).f15);
        _nw32[(new BigNumber(14)).toNumber()] = (_180_v1).fm5(_module.__default.fm14((_this).f14, globalState), _235_i9, _235_i9, globalState);
        _nw32[(new BigNumber(15)).toNumber()] = !((_this).f14);
        _nw32[(new BigNumber(16)).toNumber()] = _dafny.areEqual(_238_v43, _dafny.Seq.update(_238_v43, _module.__default.safeIndex(new BigNumber((_238_v43).length), new BigNumber((_238_v43).length)), (_this).f15));
        _nw32[(new BigNumber(17)).toNumber()] = (_this).f15;
        _nw32[(new BigNumber(18)).toNumber()] = true;
        _nw32[(new BigNumber(19)).toNumber()] = ((_180_v1).fm13((_this).f15, globalState)) === (!((_this).f15));
        _nw32[(new BigNumber(20)).toNumber()] = (_this).f14;
        _nw32[(new BigNumber(21)).toNumber()] = (_this).f15;
        _nw32[(new BigNumber(22)).toNumber()] = (new BigNumber(-262)).isLessThan(_235_i9);
        _nw32[(new BigNumber(23)).toNumber()] = (_235_i9).isLessThanOrEqualTo(new BigNumber((_241_v45).length));
        _nw32[(new BigNumber(24)).toNumber()] = (function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of _dafny.IntegerRange(new BigNumber(708), new BigNumber(526))) {
            let _243_v46 = _compr_13;
            if (((new BigNumber(708)).isLessThanOrEqualTo(_243_v46)) && ((_243_v46).isLessThan(new BigNumber(526)))) {
              _coll13.add(_module.__default.safeDivisionInt(_243_v46, _182_v3));
            }
          }
          return _coll13;
        }()).IsDisjointFrom(function () {
          let _coll14 = new _dafny.Set();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(184), new BigNumber(772))) {
            let _244_v47 = _compr_14;
            if (((new BigNumber(184)).isLessThanOrEqualTo(_244_v47)) && ((_244_v47).isLessThan(new BigNumber(772)))) {
              _coll14.add((_244_v47).multipliedBy(_182_v3));
            }
          }
          return _coll14;
        }());
        _242_v48 = _nw32;
        let _index32 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_242_v48).length));
        (_242_v48)[_index32] = _179_v0;
        let _245_v49;
        _245_v49 = _dafny.Map.Empty.slice().updateUnsafe(_179_v0,_237_v42);
        let _246_v50;
        _246_v50 = _dafny.MultiSet.fromElements(new BigNumber((_245_v49).length));
        let _247_v51;
        _247_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((_246_v50).cardinality()));
        let _index33 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_242_v48).length));
        (_242_v48)[_index33] = !((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(842))).Merge(_247_v51)).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,_235_i9));
        if ((_this).f14) {
          _237_v42 = _237_v42;
          _237_v42 = (function (_pat_let2_0) {
            return function (_248_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_250_dt__update_hcf16_h0) {
                  return _module.D3.create_DC7(_250_dt__update_hcf16_h0);
                }(_pat_let3_0);
              }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(228)), function (_249_i11) {
                return new _dafny.CodePoint('s'.codePointAt(0));
              }));
            }(_pat_let2_0);
          }(_module.D3.create_DC7(_237_v42))).dtor_cf16;
          let _index34 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_242_v48).length));
          (_242_v48)[_index34] = (_this).f15;
          let _251_v52;
          let _nw33 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _251_v52 = _nw33;
          let _index35 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_251_v52).length));
          (_251_v52)[_index35] = _235_i9;
          let _index36 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_251_v52).length));
          let _rhs30 = _182_v3;
          let _rhs31 = _182_v3;
          let _lhs20 = _251_v52;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_251_v52).length));
          r1 = _rhs30;
          _lhs20[_lhs21] = _rhs31;
          let _252_v53;
          _252_v53 = _dafny.Set.fromElements(_179_v0);
          let _index37 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_251_v52).length));
          (_251_v52)[_index37] = (_dafny.ZERO).minus((_184_v5)[_module.__default.safeIndex(new BigNumber((_252_v53).length), new BigNumber((_184_v5).length))]);
        } else {
          let _253_v54;
          _253_v54 = _dafny.MultiSet.fromElements((((_this).f14) ? (_237_v42) : (_237_v42)));
          _253_v54 = (_253_v54).Intersect(((_253_v54).update(_237_v42, _module.__default.abs(new BigNumber((_238_v43).length)))).Union(_253_v54));
          let _254_v55;
          let _nw34 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _254_v55 = _nw34;
          let _index38 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_254_v55).length));
          (_254_v55)[_index38] = _235_i9;
          let _255_v56;
          _255_v56 = _dafny.Map.Empty.slice().updateUnsafe(_182_v3,_179_v0);
          let _index39 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_254_v55).length));
          let _rhs32 = _182_v3;
          let _rhs33 = (_180_v1).fm6((((_255_v56).contains(_182_v3)) ? ((_255_v56).get(_182_v3)) : ((_this).f15)), _182_v3, globalState);
          let _lhs22 = globalState;
          let _lhs23 = _254_v55;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_254_v55).length));
          _lhs22.f1 = _rhs32;
          _lhs23[_lhs24] = _rhs33;
          _182_v3 = new BigNumber((_237_v42).length);
          let _256_v57;
          _256_v57 = _dafny.MultiSet.fromElements(_183_v4, _183_v4, _183_v4);
          let _257_v58;
          _257_v58 = _module.D0.create_DC1((_this).f15, new BigNumber((_256_v57).cardinality()));
          let _index40 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_242_v48).length));
          (_242_v48)[_index40] = (_257_v58).dtor_cf1;
          let _258_v59;
          let _nw35 = new _module.C0();
          _nw35.__ctor();
          _258_v59 = _nw35;
        }
        _238_v43 = _dafny.Seq.Concat(_238_v43, _238_v43);
      }
      let _259_v60;
      let _init4 = ((_260_v0) => function (_261_i12) {
        return (_dafny.Set.fromElements(true, _260_v0)).IsProperSubsetOf(_dafny.Set.fromElements(_260_v0));
      })(_179_v0);
      let _nw36 = Array((new BigNumber(2)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw36.length); _i0_4++) {
        _nw36[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _259_v60 = _nw36;
      r0 = _259_v60;
      r1 = new BigNumber(34);
      let _262_v61;
      _262_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
      r2 = ((_262_v61).Merge(_262_v61)).Merge(_262_v61);
      return [r0, r1, r2];
    }
    m6(p0, p1, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.Map.Empty;
      let r2 = undefined;
      let _263_v1;
      _263_v1 = _dafny.Seq.of(_module.__default.fm4(function () {
        let _coll15 = new _dafny.Set();
        for (const _compr_15 of _dafny.IntegerRange(new BigNumber(-679), new BigNumber(733))) {
          let _264_v0 = _compr_15;
          if (((new BigNumber(-679)).isLessThanOrEqualTo(_264_v0)) && ((_264_v0).isLessThan(new BigNumber(733)))) {
            _coll15.add((_264_v0).multipliedBy(p1));
          }
        }
        return _coll15;
      }(), globalState));
      if ((_263_v1)[_module.__default.safeIndex(p1, new BigNumber((_263_v1).length))]) {
        let _265_v2;
        let _init5 = ((_266_p1) => function (_267_i0) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_266_p1,false)).update(_266_p1, (_this).f14);
        })(p1);
        let _nw37 = Array((new BigNumber(25)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw37.length); _i0_5++) {
          _nw37[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _265_v2 = _nw37;
        _265_v2 = ((!((_this).f15)) ? (_265_v2) : (_265_v2));
        let _268_v3;
        let _nw38 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _268_v3 = _nw38;
        let _269_v4;
        _269_v4 = _dafny.Seq.UnicodeFromString("xcoyciqnu");
        let _index41 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_268_v3).length));
        (_268_v3)[_index41] = _269_v4;
        let _index42 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_268_v3).length));
        (_268_v3)[_index42] = _269_v4;
        let _270_v5;
        _270_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        _270_v5 = (_270_v5).Merge(_270_v5);
        let _271_v6;
        _271_v6 = _dafny.Seq.of((_268_v3)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_268_v3).length))], (_268_v3)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_268_v3).length))]);
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_271_v6, _271_v6), _271_v6)) {
          let _272_v7;
          _272_v7 = _dafny.MultiSet.fromElements((_this).f14, (_this).f15, (_this).f14, (_this).f15, (_this).f15);
          let _273_v8;
          _273_v8 = false;
          let _274_v9;
          _274_v9 = _dafny.Set.fromElements(_dafny.areEqual((_268_v3)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_268_v3).length))], _dafny.Seq.UnicodeFromString("twrc")));
          let _275_v10;
          _275_v10 = _module.D5.create_DC12(_274_v9);
          let _rhs34 = _dafny.MultiSet.FromArray(_module.__default.fm15(new _dafny.CodePoint('w'.codePointAt(0)), new BigNumber(767), globalState));
          let _rhs35 = !((_this).f14) || ((_this).f14);
          let _rhs36 = (_275_v10).dtor_cf26;
          _272_v7 = _rhs34;
          _273_v8 = _rhs35;
          _274_v9 = _rhs36;
          let _276_v11;
          _276_v11 = new _dafny.CodePoint('d'.codePointAt(0));
          let _277_v12;
          let _nw39 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _277_v12 = _nw39;
          let _278_v13;
          let _nw40 = Array((new BigNumber(11)).toNumber());
          _nw40[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw40[(_dafny.ONE).toNumber()] = (p1).minus(p1);
          _nw40[(new BigNumber(2)).toNumber()] = p1;
          _nw40[(new BigNumber(3)).toNumber()] = p1;
          _nw40[(new BigNumber(4)).toNumber()] = p1;
          _nw40[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_module.__default.fm2(_273_v8, (_dafny.ZERO).minus(p1), _276_v11, p1, globalState), _269_v4)).length);
          _nw40[(new BigNumber(6)).toNumber()] = p1;
          _nw40[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(p1, p1);
          _nw40[(new BigNumber(8)).toNumber()] = p1;
          _nw40[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_273_v8,_277_v12)).length);
          _nw40[(new BigNumber(10)).toNumber()] = p1;
          _278_v13 = _nw40;
          let _index43 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_278_v13).length));
          (_278_v13)[_index43] = p1;
          let _279_v14;
          _279_v14 = _module.D0.create_DC1((_module.D1.create_DC5(p1, (_this).f15, p1, _273_v8, p1)).dtor_cf11, p1);
          let _280_v15;
          let _nw41 = Array((new BigNumber(22)).toNumber()).fill(false);
          _280_v15 = _nw41;
          let _281_v16;
          _281_v16 = _module.D4.create_DC11(_279_v14, _280_v15, p1);
          let _index44 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_278_v13).length));
          (_278_v13)[_index44] = (_281_v16).dtor_cf25;
          let _282_v17;
          let _nw42 = new _module.C0();
          _nw42.__ctor();
          _282_v17 = _nw42;
          _276_v11 = _276_v11;
          let _283_v18;
          _283_v18 = _dafny.Seq.of((_dafny.ZERO).minus((_278_v13)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_278_v13).length))]));
          (globalState).f1 = (new BigNumber(((_this).f10).cardinality())).minus(_module.__default.safeModuloInt((_278_v13)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_278_v13).length))], (_283_v18)[_module.__default.safeIndex(p1, new BigNumber((_283_v18).length))]));
        } else {
          let _284_v19;
          _284_v19 = _module.D1.create_DC5(new BigNumber((_dafny.Seq.UnicodeFromString("o")).length), (_this).f14, p1, (_this).f15, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f14)).length));
          let _285_v20;
          _285_v20 = _dafny.Map.Empty.slice().updateUnsafe(_284_v19,(_this).f15);
          let _pat_let_tv15 = p1;
          let _pat_let_tv16 = p1;
          _285_v20 = (_285_v20).update(function (_pat_let4_0) {
            return function (_286_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_287_dt__update_hcf10_h0) {
                  return function (_pat_let6_0) {
                    return function (_288_dt__update_hcf14_h0) {
                      return function (_pat_let7_0) {
                        return function (_289_dt__update_hcf11_h0) {
                          return _module.D1.create_DC5(_287_dt__update_hcf10_h0, _289_dt__update_hcf11_h0, (_286_dt__update__tmp_h0).dtor_cf12, (_286_dt__update__tmp_h0).dtor_cf13, _288_dt__update_hcf14_h0);
                        }(_pat_let7_0);
                      }((_this).f14);
                    }(_pat_let6_0);
                  }(_pat_let_tv16);
                }(_pat_let5_0);
              }(_pat_let_tv15);
            }(_pat_let4_0);
          }(_284_v19), (_this).f14);
          let _index45 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_268_v3).length));
          (_268_v3)[_index45] = (_268_v3)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_268_v3).length))];
          let _290_v21;
          _290_v21 = _dafny.Seq.of(_263_v1);
          let _291_v22;
          _291_v22 = new _dafny.CodePoint('m'.codePointAt(0));
          let _292_v23;
          _292_v23 = _dafny.Map.Empty.slice().updateUnsafe(_291_v22,(_this).f15);
          let _293_v24;
          let _294_v25;
          let _295_v26;
          let _296_v27;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector4 = _module.__default.m0(_dafny.Seq.Concat(_dafny.Seq.update(_263_v1, _module.__default.safeIndex(p1, new BigNumber((_263_v1).length)), false), _263_v1), (_290_v21)[_module.__default.safeIndex(new BigNumber((_292_v23).length), new BigNumber((_290_v21).length))], globalState);
          _out16 = _outcollector4[0];
          _out17 = _outcollector4[1];
          _out18 = _outcollector4[2];
          _out19 = _outcollector4[3];
          _293_v24 = _out16;
          _294_v25 = _out17;
          _295_v26 = _out18;
          _296_v27 = _out19;
          _269_v4 = _dafny.Seq.update(_269_v4, _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_269_v4).length), new BigNumber(((_this).f10).cardinality()))), new BigNumber((_269_v4).length)), _291_v22);
          _293_v24 = _293_v24;
        }
        let _297_v28;
        let _nw43 = new _module.C0();
        _nw43.__ctor();
        _297_v28 = _nw43;
      } else {
        let _298_v29;
        let _nw44 = new _module.C0();
        _nw44.__ctor();
        _298_v29 = _nw44;
        let _299_v30;
        _299_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f15);
        let _300_v31;
        _300_v31 = _dafny.Set.fromElements(p1);
        _299_v30 = (_299_v30).update(_module.__default.safeModuloInt(p1, p1), ((!((_this).f14)) ? ((_this).f14) : (_module.__default.fm4(_300_v31, globalState))));
        let _301_v32;
        _301_v32 = false;
        _301_v32 = (_this).f15;
        let _302_v33;
        _302_v33 = new _dafny.CodePoint('o'.codePointAt(0));
        let _303_v34;
        _303_v34 = _dafny.Set.fromElements(_302_v33, _302_v33, _302_v33, _302_v33, _302_v33);
        let _304_v35;
        _304_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(426),p1);
        let _305_v36;
        let _nw45 = Array((new BigNumber(20)).toNumber());
        _nw45[(_dafny.ZERO).toNumber()] = p1;
        _nw45[(_dafny.ONE).toNumber()] = p1;
        _nw45[(new BigNumber(2)).toNumber()] = p1;
        _nw45[(new BigNumber(3)).toNumber()] = (_298_v29).fm6(_301_v32, p1, globalState);
        _nw45[(new BigNumber(4)).toNumber()] = p1;
        _nw45[(new BigNumber(5)).toNumber()] = p1;
        _nw45[(new BigNumber(6)).toNumber()] = p1;
        _nw45[(new BigNumber(7)).toNumber()] = p1;
        _nw45[(new BigNumber(8)).toNumber()] = (((_304_v35).contains(p1)) ? ((_304_v35).get(p1)) : (p1));
        _nw45[(new BigNumber(9)).toNumber()] = new BigNumber(746);
        _nw45[(new BigNumber(10)).toNumber()] = p1;
        _nw45[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_298_v29).fm6((_this).f15, p1, globalState));
        _nw45[(new BigNumber(12)).toNumber()] = p1;
        _nw45[(new BigNumber(13)).toNumber()] = p1;
        _nw45[(new BigNumber(14)).toNumber()] = (_298_v29).fm6((_this).f14, (_dafny.ZERO).minus(p1), globalState);
        _nw45[(new BigNumber(15)).toNumber()] = new BigNumber((_263_v1).length);
        _nw45[(new BigNumber(16)).toNumber()] = p1;
        _nw45[(new BigNumber(17)).toNumber()] = p1;
        _nw45[(new BigNumber(18)).toNumber()] = p1;
        _nw45[(new BigNumber(19)).toNumber()] = p1;
        _305_v36 = _nw45;
        let _306_v37;
        _306_v37 = _module.D0.create_DC1(!((_this).f14), p1);
        let _307_v38;
        _307_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_306_v37);
        let _308_v39;
        _308_v39 = _dafny.Map.Empty.slice().updateUnsafe(_305_v36,(_this).f15);
        let _309_v40;
        _309_v40 = _dafny.Seq.of(_307_v38, _307_v38);
        let _pat_let_tv17 = _301_v32;
        let _310_v41;
        let _nw46 = Array((new BigNumber(15)).toNumber());
        _nw46[(_dafny.ZERO).toNumber()] = _307_v38;
        _nw46[(_dafny.ONE).toNumber()] = _307_v38;
        _nw46[(new BigNumber(2)).toNumber()] = _307_v38;
        _nw46[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_308_v39).contains(_305_v36)) ? ((_308_v39).get(_305_v36)) : ((_this).f15)),_306_v37);
        _nw46[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_301_v32,function (_pat_let8_0) {
          return function (_311_dt__update__tmp_h1) {
            return function (_pat_let9_0) {
              return function (_312_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_312_dt__update_hcf1_h0, (_311_dt__update__tmp_h1).dtor_cf2);
              }(_pat_let9_0);
            }(_pat_let_tv17);
          }(_pat_let8_0);
        }(_306_v37));
        _nw46[(new BigNumber(5)).toNumber()] = (_307_v38).update(_301_v32, _306_v37);
        _nw46[(new BigNumber(6)).toNumber()] = _307_v38;
        _nw46[(new BigNumber(7)).toNumber()] = _307_v38;
        _nw46[(new BigNumber(8)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe(_301_v32,_306_v37)).update(_301_v32, _306_v37)).update((_this).f15, _306_v37);
        _nw46[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_301_v32,_306_v37);
        _nw46[(new BigNumber(10)).toNumber()] = (_309_v40)[_module.__default.safeIndex(new BigNumber((_304_v35).length), new BigNumber((_309_v40).length))];
        _nw46[(new BigNumber(11)).toNumber()] = _307_v38;
        _nw46[(new BigNumber(12)).toNumber()] = _307_v38;
        _nw46[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_301_v32,_module.D0.create_DC1((_this).f14, p1));
        _nw46[(new BigNumber(14)).toNumber()] = _307_v38;
        _310_v41 = _nw46;
        let _313_v42;
        _313_v42 = _module.D5.create_DC13((_dafny.ZERO).minus((p1).multipliedBy(new BigNumber((_303_v34).length))), _305_v36, _310_v41, p1);
        let _source7 = _313_v42;
        if (_source7.is_DC13) {
          let _314___mcc_h0 = (_source7).cf27;
          let _315___mcc_h1 = (_source7).cf28;
          let _316___mcc_h2 = (_source7).cf29;
          let _317___mcc_h3 = (_source7).cf30;
          let _318_cf30 = _317___mcc_h3;
          let _319_cf29 = _316___mcc_h2;
          let _320_cf28 = _315___mcc_h1;
          let _321_cf27 = _314___mcc_h0;
          let _322_v43;
          let _nw47 = Array((new BigNumber(17)).toNumber());
          _nw47[(_dafny.ZERO).toNumber()] = (_this).f15;
          _nw47[(_dafny.ONE).toNumber()] = !(_301_v32);
          _nw47[(new BigNumber(2)).toNumber()] = false;
          _nw47[(new BigNumber(3)).toNumber()] = _301_v32;
          _nw47[(new BigNumber(4)).toNumber()] = _301_v32;
          _nw47[(new BigNumber(5)).toNumber()] = _301_v32;
          _nw47[(new BigNumber(6)).toNumber()] = (_this).f14;
          _nw47[(new BigNumber(7)).toNumber()] = ((_this).f15) && (!(_301_v32));
          _nw47[(new BigNumber(8)).toNumber()] = (_this).f15;
          _nw47[(new BigNumber(9)).toNumber()] = (_this).f15;
          _nw47[(new BigNumber(10)).toNumber()] = (_this).f14;
          _nw47[(new BigNumber(11)).toNumber()] = _301_v32;
          _nw47[(new BigNumber(12)).toNumber()] = (_this).f14;
          _nw47[(new BigNumber(13)).toNumber()] = _301_v32;
          _nw47[(new BigNumber(14)).toNumber()] = (_this).f14;
          _nw47[(new BigNumber(15)).toNumber()] = (_this).f15;
          _nw47[(new BigNumber(16)).toNumber()] = _301_v32;
          _322_v43 = _nw47;
          let _index46 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_322_v43).length));
          (_322_v43)[_index46] = (_this).f14;
          let _323_v44;
          _323_v44 = _dafny.Seq.UnicodeFromString("t");
          let _index47 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_322_v43).length));
          (_322_v43)[_index47] = _dafny.Seq.IsPrefixOf(_323_v44, _323_v44);
          let _324_v45;
          let _nw48 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _324_v45 = _nw48;
          let _index48 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_324_v45).length));
          (_324_v45)[_index48] = new _dafny.CodePoint('c'.codePointAt(0));
          let _index49 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_324_v45).length));
          (_324_v45)[_index49] = _302_v33;
          let _325_v46;
          _325_v46 = _dafny.Map.Empty.slice().updateUnsafe(_263_v1,true);
          _299_v30 = (_299_v30).update(_321_cf27, (((_325_v46).contains(_263_v1)) ? ((_325_v46).get(_263_v1)) : (_301_v32)));
          let _326_v47;
          let _init6 = function (_327_i1) {
            return _dafny.Seq.of((_this).f15);
          };
          let _nw49 = Array((new BigNumber(2)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw49.length); _i0_6++) {
            _nw49[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _326_v47 = _nw49;
          let _index50 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_326_v47).length));
          (_326_v47)[_index50] = _dafny.Seq.Concat(_263_v1, _263_v1);
          let _index51 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_326_v47).length));
          (_326_v47)[_index51] = _263_v1;
        } else if (_source7.is_DC14) {
          let _328___mcc_h4 = (_source7).cf31;
          let _329___mcc_h5 = (_source7).cf32;
          let _330___mcc_h6 = (_source7).cf33;
          let _331___mcc_h7 = (_source7).cf34;
          let _332___mcc_h8 = (_source7).cf35;
          let _333_cf35 = _332___mcc_h8;
          let _334_cf34 = _331___mcc_h7;
          let _335_cf33 = _330___mcc_h6;
          let _336_cf32 = _329___mcc_h5;
          let _337_cf31 = _328___mcc_h4;
          let _index52 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_305_v36).length));
          (_305_v36)[_index52] = (((_this).f14) ? (_336_cf32) : (new BigNumber((_300_v31).length)));
          let _index53 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_305_v36).length));
          (_305_v36)[_index53] = (p1).multipliedBy(new BigNumber(-398));
          _300_v31 = (_300_v31).Union(function () {
            let _coll16 = new _dafny.Set();
            for (const _compr_16 of (_300_v31).Elements) {
              let _338_v48 = _compr_16;
              if ((_300_v31).contains(_338_v48)) {
                _coll16.add((_338_v48).multipliedBy(new BigNumber(721)));
              }
            }
            return _coll16;
          }());
          _301_v32 = true;
          let _339_v49;
          let _nw50 = Array((new BigNumber(15)).toNumber()).fill(_module.D0.Default());
          _339_v49 = _nw50;
          let _index54 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_339_v49).length));
          (_339_v49)[_index54] = _306_v37;
          let _340_v50;
          _340_v50 = _module.D0.create_DC0(_302_v33);
          let _index55 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_339_v49).length));
          (_339_v49)[_index55] = _module.__default.fm16(_340_v50, globalState);
        } else if (_source7.is_DC12) {
          let _341___mcc_h9 = (_source7).cf26;
          let _342_cf26 = _341___mcc_h9;
          let _343_v51;
          _343_v51 = _dafny.Seq.UnicodeFromString("uhkhhuwaq");
          _343_v51 = _343_v51;
          (globalState).f1 = ((_dafny.ZERO).minus(p1)).multipliedBy(p1);
          let _index56 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_305_v36).length));
          (_305_v36)[_index56] = new BigNumber((_300_v31).length);
          let _index57 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_305_v36).length));
          let _rhs37 = (_this).f14;
          let _rhs38 = new BigNumber(((_dafny.MultiSet.fromElements(false)).Difference(((_this).f10).Difference(_dafny.MultiSet.fromElements(_module.__default.fm4(_300_v31, globalState))))).cardinality());
          let _lhs25 = _305_v36;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_305_v36).length));
          _301_v32 = _rhs37;
          _lhs25[_lhs26] = _rhs38;
          let _index58 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_305_v36).length));
          (_305_v36)[_index58] = p1;
        } else {
          let _344___mcc_h10 = (_source7).cf36;
          let _345_cf36 = _344___mcc_h10;
          let _346_v52;
          let _init7 = function (_347_i2) {
            return (_this).f14;
          };
          let _nw51 = Array((new BigNumber(22)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw51.length); _i0_7++) {
            _nw51[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _346_v52 = _nw51;
          (globalState).f5 = ((_301_v32) ? ((((_this).f15) ? (_346_v52) : (_346_v52))) : (_346_v52));
          (globalState).f1 = p1;
          let _348_v53;
          _348_v53 = _dafny.Seq.UnicodeFromString("jfclgutmi");
          _348_v53 = _348_v53;
          _301_v32 = (!(p1).isEqualTo(p1)) && ((_263_v1)[_module.__default.safeIndex(p1, new BigNumber((_263_v1).length))]);
        }
        let _349_v54;
        _349_v54 = _dafny.Seq.UnicodeFromString("iexqde");
        let _350_v55;
        _350_v55 = _dafny.MultiSet.fromElements(p1, p1, new BigNumber((_349_v54).length));
        let _351_v56;
        _351_v56 = _dafny.Set.fromElements((_this).f15);
        let _352_v57;
        _352_v57 = _dafny.Seq.of(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(p1, p1, (((_304_v35).contains(p1)) ? ((_304_v35).get(p1)) : (new BigNumber((_351_v56).length))), p1)).length), p1), new BigNumber((_351_v56).length), p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(globalState),_349_v54)).length), p1);
        _350_v55 = _dafny.MultiSet.FromArray(_352_v57);
      }
      (globalState).f1 = p1;
      let _353_v58;
      _353_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_263_v1);
      let _354_v59;
      _354_v59 = _dafny.Set.fromElements(p1, p1);
      let _355_v60;
      _355_v60 = new _dafny.CodePoint('i'.codePointAt(0));
      let _356_v61;
      _356_v61 = _dafny.Seq.UnicodeFromString("r");
      let _357_v62;
      _357_v62 = _dafny.MultiSet.fromElements(p1, p1);
      let _rhs39 = _module.__default.safeModuloInt(p1, p1);
      let _rhs40 = ((_module.__default.fm4(_354_v59, globalState)) ? (_dafny.Seq.Concat(_module.__default.fm15(_355_v60, p1, globalState), _263_v1)) : (_263_v1));
      let _rhs41 = (_353_v58).update((_this).f14, _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_263_v1, _module.__default.safeIndex(new BigNumber((_356_v61).length), new BigNumber((_263_v1).length)), (_this).f14), _module.__default.safeIndex(new BigNumber((_357_v62).cardinality()), new BigNumber((_dafny.Seq.update(_263_v1, _module.__default.safeIndex(new BigNumber((_356_v61).length), new BigNumber((_263_v1).length)), (_this).f14)).length)), (_this).f15), _263_v1));
      let _lhs27 = globalState;
      _lhs27.f1 = _rhs39;
      _263_v1 = _rhs40;
      _353_v58 = _rhs41;
      let _358_v63;
      _358_v63 = _dafny.Set.fromElements((_this).f15);
      let _hi4 = new BigNumber(427);
      for (let _359_i3 = new BigNumber(((_358_v63).Difference(_358_v63)).length); _359_i3.isLessThan(_hi4); _359_i3 = _359_i3.plus(_dafny.ONE)) {
        let _360_v64;
        let _nw52 = new _module.C0();
        _nw52.__ctor();
        _360_v64 = _nw52;
        let _361_v65;
        _361_v65 = _module.D4.create_DC10(_360_v64);
        let _362_v67;
        _362_v67 = _module.D6.create_DC16(function () {
  let _coll17 = new _dafny.Set();
  for (const _compr_17 of _dafny.IntegerRange(new BigNumber(490), new BigNumber(676))) {
    let _363_v66 = _compr_17;
    if (((new BigNumber(490)).isLessThanOrEqualTo(_363_v66)) && ((_363_v66).isLessThan(new BigNumber(676)))) {
      _coll17.add((_363_v66).minus(new BigNumber((_356_v61).length)));
    }
  }
  return _coll17;
}());
        _361_v65 = ((_module.__default.fm4((_362_v67).dtor_cf37, globalState)) ? (_361_v65) : (_361_v65));
        let _364_v68;
        let _nw53 = new _module.C0();
        _nw53.__ctor();
        _364_v68 = _nw53;
        if ((_this).f14) {
          _263_v1 = _263_v1;
          let _365_v69;
          let _nw54 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
          _365_v69 = _nw54;
          let _366_v70;
          _366_v70 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("ey"))).length),(_this).f10);
          let _index59 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_365_v69).length));
          (_365_v69)[_index59] = _366_v70;
          let _367_v71;
          _367_v71 = _dafny.Seq.of(p1, _359_i3);
          let _368_v72;
          _368_v72 = _module.D1.create_DC4((_this).f15, false, _359_i3, _367_v71, p1);
          let _369_v74;
          let _nw55 = Array((new BigNumber(14)).toNumber());
          _nw55[(_dafny.ZERO).toNumber()] = _module.__default.fm4(function () {
            let _coll18 = new _dafny.Set();
            for (const _compr_18 of _dafny.IntegerRange(new BigNumber(771), new BigNumber(-659))) {
              let _370_v73 = _compr_18;
              if (((new BigNumber(771)).isLessThanOrEqualTo(_370_v73)) && ((_370_v73).isLessThan(new BigNumber(-659)))) {
                _coll18.add((_370_v73).multipliedBy(_359_i3));
              }
            }
            return _coll18;
          }(), globalState);
          _nw55[(_dafny.ONE).toNumber()] = !((_this).f14);
          _nw55[(new BigNumber(2)).toNumber()] = !((_this).f14);
          _nw55[(new BigNumber(3)).toNumber()] = (_this).f15;
          _nw55[(new BigNumber(4)).toNumber()] = (_this).f14;
          _nw55[(new BigNumber(5)).toNumber()] = (_this).f15;
          _nw55[(new BigNumber(6)).toNumber()] = (_this).f14;
          _nw55[(new BigNumber(7)).toNumber()] = (_this).f15;
          _nw55[(new BigNumber(8)).toNumber()] = (_this).f15;
          _nw55[(new BigNumber(9)).toNumber()] = (_this).f14;
          _nw55[(new BigNumber(10)).toNumber()] = (_this).f14;
          _nw55[(new BigNumber(11)).toNumber()] = (_this).f15;
          _nw55[(new BigNumber(12)).toNumber()] = (_this).f14;
          _nw55[(new BigNumber(13)).toNumber()] = (_this).f14;
          _369_v74 = _nw55;
          let _371_v75;
          _371_v75 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f14),_369_v74);
          let _372_v76;
          _372_v76 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_368_v72).dtor_cf6,_369_v74), _371_v75, _dafny.Map.Empty.slice().updateUnsafe(!((_this).f15),_369_v74), _371_v75);
          let _index60 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_365_v69).length));
          (_365_v69)[_index60] = _module.__default.fm17(_dafny.MultiSet.fromElements(_355_v60), new BigNumber(((_372_v76)[_module.__default.safeIndex(_359_i3, new BigNumber((_372_v76).length))]).length), globalState);
          let _373_v77;
          _373_v77 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_dafny.Seq.UnicodeFromString("jtyk"), _356_v61),(_this).f14);
          _373_v77 = (_373_v77).update((_this).f15, !(p1).isEqualTo(_359_i3));
          let _374_v78;
          _374_v78 = false;
          _374_v78 = !((((_373_v77).contains((_this).f15)) ? ((_373_v77).get((_this).f15)) : ((_354_v59).IsSubsetOf(_354_v59))));
          let _rhs42 = p1;
          let _rhs43 = _263_v1;
          let _lhs28 = globalState;
          _lhs28.f1 = _rhs42;
          _263_v1 = _rhs43;
        } else {
          (globalState).f1 = new BigNumber((((_this).f10).Union((_dafny.MultiSet.FromArray(_263_v1)).Union((_this).f10))).cardinality());
          let _375_v79;
          let _nw56 = new _module.C0();
          _nw56.__ctor();
          _375_v79 = _nw56;
          let _376_v80;
          let _nw57 = new _module.C0();
          _nw57.__ctor();
          _376_v80 = _nw57;
          (globalState).f1 = (((_357_v62).contains(_359_i3)) ? ((_357_v62).get(_359_i3)) : (_359_i3));
          let _377_v82;
          let _init8 = ((_378_i3) => function (_379_i4) {
            return (_379_i4).multipliedBy(new BigNumber((function () {
              let _coll19 = new _dafny.Map();
              for (const _compr_19 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f15),_378_i3)).Keys.Elements) {
                let _380_v81 = _compr_19;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f15),_378_i3)).contains(_380_v81)) {
                  _coll19.push([_380_v81,(_this).f14]);
                }
              }
              return _coll19;
            }()).length));
          })(_359_i3);
          let _nw58 = Array((new BigNumber(10)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw58.length); _i0_8++) {
            _nw58[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _377_v82 = _nw58;
          let _381_v83;
          _381_v83 = _dafny.Map.Empty.slice().updateUnsafe(_359_i3,new BigNumber(((_this).f10).cardinality()));
          let _382_v84;
          _382_v84 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_381_v83).length));
          let _index61 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_377_v82).length));
          (_377_v82)[_index61] = new BigNumber(((_381_v83).update(p1, (((_381_v83).contains(_359_i3)) ? ((_381_v83).get(_359_i3)) : (new BigNumber((_381_v83).length))))).length);
          let _index62 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_377_v82).length));
          (_377_v82)[_index62] = ((_360_v64).fm6((_this).f15, p1, globalState)).multipliedBy(new BigNumber((_356_v61).length));
        }
        let _383_v85;
        _383_v85 = true;
        _383_v85 = !(((_383_v85) ? (true) : (_383_v85)));
      }
      let _384_v86;
      _384_v86 = _module.D0.create_DC0(_355_v60);
      if (function (_source8) {
        if (_source8.is_DC1) {
          let _385___mcc_h11 = (_source8).cf1;
          let _386___mcc_h12 = (_source8).cf2;
          let _387_cf2 = _386___mcc_h12;
          let _388_cf1 = _385___mcc_h11;
          return false;
        } else if (_source8.is_DC0) {
          let _389___mcc_h13 = (_source8).cf0;
          let _390_cf0 = _389___mcc_h13;
          return (_this).f15;
        } else {
          let _391___mcc_h14 = (_source8).cf3;
          let _392_cf3 = _391___mcc_h14;
          return true;
        }
      }(_384_v86)) {
        if (_dafny.areEqual(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wny"), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("q"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("q")).length)), _355_v60)), _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wny"), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("q"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("q")).length)), _355_v60))).length)), _355_v60), _356_v61)) {
          let _393_v87;
          let _nw59 = new _module.C0();
          _nw59.__ctor();
          _393_v87 = _nw59;
          let _394_v88;
          _394_v88 = _dafny.Map.Empty.slice().updateUnsafe(p1,_393_v87);
          let _395_v89;
          let _nw60 = Array((new BigNumber(28)).toNumber());
          _nw60[(_dafny.ZERO).toNumber()] = _393_v87;
          _nw60[(_dafny.ONE).toNumber()] = (((_this).f15) ? (_393_v87) : (_393_v87));
          _nw60[(new BigNumber(2)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(3)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(4)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(5)).toNumber()] = (((_394_v88).contains(p1)) ? ((_394_v88).get(p1)) : (_393_v87));
          _nw60[(new BigNumber(6)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(7)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(8)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(9)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(10)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(11)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(12)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(13)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(14)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(15)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(16)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(17)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(18)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(19)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(20)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(21)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(22)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(23)).toNumber()] = (((_this).f15) ? (_393_v87) : (_393_v87));
          _nw60[(new BigNumber(24)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(25)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(26)).toNumber()] = _393_v87;
          _nw60[(new BigNumber(27)).toNumber()] = _393_v87;
          _395_v89 = _nw60;
          _395_v89 = _395_v89;
          let _396_v90;
          _396_v90 = _dafny.Seq.of(p1, p1);
          let _397_v91;
          _397_v91 = _dafny.Seq.of(_358_v63, _358_v63, _358_v63, _358_v63, _358_v63);
          let _398_v94;
          let _nw61 = Array((new BigNumber(20)).toNumber());
          _nw61[(_dafny.ZERO).toNumber()] = (p1).multipliedBy(p1);
          _nw61[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw61[(new BigNumber(2)).toNumber()] = p1;
          _nw61[(new BigNumber(3)).toNumber()] = p1;
          _nw61[(new BigNumber(4)).toNumber()] = (_396_v90)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_396_v90).length))];
          _nw61[(new BigNumber(5)).toNumber()] = p1;
          _nw61[(new BigNumber(6)).toNumber()] = p1;
          _nw61[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_module.__default.fm16(_384_v86, globalState)).dtor_cf2);
          _nw61[(new BigNumber(8)).toNumber()] = new BigNumber(((_397_v91)[_module.__default.safeIndex(p1, new BigNumber((_397_v91).length))]).length);
          _nw61[(new BigNumber(9)).toNumber()] = p1;
          _nw61[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw61[(new BigNumber(11)).toNumber()] = new BigNumber((function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of _dafny.IntegerRange(new BigNumber(861), new BigNumber(-298))) {
              let _399_v92 = _compr_20;
              if (((new BigNumber(861)).isLessThanOrEqualTo(_399_v92)) && ((_399_v92).isLessThan(new BigNumber(-298)))) {
                _coll20.push([(_399_v92).plus(p1),(_module.D0.create_DC1((_this).f15, p1)).dtor_cf2]);
              }
            }
            return _coll20;
          }()).length);
          _nw61[(new BigNumber(12)).toNumber()] = p1;
          _nw61[(new BigNumber(13)).toNumber()] = p1;
          _nw61[(new BigNumber(14)).toNumber()] = p1;
          _nw61[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_this).f10).cardinality()));
          _nw61[(new BigNumber(16)).toNumber()] = new BigNumber((_356_v61).length);
          _nw61[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(p1, new BigNumber((_module.__default.fm18((_this).f15, function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of _dafny.IntegerRange(new BigNumber(830), new BigNumber(26))) {
              let _400_v93 = _compr_21;
              if (((new BigNumber(830)).isLessThanOrEqualTo(_400_v93)) && ((_400_v93).isLessThan(new BigNumber(26)))) {
                _coll21.push([(_400_v93).multipliedBy(new BigNumber((_358_v63).length)),(_this).f14]);
              }
            }
            return _coll21;
          }(), globalState)).length));
          _nw61[(new BigNumber(18)).toNumber()] = (p1).minus(p1);
          _nw61[(new BigNumber(19)).toNumber()] = (_393_v87).fm6((_this).f15, p1, globalState);
          _398_v94 = _nw61;
          let _index63 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_398_v94).length));
          (_398_v94)[_index63] = p1;
          let _index64 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_398_v94).length));
          let _rhs44 = _355_v60;
          let _rhs45 = (new BigNumber((_dafny.Seq.Concat(_396_v90, _396_v90)).length)).plus(new BigNumber(-840));
          let _rhs46 = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_this).f15) ? ((_dafny.ZERO).minus(new BigNumber(-26))) : (p1))));
          let _rhs47 = ((true) ? (p1) : (p1));
          let _rhs48 = p1;
          let _lhs29 = _398_v94;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_398_v94).length));
          let _lhs31 = globalState;
          let _lhs32 = globalState;
          let _lhs33 = globalState;
          _355_v60 = _rhs44;
          _lhs29[_lhs30] = _rhs45;
          _lhs31.f1 = _rhs46;
          _lhs32.f1 = _rhs47;
          _lhs33.f1 = _rhs48;
          let _401_v95;
          _401_v95 = false;
          let _402_v96;
          let _nw62 = Array((new BigNumber(21)).toNumber()).fill(false);
          _402_v96 = _nw62;
          let _403_v97;
          let _nw63 = Array((new BigNumber(13)).toNumber());
          _nw63[(_dafny.ZERO).toNumber()] = _402_v96;
          _nw63[(_dafny.ONE).toNumber()] = _402_v96;
          _nw63[(new BigNumber(2)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(3)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(4)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(5)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(6)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(7)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(8)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(9)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(10)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(11)).toNumber()] = _402_v96;
          _nw63[(new BigNumber(12)).toNumber()] = _402_v96;
          _403_v97 = _nw63;
          let _404_v98;
          let _init9 = ((_405_v60) => function (_406_i5) {
            return _405_v60;
          })(_355_v60);
          let _nw64 = Array((new BigNumber(6)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw64.length); _i0_9++) {
            _nw64[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _404_v98 = _nw64;
          let _407_v99;
          _407_v99 = _dafny.Seq.of(_404_v98, _404_v98, _404_v98);
          let _rhs49 = (_dafny.Set.fromElements((_this).f15, (_this).f14)).IsProperSubsetOf(_358_v63);
          let _rhs50 = p1;
          let _rhs51 = _403_v97;
          let _rhs52 = (_393_v87).fm5(_358_v63, new BigNumber((_407_v99).length), (p1).plus(p1), globalState);
          let _lhs34 = globalState;
          _401_v95 = _rhs49;
          _lhs34.f1 = _rhs50;
          _403_v97 = _rhs51;
          _401_v95 = _rhs52;
          let _408_v100;
          let _init10 = ((_409_v61) => function (_410_i6) {
            return _409_v61;
          })(_356_v61);
          let _nw65 = Array((new BigNumber(11)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw65.length); _i0_10++) {
            _nw65[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _408_v100 = _nw65;
          let _index65 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_403_v97).length));
          (_403_v97)[_index65] = _402_v96;
          let _index66 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_403_v97).length));
          let _rhs53 = _402_v96;
          let _rhs54 = _408_v100;
          let _rhs55 = _402_v96;
          let _lhs35 = globalState;
          let _lhs36 = _403_v97;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_403_v97).length));
          _lhs35.f5 = _rhs53;
          _408_v100 = _rhs54;
          _lhs36[_lhs37] = _rhs55;
        } else {
          let _411_v101;
          let _init11 = function (_412_i7) {
            return (_this).f15;
          };
          let _nw66 = Array((new BigNumber(20)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw66.length); _i0_11++) {
            _nw66[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _411_v101 = _nw66;
          let _413_v102;
          _413_v102 = _dafny.MultiSet.fromElements(_411_v101, _411_v101);
          (globalState).f1 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(p1, new BigNumber(((_413_v102).update(_411_v101, _module.__default.abs(p1))).cardinality())), p1);
          let _414_v103;
          _414_v103 = false;
          _414_v103 = (_this).f14;
          let _415_v104;
          _415_v104 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,_module.__default.fm19(_355_v60, (_this).f14, globalState))).length));
          let _416_v105;
          _416_v105 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _417_v106;
          _417_v106 = _dafny.Seq.of(p1);
          let _418_v107;
          let _nw67 = Array((new BigNumber(29)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = p1;
          _nw67[(_dafny.ONE).toNumber()] = p1;
          _nw67[(new BigNumber(2)).toNumber()] = new BigNumber(-308);
          _nw67[(new BigNumber(3)).toNumber()] = p1;
          _nw67[(new BigNumber(4)).toNumber()] = p1;
          _nw67[(new BigNumber(5)).toNumber()] = p1;
          _nw67[(new BigNumber(6)).toNumber()] = p1;
          _nw67[(new BigNumber(7)).toNumber()] = new BigNumber((_356_v61).length);
          _nw67[(new BigNumber(8)).toNumber()] = p1;
          _nw67[(new BigNumber(9)).toNumber()] = p1;
          _nw67[(new BigNumber(10)).toNumber()] = new BigNumber((_354_v59).length);
          _nw67[(new BigNumber(11)).toNumber()] = p1;
          _nw67[(new BigNumber(12)).toNumber()] = p1;
          _nw67[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.of((_this).f14)).length);
          _nw67[(new BigNumber(14)).toNumber()] = p1;
          _nw67[(new BigNumber(15)).toNumber()] = new BigNumber(((_this).f10).cardinality());
          _nw67[(new BigNumber(16)).toNumber()] = p1;
          _nw67[(new BigNumber(17)).toNumber()] = new BigNumber((_415_v104).length);
          _nw67[(new BigNumber(18)).toNumber()] = p1;
          _nw67[(new BigNumber(19)).toNumber()] = p1;
          _nw67[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(308)), ((_419_v60) => function (_420_i8) {
            return _419_v60;
          })(_355_v60))).length);
          _nw67[(new BigNumber(21)).toNumber()] = p1;
          _nw67[(new BigNumber(22)).toNumber()] = new BigNumber((_356_v61).length);
          _nw67[(new BigNumber(23)).toNumber()] = new BigNumber((_356_v61).length);
          _nw67[(new BigNumber(24)).toNumber()] = new BigNumber((_263_v1).length);
          _nw67[(new BigNumber(25)).toNumber()] = p1;
          _nw67[(new BigNumber(26)).toNumber()] = (((_416_v105).contains(p1)) ? ((_416_v105).get(p1)) : ((_dafny.ZERO).minus(_module.__default.fm19(_355_v60, (_this).f14, globalState))));
          _nw67[(new BigNumber(27)).toNumber()] = p1;
          _nw67[(new BigNumber(28)).toNumber()] = new BigNumber((_417_v106).length);
          _418_v107 = _nw67;
          let _421_v108;
          _421_v108 = _dafny.MultiSet.fromElements(_418_v107);
          let _422_v109;
          _422_v109 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(p1))).length),_421_v108);
          let _423_v110;
          _423_v110 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_418_v107);
          let _424_v111;
          let _nw68 = Array((new BigNumber(9)).toNumber());
          _nw68[(_dafny.ZERO).toNumber()] = _421_v108;
          _nw68[(_dafny.ONE).toNumber()] = _421_v108;
          _nw68[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_418_v107);
          _nw68[(new BigNumber(3)).toNumber()] = _421_v108;
          _nw68[(new BigNumber(4)).toNumber()] = (((_422_v109).contains(p1)) ? ((_422_v109).get(p1)) : (_421_v108));
          _nw68[(new BigNumber(5)).toNumber()] = _421_v108;
          _nw68[(new BigNumber(6)).toNumber()] = _421_v108;
          _nw68[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements((((_423_v110).contains(_414_v103)) ? ((_423_v110).get(_414_v103)) : (_418_v107)));
          _nw68[(new BigNumber(8)).toNumber()] = _421_v108;
          _424_v111 = _nw68;
          let _index67 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_424_v111).length));
          (_424_v111)[_index67] = (_421_v108).Union(_421_v108);
          let _index68 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_424_v111).length));
          (_424_v111)[_index68] = _421_v108;
          _356_v61 = _dafny.Seq.UnicodeFromString("t");
          (globalState).f1 = _module.__default.safeDivisionInt(new BigNumber(-725), new BigNumber((_415_v104).length));
        }
        let _425_v112;
        _425_v112 = true;
        let _426_v113;
        _426_v113 = _module.D6.create_DC17(p1, _356_v61, (_this).f14);
        _425_v112 = !(((((_263_v1)[_module.__default.safeIndex(p1, new BigNumber((_263_v1).length))]) ? (_426_v113) : (_426_v113))).dtor_cf40);
        let _427_v114;
        _427_v114 = _dafny.Map.Empty.slice().updateUnsafe((_263_v1)[_module.__default.safeIndex(p1, new BigNumber((_263_v1).length))],p1);
        _427_v114 = _dafny.Map.Empty.slice().updateUnsafe(!(_425_v112),(p1).minus(p1));
        let _428_v116;
        _428_v116 = _module.D5.create_DC12(_358_v63);
        let _rhs56 = _module.__default.fm19(_355_v60, (_this).f15, globalState);
        let _rhs57 = p1;
        let _rhs58 = (_dafny.ZERO).minus(new BigNumber((function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of ((_dafny.Map.Empty.slice().updateUnsafe(p1,_428_v116)).update(p1, _428_v116)).Keys.Elements) {
            let _429_v115 = _compr_22;
            if (((_dafny.Map.Empty.slice().updateUnsafe(p1,_428_v116)).update(p1, _428_v116)).contains(_429_v115)) {
              _coll22.push([_module.__default.safeModuloInt(_429_v115, p1),((_425_v112) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_263_v1).length),(_this).f14)) : (_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f14)))]);
            }
          }
          return _coll22;
        }()).length));
        let _rhs59 = (((_this).f14) ? (_dafny.Seq.Concat(_356_v61, _dafny.Seq.Create(_module.__default.abs(new BigNumber(156)), function (_430_i9) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        }))) : (_356_v61));
        let _lhs38 = globalState;
        let _lhs39 = globalState;
        let _lhs40 = globalState;
        _lhs38.f1 = _rhs56;
        _lhs39.f1 = _rhs57;
        _lhs40.f1 = _rhs58;
        _356_v61 = _rhs59;
        if ((_this).f15) {
          _425_v112 = _425_v112;
          _425_v112 = (_this).f15;
          let _431_v117;
          let _432_v118;
          let _433_v119;
          let _434_v120;
          let _out20;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector5 = _module.__default.m0(_dafny.Seq.Concat(_dafny.Seq.update(_263_v1, _module.__default.safeIndex(new BigNumber(9), new BigNumber((_263_v1).length)), (_this).f15), _263_v1), _263_v1, globalState);
          _out20 = _outcollector5[0];
          _out21 = _outcollector5[1];
          _out22 = _outcollector5[2];
          _out23 = _outcollector5[3];
          _431_v117 = _out20;
          _432_v118 = _out21;
          _433_v119 = _out22;
          _434_v120 = _out23;
          _425_v112 = _431_v117;
          let _435_v121;
          _435_v121 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(381))).cardinality()),(_this).f14);
          let _436_v122;
          _436_v122 = _dafny.MultiSet.fromElements(_module.D0.create_DC1((_this).f15, new BigNumber((_435_v121).length)));
          let _437_v123;
          _437_v123 = _module.D0.create_DC1((_this).f14, p1);
          (globalState).f1 = _module.__default.safeDivisionInt(_434_v120, (((_436_v122).contains(_437_v123)) ? ((_436_v122).get(_437_v123)) : (_434_v120)));
        } else {
          let _438_v124;
          let _nw69 = new _module.C0();
          _nw69.__ctor();
          _438_v124 = _nw69;
          let _439_v125;
          let _nw70 = Array((new BigNumber(21)).toNumber()).fill([]);
          _439_v125 = _nw70;
          let _440_v126;
          let _nw71 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _440_v126 = _nw71;
          let _index69 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_439_v125).length));
          (_439_v125)[_index69] = _440_v126;
          let _441_v128;
          _441_v128 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _442_v129;
          _442_v129 = _dafny.Map.Empty.slice().updateUnsafe((function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of (_441_v128).Keys.Elements) {
              let _443_v127 = _compr_23;
              if ((_441_v128).contains(_443_v127)) {
                _coll23.push([_module.__default.safeModuloInt(_443_v127, p1),p1]);
              }
            }
            return _coll23;
          }()).Merge(_441_v128),_440_v126);
          let _index70 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_439_v125).length));
          (_439_v125)[_index70] = (((_442_v129).contains(_441_v128)) ? ((_442_v129).get(_441_v128)) : (_440_v126));
          let _444_v130;
          _444_v130 = _dafny.Seq.of(_356_v61);
          let _445_v131;
          _445_v131 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_444_v130)[_module.__default.safeIndex(p1, new BigNumber((_444_v130).length))]);
          _356_v61 = (((_445_v131).contains((_this).f15)) ? ((_445_v131).get((_this).f15)) : (_356_v61));
          _425_v112 = (_this).f15;
          _425_v112 = (_this).f14;
        }
      } else {
        let _446_v132;
        let _init12 = ((_447_v61) => function (_448_i10) {
          return _module.__default.safeDivisionInt(_448_i10, new BigNumber((_447_v61).length));
        })(_356_v61);
        let _nw72 = Array((new BigNumber(27)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw72.length); _i0_12++) {
          _nw72[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _446_v132 = _nw72;
        _446_v132 = _446_v132;
        let _449_v133;
        _449_v133 = true;
        let _450_v134;
        _450_v134 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f10).cardinality()),false);
        let _451_v135;
        _451_v135 = _dafny.Map.Empty.slice().updateUnsafe(_450_v134,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length)));
        let _452_v136;
        let _init13 = ((_453_v1, _454_p1) => function (_455_i11) {
          return (_453_v1)[_module.__default.safeIndex(_454_p1, new BigNumber((_453_v1).length))];
        })(_263_v1, p1);
        let _nw73 = Array((new BigNumber(21)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw73.length); _i0_13++) {
          _nw73[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _452_v136 = _nw73;
        let _456_v137;
        _456_v137 = _dafny.MultiSet.fromElements(_452_v136);
        _449_v133 = (new BigNumber(((_451_v135).update(_dafny.Map.Empty.slice().updateUnsafe(p1,_449_v133), (((_456_v137).contains(_452_v136)) ? ((_456_v137).get(_452_v136)) : ((_dafny.ZERO).minus(p1))))).length)).isLessThanOrEqualTo(p1);
        _449_v133 = !((_this).f14) || ((_this).f15);
        _449_v133 = (_this).f14;
        if ((_this).f15) {
          _449_v133 = (_this).f15;
          (globalState).f1 = ((new BigNumber(549)).minus(p1)).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), p1));
          let _457_v138;
          _457_v138 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
          let _458_v139;
          _458_v139 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_457_v138).length),(_357_v62).Difference(_dafny.MultiSet.fromElements(p1)));
          let _459_v140;
          _459_v140 = (_this).f10;
          let _460_v141;
          let _nw74 = new _module.C0();
          _nw74.__ctor();
          _460_v141 = _nw74;
          let _rhs60 = (p1).multipliedBy(new BigNumber((_450_v134).length));
          let _rhs61 = _dafny.Map.Empty.slice().updateUnsafe(p1,_357_v62);
          let _rhs62 = !(((_this).f10)).equals((_this).f10);
          let _rhs63 = _460_v141;
          let _rhs64 = ((new BigNumber(((_this).f10).cardinality())).multipliedBy(p1)).multipliedBy((_dafny.ZERO).minus(p1));
          let _lhs41 = globalState;
          let _lhs42 = globalState;
          _lhs41.f1 = _rhs60;
          _458_v139 = _rhs61;
          _449_v133 = _rhs62;
          r2 = _rhs63;
          _lhs42.f1 = _rhs64;
          let _461_v142;
          let _nw75 = Array((new BigNumber(26)).toNumber());
          _461_v142 = _nw75;
          _461_v142 = _461_v142;
          let _462_v143;
          _462_v143 = _module.D1.create_DC4((_this).f14, false, p1, _module.__default.fm20(p1, new BigNumber(713), globalState), new BigNumber(-757));
          let _index71 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_452_v136).length));
          (_452_v136)[_index71] = (_462_v143).dtor_cf6;
          let _463_v144;
          _463_v144 = _dafny.Map.Empty.slice().updateUnsafe(_449_v133,p1);
          let _index72 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_452_v136).length));
          (_452_v136)[_index72] = !((((_450_v134).contains(new BigNumber((_463_v144).length))) ? ((_450_v134).get(new BigNumber((_463_v144).length))) : ((_this).f15))) || (((!(!((_263_v1)[_module.__default.safeIndex(p1, new BigNumber((_263_v1).length))]))) ? (!(!(!(true)))) : ((_this).f15)));
        } else {
          let _464_v145;
          let _nw76 = Array((new BigNumber(3)).toNumber());
          _nw76[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_452_v136, _452_v136, _452_v136);
          _nw76[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(_452_v136, _452_v136);
          _nw76[(new BigNumber(2)).toNumber()] = _456_v137;
          _464_v145 = _nw76;
          let _index73 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_464_v145).length));
          (_464_v145)[_index73] = (_456_v137).Difference(_456_v137);
          let _index74 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_464_v145).length));
          (_464_v145)[_index74] = _456_v137;
          let _465_v146;
          _465_v146 = _module.D1.create_DC4(_449_v133, (_this).f14, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(81)), function (_466_i12) {
  return new BigNumber(-93);
}), p1);
          let _index75 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_446_v132).length));
          (_446_v132)[_index75] = (_465_v146).dtor_cf9;
          let _index76 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_446_v132).length));
          (_446_v132)[_index76] = _module.__default.safeDivisionInt(p1, p1);
          _449_v133 = !((_this).f14);
          let _467_v147;
          let _nw77 = new _module.C0();
          _nw77.__ctor();
          _467_v147 = _nw77;
          let _468_v148;
          let _nw78 = new _module.C0();
          _nw78.__ctor();
          _468_v148 = _nw78;
        }
      }
      if ((_this).f15) {
        if ((_263_v1)[_module.__default.safeIndex(p1, new BigNumber((_263_v1).length))]) {
          let _469_v149;
          _469_v149 = false;
          _469_v149 = (p1).isLessThanOrEqualTo((p1).multipliedBy(p1));
          let _470_v150;
          let _nw79 = new _module.C0();
          _nw79.__ctor();
          _470_v150 = _nw79;
          (globalState).f1 = p1;
          let _471_v151;
          let _nw80 = new _module.C0();
          _nw80.__ctor();
          _471_v151 = _nw80;
          _356_v61 = _dafny.Seq.Concat(_356_v61, _356_v61);
        } else {
          let _472_v152;
          let _init14 = function (_473_i13) {
            return _module.__default.safeModuloInt(_473_i13, new BigNumber(86));
          };
          let _nw81 = Array((new BigNumber(25)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw81.length); _i0_14++) {
            _nw81[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _472_v152 = _nw81;
          let _index77 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_472_v152).length));
          (_472_v152)[_index77] = _module.__default.safeDivisionInt(_module.__default.fm19(_355_v60, (_this).f15, globalState), (_dafny.ZERO).minus(p1));
          let _index78 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_472_v152).length));
          (_472_v152)[_index78] = _module.__default.fm19(_355_v60, (_this).f15, globalState);
          let _474_v153;
          let _nw82 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
          _474_v153 = _nw82;
          let _index79 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_474_v153).length));
          (_474_v153)[_index79] = _dafny.Seq.of(!((_this).f14));
          let _475_v154;
          _475_v154 = true;
          let _index80 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_474_v153).length));
          let _rhs65 = _263_v1;
          let _rhs66 = _module.__default.fm4(_354_v59, globalState);
          let _rhs67 = !((p1).isLessThan(_module.__default.fm19(_355_v60, (_this).f14, globalState)));
          let _lhs43 = _474_v153;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_474_v153).length));
          _lhs43[_lhs44] = _rhs65;
          _475_v154 = _rhs66;
          _475_v154 = _rhs67;
          let _476_v155;
          let _nw83 = new _module.C0();
          _nw83.__ctor();
          _476_v155 = _nw83;
          let _477_v156;
          let _nw84 = new _module.C0();
          _nw84.__ctor();
          _477_v156 = _nw84;
          _356_v61 = _module.__default.fm2(_475_v154, (_472_v152)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_472_v152).length))], _355_v60, (_472_v152)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_472_v152).length))], globalState);
        }
        let _478_v157;
        let _nw85 = Array((new BigNumber(29)).toNumber()).fill(false);
        _478_v157 = _nw85;
        let _index81 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_478_v157).length));
        (_478_v157)[_index81] = (_this).f15;
        let _index82 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_478_v157).length));
        (_478_v157)[_index82] = false;
        let _479_v158;
        let _init15 = ((_480_v60) => function (_481_i14) {
          return _480_v60;
        })(_355_v60);
        let _nw86 = Array((new BigNumber(26)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw86.length); _i0_15++) {
          _nw86[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _479_v158 = _nw86;
        let _index83 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_479_v158).length));
        (_479_v158)[_index83] = _355_v60;
        let _index84 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_479_v158).length));
        (_479_v158)[_index84] = _module.__default.fm1(globalState);
        let _482_v159;
        let _nw87 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _482_v159 = _nw87;
        let _483_v160;
        _483_v160 = _dafny.Map.Empty.slice().updateUnsafe(_482_v159,p1);
        _483_v160 = ((_483_v160).Merge(_483_v160)).Merge(_483_v160);
        let _484_v161;
        _484_v161 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f15);
        let _485_v163;
        _485_v163 = _module.D8.create_DC19(_484_v161);
        let _486_v164;
        _486_v164 = _dafny.Set.fromElements(_484_v161, _484_v161, (function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of _dafny.IntegerRange(new BigNumber(976), new BigNumber(896))) {
            let _487_v162 = _compr_24;
            if (((new BigNumber(976)).isLessThanOrEqualTo(_487_v162)) && ((_487_v162).isLessThan(new BigNumber(896)))) {
              _coll24.push([(_487_v162).minus(p1),true]);
            }
          }
          return _coll24;
        }()).Merge((_485_v163).dtor_cf42), _484_v161);
        let _488_v165;
        _488_v165 = _dafny.Seq.of(_486_v164);
        _486_v164 = (_486_v164).Union((_488_v165)[_module.__default.safeIndex(p1, new BigNumber((_488_v165).length))]);
      } else {
        if (_module.__default.fm4((function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of _dafny.IntegerRange(new BigNumber(52), new BigNumber(-560))) {
            let _489_v166 = _compr_25;
            if (((new BigNumber(52)).isLessThanOrEqualTo(_489_v166)) && ((_489_v166).isLessThan(new BigNumber(-560)))) {
              _coll25.add(_module.__default.safeModuloInt(_489_v166, new BigNumber((_dafny.Seq.of(p1, p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(385)), ((_490_v1) => function (_491_i15) {
                return _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_dafny.MultiSet.FromArray(_490_v1));
              })(_263_v1))).length))).length))).length)));
            }
          }
          return _coll25;
        }()).Difference(_354_v59), globalState)) {
          (globalState).f1 = p1;
          let _492_v167;
          let _nw88 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _492_v167 = _nw88;
          let _index85 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_492_v167).length));
          (_492_v167)[_index85] = _dafny.Seq.update(_dafny.Seq.Concat(_356_v61, _356_v61), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_356_v61, _356_v61)).length)), _355_v60);
          let _index86 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_492_v167).length));
          (_492_v167)[_index86] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(285)), ((_493_v60) => function (_494_i16) {
            return _493_v60;
          })(_355_v60)), _dafny.Seq.update(_356_v61, _module.__default.safeIndex(p1, new BigNumber((_356_v61).length)), _355_v60));
          let _495_v168;
          _495_v168 = true;
          _495_v168 = (_this).f14;
          (globalState).f1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p1));
          let _496_v169;
          _496_v169 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_263_v1, _263_v1), _module.__default.safeIndex(new BigNumber((_356_v61).length), new BigNumber((_dafny.Seq.Concat(_263_v1, _263_v1)).length)), (_this).f14),(p1).multipliedBy(p1));
          let _497_v170;
          _497_v170 = _module.D3.create_DC7((_492_v167)[_module.__default.safeIndex(new BigNumber(416), new BigNumber((_492_v167).length))]);
          let _498_v171;
          _498_v171 = _dafny.Map.Empty.slice().updateUnsafe(_497_v170,false);
          let _499_v172;
          _499_v172 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,!((_this).f14));
          let _500_v173;
          _500_v173 = _dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC1((_this).f14, p1)).dtor_cf1,_499_v172);
          let _rhs68 = (((_496_v169).contains(_263_v1)) ? ((_496_v169).get(_263_v1)) : (p1));
          let _rhs69 = (((((_500_v173).contains(_495_v168)) ? ((_500_v173).get(_495_v168)) : (_499_v172))).Merge(_499_v172)).contains((((_498_v171).contains(_497_v170)) ? ((_498_v171).get(_497_v170)) : ((((_499_v172).contains((_this).f15)) ? ((_499_v172).get((_this).f15)) : ((_this).f14)))));
          let _rhs70 = _module.__default.safeDivisionInt(p1, _module.__default.fm19(_355_v60, (_this).f15, globalState));
          let _lhs45 = globalState;
          let _lhs46 = globalState;
          _lhs45.f1 = _rhs68;
          _495_v168 = _rhs69;
          _lhs46.f1 = _rhs70;
        } else {
          let _501_v174;
          _501_v174 = (_this).f10;
          _501_v174 = _dafny.MultiSet.fromElements((_263_v1)[_module.__default.safeIndex(p1, new BigNumber((_263_v1).length))]);
          let _502_v175;
          let _nw89 = Array((new BigNumber(3)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _384_v86;
          _nw89[(_dafny.ONE).toNumber()] = _module.__default.fm21(globalState);
          _nw89[(new BigNumber(2)).toNumber()] = _384_v86;
          _502_v175 = _nw89;
          let _503_v176;
          _503_v176 = _dafny.Map.Empty.slice().updateUnsafe(_502_v175,(_263_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ale")).length), new BigNumber((_263_v1).length))]);
          let _504_v177;
          _504_v177 = _dafny.Seq.of(_354_v59, _354_v59, _354_v59);
          let _505_v178;
          _505_v178 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_504_v177)[_module.__default.safeIndex(p1, new BigNumber((_504_v177).length))]).length),_263_v1);
          let _506_v179;
          _506_v179 = _dafny.Map.Empty.slice().updateUnsafe(_505_v178,_502_v175);
          _503_v176 = (_503_v176).update((((_506_v179).contains(_505_v178)) ? ((_506_v179).get(_505_v178)) : (_502_v175)), _dafny.areEqual(_356_v61, _356_v61));
          let _507_v180;
          _507_v180 = false;
          let _508_v181;
          let _nw90 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _508_v181 = _nw90;
          let _index87 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length));
          (_508_v181)[_index87] = p1;
          let _index88 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length));
          let _rhs71 = false;
          let _rhs72 = p1;
          let _lhs47 = _508_v181;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length));
          _507_v180 = _rhs71;
          _lhs47[_lhs48] = _rhs72;
          let _index89 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length));
          (_508_v181)[_index89] = _module.__default.safeDivisionInt(((_508_v181)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length))]).multipliedBy((_508_v181)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length))]), _module.__default.fm19(_module.__default.fm1(globalState), (_this).f15, globalState));
          let _509_v182;
          _509_v182 = _module.D1.create_DC5(new BigNumber(-460), (_this).f15, (_508_v181)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length))], (_this).f15, p1);
          let _510_v184;
          _510_v184 = _dafny.Seq.of((_508_v181)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length))]);
          let _511_v185;
          _511_v185 = _dafny.Seq.of(_509_v182, _509_v182, _509_v182, _module.D1.create_DC5(new BigNumber((_358_v63).length), false, p1, true, new BigNumber((_510_v184).length)));
          let _512_v188;
          _512_v188 = _dafny.Set.fromElements(_509_v182);
          let _513_v189;
          _513_v189 = _dafny.Map.Empty.slice().updateUnsafe(_507_v180,_module.__default.safeModuloInt(p1, p1));
          let _514_v190;
          let _nw91 = Array((new BigNumber(21)).toNumber()).fill(false);
          _514_v190 = _nw91;
          let _515_v191;
          _515_v191 = _module.D1.create_DC3(_514_v190);
          let _516_v192;
          _516_v192 = _dafny.Set.fromElements(_515_v191, _515_v191);
          let _517_v193;
          _517_v193 = _dafny.Map.Empty.slice().updateUnsafe(_516_v192,p1);
          let _518_v194;
          _518_v194 = _dafny.Seq.of(_517_v193, _517_v193, (_dafny.Map.Empty.slice().updateUnsafe(_516_v192,new BigNumber((_263_v1).length))).update(_dafny.Set.fromElements(_515_v191), _module.__default.fm19(_355_v60, (_this).f15, globalState)));
          let _index90 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length));
          let _index91 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length));
          let _rhs73 = ((_dafny.Set.fromElements(_509_v182, _module.D1.create_DC5(_module.__default.fm19(_355_v60, (_this).f15, globalState), false, (_508_v181)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length))], (_this).f14, new BigNumber((_354_v59).length)))).Difference(function () {
            let _coll26 = new _dafny.Set();
            for (const _compr_26 of (function () {
              let _coll27 = new _dafny.Map();
              for (const _compr_27 of (function () {
                let _coll28 = new _dafny.Set();
                for (const _compr_28 of (_511_v185).Elements) {
                  let _519_v186 = _compr_28;
                  if (_dafny.Seq.contains(_511_v185, _519_v186)) {
                    _coll28.add(_519_v186);
                  }
                }
                return _coll28;
              }()).Elements) {
                let _520_v183 = _compr_27;
                if ((function () {
                  let _coll29 = new _dafny.Set();
                  for (const _compr_29 of (_511_v185).Elements) {
                    let _521_v186 = _compr_29;
                    if (_dafny.Seq.contains(_511_v185, _521_v186)) {
                      _coll29.add(_521_v186);
                    }
                  }
                  return _coll29;
                }()).contains(_520_v183)) {
                  _coll27.push([_520_v183,(_this).f15]);
                }
              }
              return _coll27;
            }()).Keys.Elements) {
              let _522_v187 = _compr_26;
              if ((function () {
                let _coll30 = new _dafny.Map();
                for (const _compr_30 of (function () {
                  let _coll31 = new _dafny.Set();
                  for (const _compr_31 of (_511_v185).Elements) {
                    let _523_v186 = _compr_31;
                    if (_dafny.Seq.contains(_511_v185, _523_v186)) {
                      _coll31.add(_523_v186);
                    }
                  }
                  return _coll31;
                }()).Elements) {
                  let _520_v183 = _compr_30;
                  if ((function () {
                    let _coll32 = new _dafny.Set();
                    for (const _compr_32 of (_511_v185).Elements) {
                      let _524_v186 = _compr_32;
                      if (_dafny.Seq.contains(_511_v185, _524_v186)) {
                        _coll32.add(_524_v186);
                      }
                    }
                    return _coll32;
                  }()).contains(_520_v183)) {
                    _coll30.push([_520_v183,(_this).f15]);
                  }
                }
                return _coll30;
              }()).contains(_522_v187)) {
                _coll26.add(_522_v187);
              }
            }
            return _coll26;
          }())).IsDisjointFrom(_512_v188);
          let _rhs74 = (((_513_v189).contains((_this).f15)) ? ((_513_v189).get((_this).f15)) : (p1));
          let _rhs75 = !(true);
          let _rhs76 = new BigNumber(((_518_v194)[_module.__default.safeIndex((_508_v181)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length))], new BigNumber((_518_v194).length))]).length);
          let _lhs49 = _508_v181;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length));
          let _lhs51 = _508_v181;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_508_v181).length));
          _507_v180 = _rhs73;
          _lhs49[_lhs50] = _rhs74;
          _507_v180 = _rhs75;
          _lhs51[_lhs52] = _rhs76;
        }
        let _525_v195;
        _525_v195 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("nthxycdvu"));
        let _526_v196;
        _526_v196 = _dafny.Map.Empty.slice().updateUnsafe(p1,_525_v195);
        let _527_v197;
        let _nw92 = Array((new BigNumber(9)).toNumber());
        _nw92[(_dafny.ZERO).toNumber()] = _525_v195;
        _nw92[(_dafny.ONE).toNumber()] = _525_v195;
        _nw92[(new BigNumber(2)).toNumber()] = _525_v195;
        _nw92[(new BigNumber(3)).toNumber()] = _525_v195;
        _nw92[(new BigNumber(4)).toNumber()] = _525_v195;
        _nw92[(new BigNumber(5)).toNumber()] = (((_526_v196).contains(p1)) ? ((_526_v196).get(p1)) : (_525_v195));
        _nw92[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(_356_v61, _356_v61);
        _nw92[(new BigNumber(7)).toNumber()] = (((_526_v196).contains(p1)) ? ((_526_v196).get(p1)) : (_525_v195));
        _nw92[(new BigNumber(8)).toNumber()] = (_525_v195).Difference(_525_v195);
        _527_v197 = _nw92;
        let _index92 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_527_v197).length));
        (_527_v197)[_index92] = (_dafny.Set.fromElements(_356_v61)).Difference(_525_v195);
        let _index93 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_527_v197).length));
        (_527_v197)[_index93] = _module.__default.fm22(globalState);
        let _528_v198;
        _528_v198 = false;
        _528_v198 = !(((p1).isEqualTo(p1)) || (_dafny.Seq.contains(_263_v1, _528_v198)));
        let _529_v199;
        let _nw93 = new _module.C0();
        _nw93.__ctor();
        _529_v199 = _nw93;
        _529_v199 = _529_v199;
        let _530_v200;
        let _init16 = ((_531_v1, _532_v61) => function (_533_i17) {
          return (_531_v1)[_module.__default.safeIndex(new BigNumber((_532_v61).length), new BigNumber((_531_v1).length))];
        })(_263_v1, _356_v61);
        let _nw94 = Array((new BigNumber(28)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw94.length); _i0_16++) {
          _nw94[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _530_v200 = _nw94;
        let _index94 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_530_v200).length));
        (_530_v200)[_index94] = _528_v198;
        let _534_v201;
        _534_v201 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_module.__default.fm19(_355_v60, true, globalState), new BigNumber(824)),!((_this).f14));
        let _index95 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_530_v200).length));
        (_530_v200)[_index95] = !((((_534_v201).contains((_dafny.ZERO).minus(_module.__default.fm19(_355_v60, _528_v198, globalState)))) ? ((_534_v201).get((_dafny.ZERO).minus(_module.__default.fm19(_355_v60, _528_v198, globalState)))) : (_528_v198)));
      }
      let _535_v202;
      let _init17 = ((_536_p1) => function (_537_i18) {
        return (_536_p1).isLessThan((_dafny.ZERO).minus(_536_p1));
      })(p1);
      let _nw95 = Array((new BigNumber(3)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw95.length); _i0_17++) {
        _nw95[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _535_v202 = _nw95;
      r0 = _535_v202;
      let _538_v204;
      _538_v204 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f14);
      r1 = (function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of _dafny.IntegerRange(new BigNumber(171), new BigNumber(252))) {
          let _539_v203 = _compr_33;
          if (((new BigNumber(171)).isLessThanOrEqualTo(_539_v203)) && ((_539_v203).isLessThan(new BigNumber(252)))) {
            _coll33.push([(_539_v203).plus(new BigNumber(429)),(_this).f14]);
          }
        }
        return _coll33;
      }()).Merge(_538_v204);
      let _540_v205;
      let _nw96 = new _module.C0();
      _nw96.__ctor();
      _540_v205 = _nw96;
      r2 = _540_v205;
      return [r0, r1, r2];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f10 = _dafny.MultiSet.Empty;
      this.f12 = false;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f12, f13, f10) {
      let _this = this;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this)._f10 = f10;
      return;
    }
    fm10(p0, p1, globalState) {
      let _this = this;
      return new _dafny.CodePoint('n'.codePointAt(0));
    };
    fm11(p0, globalState) {
      let _this = this;
      if (!(_this.f12)) {
        return (_this).f13;
      } else {
        return (_this).f13;
      }
    };
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _541_v0;
      _541_v0 = _dafny.Seq.UnicodeFromString("fsvmc");
      (globalState).f1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-614)), function (_542_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }), _541_v0)).length);
      let _543_v1;
      let _nw97 = Array((new BigNumber(14)).toNumber()).fill(false);
      _543_v1 = _nw97;
      let _544_v2;
      _544_v2 = _dafny.Map.Empty.slice().updateUnsafe(_543_v1,_module.__default.fm12(globalState));
      let _545_v3;
      _545_v3 = _dafny.Seq.of(!(_this.f12));
      _544_v2 = (_544_v2).update(_543_v1, _545_v3);
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_543_v1).length))) {
        let _546_i1 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_546_i1)) && ((_546_i1).isLessThan(new BigNumber((_543_v1).length))))) {
          (_543_v1)[(_546_i1)] = _this.f12;
        }
      }
      let _547_i2;
      _547_i2 = _dafny.ZERO;
      L2: {
        while (_this.f12) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_547_i2)) {
              break L2;
            }
            _547_i2 = (_547_i2).plus(_dafny.ONE);
            let _548_v4;
            _548_v4 = _dafny.Set.fromElements(new BigNumber((_545_v3).length), p1);
            let _549_v5;
            _549_v5 = _dafny.Set.fromElements(new BigNumber((_548_v4).length), (_this).f13);
            if ((_dafny.Set.fromElements(p0, p1)).IsDisjointFrom(_549_v5)) {
              let _index96 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_543_v1).length));
              (_543_v1)[_index96] = _this.f12;
              let _550_v6;
              _550_v6 = _dafny.Set.fromElements(_module.__default.fm4(_549_v5, globalState));
              let _index97 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_543_v1).length));
              (_543_v1)[_index97] = (_dafny.Set.fromElements(_module.__default.fm4(_dafny.Set.fromElements(p1, p1, p1), globalState), _this.f12, _this.f12, _this.f12)).IsDisjointFrom(_550_v6);
              let _551_v7;
              _551_v7 = _module.D0.create_DC1((_543_v1)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_543_v1).length))], p0);
              let _552_v8;
              _552_v8 = _module.D0.create_DC2(_551_v7);
              let _553_v9;
              _553_v9 = _dafny.MultiSet.fromElements(_552_v8);
              let _554_v10;
              _554_v10 = _dafny.Map.Empty.slice().updateUnsafe(_553_v9,p2);
              _554_v10 = (_554_v10).update(_553_v9, p2);
              let _555_v11;
              _555_v11 = _dafny.MultiSet.fromElements(p1, p1);
              let _556_v12;
              _556_v12 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(p1, p1, (_this).fm11(_dafny.MultiSet.FromArray(p2), globalState), (_this).f13)).Intersect(_549_v5),((_this).f10).update(_this.f12, _module.__default.abs(new BigNumber((_555_v11).cardinality()))));
              _556_v12 = (_556_v12).update((_dafny.Set.fromElements(p1, p0)).Difference(_548_v4), ((_this).f10).update((_543_v1)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_543_v1).length))], _module.__default.abs(p0)));
              let _557_v13;
              _557_v13 = new _dafny.CodePoint('t'.codePointAt(0));
              let _558_v14;
              _558_v14 = _dafny.Seq.of((_555_v11).update((_this).f13, _module.__default.abs((_this).f13)), _dafny.MultiSet.FromArray(p2), _555_v11);
              let _559_v15;
              let _nw98 = Array((new BigNumber(17)).toNumber());
              _nw98[(_dafny.ZERO).toNumber()] = _541_v0;
              _nw98[(_dafny.ONE).toNumber()] = _541_v0;
              _nw98[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(450)), function (_560_i3) {
                return new _dafny.CodePoint('q'.codePointAt(0));
              });
              _nw98[(new BigNumber(3)).toNumber()] = _module.__default.fm2((_543_v1)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_543_v1).length))], p0, _557_v13, new BigNumber(((_558_v14)[_module.__default.safeIndex(p1, new BigNumber((_558_v14).length))]).cardinality()), globalState);
              _nw98[(new BigNumber(4)).toNumber()] = _541_v0;
              _nw98[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_541_v0, _module.__default.safeIndex(new BigNumber((_541_v0).length), new BigNumber((_541_v0).length)), _557_v13);
              _nw98[(new BigNumber(6)).toNumber()] = _541_v0;
              _nw98[(new BigNumber(7)).toNumber()] = _541_v0;
              _nw98[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_541_v0, _dafny.Seq.UnicodeFromString("qmb"));
              _nw98[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jqnofwwqn"), _541_v0);
              _nw98[(new BigNumber(10)).toNumber()] = _541_v0;
              _nw98[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("majd");
              _nw98[(new BigNumber(12)).toNumber()] = _541_v0;
              _nw98[(new BigNumber(13)).toNumber()] = _541_v0;
              _nw98[(new BigNumber(14)).toNumber()] = _541_v0;
              _nw98[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_541_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-611)), ((_561_v13) => function (_562_i4) {
                return _561_v13;
              })(_557_v13)));
              _nw98[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-701)), ((_563_v13) => function (_564_i5) {
                return _563_v13;
              })(_557_v13));
              _559_v15 = _nw98;
              let _index98 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_559_v15).length));
              (_559_v15)[_index98] = _541_v0;
              let _index99 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_559_v15).length));
              (_559_v15)[_index99] = _541_v0;
              let _565_v16;
              let _nw99 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
              _565_v16 = _nw99;
              let _index100 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_565_v16).length));
              (_565_v16)[_index100] = (((_543_v1)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_543_v1).length))]) ? ((p2)[_module.__default.safeIndex((_this).f13, new BigNumber((p2).length))]) : (p1));
              let _index101 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_543_v1).length));
              let _index102 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_565_v16).length));
              let _index103 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_559_v15).length));
              let _rhs77 = (_this).f13;
              let _rhs78 = _this.f12;
              let _rhs79 = (p0).multipliedBy(p0);
              let _rhs80 = (_559_v15)[_module.__default.safeIndex(new BigNumber(153), new BigNumber((_559_v15).length))];
              let _lhs53 = globalState;
              let _lhs54 = _543_v1;
              let _lhs55 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_543_v1).length));
              let _lhs56 = _565_v16;
              let _lhs57 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_565_v16).length));
              let _lhs58 = _559_v15;
              let _lhs59 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_559_v15).length));
              _lhs53.f1 = _rhs77;
              _lhs54[_lhs55] = _rhs78;
              _lhs56[_lhs57] = _rhs79;
              _lhs58[_lhs59] = _rhs80;
            } else {
              let _566_v17;
              _566_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(386),p0);
              _566_v17 = _566_v17;
              _548_v4 = _548_v4;
              (_this).f12 = _this.f12;
              let _567_v18;
              _567_v18 = _module.D1.create_DC3(_543_v1);
              let _568_v19;
              let _nw100 = Array((new BigNumber(19)).toNumber());
              _nw100[(_dafny.ZERO).toNumber()] = _543_v1;
              _nw100[(_dafny.ONE).toNumber()] = _543_v1;
              _nw100[(new BigNumber(2)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(3)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(4)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(5)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(6)).toNumber()] = (_567_v18).dtor_cf4;
              _nw100[(new BigNumber(7)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(8)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(9)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(10)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(11)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(12)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(13)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(14)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(15)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(16)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(17)).toNumber()] = _543_v1;
              _nw100[(new BigNumber(18)).toNumber()] = _543_v1;
              _568_v19 = _nw100;
              let _569_v20;
              _569_v20 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_568_v19);
              _569_v20 = (_569_v20).update(_this.f12, _568_v19);
              let _570_v21;
              _570_v21 = _dafny.Map.Empty.slice().updateUnsafe(_567_v18,_this.f12);
              let _571_v22;
              _571_v22 = _dafny.MultiSet.fromElements((_this).f13, (_this).f13);
              let _572_v23;
              _572_v23 = new _dafny.CodePoint('c'.codePointAt(0));
              let _573_v24;
              _573_v24 = _module.D1.create_DC4(_this.f12, _this.f12, new BigNumber((_570_v21).length), _dafny.Seq.of(new BigNumber((_541_v0).length), new BigNumber((_module.__default.fm2(_this.f12, new BigNumber((_571_v22).cardinality()), _572_v23, p0, globalState)).length)), (_this).fm11(_dafny.MultiSet.fromElements(new BigNumber((_541_v0).length)), globalState));
              let _574_v25;
              _574_v25 = _dafny.Seq.of(_573_v24, _573_v24, _573_v24, _573_v24, _573_v24);
              let _575_v26;
              _575_v26 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_574_v25);
              let _576_v27;
              _576_v27 = _module.D0.create_DC1(!(!(new BigNumber((_575_v26).length)).isEqualTo((_this).f13)), new BigNumber(156));
              let _index104 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_543_v1).length));
              (_543_v1)[_index104] = _this.f12;
              let _index105 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_568_v19).length));
              (_568_v19)[_index105] = _543_v1;
              let _index106 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_543_v1).length));
              let _index107 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_568_v19).length));
              let _rhs81 = _576_v27;
              let _rhs82 = ((_this.f12) ? (!((_571_v22).IsSubsetOf(_571_v22))) : (_this.f12));
              let _rhs83 = _module.__default.safeDivisionInt(((_this).f13).multipliedBy((_this).fm11(_571_v22, globalState)), new BigNumber((((_this.f12) ? (_566_v17) : (_566_v17))).length));
              let _rhs84 = _543_v1;
              let _lhs60 = _543_v1;
              let _lhs61 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_543_v1).length));
              let _lhs62 = globalState;
              let _lhs63 = _568_v19;
              let _lhs64 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_568_v19).length));
              _576_v27 = _rhs81;
              _lhs60[_lhs61] = _rhs82;
              _lhs62.f1 = _rhs83;
              _lhs63[_lhs64] = _rhs84;
            }
            let _577_v28;
            _577_v28 = _module.D1.create_DC3(_543_v1);
            let _578_v29;
            let _nw101 = Array((new BigNumber(3)).toNumber());
            _nw101[(_dafny.ZERO).toNumber()] = _577_v28;
            _nw101[(_dafny.ONE).toNumber()] = _577_v28;
            _nw101[(new BigNumber(2)).toNumber()] = _577_v28;
            _578_v29 = _nw101;
            let _index108 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_578_v29).length));
            (_578_v29)[_index108] = _577_v28;
            let _index109 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_578_v29).length));
            (_578_v29)[_index109] = _577_v28;
            if (_this.f12) {
              let _579_v30;
              _579_v30 = new _dafny.CodePoint('n'.codePointAt(0));
              let _580_v31;
              _580_v31 = _dafny.Set.fromElements(false, _this.f12);
              let _581_v32;
              _581_v32 = _dafny.Set.fromElements(_module.__default.fm2(_this.f12, new BigNumber((_580_v31).length), _579_v30, p1, globalState), _dafny.Seq.UnicodeFromString("h"));
              let _582_v33;
              _582_v33 = _module.D0.create_DC1((new BigNumber(-711)).isLessThanOrEqualTo(p0), new BigNumber((_581_v32).length));
              let _rhs85 = (_this).f13;
              let _rhs86 = new _dafny.CodePoint('c'.codePointAt(0));
              let _rhs87 = _582_v33;
              let _lhs65 = globalState;
              _lhs65.f1 = _rhs85;
              _579_v30 = _rhs86;
              _582_v33 = _rhs87;
              let _rhs88 = (_this).f13;
              let _rhs89 = p1;
              let _lhs66 = globalState;
              let _lhs67 = globalState;
              _lhs66.f1 = _rhs88;
              _lhs67.f1 = _rhs89;
              let _583_v34;
              _583_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-505),_this.f12);
              let _584_v35;
              _584_v35 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,(_this).f13);
              _583_v34 = (_583_v34).update(new BigNumber(((_584_v35).update(_this.f12, p0)).length), _this.f12);
              let _585_v36;
              let _nw102 = Array((new BigNumber(26)).toNumber()).fill(_dafny.MultiSet.Empty);
              _585_v36 = _nw102;
              let _init18 = function (_586_i6) {
                return _dafny.MultiSet.fromElements((_this).f13, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_this.f12)).length))).cardinality()));
              };
              let _nw103 = Array((new BigNumber(13)).toNumber());
              for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw103.length); _i0_18++) {
                _nw103[_i0_18] = _init18(new BigNumber(_i0_18));
              }
              _585_v36 = _nw103;
              (_this).f12 = _this.f12;
            } else {
              let _587_v37;
              let _init19 = ((_588_v3) => function (_589_i7) {
                return _588_v3;
              })(_545_v3);
              let _nw104 = Array((new BigNumber(27)).toNumber());
              for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw104.length); _i0_19++) {
                _nw104[_i0_19] = _init19(new BigNumber(_i0_19));
              }
              _587_v37 = _nw104;
              let _index110 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_587_v37).length));
              (_587_v37)[_index110] = _545_v3;
              let _index111 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_587_v37).length));
              (_587_v37)[_index111] = _545_v3;
              (globalState).f1 = p1;
              let _index112 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_578_v29).length));
              (_578_v29)[_index112] = (_578_v29)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_578_v29).length))];
              (globalState).f1 = new BigNumber(66);
              let _590_v38;
              _590_v38 = _dafny.MultiSet.fromElements(_this.f12, !(_this.f12), _this.f12, _this.f12);
              let _591_v39;
              _591_v39 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("micx")).length)), p0);
              let _592_v40;
              _592_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm11(_591_v39, globalState),_this.f12);
              let _593_v41;
              _593_v41 = _module.D0.create_DC1(_this.f12, p0);
              let _rhs90 = ((_590_v38).Union(_dafny.MultiSet.FromArray(_module.__default.fm12(globalState)))).Intersect(_590_v38);
              let _rhs91 = ((((!((((_592_v40).contains((_593_v41).dtor_cf2)) ? ((_592_v40).get((_593_v41).dtor_cf2)) : (_this.f12)))) ? (false) : (_this.f12))) ? (_543_v1) : (_543_v1));
              let _lhs68 = globalState;
              _590_v38 = _rhs90;
              _lhs68.f5 = _rhs91;
            }
            if (_dafny.Seq.IsProperPrefixOf(_541_v0, _dafny.Seq.UnicodeFromString("hmmeoafg"))) {
              let _594_v42;
              _594_v42 = new _dafny.CodePoint('g'.codePointAt(0));
              let _595_v43;
              _595_v43 = _module.D0.create_DC0(_594_v42);
              let _596_v44;
              _596_v44 = _module.D0.create_DC1(!(_this.f12), (_dafny.ZERO).minus((p2)[_module.__default.safeIndex(p1, new BigNumber((p2).length))]));
              let _index113 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_543_v1).length));
              (_543_v1)[_index113] = true;
              let _597_v45;
              _597_v45 = _dafny.MultiSet.fromElements(p0, new BigNumber(756), (_this).f13, (_this).f13, (p2)[_module.__default.safeIndex(p1, new BigNumber((p2).length))]);
              let _598_v46;
              _598_v46 = _597_v45;
              let _index114 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_543_v1).length));
              let _rhs92 = (_this).fm11(((_598_v46)).Union(_597_v45), globalState);
              let _rhs93 = _595_v43;
              let _rhs94 = function (_pat_let10_0) {
                return function (_601_dt__update__tmp_h1) {
                  return function (_pat_let13_0) {
                    return function (_602_dt__update_hcf1_h1) {
                      return _module.D0.create_DC1(_602_dt__update_hcf1_h1, (_601_dt__update__tmp_h1).dtor_cf2);
                    }(_pat_let13_0);
                  }(_this.f12);
                }(_pat_let10_0);
              }(function (_pat_let11_0) {
                return function (_599_dt__update__tmp_h0) {
                  return function (_pat_let12_0) {
                    return function (_600_dt__update_hcf1_h0) {
                      return _module.D0.create_DC1(_600_dt__update_hcf1_h0, (_599_dt__update__tmp_h0).dtor_cf2);
                    }(_pat_let12_0);
                  }(_this.f12);
                }(_pat_let11_0);
              }(_module.D0.create_DC1(_this.f12, p1)));
              let _rhs95 = _this.f12;
              let _lhs69 = globalState;
              let _lhs70 = _543_v1;
              let _lhs71 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_543_v1).length));
              _lhs69.f1 = _rhs92;
              _595_v43 = _rhs93;
              _596_v44 = _rhs94;
              _lhs70[_lhs71] = _rhs95;
              let _603_v47;
              _603_v47 = _dafny.Seq.of(_541_v0, _541_v0, _541_v0, _541_v0);
              (globalState).f1 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(((_603_v47)[_module.__default.safeIndex(p1, new BigNumber((_603_v47).length))]).length)), new BigNumber(910));
              _541_v0 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bvmvcibbo"), _541_v0), _dafny.Seq.Concat(_dafny.Seq.update(_541_v0, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("hlqwmddvn")).length), new BigNumber((_541_v0).length)), _594_v42), _541_v0)), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bvmvcibbo"), _541_v0), _dafny.Seq.Concat(_dafny.Seq.update(_541_v0, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("hlqwmddvn")).length), new BigNumber((_541_v0).length)), _594_v42), _541_v0))).length)), _594_v42);
              (globalState).f1 = p0;
              let _index115 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_543_v1).length));
              (_543_v1)[_index115] = _this.f12;
            } else {
              (globalState).f1 = p1;
              (_this).f12 = _module.__default.fm4(_549_v5, globalState);
              let _604_v48;
              _604_v48 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_module.__default.fm4(_548_v4, globalState));
              _604_v48 = (_604_v48).update(_this.f12, _this.f12);
              (_this).f12 = _this.f12;
              let _index116 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_543_v1).length));
              (_543_v1)[_index116] = _this.f12;
              let _605_v49;
              _605_v49 = new _dafny.CodePoint('q'.codePointAt(0));
              let _606_v50;
              let _nw105 = Array((new BigNumber(22)).toNumber());
              _nw105[(_dafny.ZERO).toNumber()] = _605_v49;
              _nw105[(_dafny.ONE).toNumber()] = _605_v49;
              _nw105[(new BigNumber(2)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(3)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(4)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(5)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(6)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(7)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(8)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(9)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(10)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(11)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(12)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(13)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(14)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(15)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(16)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(17)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(18)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(19)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(20)).toNumber()] = _605_v49;
              _nw105[(new BigNumber(21)).toNumber()] = _605_v49;
              _606_v50 = _nw105;
              let _607_v51;
              _607_v51 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((p2).length));
              let _608_v52;
              _608_v52 = _dafny.Map.Empty.slice().updateUnsafe(_606_v50,(((_607_v51).contains(_this.f12)) ? ((_607_v51).get(_this.f12)) : (p0)));
              let _index117 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_543_v1).length));
              let _rhs96 = (true) && (!(_this.f12));
              let _rhs97 = _this.f12;
              let _rhs98 = _608_v52;
              let _rhs99 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), ((_609_v49) => function (_610_i8) {
                return _609_v49;
              })(_605_v49)), _541_v0), _dafny.Seq.UnicodeFromString("anvdwavn")), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_this.f12)).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), ((_611_v49) => function (_612_i8) {
                return _611_v49;
              })(_605_v49)), _541_v0), _dafny.Seq.UnicodeFromString("anvdwavn"))).length)), _605_v49);
              let _rhs100 = _this.f12;
              let _lhs72 = _this;
              let _lhs73 = _543_v1;
              let _lhs74 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_543_v1).length));
              let _lhs75 = _this;
              _lhs72.f12 = _rhs96;
              _lhs73[_lhs74] = _rhs97;
              _608_v52 = _rhs98;
              _541_v0 = _rhs99;
              _lhs75.f12 = _rhs100;
            }
          }
        }
      }
      let _613_v53;
      let _nw106 = new _module.C0();
      _nw106.__ctor();
      _613_v53 = _nw106;
      let _614_v54;
      _614_v54 = _dafny.Seq.of(_613_v53, _613_v53, _613_v53, _613_v53, _613_v53);
      let _615_v55;
      _615_v55 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f12);
      let _616_v56;
      _616_v56 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_615_v55).length)), p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_this.f12)).length), (_this).f13);
      let _617_v57;
      _617_v57 = _dafny.MultiSet.fromElements(new BigNumber(623), new BigNumber((_616_v56).cardinality()), (_this).f13);
      let _618_v58;
      _618_v58 = _dafny.Map.Empty.slice().updateUnsafe(false,_613_v53);
      let _619_v59;
      _619_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(802),_613_v53);
      let _620_v60;
      _620_v60 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_543_v1);
      let _621_v61;
      let _nw107 = Array((new BigNumber(20)).toNumber());
      _nw107[(_dafny.ZERO).toNumber()] = _613_v53;
      _nw107[(_dafny.ONE).toNumber()] = _613_v53;
      _nw107[(new BigNumber(2)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(3)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(4)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(5)).toNumber()] = (_614_v54)[_module.__default.safeIndex((_this).fm11(_616_v56, globalState), new BigNumber((_614_v54).length))];
      _nw107[(new BigNumber(6)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(7)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(8)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(9)).toNumber()] = (((_618_v58).contains(_this.f12)) ? ((_618_v58).get(_this.f12)) : (_613_v53));
      _nw107[(new BigNumber(10)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(11)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(12)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(13)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(14)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(15)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(16)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(17)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(18)).toNumber()] = _613_v53;
      _nw107[(new BigNumber(19)).toNumber()] = (((_619_v59).contains(new BigNumber((_620_v60).length))) ? ((_619_v59).get(new BigNumber((_620_v60).length))) : (_613_v53));
      _621_v61 = _nw107;
      let _622_v62;
      _622_v62 = _dafny.Seq.of(_621_v61, _621_v61, _621_v61, _621_v61);
      _622_v62 = _dafny.Seq.Concat(_622_v62, _dafny.Seq.Concat(_622_v62, _622_v62));
      let _623_v63;
      let _nw108 = new _module.C0();
      _nw108.__ctor();
      _623_v63 = _nw108;
      let _624_v64;
      let _nw109 = new _module.C1();
      _nw109.__ctor(_this.f12, _this.f12, (_this).f10);
      _624_v64 = _nw109;
      let _625_v65;
      _625_v65 = _dafny.Set.fromElements(_624_v64);
      let _626_v66;
      _626_v66 = _dafny.Map.Empty.slice().updateUnsafe(_625_v65,false);
      r0 = (_626_v66).Merge(_626_v66);
      return r0;
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f16 = false;
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor(f16) {
      let _this = this;
      (_this)._f16 = f16;
      return;
    }
    fm30(globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_this).f16)).length), new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-822)), function (_627_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-739),(_this).f16)).length), new BigNumber((_dafny.Seq.UnicodeFromString("nfqbkcer")).length))).length), new BigNumber(956), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f16,new BigNumber(949))).length))),_module.D9.create_DC23(new BigNumber(48)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll34 = new _dafny.Set();
        for (const _compr_34 of _dafny.IntegerRange(new BigNumber(-12), new BigNumber(-592))) {
          let _628_v0 = _compr_34;
          if (((new BigNumber(-12)).isLessThanOrEqualTo(_628_v0)) && ((_628_v0).isLessThan(new BigNumber(-592)))) {
            _coll34.add((_628_v0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ypcu")).length)),!((_this).f16))).length)));
          }
        }
        return _coll34;
      }(),_module.D9.create_DC23(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(304),new BigNumber((_dafny.Seq.UnicodeFromString("uqqfvcsyh")).length))).length))));
    };
    fm31(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(806);
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _629_v0;
      _629_v0 = _dafny.Set.fromElements(p1, new BigNumber(-38));
      let _630_v1;
      _630_v1 = _dafny.Seq.of(p1);
      (_this).m7(_module.__default.fm4(_629_v0, globalState), _630_v1, globalState);
      let _631_v2;
      let _nw110 = Array((new BigNumber(10)).toNumber()).fill(false);
      _631_v2 = _nw110;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_631_v2).length))) {
        let _632_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_632_i0)) && ((_632_i0).isLessThan(new BigNumber((_631_v2).length))))) {
          (_631_v2)[(_632_i0)] = (_this).f16;
        }
      }
      let _633_v3;
      _633_v3 = _dafny.Set.fromElements(false);
      let _634_v4;
      _634_v4 = _dafny.Seq.of(!((_this).f16));
      let _635_v5;
      let _nw111 = new _module.C1();
      _nw111.__ctor((_this).f16, (_this).f16, _dafny.MultiSet.FromArray(_634_v4));
      _635_v5 = _nw111;
      let _636_v6;
      _636_v6 = _dafny.Seq.of(_635_v5, _635_v5);
      let _637_v7;
      _637_v7 = _dafny.MultiSet.fromElements(p1, new BigNumber((_633_v3).length), p1, new BigNumber((_636_v6).length), p1);
      let _638_v8;
      let _nw112 = Array((new BigNumber(14)).toNumber()).fill([]);
      _638_v8 = _nw112;
      let _639_v9;
      _639_v9 = _dafny.Map.Empty.slice().updateUnsafe((_637_v7).Difference(_637_v7),_638_v8);
      _639_v9 = (_639_v9).update(_637_v7, _638_v8);
      if ((new BigNumber(793)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber(-839)))) {
        (globalState).f1 = p1;
        let _640_v10;
        _640_v10 = _dafny.Seq.UnicodeFromString("cxakbjhpc");
        let _641_v11;
        _641_v11 = _module.D1.create_DC5(new BigNumber(748), (_635_v5).f15, new BigNumber((_640_v10).length), (_this).f16, p1);
        let _642_v12;
        _642_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _643_v13;
        _643_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_635_v5).f15);
        let _644_v14;
        let _nw113 = Array((new BigNumber(11)).toNumber());
        _nw113[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_640_v10, _640_v10);
        _nw113[(_dafny.ONE).toNumber()] = _640_v10;
        _nw113[(new BigNumber(2)).toNumber()] = _640_v10;
        _nw113[(new BigNumber(3)).toNumber()] = _640_v10;
        _nw113[(new BigNumber(4)).toNumber()] = _640_v10;
        _nw113[(new BigNumber(5)).toNumber()] = (((_635_v5).f14) ? (_dafny.Seq.update(_dafny.Seq.update(_640_v10, _module.__default.safeIndex(new BigNumber((_module.__default.fm32((_641_v11).dtor_cf14, p0, _642_v12, (_635_v5).f15, globalState)).length), new BigNumber((_640_v10).length)), p0), _module.__default.safeIndex(new BigNumber((_643_v13).length), new BigNumber((_dafny.Seq.update(_640_v10, _module.__default.safeIndex(new BigNumber((_module.__default.fm32((_641_v11).dtor_cf14, p0, _642_v12, (_635_v5).f15, globalState)).length), new BigNumber((_640_v10).length)), p0)).length)), p0)) : (_640_v10));
        _nw113[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_640_v10, _640_v10);
        _nw113[(new BigNumber(7)).toNumber()] = (((_635_v5).f14) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(555)), ((_645_p0) => function (_646_i1) {
          return _645_p0;
        })(p0))) : (_640_v10));
        _nw113[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_640_v10, _640_v10);
        _nw113[(new BigNumber(9)).toNumber()] = _640_v10;
        _nw113[(new BigNumber(10)).toNumber()] = _640_v10;
        _644_v14 = _nw113;
        _644_v14 = _644_v14;
        let _647_v15;
        let _init20 = ((_648_p1) => function (_649_i2) {
          return _module.__default.safeModuloInt(_649_i2, _648_p1);
        })(p1);
        let _nw114 = Array((new BigNumber(28)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw114.length); _i0_20++) {
          _nw114[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _647_v15 = _nw114;
        let _index118 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_647_v15).length));
        (_647_v15)[_index118] = p1;
        let _650_v16;
        _650_v16 = _dafny.Seq.of(_640_v10);
        let _651_v17;
        _651_v17 = _dafny.MultiSet.fromElements((_635_v5).f15, (_635_v5).f14, !(false));
        let _index119 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_647_v15).length));
        let _rhs101 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((((_this).f16) ? (new BigNumber(71)) : (new BigNumber((_650_v16).length))), (p1).multipliedBy(new BigNumber((_643_v13).length))));
        let _rhs102 = p1;
        let _rhs103 = new BigNumber((_module.__default.fm33((((_651_v17).contains((_635_v5).f15)) ? ((_651_v17).get((_635_v5).f15)) : (p1)), (p1).isLessThan(p1), (_635_v5).f15, p1, globalState)).length);
        let _rhs104 = (((_637_v7).contains(p1)) ? ((_637_v7).get(p1)) : ((p1).multipliedBy(new BigNumber((_634_v4).length))));
        let _lhs76 = globalState;
        let _lhs77 = globalState;
        let _lhs78 = _647_v15;
        let _lhs79 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_647_v15).length));
        let _lhs80 = globalState;
        _lhs76.f1 = _rhs101;
        _lhs77.f1 = _rhs102;
        _lhs78[_lhs79] = _rhs103;
        _lhs80.f1 = _rhs104;
        _629_v0 = _629_v0;
        (globalState).f1 = (_647_v15)[_module.__default.safeIndex(new BigNumber(713), new BigNumber((_647_v15).length))];
      } else {
        (globalState).f1 = (((_637_v7).contains(new BigNumber(((((_this).f16) ? (_dafny.Seq.of((_this).f16)) : (_634_v4))).length))) ? ((_637_v7).get(new BigNumber(((((_this).f16) ? (_dafny.Seq.of((_this).f16)) : (_634_v4))).length))) : (p1));
        let _652_v18;
        _652_v18 = false;
        let _rhs105 = p1;
        let _rhs106 = (_635_v5).f15;
        let _lhs81 = globalState;
        _lhs81.f1 = _rhs105;
        _652_v18 = _rhs106;
        let _653_v19;
        let _nw115 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _653_v19 = _nw115;
        let _init21 = function (_654_i3) {
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dbl"), _dafny.Seq.UnicodeFromString("x"));
        };
        let _nw116 = Array((new BigNumber(5)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw116.length); _i0_21++) {
          _nw116[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _653_v19 = _nw116;
        let _655_v20;
        _655_v20 = _dafny.MultiSet.fromElements((_635_v5).f14);
        _655_v20 = (_655_v20).Difference(_655_v20);
        let _656_v21;
        _656_v21 = _dafny.Seq.UnicodeFromString("xwcdnmpw");
        let _657_v22;
        _657_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_this).f16)).length),_656_v21);
        let _rhs107 = (p1).isEqualTo(p1);
        let _rhs108 = _dafny.Map.Empty.slice().updateUnsafe((p1).plus(p1),_dafny.Seq.UnicodeFromString("oumom"));
        _652_v18 = _rhs107;
        _657_v22 = _rhs108;
      }
      let _hi5 = _module.__default.safeDivisionInt(p1, p1);
      for (let _658_i4 = _module.__default.safeModuloInt(p1, p1); _658_i4.isLessThan(_hi5); _658_i4 = _658_i4.plus(_dafny.ONE)) {
        let _659_v23;
        let _init22 = ((_660_p1) => function (_661_i5) {
          return (_661_i5).minus(_660_p1);
        })(p1);
        let _nw117 = Array((new BigNumber(27)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw117.length); _i0_22++) {
          _nw117[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _659_v23 = _nw117;
        let _index120 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_659_v23).length));
        (_659_v23)[_index120] = _658_i4;
        let _index121 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_659_v23).length));
        (_659_v23)[_index121] = p1;
        let _662_v24;
        _662_v24 = new _dafny.CodePoint('h'.codePointAt(0));
        let _663_v25;
        _663_v25 = _dafny.Seq.UnicodeFromString("qs");
        let _664_v26;
        _664_v26 = _dafny.Map.Empty.slice().updateUnsafe(!((_658_i4).isLessThanOrEqualTo(new BigNumber((_663_v25).length))),(_635_v5).f15);
        let _665_v27;
        _665_v27 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f16);
        let _666_v28;
        _666_v28 = _module.D9.create_DC22(p1, (_this).f16, _665_v27);
        let _667_v29;
        _667_v29 = _dafny.Map.Empty.slice().updateUnsafe((_635_v5).f15,_665_v27);
        let _668_v30;
        _668_v30 = _dafny.Seq.of(_665_v27, _module.__default.fm34((_666_v28).dtor_cf50, _634_v4, (_630_v1)[_module.__default.safeIndex(new BigNumber((_629_v0).length), new BigNumber((_630_v1).length))], (_635_v5).f15, globalState), _665_v27, _665_v27, (((_667_v29).contains((_635_v5).f14)) ? ((_667_v29).get((_635_v5).f14)) : ((((_667_v29).contains((_635_v5).f15)) ? ((_667_v29).get((_635_v5).f15)) : (_665_v27)))));
        let _669_v31;
        _669_v31 = _dafny.Map.Empty.slice().updateUnsafe((_635_v5).f14,p1);
        let _670_v32;
        _670_v32 = _module.D8.create_DC20(_662_v24, (_635_v5).f15, _module.__default.fm4(_dafny.Set.fromElements((_659_v23)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_659_v23).length))], _658_i4, p1), globalState), _669_v31, (_635_v5).f14);
        let _671_v33;
        _671_v33 = _dafny.Map.Empty.slice().updateUnsafe(_670_v32,p1);
        let _672_v34;
        _672_v34 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC20(_662_v24, (_635_v5).f15, (_635_v5).f15, _dafny.Map.Empty.slice().updateUnsafe((_635_v5).f15,p1), (_this).f16),_658_i4));
        let _index122 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_659_v23).length));
        let _rhs109 = (_dafny.ZERO).minus((_this).fm31(_662_v24, (_this).f16, (((_635_v5).f15) ? (_dafny.Map.Empty.slice().updateUnsafe((_659_v23)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_659_v23).length))],_671_v33)) : (_672_v34)), globalState));
        let _rhs110 = _module.__default.fm1(globalState);
        let _rhs111 = (_664_v26).Merge(_664_v26);
        let _rhs112 = _dafny.Seq.Concat(_dafny.Seq.Concat(_668_v30, _668_v30), _668_v30);
        let _lhs82 = _659_v23;
        let _lhs83 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_659_v23).length));
        _lhs82[_lhs83] = _rhs109;
        _662_v24 = _rhs110;
        _664_v26 = _rhs111;
        _668_v30 = _rhs112;
        _663_v25 = _663_v25;
        let _673_v35;
        _673_v35 = true;
        _673_v35 = (_this).f16;
      }
      let _674_v36;
      let _nw118 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _674_v36 = _nw118;
      let _675_v37;
      _675_v37 = _dafny.Seq.UnicodeFromString("nqgf");
      let _index123 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_674_v36).length));
      (_674_v36)[_index123] = new BigNumber((_675_v37).length);
      let _index124 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_674_v36).length));
      (_674_v36)[_index124] = p1;
      return;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let _676_v0;
      _676_v0 = new BigNumber(53);
      let _677_v1;
      _677_v1 = _dafny.Map.Empty.slice().updateUnsafe((_676_v0).plus(_676_v0),_676_v0);
      _677_v1 = (_677_v1).update(_676_v0, _676_v0);
      let _hi6 = _676_v0;
      for (let _678_i0 = _module.__default.safeDivisionInt(_676_v0, _676_v0); _678_i0.isLessThan(_hi6); _678_i0 = _678_i0.plus(_dafny.ONE)) {
        _676_v0 = _module.__default.safeModuloInt(_678_i0, new BigNumber(689));
        let _679_v2;
        _679_v2 = _dafny.Set.fromElements((_678_i0).minus(_676_v0), new BigNumber(597), _678_i0, _678_i0);
        (globalState).f1 = new BigNumber((_679_v2).length);
        if ((_this).f16) {
          (globalState).f1 = _module.__default.safeModuloInt(_678_i0, _676_v0);
          (globalState).f1 = _676_v0;
          (globalState).f1 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_678_i0), (_dafny.ZERO).minus(new BigNumber((((p0) ? (_679_v2) : (_dafny.Set.fromElements(_678_i0)))).length)));
          let _680_v3;
          _680_v3 = _dafny.MultiSet.fromElements(p0);
          let _681_v4;
          let _nw119 = new _module.C1();
          _nw119.__ctor(false, p0, _680_v3);
          _681_v4 = _nw119;
          let _682_v5;
          _682_v5 = _dafny.Seq.of(_681_v4, _681_v4, _681_v4);
          _676_v0 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_682_v5).length)), (_dafny.ZERO).minus(_678_i0));
          let _683_v6;
          let _init23 = function (_684_i1) {
            return _dafny.Seq.UnicodeFromString("wruaxewg");
          };
          let _nw120 = Array((new BigNumber(23)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw120.length); _i0_23++) {
            _nw120[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _683_v6 = _nw120;
          let _685_v7;
          let _nw121 = Array((new BigNumber(12)).toNumber());
          _nw121[(_dafny.ZERO).toNumber()] = _676_v0;
          _nw121[(_dafny.ONE).toNumber()] = new BigNumber(174);
          _nw121[(new BigNumber(2)).toNumber()] = _676_v0;
          _nw121[(new BigNumber(3)).toNumber()] = _676_v0;
          _nw121[(new BigNumber(4)).toNumber()] = _678_i0;
          _nw121[(new BigNumber(5)).toNumber()] = _678_i0;
          _nw121[(new BigNumber(6)).toNumber()] = _676_v0;
          _nw121[(new BigNumber(7)).toNumber()] = _678_i0;
          _nw121[(new BigNumber(8)).toNumber()] = new BigNumber(687);
          _nw121[(new BigNumber(9)).toNumber()] = _678_i0;
          _nw121[(new BigNumber(10)).toNumber()] = _678_i0;
          _nw121[(new BigNumber(11)).toNumber()] = _676_v0;
          _685_v7 = _nw121;
          let _686_v8;
          let _nw122 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _686_v8 = _nw122;
          let _687_v9;
          _687_v9 = _module.D5.create_DC13(_678_i0, _685_v7, _686_v8, _678_i0);
          let _688_v10;
          _688_v10 = _dafny.Map.Empty.slice().updateUnsafe(_683_v6,_687_v9);
          _688_v10 = (_688_v10).update(_683_v6, _687_v9);
        } else {
          let _689_v11;
          _689_v11 = _module.D9.create_DC24(_678_i0);
          let _690_v12;
          _690_v12 = _dafny.Map.Empty.slice().updateUnsafe(_689_v11,(_this).f16);
          (globalState).f1 = (p1)[_module.__default.safeIndex(new BigNumber((_690_v12).length), new BigNumber((p1).length))];
          let _691_v13;
          let _nw123 = new _module.C0();
          _nw123.__ctor();
          _691_v13 = _nw123;
          _691_v13 = _691_v13;
          let _692_v14;
          _692_v14 = _dafny.MultiSet.fromElements(p0);
          _692_v14 = _692_v14;
          let _693_v15;
          let _nw124 = new _module.C2();
          _nw124.__ctor(p0, _676_v0, _692_v14);
          _693_v15 = _nw124;
          (globalState).f1 = _676_v0;
        }
        let _694_v16;
        let _nw125 = new _module.C0();
        _nw125.__ctor();
        _694_v16 = _nw125;
      }
      let _hi7 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("edunij"), _dafny.Seq.UnicodeFromString("rmlmct"))).length);
      for (let _695_i2 = _676_v0; _695_i2.isLessThan(_hi7); _695_i2 = _695_i2.plus(_dafny.ONE)) {
        let _696_v17;
        _696_v17 = _dafny.Set.fromElements(!(true), (_this).f16);
        let _697_v18;
        _697_v18 = _module.D0.create_DC1(p0, _676_v0);
        let _698_v19;
        let _nw126 = Array((new BigNumber(4)).toNumber());
        _nw126[(_dafny.ZERO).toNumber()] = _695_i2;
        _nw126[(_dafny.ONE).toNumber()] = _695_i2;
        _nw126[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_696_v17).length));
        _nw126[(new BigNumber(3)).toNumber()] = (_697_v18).dtor_cf2;
        _698_v19 = _nw126;
        let _index125 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_698_v19).length));
        (_698_v19)[_index125] = _676_v0;
        let _699_v20;
        _699_v20 = _dafny.Seq.of(p0, (_this).f16);
        let _index126 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_698_v19).length));
        (_698_v19)[_index126] = ((_dafny.ZERO).minus(new BigNumber(((((_this).f16) ? (_699_v20) : (_dafny.Seq.update(_699_v20, _module.__default.safeIndex(_695_i2, new BigNumber((_699_v20).length)), (_this).f16)))).length))).plus(_676_v0);
        let _700_v21;
        let _init24 = ((_701_v0, _702_i2) => function (_703_i3) {
          return _dafny.Map.Empty.slice().updateUnsafe(_701_v0,_702_i2);
        })(_676_v0, _695_i2);
        let _nw127 = Array((new BigNumber(11)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw127.length); _i0_24++) {
          _nw127[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _700_v21 = _nw127;
        _700_v21 = _700_v21;
        let _704_v22;
        _704_v22 = _dafny.MultiSet.fromElements(_676_v0);
        _676_v0 = (((_704_v22).contains((p1)[_module.__default.safeIndex(_676_v0, new BigNumber((p1).length))])) ? ((_704_v22).get((p1)[_module.__default.safeIndex(_676_v0, new BigNumber((p1).length))])) : (_module.__default.safeModuloInt(_695_i2, new BigNumber(409))));
        let _705_v23;
        _705_v23 = new _dafny.CodePoint('m'.codePointAt(0));
        let _706_v24;
        _706_v24 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,p0), _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_this).f16));
        let _707_v25;
        _707_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,_676_v0);
        let _708_v26;
        _708_v26 = _module.D8.create_DC20(_705_v23, (_this).f16, !((_this).f16), _707_v25, p0);
        let _709_v27;
        _709_v27 = _dafny.Map.Empty.slice().updateUnsafe(_708_v26,_676_v0);
        let _710_v28;
        _710_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_706_v24).length),_709_v27);
        let _index127 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_698_v19).length));
        (_698_v19)[_index127] = (_this).fm31(_705_v23, ((_698_v19)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_698_v19).length))]).isLessThanOrEqualTo(new BigNumber((p1).length)), _710_v28, globalState);
      }
      let _711_v29;
      _711_v29 = _dafny.MultiSet.fromElements(_676_v0, _676_v0);
      let _712_v30;
      let _nw128 = Array((new BigNumber(4)).toNumber());
      _nw128[(_dafny.ZERO).toNumber()] = _711_v29;
      _nw128[(_dafny.ONE).toNumber()] = _711_v29;
      _nw128[(new BigNumber(2)).toNumber()] = (_711_v29).Difference(_dafny.MultiSet.FromArray(p1));
      _nw128[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(p1);
      _712_v30 = _nw128;
      _712_v30 = _712_v30;
      let _713_v31;
      _713_v31 = _dafny.Map.Empty.slice().updateUnsafe(_677_v1,_676_v0);
      let _714_v32;
      _714_v32 = true;
      let _715_v33;
      _715_v33 = _dafny.Seq.of(p0);
      let _716_v34;
      _716_v34 = _dafny.Seq.of(_715_v33);
      let _717_v35;
      _717_v35 = _dafny.Set.fromElements(p0);
      let _rhs113 = (_713_v31).update(((_677_v1).update(new BigNumber(((_716_v34)[_module.__default.safeIndex(_676_v0, new BigNumber((_716_v34).length))]).length), _676_v0)).Merge(_677_v1), new BigNumber(((_717_v35).Intersect(_717_v35)).length));
      let _rhs114 = (_this).f16;
      _713_v31 = _rhs113;
      _714_v32 = _rhs114;
      let _718_v36;
      _718_v36 = new _dafny.CodePoint('l'.codePointAt(0));
      let _719_v37;
      let _nw129 = Array((new BigNumber(5)).toNumber()).fill(false);
      _719_v37 = _nw129;
      let _index128 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_719_v37).length));
      (_719_v37)[_index128] = _714_v32;
      let _index129 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_719_v37).length));
      let _rhs115 = _module.__default.fm1(globalState);
      let _rhs116 = p0;
      let _lhs84 = _719_v37;
      let _lhs85 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_719_v37).length));
      _718_v36 = _rhs115;
      _lhs84[_lhs85] = _rhs116;
      return;
    }
    m8(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _720_v0;
      _720_v0 = false;
      _720_v0 = (true) && (((true) ? (p2) : (p2)));
      let _721_v1;
      _721_v1 = _dafny.Seq.UnicodeFromString("le");
      let _722_v2;
      _722_v2 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_721_v1).length));
      let _723_v3;
      _723_v3 = _dafny.Set.fromElements(new BigNumber((_722_v2).length));
      let _724_v4;
      _724_v4 = _dafny.MultiSet.fromElements(p3);
      let _725_v5;
      _725_v5 = _dafny.MultiSet.fromElements(p0, (_this).f16);
      let _726_v6;
      let _nw130 = new _module.C1();
      _nw130.__ctor((p2) && (_module.__default.fm4(_723_v3, globalState)), (_724_v4).IsProperSubsetOf(_724_v4), (_725_v5).update((_this).f16, _module.__default.abs((_dafny.ZERO).minus(p3))));
      _726_v6 = _nw130;
      if (_module.__default.fm4(_723_v3, globalState)) {
        let _727_v7;
        let _nw131 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _727_v7 = _nw131;
        let _728_v8;
        _728_v8 = _dafny.Map.Empty.slice().updateUnsafe(p3,false);
        let _729_v9;
        _729_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(((_728_v8).contains(new BigNumber(593))) ? ((_728_v8).get(new BigNumber(593))) : (true)));
        let _index130 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_727_v7).length));
        (_727_v7)[_index130] = new BigNumber(((_729_v9).Merge((_729_v9).update(!(p2), p2))).length);
        let _index131 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_727_v7).length));
        (_727_v7)[_index131] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_721_v1, _721_v1)).length), p1);
        let _730_v10;
        _730_v10 = _dafny.Seq.of((_726_v6).f14, (p1).isLessThanOrEqualTo((_727_v7)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_727_v7).length))]), (_726_v6).f14, !((_this).f16), (_726_v6).f14);
        _730_v10 = ((false) ? (_730_v10) : ((((_726_v6).f14) ? (_730_v10) : (_730_v10))));
        let _731_v11;
        _731_v11 = _module.D8.create_DC19(_728_v8);
        let _732_v12;
        _732_v12 = _dafny.Seq.of(_731_v11, _731_v11);
        let _733_v13;
        _733_v13 = _dafny.Seq.of(_732_v12, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-57)), ((_734_v11) => function (_735_i0) {
          return _734_v11;
        })(_731_v11)), _732_v12, _732_v12);
        let _736_v15;
        _736_v15 = _dafny.Set.fromElements(_732_v12);
        if (((_736_v15).Union(_736_v15)).IsProperSubsetOf(function () {
          let _coll35 = new _dafny.Set();
          for (const _compr_35 of (_733_v13).Elements) {
            let _737_v14 = _compr_35;
            if (_dafny.Seq.contains(_733_v13, _737_v14)) {
              _coll35.add(_737_v14);
            }
          }
          return _coll35;
        }())) {
          let _738_v16;
          let _nw132 = new _module.C1();
          _nw132.__ctor(p2, (_726_v6).f14, _dafny.MultiSet.fromElements(p2, (_726_v6).f14, p2));
          _738_v16 = _nw132;
          (globalState).f1 = (p3).multipliedBy((p3).multipliedBy((_727_v7)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_727_v7).length))]));
          let _index132 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_727_v7).length));
          (_727_v7)[_index132] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p3));
          r0 = (p1).plus(p3);
          _720_v0 = (_738_v16).f15;
        } else {
          _725_v5 = _725_v5;
          _720_v0 = (_726_v6).f14;
          let _739_v17;
          let _nw133 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _739_v17 = _nw133;
          let _index133 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_739_v17).length));
          (_739_v17)[_index133] = _721_v1;
          let _index134 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_739_v17).length));
          (_739_v17)[_index134] = _dafny.Seq.Concat(_dafny.Seq.Concat(_721_v1, _dafny.Seq.UnicodeFromString("v")), _721_v1);
          _720_v0 = p0;
          let _nw134 = new _module.C1();
          _nw134.__ctor((_726_v6).f14, (_730_v10)[_module.__default.safeIndex(p1, new BigNumber((_730_v10).length))], _725_v5);
          _726_v6 = _nw134;
        }
        let _740_v18;
        _740_v18 = _module.D0.create_DC1((_this).f16, (_727_v7)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_727_v7).length))]);
        let _741_v20;
        _741_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(395),(_727_v7)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_727_v7).length))]);
        let _742_v21;
        _742_v21 = _dafny.MultiSet.fromElements(function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of (_741_v20).Keys.Elements) {
            let _743_v19 = _compr_36;
            if ((_741_v20).contains(_743_v19)) {
              _coll36.push([_module.__default.safeModuloInt(_743_v19, p1),new BigNumber((_dafny.Seq.of(p1)).length)]);
            }
          }
          return _coll36;
        }());
        let _index135 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_727_v7).length));
        (_727_v7)[_index135] = ((_740_v18).dtor_cf2).plus((((_742_v21).contains(_741_v20)) ? ((_742_v21).get(_741_v20)) : (p3)));
        let _744_v22;
        let _nw135 = new _module.C0();
        _nw135.__ctor();
        _744_v22 = _nw135;
      } else {
        let _745_v23;
        let _init25 = ((_746_p0, _747_v3) => function (_748_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(_746_p0,new BigNumber((_747_v3).length));
        })(p0, _723_v3);
        let _nw136 = Array((new BigNumber(21)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw136.length); _i0_25++) {
          _nw136[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _745_v23 = _nw136;
        let _index136 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_745_v23).length));
        (_745_v23)[_index136] = _722_v2;
        let _index137 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_745_v23).length));
        (_745_v23)[_index137] = _722_v2;
        r0 = (p3).plus(new BigNumber(714));
        let _749_v24;
        let _nw137 = new _module.C0();
        _nw137.__ctor();
        _749_v24 = _nw137;
        let _750_v25;
        _750_v25 = _dafny.Map.Empty.slice().updateUnsafe(_749_v24,true);
        _750_v25 = (_750_v25).update(_749_v24, p2);
        let _751_v26;
        _751_v26 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(952)).isLessThan(p1),(_726_v6).f14);
        _751_v26 = (_751_v26).update(!(_dafny.areEqual(_721_v1, _721_v1)), _720_v0);
        if ((p0) === ((_726_v6).f15)) {
          _720_v0 = (_726_v6).f15;
          (globalState).f1 = p1;
          let _752_v27;
          _752_v27 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_726_v6).f14);
          let _753_v28;
          _753_v28 = _dafny.Seq.of(p1, p3);
          let _754_v29;
          _754_v29 = _module.D1.create_DC4(p2, (_726_v6).f15, p1, _753_v28, p3);
          let _755_v30;
          _755_v30 = _dafny.Seq.of(p0, (_726_v6).f14);
          let _756_v31;
          let _nw138 = Array((new BigNumber(26)).toNumber());
          _nw138[(_dafny.ZERO).toNumber()] = (new BigNumber((_752_v27).length)).minus(p1);
          _nw138[(_dafny.ONE).toNumber()] = p3;
          _nw138[(new BigNumber(2)).toNumber()] = new BigNumber(((_751_v26).Merge(_751_v26)).length);
          _nw138[(new BigNumber(3)).toNumber()] = p1;
          _nw138[(new BigNumber(4)).toNumber()] = new BigNumber((_725_v5).cardinality());
          _nw138[(new BigNumber(5)).toNumber()] = p1;
          _nw138[(new BigNumber(6)).toNumber()] = (_754_v29).dtor_cf9;
          _nw138[(new BigNumber(7)).toNumber()] = (_749_v24).fm6(p2, p1, globalState);
          _nw138[(new BigNumber(8)).toNumber()] = p1;
          _nw138[(new BigNumber(9)).toNumber()] = (((_724_v4).contains(p3)) ? ((_724_v4).get(p3)) : (p3));
          _nw138[(new BigNumber(10)).toNumber()] = p1;
          _nw138[(new BigNumber(11)).toNumber()] = ((_dafny.ZERO).minus(p1)).multipliedBy(new BigNumber((_725_v5).cardinality()));
          _nw138[(new BigNumber(12)).toNumber()] = p3;
          _nw138[(new BigNumber(13)).toNumber()] = p1;
          _nw138[(new BigNumber(14)).toNumber()] = new BigNumber((_755_v30).length);
          _nw138[(new BigNumber(15)).toNumber()] = p3;
          _nw138[(new BigNumber(16)).toNumber()] = new BigNumber((_751_v26).length);
          _nw138[(new BigNumber(17)).toNumber()] = p3;
          _nw138[(new BigNumber(18)).toNumber()] = p1;
          _nw138[(new BigNumber(19)).toNumber()] = p3;
          _nw138[(new BigNumber(20)).toNumber()] = p3;
          _nw138[(new BigNumber(21)).toNumber()] = p1;
          _nw138[(new BigNumber(22)).toNumber()] = (p1).plus(new BigNumber((_721_v1).length));
          _nw138[(new BigNumber(23)).toNumber()] = p1;
          _nw138[(new BigNumber(24)).toNumber()] = p3;
          _nw138[(new BigNumber(25)).toNumber()] = _module.__default.safeDivisionInt(p1, p1);
          _756_v31 = _nw138;
          let _index138 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_756_v31).length));
          (_756_v31)[_index138] = p1;
          let _index139 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_756_v31).length));
          (_756_v31)[_index139] = _module.__default.safeDivisionInt(p1, p3);
          _720_v0 = (_this).f16;
          _720_v0 = ((_dafny.ZERO).minus((_756_v31)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_756_v31).length))])).isLessThan(new BigNumber((_723_v3).length));
        } else {
          _726_v6 = _726_v6;
          (globalState).f1 = p3;
          let _757_v32;
          _757_v32 = _dafny.Seq.of(new BigNumber((_721_v1).length));
          let _758_v33;
          _758_v33 = _dafny.Seq.of(p2);
          let _759_v34;
          let _nw139 = new _module.C2();
          _nw139.__ctor(_720_v0, _module.__default.safeModuloInt(new BigNumber(-699), new BigNumber((_757_v32).length)), _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_758_v33, _758_v33)));
          _759_v34 = _nw139;
          let _760_v35;
          let _init26 = ((_761_p3) => function (_762_i2) {
            return (_762_i2).plus(_761_p3);
          })(p3);
          let _nw140 = Array((new BigNumber(2)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw140.length); _i0_26++) {
            _nw140[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _760_v35 = _nw140;
          let _index140 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_760_v35).length));
          (_760_v35)[_index140] = (p1).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("et")).length));
          let _763_v36;
          _763_v36 = new _dafny.CodePoint('f'.codePointAt(0));
          let _764_v37;
          _764_v37 = _module.D0.create_DC0(_763_v36);
          let _765_v38;
          _765_v38 = _module.D0.create_DC1((_726_v6).f15, p3);
          let _index141 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_760_v35).length));
          (_760_v35)[_index141] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("r"), _721_v1), _module.__default.safeIndex(_module.__default.safeModuloInt(p3, p3), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("r"), _721_v1)).length)), (_759_v34).fm10(_764_v37, _module.D0.create_DC2(_765_v38), globalState))).length);
          r0 = (_759_v34).f13;
        }
      }
      let _766_v39;
      _766_v39 = _dafny.Seq.of(p2);
      let _767_v40;
      let _nw141 = new _module.C0();
      _nw141.__ctor();
      _767_v40 = _nw141;
      let _768_v41;
      _768_v41 = _dafny.Seq.of(_767_v40);
      let _769_v42;
      _769_v42 = _dafny.Seq.of(_768_v41);
      let _770_v43;
      _770_v43 = _dafny.Map.Empty.slice().updateUnsafe(_766_v39,!_dafny.Seq.contains(_dafny.Seq.update(_769_v42, _module.__default.safeIndex(p1, new BigNumber((_769_v42).length)), _768_v41), _768_v41));
      _770_v43 = (_770_v43).update(_dafny.Seq.Concat(_module.__default.fm33(p3, p0, (_726_v6).f15, p3, globalState), _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(p2), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.of(p2)).length)), (_767_v40).fm13(p2, globalState)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(p2), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.of(p2)).length)), (_767_v40).fm13(p2, globalState))).length)), p0)), _720_v0);
      let _771_i3;
      _771_i3 = _dafny.ZERO;
      L3: {
        while ((_726_v6).f14) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_771_i3)) {
              break L3;
            }
            _771_i3 = (_771_i3).plus(_dafny.ONE);
            let _772_v44;
            _772_v44 = _dafny.Seq.of(_721_v1);
            (globalState).f1 = new BigNumber((_772_v44).length);
            let _773_v45;
            let _nw142 = new _module.C2();
            _nw142.__ctor((_726_v6).f15, new BigNumber(677), _dafny.MultiSet.FromArray(_766_v39));
            _773_v45 = _nw142;
            let _774_v46;
            let _nw143 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _774_v46 = _nw143;
            let _775_v47;
            _775_v47 = _dafny.Set.fromElements((_767_v40).fm13(_773_v45.f12, globalState));
            let _776_v48;
            _776_v48 = _dafny.Seq.of((_773_v45).f13, new BigNumber((_775_v47).length), p3);
            let _777_v49;
            let _nw144 = Array((new BigNumber(8)).toNumber());
            _nw144[(_dafny.ZERO).toNumber()] = p3;
            _nw144[(_dafny.ONE).toNumber()] = p1;
            _nw144[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_766_v39).length));
            _nw144[(new BigNumber(3)).toNumber()] = p1;
            _nw144[(new BigNumber(4)).toNumber()] = new BigNumber((_776_v48).length);
            _nw144[(new BigNumber(5)).toNumber()] = (_773_v45).f13;
            _nw144[(new BigNumber(6)).toNumber()] = (_773_v45).f13;
            _nw144[(new BigNumber(7)).toNumber()] = new BigNumber((_723_v3).length);
            _777_v49 = _nw144;
            let _778_v50;
            _778_v50 = _dafny.Map.Empty.slice().updateUnsafe(_774_v46,_777_v49);
            _778_v50 = _778_v50;
            _720_v0 = (_767_v40).fm5((_775_v47).Intersect(_775_v47), p3, new BigNumber((_722_v2).length), globalState);
          }
        }
      }
      let _779_v51;
      let _nw145 = new _module.C1();
      _nw145.__ctor(p2, (_this).f16, (_725_v5).update(p0, _module.__default.abs(p3)));
      _779_v51 = _nw145;
      r0 = p1;
      return r0;
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm9(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(921);
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _780_v0;
      _780_v0 = false;
      let _781_i0;
      _781_i0 = _dafny.ZERO;
      L4: {
        while (_780_v0) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_781_i0)) {
              break L4;
            }
            _781_i0 = (_781_i0).plus(_dafny.ONE);
            let _782_v1;
            _782_v1 = _dafny.Set.fromElements((_dafny.ZERO).minus(p1), (p1).minus(p1));
            (globalState).f1 = new BigNumber((_782_v1).length);
            let _783_v2;
            let _init27 = ((_784_v0) => function (_785_i1) {
              return _784_v0;
            })(_780_v0);
            let _nw146 = Array((new BigNumber(21)).toNumber());
            for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw146.length); _i0_27++) {
              _nw146[_i0_27] = _init27(new BigNumber(_i0_27));
            }
            _783_v2 = _nw146;
            let _index142 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_783_v2).length));
            (_783_v2)[_index142] = true;
            let _786_v3;
            _786_v3 = _dafny.MultiSet.fromElements(_780_v0, _780_v0);
            let _787_v4;
            _787_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_786_v3).IsSubsetOf(_786_v3));
            let _788_v5;
            _788_v5 = _dafny.Seq.UnicodeFromString("skejjki");
            let _789_v6;
            _789_v6 = _dafny.Map.Empty.slice().updateUnsafe(_780_v0,_788_v5);
            let _790_v7;
            _790_v7 = _dafny.MultiSet.fromElements(p1, new BigNumber(((((_789_v6).contains(!(!((p2).dtor_cf1)))) ? ((_789_v6).get(!(!((p2).dtor_cf1)))) : (_788_v5))).length));
            let _791_v8;
            let _nw147 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
            _791_v8 = _nw147;
            let _792_v9;
            _792_v9 = _dafny.Seq.of(_791_v8);
            let _index143 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_783_v2).length));
            let _rhs117 = p1;
            let _rhs118 = (((_790_v7).contains((p1).multipliedBy(p1))) ? ((_790_v7).get((p1).multipliedBy(p1))) : (new BigNumber((_792_v9).length)));
            let _rhs119 = _780_v0;
            let _rhs120 = _787_v4;
            let _lhs86 = globalState;
            let _lhs87 = globalState;
            let _lhs88 = _783_v2;
            let _lhs89 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_783_v2).length));
            _lhs86.f1 = _rhs117;
            _lhs87.f1 = _rhs118;
            _lhs88[_lhs89] = _rhs119;
            _787_v4 = _rhs120;
            let _793_v10;
            let _nw148 = new _module.C1();
            _nw148.__ctor((_783_v2)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_783_v2).length))], (_783_v2)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_783_v2).length))], _786_v3);
            _793_v10 = _nw148;
            let _794_v11;
            _794_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("eo"),!((_793_v10).f14));
            _782_v1 = (((((_794_v11).contains(_dafny.Seq.UnicodeFromString("ei"))) ? ((_794_v11).get(_dafny.Seq.UnicodeFromString("ei"))) : ((_783_v2)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_783_v2).length))]))) ? (_782_v1) : (_782_v1));
          }
        }
      }
      let _795_v12;
      _795_v12 = _dafny.Seq.of(!(_780_v0));
      let _796_v13;
      _796_v13 = _dafny.Map.Empty.slice().updateUnsafe(_780_v0,_795_v12);
      let _797_v14;
      _797_v14 = _dafny.MultiSet.fromElements(p1);
      let _798_v15;
      _798_v15 = _dafny.Seq.UnicodeFromString("ycsjd");
      let _799_v16;
      let _nw149 = Array((new BigNumber(14)).toNumber());
      _nw149[(_dafny.ZERO).toNumber()] = p1;
      _nw149[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_796_v13).length));
      _nw149[(new BigNumber(2)).toNumber()] = new BigNumber((_797_v14).cardinality());
      _nw149[(new BigNumber(3)).toNumber()] = p1;
      _nw149[(new BigNumber(4)).toNumber()] = p1;
      _nw149[(new BigNumber(5)).toNumber()] = p1;
      _nw149[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p1);
      _nw149[(new BigNumber(7)).toNumber()] = p1;
      _nw149[(new BigNumber(8)).toNumber()] = p1;
      _nw149[(new BigNumber(9)).toNumber()] = p1;
      _nw149[(new BigNumber(10)).toNumber()] = (_this).fm9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(256)), ((_800_p0) => function (_801_i2) {
        return _800_p0;
      })(p0)), p1, globalState);
      _nw149[(new BigNumber(11)).toNumber()] = new BigNumber((_798_v15).length);
      _nw149[(new BigNumber(12)).toNumber()] = p1;
      _nw149[(new BigNumber(13)).toNumber()] = (p1).multipliedBy(p1);
      _799_v16 = _nw149;
      let _802_v17;
      _802_v17 = _dafny.MultiSet.fromElements(_780_v0);
      let _803_v18;
      let _nw150 = new _module.C1();
      _nw150.__ctor(_780_v0, _780_v0, _802_v17);
      _803_v18 = _nw150;
      let _804_v19;
      _804_v19 = _dafny.Map.Empty.slice().updateUnsafe(_780_v0,_803_v18);
      let _805_v20;
      _805_v20 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_804_v19);
      let _index144 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length));
      (_799_v16)[_index144] = _module.__default.safeModuloInt(new BigNumber((_805_v20).length), p1);
      let _index145 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length));
      (_799_v16)[_index145] = (p1).multipliedBy(p1);
      let _806_v21;
      let _nw151 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
      _806_v21 = _nw151;
      let _index146 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_806_v21).length));
      (_806_v21)[_index146] = _795_v12;
      let _index147 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_806_v21).length));
      (_806_v21)[_index147] = _795_v12;
      let _hi8 = p1;
      for (let _807_i3 = (_this).fm9(_798_v15, (_this).fm9(_798_v15, (_dafny.ZERO).minus((_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))]), globalState), globalState); _807_i3.isLessThan(_hi8); _807_i3 = _807_i3.plus(_dafny.ONE)) {
        _780_v0 = !(!((false) || ((_803_v18).f15)));
        let _808_v22;
        _808_v22 = _dafny.Seq.of(_807_i3, (_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))], (_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))]);
        let _809_v23;
        _809_v23 = _dafny.Seq.of(_808_v22, _808_v22, _808_v22);
        let _index148 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length));
        let _rhs121 = (_809_v23)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_809_v23).length))];
        let _rhs122 = _module.__default.safeDivisionInt((new BigNumber(((_806_v21)[_module.__default.safeIndex(new BigNumber(824), new BigNumber((_806_v21).length))]).length)).multipliedBy(p1), (_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))]);
        let _lhs90 = _799_v16;
        let _lhs91 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length));
        _808_v22 = _rhs121;
        _lhs90[_lhs91] = _rhs122;
        let _810_v24;
        let _init28 = ((_811_i3, _812_v16) => function (_813_i4) {
          return (_811_i3).isEqualTo((_812_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_812_v16).length))]);
        })(_807_i3, _799_v16);
        let _nw152 = Array((new BigNumber(4)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw152.length); _i0_28++) {
          _nw152[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _810_v24 = _nw152;
        let _index149 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_810_v24).length));
        (_810_v24)[_index149] = (_803_v18).f14;
        let _814_v25;
        _814_v25 = _dafny.Set.fromElements((_dafny.ZERO).minus((_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))]));
        let _index150 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_810_v24).length));
        (_810_v24)[_index150] = _module.__default.fm4(_814_v25, globalState);
        let _815_v26;
        _815_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,_798_v15);
        let _816_v27;
        _816_v27 = _module.D6.create_DC17((_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))], _dafny.Seq.UnicodeFromString("hetkxtwq"), (_803_v18).f15);
        let _817_v28;
        _817_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_803_v18).f15);
        let _818_v29;
        _818_v29 = _module.D6.create_DC17((_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))], (_816_v27).dtor_cf39, (((_817_v28).contains((_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))])) ? ((_817_v28).get((_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))])) : (_780_v0)));
        (globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_807_i3, new BigNumber((_dafny.Seq.Concat((((_815_v26).contains(p0)) ? ((_815_v26).get(p0)) : (_798_v15)), (_818_v29).dtor_cf39)).length)));
      }
      let _ingredients0 = [];
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_799_v16).length))) {
        let _819_i5 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_819_i5)) && ((_819_i5).isLessThan(new BigNumber((_799_v16).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_799_v16, (_819_i5).toNumber(), _module.__default.safeModuloInt(_819_i5, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_799_v16)[_module.__default.safeIndex(new BigNumber(733), new BigNumber((_799_v16).length))],(_803_v18).f15)).Merge((_module.D8.create_DC19(_dafny.Map.Empty.slice().updateUnsafe(p1,(_803_v18).f15))).dtor_cf42)).length))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _pat_let_tv18 = p1;
      let _pat_let_tv19 = p1;
      let _pat_let_tv20 = _797_v14;
      let _pat_let_tv21 = _797_v14;
      let _rhs123 = _module.__default.safeDivisionInt(_module.__default.fm19(p0, _780_v0, globalState), _module.__default.safeDivisionInt(p1, p1));
      let _rhs124 = (((_803_v18).f14) ? (_799_v16) : (_799_v16));
      let _rhs125 = !((_803_v18).f15) || ((_803_v18).f14);
      let _rhs126 = function (_source9) {
        if (_source9.is_DC1) {
          let _820___mcc_h0 = (_source9).cf1;
          let _821___mcc_h1 = (_source9).cf2;
          let _822_cf2 = _821___mcc_h1;
          let _823_cf1 = _820___mcc_h0;
          return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_pat_let_tv18), _dafny.Seq.of(_pat_let_tv19, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(176)), function (_824_i6) {
            return new BigNumber(-747);
          })).length))));
        } else if (_source9.is_DC0) {
          let _825___mcc_h2 = (_source9).cf0;
          let _826_cf0 = _825___mcc_h2;
          return _pat_let_tv20;
        } else {
          let _827___mcc_h3 = (_source9).cf3;
          let _828_cf3 = _827___mcc_h3;
          return _pat_let_tv21;
        }
      }(_module.__default.fm21(globalState));
      let _lhs92 = globalState;
      _lhs92.f1 = _rhs123;
      _799_v16 = _rhs124;
      _780_v0 = _rhs125;
      _797_v14 = _rhs126;
      return;
    }
    m2(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _829_v0;
      _829_v0 = false;
      let _830_i0;
      _830_i0 = _dafny.ZERO;
      L5: {
        while (_829_v0) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_830_i0)) {
              break L5;
            }
            _830_i0 = (_830_i0).plus(_dafny.ONE);
            let _831_v1;
            let _nw153 = Array((new BigNumber(6)).toNumber()).fill(false);
            _831_v1 = _nw153;
            let _index151 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length));
            (_831_v1)[_index151] = _829_v0;
            let _index152 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length));
            (_831_v1)[_index152] = _829_v0;
            let _832_v2;
            _832_v2 = new BigNumber(809);
            let _833_v3;
            let _nw154 = Array((new BigNumber(4)).toNumber());
            _nw154[(_dafny.ZERO).toNumber()] = _832_v2;
            _nw154[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_832_v2);
            _nw154[(new BigNumber(2)).toNumber()] = new BigNumber(315);
            _nw154[(new BigNumber(3)).toNumber()] = _832_v2;
            _833_v3 = _nw154;
            let _834_v4;
            let _nw155 = Array((new BigNumber(2)).toNumber());
            _nw155[(_dafny.ZERO).toNumber()] = _833_v3;
            _nw155[(_dafny.ONE).toNumber()] = ((_829_v0) ? (_833_v3) : (_833_v3));
            _834_v4 = _nw155;
            _834_v4 = _834_v4;
            let _835_v5;
            _835_v5 = _dafny.Map.Empty.slice().updateUnsafe((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))],_832_v2);
            let _836_v6;
            _836_v6 = _dafny.Set.fromElements(_832_v2, _832_v2, new BigNumber((((_835_v5).update((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], new BigNumber(642))).update((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], _832_v2)).length));
            if (!(_module.__default.fm4(_836_v6, globalState))) {
              let _837_v7;
              let _838_v8;
              let _839_v9;
              let _840_v10;
              let _out24;
              let _out25;
              let _out26;
              let _out27;
              let _outcollector6 = (_this).m3(globalState);
              _out24 = _outcollector6[0];
              _out25 = _outcollector6[1];
              _out26 = _outcollector6[2];
              _out27 = _outcollector6[3];
              _837_v7 = _out24;
              _838_v8 = _out25;
              _839_v9 = _out26;
              _840_v10 = _out27;
              let _nw156 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
              _833_v3 = _nw156;
              let _841_v11;
              _841_v11 = new _dafny.CodePoint('j'.codePointAt(0));
              let _842_v12;
              _842_v12 = _module.D3.create_DC8(_841_v11, _829_v0, _module.__default.fm4(_836_v6, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(662)), ((_843_v11) => function (_844_i1) {
  return _843_v11;
})(_841_v11)));
              _829_v0 = (_842_v12).dtor_cf19;
              let _845_v13;
              _845_v13 = _dafny.Map.Empty.slice().updateUnsafe((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))],!(_829_v0));
              _837_v7 = !(!((((_845_v13).contains(_837_v7)) ? ((_845_v13).get(_837_v7)) : (_837_v7))) || (_829_v0));
              _833_v3 = _833_v3;
            } else {
              let _846_v14;
              _846_v14 = _dafny.MultiSet.fromElements((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], (_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], (_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], _829_v0, (_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))]);
              let _index153 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length));
              (_831_v1)[_index153] = (new BigNumber((_846_v14).cardinality())).isLessThan(_832_v2);
              let _847_v15;
              _847_v15 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_832_v2),(_832_v2).plus(_832_v2));
              let _848_v16;
              _848_v16 = _dafny.Seq.of(_835_v5, _835_v5, _dafny.Map.Empty.slice().updateUnsafe((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))],new BigNumber(127)));
              let _849_v17;
              _849_v17 = _dafny.MultiSet.fromElements(_832_v2);
              _847_v15 = (_847_v15).update((new BigNumber(((_848_v16)[_module.__default.safeIndex(_832_v2, new BigNumber((_848_v16).length))]).length)).plus((((_849_v17).contains(_832_v2)) ? ((_849_v17).get(_832_v2)) : (_832_v2))), (_832_v2).plus(_832_v2));
              let _850_v18;
              _850_v18 = new _dafny.CodePoint('k'.codePointAt(0));
              _850_v18 = new _dafny.CodePoint('i'.codePointAt(0));
              let _851_v19;
              _851_v19 = _dafny.Seq.UnicodeFromString("sby");
              let _852_v20;
              _852_v20 = _dafny.Map.Empty.slice().updateUnsafe(_834_v4,_851_v19);
              let _853_v21;
              _853_v21 = _dafny.Seq.of(_851_v19);
              let _854_v22;
              _854_v22 = _dafny.Seq.of(_851_v19, _851_v19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-274)), ((_855_v18) => function (_856_i5) {
                return _855_v18;
              })(_850_v18)), (_853_v21)[_module.__default.safeIndex(_832_v2, new BigNumber((_853_v21).length))], _module.__default.fm2(_829_v0, new BigNumber((_851_v19).length), _850_v18, _832_v2, globalState));
              let _857_v23;
              let _nw157 = Array((new BigNumber(24)).toNumber());
              _nw157[(_dafny.ZERO).toNumber()] = _851_v19;
              _nw157[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(757)), ((_858_v18) => function (_859_i2) {
                return _858_v18;
              })(_850_v18)), _851_v19);
              _nw157[(new BigNumber(2)).toNumber()] = _851_v19;
              _nw157[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("llcxt");
              _nw157[(new BigNumber(4)).toNumber()] = _851_v19;
              _nw157[(new BigNumber(5)).toNumber()] = (((_852_v20).contains(_834_v4)) ? ((_852_v20).get(_834_v4)) : (_851_v19));
              _nw157[(new BigNumber(6)).toNumber()] = _851_v19;
              _nw157[(new BigNumber(7)).toNumber()] = _851_v19;
              _nw157[(new BigNumber(8)).toNumber()] = _851_v19;
              _nw157[(new BigNumber(9)).toNumber()] = _851_v19;
              _nw157[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(377)), ((_860_v18) => function (_861_i3) {
                return _860_v18;
              })(_850_v18));
              _nw157[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_851_v19, _module.__default.safeIndex(new BigNumber(981), new BigNumber((_851_v19).length)), _850_v18);
              _nw157[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_851_v19, _module.__default.fm2((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], _832_v2, _850_v18, _832_v2, globalState));
              _nw157[(new BigNumber(13)).toNumber()] = _851_v19;
              _nw157[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_851_v19, _851_v19);
              _nw157[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(914)), ((_862_v18) => function (_863_i4) {
                return _862_v18;
              })(_850_v18));
              _nw157[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_851_v19, _851_v19);
              _nw157[(new BigNumber(17)).toNumber()] = _851_v19;
              _nw157[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("aqevycure"), _851_v19);
              _nw157[(new BigNumber(19)).toNumber()] = _module.__default.fm2(_829_v0, _832_v2, _850_v18, (_dafny.ZERO).minus(new BigNumber((_851_v19).length)), globalState);
              _nw157[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_851_v19, _module.__default.safeIndex(_832_v2, new BigNumber((_851_v19).length)), _850_v18);
              _nw157[(new BigNumber(21)).toNumber()] = (_853_v21)[_module.__default.safeIndex(_module.__default.fm19(_850_v18, true, globalState), new BigNumber((_853_v21).length))];
              _nw157[(new BigNumber(22)).toNumber()] = _851_v19;
              _nw157[(new BigNumber(23)).toNumber()] = _851_v19;
              _857_v23 = _nw157;
              let _index154 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_857_v23).length));
              (_857_v23)[_index154] = _dafny.Seq.update(_851_v19, _module.__default.safeIndex(_832_v2, new BigNumber((_851_v19).length)), _850_v18);
              let _index155 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_857_v23).length));
              (_857_v23)[_index155] = _851_v19;
              let _864_v24;
              _864_v24 = _module.D3.create_DC8(_850_v18, (_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], (_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], (_857_v23)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_857_v23).length))]);
              let _865_v25;
              _865_v25 = _module.D3.create_DC9(_864_v24);
              _865_v25 = _865_v25;
            }
            let _866_v26;
            _866_v26 = _dafny.Seq.UnicodeFromString("oxvklsjlo");
            let _867_v27;
            _867_v27 = new _dafny.CodePoint('u'.codePointAt(0));
            let _868_v28;
            _868_v28 = _dafny.Map.Empty.slice().updateUnsafe(_829_v0,_867_v27);
            let _869_v29;
            _869_v29 = _dafny.Map.Empty.slice().updateUnsafe((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))],(_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))]);
            let _870_v30;
            _870_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-195),_829_v0);
            let _871_v31;
            _871_v31 = _dafny.Map.Empty.slice().updateUnsafe(_832_v2,new BigNumber((_870_v30).length));
            let _872_v32;
            let _nw158 = Array((new BigNumber(19)).toNumber());
            _nw158[(_dafny.ZERO).toNumber()] = _866_v26;
            _nw158[(_dafny.ONE).toNumber()] = _866_v26;
            _nw158[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_866_v26, _866_v26);
            _nw158[(new BigNumber(3)).toNumber()] = _866_v26;
            _nw158[(new BigNumber(4)).toNumber()] = _module.__default.fm2((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], _832_v2, (((_868_v28).contains((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))])) ? ((_868_v28).get((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))])) : (_867_v27)), new BigNumber((_869_v29).length), globalState);
            _nw158[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("le");
            _nw158[(new BigNumber(6)).toNumber()] = _866_v26;
            _nw158[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_866_v26, _module.__default.safeIndex(_832_v2, new BigNumber((_866_v26).length)), _867_v27), _866_v26);
            _nw158[(new BigNumber(8)).toNumber()] = _866_v26;
            _nw158[(new BigNumber(9)).toNumber()] = _module.__default.fm2((_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], _832_v2, new _dafny.CodePoint('i'.codePointAt(0)), _832_v2, globalState);
            _nw158[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(525)), ((_873_v27) => function (_874_i6) {
              return _873_v27;
            })(_867_v27));
            _nw158[(new BigNumber(11)).toNumber()] = _866_v26;
            _nw158[(new BigNumber(12)).toNumber()] = _866_v26;
            _nw158[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_866_v26, _866_v26);
            _nw158[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(579)), ((_875_v27) => function (_876_i7) {
              return _875_v27;
            })(_867_v27));
            _nw158[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat((_module.__default.fm23(_867_v27, new BigNumber((_871_v31).length), _829_v0, (_831_v1)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_831_v1).length))], globalState)).dtor_cf20, _866_v26);
            _nw158[(new BigNumber(16)).toNumber()] = _dafny.Seq.UnicodeFromString("rkcfltsn");
            _nw158[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(818)), ((_877_v27) => function (_878_i8) {
              return _877_v27;
            })(_867_v27));
            _nw158[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_866_v26, _866_v26);
            _872_v32 = _nw158;
            let _index156 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_872_v32).length));
            (_872_v32)[_index156] = _866_v26;
            let _879_v33;
            _879_v33 = _dafny.Seq.of(_866_v26);
            let _index157 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_872_v32).length));
            (_872_v32)[_index157] = _dafny.Seq.Concat((_879_v33)[_module.__default.safeIndex(_832_v2, new BigNumber((_879_v33).length))], _866_v26);
          }
        }
      }
      let _880_i9;
      _880_i9 = _dafny.ZERO;
      L6: {
        while (!(_829_v0)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_880_i9)) {
              break L6;
            }
            _880_i9 = (_880_i9).plus(_dafny.ONE);
            let _881_v34;
            _881_v34 = new BigNumber(522);
            let _882_v35;
            _882_v35 = _dafny.MultiSet.fromElements(_829_v0);
            let _883_v36;
            let _nw159 = new _module.C1();
            _nw159.__ctor(_829_v0, _829_v0, _882_v35);
            _883_v36 = _nw159;
            let _884_v37;
            _884_v37 = _dafny.Map.Empty.slice().updateUnsafe(_829_v0,_883_v36);
            (globalState).f1 = _module.__default.safeDivisionInt(_881_v34, ((_829_v0) ? (new BigNumber((_884_v37).length)) : (_881_v34)));
            let _885_v38;
            _885_v38 = _dafny.MultiSet.fromElements(_881_v34);
            let _886_v39;
            _886_v39 = _885_v38;
            let _887_v40;
            _887_v40 = _dafny.Map.Empty.slice().updateUnsafe(_886_v39,!(_881_v34).isEqualTo(new BigNumber((_885_v38).cardinality())));
            let _888_v41;
            _888_v41 = _dafny.Set.fromElements(_881_v34, new BigNumber((_882_v35).cardinality()));
            let _889_v42;
            _889_v42 = _dafny.Map.Empty.slice().updateUnsafe(_881_v34,_module.__default.fm4(_888_v41, globalState));
            let _890_v43;
            _890_v43 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(929)), function (_891_i10) {
              return new _dafny.CodePoint('f'.codePointAt(0));
            }),_881_v34);
            _887_v40 = (_887_v40).update(_885_v38, (((_889_v42).contains(new BigNumber((_890_v43).length))) ? ((_889_v42).get(new BigNumber((_890_v43).length))) : (_829_v0)));
            _829_v0 = !(_829_v0);
            let _892_v44;
            let _nw160 = new _module.C0();
            _nw160.__ctor();
            _892_v44 = _nw160;
          }
        }
      }
      let _893_v45;
      _893_v45 = _dafny.Seq.UnicodeFromString("evwq");
      let _894_v46;
      _894_v46 = _dafny.MultiSet.fromElements(_893_v45);
      let _895_v47;
      _895_v47 = new BigNumber(502);
      let _896_v48;
      _896_v48 = _dafny.Seq.of((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_894_v46).cardinality()), (_dafny.ZERO).minus(new BigNumber((_893_v45).length)))), new BigNumber(485), _module.__default.fm19(_module.__default.fm1(globalState), _829_v0, globalState), _895_v47);
      _896_v48 = _dafny.Seq.Concat(_896_v48, _896_v48);
      let _897_i11;
      _897_i11 = _dafny.ZERO;
      L7: {
        while (true) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_897_i11)) {
              break L7;
            }
            _897_i11 = (_897_i11).plus(_dafny.ONE);
            let _898_v49;
            _898_v49 = _dafny.Map.Empty.slice().updateUnsafe(_895_v47,_829_v0);
            let _899_v50;
            _899_v50 = _dafny.Seq.of(_829_v0);
            let _900_v51;
            _900_v51 = _dafny.MultiSet.fromElements(_829_v0, _829_v0);
            let _901_v52;
            _901_v52 = _dafny.Map.Empty.slice().updateUnsafe(_829_v0,_900_v51);
            let _902_v53;
            _902_v53 = _dafny.Map.Empty.slice().updateUnsafe(_829_v0,_module.__default.fm4(_dafny.Set.fromElements(new BigNumber(-482)), globalState));
            let _903_v54;
            _903_v54 = _dafny.MultiSet.fromElements(_895_v47, new BigNumber((_dafny.Seq.of(_829_v0)).length));
            let _904_v55;
            _904_v55 = _dafny.Set.fromElements(_895_v47, _895_v47, _895_v47, _895_v47, new BigNumber((_903_v54).cardinality()));
            let _905_v56;
            _905_v56 = _dafny.Map.Empty.slice().updateUnsafe(_902_v53,_dafny.MultiSet.fromElements(new BigNumber((_904_v55).length), _895_v47));
            let _906_v57;
            _906_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_905_v56).length),new BigNumber((_898_v49).length));
            let _907_v58;
            let _nw161 = Array((new BigNumber(19)).toNumber());
            _nw161[(_dafny.ZERO).toNumber()] = (((_898_v49).contains(_895_v47)) ? ((_898_v49).get(_895_v47)) : (_829_v0));
            _nw161[(_dafny.ONE).toNumber()] = _829_v0;
            _nw161[(new BigNumber(2)).toNumber()] = _829_v0;
            _nw161[(new BigNumber(3)).toNumber()] = !(_dafny.Seq.contains(_899_v50, _829_v0));
            _nw161[(new BigNumber(4)).toNumber()] = _829_v0;
            _nw161[(new BigNumber(5)).toNumber()] = _829_v0;
            _nw161[(new BigNumber(6)).toNumber()] = (_module.__default.fm24(new BigNumber((_899_v50).length), _829_v0, _898_v49, _895_v47, globalState)).IsSubsetOf((((_901_v52).contains(_829_v0)) ? ((_901_v52).get(_829_v0)) : (_900_v51)));
            _nw161[(new BigNumber(7)).toNumber()] = false;
            _nw161[(new BigNumber(8)).toNumber()] = _829_v0;
            _nw161[(new BigNumber(9)).toNumber()] = _829_v0;
            _nw161[(new BigNumber(10)).toNumber()] = true;
            _nw161[(new BigNumber(11)).toNumber()] = _829_v0;
            _nw161[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(863),_895_v47)).equals(_906_v57);
            _nw161[(new BigNumber(13)).toNumber()] = _829_v0;
            _nw161[(new BigNumber(14)).toNumber()] = !(true) || (_829_v0);
            _nw161[(new BigNumber(15)).toNumber()] = _829_v0;
            _nw161[(new BigNumber(16)).toNumber()] = (new BigNumber(206)).isLessThan(_895_v47);
            _nw161[(new BigNumber(17)).toNumber()] = _829_v0;
            _nw161[(new BigNumber(18)).toNumber()] = _829_v0;
            _907_v58 = _nw161;
            let _index158 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_907_v58).length));
            (_907_v58)[_index158] = !((new BigNumber((_902_v53).length)).isLessThan(_895_v47));
            let _index159 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_907_v58).length));
            (_907_v58)[_index159] = ((_829_v0) ? (_829_v0) : (_829_v0));
            let _908_v59;
            _908_v59 = new _dafny.CodePoint('b'.codePointAt(0));
            let _909_v60;
            _909_v60 = _module.D3.create_DC8(_908_v59, _829_v0, _829_v0, _893_v45);
            let _910_v61;
            let _nw162 = Array((new BigNumber(21)).toNumber());
            _nw162[(_dafny.ZERO).toNumber()] = _893_v45;
            _nw162[(_dafny.ONE).toNumber()] = _893_v45;
            _nw162[(new BigNumber(2)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(3)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(4)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(5)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(6)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(7)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(8)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(9)).toNumber()] = _module.__default.fm2(_829_v0, _895_v47, _908_v59, _895_v47, globalState);
            _nw162[(new BigNumber(10)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_893_v45, _module.__default.safeIndex(_895_v47, new BigNumber((_893_v45).length)), _908_v59);
            _nw162[(new BigNumber(12)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(13)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(308)), ((_911_v59) => function (_912_i12) {
              return _911_v59;
            })(_908_v59));
            _nw162[(new BigNumber(15)).toNumber()] = (_909_v60).dtor_cf20;
            _nw162[(new BigNumber(16)).toNumber()] = _dafny.Seq.UnicodeFromString("fbkhxmvx");
            _nw162[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("ooxi");
            _nw162[(new BigNumber(18)).toNumber()] = _893_v45;
            _nw162[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("vvrnb");
            _nw162[(new BigNumber(20)).toNumber()] = _893_v45;
            _910_v61 = _nw162;
            let _913_v62;
            _913_v62 = _dafny.Map.Empty.slice().updateUnsafe(_910_v61,(_899_v50)[_module.__default.safeIndex(_895_v47, new BigNumber((_899_v50).length))]);
            _829_v0 = (((_913_v62).contains(((_829_v0) ? (_910_v61) : (_910_v61)))) ? ((_913_v62).get(((_829_v0) ? (_910_v61) : (_910_v61)))) : (_829_v0));
            let _914_v63;
            let _nw163 = new _module.C2();
            _nw163.__ctor(_829_v0, new BigNumber((_902_v53).length), _900_v51);
            _914_v63 = _nw163;
            let _915_v64;
            let _nw164 = new _module.C0();
            _nw164.__ctor();
            _915_v64 = _nw164;
          }
        }
      }
      let _916_v65;
      let _nw165 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _916_v65 = _nw165;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_916_v65).length))) {
        let _917_i13 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_917_i13)) && ((_917_i13).isLessThan(new BigNumber((_916_v65).length))))) {
          (_916_v65)[(_917_i13)] = _893_v45;
        }
      }
      let _918_v66;
      _918_v66 = _dafny.MultiSet.fromElements(_829_v0, false, _829_v0, _829_v0, _829_v0);
      let _919_v67;
      _919_v67 = new _dafny.CodePoint('b'.codePointAt(0));
      let _920_v68;
      _920_v68 = _dafny.Set.fromElements(_829_v0);
      let _921_v69;
      let _nw166 = Array((new BigNumber(28)).toNumber());
      _nw166[(_dafny.ZERO).toNumber()] = _895_v47;
      _nw166[(_dafny.ONE).toNumber()] = _895_v47;
      _nw166[(new BigNumber(2)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(3)).toNumber()] = new BigNumber((_918_v66).cardinality());
      _nw166[(new BigNumber(4)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(5)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(6)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(7)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(8)).toNumber()] = (((_918_v66).contains(_829_v0)) ? ((_918_v66).get(_829_v0)) : (_895_v47));
      _nw166[(new BigNumber(9)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(10)).toNumber()] = _module.__default.fm19(_919_v67, _829_v0, globalState);
      _nw166[(new BigNumber(11)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(12)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(13)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(14)).toNumber()] = (_this).fm9(_893_v45, (_dafny.ZERO).minus(_895_v47), globalState);
      _nw166[(new BigNumber(15)).toNumber()] = new BigNumber(240);
      _nw166[(new BigNumber(16)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_896_v48)[_module.__default.safeIndex(_895_v47, new BigNumber((_896_v48).length))]));
      _nw166[(new BigNumber(18)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(19)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(20)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(21)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(22)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(23)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(24)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(25)).toNumber()] = new BigNumber((_920_v68).length);
      _nw166[(new BigNumber(26)).toNumber()] = _895_v47;
      _nw166[(new BigNumber(27)).toNumber()] = _895_v47;
      _921_v69 = _nw166;
      let _922_v70;
      let _nw167 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
      _922_v70 = _nw167;
      let _source10 = _module.D5.create_DC13(_895_v47, _921_v69, _922_v70, (new BigNumber(200)).minus(new BigNumber((_893_v45).length)));
      if (_source10.is_DC13) {
        let _923___mcc_h0 = (_source10).cf27;
        let _924___mcc_h1 = (_source10).cf28;
        let _925___mcc_h2 = (_source10).cf29;
        let _926___mcc_h3 = (_source10).cf30;
        let _927_cf30 = _926___mcc_h3;
        let _928_cf29 = _925___mcc_h2;
        let _929_cf28 = _924___mcc_h1;
        let _930_cf27 = _923___mcc_h0;
        let _931_v71;
        _931_v71 = _module.D9.create_DC21(_dafny.Seq.of(_829_v0));
        let _932_v72;
        _932_v72 = _dafny.Seq.of(_829_v0);
        let _933_v73;
        let _934_v74;
        let _935_v75;
        let _936_v76;
        let _out28;
        let _out29;
        let _out30;
        let _out31;
        let _outcollector7 = _module.__default.m0((_931_v71).dtor_cf48, _932_v72, globalState);
        _out28 = _outcollector7[0];
        _out29 = _outcollector7[1];
        _out30 = _outcollector7[2];
        _out31 = _outcollector7[3];
        _933_v73 = _out28;
        _934_v74 = _out29;
        _935_v75 = _out30;
        _936_v76 = _out31;
        let _937_v77;
        _937_v77 = _dafny.Map.Empty.slice().updateUnsafe(_930_cf27,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("shpmce"), _module.__default.safeIndex(new BigNumber(-490), new BigNumber((_dafny.Seq.UnicodeFromString("shpmce")).length)), new _dafny.CodePoint('q'.codePointAt(0))));
        let _938_v78;
        _938_v78 = _module.D9.create_DC24(_895_v47);
        _937_v77 = _module.__default.fm25(_938_v78, globalState);
        _933_v73 = !((((_dafny.ZERO).minus(new BigNumber((_893_v45).length))).plus((_dafny.ZERO).minus(_930_cf27))).isEqualTo(_930_cf27));
        _927_cf30 = _936_v76;
      } else if (_source10.is_DC14) {
        let _939___mcc_h4 = (_source10).cf31;
        let _940___mcc_h5 = (_source10).cf32;
        let _941___mcc_h6 = (_source10).cf33;
        let _942___mcc_h7 = (_source10).cf34;
        let _943___mcc_h8 = (_source10).cf35;
        let _944_cf35 = _943___mcc_h8;
        let _945_cf34 = _942___mcc_h7;
        let _946_cf33 = _941___mcc_h6;
        let _947_cf32 = _940___mcc_h5;
        let _948_cf31 = _939___mcc_h4;
        _896_v48 = _dafny.Seq.of((_895_v47).plus(_895_v47), _895_v47);
        let _949_v79;
        let _nw168 = new _module.C1();
        _nw168.__ctor(_829_v0, _829_v0, _dafny.MultiSet.fromElements(_829_v0));
        _949_v79 = _nw168;
        let _950_v80;
        let _nw169 = Array((new BigNumber(15)).toNumber());
        _nw169[(_dafny.ZERO).toNumber()] = _949_v79;
        _nw169[(_dafny.ONE).toNumber()] = _949_v79;
        _nw169[(new BigNumber(2)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(3)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(4)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(5)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(6)).toNumber()] = ((!(false)) ? (_949_v79) : (_949_v79));
        _nw169[(new BigNumber(7)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(8)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(9)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(10)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(11)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(12)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(13)).toNumber()] = _949_v79;
        _nw169[(new BigNumber(14)).toNumber()] = _949_v79;
        _950_v80 = _nw169;
        let _index160 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_950_v80).length));
        (_950_v80)[_index160] = _949_v79;
        let _index161 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_950_v80).length));
        let _nw170 = new _module.C1();
        _nw170.__ctor(_829_v0, _dafny.Seq.IsPrefixOf(_893_v45, _893_v45), _918_v66);
        (_950_v80)[_index161] = _nw170;
        let _951_v81;
        _951_v81 = _dafny.Set.fromElements(_895_v47, _947_cf32);
        let _952_v82;
        _952_v82 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_951_v81).length)),_829_v0);
        let _953_v83;
        _953_v83 = _dafny.MultiSet.fromElements(_895_v47, _946_cf33, _946_cf33, new BigNumber((_952_v82).length));
        let _954_v84;
        _954_v84 = _dafny.Seq.of((_953_v83).Difference(_dafny.MultiSet.fromElements(_895_v47)), (_dafny.MultiSet.fromElements(_946_cf33)).Difference(_953_v83), (_953_v83).Union(_953_v83), _dafny.MultiSet.FromArray(_896_v48));
        let _955_v85;
        _955_v85 = _dafny.Seq.of(_954_v84, _954_v84, _dafny.Seq.Create(_module.__default.abs(new BigNumber(390)), ((_956_v83) => function (_957_i14) {
          return _956_v83;
        })(_953_v83)), _954_v84, _dafny.Seq.of(_953_v83));
        let _958_v86;
        _958_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_953_v83).cardinality()),_946_cf33);
        _954_v84 = _dafny.Seq.Concat(_dafny.Seq.Concat(_954_v84, _954_v84), (_955_v85)[_module.__default.safeIndex(new BigNumber((_958_v86).length), new BigNumber((_955_v85).length))]);
        let _index162 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_921_v69).length));
        (_921_v69)[_index162] = _module.__default.fm19(_module.__default.fm1(globalState), _829_v0, globalState);
        let _index163 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_921_v69).length));
        (_921_v69)[_index163] = _module.__default.fm19(_919_v67, !(_module.__default.fm4(_951_v81, globalState)), globalState);
      } else if (_source10.is_DC12) {
        let _959___mcc_h9 = (_source10).cf26;
        let _960_cf26 = _959___mcc_h9;
        let _961_v87;
        _961_v87 = _dafny.Map.Empty.slice().updateUnsafe(_895_v47,_895_v47);
        let _rhs127 = !(_829_v0);
        let _rhs128 = _961_v87;
        let _rhs129 = _919_v67;
        _829_v0 = _rhs127;
        _961_v87 = _rhs128;
        _919_v67 = _rhs129;
        let _rhs130 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_895_v47, _895_v47));
        let _rhs131 = (_895_v47).minus(new BigNumber((_893_v45).length));
        let _rhs132 = new BigNumber(561);
        let _lhs93 = globalState;
        r1 = _rhs130;
        _lhs93.f1 = _rhs131;
        r1 = _rhs132;
        let _index164 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_916_v65).length));
        (_916_v65)[_index164] = _dafny.Seq.update(_893_v45, _module.__default.safeIndex((_dafny.ZERO).minus(_895_v47), new BigNumber((_893_v45).length)), _919_v67);
        let _962_v88;
        let _nw171 = Array((new BigNumber(29)).toNumber()).fill(false);
        _962_v88 = _nw171;
        let _963_v89;
        _963_v89 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(440)), function (_964_i15) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        }));
        let _index165 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_916_v65).length));
        let _rhs133 = _962_v88;
        let _rhs134 = (_963_v89)[_module.__default.safeIndex(_895_v47, new BigNumber((_963_v89).length))];
        let _rhs135 = _dafny.Seq.Concat(_893_v45, _dafny.Seq.of(_919_v67));
        let _rhs136 = _893_v45;
        let _rhs137 = !(!(_829_v0));
        let _lhs94 = globalState;
        let _lhs95 = _916_v65;
        let _lhs96 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_916_v65).length));
        _lhs94.f5 = _rhs133;
        _lhs95[_lhs96] = _rhs134;
        r2 = _rhs135;
        _893_v45 = _rhs136;
        _829_v0 = _rhs137;
        let _965_v90;
        let _init29 = ((_966_v0) => function (_967_i16) {
          return _dafny.Seq.of(_966_v0, _966_v0, _966_v0, !(true));
        })(_829_v0);
        let _nw172 = Array((new BigNumber(7)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw172.length); _i0_29++) {
          _nw172[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _965_v90 = _nw172;
        let _968_v91;
        _968_v91 = _dafny.Seq.of(!(true));
        let _index166 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_965_v90).length));
        (_965_v90)[_index166] = _968_v91;
        let _969_v92;
        let _nw173 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
        _969_v92 = _nw173;
        let _index167 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_969_v92).length));
        (_969_v92)[_index167] = (_dafny.Map.Empty.slice().updateUnsafe(_896_v48,_829_v0)).update(_module.__default.fm20(_895_v47, _895_v47, globalState), _829_v0);
        let _970_v93;
        _970_v93 = _dafny.Set.fromElements(_895_v47, _895_v47);
        let _971_v94;
        _971_v94 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_829_v0)).length), _895_v47, _895_v47, (_this).fm9((_916_v65)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_916_v65).length))], new BigNumber(169), globalState));
        let _index168 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_965_v90).length));
        let _index169 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_969_v92).length));
        let _rhs138 = _968_v91;
        let _rhs139 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_970_v93).length)),(new BigNumber((_971_v94).cardinality())).isLessThanOrEqualTo(new BigNumber(424)));
        let _rhs140 = !(_829_v0);
        let _lhs97 = _965_v90;
        let _lhs98 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_965_v90).length));
        let _lhs99 = _969_v92;
        let _lhs100 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_969_v92).length));
        _lhs97[_lhs98] = _rhs138;
        _lhs99[_lhs100] = _rhs139;
        _829_v0 = _rhs140;
      } else {
        let _972___mcc_h10 = (_source10).cf36;
        let _973_cf36 = _972___mcc_h10;
        let _974_v95;
        let _nw174 = new _module.C2();
        _nw174.__ctor(_829_v0, _895_v47, _dafny.MultiSet.fromElements(false));
        _974_v95 = _nw174;
        _974_v95 = _974_v95;
        let _index170 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_921_v69).length));
        (_921_v69)[_index170] = _module.__default.safeDivisionInt(_895_v47, _895_v47);
        let _index171 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_921_v69).length));
        let _rhs141 = _895_v47;
        let _rhs142 = (_dafny.ZERO).minus((_974_v95).f13);
        let _rhs143 = (!(false)) === ((new BigNumber(-387)).isLessThan(_895_v47));
        let _rhs144 = !(_829_v0);
        let _lhs101 = _921_v69;
        let _lhs102 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_921_v69).length));
        let _lhs103 = _974_v95;
        _lhs101[_lhs102] = _rhs141;
        r1 = _rhs142;
        _829_v0 = _rhs143;
        _lhs103.f12 = _rhs144;
        let _975_v96;
        _975_v96 = _dafny.Seq.of(_829_v0);
        let _976_v97;
        _976_v97 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_975_v96, _module.__default.safeIndex((_921_v69)[_module.__default.safeIndex(new BigNumber(149), new BigNumber((_921_v69).length))], new BigNumber((_975_v96).length)), _974_v95.f12)).length), (_921_v69)[_module.__default.safeIndex(new BigNumber(149), new BigNumber((_921_v69).length))]);
        let _977_v98;
        let _nw175 = Array((new BigNumber(11)).toNumber());
        _nw175[(_dafny.ZERO).toNumber()] = _829_v0;
        _nw175[(_dafny.ONE).toNumber()] = _974_v95.f12;
        _nw175[(new BigNumber(2)).toNumber()] = _module.__default.fm4(_976_v97, globalState);
        _nw175[(new BigNumber(3)).toNumber()] = !((_974_v95.f12) && (true));
        _nw175[(new BigNumber(4)).toNumber()] = _974_v95.f12;
        _nw175[(new BigNumber(5)).toNumber()] = _974_v95.f12;
        _nw175[(new BigNumber(6)).toNumber()] = ((_974_v95.f12) ? (_974_v95.f12) : (true));
        _nw175[(new BigNumber(7)).toNumber()] = !(_829_v0);
        _nw175[(new BigNumber(8)).toNumber()] = _974_v95.f12;
        _nw175[(new BigNumber(9)).toNumber()] = !(((_974_v95.f12) ? ((_975_v96)[_module.__default.safeIndex((_974_v95).f13, new BigNumber((_975_v96).length))]) : (_974_v95.f12)));
        _nw175[(new BigNumber(10)).toNumber()] = ((_921_v69)[_module.__default.safeIndex(new BigNumber(149), new BigNumber((_921_v69).length))]).isLessThanOrEqualTo(new BigNumber(575));
        _977_v98 = _nw175;
        let _index172 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_977_v98).length));
        (_977_v98)[_index172] = _829_v0;
        let _index173 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_977_v98).length));
        (_977_v98)[_index173] = ((_920_v68).Union(_920_v68)).IsDisjointFrom(_920_v68);
        (_974_v95).f12 = false;
      }
      let _978_v99;
      _978_v99 = _dafny.Seq.of(_829_v0, _829_v0);
      let _979_v100;
      _979_v100 = _dafny.Map.Empty.slice().updateUnsafe(_895_v47,_895_v47);
      let _980_v101;
      _980_v101 = _dafny.Set.fromElements(_895_v47, new BigNumber((_896_v48).length), _895_v47);
      let _981_v102;
      _981_v102 = _dafny.Set.fromElements(_895_v47, _module.__default.fm19(_919_v67, (_978_v99)[_module.__default.safeIndex(_895_v47, new BigNumber((_978_v99).length))], globalState), (_895_v47).minus(_895_v47), (((_979_v100).contains((_this).fm9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(433)), ((_984_v67) => function (_985_i17) {
        return _984_v67;
      })(_919_v67)), (_module.D6.create_DC17(_895_v47, _893_v45, _module.__default.fm4(_980_v101, globalState))).dtor_cf38, globalState))) ? ((_979_v100).get((_this).fm9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(433)), ((_982_v67) => function (_983_i17) {
        return _982_v67;
      })(_919_v67)), (_module.D6.create_DC17(_895_v47, _893_v45, _module.__default.fm4(_980_v101, globalState))).dtor_cf38, globalState))) : (new BigNumber(153))));
      r0 = _981_v102;
      r1 = _895_v47;
      r2 = _dafny.Seq.Concat(_893_v45, _module.__default.fm2(_829_v0, _895_v47, _919_v67, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_895_v47,_829_v0)).length), globalState));
      return [r0, r1, r2];
    }
    m3(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      if (false) {
        let _986_v0;
        _986_v0 = _dafny.Seq.UnicodeFromString("dtaw");
        let _987_v1;
        _987_v1 = _dafny.Set.fromElements(_986_v0);
        let _988_v2;
        _988_v2 = _dafny.MultiSet.fromElements(_986_v0);
        let _989_v3;
        _989_v3 = _dafny.Seq.of(_988_v2);
        let _990_v4;
        _990_v4 = new BigNumber(794);
        _987_v1 = function () {
          let _coll37 = new _dafny.Set();
          for (const _compr_37 of ((_989_v3)[_module.__default.safeIndex(_990_v4, new BigNumber((_989_v3).length))]).Elements) {
            let _991_v5 = _compr_37;
            if (((_989_v3)[_module.__default.safeIndex(_990_v4, new BigNumber((_989_v3).length))]).contains(_991_v5)) {
              _coll37.add(_991_v5);
            }
          }
          return _coll37;
        }();
        let _992_v6;
        _992_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_986_v0).length),_990_v4);
        _992_v6 = (_992_v6).update((_dafny.ZERO).minus((_990_v4).plus(_990_v4)), (_dafny.ZERO).minus(_990_v4));
        let _993_v7;
        _993_v7 = true;
        let _994_v8;
        _994_v8 = _dafny.Seq.of(true, _993_v7);
        let _995_v9;
        _995_v9 = new _dafny.CodePoint('s'.codePointAt(0));
        (globalState).f1 = new BigNumber((_dafny.Seq.update((((_990_v4).isLessThan((_this).fm9(_986_v0, _990_v4, globalState))) ? (_dafny.Seq.update(_dafny.Seq.Concat(_986_v0, _986_v0), _module.__default.safeIndex(new BigNumber((_994_v8).length), new BigNumber((_dafny.Seq.Concat(_986_v0, _986_v0)).length)), new _dafny.CodePoint('t'.codePointAt(0)))) : (_986_v0)), _module.__default.safeIndex(_990_v4, new BigNumber(((((_990_v4).isLessThan((_this).fm9(_986_v0, _990_v4, globalState))) ? (_dafny.Seq.update(_dafny.Seq.Concat(_986_v0, _986_v0), _module.__default.safeIndex(new BigNumber((_994_v8).length), new BigNumber((_dafny.Seq.Concat(_986_v0, _986_v0)).length)), new _dafny.CodePoint('t'.codePointAt(0)))) : (_986_v0))).length)), _995_v9)).length);
        _994_v8 = _994_v8;
        let _996_v10;
        _996_v10 = _dafny.Map.Empty.slice().updateUnsafe(_993_v7,_995_v9);
        _996_v10 = (_996_v10).update(_993_v7, _995_v9);
      } else {
        let _997_v11;
        _997_v11 = new BigNumber(-824);
        (globalState).f1 = (_997_v11).multipliedBy(new BigNumber(-559));
        let _998_v12;
        _998_v12 = new _dafny.CodePoint('u'.codePointAt(0));
        let _999_v13;
        _999_v13 = true;
        let _1000_v14;
        _1000_v14 = _dafny.MultiSet.fromElements(_999_v13, !(true));
        let _1001_v15;
        _1001_v15 = _dafny.Set.fromElements(new BigNumber((_1000_v14).cardinality()), _997_v11);
        let _1002_v16;
        _1002_v16 = _dafny.Map.Empty.slice().updateUnsafe(_999_v13,_997_v11);
        let _1003_v17;
        _1003_v17 = _module.D8.create_DC20(_998_v12, _module.__default.fm4(_1001_v15, globalState), _999_v13, _1002_v16, _999_v13);
        let _1004_v18;
        _1004_v18 = _dafny.Map.Empty.slice().updateUnsafe((_1003_v17).dtor_cf43,_998_v12);
        let _1005_v19;
        _1005_v19 = _dafny.Seq.of(_998_v12, new _dafny.CodePoint('g'.codePointAt(0)), _998_v12);
        _1004_v18 = (_1004_v18).update(_998_v12, (_1005_v19)[_module.__default.safeIndex(_997_v11, new BigNumber((_1005_v19).length))]);
        let _1006_v20;
        let _init30 = ((_1007_v12) => function (_1008_i0) {
          return _1007_v12;
        })(_998_v12);
        let _nw176 = Array((new BigNumber(16)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw176.length); _i0_30++) {
          _nw176[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1006_v20 = _nw176;
        _1006_v20 = _1006_v20;
        r0 = !(_999_v13) || (_999_v13);
        let _1009_v21;
        _1009_v21 = _dafny.Map.Empty.slice().updateUnsafe(_997_v11,false);
        let _rhs145 = _997_v11;
        let _rhs146 = _999_v13;
        let _rhs147 = _module.__default.fm4(_1001_v15, globalState);
        let _rhs148 = (new BigNumber((_1009_v21).length)).multipliedBy((_997_v11).plus(_997_v11));
        let _rhs149 = _module.__default.safeDivisionInt(_997_v11, _997_v11);
        let _lhs104 = globalState;
        _lhs104.f1 = _rhs145;
        _999_v13 = _rhs146;
        _999_v13 = _rhs147;
        r3 = _rhs148;
        r3 = _rhs149;
      }
      let _1010_v22;
      _1010_v22 = true;
      let _1011_v23;
      _1011_v23 = _dafny.MultiSet.fromElements(_1010_v22);
      let _1012_v24;
      _1012_v24 = _1011_v23;
      let _source11 = _1012_v24;
      let _1013___mcc_h0 = _source11;
      let _1014_cf41 = _1013___mcc_h0;
      let _1015_v25;
      _1015_v25 = _dafny.Seq.of(false, _1010_v22);
      let _1016_v26;
      let _1017_v27;
      let _1018_v28;
      let _1019_v29;
      let _out32;
      let _out33;
      let _out34;
      let _out35;
      let _outcollector8 = _module.__default.m0(_dafny.Seq.Concat(_1015_v25, _1015_v25), _1015_v25, globalState);
      _out32 = _outcollector8[0];
      _out33 = _outcollector8[1];
      _out34 = _outcollector8[2];
      _out35 = _outcollector8[3];
      _1016_v26 = _out32;
      _1017_v27 = _out33;
      _1018_v28 = _out34;
      _1019_v29 = _out35;
      let _1020_v30;
      let _nw177 = new _module.C0();
      _nw177.__ctor();
      _1020_v30 = _nw177;
      _1019_v29 = _1019_v29;
      let _1021_v31;
      _1021_v31 = new _dafny.CodePoint('t'.codePointAt(0));
      (globalState).f1 = (_dafny.ZERO).minus(_module.__default.fm19(_1021_v31, _1016_v26, globalState));
      let _1022_v32;
      _1022_v32 = new BigNumber(-895);
      let _1023_v33;
      _1023_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1022_v32,_1010_v22);
      let _1024_v34;
      let _nw178 = Array((new BigNumber(5)).toNumber());
      _nw178[(_dafny.ZERO).toNumber()] = !(_1010_v22);
      _nw178[(_dafny.ONE).toNumber()] = (((_1023_v33).contains(_1022_v32)) ? ((_1023_v33).get(_1022_v32)) : (_1010_v22));
      _nw178[(new BigNumber(2)).toNumber()] = (_1022_v32).isLessThan(new BigNumber((_dafny.Seq.of(_1022_v32)).length));
      _nw178[(new BigNumber(3)).toNumber()] = _1010_v22;
      _nw178[(new BigNumber(4)).toNumber()] = _1010_v22;
      _1024_v34 = _nw178;
      let _1025_v35;
      _1025_v35 = _dafny.Seq.UnicodeFromString("l");
      let _1026_v36;
      _1026_v36 = _dafny.MultiSet.fromElements(_1025_v35);
      let _index174 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
      (_1024_v34)[_index174] = (_1026_v36).IsDisjointFrom(_module.__default.fm26(globalState));
      let _index175 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
      (_1024_v34)[_index175] = _1010_v22;
      let _1027_v37;
      _1027_v37 = _dafny.Set.fromElements(_1022_v32, (_dafny.ZERO).minus(_1022_v32));
      let _1028_v38;
      _1028_v38 = _dafny.Seq.of(_module.__default.fm4(_1027_v37, globalState));
      let _1029_v39;
      _1029_v39 = _module.D9.create_DC21(_1028_v38);
      let _pat_let_tv22 = _1028_v38;
      let _source12 = function (_pat_let14_0) {
        return function (_1030_dt__update__tmp_h0) {
          return function (_pat_let15_0) {
            return function (_1031_dt__update_hcf48_h0) {
              return _module.D9.create_DC21(_1031_dt__update_hcf48_h0);
            }(_pat_let15_0);
          }(_pat_let_tv22);
        }(_pat_let14_0);
      }(_1029_v39);
      if (_source12.is_DC22) {
        let _1032___mcc_h1 = (_source12).cf49;
        let _1033___mcc_h2 = (_source12).cf50;
        let _1034___mcc_h3 = (_source12).cf51;
        let _1035_cf51 = _1034___mcc_h3;
        let _1036_cf50 = _1033___mcc_h2;
        let _1037_cf49 = _1032___mcc_h1;
        let _1038_v40;
        let _nw179 = new _module.C2();
        _nw179.__ctor(false, (_1037_cf49).plus(new BigNumber(506)), (_1011_v23).Difference(_1011_v23));
        _1038_v40 = _nw179;
        let _1039_v41;
        let _nw180 = Array((new BigNumber(12)).toNumber()).fill([]);
        _1039_v41 = _nw180;
        let _1040_v42;
        _1040_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1037_cf49,(((_1011_v23).contains((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))])) ? ((_1011_v23).get((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))])) : (_1022_v32)));
        let _1041_v43;
        _1041_v43 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1042_v44;
        _1042_v44 = _dafny.Set.fromElements(_1024_v34);
        let _1043_v45;
        _1043_v45 = _dafny.MultiSet.fromElements((_1038_v40).f13);
        let _1044_v46;
        _1044_v46 = _dafny.Map.Empty.slice().updateUnsafe((_1028_v38)[_module.__default.safeIndex(_1022_v32, new BigNumber((_1028_v38).length))],(_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))]);
        let _1045_v47;
        _1045_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1025_v35,_module.__default.fm27(_1041_v43, new BigNumber(-714), _1038_v40.f12, _module.__default.fm24(_1037_cf49, (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))], _1023_v33, _1037_cf49, globalState), globalState));
        let _1046_v48;
        let _nw181 = Array((new BigNumber(16)).toNumber());
        _nw181[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.of(_1036_cf50, true, _1010_v22)).length);
        _nw181[(_dafny.ONE).toNumber()] = (((_1040_v42).contains(new BigNumber((_dafny.Seq.of(_1041_v43)).length))) ? ((_1040_v42).get(new BigNumber((_dafny.Seq.of(_1041_v43)).length))) : (new BigNumber((_1042_v44).length)));
        _nw181[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.UnicodeFromString("vtmvqyk")).length)).minus(_1022_v32));
        _nw181[(new BigNumber(3)).toNumber()] = (new BigNumber((_1025_v35).length)).plus(new BigNumber((_1025_v35).length));
        _nw181[(new BigNumber(4)).toNumber()] = _1037_cf49;
        _nw181[(new BigNumber(5)).toNumber()] = new BigNumber(720);
        _nw181[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt((((_1043_v45).contains((_1038_v40).f13)) ? ((_1043_v45).get((_1038_v40).f13)) : (new BigNumber((_1044_v46).length))), _1037_cf49);
        _nw181[(new BigNumber(7)).toNumber()] = new BigNumber(-4);
        _nw181[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(_1037_cf49, _1037_cf49);
        _nw181[(new BigNumber(9)).toNumber()] = (_1038_v40).f13;
        _nw181[(new BigNumber(10)).toNumber()] = _1022_v32;
        _nw181[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(-734), (_dafny.ZERO).minus((_1038_v40).fm11(_1043_v45, globalState)));
        _nw181[(new BigNumber(12)).toNumber()] = (_1038_v40).f13;
        _nw181[(new BigNumber(13)).toNumber()] = new BigNumber((_1045_v47).length);
        _nw181[(new BigNumber(14)).toNumber()] = (_1022_v32).multipliedBy(new BigNumber((_1025_v35).length));
        _nw181[(new BigNumber(15)).toNumber()] = new BigNumber((_1028_v38).length);
        _1046_v48 = _nw181;
        let _index176 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1046_v48).length));
        (_1046_v48)[_index176] = (_dafny.ZERO).minus((_1038_v40).f13);
        let _1047_v49;
        _1047_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1010_v22,_module.__default.fm28((_1038_v40).fm11(_dafny.MultiSet.FromArray(_dafny.Seq.of(_1037_cf49)), globalState), globalState));
        let _1048_v50;
        _1048_v50 = _dafny.Map.Empty.slice().updateUnsafe(false,_1022_v32);
        let _1049_v51;
        _1049_v51 = _dafny.Map.Empty.slice().updateUnsafe((((_1047_v49).contains(true)) ? ((_1047_v49).get(true)) : (_1048_v50)),_1039_v41);
        let _1050_v52;
        _1050_v52 = _dafny.Seq.of(((_1036_cf50) ? (_1048_v50) : (_dafny.Map.Empty.slice().updateUnsafe(false,_1037_cf49))), _1048_v50);
        let _index177 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1046_v48).length));
        let _rhs150 = (((_1049_v51).contains(_1048_v50)) ? ((_1049_v51).get(_1048_v50)) : (_1039_v41));
        let _rhs151 = new BigNumber((_1050_v52).length);
        let _lhs105 = _1046_v48;
        let _lhs106 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1046_v48).length));
        _1039_v41 = _rhs150;
        _lhs105[_lhs106] = _rhs151;
        _1010_v22 = _1010_v22;
        let _1051_v53;
        let _nw182 = new _module.C0();
        _nw182.__ctor();
        _1051_v53 = _nw182;
        let _1052_v54;
        _1052_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1051_v53,false);
        let _1053_v55;
        _1053_v55 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1028_v38, _dafny.Seq.update(_1028_v38, _module.__default.safeIndex(new BigNumber((_1052_v54).length), new BigNumber((_1028_v38).length)), _1010_v22)),(_dafny.ZERO).minus((_1046_v48)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_1046_v48).length))]));
        _1053_v55 = (_1053_v55).update(_1028_v38, _1022_v32);
      } else if (_source12.is_DC23) {
        let _1054___mcc_h4 = (_source12).cf52;
        let _1055_cf52 = _1054___mcc_h4;
        r3 = _1055_cf52;
        let _1056_v56;
        _1056_v56 = _dafny.Seq.of(_dafny.Seq.Concat((_1029_v39).dtor_cf48, _dafny.Seq.of((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))])));
        let _rhs152 = _1010_v22;
        let _rhs153 = _1056_v56;
        r0 = _rhs152;
        _1056_v56 = _rhs153;
        let _1057_v57;
        _1057_v57 = _dafny.MultiSet.fromElements(_1022_v32, _1022_v32, _1022_v32);
        let _1058_v58;
        _1058_v58 = _dafny.Seq.of(_1055_cf52);
        if ((_dafny.MultiSet.FromArray(_1058_v58)).IsProperSubsetOf((_dafny.MultiSet.FromArray(_module.__default.fm20(_1055_cf52, _1055_cf52, globalState))).Intersect(_1057_v57))) {
          let _1059_v59;
          _1059_v59 = new _dafny.CodePoint('n'.codePointAt(0));
          let _1060_v60;
          _1060_v60 = _module.D3.create_DC8(_1059_v59, (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))], _1010_v22, _1025_v35);
          let _pat_let_tv23 = _1024_v34;
          let _pat_let_tv24 = _1024_v34;
          let _pat_let_tv25 = _1059_v59;
          _1060_v60 = function (_pat_let16_0) {
            return function (_1061_dt__update__tmp_h1) {
              return function (_pat_let17_0) {
                return function (_1062_dt__update_hcf18_h0) {
                  return function (_pat_let18_0) {
                    return function (_1063_dt__update_hcf17_h0) {
                      return _module.D3.create_DC8(_1063_dt__update_hcf17_h0, _1062_dt__update_hcf18_h0, (_1061_dt__update__tmp_h1).dtor_cf19, (_1061_dt__update__tmp_h1).dtor_cf20);
                    }(_pat_let18_0);
                  }(_pat_let_tv25);
                }(_pat_let17_0);
              }((_pat_let_tv24)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_pat_let_tv23).length))]);
            }(_pat_let16_0);
          }(_1060_v60);
          let _1064_v61;
          let _nw183 = new _module.C0();
          _nw183.__ctor();
          _1064_v61 = _nw183;
          let _1065_v62;
          _1065_v62 = _dafny.Map.Empty.slice().updateUnsafe((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))],_1055_cf52);
          let _1066_v63;
          _1066_v63 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-903),_1024_v34);
          let _1067_v64;
          let _nw184 = Array((new BigNumber(6)).toNumber());
          _nw184[(_dafny.ZERO).toNumber()] = (new BigNumber((_1065_v62).length)).plus(_1022_v32);
          _nw184[(_dafny.ONE).toNumber()] = new BigNumber((_1066_v63).length);
          _nw184[(new BigNumber(2)).toNumber()] = _1022_v32;
          _nw184[(new BigNumber(3)).toNumber()] = (((_1057_v57).contains(_1022_v32)) ? ((_1057_v57).get(_1022_v32)) : (_1055_cf52));
          _nw184[(new BigNumber(4)).toNumber()] = (new BigNumber((_1027_v37).length)).multipliedBy(_1055_cf52);
          _nw184[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_1022_v32);
          _1067_v64 = _nw184;
          let _index178 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_1067_v64).length));
          (_1067_v64)[_index178] = (((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))]) ? (_1022_v32) : ((_dafny.ZERO).minus(new BigNumber((_1011_v23).cardinality()))));
          let _index179 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_1067_v64).length));
          let _rhs154 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(956)), ((_1068_v59) => function (_1069_i1) {
            return _1068_v59;
          })(_1059_v59));
          let _rhs155 = _1022_v32;
          let _lhs107 = _1067_v64;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_1067_v64).length));
          _1025_v35 = _rhs154;
          _lhs107[_lhs108] = _rhs155;
          let _1070_v65;
          let _nw185 = new _module.C2();
          _nw185.__ctor(((_1067_v64)[_module.__default.safeIndex(new BigNumber(934), new BigNumber((_1067_v64).length))]).isLessThanOrEqualTo(_1022_v32), new BigNumber(373), (_1012_v24));
          _1070_v65 = _nw185;
          let _1071_v66;
          _1071_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1070_v65.f12,(_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))]);
          _1010_v22 = (_1010_v22) === ((((_1071_v66).contains(true)) ? ((_1071_v66).get(true)) : (_1070_v65.f12)));
        } else {
          (globalState).f1 = ((_dafny.ZERO).minus(_1022_v32)).multipliedBy(_1022_v32);
          let _1072_v67;
          _1072_v67 = _dafny.Set.fromElements(_1025_v35);
          r3 = (_dafny.ZERO).minus((_1022_v32).plus(((_1058_v58)[_module.__default.safeIndex(new BigNumber((_1072_v67).length), new BigNumber((_1058_v58).length))]).multipliedBy(_1055_cf52)));
          let _init31 = ((_1073_v34) => function (_1074_i2) {
            return (_1073_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1073_v34).length))];
          })(_1024_v34);
          let _nw186 = Array((_dafny.ONE).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw186.length); _i0_31++) {
            _nw186[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1024_v34 = _nw186;
          let _1075_v68;
          _1075_v68 = new _dafny.CodePoint('a'.codePointAt(0));
          r3 = _module.__default.fm19(_1075_v68, _1010_v22, globalState);
          let _index180 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
          let _index181 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
          let _rhs156 = (_1055_cf52).plus(_1022_v32);
          let _rhs157 = (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))];
          let _rhs158 = _1055_cf52;
          let _rhs159 = (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))];
          let _rhs160 = (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))];
          let _lhs109 = _1024_v34;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
          let _lhs111 = _1024_v34;
          let _lhs112 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
          _1022_v32 = _rhs156;
          _lhs109[_lhs110] = _rhs157;
          r2 = _rhs158;
          _lhs111[_lhs112] = _rhs159;
          _1010_v22 = _rhs160;
        }
        let _1076_v69;
        let _init32 = ((_1077_v58) => function (_1078_i3) {
          return _1077_v58;
        })(_1058_v58);
        let _nw187 = Array((new BigNumber(8)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw187.length); _i0_32++) {
          _nw187[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1076_v69 = _nw187;
        let _1079_v70;
        _1079_v70 = _module.D6.create_DC17(new BigNumber(276), _1025_v35, _1010_v22);
        let _1080_v71;
        _1080_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1028_v38,_1010_v22);
        let _1081_v72;
        _1081_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1055_cf52,_1080_v71);
        let _rhs161 = _dafny.Seq.UnicodeFromString("csulwnn");
        let _rhs162 = (_1079_v70).dtor_cf39;
        let _rhs163 = _1076_v69;
        let _rhs164 = (new BigNumber((_1058_v58).length)).isLessThan(new BigNumber((_1058_v58).length));
        let _rhs165 = !(((((_1081_v72).contains(_1022_v32)) ? ((_1081_v72).get(_1022_v32)) : (_1080_v71))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1028_v38,(_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))]))).contains((((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))]) ? (_dafny.Seq.of(_1010_v22, _1010_v22, true, (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))], false)) : (_1028_v38)));
        _1025_v35 = _rhs161;
        _1025_v35 = _rhs162;
        _1076_v69 = _rhs163;
        r0 = _rhs164;
        r0 = _rhs165;
      } else if (_source12.is_DC24) {
        let _1082___mcc_h5 = (_source12).cf53;
        let _1083_cf53 = _1082___mcc_h5;
        let _1084_v73;
        let _nw188 = new _module.C1();
        _nw188.__ctor(_module.__default.fm4(_dafny.Set.fromElements(_1022_v32, new BigNumber((_1028_v38).length)), globalState), _1010_v22, (_1011_v23).Intersect(_dafny.MultiSet.FromArray(_1028_v38)));
        _1084_v73 = _nw188;
        let _index182 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
        (_1024_v34)[_index182] = (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))];
        r0 = (_1084_v73).f15;
        let _1085_v74;
        _1085_v74 = _dafny.Seq.of((_1027_v37).Intersect(_1027_v37), _1027_v37);
        _1027_v37 = (_1085_v74)[_module.__default.safeIndex(new BigNumber((_1028_v38).length), new BigNumber((_1085_v74).length))];
      } else {
        let _1086___mcc_h6 = (_source12).cf48;
        let _1087_cf48 = _1086___mcc_h6;
        r0 = _1010_v22;
        _1025_v35 = _dafny.Seq.UnicodeFromString("rd");
        let _1088_v75;
        _1088_v75 = new _dafny.CodePoint('o'.codePointAt(0));
        _1010_v22 = !(!_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_1025_v35, _module.__default.safeIndex(_1022_v32, new BigNumber((_1025_v35).length)), _1088_v75), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-326)), function (_1089_i4) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })), _module.__default.safeIndex(_1022_v32, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1025_v35, _module.__default.safeIndex(_1022_v32, new BigNumber((_1025_v35).length)), _1088_v75), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-326)), function (_1090_i4) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }))).length)), _1088_v75), _1088_v75));
        let _1091_v76;
        _1091_v76 = _dafny.Map.Empty.slice().updateUnsafe((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))],_module.__default.fm4(_1027_v37, globalState));
        let _1092_v77;
        _1092_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1022_v32,_1091_v76);
        let _1093_v78;
        _1093_v78 = _dafny.Set.fromElements((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))]);
        _1091_v76 = (((_1092_v77).contains(_module.__default.safeModuloInt((_this).fm9(_1025_v35, new BigNumber((_1093_v78).length), globalState), new BigNumber(990)))) ? ((_1092_v77).get(_module.__default.safeModuloInt((_this).fm9(_1025_v35, new BigNumber((_1093_v78).length), globalState), new BigNumber(990)))) : ((_module.__default.fm18((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))], _1023_v33, globalState)).update(_1010_v22, (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))])));
      }
      if (_1010_v22) {
        _1022_v32 = _1022_v32;
        let _1094_v79;
        _1094_v79 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1022_v32), _1022_v32, _1022_v32);
        _1094_v79 = ((_module.__default.fm29(new BigNumber(314), _1010_v22, false, globalState)).update(_1022_v32, _module.__default.abs(new BigNumber((_1025_v35).length)))).Union(((_1010_v22) ? (_1094_v79) : (_1094_v79)));
        let _1095_v80;
        _1095_v80 = new _dafny.CodePoint('b'.codePointAt(0));
        let _1096_v81;
        _1096_v81 = _module.D3.create_DC8(_1095_v80, true, _1010_v22, _module.__default.fm2(!(!(!(_1010_v22))), _1022_v32, _1095_v80, _1022_v32, globalState));
        let _1097_v82;
        _1097_v82 = _dafny.Seq.of(_1024_v34);
        let _1098_v83;
        _1098_v83 = _dafny.Set.fromElements(_1010_v22, (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))], (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))], (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))]);
        (globalState).f1 = (_this).fm9(_dafny.Seq.update(_dafny.Seq.update((_1096_v81).dtor_cf20, _module.__default.safeIndex(_1022_v32, new BigNumber(((_1096_v81).dtor_cf20).length)), _module.__default.fm1(globalState)), _module.__default.safeIndex(new BigNumber((_1097_v82).length), new BigNumber((_dafny.Seq.update((_1096_v81).dtor_cf20, _module.__default.safeIndex(_1022_v32, new BigNumber(((_1096_v81).dtor_cf20).length)), _module.__default.fm1(globalState))).length)), _1095_v80), (new BigNumber((_1098_v83).length)).plus(_1022_v32), globalState);
        let _1099_v84;
        let _nw189 = new _module.C3();
        _nw189.__ctor(_1010_v22);
        _1099_v84 = _nw189;
        let _1100_v85;
        _1100_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1099_v84,(_module.__default.fm4(_1027_v37, globalState)) === (false));
        let _index183 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
        (_1024_v34)[_index183] = (((_1100_v85).contains(_1099_v84)) ? ((_1100_v85).get(_1099_v84)) : (_1010_v22));
        let _index184 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
        (_1024_v34)[_index184] = !(_module.__default.fm4(_1027_v37, globalState));
      } else {
        let _1101_v86;
        let _nw190 = Array((new BigNumber(3)).toNumber());
        _nw190[(_dafny.ZERO).toNumber()] = _1025_v35;
        _nw190[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("rxgwn"), _module.__default.safeIndex(new BigNumber((_1011_v23).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("rxgwn")).length)), (_1025_v35)[_module.__default.safeIndex(_1022_v32, new BigNumber((_1025_v35).length))]), _1025_v35);
        _nw190[(new BigNumber(2)).toNumber()] = _1025_v35;
        _1101_v86 = _nw190;
        let _1102_v87;
        let _nw191 = Array((new BigNumber(21)).toNumber()).fill([]);
        _1102_v87 = _nw191;
        let _1103_v88;
        let _init33 = ((_1104_v37) => function (_1105_i5) {
          return _1104_v37;
        })(_1027_v37);
        let _nw192 = Array((new BigNumber(24)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw192.length); _i0_33++) {
          _nw192[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1103_v88 = _nw192;
        let _index185 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_1103_v88).length));
        (_1103_v88)[_index185] = _1027_v37;
        let _index186 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
        let _index187 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_1103_v88).length));
        let _rhs166 = (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))];
        let _rhs167 = _1101_v86;
        let _rhs168 = _1022_v32;
        let _rhs169 = _1102_v87;
        let _rhs170 = _1027_v37;
        let _lhs113 = _1024_v34;
        let _lhs114 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length));
        let _lhs115 = globalState;
        let _lhs116 = _1103_v88;
        let _lhs117 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_1103_v88).length));
        _lhs113[_lhs114] = _rhs166;
        _1101_v86 = _rhs167;
        _lhs115.f1 = _rhs168;
        _1102_v87 = _rhs169;
        _lhs116[_lhs117] = _rhs170;
        let _1106_v89;
        _1106_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1024_v34,_1022_v32);
        _1106_v89 = (_1106_v89).update(_1024_v34, (_dafny.ZERO).minus((_1022_v32).minus(_1022_v32)));
        (globalState).f1 = _1022_v32;
        let _1107_v90;
        _1107_v90 = _dafny.Set.fromElements(_1027_v37);
        let _1108_v91;
        let _nw193 = Array((new BigNumber(10)).toNumber());
        _nw193[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_1022_v32);
        _nw193[(_dafny.ONE).toNumber()] = _1022_v32;
        _nw193[(new BigNumber(2)).toNumber()] = _1022_v32;
        _nw193[(new BigNumber(3)).toNumber()] = _1022_v32;
        _nw193[(new BigNumber(4)).toNumber()] = _1022_v32;
        _nw193[(new BigNumber(5)).toNumber()] = new BigNumber((_1106_v89).length);
        _nw193[(new BigNumber(6)).toNumber()] = new BigNumber(((((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))]) ? (_1107_v90) : (_1107_v90))).length);
        _nw193[(new BigNumber(7)).toNumber()] = _1022_v32;
        _nw193[(new BigNumber(8)).toNumber()] = _1022_v32;
        _nw193[(new BigNumber(9)).toNumber()] = (new BigNumber(847)).plus(_1022_v32);
        _1108_v91 = _nw193;
        _1108_v91 = _1108_v91;
        let _index188 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1101_v86).length));
        (_1101_v86)[_index188] = _1025_v35;
        let _index189 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1101_v86).length));
        (_1101_v86)[_index189] = _1025_v35;
      }
      let _1109_v93;
      let _init34 = ((_1110_v32) => function (_1111_i7) {
        return function () {
          let _coll38 = new _dafny.Set();
          for (const _compr_38 of _dafny.IntegerRange(new BigNumber(555), new BigNumber(410))) {
            let _1112_v92 = _compr_38;
            if (((new BigNumber(555)).isLessThanOrEqualTo(_1112_v92)) && ((_1112_v92).isLessThan(new BigNumber(410)))) {
              _coll38.add((_1112_v92).multipliedBy(_1110_v32));
            }
          }
          return _coll38;
        }();
      })(_1022_v32);
      let _nw194 = Array((new BigNumber(8)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw194.length); _i0_34++) {
        _nw194[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _1109_v93 = _nw194;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1109_v93).length))) {
        let _1113_i6 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1113_i6)) && ((_1113_i6).isLessThan(new BigNumber((_1109_v93).length))))) {
          (_1109_v93)[(_1113_i6)] = (_1027_v37).Intersect(_1027_v37);
        }
      }
      let _1114_v94;
      _1114_v94 = _module.D0.create_DC0(new _dafny.CodePoint('e'.codePointAt(0)));
      let _pat_let_tv26 = _1024_v34;
      let _pat_let_tv27 = _1024_v34;
      let _pat_let_tv28 = _1025_v35;
      r0 = !(function (_source13) {
        if (_source13.is_DC1) {
          let _1115___mcc_h7 = (_source13).cf1;
          let _1116___mcc_h8 = (_source13).cf2;
          let _1117_cf2 = _1116___mcc_h8;
          let _1118_cf1 = _1115___mcc_h7;
          return !(_1118_cf1);
        } else if (_source13.is_DC0) {
          let _1119___mcc_h9 = (_source13).cf0;
          let _1120_cf0 = _1119___mcc_h9;
          return !(!((_pat_let_tv27)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_pat_let_tv26).length))]));
        } else {
          let _1121___mcc_h10 = (_source13).cf3;
          let _1122_cf3 = _1121___mcc_h10;
          return !_dafny.areEqual(_dafny.Seq.UnicodeFromString("xq"), _pat_let_tv28);
        }
      }(_module.D0.create_DC2(_1114_v94)));
      let _nw195 = Array((new BigNumber(4)).toNumber());
      _nw195[(_dafny.ZERO).toNumber()] = _1024_v34;
      _nw195[(_dafny.ONE).toNumber()] = _1024_v34;
      _nw195[(new BigNumber(2)).toNumber()] = _1024_v34;
      _nw195[(new BigNumber(3)).toNumber()] = _1024_v34;
      r1 = _nw195;
      r2 = _1022_v32;
      let _1123_v95;
      _1123_v95 = new _dafny.CodePoint('w'.codePointAt(0));
      let _1124_v96;
      _1124_v96 = _dafny.Map.Empty.slice().updateUnsafe((_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))],_module.__default.fm19(_1123_v95, (_1024_v34)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_1024_v34).length))], globalState));
      r3 = (_1022_v32).minus(_module.__default.safeModuloInt(_1022_v32, (((_1124_v96).contains(true)) ? ((_1124_v96).get(true)) : (_1022_v32))));
      return [r0, r1, r2, r3];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f10 = _dafny.MultiSet.Empty;
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1, _module.T2];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f11, f10) {
      let _this = this;
      (_this)._f11 = f11;
      (_this)._f10 = f10;
      return;
    }
    fm5(p0, p1, p2, globalState) {
      let _this = this;
      return false;
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), function (_1125_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("pqoiu"))).length);
    };
    fm7(p0, globalState) {
      let _this = this;
      return false;
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _1126_v0;
      _1126_v0 = _dafny.Seq.UnicodeFromString("wbyaw");
      let _1127_v1;
      _1127_v1 = _module.D0.create_DC0(p0);
      let _1128_v2;
      _1128_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _1129_v3;
      _1129_v3 = _dafny.MultiSet.fromElements(new BigNumber(721));
      let _1130_v4;
      _1130_v4 = _dafny.Set.fromElements(new BigNumber((_1128_v2).length), (_module.D0.create_DC1((_this).f11, new BigNumber((_1129_v3).cardinality()))).dtor_cf2);
      let _1131_v5;
      _1131_v5 = _dafny.Seq.of((_this).f11);
      let _1132_v6;
      _1132_v6 = _dafny.Seq.of(p1, p1, p1);
      let _1133_v7;
      _1133_v7 = _dafny.Set.fromElements((_this).f11);
      let _1134_v8;
      let _init35 = ((_1135_p2) => function (_1136_i0) {
        return (_1136_i0).minus((_1135_p2).dtor_cf2);
      })(p2);
      let _nw196 = Array((new BigNumber(23)).toNumber());
      for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw196.length); _i0_35++) {
        _nw196[_i0_35] = _init35(new BigNumber(_i0_35));
      }
      _1134_v8 = _nw196;
      let _1137_v9;
      _1137_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1134_v8,p1);
      let _1138_v10;
      let _nw197 = Array((new BigNumber(19)).toNumber());
      _nw197[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_this).fm6((_this).f11, new BigNumber(657), globalState), new BigNumber(641));
      _nw197[(_dafny.ONE).toNumber()] = p1;
      _nw197[(new BigNumber(2)).toNumber()] = new BigNumber(-842);
      _nw197[(new BigNumber(3)).toNumber()] = p1;
      _nw197[(new BigNumber(4)).toNumber()] = p1;
      _nw197[(new BigNumber(5)).toNumber()] = p1;
      _nw197[(new BigNumber(6)).toNumber()] = (_this).fm6(_module.__default.fm4(_1130_v4, globalState), p1, globalState);
      _nw197[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.update(_1131_v5, _module.__default.safeIndex(p1, new BigNumber((_1131_v5).length)), (_this).f11)).length), (_1132_v6)[_module.__default.safeIndex(p1, new BigNumber((_1132_v6).length))]);
      _nw197[(new BigNumber(8)).toNumber()] = p1;
      _nw197[(new BigNumber(9)).toNumber()] = new BigNumber((_1133_v7).length);
      _nw197[(new BigNumber(10)).toNumber()] = (p1).multipliedBy((_dafny.ZERO).minus(p1));
      _nw197[(new BigNumber(11)).toNumber()] = (_this).fm6(false, p1, globalState);
      _nw197[(new BigNumber(12)).toNumber()] = new BigNumber((_1137_v9).length);
      _nw197[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1126_v0, _1126_v0)).length);
      _nw197[(new BigNumber(14)).toNumber()] = (p1).multipliedBy(new BigNumber(407));
      _nw197[(new BigNumber(15)).toNumber()] = p1;
      _nw197[(new BigNumber(16)).toNumber()] = p1;
      _nw197[(new BigNumber(17)).toNumber()] = new BigNumber(954);
      _nw197[(new BigNumber(18)).toNumber()] = p1;
      _1138_v10 = _nw197;
      let _1139_v11;
      _1139_v11 = _module.D0.create_DC1((_this).f11, p1);
      let _1140_v12;
      _1140_v12 = _module.D0.create_DC2(_1139_v11);
      let _1141_v13;
      _1141_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      let _rhs171 = _dafny.Seq.Concat(_1126_v0, _1126_v0);
      let _rhs172 = _1127_v1;
      let _rhs173 = _1134_v8;
      let _rhs174 = _1140_v12;
      let _rhs175 = _1141_v13;
      _1126_v0 = _rhs171;
      _1127_v1 = _rhs172;
      _1138_v10 = _rhs173;
      _1140_v12 = _rhs174;
      _1141_v13 = _rhs175;
      if (_dafny.Seq.IsPrefixOf((((p2).dtor_cf1) ? (_1132_v6) : (_1132_v6)), _module.__default.fm8(_module.__default.fm4(_1130_v4, globalState), globalState))) {
        let _1142_v14;
        _1142_v14 = false;
        _1142_v14 = !((p1).isLessThan(new BigNumber(423)));
        let _1143_v15;
        let _1144_v16;
        let _1145_v17;
        let _1146_v18;
        let _out36;
        let _out37;
        let _out38;
        let _out39;
        let _outcollector9 = _module.__default.m0(_dafny.Seq.of(true, !(!(true))), _1131_v5, globalState);
        _out36 = _outcollector9[0];
        _out37 = _outcollector9[1];
        _out38 = _outcollector9[2];
        _out39 = _outcollector9[3];
        _1143_v15 = _out36;
        _1144_v16 = _out37;
        _1145_v17 = _out38;
        _1146_v18 = _out39;
        _1146_v18 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((new BigNumber((_1126_v0).length)).minus(_1146_v18), _1146_v18));
        _1142_v14 = !(_1146_v18).isEqualTo(p1);
        let _1147_v19;
        _1147_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1146_v18,(_this).f11);
        (globalState).f1 = (_dafny.ZERO).minus((_1132_v6)[_module.__default.safeIndex((_1146_v18).multipliedBy(new BigNumber((_1147_v19).length)), new BigNumber((_1132_v6).length))]);
      } else {
        let _1148_v20;
        let _nw198 = new _module.C3();
        _nw198.__ctor((_this).f11);
        _1148_v20 = _nw198;
        let _1149_v21;
        _1149_v21 = true;
        let _rhs176 = _1149_v21;
        let _rhs177 = (_this).f11;
        _1149_v21 = _rhs176;
        _1149_v21 = _rhs177;
        if (_1149_v21) {
          let _1150_v22;
          _1150_v22 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f11);
          let _1151_v24;
          _1151_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of _dafny.IntegerRange(new BigNumber(213), new BigNumber(198))) {
              let _1152_v23 = _compr_39;
              if (((new BigNumber(213)).isLessThanOrEqualTo(_1152_v23)) && ((_1152_v23).isLessThan(new BigNumber(198)))) {
                _coll39.push([_module.__default.safeDivisionInt(_1152_v23, p1),p1]);
              }
            }
            return _coll39;
          }()).length),_1149_v21);
          let _1153_v25;
          _1153_v25 = _dafny.Set.fromElements((_1150_v22).update((_1148_v20).f16, false), _module.__default.fm18(true, _1151_v24, globalState), ((_1149_v21) ? (_1150_v22) : (_1150_v22)));
          (globalState).f1 = new BigNumber((_1153_v25).length);
          let _1154_v26;
          let _init36 = ((_1155_v7) => function (_1156_i1) {
            return (_1155_v7).Union(_dafny.Set.fromElements(false));
          })(_1133_v7);
          let _nw199 = Array((new BigNumber(9)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw199.length); _i0_36++) {
            _nw199[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1154_v26 = _nw199;
          let _index190 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_1154_v26).length));
          (_1154_v26)[_index190] = _dafny.Set.fromElements((_1148_v20).f16, (_1148_v20).f16, (_1148_v20).f16);
          let _index191 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_1154_v26).length));
          (_1154_v26)[_index191] = _1133_v7;
          (globalState).f1 = (new BigNumber(483)).minus(new BigNumber((_1126_v0).length));
          let _1157_v27;
          let _nw200 = Array((new BigNumber(15)).toNumber()).fill(false);
          _1157_v27 = _nw200;
          let _1158_v28;
          _1158_v28 = _dafny.Seq.of(_1157_v27, _1157_v27);
          let _1159_v29;
          _1159_v29 = _dafny.Seq.of((((_1148_v20).f16) ? (_1157_v27) : (_1157_v27)), (((_1148_v20).f16) ? (_1157_v27) : (_1157_v27)), _1157_v27, (_1158_v28)[_module.__default.safeIndex(p1, new BigNumber((_1158_v28).length))]);
          let _rhs178 = p1;
          let _rhs179 = (_1158_v28)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_1158_v28).length))];
          let _rhs180 = _1134_v8;
          let _rhs181 = _1126_v0;
          let _lhs118 = globalState;
          let _lhs119 = globalState;
          _lhs118.f1 = _rhs178;
          _lhs119.f5 = _rhs179;
          _1134_v8 = _rhs180;
          _1126_v0 = _rhs181;
          (globalState).f1 = new BigNumber((_module.__default.fm14((_1148_v20).f16, globalState)).length);
        } else {
          let _1160_v30;
          let _init37 = ((_1161_v20) => function (_1162_i2) {
            return (_1161_v20).f16;
          })(_1148_v20);
          let _nw201 = Array((new BigNumber(20)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw201.length); _i0_37++) {
            _nw201[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1160_v30 = _nw201;
          let _1163_v31;
          _1163_v31 = _module.D6.create_DC17(new BigNumber((_1131_v5).length), _1126_v0, true);
          let _index192 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1160_v30).length));
          (_1160_v30)[_index192] = (_1149_v21) && ((_1163_v31).dtor_cf40);
          let _index193 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1160_v30).length));
          (_1160_v30)[_index193] = (_1148_v20).f16;
          let _1164_v32;
          _1164_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1149_v21,p1);
          let _1165_v33;
          _1165_v33 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f11);
          let _1166_v34;
          _1166_v34 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("xrghkd"));
          let _1167_v35;
          let _nw202 = new _module.C2();
          _nw202.__ctor((_this).f11, p1, (_this).f10);
          _1167_v35 = _nw202;
          let _1168_v36;
          _1168_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1148_v20).f16,(_1148_v20).f16)).length), new BigNumber((_1165_v33).length), new BigNumber((_1166_v34).length)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1148_v20).f16,(_1148_v20).f16)).length), new BigNumber((_1165_v33).length), new BigNumber((_1166_v34).length))).length)), p1),_1167_v35);
          let _index194 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1160_v30).length));
          let _index195 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1160_v30).length));
          let _rhs182 = (((_1149_v21) ? (_1164_v32) : (_1164_v32))).equals((_1164_v32).Merge(_1164_v32));
          let _rhs183 = (_1168_v36).equals(_dafny.Map.Empty.slice().updateUnsafe(_1132_v6,_1167_v35));
          let _lhs120 = _1160_v30;
          let _lhs121 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1160_v30).length));
          let _lhs122 = _1160_v30;
          let _lhs123 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1160_v30).length));
          _lhs120[_lhs121] = _rhs182;
          _lhs122[_lhs123] = _rhs183;
          let _index196 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1160_v30).length));
          let _rhs184 = ((false) === ((_1160_v30)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_1160_v30).length))])) && ((_1148_v20).f16);
          let _rhs185 = new BigNumber(730);
          let _lhs124 = _1160_v30;
          let _lhs125 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1160_v30).length));
          let _lhs126 = globalState;
          _lhs124[_lhs125] = _rhs184;
          _lhs126.f1 = _rhs185;
          let _1169_v37;
          let _nw203 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1169_v37 = _nw203;
          let _index197 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_1169_v37).length));
          (_1169_v37)[_index197] = new _dafny.CodePoint('s'.codePointAt(0));
          let _1170_v38;
          _1170_v38 = _module.D8.create_DC20(p0, (_this).f11, _1149_v21, _1164_v32, _1149_v21);
          let _index198 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_1169_v37).length));
          let _rhs186 = ((((_1164_v32).contains((_this).f11)) ? ((_1164_v32).get((_this).f11)) : (p1))).minus(_module.__default.safeDivisionInt(p1, (_1167_v35).f13));
          let _rhs187 = (((_1148_v20).f16) ? (p0) : (((_1167_v35.f12) ? (p0) : ((_1170_v38).dtor_cf43))));
          let _rhs188 = (_1167_v35).f13;
          let _lhs127 = globalState;
          let _lhs128 = _1169_v37;
          let _lhs129 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_1169_v37).length));
          let _lhs130 = globalState;
          _lhs127.f1 = _rhs186;
          _lhs128[_lhs129] = _rhs187;
          _lhs130.f1 = _rhs188;
          (_1167_v35).f12 = (_1160_v30)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_1160_v30).length))];
          let _index199 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1160_v30).length));
          (_1160_v30)[_index199] = _1149_v21;
        }
        if ((_this).f11) {
          _1131_v5 = _1131_v5;
          (globalState).f1 = (p1).multipliedBy(p1);
          _1126_v0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(260)), ((_1171_p0) => function (_1172_i3) {
            return _1171_p0;
          })(p0)), _1126_v0), _1126_v0);
          let _1173_v39;
          _1173_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1134_v8,!(_1149_v21) || ((_1148_v20).f16));
          _1173_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1134_v8,(_1148_v20).f16);
          let _1174_v40;
          _1174_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1149_v21,p1);
          let _index200 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_1134_v8).length));
          (_1134_v8)[_index200] = new BigNumber(((_1174_v40).update((_this).f11, p1)).length);
          let _index201 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_1134_v8).length));
          (_1134_v8)[_index201] = p1;
        } else {
          let _1175_v41;
          let _nw204 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1175_v41 = _nw204;
          (globalState).f5 = _1175_v41;
          let _index202 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_1175_v41).length));
          (_1175_v41)[_index202] = (_1148_v20).f16;
          let _index203 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_1175_v41).length));
          (_1175_v41)[_index203] = _1149_v21;
          let _index204 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_1175_v41).length));
          let _index205 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_1175_v41).length));
          let _rhs189 = (new BigNumber(322)).minus(p1);
          let _rhs190 = (_1148_v20).f16;
          let _rhs191 = (new BigNumber(618)).plus(p1);
          let _rhs192 = p1;
          let _rhs193 = (_this).fm5((_1133_v7).Union(_1133_v7), p1, (_1132_v6)[_module.__default.safeIndex(p1, new BigNumber((_1132_v6).length))], globalState);
          let _lhs131 = globalState;
          let _lhs132 = _1175_v41;
          let _lhs133 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_1175_v41).length));
          let _lhs134 = globalState;
          let _lhs135 = globalState;
          let _lhs136 = _1175_v41;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_1175_v41).length));
          _lhs131.f1 = _rhs189;
          _lhs132[_lhs133] = _rhs190;
          _lhs134.f1 = _rhs191;
          _lhs135.f1 = _rhs192;
          _lhs136[_lhs137] = _rhs193;
          (globalState).f1 = _module.__default.safeModuloInt(p1, p1);
          _1149_v21 = (_1148_v20).f16;
        }
        let _1176_v42;
        _1176_v42 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1148_v20).f16);
        _1176_v42 = (_1176_v42).Merge(_1176_v42);
      }
      if ((_this).f11) {
        let _1177_v43;
        let _nw205 = Array((new BigNumber(12)).toNumber());
        _nw205[(_dafny.ZERO).toNumber()] = (_this).f11;
        _nw205[(_dafny.ONE).toNumber()] = (_this).f11;
        _nw205[(new BigNumber(2)).toNumber()] = (_this).f11;
        _nw205[(new BigNumber(3)).toNumber()] = (_this).f11;
        _nw205[(new BigNumber(4)).toNumber()] = (_this).f11;
        _nw205[(new BigNumber(5)).toNumber()] = (_this).f11;
        _nw205[(new BigNumber(6)).toNumber()] = (_this).f11;
        _nw205[(new BigNumber(7)).toNumber()] = !((_this).f11);
        _nw205[(new BigNumber(8)).toNumber()] = false;
        _nw205[(new BigNumber(9)).toNumber()] = (_this).f11;
        _nw205[(new BigNumber(10)).toNumber()] = (_this).f11;
        _nw205[(new BigNumber(11)).toNumber()] = true;
        _1177_v43 = _nw205;
        let _1178_v44;
        let _nw206 = Array((new BigNumber(5)).toNumber());
        _nw206[(_dafny.ZERO).toNumber()] = (((_this).f11) ? (_1177_v43) : (_1177_v43));
        _nw206[(_dafny.ONE).toNumber()] = _1177_v43;
        _nw206[(new BigNumber(2)).toNumber()] = _1177_v43;
        _nw206[(new BigNumber(3)).toNumber()] = _1177_v43;
        _nw206[(new BigNumber(4)).toNumber()] = _1177_v43;
        _1178_v44 = _nw206;
        let _index206 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_1178_v44).length));
        (_1178_v44)[_index206] = _1177_v43;
        let _index207 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_1178_v44).length));
        (_1178_v44)[_index207] = _1177_v43;
        let _index208 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1134_v8).length));
        (_1134_v8)[_index208] = p1;
        let _1179_v45;
        _1179_v45 = false;
        let _index209 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1134_v8).length));
        let _rhs194 = _dafny.Seq.Concat(_1132_v6, _dafny.Seq.Concat(_dafny.Seq.of(p1, p1), _1132_v6));
        let _rhs195 = p1;
        let _rhs196 = (_this).fm6(_1179_v45, p1, globalState);
        let _rhs197 = (!(false)) && ((_this).f11);
        let _lhs138 = globalState;
        let _lhs139 = _1134_v8;
        let _lhs140 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1134_v8).length));
        _1132_v6 = _rhs194;
        _lhs138.f1 = _rhs195;
        _lhs139[_lhs140] = _rhs196;
        _1179_v45 = _rhs197;
        let _arr0 = (_1178_v44)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_1178_v44).length))];
        let _index210 = _module.__default.safeIndex(new BigNumber(877), new BigNumber(((_1178_v44)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_1178_v44).length))]).length));
        _arr0[_index210] = _1179_v45;
        let _arr1 = (_1178_v44)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_1178_v44).length))];
        let _index211 = _module.__default.safeIndex(new BigNumber(877), new BigNumber(((_1178_v44)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_1178_v44).length))]).length));
        _arr1[_index211] = (p1).isLessThanOrEqualTo(_module.__default.fm19(p0, _1179_v45, globalState));
        let _arr2 = (_1178_v44)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_1178_v44).length))];
        let _index212 = _module.__default.safeIndex(new BigNumber(877), new BigNumber(((_1178_v44)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_1178_v44).length))]).length));
        _arr2[_index212] = (_this).f11;
        let _1180_v46;
        _1180_v46 = new _dafny.CodePoint('x'.codePointAt(0));
        _1180_v46 = p0;
      } else {
        let _1181_v47;
        _1181_v47 = true;
        _1181_v47 = (p1).isEqualTo(p1);
        if (_1181_v47) {
          (globalState).f1 = ((!((_this).f11) || (_1181_v47)) ? ((p1).multipliedBy(p1)) : (p1));
          _1181_v47 = !(false);
          _1181_v47 = !(_1181_v47);
          (globalState).f1 = p1;
          let _1182_v48;
          let _nw207 = new _module.C0();
          _nw207.__ctor();
          _1182_v48 = _nw207;
          _1182_v48 = _1182_v48;
        } else {
          let _1183_v49;
          let _nw208 = new _module.C4();
          _nw208.__ctor();
          _1183_v49 = _nw208;
          let _1184_v50;
          _1184_v50 = _module.D3.create_DC8(new _dafny.CodePoint('m'.codePointAt(0)), (_this).f11, _module.__default.fm4(_1130_v4, globalState), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jhadxaocd"), _dafny.Seq.update(_1126_v0, _module.__default.safeIndex(p1, new BigNumber((_1126_v0).length)), new _dafny.CodePoint('k'.codePointAt(0)))));
          _1184_v50 = _1184_v50;
          (globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p1, (_dafny.ZERO).minus(new BigNumber((((_this).f10).Difference(((_this).f10).update(_1181_v47, _module.__default.abs(p1)))).cardinality()))));
          let _1185_v51;
          let _nw209 = new _module.C2();
          _nw209.__ctor(!(_1130_v4).contains(new BigNumber((_1132_v6).length)), (p1).multipliedBy(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("kwiask"), _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1126_v0, _module.__default.safeIndex(p1, new BigNumber((_1126_v0).length)), p0)).length), new BigNumber((_dafny.Seq.UnicodeFromString("kwiask")).length)), new _dafny.CodePoint('e'.codePointAt(0)))).length)), (_this).f10);
          _1185_v51 = _nw209;
          (globalState).f1 = (_1185_v51).f13;
        }
        let _1186_v52;
        let _nw210 = new _module.C2();
        _nw210.__ctor((_1129_v3).IsProperSubsetOf(_1129_v3), p1, (_this).f10);
        _1186_v52 = _nw210;
        let _1187_v53;
        let _nw211 = Array((new BigNumber(3)).toNumber());
        _nw211[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1126_v0, _1126_v0), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_1126_v0, _1126_v0)).length)), new _dafny.CodePoint('d'.codePointAt(0)));
        _nw211[(_dafny.ONE).toNumber()] = _module.__default.fm2(_1186_v52.f12, new BigNumber((_1130_v4).length), p0, (_1186_v52).f13, globalState);
        _nw211[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_1126_v0, _module.__default.safeIndex((_1186_v52).f13, new BigNumber((_1126_v0).length)), p0);
        _1187_v53 = _nw211;
        let _index213 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_1187_v53).length));
        (_1187_v53)[_index213] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xonhdkbwi"), _1126_v0);
        let _index214 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_1187_v53).length));
        (_1187_v53)[_index214] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(187)), function (_1188_i4) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        });
        let _1189_v54;
        _1189_v54 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f11),(((_this).f11) ? (_1138_v10) : (_1138_v10)));
        _1138_v10 = (((_1189_v54).contains(_dafny.areEqual(_1131_v5, _1131_v5))) ? ((_1189_v54).get(_dafny.areEqual(_1131_v5, _1131_v5))) : (_1134_v8));
      }
      let _1190_v55;
      _1190_v55 = false;
      let _1191_v56;
      let _nw212 = Array((new BigNumber(3)).toNumber()).fill(false);
      _1191_v56 = _nw212;
      let _1192_v57;
      _1192_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1191_v56,(_this).f11);
      let _1193_v58;
      _1193_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,p0);
      let _1194_v59;
      _1194_v59 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1193_v58).length)).plus(p1),_dafny.Seq.update(_1126_v0, _module.__default.safeIndex(p1, new BigNumber((_1126_v0).length)), p0));
      let _1195_v60;
      _1195_v60 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(_1130_v4, globalState),_1126_v0);
      let _pat_let_tv29 = p1;
      let _rhs198 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(((_1190_v55) ? (_dafny.Seq.of(p1)) : (_1132_v6)), _module.__default.safeIndex(p1, new BigNumber((((_1190_v55) ? (_dafny.Seq.of(p1)) : (_1132_v6))).length)), p1), _module.__default.safeIndex(new BigNumber((_1192_v57).length), new BigNumber((_dafny.Seq.update(((_1190_v55) ? (_dafny.Seq.of(p1)) : (_1132_v6)), _module.__default.safeIndex(p1, new BigNumber((((_1190_v55) ? (_dafny.Seq.of(p1)) : (_1132_v6))).length)), p1)).length)), (_1132_v6)[_module.__default.safeIndex(p1, new BigNumber((_1132_v6).length))]), (((_this).f11) ? (_1132_v6) : (_1132_v6)));
      let _rhs199 = _dafny.Seq.IsProperPrefixOf(_1131_v5, _module.__default.fm33(p1, (_this).f11, (_this).f11, p1, globalState));
      let _rhs200 = (function (_pat_let19_0) {
        return function (_1196_dt__update__tmp_h0) {
          return function (_pat_let20_0) {
            return function (_1197_dt__update_hcf2_h0) {
              return _module.D0.create_DC1((_1196_dt__update__tmp_h0).dtor_cf1, _1197_dt__update_hcf2_h0);
            }(_pat_let20_0);
          }(_pat_let_tv29);
        }(_pat_let19_0);
      }(p2)).dtor_cf1;
      let _rhs201 = _dafny.Seq.update((((_1194_v59).contains(p1)) ? ((_1194_v59).get(p1)) : ((((_1195_v60).contains((_this).f11)) ? ((_1195_v60).get((_this).f11)) : (_1126_v0)))), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1126_v0).length)), new BigNumber(((((_1194_v59).contains(p1)) ? ((_1194_v59).get(p1)) : ((((_1195_v60).contains((_this).f11)) ? ((_1195_v60).get((_this).f11)) : (_1126_v0))))).length)), p0);
      let _rhs202 = !((_this).fm7(p1, globalState));
      _1132_v6 = _rhs198;
      _1190_v55 = _rhs199;
      _1190_v55 = _rhs200;
      _1126_v0 = _rhs201;
      _1190_v55 = _rhs202;
      let _1198_v61;
      let _init38 = ((_1199_v4) => function (_1200_i6) {
        return _1199_v4;
      })(_1130_v4);
      let _nw213 = Array((new BigNumber(3)).toNumber());
      for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw213.length); _i0_38++) {
        _nw213[_i0_38] = _init38(new BigNumber(_i0_38));
      }
      _1198_v61 = _nw213;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1198_v61).length))) {
        let _1201_i5 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1201_i5)) && ((_1201_i5).isLessThan(new BigNumber((_1198_v61).length))))) {
          (_1198_v61)[(_1201_i5)] = (_dafny.Set.fromElements(p1)).Intersect((_1130_v4).Intersect(function () {
            let _coll40 = new _dafny.Set();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(102), new BigNumber(717))) {
              let _1202_v62 = _compr_40;
              if (((new BigNumber(102)).isLessThanOrEqualTo(_1202_v62)) && ((_1202_v62).isLessThan(new BigNumber(717)))) {
                _coll40.add(_module.__default.safeDivisionInt(_1202_v62, new BigNumber((_1131_v5).length)));
              }
            }
            return _coll40;
          }()));
        }
      }
      let _1203_v63;
      let _nw214 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
      _1203_v63 = _nw214;
      let _index215 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1203_v63).length));
      (_1203_v63)[_index215] = _1133_v7;
      let _index216 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1203_v63).length));
      (_1203_v63)[_index216] = _1133_v7;
      return;
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
