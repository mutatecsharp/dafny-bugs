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
    static fm0(p0, p1, p2, globalState) {
      if (!(false)) {
        return !(true);
      } else {
        return (new BigNumber(761)).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("m")).length));
      }
    };
    static fm1(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vvsslon"), _dafny.Seq.UnicodeFromString("iphflypt")), _dafny.Seq.UnicodeFromString("qc"));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false, true)).length)), new BigNumber(-55), new BigNumber((_dafny.Seq.UnicodeFromString("suxyudnjw")).length))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length))).cardinality()))).length)))).Difference((_dafny.MultiSet.fromElements(new BigNumber(367))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(265)))));
    };
    static fm3(p0, globalState) {
      return (_dafny.ZERO).minus(new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(676)), function (_0_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("byditj")))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_module.D1.create_DC4(_dafny.Seq.of(true, true), false)).dtor_cf7,_dafny.Seq.UnicodeFromString("s")))).length));
    };
    static fm6(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true));
    };
    static fm7(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(619)), function (_1_i0) {
        return (new BigNumber((_dafny.Seq.UnicodeFromString("xnuinsd")).length)).minus(new BigNumber((_dafny.Seq.of(true)).length));
      });
    };
    static fm8(p0, p1, p2, globalState) {
      return _module.D1.create_DC4(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true)), (_dafny.Set.fromElements(new BigNumber(418), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-592))).cardinality()), new BigNumber((_dafny.Seq.of(true, false, true)).length))).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((function () {
  let _coll0 = new _dafny.Set();
  for (const _compr_0 of _dafny.IntegerRange(new BigNumber(604), new BigNumber(391))) {
    let _2_v0 = _compr_0;
    if (((new BigNumber(604)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(391)))) {
      _coll0.add((_2_v0).plus(new BigNumber(262)));
    }
  }
  return _coll0;
}()).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(!(true))).length))).cardinality()), new BigNumber(893))));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return (function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(-274), new BigNumber(697))) {
          let _3_v0 = _compr_1;
          if (((new BigNumber(-274)).isLessThanOrEqualTo(_3_v0)) && ((_3_v0).isLessThan(new BigNumber(697)))) {
            _coll1.push([_module.__default.safeModuloInt(_3_v0, new BigNumber(59)),false]);
          }
        }
        return _coll1;
      }());
    };
    static fm10(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(464),false))).Intersect(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("x")).length),new BigNumber((_dafny.Seq.UnicodeFromString("aqb")).length))).length),!(true))))).Difference((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(965),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true, false))).cardinality())),true)).length),true), function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(237), new BigNumber(872))) {
          let _4_v0 = _compr_2;
          if (((new BigNumber(237)).isLessThanOrEqualTo(_4_v0)) && ((_4_v0).isLessThan(new BigNumber(872)))) {
            _coll2.push([(_4_v0).minus(new BigNumber(-833)),true]);
          }
        }
        return _coll2;
      }())).Union(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-32),false))));
    };
    static fm11(p0, p1, globalState) {
      return (_module.D8.create_DC22(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('j'.codePointAt(0))))).dtor_cf31;
    };
    static fm12(p0, p1, globalState) {
      return _module.D0.create_DC0(new _dafny.CodePoint('m'.codePointAt(0)));
    };
    static fm13(p0, globalState) {
      return new _dafny.CodePoint('b'.codePointAt(0));
    };
    static fm14(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("buhcp"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("lqxqtv")));
    };
    static fm15(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-81)), function (_5_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("rqow")).length);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(496)), function (_6_i1) {
        return new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(617),new BigNumber(128)))).length);
      }))).length),new BigNumber(-39));
    };
    static m0(p0, p1, p2, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _7_v0;
      _7_v0 = _dafny.Seq.UnicodeFromString("dfhmvgywh");
      let _hi0 = (new BigNumber((_dafny.MultiSet.fromElements(_7_v0, _dafny.Seq.UnicodeFromString("yet"), _7_v0)).cardinality())).multipliedBy(new BigNumber((_module.__default.fm1(globalState)).length));
      for (let _8_i0 = (_dafny.ZERO).minus(p2); _8_i0.isLessThan(_hi0); _8_i0 = _8_i0.plus(_dafny.ONE)) {
        if (false) {
          let _9_v1;
          let _nw0 = new _module.C0();
          _nw0.__ctor(p0);
          _9_v1 = _nw0;
          let _10_v2;
          _10_v2 = _dafny.Seq.of(p1);
          let _11_v3;
          _11_v3 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm11((_9_v1).f24, (_dafny.ZERO).minus(_8_i0), globalState)).length), _8_i0, _8_i0);
          let _12_v4;
          _12_v4 = _dafny.Map.Empty.slice().updateUnsafe(_8_i0,_11_v3);
          r1 = (_9_v1).fm4((_module.D1.create_DC4(_10_v2, p0)).dtor_cf7, p0, ((_module.__default.fm0(p2, (_9_v1).f24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(337)), ((_13_i0) => function (_14_i1) {
            return (_module.D3.create_DC10(_13_i0, new _dafny.CodePoint('m'.codePointAt(0)))).dtor_cf17;
          })(_8_i0)), globalState)) ? (_11_v3) : ((((_12_v4).contains(_8_i0)) ? ((_12_v4).get(_8_i0)) : (_11_v3)))), _module.__default.fm0(p2, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(532)), function (_15_i2) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }), globalState), globalState);
          let _16_v5;
          _16_v5 = _dafny.Seq.of(_9_v1);
          let _17_v6;
          _17_v6 = _dafny.Seq.of(_16_v5);
          let _18_v7;
          _18_v7 = _dafny.Set.fromElements(_dafny.Seq.of(_9_v1), (_17_v6)[_module.__default.safeIndex(_8_i0, new BigNumber((_17_v6).length))], _16_v5, _16_v5, _16_v5);
          let _19_v8;
          let _nw1 = Array((new BigNumber(17)).toNumber());
          _nw1[(_dafny.ZERO).toNumber()] = _18_v7;
          _nw1[(_dafny.ONE).toNumber()] = _18_v7;
          _nw1[(new BigNumber(2)).toNumber()] = _18_v7;
          _nw1[(new BigNumber(3)).toNumber()] = (_dafny.Set.fromElements(_16_v5)).Difference(_18_v7);
          _nw1[(new BigNumber(4)).toNumber()] = _18_v7;
          _nw1[(new BigNumber(5)).toNumber()] = _18_v7;
          _nw1[(new BigNumber(6)).toNumber()] = (_18_v7).Difference(_18_v7);
          _nw1[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_16_v5);
          _nw1[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_16_v5);
          _nw1[(new BigNumber(9)).toNumber()] = (_18_v7).Union(_18_v7);
          _nw1[(new BigNumber(10)).toNumber()] = _18_v7;
          _nw1[(new BigNumber(11)).toNumber()] = (_18_v7).Union(_18_v7);
          _nw1[(new BigNumber(12)).toNumber()] = _18_v7;
          _nw1[(new BigNumber(13)).toNumber()] = _18_v7;
          _nw1[(new BigNumber(14)).toNumber()] = (_dafny.Set.fromElements(_16_v5)).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(_9_v1)));
          _nw1[(new BigNumber(15)).toNumber()] = (_18_v7).Union(_18_v7);
          _nw1[(new BigNumber(16)).toNumber()] = (_18_v7).Intersect(_18_v7);
          _19_v8 = _nw1;
          let _index0 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_19_v8).length));
          (_19_v8)[_index0] = (_18_v7).Union(_dafny.Set.fromElements(_16_v5));
          let _index1 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_19_v8).length));
          (_19_v8)[_index1] = _18_v7;
          (globalState).f23 = new BigNumber(-419);
          _9_v1 = _9_v1;
        } else {
          let _20_v9;
          _20_v9 = _dafny.Seq.of(p2, p2);
          r1 = (_module.__default.safeModuloInt((_20_v9)[_module.__default.safeIndex(p2, new BigNumber((_20_v9).length))], _8_i0)).multipliedBy(new BigNumber(-624));
          let _21_v10;
          _21_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,!(p0) || (false));
          _21_v10 = _21_v10;
          (globalState).f2 = p0;
          let _22_v11;
          _22_v11 = new _dafny.CodePoint('q'.codePointAt(0));
          let _23_v12;
          _23_v12 = _module.D0.create_DC0(_22_v11);
          _23_v12 = ((p1) ? (_23_v12) : (_module.__default.fm12(p1, true, globalState)));
          let _24_v13;
          _24_v13 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber(76)));
          let _25_v15;
          _25_v15 = _dafny.Map.Empty.slice().updateUnsafe(p1,_8_i0);
          let _26_v16;
          _26_v16 = _dafny.MultiSet.fromElements(_module.__default.fm13(new BigNumber((_25_v15).length), globalState), _22_v11);
          let _27_v17;
          _27_v17 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll3 = new _dafny.Map();
            for (const _compr_3 of (_26_v16).Elements) {
              let _28_v14 = _compr_3;
              if ((_26_v16).contains(_28_v14)) {
                _coll3.push([_28_v14,_8_i0]);
              }
            }
            return _coll3;
          }(),p1);
          let _rhs0 = true;
          let _rhs1 = _dafny.Seq.Concat((_24_v13)[_module.__default.safeIndex(p2, new BigNumber((_24_v13).length))], _dafny.Seq.update(_dafny.Seq.of(_8_i0, new BigNumber(701), new BigNumber((_27_v17).length)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.of(_8_i0, new BigNumber(701), new BigNumber((_27_v17).length))).length)), p2));
          let _lhs0 = globalState;
          _lhs0.f18 = _rhs0;
          _20_v9 = _rhs1;
        }
        let _29_v18;
        let _nw2 = new _module.C0();
        _nw2.__ctor(p0);
        _29_v18 = _nw2;
        if (p0) {
          (globalState).f18 = p1;
          _7_v0 = _dafny.Seq.Concat(_module.__default.fm1(globalState), _7_v0);
          let _nw3 = new _module.C0();
          _nw3.__ctor((p1) || (p0));
          _29_v18 = _nw3;
          let _30_v19;
          let _nw4 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _30_v19 = _nw4;
          let _index2 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_30_v19).length));
          (_30_v19)[_index2] = (_dafny.ZERO).minus((_8_i0).plus(_8_i0));
          let _31_v20;
          _31_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_8_i0),p1);
          let _32_v21;
          _32_v21 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_7_v0).length))).length), new BigNumber((_31_v20).length));
          let _33_v22;
          _33_v22 = _module.D4.create_DC13(_32_v21);
          let _34_v23;
          let _nw5 = Array((new BigNumber(5)).toNumber());
          _nw5[(_dafny.ZERO).toNumber()] = _32_v21;
          _nw5[(_dafny.ONE).toNumber()] = _32_v21;
          _nw5[(new BigNumber(2)).toNumber()] = (_33_v22).dtor_cf20;
          _nw5[(new BigNumber(3)).toNumber()] = _32_v21;
          _nw5[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(799)), ((_35_p2) => function (_36_i3) {
            return _35_p2;
          })(p2));
          _34_v23 = _nw5;
          let _index3 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_34_v23).length));
          (_34_v23)[_index3] = _dafny.Seq.of(_8_i0);
          let _37_v24;
          _37_v24 = _dafny.Seq.of((_29_v18).f24);
          let _38_v25;
          _38_v25 = _dafny.Set.fromElements(_37_v24);
          let _index4 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_30_v19).length));
          let _index5 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_34_v23).length));
          let _rhs2 = new BigNumber(((_38_v25).Intersect(_38_v25)).length);
          let _rhs3 = _dafny.Seq.Concat(_32_v21, _32_v21);
          let _lhs1 = _30_v19;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_30_v19).length));
          let _lhs3 = _34_v23;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_34_v23).length));
          _lhs1[_lhs2] = _rhs2;
          _lhs3[_lhs4] = _rhs3;
          let _index6 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_30_v19).length));
          (_30_v19)[_index6] = (_30_v19)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_30_v19).length))];
        } else {
          let _39_v26;
          _39_v26 = _dafny.MultiSet.fromElements(p1, p1);
          r0 = (p2).multipliedBy((new BigNumber((_39_v26).cardinality())).multipliedBy(_8_i0));
          let _40_v27;
          let _nw6 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _40_v27 = _nw6;
          let _index7 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_40_v27).length));
          (_40_v27)[_index7] = p2;
          let _41_v28;
          _41_v28 = _dafny.Map.Empty.slice().updateUnsafe(!(_8_i0).isEqualTo(_8_i0),_7_v0);
          let _42_v29;
          _42_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(648));
          let _43_v30;
          _43_v30 = _dafny.Seq.of(new BigNumber((_42_v29).length), p2);
          let _44_v31;
          _44_v31 = new _dafny.CodePoint('t'.codePointAt(0));
          let _index8 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_40_v27).length));
          let _rhs4 = (_29_v18).fm4((_29_v18).f24, !(_dafny.Seq.IsPrefixOf(_43_v30, _43_v30)), _dafny.MultiSet.fromElements(_8_i0, p2, _8_i0, _8_i0), (_29_v18).f24, globalState);
          let _rhs5 = p0;
          let _rhs6 = p0;
          let _rhs7 = _module.__default.fm14(_44_v31, globalState);
          let _lhs5 = _40_v27;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_40_v27).length));
          let _lhs7 = globalState;
          let _lhs8 = globalState;
          _lhs5[_lhs6] = _rhs4;
          _lhs7.f2 = _rhs5;
          _lhs8.f2 = _rhs6;
          _41_v28 = _rhs7;
          let _45_v32;
          _45_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-384)), ((_46_i0) => function (_47_i4) {
            return _46_i0;
          })(_8_i0))).length),p2);
          let _48_v33;
          _48_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_29_v18).f24, p0),!(true));
          let _49_v34;
          _49_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_45_v32).length),!((((_48_v33).contains(_dafny.Seq.of(p1))) ? ((_48_v33).get(_dafny.Seq.of(p1))) : ((_29_v18).f24))));
          let _50_v35;
          _50_v35 = _dafny.MultiSet.fromElements((_49_v34).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,p1)));
          let _index9 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_40_v27).length));
          (_40_v27)[_index9] = (((_50_v35).contains(_module.__default.fm9((_dafny.ZERO).minus(new BigNumber((_39_v26).cardinality())), (_29_v18).f24, p2, p1, globalState))) ? ((_50_v35).get(_module.__default.fm9((_dafny.ZERO).minus(new BigNumber((_39_v26).cardinality())), (_29_v18).f24, p2, p1, globalState))) : (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_8_i0), _8_i0)));
          let _51_v36;
          _51_v36 = _module.D3.create_DC10(p2, _44_v31);
          let _52_v37;
          _52_v37 = _dafny.Map.Empty.slice().updateUnsafe(_29_v18,_51_v36);
          let _53_v38;
          let _nw7 = Array((new BigNumber(25)).toNumber()).fill([]);
          _53_v38 = _nw7;
          let _54_v39;
          let _nw8 = Array((new BigNumber(16)).toNumber()).fill([]);
          _54_v39 = _nw8;
          let _index10 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_53_v38).length));
          (_53_v38)[_index10] = _54_v39;
          let _55_v40;
          _55_v40 = _module.D5.create_DC15(_54_v39);
          let _56_v41;
          _56_v41 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(422)), ((_57_v31) => function (_58_i5) {
            return _57_v31;
          })(_44_v31)));
          let _index11 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_53_v38).length));
          let _rhs8 = ((_52_v37).Merge(_52_v37)).update(_29_v18, _51_v36);
          let _rhs9 = (_55_v40).dtor_cf22;
          let _rhs10 = (_56_v41)[_module.__default.safeIndex((_40_v27)[_module.__default.safeIndex(new BigNumber(365), new BigNumber((_40_v27).length))], new BigNumber((_56_v41).length))];
          let _lhs9 = _53_v38;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_53_v38).length));
          _52_v37 = _rhs8;
          _lhs9[_lhs10] = _rhs9;
          _7_v0 = _rhs10;
          _29_v18 = _29_v18;
        }
        let _59_v42;
        _59_v42 = _module.D3.create_DC11((_29_v18).f24);
        let _60_v43;
        let _nw9 = new _module.C0();
        _nw9.__ctor((new BigNumber(422)).isLessThan(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Set.fromElements((_59_v42).dtor_cf18, p0))).cardinality())));
        _60_v43 = _nw9;
      }
      (globalState).f18 = _module.__default.fm0(p2, p1, _dafny.Seq.Concat(_7_v0, _7_v0), globalState);
      let _hi1 = p2;
      for (let _61_i6 = p2; _61_i6.isLessThan(_hi1); _61_i6 = _61_i6.plus(_dafny.ONE)) {
        let _62_v44;
        _62_v44 = new _dafny.CodePoint('h'.codePointAt(0));
        let _63_v45;
        let _nw10 = new _module.C0();
        _nw10.__ctor(p1);
        _63_v45 = _nw10;
        let _64_v46;
        _64_v46 = _dafny.Map.Empty.slice().updateUnsafe(_63_v45,(_63_v45).f24);
        let _65_v47;
        _65_v47 = _dafny.Map.Empty.slice().updateUnsafe(_61_i6,(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_64_v46).length),(_63_v45).f24)).update(_61_i6, p0));
        let _66_v48;
        _66_v48 = _module.D6.create_DC18(_65_v47);
        let _67_v49;
        _67_v49 = _dafny.Map.Empty.slice().updateUnsafe((_64_v46).Merge(_64_v46),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-676)), ((_68_v44) => function (_69_i7) {
          return _68_v44;
        })(_62_v44)));
        let _rhs11 = _62_v44;
        let _rhs12 = (_66_v48).dtor_cf26;
        let _rhs13 = (((_67_v49).contains(_dafny.Map.Empty.slice().updateUnsafe(_63_v45,(_63_v45).f24))) ? ((_67_v49).get(_dafny.Map.Empty.slice().updateUnsafe(_63_v45,(_63_v45).f24))) : (_dafny.Seq.Concat(_7_v0, _7_v0)));
        let _rhs14 = _63_v45;
        let _rhs15 = !(_61_i6).isEqualTo(new BigNumber(153));
        let _lhs11 = globalState;
        _62_v44 = _rhs11;
        _65_v47 = _rhs12;
        _7_v0 = _rhs13;
        _63_v45 = _rhs14;
        _lhs11.f18 = _rhs15;
        r1 = _61_i6;
        let _70_v50;
        _70_v50 = _dafny.Seq.of(p1);
        let _71_v51;
        _71_v51 = _module.D1.create_DC4(_70_v50, !(p0));
        let _72_v52;
        _72_v52 = _dafny.Seq.of(_7_v0);
        let _73_v54;
        _73_v54 = _dafny.Seq.of(new BigNumber((_70_v50).length), p2);
        let _74_v55;
        _74_v55 = _dafny.Map.Empty.slice().updateUnsafe((_63_v45).fm4(p0, (_63_v45).f24, _dafny.MultiSet.FromArray(_73_v54), p1, globalState),_61_i6);
        let _rhs16 = p0;
        let _rhs17 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber((_72_v52).length), _61_i6), new BigNumber((((p0) ? (function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of (_module.__default.fm15(_61_i6, p0, p1, globalState)).Keys.Elements) {
            let _75_v53 = _compr_4;
            if ((_module.__default.fm15(_61_i6, p0, p1, globalState)).contains(_75_v53)) {
              _coll4.push([(_75_v53).multipliedBy(_61_i6),p2]);
            }
          }
          return _coll4;
        }()) : (_74_v55))).length));
        let _rhs18 = _71_v51;
        let _lhs12 = globalState;
        let _lhs13 = globalState;
        _lhs12.f18 = _rhs16;
        _lhs13.f23 = _rhs17;
        _71_v51 = _rhs18;
        r1 = (_dafny.ZERO).minus((_61_i6).minus(p2));
      }
      let _76_i8;
      _76_i8 = _dafny.ZERO;
      L0: {
        while (_module.__default.fm0((p2).plus(p2), p1, _7_v0, globalState)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_76_i8)) {
              break L0;
            }
            _76_i8 = (_76_i8).plus(_dafny.ONE);
            (globalState).f18 = p1;
            let _77_v56;
            _77_v56 = _module.D6.create_DC20();
            _77_v56 = _77_v56;
            let _78_v57;
            let _nw11 = Array((new BigNumber(25)).toNumber()).fill([]);
            _78_v57 = _nw11;
            let _79_v58;
            let _nw12 = Array((new BigNumber(9)).toNumber()).fill(false);
            _79_v58 = _nw12;
            let _index12 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_78_v57).length));
            (_78_v57)[_index12] = _79_v58;
            let _index13 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_78_v57).length));
            let _nw13 = Array((new BigNumber(23)).toNumber()).fill(false);
            (_78_v57)[_index13] = _nw13;
            (globalState).f2 = p0;
          }
        }
      }
      let _80_v59;
      let _nw14 = Array((new BigNumber(5)).toNumber()).fill([]);
      _80_v59 = _nw14;
      let _81_v60;
      _81_v60 = new _dafny.CodePoint('v'.codePointAt(0));
      let _82_v61;
      let _nw15 = Array((new BigNumber(9)).toNumber());
      _nw15[(_dafny.ZERO).toNumber()] = _7_v0;
      _nw15[(_dafny.ONE).toNumber()] = _7_v0;
      _nw15[(new BigNumber(2)).toNumber()] = _7_v0;
      _nw15[(new BigNumber(3)).toNumber()] = _7_v0;
      _nw15[(new BigNumber(4)).toNumber()] = _7_v0;
      _nw15[(new BigNumber(5)).toNumber()] = _7_v0;
      _nw15[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("c");
      _nw15[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("krdlpcyb");
      _nw15[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("sj"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("sj")).length)), _81_v60);
      _82_v61 = _nw15;
      let _index14 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_80_v59).length));
      (_80_v59)[_index14] = _82_v61;
      let _index15 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_80_v59).length));
      (_80_v59)[_index15] = _82_v61;
      (globalState).f2 = p1;
      let _83_v62;
      let _init0 = ((_84_p0) => function (_85_i9) {
        return _84_p0;
      })(p0);
      let _nw16 = Array((new BigNumber(9)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw16.length); _i0_0++) {
        _nw16[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _83_v62 = _nw16;
      let _86_v63;
      _86_v63 = _dafny.MultiSet.fromElements(_83_v62, _83_v62);
      r0 = new BigNumber(((_86_v63).Difference((_86_v63).update(_83_v62, _module.__default.abs(new BigNumber(177))))).cardinality());
      let _87_v64;
      _87_v64 = _dafny.Set.fromElements(_81_v60, _81_v60);
      r1 = (p2).plus((_dafny.ZERO).minus(new BigNumber(((_87_v64).Difference(function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (_87_v64).Elements) {
          let _88_v65 = _compr_5;
          if ((_87_v64).contains(_88_v65)) {
            _coll5.add(_88_v65);
          }
        }
        return _coll5;
      }())).length)));
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _89_v0;
      _89_v0 = _dafny.Seq.UnicodeFromString("spiw");
      let _90_v1;
      _90_v1 = _dafny.MultiSet.fromElements(_89_v0, _89_v0, _89_v0, _89_v0, _89_v0);
      let _91_v2;
      _91_v2 = true;
      let _92_v3;
      _92_v3 = new BigNumber(56);
      let _93_v4;
      _93_v4 = _dafny.Map.Empty.slice().updateUnsafe(_91_v2,_92_v3);
      let _94_v5;
      _94_v5 = _dafny.Map.Empty.slice().updateUnsafe(_93_v4,_92_v3);
      let _95_v6;
      let _nw17 = Array((new BigNumber(16)).toNumber()).fill([]);
      _95_v6 = _nw17;
      let _96_v7;
      _96_v7 = _dafny.Map.Empty.slice().updateUnsafe(_91_v2,_95_v6);
      let _97_v8;
      let _nw18 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _97_v8 = _nw18;
      let _98_v9;
      _98_v9 = new _dafny.CodePoint('n'.codePointAt(0));
      let _99_v10;
      _99_v10 = _dafny.Map.Empty.slice().updateUnsafe(_91_v2,_98_v9);
      let _100_v11;
      _100_v11 = _dafny.Map.Empty.slice().updateUnsafe(_92_v3,_91_v2);
      let _101_v12;
      let _nw19 = Array((new BigNumber(16)).toNumber());
      _nw19[(_dafny.ZERO).toNumber()] = _98_v9;
      _nw19[(_dafny.ONE).toNumber()] = _98_v9;
      _nw19[(new BigNumber(2)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(3)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
      _nw19[(new BigNumber(5)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(6)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(7)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(8)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(9)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(10)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(11)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(12)).toNumber()] = (((_99_v10).contains((((_100_v11).contains(_92_v3)) ? ((_100_v11).get(_92_v3)) : (_91_v2)))) ? ((_99_v10).get((((_100_v11).contains(_92_v3)) ? ((_100_v11).get(_92_v3)) : (_91_v2)))) : (_98_v9));
      _nw19[(new BigNumber(13)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(14)).toNumber()] = _98_v9;
      _nw19[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
      _101_v12 = _nw19;
      let _102_v13;
      _102_v13 = _dafny.Seq.of(_92_v3, _92_v3);
      let _103_globalState;
      let _nw20 = new _module.GlobalState();
      _nw20.__ctor(new BigNumber(580), _90_v1, false, _94_v5, new BigNumber(548), true, true, new BigNumber(376), new BigNumber(87), true, (((_96_v7).contains(false)) ? ((_96_v7).get(false)) : (_95_v6)), new _dafny.CodePoint('w'.codePointAt(0)), new BigNumber(-257), _97_v8, _dafny.Seq.Concat(_89_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-26)), function (_104_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })), _101_v12, new BigNumber(-272), _101_v12, false, true, _dafny.Seq.UnicodeFromString("dblg"), _102_v13, _102_v13, new BigNumber(-128));
      _103_globalState = _nw20;
      let _105_v14;
      let _106_v15;
      let _out0;
      let _out1;
      let _outcollector0 = _module.__default.m0(_91_v2, _91_v2, _92_v3, _103_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _105_v14 = _out0;
      _106_v15 = _out1;
      let _107_v16;
      _107_v16 = _dafny.MultiSet.fromElements(_91_v2, _91_v2);
      if ((_107_v16).IsProperSubsetOf(_107_v16)) {
        let _108_v17;
        let _109_v18;
        let _out2;
        let _out3;
        let _outcollector1 = _module.__default.m0(_91_v2, (_105_v14).isLessThanOrEqualTo(_92_v3), (_106_v15).minus(_92_v3), _103_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _108_v17 = _out2;
        _109_v18 = _out3;
        if (_91_v2) {
          let _110_v19;
          let _init1 = function (_111_i1) {
            return true;
          };
          let _nw21 = Array((new BigNumber(21)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw21.length); _i0_1++) {
            _nw21[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _110_v19 = _nw21;
          let _index16 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_110_v19).length));
          (_110_v19)[_index16] = _91_v2;
          let _112_v20;
          _112_v20 = _dafny.Seq.of(_91_v2, _91_v2);
          let _index17 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_110_v19).length));
          (_110_v19)[_index17] = (_112_v20)[_module.__default.safeIndex(_109_v18, new BigNumber((_112_v20).length))];
          let _113_v21;
          _113_v21 = _dafny.Map.Empty.slice().updateUnsafe(_91_v2,(_110_v19)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_110_v19).length))]);
          let _114_v22;
          let _nw22 = Array((new BigNumber(12)).toNumber());
          _nw22[(_dafny.ZERO).toNumber()] = _113_v21;
          _nw22[(_dafny.ONE).toNumber()] = (_113_v21).update(_91_v2, (_110_v19)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_110_v19).length))]);
          _nw22[(new BigNumber(2)).toNumber()] = _113_v21;
          _nw22[(new BigNumber(3)).toNumber()] = _113_v21;
          _nw22[(new BigNumber(4)).toNumber()] = (_113_v21).update((_110_v19)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_110_v19).length))], _91_v2);
          _nw22[(new BigNumber(5)).toNumber()] = _113_v21;
          _nw22[(new BigNumber(6)).toNumber()] = _113_v21;
          _nw22[(new BigNumber(7)).toNumber()] = _113_v21;
          _nw22[(new BigNumber(8)).toNumber()] = (_113_v21).Merge(_113_v21);
          _nw22[(new BigNumber(9)).toNumber()] = _113_v21;
          _nw22[(new BigNumber(10)).toNumber()] = (_113_v21).Merge(_113_v21);
          _nw22[(new BigNumber(11)).toNumber()] = _113_v21;
          _114_v22 = _nw22;
          _114_v22 = (((_110_v19)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_110_v19).length))]) ? (_114_v22) : (((_91_v2) ? (_114_v22) : (_114_v22))));
          _100_v11 = (_100_v11).update(_105_v14, _91_v2);
          let _115_v23;
          _115_v23 = _module.D0.create_DC0(_98_v9);
          let _index18 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_110_v19).length));
          (_110_v19)[_index18] = _module.__default.fm0((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_92_v3, _92_v3)), !(_dafny.Seq.IsPrefixOf(_dafny.Seq.update(_module.__default.fm1(_103_globalState), _module.__default.safeIndex(new BigNumber((_89_v0).length), new BigNumber((_module.__default.fm1(_103_globalState)).length)), (_115_v23).dtor_cf0), _dafny.Seq.UnicodeFromString("d"))), _89_v0, _103_globalState);
          let _116_v24;
          let _117_v25;
          let _out4;
          let _out5;
          let _outcollector2 = _module.__default.m0((_110_v19)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_110_v19).length))], _91_v2, _109_v18, _103_globalState);
          _out4 = _outcollector2[0];
          _out5 = _outcollector2[1];
          _116_v24 = _out4;
          _117_v25 = _out5;
        } else {
          _105_v14 = _92_v3;
          let _118_v26;
          _118_v26 = _module.D1.create_DC2(_91_v2);
          (_103_globalState).f18 = _module.__default.fm0(_108_v17, (_118_v26).dtor_cf2, _89_v0, _103_globalState);
          (_103_globalState).f23 = _109_v18;
          let _119_v27;
          _119_v27 = _dafny.Seq.of(_91_v2, _91_v2);
          let _120_v28;
          _120_v28 = _dafny.Map.Empty.slice().updateUnsafe(_91_v2,!(false));
          let _121_v29;
          let _nw23 = Array((new BigNumber(17)).toNumber());
          _nw23[(_dafny.ZERO).toNumber()] = !(_91_v2) || (_91_v2);
          _nw23[(_dafny.ONE).toNumber()] = false;
          _nw23[(new BigNumber(2)).toNumber()] = _91_v2;
          _nw23[(new BigNumber(3)).toNumber()] = _91_v2;
          _nw23[(new BigNumber(4)).toNumber()] = _91_v2;
          _nw23[(new BigNumber(5)).toNumber()] = !(_dafny.Seq.IsProperPrefixOf(_89_v0, _89_v0));
          _nw23[(new BigNumber(6)).toNumber()] = (_105_v14).isEqualTo(new BigNumber(-375));
          _nw23[(new BigNumber(7)).toNumber()] = _91_v2;
          _nw23[(new BigNumber(8)).toNumber()] = !(_105_v14).isEqualTo(_105_v14);
          _nw23[(new BigNumber(9)).toNumber()] = (_91_v2) && (_91_v2);
          _nw23[(new BigNumber(10)).toNumber()] = true;
          _nw23[(new BigNumber(11)).toNumber()] = _module.__default.fm0(_105_v14, _91_v2, _dafny.Seq.UnicodeFromString("n"), _103_globalState);
          _nw23[(new BigNumber(12)).toNumber()] = _91_v2;
          _nw23[(new BigNumber(13)).toNumber()] = _91_v2;
          _nw23[(new BigNumber(14)).toNumber()] = _91_v2;
          _nw23[(new BigNumber(15)).toNumber()] = true;
          _nw23[(new BigNumber(16)).toNumber()] = (_119_v27)[_module.__default.safeIndex(new BigNumber((_120_v28).length), new BigNumber((_119_v27).length))];
          _121_v29 = _nw23;
          let _rhs19 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(952)), ((_122_v9) => function (_123_i2) {
            return _122_v9;
          })(_98_v9));
          let _rhs20 = _89_v0;
          let _rhs21 = _121_v29;
          let _rhs22 = !((_dafny.ZERO).minus(_92_v3)).isEqualTo(_109_v18);
          let _rhs23 = (new BigNumber((_dafny.Seq.Concat(_89_v0, _dafny.Seq.UnicodeFromString("nboyosu"))).length)).plus((_dafny.ZERO).minus(_109_v18));
          let _lhs14 = _103_globalState;
          let _lhs15 = _103_globalState;
          _89_v0 = _rhs19;
          _lhs14.f20 = _rhs20;
          _121_v29 = _rhs21;
          _lhs15.f2 = _rhs22;
          _105_v14 = _rhs23;
          let _index19 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_121_v29).length));
          (_121_v29)[_index19] = !(_91_v2);
          let _index20 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_121_v29).length));
          (_121_v29)[_index20] = _91_v2;
        }
        let _124_v30;
        _124_v30 = _dafny.Set.fromElements(_108_v17);
        let _125_v31;
        _125_v31 = _dafny.MultiSet.fromElements(_109_v18, new BigNumber((_dafny.Seq.UnicodeFromString("nfrvenm")).length), _92_v3);
        let _126_v32;
        let _127_v33;
        let _out6;
        let _out7;
        let _outcollector3 = _module.__default.m0(_91_v2, (_module.__default.fm2(_124_v30, _106_v15, _105_v14, new BigNumber(916), _103_globalState)).IsDisjointFrom(_125_v31), _module.__default.fm3(_106_v15, _103_globalState), _103_globalState);
        _out6 = _outcollector3[0];
        _out7 = _outcollector3[1];
        _126_v32 = _out6;
        _127_v33 = _out7;
        _109_v18 = _106_v15;
        if (_91_v2) {
          let _128_v34;
          let _129_v35;
          let _out8;
          let _out9;
          let _outcollector4 = _module.__default.m0(false, _91_v2, _module.__default.safeModuloInt(_109_v18, _105_v14), _103_globalState);
          _out8 = _outcollector4[0];
          _out9 = _outcollector4[1];
          _128_v34 = _out8;
          _129_v35 = _out9;
          let _130_v36;
          _130_v36 = _module.D1.create_DC2(_91_v2);
          let _131_v37;
          let _nw24 = Array((new BigNumber(26)).toNumber()).fill(false);
          _131_v37 = _nw24;
          let _132_v38;
          _132_v38 = _dafny.Map.Empty.slice().updateUnsafe(_130_v36,_131_v37);
          let _rhs24 = _91_v2;
          let _rhs25 = _132_v38;
          let _lhs16 = _103_globalState;
          _lhs16.f18 = _rhs24;
          _132_v38 = _rhs25;
          let _133_v39;
          let _nw25 = new _module.C0();
          _nw25.__ctor(((_91_v2) ? (!(_91_v2)) : (!(_91_v2))));
          _133_v39 = _nw25;
          let _134_v40;
          let _nw26 = new _module.C0();
          _nw26.__ctor((_133_v39).f24);
          _134_v40 = _nw26;
          _134_v40 = _133_v39;
          let _index21 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_131_v37).length));
          (_131_v37)[_index21] = (_133_v39).f24;
          let _index22 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_131_v37).length));
          (_131_v37)[_index22] = ((_107_v16).Difference(_107_v16)).IsDisjointFrom((_107_v16).Difference(_107_v16));
        } else {
          let _135_v41;
          _135_v41 = _dafny.Set.fromElements(_98_v9, _98_v9);
          let _136_v42;
          _136_v42 = _dafny.Map.Empty.slice().updateUnsafe(_89_v0,_135_v41);
          let _137_v43;
          _137_v43 = _dafny.Seq.of((((_136_v42).contains(_89_v0)) ? ((_136_v42).get(_89_v0)) : (_135_v41)));
          let _138_v44;
          let _139_v45;
          let _out10;
          let _out11;
          let _outcollector5 = _module.__default.m0(_module.__default.fm0(_92_v3, _91_v2, _89_v0, _103_globalState), _dafny.Seq.IsPrefixOf(_137_v43, _137_v43), new BigNumber(495), _103_globalState);
          _out10 = _outcollector5[0];
          _out11 = _outcollector5[1];
          _138_v44 = _out10;
          _139_v45 = _out11;
          let _index23 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_97_v8).length));
          (_97_v8)[_index23] = (_105_v14).multipliedBy(_126_v32);
          let _index24 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_97_v8).length));
          (_97_v8)[_index24] = _139_v45;
          let _140_v46;
          let _init2 = ((_141_v2) => function (_142_i3) {
            return _141_v2;
          })(_91_v2);
          let _nw27 = Array((new BigNumber(11)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw27.length); _i0_2++) {
            _nw27[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _140_v46 = _nw27;
          let _index25 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_140_v46).length));
          (_140_v46)[_index25] = _91_v2;
          let _143_v47;
          _143_v47 = _module.D2.create_DC6(_89_v0);
          let _144_v48;
          let _nw28 = Array((new BigNumber(18)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = _89_v0;
          _nw28[(_dafny.ONE).toNumber()] = _89_v0;
          _nw28[(new BigNumber(2)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(3)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(4)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(5)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-932)), ((_145_v9) => function (_146_i4) {
            return _145_v9;
          })(_98_v9));
          _nw28[(new BigNumber(7)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(8)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(9)).toNumber()] = (_143_v47).dtor_cf9;
          _nw28[(new BigNumber(10)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(11)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-501)), function (_147_i5) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          });
          _nw28[(new BigNumber(13)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("awa");
          _nw28[(new BigNumber(15)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(16)).toNumber()] = _89_v0;
          _nw28[(new BigNumber(17)).toNumber()] = _89_v0;
          _144_v48 = _nw28;
          let _148_v49;
          _148_v49 = _module.D1.create_DC3((_module.D1.create_DC3(_144_v48, new BigNumber(779), _91_v2)).dtor_cf3, _105_v14, _91_v2);
          let _index26 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_140_v46).length));
          (_140_v46)[_index26] = !((_148_v49).dtor_cf5);
          let _149_v50;
          let _nw29 = new _module.C0();
          _nw29.__ctor(_module.__default.fm0(new BigNumber((_124_v30).length), (_140_v46)[_module.__default.safeIndex(new BigNumber(996), new BigNumber((_140_v46).length))], _dafny.Seq.update(_89_v0, _module.__default.safeIndex((_dafny.ZERO).minus(_127_v33), new BigNumber((_89_v0).length)), _98_v9), _103_globalState));
          _149_v50 = _nw29;
          _126_v32 = (_97_v8)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_97_v8).length))];
        }
      } else {
        _91_v2 = _91_v2;
        let _150_v51;
        let _nw30 = Array((new BigNumber(2)).toNumber()).fill(false);
        _150_v51 = _nw30;
        let _151_v52;
        _151_v52 = _dafny.Seq.of(_91_v2);
        let _index27 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_150_v51).length));
        (_150_v51)[_index27] = !_dafny.areEqual(_151_v52, _151_v52);
        let _index28 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_150_v51).length));
        (_150_v51)[_index28] = (_105_v14).isLessThanOrEqualTo(_module.__default.fm3(_106_v15, _103_globalState));
        let _152_v53;
        let _nw31 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _152_v53 = _nw31;
        let _153_v54;
        _153_v54 = _dafny.MultiSet.fromElements(_105_v14, _105_v14);
        let _154_v55;
        _154_v55 = _module.D1.create_DC3(_152_v53, new BigNumber((_153_v54).cardinality()), _module.__default.fm0(_92_v3, (_150_v51)[_module.__default.safeIndex(new BigNumber(115), new BigNumber((_150_v51).length))], _89_v0, _103_globalState));
        let _155_v56;
        let _156_v57;
        let _out12;
        let _out13;
        let _outcollector6 = _module.__default.m0(((_154_v55).dtor_cf4).isLessThanOrEqualTo(new BigNumber((_107_v16).cardinality())), (_105_v14).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("efngv")).length)), (_102_v13)[_module.__default.safeIndex(_106_v15, new BigNumber((_102_v13).length))], _103_globalState);
        _out12 = _outcollector6[0];
        _out13 = _outcollector6[1];
        _155_v56 = _out12;
        _156_v57 = _out13;
        let _157_v58;
        let _158_v59;
        let _out14;
        let _out15;
        let _outcollector7 = _module.__default.m0(_91_v2, (_150_v51)[_module.__default.safeIndex(new BigNumber(115), new BigNumber((_150_v51).length))], new BigNumber(210), _103_globalState);
        _out14 = _outcollector7[0];
        _out15 = _outcollector7[1];
        _157_v58 = _out14;
        _158_v59 = _out15;
        let _index29 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_150_v51).length));
        let _index30 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_150_v51).length));
        let _rhs26 = _150_v51;
        let _rhs27 = ((_158_v59).isLessThanOrEqualTo(_106_v15)) && ((_dafny.MultiSet.fromElements(_156_v57, _92_v3)).IsSubsetOf(_153_v54));
        let _rhs28 = (_150_v51)[_module.__default.safeIndex(new BigNumber(115), new BigNumber((_150_v51).length))];
        let _rhs29 = ((_dafny.ZERO).minus(_156_v57)).isLessThanOrEqualTo(_155_v56);
        let _rhs30 = (_dafny.ZERO).minus((_158_v59).minus(_155_v56));
        let _lhs17 = _150_v51;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_150_v51).length));
        let _lhs19 = _150_v51;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_150_v51).length));
        let _lhs21 = _103_globalState;
        let _lhs22 = _103_globalState;
        _150_v51 = _rhs26;
        _lhs17[_lhs18] = _rhs27;
        _lhs19[_lhs20] = _rhs28;
        _lhs21.f2 = _rhs29;
        _lhs22.f23 = _rhs30;
      }
      let _159_v60;
      let _nw32 = Array((new BigNumber(28)).toNumber()).fill(false);
      _159_v60 = _nw32;
      let _160_v61;
      _160_v61 = _dafny.Map.Empty.slice().updateUnsafe(_159_v60,_89_v0);
      (_103_globalState).f2 = _module.__default.fm0(_105_v14, _91_v2, (((_160_v61).contains(_159_v60)) ? ((_160_v61).get(_159_v60)) : (_89_v0)), _103_globalState);
      let _161_v62;
      let _nw33 = new _module.C0();
      _nw33.__ctor(!(_91_v2));
      _161_v62 = _nw33;
      let _162_v63;
      _162_v63 = _dafny.Map.Empty.slice().updateUnsafe(_91_v2,_161_v62);
      let _163_v64;
      _163_v64 = _dafny.Map.Empty.slice().updateUnsafe(_162_v63,_dafny.Seq.Create(_module.__default.abs(new BigNumber(950)), ((_164_v13) => function (_165_i6) {
        return new BigNumber((_164_v13).length);
      })(_102_v13)));
      _163_v64 = (_163_v64).update(_162_v63, _102_v13);
      let _166_v65;
      _166_v65 = _dafny.Map.Empty.slice().updateUnsafe((_161_v62).f24,true);
      if ((_166_v65).contains((_161_v62).f24)) {
        let _167_v66;
        _167_v66 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_106_v15, _92_v3),_97_v8);
        _167_v66 = (_167_v66).update(new BigNumber(271), _97_v8);
        let _168_v67;
        let _init3 = ((_169_v0) => function (_170_i7) {
          return _169_v0;
        })(_89_v0);
        let _nw34 = Array((new BigNumber(5)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw34.length); _i0_3++) {
          _nw34[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _168_v67 = _nw34;
        let _source0 = _module.D1.create_DC3(_168_v67, new BigNumber(712), !(((((_107_v16).contains((_161_v62).f24)) ? ((_107_v16).get((_161_v62).f24)) : (new BigNumber(-409)))).isLessThanOrEqualTo(new BigNumber(-154))));
        if (_source0.is_DC3) {
          let _171___mcc_h0 = (_source0).cf3;
          let _172___mcc_h1 = (_source0).cf4;
          let _173___mcc_h2 = (_source0).cf5;
          let _174_cf5 = _173___mcc_h2;
          let _175_cf4 = _172___mcc_h1;
          let _176_cf3 = _171___mcc_h0;
          let _177_v68;
          _177_v68 = _module.D0.create_DC1(_101_v12);
          let _178_v69;
          _178_v69 = _dafny.Map.Empty.slice().updateUnsafe(_105_v14,_177_v68);
          let _179_v70;
          let _180_v71;
          let _out16;
          let _out17;
          let _outcollector8 = _module.__default.m0((_161_v62).f24, _module.__default.fm0(new BigNumber((_178_v69).length), (_161_v62).f24, _dafny.Seq.update(_89_v0, _module.__default.safeIndex(_106_v15, new BigNumber((_89_v0).length)), _98_v9), _103_globalState), _106_v15, _103_globalState);
          _out16 = _outcollector8[0];
          _out17 = _outcollector8[1];
          _179_v70 = _out16;
          _180_v71 = _out17;
          _91_v2 = _174_cf5;
          let _index31 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_159_v60).length));
          (_159_v60)[_index31] = (false) && (_91_v2);
          let _index32 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_159_v60).length));
          let _rhs31 = (new BigNumber((_100_v11).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(_175_cf4));
          let _rhs32 = (_161_v62).f24;
          let _rhs33 = (_161_v62).f24;
          let _lhs23 = _159_v60;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_159_v60).length));
          let _lhs25 = _103_globalState;
          _lhs23[_lhs24] = _rhs31;
          _lhs25.f2 = _rhs32;
          _174_cf5 = _rhs33;
          let _181_v72;
          let _182_v73;
          let _out18;
          let _out19;
          let _outcollector9 = _module.__default.m0((_dafny.Set.fromElements((_159_v60)[_module.__default.safeIndex(new BigNumber(19), new BigNumber((_159_v60).length))])).IsSubsetOf(_dafny.Set.fromElements((_161_v62).f24, true, _174_cf5)), (_161_v62).f24, _106_v15, _103_globalState);
          _out18 = _outcollector9[0];
          _out19 = _outcollector9[1];
          _181_v72 = _out18;
          _182_v73 = _out19;
        } else if (_source0.is_DC4) {
          let _183___mcc_h3 = (_source0).cf6;
          let _184___mcc_h4 = (_source0).cf7;
          let _185_cf7 = _184___mcc_h4;
          let _186_cf6 = _183___mcc_h3;
          let _187_v74;
          _187_v74 = _module.D2.create_DC7(_106_v15, new _dafny.CodePoint('a'.codePointAt(0)), _161_v62, (_161_v62).f24);
          let _188_v75;
          _188_v75 = _dafny.MultiSet.fromElements(_161_v62, (_187_v74).dtor_cf12);
          let _189_v76;
          _189_v76 = _dafny.MultiSet.fromElements(new BigNumber((_188_v75).cardinality()));
          _106_v15 = (((_107_v16).contains((_161_v62).f24)) ? ((_107_v16).get((_161_v62).f24)) : (_module.__default.safeDivisionInt(_105_v14, new BigNumber((_189_v76).cardinality()))));
          let _pat_let_tv0 = _105_v14;
          let _pat_let_tv1 = _161_v62;
          let _190_v77;
          let _nw35 = Array((new BigNumber(16)).toNumber());
          _nw35[(_dafny.ZERO).toNumber()] = _187_v74;
          _nw35[(_dafny.ONE).toNumber()] = _module.D2.create_DC7(_106_v15, _98_v9, _161_v62, (_161_v62).f24);
          _nw35[(new BigNumber(2)).toNumber()] = _187_v74;
          _nw35[(new BigNumber(3)).toNumber()] = _187_v74;
          _nw35[(new BigNumber(4)).toNumber()] = ((!(_91_v2)) ? (_187_v74) : (_187_v74));
          _nw35[(new BigNumber(5)).toNumber()] = _module.D2.create_DC7(_106_v15, _98_v9, _161_v62, _185_cf7);
          _nw35[(new BigNumber(6)).toNumber()] = _187_v74;
          _nw35[(new BigNumber(7)).toNumber()] = _187_v74;
          _nw35[(new BigNumber(8)).toNumber()] = _187_v74;
          _nw35[(new BigNumber(9)).toNumber()] = _187_v74;
          _nw35[(new BigNumber(10)).toNumber()] = _187_v74;
          _nw35[(new BigNumber(11)).toNumber()] = _187_v74;
          _nw35[(new BigNumber(12)).toNumber()] = (((_161_v62).f24) ? (_187_v74) : (_187_v74));
          _nw35[(new BigNumber(13)).toNumber()] = _187_v74;
          _nw35[(new BigNumber(14)).toNumber()] = function (_pat_let0_0) {
            return function (_191_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_192_dt__update_hcf10_h0) {
                  return function (_pat_let2_0) {
                    return function (_193_dt__update_hcf12_h0) {
                      return _module.D2.create_DC7(_192_dt__update_hcf10_h0, (_191_dt__update__tmp_h0).dtor_cf11, _193_dt__update_hcf12_h0, (_191_dt__update__tmp_h0).dtor_cf13);
                    }(_pat_let2_0);
                  }(_pat_let_tv1);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_187_v74);
          _nw35[(new BigNumber(15)).toNumber()] = _187_v74;
          _190_v77 = _nw35;
          let _index33 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_190_v77).length));
          (_190_v77)[_index33] = _187_v74;
          let _index34 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_97_v8).length));
          (_97_v8)[_index34] = _92_v3;
          let _index35 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_190_v77).length));
          let _index36 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_97_v8).length));
          let _rhs34 = (((new BigNumber(309)).isLessThan(_105_v14)) ? (_187_v74) : (_module.D2.create_DC7(new BigNumber((_module.__default.fm6((((_189_v76).contains(_106_v15)) ? ((_189_v76).get(_106_v15)) : (_92_v3)), (_161_v62).f24, (_161_v62).f24, _103_globalState)).length), _98_v9, _161_v62, _91_v2)));
          let _rhs35 = _106_v15;
          let _rhs36 = _106_v15;
          let _rhs37 = _92_v3;
          let _lhs26 = _190_v77;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_190_v77).length));
          let _lhs28 = _103_globalState;
          let _lhs29 = _103_globalState;
          let _lhs30 = _97_v8;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_97_v8).length));
          _lhs26[_lhs27] = _rhs34;
          _lhs28.f23 = _rhs35;
          _lhs29.f23 = _rhs36;
          _lhs30[_lhs31] = _rhs37;
          _161_v62 = _161_v62;
          _92_v3 = (_dafny.ZERO).minus((_102_v13)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_89_v0).length), new BigNumber((_89_v0).length)), new BigNumber((_102_v13).length))]);
        } else if (_source0.is_DC2) {
          let _194___mcc_h5 = (_source0).cf2;
          let _195_cf2 = _194___mcc_h5;
          let _196_v78;
          let _nw36 = new _module.C0();
          _nw36.__ctor((_161_v62).f24);
          _196_v78 = _nw36;
          let _197_v79;
          let _198_v80;
          let _out20;
          let _out21;
          let _outcollector10 = _module.__default.m0((_161_v62).f24, (_196_v78).f24, _106_v15, _103_globalState);
          _out20 = _outcollector10[0];
          _out21 = _outcollector10[1];
          _197_v79 = _out20;
          _198_v80 = _out21;
          (_103_globalState).f18 = _91_v2;
          (_103_globalState).f20 = _89_v0;
        } else {
          let _199___mcc_h6 = (_source0).cf8;
          let _200_cf8 = _199___mcc_h6;
          let _index37 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((_159_v60).length));
          (_159_v60)[_index37] = (_161_v62).f24;
          let _index38 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((_159_v60).length));
          (_159_v60)[_index38] = _module.__default.fm0(_module.__default.safeDivisionInt(_106_v15, _106_v15), (_161_v62).f24, _dafny.Seq.Concat(_89_v0, _89_v0), _103_globalState);
          _92_v3 = new BigNumber((_89_v0).length);
          let _201_v81;
          let _init4 = ((_202_v62) => function (_203_i8) {
            return (((_202_v62).f24) ? (false) : ((_202_v62).f24));
          })(_161_v62);
          let _nw37 = Array((new BigNumber(21)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw37.length); _i0_4++) {
            _nw37[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _201_v81 = _nw37;
          _201_v81 = _201_v81;
          let _204_v82;
          _204_v82 = _module.D2.create_DC7(_105_v14, _98_v9, _161_v62, (_161_v62).f24);
          let _205_v83;
          _205_v83 = _dafny.Map.Empty.slice().updateUnsafe(_161_v62,!((_161_v62).f24));
          let _206_v84;
          let _207_v85;
          let _out22;
          let _out23;
          let _outcollector11 = _module.__default.m0(_module.__default.fm0(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-210)), function (_208_i9) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length), (_159_v60)[_module.__default.safeIndex(new BigNumber(362), new BigNumber((_159_v60).length))], _89_v0, _103_globalState), _module.__default.fm0(new BigNumber(65), (_204_v82).dtor_cf13, _89_v0, _103_globalState), new BigNumber((_205_v83).length), _103_globalState);
          _out22 = _outcollector11[0];
          _out23 = _outcollector11[1];
          _206_v84 = _out22;
          _207_v85 = _out23;
        }
        _107_v16 = _dafny.MultiSet.fromElements(_91_v2, (new BigNumber(741)).isEqualTo(_106_v15), (_161_v62).f24);
        let _209_v86;
        _209_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(11),new BigNumber(619));
        _92_v3 = ((((_209_v86).contains((_dafny.ZERO).minus(_105_v14))) ? ((_209_v86).get((_dafny.ZERO).minus(_105_v14))) : (_106_v15))).plus(new BigNumber((_89_v0).length));
        let _210_v87;
        let _211_v88;
        let _out24;
        let _out25;
        let _outcollector12 = _module.__default.m0((_161_v62).f24, (_161_v62).f24, _106_v15, _103_globalState);
        _out24 = _outcollector12[0];
        _out25 = _outcollector12[1];
        _210_v87 = _out24;
        _211_v88 = _out25;
      } else {
        let _212_v89;
        let _nw38 = Array((new BigNumber(9)).toNumber()).fill(_module.D1.Default());
        _212_v89 = _nw38;
        let _213_v90;
        let _nw39 = Array((new BigNumber(12)).toNumber());
        _nw39[(_dafny.ZERO).toNumber()] = _89_v0;
        _nw39[(_dafny.ONE).toNumber()] = _89_v0;
        _nw39[(new BigNumber(2)).toNumber()] = _89_v0;
        _nw39[(new BigNumber(3)).toNumber()] = _89_v0;
        _nw39[(new BigNumber(4)).toNumber()] = _89_v0;
        _nw39[(new BigNumber(5)).toNumber()] = _89_v0;
        _nw39[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), ((_214_v9) => function (_215_i10) {
          return _214_v9;
        })(_98_v9));
        _nw39[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("pqgonfj");
        _nw39[(new BigNumber(8)).toNumber()] = _89_v0;
        _nw39[(new BigNumber(9)).toNumber()] = _89_v0;
        _nw39[(new BigNumber(10)).toNumber()] = _89_v0;
        _nw39[(new BigNumber(11)).toNumber()] = _89_v0;
        _213_v90 = _nw39;
        let _216_v91;
        _216_v91 = _module.D1.create_DC3(_213_v90, _106_v15, true);
        let _index39 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_212_v89).length));
        (_212_v89)[_index39] = _216_v91;
        let _index40 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_212_v89).length));
        (_212_v89)[_index40] = _216_v91;
        _98_v9 = new _dafny.CodePoint('c'.codePointAt(0));
        let _217_v92;
        _217_v92 = _dafny.Map.Empty.slice().updateUnsafe((_105_v14).minus(_92_v3),_98_v9);
        _217_v92 = (_217_v92).update(_92_v3, _98_v9);
        (_103_globalState).f18 = (_161_v62).f24;
        let _218_v93;
        _218_v93 = _module.D1.create_DC2((new BigNumber((_102_v13).length)).isLessThanOrEqualTo(_105_v14));
        let _pat_let_tv2 = _91_v2;
        let _rhs38 = _105_v14;
        let _rhs39 = function (_pat_let3_0) {
          return function (_219_dt__update__tmp_h1) {
            return function (_pat_let4_0) {
              return function (_220_dt__update_hcf2_h0) {
                return _module.D1.create_DC2(_220_dt__update_hcf2_h0);
              }(_pat_let4_0);
            }(_pat_let_tv2);
          }(_pat_let3_0);
        }((((_161_v62).f24) ? (_218_v93) : (_218_v93)));
        let _rhs40 = _98_v9;
        _105_v14 = _rhs38;
        _218_v93 = _rhs39;
        _98_v9 = _rhs40;
      }
      _93_v4 = (_93_v4).update(_91_v2, (_dafny.ZERO).minus(((_dafny.ZERO).minus(_106_v15)).minus(_92_v3)));
      let _hi2 = _106_v15;
      for (let _221_i11 = (_105_v14).plus(_106_v15); _221_i11.isLessThan(_hi2); _221_i11 = _221_i11.plus(_dafny.ONE)) {
        let _222_v94;
        _222_v94 = _dafny.MultiSet.fromElements(_92_v3);
        if (((((_222_v94).contains(_92_v3)) ? ((_222_v94).get(_92_v3)) : (_106_v15))).isLessThan((_dafny.ZERO).minus((_106_v15).multipliedBy(_106_v15)))) {
          let _index41 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_159_v60).length));
          (_159_v60)[_index41] = _91_v2;
          let _index42 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_159_v60).length));
          (_159_v60)[_index42] = (_161_v62).f24;
          let _223_v95;
          _223_v95 = _dafny.Seq.of(_101_v12, _101_v12, _101_v12, _101_v12);
          (_103_globalState).f15 = (_223_v95)[_module.__default.safeIndex(_92_v3, new BigNumber((_223_v95).length))];
          _98_v9 = _98_v9;
          let _224_v96;
          _224_v96 = _dafny.Seq.of(true, (_161_v62).f24, _91_v2);
          let _225_v97;
          let _nw40 = Array((new BigNumber(13)).toNumber()).fill([]);
          _225_v97 = _nw40;
          let _226_v98;
          let _nw41 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Set.Empty);
          _226_v98 = _nw41;
          let _index43 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_225_v97).length));
          (_225_v97)[_index43] = _226_v98;
          let _227_v99;
          _227_v99 = _dafny.Seq.of(_226_v98, _226_v98, _226_v98);
          let _index44 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_225_v97).length));
          let _index45 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_159_v60).length));
          let _rhs41 = (_dafny.ZERO).minus(_221_i11);
          let _rhs42 = _dafny.Seq.Concat(_module.__default.fm7(_103_globalState), _102_v13);
          let _rhs43 = _dafny.Seq.Concat(_224_v96, _dafny.Seq.of((_159_v60)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_159_v60).length))], (_161_v62).f24));
          let _rhs44 = (_227_v99)[_module.__default.safeIndex(_106_v15, new BigNumber((_227_v99).length))];
          let _rhs45 = (_92_v3).isLessThan(_221_i11);
          let _lhs32 = _225_v97;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_225_v97).length));
          let _lhs34 = _159_v60;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_159_v60).length));
          _92_v3 = _rhs41;
          _102_v13 = _rhs42;
          _224_v96 = _rhs43;
          _lhs32[_lhs33] = _rhs44;
          _lhs34[_lhs35] = _rhs45;
          let _228_v100;
          _228_v100 = _dafny.Seq.of(_107_v16, _107_v16);
          let _229_v101;
          _229_v101 = _dafny.Map.Empty.slice().updateUnsafe(!((_161_v62).f24),_dafny.MultiSet.FromArray(_228_v100));
          let _230_v102;
          _230_v102 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_224_v96), _107_v16);
          let _231_v103;
          let _232_v104;
          let _out26;
          let _out27;
          let _outcollector13 = _module.__default.m0((_161_v62).f24, _91_v2, new BigNumber(((((_229_v101).contains((_159_v60)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_159_v60).length))])) ? ((_229_v101).get((_159_v60)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_159_v60).length))])) : (_230_v102))).cardinality()), _103_globalState);
          _out26 = _outcollector13[0];
          _out27 = _outcollector13[1];
          _231_v103 = _out26;
          _232_v104 = _out27;
        } else {
          _100_v11 = (_100_v11).update(_105_v14, (_161_v62).f24);
          let _233_v105;
          _233_v105 = _dafny.Set.fromElements(_91_v2);
          let _234_v106;
          _234_v106 = _dafny.Set.fromElements((_dafny.Set.fromElements((_161_v62).f24)).Difference(_233_v105), _233_v105, _233_v105, _233_v105, _233_v105);
          _234_v106 = _dafny.Set.fromElements((_233_v105).Difference(_233_v105), _233_v105);
          (_103_globalState).f20 = _dafny.Seq.Concat(_dafny.Seq.update(_89_v0, _module.__default.safeIndex(_92_v3, new BigNumber((_89_v0).length)), _98_v9), _89_v0);
          let _235_v107;
          _235_v107 = _module.D2.create_DC6(_89_v0);
          _235_v107 = _235_v107;
          _92_v3 = _221_i11;
        }
        let _236_v108;
        _236_v108 = _dafny.Seq.of((((_161_v62).f24) ? (_161_v62) : (_161_v62)), _161_v62, _161_v62);
        let _237_v109;
        let _nw42 = Array((new BigNumber(7)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = _89_v0;
        _nw42[(_dafny.ONE).toNumber()] = _89_v0;
        _nw42[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_89_v0, _89_v0);
        _nw42[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("hoeogo");
        _nw42[(new BigNumber(4)).toNumber()] = _89_v0;
        _nw42[(new BigNumber(5)).toNumber()] = _89_v0;
        _nw42[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("hsmngmhxs");
        _237_v109 = _nw42;
        let _index46 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_237_v109).length));
        (_237_v109)[_index46] = _dafny.Seq.UnicodeFromString("ite");
        let _238_v110;
        _238_v110 = _dafny.Map.Empty.slice().updateUnsafe(_89_v0,_89_v0);
        let _index47 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_237_v109).length));
        let _rhs46 = (((_238_v110).contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rvomcrdu"), _dafny.Seq.UnicodeFromString("pk")))) ? ((_238_v110).get(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rvomcrdu"), _dafny.Seq.UnicodeFromString("pk")))) : (_dafny.Seq.Concat(_89_v0, _89_v0)));
        let _rhs47 = _236_v108;
        let _rhs48 = _dafny.Seq.Concat(_dafny.Seq.Concat(_89_v0, _89_v0), _89_v0);
        let _lhs36 = _103_globalState;
        let _lhs37 = _237_v109;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_237_v109).length));
        _lhs36.f20 = _rhs46;
        _236_v108 = _rhs47;
        _lhs37[_lhs38] = _rhs48;
        let _239_v111;
        let _init5 = ((_240_v16) => function (_241_i12) {
          return _240_v16;
        })(_107_v16);
        let _nw43 = Array((new BigNumber(27)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw43.length); _i0_5++) {
          _nw43[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _239_v111 = _nw43;
        let _index48 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_239_v111).length));
        (_239_v111)[_index48] = _107_v16;
        let _index49 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_239_v111).length));
        (_239_v111)[_index49] = ((_107_v16).Union(_107_v16)).Difference((_107_v16).update((_161_v62).f24, _module.__default.abs(_221_i11)));
        let _242_v112;
        _242_v112 = _dafny.Set.fromElements((_161_v62).f24, _91_v2);
        let _source1 = _module.D1.create_DC2((_dafny.Set.fromElements(_91_v2)).IsSubsetOf(_242_v112));
        if (_source1.is_DC3) {
          let _243___mcc_h7 = (_source1).cf3;
          let _244___mcc_h8 = (_source1).cf4;
          let _245___mcc_h9 = (_source1).cf5;
          let _246_cf5 = _245___mcc_h9;
          let _247_cf4 = _244___mcc_h8;
          let _248_cf3 = _243___mcc_h7;
          let _index50 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_159_v60).length));
          (_159_v60)[_index50] = false;
          let _index51 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_159_v60).length));
          (_159_v60)[_index51] = (_221_i11).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_92_v3, _106_v15))));
          let _249_v113;
          _249_v113 = _dafny.Map.Empty.slice().updateUnsafe(_105_v14,_247_cf4);
          _249_v113 = (_249_v113).update(new BigNumber(103), (_105_v14).multipliedBy(_106_v15));
          let _250_v114;
          let _nw44 = new _module.C0();
          _nw44.__ctor((_159_v60)[_module.__default.safeIndex(new BigNumber(233), new BigNumber((_159_v60).length))]);
          _250_v114 = _nw44;
          (_103_globalState).f18 = (_246_cf5) === ((_161_v62).f24);
        } else if (_source1.is_DC4) {
          let _251___mcc_h10 = (_source1).cf6;
          let _252___mcc_h11 = (_source1).cf7;
          let _253_cf7 = _252___mcc_h11;
          let _254_cf6 = _251___mcc_h10;
          let _255_v115;
          _255_v115 = _dafny.Map.Empty.slice().updateUnsafe(_102_v13,(_161_v62).f24);
          (_103_globalState).f23 = (new BigNumber((_255_v115).length)).multipliedBy((new BigNumber(-803)).multipliedBy(new BigNumber(-924)));
          (_103_globalState).f23 = _105_v14;
          let _256_v116;
          let _257_v117;
          let _out28;
          let _out29;
          let _outcollector14 = _module.__default.m0(_91_v2, true, (_dafny.ZERO).minus(_221_i11), _103_globalState);
          _out28 = _outcollector14[0];
          _out29 = _outcollector14[1];
          _256_v116 = _out28;
          _257_v117 = _out29;
          let _index52 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_159_v60).length));
          (_159_v60)[_index52] = (_161_v62).f24;
          let _index53 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_159_v60).length));
          (_159_v60)[_index53] = _dafny.Seq.IsPrefixOf(_89_v0, _89_v0);
          let _258_v119;
          _258_v119 = _module.D1.create_DC4(_254_cf6, _91_v2);
          let _index54 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_159_v60).length));
          let _index55 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_159_v60).length));
          let _rhs49 = !(function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of _dafny.IntegerRange(new BigNumber(248), new BigNumber(-932))) {
              let _259_v118 = _compr_6;
              if (((new BigNumber(248)).isLessThanOrEqualTo(_259_v118)) && ((_259_v118).isLessThan(new BigNumber(-932)))) {
                _coll6.add((_259_v118).plus(_106_v15));
              }
            }
            return _coll6;
          }()).contains(new BigNumber(728));
          let _rhs50 = _161_v62;
          let _rhs51 = (_161_v62).f24;
          let _rhs52 = _dafny.Seq.Concat(_254_cf6, (_258_v119).dtor_cf6);
          let _rhs53 = (_161_v62).f24;
          let _lhs39 = _159_v60;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_159_v60).length));
          let _lhs41 = _103_globalState;
          let _lhs42 = _159_v60;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_159_v60).length));
          _lhs39[_lhs40] = _rhs49;
          _161_v62 = _rhs50;
          _lhs41.f2 = _rhs51;
          _254_cf6 = _rhs52;
          _lhs42[_lhs43] = _rhs53;
        } else if (_source1.is_DC2) {
          let _260___mcc_h12 = (_source1).cf2;
          let _261_cf2 = _260___mcc_h12;
          _105_v14 = (_module.__default.safeModuloInt(new BigNumber(978), _92_v3)).minus((_221_i11).plus(_221_i11));
          let _262_v120;
          let _init6 = ((_263_i11) => function (_264_i13) {
            return _dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(_263_i11));
          })(_221_i11);
          let _nw45 = Array((new BigNumber(2)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw45.length); _i0_6++) {
            _nw45[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _262_v120 = _nw45;
          let _index56 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_262_v120).length));
          (_262_v120)[_index56] = _93_v4;
          let _index57 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_262_v120).length));
          (_262_v120)[_index57] = _93_v4;
          let _265_v121;
          let _nw46 = Array((new BigNumber(28)).toNumber()).fill([]);
          _265_v121 = _nw46;
          let _266_v122;
          _266_v122 = _dafny.Map.Empty.slice().updateUnsafe(((_239_v111)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_239_v111).length))]).Difference((_239_v111)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_239_v111).length))]),_161_v62);
          let _267_v123;
          _267_v123 = _module.D3.create_DC9(_159_v60);
          let _268_v124;
          _268_v124 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pwflyn"),_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(_261_cf2, (_161_v62).f24)),_161_v62));
          let _rhs54 = _265_v121;
          let _rhs55 = _161_v62;
          let _rhs56 = (_267_v123).dtor_cf15;
          let _rhs57 = _91_v2;
          let _rhs58 = (((_268_v124).contains(_89_v0)) ? ((_268_v124).get(_89_v0)) : ((_266_v122).Merge(_266_v122)));
          _265_v121 = _rhs54;
          _161_v62 = _rhs55;
          _159_v60 = _rhs56;
          _261_cf2 = _rhs57;
          _266_v122 = _rhs58;
          (_103_globalState).f23 = (new BigNumber(810)).plus(_106_v15);
        } else {
          let _269___mcc_h13 = (_source1).cf8;
          let _270_cf8 = _269___mcc_h13;
          let _271_v125;
          _271_v125 = _dafny.Map.Empty.slice().updateUnsafe(_221_i11,_105_v14);
          let _272_v126;
          _272_v126 = _dafny.Map.Empty.slice().updateUnsafe((((_166_v65).contains((_161_v62).f24)) ? ((_166_v65).get((_161_v62).f24)) : ((_161_v62).f24)),_271_v125);
          let _273_v127;
          _273_v127 = _dafny.Map.Empty.slice().updateUnsafe(_92_v3,(((_93_v4).contains(false)) ? ((_93_v4).get(false)) : (new BigNumber((_272_v126).length))));
          (_103_globalState).f23 = _module.__default.safeDivisionInt((((_161_v62).f24) ? (new BigNumber(422)) : (new BigNumber((_89_v0).length))), (_dafny.ZERO).minus((_dafny.ZERO).minus(((_dafny.ZERO).minus(_105_v14)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_273_v127).length))))));
          (_103_globalState).f20 = ((_91_v2) ? ((_161_v62).fm5(_103_globalState)) : (_89_v0));
          let _274_v128;
          _274_v128 = _dafny.Seq.of((_161_v62).f24);
          let _275_v129;
          _275_v129 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_dafny.Seq.update(_274_v128, _module.__default.safeIndex(_106_v15, new BigNumber((_274_v128).length)), !((_161_v62).f24)), (_161_v62).f24),_106_v15);
          _275_v129 = (_275_v129).update(_module.__default.fm8(_106_v15, new BigNumber((_102_v13).length), !((_161_v62).f24), _103_globalState), (_dafny.ZERO).minus(_221_i11));
          (_103_globalState).f23 = _221_i11;
        }
      }
      let _index58 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length));
      (_159_v60)[_index58] = (_161_v62).f24;
      let _index59 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length));
      (_159_v60)[_index59] = ((_92_v3).multipliedBy(_105_v14)).isLessThan(_105_v14);
      let _276_i14;
      _276_i14 = _dafny.ZERO;
      L1: {
        while ((((_159_v60)[_module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length))]) ? ((_159_v60)[_module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length))]) : ((_159_v60)[_module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length))]))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_276_i14)) {
              break L1;
            }
            _276_i14 = (_276_i14).plus(_dafny.ONE);
            (_103_globalState).f2 = (_91_v2) || ((_161_v62).f24);
            (_103_globalState).f23 = new BigNumber(-84);
            let _277_v130;
            _277_v130 = _dafny.Seq.of((_161_v62).f24);
            let _278_v131;
            let _279_v132;
            let _out30;
            let _out31;
            let _outcollector15 = _module.__default.m0((_277_v130)[_module.__default.safeIndex(_92_v3, new BigNumber((_277_v130).length))], _91_v2, _105_v14, _103_globalState);
            _out30 = _outcollector15[0];
            _out31 = _outcollector15[1];
            _278_v131 = _out30;
            _279_v132 = _out31;
            let _280_v133;
            let _281_v134;
            let _out32;
            let _out33;
            let _outcollector16 = _module.__default.m0((_159_v60)[_module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length))], (_module.__default.fm3(new BigNumber(984), _103_globalState)).isEqualTo((_161_v62).fm4(false, _91_v2, _dafny.MultiSet.FromArray(_module.__default.fm7(_103_globalState)), _91_v2, _103_globalState)), _106_v15, _103_globalState);
            _out32 = _outcollector16[0];
            _out33 = _outcollector16[1];
            _280_v133 = _out32;
            _281_v134 = _out33;
          }
        }
      }
      let _rhs59 = _97_v8;
      let _rhs60 = new BigNumber(-668);
      _97_v8 = _rhs59;
      _92_v3 = _rhs60;
      let _282_v135;
      _282_v135 = _dafny.Set.fromElements(_91_v2, _91_v2);
      let _283_v136;
      let _nw47 = Array((new BigNumber(4)).toNumber());
      _nw47[(_dafny.ZERO).toNumber()] = _161_v62;
      _nw47[(_dafny.ONE).toNumber()] = _161_v62;
      _nw47[(new BigNumber(2)).toNumber()] = _161_v62;
      _nw47[(new BigNumber(3)).toNumber()] = _161_v62;
      _283_v136 = _nw47;
      let _index60 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_283_v136).length));
      (_283_v136)[_index60] = _161_v62;
      let _284_v137;
      _284_v137 = _module.D0.create_DC0(_98_v9);
      let _285_v138;
      _285_v138 = _dafny.Map.Empty.slice().updateUnsafe(_284_v137,_161_v62);
      let _pat_let_tv3 = _98_v9;
      let _pat_let_tv4 = _98_v9;
      let _286_v139;
      _286_v139 = _dafny.Seq.of((((_285_v138).contains(function (_pat_let7_0) {
        return function (_289_dt__update__tmp_h2) {
          return function (_pat_let8_0) {
            return function (_290_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_290_dt__update_hcf0_h0);
            }(_pat_let8_0);
          }(_pat_let_tv4);
        }(_pat_let7_0);
      }(_284_v137))) ? ((_285_v138).get(function (_pat_let5_0) {
        return function (_287_dt__update__tmp_h3) {
          return function (_pat_let6_0) {
            return function (_288_dt__update_hcf0_h1) {
              return _module.D0.create_DC0(_288_dt__update_hcf0_h1);
            }(_pat_let6_0);
          }(_pat_let_tv3);
        }(_pat_let5_0);
      }(_284_v137))) : (_161_v62)), _161_v62);
      let _index61 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_283_v136).length));
      let _rhs61 = _159_v60;
      let _rhs62 = _282_v135;
      let _rhs63 = !((((_159_v60)[_module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length))]) ? ((_159_v60)[_module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length))]) : ((_161_v62).f24)));
      let _rhs64 = (_286_v139)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_286_v139).length))];
      let _lhs44 = _103_globalState;
      let _lhs45 = _283_v136;
      let _lhs46 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_283_v136).length));
      _159_v60 = _rhs61;
      _282_v135 = _rhs62;
      _lhs44.f2 = _rhs63;
      _lhs45[_lhs46] = _rhs64;
      let _nw48 = new _module.C0();
      _nw48.__ctor(!((_159_v60)[_module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length))]));
      _161_v62 = _nw48;
      _106_v15 = (_105_v14).minus((_105_v14).plus(new BigNumber((_102_v13).length)));
      let _291_v140;
      _291_v140 = _dafny.MultiSet.fromElements(_module.__default.fm9(_106_v15, (_159_v60)[_module.__default.safeIndex(new BigNumber(794), new BigNumber((_159_v60).length))], _92_v3, _91_v2, _103_globalState), _100_v11);
      let _292_v141;
      _292_v141 = _dafny.Map.Empty.slice().updateUnsafe(_291_v140,_dafny.Seq.update((((_161_v62).f24) ? (_dafny.Seq.UnicodeFromString("balyols")) : (_89_v0)), _module.__default.safeIndex(_106_v15, new BigNumber(((((_161_v62).f24) ? (_dafny.Seq.UnicodeFromString("balyols")) : (_89_v0))).length)), _98_v9));
      let _index62 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_283_v136).length));
      let _rhs65 = _dafny.Seq.Concat(_102_v13, _dafny.Seq.of(_92_v3));
      let _rhs66 = (_283_v136)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_283_v136).length))];
      let _rhs67 = _159_v60;
      let _rhs68 = (((_292_v141).contains((_291_v140).Difference(_module.__default.fm10(_91_v2, _91_v2, (_161_v62).f24, _103_globalState)))) ? ((_292_v141).get((_291_v140).Difference(_module.__default.fm10(_91_v2, _91_v2, (_161_v62).f24, _103_globalState)))) : (_dafny.Seq.UnicodeFromString("nhrm")));
      let _lhs47 = _283_v136;
      let _lhs48 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_283_v136).length));
      let _lhs49 = _103_globalState;
      _102_v13 = _rhs65;
      _lhs47[_lhs48] = _rhs66;
      _159_v60 = _rhs67;
      _lhs49.f20 = _rhs68;
      let _293_v142;
      let _nw49 = new _module.C0();
      _nw49.__ctor(true);
      _293_v142 = _nw49;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_97_v8).length))) {
        let _294_i15 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_294_i15)) && ((_294_i15).isLessThan(new BigNumber((_97_v8).length))))) {
          (_97_v8)[(_294_i15)] = _module.__default.safeModuloInt(_294_i15, new BigNumber(((((_293_v142).f24) ? (_dafny.Map.Empty.slice().updateUnsafe(_105_v14,_92_v3)) : (_dafny.Map.Empty.slice().updateUnsafe(_92_v3,_92_v3)))).length));
        }
      }
      process.stdout.write((_89_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v1).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("spiw"), _dafny.Seq.UnicodeFromString("spiw"), _dafny.Seq.UnicodeFromString("spiw"), _dafny.Seq.UnicodeFromString("spiw"), _dafny.Seq.UnicodeFromString("spiw")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_91_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_92_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(56)).updateUnsafe(false,new BigNumber(116)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(56)),new BigNumber(56)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_96_v7).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v8)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_98_v9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_99_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('n'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(56),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v12)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_102_v13, _dafny.Seq.of(new BigNumber(56), new BigNumber(56), new BigNumber(-668)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f1).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("spiw"), _dafny.Seq.UnicodeFromString("spiw"), _dafny.Seq.UnicodeFromString("spiw"), _dafny.Seq.UnicodeFromString("spiw"), _dafny.Seq.UnicodeFromString("spiw")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_103_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f3).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(56)),new BigNumber(56)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f13)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_103_globalState).f14).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState.f15)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_103_globalState).f17)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_103_globalState.f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_103_globalState.f20).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_103_globalState).f21, _dafny.Seq.of(new BigNumber(56), new BigNumber(56)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_103_globalState).f22, _dafny.Seq.of(new BigNumber(56), new BigNumber(56)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_103_globalState.f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_v14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_106_v15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_107_v16).equals(_dafny.MultiSet.fromElements(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v60)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v60)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_160_v61).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v62).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_162_v63).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_163_v64).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v65).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_276_i14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v135).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v136)[_dafny.ZERO]).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v136)[_dafny.ONE]).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v136)[new BigNumber(2)]).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v136)[new BigNumber(3)]).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v137).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_285_v138).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_286_v139).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_291_v140).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(21),false).updateUnsafe(new BigNumber(22),false).updateUnsafe(new BigNumber(23),false).updateUnsafe(new BigNumber(24),false).updateUnsafe(new BigNumber(25),false).updateUnsafe(new BigNumber(26),false).updateUnsafe(new BigNumber(27),false).updateUnsafe(new BigNumber(28),false).updateUnsafe(new BigNumber(29),false).updateUnsafe(new BigNumber(30),false).updateUnsafe(new BigNumber(31),false).updateUnsafe(new BigNumber(32),false).updateUnsafe(new BigNumber(33),false).updateUnsafe(new BigNumber(34),false).updateUnsafe(new BigNumber(35),false).updateUnsafe(new BigNumber(36),false).updateUnsafe(new BigNumber(37),false).updateUnsafe(new BigNumber(38),false).updateUnsafe(new BigNumber(39),false).updateUnsafe(new BigNumber(40),false).updateUnsafe(new BigNumber(41),false).updateUnsafe(new BigNumber(42),false).updateUnsafe(new BigNumber(43),false).updateUnsafe(new BigNumber(44),false).updateUnsafe(new BigNumber(45),false).updateUnsafe(new BigNumber(46),false).updateUnsafe(new BigNumber(47),false).updateUnsafe(new BigNumber(48),false).updateUnsafe(new BigNumber(49),false).updateUnsafe(new BigNumber(50),false).updateUnsafe(new BigNumber(51),false).updateUnsafe(new BigNumber(52),false).updateUnsafe(new BigNumber(53),false).updateUnsafe(new BigNumber(54),false).updateUnsafe(new BigNumber(55),false).updateUnsafe(new BigNumber(56),false).updateUnsafe(new BigNumber(57),false).updateUnsafe(new BigNumber(58),false).updateUnsafe(_dafny.ZERO,false).updateUnsafe(_dafny.ONE,false).updateUnsafe(new BigNumber(2),false).updateUnsafe(new BigNumber(3),false).updateUnsafe(new BigNumber(4),false).updateUnsafe(new BigNumber(5),false).updateUnsafe(new BigNumber(6),false).updateUnsafe(new BigNumber(7),false).updateUnsafe(new BigNumber(8),false).updateUnsafe(new BigNumber(9),false).updateUnsafe(new BigNumber(10),false).updateUnsafe(new BigNumber(11),false).updateUnsafe(new BigNumber(12),false).updateUnsafe(new BigNumber(13),false).updateUnsafe(new BigNumber(14),false).updateUnsafe(new BigNumber(15),false).updateUnsafe(new BigNumber(16),false).updateUnsafe(new BigNumber(17),false).updateUnsafe(new BigNumber(18),false).updateUnsafe(new BigNumber(19),false).updateUnsafe(new BigNumber(20),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(56),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_292_v141).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(21),false).updateUnsafe(new BigNumber(22),false).updateUnsafe(new BigNumber(23),false).updateUnsafe(new BigNumber(24),false).updateUnsafe(new BigNumber(25),false).updateUnsafe(new BigNumber(26),false).updateUnsafe(new BigNumber(27),false).updateUnsafe(new BigNumber(28),false).updateUnsafe(new BigNumber(29),false).updateUnsafe(new BigNumber(30),false).updateUnsafe(new BigNumber(31),false).updateUnsafe(new BigNumber(32),false).updateUnsafe(new BigNumber(33),false).updateUnsafe(new BigNumber(34),false).updateUnsafe(new BigNumber(35),false).updateUnsafe(new BigNumber(36),false).updateUnsafe(new BigNumber(37),false).updateUnsafe(new BigNumber(38),false).updateUnsafe(new BigNumber(39),false).updateUnsafe(new BigNumber(40),false).updateUnsafe(new BigNumber(41),false).updateUnsafe(new BigNumber(42),false).updateUnsafe(new BigNumber(43),false).updateUnsafe(new BigNumber(44),false).updateUnsafe(new BigNumber(45),false).updateUnsafe(new BigNumber(46),false).updateUnsafe(new BigNumber(47),false).updateUnsafe(new BigNumber(48),false).updateUnsafe(new BigNumber(49),false).updateUnsafe(new BigNumber(50),false).updateUnsafe(new BigNumber(51),false).updateUnsafe(new BigNumber(52),false).updateUnsafe(new BigNumber(53),false).updateUnsafe(new BigNumber(54),false).updateUnsafe(new BigNumber(55),false).updateUnsafe(new BigNumber(56),false).updateUnsafe(new BigNumber(57),false).updateUnsafe(new BigNumber(58),false).updateUnsafe(_dafny.ZERO,false).updateUnsafe(_dafny.ONE,false).updateUnsafe(new BigNumber(2),false).updateUnsafe(new BigNumber(3),false).updateUnsafe(new BigNumber(4),false).updateUnsafe(new BigNumber(5),false).updateUnsafe(new BigNumber(6),false).updateUnsafe(new BigNumber(7),false).updateUnsafe(new BigNumber(8),false).updateUnsafe(new BigNumber(9),false).updateUnsafe(new BigNumber(10),false).updateUnsafe(new BigNumber(11),false).updateUnsafe(new BigNumber(12),false).updateUnsafe(new BigNumber(13),false).updateUnsafe(new BigNumber(14),false).updateUnsafe(new BigNumber(15),false).updateUnsafe(new BigNumber(16),false).updateUnsafe(new BigNumber(17),false).updateUnsafe(new BigNumber(18),false).updateUnsafe(new BigNumber(19),false).updateUnsafe(new BigNumber(20),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(56),true)),_dafny.Seq.UnicodeFromString("nalyols")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_293_v142).f24));
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
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1([]);
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
    static create_DC3(cf3, cf4, cf5) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC4(cf6, cf7) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D1(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC5(cf8) {
      let $dt = new D1(3);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf2 === other.cf2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3([], _dafny.ZERO, false);
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
    static create_DC7(cf10, cf11, cf12, cf13) {
      let $dt = new D2(0);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC6(cf9) {
      let $dt = new D2(1);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC8(cf14) {
      let $dt = new D2(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + this.cf9.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12 && this.cf13 === other.cf13;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), null, false);
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
    static create_DC10(cf16, cf17) {
      let $dt = new D3(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC11(cf18) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC9(cf15) {
      let $dt = new D3(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC12(cf19) {
      let $dt = new D3(3);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf18 === other.cf18;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf15 === other.cf15;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC14(cf21) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC13(cf20) {
      let $dt = new D4(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14(false);
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
    static create_DC16(cf23, cf24) {
      let $dt = new D5(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf22) {
      let $dt = new D5(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC17(cf25) {
      let $dt = new D5(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf22 === other.cf22;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC16(false, _dafny.Seq.of());
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
    static create_DC19(cf27, cf28, cf29) {
      let $dt = new D6(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC20() {
      let $dt = new D6(1);
      return $dt;
    }
    static create_DC18(cf26) {
      let $dt = new D6(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC20";
      } else if (this.$tag === 2) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC19(_dafny.Map.Empty, _dafny.Set.Empty, false);
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
    static create_DC21(cf30) {
      let $dt = new D7(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
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
    static create_DC23(cf32, cf33, cf34, cf35, cf36) {
      let $dt = new D8(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC24(cf37, cf38) {
      let $dt = new D8(1);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC25(cf39, cf40) {
      let $dt = new D8(2);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC22(cf31) {
      let $dt = new D8(3);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get is_DC22() { return this.$tag === 3; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf32) + ", " + this.cf33.toVerbatimString(true) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37) && this.cf38 === other.cf38;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf39 === other.cf39 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.MultiSet.Empty;
      this.f2 = false;
      this.f15 = [];
      this.f18 = false;
      this.f20 = _dafny.Seq.UnicodeFromString("");
      this.f23 = _dafny.ZERO;
      this._f0 = _dafny.ZERO;
      this._f3 = _dafny.Map.Empty;
      this._f4 = _dafny.ZERO;
      this._f5 = false;
      this._f6 = false;
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this._f9 = false;
      this._f10 = [];
      this._f11 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f12 = _dafny.ZERO;
      this._f13 = [];
      this._f14 = _dafny.Seq.UnicodeFromString("");
      this._f16 = _dafny.ZERO;
      this._f17 = [];
      this._f19 = false;
      this._f21 = _dafny.Seq.of();
      this._f22 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this).f18 = f18;
      (_this)._f19 = f19;
      (_this).f20 = f20;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      (_this).f23 = f23;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f3() {
      let _this = this;
      return _this._f3;
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
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f24 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f24) {
      let _this = this;
      (_this)._f24 = f24;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (((true) ? (new BigNumber(845)) : (new BigNumber((_dafny.Seq.UnicodeFromString("vwwnotdl")).length)))).multipliedBy((_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), function (_295_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-399),!((_this).f24))).length));
      }))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(!((_this).f24))).length), new BigNumber(980)))).cardinality())));
    };
    fm5(globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("lffyht");
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
