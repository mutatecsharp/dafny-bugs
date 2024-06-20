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
    static fm0(p0, p1, p2, p3, globalState) {
      return (true) === (!(!(((true) ? (true) : (false)))));
    };
    static fm1(p0, p1, globalState) {
      let _source0 = _module.D10.create_DC23(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(11), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(473), new BigNumber((_dafny.Seq.UnicodeFromString("mwihgjxua")).length))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-810), new BigNumber((_dafny.Seq.of(true, false)).length))).length)))).length), !(true), true);
      if (_source0.is_DC23) {
        let _0___mcc_h0 = (_source0).cf33;
        let _1___mcc_h1 = (_source0).cf34;
        let _2___mcc_h2 = (_source0).cf35;
        let _3_cf35 = _2___mcc_h2;
        let _4_cf34 = _1___mcc_h1;
        let _5_cf33 = _0___mcc_h0;
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_5_cf33, new BigNumber(841))).length),_4_cf34)).length);
      } else {
        let _6___mcc_h3 = (_source0).cf32;
        let _7_cf32 = _6___mcc_h3;
        if ((_7_cf32).f16) {
          return new BigNumber((_dafny.Seq.of(new BigNumber(-846))).length);
        } else {
          return new BigNumber((_dafny.Seq.UnicodeFromString("dhf")).length);
        }
      }
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.UnicodeFromString("g");
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(933), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(200),true)).length))).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-918),!(false))).length)))).Union(_dafny.Set.fromElements(new BigNumber(-219), new BigNumber((_dafny.Seq.of(false, !(!(true)))).length), new BigNumber(623), new BigNumber(881)));
    };
    static fm10(p0, globalState) {
      return new BigNumber((function (_source1) {
        if (_source1.is_DC2) {
          let _8___mcc_h0 = (_source1).cf2;
          let _9_cf2 = _8___mcc_h0;
          return (_dafny.MultiSet.fromElements(new BigNumber(-14), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(541)), function (_10_i0) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })).length), new BigNumber(958), new BigNumber((function () {
            let _coll0 = new _dafny.Map();
            for (const _compr_0 of _dafny.IntegerRange(new BigNumber(892), new BigNumber(185))) {
              let _11_v0 = _compr_0;
              if (((new BigNumber(892)).isLessThanOrEqualTo(_11_v0)) && ((_11_v0).isLessThan(new BigNumber(185)))) {
                _coll0.push([(_11_v0).plus(new BigNumber(712)),new BigNumber((_dafny.Seq.of(_9_cf2, _9_cf2)).length)]);
              }
            }
            return _coll0;
          }()).length))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-142)), function (_12_i1) {
            return new BigNumber(-701);
          })));
        } else if (_source1.is_DC3) {
          let _13___mcc_h1 = (_source1).cf3;
          let _14___mcc_h2 = (_source1).cf4;
          let _15_cf4 = _14___mcc_h2;
          let _16_cf3 = _13___mcc_h1;
          return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(911)), function (_17_i2) {
            return new BigNumber(944);
          }), _dafny.Seq.of(new BigNumber(290), new BigNumber(674))));
        } else {
          let _18___mcc_h3 = (_source1).cf1;
          let _19_cf1 = _18___mcc_h3;
          return (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_19_cf1, _19_cf1, true)).length)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_19_cf1)).length)));
        }
      }(_module.D1.create_DC3(!(!(false)), new BigNumber(18)))).cardinality());
    };
    static fm11(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(756)), function (_20_i0) {
        return (_dafny.MultiSet.fromElements(!(!(false)), !(!(true)))).Union(_dafny.MultiSet.fromElements(true));
      });
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(false)).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)));
    };
    static fm15(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC8(true, new BigNumber(257), new BigNumber(9)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC8(true, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length), new BigNumber((_dafny.Seq.of(true)).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC8(true, new BigNumber((_dafny.Seq.of(true)).length), (_dafny.ZERO).minus(new BigNumber(-201)))));
    };
    static fm16(globalState) {
      return _module.D3.create_DC8(false, (new BigNumber(633)).multipliedBy((_module.D15.create_DC39(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(_module.D8.create_DC18(), _module.D8.create_DC18()), _dafny.Seq.of(_module.D8.create_DC18()))).cardinality()),true)).length), !(!(false)), new BigNumber(981), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(470)), function (_21_i0) {
  return new BigNumber((_dafny.Seq.of(new BigNumber(263))).length);
})).length))).dtor_cf59), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lriyttbrc"), _dafny.Seq.UnicodeFromString("jftuf"))).length));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return _module.__default.safeModuloInt((new BigNumber((_dafny.MultiSet.fromElements(false, false, false, true, true)).cardinality())).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-45)), function (_22_i0) {
        return (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length))));
      })).length)), (new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)))).length)).plus(new BigNumber(271)));
    };
    static fm18(p0, p1, p2, globalState) {
      return _dafny.Seq.of(false, !((new BigNumber(583)).isLessThan(new BigNumber(-922))), (false) || (!(false)));
    };
    static fm19(p0, globalState) {
      return ((_module.D15.create_DC39(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).length), true, new BigNumber((_dafny.Seq.UnicodeFromString("cynnxklv")).length), new BigNumber(-334))).dtor_cf59).plus(new BigNumber(-349));
    };
    static fm20(globalState) {
      if (!(!(!(false))) || (false)) {
        return function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of _dafny.IntegerRange(new BigNumber(650), new BigNumber(968))) {
              let _23_v1 = _compr_2;
              if (((new BigNumber(650)).isLessThanOrEqualTo(_23_v1)) && ((_23_v1).isLessThan(new BigNumber(968)))) {
                _coll2.push([_module.__default.safeDivisionInt(_23_v1, new BigNumber((_dafny.Seq.UnicodeFromString("ubretlxf")).length)),!(true)]);
              }
            }
            return _coll2;
          }()).Keys.Elements) {
            let _24_v0 = _compr_1;
            if ((function () {
              let _coll3 = new _dafny.Map();
              for (const _compr_3 of _dafny.IntegerRange(new BigNumber(650), new BigNumber(968))) {
                let _23_v1 = _compr_3;
                if (((new BigNumber(650)).isLessThanOrEqualTo(_23_v1)) && ((_23_v1).isLessThan(new BigNumber(968)))) {
                  _coll3.push([_module.__default.safeDivisionInt(_23_v1, new BigNumber((_dafny.Seq.UnicodeFromString("ubretlxf")).length)),!(true)]);
                }
              }
              return _coll3;
            }()).contains(_24_v0)) {
              _coll1.push([_module.__default.safeModuloInt(_24_v0, new BigNumber((_dafny.Seq.of(false)).length)),_dafny.Seq.of(new BigNumber(809))]);
            }
          }
          return _coll1;
        }();
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(!(false), !(false))).length),_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, true, true, false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("kjccwlgyi")).length), new BigNumber(12)));
      }
    };
    static fm21(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(!(((true) ? (!(false)) : (true))), false, !_dafny.areEqual((_dafny.Seq.Create(_module.__default.abs(new BigNumber(309)), function (_25_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(613)), function (_26_i1) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })));
    };
    static fm22(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(false),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(!(false))));
    };
    static fm23(p0, p1, globalState) {
      return _dafny.Set.fromElements(!(new BigNumber(-93)).isEqualTo(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, !(false)))).cardinality())), (_dafny.Set.fromElements(new BigNumber(235), new BigNumber(10))).IsSubsetOf(_dafny.Set.fromElements(new BigNumber(-751))));
    };
    static fm24(p0, p1, p2, globalState) {
      let _source2 = _module.D3.create_DC6(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ap"), (_dafny.Seq.UnicodeFromString("mthdwg")))).cardinality()), new BigNumber(-666))).length), true, false, new BigNumber((function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of _dafny.IntegerRange(new BigNumber(188), new BigNumber(665))) {
    let _27_v0 = _compr_4;
    if (((new BigNumber(188)).isLessThanOrEqualTo(_27_v0)) && ((_27_v0).isLessThan(new BigNumber(665)))) {
      _coll4.push([(_27_v0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-960)), function (_29_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(954)), function (_28_i1) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length)]);
    }
  }
  return _coll4;
}()).length), false);
      if (_source2.is_DC6) {
        let _30___mcc_h0 = (_source2).cf7;
        let _31___mcc_h1 = (_source2).cf8;
        let _32___mcc_h2 = (_source2).cf9;
        let _33___mcc_h3 = (_source2).cf10;
        let _34___mcc_h4 = (_source2).cf11;
        let _35_cf11 = _34___mcc_h4;
        let _36_cf10 = _33___mcc_h3;
        let _37_cf9 = _32___mcc_h2;
        let _38_cf8 = _31___mcc_h1;
        let _39_cf7 = _30___mcc_h0;
        return _dafny.MultiSet.fromElements(new BigNumber(608));
      } else if (_source2.is_DC7) {
        let _40___mcc_h5 = (_source2).cf12;
        let _41_cf12 = _40___mcc_h5;
        return (_dafny.MultiSet.fromElements(_41_cf12, _41_cf12)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(321), _41_cf12)));
      } else if (_source2.is_DC8) {
        let _42___mcc_h6 = (_source2).cf13;
        let _43___mcc_h7 = (_source2).cf14;
        let _44___mcc_h8 = (_source2).cf15;
        let _45_cf15 = _44___mcc_h8;
        let _46_cf14 = _43___mcc_h7;
        let _47_cf13 = _42___mcc_h6;
        return _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_module.D14.create_DC36(_45_cf15, _47_cf13, _47_cf13, new BigNumber((_dafny.Seq.UnicodeFromString("gqwuhixo")).length), _47_cf13)).dtor_cf51)).length)));
      } else {
        let _48___mcc_h9 = (_source2).cf6;
        let _49_cf6 = _48___mcc_h9;
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("wiaxnjh")).length));
      }
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(841)), function (_50_i0) {
        return (new BigNumber((_dafny.Seq.UnicodeFromString("gmhrvhsnu")).length)).plus(new BigNumber(207));
      });
    };
    static fm26(p0, globalState) {
      let _source3 = _dafny.Seq.UnicodeFromString("r");
      let _51___mcc_h0 = _source3;
      let _52_cf20 = _51___mcc_h0;
      return new _dafny.CodePoint('w'.codePointAt(0));
    };
    static fm27(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-319),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()),_dafny.Seq.Create(_module.__default.abs(new BigNumber(36)), function (_53_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }))).length))).cardinality())),false));
    };
    static fm28(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("butfuvwqw")).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(170),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false, !(false), false, true)).cardinality()),function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of ((_module.D17.create_DC44(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))).cardinality()))).length), new BigNumber(728)))).dtor_cf70).Elements) {
          let _54_v0 = _compr_5;
          if (((_module.D17.create_DC44(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))).cardinality()))).length), new BigNumber(728)))).dtor_cf70).contains(_54_v0)) {
            _coll5.push([(_54_v0).plus(new BigNumber(866)),true]);
          }
        }
        return _coll5;
      }()))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-349))),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_module.D9.create_DC20(true, new _dafny.CodePoint('e'.codePointAt(0)))).dtor_cf29, true)).length),new BigNumber((_dafny.Seq.UnicodeFromString("nwvavjx")).length))).length),false)));
    };
    static fm29(p0, p1, p2, p3, globalState) {
      let _source4 = _module.D8.create_DC18();
      if (_source4.is_DC18) {
        return _module.D1.create_DC1(true);
      } else {
        let _55___mcc_h0 = (_source4).cf27;
        let _56_cf27 = _55___mcc_h0;
        return _module.D1.create_DC1(true);
      }
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return _module.D9.create_DC20(((!(false)) ? (true) : (!(false))), new _dafny.CodePoint('j'.codePointAt(0)));
    };
    static m0(globalState) {
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.Map.Empty;
      let _57_v0;
      _57_v0 = false;
      let _58_v1;
      let _nw0 = Array((new BigNumber(26)).toNumber()).fill(false);
      _58_v1 = _nw0;
      let _59_v2;
      let _nw1 = Array((new BigNumber(6)).toNumber());
      _nw1[(_dafny.ZERO).toNumber()] = _58_v1;
      _nw1[(_dafny.ONE).toNumber()] = _58_v1;
      _nw1[(new BigNumber(2)).toNumber()] = _58_v1;
      _nw1[(new BigNumber(3)).toNumber()] = _58_v1;
      _nw1[(new BigNumber(4)).toNumber()] = _58_v1;
      _nw1[(new BigNumber(5)).toNumber()] = _58_v1;
      _59_v2 = _nw1;
      let _60_v3;
      _60_v3 = _module.D3.create_DC5(_59_v2);
      let _61_v4;
      let _nw2 = new _module.C1();
      _nw2.__ctor(((_57_v0) ? (_57_v0) : (_57_v0)), _60_v3);
      _61_v4 = _nw2;
      let _62_v5;
      _62_v5 = new BigNumber(188);
      if (((_57_v0) ? (_57_v0) : (_module.__default.fm0(_62_v5, _62_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(635)), ((_63_v4) => function (_64_i0) {
        return _dafny.MultiSet.fromElements((_63_v4).f16);
      })(_61_v4)), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rkrihtkwb"),!((_61_v4).f16)), globalState)))) {
        let _65_v6;
        _65_v6 = _dafny.MultiSet.fromElements(_57_v0, _57_v0);
        let _66_v7;
        _66_v7 = _dafny.Seq.of(_65_v6);
        let _67_v8;
        _67_v8 = _dafny.Seq.UnicodeFromString("gmqmsbnsb");
        let _68_v9;
        _68_v9 = new _dafny.CodePoint('d'.codePointAt(0));
        let _69_v10;
        _69_v10 = _dafny.Map.Empty.slice().updateUnsafe(_67_v8,(_61_v4).fm4(_62_v5, true, _68_v9, globalState));
        if (((((_61_v4).f16) ? (false) : (_module.__default.fm0(_62_v5, _62_v5, _66_v7, _69_v10, globalState)))) === ((_61_v4).f16)) {
          r1 = !(_57_v0);
          _58_v1 = _58_v1;
          let _index0 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_58_v1).length));
          (_58_v1)[_index0] = _57_v0;
          let _70_v11;
          _70_v11 = _dafny.Seq.of((_61_v4).f16);
          let _index1 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_58_v1).length));
          let _rhs0 = _68_v9;
          let _rhs1 = (_70_v11)[_module.__default.safeIndex(_62_v5, new BigNumber((_70_v11).length))];
          let _rhs2 = (_61_v4).f16;
          let _lhs0 = _58_v1;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_58_v1).length));
          _68_v9 = _rhs0;
          _lhs0[_lhs1] = _rhs1;
          _57_v0 = _rhs2;
          let _index2 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_58_v1).length));
          (_58_v1)[_index2] = (_61_v4).f16;
          let _71_v12;
          let _nw3 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _71_v12 = _nw3;
          _71_v12 = _71_v12;
        } else {
          r2 = !(!((_61_v4).f16)) || ((_61_v4).f16);
          let _72_v13;
          _72_v13 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_62_v5),(_61_v4).f16);
          let _73_v14;
          _73_v14 = _dafny.Seq.of(new BigNumber((_72_v13).length));
          let _74_v15;
          _74_v15 = _dafny.Set.fromElements(_62_v5, new BigNumber((_73_v14).length), _62_v5);
          let _75_v16;
          _75_v16 = _dafny.Map.Empty.slice().updateUnsafe(_58_v1,_67_v8);
          let _76_v17;
          _76_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_74_v15).length),_75_v16);
          _76_v17 = (_76_v17).update(new BigNumber(184), _75_v16);
          _69_v10 = (_69_v10).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), ((_77_v9) => function (_78_i1) {
            return _77_v9;
          })(_68_v9)), (_61_v4).f16);
          let _79_v18;
          _79_v18 = _dafny.MultiSet.fromElements(_68_v9);
          let _80_v19;
          _80_v19 = _module.D15.create_DC39((_61_v4).f16, new BigNumber((_79_v18).cardinality()), false, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), ((_81_v9) => function (_82_i2) {
  return _81_v9;
})(_68_v9))).length), _62_v5);
          let _83_v20;
          _83_v20 = _dafny.Map.Empty.slice().updateUnsafe(_62_v5,_80_v19);
          _83_v20 = (_83_v20).update((_62_v5).multipliedBy(_62_v5), _80_v19);
          (globalState).f2 = _62_v5;
        }
        let _84_v21;
        let _init0 = ((_85_v5) => function (_86_i3) {
          return (_86_i3).multipliedBy(_85_v5);
        })(_62_v5);
        let _nw4 = Array((new BigNumber(3)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw4.length); _i0_0++) {
          _nw4[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _84_v21 = _nw4;
        let _index3 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length));
        (_84_v21)[_index3] = _module.__default.safeModuloInt(new BigNumber(127), _62_v5);
        let _index4 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length));
        (_84_v21)[_index4] = _module.__default.safeModuloInt(_62_v5, _62_v5);
        r1 = (_61_v4).f16;
        if ((_61_v4).f16) {
          (globalState).f4 = _84_v21;
          let _index5 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_59_v2).length));
          (_59_v2)[_index5] = _58_v1;
          let _index6 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_59_v2).length));
          (_59_v2)[_index6] = _58_v1;
          let _index7 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length));
          (_84_v21)[_index7] = (_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))];
          let _87_v22;
          _87_v22 = _dafny.Map.Empty.slice().updateUnsafe(_57_v0,new BigNumber((_65_v6).cardinality()));
          let _88_v23;
          _88_v23 = _dafny.Set.fromElements(new BigNumber((_67_v8).length));
          let _89_v24;
          _89_v24 = _dafny.Seq.of(_88_v23, _88_v23);
          let _90_v25;
          let _nw5 = new _module.C0();
          _nw5.__ctor(_module.__default.fm0(_module.__default.fm19(new BigNumber((_87_v22).length), globalState), (_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))], _66_v7, _69_v10, globalState), (new BigNumber((_89_v24).length)).minus(_62_v5));
          _90_v25 = _nw5;
          let _91_v26;
          let _nw6 = new _module.C2();
          _nw6.__ctor(_57_v0, ((_90_v25).f12) && (true));
          _91_v26 = _nw6;
          let _arr0 = (_59_v2)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_59_v2).length))];
          let _index8 = _module.__default.safeIndex(new BigNumber(136), new BigNumber(((_59_v2)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_59_v2).length))]).length));
          _arr0[_index8] = (_61_v4).f16;
          let _index9 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length));
          let _arr1 = (_59_v2)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_59_v2).length))];
          let _index10 = _module.__default.safeIndex(new BigNumber(136), new BigNumber(((_59_v2)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_59_v2).length))]).length));
          let _rhs3 = (_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))];
          let _rhs4 = _90_v25;
          let _rhs5 = _91_v26;
          let _rhs6 = !((_61_v4).f16);
          let _rhs7 = (_91_v26).fm4(new BigNumber(994), !(false), new _dafny.CodePoint('h'.codePointAt(0)), globalState);
          let _lhs2 = _84_v21;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length));
          let _lhs4 = (_59_v2)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_59_v2).length))];
          let _lhs5 = _module.__default.safeIndex(new BigNumber(136), new BigNumber(((_59_v2)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_59_v2).length))]).length));
          _lhs2[_lhs3] = _rhs3;
          _90_v25 = _rhs4;
          _91_v26 = _rhs5;
          r1 = _rhs6;
          _lhs4[_lhs5] = _rhs7;
          let _92_v27;
          let _nw7 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
          _92_v27 = _nw7;
          let _index11 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_92_v27).length));
          (_92_v27)[_index11] = _69_v10;
          let _93_v28;
          _93_v28 = _dafny.Seq.of((_90_v25).f13, _62_v5);
          let _index12 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length));
          let _index13 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_92_v27).length));
          let _index14 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length));
          let _rhs8 = (((_90_v25).f12) ? ((_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))]) : ((_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))]));
          let _rhs9 = _69_v10;
          let _rhs10 = ((((false) ? (_module.__default.fm0((_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))], _62_v5, _66_v7, _69_v10, globalState)) : ((_61_v4).f16))) ? ((_61_v4).f16) : (!((_62_v5).isLessThan((_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))]))));
          let _rhs11 = ((_90_v25).f13).plus((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(204)), ((_94_v9) => function (_95_i4) {
            return _94_v9;
          })(_68_v9))).length)).minus((_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))]));
          let _rhs12 = ((_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))]).plus((_93_v28)[_module.__default.safeIndex((_dafny.ZERO).minus((_90_v25).f13), new BigNumber((_93_v28).length))]);
          let _lhs6 = _84_v21;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length));
          let _lhs8 = _92_v27;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_92_v27).length));
          let _lhs10 = _84_v21;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length));
          _lhs6[_lhs7] = _rhs8;
          _lhs8[_lhs9] = _rhs9;
          _57_v0 = _rhs10;
          _lhs10[_lhs11] = _rhs11;
          _62_v5 = _rhs12;
        } else {
          _67_v8 = _67_v8;
          let _96_v29;
          _96_v29 = _dafny.Set.fromElements(_58_v1);
          let _97_v30;
          _97_v30 = _dafny.Seq.of(_module.D7.create_DC14(_96_v29));
          _97_v30 = _97_v30;
          let _98_v31;
          let _nw8 = new _module.C2();
          _nw8.__ctor((((_61_v4).fm4(new BigNumber(567), (_61_v4).f16, _68_v9, globalState)) ? (!(false)) : (_57_v0)), true);
          _98_v31 = _nw8;
          _98_v31 = _98_v31;
          let _99_v32;
          _99_v32 = _dafny.Seq.of(_67_v8);
          let _100_v33;
          let _nw9 = new _module.C1();
          _nw9.__ctor(_dafny.Seq.IsProperPrefixOf(_67_v8, (_99_v32)[_module.__default.safeIndex(new BigNumber((_67_v8).length), new BigNumber((_99_v32).length))]), _61_v4.f17);
          _100_v33 = _nw9;
          _68_v9 = _68_v9;
        }
        let _101_v34;
        _101_v34 = _dafny.Map.Empty.slice().updateUnsafe(_57_v0,(_62_v5).isLessThanOrEqualTo((_84_v21)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_84_v21).length))]));
        _101_v34 = (_101_v34).update(!((_61_v4).f16) || (_57_v0), !(_57_v0) || ((_61_v4).f16));
      } else {
        let _102_v35;
        let _nw10 = new _module.C2();
        _nw10.__ctor((_61_v4).f16, (_61_v4).f16);
        _102_v35 = _nw10;
        _102_v35 = _102_v35;
        let _103_v36;
        _103_v36 = _dafny.Map.Empty.slice().updateUnsafe((_61_v4).f16,_102_v35.f14);
        let _104_v37;
        _104_v37 = _module.D9.create_DC20(!((((_103_v36).contains((_102_v35).f15)) ? ((_103_v36).get((_102_v35).f15)) : ((_102_v35).f15))), new _dafny.CodePoint('t'.codePointAt(0)));
        let _105_v38;
        _105_v38 = _dafny.Seq.UnicodeFromString("kplyjx");
        let _106_v39;
        _106_v39 = _dafny.Seq.of(new BigNumber(647), _62_v5);
        let _107_v40;
        _107_v40 = _module.D16.create_DC42((_106_v39)[_module.__default.safeIndex(new BigNumber(36), new BigNumber((_106_v39).length))], (_61_v4).f16);
        let _108_v41;
        _108_v41 = new _dafny.CodePoint('a'.codePointAt(0));
        let _109_v42;
        _109_v42 = _dafny.Map.Empty.slice().updateUnsafe(_62_v5,_dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), ((_110_v41) => function (_111_i5) {
          return _110_v41;
        })(_108_v41)));
        let _rhs13 = !(false);
        let _rhs14 = _module.__default.fm30((_107_v40).dtor_cf65, new BigNumber(210), _108_v41, _62_v5, globalState);
        let _rhs15 = (((_109_v42).contains((_dafny.ZERO).minus(new BigNumber((_105_v38).length)))) ? ((_109_v42).get((_dafny.ZERO).minus(new BigNumber((_105_v38).length)))) : (_module.__default.fm2(_102_v35.f14, _62_v5, new BigNumber(875), true, globalState)));
        let _rhs16 = ((new BigNumber(970)).minus(_62_v5)).isLessThan((_dafny.ZERO).minus(_62_v5));
        let _lhs12 = _102_v35;
        r1 = _rhs13;
        _104_v37 = _rhs14;
        _105_v38 = _rhs15;
        _lhs12.f14 = _rhs16;
        r1 = (_61_v4).f16;
        if (_57_v0) {
          let _112_v43;
          let _nw11 = new _module.C1();
          _nw11.__ctor((_61_v4).f16, _61_v4.f17);
          _112_v43 = _nw11;
          let _113_v44;
          _113_v44 = _dafny.Map.Empty.slice().updateUnsafe(_105_v38,_module.__default.safeDivisionInt(_62_v5, _62_v5));
          let _114_v45;
          _114_v45 = _dafny.Seq.of(_112_v43, _112_v43, _112_v43);
          let _115_v47;
          _115_v47 = _dafny.Set.fromElements(_105_v38);
          let _rhs17 = (_114_v45)[_module.__default.safeIndex(_62_v5, new BigNumber((_114_v45).length))];
          let _rhs18 = _62_v5;
          let _rhs19 = _60_v3;
          let _rhs20 = _62_v5;
          let _rhs21 = (function () {
            let _coll6 = new _dafny.Map();
            for (const _compr_6 of (_115_v47).Elements) {
              let _116_v46 = _compr_6;
              if ((_115_v47).contains(_116_v46)) {
                _coll6.push([_116_v46,(_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(_62_v5))).length))]);
              }
            }
            return _coll6;
          }()).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-734)), ((_117_v41) => function (_118_i6) {
            return _117_v41;
          })(_108_v41)), _62_v5);
          let _lhs13 = globalState;
          let _lhs14 = globalState;
          _112_v43 = _rhs17;
          _lhs13.f2 = _rhs18;
          _60_v3 = _rhs19;
          _lhs14.f2 = _rhs20;
          _113_v44 = _rhs21;
          let _119_v48;
          _119_v48 = _dafny.MultiSet.fromElements(_112_v43);
          let _120_v49;
          _120_v49 = _dafny.Seq.of((_61_v4).f16, (_102_v35).fm12(true, (((_119_v48).contains(_112_v43)) ? ((_119_v48).get(_112_v43)) : (_62_v5)), _dafny.MultiSet.fromElements(_62_v5), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), ((_121_v41) => function (_122_i7) {
            return _121_v41;
          })(_108_v41))).length), globalState));
          let _123_v50;
          _123_v50 = _dafny.Map.Empty.slice().updateUnsafe(_62_v5,_57_v0);
          let _124_v51;
          let _nw12 = new _module.C3();
          _nw12.__ctor(_62_v5, _123_v50);
          _124_v51 = _nw12;
          let _125_v52;
          _125_v52 = _dafny.Map.Empty.slice().updateUnsafe(_57_v0,new BigNumber((_120_v49).length));
          let _126_v53;
          _126_v53 = _dafny.Set.fromElements(new BigNumber(198));
          let _127_v54;
          _127_v54 = _dafny.MultiSet.fromElements(_102_v35.f14, (_102_v35).f15);
          let _128_v55;
          _128_v55 = _module.D15.create_DC39(_102_v35.f14, new BigNumber((_105_v38).length), (_61_v4).f16, new BigNumber((_127_v54).cardinality()), _124_v51.f10);
          let _129_v56;
          _129_v56 = _124_v51.f10;
          let _130_v57;
          _130_v57 = _dafny.Map.Empty.slice().updateUnsafe(_129_v56,_102_v35.f14);
          let _131_v58;
          let _nw13 = Array((new BigNumber(9)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _62_v5;
          _nw13[(_dafny.ONE).toNumber()] = new BigNumber((_103_v36).length);
          _nw13[(new BigNumber(2)).toNumber()] = (_62_v5).minus((((_125_v52).contains(_102_v35.f14)) ? ((_125_v52).get(_102_v35.f14)) : (new BigNumber((_105_v38).length))));
          _nw13[(new BigNumber(3)).toNumber()] = new BigNumber((_module.__default.fm18(new BigNumber((_126_v53).length), (_128_v55).dtor_cf58, _124_v51.f10, globalState)).length);
          _nw13[(new BigNumber(4)).toNumber()] = _62_v5;
          _nw13[(new BigNumber(5)).toNumber()] = (_module.__default.fm10((_102_v35).f15, globalState)).minus(new BigNumber((_module.__default.fm25(_62_v5, _module.__default.fm17(_102_v35.f14, _124_v51.f10, (_102_v35).f15, new BigNumber((_130_v57).length), globalState), _124_v51.f10, _62_v5, globalState)).length));
          _nw13[(new BigNumber(6)).toNumber()] = _124_v51.f10;
          _nw13[(new BigNumber(7)).toNumber()] = _62_v5;
          _nw13[(new BigNumber(8)).toNumber()] = _62_v5;
          _131_v58 = _nw13;
          let _index15 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_131_v58).length));
          (_131_v58)[_index15] = new BigNumber(((_124_v51).f11).length);
          let _index16 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_131_v58).length));
          let _rhs22 = _dafny.Seq.update(_dafny.Seq.update(_120_v49, _module.__default.safeIndex(_124_v51.f10, new BigNumber((_120_v49).length)), _102_v35.f14), _module.__default.safeIndex(new BigNumber((_127_v54).cardinality()), new BigNumber((_dafny.Seq.update(_120_v49, _module.__default.safeIndex(_124_v51.f10, new BigNumber((_120_v49).length)), _102_v35.f14)).length)), _dafny.Seq.IsProperPrefixOf(_105_v38, _105_v38));
          let _rhs23 = _124_v51.f10;
          let _rhs24 = _124_v51;
          let _rhs25 = _62_v5;
          let _rhs26 = _124_v51.f10;
          let _lhs15 = globalState;
          let _lhs16 = _131_v58;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_131_v58).length));
          _120_v49 = _rhs22;
          _lhs15.f2 = _rhs23;
          _124_v51 = _rhs24;
          _lhs16[_lhs17] = _rhs25;
          _62_v5 = _rhs26;
          _106_v39 = _106_v39;
          r0 = (_dafny.ZERO).minus((_106_v39)[_module.__default.safeIndex((_dafny.ZERO).minus(_62_v5), new BigNumber((_106_v39).length))]);
          let _132_v59;
          _132_v59 = _dafny.MultiSet.fromElements(_131_v58, _131_v58);
          let _index17 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_58_v1).length));
          (_58_v1)[_index17] = ((_131_v58)[_module.__default.safeIndex(new BigNumber(82), new BigNumber((_131_v58).length))]).isEqualTo(new BigNumber((_132_v59).cardinality()));
          let _index18 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_58_v1).length));
          let _rhs27 = _58_v1;
          let _rhs28 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_module.D3.create_DC8((_61_v4).f16, _124_v51.f10, (_dafny.ZERO).minus(_124_v51.f10))).dtor_cf13,(_61_v4).f16)).update(false, !((_module.D3.create_DC6((_131_v58)[_module.__default.safeIndex(new BigNumber(82), new BigNumber((_131_v58).length))], _57_v0, (_61_v4).f16, _124_v51.f10, (_102_v35).f15)).dtor_cf8))).length);
          let _rhs29 = _102_v35.f14;
          let _rhs30 = _module.__default.safeDivisionInt(_124_v51.f10, _124_v51.f10);
          let _lhs18 = _58_v1;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_58_v1).length));
          _58_v1 = _rhs27;
          _62_v5 = _rhs28;
          _lhs18[_lhs19] = _rhs29;
          r0 = _rhs30;
        } else {
          (_102_v35).f14 = ((((_61_v4).f16) ? (new BigNumber(376)) : (new BigNumber((_105_v38).length)))).isLessThanOrEqualTo(_62_v5);
          let _133_v60;
          _133_v60 = _dafny.Set.fromElements(_106_v39, _106_v39);
          r0 = (_dafny.ZERO).minus(new BigNumber((_133_v60).length));
          let _134_v61;
          _134_v61 = _dafny.Map.Empty.slice().updateUnsafe((_62_v5).minus(_62_v5),_62_v5);
          let _rhs31 = _62_v5;
          let _rhs32 = _134_v61;
          let _lhs20 = globalState;
          _lhs20.f2 = _rhs31;
          _134_v61 = _rhs32;
          let _135_v62;
          let _nw14 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _135_v62 = _nw14;
          let _index19 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_135_v62).length));
          (_135_v62)[_index19] = new BigNumber((_106_v39).length);
          let _136_v63;
          _136_v63 = _dafny.Map.Empty.slice().updateUnsafe((_61_v4).f16,(_dafny.ZERO).minus(_62_v5));
          let _137_v64;
          _137_v64 = _module.D9.create_DC19(_136_v63);
          let _138_v65;
          _138_v65 = _module.D9.create_DC21(_137_v64);
          let _139_v66;
          let _nw15 = Array((new BigNumber(2)).toNumber());
          _nw15[(_dafny.ZERO).toNumber()] = _module.D9.create_DC21(_137_v64);
          _nw15[(_dafny.ONE).toNumber()] = _138_v65;
          _139_v66 = _nw15;
          let _index20 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_139_v66).length));
          (_139_v66)[_index20] = _138_v65;
          let _140_v67;
          _140_v67 = _dafny.Map.Empty.slice().updateUnsafe(_62_v5,(_61_v4).f16);
          let _index21 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_135_v62).length));
          let _index22 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_139_v66).length));
          let _rhs33 = _61_v4;
          let _rhs34 = (_102_v35.f14) === (!(_57_v0) || ((((_140_v67).contains(_62_v5)) ? ((_140_v67).get(_62_v5)) : (_102_v35.f14))));
          let _rhs35 = _module.__default.fm19(_62_v5, globalState);
          let _rhs36 = _138_v65;
          let _lhs21 = _135_v62;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_135_v62).length));
          let _lhs23 = _139_v66;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_139_v66).length));
          _61_v4 = _rhs33;
          _57_v0 = _rhs34;
          _lhs21[_lhs22] = _rhs35;
          _lhs23[_lhs24] = _rhs36;
          let _141_v68;
          _141_v68 = _dafny.Seq.of((_102_v35).f15, true);
          let _index23 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_135_v62).length));
          (_135_v62)[_index23] = new BigNumber((_141_v68).length);
        }
        _58_v1 = _58_v1;
      }
      let _142_v69;
      let _nw16 = new _module.C0();
      _nw16.__ctor(_57_v0, _62_v5);
      _142_v69 = _nw16;
      let _143_v70;
      _143_v70 = _dafny.Map.Empty.slice().updateUnsafe((_61_v4).f16,_142_v69);
      _143_v70 = (_143_v70).update((_61_v4).f16, _142_v69);
      r1 = false;
      r2 = _57_v0;
      let _index24 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_58_v1).length));
      (_58_v1)[_index24] = _57_v0;
      let _144_v71;
      let _init1 = ((_145_v0) => function (_146_i8) {
        return (_146_i8).plus(new BigNumber((_dafny.Seq.of(_145_v0)).length));
      })(_57_v0);
      let _nw17 = Array((new BigNumber(4)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw17.length); _i0_1++) {
        _nw17[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _144_v71 = _nw17;
      let _147_v72;
      _147_v72 = _dafny.Seq.of((_142_v69).f13);
      let _index25 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_58_v1).length));
      let _rhs37 = _144_v71;
      let _rhs38 = (_module.__default.fm10(_57_v0, globalState)).plus(((_142_v69).f13).minus((_147_v72)[_module.__default.safeIndex((_142_v69).f13, new BigNumber((_147_v72).length))]));
      let _rhs39 = _57_v0;
      let _lhs25 = globalState;
      let _lhs26 = _58_v1;
      let _lhs27 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_58_v1).length));
      _lhs25.f4 = _rhs37;
      _62_v5 = _rhs38;
      _lhs26[_lhs27] = _rhs39;
      r0 = new BigNumber(809);
      let _148_v73;
      _148_v73 = new _dafny.CodePoint('c'.codePointAt(0));
      let _149_v74;
      _149_v74 = _dafny.Seq.of(_148_v73);
      r1 = (new BigNumber((_149_v74).length)).isLessThan(_62_v5);
      r2 = !((((_142_v69).f12) ? (((_58_v1)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_58_v1).length))]) === (_57_v0)) : (_57_v0)));
      let _150_v75;
      _150_v75 = _dafny.Map.Empty.slice().updateUnsafe((_58_v1)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_58_v1).length))],(_142_v69).f13);
      r3 = (((_142_v69).f12) ? ((((_61_v4).f16) ? (_150_v75) : (_150_v75))) : (_150_v75));
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _151_v0;
      let _nw18 = Array((new BigNumber(4)).toNumber()).fill([]);
      _151_v0 = _nw18;
      let _152_v1;
      _152_v1 = new BigNumber(581);
      let _153_v2;
      _153_v2 = _dafny.MultiSet.fromElements(_152_v1, _152_v1);
      let _154_v3;
      _154_v3 = _dafny.Seq.of(new BigNumber((_153_v2).cardinality()));
      let _155_v4;
      _155_v4 = false;
      let _156_v5;
      _156_v5 = _dafny.Map.Empty.slice().updateUnsafe(_152_v1,_155_v4);
      let _157_v6;
      _157_v6 = _dafny.Seq.of(_155_v4);
      let _158_v7;
      _158_v7 = _dafny.Set.fromElements(new BigNumber((_157_v6).length));
      let _159_v8;
      _159_v8 = _152_v1;
      let _160_v9;
      _160_v9 = _dafny.Seq.UnicodeFromString("joueqk");
      let _161_v10;
      let _nw19 = Array((new BigNumber(29)).toNumber());
      _nw19[(_dafny.ZERO).toNumber()] = new BigNumber(-822);
      _nw19[(_dafny.ONE).toNumber()] = _152_v1;
      _nw19[(new BigNumber(2)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(3)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(4)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(5)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_152_v1);
      _nw19[(new BigNumber(7)).toNumber()] = new BigNumber(288);
      _nw19[(new BigNumber(8)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(9)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(10)).toNumber()] = new BigNumber((_154_v3).length);
      _nw19[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_156_v5, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(55),_155_v4), _dafny.Map.Empty.slice().updateUnsafe(_152_v1,_155_v4)), _module.__default.safeIndex(new BigNumber((_158_v7).length), new BigNumber((_dafny.Seq.of(_156_v5, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(55),_155_v4), _dafny.Map.Empty.slice().updateUnsafe(_152_v1,_155_v4))).length)), _156_v5)).length);
      _nw19[(new BigNumber(12)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(13)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(14)).toNumber()] = new BigNumber((_156_v5).length);
      _nw19[(new BigNumber(15)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(16)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(17)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(18)).toNumber()] = (_159_v8);
      _nw19[(new BigNumber(19)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(20)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(21)).toNumber()] = new BigNumber(576);
      _nw19[(new BigNumber(22)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(23)).toNumber()] = new BigNumber((_160_v9).length);
      _nw19[(new BigNumber(24)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(25)).toNumber()] = new BigNumber(-713);
      _nw19[(new BigNumber(26)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(27)).toNumber()] = _152_v1;
      _nw19[(new BigNumber(28)).toNumber()] = _152_v1;
      _161_v10 = _nw19;
      let _162_globalState;
      let _nw20 = new _module.GlobalState();
      _nw20.__ctor(true, false, new BigNumber(572), _151_v0, _161_v10, new BigNumber(807), new BigNumber(286), _161_v10, new BigNumber(-297), new BigNumber(926));
      _162_globalState = _nw20;
      let _163_v11;
      let _nw21 = Array((new BigNumber(2)).toNumber()).fill(_dafny.MultiSet.Empty);
      _163_v11 = _nw21;
      let _164_v12;
      _164_v12 = _dafny.Map.Empty.slice().updateUnsafe(_152_v1,_163_v11);
      _164_v12 = (_164_v12).update(_152_v1, _163_v11);
      let _165_v13;
      _165_v13 = _dafny.MultiSet.fromElements(_155_v4, _155_v4, _155_v4, _155_v4);
      let _166_v14;
      _166_v14 = _dafny.Seq.of(_165_v13);
      let _167_v15;
      _167_v15 = _dafny.Map.Empty.slice().updateUnsafe(_160_v9,_155_v4);
      if (_module.__default.fm0(_152_v1, _152_v1, _166_v14, _167_v15, _162_globalState)) {
        let _168_v16;
        let _169_v17;
        let _170_v18;
        let _171_v19;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = _module.__default.m0(_162_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _168_v16 = _out0;
        _169_v17 = _out1;
        _170_v18 = _out2;
        _171_v19 = _out3;
        let _172_v20;
        let _init2 = ((_173_v17) => function (_174_i0) {
          return (_module.D1.create_DC2(_173_v17)).dtor_cf2;
        })(_169_v17);
        let _nw22 = Array((new BigNumber(11)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw22.length); _i0_2++) {
          _nw22[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _172_v20 = _nw22;
        let _index26 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_172_v20).length));
        (_172_v20)[_index26] = (_152_v1).isLessThanOrEqualTo(new BigNumber((_157_v6).length));
        let _index27 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_172_v20).length));
        let _rhs40 = false;
        let _rhs41 = _170_v18;
        let _rhs42 = new BigNumber(261);
        let _rhs43 = (_152_v1).isLessThanOrEqualTo(_168_v16);
        let _lhs28 = _162_globalState;
        let _lhs29 = _172_v20;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_172_v20).length));
        _169_v17 = _rhs40;
        _169_v17 = _rhs41;
        _lhs28.f2 = _rhs42;
        _lhs29[_lhs30] = _rhs43;
        let _175_v21;
        _175_v21 = _module.D1.create_DC2(_169_v17);
        _175_v21 = _module.D1.create_DC2(true);
        _170_v18 = _dafny.areEqual(_154_v3, _dafny.Seq.of(_168_v16));
        let _176_v22;
        let _177_v23;
        let _178_v24;
        let _179_v25;
        let _out4;
        let _out5;
        let _out6;
        let _out7;
        let _outcollector1 = _module.__default.m0(_162_globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _out6 = _outcollector1[2];
        _out7 = _outcollector1[3];
        _176_v22 = _out4;
        _177_v23 = _out5;
        _178_v24 = _out6;
        _179_v25 = _out7;
      } else {
        let _180_v26;
        let _181_v27;
        let _182_v28;
        let _183_v29;
        let _out8;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector2 = _module.__default.m0(_162_globalState);
        _out8 = _outcollector2[0];
        _out9 = _outcollector2[1];
        _out10 = _outcollector2[2];
        _out11 = _outcollector2[3];
        _180_v26 = _out8;
        _181_v27 = _out9;
        _182_v28 = _out10;
        _183_v29 = _out11;
        _155_v4 = _module.__default.fm0((_152_v1).minus(_152_v1), _180_v26, _dafny.Seq.Concat(_166_v14, _166_v14), _167_v15, _162_globalState);
        _161_v10 = ((_181_v27) ? (_161_v10) : (_161_v10));
        if (_module.__default.fm0(_152_v1, _152_v1, _166_v14, (_167_v15).Merge(_167_v15), _162_globalState)) {
          let _184_v30;
          let _init3 = ((_185_v1, _186_v28) => function (_187_i1) {
            return _dafny.Map.Empty.slice().updateUnsafe(_185_v1,!(_186_v28));
          })(_152_v1, _182_v28);
          let _nw23 = Array((new BigNumber(23)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw23.length); _i0_3++) {
            _nw23[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _184_v30 = _nw23;
          _184_v30 = _184_v30;
          let _index28 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_161_v10).length));
          (_161_v10)[_index28] = _180_v26;
          let _188_v31;
          let _nw24 = Array((new BigNumber(21)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = _159_v8;
          _nw24[(_dafny.ONE).toNumber()] = _159_v8;
          _nw24[(new BigNumber(2)).toNumber()] = _180_v26;
          _nw24[(new BigNumber(3)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(4)).toNumber()] = new BigNumber(386);
          _nw24[(new BigNumber(5)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(6)).toNumber()] = _module.__default.fm1(_181_v27, _155_v4, _162_globalState);
          _nw24[(new BigNumber(7)).toNumber()] = new BigNumber(911);
          _nw24[(new BigNumber(8)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(9)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(10)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(11)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(12)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(13)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(14)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(15)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(16)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(17)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(18)).toNumber()] = _159_v8;
          _nw24[(new BigNumber(19)).toNumber()] = _152_v1;
          _nw24[(new BigNumber(20)).toNumber()] = new BigNumber((_183_v29).length);
          _188_v31 = _nw24;
          let _index29 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_188_v31).length));
          (_188_v31)[_index29] = _159_v8;
          let _189_v32;
          let _nw25 = Array((new BigNumber(4)).toNumber()).fill(false);
          _189_v32 = _nw25;
          let _index30 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_189_v32).length));
          (_189_v32)[_index30] = _181_v27;
          let _index31 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_161_v10).length));
          let _index32 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_188_v31).length));
          let _index33 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_189_v32).length));
          let _rhs44 = (_180_v26).multipliedBy(_180_v26);
          let _rhs45 = _159_v8;
          let _rhs46 = _module.__default.fm0(new BigNumber((_156_v5).length), _152_v1, _166_v14, _dafny.Map.Empty.slice().updateUnsafe(_160_v9,_155_v4), _162_globalState);
          let _rhs47 = _180_v26;
          let _lhs31 = _161_v10;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_161_v10).length));
          let _lhs33 = _188_v31;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_188_v31).length));
          let _lhs35 = _189_v32;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_189_v32).length));
          _lhs31[_lhs32] = _rhs44;
          _lhs33[_lhs34] = _rhs45;
          _lhs35[_lhs36] = _rhs46;
          _152_v1 = _rhs47;
          let _190_v33;
          _190_v33 = _dafny.Seq.of(_160_v9, _160_v9);
          let _191_v34;
          _191_v34 = _dafny.MultiSet.fromElements(_160_v9);
          _182_v28 = (_191_v34).IsSubsetOf(_dafny.MultiSet.FromArray(_190_v33));
          let _index34 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_189_v32).length));
          (_189_v32)[_index34] = _182_v28;
          let _192_v35;
          let _193_v36;
          let _194_v37;
          let _195_v38;
          let _out12;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = _module.__default.m0(_162_globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _out15 = _outcollector3[3];
          _192_v35 = _out12;
          _193_v36 = _out13;
          _194_v37 = _out14;
          _195_v38 = _out15;
        } else {
          _182_v28 = _155_v4;
          (_162_globalState).f2 = _180_v26;
          let _index35 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_161_v10).length));
          (_161_v10)[_index35] = _152_v1;
          let _index36 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_161_v10).length));
          (_161_v10)[_index36] = (_152_v1).multipliedBy(new BigNumber((_dafny.Seq.Concat(_160_v9, _160_v9)).length));
          _182_v28 = ((false) ? (_182_v28) : (_182_v28));
          let _196_v39;
          _196_v39 = new _dafny.CodePoint('r'.codePointAt(0));
          let _197_v40;
          _197_v40 = _dafny.Map.Empty.slice().updateUnsafe(_196_v39,_181_v27);
          let _index37 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_161_v10).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_161_v10).length));
          let _rhs48 = _155_v4;
          let _rhs49 = new BigNumber(137);
          let _rhs50 = ((new BigNumber(-485)).multipliedBy(_180_v26)).isLessThanOrEqualTo(_180_v26);
          let _rhs51 = true;
          let _rhs52 = (new BigNumber((_197_v40).length)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("d")).length)));
          let _lhs37 = _161_v10;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_161_v10).length));
          let _lhs39 = _161_v10;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_161_v10).length));
          _182_v28 = _rhs48;
          _lhs37[_lhs38] = _rhs49;
          _155_v4 = _rhs50;
          _182_v28 = _rhs51;
          _lhs39[_lhs40] = _rhs52;
        }
        _181_v27 = !(_155_v4) || ((_180_v26).isLessThanOrEqualTo(_152_v1));
      }
      let _hi0 = _152_v1;
      for (let _198_i2 = _152_v1; _198_i2.isLessThan(_hi0); _198_i2 = _198_i2.plus(_dafny.ONE)) {
        let _source5 = _159_v8;
        let _199___mcc_h0 = _source5;
        let _200_cf0 = _199___mcc_h0;
        _155_v4 = (false) || (_155_v4);
        let _201_v41;
        let _nw26 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _201_v41 = _nw26;
        let _202_v42;
        _202_v42 = _dafny.Seq.of(_161_v10);
        let _203_v43;
        _203_v43 = _dafny.Map.Empty.slice().updateUnsafe(_155_v4,_155_v4);
        let _rhs53 = (_202_v42)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_157_v6)[_module.__default.safeIndex(new BigNumber((_module.__default.fm2(!(_155_v4), _200_cf0, new BigNumber((_203_v43).length), _155_v4, _162_globalState)).length), new BigNumber((_157_v6).length))],false)).length), new BigNumber((_202_v42).length))];
        let _rhs54 = _201_v41;
        let _rhs55 = false;
        let _lhs41 = _162_globalState;
        _lhs41.f4 = _rhs53;
        _201_v41 = _rhs54;
        _155_v4 = _rhs55;
        _203_v43 = (_203_v43).update(_155_v4, _155_v4);
        _160_v9 = _160_v9;
        let _204_v44;
        _204_v44 = _dafny.Set.fromElements(_160_v9);
        _204_v44 = (_204_v44).Difference(_204_v44);
        (_162_globalState).f2 = (_152_v1).plus(_module.__default.safeDivisionInt(_198_i2, new BigNumber(566)));
        let _205_v45;
        let _nw27 = new _module.C0();
        _nw27.__ctor(false, _152_v1);
        _205_v45 = _nw27;
        let _206_v46;
        _206_v46 = _dafny.MultiSet.fromElements(_205_v45, _205_v45);
        let _207_v47;
        let _nw28 = Array((new BigNumber(27)).toNumber()).fill([]);
        _207_v47 = _nw28;
        let _208_v48;
        let _nw29 = new _module.C1();
        _nw29.__ctor((_206_v46).IsSubsetOf(_206_v46), _module.D3.create_DC5(_207_v47));
        _208_v48 = _nw29;
      }
      let _209_v49;
      let _nw30 = Array((new BigNumber(7)).toNumber()).fill([]);
      _209_v49 = _nw30;
      let _210_v50;
      _210_v50 = _module.D3.create_DC5(_209_v49);
      let _211_v51;
      let _nw31 = new _module.C1();
      _nw31.__ctor(_155_v4, ((_155_v4) ? (_210_v50) : (_210_v50)));
      _211_v51 = _nw31;
      if (_155_v4) {
        _160_v9 = _160_v9;
        let _212_v52;
        let _213_v53;
        let _214_v54;
        let _215_v55;
        let _out16;
        let _out17;
        let _out18;
        let _out19;
        let _outcollector4 = _module.__default.m0(_162_globalState);
        _out16 = _outcollector4[0];
        _out17 = _outcollector4[1];
        _out18 = _outcollector4[2];
        _out19 = _outcollector4[3];
        _212_v52 = _out16;
        _213_v53 = _out17;
        _214_v54 = _out18;
        _215_v55 = _out19;
        let _216_v56;
        let _nw32 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
        _216_v56 = _nw32;
        _216_v56 = _216_v56;
        _211_v51 = _211_v51;
        _214_v54 = false;
      } else {
        let _217_v57;
        let _nw33 = Array((new BigNumber(8)).toNumber()).fill(false);
        _217_v57 = _nw33;
        _217_v57 = _217_v57;
        _211_v51 = _211_v51;
        let _218_v58;
        let _nw34 = new _module.C2();
        _nw34.__ctor((_158_v7).IsSubsetOf(_dafny.Set.fromElements(new BigNumber(414), _152_v1, _152_v1, _152_v1, _152_v1)), true);
        _218_v58 = _nw34;
        _218_v58 = _218_v58;
        let _219_v59;
        _219_v59 = new _dafny.CodePoint('q'.codePointAt(0));
        let _220_v60;
        _220_v60 = _dafny.Map.Empty.slice().updateUnsafe(_219_v59,_152_v1);
        let _221_v61;
        _221_v61 = _dafny.Map.Empty.slice().updateUnsafe((_220_v60).Merge(_220_v60),_module.__default.safeModuloInt(new BigNumber(-347), _152_v1));
        _221_v61 = (_221_v61).Merge(_221_v61);
        let _222_v62;
        let _nw35 = Array((new BigNumber(10)).toNumber()).fill(_module.D3.Default());
        _222_v62 = _nw35;
        let _223_v63;
        _223_v63 = _module.D3.create_DC8((_211_v51).f16, _152_v1, _152_v1);
        let _index39 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_222_v62).length));
        (_222_v62)[_index39] = _223_v63;
        let _index40 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_222_v62).length));
        (_222_v62)[_index40] = _223_v63;
      }
      let _hi1 = _module.__default.safeModuloInt(_152_v1, _152_v1);
      for (let _224_i3 = _module.__default.fm17(!(false), _152_v1, _155_v4, _152_v1, _162_globalState); _224_i3.isLessThan(_hi1); _224_i3 = _224_i3.plus(_dafny.ONE)) {
        let _225_v64;
        _225_v64 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_152_v1),_152_v1);
        let _226_v65;
        _226_v65 = _dafny.Set.fromElements(_160_v9, _160_v9);
        _225_v64 = (_225_v64).update(new BigNumber(((_226_v65).Difference(_226_v65)).length), (_224_i3).minus(new BigNumber((_225_v64).length)));
        let _227_v66;
        let _nw36 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _227_v66 = _nw36;
        let _228_v67;
        _228_v67 = _dafny.Map.Empty.slice().updateUnsafe(_227_v66,_224_i3);
        _228_v67 = (_228_v67).update(_227_v66, _152_v1);
        let _index41 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_161_v10).length));
        (_161_v10)[_index41] = _224_i3;
        let _index42 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_161_v10).length));
        (_161_v10)[_index42] = _224_i3;
        let _rhs56 = _module.__default.fm2(_155_v4, (_dafny.ZERO).minus((_161_v10)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_161_v10).length))]), _152_v1, !(_155_v4), _162_globalState);
        let _rhs57 = _module.__default.safeDivisionInt(new BigNumber(897), _224_i3);
        let _rhs58 = (_161_v10)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_161_v10).length))];
        let _lhs42 = _162_globalState;
        let _lhs43 = _162_globalState;
        _160_v9 = _rhs56;
        _lhs42.f2 = _rhs57;
        _lhs43.f2 = _rhs58;
      }
      if (_155_v4) {
        let _229_v68;
        let _nw37 = new _module.C2();
        _nw37.__ctor((_211_v51).f16, (_211_v51).f16);
        _229_v68 = _nw37;
        let _230_v69;
        _230_v69 = _module.D4.create_DC10(_229_v68, new BigNumber(84), _152_v1);
        let _source6 = _230_v69;
        if (_source6.is_DC10) {
          let _231___mcc_h1 = (_source6).cf17;
          let _232___mcc_h2 = (_source6).cf18;
          let _233___mcc_h3 = (_source6).cf19;
          let _234_cf19 = _233___mcc_h3;
          let _235_cf18 = _232___mcc_h2;
          let _236_cf17 = _231___mcc_h1;
          _160_v9 = _160_v9;
          let _237_v70;
          _237_v70 = _module.D8.create_DC18();
          let _238_v71;
          _238_v71 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_235_cf18),_237_v70);
          let _239_v72;
          _239_v72 = new _dafny.CodePoint('a'.codePointAt(0));
          let _rhs59 = ((_dafny.Seq.IsProperPrefixOf((_dafny.Seq.UnicodeFromString("bonxg")), _dafny.Seq.update(_160_v9, _module.__default.safeIndex(_152_v1, new BigNumber((_160_v9).length)), _239_v72))) ? ((_211_v51).f16) : ((_211_v51).f16));
          let _rhs60 = _238_v71;
          let _rhs61 = ((_234_cf19).minus(_152_v1)).isLessThanOrEqualTo(_152_v1);
          _155_v4 = _rhs59;
          _238_v71 = _rhs60;
          _155_v4 = _rhs61;
          let _240_v73;
          let _nw38 = Array((new BigNumber(26)).toNumber());
          _nw38[(_dafny.ZERO).toNumber()] = _211_v51;
          _nw38[(_dafny.ONE).toNumber()] = _211_v51;
          _nw38[(new BigNumber(2)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(3)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(4)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(5)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(6)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(7)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(8)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(9)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(10)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(11)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(12)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(13)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(14)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(15)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(16)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(17)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(18)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(19)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(20)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(21)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(22)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(23)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(24)).toNumber()] = _211_v51;
          _nw38[(new BigNumber(25)).toNumber()] = ((_155_v4) ? (_211_v51) : (_211_v51));
          _240_v73 = _nw38;
          let _index43 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_240_v73).length));
          (_240_v73)[_index43] = _211_v51;
          let _index44 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_240_v73).length));
          (_240_v73)[_index44] = _211_v51;
          let _241_v74;
          _241_v74 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_152_v1),_235_cf18);
          _241_v74 = (_241_v74).update((_152_v1).plus(_152_v1), _152_v1);
        } else {
          let _242___mcc_h4 = (_source6).cf16;
          let _243_cf16 = _242___mcc_h4;
          _155_v4 = (_211_v51).f16;
          (_162_globalState).f2 = new BigNumber(412);
          let _244_v75;
          let _245_v76;
          let _out20;
          let _out21;
          let _outcollector5 = (_211_v51).m5(_152_v1, _162_globalState);
          _out20 = _outcollector5[0];
          _out21 = _outcollector5[1];
          _244_v75 = _out20;
          _245_v76 = _out21;
          let _246_v77;
          let _247_v78;
          let _out22;
          let _out23;
          let _outcollector6 = (_211_v51).m5((_154_v3)[_module.__default.safeIndex(_152_v1, new BigNumber((_154_v3).length))], _162_globalState);
          _out22 = _outcollector6[0];
          _out23 = _outcollector6[1];
          _246_v77 = _out22;
          _247_v78 = _out23;
        }
        let _248_v79;
        let _nw39 = Array((new BigNumber(6)).toNumber()).fill(false);
        _248_v79 = _nw39;
        let _index45 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
        (_248_v79)[_index45] = (((_211_v51).f16) ? ((_211_v51).f16) : ((_211_v51).f16));
        let _index46 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
        (_248_v79)[_index46] = !((_211_v51).f16);
        (_162_globalState).f2 = _152_v1;
        let _source7 = _module.D11.create_DC26();
        if (_source7.is_DC25) {
          let _249___mcc_h5 = (_source7).cf37;
          let _250_cf37 = _249___mcc_h5;
          let _251_v80;
          _251_v80 = _module.D6.create_DC13((_248_v79)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length))]);
          let _252_v81;
          let _nw40 = new _module.C1();
          _nw40.__ctor((_251_v80).dtor_cf22, _211_v51.f17);
          _252_v81 = _nw40;
          let _index47 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
          (_248_v79)[_index47] = (_248_v79)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length))];
          let _index48 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
          (_248_v79)[_index48] = (_211_v51).f16;
          let _253_v82;
          let _nw41 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Set.Empty);
          _253_v82 = _nw41;
          let _254_v83;
          _254_v83 = _dafny.Map.Empty.slice().updateUnsafe(_154_v3,_152_v1);
          let _255_v84;
          _255_v84 = _dafny.MultiSet.fromElements(_250_cf37);
          let _index49 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_253_v82).length));
          (_253_v82)[_index49] = function () {
            let _coll7 = new _dafny.Set();
            for (const _compr_7 of ((_254_v83).update(_154_v3, new BigNumber((_255_v84).cardinality()))).Keys.Elements) {
              let _256_v85 = _compr_7;
              if (((_254_v83).update(_154_v3, new BigNumber((_255_v84).cardinality()))).contains(_256_v85)) {
                _coll7.add(_256_v85);
              }
            }
            return _coll7;
          }();
          let _index50 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_253_v82).length));
          (_253_v82)[_index50] = _dafny.Set.fromElements(_154_v3, _dafny.Seq.Concat(_dafny.Seq.update(_154_v3, _module.__default.safeIndex(_152_v1, new BigNumber((_154_v3).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), ((_257_cf37) => function (_258_i4) {
            return _257_cf37;
          })(_250_cf37))).length)), _154_v3), _module.__default.fm25(_152_v1, _152_v1, _152_v1, new BigNumber((_158_v7).length), _162_globalState));
        } else if (_source7.is_DC26) {
          let _index51 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_161_v10).length));
          (_161_v10)[_index51] = new BigNumber((_160_v9).length);
          let _index52 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_161_v10).length));
          (_161_v10)[_index52] = _152_v1;
          let _index53 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
          (_248_v79)[_index53] = (_211_v51).f16;
          let _259_v86;
          _259_v86 = new _dafny.CodePoint('b'.codePointAt(0));
          _155_v4 = (((_248_v79)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length))]) ? (!(!((_211_v51).fm4(_152_v1, (_211_v51).f16, _259_v86, _162_globalState)))) : ((_248_v79)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length))]));
          let _260_v87;
          let _nw42 = Array((_dafny.ONE).toNumber());
          _nw42[(_dafny.ZERO).toNumber()] = _160_v9;
          _260_v87 = _nw42;
          let _index54 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_260_v87).length));
          (_260_v87)[_index54] = _160_v9;
          let _261_v88;
          let _nw43 = new _module.C2();
          _nw43.__ctor((_211_v51).f16, (_248_v79)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length))]);
          _261_v88 = _nw43;
          let _262_v89;
          _262_v89 = _dafny.Map.Empty.slice().updateUnsafe(_261_v88,_261_v88.f14);
          let _263_v90;
          _263_v90 = _160_v9;
          let _264_v91;
          _264_v91 = _module.D12.create_DC28(_262_v89);
          let _265_v92;
          _265_v92 = _dafny.Map.Empty.slice().updateUnsafe(_152_v1,_211_v51);
          let _index55 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_260_v87).length));
          let _rhs62 = _dafny.Seq.Concat(_dafny.Seq.Concat(_160_v9, (_263_v90)), _160_v9);
          let _rhs63 = (_264_v91).dtor_cf39;
          let _rhs64 = (_161_v10)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_161_v10).length))];
          let _rhs65 = ((_155_v4) ? (_211_v51) : ((((_265_v92).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_157_v6).length))))) ? ((_265_v92).get((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_157_v6).length))))) : (_211_v51))));
          let _rhs66 = _259_v86;
          let _lhs44 = _260_v87;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_260_v87).length));
          let _lhs46 = _162_globalState;
          _lhs44[_lhs45] = _rhs62;
          _262_v89 = _rhs63;
          _lhs46.f2 = _rhs64;
          _211_v51 = _rhs65;
          _259_v86 = _rhs66;
        } else if (_source7.is_DC24) {
          let _266___mcc_h6 = (_source7).cf36;
          let _267_cf36 = _266___mcc_h6;
          let _268_v93;
          _268_v93 = _dafny.Set.fromElements((_211_v51).f16);
          let _269_v94;
          _269_v94 = _dafny.Map.Empty.slice().updateUnsafe(_152_v1,_152_v1);
          let _index56 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
          (_248_v79)[_index56] = ((_module.__default.fm23(_269_v94, (_211_v51).f16, _162_globalState)).Union(_dafny.Set.fromElements(true, (_211_v51).f16))).IsProperSubsetOf(_268_v93);
          let _270_v95;
          _270_v95 = _dafny.Map.Empty.slice().updateUnsafe(false,_152_v1);
          _270_v95 = (_270_v95).update((((_211_v51).f16) ? (_155_v4) : ((_211_v51).f16)), _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(_152_v1)), _152_v1));
          (_162_globalState).f2 = new BigNumber((_156_v5).length);
          (_162_globalState).f2 = _152_v1;
        } else {
          let _271___mcc_h7 = (_source7).cf38;
          let _272_cf38 = _271___mcc_h7;
          let _273_v96;
          _273_v96 = new _dafny.CodePoint('y'.codePointAt(0));
          _273_v96 = _273_v96;
          let _274_v97;
          let _nw44 = new _module.C3();
          _nw44.__ctor((((_211_v51).f16) ? (_152_v1) : ((_152_v1))), _module.__default.fm27(_162_globalState));
          _274_v97 = _nw44;
          let _index57 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
          let _index58 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
          let _rhs67 = (_module.__default.fm19(_274_v97.f10, _162_globalState)).isEqualTo(_274_v97.f10);
          let _rhs68 = (_152_v1).isLessThanOrEqualTo(_274_v97.f10);
          let _rhs69 = _274_v97;
          let _lhs47 = _248_v79;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
          let _lhs49 = _248_v79;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
          _lhs47[_lhs48] = _rhs67;
          _lhs49[_lhs50] = _rhs68;
          _274_v97 = _rhs69;
          let _index59 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
          (_248_v79)[_index59] = (_211_v51).f16;
          let _index60 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_161_v10).length));
          (_161_v10)[_index60] = _152_v1;
          let _index61 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_161_v10).length));
          (_161_v10)[_index61] = _274_v97.f10;
        }
        let _275_v98;
        _275_v98 = _dafny.Map.Empty.slice().updateUnsafe((_248_v79)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length))],new BigNumber((_dafny.Seq.update(_154_v3, _module.__default.safeIndex(_152_v1, new BigNumber((_154_v3).length)), _152_v1)).length));
        let _276_v99;
        _276_v99 = new _dafny.CodePoint('e'.codePointAt(0));
        let _index62 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
        let _rhs70 = (_152_v1).isLessThanOrEqualTo(_152_v1);
        let _rhs71 = (_248_v79)[_module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length))];
        let _rhs72 = _160_v9;
        let _rhs73 = (_229_v68).fm4((_dafny.ZERO).minus(_152_v1), _dafny.areEqual(_160_v9, _160_v9), _276_v99, _162_globalState);
        let _rhs74 = _275_v98;
        let _lhs51 = _248_v79;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_248_v79).length));
        _155_v4 = _rhs70;
        _155_v4 = _rhs71;
        _160_v9 = _rhs72;
        _lhs51[_lhs52] = _rhs73;
        _275_v98 = _rhs74;
      } else {
        let _277_v100;
        let _nw45 = Array((new BigNumber(29)).toNumber()).fill(false);
        _277_v100 = _nw45;
        let _278_v101;
        _278_v101 = _module.D12.create_DC29(_277_v100, _155_v4);
        _155_v4 = (_278_v101).dtor_cf41;
        let _279_v102;
        let _nw46 = new _module.C2();
        _nw46.__ctor((_152_v1).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.update(_154_v3, _module.__default.safeIndex(new BigNumber((_160_v9).length), new BigNumber((_154_v3).length)), _152_v1)).length)), (_152_v1).isLessThanOrEqualTo(_152_v1));
        _279_v102 = _nw46;
        let _280_v103;
        _280_v103 = _dafny.Map.Empty.slice().updateUnsafe(_152_v1,_156_v5);
        let _rhs75 = _279_v102;
        let _rhs76 = _module.__default.fm28(new BigNumber((_dafny.Seq.Concat(_157_v6, _dafny.Seq.of((_211_v51).f16, _module.__default.fm0(new BigNumber(891), _152_v1, _166_v14, _dafny.Map.Empty.slice().updateUnsafe(_160_v9,(_211_v51).f16), _162_globalState)))).length), true, _162_globalState);
        _279_v102 = _rhs75;
        _280_v103 = _rhs76;
        _155_v4 = (_module.D1.create_DC1(_155_v4)).dtor_cf1;
        let _281_v104;
        let _nw47 = new _module.C3();
        _nw47.__ctor(new BigNumber((_154_v3).length), _156_v5);
        _281_v104 = _nw47;
        let _282_v105;
        let _nw48 = new _module.C2();
        _nw48.__ctor(_155_v4, (_157_v6)[_module.__default.safeIndex(_281_v104.f10, new BigNumber((_157_v6).length))]);
        _282_v105 = _nw48;
        let _283_v106;
        _283_v106 = _dafny.Map.Empty.slice().updateUnsafe(_281_v104,_282_v105);
        let _284_v107;
        _284_v107 = _dafny.Seq.of(_283_v106);
        let _285_v108;
        let _nw49 = new _module.C0();
        _nw49.__ctor((_211_v51).f16, new BigNumber((_284_v107).length));
        _285_v108 = _nw49;
        if (!((_211_v51).f16)) {
          let _rhs77 = (((_211_v51).f16) ? ((_211_v51).f16) : (_282_v105.f14));
          let _rhs78 = ((new BigNumber(771)).multipliedBy(new BigNumber(265))).minus((_285_v108).f13);
          let _rhs79 = _277_v100;
          let _rhs80 = _285_v108;
          let _lhs53 = _282_v105;
          let _lhs54 = _162_globalState;
          _lhs53.f14 = _rhs77;
          _lhs54.f2 = _rhs78;
          _277_v100 = _rhs79;
          _285_v108 = _rhs80;
          let _286_v109;
          _286_v109 = new _dafny.CodePoint('s'.codePointAt(0));
          let _rhs81 = _281_v104.f10;
          let _rhs82 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(111)), function (_287_i5) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          }), _module.__default.fm2((_211_v51).f16, (_285_v108).f13, (_dafny.ZERO).minus((_285_v108).f13), _282_v105.f14, _162_globalState)), _module.__default.safeIndex(new BigNumber((_157_v6).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(111)), function (_288_i5) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          }), _module.__default.fm2((_211_v51).f16, (_285_v108).f13, (_dafny.ZERO).minus((_285_v108).f13), _282_v105.f14, _162_globalState))).length)), _286_v109);
          let _rhs83 = _160_v9;
          let _lhs55 = _162_globalState;
          _lhs55.f2 = _rhs81;
          _160_v9 = _rhs82;
          _160_v9 = _rhs83;
          let _index63 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_277_v100).length));
          (_277_v100)[_index63] = !(true);
          let _index64 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_277_v100).length));
          (_277_v100)[_index64] = (_282_v105).f15;
          let _289_v110;
          let _nw50 = new _module.C3();
          _nw50.__ctor((_154_v3)[_module.__default.safeIndex((_285_v108).f13, new BigNumber((_154_v3).length))], ((_281_v104).f11).update((_285_v108).f13, _155_v4));
          _289_v110 = _nw50;
          _289_v110 = _289_v110;
          let _290_v111;
          let _nw51 = new _module.C0();
          _nw51.__ctor(!((_285_v108).f12), _281_v104.f10);
          _290_v111 = _nw51;
        } else {
          (_282_v105).f14 = (_282_v105).f15;
          let _291_v112;
          _291_v112 = _module.D10.create_DC23((_285_v108).f13, _282_v105.f14, (_211_v51).f16);
          let _292_v113;
          _292_v113 = _dafny.Map.Empty.slice().updateUnsafe(((_291_v112).dtor_cf33).isLessThanOrEqualTo(new BigNumber((_157_v6).length)),(_282_v105.f14) && ((_285_v108).f12));
          _292_v113 = (_292_v113).Merge(_module.__default.fm22(_162_globalState));
          (_162_globalState).f2 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_152_v1, _152_v1));
          (_281_v104).f10 = (_285_v108).f13;
          let _293_v114;
          let _294_v115;
          let _295_v116;
          let _296_v117;
          let _out24;
          let _out25;
          let _out26;
          let _out27;
          let _outcollector7 = _module.__default.m0(_162_globalState);
          _out24 = _outcollector7[0];
          _out25 = _outcollector7[1];
          _out26 = _outcollector7[2];
          _out27 = _outcollector7[3];
          _293_v114 = _out24;
          _294_v115 = _out25;
          _295_v116 = _out26;
          _296_v117 = _out27;
        }
      }
      let _index65 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length));
      (_161_v10)[_index65] = new BigNumber(192);
      let _index66 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length));
      (_161_v10)[_index66] = _152_v1;
      let _297_v118;
      let _nw52 = new _module.C3();
      _nw52.__ctor(_152_v1, _156_v5);
      _297_v118 = _nw52;
      let _298_v119;
      _298_v119 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),_297_v118);
      let _299_v120;
      _299_v120 = new _dafny.CodePoint('j'.codePointAt(0));
      _298_v119 = (_298_v119).update(_299_v120, _297_v118);
      let _300_v121;
      let _nw53 = Array((new BigNumber(5)).toNumber());
      _nw53[(_dafny.ZERO).toNumber()] = _155_v4;
      _nw53[(_dafny.ONE).toNumber()] = ((true) ? (_155_v4) : (_155_v4));
      _nw53[(new BigNumber(2)).toNumber()] = (_211_v51).f16;
      _nw53[(new BigNumber(3)).toNumber()] = (_211_v51).f16;
      _nw53[(new BigNumber(4)).toNumber()] = (_155_v4) || (_155_v4);
      _300_v121 = _nw53;
      let _index67 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length));
      (_300_v121)[_index67] = true;
      let _301_v122;
      _301_v122 = _module.D9.create_DC20(_155_v4, _299_v120);
      let _302_v123;
      _302_v123 = _dafny.MultiSet.fromElements(_301_v122, _301_v122, _301_v122, _301_v122);
      let _index68 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length));
      (_300_v121)[_index68] = (_302_v123).IsDisjointFrom(_302_v123);
      let _303_i6;
      _303_i6 = _dafny.ZERO;
      L0: {
        while (_155_v4) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_303_i6)) {
              break L0;
            }
            _303_i6 = (_303_i6).plus(_dafny.ONE);
            let _index69 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length));
            (_161_v10)[_index69] = new BigNumber((_160_v9).length);
            let _304_v125;
            _304_v125 = _dafny.MultiSet.fromElements(_299_v120, _299_v120);
            let _305_v126;
            _305_v126 = _dafny.Map.Empty.slice().updateUnsafe((_211_v51).f16,_module.__default.fm29(_152_v1, (function () {
              let _coll8 = new _dafny.Map();
              for (const _compr_8 of (_304_v125).Elements) {
                let _306_v124 = _compr_8;
                if ((_304_v125).contains(_306_v124)) {
                  _coll8.push([_306_v124,_152_v1]);
                }
              }
              return _coll8;
            }()).update(_299_v120, new BigNumber((_154_v3).length)), new BigNumber((_160_v9).length), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_211_v51).f16))).cardinality()), _162_globalState));
            let _307_v127;
            _307_v127 = _module.D3.create_DC7(new BigNumber((_156_v5).length));
            let _308_v128;
            _308_v128 = _dafny.Map.Empty.slice().updateUnsafe((_297_v118).fm7(_305_v126, true, _152_v1, _162_globalState),_307_v127);
            let _309_v129;
            _309_v129 = _dafny.Map.Empty.slice().updateUnsafe((_211_v51).f16,new BigNumber((_160_v9).length));
            let _310_v130;
            _310_v130 = _module.D4.create_DC9(_module.__default.fm14(_308_v128, _155_v4, _309_v129, (_300_v121)[_module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length))], _162_globalState));
            _310_v130 = _310_v130;
            let _311_v131;
            _311_v131 = _dafny.Set.fromElements(_dafny.Seq.Concat(_157_v6, _157_v6), _157_v6);
            let _312_v132;
            _312_v132 = _dafny.Seq.of(_309_v129);
            let _313_v133;
            _313_v133 = _dafny.Map.Empty.slice().updateUnsafe(_155_v4,_dafny.Seq.update(_157_v6, _module.__default.safeIndex(_152_v1, new BigNumber((_157_v6).length)), (_300_v121)[_module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length))]));
            let _314_v134;
            _314_v134 = _dafny.Seq.of((((_313_v133).contains(false)) ? ((_313_v133).get(false)) : (_157_v6)), _157_v6);
            let _315_v136;
            _315_v136 = _dafny.Seq.of((_dafny.Set.fromElements(_dafny.Seq.of((_211_v51).f16), _dafny.Seq.of(_155_v4, (_211_v51).f16), _157_v6, _dafny.Seq.of((_297_v118).fm5(_299_v120, _312_v132, _155_v4, _162_globalState)), _157_v6)).Union(_311_v131), ((_module.D13.create_DC31(_311_v131)).dtor_cf43).Intersect(function () {
              let _coll9 = new _dafny.Set();
              for (const _compr_9 of (_314_v134).Elements) {
                let _316_v135 = _compr_9;
                if (_dafny.Seq.contains(_314_v134, _316_v135)) {
                  _coll9.add(_316_v135);
                }
              }
              return _coll9;
            }()), _311_v131, _311_v131);
            let _317_v137;
            _317_v137 = _dafny.Map.Empty.slice().updateUnsafe((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))],(_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))]);
            let _318_v138;
            _318_v138 = _dafny.Map.Empty.slice().updateUnsafe((_211_v51).f16,_317_v137);
            _311_v131 = (_315_v136)[_module.__default.safeIndex(new BigNumber(((_318_v138).Merge(_318_v138)).length), new BigNumber((_315_v136).length))];
            let _319_v139;
            _319_v139 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_154_v3)[_module.__default.safeIndex(_152_v1, new BigNumber((_154_v3).length))],_299_v120)).length)), _153_v2, _153_v2);
            _319_v139 = _319_v139;
          }
        }
      }
      let _index70 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length));
      (_300_v121)[_index70] = true;
      let _320_v140;
      let _nw54 = Array((new BigNumber(18)).toNumber());
      _320_v140 = _nw54;
      let _321_v141;
      _321_v141 = _dafny.Seq.of(_320_v140);
      if ((new BigNumber((_321_v141).length)).isLessThanOrEqualTo(_152_v1)) {
        let _322_v142;
        _322_v142 = _module.D14.create_DC35(_dafny.Seq.of(_211_v51));
        let _323_v143;
        _323_v143 = _dafny.Seq.of(_211_v51);
        let _pat_let_tv0 = _323_v143;
        let _pat_let_tv1 = _323_v143;
        _155_v4 = !(((_158_v7).Difference(_158_v7)).IsSubsetOf((_158_v7).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.update((function (_pat_let0_0) {
          return function (_324_dt__update__tmp_h5) {
            return function (_pat_let1_0) {
              return function (_325_dt__update_hcf50_h0) {
                return _module.D14.create_DC35(_325_dt__update_hcf50_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_322_v142)).dtor_cf50, _module.__default.safeIndex(_152_v1, new BigNumber(((function (_pat_let2_0) {
          return function (_326_dt__update__tmp_h6) {
            return function (_pat_let3_0) {
              return function (_327_dt__update_hcf50_h1) {
                return _module.D14.create_DC35(_327_dt__update_hcf50_h1);
              }(_pat_let3_0);
            }(_pat_let_tv1);
          }(_pat_let2_0);
        }(_322_v142)).dtor_cf50).length)), _211_v51)).length)))));
        let _index71 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length));
        (_161_v10)[_index71] = new BigNumber((function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of _dafny.IntegerRange(new BigNumber(63), new BigNumber(992))) {
            let _328_v144 = _compr_10;
            if (((new BigNumber(63)).isLessThanOrEqualTo(_328_v144)) && ((_328_v144).isLessThan(new BigNumber(992)))) {
              _coll10.add(_module.__default.safeModuloInt(_328_v144, new BigNumber((_160_v9).length)));
            }
          }
          return _coll10;
        }()).length);
        let _329_v145;
        _329_v145 = _module.D11.create_DC26();
        _329_v145 = _329_v145;
        let _330_v146;
        let _331_v147;
        let _out28;
        let _out29;
        let _outcollector8 = (_211_v51).m5((_dafny.ZERO).minus((_152_v1).minus(new BigNumber((_160_v9).length))), _162_globalState);
        _out28 = _outcollector8[0];
        _out29 = _outcollector8[1];
        _330_v146 = _out28;
        _331_v147 = _out29;
        if ((_211_v51).f16) {
          let _332_v148;
          let _nw55 = Array((new BigNumber(5)).toNumber());
          _332_v148 = _nw55;
          _332_v148 = ((true) ? (_332_v148) : (_332_v148));
          let _rhs84 = _331_v147;
          let _rhs85 = !((_211_v51).f16);
          let _rhs86 = (_dafny.ZERO).minus((_152_v1).minus(_152_v1));
          _331_v147 = _rhs84;
          _155_v4 = _rhs85;
          _152_v1 = _rhs86;
          let _333_v149;
          _333_v149 = _dafny.Seq.of(_297_v118);
          let _334_v150;
          _334_v150 = _dafny.Map.Empty.slice().updateUnsafe((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))],(_333_v149)[_module.__default.safeIndex((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))], new BigNumber((_333_v149).length))]);
          _334_v150 = (_334_v150).update(_module.__default.fm10((_211_v51).f16, _162_globalState), _297_v118);
          _155_v4 = true;
          let _335_v151;
          _335_v151 = _dafny.Seq.of(_dafny.Seq.update(_160_v9, _module.__default.safeIndex(_330_v146, new BigNumber((_160_v9).length)), _299_v120));
          _152_v1 = (((_153_v2).contains(_152_v1)) ? ((_153_v2).get(_152_v1)) : (new BigNumber(((_335_v151)[_module.__default.safeIndex(_330_v146, new BigNumber((_335_v151).length))]).length)));
        } else {
          let _336_v152;
          let _337_v153;
          let _out30;
          let _out31;
          let _outcollector9 = (_211_v51).m5((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))], _162_globalState);
          _out30 = _outcollector9[0];
          _out31 = _outcollector9[1];
          _336_v152 = _out30;
          _337_v153 = _out31;
          let _338_v154;
          let _nw56 = new _module.C3();
          _nw56.__ctor(_337_v153, _156_v5);
          _338_v154 = _nw56;
          let _339_v155;
          _339_v155 = _dafny.MultiSet.fromElements(_338_v154);
          let _340_v156;
          _340_v156 = _module.D15.create_DC38(_338_v154);
          let _341_v157;
          _341_v157 = _dafny.MultiSet.fromElements(_339_v155, _339_v155, _dafny.MultiSet.FromArray(_dafny.Seq.of((_340_v156).dtor_cf57, _338_v154, _338_v154)), _339_v155);
          _341_v157 = _341_v157;
          let _342_v159;
          _342_v159 = _module.D7.create_DC15(new BigNumber((function () {
  let _coll11 = new _dafny.Set();
  for (const _compr_11 of _dafny.IntegerRange(new BigNumber(941), new BigNumber(43))) {
    let _343_v158 = _compr_11;
    if (((new BigNumber(941)).isLessThanOrEqualTo(_343_v158)) && ((_343_v158).isLessThan(new BigNumber(43)))) {
      _coll11.add((_343_v158).plus(_331_v147));
    }
  }
  return _coll11;
}()).length), (_211_v51).f16);
          let _344_v160;
          let _nw57 = new _module.C2();
          _nw57.__ctor((((_211_v51).f16) ? ((_342_v159).dtor_cf25) : ((_211_v51).f16)), !(_155_v4));
          _344_v160 = _nw57;
          let _345_v161;
          _345_v161 = _dafny.Map.Empty.slice().updateUnsafe(_155_v4,(_211_v51).f16);
          let _346_v162;
          _346_v162 = _dafny.Map.Empty.slice().updateUnsafe(_300_v121,!((((_345_v161).contains(_155_v4)) ? ((_345_v161).get(_155_v4)) : ((_211_v51).f16))));
          let _rhs87 = (_211_v51).f16;
          let _rhs88 = _344_v160;
          let _rhs89 = (_346_v162).Merge(_346_v162);
          _155_v4 = _rhs87;
          _344_v160 = _rhs88;
          _346_v162 = _rhs89;
          let _index72 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length));
          (_161_v10)[_index72] = _331_v147;
          let _347_v163;
          _347_v163 = _dafny.Set.fromElements(_299_v120);
          let _index73 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length));
          (_300_v121)[_index73] = !((_dafny.Set.fromElements(_299_v120)).IsSubsetOf((_347_v163).Union(_347_v163)));
        }
      } else {
        let _348_v164;
        _348_v164 = _module.D12.create_DC29(_300_v121, _155_v4);
        let _349_v165;
        _349_v165 = _module.D1.create_DC1((_211_v51).f16);
        let _350_v166;
        _350_v166 = _dafny.Map.Empty.slice().updateUnsafe((_211_v51).f16,_349_v165);
        _348_v164 = ((((_297_v118).fm7(_350_v166, (_211_v51).f16, new BigNumber(575), _162_globalState)) && (true)) ? (_module.D12.create_DC29(_300_v121, (_211_v51).f16)) : (_348_v164));
        let _index74 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length));
        (_300_v121)[_index74] = _155_v4;
        let _351_v167;
        _351_v167 = _dafny.Set.fromElements(_156_v5);
        _155_v4 = !(_351_v167).equals((_351_v167).Intersect(_351_v167));
        if ((_157_v6)[_module.__default.safeIndex((_152_v1).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(!(_module.__default.fm0((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))], _152_v1, _166_v14, _dafny.Map.Empty.slice().updateUnsafe(_160_v9,_module.__default.fm0((_154_v3)[_module.__default.safeIndex(new BigNumber(-797), new BigNumber((_154_v3).length))], new BigNumber((_160_v9).length), _166_v14, _167_v15, _162_globalState)), _162_globalState)), (_211_v51).f16, _155_v4)).cardinality())), new BigNumber((_157_v6).length))]) {
          let _352_v168;
          let _nw58 = new _module.C1();
          _nw58.__ctor(!(!(!(_153_v2).equals(_dafny.MultiSet.fromElements(_152_v1)))), _210_v50);
          _352_v168 = _nw58;
          let _353_v169;
          let _nw59 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Set.Empty);
          _353_v169 = _nw59;
          let _354_v170;
          _354_v170 = _dafny.Set.fromElements(_160_v9);
          let _index75 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_353_v169).length));
          (_353_v169)[_index75] = (_354_v170).Intersect(_dafny.Set.fromElements(_160_v9));
          let _index76 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_353_v169).length));
          (_353_v169)[_index76] = _354_v170;
          _152_v1 = _module.__default.fm17(_155_v4, _152_v1, (_211_v51).f16, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))],false)).length), _162_globalState);
          _300_v121 = _300_v121;
          let _355_v171;
          _355_v171 = _dafny.Seq.of(_297_v118, _297_v118, _297_v118, _297_v118, _297_v118);
          let _356_v172;
          _356_v172 = _dafny.Set.fromElements(_355_v171);
          let _357_v173;
          _357_v173 = _dafny.Map.Empty.slice().updateUnsafe(((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))]).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(_152_v1))),(_dafny.ZERO).minus(new BigNumber((_356_v172).length)));
          _357_v173 = function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_154_v3).Elements) {
              let _358_v174 = _compr_12;
              if (_dafny.Seq.contains(_154_v3, _358_v174)) {
                _coll12.push([(_358_v174).minus((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))]),new BigNumber(39)]);
              }
            }
            return _coll12;
          }();
        } else {
          _155_v4 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("yx"), _160_v9);
          _155_v4 = (_300_v121)[_module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length))];
          let _359_v175;
          _359_v175 = _module.D16.create_DC41(_166_v14);
          (_162_globalState).f2 = new BigNumber(((_359_v175).dtor_cf64).length);
          let _360_v176;
          _360_v176 = _dafny.Map.Empty.slice().updateUnsafe((_211_v51).f16,(_211_v51).f16);
          let _361_v177;
          _361_v177 = _dafny.Map.Empty.slice().updateUnsafe((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))],(new BigNumber((_360_v176).length)).minus(new BigNumber((_154_v3).length)));
          _361_v177 = (_361_v177).update(new BigNumber(((_361_v177).Merge(_361_v177)).length), new BigNumber(881));
          let _362_v178;
          let _363_v179;
          let _364_v180;
          let _365_v181;
          let _out32;
          let _out33;
          let _out34;
          let _out35;
          let _outcollector10 = (_297_v118).m2(_162_globalState);
          _out32 = _outcollector10[0];
          _out33 = _outcollector10[1];
          _out34 = _outcollector10[2];
          _out35 = _outcollector10[3];
          _362_v178 = _out32;
          _363_v179 = _out33;
          _364_v180 = _out34;
          _365_v181 = _out35;
        }
        let _366_v182;
        let _nw60 = new _module.C2();
        _nw60.__ctor((_211_v51).f16, (_211_v51).f16);
        _366_v182 = _nw60;
        let _367_v183;
        _367_v183 = _dafny.Map.Empty.slice().updateUnsafe(_366_v182,_155_v4);
        let _source8 = _module.D12.create_DC28((_367_v183).Merge(_367_v183));
        if (_source8.is_DC29) {
          let _368___mcc_h8 = (_source8).cf40;
          let _369___mcc_h9 = (_source8).cf41;
          let _370_cf41 = _369___mcc_h9;
          let _371_cf40 = _368___mcc_h8;
          let _rhs90 = _152_v1;
          let _rhs91 = (_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))];
          let _lhs56 = _162_globalState;
          let _lhs57 = _162_globalState;
          _lhs56.f2 = _rhs90;
          _lhs57.f2 = _rhs91;
          _297_v118 = _297_v118;
          let _372_v184;
          let _373_v185;
          let _out36;
          let _out37;
          let _outcollector11 = (_211_v51).m5(new BigNumber(-863), _162_globalState);
          _out36 = _outcollector11[0];
          _out37 = _outcollector11[1];
          _372_v184 = _out36;
          _373_v185 = _out37;
          (_366_v182).m4((_366_v182).f15, _372_v184, _162_globalState);
        } else if (_source8.is_DC28) {
          let _374___mcc_h10 = (_source8).cf39;
          let _375_cf39 = _374___mcc_h10;
          let _376_v186;
          let _nw61 = new _module.C1();
          _nw61.__ctor((new BigNumber(-981)).isEqualTo(_module.__default.fm17(!((_366_v182).f15), new BigNumber(-403), (_211_v51).f16, (_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))], _162_globalState)), _211_v51.f17);
          _376_v186 = _nw61;
          _366_v182 = _366_v182;
          (_162_globalState).f2 = _152_v1;
          let _377_v187;
          _377_v187 = _161_v10;
          let _378_v188;
          let _nw62 = Array((new BigNumber(2)).toNumber());
          _nw62[(_dafny.ZERO).toNumber()] = _377_v187;
          _nw62[(_dafny.ONE).toNumber()] = _161_v10;
          _378_v188 = _nw62;
          let _index77 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_378_v188).length));
          (_378_v188)[_index77] = _377_v187;
          let _379_v190;
          let _nw63 = new _module.C2();
          _nw63.__ctor(!(_158_v7).equals(function () {
            let _coll13 = new _dafny.Set();
            for (const _compr_13 of (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10((_376_v186).f16, _162_globalState),_153_v2)).Keys.Elements) {
              let _380_v189 = _compr_13;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10((_376_v186).f16, _162_globalState),_153_v2)).contains(_380_v189)) {
                _coll13.add((_380_v189).minus(new BigNumber((_dafny.Set.fromElements(false)).length)));
              }
            }
            return _coll13;
          }()), true);
          _379_v190 = _nw63;
          let _index78 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_378_v188).length));
          let _index79 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length));
          let _rhs92 = ((!((_366_v182).f15)) ? (_377_v187) : (_377_v187));
          let _rhs93 = new BigNumber(286);
          let _rhs94 = _379_v190;
          let _rhs95 = _165_v13;
          let _rhs96 = _152_v1;
          let _lhs58 = _378_v188;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_378_v188).length));
          let _lhs60 = _161_v10;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length));
          let _lhs62 = _162_globalState;
          _lhs58[_lhs59] = _rhs92;
          _lhs60[_lhs61] = _rhs93;
          _379_v190 = _rhs94;
          _165_v13 = _rhs95;
          _lhs62.f2 = _rhs96;
        } else {
          let _381___mcc_h11 = (_source8).cf42;
          let _382_cf42 = _381___mcc_h11;
          let _383_v191;
          let _nw64 = new _module.C1();
          _nw64.__ctor((((_156_v5).contains(new BigNumber(128))) ? ((_156_v5).get(new BigNumber(128))) : (true)), _211_v51.f17);
          _383_v191 = _nw64;
          let _rhs97 = _383_v191;
          let _rhs98 = _module.__default.safeDivisionInt(_152_v1, (_dafny.ZERO).minus((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))]));
          let _rhs99 = ((_152_v1).minus((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))])).multipliedBy(_152_v1);
          let _rhs100 = _158_v7;
          let _lhs63 = _162_globalState;
          let _lhs64 = _162_globalState;
          _383_v191 = _rhs97;
          _lhs63.f2 = _rhs98;
          _lhs64.f2 = _rhs99;
          _158_v7 = _rhs100;
          let _384_v192;
          let _nw65 = new _module.C0();
          _nw65.__ctor(_366_v182.f14, _152_v1);
          _384_v192 = _nw65;
          let _385_v193;
          _385_v193 = _dafny.Map.Empty.slice().updateUnsafe(_366_v182.f14,_384_v192);
          let _386_v194;
          _386_v194 = _module.D13.create_DC33(_300_v121, _366_v182, _385_v193, (_211_v51).f16);
          let _387_v195;
          _387_v195 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((((_153_v2).contains((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))])) ? ((_153_v2).get((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))])) : (_module.__default.fm19(_152_v1, _162_globalState))))),(_386_v194).dtor_cf45);
          _387_v195 = (_387_v195).update(_152_v1, _300_v121);
          let _388_v196;
          let _389_v197;
          let _390_v198;
          let _391_v199;
          let _out38;
          let _out39;
          let _out40;
          let _out41;
          let _outcollector12 = (_297_v118).m2(_162_globalState);
          _out38 = _outcollector12[0];
          _out39 = _outcollector12[1];
          _out40 = _outcollector12[2];
          _out41 = _outcollector12[3];
          _388_v196 = _out38;
          _389_v197 = _out39;
          _390_v198 = _out40;
          _391_v199 = _out41;
          (_162_globalState).f2 = (_dafny.ZERO).minus((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))]);
        }
      }
      let _392_i7;
      _392_i7 = _dafny.ZERO;
      L1: {
        while ((_300_v121)[_module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length))]) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_392_i7)) {
              break L1;
            }
            _392_i7 = (_392_i7).plus(_dafny.ONE);
            let _393_v200;
            let _nw66 = Array((new BigNumber(13)).toNumber()).fill(_module.D3.Default());
            _393_v200 = _nw66;
            let _394_v201;
            _394_v201 = _module.D3.create_DC6((_161_v10)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_161_v10).length))], (_211_v51).f16, (_211_v51).f16, _152_v1, (_211_v51).f16);
            let _index80 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_393_v200).length));
            (_393_v200)[_index80] = _394_v201;
            let _index81 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_393_v200).length));
            (_393_v200)[_index81] = _394_v201;
            _155_v4 = (_300_v121)[_module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length))];
            let _395_v202;
            _395_v202 = _dafny.Set.fromElements((_211_v51).f16);
            let _396_v203;
            _396_v203 = _dafny.Seq.of((_395_v202).Union(_395_v202), _395_v202);
            _396_v203 = _dafny.Seq.Concat(_396_v203, _dafny.Seq.of(_dafny.Set.fromElements((_211_v51).f16), _395_v202, _395_v202, _395_v202));
            _155_v4 = _dafny.Seq.IsProperPrefixOf(_160_v9, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yecri"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(156)), function (_397_i8) {
              return new _dafny.CodePoint('c'.codePointAt(0));
            })));
          }
        }
      }
      _155_v4 = (_300_v121)[_module.__default.safeIndex(new BigNumber(366), new BigNumber((_300_v121).length))];
      _155_v4 = !((_157_v6)[_module.__default.safeIndex((_dafny.ZERO).minus(_152_v1), new BigNumber((_157_v6).length))]);
      process.stdout.write(_dafny.toString(_152_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v2).equals(_dafny.MultiSet.fromElements(new BigNumber(581), new BigNumber(581)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_154_v3, _dafny.Seq.of(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_155_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(581),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_157_v6, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v7).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v8)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_160_v9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_162_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState.f4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState.f4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState.f4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState.f4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f7)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_164_v12).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v13).equals(_dafny.MultiSet.fromElements(false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_166_v14, _dafny.Seq.of(_dafny.MultiSet.fromElements(false, false, false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("joueqk"),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_211_v51).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_298_v119).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_299_v120));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_300_v121)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_300_v121)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_300_v121)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_300_v121)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_300_v121)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_301_v122).dtor_cf29));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_301_v122).dtor_cf30));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v123).equals(_dafny.MultiSet.fromElements(_module.D9.create_DC20(false, new _dafny.CodePoint('j'.codePointAt(0))), _module.D9.create_DC20(false, new _dafny.CodePoint('j'.codePointAt(0))), _module.D9.create_DC20(false, new _dafny.CodePoint('j'.codePointAt(0))), _module.D9.create_DC20(false, new _dafny.CodePoint('j'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_303_i6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_321_v141).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_392_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.ZERO;
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
    static create_DC2(cf2) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC3(cf3, cf4) {
      let $dt = new D1(1);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(2);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC1() { return this.$tag === 2; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf2 === other.cf2;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf1 === other.cf1;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(false);
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
    static create_DC4(cf5) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf5 === other.cf5;
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf7, cf8, cf9, cf10, cf11) {
      let $dt = new D3(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC8(cf13, cf14, cf15) {
      let $dt = new D3(2);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC5(cf6) {
      let $dt = new D3(3);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8 && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf6 === other.cf6;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC6(_dafny.ZERO, false, false, _dafny.ZERO, false);
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
    static create_DC10(cf17, cf18, cf19) {
      let $dt = new D4(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC9(cf16) {
      let $dt = new D4(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(null, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC11(cf20) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC11" + "(" + this.cf20.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
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
    static create_DC13(cf22) {
      let $dt = new D6(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC12(cf21) {
      let $dt = new D6(1);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC13(false);
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
    static create_DC15(cf24, cf25) {
      let $dt = new D7(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC14(cf23) {
      let $dt = new D7(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC16(cf26) {
      let $dt = new D7(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24) && this.cf25 === other.cf25;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC15(_dafny.ZERO, false);
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
    static create_DC18() {
      let $dt = new D8(0);
      return $dt;
    }
    static create_DC17(cf27) {
      let $dt = new D8(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC18";
      } else if (this.$tag === 1) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC18();
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
    static create_DC20(cf29, cf30) {
      let $dt = new D9(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC19(cf28) {
      let $dt = new D9(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC21(cf31) {
      let $dt = new D9(2);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf29 === other.cf29 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC20(false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC23(cf33, cf34, cf35) {
      let $dt = new D10(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC22(cf32) {
      let $dt = new D10(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC22" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33) && this.cf34 === other.cf34 && this.cf35 === other.cf35;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf32 === other.cf32;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC23(_dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC25(cf37) {
      let $dt = new D11(0);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC26() {
      let $dt = new D11(1);
      return $dt;
    }
    static create_DC24(cf36) {
      let $dt = new D11(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC27(cf38) {
      let $dt = new D11(3);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get is_DC27() { return this.$tag === 3; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC26";
      } else if (this.$tag === 2) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf38) + ")";
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
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf36 === other.cf36;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC25(new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC29(cf40, cf41) {
      let $dt = new D12(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC28(cf39) {
      let $dt = new D12(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC30(cf42) {
      let $dt = new D12(2);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf40 === other.cf40 && this.cf41 === other.cf41;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC29([], false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf44) {
      let $dt = new D13(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC33(cf45, cf46, cf47, cf48) {
      let $dt = new D13(1);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC31(cf43) {
      let $dt = new D13(2);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC34(cf49) {
      let $dt = new D13(3);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get is_DC34() { return this.$tag === 3; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC32" + "(" + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC31" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf45 === other.cf45 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47) && this.cf48 === other.cf48;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC32(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC36(cf51, cf52, cf53, cf54, cf55) {
      let $dt = new D14(0);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC37(cf56) {
      let $dt = new D14(1);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC35(cf50) {
      let $dt = new D14(2);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52 && this.cf53 === other.cf53 && _dafny.areEqual(this.cf54, other.cf54) && this.cf55 === other.cf55;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf56 === other.cf56;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC36(_dafny.ZERO, false, false, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC39(cf58, cf59, cf60, cf61, cf62) {
      let $dt = new D15(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC38(cf57) {
      let $dt = new D15(1);
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC40(cf63) {
      let $dt = new D15(2);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC39" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC40" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60 && _dafny.areEqual(this.cf61, other.cf61) && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf57 === other.cf57;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf63, other.cf63);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC39(false, _dafny.ZERO, false, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC42(cf65, cf66) {
      let $dt = new D16(0);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC43(cf67, cf68, cf69) {
      let $dt = new D16(1);
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC41(cf64) {
      let $dt = new D16(2);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC41() { return this.$tag === 2; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC43" + "(" + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf64) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf65, other.cf65) && this.cf66 === other.cf66;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf67 === other.cf67 && _dafny.areEqual(this.cf68, other.cf68) && _dafny.areEqual(this.cf69, other.cf69);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC42(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC45(cf71, cf72) {
      let $dt = new D17(0);
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC44(cf70) {
      let $dt = new D17(1);
      $dt.cf70 = cf70;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf70() { return this.cf70; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC45" + "(" + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC44" + "(" + _dafny.toString(this.cf70) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf71, other.cf71) && this.cf72 === other.cf72;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf70, other.cf70);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC45(_dafny.Map.Empty, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
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
      this.f2 = _dafny.ZERO;
      this.f4 = [];
      this._f0 = false;
      this._f1 = false;
      this._f3 = [];
      this._f5 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f7 = [];
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
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
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f3() {
      let _this = this;
      return _this._f3;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f12 = false;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f12, f13) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("aqbqv"), _dafny.Seq.UnicodeFromString("x")), _dafny.Seq.UnicodeFromString("g"));
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

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f17 = _module.D3.Default();
      this._f16 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f16, f17) {
      let _this = this;
      (_this)._f16 = f16;
      (_this).f17 = f17;
      return;
    }
    fm3(globalState) {
      let _this = this;
      let _source9 = new BigNumber(310);
      let _398___mcc_h0 = _source9;
      let _399_cf0 = _398___mcc_h0;
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-795)),_dafny.Map.Empty.slice().updateUnsafe((_this).f16,_399_cf0));
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f16;
    };
    m5(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _400_v0;
      _400_v0 = _module.D1.create_DC3((_this).f16, _module.__default.fm1((_this).f16, true, globalState));
      let _source10 = function (_source11) {
        if (_source11.is_DC2) {
          let _401___mcc_h4 = (_source11).cf2;
          let _402_cf2 = _401___mcc_h4;
          return _module.D4.create_DC9(_dafny.MultiSet.FromArray(_dafny.Seq.of(_402_cf2)));
        } else if (_source11.is_DC3) {
          let _403___mcc_h5 = (_source11).cf3;
          let _404___mcc_h6 = (_source11).cf4;
          let _405_cf4 = _404___mcc_h6;
          let _406_cf3 = _403___mcc_h5;
          return _module.D4.create_DC9(_dafny.MultiSet.fromElements(_406_cf3, (_this).f16));
        } else {
          let _407___mcc_h7 = (_source11).cf1;
          let _408_cf1 = _407___mcc_h7;
          return _module.D4.create_DC9(_dafny.MultiSet.fromElements(true));
        }
      }(_400_v0);
      if (_source10.is_DC10) {
        let _409___mcc_h0 = (_source10).cf17;
        let _410___mcc_h1 = (_source10).cf18;
        let _411___mcc_h2 = (_source10).cf19;
        let _412_cf19 = _411___mcc_h2;
        let _413_cf18 = _410___mcc_h1;
        let _414_cf17 = _409___mcc_h0;
        let _415_v1;
        _415_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(((_this).f16) ? (_412_cf19) : (_413_cf18)));
        _415_v1 = (_415_v1).update((_this).f16, _412_cf19);
        let _416_v2;
        _416_v2 = _module.D3.create_DC7((p0).minus(_413_cf18));
        let _417_v3;
        _417_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_this).f16);
        _416_v2 = (((p0).isLessThanOrEqualTo(new BigNumber((_417_v3).length))) ? (_416_v2) : ((((_this).f16) ? (_416_v2) : (_416_v2))));
        let _418_v4;
        _418_v4 = _dafny.Seq.of(new BigNumber(980));
        let _419_v5;
        _419_v5 = _dafny.Seq.of(new BigNumber(-191), (_418_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_418_v4).length))]);
        _419_v5 = _dafny.Seq.update(_418_v4, _module.__default.safeIndex(p0, new BigNumber((_418_v4).length)), p0);
        let _420_v6;
        _420_v6 = _dafny.Seq.UnicodeFromString("raf");
        (globalState).f2 = new BigNumber((_module.__default.fm18((((_this).f16) ? (new BigNumber((_420_v6).length)) : (_413_cf18)), (_this).f16, _412_cf19, globalState)).length);
      } else {
        let _421___mcc_h3 = (_source10).cf16;
        let _422_cf16 = _421___mcc_h3;
        let _423_v7;
        _423_v7 = true;
        _423_v7 = (_this).f16;
        let _424_v8;
        _424_v8 = _dafny.Seq.UnicodeFromString("fg");
        let _425_v9;
        _425_v9 = new _dafny.CodePoint('t'.codePointAt(0));
        let _426_v10;
        _426_v10 = _dafny.MultiSet.fromElements(_425_v9, new _dafny.CodePoint('h'.codePointAt(0)), _425_v9);
        let _427_v11;
        _427_v11 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0));
        let _428_v12;
        _428_v12 = _dafny.Set.fromElements(new BigNumber((_427_v11).cardinality()));
        let _429_v14;
        _429_v14 = _dafny.Seq.of(_424_v8);
        let _430_v15;
        _430_v15 = _dafny.Seq.of(_module.__default.fm0(p0, p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(42)), ((_431_cf16) => function (_432_i0) {
          return _431_cf16;
        })(_422_cf16)), function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of (_dafny.Seq.update(_429_v14, _module.__default.safeIndex(p0, new BigNumber((_429_v14).length)), _424_v8)).Elements) {
            let _433_v13 = _compr_14;
            if (_dafny.Seq.contains(_dafny.Seq.update(_429_v14, _module.__default.safeIndex(p0, new BigNumber((_429_v14).length)), _424_v8), _433_v13)) {
              _coll14.push([_433_v13,(_this).f16]);
            }
          }
          return _coll14;
        }(), globalState), _423_v7);
        let _434_v16;
        let _nw67 = Array((new BigNumber(22)).toNumber());
        _nw67[(_dafny.ZERO).toNumber()] = (_this).f16;
        _nw67[(_dafny.ONE).toNumber()] = ((!((_this).f16)) ? ((_this).f16) : ((_this).f16));
        _nw67[(new BigNumber(2)).toNumber()] = (_this).f16;
        _nw67[(new BigNumber(3)).toNumber()] = (_this).f16;
        _nw67[(new BigNumber(4)).toNumber()] = _dafny.areEqual(_424_v8, _dafny.Seq.UnicodeFromString("qmuiq"));
        _nw67[(new BigNumber(5)).toNumber()] = (_426_v10).IsDisjointFrom(_426_v10);
        _nw67[(new BigNumber(6)).toNumber()] = (_this).f16;
        _nw67[(new BigNumber(7)).toNumber()] = !(_423_v7);
        _nw67[(new BigNumber(8)).toNumber()] = (_428_v12).IsSubsetOf(_428_v12);
        _nw67[(new BigNumber(9)).toNumber()] = (_423_v7) === (_423_v7);
        _nw67[(new BigNumber(10)).toNumber()] = _423_v7;
        _nw67[(new BigNumber(11)).toNumber()] = _423_v7;
        _nw67[(new BigNumber(12)).toNumber()] = _423_v7;
        _nw67[(new BigNumber(13)).toNumber()] = _423_v7;
        _nw67[(new BigNumber(14)).toNumber()] = !(!(((_this).f16) && (_423_v7)));
        _nw67[(new BigNumber(15)).toNumber()] = _423_v7;
        _nw67[(new BigNumber(16)).toNumber()] = _423_v7;
        _nw67[(new BigNumber(17)).toNumber()] = (_this).f16;
        _nw67[(new BigNumber(18)).toNumber()] = _423_v7;
        _nw67[(new BigNumber(19)).toNumber()] = (_this).f16;
        _nw67[(new BigNumber(20)).toNumber()] = _dafny.areEqual(_430_v15, _dafny.Seq.update(_430_v15, _module.__default.safeIndex(p0, new BigNumber((_430_v15).length)), (_this).f16));
        _nw67[(new BigNumber(21)).toNumber()] = _423_v7;
        _434_v16 = _nw67;
        let _index82 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_434_v16).length));
        (_434_v16)[_index82] = (_this).f16;
        let _index83 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_434_v16).length));
        (_434_v16)[_index83] = true;
        let _435_v17;
        _435_v17 = _dafny.Seq.of(_422_cf16);
        let _436_v18;
        _436_v18 = _dafny.Map.Empty.slice().updateUnsafe(_424_v8,_423_v7);
        let _437_v19;
        _437_v19 = _dafny.Map.Empty.slice().updateUnsafe(_425_v9,_436_v18);
        let _438_v20;
        _438_v20 = _dafny.Map.Empty.slice().updateUnsafe((_434_v16)[_module.__default.safeIndex(new BigNumber(295), new BigNumber((_434_v16).length))],_module.__default.fm0(p0, new BigNumber((_module.__default.fm18(p0, (_this).f16, p0, globalState)).length), _435_v17, (((_437_v19).contains(new _dafny.CodePoint('b'.codePointAt(0)))) ? ((_437_v19).get(new _dafny.CodePoint('b'.codePointAt(0)))) : (_436_v18)), globalState));
        let _439_v21;
        _439_v21 = _dafny.Set.fromElements((_434_v16)[_module.__default.safeIndex(new BigNumber(295), new BigNumber((_434_v16).length))]);
        let _440_v22;
        _440_v22 = _dafny.Seq.of(_dafny.Set.fromElements((_434_v16)[_module.__default.safeIndex(new BigNumber(295), new BigNumber((_434_v16).length))], false), _439_v21, _dafny.Set.fromElements((_434_v16)[_module.__default.safeIndex(new BigNumber(295), new BigNumber((_434_v16).length))], !((_this).f16)));
        _438_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber(((_440_v22)[_module.__default.safeIndex(p0, new BigNumber((_440_v22).length))]).length), p0, _435_v17, _436_v18, globalState),(_this).f16);
        let _441_v23;
        let _nw68 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
        _441_v23 = _nw68;
        let _index84 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_441_v23).length));
        (_441_v23)[_index84] = _430_v15;
        let _index85 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_441_v23).length));
        (_441_v23)[_index85] = _430_v15;
      }
      let _442_i1;
      _442_i1 = _dafny.ZERO;
      L2: {
        while ((p0).isLessThanOrEqualTo(new BigNumber(-271))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_442_i1)) {
              break L2;
            }
            _442_i1 = (_442_i1).plus(_dafny.ONE);
            let _443_v24;
            _443_v24 = false;
            _443_v24 = (_this).f16;
            r1 = p0;
            let _444_v25;
            _444_v25 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_module.__default.fm19(p0, globalState)));
            let _445_v26;
            _445_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
            let _446_v27;
            let _nw69 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            _446_v27 = _nw69;
            let _447_v28;
            _447_v28 = _dafny.Map.Empty.slice().updateUnsafe(!(_443_v24),_446_v27);
            let _448_v29;
            let _nw70 = Array((new BigNumber(11)).toNumber());
            _nw70[(_dafny.ZERO).toNumber()] = (_this).f16;
            _nw70[(_dafny.ONE).toNumber()] = _443_v24;
            _nw70[(new BigNumber(2)).toNumber()] = true;
            _nw70[(new BigNumber(3)).toNumber()] = (_this).f16;
            _nw70[(new BigNumber(4)).toNumber()] = _443_v24;
            _nw70[(new BigNumber(5)).toNumber()] = _443_v24;
            _nw70[(new BigNumber(6)).toNumber()] = (_this).f16;
            _nw70[(new BigNumber(7)).toNumber()] = (_this).f16;
            _nw70[(new BigNumber(8)).toNumber()] = (_this).f16;
            _nw70[(new BigNumber(9)).toNumber()] = true;
            _nw70[(new BigNumber(10)).toNumber()] = (_this).f16;
            _448_v29 = _nw70;
            let _449_v30;
            _449_v30 = _dafny.Set.fromElements(_448_v29, _448_v29);
            let _450_v31;
            _450_v31 = _module.D7.create_DC14(_449_v30);
            let _451_v32;
            let _nw71 = new _module.C0();
            _nw71.__ctor(_443_v24, p0);
            _451_v32 = _nw71;
            let _452_v33;
            _452_v33 = _dafny.Map.Empty.slice().updateUnsafe(_451_v32,_module.__default.fm19(p0, globalState));
            let _453_v34;
            _453_v34 = _dafny.Seq.of(_443_v24, (_451_v32).f12, (_451_v32).f12, (_this).f16, _443_v24);
            let _454_v35;
            _454_v35 = _dafny.Seq.UnicodeFromString("ie");
            let _455_v36;
            _455_v36 = _dafny.Map.Empty.slice().updateUnsafe((_451_v32).f12,_454_v35);
            let _456_v37;
            _456_v37 = _dafny.Map.Empty.slice().updateUnsafe(_443_v24,new BigNumber((_455_v36).length));
            let _457_v38;
            let _nw72 = Array((new BigNumber(24)).toNumber());
            _nw72[(_dafny.ZERO).toNumber()] = p0;
            _nw72[(_dafny.ONE).toNumber()] = (((_445_v26).contains(new BigNumber(452))) ? ((_445_v26).get(new BigNumber(452))) : (new BigNumber((_447_v28).length)));
            _nw72[(new BigNumber(2)).toNumber()] = p0;
            _nw72[(new BigNumber(3)).toNumber()] = p0;
            _nw72[(new BigNumber(4)).toNumber()] = p0;
            _nw72[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p0);
            _nw72[(new BigNumber(6)).toNumber()] = p0;
            _nw72[(new BigNumber(7)).toNumber()] = p0;
            _nw72[(new BigNumber(8)).toNumber()] = new BigNumber(689);
            _nw72[(new BigNumber(9)).toNumber()] = new BigNumber(((_450_v31).dtor_cf23).length);
            _nw72[(new BigNumber(10)).toNumber()] = p0;
            _nw72[(new BigNumber(11)).toNumber()] = p0;
            _nw72[(new BigNumber(12)).toNumber()] = p0;
            _nw72[(new BigNumber(13)).toNumber()] = new BigNumber((_452_v33).length);
            _nw72[(new BigNumber(14)).toNumber()] = p0;
            _nw72[(new BigNumber(15)).toNumber()] = new BigNumber((_module.__default.fm20(globalState)).length);
            _nw72[(new BigNumber(16)).toNumber()] = new BigNumber(499);
            _nw72[(new BigNumber(17)).toNumber()] = p0;
            _nw72[(new BigNumber(18)).toNumber()] = (_451_v32).f13;
            _nw72[(new BigNumber(19)).toNumber()] = p0;
            _nw72[(new BigNumber(20)).toNumber()] = new BigNumber((_453_v34).length);
            _nw72[(new BigNumber(21)).toNumber()] = (((_456_v37).contains(true)) ? ((_456_v37).get(true)) : (_module.__default.fm19(p0, globalState)));
            _nw72[(new BigNumber(22)).toNumber()] = (_451_v32).f13;
            _nw72[(new BigNumber(23)).toNumber()] = p0;
            _457_v38 = _nw72;
            let _458_v39;
            _458_v39 = _457_v38;
            let _459_v40;
            _459_v40 = _dafny.Seq.of(new BigNumber((_454_v35).length));
            let _rhs101 = _443_v24;
            let _rhs102 = false;
            let _rhs103 = (_dafny.MultiSet.FromArray(_459_v40)).Difference(((_444_v25).update(new BigNumber((_454_v35).length), _module.__default.abs(p0))).Intersect(_444_v25));
            let _rhs104 = _446_v27;
            _443_v24 = _rhs101;
            _443_v24 = _rhs102;
            _444_v25 = _rhs103;
            _458_v39 = _rhs104;
            (globalState).f2 = p0;
          }
        }
      }
      let _460_i2;
      _460_i2 = _dafny.ZERO;
      L3: {
        while (!(((_this).f16) && (!((_this).f16)))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_460_i2)) {
              break L3;
            }
            _460_i2 = (_460_i2).plus(_dafny.ONE);
            (globalState).f2 = p0;
            let _461_v41;
            let _nw73 = new _module.C0();
            _nw73.__ctor((_this).f16, p0);
            _461_v41 = _nw73;
            if ((_this).f16) {
              let _462_v42;
              _462_v42 = false;
              let _463_v43;
              _463_v43 = _dafny.MultiSet.fromElements((_461_v41).f13);
              let _464_v44;
              _464_v44 = new _dafny.CodePoint('o'.codePointAt(0));
              _462_v42 = (_this).fm4(((_461_v41).f13).multipliedBy(new BigNumber((_463_v43).cardinality())), (p0).isLessThanOrEqualTo((_dafny.ZERO).minus((_461_v41).f13)), _464_v44, globalState);
              _462_v42 = !((_461_v41).f12);
              let _465_v45;
              _465_v45 = _dafny.Seq.of((_461_v41).f13, (_461_v41).f13);
              let _466_v46;
              _466_v46 = _dafny.Seq.of((_465_v45)[_module.__default.safeIndex(p0, new BigNumber((_465_v45).length))], _module.__default.fm19((_461_v41).f13, globalState), p0, _module.__default.fm19((_461_v41).f13, globalState));
              let _467_v47;
              let _init4 = ((_468_v41) => function (_469_i3) {
                return (_469_i3).plus((_468_v41).f13);
              })(_461_v41);
              let _nw74 = Array((new BigNumber(12)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw74.length); _i0_4++) {
                _nw74[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _467_v47 = _nw74;
              let _index86 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_467_v47).length));
              (_467_v47)[_index86] = (_461_v41).f13;
              let _470_v48;
              let _init5 = ((_471_v42, _472_v41) => function (_473_i4) {
                return (_dafny.MultiSet.fromElements(_471_v42)).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of((_472_v41).f12)));
              })(_462_v42, _461_v41);
              let _nw75 = Array((new BigNumber(12)).toNumber());
              for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw75.length); _i0_5++) {
                _nw75[_i0_5] = _init5(new BigNumber(_i0_5));
              }
              _470_v48 = _nw75;
              let _474_v49;
              _474_v49 = _dafny.MultiSet.fromElements((_461_v41).f12);
              let _index87 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_470_v48).length));
              (_470_v48)[_index87] = (_474_v49).Union(_module.__default.fm21((_461_v41).f13, false, (_461_v41).f13, globalState));
              let _475_v50;
              _475_v50 = _dafny.Set.fromElements(new BigNumber(53));
              let _index88 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_467_v47).length));
              let _index89 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_470_v48).length));
              let _rhs105 = !((_475_v50).IsDisjointFrom(_475_v50));
              let _rhs106 = _dafny.Seq.Concat(_dafny.Seq.Concat(_465_v45, _465_v45), _466_v46);
              let _rhs107 = _module.__default.fm19((new BigNumber((_463_v43).cardinality())).minus(new BigNumber((_474_v49).cardinality())), globalState);
              let _rhs108 = _474_v49;
              let _lhs65 = _467_v47;
              let _lhs66 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_467_v47).length));
              let _lhs67 = _470_v48;
              let _lhs68 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_470_v48).length));
              _462_v42 = _rhs105;
              _465_v45 = _rhs106;
              _lhs65[_lhs66] = _rhs107;
              _lhs67[_lhs68] = _rhs108;
              let _476_v51;
              let _nw76 = new _module.C0();
              _nw76.__ctor((_461_v41).f12, (_dafny.ZERO).minus((_467_v47)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_467_v47).length))]));
              _476_v51 = _nw76;
              let _477_v52;
              _477_v52 = _dafny.Map.Empty.slice().updateUnsafe(_462_v42,(_this).f16);
              let _478_v53;
              let _nw77 = Array((new BigNumber(29)).toNumber());
              _nw77[(_dafny.ZERO).toNumber()] = _module.__default.fm22(globalState);
              _nw77[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_461_v41).f12);
              _nw77[(new BigNumber(2)).toNumber()] = (_module.__default.fm22(globalState)).Merge(_477_v52);
              _nw77[(new BigNumber(3)).toNumber()] = (_module.__default.fm22(globalState)).Merge(_477_v52);
              _nw77[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_461_v41).f12);
              _nw77[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_476_v51).f12,(_this).f16);
              _nw77[(new BigNumber(6)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(7)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(8)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(9)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(10)).toNumber()] = (_477_v52).Merge(_477_v52);
              _nw77[(new BigNumber(11)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(12)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(13)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(14)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(15)).toNumber()] = (_477_v52).Merge(_477_v52);
              _nw77[(new BigNumber(16)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(17)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(18)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(19)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(20)).toNumber()] = (((_476_v51).f12) ? (_477_v52) : (_477_v52));
              _nw77[(new BigNumber(21)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(22)).toNumber()] = (_477_v52).Merge(_477_v52);
              _nw77[(new BigNumber(23)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_461_v41).f12,(_461_v41).f12)).Merge(_477_v52);
              _nw77[(new BigNumber(24)).toNumber()] = (_477_v52).Merge(_477_v52);
              _nw77[(new BigNumber(25)).toNumber()] = (_477_v52).update(_462_v42, (_this).f16);
              _nw77[(new BigNumber(26)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(27)).toNumber()] = _477_v52;
              _nw77[(new BigNumber(28)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_461_v41).f12,(_476_v51).f12);
              _478_v53 = _nw77;
              let _index90 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_478_v53).length));
              (_478_v53)[_index90] = (_477_v52).Merge(_477_v52);
              let _index91 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_478_v53).length));
              (_478_v53)[_index91] = _477_v52;
            } else {
              r1 = p0;
              let _479_v54;
              let _nw78 = Array((new BigNumber(25)).toNumber()).fill(false);
              _479_v54 = _nw78;
              _479_v54 = _479_v54;
              _479_v54 = _479_v54;
              let _480_v55;
              _480_v55 = true;
              let _481_v56;
              _481_v56 = new _dafny.CodePoint('m'.codePointAt(0));
              let _482_v57;
              _482_v57 = _dafny.Seq.UnicodeFromString("m");
              _480_v55 = _dafny.Seq.contains(_482_v57, _481_v56);
              let _483_v58;
              let _nw79 = new _module.C0();
              _nw79.__ctor(!((_461_v41).f13).isEqualTo(new BigNumber(734)), (p0).plus(p0));
              _483_v58 = _nw79;
            }
            if ((_461_v41).f12) {
              let _484_v59;
              _484_v59 = true;
              let _485_v60;
              _485_v60 = new _dafny.CodePoint('v'.codePointAt(0));
              let _486_v61;
              _486_v61 = _dafny.Seq.of(_485_v60, _485_v60, _485_v60, _485_v60, _485_v60);
              _484_v59 = _dafny.Seq.IsPrefixOf(_486_v61, _486_v61);
              _485_v60 = _485_v60;
              let _487_v62;
              let _init6 = ((_488_v61, _489_v41) => function (_490_i5) {
                return _module.D7.create_DC15(new BigNumber((_488_v61).length), (_489_v41).f12);
              })(_486_v61, _461_v41);
              let _nw80 = Array((new BigNumber(7)).toNumber());
              for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw80.length); _i0_6++) {
                _nw80[_i0_6] = _init6(new BigNumber(_i0_6));
              }
              _487_v62 = _nw80;
              let _491_v63;
              _491_v63 = _module.D7.create_DC15((_461_v41).f13, (_this).f16);
              let _index92 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_487_v62).length));
              (_487_v62)[_index92] = _491_v63;
              let _index93 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_487_v62).length));
              (_487_v62)[_index93] = _491_v63;
              _484_v59 = (_461_v41).f12;
              let _rhs109 = (p0).multipliedBy((_461_v41).f13);
              let _rhs110 = (p0).isLessThanOrEqualTo(_module.__default.fm19((_461_v41).f13, globalState));
              let _rhs111 = (_461_v41).f13;
              let _rhs112 = (_module.__default.safeModuloInt((_461_v41).f13, (_461_v41).f13)).plus((_461_v41).f13);
              let _lhs69 = globalState;
              let _lhs70 = globalState;
              _lhs69.f2 = _rhs109;
              _484_v59 = _rhs110;
              _lhs70.f2 = _rhs111;
              r1 = _rhs112;
            } else {
              let _492_v64;
              _492_v64 = true;
              _492_v64 = _492_v64;
              let _493_v65;
              _493_v65 = _dafny.Map.Empty.slice().updateUnsafe((_461_v41).f12,(_461_v41).f12);
              let _494_v66;
              _494_v66 = _dafny.Set.fromElements(p0, (_461_v41).f13);
              let _495_v67;
              _495_v67 = _module.D3.create_DC8((_this).f16, new BigNumber((_494_v66).length), p0);
              let _496_v68;
              _496_v68 = _dafny.MultiSet.fromElements((_495_v67).dtor_cf15);
              let _497_v69;
              _497_v69 = _dafny.Map.Empty.slice().updateUnsafe((_461_v41).f13,(_this).f16);
              let _498_v70;
              _498_v70 = _module.D3.create_DC8((_461_v41).f12, (_461_v41).f13, (((_496_v68).contains(p0)) ? ((_496_v68).get(p0)) : (new BigNumber((_497_v69).length))));
              let _499_v71;
              _499_v71 = _dafny.MultiSet.fromElements((new BigNumber((_493_v65).length)).multipliedBy((_498_v70).dtor_cf14), (_461_v41).f13);
              let _rhs113 = (_499_v71).Union(_dafny.MultiSet.fromElements((_461_v41).f13));
              let _rhs114 = p0;
              let _rhs115 = _module.__default.fm19((_461_v41).f13, globalState);
              let _rhs116 = (_this).f16;
              let _lhs71 = globalState;
              _499_v71 = _rhs113;
              _lhs71.f2 = _rhs114;
              r0 = _rhs115;
              _492_v64 = _rhs116;
              let _500_v72;
              _500_v72 = (_461_v41).f13;
              let _501_v73;
              _501_v73 = _dafny.MultiSet.fromElements(_module.D1.create_DC3((_461_v41).f12, _500_v72), _400_v0);
              r0 = ((((_501_v73).contains(_400_v0)) ? ((_501_v73).get(_400_v0)) : (_module.__default.fm19(p0, globalState)))).minus((_dafny.ZERO).minus((_461_v41).f13));
              let _502_v74;
              let _nw81 = Array((new BigNumber(10)).toNumber()).fill(false);
              _502_v74 = _nw81;
              let _index94 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_502_v74).length));
              (_502_v74)[_index94] = (_461_v41).f12;
              let _503_v75;
              _503_v75 = _dafny.Seq.UnicodeFromString("eotb");
              let _504_v76;
              _504_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_503_v75).length),_495_v67);
              let _505_v78;
              _505_v78 = new _dafny.CodePoint('d'.codePointAt(0));
              let _506_v79;
              _506_v79 = _dafny.Map.Empty.slice().updateUnsafe((_461_v41).f13,_505_v78);
              let _index95 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_502_v74).length));
              let _rhs117 = (p0).isLessThan(p0);
              let _rhs118 = (((_492_v64) ? (_504_v76) : (_504_v76))).Merge((_504_v76).Merge(function () {
                let _coll15 = new _dafny.Map();
                for (const _compr_15 of (_506_v79).Keys.Elements) {
                  let _507_v77 = _compr_15;
                  if ((_506_v79).contains(_507_v77)) {
                    _coll15.push([(_507_v77).plus(new BigNumber(748)),_495_v67]);
                  }
                }
                return _coll15;
              }()));
              let _lhs72 = _502_v74;
              let _lhs73 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_502_v74).length));
              _lhs72[_lhs73] = _rhs117;
              _504_v76 = _rhs118;
              let _508_v80;
              let _nw82 = new _module.C0();
              _nw82.__ctor((_461_v41).f12, (_461_v41).f13);
              _508_v80 = _nw82;
            }
          }
        }
      }
      let _509_v81;
      _509_v81 = true;
      let _rhs119 = (_this).f16;
      let _rhs120 = p0;
      let _lhs74 = globalState;
      _509_v81 = _rhs119;
      _lhs74.f2 = _rhs120;
      let _hi2 = p0;
      for (let _510_i6 = p0; _510_i6.isLessThan(_hi2); _510_i6 = _510_i6.plus(_dafny.ONE)) {
        let _511_v82;
        let _512_v83;
        let _513_v84;
        let _514_v85;
        let _out42;
        let _out43;
        let _out44;
        let _out45;
        let _outcollector13 = _module.__default.m0(globalState);
        _out42 = _outcollector13[0];
        _out43 = _outcollector13[1];
        _out44 = _outcollector13[2];
        _out45 = _outcollector13[3];
        _511_v82 = _out42;
        _512_v83 = _out43;
        _513_v84 = _out44;
        _514_v85 = _out45;
        let _515_v86;
        let _nw83 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _515_v86 = _nw83;
        let _index96 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_515_v86).length));
        (_515_v86)[_index96] = _module.__default.fm19((_dafny.ZERO).minus(p0), globalState);
        let _516_v87;
        let _nw84 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _516_v87 = _nw84;
        let _517_v88;
        _517_v88 = _dafny.Seq.UnicodeFromString("jasbi");
        let _518_v89;
        _518_v89 = _517_v88;
        let _index97 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_516_v87).length));
        (_516_v87)[_index97] = _518_v89;
        let _519_v90;
        _519_v90 = _dafny.Map.Empty.slice().updateUnsafe(_510_i6,(_this).f16);
        let _520_v91;
        _520_v91 = _dafny.Seq.of(_512_v83);
        let _521_v92;
        _521_v92 = _dafny.Map.Empty.slice().updateUnsafe(_510_i6,p0);
        let _522_v93;
        _522_v93 = _dafny.MultiSet.fromElements((((_521_v92).contains(_510_i6)) ? ((_521_v92).get(_510_i6)) : (new BigNumber(267))));
        let _523_v94;
        _523_v94 = _dafny.MultiSet.fromElements(_509_v81);
        let _524_v95;
        _524_v95 = _dafny.Seq.of(_523_v94);
        let _525_v96;
        _525_v96 = _dafny.Map.Empty.slice().updateUnsafe(_517_v88,_509_v81);
        let _526_v97;
        _526_v97 = _module.D3.create_DC6(_511_v82, (((_519_v90).contains(new BigNumber((_520_v91).length))) ? ((_519_v90).get(new BigNumber((_520_v91).length))) : (false)), !(_module.__default.fm0(new BigNumber((_522_v93).cardinality()), new BigNumber((_520_v91).length), _524_v95, _525_v96, globalState)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_511_v82,(_this).f16)).length), true);
        let _527_v98;
        _527_v98 = _dafny.Seq.of(_511_v82, _510_i6);
        let _index98 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_515_v86).length));
        let _index99 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_516_v87).length));
        let _rhs121 = (_526_v97).dtor_cf7;
        let _rhs122 = _510_i6;
        let _rhs123 = new BigNumber((_dafny.Seq.Concat(_527_v98, ((_513_v84) ? (_527_v98) : (_527_v98)))).length);
        let _rhs124 = p0;
        let _rhs125 = _518_v89;
        let _lhs75 = _515_v86;
        let _lhs76 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_515_v86).length));
        let _lhs77 = globalState;
        let _lhs78 = globalState;
        let _lhs79 = _516_v87;
        let _lhs80 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_516_v87).length));
        _511_v82 = _rhs121;
        _lhs75[_lhs76] = _rhs122;
        _lhs77.f2 = _rhs123;
        _lhs78.f2 = _rhs124;
        _lhs79[_lhs80] = _rhs125;
        _513_v84 = (_module.__default.safeDivisionInt(_511_v82, new BigNumber(965))).isLessThanOrEqualTo(p0);
        let _528_v99;
        _528_v99 = new _dafny.CodePoint('e'.codePointAt(0));
        let _529_v100;
        _529_v100 = _dafny.Map.Empty.slice().updateUnsafe(!(_509_v81),false);
        _528_v99 = (((((_529_v100).contains(_513_v84)) ? ((_529_v100).get(_513_v84)) : (true))) ? (_528_v99) : (new _dafny.CodePoint('m'.codePointAt(0))));
      }
      let _530_v101;
      _530_v101 = _dafny.Seq.UnicodeFromString("tri");
      r0 = new BigNumber((_530_v101).length);
      r0 = p0;
      r1 = p0;
      return [r0, r1];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this.f14 = false;
      this._f15 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f14, f15) {
      let _this = this;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(531)).plus(new BigNumber((_dafny.Seq.of((_this).f15, _this.f14, false)).length)),(_dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), function (_531_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber(730))));
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f14;
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_this.f14)).cardinality()), new BigNumber(73))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), function (_532_i0) {
        return new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_this.f14)).length)))).length);
      })).length)));
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return ((new BigNumber((_dafny.Seq.UnicodeFromString("lgislaabn")).length)).minus(new BigNumber(169))).isLessThan(new BigNumber(406));
    };
    m3(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let _533_v0;
      _533_v0 = new BigNumber(55);
      let _534_v1;
      let _nw85 = new _module.C0();
      _nw85.__ctor((_this).f15, _533_v0);
      _534_v1 = _nw85;
      let _535_v2;
      _535_v2 = _dafny.Seq.of(_534_v1, _534_v1);
      let _536_v3;
      _536_v3 = _dafny.MultiSet.fromElements(!((_this).f15));
      let _537_v4;
      _537_v4 = _module.D3.create_DC7(_533_v0);
      let _538_v5;
      _538_v5 = _dafny.Map.Empty.slice().updateUnsafe(true,_537_v4);
      let _539_v6;
      _539_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_533_v0);
      let _540_v7;
      _540_v7 = _dafny.Seq.UnicodeFromString("vedhrd");
      let _541_v8;
      _541_v8 = _dafny.Map.Empty.slice().updateUnsafe((_534_v1).f13,new BigNumber((_539_v6).length));
      let _542_v10;
      _542_v10 = _dafny.Set.fromElements((_534_v1).f13, (_534_v1).f13);
      let _543_v11;
      _543_v11 = new _dafny.CodePoint('i'.codePointAt(0));
      let _544_v12;
      _544_v12 = _dafny.Seq.of(new BigNumber((_542_v10).length), (_534_v1).f13, new BigNumber((_dafny.MultiSet.fromElements(_543_v11, _543_v11)).cardinality()), _533_v0);
      let _545_v13;
      let _nw86 = Array((new BigNumber(26)).toNumber());
      _nw86[(_dafny.ZERO).toNumber()] = (_533_v0).minus(_533_v0);
      _nw86[(_dafny.ONE).toNumber()] = _533_v0;
      _nw86[(new BigNumber(2)).toNumber()] = _533_v0;
      _nw86[(new BigNumber(3)).toNumber()] = new BigNumber(-457);
      _nw86[(new BigNumber(4)).toNumber()] = _533_v0;
      _nw86[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_533_v0);
      _nw86[(new BigNumber(6)).toNumber()] = _533_v0;
      _nw86[(new BigNumber(7)).toNumber()] = (_533_v0).minus(new BigNumber(452));
      _nw86[(new BigNumber(8)).toNumber()] = _533_v0;
      _nw86[(new BigNumber(9)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_535_v2)[_module.__default.safeIndex((_534_v1).f13, new BigNumber((_535_v2).length))],(_534_v1).f13)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_534_v1,_533_v0))).length);
      _nw86[(new BigNumber(10)).toNumber()] = new BigNumber(((_536_v3).update((_534_v1).f12, _module.__default.abs(_533_v0))).cardinality());
      _nw86[(new BigNumber(11)).toNumber()] = (_534_v1).f13;
      _nw86[(new BigNumber(12)).toNumber()] = new BigNumber((_module.__default.fm14(_538_v5, _this.f14, _539_v6, (_this).f15, globalState)).cardinality());
      _nw86[(new BigNumber(13)).toNumber()] = _533_v0;
      _nw86[(new BigNumber(14)).toNumber()] = (new BigNumber((_dafny.MultiSet.fromElements((_this).f15, (_this).f15)).cardinality())).plus(new BigNumber((_540_v7).length));
      _nw86[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_533_v0);
      _nw86[(new BigNumber(16)).toNumber()] = (_534_v1).f13;
      _nw86[(new BigNumber(17)).toNumber()] = new BigNumber((((_541_v8).update((_534_v1).f13, (_534_v1).f13)).update(new BigNumber((function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.Map.Empty.slice().updateUnsafe(p0,(_534_v1).f13)).Keys.Elements) {
          let _546_v9 = _compr_16;
          if ((_dafny.Map.Empty.slice().updateUnsafe(p0,(_534_v1).f13)).contains(_546_v9)) {
            _coll16.push([_546_v9,false]);
          }
        }
        return _coll16;
      }()).length), new BigNumber(-304))).length);
      _nw86[(new BigNumber(18)).toNumber()] = (_534_v1).f13;
      _nw86[(new BigNumber(19)).toNumber()] = (_534_v1).f13;
      _nw86[(new BigNumber(20)).toNumber()] = ((_534_v1).f13).minus(new BigNumber((_544_v12).length));
      _nw86[(new BigNumber(21)).toNumber()] = new BigNumber(-752);
      _nw86[(new BigNumber(22)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(117), (_534_v1).f13);
      _nw86[(new BigNumber(23)).toNumber()] = new BigNumber(((((_this).f15) ? (_536_v3) : (_536_v3))).cardinality());
      _nw86[(new BigNumber(24)).toNumber()] = _533_v0;
      _nw86[(new BigNumber(25)).toNumber()] = ((_this.f14) ? (new BigNumber(293)) : (_533_v0));
      _545_v13 = _nw86;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_545_v13).length))) {
        let _547_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_547_i0)) && ((_547_i0).isLessThan(new BigNumber((_545_v13).length))))) {
          (_545_v13)[(_547_i0)] = (_547_i0).minus((_dafny.ZERO).minus(_533_v0));
        }
      }
      let _548_v14;
      let _nw87 = Array((new BigNumber(13)).toNumber());
      _548_v14 = _nw87;
      _548_v14 = _548_v14;
      let _549_v15;
      let _nw88 = new _module.C0();
      _nw88.__ctor(true, _module.__default.safeDivisionInt(_533_v0, _533_v0));
      _549_v15 = _nw88;
      let _550_v16;
      _550_v16 = _dafny.Map.Empty.slice().updateUnsafe((_549_v15).f13,false);
      let _551_v17;
      _551_v17 = _dafny.Seq.of(_550_v16, (_550_v16).Merge(_550_v16), _550_v16);
      _551_v17 = _551_v17;
      let _552_v18;
      _552_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_module.D3.create_DC8(_this.f14, new BigNumber((_dafny.Seq.update(_544_v12, _module.__default.safeIndex((_549_v15).f13, new BigNumber((_544_v12).length)), (_534_v1).f13)).length), _533_v0));
      let _553_v19;
      _553_v19 = _module.D3.create_DC8((_this).fm4((_549_v15).f13, (_549_v15).f12, _543_v11, globalState), _533_v0, (_549_v15).f13);
      let _554_v20;
      _554_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,(_549_v15).f12);
      let _555_v21;
      _555_v21 = _dafny.Set.fromElements(true, (_549_v15).f12);
      let _556_v22;
      let _nw89 = Array((new BigNumber(25)).toNumber());
      _nw89[(_dafny.ZERO).toNumber()] = (_552_v18).Merge(_552_v18);
      _nw89[(_dafny.ONE).toNumber()] = (_552_v18).Merge(_552_v18);
      _nw89[(new BigNumber(2)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(3)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(4)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(5)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_this.f14,_553_v19)).Merge(_552_v18);
      _nw89[(new BigNumber(6)).toNumber()] = ((_552_v18).update(!((_534_v1).f12), _553_v19)).Merge(_552_v18);
      _nw89[(new BigNumber(7)).toNumber()] = (_module.__default.fm15(globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_534_v1).f12,_module.D3.create_DC8(_this.f14, (_549_v15).f13, new BigNumber(-421))));
      _nw89[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_549_v15).f12,_module.__default.fm16(globalState));
      _nw89[(new BigNumber(9)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(10)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(11)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(12)).toNumber()] = (_552_v18).Merge((_552_v18).update(true, _module.D3.create_DC8((_549_v15).f12, _533_v0, _533_v0)));
      _nw89[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_549_v15).f12,_module.D3.create_DC8(false, _module.__default.fm17((((_554_v20).contains(_this.f14)) ? ((_554_v20).get(_this.f14)) : (_this.f14)), _533_v0, (_549_v15).f12, (_534_v1).f13, globalState), new BigNumber((_555_v21).length)));
      _nw89[(new BigNumber(14)).toNumber()] = (_552_v18).Merge(_552_v18);
      _nw89[(new BigNumber(15)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(16)).toNumber()] = (_552_v18).Merge(_552_v18);
      _nw89[(new BigNumber(17)).toNumber()] = (_552_v18).Merge(_552_v18);
      _nw89[(new BigNumber(18)).toNumber()] = ((_module.D6.create_DC12(_552_v18)).dtor_cf21).Merge(_dafny.Map.Empty.slice().updateUnsafe((_534_v1).f12,_553_v19));
      _nw89[(new BigNumber(19)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(20)).toNumber()] = (_552_v18).update((_this).f15, _553_v19);
      _nw89[(new BigNumber(21)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(22)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_534_v1).f12,_553_v19)).update((_this).f15, _553_v19);
      _nw89[(new BigNumber(23)).toNumber()] = _552_v18;
      _nw89[(new BigNumber(24)).toNumber()] = _552_v18;
      _556_v22 = _nw89;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_556_v22).length))) {
        let _557_i1 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_557_i1)) && ((_557_i1).isLessThan(new BigNumber((_556_v22).length))))) {
          (_556_v22)[(_557_i1)] = _552_v18;
        }
      }
      let _558_v23;
      _558_v23 = _540_v7;
      _558_v23 = _540_v7;
      r0 = (_534_v1).f12;
      let _nw90 = Array((new BigNumber(27)).toNumber()).fill(false);
      r1 = _nw90;
      return [r0, r1];
    }
    m4(p0, p1, globalState) {
      let _this = this;
      (globalState).f2 = (p1).plus(p1);
      let _559_v0;
      _559_v0 = new _dafny.CodePoint('v'.codePointAt(0));
      let _560_v1;
      _560_v1 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
      let _561_v2;
      _561_v2 = _dafny.Map.Empty.slice().updateUnsafe((((_560_v1).contains(p0)) ? ((_560_v1).get(p0)) : (p1)),_559_v0);
      let _562_v3;
      _562_v3 = _dafny.MultiSet.fromElements((_this).f15);
      _559_v0 = (((_561_v2).contains((new BigNumber(692)).multipliedBy((((_562_v3).contains(_this.f14)) ? ((_562_v3).get(_this.f14)) : (new BigNumber(648)))))) ? ((_561_v2).get((new BigNumber(692)).multipliedBy((((_562_v3).contains(_this.f14)) ? ((_562_v3).get(_this.f14)) : (new BigNumber(648)))))) : (_559_v0));
      let _563_v4;
      _563_v4 = p1;
      let _564_i0;
      _564_i0 = _dafny.ZERO;
      L4: {
        while (function (_source12) {
          let _574___mcc_h0 = _source12;
          let _575_cf0 = _574___mcc_h0;
          return _this.f14;
        }(_563_v4)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_564_i0)) {
              break L4;
            }
            _564_i0 = (_564_i0).plus(_dafny.ONE);
            let _565_v5;
            _565_v5 = _dafny.MultiSet.fromElements(new BigNumber(685), new BigNumber((_562_v3).cardinality()), new BigNumber(525), p1);
            let _566_v6;
            let _nw91 = Array((new BigNumber(3)).toNumber()).fill([]);
            _566_v6 = _nw91;
            let _567_v7;
            _567_v7 = _module.D3.create_DC5(_566_v6);
            let _568_v8;
            let _nw92 = new _module.C1();
            _nw92.__ctor(_this.f14, _567_v7);
            _568_v8 = _nw92;
            let _569_v9;
            _569_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_568_v8);
            (globalState).f2 = (((((_565_v5).contains(new BigNumber((_569_v9).length))) ? ((_565_v5).get(new BigNumber((_569_v9).length))) : (p1))).plus(p1)).multipliedBy((_dafny.ZERO).minus(p1));
            (_this).f14 = p0;
            let _570_v10;
            _570_v10 = _dafny.Seq.UnicodeFromString("e");
            let _571_v11;
            _571_v11 = _dafny.Map.Empty.slice().updateUnsafe((p0) || (p0),!((_dafny.ZERO).minus(new BigNumber((_570_v10).length))).isEqualTo(p1));
            _571_v11 = _module.__default.fm22(globalState);
            let _572_v12;
            let _nw93 = Array((new BigNumber(3)).toNumber());
            _nw93[(_dafny.ZERO).toNumber()] = (_this).f15;
            _nw93[(_dafny.ONE).toNumber()] = !(p0);
            _nw93[(new BigNumber(2)).toNumber()] = p0;
            _572_v12 = _nw93;
            let _573_v13;
            _573_v13 = _dafny.MultiSet.fromElements(_572_v12);
            _573_v13 = (_573_v13).Difference((_573_v13).update(_572_v12, _module.__default.abs(new BigNumber((_560_v1).length))));
          }
        }
      }
      let _576_i1;
      _576_i1 = _dafny.ZERO;
      L5: {
        while (p0) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_576_i1)) {
              break L5;
            }
            _576_i1 = (_576_i1).plus(_dafny.ONE);
            (_this).f14 = (_this).f15;
            let _577_v14;
            _577_v14 = _dafny.Set.fromElements(_559_v0, new _dafny.CodePoint('y'.codePointAt(0)));
            let _578_v15;
            _578_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,p0);
            let _579_v16;
            let _nw94 = new _module.C0();
            _nw94.__ctor(p0, new BigNumber((_578_v15).length));
            _579_v16 = _nw94;
            let _580_v17;
            _580_v17 = _dafny.Seq.of(_579_v16);
            let _581_v18;
            _581_v18 = _module.D3.create_DC8((_this).f15, (_579_v16).f13, (_579_v16).f13);
            let _582_v19;
            let _nw95 = Array((new BigNumber(20)).toNumber());
            _nw95[(_dafny.ZERO).toNumber()] = (_this).f15;
            _nw95[(_dafny.ONE).toNumber()] = _this.f14;
            _nw95[(new BigNumber(2)).toNumber()] = true;
            _nw95[(new BigNumber(3)).toNumber()] = p0;
            _nw95[(new BigNumber(4)).toNumber()] = p0;
            _nw95[(new BigNumber(5)).toNumber()] = (p1).isLessThan(p1);
            _nw95[(new BigNumber(6)).toNumber()] = (_this).f15;
            _nw95[(new BigNumber(7)).toNumber()] = (_this).f15;
            _nw95[(new BigNumber(8)).toNumber()] = (_this).fm4(p1, p0, _559_v0, globalState);
            _nw95[(new BigNumber(9)).toNumber()] = false;
            _nw95[(new BigNumber(10)).toNumber()] = (_577_v14).IsDisjointFrom(_577_v14);
            _nw95[(new BigNumber(11)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_579_v16, _579_v16), _580_v17);
            _nw95[(new BigNumber(12)).toNumber()] = p0;
            _nw95[(new BigNumber(13)).toNumber()] = (true) && ((_579_v16).f12);
            _nw95[(new BigNumber(14)).toNumber()] = true;
            _nw95[(new BigNumber(15)).toNumber()] = (_this).f15;
            _nw95[(new BigNumber(16)).toNumber()] = (_579_v16).f12;
            _nw95[(new BigNumber(17)).toNumber()] = (_579_v16).f12;
            _nw95[(new BigNumber(18)).toNumber()] = (((_581_v18).dtor_cf13) ? ((_this).f15) : ((_579_v16).f12));
            _nw95[(new BigNumber(19)).toNumber()] = _this.f14;
            _582_v19 = _nw95;
            let _index100 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_582_v19).length));
            (_582_v19)[_index100] = (_this).f15;
            let _583_v20;
            _583_v20 = _module.D7.create_DC15((_579_v16).f13, _this.f14);
            let _index101 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_582_v19).length));
            (_582_v19)[_index101] = ((!(_this.f14)) ? (_this.f14) : (((_583_v20).dtor_cf24).isLessThanOrEqualTo(p1)));
            let _584_v21;
            _584_v21 = _dafny.Map.Empty.slice().updateUnsafe((_579_v16).f13,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-457)), ((_585_v0) => function (_586_i2) {
              return _585_v0;
            })(_559_v0))).length));
            let _587_v22;
            _587_v22 = _dafny.Set.fromElements((_579_v16).f13, (_dafny.ZERO).minus((((_584_v21).contains((_579_v16).f13)) ? ((_584_v21).get((_579_v16).f13)) : (new BigNumber(-421)))));
            let _588_v23;
            let _nw96 = Array((new BigNumber(19)).toNumber());
            _nw96[(_dafny.ZERO).toNumber()] = _582_v19;
            _nw96[(_dafny.ONE).toNumber()] = _582_v19;
            _nw96[(new BigNumber(2)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(3)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(4)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(5)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(6)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(7)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(8)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(9)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(10)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(11)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(12)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(13)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(14)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(15)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(16)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(17)).toNumber()] = _582_v19;
            _nw96[(new BigNumber(18)).toNumber()] = _582_v19;
            _588_v23 = _nw96;
            let _589_v24;
            _589_v24 = _module.D3.create_DC5(_588_v23);
            let _590_v25;
            let _nw97 = new _module.C1();
            _nw97.__ctor(false, _589_v24);
            _590_v25 = _nw97;
            let _591_v26;
            _591_v26 = _dafny.Map.Empty.slice().updateUnsafe(_590_v25,_587_v22);
            _587_v22 = (((_591_v26).contains(_590_v25)) ? ((_591_v26).get(_590_v25)) : ((_587_v22).Intersect(_587_v22)));
            let _592_v28;
            _592_v28 = _dafny.Seq.of(p1);
            let _593_v29;
            _593_v29 = _dafny.Seq.of(_562_v3);
            let _594_v30;
            _594_v30 = _dafny.Seq.UnicodeFromString("leigq");
            let _595_v31;
            _595_v31 = _dafny.Map.Empty.slice().updateUnsafe(_594_v30,(_582_v19)[_module.__default.safeIndex(new BigNumber(66), new BigNumber((_582_v19).length))]);
            let _596_v32;
            _596_v32 = _dafny.Set.fromElements(_this.f14, _module.__default.fm0(p1, new BigNumber((function () {
              let _coll17 = new _dafny.Map();
              for (const _compr_17 of (_592_v28).Elements) {
                let _597_v27 = _compr_17;
                if (_dafny.Seq.contains(_592_v28, _597_v27)) {
                  _coll17.push([(_597_v27).multipliedBy((_579_v16).f13),(_this).f15]);
                }
              }
              return _coll17;
            }()).length), _593_v29, _595_v31, globalState));
            let _598_v33;
            _598_v33 = _dafny.Seq.of(_596_v32);
            let _599_v34;
            _599_v34 = _dafny.MultiSet.fromElements((new BigNumber((_562_v3).cardinality())).multipliedBy(new BigNumber((_596_v32).length)), p1, p1, _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(((_598_v33)[_module.__default.safeIndex((_579_v16).f13, new BigNumber((_598_v33).length))]).length)), p1), (p1).minus(p1));
            (globalState).f2 = (_dafny.ZERO).minus((((_599_v34).contains(p1)) ? ((_599_v34).get(p1)) : (p1)));
          }
        }
      }
      let _600_v35;
      _600_v35 = _dafny.Seq.of(_this.f14, _this.f14);
      let _601_v36;
      let _602_v37;
      let _out46;
      let _out47;
      let _outcollector14 = (_this).m3(_600_v35, globalState);
      _out46 = _outcollector14[0];
      _out47 = _outcollector14[1];
      _601_v36 = _out46;
      _602_v37 = _out47;
      let _603_v38;
      _603_v38 = _module.D4.create_DC9((_562_v3).update(_601_v36, _module.__default.abs(p1)));
      _603_v38 = ((p0) ? (_603_v38) : (_603_v38));
      return;
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f10 = _dafny.ZERO;
      this._f11 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    set f10(value) {
      let _this = this;
      _this._f10 = value;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f10, f11) {
      let _this = this;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, p1, p2, globalState) {
      let _this = this;
      return !(true) || ((_dafny.Set.fromElements(new BigNumber(246))).IsProperSubsetOf(function () {
        let _coll18 = new _dafny.Set();
        for (const _compr_18 of (_dafny.Set.fromElements(_this.f10)).Elements) {
          let _604_v0 = _compr_18;
          if ((_dafny.Set.fromElements(_this.f10)).contains(_604_v0)) {
            _coll18.add(_module.__default.safeModuloInt(_604_v0, new BigNumber(777)));
          }
        }
        return _coll18;
      }()));
    };
    fm6(p0, globalState) {
      let _this = this;
      let _source13 = _module.D1.create_DC2(!(true));
      if (_source13.is_DC2) {
        let _605___mcc_h0 = (_source13).cf2;
        let _606_cf2 = _605___mcc_h0;
        return _module.D1.create_DC2(_606_cf2);
      } else if (_source13.is_DC3) {
        let _607___mcc_h1 = (_source13).cf3;
        let _608___mcc_h2 = (_source13).cf4;
        let _609_cf4 = _608___mcc_h2;
        let _610_cf3 = _607___mcc_h1;
        return _module.D1.create_DC2(_610_cf3);
      } else {
        let _611___mcc_h3 = (_source13).cf1;
        let _612_cf1 = _611___mcc_h3;
        return _module.D1.create_DC2(false);
      }
    };
    fm3(globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_this.f10,_dafny.Map.Empty.slice().updateUnsafe(true,_this.f10))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f11).length),_dafny.Map.Empty.slice().updateUnsafe(true,_this.f10)));
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(_this.f10)).isLessThanOrEqualTo(_this.f10);
    };
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return !(((true) ? (((!(true)) ? (!(true)) : (false))) : (false)));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _613_v0;
      _613_v0 = true;
      _613_v0 = ((_613_v0) ? (_613_v0) : (_613_v0));
      let _614_v1;
      _614_v1 = _dafny.Seq.of(_this.f10);
      let _615_v2;
      _615_v2 = _dafny.Set.fromElements(_this.f10, _this.f10, _this.f10);
      if (((_613_v0) ? (((_614_v1)[_module.__default.safeIndex(new BigNumber((_615_v2).length), new BigNumber((_614_v1).length))]).isLessThanOrEqualTo(_this.f10)) : (_613_v0))) {
        let _616_v3;
        let _nw98 = new _module.C0();
        _nw98.__ctor(((_dafny.ZERO).minus(_this.f10)).isLessThanOrEqualTo(_this.f10), (_dafny.ZERO).minus(_this.f10));
        _616_v3 = _nw98;
        if ((_613_v0) || ((_616_v3).f12)) {
          (_this).f10 = (_616_v3).f13;
          let _617_v4;
          let _618_v5;
          let _619_v6;
          let _620_v7;
          let _out48;
          let _out49;
          let _out50;
          let _out51;
          let _outcollector15 = _module.__default.m0(globalState);
          _out48 = _outcollector15[0];
          _out49 = _outcollector15[1];
          _out50 = _outcollector15[2];
          _out51 = _outcollector15[3];
          _617_v4 = _out48;
          _618_v5 = _out49;
          _619_v6 = _out50;
          _620_v7 = _out51;
          let _621_v8;
          let _init7 = ((_622_v3) => function (_623_i0) {
            return (new BigNumber((_dafny.Seq.of((_622_v3).f12)).length)).isEqualTo((_622_v3).f13);
          })(_616_v3);
          let _nw99 = Array((new BigNumber(21)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw99.length); _i0_7++) {
            _nw99[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _621_v8 = _nw99;
          _621_v8 = _621_v8;
          let _624_v9;
          let _init8 = ((_625_v1) => function (_626_i1) {
            return _module.__default.safeModuloInt(_626_i1, new BigNumber((_625_v1).length));
          })(_614_v1);
          let _nw100 = Array((new BigNumber(29)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw100.length); _i0_8++) {
            _nw100[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _624_v9 = _nw100;
          let _627_v10;
          let _nw101 = Array((new BigNumber(23)).toNumber());
          _627_v10 = _nw101;
          let _628_v11;
          _628_v11 = _dafny.Map.Empty.slice().updateUnsafe(_627_v10,_624_v9);
          let _629_v12;
          _629_v12 = _624_v9;
          let _630_v13;
          _630_v13 = _dafny.Seq.of(_624_v9, _624_v9, _624_v9, _624_v9, _624_v9);
          let _631_v14;
          let _nw102 = Array((new BigNumber(29)).toNumber());
          _nw102[(_dafny.ZERO).toNumber()] = ((_619_v6) ? (_624_v9) : (_624_v9));
          _nw102[(_dafny.ONE).toNumber()] = _624_v9;
          _nw102[(new BigNumber(2)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(3)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(4)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(5)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(6)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(7)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(8)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(9)).toNumber()] = (((_628_v11).contains(_627_v10)) ? ((_628_v11).get(_627_v10)) : (_624_v9));
          _nw102[(new BigNumber(10)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(11)).toNumber()] = (((_616_v3).f12) ? (_624_v9) : (_624_v9));
          _nw102[(new BigNumber(12)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(13)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(14)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(15)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(16)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(17)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(18)).toNumber()] = (_629_v12);
          _nw102[(new BigNumber(19)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(20)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(21)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(22)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(23)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(24)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(25)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(26)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(27)).toNumber()] = _624_v9;
          _nw102[(new BigNumber(28)).toNumber()] = (_630_v13)[_module.__default.safeIndex(_617_v4, new BigNumber((_630_v13).length))];
          _631_v14 = _nw102;
          let _index102 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_631_v14).length));
          (_631_v14)[_index102] = _624_v9;
          let _index103 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_631_v14).length));
          (_631_v14)[_index103] = _624_v9;
          let _index104 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_621_v8).length));
          (_621_v8)[_index104] = _618_v5;
          let _632_v15;
          let _nw103 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Set.Empty);
          _632_v15 = _nw103;
          let _633_v16;
          _633_v16 = _dafny.Seq.of((_616_v3).f12);
          let _634_v17;
          _634_v17 = _dafny.Map.Empty.slice().updateUnsafe(_633_v16,_this.f10);
          let _index105 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_632_v15).length));
          (_632_v15)[_index105] = _module.__default.fm9(_613_v0, _634_v17, _617_v4, new _dafny.CodePoint('u'.codePointAt(0)), globalState);
          let _635_v18;
          let _nw104 = Array((new BigNumber(3)).toNumber()).fill([]);
          _635_v18 = _nw104;
          let _636_v19;
          _636_v19 = _dafny.Seq.of(_635_v18);
          let _637_v20;
          let _nw105 = Array((new BigNumber(12)).toNumber());
          _nw105[(_dafny.ZERO).toNumber()] = (_module.D3.create_DC5(_635_v18)).dtor_cf6;
          _nw105[(_dafny.ONE).toNumber()] = _635_v18;
          _nw105[(new BigNumber(2)).toNumber()] = _635_v18;
          _nw105[(new BigNumber(3)).toNumber()] = (_636_v19)[_module.__default.safeIndex((_616_v3).f13, new BigNumber((_636_v19).length))];
          _nw105[(new BigNumber(4)).toNumber()] = _635_v18;
          _nw105[(new BigNumber(5)).toNumber()] = _635_v18;
          _nw105[(new BigNumber(6)).toNumber()] = _635_v18;
          _nw105[(new BigNumber(7)).toNumber()] = _635_v18;
          _nw105[(new BigNumber(8)).toNumber()] = _635_v18;
          _nw105[(new BigNumber(9)).toNumber()] = _635_v18;
          _nw105[(new BigNumber(10)).toNumber()] = ((true) ? (_635_v18) : (_635_v18));
          _nw105[(new BigNumber(11)).toNumber()] = _635_v18;
          _637_v20 = _nw105;
          let _index106 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_637_v20).length));
          (_637_v20)[_index106] = _635_v18;
          let _638_v21;
          _638_v21 = new _dafny.CodePoint('m'.codePointAt(0));
          let _639_v22;
          _639_v22 = _dafny.Seq.of(_638_v21, _638_v21);
          let _640_v23;
          _640_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10(_618_v5, globalState),new BigNumber(-133));
          let _641_v24;
          _641_v24 = _dafny.MultiSet.fromElements(_613_v0);
          let _642_v25;
          _642_v25 = _module.D4.create_DC9(_641_v24);
          let _643_v26;
          let _nw106 = Array((new BigNumber(20)).toNumber());
          _nw106[(_dafny.ZERO).toNumber()] = _638_v21;
          _nw106[(_dafny.ONE).toNumber()] = _638_v21;
          _nw106[(new BigNumber(2)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(3)).toNumber()] = (_639_v22)[_module.__default.safeIndex((_616_v3).f13, new BigNumber((_639_v22).length))];
          _nw106[(new BigNumber(4)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(5)).toNumber()] = (_639_v22)[_module.__default.safeIndex((((_640_v23).contains((((_620_v7).contains((_616_v3).f12)) ? ((_620_v7).get((_616_v3).f12)) : (new BigNumber(((_642_v25).dtor_cf16).cardinality()))))) ? ((_640_v23).get((((_620_v7).contains((_616_v3).f12)) ? ((_620_v7).get((_616_v3).f12)) : (new BigNumber(((_642_v25).dtor_cf16).cardinality()))))) : ((_616_v3).f13)), new BigNumber((_639_v22).length))];
          _nw106[(new BigNumber(6)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(7)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(8)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(9)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(10)).toNumber()] = ((_618_v5) ? (_638_v21) : (_638_v21));
          _nw106[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
          _nw106[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _nw106[(new BigNumber(13)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(14)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(15)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(16)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(17)).toNumber()] = _638_v21;
          _nw106[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
          _nw106[(new BigNumber(19)).toNumber()] = _638_v21;
          _643_v26 = _nw106;
          let _index107 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_643_v26).length));
          (_643_v26)[_index107] = _638_v21;
          let _644_v27;
          _644_v27 = _dafny.Seq.of(_641_v24, _dafny.MultiSet.FromArray(_dafny.Seq.of(_613_v0)));
          let _645_v28;
          _645_v28 = _dafny.Map.Empty.slice().updateUnsafe(_639_v22,_619_v6);
          let _646_v29;
          _646_v29 = _dafny.Set.fromElements(_613_v0, _module.__default.fm0(new BigNumber((_614_v1).length), _this.f10, _644_v27, _645_v28, globalState), (_616_v3).f12, _613_v0);
          let _647_v31;
          _647_v31 = _dafny.Map.Empty.slice().updateUnsafe(_639_v22,(_616_v3).f13);
          let _648_v32;
          _648_v32 = _dafny.Seq.of(function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of (_647_v31).Keys.Elements) {
              let _649_v30 = _compr_19;
              if ((_647_v31).contains(_649_v30)) {
                _coll19.push([_649_v30,_618_v5]);
              }
            }
            return _coll19;
          }());
          let _index108 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_621_v8).length));
          let _index109 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_632_v15).length));
          let _index110 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_637_v20).length));
          let _index111 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_643_v26).length));
          let _rhs126 = ((_618_v5) ? (_module.__default.fm0((_dafny.ZERO).minus(new BigNumber((_646_v29).length)), _this.f10, _module.__default.fm11(globalState), (_648_v32)[_module.__default.safeIndex((_616_v3).f13, new BigNumber((_648_v32).length))], globalState)) : (_613_v0));
          let _rhs127 = ((_616_v3).f13).multipliedBy(_this.f10);
          let _rhs128 = _615_v2;
          let _rhs129 = _635_v18;
          let _rhs130 = (_639_v22)[_module.__default.safeIndex(_617_v4, new BigNumber((_639_v22).length))];
          let _lhs81 = _621_v8;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_621_v8).length));
          let _lhs83 = _632_v15;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_632_v15).length));
          let _lhs85 = _637_v20;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_637_v20).length));
          let _lhs87 = _643_v26;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_643_v26).length));
          _lhs81[_lhs82] = _rhs126;
          _617_v4 = _rhs127;
          _lhs83[_lhs84] = _rhs128;
          _lhs85[_lhs86] = _rhs129;
          _lhs87[_lhs88] = _rhs130;
        } else {
          _614_v1 = _614_v1;
          r0 = ((_616_v3).f13).minus(new BigNumber(-909));
          let _650_v33;
          let _nw107 = Array((new BigNumber(3)).toNumber()).fill(false);
          _650_v33 = _nw107;
          let _rhs131 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_616_v3).f13, (_616_v3).f13), _module.__default.safeModuloInt((_616_v3).f13, _this.f10));
          let _rhs132 = _650_v33;
          r0 = _rhs131;
          _650_v33 = _rhs132;
          let _651_v34;
          _651_v34 = _dafny.Seq.UnicodeFromString("ajenxyshb");
          r0 = _module.__default.safeModuloInt((_614_v1)[_module.__default.safeIndex((_616_v3).f13, new BigNumber((_614_v1).length))], new BigNumber((_651_v34).length));
          let _652_v35;
          _652_v35 = _651_v34;
          let _653_v36;
          _653_v36 = _dafny.Seq.of(_module.__default.fm2(false, new BigNumber(996), (_616_v3).f13, !(_613_v0), globalState), _dafny.Seq.Concat((_651_v34), _651_v34));
          _653_v36 = _653_v36;
        }
        let _654_v37;
        _654_v37 = _dafny.Seq.UnicodeFromString("adcgwephx");
        _654_v37 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lndvj"), _654_v37);
        if ((_module.D3.create_DC8(!((_616_v3).f12), new BigNumber((_615_v2).length), (_616_v3).f13)).dtor_cf13) {
          _613_v0 = (_616_v3).f12;
          let _655_v38;
          let _init9 = ((_656_v0) => function (_657_i2) {
            return _656_v0;
          })(_613_v0);
          let _nw108 = Array((new BigNumber(8)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw108.length); _i0_9++) {
            _nw108[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _655_v38 = _nw108;
          let _index112 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_655_v38).length));
          (_655_v38)[_index112] = _613_v0;
          let _index113 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_655_v38).length));
          (_655_v38)[_index113] = _613_v0;
          let _658_v39;
          _658_v39 = _module.D3.create_DC8((_655_v38)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_655_v38).length))], _this.f10, new BigNumber(704));
          let _index114 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_655_v38).length));
          (_655_v38)[_index114] = (_module.__default.safeModuloInt((_658_v39).dtor_cf15, new BigNumber((_654_v37).length))).isLessThan(new BigNumber((_614_v1).length));
          _613_v0 = !(_613_v0);
          let _659_v40;
          _659_v40 = _654_v37;
          let _660_v41;
          _660_v41 = _dafny.Map.Empty.slice().updateUnsafe((_659_v40),(_dafny.ZERO).minus((_616_v3).f13));
          _660_v41 = (_660_v41).update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wayojdhbc"), _dafny.Seq.UnicodeFromString("kwne")), (_616_v3).f13);
        } else {
          let _661_v42;
          let _init10 = ((_662_v0) => function (_663_i3) {
            return _662_v0;
          })(_613_v0);
          let _nw109 = Array((new BigNumber(8)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw109.length); _i0_10++) {
            _nw109[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _661_v42 = _nw109;
          let _index115 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length));
          (_661_v42)[_index115] = !(!(!((_616_v3).f12)));
          let _index116 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length));
          (_661_v42)[_index116] = false;
          _616_v3 = _616_v3;
          _613_v0 = (_661_v42)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length))];
          let _index117 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length));
          let _index118 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length));
          let _rhs133 = (_616_v3).f12;
          let _rhs134 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(377)), function (_664_i4) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          }), _654_v37);
          let _rhs135 = (_616_v3).f12;
          let _rhs136 = (((_616_v3).f12) && ((_661_v42)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length))])) || ((_616_v3).f12);
          let _lhs89 = _661_v42;
          let _lhs90 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length));
          let _lhs91 = _661_v42;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length));
          _613_v0 = _rhs133;
          _654_v37 = _rhs134;
          _lhs89[_lhs90] = _rhs135;
          _lhs91[_lhs92] = _rhs136;
          let _665_v43;
          let _init11 = function (_666_i5) {
            return (_666_i5).minus(_this.f10);
          };
          let _nw110 = Array((new BigNumber(6)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw110.length); _i0_11++) {
            _nw110[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _665_v43 = _nw110;
          let _667_v44;
          _667_v44 = _dafny.Map.Empty.slice().updateUnsafe((_661_v42)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length))],(_661_v42)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_661_v42).length))]);
          let _668_v45;
          _668_v45 = _module.D1.create_DC3(_613_v0, _module.__default.fm1((((_667_v44).contains((_616_v3).f12)) ? ((_667_v44).get((_616_v3).f12)) : (true)), _613_v0, globalState));
          let _669_v46;
          _669_v46 = _dafny.Map.Empty.slice().updateUnsafe(_665_v43,function (_pat_let4_0) {
            return function (_670_dt__update__tmp_h1) {
              return function (_pat_let5_0) {
                return function (_671_dt__update_hcf3_h0) {
                  return _module.D1.create_DC3(_671_dt__update_hcf3_h0, (_670_dt__update__tmp_h1).dtor_cf4);
                }(_pat_let5_0);
              }(!(false));
            }(_pat_let4_0);
          }(_668_v45));
          (globalState).f2 = new BigNumber((_669_v46).length);
        }
        let _672_v47;
        _672_v47 = _module.D1.create_DC2(_613_v0);
        let _source14 = _672_v47;
        if (_source14.is_DC2) {
          let _673___mcc_h0 = (_source14).cf2;
          let _674_cf2 = _673___mcc_h0;
          (_this).f10 = (_dafny.ZERO).minus((((_615_v2).IsProperSubsetOf(_615_v2)) ? (new BigNumber((_dafny.Seq.Concat(_614_v1, _dafny.Seq.of((_614_v1)[_module.__default.safeIndex(new BigNumber(-325), new BigNumber((_614_v1).length))]))).length)) : (_module.__default.fm10((_616_v3).f12, globalState))));
          let _675_v48;
          _675_v48 = _dafny.Map.Empty.slice().updateUnsafe(_674_cf2,_module.__default.fm10((_616_v3).f12, globalState));
          let _676_v49;
          _676_v49 = _dafny.Map.Empty.slice().updateUnsafe(_654_v37,_674_cf2);
          let _677_v50;
          _677_v50 = _dafny.Map.Empty.slice().updateUnsafe((_616_v3).f12,_module.__default.fm0(new BigNumber((_675_v48).length), (_616_v3).f13, _module.__default.fm11(globalState), _676_v49, globalState));
          (_this).f10 = (_dafny.ZERO).minus(new BigNumber((_677_v50).length));
          _675_v48 = (_675_v48).update(true, (_616_v3).f13);
          let _678_v51;
          _678_v51 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(_613_v0),(_616_v3).f13),(_616_v3).f12);
          _678_v51 = (_678_v51).update(_675_v48, !(_613_v0) || ((_616_v3).f12));
        } else if (_source14.is_DC3) {
          let _679___mcc_h1 = (_source14).cf3;
          let _680___mcc_h2 = (_source14).cf4;
          let _681_cf4 = _680___mcc_h2;
          let _682_cf3 = _679___mcc_h1;
          let _683_v52;
          _683_v52 = _dafny.Seq.of(_682_cf3, _613_v0);
          let _684_v53;
          _684_v53 = _dafny.MultiSet.fromElements((_616_v3).f13, _this.f10);
          let _685_v54;
          _685_v54 = _module.D3.create_DC8(_dafny.Seq.IsPrefixOf(_683_v52, _683_v52), (_616_v3).f13, (((_684_v53).contains(_this.f10)) ? ((_684_v53).get(_this.f10)) : (_this.f10)));
          _685_v54 = _685_v54;
          (globalState).f2 = _this.f10;
          let _686_v55;
          _686_v55 = _dafny.MultiSet.fromElements(_613_v0, _682_cf3);
          let _687_v56;
          _687_v56 = _dafny.Map.Empty.slice().updateUnsafe(_686_v55,(_616_v3).f12);
          let _688_v57;
          _688_v57 = _module.D3.create_DC6((_616_v3).f13, _613_v0, (_616_v3).f12, _this.f10, (_616_v3).f12);
          _687_v56 = (_687_v56).update((_686_v55).update((_616_v3).f12, _module.__default.abs(new BigNumber((_615_v2).length))), ((_613_v0) ? ((_688_v57).dtor_cf8) : ((_616_v3).f12)));
          _613_v0 = (_616_v3).f12;
        } else {
          let _689___mcc_h3 = (_source14).cf1;
          let _690_cf1 = _689___mcc_h3;
          let _691_v58;
          _691_v58 = _dafny.Map.Empty.slice().updateUnsafe((_616_v3).f13,!((_616_v3).f12));
          _691_v58 = (_691_v58).update((_616_v3).f13, (_616_v3).f12);
          let _692_v59;
          _692_v59 = _module.D1.create_DC1((_616_v3).f12);
          let _693_v60;
          _693_v60 = _dafny.Seq.of(_613_v0, (_616_v3).f12);
          let _694_v61;
          _694_v61 = _dafny.Map.Empty.slice().updateUnsafe((_693_v60)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f10), new BigNumber((_693_v60).length))],!((_616_v3).f12));
          _690_cf1 = ((_this).fm7(_dafny.Map.Empty.slice().updateUnsafe((_616_v3).f12,_692_v59), _690_cf1, (_616_v3).f13, globalState)) || (!(new BigNumber((_694_v61).length)).isEqualTo(_this.f10));
          (globalState).f2 = new BigNumber((_654_v37).length);
          let _695_v62;
          let _init12 = function (_696_i6) {
            return _module.__default.safeDivisionInt(_696_i6, new BigNumber(850));
          };
          let _nw111 = Array((new BigNumber(15)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw111.length); _i0_12++) {
            _nw111[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _695_v62 = _nw111;
          let _697_v63;
          _697_v63 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,_695_v62);
          let _698_v64;
          let _nw112 = Array((new BigNumber(11)).toNumber());
          _nw112[(_dafny.ZERO).toNumber()] = _695_v62;
          _nw112[(_dafny.ONE).toNumber()] = _695_v62;
          _nw112[(new BigNumber(2)).toNumber()] = _695_v62;
          _nw112[(new BigNumber(3)).toNumber()] = _695_v62;
          _nw112[(new BigNumber(4)).toNumber()] = _695_v62;
          _nw112[(new BigNumber(5)).toNumber()] = _695_v62;
          _nw112[(new BigNumber(6)).toNumber()] = (((_697_v63).contains(_this.f10)) ? ((_697_v63).get(_this.f10)) : (_695_v62));
          _nw112[(new BigNumber(7)).toNumber()] = _695_v62;
          _nw112[(new BigNumber(8)).toNumber()] = _695_v62;
          _nw112[(new BigNumber(9)).toNumber()] = _695_v62;
          _nw112[(new BigNumber(10)).toNumber()] = _695_v62;
          _698_v64 = _nw112;
          let _index119 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_698_v64).length));
          (_698_v64)[_index119] = _695_v62;
          let _index120 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_698_v64).length));
          (_698_v64)[_index120] = _695_v62;
        }
      } else {
        let _699_v65;
        let _init13 = ((_700_v0) => function (_701_i7) {
          return _module.__default.safeDivisionInt(_701_i7, new BigNumber((_dafny.Seq.of(false, _700_v0)).length));
        })(_613_v0);
        let _nw113 = Array((new BigNumber(5)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw113.length); _i0_13++) {
          _nw113[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _699_v65 = _nw113;
        let _index121 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_699_v65).length));
        (_699_v65)[_index121] = _this.f10;
        let _702_v66;
        _702_v66 = _dafny.Set.fromElements(_613_v0, _613_v0, _613_v0);
        let _703_v67;
        _703_v67 = _dafny.Map.Empty.slice().updateUnsafe(_613_v0,_this.f10);
        let _index122 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_699_v65).length));
        let _rhs137 = _613_v0;
        let _rhs138 = (_702_v66).IsProperSubsetOf(_702_v66);
        let _rhs139 = (_this.f10).plus(new BigNumber(-305));
        let _rhs140 = (_dafny.ZERO).minus(new BigNumber((_703_v67).length));
        let _lhs93 = _699_v65;
        let _lhs94 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_699_v65).length));
        let _lhs95 = globalState;
        _613_v0 = _rhs137;
        _613_v0 = _rhs138;
        _lhs93[_lhs94] = _rhs139;
        _lhs95.f2 = _rhs140;
        let _704_v68;
        let _nw114 = Array((new BigNumber(21)).toNumber()).fill(false);
        _704_v68 = _nw114;
        let _705_v69;
        _705_v69 = _dafny.Seq.UnicodeFromString("cirg");
        let _index123 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_704_v68).length));
        (_704_v68)[_index123] = _dafny.areEqual(_705_v69, _705_v69);
        let _706_v70;
        let _nw115 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _706_v70 = _nw115;
        let _707_v71;
        let _nw116 = Array((new BigNumber(15)).toNumber()).fill([]);
        _707_v71 = _nw116;
        let _708_v72;
        let _nw117 = new _module.C1();
        _nw117.__ctor(_613_v0, _module.D3.create_DC5(_707_v71));
        _708_v72 = _nw117;
        let _709_v73;
        _709_v73 = _dafny.MultiSet.fromElements(_708_v72);
        let _710_v74;
        _710_v74 = _dafny.MultiSet.fromElements(_613_v0);
        let _711_v75;
        _711_v75 = _dafny.Map.Empty.slice().updateUnsafe(_710_v74,_613_v0);
        let _712_v76;
        _712_v76 = _dafny.Seq.of(_613_v0, _613_v0, _613_v0, _613_v0, _613_v0);
        let _713_v77;
        _713_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_711_v75).length),_dafny.MultiSet.fromElements((_699_v65)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_699_v65).length))], new BigNumber((_712_v76).length)));
        let _714_v78;
        _714_v78 = _dafny.Map.Empty.slice().updateUnsafe(_713_v77,_706_v70);
        let _715_v81;
        _715_v81 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,new BigNumber((_705_v69).length));
        let _716_v82;
        _716_v82 = new _dafny.CodePoint('a'.codePointAt(0));
        let _717_v83;
        _717_v83 = _dafny.Set.fromElements(_716_v82);
        let _index124 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_704_v68).length));
        let _rhs141 = (_615_v2).IsDisjointFrom(_615_v2);
        let _rhs142 = (((_709_v73).contains(_708_v72)) ? ((_709_v73).get(_708_v72)) : (new BigNumber(265)));
        let _rhs143 = (((_714_v78).contains((function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of _dafny.IntegerRange(new BigNumber(-398), new BigNumber(496))) {
            let _720_v79 = _compr_22;
            if (((new BigNumber(-398)).isLessThanOrEqualTo(_720_v79)) && ((_720_v79).isLessThan(new BigNumber(496)))) {
              _coll22.push([(_720_v79).minus(_this.f10),_dafny.MultiSet.fromElements(new BigNumber((_702_v66).length), (_699_v65)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_699_v65).length))])]);
            }
          }
          return _coll22;
        }()).Merge(function () {
          let _coll23 = new _dafny.Map();
          for (const _compr_23 of (_715_v81).Keys.Elements) {
            let _721_v80 = _compr_23;
            if ((_715_v81).contains(_721_v80)) {
              _coll23.push([_module.__default.safeModuloInt(_721_v80, new BigNumber(-284)),_dafny.MultiSet.fromElements(_this.f10)]);
            }
          }
          return _coll23;
        }()))) ? ((_714_v78).get((function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of _dafny.IntegerRange(new BigNumber(-398), new BigNumber(496))) {
            let _718_v79 = _compr_20;
            if (((new BigNumber(-398)).isLessThanOrEqualTo(_718_v79)) && ((_718_v79).isLessThan(new BigNumber(496)))) {
              _coll20.push([(_718_v79).minus(_this.f10),_dafny.MultiSet.fromElements(new BigNumber((_702_v66).length), (_699_v65)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_699_v65).length))])]);
            }
          }
          return _coll20;
        }()).Merge(function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of (_715_v81).Keys.Elements) {
            let _719_v80 = _compr_21;
            if ((_715_v81).contains(_719_v80)) {
              _coll21.push([_module.__default.safeModuloInt(_719_v80, new BigNumber(-284)),_dafny.MultiSet.fromElements(_this.f10)]);
            }
          }
          return _coll21;
        }()))) : (_706_v70));
        let _rhs144 = (((_613_v0) ? (_717_v83) : (_dafny.Set.fromElements(_716_v82, new _dafny.CodePoint('i'.codePointAt(0)))))).IsProperSubsetOf(_717_v83);
        let _rhs145 = (_module.__default.fm23(_715_v81, _613_v0, globalState)).IsDisjointFrom(_702_v66);
        let _lhs96 = _704_v68;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_704_v68).length));
        _lhs96[_lhs97] = _rhs141;
        r0 = _rhs142;
        _706_v70 = _rhs143;
        _613_v0 = _rhs144;
        _613_v0 = _rhs145;
        let _index125 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_704_v68).length));
        (_704_v68)[_index125] = (new BigNumber(((_module.__default.fm24(_613_v0, _712_v76, (_704_v68)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_704_v68).length))], globalState)).Union(_dafny.MultiSet.FromArray(_614_v1))).cardinality())).isLessThan((new BigNumber((_dafny.Seq.UnicodeFromString("ethwsioyp")).length)).minus((_699_v65)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_699_v65).length))]));
        let _722_v84;
        _722_v84 = _dafny.Map.Empty.slice().updateUnsafe(true,false);
        let _index126 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_704_v68).length));
        (_704_v68)[_index126] = (new BigNumber((_722_v84).length)).isLessThan(_this.f10);
        let _index127 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_704_v68).length));
        (_704_v68)[_index127] = (_704_v68)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_704_v68).length))];
      }
      (globalState).f2 = _module.__default.fm10(_613_v0, globalState);
      let _723_v85;
      _723_v85 = _dafny.Seq.of(_613_v0, true);
      let _724_v86;
      _724_v86 = _dafny.Map.Empty.slice().updateUnsafe(_723_v85,(_dafny.ZERO).minus(_this.f10));
      let _725_v88;
      _725_v88 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_723_v85));
      let _726_v89;
      _726_v89 = _dafny.Seq.UnicodeFromString("yjyrhbv");
      let _727_v90;
      _727_v90 = _dafny.Map.Empty.slice().updateUnsafe(_726_v89,true);
      let _728_v91;
      let _nw118 = new _module.C2();
      _nw118.__ctor(_613_v0, _module.__default.fm0(new BigNumber((function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of (_615_v2).Elements) {
          let _729_v87 = _compr_24;
          if ((_615_v2).contains(_729_v87)) {
            _coll24.push([_module.__default.safeDivisionInt(_729_v87, _this.f10),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)))).length))).cardinality())]);
          }
        }
        return _coll24;
      }()).length), _this.f10, _725_v88, _727_v90, globalState));
      _728_v91 = _nw118;
      let _730_v92;
      _730_v92 = _dafny.Seq.of(_728_v91, _728_v91);
      _724_v86 = (_724_v86).update(_723_v85, new BigNumber((_dafny.Seq.Concat(_730_v92, _730_v92)).length));
      let _731_v93;
      _731_v93 = _dafny.MultiSet.fromElements(new BigNumber((_723_v85).length));
      let _732_v94;
      _732_v94 = _dafny.MultiSet.fromElements(_this.f10, _this.f10, new BigNumber((_726_v89).length), new BigNumber((_731_v93).cardinality()));
      let _733_v95;
      _733_v95 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(567)), function (_734_i9) {
        return new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0)))).length);
      })));
      let _735_v96;
      let _nw119 = Array((new BigNumber(15)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(339)), function (_736_i8) {
        return new BigNumber(22);
      }), _614_v1));
      _nw119[(_dafny.ONE).toNumber()] = _732_v94;
      _nw119[(new BigNumber(2)).toNumber()] = _732_v94;
      _nw119[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_this.f10, new BigNumber(((_this).f11).length), new BigNumber((_dafny.MultiSet.fromElements((_728_v91).f15)).cardinality()), new BigNumber((_615_v2).length));
      _nw119[(new BigNumber(4)).toNumber()] = ((false) ? (_732_v94) : (_732_v94));
      _nw119[(new BigNumber(5)).toNumber()] = (_731_v93).update(_this.f10, _module.__default.abs(_this.f10));
      _nw119[(new BigNumber(6)).toNumber()] = (_732_v94).Intersect(_731_v93);
      _nw119[(new BigNumber(7)).toNumber()] = _module.__default.fm24((_728_v91).f15, _723_v85, _728_v91.f14, globalState);
      _nw119[(new BigNumber(8)).toNumber()] = (_732_v94).Intersect((_dafny.MultiSet.fromElements(_this.f10)).update(new BigNumber((_733_v95).cardinality()), _module.__default.abs(_this.f10)));
      _nw119[(new BigNumber(9)).toNumber()] = (_732_v94).Intersect(_731_v93);
      _nw119[(new BigNumber(10)).toNumber()] = (_module.__default.fm24((_728_v91).f15, _module.__default.fm18(_this.f10, (_728_v91).f15, new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_728_v91).f15), _module.__default.safeIndex(new BigNumber(258), new BigNumber((_dafny.Seq.of((_728_v91).f15)).length)), true)).length), globalState), (_728_v91).f15, globalState)).Intersect(_732_v94);
      _nw119[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_this.f10, (_dafny.ZERO).minus(_this.f10));
      _nw119[(new BigNumber(12)).toNumber()] = _731_v93;
      _nw119[(new BigNumber(13)).toNumber()] = _731_v93;
      _nw119[(new BigNumber(14)).toNumber()] = _731_v93;
      _735_v96 = _nw119;
      let _index128 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_735_v96).length));
      (_735_v96)[_index128] = _dafny.MultiSet.FromArray(_614_v1);
      let _index129 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_735_v96).length));
      (_735_v96)[_index129] = ((_dafny.MultiSet.fromElements(_this.f10, new BigNumber(552))).Intersect(_732_v94)).Intersect(_731_v93);
      let _737_v97;
      _737_v97 = _726_v89;
      let _738_v98;
      let _init14 = function (_739_i10) {
        return _module.__default.safeModuloInt(_739_i10, _this.f10);
      };
      let _nw120 = Array((new BigNumber(15)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw120.length); _i0_14++) {
        _nw120[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _738_v98 = _nw120;
      let _740_v99;
      _740_v99 = _dafny.Map.Empty.slice().updateUnsafe(_737_v97,_738_v98);
      let _741_v100;
      _741_v100 = _module.D8.create_DC17(_740_v99);
      (globalState).f2 = new BigNumber((((_740_v99).Merge((_741_v100).dtor_cf27)).Merge(_740_v99)).length);
      r0 = _this.f10;
      return r0;
    }
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _742_v0;
      _742_v0 = _module.D6.create_DC13(((_dafny.ZERO).minus(_this.f10)).isLessThan(_this.f10));
      let _source15 = _742_v0;
      if (_source15.is_DC13) {
        let _743___mcc_h0 = (_source15).cf22;
        let _744_cf22 = _743___mcc_h0;
        r0 = _744_cf22;
        let _745_v1;
        let _nw121 = Array((new BigNumber(4)).toNumber()).fill(false);
        _745_v1 = _nw121;
        let _746_v2;
        _746_v2 = _module.D7.create_DC14(_dafny.Set.fromElements(_745_v1, _745_v1, _745_v1, _745_v1));
        let _747_v3;
        _747_v3 = _dafny.Set.fromElements(_745_v1);
        let _pat_let_tv2 = _747_v3;
        let _source16 = function (_pat_let6_0) {
          return function (_748_dt__update__tmp_h0) {
            return function (_pat_let7_0) {
              return function (_749_dt__update_hcf23_h0) {
                return _module.D7.create_DC14(_749_dt__update_hcf23_h0);
              }(_pat_let7_0);
            }(_pat_let_tv2);
          }(_pat_let6_0);
        }(_746_v2);
        if (_source16.is_DC15) {
          let _750___mcc_h2 = (_source16).cf24;
          let _751___mcc_h3 = (_source16).cf25;
          let _752_cf25 = _751___mcc_h3;
          let _753_cf24 = _750___mcc_h2;
          (globalState).f2 = _753_cf24;
          (_this).f10 = (_module.__default.safeModuloInt(_this.f10, new BigNumber(317))).minus(((_dafny.ZERO).minus((_dafny.ZERO).minus(_753_cf24))).multipliedBy(_this.f10));
          let _754_v4;
          _754_v4 = _dafny.Map.Empty.slice().updateUnsafe(_752_cf25,_753_cf24);
          let _755_v5;
          _755_v5 = _dafny.MultiSet.fromElements(_753_cf24, _this.f10);
          let _756_v6;
          _756_v6 = _dafny.MultiSet.fromElements(_752_cf25);
          let _757_v7;
          _757_v7 = _module.D9.create_DC19(_754_v4);
          let _758_v8;
          _758_v8 = _dafny.MultiSet.fromElements(_754_v4, (_754_v4).Merge(_754_v4), (_754_v4).update(_752_cf25, _this.f10), _754_v4, (_dafny.Map.Empty.slice().updateUnsafe(_744_cf22,(((_755_v5).contains((((_756_v6).contains(_744_cf22)) ? ((_756_v6).get(_744_cf22)) : (_753_cf24)))) ? ((_755_v5).get((((_756_v6).contains(_744_cf22)) ? ((_756_v6).get(_744_cf22)) : (_753_cf24)))) : (_753_cf24)))).Merge((_757_v7).dtor_cf28));
          let _759_v9;
          _759_v9 = _dafny.Seq.UnicodeFromString("wiwrmiodu");
          let _rhs146 = _758_v8;
          let _rhs147 = _dafny.Seq.UnicodeFromString("qjdplvq");
          _758_v8 = _rhs146;
          _759_v9 = _rhs147;
          let _760_v10;
          let _761_v11;
          let _762_v12;
          let _763_v13;
          let _out52;
          let _out53;
          let _out54;
          let _out55;
          let _outcollector16 = _module.__default.m0(globalState);
          _out52 = _outcollector16[0];
          _out53 = _outcollector16[1];
          _out54 = _outcollector16[2];
          _out55 = _outcollector16[3];
          _760_v10 = _out52;
          _761_v11 = _out53;
          _762_v12 = _out54;
          _763_v13 = _out55;
        } else if (_source16.is_DC14) {
          let _764___mcc_h4 = (_source16).cf23;
          let _765_cf23 = _764___mcc_h4;
          let _766_v14;
          _766_v14 = new _dafny.CodePoint('j'.codePointAt(0));
          let _767_v15;
          _767_v15 = _module.D9.create_DC20(_744_cf22, _766_v14);
          let _768_v16;
          let _nw122 = Array((new BigNumber(13)).toNumber()).fill([]);
          _768_v16 = _nw122;
          let _769_v17;
          _769_v17 = _module.D3.create_DC5(_768_v16);
          let _770_v18;
          let _nw123 = new _module.C1();
          _nw123.__ctor(((_767_v15).dtor_cf29) || (_744_cf22), _769_v17);
          _770_v18 = _nw123;
          let _771_v19;
          _771_v19 = _module.D10.create_DC22(_770_v18);
          let _pat_let_tv3 = _770_v18;
          _770_v18 = (function (_pat_let8_0) {
            return function (_772_dt__update__tmp_h1) {
              return function (_pat_let9_0) {
                return function (_773_dt__update_hcf32_h0) {
                  return _module.D10.create_DC22(_773_dt__update_hcf32_h0);
                }(_pat_let9_0);
              }(_pat_let_tv3);
            }(_pat_let8_0);
          }(_771_v19)).dtor_cf32;
          let _774_v20;
          _774_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,(_this).f11);
          let _775_v21;
          _775_v21 = _dafny.Seq.of(_this.f10, _this.f10);
          let _776_v22;
          _776_v22 = _dafny.Map.Empty.slice().updateUnsafe(_774_v20,(_775_v21)[_module.__default.safeIndex(_this.f10, new BigNumber((_775_v21).length))]);
          _776_v22 = (_776_v22).update(function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of ((_this).f11).Keys.Elements) {
              let _777_v23 = _compr_25;
              if (((_this).f11).contains(_777_v23)) {
                _coll25.push([(_777_v23).minus(_this.f10),(_this).f11]);
              }
            }
            return _coll25;
          }(), _this.f10);
          let _778_v24;
          _778_v24 = _dafny.Seq.UnicodeFromString("d");
          let _779_v25;
          let _nw124 = Array((new BigNumber(18)).toNumber());
          _nw124[(_dafny.ZERO).toNumber()] = _775_v21;
          _nw124[(_dafny.ONE).toNumber()] = _module.__default.fm25(_this.f10, _this.f10, new BigNumber((_778_v24).length), _this.f10, globalState);
          _nw124[(new BigNumber(2)).toNumber()] = _775_v21;
          _nw124[(new BigNumber(3)).toNumber()] = _775_v21;
          _nw124[(new BigNumber(4)).toNumber()] = _775_v21;
          _nw124[(new BigNumber(5)).toNumber()] = _775_v21;
          _nw124[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_this.f10);
          _nw124[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_this.f10);
          _nw124[(new BigNumber(8)).toNumber()] = _775_v21;
          _nw124[(new BigNumber(9)).toNumber()] = _775_v21;
          _nw124[(new BigNumber(10)).toNumber()] = _775_v21;
          _nw124[(new BigNumber(11)).toNumber()] = _775_v21;
          _nw124[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(_this.f10, new BigNumber(758));
          _nw124[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(903)), function (_780_i0) {
            return _this.f10;
          });
          _nw124[(new BigNumber(14)).toNumber()] = _775_v21;
          _nw124[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_this.f10);
          _nw124[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(_this.f10);
          _nw124[(new BigNumber(17)).toNumber()] = _775_v21;
          _779_v25 = _nw124;
          let _781_v26;
          let _nw125 = new _module.C2();
          _nw125.__ctor((_770_v18).f16, ((((_this).f11).contains((_dafny.ZERO).minus(_this.f10))) ? (((_this).f11).get((_dafny.ZERO).minus(_this.f10))) : ((_770_v18).f16)));
          _781_v26 = _nw125;
          let _782_v27;
          _782_v27 = _dafny.Map.Empty.slice().updateUnsafe(_779_v25,_781_v26);
          _782_v27 = (_782_v27).update(_779_v25, _781_v26);
          let _783_v28;
          let _init15 = ((_784_v21) => function (_785_i1) {
            return (_785_i1).minus((_784_v21)[_module.__default.safeIndex(_this.f10, new BigNumber((_784_v21).length))]);
          })(_775_v21);
          let _nw126 = Array((new BigNumber(26)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw126.length); _i0_15++) {
            _nw126[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _783_v28 = _nw126;
          let _index130 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_783_v28).length));
          (_783_v28)[_index130] = (_this.f10).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("dq")).length));
          let _786_v29;
          _786_v29 = _dafny.Seq.of(_744_cf22);
          let _index131 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_783_v28).length));
          (_783_v28)[_index131] = ((((_775_v21)[_module.__default.safeIndex(_this.f10, new BigNumber((_775_v21).length))]).isLessThan((_dafny.ZERO).minus(_this.f10))) ? ((_module.__default.fm10((_781_v26).f15, globalState)).plus(new BigNumber((_786_v29).length))) : (new BigNumber((_778_v24).length)));
        } else {
          let _787___mcc_h5 = (_source16).cf26;
          let _788_cf26 = _787___mcc_h5;
          r2 = !((_744_cf22) && (_744_cf22));
          (globalState).f2 = (new BigNumber(187)).multipliedBy(_this.f10);
          let _789_v30;
          let _nw127 = new _module.C2();
          _nw127.__ctor(_744_cf22, false);
          _789_v30 = _nw127;
          let _790_v31;
          _790_v31 = _dafny.Seq.of((_789_v30).f15);
          let _791_v32;
          let _792_v33;
          let _out56;
          let _out57;
          let _outcollector17 = (_789_v30).m3(_dafny.Seq.Concat(_dafny.Seq.update(_790_v31, _module.__default.safeIndex(_this.f10, new BigNumber((_790_v31).length)), _744_cf22), _dafny.Seq.of(!((_789_v30).f15))), globalState);
          _out56 = _outcollector17[0];
          _out57 = _outcollector17[1];
          _791_v32 = _out56;
          _792_v33 = _out57;
        }
        let _793_v34;
        _793_v34 = _dafny.Seq.of(_744_cf22);
        let _794_v35;
        _794_v35 = _module.D9.create_DC20(_dafny.areEqual(_793_v34, _793_v34), _module.__default.fm26(((((_this).f11).contains(_this.f10)) ? (((_this).f11).get(_this.f10)) : (_744_cf22)), globalState));
        let _795_v36;
        _795_v36 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm27(globalState)).length));
        let _796_v37;
        _796_v37 = new _dafny.CodePoint('f'.codePointAt(0));
        _794_v35 = _module.D9.create_DC20((_dafny.MultiSet.fromElements(_this.f10)).IsSubsetOf(_795_v36), _796_v37);
        let _797_v38;
        _797_v38 = _dafny.Seq.UnicodeFromString("bicxihw");
        let _798_v39;
        _798_v39 = _dafny.Map.Empty.slice().updateUnsafe(_744_cf22,_this.f10);
        let _799_v40;
        _799_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-765)), ((_800_v37) => function (_801_i2) {
          return _800_v37;
        })(_796_v37))).length));
        r1 = (new BigNumber((_dafny.Seq.Concat(_797_v38, _dafny.Seq.update(_797_v38, _module.__default.safeIndex(new BigNumber(617), new BigNumber((_797_v38).length)), _796_v37))).length)).multipliedBy(((((_798_v39).contains(_744_cf22)) ? ((_798_v39).get(_744_cf22)) : (_this.f10))).multipliedBy((((_799_v40).contains(new BigNumber((_795_v36).cardinality()))) ? ((_799_v40).get(new BigNumber((_795_v36).cardinality()))) : (_this.f10))));
      } else {
        let _802___mcc_h1 = (_source15).cf21;
        let _803_cf21 = _802___mcc_h1;
        (globalState).f2 = new BigNumber(447);
        (globalState).f2 = _this.f10;
        let _804_v41;
        _804_v41 = false;
        let _805_v42;
        let _nw128 = Array((new BigNumber(3)).toNumber());
        _nw128[(_dafny.ZERO).toNumber()] = _804_v41;
        _nw128[(_dafny.ONE).toNumber()] = _804_v41;
        _nw128[(new BigNumber(2)).toNumber()] = _804_v41;
        _805_v42 = _nw128;
        let _806_v43;
        let _nw129 = Array((new BigNumber(15)).toNumber());
        _nw129[(_dafny.ZERO).toNumber()] = _805_v42;
        _nw129[(_dafny.ONE).toNumber()] = _805_v42;
        _nw129[(new BigNumber(2)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(3)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(4)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(5)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(6)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(7)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(8)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(9)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(10)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(11)).toNumber()] = (_module.D11.create_DC24(_805_v42)).dtor_cf36;
        _nw129[(new BigNumber(12)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(13)).toNumber()] = _805_v42;
        _nw129[(new BigNumber(14)).toNumber()] = _805_v42;
        _806_v43 = _nw129;
        let _rhs148 = _806_v43;
        let _rhs149 = _this.f10;
        let _lhs98 = globalState;
        _806_v43 = _rhs148;
        _lhs98.f2 = _rhs149;
        let _807_v44;
        let _808_v45;
        let _809_v46;
        let _810_v47;
        let _out58;
        let _out59;
        let _out60;
        let _out61;
        let _outcollector18 = _module.__default.m0(globalState);
        _out58 = _outcollector18[0];
        _out59 = _outcollector18[1];
        _out60 = _outcollector18[2];
        _out61 = _outcollector18[3];
        _807_v44 = _out58;
        _808_v45 = _out59;
        _809_v46 = _out60;
        _810_v47 = _out61;
      }
      let _811_v48;
      _811_v48 = false;
      let _812_v49;
      let _nw130 = new _module.C0();
      _nw130.__ctor(_811_v48, _this.f10);
      _812_v49 = _nw130;
      let _813_v50;
      _813_v50 = _module.D1.create_DC1(_811_v48);
      let _814_v51;
      let _nw131 = new _module.C2();
      _nw131.__ctor((_812_v49).f12, (_812_v49).f12);
      _814_v51 = _nw131;
      let _815_v52;
      _815_v52 = _dafny.Map.Empty.slice().updateUnsafe((_813_v50).dtor_cf1,_814_v51);
      let _816_v53;
      _816_v53 = _dafny.Seq.UnicodeFromString("ntqkq");
      let _817_v54;
      _817_v54 = _dafny.Set.fromElements(new BigNumber((_815_v52).length), (_812_v49).f13, (_812_v49).f13, new BigNumber((_816_v53).length), _this.f10);
      let _818_v55;
      _818_v55 = _dafny.Map.Empty.slice().updateUnsafe(!((_812_v49).f12),new BigNumber((_817_v54).length));
      let _819_v56;
      _819_v56 = _dafny.Seq.of(_818_v55);
      let _820_v57;
      let _nw132 = Array((new BigNumber(13)).toNumber());
      _nw132[(_dafny.ZERO).toNumber()] = _818_v55;
      _nw132[(_dafny.ONE).toNumber()] = _818_v55;
      _nw132[(new BigNumber(2)).toNumber()] = _818_v55;
      _nw132[(new BigNumber(3)).toNumber()] = _818_v55;
      _nw132[(new BigNumber(4)).toNumber()] = _818_v55;
      _nw132[(new BigNumber(5)).toNumber()] = ((_819_v56)[_module.__default.safeIndex((_812_v49).f13, new BigNumber((_819_v56).length))]).Merge(_818_v55);
      _nw132[(new BigNumber(6)).toNumber()] = _818_v55;
      _nw132[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_811_v48,(_812_v49).f13);
      _nw132[(new BigNumber(8)).toNumber()] = (_818_v55).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_811_v48),_this.f10));
      _nw132[(new BigNumber(9)).toNumber()] = (_818_v55).Merge(_818_v55);
      _nw132[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_814_v51).f15,new BigNumber((_816_v53).length));
      _nw132[(new BigNumber(11)).toNumber()] = (_818_v55).Merge(_dafny.Map.Empty.slice().updateUnsafe((_814_v51).f15,(_812_v49).f13));
      _nw132[(new BigNumber(12)).toNumber()] = _818_v55;
      _820_v57 = _nw132;
      let _index132 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_820_v57).length));
      (_820_v57)[_index132] = _818_v55;
      let _821_v58;
      let _nw133 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _821_v58 = _nw133;
      let _822_v59;
      _822_v59 = new _dafny.CodePoint('g'.codePointAt(0));
      let _index133 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_821_v58).length));
      (_821_v58)[_index133] = _822_v59;
      let _823_v60;
      _823_v60 = _module.D9.create_DC19(_818_v55);
      let _pat_let_tv4 = _814_v51;
      let _pat_let_tv5 = _812_v49;
      let _pat_let_tv6 = _812_v49;
      let _index134 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_820_v57).length));
      let _index135 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_821_v58).length));
      let _rhs150 = function (_source17) {
        if (_source17.is_DC20) {
          let _824___mcc_h6 = (_source17).cf29;
          let _825___mcc_h7 = (_source17).cf30;
          let _826_cf30 = _825___mcc_h7;
          let _827_cf29 = _824___mcc_h6;
          return (_pat_let_tv4).f15;
        } else if (_source17.is_DC19) {
          let _828___mcc_h8 = (_source17).cf28;
          let _829_cf28 = _828___mcc_h8;
          return (_pat_let_tv5).f12;
        } else {
          let _830___mcc_h9 = (_source17).cf31;
          let _831_cf31 = _830___mcc_h9;
          return !(_this.f10).isEqualTo((_pat_let_tv6).f13);
        }
      }(((_811_v48) ? (_823_v60) : (_823_v60)));
      let _rhs151 = (_818_v55).Merge(_818_v55);
      let _rhs152 = _822_v59;
      let _rhs153 = ((_this.f10).plus(_this.f10)).minus(_module.__default.fm19(new BigNumber((_816_v53).length), globalState));
      let _lhs99 = _820_v57;
      let _lhs100 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_820_v57).length));
      let _lhs101 = _821_v58;
      let _lhs102 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_821_v58).length));
      let _lhs103 = _this;
      r0 = _rhs150;
      _lhs99[_lhs100] = _rhs151;
      _lhs101[_lhs102] = _rhs152;
      _lhs103.f10 = _rhs153;
      (globalState).f2 = (_812_v49).f13;
      let _832_v61;
      _832_v61 = _dafny.Seq.of(_816_v53, _dafny.Seq.UnicodeFromString("hjimyu"));
      let _833_v62;
      _833_v62 = (_dafny.ZERO).minus((_812_v49).f13);
      let _source18 = _module.D1.create_DC3(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(924)), ((_834_v58) => function (_835_i3) {
  return (_module.D11.create_DC25((_834_v58)[_module.__default.safeIndex(new BigNumber(966), new BigNumber((_834_v58).length))])).dtor_cf37;
})(_821_v58)), (_832_v61)[_module.__default.safeIndex(_this.f10, new BigNumber((_832_v61).length))]), _833_v62);
      if (_source18.is_DC2) {
        let _836___mcc_h10 = (_source18).cf2;
        let _837_cf2 = _836___mcc_h10;
        _811_v48 = (_812_v49).f12;
        (globalState).f2 = (((_812_v49).f12) ? ((_812_v49).f13) : ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_812_v49).f13,_814_v51.f14)).length)).minus((_812_v49).f13)));
        (globalState).f2 = (_812_v49).f13;
        let _838_v63;
        _838_v63 = _dafny.Seq.of((_814_v51).f15, _814_v51.f14, (_814_v51).f15, _811_v48);
        let _839_v64;
        let _nw134 = new _module.C2();
        _nw134.__ctor(!_dafny.areEqual(_838_v63, _838_v63), (_812_v49).f12);
        _839_v64 = _nw134;
      } else if (_source18.is_DC3) {
        let _840___mcc_h11 = (_source18).cf3;
        let _841___mcc_h12 = (_source18).cf4;
        let _842_cf4 = _841___mcc_h12;
        let _843_cf3 = _840___mcc_h11;
        let _844_v65;
        let _nw135 = new _module.C0();
        _nw135.__ctor(_843_cf3, _this.f10);
        _844_v65 = _nw135;
        let _845_v66;
        _845_v66 = _module.D10.create_DC23((_812_v49).f13, (_812_v49).f12, _843_cf3);
        _845_v66 = _845_v66;
        _822_v59 = _822_v59;
        _811_v48 = _843_cf3;
      } else {
        let _846___mcc_h13 = (_source18).cf1;
        let _847_cf1 = _846___mcc_h13;
        let _848_v67;
        _848_v67 = _dafny.Set.fromElements((_812_v49).f12);
        (_814_v51).f14 = (!(_847_cf1)) || ((_848_v67).IsDisjointFrom(_848_v67));
        r3 = !((_814_v51).f15);
        let _849_v68;
        let _init16 = ((_850_v49) => function (_851_i4) {
          return !((_850_v49).f12);
        })(_812_v49);
        let _nw136 = Array((new BigNumber(21)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw136.length); _i0_16++) {
          _nw136[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _849_v68 = _nw136;
        _849_v68 = _849_v68;
        let _852_v69;
        let _853_v70;
        let _854_v71;
        let _855_v72;
        let _out62;
        let _out63;
        let _out64;
        let _out65;
        let _outcollector19 = _module.__default.m0(globalState);
        _out62 = _outcollector19[0];
        _out63 = _outcollector19[1];
        _out64 = _outcollector19[2];
        _out65 = _outcollector19[3];
        _852_v69 = _out62;
        _853_v70 = _out63;
        _854_v71 = _out64;
        _855_v72 = _out65;
      }
      (_814_v51).f14 = (_814_v51).f15;
      r0 = ((_812_v49).f13).isLessThanOrEqualTo((_812_v49).f13);
      r1 = (_812_v49).f13;
      r2 = _811_v48;
      let _856_v73;
      _856_v73 = _dafny.MultiSet.fromElements(_811_v48);
      r3 = (_814_v51).fm13(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_814_v51.f14), _dafny.Seq.update(_dafny.Seq.of(_811_v48, !(_814_v51.f14), !((_812_v49).f12), (_814_v51).f15), _module.__default.safeIndex((_dafny.ZERO).minus((_812_v49).f13), new BigNumber((_dafny.Seq.of(_811_v48, !(_814_v51.f14), !((_812_v49).f12), (_814_v51).f15)).length)), (_814_v51).f15))).length), (_856_v73).IsSubsetOf(_856_v73), new BigNumber(22), globalState);
      return [r0, r1, r2, r3];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
