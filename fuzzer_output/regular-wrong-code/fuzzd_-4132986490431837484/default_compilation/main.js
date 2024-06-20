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
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false)));
    };
    static fm1(p0, p1, globalState) {
      return !(!(!(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(750)), function (_0_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("uhvt")))));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(false))).length), new BigNumber(242)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-242)))), ((true) ? (_dafny.Seq.of(new BigNumber(-647), new BigNumber((_dafny.Seq.UnicodeFromString("p")).length), new BigNumber(578), new BigNumber(227))) : (_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-573)), function (_1_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }))).cardinality()), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), function (_2_i1) {
        return new BigNumber((_dafny.Seq.of(new BigNumber(395), new BigNumber(-527))).length);
      })).length)))))));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      if (false) {
        return (_dafny.Set.fromElements(_dafny.Set.fromElements(!(false)))).Intersect(_dafny.Set.fromElements(_dafny.Set.fromElements((_module.D10.create_DC29(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber(-791))).length), false)).dtor_cf57), _dafny.Set.fromElements(false), _dafny.Set.fromElements(true, true), _dafny.Set.fromElements(false), _dafny.Set.fromElements(true)));
      } else {
        return _dafny.Set.fromElements(_dafny.Set.fromElements(false, false), _dafny.Set.fromElements(false));
      }
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return (((true) ? (_dafny.Set.fromElements(true, true, true)) : (_dafny.Set.fromElements(false)))).Difference((_dafny.Set.fromElements(true, true, true, false)).Intersect(_dafny.Set.fromElements(false)));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC6((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(207), new BigNumber(-76)))), (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(678),_dafny.Seq.of(true))).length))).length), new BigNumber(590)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("tx")).length)), _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(25)), function (_3_i0) {
  return new BigNumber(-348);
})))).IsSubsetOf(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(732)))), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), function (_4_i1) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length);
})).length), false, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Seq.of(false, false)).length))).length)));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      if ((!(false)) || (!(!(false)))) {
        if (false) {
          return _module.D3.create_DC8(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)))), !(true), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-603),false)).length)));
        } else {
          return _module.D3.create_DC8(_dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0))), false, new BigNumber(251));
        }
      } else {
        return _module.D3.create_DC8(_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0))), !(false), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()));
      }
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('u'.codePointAt(0));
    };
    static fm15(p0, p1, globalState) {
      let _source0 = _module.D4.create_DC10(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-492))));
      if (_source0.is_DC11) {
        let _5___mcc_h0 = (_source0).cf23;
        let _6___mcc_h1 = (_source0).cf24;
        let _7___mcc_h2 = (_source0).cf25;
        let _8_cf25 = _7___mcc_h2;
        let _9_cf24 = _6___mcc_h1;
        let _10_cf23 = _5___mcc_h0;
        return _module.D0.create_DC1(_dafny.MultiSet.fromElements(_9_cf24));
      } else {
        let _11___mcc_h3 = (_source0).cf22;
        let _12_cf22 = _11___mcc_h3;
        return _module.D0.create_DC1(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)));
      }
    };
    static fm18(p0, p1, globalState) {
      let _source1 = _module.D21.create_DC52(new _dafny.CodePoint('k'.codePointAt(0)));
      if (_source1.is_DC52) {
        let _13___mcc_h0 = (_source1).cf89;
        let _14_cf89 = _13___mcc_h0;
        return _module.D0.create_DC0(true);
      } else if (_source1.is_DC51) {
        let _15___mcc_h1 = (_source1).cf88;
        let _16_cf88 = _15___mcc_h1;
        return _module.D0.create_DC0(false);
      } else {
        let _17___mcc_h2 = (_source1).cf90;
        let _18_cf90 = _17___mcc_h2;
        if (!(false)) {
          return _module.D0.create_DC0(!(false));
        } else {
          return _module.D0.create_DC0(true);
        }
      }
    };
    static fm21(p0, p1, globalState) {
      if (!(_dafny.ONE).isEqualTo((_dafny.ZERO).minus(new BigNumber(-669)))) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }
    };
    static fm23(p0, p1, globalState) {
      return _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_module.D4.create_DC10(_dafny.MultiSet.fromElements(new BigNumber(656))), _module.D4.create_DC10(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-781)), function (_19_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length), new BigNumber(413))))).length));
    };
    static fm24(p0, p1, globalState) {
      let _source2 = _module.D20.create_DC49(false, new _dafny.CodePoint('w'.codePointAt(0)));
      if (_source2.is_DC48) {
        let _20___mcc_h0 = (_source2).cf82;
        let _21___mcc_h1 = (_source2).cf83;
        let _22___mcc_h2 = (_source2).cf84;
        let _23_cf84 = _22___mcc_h2;
        let _24_cf83 = _21___mcc_h1;
        let _25_cf82 = _20___mcc_h0;
        return (new BigNumber(470)).minus(new BigNumber(613));
      } else if (_source2.is_DC49) {
        let _26___mcc_h3 = (_source2).cf85;
        let _27___mcc_h4 = (_source2).cf86;
        let _28_cf86 = _27___mcc_h4;
        let _29_cf85 = _26___mcc_h3;
        return _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_29_cf85)).cardinality()), new BigNumber(-440))).length), new BigNumber(609));
      } else if (_source2.is_DC47) {
        let _30___mcc_h5 = (_source2).cf81;
        let _31_cf81 = _30___mcc_h5;
        return new BigNumber(132);
      } else {
        let _32___mcc_h6 = (_source2).cf87;
        let _33_cf87 = _32___mcc_h6;
        return new BigNumber(626);
      }
    };
    static fm25(globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-198))).length))).length), new BigNumber(496))).Union((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(550), new BigNumber(-304))) {
          let _34_v0 = _compr_0;
          if (((new BigNumber(550)).isLessThanOrEqualTo(_34_v0)) && ((_34_v0).isLessThan(new BigNumber(-304)))) {
            _coll0.push([_module.__default.safeModuloInt(_34_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-980)), function (_35_i0) {
              return new _dafny.CodePoint('t'.codePointAt(0));
            })).length)),_dafny.Map.Empty.slice().updateUnsafe(false,false)]);
          }
        }
        return _coll0;
      }()).length), new BigNumber((_dafny.Seq.of(new BigNumber(165), new BigNumber(557))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).length))).Difference(_dafny.MultiSet.fromElements(new BigNumber(297), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(68),true)).length))).length)))));
    };
    static fm26(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(true))).Merge((_module.D26.create_DC62(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(false, true, false, !(true))))).dtor_cf100);
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return (function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),!(true)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true))).length),false))).Elements) {
          let _36_v0 = _compr_1;
          if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),!(true)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true))).length),false))).contains(_36_v0)) {
            _coll1.push([_36_v0,_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))]);
          }
        }
        return _coll1;
      }()).Merge((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(484)), function (_37_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(727),false);
        })).Elements) {
          let _38_v1 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(484)), function (_37_i0) {
            return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(727),false);
          }), _38_v1)) {
            _coll2.push([_38_v1,_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0))))]);
          }
        }
        return _coll2;
      }()).Merge(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(107),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-709),(_module.D10.create_DC29(new BigNumber(433), true)).dtor_cf57))).Elements) {
          let _39_v2 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(107),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-709),(_module.D10.create_DC29(new BigNumber(433), true)).dtor_cf57)), _39_v2)) {
            _coll3.push([_39_v2,_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)))]);
          }
        }
        return _coll3;
      }()));
    };
    static fm32(p0, globalState) {
      return (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(323), _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(826))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("cvfvdpjmg")).length))));
    };
    static fm33(p0, globalState) {
      return (_dafny.MultiSet.fromElements(true, false, true, true)).Difference(_dafny.MultiSet.fromElements(false));
    };
    static fm34(p0, p1, p2, p3, globalState) {
      let _source3 = _module.D8.create_DC22((_dafny.ZERO).minus(new BigNumber(-494)), new BigNumber((_dafny.Seq.of(new BigNumber(772))).length));
      if (_source3.is_DC22) {
        let _40___mcc_h0 = (_source3).cf41;
        let _41___mcc_h1 = (_source3).cf42;
        let _42_cf42 = _41___mcc_h1;
        let _43_cf41 = _40___mcc_h0;
        return new _dafny.CodePoint('w'.codePointAt(0));
      } else if (_source3.is_DC23) {
        let _44___mcc_h2 = (_source3).cf43;
        let _45___mcc_h3 = (_source3).cf44;
        let _46___mcc_h4 = (_source3).cf45;
        let _47_cf45 = _46___mcc_h4;
        let _48_cf44 = _45___mcc_h3;
        let _49_cf43 = _44___mcc_h2;
        if (true) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }
      } else if (_source3.is_DC24) {
        let _50___mcc_h5 = (_source3).cf46;
        let _51___mcc_h6 = (_source3).cf47;
        let _52___mcc_h7 = (_source3).cf48;
        let _53___mcc_h8 = (_source3).cf49;
        let _54___mcc_h9 = (_source3).cf50;
        let _55_cf50 = _54___mcc_h9;
        let _56_cf49 = _53___mcc_h8;
        let _57_cf48 = _52___mcc_h7;
        let _58_cf47 = _51___mcc_h6;
        let _59_cf46 = _50___mcc_h5;
        return new _dafny.CodePoint('r'.codePointAt(0));
      } else if (_source3.is_DC21) {
        let _60___mcc_h10 = (_source3).cf40;
        let _61_cf40 = _60___mcc_h10;
        if (true) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }
      } else {
        let _62___mcc_h11 = (_source3).cf51;
        let _63_cf51 = _62___mcc_h11;
        return new _dafny.CodePoint('l'.codePointAt(0));
      }
    };
    static fm35(p0, globalState) {
      let _source4 = _module.D8.create_DC24(true, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-792)), function (_64_i0) {
  return new _dafny.CodePoint('a'.codePointAt(0));
})).length), _dafny.Seq.of(true), false, _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)),new BigNumber(905))).length)));
      if (_source4.is_DC22) {
        let _65___mcc_h0 = (_source4).cf41;
        let _66___mcc_h1 = (_source4).cf42;
        let _67_cf42 = _66___mcc_h1;
        let _68_cf41 = _65___mcc_h0;
        return ((_module.D8.create_DC21(function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), function (_69_i1) {
    return _module.D0.create_DC1(_dafny.MultiSet.fromElements(true));
  })).Elements) {
    let _70_v0 = _compr_4;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), function (_69_i1) {
      return _module.D0.create_DC1(_dafny.MultiSet.fromElements(true));
    }), _70_v0)) {
      _coll4.push([_70_v0,false]);
    }
  }
  return _coll4;
}())).dtor_cf40).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false))),false));
      } else if (_source4.is_DC23) {
        let _71___mcc_h2 = (_source4).cf43;
        let _72___mcc_h3 = (_source4).cf44;
        let _73___mcc_h4 = (_source4).cf45;
        let _74_cf45 = _73___mcc_h4;
        let _75_cf44 = _72___mcc_h3;
        let _76_cf43 = _71___mcc_h2;
        return _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(false)),false);
      } else if (_source4.is_DC24) {
        let _77___mcc_h5 = (_source4).cf46;
        let _78___mcc_h6 = (_source4).cf47;
        let _79___mcc_h7 = (_source4).cf48;
        let _80___mcc_h8 = (_source4).cf49;
        let _81___mcc_h9 = (_source4).cf50;
        let _82_cf50 = _81___mcc_h9;
        let _83_cf49 = _80___mcc_h8;
        let _84_cf48 = _79___mcc_h7;
        let _85_cf47 = _78___mcc_h6;
        let _86_cf46 = _77___mcc_h5;
        if (_86_cf46) {
          return _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(_83_cf49, _83_cf49)),_86_cf46);
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.FromArray(_84_cf48)),_83_cf49);
        }
      } else if (_source4.is_DC21) {
        let _87___mcc_h10 = (_source4).cf40;
        let _88_cf40 = _87___mcc_h10;
        return _88_cf40;
      } else {
        let _89___mcc_h11 = (_source4).cf51;
        let _90_cf51 = _89___mcc_h11;
        return (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(true)),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(false)),true));
      }
    };
    static fm36(p0, p1, p2, globalState) {
      return _module.D0.create_DC1(_dafny.MultiSet.fromElements(false));
    };
    static fm37(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,false));
    };
    static fm38(p0, p1, globalState) {
      return (function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("xsk")).length))).Elements) {
          let _91_v0 = _compr_5;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("xsk")).length)), _91_v0)) {
            _coll5.push([_module.__default.safeDivisionInt(_91_v0, new BigNumber(-244)),(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-729)))]);
          }
        }
        return _coll5;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(312),new BigNumber(669)));
    };
    static fm39(p0, p1, p2, globalState) {
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC7(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-880)), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-879)))),_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0))))).length);
    };
    static fm40(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(!(true)), false, true, false))).Union(_dafny.MultiSet.fromElements(false)));
    };
    static fm41(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('d'.codePointAt(0));
    };
    static fm42(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(988),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(59)), function (_92_i0) {
        return new BigNumber(202);
      })).length))).length))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true, !(false), true)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(5)), function (_93_i1) {
        return (_dafny.ZERO).minus(new BigNumber(-690));
      })));
    };
    static fm43(globalState) {
      let _source5 = _module.D8.create_DC25(_module.D8.create_DC24(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gbqytxrv"),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).length))).length), _dafny.Seq.of(true), !(true), _dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), function (_94_i0) {
  return new BigNumber(-972);
})));
      if (_source5.is_DC22) {
        let _95___mcc_h0 = (_source5).cf41;
        let _96___mcc_h1 = (_source5).cf42;
        let _97_cf42 = _96___mcc_h1;
        let _98_cf41 = _95___mcc_h0;
        return _dafny.MultiSet.fromElements(new BigNumber(-766), new BigNumber(342), _98_cf41, _97_cf42);
      } else if (_source5.is_DC23) {
        let _99___mcc_h2 = (_source5).cf43;
        let _100___mcc_h3 = (_source5).cf44;
        let _101___mcc_h4 = (_source5).cf45;
        let _102_cf45 = _101___mcc_h4;
        let _103_cf44 = _100___mcc_h3;
        let _104_cf43 = _99___mcc_h2;
        return (_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-246)))))).Union(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(427)), function (_105_i1) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).length))));
      } else if (_source5.is_DC24) {
        let _106___mcc_h5 = (_source5).cf46;
        let _107___mcc_h6 = (_source5).cf47;
        let _108___mcc_h7 = (_source5).cf48;
        let _109___mcc_h8 = (_source5).cf49;
        let _110___mcc_h9 = (_source5).cf50;
        let _111_cf50 = _110___mcc_h9;
        let _112_cf49 = _109___mcc_h8;
        let _113_cf48 = _108___mcc_h7;
        let _114_cf47 = _107___mcc_h6;
        let _115_cf46 = _106___mcc_h5;
        return _dafny.MultiSet.fromElements(_114_cf47, _114_cf47, _114_cf47, _114_cf47);
      } else if (_source5.is_DC21) {
        let _116___mcc_h10 = (_source5).cf40;
        let _117_cf40 = _116___mcc_h10;
        return (_dafny.MultiSet.fromElements(new BigNumber(832), new BigNumber((_dafny.Set.fromElements(!(false))).length), new BigNumber(789), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("fuonpos")).length),false)).length))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(351)), function (_118_i2) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length)));
      } else {
        let _119___mcc_h11 = (_source5).cf51;
        let _120_cf51 = _119___mcc_h11;
        return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("schrvre")).length), new BigNumber(887))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("yrovpuj")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, false)).length))));
      }
    };
    static fm45(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements((new BigNumber((_dafny.Seq.of(true)).length)).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("g")).length)));
    };
    static fm46(p0, p1, p2, globalState) {
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat((_module.D16.create_DC40(_dafny.Seq.UnicodeFromString("llk"), new BigNumber(894))).dtor_cf73, _dafny.Seq.UnicodeFromString("hf")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(401)), function (_121_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("lacfjew")))).length);
    };
    static fm47(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(668),false)).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-920)), function (_122_i0) {
        return (_module.D21.create_DC52(new _dafny.CodePoint('j'.codePointAt(0)))).dtor_cf89;
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(!((_module.D10.create_DC31(false, !(!(true)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("bbpvswxo")).length), new BigNumber(340))).length), new BigNumber((_dafny.Set.fromElements(true, true)).length), new BigNumber(481), new BigNumber((function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(-128), new BigNumber(591))) {
    let _123_v0 = _compr_6;
    if (((new BigNumber(-128)).isLessThanOrEqualTo(_123_v0)) && ((_123_v0).isLessThan(new BigNumber(591)))) {
      _coll6.push([(_123_v0).multipliedBy(new BigNumber(-966)),true]);
    }
  }
  return _coll6;
}()).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("pnyamwon")).length), new BigNumber((_dafny.Seq.UnicodeFromString("kd")).length))).length)))).length))).dtor_cf61), false))).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(396)), function (_124_i1) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })));
    };
    static fm48(p0, p1, p2, globalState) {
      if (_dafny.areEqual(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(937)), function (_125_i0) {
        return new BigNumber(5);
      }), _dafny.Seq.of(new BigNumber(753)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(371)), function (_126_i1) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-478)), function (_127_i2) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length));
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(111)), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(916), new BigNumber(995))).length))), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), function (_128_i3) {
        return new BigNumber(908);
      })).length)), _dafny.MultiSet.fromElements(new BigNumber(989)), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(886))))).length))), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(382)), _dafny.Seq.of(new BigNumber(-56), new BigNumber(776)), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), new BigNumber(36))))) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }
    };
    static fm49(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.FromArray(((false) ? (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(!(true)))).length))) : (_dafny.Seq.of(new BigNumber(220), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("yxr")).length)),true)).length)))))).Union((_dafny.MultiSet.fromElements(new BigNumber(18), new BigNumber(-111))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("n")).length), new BigNumber(478), new BigNumber(708), new BigNumber(-3))));
    };
    static fm50(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(112), new BigNumber(890))) {
          let _129_v0 = _compr_7;
          if (((new BigNumber(112)).isLessThanOrEqualTo(_129_v0)) && ((_129_v0).isLessThan(new BigNumber(890)))) {
            _coll7.push([(_129_v0).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qxkxsgkqe")).length),new BigNumber((_dafny.Set.fromElements(false)).length))).length)),true]);
          }
        }
        return _coll7;
      }()).length),new BigNumber((_dafny.Seq.UnicodeFromString("jpyifw")).length))));
    };
    static fm51(p0, p1, globalState) {
      if (!(!((_dafny.ZERO).minus(new BigNumber(-184))).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(284),false)).length)))) {
        return _dafny.Seq.of(_dafny.Seq.of(true, true, true, !((_module.D26.create_DC63(new BigNumber(-626), new BigNumber(673), true)).dtor_cf103)));
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(397)), function (_130_i0) {
          return _dafny.Seq.of(true);
        }), _dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(!(false)), _dafny.Seq.of(false)));
      }
    };
    static fm52(p0, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.of(false))).Intersect(function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of (_dafny.Seq.of(_dafny.Seq.of(true, true))).Elements) {
          let _131_v0 = _compr_8;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(true, true)), _131_v0)) {
            _coll8.add(_131_v0);
          }
        }
        return _coll8;
      }())).Difference((_dafny.Set.fromElements(_dafny.Seq.of(false, false), _dafny.Seq.of(!(false)))).Difference(_dafny.Set.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(!(false), false), _dafny.Seq.of(true, false))));
    };
    static fm53(p0, p1, globalState) {
      return _module.D10.create_DC31(!(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("cvutgcy"), _dafny.Seq.UnicodeFromString("g"))), !(false) || (true), (new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(415)), function (_132_i0) {
  return new _dafny.CodePoint('p'.codePointAt(0));
})).length)));
    };
    static fm54(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(746)), function (_133_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("vqw")).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("pqg")).length)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length)))).Difference(function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(276)), function (_134_i1) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(272),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(639)), function (_135_i2) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        })).length))).length))).length)))).Elements) {
          let _136_v0 = _compr_9;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(276)), function (_134_i1) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(272),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(639)), function (_135_i2) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length))).length))).length)))).contains(_136_v0)) {
            _coll9.add(_136_v0);
          }
        }
        return _coll9;
      }());
    };
    static fm55(p0, p1, p2, globalState) {
      let _source6 = _module.D10.create_DC29(new BigNumber(100), true);
      if (_source6.is_DC29) {
        let _137___mcc_h0 = (_source6).cf56;
        let _138___mcc_h1 = (_source6).cf57;
        let _139_cf57 = _138___mcc_h1;
        let _140_cf56 = _137___mcc_h0;
        return _dafny.Set.fromElements(false);
      } else if (_source6.is_DC30) {
        let _141___mcc_h2 = (_source6).cf58;
        let _142___mcc_h3 = (_source6).cf59;
        let _143_cf59 = _142___mcc_h3;
        let _144_cf58 = _141___mcc_h2;
        return _dafny.Set.fromElements(false);
      } else if (_source6.is_DC31) {
        let _145___mcc_h4 = (_source6).cf60;
        let _146___mcc_h5 = (_source6).cf61;
        let _147___mcc_h6 = (_source6).cf62;
        let _148_cf62 = _147___mcc_h6;
        let _149_cf61 = _146___mcc_h5;
        let _150_cf60 = _145___mcc_h4;
        if (!(_150_cf60)) {
          return _dafny.Set.fromElements(_149_cf61);
        } else {
          return _dafny.Set.fromElements(_149_cf61);
        }
      } else if (_source6.is_DC28) {
        let _151___mcc_h7 = (_source6).cf55;
        let _152_cf55 = _151___mcc_h7;
        return _dafny.Set.fromElements((_152_cf55).f43, (_152_cf55).f43, (_152_cf55).f43, !((_152_cf55).f43), (_152_cf55).f43);
      } else {
        let _153___mcc_h8 = (_source6).cf63;
        let _154_cf63 = _153___mcc_h8;
        return _dafny.Set.fromElements(false, !(true));
      }
    };
    static fm56(p0, p1, p2, globalState) {
      return _module.D2.create_DC6(_module.__default.safeDivisionInt(new BigNumber(875), new BigNumber(-348)), ((!(true)) ? (true) : (!(!(!(!(true)))))), new BigNumber(-364), !((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(232),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length))).length)).isLessThan(new BigNumber(55))), new BigNumber(667));
    };
    static fm57(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(833),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("bcbb")).length),false));
    };
    static fm58(p0, p1, globalState) {
      return _module.D14.create_DC37(_module.__default.safeModuloInt((_module.D26.create_DC63(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((function () {
  let _coll10 = new _dafny.Set();
  for (const _compr_10 of (_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)))).Elements) {
    let _155_v0 = _compr_10;
    if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0))), _155_v0)) {
      _coll10.add(_155_v0);
    }
  }
  return _coll10;
}()).length))).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(70))).length), true)).dtor_cf101, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(323)), function (_156_i0) {
  return new _dafny.CodePoint('b'.codePointAt(0));
})).length)), true, _module.__default.safeDivisionInt(new BigNumber(723), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-464),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),(_module.D8.create_DC22(new BigNumber((_dafny.Seq.of(new BigNumber(-737))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).dtor_cf41)).length))).length)));
    };
    static fm59(p0, p1, p2, globalState) {
      return ((function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(991)), function (_157_i0) {
          return new BigNumber((_dafny.Seq.of(false)).length);
        })).Elements) {
          let _158_v0 = _compr_11;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(991)), function (_157_i0) {
            return new BigNumber((_dafny.Seq.of(false)).length);
          }), _158_v0)) {
            _coll11.push([_module.__default.safeModuloInt(_158_v0, new BigNumber((_dafny.Set.fromElements(false)).length)),_dafny.Set.fromElements(new BigNumber(326))]);
          }
        }
        return _coll11;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-574),_dafny.Set.fromElements(new BigNumber(88))))).Merge(function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ekgojohii")).length), new BigNumber(374), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, false)).length), new BigNumber(185), new BigNumber(958))).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(976)), function (_159_i1) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        })).length))).length), new BigNumber(248), new BigNumber(565))).Elements) {
          let _160_v1 = _compr_12;
          if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ekgojohii")).length), new BigNumber(374), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, false)).length), new BigNumber(185), new BigNumber(958))).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(976)), function (_159_i1) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          })).length))).length), new BigNumber(248), new BigNumber(565))).contains(_160_v1)) {
            _coll12.push([(_160_v1).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("vrr")).length)),function () {
              let _coll13 = new _dafny.Set();
              for (const _compr_13 of _dafny.IntegerRange(new BigNumber(901), new BigNumber(58))) {
                let _161_v2 = _compr_13;
                if (((new BigNumber(901)).isLessThanOrEqualTo(_161_v2)) && ((_161_v2).isLessThan(new BigNumber(58)))) {
                  _coll13.add((_161_v2).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(789)), function (_162_i2) {
                    return new _dafny.CodePoint('d'.codePointAt(0));
                  })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true))).cardinality()),new BigNumber(-470))).length))).cardinality())));
                }
              }
              return _coll13;
            }()]);
          }
        }
        return _coll12;
      }());
    };
    static fm60(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(110))).length))).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("vcr")).length), new BigNumber((_dafny.Set.fromElements((_module.D22.create_DC55(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))).length), new BigNumber(372), true)).dtor_cf93)).length), new BigNumber(-613)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Create(_module.__default.abs(new BigNumber(342)), function (_163_i0) {
        return new BigNumber(543);
      })))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(746)), function (_164_i1) {
        return (_dafny.ZERO).minus(new BigNumber(-801));
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(123)), function (_165_i2) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("vgo")).length);
      }))));
    };
    static fm61(p0, p1, p2, globalState) {
      let _source7 = _module.D16.create_DC41(_module.D16.create_DC39(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(522)), function (_166_i0) {
  return new BigNumber(707);
}), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),false)).length), new BigNumber(395), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-512)), function (_167_i1) {
  return new _dafny.CodePoint('u'.codePointAt(0));
})).length), new BigNumber((_dafny.Seq.UnicodeFromString("gkgia")).length), new BigNumber(806)), _dafny.Seq.of(new BigNumber(124)))));
      if (_source7.is_DC40) {
        let _168___mcc_h0 = (_source7).cf73;
        let _169___mcc_h1 = (_source7).cf74;
        let _170_cf74 = _169___mcc_h1;
        let _171_cf73 = _168___mcc_h0;
        return _module.D8.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(false)),false));
      } else if (_source7.is_DC39) {
        let _172___mcc_h2 = (_source7).cf72;
        let _173_cf72 = _172___mcc_h2;
        return _module.D8.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(!(false), !(false))),!(true)));
      } else {
        let _174___mcc_h3 = (_source7).cf75;
        let _175_cf75 = _174___mcc_h3;
        return _module.D8.create_DC21(function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(!(false))),new BigNumber(618))).Keys.Elements) {
    let _176_v0 = _compr_14;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(!(false))),new BigNumber(618))).contains(_176_v0)) {
      _coll14.push([_176_v0,true]);
    }
  }
  return _coll14;
}());
      }
    };
    static fm62(p0, p1, p2, p3, globalState) {
      return _module.D21.create_DC51(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(285),_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(902))).length))));
    };
    static fm63(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(687)), function (_177_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length)), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(385)), function (_178_i1) {
        return _dafny.Seq.UnicodeFromString("eix");
      })).length))));
    };
    static fm64(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(279),(_module.D16.create_DC40(_dafny.Seq.UnicodeFromString("stc"), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).dtor_cf73)).Keys.Elements) {
          let _179_v0 = _compr_15;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(279),(_module.D16.create_DC40(_dafny.Seq.UnicodeFromString("stc"), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).dtor_cf73)).contains(_179_v0)) {
            _coll15.push([(_179_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)),_module.D0.create_DC1(_dafny.MultiSet.fromElements(false, true))]);
          }
        }
        return _coll15;
      }();
    };
    static fm65(p0, p1, p2, globalState) {
      return (((!(true)) ? (_dafny.Map.Empty.slice().updateUnsafe(true,_module.D8.create_DC22(new BigNumber(-188), _dafny.ZERO))) : (_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),_module.D8.create_DC22(new BigNumber((_dafny.Seq.UnicodeFromString("n")).length), new BigNumber(398)))))).Merge((_module.D29.create_DC67(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D8.create_DC22((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)), new BigNumber(792))))).dtor_cf107);
    };
    static fm66(globalState) {
      return _module.D6.create_DC15(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("fflctru")).length)), _dafny.Seq.of(new BigNumber(137))));
    };
    static fm67(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,true)), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(!(!(!(true))),false))),_dafny.Map.Empty.slice().updateUnsafe(true,!(false)));
    };
    static fm68(p0, p1, p2, globalState) {
      return _module.D20.create_DC49(_dafny.areEqual(_dafny.Seq.of(!(false)), _dafny.Seq.of(false, false, true, false)), new _dafny.CodePoint('q'.codePointAt(0)));
    };
    static fm69(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_module.D16.create_DC39(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(814)), _dafny.Seq.of(new BigNumber(920)))),_dafny.Seq.UnicodeFromString("wrtq"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D16.create_DC39(_dafny.Set.fromElements((_module.D8.create_DC24(false, new BigNumber((function () {
  let _coll16 = new _dafny.Map();
  for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-784), new BigNumber(969))) {
    let _180_v0 = _compr_16;
    if (((new BigNumber(-784)).isLessThanOrEqualTo(_180_v0)) && ((_180_v0).isLessThan(new BigNumber(969)))) {
      _coll16.push([(_180_v0).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(689), new BigNumber(252))).length)),false]);
    }
  }
  return _coll16;
}()).length), _dafny.Seq.of(false), !(!(true)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length)))).dtor_cf50, _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), function (_181_i0) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("neweldmjf")).length),new BigNumber(783))).length);
})).length)))),_dafny.Seq.UnicodeFromString("dbhwx")));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = _dafny.ZERO;
      let _182_v0;
      _182_v0 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(555),_module.__default.safeModuloInt(p2, p2));
      _182_v0 = _182_v0;
      let _183_v1;
      _183_v1 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
      let _184_v2;
      _184_v2 = new _dafny.CodePoint('y'.codePointAt(0));
      let _185_v3;
      _185_v3 = _module.D20.create_DC49(p3, _184_v2);
      let _186_v4;
      _186_v4 = _dafny.MultiSet.fromElements(_185_v3);
      let _187_v5;
      _187_v5 = _dafny.MultiSet.fromElements(p1);
      let _188_v6;
      _188_v6 = _dafny.Set.fromElements((((_187_v5).contains(p3)) ? ((_187_v5).get(p3)) : (p2)));
      let _hi0 = new BigNumber(312);
      for (let _189_i0 = ((_dafny.ZERO).minus(new BigNumber((_183_v1).length))).multipliedBy(_module.__default.fm46((((_186_v4).contains(_module.__default.fm68(p3, p1, p1, globalState))) ? ((_186_v4).get(_module.__default.fm68(p3, p1, p1, globalState))) : (p2)), new BigNumber((_188_v6).length), p1, globalState)); _189_i0.isLessThan(_hi0); _189_i0 = _189_i0.plus(_dafny.ONE)) {
        let _190_v7;
        let _init0 = ((_191_p2, _192_p3, _193_p1) => function (_194_i1) {
          return _module.D2.create_DC6((_dafny.ZERO).minus(_191_p2), _192_p3, new BigNumber(597), _193_p1, _191_p2);
        })(p2, p3, p1);
        let _nw0 = Array((new BigNumber(15)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _190_v7 = _nw0;
        let _195_v8;
        _195_v8 = _dafny.MultiSet.fromElements(_190_v7, _190_v7);
        let _196_v9;
        _196_v9 = _dafny.Seq.UnicodeFromString("kuxvyifu");
        (globalState).f26 = (_dafny.MultiSet.fromElements(_190_v7)).IsDisjointFrom((_195_v8).update(_190_v7, _module.__default.abs(new BigNumber((_196_v9).length))));
        let _197_v10;
        _197_v10 = _dafny.MultiSet.fromElements(_189_i0);
        let _198_v11;
        let _nw1 = Array((new BigNumber(16)).toNumber());
        _nw1[(_dafny.ZERO).toNumber()] = new BigNumber(537);
        _nw1[(_dafny.ONE).toNumber()] = _189_i0;
        _nw1[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Set.fromElements(p2)).length);
        _nw1[(new BigNumber(3)).toNumber()] = (p2).plus((_dafny.ZERO).minus(new BigNumber((_196_v9).length)));
        _nw1[(new BigNumber(4)).toNumber()] = ((p3) ? (_189_i0) : (new BigNumber((_197_v10).cardinality())));
        _nw1[(new BigNumber(5)).toNumber()] = p2;
        _nw1[(new BigNumber(6)).toNumber()] = p2;
        _nw1[(new BigNumber(7)).toNumber()] = p2;
        _nw1[(new BigNumber(8)).toNumber()] = _module.__default.fm39(p1, _197_v10, p3, globalState);
        _nw1[(new BigNumber(9)).toNumber()] = new BigNumber((_196_v9).length);
        _nw1[(new BigNumber(10)).toNumber()] = p2;
        _nw1[(new BigNumber(11)).toNumber()] = _189_i0;
        _nw1[(new BigNumber(12)).toNumber()] = p2;
        _nw1[(new BigNumber(13)).toNumber()] = (p2).multipliedBy(p2);
        _nw1[(new BigNumber(14)).toNumber()] = new BigNumber(((_182_v0).Merge(_182_v0)).length);
        _nw1[(new BigNumber(15)).toNumber()] = ((p0)[_module.__default.safeIndex(_189_i0, new BigNumber((p0).length))]).multipliedBy(p2);
        _198_v11 = _nw1;
        let _index0 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_198_v11).length));
        (_198_v11)[_index0] = _189_i0;
        let _index1 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_198_v11).length));
        (_198_v11)[_index1] = _189_i0;
        if (p1) {
          (globalState).f13 = p2;
          let _199_v12;
          _199_v12 = _dafny.Map.Empty.slice().updateUnsafe(p3,_184_v2);
          let _200_v13;
          _200_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,false);
          let _201_v14;
          _201_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_199_v12).length),_200_v13);
          let _202_v15;
          _202_v15 = _dafny.Seq.of(p1);
          let _203_v16;
          _203_v16 = _dafny.Seq.of(_202_v15);
          let _204_v17;
          let _nw2 = new _module.C9();
          _nw2.__ctor(new BigNumber((_187_v5).cardinality()), _201_v14, _184_v2, (new BigNumber((_196_v9).length)).multipliedBy(new BigNumber((_200_v13).length)), p1, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(p3), _202_v15), _module.__default.safeIndex(new BigNumber(708), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p3), _202_v15)).length)), p3), _dafny.Seq.Concat(_203_v16, _203_v16), _189_i0, (p2).isLessThan((_198_v11)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_198_v11).length))]), p3, p3);
          _204_v17 = _nw2;
          (globalState).f13 = _module.__default.safeDivisionInt(new BigNumber((_202_v15).length), ((_204_v17).f38).multipliedBy(_189_i0));
          (globalState).f26 = p3;
          let _205_v18;
          let _init1 = ((_206_p1) => function (_207_i2) {
            return _206_p1;
          })(p1);
          let _nw3 = Array((new BigNumber(17)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw3.length); _i0_1++) {
            _nw3[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _205_v18 = _nw3;
          let _208_v19;
          _208_v19 = _dafny.Map.Empty.slice().updateUnsafe(_196_v9,_205_v18);
          _208_v19 = (_208_v19).update(_dafny.Seq.Concat(_196_v9, _dafny.Seq.update(_196_v9, _module.__default.safeIndex(_189_i0, new BigNumber((_196_v9).length)), _184_v2)), _205_v18);
        } else {
          let _209_v20;
          _209_v20 = _dafny.Seq.of(p3, p3);
          let _210_v21;
          _210_v21 = _dafny.Seq.of(_209_v20);
          let _211_v22;
          let _nw4 = Array((new BigNumber(6)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _211_v22 = _nw4;
          let _212_v23;
          let _nw5 = new _module.C0();
          _nw5.__ctor(_211_v22);
          _212_v23 = _nw5;
          let _213_v24;
          _213_v24 = _dafny.Seq.of(_212_v23);
          let _214_v25;
          let _nw6 = new _module.C6();
          _nw6.__ctor(_184_v2, (((_182_v0).contains(new BigNumber(220))) ? ((_182_v0).get(new BigNumber(220))) : (new BigNumber((_dafny.Seq.UnicodeFromString("f")).length))), p3, _dafny.Seq.of(p1), _210_v21, _189_i0, p1, !_dafny.areEqual(_213_v24, _213_v24), p3);
          _214_v25 = _nw6;
          let _215_v26;
          let _nw7 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _215_v26 = _nw7;
          let _216_v27;
          let _nw8 = Array((new BigNumber(16)).toNumber()).fill(false);
          _216_v27 = _nw8;
          let _217_v28;
          _217_v28 = _dafny.Map.Empty.slice().updateUnsafe(_216_v27,_dafny.MultiSet.fromElements(p3, (((_183_v1).contains(_189_i0)) ? ((_183_v1).get(_189_i0)) : (false))));
          let _index2 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_215_v26).length));
          (_215_v26)[_index2] = _217_v28;
          let _index3 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_215_v26).length));
          let _rhs0 = _217_v28;
          let _rhs1 = _189_i0;
          let _lhs0 = _215_v26;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_215_v26).length));
          let _lhs2 = globalState;
          _lhs0[_lhs1] = _rhs0;
          _lhs2.f13 = _rhs1;
          _197_v10 = (_197_v10).Intersect((_module.D4.create_DC10(_197_v10)).dtor_cf22);
          let _218_v29;
          _218_v29 = _dafny.Set.fromElements(_dafny.Seq.of(_189_i0), p0, _dafny.Seq.of((_198_v11)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_198_v11).length))], p2), p0);
          let _219_v30;
          _219_v30 = _module.D16.create_DC39(_218_v29);
          let _220_v31;
          _220_v31 = _dafny.Map.Empty.slice().updateUnsafe(_219_v30,_196_v9);
          let _221_v32;
          _221_v32 = _module.D0.create_DC0(p3);
          let _index4 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_198_v11).length));
          let _rhs2 = _module.__default.fm69(p1, p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), ((_222_v2) => function (_223_i3) {
            return _222_v2;
          })(_184_v2))).length), !((_221_v32).dtor_cf0), globalState);
          let _rhs3 = (_dafny.ZERO).minus(p2);
          let _lhs3 = _198_v11;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_198_v11).length));
          _220_v31 = _rhs2;
          _lhs3[_lhs4] = _rhs3;
          let _224_v33;
          let _nw9 = new _module.C8();
          _nw9.__ctor(p3, _184_v2, new BigNumber(689), p1, _209_v20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(418)), ((_225_v20) => function (_226_i4) {
            return _225_v20;
          })(_209_v20)), p2, p3, true, p3);
          _224_v33 = _nw9;
          let _227_v34;
          _227_v34 = _dafny.Seq.of(_224_v33);
          let _228_v35;
          _228_v35 = _dafny.Seq.of(p0);
          let _229_v36;
          let _nw10 = new _module.C5();
          _nw10.__ctor(p1, new BigNumber((_227_v34).length), new BigNumber((_228_v35).length), (_214_v25).fm6(globalState), _209_v20, _210_v21, (_dafny.ZERO).minus(new BigNumber((_196_v9).length)), _224_v33.f28, true, true);
          _229_v36 = _nw10;
          let _230_v37;
          let _nw11 = Array((new BigNumber(3)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = _229_v36;
          _nw11[(_dafny.ONE).toNumber()] = _229_v36;
          _nw11[(new BigNumber(2)).toNumber()] = _229_v36;
          _230_v37 = _nw11;
          let _231_v38;
          _231_v38 = _dafny.MultiSet.fromElements(((_224_v33.f28) ? (_198_v11) : (_198_v11)), _198_v11, _198_v11, _198_v11, ((p3) ? (_198_v11) : (_198_v11)));
          let _rhs4 = new BigNumber((_231_v38).cardinality());
          let _rhs5 = _189_i0;
          let _rhs6 = _230_v37;
          let _rhs7 = (_229_v36).fm6(globalState);
          let _rhs8 = !(((_198_v11)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_198_v11).length))]).isLessThanOrEqualTo(_189_i0)) || ((_209_v20)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_224_v33.f29)).length), new BigNumber((_209_v20).length))]);
          let _lhs5 = globalState;
          let _lhs6 = globalState;
          r0 = _rhs4;
          r0 = _rhs5;
          _230_v37 = _rhs6;
          _lhs5.f26 = _rhs7;
          _lhs6.f12 = _rhs8;
        }
        let _232_v39;
        _232_v39 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        let _233_v40;
        _233_v40 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(p1, new BigNumber((_196_v9).length), globalState),p2)).update(p1, p2), (_232_v39).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,_189_i0)), (_232_v39).Merge(_232_v39), _232_v39);
        _233_v40 = _233_v40;
      }
      let _234_v41;
      _234_v41 = _dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)));
      let _235_v42;
      _235_v42 = _dafny.Seq.of(p3, p1, true);
      let _236_v43;
      _236_v43 = _dafny.MultiSet.fromElements(_234_v41);
      let _237_v44;
      _237_v44 = _dafny.Seq.of(_235_v42);
      let _238_v45;
      let _nw12 = new _module.C10();
      _nw12.__ctor(_184_v2, p2, p3, _235_v42, _237_v44, p2, (_235_v42)[_module.__default.safeIndex(p2, new BigNumber((_235_v42).length))], p1, p3);
      _238_v45 = _nw12;
      let _239_v46;
      let _nw13 = new _module.C4();
      _nw13.__ctor(p2, false, _dafny.Seq.of(!(p3), p3), _237_v44, p2, p1, p1, p3);
      _239_v46 = _nw13;
      let _240_v47;
      _240_v47 = _dafny.Map.Empty.slice().updateUnsafe(_238_v45,_239_v46);
      let _241_v48;
      _241_v48 = _dafny.Seq.of(_240_v47, _240_v47, _240_v47, _240_v47, _240_v47);
      let _242_v49;
      _242_v49 = _dafny.MultiSet.fromElements((_241_v48)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_241_v48).length))], _240_v47);
      let _243_v50;
      let _nw14 = Array((new BigNumber(23)).toNumber());
      _nw14[(_dafny.ZERO).toNumber()] = p1;
      _nw14[(_dafny.ONE).toNumber()] = ((p1) ? (true) : (p3));
      _nw14[(new BigNumber(2)).toNumber()] = !(p1) || (p3);
      _nw14[(new BigNumber(3)).toNumber()] = (_183_v1).contains(_module.__default.fm39(!(p1), _dafny.MultiSet.fromElements(p2, p2, p2, (_dafny.ZERO).minus(p2), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(574)), ((_244_v2) => function (_245_i5) {
        return _244_v2;
      })(_184_v2))).length)), p1, globalState));
      _nw14[(new BigNumber(4)).toNumber()] = p1;
      _nw14[(new BigNumber(5)).toNumber()] = p3;
      _nw14[(new BigNumber(6)).toNumber()] = p3;
      _nw14[(new BigNumber(7)).toNumber()] = p3;
      _nw14[(new BigNumber(8)).toNumber()] = p3;
      _nw14[(new BigNumber(9)).toNumber()] = (p3) && (p1);
      _nw14[(new BigNumber(10)).toNumber()] = (p3) === (p3);
      _nw14[(new BigNumber(11)).toNumber()] = p3;
      _nw14[(new BigNumber(12)).toNumber()] = p3;
      _nw14[(new BigNumber(13)).toNumber()] = true;
      _nw14[(new BigNumber(14)).toNumber()] = p1;
      _nw14[(new BigNumber(15)).toNumber()] = (_234_v41).contains(_184_v2);
      _nw14[(new BigNumber(16)).toNumber()] = p1;
      _nw14[(new BigNumber(17)).toNumber()] = (((_183_v1).contains((_dafny.ZERO).minus(new BigNumber((_235_v42).length)))) ? ((_183_v1).get((_dafny.ZERO).minus(new BigNumber((_235_v42).length)))) : (p1));
      _nw14[(new BigNumber(18)).toNumber()] = (p2).isLessThan(p2);
      _nw14[(new BigNumber(19)).toNumber()] = !((_236_v43).IsDisjointFrom(_236_v43));
      _nw14[(new BigNumber(20)).toNumber()] = p3;
      _nw14[(new BigNumber(21)).toNumber()] = p1;
      _nw14[(new BigNumber(22)).toNumber()] = (_242_v49).equals(_242_v49);
      _243_v50 = _nw14;
      let _index5 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_243_v50).length));
      (_243_v50)[_index5] = _239_v46.f29;
      let _246_v51;
      _246_v51 = _dafny.Seq.UnicodeFromString("ujnosv");
      let _247_v52;
      _247_v52 = _dafny.MultiSet.fromElements(_246_v51);
      let _248_v53;
      _248_v53 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(518)), ((_249_v51) => function (_250_i6) {
        return _249_v51;
      })(_246_v51)));
      let _251_v54;
      _251_v54 = _dafny.MultiSet.fromElements(p2);
      let _252_v55;
      let _nw15 = new _module.C6();
      _nw15.__ctor(_184_v2, (_239_v46).f30, p3, _235_v42, (_239_v46).f33, p2, _239_v46.f35, _239_v46.f35, p3);
      _252_v55 = _nw15;
      let _253_v56;
      _253_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_246_v51).length),_252_v55);
      let _254_v57;
      _254_v57 = _dafny.Seq.of((((_253_v56).contains(new BigNumber(330))) ? ((_253_v56).get(new BigNumber(330))) : (_252_v55)));
      let _255_v58;
      _255_v58 = _dafny.Map.Empty.slice().updateUnsafe(_184_v2,new BigNumber((_254_v57).length));
      let _index6 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_243_v50).length));
      let _rhs9 = (_239_v46).fm6(globalState);
      let _rhs10 = ((_247_v52).Intersect(_247_v52)).Union(_dafny.MultiSet.FromArray((_248_v53)[_module.__default.safeIndex((((_251_v54).contains((_dafny.ZERO).minus((_dafny.ZERO).minus((((_255_v58).contains(_184_v2)) ? ((_255_v58).get(_184_v2)) : ((_239_v46).f30)))))) ? ((_251_v54).get((_dafny.ZERO).minus((_dafny.ZERO).minus((((_255_v58).contains(_184_v2)) ? ((_255_v58).get(_184_v2)) : ((_239_v46).f30)))))) : (new BigNumber((_246_v51).length))), new BigNumber((_248_v53).length))]));
      let _lhs7 = _243_v50;
      let _lhs8 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_243_v50).length));
      _lhs7[_lhs8] = _rhs9;
      _247_v52 = _rhs10;
      r0 = (_239_v46).f30;
      let _256_v59;
      let _nw16 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _256_v59 = _nw16;
      _256_v59 = _256_v59;
      let _hi1 = (_239_v46).f30;
      for (let _257_i7 = (_252_v55).f30; _257_i7.isLessThan(_hi1); _257_i7 = _257_i7.plus(_dafny.ONE)) {
        (_239_v46).f31 = false;
        (globalState).f13 = ((_239_v46).f30).minus((_239_v46).f34);
        let _258_v60;
        _258_v60 = _module.D0.create_DC1(_187_v5);
        let _259_v61;
        let _nw17 = new _module.C10();
        _nw17.__ctor(_module.__default.fm14(new BigNumber(386), ((_252_v55).f32)[_module.__default.safeIndex((_239_v46).f34, new BigNumber(((_252_v55).f32).length))], (_252_v55).f30, _258_v60, globalState), new BigNumber(-339), p1, _dafny.Seq.Concat((_239_v46).f32, _235_v42), _dafny.Seq.Concat((_239_v46).f33, _237_v44), (_239_v46).f30, (_243_v50)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_243_v50).length))], _239_v46.f31, ((true) ? ((_238_v45).fm7(_251_v54, (_dafny.ZERO).minus(new BigNumber(-158)), globalState)) : (_252_v55.f29)));
        _259_v61 = _nw17;
        let _260_v62;
        let _nw18 = new _module.C2();
        _nw18.__ctor((_239_v46).f34, (_252_v55).f30, _252_v55.f29, _252_v55.f29, false);
        _260_v62 = _nw18;
        let _261_v63;
        _261_v63 = _dafny.Seq.of(_260_v62, _260_v62);
        let _262_v64;
        _262_v64 = _dafny.Seq.Concat(_dafny.Seq.of(_260_v62), _261_v63);
        _262_v64 = _262_v64;
      }
      r0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p2, new BigNumber(((((_243_v50)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_243_v50).length))]) ? (_246_v51) : (_246_v51))).length)));
      return r0;
    }
    static Main(__noArgsParameter) {
      let _263_v0;
      let _nw19 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _263_v0 = _nw19;
      let _264_v1;
      _264_v1 = new BigNumber(20);
      let _265_v2;
      _265_v2 = _dafny.Seq.of(_264_v1);
      let _266_v3;
      _266_v3 = _dafny.Map.Empty.slice().updateUnsafe(_263_v0,_265_v2);
      let _267_v4;
      _267_v4 = true;
      let _268_v5;
      _268_v5 = _dafny.Map.Empty.slice().updateUnsafe((((_266_v3).contains(_263_v0)) ? ((_266_v3).get(_263_v0)) : (_265_v2)),_267_v4);
      let _269_v6;
      _269_v6 = _dafny.Seq.UnicodeFromString("rhhbc");
      let _270_v7;
      _270_v7 = _dafny.Seq.of(_269_v6);
      let _271_v8;
      let _nw20 = Array((new BigNumber(21)).toNumber()).fill([]);
      _271_v8 = _nw20;
      let _272_v9;
      let _init2 = ((_273_v4) => function (_274_i0) {
        return _273_v4;
      })(_267_v4);
      let _nw21 = Array((_dafny.ONE).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw21.length); _i0_2++) {
        _nw21[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _272_v9 = _nw21;
      let _275_globalState;
      let _nw22 = new _module.GlobalState();
      _nw22.__ctor(new BigNumber(324), false, true, false, new BigNumber(862), (_268_v5).update(_265_v2, _267_v4), new BigNumber(38), _269_v6, (_270_v7)[_module.__default.safeIndex(_264_v1, new BigNumber((_270_v7).length))], true, new BigNumber(167), new BigNumber(-961), true, new BigNumber(457), true, _271_v8, new BigNumber(745), new _dafny.CodePoint('d'.codePointAt(0)), _272_v9, true, true, _272_v9, _263_v0, true, new _dafny.CodePoint('y'.codePointAt(0)), _265_v2, false, new BigNumber(426));
      _275_globalState = _nw22;
      _264_v1 = new BigNumber(569);
      if (_267_v4) {
        let _276_v10;
        _276_v10 = _dafny.MultiSet.fromElements(_264_v1);
        let _277_v11;
        _277_v11 = _dafny.Map.Empty.slice().updateUnsafe(_267_v4,_264_v1);
        let _278_v12;
        _278_v12 = _dafny.Seq.of(_267_v4);
        let _279_v13;
        _279_v13 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_264_v1),new BigNumber((_278_v12).length));
        _264_v1 = _module.__default.safeModuloInt((((_276_v10).contains((((_277_v11).contains(_267_v4)) ? ((_277_v11).get(_267_v4)) : (_264_v1)))) ? ((_276_v10).get((((_277_v11).contains(_267_v4)) ? ((_277_v11).get(_267_v4)) : (_264_v1)))) : (_264_v1)), (((_279_v13).contains(_264_v1)) ? ((_279_v13).get(_264_v1)) : (_264_v1)));
        let _280_v14;
        let _out0;
        _out0 = _module.__default.m0(_265_v2, _267_v4, _264_v1, !((_dafny.Map.Empty.slice().updateUnsafe(_267_v4,_263_v0)).update(true, _263_v0)).contains(_267_v4), _275_globalState);
        _280_v14 = _out0;
        (_275_globalState).f1 = _dafny.areEqual(_265_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(940)), ((_281_v14) => function (_282_i1) {
          return (_dafny.ZERO).minus(_281_v14);
        })(_280_v14)));
        let _283_v15;
        _283_v15 = _dafny.Set.fromElements(_267_v4, _267_v4);
        let _284_v16;
        _284_v16 = _dafny.Map.Empty.slice().updateUnsafe(_272_v9,(_283_v15).contains(_267_v4));
        _284_v16 = _284_v16;
        let _285_v17;
        _285_v17 = _dafny.MultiSet.fromElements(_267_v4, _267_v4);
        let _286_v18;
        _286_v18 = new _dafny.CodePoint('b'.codePointAt(0));
        let _287_v19;
        _287_v19 = _dafny.Map.Empty.slice().updateUnsafe(_267_v4,_286_v18);
        let _288_v20;
        let _nw23 = Array((new BigNumber(11)).toNumber());
        _nw23[(_dafny.ZERO).toNumber()] = _278_v12;
        _nw23[(_dafny.ONE).toNumber()] = _278_v12;
        _nw23[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_278_v12, _module.__default.fm0(_280_v14, _267_v4, _267_v4, new BigNumber((_285_v17).cardinality()), _275_globalState));
        _nw23[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_267_v4);
        _nw23[(new BigNumber(4)).toNumber()] = ((_267_v4) ? (_278_v12) : (_278_v12));
        _nw23[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_278_v12, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_287_v19).length)), new BigNumber((_278_v12).length)), _267_v4);
        _nw23[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_267_v4, _267_v4), _278_v12);
        _nw23[(new BigNumber(7)).toNumber()] = _278_v12;
        _nw23[(new BigNumber(8)).toNumber()] = _module.__default.fm0(_264_v1, _267_v4, !(_267_v4), new BigNumber((_285_v17).cardinality()), _275_globalState);
        _nw23[(new BigNumber(9)).toNumber()] = _278_v12;
        _nw23[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_278_v12, _278_v12);
        _288_v20 = _nw23;
        let _index7 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_288_v20).length));
        (_288_v20)[_index7] = _278_v12;
        let _index8 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_288_v20).length));
        (_288_v20)[_index8] = _278_v12;
      } else {
        let _index9 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_272_v9).length));
        (_272_v9)[_index9] = _267_v4;
        let _289_v21;
        _289_v21 = _dafny.MultiSet.fromElements(!(_267_v4));
        let _290_v22;
        _290_v22 = _dafny.Map.Empty.slice().updateUnsafe(_267_v4,!(_289_v21).contains(_267_v4));
        let _index10 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_272_v9).length));
        let _rhs11 = (((_290_v22).contains(!(_267_v4) || (_267_v4))) ? ((_290_v22).get(!(_267_v4) || (_267_v4))) : (_267_v4));
        let _rhs12 = !(_module.__default.fm1(((_267_v4) ? (false) : (_267_v4)), _264_v1, _275_globalState));
        let _lhs9 = _272_v9;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_272_v9).length));
        let _lhs11 = _275_globalState;
        _lhs9[_lhs10] = _rhs11;
        _lhs11.f26 = _rhs12;
        let _291_v23;
        _291_v23 = _dafny.MultiSet.fromElements(_264_v1);
        let _292_v24;
        let _init3 = function (_293_i2) {
          return false;
        };
        let _nw24 = Array((new BigNumber(28)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw24.length); _i0_3++) {
          _nw24[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _292_v24 = _nw24;
        let _294_v25;
        _294_v25 = _dafny.Set.fromElements(_292_v24);
        let _295_v26;
        _295_v26 = _dafny.Map.Empty.slice().updateUnsafe(!((_272_v9)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_272_v9).length))]),_292_v24);
        let _296_v27;
        _296_v27 = _dafny.Map.Empty.slice().updateUnsafe(_291_v23,(_294_v25).IsDisjointFrom(_dafny.Set.fromElements(_292_v24, (((_295_v26).contains(_267_v4)) ? ((_295_v26).get(_267_v4)) : (_292_v24)))));
        (_275_globalState).f1 = (((_296_v27).contains(_291_v23)) ? ((_296_v27).get(_291_v23)) : (_267_v4));
        let _297_v28;
        _297_v28 = _dafny.Seq.of((_272_v9)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_272_v9).length))]);
        let _298_v29;
        let _out1;
        _out1 = _module.__default.m0(_265_v2, (_297_v28)[_module.__default.safeIndex(new BigNumber(780), new BigNumber((_297_v28).length))], _264_v1, !(_267_v4), _275_globalState);
        _298_v29 = _out1;
        _297_v28 = _297_v28;
        let _299_v30;
        _299_v30 = _dafny.Set.fromElements(new BigNumber((_297_v28).length));
        let _300_v31;
        _300_v31 = _dafny.Set.fromElements(_297_v28, _297_v28, _297_v28);
        let _301_v33;
        _301_v33 = _dafny.MultiSet.fromElements(function () {
          let _coll17 = new _dafny.Set();
          for (const _compr_17 of (_300_v31).Elements) {
            let _302_v32 = _compr_17;
            if ((_300_v31).contains(_302_v32)) {
              _coll17.add(_302_v32);
            }
          }
          return _coll17;
        }(), _300_v31);
        (_275_globalState).f26 = !(((_299_v30).IsDisjointFrom(_299_v30)) === (((((_301_v33).contains(_300_v31)) ? ((_301_v33).get(_300_v31)) : (_298_v29))).isLessThan(_298_v29)));
      }
      let _hi2 = _264_v1;
      for (let _303_i3 = _264_v1; _303_i3.isLessThan(_hi2); _303_i3 = _303_i3.plus(_dafny.ONE)) {
        (_275_globalState).f1 = _267_v4;
        let _304_v34;
        _304_v34 = _dafny.MultiSet.fromElements(_303_i3);
        let _305_v35;
        let _nw25 = new _module.C7();
        _nw25.__ctor(_303_i3, _304_v34, (_264_v1).multipliedBy(new BigNumber(568)), !(_267_v4) || (_267_v4), _267_v4, _267_v4);
        _305_v35 = _nw25;
        (_275_globalState).f12 = ((_305_v35).f41).isLessThanOrEqualTo(_303_i3);
        (_275_globalState).f13 = _303_i3;
      }
      let _rhs13 = _264_v1;
      let _rhs14 = new BigNumber(417);
      let _lhs12 = _275_globalState;
      _lhs12.f13 = _rhs13;
      _264_v1 = _rhs14;
      let _306_v36;
      let _nw26 = new _module.C7();
      _nw26.__ctor(_264_v1, _dafny.MultiSet.FromArray(_265_v2), new BigNumber(780), _267_v4, false, _267_v4);
      _306_v36 = _nw26;
      let _307_v37;
      _307_v37 = _dafny.Map.Empty.slice().updateUnsafe(_272_v9,_306_v36);
      let _308_v38;
      _308_v38 = _dafny.Map.Empty.slice().updateUnsafe((_306_v36).f41,(_307_v37).update(_272_v9, _306_v36));
      _307_v37 = (_307_v37).Merge((((_308_v38).contains(_module.__default.fm46(new BigNumber(811), _264_v1, _267_v4, _275_globalState))) ? ((_308_v38).get(_module.__default.fm46(new BigNumber(811), _264_v1, _267_v4, _275_globalState))) : (_307_v37)));
      let _309_v39;
      _309_v39 = _dafny.Map.Empty.slice().updateUnsafe((_265_v2)[_module.__default.safeIndex(_264_v1, new BigNumber((_265_v2).length))],((_306_v36).f41).minus(_264_v1));
      let _310_v40;
      _310_v40 = new _dafny.CodePoint('m'.codePointAt(0));
      let _311_v41;
      _311_v41 = _dafny.Set.fromElements(_310_v40);
      _309_v39 = (_309_v39).update(_module.__default.safeModuloInt((_dafny.ZERO).minus(_264_v1), _264_v1), new BigNumber(((_311_v41).Union(_311_v41)).length));
      let _312_v42;
      let _nw27 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
      _312_v42 = _nw27;
      let _313_v43;
      _313_v43 = _312_v42;
      _312_v42 = (_313_v43);
      let _hi3 = (_306_v36).f41;
      for (let _314_i4 = _264_v1; _314_i4.isLessThan(_hi3); _314_i4 = _314_i4.plus(_dafny.ONE)) {
        if (_267_v4) {
          let _315_v44;
          _315_v44 = _module.D22.create_DC55(new BigNumber(263), _264_v1, _267_v4);
          let _316_v45;
          _316_v45 = _module.D3.create_DC9((_315_v44).dtor_cf94, _267_v4, _272_v9, _267_v4);
          let _317_v46;
          let _318_v47;
          let _out2;
          let _out3;
          let _outcollector0 = (_306_v36).m8((_316_v45).dtor_cf20, _275_globalState);
          _out2 = _outcollector0[0];
          _out3 = _outcollector0[1];
          _317_v46 = _out2;
          _318_v47 = _out3;
          _264_v1 = _314_i4;
          (_275_globalState).f12 = _267_v4;
          let _319_v48;
          let _320_v49;
          let _out4;
          let _out5;
          let _outcollector1 = (_306_v36).m8(_272_v9, _275_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _319_v48 = _out4;
          _320_v49 = _out5;
          (_275_globalState).f26 = _267_v4;
        } else {
          (_275_globalState).f1 = (new BigNumber(934)).isLessThanOrEqualTo(_314_i4);
          let _321_v50;
          let _nw28 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Set.Empty);
          _321_v50 = _nw28;
          let _322_v51;
          _322_v51 = _dafny.Set.fromElements(_267_v4);
          let _323_v52;
          _323_v52 = _dafny.Seq.of(_267_v4);
          let _324_v53;
          _324_v53 = _dafny.Seq.of(_323_v52);
          let _325_v54;
          _325_v54 = _module.D14.create_DC37((_306_v36).f41, _267_v4, (_306_v36).f41);
          let _326_v55;
          _326_v55 = _dafny.MultiSet.fromElements(true);
          let _327_v56;
          _327_v56 = _dafny.Map.Empty.slice().updateUnsafe(_267_v4,!(false));
          let _nw29 = Array((new BigNumber(28)).toNumber());
          _nw29[(_dafny.ZERO).toNumber()] = _322_v51;
          _nw29[(_dafny.ONE).toNumber()] = _322_v51;
          _nw29[(new BigNumber(2)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(3)).toNumber()] = (_322_v51).Difference(_dafny.Set.fromElements(_267_v4, _267_v4));
          _nw29[(new BigNumber(4)).toNumber()] = (_322_v51).Union(_322_v51);
          _nw29[(new BigNumber(5)).toNumber()] = (_322_v51).Difference(_module.__default.fm11(_324_v53, !(_267_v4), _267_v4, (_306_v36).f41, _275_globalState));
          _nw29[(new BigNumber(6)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(false);
          _nw29[(new BigNumber(8)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(9)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(10)).toNumber()] = (_322_v51).Difference(_322_v51);
          _nw29[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements(_267_v4, _267_v4, _267_v4);
          _nw29[(new BigNumber(12)).toNumber()] = (_322_v51).Difference(_322_v51);
          _nw29[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(_267_v4, _267_v4, _267_v4);
          _nw29[(new BigNumber(14)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(15)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(16)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(17)).toNumber()] = (_322_v51);
          _nw29[(new BigNumber(18)).toNumber()] = _module.__default.fm45((_325_v54).dtor_cf70, _326_v55, (((_327_v56).contains((_306_v36).fm19(_275_globalState))) ? ((_327_v56).get((_306_v36).fm19(_275_globalState))) : (true)), _314_i4, _275_globalState);
          _nw29[(new BigNumber(19)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(20)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(21)).toNumber()] = (_322_v51).Intersect(_322_v51);
          _nw29[(new BigNumber(22)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(23)).toNumber()] = _dafny.Set.fromElements((_323_v52)[_module.__default.safeIndex((_306_v36).f41, new BigNumber((_323_v52).length))], _267_v4);
          _nw29[(new BigNumber(24)).toNumber()] = (_322_v51).Intersect(_322_v51);
          _nw29[(new BigNumber(25)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(26)).toNumber()] = _322_v51;
          _nw29[(new BigNumber(27)).toNumber()] = (_module.__default.fm45(new BigNumber((_269_v6).length), _326_v55, (_306_v36).fm19(_275_globalState), new BigNumber(492), _275_globalState)).Union(_322_v51);
          _321_v50 = _nw29;
          let _328_v57;
          _328_v57 = _module.D1.create_DC2(_272_v9);
          let _pat_let_tv0 = _272_v9;
          _328_v57 = function (_pat_let0_0) {
            return function (_329_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_330_dt__update_hcf2_h0) {
                  return _module.D1.create_DC2(_330_dt__update_hcf2_h0);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_module.D1.create_DC2(_272_v9));
          let _331_v58;
          let _332_v59;
          let _out6;
          let _out7;
          let _outcollector2 = (_306_v36).m8(_272_v9, _275_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _331_v58 = _out6;
          _332_v59 = _out7;
          let _333_v60;
          let _init4 = ((_334_v40) => function (_335_i5) {
            return _334_v40;
          })(_310_v40);
          let _nw30 = Array((new BigNumber(11)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw30.length); _i0_4++) {
            _nw30[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _333_v60 = _nw30;
          let _336_v61;
          let _nw31 = new _module.C0();
          _nw31.__ctor(_333_v60);
          _336_v61 = _nw31;
          _336_v61 = _336_v61;
        }
        let _337_v62;
        let _nw32 = new _module.C2();
        _nw32.__ctor((new BigNumber((_dafny.Seq.of(_269_v6)).length)).multipliedBy(new BigNumber((_269_v6).length)), _264_v1, _267_v4, _267_v4, _267_v4);
        _337_v62 = _nw32;
        let _338_v63;
        let _init5 = ((_339_v4) => function (_340_i6) {
          return _dafny.MultiSet.fromElements(_339_v4);
        })(_267_v4);
        let _nw33 = Array((new BigNumber(23)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw33.length); _i0_5++) {
          _nw33[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _338_v63 = _nw33;
        let _341_v64;
        _341_v64 = _dafny.MultiSet.fromElements(_267_v4);
        let _index11 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_338_v63).length));
        (_338_v63)[_index11] = (_341_v64).Union(_341_v64);
        let _342_v65;
        _342_v65 = _dafny.Seq.of(_267_v4, _267_v4);
        let _index12 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_338_v63).length));
        (_338_v63)[_index12] = ((((_342_v65)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(292)), ((_343_v4) => function (_344_i7) {
          return _dafny.Map.Empty.slice().updateUnsafe(_343_v4,_343_v4);
        })(_267_v4))).length), new BigNumber((_342_v65).length))]) ? (_dafny.MultiSet.fromElements(_267_v4, _267_v4)) : (_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.update(_342_v65, _module.__default.safeIndex(new BigNumber((_269_v6).length), new BigNumber((_342_v65).length)), _267_v4), _module.__default.safeIndex((_306_v36).f41, new BigNumber((_dafny.Seq.update(_342_v65, _module.__default.safeIndex(new BigNumber((_269_v6).length), new BigNumber((_342_v65).length)), _267_v4)).length)), _267_v4))))).Intersect(_341_v64);
        (_275_globalState).f13 = (_dafny.ZERO).minus(new BigNumber((_269_v6).length));
      }
      let _345_v66;
      _345_v66 = _dafny.Seq.of(_267_v4);
      let _index13 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
      (_263_v0)[_index13] = (new BigNumber((_345_v66).length)).multipliedBy((_306_v36).f41);
      let _346_v67;
      _346_v67 = _module.D8.create_DC24(false, new BigNumber((_269_v6).length), _345_v66, true, _265_v2);
      let _347_v68;
      _347_v68 = _dafny.MultiSet.fromElements(_310_v40);
      let _348_v69;
      _348_v69 = _module.D3.create_DC8(_347_v68, _267_v4, (_306_v36).f41);
      let _349_v70;
      _349_v70 = _dafny.Set.fromElements(false, _267_v4, _267_v4, (_346_v67).dtor_cf49, (_348_v69).dtor_cf16);
      let _index14 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
      (_263_v0)[_index14] = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_306_v36).f41, new BigNumber((_349_v70).length)), _264_v1);
      (_275_globalState).f1 = ((_267_v4) ? (!(false)) : (_267_v4));
      let _350_i8;
      _350_i8 = _dafny.ZERO;
      L0: {
        while (_267_v4) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_350_i8)) {
              break L0;
            }
            _350_i8 = (_350_i8).plus(_dafny.ONE);
            let _351_v71;
            let _352_v72;
            let _353_v73;
            let _354_v74;
            let _out8;
            let _out9;
            let _out10;
            let _out11;
            let _outcollector3 = (_306_v36).m7(false, _272_v9, (_349_v70).IsSubsetOf(_dafny.Set.fromElements((_306_v36).fm19(_275_globalState))), _275_globalState);
            _out8 = _outcollector3[0];
            _out9 = _outcollector3[1];
            _out10 = _outcollector3[2];
            _out11 = _outcollector3[3];
            _351_v71 = _out8;
            _352_v72 = _out9;
            _353_v73 = _out10;
            _354_v74 = _out11;
            (_275_globalState).f1 = _module.__default.fm1((_345_v66)[_module.__default.safeIndex(new BigNumber(-315), new BigNumber((_345_v66).length))], (_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))], _275_globalState);
            _353_v73 = (_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))];
            let _355_v75;
            let _nw34 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _355_v75 = _nw34;
            let _356_v76;
            _356_v76 = _module.D20.create_DC50(_module.D20.create_DC47(_355_v75));
            let _357_v77;
            _357_v77 = _dafny.Map.Empty.slice().updateUnsafe(_267_v4,_356_v76);
            let _358_v78;
            _358_v78 = _module.D20.create_DC50((((_357_v77).contains(_267_v4)) ? ((_357_v77).get(_267_v4)) : (_356_v76)));
            let _source8 = _358_v78;
            if (_source8.is_DC48) {
              let _359___mcc_h0 = (_source8).cf82;
              let _360___mcc_h1 = (_source8).cf83;
              let _361___mcc_h2 = (_source8).cf84;
              let _362_cf84 = _361___mcc_h2;
              let _363_cf83 = _360___mcc_h1;
              let _364_cf82 = _359___mcc_h0;
              let _365_v79;
              _365_v79 = _dafny.Seq.of((_module.D1.create_DC2(_272_v9)).dtor_cf2, _272_v9, _272_v9, _272_v9);
              (_275_globalState).f21 = (_365_v79)[_module.__default.safeIndex(_364_cf82, new BigNumber((_365_v79).length))];
              _270_v7 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_270_v7, _module.__default.safeIndex(new BigNumber(507), new BigNumber((_270_v7).length)), (_306_v36).fm3(_267_v4, false, _275_globalState)), _module.__default.safeIndex(_module.__default.safeDivisionInt(_module.__default.fm39(_267_v4, (_306_v36).f42, _267_v4, _275_globalState), new BigNumber(-975)), new BigNumber((_dafny.Seq.update(_270_v7, _module.__default.safeIndex(new BigNumber(507), new BigNumber((_270_v7).length)), (_306_v36).fm3(_267_v4, false, _275_globalState))).length)), _354_v74), _module.__default.safeIndex((_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))], new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_270_v7, _module.__default.safeIndex(new BigNumber(507), new BigNumber((_270_v7).length)), (_306_v36).fm3(_267_v4, false, _275_globalState)), _module.__default.safeIndex(_module.__default.safeDivisionInt(_module.__default.fm39(_267_v4, (_306_v36).f42, _267_v4, _275_globalState), new BigNumber(-975)), new BigNumber((_dafny.Seq.update(_270_v7, _module.__default.safeIndex(new BigNumber(507), new BigNumber((_270_v7).length)), (_306_v36).fm3(_267_v4, false, _275_globalState))).length)), _354_v74)).length)), _dafny.Seq.UnicodeFromString("ydtxhh"));
              _362_cf84 = _362_cf84;
              _309_v39 = (_309_v39).update(((_306_v36).f41).plus((_306_v36).f41), (_306_v36).f41);
            } else if (_source8.is_DC49) {
              let _366___mcc_h3 = (_source8).cf85;
              let _367___mcc_h4 = (_source8).cf86;
              let _368_cf86 = _367___mcc_h4;
              let _369_cf85 = _366___mcc_h3;
              let _index15 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_272_v9).length));
              (_272_v9)[_index15] = (_306_v36).fm19(_275_globalState);
              let _370_v81;
              _370_v81 = _dafny.Map.Empty.slice().updateUnsafe(_368_cf86,_351_v71);
              let _371_v82;
              let _nw35 = new _module.C2();
              _nw35.__ctor(((((_306_v36).f42).contains(_module.__default.fm46((_dafny.ZERO).minus(new BigNumber((function () {
                let _coll19 = new _dafny.Map();
                for (const _compr_19 of (_370_v81).Keys.Elements) {
                  let _373_v80 = _compr_19;
                  if ((_370_v81).contains(_373_v80)) {
                    _coll19.push([_373_v80,_267_v4]);
                  }
                }
                return _coll19;
              }()).length)), (_306_v36).f41, _369_cf85, _275_globalState))) ? (((_306_v36).f42).get(_module.__default.fm46((_dafny.ZERO).minus(new BigNumber((function () {
                let _coll18 = new _dafny.Map();
                for (const _compr_18 of (_370_v81).Keys.Elements) {
                  let _372_v80 = _compr_18;
                  if ((_370_v81).contains(_372_v80)) {
                    _coll18.push([_372_v80,_267_v4]);
                  }
                }
                return _coll18;
              }()).length)), (_306_v36).f41, _369_cf85, _275_globalState))) : (((((_306_v36).f42).contains((_306_v36).f41)) ? (((_306_v36).f42).get((_306_v36).f41)) : ((_306_v36).f41)))), new BigNumber((_354_v74).length), _267_v4, !((_345_v66)[_module.__default.safeIndex(new BigNumber(393), new BigNumber((_345_v66).length))]), _267_v4);
              _371_v82 = _nw35;
              let _374_v83;
              _374_v83 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_265_v2).length),_371_v82);
              let _375_v84;
              let _nw36 = Array((new BigNumber(23)).toNumber());
              _nw36[(_dafny.ZERO).toNumber()] = _371_v82;
              _nw36[(_dafny.ONE).toNumber()] = _371_v82;
              _nw36[(new BigNumber(2)).toNumber()] = (((_374_v83).contains((_306_v36).f41)) ? ((_374_v83).get((_306_v36).f41)) : (_371_v82));
              _nw36[(new BigNumber(3)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(4)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(5)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(6)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(7)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(8)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(9)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(10)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(11)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(12)).toNumber()] = (((_374_v83).contains(new BigNumber(22))) ? ((_374_v83).get(new BigNumber(22))) : (_371_v82));
              _nw36[(new BigNumber(13)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(14)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(15)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(16)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(17)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(18)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(19)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(20)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(21)).toNumber()] = _371_v82;
              _nw36[(new BigNumber(22)).toNumber()] = _371_v82;
              _375_v84 = _nw36;
              let _index16 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_375_v84).length));
              (_375_v84)[_index16] = _371_v82;
              let _376_v85;
              _376_v85 = _module.D20.create_DC49(_267_v4, _310_v40);
              let _377_v86;
              _377_v86 = _module.D5.create_DC13(_369_cf85, _dafny.Map.Empty.slice().updateUnsafe(!(_369_cf85),(_371_v82).f46), _310_v40);
              let _378_v87;
              let _nw37 = Array((new BigNumber(29)).toNumber());
              _nw37[(_dafny.ZERO).toNumber()] = _351_v71;
              _nw37[(_dafny.ONE).toNumber()] = _310_v40;
              _nw37[(new BigNumber(2)).toNumber()] = _module.__default.fm41(_369_cf85, _368_cf86, _267_v4, (_306_v36).f41, _275_globalState);
              _nw37[(new BigNumber(3)).toNumber()] = _310_v40;
              _nw37[(new BigNumber(4)).toNumber()] = _351_v71;
              _nw37[(new BigNumber(5)).toNumber()] = _351_v71;
              _nw37[(new BigNumber(6)).toNumber()] = _351_v71;
              _nw37[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
              _nw37[(new BigNumber(8)).toNumber()] = _310_v40;
              _nw37[(new BigNumber(9)).toNumber()] = _351_v71;
              _nw37[(new BigNumber(10)).toNumber()] = _368_cf86;
              _nw37[(new BigNumber(11)).toNumber()] = (_376_v85).dtor_cf86;
              _nw37[(new BigNumber(12)).toNumber()] = _351_v71;
              _nw37[(new BigNumber(13)).toNumber()] = _310_v40;
              _nw37[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
              _nw37[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
              _nw37[(new BigNumber(16)).toNumber()] = (_377_v86).dtor_cf29;
              _nw37[(new BigNumber(17)).toNumber()] = _351_v71;
              _nw37[(new BigNumber(18)).toNumber()] = _351_v71;
              _nw37[(new BigNumber(19)).toNumber()] = _351_v71;
              _nw37[(new BigNumber(20)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
              _nw37[(new BigNumber(21)).toNumber()] = _351_v71;
              _nw37[(new BigNumber(22)).toNumber()] = _310_v40;
              _nw37[(new BigNumber(23)).toNumber()] = _368_cf86;
              _nw37[(new BigNumber(24)).toNumber()] = _310_v40;
              _nw37[(new BigNumber(25)).toNumber()] = _310_v40;
              _nw37[(new BigNumber(26)).toNumber()] = _368_cf86;
              _nw37[(new BigNumber(27)).toNumber()] = _368_cf86;
              _nw37[(new BigNumber(28)).toNumber()] = _module.__default.fm48(_369_cf85, _369_cf85, _264_v1, _275_globalState);
              _378_v87 = _nw37;
              let _index17 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_378_v87).length));
              (_378_v87)[_index17] = new _dafny.CodePoint('m'.codePointAt(0));
              let _index18 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_272_v9).length));
              let _index19 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_375_v84).length));
              let _index20 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_378_v87).length));
              let _rhs15 = (_371_v82).fm30(_369_cf85, _275_globalState);
              let _rhs16 = _371_v82;
              let _rhs17 = _module.__default.safeModuloInt((_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))], new BigNumber(908));
              let _rhs18 = !(_267_v4);
              let _rhs19 = _368_cf86;
              let _lhs13 = _272_v9;
              let _lhs14 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_272_v9).length));
              let _lhs15 = _375_v84;
              let _lhs16 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_375_v84).length));
              let _lhs17 = _275_globalState;
              let _lhs18 = _378_v87;
              let _lhs19 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_378_v87).length));
              _lhs13[_lhs14] = _rhs15;
              _lhs15[_lhs16] = _rhs16;
              _353_v73 = _rhs17;
              _lhs17.f1 = _rhs18;
              _lhs18[_lhs19] = _rhs19;
              (_275_globalState).f26 = (_272_v9)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_272_v9).length))];
              _306_v36 = _306_v36;
              _368_cf86 = (_378_v87)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_378_v87).length))];
            } else if (_source8.is_DC47) {
              let _379___mcc_h5 = (_source8).cf81;
              let _380_cf81 = _379___mcc_h5;
              let _381_v88;
              let _382_v89;
              let _out12;
              let _out13;
              let _outcollector4 = (_306_v36).m8(_272_v9, _275_globalState);
              _out12 = _outcollector4[0];
              _out13 = _outcollector4[1];
              _381_v88 = _out12;
              _382_v89 = _out13;
              let _383_v90;
              _383_v90 = _dafny.Map.Empty.slice().updateUnsafe(_267_v4,_267_v4);
              let _384_v91;
              _384_v91 = _dafny.Map.Empty.slice().updateUnsafe((_306_v36).f41,_383_v90);
              let _385_v92;
              _385_v92 = _dafny.Set.fromElements(_263_v0);
              let _386_v93;
              let _nw38 = new _module.C9();
              _nw38.__ctor(_382_v89, _384_v91, _310_v40, _382_v89, !(_267_v4), _dafny.Seq.update(_345_v66, _module.__default.safeIndex(new BigNumber(74), new BigNumber((_345_v66).length)), (_306_v36).fm19(_275_globalState)), _dafny.Seq.update(_dafny.Seq.of(_345_v66, _345_v66), _module.__default.safeIndex((_306_v36).f41, new BigNumber((_dafny.Seq.of(_345_v66, _345_v66)).length)), _345_v66), (_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))], _267_v4, (_385_v92).IsProperSubsetOf(_385_v92), _267_v4);
              _386_v93 = _nw38;
              _386_v93 = _386_v93;
              (_275_globalState).f1 = !(!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), ((_387_v40) => function (_388_i9) {
                return _387_v40;
              })(_310_v40)), _module.__default.safeIndex((_306_v36).f41, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), ((_389_v40) => function (_390_i9) {
                return _389_v40;
              })(_310_v40))).length)), _351_v71), _dafny.Seq.Concat(_354_v74, _354_v74))));
              let _index21 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
              (_263_v0)[_index21] = ((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(265)), ((_391_v71, _392_v36, _393_v89) => function (_394_i10) {
                return new BigNumber((_dafny.Seq.of(_module.D21.create_DC53(_module.D21.create_DC53(_module.D21.create_DC52(_391_v71))), _module.D21.create_DC53(_module.D21.create_DC51(_dafny.Map.Empty.slice().updateUnsafe((_392_v36).f41,_dafny.Set.fromElements((_392_v36).f41, _393_v89)))))).length);
              })(_351_v71, _306_v36, _382_v89))).length)).minus(new BigNumber((_269_v6).length))).multipliedBy((_306_v36).f41);
            } else {
              let _395___mcc_h6 = (_source8).cf87;
              let _396_cf87 = _395___mcc_h6;
              let _397_v94;
              let _nw39 = Array((new BigNumber(23)).toNumber()).fill([]);
              _397_v94 = _nw39;
              let _index22 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_397_v94).length));
              (_397_v94)[_index22] = _272_v9;
              let _index23 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_397_v94).length));
              (_397_v94)[_index23] = _272_v9;
              let _398_v95;
              _398_v95 = _dafny.Set.fromElements(_306_v36, _306_v36, _306_v36);
              let _399_v96;
              _399_v96 = _dafny.Set.fromElements(_398_v95);
              let _400_v97;
              _400_v97 = _dafny.Seq.of(_399_v96);
              let _401_v98;
              _401_v98 = _module.D24.create_DC57(_dafny.Set.fromElements(_398_v95, _398_v95));
              let _402_v99;
              let _nw40 = Array((new BigNumber(15)).toNumber());
              _nw40[(_dafny.ZERO).toNumber()] = _399_v96;
              _nw40[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_398_v95, _398_v95);
              _nw40[(new BigNumber(2)).toNumber()] = _399_v96;
              _nw40[(new BigNumber(3)).toNumber()] = (_399_v96).Union(_dafny.Set.fromElements(_398_v95));
              _nw40[(new BigNumber(4)).toNumber()] = _399_v96;
              _nw40[(new BigNumber(5)).toNumber()] = _399_v96;
              _nw40[(new BigNumber(6)).toNumber()] = (_399_v96).Difference((_400_v97)[_module.__default.safeIndex(new BigNumber((_354_v74).length), new BigNumber((_400_v97).length))]);
              _nw40[(new BigNumber(7)).toNumber()] = _399_v96;
              _nw40[(new BigNumber(8)).toNumber()] = _399_v96;
              _nw40[(new BigNumber(9)).toNumber()] = _399_v96;
              _nw40[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_398_v95, _398_v95);
              _nw40[(new BigNumber(11)).toNumber()] = _399_v96;
              _nw40[(new BigNumber(12)).toNumber()] = _399_v96;
              _nw40[(new BigNumber(13)).toNumber()] = (_401_v98).dtor_cf96;
              _nw40[(new BigNumber(14)).toNumber()] = _399_v96;
              _402_v99 = _nw40;
              let _index24 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_402_v99).length));
              (_402_v99)[_index24] = _399_v96;
              let _index25 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_402_v99).length));
              (_402_v99)[_index25] = _399_v96;
              let _403_v100;
              _403_v100 = _module.D3.create_DC9(_267_v4, _267_v4, (_397_v94)[_module.__default.safeIndex(new BigNumber(486), new BigNumber((_397_v94).length))], _267_v4);
              let _pat_let_tv1 = _267_v4;
              let _pat_let_tv2 = _267_v4;
              let _404_v101;
              _404_v101 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let2_0) {
                return function (_405_dt__update__tmp_h1) {
                  return function (_pat_let3_0) {
                    return function (_406_dt__update_hcf21_h0) {
                      return function (_pat_let4_0) {
                        return function (_407_dt__update_hcf18_h0) {
                          return _module.D3.create_DC9(_407_dt__update_hcf18_h0, (_405_dt__update__tmp_h1).dtor_cf19, (_405_dt__update__tmp_h1).dtor_cf20, _406_dt__update_hcf21_h0);
                        }(_pat_let4_0);
                      }(_pat_let_tv2);
                    }(_pat_let3_0);
                  }(_pat_let_tv1);
                }(_pat_let2_0);
              }(_403_v100),(_306_v36).f41);
              let _408_v102;
              _408_v102 = _dafny.Map.Empty.slice().updateUnsafe((_306_v36).f41,_310_v40);
              let _409_v103;
              _409_v103 = _module.D2.create_DC6((_306_v36).f41, _267_v4, _module.__default.fm46((_306_v36).f41, new BigNumber((_404_v101).length), true, _275_globalState), _267_v4, new BigNumber((_408_v102).length));
              _264_v1 = ((_409_v103).dtor_cf13).minus((new BigNumber((_354_v74).length)).multipliedBy(_264_v1));
              let _410_v104;
              let _411_v105;
              let _412_v106;
              let _413_v107;
              let _out14;
              let _out15;
              let _out16;
              let _out17;
              let _outcollector5 = (_306_v36).m7(_dafny.Seq.contains(_345_v66, _267_v4), (_397_v94)[_module.__default.safeIndex(new BigNumber(486), new BigNumber((_397_v94).length))], _module.__default.fm1(true, new BigNumber((_module.__default.fm33(_267_v4, _275_globalState)).cardinality()), _275_globalState), _275_globalState);
              _out14 = _outcollector5[0];
              _out15 = _outcollector5[1];
              _out16 = _outcollector5[2];
              _out17 = _outcollector5[3];
              _410_v104 = _out14;
              _411_v105 = _out15;
              _412_v106 = _out16;
              _413_v107 = _out17;
            }
          }
        }
      }
      let _414_v108;
      _414_v108 = _dafny.Seq.of(_345_v66);
      let _415_v109;
      let _nw41 = new _module.C4();
      _nw41.__ctor((_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))], _267_v4, _345_v66, _414_v108, ((_306_v36).f41).plus(new BigNumber((_349_v70).length)), true, _267_v4, _267_v4);
      _415_v109 = _nw41;
      _270_v7 = _dafny.Seq.Concat(_dafny.Seq.update(_270_v7, _module.__default.safeIndex(_264_v1, new BigNumber((_270_v7).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), ((_416_v40) => function (_417_i11) {
        return _416_v40;
      })(_310_v40))), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ssisbjr")));
      let _418_v110;
      _418_v110 = _dafny.MultiSet.fromElements(!(true));
      if ((_418_v110).IsProperSubsetOf(_418_v110)) {
        let _419_v111;
        let _nw42 = Array((new BigNumber(2)).toNumber());
        _419_v111 = _nw42;
        let _420_v112;
        let _nw43 = new _module.C1();
        _nw43.__ctor(false, _267_v4, _267_v4);
        _420_v112 = _nw43;
        let _index26 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_419_v111).length));
        (_419_v111)[_index26] = _420_v112;
        let _421_v114;
        let _init6 = ((_422_v1, _423_v36) => function (_424_i12) {
          return (_dafny.Set.fromElements(new BigNumber(646), (_dafny.ZERO).minus(_422_v1))).Union(function () {
            let _coll20 = new _dafny.Set();
            for (const _compr_20 of _dafny.IntegerRange(new BigNumber(718), new BigNumber(593))) {
              let _425_v113 = _compr_20;
              if (((new BigNumber(718)).isLessThanOrEqualTo(_425_v113)) && ((_425_v113).isLessThan(new BigNumber(593)))) {
                _coll20.add((_425_v113).plus((_423_v36).f41));
              }
            }
            return _coll20;
          }());
        })(_264_v1, _306_v36);
        let _nw44 = Array((new BigNumber(20)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw44.length); _i0_6++) {
          _nw44[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _421_v114 = _nw44;
        let _index27 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_419_v111).length));
        let _rhs20 = _420_v112;
        let _rhs21 = _421_v114;
        let _rhs22 = _418_v110;
        let _lhs20 = _419_v111;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_419_v111).length));
        _lhs20[_lhs21] = _rhs20;
        _421_v114 = _rhs21;
        _418_v110 = _rhs22;
        let _426_v115;
        let _427_v116;
        let _428_v117;
        let _429_v118;
        let _out18;
        let _out19;
        let _out20;
        let _out21;
        let _outcollector6 = (_306_v36).m7(false, _272_v9, _420_v112.f29, _275_globalState);
        _out18 = _outcollector6[0];
        _out19 = _outcollector6[1];
        _out20 = _outcollector6[2];
        _out21 = _outcollector6[3];
        _426_v115 = _out18;
        _427_v116 = _out19;
        _428_v117 = _out20;
        _429_v118 = _out21;
        let _430_v119;
        _430_v119 = _dafny.Map.Empty.slice().updateUnsafe(_267_v4,(_306_v36).f41);
        let _431_v120;
        let _nw45 = new _module.C4();
        _nw45.__ctor(new BigNumber((_429_v118).length), _420_v112.f28, _345_v66, _414_v108, (((_430_v119).contains(_267_v4)) ? ((_430_v119).get(_267_v4)) : ((_306_v36).f41)), _267_v4, (_415_v109).fm6(_275_globalState), _420_v112.f28);
        _431_v120 = _nw45;
        let _432_v121;
        _432_v121 = _dafny.MultiSet.fromElements(_431_v120);
        let _433_v122;
        _433_v122 = _module.D10.create_DC30(_310_v40, _432_v121);
        let _source9 = ((_267_v4) ? (_433_v122) : (_433_v122));
        if (_source9.is_DC29) {
          let _434___mcc_h7 = (_source9).cf56;
          let _435___mcc_h8 = (_source9).cf57;
          let _436_cf57 = _435___mcc_h8;
          let _437_cf56 = _434___mcc_h7;
          let _438_v123;
          let _439_v124;
          let _440_v125;
          let _441_v126;
          let _out22;
          let _out23;
          let _out24;
          let _out25;
          let _outcollector7 = (_306_v36).m7(_267_v4, _272_v9, _267_v4, _275_globalState);
          _out22 = _outcollector7[0];
          _out23 = _outcollector7[1];
          _out24 = _outcollector7[2];
          _out25 = _outcollector7[3];
          _438_v123 = _out22;
          _439_v124 = _out23;
          _440_v125 = _out24;
          _441_v126 = _out25;
          let _index28 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
          (_263_v0)[_index28] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_431_v120).f30)).length));
          let _442_v127;
          let _nw46 = new _module.C7();
          _nw46.__ctor((_265_v2)[_module.__default.safeIndex(_440_v125, new BigNumber((_265_v2).length))], ((_306_v36).f42).Union((_306_v36).f42), (new BigNumber(-421)).minus(new BigNumber((_269_v6).length)), _431_v120.f35, _436_cf57, _431_v120.f31);
          _442_v127 = _nw46;
          let _rhs23 = (_306_v36).f41;
          let _rhs24 = (_431_v120).fm3(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_431_v120.f28), _345_v66), _420_v112.f28, _275_globalState);
          let _rhs25 = !(_420_v112.f28) || (!(_420_v112.f29) || (_267_v4));
          let _rhs26 = ((_306_v36).f42).IsProperSubsetOf((_442_v127).f42);
          let _lhs22 = _275_globalState;
          let _lhs23 = _431_v120;
          _264_v1 = _rhs23;
          _441_v126 = _rhs24;
          _lhs22.f12 = _rhs25;
          _lhs23.f35 = _rhs26;
        } else if (_source9.is_DC30) {
          let _443___mcc_h9 = (_source9).cf58;
          let _444___mcc_h10 = (_source9).cf59;
          let _445_cf59 = _444___mcc_h10;
          let _446_cf58 = _443___mcc_h9;
          let _rhs27 = _429_v118;
          let _rhs28 = _272_v9;
          let _lhs24 = _275_globalState;
          _429_v118 = _rhs27;
          _lhs24.f21 = _rhs28;
          let _447_v128;
          _447_v128 = _dafny.Set.fromElements((_306_v36).f41);
          let _rhs29 = _431_v120.f31;
          let _rhs30 = (_447_v128).IsSubsetOf(_447_v128);
          let _lhs25 = _275_globalState;
          let _lhs26 = _275_globalState;
          _lhs25.f26 = _rhs29;
          _lhs26.f26 = _rhs30;
          let _448_v129;
          _448_v129 = _dafny.Seq.of(_265_v2);
          _430_v119 = (_430_v119).update(_431_v120.f35, new BigNumber((_448_v129).length));
          let _449_v130;
          _449_v130 = _dafny.Seq.of(_module.__default.fm66(_275_globalState));
          let _450_v131;
          let _nw47 = new _module.C5();
          _nw47.__ctor(_431_v120.f31, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(535)), ((_451_v40) => function (_452_i13) {
            return _451_v40;
          })(_310_v40))).length), (_dafny.ZERO).minus((_306_v36).f41), _431_v120.f35, (_431_v120).f32, _dafny.Seq.of((_431_v120).f32, (_431_v120).f32, (_431_v120).f32, _dafny.Seq.of(_431_v120.f29, _431_v120.f29, false, _267_v4, _431_v120.f31), (_431_v120).f32), new BigNumber((_449_v130).length), true, _420_v112.f29, _431_v120.f35);
          _450_v131 = _nw47;
          let _453_v132;
          _453_v132 = _dafny.MultiSet.fromElements(_450_v131);
          _453_v132 = (_dafny.MultiSet.fromElements(_450_v131)).Difference(_453_v132);
        } else if (_source9.is_DC31) {
          let _454___mcc_h11 = (_source9).cf60;
          let _455___mcc_h12 = (_source9).cf61;
          let _456___mcc_h13 = (_source9).cf62;
          let _457_cf62 = _456___mcc_h13;
          let _458_cf61 = _455___mcc_h12;
          let _459_cf60 = _454___mcc_h11;
          let _460_v133;
          let _nw48 = new _module.C2();
          _nw48.__ctor(_264_v1, new BigNumber(((_431_v120).f32).length), _431_v120.f28, _431_v120.f29, _431_v120.f29);
          _460_v133 = _nw48;
          let _461_v134;
          _461_v134 = _module.D9.create_DC26(_460_v133);
          let _462_v135;
          let _nw49 = Array((new BigNumber(22)).toNumber());
          _462_v135 = _nw49;
          let _463_v136;
          let _nw50 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
          _463_v136 = _nw50;
          let _464_v137;
          let _nw51 = new _module.C3();
          _nw51.__ctor(new BigNumber(-956), _463_v136, _310_v40, new BigNumber(-341), _431_v120.f31, _345_v66, _dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), ((_465_v66) => function (_466_i14) {
            return _465_v66;
          })(_345_v66)), new BigNumber(775), _431_v120.f28, _459_cf60, _459_cf60);
          _464_v137 = _nw51;
          let _index29 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_462_v135).length));
          (_462_v135)[_index29] = _464_v137;
          let _467_v138;
          let _nw52 = Array((_dafny.ONE).toNumber());
          _nw52[(_dafny.ZERO).toNumber()] = _429_v118;
          _467_v138 = _nw52;
          let _index30 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_467_v138).length));
          (_467_v138)[_index30] = _429_v118;
          let _index31 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_462_v135).length));
          let _index32 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_467_v138).length));
          let _rhs31 = (_265_v2)[_module.__default.safeIndex((_431_v120).f34, new BigNumber((_265_v2).length))];
          let _rhs32 = _module.D9.create_DC26(_460_v133);
          let _rhs33 = _464_v137;
          let _rhs34 = _429_v118;
          let _rhs35 = (_431_v120).f34;
          let _lhs27 = _275_globalState;
          let _lhs28 = _462_v135;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_462_v135).length));
          let _lhs30 = _467_v138;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_467_v138).length));
          _lhs27.f13 = _rhs31;
          _461_v134 = _rhs32;
          _lhs28[_lhs29] = _rhs33;
          _lhs30[_lhs31] = _rhs34;
          _264_v1 = _rhs35;
          let _468_v139;
          let _469_v140;
          let _470_v141;
          let _out26;
          let _out27;
          let _out28;
          let _outcollector8 = (_431_v120).m1(_275_globalState);
          _out26 = _outcollector8[0];
          _out27 = _outcollector8[1];
          _out28 = _outcollector8[2];
          _468_v139 = _out26;
          _469_v140 = _out27;
          _470_v141 = _out28;
          _469_v140 = _264_v1;
          let _471_v142;
          _471_v142 = _module.D8.create_DC24(_420_v112.f28, _469_v140, _dafny.Seq.of(_459_cf60), _267_v4, _265_v2);
          let _472_v143;
          _472_v143 = _module.D8.create_DC25(_471_v142);
          let _473_v144;
          _473_v144 = _module.D8.create_DC25(_472_v143);
          let _474_v145;
          _474_v145 = _dafny.Map.Empty.slice().updateUnsafe(true,_426_v115);
          let _rhs36 = _473_v144;
          let _rhs37 = (_474_v145).contains(_431_v120.f28);
          let _lhs32 = _431_v120;
          _473_v144 = _rhs36;
          _lhs32.f29 = _rhs37;
        } else if (_source9.is_DC28) {
          let _475___mcc_h14 = (_source9).cf55;
          let _476_cf55 = _475___mcc_h14;
          let _index33 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_272_v9).length));
          (_272_v9)[_index33] = (new BigNumber((_418_v110).cardinality())).isLessThanOrEqualTo(_428_v117);
          let _index34 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_272_v9).length));
          (_272_v9)[_index34] = !_dafny.areEqual((_431_v120).f32, (_431_v120).f32);
          let _477_v146;
          let _nw53 = Array((new BigNumber(4)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _477_v146 = _nw53;
          let _478_v147;
          let _nw54 = new _module.C0();
          _nw54.__ctor(_477_v146);
          _478_v147 = _nw54;
          let _479_v148;
          _479_v148 = _dafny.Map.Empty.slice().updateUnsafe(((_306_v36).f41).plus((_306_v36).f41),(_415_v109).fm6(_275_globalState));
          let _480_v149;
          _480_v149 = _dafny.Seq.of((_306_v36).f42);
          let _481_v150;
          _481_v150 = _module.D5.create_DC13(((_306_v36).f42).IsProperSubsetOf((_480_v149)[_module.__default.safeIndex((_431_v120).f34, new BigNumber((_480_v149).length))]), _430_v119, new _dafny.CodePoint('e'.codePointAt(0)));
          let _482_v151;
          _482_v151 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_431_v120).f30, (_431_v120).f30),_478_v147);
          let _rhs38 = (((_482_v151).contains(((_431_v120.f31) ? ((_306_v36).f41) : ((_431_v120).f30)))) ? ((_482_v151).get(((_431_v120.f31) ? ((_306_v36).f41) : ((_431_v120).f30)))) : (_478_v147));
          let _rhs39 = (_479_v148).update(_428_v117, (_420_v112.f28) && (_431_v120.f29));
          let _rhs40 = !((_dafny.Set.fromElements(_431_v120.f35, (_476_cf55).f43, _431_v120.f29, _420_v112.f28)).IsProperSubsetOf((_349_v70).Difference(_module.__default.fm45((_306_v36).f41, _418_v110, _420_v112.f29, _264_v1, _275_globalState))));
          let _rhs41 = ((((_306_v36).fm20((_431_v120).f30, (_272_v9)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_272_v9).length))], _275_globalState)).isEqualTo((_306_v36).f41)) ? (((_306_v36).f41).minus(new BigNumber((_429_v118).length))) : (_264_v1));
          let _rhs42 = _481_v150;
          let _lhs33 = _275_globalState;
          let _lhs34 = _275_globalState;
          _478_v147 = _rhs38;
          _479_v148 = _rhs39;
          _lhs33.f26 = _rhs40;
          _lhs34.f13 = _rhs41;
          _481_v150 = _rhs42;
          let _483_v152;
          let _nw55 = Array((new BigNumber(11)).toNumber()).fill(false);
          _483_v152 = _nw55;
          let _484_v153;
          _484_v153 = _dafny.Map.Empty.slice().updateUnsafe(_483_v152,(_272_v9)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_272_v9).length))]);
          let _index35 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_272_v9).length));
          (_272_v9)[_index35] = (_module.__default.safeModuloInt(new BigNumber(433), new BigNumber((_484_v153).length))).isLessThanOrEqualTo((_dafny.ZERO).minus(_module.__default.safeModuloInt((_306_v36).f41, _264_v1)));
          let _index36 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_271_v8).length));
          (_271_v8)[_index36] = _263_v0;
          let _485_v154;
          _485_v154 = _dafny.Seq.of(_263_v0, _263_v0, _263_v0);
          let _index37 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_271_v8).length));
          (_271_v8)[_index37] = (_485_v154)[_module.__default.safeIndex((_306_v36).f41, new BigNumber((_485_v154).length))];
        } else {
          let _486___mcc_h15 = (_source9).cf63;
          let _487_cf63 = _486___mcc_h15;
          _418_v110 = _dafny.MultiSet.fromElements((_dafny.Set.fromElements(_420_v112.f29)).IsSubsetOf(_349_v70), true);
          (_275_globalState).f13 = new BigNumber(481);
          let _488_v155;
          let _out29;
          _out29 = _module.__default.m0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(555)), ((_489_v120) => function (_490_i15) {
            return (_489_v120).f34;
          })(_431_v120)), !(_420_v112.f28), new BigNumber((_dafny.MultiSet.fromElements(_420_v112.f29, _431_v120.f35)).cardinality()), _431_v120.f29, _275_globalState);
          _488_v155 = _out29;
          let _491_v156;
          _491_v156 = _dafny.Set.fromElements((_431_v120).f30, _428_v117, (_431_v120).f30, _264_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_420_v112.f28,_415_v109)).length));
          let _492_v157;
          let _nw56 = new _module.C6();
          _nw56.__ctor(_426_v115, new BigNumber((_dafny.Seq.Concat(_269_v6, _429_v118)).length), (_dafny.Set.fromElements((_306_v36).f41, new BigNumber(313))).IsSubsetOf(_491_v156), _345_v66, (_431_v120).f33, _428_v117, _431_v120.f35, ((_431_v120).f32)[_module.__default.safeIndex(_428_v117, new BigNumber(((_431_v120).f32).length))], _420_v112.f28);
          _492_v157 = _nw56;
          let _index38 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
          let _rhs43 = _488_v155;
          let _rhs44 = false;
          let _rhs45 = _492_v157;
          let _rhs46 = _492_v157.f29;
          let _lhs35 = _263_v0;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
          let _lhs37 = _275_globalState;
          let _lhs38 = _492_v157;
          _lhs35[_lhs36] = _rhs43;
          _lhs37.f12 = _rhs44;
          _492_v157 = _rhs45;
          _lhs38.f31 = _rhs46;
        }
        if (false) {
          let _493_v158;
          let _nw57 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _493_v158 = _nw57;
          let _494_v159;
          let _nw58 = new _module.C0();
          _nw58.__ctor(_493_v158);
          _494_v159 = _nw58;
          let _495_v160;
          _495_v160 = _dafny.Seq.of(_494_v159, _494_v159);
          let _rhs47 = (_dafny.ZERO).minus(((_306_v36).fm20(new BigNumber((_495_v160).length), _431_v120.f28, _275_globalState)).minus((_306_v36).f41));
          let _rhs48 = (_dafny.ZERO).minus((_431_v120).f34);
          let _lhs39 = _275_globalState;
          _lhs39.f13 = _rhs47;
          _264_v1 = _rhs48;
          let _496_v161;
          _496_v161 = _module.D1.create_DC3(_431_v120.f31, _428_v117, _494_v159, _module.__default.fm32(_420_v112.f29, _275_globalState));
          (_275_globalState).f13 = (_496_v161).dtor_cf4;
          let _497_v162;
          _497_v162 = _dafny.Set.fromElements(_265_v2, _dafny.Seq.of(_264_v1, (_431_v120).f34));
          let _498_v163;
          _498_v163 = _module.D16.create_DC39((_497_v162).Difference(_497_v162));
          _498_v163 = _module.D16.create_DC39(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(405)), ((_499_v120) => function (_500_i16) {
  return (_499_v120).f34;
})(_431_v120))));
          (_431_v120).f31 = (_431_v120).fm6(_275_globalState);
          let _501_v164;
          let _nw59 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _501_v164 = _nw59;
          let _502_v165;
          let _nw60 = new _module.C3();
          _nw60.__ctor(new BigNumber((_265_v2).length), _501_v164, new _dafny.CodePoint('d'.codePointAt(0)), (_dafny.ZERO).minus((_306_v36).f41), _431_v120.f28, (_431_v120).f32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(18)), ((_503_v120) => function (_504_i17) {
            return _dafny.Seq.of(_503_v120.f28, _503_v120.f31);
          })(_431_v120)), (new BigNumber(((_431_v120).f32).length)).plus((_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))]), _420_v112.f28, !(_431_v120.f29), !(_264_v1).isEqualTo((_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))]));
          _502_v165 = _nw60;
        } else {
          let _505_v166;
          let _nw61 = new _module.C10();
          _nw61.__ctor((((_306_v36).fm19(_275_globalState)) ? (_310_v40) : ((_429_v118)[_module.__default.safeIndex((_431_v120).f34, new BigNumber((_429_v118).length))])), (_431_v120).f30, _420_v112.f28, (_431_v120).f32, (_431_v120).f33, (_431_v120).f34, _420_v112.f29, _420_v112.f29, ((_306_v36).f42).equals(_module.__default.fm25(_275_globalState)));
          _505_v166 = _nw61;
          let _rhs49 = (((_306_v36).f42).Intersect((_306_v36).f42)).IsProperSubsetOf((_306_v36).f42);
          let _rhs50 = ((_431_v120.f29) ? (_431_v120.f31) : (_module.__default.fm1(_267_v4, new BigNumber(419), _275_globalState)));
          let _rhs51 = ((_306_v36).f41).isLessThanOrEqualTo((_306_v36).f41);
          let _rhs52 = _505_v166;
          let _lhs40 = _275_globalState;
          let _lhs41 = _275_globalState;
          let _lhs42 = _420_v112;
          _lhs40.f1 = _rhs49;
          _lhs41.f12 = _rhs50;
          _lhs42.f28 = _rhs51;
          _505_v166 = _rhs52;
          let _506_v167;
          let _nw62 = new _module.C8();
          _nw62.__ctor(_431_v120.f29, _module.__default.fm48(_431_v120.f28, _431_v120.f29, (_431_v120).f30, _275_globalState), ((_306_v36).f41).plus((_306_v36).f41), _420_v112.f28, _345_v66, _dafny.Seq.Concat(_414_v108, (_431_v120).f33), (_306_v36).f41, _420_v112.f28, (_311_v41).IsProperSubsetOf(_311_v41), false);
          _506_v167 = _nw62;
          let _507_v168;
          let _nw63 = new _module.C8();
          _nw63.__ctor(_420_v112.f28, _426_v115, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_431_v120.f29,_431_v120.f31)).length), _431_v120.f29, (_431_v120).f32, _dafny.Seq.Concat(_414_v108, (_431_v120).f33), (_306_v36).f41, _431_v120.f28, ((_306_v36).f41).isLessThanOrEqualTo((_431_v120).f30), _431_v120.f31);
          _507_v168 = _nw63;
          _507_v168 = _507_v168;
          let _508_v169;
          _508_v169 = _dafny.Map.Empty.slice().updateUnsafe(_431_v120.f35,(_507_v168).fm7((_306_v36).f42, (_431_v120).f34, _275_globalState));
          let _509_v170;
          _509_v170 = _dafny.Map.Empty.slice().updateUnsafe(_431_v120.f28,(_508_v169).Merge(_508_v169));
          let _510_v171;
          let _nw64 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
          _510_v171 = _nw64;
          let _511_v172;
          let _init7 = ((_512_v118) => function (_513_i18) {
            return _512_v118;
          })(_429_v118);
          let _nw65 = Array((new BigNumber(4)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw65.length); _i0_7++) {
            _nw65[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _511_v172 = _nw65;
          let _index39 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_511_v172).length));
          (_511_v172)[_index39] = (_431_v120).fm3(_431_v120.f35, !(_431_v120.f31), _275_globalState);
          let _514_v173;
          _514_v173 = _dafny.Map.Empty.slice().updateUnsafe(_309_v39,true);
          let _515_v175;
          let _nw66 = Array((new BigNumber(15)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = _310_v40;
          _nw66[(_dafny.ONE).toNumber()] = _426_v115;
          _nw66[(new BigNumber(2)).toNumber()] = _426_v115;
          _nw66[(new BigNumber(3)).toNumber()] = _310_v40;
          _nw66[(new BigNumber(4)).toNumber()] = _507_v168.f36;
          _nw66[(new BigNumber(5)).toNumber()] = _310_v40;
          _nw66[(new BigNumber(6)).toNumber()] = _426_v115;
          _nw66[(new BigNumber(7)).toNumber()] = _426_v115;
          _nw66[(new BigNumber(8)).toNumber()] = _507_v168.f36;
          _nw66[(new BigNumber(9)).toNumber()] = _507_v168.f36;
          _nw66[(new BigNumber(10)).toNumber()] = _507_v168.f36;
          _nw66[(new BigNumber(11)).toNumber()] = _507_v168.f36;
          _nw66[(new BigNumber(12)).toNumber()] = _426_v115;
          _nw66[(new BigNumber(13)).toNumber()] = _310_v40;
          _nw66[(new BigNumber(14)).toNumber()] = _310_v40;
          _515_v175 = _nw66;
          let _516_v176;
          let _nw67 = new _module.C0();
          _nw67.__ctor(_515_v175);
          _516_v176 = _nw67;
          let _517_v177;
          _517_v177 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm39(_431_v120.f35, (_306_v36).f42, _431_v120.f28, _275_globalState),_516_v176);
          let _518_v178;
          _518_v178 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_517_v177).length),_431_v120.f29);
          let _index40 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_511_v172).length));
          let _rhs53 = _module.__default.fm67(_275_globalState);
          let _rhs54 = !(((_431_v120.f31) ? ((((_514_v173).contains(function () {
            let _coll22 = new _dafny.Map();
            for (const _compr_22 of _dafny.IntegerRange(new BigNumber(783), new BigNumber(265))) {
              let _520_v174 = _compr_22;
              if (((new BigNumber(783)).isLessThanOrEqualTo(_520_v174)) && ((_520_v174).isLessThan(new BigNumber(265)))) {
                _coll22.push([_module.__default.safeDivisionInt(_520_v174, (_507_v168).f34),(_507_v168).f30]);
              }
            }
            return _coll22;
          }())) ? ((_514_v173).get(function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of _dafny.IntegerRange(new BigNumber(783), new BigNumber(265))) {
              let _519_v174 = _compr_21;
              if (((new BigNumber(783)).isLessThanOrEqualTo(_519_v174)) && ((_519_v174).isLessThan(new BigNumber(265)))) {
                _coll21.push([_module.__default.safeDivisionInt(_519_v174, (_507_v168).f34),(_507_v168).f30]);
              }
            }
            return _coll21;
          }())) : (_420_v112.f29))) : ((((_518_v178).contains((_507_v168).f30)) ? ((_518_v178).get((_507_v168).f30)) : (_420_v112.f29)))));
          let _rhs55 = _510_v171;
          let _rhs56 = _429_v118;
          let _rhs57 = _dafny.Seq.Concat(((!(_420_v112.f28)) ? (_429_v118) : (_429_v118)), ((_507_v168.f29) ? (_dafny.Seq.UnicodeFromString("lqm")) : (_dafny.Seq.UnicodeFromString("xb"))));
          let _lhs43 = _431_v120;
          let _lhs44 = _511_v172;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_511_v172).length));
          _509_v170 = _rhs53;
          _lhs43.f29 = _rhs54;
          _510_v171 = _rhs55;
          _429_v118 = _rhs56;
          _lhs44[_lhs45] = _rhs57;
          let _index41 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_272_v9).length));
          (_272_v9)[_index41] = false;
          let _index42 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_272_v9).length));
          (_272_v9)[_index42] = ((_431_v120).f32)[_module.__default.safeIndex((_306_v36).f41, new BigNumber(((_431_v120).f32).length))];
        }
        let _521_v179;
        let _nw68 = new _module.C1();
        _nw68.__ctor(_431_v120.f35, _431_v120.f29, _420_v112.f29);
        _521_v179 = _nw68;
        let _522_v180;
        _522_v180 = _module.D10.create_DC28(_521_v179);
        let _pat_let_tv3 = _521_v179;
        let _523_v181;
        let _nw69 = Array((new BigNumber(26)).toNumber());
        _nw69[(_dafny.ZERO).toNumber()] = _522_v180;
        _nw69[(_dafny.ONE).toNumber()] = _module.D10.create_DC28(_521_v179);
        _nw69[(new BigNumber(2)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(3)).toNumber()] = _module.D10.create_DC28(_521_v179);
        _nw69[(new BigNumber(4)).toNumber()] = _module.D10.create_DC28(_521_v179);
        _nw69[(new BigNumber(5)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(6)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(7)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(8)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(9)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(10)).toNumber()] = function (_pat_let5_0) {
          return function (_524_dt__update__tmp_h2) {
            return function (_pat_let6_0) {
              return function (_525_dt__update_hcf55_h0) {
                return _module.D10.create_DC28(_525_dt__update_hcf55_h0);
              }(_pat_let6_0);
            }(_pat_let_tv3);
          }(_pat_let5_0);
        }(_522_v180);
        _nw69[(new BigNumber(11)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(12)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(13)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(14)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(15)).toNumber()] = _module.D10.create_DC28((_522_v180).dtor_cf55);
        _nw69[(new BigNumber(16)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(17)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(18)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(19)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(20)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(21)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(22)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(23)).toNumber()] = _522_v180;
        _nw69[(new BigNumber(24)).toNumber()] = _module.D10.create_DC28(_521_v179);
        _nw69[(new BigNumber(25)).toNumber()] = _522_v180;
        _523_v181 = _nw69;
        let _526_v182;
        _526_v182 = _module.D20.create_DC48((_431_v120).f30, _523_v181, _431_v120.f29);
        _526_v182 = _526_v182;
      } else {
        let _527_v183;
        _527_v183 = _module.D3.create_DC9(false, !(_267_v4), _272_v9, _267_v4);
        let _528_v184;
        let _529_v185;
        let _out30;
        let _out31;
        let _outcollector9 = (_306_v36).m8((_527_v183).dtor_cf20, _275_globalState);
        _out30 = _outcollector9[0];
        _out31 = _outcollector9[1];
        _528_v184 = _out30;
        _529_v185 = _out31;
        let _530_v186;
        let _nw70 = Array((new BigNumber(27)).toNumber());
        _530_v186 = _nw70;
        let _531_v187;
        let _nw71 = new _module.C5();
        _nw71.__ctor(!(true), _264_v1, _529_v185, _267_v4, _345_v66, _414_v108, (_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))], _267_v4, _267_v4, _267_v4);
        _531_v187 = _nw71;
        let _index43 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_530_v186).length));
        (_530_v186)[_index43] = _531_v187;
        let _index44 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_530_v186).length));
        (_530_v186)[_index44] = _531_v187;
        let _532_v188;
        _532_v188 = _dafny.Seq.of(_415_v109, _415_v109, _415_v109, _415_v109, _415_v109);
        let _533_v189;
        _533_v189 = _module.D17.create_DC42(_dafny.Seq.Concat(_532_v188, _532_v188));
        let _534_v190;
        _534_v190 = _dafny.Map.Empty.slice().updateUnsafe(_531_v187.f31,_532_v188);
        _533_v189 = _module.D17.create_DC42((((_534_v190).contains((_531_v187).fm6(_275_globalState))) ? ((_534_v190).get((_531_v187).fm6(_275_globalState))) : (_532_v188)));
        let _535_v191;
        let _536_v192;
        let _out32;
        let _out33;
        let _outcollector10 = (_306_v36).m8(_272_v9, _275_globalState);
        _out32 = _outcollector10[0];
        _out33 = _outcollector10[1];
        _535_v191 = _out32;
        _536_v192 = _out33;
        let _537_v193;
        let _init8 = ((_538_v40) => function (_539_i19) {
          return _538_v40;
        })(_310_v40);
        let _nw72 = Array((new BigNumber(8)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw72.length); _i0_8++) {
          _nw72[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _537_v193 = _nw72;
        let _540_v194;
        let _nw73 = new _module.C0();
        _nw73.__ctor(_537_v193);
        _540_v194 = _nw73;
        let _541_v195;
        _541_v195 = _module.D1.create_DC3(_531_v187.f28, (_dafny.ZERO).minus(_529_v185), _540_v194, (_265_v2)[_module.__default.safeIndex((_306_v36).f41, new BigNumber((_265_v2).length))]);
        let _542_v196;
        let _nw74 = new _module.C1();
        _nw74.__ctor(_531_v187.f28, _531_v187.f28, !((_541_v195).dtor_cf3));
        _542_v196 = _nw74;
      }
      let _543_v197;
      _543_v197 = _dafny.Seq.of(_348_v69);
      if ((_dafny.MultiSet.fromElements(_543_v197)).equals(_dafny.MultiSet.fromElements(_543_v197, _543_v197, _543_v197))) {
        _309_v39 = (_309_v39).update((_306_v36).f41, new BigNumber((_dafny.Seq.UnicodeFromString("fy")).length));
        let _544_v198;
        let _nw75 = new _module.C7();
        _nw75.__ctor(((((_306_v36).f42).contains(new BigNumber(-183))) ? (((_306_v36).f42).get(new BigNumber(-183))) : (_264_v1)), (_306_v36).f42, new BigNumber((_349_v70).length), _267_v4, _267_v4, _267_v4);
        _544_v198 = _nw75;
        let _545_v199;
        _545_v199 = _dafny.Seq.of(_544_v198, _544_v198, _544_v198, _544_v198);
        let _546_v200;
        _546_v200 = _dafny.Map.Empty.slice().updateUnsafe(_415_v109,_545_v199);
        let _547_v201;
        _547_v201 = _dafny.Set.fromElements(new BigNumber((_269_v6).length));
        let _548_v202;
        _548_v202 = _dafny.Seq.of(_545_v199, _545_v199);
        let _549_v203;
        _549_v203 = _545_v199;
        let _550_v204;
        let _nw76 = Array((new BigNumber(27)).toNumber());
        _nw76[(_dafny.ZERO).toNumber()] = ((_267_v4) ? (_545_v199) : (_545_v199));
        _nw76[(_dafny.ONE).toNumber()] = _545_v199;
        _nw76[(new BigNumber(2)).toNumber()] = _545_v199;
        _nw76[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_545_v199, _545_v199);
        _nw76[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_545_v199, _545_v199);
        _nw76[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_545_v199, _545_v199);
        _nw76[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_545_v199, _dafny.Seq.of(_544_v198));
        _nw76[(new BigNumber(7)).toNumber()] = _545_v199;
        _nw76[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_545_v199, _545_v199);
        _nw76[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_545_v199, _module.__default.safeIndex(new BigNumber(473), new BigNumber((_545_v199).length)), _544_v198), _545_v199);
        _nw76[(new BigNumber(10)).toNumber()] = _545_v199;
        _nw76[(new BigNumber(11)).toNumber()] = _545_v199;
        _nw76[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(_544_v198);
        _nw76[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_545_v199, _module.__default.safeIndex((_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))], new BigNumber((_545_v199).length)), _544_v198);
        _nw76[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_545_v199, _545_v199);
        _nw76[(new BigNumber(15)).toNumber()] = _545_v199;
        _nw76[(new BigNumber(16)).toNumber()] = _545_v199;
        _nw76[(new BigNumber(17)).toNumber()] = _545_v199;
        _nw76[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_545_v199, (((_546_v200).contains(_415_v109)) ? ((_546_v200).get(_415_v109)) : (_545_v199)));
        _nw76[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(_544_v198, _544_v198, _544_v198, _544_v198, _544_v198);
        _nw76[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_545_v199, _dafny.Seq.update(_dafny.Seq.update(_545_v199, _module.__default.safeIndex(new BigNumber((_547_v201).length), new BigNumber((_545_v199).length)), _544_v198), _module.__default.safeIndex((_306_v36).f41, new BigNumber((_dafny.Seq.update(_545_v199, _module.__default.safeIndex(new BigNumber((_547_v201).length), new BigNumber((_545_v199).length)), _544_v198)).length)), _544_v198));
        _nw76[(new BigNumber(21)).toNumber()] = (_548_v202)[_module.__default.safeIndex(((((_306_v36).f42).contains(new BigNumber(576))) ? (((_306_v36).f42).get(new BigNumber(576))) : (new BigNumber(-151))), new BigNumber((_548_v202).length))];
        _nw76[(new BigNumber(22)).toNumber()] = (_549_v203);
        _nw76[(new BigNumber(23)).toNumber()] = _dafny.Seq.of(_544_v198);
        _nw76[(new BigNumber(24)).toNumber()] = _545_v199;
        _nw76[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_545_v199, _545_v199);
        _nw76[(new BigNumber(26)).toNumber()] = _dafny.Seq.Concat(_545_v199, _545_v199);
        _550_v204 = _nw76;
        let _index45 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_550_v204).length));
        (_550_v204)[_index45] = _dafny.Seq.Concat(_545_v199, _dafny.Seq.of(_544_v198));
        let _index46 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_550_v204).length));
        (_550_v204)[_index46] = _dafny.Seq.Concat(((_544_v198.f28) ? (_545_v199) : (_545_v199)), _545_v199);
        let _551_v205;
        let _552_v206;
        let _553_v207;
        let _554_v208;
        let _out34;
        let _out35;
        let _out36;
        let _out37;
        let _outcollector11 = (_306_v36).m7(_544_v198.f31, _272_v9, _544_v198.f29, _275_globalState);
        _out34 = _outcollector11[0];
        _out35 = _outcollector11[1];
        _out36 = _outcollector11[2];
        _out37 = _outcollector11[3];
        _551_v205 = _out34;
        _552_v206 = _out35;
        _553_v207 = _out36;
        _554_v208 = _out37;
        let _555_v209;
        let _nw77 = new _module.C1();
        _nw77.__ctor(!((_306_v36).fm19(_275_globalState)), false, _544_v198.f28);
        _555_v209 = _nw77;
        let _556_v210;
        _556_v210 = _module.D10.create_DC28(_555_v209);
        let _pat_let_tv4 = _555_v209;
        let _557_v211;
        let _nw78 = Array((new BigNumber(17)).toNumber());
        _nw78[(_dafny.ZERO).toNumber()] = _556_v210;
        _nw78[(_dafny.ONE).toNumber()] = _556_v210;
        _nw78[(new BigNumber(2)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(3)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(4)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(5)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(6)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(7)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(8)).toNumber()] = function (_pat_let7_0) {
          return function (_558_dt__update__tmp_h3) {
            return function (_pat_let8_0) {
              return function (_559_dt__update_hcf55_h1) {
                return _module.D10.create_DC28(_559_dt__update_hcf55_h1);
              }(_pat_let8_0);
            }(_pat_let_tv4);
          }(_pat_let7_0);
        }(_556_v210);
        _nw78[(new BigNumber(9)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(10)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(11)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(12)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(13)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(14)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(15)).toNumber()] = _556_v210;
        _nw78[(new BigNumber(16)).toNumber()] = _556_v210;
        _557_v211 = _nw78;
        let _560_v212;
        _560_v212 = _module.D20.create_DC50(_module.D20.create_DC48((_306_v36).f41, _557_v211, (_555_v209).f43));
        let _source10 = _560_v212;
        if (_source10.is_DC48) {
          let _561___mcc_h16 = (_source10).cf82;
          let _562___mcc_h17 = (_source10).cf83;
          let _563___mcc_h18 = (_source10).cf84;
          let _564_cf84 = _563___mcc_h18;
          let _565_cf83 = _562___mcc_h17;
          let _566_cf82 = _561___mcc_h16;
          let _567_v213;
          let _568_v214;
          let _569_v215;
          let _out38;
          let _out39;
          let _out40;
          let _outcollector12 = (_415_v109).m1(_275_globalState);
          _out38 = _outcollector12[0];
          _out39 = _outcollector12[1];
          _out40 = _outcollector12[2];
          _567_v213 = _out38;
          _568_v214 = _out39;
          _569_v215 = _out40;
          _568_v214 = (_module.__default.safeModuloInt((_306_v36).f41, new BigNumber(50))).minus(new BigNumber((_345_v66).length));
          let _570_v216;
          _570_v216 = _dafny.Seq.of(_349_v70);
          _349_v70 = (_570_v216)[_module.__default.safeIndex(_553_v207, new BigNumber((_570_v216).length))];
          let _index47 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
          (_263_v0)[_index47] = (_306_v36).f41;
        } else if (_source10.is_DC49) {
          let _571___mcc_h19 = (_source10).cf85;
          let _572___mcc_h20 = (_source10).cf86;
          let _573_cf86 = _572___mcc_h20;
          let _574_cf85 = _571___mcc_h19;
          (_544_v198).f31 = (_555_v209).f43;
          let _575_v217;
          let _576_v218;
          let _577_v219;
          let _out41;
          let _out42;
          let _out43;
          let _outcollector13 = (_415_v109).m1(_275_globalState);
          _out41 = _outcollector13[0];
          _out42 = _outcollector13[1];
          _out43 = _outcollector13[2];
          _575_v217 = _out41;
          _576_v218 = _out42;
          _577_v219 = _out43;
          (_544_v198).f29 = _dafny.areEqual(_269_v6, ((_267_v4) ? (_577_v219) : (_269_v6)));
          let _578_v220;
          _578_v220 = _module.D4.create_DC11((_544_v198).f30, _574_cf85, new BigNumber(854));
          let _579_v221;
          _579_v221 = _dafny.Map.Empty.slice().updateUnsafe(_551_v205,(_555_v209).f43);
          let _580_v222;
          let _nw79 = new _module.C10();
          _nw79.__ctor(_551_v205, (_578_v220).dtor_cf25, (((_555_v209).fm22(_576_v218, _module.__default.fm24(false, (_306_v36).f41, _275_globalState), _544_v198.f29, _275_globalState)) ? (_544_v198.f29) : ((_555_v209).f43)), _345_v66, _414_v108, _264_v1, _544_v198.f31, _544_v198.f31, (((_579_v221).contains(_573_cf86)) ? ((_579_v221).get(_573_cf86)) : (_544_v198.f28)));
          _580_v222 = _nw79;
        } else if (_source10.is_DC47) {
          let _581___mcc_h21 = (_source10).cf81;
          let _582_cf81 = _581___mcc_h21;
          let _583_v223;
          let _nw80 = new _module.C4();
          _nw80.__ctor(_264_v1, _544_v198.f31, _345_v66, _414_v108, (_306_v36).f41, (_555_v209).f43, _544_v198.f31, _544_v198.f29);
          _583_v223 = _nw80;
          let _584_v224;
          _584_v224 = _dafny.Seq.of(_583_v223);
          let _585_v225;
          let _nw81 = new _module.C10();
          _nw81.__ctor(_310_v40, new BigNumber((_584_v224).length), _583_v223.f31, (_583_v223).f32, _414_v108, _553_v207, _583_v223.f28, _267_v4, _583_v223.f29);
          _585_v225 = _nw81;
          let _586_v226;
          let _nw82 = Array((new BigNumber(20)).toNumber());
          _nw82[(_dafny.ZERO).toNumber()] = _585_v225;
          _nw82[(_dafny.ONE).toNumber()] = _585_v225;
          _nw82[(new BigNumber(2)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(3)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(4)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(5)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(6)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(7)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(8)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(9)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(10)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(11)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(12)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(13)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(14)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(15)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(16)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(17)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(18)).toNumber()] = _585_v225;
          _nw82[(new BigNumber(19)).toNumber()] = _585_v225;
          _586_v226 = _nw82;
          _586_v226 = _586_v226;
          (_275_globalState).f13 = (_583_v223).f30;
          let _587_v227;
          _587_v227 = _dafny.Map.Empty.slice().updateUnsafe(_583_v223.f28,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(834)), ((_588_v40) => function (_589_i20) {
            return _588_v40;
          })(_310_v40))).length));
          let _590_v228;
          _590_v228 = _dafny.Set.fromElements(_544_v198);
          let _591_v229;
          _591_v229 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-925),_module.__default.fm41(_583_v223.f31, _310_v40, _544_v198.f28, new BigNumber((_590_v228).length), _275_globalState));
          let _592_v230;
          let _nw83 = new _module.C6();
          _nw83.__ctor(_310_v40, new BigNumber(925), _583_v223.f35, _dafny.Seq.of(_583_v223.f31), (_583_v223).f33, (((_587_v227).contains(false)) ? ((_587_v227).get(false)) : (new BigNumber((_591_v229).length))), _583_v223.f31, false, _544_v198.f31);
          _592_v230 = _nw83;
          let _593_v231;
          _593_v231 = _dafny.Map.Empty.slice().updateUnsafe(_592_v230,(_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))]);
          let _594_v232;
          _594_v232 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_592_v230,new BigNumber(((_306_v36).f42).cardinality())), _593_v231);
          let _595_v233;
          let _nw84 = new _module.C6();
          _nw84.__ctor(_310_v40, _module.__default.safeModuloInt(new BigNumber(723), new BigNumber((_587_v227).length)), (_547_v201).IsSubsetOf(_547_v201), _345_v66, _module.__default.fm51(_544_v198.f31, _310_v40, _275_globalState), _264_v1, _544_v198.f28, _544_v198.f29, _dafny.Seq.contains(_594_v232, _593_v231));
          _595_v233 = _nw84;
          let _index48 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_272_v9).length));
          (_272_v9)[_index48] = !(_583_v223.f31);
          let _index49 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_272_v9).length));
          (_272_v9)[_index49] = _592_v230.f31;
        } else {
          let _596___mcc_h22 = (_source10).cf87;
          let _597_cf87 = _596___mcc_h22;
          let _598_v234;
          let _init9 = function (_599_i21) {
            return _dafny.Seq.UnicodeFromString("vtf");
          };
          let _nw85 = Array((new BigNumber(15)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw85.length); _i0_9++) {
            _nw85[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _598_v234 = _nw85;
          _598_v234 = _598_v234;
          (_275_globalState).f13 = (_553_v207).plus(((_544_v198.f31) ? ((_dafny.ZERO).minus((_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))])) : (_553_v207)));
          _553_v207 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_544_v198).f30), (_306_v36).f41));
          let _600_v235;
          _600_v235 = _dafny.Map.Empty.slice().updateUnsafe(_547_v201,_269_v6);
          (_275_globalState).f13 = _module.__default.fm32(!((_600_v235).equals(_600_v235)), _275_globalState);
        }
        let _601_v236;
        let _init10 = ((_602_v198) => function (_603_i22) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_602_v198.f31,!(_602_v198.f29))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_602_v198.f28,!(false)));
        })(_544_v198);
        let _nw86 = Array((new BigNumber(22)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw86.length); _i0_10++) {
          _nw86[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _601_v236 = _nw86;
        let _604_v237;
        _604_v237 = _module.D22.create_DC55((_544_v198).f30, _264_v1, (_555_v209).f43);
        let _605_v238;
        _605_v238 = _dafny.Map.Empty.slice().updateUnsafe(_544_v198.f29,(_604_v237).dtor_cf94);
        let _index50 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_601_v236).length));
        (_601_v236)[_index50] = _605_v238;
        let _606_v239;
        let _nw87 = new _module.C8();
        _nw87.__ctor(false, _551_v205, new BigNumber(694), !((_555_v209).f43), _dafny.Seq.update(_345_v66, _module.__default.safeIndex(_264_v1, new BigNumber((_345_v66).length)), (_555_v209).f43), _dafny.Seq.of(_345_v66), (_306_v36).f41, (_555_v209).f43, _544_v198.f31, false);
        _606_v239 = _nw87;
        let _607_v240;
        _607_v240 = _dafny.Map.Empty.slice().updateUnsafe(_544_v198.f31,_606_v239);
        let _608_v241;
        let _nw88 = Array((new BigNumber(6)).toNumber());
        _nw88[(_dafny.ZERO).toNumber()] = (_607_v240).update(_606_v239.f31, _606_v239);
        _nw88[(_dafny.ONE).toNumber()] = (_607_v240).Merge(_607_v240);
        _nw88[(new BigNumber(2)).toNumber()] = ((_544_v198.f28) ? (_607_v240) : (_607_v240));
        _nw88[(new BigNumber(3)).toNumber()] = _607_v240;
        _nw88[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_606_v239.f31,_606_v239);
        _nw88[(new BigNumber(5)).toNumber()] = _607_v240;
        _608_v241 = _nw88;
        let _609_v242;
        let _nw89 = Array((new BigNumber(27)).toNumber()).fill([]);
        _609_v242 = _nw89;
        let _610_v243;
        _610_v243 = _dafny.Map.Empty.slice().updateUnsafe(_609_v242,!(!(_606_v239.f28) || (_606_v239.f35)));
        let _index51 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_601_v236).length));
        let _index52 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
        let _rhs58 = _605_v238;
        let _rhs59 = _608_v241;
        let _rhs60 = (((_267_v4) ? (new BigNumber((_547_v201).length)) : (new BigNumber(797)))).isLessThan((_606_v239).f34);
        let _rhs61 = ((_306_v36).f41).plus((_263_v0)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length))]);
        let _rhs62 = (((_610_v243).contains(_609_v242)) ? ((_610_v243).get(_609_v242)) : (_606_v239.f29));
        let _lhs46 = _601_v236;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_601_v236).length));
        let _lhs48 = _275_globalState;
        let _lhs49 = _263_v0;
        let _lhs50 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
        let _lhs51 = _275_globalState;
        _lhs46[_lhs47] = _rhs58;
        _608_v241 = _rhs59;
        _lhs48.f1 = _rhs60;
        _lhs49[_lhs50] = _rhs61;
        _lhs51.f12 = _rhs62;
      } else {
        let _rhs63 = _263_v0;
        let _rhs64 = (_306_v36).fm19(_275_globalState);
        let _lhs52 = _275_globalState;
        _263_v0 = _rhs63;
        _lhs52.f1 = _rhs64;
        let _611_v244;
        _611_v244 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false),(_306_v36).f41)).length), new BigNumber((_269_v6).length));
        (_275_globalState).f1 = !((new BigNumber((_611_v244).length)).isLessThan((_306_v36).f41));
        let _index53 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
        (_263_v0)[_index53] = ((_306_v36).f41).plus((_306_v36).f41);
        let _index54 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length));
        (_272_v9)[_index54] = _267_v4;
        let _index55 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length));
        (_272_v9)[_index55] = !(((_267_v4) ? (!(_267_v4) || (_267_v4)) : (_267_v4)));
        if ((_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))]) {
          (_275_globalState).f1 = (_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))];
          let _index56 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
          (_263_v0)[_index56] = new BigNumber(501);
          let _612_v245;
          let _nw90 = new _module.C1();
          _nw90.__ctor((_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))], (_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))], _267_v4);
          _612_v245 = _nw90;
          _264_v1 = (_306_v36).f41;
          let _613_v246;
          _613_v246 = _module.D7.create_DC20(_module.D7.create_DC19((_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))], (_306_v36).f41, new BigNumber(-249)));
          let _614_v247;
          _614_v247 = _dafny.Map.Empty.slice().updateUnsafe(_613_v246,_345_v66);
          let _615_v248;
          _615_v248 = _module.D7.create_DC19((_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))], (_306_v36).f41, (_306_v36).f41);
          _614_v247 = (_614_v247).update(_module.D7.create_DC20(_615_v248), _dafny.Seq.of(_267_v4));
        } else {
          let _616_v249;
          let _init11 = ((_617_v9) => function (_618_i23) {
            return (_617_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_617_v9).length))];
          })(_272_v9);
          let _nw91 = Array((new BigNumber(16)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw91.length); _i0_11++) {
            _nw91[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _616_v249 = _nw91;
          let _619_v250;
          let _620_v251;
          let _621_v252;
          let _622_v253;
          let _out44;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector14 = (_306_v36).m7(_267_v4, _616_v249, (_306_v36).fm19(_275_globalState), _275_globalState);
          _out44 = _outcollector14[0];
          _out45 = _outcollector14[1];
          _out46 = _outcollector14[2];
          _out47 = _outcollector14[3];
          _619_v250 = _out44;
          _620_v251 = _out45;
          _621_v252 = _out46;
          _622_v253 = _out47;
          let _index57 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
          let _rhs65 = (_306_v36).f41;
          let _rhs66 = _dafny.Seq.Concat(_269_v6, _622_v253);
          let _rhs67 = _269_v6;
          let _rhs68 = _module.__default.safeDivisionInt((_306_v36).f41, _module.__default.fm32((_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))], _275_globalState));
          let _lhs53 = _263_v0;
          let _lhs54 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
          _lhs53[_lhs54] = _rhs65;
          _269_v6 = _rhs66;
          _622_v253 = _rhs67;
          _621_v252 = _rhs68;
          let _623_v254;
          let _nw92 = new _module.C5();
          _nw92.__ctor((_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))], _module.__default.safeModuloInt(_module.__default.fm24(false, (_306_v36).f41, _275_globalState), _621_v252), _264_v1, (_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))], _345_v66, _dafny.Seq.Concat(_414_v108, _414_v108), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))])).length)), !(!((_267_v4) || (_267_v4))), (_264_v1).isEqualTo((_306_v36).f41), !((_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))]) || (false));
          _623_v254 = _nw92;
          let _624_v255;
          let _nw93 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _624_v255 = _nw93;
          let _index58 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_624_v255).length));
          (_624_v255)[_index58] = ((_267_v4) ? (_310_v40) : (_310_v40));
          let _625_v256;
          _625_v256 = _dafny.Map.Empty.slice().updateUnsafe(_415_v109,(_306_v36).f41);
          let _index59 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_624_v255).length));
          (_624_v255)[_index59] = ((((_623_v254).f48).isLessThan(new BigNumber(99))) ? (_619_v250) : (_module.__default.fm48((_272_v9)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_272_v9).length))], true, new BigNumber((_625_v256).length), _275_globalState)));
          (_275_globalState).f13 = new BigNumber(-574);
        }
      }
      let _626_v257;
      _626_v257 = _dafny.MultiSet.fromElements(_272_v9, _272_v9);
      let _index60 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_263_v0).length));
      (_263_v0)[_index60] = (_module.__default.safeModuloInt(new BigNumber((_626_v257).cardinality()), (_306_v36).f41)).plus(new BigNumber(329));
      process.stdout.write(_dafny.toString((_263_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_264_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_265_v2, _dafny.Seq.of(new BigNumber(20)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_266_v3).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_267_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(20)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_269_v6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_270_v7, _dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0))), _dafny.Seq.UnicodeFromString("ssisbjr")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_275_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_275_globalState).f5).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(20)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_275_globalState).f7).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_275_globalState).f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_275_globalState.f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_275_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_275_globalState).f18)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState.f21)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_275_globalState).f22)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_275_globalState).f25, _dafny.Seq.of(new BigNumber(20)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_275_globalState.f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_globalState).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v36).f41));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_306_v36).f42).equals(_dafny.MultiSet.fromElements(new BigNumber(20)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_307_v37).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_308_v38).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_v39).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(20),_dafny.ZERO).updateUnsafe(_dafny.ZERO,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_310_v40));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_311_v41).equals(_dafny.Set.fromElements(new _dafny.CodePoint('m'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_345_v66, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_346_v67).dtor_cf46));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_346_v67).dtor_cf47));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_346_v67).dtor_cf48, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_346_v67).dtor_cf49));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_346_v67).dtor_cf50, _dafny.Seq.of(new BigNumber(20)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_347_v68).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_348_v69).dtor_cf15).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_348_v69).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_348_v69).dtor_cf17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_349_v70).equals(_dafny.Set.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_350_i8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_414_v108, _dafny.Seq.of(_dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_418_v110).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_543_v197, _dafny.Seq.of(_module.D3.create_DC8(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0))), true, new BigNumber(417))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_626_v257).cardinality())));
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
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.MultiSet.Empty);
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
    static create_DC3(cf3, cf4, cf5, cf6) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D1(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC4(cf7) {
      let $dt = new D1(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(false, _dafny.ZERO, null, _dafny.ZERO);
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
    static create_DC6(cf9, cf10, cf11, cf12, cf13) {
      let $dt = new D2(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC5(cf8) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf8 === other.cf8;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.ZERO, false, _dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC8(cf15, cf16, cf17) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC9(cf18, cf19, cf20, cf21) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC7(cf14) {
      let $dt = new D3(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf18 === other.cf18 && this.cf19 === other.cf19 && this.cf20 === other.cf20 && this.cf21 === other.cf21;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(_dafny.MultiSet.Empty, false, _dafny.ZERO);
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
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC13(cf27, cf28, cf29) {
      let $dt = new D5(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC14(cf30, cf31) {
      let $dt = new D5(1);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC12(cf26) {
      let $dt = new D5(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC12" + "(" + this.cf26.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf27 === other.cf27 && _dafny.areEqual(this.cf28, other.cf28) && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(false, _dafny.Map.Empty, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC16(cf33) {
      let $dt = new D6(0);
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC15(cf32) {
      let $dt = new D6(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC17(cf34) {
      let $dt = new D6(2);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf33 === other.cf33;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16([]);
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
    static create_DC19(cf36, cf37, cf38) {
      let $dt = new D7(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC18(cf35) {
      let $dt = new D7(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC20(cf39) {
      let $dt = new D7(2);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37) && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC22(cf41, cf42) {
      let $dt = new D8(0);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC23(cf43, cf44, cf45) {
      let $dt = new D8(1);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC24(cf46, cf47, cf48, cf49, cf50) {
      let $dt = new D8(2);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC21(cf40) {
      let $dt = new D8(3);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC25(cf51) {
      let $dt = new D8(4);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get is_DC21() { return this.$tag === 3; }
    get is_DC25() { return this.$tag === 4; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 4) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf43 === other.cf43 && this.cf44 === other.cf44 && this.cf45 === other.cf45;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48) && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC22(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC27(cf53, cf54) {
      let $dt = new D9(0);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC26(cf52) {
      let $dt = new D9(1);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf53) + ", " + this.cf54.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf53 === other.cf53 && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf52 === other.cf52;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC27(null, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC29(cf56, cf57) {
      let $dt = new D10(0);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC30(cf58, cf59) {
      let $dt = new D10(1);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC31(cf60, cf61, cf62) {
      let $dt = new D10(2);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC28(cf55) {
      let $dt = new D10(3);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC32(cf63) {
      let $dt = new D10(4);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get is_DC28() { return this.$tag === 3; }
    get is_DC32() { return this.$tag === 4; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC30" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC31" + "(" + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 3) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 4) {
        return "D10.DC32" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf60 === other.cf60 && this.cf61 === other.cf61 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf55 === other.cf55;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf63, other.cf63);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29(_dafny.ZERO, false);
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
    static create_DC33(cf64) {
      let $dt = new D11(0);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf64) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
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
    static create_DC34(cf65) {
      let $dt = new D12(0);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf65, other.cf65);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
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
    static create_DC35(cf66) {
      let $dt = new D13(0);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf66, other.cf66);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC37(cf68, cf69, cf70) {
      let $dt = new D14(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC36(cf67) {
      let $dt = new D14(1);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68) && this.cf69 === other.cf69 && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf67, other.cf67);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC37(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC38(cf71) {
      let $dt = new D15(0);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf71 === other.cf71;
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
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC40(cf73, cf74) {
      let $dt = new D16(0);
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC39(cf72) {
      let $dt = new D16(1);
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC41(cf75) {
      let $dt = new D16(2);
      $dt.cf75 = cf75;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get is_DC41() { return this.$tag === 2; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf75() { return this.cf75; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC40" + "(" + this.cf73.toVerbatimString(true) + ", " + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf75) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf73, other.cf73) && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf72, other.cf72);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf75, other.cf75);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC40(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO);
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
    static create_DC43() {
      let $dt = new D17(0);
      return $dt;
    }
    static create_DC44(cf77, cf78) {
      let $dt = new D17(1);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC42(cf76) {
      let $dt = new D17(2);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get is_DC42() { return this.$tag === 2; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC43";
      } else if (this.$tag === 1) {
        return "D17.DC44" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC42" + "(" + _dafny.toString(this.cf76) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77) && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf76, other.cf76);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC43();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC45(cf79) {
      let $dt = new D18(0);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC45" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC46(cf80) {
      let $dt = new D19(0);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC46" + "(" + _dafny.toString(this.cf80) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf80, other.cf80);
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
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC48(cf82, cf83, cf84) {
      let $dt = new D20(0);
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      return $dt;
    }
    static create_DC49(cf85, cf86) {
      let $dt = new D20(1);
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC47(cf81) {
      let $dt = new D20(2);
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC50(cf87) {
      let $dt = new D20(3);
      $dt.cf87 = cf87;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get is_DC47() { return this.$tag === 2; }
    get is_DC50() { return this.$tag === 3; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf87() { return this.cf87; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC48" + "(" + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC49" + "(" + _dafny.toString(this.cf85) + ", " + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC47" + "(" + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 3) {
        return "D20.DC50" + "(" + _dafny.toString(this.cf87) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf82, other.cf82) && this.cf83 === other.cf83 && this.cf84 === other.cf84;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf85 === other.cf85 && _dafny.areEqual(this.cf86, other.cf86);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf81 === other.cf81;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf87, other.cf87);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC48(_dafny.ZERO, [], false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC52(cf89) {
      let $dt = new D21(0);
      $dt.cf89 = cf89;
      return $dt;
    }
    static create_DC51(cf88) {
      let $dt = new D21(1);
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC53(cf90) {
      let $dt = new D21(2);
      $dt.cf90 = cf90;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC51() { return this.$tag === 1; }
    get is_DC53() { return this.$tag === 2; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf90() { return this.cf90; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC52" + "(" + _dafny.toString(this.cf89) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC51" + "(" + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC53" + "(" + _dafny.toString(this.cf90) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf89, other.cf89);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf88, other.cf88);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf90, other.cf90);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC52(new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf92, cf93, cf94) {
      let $dt = new D22(0);
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC54(cf91) {
      let $dt = new D22(1);
      $dt.cf91 = cf91;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf91() { return this.cf91; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC55" + "(" + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC54" + "(" + _dafny.toString(this.cf91) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf92, other.cf92) && _dafny.areEqual(this.cf93, other.cf93) && this.cf94 === other.cf94;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf91 === other.cf91;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC55(_dafny.ZERO, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC56(cf95) {
      let $dt = new D23(0);
      $dt.cf95 = cf95;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get dtor_cf95() { return this.cf95; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC56" + "(" + _dafny.toString(this.cf95) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf95 === other.cf95;
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
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC58() {
      let $dt = new D24(0);
      return $dt;
    }
    static create_DC59(cf97) {
      let $dt = new D24(1);
      $dt.cf97 = cf97;
      return $dt;
    }
    static create_DC57(cf96) {
      let $dt = new D24(2);
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC60(cf98) {
      let $dt = new D24(3);
      $dt.cf98 = cf98;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC59() { return this.$tag === 1; }
    get is_DC57() { return this.$tag === 2; }
    get is_DC60() { return this.$tag === 3; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf98() { return this.cf98; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC58";
      } else if (this.$tag === 1) {
        return "D24.DC59" + "(" + _dafny.toString(this.cf97) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC57" + "(" + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 3) {
        return "D24.DC60" + "(" + _dafny.toString(this.cf98) + ")";
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
        return other.$tag === 1 && this.cf97 === other.cf97;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf98, other.cf98);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC58();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC61(cf99) {
      let $dt = new D25(0);
      $dt.cf99 = cf99;
      return $dt;
    }
    get is_DC61() { return this.$tag === 0; }
    get dtor_cf99() { return this.cf99; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC61" + "(" + _dafny.toString(this.cf99) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf99, other.cf99);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC63(cf101, cf102, cf103) {
      let $dt = new D26(0);
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC62(cf100) {
      let $dt = new D26(1);
      $dt.cf100 = cf100;
      return $dt;
    }
    static create_DC64(cf104) {
      let $dt = new D26(2);
      $dt.cf104 = cf104;
      return $dt;
    }
    get is_DC63() { return this.$tag === 0; }
    get is_DC62() { return this.$tag === 1; }
    get is_DC64() { return this.$tag === 2; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf104() { return this.cf104; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC63" + "(" + _dafny.toString(this.cf101) + ", " + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC62" + "(" + _dafny.toString(this.cf100) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC64" + "(" + _dafny.toString(this.cf104) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf101, other.cf101) && _dafny.areEqual(this.cf102, other.cf102) && this.cf103 === other.cf103;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf100, other.cf100);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf104, other.cf104);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC63(_dafny.ZERO, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC65(cf105) {
      let $dt = new D27(0);
      $dt.cf105 = cf105;
      return $dt;
    }
    get is_DC65() { return this.$tag === 0; }
    get dtor_cf105() { return this.cf105; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC65" + "(" + _dafny.toString(this.cf105) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf105, other.cf105);
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
          return D27.Default();
        }
      };
    }
  }

  $module.D28 = class D28 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC66(cf106) {
      let $dt = new D28(0);
      $dt.cf106 = cf106;
      return $dt;
    }
    get is_DC66() { return this.$tag === 0; }
    get dtor_cf106() { return this.cf106; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC66" + "(" + _dafny.toString(this.cf106) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf106, other.cf106);
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
          return D28.Default();
        }
      };
    }
  }

  $module.D29 = class D29 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC68(cf108, cf109, cf110, cf111) {
      let $dt = new D29(0);
      $dt.cf108 = cf108;
      $dt.cf109 = cf109;
      $dt.cf110 = cf110;
      $dt.cf111 = cf111;
      return $dt;
    }
    static create_DC67(cf107) {
      let $dt = new D29(1);
      $dt.cf107 = cf107;
      return $dt;
    }
    static create_DC69(cf112) {
      let $dt = new D29(2);
      $dt.cf112 = cf112;
      return $dt;
    }
    get is_DC68() { return this.$tag === 0; }
    get is_DC67() { return this.$tag === 1; }
    get is_DC69() { return this.$tag === 2; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf112() { return this.cf112; }
    toString() {
      if (this.$tag === 0) {
        return "D29.DC68" + "(" + _dafny.toString(this.cf108) + ", " + _dafny.toString(this.cf109) + ", " + _dafny.toString(this.cf110) + ", " + _dafny.toString(this.cf111) + ")";
      } else if (this.$tag === 1) {
        return "D29.DC67" + "(" + _dafny.toString(this.cf107) + ")";
      } else if (this.$tag === 2) {
        return "D29.DC69" + "(" + _dafny.toString(this.cf112) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf108, other.cf108) && _dafny.areEqual(this.cf109, other.cf109) && this.cf110 === other.cf110 && this.cf111 === other.cf111;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf107, other.cf107);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf112, other.cf112);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D29.create_DC68(_dafny.Seq.of(), _dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D29.Default();
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

  $module.T3 = class T3 {
  };

  $module.T4 = class T4 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = false;
      this.f12 = false;
      this.f13 = _dafny.ZERO;
      this.f21 = [];
      this.f26 = false;
      this._f0 = _dafny.ZERO;
      this._f2 = false;
      this._f3 = false;
      this._f4 = _dafny.ZERO;
      this._f5 = _dafny.Map.Empty;
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.Seq.UnicodeFromString("");
      this._f8 = _dafny.Seq.UnicodeFromString("");
      this._f9 = false;
      this._f10 = _dafny.ZERO;
      this._f11 = _dafny.ZERO;
      this._f14 = false;
      this._f15 = [];
      this._f16 = _dafny.ZERO;
      this._f17 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f18 = [];
      this._f19 = false;
      this._f20 = false;
      this._f22 = [];
      this._f23 = false;
      this._f24 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f25 = _dafny.Seq.of();
      this._f27 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this).f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      (_this).f26 = f26;
      (_this)._f27 = f27;
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
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f37 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f37) {
      let _this = this;
      (_this).f37 = f37;
      return;
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f28 = false;
      this._f29 = false;
      this._f43 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    __ctor(f43, f28, f29) {
      let _this = this;
      (_this)._f43 = f43;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm2(globalState) {
      let _this = this;
      let _source11 = _module.D3.create_DC8(_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))), false, new BigNumber(611));
      if (_source11.is_DC8) {
        let _627___mcc_h0 = (_source11).cf15;
        let _628___mcc_h1 = (_source11).cf16;
        let _629___mcc_h2 = (_source11).cf17;
        let _630_cf17 = _629___mcc_h2;
        let _631_cf16 = _628___mcc_h1;
        let _632_cf15 = _627___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe((_this).f43,_630_cf17)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f43,(_dafny.ZERO).minus(_630_cf17)));
      } else if (_source11.is_DC9) {
        let _633___mcc_h3 = (_source11).cf18;
        let _634___mcc_h4 = (_source11).cf19;
        let _635___mcc_h5 = (_source11).cf20;
        let _636___mcc_h6 = (_source11).cf21;
        let _637_cf21 = _636___mcc_h6;
        let _638_cf20 = _635___mcc_h5;
        let _639_cf19 = _634___mcc_h4;
        let _640_cf18 = _633___mcc_h3;
        return (_dafny.Map.Empty.slice().updateUnsafe((_this).f43,new BigNumber(513))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(615)));
      } else {
        let _641___mcc_h7 = (_source11).cf14;
        let _642_cf14 = _641___mcc_h7;
        return (_dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC6(new BigNumber(643), _this.f28, new BigNumber(167), (_this).f43, new BigNumber((_dafny.Seq.of(new BigNumber(-627))).length))).dtor_cf12,new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f43,new BigNumber(342)));
      }
    };
    fm22(p0, p1, p2, globalState) {
      let _this = this;
      return true;
    };
    m9(globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = new _dafny.CodePoint('D'.codePointAt(0));
      let r3 = false;
      let _643_v0;
      _643_v0 = new BigNumber(690);
      (_this).f28 = !(_module.__default.fm1(_this.f28, _643_v0, globalState));
      let _644_v1;
      _644_v1 = _dafny.Map.Empty.slice().updateUnsafe(_643_v0,_643_v0);
      let _645_v2;
      _645_v2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(((_644_v1).contains(new BigNumber(193))) ? ((_644_v1).get(new BigNumber(193))) : (new BigNumber((_dafny.Set.fromElements(_643_v0)).length))));
      let _646_v3;
      _646_v3 = _dafny.MultiSet.fromElements(new BigNumber((_645_v2).length));
      r3 = (_646_v3).equals((_646_v3).Union(_646_v3));
      _644_v1 = (_644_v1).Merge(_dafny.Map.Empty.slice().updateUnsafe(_643_v0,new BigNumber(854)));
      let _647_v4;
      _647_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(410),_this.f28);
      let _648_v5;
      _648_v5 = _dafny.Seq.of(_643_v0);
      _643_v0 = ((_dafny.ZERO).minus(new BigNumber(((_647_v4).update(_643_v0, _this.f28)).length))).multipliedBy((new BigNumber(545)).minus(new BigNumber((_648_v5).length)));
      let _649_v6;
      _649_v6 = _dafny.Seq.UnicodeFromString("avexbc");
      let _650_v7;
      _650_v7 = _dafny.MultiSet.fromElements((_this).f43, _this.f29, ((_this).f43) === ((_this).f43), _module.__default.fm1((_this).f43, _643_v0, globalState));
      let _651_v8;
      _651_v8 = _dafny.Set.fromElements(false);
      let _652_v9;
      _652_v9 = _dafny.Seq.of(_645_v2, _645_v2);
      let _653_v10;
      _653_v10 = _module.D3.create_DC7(_652_v9);
      let _654_v11;
      _654_v11 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_643_v0,_643_v0));
      let _655_v12;
      _655_v12 = _dafny.Set.fromElements(new BigNumber(297));
      let _pat_let_tv5 = _650_v7;
      let _pat_let_tv6 = _650_v7;
      let _pat_let_tv7 = _650_v7;
      let _pat_let_tv8 = _650_v7;
      let _rhs69 = _dafny.Seq.Concat(_649_v6, _649_v6);
      let _rhs70 = (_649_v6)[_module.__default.safeIndex(new BigNumber((_651_v8).length), new BigNumber((_649_v6).length))];
      let _rhs71 = function (_source12) {
        if (_source12.is_DC8) {
          let _656___mcc_h0 = (_source12).cf15;
          let _657___mcc_h1 = (_source12).cf16;
          let _658___mcc_h2 = (_source12).cf17;
          let _659_cf17 = _658___mcc_h2;
          let _660_cf16 = _657___mcc_h1;
          let _661_cf15 = _656___mcc_h0;
          return (_pat_let_tv5).Intersect((_module.D0.create_DC1(_pat_let_tv6)).dtor_cf1);
        } else if (_source12.is_DC9) {
          let _662___mcc_h3 = (_source12).cf18;
          let _663___mcc_h4 = (_source12).cf19;
          let _664___mcc_h5 = (_source12).cf20;
          let _665___mcc_h6 = (_source12).cf21;
          let _666_cf21 = _665___mcc_h6;
          let _667_cf20 = _664___mcc_h5;
          let _668_cf19 = _663___mcc_h4;
          let _669_cf18 = _662___mcc_h3;
          return _pat_let_tv7;
        } else {
          let _670___mcc_h7 = (_source12).cf14;
          let _671_cf14 = _670___mcc_h7;
          return _pat_let_tv8;
        }
      }(_653_v10);
      let _rhs72 = ((_644_v1).Merge(_644_v1)).Merge(((_654_v11)[_module.__default.safeIndex(new BigNumber((_649_v6).length), new BigNumber((_654_v11).length))]).Merge((_654_v11)[_module.__default.safeIndex(_643_v0, new BigNumber((_654_v11).length))]));
      let _rhs73 = (_this).fm22(new BigNumber((_649_v6).length), _643_v0, (_655_v12).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_648_v5).length))), globalState);
      let _lhs55 = globalState;
      _649_v6 = _rhs69;
      r2 = _rhs70;
      _650_v7 = _rhs71;
      _644_v1 = _rhs72;
      _lhs55.f12 = _rhs73;
      if (!(_module.__default.fm1((_this).f43, (_643_v0).multipliedBy(_643_v0), globalState))) {
        let _672_v13;
        _672_v13 = new _dafny.CodePoint('t'.codePointAt(0));
        let _673_v14;
        _673_v14 = _dafny.Map.Empty.slice().updateUnsafe(_672_v13,_648_v5);
        let _674_v15;
        let _nw94 = Array((new BigNumber(24)).toNumber());
        _nw94[(_dafny.ZERO).toNumber()] = _648_v5;
        _nw94[(_dafny.ONE).toNumber()] = _648_v5;
        _nw94[(new BigNumber(2)).toNumber()] = _module.__default.fm23(_this.f28, _module.__default.fm24(_this.f29, new BigNumber(489), globalState), globalState);
        _nw94[(new BigNumber(3)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_648_v5, _module.__default.safeIndex(_643_v0, new BigNumber((_648_v5).length)), _module.__default.fm24(false, _module.__default.fm24(false, _643_v0, globalState), globalState));
        _nw94[(new BigNumber(5)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(6)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(7)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(8)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_643_v0);
        _nw94[(new BigNumber(10)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(11)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(376)), ((_675_v6) => function (_676_i0) {
          return new BigNumber((_675_v6).length);
        })(_649_v6));
        _nw94[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(955)), ((_677_v0) => function (_678_i1) {
          return _677_v0;
        })(_643_v0));
        _nw94[(new BigNumber(14)).toNumber()] = _module.__default.fm23((_this).f43, new BigNumber((_649_v6).length), globalState);
        _nw94[(new BigNumber(15)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(16)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(17)).toNumber()] = (((_673_v14).contains(_672_v13)) ? ((_673_v14).get(_672_v13)) : (_648_v5));
        _nw94[(new BigNumber(18)).toNumber()] = _module.__default.fm23(_this.f29, _643_v0, globalState);
        _nw94[(new BigNumber(19)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(20)).toNumber()] = _648_v5;
        _nw94[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_643_v0, new BigNumber(698), new BigNumber(327)), _module.__default.safeIndex((_648_v5)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_648_v5).length))], new BigNumber((_dafny.Seq.of(_643_v0, new BigNumber(698), new BigNumber(327))).length)), new BigNumber((_646_v3).cardinality()));
        _nw94[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_648_v5, _648_v5);
        _nw94[(new BigNumber(23)).toNumber()] = ((_this.f28) ? (_648_v5) : (_648_v5));
        _674_v15 = _nw94;
        let _index61 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_674_v15).length));
        (_674_v15)[_index61] = _648_v5;
        let _679_v16;
        _679_v16 = _module.D6.create_DC15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-544)), ((_680_v0) => function (_681_i2) {
  return _680_v0;
})(_643_v0)));
        let _index62 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_674_v15).length));
        (_674_v15)[_index62] = (_679_v16).dtor_cf32;
        let _682_v17;
        _682_v17 = _module.D4.create_DC10((_646_v3).Union(_646_v3));
        let _source13 = _682_v17;
        if (_source13.is_DC11) {
          let _683___mcc_h8 = (_source13).cf23;
          let _684___mcc_h9 = (_source13).cf24;
          let _685___mcc_h10 = (_source13).cf25;
          let _686_cf25 = _685___mcc_h10;
          let _687_cf24 = _684___mcc_h9;
          let _688_cf23 = _683___mcc_h8;
          let _689_v18;
          _689_v18 = _dafny.MultiSet.fromElements(_672_v13);
          let _690_v19;
          _690_v19 = _dafny.Seq.of((_this).f43, _this.f28);
          (globalState).f12 = (_646_v3).IsDisjointFrom((_646_v3).update((_module.D3.create_DC8(_689_v18, (_690_v19)[_module.__default.safeIndex(_686_cf25, new BigNumber((_690_v19).length))], _688_cf23)).dtor_cf17, _module.__default.abs(new BigNumber((_649_v6).length))));
          let _691_v20;
          let _nw95 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _691_v20 = _nw95;
          let _index63 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_691_v20).length));
          (_691_v20)[_index63] = _module.__default.fm24(_687_cf24, new BigNumber(818), globalState);
          let _692_v21;
          _692_v21 = _dafny.Seq.of(_649_v6, _dafny.Seq.update(_dafny.Seq.Concat(_649_v6, _649_v6), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_643_v0,_690_v19)).length), new BigNumber((_dafny.Seq.Concat(_649_v6, _649_v6)).length)), _672_v13));
          let _index64 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_691_v20).length));
          let _rhs74 = new BigNumber(((_692_v21)[_module.__default.safeIndex((_686_cf25).minus(_module.__default.fm24((_this).f43, _module.__default.fm24(_this.f28, _688_cf23, globalState), globalState)), new BigNumber((_692_v21).length))]).length);
          let _rhs75 = _690_v19;
          let _lhs56 = _691_v20;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_691_v20).length));
          _lhs56[_lhs57] = _rhs74;
          _690_v19 = _rhs75;
          (globalState).f12 = _this.f29;
          _686_cf25 = ((_691_v20)[_module.__default.safeIndex(new BigNumber(251), new BigNumber((_691_v20).length))]).minus((_643_v0).plus(_686_cf25));
        } else {
          let _693___mcc_h11 = (_source13).cf22;
          let _694_cf22 = _693___mcc_h11;
          let _695_v22;
          let _nw96 = Array((new BigNumber(19)).toNumber()).fill(false);
          _695_v22 = _nw96;
          let _index65 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_695_v22).length));
          (_695_v22)[_index65] = _this.f29;
          let _index66 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_695_v22).length));
          (_695_v22)[_index66] = (_this).f43;
          (globalState).f12 = _this.f29;
          r1 = (_643_v0).plus((_643_v0).minus(new BigNumber((_dafny.Set.fromElements((((_644_v1).contains(_643_v0)) ? ((_644_v1).get(_643_v0)) : (_643_v0)))).length)));
          (globalState).f13 = new BigNumber((_649_v6).length);
        }
        _649_v6 = _649_v6;
        (globalState).f1 = ((_this).f43) === (true);
        (globalState).f13 = (_643_v0).minus(_643_v0);
      } else {
        if (_this.f29) {
          (globalState).f26 = (_this).f43;
          let _696_v23;
          let _nw97 = Array((new BigNumber(8)).toNumber()).fill(false);
          _696_v23 = _nw97;
          let _index67 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_696_v23).length));
          (_696_v23)[_index67] = _module.__default.fm1(_this.f28, _643_v0, globalState);
          let _index68 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_696_v23).length));
          (_696_v23)[_index68] = ((_this.f28) ? (_this.f29) : ((_this).f43));
          let _697_v24;
          _697_v24 = _dafny.MultiSet.fromElements(_648_v5);
          (globalState).f1 = (_697_v24).equals(_697_v24);
          r1 = _643_v0;
          r1 = _643_v0;
        } else {
          let _698_v25;
          let _nw98 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _698_v25 = _nw98;
          let _index69 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_698_v25).length));
          (_698_v25)[_index69] = (_643_v0).minus(_643_v0);
          let _index70 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_698_v25).length));
          (_698_v25)[_index70] = _module.__default.safeModuloInt(_643_v0, new BigNumber(15));
          r1 = _643_v0;
          let _index71 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_698_v25).length));
          (_698_v25)[_index71] = ((_this.f29) ? ((_698_v25)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_698_v25).length))]) : (((_698_v25)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_698_v25).length))]).multipliedBy((_698_v25)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_698_v25).length))])));
          let _699_v26;
          _699_v26 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_698_v25)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_698_v25).length))], (_698_v25)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_698_v25).length))])),_649_v6);
          _699_v26 = (_699_v26).update((_698_v25)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_698_v25).length))], _649_v6);
          let _700_v27;
          _700_v27 = _dafny.Seq.of((_this).f43);
          let _701_v28;
          _701_v28 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_649_v6).length)).isLessThan(_643_v0),(_700_v27)[_module.__default.safeIndex(new BigNumber((_650_v7).cardinality()), new BigNumber((_700_v27).length))]);
          _701_v28 = (_701_v28).update(_this.f29, _this.f28);
        }
        r2 = new _dafny.CodePoint('h'.codePointAt(0));
        let _702_v29;
        let _nw99 = Array((new BigNumber(12)).toNumber()).fill(false);
        _702_v29 = _nw99;
        let _index72 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_702_v29).length));
        (_702_v29)[_index72] = _this.f28;
        let _index73 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_702_v29).length));
        (_702_v29)[_index73] = !((_643_v0).isLessThanOrEqualTo(_643_v0));
        r3 = ((!(_this.f28) || ((_this).f43)) ? (!((_this).fm22(_643_v0, _643_v0, _this.f28, globalState))) : (_this.f29));
        let _703_v30;
        let _init12 = ((_704_v0) => function (_705_i3) {
          return (_705_i3).multipliedBy(_704_v0);
        })(_643_v0);
        let _nw100 = Array((new BigNumber(20)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw100.length); _i0_12++) {
          _nw100[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _703_v30 = _nw100;
        let _index74 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_703_v30).length));
        (_703_v30)[_index74] = (_643_v0).multipliedBy(_643_v0);
        let _706_v31;
        _706_v31 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_this.f28);
        let _707_v32;
        _707_v32 = _dafny.Set.fromElements(_706_v31, _706_v31);
        let _index75 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_703_v30).length));
        (_703_v30)[_index75] = new BigNumber((_707_v32).length);
      }
      let _708_v33;
      _708_v33 = new _dafny.CodePoint('r'.codePointAt(0));
      r0 = _708_v33;
      r1 = (_643_v0).minus(_643_v0);
      r2 = _708_v33;
      r3 = function (_source14) {
        if (_source14.is_DC16) {
          let _709___mcc_h12 = (_source14).cf33;
          let _710_cf33 = _709___mcc_h12;
          return (_this).f43;
        } else if (_source14.is_DC15) {
          let _711___mcc_h13 = (_source14).cf32;
          let _712_cf32 = _711___mcc_h13;
          return !(_this.f28);
        } else {
          let _713___mcc_h14 = (_source14).cf34;
          let _714_cf34 = _713___mcc_h14;
          return _this.f29;
        }
      }(_module.D6.create_DC15(_648_v5));
      return [r0, r1, r2, r3];
    }
    m10(globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _715_v0;
      _715_v0 = _dafny.Seq.UnicodeFromString("cjeniah");
      let _716_v1;
      _716_v1 = new BigNumber(229);
      (globalState).f13 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_715_v0).length))).multipliedBy(_716_v1));
      let _717_v2;
      let _init13 = ((_718_v1) => function (_719_i0) {
        return _module.__default.safeModuloInt(_719_i0, _718_v1);
      })(_716_v1);
      let _nw101 = Array((new BigNumber(20)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw101.length); _i0_13++) {
        _nw101[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _717_v2 = _nw101;
      let _720_v3;
      _720_v3 = _dafny.Set.fromElements(_717_v2);
      let _721_v4;
      let _nw102 = Array((new BigNumber(26)).toNumber());
      _721_v4 = _nw102;
      let _722_v5;
      _722_v5 = _module.D2.create_DC5(_721_v4);
      let _723_v6;
      _723_v6 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_717_v2)).IsSubsetOf(_720_v3),((_this.f29) ? (_722_v5) : (_module.D2.create_DC5(_721_v4))));
      let _rhs76 = (_dafny.ZERO).minus((_716_v1).minus(_module.__default.safeModuloInt(_716_v1, new BigNumber(390))));
      let _rhs77 = false;
      let _rhs78 = _723_v6;
      let _lhs58 = globalState;
      _716_v1 = _rhs76;
      _lhs58.f12 = _rhs77;
      _723_v6 = _rhs78;
      let _724_i1;
      _724_i1 = _dafny.ZERO;
      L1: {
        while (!(_this.f28)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_724_i1)) {
              break L1;
            }
            _724_i1 = (_724_i1).plus(_dafny.ONE);
            let _725_v7;
            let _nw103 = Array((new BigNumber(16)).toNumber()).fill(false);
            _725_v7 = _nw103;
            (globalState).f21 = _725_v7;
            let _726_v8;
            let _init14 = function (_727_i2) {
              return new _dafny.CodePoint('i'.codePointAt(0));
            };
            let _nw104 = Array((new BigNumber(9)).toNumber());
            for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw104.length); _i0_14++) {
              _nw104[_i0_14] = _init14(new BigNumber(_i0_14));
            }
            _726_v8 = _nw104;
            let _728_v9;
            let _nw105 = new _module.C0();
            _nw105.__ctor(_726_v8);
            _728_v9 = _nw105;
            let _729_v10;
            _729_v10 = _dafny.Seq.of(new BigNumber(111), _716_v1, _716_v1);
            let _730_v11;
            _730_v11 = _dafny.Seq.of(_this.f29);
            let _731_v12;
            let _out48;
            _out48 = _module.__default.m0(_729_v10, (_730_v11)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), function (_732_i3) {
              return new BigNumber(758);
            })).length), new BigNumber((_730_v11).length))], (_dafny.ZERO).minus(_716_v1), _this.f29, globalState);
            _731_v12 = _out48;
            let _733_v13;
            _733_v13 = _dafny.MultiSet.fromElements(_731_v12, _731_v12);
            let _734_v14;
            _734_v14 = _dafny.Set.fromElements(_733_v13);
            let _735_v15;
            _735_v15 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f28) ? (_734_v14) : (_734_v14)),_731_v12);
            _735_v15 = (_735_v15).update((_dafny.Set.fromElements(_module.__default.fm25(globalState))).Difference(_734_v14), new BigNumber((_729_v10).length));
          }
        }
      }
      let _736_v16;
      let _nw106 = Array((_dafny.ONE).toNumber()).fill(false);
      _736_v16 = _nw106;
      let _index76 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_736_v16).length));
      (_736_v16)[_index76] = (_this).f43;
      let _index77 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_736_v16).length));
      (_736_v16)[_index77] = _this.f29;
      let _737_v17;
      let _nw107 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _737_v17 = _nw107;
      let _738_v18;
      let _nw108 = new _module.C0();
      _nw108.__ctor(_737_v17);
      _738_v18 = _nw108;
      let _739_v19;
      _739_v19 = _dafny.Set.fromElements(_716_v1, _716_v1);
      _716_v1 = _module.__default.safeDivisionInt(new BigNumber(442), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_739_v19,_716_v1)).length));
      let _740_v20;
      _740_v20 = _dafny.Map.Empty.slice().updateUnsafe(!((_736_v16)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_736_v16).length))]),_dafny.Seq.of(_this.f28, !(!(!(_this.f29))), _this.f29, _this.f28, _this.f29));
      r0 = (_740_v20).Merge((_740_v20).Merge(_module.__default.fm26(_this.f29, globalState)));
      return r0;
    }
    get f43() {
      let _this = this;
      return _this._f43;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f28 = false;
      this._f29 = false;
      this._f31 = false;
      this._f30 = _dafny.ZERO;
      this._f46 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    set f31(value) {
      let _this = this;
      _this._f31 = value;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    __ctor(f46, f30, f31, f28, f29) {
      let _this = this;
      (_this)._f46 = f46;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm3(p0, p1, globalState) {
      let _this = this;
      return (_module.D5.create_DC12(_dafny.Seq.UnicodeFromString("xlpme"))).dtor_cf26;
    };
    fm2(globalState) {
      let _this = this;
      return (((_this.f28) ? (_dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_this).f30)) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_this).f30)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f29,new BigNumber(-74)));
    };
    fm29(globalState) {
      let _this = this;
      let _source15 = ((_this.f29) ? (_module.D5.create_DC13(_this.f29, _dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_this).f30), new _dafny.CodePoint('k'.codePointAt(0)))) : (_module.D5.create_DC13(_this.f31, _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f46), new _dafny.CodePoint('m'.codePointAt(0)))));
      if (_source15.is_DC13) {
        let _741___mcc_h0 = (_source15).cf27;
        let _742___mcc_h1 = (_source15).cf28;
        let _743___mcc_h2 = (_source15).cf29;
        let _744_cf29 = _743___mcc_h2;
        let _745_cf28 = _742___mcc_h1;
        let _746_cf27 = _741___mcc_h0;
        return (_this).f46;
      } else if (_source15.is_DC14) {
        let _747___mcc_h3 = (_source15).cf30;
        let _748___mcc_h4 = (_source15).cf31;
        let _749_cf31 = _748___mcc_h4;
        let _750_cf30 = _747___mcc_h3;
        return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_749_cf31))).cardinality());
      } else {
        let _751___mcc_h5 = (_source15).cf26;
        let _752_cf26 = _751___mcc_h5;
        return (_this).f46;
      }
    };
    fm30(p0, globalState) {
      let _this = this;
      return _this.f28;
    };
    m13(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let _753_v0;
      _753_v0 = _dafny.Seq.of((_dafny.ZERO).minus(p3));
      let _754_i0;
      _754_i0 = _dafny.ZERO;
      L2: {
        while (_dafny.Seq.IsProperPrefixOf(_753_v0, _753_v0)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_754_i0)) {
              break L2;
            }
            _754_i0 = (_754_i0).plus(_dafny.ONE);
            let _755_v1;
            let _nw109 = Array((new BigNumber(11)).toNumber()).fill(false);
            _755_v1 = _nw109;
            let _756_v2;
            _756_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_755_v1)).cardinality())));
            (globalState).f13 = ((p2) ? (p3) : ((((_756_v2).contains((_this).f46)) ? ((_756_v2).get((_this).f46)) : ((_this).f30))));
            (globalState).f13 = (_this).f30;
            if (_this.f31) {
              (globalState).f12 = p2;
              let _757_v3;
              _757_v3 = _dafny.Seq.UnicodeFromString("lgyuevq");
              _757_v3 = (_this).fm3(_this.f28, (p3).isLessThanOrEqualTo(p3), globalState);
              (globalState).f13 = (((p1).contains(_this.f28)) ? ((p1).get(_this.f28)) : (((!(true)) ? ((_this).f46) : (new BigNumber((_757_v3).length)))));
              _753_v0 = _dafny.Seq.of((new BigNumber(-583)).minus((_this).f46));
              (globalState).f13 = (_this).f46;
            } else {
              (globalState).f13 = (_this).f46;
              (globalState).f13 = (_this).f46;
              let _index78 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_755_v1).length));
              (_755_v1)[_index78] = _this.f31;
              let _758_v4;
              _758_v4 = _dafny.Seq.UnicodeFromString("d");
              let _index79 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_755_v1).length));
              (_755_v1)[_index79] = (new BigNumber((_758_v4).length)).isLessThan((_this).f46);
              let _index80 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_755_v1).length));
              (_755_v1)[_index80] = (_755_v1)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_755_v1).length))];
              let _759_v5;
              let _nw110 = Array((new BigNumber(4)).toNumber()).fill(false);
              _759_v5 = _nw110;
              let _760_v6;
              _760_v6 = _dafny.Map.Empty.slice().updateUnsafe(_759_v5,_dafny.Seq.Concat(_758_v4, _758_v4));
              _760_v6 = (_760_v6).update(_759_v5, _758_v4);
            }
            let _761_v7;
            let _nw111 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Set.Empty);
            _761_v7 = _nw111;
            let _762_v8;
            _762_v8 = _dafny.Set.fromElements(_this.f31, _this.f28);
            let _index81 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_761_v7).length));
            (_761_v7)[_index81] = _762_v8;
            let _763_v9;
            _763_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_this.f29);
            let _index82 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_761_v7).length));
            let _rhs79 = (_dafny.Set.fromElements(_this.f28, _this.f29)).Union(_762_v8);
            let _rhs80 = (((true) === (_this.f31)) ? ((_this).f46) : ((_this).f46));
            let _rhs81 = ((_this).f46).plus((new BigNumber((_763_v9).length)).multipliedBy((_this).f30));
            let _lhs59 = _761_v7;
            let _lhs60 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_761_v7).length));
            let _lhs61 = globalState;
            let _lhs62 = globalState;
            _lhs59[_lhs60] = _rhs79;
            _lhs61.f13 = _rhs80;
            _lhs62.f13 = _rhs81;
          }
        }
      }
      let _764_v10;
      let _init15 = ((_765_p3) => function (_766_i2) {
        return (_dafny.Set.fromElements((_this).f46, new BigNumber(-550))).Union(_dafny.Set.fromElements(_765_p3));
      })(p3);
      let _nw112 = Array((new BigNumber(27)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw112.length); _i0_15++) {
        _nw112[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _764_v10 = _nw112;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_764_v10).length))) {
        let _767_i1 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_767_i1)) && ((_767_i1).isLessThan(new BigNumber((_764_v10).length))))) {
          (_764_v10)[(_767_i1)] = ((_dafny.Set.fromElements((_this).f30, (_this).f30)).Difference(_dafny.Set.fromElements((_753_v0)[_module.__default.safeIndex((_this).f46, new BigNumber((_753_v0).length))]))).Intersect((_dafny.Set.fromElements(p3)).Union(_dafny.Set.fromElements(new BigNumber(950))));
        }
      }
      let _768_v11;
      let _nw113 = new _module.C1();
      _nw113.__ctor(_this.f29, _this.f29, !(true) || (!(_this.f28)));
      _768_v11 = _nw113;
      let _769_v12;
      let _nw114 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _769_v12 = _nw114;
      let _770_v13;
      let _nw115 = new _module.C0();
      _nw115.__ctor(_769_v12);
      _770_v13 = _nw115;
      let _771_v14;
      _771_v14 = _dafny.Seq.of(true);
      let _772_v15;
      _772_v15 = _dafny.Seq.UnicodeFromString("xuygtgj");
      _771_v14 = _dafny.Seq.update(_module.__default.fm0(_module.__default.safeModuloInt(new BigNumber((_772_v15).length), (_this).f30), _this.f29, _this.f29, new BigNumber(808), globalState), _module.__default.safeIndex((_this).f30, new BigNumber((_module.__default.fm0(_module.__default.safeModuloInt(new BigNumber((_772_v15).length), (_this).f30), _this.f29, _this.f29, new BigNumber(808), globalState)).length)), true);
      let _773_v16;
      _773_v16 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(p2, new BigNumber((p1).length), globalState),_770_v13.f37);
      let _774_v17;
      let _nw116 = new _module.C0();
      _nw116.__ctor((((_773_v16).contains(p2)) ? ((_773_v16).get(p2)) : (_770_v13.f37)));
      _774_v17 = _nw116;
      r0 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(61)), function (_775_i3) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), _772_v15)).length)).isLessThanOrEqualTo((_753_v0)[_module.__default.safeIndex((_this).f30, new BigNumber((_753_v0).length))]);
      return r0;
    }
    get f46() {
      let _this = this;
      return _this._f46;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f28 = false;
      this._f29 = false;
      this._f31 = false;
      this._f35 = false;
      this._f36 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f34 = _dafny.ZERO;
      this._f32 = _dafny.Seq.of();
      this._f33 = _dafny.Seq.of();
      this._f30 = _dafny.ZERO;
      this._f44 = _dafny.ZERO;
      this._f45 = [];
    }
    _parentTraits() {
      return [_module.T4, _module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    set f31(value) {
      let _this = this;
      _this._f31 = value;
    };
    get f35() {
      let _this = this;
      return _this._f35;
    };
    set f35(value) {
      let _this = this;
      _this._f35 = value;
    };
    get f36() {
      let _this = this;
      return _this._f36;
    };
    set f36(value) {
      let _this = this;
      _this._f36 = value;
    };
    get f34() {
      let _this = this;
      return _this._f34;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    __ctor(f44, f45, f36, f34, f35, f32, f33, f30, f31, f28, f29) {
      let _this = this;
      (_this)._f44 = f44;
      (_this)._f45 = f45;
      (_this)._f36 = f36;
      (_this)._f34 = f34;
      (_this)._f35 = f35;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm7(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vsrndyx"), _dafny.Seq.UnicodeFromString("jsqkph")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mfxuo"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(673)), function (_776_i0) {
        return _this.f36;
      })));
    };
    fm6(globalState) {
      let _this = this;
      return _this.f29;
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Set.fromElements(((_this.f35) ? ((_this).f34) : ((_dafny.ZERO).minus((_this).f34))), (_this).f34, ((_this).f34).minus((_this).f30), (_this).f30, _module.__default.safeDivisionInt((_this).f30, (_this).f30));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f32, (_this).f32), (_this).f32);
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      let _source16 = ((_this.f28) ? (_module.D5.create_DC14(_this.f35, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f44,(_this).f44)).length))) : (_module.D5.create_DC14(_this.f29, (_this).f44)));
      if (_source16.is_DC13) {
        let _777___mcc_h0 = (_source16).cf27;
        let _778___mcc_h1 = (_source16).cf28;
        let _779___mcc_h2 = (_source16).cf29;
        let _780_cf29 = _779___mcc_h2;
        let _781_cf28 = _778___mcc_h1;
        let _782_cf27 = _777___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pcm"), _dafny.Seq.UnicodeFromString("cca"));
      } else if (_source16.is_DC14) {
        let _783___mcc_h3 = (_source16).cf30;
        let _784___mcc_h4 = (_source16).cf31;
        let _785_cf31 = _784___mcc_h4;
        let _786_cf30 = _783___mcc_h3;
        return _dafny.Seq.UnicodeFromString("auph");
      } else {
        let _787___mcc_h5 = (_source16).cf26;
        let _788_cf26 = _787___mcc_h5;
        return _788_cf26;
      }
    };
    fm2(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,new BigNumber(230))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_this).f44))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_this).f34));
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      (globalState).f26 = (new BigNumber(66)).isLessThanOrEqualTo((_this).f44);
      let _789_v0;
      let _nw117 = Array((new BigNumber(4)).toNumber());
      _nw117[(_dafny.ZERO).toNumber()] = _this.f36;
      _nw117[(_dafny.ONE).toNumber()] = _this.f36;
      _nw117[(new BigNumber(2)).toNumber()] = _this.f36;
      _nw117[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
      _789_v0 = _nw117;
      let _790_v1;
      _790_v1 = _dafny.Set.fromElements(_789_v0, _789_v0);
      if (((_790_v1).Difference(_790_v1)).IsProperSubsetOf(_790_v1)) {
        let _791_v2;
        _791_v2 = _dafny.Seq.UnicodeFromString("mmi");
        let _792_v3;
        _792_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,new BigNumber((_dafny.Seq.update(_791_v2, _module.__default.safeIndex((_this).f30, new BigNumber((_791_v2).length)), (_791_v2)[_module.__default.safeIndex((_this).f30, new BigNumber((_791_v2).length))])).length));
        let _793_v4;
        _793_v4 = _dafny.MultiSet.fromElements((_this).f44);
        let _794_v5;
        let _nw118 = new _module.C2();
        _nw118.__ctor(new BigNumber((_module.__default.fm31(_module.__default.fm32(!(_this.f28), globalState), _792_v3, (function (_pat_let9_0) {
          return function (_795_dt__update__tmp_h0) {
            return function (_pat_let10_0) {
              return function (_796_dt__update_hcf31_h0) {
                return _module.D5.create_DC14((_795_dt__update__tmp_h0).dtor_cf30, _796_dt__update_hcf31_h0);
              }(_pat_let10_0);
            }(new BigNumber(-18));
          }(_pat_let9_0);
        }(_module.D5.create_DC14(_this.f35, (_this).f44))).dtor_cf30, ((_this).f32)[_module.__default.safeIndex(new BigNumber((_793_v4).cardinality()), new BigNumber(((_this).f32).length))], globalState)).length), new BigNumber(454), ((_this).f34).isEqualTo((_this).f30), _this.f35, ((_this).f32)[_module.__default.safeIndex(new BigNumber(900), new BigNumber(((_this).f32).length))]);
        _794_v5 = _nw118;
        _794_v5 = ((_this.f31) ? (_794_v5) : (_794_v5));
        let _797_v6;
        _797_v6 = _dafny.Map.Empty.slice().updateUnsafe(_794_v5.f31,(_this).f44);
        let _798_v7;
        _798_v7 = _module.D5.create_DC13(_794_v5.f31, _797_v6, _this.f36);
        let _source17 = _798_v7;
        if (_source17.is_DC13) {
          let _799___mcc_h0 = (_source17).cf27;
          let _800___mcc_h1 = (_source17).cf28;
          let _801___mcc_h2 = (_source17).cf29;
          let _802_cf29 = _801___mcc_h2;
          let _803_cf28 = _800___mcc_h1;
          let _804_cf27 = _799___mcc_h0;
          (globalState).f13 = _module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.fromElements((_this).f34)).cardinality()), (_this).f34);
          let _805_v8;
          _805_v8 = _dafny.Seq.of((_this).f34, (_794_v5).f30);
          let _806_v9;
          let _nw119 = Array((new BigNumber(10)).toNumber());
          _nw119[(_dafny.ZERO).toNumber()] = _module.__default.fm32(_this.f28, globalState);
          _nw119[(_dafny.ONE).toNumber()] = (_805_v8)[_module.__default.safeIndex((_794_v5).f30, new BigNumber((_805_v8).length))];
          _nw119[(new BigNumber(2)).toNumber()] = ((_794_v5).f30).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(238)), ((_807_cf29) => function (_808_i0) {
            return _807_cf29;
          })(_802_cf29))).length));
          _nw119[(new BigNumber(3)).toNumber()] = new BigNumber((_805_v8).length);
          _nw119[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).f30);
          _nw119[(new BigNumber(5)).toNumber()] = (_794_v5).f30;
          _nw119[(new BigNumber(6)).toNumber()] = (_this).f34;
          _nw119[(new BigNumber(7)).toNumber()] = new BigNumber(-91);
          _nw119[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt((_794_v5).f30, (_794_v5).f30);
          _nw119[(new BigNumber(9)).toNumber()] = (_this).f30;
          _806_v9 = _nw119;
          let _nw120 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _806_v9 = _nw120;
          let _809_v10;
          let _nw121 = new _module.C0();
          _nw121.__ctor(_789_v0);
          _809_v10 = _nw121;
          let _810_v11;
          _810_v11 = _module.D1.create_DC3(_794_v5.f29, (_this).f34, _809_v10, (_this).f44);
          let _811_v12;
          _811_v12 = _module.D1.create_DC4(_810_v11);
          let _812_v13;
          _812_v13 = _dafny.MultiSet.fromElements(_811_v12, _811_v12);
          let _813_v14;
          _813_v14 = _dafny.Seq.of(_811_v12);
          let _pat_let_tv9 = _810_v11;
          let _814_v15;
          let _nw122 = Array((new BigNumber(14)).toNumber());
          _nw122[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_811_v12, function (_pat_let11_0) {
            return function (_815_dt__update__tmp_h1) {
              return function (_pat_let12_0) {
                return function (_816_dt__update_hcf7_h0) {
                  return _module.D1.create_DC4(_816_dt__update_hcf7_h0);
                }(_pat_let12_0);
              }(_pat_let_tv9);
            }(_pat_let11_0);
          }(_811_v12), _811_v12, _module.D1.create_DC4(_810_v11), _811_v12);
          _nw122[(_dafny.ONE).toNumber()] = (_812_v13).Union(_812_v13);
          _nw122[(new BigNumber(2)).toNumber()] = _812_v13;
          _nw122[(new BigNumber(3)).toNumber()] = _812_v13;
          _nw122[(new BigNumber(4)).toNumber()] = _812_v13;
          _nw122[(new BigNumber(5)).toNumber()] = (_812_v13).update(_module.D1.create_DC4(_810_v11), _module.__default.abs(new BigNumber((_791_v2).length)));
          _nw122[(new BigNumber(6)).toNumber()] = _812_v13;
          _nw122[(new BigNumber(7)).toNumber()] = _812_v13;
          _nw122[(new BigNumber(8)).toNumber()] = (_812_v13).Intersect(_812_v13);
          _nw122[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_module.D1.create_DC4(_810_v11));
          _nw122[(new BigNumber(10)).toNumber()] = _812_v13;
          _nw122[(new BigNumber(11)).toNumber()] = _812_v13;
          _nw122[(new BigNumber(12)).toNumber()] = ((_794_v5.f29) ? (_dafny.MultiSet.FromArray(_813_v14)) : (_812_v13));
          _nw122[(new BigNumber(13)).toNumber()] = _812_v13;
          _814_v15 = _nw122;
          let _index83 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_814_v15).length));
          (_814_v15)[_index83] = _dafny.MultiSet.fromElements(_811_v12);
          let _817_v16;
          _817_v16 = _dafny.Seq.of(_813_v14, _813_v14);
          let _index84 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_814_v15).length));
          (_814_v15)[_index84] = _dafny.MultiSet.FromArray(((_794_v5.f31) ? (_813_v14) : ((_817_v16)[_module.__default.safeIndex((_794_v5).f30, new BigNumber((_817_v16).length))])));
          (globalState).f12 = _794_v5.f29;
        } else if (_source17.is_DC14) {
          let _818___mcc_h3 = (_source17).cf30;
          let _819___mcc_h4 = (_source17).cf31;
          let _820_cf31 = _819___mcc_h4;
          let _821_cf30 = _818___mcc_h3;
          (_794_v5).f31 = !(_794_v5.f31);
          (globalState).f1 = !(_this.f28);
          let _822_v17;
          let _nw123 = Array((new BigNumber(28)).toNumber()).fill(false);
          _822_v17 = _nw123;
          let _index85 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_822_v17).length));
          (_822_v17)[_index85] = _794_v5.f31;
          let _823_v18;
          _823_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(242),false);
          let _824_v19;
          _824_v19 = _module.D2.create_DC6((_this).f34, _794_v5.f28, new BigNumber((_823_v18).length), _this.f31, (_794_v5).f30);
          let _825_v20;
          _825_v20 = _dafny.Seq.of(_824_v19, _824_v19);
          let _826_v21;
          _826_v21 = _module.D2.create_DC6(_module.__default.safeModuloInt(_module.__default.fm32(_this.f31, globalState), (_this).f30), !_dafny.Seq.contains(_825_v20, _824_v19), (_794_v5).f30, _794_v5.f29, (_this).f44);
          let _827_v22;
          _827_v22 = _module.D5.create_DC14(_794_v5.f31, (_this).f30);
          let _index86 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_822_v17).length));
          let _rhs82 = _this.f31;
          let _rhs83 = ((true) ? (_826_v21) : (_824_v19));
          let _rhs84 = (new BigNumber(227)).multipliedBy((_this).f34);
          let _rhs85 = _module.__default.fm1((_827_v22).dtor_cf30, ((_this).f30).plus((_this).f30), globalState);
          let _rhs86 = ((((_this.f35) ? (_this.f28) : (_794_v5.f31))) ? (_794_v5.f31) : (_dafny.areEqual(_dafny.Seq.of(_821_cf30), (_this).f32)));
          let _lhs63 = _822_v17;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_822_v17).length));
          let _lhs65 = globalState;
          let _lhs66 = _794_v5;
          _lhs63[_lhs64] = _rhs82;
          _824_v19 = _rhs83;
          _820_cf31 = _rhs84;
          _lhs65.f12 = _rhs85;
          _lhs66.f28 = _rhs86;
          (globalState).f13 = _module.__default.safeDivisionInt(((((_this).f32)[_module.__default.safeIndex(_820_cf31, new BigNumber(((_this).f32).length))]) ? (new BigNumber(-11)) : ((_794_v5).f30)), (_this).f30);
        } else {
          let _828___mcc_h5 = (_source17).cf26;
          let _829_cf26 = _828___mcc_h5;
          let _830_v23;
          _830_v23 = _dafny.Map.Empty.slice().updateUnsafe(_792_v3,_794_v5.f28);
          let _831_v24;
          let _832_v25;
          let _out49;
          let _out50;
          let _outcollector15 = (_this).m12((_this).f30, _830_v23, globalState);
          _out49 = _outcollector15[0];
          _out50 = _outcollector15[1];
          _831_v24 = _out49;
          _832_v25 = _out50;
          (globalState).f13 = ((_this).f34).minus((_794_v5).f30);
          let _833_v26;
          let _nw124 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _833_v26 = _nw124;
          let _index87 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_833_v26).length));
          (_833_v26)[_index87] = new BigNumber((_797_v6).length);
          let _index88 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_833_v26).length));
          let _rhs87 = _module.__default.fm32(_832_v25, globalState);
          let _rhs88 = _794_v5.f29;
          let _lhs67 = _833_v26;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_833_v26).length));
          let _lhs69 = _794_v5;
          _lhs67[_lhs68] = _rhs87;
          _lhs69.f31 = _rhs88;
          let _834_v27;
          let _nw125 = new _module.C1();
          _nw125.__ctor(_this.f31, _794_v5.f29, _794_v5.f28);
          _834_v27 = _nw125;
          let _835_v28;
          _835_v28 = _dafny.MultiSet.fromElements(_834_v27);
          (globalState).f1 = ((_835_v28).update(_834_v27, _module.__default.abs(new BigNumber(895)))).IsDisjointFrom((_835_v28).update(_834_v27, _module.__default.abs((_this).f34)));
        }
        (globalState).f1 = _794_v5.f31;
        let _836_v29;
        let _nw126 = new _module.C0();
        _nw126.__ctor(_789_v0);
        _836_v29 = _nw126;
        (_794_v5).f28 = true;
      } else {
        let _837_v30;
        _837_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_this.f28);
        let _838_v31;
        _838_v31 = _module.D0.create_DC1(_dafny.MultiSet.fromElements(!((((_837_v30).contains(false)) ? ((_837_v30).get(false)) : (!(_this.f28))))));
        _838_v31 = _838_v31;
        (globalState).f13 = (_this).f30;
        (_this).f31 = ((_this).f34).isLessThanOrEqualTo((_this).f30);
        let _839_v32;
        _839_v32 = _dafny.Set.fromElements(_this.f28, _this.f35);
        let _840_v33;
        _840_v33 = _dafny.Seq.of(new BigNumber((_839_v32).length));
        let _841_v34;
        _841_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-187),new BigNumber((_840_v33).length));
        let _842_v35;
        _842_v35 = _dafny.Seq.of((_841_v34).update((_this).f44, (_this).f34));
        let _rhs89 = _842_v35;
        let _rhs90 = _this.f28;
        let _rhs91 = _this.f29;
        let _lhs70 = globalState;
        let _lhs71 = globalState;
        _842_v35 = _rhs89;
        _lhs70.f26 = _rhs90;
        _lhs71.f1 = _rhs91;
        let _843_v36;
        let _nw127 = new _module.C0();
        _nw127.__ctor(_789_v0);
        _843_v36 = _nw127;
      }
      let _844_v37;
      _844_v37 = _dafny.MultiSet.fromElements(_this.f31, _this.f31);
      let _hi4 = ((_this).f44).minus((((_844_v37).contains(_this.f31)) ? ((_844_v37).get(_this.f31)) : ((_this).f44)));
      for (let _845_i1 = (_this).f30; _845_i1.isLessThan(_hi4); _845_i1 = _845_i1.plus(_dafny.ONE)) {
        (globalState).f12 = _this.f28;
        let _846_v38;
        _846_v38 = _module.D2.create_DC6(new BigNumber(51), _this.f28, (_this).f30, _this.f35, new BigNumber(290));
        let _847_v39;
        _847_v39 = _dafny.MultiSet.fromElements((_this).f34);
        let _848_v40;
        _848_v40 = _dafny.Seq.UnicodeFromString("ksmswsjps");
        let _849_v41;
        _849_v41 = _dafny.Set.fromElements(_848_v40, _848_v40, _848_v40, _848_v40, _848_v40);
        let _850_v42;
        _850_v42 = _dafny.Seq.of((_this).f34, (_this).f44);
        let _851_v43;
        _851_v43 = _dafny.Seq.of(new BigNumber((_850_v42).length));
        let _852_v44;
        _852_v44 = _dafny.Set.fromElements(new BigNumber((_850_v42).length));
        let _853_v45;
        _853_v45 = _dafny.Set.fromElements(_this.f28);
        let _854_v46;
        _854_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f44,_this.f28);
        let _855_v47;
        let _nw128 = Array((new BigNumber(25)).toNumber());
        _nw128[(_dafny.ZERO).toNumber()] = !((_846_v38).dtor_cf10);
        _nw128[(_dafny.ONE).toNumber()] = (_this).fm7(_847_v39, new BigNumber(125), globalState);
        _nw128[(new BigNumber(2)).toNumber()] = (_849_v41).IsSubsetOf(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("aa"), _848_v40));
        _nw128[(new BigNumber(3)).toNumber()] = _this.f29;
        _nw128[(new BigNumber(4)).toNumber()] = _this.f31;
        _nw128[(new BigNumber(5)).toNumber()] = _this.f29;
        _nw128[(new BigNumber(6)).toNumber()] = ((_this).f34).isLessThan(_845_i1);
        _nw128[(new BigNumber(7)).toNumber()] = _this.f31;
        _nw128[(new BigNumber(8)).toNumber()] = true;
        _nw128[(new BigNumber(9)).toNumber()] = false;
        _nw128[(new BigNumber(10)).toNumber()] = _this.f35;
        _nw128[(new BigNumber(11)).toNumber()] = _this.f35;
        _nw128[(new BigNumber(12)).toNumber()] = _this.f35;
        _nw128[(new BigNumber(13)).toNumber()] = _this.f35;
        _nw128[(new BigNumber(14)).toNumber()] = !(_this.f29);
        _nw128[(new BigNumber(15)).toNumber()] = _this.f28;
        _nw128[(new BigNumber(16)).toNumber()] = ((_this.f31) ? (false) : (_this.f31));
        _nw128[(new BigNumber(17)).toNumber()] = (_852_v44).equals(_852_v44);
        _nw128[(new BigNumber(18)).toNumber()] = !(_this.f31);
        _nw128[(new BigNumber(19)).toNumber()] = _this.f31;
        _nw128[(new BigNumber(20)).toNumber()] = (_853_v45).IsProperSubsetOf(_853_v45);
        _nw128[(new BigNumber(21)).toNumber()] = (_854_v46).contains(new BigNumber(135));
        _nw128[(new BigNumber(22)).toNumber()] = _this.f29;
        _nw128[(new BigNumber(23)).toNumber()] = !(_dafny.MultiSet.fromElements(new BigNumber((_848_v40).length), (_this).f34)).contains(new BigNumber((_850_v42).length));
        _nw128[(new BigNumber(24)).toNumber()] = false;
        _855_v47 = _nw128;
        let _index89 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_855_v47).length));
        (_855_v47)[_index89] = (_module.__default.fm32(_this.f29, globalState)).isEqualTo(_845_i1);
        let _index90 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_855_v47).length));
        (_855_v47)[_index90] = (_this).fm7(_847_v39, _module.__default.safeDivisionInt(new BigNumber((_853_v45).length), (_this).f44), globalState);
        let _856_v48;
        _856_v48 = _dafny.MultiSet.fromElements(_this.f36, new _dafny.CodePoint('o'.codePointAt(0)));
        let _857_v49;
        let _out51;
        _out51 = _module.__default.m0(_851_v43, !(_dafny.Seq.contains(_848_v40, _this.f36)), new BigNumber((_856_v48).cardinality()), !(_this.f28), globalState);
        _857_v49 = _out51;
        let _858_v50;
        _858_v50 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),_789_v0);
        let _859_v51;
        _859_v51 = _module.D0.create_DC1(_module.__default.fm33(_this.f29, globalState));
        _858_v50 = (_858_v50).update((_859_v51).dtor_cf1, _789_v0);
      }
      let _860_v52;
      _860_v52 = _dafny.Seq.of(_844_v37, _844_v37, _844_v37, _844_v37);
      let _861_v53;
      let _nw129 = new _module.C0();
      _nw129.__ctor(_789_v0);
      _861_v53 = _nw129;
      let _862_v54;
      _862_v54 = _dafny.Seq.UnicodeFromString("jrneeb");
      let _863_v55;
      _863_v55 = _module.D1.create_DC3(_this.f35, (_this).f34, _861_v53, new BigNumber((_862_v54).length));
      let _864_v56;
      _864_v56 = _dafny.MultiSet.fromElements(new BigNumber((_844_v37).cardinality()), (_this).f34);
      let _865_v57;
      let _nw130 = Array((new BigNumber(26)).toNumber());
      _nw130[(_dafny.ZERO).toNumber()] = (_this).f34;
      _nw130[(_dafny.ONE).toNumber()] = (_this).f30;
      _nw130[(new BigNumber(2)).toNumber()] = new BigNumber(((_860_v52)[_module.__default.safeIndex((_863_v55).dtor_cf4, new BigNumber((_860_v52).length))]).cardinality());
      _nw130[(new BigNumber(3)).toNumber()] = (_this).f34;
      _nw130[(new BigNumber(4)).toNumber()] = ((_this).f44).plus(new BigNumber((_862_v54).length));
      _nw130[(new BigNumber(5)).toNumber()] = ((_this).f44).multipliedBy((_this).f34);
      _nw130[(new BigNumber(6)).toNumber()] = (_this).f30;
      _nw130[(new BigNumber(7)).toNumber()] = new BigNumber(-578);
      _nw130[(new BigNumber(8)).toNumber()] = (_this).f44;
      _nw130[(new BigNumber(9)).toNumber()] = (_this).f30;
      _nw130[(new BigNumber(10)).toNumber()] = (_this).f34;
      _nw130[(new BigNumber(11)).toNumber()] = (_this).f30;
      _nw130[(new BigNumber(12)).toNumber()] = (_this).f44;
      _nw130[(new BigNumber(13)).toNumber()] = (_this).f34;
      _nw130[(new BigNumber(14)).toNumber()] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f30,new BigNumber(337))).length)).multipliedBy((_this).f30);
      _nw130[(new BigNumber(15)).toNumber()] = ((_this).f30).plus((_this).f44);
      _nw130[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus((_this).f34);
      _nw130[(new BigNumber(17)).toNumber()] = (_this).f34;
      _nw130[(new BigNumber(18)).toNumber()] = ((_this).f34).minus((_this).f30);
      _nw130[(new BigNumber(19)).toNumber()] = (_this).f34;
      _nw130[(new BigNumber(20)).toNumber()] = ((_this.f28) ? ((_dafny.ZERO).minus((_this).f34)) : ((_this).f44));
      _nw130[(new BigNumber(21)).toNumber()] = (((_864_v56).contains((_dafny.ZERO).minus((_this).f44))) ? ((_864_v56).get((_dafny.ZERO).minus((_this).f44))) : (new BigNumber((_dafny.Seq.of(_this.f36)).length)));
      _nw130[(new BigNumber(22)).toNumber()] = _module.__default.safeDivisionInt((_this).f44, new BigNumber((_862_v54).length));
      _nw130[(new BigNumber(23)).toNumber()] = (_this).f44;
      _nw130[(new BigNumber(24)).toNumber()] = (_this).f44;
      _nw130[(new BigNumber(25)).toNumber()] = (((_844_v37).contains(!(_this.f35))) ? ((_844_v37).get(!(_this.f35))) : ((_this).f44));
      _865_v57 = _nw130;
      _865_v57 = _865_v57;
      if ((_this).fm6(globalState)) {
        if (_this.f35) {
          let _866_v58;
          let _nw131 = Array((new BigNumber(23)).toNumber()).fill(false);
          _866_v58 = _nw131;
          let _rhs92 = _this.f31;
          let _rhs93 = _866_v58;
          let _rhs94 = _865_v57;
          let _rhs95 = _this.f28;
          let _rhs96 = false;
          let _lhs72 = globalState;
          let _lhs73 = globalState;
          let _lhs74 = _this;
          r0 = _rhs92;
          _lhs72.f21 = _rhs93;
          _865_v57 = _rhs94;
          _lhs73.f12 = _rhs95;
          _lhs74.f28 = _rhs96;
          let _867_v59;
          _867_v59 = _dafny.Seq.of(_865_v57, _865_v57);
          let _rhs97 = _dafny.Seq.UnicodeFromString("d");
          let _rhs98 = _866_v58;
          let _rhs99 = _dafny.Seq.update((_this).fm3(!(!_dafny.Seq.contains(_867_v59, _865_v57)), _this.f29, globalState), _module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_864_v56).cardinality()), (_this).f30), new BigNumber(((_this).fm3(!(!_dafny.Seq.contains(_867_v59, _865_v57)), _this.f29, globalState)).length)), (_862_v54)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_860_v52).length)), new BigNumber((_862_v54).length))]);
          let _lhs75 = globalState;
          _862_v54 = _rhs97;
          _lhs75.f21 = _rhs98;
          _862_v54 = _rhs99;
          let _index91 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_865_v57).length));
          (_865_v57)[_index91] = (_this).f30;
          let _index92 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_865_v57).length));
          (_865_v57)[_index92] = ((_this).f34).minus((_this).f30);
          let _868_v60;
          _868_v60 = _dafny.Seq.of((_dafny.ZERO).minus((_865_v57)[_module.__default.safeIndex(new BigNumber(616), new BigNumber((_865_v57).length))]));
          (globalState).f12 = ((_this).f34).isLessThan((_dafny.ZERO).minus((new BigNumber((_868_v60).length)).multipliedBy((_865_v57)[_module.__default.safeIndex(new BigNumber(616), new BigNumber((_865_v57).length))])));
          let _index93 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_866_v58).length));
          (_866_v58)[_index93] = !(false);
          let _index94 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_866_v58).length));
          (_866_v58)[_index94] = ((_this).f44).isLessThan((_this).f30);
        } else {
          r2 = _862_v54;
          let _869_v61;
          _869_v61 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_this.f36);
          let _870_v62;
          _870_v62 = _dafny.Map.Empty.slice().updateUnsafe((((_869_v61).contains(_this.f35)) ? ((_869_v61).get(_this.f35)) : (_this.f36)),new BigNumber((_862_v54).length));
          let _index95 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_865_v57).length));
          (_865_v57)[_index95] = _module.__default.safeDivisionInt(new BigNumber((_870_v62).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),(_this).f34)).length));
          let _871_v63;
          _871_v63 = _module.D2.create_DC6((_this).f34, _this.f31, (_this).f34, _this.f29, (_this).f30);
          let _872_v64;
          _872_v64 = _dafny.MultiSet.fromElements(_this.f36, _this.f36, _this.f36, _this.f36);
          let _873_v65;
          _873_v65 = _module.D3.create_DC8(_872_v64, _this.f28, (_this).f44);
          let _874_v66;
          let _nw132 = Array((new BigNumber(11)).toNumber());
          _nw132[(_dafny.ZERO).toNumber()] = (_871_v63).dtor_cf10;
          _nw132[(_dafny.ONE).toNumber()] = _this.f31;
          _nw132[(new BigNumber(2)).toNumber()] = !(_844_v37).equals(_dafny.MultiSet.fromElements(_this.f31));
          _nw132[(new BigNumber(3)).toNumber()] = false;
          _nw132[(new BigNumber(4)).toNumber()] = !((_this).f34).isEqualTo(new BigNumber(((_873_v65).dtor_cf15).cardinality()));
          _nw132[(new BigNumber(5)).toNumber()] = _this.f35;
          _nw132[(new BigNumber(6)).toNumber()] = _this.f29;
          _nw132[(new BigNumber(7)).toNumber()] = _this.f35;
          _nw132[(new BigNumber(8)).toNumber()] = _this.f29;
          _nw132[(new BigNumber(9)).toNumber()] = (new BigNumber((_862_v54).length)).isLessThanOrEqualTo((_this).f30);
          _nw132[(new BigNumber(10)).toNumber()] = _this.f28;
          _874_v66 = _nw132;
          let _index96 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_874_v66).length));
          (_874_v66)[_index96] = (_this.f35) || (_this.f29);
          let _875_v67;
          _875_v67 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), function (_876_i2) {
            return _this.f36;
          })).length));
          let _index97 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_865_v57).length));
          let _index98 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_874_v66).length));
          let _rhs100 = (_this).f44;
          let _rhs101 = (function () {
            let _coll23 = new _dafny.Set();
            for (const _compr_23 of (_875_v67).Elements) {
              let _877_v68 = _compr_23;
              if ((_875_v67).contains(_877_v68)) {
                _coll23.add(_module.__default.safeModuloInt(_877_v68, new BigNumber((_dafny.Seq.UnicodeFromString("aohlb")).length)));
              }
            }
            return _coll23;
          }()).IsSubsetOf(_dafny.Set.fromElements((_this).f30));
          let _lhs76 = _865_v57;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_865_v57).length));
          let _lhs78 = _874_v66;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_874_v66).length));
          _lhs76[_lhs77] = _rhs100;
          _lhs78[_lhs79] = _rhs101;
          (globalState).f13 = ((_871_v63).dtor_cf11).multipliedBy((_865_v57)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_865_v57).length))]);
          let _878_v69;
          _878_v69 = _dafny.Map.Empty.slice().updateUnsafe((_874_v66)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_874_v66).length))],(_865_v57)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_865_v57).length))]);
          let _879_v70;
          _879_v70 = _dafny.MultiSet.fromElements(_878_v69, _878_v69, ((_this.f35) ? (_878_v69) : (_878_v69)));
          _879_v70 = _879_v70;
          (globalState).f12 = (_874_v66)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_874_v66).length))];
        }
        let _880_v71;
        let _init16 = function (_881_i3) {
          return ((_this.f29) ? (_this.f28) : (_this.f35));
        };
        let _nw133 = Array((_dafny.ONE).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw133.length); _i0_16++) {
          _nw133[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _880_v71 = _nw133;
        let _index99 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_880_v71).length));
        (_880_v71)[_index99] = _this.f29;
        let _index100 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_880_v71).length));
        (_880_v71)[_index100] = !_dafny.Seq.contains((_this).f32, !(_this.f28) || (_this.f28));
        (globalState).f13 = _module.__default.fm32(true, globalState);
        let _882_v72;
        _882_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_this.f31);
        (_this).f36 = _module.__default.fm34(new BigNumber(((_882_v72).update((_this).f44, _this.f28)).length), (_this).f34, (_this).f44, _this.f31, globalState);
        let _index101 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_865_v57).length));
        (_865_v57)[_index101] = (_this).f30;
        let _index102 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_865_v57).length));
        (_865_v57)[_index102] = _module.__default.safeModuloInt(_module.__default.fm32((_880_v71)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((_880_v71).length))], globalState), (_this).f30);
      } else {
        (globalState).f13 = (_module.__default.fm32(_this.f28, globalState)).multipliedBy((_this).f30);
        r0 = _this.f28;
        let _883_v73;
        _883_v73 = _dafny.Seq.of((_this).f34, new BigNumber(307), _module.__default.safeDivisionInt((_this).f34, new BigNumber(381)), (_this).f44, (_this).f30);
        let _884_v74;
        _884_v74 = _dafny.MultiSet.fromElements(_883_v73);
        let _rhs102 = _dafny.Seq.of(((_this.f31) ? (new BigNumber((_884_v74).cardinality())) : ((_this).f44)), (_dafny.ZERO).minus((_this).f30), (_this).f44, (_this).f34, _module.__default.safeModuloInt((_this).f34, new BigNumber(-445)));
        let _rhs103 = !(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_883_v73, _883_v73))).contains((_this).f30);
        let _rhs104 = !(((_this).f44).isLessThan(((_this).f30).plus((_this).f34)));
        let _lhs80 = _this;
        let _lhs81 = globalState;
        _883_v73 = _rhs102;
        _lhs80.f29 = _rhs103;
        _lhs81.f26 = _rhs104;
        let _885_v75;
        let _nw134 = Array((new BigNumber(10)).toNumber());
        _nw134[(_dafny.ZERO).toNumber()] = true;
        _nw134[(_dafny.ONE).toNumber()] = _this.f28;
        _nw134[(new BigNumber(2)).toNumber()] = (_this).fm7(_dafny.MultiSet.FromArray(_883_v73), (_this).f30, globalState);
        _nw134[(new BigNumber(3)).toNumber()] = !(_this.f29);
        _nw134[(new BigNumber(4)).toNumber()] = true;
        _nw134[(new BigNumber(5)).toNumber()] = !(_this.f29);
        _nw134[(new BigNumber(6)).toNumber()] = _this.f28;
        _nw134[(new BigNumber(7)).toNumber()] = _this.f31;
        _nw134[(new BigNumber(8)).toNumber()] = _this.f28;
        _nw134[(new BigNumber(9)).toNumber()] = _this.f35;
        _885_v75 = _nw134;
        let _886_v76;
        _886_v76 = _dafny.Map.Empty.slice().updateUnsafe((_this).f44,_885_v75);
        let _887_v77;
        _887_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f44,_862_v54);
        _886_v76 = (_886_v76).update(((_this).f34).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.update((_this).f32, _module.__default.safeIndex(new BigNumber((_887_v77).length), new BigNumber(((_this).f32).length)), _this.f29), _module.__default.safeIndex((_this).f30, new BigNumber((_dafny.Seq.update((_this).f32, _module.__default.safeIndex(new BigNumber((_887_v77).length), new BigNumber(((_this).f32).length)), _this.f29)).length)), _this.f31)).length)), _885_v75);
        (globalState).f13 = (_this).f30;
      }
      r2 = _862_v54;
      let _888_v78;
      _888_v78 = _dafny.Map.Empty.slice().updateUnsafe(!(_864_v56).equals(_864_v56),_this.f28);
      let _889_v79;
      _889_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f30);
      r0 = (((_888_v78).contains((_889_v79).contains((_this).f30))) ? ((_888_v78).get((_889_v79).contains((_this).f30))) : (_this.f28));
      let _890_v80;
      _890_v80 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,!(_this.f31));
      let _891_v81;
      _891_v81 = _dafny.Seq.of(new BigNumber((_890_v80).length));
      let _892_v82;
      _892_v82 = _module.D6.create_DC15(_891_v81);
      let _893_v83;
      _893_v83 = _dafny.Map.Empty.slice().updateUnsafe(_892_v82,(_this).f34);
      let _894_v84;
      _894_v84 = _dafny.Seq.of((_this).f44, (((_893_v83).contains(_892_v82)) ? ((_893_v83).get(_892_v82)) : (new BigNumber(601))));
      r1 = (_894_v84)[_module.__default.safeIndex((_this).f34, new BigNumber((_894_v84).length))];
      r2 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(926)), function (_895_i4) {
        return _this.f36;
      }), _module.__default.safeIndex((_this).f44, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(926)), function (_896_i4) {
        return _this.f36;
      })).length)), _this.f36), _dafny.Seq.update(_dafny.Seq.Concat(_862_v54, _862_v54), _module.__default.safeIndex((_this).f44, new BigNumber((_dafny.Seq.Concat(_862_v54, _862_v54)).length)), _this.f36));
      return [r0, r1, r2];
    }
    m12(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let _897_v0;
      let _nw135 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _897_v0 = _nw135;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_897_v0).length))) {
        let _898_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_898_i0)) && ((_898_i0).isLessThan(new BigNumber((_897_v0).length))))) {
          (_897_v0)[(_898_i0)] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pftnlhv"), _dafny.Seq.UnicodeFromString("mekvfibbt"));
        }
      }
      let _899_v1;
      _899_v1 = _dafny.Seq.of((_this).f34);
      let _900_v2;
      let _out52;
      _out52 = _module.__default.m0(_dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus((_this).f30), (_this).f44), _899_v1), _this.f35, (_this).f44, _this.f31, globalState);
      _900_v2 = _out52;
      let _901_v3;
      _901_v3 = _module.D3.create_DC8(_dafny.MultiSet.fromElements(_this.f36), _this.f28, p0);
      if ((_this.f29) || ((_901_v3).dtor_cf16)) {
        let _902_v4;
        let _init17 = function (_903_i1) {
          return _this.f36;
        };
        let _nw136 = Array((new BigNumber(16)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw136.length); _i0_17++) {
          _nw136[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _902_v4 = _nw136;
        let _904_v5;
        let _nw137 = new _module.C0();
        _nw137.__ctor(_902_v4);
        _904_v5 = _nw137;
        let _905_v6;
        let _nw138 = new _module.C0();
        _nw138.__ctor(_904_v5.f37);
        _905_v6 = _nw138;
        let _906_v7;
        _906_v7 = _dafny.Seq.UnicodeFromString("eaganir");
        (globalState).f13 = ((_dafny.ZERO).minus(new BigNumber((_906_v7).length))).plus((_this).f34);
        let _907_v8;
        _907_v8 = _dafny.Set.fromElements(_902_v4);
        let _908_v9;
        let _out53;
        _out53 = _module.__default.m0(_dafny.Seq.of((_this).f44), !((_907_v8).IsProperSubsetOf(_907_v8)), p0, _this.f28, globalState);
        _908_v9 = _out53;
        let _909_v10;
        _909_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f35);
        let _910_v11;
        let _nw139 = new _module.C2();
        _nw139.__ctor((_this).f30, (_dafny.ZERO).minus(new BigNumber((_909_v10).length)), _this.f35, (_this.f31) || (_this.f29), _this.f28);
        _910_v11 = _nw139;
      } else {
        let _911_v12;
        let _nw140 = Array((new BigNumber(11)).toNumber()).fill([]);
        _911_v12 = _nw140;
        let _912_v13;
        let _init18 = function (_913_i2) {
          return _this.f29;
        };
        let _nw141 = Array((new BigNumber(17)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw141.length); _i0_18++) {
          _nw141[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _912_v13 = _nw141;
        let _nw142 = Array((new BigNumber(13)).toNumber());
        _nw142[(_dafny.ZERO).toNumber()] = _912_v13;
        _nw142[(_dafny.ONE).toNumber()] = _912_v13;
        _nw142[(new BigNumber(2)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(3)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(4)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(5)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(6)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(7)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(8)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(9)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(10)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(11)).toNumber()] = _912_v13;
        _nw142[(new BigNumber(12)).toNumber()] = _912_v13;
        _911_v12 = _nw142;
        let _914_v14;
        _914_v14 = _dafny.Seq.of(_this.f28, _this.f35);
        let _915_v15;
        _915_v15 = _dafny.MultiSet.fromElements((_this).f34);
        let _916_v16;
        let _nw143 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _916_v16 = _nw143;
        let _917_v17;
        _917_v17 = _dafny.Map.Empty.slice().updateUnsafe(_916_v16,_912_v13);
        let _918_v18;
        _918_v18 = _dafny.Seq.UnicodeFromString("l");
        let _919_v19;
        _919_v19 = _dafny.MultiSet.fromElements(_918_v18);
        let _rhs105 = new BigNumber(((_dafny.MultiSet.fromElements(_918_v18)).Intersect((_919_v19).update(_918_v18, _module.__default.abs((_this).f34)))).cardinality());
        let _rhs106 = _this.f28;
        let _rhs107 = (_this).f32;
        let _rhs108 = (_915_v15).update(_900_v2, _module.__default.abs(new BigNumber((_dafny.Seq.of(p0, (_this).f30, (_899_v1)[_module.__default.safeIndex(p0, new BigNumber((_899_v1).length))])).length)));
        let _rhs109 = _917_v17;
        let _lhs82 = globalState;
        _900_v2 = _rhs105;
        _lhs82.f1 = _rhs106;
        _914_v14 = _rhs107;
        _915_v15 = _rhs108;
        _917_v17 = _rhs109;
        let _index103 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_912_v13).length));
        (_912_v13)[_index103] = !((((_this).f32)[_module.__default.safeIndex((_this).f44, new BigNumber(((_this).f32).length))]) === (_this.f29));
        let _index104 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_912_v13).length));
        let _rhs110 = _this.f35;
        let _rhs111 = _dafny.Seq.Concat(_914_v14, _dafny.Seq.Concat((_this).f32, (_this).f32));
        let _rhs112 = _900_v2;
        let _rhs113 = (_918_v18)[_module.__default.safeIndex(_900_v2, new BigNumber((_918_v18).length))];
        let _rhs114 = _this.f31;
        let _lhs83 = globalState;
        let _lhs84 = _this;
        let _lhs85 = _912_v13;
        let _lhs86 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_912_v13).length));
        _lhs83.f26 = _rhs110;
        _914_v14 = _rhs111;
        _900_v2 = _rhs112;
        _lhs84.f36 = _rhs113;
        _lhs85[_lhs86] = _rhs114;
        let _920_v20;
        let _nw144 = new _module.C1();
        _nw144.__ctor(_this.f35, ((_this).f32)[_module.__default.safeIndex(new BigNumber(729), new BigNumber(((_this).f32).length))], _this.f31);
        _920_v20 = _nw144;
        (globalState).f1 = _this.f28;
      }
      let _921_i3;
      _921_i3 = _dafny.ZERO;
      L3: {
        while (_this.f29) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_921_i3)) {
              break L3;
            }
            _921_i3 = (_921_i3).plus(_dafny.ONE);
            let _922_v21;
            _922_v21 = _dafny.Seq.of(_this.f28);
            let _rhs115 = ((_this.f28) ? ((_this).f32) : (_922_v21));
            let _rhs116 = _900_v2;
            let _rhs117 = !(_this.f31) || (_this.f28);
            let _lhs87 = globalState;
            _922_v21 = _rhs115;
            _900_v2 = _rhs116;
            _lhs87.f12 = _rhs117;
            let _923_v22;
            let _init19 = ((_924_p0) => function (_925_i4) {
              return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f35,_this.f28))).cardinality()))).IsSubsetOf(_dafny.MultiSet.fromElements(_924_p0));
            })(p0);
            let _nw145 = Array((new BigNumber(18)).toNumber());
            for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw145.length); _i0_19++) {
              _nw145[_i0_19] = _init19(new BigNumber(_i0_19));
            }
            _923_v22 = _nw145;
            let _index105 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_923_v22).length));
            (_923_v22)[_index105] = _this.f29;
            let _index106 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_923_v22).length));
            (_923_v22)[_index106] = _this.f28;
            let _926_v23;
            _926_v23 = _dafny.MultiSet.fromElements((_this).f34, (_this).f34, _module.__default.fm32(!((_923_v22)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_923_v22).length))]), globalState), (_this).f44);
            r1 = (_this).fm7(_926_v23, (_this).f44, globalState);
            let _927_v24;
            let _nw146 = Array((new BigNumber(4)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _927_v24 = _nw146;
            let _index107 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_927_v24).length));
            (_927_v24)[_index107] = _this.f36;
            let _928_v25;
            _928_v25 = _module.D5.create_DC13(_this.f29, _dafny.Map.Empty.slice().updateUnsafe((_923_v22)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_923_v22).length))],(_this).f30), new _dafny.CodePoint('y'.codePointAt(0)));
            let _index108 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_927_v24).length));
            (_927_v24)[_index108] = (_928_v25).dtor_cf29;
          }
        }
      }
      let _929_v26;
      let _nw147 = new _module.C2();
      _nw147.__ctor((_this).f34, (_dafny.ZERO).minus(_900_v2), !(_this.f35), !(_this.f31), _this.f31);
      _929_v26 = _nw147;
      let _930_v27;
      _930_v27 = _dafny.Seq.of(_this.f36, _this.f36);
      let _931_v28;
      _931_v28 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,_dafny.Seq.IsPrefixOf(_930_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(377)), function (_932_i5) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })));
      let _933_v29;
      let _nw148 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
      _933_v29 = _nw148;
      let _934_v30;
      _934_v30 = _dafny.MultiSet.fromElements(_this.f36);
      let _935_v31;
      _935_v31 = _dafny.Seq.of(_933_v29);
      let _rhs118 = (((_934_v30).contains(_this.f36)) ? ((_934_v30).get(_this.f36)) : ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f35,_this.f35)).length)).multipliedBy(p0)));
      let _rhs119 = _931_v28;
      let _rhs120 = (_935_v31)[_module.__default.safeIndex(new BigNumber(((_this).f32).length), new BigNumber((_935_v31).length))];
      let _lhs88 = globalState;
      _lhs88.f13 = _rhs118;
      _931_v28 = _rhs119;
      _933_v29 = _rhs120;
      let _936_v32;
      _936_v32 = _dafny.Set.fromElements(_900_v2);
      r0 = (_936_v32).Union((_936_v32).Union(_936_v32));
      r1 = _this.f28;
      return [r0, r1];
    }
    get f44() {
      let _this = this;
      return _this._f44;
    };
    get f45() {
      let _this = this;
      return _this._f45;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f28 = false;
      this._f29 = false;
      this._f31 = false;
      this._f35 = false;
      this._f34 = _dafny.ZERO;
      this._f32 = _dafny.Seq.of();
      this._f33 = _dafny.Seq.of();
      this._f30 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    set f31(value) {
      let _this = this;
      _this._f31 = value;
    };
    get f35() {
      let _this = this;
      return _this._f35;
    };
    set f35(value) {
      let _this = this;
      _this._f35 = value;
    };
    get f34() {
      let _this = this;
      return _this._f34;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    __ctor(f34, f35, f32, f33, f30, f31, f28, f29) {
      let _this = this;
      (_this)._f34 = f34;
      (_this)._f35 = f35;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm6(globalState) {
      let _this = this;
      return !(_this.f31);
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((function () {
        let _coll24 = new _dafny.Set();
        for (const _compr_24 of _dafny.IntegerRange(new BigNumber(460), new BigNumber(658))) {
          let _937_v0 = _compr_24;
          if (((new BigNumber(460)).isLessThanOrEqualTo(_937_v0)) && ((_937_v0).isLessThan(new BigNumber(658)))) {
            _coll24.add(_module.__default.safeModuloInt(_937_v0, (_this).f34));
          }
        }
        return _coll24;
      }()).Difference(_dafny.Set.fromElements(new BigNumber(972)))).Union((_dafny.Set.fromElements((_this).f34)).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_this.f28)).length))));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f32;
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      if ((_dafny.MultiSet.fromElements((_this).f30)).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(845)), function (_938_i0) {
        return (_this).f30;
      })))) {
        return _dafny.Seq.UnicodeFromString("fe");
      } else {
        return _dafny.Seq.UnicodeFromString("sjcbfbtrq");
      }
    };
    fm2(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(!(!(((_this).f32)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length))])))).IsSubsetOf(_dafny.Set.fromElements(_this.f31)),(_this).f34);
    };
    fm44(p0, globalState) {
      let _this = this;
      return ((_this).f30).isLessThan(_module.__default.safeModuloInt((_this).f30, (_dafny.ZERO).minus((_this).f30)));
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      (_this).f35 = _this.f31;
      let _939_i0;
      _939_i0 = _dafny.ZERO;
      L4: {
        while ((_this).fm6(globalState)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_939_i0)) {
              break L4;
            }
            _939_i0 = (_939_i0).plus(_dafny.ONE);
            let _940_v0;
            _940_v0 = _dafny.MultiSet.fromElements(_this.f28);
            let _941_v1;
            _941_v1 = _dafny.Set.fromElements(_this.f31);
            let _942_v2;
            let _nw149 = Array((new BigNumber(20)).toNumber());
            _nw149[(_dafny.ZERO).toNumber()] = _this.f31;
            _nw149[(_dafny.ONE).toNumber()] = _this.f28;
            _nw149[(new BigNumber(2)).toNumber()] = _this.f28;
            _nw149[(new BigNumber(3)).toNumber()] = _this.f29;
            _nw149[(new BigNumber(4)).toNumber()] = _this.f28;
            _nw149[(new BigNumber(5)).toNumber()] = _this.f29;
            _nw149[(new BigNumber(6)).toNumber()] = _this.f28;
            _nw149[(new BigNumber(7)).toNumber()] = _this.f28;
            _nw149[(new BigNumber(8)).toNumber()] = _this.f28;
            _nw149[(new BigNumber(9)).toNumber()] = false;
            _nw149[(new BigNumber(10)).toNumber()] = _this.f29;
            _nw149[(new BigNumber(11)).toNumber()] = _this.f29;
            _nw149[(new BigNumber(12)).toNumber()] = _this.f35;
            _nw149[(new BigNumber(13)).toNumber()] = _this.f29;
            _nw149[(new BigNumber(14)).toNumber()] = _this.f29;
            _nw149[(new BigNumber(15)).toNumber()] = true;
            _nw149[(new BigNumber(16)).toNumber()] = true;
            _nw149[(new BigNumber(17)).toNumber()] = (_this).fm6(globalState);
            _nw149[(new BigNumber(18)).toNumber()] = _this.f35;
            _nw149[(new BigNumber(19)).toNumber()] = (_this).fm6(globalState);
            _942_v2 = _nw149;
            let _943_v3;
            _943_v3 = _dafny.Map.Empty.slice().updateUnsafe(_942_v2,_941_v1);
            let _944_v4;
            _944_v4 = _dafny.Seq.of(_941_v1);
            let _945_v5;
            let _nw150 = Array((new BigNumber(23)).toNumber());
            _nw150[(_dafny.ZERO).toNumber()] = _module.__default.fm45((_this).f30, _940_v0, _this.f35, new BigNumber(((_this).f32).length), globalState);
            _nw150[(_dafny.ONE).toNumber()] = ((_941_v1)).Difference(_dafny.Set.fromElements(_this.f29, false, _this.f29, _this.f35, _this.f31));
            _nw150[(new BigNumber(2)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(3)).toNumber()] = (((_943_v3).contains(_942_v2)) ? ((_943_v3).get(_942_v2)) : (_941_v1));
            _nw150[(new BigNumber(4)).toNumber()] = (_941_v1).Difference(_941_v1);
            _nw150[(new BigNumber(5)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(6)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(7)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(8)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(9)).toNumber()] = (_dafny.Set.fromElements(_this.f35)).Difference(_module.__default.fm45((_dafny.ZERO).minus((_this).f30), _940_v0, _this.f35, (_this).f34, globalState));
            _nw150[(new BigNumber(10)).toNumber()] = (_941_v1).Intersect(_941_v1);
            _nw150[(new BigNumber(11)).toNumber()] = (_944_v4)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_944_v4).length))];
            _nw150[(new BigNumber(12)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(13)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(14)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(15)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(16)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(17)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(18)).toNumber()] = _module.__default.fm45((_this).f30, _940_v0, _this.f35, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).fm6(globalState),(_this).f34)).length), globalState);
            _nw150[(new BigNumber(19)).toNumber()] = _dafny.Set.fromElements(true);
            _nw150[(new BigNumber(20)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(21)).toNumber()] = _941_v1;
            _nw150[(new BigNumber(22)).toNumber()] = _941_v1;
            _945_v5 = _nw150;
            _945_v5 = _945_v5;
            let _946_v6;
            _946_v6 = _dafny.MultiSet.fromElements((_this).f30, new BigNumber(355));
            if (((_this).f32)[_module.__default.safeIndex(_module.__default.fm46((_this).f34, new BigNumber((_946_v6).cardinality()), _this.f28, globalState), new BigNumber(((_this).f32).length))]) {
              let _947_v7;
              _947_v7 = _dafny.Set.fromElements((_this).f34, (_this).f34);
              r1 = _module.__default.safeDivisionInt(new BigNumber(((_947_v7).Difference(_947_v7)).length), new BigNumber((_946_v6).cardinality()));
              let _948_v8;
              _948_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_this.f29);
              let _949_v9;
              let _nw151 = Array((new BigNumber(12)).toNumber());
              _nw151[(_dafny.ZERO).toNumber()] = new BigNumber(((_this).f32).length);
              _nw151[(_dafny.ONE).toNumber()] = (_this).f30;
              _nw151[(new BigNumber(2)).toNumber()] = (_this).f30;
              _nw151[(new BigNumber(3)).toNumber()] = ((_this).f34).multipliedBy((_this).f30);
              _nw151[(new BigNumber(4)).toNumber()] = new BigNumber((_948_v8).length);
              _nw151[(new BigNumber(5)).toNumber()] = (_this).f30;
              _nw151[(new BigNumber(6)).toNumber()] = (_this).f30;
              _nw151[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((_this).f30, (_this).f30);
              _nw151[(new BigNumber(8)).toNumber()] = (_this).f34;
              _nw151[(new BigNumber(9)).toNumber()] = ((_this.f29) ? ((_this).f34) : (new BigNumber((_dafny.Seq.UnicodeFromString("khil")).length)));
              _nw151[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f34, (_this).f34));
              _nw151[(new BigNumber(11)).toNumber()] = (_this).f34;
              _949_v9 = _nw151;
              let _index109 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_949_v9).length));
              (_949_v9)[_index109] = ((_this).f30).multipliedBy((_dafny.ZERO).minus((_this).f30));
              let _index110 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_949_v9).length));
              (_949_v9)[_index110] = new BigNumber(802);
              (globalState).f13 = new BigNumber(((_947_v7).Union(_947_v7)).length);
              let _index111 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_949_v9).length));
              (_949_v9)[_index111] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_949_v9)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_949_v9).length))]));
              (globalState).f1 = ((_this).f30).isLessThan((_949_v9)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_949_v9).length))]);
            } else {
              let _index112 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_942_v2).length));
              (_942_v2)[_index112] = ((_this).f30).isEqualTo((_this).f30);
              let _index113 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_942_v2).length));
              (_942_v2)[_index113] = _this.f29;
              (globalState).f13 = new BigNumber(((_946_v6).Difference(_946_v6)).cardinality());
              (globalState).f13 = (_this).f30;
              r1 = new BigNumber(((_this).f32).length);
              let _index114 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_942_v2).length));
              (_942_v2)[_index114] = _this.f28;
            }
            let _950_v10;
            _950_v10 = new _dafny.CodePoint('p'.codePointAt(0));
            _950_v10 = _950_v10;
            let _951_v11;
            let _nw152 = Array((new BigNumber(18)).toNumber());
            _nw152[(_dafny.ZERO).toNumber()] = _942_v2;
            _nw152[(_dafny.ONE).toNumber()] = _942_v2;
            _nw152[(new BigNumber(2)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(3)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(4)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(5)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(6)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(7)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(8)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(9)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(10)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(11)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(12)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(13)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(14)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(15)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(16)).toNumber()] = _942_v2;
            _nw152[(new BigNumber(17)).toNumber()] = _942_v2;
            _951_v11 = _nw152;
            let _index115 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_951_v11).length));
            (_951_v11)[_index115] = _942_v2;
            let _952_v12;
            _952_v12 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat((_this).f32, (_this).f32), _dafny.Seq.Concat((_this).f32, (_this).f32), (_this).f32, _dafny.Seq.Concat((_this).f32, (_this).fm5((_this).f30, (_this).f34, !(_this.f35), (_this).f30, globalState)));
            let _953_v13;
            _953_v13 = _dafny.Set.fromElements((_this).f34);
            let _954_v14;
            _954_v14 = _953_v13;
            let _955_v15;
            _955_v15 = _dafny.Seq.of((_954_v14), _953_v13, _953_v13);
            let _956_v16;
            _956_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,(_953_v13).IsSubsetOf((_955_v15)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_955_v15).length))]));
            let _index116 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_951_v11).length));
            let _rhs121 = _942_v2;
            let _rhs122 = (_952_v12).Difference((_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of((_this).f32, (_this).f32), _module.__default.safeIndex((_this).f30, new BigNumber((_dafny.Seq.of((_this).f32, (_this).f32)).length)), (_this).f32))).Union(_952_v12));
            let _rhs123 = _953_v13;
            let _rhs124 = (_956_v16).Merge(_956_v16);
            let _lhs89 = _951_v11;
            let _lhs90 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_951_v11).length));
            _lhs89[_lhs90] = _rhs121;
            _952_v12 = _rhs122;
            _953_v13 = _rhs123;
            _956_v16 = _rhs124;
          }
        }
      }
      r1 = (_this).f30;
      if (_this.f35) {
        let _957_v17;
        let _nw153 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
        _957_v17 = _nw153;
        let _958_v18;
        _958_v18 = new _dafny.CodePoint('p'.codePointAt(0));
        let _959_v19;
        _959_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(768),_958_v18);
        let _960_v20;
        _960_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_dafny.Seq.UnicodeFromString("qefsgg"));
        let _index117 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_957_v17).length));
        (_957_v17)[_index117] = (_module.__default.fm47(_this.f29, _959_v19, globalState)).Merge(_960_v20);
        let _961_v21;
        let _nw154 = Array((new BigNumber(26)).toNumber()).fill(false);
        _961_v21 = _nw154;
        let _962_v22;
        _962_v22 = _dafny.Map.Empty.slice().updateUnsafe(_961_v21,_this.f28);
        let _963_v23;
        _963_v23 = _dafny.Seq.of((_this).f34);
        let _964_v24;
        _964_v24 = _module.D8.create_DC24(_this.f28, (_this).f34, (_this).f32, _this.f31, _963_v23);
        let _965_v25;
        let _nw155 = Array((new BigNumber(9)).toNumber());
        _nw155[(_dafny.ZERO).toNumber()] = ((_this).f34).isLessThan(new BigNumber((_962_v22).length));
        _nw155[(_dafny.ONE).toNumber()] = _this.f31;
        _nw155[(new BigNumber(2)).toNumber()] = ((_this.f31) ? (true) : (_module.__default.fm1(_this.f31, (_this).f30, globalState)));
        _nw155[(new BigNumber(3)).toNumber()] = (new BigNumber(711)).isEqualTo(new BigNumber(238));
        _nw155[(new BigNumber(4)).toNumber()] = !_dafny.Seq.contains((_964_v24).dtor_cf50, (_963_v23)[_module.__default.safeIndex((_this).f34, new BigNumber((_963_v23).length))]);
        _nw155[(new BigNumber(5)).toNumber()] = true;
        _nw155[(new BigNumber(6)).toNumber()] = _this.f29;
        _nw155[(new BigNumber(7)).toNumber()] = _this.f31;
        _nw155[(new BigNumber(8)).toNumber()] = !_dafny.areEqual((_this).f32, _dafny.Seq.of(_this.f31, _this.f31));
        _965_v25 = _nw155;
        let _index118 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_957_v17).length));
        let _rhs125 = _this.f31;
        let _rhs126 = ((_this).f30).plus(new BigNumber(845));
        let _rhs127 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.UnicodeFromString("yowl")).length)).plus((_this).f34),_dafny.Seq.Create(_module.__default.abs(new BigNumber(526)), ((_966_v18) => function (_967_i1) {
          return _966_v18;
        })(_958_v18)));
        let _rhs128 = _965_v25;
        let _rhs129 = true;
        let _lhs91 = globalState;
        let _lhs92 = _957_v17;
        let _lhs93 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_957_v17).length));
        let _lhs94 = globalState;
        let _lhs95 = globalState;
        _lhs91.f26 = _rhs125;
        r1 = _rhs126;
        _lhs92[_lhs93] = _rhs127;
        _lhs94.f21 = _rhs128;
        _lhs95.f12 = _rhs129;
        (globalState).f13 = ((_this).f34).plus((_this).f34);
        let _968_v26;
        _968_v26 = _dafny.Seq.UnicodeFromString("uiiih");
        (_this).f31 = ((_this).f32)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_968_v26, _968_v26)).length))), new BigNumber(((_this).f32).length))];
        let _969_v27;
        _969_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_this.f29);
        let _970_v28;
        let _nw156 = Array((new BigNumber(5)).toNumber());
        _nw156[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Set.fromElements(new BigNumber(931))).length);
        _nw156[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_this).f34);
        _nw156[(new BigNumber(2)).toNumber()] = new BigNumber(((_969_v27).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-557),_this.f35))).length);
        _nw156[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((_this).f30, (_this).f34);
        _nw156[(new BigNumber(4)).toNumber()] = (_this).f30;
        _970_v28 = _nw156;
        let _index119 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_970_v28).length));
        (_970_v28)[_index119] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_this.f28),(_this).f34)).length);
        let _index120 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_970_v28).length));
        (_970_v28)[_index120] = (_this).f30;
        let _971_v29;
        let _nw157 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
        _971_v29 = _nw157;
        let _972_v30;
        _972_v30 = _dafny.MultiSet.fromElements(new BigNumber((_963_v23).length));
        let _973_v31;
        _973_v31 = _dafny.Set.fromElements((_this).f34, new BigNumber((_972_v30).cardinality()));
        let _974_v32;
        _974_v32 = _module.D1.create_DC2(_961_v21);
        let _975_v33;
        _975_v33 = _module.D1.create_DC4(_974_v32);
        let _976_v34;
        _976_v34 = _module.D1.create_DC4(_974_v32);
        let _977_v35;
        _977_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_976_v34);
        let _978_v36;
        _978_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_973_v31).length),_977_v35);
        let _979_v37;
        _979_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,(_dafny.ZERO).minus(new BigNumber((_978_v36).length)));
        let _980_v38;
        let _nw158 = Array((new BigNumber(27)).toNumber());
        _nw158[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
        _nw158[(_dafny.ONE).toNumber()] = _958_v18;
        _nw158[(new BigNumber(2)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(3)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(4)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(5)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(6)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(7)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(8)).toNumber()] = (((_959_v19).contains(new BigNumber(125))) ? ((_959_v19).get(new BigNumber(125))) : (_958_v18));
        _nw158[(new BigNumber(9)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
        _nw158[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _nw158[(new BigNumber(12)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(13)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(14)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(15)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(16)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(17)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(18)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(19)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(20)).toNumber()] = _module.__default.fm48(true, _module.__default.fm1(_this.f28, (_this).f34, globalState), new BigNumber(((_this).f32).length), globalState);
        _nw158[(new BigNumber(21)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(22)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(23)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(24)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
        _nw158[(new BigNumber(25)).toNumber()] = _958_v18;
        _nw158[(new BigNumber(26)).toNumber()] = _958_v18;
        _980_v38 = _nw158;
        let _981_v39;
        let _nw159 = new _module.C0();
        _nw159.__ctor(_980_v38);
        _981_v39 = _nw159;
        let _982_v40;
        _982_v40 = _module.D1.create_DC3(_this.f35, (_this).f30, _981_v39, new BigNumber(((_this).f32).length));
        let _983_v41;
        let _nw160 = new _module.C3();
        _nw160.__ctor((_this).f34, _971_v29, _958_v18, (_this).f34, !(_this.f35), _dafny.Seq.update((_this).f32, _module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length)), _this.f29), (_this).f33, (((_979_v37).contains(_this.f29)) ? ((_979_v37).get(_this.f29)) : ((_this).f30)), (_982_v40).dtor_cf3, !(_this.f29), _this.f31);
        _983_v41 = _nw160;
      } else {
        let _984_v42;
        _984_v42 = _dafny.Set.fromElements(_this.f31);
        let _985_v43;
        _985_v43 = _984_v42;
        let _986_v44;
        _986_v44 = new _dafny.CodePoint('k'.codePointAt(0));
        let _987_v45;
        _987_v45 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f34));
        let _988_v46;
        _988_v46 = _module.D4.create_DC10(_module.__default.fm49(_dafny.MultiSet.fromElements(new BigNumber(977), new BigNumber(((_985_v43)).length)), _986_v44, _dafny.Seq.of(_dafny.Seq.of((_dafny.ZERO).minus((_this).f34)), _987_v45, _987_v45), globalState));
        _988_v46 = _988_v46;
        let _989_v47;
        let _nw161 = Array((new BigNumber(18)).toNumber()).fill(false);
        _989_v47 = _nw161;
        (globalState).f21 = _989_v47;
        (globalState).f1 = _this.f35;
        let _990_v48;
        _990_v48 = _dafny.Seq.UnicodeFromString("hngyxcapn");
        let _991_v49;
        let _nw162 = Array((new BigNumber(3)).toNumber());
        _nw162[(_dafny.ZERO).toNumber()] = _986_v44;
        _nw162[(_dafny.ONE).toNumber()] = _986_v44;
        _nw162[(new BigNumber(2)).toNumber()] = (_990_v48)[_module.__default.safeIndex((_this).f30, new BigNumber((_990_v48).length))];
        _991_v49 = _nw162;
        let _index121 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_991_v49).length));
        (_991_v49)[_index121] = _986_v44;
        let _index122 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_991_v49).length));
        (_991_v49)[_index122] = (((_this).fm6(globalState)) ? (_986_v44) : (_986_v44));
        let _992_v50;
        _992_v50 = _dafny.Map.Empty.slice().updateUnsafe(_989_v47,!(true));
        (_this).f28 = !(new BigNumber((_992_v50).length)).isEqualTo((_dafny.ZERO).minus((_this).f34));
      }
      let _993_v51;
      _993_v51 = _dafny.Set.fromElements((_this).f30);
      let _994_v52;
      _994_v52 = _dafny.MultiSet.fromElements((_this).f34, (_this).f30);
      let _995_v53;
      _995_v53 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,new BigNumber(((_this).f32).length));
      let _rhs130 = _module.__default.safeDivisionInt((_this).f30, (_this).f30);
      let _rhs131 = _module.__default.safeModuloInt(new BigNumber(((_994_v52).Intersect(_994_v52)).cardinality()), ((_this.f29) ? (new BigNumber((_995_v53).length)) : ((_this).f30)));
      let _rhs132 = (_993_v51).Intersect(_993_v51);
      let _lhs96 = globalState;
      r1 = _rhs130;
      _lhs96.f13 = _rhs131;
      _993_v51 = _rhs132;
      let _996_v54;
      _996_v54 = _dafny.Seq.UnicodeFromString("cmjcaq");
      let _997_v55;
      _997_v55 = _dafny.Set.fromElements(_996_v54);
      let _998_v56;
      _998_v56 = new _dafny.CodePoint('h'.codePointAt(0));
      (globalState).f13 = new BigNumber((((_997_v55).Intersect(_dafny.Set.fromElements(_996_v54, _996_v54, _dafny.Seq.update(_996_v54, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f34), new BigNumber((_996_v54).length)), _998_v56)))).Difference((_997_v55).Union(_997_v55))).length);
      let _999_v57;
      _999_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_this.f31);
      r0 = !((((_999_v57).contains(((_this).f30).plus((_this).f34))) ? ((_999_v57).get(((_this).f30).plus((_this).f34))) : (_this.f35)));
      r1 = (_this).f30;
      let _1000_v58;
      _1000_v58 = _module.D5.create_DC12(_dafny.Seq.UnicodeFromString("rpyunkfoa"));
      r2 = _dafny.Seq.Concat(_996_v54, (_1000_v58).dtor_cf26);
      return [r0, r1, r2];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f28 = false;
      this._f29 = false;
      this._f31 = false;
      this._f35 = false;
      this._f34 = _dafny.ZERO;
      this._f32 = _dafny.Seq.of();
      this._f33 = _dafny.Seq.of();
      this._f30 = _dafny.ZERO;
      this._f47 = false;
      this._f48 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    set f31(value) {
      let _this = this;
      _this._f31 = value;
    };
    get f35() {
      let _this = this;
      return _this._f35;
    };
    set f35(value) {
      let _this = this;
      _this._f35 = value;
    };
    get f34() {
      let _this = this;
      return _this._f34;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    __ctor(f47, f48, f34, f35, f32, f33, f30, f31, f28, f29) {
      let _this = this;
      (_this)._f47 = f47;
      (_this)._f48 = f48;
      (_this)._f34 = f34;
      (_this)._f35 = f35;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm6(globalState) {
      let _this = this;
      return ((_this).f34).isLessThan((_this).f30);
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (function () {
        let _coll25 = new _dafny.Set();
        for (const _compr_25 of (_dafny.Seq.of(new BigNumber(-735))).Elements) {
          let _1001_v0 = _compr_25;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-735)), _1001_v0)) {
            _coll25.add((_1001_v0).minus(new BigNumber(599)));
          }
        }
        return _coll25;
      }()).Union((_dafny.Set.fromElements((_this).f48)).Difference(_dafny.Set.fromElements(new BigNumber(255))));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f32;
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(741)), function (_1002_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("bfcfykdh"));
    };
    fm2(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f48)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_this).f48))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_this).f30));
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      r1 = (_this).f48;
      let _1003_v0;
      _1003_v0 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1004_v1;
      _1004_v1 = _dafny.MultiSet.fromElements(_1003_v0, _1003_v0);
      let _1005_v2;
      _1005_v2 = _dafny.MultiSet.fromElements((_this).f48);
      let _1006_v3;
      _1006_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1004_v1,_module.__default.fm39((_this).f47, _1005_v2, _this.f31, globalState));
      let _1007_v4;
      _1007_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1006_v3,(_this).f47);
      _1007_v4 = (_1007_v4).update(_dafny.Map.Empty.slice().updateUnsafe(_1004_v1,(_this).f34), !((_this).fm6(globalState)));
      let _1008_v5;
      _1008_v5 = _dafny.Seq.of((_this).f48);
      let _1009_v6;
      _1009_v6 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f28),_dafny.Seq.IsPrefixOf(_1008_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), function (_1010_i0) {
        return (_this).f48;
      })));
      let _1011_v7;
      _1011_v7 = _dafny.Seq.of(((_this).f32)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length))]);
      let _1012_v8;
      let _nw163 = Array((new BigNumber(13)).toNumber()).fill(_module.D4.Default());
      _1012_v8 = _nw163;
      let _1013_v9;
      _1013_v9 = _module.D4.create_DC11(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), ((_1014_v0) => function (_1015_i1) {
  return _1014_v0;
})(_1003_v0))).length), true, (_this).f30);
      let _index123 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1012_v8).length));
      (_1012_v8)[_index123] = _1013_v9;
      let _index124 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1012_v8).length));
      let _rhs133 = _1009_v6;
      let _rhs134 = _1011_v7;
      let _rhs135 = _module.D4.create_DC11((_this).f30, ((_this).f48).isLessThanOrEqualTo((_this).f34), new BigNumber(-352));
      let _rhs136 = !(_this.f35) || (((_this).f34).isEqualTo((_this).f30));
      let _lhs97 = _1012_v8;
      let _lhs98 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1012_v8).length));
      let _lhs99 = _this;
      _1009_v6 = _rhs133;
      _1011_v7 = _rhs134;
      _lhs97[_lhs98] = _rhs135;
      _lhs99.f35 = _rhs136;
      let _1016_v10;
      _1016_v10 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(941));
      let _1017_v11;
      _1017_v11 = _dafny.MultiSet.fromElements(_this.f29);
      let _1018_v12;
      _1018_v12 = _dafny.Map.Empty.slice().updateUnsafe(true,_1003_v0);
      let _1019_v13;
      let _nw164 = Array((new BigNumber(29)).toNumber());
      _nw164[(_dafny.ZERO).toNumber()] = (_this).f30;
      _nw164[(_dafny.ONE).toNumber()] = (_this).f30;
      _nw164[(new BigNumber(2)).toNumber()] = (_this).f48;
      _nw164[(new BigNumber(3)).toNumber()] = (_1008_v5)[_module.__default.safeIndex((((_1005_v2).contains((_this).f34)) ? ((_1005_v2).get((_this).f34)) : ((_this).f30)), new BigNumber((_1008_v5).length))];
      _nw164[(new BigNumber(4)).toNumber()] = (_this).f30;
      _nw164[(new BigNumber(5)).toNumber()] = (_this).f30;
      _nw164[(new BigNumber(6)).toNumber()] = (_this).f34;
      _nw164[(new BigNumber(7)).toNumber()] = (_this).f34;
      _nw164[(new BigNumber(8)).toNumber()] = new BigNumber(381);
      _nw164[(new BigNumber(9)).toNumber()] = (_this).f30;
      _nw164[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((_this).f34);
      _nw164[(new BigNumber(11)).toNumber()] = (_this).f48;
      _nw164[(new BigNumber(12)).toNumber()] = (_this).f48;
      _nw164[(new BigNumber(13)).toNumber()] = (_this).f48;
      _nw164[(new BigNumber(14)).toNumber()] = new BigNumber((_1008_v5).length);
      _nw164[(new BigNumber(15)).toNumber()] = new BigNumber(803);
      _nw164[(new BigNumber(16)).toNumber()] = (_this).f34;
      _nw164[(new BigNumber(17)).toNumber()] = (_this).f34;
      _nw164[(new BigNumber(18)).toNumber()] = (_this).f48;
      _nw164[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_this.f29, false, _module.__default.fm1(_this.f28, (((_1016_v10).contains(_this.f28)) ? ((_1016_v10).get(_this.f28)) : ((_this).f48)), globalState), _this.f35, true)).cardinality());
      _nw164[(new BigNumber(20)).toNumber()] = (_this).f48;
      _nw164[(new BigNumber(21)).toNumber()] = new BigNumber((_1017_v11).cardinality());
      _nw164[(new BigNumber(22)).toNumber()] = (_this).f34;
      _nw164[(new BigNumber(23)).toNumber()] = (_this).f48;
      _nw164[(new BigNumber(24)).toNumber()] = (_this).f48;
      _nw164[(new BigNumber(25)).toNumber()] = (_this).f34;
      _nw164[(new BigNumber(26)).toNumber()] = new BigNumber((_1018_v12).length);
      _nw164[(new BigNumber(27)).toNumber()] = (_this).f34;
      _nw164[(new BigNumber(28)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(656)), ((_1020_v0) => function (_1021_i2) {
        return _1020_v0;
      })(_1003_v0))).length);
      _1019_v13 = _nw164;
      let _1022_v14;
      _1022_v14 = _dafny.Set.fromElements(_1019_v13, _1019_v13, _1019_v13, _1019_v13, _1019_v13);
      _1022_v14 = _1022_v14;
      let _1023_v15;
      _1023_v15 = _module.D8.create_DC22(new BigNumber((_dafny.Seq.UnicodeFromString("djojhst")).length), _module.__default.safeModuloInt(new BigNumber(948), (_this).f48));
      let _source18 = _1023_v15;
      if (_source18.is_DC22) {
        let _1024___mcc_h0 = (_source18).cf41;
        let _1025___mcc_h1 = (_source18).cf42;
        let _1026_cf42 = _1025___mcc_h1;
        let _1027_cf41 = _1024___mcc_h0;
        let _rhs137 = _this.f31;
        let _rhs138 = (_this).fm3(_this.f28, _this.f29, globalState);
        let _lhs100 = globalState;
        _lhs100.f1 = _rhs137;
        r2 = _rhs138;
        (globalState).f13 = _1026_cf42;
        if (!(_this.f28)) {
          let _1028_v16;
          _1028_v16 = _dafny.Set.fromElements(_1017_v11, _1017_v11);
          let _1029_v17;
          _1029_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1017_v11,(_this).f48);
          let _1030_v19;
          _1030_v19 = _module.D8.create_DC24(true, new BigNumber((_1009_v6).length), (_this).f32, _this.f35, _1008_v5);
          let _1031_v20;
          _1031_v20 = _module.D8.create_DC25(_1030_v19);
          let _1032_v21;
          let _nw165 = Array((new BigNumber(8)).toNumber());
          _nw165[(_dafny.ZERO).toNumber()] = _1028_v16;
          _nw165[(_dafny.ONE).toNumber()] = function () {
            let _coll26 = new _dafny.Set();
            for (const _compr_26 of (_1029_v17).Keys.Elements) {
              let _1033_v18 = _compr_26;
              if ((_1029_v17).contains(_1033_v18)) {
                _coll26.add(_1033_v18);
              }
            }
            return _coll26;
          }();
          _nw165[(new BigNumber(2)).toNumber()] = (_1028_v16).Intersect(_module.__default.fm40(_1027_cf41, _this.f29, _1031_v20, globalState));
          _nw165[(new BigNumber(3)).toNumber()] = _1028_v16;
          _nw165[(new BigNumber(4)).toNumber()] = _1028_v16;
          _nw165[(new BigNumber(5)).toNumber()] = _1028_v16;
          _nw165[(new BigNumber(6)).toNumber()] = (_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f31, _this.f35, _this.f35)))).Intersect(_1028_v16);
          _nw165[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_1017_v11, _dafny.MultiSet.fromElements(false, _this.f28), _1017_v11, _1017_v11);
          _1032_v21 = _nw165;
          let _index125 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1032_v21).length));
          (_1032_v21)[_index125] = _1028_v16;
          let _index126 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1032_v21).length));
          (_1032_v21)[_index126] = _1028_v16;
          (globalState).f13 = _module.__default.fm39(_this.f35, _1005_v2, (!(_this.f35)) || (_this.f29), globalState);
          let _1034_v22;
          let _nw166 = Array((new BigNumber(15)).toNumber()).fill([]);
          _1034_v22 = _nw166;
          let _1035_v23;
          let _init20 = function (_1036_i3) {
            return _this.f31;
          };
          let _nw167 = Array((new BigNumber(27)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw167.length); _i0_20++) {
            _nw167[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _1035_v23 = _nw167;
          let _index127 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_1034_v22).length));
          (_1034_v22)[_index127] = _1035_v23;
          let _index128 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1019_v13).length));
          (_1019_v13)[_index128] = _1027_cf41;
          let _1037_v24;
          let _nw168 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1037_v24 = _nw168;
          let _index129 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_1034_v22).length));
          let _index130 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1019_v13).length));
          let _rhs139 = ((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f47)).update(_this.f35, _this.f35)).update(_this.f35, (_this).f47);
          let _rhs140 = _1035_v23;
          let _rhs141 = _1027_cf41;
          let _rhs142 = _1037_v24;
          let _rhs143 = _dafny.Seq.Concat((_this).f32, (_this).f32);
          let _lhs101 = _1034_v22;
          let _lhs102 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_1034_v22).length));
          let _lhs103 = _1019_v13;
          let _lhs104 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1019_v13).length));
          _1009_v6 = _rhs139;
          _lhs101[_lhs102] = _rhs140;
          _lhs103[_lhs104] = _rhs141;
          _1037_v24 = _rhs142;
          _1011_v7 = _rhs143;
          let _1038_v25;
          let _nw169 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _1038_v25 = _nw169;
          let _1039_v26;
          _1039_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(183),_this.f35);
          let _1040_v27;
          _1040_v27 = _dafny.Seq.of(_1003_v0);
          let _1041_v28;
          let _nw170 = new _module.C1();
          _nw170.__ctor(_this.f29, _this.f31, _this.f35);
          _1041_v28 = _nw170;
          let _1042_v29;
          _1042_v29 = _dafny.Set.fromElements((_this).f48, new BigNumber(696), _1027_cf41, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1041_v28,_module.__default.fm41(_this.f31, _1003_v0, _this.f28, (_this).f48, globalState))).length));
          let _1043_v30;
          let _init21 = ((_1044_v0) => function (_1045_i4) {
            return _1044_v0;
          })(_1003_v0);
          let _nw171 = Array((new BigNumber(2)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw171.length); _i0_21++) {
            _nw171[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _1043_v30 = _nw171;
          let _1046_v31;
          let _nw172 = new _module.C0();
          _nw172.__ctor(_1043_v30);
          _1046_v31 = _nw172;
          let _1047_v32;
          _1047_v32 = _dafny.Map.Empty.slice().updateUnsafe(_module.D9.create_DC27(_1046_v31, _dafny.Seq.UnicodeFromString("dxwqwtkoe")),_1035_v23);
          let _nw173 = Array((new BigNumber(22)).toNumber());
          _nw173[(_dafny.ZERO).toNumber()] = _1027_cf41;
          _nw173[(_dafny.ONE).toNumber()] = _1026_cf42;
          _nw173[(new BigNumber(2)).toNumber()] = ((_this).f30).plus(_1026_cf42);
          _nw173[(new BigNumber(3)).toNumber()] = _module.__default.fm39((((_1039_v26).contains(new BigNumber((_1039_v26).length))) ? ((_1039_v26).get(new BigNumber((_1039_v26).length))) : (_this.f28)), _1005_v2, _this.f31, globalState);
          _nw173[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(_1026_cf42, new BigNumber((_1040_v27).length));
          _nw173[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1008_v5, _dafny.Seq.of((_1019_v13)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_1019_v13).length))]))).length);
          _nw173[(new BigNumber(6)).toNumber()] = (_this).f30;
          _nw173[(new BigNumber(7)).toNumber()] = (_module.__default.fm39(false, _1005_v2, _this.f29, globalState)).multipliedBy((_this).f48);
          _nw173[(new BigNumber(8)).toNumber()] = (_this).f30;
          _nw173[(new BigNumber(9)).toNumber()] = new BigNumber(-634);
          _nw173[(new BigNumber(10)).toNumber()] = _1026_cf42;
          _nw173[(new BigNumber(11)).toNumber()] = _1027_cf41;
          _nw173[(new BigNumber(12)).toNumber()] = (_this).f34;
          _nw173[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt((_this).f30, new BigNumber(((_this).fm4(_module.__default.fm1(_this.f35, (_1019_v13)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_1019_v13).length))], globalState), _1042_v29, _1026_cf42, _this.f28, globalState)).length));
          _nw173[(new BigNumber(14)).toNumber()] = _module.__default.fm39(_this.f31, _1005_v2, !((_1041_v28).f43), globalState);
          _nw173[(new BigNumber(15)).toNumber()] = (_this).f30;
          _nw173[(new BigNumber(16)).toNumber()] = new BigNumber(((_1047_v32).Merge(_1047_v32)).length);
          _nw173[(new BigNumber(17)).toNumber()] = (_this).f48;
          _nw173[(new BigNumber(18)).toNumber()] = (_this).f30;
          _nw173[(new BigNumber(19)).toNumber()] = (_this).f34;
          _nw173[(new BigNumber(20)).toNumber()] = (_1019_v13)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_1019_v13).length))];
          _nw173[(new BigNumber(21)).toNumber()] = (_this).f34;
          _1038_v25 = _nw173;
          _1026_cf42 = _module.__default.fm39(_this.f29, _dafny.MultiSet.FromArray(_module.__default.fm42(_1040_v27, (_this).f48, _dafny.Seq.UnicodeFromString("gkp"), globalState)), (_1041_v28).f43, globalState);
        } else {
          (_this).f35 = ((_this).f30).isEqualTo(_module.__default.fm39((_this).fm6(globalState), _1005_v2, (_this).f47, globalState));
          _1026_cf42 = new BigNumber(441);
          r1 = ((!(_this.f28)) ? ((_dafny.ZERO).minus(((((_1017_v11).contains(!(_this.f29))) ? ((_1017_v11).get(!(_this.f29))) : (new BigNumber((_1008_v5).length)))).plus(new BigNumber(545)))) : ((_this).f30));
          let _1048_v33;
          let _nw174 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1048_v33 = _nw174;
          let _1049_v34;
          _1049_v34 = _dafny.Set.fromElements((_this).f48);
          let _1050_v35;
          _1050_v35 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm41(_this.f35, _1003_v0, _this.f29, (_this).f34, globalState),(_1005_v2).update((_dafny.ZERO).minus((_this).f30), _module.__default.abs(new BigNumber((_1049_v34).length))));
          let _index131 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1048_v33).length));
          (_1048_v33)[_index131] = (((_1050_v35).contains(_module.__default.fm41(false, _1003_v0, _this.f28, _module.__default.fm39(false, _1005_v2, _this.f35, globalState), globalState))) ? ((_1050_v35).get(_module.__default.fm41(false, _1003_v0, _this.f28, _module.__default.fm39(false, _1005_v2, _this.f35, globalState), globalState))) : (_1005_v2));
          let _index132 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1048_v33).length));
          (_1048_v33)[_index132] = _module.__default.fm43(globalState);
          let _1051_v36;
          let _nw175 = Array((_dafny.ONE).toNumber());
          _nw175[(_dafny.ZERO).toNumber()] = _this.f35;
          _1051_v36 = _nw175;
          let _index133 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1051_v36).length));
          (_1051_v36)[_index133] = (_this).f47;
          let _index134 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1051_v36).length));
          (_1051_v36)[_index134] = ((_this).f30).isEqualTo((_this).f48);
        }
        let _1052_v37;
        _1052_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,(_this).f30);
        _1052_v37 = (_1052_v37).update(_1027_cf41, (_this).f34);
      } else if (_source18.is_DC23) {
        let _1053___mcc_h2 = (_source18).cf43;
        let _1054___mcc_h3 = (_source18).cf44;
        let _1055___mcc_h4 = (_source18).cf45;
        let _1056_cf45 = _1055___mcc_h4;
        let _1057_cf44 = _1054___mcc_h3;
        let _1058_cf43 = _1053___mcc_h2;
        _1003_v0 = _1003_v0;
        let _1059_v38;
        _1059_v38 = _dafny.Seq.UnicodeFromString("nujp");
        (globalState).f1 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_1059_v38, _module.__default.safeIndex((_this).f48, new BigNumber((_1059_v38).length)), _1003_v0), _1059_v38);
        (_this).f35 = false;
        (globalState).f13 = (_this).f48;
      } else if (_source18.is_DC24) {
        let _1060___mcc_h5 = (_source18).cf46;
        let _1061___mcc_h6 = (_source18).cf47;
        let _1062___mcc_h7 = (_source18).cf48;
        let _1063___mcc_h8 = (_source18).cf49;
        let _1064___mcc_h9 = (_source18).cf50;
        let _1065_cf50 = _1064___mcc_h9;
        let _1066_cf49 = _1063___mcc_h8;
        let _1067_cf48 = _1062___mcc_h7;
        let _1068_cf47 = _1061___mcc_h6;
        let _1069_cf46 = _1060___mcc_h5;
        _1005_v2 = (_1005_v2).Difference(_1005_v2);
        _1003_v0 = new _dafny.CodePoint('i'.codePointAt(0));
        let _1070_v39;
        _1070_v39 = _dafny.Seq.UnicodeFromString("vvj");
        let _1071_v40;
        let _nw176 = Array((new BigNumber(5)).toNumber());
        _nw176[(_dafny.ZERO).toNumber()] = _1070_v39;
        _nw176[(_dafny.ONE).toNumber()] = _1070_v39;
        _nw176[(new BigNumber(2)).toNumber()] = _1070_v39;
        _nw176[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1070_v39, _1070_v39);
        _nw176[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_1070_v39, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f34), new BigNumber((_1070_v39).length)), _1003_v0);
        _1071_v40 = _nw176;
        let _index135 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1071_v40).length));
        (_1071_v40)[_index135] = _1070_v39;
        let _1072_v41;
        _1072_v41 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(849)), ((_1073_v0) => function (_1074_i5) {
          return _1073_v0;
        })(_1003_v0)));
        let _index136 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1071_v40).length));
        (_1071_v40)[_index136] = (_1072_v41)[_module.__default.safeIndex((_this).f30, new BigNumber((_1072_v41).length))];
        let _1075_v42;
        _1075_v42 = _dafny.Seq.of(_1008_v5, _1065_cf50, _1065_cf50);
        let _1076_v43;
        _1076_v43 = _dafny.Map.Empty.slice().updateUnsafe((_1075_v42)[_module.__default.safeIndex((_dafny.ZERO).minus(_1068_cf47), new BigNumber((_1075_v42).length))],_1066_cf49);
        let _1077_v44;
        _1077_v44 = _module.D6.create_DC15(_1008_v5);
        _1076_v43 = ((_1076_v43).Merge(_1076_v43)).update((_1077_v44).dtor_cf32, _1066_cf49);
      } else if (_source18.is_DC21) {
        let _1078___mcc_h10 = (_source18).cf40;
        let _1079_cf40 = _1078___mcc_h10;
        let _1080_v45;
        let _nw177 = new _module.C2();
        _nw177.__ctor((_this).f48, (_this).f30, _this.f31, _this.f28, _this.f35);
        _1080_v45 = _nw177;
        let _1081_v46;
        let _nw178 = Array((new BigNumber(8)).toNumber());
        _nw178[(_dafny.ZERO).toNumber()] = (_this).f47;
        _nw178[(_dafny.ONE).toNumber()] = _this.f35;
        _nw178[(new BigNumber(2)).toNumber()] = _this.f29;
        _nw178[(new BigNumber(3)).toNumber()] = true;
        _nw178[(new BigNumber(4)).toNumber()] = _this.f35;
        _nw178[(new BigNumber(5)).toNumber()] = _this.f35;
        _nw178[(new BigNumber(6)).toNumber()] = _this.f28;
        _nw178[(new BigNumber(7)).toNumber()] = _this.f35;
        _1081_v46 = _nw178;
        let _1082_v47;
        _1082_v47 = _module.D3.create_DC9((_1080_v45).fm30(_this.f35, globalState), (_this).fm6(globalState), _1081_v46, (_1011_v7)[_module.__default.safeIndex((_this).f48, new BigNumber((_1011_v7).length))]);
        let _1083_v48;
        _1083_v48 = _module.D6.create_DC15(_1008_v5);
        let _1084_v49;
        _1084_v49 = _dafny.Seq.of(_1008_v5);
        let _1085_v50;
        _1085_v50 = _module.D8.create_DC24(_this.f35, (_this).f34, (_this).f32, _this.f28, (_1084_v49)[_module.__default.safeIndex((_this).f34, new BigNumber((_1084_v49).length))]);
        let _pat_let_tv10 = _1085_v50;
        let _1086_v51;
        let _1087_v52;
        let _out54;
        let _out55;
        let _outcollector16 = (_this).m14(_1003_v0, (_1082_v47).dtor_cf20, function (_pat_let13_0) {
          return function (_1088_dt__update__tmp_h0) {
            return function (_pat_let14_0) {
              return function (_1089_dt__update_hcf32_h0) {
                return _module.D6.create_DC15(_1089_dt__update_hcf32_h0);
              }(_pat_let14_0);
            }((_pat_let_tv10).dtor_cf50);
          }(_pat_let13_0);
        }(_1083_v48), globalState);
        _out54 = _outcollector16[0];
        _out55 = _outcollector16[1];
        _1086_v51 = _out54;
        _1087_v52 = _out55;
        (_this).f29 = _this.f29;
        (globalState).f13 = (_this).f34;
      } else {
        let _1090___mcc_h11 = (_source18).cf51;
        let _1091_cf51 = _1090___mcc_h11;
        let _1092_v53;
        _1092_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f34);
        let _1093_v54;
        _1093_v54 = _dafny.Seq.UnicodeFromString("e");
        let _1094_v55;
        _1094_v55 = _dafny.Seq.of(_1092_v53);
        _1092_v53 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1093_v54).length),new BigNumber((_1093_v54).length))).Merge((_1094_v55)[_module.__default.safeIndex((_this).f48, new BigNumber((_1094_v55).length))]);
        (globalState).f13 = new BigNumber(342);
        let _source19 = _1023_v15;
        if (_source19.is_DC22) {
          let _1095___mcc_h12 = (_source19).cf41;
          let _1096___mcc_h13 = (_source19).cf42;
          let _1097_cf42 = _1096___mcc_h13;
          let _1098_cf41 = _1095___mcc_h12;
          let _1099_v56;
          _1099_v56 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f30).minus(_1097_cf42),_this.f28);
          _1099_v56 = (_1099_v56).update((new BigNumber(991)).multipliedBy((_this).f34), ((_this).f34).isEqualTo(_1097_cf42));
          let _1100_v58;
          _1100_v58 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_1099_v56);
          let _1101_v59;
          let _nw179 = Array((new BigNumber(10)).toNumber());
          _nw179[(_dafny.ZERO).toNumber()] = function () {
            let _coll27 = new _dafny.Map();
            for (const _compr_27 of (_1008_v5).Elements) {
              let _1102_v57 = _compr_27;
              if (_dafny.Seq.contains(_1008_v5, _1102_v57)) {
                _coll27.push([_module.__default.safeDivisionInt(_1102_v57, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f35,new BigNumber((_dafny.Set.fromElements(_1008_v5)).length))).length)),_this.f31]);
              }
            }
            return _coll27;
          }();
          _nw179[(_dafny.ONE).toNumber()] = _1099_v56;
          _nw179[(new BigNumber(2)).toNumber()] = _1099_v56;
          _nw179[(new BigNumber(3)).toNumber()] = _1099_v56;
          _nw179[(new BigNumber(4)).toNumber()] = (((_1100_v58).contains((_this).f47)) ? ((_1100_v58).get((_this).f47)) : (_1099_v56));
          _nw179[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_this.f35);
          _nw179[(new BigNumber(6)).toNumber()] = _1099_v56;
          _nw179[(new BigNumber(7)).toNumber()] = _1099_v56;
          _nw179[(new BigNumber(8)).toNumber()] = _1099_v56;
          _nw179[(new BigNumber(9)).toNumber()] = _1099_v56;
          _1101_v59 = _nw179;
          let _1103_v60;
          let _nw180 = Array((new BigNumber(9)).toNumber());
          _nw180[(_dafny.ZERO).toNumber()] = _1101_v59;
          _nw180[(_dafny.ONE).toNumber()] = _1101_v59;
          _nw180[(new BigNumber(2)).toNumber()] = _1101_v59;
          _nw180[(new BigNumber(3)).toNumber()] = _1101_v59;
          _nw180[(new BigNumber(4)).toNumber()] = _1101_v59;
          _nw180[(new BigNumber(5)).toNumber()] = _1101_v59;
          _nw180[(new BigNumber(6)).toNumber()] = _1101_v59;
          _nw180[(new BigNumber(7)).toNumber()] = ((_this.f35) ? (_1101_v59) : (_1101_v59));
          _nw180[(new BigNumber(8)).toNumber()] = _1101_v59;
          _1103_v60 = _nw180;
          let _index137 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_1103_v60).length));
          (_1103_v60)[_index137] = _1101_v59;
          let _1104_v61;
          _1104_v61 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,new BigNumber(-133))).length), new BigNumber((_1009_v6).length), (_this).f48, _1098_cf41);
          let _1105_v63;
          let _nw181 = Array((new BigNumber(17)).toNumber());
          _nw181[(_dafny.ZERO).toNumber()] = _this.f28;
          _nw181[(_dafny.ONE).toNumber()] = (function () {
            let _coll28 = new _dafny.Set();
            for (const _compr_28 of _dafny.IntegerRange(new BigNumber(57), new BigNumber(761))) {
              let _1106_v62 = _compr_28;
              if (((new BigNumber(57)).isLessThanOrEqualTo(_1106_v62)) && ((_1106_v62).isLessThan(new BigNumber(761)))) {
                _coll28.add((_1106_v62).multipliedBy((_this).f30));
              }
            }
            return _coll28;
          }()).IsProperSubsetOf(_1104_v61);
          _nw181[(new BigNumber(2)).toNumber()] = (!((_this).f47)) || ((_this).f47);
          _nw181[(new BigNumber(3)).toNumber()] = (_module.__default.fm1(_this.f35, _1098_cf41, globalState)) || (_this.f28);
          _nw181[(new BigNumber(4)).toNumber()] = false;
          _nw181[(new BigNumber(5)).toNumber()] = _this.f28;
          _nw181[(new BigNumber(6)).toNumber()] = _this.f31;
          _nw181[(new BigNumber(7)).toNumber()] = (_this).f47;
          _nw181[(new BigNumber(8)).toNumber()] = false;
          _nw181[(new BigNumber(9)).toNumber()] = (_this).f47;
          _nw181[(new BigNumber(10)).toNumber()] = _this.f35;
          _nw181[(new BigNumber(11)).toNumber()] = (_this).fm6(globalState);
          _nw181[(new BigNumber(12)).toNumber()] = _this.f28;
          _nw181[(new BigNumber(13)).toNumber()] = !((_dafny.MultiSet.FromArray(_1011_v7)).update(!(_this.f31), _module.__default.abs(_1097_cf42))).contains(_this.f29);
          _nw181[(new BigNumber(14)).toNumber()] = _this.f29;
          _nw181[(new BigNumber(15)).toNumber()] = _this.f35;
          _nw181[(new BigNumber(16)).toNumber()] = _this.f31;
          _1105_v63 = _nw181;
          let _index138 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1105_v63).length));
          (_1105_v63)[_index138] = _this.f29;
          let _1107_v64;
          _1107_v64 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(889)), ((_1108_v0) => function (_1109_i6) {
            return _1108_v0;
          })(_1003_v0)));
          let _index139 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_1103_v60).length));
          let _index140 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1105_v63).length));
          let _rhs144 = new BigNumber(788);
          let _rhs145 = _1101_v59;
          let _rhs146 = (new BigNumber((_dafny.Seq.Concat((_1107_v64)[_module.__default.safeIndex(_1098_cf41, new BigNumber((_1107_v64).length))], _1093_v54)).length)).isLessThanOrEqualTo(new BigNumber((_1017_v11).cardinality()));
          let _lhs105 = globalState;
          let _lhs106 = _1103_v60;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_1103_v60).length));
          let _lhs108 = _1105_v63;
          let _lhs109 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1105_v63).length));
          _lhs105.f13 = _rhs144;
          _lhs106[_lhs107] = _rhs145;
          _lhs108[_lhs109] = _rhs146;
          (globalState).f13 = (_this).f34;
          let _1110_v65;
          let _init22 = ((_1111_v11) => function (_1112_i7) {
            return _1111_v11;
          })(_1017_v11);
          let _nw182 = Array((new BigNumber(17)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw182.length); _i0_22++) {
            _nw182[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _1110_v65 = _nw182;
          let _1113_v66;
          _1113_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1110_v65,_1098_cf41);
          _1113_v66 = (_1113_v66).update(_1110_v65, (_this).f34);
        } else if (_source19.is_DC23) {
          let _1114___mcc_h14 = (_source19).cf43;
          let _1115___mcc_h15 = (_source19).cf44;
          let _1116___mcc_h16 = (_source19).cf45;
          let _1117_cf45 = _1116___mcc_h16;
          let _1118_cf44 = _1115___mcc_h15;
          let _1119_cf43 = _1114___mcc_h14;
          (globalState).f13 = (_this).f34;
          let _1120_v67;
          _1120_v67 = _dafny.Seq.of(_1093_v54, _dafny.Seq.Concat(_1093_v54, _1093_v54));
          _1120_v67 = _dafny.Seq.Concat(_1120_v67, _1120_v67);
          let _1121_v68;
          _1121_v68 = _dafny.Set.fromElements(_1117_cf45, (_this).f47);
          let _1122_v69;
          let _nw183 = new _module.C1();
          _nw183.__ctor((new BigNumber((_1121_v68).length)).isEqualTo((_this).f30), _this.f28, _this.f31);
          _1122_v69 = _nw183;
          _1019_v13 = _1019_v13;
        } else if (_source19.is_DC24) {
          let _1123___mcc_h17 = (_source19).cf46;
          let _1124___mcc_h18 = (_source19).cf47;
          let _1125___mcc_h19 = (_source19).cf48;
          let _1126___mcc_h20 = (_source19).cf49;
          let _1127___mcc_h21 = (_source19).cf50;
          let _1128_cf50 = _1127___mcc_h21;
          let _1129_cf49 = _1126___mcc_h20;
          let _1130_cf48 = _1125___mcc_h19;
          let _1131_cf47 = _1124___mcc_h18;
          let _1132_cf46 = _1123___mcc_h17;
          (globalState).f13 = (_module.__default.safeModuloInt((_this).f48, (_this).f34)).minus((_this).f34);
          let _1133_v70;
          _1133_v70 = _module.D3.create_DC8(_dafny.MultiSet.fromElements(_1003_v0), (_this).f47, new BigNumber(735));
          let _pat_let_tv11 = _1129_cf49;
          _1133_v70 = function (_pat_let15_0) {
            return function (_1134_dt__update__tmp_h1) {
              return function (_pat_let16_0) {
                return function (_1135_dt__update_hcf16_h0) {
                  return _module.D3.create_DC8((_1134_dt__update__tmp_h1).dtor_cf15, _1135_dt__update_hcf16_h0, (_1134_dt__update__tmp_h1).dtor_cf17);
                }(_pat_let16_0);
              }(_pat_let_tv11);
            }(_pat_let15_0);
          }(_1133_v70);
          let _1136_v71;
          let _init23 = function (_1137_i8) {
            return _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_this.f31);
          };
          let _nw184 = Array((_dafny.ONE).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw184.length); _i0_23++) {
            _nw184[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1136_v71 = _nw184;
          let _1138_v72;
          _1138_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_1129_cf49);
          let _index141 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1136_v71).length));
          (_1136_v71)[_index141] = (_1138_v72).Merge(_1138_v72);
          let _index142 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1136_v71).length));
          (_1136_v71)[_index142] = ((((_this.f29) ? (true) : (((_this).f32)[_module.__default.safeIndex((_this).f48, new BigNumber(((_this).f32).length))]))) ? (_1138_v72) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f47)));
          let _1139_v74;
          _1139_v74 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of ((_1136_v71)[_module.__default.safeIndex(new BigNumber(224), new BigNumber((_1136_v71).length))]).Keys.Elements) {
              let _1140_v73 = _compr_29;
              if (((_1136_v71)[_module.__default.safeIndex(new BigNumber(224), new BigNumber((_1136_v71).length))]).contains(_1140_v73)) {
                _coll29.add((_1140_v73).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("hmcvtul")).length)));
              }
            }
            return _coll29;
          }()).length)), (_this).f30);
          let _1141_v76;
          _1141_v76 = _dafny.Seq.of(_1139_v74);
          let _1142_v77;
          let _nw185 = Array((new BigNumber(13)).toNumber());
          _nw185[(_dafny.ZERO).toNumber()] = _1139_v74;
          _nw185[(_dafny.ONE).toNumber()] = (_1139_v74).Difference(_dafny.Set.fromElements((_this).f34, (_this).f34));
          _nw185[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f30), (_this).f30);
          _nw185[(new BigNumber(3)).toNumber()] = function () {
            let _coll30 = new _dafny.Set();
            for (const _compr_30 of _dafny.IntegerRange(new BigNumber(718), new BigNumber(-585))) {
              let _1143_v75 = _compr_30;
              if (((new BigNumber(718)).isLessThanOrEqualTo(_1143_v75)) && ((_1143_v75).isLessThan(new BigNumber(-585)))) {
                _coll30.add(_module.__default.safeModuloInt(_1143_v75, new BigNumber(470)));
              }
            }
            return _coll30;
          }();
          _nw185[(new BigNumber(4)).toNumber()] = _1139_v74;
          _nw185[(new BigNumber(5)).toNumber()] = (_1139_v74).Difference(_1139_v74);
          _nw185[(new BigNumber(6)).toNumber()] = (_1139_v74).Difference(_1139_v74);
          _nw185[(new BigNumber(7)).toNumber()] = (_1141_v76)[_module.__default.safeIndex(new BigNumber((_1128_cf50).length), new BigNumber((_1141_v76).length))];
          _nw185[(new BigNumber(8)).toNumber()] = _1139_v74;
          _nw185[(new BigNumber(9)).toNumber()] = _1139_v74;
          _nw185[(new BigNumber(10)).toNumber()] = _1139_v74;
          _nw185[(new BigNumber(11)).toNumber()] = _1139_v74;
          _nw185[(new BigNumber(12)).toNumber()] = _1139_v74;
          _1142_v77 = _nw185;
          let _index143 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1142_v77).length));
          (_1142_v77)[_index143] = _dafny.Set.fromElements(_1131_cf47, (_this).f30);
          let _index144 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1142_v77).length));
          (_1142_v77)[_index144] = (_this).fm4(_1129_cf49, function () {
            let _coll31 = new _dafny.Set();
            for (const _compr_31 of _dafny.IntegerRange(new BigNumber(-346), new BigNumber(278))) {
              let _1144_v78 = _compr_31;
              if (((new BigNumber(-346)).isLessThanOrEqualTo(_1144_v78)) && ((_1144_v78).isLessThan(new BigNumber(278)))) {
                _coll31.add(_module.__default.safeDivisionInt(_1144_v78, _1131_cf47));
              }
            }
            return _coll31;
          }(), (_this).f34, _1132_cf46, globalState);
        } else if (_source19.is_DC21) {
          let _1145___mcc_h22 = (_source19).cf40;
          let _1146_cf40 = _1145___mcc_h22;
          let _1147_v79;
          _1147_v79 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_1005_v2).Union(_dafny.MultiSet.fromElements((_this).f30)));
          let _rhs147 = (((_1147_v79).contains((_this.f28) === (_this.f29))) ? ((_1147_v79).get((_this.f28) === (_this.f29))) : ((_1005_v2).update((_this).f48, _module.__default.abs(new BigNumber((_dafny.Seq.UnicodeFromString("gonoudi")).length)))));
          let _rhs148 = (_this).f47;
          let _rhs149 = ((_this).f48).minus((_this).f30);
          let _lhs110 = globalState;
          let _lhs111 = globalState;
          _1005_v2 = _rhs147;
          _lhs110.f12 = _rhs148;
          _lhs111.f13 = _rhs149;
          let _1148_v80;
          let _nw186 = Array((new BigNumber(6)).toNumber());
          _nw186[(_dafny.ZERO).toNumber()] = false;
          _nw186[(_dafny.ONE).toNumber()] = ((_this).f47) === (_this.f29);
          _nw186[(new BigNumber(2)).toNumber()] = ((_this).f48).isLessThan((_this).f30);
          _nw186[(new BigNumber(3)).toNumber()] = !(_this.f28);
          _nw186[(new BigNumber(4)).toNumber()] = true;
          _nw186[(new BigNumber(5)).toNumber()] = _this.f28;
          _1148_v80 = _nw186;
          let _index145 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1148_v80).length));
          (_1148_v80)[_index145] = ((_this.f31) ? (_this.f31) : (_this.f29));
          let _index146 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1148_v80).length));
          (_1148_v80)[_index146] = false;
          let _1149_v81;
          _1149_v81 = _dafny.MultiSet.fromElements((_1012_v8)[_module.__default.safeIndex(new BigNumber(423), new BigNumber((_1012_v8).length))]);
          let _index147 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1148_v80).length));
          let _rhs150 = _1148_v80;
          let _rhs151 = !((new BigNumber((_1011_v7).length)).isLessThanOrEqualTo((new BigNumber((function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of (_1008_v5).Elements) {
              let _1150_v82 = _compr_32;
              if (_dafny.Seq.contains(_1008_v5, _1150_v82)) {
                _coll32.push([_module.__default.safeModuloInt(_1150_v82, (_this).f30),_1093_v54]);
              }
            }
            return _coll32;
          }()).length)).minus(new BigNumber((_1093_v54).length))));
          let _rhs152 = _dafny.MultiSet.fromElements(_module.D4.create_DC11((_this).f30, (_1148_v80)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_1148_v80).length))], (_this).f30));
          let _lhs112 = _1148_v80;
          let _lhs113 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1148_v80).length));
          _1148_v80 = _rhs150;
          _lhs112[_lhs113] = _rhs151;
          _1149_v81 = _rhs152;
          (_this).f29 = false;
        } else {
          let _1151___mcc_h23 = (_source19).cf51;
          let _1152_cf51 = _1151___mcc_h23;
          (_this).f35 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_1011_v7, (_this).f32), _dafny.Seq.Concat((_this).fm5((_this).f48, (_this).f48, (_this).f47, (_this).f48, globalState), (_this).f32));
          let _1153_v83;
          let _nw187 = new _module.C1();
          _nw187.__ctor(_dafny.Seq.IsPrefixOf(_1093_v54, _dafny.Seq.Create(_module.__default.abs(new BigNumber(314)), function (_1154_i9) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          })), ((_this).f34).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("tsm")).length)), _this.f31);
          _1153_v83 = _nw187;
          (globalState).f26 = (_1153_v83).f43;
          let _1155_v84;
          let _nw188 = Array((new BigNumber(29)).toNumber()).fill(false);
          _1155_v84 = _nw188;
          (globalState).f21 = _1155_v84;
        }
        let _rhs153 = _module.__default.fm39(_this.f28, _1005_v2, !(_this.f31), globalState);
        let _rhs154 = ((_this).f30).isLessThan((_this).f48);
        let _rhs155 = (new BigNumber(358)).isEqualTo((((_1017_v11).contains(_this.f28)) ? ((_1017_v11).get(_this.f28)) : (new BigNumber((_dafny.Seq.update(_1008_v5, _module.__default.safeIndex((_this).f34, new BigNumber((_1008_v5).length)), (_this).f48)).length))));
        let _lhs114 = globalState;
        let _lhs115 = _this;
        let _lhs116 = globalState;
        _lhs114.f13 = _rhs153;
        _lhs115.f29 = _rhs154;
        _lhs116.f12 = _rhs155;
      }
      let _hi5 = (_dafny.ZERO).minus(new BigNumber(-182));
      for (let _1156_i10 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f34,_1005_v2)).length)).minus((_dafny.ZERO).minus((_this).f48)); _1156_i10.isLessThan(_hi5); _1156_i10 = _1156_i10.plus(_dafny.ONE)) {
        (globalState).f12 = _module.__default.fm1(_this.f35, new BigNumber((_1005_v2).cardinality()), globalState);
        (globalState).f12 = (!(new BigNumber(169)).isEqualTo(_1156_i10)) && (_this.f35);
        (globalState).f13 = ((_this).f34).multipliedBy((_this).f48);
        let _1157_v85;
        _1157_v85 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_1017_v11);
        _1157_v85 = (_1157_v85).update(_this.f28, _dafny.MultiSet.fromElements(_this.f31));
      }
      let _1158_v86;
      _1158_v86 = _dafny.Seq.UnicodeFromString("jjtojxk");
      r0 = ((_this).f34).isLessThan(new BigNumber((_1158_v86).length));
      r1 = ((!(_this.f28)) ? (((_this).f34).plus((_this).f30)) : ((_dafny.ZERO).minus((_this).f34)));
      r2 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("l"), _1158_v86), _module.__default.safeIndex(new BigNumber(213), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("l"), _1158_v86)).length)), _1003_v0);
      return [r0, r1, r2];
    }
    m14(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = new _dafny.CodePoint('D'.codePointAt(0));
      let _hi6 = new BigNumber((_dafny.Seq.UnicodeFromString("bxgmww")).length);
      for (let _1159_i0 = (_this).f30; _1159_i0.isLessThan(_hi6); _1159_i0 = _1159_i0.plus(_dafny.ONE)) {
        (globalState).f12 = _this.f28;
        (globalState).f13 = _module.__default.safeDivisionInt(_1159_i0, (_1159_i0).minus(_1159_i0));
        let _1160_v0;
        _1160_v0 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(73),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(137)), ((_1161_p0) => function (_1162_i1) {
          return _1161_p0;
        })(p0))).length));
        let _1163_v1;
        let _nw189 = new _module.C4();
        _nw189.__ctor((_this).f30, _this.f29, (_this).f32, (_this).f33, new BigNumber((_1160_v0).length), _this.f35, _module.__default.fm1(false, (_this).f48, globalState), (_this).f47);
        _1163_v1 = _nw189;
        let _1164_v2;
        _1164_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1159_i0,_1163_v1);
        let _1165_v3;
        _1165_v3 = _module.D10.create_DC30(_module.__default.fm41(_this.f31, p0, _this.f31, (_this).f34, globalState), _dafny.MultiSet.fromElements(_1163_v1, _1163_v1, (((_1164_v2).contains(new BigNumber(764))) ? ((_1164_v2).get(new BigNumber(764))) : (_1163_v1)), _1163_v1));
        let _1166_v4;
        _1166_v4 = _dafny.MultiSet.fromElements(_1163_v1, _1163_v1, _1163_v1, _1163_v1, _1163_v1);
        let _pat_let_tv12 = _1166_v4;
        let _source20 = function (_pat_let17_0) {
          return function (_1167_dt__update__tmp_h0) {
            return function (_pat_let18_0) {
              return function (_1168_dt__update_hcf59_h0) {
                return _module.D10.create_DC30((_1167_dt__update__tmp_h0).dtor_cf58, _1168_dt__update_hcf59_h0);
              }(_pat_let18_0);
            }(_pat_let_tv12);
          }(_pat_let17_0);
        }(_1165_v3);
        if (_source20.is_DC29) {
          let _1169___mcc_h0 = (_source20).cf56;
          let _1170___mcc_h1 = (_source20).cf57;
          let _1171_cf57 = _1170___mcc_h1;
          let _1172_cf56 = _1169___mcc_h0;
          (globalState).f13 = (_1172_cf56).minus(((_this).f34).multipliedBy((_this).f34));
          let _1173_v5;
          _1173_v5 = _dafny.Set.fromElements(p0, p0, p0);
          let _index148 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((p1).length));
          (p1)[_index148] = !(_dafny.Map.Empty.slice().updateUnsafe(_1163_v1.f31,_1173_v5)).contains(_this.f31);
          let _index149 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((p1).length));
          (p1)[_index149] = _1163_v1.f29;
          let _1174_v6;
          _1174_v6 = _dafny.Seq.UnicodeFromString("xji");
          let _1175_v7;
          _1175_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1174_v6,(_this).f30);
          _1175_v7 = _1175_v7;
          let _1176_v8;
          let _init24 = ((_1177_p0) => function (_1178_i2) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(315)), ((_1179_p0) => function (_1180_i3) {
              return _1179_p0;
            })(_1177_p0));
          })(p0);
          let _nw190 = Array((new BigNumber(29)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw190.length); _i0_24++) {
            _nw190[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1176_v8 = _nw190;
          let _index150 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1176_v8).length));
          (_1176_v8)[_index150] = _1174_v6;
          let _index151 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1176_v8).length));
          (_1176_v8)[_index151] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), ((_1181_p0) => function (_1182_i4) {
            return _1181_p0;
          })(p0));
        } else if (_source20.is_DC30) {
          let _1183___mcc_h2 = (_source20).cf58;
          let _1184___mcc_h3 = (_source20).cf59;
          let _1185_cf59 = _1184___mcc_h3;
          let _1186_cf58 = _1183___mcc_h2;
          (globalState).f13 = (_this).f48;
          let _1187_v9;
          let _nw191 = Array((new BigNumber(4)).toNumber());
          _nw191[(_dafny.ZERO).toNumber()] = (_1163_v1).f34;
          _nw191[(_dafny.ONE).toNumber()] = (_1163_v1).f34;
          _nw191[(new BigNumber(2)).toNumber()] = (_this).f34;
          _nw191[(new BigNumber(3)).toNumber()] = (_1163_v1).f34;
          _1187_v9 = _nw191;
          let _1188_v10;
          _1188_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1187_v9,_1163_v1.f28);
          let _1189_v11;
          _1189_v11 = _module.D7.create_DC19(_1163_v1.f31, (_dafny.ZERO).minus((_1163_v1).f34), new BigNumber((_1188_v10).length));
          (globalState).f13 = (_1189_v11).dtor_cf38;
          let _1190_v12;
          _1190_v12 = _dafny.Seq.UnicodeFromString("dplmqrvf");
          let _1191_v13;
          _1191_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f48,_1190_v12);
          let _1192_v14;
          _1192_v14 = _dafny.Set.fromElements(_this.f31);
          let _1193_v16;
          _1193_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_this.f28);
          let _1194_v17;
          _1194_v17 = _dafny.Seq.of(new BigNumber(((_1163_v1).f32).length), new BigNumber((_1193_v16).length));
          let _1195_v18;
          let _nw192 = Array((new BigNumber(26)).toNumber());
          _nw192[(_dafny.ZERO).toNumber()] = (_1163_v1).f32;
          _nw192[(_dafny.ONE).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(2)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(3)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(4)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(5)).toNumber()] = (_1163_v1).fm5(new BigNumber(((((_1191_v13).contains((_this).f48)) ? ((_1191_v13).get((_this).f48)) : (_1190_v12))).length), (_this).f34, (_this).f47, (_1163_v1).f30, globalState);
          _nw192[(new BigNumber(6)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(7)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(8)).toNumber()] = (_this).f32;
          _nw192[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_1163_v1.f28, _1163_v1.f29);
          _nw192[(new BigNumber(10)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(11)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(12)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(13)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(14)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(15)).toNumber()] = (_1163_v1).fm5(new BigNumber((_1192_v14).length), (_this).f30, _1163_v1.f29, new BigNumber((function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of (_1194_v17).Elements) {
              let _1196_v15 = _compr_33;
              if (_dafny.Seq.contains(_1194_v17, _1196_v15)) {
                _coll33.push([_module.__default.safeModuloInt(_1196_v15, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f30)).length)),_1163_v1.f28]);
              }
            }
            return _coll33;
          }()).length), globalState);
          _nw192[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(_this.f29);
          _nw192[(new BigNumber(17)).toNumber()] = (_this).f32;
          _nw192[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(_1163_v1.f35);
          _nw192[(new BigNumber(19)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(20)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(21)).toNumber()] = (_this).f32;
          _nw192[(new BigNumber(22)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(23)).toNumber()] = (_1163_v1).f32;
          _nw192[(new BigNumber(24)).toNumber()] = _dafny.Seq.of((_this).f47, _1163_v1.f35, _this.f28);
          _nw192[(new BigNumber(25)).toNumber()] = (_1163_v1).f32;
          _1195_v18 = _nw192;
          let _1197_v19;
          let _nw193 = new _module.C3();
          _nw193.__ctor(new BigNumber(893), _1195_v18, p0, (_module.__default.fm46(new BigNumber((_1190_v12).length), (_this).f34, _this.f31, globalState)).multipliedBy((_1163_v1).f34), !(!(((_this).f32)[_module.__default.safeIndex((_1163_v1).f30, new BigNumber(((_this).f32).length))])), _dafny.Seq.Concat(_dafny.Seq.of(false), (_1163_v1).f32), _dafny.Seq.Concat((_1163_v1).f33, (_this).f33), (_1163_v1).f30, !(((_1163_v1.f35) ? (_1163_v1.f31) : (_this.f35))), ((_1163_v1.f31) ? (true) : (false)), true);
          _1197_v19 = _nw193;
          _1197_v19 = _1197_v19;
          let _1198_v20;
          _1198_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), ((_1199_p0) => function (_1200_i5) {
            return _1199_p0;
          })(p0)),_1192_v14);
          _1198_v20 = (_1198_v20).update(_dafny.Seq.Concat(_1190_v12, _dafny.Seq.update((_this).fm3((_this).f47, _1163_v1.f28, globalState), _module.__default.safeIndex((((_1160_v0).contains(new BigNumber((_1160_v0).length))) ? ((_1160_v0).get(new BigNumber((_1160_v0).length))) : (new BigNumber(((_this).f32).length))), new BigNumber(((_this).fm3((_this).f47, _1163_v1.f28, globalState)).length)), _1197_v19.f36)), _1192_v14);
        } else if (_source20.is_DC31) {
          let _1201___mcc_h4 = (_source20).cf60;
          let _1202___mcc_h5 = (_source20).cf61;
          let _1203___mcc_h6 = (_source20).cf62;
          let _1204_cf62 = _1203___mcc_h6;
          let _1205_cf61 = _1202___mcc_h5;
          let _1206_cf60 = _1201___mcc_h4;
          let _index152 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((p1).length));
          (p1)[_index152] = ((_1163_v1).f32)[_module.__default.safeIndex(new BigNumber(959), new BigNumber(((_1163_v1).f32).length))];
          let _1207_v21;
          _1207_v21 = _module.D7.create_DC19((_this).f47, (_1163_v1).f34, (_1163_v1).f34);
          let _index153 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((p1).length));
          (p1)[_index153] = (_1207_v21).dtor_cf36;
          let _1208_v22;
          let _nw194 = Array((new BigNumber(21)).toNumber()).fill([]);
          _1208_v22 = _nw194;
          _1208_v22 = _1208_v22;
          let _1209_v23;
          let _nw195 = new _module.C1();
          _nw195.__ctor(false, (p1)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((p1).length))], false);
          _1209_v23 = _nw195;
          let _1210_v24;
          _1210_v24 = _module.D10.create_DC28(_1209_v23);
          let _1211_v25;
          _1211_v25 = _dafny.MultiSet.fromElements(_1210_v24);
          let _1212_v26;
          _1212_v26 = _dafny.Set.fromElements((_1211_v25).Difference(_1211_v25));
          _1212_v26 = _1212_v26;
          (globalState).f13 = _1159_i0;
        } else if (_source20.is_DC28) {
          let _1213___mcc_h7 = (_source20).cf55;
          let _1214_cf55 = _1213___mcc_h7;
          r1 = p0;
          let _1215_v27;
          let _nw196 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1215_v27 = _nw196;
          let _1216_v28;
          _1216_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1215_v27,(_this.f28) || (_this.f31));
          _1216_v28 = (_1216_v28);
          let _nw197 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _1215_v27 = _nw197;
          (globalState).f12 = !((_1214_cf55).f43);
        } else {
          let _1217___mcc_h8 = (_source20).cf63;
          let _1218_cf63 = _1217___mcc_h8;
          let _1219_v29;
          _1219_v29 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f28);
          let _1220_v30;
          _1220_v30 = _dafny.MultiSet.fromElements(((_this).f48).isLessThanOrEqualTo((_1163_v1).f34), (((_1219_v29).contains(_1163_v1.f28)) ? ((_1219_v29).get(_1163_v1.f28)) : (_1163_v1.f28)), ((_1163_v1.f35) ? (_this.f28) : (_1163_v1.f28)), _1163_v1.f28);
          (globalState).f13 = new BigNumber((_1220_v30).cardinality());
          (_1163_v1).f35 = _1163_v1.f35;
          r1 = p0;
          (globalState).f13 = _module.__default.fm46((_1163_v1).f30, ((_this.f31) ? ((_this).f48) : ((_1163_v1).f30)), _1163_v1.f28, globalState);
        }
        let _1221_v31;
        _1221_v31 = _dafny.MultiSet.fromElements(false);
        if ((((_1221_v31).IsSubsetOf(_dafny.MultiSet.fromElements(_this.f35, (_this).f47))) ? (_dafny.areEqual((_this).f32, (_this).f32)) : (true))) {
          (globalState).f13 = (_this).f30;
          let _1222_v32;
          _1222_v32 = _dafny.Seq.of((_1163_v1).f34, _1159_i0, (_1163_v1).f34);
          let _1223_v33;
          _1223_v33 = _module.D8.create_DC24(_1163_v1.f35, (_1163_v1).f30, (_this).f32, _1163_v1.f31, _dafny.Seq.update(_1222_v32, _module.__default.safeIndex((_this).f48, new BigNumber((_1222_v32).length)), (_this).f48));
          let _1224_v34;
          _1224_v34 = _module.D0.create_DC0(_1163_v1.f31);
          let _1225_v35;
          let _nw198 = new _module.C2();
          _nw198.__ctor(new BigNumber((_1160_v0).length), (_1223_v33).dtor_cf47, (_1224_v34).dtor_cf0, !(_1163_v1.f29), _this.f29);
          _1225_v35 = _nw198;
          let _1226_v36;
          _1226_v36 = _dafny.Map.Empty.slice().updateUnsafe(!((_1163_v1.f29) === (_this.f35)),_1163_v1.f35);
          _1226_v36 = (_1226_v36).update((_this).f47, (_this).f47);
          let _1227_v37;
          let _nw199 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
          _1227_v37 = _nw199;
          let _1228_v38;
          _1228_v38 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f34);
          let _1229_v39;
          _1229_v39 = _dafny.Set.fromElements((_1163_v1).f30, (_1163_v1).f34, new BigNumber((_1228_v38).length), (_this).f48, (_1225_v35).f46);
          let _1230_v41;
          _1230_v41 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll34 = new _dafny.Set();
            for (const _compr_34 of _dafny.IntegerRange(new BigNumber(127), new BigNumber(328))) {
              let _1231_v40 = _compr_34;
              if (((new BigNumber(127)).isLessThanOrEqualTo(_1231_v40)) && ((_1231_v40).isLessThan(new BigNumber(328)))) {
                _coll34.add(_module.__default.safeModuloInt(_1231_v40, (_1163_v1).f34));
              }
            }
            return _coll34;
          }(),p1);
          let _index154 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1227_v37).length));
          (_1227_v37)[_index154] = (_dafny.Map.Empty.slice().updateUnsafe(_1229_v39,p1)).Merge(_1230_v41);
          let _index155 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1227_v37).length));
          (_1227_v37)[_index155] = _1230_v41;
          let _1232_v42;
          let _nw200 = new _module.C1();
          _nw200.__ctor(true, (((_1226_v36).contains(_this.f35)) ? ((_1226_v36).get(_this.f35)) : (_this.f29)), _1163_v1.f35);
          _1232_v42 = _nw200;
          let _1233_v43;
          _1233_v43 = _dafny.Map.Empty.slice().updateUnsafe((_1163_v1).f34,_module.D10.create_DC28(_1232_v42));
          let _1234_v44;
          _1234_v44 = _dafny.Seq.UnicodeFromString("mo");
          (_1163_v1).f29 = (new BigNumber((_1233_v43).length)).isLessThanOrEqualTo(new BigNumber((_1234_v44).length));
        } else {
          (globalState).f26 = !(_1163_v1.f29);
          let _1235_v45;
          _1235_v45 = _dafny.Seq.UnicodeFromString("sqrfmscn");
          let _index156 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((p1).length));
          (p1)[_index156] = !_dafny.areEqual(_1235_v45, _1235_v45);
          let _index157 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((p1).length));
          (p1)[_index157] = _1163_v1.f31;
          let _1236_v46;
          _1236_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1221_v31,p0);
          _1236_v46 = (_1236_v46).update((_1221_v31).update(!(_1163_v1.f31), _module.__default.abs((_1163_v1).f30)), (_1235_v45)[_module.__default.safeIndex((_1163_v1).f34, new BigNumber((_1235_v45).length))]);
          (globalState).f26 = _dafny.Seq.contains((_this).f32, _module.__default.fm1(_1163_v1.f28, (_1163_v1).f34, globalState));
          (_this).f35 = _1163_v1.f28;
        }
      }
      if (!(((_this).f32)[_module.__default.safeIndex(new BigNumber(-172), new BigNumber(((_this).f32).length))])) {
        r1 = p0;
        (globalState).f13 = (new BigNumber(560)).multipliedBy((_this).f34);
        let _1237_v47;
        _1237_v47 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f32);
        let _1238_v48;
        _1238_v48 = _module.D3.create_DC9(_this.f35, (_this).f47, p1, _this.f28);
        let _1239_v49;
        let _nw201 = Array((new BigNumber(21)).toNumber());
        _nw201[(_dafny.ZERO).toNumber()] = (_this).f32;
        _nw201[(_dafny.ONE).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(2)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(3)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(4)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(5)).toNumber()] = (((_1237_v47).contains((_1238_v48).dtor_cf21)) ? ((_1237_v47).get((_1238_v48).dtor_cf21)) : ((_this).f32));
        _nw201[(new BigNumber(6)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(7)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(false, (_this).f47, _this.f35, !(true));
        _nw201[(new BigNumber(9)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(10)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(11)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(12)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(true, _this.f31);
        _nw201[(new BigNumber(14)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(15)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(16)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(17)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(true, _this.f28, _this.f28);
        _nw201[(new BigNumber(19)).toNumber()] = (_this).f32;
        _nw201[(new BigNumber(20)).toNumber()] = _dafny.Seq.of((_this).f47);
        _1239_v49 = _nw201;
        let _1240_v50;
        let _nw202 = new _module.C3();
        _nw202.__ctor(_module.__default.safeDivisionInt((_this).f30, (_this).f30), _1239_v49, p0, ((_this).f48).multipliedBy((_this).f30), ((_this).f47) === (_this.f28), (_this).f32, (_this).f33, (_this).f30, false, (_this).f47, _this.f35);
        _1240_v50 = _nw202;
        (globalState).f12 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), function (_1241_i6) {
          return (_this).f30;
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(685)), function (_1242_i7) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("oaqupmrc")).length);
        }));
        let _1243_v51;
        _1243_v51 = _dafny.Seq.of((_1240_v50).f44);
        let _1244_v52;
        _1244_v52 = _dafny.MultiSet.fromElements(_1243_v51, _dafny.Seq.Create(_module.__default.abs(new BigNumber(618)), ((_1245_v50) => function (_1246_i8) {
          return (_1245_v50).f44;
        })(_1240_v50)), (p2).dtor_cf32);
        let _1247_v53;
        _1247_v53 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_1244_v52);
        let _1248_v54;
        _1248_v54 = _dafny.Set.fromElements((_this).f34, new BigNumber((_1243_v51).length), (_this).f34);
        _1247_v53 = (_1247_v53).update((_dafny.Set.fromElements((_this).f34)).IsDisjointFrom(_1248_v54), _1244_v52);
      } else {
        let _1249_v55;
        _1249_v55 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,((_dafny.ZERO).minus((_this).f48)).minus((_this).f30));
        let _1250_v56;
        _1250_v56 = _dafny.Seq.UnicodeFromString("axjmgmig");
        _1249_v55 = (_1249_v55).update(_dafny.Seq.IsPrefixOf(_1250_v56, _1250_v56), (_this).f30);
        let _1251_v57;
        _1251_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f48,p0);
        _1251_v57 = (_1251_v57).update((_this).f30, p0);
        let _1252_v58;
        _1252_v58 = _dafny.Seq.of((_this).f48, (_this).f48);
        let _1253_v59;
        _1253_v59 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Set.fromElements(_this.f35)).length)).minus((_this).f34),_1252_v58);
        let _1254_v60;
        _1254_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_this.f35);
        _1253_v59 = (_1253_v59).update(new BigNumber((_1254_v60).length), _1252_v58);
        (globalState).f1 = !(_this.f35);
        let _1255_v61;
        _1255_v61 = _dafny.Seq.of(_1250_v56, _dafny.Seq.UnicodeFromString("yumismnte"));
        let _1256_v63;
        _1256_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,new BigNumber(692));
        let _1257_v64;
        _1257_v64 = _dafny.MultiSet.fromElements((_this).f47);
        let _1258_v65;
        let _init25 = function (_1259_i9) {
          return (_1259_i9).plus((_module.D7.create_DC19(_this.f35, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(146)), function (_1260_i10) {
  return (_this).f30;
})).length), (_this).f34)).dtor_cf37);
        };
        let _nw203 = Array((new BigNumber(9)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw203.length); _i0_25++) {
          _nw203[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1258_v65 = _nw203;
        let _1261_v66;
        _1261_v66 = _module.D0.create_DC0(_this.f35);
        let _1262_v67;
        _1262_v67 = _dafny.Seq.of(_1261_v66);
        let _1263_v68;
        _1263_v68 = _dafny.Set.fromElements((_this).f48);
        let _1264_v69;
        _1264_v69 = _module.D2.create_DC6((_dafny.ZERO).minus((_this).f30), _this.f35, (_this).f34, (_this).f47, (_this).f34);
        let _1265_v70;
        _1265_v70 = _dafny.MultiSet.fromElements((_this).f48, (_this).f34);
        let _1266_v71;
        let _nw204 = Array((new BigNumber(29)).toNumber());
        _nw204[(_dafny.ZERO).toNumber()] = (new BigNumber(((_this).fm4(true, _dafny.Set.fromElements((_this).f48, (_this).f48), new BigNumber((_dafny.Seq.update((_1255_v61)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("kuntapvl")).length), new BigNumber((_1255_v61).length))], _module.__default.safeIndex((_this).f48, new BigNumber(((_1255_v61)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("kuntapvl")).length), new BigNumber((_1255_v61).length))]).length)), new _dafny.CodePoint('j'.codePointAt(0)))).length), _this.f28, globalState)).length)).plus(new BigNumber((function () {
          let _coll35 = new _dafny.Set();
          for (const _compr_35 of (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("wa"))).Elements) {
            let _1267_v62 = _compr_35;
            if ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("wa"))).contains(_1267_v62)) {
              _coll35.add(_1267_v62);
            }
          }
          return _coll35;
        }()).length));
        _nw204[(_dafny.ONE).toNumber()] = (((_1256_v63).contains((_this).f48)) ? ((_1256_v63).get((_this).f48)) : ((_this).f30));
        _nw204[(new BigNumber(2)).toNumber()] = ((_this).f48).minus((_this).f30);
        _nw204[(new BigNumber(3)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_1252_v58).length))).multipliedBy((((_1257_v64).contains((_module.D10.create_DC31(_this.f29, (_this).f47, (_this).f48)).dtor_cf61)) ? ((_1257_v64).get((_module.D10.create_DC31(_this.f29, (_this).f47, (_this).f48)).dtor_cf61)) : ((_this).f34)));
        _nw204[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f34, (_this).f30));
        _nw204[(new BigNumber(5)).toNumber()] = (_this).f30;
        _nw204[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt((_this).f30, (_this).f34);
        _nw204[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.of(_1258_v65, _1258_v65, _1258_v65, _1258_v65)).length);
        _nw204[(new BigNumber(8)).toNumber()] = new BigNumber((_1262_v67).length);
        _nw204[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt((_this).f30, (_this).f48);
        _nw204[(new BigNumber(10)).toNumber()] = (_this).f48;
        _nw204[(new BigNumber(11)).toNumber()] = ((_this).f34).minus((_this).f34);
        _nw204[(new BigNumber(12)).toNumber()] = new BigNumber((_1263_v68).length);
        _nw204[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f48))).length)), (_this).f30);
        _nw204[(new BigNumber(14)).toNumber()] = ((_1264_v69).dtor_cf11).multipliedBy((_dafny.ZERO).minus((_this).f30));
        _nw204[(new BigNumber(15)).toNumber()] = (_this).f34;
        _nw204[(new BigNumber(16)).toNumber()] = (_this).f34;
        _nw204[(new BigNumber(17)).toNumber()] = (_this).f48;
        _nw204[(new BigNumber(18)).toNumber()] = (_this).f30;
        _nw204[(new BigNumber(19)).toNumber()] = (_this).f30;
        _nw204[(new BigNumber(20)).toNumber()] = (_this).f30;
        _nw204[(new BigNumber(21)).toNumber()] = (_this).f30;
        _nw204[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_this).f34);
        _nw204[(new BigNumber(23)).toNumber()] = (_this).f48;
        _nw204[(new BigNumber(24)).toNumber()] = new BigNumber((_1263_v68).length);
        _nw204[(new BigNumber(25)).toNumber()] = new BigNumber(-628);
        _nw204[(new BigNumber(26)).toNumber()] = (((_1265_v70).contains((_this).f30)) ? ((_1265_v70).get((_this).f30)) : ((_this).f34));
        _nw204[(new BigNumber(27)).toNumber()] = ((_this).f48).multipliedBy(_module.__default.fm39((_this).f47, _1265_v70, _this.f31, globalState));
        _nw204[(new BigNumber(28)).toNumber()] = new BigNumber((_1249_v55).length);
        _1266_v71 = _nw204;
        let _index158 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1266_v71).length));
        (_1266_v71)[_index158] = (_this).f48;
        let _1268_v72;
        _1268_v72 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_1258_v65);
        let _index159 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1266_v71).length));
        let _rhs156 = new BigNumber(808);
        let _rhs157 = (new BigNumber((_1268_v72).length)).multipliedBy(_module.__default.fm39((_this).f47, _1265_v70, _this.f31, globalState));
        let _rhs158 = (_dafny.ZERO).minus((_this).f34);
        let _lhs117 = _1266_v71;
        let _lhs118 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_1266_v71).length));
        let _lhs119 = globalState;
        let _lhs120 = globalState;
        _lhs117[_lhs118] = _rhs156;
        _lhs119.f13 = _rhs157;
        _lhs120.f13 = _rhs158;
      }
      let _1269_v73;
      _1269_v73 = _module.D6.create_DC16(p1);
      let _1270_v74;
      _1270_v74 = _module.D6.create_DC17(_1269_v73);
      _1270_v74 = _1270_v74;
      (_this).f28 = function (_source21) {
        if (_source21.is_DC16) {
          let _1271___mcc_h9 = (_source21).cf33;
          let _1272_cf33 = _1271___mcc_h9;
          return (_this).f47;
        } else if (_source21.is_DC15) {
          let _1273___mcc_h10 = (_source21).cf32;
          let _1274_cf32 = _1273___mcc_h10;
          return false;
        } else {
          let _1275___mcc_h11 = (_source21).cf34;
          let _1276_cf34 = _1275___mcc_h11;
          return (_this).f47;
        }
      }(p2);
      let _1277_v75;
      _1277_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).f48,_this.f35);
      _1277_v75 = _1277_v75;
      let _1278_v76;
      _1278_v76 = _dafny.Set.fromElements(p1);
      let _1279_v77;
      _1279_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,(_this).f34);
      let _1280_v78;
      _1280_v78 = _dafny.Map.Empty.slice().updateUnsafe(!(_1278_v76).contains(p1),_1279_v77);
      _1280_v78 = (_module.__default.fm50(_1277_v75, globalState)).Merge(_1280_v78);
      let _1281_v79;
      _1281_v79 = _dafny.Set.fromElements((_this).f47, true);
      let _1282_v80;
      _1282_v80 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_this.f31, (_this).f30, globalState),(_this).f30);
      let _1283_v81;
      _1283_v81 = _dafny.MultiSet.fromElements((((_1282_v80).contains((_this).f47)) ? ((_1282_v80).get((_this).f47)) : ((_this).f30)), (_this).f30, (_this).f34, (_dafny.ZERO).minus((_this).f30), (_this).f30);
      r0 = (new BigNumber((_1281_v79).length)).isEqualTo(new BigNumber((_1283_v81).cardinality()));
      let _1284_v82;
      _1284_v82 = _dafny.Seq.of(p0, p0);
      r1 = (_1284_v82)[_module.__default.safeIndex((_this).f48, new BigNumber((_1284_v82).length))];
      return [r0, r1];
    }
    get f47() {
      let _this = this;
      return _this._f47;
    };
    get f48() {
      let _this = this;
      return _this._f48;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f28 = false;
      this._f29 = false;
      this._f31 = false;
      this._f35 = false;
      this._f36 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f34 = _dafny.ZERO;
      this._f32 = _dafny.Seq.of();
      this._f33 = _dafny.Seq.of();
      this._f30 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T4, _module.T2, _module.T3, _module.T1, _module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    set f31(value) {
      let _this = this;
      _this._f31 = value;
    };
    get f35() {
      let _this = this;
      return _this._f35;
    };
    set f35(value) {
      let _this = this;
      _this._f35 = value;
    };
    get f36() {
      let _this = this;
      return _this._f36;
    };
    set f36(value) {
      let _this = this;
      _this._f36 = value;
    };
    get f34() {
      let _this = this;
      return _this._f34;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    __ctor(f36, f34, f35, f32, f33, f30, f31, f28, f29) {
      let _this = this;
      (_this)._f36 = f36;
      (_this)._f34 = f34;
      (_this)._f35 = f35;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm7(p0, p1, globalState) {
      let _this = this;
      return _this.f31;
    };
    fm6(globalState) {
      let _this = this;
      return _dafny.areEqual(_dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0))));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((function () {
        let _coll36 = new _dafny.Set();
        for (const _compr_36 of _dafny.IntegerRange(new BigNumber(924), new BigNumber(550))) {
          let _1285_v0 = _compr_36;
          if (((new BigNumber(924)).isLessThanOrEqualTo(_1285_v0)) && ((_1285_v0).isLessThan(new BigNumber(550)))) {
            _coll36.add((_1285_v0).multipliedBy(new BigNumber(((_this).f32).length)));
          }
        }
        return _coll36;
      }()).Difference(_dafny.Set.fromElements(new BigNumber(-281)))).Intersect(_dafny.Set.fromElements(new BigNumber(543), new BigNumber((_dafny.Seq.of((_this).f34, (_this).f34, (_this).f30)).length), (_this).f30, (_this).f30));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_this).f33)[_module.__default.safeIndex(((false) ? ((_this).f30) : ((_this).f30)), new BigNumber(((_this).f33).length))];
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("crmrunqsx");
    };
    fm2(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f30);
    };
    fm27(p0, p1, p2, globalState) {
      let _this = this;
      if (((_this.f29) ? (false) : (_this.f29))) {
        return (_dafny.ZERO).minus((_this).f30);
      } else {
        return (_this).f34;
      }
    };
    fm28(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(-675);
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _hi7 = (_this).f34;
      for (let _1286_i0 = (_this).f30; _1286_i0.isLessThan(_hi7); _1286_i0 = _1286_i0.plus(_dafny.ONE)) {
        let _1287_v0;
        let _nw205 = Array((new BigNumber(19)).toNumber()).fill([]);
        _1287_v0 = _nw205;
        let _1288_v1;
        _1288_v1 = _dafny.Seq.of(_1287_v0);
        let _1289_v2;
        _1289_v2 = _dafny.Map.Empty.slice().updateUnsafe((_1286_i0).multipliedBy(_1286_i0),(_1288_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber(380))).length), new BigNumber((_1288_v1).length))]);
        let _1290_v3;
        _1290_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_1287_v0);
        _1289_v2 = (_1289_v2).update(_module.__default.safeDivisionInt(_1286_i0, new BigNumber(-558)), (((_1290_v3).contains(_this.f28)) ? ((_1290_v3).get(_this.f28)) : (_1287_v0)));
        let _1291_v4;
        _1291_v4 = _dafny.Seq.UnicodeFromString("xku");
        let _1292_v5;
        _1292_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-301),_1291_v4);
        if (!(_1286_i0).isEqualTo(new BigNumber((_1292_v5).length))) {
          let _1293_v6;
          _1293_v6 = _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(476), _1286_i0));
          let _1294_v7;
          _1294_v7 = _dafny.Map.Empty.slice().updateUnsafe((_1293_v6)[_module.__default.safeIndex((_this).f30, new BigNumber((_1293_v6).length))],_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_this).f34)).length)));
          let _1295_v8;
          _1295_v8 = _dafny.Set.fromElements(new BigNumber(12), _1286_i0, _1286_i0);
          let _1296_v9;
          let _nw206 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _1296_v9 = _nw206;
          let _1297_v10;
          _1297_v10 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_this.f28);
          let _1298_v11;
          let _nw207 = new _module.C3();
          _nw207.__ctor(_1286_i0, _1296_v9, (_1291_v4)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f34)).length), new BigNumber((_1291_v4).length))], (_this).f34, _this.f29, (_this).f32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(658)), function (_1299_i1) {
            return (_this).f32;
          }), new BigNumber((_dafny.Seq.of(new BigNumber((_1297_v10).length), new BigNumber(387))).length), _this.f29, _this.f31, _this.f35);
          _1298_v11 = _nw207;
          let _1300_v12;
          _1300_v12 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("yor")).length), (_this).f30, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f34,_1298_v11)).length));
          _1294_v7 = (_1294_v7).update(((_this.f29) ? (_1295_v8) : (_1295_v8)), _dafny.Seq.Concat(_1300_v12, _1300_v12));
          let _1301_v13;
          _1301_v13 = _dafny.Seq.of(_1296_v9, _1296_v9);
          let _1302_v14;
          _1302_v14 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,new _dafny.CodePoint('i'.codePointAt(0)));
          let _1303_v15;
          let _nw208 = new _module.C3();
          _nw208.__ctor(_1286_i0, (_1301_v13)[_module.__default.safeIndex((_this).f34, new BigNumber((_1301_v13).length))], (((_1302_v14).contains(_1298_v11.f36)) ? ((_1302_v14).get(_1298_v11.f36)) : (new _dafny.CodePoint('u'.codePointAt(0)))), new BigNumber(35), _this.f28, _dafny.Seq.of(_1298_v11.f29, _1298_v11.f28), (_this).f33, _1286_i0, (!(_1298_v11.f31)) || (_this.f28), true, _this.f31);
          _1303_v15 = _nw208;
          let _1304_v16;
          let _nw209 = Array((new BigNumber(10)).toNumber());
          _1304_v16 = _nw209;
          let _1305_v17;
          _1305_v17 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("twqhca"));
          let _1306_v18;
          _1306_v18 = _dafny.Map.Empty.slice().updateUnsafe(!((_1305_v17).contains(_1291_v4)),_1304_v16);
          let _rhs159 = ((_this).f30).multipliedBy(((_1298_v11.f29) ? ((_1298_v11).f34) : ((_1303_v15).f44)));
          let _rhs160 = (((_1306_v18).contains(!(_1298_v11.f28))) ? ((_1306_v18).get(!(_1298_v11.f28))) : (_1304_v16));
          let _rhs161 = !(!(!(_dafny.Seq.contains(_1291_v4, _1298_v11.f36))));
          let _lhs121 = globalState;
          r1 = _rhs159;
          _1304_v16 = _rhs160;
          _lhs121.f26 = _rhs161;
          let _1307_v19;
          let _nw210 = new _module.C2();
          _nw210.__ctor((_this).f30, (_this).f30, _1298_v11.f29, _1298_v11.f31, _1298_v11.f35);
          _1307_v19 = _nw210;
          let _1308_v20;
          _1308_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1298_v11.f31,_1291_v4);
          let _1309_v21;
          _1309_v21 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1307_v19,new BigNumber((_1291_v4).length))).length)).plus((_1307_v19).f46),((_this.f31) ? (_1308_v20) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_1291_v4))));
          let _1310_v22;
          _1310_v22 = _module.D5.create_DC12(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-852)), ((_1311_v11) => function (_1312_i2) {
  return _1311_v11.f36;
})(_1298_v11)));
          _1309_v21 = (_1309_v21).update((_1303_v15).f44, ((!(_this.f31)) ? (_dafny.Map.Empty.slice().updateUnsafe(_1298_v11.f28,(_1310_v22).dtor_cf26)) : (_1308_v20)));
          let _rhs162 = _this.f31;
          let _rhs163 = _1286_i0;
          let _rhs164 = (new BigNumber(308)).multipliedBy((_this).f34);
          let _lhs122 = _this;
          let _lhs123 = globalState;
          let _lhs124 = globalState;
          _lhs122.f35 = _rhs162;
          _lhs123.f13 = _rhs163;
          _lhs124.f13 = _rhs164;
        } else {
          let _1313_v23;
          let _nw211 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _1313_v23 = _nw211;
          let _1314_v24;
          let _nw212 = new _module.C3();
          _nw212.__ctor(new BigNumber(-200), _1313_v23, new _dafny.CodePoint('s'.codePointAt(0)), (_this).fm28((_this).f30, true, globalState), _this.f35, (_this).f32, (_this).f33, _1286_i0, _dafny.Seq.IsPrefixOf((_this).f32, (_this).f32), _this.f29, ((_this).f34).isEqualTo(new BigNumber((_1291_v4).length)));
          _1314_v24 = _nw212;
          _1291_v4 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), function (_1315_i3) {
            return _this.f36;
          });
          let _1316_v25;
          _1316_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1314_v24,new BigNumber(((_this).f32).length));
          let _1317_v26;
          _1317_v26 = _dafny.Seq.of(_1316_v25, _1316_v25, _1316_v25, _1316_v25);
          let _1318_v27;
          _1318_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f30);
          r1 = (_dafny.ZERO).minus(((_dafny.Seq.IsPrefixOf(_1317_v26, _1317_v26)) ? (new BigNumber((_1318_v27).length)) : (_1286_i0)));
          let _1319_v28;
          _1319_v28 = _module.D5.create_DC14(_this.f35, (_this).f30);
          let _1320_v29;
          let _init26 = function (_1321_i4) {
            return _this.f36;
          };
          let _nw213 = Array((new BigNumber(19)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw213.length); _i0_26++) {
            _nw213[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1320_v29 = _nw213;
          let _1322_v30;
          let _nw214 = new _module.C0();
          _nw214.__ctor((((_1319_v28).dtor_cf30) ? (_1320_v29) : (_1320_v29)));
          _1322_v30 = _nw214;
          (globalState).f13 = (_1314_v24).f44;
        }
        let _1323_v31;
        _1323_v31 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_this.f29);
        let _1324_v32;
        _1324_v32 = _dafny.MultiSet.fromElements(new BigNumber(605), (_this).f34, (_this).f30, new BigNumber((_1323_v31).length));
        let _1325_v33;
        _1325_v33 = _dafny.Seq.of(true, (_1324_v32).IsDisjointFrom(_1324_v32));
        _1325_v33 = _dafny.Seq.Concat((_this).f32, _dafny.Seq.of(false, !(_this.f35)));
        let _1326_v34;
        _1326_v34 = _dafny.Set.fromElements((_this).f30);
        _1326_v34 = ((_1326_v34).Difference(function () {
          let _coll37 = new _dafny.Set();
          for (const _compr_37 of (_1326_v34).Elements) {
            let _1327_v35 = _compr_37;
            if ((_1326_v34).contains(_1327_v35)) {
              _coll37.add(_module.__default.safeModuloInt(_1327_v35, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(-246))).length)));
            }
          }
          return _coll37;
        }())).Difference(_1326_v34);
      }
      let _1328_v36;
      let _nw215 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
      _1328_v36 = _nw215;
      let _1329_v37;
      _1329_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1328_v36,(_this).f30);
      let _1330_v38;
      _1330_v38 = _module.D7.create_DC18((_1329_v37).update(_1328_v36, (_this).f30));
      let _pat_let_tv13 = _1329_v37;
      let _index160 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1328_v36).length));
      (_1328_v36)[_index160] = new BigNumber(((function (_pat_let19_0) {
        return function (_1331_dt__update__tmp_h0) {
          return function (_pat_let20_0) {
            return function (_1332_dt__update_hcf35_h0) {
              return _module.D7.create_DC18(_1332_dt__update_hcf35_h0);
            }(_pat_let20_0);
          }(_pat_let_tv13);
        }(_pat_let19_0);
      }(_1330_v38)).dtor_cf35).length);
      let _1333_v39;
      _1333_v39 = _dafny.MultiSet.fromElements(true);
      let _1334_v40;
      _1334_v40 = _module.D0.create_DC1(_1333_v39);
      let _1335_v41;
      _1335_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1334_v40,_this.f29);
      let _1336_v42;
      _1336_v42 = _module.D8.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_1334_v40,_this.f35));
      let _1337_v44;
      _1337_v44 = _dafny.Set.fromElements(_1334_v40);
      let _1338_v46;
      _1338_v46 = _dafny.Seq.of((_module.D8.create_DC21(_1335_v41)).dtor_cf40);
      let _pat_let_tv14 = _1333_v39;
      let _1339_v47;
      let _nw216 = Array((new BigNumber(25)).toNumber());
      _nw216[(_dafny.ZERO).toNumber()] = _1335_v41;
      _nw216[(_dafny.ONE).toNumber()] = ((_this.f28) ? (_dafny.Map.Empty.slice().updateUnsafe(_1334_v40,_this.f28)) : (_1335_v41));
      _nw216[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1334_v40,!(!(_this.f35)))).Merge(_1335_v41);
      _nw216[(new BigNumber(3)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(4)).toNumber()] = ((_1336_v42).dtor_cf40).Merge(_1335_v41);
      _nw216[(new BigNumber(5)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(6)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(7)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(8)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_1333_v39),_module.__default.fm1(_this.f31, (_this).f34, globalState));
      _nw216[(new BigNumber(10)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(11)).toNumber()] = _module.__default.fm35(_this.f35, globalState);
      _nw216[(new BigNumber(12)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(13)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm36(new BigNumber(893), (_this).f34, _this.f31, globalState),!(_this.f35))).Merge(_1335_v41);
      _nw216[(new BigNumber(14)).toNumber()] = ((_this.f28) ? (_1335_v41) : (function () {
        let _coll38 = new _dafny.Map();
        for (const _compr_38 of (_1337_v44).Elements) {
          let _1340_v43 = _compr_38;
          if ((_1337_v44).contains(_1340_v43)) {
            _coll38.push([_1340_v43,_this.f35]);
          }
        }
        return _coll38;
      }()));
      _nw216[(new BigNumber(15)).toNumber()] = (_1335_v41).Merge(function () {
        let _coll39 = new _dafny.Map();
        for (const _compr_39 of (_1335_v41).Keys.Elements) {
          let _1341_v45 = _compr_39;
          if ((_1335_v41).contains(_1341_v45)) {
            _coll39.push([_1341_v45,true]);
          }
        }
        return _coll39;
      }());
      _nw216[(new BigNumber(16)).toNumber()] = (_1338_v46)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f34), new BigNumber((_1338_v46).length))];
      _nw216[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let21_0) {
        return function (_1342_dt__update__tmp_h1) {
          return function (_pat_let22_0) {
            return function (_1343_dt__update_hcf1_h0) {
              return _module.D0.create_DC1(_1343_dt__update_hcf1_h0);
            }(_pat_let22_0);
          }(_pat_let_tv14);
        }(_pat_let21_0);
      }(_module.D0.create_DC1(_1333_v39)),_this.f28);
      _nw216[(new BigNumber(18)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(19)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(20)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(21)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(22)).toNumber()] = _module.__default.fm35(_this.f35, globalState);
      _nw216[(new BigNumber(23)).toNumber()] = _1335_v41;
      _nw216[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1334_v40,_this.f28);
      _1339_v47 = _nw216;
      let _1344_v48;
      let _nw217 = Array((new BigNumber(13)).toNumber()).fill(false);
      _1344_v48 = _nw217;
      let _index161 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_1344_v48).length));
      (_1344_v48)[_index161] = _this.f35;
      let _1345_v49;
      let _nw218 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1345_v49 = _nw218;
      let _1346_v50;
      _1346_v50 = _dafny.Seq.UnicodeFromString("ehackpxcv");
      let _1347_v51;
      _1347_v51 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(781)), function (_1348_i5) {
        return _this.f36;
      }),_this.f29);
      let _index162 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_1345_v49).length));
      (_1345_v49)[_index162] = _module.__default.fm34(new BigNumber((_1346_v50).length), (((_1333_v39).contains(_this.f35)) ? ((_1333_v39).get(_this.f35)) : (new BigNumber((_module.__default.fm37(_this.f36, _this.f31, globalState)).length))), (_this).f30, (((_1347_v51).contains(_dafny.Seq.UnicodeFromString("exfveuas"))) ? ((_1347_v51).get(_dafny.Seq.UnicodeFromString("exfveuas"))) : (true)), globalState);
      let _index163 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1328_v36).length));
      let _index164 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_1344_v48).length));
      let _index165 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_1345_v49).length));
      let _rhs165 = ((_this).fm28((_this).f34, _this.f35, globalState)).plus((_this).fm28((_this).f30, _this.f29, globalState));
      let _rhs166 = _1339_v47;
      let _rhs167 = !(true);
      let _rhs168 = (_this.f29) || (_this.f28);
      let _rhs169 = new _dafny.CodePoint('u'.codePointAt(0));
      let _lhs125 = _1328_v36;
      let _lhs126 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1328_v36).length));
      let _lhs127 = globalState;
      let _lhs128 = _1344_v48;
      let _lhs129 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_1344_v48).length));
      let _lhs130 = _1345_v49;
      let _lhs131 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_1345_v49).length));
      _lhs125[_lhs126] = _rhs165;
      _1339_v47 = _rhs166;
      _lhs127.f26 = _rhs167;
      _lhs128[_lhs129] = _rhs168;
      _lhs130[_lhs131] = _rhs169;
      let _1349_v52;
      _1349_v52 = _dafny.MultiSet.fromElements((_this).fm27((_1344_v48)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_1344_v48).length))], (_1328_v36)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((_1328_v36).length))], (_this).f30, globalState));
      let _1350_v53;
      _1350_v53 = _dafny.Seq.of((_1328_v36)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((_1328_v36).length))], (_1328_v36)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((_1328_v36).length))], (_this).f30, new BigNumber((_dafny.Seq.of(new BigNumber(-627), (_this).f34)).length), (_dafny.ZERO).minus((_this).f30));
      let _1351_v54;
      let _nw219 = new _module.C2();
      _nw219.__ctor((_1328_v36)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((_1328_v36).length))], (((_1349_v52).contains((_1350_v53)[_module.__default.safeIndex((_1328_v36)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((_1328_v36).length))], new BigNumber((_1350_v53).length))])) ? ((_1349_v52).get((_1350_v53)[_module.__default.safeIndex((_1328_v36)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((_1328_v36).length))], new BigNumber((_1350_v53).length))])) : ((_this).f34)), _this.f28, _this.f35, !(_this.f29));
      _1351_v54 = _nw219;
      let _1352_v55;
      _1352_v55 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(799)), function (_1353_i6) {
        return (_this).f30;
      }),new BigNumber((_1350_v53).length));
      let _rhs170 = new BigNumber((_dafny.Seq.update(((((_this).f30).isLessThanOrEqualTo(new BigNumber(((_this).f32).length))) ? (_1346_v50) : (_1346_v50)), _module.__default.safeIndex(((_1351_v54).f46).minus((_dafny.ZERO).minus(new BigNumber((_1352_v55).length))), new BigNumber((((((_this).f30).isLessThanOrEqualTo(new BigNumber(((_this).f32).length))) ? (_1346_v50) : (_1346_v50))).length)), (_1345_v49)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_1345_v49).length))])).length);
      let _rhs171 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1346_v50, _1346_v50), _dafny.Seq.of(_this.f36));
      let _rhs172 = ((_this.f35) ? (new BigNumber((_dafny.Seq.Concat(_1346_v50, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(178)), ((_1354_v49) => function (_1355_i7) {
        return (_1354_v49)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_1354_v49).length))];
      })(_1345_v49)), _module.__default.safeIndex((_this).f34, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(178)), ((_1356_v49) => function (_1357_i7) {
        return (_1356_v49)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_1356_v49).length))];
      })(_1345_v49))).length)), (_1345_v49)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_1345_v49).length))]))).length)) : ((_this).f34));
      let _rhs173 = !(_this.f35) || (_this.f28);
      let _lhs132 = globalState;
      let _lhs133 = globalState;
      r1 = _rhs170;
      _1346_v50 = _rhs171;
      _lhs132.f13 = _rhs172;
      _lhs133.f12 = _rhs173;
      let _index166 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_1344_v48).length));
      (_1344_v48)[_index166] = _this.f29;
      let _index167 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_1344_v48).length));
      (_1344_v48)[_index167] = (_1351_v54).fm30(false, globalState);
      r0 = _this.f31;
      r1 = (_this).f34;
      r2 = _1346_v50;
      return [r0, r1, r2];
    }
    m11(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      let _1358_v0;
      _1358_v0 = _module.D8.create_DC22((_this).f34, (_this).f34);
      let _1359_v1;
      _1359_v1 = _module.D8.create_DC25(_module.D8.create_DC25(_1358_v0));
      let _source22 = _1359_v1;
      if (_source22.is_DC22) {
        let _1360___mcc_h0 = (_source22).cf41;
        let _1361___mcc_h1 = (_source22).cf42;
        let _1362_cf42 = _1361___mcc_h1;
        let _1363_cf41 = _1360___mcc_h0;
        let _1364_v2;
        let _nw220 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _1364_v2 = _nw220;
        let _1365_v3;
        let _nw221 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1365_v3 = _nw221;
        let _index168 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1364_v2).length));
        (_1364_v2)[_index168] = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_1365_v3);
        let _1366_v4;
        _1366_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_1365_v3);
        let _index169 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1364_v2).length));
        (_1364_v2)[_index169] = (_1366_v4).Merge(_1366_v4);
        let _1367_v5;
        let _init27 = function (_1368_i0) {
          return false;
        };
        let _nw222 = Array((new BigNumber(3)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw222.length); _i0_27++) {
          _nw222[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1367_v5 = _nw222;
        let _1369_v6;
        _1369_v6 = _module.D1.create_DC2(_1367_v5);
        let _1370_v7;
        _1370_v7 = _module.D1.create_DC4(_1369_v6);
        let _source23 = _1370_v7;
        if (_source23.is_DC3) {
          let _1371___mcc_h12 = (_source23).cf3;
          let _1372___mcc_h13 = (_source23).cf4;
          let _1373___mcc_h14 = (_source23).cf5;
          let _1374___mcc_h15 = (_source23).cf6;
          let _1375_cf6 = _1374___mcc_h15;
          let _1376_cf5 = _1373___mcc_h14;
          let _1377_cf4 = _1372___mcc_h13;
          let _1378_cf3 = _1371___mcc_h12;
          let _1379_v8;
          _1379_v8 = _dafny.Seq.UnicodeFromString("tyollv");
          let _1380_v9;
          _1380_v9 = _dafny.Seq.of(new BigNumber((_1379_v8).length));
          let _1381_v10;
          let _out56;
          _out56 = _module.__default.m0(_1380_v9, !(_this.f29), _1363_cf41, p0, globalState);
          _1381_v10 = _out56;
          let _1382_v11;
          let _nw223 = Array((new BigNumber(27)).toNumber()).fill([]);
          _1382_v11 = _nw223;
          let _index170 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1382_v11).length));
          (_1382_v11)[_index170] = ((_this.f31) ? (_1367_v5) : (_1367_v5));
          let _index171 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_1365_v3).length));
          (_1365_v3)[_index171] = _1377_cf4;
          let _1383_v12;
          _1383_v12 = _dafny.MultiSet.fromElements(!(_this.f35));
          let _index172 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1382_v11).length));
          let _index173 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_1365_v3).length));
          let _rhs174 = _1367_v5;
          let _rhs175 = (((_1383_v12).contains(_this.f31)) ? ((_1383_v12).get(_this.f31)) : ((_this).f30));
          let _rhs176 = _1375_cf6;
          let _rhs177 = _1359_v1;
          let _lhs134 = _1382_v11;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1382_v11).length));
          let _lhs136 = _1365_v3;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_1365_v3).length));
          _lhs134[_lhs135] = _rhs174;
          r2 = _rhs175;
          _lhs136[_lhs137] = _rhs176;
          _1359_v1 = _rhs177;
          _1379_v8 = _1379_v8;
          r2 = _1363_cf41;
        } else if (_source23.is_DC2) {
          let _1384___mcc_h16 = (_source23).cf2;
          let _1385_cf2 = _1384___mcc_h16;
          let _1386_v13;
          let _nw224 = Array((new BigNumber(6)).toNumber());
          _1386_v13 = _nw224;
          let _1387_v14;
          _1387_v14 = _dafny.MultiSet.fromElements(_1362_cf42);
          let _1388_v15;
          _1388_v15 = _dafny.Set.fromElements(_1387_v14);
          let _1389_v16;
          let _nw225 = new _module.C2();
          _nw225.__ctor(new BigNumber(643), new BigNumber((_1388_v15).length), _this.f28, p0, !(_this.f28));
          _1389_v16 = _nw225;
          let _index174 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1386_v13).length));
          (_1386_v13)[_index174] = _1389_v16;
          let _1390_v17;
          _1390_v17 = _module.D9.create_DC26(_1389_v16);
          let _index175 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1386_v13).length));
          (_1386_v13)[_index175] = (_1390_v17).dtor_cf52;
          let _1391_v18;
          _1391_v18 = _dafny.MultiSet.fromElements((_this).f32, (_this).f32);
          (globalState).f13 = ((_this).f30).minus(new BigNumber((_1391_v18).cardinality()));
          let _1392_v19;
          _1392_v19 = _dafny.Set.fromElements((_this).f30);
          let _index176 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1365_v3).length));
          (_1365_v3)[_index176] = new BigNumber((_1392_v19).length);
          let _1393_v20;
          _1393_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_1389_v16).f46);
          let _1394_v21;
          _1394_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1365_v3,new BigNumber((_1393_v20).length));
          let _index177 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1365_v3).length));
          (_1365_v3)[_index177] = (((_1394_v21).contains(_1365_v3)) ? ((_1394_v21).get(_1365_v3)) : ((_this).fm27(_this.f29, _1363_cf41, new BigNumber(736), globalState)));
          let _1395_v22;
          _1395_v22 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_this.f31);
          let _1396_v23;
          _1396_v23 = _dafny.Seq.of(_1362_cf42);
          _1395_v22 = (_1395_v22).update(!(((false) ? (_this.f31) : (_this.f31))), _dafny.areEqual(_1396_v23, _1396_v23));
        } else {
          let _1397___mcc_h17 = (_source23).cf7;
          let _1398_cf7 = _1397___mcc_h17;
          let _1399_v24;
          _1399_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1367_v5,_this.f36);
          let _1400_v25;
          _1400_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1399_v24);
          let _1401_v26;
          let _nw226 = Array((new BigNumber(15)).toNumber());
          _nw226[(_dafny.ZERO).toNumber()] = _1399_v24;
          _nw226[(_dafny.ONE).toNumber()] = _1399_v24;
          _nw226[(new BigNumber(2)).toNumber()] = _1399_v24;
          _nw226[(new BigNumber(3)).toNumber()] = _1399_v24;
          _nw226[(new BigNumber(4)).toNumber()] = (_1399_v24).Merge(_1399_v24);
          _nw226[(new BigNumber(5)).toNumber()] = (_1399_v24).Merge(_1399_v24);
          _nw226[(new BigNumber(6)).toNumber()] = _1399_v24;
          _nw226[(new BigNumber(7)).toNumber()] = _1399_v24;
          _nw226[(new BigNumber(8)).toNumber()] = ((_1399_v24).update(_1367_v5, _this.f36)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1367_v5,_this.f36));
          _nw226[(new BigNumber(9)).toNumber()] = _1399_v24;
          _nw226[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1367_v5,_this.f36);
          _nw226[(new BigNumber(11)).toNumber()] = (((_1400_v25).contains(_this.f31)) ? ((_1400_v25).get(_this.f31)) : (_1399_v24));
          _nw226[(new BigNumber(12)).toNumber()] = (_1399_v24).Merge(_1399_v24);
          _nw226[(new BigNumber(13)).toNumber()] = _1399_v24;
          _nw226[(new BigNumber(14)).toNumber()] = _1399_v24;
          _1401_v26 = _nw226;
          let _index178 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1401_v26).length));
          (_1401_v26)[_index178] = _1399_v24;
          let _1402_v27;
          _1402_v27 = _module.D1.create_DC2(_1367_v5);
          let _1403_v28;
          let _nw227 = new _module.C1();
          _nw227.__ctor(_this.f35, _this.f29, false);
          _1403_v28 = _nw227;
          let _1404_v29;
          _1404_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1403_v28,new BigNumber((_dafny.Seq.UnicodeFromString("auuk")).length));
          let _index179 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1401_v26).length));
          let _rhs178 = _1399_v24;
          let _rhs179 = false;
          let _rhs180 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(781)), ((_1405_cf41) => function (_1406_i1) {
            return (_module.D5.create_DC13(_this.f28, _dafny.Map.Empty.slice().updateUnsafe(true,_1405_cf41), _this.f36)).dtor_cf29;
          })(_1363_cf41))).length), ((_this).f34).plus((((_1404_v29).contains(_1403_v28)) ? ((_1404_v29).get(_1403_v28)) : ((_this).fm28(new BigNumber((_dafny.MultiSet.fromElements(_this.f35)).cardinality()), _this.f31, globalState)))))));
          let _rhs181 = _1362_cf42;
          let _rhs182 = _1402_v27;
          let _lhs138 = _1401_v26;
          let _lhs139 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1401_v26).length));
          let _lhs140 = _this;
          let _lhs141 = globalState;
          let _lhs142 = globalState;
          _lhs138[_lhs139] = _rhs178;
          _lhs140.f29 = _rhs179;
          _lhs141.f13 = _rhs180;
          _lhs142.f13 = _rhs181;
          _1402_v27 = _rhs182;
          let _1407_v30;
          let _init28 = function (_1408_i2) {
            return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jhgvtevm"), _dafny.Seq.UnicodeFromString("chiwr"));
          };
          let _nw228 = Array((new BigNumber(15)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw228.length); _i0_28++) {
            _nw228[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1407_v30 = _nw228;
          let _1409_v31;
          _1409_v31 = _dafny.Seq.UnicodeFromString("yguw");
          let _index180 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_1407_v30).length));
          (_1407_v30)[_index180] = ((p0) ? (_1409_v31) : (_1409_v31));
          let _index181 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_1407_v30).length));
          (_1407_v30)[_index181] = _dafny.Seq.Concat(_1409_v31, _dafny.Seq.UnicodeFromString("xhyw"));
          _1363_cf41 = _1363_cf41;
          _1363_cf41 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(554)), function (_1410_i3) {
            return _this.f36;
          }), _1409_v31)).length);
        }
        let _1411_v33;
        _1411_v33 = _dafny.MultiSet.fromElements(_1363_cf41);
        let _1412_v34;
        _1412_v34 = _module.D4.create_DC10(_1411_v33);
        let _1413_v35;
        _1413_v35 = _dafny.Map.Empty.slice().updateUnsafe((_1412_v34).dtor_cf22,true);
        let _1414_v36;
        _1414_v36 = _dafny.Seq.of(new BigNumber((function () {
          let _coll40 = new _dafny.Map();
          for (const _compr_40 of (_1413_v35).Keys.Elements) {
            let _1415_v32 = _compr_40;
            if ((_1413_v35).contains(_1415_v32)) {
              _coll40.push([_1415_v32,_this.f29]);
            }
          }
          return _coll40;
        }()).length), _1363_cf41);
        let _1416_v37;
        _1416_v37 = _dafny.Set.fromElements((_this).f34);
        (globalState).f13 = new BigNumber((_dafny.Seq.update(_1414_v36, _module.__default.safeIndex(new BigNumber((_1416_v37).length), new BigNumber((_1414_v36).length)), ((_this).f30).plus(_1363_cf41))).length);
        (globalState).f13 = (_dafny.ZERO).minus((_this).f30);
      } else if (_source22.is_DC23) {
        let _1417___mcc_h2 = (_source22).cf43;
        let _1418___mcc_h3 = (_source22).cf44;
        let _1419___mcc_h4 = (_source22).cf45;
        let _1420_cf45 = _1419___mcc_h4;
        let _1421_cf44 = _1418___mcc_h3;
        let _1422_cf43 = _1417___mcc_h2;
        let _1423_v38;
        let _nw229 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1423_v38 = _nw229;
        let _1424_v39;
        let _nw230 = new _module.C0();
        _nw230.__ctor(_1423_v38);
        _1424_v39 = _nw230;
        let _1425_v40;
        _1425_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_1424_v39);
        let _1426_v41;
        _1426_v41 = _dafny.Seq.of(new BigNumber(835));
        let _1427_v42;
        _1427_v42 = _dafny.Seq.of(new BigNumber(161), new BigNumber((_1426_v41).length));
        let _1428_v43;
        _1428_v43 = _dafny.Set.fromElements(_1427_v42);
        let _1429_v44;
        _1429_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1428_v43).length),_1421_cf44);
        let _1430_v45;
        _1430_v45 = _dafny.Set.fromElements(true);
        let _1431_v46;
        _1431_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1429_v44,_1430_v45);
        let _1432_v47;
        _1432_v47 = _dafny.Seq.of((((_1431_v46).contains(_1429_v44)) ? ((_1431_v46).get(_1429_v44)) : (_1430_v45)), _dafny.Set.fromElements(_1421_cf44, _1420_cf45), _1430_v45);
        let _1433_v48;
        let _nw231 = Array((new BigNumber(28)).toNumber());
        _nw231[(_dafny.ZERO).toNumber()] = _1424_v39;
        _nw231[(_dafny.ONE).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(2)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(3)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(4)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(5)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(6)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(7)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(8)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(9)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(10)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(11)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(12)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(13)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(14)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(15)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(16)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(17)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(18)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(19)).toNumber()] = (((_1425_v40).contains(new BigNumber((_1432_v47).length))) ? ((_1425_v40).get(new BigNumber((_1432_v47).length))) : (_1424_v39));
        _nw231[(new BigNumber(20)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(21)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(22)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(23)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(24)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(25)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(26)).toNumber()] = _1424_v39;
        _nw231[(new BigNumber(27)).toNumber()] = _1424_v39;
        _1433_v48 = _nw231;
        _1433_v48 = _1433_v48;
        r2 = (new BigNumber(((_this).f32).length)).minus((_this).f30);
        (globalState).f13 = new BigNumber(262);
        (globalState).f12 = _1422_cf43;
      } else if (_source22.is_DC24) {
        let _1434___mcc_h5 = (_source22).cf46;
        let _1435___mcc_h6 = (_source22).cf47;
        let _1436___mcc_h7 = (_source22).cf48;
        let _1437___mcc_h8 = (_source22).cf49;
        let _1438___mcc_h9 = (_source22).cf50;
        let _1439_cf50 = _1438___mcc_h9;
        let _1440_cf49 = _1437___mcc_h8;
        let _1441_cf48 = _1436___mcc_h7;
        let _1442_cf47 = _1435___mcc_h6;
        let _1443_cf46 = _1434___mcc_h5;
        let _1444_v49;
        _1444_v49 = _dafny.Seq.UnicodeFromString("c");
        let _1445_v50;
        _1445_v50 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f34),_this.f36);
        let _1446_v51;
        let _nw232 = Array((new BigNumber(16)).toNumber());
        _nw232[(_dafny.ZERO).toNumber()] = _this.f36;
        _nw232[(_dafny.ONE).toNumber()] = _this.f36;
        _nw232[(new BigNumber(2)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(3)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
        _nw232[(new BigNumber(5)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(6)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(7)).toNumber()] = (((_1445_v50).contains(new BigNumber((_1444_v49).length))) ? ((_1445_v50).get(new BigNumber((_1444_v49).length))) : (_this.f36));
        _nw232[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
        _nw232[(new BigNumber(9)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(10)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(11)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(12)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(13)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(14)).toNumber()] = _this.f36;
        _nw232[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
        _1446_v51 = _nw232;
        let _1447_v52;
        let _nw233 = new _module.C0();
        _nw233.__ctor(_1446_v51);
        _1447_v52 = _nw233;
        let _source24 = _module.D1.create_DC3(_this.f35, (_1439_cf50)[_module.__default.safeIndex(new BigNumber((_1444_v49).length), new BigNumber((_1439_cf50).length))], _1447_v52, _module.__default.safeModuloInt(_1442_cf47, _1442_cf47));
        if (_source24.is_DC3) {
          let _1448___mcc_h18 = (_source24).cf3;
          let _1449___mcc_h19 = (_source24).cf4;
          let _1450___mcc_h20 = (_source24).cf5;
          let _1451___mcc_h21 = (_source24).cf6;
          let _1452_cf6 = _1451___mcc_h21;
          let _1453_cf5 = _1450___mcc_h20;
          let _1454_cf4 = _1449___mcc_h19;
          let _1455_cf3 = _1448___mcc_h18;
          let _1456_v53;
          let _nw234 = new _module.C0();
          _nw234.__ctor(_1447_v52.f37);
          _1456_v53 = _nw234;
          let _1457_v54;
          let _nw235 = new _module.C1();
          _nw235.__ctor(_1440_cf49, false, true);
          _1457_v54 = _nw235;
          let _1458_v55;
          _1458_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1457_v54,_this.f36);
          _1454_cf4 = (_module.__default.safeModuloInt(new BigNumber((_1458_v55).length), _1452_cf6)).minus(_1454_cf4);
          let _1459_v56;
          let _init29 = function (_1460_i4) {
            return (_1460_i4).minus((_this).f34);
          };
          let _nw236 = Array((new BigNumber(29)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw236.length); _i0_29++) {
            _nw236[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1459_v56 = _nw236;
          let _1461_v57;
          _1461_v57 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f36);
          let _1462_v58;
          _1462_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1459_v56,new BigNumber((_1461_v57).length));
          let _1463_v59;
          _1463_v59 = _dafny.MultiSet.fromElements(_1457_v54.f28);
          let _1464_v60;
          let _nw237 = new _module.C2();
          _nw237.__ctor((_this).f30, new BigNumber((_1462_v58).length), (_1463_v59).IsDisjointFrom(_dafny.MultiSet.fromElements(_1457_v54.f28, _this.f28)), !(!(_1457_v54.f28)), !(_1457_v54.f28));
          _1464_v60 = _nw237;
          r3 = (_this).fm27(((_this).f32)[_module.__default.safeIndex(_1452_cf6, new BigNumber(((_this).f32).length))], (new BigNumber((_1439_cf50).length)).multipliedBy(_1454_cf4), _1454_cf4, globalState);
        } else if (_source24.is_DC2) {
          let _1465___mcc_h22 = (_source24).cf2;
          let _1466_cf2 = _1465___mcc_h22;
          (_this).f36 = _this.f36;
          let _1467_v61;
          let _nw238 = new _module.C0();
          _nw238.__ctor(_1447_v52.f37);
          _1467_v61 = _nw238;
          r1 = !((_1441_cf48)[_module.__default.safeIndex((_this).f30, new BigNumber((_1441_cf48).length))]);
          (globalState).f13 = (_this).f30;
        } else {
          let _1468___mcc_h23 = (_source24).cf7;
          let _1469_cf7 = _1468___mcc_h23;
          let _1470_v62;
          let _nw239 = new _module.C0();
          _nw239.__ctor(_1447_v52.f37);
          _1470_v62 = _nw239;
          r2 = _module.__default.safeModuloInt((_this).f34, _1442_cf47);
          let _1471_v63;
          _1471_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,!(_1443_cf46));
          let _1472_v64;
          _1472_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,new BigNumber((_1471_v63).length));
          let _1473_v65;
          let _nw240 = Array((new BigNumber(23)).toNumber());
          _nw240[(_dafny.ZERO).toNumber()] = _1441_cf48;
          _nw240[(_dafny.ONE).toNumber()] = (_this).f32;
          _nw240[(new BigNumber(2)).toNumber()] = _1441_cf48;
          _nw240[(new BigNumber(3)).toNumber()] = (_this).f32;
          _nw240[(new BigNumber(4)).toNumber()] = _1441_cf48;
          _nw240[(new BigNumber(5)).toNumber()] = (_this).f32;
          _nw240[(new BigNumber(6)).toNumber()] = _1441_cf48;
          _nw240[(new BigNumber(7)).toNumber()] = _module.__default.fm0(_1442_cf47, _this.f35, p0, new BigNumber((_1472_v64).length), globalState);
          _nw240[(new BigNumber(8)).toNumber()] = (_this).f32;
          _nw240[(new BigNumber(9)).toNumber()] = (_this).f32;
          _nw240[(new BigNumber(10)).toNumber()] = _1441_cf48;
          _nw240[(new BigNumber(11)).toNumber()] = (_this).f32;
          _nw240[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(_this.f35, _1440_cf49, _1443_cf46, _this.f31, _this.f29);
          _nw240[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update((_this).f32, _module.__default.safeIndex(_1442_cf47, new BigNumber(((_this).f32).length)), p0), _module.__default.safeIndex((_this).f34, new BigNumber((_dafny.Seq.update((_this).f32, _module.__default.safeIndex(_1442_cf47, new BigNumber(((_this).f32).length)), p0)).length)), _this.f35);
          _nw240[(new BigNumber(14)).toNumber()] = _1441_cf48;
          _nw240[(new BigNumber(15)).toNumber()] = _dafny.Seq.update((_this).f32, _module.__default.safeIndex(_1442_cf47, new BigNumber(((_this).f32).length)), !(_1443_cf46));
          _nw240[(new BigNumber(16)).toNumber()] = (_this).f32;
          _nw240[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(p0);
          _nw240[(new BigNumber(18)).toNumber()] = (_this).f32;
          _nw240[(new BigNumber(19)).toNumber()] = (_this).f32;
          _nw240[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(_this.f35);
          _nw240[(new BigNumber(21)).toNumber()] = _1441_cf48;
          _nw240[(new BigNumber(22)).toNumber()] = _dafny.Seq.of(p0, _1443_cf46, true, _this.f28, false);
          _1473_v65 = _nw240;
          let _1474_v66;
          _1474_v66 = _dafny.Set.fromElements(_this.f31, _1440_cf49);
          let _1475_v67;
          _1475_v67 = _module.D2.create_DC6(_1442_cf47, false, (_this).f30, _1443_cf46, _1442_cf47);
          let _1476_v68;
          let _nw241 = new _module.C3();
          _nw241.__ctor((_this).f34, _1473_v65, _this.f36, (_this).f34, (_module.__default.fm1(_this.f35, new BigNumber((_module.__default.fm38(false, new BigNumber((_1474_v66).length), globalState)).length), globalState)) && (_1440_cf49), _dafny.Seq.update(_dafny.Seq.update(((_this).f33)[_module.__default.safeIndex(_1442_cf47, new BigNumber(((_this).f33).length))], _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_this.f36)).length), new BigNumber((((_this).f33)[_module.__default.safeIndex(_1442_cf47, new BigNumber(((_this).f33).length))]).length)), p0), _module.__default.safeIndex(new BigNumber(540), new BigNumber((_dafny.Seq.update(((_this).f33)[_module.__default.safeIndex(_1442_cf47, new BigNumber(((_this).f33).length))], _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_this.f36)).length), new BigNumber((((_this).f33)[_module.__default.safeIndex(_1442_cf47, new BigNumber(((_this).f33).length))]).length)), p0)).length)), _this.f35), (_this).f33, (_this).f34, _this.f35, (_1475_v67).dtor_cf10, p0);
          _1476_v68 = _nw241;
          r3 = (_1442_cf47).multipliedBy((_1476_v68).f44);
        }
        if (_1443_cf46) {
          let _1477_v69;
          _1477_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,!(_1440_cf49));
          let _1478_v70;
          _1478_v70 = _dafny.Set.fromElements(_1477_v69);
          let _1479_v71;
          _1479_v71 = _module.D7.create_DC19(_this.f29, (_dafny.ZERO).minus(_1442_cf47), (_this).f30);
          _1478_v70 = (((_1479_v71).dtor_cf36) ? ((_1478_v70).Difference(_1478_v70)) : (_1478_v70));
          let _1480_v72;
          _1480_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1442_cf47,_dafny.Set.fromElements((_this).f34));
          let _1481_v73;
          _1481_v73 = _dafny.Set.fromElements(_this.f35, _this.f31);
          let _1482_v74;
          let _nw242 = Array((new BigNumber(20)).toNumber());
          _nw242[(_dafny.ZERO).toNumber()] = (_this).f34;
          _nw242[(_dafny.ONE).toNumber()] = (_this).f34;
          _nw242[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_1442_cf47);
          _nw242[(new BigNumber(3)).toNumber()] = (_this).f30;
          _nw242[(new BigNumber(4)).toNumber()] = new BigNumber((_1480_v72).length);
          _nw242[(new BigNumber(5)).toNumber()] = (_this).f30;
          _nw242[(new BigNumber(6)).toNumber()] = (_this).f34;
          _nw242[(new BigNumber(7)).toNumber()] = _1442_cf47;
          _nw242[(new BigNumber(8)).toNumber()] = new BigNumber((_1481_v73).length);
          _nw242[(new BigNumber(9)).toNumber()] = _1442_cf47;
          _nw242[(new BigNumber(10)).toNumber()] = (_this).f34;
          _nw242[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_this).f30);
          _nw242[(new BigNumber(12)).toNumber()] = (_this).f30;
          _nw242[(new BigNumber(13)).toNumber()] = (_this).f30;
          _nw242[(new BigNumber(14)).toNumber()] = new BigNumber(300);
          _nw242[(new BigNumber(15)).toNumber()] = _1442_cf47;
          _nw242[(new BigNumber(16)).toNumber()] = (_this).f30;
          _nw242[(new BigNumber(17)).toNumber()] = (_this).f34;
          _nw242[(new BigNumber(18)).toNumber()] = (_this).f34;
          _nw242[(new BigNumber(19)).toNumber()] = new BigNumber(162);
          _1482_v74 = _nw242;
          let _1483_v75;
          _1483_v75 = _dafny.Seq.of(_1482_v74);
          let _1484_v76;
          _1484_v76 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f34, (_this).f30),(_this).fm27(_this.f31, new BigNumber((_1483_v75).length), (_this).fm28((_this).f34, (_this).fm7(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f34)), _1442_cf47, globalState), globalState), globalState));
          _1484_v76 = (_1484_v76).update((_dafny.ZERO).minus(_1442_cf47), _module.__default.safeModuloInt(_1442_cf47, _1442_cf47));
          r3 = ((_this.f35) ? (_module.__default.safeModuloInt(new BigNumber((_1444_v49).length), (_this).f34)) : ((new BigNumber(636)).minus((_this).f30)));
          let _1485_v77;
          let _nw243 = Array((new BigNumber(11)).toNumber());
          _1485_v77 = _nw243;
          let _1486_v78;
          _1486_v78 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1442_cf47));
          let _1487_v79;
          let _nw244 = new _module.C1();
          _nw244.__ctor(_this.f31, _this.f35, (_this).fm7(_1486_v78, (_this).f34, globalState));
          _1487_v79 = _nw244;
          let _index182 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_1485_v77).length));
          (_1485_v77)[_index182] = _1487_v79;
          let _1488_v80;
          let _nw245 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _1488_v80 = _nw245;
          let _1489_v81;
          let _init30 = function (_1490_i5) {
            return _this.f29;
          };
          let _nw246 = Array((new BigNumber(26)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw246.length); _i0_30++) {
            _nw246[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1489_v81 = _nw246;
          let _1491_v82;
          _1491_v82 = _dafny.Map.Empty.slice().updateUnsafe((((_1484_v76).contains((_this).f30)) ? ((_1484_v76).get((_this).f30)) : ((_this).f30)),_1489_v81);
          let _index183 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_1488_v80).length));
          (_1488_v80)[_index183] = _1491_v82;
          let _1492_v83;
          _1492_v83 = _dafny.Map.Empty.slice().updateUnsafe(false,_1487_v79);
          let _1493_v84;
          _1493_v84 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_1491_v82);
          let _index184 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_1485_v77).length));
          let _index185 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_1488_v80).length));
          let _rhs183 = _module.__default.safeDivisionInt((_this).f34, _1442_cf47);
          let _rhs184 = (((_1492_v83).contains(_this.f29)) ? ((_1492_v83).get(_this.f29)) : ((_module.D10.create_DC28(_1487_v79)).dtor_cf55));
          let _rhs185 = (((((_1493_v84).contains((_1487_v79).f43)) ? ((_1493_v84).get((_1487_v79).f43)) : (_1491_v82))).Merge(_1491_v82)).Merge(_1491_v82);
          let _lhs143 = globalState;
          let _lhs144 = _1485_v77;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_1485_v77).length));
          let _lhs146 = _1488_v80;
          let _lhs147 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_1488_v80).length));
          _lhs143.f13 = _rhs183;
          _lhs144[_lhs145] = _rhs184;
          _lhs146[_lhs147] = _rhs185;
          let _1494_v85;
          let _nw247 = Array((new BigNumber(11)).toNumber());
          _nw247[(_dafny.ZERO).toNumber()] = _1489_v81;
          _nw247[(_dafny.ONE).toNumber()] = _1489_v81;
          _nw247[(new BigNumber(2)).toNumber()] = _1489_v81;
          _nw247[(new BigNumber(3)).toNumber()] = _1489_v81;
          _nw247[(new BigNumber(4)).toNumber()] = _1489_v81;
          _nw247[(new BigNumber(5)).toNumber()] = _1489_v81;
          _nw247[(new BigNumber(6)).toNumber()] = _1489_v81;
          _nw247[(new BigNumber(7)).toNumber()] = _1489_v81;
          _nw247[(new BigNumber(8)).toNumber()] = _1489_v81;
          _nw247[(new BigNumber(9)).toNumber()] = ((_this.f31) ? (_1489_v81) : (_1489_v81));
          _nw247[(new BigNumber(10)).toNumber()] = _1489_v81;
          _1494_v85 = _nw247;
          let _1495_v86;
          _1495_v86 = _dafny.Seq.of(_1489_v81, _1489_v81);
          let _index186 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1494_v85).length));
          (_1494_v85)[_index186] = (_1495_v86)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1444_v49).length)), new BigNumber((_1495_v86).length))];
          let _index187 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1494_v85).length));
          (_1494_v85)[_index187] = _1489_v81;
        } else {
          (globalState).f1 = _this.f28;
          (globalState).f1 = _module.__default.fm1(false, (_this).fm28((_this).f30, _1440_cf49, globalState), globalState);
          let _1496_v87;
          _1496_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1443_cf46,!(_this.f31));
          _1496_v87 = (_1496_v87).update(_this.f29, ((_this).f34).isLessThanOrEqualTo((_this).f34));
          let _1497_v88;
          let _nw248 = new _module.C0();
          _nw248.__ctor(_1447_v52.f37);
          _1497_v88 = _nw248;
          (globalState).f26 = (_this.f31) && (_this.f31);
        }
        let _1498_v89;
        _1498_v89 = _module.D10.create_DC29(_1442_cf47, _module.__default.fm1(_this.f29, _1442_cf47, globalState));
        (_this).f28 = (_1498_v89).dtor_cf57;
        (globalState).f1 = !(_this.f35) || (_1440_cf49);
      } else if (_source22.is_DC21) {
        let _1499___mcc_h10 = (_source22).cf40;
        let _1500_cf40 = _1499___mcc_h10;
        let _1501_v90;
        _1501_v90 = _dafny.Seq.UnicodeFromString("ggytjvvm");
        let _1502_v91;
        let _nw249 = Array((new BigNumber(29)).toNumber());
        _nw249[(_dafny.ZERO).toNumber()] = _1501_v90;
        _nw249[(_dafny.ONE).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(2)).toNumber()] = ((false) ? (_1501_v90) : (_1501_v90));
        _nw249[(new BigNumber(3)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(4)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(5)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(6)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1501_v90, _1501_v90);
        _nw249[(new BigNumber(8)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(9)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_1501_v90, _module.__default.safeIndex((_this).f30, new BigNumber((_1501_v90).length)), _this.f36);
        _nw249[(new BigNumber(11)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("nqoj");
        _nw249[(new BigNumber(13)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(14)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(15)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(16)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(17)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(18)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(19)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_1501_v90, _1501_v90);
        _nw249[(new BigNumber(21)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-167)), function (_1503_i6) {
          return _this.f36;
        });
        _nw249[(new BigNumber(22)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(23)).toNumber()] = (_this).fm3(_this.f29, p0, globalState);
        _nw249[(new BigNumber(24)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(25)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(26)).toNumber()] = _dafny.Seq.UnicodeFromString("lskefdq");
        _nw249[(new BigNumber(27)).toNumber()] = _1501_v90;
        _nw249[(new BigNumber(28)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1501_v90, _dafny.Seq.UnicodeFromString("k")), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f34), new BigNumber((_dafny.Seq.Concat(_1501_v90, _dafny.Seq.UnicodeFromString("k"))).length)), _this.f36);
        _1502_v91 = _nw249;
        let _index188 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length));
        (_1502_v91)[_index188] = _dafny.Seq.Concat(_1501_v90, _1501_v90);
        let _index189 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length));
        let _rhs186 = _1501_v90;
        let _rhs187 = new BigNumber(852);
        let _lhs148 = _1502_v91;
        let _lhs149 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length));
        let _lhs150 = globalState;
        _lhs148[_lhs149] = _rhs186;
        _lhs150.f13 = _rhs187;
        if (_this.f35) {
          (_this).f36 = _this.f36;
          let _1504_v92;
          let _nw250 = Array((new BigNumber(19)).toNumber()).fill(false);
          _1504_v92 = _nw250;
          let _1505_v93;
          _1505_v93 = _dafny.Seq.of((_this).f30, (_dafny.ZERO).minus((_this).f34));
          let _1506_v94;
          _1506_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1505_v93,true);
          let _1507_v95;
          let _init31 = function (_1508_i7) {
            return _this.f36;
          };
          let _nw251 = Array((new BigNumber(23)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw251.length); _i0_31++) {
            _nw251[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1507_v95 = _nw251;
          let _1509_v96;
          let _nw252 = new _module.C0();
          _nw252.__ctor(_1507_v95);
          _1509_v96 = _nw252;
          let _1510_v97;
          _1510_v97 = _dafny.Seq.of(_1509_v96);
          let _1511_v98;
          _1511_v98 = _module.D9.create_DC27((_1510_v97)[_module.__default.safeIndex((_this).f34, new BigNumber((_1510_v97).length))], _1501_v90);
          let _1512_v99;
          _1512_v99 = _dafny.MultiSet.fromElements(_1511_v98);
          let _index190 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1504_v92).length));
          (_1504_v92)[_index190] = (new BigNumber((_1512_v99).cardinality())).isLessThan(new BigNumber((_1506_v94).length));
          let _index191 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1504_v92).length));
          (_1504_v92)[_index191] = _dafny.Seq.IsPrefixOf(_1501_v90, (_1502_v91)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length))]);
          let _1513_v100;
          let _init32 = function (_1514_i8) {
            return _module.D8.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.FromArray((_this).f32)),_this.f35));
          };
          let _nw253 = Array((new BigNumber(27)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw253.length); _i0_32++) {
            _nw253[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1513_v100 = _nw253;
          let _1515_v101;
          _1515_v101 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_this.f31);
          let _1516_v102;
          _1516_v102 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1504_v92)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_1504_v92).length))]);
          let _1517_v103;
          _1517_v103 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1515_v101).length),_1516_v102);
          let _1518_v104;
          _1518_v104 = _dafny.Map.Empty.slice().updateUnsafe(_1513_v100,new BigNumber(((_1517_v103).Merge(_1517_v103)).length));
          _1518_v104 = _1518_v104;
          let _1519_v105;
          _1519_v105 = _dafny.MultiSet.fromElements((_this).f30);
          let _1520_v106;
          let _nw254 = new _module.C5();
          _nw254.__ctor(_this.f29, (_this).f34, (_this).f34, _this.f35, _module.__default.fm0(new BigNumber((_1519_v105).cardinality()), p0, _this.f29, (_this).f34, globalState), (_module.D14.create_DC36(_module.__default.fm51(p0, _this.f36, globalState))).dtor_cf67, new BigNumber((_1515_v101).length), _this.f28, _this.f29, p0);
          _1520_v106 = _nw254;
          let _1521_v107;
          _1521_v107 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_1520_v106);
          let _1522_v108;
          _1522_v108 = _dafny.Seq.of(_1520_v106.f29, true, _1520_v106.f28, _this.f28, (_1504_v92)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_1504_v92).length))]);
          let _1523_v109;
          _1523_v109 = _dafny.MultiSet.fromElements((_1520_v106).f32, _1522_v108);
          let _rhs188 = _1521_v107;
          let _rhs189 = _dafny.Seq.update((_1520_v106).f32, _module.__default.safeIndex((_this).f34, new BigNumber(((_1520_v106).f32).length)), !_dafny.Seq.contains((_1502_v91)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length))], _this.f36));
          let _rhs190 = (_module.__default.fm52(_module.D8.create_DC21((_1500_cf40).update(_module.__default.fm36((_1520_v106).f30, new BigNumber((_1519_v105).cardinality()), _this.f29, globalState), _1520_v106.f29)), globalState)).IsSubsetOf((_dafny.Set.fromElements(_1522_v108, (_this).f32, (_1520_v106).f32)).Difference(function () {
            let _coll41 = new _dafny.Set();
            for (const _compr_41 of (_1523_v109).Elements) {
              let _1524_v110 = _compr_41;
              if ((_1523_v109).contains(_1524_v110)) {
                _coll41.add(_1524_v110);
              }
            }
            return _coll41;
          }()));
          let _lhs151 = _1520_v106;
          _1521_v107 = _rhs188;
          _1522_v108 = _rhs189;
          _lhs151.f29 = _rhs190;
          _1516_v102 = (_1516_v102).update(!((((_1515_v101).contains((_this).f34)) ? ((_1515_v101).get((_this).f34)) : (!(_this.f28)))), p0);
        } else {
          let _1525_v111;
          let _nw255 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _1525_v111 = _nw255;
          let _index192 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1525_v111).length));
          (_1525_v111)[_index192] = (_this).f30;
          let _index193 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1525_v111).length));
          (_1525_v111)[_index193] = ((_this).f34).minus(_module.__default.safeDivisionInt(new BigNumber(((_this).f32).length), (_this).f34));
          let _1526_v112;
          _1526_v112 = _dafny.Set.fromElements(((p0) ? (_this.f29) : (_this.f28)));
          let _1527_v113;
          _1527_v113 = _dafny.MultiSet.fromElements(_this.f28);
          let _1528_v114;
          _1528_v114 = _dafny.Seq.of((((_1527_v113).contains(_this.f35)) ? ((_1527_v113).get(_this.f35)) : ((_this).f34)), (_1525_v111)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_1525_v111).length))]);
          _1526_v112 = _dafny.Set.fromElements((new BigNumber((_1528_v114).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("pk")).length)), _this.f28, false);
          let _1529_v115;
          let _out57;
          _out57 = _module.__default.m0(_1528_v114, _this.f29, _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(-105)), (_this).f34), _this.f35, globalState);
          _1529_v115 = _out57;
          let _1530_v116;
          _1530_v116 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(66),_this.f28);
          let _1531_v117;
          _1531_v117 = _dafny.MultiSet.fromElements((_1525_v111)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_1525_v111).length))]);
          let _1532_v118;
          _1532_v118 = _dafny.Map.Empty.slice().updateUnsafe((_1525_v111)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_1525_v111).length))],(_this).f30);
          let _1533_v119;
          let _nw256 = new _module.C4();
          _nw256.__ctor(new BigNumber(((_1502_v91)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length))]).length), (true) && ((((_1530_v116).contains(new BigNumber(((_this).f32).length))) ? ((_1530_v116).get(new BigNumber(((_this).f32).length))) : ((_this).fm7(_1531_v117, _1529_v115, globalState)))), (_this).f32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(846)), function (_1534_i9) {
            return (_this).f32;
          }), new BigNumber((_1527_v113).cardinality()), _this.f35, _this.f29, (new BigNumber((_1532_v118).length)).isLessThanOrEqualTo(new BigNumber(((_1502_v91)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length))]).length)));
          _1533_v119 = _nw256;
          let _1535_v120;
          let _nw257 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1535_v120 = _nw257;
          let _1536_v121;
          let _nw258 = new _module.C0();
          _nw258.__ctor(_1535_v120);
          _1536_v121 = _nw258;
          let _index194 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length));
          let _rhs191 = _dafny.Seq.Concat(_1501_v90, _dafny.Seq.UnicodeFromString("gobhvtjm"));
          let _rhs192 = _1533_v119;
          let _rhs193 = (_module.D1.create_DC3(((_this).f32)[_module.__default.safeIndex((_1525_v111)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_1525_v111).length))], new BigNumber(((_this).f32).length))], new BigNumber(302), _1536_v121, new BigNumber(((_this).f32).length))).dtor_cf6;
          let _lhs152 = _1502_v91;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length));
          _lhs152[_lhs153] = _rhs191;
          _1533_v119 = _rhs192;
          r3 = _rhs193;
          let _index195 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1525_v111).length));
          (_1525_v111)[_index195] = new BigNumber(((_1526_v112).Intersect((_1526_v112).Difference(_1526_v112))).length);
        }
        let _1537_v122;
        let _nw259 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _1537_v122 = _nw259;
        let _index196 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_1537_v122).length));
        (_1537_v122)[_index196] = (_this).f34;
        let _index197 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_1537_v122).length));
        let _rhs194 = ((new BigNumber(-486)).multipliedBy(new BigNumber((_1501_v90).length))).minus(_module.__default.safeDivisionInt((_this).f34, (_this).f34));
        let _rhs195 = _this.f29;
        let _rhs196 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1501_v90, (_1502_v91)[_module.__default.safeIndex(new BigNumber(600), new BigNumber((_1502_v91).length))]), _1501_v90);
        let _rhs197 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f30), (_dafny.ZERO).minus((_this).f30));
        let _lhs154 = _this;
        let _lhs155 = _1537_v122;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_1537_v122).length));
        r2 = _rhs194;
        _lhs154.f35 = _rhs195;
        _1501_v90 = _rhs196;
        _lhs155[_lhs156] = _rhs197;
        (_this).f36 = new _dafny.CodePoint('a'.codePointAt(0));
      } else {
        let _1538___mcc_h11 = (_source22).cf51;
        let _1539_cf51 = _1538___mcc_h11;
        (_this).f35 = !(_this.f31) || (p0);
        r1 = !(true) || (!(p0));
        (_this).f28 = !(p0);
        (globalState).f26 = p0;
      }
      if (((_this).f30).isEqualTo((_this).f34)) {
        let _1540_v123;
        _1540_v123 = _dafny.Seq.of((_this).f34, (_this).f30, (_this).f30);
        let _1541_v124;
        let _out58;
        _out58 = _module.__default.m0(_1540_v123, !(p0), (_this).f30, _this.f35, globalState);
        _1541_v124 = _out58;
        let _1542_v125;
        _1542_v125 = _dafny.Seq.UnicodeFromString("lsw");
        _1542_v125 = _1542_v125;
        let _1543_v126;
        _1543_v126 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,(_this).f34);
        let _1544_v127;
        _1544_v127 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), function (_1545_i10) {
          return (_this).f34;
        }));
        let _1546_v128;
        _1546_v128 = _module.D14.create_DC36((_this).f33);
        let _1547_v129;
        _1547_v129 = _dafny.MultiSet.fromElements(_1541_v124);
        let _1548_v130;
        _1548_v130 = _dafny.Seq.of(_1547_v129, _dafny.MultiSet.fromElements((_this).f30, (_this).f34));
        let _1549_v131;
        let _nw260 = new _module.C5();
        _nw260.__ctor(_this.f31, (((_1543_v126).contains(new BigNumber(((_this).f32).length))) ? ((_1543_v126).get(new BigNumber(((_this).f32).length))) : (_module.__default.fm32(true, globalState))), (_this).f30, _dafny.Seq.IsProperPrefixOf(_1540_v123, (_1544_v127)[_module.__default.safeIndex((_this).f30, new BigNumber((_1544_v127).length))]), (_this).f32, (_1546_v128).dtor_cf67, _module.__default.safeDivisionInt((_this).f34, _module.__default.fm46(new BigNumber(((_1548_v130)[_module.__default.safeIndex((_this).f34, new BigNumber((_1548_v130).length))]).cardinality()), (_this).f34, false, globalState)), _module.__default.fm1(true, (_1540_v123)[_module.__default.safeIndex(new BigNumber((_1540_v123).length), new BigNumber((_1540_v123).length))], globalState), ((_this).f32)[_module.__default.safeIndex(_1541_v124, new BigNumber(((_this).f32).length))], (!(_this.f28)) || (p0));
        _1549_v131 = _nw260;
        let _1550_v132;
        _1550_v132 = _module.D4.create_DC11(new BigNumber((_1542_v125).length), _this.f29, (_1549_v131).f48);
        let _source25 = _1550_v132;
        if (_source25.is_DC11) {
          let _1551___mcc_h24 = (_source25).cf23;
          let _1552___mcc_h25 = (_source25).cf24;
          let _1553___mcc_h26 = (_source25).cf25;
          let _1554_cf25 = _1553___mcc_h26;
          let _1555_cf24 = _1552___mcc_h25;
          let _1556_cf23 = _1551___mcc_h24;
          let _1557_v133;
          let _nw261 = Array((new BigNumber(10)).toNumber());
          _nw261[(_dafny.ZERO).toNumber()] = _this.f36;
          _nw261[(_dafny.ONE).toNumber()] = _this.f36;
          _nw261[(new BigNumber(2)).toNumber()] = _this.f36;
          _nw261[(new BigNumber(3)).toNumber()] = _this.f36;
          _nw261[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
          _nw261[(new BigNumber(5)).toNumber()] = _this.f36;
          _nw261[(new BigNumber(6)).toNumber()] = _this.f36;
          _nw261[(new BigNumber(7)).toNumber()] = _this.f36;
          _nw261[(new BigNumber(8)).toNumber()] = _this.f36;
          _nw261[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
          _1557_v133 = _nw261;
          let _1558_v134;
          let _nw262 = new _module.C0();
          _nw262.__ctor(_1557_v133);
          _1558_v134 = _nw262;
          let _1559_v135;
          _1559_v135 = _dafny.Map.Empty.slice().updateUnsafe(_1542_v125,_1558_v134);
          _1559_v135 = _dafny.Map.Empty.slice().updateUnsafe(_1542_v125,(((_1549_v131).f47) ? (_1558_v134) : ((_module.D1.create_DC3(p0, (_dafny.ZERO).minus((_this).f34), _1558_v134, _1541_v124)).dtor_cf5)));
          let _1560_v136;
          _1560_v136 = _dafny.Map.Empty.slice().updateUnsafe((_1549_v131).f47,true);
          _1560_v136 = (_1560_v136).update(_this.f28, (_1541_v124).isLessThan(_1556_cf23));
          _1542_v125 = _dafny.Seq.Concat(_1542_v125, _1542_v125);
          (globalState).f1 = ((_this).f32)[_module.__default.safeIndex(_1554_cf25, new BigNumber(((_this).f32).length))];
        } else {
          let _1561___mcc_h27 = (_source25).cf22;
          let _1562_cf22 = _1561___mcc_h27;
          let _1563_v137;
          _1563_v137 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_this.f31);
          let _1564_v139;
          let _nw263 = new _module.C4();
          _nw263.__ctor(_1541_v124, true, _dafny.Seq.of(!(_this.f29)), _dafny.Seq.Concat(_dafny.Seq.update((_this).f33, _module.__default.safeIndex(_1541_v124, new BigNumber(((_this).f33).length)), (_this).f32), _dafny.Seq.of((_this).f32, (_this).f32, (_this).fm5((_1549_v131).f48, (_dafny.ZERO).minus(new BigNumber((_1563_v137).length)), _this.f28, _1541_v124, globalState))), (_this).f34, !(function () {
            let _coll42 = new _dafny.Map();
            for (const _compr_42 of _dafny.IntegerRange(new BigNumber(910), new BigNumber(734))) {
              let _1565_v138 = _compr_42;
              if (((new BigNumber(910)).isLessThanOrEqualTo(_1565_v138)) && ((_1565_v138).isLessThan(new BigNumber(734)))) {
                _coll42.push([_module.__default.safeModuloInt(_1565_v138, (_this).f34),new BigNumber(341)]);
              }
            }
            return _coll42;
          }()).contains((_this).f30), (_1549_v131).f47, _this.f29);
          _1564_v139 = _nw263;
          r1 = ((_this).f34).isLessThanOrEqualTo((_1549_v131).f48);
          _1542_v125 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1542_v125, _dafny.Seq.UnicodeFromString("mhye")), _1542_v125);
          let _1566_v140;
          let _init33 = ((_1567_p0) => function (_1568_i11) {
            return _1567_p0;
          })(p0);
          let _nw264 = Array((new BigNumber(24)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw264.length); _i0_33++) {
            _nw264[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1566_v140 = _nw264;
          let _1569_v141;
          _1569_v141 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_1549_v131).f47);
          let _1570_v142;
          _1570_v142 = _dafny.MultiSet.fromElements(_this.f29);
          let _index198 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1566_v140).length));
          (_1566_v140)[_index198] = (new BigNumber((_1569_v141).length)).isLessThanOrEqualTo(new BigNumber((_1570_v142).cardinality()));
          let _index199 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1566_v140).length));
          (_1566_v140)[_index199] = (((_this).f30).isEqualTo(new BigNumber(166))) === ((_1549_v131).f47);
        }
        let _1571_v143;
        let _init34 = ((_1572_v125) => function (_1573_i12) {
          return _1572_v125;
        })(_1542_v125);
        let _nw265 = Array((new BigNumber(28)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw265.length); _i0_34++) {
          _nw265[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1571_v143 = _nw265;
        let _rhs198 = (_this).f30;
        let _rhs199 = _1541_v124;
        let _rhs200 = _1571_v143;
        let _rhs201 = (_this).f34;
        let _rhs202 = _dafny.Seq.Concat((_1549_v131).fm3(_this.f29, (_1549_v131).f47, globalState), _1542_v125);
        let _lhs157 = globalState;
        r3 = _rhs198;
        _lhs157.f13 = _rhs199;
        _1571_v143 = _rhs200;
        r2 = _rhs201;
        _1542_v125 = _rhs202;
      } else {
        let _1574_v144;
        _1574_v144 = _dafny.Seq.of((_this).f34);
        let _1575_v145;
        let _out59;
        _out59 = _module.__default.m0(_dafny.Seq.Concat(_1574_v144, _dafny.Seq.of((_this).f30)), _this.f29, (_this).f34, p0, globalState);
        _1575_v145 = _out59;
        r2 = _1575_v145;
        let _1576_v146;
        _1576_v146 = _dafny.MultiSet.fromElements((_this).f30, _1575_v145, (_this).f30, _1575_v145);
        let _1577_v147;
        _1577_v147 = _module.D4.create_DC10(_1576_v146);
        let _1578_v148;
        _1578_v148 = _dafny.Seq.of(_1576_v146, (_1577_v147).dtor_cf22, (_1576_v146).Difference(_1576_v146));
        let _1579_v149;
        let _init35 = function (_1580_i13) {
          return !(_this.f31);
        };
        let _nw266 = Array((new BigNumber(22)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw266.length); _i0_35++) {
          _nw266[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1579_v149 = _nw266;
        let _rhs203 = _1579_v149;
        let _rhs204 = new BigNumber(354);
        let _rhs205 = _module.__default.fm32(!(p0) || (p0), globalState);
        let _rhs206 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1578_v148, _1578_v148), _1578_v148);
        let _lhs158 = globalState;
        let _lhs159 = globalState;
        _lhs158.f21 = _rhs203;
        _lhs159.f13 = _rhs204;
        _1575_v145 = _rhs205;
        _1578_v148 = _rhs206;
        _1575_v145 = _1575_v145;
        r3 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_this.f29)).cardinality())), ((_this).f34).multipliedBy((_this).f34));
      }
      let _1581_v150;
      let _init36 = function (_1582_i14) {
        return _this.f36;
      };
      let _nw267 = Array((new BigNumber(2)).toNumber());
      for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw267.length); _i0_36++) {
        _nw267[_i0_36] = _init36(new BigNumber(_i0_36));
      }
      _1581_v150 = _nw267;
      let _1583_v151;
      let _nw268 = new _module.C0();
      _nw268.__ctor(_1581_v150);
      _1583_v151 = _nw268;
      let _1584_v152;
      _1584_v152 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f35);
      let _1585_v153;
      _1585_v153 = _dafny.Seq.of(_1584_v152, _1584_v152, _1584_v152, _1584_v152, _1584_v152);
      let _1586_v154;
      _1586_v154 = _module.D1.create_DC3(_this.f29, (_this).f30, _1583_v151, ((_this).f30).plus(new BigNumber((_1585_v153).length)));
      let _source26 = _1586_v154;
      if (_source26.is_DC3) {
        let _1587___mcc_h28 = (_source26).cf3;
        let _1588___mcc_h29 = (_source26).cf4;
        let _1589___mcc_h30 = (_source26).cf5;
        let _1590___mcc_h31 = (_source26).cf6;
        let _1591_cf6 = _1590___mcc_h31;
        let _1592_cf5 = _1589___mcc_h30;
        let _1593_cf4 = _1588___mcc_h29;
        let _1594_cf3 = _1587___mcc_h28;
        let _1595_v155;
        let _nw269 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1595_v155 = _nw269;
        let _1596_v156;
        _1596_v156 = _dafny.Seq.UnicodeFromString("xiirlaq");
        let _index200 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length));
        (_1595_v155)[_index200] = _dafny.Seq.Concat(_1596_v156, _1596_v156);
        let _1597_v157;
        _1597_v157 = _dafny.Seq.of((_this).f30);
        let _index201 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length));
        let _rhs207 = ((_this).f32)[_module.__default.safeIndex((_1597_v157)[_module.__default.safeIndex((_this).f34, new BigNumber((_1597_v157).length))], new BigNumber(((_this).f32).length))];
        let _rhs208 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), function (_1598_i15) {
          return _this.f36;
        });
        let _rhs209 = _this.f28;
        let _lhs160 = _1595_v155;
        let _lhs161 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length));
        let _lhs162 = _this;
        r0 = _rhs207;
        _lhs160[_lhs161] = _rhs208;
        _lhs162.f29 = _rhs209;
        let _1599_v158;
        _1599_v158 = _dafny.Set.fromElements(((_this).f30).isEqualTo(_1593_cf4), true);
        let _1600_v159;
        _1600_v159 = _dafny.Set.fromElements(_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f30))));
        let _rhs210 = ((_dafny.ZERO).minus(_1593_cf4)).multipliedBy(new BigNumber(((_dafny.Set.fromElements(_1597_v157)).Intersect(_1600_v159)).length));
        let _rhs211 = !(!((_this).f30).isEqualTo(((_this).f34).plus(_1591_cf6)));
        let _rhs212 = _this.f36;
        let _rhs213 = (_1599_v158).Difference(_1599_v158);
        let _rhs214 = ((_this).f34).isLessThanOrEqualTo(new BigNumber((_1599_v158).length));
        let _lhs163 = globalState;
        let _lhs164 = globalState;
        let _lhs165 = _this;
        let _lhs166 = _this;
        _lhs163.f13 = _rhs210;
        _lhs164.f26 = _rhs211;
        _lhs165.f36 = _rhs212;
        _1599_v158 = _rhs213;
        _lhs166.f29 = _rhs214;
        let _1601_v160;
        let _nw270 = new _module.C2();
        _nw270.__ctor(_1591_cf6, _module.__default.fm32(!(p0), globalState), !(_this.f35), ((_this).f32)[_module.__default.safeIndex(_1593_cf4, new BigNumber(((_this).f32).length))], _this.f31);
        _1601_v160 = _nw270;
        let _1602_v161;
        _1602_v161 = _module.D9.create_DC26(_1601_v160);
        let _1603_v162;
        _1603_v162 = _module.D9.create_DC26((_1602_v161).dtor_cf52);
        let _source27 = _1602_v161;
        if (_source27.is_DC27) {
          let _1604___mcc_h34 = (_source27).cf53;
          let _1605___mcc_h35 = (_source27).cf54;
          let _1606_cf54 = _1605___mcc_h35;
          let _1607_cf53 = _1604___mcc_h34;
          let _1608_v163;
          let _nw271 = new _module.C4();
          _nw271.__ctor((_1601_v160).f46, _1594_cf3, (_this).f32, (_this).f33, new BigNumber((_1584_v152).length), _this.f35, _this.f35, _this.f29);
          _1608_v163 = _nw271;
          let _1609_v164;
          _1609_v164 = _dafny.MultiSet.fromElements(_1608_v163, _1608_v163, _1608_v163, _1608_v163, _1608_v163);
          let _1610_v165;
          _1610_v165 = _module.D10.create_DC30(_this.f36, (_1609_v164).update(_1608_v163, _module.__default.abs((_1601_v160).f46)));
          _1610_v165 = _1610_v165;
          let _1611_v166;
          let _nw272 = Array((new BigNumber(8)).toNumber()).fill(false);
          _1611_v166 = _nw272;
          let _index202 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_1611_v166).length));
          (_1611_v166)[_index202] = _1608_v163.f31;
          let _index203 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_1611_v166).length));
          (_1611_v166)[_index203] = p0;
          let _1612_v167;
          _1612_v167 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(556),_1608_v163.f28);
          _1612_v167 = _1612_v167;
          let _1613_v168;
          _1613_v168 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus(new BigNumber(-700))).isLessThanOrEqualTo((_1597_v157)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1606_cf54, _module.__default.safeIndex((_1601_v160).f46, new BigNumber((_1606_cf54).length)), _this.f36)).length), new BigNumber((_1597_v157).length))]),(_dafny.MultiSet.fromElements(!(_1608_v163.f35), false, (_1611_v166)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_1611_v166).length))])).Intersect(_dafny.MultiSet.fromElements(_1608_v163.f31)));
          let _1614_v169;
          _1614_v169 = _dafny.MultiSet.fromElements(false);
          let _1615_v170;
          _1615_v170 = _dafny.Map.Empty.slice().updateUnsafe((_1595_v155)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length))],_1591_cf6);
          let _1616_v171;
          _1616_v171 = _module.D9.create_DC27(_1583_v151, (_1595_v155)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length))]);
          let _1617_v172;
          _1617_v172 = _dafny.Seq.of(_1616_v171);
          let _rhs215 = ((_1601_v160).f46).multipliedBy(((_1601_v160).fm29(globalState)).minus(new BigNumber(564)));
          let _rhs216 = ((_1614_v169).update(_1608_v163.f29, _module.__default.abs(new BigNumber((_1615_v170).length)))).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1608_v163.f28, _this.f28, (_1611_v166)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_1611_v166).length))]));
          let _rhs217 = _1594_cf3;
          let _rhs218 = ((_1613_v168).Merge(_1613_v168)).update(!(_1608_v163.f28) || (_1608_v163.f29), _dafny.MultiSet.FromArray((_this).f32));
          let _rhs219 = new BigNumber((_dafny.Seq.update((_this).fm3(_dafny.Seq.IsPrefixOf(_1617_v172, _1617_v172), _1594_cf3, globalState), _module.__default.safeIndex((_1601_v160).f46, new BigNumber(((_this).fm3(_dafny.Seq.IsPrefixOf(_1617_v172, _1617_v172), _1594_cf3, globalState)).length)), _this.f36)).length);
          let _lhs167 = globalState;
          let _lhs168 = globalState;
          _1591_cf6 = _rhs215;
          _lhs167.f26 = _rhs216;
          _lhs168.f1 = _rhs217;
          _1613_v168 = _rhs218;
          _1591_cf6 = _rhs219;
        } else {
          let _1618___mcc_h36 = (_source27).cf52;
          let _1619_cf52 = _1618___mcc_h36;
          _1596_v156 = (_1595_v155)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length))];
          (globalState).f12 = true;
          let _1620_v173;
          _1620_v173 = _dafny.MultiSet.fromElements(_1596_v156);
          let _1621_v174;
          _1621_v174 = _dafny.Map.Empty.slice().updateUnsafe((_1620_v173).Union(_dafny.MultiSet.fromElements((_1595_v155)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length))], (_1595_v155)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length))], (_1595_v155)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length))], (_1595_v155)[_module.__default.safeIndex(new BigNumber(860), new BigNumber((_1595_v155).length))])),(_1601_v160).f46);
          _1621_v174 = (_1621_v174).update(_1620_v173, _1591_cf6);
          let _1622_v175;
          _1622_v175 = _dafny.MultiSet.fromElements(_this.f28, _this.f35);
          _1622_v175 = (_1622_v175).Difference(_1622_v175);
        }
        let _1623_v176;
        _1623_v176 = _dafny.Seq.of(_1596_v156);
        _1623_v176 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-677)), ((_1624_v156) => function (_1625_i16) {
          return _1624_v156;
        })(_1596_v156)), _dafny.Seq.of(_1596_v156));
      } else if (_source26.is_DC2) {
        let _1626___mcc_h32 = (_source26).cf2;
        let _1627_cf2 = _1626___mcc_h32;
        let _1628_v177;
        _1628_v177 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), function (_1629_i17) {
          return _this.f36;
        })).length)), (_this).f34, (_this).f30, _module.__default.fm32(!(_this.f35), globalState));
        _1628_v177 = (_1628_v177).Intersect(_1628_v177);
        (globalState).f1 = _this.f35;
        (globalState).f13 = new BigNumber(804);
        let _1630_v178;
        _1630_v178 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,(_this).f30);
        _1630_v178 = _1630_v178;
      } else {
        let _1631___mcc_h33 = (_source26).cf7;
        let _1632_cf7 = _1631___mcc_h33;
        (_this).f31 = _this.f29;
        let _1633_v179;
        _1633_v179 = _dafny.MultiSet.fromElements(true);
        let _1634_v180;
        _1634_v180 = _module.D0.create_DC1(_1633_v179);
        let _1635_v181;
        _1635_v181 = _dafny.Map.Empty.slice().updateUnsafe(_1634_v180,_this.f29);
        let _1636_v182;
        _1636_v182 = _module.D8.create_DC21((_1635_v181).Merge(_1635_v181));
        let _source28 = _1636_v182;
        if (_source28.is_DC22) {
          let _1637___mcc_h37 = (_source28).cf41;
          let _1638___mcc_h38 = (_source28).cf42;
          let _1639_cf42 = _1638___mcc_h38;
          let _1640_cf41 = _1637___mcc_h37;
          let _1641_v183;
          let _init37 = ((_1642_cf41) => function (_1643_i18) {
            return _module.__default.safeModuloInt(_1643_i18, _1642_cf41);
          })(_1640_cf41);
          let _nw273 = Array((new BigNumber(8)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw273.length); _i0_37++) {
            _nw273[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1641_v183 = _nw273;
          let _index204 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1641_v183).length));
          (_1641_v183)[_index204] = new BigNumber(593);
          let _index205 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1641_v183).length));
          (_1641_v183)[_index205] = (_1639_cf42).plus((_1639_cf42).multipliedBy((_this).f34));
          let _1644_v184;
          let _nw274 = Array((new BigNumber(25)).toNumber()).fill(false);
          _1644_v184 = _nw274;
          let _index206 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_1644_v184).length));
          (_1644_v184)[_index206] = p0;
          let _1645_v185;
          _1645_v185 = _dafny.Set.fromElements(_this.f28);
          let _1646_v186;
          _1646_v186 = _module.D7.create_DC19(_this.f31, new BigNumber(530), new BigNumber(368));
          let _index207 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_1644_v184).length));
          let _index208 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1641_v183).length));
          let _rhs220 = (_1645_v185).IsProperSubsetOf(((p0) ? (_1645_v185) : (_1645_v185)));
          let _rhs221 = _this.f31;
          let _rhs222 = (_dafny.ZERO).minus((_this).f34);
          let _rhs223 = !((_1646_v186).dtor_cf36);
          let _lhs169 = _1644_v184;
          let _lhs170 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_1644_v184).length));
          let _lhs171 = _this;
          let _lhs172 = _1641_v183;
          let _lhs173 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1641_v183).length));
          let _lhs174 = globalState;
          _lhs169[_lhs170] = _rhs220;
          _lhs171.f28 = _rhs221;
          _lhs172[_lhs173] = _rhs222;
          _lhs174.f26 = _rhs223;
          let _1647_v187;
          _1647_v187 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_this).f34));
          let _1648_v188;
          _1648_v188 = _dafny.Map.Empty.slice().updateUnsafe((_1644_v184)[_module.__default.safeIndex(new BigNumber(803), new BigNumber((_1644_v184).length))],_1640_cf41);
          (globalState).f13 = new BigNumber((((_1647_v187)[_module.__default.safeIndex(new BigNumber(-962), new BigNumber((_1647_v187).length))]).Merge((_1648_v188).Merge(_1648_v188))).length);
          let _1649_v189;
          _1649_v189 = _dafny.Seq.UnicodeFromString("mwd");
          let _rhs224 = new BigNumber(742);
          let _rhs225 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uy"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), function (_1650_i19) {
            return _this.f36;
          })), _dafny.Seq.UnicodeFromString("gvqi"));
          let _rhs226 = (_1644_v184)[_module.__default.safeIndex(new BigNumber(803), new BigNumber((_1644_v184).length))];
          let _lhs175 = globalState;
          r3 = _rhs224;
          _1649_v189 = _rhs225;
          _lhs175.f1 = _rhs226;
        } else if (_source28.is_DC23) {
          let _1651___mcc_h39 = (_source28).cf43;
          let _1652___mcc_h40 = (_source28).cf44;
          let _1653___mcc_h41 = (_source28).cf45;
          let _1654_cf45 = _1653___mcc_h41;
          let _1655_cf44 = _1652___mcc_h40;
          let _1656_cf43 = _1651___mcc_h39;
          let _1657_v190;
          _1657_v190 = _dafny.Seq.of((_this).f34);
          let _1658_v191;
          let _nw275 = new _module.C4();
          _nw275.__ctor(new BigNumber((_1633_v179).cardinality()), _this.f28, _dafny.Seq.Concat((_this).f32, ((_this).f33)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(229)), function (_1659_i20) {
            return _this.f36;
          })).length), new BigNumber(((_this).f33).length))]), (_this).f33, (_this).f30, false, _1655_cf44, _dafny.Seq.IsProperPrefixOf(_1657_v190, _dafny.Seq.update(_1657_v190, _module.__default.safeIndex(new BigNumber(526), new BigNumber((_1657_v190).length)), (_this).f30)));
          _1658_v191 = _nw275;
          let _1660_v192;
          let _nw276 = new _module.C1();
          _nw276.__ctor(false, _this.f31, _1656_cf43);
          _1660_v192 = _nw276;
          _1660_v192 = _1660_v192;
          let _1661_v193;
          _1661_v193 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,_1655_cf44);
          let _1662_v194;
          _1662_v194 = _dafny.Map.Empty.slice().updateUnsafe(_1661_v193,_1656_cf43);
          let _1663_v195;
          _1663_v195 = _dafny.Map.Empty.slice().updateUnsafe(_1662_v194,_this.f28);
          let _1664_v196;
          _1664_v196 = _dafny.Seq.UnicodeFromString("nvyflppiw");
          let _1665_v197;
          _1665_v197 = _dafny.MultiSet.fromElements((_this).f30);
          let _1666_v198;
          _1666_v198 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_this.f35);
          let _1667_v199;
          let _nw277 = Array((new BigNumber(28)).toNumber());
          _nw277[(_dafny.ZERO).toNumber()] = _this.f31;
          _nw277[(_dafny.ONE).toNumber()] = (((_1663_v195).contains(_1662_v194)) ? ((_1663_v195).get(_1662_v194)) : (_this.f29));
          _nw277[(new BigNumber(2)).toNumber()] = _this.f29;
          _nw277[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_1664_v196, _1664_v196);
          _nw277[(new BigNumber(4)).toNumber()] = (_1660_v192).f43;
          _nw277[(new BigNumber(5)).toNumber()] = _1654_cf45;
          _nw277[(new BigNumber(6)).toNumber()] = !((_1660_v192).f43) || (p0);
          _nw277[(new BigNumber(7)).toNumber()] = _1655_cf44;
          _nw277[(new BigNumber(8)).toNumber()] = _1656_cf43;
          _nw277[(new BigNumber(9)).toNumber()] = _this.f35;
          _nw277[(new BigNumber(10)).toNumber()] = _this.f28;
          _nw277[(new BigNumber(11)).toNumber()] = (_1660_v192).f43;
          _nw277[(new BigNumber(12)).toNumber()] = _this.f31;
          _nw277[(new BigNumber(13)).toNumber()] = (_1660_v192).f43;
          _nw277[(new BigNumber(14)).toNumber()] = _dafny.Seq.contains((_this).f32, _1654_cf45);
          _nw277[(new BigNumber(15)).toNumber()] = (_1656_cf43) || ((((_1584_v152).contains(_this.f28)) ? ((_1584_v152).get(_this.f28)) : ((_this).fm7(_1665_v197, (_this).f34, globalState))));
          _nw277[(new BigNumber(16)).toNumber()] = (_1660_v192).f43;
          _nw277[(new BigNumber(17)).toNumber()] = _1656_cf43;
          _nw277[(new BigNumber(18)).toNumber()] = (_1660_v192).f43;
          _nw277[(new BigNumber(19)).toNumber()] = true;
          _nw277[(new BigNumber(20)).toNumber()] = (_1660_v192).f43;
          _nw277[(new BigNumber(21)).toNumber()] = _this.f29;
          _nw277[(new BigNumber(22)).toNumber()] = _1656_cf43;
          _nw277[(new BigNumber(23)).toNumber()] = !((_1660_v192).f43);
          _nw277[(new BigNumber(24)).toNumber()] = _this.f29;
          _nw277[(new BigNumber(25)).toNumber()] = (((_1666_v198).contains((_this).f34)) ? ((_1666_v198).get((_this).f34)) : (_this.f28));
          _nw277[(new BigNumber(26)).toNumber()] = _this.f35;
          _nw277[(new BigNumber(27)).toNumber()] = true;
          _1667_v199 = _nw277;
          let _index209 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_1667_v199).length));
          (_1667_v199)[_index209] = _this.f28;
          let _index210 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_1667_v199).length));
          let _rhs227 = (((_1665_v197).contains((_this).f34)) ? ((_1665_v197).get((_this).f34)) : ((_dafny.ZERO).minus((_this).f34)));
          let _rhs228 = _1656_cf43;
          let _lhs176 = globalState;
          let _lhs177 = _1667_v199;
          let _lhs178 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_1667_v199).length));
          _lhs176.f13 = _rhs227;
          _lhs177[_lhs178] = _rhs228;
          let _index211 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_1667_v199).length));
          (_1667_v199)[_index211] = (_1667_v199)[_module.__default.safeIndex(new BigNumber(577), new BigNumber((_1667_v199).length))];
        } else if (_source28.is_DC24) {
          let _1668___mcc_h42 = (_source28).cf46;
          let _1669___mcc_h43 = (_source28).cf47;
          let _1670___mcc_h44 = (_source28).cf48;
          let _1671___mcc_h45 = (_source28).cf49;
          let _1672___mcc_h46 = (_source28).cf50;
          let _1673_cf50 = _1672___mcc_h46;
          let _1674_cf49 = _1671___mcc_h45;
          let _1675_cf48 = _1670___mcc_h44;
          let _1676_cf47 = _1669___mcc_h43;
          let _1677_cf46 = _1668___mcc_h42;
          let _1678_v200;
          let _nw278 = Array((new BigNumber(2)).toNumber());
          _nw278[(_dafny.ZERO).toNumber()] = _1674_cf49;
          _nw278[(_dafny.ONE).toNumber()] = ((_this).f30).isLessThanOrEqualTo(_1676_cf47);
          _1678_v200 = _nw278;
          let _index212 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1678_v200).length));
          (_1678_v200)[_index212] = _this.f28;
          let _index213 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1678_v200).length));
          (_1678_v200)[_index213] = !(_this.f35);
          let _1679_v201;
          let _init38 = function (_1680_i21) {
            return _module.__default.safeModuloInt(_1680_i21, (_this).f30);
          };
          let _nw279 = Array((_dafny.ONE).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw279.length); _i0_38++) {
            _nw279[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1679_v201 = _nw279;
          let _index214 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1678_v200).length));
          let _rhs229 = _1674_cf49;
          let _rhs230 = (_module.__default.safeModuloInt(_1676_cf47, new BigNumber(22))).minus((new BigNumber((_1633_v179).cardinality())).plus(_1676_cf47));
          let _rhs231 = _1679_v201;
          let _rhs232 = (new BigNumber(645)).isLessThanOrEqualTo(_1676_cf47);
          let _lhs179 = _1678_v200;
          let _lhs180 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1678_v200).length));
          _1677_cf46 = _rhs229;
          r3 = _rhs230;
          _1679_v201 = _rhs231;
          _lhs179[_lhs180] = _rhs232;
          let _1681_v202;
          let _nw280 = new _module.C4();
          _nw280.__ctor((_this).f30, _this.f35, _dafny.Seq.of(_this.f31, p0, _this.f35, _1674_cf49, _1677_cf46), (_this).f33, (_this).f34, _this.f35, _this.f31, (_1678_v200)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_1678_v200).length))]);
          _1681_v202 = _nw280;
          let _1682_v203;
          _1682_v203 = _dafny.MultiSet.fromElements(_1681_v202, _1681_v202);
          let _1683_v204;
          _1683_v204 = _dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC30(_this.f36, _1682_v203),((_1681_v202).f34).minus((_dafny.ZERO).minus(_1676_cf47)));
          _1683_v204 = (_1683_v204).update(_module.D10.create_DC30(_this.f36, _dafny.MultiSet.fromElements(_1681_v202)), _1676_cf47);
          let _1684_v205;
          _1684_v205 = _dafny.Seq.of(_1583_v151.f37, _1583_v151.f37, _1583_v151.f37);
          _1673_cf50 = _dafny.Seq.of(new BigNumber((_1684_v205).length));
        } else if (_source28.is_DC21) {
          let _1685___mcc_h47 = (_source28).cf40;
          let _1686_cf40 = _1685___mcc_h47;
          let _1687_v206;
          _1687_v206 = _dafny.Seq.UnicodeFromString("jkeeiw");
          (globalState).f12 = _dafny.Seq.contains(_1687_v206, _this.f36);
          let _1688_v207;
          _1688_v207 = _dafny.MultiSet.fromElements((_this).f30);
          let _1689_v208;
          _1689_v208 = _dafny.Set.fromElements(new BigNumber((_1584_v152).length), (_this).f34, (_this).f30);
          let _1690_v209;
          _1690_v209 = _dafny.Seq.of((_this).f34, new BigNumber((_1633_v179).cardinality()), (_this).f30, new BigNumber((_1633_v179).cardinality()), (_this).f30);
          let _1691_v210;
          _1691_v210 = _dafny.Set.fromElements(!(!(_this.f28)), _this.f28, false);
          let _1692_v211;
          let _nw281 = Array((new BigNumber(8)).toNumber());
          _nw281[(_dafny.ZERO).toNumber()] = (_1688_v207).contains((_this).f34);
          _nw281[(_dafny.ONE).toNumber()] = _dafny.areEqual(_module.__default.fm42(_1687_v206, (((_1688_v207).contains(new BigNumber((_1689_v208).length))) ? ((_1688_v207).get(new BigNumber((_1689_v208).length))) : ((_this).f30)), _dafny.Seq.UnicodeFromString("yvyovs"), globalState), _1690_v209);
          _nw281[(new BigNumber(2)).toNumber()] = _this.f31;
          _nw281[(new BigNumber(3)).toNumber()] = (_1691_v210).IsDisjointFrom(_1691_v210);
          _nw281[(new BigNumber(4)).toNumber()] = _this.f31;
          _nw281[(new BigNumber(5)).toNumber()] = _this.f29;
          _nw281[(new BigNumber(6)).toNumber()] = p0;
          _nw281[(new BigNumber(7)).toNumber()] = _this.f28;
          _1692_v211 = _nw281;
          let _index215 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_1692_v211).length));
          (_1692_v211)[_index215] = ((_module.__default.fm53(new _dafny.CodePoint('v'.codePointAt(0)), (_this).f30, globalState)).dtor_cf62).isLessThan((_this).f30);
          let _index216 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_1692_v211).length));
          (_1692_v211)[_index216] = _module.__default.fm1(p0, (_this).f30, globalState);
          let _1693_v212;
          let _nw282 = Array((new BigNumber(13)).toNumber()).fill([]);
          _1693_v212 = _nw282;
          let _1694_v213;
          let _nw283 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1694_v213 = _nw283;
          let _index217 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_1693_v212).length));
          (_1693_v212)[_index217] = _1694_v213;
          let _index218 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_1693_v212).length));
          let _rhs233 = _1694_v213;
          let _rhs234 = (_this).f30;
          let _rhs235 = (_this).f34;
          let _rhs236 = (_this).fm6(globalState);
          let _lhs181 = _1693_v212;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_1693_v212).length));
          let _lhs183 = globalState;
          let _lhs184 = globalState;
          let _lhs185 = _this;
          _lhs181[_lhs182] = _rhs233;
          _lhs183.f13 = _rhs234;
          _lhs184.f13 = _rhs235;
          _lhs185.f35 = _rhs236;
          let _1695_v214;
          _1695_v214 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f30);
          _1695_v214 = (_1695_v214).update(_this.f35, (_this).f34);
        } else {
          let _1696___mcc_h48 = (_source28).cf51;
          let _1697_cf51 = _1696___mcc_h48;
          (globalState).f13 = (_this).f34;
          (globalState).f1 = p0;
          (globalState).f13 = _module.__default.safeModuloInt(((_this).f34).minus(new BigNumber(669)), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-218))));
          let _1698_v215;
          _1698_v215 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,p0);
          let _rhs237 = !(p0);
          let _rhs238 = (_dafny.ZERO).minus(_module.__default.fm46((_dafny.ZERO).minus((_this).f34), (_this).f34, (((_1698_v215).contains((_this).f30)) ? ((_1698_v215).get((_this).f30)) : (p0)), globalState));
          let _lhs186 = _this;
          _lhs186.f31 = _rhs237;
          r3 = _rhs238;
        }
        let _1699_v216;
        _1699_v216 = _dafny.Seq.UnicodeFromString("ylj");
        _1699_v216 = _1699_v216;
        let _rhs239 = (_this).f34;
        let _rhs240 = (_this).f30;
        let _lhs187 = globalState;
        r2 = _rhs239;
        _lhs187.f13 = _rhs240;
      }
      let _1700_v217;
      let _nw284 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _1700_v217 = _nw284;
      let _1701_v218;
      let _nw285 = Array((new BigNumber(24)).toNumber()).fill(false);
      _1701_v218 = _nw285;
      let _1702_v219;
      _1702_v219 = _dafny.Seq.of(_1701_v218, _1701_v218, _1701_v218, _1701_v218);
      let _index219 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_1700_v217).length));
      (_1700_v217)[_index219] = new BigNumber((_1702_v219).length);
      let _index220 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_1700_v217).length));
      (_1700_v217)[_index220] = (_this).f30;
      r2 = (_1700_v217)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_1700_v217).length))];
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1700_v217).length))) {
        let _1703_i22 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1703_i22)) && ((_1703_i22).isLessThan(new BigNumber((_1700_v217).length))))) {
          (_1700_v217)[(_1703_i22)] = (_1703_i22).multipliedBy((_this).f30);
        }
      }
      r0 = true;
      r1 = _this.f31;
      let _1704_v220;
      _1704_v220 = _dafny.Set.fromElements((_1700_v217)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_1700_v217).length))]);
      let _1705_v221;
      _1705_v221 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_1704_v220);
      let _1706_v222;
      _1706_v222 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1705_v221).Merge(_1705_v221));
      r2 = (_dafny.ZERO).minus(new BigNumber(((((_1706_v222).contains(_this.f31)) ? ((_1706_v222).get(_this.f31)) : ((_1705_v221).Merge(_1705_v221)))).length));
      let _1707_v223;
      _1707_v223 = _dafny.Seq.UnicodeFromString("urgiph");
      let _1708_v224;
      _1708_v224 = _dafny.Seq.of((_this).f34, new BigNumber((_1707_v223).length), ((_1700_v217)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_1700_v217).length))]).multipliedBy((_1700_v217)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_1700_v217).length))]));
      r3 = (_1708_v224)[_module.__default.safeIndex((_1700_v217)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_1700_v217).length))], new BigNumber((_1708_v224).length))];
      return [r0, r1, r2, r3];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f28 = false;
      this._f29 = false;
      this._f31 = false;
      this._f30 = _dafny.ZERO;
      this._f41 = _dafny.ZERO;
      this._f42 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    set f31(value) {
      let _this = this;
      _this._f31 = value;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    __ctor(f41, f42, f30, f31, f28, f29) {
      let _this = this;
      (_this)._f41 = f41;
      (_this)._f42 = f42;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm3(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(557)), function (_1709_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("lytqd")), ((_this.f28) ? (_dafny.Seq.UnicodeFromString("gig")) : (_dafny.Seq.UnicodeFromString("cbkwbpt"))));
    };
    fm2(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_this).f41);
    };
    fm19(globalState) {
      let _this = this;
      return _this.f29;
    };
    fm20(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f30));
    };
    m7(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Seq.UnicodeFromString("");
      (_this).f29 = !(_this.f31) || (p2);
      (_this).f31 = p2;
      let _1710_v0;
      _1710_v0 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,false),(_this).f41);
      let _hi8 = (_this).f41;
      for (let _1711_i0 = new BigNumber((_1710_v0).length); _1711_i0.isLessThan(_hi8); _1711_i0 = _1711_i0.plus(_dafny.ONE)) {
        let _1712_v1;
        _1712_v1 = _dafny.Seq.UnicodeFromString("fiw");
        r3 = _dafny.Seq.Concat((_this).fm3(p0, _this.f31, globalState), _dafny.Seq.Concat(_1712_v1, _dafny.Seq.UnicodeFromString("phvd")));
        (globalState).f13 = ((_this).f30).minus(_1711_i0);
        r2 = new BigNumber(445);
        let _1713_v2;
        _1713_v2 = _dafny.MultiSet.fromElements((_this).fm20((_this).f41, false, globalState), (_this).f30, _1711_i0);
        _1713_v2 = (_this).f42;
      }
      let _1714_v3;
      _1714_v3 = _dafny.Seq.of((_this).f41);
      let _1715_v4;
      _1715_v4 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),new BigNumber(-371));
      (globalState).f13 = (_1714_v3)[_module.__default.safeIndex((((_1715_v4).contains(_this.f29)) ? ((_1715_v4).get(_this.f29)) : ((_this).f41)), new BigNumber((_1714_v3).length))];
      let _1716_v5;
      _1716_v5 = new _dafny.CodePoint('r'.codePointAt(0));
      let _1717_v6;
      _1717_v6 = _dafny.Seq.UnicodeFromString("pdeidmrnl");
      let _1718_v7;
      let _nw286 = Array((new BigNumber(28)).toNumber());
      _nw286[(_dafny.ZERO).toNumber()] = _1716_v5;
      _nw286[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
      _nw286[(new BigNumber(2)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(3)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(4)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(5)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(6)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(7)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(8)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(9)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
      _nw286[(new BigNumber(11)).toNumber()] = _module.__default.fm21(_1716_v5, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f41,false)).length), globalState);
      _nw286[(new BigNumber(12)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(13)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(14)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(15)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(16)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(17)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(18)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(19)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(20)).toNumber()] = _module.__default.fm21(_module.__default.fm21(_1716_v5, (_this).f30, globalState), (_this).f30, globalState);
      _nw286[(new BigNumber(21)).toNumber()] = (_1717_v6)[_module.__default.safeIndex((_this).f30, new BigNumber((_1717_v6).length))];
      _nw286[(new BigNumber(22)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
      _nw286[(new BigNumber(23)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(24)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(25)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(26)).toNumber()] = _1716_v5;
      _nw286[(new BigNumber(27)).toNumber()] = _1716_v5;
      _1718_v7 = _nw286;
      let _1719_v8;
      let _nw287 = new _module.C0();
      _nw287.__ctor(_1718_v7);
      _1719_v8 = _nw287;
      let _1720_v9;
      _1720_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f41,(_this).f30);
      let _1721_v10;
      let _nw288 = Array((new BigNumber(20)).toNumber());
      _nw288[(_dafny.ZERO).toNumber()] = (_this).f41;
      _nw288[(_dafny.ONE).toNumber()] = ((p0) ? ((_dafny.ZERO).minus((_this).f41)) : (new BigNumber(277)));
      _nw288[(new BigNumber(2)).toNumber()] = (_this).f41;
      _nw288[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_this).f41);
      _nw288[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).f41);
      _nw288[(new BigNumber(5)).toNumber()] = (_this).fm20((_this).f41, !(true), globalState);
      _nw288[(new BigNumber(6)).toNumber()] = (_this).f30;
      _nw288[(new BigNumber(7)).toNumber()] = (_this).f30;
      _nw288[(new BigNumber(8)).toNumber()] = (_this).f30;
      _nw288[(new BigNumber(9)).toNumber()] = (_this).fm20((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f30)), p0, globalState);
      _nw288[(new BigNumber(10)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_1720_v9).length))).multipliedBy((_this).f41);
      _nw288[(new BigNumber(11)).toNumber()] = (_this).fm20((_this).f41, _this.f31, globalState);
      _nw288[(new BigNumber(12)).toNumber()] = (_this).f30;
      _nw288[(new BigNumber(13)).toNumber()] = ((_this).f41).plus((_this).fm20((_this).f41, p0, globalState));
      _nw288[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((((_1715_v4).contains(_this.f28)) ? ((_1715_v4).get(_this.f28)) : ((_this).f30)));
      _nw288[(new BigNumber(15)).toNumber()] = (_this).f30;
      _nw288[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_this).f41)).length);
      _nw288[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((_this).f41);
      _nw288[(new BigNumber(18)).toNumber()] = (_this).f30;
      _nw288[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt((_this).fm20((_this).f30, (_this).fm19(globalState), globalState), (_this).f41);
      _1721_v10 = _nw288;
      let _index221 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_1721_v10).length));
      (_1721_v10)[_index221] = _module.__default.safeDivisionInt((_this).f30, (_dafny.ZERO).minus((_this).f30));
      let _index222 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((p1).length));
      (p1)[_index222] = p0;
      let _index223 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1721_v10).length));
      (_1721_v10)[_index223] = (_this).f41;
      let _1722_v11;
      _1722_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f41,p2);
      let _1723_v12;
      _1723_v12 = _dafny.MultiSet.fromElements(_this.f28);
      let _index224 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_1721_v10).length));
      let _index225 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((p1).length));
      let _index226 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1721_v10).length));
      let _rhs241 = (_this).f30;
      let _rhs242 = (_this).f41;
      let _rhs243 = !((((_this).f30).minus(new BigNumber((_1722_v11).length))).isLessThan((_this).f30));
      let _rhs244 = (_dafny.MultiSet.fromElements(p2)).IsProperSubsetOf(_1723_v12);
      let _rhs245 = ((_this).f30).minus((_this).f30);
      let _lhs188 = globalState;
      let _lhs189 = _1721_v10;
      let _lhs190 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_1721_v10).length));
      let _lhs191 = p1;
      let _lhs192 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((p1).length));
      let _lhs193 = _this;
      let _lhs194 = _1721_v10;
      let _lhs195 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1721_v10).length));
      _lhs188.f13 = _rhs241;
      _lhs189[_lhs190] = _rhs242;
      _lhs191[_lhs192] = _rhs243;
      _lhs193.f31 = _rhs244;
      _lhs194[_lhs195] = _rhs245;
      let _1724_v13;
      _1724_v13 = _module.D5.create_DC13(p0, _1715_v4, _1716_v5);
      let _1725_v14;
      _1725_v14 = _module.D5.create_DC13(_this.f28, (_1724_v13).dtor_cf28, _1716_v5);
      let _pat_let_tv15 = _1716_v5;
      let _pat_let_tv16 = _1716_v5;
      r0 = function (_source29) {
        if (_source29.is_DC13) {
          let _1726___mcc_h0 = (_source29).cf27;
          let _1727___mcc_h1 = (_source29).cf28;
          let _1728___mcc_h2 = (_source29).cf29;
          let _1729_cf29 = _1728___mcc_h2;
          let _1730_cf28 = _1727___mcc_h1;
          let _1731_cf27 = _1726___mcc_h0;
          return new _dafny.CodePoint('e'.codePointAt(0));
        } else if (_source29.is_DC14) {
          let _1732___mcc_h3 = (_source29).cf30;
          let _1733___mcc_h4 = (_source29).cf31;
          let _1734_cf31 = _1733___mcc_h4;
          let _1735_cf30 = _1732___mcc_h3;
          return _pat_let_tv15;
        } else {
          let _1736___mcc_h5 = (_source29).cf26;
          let _1737_cf26 = _1736___mcc_h5;
          return _pat_let_tv16;
        }
      }(_1724_v13);
      let _1738_v15;
      _1738_v15 = _dafny.Seq.of(_1720_v9, _1720_v9, _1720_v9, _dafny.Map.Empty.slice().updateUnsafe((_this).fm20((_this).f41, p0, globalState),new BigNumber((_1722_v11).length)));
      r1 = ((p0) ? (_1738_v15) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), ((_1739_v9) => function (_1740_i1) {
        return _1739_v9;
      })(_1720_v9))));
      r2 = _module.__default.safeModuloInt(new BigNumber((_1717_v6).length), (_1721_v10)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_1721_v10).length))]);
      r3 = _1717_v6;
      return [r0, r1, r2, r3];
    }
    m8(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _1741_v0;
      let _nw289 = new _module.C1();
      _nw289.__ctor(_this.f29, _this.f29, true);
      _1741_v0 = _nw289;
      let _1742_v1;
      _1742_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1741_v0,new BigNumber(161));
      let _1743_v2;
      _1743_v2 = new _dafny.CodePoint('w'.codePointAt(0));
      let _1744_v3;
      _1744_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1741_v0.f28,(_this).fm19(globalState));
      let _1745_v4;
      let _nw290 = new _module.C6();
      _nw290.__ctor(_1743_v2, (_this).f30, _this.f29, _dafny.Seq.of(_this.f29), _module.__default.fm51(_1741_v0.f29, _1743_v2, globalState), new BigNumber((_1744_v3).length), _this.f29, _this.f29, _this.f29);
      _1745_v4 = _nw290;
      let _1746_v5;
      _1746_v5 = _dafny.Map.Empty.slice().updateUnsafe(((((_1742_v1).contains(_1741_v0)) ? ((_1742_v1).get(_1741_v0)) : ((_this).f30))).isLessThanOrEqualTo((_this).f41),(new BigNumber((_dafny.Seq.of(_1745_v4, _1745_v4)).length)).plus((_this).f30));
      _1746_v5 = (_1746_v5).update(_this.f31, (_1745_v4).f30);
      let _1747_v6;
      _1747_v6 = _module.D5.create_DC13(_1745_v4.f31, _1746_v5, _1743_v2);
      let _1748_v7;
      _1748_v7 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm24(_1741_v0.f29, new BigNumber(487), globalState),(_1745_v4).f30);
      let _1749_v8;
      _1749_v8 = _dafny.Seq.UnicodeFromString("xbffckyc");
      let _1750_v9;
      let _nw291 = new _module.C5();
      _nw291.__ctor((_1747_v6).dtor_cf27, new BigNumber(593), (_this).f41, _this.f31, (_1745_v4).f32, (_1745_v4).f33, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1748_v7,(_dafny.ZERO).minus((_1745_v4).f30))).length), true, ((_1745_v4).f32)[_module.__default.safeIndex(new BigNumber(((_this).f42).cardinality()), new BigNumber(((_1745_v4).f32).length))], !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(418)), ((_1751_v2) => function (_1752_i0) {
        return _1751_v2;
      })(_1743_v2)), _1749_v8));
      _1750_v9 = _nw291;
      (_1741_v0).f29 = ((_1745_v4).f32)[_module.__default.safeIndex((_this).f30, new BigNumber(((_1745_v4).f32).length))];
      let _1753_v10;
      _1753_v10 = _dafny.Seq.of(new BigNumber((_1746_v5).length));
      let _hi9 = (_1750_v9).f48;
      for (let _1754_i1 = (new BigNumber((_1753_v10).length)).minus((_this).f30); _1754_i1.isLessThan(_hi9); _1754_i1 = _1754_i1.plus(_dafny.ONE)) {
        let _1755_v11;
        _1755_v11 = _module.D8.create_DC24((_1750_v9).f47, (_this).f41, (_1745_v4).f32, (_1750_v9).f47, _1753_v10);
        let _1756_v12;
        let _nw292 = new _module.C4();
        _nw292.__ctor(new BigNumber((_1749_v8).length), (_1750_v9).f47, (_1755_v11).dtor_cf48, _dafny.Seq.Concat((_1745_v4).f33, (_1745_v4).f33), (_1745_v4).f30, ((_1745_v4).f30).isLessThan((_1750_v9).f48), (_1750_v9).f47, _1741_v0.f29);
        _1756_v12 = _nw292;
        let _index227 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((p0).length));
        (p0)[_index227] = _this.f29;
        let _index228 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((p0).length));
        (p0)[_index228] = (_1756_v12).fm44((_dafny.ZERO).minus(_1754_i1), globalState);
        let _1757_v13;
        _1757_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1754_i1,_1741_v0.f28);
        _1757_v13 = (_1757_v13).update(new BigNumber((_dafny.Seq.update((_1745_v4).f32, _module.__default.safeIndex(new BigNumber(-685), new BigNumber(((_1745_v4).f32).length)), _1741_v0.f29)).length), _1745_v4.f29);
        let _1758_v14;
        let _init39 = function (_1759_i2) {
          return (_1759_i2).plus(new BigNumber(962));
        };
        let _nw293 = Array((new BigNumber(13)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw293.length); _i0_39++) {
          _nw293[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1758_v14 = _nw293;
        let _1760_v15;
        _1760_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v14,new BigNumber(((_1745_v4).f32).length));
        let _1761_v16;
        _1761_v16 = _module.D7.create_DC20(_module.D7.create_DC20(_module.D7.create_DC18(_1760_v15)));
        let _1762_v17;
        _1762_v17 = _module.D7.create_DC20(_1761_v16);
        let _pat_let_tv17 = _1761_v16;
        let _1763_v18;
        let _nw294 = Array((new BigNumber(19)).toNumber());
        _nw294[(_dafny.ZERO).toNumber()] = _1762_v17;
        _nw294[(_dafny.ONE).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(2)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(3)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(4)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(5)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(6)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(7)).toNumber()] = _module.D7.create_DC20(_1761_v16);
        _nw294[(new BigNumber(8)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(9)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(10)).toNumber()] = _module.D7.create_DC20(_1761_v16);
        _nw294[(new BigNumber(11)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(12)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(13)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(14)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(15)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(16)).toNumber()] = function (_pat_let23_0) {
          return function (_1764_dt__update__tmp_h0) {
            return function (_pat_let24_0) {
              return function (_1765_dt__update_hcf39_h0) {
                return _module.D7.create_DC20(_1765_dt__update_hcf39_h0);
              }(_pat_let24_0);
            }(_pat_let_tv17);
          }(_pat_let23_0);
        }(_1762_v17);
        _nw294[(new BigNumber(17)).toNumber()] = _1762_v17;
        _nw294[(new BigNumber(18)).toNumber()] = _1762_v17;
        _1763_v18 = _nw294;
        let _1766_v19;
        _1766_v19 = _1763_v18;
        let _1767_v20;
        let _nw295 = Array((new BigNumber(17)).toNumber());
        _nw295[(_dafny.ZERO).toNumber()] = (((p0)[_module.__default.safeIndex(new BigNumber(581), new BigNumber((p0).length))]) ? (_1763_v18) : (_1763_v18));
        _nw295[(_dafny.ONE).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(2)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(3)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(4)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(5)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(6)).toNumber()] = (_1766_v19);
        _nw295[(new BigNumber(7)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(8)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(9)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(10)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(11)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(12)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(13)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(14)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(15)).toNumber()] = _1763_v18;
        _nw295[(new BigNumber(16)).toNumber()] = _1763_v18;
        _1767_v20 = _nw295;
        let _1768_v21;
        let _nw296 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1768_v21 = _nw296;
        let _index229 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1768_v21).length));
        (_1768_v21)[_index229] = _1743_v2;
        let _index230 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1768_v21).length));
        let _rhs246 = _1767_v20;
        let _rhs247 = _module.__default.safeDivisionInt((_1750_v9).f48, _module.__default.safeModuloInt((_this).f41, _1754_i1));
        let _rhs248 = new _dafny.CodePoint('x'.codePointAt(0));
        let _rhs249 = _1745_v4.f29;
        let _rhs250 = _1745_v4.f31;
        let _lhs196 = _1768_v21;
        let _lhs197 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1768_v21).length));
        let _lhs198 = globalState;
        let _lhs199 = _1741_v0;
        _1767_v20 = _rhs246;
        r1 = _rhs247;
        _lhs196[_lhs197] = _rhs248;
        _lhs198.f12 = _rhs249;
        _lhs199.f29 = _rhs250;
      }
      if (_this.f29) {
        if ((_1750_v9).f47) {
          let _1769_v22;
          let _init40 = function (_1770_i3) {
            return _dafny.Seq.of(false);
          };
          let _nw297 = Array((new BigNumber(27)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw297.length); _i0_40++) {
            _nw297[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1769_v22 = _nw297;
          let _rhs251 = _1769_v22;
          let _rhs252 = (_1745_v4).f30;
          let _rhs253 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hgo"), _1749_v8);
          let _lhs200 = globalState;
          _1769_v22 = _rhs251;
          _lhs200.f13 = _rhs252;
          _1749_v8 = _rhs253;
          let _1771_v23;
          _1771_v23 = _dafny.Map.Empty.slice().updateUnsafe((_1745_v4).f30,_1746_v5);
          let _rhs254 = _module.__default.fm46(new BigNumber(((((_1771_v23).contains(new BigNumber(((_module.D8.create_DC24((_1750_v9).f47, (_1750_v9).f48, (_1745_v4).f32, _1741_v0.f29, _1753_v10)).dtor_cf48).length))) ? ((_1771_v23).get(new BigNumber(((_module.D8.create_DC24((_1750_v9).f47, (_1750_v9).f48, (_1745_v4).f32, _1741_v0.f29, _1753_v10)).dtor_cf48).length))) : (_1746_v5))).length), (_1750_v9).f48, _1741_v0.f28, globalState);
          let _rhs255 = _1743_v2;
          let _lhs201 = globalState;
          _lhs201.f13 = _rhs254;
          _1743_v2 = _rhs255;
          (globalState).f26 = false;
          let _1772_v24;
          let _nw298 = new _module.C1();
          _nw298.__ctor((_1750_v9).fm6(globalState), _1741_v0.f28, _1745_v4.f31);
          _1772_v24 = _nw298;
          let _1773_v25;
          _1773_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1772_v24,new BigNumber((_1749_v8).length));
          _1773_v25 = (_1773_v25).update(((true) ? (_1772_v24) : (_1772_v24)), (_1745_v4).f30);
          (globalState).f26 = (_1750_v9).f47;
        } else {
          let _1774_v26;
          let _nw299 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
          _1774_v26 = _nw299;
          _1774_v26 = _1774_v26;
          let _rhs256 = _1741_v0.f28;
          let _rhs257 = _1743_v2;
          let _rhs258 = _1743_v2;
          let _lhs202 = _1741_v0;
          _lhs202.f28 = _rhs256;
          _1743_v2 = _rhs257;
          _1743_v2 = _rhs258;
          let _index231 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((p0).length));
          (p0)[_index231] = (_1750_v9).f47;
          let _index232 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((p0).length));
          (p0)[_index232] = (_1750_v9).f47;
          let _1775_v27;
          let _nw300 = new _module.C1();
          _nw300.__ctor(_1741_v0.f29, false, ((_1750_v9).f48).isLessThanOrEqualTo((_1745_v4).f30));
          _1775_v27 = _nw300;
          let _1776_v28;
          let _nw301 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1776_v28 = _nw301;
          let _index233 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1776_v28).length));
          (_1776_v28)[_index233] = (_this).f41;
          let _index234 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1776_v28).length));
          (_1776_v28)[_index234] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1745_v4).f30));
        }
        let _1777_v29;
        _1777_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_dafny.Seq.Concat(_1749_v8, _1749_v8));
        _1777_v29 = ((_1777_v29).Merge(_1777_v29)).Merge(_1777_v29);
        let _1778_v30;
        let _init41 = ((_1779_v2) => function (_1780_i4) {
          return _1779_v2;
        })(_1743_v2);
        let _nw302 = Array((new BigNumber(5)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw302.length); _i0_41++) {
          _nw302[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1778_v30 = _nw302;
        let _index235 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_1778_v30).length));
        (_1778_v30)[_index235] = new _dafny.CodePoint('a'.codePointAt(0));
        let _index236 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_1778_v30).length));
        (_1778_v30)[_index236] = _1743_v2;
        let _1781_v31;
        _1781_v31 = _dafny.MultiSet.fromElements(_this.f28);
        let _1782_v32;
        _1782_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1745_v4).f30,false);
        r1 = (new BigNumber(((_1781_v31).update(_1741_v0.f29, _module.__default.abs(new BigNumber(((_1745_v4).f32).length)))).cardinality())).minus(_module.__default.safeDivisionInt((((_1746_v5).contains(!((_module.D2.create_DC6((_1745_v4).f30, true, (_1745_v4).f30, _1745_v4.f29, (_1745_v4).f30)).dtor_cf12))) ? ((_1746_v5).get(!((_module.D2.create_DC6((_1745_v4).f30, true, (_1745_v4).f30, _1745_v4.f29, (_1745_v4).f30)).dtor_cf12))) : (new BigNumber((_1782_v32).length))), (_1745_v4).f30));
        let _index237 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((p0).length));
        (p0)[_index237] = _1745_v4.f29;
        let _index238 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((p0).length));
        (p0)[_index238] = _1741_v0.f29;
      } else {
        let _1783_v33;
        _1783_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(665)), function (_1784_i5) {
          return (_this).f30;
        }),(_1745_v4).f30);
        (globalState).f13 = new BigNumber(((_1783_v33).Merge(_1783_v33)).length);
        (_1745_v4).f28 = _1745_v4.f28;
        let _1785_v34;
        _1785_v34 = _module.D4.create_DC11((_1745_v4).f30, _1741_v0.f28, (_this).f41);
        let _1786_v35;
        _1786_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1785_v34,(_this).f41);
        let _1787_v36;
        _1787_v36 = _dafny.Seq.of(_1786_v35);
        let _rhs259 = _module.__default.safeModuloInt(new BigNumber(96), (_1750_v9).f48);
        let _rhs260 = _module.__default.safeModuloInt((_dafny.ZERO).minus(((_1745_v4).f30).minus(_module.__default.fm46(new BigNumber((_1749_v8).length), new BigNumber((_dafny.Seq.UnicodeFromString("gtbdkgdg")).length), _1741_v0.f28, globalState))), (_1745_v4).f30);
        let _rhs261 = _1749_v8;
        let _rhs262 = _module.__default.fm1(_dafny.Seq.IsProperPrefixOf(_1787_v36, _dafny.Seq.update(_1787_v36, _module.__default.safeIndex((_this).f30, new BigNumber((_1787_v36).length)), _1786_v35)), (_1750_v9).f48, globalState);
        let _lhs203 = globalState;
        let _lhs204 = globalState;
        let _lhs205 = _1745_v4;
        _lhs203.f13 = _rhs259;
        _lhs204.f13 = _rhs260;
        _1749_v8 = _rhs261;
        _lhs205.f31 = _rhs262;
        let _1788_v37;
        let _1789_v38;
        let _1790_v39;
        let _1791_v40;
        let _out60;
        let _out61;
        let _out62;
        let _out63;
        let _outcollector17 = (_this).m7(!(_1741_v0.f28), p0, _module.__default.fm1((_1750_v9).f47, (_this).f30, globalState), globalState);
        _out60 = _outcollector17[0];
        _out61 = _outcollector17[1];
        _out62 = _outcollector17[2];
        _out63 = _outcollector17[3];
        _1788_v37 = _out60;
        _1789_v38 = _out61;
        _1790_v39 = _out62;
        _1791_v40 = _out63;
        (_1741_v0).f29 = !(((_1741_v0.f28) ? ((_1750_v9).f47) : ((_1745_v4.f29) || (_1745_v4.f31))));
      }
      let _1792_v41;
      let _nw303 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _1792_v41 = _nw303;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1792_v41).length))) {
        let _1793_i6 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1793_i6)) && ((_1793_i6).isLessThan(new BigNumber((_1792_v41).length))))) {
          (_1792_v41)[(_1793_i6)] = _module.__default.safeDivisionInt(_1793_i6, _module.__default.safeDivisionInt((_1745_v4).f30, new BigNumber((_1749_v8).length)));
        }
      }
      r0 = p0;
      let _1794_v42;
      _1794_v42 = _dafny.Set.fromElements(new BigNumber((_1749_v8).length));
      let _1795_v43;
      _1795_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1745_v4.f29,_1794_v42);
      r1 = (_dafny.ZERO).minus(new BigNumber((((_1794_v42).Intersect(_1794_v42)).Intersect((((_1795_v43).contains(_1741_v0.f29)) ? ((_1795_v43).get(_1741_v0.f29)) : (_dafny.Set.fromElements((_this).f30))))).length));
      return [r0, r1];
    }
    get f41() {
      let _this = this;
      return _this._f41;
    };
    get f42() {
      let _this = this;
      return _this._f42;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f28 = false;
      this._f29 = false;
      this._f31 = false;
      this._f35 = false;
      this._f36 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f34 = _dafny.ZERO;
      this._f32 = _dafny.Seq.of();
      this._f33 = _dafny.Seq.of();
      this._f30 = _dafny.ZERO;
      this.f40 = false;
    }
    _parentTraits() {
      return [_module.T4, _module.T2, _module.T3, _module.T1, _module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    set f31(value) {
      let _this = this;
      _this._f31 = value;
    };
    get f35() {
      let _this = this;
      return _this._f35;
    };
    set f35(value) {
      let _this = this;
      _this._f35 = value;
    };
    get f36() {
      let _this = this;
      return _this._f36;
    };
    set f36(value) {
      let _this = this;
      _this._f36 = value;
    };
    get f34() {
      let _this = this;
      return _this._f34;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    __ctor(f40, f36, f34, f35, f32, f33, f30, f31, f28, f29) {
      let _this = this;
      (_this).f40 = f40;
      (_this)._f36 = f36;
      (_this)._f34 = f34;
      (_this)._f35 = f35;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm7(p0, p1, globalState) {
      let _this = this;
      return _this.f35;
    };
    fm6(globalState) {
      let _this = this;
      return _this.f35;
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Set.fromElements((_this).f34);
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (_this.f40) {
        return (_this).f32;
      } else {
        return _dafny.Seq.Concat((_this).f32, _dafny.Seq.of(_this.f29));
      }
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(((_this.f28) ? (_dafny.Seq.UnicodeFromString("qlucvm")) : (_dafny.Seq.UnicodeFromString("ef"))), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sgolrgbnu"), _dafny.Seq.UnicodeFromString("tbm")));
    };
    fm2(globalState) {
      let _this = this;
      let _source30 = _module.D4.create_DC10(_dafny.MultiSet.fromElements(new BigNumber(747), new BigNumber(611)));
      if (_source30.is_DC11) {
        let _1796___mcc_h0 = (_source30).cf23;
        let _1797___mcc_h1 = (_source30).cf24;
        let _1798___mcc_h2 = (_source30).cf25;
        let _1799_cf25 = _1798___mcc_h2;
        let _1800_cf24 = _1797___mcc_h1;
        let _1801_cf23 = _1796___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus((_this).f30));
      } else {
        let _1802___mcc_h3 = (_source30).cf22;
        let _1803_cf22 = _1802___mcc_h3;
        return (_dafny.Map.Empty.slice().updateUnsafe(_this.f35,(_dafny.ZERO).minus((_this).f30))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f40,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f31)).length)));
      }
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      (_this).f36 = new _dafny.CodePoint('a'.codePointAt(0));
      if ((((_this).f30).isLessThan((_this).f30)) && (_this.f28)) {
        let _1804_v0;
        _1804_v0 = _dafny.Seq.UnicodeFromString("suoxd");
        let _1805_v1;
        _1805_v1 = _dafny.MultiSet.fromElements((_this).f34, new BigNumber((_1804_v0).length));
        let _1806_v2;
        _1806_v2 = _dafny.Seq.of((_this).f30);
        let _1807_v3;
        _1807_v3 = _dafny.Seq.of(_dafny.Seq.of((_this).f34, new BigNumber((_1805_v1).cardinality()), (_this).f30), _1806_v2, _1806_v2, _1806_v2);
        let _1808_v4;
        let _init42 = function (_1809_i0) {
          return (_1809_i0).multipliedBy((_this).f34);
        };
        let _nw304 = Array((new BigNumber(28)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw304.length); _i0_42++) {
          _nw304[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1808_v4 = _nw304;
        let _1810_v5;
        _1810_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1808_v4,_1804_v0);
        let _1811_v6;
        _1811_v6 = _dafny.MultiSet.fromElements(_this.f31);
        let _1812_v7;
        let _nw305 = new _module.C7();
        _nw305.__ctor(_module.__default.safeModuloInt(new BigNumber(-494), (_this).f30), (_module.__default.fm49(_1805_v1, _this.f36, _1807_v3, globalState)).Difference(_1805_v1), new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(_1808_v4,_1804_v0)).update(_1808_v4, _1804_v0)).Merge(_1810_v5)).length), _this.f40, _dafny.Seq.contains(_1804_v0, _module.__default.fm34((_this).f34, (_this).f30, new BigNumber((_1811_v6).cardinality()), !(_this.f35), globalState)), _this.f29);
        _1812_v7 = _nw305;
        _1812_v7 = _1812_v7;
        r1 = (_dafny.ZERO).minus(_module.__default.fm39(true, _1805_v1, _this.f40, globalState));
        let _index239 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1808_v4).length));
        (_1808_v4)[_index239] = new BigNumber((_dafny.Seq.Concat((_this).f32, (_this).f32)).length);
        let _index240 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1808_v4).length));
        (_1808_v4)[_index240] = (_this).f34;
        let _1813_v9;
        _1813_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_1808_v4);
        let _1814_v10;
        _1814_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,true);
        let _1815_v11;
        _1815_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1813_v9).length),_1814_v10);
        let _1816_v13;
        _1816_v13 = _dafny.Set.fromElements(function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of ((_1815_v11).update((_this).f34, _1814_v10)).Keys.Elements) {
            let _1817_v8 = _compr_43;
            if (((_1815_v11).update((_this).f34, _1814_v10)).contains(_1817_v8)) {
              _coll43.push([(_1817_v8).plus(new BigNumber(24)),_this.f29]);
            }
          }
          return _coll43;
        }(), function () {
          let _coll44 = new _dafny.Map();
          for (const _compr_44 of _dafny.IntegerRange(new BigNumber(985), new BigNumber(854))) {
            let _1818_v12 = _compr_44;
            if (((new BigNumber(985)).isLessThanOrEqualTo(_1818_v12)) && ((_1818_v12).isLessThan(new BigNumber(854)))) {
              _coll44.push([_module.__default.safeDivisionInt(_1818_v12, (_1808_v4)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_1808_v4).length))]),false]);
            }
          }
          return _coll44;
        }(), _1814_v10);
        _1816_v13 = _dafny.Set.fromElements(_1814_v10, function () {
          let _coll45 = new _dafny.Map();
          for (const _compr_45 of _dafny.IntegerRange(new BigNumber(-637), new BigNumber(477))) {
            let _1819_v14 = _compr_45;
            if (((new BigNumber(-637)).isLessThanOrEqualTo(_1819_v14)) && ((_1819_v14).isLessThan(new BigNumber(477)))) {
              _coll45.push([(_1819_v14).multipliedBy((_this).f30),_1812_v7.f28]);
            }
          }
          return _coll45;
        }());
        _1811_v6 = _1811_v6;
      } else {
        let _1820_v15;
        _1820_v15 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f35);
        let _1821_v16;
        _1821_v16 = _dafny.Seq.of((_this).f30, (_this).f34);
        let _1822_v17;
        _1822_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1820_v15,_1821_v16);
        let _1823_v18;
        let _out64;
        _out64 = _module.__default.m0((((_1822_v17).contains(_1820_v15)) ? ((_1822_v17).get(_1820_v15)) : (_1821_v16)), _this.f28, ((_this).f30).minus(new BigNumber(-749)), _this.f29, globalState);
        _1823_v18 = _out64;
        (globalState).f13 = (_this).f34;
        let _1824_v19;
        let _init43 = function (_1825_i1) {
          return (_1825_i1).multipliedBy(new BigNumber((_dafny.Seq.of(_this.f28, _this.f29)).length));
        };
        let _nw306 = Array((new BigNumber(6)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw306.length); _i0_43++) {
          _nw306[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1824_v19 = _nw306;
        let _index241 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1824_v19).length));
        (_1824_v19)[_index241] = (_this).f30;
        let _index242 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1824_v19).length));
        (_1824_v19)[_index242] = (_this).f34;
        (globalState).f1 = _this.f28;
        _1820_v15 = (_1820_v15).update(_this.f40, !(((_this.f28) ? (_this.f40) : (_this.f31))));
      }
      let _1826_i2;
      _1826_i2 = _dafny.ZERO;
      L5: {
        while (((_this.f29) ? (!((_this.f40) === (true))) : (true))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1826_i2)) {
              break L5;
            }
            _1826_i2 = (_1826_i2).plus(_dafny.ONE);
            let _1827_v20;
            _1827_v20 = _dafny.Seq.UnicodeFromString("idpxt");
            let _1828_v21;
            _1828_v21 = _dafny.Seq.of(_this.f35, _this.f35, (new BigNumber((_1827_v20).length)).isLessThanOrEqualTo((_this).f30));
            _1828_v21 = _1828_v21;
            let _1829_v22;
            let _init44 = function (_1830_i3) {
              return new _dafny.CodePoint('u'.codePointAt(0));
            };
            let _nw307 = Array((new BigNumber(13)).toNumber());
            for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw307.length); _i0_44++) {
              _nw307[_i0_44] = _init44(new BigNumber(_i0_44));
            }
            _1829_v22 = _nw307;
            let _1831_v23;
            let _nw308 = new _module.C0();
            _nw308.__ctor(_1829_v22);
            _1831_v23 = _nw308;
            let _source31 = _module.D9.create_DC27(_1831_v23, _1827_v20);
            if (_source31.is_DC27) {
              let _1832___mcc_h0 = (_source31).cf53;
              let _1833___mcc_h1 = (_source31).cf54;
              let _1834_cf54 = _1833___mcc_h1;
              let _1835_cf53 = _1832___mcc_h0;
              let _1836_v24;
              _1836_v24 = _module.D9.create_DC27(_1835_cf53, ((((_this).f32)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length))]) ? (_1827_v20) : (_1834_cf54)));
              let _rhs263 = _1836_v24;
              let _rhs264 = (_this).fm3(_this.f31, _this.f40, globalState);
              let _rhs265 = _this.f29;
              _1836_v24 = _rhs263;
              r2 = _rhs264;
              r0 = _rhs265;
              (globalState).f12 = _this.f35;
              (globalState).f13 = ((_this.f28) ? ((_this).f34) : ((_dafny.ZERO).minus((_this).f30)));
              let _1837_v25;
              _1837_v25 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), function (_1838_i4) {
                return (_this).f30;
              }));
              let _1839_v26;
              _1839_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_1837_v25);
              let _1840_v27;
              _1840_v27 = _dafny.Seq.of((_this).f34);
              let _1841_v29;
              _1841_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1834_cf54).length),_this.f36);
              let _1842_v31;
              _1842_v31 = _module.D16.create_DC39(_1837_v25);
              let _1843_v32;
              _1843_v32 = _dafny.Seq.of(_dafny.Seq.of((_this).f30, (_this).f34), _1840_v27, _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("emyatjs")).length)), (_this).f34));
              let _1844_v34;
              let _nw309 = Array((new BigNumber(23)).toNumber());
              _nw309[(_dafny.ZERO).toNumber()] = _1837_v25;
              _nw309[(_dafny.ONE).toNumber()] = (_1837_v25).Intersect(_1837_v25);
              _nw309[(new BigNumber(2)).toNumber()] = (((_1839_v26).contains((_this).f34)) ? ((_1839_v26).get((_this).f34)) : (_1837_v25));
              _nw309[(new BigNumber(3)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(4)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(5)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(6)).toNumber()] = ((_this.f40) ? (_dafny.Set.fromElements(_1840_v27, _1840_v27)) : (_1837_v25));
              _nw309[(new BigNumber(7)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(8)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(9)).toNumber()] = (_dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.of((_this).f34), _module.__default.safeIndex(new BigNumber(-949), new BigNumber((_dafny.Seq.of((_this).f34)).length)), (_this).f34), _1840_v27, _1840_v27)).Intersect(_1837_v25);
              _nw309[(new BigNumber(10)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(11)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(12)).toNumber()] = function () {
                let _coll46 = new _dafny.Set();
                for (const _compr_46 of (_1837_v25).Elements) {
                  let _1845_v28 = _compr_46;
                  if ((_1837_v25).contains(_1845_v28)) {
                    _coll46.add(_1845_v28);
                  }
                }
                return _coll46;
              }();
              _nw309[(new BigNumber(13)).toNumber()] = (_dafny.Set.fromElements(_dafny.Seq.of((_this).f34, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f35,false)).length), new BigNumber((function () {
                let _coll47 = new _dafny.Set();
                for (const _compr_47 of (_1841_v29).Keys.Elements) {
                  let _1846_v30 = _compr_47;
                  if ((_1841_v29).contains(_1846_v30)) {
                    _coll47.add((_1846_v30).multipliedBy(new BigNumber(-831)));
                  }
                }
                return _coll47;
              }()).length), new BigNumber((_dafny.Seq.UnicodeFromString("wt")).length)))).Difference((_1842_v31).dtor_cf72);
              _nw309[(new BigNumber(14)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(15)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(16)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(17)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(18)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(19)).toNumber()] = (_module.__default.fm54(new BigNumber(536), _1840_v27, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), function (_1847_i5) {
                return new _dafny.CodePoint('w'.codePointAt(0));
              })).length), globalState)).Difference(_1837_v25);
              _nw309[(new BigNumber(20)).toNumber()] = _1837_v25;
              _nw309[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements(_1840_v27);
              _nw309[(new BigNumber(22)).toNumber()] = function () {
                let _coll48 = new _dafny.Set();
                for (const _compr_48 of (_1843_v32).Elements) {
                  let _1848_v33 = _compr_48;
                  if (_dafny.Seq.contains(_1843_v32, _1848_v33)) {
                    _coll48.add(_1848_v33);
                  }
                }
                return _coll48;
              }();
              _1844_v34 = _nw309;
              let _index243 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_1844_v34).length));
              (_1844_v34)[_index243] = _1837_v25;
              let _index244 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_1844_v34).length));
              let _rhs266 = _1837_v25;
              let _rhs267 = ((_this).f30).minus((_this).f30);
              let _rhs268 = _module.__default.safeModuloInt(((_this).f30).minus(new BigNumber((function () {
                let _coll49 = new _dafny.Map();
                for (const _compr_49 of _dafny.IntegerRange(new BigNumber(642), new BigNumber(927))) {
                  let _1849_v35 = _compr_49;
                  if (((new BigNumber(642)).isLessThanOrEqualTo(_1849_v35)) && ((_1849_v35).isLessThan(new BigNumber(927)))) {
                    _coll49.push([(_1849_v35).multipliedBy((_this).f34),(_this).f30]);
                  }
                }
                return _coll49;
              }()).length)), (_module.__default.fm32(_this.f40, globalState)).minus(_module.__default.fm32(_this.f35, globalState)));
              let _rhs269 = (_this).f34;
              let _lhs206 = _1844_v34;
              let _lhs207 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_1844_v34).length));
              let _lhs208 = globalState;
              let _lhs209 = globalState;
              _lhs206[_lhs207] = _rhs266;
              _lhs208.f13 = _rhs267;
              r1 = _rhs268;
              _lhs209.f13 = _rhs269;
            } else {
              let _1850___mcc_h2 = (_source31).cf52;
              let _1851_cf52 = _1850___mcc_h2;
              let _1852_v36;
              let _nw310 = Array((new BigNumber(22)).toNumber()).fill(false);
              _1852_v36 = _nw310;
              (globalState).f21 = ((_this.f40) ? (_1852_v36) : (_1852_v36));
              let _1853_v37;
              let _init45 = ((_1854_v20) => function (_1855_i6) {
                return (_1855_i6).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_1854_v20).length), (_this).f30)).length));
              })(_1827_v20);
              let _nw311 = Array((new BigNumber(7)).toNumber());
              for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw311.length); _i0_45++) {
                _nw311[_i0_45] = _init45(new BigNumber(_i0_45));
              }
              _1853_v37 = _nw311;
              let _index245 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1853_v37).length));
              (_1853_v37)[_index245] = ((_this).f34).minus((_this).f30);
              let _index246 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1853_v37).length));
              let _rhs270 = (_module.__default.safeModuloInt((_1851_cf52).f46, (_1851_cf52).f46)).isLessThan((_this).f30);
              let _rhs271 = (_this).f30;
              let _lhs210 = globalState;
              let _lhs211 = _1853_v37;
              let _lhs212 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1853_v37).length));
              _lhs210.f1 = _rhs270;
              _lhs211[_lhs212] = _rhs271;
              let _1856_v38;
              _1856_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_this.f40);
              let _1857_v39;
              let _nw312 = new _module.C2();
              _nw312.__ctor(_module.__default.safeDivisionInt(new BigNumber((_1856_v38).length), (_1853_v37)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_1853_v37).length))]), _module.__default.fm24(_this.f31, new BigNumber((_1827_v20).length), globalState), _this.f29, _this.f35, true);
              _1857_v39 = _nw312;
              (globalState).f13 = ((_1857_v39).f46).minus(new BigNumber(514));
            }
            (globalState).f12 = _this.f40;
            let _1858_v40;
            let _nw313 = Array((new BigNumber(13)).toNumber()).fill([]);
            _1858_v40 = _nw313;
            let _1859_v41;
            let _nw314 = Array((new BigNumber(4)).toNumber()).fill(false);
            _1859_v41 = _nw314;
            let _nw315 = Array((new BigNumber(24)).toNumber());
            _nw315[(_dafny.ZERO).toNumber()] = ((_this.f31) ? (_1859_v41) : (_1859_v41));
            _nw315[(_dafny.ONE).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(2)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(3)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(4)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(5)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(6)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(7)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(8)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(9)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(10)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(11)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(12)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(13)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(14)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(15)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(16)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(17)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(18)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(19)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(20)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(21)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(22)).toNumber()] = _1859_v41;
            _nw315[(new BigNumber(23)).toNumber()] = _1859_v41;
            _1858_v40 = _nw315;
          }
        }
      }
      (_this).f40 = _this.f28;
      (_this).f35 = !((_this).f30).isEqualTo((_this).f30);
      let _1860_v42;
      _1860_v42 = _dafny.MultiSet.fromElements(_module.__default.fm1(_this.f29, (_this).f30, globalState));
      let _1861_v43;
      _1861_v43 = _module.D0.create_DC1(_1860_v42);
      let _1862_v44;
      _1862_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1861_v43,_this.f40);
      let _1863_v45;
      _1863_v45 = _module.D8.create_DC21(_1862_v44);
      let _1864_v46;
      _1864_v46 = _dafny.Seq.of(_1863_v45);
      let _1865_v47;
      _1865_v47 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), function (_1866_i7) {
        return _module.D8.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(_this.f31)),_this.f29));
      }), _1864_v46, _dafny.Seq.Create(_module.__default.abs(new BigNumber(47)), ((_1867_v45) => function (_1868_i8) {
        return _1867_v45;
      })(_1863_v45)), _dafny.Seq.of(_1863_v45, _1863_v45, _1863_v45, _1863_v45));
      if (_dafny.Seq.IsProperPrefixOf(_1865_v47, _dafny.Seq.of(_dafny.Seq.of(_1863_v45)))) {
        let _1869_v48;
        let _nw316 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
        _1869_v48 = _nw316;
        let _1870_v49;
        let _nw317 = new _module.C3();
        _nw317.__ctor((_this).f34, _1869_v48, new _dafny.CodePoint('i'.codePointAt(0)), new BigNumber(101), _this.f35, _dafny.Seq.Concat(_dafny.Seq.of(true), (_this).f32), (_this).f33, new BigNumber((_dafny.Seq.UnicodeFromString("l")).length), _this.f28, _this.f29, (false) && (false));
        _1870_v49 = _nw317;
        let _1871_v50;
        _1871_v50 = _dafny.Set.fromElements(_this.f35);
        let _1872_v51;
        let _init46 = function (_1873_i9) {
          return _this.f36;
        };
        let _nw318 = Array((new BigNumber(18)).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw318.length); _i0_46++) {
          _nw318[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _1872_v51 = _nw318;
        let _1874_v52;
        let _nw319 = new _module.C0();
        _nw319.__ctor(_1872_v51);
        _1874_v52 = _nw319;
        let _1875_v53;
        let _1876_v54;
        let _1877_v55;
        let _1878_v56;
        let _out65;
        let _out66;
        let _out67;
        let _out68;
        let _outcollector18 = (_this).m6(_module.__default.safeDivisionInt((_this).f34, (_1870_v49).f44), _module.__default.fm39(true, _dafny.MultiSet.fromElements(new BigNumber((_1871_v50).length)), _this.f31, globalState), _this.f36, _1874_v52, globalState);
        _out65 = _outcollector18[0];
        _out66 = _outcollector18[1];
        _out67 = _outcollector18[2];
        _out68 = _outcollector18[3];
        _1875_v53 = _out65;
        _1876_v54 = _out66;
        _1877_v55 = _out67;
        _1878_v56 = _out68;
        let _1879_v57;
        _1879_v57 = _dafny.Seq.UnicodeFromString("fu");
        r2 = _dafny.Seq.Concat(_1879_v57, _1879_v57);
        (_this).f36 = _this.f36;
        let _1880_v58;
        _1880_v58 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,new BigNumber((_1860_v42).cardinality()));
        _1880_v58 = (_1880_v58).update((_this).fm6(globalState), (_this).f34);
      } else {
        r1 = ((_this).f30).multipliedBy((_this).f30);
        r1 = (new BigNumber(((_this).fm3(_this.f31, _this.f35, globalState)).length)).minus((_this).f30);
        (globalState).f1 = _this.f35;
        let _1881_v59;
        let _nw320 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
        _1881_v59 = _nw320;
        let _1882_v60;
        let _nw321 = new _module.C4();
        _nw321.__ctor(new BigNumber(657), false, (_this).f32, (_this).f33, (_this).f34, false, _this.f31, _this.f28);
        _1882_v60 = _nw321;
        let _1883_v61;
        _1883_v61 = _dafny.Seq.of(_1882_v60, _1882_v60, _1882_v60, _1882_v60, _1882_v60);
        let _index247 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1881_v59).length));
        (_1881_v59)[_index247] = _dafny.Seq.Concat(_1883_v61, _1883_v61);
        let _1884_v62;
        _1884_v62 = _module.D17.create_DC42(_1883_v61);
        let _index248 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1881_v59).length));
        let _rhs272 = (_1884_v62).dtor_cf76;
        let _rhs273 = ((_this).f34).isLessThanOrEqualTo((_this).f30);
        let _lhs213 = _1881_v59;
        let _lhs214 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1881_v59).length));
        let _lhs215 = globalState;
        _lhs213[_lhs214] = _rhs272;
        _lhs215.f1 = _rhs273;
        let _1885_v63;
        _1885_v63 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_this).f30);
        _1885_v63 = (_1885_v63).update(_this.f28, (_this).f34);
      }
      r0 = _this.f35;
      let _1886_v64;
      _1886_v64 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,_this.f36);
      r1 = (new BigNumber((_1886_v64).length)).multipliedBy((_this).f34);
      let _1887_v65;
      _1887_v65 = _dafny.Seq.UnicodeFromString("v");
      r2 = _1887_v65;
      return [r0, r1, r2];
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _1888_v0;
      _1888_v0 = _dafny.Set.fromElements(_this.f35, _this.f29, _this.f35);
      let _hi10 = _module.__default.safeDivisionInt(new BigNumber((_1888_v0).length), (_this).f30);
      for (let _1889_i0 = ((_this.f31) ? (p1) : ((_this).f30)); _1889_i0.isLessThan(_hi10); _1889_i0 = _1889_i0.plus(_dafny.ONE)) {
        let _1890_v1;
        _1890_v1 = _dafny.Seq.UnicodeFromString("vq");
        let _1891_v2;
        _1891_v2 = _dafny.Seq.of(p1, _1889_i0, new BigNumber(43), _1889_i0, p0);
        let _1892_v4;
        _1892_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f34)).length),p1);
        let _1893_v5;
        let _init47 = function (_1894_i1) {
          return (_1894_i1).plus((_dafny.ZERO).minus((_this).f34));
        };
        let _nw322 = Array((new BigNumber(7)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw322.length); _i0_47++) {
          _nw322[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1893_v5 = _nw322;
        let _1895_v6;
        _1895_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1893_v5,_this.f29);
        let _1896_v7;
        let _nw323 = new _module.C5();
        _nw323.__ctor(_this.f35, new BigNumber((function () {
          let _coll50 = new _dafny.Map();
          for (const _compr_50 of (_1892_v4).Keys.Elements) {
            let _1897_v3 = _compr_50;
            if ((_1892_v4).contains(_1897_v3)) {
              _coll50.push([(_1897_v3).minus((_this).f30),_this.f28]);
            }
          }
          return _coll50;
        }()).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), _this.f28, _module.__default.fm0((_this).f34, false, _this.f40, new BigNumber((_1890_v1).length), globalState), _dafny.Seq.update((_this).f33, _module.__default.safeIndex((_this).f30, new BigNumber(((_this).f33).length)), (_this).f32), new BigNumber((_1895_v6).length), false, _this.f31, true);
        _1896_v7 = _nw323;
        let _1898_v8;
        _1898_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1891_v2,_1896_v7);
        (globalState).f13 = new BigNumber((_dafny.Seq.update(_1890_v1, _module.__default.safeIndex(new BigNumber((_1898_v8).length), new BigNumber((_1890_v1).length)), _this.f36)).length);
        (_this).f35 = false;
        (_this).f36 = _this.f36;
        let _1899_v9;
        _1899_v9 = _module.D16.create_DC40(_dafny.Seq.UnicodeFromString("rcj"), (new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_1891_v2)).cardinality()), (_this).f34)).length)).plus((_1896_v7).f30));
        let _source32 = _1899_v9;
        if (_source32.is_DC40) {
          let _1900___mcc_h0 = (_source32).cf73;
          let _1901___mcc_h1 = (_source32).cf74;
          let _1902_cf74 = _1901___mcc_h1;
          let _1903_cf73 = _1900___mcc_h0;
          let _1904_v10;
          let _nw324 = Array((new BigNumber(19)).toNumber()).fill(false);
          _1904_v10 = _nw324;
          let _1905_v11;
          _1905_v11 = _dafny.Seq.of(_1904_v10, _1904_v10);
          let _1906_v12;
          _1906_v12 = _1905_v11;
          _1905_v11 = (_1906_v12);
          (globalState).f26 = _this.f35;
          let _1907_v13;
          _1907_v13 = _dafny.Set.fromElements(_module.__default.fm55((_dafny.ZERO).minus(p0), _1896_v7.f29, new BigNumber((_dafny.Set.fromElements(true, _1896_v7.f31)).length), globalState), _1888_v0);
          _1907_v13 = _1907_v13;
          let _1908_v14;
          _1908_v14 = _dafny.Seq.of(_1888_v0);
          let _1909_v15;
          _1909_v15 = _module.D2.create_DC6(new BigNumber((_dafny.Seq.of((_this).f34)).length), _1896_v7.f29, new BigNumber(423), _this.f29, new BigNumber(((_1908_v14)[_module.__default.safeIndex((_1896_v7).f30, new BigNumber((_1908_v14).length))]).length));
          _1909_v15 = _1909_v15;
        } else if (_source32.is_DC39) {
          let _1910___mcc_h2 = (_source32).cf72;
          let _1911_cf72 = _1910___mcc_h2;
          let _1912_v16;
          let _nw325 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
          _1912_v16 = _nw325;
          let _1913_v17;
          let _nw326 = new _module.C4();
          _nw326.__ctor(_1889_i0, _module.__default.fm1(_1896_v7.f31, p0, globalState), _dafny.Seq.of(_1896_v7.f31), (_1896_v7).f33, new BigNumber((_module.__default.fm0(new BigNumber(553), _this.f35, _this.f29, p1, globalState)).length), _1896_v7.f31, false, _1896_v7.f28);
          _1913_v17 = _nw326;
          let _1914_v18;
          let _nw327 = new _module.C3();
          _nw327.__ctor((_this).f34, ((_1896_v7.f31) ? (_1912_v16) : (_1912_v16)), _this.f36, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-686)), ((_1915_v7) => function (_1916_i2) {
            return new BigNumber(((_1915_v7).f32).length);
          })(_1896_v7))).length), _this.f31, _dafny.Seq.Concat((_1896_v7).f32, _dafny.Seq.of(_1896_v7.f31, _this.f31, !(_1896_v7.f29))), (_1896_v7).f33, (_1896_v7).f30, _module.__default.fm1(_this.f35, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p1,_1913_v17)).update((_dafny.ZERO).minus(new BigNumber((_1890_v1).length)), _1913_v17)).length), globalState), _this.f40, _1896_v7.f29);
          _1914_v18 = _nw327;
          let _index249 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1893_v5).length));
          (_1893_v5)[_index249] = (_this).f30;
          let _index250 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1893_v5).length));
          (_1893_v5)[_index250] = _module.__default.safeDivisionInt((_1896_v7).f30, (_1896_v7).f30);
          let _1917_v19;
          let _init48 = ((_1918_v7) => function (_1919_i3) {
            return (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_this.f29, _this.f40))).IsProperSubsetOf((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_1918_v7.f28))));
          })(_1896_v7);
          let _nw328 = Array((new BigNumber(19)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw328.length); _i0_48++) {
            _nw328[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1917_v19 = _nw328;
          let _index251 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_1917_v19).length));
          (_1917_v19)[_index251] = _this.f28;
          let _index252 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_1917_v19).length));
          (_1917_v19)[_index252] = (p1).isLessThan(p0);
          _1917_v19 = _1917_v19;
        } else {
          let _1920___mcc_h3 = (_source32).cf75;
          let _1921_cf75 = _1920___mcc_h3;
          (globalState).f1 = _this.f40;
          let _1922_v20;
          let _nw329 = Array((new BigNumber(16)).toNumber()).fill([]);
          _1922_v20 = _nw329;
          let _1923_v21;
          let _init49 = ((_1924_v7) => function (_1925_i4) {
            return _1924_v7.f28;
          })(_1896_v7);
          let _nw330 = Array((new BigNumber(13)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw330.length); _i0_49++) {
            _nw330[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1923_v21 = _nw330;
          let _index253 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1922_v20).length));
          (_1922_v20)[_index253] = _1923_v21;
          let _index254 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1922_v20).length));
          let _rhs274 = new BigNumber(522);
          let _rhs275 = _1923_v21;
          let _lhs216 = _1922_v20;
          let _lhs217 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1922_v20).length));
          r3 = _rhs274;
          _lhs216[_lhs217] = _rhs275;
          let _1926_v22;
          _1926_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1923_v21,new BigNumber((_1890_v1).length));
          let _1927_v23;
          let _nw331 = new _module.C4();
          _nw331.__ctor((_1896_v7).f30, _1896_v7.f28, (_1896_v7).f32, _dafny.Seq.Concat(_dafny.Seq.update((_1896_v7).f33, _module.__default.safeIndex(p1, new BigNumber(((_1896_v7).f33).length)), (_this).f32), (_this).f33), _module.__default.safeModuloInt((_1896_v7).f30, new BigNumber(143)), ((_1896_v7).f30).isEqualTo(new BigNumber((_1926_v22).length)), _this.f31, _1896_v7.f28);
          _1927_v23 = _nw331;
          let _1928_v24;
          _1928_v24 = _dafny.Set.fromElements(_this.f36, _module.__default.fm21(p2, p0, globalState));
          let _1929_v25;
          _1929_v25 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f30);
          let _1930_v26;
          _1930_v26 = _module.D2.create_DC6(new BigNumber((_1928_v24).length), _1896_v7.f29, new BigNumber(230), _1896_v7.f31, (((_1929_v25).contains(_this.f28)) ? ((_1929_v25).get(_this.f28)) : (new BigNumber((_1891_v2).length))));
          let _1931_v27;
          let _nw332 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1931_v27 = _nw332;
          let _1932_v28;
          _1932_v28 = _dafny.Seq.of(_1890_v1);
          let _rhs276 = _module.__default.fm56((_this).f30, _dafny.Seq.IsPrefixOf(_1890_v1, _1890_v1), ((_dafny.ZERO).minus((_1896_v7).f30)).multipliedBy(p0), globalState);
          let _rhs277 = _module.__default.safeDivisionInt(p1, (_this).f30);
          let _rhs278 = p1;
          let _rhs279 = _1931_v27;
          let _rhs280 = !_dafny.areEqual((_1932_v28)[_module.__default.safeIndex(new BigNumber((_1890_v1).length), new BigNumber((_1932_v28).length))], _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(443)), ((_1933_p2) => function (_1934_i5) {
            return _1933_p2;
          })(p2)), _1890_v1));
          let _lhs218 = globalState;
          let _lhs219 = globalState;
          let _lhs220 = _1896_v7;
          _1930_v26 = _rhs276;
          _lhs218.f13 = _rhs277;
          _lhs219.f13 = _rhs278;
          _1931_v27 = _rhs279;
          _lhs220.f28 = _rhs280;
        }
      }
      let _1935_v29;
      let _nw333 = Array((new BigNumber(14)).toNumber()).fill(false);
      _1935_v29 = _nw333;
      let _1936_v30;
      _1936_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_this.f29);
      let _1937_v31;
      _1937_v31 = _dafny.Map.Empty.slice().updateUnsafe(_this.f40,new BigNumber((_1936_v30).length));
      let _1938_v32;
      _1938_v32 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f40),_module.__default.fm57(p2, new BigNumber((_1937_v31).length), globalState));
      let _index255 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1935_v29).length));
      (_1935_v29)[_index255] = _module.__default.fm1(_this.f29, _module.__default.fm24(false, new BigNumber((_1938_v32).length), globalState), globalState);
      let _index256 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1935_v29).length));
      (_1935_v29)[_index256] = ((_this).f32)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f34)), new BigNumber(((_this).f32).length))];
      let _1939_i6;
      _1939_i6 = _dafny.ZERO;
      L6: {
        while (!(_this.f28)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1939_i6)) {
              break L6;
            }
            _1939_i6 = (_1939_i6).plus(_dafny.ONE);
            let _1940_v33;
            _1940_v33 = _dafny.MultiSet.fromElements(_this.f40);
            let _1941_v34;
            _1941_v34 = _dafny.MultiSet.fromElements((_this).f34, p1);
            let _1942_v35;
            _1942_v35 = _dafny.Seq.of((_this).f34);
            let _1943_v36;
            _1943_v36 = _module.D4.create_DC11(p1, _this.f40, (((_1940_v33).contains(_this.f28)) ? ((_1940_v33).get(_this.f28)) : (p0)));
            let _1944_v37;
            _1944_v37 = _dafny.MultiSet.fromElements(_1943_v36);
            let _1945_v38;
            _1945_v38 = _dafny.Seq.UnicodeFromString("ff");
            let _1946_v39;
            let _nw334 = new _module.C7();
            _nw334.__ctor((((_1940_v33).contains((_1935_v29)[_module.__default.safeIndex(new BigNumber(82), new BigNumber((_1935_v29).length))])) ? ((_1940_v33).get((_1935_v29)[_module.__default.safeIndex(new BigNumber(82), new BigNumber((_1935_v29).length))])) : (p0)), (_1941_v34).Difference(_dafny.MultiSet.fromElements(new BigNumber((_1942_v35).length))), p0, (_1944_v37).IsProperSubsetOf(_1944_v37), ((_this).f32)[_module.__default.safeIndex(new BigNumber((_1945_v38).length), new BigNumber(((_this).f32).length))], _this.f29);
            _1946_v39 = _nw334;
            let _1947_v40;
            let _nw335 = Array((new BigNumber(9)).toNumber());
            _1947_v40 = _nw335;
            let _1948_v41;
            _1948_v41 = _dafny.Seq.of(_module.__default.fm49(_1941_v34, _this.f36, _dafny.Seq.of(_1942_v35, _dafny.Seq.of(new BigNumber((_1945_v38).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1947_v40,_this.f31)).length))), globalState), _1941_v34);
            let _rhs281 = _1946_v39;
            let _rhs282 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_1948_v41, _dafny.Seq.Create(_module.__default.abs(new BigNumber(963)), ((_1949_v35) => function (_1950_i7) {
              return _dafny.MultiSet.FromArray(_1949_v35);
            })(_1942_v35))), _dafny.Seq.Concat(_dafny.Seq.update(_1948_v41, _module.__default.safeIndex((_this).f34, new BigNumber((_1948_v41).length)), _1941_v34), _dafny.Seq.update(_1948_v41, _module.__default.safeIndex(new BigNumber(492), new BigNumber((_1948_v41).length)), _dafny.MultiSet.fromElements((_this).f34))));
            _1946_v39 = _rhs281;
            r1 = _rhs282;
            let _index257 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1935_v29).length));
            let _rhs283 = (_this).fm7(_1941_v34, (_dafny.ZERO).minus((_this).f30), globalState);
            let _rhs284 = ((_this).f30).multipliedBy((_dafny.ZERO).minus((_this).f34));
            let _lhs221 = _1935_v29;
            let _lhs222 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1935_v29).length));
            let _lhs223 = globalState;
            _lhs221[_lhs222] = _rhs283;
            _lhs223.f13 = _rhs284;
            (globalState).f21 = _1935_v29;
            let _1951_v42;
            let _nw336 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
            _1951_v42 = _nw336;
            let _index258 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1951_v42).length));
            (_1951_v42)[_index258] = (_this).f32;
            let _index259 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1951_v42).length));
            (_1951_v42)[_index259] = (_this).f32;
          }
        }
      }
      let _1952_v43;
      _1952_v43 = _dafny.Seq.UnicodeFromString("opebgfkic");
      let _1953_v44;
      let _nw337 = Array((new BigNumber(13)).toNumber());
      _nw337[(_dafny.ZERO).toNumber()] = p1;
      _nw337[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f30));
      _nw337[(new BigNumber(2)).toNumber()] = p0;
      _nw337[(new BigNumber(3)).toNumber()] = new BigNumber((_1952_v43).length);
      _nw337[(new BigNumber(4)).toNumber()] = p0;
      _nw337[(new BigNumber(5)).toNumber()] = (_this).f34;
      _nw337[(new BigNumber(6)).toNumber()] = (_this).f30;
      _nw337[(new BigNumber(7)).toNumber()] = (_this).f34;
      _nw337[(new BigNumber(8)).toNumber()] = (_this).f30;
      _nw337[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1952_v43).length));
      _nw337[(new BigNumber(10)).toNumber()] = p1;
      _nw337[(new BigNumber(11)).toNumber()] = (_this).f30;
      _nw337[(new BigNumber(12)).toNumber()] = p0;
      _1953_v44 = _nw337;
      let _1954_v45;
      _1954_v45 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_this.f28)).length), (_this).f30);
      let _1955_v46;
      _1955_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1953_v44,_dafny.Seq.Concat(_module.__default.fm23((_1935_v29)[_module.__default.safeIndex(new BigNumber(82), new BigNumber((_1935_v29).length))], p0, globalState), _1954_v45));
      let _1956_v47;
      _1956_v47 = _dafny.Set.fromElements(new BigNumber((_1954_v45).length));
      let _1957_v48;
      _1957_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1956_v47,_1955_v46);
      let _1958_v49;
      _1958_v49 = _1888_v0;
      let _1959_v50;
      _1959_v50 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f28),_1955_v46);
      let _1960_v51;
      _1960_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1958_v49)).length),(((_1959_v50).contains(_this.f40)) ? ((_1959_v50).get(_this.f40)) : (_1955_v46)));
      _1955_v46 = (((_1957_v48).contains(_1956_v47)) ? ((_1957_v48).get(_1956_v47)) : ((((_1960_v51).contains(p0)) ? ((_1960_v51).get(p0)) : (_1955_v46))));
      let _index260 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1953_v44).length));
      (_1953_v44)[_index260] = (_this).f30;
      let _index261 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1953_v44).length));
      (_1953_v44)[_index261] = (_dafny.ZERO).minus((((_this.f29) ? ((_this).f30) : (p0))).multipliedBy(new BigNumber(-661)));
      (_this).f29 = _this.f29;
      let _1961_v52;
      _1961_v52 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.of(_1937_v31, _1937_v31));
      let _1962_v53;
      _1962_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(26),_this.f31);
      let _1963_v54;
      _1963_v54 = _dafny.Seq.of(_1937_v31);
      let _1964_v55;
      _1964_v55 = _module.D3.create_DC7((((_1961_v52).contains((_dafny.ZERO).minus(new BigNumber((_1962_v53).length)))) ? ((_1961_v52).get((_dafny.ZERO).minus(new BigNumber((_1962_v53).length)))) : (_1963_v54)));
      let _1965_v56;
      _1965_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1964_v55,(_this).f30);
      let _1966_v57;
      _1966_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1965_v56,_1953_v44);
      r0 = _1966_v57;
      r1 = (_this).fm7(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(868)), ((_1967_v45) => function (_1968_i8) {
        return (_module.D2.create_DC6(new BigNumber(438), _this.f35, new BigNumber(871), _this.f28, new BigNumber((_1967_v45).length))).dtor_cf11;
      })(_1954_v45))), _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements((_1935_v29)[_module.__default.safeIndex(new BigNumber(82), new BigNumber((_1935_v29).length))], _this.f40, _this.f35)).cardinality()), new BigNumber((_1952_v43).length)), globalState);
      r2 = _this.f28;
      r3 = p1;
      return [r0, r1, r2, r3];
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f28 = false;
      this._f29 = false;
      this._f31 = false;
      this._f35 = false;
      this._f36 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f34 = _dafny.ZERO;
      this._f32 = _dafny.Seq.of();
      this._f33 = _dafny.Seq.of();
      this._f30 = _dafny.ZERO;
      this._f38 = _dafny.ZERO;
      this._f39 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T4, _module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    set f31(value) {
      let _this = this;
      _this._f31 = value;
    };
    get f35() {
      let _this = this;
      return _this._f35;
    };
    set f35(value) {
      let _this = this;
      _this._f35 = value;
    };
    get f36() {
      let _this = this;
      return _this._f36;
    };
    set f36(value) {
      let _this = this;
      _this._f36 = value;
    };
    get f34() {
      let _this = this;
      return _this._f34;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    __ctor(f38, f39, f36, f34, f35, f32, f33, f30, f31, f28, f29) {
      let _this = this;
      (_this)._f38 = f38;
      (_this)._f39 = f39;
      (_this)._f36 = f36;
      (_this)._f34 = f34;
      (_this)._f35 = f35;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm7(p0, p1, globalState) {
      let _this = this;
      if (true) {
        return false;
      } else {
        return (_module.D0.create_DC0(_this.f35)).dtor_cf0;
      }
    };
    fm6(globalState) {
      let _this = this;
      return !(((_this.f29) ? (_dafny.Map.Empty.slice().updateUnsafe(!(false),_this.f35)) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_module.D0.create_DC0(_this.f35)).dtor_cf0)))).equals((_dafny.Map.Empty.slice().updateUnsafe(!(((_this).f32)[_module.__default.safeIndex((_this).f30, new BigNumber(((_this).f32).length))]),_this.f31)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f31,_this.f28)));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Set.fromElements(new BigNumber((((_module.D4.create_DC10(_dafny.MultiSet.fromElements((_this).f34))).dtor_cf22).Difference(_dafny.MultiSet.fromElements(new BigNumber(571), new BigNumber(((_module.D0.create_DC1(_dafny.MultiSet.fromElements(_this.f28))).dtor_cf1).cardinality()), (_this).f34, (_this).f30, new BigNumber((function () {
        let _coll51 = new _dafny.Map();
        for (const _compr_51 of (_dafny.Set.fromElements((_this).f38)).Elements) {
          let _1969_v0 = _compr_51;
          if ((_dafny.Set.fromElements((_this).f38)).contains(_1969_v0)) {
            _coll51.push([_module.__default.safeDivisionInt(_1969_v0, (_this).f38),_this.f31]);
          }
        }
        return _coll51;
      }()).length)))).cardinality()));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f32;
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nxjjkbq"), _dafny.Seq.UnicodeFromString("buh"));
    };
    fm2(globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_this).f38)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f30)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_this.f35),(_this).f34)));
    };
    fm16(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements((_dafny.ZERO).minus((_this).f30))).Intersect(_dafny.Set.fromElements((_this).f38, (_this).f30))).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(715), (_this).f30));
    };
    fm17(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_this.f35),_dafny.MultiSet.fromElements(_this.f28, !(_this.f31), _this.f31))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true, _this.f31),_dafny.MultiSet.fromElements(_this.f28)))).length)).plus((_this).f34);
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _1970_v0;
      _1970_v0 = _dafny.Set.fromElements(_this.f35, _this.f35);
      if ((_1970_v0).IsDisjointFrom(_dafny.Set.fromElements(_this.f29, _this.f35))) {
        let _1971_v1;
        _1971_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,(_this).f30);
        if (((((_1971_v1).contains(_this.f31)) ? ((_1971_v1).get(_this.f31)) : ((_this).f34))).isLessThanOrEqualTo((_this).f30)) {
          let _1972_v2;
          let _init50 = function (_1973_i0) {
            return _this.f36;
          };
          let _nw338 = Array((_dafny.ONE).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw338.length); _i0_50++) {
            _nw338[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _1972_v2 = _nw338;
          let _1974_v3;
          let _nw339 = new _module.C0();
          _nw339.__ctor(_1972_v2);
          _1974_v3 = _nw339;
          let _1975_v4;
          _1975_v4 = _module.D1.create_DC3(_this.f35, new BigNumber(777), _1974_v3, (_dafny.ZERO).minus((_this).f34));
          _1975_v4 = _1975_v4;
          let _1976_v5;
          _1976_v5 = _dafny.Seq.UnicodeFromString("iyotbppo");
          let _1977_v6;
          _1977_v6 = _module.D5.create_DC12(_1976_v5);
          let _1978_v7;
          _1978_v7 = _module.D0.create_DC0(_this.f35);
          let _1979_v8;
          _1979_v8 = _module.D2.create_DC6((_this).f30, _this.f35, new BigNumber((_1976_v5).length), !((_1978_v7).dtor_cf0), new BigNumber(-57));
          let _1980_v9;
          _1980_v9 = _dafny.MultiSet.fromElements(_this.f28, _this.f35, _this.f28);
          let _1981_v10;
          let _nw340 = Array((new BigNumber(10)).toNumber());
          _nw340[(_dafny.ZERO).toNumber()] = (_this).f30;
          _nw340[(_dafny.ONE).toNumber()] = (_this).f30;
          _nw340[(new BigNumber(2)).toNumber()] = ((_this).f30).multipliedBy((_this).f34);
          _nw340[(new BigNumber(3)).toNumber()] = (new BigNumber(((_1977_v6).dtor_cf26).length)).plus((_this).f38);
          _nw340[(new BigNumber(4)).toNumber()] = (_this).f34;
          _nw340[(new BigNumber(5)).toNumber()] = (_1979_v8).dtor_cf11;
          _nw340[(new BigNumber(6)).toNumber()] = (_this).f34;
          _nw340[(new BigNumber(7)).toNumber()] = new BigNumber((_1980_v9).cardinality());
          _nw340[(new BigNumber(8)).toNumber()] = (_this).f34;
          _nw340[(new BigNumber(9)).toNumber()] = (_this).f38;
          _1981_v10 = _nw340;
          let _1982_v11;
          _1982_v11 = _dafny.Set.fromElements((_this).f34, (_this).f34);
          let _1983_v12;
          _1983_v12 = _dafny.MultiSet.fromElements((_this).f30);
          let _1984_v13;
          let _nw341 = Array((new BigNumber(16)).toNumber());
          _nw341[(_dafny.ZERO).toNumber()] = true;
          _nw341[(_dafny.ONE).toNumber()] = _this.f31;
          _nw341[(new BigNumber(2)).toNumber()] = _this.f28;
          _nw341[(new BigNumber(3)).toNumber()] = (_this).fm16(_this.f31, _1978_v7, _1982_v11, (_this).f38, globalState);
          _nw341[(new BigNumber(4)).toNumber()] = _this.f35;
          _nw341[(new BigNumber(5)).toNumber()] = _this.f31;
          _nw341[(new BigNumber(6)).toNumber()] = _this.f29;
          _nw341[(new BigNumber(7)).toNumber()] = _this.f35;
          _nw341[(new BigNumber(8)).toNumber()] = _this.f31;
          _nw341[(new BigNumber(9)).toNumber()] = !(true);
          _nw341[(new BigNumber(10)).toNumber()] = _this.f28;
          _nw341[(new BigNumber(11)).toNumber()] = _this.f28;
          _nw341[(new BigNumber(12)).toNumber()] = (_this).fm7((_1983_v12).update((_this).f34, _module.__default.abs((_this).f38)), (_this).f30, globalState);
          _nw341[(new BigNumber(13)).toNumber()] = _this.f35;
          _nw341[(new BigNumber(14)).toNumber()] = _this.f29;
          _nw341[(new BigNumber(15)).toNumber()] = _this.f28;
          _1984_v13 = _nw341;
          let _1985_v14;
          _1985_v14 = _module.D3.create_DC9(_this.f28, false, _1984_v13, _this.f31);
          let _index262 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1981_v10).length));
          (_1981_v10)[_index262] = (_dafny.ZERO).minus((((_1985_v14).dtor_cf21) ? ((_this).f30) : ((_this).f30)));
          let _index263 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1981_v10).length));
          (_1981_v10)[_index263] = (_this).f38;
          let _1986_v15;
          _1986_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,_this.f35);
          _1986_v15 = (_1986_v15).update(_this.f36, _this.f35);
          let _1987_v16;
          _1987_v16 = _dafny.Seq.of(((_this).f34).minus((_1981_v10)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_1981_v10).length))]));
          (globalState).f13 = (_1987_v16)[_module.__default.safeIndex((_this).f34, new BigNumber((_1987_v16).length))];
          let _1988_v18;
          _1988_v18 = _dafny.Seq.of(function () {
            let _coll52 = new _dafny.Set();
            for (const _compr_52 of _dafny.IntegerRange(new BigNumber(144), new BigNumber(351))) {
              let _1989_v17 = _compr_52;
              if (((new BigNumber(144)).isLessThanOrEqualTo(_1989_v17)) && ((_1989_v17).isLessThan(new BigNumber(351)))) {
                _coll52.add((_1989_v17).minus((_this).f38));
              }
            }
            return _coll52;
          }());
          let _1990_v21;
          _1990_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,function () {
            let _coll53 = new _dafny.Set();
            for (const _compr_53 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1976_v5).length),_this.f28)).Keys.Elements) {
              let _1991_v20 = _compr_53;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1976_v5).length),_this.f28)).contains(_1991_v20)) {
                _coll53.add((_1991_v20).plus(new BigNumber((_dafny.Seq.of(new BigNumber(454), new BigNumber(198))).length)));
              }
            }
            return _coll53;
          }());
          let _1992_v22;
          _1992_v22 = _dafny.Set.fromElements(_dafny.Set.fromElements((_this).f38), (((_1990_v21).contains((_this).f38)) ? ((_1990_v21).get((_this).f38)) : ((_1988_v18)[_module.__default.safeIndex((_this).f38, new BigNumber((_1988_v18).length))])), _1982_v11, _1982_v11, _dafny.Set.fromElements(new BigNumber(-489), (_this).f38));
          let _index264 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1981_v10).length));
          (_1981_v10)[_index264] = new BigNumber((_dafny.Seq.of((_1992_v22).IsSubsetOf(function () {
            let _coll54 = new _dafny.Set();
            for (const _compr_54 of (_1988_v18).Elements) {
              let _1993_v19 = _compr_54;
              if (_dafny.Seq.contains(_1988_v18, _1993_v19)) {
                _coll54.add(_1993_v19);
              }
            }
            return _coll54;
          }()), ((_this.f28) ? (_this.f31) : (false)), _this.f31)).length);
        } else {
          let _1994_v23;
          _1994_v23 = _dafny.Set.fromElements((_this).f30);
          (globalState).f1 = (_this).fm16(((_this).f34).isLessThanOrEqualTo((_this).f38), ((_this.f29) ? (_module.__default.fm18(_this.f31, _this.f29, globalState)) : (_module.D0.create_DC0(_this.f35))), _1994_v23, (_this).f38, globalState);
          let _1995_v24;
          _1995_v24 = _dafny.Seq.UnicodeFromString("srednb");
          let _1996_v25;
          _1996_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_1995_v24, _1995_v24),!(((_this.f31) ? (_this.f35) : (_this.f28))));
          _1996_v25 = (_1996_v25).update(_this.f31, _this.f29);
          (_this).f36 = _this.f36;
          let _1997_v26;
          _1997_v26 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-92)), function (_1998_i1) {
            return _this.f36;
          })).length)) : (new BigNumber((_dafny.Seq.UnicodeFromString("im")).length))),!(_this.f35));
          r0 = !((((_1997_v26).contains((_this).f38)) ? ((_1997_v26).get((_this).f38)) : (_this.f31)));
          let _1999_v27;
          let _nw342 = Array((new BigNumber(12)).toNumber()).fill(false);
          _1999_v27 = _nw342;
          let _index265 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1999_v27).length));
          (_1999_v27)[_index265] = _this.f29;
          let _index266 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1999_v27).length));
          (_1999_v27)[_index266] = _this.f28;
        }
        let _2000_v28;
        _2000_v28 = _dafny.Seq.UnicodeFromString("bd");
        let _2001_v29;
        let _nw343 = Array((new BigNumber(2)).toNumber());
        _nw343[(_dafny.ZERO).toNumber()] = _2000_v28;
        _nw343[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2000_v28, _2000_v28);
        _2001_v29 = _nw343;
        let _index267 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_2001_v29).length));
        (_2001_v29)[_index267] = _2000_v28;
        let _2002_v30;
        let _nw344 = Array((new BigNumber(16)).toNumber());
        _2002_v30 = _nw344;
        let _2003_v31;
        let _init51 = function (_2004_i2) {
          return _this.f36;
        };
        let _nw345 = Array((new BigNumber(22)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw345.length); _i0_51++) {
          _nw345[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _2003_v31 = _nw345;
        let _2005_v32;
        let _nw346 = new _module.C0();
        _nw346.__ctor(_2003_v31);
        _2005_v32 = _nw346;
        let _index268 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_2002_v30).length));
        (_2002_v30)[_index268] = _2005_v32;
        let _2006_v33;
        _2006_v33 = _dafny.Seq.of((_this).f38, (_this).f30);
        let _2007_v34;
        let _nw347 = new _module.C5();
        _nw347.__ctor(_this.f28, (_this).f34, (_this).f30, _this.f31, (_this).f32, (_this).f33, (_2006_v33)[_module.__default.safeIndex((_this).f30, new BigNumber((_2006_v33).length))], _this.f28, _this.f35, !(_this.f35));
        _2007_v34 = _nw347;
        let _2008_v35;
        _2008_v35 = _dafny.Map.Empty.slice().updateUnsafe(true,_2007_v34);
        let _2009_v36;
        _2009_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2000_v28).length),_module.__default.fm46((_this).f34, (_2007_v34).f30, _this.f35, globalState));
        let _index269 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_2001_v29).length));
        let _index270 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_2002_v30).length));
        let _rhs285 = (_this).fm3(_this.f28, _this.f28, globalState);
        let _rhs286 = (new BigNumber((_2008_v35).length)).multipliedBy(((_this).f38).multipliedBy((((_2009_v36).contains((_2006_v33)[_module.__default.safeIndex((_2007_v34).f30, new BigNumber((_2006_v33).length))])) ? ((_2009_v36).get((_2006_v33)[_module.__default.safeIndex((_2007_v34).f30, new BigNumber((_2006_v33).length))])) : ((_2007_v34).f30))));
        let _rhs287 = (_2007_v34).f30;
        let _rhs288 = _2005_v32;
        let _lhs224 = _2001_v29;
        let _lhs225 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_2001_v29).length));
        let _lhs226 = globalState;
        let _lhs227 = globalState;
        let _lhs228 = _2002_v30;
        let _lhs229 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_2002_v30).length));
        _lhs224[_lhs225] = _rhs285;
        _lhs226.f13 = _rhs286;
        _lhs227.f13 = _rhs287;
        _lhs228[_lhs229] = _rhs288;
        r1 = ((new BigNumber(620)).minus((_this).f38)).multipliedBy((_2007_v34).f30);
        let _2010_v37;
        let _nw348 = Array((new BigNumber(19)).toNumber()).fill(false);
        _2010_v37 = _nw348;
        (globalState).f21 = _2010_v37;
        let _2011_v38;
        _2011_v38 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(_this.f31, _this.f28, _2007_v34.f29)),_this.f29);
        let _2012_v39;
        _2012_v39 = _module.D8.create_DC21(_2011_v38);
        let _2013_v40;
        _2013_v40 = _module.D8.create_DC25(_2012_v39);
        let _rhs289 = _2010_v37;
        let _rhs290 = _2010_v37;
        let _rhs291 = _2013_v40;
        let _rhs292 = _2007_v34.f28;
        let _lhs230 = _this;
        _2010_v37 = _rhs289;
        _2010_v37 = _rhs290;
        _2013_v40 = _rhs291;
        _lhs230.f28 = _rhs292;
      } else {
        let _2014_v41;
        _2014_v41 = _dafny.MultiSet.fromElements((_this).f38, (_this).f38);
        let _2015_v42;
        let _nw349 = Array((new BigNumber(4)).toNumber());
        _nw349[(_dafny.ZERO).toNumber()] = new BigNumber((_2014_v41).cardinality());
        _nw349[(_dafny.ONE).toNumber()] = (_this).f34;
        _nw349[(new BigNumber(2)).toNumber()] = (_this).f34;
        _nw349[(new BigNumber(3)).toNumber()] = (_this).f34;
        _2015_v42 = _nw349;
        let _index271 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length));
        (_2015_v42)[_index271] = (_this).f30;
        let _index272 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length));
        (_2015_v42)[_index272] = (_this).f34;
        (globalState).f1 = false;
        (globalState).f26 = _this.f35;
        let _2016_v43;
        let _nw350 = new _module.C8();
        _nw350.__ctor(_this.f35, _this.f36, (_this).f34, _this.f28, (_this).f32, (_this).f33, (_this).f30, _this.f31, _this.f29, _this.f31);
        _2016_v43 = _nw350;
        let _2017_v44;
        _2017_v44 = _dafny.Map.Empty.slice().updateUnsafe(_2016_v43,(_this).f30);
        let _2018_v45;
        _2018_v45 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f34), (_this).f30, _module.__default.fm24(_this.f35, (_2015_v42)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length))], globalState), (_this).f38, new BigNumber((_2017_v44).length));
        if ((new BigNumber((_2018_v45).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(((_this).f34).plus((_this).f34)))) {
          let _index273 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length));
          let _rhs293 = ((_this).f38).isEqualTo((_this).f34);
          let _rhs294 = (_this).f38;
          let _lhs231 = globalState;
          let _lhs232 = _2015_v42;
          let _lhs233 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length));
          _lhs231.f1 = _rhs293;
          _lhs232[_lhs233] = _rhs294;
          let _index274 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length));
          (_2015_v42)[_index274] = (_this).f34;
          let _2019_v46;
          let _init52 = function (_2020_i3) {
            return _this.f28;
          };
          let _nw351 = Array((new BigNumber(18)).toNumber());
          for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw351.length); _i0_52++) {
            _nw351[_i0_52] = _init52(new BigNumber(_i0_52));
          }
          _2019_v46 = _nw351;
          let _index275 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2019_v46).length));
          (_2019_v46)[_index275] = _this.f35;
          let _2021_v47;
          _2021_v47 = _dafny.Set.fromElements(_dafny.Set.fromElements((_this).f30));
          let _2022_v48;
          let _nw352 = new _module.C7();
          _nw352.__ctor((_this).f34, _2014_v41, (_this).f38, _this.f35, _this.f35, _2016_v43.f40);
          _2022_v48 = _nw352;
          let _2023_v49;
          _2023_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((((_2022_v48.f29) ? ((_this).f32) : (_dafny.Seq.update((_this).f32, _module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length)), _2016_v43.f40)))).length),_2022_v48);
          let _index276 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2019_v46).length));
          let _rhs295 = _this.f29;
          let _rhs296 = (_2021_v47).Intersect(_2021_v47);
          let _rhs297 = (((_2023_v49).contains(_module.__default.safeDivisionInt((_2015_v42)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length))], new BigNumber((_dafny.Seq.UnicodeFromString("okb")).length)))) ? ((_2023_v49).get(_module.__default.safeDivisionInt((_2015_v42)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length))], new BigNumber((_dafny.Seq.UnicodeFromString("okb")).length)))) : (_2022_v48));
          let _lhs234 = _2019_v46;
          let _lhs235 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_2019_v46).length));
          _lhs234[_lhs235] = _rhs295;
          _2021_v47 = _rhs296;
          _2022_v48 = _rhs297;
          let _2024_v50;
          let _nw353 = new _module.C1();
          _nw353.__ctor(_this.f31, true, _this.f29);
          _2024_v50 = _nw353;
          let _2025_v51;
          _2025_v51 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,_dafny.Map.Empty.slice().updateUnsafe(_2024_v50,_module.__default.fm39(_2016_v43.f40, _2014_v41, ((_this).f32)[_module.__default.safeIndex((_this).f38, new BigNumber(((_this).f32).length))], globalState)));
          let _2026_v52;
          _2026_v52 = _dafny.Map.Empty.slice().updateUnsafe(_2024_v50,(_this).f38);
          _2025_v51 = (_2025_v51).update(_this.f36, _2026_v52);
          let _2027_v53;
          _2027_v53 = _dafny.Seq.UnicodeFromString("mhctyq");
          let _2028_v54;
          let _nw354 = Array((new BigNumber(15)).toNumber());
          _nw354[(_dafny.ZERO).toNumber()] = _this.f36;
          _nw354[(_dafny.ONE).toNumber()] = _this.f36;
          _nw354[(new BigNumber(2)).toNumber()] = _this.f36;
          _nw354[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
          _nw354[(new BigNumber(4)).toNumber()] = _this.f36;
          _nw354[(new BigNumber(5)).toNumber()] = _module.__default.fm34((_this).f38, (_this).f38, new BigNumber((_2027_v53).length), _this.f28, globalState);
          _nw354[(new BigNumber(6)).toNumber()] = _this.f36;
          _nw354[(new BigNumber(7)).toNumber()] = _module.__default.fm41(_this.f29, _this.f36, _this.f29, (_this).f38, globalState);
          _nw354[(new BigNumber(8)).toNumber()] = (_2027_v53)[_module.__default.safeIndex((_this).f30, new BigNumber((_2027_v53).length))];
          _nw354[(new BigNumber(9)).toNumber()] = _this.f36;
          _nw354[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
          _nw354[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
          _nw354[(new BigNumber(12)).toNumber()] = (_2027_v53)[_module.__default.safeIndex(new BigNumber((_2014_v41).cardinality()), new BigNumber((_2027_v53).length))];
          _nw354[(new BigNumber(13)).toNumber()] = _this.f36;
          _nw354[(new BigNumber(14)).toNumber()] = _this.f36;
          _2028_v54 = _nw354;
          let _index277 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_2028_v54).length));
          (_2028_v54)[_index277] = _this.f36;
          let _index278 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_2028_v54).length));
          let _rhs298 = _this.f36;
          let _rhs299 = _this.f36;
          let _rhs300 = !(_this.f28) || ((_2019_v46)[_module.__default.safeIndex(new BigNumber(426), new BigNumber((_2019_v46).length))]);
          let _lhs236 = _this;
          let _lhs237 = _2028_v54;
          let _lhs238 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_2028_v54).length));
          let _lhs239 = globalState;
          _lhs236.f36 = _rhs298;
          _lhs237[_lhs238] = _rhs299;
          _lhs239.f26 = _rhs300;
        } else {
          let _2029_v55;
          _2029_v55 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f35);
          let _2030_v56;
          _2030_v56 = _dafny.Set.fromElements(_this.f36, _module.__default.fm41(_this.f31, _this.f36, _this.f35, (_this).f30, globalState));
          let _2031_v57;
          _2031_v57 = _module.D0.create_DC0(_this.f28);
          let _2032_v58;
          _2032_v58 = _dafny.Map.Empty.slice().updateUnsafe((_2015_v42)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length))],_this.f31);
          let _2033_v59;
          _2033_v59 = _dafny.Map.Empty.slice().updateUnsafe((_2015_v42)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length))],_module.__default.fm1(_2016_v43.f40, new BigNumber((_2032_v58).length), globalState));
          let _2034_v60;
          let _nw355 = Array((new BigNumber(26)).toNumber());
          _nw355[(_dafny.ZERO).toNumber()] = _2016_v43.f40;
          _nw355[(_dafny.ONE).toNumber()] = _this.f28;
          _nw355[(new BigNumber(2)).toNumber()] = !(_this.f35);
          _nw355[(new BigNumber(3)).toNumber()] = _this.f35;
          _nw355[(new BigNumber(4)).toNumber()] = _this.f35;
          _nw355[(new BigNumber(5)).toNumber()] = _this.f31;
          _nw355[(new BigNumber(6)).toNumber()] = (((_2029_v55).contains(_this.f35)) ? ((_2029_v55).get(_this.f35)) : (true));
          _nw355[(new BigNumber(7)).toNumber()] = _2016_v43.f40;
          _nw355[(new BigNumber(8)).toNumber()] = !(_2016_v43.f40);
          _nw355[(new BigNumber(9)).toNumber()] = (_2014_v41).contains((_2015_v42)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length))]);
          _nw355[(new BigNumber(10)).toNumber()] = (_2030_v56).contains(new _dafny.CodePoint('s'.codePointAt(0)));
          _nw355[(new BigNumber(11)).toNumber()] = _2016_v43.f40;
          _nw355[(new BigNumber(12)).toNumber()] = _this.f35;
          _nw355[(new BigNumber(13)).toNumber()] = _this.f31;
          _nw355[(new BigNumber(14)).toNumber()] = _this.f29;
          _nw355[(new BigNumber(15)).toNumber()] = (_this).fm16(!(_this.f28), _2031_v57, (_this).fm4(_this.f31, _dafny.Set.fromElements((_this).f30), (_this).f30, _this.f29, globalState), new BigNumber(((_this).f32).length), globalState);
          _nw355[(new BigNumber(16)).toNumber()] = _2016_v43.f40;
          _nw355[(new BigNumber(17)).toNumber()] = _this.f35;
          _nw355[(new BigNumber(18)).toNumber()] = _this.f28;
          _nw355[(new BigNumber(19)).toNumber()] = ((_this).f34).isEqualTo(new BigNumber(848));
          _nw355[(new BigNumber(20)).toNumber()] = !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), function (_2035_i4) {
            return _this.f36;
          }), _dafny.Seq.UnicodeFromString("vdyxsb"));
          _nw355[(new BigNumber(21)).toNumber()] = _this.f29;
          _nw355[(new BigNumber(22)).toNumber()] = (_this.f31) || (_this.f35);
          _nw355[(new BigNumber(23)).toNumber()] = (_2032_v58).contains((_this).f30);
          _nw355[(new BigNumber(24)).toNumber()] = (false) || (!(_this.f35));
          _nw355[(new BigNumber(25)).toNumber()] = !(_this.f31) || (_this.f29);
          _2034_v60 = _nw355;
          let _index279 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_2034_v60).length));
          (_2034_v60)[_index279] = ((true) ? (_this.f35) : (true));
          let _2036_v61;
          _2036_v61 = _dafny.MultiSet.fromElements(_this.f31);
          let _index280 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_2034_v60).length));
          (_2034_v60)[_index280] = (((_2036_v61).IsProperSubsetOf(_2036_v61)) ? (_this.f31) : (_this.f29));
          let _index281 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length));
          (_2015_v42)[_index281] = (_this).f34;
          (_this).f36 = _this.f36;
          (globalState).f12 = ((_2029_v55).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f31,_2016_v43.f40))).equals((_2029_v55).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_2016_v43.f40)));
          (_this).f36 = _this.f36;
        }
        let _index282 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length));
        let _rhs301 = !((!(_this.f35)) || (_this.f28));
        let _rhs302 = (_2015_v42)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length))];
        let _rhs303 = _2016_v43.f40;
        let _lhs240 = _2016_v43;
        let _lhs241 = _2015_v42;
        let _lhs242 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2015_v42).length));
        let _lhs243 = globalState;
        _lhs240.f40 = _rhs301;
        _lhs241[_lhs242] = _rhs302;
        _lhs243.f12 = _rhs303;
      }
      let _2037_v62;
      let _init53 = function (_2038_i5) {
        return ((_this.f35) ? (false) : (_this.f28));
      };
      let _nw356 = Array((new BigNumber(23)).toNumber());
      for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw356.length); _i0_53++) {
        _nw356[_i0_53] = _init53(new BigNumber(_i0_53));
      }
      _2037_v62 = _nw356;
      let _index283 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_2037_v62).length));
      (_2037_v62)[_index283] = _this.f28;
      let _2039_v63;
      _2039_v63 = _dafny.Seq.UnicodeFromString("gxqix");
      let _2040_v64;
      _2040_v64 = _dafny.Seq.of(_2039_v63, _2039_v63);
      let _index284 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_2037_v62).length));
      (_2037_v62)[_index284] = _dafny.areEqual(_2040_v64, _2040_v64);
      let _2041_v65;
      let _nw357 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _2041_v65 = _nw357;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2041_v65).length))) {
        let _2042_i6 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2042_i6)) && ((_2042_i6).isLessThan(new BigNumber((_2041_v65).length))))) {
          (_2041_v65)[(_2042_i6)] = (_2042_i6).minus((_this).f30);
        }
      }
      r1 = new BigNumber((_1970_v0).length);
      r1 = new BigNumber(-371);
      let _2043_v66;
      _2043_v66 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,new BigNumber(((_this).f32).length));
      if ((_2043_v66).equals((_2043_v66).Merge(_2043_v66))) {
        r0 = _this.f28;
        let _2044_v67;
        let _init54 = function (_2045_i7) {
          return _this.f36;
        };
        let _nw358 = Array((new BigNumber(17)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw358.length); _i0_54++) {
          _nw358[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _2044_v67 = _nw358;
        let _2046_v68;
        let _nw359 = new _module.C0();
        _nw359.__ctor(_2044_v67);
        _2046_v68 = _nw359;
        let _2047_v69;
        _2047_v69 = _module.D1.create_DC3(_this.f29, new BigNumber((_2039_v63).length), _2046_v68, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xevgdv")).length)));
        let _2048_v70;
        let _nw360 = new _module.C4();
        _nw360.__ctor((_this).f34, (_this).fm6(globalState), _dafny.Seq.Concat(_dafny.Seq.of(_this.f35), (_this).f32), _dafny.Seq.Concat(_dafny.Seq.of((_this).f32), _dafny.Seq.of((_this).f32, (_this).f32)), _module.__default.safeModuloInt((_this).f34, (_this).f34), (_2047_v69).dtor_cf3, _this.f35, _this.f31);
        _2048_v70 = _nw360;
        let _rhs304 = _2048_v70;
        let _rhs305 = _this.f31;
        let _rhs306 = _2041_v65;
        let _lhs244 = _this;
        _2048_v70 = _rhs304;
        _lhs244.f31 = _rhs305;
        _2041_v65 = _rhs306;
        let _2049_v71;
        _2049_v71 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,_2041_v65);
        _2041_v65 = (((_2049_v71).contains((_this).f38)) ? ((_2049_v71).get((_this).f38)) : (_2041_v65));
        (_this).f28 = (_1970_v0).IsProperSubsetOf(_1970_v0);
        let _2050_v72;
        let _nw361 = Array((new BigNumber(14)).toNumber()).fill([]);
        _2050_v72 = _nw361;
        let _index285 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_2050_v72).length));
        (_2050_v72)[_index285] = _2037_v62;
        let _index286 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_2050_v72).length));
        (_2050_v72)[_index286] = _2037_v62;
      } else {
        let _2051_v73;
        _2051_v73 = _dafny.Map.Empty.slice().updateUnsafe((_2037_v62)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_2037_v62).length))],(_dafny.ZERO).minus((_this).f34));
        let _2052_v74;
        _2052_v74 = _dafny.Seq.of(new BigNumber((_2051_v73).length), _module.__default.safeDivisionInt((_this).f34, new BigNumber((_dafny.MultiSet.fromElements(_2041_v65)).cardinality())), (_this).f38);
        let _2053_v75;
        _2053_v75 = _dafny.Seq.of(_2052_v74, _2052_v74);
        _2052_v74 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2052_v74, _dafny.Seq.Create(_module.__default.abs(new BigNumber(832)), function (_2054_i8) {
          return (_this).f38;
        })), _dafny.Seq.Concat(_2052_v74, (_2053_v75)[_module.__default.safeIndex((_this).f30, new BigNumber((_2053_v75).length))]));
        let _index287 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_2041_v65).length));
        (_2041_v65)[_index287] = new BigNumber((_1970_v0).length);
        let _index288 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_2041_v65).length));
        (_2041_v65)[_index288] = _module.__default.safeDivisionInt(new BigNumber(-825), (_this).f34);
        let _2055_v76;
        _2055_v76 = _dafny.Map.Empty.slice().updateUnsafe((_2052_v74)[_module.__default.safeIndex((_this).f30, new BigNumber((_2052_v74).length))],_this.f29);
        let _2056_v77;
        _2056_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2037_v62,new BigNumber((_2055_v76).length));
        let _2057_v78;
        _2057_v78 = _dafny.Map.Empty.slice().updateUnsafe((_2041_v65)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_2041_v65).length))],_dafny.Seq.Create(_module.__default.abs(new BigNumber(836)), function (_2058_i9) {
          return _this.f36;
        }));
        let _2059_v79;
        _2059_v79 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,new _dafny.CodePoint('h'.codePointAt(0)));
        let _2060_v80;
        let _nw362 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _2060_v80 = _nw362;
        let _2061_v81;
        let _nw363 = new _module.C5();
        _nw363.__ctor(((_this).f34).isEqualTo(new BigNumber(793)), new BigNumber((_2056_v77).length), new BigNumber(((((_2057_v78).contains((_this).f38)) ? ((_2057_v78).get((_this).f38)) : (_2039_v63))).length), (_2037_v62)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_2037_v62).length))], (_this).f32, _dafny.Seq.of(_dafny.Seq.of(_this.f31, _this.f29), _dafny.Seq.update((_this).f32, _module.__default.safeIndex(new BigNumber((_2059_v79).length), new BigNumber(((_this).f32).length)), _this.f28), (_this).f32), new BigNumber((_dafny.Seq.update(_2052_v74, _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_2060_v80)).length), new BigNumber((_2052_v74).length)), (_2052_v74)[_module.__default.safeIndex((_this).f34, new BigNumber((_2052_v74).length))])).length), (new BigNumber(((_this).f32).length)).isLessThan((_this).f30), _this.f28, ((_this).f32)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length))]);
        _2061_v81 = _nw363;
        let _2062_v82;
        _2062_v82 = _dafny.MultiSet.fromElements((_this).f30, new BigNumber((_2039_v63).length));
        let _2063_v83;
        _2063_v83 = _dafny.MultiSet.fromElements(new BigNumber(290), new BigNumber((_2062_v82).cardinality()));
        (globalState).f12 = (_this).fm7(_2062_v82, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_2052_v74, _2052_v74)).length)), globalState);
        let _2064_v84;
        _2064_v84 = _module.D8.create_DC22((_this).f34, (_this).f30);
        let _2065_v85;
        _2065_v85 = _module.D2.create_DC6((_2064_v84).dtor_cf41, _this.f28, (_2061_v81).f48, (_2061_v81).f47, new BigNumber(-48));
        let _pat_let_tv18 = _2055_v76;
        let _pat_let_tv19 = _2055_v76;
        _2065_v85 = function (_pat_let25_0) {
          return function (_2066_dt__update__tmp_h0) {
            return function (_pat_let26_0) {
              return function (_2067_dt__update_hcf9_h0) {
                return function (_pat_let27_0) {
                  return function (_2068_dt__update_hcf12_h0) {
                    return function (_pat_let28_0) {
                      return function (_2069_dt__update_hcf13_h0) {
                        return _module.D2.create_DC6(_2067_dt__update_hcf9_h0, (_2066_dt__update__tmp_h0).dtor_cf10, (_2066_dt__update__tmp_h0).dtor_cf11, _2068_dt__update_hcf12_h0, _2069_dt__update_hcf13_h0);
                      }(_pat_let28_0);
                    }(new BigNumber(((_pat_let_tv18).Merge(_pat_let_tv19)).length));
                  }(_pat_let27_0);
                }(!(_this.f29));
              }(_pat_let26_0);
            }((_this).f38);
          }(_pat_let25_0);
        }(_2065_v85);
      }
      r0 = _this.f31;
      r1 = (_this).f30;
      let _2070_v86;
      _2070_v86 = _dafny.MultiSet.fromElements((_this).f34, (_dafny.ZERO).minus((_this).f38));
      let _2071_v87;
      _2071_v87 = _module.D14.create_DC37(_module.__default.fm46((_this).f34, (_this).f38, _this.f29, globalState), (_this).fm7(_2070_v86, (_this).f30, globalState), new BigNumber((_dafny.Seq.UnicodeFromString("mfybmsjun")).length));
      let _pat_let_tv20 = _2039_v63;
      let _pat_let_tv21 = _2039_v63;
      r2 = function (_source33) {
        if (_source33.is_DC37) {
          let _2072___mcc_h0 = (_source33).cf68;
          let _2073___mcc_h1 = (_source33).cf69;
          let _2074___mcc_h2 = (_source33).cf70;
          let _2075_cf70 = _2074___mcc_h2;
          let _2076_cf69 = _2073___mcc_h1;
          let _2077_cf68 = _2072___mcc_h0;
          return _pat_let_tv20;
        } else {
          let _2078___mcc_h3 = (_source33).cf67;
          let _2079_cf67 = _2078___mcc_h3;
          return _pat_let_tv21;
        }
      }(_2071_v87);
      return [r0, r1, r2];
    }
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      (_this).f36 = _this.f36;
      (globalState).f13 = (_dafny.ZERO).minus((_this).f38);
      (_this).f28 = _this.f29;
      let _2080_v0;
      _2080_v0 = _module.D0.create_DC0(_this.f35);
      if ((_2080_v0).dtor_cf0) {
        let _2081_v1;
        _2081_v1 = _dafny.Seq.of((_this).f34, (_this).f34);
        let _2082_v2;
        _2082_v2 = _dafny.Set.fromElements(_2081_v1);
        let _2083_v3;
        _2083_v3 = _module.D16.create_DC39(_2082_v2);
        let _2084_v4;
        _2084_v4 = _module.D16.create_DC41(_2083_v3);
        let _2085_v5;
        _2085_v5 = _module.D16.create_DC41(_module.D16.create_DC41(_2084_v4));
        let _2086_v6;
        let _nw364 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2086_v6 = _nw364;
        let _2087_v7;
        _2087_v7 = _module.D20.create_DC47(_2086_v6);
        let _pat_let_tv22 = _2086_v6;
        let _rhs307 = _2085_v5;
        let _rhs308 = (_this).f38;
        let _rhs309 = ((_this).f32)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length))];
        let _rhs310 = (function (_pat_let29_0) {
          return function (_2088_dt__update__tmp_h0) {
            return function (_pat_let30_0) {
              return function (_2089_dt__update_hcf81_h0) {
                return _module.D20.create_DC47(_2089_dt__update_hcf81_h0);
              }(_pat_let30_0);
            }(_pat_let_tv22);
          }(_pat_let29_0);
        }(_2087_v7)).dtor_cf81;
        let _rhs311 = (_this).f30;
        let _lhs245 = globalState;
        let _lhs246 = globalState;
        let _lhs247 = globalState;
        _2085_v5 = _rhs307;
        _lhs245.f13 = _rhs308;
        _lhs246.f1 = _rhs309;
        _2086_v6 = _rhs310;
        _lhs247.f13 = _rhs311;
        let _2090_v8;
        _2090_v8 = _dafny.MultiSet.fromElements(false, _this.f28);
        if ((_dafny.MultiSet.fromElements(_this.f31, _this.f35, _this.f35)).IsProperSubsetOf(_2090_v8)) {
          (globalState).f13 = (_this).f34;
          let _2091_v9;
          _2091_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,new BigNumber((((_this).f33)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f30), new BigNumber(((_this).f33).length))]).length));
          let _2092_v10;
          _2092_v10 = _dafny.Seq.of(_2091_v9);
          (_this).f28 = (new BigNumber((_2092_v10).length)).isLessThanOrEqualTo((_this).f38);
          let _2093_v11;
          _2093_v11 = _module.D20.create_DC49(_this.f29, _this.f36);
          let _2094_v12;
          let _nw365 = new _module.C1();
          _nw365.__ctor(_this.f29, _dafny.Seq.contains((_this).f32, _this.f28), !((_2093_v11).dtor_cf85));
          _2094_v12 = _nw365;
          (globalState).f1 = _this.f35;
          let _2095_v13;
          let _nw366 = Array((new BigNumber(10)).toNumber()).fill(false);
          _2095_v13 = _nw366;
          let _index289 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_2095_v13).length));
          (_2095_v13)[_index289] = true;
          let _index290 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_2095_v13).length));
          (_2095_v13)[_index290] = _this.f35;
          let _2096_v14;
          _2096_v14 = _dafny.Seq.UnicodeFromString("wc");
          let _index291 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_2095_v13).length));
          let _index292 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_2095_v13).length));
          let _rhs312 = ((_this).f34).isLessThanOrEqualTo((_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_2081_v1).length))).minus((_dafny.ZERO).minus((_this).f38))));
          let _rhs313 = _this.f31;
          let _rhs314 = (_module.__default.safeDivisionInt((_this).f38, new BigNumber((_2081_v1).length))).multipliedBy((_this).f38);
          let _rhs315 = _dafny.Seq.UnicodeFromString("e");
          let _rhs316 = (_this).f30;
          let _lhs248 = _2095_v13;
          let _lhs249 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_2095_v13).length));
          let _lhs250 = _2095_v13;
          let _lhs251 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_2095_v13).length));
          let _lhs252 = globalState;
          let _lhs253 = globalState;
          _lhs248[_lhs249] = _rhs312;
          _lhs250[_lhs251] = _rhs313;
          _lhs252.f13 = _rhs314;
          _2096_v14 = _rhs315;
          _lhs253.f13 = _rhs316;
        } else {
          let _2097_v15;
          let _nw367 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2097_v15 = _nw367;
          let _index293 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_2097_v15).length));
          (_2097_v15)[_index293] = _2090_v8;
          let _index294 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_2097_v15).length));
          (_2097_v15)[_index294] = _2090_v8;
          (_this).f28 = _this.f35;
          let _2098_v16;
          _2098_v16 = _module.D8.create_DC24(_this.f31, (_this).f34, (_this).f32, _this.f35, _2081_v1);
          let _2099_v17;
          _2099_v17 = _module.D8.create_DC25(_2098_v16);
          let _2100_v18;
          _2100_v18 = _dafny.Seq.of(_2099_v17, _2099_v17);
          let _2101_v19;
          _2101_v19 = _dafny.MultiSet.fromElements(_2100_v18, _2100_v18);
          (globalState).f26 = ((_2101_v19).Difference(_2101_v19)).IsProperSubsetOf(_2101_v19);
          let _2102_v20;
          let _nw368 = Array((new BigNumber(29)).toNumber());
          _2102_v20 = _nw368;
          let _2103_v21;
          _2103_v21 = _dafny.MultiSet.fromElements(_2102_v20);
          (globalState).f26 = !((_2103_v21).IsProperSubsetOf(_2103_v21));
          let _2104_v22;
          _2104_v22 = _dafny.MultiSet.fromElements((_this).f34, (_this).f30);
          let _2105_v23;
          _2105_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(((_2104_v22).update(new BigNumber(-797), _module.__default.abs((_this).f38))).cardinality())), (_this).f38));
          let _2106_v24;
          _2106_v24 = _dafny.Set.fromElements(_this.f31);
          (globalState).f13 = (((_2105_v23).contains(new BigNumber(((_2106_v24).Intersect(_dafny.Set.fromElements(_this.f35))).length))) ? ((_2105_v23).get(new BigNumber(((_2106_v24).Intersect(_dafny.Set.fromElements(_this.f35))).length))) : (((_this).f34).multipliedBy((_this).f38)));
        }
        let _2107_v25;
        let _init55 = ((_2108_v1) => function (_2109_i0) {
          return (_dafny.Map.Empty.slice().updateUnsafe((_this).f34,new BigNumber(238))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_2108_v1)[_module.__default.safeIndex(new BigNumber(-294), new BigNumber((_2108_v1).length))],(_this).f30));
        })(_2081_v1);
        let _nw369 = Array((new BigNumber(13)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw369.length); _i0_55++) {
          _nw369[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _2107_v25 = _nw369;
        let _2110_v26;
        _2110_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(713),(_this).f34);
        let _index295 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_2107_v25).length));
        (_2107_v25)[_index295] = (_2110_v26).Merge(_2110_v26);
        let _index296 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_2107_v25).length));
        (_2107_v25)[_index296] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f32).length),(_this).f38);
        let _2111_v27;
        _2111_v27 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f30));
        let _2112_v28;
        _2112_v28 = _module.D14.create_DC37((((_2111_v27).contains((_this).f38)) ? ((_2111_v27).get((_this).f38)) : ((_this).f34)), _this.f35, _module.__default.safeModuloInt(new BigNumber(625), (_this).f38));
        _2112_v28 = _module.__default.fm58(_this.f28, ((_this).f33)[_module.__default.safeIndex((_this).f38, new BigNumber(((_this).f33).length))], globalState);
        let _2113_v29;
        _2113_v29 = _dafny.Set.fromElements((_this).f30);
        let _2114_v30;
        let _out69;
        _out69 = _module.__default.m0(_2081_v1, _this.f28, (_this).f34, (_2113_v29).contains((_this).f38), globalState);
        _2114_v30 = _out69;
      } else {
        let _nw370 = Array((new BigNumber(27)).toNumber()).fill(false);
        (globalState).f21 = _nw370;
        if (_this.f35) {
          let _2115_v31;
          _2115_v31 = _dafny.Seq.of((_this).f34);
          let _2116_v32;
          let _out70;
          _out70 = _module.__default.m0(_dafny.Seq.Concat(_2115_v31, _2115_v31), _this.f31, (_this).f38, _this.f35, globalState);
          _2116_v32 = _out70;
          (globalState).f26 = false;
          (_this).f35 = !(_this.f31);
          let _2117_v33;
          _2117_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f38);
          let _2118_v35;
          _2118_v35 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,function () {
            let _coll55 = new _dafny.Set();
            for (const _compr_55 of (_2117_v33).Keys.Elements) {
              let _2119_v34 = _compr_55;
              if ((_2117_v33).contains(_2119_v34)) {
                _coll55.add(_module.__default.safeDivisionInt(_2119_v34, new BigNumber((_dafny.Seq.of(true, !(true))).length)));
              }
            }
            return _coll55;
          }());
          let _2120_v36;
          _2120_v36 = _module.D2.create_DC6((_dafny.ZERO).minus(new BigNumber((((_this.f31) ? (_2118_v35) : (_2118_v35))).length)), ((_this).f32)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length))], (_this).f34, _this.f31, _module.__default.safeDivisionInt(_2116_v32, new BigNumber(423)));
          let _2121_v37;
          _2121_v37 = _dafny.Set.fromElements(_this.f31);
          let _2122_v38;
          _2122_v38 = _dafny.Seq.of(_2121_v37, _dafny.Set.fromElements(_this.f28), _2121_v37, _module.__default.fm45((_this).f38, _dafny.MultiSet.FromArray((_this).f32), _this.f28, (_dafny.ZERO).minus(_2116_v32), globalState), _2121_v37);
          let _2123_v39;
          _2123_v39 = _dafny.Set.fromElements((_2122_v38)[_module.__default.safeIndex(new BigNumber(((_this).f32).length), new BigNumber((_2122_v38).length))]);
          let _2124_v40;
          let _init56 = function (_2125_i1) {
            return _this.f35;
          };
          let _nw371 = Array((new BigNumber(3)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw371.length); _i0_56++) {
            _nw371[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _2124_v40 = _nw371;
          let _2126_v41;
          _2126_v41 = _dafny.Seq.of(_2124_v40);
          let _2127_v42;
          _2127_v42 = _dafny.Seq.of((_2126_v41)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_2126_v41).length))], _2124_v40);
          _2120_v36 = _module.D2.create_DC6(new BigNumber((_2123_v39).length), _this.f28, ((!(_this.f29)) ? (new BigNumber(((_this).f32).length)) : (new BigNumber(968))), _this.f29, new BigNumber((_2127_v42).length));
          let _2128_v43;
          _2128_v43 = _dafny.Map.Empty.slice().updateUnsafe(_2115_v31,true);
          _2128_v43 = ((_2128_v43).update(_2115_v31, _this.f28)).update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(991)), function (_2129_i2) {
            return (_this).f30;
          }), _2115_v31), _this.f35);
        } else {
          let _2130_v44;
          _2130_v44 = _dafny.Set.fromElements(_this.f35);
          let _2131_v45;
          _2131_v45 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,(_this).f30);
          let _2132_v46;
          _2132_v46 = _dafny.Seq.UnicodeFromString("sk");
          let _2133_v47;
          let _nw372 = Array((new BigNumber(11)).toNumber());
          _nw372[(_dafny.ZERO).toNumber()] = (_this).f38;
          _nw372[(_dafny.ONE).toNumber()] = (_this).f34;
          _nw372[(new BigNumber(2)).toNumber()] = (_this).f38;
          _nw372[(new BigNumber(3)).toNumber()] = new BigNumber((_2130_v44).length);
          _nw372[(new BigNumber(4)).toNumber()] = new BigNumber((_2131_v45).length);
          _nw372[(new BigNumber(5)).toNumber()] = (_this).f38;
          _nw372[(new BigNumber(6)).toNumber()] = new BigNumber((_2132_v46).length);
          _nw372[(new BigNumber(7)).toNumber()] = (_this).f38;
          _nw372[(new BigNumber(8)).toNumber()] = new BigNumber((_2132_v46).length);
          _nw372[(new BigNumber(9)).toNumber()] = (_this).f30;
          _nw372[(new BigNumber(10)).toNumber()] = (_this).f38;
          _2133_v47 = _nw372;
          let _2134_v48;
          _2134_v48 = _dafny.MultiSet.fromElements(_2133_v47);
          let _2135_v50;
          let _nw373 = Array((new BigNumber(16)).toNumber());
          _nw373[(_dafny.ZERO).toNumber()] = (_this).f34;
          _nw373[(_dafny.ONE).toNumber()] = (new BigNumber((_2134_v48).cardinality())).minus((_this).f38);
          _nw373[(new BigNumber(2)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("lyvg")).length)).plus(new BigNumber(((_this).f33).length));
          _nw373[(new BigNumber(3)).toNumber()] = new BigNumber(140);
          _nw373[(new BigNumber(4)).toNumber()] = (_this).f38;
          _nw373[(new BigNumber(5)).toNumber()] = ((_this).f30).multipliedBy((_this).f30);
          _nw373[(new BigNumber(6)).toNumber()] = (_this).f30;
          _nw373[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), function (_2136_i3) {
            return _this.f36;
          }), _2132_v46)).length);
          _nw373[(new BigNumber(8)).toNumber()] = (_this).f34;
          _nw373[(new BigNumber(9)).toNumber()] = (_this).f30;
          _nw373[(new BigNumber(10)).toNumber()] = (_this).f38;
          _nw373[(new BigNumber(11)).toNumber()] = new BigNumber((_2132_v46).length);
          _nw373[(new BigNumber(12)).toNumber()] = (_this).f34;
          _nw373[(new BigNumber(13)).toNumber()] = (_this).f30;
          _nw373[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt((_this).f30, (_this).f30);
          _nw373[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,new BigNumber((function () {
            let _coll56 = new _dafny.Map();
            for (const _compr_56 of _dafny.IntegerRange(new BigNumber(-644), new BigNumber(310))) {
              let _2137_v49 = _compr_56;
              if (((new BigNumber(-644)).isLessThanOrEqualTo(_2137_v49)) && ((_2137_v49).isLessThan(new BigNumber(310)))) {
                _coll56.push([(_2137_v49).minus((_this).f30),_this.f31]);
              }
            }
            return _coll56;
          }()).length))).length);
          _2135_v50 = _nw373;
          let _index297 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_2135_v50).length));
          (_2135_v50)[_index297] = (_this).f38;
          let _2138_v51;
          _2138_v51 = _dafny.Set.fromElements((_this).f30);
          let _2139_v52;
          _2139_v52 = _module.D21.create_DC51(_module.__default.fm59(_this.f35, _this.f35, _this.f28, globalState));
          let _2140_v53;
          _2140_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_2138_v51);
          let _index298 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_2135_v50).length));
          let _rhs317 = !(_2138_v51).contains((_this).f34);
          let _rhs318 = false;
          let _rhs319 = _module.__default.fm48(!((_2139_v52).dtor_cf88).equals((_2140_v53).update((_this).f38, function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of _dafny.IntegerRange(new BigNumber(-564), new BigNumber(901))) {
              let _2141_v54 = _compr_57;
              if (((new BigNumber(-564)).isLessThanOrEqualTo(_2141_v54)) && ((_2141_v54).isLessThan(new BigNumber(901)))) {
                _coll57.add(_module.__default.safeModuloInt(_2141_v54, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f35)).length)));
              }
            }
            return _coll57;
          }())), _this.f35, (_this).f38, globalState);
          let _rhs320 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(756)), function (_2142_i4) {
            return (_this).f38;
          })).length);
          let _lhs254 = globalState;
          let _lhs255 = globalState;
          let _lhs256 = _this;
          let _lhs257 = _2135_v50;
          let _lhs258 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_2135_v50).length));
          _lhs254.f12 = _rhs317;
          _lhs255.f12 = _rhs318;
          _lhs256.f36 = _rhs319;
          _lhs257[_lhs258] = _rhs320;
          (_this).f36 = _this.f36;
          let _2143_v55;
          let _init57 = ((_2144_v46) => function (_2145_i5) {
            return _2144_v46;
          })(_2132_v46);
          let _nw374 = Array((new BigNumber(3)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw374.length); _i0_57++) {
            _nw374[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _2143_v55 = _nw374;
          let _index299 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2143_v55).length));
          (_2143_v55)[_index299] = _2132_v46;
          let _index300 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2143_v55).length));
          (_2143_v55)[_index300] = (_this).fm3(!(!(_this.f35)), false, globalState);
          let _2146_v56;
          let _init58 = function (_2147_i6) {
            return _dafny.MultiSet.fromElements(_this.f29);
          };
          let _nw375 = Array((new BigNumber(6)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw375.length); _i0_58++) {
            _nw375[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _2146_v56 = _nw375;
          let _2148_v57;
          _2148_v57 = _dafny.MultiSet.fromElements(_this.f28, _this.f31, _this.f31);
          let _index301 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_2146_v56).length));
          (_2146_v56)[_index301] = _2148_v57;
          let _2149_v58;
          let _nw376 = Array((new BigNumber(25)).toNumber()).fill(false);
          _2149_v58 = _nw376;
          let _index302 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_2149_v58).length));
          (_2149_v58)[_index302] = _this.f28;
          let _index303 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_2146_v56).length));
          let _index304 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_2135_v50).length));
          let _index305 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_2149_v58).length));
          let _rhs321 = (_2143_v55)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_2143_v55).length))];
          let _rhs322 = (_2148_v57).update(_this.f31, _module.__default.abs((_this).f38));
          let _rhs323 = (new BigNumber(939)).isLessThan((_dafny.ZERO).minus((_this).f38));
          let _rhs324 = (_2135_v50)[_module.__default.safeIndex(new BigNumber(832), new BigNumber((_2135_v50).length))];
          let _rhs325 = !(!(_this.f28));
          let _lhs259 = _2146_v56;
          let _lhs260 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_2146_v56).length));
          let _lhs261 = globalState;
          let _lhs262 = _2135_v50;
          let _lhs263 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_2135_v50).length));
          let _lhs264 = _2149_v58;
          let _lhs265 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_2149_v58).length));
          _2132_v46 = _rhs321;
          _lhs259[_lhs260] = _rhs322;
          _lhs261.f1 = _rhs323;
          _lhs262[_lhs263] = _rhs324;
          _lhs264[_lhs265] = _rhs325;
          let _2150_v59;
          _2150_v59 = _dafny.MultiSet.fromElements((_this).f34);
          let _2151_v60;
          _2151_v60 = _module.D17.create_DC44(_2150_v59, (_this).f34);
          let _2152_v61;
          _2152_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_this.f28), (_this).f32),_this.f29);
          let _index306 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_2135_v50).length));
          let _rhs326 = function (_pat_let31_0) {
            return function (_2153_dt__update__tmp_h1) {
              return function (_pat_let32_0) {
                return function (_2154_dt__update_hcf78_h0) {
                  return _module.D17.create_DC44((_2153_dt__update__tmp_h1).dtor_cf77, _2154_dt__update_hcf78_h0);
                }(_pat_let32_0);
              }(new BigNumber(695));
            }(_pat_let31_0);
          }(_2151_v60);
          let _rhs327 = new BigNumber((_2152_v61).length);
          let _lhs266 = _2135_v50;
          let _lhs267 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_2135_v50).length));
          _2151_v60 = _rhs326;
          _lhs266[_lhs267] = _rhs327;
        }
        let _2155_v62;
        _2155_v62 = _dafny.Seq.UnicodeFromString("rvpm");
        _2155_v62 = _dafny.Seq.UnicodeFromString("nnhqon");
        let _2156_v63;
        let _init59 = ((_2157_v62) => function (_2158_i7) {
          return (_dafny.MultiSet.fromElements(_this.f36)).IsProperSubsetOf(_dafny.MultiSet.fromElements((_2157_v62)[_module.__default.safeIndex((_this).f38, new BigNumber((_2157_v62).length))]));
        })(_2155_v62);
        let _nw377 = Array((new BigNumber(26)).toNumber());
        for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw377.length); _i0_59++) {
          _nw377[_i0_59] = _init59(new BigNumber(_i0_59));
        }
        _2156_v63 = _nw377;
        let _index307 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_2156_v63).length));
        (_2156_v63)[_index307] = _this.f28;
        let _index308 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2156_v63).length));
        (_2156_v63)[_index308] = _this.f31;
        let _2159_v64;
        _2159_v64 = _dafny.MultiSet.fromElements(_this.f29, _this.f29);
        let _2160_v65;
        _2160_v65 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f34), (((_2159_v64).contains(_this.f31)) ? ((_2159_v64).get(_this.f31)) : (new BigNumber((_dafny.Set.fromElements(false)).length))), new BigNumber(431));
        let _index309 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_2156_v63).length));
        let _index310 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2156_v63).length));
        let _rhs328 = _module.__default.safeDivisionInt(new BigNumber((_2160_v65).length), (_2160_v65)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_2160_v65).length))]);
        let _rhs329 = ((_this).f34).isLessThanOrEqualTo((_dafny.ZERO).minus((_this).f38));
        let _rhs330 = _this.f28;
        let _lhs268 = globalState;
        let _lhs269 = _2156_v63;
        let _lhs270 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_2156_v63).length));
        let _lhs271 = _2156_v63;
        let _lhs272 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2156_v63).length));
        _lhs268.f13 = _rhs328;
        _lhs269[_lhs270] = _rhs329;
        _lhs271[_lhs272] = _rhs330;
        let _2161_v66;
        _2161_v66 = _dafny.MultiSet.fromElements((_this).f30);
        let _rhs331 = (_dafny.MultiSet.fromElements((_this).f38)).Difference(_2161_v66);
        let _rhs332 = new BigNumber(224);
        let _lhs273 = globalState;
        _2161_v66 = _rhs331;
        _lhs273.f13 = _rhs332;
      }
      (globalState).f26 = _this.f28;
      let _2162_v67;
      _2162_v67 = _dafny.Seq.of(_this.f29);
      let _2163_v68;
      let _init60 = function (_2164_i8) {
        return (_2164_i8).plus((_this).f30);
      };
      let _nw378 = Array((new BigNumber(8)).toNumber());
      for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw378.length); _i0_60++) {
        _nw378[_i0_60] = _init60(new BigNumber(_i0_60));
      }
      _2163_v68 = _nw378;
      let _2165_v69;
      _2165_v69 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f34), (_this).f38);
      let _index311 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_2163_v68).length));
      (_2163_v68)[_index311] = new BigNumber((_2165_v69).cardinality());
      let _index312 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_2163_v68).length));
      let _rhs333 = _2162_v67;
      let _rhs334 = (_this).f30;
      let _rhs335 = (_dafny.ZERO).minus((_this).f34);
      let _rhs336 = (_this).f38;
      let _rhs337 = (_this).f30;
      let _lhs274 = globalState;
      let _lhs275 = globalState;
      let _lhs276 = globalState;
      let _lhs277 = _2163_v68;
      let _lhs278 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_2163_v68).length));
      _2162_v67 = _rhs333;
      _lhs274.f13 = _rhs334;
      _lhs275.f13 = _rhs335;
      _lhs276.f13 = _rhs336;
      _lhs277[_lhs278] = _rhs337;
      let _2166_v70;
      _2166_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,_this.f35);
      let _2167_v71;
      _2167_v71 = _dafny.Map.Empty.slice().updateUnsafe(_2166_v70,(_2163_v68)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_2163_v68).length))]);
      let _2168_v72;
      _2168_v72 = _dafny.Seq.of(_2167_v71, (_2167_v71).Merge(_2167_v71), _2167_v71, _2167_v71);
      r0 = (_2168_v72)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_2168_v72).length))];
      return r0;
    }
    m5(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = [];
      let _2169_v0;
      let _nw379 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _2169_v0 = _nw379;
      let _2170_v1;
      _2170_v1 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(778)), function (_2171_i0) {
        return _this.f36;
      }),((false) ? (_2169_v0) : (_2169_v0)));
      _2170_v1 = (_2170_v1).update((_this).fm3(_this.f28, _this.f31, globalState), _2169_v0);
      let _2172_v2;
      let _init61 = function (_2173_i1) {
        return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f28,_this.f31)).length)).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f34)));
      };
      let _nw380 = Array((new BigNumber(4)).toNumber());
      for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw380.length); _i0_61++) {
        _nw380[_i0_61] = _init61(new BigNumber(_i0_61));
      }
      _2172_v2 = _nw380;
      let _index313 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length));
      (_2172_v2)[_index313] = !((_this.f31) && (true));
      let _2174_v3;
      _2174_v3 = _dafny.MultiSet.fromElements(_this.f28);
      let _2175_v4;
      _2175_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_2174_v3);
      let _index314 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length));
      (_2169_v0)[_index314] = new BigNumber(((_this).fm3(_this.f35, ((_this).f32)[_module.__default.safeIndex((_this).f30, new BigNumber(((_this).f32).length))], globalState)).length);
      let _2176_v5;
      _2176_v5 = _dafny.MultiSet.fromElements((_this).f38, new BigNumber(325));
      let _index315 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2169_v0).length));
      (_2169_v0)[_index315] = _module.__default.fm39(_this.f31, _2176_v5, _this.f28, globalState);
      let _index316 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length));
      let _index317 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length));
      let _index318 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2169_v0).length));
      let _rhs338 = _this.f36;
      let _rhs339 = _this.f29;
      let _rhs340 = ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-867),_2174_v3)).Merge(_2175_v4)).Merge((_2175_v4).Merge(_2175_v4));
      let _rhs341 = (_this).f30;
      let _rhs342 = (_this).f38;
      let _lhs279 = _this;
      let _lhs280 = _2172_v2;
      let _lhs281 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length));
      let _lhs282 = _2169_v0;
      let _lhs283 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length));
      let _lhs284 = _2169_v0;
      let _lhs285 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2169_v0).length));
      _lhs279.f36 = _rhs338;
      _lhs280[_lhs281] = _rhs339;
      _2175_v4 = _rhs340;
      _lhs282[_lhs283] = _rhs341;
      _lhs284[_lhs285] = _rhs342;
      if (_dafny.Seq.IsProperPrefixOf((_this).f32, (_this).f32)) {
        (globalState).f13 = (_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))];
        _2169_v0 = _2169_v0;
        let _2177_v6;
        _2177_v6 = _dafny.Seq.of(_this.f36, _this.f36);
        let _2178_v7;
        _2178_v7 = _dafny.Seq.of(true, false, ((!(_this.f35)) ? (_this.f29) : (_this.f29)), _this.f28);
        let _2179_v8;
        _2179_v8 = _dafny.Set.fromElements((_this).f34);
        let _rhs343 = _dafny.Seq.Concat(_dafny.Seq.of(_this.f36), _2177_v6);
        let _rhs344 = ((_this).f34).isLessThan((_dafny.ZERO).minus((_this).f38));
        let _rhs345 = _dafny.Seq.Concat((_this).f32, _2178_v7);
        let _rhs346 = (_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))];
        let _rhs347 = _2179_v8;
        let _lhs286 = globalState;
        let _lhs287 = globalState;
        _2177_v6 = _rhs343;
        _lhs286.f26 = _rhs344;
        _2178_v7 = _rhs345;
        _lhs287.f13 = _rhs346;
        _2179_v8 = _rhs347;
        let _2180_v9;
        _2180_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2169_v0,(_this).fm3(_this.f31, (_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))], globalState));
        _2180_v9 = (_2180_v9).update(_2169_v0, _2177_v6);
        let _2181_v10;
        _2181_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(995),!((_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))]));
        _2181_v10 = (_2181_v10).update((_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))], _this.f35);
      } else {
        let _2182_v11;
        _2182_v11 = _dafny.Seq.UnicodeFromString("aki");
        let _2183_v12;
        _2183_v12 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ox")).length));
        let _2184_v13;
        _2184_v13 = _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus((_2183_v12)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f30), new BigNumber((_2183_v12).length))])));
        let _rhs348 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), function (_2185_i2) {
          return _this.f36;
        }), _2182_v11);
        let _rhs349 = !((_this.f28) || ((_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))])) || (!(_module.__default.fm60(globalState)).equals((_dafny.Map.Empty.slice().updateUnsafe(_this.f28,_2183_v12)).update((_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))], _2183_v12)));
        let _rhs350 = (_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))];
        let _lhs288 = _this;
        _2182_v11 = _rhs348;
        _lhs288.f35 = _rhs349;
        r0 = _rhs350;
        let _index319 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length));
        (_2169_v0)[_index319] = (_dafny.ZERO).minus(((_this).f38).plus(new BigNumber(((_this).f32).length)));
        let _2186_v14;
        let _out71;
        _out71 = (_this).m4(globalState);
        _2186_v14 = _out71;
        (globalState).f13 = (_this).f30;
        let _2187_v15;
        _2187_v15 = _module.D0.create_DC1(_2174_v3);
        let _pat_let_tv23 = _2174_v3;
        let _2188_v16;
        _2188_v16 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let33_0) {
          return function (_2189_dt__update__tmp_h0) {
            return function (_pat_let34_0) {
              return function (_2190_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_2190_dt__update_hcf1_h0);
              }(_pat_let34_0);
            }(_pat_let_tv23);
          }(_pat_let33_0);
        }(_2187_v15),true);
        let _2191_v17;
        _2191_v17 = _module.D8.create_DC21(_2188_v16);
        let _2192_v18;
        _2192_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f38,_2188_v16);
        let _2193_v19;
        _2193_v19 = _dafny.Set.fromElements((_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))], _this.f29);
        let _2194_v21;
        _2194_v21 = _dafny.MultiSet.fromElements(_2187_v15);
        let _pat_let_tv24 = _2194_v21;
        let _pat_let_tv25 = _2194_v21;
        let _2195_v22;
        let _nw381 = Array((new BigNumber(18)).toNumber());
        _nw381[(_dafny.ZERO).toNumber()] = _2191_v17;
        _nw381[(_dafny.ONE).toNumber()] = function (_pat_let35_0) {
          return function (_2196_dt__update__tmp_h1) {
            return function (_pat_let36_0) {
              return function (_2198_dt__update_hcf40_h0) {
                return _module.D8.create_DC21(_2198_dt__update_hcf40_h0);
              }(_pat_let36_0);
            }(function () {
              let _coll58 = new _dafny.Map();
              for (const _compr_58 of (_pat_let_tv24).Elements) {
                let _2197_v20 = _compr_58;
                if ((_pat_let_tv25).contains(_2197_v20)) {
                  _coll58.push([_2197_v20,_this.f29]);
                }
              }
              return _coll58;
            }());
          }(_pat_let35_0);
        }(_module.D8.create_DC21((((_2192_v18).contains(new BigNumber((_2193_v19).length))) ? ((_2192_v18).get(new BigNumber((_2193_v19).length))) : (_dafny.Map.Empty.slice().updateUnsafe(_2187_v15,(_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))])))));
        _nw381[(new BigNumber(2)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(3)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(4)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(5)).toNumber()] = _module.__default.fm61(_2182_v11, (_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))], (_this).f30, globalState);
        _nw381[(new BigNumber(6)).toNumber()] = _module.D8.create_DC21(_2188_v16);
        _nw381[(new BigNumber(7)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(8)).toNumber()] = _module.D8.create_DC21(_2188_v16);
        _nw381[(new BigNumber(9)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(10)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(11)).toNumber()] = _module.D8.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_2187_v15,_this.f35));
        _nw381[(new BigNumber(12)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(13)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(14)).toNumber()] = _module.D8.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.MultiSet.fromElements(_this.f31, true)),_this.f29));
        _nw381[(new BigNumber(15)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(16)).toNumber()] = _2191_v17;
        _nw381[(new BigNumber(17)).toNumber()] = _2191_v17;
        _2195_v22 = _nw381;
        let _index320 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_2195_v22).length));
        (_2195_v22)[_index320] = _2191_v17;
        let _index321 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_2195_v22).length));
        let _rhs351 = _2191_v17;
        let _rhs352 = ((_2176_v5).Union(_dafny.MultiSet.FromArray(_2184_v13))).Difference((_2176_v5).Difference(_2176_v5));
        let _lhs289 = _2195_v22;
        let _lhs290 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_2195_v22).length));
        _lhs289[_lhs290] = _rhs351;
        _2176_v5 = _rhs352;
      }
      (globalState).f13 = (_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))];
      let _2199_i3;
      _2199_i3 = _dafny.ZERO;
      L7: {
        while ((new BigNumber(((_2174_v3).update(((_this).f32)[_module.__default.safeIndex((((_2174_v3).contains((_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))])) ? ((_2174_v3).get((_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))])) : ((_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))])), new BigNumber(((_this).f32).length))], _module.__default.abs((_this).f34))).cardinality())).isLessThan((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2169_v0,(_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))])).length)).multipliedBy((_this).f34))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2199_i3)) {
              break L7;
            }
            _2199_i3 = (_2199_i3).plus(_dafny.ONE);
            let _2200_v23;
            _2200_v23 = _dafny.Seq.UnicodeFromString("wejhofhx");
            let _2201_v24;
            _2201_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2200_v23,(_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))]);
            let _2202_v25;
            let _nw382 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _2202_v25 = _nw382;
            let _2203_v26;
            let _nw383 = new _module.C0();
            _nw383.__ctor(_2202_v25);
            _2203_v26 = _nw383;
            let _2204_v27;
            _2204_v27 = _module.D9.create_DC27(_2203_v26, _2200_v23);
            let _2205_v28;
            _2205_v28 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_2204_v27);
            let _2206_v29;
            _2206_v29 = _dafny.Map.Empty.slice().updateUnsafe((_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))],_2205_v28);
            let _rhs353 = _2201_v24;
            let _rhs354 = _module.__default.safeModuloInt(new BigNumber(((_2206_v29).Merge(_2206_v29)).length), (_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))]);
            let _rhs355 = _2169_v0;
            let _lhs291 = globalState;
            _2201_v24 = _rhs353;
            _lhs291.f13 = _rhs354;
            _2169_v0 = _rhs355;
            let _2207_v30;
            _2207_v30 = _dafny.Seq.of((_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))], (_this).f38);
            let _2208_v31;
            let _nw384 = Array((new BigNumber(16)).toNumber());
            _nw384[(_dafny.ZERO).toNumber()] = _2207_v30;
            _nw384[(_dafny.ONE).toNumber()] = _2207_v30;
            _nw384[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_2207_v30, _module.__default.safeIndex((_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))], new BigNumber((_2207_v30).length)), (_this).f38);
            _nw384[(new BigNumber(3)).toNumber()] = _2207_v30;
            _nw384[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(920)), ((_2209_v5) => function (_2210_i4) {
              return new BigNumber((_2209_v5).cardinality());
            })(_2176_v5));
            _nw384[(new BigNumber(5)).toNumber()] = _2207_v30;
            _nw384[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_2207_v30, _2207_v30);
            _nw384[(new BigNumber(7)).toNumber()] = _2207_v30;
            _nw384[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_2207_v30, _dafny.Seq.of((_this).f30, (_this).f30));
            _nw384[(new BigNumber(9)).toNumber()] = _module.__default.fm23(_this.f31, new BigNumber((_2174_v3).cardinality()), globalState);
            _nw384[(new BigNumber(10)).toNumber()] = _2207_v30;
            _nw384[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_2207_v30, _2207_v30), _module.__default.safeIndex((_this).f38, new BigNumber((_dafny.Seq.Concat(_2207_v30, _2207_v30)).length)), (_this).f34);
            _nw384[(new BigNumber(12)).toNumber()] = ((_this.f29) ? (_module.__default.fm23(!(_this.f35), (_this).f38, globalState)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(107)), function (_2211_i5) {
              return new BigNumber((_dafny.Seq.UnicodeFromString("idwc")).length);
            })));
            _nw384[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_2207_v30, _2207_v30);
            _nw384[(new BigNumber(14)).toNumber()] = _2207_v30;
            _nw384[(new BigNumber(15)).toNumber()] = _2207_v30;
            _2208_v31 = _nw384;
            let _index322 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length));
            let _rhs356 = _2208_v31;
            let _rhs357 = _2169_v0;
            let _rhs358 = _this.f31;
            let _rhs359 = (((_this).fm6(globalState)) ? (new BigNumber((_2200_v23).length)) : ((_dafny.ZERO).minus(((_this).f38).multipliedBy((_this).f34))));
            let _lhs292 = globalState;
            let _lhs293 = _2169_v0;
            let _lhs294 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length));
            _2208_v31 = _rhs356;
            _2169_v0 = _rhs357;
            _lhs292.f1 = _rhs358;
            _lhs293[_lhs294] = _rhs359;
            let _2212_v32;
            _2212_v32 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_2200_v23).length)),((_this).f34).plus((_this).f38));
            let _2213_v33;
            _2213_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2212_v32,(_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))]);
            _2212_v32 = (_2212_v32).update(new BigNumber(((_2213_v33).Merge(_2213_v33)).length), (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-265), (_this).f38)));
            let _2214_v34;
            _2214_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f31) || (_this.f31),_module.__default.fm62(_this.f36, (_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))], _this.f29, _this.f28, globalState));
            let _2215_v35;
            _2215_v35 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_this).f34);
            let _2216_v36;
            _2216_v36 = _dafny.Set.fromElements((_this).f38, (_this).f34);
            let _2217_v37;
            _2217_v37 = _module.D21.create_DC51(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2215_v35).length),_2216_v36));
            _2214_v34 = (_2214_v34).update(_this.f35, _2217_v37);
          }
        }
      }
      let _hi11 = (_this).f38;
      for (let _2218_i6 = (_this).f34; _2218_i6.isLessThan(_hi11); _2218_i6 = _2218_i6.plus(_dafny.ONE)) {
        let _2219_v38;
        _2219_v38 = _dafny.Seq.of(_2172_v2);
        let _2220_v39;
        _2220_v39 = _2219_v38;
        _2220_v39 = _2220_v39;
        let _index323 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length));
        (_2169_v0)[_index323] = (_dafny.ZERO).minus((new BigNumber(88)).plus(_2218_i6));
        let _2221_v40;
        _2221_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,_2172_v2);
        _2221_v40 = (_2221_v40).update(_this.f36, _2172_v2);
        if (_this.f35) {
          let _index324 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length));
          (_2169_v0)[_index324] = new BigNumber(55);
          (globalState).f21 = _2172_v2;
          let _2222_v41;
          _2222_v41 = _dafny.Seq.UnicodeFromString("fvq");
          let _2223_v42;
          let _nw385 = new _module.C4();
          _nw385.__ctor(new BigNumber(718), true, _dafny.Seq.Concat(_dafny.Seq.of(!(_this.f31)), (_this).f32), (_this).f33, new BigNumber((_2222_v41).length), (_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))], _this.f35, _this.f29);
          _2223_v42 = _nw385;
          let _2224_v43;
          _2224_v43 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f35);
          let _rhs360 = (_2172_v2)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length))];
          let _rhs361 = _2222_v41;
          let _rhs362 = _2169_v0;
          let _rhs363 = !((_2174_v3).IsSubsetOf(_dafny.MultiSet.fromElements((((_2224_v43).contains(_this.f28)) ? ((_2224_v43).get(_this.f28)) : (_this.f29)), _this.f31, false, _this.f29)));
          let _rhs364 = _2223_v42;
          let _lhs295 = _this;
          let _lhs296 = _this;
          _lhs295.f31 = _rhs360;
          _2222_v41 = _rhs361;
          _2169_v0 = _rhs362;
          _lhs296.f29 = _rhs363;
          _2223_v42 = _rhs364;
          let _2225_v44;
          _2225_v44 = _dafny.Seq.of(_2174_v3);
          let _index325 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_2172_v2).length));
          (_2172_v2)[_index325] = (((((_2225_v44)[_module.__default.safeIndex((_this).f38, new BigNumber((_2225_v44).length))]).update(_this.f31, _module.__default.abs((_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))]))).IsDisjointFrom(_2174_v3)) ? (_this.f35) : (true));
          let _2226_v45;
          let _init62 = ((_2227_v41) => function (_2228_i7) {
            return _2227_v41;
          })(_2222_v41);
          let _nw386 = Array((new BigNumber(8)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw386.length); _i0_62++) {
            _nw386[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _2226_v45 = _nw386;
          let _index326 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2226_v45).length));
          (_2226_v45)[_index326] = _2222_v41;
          let _index327 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2226_v45).length));
          (_2226_v45)[_index327] = _dafny.Seq.Concat(_dafny.Seq.Concat(_2222_v41, _2222_v41), (_this).fm3(_this.f28, _this.f31, globalState));
        } else {
          let _2229_v46;
          _2229_v46 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f31);
          (globalState).f13 = (_this).fm17(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_this.f28,_this.f28)).Merge((_2229_v46).update(_this.f28, true))).length), _2218_i6, globalState);
          let _2230_v47;
          _2230_v47 = _dafny.Seq.UnicodeFromString("gmac");
          _2230_v47 = _2230_v47;
          let _2231_v48;
          let _nw387 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2231_v48 = _nw387;
          let _2232_v49;
          let _nw388 = new _module.C0();
          _nw388.__ctor(_2231_v48);
          _2232_v49 = _nw388;
          _2232_v49 = _2232_v49;
          (globalState).f13 = _2218_i6;
          let _index328 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length));
          (_2169_v0)[_index328] = (_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))];
        }
      }
      let _2233_v50;
      _2233_v50 = _dafny.MultiSet.fromElements(_2174_v3, _2174_v3, _module.__default.fm33(_this.f29, globalState));
      r0 = new BigNumber(((_2233_v50).Union(((_2233_v50).update(_2174_v3, _module.__default.abs((_2169_v0)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_2169_v0).length))]))).update(_2174_v3, _module.__default.abs(new BigNumber(444))))).cardinality());
      let _nw389 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Set.Empty);
      r1 = _nw389;
      r2 = _2172_v2;
      return [r0, r1, r2];
    }
    get f38() {
      let _this = this;
      return _this._f38;
    };
    get f39() {
      let _this = this;
      return _this._f39;
    };
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f28 = false;
      this._f29 = false;
      this._f31 = false;
      this._f35 = false;
      this._f36 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f34 = _dafny.ZERO;
      this._f32 = _dafny.Seq.of();
      this._f33 = _dafny.Seq.of();
      this._f30 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T4, _module.T3, _module.T2, _module.T1, _module.T0];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    set f28(value) {
      let _this = this;
      _this._f28 = value;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f31() {
      let _this = this;
      return _this._f31;
    };
    set f31(value) {
      let _this = this;
      _this._f31 = value;
    };
    get f35() {
      let _this = this;
      return _this._f35;
    };
    set f35(value) {
      let _this = this;
      _this._f35 = value;
    };
    get f36() {
      let _this = this;
      return _this._f36;
    };
    set f36(value) {
      let _this = this;
      _this._f36 = value;
    };
    get f34() {
      let _this = this;
      return _this._f34;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
    __ctor(f36, f34, f35, f32, f33, f30, f31, f28, f29) {
      let _this = this;
      (_this)._f36 = f36;
      (_this)._f34 = f34;
      (_this)._f35 = f35;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f30 = f30;
      (_this)._f31 = f31;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm7(p0, p1, globalState) {
      let _this = this;
      return !(!((!(!(_this.f29))) || (_this.f35)));
    };
    fm6(globalState) {
      let _this = this;
      return (((_this.f35) ? (_module.D0.create_DC0(_this.f28)) : (_module.D0.create_DC0(_this.f31)))).dtor_cf0;
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements((_this).f34, (_this).f34)).Difference(function () {
        let _coll59 = new _dafny.Set();
        for (const _compr_59 of _dafny.IntegerRange(new BigNumber(507), new BigNumber(18))) {
          let _2234_v0 = _compr_59;
          if (((new BigNumber(507)).isLessThanOrEqualTo(_2234_v0)) && ((_2234_v0).isLessThan(new BigNumber(18)))) {
            _coll59.add((_2234_v0).minus((_this).f34));
          }
        }
        return _coll59;
      }());
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source34 = _module.D0.create_DC1(_dafny.MultiSet.FromArray((_this).f32));
      if (_source34.is_DC1) {
        let _2235___mcc_h0 = (_source34).cf1;
        let _2236_cf1 = _2235___mcc_h0;
        return _dafny.Seq.Concat((_this).f32, _dafny.Seq.of(_this.f28));
      } else {
        let _2237___mcc_h1 = (_source34).cf0;
        let _2238_cf0 = _2237___mcc_h1;
        return _dafny.Seq.Concat(((_this).f33)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f33).length))], (_this).f32);
      }
    };
    fm3(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lgpghnx"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(991)), function (_2239_i0) {
        return _this.f36;
      }));
    };
    fm2(globalState) {
      let _this = this;
      if (_this.f31) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_dafny.ZERO).minus((_this).f30))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_this).f30));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_dafny.ZERO).minus((_this).f34))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f35,new BigNumber(936)));
      }
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return (new BigNumber(361)).multipliedBy((new BigNumber(369)).multipliedBy((_this).f30));
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      (_this).f31 = _this.f28;
      let _2240_v0;
      _2240_v0 = _dafny.MultiSet.fromElements(_this.f29, _this.f31, _this.f29);
      let _2241_v1;
      _2241_v1 = _module.D0.create_DC1(_2240_v0);
      let _source35 = _2241_v1;
      if (_source35.is_DC1) {
        let _2242___mcc_h0 = (_source35).cf1;
        let _2243_cf1 = _2242___mcc_h0;
        let _2244_v2;
        let _nw390 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2244_v2 = _nw390;
        let _2245_v3;
        let _nw391 = new _module.C0();
        _nw391.__ctor(_2244_v2);
        _2245_v3 = _nw391;
        let _2246_v4;
        _2246_v4 = _dafny.Seq.UnicodeFromString("hdvneponb");
        let _2247_v5;
        _2247_v5 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f34));
        let _2248_v6;
        _2248_v6 = _dafny.MultiSet.fromElements((_this).f30);
        let _2249_v7;
        let _init63 = function (_2250_i0) {
          return _this.f28;
        };
        let _nw392 = Array((new BigNumber(19)).toNumber());
        for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw392.length); _i0_63++) {
          _nw392[_i0_63] = _init63(new BigNumber(_i0_63));
        }
        _2249_v7 = _nw392;
        let _2251_v8;
        _2251_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2249_v7,_2248_v6);
        let _2252_v9;
        let _nw393 = Array((new BigNumber(12)).toNumber());
        _nw393[(_dafny.ZERO).toNumber()] = new BigNumber(15);
        _nw393[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2246_v4, _2246_v4)).length);
        _nw393[(new BigNumber(2)).toNumber()] = new BigNumber((_2247_v5).length);
        _nw393[(new BigNumber(3)).toNumber()] = new BigNumber((_2248_v6).cardinality());
        _nw393[(new BigNumber(4)).toNumber()] = (_this).f34;
        _nw393[(new BigNumber(5)).toNumber()] = (_this).f34;
        _nw393[(new BigNumber(6)).toNumber()] = (_this).f30;
        _nw393[(new BigNumber(7)).toNumber()] = (_this).f30;
        _nw393[(new BigNumber(8)).toNumber()] = (_this).f30;
        _nw393[(new BigNumber(9)).toNumber()] = new BigNumber((_2251_v8).length);
        _nw393[(new BigNumber(10)).toNumber()] = (_this).f30;
        _nw393[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt((_this).f34, new BigNumber(590));
        _2252_v9 = _nw393;
        let _index329 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_2252_v9).length));
        (_2252_v9)[_index329] = (_this).f30;
        let _index330 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_2252_v9).length));
        (_2252_v9)[_index330] = (_dafny.ZERO).minus(((_this).fm8(false, _this.f36, _this.f29, globalState)).plus((_2247_v5)[_module.__default.safeIndex((_this).f34, new BigNumber((_2247_v5).length))]));
        let _index331 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_2252_v9).length));
        (_2252_v9)[_index331] = (_this).f30;
        let _2253_v10;
        _2253_v10 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_module.D1.create_DC2(_2249_v7)).dtor_cf2);
        let _2254_v11;
        _2254_v11 = _dafny.MultiSet.fromElements((((_2253_v10).contains(_this.f28)) ? ((_2253_v10).get(_this.f28)) : (_2249_v7)));
        let _2255_v12;
        _2255_v12 = _dafny.Set.fromElements((_2252_v9)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_2252_v9).length))]);
        let _2256_v14;
        let _out72;
        _out72 = _module.__default.m0(_dafny.Seq.Concat(_dafny.Seq.of((((_2254_v11).contains(_2249_v7)) ? ((_2254_v11).get(_2249_v7)) : ((_2252_v9)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_2252_v9).length))]))), _2247_v5), _this.f29, (((_2248_v6).contains(new BigNumber((_dafny.Seq.of(_this.f28)).length))) ? ((_2248_v6).get(new BigNumber((_dafny.Seq.of(_this.f28)).length))) : (new BigNumber((_2247_v5).length))), (function () {
          let _coll60 = new _dafny.Set();
          for (const _compr_60 of _dafny.IntegerRange(new BigNumber(580), new BigNumber(286))) {
            let _2257_v13 = _compr_60;
            if (((new BigNumber(580)).isLessThanOrEqualTo(_2257_v13)) && ((_2257_v13).isLessThan(new BigNumber(286)))) {
              _coll60.add((_2257_v13).multipliedBy(new BigNumber(-415)));
            }
          }
          return _coll60;
        }()).IsProperSubsetOf(_2255_v12), globalState);
        _2256_v14 = _out72;
      } else {
        let _2258___mcc_h1 = (_source35).cf0;
        let _2259_cf0 = _2258___mcc_h1;
        (globalState).f26 = _this.f35;
        let _2260_v15;
        let _init64 = function (_2261_i1) {
          return (_this).f32;
        };
        let _nw394 = Array((new BigNumber(28)).toNumber());
        for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw394.length); _i0_64++) {
          _nw394[_i0_64] = _init64(new BigNumber(_i0_64));
        }
        _2260_v15 = _nw394;
        let _2262_v16;
        _2262_v16 = _dafny.Seq.UnicodeFromString("erxeyj");
        let _2263_v17;
        _2263_v17 = _dafny.Seq.of((_this).f30, new BigNumber((_2262_v16).length));
        let _2264_v18;
        _2264_v18 = _dafny.Set.fromElements(new BigNumber((_2263_v17).length));
        let _rhs365 = _2260_v15;
        let _rhs366 = new BigNumber(((((_this).fm4(_this.f28, _2264_v18, (_this).f30, _2259_cf0, globalState)).Union(_2264_v18)).Union((_2264_v18).Difference(_2264_v18))).length);
        let _rhs367 = _2241_v1;
        let _rhs368 = ((_this).f30).plus((_this).f34);
        let _lhs297 = globalState;
        _2260_v15 = _rhs365;
        r1 = _rhs366;
        _2241_v1 = _rhs367;
        _lhs297.f13 = _rhs368;
        let _2265_v19;
        let _nw395 = Array((new BigNumber(24)).toNumber());
        _2265_v19 = _nw395;
        let _2266_v20;
        _2266_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-720),new BigNumber(212));
        let _2267_v21;
        _2267_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2265_v19,new BigNumber((_dafny.Seq.Concat(_module.__default.fm9(_this.f36, new BigNumber((_2266_v20).length), (_this).f34, _this.f35, globalState), _2263_v17)).length));
        _2267_v21 = (_2267_v21).update((_module.D2.create_DC5(_2265_v19)).dtor_cf8, new BigNumber(206));
        _2259_cf0 = _this.f31;
      }
      let _2268_v22;
      let _init65 = function (_2269_i2) {
        return _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(866)), function (_2270_i3) {
          return _this.f36;
        }), _dafny.Seq.UnicodeFromString("fxsini"));
      };
      let _nw396 = Array((new BigNumber(6)).toNumber());
      for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw396.length); _i0_65++) {
        _nw396[_i0_65] = _init65(new BigNumber(_i0_65));
      }
      _2268_v22 = _nw396;
      let _index332 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length));
      (_2268_v22)[_index332] = false;
      let _index333 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length));
      let _rhs369 = function (_source36) {
        if (_source36.is_DC1) {
          let _2271___mcc_h2 = (_source36).cf1;
          let _2272_cf1 = _2271___mcc_h2;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), function (_2273_i4) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          });
        } else {
          let _2274___mcc_h3 = (_source36).cf0;
          let _2275_cf0 = _2274___mcc_h3;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("srxfo"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(465)), function (_2276_i5) {
            return _this.f36;
          }));
        }
      }(_module.D0.create_DC0(_this.f29));
      let _rhs370 = _this.f31;
      let _lhs298 = _2268_v22;
      let _lhs299 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length));
      r2 = _rhs369;
      _lhs298[_lhs299] = _rhs370;
      let _2277_v23;
      _2277_v23 = _dafny.Seq.of(_module.__default.safeModuloInt((_this).f34, new BigNumber(106)), (_this).f34);
      let _2278_v24;
      _2278_v24 = _dafny.Seq.UnicodeFromString("p");
      let _rhs371 = _2277_v23;
      let _rhs372 = !(!(_this.f28));
      let _rhs373 = _2278_v24;
      _2277_v23 = _rhs371;
      r0 = _rhs372;
      r2 = _rhs373;
      let _2279_i6;
      _2279_i6 = _dafny.ZERO;
      L8: {
        while (!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(241)), function (_2315_i7) {
          return _this.f36;
        }), (_this).fm3((_2268_v22)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length))], (_2268_v22)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length))], globalState)))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2279_i6)) {
              break L8;
            }
            _2279_i6 = (_2279_i6).plus(_dafny.ONE);
            let _2280_v25;
            _2280_v25 = _dafny.Set.fromElements(false, _this.f35, _this.f29);
            _2280_v25 = _dafny.Set.fromElements(_this.f35, ((_this).f34).isLessThan((_this).f30), (_2268_v22)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length))], !(_2280_v25).contains(_this.f28));
            let _2281_v26;
            _2281_v26 = _module.D0.create_DC0(true);
            let _2282_v27;
            let _nw397 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2282_v27 = _nw397;
            let _2283_v28;
            _2283_v28 = _dafny.Map.Empty.slice().updateUnsafe(false,(((_2281_v26).dtor_cf0) ? (_2282_v27) : (_2282_v27)));
            _2283_v28 = (_2283_v28).update(_this.f28, (((_2268_v22)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length))]) ? (_2282_v27) : ((((_2283_v28).contains(_this.f35)) ? ((_2283_v28).get(_this.f35)) : (_2282_v27)))));
            (globalState).f21 = _2268_v22;
            let _2284_v29;
            let _nw398 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _2284_v29 = _nw398;
            let _2285_v30;
            let _nw399 = new _module.C0();
            _nw399.__ctor(_2284_v29);
            _2285_v30 = _nw399;
            let _2286_v31;
            _2286_v31 = _module.D1.create_DC3(_this.f35, ((_this).f30).multipliedBy((_this).f34), _2285_v30, new BigNumber(929));
            let _source37 = _2286_v31;
            if (_source37.is_DC3) {
              let _2287___mcc_h4 = (_source37).cf3;
              let _2288___mcc_h5 = (_source37).cf4;
              let _2289___mcc_h6 = (_source37).cf5;
              let _2290___mcc_h7 = (_source37).cf6;
              let _2291_cf6 = _2290___mcc_h7;
              let _2292_cf5 = _2289___mcc_h6;
              let _2293_cf4 = _2288___mcc_h5;
              let _2294_cf3 = _2287___mcc_h4;
              let _2295_v32;
              _2295_v32 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f30).multipliedBy((_dafny.ZERO).minus(_2291_cf6)),(_this).f34);
              let _2296_v33;
              _2296_v33 = _module.D2.create_DC6((_this).f34, (_2268_v22)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length))], new BigNumber((_2278_v24).length), _this.f29, (_this).f30);
              _2295_v32 = (_2295_v32).update(_2291_cf6, (_2296_v33).dtor_cf9);
              (globalState).f1 = (_dafny.Set.fromElements(_module.__default.fm11((_this).f33, _this.f35, true, (_this).f34, globalState), _2280_v25, _2280_v25)).IsProperSubsetOf(_module.__default.fm10(_2280_v25, new _dafny.CodePoint('y'.codePointAt(0)), _2293_cf4, !(_this.f31), globalState));
              r2 = _dafny.Seq.update(_dafny.Seq.Concat(_2278_v24, _2278_v24), _module.__default.safeIndex(_2293_cf4, new BigNumber((_dafny.Seq.Concat(_2278_v24, _2278_v24)).length)), new _dafny.CodePoint('r'.codePointAt(0)));
              (globalState).f12 = _this.f31;
            } else if (_source37.is_DC2) {
              let _2297___mcc_h8 = (_source37).cf2;
              let _2298_cf2 = _2297___mcc_h8;
              let _2299_v34;
              _2299_v34 = _module.D2.create_DC6((_this).f30, _this.f35, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(((_this).f32).length), (_this).f30),(_this).f34)).length), !(_module.__default.fm1(true, (_this).f30, globalState)), (_this).f34);
              let _2300_v35;
              _2300_v35 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f30);
              _2299_v34 = _module.__default.fm12(_this.f36, (_dafny.ZERO).minus(((true) ? ((_this).f30) : ((_this).f30))), _2300_v35, _dafny.Seq.IsProperPrefixOf(_2278_v24, _2278_v24), globalState);
              let _2301_v36;
              _2301_v36 = _dafny.Seq.of(_2300_v35);
              let _2302_v37;
              _2302_v37 = _module.D3.create_DC7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-446)), function (_2303_i8) {
  return _dafny.Map.Empty.slice().updateUnsafe(_this.f28,(_dafny.ZERO).minus((_this).f34));
}));
              let _2304_v38;
              let _nw400 = Array((new BigNumber(6)).toNumber());
              _nw400[(_dafny.ZERO).toNumber()] = _2301_v36;
              _nw400[(_dafny.ONE).toNumber()] = (_2302_v37).dtor_cf14;
              _nw400[(new BigNumber(2)).toNumber()] = _2301_v36;
              _nw400[(new BigNumber(3)).toNumber()] = _2301_v36;
              _nw400[(new BigNumber(4)).toNumber()] = _2301_v36;
              _nw400[(new BigNumber(5)).toNumber()] = _2301_v36;
              _2304_v38 = _nw400;
              let _index334 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_2304_v38).length));
              (_2304_v38)[_index334] = _2301_v36;
              let _index335 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length));
              let _index336 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_2304_v38).length));
              let _rhs374 = ((_this).f30).isLessThan((_this).fm8(_this.f31, new _dafny.CodePoint('p'.codePointAt(0)), _this.f31, globalState));
              let _rhs375 = (_2268_v22)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length))];
              let _rhs376 = new _dafny.CodePoint('d'.codePointAt(0));
              let _rhs377 = (_this).f30;
              let _rhs378 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2301_v36, _2301_v36), _2301_v36);
              let _lhs300 = _2268_v22;
              let _lhs301 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length));
              let _lhs302 = _this;
              let _lhs303 = globalState;
              let _lhs304 = _2304_v38;
              let _lhs305 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_2304_v38).length));
              r0 = _rhs374;
              _lhs300[_lhs301] = _rhs375;
              _lhs302.f36 = _rhs376;
              _lhs303.f13 = _rhs377;
              _lhs304[_lhs305] = _rhs378;
              _2300_v35 = (_2300_v35).update((_2268_v22)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_2268_v22).length))], (_this).f34);
              _2285_v30 = _2285_v30;
            } else {
              let _2305___mcc_h9 = (_source37).cf7;
              let _2306_cf7 = _2305___mcc_h9;
              (globalState).f26 = _this.f31;
              let _2307_v39;
              _2307_v39 = _dafny.MultiSet.fromElements(_this.f36, _this.f36, _this.f36);
              let _2308_v40;
              _2308_v40 = _module.D3.create_DC8((_2307_v39).Union(_dafny.MultiSet.fromElements(_this.f36, _this.f36, _this.f36)), true, (_this).f30);
              let _2309_v41;
              _2309_v41 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f33).length),(_this).f30),(_this).f34);
              let _2310_v43;
              _2310_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f30);
              let _2311_v44;
              _2311_v44 = _dafny.Set.fromElements(_2310_v43);
              let _2312_v46;
              _2312_v46 = _dafny.MultiSet.fromElements(_2280_v25);
              _2308_v40 = _module.__default.fm13(new BigNumber(((_2309_v41).Merge(function () {
                let _coll61 = new _dafny.Map();
                for (const _compr_61 of (_2311_v44).Elements) {
                  let _2313_v42 = _compr_61;
                  if ((_2311_v44).contains(_2313_v42)) {
                    _coll61.push([_2313_v42,new BigNumber((function () {
                      let _coll62 = new _dafny.Set();
                      for (const _compr_62 of _dafny.IntegerRange(new BigNumber(590), new BigNumber(316))) {
                        let _2314_v45 = _compr_62;
                        if (((new BigNumber(590)).isLessThanOrEqualTo(_2314_v45)) && ((_2314_v45).isLessThan(new BigNumber(316)))) {
                          _coll62.add((_2314_v45).plus(new BigNumber((_2277_v23).length)));
                        }
                      }
                      return _coll62;
                    }()).length)]);
                  }
                }
                return _coll61;
              }())).length), (_this).f34, (_this).f34, (_2312_v46).Intersect(_2312_v46), globalState);
              (globalState).f13 = (_2277_v23)[_module.__default.safeIndex((_this).f34, new BigNumber((_2277_v23).length))];
              (globalState).f1 = false;
            }
          }
        }
      }
      let _2316_i9;
      _2316_i9 = _dafny.ZERO;
      L9: {
        while (_this.f31) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2316_i9)) {
              break L9;
            }
            _2316_i9 = (_2316_i9).plus(_dafny.ONE);
            (_this).f31 = !((_this).fm7(_dafny.MultiSet.fromElements((_this).f34), (_this).f30, globalState));
            r1 = ((_this).f30).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),_this.f29)).length)));
            let _2317_v47;
            _2317_v47 = _dafny.MultiSet.fromElements(_module.__default.fm14((_this).f30, false, new BigNumber(427), _module.__default.fm15((_this).f34, (_this).f30, globalState), globalState), _this.f36);
            let _2318_v48;
            _2318_v48 = _dafny.Set.fromElements(_2317_v47);
            _2318_v48 = _dafny.Set.fromElements(_2317_v47);
            let _2319_v49;
            let _nw401 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _2319_v49 = _nw401;
            let _2320_v50;
            let _nw402 = new _module.C0();
            _nw402.__ctor(_2319_v49);
            _2320_v50 = _nw402;
          }
        }
      }
      r0 = _this.f31;
      r1 = (_this).f34;
      r2 = _dafny.Seq.Concat((_this).fm3(((_this).f32)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length))], !(true), globalState), _dafny.Seq.Concat(_2278_v24, _2278_v24));
      return [r0, r1, r2];
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _module.D0.Default();
      let r2 = _module.D0.Default();
      let r3 = false;
      let _2321_v0;
      let _init66 = function (_2322_i0) {
        return _this.f31;
      };
      let _nw403 = Array((new BigNumber(16)).toNumber());
      for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw403.length); _i0_66++) {
        _nw403[_i0_66] = _init66(new BigNumber(_i0_66));
      }
      _2321_v0 = _nw403;
      let _2323_v1;
      _2323_v1 = _dafny.MultiSet.fromElements(_this.f36, _this.f36, _this.f36, _this.f36);
      let _2324_v2;
      _2324_v2 = _module.D3.create_DC8(_2323_v1, _this.f28, (_this).f34);
      let _index337 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length));
      (_2321_v0)[_index337] = ((_this).f30).isEqualTo((_2324_v2).dtor_cf17);
      let _index338 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length));
      (_2321_v0)[_index338] = _this.f29;
      let _2325_i1;
      _2325_i1 = _dafny.ZERO;
      L10: {
        while (_this.f35) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2325_i1)) {
              break L10;
            }
            _2325_i1 = (_2325_i1).plus(_dafny.ONE);
            let _2326_v3;
            _2326_v3 = _module.D3.create_DC9(!(_this.f28), _this.f35, _2321_v0, _this.f28);
            let _2327_v4;
            _2327_v4 = _dafny.Seq.UnicodeFromString("kddokf");
            let _index339 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length));
            let _rhs379 = _module.D3.create_DC9(_this.f28, !_dafny.Seq.contains(_dafny.Seq.of(false, (_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))], _this.f28, (_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))], _this.f35), !(_this.f35)), _2321_v0, (_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))]);
            let _rhs380 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_2327_v4, _2327_v4)).length), p0);
            let _rhs381 = !(_this.f28) || (_this.f28);
            let _rhs382 = _2326_v3;
            let _rhs383 = _this.f31;
            let _lhs306 = globalState;
            let _lhs307 = _2321_v0;
            let _lhs308 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length));
            let _lhs309 = _this;
            _2326_v3 = _rhs379;
            _lhs306.f13 = _rhs380;
            _lhs307[_lhs308] = _rhs381;
            _2326_v3 = _rhs382;
            _lhs309.f31 = _rhs383;
            let _2328_v5;
            _2328_v5 = _dafny.MultiSet.fromElements(p2);
            let _2329_v6;
            _2329_v6 = _dafny.Map.Empty.slice().updateUnsafe((_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))],new BigNumber(((_this).f32).length));
            (globalState).f26 = !((_2328_v5).IsProperSubsetOf(_2328_v5)) || (!(_2329_v6).equals(_2329_v6));
            let _2330_v7;
            let _init67 = ((_2331_p0) => function (_2332_i2) {
              return _module.__default.safeModuloInt(_2332_i2, _2331_p0);
            })(p0);
            let _nw404 = Array((new BigNumber(10)).toNumber());
            for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw404.length); _i0_67++) {
              _nw404[_i0_67] = _init67(new BigNumber(_i0_67));
            }
            _2330_v7 = _nw404;
            _2330_v7 = _2330_v7;
            (globalState).f1 = !(((_this.f28) ? (_this.f31) : ((_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))])));
          }
        }
      }
      let _2333_v8;
      let _nw405 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2333_v8 = _nw405;
      let _2334_v9;
      _2334_v9 = _dafny.Seq.of(_2333_v8);
      let _2335_v10;
      _2335_v10 = _dafny.Set.fromElements((_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))]);
      let _2336_v11;
      _2336_v11 = _dafny.MultiSet.fromElements(_2335_v10);
      let _2337_v12;
      let _nw406 = Array((new BigNumber(28)).toNumber());
      _nw406[(_dafny.ZERO).toNumber()] = (_2334_v9)[_module.__default.safeIndex((((_2336_v11).contains(_dafny.Set.fromElements((_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))]))) ? ((_2336_v11).get(_dafny.Set.fromElements((_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))]))) : ((_this).f34)), new BigNumber((_2334_v9).length))];
      _nw406[(_dafny.ONE).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(2)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(3)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(4)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(5)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(6)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(7)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(8)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(9)).toNumber()] = ((_this.f31) ? (_2333_v8) : (_2333_v8));
      _nw406[(new BigNumber(10)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(11)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(12)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(13)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(14)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(15)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(16)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(17)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(18)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(19)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(20)).toNumber()] = ((_this.f28) ? (_2333_v8) : (_2333_v8));
      _nw406[(new BigNumber(21)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(22)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(23)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(24)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(25)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(26)).toNumber()] = _2333_v8;
      _nw406[(new BigNumber(27)).toNumber()] = _2333_v8;
      _2337_v12 = _nw406;
      let _index340 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_2337_v12).length));
      (_2337_v12)[_index340] = _2333_v8;
      let _index341 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_2337_v12).length));
      (_2337_v12)[_index341] = _2333_v8;
      let _2338_i3;
      _2338_i3 = _dafny.ZERO;
      L11: {
        while (true) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2338_i3)) {
              break L11;
            }
            _2338_i3 = (_2338_i3).plus(_dafny.ONE);
            let _index342 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length));
            (_2321_v0)[_index342] = (((_this).f34).minus(p1)).isLessThanOrEqualTo((_this).f34);
            let _2339_v13;
            _2339_v13 = _dafny.MultiSet.fromElements(_this.f35, _this.f28);
            let _2340_v14;
            _2340_v14 = _dafny.Seq.of(_2339_v13);
            let _2341_v15;
            let _nw407 = Array((new BigNumber(14)).toNumber());
            _nw407[(_dafny.ZERO).toNumber()] = _2339_v13;
            _nw407[(_dafny.ONE).toNumber()] = ((_this.f28) ? (_2339_v13) : (_2339_v13));
            _nw407[(new BigNumber(2)).toNumber()] = _2339_v13;
            _nw407[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_this.f31, ((_this).f32)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length))], _this.f29, _this.f28, _this.f28);
            _nw407[(new BigNumber(4)).toNumber()] = _2339_v13;
            _nw407[(new BigNumber(5)).toNumber()] = _2339_v13;
            _nw407[(new BigNumber(6)).toNumber()] = _2339_v13;
            _nw407[(new BigNumber(7)).toNumber()] = ((_this.f31) ? (_2339_v13) : (_2339_v13));
            _nw407[(new BigNumber(8)).toNumber()] = (_2340_v14)[_module.__default.safeIndex(p2, new BigNumber((_2340_v14).length))];
            _nw407[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(((_this).f32)[_module.__default.safeIndex((_this).f30, new BigNumber(((_this).f32).length))]);
            _nw407[(new BigNumber(10)).toNumber()] = (_2339_v13).Union(_2339_v13);
            _nw407[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_this.f35, _this.f35, (_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))], _this.f28, false);
            _nw407[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(!(_this.f31)));
            _nw407[(new BigNumber(13)).toNumber()] = _2339_v13;
            _2341_v15 = _nw407;
            let _index343 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_2341_v15).length));
            (_2341_v15)[_index343] = (_2339_v13).update(false, _module.__default.abs((_dafny.ZERO).minus((_this).f30)));
            let _2342_v16;
            _2342_v16 = _dafny.Seq.UnicodeFromString("unmvjnend");
            let _2343_v17;
            _2343_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_2342_v16);
            let _index344 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_2341_v15).length));
            (_2341_v15)[_index344] = _dafny.MultiSet.fromElements(_this.f29, _dafny.areEqual(_dafny.Seq.UnicodeFromString("jfc"), (((_2343_v17).contains((_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))])) ? ((_2343_v17).get((_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))])) : (_2342_v16))));
            let _2344_v18;
            let _nw408 = Array((new BigNumber(6)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _2344_v18 = _nw408;
            let _2345_v19;
            _2345_v19 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_2344_v18);
            let _2346_v20;
            let _nw409 = new _module.C0();
            _nw409.__ctor((((_2345_v19).contains(_this.f29)) ? ((_2345_v19).get(_this.f29)) : (_2344_v18)));
            _2346_v20 = _nw409;
            let _2347_v21;
            let _init68 = function (_2348_i4) {
              return (_this).f32;
            };
            let _nw410 = Array((new BigNumber(13)).toNumber());
            for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw410.length); _i0_68++) {
              _nw410[_i0_68] = _init68(new BigNumber(_i0_68));
            }
            _2347_v21 = _nw410;
            let _2349_v22;
            _2349_v22 = _module.D1.create_DC3((_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))], p0, _2346_v20, (_this).f30);
            let _2350_v23;
            let _nw411 = new _module.C3();
            _nw411.__ctor(new BigNumber(((_this).f32).length), _2347_v21, _this.f36, (_dafny.ZERO).minus((_this).f30), (_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))], _dafny.Seq.of(_this.f31, (_2349_v22).dtor_cf3), _dafny.Seq.of(_dafny.Seq.of(_this.f31)), new BigNumber(((_this).f32).length), _this.f35, _this.f28, !(_this.f29));
            _2350_v23 = _nw411;
            let _2351_v24;
            _2351_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2350_v23,p1);
            let _2352_v25;
            _2352_v25 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f28) ? (_this.f36) : (_this.f36)),(((_2351_v24).contains(_2350_v23)) ? ((_2351_v24).get(_2350_v23)) : ((_dafny.ZERO).minus(p2))));
            let _2353_v26;
            _2353_v26 = _dafny.Seq.of((_this).f30);
            let _2354_v27;
            _2354_v27 = _dafny.MultiSet.fromElements((_2353_v26)[_module.__default.safeIndex((_this).f30, new BigNumber((_2353_v26).length))], (_2350_v23).f34, (_2350_v23).f30);
            _2352_v25 = (_2352_v25).update(_module.__default.fm34((_this).f34, p0, (((_2354_v27).contains(_module.__default.fm46((_2350_v23).f34, p1, _this.f31, globalState))) ? ((_2354_v27).get(_module.__default.fm46((_2350_v23).f34, p1, _this.f31, globalState))) : (p2)), false, globalState), (_2350_v23).f34);
          }
        }
      }
      let _2355_v28;
      let _nw412 = Array((new BigNumber(18)).toNumber()).fill([]);
      _2355_v28 = _nw412;
      let _2356_v29;
      let _nw413 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Set.Empty);
      _2356_v29 = _nw413;
      let _2357_v30;
      _2357_v30 = _module.D22.create_DC54(_2356_v29);
      let _index345 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2355_v28).length));
      (_2355_v28)[_index345] = (_2357_v30).dtor_cf91;
      let _index346 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2355_v28).length));
      (_2355_v28)[_index346] = _2356_v29;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2321_v0).length))) {
        let _2358_i5 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2358_i5)) && ((_2358_i5).isLessThan(new BigNumber((_2321_v0).length))))) {
          (_2321_v0)[(_2358_i5)] = false;
        }
      }
      let _2359_v31;
      _2359_v31 = _dafny.Seq.UnicodeFromString("yciqas");
      let _2360_v32;
      _2360_v32 = _module.D16.create_DC40(_2359_v31, p1);
      let _2361_v33;
      _2361_v33 = _module.D16.create_DC41(_2360_v32);
      let _2362_v34;
      _2362_v34 = _dafny.MultiSet.fromElements(_2361_v33, _2361_v33);
      r0 = ((_dafny.MultiSet.fromElements(_2361_v33)).update(_2361_v33, _module.__default.abs((_this).f30))).IsProperSubsetOf((_2362_v34).update(_2361_v33, _module.__default.abs(p1)));
      let _2363_v35;
      _2363_v35 = _dafny.MultiSet.fromElements(_this.f28, _this.f28, _this.f28, (_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))], true);
      r1 = _module.D0.create_DC1(_2363_v35);
      let _2364_v36;
      _2364_v36 = _module.D0.create_DC0(!(!(_2335_v10).equals(_dafny.Set.fromElements((_2321_v0)[_module.__default.safeIndex(new BigNumber(695), new BigNumber((_2321_v0).length))], _this.f31, _this.f28))));
      r2 = _2364_v36;
      r3 = _this.f28;
      return [r0, r1, r2, r3];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _2365_v0;
      _2365_v0 = _dafny.Set.fromElements(new BigNumber(-512), (_this).f30);
      let _2366_v1;
      _2366_v1 = _module.D7.create_DC19(p0, new BigNumber(608), (_this).f30);
      let _2367_v2;
      let _nw414 = Array((new BigNumber(8)).toNumber());
      _nw414[(_dafny.ZERO).toNumber()] = (_this).f34;
      _nw414[(_dafny.ONE).toNumber()] = (_this).f30;
      _nw414[(new BigNumber(2)).toNumber()] = (_2366_v1).dtor_cf38;
      _nw414[(new BigNumber(3)).toNumber()] = (_this).f30;
      _nw414[(new BigNumber(4)).toNumber()] = (_this).f34;
      _nw414[(new BigNumber(5)).toNumber()] = ((_module.__default.fm1(p0, (_this).f30, globalState)) ? ((_this).f30) : ((_this).f34));
      _nw414[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_this).f30);
      _nw414[(new BigNumber(7)).toNumber()] = new BigNumber(135);
      _2367_v2 = _nw414;
      let _index347 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
      (_2367_v2)[_index347] = (_this).f30;
      let _2368_v3;
      _2368_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,new BigNumber(((_this).f32).length));
      let _2369_v4;
      _2369_v4 = _module.D5.create_DC13(p0, _2368_v3, _this.f36);
      let _2370_v5;
      _2370_v5 = _dafny.MultiSet.fromElements(_2368_v3, _2368_v3);
      let _2371_v6;
      _2371_v6 = _dafny.Seq.UnicodeFromString("vilf");
      let _2372_v7;
      _2372_v7 = _dafny.MultiSet.fromElements(new BigNumber((_2371_v6).length));
      let _2373_v8;
      _2373_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,p1);
      let _2374_v9;
      _2374_v9 = _dafny.Seq.of((_this).f34, new BigNumber((_2373_v8).length));
      let _2375_v10;
      _2375_v10 = _dafny.MultiSet.fromElements(p1);
      let _pat_let_tv26 = _2365_v0;
      let _pat_let_tv27 = _2365_v0;
      let _pat_let_tv28 = _2365_v0;
      let _index348 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
      let _rhs384 = p0;
      let _rhs385 = function (_source38) {
        if (_source38.is_DC13) {
          let _2376___mcc_h0 = (_source38).cf27;
          let _2377___mcc_h1 = (_source38).cf28;
          let _2378___mcc_h2 = (_source38).cf29;
          let _2379_cf29 = _2378___mcc_h2;
          let _2380_cf28 = _2377___mcc_h1;
          let _2381_cf27 = _2376___mcc_h0;
          return _dafny.Set.fromElements((_this).f30, new BigNumber(-653));
        } else if (_source38.is_DC14) {
          let _2382___mcc_h3 = (_source38).cf30;
          let _2383___mcc_h4 = (_source38).cf31;
          let _2384_cf31 = _2383___mcc_h4;
          let _2385_cf30 = _2382___mcc_h3;
          return _pat_let_tv26;
        } else {
          let _2386___mcc_h5 = (_source38).cf26;
          let _2387_cf26 = _2386___mcc_h5;
          if (_this.f35) {
            return _pat_let_tv27;
          } else {
            return _pat_let_tv28;
          }
        }
      }(_2369_v4);
      let _rhs386 = ((_dafny.ZERO).minus(new BigNumber(((_2370_v5).update((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f30)).update(false, (_this).f34), _module.__default.abs(new BigNumber((_dafny.Seq.of((_this).fm7(_module.__default.fm49(_2372_v7, _this.f36, _dafny.Seq.of(_2374_v9), globalState), (_this).f30, globalState))).length)))).cardinality()))).plus(new BigNumber((_2375_v10).cardinality()));
      let _lhs310 = _this;
      let _lhs311 = _2367_v2;
      let _lhs312 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
      _lhs310.f31 = _rhs384;
      _2365_v0 = _rhs385;
      _lhs311[_lhs312] = _rhs386;
      if (!(p0)) {
        let _2388_v11;
        _2388_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f35,_2374_v9);
        let _2389_v12;
        _2389_v12 = _dafny.Seq.of((new BigNumber((_dafny.Seq.update(_2374_v9, _module.__default.safeIndex((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], new BigNumber((_2374_v9).length)), new BigNumber((_2371_v6).length))).length)).isLessThan(new BigNumber((_2388_v11).length)), _this.f29);
        _2389_v12 = _module.__default.fm0(((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]).minus((_this).f34), _this.f31, p0, ((_this).f34).plus(new BigNumber(((_this).fm3(_this.f29, _this.f29, globalState)).length)), globalState);
        (_this).f29 = _this.f35;
        let _2390_v13;
        _2390_v13 = _dafny.Seq.of(_2365_v0);
        _2365_v0 = ((_2390_v13)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_this.f28, _this.f28)).cardinality()), new BigNumber((_2390_v13).length))]).Intersect((_2365_v0).Union(_2365_v0));
        let _2391_v14;
        _2391_v14 = _dafny.MultiSet.fromElements((_2375_v10).update(p0, _module.__default.abs(new BigNumber(178))));
        let _2392_v15;
        _2392_v15 = _2391_v14;
        let _source39 = _2392_v15;
        let _2393___mcc_h6 = _source39;
        let _2394_cf80 = _2393___mcc_h6;
        let _2395_v16;
        let _nw415 = Array((new BigNumber(22)).toNumber()).fill(_dafny.MultiSet.Empty);
        _2395_v16 = _nw415;
        let _index349 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_2395_v16).length));
        (_2395_v16)[_index349] = _2375_v10;
        let _2396_v17;
        _2396_v17 = _dafny.Seq.of(_2375_v10, _2375_v10);
        let _index350 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_2395_v16).length));
        (_2395_v16)[_index350] = (_2396_v17)[_module.__default.safeIndex(((_this).f30).minus((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]), new BigNumber((_2396_v17).length))];
        (_this).f36 = ((_this.f29) ? (_this.f36) : (_this.f36));
        let _2397_v18;
        let _nw416 = Array((new BigNumber(21)).toNumber()).fill(_module.D7.Default());
        _2397_v18 = _nw416;
        let _2398_v19;
        _2398_v19 = _2397_v18;
        _2398_v19 = _2398_v19;
        (globalState).f12 = p1;
        let _2399_v20;
        let _init69 = function (_2400_i0) {
          return _this.f29;
        };
        let _nw417 = Array((new BigNumber(11)).toNumber());
        for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw417.length); _i0_69++) {
          _nw417[_i0_69] = _init69(new BigNumber(_i0_69));
        }
        _2399_v20 = _nw417;
        let _2401_v21;
        _2401_v21 = _module.D6.create_DC16((_module.D3.create_DC9(_module.__default.fm1(p0, (_this).f30, globalState), _module.__default.fm1(_this.f35, (_this).f30, globalState), _2399_v20, false)).dtor_cf20);
        let _2402_v22;
        _2402_v22 = _module.D20.create_DC50(_module.D20.create_DC49(_this.f35, _this.f36));
        let _2403_v23;
        _2403_v23 = _module.D3.create_DC9(_module.__default.fm1(p1, (_this).f34, globalState), false, _2399_v20, p0);
        let _2404_v24;
        _2404_v24 = _dafny.Seq.of((_2403_v23).dtor_cf20);
        let _2405_v25;
        let _nw418 = new _module.C1();
        _nw418.__ctor(_this.f28, p1, p1);
        _2405_v25 = _nw418;
        let _2406_v26;
        _2406_v26 = _module.D10.create_DC28(_2405_v25);
        let _pat_let_tv29 = _2405_v25;
        let _pat_let_tv30 = _2405_v25;
        let _2407_v27;
        let _nw419 = Array((new BigNumber(21)).toNumber());
        _nw419[(_dafny.ZERO).toNumber()] = _2406_v26;
        _nw419[(_dafny.ONE).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(2)).toNumber()] = _module.D10.create_DC28(_2405_v25);
        _nw419[(new BigNumber(3)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(4)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(5)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(6)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(7)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(8)).toNumber()] = _module.D10.create_DC28(_2405_v25);
        _nw419[(new BigNumber(9)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(10)).toNumber()] = _module.D10.create_DC28(_2405_v25);
        _nw419[(new BigNumber(11)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(12)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(13)).toNumber()] = function (_pat_let37_0) {
          return function (_2408_dt__update__tmp_h0) {
            return function (_pat_let38_0) {
              return function (_2409_dt__update_hcf55_h0) {
                return _module.D10.create_DC28(_2409_dt__update_hcf55_h0);
              }(_pat_let38_0);
            }(_pat_let_tv29);
          }(_pat_let37_0);
        }(_2406_v26);
        _nw419[(new BigNumber(14)).toNumber()] = function (_pat_let39_0) {
          return function (_2410_dt__update__tmp_h1) {
            return function (_pat_let40_0) {
              return function (_2411_dt__update_hcf55_h1) {
                return _module.D10.create_DC28(_2411_dt__update_hcf55_h1);
              }(_pat_let40_0);
            }(_pat_let_tv30);
          }(_pat_let39_0);
        }(_2406_v26);
        _nw419[(new BigNumber(15)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(16)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(17)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(18)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(19)).toNumber()] = _2406_v26;
        _nw419[(new BigNumber(20)).toNumber()] = _module.D10.create_DC28(_2405_v25);
        _2407_v27 = _nw419;
        let _2412_v28;
        _2412_v28 = _module.D20.create_DC48((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], _2407_v27, _this.f28);
        let _2413_v29;
        _2413_v29 = _module.D20.create_DC50(_2412_v28);
        let _pat_let_tv31 = _2404_v24;
        let _pat_let_tv32 = _2404_v24;
        let _rhs387 = function (_pat_let41_0) {
          return function (_2414_dt__update__tmp_h2) {
            return function (_pat_let42_0) {
              return function (_2415_dt__update_hcf33_h0) {
                return _module.D6.create_DC16(_2415_dt__update_hcf33_h0);
              }(_pat_let42_0);
            }((_pat_let_tv31)[_module.__default.safeIndex((_this).f34, new BigNumber((_pat_let_tv32).length))]);
          }(_pat_let41_0);
        }(((p0) ? (_2401_v21) : (_2401_v21)));
        let _rhs388 = _dafny.Seq.contains(_2371_v6, _this.f36);
        let _rhs389 = ((true) ? (_module.D20.create_DC50(_2413_v29)) : (_2402_v22));
        let _rhs390 = _module.__default.safeModuloInt((_this).f34, new BigNumber(-651));
        let _lhs313 = globalState;
        _2401_v21 = _rhs387;
        _lhs313.f26 = _rhs388;
        _2402_v22 = _rhs389;
        r1 = _rhs390;
      } else {
        let _index351 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
        let _rhs391 = _2371_v6;
        let _rhs392 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]));
        let _lhs314 = _2367_v2;
        let _lhs315 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
        _2371_v6 = _rhs391;
        _lhs314[_lhs315] = _rhs392;
        let _2416_v30;
        _2416_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,new BigNumber(717));
        let _2417_v31;
        _2417_v31 = _dafny.Map.Empty.slice().updateUnsafe(_2416_v30,_2371_v6);
        (globalState).f1 = (_module.__default.safeModuloInt(new BigNumber((_2417_v31).length), (_this).f34)).isLessThanOrEqualTo((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]);
        if (_this.f35) {
          _2374_v9 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2374_v9, _2374_v9), _dafny.Seq.Concat(_2374_v9, _2374_v9));
          let _2418_v32;
          let _init70 = function (_2419_i1) {
            return true;
          };
          let _nw420 = Array((new BigNumber(6)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw420.length); _i0_70++) {
            _nw420[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2418_v32 = _nw420;
          let _2420_v33;
          _2420_v33 = _dafny.Map.Empty.slice().updateUnsafe(_this.f36,_2373_v8);
          let _index352 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_2418_v32).length));
          (_2418_v32)[_index352] = _dafny.Seq.IsProperPrefixOf(_module.__default.fm0((_this).f34, true, p0, (_this).f34, globalState), _module.__default.fm0(new BigNumber(((((_2420_v33).contains(new _dafny.CodePoint('e'.codePointAt(0)))) ? ((_2420_v33).get(new _dafny.CodePoint('e'.codePointAt(0)))) : (_2373_v8))).length), false, p0, (_this).f30, globalState));
          let _2421_v34;
          _2421_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,_this.f36);
          let _2422_v35;
          let _init71 = function (_2423_i2) {
            return _this.f36;
          };
          let _nw421 = Array((_dafny.ONE).toNumber());
          for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw421.length); _i0_71++) {
            _nw421[_i0_71] = _init71(new BigNumber(_i0_71));
          }
          _2422_v35 = _nw421;
          let _2424_v36;
          let _nw422 = new _module.C0();
          _nw422.__ctor(_2422_v35);
          _2424_v36 = _nw422;
          let _2425_v37;
          _2425_v37 = _module.D9.create_DC27(_2424_v36, _2371_v6);
          let _2426_v38;
          _2426_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2425_v37,_this.f29);
          let _2427_v39;
          let _nw423 = new _module.C6();
          _nw423.__ctor((((_2421_v34).contains(_module.__default.fm46(new BigNumber((_2426_v38).length), (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], _this.f28, globalState))) ? ((_2421_v34).get(_module.__default.fm46(new BigNumber((_2426_v38).length), (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], _this.f28, globalState))) : (_this.f36)), (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], !(!(p1)), (_this).f32, (_this).f33, (_this).f30, p0, p0, _this.f35);
          _2427_v39 = _nw423;
          let _2428_v40;
          _2428_v40 = _dafny.Set.fromElements(_2427_v39, _2427_v39);
          let _2429_v41;
          _2429_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_2428_v40).length), (_this).f30),p0);
          let _index353 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_2418_v32).length));
          (_2418_v32)[_index353] = (((_2429_v41).contains((_this).f34)) ? ((_2429_v41).get((_this).f34)) : (p0));
          (globalState).f21 = _2418_v32;
          r2 = (_dafny.ZERO).minus((_this).f30);
          let _2430_v42;
          let _nw424 = Array((new BigNumber(28)).toNumber());
          _2430_v42 = _nw424;
          let _2431_v43;
          _2431_v43 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2430_v42);
          _2431_v43 = (_2431_v43).update(_this.f29, _2430_v42);
        } else {
          let _2432_v44;
          let _init72 = function (_2433_i3) {
            return _this.f28;
          };
          let _nw425 = Array((new BigNumber(18)).toNumber());
          for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw425.length); _i0_72++) {
            _nw425[_i0_72] = _init72(new BigNumber(_i0_72));
          }
          _2432_v44 = _nw425;
          let _index354 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_2432_v44).length));
          (_2432_v44)[_index354] = _this.f28;
          let _2434_v45;
          _2434_v45 = _dafny.Seq.of(_2367_v2);
          let _index355 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_2432_v44).length));
          let _rhs393 = _module.__default.fm1((_this.f28) || (((_this).f32)[_module.__default.safeIndex((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], new BigNumber(((_this).f32).length))]), new BigNumber(96), globalState);
          let _rhs394 = _dafny.Seq.IsProperPrefixOf(_2434_v45, _2434_v45);
          let _rhs395 = _module.__default.fm1(_this.f29, _module.__default.safeModuloInt((_this).f34, (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]), globalState);
          let _rhs396 = _this.f31;
          let _rhs397 = (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))];
          let _lhs316 = globalState;
          let _lhs317 = _2432_v44;
          let _lhs318 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_2432_v44).length));
          let _lhs319 = globalState;
          let _lhs320 = _this;
          let _lhs321 = globalState;
          _lhs316.f26 = _rhs393;
          _lhs317[_lhs318] = _rhs394;
          _lhs319.f26 = _rhs395;
          _lhs320.f28 = _rhs396;
          _lhs321.f13 = _rhs397;
          let _2435_v47;
          _2435_v47 = function () {
            let _coll63 = new _dafny.Set();
            for (const _compr_63 of (_2365_v0).Elements) {
              let _2436_v46 = _compr_63;
              if ((_2365_v0).contains(_2436_v46)) {
                _coll63.add(_module.__default.safeModuloInt(_2436_v46, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(672), new BigNumber((_dafny.Seq.of(!(!(true)))).length), new BigNumber(866))).length)));
              }
            }
            return _coll63;
          }();
          let _2437_v48;
          let _nw426 = Array((new BigNumber(7)).toNumber());
          _nw426[(_dafny.ZERO).toNumber()] = _2435_v47;
          _nw426[(_dafny.ONE).toNumber()] = _2435_v47;
          _nw426[(new BigNumber(2)).toNumber()] = _2365_v0;
          _nw426[(new BigNumber(3)).toNumber()] = _2435_v47;
          _nw426[(new BigNumber(4)).toNumber()] = ((false) ? (_2435_v47) : (_2435_v47));
          _nw426[(new BigNumber(5)).toNumber()] = _2435_v47;
          _nw426[(new BigNumber(6)).toNumber()] = _2435_v47;
          _2437_v48 = _nw426;
          let _index356 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_2437_v48).length));
          (_2437_v48)[_index356] = _2435_v47;
          let _index357 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_2437_v48).length));
          let _rhs398 = new BigNumber(292);
          let _rhs399 = _2371_v6;
          let _rhs400 = !(true);
          let _rhs401 = _module.__default.fm63((_this.f28) && ((_2432_v44)[_module.__default.safeIndex(new BigNumber(932), new BigNumber((_2432_v44).length))]), p1, p0, globalState);
          let _lhs322 = globalState;
          let _lhs323 = globalState;
          let _lhs324 = _2437_v48;
          let _lhs325 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_2437_v48).length));
          _lhs322.f13 = _rhs398;
          _2371_v6 = _rhs399;
          _lhs323.f1 = _rhs400;
          _lhs324[_lhs325] = _rhs401;
          let _2438_v49;
          _2438_v49 = _dafny.Seq.of(_dafny.Seq.Concat((_this).f32, (_this).f32));
          let _index358 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
          let _rhs402 = new BigNumber(759);
          let _rhs403 = ((_this).f30).plus(((_this).f30).multipliedBy((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]));
          let _rhs404 = (_this).f34;
          let _rhs405 = _this.f36;
          let _rhs406 = _dafny.Seq.Concat(_2438_v49, (_this).f33);
          let _lhs326 = globalState;
          let _lhs327 = _2367_v2;
          let _lhs328 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
          let _lhs329 = _this;
          r1 = _rhs402;
          _lhs326.f13 = _rhs403;
          _lhs327[_lhs328] = _rhs404;
          _lhs329.f36 = _rhs405;
          _2438_v49 = _rhs406;
          let _2439_v50;
          _2439_v50 = _dafny.Seq.of(true);
          _2439_v50 = _2439_v50;
          let _2440_v51;
          _2440_v51 = _module.D4.create_DC11(new BigNumber((_2373_v8).length), (_2432_v44)[_module.__default.safeIndex(new BigNumber(932), new BigNumber((_2432_v44).length))], new BigNumber((_dafny.Seq.UnicodeFromString("tuwhsf")).length));
          let _index359 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
          (_2367_v2)[_index359] = (((_2440_v51).dtor_cf25).minus((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))])).plus((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]);
        }
        let _2441_v52;
        _2441_v52 = _module.D0.create_DC0(p1);
        let _2442_v53;
        _2442_v53 = _dafny.Set.fromElements(p0, (_2441_v52).dtor_cf0);
        let _2443_v54;
        _2443_v54 = _dafny.Seq.of(_2442_v53);
        let _2444_v55;
        let _2445_v56;
        let _2446_v57;
        let _2447_v58;
        let _out73;
        let _out74;
        let _out75;
        let _out76;
        let _outcollector19 = (_this).m2((_this).f30, (_this).f34, (_this).f30, _2443_v54, globalState);
        _out73 = _outcollector19[0];
        _out74 = _outcollector19[1];
        _out75 = _outcollector19[2];
        _out76 = _outcollector19[3];
        _2444_v55 = _out73;
        _2445_v56 = _out74;
        _2446_v57 = _out75;
        _2447_v58 = _out76;
        let _2448_v59;
        let _nw427 = Array((new BigNumber(25)).toNumber()).fill(false);
        _2448_v59 = _nw427;
        let _index360 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_2448_v59).length));
        (_2448_v59)[_index360] = _this.f28;
        let _2449_v60;
        _2449_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f34,(((_2375_v10).contains(_this.f28)) ? ((_2375_v10).get(_this.f28)) : ((_this).f34)));
        let _index361 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_2448_v59).length));
        (_2448_v59)[_index361] = ((_this).f32)[_module.__default.safeIndex((new BigNumber((_2449_v60).length)).plus((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]), new BigNumber(((_this).f32).length))];
      }
      if (p1) {
        r0 = p0;
        if (_this.f31) {
          let _2450_v61;
          let _init73 = function (_2451_i4) {
            return (_this.f28) === (_this.f29);
          };
          let _nw428 = Array((new BigNumber(25)).toNumber());
          for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw428.length); _i0_73++) {
            _nw428[_i0_73] = _init73(new BigNumber(_i0_73));
          }
          _2450_v61 = _nw428;
          (globalState).f21 = _2450_v61;
          let _2452_v62;
          _2452_v62 = _module.D16.create_DC40(_2371_v6, (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]);
          let _2453_v63;
          let _nw429 = new _module.C7();
          _nw429.__ctor((_this).f34, (_2372_v7).Intersect(_2372_v7), (_this).f30, _dafny.Seq.IsProperPrefixOf((_2452_v62).dtor_cf73, _2371_v6), _this.f29, _this.f28);
          _2453_v63 = _nw429;
          _2453_v63 = _2453_v63;
          let _2454_v64;
          _2454_v64 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_this.f36));
          let _index362 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_2450_v61).length));
          (_2450_v61)[_index362] = (_2454_v64).IsProperSubsetOf(_2454_v64);
          let _2455_v65;
          _2455_v65 = _module.D3.create_DC9(_2453_v63.f28, _this.f31, _2450_v61, p1);
          let _2456_v66;
          _2456_v66 = _dafny.Set.fromElements(!(p1), _this.f28, !(_2453_v63.f28));
          let _2457_v67;
          _2457_v67 = _dafny.Map.Empty.slice().updateUnsafe((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))],_this.f31);
          let _2458_v69;
          _2458_v69 = _dafny.Set.fromElements(_2457_v67, function () {
            let _coll64 = new _dafny.Map();
            for (const _compr_64 of (_dafny.Set.fromElements((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))])).Elements) {
              let _2459_v68 = _compr_64;
              if ((_dafny.Set.fromElements((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))])).contains(_2459_v68)) {
                _coll64.push([(_2459_v68).multipliedBy((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]),_2453_v63.f28]);
              }
            }
            return _coll64;
          }());
          let _2460_v70;
          _2460_v70 = _dafny.Seq.of(!(_dafny.Set.fromElements(_this.f35, false)).equals(_2456_v66), (_2458_v69).IsDisjointFrom(_2458_v69), !(true));
          let _2461_v71;
          let _nw430 = Array((new BigNumber(8)).toNumber());
          _nw430[(_dafny.ZERO).toNumber()] = _this.f36;
          _nw430[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
          _nw430[(new BigNumber(2)).toNumber()] = _this.f36;
          _nw430[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
          _nw430[(new BigNumber(4)).toNumber()] = _this.f36;
          _nw430[(new BigNumber(5)).toNumber()] = _this.f36;
          _nw430[(new BigNumber(6)).toNumber()] = (_2371_v6)[_module.__default.safeIndex((_this).f34, new BigNumber((_2371_v6).length))];
          _nw430[(new BigNumber(7)).toNumber()] = _this.f36;
          _2461_v71 = _nw430;
          let _2462_v72;
          _2462_v72 = _dafny.Seq.of((_this).f33, _dafny.Seq.Create(_module.__default.abs(new BigNumber(414)), function (_2463_i5) {
            return (_this).f32;
          }));
          let _2464_v73;
          let _nw431 = new _module.C5();
          _nw431.__ctor(_2453_v63.f29, (_this).f34, (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], p1, (_this).f32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(210)), function (_2465_i6) {
            return (_this).f32;
          }), (_2453_v63).f30, !(p0), _this.f31, p1);
          _2464_v73 = _nw431;
          let _2466_v74;
          _2466_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2464_v73,new BigNumber((_2368_v3).length));
          let _2467_v75;
          let _nw432 = new _module.C5();
          _nw432.__ctor(_this.f35, (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], (_2453_v63).f30, _this.f35, _2460_v70, (_2462_v72)[_module.__default.safeIndex(new BigNumber((_2466_v74).length), new BigNumber((_2462_v72).length))], new BigNumber((_2460_v70).length), _2464_v73.f28, p0, p1);
          _2467_v75 = _nw432;
          let _2468_v76;
          _2468_v76 = _dafny.Map.Empty.slice().updateUnsafe(_2461_v71,_2467_v75);
          let _pat_let_tv33 = _2450_v61;
          let _pat_let_tv34 = _2450_v61;
          let _index363 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_2450_v61).length));
          let _rhs407 = (_this.f29) || (p0);
          let _rhs408 = function (_pat_let43_0) {
            return function (_2469_dt__update__tmp_h3) {
              return function (_pat_let44_0) {
                return function (_2470_dt__update_hcf20_h0) {
                  return _module.D3.create_DC9((_2469_dt__update__tmp_h3).dtor_cf18, (_2469_dt__update__tmp_h3).dtor_cf19, _2470_dt__update_hcf20_h0, (_2469_dt__update__tmp_h3).dtor_cf21);
                }(_pat_let44_0);
              }(((_this.f29) ? (_pat_let_tv33) : (_pat_let_tv34)));
            }(_pat_let43_0);
          }(_2455_v65);
          let _rhs409 = _module.__default.safeDivisionInt((_this).f34, new BigNumber((_2468_v76).length));
          let _rhs410 = (_2467_v75).f32;
          let _lhs330 = _2450_v61;
          let _lhs331 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_2450_v61).length));
          _lhs330[_lhs331] = _rhs407;
          _2455_v65 = _rhs408;
          r2 = _rhs409;
          _2460_v70 = _rhs410;
          let _index364 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
          (_2367_v2)[_index364] = new BigNumber(((_2456_v66).Intersect((_dafny.Set.fromElements((_2450_v61)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_2450_v61).length))], _2453_v63.f29, _2453_v63.f31, _2453_v63.f29, (_this).fm7((_2372_v7).update((_dafny.ZERO).minus((_this).f34), _module.__default.abs(new BigNumber(154))), (_2467_v75).f34, globalState))).Difference(_2456_v66))).length);
          let _2471_v77;
          let _nw433 = new _module.C4();
          _nw433.__ctor((_2467_v75).f30, _this.f28, (_2467_v75).f32, (_2464_v73).f33, new BigNumber((_2457_v67).length), _2464_v73.f31, _2467_v75.f31, _this.f35);
          _2471_v77 = _nw433;
          let _2472_v78;
          _2472_v78 = _dafny.MultiSet.fromElements(_2471_v77, _2471_v77);
          let _2473_v79;
          let _nw434 = Array((new BigNumber(6)).toNumber()).fill(_module.D10.Default());
          _2473_v79 = _nw434;
          let _2474_v80;
          _2474_v80 = _module.D20.create_DC48(new BigNumber((_2472_v78).cardinality()), _2473_v79, _2467_v75.f29);
          (globalState).f13 = (_2474_v80).dtor_cf82;
        } else {
          let _2475_v81;
          _2475_v81 = _dafny.Map.Empty.slice().updateUnsafe(_2375_v10,p1);
          let _2476_v82;
          let _init74 = function (_2477_i7) {
            return (_this).f32;
          };
          let _nw435 = Array((_dafny.ONE).toNumber());
          for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw435.length); _i0_74++) {
            _nw435[_i0_74] = _init74(new BigNumber(_i0_74));
          }
          _2476_v82 = _nw435;
          let _2478_v83;
          _2478_v83 = _dafny.Set.fromElements((((_2373_v8).contains(p0)) ? ((_2373_v8).get(p0)) : (_this.f29)), _this.f28);
          let _2479_v84;
          let _nw436 = new _module.C3();
          _nw436.__ctor(new BigNumber((_2478_v83).length), _2476_v82, _this.f36, (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], false, (_this).f32, (_this).f33, new BigNumber((_2365_v0).length), _module.__default.fm1(p1, (_this).f30, globalState), true, p1);
          _2479_v84 = _nw436;
          let _2480_v85;
          _2480_v85 = _dafny.Map.Empty.slice().updateUnsafe((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))],_2479_v84);
          let _2481_v86;
          _2481_v86 = _dafny.Seq.of(_2479_v84, (((_2480_v85).contains((_this).f30)) ? ((_2480_v85).get((_this).f30)) : (_2479_v84)));
          let _2482_v87;
          _2482_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f28,_this.f36)).length),(_this).f34);
          let _2483_v88;
          let _nw437 = new _module.C3();
          _nw437.__ctor(new BigNumber((_2475_v81).length), _2476_v82, _this.f36, ((_this).f34).minus((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]), (_this).fm7(_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))])), new BigNumber((_2481_v86).length), globalState), (_this).f32, (_this).f33, (((_2482_v87).contains((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))])) ? ((_2482_v87).get((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))])) : ((_2479_v84).f44)), _this.f28, ((_this).f34).isLessThanOrEqualTo((_2479_v84).f44), (_2375_v10).IsProperSubsetOf(_2375_v10));
          _2483_v88 = _nw437;
          let _2484_v89;
          let _nw438 = new _module.C4();
          _nw438.__ctor(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))])).length), (_2479_v84).f44), (new BigNumber(213)).isLessThanOrEqualTo(new BigNumber(452)), (_this).f32, _dafny.Seq.Concat(_dafny.Seq.of((_this).f32, (_this).f32, (_this).f32), (_this).f33), (_2479_v84).f44, !(false) || (p1), _this.f28, p1);
          _2484_v89 = _nw438;
          let _index365 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
          (_2367_v2)[_index365] = (_dafny.ZERO).minus((_2479_v84).f44);
          (globalState).f26 = (((_this).f30).isEqualTo((_this).f34)) === ((_2484_v89).fm6(globalState));
          let _2485_v90;
          _2485_v90 = _dafny.Seq.of(_2372_v7, _2372_v7);
          let _2486_v91;
          let _nw439 = new _module.C7();
          _nw439.__ctor((_2479_v84).f44, (_2485_v90)[_module.__default.safeIndex((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], new BigNumber((_2485_v90).length))], new BigNumber((_dafny.Seq.update(_2371_v6, _module.__default.safeIndex((_2479_v84).f44, new BigNumber((_2371_v6).length)), _this.f36)).length), _this.f28, p1, (_2484_v89).fm44((_this).f30, globalState));
          _2486_v91 = _nw439;
          let _nw440 = new _module.C7();
          _nw440.__ctor((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_2486_v91).f30, (_2483_v88).f44)), _module.__default.fm25(globalState), (_2486_v91).f30, true, !((_2479_v84).f44).isEqualTo((_2483_v88).f44), _2486_v91.f28);
          _2486_v91 = _nw440;
        }
        _2373_v8 = (_2373_v8).update(p0, !(!(_this.f31)));
        (globalState).f13 = (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))];
        let _2487_v92;
        _2487_v92 = _module.D0.create_DC1(_2375_v10);
        let _2488_v93;
        _2488_v93 = _dafny.Map.Empty.slice().updateUnsafe((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))],_2487_v92);
        let _2489_v94;
        _2489_v94 = _dafny.Set.fromElements(_this.f28);
        let _2490_v96;
        let _nw441 = Array((new BigNumber(15)).toNumber());
        _nw441[(_dafny.ZERO).toNumber()] = (_2488_v93).update(new BigNumber((_2489_v94).length), _2487_v92);
        _nw441[(_dafny.ONE).toNumber()] = _2488_v93;
        _nw441[(new BigNumber(2)).toNumber()] = _2488_v93;
        _nw441[(new BigNumber(3)).toNumber()] = _2488_v93;
        _nw441[(new BigNumber(4)).toNumber()] = _module.__default.fm64((_this).f30, _this.f35, (_dafny.ZERO).minus((_this).f34), _this.f28, globalState);
        _nw441[(new BigNumber(5)).toNumber()] = ((_this.f31) ? (_2488_v93) : (_2488_v93));
        _nw441[(new BigNumber(6)).toNumber()] = _2488_v93;
        _nw441[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f34)),_2487_v92);
        _nw441[(new BigNumber(8)).toNumber()] = _2488_v93;
        _nw441[(new BigNumber(9)).toNumber()] = _2488_v93;
        _nw441[(new BigNumber(10)).toNumber()] = function () {
          let _coll65 = new _dafny.Map();
          for (const _compr_65 of (_2374_v9).Elements) {
            let _2491_v95 = _compr_65;
            if (_dafny.Seq.contains(_2374_v9, _2491_v95)) {
              _coll65.push([(_2491_v95).minus((_this).f30),_module.D0.create_DC1(_dafny.MultiSet.fromElements(_this.f28))]);
            }
          }
          return _coll65;
        }();
        _nw441[(new BigNumber(11)).toNumber()] = (_2488_v93).Merge(_2488_v93);
        _nw441[(new BigNumber(12)).toNumber()] = _2488_v93;
        _nw441[(new BigNumber(13)).toNumber()] = (_2488_v93).Merge(_2488_v93);
        _nw441[(new BigNumber(14)).toNumber()] = (_2488_v93).Merge(_2488_v93);
        _2490_v96 = _nw441;
        let _index366 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2490_v96).length));
        (_2490_v96)[_index366] = _2488_v93;
        let _2492_v97;
        _2492_v97 = _module.D8.create_DC22((_this).f30, new BigNumber((_dafny.Seq.update(_2371_v6, _module.__default.safeIndex((_this).f34, new BigNumber((_2371_v6).length)), _this.f36)).length));
        let _2493_v98;
        _2493_v98 = _dafny.Map.Empty.slice().updateUnsafe(_this.f28,_2492_v97);
        let _2494_v100;
        _2494_v100 = _dafny.Map.Empty.slice().updateUnsafe((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))],p0);
        let _index367 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2490_v96).length));
        let _rhs411 = (_this).f34;
        let _rhs412 = (_2488_v93).Merge(function () {
          let _coll66 = new _dafny.Map();
          for (const _compr_66 of (_2494_v100).Keys.Elements) {
            let _2495_v99 = _compr_66;
            if ((_2494_v100).contains(_2495_v99)) {
              _coll66.push([_module.__default.safeDivisionInt(_2495_v99, (_this).f30),_2487_v92]);
            }
          }
          return _coll66;
        }());
        let _rhs413 = (_this).f34;
        let _rhs414 = ((p1) ? (_2493_v98) : ((_module.__default.fm65((_this).f34, _this.f35, (_this).f34, globalState)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_this.f35,_2492_v97)).update(_this.f35, _module.D8.create_DC22((_this).f30, (_this).f34)))));
        let _rhs415 = _this.f28;
        let _lhs332 = _2490_v96;
        let _lhs333 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2490_v96).length));
        let _lhs334 = globalState;
        let _lhs335 = globalState;
        r1 = _rhs411;
        _lhs332[_lhs333] = _rhs412;
        _lhs334.f13 = _rhs413;
        _2493_v98 = _rhs414;
        _lhs335.f1 = _rhs415;
      } else {
        if (_this.f28) {
          (globalState).f13 = (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))];
          (_this).f36 = (_2371_v6)[_module.__default.safeIndex((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], new BigNumber((_2371_v6).length))];
          let _init75 = ((_2496_p1) => function (_2497_i8) {
            return _2496_p1;
          })(p1);
          let _nw442 = Array((new BigNumber(23)).toNumber());
          for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw442.length); _i0_75++) {
            _nw442[_i0_75] = _init75(new BigNumber(_i0_75));
          }
          (globalState).f21 = _nw442;
          let _2498_v101;
          _2498_v101 = _module.D0.create_DC1(_2375_v10);
          let _2499_v102;
          _2499_v102 = _dafny.Map.Empty.slice().updateUnsafe((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))],_this.f31);
          let _2500_v103;
          _2500_v103 = _dafny.Seq.of(_2499_v102);
          let _2501_v104;
          _2501_v104 = _dafny.Set.fromElements(false);
          let _2502_v105;
          _2502_v105 = _module.D2.create_DC6((_this).f34, _this.f35, new BigNumber((_2501_v104).length), _this.f35, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f31,(_this).f30)).length));
          let _index368 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
          let _rhs416 = _module.__default.fm15(new BigNumber((_2374_v9).length), new BigNumber((_2500_v103).length), globalState);
          let _rhs417 = (_dafny.Map.Empty.slice().updateUnsafe(_this.f35,(_2502_v105).dtor_cf10)).Merge(_2373_v8);
          let _rhs418 = new BigNumber(121);
          let _rhs419 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_2371_v6, _2371_v6), _2371_v6);
          let _lhs336 = _2367_v2;
          let _lhs337 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
          _2498_v101 = _rhs416;
          _2373_v8 = _rhs417;
          _lhs336[_lhs337] = _rhs418;
          r3 = _rhs419;
          (_this).f36 = _this.f36;
        } else {
          let _2503_v106;
          _2503_v106 = _dafny.Map.Empty.slice().updateUnsafe(_2367_v2,(_this).fm4(p0, _2365_v0, (_this).f30, p1, globalState));
          let _2504_v107;
          _2504_v107 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_this.f31);
          let _2505_v108;
          let _nw443 = new _module.C8();
          _nw443.__ctor(p1, _this.f36, new BigNumber((_2371_v6).length), _this.f31, (_this).f32, (_this).f33, new BigNumber((_2374_v9).length), false, _this.f28, _this.f35);
          _2505_v108 = _nw443;
          let _2506_v109;
          _2506_v109 = _dafny.MultiSet.fromElements(_2505_v108, _2505_v108, _2505_v108, _2505_v108);
          let _2507_v110;
          let _nw444 = new _module.C7();
          _nw444.__ctor((new BigNumber((_2503_v106).length)).minus(new BigNumber((_2372_v7).cardinality())), (_2372_v7).Intersect(_2372_v7), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-618)), function (_2508_i9) {
            return _this.f36;
          })).length)), (((_2504_v107).contains(new BigNumber((_dafny.Seq.update(_2371_v6, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f30), new BigNumber((_2371_v6).length)), _this.f36)).length))) ? ((_2504_v107).get(new BigNumber((_dafny.Seq.update(_2371_v6, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f30), new BigNumber((_2371_v6).length)), _this.f36)).length))) : (_this.f28)), _dafny.Seq.contains((_this).f32, _this.f31), (_2506_v109).IsDisjointFrom(_2506_v109));
          _2507_v110 = _nw444;
          _2507_v110 = _2507_v110;
          let _2509_v111;
          let _init76 = function (_2510_i10) {
            return true;
          };
          let _nw445 = Array((new BigNumber(6)).toNumber());
          for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw445.length); _i0_76++) {
            _nw445[_i0_76] = _init76(new BigNumber(_i0_76));
          }
          _2509_v111 = _nw445;
          let _index369 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_2509_v111).length));
          (_2509_v111)[_index369] = p1;
          let _index370 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_2509_v111).length));
          (_2509_v111)[_index370] = ((_2507_v110.f31) || (!(_2507_v110.f28))) && (_2507_v110.f31);
          let _2511_v112;
          let _nw446 = new _module.C4();
          _nw446.__ctor(_module.__default.safeDivisionInt((_module.D10.create_DC31(_2505_v108.f29, _this.f28, (_2507_v110).f30)).dtor_cf62, new BigNumber(754)), true, (_this).f32, (_2505_v108).f33, (_2505_v108).f30, (_2509_v111)[_module.__default.safeIndex(new BigNumber(527), new BigNumber((_2509_v111).length))], (_2370_v5).contains(_2368_v3), !(_this.f31));
          _2511_v112 = _nw446;
          let _2512_v113;
          _2512_v113 = _dafny.Seq.of((_dafny.MultiSet.fromElements((_2509_v111)[_module.__default.safeIndex(new BigNumber(527), new BigNumber((_2509_v111).length))])).update(_2507_v110.f31, _module.__default.abs(_module.__default.fm24(_2507_v110.f31, (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], globalState))), (_dafny.MultiSet.fromElements(_2505_v108.f28, _2507_v110.f29)).Intersect((_dafny.MultiSet.fromElements(_this.f28)).update(_2505_v108.f28, _module.__default.abs((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]))), (_2375_v10).update(_2507_v110.f29, _module.__default.abs((_2505_v108).f30)));
          _2512_v113 = _dafny.Seq.of(((_2507_v110.f29) ? (_2375_v10) : ((_2375_v10).update(false, _module.__default.abs((_this).f30)))));
          let _2513_v114;
          let _nw447 = new _module.C4();
          _nw447.__ctor(_module.__default.safeDivisionInt((_2505_v108).f30, (_this).f30), p1, (_this).f32, (_this).f33, (_dafny.ZERO).minus((_this).f30), _2505_v108.f28, _this.f29, false);
          _2513_v114 = _nw447;
          let _rhs420 = _2513_v114;
          let _rhs421 = _2371_v6;
          _2513_v114 = _rhs420;
          _2371_v6 = _rhs421;
        }
        let _2514_v115;
        let _nw448 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2514_v115 = _nw448;
        let _index371 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_2514_v115).length));
        (_2514_v115)[_index371] = _2371_v6;
        let _2515_v116;
        _2515_v116 = _module.D16.create_DC40((_this).fm3(_this.f29, p0, globalState), (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]);
        let _index372 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_2514_v115).length));
        (_2514_v115)[_index372] = (_2515_v116).dtor_cf73;
        (globalState).f1 = _this.f31;
        let _2516_v117;
        let _nw449 = Array((new BigNumber(24)).toNumber()).fill(false);
        _2516_v117 = _nw449;
        let _index373 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_2516_v117).length));
        (_2516_v117)[_index373] = ((_this).f32)[_module.__default.safeIndex((_this).f34, new BigNumber(((_this).f32).length))];
        let _2517_v118;
        let _init77 = function (_2518_i11) {
          return _this.f36;
        };
        let _nw450 = Array((new BigNumber(21)).toNumber());
        for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw450.length); _i0_77++) {
          _nw450[_i0_77] = _init77(new BigNumber(_i0_77));
        }
        _2517_v118 = _nw450;
        let _index374 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_2517_v118).length));
        (_2517_v118)[_index374] = _this.f36;
        let _2519_v119;
        _2519_v119 = _module.D0.create_DC1(_2375_v10);
        let _2520_v120;
        _2520_v120 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-393),_this.f28);
        let _pat_let_tv35 = _2375_v10;
        let _index375 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_2516_v117).length));
        let _index376 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_2517_v118).length));
        let _index377 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
        let _rhs422 = _this.f28;
        let _rhs423 = _module.__default.fm14((_this).f34, !(_this.f29) || (p0), _module.__default.fm39(_this.f28, _dafny.MultiSet.fromElements((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]), _this.f35, globalState), function (_pat_let45_0) {
          return function (_2521_dt__update__tmp_h4) {
            return function (_pat_let46_0) {
              return function (_2522_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_2522_dt__update_hcf1_h0);
              }(_pat_let46_0);
            }(_pat_let_tv35);
          }(_pat_let45_0);
        }(_2519_v119), globalState);
        let _rhs424 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(614)), function (_2523_i12) {
          return _this.f36;
        }), (_2514_v115)[_module.__default.safeIndex(new BigNumber(349), new BigNumber((_2514_v115).length))])).length)).plus((_this).f34);
        let _rhs425 = (_this).f34;
        let _rhs426 = (new BigNumber((((_2520_v120).update(new BigNumber(743), false)).update((_dafny.ZERO).minus((_this).f30), _this.f35)).length)).plus((_this).f34);
        let _lhs338 = _2516_v117;
        let _lhs339 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_2516_v117).length));
        let _lhs340 = _2517_v118;
        let _lhs341 = _module.__default.safeIndex(new BigNumber(786), new BigNumber((_2517_v118).length));
        let _lhs342 = _2367_v2;
        let _lhs343 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length));
        _lhs338[_lhs339] = _rhs422;
        _lhs340[_lhs341] = _rhs423;
        _lhs342[_lhs343] = _rhs424;
        r1 = _rhs425;
        r1 = _rhs426;
        _2367_v2 = _2367_v2;
      }
      let _2524_v121;
      let _nw451 = new _module.C5();
      _nw451.__ctor(p1, (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], (((_2372_v7).contains((_this).f34)) ? ((_2372_v7).get((_this).f34)) : (new BigNumber(444))), _this.f28, (_this).f32, (_this).f33, (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))], _this.f31, _this.f31, !(_this.f29));
      _2524_v121 = _nw451;
      let _2525_v122;
      _2525_v122 = _2365_v0;
      _2525_v122 = ((!(p1)) ? (_2525_v122) : (_2525_v122));
      let _2526_v123;
      _2526_v123 = _dafny.Map.Empty.slice().updateUnsafe((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))],(_2373_v8).update(p0, p0));
      let _2527_v124;
      let _nw452 = Array((new BigNumber(25)).toNumber()).fill([]);
      _2527_v124 = _nw452;
      let _2528_v125;
      _2528_v125 = _dafny.Map.Empty.slice().updateUnsafe(_2527_v124,(_this).f33);
      let _2529_v126;
      _2529_v126 = _module.D14.create_DC36((_this).f33);
      let _2530_v127;
      _2530_v127 = _dafny.Set.fromElements(p1);
      let _2531_v128;
      let _nw453 = new _module.C9();
      _nw453.__ctor((_dafny.ZERO).minus((_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))]), _2526_v123, _this.f36, new BigNumber(((_this).f32).length), ((_this).f34).isLessThanOrEqualTo((_this).f30), (_this).f32, _dafny.Seq.Concat((((_2528_v125).contains(_2527_v124)) ? ((_2528_v125).get(_2527_v124)) : ((_2529_v126).dtor_cf67)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(890)), function (_2532_i13) {
        return (_this).f32;
      })), (_this).f30, (_2524_v121).f47, (_2530_v127).IsSubsetOf(_module.__default.fm11((_this).f33, p1, (_2524_v121).f47, new BigNumber(195), globalState)), _dafny.Seq.contains(_dafny.Seq.of(_2366_v1), _2366_v1));
      _2531_v128 = _nw453;
      r0 = (_2524_v121).fm6(globalState);
      r1 = (((_2368_v3).contains((_2524_v121).f47)) ? ((_2368_v3).get((_2524_v121).f47)) : ((new BigNumber(((_this).f32).length)).minus((_this).f34)));
      r2 = (_2367_v2)[_module.__default.safeIndex(new BigNumber(458), new BigNumber((_2367_v2).length))];
      let _2533_v129;
      _2533_v129 = _dafny.Set.fromElements(_2367_v2, _2367_v2);
      r3 = (_this).fm7(_dafny.MultiSet.FromArray(_2374_v9), new BigNumber(((_2533_v129).Difference(_2533_v129)).length), globalState);
      return [r0, r1, r2, r3];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
