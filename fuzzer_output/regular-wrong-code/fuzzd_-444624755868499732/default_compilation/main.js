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
    static fm3(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC6(_dafny.Seq.UnicodeFromString("t"), new BigNumber(-80));
    };
    static fm5(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qt")).length)))).cardinality()), new BigNumber((_dafny.Set.fromElements(true)).length))).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber(-958))), !(!(false)) || (false), !(false));
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('j'.codePointAt(0));
    };
    static fm7(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D1.create_DC4(_dafny.Set.fromElements(new BigNumber(934)));
      if (_source0.is_DC5) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(207)), function (_0_i0) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(626)), function (_1_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }));
      } else if (_source0.is_DC6) {
        let _2___mcc_h0 = (_source0).cf5;
        let _3___mcc_h1 = (_source0).cf6;
        let _4_cf6 = _3___mcc_h1;
        let _5_cf5 = _2___mcc_h0;
        return _5_cf5;
      } else {
        let _6___mcc_h2 = (_source0).cf4;
        let _7_cf4 = _6___mcc_h2;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(997)), function (_8_i2) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        });
      }
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!(true) || (false),(new BigNumber(984)).isLessThanOrEqualTo(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-626))).length))).cardinality())));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("je"), _dafny.Seq.UnicodeFromString("rfsuhwou")), ((true) ? (_dafny.Seq.UnicodeFromString("a")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), function (_9_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }))));
    };
    static fm14(p0, p1, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(780), new BigNumber(411), new BigNumber(574), new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber(362)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length)), new BigNumber(986), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, false)).length)),false)).length))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(187)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), function (_10_i0) {
        return new BigNumber((_dafny.Set.fromElements(false, true)).length);
      })), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(338)), function (_11_i1) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      })).length), new BigNumber(623), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-883)))).length))).length))).length)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(940)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(221)), function (_12_i2) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),false)).length);
      })));
    };
    static fm15(p0, globalState) {
      return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(710))).length),new BigNumber(-724)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-115),new BigNumber(120)));
    };
    static fm16(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(false)).Intersect((_dafny.MultiSet.fromElements(!(true), true)).Intersect(_dafny.MultiSet.fromElements(false, false, false)));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC3(false)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-390)), function (_13_i0) {
          return _module.D0.create_DC3(false);
        })))).Elements) {
          let _14_v0 = _compr_0;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC3(false)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-390)), function (_13_i0) {
            return _module.D0.create_DC3(false);
          })))).contains(_14_v0)) {
            _coll0.push([_14_v0,!((new BigNumber((_dafny.Seq.of(true)).length)).isEqualTo(new BigNumber(713)))]);
          }
        }
        return _coll0;
      }();
    };
    static fm18(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('y'.codePointAt(0));
    };
    static fm19(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("jx")).length))).length), new BigNumber(-653), new BigNumber((_dafny.Seq.UnicodeFromString("pbdbqo")).length), new BigNumber(-821), new BigNumber((_dafny.Seq.UnicodeFromString("u")).length))))).cardinality()))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber(387))).length)))).length)))).Union(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC30(new BigNumber((_dafny.Seq.UnicodeFromString("wrnolacmv")).length)),new BigNumber(291))).length))));
    };
    static fm20(p0, p1, p2, globalState) {
      return _module.D0.create_DC2();
    };
    static fm21(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(true, !(false), false, false)).Intersect((_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(true)));
    };
    static fm22(p0, p1, globalState) {
      let _source1 = _dafny.MultiSet.fromElements(false);
      let _15___mcc_h0 = _source1;
      let _16_cf64 = _15___mcc_h0;
      return _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(!(false),false));
    };
    static fm23(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,!(true))),false)).Merge(((false) ? (_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,true)),false)) : (_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,false)),true))));
    };
    static fm24(p0, p1, p2, globalState) {
      return _module.D7.create_DC21((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(true,true)),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(true,true)),true)));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC1((true) || (true), (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(792)))).cardinality()))));
    };
    static fm26(p0, p1, p2, globalState) {
      let _source2 = _module.D4.create_DC14(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(315)), function (_17_i0) {
  return new _dafny.CodePoint('u'.codePointAt(0));
})).length))).length), !(!(true)), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(807))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),new BigNumber(865))).length))).cardinality()), !(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),true)).length));
      if (_source2.is_DC12) {
        let _18___mcc_h0 = (_source2).cf14;
        let _19___mcc_h1 = (_source2).cf15;
        let _20___mcc_h2 = (_source2).cf16;
        let _21_cf16 = _20___mcc_h2;
        let _22_cf15 = _19___mcc_h1;
        let _23_cf14 = _18___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_21_cf16,_22_cf15),_22_cf15);
      } else if (_source2.is_DC13) {
        let _24___mcc_h3 = (_source2).cf17;
        let _25___mcc_h4 = (_source2).cf18;
        let _26___mcc_h5 = (_source2).cf19;
        let _27_cf19 = _26___mcc_h5;
        let _28_cf18 = _25___mcc_h4;
        let _29_cf17 = _24___mcc_h3;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(_29_cf17),new BigNumber((_dafny.Set.fromElements(_27_cf19)).length)),_28_cf18);
      } else if (_source2.is_DC14) {
        let _30___mcc_h6 = (_source2).cf20;
        let _31___mcc_h7 = (_source2).cf21;
        let _32___mcc_h8 = (_source2).cf22;
        let _33___mcc_h9 = (_source2).cf23;
        let _34___mcc_h10 = (_source2).cf24;
        let _35_cf24 = _34___mcc_h10;
        let _36_cf23 = _33___mcc_h9;
        let _37_cf22 = _32___mcc_h8;
        let _38_cf21 = _31___mcc_h7;
        let _39_cf20 = _30___mcc_h6;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_36_cf23,new BigNumber((_dafny.Set.fromElements(_35_cf24)).length)),_39_cf20)).Merge(function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(_37_cf22)),!(_38_cf21))).Keys.Elements) {
            let _40_v0 = _compr_1;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(_37_cf22)),!(_38_cf21))).contains(_40_v0)) {
              _coll1.push([_40_v0,_39_cf20]);
            }
          }
          return _coll1;
        }());
      } else if (_source2.is_DC11) {
        let _41___mcc_h11 = (_source2).cf13;
        let _42_cf13 = _41___mcc_h11;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(438)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(90),true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(345)),new BigNumber(705)));
      } else {
        let _43___mcc_h12 = (_source2).cf25;
        let _44_cf25 = _43___mcc_h12;
        if (true) {
          return function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(322)))).Elements) {
              let _45_v1 = _compr_2;
              if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(322)))).contains(_45_v1)) {
                _coll2.push([_45_v1,new BigNumber(288)]);
              }
            }
            return _coll2;
          }();
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cjrwfbgy")).length),new BigNumber(-569))).length))),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(486),new BigNumber(-839))).length)));
        }
      }
    };
    static fm27(p0, globalState) {
      return _module.D4.create_DC13(!(true) || (true), new BigNumber(915), ((true) ? (false) : (false)));
    };
    static fm28(p0, globalState) {
      return _dafny.Seq.of(new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber(475))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(619), new BigNumber(473), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))))).cardinality()), new BigNumber(547));
    };
    static fm29(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(611)), function (_46_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),false);
      }), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),!(true))));
    };
    static fm30(globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("qqwr"));
    };
    static fm31(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(832),true)).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(828),!(false))).Merge(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-711), new BigNumber(941))) {
          let _47_v0 = _compr_3;
          if (((new BigNumber(-711)).isLessThanOrEqualTo(_47_v0)) && ((_47_v0).isLessThan(new BigNumber(941)))) {
            _coll3.push([_module.__default.safeModuloInt(_47_v0, new BigNumber(433)),true]);
          }
        }
        return _coll3;
      }()));
    };
    static fm32(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(true,!(false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(true,true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Map.Empty.slice().updateUnsafe(false,true))));
    };
    static fm33(globalState) {
      return _dafny.Seq.of((true) || (false), true);
    };
    static fm34(p0, globalState) {
      return ((function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of (function () {
            let _coll6 = new _dafny.Map();
            for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),false)).Keys.Elements) {
              let _48_v1 = _compr_6;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),false)).contains(_48_v1)) {
                _coll6.push([_module.__default.safeDivisionInt(_48_v1, new BigNumber((_dafny.Seq.UnicodeFromString("ajo")).length)),new BigNumber(229)]);
              }
            }
            return _coll6;
          }()).Keys.Elements) {
            let _49_v0 = _compr_5;
            if ((function () {
              let _coll7 = new _dafny.Map();
              for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),false)).Keys.Elements) {
                let _48_v1 = _compr_7;
                if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),false)).contains(_48_v1)) {
                  _coll7.push([_module.__default.safeDivisionInt(_48_v1, new BigNumber((_dafny.Seq.UnicodeFromString("ajo")).length)),new BigNumber(229)]);
                }
              }
              return _coll7;
            }()).contains(_49_v0)) {
              _coll5.push([(_49_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,false)))).length)),_dafny.Seq.UnicodeFromString("jvslhv")]);
            }
          }
          return _coll5;
        }()).Keys.Elements) {
          let _50_v2 = _compr_4;
          if ((function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (function () {
              let _coll9 = new _dafny.Map();
              for (const _compr_9 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),false)).Keys.Elements) {
                let _48_v1 = _compr_9;
                if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),false)).contains(_48_v1)) {
                  _coll9.push([_module.__default.safeDivisionInt(_48_v1, new BigNumber((_dafny.Seq.UnicodeFromString("ajo")).length)),new BigNumber(229)]);
                }
              }
              return _coll9;
            }()).Keys.Elements) {
              let _49_v0 = _compr_8;
              if ((function () {
                let _coll10 = new _dafny.Map();
                for (const _compr_10 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),false)).Keys.Elements) {
                  let _48_v1 = _compr_10;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),false)).contains(_48_v1)) {
                    _coll10.push([_module.__default.safeDivisionInt(_48_v1, new BigNumber((_dafny.Seq.UnicodeFromString("ajo")).length)),new BigNumber(229)]);
                  }
                }
                return _coll10;
              }()).contains(_49_v0)) {
                _coll8.push([(_49_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,false)))).length)),_dafny.Seq.UnicodeFromString("jvslhv")]);
              }
            }
            return _coll8;
          }()).contains(_50_v2)) {
            _coll4.add((_50_v2).minus((_dafny.ZERO).minus(new BigNumber(-61))));
          }
        }
        return _coll4;
      }()).Union(function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(567), new BigNumber(438))) {
          let _51_v3 = _compr_11;
          if (((new BigNumber(567)).isLessThanOrEqualTo(_51_v3)) && ((_51_v3).isLessThan(new BigNumber(438)))) {
            _coll11.add((_51_v3).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-116)), function (_52_i0) {
              return new _dafny.CodePoint('s'.codePointAt(0));
            })).length)));
          }
        }
        return _coll11;
      }())).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("xjjpql")).length))).cardinality()), new BigNumber(404), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(508),new _dafny.CodePoint('y'.codePointAt(0)))).length))).cardinality()), new BigNumber(944)));
    };
    static fm35(p0, p1, p2, p3, globalState) {
      return ((function () {
        let _coll12 = new _dafny.Set();
        for (const _compr_12 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(570)), function (_53_i0) {
          return _dafny.Seq.UnicodeFromString("a");
        })).Elements) {
          let _54_v0 = _compr_12;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(570)), function (_53_i0) {
            return _dafny.Seq.UnicodeFromString("a");
          }), _54_v0)) {
            _coll12.add(_54_v0);
          }
        }
        return _coll12;
      }()).Difference(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("kejowhwr")))).Union((function () {
        let _coll13 = new _dafny.Set();
        for (const _compr_13 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("iofop"),_dafny.Seq.of(true, false))).Keys.Elements) {
          let _55_v1 = _compr_13;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("iofop"),_dafny.Seq.of(true, false))).contains(_55_v1)) {
            _coll13.add(_55_v1);
          }
        }
        return _coll13;
      }()).Intersect(function () {
        let _coll14 = new _dafny.Set();
        for (const _compr_14 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("thbw"))).Elements) {
          let _56_v2 = _compr_14;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("thbw"))).contains(_56_v2)) {
            _coll14.add(_56_v2);
          }
        }
        return _coll14;
      }()));
    };
    static fm36(p0, p1, p2, globalState) {
      return _module.D3.create_DC9((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(949)))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(-759), new BigNumber(320))), _module.D0.create_DC2(), new BigNumber(829));
    };
    static fm37(p0, p1, p2, globalState) {
      return (_module.D3.create_DC9(_dafny.MultiSet.fromElements(new BigNumber(198), new BigNumber((function () {
  let _coll15 = new _dafny.Set();
  for (const _compr_15 of (function () {
    let _coll16 = new _dafny.Map();
    for (const _compr_16 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)))).Elements) {
      let _57_v0 = _compr_16;
      if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)))).contains(_57_v0)) {
        _coll16.push([_57_v0,false]);
      }
    }
    return _coll16;
  }()).Keys.Elements) {
    let _58_v1 = _compr_15;
    if ((function () {
      let _coll17 = new _dafny.Map();
      for (const _compr_17 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)))).Elements) {
        let _57_v0 = _compr_17;
        if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)))).contains(_57_v0)) {
          _coll17.push([_57_v0,false]);
        }
      }
      return _coll17;
    }()).contains(_58_v1)) {
      _coll15.add(_58_v1);
    }
  }
  return _coll15;
}()).length)), _module.D0.create_DC2(), new BigNumber(690))).dtor_cf11;
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return (_module.D12.create_DC35(true, new BigNumber(723), new BigNumber(135))).dtor_cf57;
    };
    static fm39(globalState) {
      if (!((new BigNumber(-773)).isEqualTo(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(11))).length), new BigNumber((_dafny.Set.fromElements(true)).length))).length)))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("gnr")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(643)));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(824))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("pmlngaghv")).length)));
      }
    };
    static m0(p0, globalState) {
      let r0 = [];
      let r1 = false;
      let r2 = _dafny.MultiSet.Empty;
      let r3 = _dafny.ZERO;
      let _59_v0;
      _59_v0 = true;
      if ((_59_v0) || ((_59_v0) || (_59_v0))) {
        let _60_v1;
        _60_v1 = new BigNumber(125);
        let _61_v2;
        _61_v2 = _dafny.Seq.UnicodeFromString("or");
        let _62_v3;
        _62_v3 = _dafny.MultiSet.fromElements(p0);
        let _63_v4;
        _63_v4 = _module.D9.create_DC27(_62_v3);
        let _64_v5;
        _64_v5 = _dafny.Seq.of(_63_v4);
        let _65_v6;
        _65_v6 = _dafny.MultiSet.fromElements(_59_v0);
        let _66_v7;
        _66_v7 = _dafny.Map.Empty.slice().updateUnsafe(true,_65_v6);
        let _67_v8;
        _67_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_66_v7).length),_60_v1);
        let _nw0 = Array((new BigNumber(14)).toNumber());
        _nw0[(_dafny.ZERO).toNumber()] = _60_v1;
        _nw0[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_61_v2, _61_v2)).length);
        _nw0[(new BigNumber(2)).toNumber()] = _60_v1;
        _nw0[(new BigNumber(3)).toNumber()] = _60_v1;
        _nw0[(new BigNumber(4)).toNumber()] = new BigNumber(-822);
        _nw0[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_64_v5).length));
        _nw0[(new BigNumber(6)).toNumber()] = _60_v1;
        _nw0[(new BigNumber(7)).toNumber()] = _60_v1;
        _nw0[(new BigNumber(8)).toNumber()] = _60_v1;
        _nw0[(new BigNumber(9)).toNumber()] = _60_v1;
        _nw0[(new BigNumber(10)).toNumber()] = new BigNumber(-244);
        _nw0[(new BigNumber(11)).toNumber()] = _60_v1;
        _nw0[(new BigNumber(12)).toNumber()] = _60_v1;
        _nw0[(new BigNumber(13)).toNumber()] = new BigNumber((_67_v8).length);
        r0 = _nw0;
        let _68_v9;
        let _nw1 = Array((new BigNumber(6)).toNumber()).fill([]);
        _68_v9 = _nw1;
        let _69_v10;
        let _init0 = ((_70_v1) => function (_71_i0) {
          return (_71_i0).multipliedBy(_70_v1);
        })(_60_v1);
        let _nw2 = Array((new BigNumber(8)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
          _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _69_v10 = _nw2;
        let _index0 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_68_v9).length));
        (_68_v9)[_index0] = _69_v10;
        let _index1 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_68_v9).length));
        let _nw3 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        (_68_v9)[_index1] = _nw3;
        r3 = _module.__default.safeDivisionInt(_60_v1, ((_59_v0) ? (_60_v1) : (_60_v1)));
        let _72_v11;
        _72_v11 = _dafny.Seq.of(_60_v1);
        let _73_v12;
        _73_v12 = _dafny.Map.Empty.slice().updateUnsafe(_59_v0,(_72_v11)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_72_v11).length))]);
        let _74_v13;
        _74_v13 = _dafny.MultiSet.fromElements(new BigNumber(-256), new BigNumber(656));
        let _rhs0 = (_74_v13).IsProperSubsetOf(_74_v13);
        let _rhs1 = (_73_v12).Merge(_73_v12);
        let _rhs2 = _59_v0;
        let _rhs3 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-981)), ((_75_p0) => function (_76_i1) {
          return _75_p0;
        })(p0));
        let _rhs4 = _60_v1;
        r1 = _rhs0;
        _73_v12 = _rhs1;
        r1 = _rhs2;
        _61_v2 = _rhs3;
        _60_v1 = _rhs4;
        r1 = _59_v0;
      } else {
        let _77_v14;
        _77_v14 = _dafny.ONE;
        let _78_v15;
        _78_v15 = _dafny.Seq.of(_59_v0);
        let _rhs5 = (_dafny.ZERO).minus(_77_v14);
        let _rhs6 = new BigNumber(844);
        let _rhs7 = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Concat(_78_v15, _78_v15)).length)).plus(_77_v14));
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        let _lhs2 = globalState;
        _lhs0.f2 = _rhs5;
        _lhs1.f2 = _rhs6;
        _lhs2.f2 = _rhs7;
        _77_v14 = new BigNumber(148);
        let _79_v16;
        _79_v16 = _dafny.Seq.UnicodeFromString("tuyexyypg");
        let _80_v17;
        let _nw4 = new _module.C2();
        _nw4.__ctor(new BigNumber((_79_v16).length), _77_v14);
        _80_v17 = _nw4;
        (globalState).f2 = _77_v14;
        (globalState).f2 = (_module.__default.safeDivisionInt(_module.__default.fm37(_59_v0, new BigNumber(805), !(_59_v0), globalState), _77_v14)).minus(new BigNumber(491));
      }
      let _81_v18;
      _81_v18 = new BigNumber(-649);
      let _hi0 = _81_v18;
      for (let _82_i2 = _81_v18; _82_i2.isLessThan(_hi0); _82_i2 = _82_i2.plus(_dafny.ONE)) {
        let _83_v19;
        let _nw5 = new _module.C3();
        _nw5.__ctor(false, (_82_i2).multipliedBy(_81_v18), _module.__default.safeModuloInt(_81_v18, _82_i2));
        _83_v19 = _nw5;
        let _nw6 = new _module.C1();
        _nw6.__ctor(_82_i2, (_82_i2).multipliedBy(new BigNumber(-405)));
        _83_v19 = _nw6;
        let _84_v20;
        let _nw7 = new _module.C2();
        _nw7.__ctor(_81_v18, _83_v19.f6);
        _84_v20 = _nw7;
        let _85_v21;
        _85_v21 = _dafny.Seq.UnicodeFromString("krfj");
        _85_v21 = _dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm7(p0, true, _59_v0, _81_v18, globalState), _dafny.Seq.Concat(_85_v21, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("bwr"), _module.__default.safeIndex(_82_i2, new BigNumber((_dafny.Seq.UnicodeFromString("bwr")).length)), p0))), _module.__default.safeIndex(new BigNumber(446), new BigNumber((_dafny.Seq.Concat(_module.__default.fm7(p0, true, _59_v0, _81_v18, globalState), _dafny.Seq.Concat(_85_v21, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("bwr"), _module.__default.safeIndex(_82_i2, new BigNumber((_dafny.Seq.UnicodeFromString("bwr")).length)), p0)))).length)), (_85_v21)[_module.__default.safeIndex(_82_i2, new BigNumber((_85_v21).length))]);
        let _86_v22;
        let _init1 = ((_87_v0) => function (_88_i3) {
          return _87_v0;
        })(_59_v0);
        let _nw8 = Array((new BigNumber(27)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw8.length); _i0_1++) {
          _nw8[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _86_v22 = _nw8;
        let _89_v23;
        _89_v23 = _module.D7.create_DC22(_86_v22, _85_v21);
        _59_v0 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("vuhemdl"), (_89_v23).dtor_cf39);
      }
      let _90_v24;
      let _nw9 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _90_v24 = _nw9;
      let _91_v25;
      _91_v25 = _module.D5.create_DC16(_90_v24);
      let _92_v26;
      _92_v26 = _dafny.Map.Empty.slice().updateUnsafe(_91_v25,_81_v18);
      let _93_v27;
      _93_v27 = _dafny.MultiSet.fromElements(_92_v26, _dafny.Map.Empty.slice().updateUnsafe(_91_v25,_81_v18), _dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC16(_90_v24),_81_v18));
      _93_v27 = (_93_v27).Intersect(_93_v27);
      let _94_v28;
      _94_v28 = _dafny.Seq.of(!(_59_v0));
      let _95_v29;
      _95_v29 = _dafny.Set.fromElements(_94_v28);
      let _96_v30;
      _96_v30 = _dafny.MultiSet.fromElements(_59_v0);
      let _hi1 = new BigNumber((_96_v30).cardinality());
      for (let _97_i4 = new BigNumber(((_95_v29).Union(_95_v29)).length); _97_i4.isLessThan(_hi1); _97_i4 = _97_i4.plus(_dafny.ONE)) {
        let _98_v31;
        _98_v31 = _module.D4.create_DC14(_81_v18, _59_v0, _97_i4, _59_v0, _97_i4);
        _96_v30 = _dafny.MultiSet.fromElements(!((function (_pat_let0_0) {
          return function (_99_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_100_dt__update_hcf20_h0) {
                return function (_pat_let2_0) {
                  return function (_101_dt__update_hcf24_h0) {
                    return _module.D4.create_DC14(_100_dt__update_hcf20_h0, (_99_dt__update__tmp_h0).dtor_cf21, (_99_dt__update__tmp_h0).dtor_cf22, (_99_dt__update__tmp_h0).dtor_cf23, _101_dt__update_hcf24_h0);
                  }(_pat_let2_0);
                }(_97_i4);
              }(_pat_let1_0);
            }(_97_i4);
          }(_pat_let0_0);
        }(_98_v31)).dtor_cf21));
        (globalState).f2 = _81_v18;
        let _102_v32;
        _102_v32 = _dafny.Map.Empty.slice().updateUnsafe(_97_i4,_59_v0);
        let _103_v33;
        _103_v33 = _dafny.Seq.of(_97_i4, _81_v18);
        r1 = _module.__default.fm38(_102_v32, _dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_103_v33).length)), _97_i4), true, (_97_i4).isLessThan(_81_v18), globalState);
        let _104_v34;
        let _nw10 = new _module.C1();
        _nw10.__ctor(new BigNumber(-663), _module.__default.fm37(_59_v0, _81_v18, _59_v0, globalState));
        _104_v34 = _nw10;
      }
      let _105_v35;
      _105_v35 = _dafny.Seq.UnicodeFromString("oyt");
      _105_v35 = _dafny.Seq.Concat(_105_v35, _105_v35);
      let _106_v36;
      _106_v36 = _module.D0.create_DC0(_59_v0);
      let _107_v37;
      _107_v37 = _module.D10.create_DC29(_dafny.Map.Empty.slice().updateUnsafe(_81_v18,new BigNumber((_105_v35).length)));
      let _rhs8 = _106_v36;
      let _rhs9 = _107_v37;
      let _rhs10 = _59_v0;
      let _rhs11 = _59_v0;
      _106_v36 = _rhs8;
      _107_v37 = _rhs9;
      r1 = _rhs10;
      _59_v0 = _rhs11;
      let _108_v38;
      _108_v38 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_105_v35, _105_v35),_90_v24);
      r0 = (((_108_v38).contains(_dafny.Seq.Concat(_105_v35, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("kcykgfojv"), _module.__default.safeIndex(new BigNumber(-354), new BigNumber((_dafny.Seq.UnicodeFromString("kcykgfojv")).length)), new _dafny.CodePoint('r'.codePointAt(0)))))) ? ((_108_v38).get(_dafny.Seq.Concat(_105_v35, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("kcykgfojv"), _module.__default.safeIndex(new BigNumber(-354), new BigNumber((_dafny.Seq.UnicodeFromString("kcykgfojv")).length)), new _dafny.CodePoint('r'.codePointAt(0)))))) : (_90_v24));
      r1 = _59_v0;
      let _109_v39;
      _109_v39 = _dafny.MultiSet.fromElements(_105_v35, _105_v35);
      r2 = _109_v39;
      let _110_v40;
      _110_v40 = _dafny.MultiSet.fromElements(_81_v18);
      let _111_v41;
      _111_v41 = _dafny.Seq.of(new BigNumber(747), (_dafny.ZERO).minus((((_110_v40).contains(_81_v18)) ? ((_110_v40).get(_81_v18)) : (_81_v18))));
      r3 = new BigNumber((_111_v41).length);
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _112_v0;
      _112_v0 = new BigNumber(-723);
      let _113_v1;
      _113_v1 = new _dafny.CodePoint('i'.codePointAt(0));
      let _114_v2;
      _114_v2 = true;
      let _115_v3;
      _115_v3 = _dafny.Map.Empty.slice().updateUnsafe(_113_v1,_114_v2);
      let _116_v4;
      _116_v4 = _dafny.MultiSet.fromElements(_112_v0, _112_v0, _112_v0);
      let _117_v5;
      _117_v5 = _dafny.Map.Empty.slice().updateUnsafe(_113_v1,_116_v4);
      let _118_globalState;
      let _nw11 = new _module.GlobalState();
      _nw11.__ctor(new BigNumber(275), new BigNumber(628), new BigNumber(-930), new BigNumber(254), (_dafny.MultiSet.fromElements(_112_v0, (_dafny.ZERO).minus(new BigNumber((_115_v3).length)), _112_v0, _112_v0, _112_v0)).Union((((_117_v5).contains(_113_v1)) ? ((_117_v5).get(_113_v1)) : (_116_v4))));
      _118_globalState = _nw11;
      if (_114_v2) {
        let _119_v6;
        let _120_v7;
        let _121_v8;
        let _122_v9;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = _module.__default.m0(_113_v1, _118_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _119_v6 = _out0;
        _120_v7 = _out1;
        _121_v8 = _out2;
        _122_v9 = _out3;
        let _123_v10;
        let _nw12 = new _module.C3();
        _nw12.__ctor(_120_v7, _112_v0, _122_v9);
        _123_v10 = _nw12;
        let _124_v11;
        let _nw13 = new _module.C1();
        _nw13.__ctor(_module.__default.safeModuloInt(_122_v9, _122_v9), _112_v0);
        _124_v11 = _nw13;
        let _nw14 = new _module.C1();
        _nw14.__ctor(_module.__default.safeDivisionInt(_122_v9, (_dafny.ZERO).minus(_112_v0)), _112_v0);
        _124_v11 = _nw14;
        let _125_v12;
        _125_v12 = _dafny.Seq.UnicodeFromString("a");
        let _126_v13;
        _126_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_125_v12, _125_v12),_112_v0);
        let _127_v14;
        _127_v14 = _dafny.MultiSet.fromElements(_124_v11);
        _126_v13 = (_126_v13).update((!((_123_v10).f7)) || (_114_v2), (((_127_v14).contains(_124_v11)) ? ((_127_v14).get(_124_v11)) : (_112_v0)));
        _120_v7 = !(_120_v7);
      } else {
        let _128_v15;
        _128_v15 = _dafny.Seq.UnicodeFromString("mw");
        let _129_v16;
        _129_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(933),_112_v0);
        let _130_v17;
        _130_v17 = _dafny.Seq.of(_114_v2);
        let _131_v18;
        _131_v18 = _dafny.MultiSet.fromElements(_114_v2, _114_v2);
        let _132_v19;
        _132_v19 = _131_v18;
        let _133_v20;
        let _nw15 = Array((new BigNumber(21)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = new BigNumber((_128_v15).length);
        _nw15[(_dafny.ONE).toNumber()] = new BigNumber((_129_v16).length);
        _nw15[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(_112_v0, _112_v0);
        _nw15[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_112_v0, new BigNumber((_dafny.Set.fromElements(_112_v0)).length));
        _nw15[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_130_v17, _130_v17)).length);
        _nw15[(new BigNumber(5)).toNumber()] = _112_v0;
        _nw15[(new BigNumber(6)).toNumber()] = new BigNumber(-903);
        _nw15[(new BigNumber(7)).toNumber()] = (_112_v0).plus(new BigNumber((_116_v4).cardinality()));
        _nw15[(new BigNumber(8)).toNumber()] = _112_v0;
        _nw15[(new BigNumber(9)).toNumber()] = _112_v0;
        _nw15[(new BigNumber(10)).toNumber()] = _112_v0;
        _nw15[(new BigNumber(11)).toNumber()] = new BigNumber((_128_v15).length);
        _nw15[(new BigNumber(12)).toNumber()] = _112_v0;
        _nw15[(new BigNumber(13)).toNumber()] = new BigNumber((_module.__default.fm35(true, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(813)), ((_134_v0) => function (_135_i0) {
          return _134_v0;
        })(_112_v0))).length), _112_v0, _112_v0, _118_globalState)).length);
        _nw15[(new BigNumber(14)).toNumber()] = new BigNumber(((_132_v19)).cardinality());
        _nw15[(new BigNumber(15)).toNumber()] = new BigNumber(281);
        _nw15[(new BigNumber(16)).toNumber()] = _112_v0;
        _nw15[(new BigNumber(17)).toNumber()] = _112_v0;
        _nw15[(new BigNumber(18)).toNumber()] = new BigNumber(-225);
        _nw15[(new BigNumber(19)).toNumber()] = new BigNumber((_130_v17).length);
        _nw15[(new BigNumber(20)).toNumber()] = _112_v0;
        _133_v20 = _nw15;
        let _index2 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length));
        (_133_v20)[_index2] = _112_v0;
        let _136_v21;
        _136_v21 = _module.D0.create_DC2();
        let _137_v22;
        _137_v22 = _module.D3.create_DC9(_116_v4, _136_v21, _112_v0);
        let _138_v23;
        let _nw16 = Array((_dafny.ONE).toNumber());
        _nw16[(_dafny.ZERO).toNumber()] = _137_v22;
        _138_v23 = _nw16;
        let _pat_let_tv0 = _136_v21;
        let _index3 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_138_v23).length));
        (_138_v23)[_index3] = function (_pat_let3_0) {
          return function (_139_dt__update__tmp_h0) {
            return function (_pat_let4_0) {
              return function (_140_dt__update_hcf10_h0) {
                return _module.D3.create_DC9((_139_dt__update__tmp_h0).dtor_cf9, _140_dt__update_hcf10_h0, (_139_dt__update__tmp_h0).dtor_cf11);
              }(_pat_let4_0);
            }(_pat_let_tv0);
          }(_pat_let3_0);
        }(_module.__default.fm36(_module.D5.create_DC19(_112_v0, _113_v1, _112_v0), _113_v1, _114_v2, _118_globalState));
        let _index4 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length));
        let _index5 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_138_v23).length));
        let _rhs12 = !(_114_v2);
        let _rhs13 = ((new BigNumber(524)).minus(new BigNumber((_dafny.Seq.update(_128_v15, _module.__default.safeIndex(_112_v0, new BigNumber((_128_v15).length)), _113_v1)).length))).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_130_v17).length))));
        let _rhs14 = _137_v22;
        let _lhs3 = _133_v20;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length));
        let _lhs5 = _138_v23;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_138_v23).length));
        _114_v2 = _rhs12;
        _lhs3[_lhs4] = _rhs13;
        _lhs5[_lhs6] = _rhs14;
        let _141_v24;
        _141_v24 = _module.D7.create_DC23((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))], _128_v15, new BigNumber((_129_v16).length));
        let _142_v25;
        _142_v25 = _dafny.MultiSet.fromElements(_128_v15, (_141_v24).dtor_cf41, _128_v15);
        if ((_142_v25).equals(_142_v25)) {
          let _143_v26;
          let _144_v27;
          let _145_v28;
          let _146_v29;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = _module.__default.m0(_113_v1, _118_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _143_v26 = _out4;
          _144_v27 = _out5;
          _145_v28 = _out6;
          _146_v29 = _out7;
          let _147_v30;
          let _nw17 = new _module.C3();
          _nw17.__ctor(_144_v27, _112_v0, _112_v0);
          _147_v30 = _nw17;
          let _148_v31;
          _148_v31 = _dafny.Map.Empty.slice().updateUnsafe(((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))]).plus(new BigNumber(608)),_147_v30);
          let _149_v32;
          let _nw18 = Array((new BigNumber(3)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = (_147_v30).f7;
          _nw18[(_dafny.ONE).toNumber()] = (((_115_v3).contains(_113_v1)) ? ((_115_v3).get(_113_v1)) : ((_147_v30).f7));
          _nw18[(new BigNumber(2)).toNumber()] = _144_v27;
          _149_v32 = _nw18;
          let _index6 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_149_v32).length));
          (_149_v32)[_index6] = !(_114_v2);
          let _150_v33;
          _150_v33 = _dafny.Map.Empty.slice().updateUnsafe(_113_v1,_128_v15);
          let _151_v34;
          _151_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_150_v33).contains(_113_v1)) ? ((_150_v33).get(_113_v1)) : (_128_v15))).length),(_147_v30).f7);
          let _152_v35;
          _152_v35 = _dafny.MultiSet.fromElements(_151_v34, _151_v34);
          let _153_v36;
          _153_v36 = _dafny.Set.fromElements(_114_v2);
          let _154_v37;
          _154_v37 = _dafny.Seq.of((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))], new BigNumber((_dafny.Set.fromElements(true, !(_114_v2))).length));
          let _index7 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_149_v32).length));
          let _rhs15 = _148_v31;
          let _rhs16 = (_147_v30).fm0((_152_v35).Difference(_152_v35), _114_v2, _128_v15, _118_globalState);
          let _rhs17 = (new BigNumber((_154_v37).length)).isLessThanOrEqualTo((_147_v30).fm2(new BigNumber(209), _153_v36, _118_globalState));
          let _rhs18 = _146_v29;
          let _lhs7 = _149_v32;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_149_v32).length));
          let _lhs9 = _118_globalState;
          _148_v31 = _rhs15;
          _lhs7[_lhs8] = _rhs16;
          _144_v27 = _rhs17;
          _lhs9.f2 = _rhs18;
          let _155_v38;
          let _nw19 = new _module.C3();
          _nw19.__ctor(_114_v2, (_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))], _112_v0);
          _155_v38 = _nw19;
          let _156_v39;
          let _nw20 = new _module.C1();
          _nw20.__ctor((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))], (_dafny.ZERO).minus((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))]));
          _156_v39 = _nw20;
          let _157_v40;
          _157_v40 = _dafny.Seq.of(_149_v32, _149_v32, _149_v32, _149_v32);
          let _158_v41;
          _158_v41 = _module.D7.create_DC22(_149_v32, _128_v15);
          let _159_v42;
          let _nw21 = Array((new BigNumber(13)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = (_157_v40)[_module.__default.safeIndex((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))], new BigNumber((_157_v40).length))];
          _nw21[(_dafny.ONE).toNumber()] = _149_v32;
          _nw21[(new BigNumber(2)).toNumber()] = _149_v32;
          _nw21[(new BigNumber(3)).toNumber()] = _149_v32;
          _nw21[(new BigNumber(4)).toNumber()] = (_157_v40)[_module.__default.safeIndex((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))], new BigNumber((_157_v40).length))];
          _nw21[(new BigNumber(5)).toNumber()] = _149_v32;
          _nw21[(new BigNumber(6)).toNumber()] = (_158_v41).dtor_cf38;
          _nw21[(new BigNumber(7)).toNumber()] = _149_v32;
          _nw21[(new BigNumber(8)).toNumber()] = _149_v32;
          _nw21[(new BigNumber(9)).toNumber()] = (_158_v41).dtor_cf38;
          _nw21[(new BigNumber(10)).toNumber()] = _149_v32;
          _nw21[(new BigNumber(11)).toNumber()] = _149_v32;
          _nw21[(new BigNumber(12)).toNumber()] = _149_v32;
          _159_v42 = _nw21;
          let _160_v43;
          _160_v43 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(167)), ((_161_v1) => function (_162_i1) {
            return _161_v1;
          })(_113_v1)));
          let _index8 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length));
          let _rhs19 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_146_v29, new BigNumber((_dafny.Seq.of(_114_v2, (_155_v38).f7, _114_v2, (_155_v38).f7, true)).length)));
          let _rhs20 = (_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))];
          let _rhs21 = _159_v42;
          let _rhs22 = new BigNumber(((_160_v43)[_module.__default.safeIndex((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))], new BigNumber((_160_v43).length))]).length);
          let _lhs10 = _133_v20;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length));
          let _lhs12 = _118_globalState;
          _lhs10[_lhs11] = _rhs19;
          _lhs12.f2 = _rhs20;
          _159_v42 = _rhs21;
          _112_v0 = _rhs22;
        } else {
          (_118_globalState).f2 = _112_v0;
          let _163_v44;
          let _nw22 = new _module.C3();
          _nw22.__ctor(_114_v2, ((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))]).plus(new BigNumber(11)), (_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))]);
          _163_v44 = _nw22;
          let _164_v45;
          _164_v45 = _163_v44;
          _163_v44 = (_164_v45);
          let _165_v46;
          let _nw23 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
          _165_v46 = _nw23;
          let _index9 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_165_v46).length));
          (_165_v46)[_index9] = _131_v18;
          let _index10 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_165_v46).length));
          let _rhs23 = (_dafny.ZERO).minus((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))]);
          let _rhs24 = (_131_v18).Union(_131_v18);
          let _lhs13 = _165_v46;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_165_v46).length));
          _112_v0 = _rhs23;
          _lhs13[_lhs14] = _rhs24;
          let _166_v47;
          let _init2 = ((_167_v44) => function (_168_i2) {
            return (_167_v44).f7;
          })(_163_v44);
          let _nw24 = Array((new BigNumber(10)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw24.length); _i0_2++) {
            _nw24[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _166_v47 = _nw24;
          let _169_v48;
          _169_v48 = _dafny.Map.Empty.slice().updateUnsafe(_113_v1,_166_v47);
          let _170_v49;
          _170_v49 = _dafny.Map.Empty.slice().updateUnsafe((((_169_v48).contains(_113_v1)) ? ((_169_v48).get(_113_v1)) : (_166_v47)),(_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))]);
          _170_v49 = (_170_v49).Merge((_170_v49).Merge(_dafny.Map.Empty.slice().updateUnsafe(_166_v47,new BigNumber((_128_v15).length))));
          let _index11 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length));
          (_133_v20)[_index11] = (_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))];
        }
        let _171_v50;
        let _nw25 = new _module.C0();
        _nw25.__ctor(new BigNumber(-984));
        _171_v50 = _nw25;
        let _172_v51;
        _172_v51 = _dafny.MultiSet.fromElements(_171_v50);
        let _173_v52;
        let _nw26 = new _module.C1();
        _nw26.__ctor(((((_172_v51).contains(_171_v50)) ? ((_172_v51).get(_171_v50)) : (new BigNumber((_128_v15).length)))).multipliedBy((_171_v50).f8), ((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))]).plus(_112_v0));
        _173_v52 = _nw26;
        _173_v52 = _173_v52;
        let _174_v53;
        _174_v53 = _dafny.Seq.of(_112_v0, (((_129_v16).contains((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))])) ? ((_129_v16).get((_133_v20)[_module.__default.safeIndex(new BigNumber(190), new BigNumber((_133_v20).length))])) : ((_173_v52).f5)), (_173_v52).f5);
        let _175_v54;
        let _nw27 = new _module.C3();
        _nw27.__ctor(_114_v2, new BigNumber((_174_v53).length), (_171_v50).f8);
        _175_v54 = _nw27;
        _128_v15 = _128_v15;
      }
      let _176_i3;
      _176_i3 = _dafny.ZERO;
      L0: {
        while (_114_v2) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_176_i3)) {
              break L0;
            }
            _176_i3 = (_176_i3).plus(_dafny.ONE);
            let _rhs25 = ((_114_v2) ? (_114_v2) : (!((_116_v4).IsProperSubsetOf(_116_v4))));
            let _rhs26 = _114_v2;
            let _rhs27 = _114_v2;
            _114_v2 = _rhs25;
            _114_v2 = _rhs26;
            _114_v2 = _rhs27;
            let _177_v55;
            _177_v55 = _dafny.MultiSet.fromElements(_114_v2, _114_v2);
            let _178_v56;
            _178_v56 = _dafny.Map.Empty.slice().updateUnsafe(_114_v2,_112_v0);
            (_118_globalState).f2 = (((_177_v55).contains((((_115_v3).contains(_113_v1)) ? ((_115_v3).get(_113_v1)) : (_114_v2)))) ? ((_177_v55).get((((_115_v3).contains(_113_v1)) ? ((_115_v3).get(_113_v1)) : (_114_v2)))) : (((((_178_v56).contains(_114_v2)) ? ((_178_v56).get(_114_v2)) : (_112_v0))).plus(_112_v0)));
            _114_v2 = !(_114_v2);
            let _179_v57;
            _179_v57 = _dafny.Seq.UnicodeFromString("lqy");
            let _180_v58;
            _180_v58 = _dafny.Map.Empty.slice().updateUnsafe(_114_v2,_113_v1);
            let _181_v59;
            _181_v59 = _dafny.Seq.of(_114_v2);
            _179_v57 = _module.__default.fm7((((_180_v58).contains(_114_v2)) ? ((_180_v58).get(_114_v2)) : (_113_v1)), _114_v2, (new BigNumber((_181_v59).length)).isEqualTo((_dafny.ZERO).minus(_112_v0)), _module.__default.safeModuloInt((_dafny.ZERO).minus(_112_v0), _112_v0), _118_globalState);
          }
        }
      }
      let _182_v60;
      _182_v60 = _dafny.Seq.of(false);
      let _183_v61;
      _183_v61 = _dafny.Set.fromElements(!((_182_v60)[_module.__default.safeIndex(_112_v0, new BigNumber((_182_v60).length))]));
      let _184_v62;
      _184_v62 = _dafny.Seq.UnicodeFromString("uapcj");
      let _185_v63;
      let _nw28 = Array((new BigNumber(9)).toNumber());
      _nw28[(_dafny.ZERO).toNumber()] = ((_114_v2) ? (_112_v0) : ((_dafny.ZERO).minus(_112_v0)));
      _nw28[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(((_dafny.ZERO).minus(_112_v0)).multipliedBy(_112_v0));
      _nw28[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(804)), ((_186_v1) => function (_187_i5) {
        return _186_v1;
      })(_113_v1))).length);
      _nw28[(new BigNumber(3)).toNumber()] = _112_v0;
      _nw28[(new BigNumber(4)).toNumber()] = new BigNumber((_183_v61).length);
      _nw28[(new BigNumber(5)).toNumber()] = _112_v0;
      _nw28[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_184_v62, _dafny.Seq.UnicodeFromString("qmhnk"))).length);
      _nw28[(new BigNumber(7)).toNumber()] = new BigNumber(((_183_v61).Intersect(_183_v61)).length);
      _nw28[(new BigNumber(8)).toNumber()] = _112_v0;
      _185_v63 = _nw28;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_185_v63).length))) {
        let _188_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_188_i4)) && ((_188_i4).isLessThan(new BigNumber((_185_v63).length))))) {
          (_185_v63)[(_188_i4)] = (_188_i4).plus(new BigNumber((_182_v60).length));
        }
      }
      let _189_v64;
      _189_v64 = _module.D1.create_DC6(_184_v62, (_dafny.ZERO).minus(new BigNumber((((_114_v2) ? (_183_v61) : (_183_v61))).length)));
      let _source3 = _189_v64;
      if (_source3.is_DC5) {
        let _rhs28 = _113_v1;
        let _rhs29 = _114_v2;
        _113_v1 = _rhs28;
        _114_v2 = _rhs29;
        _112_v0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-266), (_112_v0).plus(_112_v0)));
        let _190_v65;
        let _191_v66;
        let _192_v67;
        let _193_v68;
        let _out8;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector2 = _module.__default.m0(_113_v1, _118_globalState);
        _out8 = _outcollector2[0];
        _out9 = _outcollector2[1];
        _out10 = _outcollector2[2];
        _out11 = _outcollector2[3];
        _190_v65 = _out8;
        _191_v66 = _out9;
        _192_v67 = _out10;
        _193_v68 = _out11;
        _191_v66 = true;
      } else if (_source3.is_DC6) {
        let _194___mcc_h0 = (_source3).cf5;
        let _195___mcc_h1 = (_source3).cf6;
        let _196_cf6 = _195___mcc_h1;
        let _197_cf5 = _194___mcc_h0;
        _114_v2 = ((_114_v2) ? ((_183_v61).IsSubsetOf(_183_v61)) : (false));
        let _198_v69;
        let _nw29 = new _module.C1();
        _nw29.__ctor(new BigNumber(((_dafny.Set.fromElements(_module.__default.fm37(_114_v2, _112_v0, _114_v2, _118_globalState))).Difference(_dafny.Set.fromElements(_112_v0, _112_v0, _112_v0, new BigNumber(678), _196_cf6))).length), _196_cf6);
        _198_v69 = _nw29;
        _198_v69 = _198_v69;
        let _199_v70;
        _199_v70 = _module.D9.create_DC28(_114_v2, _196_cf6, _112_v0);
        _114_v2 = !((_199_v70).dtor_cf46);
        let _200_v71;
        let _nw30 = Array((new BigNumber(9)).toNumber()).fill(false);
        _200_v71 = _nw30;
        let _index12 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_200_v71).length));
        (_200_v71)[_index12] = _114_v2;
        let _index13 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_200_v71).length));
        (_200_v71)[_index13] = _114_v2;
      } else {
        let _201___mcc_h2 = (_source3).cf4;
        let _202_cf4 = _201___mcc_h2;
        if (((_114_v2) ? ((_112_v0).isLessThanOrEqualTo(_112_v0)) : (_114_v2))) {
          let _203_v72;
          _203_v72 = _dafny.Seq.of(new BigNumber((_module.__default.fm7(_113_v1, _114_v2, _114_v2, new BigNumber(373), _118_globalState)).length));
          _203_v72 = _203_v72;
          let _204_v73;
          _204_v73 = _dafny.Map.Empty.slice().updateUnsafe(_184_v62,_114_v2);
          _204_v73 = (_204_v73).update(_dafny.Seq.UnicodeFromString("xajwiulrn"), _114_v2);
          _114_v2 = false;
          (_118_globalState).f2 = _112_v0;
          let _205_v74;
          let _nw31 = Array((new BigNumber(26)).toNumber());
          _205_v74 = _nw31;
          let _206_v75;
          _206_v75 = _dafny.Map.Empty.slice().updateUnsafe(_112_v0,true);
          let _207_v76;
          _207_v76 = _dafny.MultiSet.fromElements((((_206_v75).contains(new BigNumber((_202_cf4).length))) ? ((_206_v75).get(new BigNumber((_202_cf4).length))) : (true)));
          let _208_v77;
          _208_v77 = _dafny.Map.Empty.slice().updateUnsafe(_114_v2,_114_v2);
          let _209_v78;
          _209_v78 = _dafny.Map.Empty.slice().updateUnsafe(false,(((_207_v76).contains(_114_v2)) ? ((_207_v76).get(_114_v2)) : (new BigNumber((_208_v77).length))));
          let _210_v79;
          _210_v79 = (_209_v78).Merge(_209_v78);
          let _rhs30 = _205_v74;
          let _rhs31 = (_dafny.ZERO).minus(_112_v0);
          let _rhs32 = _209_v78;
          let _rhs33 = (_module.__default.safeDivisionInt(_112_v0, _112_v0)).isEqualTo(_112_v0);
          let _lhs15 = _118_globalState;
          _205_v74 = _rhs30;
          _lhs15.f2 = _rhs31;
          _210_v79 = _rhs32;
          _114_v2 = _rhs33;
        } else {
          _114_v2 = _114_v2;
          let _211_v80;
          _211_v80 = _dafny.Seq.of(_112_v0);
          let _212_v81;
          _212_v81 = _dafny.Map.Empty.slice().updateUnsafe(_211_v80,_112_v0);
          let _213_v82;
          _213_v82 = _dafny.Map.Empty.slice().updateUnsafe(_184_v62,_112_v0);
          _112_v0 = new BigNumber(((_212_v81).update(_dafny.Seq.Concat(_211_v80, _dafny.Seq.of(_112_v0)), (((_213_v82).contains(_dafny.Seq.update(_184_v62, _module.__default.safeIndex(_112_v0, new BigNumber((_184_v62).length)), _113_v1))) ? ((_213_v82).get(_dafny.Seq.update(_184_v62, _module.__default.safeIndex(_112_v0, new BigNumber((_184_v62).length)), _113_v1))) : (_112_v0)))).length);
          _182_v60 = _182_v60;
          let _214_v83;
          let _nw32 = new _module.C1();
          _nw32.__ctor(new BigNumber(896), new BigNumber(501));
          _214_v83 = _nw32;
          let _215_v84;
          _215_v84 = _dafny.Map.Empty.slice().updateUnsafe(_112_v0,_214_v83);
          let _216_v85;
          let _nw33 = Array((new BigNumber(6)).toNumber());
          _nw33[(_dafny.ZERO).toNumber()] = _214_v83;
          _nw33[(_dafny.ONE).toNumber()] = (((_215_v84).contains(_112_v0)) ? ((_215_v84).get(_112_v0)) : (_214_v83));
          _nw33[(new BigNumber(2)).toNumber()] = _214_v83;
          _nw33[(new BigNumber(3)).toNumber()] = _214_v83;
          _nw33[(new BigNumber(4)).toNumber()] = _214_v83;
          _nw33[(new BigNumber(5)).toNumber()] = _214_v83;
          _216_v85 = _nw33;
          let _index14 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_216_v85).length));
          (_216_v85)[_index14] = _214_v83;
          let _index15 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_216_v85).length));
          (_216_v85)[_index15] = _214_v83;
          let _217_v86;
          _217_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(602),_114_v2);
          let _218_v87;
          _218_v87 = _dafny.MultiSet.fromElements(_217_v86);
          (_118_globalState).f2 = (_module.__default.safeModuloInt(_112_v0, _112_v0)).plus(_module.__default.fm37(_114_v2, _112_v0, (_214_v83).fm0(_218_v87, _114_v2, _184_v62, _118_globalState), _118_globalState));
        }
        let _219_v88;
        let _220_v89;
        let _221_v90;
        let _222_v91;
        let _out12;
        let _out13;
        let _out14;
        let _out15;
        let _outcollector3 = _module.__default.m0(_113_v1, _118_globalState);
        _out12 = _outcollector3[0];
        _out13 = _outcollector3[1];
        _out14 = _outcollector3[2];
        _out15 = _outcollector3[3];
        _219_v88 = _out12;
        _220_v89 = _out13;
        _221_v90 = _out14;
        _222_v91 = _out15;
        let _223_v92;
        let _init3 = ((_224_v89) => function (_225_i6) {
          return _224_v89;
        })(_220_v89);
        let _nw34 = Array((new BigNumber(29)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw34.length); _i0_3++) {
          _nw34[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _223_v92 = _nw34;
        let _226_v93;
        _226_v93 = _module.D7.create_DC22(_223_v92, _dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), ((_227_v1) => function (_228_i7) {
  return _227_v1;
})(_113_v1)));
        let _229_v94;
        _229_v94 = _dafny.Seq.of(_223_v92, (_226_v93).dtor_cf38, _223_v92);
        (_118_globalState).f2 = new BigNumber((((_114_v2) ? (_229_v94) : (_229_v94))).length);
        let _230_v95;
        let _nw35 = new _module.C1();
        _nw35.__ctor(new BigNumber(220), _222_v91);
        _230_v95 = _nw35;
        _230_v95 = _230_v95;
      }
      _112_v0 = ((_dafny.ZERO).minus(_112_v0)).plus(_112_v0);
      let _231_v96;
      _231_v96 = _dafny.Map.Empty.slice().updateUnsafe(!(_114_v2),_114_v2);
      let _232_v97;
      _232_v97 = _dafny.Seq.of(new BigNumber((_116_v4).cardinality()));
      let _233_v98;
      _233_v98 = _dafny.Map.Empty.slice().updateUnsafe((_232_v97)[_module.__default.safeIndex(_112_v0, new BigNumber((_232_v97).length))],_114_v2);
      let _234_v99;
      _234_v99 = _dafny.MultiSet.fromElements(_114_v2);
      let _235_i8;
      _235_i8 = _dafny.ZERO;
      L1: {
        let _pat_let_tv1 = _114_v2;
        let _pat_let_tv2 = _114_v2;
        let _pat_let_tv3 = _114_v2;
        let _pat_let_tv4 = _114_v2;
        while (function (_source4) {
          if (_source4.is_DC5) {
            return _pat_let_tv1;
          } else if (_source4.is_DC6) {
            let _241___mcc_h3 = (_source4).cf5;
            let _242___mcc_h4 = (_source4).cf6;
            let _243_cf6 = _242___mcc_h4;
            let _244_cf5 = _241___mcc_h3;
            return _pat_let_tv2;
          } else {
            let _245___mcc_h5 = (_source4).cf4;
            let _246_cf4 = _245___mcc_h5;
            return !(_pat_let_tv3) || (_pat_let_tv4);
          }
        }(_module.__default.fm3(_module.__default.fm37(false, new BigNumber((_231_v96).length), _module.__default.fm38(_233_v98, _114_v2, _114_v2, _114_v2, _118_globalState), _118_globalState), new BigNumber(377), _dafny.Seq.of(_234_v99), _dafny.Seq.update(_232_v97, _module.__default.safeIndex(_112_v0, new BigNumber((_232_v97).length)), new BigNumber((_dafny.Seq.UnicodeFromString("daff")).length)), _118_globalState))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_235_i8)) {
              break L1;
            }
            _235_i8 = (_235_i8).plus(_dafny.ONE);
            let _236_v100;
            let _237_v101;
            let _238_v102;
            let _239_v103;
            let _out16;
            let _out17;
            let _out18;
            let _out19;
            let _outcollector4 = _module.__default.m0(_113_v1, _118_globalState);
            _out16 = _outcollector4[0];
            _out17 = _outcollector4[1];
            _out18 = _outcollector4[2];
            _out19 = _outcollector4[3];
            _236_v100 = _out16;
            _237_v101 = _out17;
            _238_v102 = _out18;
            _239_v103 = _out19;
            (_118_globalState).f2 = (((_239_v103).isLessThan(_239_v103)) ? (_239_v103) : (new BigNumber((_234_v99).cardinality())));
            let _240_v104;
            let _nw36 = new _module.C1();
            _nw36.__ctor(new BigNumber(623), (_112_v0).multipliedBy(_112_v0));
            _240_v104 = _nw36;
            _185_v63 = _236_v100;
          }
        }
      }
      let _247_v105;
      _247_v105 = _dafny.Seq.of(_116_v4);
      let _hi2 = new BigNumber((_184_v62).length);
      for (let _248_i9 = new BigNumber((_dafny.Seq.Concat(_247_v105, _247_v105)).length); _248_i9.isLessThan(_hi2); _248_i9 = _248_i9.plus(_dafny.ONE)) {
        _114_v2 = _114_v2;
        let _249_v107;
        _249_v107 = _dafny.Set.fromElements(_module.__default.fm37(_114_v2, new BigNumber((_184_v62).length), _114_v2, _118_globalState));
        _233_v98 = ((_114_v2) ? ((function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of (_249_v107).Elements) {
            let _250_v106 = _compr_18;
            if ((_249_v107).contains(_250_v106)) {
              _coll18.push([_module.__default.safeDivisionInt(_250_v106, _112_v0),_114_v2]);
            }
          }
          return _coll18;
        }()).Merge(_233_v98)) : (_dafny.Map.Empty.slice().updateUnsafe(_112_v0,_114_v2)));
        let _251_v108;
        _251_v108 = _dafny.Map.Empty.slice().updateUnsafe(_182_v60,_112_v0);
        _251_v108 = (_251_v108).update(_182_v60, _112_v0);
        _114_v2 = _114_v2;
      }
      let _252_v109;
      let _nw37 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
      _252_v109 = _nw37;
      let _253_v110;
      _253_v110 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_112_v0),_module.__default.fm37(_114_v2, new BigNumber((_231_v96).length), _114_v2, _118_globalState));
      let _index16 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_252_v109).length));
      (_252_v109)[_index16] = (_253_v110).update(_112_v0, _112_v0);
      let _index17 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_252_v109).length));
      (_252_v109)[_index17] = ((!(_114_v2) || (_114_v2)) ? ((_253_v110).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(536),_112_v0))) : (_253_v110));
      let _index18 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
      (_185_v63)[_index18] = (new BigNumber(-644)).plus(_112_v0);
      let _index19 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
      let _rhs34 = !(true);
      let _rhs35 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt(_112_v0, _112_v0))).multipliedBy((_112_v0).plus(new BigNumber((_184_v62).length)));
      let _rhs36 = (_dafny.ZERO).minus((_112_v0).minus(_112_v0));
      let _rhs37 = _112_v0;
      let _lhs16 = _118_globalState;
      let _lhs17 = _185_v63;
      let _lhs18 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
      let _lhs19 = _118_globalState;
      _114_v2 = _rhs34;
      _lhs16.f2 = _rhs35;
      _lhs17[_lhs18] = _rhs36;
      _lhs19.f2 = _rhs37;
      let _254_v111;
      _254_v111 = _dafny.Map.Empty.slice().updateUnsafe(_114_v2,(_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]);
      _254_v111 = (_254_v111).update(!(_114_v2), _112_v0);
      let _255_v112;
      _255_v112 = _module.D3.create_DC8(_231_v96);
      let _256_v113;
      _256_v113 = _dafny.Seq.of(_255_v112, _255_v112);
      let _257_v114;
      _257_v114 = _module.D16.create_DC40(_256_v113);
      _256_v113 = _dafny.Seq.Concat(_256_v113, _dafny.Seq.Concat(_256_v113, (_257_v114).dtor_cf66));
      let _258_v117;
      _258_v117 = _dafny.Map.Empty.slice().updateUnsafe((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))],_233_v98);
      let _259_v119;
      _259_v119 = _dafny.Set.fromElements(_112_v0);
      let _260_v121;
      _260_v121 = _dafny.Map.Empty.slice().updateUnsafe(_233_v98,_233_v98);
      let _261_v122;
      let _nw38 = Array((new BigNumber(23)).toNumber());
      _nw38[(_dafny.ZERO).toNumber()] = _233_v98;
      _nw38[(_dafny.ONE).toNumber()] = _233_v98;
      _nw38[(new BigNumber(2)).toNumber()] = _233_v98;
      _nw38[(new BigNumber(3)).toNumber()] = _233_v98;
      _nw38[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(284),false);
      _nw38[(new BigNumber(5)).toNumber()] = (_233_v98).Merge(_233_v98);
      _nw38[(new BigNumber(6)).toNumber()] = _233_v98;
      _nw38[(new BigNumber(7)).toNumber()] = _233_v98;
      _nw38[(new BigNumber(8)).toNumber()] = (function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of (_116_v4).Elements) {
          let _262_v115 = _compr_19;
          if ((_116_v4).contains(_262_v115)) {
            _coll19.push([_module.__default.safeDivisionInt(_262_v115, new BigNumber(-416)),_114_v2]);
          }
        }
        return _coll19;
      }()).Merge(_233_v98);
      _nw38[(new BigNumber(9)).toNumber()] = (_233_v98).Merge(_233_v98);
      _nw38[(new BigNumber(10)).toNumber()] = (_233_v98).Merge(_233_v98);
      _nw38[(new BigNumber(11)).toNumber()] = _233_v98;
      _nw38[(new BigNumber(12)).toNumber()] = ((_233_v98).update((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))], _114_v2)).Merge(_233_v98);
      _nw38[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(197),_114_v2);
      _nw38[(new BigNumber(14)).toNumber()] = ((_114_v2) ? (_233_v98) : (function () {
        let _coll20 = new _dafny.Map();
        for (const _compr_20 of (_116_v4).Elements) {
          let _263_v116 = _compr_20;
          if ((_116_v4).contains(_263_v116)) {
            _coll20.push([_module.__default.safeDivisionInt(_263_v116, (_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]),_114_v2]);
          }
        }
        return _coll20;
      }()));
      _nw38[(new BigNumber(15)).toNumber()] = (((_258_v117).contains(new BigNumber((_184_v62).length))) ? ((_258_v117).get(new BigNumber((_184_v62).length))) : (_233_v98));
      _nw38[(new BigNumber(16)).toNumber()] = function () {
        let _coll21 = new _dafny.Map();
        for (const _compr_21 of (_259_v119).Elements) {
          let _264_v118 = _compr_21;
          if ((_259_v119).contains(_264_v118)) {
            _coll21.push([_module.__default.safeModuloInt(_264_v118, _112_v0),_114_v2]);
          }
        }
        return _coll21;
      }();
      _nw38[(new BigNumber(17)).toNumber()] = function () {
        let _coll22 = new _dafny.Map();
        for (const _compr_22 of _dafny.IntegerRange(new BigNumber(831), new BigNumber(188))) {
          let _265_v120 = _compr_22;
          if (((new BigNumber(831)).isLessThanOrEqualTo(_265_v120)) && ((_265_v120).isLessThan(new BigNumber(188)))) {
            _coll22.push([(_265_v120).plus((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]),_114_v2]);
          }
        }
        return _coll22;
      }();
      _nw38[(new BigNumber(18)).toNumber()] = _233_v98;
      _nw38[(new BigNumber(19)).toNumber()] = _233_v98;
      _nw38[(new BigNumber(20)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))],_114_v2)).Merge(_233_v98);
      _nw38[(new BigNumber(21)).toNumber()] = (((_260_v121).contains(_233_v98)) ? ((_260_v121).get(_233_v98)) : (_233_v98));
      _nw38[(new BigNumber(22)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))],!(_114_v2));
      _261_v122 = _nw38;
      let _266_v123;
      let _nw39 = new _module.C2();
      _nw39.__ctor(_112_v0, _module.__default.fm37(_114_v2, new BigNumber(799), _module.__default.fm38(_dafny.Map.Empty.slice().updateUnsafe(_112_v0,_114_v2), _114_v2, _114_v2, _114_v2, _118_globalState), _118_globalState));
      _266_v123 = _nw39;
      let _267_v124;
      _267_v124 = _dafny.MultiSet.fromElements(_266_v123);
      let _rhs38 = _261_v122;
      let _rhs39 = (_267_v124).IsSubsetOf(((_114_v2) ? (_dafny.MultiSet.fromElements(_266_v123)) : (_267_v124)));
      _261_v122 = _rhs38;
      _114_v2 = _rhs39;
      let _268_v125;
      let _269_v126;
      let _out20;
      let _out21;
      let _outcollector5 = (_266_v123).m1(_118_globalState);
      _out20 = _outcollector5[0];
      _out21 = _outcollector5[1];
      _268_v125 = _out20;
      _269_v126 = _out21;
      let _hi3 = _268_v125;
      for (let _270_i10 = _268_v125; _270_i10.isLessThan(_hi3); _270_i10 = _270_i10.plus(_dafny.ONE)) {
        if (_114_v2) {
          _231_v96 = (_231_v96).update(!(!(_269_v126)), ((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]).isLessThan(new BigNumber(106)));
          _114_v2 = ((function () {
            let _coll23 = new _dafny.Set();
            for (const _compr_23 of _dafny.IntegerRange(new BigNumber(816), new BigNumber(77))) {
              let _271_v128 = _compr_23;
              if (((new BigNumber(816)).isLessThanOrEqualTo(_271_v128)) && ((_271_v128).isLessThan(new BigNumber(77)))) {
                _coll23.add(_module.__default.safeDivisionInt(_271_v128, _270_i10));
              }
            }
            return _coll23;
          }()).Union(_259_v119)).IsSubsetOf((function () {
            let _coll24 = new _dafny.Set();
            for (const _compr_24 of _dafny.IntegerRange(new BigNumber(276), new BigNumber(504))) {
              let _272_v127 = _compr_24;
              if (((new BigNumber(276)).isLessThanOrEqualTo(_272_v127)) && ((_272_v127).isLessThan(new BigNumber(504)))) {
                _coll24.add(_module.__default.safeModuloInt(_272_v127, (_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]));
              }
            }
            return _coll24;
          }()).Union(_dafny.Set.fromElements((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))], new BigNumber((_dafny.MultiSet.fromElements(_112_v0, _268_v125)).cardinality()))));
          _114_v2 = ((_266_v123).fm2((_266_v123).fm2((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))], _183_v61, _118_globalState), _dafny.Set.fromElements(true), _118_globalState)).isLessThan(new BigNumber((_116_v4).cardinality()));
          let _273_v129;
          let _nw40 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
          _273_v129 = _nw40;
          let _index20 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_273_v129).length));
          (_273_v129)[_index20] = _dafny.Seq.Concat(_182_v60, _182_v60);
          let _index21 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_273_v129).length));
          (_273_v129)[_index21] = _dafny.Seq.of(_269_v126, !(false) || (_114_v2));
          let _274_v130;
          let _init4 = ((_275_v2) => function (_276_i11) {
            return (!(false)) || (!(_275_v2));
          })(_114_v2);
          let _nw41 = Array((new BigNumber(16)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw41.length); _i0_4++) {
            _nw41[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _274_v130 = _nw41;
          let _index22 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_274_v130).length));
          (_274_v130)[_index22] = _269_v126;
          let _277_v131;
          _277_v131 = _module.D16.create_DC42((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))], new BigNumber((_184_v62).length), _113_v1);
          let _278_v132;
          _278_v132 = _dafny.Map.Empty.slice().updateUnsafe((_273_v129)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_273_v129).length))],new BigNumber((_184_v62).length));
          let _index23 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_274_v130).length));
          let _rhs40 = (_277_v131).dtor_cf70;
          let _rhs41 = _114_v2;
          let _rhs42 = !((new BigNumber((_278_v132).length)).isLessThan(_268_v125));
          let _lhs20 = _274_v130;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_274_v130).length));
          _113_v1 = _rhs40;
          _114_v2 = _rhs41;
          _lhs20[_lhs21] = _rhs42;
        } else {
          let _279_v133;
          _279_v133 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_232_v97, _dafny.Seq.update(_232_v97, _module.__default.safeIndex(new BigNumber((_231_v96).length), new BigNumber((_232_v97).length)), _268_v125)),_268_v125);
          _279_v133 = (_279_v133).update(_232_v97, new BigNumber((_232_v97).length));
          let _index24 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
          (_185_v63)[_index24] = _270_i10;
          (_266_v123).m2(_268_v125, _118_globalState);
          let _280_v134;
          _280_v134 = _module.D7.create_DC23(_module.__default.fm37(_114_v2, _112_v0, _269_v126, _118_globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), ((_281_v1) => function (_282_i12) {
  return _281_v1;
})(_113_v1)), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("nltodsjko")).length), (_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]));
          _280_v134 = _280_v134;
          _185_v63 = _185_v63;
        }
        if (_114_v2) {
          let _283_v135;
          _283_v135 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xmmvfaq")),_114_v2);
          let _284_v136;
          _284_v136 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xsltdk"), _184_v62, _dafny.Seq.UnicodeFromString("etaeaj"));
          let _285_v137;
          _285_v137 = _dafny.Map.Empty.slice().updateUnsafe(_269_v126,_184_v62);
          _283_v135 = (_283_v135).update((_284_v136).update((((_285_v137).contains(_114_v2)) ? ((_285_v137).get(_114_v2)) : (_184_v62)), _module.__default.abs(_270_i10)), _dafny.Seq.IsPrefixOf(_184_v62, _184_v62));
          let _286_v138;
          let _nw42 = new _module.C0();
          _nw42.__ctor(_270_i10);
          _286_v138 = _nw42;
          let _287_v139;
          let _nw43 = new _module.C1();
          _nw43.__ctor(new BigNumber((_184_v62).length), (_270_i10).minus(_268_v125));
          _287_v139 = _nw43;
          let _288_v140;
          let _nw44 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
          _288_v140 = _nw44;
          let _index25 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_288_v140).length));
          (_288_v140)[_index25] = _dafny.Seq.of(_270_i10);
          let _289_v141;
          _289_v141 = _dafny.Map.Empty.slice().updateUnsafe(_287_v139,_112_v0);
          let _index26 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_288_v140).length));
          let _rhs43 = _287_v139;
          let _rhs44 = _232_v97;
          let _rhs45 = (_114_v2) === (_114_v2);
          let _rhs46 = ((_114_v2) ? ((_232_v97)[_module.__default.safeIndex(new BigNumber((_289_v141).length), new BigNumber((_232_v97).length))]) : (_112_v0));
          let _lhs22 = _288_v140;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_288_v140).length));
          let _lhs24 = _118_globalState;
          _287_v139 = _rhs43;
          _lhs22[_lhs23] = _rhs44;
          _114_v2 = _rhs45;
          _lhs24.f2 = _rhs46;
          _114_v2 = _269_v126;
          _113_v1 = _113_v1;
        } else {
          (_266_v123).m4((new BigNumber((_dafny.MultiSet.fromElements(_269_v126)).cardinality())).isLessThanOrEqualTo(_270_i10), _118_globalState);
          let _290_v142;
          let _nw45 = Array((new BigNumber(2)).toNumber());
          _nw45[(_dafny.ZERO).toNumber()] = _184_v62;
          _nw45[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_184_v62, _dafny.Seq.UnicodeFromString("dmsgbdkrd"));
          _290_v142 = _nw45;
          let _index27 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_290_v142).length));
          (_290_v142)[_index27] = _module.__default.fm11(new BigNumber(74), new BigNumber((_259_v119).length), _112_v0, _112_v0, _118_globalState);
          let _291_v143;
          _291_v143 = _dafny.Seq.of(_184_v62);
          let _index28 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_290_v142).length));
          (_290_v142)[_index28] = _dafny.Seq.Concat(_dafny.Seq.Concat(_184_v62, _dafny.Seq.UnicodeFromString("laqun")), _dafny.Seq.Concat(_184_v62, (_291_v143)[_module.__default.safeIndex(_270_i10, new BigNumber((_291_v143).length))]));
          let _292_v144;
          let _nw46 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
          _292_v144 = _nw46;
          _292_v144 = _292_v144;
          let _293_v145;
          let _nw47 = new _module.C1();
          _nw47.__ctor(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_269_v126,_268_v125)).Merge(_254_v111)).length), new BigNumber(296));
          _293_v145 = _nw47;
          let _rhs47 = _269_v126;
          let _rhs48 = _module.__default.fm28(_268_v125, _118_globalState);
          let _rhs49 = _269_v126;
          let _rhs50 = _293_v145;
          _114_v2 = _rhs47;
          _232_v97 = _rhs48;
          _269_v126 = _rhs49;
          _293_v145 = _rhs50;
          let _294_v146;
          _294_v146 = (_254_v111).update(_114_v2, _268_v125);
          let _295_v147;
          let _nw48 = Array((new BigNumber(26)).toNumber());
          _nw48[(_dafny.ZERO).toNumber()] = _294_v146;
          _nw48[(_dafny.ONE).toNumber()] = _294_v146;
          _nw48[(new BigNumber(2)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(3)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(4)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(5)).toNumber()] = _254_v111;
          _nw48[(new BigNumber(6)).toNumber()] = _254_v111;
          _nw48[(new BigNumber(7)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(8)).toNumber()] = _254_v111;
          _nw48[(new BigNumber(9)).toNumber()] = _254_v111;
          _nw48[(new BigNumber(10)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_269_v126,new BigNumber(-772));
          _nw48[(new BigNumber(12)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(13)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(14)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(15)).toNumber()] = ((_269_v126) ? (_294_v146) : (_294_v146));
          _nw48[(new BigNumber(16)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(17)).toNumber()] = _254_v111;
          _nw48[(new BigNumber(18)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(19)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(20)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(21)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(22)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(23)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(24)).toNumber()] = _294_v146;
          _nw48[(new BigNumber(25)).toNumber()] = _294_v146;
          _295_v147 = _nw48;
          let _index29 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_295_v147).length));
          (_295_v147)[_index29] = _254_v111;
          let _index30 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_295_v147).length));
          (_295_v147)[_index30] = _294_v146;
        }
        let _296_v148;
        _296_v148 = _dafny.Map.Empty.slice().updateUnsafe(_269_v126,_dafny.Set.fromElements(_114_v2, _269_v126, _269_v126));
        let _297_v149;
        let _nw49 = new _module.C1();
        _nw49.__ctor(new BigNumber(((((_296_v148).contains(_114_v2)) ? ((_296_v148).get(_114_v2)) : (_dafny.Set.fromElements(_114_v2)))).length), _112_v0);
        _297_v149 = _nw49;
        let _298_v150;
        _298_v150 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(640),_114_v2)).length))),_297_v149);
        let _299_v151;
        _299_v151 = _module.D11.create_DC31((((_298_v150).contains((_dafny.ZERO).minus(_112_v0))) ? ((_298_v150).get((_dafny.ZERO).minus(_112_v0))) : (_297_v149)));
        let _pat_let_tv5 = _297_v149;
        let _source5 = function (_pat_let5_0) {
          return function (_300_dt__update__tmp_h2) {
            return function (_pat_let6_0) {
              return function (_301_dt__update_hcf51_h0) {
                return _module.D11.create_DC31(_301_dt__update_hcf51_h0);
              }(_pat_let6_0);
            }(_pat_let_tv5);
          }(_pat_let5_0);
        }(_299_v151);
        if (_source5.is_DC32) {
          let _302___mcc_h6 = (_source5).cf52;
          let _303___mcc_h7 = (_source5).cf53;
          let _304___mcc_h8 = (_source5).cf54;
          let _305_cf54 = _304___mcc_h8;
          let _306_cf53 = _303___mcc_h7;
          let _307_cf52 = _302___mcc_h6;
          let _308_v152;
          let _nw50 = Array((new BigNumber(12)).toNumber()).fill([]);
          _308_v152 = _nw50;
          let _309_v153;
          _309_v153 = _module.D1.create_DC5();
          let _rhs51 = (_266_v123).fm9(_309_v153, _114_v2, _118_globalState);
          let _rhs52 = _114_v2;
          let _rhs53 = _112_v0;
          let _rhs54 = ((_259_v119).Union(_259_v119)).Difference(_259_v119);
          let _rhs55 = _308_v152;
          let _lhs25 = _118_globalState;
          _305_cf54 = _rhs51;
          _269_v126 = _rhs52;
          _lhs25.f2 = _rhs53;
          _259_v119 = _rhs54;
          _308_v152 = _rhs55;
          _306_cf53 = (_268_v125).isLessThanOrEqualTo(_270_i10);
          let _310_v154;
          let _nw51 = new _module.C2();
          _nw51.__ctor(new BigNumber(320), (_dafny.ZERO).minus(_112_v0));
          _310_v154 = _nw51;
          let _311_v155;
          let _nw52 = Array((new BigNumber(10)).toNumber()).fill(false);
          _311_v155 = _nw52;
          let _312_v156;
          _312_v156 = _module.D7.create_DC22(_311_v155, _184_v62);
          _311_v155 = (_312_v156).dtor_cf38;
        } else {
          let _313___mcc_h9 = (_source5).cf51;
          let _314_cf51 = _313___mcc_h9;
          _114_v2 = !(_270_i10).isEqualTo(_268_v125);
          _114_v2 = (_234_v99).equals(_module.__default.fm5(_232_v97, !(_269_v126), (_314_cf51).fm13(_270_i10, _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sbeh")), _118_globalState), _118_globalState));
          let _315_v157;
          let _nw53 = new _module.C3();
          _nw53.__ctor(_114_v2, (_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))], _270_i10);
          _315_v157 = _nw53;
          let _316_v158;
          _316_v158 = _dafny.Map.Empty.slice().updateUnsafe(_184_v62,_315_v157);
          let _317_v159;
          let _nw54 = new _module.C2();
          _nw54.__ctor(_268_v125, (((_254_v111).contains(_114_v2)) ? ((_254_v111).get(_114_v2)) : (new BigNumber(((_316_v158).update(_184_v62, _315_v157)).length))));
          _317_v159 = _nw54;
          let _318_v160;
          _318_v160 = _317_v159;
          let _rhs56 = false;
          let _rhs57 = _269_v126;
          let _rhs58 = _318_v160;
          _269_v126 = _rhs56;
          _114_v2 = _rhs57;
          _318_v160 = _rhs58;
          let _319_v161;
          _319_v161 = _dafny.MultiSet.fromElements(_184_v62, _184_v62);
          let _320_v162;
          let _nw55 = new _module.C0();
          _nw55.__ctor((_297_v149).fm13((_314_cf51).fm13(_315_v157.f6, _319_v161, _118_globalState), _319_v161, _118_globalState));
          _320_v162 = _nw55;
          let _321_v163;
          _321_v163 = _dafny.Seq.of(_320_v162, _320_v162, _320_v162);
          let _322_v164;
          _322_v164 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_321_v163, _321_v163),_112_v0);
          let _index31 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
          let _rhs59 = (_dafny.ZERO).minus(((_320_v162).f8).multipliedBy((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]));
          let _rhs60 = (_322_v164).update(_dafny.Seq.of(_320_v162), (_dafny.ZERO).minus((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]));
          let _lhs26 = _185_v63;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
          _lhs26[_lhs27] = _rhs59;
          _322_v164 = _rhs60;
        }
        let _323_v165;
        let _nw56 = new _module.C3();
        _nw56.__ctor((_183_v61).IsDisjointFrom(_dafny.Set.fromElements(_114_v2)), (_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))], _112_v0);
        _323_v165 = _nw56;
        _323_v165 = _323_v165;
      }
      _112_v0 = _module.__default.fm37(_module.__default.fm38(_233_v98, _114_v2, _114_v2, _269_v126, _118_globalState), (_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))], _114_v2, _118_globalState);
      if ((_259_v119).equals(_259_v119)) {
        let _324_v166;
        let _nw57 = Array((new BigNumber(13)).toNumber());
        _nw57[(_dafny.ZERO).toNumber()] = _185_v63;
        _nw57[(_dafny.ONE).toNumber()] = _185_v63;
        _nw57[(new BigNumber(2)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(3)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(4)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(5)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(6)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(7)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(8)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(9)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(10)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(11)).toNumber()] = _185_v63;
        _nw57[(new BigNumber(12)).toNumber()] = _185_v63;
        _324_v166 = _nw57;
        let _index32 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_324_v166).length));
        (_324_v166)[_index32] = _185_v63;
        let _index33 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_324_v166).length));
        (_324_v166)[_index33] = _185_v63;
        let _325_v167;
        _325_v167 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_183_v61).length),_185_v63);
        let _326_v168;
        _326_v168 = _dafny.Map.Empty.slice().updateUnsafe((_325_v167).Merge(_325_v167),_269_v126);
        if ((((_326_v168).contains(_325_v167)) ? ((_326_v168).get(_325_v167)) : (_269_v126))) {
          let _327_v169;
          _327_v169 = _module.D4.create_DC14((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))], _114_v2, new BigNumber((_184_v62).length), !(!(_269_v126)), _112_v0);
          let _328_v170;
          _328_v170 = _dafny.Map.Empty.slice().updateUnsafe(_327_v169,new BigNumber(8));
          let _329_v171;
          _329_v171 = _dafny.Map.Empty.slice().updateUnsafe(_114_v2,_328_v170);
          let _330_v172;
          let _nw58 = new _module.C1();
          _nw58.__ctor(_268_v125, new BigNumber((_329_v171).length));
          _330_v172 = _nw58;
          let _331_v173;
          _331_v173 = _dafny.Map.Empty.slice().updateUnsafe(_268_v125,_231_v96);
          let _332_v174;
          let _init5 = ((_333_v2) => function (_334_i13) {
            return _333_v2;
          })(_114_v2);
          let _nw59 = Array((new BigNumber(28)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw59.length); _i0_5++) {
            _nw59[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _332_v174 = _nw59;
          let _index34 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_332_v174).length));
          (_332_v174)[_index34] = (_112_v0).isEqualTo(_268_v125);
          let _335_v175;
          let _nw60 = new _module.C1();
          _nw60.__ctor(_112_v0, _112_v0);
          _335_v175 = _nw60;
          let _336_v176;
          _336_v176 = _module.D9.create_DC28(_269_v126, new BigNumber((_dafny.Seq.of(_335_v175, _335_v175, _335_v175, _335_v175)).length), _112_v0);
          let _337_v177;
          _337_v177 = _dafny.Seq.of(_336_v176);
          let _index35 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_332_v174).length));
          let _index36 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
          let _rhs61 = _330_v172;
          let _rhs62 = _331_v173;
          let _rhs63 = !(_114_v2);
          let _rhs64 = ((_114_v2) ? ((new BigNumber((_337_v177).length)).minus(new BigNumber((_182_v60).length))) : (_335_v175.f6));
          let _lhs28 = _332_v174;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_332_v174).length));
          let _lhs30 = _185_v63;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
          _330_v172 = _rhs61;
          _331_v173 = _rhs62;
          _lhs28[_lhs29] = _rhs63;
          _lhs30[_lhs31] = _rhs64;
          let _index37 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_332_v174).length));
          (_332_v174)[_index37] = _269_v126;
          let _338_v178;
          let _339_v179;
          let _out22;
          let _out23;
          let _outcollector6 = (_266_v123).m1(_118_globalState);
          _out22 = _outcollector6[0];
          _out23 = _outcollector6[1];
          _338_v178 = _out22;
          _339_v179 = _out23;
          let _index38 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_332_v174).length));
          let _rhs65 = ((_267_v124).Union(_267_v124)).update(_266_v123, _module.__default.abs(new BigNumber(328)));
          let _rhs66 = (_332_v174)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((_332_v174).length))];
          let _rhs67 = _266_v123;
          let _rhs68 = _114_v2;
          let _rhs69 = (_182_v60)[_module.__default.safeIndex(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_335_v175).f5)).length),_113_v1)).update((_dafny.ZERO).minus(_335_v175.f6), _113_v1)).length), new BigNumber((_182_v60).length))];
          let _lhs32 = _332_v174;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_332_v174).length));
          _267_v124 = _rhs65;
          _lhs32[_lhs33] = _rhs66;
          _266_v123 = _rhs67;
          _114_v2 = _rhs68;
          _339_v179 = _rhs69;
          _114_v2 = !(new BigNumber((_184_v62).length)).isEqualTo(new BigNumber(998));
        } else {
          (_266_v123).m4(!(true), _118_globalState);
          let _340_v180;
          let _nw61 = Array((new BigNumber(29)).toNumber()).fill(false);
          _340_v180 = _nw61;
          let _index39 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_340_v180).length));
          (_340_v180)[_index39] = !(!(_114_v2) || (true));
          let _index40 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_340_v180).length));
          (_340_v180)[_index40] = _114_v2;
          (_118_globalState).f2 = _112_v0;
          let _341_v181;
          let _342_v182;
          let _out24;
          let _out25;
          let _outcollector7 = (_266_v123).m1(_118_globalState);
          _out24 = _outcollector7[0];
          _out25 = _outcollector7[1];
          _341_v181 = _out24;
          _342_v182 = _out25;
          _116_v4 = ((_116_v4).update(new BigNumber((_184_v62).length), _module.__default.abs(_341_v181))).Difference(_116_v4);
        }
        let _343_v183;
        let _nw62 = Array((new BigNumber(2)).toNumber()).fill([]);
        _343_v183 = _nw62;
        let _344_v184;
        let _nw63 = Array((new BigNumber(16)).toNumber()).fill(false);
        _344_v184 = _nw63;
        let _index41 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_343_v183).length));
        (_343_v183)[_index41] = _344_v184;
        let _index42 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_343_v183).length));
        let _index43 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
        let _rhs70 = _268_v125;
        let _rhs71 = _344_v184;
        let _rhs72 = _268_v125;
        let _rhs73 = _184_v62;
        let _rhs74 = ((_114_v2) ? (_184_v62) : (_184_v62));
        let _lhs34 = _118_globalState;
        let _lhs35 = _343_v183;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_343_v183).length));
        let _lhs37 = _185_v63;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
        _lhs34.f2 = _rhs70;
        _lhs35[_lhs36] = _rhs71;
        _lhs37[_lhs38] = _rhs72;
        _184_v62 = _rhs73;
        _184_v62 = _rhs74;
        let _345_v185;
        let _nw64 = Array((new BigNumber(2)).toNumber());
        _345_v185 = _nw64;
        let _346_v186;
        let _nw65 = new _module.C1();
        _nw65.__ctor((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))], new BigNumber((_253_v110).length));
        _346_v186 = _nw65;
        let _index44 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_345_v185).length));
        (_345_v185)[_index44] = _346_v186;
        let _index45 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_345_v185).length));
        (_345_v185)[_index45] = _346_v186;
        let _rhs75 = (_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))];
        let _rhs76 = _113_v1;
        let _lhs39 = _118_globalState;
        _lhs39.f2 = _rhs75;
        _113_v1 = _rhs76;
      } else {
        _184_v62 = _dafny.Seq.UnicodeFromString("qcqcxq");
        (_266_v123).m2(_module.__default.safeDivisionInt(new BigNumber((_184_v62).length), (_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]), _118_globalState);
        _269_v126 = !(_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("miaw"), new _dafny.CodePoint('a'.codePointAt(0)))) || ((((_233_v98).contains(new BigNumber(461))) ? ((_233_v98).get(new BigNumber(461))) : (_114_v2)));
        if ((_269_v126) || (!(_114_v2))) {
          let _347_v187;
          let _nw66 = new _module.C0();
          _nw66.__ctor((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]);
          _347_v187 = _nw66;
          let _348_v188;
          _348_v188 = _dafny.Seq.of(_254_v111, _254_v111);
          let _349_v189;
          _349_v189 = _dafny.Set.fromElements(_113_v1);
          let _350_v190;
          let _nw67 = Array((new BigNumber(11)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _254_v111;
          _nw67[(_dafny.ONE).toNumber()] = (_254_v111).Merge(_254_v111);
          _nw67[(new BigNumber(2)).toNumber()] = (_254_v111).Merge(_254_v111);
          _nw67[(new BigNumber(3)).toNumber()] = _254_v111;
          _nw67[(new BigNumber(4)).toNumber()] = (_348_v188)[_module.__default.safeIndex((_347_v187).f8, new BigNumber((_348_v188).length))];
          _nw67[(new BigNumber(5)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_269_v126,_268_v125)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_269_v126,new BigNumber((_349_v189).length)));
          _nw67[(new BigNumber(6)).toNumber()] = _module.__default.fm39(_118_globalState);
          _nw67[(new BigNumber(7)).toNumber()] = _254_v111;
          _nw67[(new BigNumber(8)).toNumber()] = _254_v111;
          _nw67[(new BigNumber(9)).toNumber()] = _254_v111;
          _nw67[(new BigNumber(10)).toNumber()] = _254_v111;
          _350_v190 = _nw67;
          let _index46 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_350_v190).length));
          (_350_v190)[_index46] = (_dafny.Map.Empty.slice().updateUnsafe(!(_269_v126),_268_v125)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_269_v126,new BigNumber(193)));
          let _index47 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_350_v190).length));
          let _rhs77 = _module.__default.safeModuloInt(_268_v125, new BigNumber((_dafny.Seq.UnicodeFromString("awmjrtbl")).length));
          let _rhs78 = _232_v97;
          let _rhs79 = (_348_v188)[_module.__default.safeIndex(_268_v125, new BigNumber((_348_v188).length))];
          let _rhs80 = (_266_v123).fm2((_dafny.ZERO).minus((_dafny.ZERO).minus((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))])), _183_v61, _118_globalState);
          let _rhs81 = _269_v126;
          let _lhs40 = _118_globalState;
          let _lhs41 = _350_v190;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_350_v190).length));
          let _lhs43 = _118_globalState;
          _lhs40.f2 = _rhs77;
          _232_v97 = _rhs78;
          _lhs41[_lhs42] = _rhs79;
          _lhs43.f2 = _rhs80;
          _114_v2 = _rhs81;
          let _rhs82 = (_112_v0).plus((_185_v63)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length))]);
          let _rhs83 = _113_v1;
          let _rhs84 = function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of _dafny.IntegerRange(new BigNumber(243), new BigNumber(405))) {
              let _351_v191 = _compr_25;
              if (((new BigNumber(243)).isLessThanOrEqualTo(_351_v191)) && ((_351_v191).isLessThan(new BigNumber(405)))) {
                _coll25.push([_module.__default.safeModuloInt(_351_v191, new BigNumber(((_116_v4).update(new BigNumber((_183_v61).length), _module.__default.abs(new BigNumber((_231_v96).length)))).cardinality())),(_347_v187).f8]);
              }
            }
            return _coll25;
          }();
          let _lhs44 = _118_globalState;
          _lhs44.f2 = _rhs82;
          _113_v1 = _rhs83;
          _253_v110 = _rhs84;
          _184_v62 = _dafny.Seq.update(_184_v62, _module.__default.safeIndex(_268_v125, new BigNumber((_184_v62).length)), _113_v1);
          let _352_v192;
          _352_v192 = (_183_v61).Difference(_183_v61);
          _352_v192 = _183_v61;
        } else {
          let _353_v193;
          let _354_v194;
          let _out26;
          let _out27;
          let _outcollector8 = (_266_v123).m1(_118_globalState);
          _out26 = _outcollector8[0];
          _out27 = _outcollector8[1];
          _353_v193 = _out26;
          _354_v194 = _out27;
          let _355_v195;
          let _nw68 = new _module.C2();
          _nw68.__ctor((_232_v97)[_module.__default.safeIndex(_112_v0, new BigNumber((_232_v97).length))], new BigNumber(136));
          _355_v195 = _nw68;
          let _356_v196;
          _356_v196 = _dafny.Seq.of(_184_v62);
          let _357_v197;
          _357_v197 = _dafny.Seq.of(_259_v119);
          let _index48 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
          let _index49 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
          let _rhs85 = (((_233_v98).contains(new BigNumber(((_356_v196)[_module.__default.safeIndex(new BigNumber(((_357_v197)[_module.__default.safeIndex(_353_v193, new BigNumber((_357_v197).length))]).length), new BigNumber((_356_v196).length))]).length))) ? ((_233_v98).get(new BigNumber(((_356_v196)[_module.__default.safeIndex(new BigNumber(((_357_v197)[_module.__default.safeIndex(_353_v193, new BigNumber((_357_v197).length))]).length), new BigNumber((_356_v196).length))]).length))) : (_354_v194));
          let _rhs86 = (_112_v0).plus((_353_v193).multipliedBy(new BigNumber(955)));
          let _rhs87 = _268_v125;
          let _lhs45 = _185_v63;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
          let _lhs47 = _185_v63;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_185_v63).length));
          _354_v194 = _rhs85;
          _lhs45[_lhs46] = _rhs86;
          _lhs47[_lhs48] = _rhs87;
          _268_v125 = _268_v125;
          (_118_globalState).f2 = _module.__default.safeModuloInt(_112_v0, _268_v125);
        }
        _114_v2 = !(!(_114_v2)) || (_269_v126);
      }
      process.stdout.write(_dafny.toString(_112_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_113_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_114_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_115_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v4).equals(_dafny.MultiSet.fromElements(new BigNumber(-723), new BigNumber(-723), new BigNumber(-723)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_117_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),_dafny.MultiSet.fromElements(new BigNumber(-723), new BigNumber(-723), new BigNumber(-723))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_118_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_118_globalState).f4).equals(_dafny.MultiSet.fromElements(new BigNumber(-723), new BigNumber(-723), new BigNumber(-723), new BigNumber(-723), new BigNumber(-723), new BigNumber(-723), new BigNumber(-723), new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_176_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_182_v60, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_183_v61).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_184_v62).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_v63)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_189_v64).dtor_cf5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v64).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_231_v96).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_232_v97, _dafny.Seq.of(new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_233_v98).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_v99).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_235_i8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_247_v105, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(-723), new BigNumber(-723), new BigNumber(-723))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_252_v109)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,new BigNumber(690)).updateUnsafe(new BigNumber(536),_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v110).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,new BigNumber(690)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v111).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO).updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_255_v112).dtor_cf8).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_256_v113, _dafny.Seq.of(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,true)), _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,true)), _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,true)), _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,true)), _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,true)), _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,true))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_257_v114).dtor_cf66, _dafny.Seq.of(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,true)), _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(false,true))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v117).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v119).equals(_dafny.Set.fromElements(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v121).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(284),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),false).updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true).updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(197),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-723),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false).updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_261_v122)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_267_v124).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_268_v125));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_269_v126));
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
    static create_DC2() {
      let $dt = new D0(1);
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D0(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf3 === other.cf3;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, false);
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
    static create_DC5() {
      let $dt = new D1(0);
      return $dt;
    }
    static create_DC6(cf5, cf6) {
      let $dt = new D1(1);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC4(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5";
      } else if (this.$tag === 1) {
        return "D1.DC6" + "(" + this.cf5.toVerbatimString(true) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5();
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
    static create_DC7(cf7) {
      let $dt = new D2(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf9, cf10, cf11) {
      let $dt = new D3(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC8(cf8) {
      let $dt = new D3(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC10(cf12) {
      let $dt = new D3(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(_dafny.MultiSet.Empty, _module.D0.Default(), _dafny.ZERO);
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
    static create_DC12(cf14, cf15, cf16) {
      let $dt = new D4(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC13(cf17, cf18, cf19) {
      let $dt = new D4(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC14(cf20, cf21, cf22, cf23, cf24) {
      let $dt = new D4(2);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC11(cf13) {
      let $dt = new D4(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC15(cf25) {
      let $dt = new D4(4);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get is_DC15() { return this.$tag === 4; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 4) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18) && this.cf19 === other.cf19;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21 && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(false, _dafny.ZERO, false);
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
    static create_DC17(cf27, cf28, cf29) {
      let $dt = new D5(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC18(cf30, cf31, cf32) {
      let $dt = new D5(1);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC19(cf33, cf34, cf35) {
      let $dt = new D5(2);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC16(cf26) {
      let $dt = new D5(3);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get is_DC16() { return this.$tag === 3; }
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
    toString() {
      if (this.$tag === 0) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC19" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf26 === other.cf26;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC17(_dafny.ZERO, [], _dafny.Map.Empty);
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
    static create_DC20(cf36) {
      let $dt = new D6(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36;
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf38, cf39) {
      let $dt = new D7(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC23(cf40, cf41, cf42) {
      let $dt = new D7(1);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC24() {
      let $dt = new D7(2);
      return $dt;
    }
    static create_DC21(cf37) {
      let $dt = new D7(3);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC25(cf43) {
      let $dt = new D7(4);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get is_DC21() { return this.$tag === 3; }
    get is_DC25() { return this.$tag === 4; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf38) + ", " + this.cf39.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC23" + "(" + _dafny.toString(this.cf40) + ", " + this.cf41.toVerbatimString(true) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC24";
      } else if (this.$tag === 3) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 4) {
        return "D7.DC25" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC22([], _dafny.Seq.UnicodeFromString(""));
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
    static create_DC26(cf44) {
      let $dt = new D8(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC26" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28(cf46, cf47, cf48) {
      let $dt = new D9(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC27(cf45) {
      let $dt = new D9(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC28(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC30(cf50) {
      let $dt = new D10(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC29(cf49) {
      let $dt = new D10(1);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC30" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC30(_dafny.ZERO);
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
    static create_DC32(cf52, cf53, cf54) {
      let $dt = new D11(0);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC31(cf51) {
      let $dt = new D11(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf52 === other.cf52 && this.cf53 === other.cf53 && this.cf54 === other.cf54;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf51 === other.cf51;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC32(false, false, false);
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
    static create_DC34(cf56) {
      let $dt = new D12(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC35(cf57, cf58, cf59) {
      let $dt = new D12(1);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC36(cf60, cf61, cf62) {
      let $dt = new D12(2);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC33(cf55) {
      let $dt = new D12(3);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC35" + "(" + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + this.cf62.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf56 === other.cf56;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf60 === other.cf60 && this.cf61 === other.cf61 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC34(false);
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
    static create_DC37(cf63) {
      let $dt = new D13(0);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC37" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf63, other.cf63);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC38(cf64) {
      let $dt = new D14(0);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC38" + "(" + _dafny.toString(this.cf64) + ")";
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
      return _dafny.MultiSet.Empty;
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
    static create_DC39(cf65) {
      let $dt = new D15(0);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC39" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf65 === other.cf65;
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
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf67) {
      let $dt = new D16(0);
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC42(cf68, cf69, cf70) {
      let $dt = new D16(1);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC40(cf66) {
      let $dt = new D16(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC43(cf71) {
      let $dt = new D16(3);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get is_DC43() { return this.$tag === 3; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC40" + "(" + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 3) {
        return "D16.DC43" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf68, other.cf68) && _dafny.areEqual(this.cf69, other.cf69) && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf66, other.cf66);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf71, other.cf71);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC41(_dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
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
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f4 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f8 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f8) {
      let _this = this;
      (_this)._f8 = f8;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f8;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f6 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    set f6(value) {
      let _this = this;
      _this._f6 = value;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f5, f6) {
      let _this = this;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm0(p0, p1, p2, globalState) {
      let _this = this;
      return (new BigNumber(188)).isLessThan(_this.f6);
    };
    fm1(globalState) {
      let _this = this;
      return ((function () {
        let _coll26 = new _dafny.Set();
        for (const _compr_26 of _dafny.IntegerRange(new BigNumber(110), new BigNumber(681))) {
          let _358_v0 = _compr_26;
          if (((new BigNumber(110)).isLessThanOrEqualTo(_358_v0)) && ((_358_v0).isLessThan(new BigNumber(681)))) {
            _coll26.add((_358_v0).multipliedBy(_this.f6));
          }
        }
        return _coll26;
      }()).Difference(function () {
        let _coll27 = new _dafny.Set();
        for (const _compr_27 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_359_i0) {
          return _this.f6;
        })).length), _this.f6)).Elements) {
          let _360_v1 = _compr_27;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_359_i0) {
            return _this.f6;
          })).length), _this.f6), _360_v1)) {
            _coll27.add(_module.__default.safeDivisionInt(_360_v1, new BigNumber(499)));
          }
        }
        return _coll27;
      }())).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(231)));
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return _this.f6;
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(((_dafny.Set.fromElements((_this).f5)).Intersect(_dafny.Set.fromElements((_this).f5))).length);
    };
    fm13(p0, p1, globalState) {
      let _this = this;
      return (_this).f5;
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _361_v0;
      let _init6 = function (_362_i0) {
        return _dafny.Seq.UnicodeFromString("yjwhec");
      };
      let _nw69 = Array((_dafny.ONE).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw69.length); _i0_6++) {
        _nw69[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _361_v0 = _nw69;
      let _363_v1;
      _363_v1 = _dafny.Seq.UnicodeFromString("oll");
      let _index50 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length));
      (_361_v0)[_index50] = _363_v1;
      let _364_v2;
      _364_v2 = new _dafny.CodePoint('r'.codePointAt(0));
      let _index51 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length));
      let _rhs88 = _dafny.Seq.UnicodeFromString("alijtn");
      let _rhs89 = _364_v2;
      let _lhs49 = _361_v0;
      let _lhs50 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length));
      _lhs49[_lhs50] = _rhs88;
      _364_v2 = _rhs89;
      let _365_v3;
      let _nw70 = Array((new BigNumber(29)).toNumber()).fill(false);
      _365_v3 = _nw70;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_365_v3).length))) {
        let _366_i1 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_366_i1)) && ((_366_i1).isLessThan(new BigNumber((_365_v3).length))))) {
          (_365_v3)[(_366_i1)] = !(!(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-399)), ((_367_v2) => function (_368_i2) {
            return _367_v2;
          })(_364_v2)), _dafny.Seq.UnicodeFromString("djet")))) || (false);
        }
      }
      let _369_v4;
      let _nw71 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _369_v4 = _nw71;
      let _370_v5;
      _370_v5 = _dafny.Seq.of(_369_v4);
      let _371_v6;
      _371_v6 = _dafny.Set.fromElements((_370_v5)[_module.__default.safeIndex((_this).f5, new BigNumber((_370_v5).length))]);
      _371_v6 = (_dafny.Set.fromElements(_369_v4, _369_v4)).Intersect(_371_v6);
      let _index52 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length));
      (_369_v4)[_index52] = (new BigNumber(85)).multipliedBy((_this).f5);
      let _index53 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length));
      (_369_v4)[_index53] = (_this).f5;
      let _372_v7;
      _372_v7 = _dafny.Seq.of((_this).f5);
      let _373_v8;
      _373_v8 = _dafny.Set.fromElements(_372_v7);
      let _374_v9;
      _374_v9 = true;
      let _375_i3;
      _375_i3 = _dafny.ZERO;
      L2: {
        while (((_373_v8).Intersect(_373_v8)).IsProperSubsetOf(_module.__default.fm14(new BigNumber((_dafny.Seq.UnicodeFromString("iwogcvy")).length), _374_v9, globalState))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_375_i3)) {
              break L2;
            }
            _375_i3 = (_375_i3).plus(_dafny.ONE);
            (_this).f6 = new BigNumber(131);
            r1 = !(!(_374_v9) || (_374_v9));
            let _376_v10;
            _376_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f6);
            let _377_v11;
            _377_v11 = _dafny.Map.Empty.slice().updateUnsafe((_module.D1.create_DC6(_dafny.Seq.UnicodeFromString("tyctsial"), (_this).f5)).dtor_cf6,_dafny.Seq.of(_376_v10, _376_v10));
            _377_v11 = (_377_v11).update((_this).f5, _module.__default.fm15(_364_v2, globalState));
            let _index54 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_365_v3).length));
            (_365_v3)[_index54] = _374_v9;
            let _index55 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_365_v3).length));
            (_365_v3)[_index55] = false;
          }
        }
      }
      if (_dafny.areEqual(_372_v7, _dafny.Seq.of(_this.f6))) {
        let _index56 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length));
        (_365_v3)[_index56] = !(_374_v9);
        let _index57 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length));
        (_365_v3)[_index57] = _374_v9;
        let _378_v12;
        _378_v12 = _module.D0.create_DC0(!(false));
        let _source6 = _378_v12;
        if (_source6.is_DC1) {
          let _379___mcc_h0 = (_source6).cf1;
          let _380___mcc_h1 = (_source6).cf2;
          let _381_cf2 = _380___mcc_h1;
          let _382_cf1 = _379___mcc_h0;
          (_this).f6 = (_this).f5;
          let _index58 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length));
          (_365_v3)[_index58] = _374_v9;
          let _383_v13;
          let _nw72 = new _module.C0();
          _nw72.__ctor(new BigNumber(85));
          _383_v13 = _nw72;
          let _384_v14;
          let _nw73 = new _module.C0();
          _nw73.__ctor((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          _384_v14 = _nw73;
        } else if (_source6.is_DC2) {
          _363_v1 = (_361_v0)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length))];
          let _385_v15;
          _385_v15 = _dafny.Seq.of(_374_v9);
          let _386_v16;
          _386_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-348),!((_385_v15)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f5), new BigNumber((_385_v15).length))]));
          let _387_v18;
          let _nw74 = new _module.C0();
          _nw74.__ctor(_this.f6);
          _387_v18 = _nw74;
          let _388_v19;
          _388_v19 = _dafny.Map.Empty.slice().updateUnsafe(_387_v18,(_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          let _389_v20;
          _389_v20 = _dafny.Set.fromElements(new BigNumber((_388_v19).length));
          let _390_v21;
          _390_v21 = _dafny.MultiSet.fromElements(_374_v9, _374_v9, (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))]);
          _374_v9 = ((_dafny.MultiSet.fromElements(false)).Union(_390_v21)).IsProperSubsetOf(_module.__default.fm16(_386_v16, (_this).f5, function () {
            let _coll28 = new _dafny.Map();
            for (const _compr_28 of (_389_v20).Elements) {
              let _391_v17 = _compr_28;
              if ((_389_v20).contains(_391_v17)) {
                _coll28.push([_module.__default.safeModuloInt(_391_v17, (_387_v18).f8),false]);
              }
            }
            return _coll28;
          }(), globalState));
          let _392_v22;
          let _init7 = ((_393_v3) => function (_394_i4) {
            return (_393_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_393_v3).length))];
          })(_365_v3);
          let _nw75 = Array((new BigNumber(11)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw75.length); _i0_7++) {
            _nw75[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _392_v22 = _nw75;
          let _395_v23;
          _395_v23 = _dafny.Map.Empty.slice().updateUnsafe(_387_v18,_392_v22);
          (globalState).f2 = _module.__default.safeDivisionInt(_this.f6, (_this).fm12(_374_v9, new BigNumber((_395_v23).length), (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))], new BigNumber((_module.__default.fm17((_387_v18).f8, _module.D0.create_DC0(_374_v9), _this.f6, true, globalState)).length), globalState));
          let _396_v24;
          let _nw76 = Array((new BigNumber(11)).toNumber()).fill([]);
          _396_v24 = _nw76;
          let _397_v25;
          _397_v25 = _module.D0.create_DC1(false, (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))]);
          let _pat_let_tv6 = _374_v9;
          let _index59 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length));
          let _index60 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length));
          let _rhs90 = (_this.f6).isLessThanOrEqualTo(_module.__default.safeModuloInt((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))], (_this).f5));
          let _rhs91 = (((function (_pat_let7_0) {
            return function (_398_dt__update__tmp_h0) {
              return function (_pat_let8_0) {
                return function (_399_dt__update_hcf2_h0) {
                  return _module.D0.create_DC1((_398_dt__update__tmp_h0).dtor_cf1, _399_dt__update_hcf2_h0);
                }(_pat_let8_0);
              }(_pat_let_tv6);
            }(_pat_let7_0);
          }(_397_v25)).dtor_cf2) ? (_396_v24) : (_396_v24));
          let _rhs92 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]), _this.f6);
          let _rhs93 = _392_v22;
          let _rhs94 = _372_v7;
          let _lhs51 = _365_v3;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length));
          let _lhs53 = _369_v4;
          let _lhs54 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length));
          _lhs51[_lhs52] = _rhs90;
          _396_v24 = _rhs91;
          _lhs53[_lhs54] = _rhs92;
          _392_v22 = _rhs93;
          _372_v7 = _rhs94;
        } else if (_source6.is_DC3) {
          let _400___mcc_h2 = (_source6).cf3;
          let _401_cf3 = _400___mcc_h2;
          _401_cf3 = (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))];
          _363_v1 = (_361_v0)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length))];
          let _init8 = function (_402_i5) {
            return _module.__default.safeModuloInt(_402_i5, (_this).f5);
          };
          let _nw77 = Array((new BigNumber(27)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw77.length); _i0_8++) {
            _nw77[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _369_v4 = _nw77;
          let _403_v26;
          let _init9 = ((_404_v9) => function (_405_i6) {
            return _404_v9;
          })(_374_v9);
          let _nw78 = Array((new BigNumber(18)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw78.length); _i0_9++) {
            _nw78[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _403_v26 = _nw78;
          let _406_v27;
          _406_v27 = _dafny.Map.Empty.slice().updateUnsafe(_401_cf3,(_this).f5);
          let _407_v28;
          _407_v28 = _dafny.Map.Empty.slice().updateUnsafe(_403_v26,_dafny.Map.Empty.slice().updateUnsafe(_401_cf3,(_this).f5));
          let _408_v29;
          let _nw79 = Array((new BigNumber(3)).toNumber());
          _nw79[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_403_v26,_dafny.Map.Empty.slice().updateUnsafe(true,_this.f6))).update(_403_v26, _406_v27);
          _nw79[(_dafny.ONE).toNumber()] = _407_v28;
          _nw79[(new BigNumber(2)).toNumber()] = (_407_v28).Merge(_407_v28);
          _408_v29 = _nw79;
          let _index61 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_408_v29).length));
          (_408_v29)[_index61] = _407_v28;
          let _409_v30;
          _409_v30 = _dafny.Map.Empty.slice().updateUnsafe(_374_v9,_374_v9);
          let _410_v31;
          _410_v31 = _module.D3.create_DC8(_409_v30);
          let _index62 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_408_v29).length));
          let _index63 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length));
          let _rhs95 = new BigNumber(((_410_v31).dtor_cf8).length);
          let _rhs96 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber(((_361_v0)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length))]).length), _this.f6), new BigNumber((_dafny.MultiSet.fromElements((_361_v0)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length))], _363_v1)).cardinality())));
          let _rhs97 = (_dafny.Map.Empty.slice().updateUnsafe(_403_v26,_406_v27)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_403_v26,_406_v27));
          let _rhs98 = (_this.f6).isLessThanOrEqualTo((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          let _lhs55 = _this;
          let _lhs56 = globalState;
          let _lhs57 = _408_v29;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_408_v29).length));
          let _lhs59 = _365_v3;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length));
          _lhs55.f6 = _rhs95;
          _lhs56.f2 = _rhs96;
          _lhs57[_lhs58] = _rhs97;
          _lhs59[_lhs60] = _rhs98;
        } else {
          let _411___mcc_h3 = (_source6).cf0;
          let _412_cf0 = _411___mcc_h3;
          let _index64 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length));
          (_369_v4)[_index64] = ((_dafny.ZERO).minus(_this.f6)).minus(_this.f6);
          let _413_v32;
          _413_v32 = _dafny.MultiSet.fromElements(_372_v7);
          _413_v32 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))], (_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]), _dafny.Seq.of((_this).fm12(_374_v9, (_this).f5, _374_v9, (_this).fm12((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))], (_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))], _374_v9, _this.f6, globalState), globalState)), (_module.D4.create_DC11(_372_v7)).dtor_cf13, _372_v7);
          let _414_v33;
          let _nw80 = new _module.C0();
          _nw80.__ctor((_this).f5);
          _414_v33 = _nw80;
          let _415_v34;
          _415_v34 = _dafny.Seq.of(_412_cf0);
          let _416_v35;
          _416_v35 = _dafny.MultiSet.fromElements(new BigNumber((_415_v34).length), new BigNumber(979), new BigNumber(490));
          _416_v35 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_372_v7, _dafny.Seq.Concat(_372_v7, _372_v7)));
        }
        let _417_v36;
        _417_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(134),_374_v9);
        let _418_v37;
        _418_v37 = _dafny.MultiSet.fromElements(_417_v36);
        let _419_v38;
        _419_v38 = _dafny.Map.Empty.slice().updateUnsafe(false,(_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))]);
        let _420_v39;
        _420_v39 = _dafny.Set.fromElements((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))]);
        let _421_v40;
        let _nw81 = Array((new BigNumber(21)).toNumber());
        _nw81[(_dafny.ZERO).toNumber()] = _374_v9;
        _nw81[(_dafny.ONE).toNumber()] = _374_v9;
        _nw81[(new BigNumber(2)).toNumber()] = !((_this).fm0(_418_v37, true, _363_v1, globalState));
        _nw81[(new BigNumber(3)).toNumber()] = false;
        _nw81[(new BigNumber(4)).toNumber()] = !(true);
        _nw81[(new BigNumber(5)).toNumber()] = (_this).fm1(globalState);
        _nw81[(new BigNumber(6)).toNumber()] = !(_419_v38).contains(_374_v9);
        _nw81[(new BigNumber(7)).toNumber()] = _374_v9;
        _nw81[(new BigNumber(8)).toNumber()] = _374_v9;
        _nw81[(new BigNumber(9)).toNumber()] = (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))];
        _nw81[(new BigNumber(10)).toNumber()] = _374_v9;
        _nw81[(new BigNumber(11)).toNumber()] = (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))];
        _nw81[(new BigNumber(12)).toNumber()] = _374_v9;
        _nw81[(new BigNumber(13)).toNumber()] = (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))];
        _nw81[(new BigNumber(14)).toNumber()] = _374_v9;
        _nw81[(new BigNumber(15)).toNumber()] = (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))];
        _nw81[(new BigNumber(16)).toNumber()] = (_420_v39).IsSubsetOf(_420_v39);
        _nw81[(new BigNumber(17)).toNumber()] = (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))];
        _nw81[(new BigNumber(18)).toNumber()] = false;
        _nw81[(new BigNumber(19)).toNumber()] = _374_v9;
        _nw81[(new BigNumber(20)).toNumber()] = _374_v9;
        _421_v40 = _nw81;
        _421_v40 = _421_v40;
        let _422_v41;
        let _nw82 = Array((new BigNumber(15)).toNumber()).fill([]);
        _422_v41 = _nw82;
        let _index65 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_422_v41).length));
        (_422_v41)[_index65] = _421_v40;
        let _index66 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_422_v41).length));
        (_422_v41)[_index66] = _421_v40;
        let _423_v42;
        _423_v42 = _dafny.Map.Empty.slice().updateUnsafe(_369_v4,(_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
        let _424_v43;
        _424_v43 = _dafny.MultiSet.fromElements(_this.f6);
        let _425_v44;
        _425_v44 = _dafny.Map.Empty.slice().updateUnsafe(_423_v42,_424_v43);
        let _426_v45;
        _426_v45 = _module.D3.create_DC9((((_425_v44).contains((_423_v42).update(_369_v4, new BigNumber(-664)))) ? ((_425_v44).get((_423_v42).update(_369_v4, new BigNumber(-664)))) : (_424_v43)), _module.D0.create_DC2(), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(29)), ((_427_v2) => function (_428_i7) {
  return _427_v2;
})(_364_v2))).length));
        let _source7 = _426_v45;
        if (_source7.is_DC9) {
          let _429___mcc_h4 = (_source7).cf9;
          let _430___mcc_h5 = (_source7).cf10;
          let _431___mcc_h6 = (_source7).cf11;
          let _432_cf11 = _431___mcc_h6;
          let _433_cf10 = _430___mcc_h5;
          let _434_cf9 = _429___mcc_h4;
          _369_v4 = _369_v4;
          let _435_v46;
          _435_v46 = _dafny.Map.Empty.slice().updateUnsafe(_374_v9,_372_v7);
          let _436_v47;
          let _nw83 = new _module.C0();
          _nw83.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), ((_437_v2) => function (_438_i8) {
            return _437_v2;
          })(_364_v2)),new BigNumber(((((_435_v46).contains((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))])) ? ((_435_v46).get((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))])) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(10)), function (_439_i9) {
            return (_dafny.ZERO).minus(_this.f6);
          })))).length))).length));
          _436_v47 = _nw83;
          let _index67 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length));
          (_361_v0)[_index67] = _363_v1;
          (globalState).f2 = (_436_v47).f8;
        } else if (_source7.is_DC8) {
          let _440___mcc_h7 = (_source7).cf8;
          let _441_cf8 = _440___mcc_h7;
          let _index68 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length));
          (_369_v4)[_index68] = (_dafny.ZERO).minus((_this).fm12(true, (_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))], (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))], _this.f6, globalState));
          _374_v9 = (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))];
          let _442_v48;
          let _nw84 = new _module.C0();
          _nw84.__ctor((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          _442_v48 = _nw84;
          let _443_v49;
          _443_v49 = _dafny.Map.Empty.slice().updateUnsafe(!(_374_v9) || ((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))]),_module.__default.safeDivisionInt((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))], (_372_v7)[_module.__default.safeIndex((_this).f5, new BigNumber((_372_v7).length))]));
          let _rhs99 = (_dafny.ZERO).minus(_this.f6);
          let _rhs100 = _443_v49;
          let _rhs101 = _364_v2;
          let _lhs61 = globalState;
          _lhs61.f2 = _rhs99;
          _443_v49 = _rhs100;
          _364_v2 = _rhs101;
        } else {
          let _444___mcc_h8 = (_source7).cf12;
          let _445_cf12 = _444___mcc_h8;
          let _446_v50;
          _446_v50 = _dafny.Seq.of(_374_v9);
          _446_v50 = _446_v50;
          let _rhs102 = _369_v4;
          let _rhs103 = !(_424_v43).contains((_dafny.ZERO).minus((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]));
          _369_v4 = _rhs102;
          _374_v9 = _rhs103;
          let _447_v51;
          let _nw85 = Array((new BigNumber(23)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _364_v2;
          _nw85[(_dafny.ONE).toNumber()] = _364_v2;
          _nw85[(new BigNumber(2)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(3)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(4)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(5)).toNumber()] = _module.__default.fm18((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))], (_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))], _this.f6, (_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))], globalState);
          _nw85[(new BigNumber(6)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(7)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(8)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(9)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(10)).toNumber()] = (((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))]) ? (_364_v2) : (_module.__default.fm18((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))], _this.f6, new BigNumber(-788), _374_v9, globalState)));
          _nw85[(new BigNumber(11)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(12)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(13)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(14)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(15)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(16)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(17)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(18)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(19)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(20)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(21)).toNumber()] = _364_v2;
          _nw85[(new BigNumber(22)).toNumber()] = _364_v2;
          _447_v51 = _nw85;
          _447_v51 = _447_v51;
          let _448_v52;
          _448_v52 = _module.D0.create_DC3((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))]);
          let _449_v53;
          let _nw86 = Array((new BigNumber(10)).toNumber());
          _nw86[(_dafny.ZERO).toNumber()] = _448_v52;
          _nw86[(_dafny.ONE).toNumber()] = _448_v52;
          _nw86[(new BigNumber(2)).toNumber()] = _448_v52;
          _nw86[(new BigNumber(3)).toNumber()] = _module.D0.create_DC3(_374_v9);
          _nw86[(new BigNumber(4)).toNumber()] = _448_v52;
          _nw86[(new BigNumber(5)).toNumber()] = _448_v52;
          _nw86[(new BigNumber(6)).toNumber()] = _module.D0.create_DC3((_365_v3)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_365_v3).length))]);
          _nw86[(new BigNumber(7)).toNumber()] = _448_v52;
          _nw86[(new BigNumber(8)).toNumber()] = _448_v52;
          _nw86[(new BigNumber(9)).toNumber()] = _448_v52;
          _449_v53 = _nw86;
          let _index69 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_449_v53).length));
          (_449_v53)[_index69] = _448_v52;
          let _index70 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_449_v53).length));
          (_449_v53)[_index70] = _448_v52;
        }
      } else {
        let _index71 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length));
        (_361_v0)[_index71] = (_361_v0)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length))];
        (globalState).f2 = _module.__default.safeModuloInt((_this).fm13(_this.f6, _dafny.MultiSet.fromElements(_363_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-505)), ((_450_v2) => function (_451_i10) {
          return _450_v2;
        })(_364_v2)), _363_v1, (_361_v0)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length))]), globalState), new BigNumber(393));
        if (_374_v9) {
          let _452_v54;
          _452_v54 = _dafny.Seq.of(((_374_v9) ? (_363_v1) : (_dafny.Seq.UnicodeFromString("ltkjnhdi"))));
          let _453_v55;
          _453_v55 = _dafny.Map.Empty.slice().updateUnsafe(_374_v9,_374_v9);
          let _index72 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length));
          (_369_v4)[_index72] = new BigNumber((_dafny.Seq.update(_452_v54, _module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_453_v55).length), new BigNumber(835)), new BigNumber((_452_v54).length)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(145)), ((_454_v2) => function (_455_i11) {
            return _454_v2;
          })(_364_v2)), (_361_v0)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length))]))).length);
          let _index73 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_365_v3).length));
          (_365_v3)[_index73] = !((_this).f5).isEqualTo((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          let _index74 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_365_v3).length));
          (_365_v3)[_index74] = (_this).fm1(globalState);
          let _456_v56;
          let _nw87 = new _module.C0();
          _nw87.__ctor((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          _456_v56 = _nw87;
          let _457_v57;
          _457_v57 = _dafny.MultiSet.fromElements((_456_v56).f8, (_this).f5, _this.f6, (_dafny.ZERO).minus(new BigNumber((_363_v1).length)), (_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          let _458_v58;
          _458_v58 = _dafny.Map.Empty.slice().updateUnsafe(_363_v1,(_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          let _459_v59;
          _459_v59 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,(_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          let _460_v60;
          _460_v60 = _dafny.Map.Empty.slice().updateUnsafe((_365_v3)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_365_v3).length))],_dafny.MultiSet.FromArray(_372_v7));
          let _461_v61;
          let _nw88 = Array((new BigNumber(29)).toNumber());
          _nw88[(_dafny.ZERO).toNumber()] = _457_v57;
          _nw88[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements((_this).f5);
          _nw88[(new BigNumber(2)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(3)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(4)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(5)).toNumber()] = (_dafny.MultiSet.FromArray(_372_v7)).Difference(_457_v57);
          _nw88[(new BigNumber(6)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.FromArray(_372_v7)).Union(_457_v57);
          _nw88[(new BigNumber(8)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(9)).toNumber()] = (_module.__default.fm19((_365_v3)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_365_v3).length))], (_456_v56).f8, (_365_v3)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_365_v3).length))], globalState)).update(_this.f6, _module.__default.abs((((_458_v58).contains((_361_v0)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length))])) ? ((_458_v58).get((_361_v0)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_361_v0).length))])) : ((_this).f5))));
          _nw88[(new BigNumber(10)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(11)).toNumber()] = (_dafny.MultiSet.fromElements((_this).f5)).update((((_459_v59).contains((_this).f5)) ? ((_459_v59).get((_this).f5)) : (_this.f6)), _module.__default.abs(new BigNumber(690)));
          _nw88[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_363_v1).length)), new BigNumber((_457_v57).cardinality()));
          _nw88[(new BigNumber(13)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(14)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(15)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(16)).toNumber()] = (_457_v57).Union((((_460_v60).contains(false)) ? ((_460_v60).get(false)) : (_457_v57)));
          _nw88[(new BigNumber(17)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(18)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(19)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(20)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(21)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(22)).toNumber()] = (_457_v57).Union(_457_v57);
          _nw88[(new BigNumber(23)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(24)).toNumber()] = _dafny.MultiSet.fromElements((_this).f5);
          _nw88[(new BigNumber(25)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(26)).toNumber()] = _457_v57;
          _nw88[(new BigNumber(27)).toNumber()] = (((_365_v3)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_365_v3).length))]) ? (_457_v57) : (_dafny.MultiSet.FromArray(_372_v7)));
          _nw88[(new BigNumber(28)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_372_v7, _dafny.Seq.of((_456_v56).f8)));
          _461_v61 = _nw88;
          let _index75 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_461_v61).length));
          (_461_v61)[_index75] = (_457_v57).Intersect(_457_v57);
          let _index76 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_461_v61).length));
          let _index77 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_365_v3).length));
          let _rhs104 = _457_v57;
          let _rhs105 = !(!(_374_v9)) || (!(_374_v9) || (!(!(_374_v9))));
          let _rhs106 = _dafny.Seq.of(_363_v1);
          let _lhs62 = _461_v61;
          let _lhs63 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_461_v61).length));
          let _lhs64 = _365_v3;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_365_v3).length));
          _lhs62[_lhs63] = _rhs104;
          _lhs64[_lhs65] = _rhs105;
          _452_v54 = _rhs106;
          let _462_v62;
          let _463_v63;
          let _464_v64;
          let _465_v65;
          let _out28;
          let _out29;
          let _out30;
          let _out31;
          let _outcollector9 = _module.__default.m0(_364_v2, globalState);
          _out28 = _outcollector9[0];
          _out29 = _outcollector9[1];
          _out30 = _outcollector9[2];
          _out31 = _outcollector9[3];
          _462_v62 = _out28;
          _463_v63 = _out29;
          _464_v64 = _out30;
          _465_v65 = _out31;
        } else {
          let _466_v66;
          let _nw89 = new _module.C0();
          _nw89.__ctor((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))]);
          _466_v66 = _nw89;
          let _467_v67;
          _467_v67 = _module.D4.create_DC14((_369_v4)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_369_v4).length))], false, (_this).f5, _374_v9, (_466_v66).f8);
          (globalState).f2 = (_467_v67).dtor_cf22;
          let _468_v68;
          let _469_v69;
          let _470_v70;
          let _471_v71;
          let _out32;
          let _out33;
          let _out34;
          let _out35;
          let _outcollector10 = _module.__default.m0(new _dafny.CodePoint('p'.codePointAt(0)), globalState);
          _out32 = _outcollector10[0];
          _out33 = _outcollector10[1];
          _out34 = _outcollector10[2];
          _out35 = _outcollector10[3];
          _468_v68 = _out32;
          _469_v69 = _out33;
          _470_v70 = _out34;
          _471_v71 = _out35;
          r1 = _469_v69;
          r1 = _374_v9;
        }
        r1 = _374_v9;
        _374_v9 = !(_374_v9);
      }
      r0 = _module.__default.safeModuloInt((_this).f5, (_this).f5);
      r1 = _374_v9;
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let _472_v0;
      let _nw90 = new _module.C0();
      _nw90.__ctor(_this.f6);
      _472_v0 = _nw90;
      let _473_v1;
      let _init10 = ((_474_v0) => function (_475_i1) {
        return _module.__default.safeModuloInt(_475_i1, (_474_v0).f8);
      })(_472_v0);
      let _nw91 = Array((new BigNumber(27)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw91.length); _i0_10++) {
        _nw91[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _473_v1 = _nw91;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_473_v1).length))) {
        let _476_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_476_i0)) && ((_476_i0).isLessThan(new BigNumber((_473_v1).length))))) {
          (_473_v1)[(_476_i0)] = (_476_i0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-155)), function (_477_i2) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          })).length));
        }
      }
      let _478_v2;
      _478_v2 = true;
      (_this).f6 = ((_478_v2) ? ((_472_v0).f8) : (_this.f6));
      let _479_v3;
      let _init11 = ((_480_v0) => function (_481_i4) {
        return ((_480_v0).f8).isEqualTo(_this.f6);
      })(_472_v0);
      let _nw92 = Array((new BigNumber(19)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw92.length); _i0_11++) {
        _nw92[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _479_v3 = _nw92;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_479_v3).length))) {
        let _482_i3 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_482_i3)) && ((_482_i3).isLessThan(new BigNumber((_479_v3).length))))) {
          (_479_v3)[(_482_i3)] = _478_v2;
        }
      }
      let _483_v4;
      let _init12 = ((_484_v2, _485_v0) => function (_486_i6) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_484_v2,(_485_v0).f8)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_484_v2),(_485_v0).f8));
      })(_478_v2, _472_v0);
      let _nw93 = Array((new BigNumber(8)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw93.length); _i0_12++) {
        _nw93[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _483_v4 = _nw93;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_483_v4).length))) {
        let _487_i5 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_487_i5)) && ((_487_i5).isLessThan(new BigNumber((_483_v4).length))))) {
          (_483_v4)[(_487_i5)] = _dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((((!(_478_v2)) ? (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber((_dafny.Seq.of(_module.D1.create_DC4(_dafny.Set.fromElements((_dafny.ZERO).minus((_this).f5))), _module.D1.create_DC4(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("bof")).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_478_v2,_478_v2)).length), (_dafny.ZERO).minus((_472_v0).f8), (_472_v0).f8, _this.f6)))).length))) : (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),(_dafny.ZERO).minus((_472_v0).f8))))).length));
        }
      }
      let _488_v5;
      _488_v5 = new _dafny.CodePoint('m'.codePointAt(0));
      let _489_v6;
      _489_v6 = _dafny.Set.fromElements(_478_v2, _478_v2, true, _478_v2, _478_v2);
      let _rhs107 = _479_v3;
      let _rhs108 = ((true) ? (_488_v5) : (_488_v5));
      let _rhs109 = (_dafny.ZERO).minus(((_478_v2) ? ((_this).fm2((_472_v0).f8, _489_v6, globalState)) : ((_472_v0).f8)));
      let _rhs110 = (_472_v0).f8;
      let _rhs111 = (_472_v0).f8;
      let _lhs66 = globalState;
      let _lhs67 = _this;
      let _lhs68 = _this;
      _479_v3 = _rhs107;
      _488_v5 = _rhs108;
      _lhs66.f2 = _rhs109;
      _lhs67.f6 = _rhs110;
      _lhs68.f6 = _rhs111;
      return;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f6 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    set f6(value) {
      let _this = this;
      _this._f6 = value;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f5, f6) {
      let _this = this;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm0(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(-363), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(true, true)).length))).length)))).IsDisjointFrom((_dafny.MultiSet.fromElements(_dafny.Seq.of(_this.f6))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(567)), function (_490_i0) {
        return _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, false)).length), new BigNumber(-401));
      }))));
    };
    fm1(globalState) {
      let _this = this;
      return _dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(874)), function (_491_i0) {
        return (_this).f5;
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), function (_492_i1) {
        return (_this).f5;
      })), (_this).f5);
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(-302);
    };
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("j")).length))).isLessThan((_dafny.ZERO).minus((_this).f5));
    };
    fm9(p0, p1, globalState) {
      let _this = this;
      return true;
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _493_v0;
      _493_v0 = true;
      r1 = _493_v0;
      let _494_v1;
      _494_v1 = new _dafny.CodePoint('q'.codePointAt(0));
      let _495_v2;
      let _496_v3;
      let _497_v4;
      let _498_v5;
      let _out36;
      let _out37;
      let _out38;
      let _out39;
      let _outcollector11 = _module.__default.m0(_494_v1, globalState);
      _out36 = _outcollector11[0];
      _out37 = _outcollector11[1];
      _out38 = _outcollector11[2];
      _out39 = _outcollector11[3];
      _495_v2 = _out36;
      _496_v3 = _out37;
      _497_v4 = _out38;
      _498_v5 = _out39;
      let _index78 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length));
      (_495_v2)[_index78] = _498_v5;
      let _499_v6;
      _499_v6 = _module.D1.create_DC5();
      let _500_v7;
      _500_v7 = _dafny.MultiSet.fromElements(_496_v3);
      let _index79 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length));
      (_495_v2)[_index79] = (((_this).fm9(_499_v6, _493_v0, globalState)) ? (new BigNumber((_500_v7).cardinality())) : ((_this).f5));
      let _hi4 = _module.__default.safeModuloInt(new BigNumber(247), _498_v5);
      for (let _501_i0 = (new BigNumber(282)).plus((_dafny.ZERO).minus(_498_v5)); _501_i0.isLessThan(_hi4); _501_i0 = _501_i0.plus(_dafny.ONE)) {
        let _502_v8;
        _502_v8 = _dafny.Seq.of(!(_this.f6).isEqualTo(_501_i0));
        if ((_502_v8)[_module.__default.safeIndex((_this.f6).minus(_501_i0), new BigNumber((_502_v8).length))]) {
          let _503_v9;
          _503_v9 = _dafny.MultiSet.fromElements(_501_i0);
          let _504_v10;
          let _nw94 = new _module.C0();
          _nw94.__ctor(new BigNumber((_503_v9).cardinality()));
          _504_v10 = _nw94;
          let _505_v11;
          _505_v11 = _dafny.Seq.UnicodeFromString("ij");
          let _506_v12;
          _506_v12 = _dafny.Map.Empty.slice().updateUnsafe(_493_v0,_505_v11);
          _505_v11 = _dafny.Seq.Concat(_dafny.Seq.update((((_506_v12).contains(_496_v3)) ? ((_506_v12).get(_496_v3)) : (_505_v11)), _module.__default.safeIndex((_this).f5, new BigNumber(((((_506_v12).contains(_496_v3)) ? ((_506_v12).get(_496_v3)) : (_505_v11))).length)), _494_v1), _505_v11);
          let _507_v13;
          _507_v13 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_498_v5),_493_v0);
          let _508_v14;
          _508_v14 = _dafny.MultiSet.fromElements(_507_v13, _507_v13, _507_v13);
          let _509_v15;
          let _nw95 = Array((new BigNumber(8)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = !((((_507_v13).contains(new BigNumber((_dafny.Seq.update(_502_v8, _module.__default.safeIndex((_495_v2)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length))], new BigNumber((_502_v8).length)), _496_v3)).length))) ? ((_507_v13).get(new BigNumber((_dafny.Seq.update(_502_v8, _module.__default.safeIndex((_495_v2)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length))], new BigNumber((_502_v8).length)), _496_v3)).length))) : (true))) || (_496_v3);
          _nw95[(_dafny.ONE).toNumber()] = _496_v3;
          _nw95[(new BigNumber(2)).toNumber()] = ((_496_v3) ? (false) : (_493_v0));
          _nw95[(new BigNumber(3)).toNumber()] = _496_v3;
          _nw95[(new BigNumber(4)).toNumber()] = (_this).fm8(_505_v11, _496_v3, (_this).fm0(_508_v14, _496_v3, _505_v11, globalState), _493_v0, globalState);
          _nw95[(new BigNumber(5)).toNumber()] = _496_v3;
          _nw95[(new BigNumber(6)).toNumber()] = false;
          _nw95[(new BigNumber(7)).toNumber()] = _496_v3;
          _509_v15 = _nw95;
          let _index80 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_509_v15).length));
          (_509_v15)[_index80] = _493_v0;
          let _index81 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_509_v15).length));
          (_509_v15)[_index81] = !((_503_v9).Union(_503_v9)).equals(_503_v9);
          let _510_v16;
          _510_v16 = _dafny.Set.fromElements(_496_v3, _493_v0, _493_v0);
          let _511_v17;
          let _nw96 = new _module.C0();
          _nw96.__ctor(((_493_v0) ? (new BigNumber((_510_v16).length)) : (_501_i0)));
          _511_v17 = _nw96;
          _496_v3 = (_509_v15)[_module.__default.safeIndex(new BigNumber(868), new BigNumber((_509_v15).length))];
        } else {
          let _512_v18;
          let _513_v19;
          let _514_v20;
          let _515_v21;
          let _out40;
          let _out41;
          let _out42;
          let _out43;
          let _outcollector12 = _module.__default.m0(_494_v1, globalState);
          _out40 = _outcollector12[0];
          _out41 = _outcollector12[1];
          _out42 = _outcollector12[2];
          _out43 = _outcollector12[3];
          _512_v18 = _out40;
          _513_v19 = _out41;
          _514_v20 = _out42;
          _515_v21 = _out43;
          r1 = _493_v0;
          let _516_v22;
          let _nw97 = new _module.C0();
          _nw97.__ctor((_495_v2)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length))]);
          _516_v22 = _nw97;
          let _517_v23;
          _517_v23 = _dafny.Seq.of((_516_v22).f8, _501_i0);
          let _518_v24;
          _518_v24 = _dafny.Seq.UnicodeFromString("gygdbkdr");
          let _519_v25;
          _519_v25 = _dafny.Set.fromElements(_501_i0);
          let _520_v26;
          let _nw98 = Array((new BigNumber(12)).toNumber());
          _nw98[(_dafny.ZERO).toNumber()] = true;
          _nw98[(_dafny.ONE).toNumber()] = ((_495_v2)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length))]).isLessThanOrEqualTo((_517_v23)[_module.__default.safeIndex(_501_i0, new BigNumber((_517_v23).length))]);
          _nw98[(new BigNumber(2)).toNumber()] = (false) && (_496_v3);
          _nw98[(new BigNumber(3)).toNumber()] = (_this).fm8(_518_v24, (_502_v8)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-830)), function (_521_i1) {
            return new BigNumber(163);
          })).length), new BigNumber((_502_v8).length))], _513_v19, _496_v3, globalState);
          _nw98[(new BigNumber(4)).toNumber()] = (_dafny.Set.fromElements(new BigNumber((_500_v7).cardinality()), new BigNumber((_502_v8).length))).IsDisjointFrom(_519_v25);
          _nw98[(new BigNumber(5)).toNumber()] = _493_v0;
          _nw98[(new BigNumber(6)).toNumber()] = (_this).fm9(_499_v6, _513_v19, globalState);
          _nw98[(new BigNumber(7)).toNumber()] = (!(_496_v3)) === (_496_v3);
          _nw98[(new BigNumber(8)).toNumber()] = !(_493_v0);
          _nw98[(new BigNumber(9)).toNumber()] = _493_v0;
          _nw98[(new BigNumber(10)).toNumber()] = _513_v19;
          _nw98[(new BigNumber(11)).toNumber()] = !(_513_v19);
          _520_v26 = _nw98;
          let _nw99 = Array((new BigNumber(12)).toNumber()).fill(false);
          _520_v26 = _nw99;
          _496_v3 = !(true);
        }
        let _522_v27;
        _522_v27 = _dafny.MultiSet.fromElements(_501_i0, new BigNumber((_502_v8).length));
        let _523_v28;
        _523_v28 = _dafny.Set.fromElements(_522_v27);
        let _524_v29;
        _524_v29 = _dafny.Map.Empty.slice().updateUnsafe(_498_v5,_523_v28);
        let _525_v30;
        _525_v30 = _dafny.Seq.of(_522_v27);
        _524_v29 = (_524_v29).update((_this).f5, function () {
          let _coll29 = new _dafny.Set();
          for (const _compr_29 of (_525_v30).Elements) {
            let _526_v31 = _compr_29;
            if (_dafny.Seq.contains(_525_v30, _526_v31)) {
              _coll29.add(_526_v31);
            }
          }
          return _coll29;
        }());
        let _527_v32;
        _527_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(516),_493_v0);
        let _528_v33;
        _528_v33 = _dafny.Seq.of(_527_v32);
        let _529_v34;
        _529_v34 = _dafny.Seq.of(_528_v33);
        let _530_v35;
        _530_v35 = _dafny.Seq.UnicodeFromString("cldrgrfef");
        let _531_v36;
        let _nw100 = Array((new BigNumber(12)).toNumber());
        _nw100[(_dafny.ZERO).toNumber()] = _496_v3;
        _nw100[(_dafny.ONE).toNumber()] = ((_496_v3) ? (_493_v0) : ((_502_v8)[_module.__default.safeIndex(new BigNumber((_500_v7).cardinality()), new BigNumber((_502_v8).length))]));
        _nw100[(new BigNumber(2)).toNumber()] = _496_v3;
        _nw100[(new BigNumber(3)).toNumber()] = (_this).fm0(_dafny.MultiSet.FromArray((_529_v34)[_module.__default.safeIndex((_495_v2)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length))], new BigNumber((_529_v34).length))]), false, _530_v35, globalState);
        _nw100[(new BigNumber(4)).toNumber()] = ((_493_v0) ? (false) : (_493_v0));
        _nw100[(new BigNumber(5)).toNumber()] = (_493_v0) && (_496_v3);
        _nw100[(new BigNumber(6)).toNumber()] = _493_v0;
        _nw100[(new BigNumber(7)).toNumber()] = _493_v0;
        _nw100[(new BigNumber(8)).toNumber()] = ((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f5))).isLessThan(_498_v5);
        _nw100[(new BigNumber(9)).toNumber()] = _496_v3;
        _nw100[(new BigNumber(10)).toNumber()] = _496_v3;
        _nw100[(new BigNumber(11)).toNumber()] = ((_493_v0) ? (_493_v0) : (_496_v3));
        _531_v36 = _nw100;
        let _index82 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_531_v36).length));
        (_531_v36)[_index82] = _496_v3;
        let _index83 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_531_v36).length));
        (_531_v36)[_index83] = ((_500_v7).Intersect(_500_v7)).IsSubsetOf((_dafny.MultiSet.FromArray(_502_v8)).Union(_500_v7));
        let _532_v37;
        let _init13 = ((_533_v36, _534_v0) => function (_535_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe((_533_v36)[_module.__default.safeIndex(new BigNumber(486), new BigNumber((_533_v36).length))],_534_v0);
        })(_531_v36, _493_v0);
        let _nw101 = Array((new BigNumber(29)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw101.length); _i0_13++) {
          _nw101[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _532_v37 = _nw101;
        let _536_v38;
        _536_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f5),(_495_v2)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length))]);
        let _537_v39;
        _537_v39 = _dafny.Map.Empty.slice().updateUnsafe(_493_v0,true);
        let _index84 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_532_v37).length));
        (_532_v37)[_index84] = (_module.__default.fm10((_495_v2)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length))], _536_v38, !(_496_v3), (_495_v2)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length))], globalState)).Merge(_537_v39);
        let _index85 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_532_v37).length));
        (_532_v37)[_index85] = _537_v39;
      }
      let _index86 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length));
      (_495_v2)[_index86] = _this.f6;
      let _538_v40;
      _538_v40 = _dafny.Seq.of(new BigNumber(777), (_495_v2)[_module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length))]);
      let _539_v41;
      _539_v41 = _dafny.Map.Empty.slice().updateUnsafe(_493_v0,_496_v3);
      let _540_v42;
      _540_v42 = _dafny.Set.fromElements(_493_v0);
      let _index87 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_495_v2).length));
      (_495_v2)[_index87] = (_498_v5).minus(((_538_v40)[_module.__default.safeIndex(_this.f6, new BigNumber((_538_v40).length))]).plus((_this).fm2(new BigNumber((_539_v41).length), _540_v42, globalState)));
      r0 = (_538_v40)[_module.__default.safeIndex(_498_v5, new BigNumber((_538_v40).length))];
      r1 = _493_v0;
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let _541_v0;
      _541_v0 = false;
      if (_541_v0) {
        let _542_v1;
        _542_v1 = _dafny.Seq.of(p0, _this.f6);
        let _543_v2;
        _543_v2 = new _dafny.CodePoint('q'.codePointAt(0));
        let _544_v3;
        _544_v3 = _dafny.Map.Empty.slice().updateUnsafe(_543_v2,(_dafny.ZERO).minus((_this).f5));
        let _545_v4;
        _545_v4 = _dafny.Seq.UnicodeFromString("rwfr");
        let _rhs112 = _542_v1;
        let _rhs113 = _dafny.areEqual(_module.__default.fm11((_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), (_this).f5, _this.f6, _this.f6, globalState), _module.__default.fm11((_this).f5, new BigNumber(((_544_v3).update((_545_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_545_v4).length))], new BigNumber(996))).length), p0, new BigNumber((_545_v4).length), globalState));
        _542_v1 = _rhs112;
        _541_v0 = _rhs113;
        let _546_v5;
        _546_v5 = _module.D1.create_DC6(_545_v4, _this.f6);
        let _547_v6;
        _547_v6 = _module.D0.create_DC0(false);
        let _548_v7;
        _548_v7 = _dafny.MultiSet.fromElements((_547_v6).dtor_cf0, _541_v0);
        let _549_v8;
        _549_v8 = _dafny.MultiSet.fromElements(_548_v7);
        _541_v0 = (_this).fm8(_dafny.Seq.Concat(_545_v4, (_546_v5).dtor_cf5), !(_541_v0) || (!(false)), _541_v0, (_549_v8).IsDisjointFrom(_549_v8), globalState);
        let _550_v9;
        let _nw102 = Array((new BigNumber(3)).toNumber()).fill(_dafny.MultiSet.Empty);
        _550_v9 = _nw102;
        let _551_v10;
        _551_v10 = _dafny.Set.fromElements(_541_v0, _541_v0, false);
        let _552_v11;
        _552_v11 = _dafny.MultiSet.fromElements(_543_v2, _543_v2);
        let _553_v12;
        _553_v12 = _dafny.MultiSet.fromElements((_this).f5, (_this).fm2((_this).f5, _551_v10, globalState), (((_552_v11).contains(_543_v2)) ? ((_552_v11).get(_543_v2)) : ((((_544_v3).contains(_543_v2)) ? ((_544_v3).get(_543_v2)) : (new BigNumber(660))))), p0, (_this).f5);
        let _index88 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_550_v9).length));
        (_550_v9)[_index88] = _553_v12;
        let _index89 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_550_v9).length));
        (_550_v9)[_index89] = _dafny.MultiSet.fromElements(p0);
        let _554_v13;
        _554_v13 = _dafny.Set.fromElements(p0);
        let _555_v14;
        _555_v14 = _module.D1.create_DC4(_554_v13);
        let _source8 = _555_v14;
        if (_source8.is_DC5) {
          let _556_v15;
          let _nw103 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _556_v15 = _nw103;
          let _index90 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_556_v15).length));
          (_556_v15)[_index90] = (_this).fm2(_this.f6, _551_v10, globalState);
          let _index91 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_556_v15).length));
          let _rhs114 = _541_v0;
          let _rhs115 = (_552_v11).IsDisjointFrom(_552_v11);
          let _rhs116 = ((_541_v0) ? (_556_v15) : (_556_v15));
          let _rhs117 = (_dafny.ZERO).minus((_this).fm2(_this.f6, (_dafny.Set.fromElements(_541_v0)).Difference(_551_v10), globalState));
          let _lhs69 = _556_v15;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_556_v15).length));
          _541_v0 = _rhs114;
          _541_v0 = _rhs115;
          _556_v15 = _rhs116;
          _lhs69[_lhs70] = _rhs117;
          (_this).f6 = (_this).f5;
          let _557_v16;
          let _nw104 = Array((new BigNumber(3)).toNumber()).fill(false);
          _557_v16 = _nw104;
          let _index92 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_557_v16).length));
          (_557_v16)[_index92] = !(_541_v0);
          let _index93 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_557_v16).length));
          (_557_v16)[_index93] = _541_v0;
          let _index94 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_556_v15).length));
          (_556_v15)[_index94] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("npeqafbc"), _545_v4)).length);
        } else if (_source8.is_DC6) {
          let _558___mcc_h0 = (_source8).cf5;
          let _559___mcc_h1 = (_source8).cf6;
          let _560_cf6 = _559___mcc_h1;
          let _561_cf5 = _558___mcc_h0;
          let _562_v17;
          let _nw105 = Array((new BigNumber(3)).toNumber());
          _nw105[(_dafny.ZERO).toNumber()] = _560_cf6;
          _nw105[(_dafny.ONE).toNumber()] = _this.f6;
          _nw105[(new BigNumber(2)).toNumber()] = _this.f6;
          _562_v17 = _nw105;
          let _index95 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_562_v17).length));
          (_562_v17)[_index95] = new BigNumber(728);
          let _index96 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_562_v17).length));
          (_562_v17)[_index96] = (_this).f5;
          let _563_v18;
          let _nw106 = new _module.C0();
          _nw106.__ctor(p0);
          _563_v18 = _nw106;
          let _564_v19;
          _564_v19 = _dafny.MultiSet.fromElements(((_541_v0) ? (_563_v18) : (_563_v18)), _563_v18);
          _564_v19 = (_564_v19).Union(_564_v19);
          (globalState).f2 = (_this).f5;
        } else {
          let _565___mcc_h2 = (_source8).cf4;
          let _566_cf4 = _565___mcc_h2;
          let _567_v20;
          _567_v20 = _module.D0.create_DC2();
          let _568_v21;
          let _nw107 = Array((new BigNumber(15)).toNumber());
          _nw107[(_dafny.ZERO).toNumber()] = ((_541_v0) ? (_567_v20) : (_567_v20));
          _nw107[(_dafny.ONE).toNumber()] = _module.D0.create_DC2();
          _nw107[(new BigNumber(2)).toNumber()] = _567_v20;
          _nw107[(new BigNumber(3)).toNumber()] = _567_v20;
          _nw107[(new BigNumber(4)).toNumber()] = _567_v20;
          _nw107[(new BigNumber(5)).toNumber()] = _567_v20;
          _nw107[(new BigNumber(6)).toNumber()] = _567_v20;
          _nw107[(new BigNumber(7)).toNumber()] = _567_v20;
          _nw107[(new BigNumber(8)).toNumber()] = _567_v20;
          _nw107[(new BigNumber(9)).toNumber()] = _module.D0.create_DC2();
          _nw107[(new BigNumber(10)).toNumber()] = _567_v20;
          _nw107[(new BigNumber(11)).toNumber()] = _567_v20;
          _nw107[(new BigNumber(12)).toNumber()] = _module.D0.create_DC2();
          _nw107[(new BigNumber(13)).toNumber()] = _module.D0.create_DC2();
          _nw107[(new BigNumber(14)).toNumber()] = ((_541_v0) ? (_567_v20) : (_567_v20));
          _568_v21 = _nw107;
          let _index97 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_568_v21).length));
          (_568_v21)[_index97] = _module.D0.create_DC2();
          let _index98 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_568_v21).length));
          (_568_v21)[_index98] = _567_v20;
          let _569_v22;
          _569_v22 = _dafny.Map.Empty.slice().updateUnsafe(_545_v4,(new BigNumber(((_550_v9)[_module.__default.safeIndex(new BigNumber(379), new BigNumber((_550_v9).length))]).cardinality())).multipliedBy(p0));
          _569_v22 = (_569_v22).update(((_541_v0) ? (_545_v4) : (_545_v4)), (new BigNumber((_554_v13).length)).multipliedBy((_this).f5));
          let _570_v23;
          let _nw108 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _570_v23 = _nw108;
          let _571_v24;
          let _nw109 = new _module.C1();
          _nw109.__ctor((_dafny.ZERO).minus(p0), p0);
          _571_v24 = _nw109;
          let _572_v25;
          _572_v25 = _dafny.Set.fromElements(_571_v24, _571_v24, _571_v24, _571_v24, _571_v24);
          let _573_v26;
          let _nw110 = Array((new BigNumber(4)).toNumber()).fill([]);
          _573_v26 = _nw110;
          let _index99 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_573_v26).length));
          (_573_v26)[_index99] = _570_v23;
          let _574_v27;
          _574_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_541_v0);
          let _575_v28;
          _575_v28 = _module.D5.create_DC17(_this.f6, _570_v23, _574_v27);
          let _index100 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_573_v26).length));
          let _rhs118 = _570_v23;
          let _rhs119 = _572_v25;
          let _rhs120 = (_575_v28).dtor_cf28;
          let _rhs121 = ((p0).isLessThan((_this).fm2((_this).f5, _551_v10, globalState))) && (((_this).f5).isEqualTo(p0));
          let _rhs122 = ((_548_v7).Difference(_dafny.MultiSet.fromElements(!(true)))).IsProperSubsetOf(_548_v7);
          let _lhs71 = _573_v26;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_573_v26).length));
          _570_v23 = _rhs118;
          _572_v25 = _rhs119;
          _lhs71[_lhs72] = _rhs120;
          _541_v0 = _rhs121;
          _541_v0 = _rhs122;
          (_this).f6 = (_571_v24).f5;
        }
        _541_v0 = _dafny.Seq.IsProperPrefixOf(_545_v4, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gyqisghb"), _545_v4));
      } else {
        let _576_v30;
        _576_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,_541_v0);
        let _577_v31;
        let _nw111 = new _module.C1();
        _nw111.__ctor((_this).f5, (_this).f5);
        _577_v31 = _nw111;
        let _578_v32;
        _578_v32 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_577_v31);
        let _579_v33;
        _579_v33 = _dafny.Seq.UnicodeFromString("ynpxt");
        let _580_v34;
        _580_v34 = _dafny.Map.Empty.slice().updateUnsafe(_579_v33,_578_v32);
        let _581_v35;
        let _nw112 = Array((new BigNumber(28)).toNumber());
        _nw112[(_dafny.ZERO).toNumber()] = p0;
        _nw112[(_dafny.ONE).toNumber()] = p0;
        _nw112[(new BigNumber(2)).toNumber()] = _this.f6;
        _nw112[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_this).f5, _this.f6);
        _nw112[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).f5);
        _nw112[(new BigNumber(5)).toNumber()] = (_this).f5;
        _nw112[(new BigNumber(6)).toNumber()] = new BigNumber((function () {
          let _coll30 = new _dafny.Set();
          for (const _compr_30 of _dafny.IntegerRange(new BigNumber(590), new BigNumber(716))) {
            let _582_v29 = _compr_30;
            if (((new BigNumber(590)).isLessThanOrEqualTo(_582_v29)) && ((_582_v29).isLessThan(new BigNumber(716)))) {
              _coll30.add((_582_v29).multipliedBy(_this.f6));
            }
          }
          return _coll30;
        }()).length);
        _nw112[(new BigNumber(7)).toNumber()] = (new BigNumber((_576_v30).length)).minus(new BigNumber(775));
        _nw112[(new BigNumber(8)).toNumber()] = _this.f6;
        _nw112[(new BigNumber(9)).toNumber()] = new BigNumber(((_578_v32).Merge((((_580_v34).contains(_579_v33)) ? ((_580_v34).get(_579_v33)) : (_578_v32)))).length);
        _nw112[(new BigNumber(10)).toNumber()] = p0;
        _nw112[(new BigNumber(11)).toNumber()] = _this.f6;
        _nw112[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt((_this).f5, (_this).f5);
        _nw112[(new BigNumber(13)).toNumber()] = ((_this).f5).plus(new BigNumber(57));
        _nw112[(new BigNumber(14)).toNumber()] = p0;
        _nw112[(new BigNumber(15)).toNumber()] = (_this).f5;
        _nw112[(new BigNumber(16)).toNumber()] = p0;
        _nw112[(new BigNumber(17)).toNumber()] = ((_541_v0) ? (_this.f6) : ((_this).f5));
        _nw112[(new BigNumber(18)).toNumber()] = _this.f6;
        _nw112[(new BigNumber(19)).toNumber()] = (_this).f5;
        _nw112[(new BigNumber(20)).toNumber()] = p0;
        _nw112[(new BigNumber(21)).toNumber()] = (_this).f5;
        _nw112[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(((true) ? (_this.f6) : (p0)));
        _nw112[(new BigNumber(23)).toNumber()] = (_this.f6).plus(p0);
        _nw112[(new BigNumber(24)).toNumber()] = _this.f6;
        _nw112[(new BigNumber(25)).toNumber()] = _this.f6;
        _nw112[(new BigNumber(26)).toNumber()] = (p0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("fmkmdiv")).length));
        _nw112[(new BigNumber(27)).toNumber()] = _this.f6;
        _581_v35 = _nw112;
        _581_v35 = (_module.D5.create_DC17(_this.f6, _581_v35, _576_v30)).dtor_cf28;
        let _583_v36;
        _583_v36 = _module.D5.create_DC16(_581_v35);
        let _584_v37;
        _584_v37 = _dafny.Seq.of(_dafny.Set.fromElements(_583_v36, _583_v36, _583_v36, _583_v36, _583_v36));
        let _rhs123 = _541_v0;
        let _rhs124 = _541_v0;
        let _rhs125 = _584_v37;
        let _rhs126 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tomjydsx"), _dafny.Seq.Concat(_579_v33, _579_v33))).length);
        let _rhs127 = !(true) || (_541_v0);
        let _lhs73 = globalState;
        _541_v0 = _rhs123;
        _541_v0 = _rhs124;
        _584_v37 = _rhs125;
        _lhs73.f2 = _rhs126;
        _541_v0 = _rhs127;
        let _585_v38;
        _585_v38 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_this.f6);
        let _586_v39;
        _586_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_585_v38);
        _586_v39 = (_586_v39).update(p0, ((!(!(_541_v0))) ? (_dafny.Map.Empty.slice().updateUnsafe(_541_v0,_this.f6)) : (_dafny.Map.Empty.slice().updateUnsafe(_541_v0,_this.f6))));
        let _587_v40;
        _587_v40 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(47)), function (_588_i0) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        })).length));
        let _589_v41;
        let _nw113 = Array((new BigNumber(3)).toNumber());
        _nw113[(_dafny.ZERO).toNumber()] = _541_v0;
        _nw113[(_dafny.ONE).toNumber()] = !_dafny.areEqual(_587_v40, _dafny.Seq.update(_587_v40, _module.__default.safeIndex(_this.f6, new BigNumber((_587_v40).length)), new BigNumber((_587_v40).length)));
        _nw113[(new BigNumber(2)).toNumber()] = _541_v0;
        _589_v41 = _nw113;
        _589_v41 = _589_v41;
        let _index101 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_581_v35).length));
        (_581_v35)[_index101] = _this.f6;
        let _index102 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_581_v35).length));
        (_581_v35)[_index102] = p0;
      }
      let _590_v42;
      let _nw114 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _590_v42 = _nw114;
      let _591_v43;
      _591_v43 = new _dafny.CodePoint('a'.codePointAt(0));
      let _index103 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length));
      (_590_v42)[_index103] = _591_v43;
      let _592_v44;
      _592_v44 = _dafny.MultiSet.fromElements(_541_v0, _541_v0);
      let _593_v45;
      _593_v45 = _dafny.MultiSet.fromElements(_592_v44);
      let _index104 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length));
      (_590_v42)[_index104] = (((_593_v45).IsDisjointFrom(_593_v45)) ? (_591_v43) : (_591_v43));
      let _594_v46;
      _594_v46 = _dafny.Seq.of(_541_v0, _541_v0);
      if ((_594_v46)[_module.__default.safeIndex(new BigNumber(-612), new BigNumber((_594_v46).length))]) {
        let _595_v47;
        let _nw115 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _595_v47 = _nw115;
        let _596_v48;
        _596_v48 = _dafny.Map.Empty.slice().updateUnsafe(_592_v44,_541_v0);
        let _597_v49;
        _597_v49 = _dafny.Seq.of(p0, new BigNumber((_596_v48).length));
        let _598_v50;
        _598_v50 = _dafny.Seq.of((_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))], _591_v43);
        let _599_v51;
        _599_v51 = _dafny.Seq.of((_597_v49)[_module.__default.safeIndex(new BigNumber((_598_v50).length), new BigNumber((_597_v49).length))]);
        let _600_v52;
        _600_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_597_v49).length),_541_v0);
        let _601_v53;
        _601_v53 = _module.D5.create_DC17((_this).f5, _595_v47, (_600_v52).Merge(_600_v52));
        _601_v53 = _601_v53;
        (globalState).f2 = (_this).f5;
        if ((_541_v0) === ((_this.f6).isLessThanOrEqualTo(p0))) {
          let _602_v54;
          _602_v54 = _dafny.Set.fromElements(_592_v44);
          let _603_v55;
          _603_v55 = _dafny.MultiSet.fromElements(_602_v54);
          let _604_v57;
          _604_v57 = _dafny.Map.Empty.slice().updateUnsafe(_591_v43,_541_v0);
          (_this).f6 = ((p0).plus((_this).f5)).multipliedBy((((_603_v55).contains(_602_v54)) ? ((_603_v55).get(_602_v54)) : (new BigNumber((function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of (_604_v57).Keys.Elements) {
              let _605_v56 = _compr_31;
              if ((_604_v57).contains(_605_v56)) {
                _coll31.push([_605_v56,_541_v0]);
              }
            }
            return _coll31;
          }()).length))));
          let _606_v58;
          _606_v58 = _dafny.Map.Empty.slice().updateUnsafe(_595_v47,!(_541_v0) || (_541_v0));
          let _607_v59;
          _607_v59 = _dafny.Seq.of(_606_v58);
          _606_v58 = (_606_v58).Merge((_607_v59)[_module.__default.safeIndex(new BigNumber(-172), new BigNumber((_607_v59).length))]);
          let _index105 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length));
          (_590_v42)[_index105] = (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))];
          _598_v50 = _598_v50;
          let _608_v60;
          _608_v60 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,(_this).f5);
          _608_v60 = (_608_v60).update(p0, new BigNumber((_dafny.Seq.Concat(_598_v50, _598_v50)).length));
        } else {
          (globalState).f2 = (_this).f5;
          let _609_v61;
          _609_v61 = _dafny.MultiSet.fromElements(new BigNumber((_598_v50).length), _this.f6);
          let _610_v62;
          _610_v62 = _module.D0.create_DC2();
          let _611_v63;
          _611_v63 = _module.D3.create_DC9(_dafny.MultiSet.fromElements(p0), _610_v62, (_this).f5);
          let _612_v64;
          _612_v64 = _dafny.Set.fromElements(_this.f6);
          let _613_v65;
          _613_v65 = _dafny.MultiSet.fromElements((_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))], (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))]);
          let _614_v66;
          let _nw116 = Array((new BigNumber(20)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = _541_v0;
          _nw116[(_dafny.ONE).toNumber()] = _dafny.Seq.IsPrefixOf(_597_v49, _599_v51);
          _nw116[(new BigNumber(2)).toNumber()] = false;
          _nw116[(new BigNumber(3)).toNumber()] = !(false);
          _nw116[(new BigNumber(4)).toNumber()] = (_dafny.MultiSet.fromElements((_611_v63).dtor_cf11)).IsSubsetOf(_609_v61);
          _nw116[(new BigNumber(5)).toNumber()] = !(_612_v64).equals(_dafny.Set.fromElements((_this).f5));
          _nw116[(new BigNumber(6)).toNumber()] = _541_v0;
          _nw116[(new BigNumber(7)).toNumber()] = _541_v0;
          _nw116[(new BigNumber(8)).toNumber()] = (_592_v44).IsDisjointFrom(_592_v44);
          _nw116[(new BigNumber(9)).toNumber()] = _dafny.Seq.IsPrefixOf(_598_v50, _598_v50);
          _nw116[(new BigNumber(10)).toNumber()] = (_609_v61).IsDisjointFrom(_dafny.MultiSet.fromElements((_this).f5, p0));
          _nw116[(new BigNumber(11)).toNumber()] = _541_v0;
          _nw116[(new BigNumber(12)).toNumber()] = (new BigNumber(888)).isLessThan(_this.f6);
          _nw116[(new BigNumber(13)).toNumber()] = _541_v0;
          _nw116[(new BigNumber(14)).toNumber()] = _541_v0;
          _nw116[(new BigNumber(15)).toNumber()] = !(_541_v0);
          _nw116[(new BigNumber(16)).toNumber()] = ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), _591_v43)).update(_591_v43, _module.__default.abs(p0))).IsSubsetOf(_613_v65);
          _nw116[(new BigNumber(17)).toNumber()] = _541_v0;
          _nw116[(new BigNumber(18)).toNumber()] = _541_v0;
          _nw116[(new BigNumber(19)).toNumber()] = _541_v0;
          _614_v66 = _nw116;
          let _index106 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_614_v66).length));
          (_614_v66)[_index106] = _541_v0;
          let _index107 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_614_v66).length));
          (_614_v66)[_index107] = ((_this).f5).isLessThanOrEqualTo((_this).f5);
          let _index108 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_614_v66).length));
          (_614_v66)[_index108] = _541_v0;
          let _615_v67;
          _615_v67 = _dafny.MultiSet.fromElements(_595_v47);
          let _616_v68;
          _616_v68 = _dafny.Map.Empty.slice().updateUnsafe(_615_v67,_597_v49);
          _616_v68 = (_616_v68).update(_615_v67, _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_this).f5), _module.__default.safeIndex(new BigNumber((_599_v51).length), new BigNumber((_dafny.Seq.of((_this).f5)).length)), p0), _597_v49));
          let _617_v69;
          _617_v69 = _dafny.Set.fromElements(_614_v66, _614_v66);
          _617_v69 = _617_v69;
        }
        let _618_v70;
        let _619_v71;
        let _620_v72;
        let _621_v73;
        let _out44;
        let _out45;
        let _out46;
        let _out47;
        let _outcollector13 = _module.__default.m0(new _dafny.CodePoint('l'.codePointAt(0)), globalState);
        _out44 = _outcollector13[0];
        _out45 = _outcollector13[1];
        _out46 = _outcollector13[2];
        _out47 = _outcollector13[3];
        _618_v70 = _out44;
        _619_v71 = _out45;
        _620_v72 = _out46;
        _621_v73 = _out47;
        let _622_v74;
        _622_v74 = _dafny.Set.fromElements((_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))]);
        _622_v74 = _622_v74;
      } else {
        let _623_v75;
        _623_v75 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_541_v0);
        let _624_v76;
        _624_v76 = _dafny.Seq.UnicodeFromString("ycjiany");
        _541_v0 = (((_623_v75).contains((_this.f6).isLessThan(new BigNumber((_624_v76).length)))) ? ((_623_v75).get((_this.f6).isLessThan(new BigNumber((_624_v76).length)))) : ((new BigNumber((_dafny.Seq.UnicodeFromString("wml")).length)).isLessThan((_dafny.ZERO).minus(_this.f6))));
        let _625_v77;
        let _nw117 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
        _625_v77 = _nw117;
        let _626_v78;
        _626_v78 = _dafny.Set.fromElements((_this).f5, p0, (_this).f5, p0);
        let _627_v79;
        _627_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_626_v78).length),new BigNumber(952));
        let _628_v80;
        _628_v80 = _dafny.Seq.of(p0);
        let _629_v81;
        let _nw118 = new _module.C0();
        _nw118.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_541_v0,(((_627_v79).contains(p0)) ? ((_627_v79).get(p0)) : (new BigNumber((_628_v80).length))))).length));
        _629_v81 = _nw118;
        let _630_v82;
        _630_v82 = _dafny.Seq.of(_629_v81, _629_v81);
        let _index109 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_625_v77).length));
        (_625_v77)[_index109] = _630_v82;
        let _index110 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_625_v77).length));
        (_625_v77)[_index110] = _630_v82;
        let _631_v83;
        let _nw119 = new _module.C1();
        _nw119.__ctor(new BigNumber(579), new BigNumber(464));
        _631_v83 = _nw119;
        _541_v0 = !(((_541_v0) ? ((_594_v46)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_626_v78).length)), new BigNumber((_594_v46).length))]) : (_541_v0)));
        let _632_v84;
        let _init14 = function (_633_i1) {
          return _module.__default.safeModuloInt(_633_i1, (_this).f5);
        };
        let _nw120 = Array((new BigNumber(4)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw120.length); _i0_14++) {
          _nw120[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _632_v84 = _nw120;
        let _index111 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_632_v84).length));
        (_632_v84)[_index111] = (_dafny.ZERO).minus((_this.f6).multipliedBy(_this.f6));
        let _index112 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_632_v84).length));
        (_632_v84)[_index112] = (_dafny.ZERO).minus((p0).minus(_this.f6));
      }
      let _634_i2;
      _634_i2 = _dafny.ZERO;
      L3: {
        while (!(!(_541_v0)) || (_541_v0)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_634_i2)) {
              break L3;
            }
            _634_i2 = (_634_i2).plus(_dafny.ONE);
            (_this).m4((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(89)), function (_635_i3) {
              return (_this).f5;
            }))).IsProperSubsetOf(_dafny.MultiSet.fromElements((_this).f5)), globalState);
            (_this).f6 = new BigNumber(841);
            let _636_v86;
            _636_v86 = _dafny.Seq.of(function () {
              let _coll32 = new _dafny.Set();
              for (const _compr_32 of _dafny.IntegerRange(new BigNumber(379), new BigNumber(-220))) {
                let _637_v85 = _compr_32;
                if (((new BigNumber(379)).isLessThanOrEqualTo(_637_v85)) && ((_637_v85).isLessThan(new BigNumber(-220)))) {
                  _coll32.add(_module.__default.safeModuloInt(_637_v85, new BigNumber((_dafny.Seq.UnicodeFromString("wxvn")).length)));
                }
              }
              return _coll32;
            }());
            let _source9 = _module.__default.fm20((_636_v86)[_module.__default.safeIndex(_this.f6, new BigNumber((_636_v86).length))], !((_this.f6).isEqualTo((_this).f5)), false, globalState);
            if (_source9.is_DC1) {
              let _638___mcc_h3 = (_source9).cf1;
              let _639___mcc_h4 = (_source9).cf2;
              let _640_cf2 = _639___mcc_h4;
              let _641_cf1 = _638___mcc_h3;
              let _642_v87;
              _642_v87 = _dafny.Seq.UnicodeFromString("ekgniqmts");
              let _index113 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length));
              (_590_v42)[_index113] = (_642_v87)[_module.__default.safeIndex((_this.f6).plus(new BigNumber(862)), new BigNumber((_642_v87).length))];
              _640_cf2 = _641_cf1;
              let _643_v88;
              let _init15 = ((_644_p0) => function (_645_i4) {
                return (_dafny.MultiSet.fromElements(new BigNumber(729))).Union(_dafny.MultiSet.fromElements(_644_p0, (_this).f5, new BigNumber(180)));
              })(p0);
              let _nw121 = Array((new BigNumber(7)).toNumber());
              for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw121.length); _i0_15++) {
                _nw121[_i0_15] = _init15(new BigNumber(_i0_15));
              }
              _643_v88 = _nw121;
              _643_v88 = _643_v88;
              let _646_v89;
              _646_v89 = _dafny.Map.Empty.slice().updateUnsafe(_641_cf1,p0);
              let _647_v90;
              _647_v90 = _dafny.Set.fromElements(_641_cf1);
              let _648_v91;
              _648_v91 = _dafny.Seq.of(_647_v90, _647_v90, _647_v90, _647_v90);
              let _649_v92;
              _649_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(473),_this.f6);
              let _650_v93;
              _650_v93 = _dafny.MultiSet.fromElements(_this.f6);
              let _651_v94;
              _651_v94 = _dafny.Map.Empty.slice().updateUnsafe(_650_v93,_541_v0);
              let _652_v95;
              _652_v95 = _dafny.Set.fromElements((_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))]);
              let _653_v96;
              _653_v96 = _dafny.Seq.of(_this.f6, new BigNumber((_642_v87).length));
              let _654_v98;
              _654_v98 = _dafny.MultiSet.fromElements(_652_v95, function () {
                let _coll33 = new _dafny.Set();
                for (const _compr_33 of (_652_v95).Elements) {
                  let _655_v97 = _compr_33;
                  if ((_652_v95).contains(_655_v97)) {
                    _coll33.add(_655_v97);
                  }
                }
                return _coll33;
              }());
              let _656_v99;
              let _nw122 = Array((new BigNumber(28)).toNumber());
              _nw122[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_this).fm2(new BigNumber((_dafny.Set.fromElements(_641_cf1, _641_cf1)).length), _dafny.Set.fromElements(_541_v0), globalState));
              _nw122[(_dafny.ONE).toNumber()] = (_this).f5;
              _nw122[(new BigNumber(2)).toNumber()] = ((_640_cf2) ? ((_this).f5) : ((_this).f5));
              _nw122[(new BigNumber(3)).toNumber()] = new BigNumber((_646_v89).length);
              _nw122[(new BigNumber(4)).toNumber()] = _this.f6;
              _nw122[(new BigNumber(5)).toNumber()] = _this.f6;
              _nw122[(new BigNumber(6)).toNumber()] = (_this).f5;
              _nw122[(new BigNumber(7)).toNumber()] = p0;
              _nw122[(new BigNumber(8)).toNumber()] = p0;
              _nw122[(new BigNumber(9)).toNumber()] = (_this).f5;
              _nw122[(new BigNumber(10)).toNumber()] = (_this).f5;
              _nw122[(new BigNumber(11)).toNumber()] = new BigNumber(-309);
              _nw122[(new BigNumber(12)).toNumber()] = _this.f6;
              _nw122[(new BigNumber(13)).toNumber()] = (_this).fm2(_this.f6, _647_v90, globalState);
              _nw122[(new BigNumber(14)).toNumber()] = new BigNumber((_642_v87).length);
              _nw122[(new BigNumber(15)).toNumber()] = ((!(_640_cf2)) ? (new BigNumber((_648_v91).length)) : (_this.f6));
              _nw122[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(p0);
              _nw122[(new BigNumber(17)).toNumber()] = _this.f6;
              _nw122[(new BigNumber(18)).toNumber()] = (new BigNumber((_649_v92).length)).multipliedBy(new BigNumber((_651_v94).length));
              _nw122[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_594_v46, _594_v46)).length);
              _nw122[(new BigNumber(20)).toNumber()] = (_this).f5;
              _nw122[(new BigNumber(21)).toNumber()] = (p0).multipliedBy(_this.f6);
              _nw122[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_this).fm2(new BigNumber((_652_v95).length), _647_v90, globalState));
              _nw122[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus((_this).f5);
              _nw122[(new BigNumber(24)).toNumber()] = (_653_v96)[_module.__default.safeIndex((_this).f5, new BigNumber((_653_v96).length))];
              _nw122[(new BigNumber(25)).toNumber()] = (_this).f5;
              _nw122[(new BigNumber(26)).toNumber()] = _this.f6;
              _nw122[(new BigNumber(27)).toNumber()] = _module.__default.safeDivisionInt((_this).f5, (_dafny.ZERO).minus((((_654_v98).contains(_652_v95)) ? ((_654_v98).get(_652_v95)) : ((_this).f5))));
              _656_v99 = _nw122;
              let _657_v100;
              _657_v100 = _module.D4.create_DC13(!(_541_v0), new BigNumber((_649_v92).length), _640_cf2);
              let _index114 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length));
              let _rhs128 = _656_v99;
              let _rhs129 = (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))];
              let _rhs130 = _656_v99;
              let _rhs131 = new BigNumber((_dafny.Seq.Concat(_653_v96, _653_v96)).length);
              let _rhs132 = (_657_v100).dtor_cf17;
              let _lhs74 = _590_v42;
              let _lhs75 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length));
              let _lhs76 = globalState;
              _656_v99 = _rhs128;
              _lhs74[_lhs75] = _rhs129;
              _656_v99 = _rhs130;
              _lhs76.f2 = _rhs131;
              _541_v0 = _rhs132;
            } else if (_source9.is_DC2) {
              let _658_v101;
              _658_v101 = _dafny.Set.fromElements(_541_v0);
              _541_v0 = ((_658_v101).IsSubsetOf(_658_v101)) === (_541_v0);
              let _659_v102;
              let _nw123 = Array((new BigNumber(16)).toNumber()).fill([]);
              _659_v102 = _nw123;
              let _660_v103;
              let _nw124 = Array((new BigNumber(20)).toNumber()).fill(false);
              _660_v103 = _nw124;
              let _661_v104;
              _661_v104 = _dafny.Seq.of(_660_v103);
              let _index115 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_659_v102).length));
              (_659_v102)[_index115] = (_661_v104)[_module.__default.safeIndex(new BigNumber(390), new BigNumber((_661_v104).length))];
              let _index116 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_659_v102).length));
              (_659_v102)[_index116] = _660_v103;
              let _662_v105;
              _662_v105 = _dafny.MultiSet.fromElements(_658_v101);
              let _663_v106;
              let _nw125 = new _module.C1();
              _nw125.__ctor(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(107)), ((_664_v43) => function (_665_i5) {
                return _664_v43;
              })(_591_v43))).length), new BigNumber(((_662_v105).update(_module.__default.fm21(p0, p0, _541_v0, globalState), _module.__default.abs(_this.f6))).cardinality()));
              _663_v106 = _nw125;
              (_this).f6 = (_dafny.ZERO).minus(p0);
            } else if (_source9.is_DC3) {
              let _666___mcc_h5 = (_source9).cf3;
              let _667_cf3 = _666___mcc_h5;
              let _668_v107;
              _668_v107 = _dafny.Seq.UnicodeFromString("lsid");
              let _669_v108;
              _669_v108 = _dafny.Seq.of(_668_v107);
              let _670_v109;
              _670_v109 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,_669_v108);
              _670_v109 = (_670_v109).update(p0, _669_v108);
              _668_v107 = _dafny.Seq.update(((_667_cf3) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-803)), function (_671_i6) {
                return new _dafny.CodePoint('r'.codePointAt(0));
              })) : (_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_668_v107, _dafny.Seq.UnicodeFromString("geuucj")), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_668_v107, _dafny.Seq.UnicodeFromString("geuucj"))).length)), (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))]), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_668_v107, _dafny.Seq.UnicodeFromString("geuucj")), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_668_v107, _dafny.Seq.UnicodeFromString("geuucj"))).length)), (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))])).length)), _591_v43))), _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((((_667_cf3) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-803)), function (_672_i6) {
                return new _dafny.CodePoint('r'.codePointAt(0));
              })) : (_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_668_v107, _dafny.Seq.UnicodeFromString("geuucj")), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_668_v107, _dafny.Seq.UnicodeFromString("geuucj"))).length)), (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))]), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_668_v107, _dafny.Seq.UnicodeFromString("geuucj")), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_668_v107, _dafny.Seq.UnicodeFromString("geuucj"))).length)), (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))])).length)), _591_v43)))).length)), _591_v43);
              let _673_v110;
              _673_v110 = _dafny.Seq.of(p0, _this.f6);
              let _674_v111;
              let _nw126 = new _module.C0();
              _nw126.__ctor(new BigNumber((_dafny.Set.fromElements(_673_v110)).length));
              _674_v111 = _nw126;
              let _675_v112;
              let _nw127 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
              _675_v112 = _nw127;
              let _676_v113;
              _676_v113 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,(_594_v46)[_module.__default.safeIndex(_this.f6, new BigNumber((_594_v46).length))]);
              let _677_v114;
              _677_v114 = _module.D5.create_DC17(new BigNumber((_dafny.Seq.of(_this.f6, new BigNumber((_dafny.Seq.of(_667_cf3)).length), new BigNumber(288))).length), _675_v112, _676_v113);
              let _678_v115;
              _678_v115 = _dafny.MultiSet.fromElements(_module.D5.create_DC17(new BigNumber((_594_v46).length), _675_v112, _676_v113), _677_v114);
              let _679_v116;
              let _nw128 = Array((new BigNumber(28)).toNumber()).fill(false);
              _679_v116 = _nw128;
              let _index117 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_679_v116).length));
              (_679_v116)[_index117] = _667_cf3;
              let _index118 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_679_v116).length));
              (_679_v116)[_index118] = _667_cf3;
              let _index119 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_679_v116).length));
              let _index120 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_679_v116).length));
              let _rhs133 = _678_v115;
              let _rhs134 = _541_v0;
              let _rhs135 = !(_541_v0);
              let _lhs77 = _679_v116;
              let _lhs78 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_679_v116).length));
              let _lhs79 = _679_v116;
              let _lhs80 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_679_v116).length));
              _678_v115 = _rhs133;
              _lhs77[_lhs78] = _rhs134;
              _lhs79[_lhs80] = _rhs135;
            } else {
              let _680___mcc_h6 = (_source9).cf0;
              let _681_cf0 = _680___mcc_h6;
              let _682_v117;
              _682_v117 = _dafny.Set.fromElements(p0);
              _682_v117 = ((_682_v117).Intersect(_682_v117)).Intersect(_682_v117);
              (globalState).f2 = p0;
              let _683_v118;
              let _nw129 = new _module.C1();
              _nw129.__ctor((_this).f5, (_this).f5);
              _683_v118 = _nw129;
              let _684_v119;
              let _nw130 = new _module.C1();
              _nw130.__ctor(p0, (_dafny.ZERO).minus(_this.f6));
              _684_v119 = _nw130;
              let _685_v120;
              _685_v120 = _684_v119;
              _684_v119 = (_685_v120);
            }
            if (_541_v0) {
              _541_v0 = _541_v0;
              let _686_v121;
              _686_v121 = _dafny.Seq.of(_this.f6);
              let _687_v123;
              _687_v123 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm22(function () {
                let _coll34 = new _dafny.Set();
                for (const _compr_34 of (_686_v121).Elements) {
                  let _688_v122 = _compr_34;
                  if (_dafny.Seq.contains(_686_v121, _688_v122)) {
                    _coll34.add((_688_v122).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(562)), function (_689_i7) {
                      return (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(561))).length));
                    })).length),new BigNumber((_dafny.Seq.UnicodeFromString("vudi")).length))).length)));
                  }
                }
                return _coll34;
              }(), _541_v0, globalState),_541_v0);
              let _690_v124;
              _690_v124 = _module.D7.create_DC21(_687_v123);
              let _691_v125;
              _691_v125 = _dafny.Set.fromElements(_541_v0);
              let _692_v126;
              _692_v126 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,(_this).fm1(globalState));
              let _693_v127;
              _693_v127 = _module.D3.create_DC8(_692_v126);
              let _694_v129;
              _694_v129 = _dafny.Seq.of(_693_v127);
              let _695_v131;
              _695_v131 = _dafny.Seq.of(_687_v123);
              let _696_v132;
              let _nw131 = Array((new BigNumber(28)).toNumber());
              _nw131[(_dafny.ZERO).toNumber()] = (_687_v123).Merge(_687_v123);
              _nw131[(_dafny.ONE).toNumber()] = (_690_v124).dtor_cf37;
              _nw131[(new BigNumber(2)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(3)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(4)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(5)).toNumber()] = (_687_v123).Merge((_module.__default.fm23((_this).fm2(p0, _691_v125, globalState), globalState)).update(_693_v127, true));
              _nw131[(new BigNumber(6)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(7)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_693_v127,_541_v0)).Merge(_687_v123);
              _nw131[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_693_v127,_541_v0);
              _nw131[(new BigNumber(9)).toNumber()] = (_687_v123).Merge(_687_v123);
              _nw131[(new BigNumber(10)).toNumber()] = ((function () {
                let _coll35 = new _dafny.Map();
                for (const _compr_35 of (_694_v129).Elements) {
                  let _697_v128 = _compr_35;
                  if (_dafny.Seq.contains(_694_v129, _697_v128)) {
                    _coll35.push([_697_v128,_541_v0]);
                  }
                }
                return _coll35;
              }()).update(_693_v127, (_this).fm1(globalState))).Merge(_687_v123);
              _nw131[(new BigNumber(11)).toNumber()] = function () {
                let _coll36 = new _dafny.Map();
                for (const _compr_36 of (_687_v123).Keys.Elements) {
                  let _698_v130 = _compr_36;
                  if ((_687_v123).contains(_698_v130)) {
                    _coll36.push([_698_v130,_541_v0]);
                  }
                }
                return _coll36;
              }();
              _nw131[(new BigNumber(12)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(13)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(14)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(15)).toNumber()] = ((_541_v0) ? (_687_v123) : (_687_v123));
              _nw131[(new BigNumber(16)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(17)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(18)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(19)).toNumber()] = (_695_v131)[_module.__default.safeIndex(p0, new BigNumber((_695_v131).length))];
              _nw131[(new BigNumber(20)).toNumber()] = (_687_v123).Merge(_687_v123);
              _nw131[(new BigNumber(21)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(22)).toNumber()] = (_687_v123).Merge(_dafny.Map.Empty.slice().updateUnsafe(_693_v127,_541_v0));
              _nw131[(new BigNumber(23)).toNumber()] = (_687_v123).Merge(_dafny.Map.Empty.slice().updateUnsafe(_693_v127,(_594_v46)[_module.__default.safeIndex(_this.f6, new BigNumber((_594_v46).length))]));
              _nw131[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_693_v127,!(_541_v0));
              _nw131[(new BigNumber(25)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(26)).toNumber()] = _687_v123;
              _nw131[(new BigNumber(27)).toNumber()] = _687_v123;
              _696_v132 = _nw131;
              let _index121 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_696_v132).length));
              (_696_v132)[_index121] = (_module.__default.fm23(_this.f6, globalState)).Merge(_687_v123);
              let _699_v133;
              _699_v133 = _dafny.Set.fromElements(_dafny.Seq.of((_this).f5, (_dafny.ZERO).minus(p0), _this.f6));
              let _index122 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_696_v132).length));
              let _rhs136 = ((_699_v133).Union(_699_v133)).IsProperSubsetOf((_699_v133).Union(_699_v133));
              let _rhs137 = _541_v0;
              let _rhs138 = (p0).isLessThanOrEqualTo((_dafny.ZERO).minus(p0));
              let _rhs139 = (new BigNumber((_594_v46).length)).multipliedBy(((_this).f5).multipliedBy((_this).f5));
              let _rhs140 = (_687_v123).Merge((_module.__default.fm23((_dafny.ZERO).minus(new BigNumber((_594_v46).length)), globalState)).Merge((_687_v123).update(_693_v127, _541_v0)));
              let _lhs81 = globalState;
              let _lhs82 = _696_v132;
              let _lhs83 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_696_v132).length));
              _541_v0 = _rhs136;
              _541_v0 = _rhs137;
              _541_v0 = _rhs138;
              _lhs81.f2 = _rhs139;
              _lhs82[_lhs83] = _rhs140;
              let _700_v134;
              let _nw132 = Array((new BigNumber(10)).toNumber());
              _nw132[(_dafny.ZERO).toNumber()] = _541_v0;
              _nw132[(_dafny.ONE).toNumber()] = _541_v0;
              _nw132[(new BigNumber(2)).toNumber()] = _541_v0;
              _nw132[(new BigNumber(3)).toNumber()] = _541_v0;
              _nw132[(new BigNumber(4)).toNumber()] = _541_v0;
              _nw132[(new BigNumber(5)).toNumber()] = !(p0).isEqualTo(_this.f6);
              _nw132[(new BigNumber(6)).toNumber()] = _541_v0;
              _nw132[(new BigNumber(7)).toNumber()] = _541_v0;
              _nw132[(new BigNumber(8)).toNumber()] = _541_v0;
              _nw132[(new BigNumber(9)).toNumber()] = _541_v0;
              _700_v134 = _nw132;
              _700_v134 = _700_v134;
              _541_v0 = !(_541_v0);
              let _701_v135;
              let _nw133 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
              _701_v135 = _nw133;
              let _index123 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_701_v135).length));
              (_701_v135)[_index123] = ((_541_v0) ? (_686_v121) : (_686_v121));
              let _index124 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_701_v135).length));
              (_701_v135)[_index124] = _686_v121;
            } else {
              let _702_v136;
              let _703_v137;
              let _704_v138;
              let _705_v139;
              let _out48;
              let _out49;
              let _out50;
              let _out51;
              let _outcollector14 = _module.__default.m0(_module.__default.fm18(_541_v0, _this.f6, _this.f6, true, globalState), globalState);
              _out48 = _outcollector14[0];
              _out49 = _outcollector14[1];
              _out50 = _outcollector14[2];
              _out51 = _outcollector14[3];
              _702_v136 = _out48;
              _703_v137 = _out49;
              _704_v138 = _out50;
              _705_v139 = _out51;
              let _706_v140;
              _706_v140 = _dafny.Set.fromElements(!(_541_v0));
              _705_v139 = (_this).fm2((_this).fm2(_this.f6, _706_v140, globalState), _706_v140, globalState);
              let _707_v141;
              _707_v141 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_702_v136);
              let _708_v142;
              _708_v142 = _dafny.Set.fromElements(new BigNumber(616), _705_v139);
              let _709_v143;
              let _nw134 = new _module.C1();
              _nw134.__ctor(_module.__default.safeDivisionInt(_this.f6, new BigNumber((_707_v141).length)), new BigNumber((_708_v142).length));
              _709_v143 = _nw134;
              let _rhs141 = _591_v43;
              let _rhs142 = _709_v143;
              let _rhs143 = _703_v137;
              let _rhs144 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_705_v139), _705_v139);
              let _lhs84 = globalState;
              _591_v43 = _rhs141;
              _709_v143 = _rhs142;
              _541_v0 = _rhs143;
              _lhs84.f2 = _rhs144;
              let _710_v144;
              _710_v144 = _dafny.Seq.UnicodeFromString("bnmpaxw");
              let _711_v145;
              let _nw135 = new _module.C1();
              _nw135.__ctor((_dafny.ZERO).minus(_705_v139), (_dafny.ZERO).minus(((_this).f5).minus(new BigNumber((_710_v144).length))));
              _711_v145 = _nw135;
              let _712_v146;
              _712_v146 = _module.D0.create_DC3(_541_v0);
              let _rhs145 = _this.f6;
              let _rhs146 = _712_v146;
              let _rhs147 = _703_v137;
              let _rhs148 = (_709_v143).fm2((_this.f6).minus(new BigNumber(922)), _706_v140, globalState);
              let _rhs149 = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Concat(_710_v144, _dafny.Seq.update(_710_v144, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_703_v137)).cardinality()), new BigNumber((_710_v144).length)), (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))]))).length)).multipliedBy(p0));
              let _lhs85 = globalState;
              let _lhs86 = globalState;
              _lhs85.f2 = _rhs145;
              _712_v146 = _rhs146;
              _703_v137 = _rhs147;
              _705_v139 = _rhs148;
              _lhs86.f2 = _rhs149;
            }
          }
        }
      }
      let _713_v147;
      let _init16 = ((_714_v0) => function (_715_i8) {
        return (_715_i8).multipliedBy(new BigNumber((_dafny.Set.fromElements(_714_v0, _714_v0, !(_714_v0))).length));
      })(_541_v0);
      let _nw136 = Array((new BigNumber(2)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw136.length); _i0_16++) {
        _nw136[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _713_v147 = _nw136;
      let _716_v148;
      _716_v148 = _dafny.Map.Empty.slice().updateUnsafe(_713_v147,_this.f6);
      let _717_v149;
      _717_v149 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_716_v148);
      _717_v149 = (_717_v149).update((((_592_v44).contains(_541_v0)) ? ((_592_v44).get(_541_v0)) : ((_this).f5)), _716_v148);
      if (true) {
        _591_v43 = _591_v43;
        let _718_v150;
        _718_v150 = _dafny.Set.fromElements(_713_v147, _713_v147);
        _718_v150 = _718_v150;
        let _719_v151;
        _719_v151 = _dafny.Seq.UnicodeFromString("i");
        let _index125 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length));
        (_713_v147)[_index125] = new BigNumber((_719_v151).length);
        let _index126 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length));
        (_713_v147)[_index126] = _this.f6;
        let _720_v152;
        _720_v152 = _module.D0.create_DC2();
        let _source10 = _720_v152;
        if (_source10.is_DC1) {
          let _721___mcc_h7 = (_source10).cf1;
          let _722___mcc_h8 = (_source10).cf2;
          let _723_cf2 = _722___mcc_h8;
          let _724_cf1 = _721___mcc_h7;
          let _725_v153;
          let _nw137 = Array((new BigNumber(3)).toNumber()).fill(false);
          _725_v153 = _nw137;
          let _726_v154;
          _726_v154 = _dafny.Map.Empty.slice().updateUnsafe(_725_v153,(_this).f5);
          _726_v154 = (_726_v154).update(_725_v153, _this.f6);
          (_this).f6 = _module.__default.safeDivisionInt(new BigNumber(234), (_this).f5);
          let _index127 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length));
          (_713_v147)[_index127] = new BigNumber(236);
          let _index128 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_725_v153).length));
          (_725_v153)[_index128] = _541_v0;
          let _index129 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_725_v153).length));
          (_725_v153)[_index129] = _724_cf1;
        } else if (_source10.is_DC2) {
          _541_v0 = (_541_v0) || (_541_v0);
          let _727_v155;
          _727_v155 = _dafny.Map.Empty.slice().updateUnsafe(true,_541_v0);
          let _index130 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length));
          (_713_v147)[_index130] = (new BigNumber(((_727_v155).Merge(_module.__default.fm10((_this).f5, _dafny.Map.Empty.slice().updateUnsafe(_this.f6,p0), _541_v0, p0, globalState))).length)).plus((p0).minus(p0));
          let _728_v156;
          let _nw138 = Array((new BigNumber(6)).toNumber()).fill(_module.D7.Default());
          _728_v156 = _nw138;
          let _729_v157;
          _729_v157 = _module.D3.create_DC8(_727_v155);
          let _730_v158;
          _730_v158 = _dafny.Map.Empty.slice().updateUnsafe(_729_v157,_541_v0);
          let _731_v159;
          _731_v159 = _dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC21(_730_v158),_728_v156);
          let _732_v160;
          _732_v160 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_541_v0,_541_v0));
          let _index131 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length));
          let _rhs150 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))], (_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))]), (_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))]);
          let _rhs151 = (((_731_v159).contains(_module.__default.fm24(_732_v160, new BigNumber(-597), _541_v0, globalState))) ? ((_731_v159).get(_module.__default.fm24(_732_v160, new BigNumber(-597), _541_v0, globalState))) : (_728_v156));
          let _lhs87 = _713_v147;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length));
          _lhs87[_lhs88] = _rhs150;
          _728_v156 = _rhs151;
          let _733_v161;
          let _nw139 = Array((new BigNumber(16)).toNumber()).fill([]);
          _733_v161 = _nw139;
          let _734_v162;
          let _nw140 = Array((new BigNumber(17)).toNumber());
          _nw140[(_dafny.ZERO).toNumber()] = _541_v0;
          _nw140[(_dafny.ONE).toNumber()] = _541_v0;
          _nw140[(new BigNumber(2)).toNumber()] = _541_v0;
          _nw140[(new BigNumber(3)).toNumber()] = _541_v0;
          _nw140[(new BigNumber(4)).toNumber()] = _541_v0;
          _nw140[(new BigNumber(5)).toNumber()] = false;
          _nw140[(new BigNumber(6)).toNumber()] = !(false);
          _nw140[(new BigNumber(7)).toNumber()] = _541_v0;
          _nw140[(new BigNumber(8)).toNumber()] = _541_v0;
          _nw140[(new BigNumber(9)).toNumber()] = true;
          _nw140[(new BigNumber(10)).toNumber()] = _541_v0;
          _nw140[(new BigNumber(11)).toNumber()] = _541_v0;
          _nw140[(new BigNumber(12)).toNumber()] = !(_541_v0);
          _nw140[(new BigNumber(13)).toNumber()] = _541_v0;
          _nw140[(new BigNumber(14)).toNumber()] = false;
          _nw140[(new BigNumber(15)).toNumber()] = false;
          _nw140[(new BigNumber(16)).toNumber()] = false;
          _734_v162 = _nw140;
          let _index132 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_733_v161).length));
          (_733_v161)[_index132] = _734_v162;
          let _index133 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_733_v161).length));
          let _rhs152 = ((_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))]).isLessThan((_dafny.ZERO).minus((_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))]));
          let _rhs153 = _734_v162;
          let _rhs154 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(362)), ((_735_v43) => function (_736_i9) {
            return _735_v43;
          })(_591_v43));
          let _lhs89 = _733_v161;
          let _lhs90 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_733_v161).length));
          _541_v0 = _rhs152;
          _lhs89[_lhs90] = _rhs153;
          _719_v151 = _rhs154;
        } else if (_source10.is_DC3) {
          let _737___mcc_h9 = (_source10).cf3;
          let _738_cf3 = _737___mcc_h9;
          let _739_v163;
          _739_v163 = _dafny.Set.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))], (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))], (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))], (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))]);
          let _740_v164;
          _740_v164 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_719_v151, _module.__default.safeIndex(new BigNumber(210), new BigNumber((_719_v151).length)), (_590_v42)[_module.__default.safeIndex(new BigNumber(508), new BigNumber((_590_v42).length))])).length));
          let _rhs155 = _739_v163;
          let _rhs156 = _module.__default.safeDivisionInt((new BigNumber((_740_v164).cardinality())).minus(_this.f6), (_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))]);
          let _lhs91 = globalState;
          _739_v163 = _rhs155;
          _lhs91.f2 = _rhs156;
          let _741_v165;
          let _init17 = function (_742_i10) {
            return (_742_i10).plus(new BigNumber((_dafny.Seq.UnicodeFromString("jsc")).length));
          };
          let _nw141 = Array((new BigNumber(28)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw141.length); _i0_17++) {
            _nw141[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _741_v165 = _nw141;
          let _743_v166;
          _743_v166 = _dafny.Seq.of(_741_v165, _741_v165, _741_v165);
          _741_v165 = (_743_v166)[_module.__default.safeIndex(_this.f6, new BigNumber((_743_v166).length))];
          let _744_v167;
          _744_v167 = _module.D0.create_DC1(_738_cf3, _541_v0);
          let _745_v168;
          _745_v168 = _dafny.Seq.of(_740_v164);
          let _746_v169;
          _746_v169 = _dafny.Set.fromElements(!(_541_v0));
          let _747_v170;
          _747_v170 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,(_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))]);
          let _748_v171;
          _748_v171 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2(p0, _746_v169, globalState),(_747_v170).update(_738_cf3, p0));
          let _749_v172;
          let _nw142 = Array((new BigNumber(8)).toNumber());
          _nw142[(_dafny.ZERO).toNumber()] = _module.D0.create_DC1(false, _541_v0);
          _nw142[(_dafny.ONE).toNumber()] = _module.D0.create_DC1(_541_v0, _541_v0);
          _nw142[(new BigNumber(2)).toNumber()] = _module.D0.create_DC1((_this).fm8(_dafny.Seq.UnicodeFromString("yotfvv"), _738_cf3, _541_v0, _541_v0, globalState), _541_v0);
          _nw142[(new BigNumber(3)).toNumber()] = _744_v167;
          _nw142[(new BigNumber(4)).toNumber()] = _module.D0.create_DC1(_738_cf3, _738_cf3);
          _nw142[(new BigNumber(5)).toNumber()] = _module.__default.fm25(_738_cf3, (_this).f5, _745_v168, _748_v171, globalState);
          _nw142[(new BigNumber(6)).toNumber()] = _744_v167;
          _nw142[(new BigNumber(7)).toNumber()] = _744_v167;
          _749_v172 = _nw142;
          let _index134 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_749_v172).length));
          (_749_v172)[_index134] = _module.__default.fm25(_738_cf3, _this.f6, _745_v168, _748_v171, globalState);
          let _index135 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_749_v172).length));
          (_749_v172)[_index135] = _module.__default.fm25(((_541_v0) ? (_738_cf3) : (true)), _this.f6, _dafny.Seq.Concat(_745_v168, _745_v168), _748_v171, globalState);
          let _index136 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length));
          (_713_v147)[_index136] = (_this).f5;
        } else {
          let _750___mcc_h10 = (_source10).cf0;
          let _751_cf0 = _750___mcc_h10;
          let _752_v173;
          _752_v173 = _module.D0.create_DC3(_541_v0);
          let _753_v174;
          _753_v174 = _dafny.MultiSet.fromElements(_591_v43, new _dafny.CodePoint('s'.codePointAt(0)), _591_v43);
          let _754_v175;
          _754_v175 = _dafny.Map.Empty.slice().updateUnsafe((_752_v173).dtor_cf3,new BigNumber((_753_v174).cardinality()));
          let _755_v176;
          _755_v176 = _dafny.Seq.of(_754_v175);
          let _756_v177;
          _756_v177 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_755_v176).length),_541_v0);
          let _757_v178;
          _757_v178 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_756_v177).length),_751_cf0);
          let _758_v179;
          let _nw143 = new _module.C0();
          _nw143.__ctor(new BigNumber((_756_v177).length));
          _758_v179 = _nw143;
          let _759_v182;
          _759_v182 = _module.D1.create_DC5();
          let _760_v183;
          _760_v183 = _dafny.Map.Empty.slice().updateUnsafe(_754_v175,(_this).fm9(_759_v182, _541_v0, globalState));
          let _761_v184;
          _761_v184 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_594_v46).length),_754_v175);
          let _762_v185;
          _762_v185 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_751_cf0,new BigNumber(352)),p0);
          let _763_v186;
          _763_v186 = _762_v185;
          let _764_v190;
          _764_v190 = _dafny.Map.Empty.slice().updateUnsafe(_754_v175,_719_v151);
          let _765_v191;
          let _nw144 = Array((new BigNumber(14)).toNumber());
          _nw144[(_dafny.ZERO).toNumber()] = function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of (function () {
              let _coll38 = new _dafny.Map();
              for (const _compr_38 of (_760_v183).Keys.Elements) {
                let _766_v181 = _compr_38;
                if ((_760_v183).contains(_766_v181)) {
                  _coll38.push([_766_v181,new BigNumber((_592_v44).cardinality())]);
                }
              }
              return _coll38;
            }()).Keys.Elements) {
              let _767_v180 = _compr_37;
              if ((function () {
                let _coll39 = new _dafny.Map();
                for (const _compr_39 of (_760_v183).Keys.Elements) {
                  let _766_v181 = _compr_39;
                  if ((_760_v183).contains(_766_v181)) {
                    _coll39.push([_766_v181,new BigNumber((_592_v44).cardinality())]);
                  }
                }
                return _coll39;
              }()).contains(_767_v180)) {
                _coll37.push([_767_v180,(_758_v179).f8]);
              }
            }
            return _coll37;
          }();
          _nw144[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_761_v184).contains((_758_v179).f8)) ? ((_761_v184).get((_758_v179).f8)) : (_754_v175)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f6,_541_v0)).length));
          _nw144[(new BigNumber(2)).toNumber()] = _762_v185;
          _nw144[(new BigNumber(3)).toNumber()] = (_763_v186);
          _nw144[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_754_v175,p0);
          _nw144[(new BigNumber(5)).toNumber()] = (function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of (_760_v183).Keys.Elements) {
              let _768_v187 = _compr_40;
              if ((_760_v183).contains(_768_v187)) {
                _coll40.push([_768_v187,new BigNumber((_dafny.Seq.of(function () {
                  let _coll41 = new _dafny.Map();
                  for (const _compr_41 of (_dafny.Seq.of(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(_541_v0,_751_cf0)))).Elements) {
                    let _769_v188 = _compr_41;
                    if (_dafny.Seq.contains(_dafny.Seq.of(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(_541_v0,_751_cf0))), _769_v188)) {
                      _coll41.push([_769_v188,(_758_v179).f8]);
                    }
                  }
                  return _coll41;
                }())).length)]);
              }
            }
            return _coll40;
          }()).Merge(_762_v185);
          _nw144[(new BigNumber(6)).toNumber()] = _762_v185;
          _nw144[(new BigNumber(7)).toNumber()] = _762_v185;
          _nw144[(new BigNumber(8)).toNumber()] = _762_v185;
          _nw144[(new BigNumber(9)).toNumber()] = (_762_v185).Merge(_dafny.Map.Empty.slice().updateUnsafe(_754_v175,(_758_v179).f8));
          _nw144[(new BigNumber(10)).toNumber()] = function () {
            let _coll42 = new _dafny.Map();
            for (const _compr_42 of (_764_v190).Keys.Elements) {
              let _770_v189 = _compr_42;
              if ((_764_v190).contains(_770_v189)) {
                _coll42.push([_770_v189,(_this).f5]);
              }
            }
            return _coll42;
          }();
          _nw144[(new BigNumber(11)).toNumber()] = (_762_v185).update((_754_v175).update(_541_v0, p0), (_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))]);
          _nw144[(new BigNumber(12)).toNumber()] = _762_v185;
          _nw144[(new BigNumber(13)).toNumber()] = _762_v185;
          _765_v191 = _nw144;
          let _index137 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_765_v191).length));
          (_765_v191)[_index137] = _762_v185;
          let _771_v192;
          _771_v192 = _module.D4.create_DC15(_module.__default.fm27(new BigNumber(437), globalState));
          let _index138 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_765_v191).length));
          let _rhs157 = _758_v179;
          let _rhs158 = _module.__default.fm26(_541_v0, (_719_v151)[_module.__default.safeIndex((_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))], new BigNumber((_719_v151).length))], _771_v192, globalState);
          let _lhs92 = _765_v191;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_765_v191).length));
          _758_v179 = _rhs157;
          _lhs92[_lhs93] = _rhs158;
          let _772_v193;
          let _nw145 = new _module.C0();
          _nw145.__ctor(p0);
          _772_v193 = _nw145;
          _541_v0 = _541_v0;
          let _773_v194;
          _773_v194 = _dafny.Set.fromElements(_751_cf0);
          let _774_v195;
          _774_v195 = _dafny.Seq.of(p0);
          let _775_v196;
          _775_v196 = _dafny.Set.fromElements(new BigNumber((_774_v195).length), (_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))]);
          let _776_v197;
          _776_v197 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_775_v196)).length),(_758_v179).f8)).length), new BigNumber(-505), (_dafny.ZERO).minus((_772_v193).f8), (_this).f5);
          let _777_v198;
          _777_v198 = _dafny.Seq.of(_756_v177, _756_v177, _757_v178, (_757_v178).update(new BigNumber((_774_v195).length), _751_cf0));
          let _778_v199;
          _778_v199 = _module.D5.create_DC18(_751_cf0, p0, false);
          let _pat_let_tv7 = _751_cf0;
          let _779_v200;
          _779_v200 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let9_0) {
            return function (_780_dt__update__tmp_h0) {
              return function (_pat_let10_0) {
                return function (_781_dt__update_hcf30_h0) {
                  return _module.D5.create_DC18(_781_dt__update_hcf30_h0, (_780_dt__update__tmp_h0).dtor_cf31, (_780_dt__update__tmp_h0).dtor_cf32);
                }(_pat_let10_0);
              }(_pat_let_tv7);
            }(_pat_let9_0);
          }(_778_v199),(_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))]);
          let _782_v201;
          let _nw146 = Array((new BigNumber(27)).toNumber()).fill(false);
          _782_v201 = _nw146;
          let _783_v202;
          let _nw147 = Array((new BigNumber(23)).toNumber());
          _nw147[(_dafny.ZERO).toNumber()] = (_this).f5;
          _nw147[(_dafny.ONE).toNumber()] = _this.f6;
          _nw147[(new BigNumber(2)).toNumber()] = (_this).fm2(p0, _773_v194, globalState);
          _nw147[(new BigNumber(3)).toNumber()] = (_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))];
          _nw147[(new BigNumber(4)).toNumber()] = new BigNumber(195);
          _nw147[(new BigNumber(5)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(6)).toNumber()] = _this.f6;
          _nw147[(new BigNumber(7)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(8)).toNumber()] = new BigNumber((_777_v198).length);
          _nw147[(new BigNumber(9)).toNumber()] = new BigNumber((_719_v151).length);
          _nw147[(new BigNumber(10)).toNumber()] = (_772_v193).f8;
          _nw147[(new BigNumber(11)).toNumber()] = (_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))];
          _nw147[(new BigNumber(12)).toNumber()] = _this.f6;
          _nw147[(new BigNumber(13)).toNumber()] = _this.f6;
          _nw147[(new BigNumber(14)).toNumber()] = new BigNumber(((_779_v200).update(_778_v199, _this.f6)).length);
          _nw147[(new BigNumber(15)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(16)).toNumber()] = (_this).f5;
          _nw147[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_591_v43,_782_v201)).length);
          _nw147[(new BigNumber(18)).toNumber()] = new BigNumber((_757_v178).length);
          _nw147[(new BigNumber(19)).toNumber()] = (_713_v147)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_713_v147).length))];
          _nw147[(new BigNumber(20)).toNumber()] = new BigNumber((_776_v197).length);
          _nw147[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus((_772_v193).f8);
          _nw147[(new BigNumber(22)).toNumber()] = new BigNumber(79);
          _783_v202 = _nw147;
          let _784_v203;
          _784_v203 = _module.D5.create_DC16(_783_v202);
          let _pat_let_tv8 = _783_v202;
          _784_v203 = function (_pat_let11_0) {
            return function (_785_dt__update__tmp_h1) {
              return function (_pat_let12_0) {
                return function (_786_dt__update_hcf26_h0) {
                  return _module.D5.create_DC16(_786_dt__update_hcf26_h0);
                }(_pat_let12_0);
              }(_pat_let_tv8);
            }(_pat_let11_0);
          }(_784_v203);
        }
        let _787_v204;
        let _nw148 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
        _787_v204 = _nw148;
        _787_v204 = _787_v204;
      } else {
        _541_v0 = ((_592_v44).Union(_592_v44)).IsSubsetOf(_592_v44);
        let _788_v205;
        _788_v205 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f6);
        _788_v205 = _788_v205;
        if (true) {
          (globalState).f2 = _this.f6;
          let _789_v206;
          let _nw149 = new _module.C0();
          _nw149.__ctor(_this.f6);
          _789_v206 = _nw149;
          let _790_v207;
          _790_v207 = _dafny.Set.fromElements(_789_v206);
          _541_v0 = !(((new BigNumber((_594_v46).length)).multipliedBy(new BigNumber((_790_v207).length))).isLessThanOrEqualTo(_this.f6));
          let _791_v208;
          _791_v208 = _dafny.Set.fromElements(_541_v0);
          let _792_v209;
          let _nw150 = Array((new BigNumber(25)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = _594_v46;
          _nw150[(_dafny.ONE).toNumber()] = _594_v46;
          _nw150[(new BigNumber(2)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(3)).toNumber()] = ((_541_v0) ? (_dafny.Seq.update(_594_v46, _module.__default.safeIndex(new BigNumber(-17), new BigNumber((_594_v46).length)), _541_v0)) : (_594_v46));
          _nw150[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_594_v46, _dafny.Seq.of(_541_v0, _541_v0));
          _nw150[(new BigNumber(5)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(6)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(7)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(8)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(9)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_541_v0), _594_v46);
          _nw150[(new BigNumber(11)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(12)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_594_v46, _module.__default.safeIndex(p0, new BigNumber((_594_v46).length)), _541_v0);
          _nw150[(new BigNumber(14)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_594_v46, _594_v46);
          _nw150[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_594_v46, _594_v46);
          _nw150[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_594_v46, _module.__default.safeIndex(new BigNumber((_791_v208).length), new BigNumber((_594_v46).length)), false);
          _nw150[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_594_v46, _module.__default.safeIndex(p0, new BigNumber((_594_v46).length)), _541_v0);
          _nw150[(new BigNumber(19)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(_541_v0);
          _nw150[(new BigNumber(21)).toNumber()] = _dafny.Seq.of(_541_v0, _541_v0, _541_v0, _541_v0, true);
          _nw150[(new BigNumber(22)).toNumber()] = _594_v46;
          _nw150[(new BigNumber(23)).toNumber()] = _dafny.Seq.of(false, _541_v0, _541_v0);
          _nw150[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_594_v46)[_module.__default.safeIndex((_789_v206).f8, new BigNumber((_594_v46).length))]), _dafny.Seq.of(_541_v0));
          _792_v209 = _nw150;
          let _index139 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_792_v209).length));
          (_792_v209)[_index139] = _594_v46;
          let _index140 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_792_v209).length));
          (_792_v209)[_index140] = _594_v46;
          let _793_v210;
          let _nw151 = Array((new BigNumber(24)).toNumber()).fill(false);
          _793_v210 = _nw151;
          let _index141 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_793_v210).length));
          (_793_v210)[_index141] = !_dafny.Seq.contains(_594_v46, _541_v0);
          let _index142 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_793_v210).length));
          let _rhs159 = _dafny.Seq.contains((_792_v209)[_module.__default.safeIndex(new BigNumber(914), new BigNumber((_792_v209).length))], !(_541_v0));
          let _rhs160 = (_789_v206).f8;
          let _rhs161 = _541_v0;
          let _rhs162 = _module.__default.safeModuloInt(((_789_v206).f8).plus((_789_v206).fm4(_591_v43, _592_v44, _541_v0, globalState)), (_789_v206).f8);
          let _lhs94 = globalState;
          let _lhs95 = _793_v210;
          let _lhs96 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_793_v210).length));
          let _lhs97 = globalState;
          _541_v0 = _rhs159;
          _lhs94.f2 = _rhs160;
          _lhs95[_lhs96] = _rhs161;
          _lhs97.f2 = _rhs162;
          let _794_v211;
          _794_v211 = _dafny.Seq.UnicodeFromString("nep");
          (_this).f6 = (_dafny.ZERO).minus((new BigNumber((_794_v211).length)).multipliedBy(new BigNumber(449)));
        } else {
          _713_v147 = _713_v147;
          let _795_v212;
          _795_v212 = _module.D7.create_DC24();
          _795_v212 = _795_v212;
          _541_v0 = !(_541_v0);
          let _796_v213;
          let _nw152 = new _module.C1();
          _nw152.__ctor(p0, (_this).f5);
          _796_v213 = _nw152;
          let _797_v214;
          _797_v214 = _dafny.Map.Empty.slice().updateUnsafe(_796_v213,_541_v0);
          let _index143 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_713_v147).length));
          (_713_v147)[_index143] = new BigNumber((_797_v214).length);
          let _798_v215;
          let _nw153 = Array((new BigNumber(23)).toNumber()).fill(false);
          _798_v215 = _nw153;
          let _index144 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_798_v215).length));
          (_798_v215)[_index144] = _541_v0;
          let _index145 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_713_v147).length));
          let _index146 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_798_v215).length));
          let _rhs163 = (_this).f5;
          let _rhs164 = _541_v0;
          let _lhs98 = _713_v147;
          let _lhs99 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_713_v147).length));
          let _lhs100 = _798_v215;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_798_v215).length));
          _lhs98[_lhs99] = _rhs163;
          _lhs100[_lhs101] = _rhs164;
          let _799_v216;
          _799_v216 = _dafny.Map.Empty.slice().updateUnsafe(_798_v215,_592_v44);
          _799_v216 = (_799_v216).update(_798_v215, _592_v44);
        }
        let _800_v217;
        let _nw154 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _800_v217 = _nw154;
        let _801_v218;
        _801_v218 = _dafny.Seq.UnicodeFromString("chcl");
        let _index147 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_800_v217).length));
        (_800_v217)[_index147] = _801_v218;
        let _index148 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_800_v217).length));
        (_800_v217)[_index148] = _801_v218;
        let _802_v219;
        _802_v219 = _dafny.Map.Empty.slice().updateUnsafe(_541_v0,_this.f6);
        let _803_v220;
        let _nw155 = new _module.C0();
        _nw155.__ctor((((_802_v219).contains(_541_v0)) ? ((_802_v219).get(_541_v0)) : (new BigNumber((_801_v218).length))));
        _803_v220 = _nw155;
      }
      return;
    }
    m4(p0, globalState) {
      let _this = this;
      if (((p0) ? (p0) : (p0))) {
        let _804_v0;
        _804_v0 = _dafny.Set.fromElements(_this.f6);
        let _805_v1;
        _805_v1 = _module.D1.create_DC4(_dafny.Set.fromElements(_this.f6));
        let _806_v2;
        _806_v2 = new _dafny.CodePoint('a'.codePointAt(0));
        let _807_v3;
        _807_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f6);
        let _808_v4;
        _808_v4 = _dafny.Map.Empty.slice().updateUnsafe(_806_v2,_807_v3);
        let _809_v5;
        let _nw156 = Array((new BigNumber(21)).toNumber());
        _nw156[(_dafny.ZERO).toNumber()] = p0;
        _nw156[(_dafny.ONE).toNumber()] = (_804_v0).contains((_this).f5);
        _nw156[(new BigNumber(2)).toNumber()] = p0;
        _nw156[(new BigNumber(3)).toNumber()] = p0;
        _nw156[(new BigNumber(4)).toNumber()] = false;
        _nw156[(new BigNumber(5)).toNumber()] = (!(p0)) && (p0);
        _nw156[(new BigNumber(6)).toNumber()] = (_804_v0).IsDisjointFrom((_805_v1).dtor_cf4);
        _nw156[(new BigNumber(7)).toNumber()] = !(true) || (p0);
        _nw156[(new BigNumber(8)).toNumber()] = p0;
        _nw156[(new BigNumber(9)).toNumber()] = !((_this.f6).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(false, false)).length)));
        _nw156[(new BigNumber(10)).toNumber()] = !((_808_v4).contains(_806_v2));
        _nw156[(new BigNumber(11)).toNumber()] = true;
        _nw156[(new BigNumber(12)).toNumber()] = p0;
        _nw156[(new BigNumber(13)).toNumber()] = p0;
        _nw156[(new BigNumber(14)).toNumber()] = true;
        _nw156[(new BigNumber(15)).toNumber()] = p0;
        _nw156[(new BigNumber(16)).toNumber()] = p0;
        _nw156[(new BigNumber(17)).toNumber()] = p0;
        _nw156[(new BigNumber(18)).toNumber()] = p0;
        _nw156[(new BigNumber(19)).toNumber()] = p0;
        _nw156[(new BigNumber(20)).toNumber()] = p0;
        _809_v5 = _nw156;
        let _index149 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length));
        (_809_v5)[_index149] = ((_dafny.ZERO).minus((_this).f5)).isLessThanOrEqualTo((_this).f5);
        let _index150 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length));
        (_809_v5)[_index150] = (_this).fm9(_module.D1.create_DC5(), ((_this).fm1(globalState)) === (p0), globalState);
        let _index151 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length));
        (_809_v5)[_index151] = true;
        let _810_v6;
        _810_v6 = _dafny.Set.fromElements(p0, (_this.f6).isEqualTo(new BigNumber((_804_v0).length)), p0);
        _810_v6 = ((_dafny.Set.fromElements((_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))], (_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))])).Difference(_810_v6)).Union(_810_v6);
        let _811_v7;
        _811_v7 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f5)).length));
        let _812_v8;
        _812_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,new BigNumber((_811_v7).length));
        let _813_v9;
        _813_v9 = _dafny.Seq.of(p0, (_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))], (_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))]);
        let _814_v10;
        _814_v10 = _dafny.MultiSet.fromElements(p0, !((_813_v9)[_module.__default.safeIndex(_this.f6, new BigNumber((_813_v9).length))]), (_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))]);
        let _815_v11;
        let _nw157 = Array((new BigNumber(18)).toNumber());
        _nw157[(_dafny.ZERO).toNumber()] = _this.f6;
        _nw157[(_dafny.ONE).toNumber()] = new BigNumber((_812_v8).length);
        _nw157[(new BigNumber(2)).toNumber()] = (_this).f5;
        _nw157[(new BigNumber(3)).toNumber()] = (_this).f5;
        _nw157[(new BigNumber(4)).toNumber()] = (_this).f5;
        _nw157[(new BigNumber(5)).toNumber()] = (_this).f5;
        _nw157[(new BigNumber(6)).toNumber()] = _this.f6;
        _nw157[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_811_v7)).cardinality());
        _nw157[(new BigNumber(8)).toNumber()] = _this.f6;
        _nw157[(new BigNumber(9)).toNumber()] = new BigNumber(199);
        _nw157[(new BigNumber(10)).toNumber()] = (_this).f5;
        _nw157[(new BigNumber(11)).toNumber()] = (_this).f5;
        _nw157[(new BigNumber(12)).toNumber()] = _this.f6;
        _nw157[(new BigNumber(13)).toNumber()] = (_this).f5;
        _nw157[(new BigNumber(14)).toNumber()] = _this.f6;
        _nw157[(new BigNumber(15)).toNumber()] = (_this).f5;
        _nw157[(new BigNumber(16)).toNumber()] = new BigNumber((_814_v10).cardinality());
        _nw157[(new BigNumber(17)).toNumber()] = (_this).f5;
        _815_v11 = _nw157;
        let _816_v12;
        _816_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,!((_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))]));
        let _817_v13;
        _817_v13 = _module.D5.create_DC17(new BigNumber((_module.__default.fm11((_this).f5, (_this).f5, (_this).f5, (_dafny.ZERO).minus(new BigNumber((_811_v7).length)), globalState)).length), _815_v11, _816_v12);
        let _818_v16;
        _818_v16 = _dafny.Map.Empty.slice().updateUnsafe(_806_v2,(_this).f5);
        let _pat_let_tv9 = _815_v11;
        let _819_v18;
        let _nw158 = Array((new BigNumber(28)).toNumber());
        _nw158[(_dafny.ZERO).toNumber()] = _817_v13;
        _nw158[(_dafny.ONE).toNumber()] = _module.D5.create_DC17((_this).f5, _815_v11, function () {
  let _coll43 = new _dafny.Map();
  for (const _compr_43 of _dafny.IntegerRange(new BigNumber(171), new BigNumber(954))) {
    let _820_v14 = _compr_43;
    if (((new BigNumber(171)).isLessThanOrEqualTo(_820_v14)) && ((_820_v14).isLessThan(new BigNumber(954)))) {
      _coll43.push([(_820_v14).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("ljpgjdur")).length)),(_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))]]);
    }
  }
  return _coll43;
}());
        _nw158[(new BigNumber(2)).toNumber()] = _module.D5.create_DC17(new BigNumber((_814_v10).cardinality()), _815_v11, function () {
  let _coll44 = new _dafny.Map();
  for (const _compr_44 of (_804_v0).Elements) {
    let _821_v15 = _compr_44;
    if ((_804_v0).contains(_821_v15)) {
      _coll44.push([(_821_v15).plus(new BigNumber((_816_v12).length)),(_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))]]);
    }
  }
  return _coll44;
}());
        _nw158[(new BigNumber(3)).toNumber()] = (((_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))]) ? (_817_v13) : (_817_v13));
        _nw158[(new BigNumber(4)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(5)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(6)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(7)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(8)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(9)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(10)).toNumber()] = function (_pat_let13_0) {
          return function (_822_dt__update__tmp_h0) {
            return function (_pat_let14_0) {
              return function (_823_dt__update_hcf28_h0) {
                return _module.D5.create_DC17((_822_dt__update__tmp_h0).dtor_cf27, _823_dt__update_hcf28_h0, (_822_dt__update__tmp_h0).dtor_cf29);
              }(_pat_let14_0);
            }(_pat_let_tv9);
          }(_pat_let13_0);
        }(_817_v13);
        _nw158[(new BigNumber(11)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(12)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(13)).toNumber()] = _module.D5.create_DC17((_dafny.ZERO).minus(_this.f6), _815_v11, _816_v12);
        _nw158[(new BigNumber(14)).toNumber()] = _module.D5.create_DC17((((_818_v16).contains(_806_v2)) ? ((_818_v16).get(_806_v2)) : ((_this).f5)), _815_v11, function () {
  let _coll45 = new _dafny.Map();
  for (const _compr_45 of (_816_v12).Keys.Elements) {
    let _824_v17 = _compr_45;
    if ((_816_v12).contains(_824_v17)) {
      _coll45.push([_module.__default.safeModuloInt(_824_v17, (_this).f5),(_809_v5)[_module.__default.safeIndex(new BigNumber(24), new BigNumber((_809_v5).length))]]);
    }
  }
  return _coll45;
}());
        _nw158[(new BigNumber(15)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(16)).toNumber()] = _module.D5.create_DC17(new BigNumber(639), _815_v11, _816_v12);
        _nw158[(new BigNumber(17)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(18)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(19)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(20)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(21)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(22)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(23)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(24)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(25)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(26)).toNumber()] = _817_v13;
        _nw158[(new BigNumber(27)).toNumber()] = _module.D5.create_DC17(_this.f6, _815_v11, _816_v12);
        _819_v18 = _nw158;
        let _nw159 = Array((new BigNumber(6)).toNumber());
        _nw159[(_dafny.ZERO).toNumber()] = _817_v13;
        _nw159[(_dafny.ONE).toNumber()] = _817_v13;
        _nw159[(new BigNumber(2)).toNumber()] = _817_v13;
        _nw159[(new BigNumber(3)).toNumber()] = _817_v13;
        _nw159[(new BigNumber(4)).toNumber()] = _817_v13;
        _nw159[(new BigNumber(5)).toNumber()] = _817_v13;
        _819_v18 = _nw159;
        _814_v10 = _814_v10;
      } else {
        let _825_v19;
        _825_v19 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,_this.f6);
        let _826_v20;
        _826_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        (globalState).f2 = _module.__default.safeModuloInt((((_825_v19).contains(new BigNumber((_module.__default.fm10(new BigNumber(-670), _825_v19, p0, _this.f6, globalState)).length))) ? ((_825_v19).get(new BigNumber((_module.__default.fm10(new BigNumber(-670), _825_v19, p0, _this.f6, globalState)).length))) : (_this.f6)), new BigNumber((_826_v20).length));
        let _827_v21;
        _827_v21 = new _dafny.CodePoint('h'.codePointAt(0));
        _827_v21 = _827_v21;
        (globalState).f2 = _this.f6;
        let _828_v22;
        _828_v22 = false;
        let _829_v23;
        _829_v23 = _dafny.Seq.UnicodeFromString("d");
        let _830_v24;
        _830_v24 = _dafny.Seq.of(_this.f6, new BigNumber((_dafny.Seq.UnicodeFromString("ml")).length), (_this).f5, new BigNumber((_829_v23).length));
        let _831_v25;
        _831_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,p0);
        let _832_v26;
        _832_v26 = _dafny.Seq.of(!(!(_828_v22)), (((_831_v25).contains((_this).f5)) ? ((_831_v25).get((_this).f5)) : (p0)));
        _828_v22 = (((_this).f5).minus(new BigNumber((_830_v24).length))).isEqualTo(new BigNumber((_832_v26).length));
        let _833_v27;
        _833_v27 = _dafny.Set.fromElements(new BigNumber(880), new BigNumber((_module.__default.fm11((_this).f5, new BigNumber(574), _this.f6, _this.f6, globalState)).length), _this.f6);
        _833_v27 = _dafny.Set.fromElements((_this).f5, (_dafny.ZERO).minus(_this.f6), _module.__default.safeModuloInt((_this).f5, (_dafny.ZERO).minus(_this.f6)));
      }
      (_this).f6 = (_this).f5;
      let _834_v28;
      _834_v28 = false;
      _834_v28 = p0;
      let _835_v30;
      _835_v30 = _dafny.Seq.of(_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(true,true)));
      let _836_v31;
      _836_v31 = _module.D7.create_DC21(function () {
  let _coll46 = new _dafny.Map();
  for (const _compr_46 of (_835_v30).Elements) {
    let _837_v29 = _compr_46;
    if (_dafny.Seq.contains(_835_v30, _837_v29)) {
      _coll46.push([_837_v29,false]);
    }
  }
  return _coll46;
}());
      let _pat_let_tv10 = _834_v28;
      let _pat_let_tv11 = _834_v28;
      let _pat_let_tv12 = p0;
      let _pat_let_tv13 = p0;
      if (function (_source11) {
        if (_source11.is_DC22) {
          let _838___mcc_h1 = (_source11).cf38;
          let _839___mcc_h2 = (_source11).cf39;
          let _840_cf39 = _839___mcc_h2;
          let _841_cf38 = _838___mcc_h1;
          return _pat_let_tv10;
        } else if (_source11.is_DC23) {
          let _842___mcc_h3 = (_source11).cf40;
          let _843___mcc_h4 = (_source11).cf41;
          let _844___mcc_h5 = (_source11).cf42;
          let _845_cf42 = _844___mcc_h5;
          let _846_cf41 = _843___mcc_h4;
          let _847_cf40 = _842___mcc_h3;
          return !(_pat_let_tv11);
        } else if (_source11.is_DC24) {
          return _pat_let_tv12;
        } else if (_source11.is_DC21) {
          let _848___mcc_h6 = (_source11).cf37;
          let _849_cf37 = _848___mcc_h6;
          return true;
        } else {
          let _850___mcc_h7 = (_source11).cf43;
          let _851_cf43 = _850___mcc_h7;
          return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(new BigNumber(-767), new BigNumber((_dafny.Set.fromElements(_pat_let_tv13)).length)), _dafny.Seq.of(_this.f6));
        }
      }(_836_v31)) {
        if (p0) {
          _834_v28 = _834_v28;
          let _852_v32;
          _852_v32 = new _dafny.CodePoint('b'.codePointAt(0));
          let _853_v33;
          _853_v33 = _dafny.MultiSet.fromElements(_852_v32, _852_v32, _852_v32);
          let _854_v34;
          _854_v34 = _module.D9.create_DC27(_853_v33);
          _853_v33 = (_853_v33).Union((_854_v34).dtor_cf45);
          let _855_v35;
          let _nw160 = new _module.C0();
          _nw160.__ctor(_module.__default.safeModuloInt((_this).f5, new BigNumber(541)));
          _855_v35 = _nw160;
          _834_v28 = p0;
          _834_v28 = ((_855_v35).f8).isLessThan((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_this.f6)).length)).multipliedBy((_this).f5));
        } else {
          let _856_v36;
          _856_v36 = _dafny.Seq.UnicodeFromString("b");
          _856_v36 = _856_v36;
          let _857_v37;
          let _init18 = ((_858_v36) => function (_859_i0) {
            return (_859_i0).minus(new BigNumber((_858_v36).length));
          })(_856_v36);
          let _nw161 = Array((new BigNumber(6)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw161.length); _i0_18++) {
            _nw161[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _857_v37 = _nw161;
          _857_v37 = _857_v37;
          let _860_v38;
          _860_v38 = _dafny.Set.fromElements(false, _834_v28);
          let _861_v39;
          let _nw162 = new _module.C1();
          _nw162.__ctor(((_this).fm2(_this.f6, _860_v38, globalState)).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f6))), ((_834_v28) ? (_this.f6) : (new BigNumber(941))));
          _861_v39 = _nw162;
          let _862_v40;
          let _nw163 = new _module.C1();
          _nw163.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("srfo")).length), (_dafny.ZERO).minus((_861_v39).fm13((_dafny.ZERO).minus(_this.f6), _dafny.MultiSet.fromElements(_856_v36, _856_v36), globalState)));
          _862_v40 = _nw163;
          let _863_v41;
          _863_v41 = _module.D0.create_DC1(false, _834_v28);
          let _864_v42;
          _864_v42 = _module.D9.create_DC28(_834_v28, _this.f6, (_862_v40).fm13(new BigNumber(971), _dafny.MultiSet.fromElements(_856_v36, _856_v36), globalState));
          let _865_v43;
          _865_v43 = _dafny.Seq.of(p0);
          let _866_v44;
          let _nw164 = Array((new BigNumber(26)).toNumber());
          _nw164[(_dafny.ZERO).toNumber()] = _834_v28;
          _nw164[(_dafny.ONE).toNumber()] = _834_v28;
          _nw164[(new BigNumber(2)).toNumber()] = _834_v28;
          _nw164[(new BigNumber(3)).toNumber()] = _834_v28;
          _nw164[(new BigNumber(4)).toNumber()] = p0;
          _nw164[(new BigNumber(5)).toNumber()] = (_863_v41).dtor_cf2;
          _nw164[(new BigNumber(6)).toNumber()] = _834_v28;
          _nw164[(new BigNumber(7)).toNumber()] = p0;
          _nw164[(new BigNumber(8)).toNumber()] = false;
          _nw164[(new BigNumber(9)).toNumber()] = _834_v28;
          _nw164[(new BigNumber(10)).toNumber()] = p0;
          _nw164[(new BigNumber(11)).toNumber()] = p0;
          _nw164[(new BigNumber(12)).toNumber()] = p0;
          _nw164[(new BigNumber(13)).toNumber()] = _834_v28;
          _nw164[(new BigNumber(14)).toNumber()] = p0;
          _nw164[(new BigNumber(15)).toNumber()] = p0;
          _nw164[(new BigNumber(16)).toNumber()] = (_864_v42).dtor_cf46;
          _nw164[(new BigNumber(17)).toNumber()] = _834_v28;
          _nw164[(new BigNumber(18)).toNumber()] = true;
          _nw164[(new BigNumber(19)).toNumber()] = p0;
          _nw164[(new BigNumber(20)).toNumber()] = p0;
          _nw164[(new BigNumber(21)).toNumber()] = (_865_v43)[_module.__default.safeIndex((_this).f5, new BigNumber((_865_v43).length))];
          _nw164[(new BigNumber(22)).toNumber()] = p0;
          _nw164[(new BigNumber(23)).toNumber()] = p0;
          _nw164[(new BigNumber(24)).toNumber()] = p0;
          _nw164[(new BigNumber(25)).toNumber()] = _834_v28;
          _866_v44 = _nw164;
          let _867_v45;
          _867_v45 = _dafny.Seq.of(_866_v44, _866_v44);
          let _868_v46;
          _868_v46 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f5);
          let _869_v48;
          _869_v48 = new _dafny.CodePoint('b'.codePointAt(0));
          let _870_v49;
          _870_v49 = _dafny.Map.Empty.slice().updateUnsafe(_868_v46,_869_v48);
          let _871_v50;
          _871_v50 = _dafny.Map.Empty.slice().updateUnsafe(_867_v45,(_dafny.Map.Empty.slice().updateUnsafe(_868_v46,(_856_v36)[_module.__default.safeIndex(new BigNumber((function () {
            let _coll47 = new _dafny.Set();
            for (const _compr_47 of _dafny.IntegerRange(new BigNumber(340), new BigNumber(549))) {
              let _872_v47 = _compr_47;
              if (((new BigNumber(340)).isLessThanOrEqualTo(_872_v47)) && ((_872_v47).isLessThan(new BigNumber(549)))) {
                _coll47.add(_module.__default.safeModuloInt(_872_v47, new BigNumber((_865_v43).length)));
              }
            }
            return _coll47;
          }()).length), new BigNumber((_856_v36).length))])).equals(_870_v49));
          _871_v50 = (_871_v50).update(_867_v45, _834_v28);
        }
        _834_v28 = !(p0) || ((_dafny.MultiSet.fromElements(_834_v28)).IsSubsetOf(_dafny.MultiSet.fromElements(_834_v28)));
        let _873_v51;
        _873_v51 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f5);
        let _874_v52;
        _874_v52 = (_873_v51).Merge(_873_v51);
        let _source12 = _874_v52;
        let _875___mcc_h0 = _source12;
        let _876_cf7 = _875___mcc_h0;
        let _877_v53;
        _877_v53 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f5), (_this).f5);
        let _878_v54;
        _878_v54 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,new BigNumber((_877_v53).length));
        let _879_v55;
        _879_v55 = new _dafny.CodePoint('a'.codePointAt(0));
        let _880_v56;
        _880_v56 = _dafny.Map.Empty.slice().updateUnsafe(_878_v54,_879_v55);
        _880_v56 = _880_v56;
        let _881_v58;
        _881_v58 = _module.D10.create_DC29(_878_v54);
        let _882_v59;
        _882_v59 = _dafny.Seq.UnicodeFromString("dcfhvxpm");
        let _883_v60;
        _883_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,_882_v59);
        let _884_v61;
        let _nw165 = Array((new BigNumber(26)).toNumber());
        _nw165[(_dafny.ZERO).toNumber()] = _878_v54;
        _nw165[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f6);
        _nw165[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_876_cf7).length),(_this).f5);
        _nw165[(new BigNumber(3)).toNumber()] = function () {
          let _coll48 = new _dafny.Map();
          for (const _compr_48 of _dafny.IntegerRange(new BigNumber(353), new BigNumber(769))) {
            let _885_v57 = _compr_48;
            if (((new BigNumber(353)).isLessThanOrEqualTo(_885_v57)) && ((_885_v57).isLessThan(new BigNumber(769)))) {
              _coll48.push([(_885_v57).plus(_this.f6),_this.f6]);
            }
          }
          return _coll48;
        }();
        _nw165[(new BigNumber(4)).toNumber()] = (_881_v58).dtor_cf49;
        _nw165[(new BigNumber(5)).toNumber()] = (_878_v54).Merge(_878_v54);
        _nw165[(new BigNumber(6)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(7)).toNumber()] = (_878_v54).Merge(_878_v54);
        _nw165[(new BigNumber(8)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(9)).toNumber()] = (_878_v54).Merge(_878_v54);
        _nw165[(new BigNumber(10)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(11)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(12)).toNumber()] = (_878_v54).update(_this.f6, new BigNumber((_883_v60).length));
        _nw165[(new BigNumber(13)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(14)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(15)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(16)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(17)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(18)).toNumber()] = (_878_v54).Merge(_878_v54);
        _nw165[(new BigNumber(19)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(20)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(21)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(22)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(23)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(24)).toNumber()] = _878_v54;
        _nw165[(new BigNumber(25)).toNumber()] = _878_v54;
        _884_v61 = _nw165;
        let _886_v63;
        _886_v63 = _dafny.Set.fromElements(_this.f6);
        let _index152 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_884_v61).length));
        (_884_v61)[_index152] = function () {
          let _coll49 = new _dafny.Map();
          for (const _compr_49 of (_886_v63).Elements) {
            let _887_v62 = _compr_49;
            if ((_886_v63).contains(_887_v62)) {
              _coll49.push([_module.__default.safeDivisionInt(_887_v62, _this.f6),(_this).f5]);
            }
          }
          return _coll49;
        }();
        let _index153 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_884_v61).length));
        (_884_v61)[_index153] = (_878_v54).Merge(function () {
          let _coll50 = new _dafny.Map();
          for (const _compr_50 of (_878_v54).Keys.Elements) {
            let _888_v64 = _compr_50;
            if ((_878_v54).contains(_888_v64)) {
              _coll50.push([_module.__default.safeModuloInt(_888_v64, (_this).f5),(_this).f5]);
            }
          }
          return _coll50;
        }());
        _882_v59 = _882_v59;
        (globalState).f2 = (_this).f5;
        let _889_v65;
        _889_v65 = _dafny.MultiSet.fromElements((_this).f5);
        if ((_889_v65).IsDisjointFrom(_889_v65)) {
          _889_v65 = _889_v65;
          let _890_v66;
          let _nw166 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _890_v66 = _nw166;
          let _891_v67;
          _891_v67 = _dafny.Set.fromElements(p0, p0, _834_v28, !(p0));
          let _index154 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_890_v66).length));
          (_890_v66)[_index154] = (_this).fm2(_this.f6, _891_v67, globalState);
          let _892_v68;
          _892_v68 = _dafny.Seq.of(p0);
          let _index155 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_890_v66).length));
          (_890_v66)[_index155] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_892_v68),_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-838)), function (_893_i1) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }), _module.__default.safeIndex(new BigNumber(620), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-838)), function (_894_i1) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          })).length)), new _dafny.CodePoint('l'.codePointAt(0))))).length)).multipliedBy((_this).f5);
          let _895_v69;
          _895_v69 = new _dafny.CodePoint('g'.codePointAt(0));
          let _896_v70;
          _896_v70 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,_895_v69);
          _895_v69 = (((_896_v70).contains(((p0) ? ((_890_v66)[_module.__default.safeIndex(new BigNumber(761), new BigNumber((_890_v66).length))]) : (_this.f6)))) ? ((_896_v70).get(((p0) ? ((_890_v66)[_module.__default.safeIndex(new BigNumber(761), new BigNumber((_890_v66).length))]) : (_this.f6)))) : (new _dafny.CodePoint('e'.codePointAt(0))));
          let _897_v71;
          let _init19 = ((_898_p0) => function (_899_i2) {
            return _898_p0;
          })(p0);
          let _nw167 = Array((new BigNumber(7)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw167.length); _i0_19++) {
            _nw167[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _897_v71 = _nw167;
          let _index156 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_897_v71).length));
          (_897_v71)[_index156] = (_this).fm1(globalState);
          let _900_v72;
          _900_v72 = _module.D1.create_DC5();
          let _index157 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_897_v71).length));
          let _rhs165 = p0;
          let _rhs166 = (_this).f5;
          let _rhs167 = ((_this).fm9(_900_v72, true, globalState)) && (_834_v28);
          let _lhs102 = _897_v71;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_897_v71).length));
          let _lhs104 = globalState;
          _lhs102[_lhs103] = _rhs165;
          _lhs104.f2 = _rhs166;
          _834_v28 = _rhs167;
          let _901_v73;
          let _902_v74;
          let _903_v75;
          let _904_v76;
          let _out52;
          let _out53;
          let _out54;
          let _out55;
          let _outcollector15 = _module.__default.m0(new _dafny.CodePoint('m'.codePointAt(0)), globalState);
          _out52 = _outcollector15[0];
          _out53 = _outcollector15[1];
          _out54 = _outcollector15[2];
          _out55 = _outcollector15[3];
          _901_v73 = _out52;
          _902_v74 = _out53;
          _903_v75 = _out54;
          _904_v76 = _out55;
        } else {
          let _905_v77;
          _905_v77 = _module.D1.create_DC5();
          let _906_v78;
          _906_v78 = _dafny.Map.Empty.slice().updateUnsafe(_905_v77,((_this).f5).minus((_this).f5));
          _906_v78 = (_906_v78).update(_905_v77, (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f6)));
          let _907_v79;
          let _nw168 = new _module.C0();
          _nw168.__ctor(new BigNumber(780));
          _907_v79 = _nw168;
          let _908_v80;
          _908_v80 = _dafny.Seq.of(_834_v28, p0);
          let _909_v81;
          _909_v81 = _dafny.Seq.of(new BigNumber((_908_v80).length), _this.f6);
          let _910_v82;
          let _nw169 = new _module.C0();
          _nw169.__ctor((_this).fm2((_this).f5, _module.__default.fm21(new BigNumber(-113), new BigNumber((_909_v81).length), false, globalState), globalState));
          _910_v82 = _nw169;
          let _911_v83;
          _911_v83 = _dafny.Seq.of(_907_v79);
          _910_v82 = (_911_v83)[_module.__default.safeIndex((_907_v79).f8, new BigNumber((_911_v83).length))];
          let _912_v84;
          let _nw170 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
          _912_v84 = _nw170;
          let _index158 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_912_v84).length));
          (_912_v84)[_index158] = _909_v81;
          let _913_v85;
          _913_v85 = _dafny.Set.fromElements(_889_v65, _889_v65);
          let _914_v86;
          let _nw171 = new _module.C1();
          _nw171.__ctor(new BigNumber((_913_v85).length), (_910_v82).f8);
          _914_v86 = _nw171;
          let _915_v87;
          _915_v87 = _dafny.Seq.of(_914_v86);
          let _916_v88;
          _916_v88 = _module.D11.create_DC31(_914_v86);
          let _917_v89;
          let _nw172 = Array((new BigNumber(23)).toNumber());
          _nw172[(_dafny.ZERO).toNumber()] = _914_v86;
          _nw172[(_dafny.ONE).toNumber()] = _914_v86;
          _nw172[(new BigNumber(2)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(3)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(4)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(5)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(6)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(7)).toNumber()] = (_915_v87)[_module.__default.safeIndex((_910_v82).f8, new BigNumber((_915_v87).length))];
          _nw172[(new BigNumber(8)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(9)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(10)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(11)).toNumber()] = (_916_v88).dtor_cf51;
          _nw172[(new BigNumber(12)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(13)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(14)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(15)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(16)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(17)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(18)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(19)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(20)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(21)).toNumber()] = _914_v86;
          _nw172[(new BigNumber(22)).toNumber()] = _914_v86;
          _917_v89 = _nw172;
          let _918_v90;
          _918_v90 = _dafny.Set.fromElements(_917_v89, _917_v89, _917_v89, _917_v89, _917_v89);
          let _index159 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_912_v84).length));
          (_912_v84)[_index159] = _dafny.Seq.of(new BigNumber(((_dafny.Set.fromElements(_917_v89)).Difference(_918_v90)).length), (_910_v82).f8);
          let _919_v91;
          let _init20 = function (_920_i3) {
            return _module.__default.safeModuloInt(_920_i3, _this.f6);
          };
          let _nw173 = Array((new BigNumber(28)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw173.length); _i0_20++) {
            _nw173[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _919_v91 = _nw173;
          let _921_v92;
          _921_v92 = _module.D5.create_DC16(_919_v91);
          let _922_v93;
          let _nw174 = Array((new BigNumber(24)).toNumber());
          _nw174[(_dafny.ZERO).toNumber()] = _919_v91;
          _nw174[(_dafny.ONE).toNumber()] = _919_v91;
          _nw174[(new BigNumber(2)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(3)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(4)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(5)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(6)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(7)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(8)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(9)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(10)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(11)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(12)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(13)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(14)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(15)).toNumber()] = (_921_v92).dtor_cf26;
          _nw174[(new BigNumber(16)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(17)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(18)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(19)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(20)).toNumber()] = (_921_v92).dtor_cf26;
          _nw174[(new BigNumber(21)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(22)).toNumber()] = _919_v91;
          _nw174[(new BigNumber(23)).toNumber()] = _919_v91;
          _922_v93 = _nw174;
          let _index160 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_922_v93).length));
          (_922_v93)[_index160] = _919_v91;
          let _index161 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_922_v93).length));
          let _rhs168 = (_dafny.ZERO).minus((_907_v79).f8);
          let _rhs169 = _919_v91;
          let _lhs105 = _this;
          let _lhs106 = _922_v93;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_922_v93).length));
          _lhs105.f6 = _rhs168;
          _lhs106[_lhs107] = _rhs169;
        }
        let _923_v94;
        _923_v94 = new _dafny.CodePoint('u'.codePointAt(0));
        _923_v94 = _923_v94;
      } else {
        (globalState).f2 = (_this).f5;
        let _924_v95;
        let _nw175 = Array((_dafny.ONE).toNumber());
        _nw175[(_dafny.ZERO).toNumber()] = p0;
        _924_v95 = _nw175;
        let _index162 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length));
        (_924_v95)[_index162] = (_this.f6).isLessThanOrEqualTo(_this.f6);
        let _index163 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length));
        (_924_v95)[_index163] = p0;
        _834_v28 = p0;
        let _925_v96;
        _925_v96 = _dafny.Seq.UnicodeFromString("hjndu");
        _925_v96 = _925_v96;
        if ((_924_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length))]) {
          _834_v28 = ((_924_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length))]) === (p0);
          let _926_v97;
          _926_v97 = _dafny.MultiSet.fromElements((_924_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length))]);
          let _927_v98;
          _927_v98 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_925_v96).length)));
          let _928_v99;
          _928_v99 = _dafny.MultiSet.fromElements((_926_v97).IsSubsetOf(_926_v97), (_927_v98).IsDisjointFrom(_927_v98));
          let _929_v100;
          _929_v100 = _dafny.Seq.of((_924_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length))]);
          let _930_v101;
          _930_v101 = _dafny.Seq.of(_927_v98);
          _926_v97 = ((_928_v99).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.update(_929_v100, _module.__default.safeIndex((_dafny.ZERO).minus(_this.f6), new BigNumber((_929_v100).length)), p0)))).update((_924_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length))], _module.__default.abs(((false) ? ((_this).f5) : (new BigNumber(((_930_v101)[_module.__default.safeIndex((_this).f5, new BigNumber((_930_v101).length))]).length)))));
          let _index164 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length));
          (_924_v95)[_index164] = !(_834_v28) || ((_924_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length))]);
          let _931_v102;
          let _nw176 = new _module.C1();
          _nw176.__ctor((_this.f6).plus((_this).f5), _this.f6);
          _931_v102 = _nw176;
          let _index165 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length));
          (_924_v95)[_index165] = _834_v28;
        } else {
          let _932_v103;
          let _nw177 = new _module.C1();
          _nw177.__ctor((_this).f5, (_this).f5);
          _932_v103 = _nw177;
          let _933_v104;
          _933_v104 = _dafny.Map.Empty.slice().updateUnsafe(_932_v103,p0);
          let _934_v105;
          _934_v105 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(808)), function (_935_i4) {
            return (_this).f5;
          }),_933_v104);
          let _936_v106;
          let _init21 = ((_937_v95) => function (_938_i5) {
            return _dafny.Seq.Concat(_dafny.Seq.of((_937_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_937_v95).length))], (_937_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_937_v95).length))]), _dafny.Seq.of((_937_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_937_v95).length))]));
          })(_924_v95);
          let _nw178 = Array((new BigNumber(3)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw178.length); _i0_21++) {
            _nw178[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _936_v106 = _nw178;
          let _939_v107;
          _939_v107 = _dafny.Seq.of(_this.f6);
          let _rhs170 = _this.f6;
          let _rhs171 = _dafny.Map.Empty.slice().updateUnsafe(_939_v107,_933_v104);
          let _rhs172 = _936_v106;
          let _lhs108 = globalState;
          _lhs108.f2 = _rhs170;
          _934_v105 = _rhs171;
          _936_v106 = _rhs172;
          _925_v96 = _dafny.Seq.UnicodeFromString("kobrta");
          let _940_v108;
          _940_v108 = _dafny.Seq.of((_924_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length))], false);
          _940_v108 = (((p0) === ((_924_v95)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length))])) ? (_dafny.Seq.Concat(_940_v108, _940_v108)) : (_940_v108));
          let _941_v109;
          _941_v109 = _module.D5.create_DC18(_834_v28, (_this).f5, !(_834_v28));
          let _rhs173 = (new BigNumber((_939_v107).length)).minus(_this.f6);
          let _rhs174 = _941_v109;
          let _lhs109 = _this;
          _lhs109.f6 = _rhs173;
          _941_v109 = _rhs174;
          let _index166 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_924_v95).length));
          (_924_v95)[_index166] = p0;
        }
      }
      let _942_v110;
      _942_v110 = _dafny.Seq.UnicodeFromString("gahubgf");
      let _943_v111;
      _943_v111 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_942_v110).length),true);
      let _944_v112;
      _944_v112 = _module.D0.create_DC1(_834_v28, _834_v28);
      let _945_v113;
      _945_v113 = _dafny.MultiSet.fromElements(false);
      let _946_v114;
      let _nw179 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _946_v114 = _nw179;
      let _947_v115;
      _947_v115 = _dafny.Map.Empty.slice().updateUnsafe(_946_v114,_834_v28);
      let _948_v116;
      _948_v116 = new _dafny.CodePoint('a'.codePointAt(0));
      let _949_v117;
      _949_v117 = _dafny.Set.fromElements(_this.f6, _this.f6);
      let _950_v118;
      _950_v118 = _dafny.Seq.of((_this).f5, _this.f6, (_this).f5);
      let _951_v119;
      _951_v119 = _dafny.Map.Empty.slice().updateUnsafe(p0,_943_v111);
      let _952_v120;
      _952_v120 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_950_v118).length), _this.f6)).cardinality()),p0), (((_951_v119).contains(p0)) ? ((_951_v119).get(p0)) : (_943_v111)));
      let _953_v121;
      let _nw180 = Array((new BigNumber(26)).toNumber());
      _nw180[(_dafny.ZERO).toNumber()] = _834_v28;
      _nw180[(_dafny.ONE).toNumber()] = (_this).fm0(_dafny.MultiSet.fromElements(_943_v111, _943_v111), p0, _942_v110, globalState);
      _nw180[(new BigNumber(2)).toNumber()] = !(p0);
      _nw180[(new BigNumber(3)).toNumber()] = (_944_v112).dtor_cf1;
      _nw180[(new BigNumber(4)).toNumber()] = _834_v28;
      _nw180[(new BigNumber(5)).toNumber()] = !(!(_834_v28));
      _nw180[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.fromElements(p0)).equals(_945_v113);
      _nw180[(new BigNumber(7)).toNumber()] = (((_947_v115).contains(_946_v114)) ? ((_947_v115).get(_946_v114)) : (_834_v28));
      _nw180[(new BigNumber(8)).toNumber()] = p0;
      _nw180[(new BigNumber(9)).toNumber()] = true;
      _nw180[(new BigNumber(10)).toNumber()] = !(p0);
      _nw180[(new BigNumber(11)).toNumber()] = !_dafny.areEqual(_dafny.Seq.update(_module.__default.fm11(_this.f6, (_this).f5, new BigNumber((_945_v113).cardinality()), _this.f6, globalState), _module.__default.safeIndex((_this).f5, new BigNumber((_module.__default.fm11(_this.f6, (_this).f5, new BigNumber((_945_v113).cardinality()), _this.f6, globalState)).length)), _948_v116), _942_v110);
      _nw180[(new BigNumber(12)).toNumber()] = (_949_v117).IsSubsetOf(_949_v117);
      _nw180[(new BigNumber(13)).toNumber()] = p0;
      _nw180[(new BigNumber(14)).toNumber()] = p0;
      _nw180[(new BigNumber(15)).toNumber()] = p0;
      _nw180[(new BigNumber(16)).toNumber()] = (_this).fm0(_952_v120, p0, _942_v110, globalState);
      _nw180[(new BigNumber(17)).toNumber()] = _834_v28;
      _nw180[(new BigNumber(18)).toNumber()] = _834_v28;
      _nw180[(new BigNumber(19)).toNumber()] = _834_v28;
      _nw180[(new BigNumber(20)).toNumber()] = (_945_v113).IsDisjointFrom(_945_v113);
      _nw180[(new BigNumber(21)).toNumber()] = p0;
      _nw180[(new BigNumber(22)).toNumber()] = p0;
      _nw180[(new BigNumber(23)).toNumber()] = (p0) || (false);
      _nw180[(new BigNumber(24)).toNumber()] = _834_v28;
      _nw180[(new BigNumber(25)).toNumber()] = _834_v28;
      _953_v121 = _nw180;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_953_v121).length))) {
        let _954_i6 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_954_i6)) && ((_954_i6).isLessThan(new BigNumber((_953_v121).length))))) {
          (_953_v121)[(_954_i6)] = ((_this).f5).isEqualTo(_this.f6);
        }
      }
      let _955_v122;
      _955_v122 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,_942_v110);
      let _956_v123;
      _956_v123 = _dafny.Set.fromElements(!(_834_v28));
      _955_v122 = (_955_v122).update((_dafny.ZERO).minus((_this).fm2(_this.f6, _956_v123, globalState)), _942_v110);
      return;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f6 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
      this._f7 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    set f6(value) {
      let _this = this;
      _this._f6 = value;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f7, f5, f6) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm2(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((((function () {
        let _coll51 = new _dafny.Map();
        for (const _compr_51 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f5,_dafny.Seq.UnicodeFromString("bbwcwhig"))).Keys.Elements) {
          let _957_v0 = _compr_51;
          if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f5,_dafny.Seq.UnicodeFromString("bbwcwhig"))).contains(_957_v0)) {
            _coll51.push([(_957_v0).minus((_this).f5),new _dafny.CodePoint('a'.codePointAt(0))]);
          }
        }
        return _coll51;
      }()).Merge(function () {
        let _coll52 = new _dafny.Map();
        for (const _compr_52 of _dafny.IntegerRange(new BigNumber(649), new BigNumber(978))) {
          let _958_v1 = _compr_52;
          if (((new BigNumber(649)).isLessThanOrEqualTo(_958_v1)) && ((_958_v1).isLessThan(new BigNumber(978)))) {
            _coll52.push([(_958_v1).minus((_this).f5),new _dafny.CodePoint('h'.codePointAt(0))]);
          }
        }
        return _coll52;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_this.f6)).cardinality()),new _dafny.CodePoint('x'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_this).f7)).cardinality()),new _dafny.CodePoint('u'.codePointAt(0)))))).length);
    };
    fm0(p0, p1, p2, globalState) {
      let _this = this;
      if (false) {
        return (_this).f7;
      } else {
        return (_this).f7;
      }
    };
    fm1(globalState) {
      let _this = this;
      return (_this).f7;
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _959_i0;
      _959_i0 = _dafny.ZERO;
      L4: {
        while ((_module.D0.create_DC1((_this).f7, (_this).f7)).dtor_cf2) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_959_i0)) {
              break L4;
            }
            _959_i0 = (_959_i0).plus(_dafny.ONE);
            let _960_v0;
            _960_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,false);
            r1 = (((_this).f7) ? ((_this).f7) : ((((_960_v0).contains((_this).f7)) ? ((_960_v0).get((_this).f7)) : ((_this).f7))));
            let _961_v1;
            _961_v1 = _module.D0.create_DC3((_this).f7);
            r1 = (_961_v1).dtor_cf3;
            let _962_v2;
            _962_v2 = _dafny.Set.fromElements(_this.f6);
            let _963_v3;
            let _nw181 = Array((new BigNumber(9)).toNumber()).fill(false);
            _963_v3 = _nw181;
            let _964_v4;
            _964_v4 = _dafny.Set.fromElements((_this).f7, (_this).f7, (_this).f7);
            let _index167 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_963_v3).length));
            (_963_v3)[_index167] = (_964_v4).IsSubsetOf(_964_v4);
            let _965_v5;
            _965_v5 = _module.D1.create_DC4(_dafny.Set.fromElements(_this.f6, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,new BigNumber(836))).length), _this.f6, (_this).f5, new BigNumber(337)));
            let _966_v6;
            _966_v6 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(853),(_this).f7));
            let _index168 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_963_v3).length));
            let _rhs175 = (_965_v5).dtor_cf4;
            let _rhs176 = (_this).f7;
            let _rhs177 = _dafny.Map.Empty.slice().updateUnsafe((_966_v6).IsDisjointFrom(_966_v6),(_this).f7);
            let _lhs110 = _963_v3;
            let _lhs111 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_963_v3).length));
            _962_v2 = _rhs175;
            _lhs110[_lhs111] = _rhs176;
            _960_v0 = _rhs177;
            let _967_v7;
            _967_v7 = _dafny.Seq.UnicodeFromString("eijplewq");
            let _968_v8;
            _968_v8 = _module.D1.create_DC6(_967_v7, (_this).f5);
            let _969_v9;
            _969_v9 = _dafny.Seq.of((_this).f5, (_this).f5, (_dafny.ZERO).minus(_this.f6), (_this).f5, (_this).fm2((_this).f5, _964_v4, globalState));
            let _970_v10;
            _970_v10 = _dafny.Seq.of((_dafny.ZERO).minus(_this.f6), (_968_v8).dtor_cf6, (_969_v9)[_module.__default.safeIndex(new BigNumber(168), new BigNumber((_969_v9).length))], new BigNumber(240), _this.f6);
            r0 = (_970_v10)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(877)), function (_971_i1) {
              return _dafny.Seq.UnicodeFromString("fnkhhatg");
            })).length)), new BigNumber((_970_v10).length))];
          }
        }
      }
      r1 = (_this).f7;
      let _972_v11;
      let _nw182 = Array((new BigNumber(4)).toNumber()).fill(false);
      _972_v11 = _nw182;
      let _973_v12;
      _973_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_module.D0.create_DC3((_this).f7));
      let _974_v13;
      _974_v13 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,_973_v12);
      let _975_v14;
      _975_v14 = new _dafny.CodePoint('q'.codePointAt(0));
      let _976_v15;
      _976_v15 = _dafny.Map.Empty.slice().updateUnsafe(_975_v14,new BigNumber(916));
      let _977_v16;
      _977_v16 = _dafny.Seq.of(_976_v15);
      let _978_v17;
      _978_v17 = _dafny.MultiSet.fromElements(_973_v12, (((_974_v13).contains((_dafny.ZERO).minus(new BigNumber((_977_v16).length)))) ? ((_974_v13).get((_dafny.ZERO).minus(new BigNumber((_977_v16).length)))) : (_973_v12)));
      let _index169 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length));
      (_972_v11)[_index169] = !((_978_v17).IsProperSubsetOf(_dafny.MultiSet.fromElements(_973_v12)));
      let _index170 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length));
      (_972_v11)[_index170] = (_this).f7;
      let _979_v18;
      _979_v18 = _dafny.MultiSet.fromElements((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))], (_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))]);
      let _980_i2;
      _980_i2 = _dafny.ZERO;
      L5: {
        while (!(new BigNumber((_979_v18).cardinality())).isEqualTo((_dafny.ZERO).minus(_this.f6))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_980_i2)) {
              break L5;
            }
            _980_i2 = (_980_i2).plus(_dafny.ONE);
            let _981_v19;
            let _nw183 = Array((new BigNumber(11)).toNumber()).fill([]);
            _981_v19 = _nw183;
            let _982_v20;
            let _nw184 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            _982_v20 = _nw184;
            let _index171 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_981_v19).length));
            (_981_v19)[_index171] = _982_v20;
            let _983_v21;
            _983_v21 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f5),_982_v20);
            let _index172 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_981_v19).length));
            (_981_v19)[_index172] = (((_983_v21).contains(_this.f6)) ? ((_983_v21).get(_this.f6)) : (_982_v20));
            let _984_v22;
            _984_v22 = _dafny.Seq.UnicodeFromString("wgosnrps");
            let _985_v23;
            _985_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(127)), ((_986_v14) => function (_987_i3) {
              return _986_v14;
            })(_975_v14)), _984_v22),(_981_v19)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_981_v19).length))]);
            let _988_v24;
            _988_v24 = _module.D1.create_DC6(_984_v22, _this.f6);
            let _index173 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_981_v19).length));
            (_981_v19)[_index173] = (((_985_v23).contains(_dafny.Seq.Concat((_988_v24).dtor_cf5, _984_v22))) ? ((_985_v23).get(_dafny.Seq.Concat((_988_v24).dtor_cf5, _984_v22))) : (_982_v20));
            _975_v14 = _975_v14;
            if ((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))]) {
              let _989_v25;
              _989_v25 = _dafny.Seq.of((_this).f5);
              let _990_v26;
              _990_v26 = _dafny.Seq.of(_dafny.Seq.of((_this).f5, _this.f6));
              let _991_v27;
              _991_v27 = _dafny.Seq.of((_this).f7);
              r1 = (((((_this).f7) ? ((_this).f7) : ((_this).f7))) ? (!(_dafny.Seq.IsProperPrefixOf(_989_v25, (_990_v26)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(new BigNumber((_991_v27).length))).length), new BigNumber((_990_v26).length))]))) : (false));
              let _992_v28;
              _992_v28 = _dafny.Set.fromElements(_972_v11, _972_v11, _972_v11);
              let _993_v29;
              let _nw185 = Array((new BigNumber(15)).toNumber());
              _nw185[(_dafny.ZERO).toNumber()] = _992_v28;
              _nw185[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_972_v11);
              _nw185[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_972_v11);
              _nw185[(new BigNumber(3)).toNumber()] = _992_v28;
              _nw185[(new BigNumber(4)).toNumber()] = (_992_v28).Union(_992_v28);
              _nw185[(new BigNumber(5)).toNumber()] = _992_v28;
              _nw185[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(_972_v11, _972_v11, _972_v11, _972_v11);
              _nw185[(new BigNumber(7)).toNumber()] = (_992_v28).Difference(_992_v28);
              _nw185[(new BigNumber(8)).toNumber()] = _992_v28;
              _nw185[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(_972_v11);
              _nw185[(new BigNumber(10)).toNumber()] = _992_v28;
              _nw185[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements(_972_v11);
              _nw185[(new BigNumber(12)).toNumber()] = _992_v28;
              _nw185[(new BigNumber(13)).toNumber()] = _992_v28;
              _nw185[(new BigNumber(14)).toNumber()] = _992_v28;
              _993_v29 = _nw185;
              let _index174 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_993_v29).length));
              (_993_v29)[_index174] = _992_v28;
              let _index175 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_993_v29).length));
              (_993_v29)[_index175] = _992_v28;
              let _994_v30;
              _994_v30 = _dafny.Set.fromElements((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))]);
              _994_v30 = ((((_this).f7) ? (_994_v30) : (_994_v30))).Intersect(_994_v30);
              let _995_v31;
              _995_v31 = _dafny.Map.Empty.slice().updateUnsafe(_972_v11,(_this).f5);
              let _996_v32;
              _996_v32 = _dafny.Map.Empty.slice().updateUnsafe((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))],(_995_v31).equals(_995_v31));
              _996_v32 = (_996_v32).Merge(_996_v32);
              let _997_v33;
              _997_v33 = _dafny.Seq.of(_984_v22, _dafny.Seq.UnicodeFromString("dajoybt"), _dafny.Seq.update(_984_v22, _module.__default.safeIndex(_this.f6, new BigNumber((_984_v22).length)), _975_v14), _984_v22);
              _984_v22 = _dafny.Seq.Concat(_dafny.Seq.Concat((_997_v33)[_module.__default.safeIndex(_this.f6, new BigNumber((_997_v33).length))], _dafny.Seq.UnicodeFromString("iugifadp")), _dafny.Seq.UnicodeFromString("viaetyyar"));
            } else {
              let _998_v34;
              _998_v34 = _dafny.Seq.of(_this.f6, _this.f6);
              let _999_v35;
              _999_v35 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f5).isEqualTo((_this).f5),(_998_v34)[_module.__default.safeIndex((_this).f5, new BigNumber((_998_v34).length))]);
              let _1000_v36;
              _1000_v36 = _999_v35;
              _999_v35 = (_999_v35).Merge((_1000_v36));
              let _1001_v37;
              _1001_v37 = _dafny.Seq.of(_979_v18, _979_v18);
              (globalState).f2 = (_module.__default.fm3((_this).f5, new BigNumber(434), _1001_v37, _998_v34, globalState)).dtor_cf6;
              r1 = (_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))];
              let _1002_v38;
              let _nw186 = new _module.C0();
              _nw186.__ctor(new BigNumber((_984_v22).length));
              _1002_v38 = _nw186;
              r0 = (_1002_v38).f8;
            }
          }
        }
      }
      let _index176 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length));
      (_972_v11)[_index176] = (_this).f7;
      if ((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))]) {
        let _1003_v39;
        let _nw187 = new _module.C0();
        _nw187.__ctor((_this).f5);
        _1003_v39 = _nw187;
        _1003_v39 = _1003_v39;
        let _nw188 = new _module.C0();
        _nw188.__ctor((_1003_v39).f8);
        _1003_v39 = _nw188;
        let _1004_v40;
        _1004_v40 = _dafny.Seq.of((_this).f5);
        let _1005_v41;
        _1005_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1003_v39,(_this).f7);
        let _1006_v42;
        _1006_v42 = _dafny.Map.Empty.slice().updateUnsafe((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))],_975_v14);
        let _1007_v43;
        let _nw189 = Array((new BigNumber(11)).toNumber());
        _nw189[(_dafny.ZERO).toNumber()] = ((_1003_v39).f8).minus(new BigNumber((_1004_v40).length));
        _nw189[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("sksxj")).length);
        _nw189[(new BigNumber(2)).toNumber()] = ((_1003_v39).f8).plus((_dafny.ZERO).minus((_1003_v39).f8));
        _nw189[(new BigNumber(3)).toNumber()] = (_this).f5;
        _nw189[(new BigNumber(4)).toNumber()] = (_1004_v40)[_module.__default.safeIndex(new BigNumber((_1005_v41).length), new BigNumber((_1004_v40).length))];
        _nw189[(new BigNumber(5)).toNumber()] = _this.f6;
        _nw189[(new BigNumber(6)).toNumber()] = _this.f6;
        _nw189[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_1003_v39).fm4(_975_v14, _module.__default.fm5(_1004_v40, (_this).f7, _this.f6, globalState), (_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))], globalState));
        _nw189[(new BigNumber(8)).toNumber()] = new BigNumber((_1006_v42).length);
        _nw189[(new BigNumber(9)).toNumber()] = ((_this).f5).minus(_this.f6);
        _nw189[(new BigNumber(10)).toNumber()] = (_1003_v39).f8;
        _1007_v43 = _nw189;
        let _1008_v44;
        _1008_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1007_v43,_this.f6);
        let _1009_v45;
        let _nw190 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1009_v45 = _nw190;
        let _1010_v46;
        _1010_v46 = _dafny.Seq.UnicodeFromString("gfn");
        let _1011_v47;
        _1011_v47 = _dafny.Map.Empty.slice().updateUnsafe((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))],_1007_v43);
        let _1012_v48;
        _1012_v48 = _dafny.MultiSet.fromElements(_972_v11);
        let _rhs178 = (((_1011_v47).contains((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))])) ? ((_1011_v47).get((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))])) : (_1007_v43));
        let _rhs179 = _1008_v44;
        let _rhs180 = _1009_v45;
        let _rhs181 = _1010_v46;
        let _rhs182 = (((_1012_v48).contains(_972_v11)) ? ((_1012_v48).get(_972_v11)) : ((_1003_v39).fm4(_975_v14, _module.__default.fm5(_dafny.Seq.Create(_module.__default.abs(new BigNumber(530)), ((_1013_v39) => function (_1014_i4) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_1013_v39).f8)).length);
        })(_1003_v39)), (_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))], (_this).f5, globalState), (_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))], globalState)));
        _1007_v43 = _rhs178;
        _1008_v44 = _rhs179;
        _1009_v45 = _rhs180;
        _1010_v46 = _rhs181;
        r0 = _rhs182;
        let _1015_v49;
        _1015_v49 = _dafny.Map.Empty.slice().updateUnsafe((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))],_this.f6);
        (globalState).f2 = _module.__default.safeDivisionInt((new BigNumber((_1015_v49).length)).minus((_1003_v39).f8), (_1003_v39).f8);
        let _index177 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length));
        (_972_v11)[_index177] = (_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))];
      } else {
        (globalState).f2 = (_this).f5;
        let _index178 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length));
        (_972_v11)[_index178] = true;
        r1 = ((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))]) && (false);
        let _1016_v50;
        let _nw191 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
        _1016_v50 = _nw191;
        let _1017_v51;
        _1017_v51 = _dafny.Seq.of(_1016_v50, _1016_v50, _1016_v50);
        _1016_v50 = (_1017_v51)[_module.__default.safeIndex(_this.f6, new BigNumber((_1017_v51).length))];
        let _1018_v52;
        let _init22 = ((_1019_v11) => function (_1020_i5) {
          return _dafny.Seq.Concat(_dafny.Seq.of((_this).f7), _dafny.Seq.of((_module.D0.create_DC1(!((_1019_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_1019_v11).length))]), (_1019_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_1019_v11).length))])).dtor_cf1));
        })(_972_v11);
        let _nw192 = Array((new BigNumber(7)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw192.length); _i0_22++) {
          _nw192[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _1018_v52 = _nw192;
        let _1021_v53;
        _1021_v53 = _dafny.Seq.of((_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))], (_972_v11)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_972_v11).length))]);
        let _index179 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_1018_v52).length));
        (_1018_v52)[_index179] = _1021_v53;
        let _index180 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_1018_v52).length));
        (_1018_v52)[_index180] = _1021_v53;
      }
      let _1022_v54;
      let _nw193 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _1022_v54 = _nw193;
      let _1023_v55;
      _1023_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_1022_v54);
      r0 = ((_this).f5).multipliedBy(_module.__default.safeDivisionInt(_this.f6, new BigNumber((_1023_v55).length)));
      r1 = (_this).f7;
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      if ((((_this).f7) ? ((true) === ((_this).f7)) : ((_this).f7))) {
        let _1024_v0;
        let _nw194 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _1024_v0 = _nw194;
        let _1025_v1;
        let _init23 = function (_1026_i0) {
          return _module.__default.safeModuloInt(_1026_i0, _this.f6);
        };
        let _nw195 = Array((new BigNumber(29)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw195.length); _i0_23++) {
          _nw195[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _1025_v1 = _nw195;
        let _1027_v2;
        _1027_v2 = _dafny.Seq.of(_1025_v1, _1025_v1);
        let _1028_v3;
        _1028_v3 = _dafny.Seq.of(_1025_v1, (_1027_v2)[_module.__default.safeIndex(p0, new BigNumber((_1027_v2).length))]);
        let _index181 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_1024_v0).length));
        (_1024_v0)[_index181] = _1027_v2;
        let _index182 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((_1024_v0).length));
        (_1024_v0)[_index182] = _1028_v3;
        let _1029_v4;
        _1029_v4 = _module.D0.create_DC1((_this).f7, (_this).f7);
        let _1030_v5;
        _1030_v5 = _dafny.MultiSet.fromElements(_1029_v4);
        let _1031_v6;
        _1031_v6 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1025_v1,_this.f6)).length));
        let _1032_v7;
        _1032_v7 = _dafny.Set.fromElements(_1030_v5, (_1030_v5).update(_1029_v4, _module.__default.abs(new BigNumber((_1031_v6).length))));
        if ((_1032_v7).IsSubsetOf(((false) ? (_1032_v7) : (_1032_v7)))) {
          let _1033_v8;
          _1033_v8 = new _dafny.CodePoint('f'.codePointAt(0));
          _1033_v8 = new _dafny.CodePoint('t'.codePointAt(0));
          let _1034_v9;
          _1034_v9 = _dafny.Seq.UnicodeFromString("jxjhrkcig");
          let _1035_v10;
          _1035_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,p0);
          let _1036_v11;
          _1036_v11 = _dafny.Set.fromElements((_this).f7);
          let _1037_v12;
          _1037_v12 = _module.D1.create_DC6(_1034_v9, (_module.D1.create_DC6(_module.__default.fm7(_1033_v8, (_this).f7, (_this).f7, (_this).f5, globalState), new BigNumber((_1036_v11).length))).dtor_cf6);
          let _1038_v13;
          _1038_v13 = _dafny.Seq.of((_this).f7);
          let _1039_v14;
          let _nw196 = Array((new BigNumber(21)).toNumber());
          _nw196[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-602)), ((_1040_v8) => function (_1041_i1) {
            return _1040_v8;
          })(_1033_v8)), _1034_v9), _module.__default.safeIndex(new BigNumber(806), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-602)), ((_1042_v8) => function (_1043_i1) {
            return _1042_v8;
          })(_1033_v8)), _1034_v9)).length)), _module.__default.fm6(new BigNumber((_1035_v10).length), new BigNumber((_dafny.Seq.UnicodeFromString("kjqkqeiu")).length), (_this).f7, (_this).f7, globalState));
          _nw196[(_dafny.ONE).toNumber()] = ((!((_this).f7)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), ((_1044_v8) => function (_1045_i2) {
            return _1044_v8;
          })(_1033_v8))) : (_1034_v9));
          _nw196[(new BigNumber(2)).toNumber()] = _1034_v9;
          _nw196[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("me"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("me")).length)), _1033_v8);
          _nw196[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_module.__default.fm7(_1033_v8, (_this).f7, (_this).f7, _this.f6, globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm7(_1033_v8, (_this).f7, (_this).f7, _this.f6, globalState)).length)), _1033_v8);
          _nw196[(new BigNumber(5)).toNumber()] = (_1037_v12).dtor_cf5;
          _nw196[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(430)), ((_1046_v8) => function (_1047_i3) {
            return _1046_v8;
          })(_1033_v8));
          _nw196[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-983)), ((_1048_v8) => function (_1049_i4) {
            return _1048_v8;
          })(_1033_v8));
          _nw196[(new BigNumber(8)).toNumber()] = _1034_v9;
          _nw196[(new BigNumber(9)).toNumber()] = _1034_v9;
          _nw196[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("hcns");
          _nw196[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("gmyra");
          _nw196[(new BigNumber(12)).toNumber()] = _1034_v9;
          _nw196[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("dyonfg");
          _nw196[(new BigNumber(14)).toNumber()] = _1034_v9;
          _nw196[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_1034_v9, _dafny.Seq.UnicodeFromString("npwl"));
          _nw196[(new BigNumber(16)).toNumber()] = _1034_v9;
          _nw196[(new BigNumber(17)).toNumber()] = _1034_v9;
          _nw196[(new BigNumber(18)).toNumber()] = _module.__default.fm7(new _dafny.CodePoint('i'.codePointAt(0)), (_1038_v13)[_module.__default.safeIndex(new BigNumber((_1034_v9).length), new BigNumber((_1038_v13).length))], (_this).f7, new BigNumber((_dafny.Seq.UnicodeFromString("biesqdee")).length), globalState);
          _nw196[(new BigNumber(19)).toNumber()] = _1034_v9;
          _nw196[(new BigNumber(20)).toNumber()] = _1034_v9;
          _1039_v14 = _nw196;
          let _1050_v15;
          _1050_v15 = _dafny.MultiSet.fromElements(_this.f6, _this.f6);
          let _index183 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1025_v1).length));
          (_1025_v1)[_index183] = new BigNumber(((_dafny.MultiSet.fromElements(p0)).Union(_1050_v15)).cardinality());
          let _index184 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1025_v1).length));
          let _rhs183 = _1039_v14;
          let _rhs184 = (_this).fm2(new BigNumber((_1031_v6).length), _1036_v11, globalState);
          let _lhs112 = _1025_v1;
          let _lhs113 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_1025_v1).length));
          _1039_v14 = _rhs183;
          _lhs112[_lhs113] = _rhs184;
          let _1051_v16;
          _1051_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1033_v8,_1033_v8);
          let _1052_v17;
          let _nw197 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1052_v17 = _nw197;
          let _1053_v18;
          _1053_v18 = _dafny.MultiSet.fromElements(_1052_v17, _1052_v17, _1052_v17);
          let _1054_v19;
          let _nw198 = new _module.C1();
          _nw198.__ctor(new BigNumber((_1038_v13).length), new BigNumber((_1053_v18).cardinality()));
          _1054_v19 = _nw198;
          let _1055_v20;
          _1055_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1054_v19,new _dafny.CodePoint('n'.codePointAt(0)));
          _1033_v8 = (((_1051_v16).contains((((_1055_v20).contains(_1054_v19)) ? ((_1055_v20).get(_1054_v19)) : (_1033_v8)))) ? ((_1051_v16).get((((_1055_v20).contains(_1054_v19)) ? ((_1055_v20).get(_1054_v19)) : (_1033_v8)))) : ((_1034_v9)[_module.__default.safeIndex((_this).f5, new BigNumber((_1034_v9).length))]));
          let _1056_v21;
          let _nw199 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
          _1056_v21 = _nw199;
          let _1057_v22;
          _1057_v22 = _module.D4.create_DC11(_1031_v6);
          let _1058_v23;
          _1058_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1056_v21,_1057_v22);
          _1058_v23 = (_1058_v23).update(_1056_v21, function (_pat_let15_0) {
            return function (_1059_dt__update__tmp_h0) {
              return function (_pat_let16_0) {
                return function (_1061_dt__update_hcf13_h0) {
                  return _module.D4.create_DC11(_1061_dt__update_hcf13_h0);
                }(_pat_let16_0);
              }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-764)), function (_1060_i5) {
                return new BigNumber(456);
              }));
            }(_pat_let15_0);
          }(_1057_v22));
          _1038_v13 = _1038_v13;
        } else {
          let _1062_v24;
          let _nw200 = new _module.C1();
          _nw200.__ctor((_this).f5, (_this).f5);
          _1062_v24 = _nw200;
          let _1063_v25;
          let _nw201 = Array((new BigNumber(8)).toNumber());
          _nw201[(_dafny.ZERO).toNumber()] = _1062_v24;
          _nw201[(_dafny.ONE).toNumber()] = _1062_v24;
          _nw201[(new BigNumber(2)).toNumber()] = _1062_v24;
          _nw201[(new BigNumber(3)).toNumber()] = _1062_v24;
          _nw201[(new BigNumber(4)).toNumber()] = _1062_v24;
          _nw201[(new BigNumber(5)).toNumber()] = _1062_v24;
          _nw201[(new BigNumber(6)).toNumber()] = _1062_v24;
          _nw201[(new BigNumber(7)).toNumber()] = _1062_v24;
          _1063_v25 = _nw201;
          let _index185 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1063_v25).length));
          (_1063_v25)[_index185] = _1062_v24;
          let _index186 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1063_v25).length));
          (_1063_v25)[_index186] = _1062_v24;
          let _1064_v26;
          let _nw202 = new _module.C0();
          _nw202.__ctor(_this.f6);
          _1064_v26 = _nw202;
          let _1065_v27;
          let _1066_v28;
          let _out56;
          let _out57;
          let _outcollector16 = (_1062_v24).m1(globalState);
          _out56 = _outcollector16[0];
          _out57 = _outcollector16[1];
          _1065_v27 = _out56;
          _1066_v28 = _out57;
          let _1067_v29;
          _1067_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1027_v2,(_1064_v26).f8);
          (_1062_v24).f6 = (((_1067_v29).contains(_dafny.Seq.of(_1025_v1, _1025_v1))) ? ((_1067_v29).get(_dafny.Seq.of(_1025_v1, _1025_v1))) : ((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(351)), function (_1068_i6) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          })).length)).minus(new BigNumber(40))));
          let _1069_v30;
          _1069_v30 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1070_v31;
          _1070_v31 = _dafny.Seq.UnicodeFromString("yern");
          _1066_v28 = _dafny.Seq.contains(_1070_v31, _1069_v30);
        }
        _1029_v4 = (((_this).f7) ? (_1029_v4) : (function (_pat_let17_0) {
          return function (_1071_dt__update__tmp_h1) {
            return function (_pat_let18_0) {
              return function (_1072_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_1072_dt__update_hcf1_h0, (_1071_dt__update__tmp_h1).dtor_cf2);
              }(_pat_let18_0);
            }((_this).f7);
          }(_pat_let17_0);
        }(_1029_v4)));
        let _1073_v32;
        _1073_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f7);
        let _1074_v33;
        _1074_v33 = _module.D3.create_DC8(_1073_v32);
        let _pat_let_tv14 = _1073_v32;
        let _source13 = function (_pat_let19_0) {
          return function (_1075_dt__update__tmp_h2) {
            return function (_pat_let20_0) {
              return function (_1076_dt__update_hcf8_h0) {
                return _module.D3.create_DC8(_1076_dt__update_hcf8_h0);
              }(_pat_let20_0);
            }(_pat_let_tv14);
          }(_pat_let19_0);
        }(_1074_v33);
        if (_source13.is_DC9) {
          let _1077___mcc_h0 = (_source13).cf9;
          let _1078___mcc_h1 = (_source13).cf10;
          let _1079___mcc_h2 = (_source13).cf11;
          let _1080_cf11 = _1079___mcc_h2;
          let _1081_cf10 = _1078___mcc_h1;
          let _1082_cf9 = _1077___mcc_h0;
          let _1083_v34;
          let _nw203 = Array((new BigNumber(12)).toNumber());
          _1083_v34 = _nw203;
          let _1084_v35;
          let _nw204 = new _module.C2();
          _nw204.__ctor((_this).f5, _this.f6);
          _1084_v35 = _nw204;
          let _index187 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_1083_v34).length));
          (_1083_v34)[_index187] = _1084_v35;
          let _index188 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_1083_v34).length));
          (_1083_v34)[_index188] = _1084_v35;
          let _1085_v36;
          _1085_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,p0);
          let _1086_v37;
          _1086_v37 = _1085_v36;
          _1086_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_this.f6);
          let _1087_v38;
          let _nw205 = new _module.C2();
          _nw205.__ctor(((_this).f5).plus(_this.f6), _this.f6);
          _1087_v38 = _nw205;
          let _1088_v39;
          _1088_v39 = _dafny.Seq.of((_module.D9.create_DC28((_this).f7, (_this).f5, (_this).f5)).dtor_cf46);
          _1088_v39 = _dafny.Seq.of((_this).f7, (_this).f7, (_this).f7);
        } else if (_source13.is_DC8) {
          let _1089___mcc_h3 = (_source13).cf8;
          let _1090_cf8 = _1089___mcc_h3;
          let _1091_v40;
          _1091_v40 = false;
          let _1092_v41;
          _1092_v41 = _dafny.Set.fromElements(_1073_v32, _1073_v32, _1073_v32, _1090_cf8);
          _1091_v40 = !(((_1092_v41).Union(_1092_v41)).IsDisjointFrom(_1092_v41));
          let _1093_v42;
          _1093_v42 = _dafny.Seq.of(true);
          let _index189 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1025_v1).length));
          (_1025_v1)[_index189] = _module.__default.safeDivisionInt((_this).f5, new BigNumber((_1093_v42).length));
          let _1094_v43;
          _1094_v43 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),_1025_v1);
          let _1095_v44;
          _1095_v44 = _dafny.Seq.UnicodeFromString("nxsi");
          let _1096_v45;
          _1096_v45 = _dafny.MultiSet.fromElements(_1095_v44);
          let _1097_v46;
          _1097_v46 = new _dafny.CodePoint('d'.codePointAt(0));
          let _index190 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1025_v1).length));
          let _rhs185 = _this.f6;
          let _rhs186 = new BigNumber(-212);
          let _rhs187 = (((((_1096_v45).contains(_1095_v44)) ? ((_1096_v45).get(_1095_v44)) : (new BigNumber(291)))).minus(p0)).isLessThanOrEqualTo((new BigNumber((_1031_v6).length)).plus(_this.f6));
          let _rhs188 = _dafny.Map.Empty.slice().updateUnsafe(_1097_v46,_1025_v1);
          let _lhs114 = _1025_v1;
          let _lhs115 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1025_v1).length));
          let _lhs116 = _this;
          _lhs114[_lhs115] = _rhs185;
          _lhs116.f6 = _rhs186;
          _1091_v40 = _rhs187;
          _1094_v43 = _rhs188;
          let _1098_v47;
          let _init24 = ((_1099_v40) => function (_1100_i7) {
            return _module.D0.create_DC0(_1099_v40);
          })(_1091_v40);
          let _nw206 = Array((new BigNumber(10)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw206.length); _i0_24++) {
            _nw206[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1098_v47 = _nw206;
          let _1101_v48;
          _1101_v48 = _dafny.Seq.of(_1098_v47);
          let _1102_v49;
          _1102_v49 = _module.D1.create_DC5();
          let _1103_v50;
          let _nw207 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
          _1103_v50 = _nw207;
          let _rhs189 = _dafny.Seq.Concat(_dafny.Seq.of(_1098_v47), _1101_v48);
          let _rhs190 = _1102_v49;
          let _rhs191 = _1103_v50;
          _1101_v48 = _rhs189;
          _1102_v49 = _rhs190;
          _1103_v50 = _rhs191;
          (globalState).f2 = (_this).f5;
        } else {
          let _1104___mcc_h4 = (_source13).cf12;
          let _1105_cf12 = _1104___mcc_h4;
          let _1106_v51;
          _1106_v51 = _dafny.Seq.UnicodeFromString("b");
          let _1107_v52;
          _1107_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1031_v6,p0);
          let _1108_v53;
          _1108_v53 = _module.D1.create_DC6(_1106_v51, _this.f6);
          let _1109_v54;
          let _nw208 = new _module.C2();
          _nw208.__ctor(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(826)), function (_1110_i8) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })).length), new BigNumber((_module.__default.fm5(_dafny.Seq.of(p0, _this.f6, p0), (_this).f7, new BigNumber(((_1108_v53).dtor_cf5).length), globalState)).cardinality()));
          _1109_v54 = _nw208;
          let _1111_v55;
          let _nw209 = Array((new BigNumber(14)).toNumber());
          _nw209[(_dafny.ZERO).toNumber()] = (_this).f7;
          _nw209[(_dafny.ONE).toNumber()] = ((_this).f7) || ((_this).f7);
          _nw209[(new BigNumber(2)).toNumber()] = (_this).f7;
          _nw209[(new BigNumber(3)).toNumber()] = (_this).f7;
          _nw209[(new BigNumber(4)).toNumber()] = (_this).f7;
          _nw209[(new BigNumber(5)).toNumber()] = (_this).f7;
          _nw209[(new BigNumber(6)).toNumber()] = _dafny.Seq.IsPrefixOf(_1106_v51, _1106_v51);
          _nw209[(new BigNumber(7)).toNumber()] = (_this).f7;
          _nw209[(new BigNumber(8)).toNumber()] = (_this).f7;
          _nw209[(new BigNumber(9)).toNumber()] = !(((_this).f7) || ((_this).f7));
          _nw209[(new BigNumber(10)).toNumber()] = (p0).isLessThan(_this.f6);
          _nw209[(new BigNumber(11)).toNumber()] = (_1107_v52).contains(_module.__default.fm28(p0, globalState));
          _nw209[(new BigNumber(12)).toNumber()] = (_dafny.Set.fromElements(_1109_v54, _1109_v54, _1109_v54, _1109_v54, _1109_v54)).IsProperSubsetOf(_dafny.Set.fromElements(_1109_v54));
          _nw209[(new BigNumber(13)).toNumber()] = !(((_this).f7) || (false));
          _1111_v55 = _nw209;
          let _1112_v56;
          _1112_v56 = _dafny.Set.fromElements(new BigNumber((_1106_v51).length));
          let _index191 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_1111_v55).length));
          (_1111_v55)[_index191] = (_1112_v56).contains((_this).f5);
          let _index192 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_1111_v55).length));
          (_1111_v55)[_index192] = !((_this).f7);
          let _1113_v57;
          _1113_v57 = _dafny.Seq.of((_1111_v55)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_1111_v55).length))]);
          let _1114_v58;
          _1114_v58 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of((_1111_v55)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_1111_v55).length))]), _1113_v57)));
          (globalState).f2 = new BigNumber((_1114_v58).cardinality());
          let _1115_v59;
          _1115_v59 = _dafny.Map.Empty.slice().updateUnsafe((_1111_v55)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_1111_v55).length))],new BigNumber((_dafny.Seq.UnicodeFromString("ffv")).length));
          let _index193 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_1111_v55).length));
          (_1111_v55)[_index193] = (new BigNumber(-77)).isLessThanOrEqualTo(_module.__default.safeModuloInt((_this).fm2(new BigNumber((_1115_v59).length), _dafny.Set.fromElements((_this).f7, (_1111_v55)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_1111_v55).length))], (_this).f7), globalState), p0));
          let _index194 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_1111_v55).length));
          (_1111_v55)[_index194] = (_this).f7;
        }
        let _1116_v60;
        let _nw210 = new _module.C2();
        _nw210.__ctor((_1031_v6)[_module.__default.safeIndex(_this.f6, new BigNumber((_1031_v6).length))], (_this).f5);
        _1116_v60 = _nw210;
      } else {
        let _1117_v61;
        _1117_v61 = new _dafny.CodePoint('a'.codePointAt(0));
        _1117_v61 = _1117_v61;
        let _1118_v62;
        let _init25 = function (_1119_i9) {
          return _module.__default.safeModuloInt(_1119_i9, (_this).f5);
        };
        let _nw211 = Array((new BigNumber(10)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw211.length); _i0_25++) {
          _nw211[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1118_v62 = _nw211;
        let _1120_v63;
        _1120_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1117_v61,(_this).f7);
        let _1121_v64;
        _1121_v64 = _dafny.MultiSet.fromElements(_1120_v63, _1120_v63);
        let _1122_v65;
        _1122_v65 = _module.D12.create_DC33(_1121_v64);
        let _1123_v66;
        _1123_v66 = _dafny.Seq.of(_1120_v63);
        let _1124_v67;
        let _nw212 = Array((new BigNumber(11)).toNumber());
        _nw212[(_dafny.ZERO).toNumber()] = _1121_v64;
        _nw212[(_dafny.ONE).toNumber()] = _1121_v64;
        _nw212[(new BigNumber(2)).toNumber()] = _1121_v64;
        _nw212[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_module.__default.fm29(_this.f6, globalState));
        _nw212[(new BigNumber(4)).toNumber()] = _1121_v64;
        _nw212[(new BigNumber(5)).toNumber()] = (_1122_v65).dtor_cf55;
        _nw212[(new BigNumber(6)).toNumber()] = (_1121_v64).Difference(_dafny.MultiSet.FromArray(_1123_v66));
        _nw212[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements(_1120_v63);
        _nw212[(new BigNumber(8)).toNumber()] = (_1121_v64).Intersect(_1121_v64);
        _nw212[(new BigNumber(9)).toNumber()] = _1121_v64;
        _nw212[(new BigNumber(10)).toNumber()] = _1121_v64;
        _1124_v67 = _nw212;
        let _index195 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_1124_v67).length));
        (_1124_v67)[_index195] = _1121_v64;
        let _1125_v68;
        _1125_v68 = _dafny.Seq.UnicodeFromString("tsvv");
        let _index196 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_1124_v67).length));
        let _rhs192 = _1118_v62;
        let _rhs193 = (_1125_v68)[_module.__default.safeIndex(new BigNumber((_1125_v68).length), new BigNumber((_1125_v68).length))];
        let _rhs194 = p0;
        let _rhs195 = _dafny.MultiSet.fromElements(_1120_v63);
        let _lhs117 = _this;
        let _lhs118 = _1124_v67;
        let _lhs119 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_1124_v67).length));
        _1118_v62 = _rhs192;
        _1117_v61 = _rhs193;
        _lhs117.f6 = _rhs194;
        _lhs118[_lhs119] = _rhs195;
        _1125_v68 = _1125_v68;
        let _1126_v69;
        _1126_v69 = true;
        let _1127_v70;
        _1127_v70 = _dafny.Set.fromElements(_this.f6);
        let _rhs196 = _1117_v61;
        let _rhs197 = (_dafny.Set.fromElements(_this.f6)).IsProperSubsetOf((_1127_v70).Difference(_1127_v70));
        let _rhs198 = (((new BigNumber((_1125_v68).length)).isLessThanOrEqualTo((_this).f5)) ? (_1117_v61) : (_1117_v61));
        _1117_v61 = _rhs196;
        _1126_v69 = _rhs197;
        _1117_v61 = _rhs198;
        _1117_v61 = (((_this).f7) ? (_1117_v61) : (_1117_v61));
      }
      let _1128_i10;
      _1128_i10 = _dafny.ZERO;
      L6: {
        while (!((_this).f7) || (((_this).f7) && ((_this).f7))) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1128_i10)) {
              break L6;
            }
            _1128_i10 = (_1128_i10).plus(_dafny.ONE);
            let _1129_v71;
            let _nw213 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Set.Empty);
            _1129_v71 = _nw213;
            let _1130_v72;
            _1130_v72 = _dafny.Set.fromElements((_this).f7);
            let _index197 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_1129_v71).length));
            (_1129_v71)[_index197] = _1130_v72;
            let _index198 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_1129_v71).length));
            (_1129_v71)[_index198] = _1130_v72;
            let _1131_v73;
            _1131_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,p0);
            let _1132_v74;
            _1132_v74 = _dafny.Seq.of(new BigNumber(-750), ((((_1131_v73).contains((_this).f7)) ? ((_1131_v73).get((_this).f7)) : (new BigNumber(-885)))).multipliedBy(new BigNumber(537)));
            (globalState).f2 = (_1132_v74)[_module.__default.safeIndex((_this).f5, new BigNumber((_1132_v74).length))];
            let _1133_v75;
            let _nw214 = new _module.C1();
            _nw214.__ctor((_dafny.ZERO).minus((((_module.D12.create_DC35((_this).f7, _this.f6, new BigNumber(388))).dtor_cf57) ? (p0) : (_this.f6))), p0);
            _1133_v75 = _nw214;
            let _1134_v76;
            let _nw215 = new _module.C2();
            _nw215.__ctor(p0, _this.f6);
            _1134_v76 = _nw215;
            _1134_v76 = _1134_v76;
          }
        }
      }
      if (!(!((_this).f7))) {
        let _1135_v77;
        _1135_v77 = _dafny.Set.fromElements(!((_this).f7));
        (globalState).f2 = ((_this).f5).minus((new BigNumber((_1135_v77).length)).minus((_this).f5));
        let _1136_v78;
        let _nw216 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _1136_v78 = _nw216;
        let _index199 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1136_v78).length));
        (_1136_v78)[_index199] = _this.f6;
        let _1137_v79;
        _1137_v79 = _dafny.MultiSet.fromElements((_this).f7, (_this).f7, (_this).f7, (_this).f7, (_this).f7);
        let _index200 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1136_v78).length));
        (_1136_v78)[_index200] = (((_1137_v79).contains((_this).f7)) ? ((_1137_v79).get((_this).f7)) : (p0));
        let _1138_v80;
        let _nw217 = Array((new BigNumber(9)).toNumber()).fill(false);
        _1138_v80 = _nw217;
        let _1139_v81;
        _1139_v81 = _dafny.Seq.of((_this).f5);
        let _index201 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1138_v80).length));
        (_1138_v80)[_index201] = _dafny.areEqual(_1139_v81, _1139_v81);
        let _1140_v82;
        let _nw218 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1140_v82 = _nw218;
        let _1141_v83;
        _1141_v83 = _dafny.Seq.UnicodeFromString("thwvihdbx");
        let _1142_v84;
        _1142_v84 = new _dafny.CodePoint('m'.codePointAt(0));
        let _index202 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_1140_v82).length));
        (_1140_v82)[_index202] = _dafny.Seq.update(_dafny.Seq.Concat(_1141_v83, _dafny.Seq.UnicodeFromString("no")), _module.__default.safeIndex((_this).f5, new BigNumber((_dafny.Seq.Concat(_1141_v83, _dafny.Seq.UnicodeFromString("no"))).length)), _1142_v84);
        let _index203 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1138_v80).length));
        let _index204 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_1140_v82).length));
        let _rhs199 = (_module.__default.safeModuloInt((_this).f5, p0)).plus((_this).fm2(p0, _dafny.Set.fromElements((_this).f7, (_this).f7, (_this).f7), globalState));
        let _rhs200 = (_this).f7;
        let _rhs201 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("rsn"), _module.__default.safeIndex(_this.f6, new BigNumber((_dafny.Seq.UnicodeFromString("rsn")).length)), (_1141_v83)[_module.__default.safeIndex(p0, new BigNumber((_1141_v83).length))]), _dafny.Seq.Concat(_1141_v83, _1141_v83));
        let _lhs120 = globalState;
        let _lhs121 = _1138_v80;
        let _lhs122 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1138_v80).length));
        let _lhs123 = _1140_v82;
        let _lhs124 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_1140_v82).length));
        _lhs120.f2 = _rhs199;
        _lhs121[_lhs122] = _rhs200;
        _lhs123[_lhs124] = _rhs201;
        let _index205 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1138_v80).length));
        (_1138_v80)[_index205] = false;
        let _1143_v85;
        _1143_v85 = _dafny.Map.Empty.slice().updateUnsafe((_1136_v78)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_1136_v78).length))],new BigNumber(525));
        let _1144_v86;
        _1144_v86 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_this.f6, (_dafny.ZERO).minus((_1136_v78)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_1136_v78).length))])),((_this).fm2(new BigNumber(301), _dafny.Set.fromElements((_this).f7, (_1138_v80)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_1138_v80).length))]), globalState)).isEqualTo(new BigNumber((_1143_v85).length)));
        _1144_v86 = (_1144_v86).update((_this).f5, _dafny.areEqual(_dafny.Seq.of((_dafny.ZERO).minus((_1136_v78)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_1136_v78).length))])), _1139_v81));
      } else {
        let _1145_v87;
        _1145_v87 = _dafny.Seq.of((_this).f7, !((_this).f7), (_this).f7, (_this).f7);
        let _1146_v88;
        _1146_v88 = _dafny.Set.fromElements(false);
        let _1147_v89;
        let _nw219 = Array((new BigNumber(13)).toNumber());
        _nw219[(_dafny.ZERO).toNumber()] = (_this).f7;
        _nw219[(_dafny.ONE).toNumber()] = (_this).f7;
        _nw219[(new BigNumber(2)).toNumber()] = (_this).f7;
        _nw219[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_1145_v87, _1145_v87);
        _nw219[(new BigNumber(4)).toNumber()] = (_1146_v88).IsDisjointFrom(_1146_v88);
        _nw219[(new BigNumber(5)).toNumber()] = (_this).f7;
        _nw219[(new BigNumber(6)).toNumber()] = (_this).f7;
        _nw219[(new BigNumber(7)).toNumber()] = !(_1146_v88).equals(_dafny.Set.fromElements((_this).f7, !((_this).f7), (_this).f7, (_this).f7, (_this).f7));
        _nw219[(new BigNumber(8)).toNumber()] = (_this).f7;
        _nw219[(new BigNumber(9)).toNumber()] = (_this).f7;
        _nw219[(new BigNumber(10)).toNumber()] = (((_this).f7) ? (!(false)) : ((_this).f7));
        _nw219[(new BigNumber(11)).toNumber()] = (_this).f7;
        _nw219[(new BigNumber(12)).toNumber()] = (_this).f7;
        _1147_v89 = _nw219;
        _1147_v89 = _1147_v89;
        if ((_this).f7) {
          let _1148_v90;
          _1148_v90 = false;
          _1148_v90 = ((_this).fm2(_this.f6, _1146_v88, globalState)).isLessThanOrEqualTo((_this).f5);
          let _index206 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1147_v89).length));
          (_1147_v89)[_index206] = (_this).f7;
          let _index207 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_1147_v89).length));
          (_1147_v89)[_index207] = (_this).fm1(globalState);
          let _1149_v91;
          let _nw220 = Array((new BigNumber(4)).toNumber());
          _nw220[(_dafny.ZERO).toNumber()] = ((_this).f5).plus(new BigNumber(-587));
          _nw220[(_dafny.ONE).toNumber()] = (_this).f5;
          _nw220[(new BigNumber(2)).toNumber()] = p0;
          _nw220[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(p0, new BigNumber(35));
          _1149_v91 = _nw220;
          let _index208 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1149_v91).length));
          (_1149_v91)[_index208] = new BigNumber(428);
          let _index209 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1149_v91).length));
          (_1149_v91)[_index209] = (((_1148_v90) ? (new BigNumber((_dafny.Seq.UnicodeFromString("qiamx")).length)) : (_this.f6))).plus((p0).multipliedBy(new BigNumber(140)));
          let _1150_v92;
          _1150_v92 = _module.D4.create_DC14(p0, _1148_v90, (_1149_v91)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_1149_v91).length))], true, _this.f6);
          let _1151_v93;
          _1151_v93 = _dafny.Map.Empty.slice().updateUnsafe((_1150_v92).dtor_cf22,_1149_v91);
          _1151_v93 = (_1151_v93).update((_this).f5, _1149_v91);
          let _1152_v94;
          _1152_v94 = _dafny.Seq.of(p0);
          let _1153_v96;
          _1153_v96 = _module.D1.create_DC4(function () {
  let _coll53 = new _dafny.Set();
  for (const _compr_53 of (_1152_v94).Elements) {
    let _1154_v95 = _compr_53;
    if (_dafny.Seq.contains(_1152_v94, _1154_v95)) {
      _coll53.add(_module.__default.safeDivisionInt(_1154_v95, new BigNumber((_dafny.Seq.of(!(false), true)).length)));
    }
  }
  return _coll53;
}());
          let _1155_v97;
          _1155_v97 = _dafny.Set.fromElements((_this).f5, p0, (_1149_v91)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_1149_v91).length))], p0);
          let _pat_let_tv15 = _1155_v97;
          _1153_v96 = function (_pat_let21_0) {
            return function (_1156_dt__update__tmp_h4) {
              return function (_pat_let22_0) {
                return function (_1157_dt__update_hcf4_h0) {
                  return _module.D1.create_DC4(_1157_dt__update_hcf4_h0);
                }(_pat_let22_0);
              }(_pat_let_tv15);
            }(_pat_let21_0);
          }(_1153_v96);
        } else {
          let _1158_v98;
          _1158_v98 = false;
          let _1159_v99;
          _1159_v99 = _dafny.Seq.UnicodeFromString("ympy");
          let _1160_v100;
          let _nw221 = new _module.C2();
          _nw221.__ctor(p0, _this.f6);
          _1160_v100 = _nw221;
          let _1161_v101;
          _1161_v101 = _dafny.MultiSet.fromElements(_1160_v100);
          let _1162_v103;
          _1162_v103 = _dafny.Seq.of(_this.f6);
          let _1163_v104;
          _1163_v104 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f5);
          let _1164_v105;
          _1164_v105 = new _dafny.CodePoint('n'.codePointAt(0));
          let _1165_v106;
          _1165_v106 = _module.D5.create_DC19(new BigNumber((_1146_v88).length), _1164_v105, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,true)).length));
          let _rhs202 = !(_1161_v101).contains(_1160_v100);
          let _rhs203 = ((function () {
            let _coll54 = new _dafny.Map();
            for (const _compr_54 of (_1162_v103).Elements) {
              let _1166_v102 = _compr_54;
              if (_dafny.Seq.contains(_1162_v103, _1166_v102)) {
                _coll54.push([_module.__default.safeModuloInt(_1166_v102, p0),new BigNumber(825)]);
              }
            }
            return _coll54;
          }()).Merge(_1163_v104)).contains((_this).fm2((_this).f5, _dafny.Set.fromElements((_this).f7), globalState));
          let _rhs204 = _1158_v98;
          let _rhs205 = ((true) ? (_dafny.Seq.update(_1159_v99, _module.__default.safeIndex(p0, new BigNumber((_1159_v99).length)), (_1165_v106).dtor_cf34)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), ((_1167_v105) => function (_1168_i11) {
            return _1167_v105;
          })(_1164_v105))));
          _1158_v98 = _rhs202;
          _1158_v98 = _rhs203;
          _1158_v98 = _rhs204;
          _1159_v99 = _rhs205;
          _1158_v98 = (_this).f7;
          let _1169_v107;
          let _nw222 = new _module.C2();
          _nw222.__ctor(new BigNumber(-813), (_this).f5);
          _1169_v107 = _nw222;
          let _1170_v108;
          let _init26 = function (_1171_i12) {
            return _dafny.Seq.UnicodeFromString("pfkrct");
          };
          let _nw223 = Array((new BigNumber(29)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw223.length); _i0_26++) {
            _nw223[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1170_v108 = _nw223;
          _1170_v108 = (((_1158_v98) || (false)) ? (_1170_v108) : (_1170_v108));
          _1158_v98 = _1158_v98;
        }
        let _index210 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1147_v89).length));
        (_1147_v89)[_index210] = (((_this).f7) ? (true) : (!((_this).f7)));
        let _index211 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1147_v89).length));
        (_1147_v89)[_index211] = (_this).f7;
        let _1172_v109;
        _1172_v109 = _dafny.Map.Empty.slice().updateUnsafe((_1147_v89)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1147_v89).length))],(_1147_v89)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1147_v89).length))]);
        _1172_v109 = (_1172_v109).update((_1147_v89)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1147_v89).length))], (_this).f7);
        let _index212 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1147_v89).length));
        (_1147_v89)[_index212] = (_1147_v89)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_1147_v89).length))];
      }
      let _1173_v110;
      _1173_v110 = _dafny.MultiSet.fromElements((_this).f5);
      (_this).f6 = _module.__default.safeModuloInt((_this).f5, new BigNumber((_1173_v110).cardinality()));
      let _1174_v111;
      _1174_v111 = _dafny.Seq.UnicodeFromString("pfpqnji");
      let _1175_v112;
      _1175_v112 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1176_v113;
      _1176_v113 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_this).f7);
      let _1177_v114;
      _1177_v114 = _dafny.Seq.of(_1176_v113);
      let _hi5 = new BigNumber((_module.__default.fm7(_1175_v112, (_this).fm0(_dafny.MultiSet.FromArray(_1177_v114), (_this).f7, _dafny.Seq.UnicodeFromString("pcfqhi"), globalState), (_this).f7, (_this).f5, globalState)).length);
      for (let _1178_i13 = (p0).minus(new BigNumber((_1174_v111).length)); _1178_i13.isLessThan(_hi5); _1178_i13 = _1178_i13.plus(_dafny.ONE)) {
        let _1179_v115;
        _1179_v115 = _module.D4.create_DC12((_this).f7, _1178_i13, (_this).f7);
        if (!((_1173_v110).Intersect(_dafny.MultiSet.fromElements(_this.f6))).contains((_1179_v115).dtor_cf15)) {
          let _1180_v116;
          _1180_v116 = _dafny.Seq.of((_this).f5, (_this).f5);
          let _1181_v117;
          _1181_v117 = _dafny.Map.Empty.slice().updateUnsafe(_1174_v111,_1180_v116);
          _1181_v117 = _1181_v117;
          (globalState).f2 = _1178_i13;
          let _1182_v118;
          _1182_v118 = _dafny.Seq.of((_this).f7);
          let _1183_v119;
          _1183_v119 = _dafny.Seq.of(_dafny.Seq.update(_1182_v118, _module.__default.safeIndex(_this.f6, new BigNumber((_1182_v118).length)), (_this).f7));
          let _1184_v120;
          _1184_v120 = _module.D12.create_DC34(!((_this).f7));
          let _1185_v121;
          _1185_v121 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,function (_pat_let23_0) {
            return function (_1186_dt__update__tmp_h5) {
              return function (_pat_let24_0) {
                return function (_1187_dt__update_hcf56_h0) {
                  return _module.D12.create_DC34(_1187_dt__update_hcf56_h0);
                }(_pat_let24_0);
              }(!((_this).f7));
            }(_pat_let23_0);
          }(_1184_v120));
          let _1188_v122;
          _1188_v122 = _dafny.Map.Empty.slice().updateUnsafe((_1183_v119)[_module.__default.safeIndex(new BigNumber((_1185_v121).length), new BigNumber((_1183_v119).length))],_module.__default.safeDivisionInt(_1178_i13, _1178_i13));
          _1188_v122 = (_1188_v122).update(_dafny.Seq.update(_dafny.Seq.of(true), _module.__default.safeIndex(new BigNumber((_1174_v111).length), new BigNumber((_dafny.Seq.of(true)).length)), true), _module.__default.safeModuloInt(new BigNumber((_1176_v113).length), _this.f6));
          let _1189_v123;
          _1189_v123 = true;
          let _1190_v124;
          _1190_v124 = _dafny.MultiSet.fromElements(_1175_v112, _1175_v112);
          _1189_v123 = (_1190_v124).IsSubsetOf(_1190_v124);
          _1189_v123 = _1189_v123;
        } else {
          let _1191_v125;
          let _nw224 = new _module.C1();
          _nw224.__ctor(_1178_i13, _this.f6);
          _1191_v125 = _nw224;
          let _1192_v126;
          _1192_v126 = true;
          let _1193_v127;
          _1193_v127 = _dafny.Seq.of((_this).f7);
          _1192_v126 = !((_1193_v127)[_module.__default.safeIndex(_1178_i13, new BigNumber((_1193_v127).length))]);
          let _1194_v128;
          let _init27 = ((_1195_v126) => function (_1196_i14) {
            return _1195_v126;
          })(_1192_v126);
          let _nw225 = Array((new BigNumber(4)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw225.length); _i0_27++) {
            _nw225[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1194_v128 = _nw225;
          let _1197_v129;
          _1197_v129 = _module.D7.create_DC22(_1194_v128, _1174_v111);
          let _1198_v130;
          _1198_v130 = _module.D7.create_DC25(_1197_v129);
          let _1199_v131;
          _1199_v131 = _dafny.Set.fromElements(_1198_v130);
          let _1200_v132;
          _1200_v132 = _dafny.Seq.of(_1199_v131, _1199_v131);
          let _1201_v133;
          let _nw226 = new _module.C2();
          _nw226.__ctor(new BigNumber((_1174_v111).length), new BigNumber(((_1199_v131).Union((_1200_v132)[_module.__default.safeIndex(new BigNumber(857), new BigNumber((_1200_v132).length))])).length));
          _1201_v133 = _nw226;
          let _1202_v134;
          let _nw227 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1202_v134 = _nw227;
          let _1203_v135;
          _1203_v135 = _dafny.Map.Empty.slice().updateUnsafe(_1202_v134,(_this).f5);
          _1203_v135 = (_1203_v135).update(_1202_v134, _1178_i13);
          (globalState).f2 = (_this).f5;
        }
        let _1204_v136;
        _1204_v136 = _module.D12.create_DC35((_this).f7, new BigNumber(90), (new BigNumber(222)).multipliedBy(_this.f6));
        _1204_v136 = _1204_v136;
        let _1205_v137;
        _1205_v137 = _dafny.Set.fromElements(true, true);
        let _1206_v138;
        _1206_v138 = _dafny.MultiSet.fromElements(_1205_v137, _1205_v137, _1205_v137);
        let _1207_v139;
        _1207_v139 = _dafny.Seq.of(_1205_v137, _1205_v137);
        if (((_1206_v138).Intersect(_dafny.MultiSet.FromArray(_1207_v139))).contains((_1205_v137).Difference(_1205_v137))) {
          _1205_v137 = _1205_v137;
          (globalState).f2 = _module.__default.safeDivisionInt((p0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_module.__default.fm28(_this.f6, globalState)).length))), (_this).f5);
          let _1208_v140;
          _1208_v140 = _dafny.Seq.of(false, (_1204_v136).dtor_cf57);
          let _1209_v141;
          _1209_v141 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f7, (_this).f7, false)).length), _1178_i13, (_this).fm2((_this).f5, _module.__default.fm21(new BigNumber((_dafny.MultiSet.FromArray(_1208_v140)).cardinality()), (_dafny.ZERO).minus((_this).f5), false, globalState), globalState), _this.f6);
          (_this).f6 = new BigNumber((_module.__default.fm5(_1209_v141, ((_1205_v137)).IsSubsetOf(_1205_v137), (_1178_i13).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(103)), ((_1210_v112) => function (_1211_i15) {
            return _1210_v112;
          })(_1175_v112))).length)), globalState)).cardinality());
          let _1212_v142;
          let _nw228 = Array((new BigNumber(5)).toNumber());
          _nw228[(_dafny.ZERO).toNumber()] = !((_this).f7);
          _nw228[(_dafny.ONE).toNumber()] = !(true);
          _nw228[(new BigNumber(2)).toNumber()] = (_this).f7;
          _nw228[(new BigNumber(3)).toNumber()] = !(_1178_i13).isEqualTo(new BigNumber((_1174_v111).length));
          _nw228[(new BigNumber(4)).toNumber()] = (_this).f7;
          _1212_v142 = _nw228;
          let _index213 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1212_v142).length));
          (_1212_v142)[_index213] = (_this).f7;
          let _index214 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1212_v142).length));
          let _rhs206 = _1174_v111;
          let _rhs207 = true;
          let _lhs125 = _1212_v142;
          let _lhs126 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1212_v142).length));
          _1174_v111 = _rhs206;
          _lhs125[_lhs126] = _rhs207;
          _1208_v140 = _dafny.Seq.of(false);
        } else {
          (globalState).f2 = _1178_i13;
          let _1213_v143;
          _1213_v143 = _module.D4.create_DC13((_this).f7, _1178_i13, (_this).f7);
          let _pat_let_tv16 = p0;
          let _1214_v144;
          let _nw229 = Array((new BigNumber(24)).toNumber());
          _nw229[(_dafny.ZERO).toNumber()] = _1213_v143;
          _nw229[(_dafny.ONE).toNumber()] = (((_this).f7) ? (function (_pat_let25_0) {
            return function (_1215_dt__update__tmp_h6) {
              return function (_pat_let26_0) {
                return function (_1216_dt__update_hcf17_h0) {
                  return _module.D4.create_DC13(_1216_dt__update_hcf17_h0, (_1215_dt__update__tmp_h6).dtor_cf18, (_1215_dt__update__tmp_h6).dtor_cf19);
                }(_pat_let26_0);
              }(false);
            }(_pat_let25_0);
          }(_1213_v143)) : (_module.__default.fm27(_this.f6, globalState)));
          _nw229[(new BigNumber(2)).toNumber()] = _module.__default.fm27(_this.f6, globalState);
          _nw229[(new BigNumber(3)).toNumber()] = function (_pat_let27_0) {
            return function (_1217_dt__update__tmp_h7) {
              return function (_pat_let28_0) {
                return function (_1218_dt__update_hcf18_h0) {
                  return _module.D4.create_DC13((_1217_dt__update__tmp_h7).dtor_cf17, _1218_dt__update_hcf18_h0, (_1217_dt__update__tmp_h7).dtor_cf19);
                }(_pat_let28_0);
              }(_pat_let_tv16);
            }(_pat_let27_0);
          }(_1213_v143);
          _nw229[(new BigNumber(4)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(5)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(6)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(7)).toNumber()] = _module.__default.fm27(_this.f6, globalState);
          _nw229[(new BigNumber(8)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(9)).toNumber()] = _module.D4.create_DC13((_this).f7, _this.f6, (_this).f7);
          _nw229[(new BigNumber(10)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(11)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(12)).toNumber()] = _module.D4.create_DC13((_this).f7, p0, (_this).f7);
          _nw229[(new BigNumber(13)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(14)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(15)).toNumber()] = _module.__default.fm27(_this.f6, globalState);
          _nw229[(new BigNumber(16)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(17)).toNumber()] = (((_this).f7) ? (_module.D4.create_DC13((_this).f7, _this.f6, !((_this).f7))) : (_module.D4.create_DC13((_this).f7, p0, (_this).f7)));
          _nw229[(new BigNumber(18)).toNumber()] = _module.D4.create_DC13((_this).fm1(globalState), _this.f6, (_this).f7);
          _nw229[(new BigNumber(19)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(20)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(21)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(22)).toNumber()] = _1213_v143;
          _nw229[(new BigNumber(23)).toNumber()] = _1213_v143;
          _1214_v144 = _nw229;
          let _index215 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1214_v144).length));
          (_1214_v144)[_index215] = (((_this).f7) ? (_1213_v143) : (_1213_v143));
          let _index216 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1214_v144).length));
          (_1214_v144)[_index216] = _1213_v143;
          let _1219_v145;
          let _nw230 = Array((new BigNumber(4)).toNumber());
          _nw230[(_dafny.ZERO).toNumber()] = _1178_i13;
          _nw230[(_dafny.ONE).toNumber()] = p0;
          _nw230[(new BigNumber(2)).toNumber()] = _this.f6;
          _nw230[(new BigNumber(3)).toNumber()] = p0;
          _1219_v145 = _nw230;
          let _1220_v146;
          _1220_v146 = _dafny.Map.Empty.slice().updateUnsafe(_1205_v137,_1178_i13);
          let _1221_v147;
          _1221_v147 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,new BigNumber(256));
          let _1222_v148;
          let _nw231 = new _module.C2();
          _nw231.__ctor(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-827)), function (_1223_i16) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          })).length), (_this).f5);
          _1222_v148 = _nw231;
          let _1224_v149;
          _1224_v149 = _dafny.Set.fromElements(_1222_v148, _1222_v148);
          let _nw232 = Array((new BigNumber(15)).toNumber());
          _nw232[(_dafny.ZERO).toNumber()] = p0;
          _nw232[(_dafny.ONE).toNumber()] = ((_this).f5).minus(_1178_i13);
          _nw232[(new BigNumber(2)).toNumber()] = _1178_i13;
          _nw232[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_1178_i13, new BigNumber(-158));
          _nw232[(new BigNumber(4)).toNumber()] = (((_1220_v146).contains(_1205_v137)) ? ((_1220_v146).get(_1205_v137)) : ((_this).f5));
          _nw232[(new BigNumber(5)).toNumber()] = new BigNumber(((_1221_v147).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f7,new BigNumber(627)))).length);
          _nw232[(new BigNumber(6)).toNumber()] = (_this).f5;
          _nw232[(new BigNumber(7)).toNumber()] = p0;
          _nw232[(new BigNumber(8)).toNumber()] = new BigNumber((_1224_v149).length);
          _nw232[(new BigNumber(9)).toNumber()] = _1222_v148.f6;
          _nw232[(new BigNumber(10)).toNumber()] = _1222_v148.f6;
          _nw232[(new BigNumber(11)).toNumber()] = (((_1221_v147).contains((_this).f7)) ? ((_1221_v147).get((_this).f7)) : ((_1179_v115).dtor_cf15));
          _nw232[(new BigNumber(12)).toNumber()] = (((_this).f7) ? (_this.f6) : (_1178_i13));
          _nw232[(new BigNumber(13)).toNumber()] = _1222_v148.f6;
          _nw232[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_1222_v148).fm2(p0, _1205_v137, globalState));
          _1219_v145 = _nw232;
          let _index217 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1219_v145).length));
          (_1219_v145)[_index217] = _this.f6;
          let _1225_v150;
          _1225_v150 = _dafny.MultiSet.fromElements((_this).f7);
          let _index218 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_1219_v145).length));
          (_1219_v145)[_index218] = (_this).fm2(_1222_v148.f6, _module.__default.fm21((_dafny.ZERO).minus(new BigNumber((_1174_v111).length)), new BigNumber((_1225_v150).cardinality()), true, globalState), globalState);
          let _1226_v151;
          let _nw233 = new _module.C2();
          _nw233.__ctor((_this).f5, _1178_i13);
          _1226_v151 = _nw233;
        }
        let _1227_v152;
        _1227_v152 = _dafny.Seq.of(_1178_i13);
        (_this).f6 = new BigNumber((_dafny.Seq.update(_1227_v152, _module.__default.safeIndex(p0, new BigNumber((_1227_v152).length)), _1178_i13)).length);
      }
      if ((_this.f6).isLessThanOrEqualTo(new BigNumber(918))) {
        let _1228_v153;
        _1228_v153 = false;
        _1228_v153 = (_this).f7;
        let _1229_v154;
        let _nw234 = new _module.C0();
        _nw234.__ctor(p0);
        _1229_v154 = _nw234;
        let _1230_v155;
        _1230_v155 = _dafny.Seq.of(_1229_v154);
        let _1231_v156;
        _1231_v156 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1174_v111).length),_1230_v155);
        let _1232_v157;
        _1232_v157 = _dafny.Set.fromElements(_1228_v153);
        (_this).f6 = (_this).fm2(new BigNumber((_1231_v156).length), _1232_v157, globalState);
        let _1233_v158;
        _1233_v158 = _module.D0.create_DC2();
        let _1234_v159;
        _1234_v159 = _module.D3.create_DC9(_1173_v110, _1233_v158, _this.f6);
        let _1235_v160;
        _1235_v160 = _module.D3.create_DC10(_1234_v159);
        _1235_v160 = _module.D3.create_DC10(_1234_v159);
        let _1236_v161;
        let _nw235 = Array((new BigNumber(6)).toNumber()).fill(false);
        _1236_v161 = _nw235;
        let _1237_v162;
        _1237_v162 = _dafny.Set.fromElements(_1236_v161);
        let _1238_v163;
        _1238_v163 = _dafny.Map.Empty.slice().updateUnsafe(_1175_v112,_1237_v162);
        let _1239_v164;
        _1239_v164 = _dafny.Map.Empty.slice().updateUnsafe(_1228_v153,_1236_v161);
        let _1240_v165;
        let _nw236 = Array((new BigNumber(15)).toNumber());
        _nw236[(_dafny.ZERO).toNumber()] = _1237_v162;
        _nw236[(_dafny.ONE).toNumber()] = _1237_v162;
        _nw236[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_1236_v161, _1236_v161, _1236_v161, _1236_v161, _1236_v161);
        _nw236[(new BigNumber(3)).toNumber()] = _1237_v162;
        _nw236[(new BigNumber(4)).toNumber()] = _1237_v162;
        _nw236[(new BigNumber(5)).toNumber()] = (_1237_v162).Union((((_1238_v163).contains(_1175_v112)) ? ((_1238_v163).get(_1175_v112)) : (_1237_v162)));
        _nw236[(new BigNumber(6)).toNumber()] = _1237_v162;
        _nw236[(new BigNumber(7)).toNumber()] = (_dafny.Set.fromElements((((_1239_v164).contains(!(_1228_v153))) ? ((_1239_v164).get(!(_1228_v153))) : (_1236_v161)))).Union(_1237_v162);
        _nw236[(new BigNumber(8)).toNumber()] = _1237_v162;
        _nw236[(new BigNumber(9)).toNumber()] = _1237_v162;
        _nw236[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_1236_v161, _1236_v161, _1236_v161);
        _nw236[(new BigNumber(11)).toNumber()] = _1237_v162;
        _nw236[(new BigNumber(12)).toNumber()] = _1237_v162;
        _nw236[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(_1236_v161);
        _nw236[(new BigNumber(14)).toNumber()] = (_1237_v162).Union(_1237_v162);
        _1240_v165 = _nw236;
        let _1241_v166;
        _1241_v166 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1174_v111).length),_1237_v162);
        let _1242_v167;
        _1242_v167 = _dafny.Seq.of((((_1241_v166).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-798)), ((_1245_v112) => function (_1246_i17) {
          return _1245_v112;
        })(_1175_v112))).length))) ? ((_1241_v166).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-798)), ((_1243_v112) => function (_1244_i17) {
          return _1243_v112;
        })(_1175_v112))).length))) : (_1237_v162)));
        let _index219 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1240_v165).length));
        (_1240_v165)[_index219] = (_1242_v167)[_module.__default.safeIndex((_this).f5, new BigNumber((_1242_v167).length))];
        let _index220 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1240_v165).length));
        (_1240_v165)[_index220] = _1237_v162;
        if (!(_1228_v153)) {
          let _1247_v168;
          _1247_v168 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("hkkcfpn")).length));
          let _1248_v169;
          _1248_v169 = _dafny.Set.fromElements(new BigNumber(-116), new BigNumber((_1247_v168).length), (_1229_v154).f8, (p0).minus((_this).f5));
          _1247_v168 = ((((_this).f7) ? (_1247_v168) : (_1248_v169))).Intersect(_1247_v168);
          let _1249_v170;
          let _nw237 = new _module.C1();
          _nw237.__ctor(_this.f6, ((!(_1228_v153)) ? (new BigNumber(615)) : ((_1229_v154).f8)));
          _1249_v170 = _nw237;
          let _1250_v171;
          _1250_v171 = _dafny.MultiSet.fromElements(_1228_v153, (_this).f7, _1228_v153, _1228_v153, (_this).f7);
          let _1251_v172;
          let _nw238 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _1251_v172 = _nw238;
          let _1252_v173;
          _1252_v173 = _dafny.Map.Empty.slice().updateUnsafe(_1251_v172,_1228_v153);
          let _1253_v174;
          let _nw239 = new _module.C1();
          _nw239.__ctor(_module.__default.safeDivisionInt((_1229_v154).f8, (((_1250_v171).contains((_this).f7)) ? ((_1250_v171).get((_this).f7)) : ((_1229_v154).f8))), (new BigNumber((_1252_v173).length)).minus(_this.f6));
          _1253_v174 = _nw239;
          let _1254_v175;
          let _nw240 = Array((new BigNumber(2)).toNumber()).fill(_module.D4.Default());
          _1254_v175 = _nw240;
          let _1255_v176;
          _1255_v176 = _dafny.Seq.of(new BigNumber((_1174_v111).length), new BigNumber(345));
          let _1256_v177;
          _1256_v177 = _module.D4.create_DC11(_1255_v176);
          let _index221 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1254_v175).length));
          (_1254_v175)[_index221] = _1256_v177;
          let _1257_v178;
          _1257_v178 = _dafny.Map.Empty.slice().updateUnsafe(_1251_v172,_1253_v174);
          let _1258_v179;
          _1258_v179 = _dafny.Map.Empty.slice().updateUnsafe(_1229_v154,(_1229_v154).f8);
          let _1259_v180;
          _1259_v180 = _dafny.Map.Empty.slice().updateUnsafe(_1258_v179,(_this).f7);
          let _index222 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1254_v175).length));
          let _rhs208 = _1249_v170;
          let _rhs209 = _1253_v174;
          let _rhs210 = _1256_v177;
          let _rhs211 = (((_1259_v180).contains((_1258_v179).Merge(_1258_v179))) ? ((_1259_v180).get((_1258_v179).Merge(_1258_v179))) : (_1228_v153));
          let _rhs212 = _1257_v178;
          let _lhs127 = _1254_v175;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1254_v175).length));
          _1249_v170 = _rhs208;
          _1253_v174 = _rhs209;
          _lhs127[_lhs128] = _rhs210;
          _1228_v153 = _rhs211;
          _1257_v178 = _rhs212;
          let _1260_v181;
          _1260_v181 = _dafny.Seq.of(_1251_v172, _1251_v172, _1251_v172, _1251_v172, _1251_v172);
          _1228_v153 = ((_1228_v153) ? (!_dafny.areEqual(_1260_v181, _1260_v181)) : ((_1232_v157).IsProperSubsetOf(_module.__default.fm21((_1229_v154).f8, (_1229_v154).f8, false, globalState))));
          let _1261_v182;
          _1261_v182 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("nduue"), _1174_v111);
          let _1262_v183;
          let _nw241 = new _module.C1();
          _nw241.__ctor(_this.f6, (_1253_v174).fm13(new BigNumber((_1174_v111).length), _1261_v182, globalState));
          _1262_v183 = _nw241;
          let _1263_v184;
          let _nw242 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Set.Empty);
          _1263_v184 = _nw242;
          _1263_v184 = _1263_v184;
        } else {
          let _1264_v185;
          let _nw243 = Array((new BigNumber(13)).toNumber());
          _1264_v185 = _nw243;
          _1264_v185 = _1264_v185;
          let _1265_v186;
          _1265_v186 = _dafny.MultiSet.fromElements(_1228_v153);
          let _1266_v187;
          let _nw244 = new _module.C2();
          _nw244.__ctor(p0, new BigNumber(((_1265_v186).Difference(_dafny.MultiSet.fromElements(true))).cardinality()));
          _1266_v187 = _nw244;
          _1266_v187 = _1266_v187;
          let _1267_v188;
          _1267_v188 = _dafny.Seq.of((_this).f7);
          _1228_v153 = !(!((_1267_v188)[_module.__default.safeIndex((_1229_v154).f8, new BigNumber((_1267_v188).length))]));
          let _index223 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_1236_v161).length));
          (_1236_v161)[_index223] = (_this).f7;
          let _1268_v189;
          _1268_v189 = _dafny.Seq.of((_1229_v154).f8, (_1229_v154).f8);
          let _index224 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_1236_v161).length));
          (_1236_v161)[_index224] = !((new BigNumber((_1268_v189).length)).isLessThanOrEqualTo((_1229_v154).f8)) || (_1228_v153);
          let _index225 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_1236_v161).length));
          (_1236_v161)[_index225] = (_this).f7;
        }
      } else {
        if (!((_module.__default.fm30(globalState)).IsSubsetOf((_dafny.MultiSet.fromElements(_1174_v111, _1174_v111)).Intersect(_dafny.MultiSet.fromElements(_1174_v111))))) {
          let _1269_v190;
          let _nw245 = new _module.C0();
          _nw245.__ctor(new BigNumber(389));
          _1269_v190 = _nw245;
          let _1270_v191;
          _1270_v191 = _dafny.Seq.of((_1269_v190).f8);
          let _1271_v192;
          _1271_v192 = _dafny.MultiSet.fromElements((_this).f7);
          let _1272_v193;
          let _nw246 = new _module.C1();
          _nw246.__ctor((_1270_v191)[_module.__default.safeIndex((_1269_v190).fm4(_1175_v112, _1271_v192, (_this).f7, globalState), new BigNumber((_1270_v191).length))], (_this).f5);
          _1272_v193 = _nw246;
          let _1273_v194;
          _1273_v194 = _dafny.Map.Empty.slice().updateUnsafe(_1272_v193,(_this).f7);
          let _1274_v195;
          _1274_v195 = _module.D12.create_DC35((_this).f7, new BigNumber(((_1273_v194).Merge(_1273_v194)).length), new BigNumber((_1174_v111).length));
          _1274_v195 = _1274_v195;
          (_this).f6 = _module.__default.safeModuloInt(p0, (_this).f5);
          let _1275_v196;
          _1275_v196 = false;
          _1275_v196 = (_this).f7;
          let _1276_v197;
          let _1277_v198;
          let _1278_v199;
          let _1279_v200;
          let _out58;
          let _out59;
          let _out60;
          let _out61;
          let _outcollector17 = _module.__default.m0(new _dafny.CodePoint('r'.codePointAt(0)), globalState);
          _out58 = _outcollector17[0];
          _out59 = _outcollector17[1];
          _out60 = _outcollector17[2];
          _out61 = _outcollector17[3];
          _1276_v197 = _out58;
          _1277_v198 = _out59;
          _1278_v199 = _out60;
          _1279_v200 = _out61;
        } else {
          (globalState).f2 = new BigNumber(34);
          let _1280_v201;
          let _nw247 = Array((_dafny.ONE).toNumber()).fill(false);
          _1280_v201 = _nw247;
          let _index226 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1280_v201).length));
          (_1280_v201)[_index226] = (((_this).f7) ? ((_this).f7) : (true));
          let _1281_v202;
          _1281_v202 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), ((_1282_v112) => function (_1283_i18) {
            return _1282_v112;
          })(_1175_v112)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(773)), ((_1284_v112) => function (_1285_i19) {
            return _1284_v112;
          })(_1175_v112)));
          let _1286_v203;
          _1286_v203 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(((_1281_v202)[_module.__default.safeIndex((_this).f5, new BigNumber((_1281_v202).length))]).length));
          let _1287_v204;
          _1287_v204 = _dafny.Seq.of(new BigNumber((_1286_v203).length), p0);
          let _index227 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1280_v201).length));
          (_1280_v201)[_index227] = ((_1287_v204)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f7)).length), new BigNumber((_1287_v204).length))]).isLessThan((_this).f5);
          let _1288_v205;
          let _nw248 = new _module.C0();
          _nw248.__ctor(_this.f6);
          _1288_v205 = _nw248;
          let _1289_v206;
          let _nw249 = Array((new BigNumber(5)).toNumber());
          _nw249[(_dafny.ZERO).toNumber()] = _1288_v205;
          _nw249[(_dafny.ONE).toNumber()] = _1288_v205;
          _nw249[(new BigNumber(2)).toNumber()] = _1288_v205;
          _nw249[(new BigNumber(3)).toNumber()] = _1288_v205;
          _nw249[(new BigNumber(4)).toNumber()] = _1288_v205;
          _1289_v206 = _nw249;
          let _index228 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_1289_v206).length));
          (_1289_v206)[_index228] = (((_1280_v201)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_1280_v201).length))]) ? (_1288_v205) : (_1288_v205));
          let _index229 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_1289_v206).length));
          (_1289_v206)[_index229] = _1288_v205;
          _1286_v203 = (_1286_v203).update((_this).f5, _module.__default.safeModuloInt(p0, _this.f6));
          let _1290_v207;
          let _nw250 = new _module.C1();
          _nw250.__ctor((((_this).f7) ? (_this.f6) : ((_this).f5)), ((_1288_v205).f8).minus((_1288_v205).f8));
          _1290_v207 = _nw250;
          let _1291_v208;
          _1291_v208 = _dafny.Seq.of((_1280_v201)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_1280_v201).length))], (_this).f7, (_this).f7);
          let _1292_v209;
          let _nw251 = new _module.C2();
          _nw251.__ctor(new BigNumber((_1174_v111).length), new BigNumber((_1291_v208).length));
          _1292_v209 = _nw251;
          let _1293_v210;
          _1293_v210 = _dafny.Seq.of(_1292_v209);
          let _index230 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1280_v201).length));
          let _rhs213 = _1290_v207;
          let _rhs214 = !_dafny.Seq.contains(_1293_v210, _1292_v209);
          let _rhs215 = _1175_v112;
          let _rhs216 = _1287_v204;
          let _lhs129 = _1280_v201;
          let _lhs130 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1280_v201).length));
          _1290_v207 = _rhs213;
          _lhs129[_lhs130] = _rhs214;
          _1175_v112 = _rhs215;
          _1287_v204 = _rhs216;
        }
        let _1294_v211;
        _1294_v211 = _dafny.MultiSet.fromElements((_this).f7, (_this).f7);
        let _1295_v212;
        let _nw252 = new _module.C1();
        _nw252.__ctor(new BigNumber(((_dafny.MultiSet.fromElements((_this).f7)).Union(_1294_v211)).cardinality()), (((_this).f7) ? ((_dafny.ZERO).minus((_this).f5)) : (p0)));
        _1295_v212 = _nw252;
        let _1296_v213;
        let _nw253 = new _module.C2();
        _nw253.__ctor((_this).f5, _this.f6);
        _1296_v213 = _nw253;
        _1296_v213 = _1296_v213;
        let _1297_v214;
        _1297_v214 = true;
        _1297_v214 = (_this).fm1(globalState);
        _1296_v213 = _1296_v213;
      }
      return;
    }
    m3(globalState) {
      let _this = this;
      let r0 = undefined;
      let _1298_v0;
      _1298_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f7);
      _1298_v0 = (_1298_v0).update((_this).f7, (_this).f7);
      let _1299_v1;
      _1299_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,new BigNumber(683));
      (_this).f6 = (new BigNumber(-594)).multipliedBy(new BigNumber((_module.__default.fm10((_this).f5, _1299_v1, (_this).f7, _this.f6, globalState)).length));
      let _1300_v2;
      _1300_v2 = _dafny.MultiSet.fromElements((_this).f7);
      let _1301_v3;
      _1301_v3 = _dafny.Seq.of((_this).f7);
      let _1302_v4;
      _1302_v4 = _dafny.Set.fromElements((_this).f7, !((_this).f7), (_this).f7);
      let _1303_v5;
      _1303_v5 = _module.D9.create_DC28((_1300_v2).IsSubsetOf(_1300_v2), (_this).f5, (_this).fm2(new BigNumber((_1301_v3).length), _1302_v4, globalState));
      _1303_v5 = _module.D9.create_DC28((_this).f7, _this.f6, _this.f6);
      let _1304_v6;
      _1304_v6 = _dafny.Set.fromElements(new BigNumber((_1298_v0).length));
      let _1305_v7;
      _1305_v7 = _dafny.Seq.UnicodeFromString("exw");
      let _1306_v8;
      _1306_v8 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-957)), function (_1307_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }), _1305_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_1308_i1) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }), _1305_v7);
      let _1309_v9;
      _1309_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_this).f7);
      let _1310_v10;
      _1310_v10 = _dafny.Seq.of(_dafny.Set.fromElements((_dafny.ZERO).minus(_this.f6)), (_1304_v6).Union(_dafny.Set.fromElements(new BigNumber((_1306_v8).cardinality()))), (_dafny.Set.fromElements(_this.f6, (_this).f5, new BigNumber((_1309_v9).length))).Intersect(_1304_v6));
      let _1311_v11;
      _1311_v11 = true;
      let _1312_v12;
      _1312_v12 = _dafny.Set.fromElements(_module.__default.fm18(!((_this).f7), _this.f6, (_this).f5, _1311_v11, globalState));
      let _1313_v13;
      _1313_v13 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1314_v14;
      let _nw254 = new _module.C2();
      _nw254.__ctor(_this.f6, new BigNumber((_dafny.Seq.of(new BigNumber(43))).length));
      _1314_v14 = _nw254;
      let _1315_v15;
      _1315_v15 = _dafny.Seq.of(_1314_v14);
      let _1316_v16;
      _1316_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1313_v13,new BigNumber((_dafny.MultiSet.FromArray(_1315_v15)).cardinality()));
      let _rhs217 = (_this).f5;
      let _rhs218 = _1310_v10;
      let _rhs219 = ((_this).f5).isLessThan((((_1316_v16).contains(_1313_v13)) ? ((_1316_v16).get(_1313_v13)) : (new BigNumber(149))));
      let _rhs220 = _dafny.Set.fromElements(_1313_v13);
      let _rhs221 = ((_this).f7) || ((((_1298_v0).contains((_this).f7)) ? ((_1298_v0).get((_this).f7)) : ((_this).f7)));
      let _lhs131 = globalState;
      _lhs131.f2 = _rhs217;
      _1310_v10 = _rhs218;
      _1311_v11 = _rhs219;
      _1312_v12 = _rhs220;
      _1311_v11 = _rhs221;
      let _1317_v17;
      let _nw255 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _1317_v17 = _nw255;
      let _1318_v18;
      _1318_v18 = _module.D5.create_DC17(_this.f6, _1317_v17, (_module.__default.fm31(_this.f6, false, true, globalState)).update(_this.f6, (_this).f7));
      let _source14 = _1318_v18;
      if (_source14.is_DC17) {
        let _1319___mcc_h0 = (_source14).cf27;
        let _1320___mcc_h1 = (_source14).cf28;
        let _1321___mcc_h2 = (_source14).cf29;
        let _1322_cf29 = _1321___mcc_h2;
        let _1323_cf28 = _1320___mcc_h1;
        let _1324_cf27 = _1319___mcc_h0;
        (globalState).f2 = new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1311_v11,_1311_v11), _1298_v0)).length);
        let _1325_v19;
        _1325_v19 = _dafny.Seq.of((new BigNumber((_1305_v7).length)).plus(_1324_cf27));
        let _rhs222 = _1311_v11;
        let _rhs223 = (_1325_v19)[_module.__default.safeIndex((_dafny.ZERO).minus(_1324_cf27), new BigNumber((_1325_v19).length))];
        let _rhs224 = new BigNumber(928);
        let _rhs225 = !(new BigNumber(402)).isEqualTo(new BigNumber((_1301_v3).length));
        let _rhs226 = _dafny.Map.Empty.slice().updateUnsafe(((_1311_v11) ? (_1313_v13) : (_1313_v13)),((_this).f5).minus(_this.f6));
        let _lhs132 = _this;
        let _lhs133 = globalState;
        _1311_v11 = _rhs222;
        _lhs132.f6 = _rhs223;
        _lhs133.f2 = _rhs224;
        _1311_v11 = _rhs225;
        _1316_v16 = _rhs226;
        _1300_v2 = (_1300_v2).Intersect(_1300_v2);
        (globalState).f2 = (_this).f5;
      } else if (_source14.is_DC18) {
        let _1326___mcc_h3 = (_source14).cf30;
        let _1327___mcc_h4 = (_source14).cf31;
        let _1328___mcc_h5 = (_source14).cf32;
        let _1329_cf32 = _1328___mcc_h5;
        let _1330_cf31 = _1327___mcc_h4;
        let _1331_cf30 = _1326___mcc_h3;
        (globalState).f2 = (_this).f5;
        if ((_this.f6).isEqualTo(_module.__default.safeModuloInt(new BigNumber(135), _this.f6))) {
          (_this).f6 = (_this).f5;
          let _1332_v20;
          let _nw256 = new _module.C0();
          _nw256.__ctor((((_this).f7) ? (_this.f6) : (_this.f6)));
          _1332_v20 = _nw256;
          let _1333_v21;
          let _nw257 = new _module.C2();
          _nw257.__ctor(new BigNumber(868), new BigNumber((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(107)), ((_1334_v13) => function (_1335_i2) {
            return _1334_v13;
          })(_1313_v13)))).length));
          _1333_v21 = _nw257;
          let _1336_v22;
          _1336_v22 = _dafny.Seq.of(_1333_v21, _1333_v21, _1333_v21);
          _1333_v21 = (_1336_v22)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(495), (_1332_v20).f8)), new BigNumber((_1336_v22).length))];
          let _1337_v23;
          let _init28 = ((_1338_cf30) => function (_1339_i3) {
            return _1338_cf30;
          })(_1331_cf30);
          let _nw258 = Array((new BigNumber(22)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw258.length); _i0_28++) {
            _nw258[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1337_v23 = _nw258;
          let _index231 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1337_v23).length));
          (_1337_v23)[_index231] = _1329_cf32;
          let _index232 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1337_v23).length));
          (_1337_v23)[_index232] = (_1304_v6).IsSubsetOf(_1304_v6);
          let _1340_v24;
          _1340_v24 = _dafny.Map.Empty.slice().updateUnsafe(!((_1337_v23)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((_1337_v23).length))]),_module.__default.safeDivisionInt(_1330_cf31, (_this).fm2((_1332_v20).f8, _1302_v4, globalState)));
          let _1341_v25;
          let _nw259 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1341_v25 = _nw259;
          let _index233 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1341_v25).length));
          (_1341_v25)[_index233] = _dafny.Seq.UnicodeFromString("ucngrye");
          let _1342_v26;
          let _nw260 = Array((new BigNumber(20)).toNumber());
          _1342_v26 = _nw260;
          let _1343_v27;
          _1343_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,_1342_v26);
          let _1344_v28;
          let _nw261 = Array((new BigNumber(4)).toNumber());
          _nw261[(_dafny.ZERO).toNumber()] = (((_1343_v27).contains(_1330_cf31)) ? ((_1343_v27).get(_1330_cf31)) : (_1342_v26));
          _nw261[(_dafny.ONE).toNumber()] = _1342_v26;
          _nw261[(new BigNumber(2)).toNumber()] = _1342_v26;
          _nw261[(new BigNumber(3)).toNumber()] = _1342_v26;
          _1344_v28 = _nw261;
          let _index234 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1344_v28).length));
          (_1344_v28)[_index234] = _1342_v26;
          let _1345_v29;
          _1345_v29 = _module.D1.create_DC5();
          let _1346_v30;
          _1346_v30 = _module.D5.create_DC19((_1332_v20).f8, _1313_v13, (_this).f5);
          let _1347_v31;
          _1347_v31 = _dafny.MultiSet.fromElements(_this.f6, new BigNumber(148), _1330_cf31);
          let _index235 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1337_v23).length));
          let _index236 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1341_v25).length));
          let _index237 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1344_v28).length));
          let _rhs227 = (_1333_v21).fm9(_1345_v29, (new BigNumber((_1304_v6).length)).isLessThanOrEqualTo((_this).f5), globalState);
          let _rhs228 = (_1346_v30).dtor_cf34;
          let _rhs229 = _dafny.Map.Empty.slice().updateUnsafe(_1331_cf30,(new BigNumber(603)).plus(_this.f6));
          let _rhs230 = _module.__default.fm11(_1330_cf31, (((_1347_v31).contains((_this).f5)) ? ((_1347_v31).get((_this).f5)) : ((_1332_v20).f8)), _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_1298_v0).length)), (_1332_v20).fm4(_1313_v13, _1300_v2, (_this).f7, globalState)), _module.__default.safeModuloInt(_1330_cf31, _this.f6), globalState);
          let _rhs231 = _1342_v26;
          let _lhs134 = _1337_v23;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1337_v23).length));
          let _lhs136 = _1341_v25;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1341_v25).length));
          let _lhs138 = _1344_v28;
          let _lhs139 = _module.__default.safeIndex(new BigNumber(143), new BigNumber((_1344_v28).length));
          _lhs134[_lhs135] = _rhs227;
          _1313_v13 = _rhs228;
          _1340_v24 = _rhs229;
          _lhs136[_lhs137] = _rhs230;
          _lhs138[_lhs139] = _rhs231;
        } else {
          _1313_v13 = _1313_v13;
          let _1348_v32;
          let _nw262 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1348_v32 = _nw262;
          let _index238 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_1348_v32).length));
          (_1348_v32)[_index238] = _dafny.Seq.Concat(_1305_v7, _1305_v7);
          let _index239 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_1348_v32).length));
          (_1348_v32)[_index239] = _1305_v7;
          let _1349_v33;
          let _init29 = ((_1350_v7, _1351_v6, _1352_cf32, _1353_v2, _1354_v32) => function (_1355_i4) {
            return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D7.create_DC23((((_1353_v2).contains(_1352_cf32)) ? ((_1353_v2).get(_1352_cf32)) : ((_this).f5)), (_1354_v32)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_1354_v32).length))], (_this).f5)))).IsSubsetOf(_dafny.MultiSet.fromElements(_module.D7.create_DC23(_this.f6, _1350_v7, new BigNumber((_1351_v6).length))));
          })(_1305_v7, _1304_v6, _1329_cf32, _1300_v2, _1348_v32);
          let _nw263 = Array((new BigNumber(6)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw263.length); _i0_29++) {
            _nw263[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1349_v33 = _nw263;
          _1349_v33 = _1349_v33;
          let _1356_v34;
          _1356_v34 = _dafny.Seq.of((_module.D1.create_DC6(_dafny.Seq.UnicodeFromString("htofcuax"), _this.f6)).dtor_cf5, (_1348_v32)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_1348_v32).length))], (_1348_v32)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_1348_v32).length))]);
          let _1357_v35;
          let _nw264 = new _module.C2();
          _nw264.__ctor((_this).f5, new BigNumber((_1356_v34).length));
          _1357_v35 = _nw264;
          let _1358_v36;
          let _nw265 = new _module.C2();
          _nw265.__ctor(new BigNumber(-392), new BigNumber((_1299_v1).length));
          _1358_v36 = _nw265;
        }
        let _1359_v37;
        _1359_v37 = _dafny.Set.fromElements(_1300_v2);
        if ((_1359_v37).IsProperSubsetOf(_1359_v37)) {
          _1329_cf32 = _1311_v11;
          let _1360_v38;
          _1360_v38 = _module.D4.create_DC13(_1329_cf32, (_this).f5, _1329_cf32);
          let _rhs232 = true;
          let _rhs233 = _1331_cf30;
          let _rhs234 = !(((_this).f5).isLessThan((_1360_v38).dtor_cf18));
          _1311_v11 = _rhs232;
          _1311_v11 = _rhs233;
          _1331_cf30 = _rhs234;
          let _1361_v39;
          _1361_v39 = _module.D12.create_DC36(_1311_v11, _1331_cf30, _1305_v7);
          let _1362_v40;
          let _nw266 = new _module.C2();
          _nw266.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,_1330_cf31)).length), new BigNumber((_1309_v9).length));
          _1362_v40 = _nw266;
          _1309_v9 = (_1309_v9).update(_module.__default.safeDivisionInt(_this.f6, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1361_v39,_1362_v40)).length)), (_this).f7);
          let _1363_v41;
          let _nw267 = Array((new BigNumber(27)).toNumber()).fill(false);
          _1363_v41 = _nw267;
          let _1364_v42;
          _1364_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1362_v40,_1363_v41);
          _1364_v42 = _1364_v42;
          (globalState).f2 = (_this).f5;
        } else {
          let _1365_v43;
          _1365_v43 = _dafny.MultiSet.fromElements((_this).f5);
          let _1366_v44;
          let _nw268 = Array((new BigNumber(9)).toNumber());
          _nw268[(_dafny.ZERO).toNumber()] = _1365_v43;
          _nw268[(_dafny.ONE).toNumber()] = _1365_v43;
          _nw268[(new BigNumber(2)).toNumber()] = _1365_v43;
          _nw268[(new BigNumber(3)).toNumber()] = (_1365_v43).update(new BigNumber((_1365_v43).cardinality()), _module.__default.abs((_this).f5));
          _nw268[(new BigNumber(4)).toNumber()] = (_1365_v43).Intersect(_1365_v43);
          _nw268[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber(-390), _this.f6, (_this).f5);
          _nw268[(new BigNumber(6)).toNumber()] = _1365_v43;
          _nw268[(new BigNumber(7)).toNumber()] = (_1365_v43).Difference(_1365_v43);
          _nw268[(new BigNumber(8)).toNumber()] = (_1365_v43).update((_this).f5, _module.__default.abs((_this).f5));
          _1366_v44 = _nw268;
          _1366_v44 = _1366_v44;
          let _1367_v45;
          let _nw269 = new _module.C0();
          _nw269.__ctor((_module.D4.create_DC13(_1329_cf32, (_this).f5, _1329_cf32)).dtor_cf18);
          _1367_v45 = _nw269;
          let _1368_v46;
          _1368_v46 = _dafny.Seq.of(_this.f6, (_1367_v45).f8, _1330_cf31, (_this).f5, (_1367_v45).f8);
          _1368_v46 = _1368_v46;
          let _1369_v47;
          let _nw270 = new _module.C1();
          _nw270.__ctor((_1367_v45).fm4(_1313_v13, _1300_v2, _1329_cf32, globalState), new BigNumber(578));
          _1369_v47 = _nw270;
          _1369_v47 = _1369_v47;
          let _1370_v48;
          let _nw271 = Array((new BigNumber(18)).toNumber());
          _1370_v48 = _nw271;
          let _1371_v49;
          let _nw272 = new _module.C2();
          _nw272.__ctor(new BigNumber(326), new BigNumber((_1301_v3).length));
          _1371_v49 = _nw272;
          let _index240 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1370_v48).length));
          (_1370_v48)[_index240] = _1371_v49;
          let _index241 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1370_v48).length));
          let _rhs235 = _1371_v49;
          let _rhs236 = _this.f6;
          let _rhs237 = _1369_v47;
          let _lhs140 = _1370_v48;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_1370_v48).length));
          let _lhs142 = globalState;
          _lhs140[_lhs141] = _rhs235;
          _lhs142.f2 = _rhs236;
          _1369_v47 = _rhs237;
        }
        let _1372_v50;
        _1372_v50 = _module.D12.create_DC36(_1311_v11, _1329_cf32, _1305_v7);
        let _1373_v51;
        _1373_v51 = _dafny.Seq.of(_1305_v7, _1305_v7, _dafny.Seq.UnicodeFromString("ey"));
        let _1374_v52;
        let _nw273 = Array((new BigNumber(18)).toNumber());
        _nw273[(_dafny.ZERO).toNumber()] = _1305_v7;
        _nw273[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_module.__default.fm11((_this).f5, _this.f6, _1330_cf31, _this.f6, globalState), _1305_v7);
        _nw273[(new BigNumber(2)).toNumber()] = (_1372_v50).dtor_cf62;
        _nw273[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_1305_v7, _module.__default.safeIndex(_this.f6, new BigNumber((_1305_v7).length)), _1313_v13);
        _nw273[(new BigNumber(4)).toNumber()] = _1305_v7;
        _nw273[(new BigNumber(5)).toNumber()] = _1305_v7;
        _nw273[(new BigNumber(6)).toNumber()] = _1305_v7;
        _nw273[(new BigNumber(7)).toNumber()] = _1305_v7;
        _nw273[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1305_v7, _dafny.Seq.UnicodeFromString("iaddku"));
        _nw273[(new BigNumber(9)).toNumber()] = _1305_v7;
        _nw273[(new BigNumber(10)).toNumber()] = _1305_v7;
        _nw273[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("qpastef");
        _nw273[(new BigNumber(12)).toNumber()] = _1305_v7;
        _nw273[(new BigNumber(13)).toNumber()] = _1305_v7;
        _nw273[(new BigNumber(14)).toNumber()] = _1305_v7;
        _nw273[(new BigNumber(15)).toNumber()] = (_1373_v51)[_module.__default.safeIndex(_1330_cf31, new BigNumber((_1373_v51).length))];
        _nw273[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(5)), ((_1375_v13) => function (_1376_i5) {
          return _1375_v13;
        })(_1313_v13));
        _nw273[(new BigNumber(17)).toNumber()] = _1305_v7;
        _1374_v52 = _nw273;
        let _nw274 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1374_v52 = _nw274;
      } else if (_source14.is_DC19) {
        let _1377___mcc_h6 = (_source14).cf33;
        let _1378___mcc_h7 = (_source14).cf34;
        let _1379___mcc_h8 = (_source14).cf35;
        let _1380_cf35 = _1379___mcc_h8;
        let _1381_cf34 = _1378___mcc_h7;
        let _1382_cf33 = _1377___mcc_h6;
        _1382_cf33 = _1382_cf33;
        let _1383_v53;
        _1383_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1311_v11,(_1298_v0).Merge(_1298_v0));
        _1383_v53 = _module.__default.fm32((_this).f5, !(!(true)) || ((_this).f7), globalState);
        let _1384_v54;
        let _nw275 = new _module.C0();
        _nw275.__ctor(_1382_cf33);
        _1384_v54 = _nw275;
        _1384_v54 = _1384_v54;
        let _1385_v55;
        let _init30 = ((_1386_v0) => function (_1387_i6) {
          return _dafny.Seq.of((_module.D0.create_DC3((_module.D4.create_DC14(new BigNumber((_1386_v0).length), (_this).f7, new BigNumber(675), (_this).f7, _this.f6)).dtor_cf21)).dtor_cf3);
        })(_1298_v0);
        let _nw276 = Array((new BigNumber(2)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw276.length); _i0_30++) {
          _nw276[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1385_v55 = _nw276;
        let _index242 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_1385_v55).length));
        (_1385_v55)[_index242] = _dafny.Seq.of(_1311_v11);
        let _1388_v56;
        _1388_v56 = _dafny.Seq.of(_1301_v3, _dafny.Seq.of(true, (_this).f7), _1301_v3, _1301_v3, _1301_v3);
        let _1389_v57;
        _1389_v57 = _dafny.Seq.of((_1388_v56)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_1388_v56).length))], _1301_v3);
        let _1390_v58;
        let _nw277 = Array((new BigNumber(3)).toNumber()).fill(false);
        _1390_v58 = _nw277;
        let _1391_v59;
        _1391_v59 = _dafny.MultiSet.fromElements(_1390_v58);
        let _index243 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_1385_v55).length));
        (_1385_v55)[_index243] = _dafny.Seq.Concat(_dafny.Seq.Concat((_1388_v56)[_module.__default.safeIndex(new BigNumber((_1391_v59).cardinality()), new BigNumber((_1388_v56).length))], _1301_v3), _module.__default.fm33(globalState));
      } else {
        let _1392___mcc_h9 = (_source14).cf26;
        let _1393_cf26 = _1392___mcc_h9;
        _1313_v13 = _1313_v13;
        let _1394_v60;
        let _nw278 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
        _1394_v60 = _nw278;
        let _index244 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_1394_v60).length));
        (_1394_v60)[_index244] = _dafny.Seq.of((_this).f7);
        let _index245 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_1394_v60).length));
        (_1394_v60)[_index245] = _dafny.Seq.Concat(_1301_v3, _module.__default.fm33(globalState));
        let _1395_v61;
        _1395_v61 = _dafny.Map.Empty.slice().updateUnsafe(_this.f6,_1305_v7);
        _1311_v11 = !(_1395_v61).equals(_1395_v61);
        let _1396_v62;
        _1396_v62 = _dafny.Seq.of((_this).f5, new BigNumber(-235));
        let _1397_v63;
        _1397_v63 = _module.D4.create_DC11(_1396_v62);
        let _pat_let_tv17 = _1396_v62;
        let _1398_v64;
        let _nw279 = Array((_dafny.ONE).toNumber());
        _nw279[(_dafny.ZERO).toNumber()] = function (_pat_let29_0) {
          return function (_1399_dt__update__tmp_h0) {
            return function (_pat_let30_0) {
              return function (_1400_dt__update_hcf13_h0) {
                return _module.D4.create_DC11(_1400_dt__update_hcf13_h0);
              }(_pat_let30_0);
            }(_pat_let_tv17);
          }(_pat_let29_0);
        }(_1397_v63);
        _1398_v64 = _nw279;
        let _index246 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_1398_v64).length));
        (_1398_v64)[_index246] = _1397_v63;
        let _index247 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_1398_v64).length));
        (_1398_v64)[_index247] = _module.__default.fm27(_this.f6, globalState);
      }
      if (!((_1302_v4).IsDisjointFrom(_dafny.Set.fromElements(_1311_v11, _1311_v11, (_this).f7, !((_this).f7))))) {
        let _1401_v65;
        _1401_v65 = _dafny.MultiSet.fromElements(_1309_v9);
        let _1402_v66;
        _1402_v66 = _dafny.Seq.of(_1401_v65);
        if ((_this).fm0((_1402_v66)[_module.__default.safeIndex(_this.f6, new BigNumber((_1402_v66).length))], (_1300_v2).IsProperSubsetOf(_1300_v2), _dafny.Seq.Concat(_1305_v7, _1305_v7), globalState)) {
          let _1403_v67;
          _1403_v67 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(58)), function (_1404_i7) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }),_module.__default.safeModuloInt((_this).f5, (_this).f5));
          let _1405_v68;
          let _nw280 = Array((new BigNumber(11)).toNumber()).fill(false);
          _1405_v68 = _nw280;
          let _1406_v69;
          _1406_v69 = _dafny.MultiSet.fromElements(_1405_v68);
          let _1407_v70;
          _1407_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1314_v14,_1403_v67);
          let _rhs238 = (((_1407_v70).contains(_1314_v14)) ? ((_1407_v70).get(_1314_v14)) : ((_1403_v67).Merge(_1403_v67)));
          let _rhs239 = _1406_v69;
          _1403_v67 = _rhs238;
          _1406_v69 = _rhs239;
          let _index248 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_1405_v68).length));
          (_1405_v68)[_index248] = _1311_v11;
          let _index249 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_1405_v68).length));
          (_1405_v68)[_index249] = (_this.f6).isLessThanOrEqualTo(_this.f6);
          let _1408_v71;
          _1408_v71 = _module.D0.create_DC3((_this).f7);
          let _1409_v73;
          _1409_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1311_v11,_this.f6);
          let _1410_v74;
          _1410_v74 = _dafny.Seq.of((new BigNumber(687)).minus((_this).f5), _module.__default.safeDivisionInt(_this.f6, _this.f6), (_this.f6).plus(_this.f6), (_this).f5, (_this).fm2(_this.f6, _1302_v4, globalState));
          let _index250 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_1405_v68).length));
          let _index251 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_1405_v68).length));
          let _rhs240 = (_1408_v71).dtor_cf3;
          let _rhs241 = function () {
            let _coll55 = new _dafny.Map();
            for (const _compr_55 of _dafny.IntegerRange(new BigNumber(-393), new BigNumber(743))) {
              let _1411_v72 = _compr_55;
              if (((new BigNumber(-393)).isLessThanOrEqualTo(_1411_v72)) && ((_1411_v72).isLessThan(new BigNumber(743)))) {
                _coll55.push([_module.__default.safeModuloInt(_1411_v72, new BigNumber(513)),_this.f6]);
              }
            }
            return _coll55;
          }();
          let _rhs242 = ((_this.f6).minus((((_1409_v73).contains(_1311_v11)) ? ((_1409_v73).get(_1311_v11)) : ((_this).f5)))).multipliedBy((_dafny.ZERO).minus(_this.f6));
          let _rhs243 = (_1410_v74)[_module.__default.safeIndex((_this).f5, new BigNumber((_1410_v74).length))];
          let _rhs244 = (_1301_v3)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_1305_v7, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("esd"), _module.__default.safeIndex((_this).fm2(_this.f6, _1302_v4, globalState), new BigNumber((_dafny.Seq.UnicodeFromString("esd")).length)), _1313_v13))).length), new BigNumber((_1301_v3).length))];
          let _lhs143 = _1405_v68;
          let _lhs144 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_1405_v68).length));
          let _lhs145 = globalState;
          let _lhs146 = globalState;
          let _lhs147 = _1405_v68;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_1405_v68).length));
          _lhs143[_lhs144] = _rhs240;
          _1299_v1 = _rhs241;
          _lhs145.f2 = _rhs242;
          _lhs146.f2 = _rhs243;
          _lhs147[_lhs148] = _rhs244;
          (globalState).f2 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((((_1405_v68)[_module.__default.safeIndex(new BigNumber(950), new BigNumber((_1405_v68).length))]) ? ((_this).f5) : ((_dafny.ZERO).minus(_this.f6))), (_this).fm2((_this).f5, _1302_v4, globalState)));
          _1309_v9 = (_1309_v9).update((_this).f5, (_this).fm0(_1401_v65, (_this).f7, _1305_v7, globalState));
          _1305_v7 = _module.__default.fm7(_1313_v13, (_1405_v68)[_module.__default.safeIndex(new BigNumber(950), new BigNumber((_1405_v68).length))], _1311_v11, (_dafny.ZERO).minus(new BigNumber((_1410_v74).length)), globalState);
        } else {
          let _1412_v75;
          let _nw281 = Array((new BigNumber(13)).toNumber());
          _1412_v75 = _nw281;
          let _1413_v76;
          let _nw282 = Array((new BigNumber(2)).toNumber());
          _nw282[(_dafny.ZERO).toNumber()] = _1412_v75;
          _nw282[(_dafny.ONE).toNumber()] = _1412_v75;
          _1413_v76 = _nw282;
          let _1414_v77;
          let _nw283 = Array((_dafny.ONE).toNumber());
          _nw283[(_dafny.ZERO).toNumber()] = _1413_v76;
          _1414_v77 = _nw283;
          let _index252 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_1414_v77).length));
          (_1414_v77)[_index252] = _1413_v76;
          let _index253 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_1414_v77).length));
          let _nw284 = Array((new BigNumber(16)).toNumber()).fill([]);
          (_1414_v77)[_index253] = _nw284;
          (globalState).f2 = new BigNumber(-132);
          let _index254 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_1317_v17).length));
          (_1317_v17)[_index254] = (_this.f6).minus(_this.f6);
          let _index255 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_1317_v17).length));
          (_1317_v17)[_index255] = (_this).f5;
          let _1415_v78;
          let _nw285 = new _module.C0();
          _nw285.__ctor(((_this).f5).plus(new BigNumber((_1305_v7).length)));
          _1415_v78 = _nw285;
          _1313_v13 = _1313_v13;
        }
        if (_1311_v11) {
          let _1416_v79;
          _1416_v79 = _module.D10.create_DC30((_this).f5);
          (globalState).f2 = (_1416_v79).dtor_cf50;
          _1311_v11 = (_module.D12.create_DC34((_this).f7)).dtor_cf56;
          let _1417_v80;
          let _init31 = ((_1418_v11, _1419_v7) => function (_1420_i8) {
            return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1418_v11,_module.D7.create_DC23(_this.f6, _1419_v7, (_this).f5)),new _dafny.CodePoint('c'.codePointAt(0)));
          })(_1311_v11, _1305_v7);
          let _nw286 = Array((new BigNumber(14)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw286.length); _i0_31++) {
            _nw286[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1417_v80 = _nw286;
          let _1421_v81;
          _1421_v81 = _module.D12.create_DC36((((_1298_v0).contains(_1311_v11)) ? ((_1298_v0).get(_1311_v11)) : ((_this).f7)), _1311_v11, _1305_v7);
          let _1422_v82;
          _1422_v82 = _module.D7.create_DC23((_this).f5, _1305_v7, new BigNumber(845));
          let _1423_v83;
          _1423_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_1421_v81).dtor_cf60,_1422_v82),_1313_v13);
          let _index256 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1417_v80).length));
          (_1417_v80)[_index256] = (_1423_v83).Merge(_1423_v83);
          let _1424_v84;
          _1424_v84 = _dafny.Seq.of(_dafny.Seq.Concat(_1301_v3, _1301_v3), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_this).f7), _module.__default.safeIndex((_this).f5, new BigNumber((_dafny.Seq.of((_this).f7)).length)), !((_this).f7)), _1301_v3), _dafny.Seq.Concat(_1301_v3, _dafny.Seq.update(_1301_v3, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("bsnlo")).length), new BigNumber((_1301_v3).length)), false)));
          let _1425_v85;
          let _init32 = ((_1426_v13) => function (_1427_i9) {
            return (((_this).f7) ? (_1426_v13) : (new _dafny.CodePoint('u'.codePointAt(0))));
          })(_1313_v13);
          let _nw287 = Array((new BigNumber(15)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw287.length); _i0_32++) {
            _nw287[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1425_v85 = _nw287;
          let _index257 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_1425_v85).length));
          (_1425_v85)[_index257] = _1313_v13;
          let _index258 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1317_v17).length));
          (_1317_v17)[_index258] = (_this).f5;
          let _index259 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1417_v80).length));
          let _index260 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_1425_v85).length));
          let _index261 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1317_v17).length));
          let _rhs245 = _1313_v13;
          let _rhs246 = _1423_v83;
          let _rhs247 = _dafny.Seq.update(_1424_v84, _module.__default.safeIndex(new BigNumber((_1300_v2).cardinality()), new BigNumber((_1424_v84).length)), _1301_v3);
          let _rhs248 = _module.__default.fm18((_1311_v11) || (_1311_v11), (_this.f6).plus((_dafny.ZERO).minus(new BigNumber((_1304_v6).length))), (_this).f5, _1311_v11, globalState);
          let _rhs249 = _module.__default.safeModuloInt((_this.f6).plus(new BigNumber(580)), _module.__default.safeModuloInt(new BigNumber((_1300_v2).cardinality()), _this.f6));
          let _lhs149 = _1417_v80;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1417_v80).length));
          let _lhs151 = _1425_v85;
          let _lhs152 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_1425_v85).length));
          let _lhs153 = _1317_v17;
          let _lhs154 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1317_v17).length));
          _1313_v13 = _rhs245;
          _lhs149[_lhs150] = _rhs246;
          _1424_v84 = _rhs247;
          _lhs151[_lhs152] = _rhs248;
          _lhs153[_lhs154] = _rhs249;
          let _1428_v86;
          _1428_v86 = _dafny.Seq.of(_1300_v2);
          let _1429_v87;
          _1429_v87 = _module.D0.create_DC1(!(true), (_this).f7);
          let _1430_v88;
          let _nw288 = new _module.C0();
          _nw288.__ctor((_1317_v17)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1317_v17).length))]);
          _1430_v88 = _nw288;
          let _1431_v89;
          _1431_v89 = _dafny.Set.fromElements(_1430_v88);
          let _1432_v90;
          _1432_v90 = _dafny.Seq.of((_1430_v88).fm4(new _dafny.CodePoint('m'.codePointAt(0)), _1300_v2, (_this).f7, globalState), (_1317_v17)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1317_v17).length))]);
          let _1433_v91;
          let _nw289 = Array((new BigNumber(23)).toNumber());
          _nw289[(_dafny.ZERO).toNumber()] = (_this).f7;
          _nw289[(_dafny.ONE).toNumber()] = !(!(!((_this).f7)) || ((_this).f7));
          _nw289[(new BigNumber(2)).toNumber()] = _1311_v11;
          _nw289[(new BigNumber(3)).toNumber()] = !(_1311_v11);
          _nw289[(new BigNumber(4)).toNumber()] = _1311_v11;
          _nw289[(new BigNumber(5)).toNumber()] = _1311_v11;
          _nw289[(new BigNumber(6)).toNumber()] = (_this).f7;
          _nw289[(new BigNumber(7)).toNumber()] = _1311_v11;
          _nw289[(new BigNumber(8)).toNumber()] = _1311_v11;
          _nw289[(new BigNumber(9)).toNumber()] = _1311_v11;
          _nw289[(new BigNumber(10)).toNumber()] = ((_1428_v86)[_module.__default.safeIndex(new BigNumber(776), new BigNumber((_1428_v86).length))]).IsProperSubsetOf(_1300_v2);
          _nw289[(new BigNumber(11)).toNumber()] = (_this).f7;
          _nw289[(new BigNumber(12)).toNumber()] = (_1301_v3)[_module.__default.safeIndex((_this).f5, new BigNumber((_1301_v3).length))];
          _nw289[(new BigNumber(13)).toNumber()] = (_this).f7;
          _nw289[(new BigNumber(14)).toNumber()] = ((_this).f7) || ((((_1309_v9).contains((_this).f5)) ? ((_1309_v9).get((_this).f5)) : (_1311_v11)));
          _nw289[(new BigNumber(15)).toNumber()] = !(false) || (_1311_v11);
          _nw289[(new BigNumber(16)).toNumber()] = !(true);
          _nw289[(new BigNumber(17)).toNumber()] = (_1429_v87).dtor_cf1;
          _nw289[(new BigNumber(18)).toNumber()] = (_1431_v89).IsProperSubsetOf(_1431_v89);
          _nw289[(new BigNumber(19)).toNumber()] = ((_1310_v10)[_module.__default.safeIndex(new BigNumber(759), new BigNumber((_1310_v10).length))]).IsSubsetOf(_module.__default.fm34(_1432_v90, globalState));
          _nw289[(new BigNumber(20)).toNumber()] = !(!((_this).fm1(globalState)));
          _nw289[(new BigNumber(21)).toNumber()] = false;
          _nw289[(new BigNumber(22)).toNumber()] = (_this).f7;
          _1433_v91 = _nw289;
          _1433_v91 = _1433_v91;
          let _index262 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1317_v17).length));
          (_1317_v17)[_index262] = _module.__default.safeModuloInt(new BigNumber(29), (_1430_v88).f8);
        } else {
          (globalState).f2 = ((_this).fm2((_this).f5, _1302_v4, globalState)).multipliedBy(_this.f6);
          let _1434_v92;
          let _nw290 = Array((new BigNumber(7)).toNumber());
          _nw290[(_dafny.ZERO).toNumber()] = _1317_v17;
          _nw290[(_dafny.ONE).toNumber()] = _1317_v17;
          _nw290[(new BigNumber(2)).toNumber()] = _1317_v17;
          _nw290[(new BigNumber(3)).toNumber()] = _1317_v17;
          _nw290[(new BigNumber(4)).toNumber()] = _1317_v17;
          _nw290[(new BigNumber(5)).toNumber()] = _1317_v17;
          _nw290[(new BigNumber(6)).toNumber()] = _1317_v17;
          _1434_v92 = _nw290;
          let _index263 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1434_v92).length));
          (_1434_v92)[_index263] = _1317_v17;
          let _1435_v93;
          _1435_v93 = _module.D5.create_DC16(_1317_v17);
          let _index264 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1434_v92).length));
          (_1434_v92)[_index264] = (_1435_v93).dtor_cf26;
          _1311_v11 = ((_this).f7) && (!((_this).f5).isEqualTo((_this).f5));
          let _1436_v94;
          _1436_v94 = _dafny.MultiSet.fromElements(new BigNumber(207), (_this).f5);
          let _1437_v95;
          _1437_v95 = _module.D0.create_DC2();
          let _1438_v96;
          _1438_v96 = _module.D3.create_DC9(_1436_v94, _1437_v95, _this.f6);
          let _1439_v97;
          _1439_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1311_v11,_1438_v96);
          let _1440_v98;
          _1440_v98 = _module.D3.create_DC10((((_1439_v97).contains(_1311_v11)) ? ((_1439_v97).get(_1311_v11)) : (_1438_v96)));
          let _1441_v99;
          _1441_v99 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_1438_v96);
          let _1442_v100;
          _1442_v100 = _dafny.Seq.of((_this).f5);
          let _1443_v101;
          _1443_v101 = _dafny.Set.fromElements(_1440_v98, _module.D3.create_DC10((((_1441_v99).contains(new BigNumber((_1442_v100).length))) ? ((_1441_v99).get(new BigNumber((_1442_v100).length))) : (_1438_v96))));
          (_this).f6 = new BigNumber((_1443_v101).length);
          _1313_v13 = _1313_v13;
        }
        let _1444_v102;
        _1444_v102 = _module.D5.create_DC19(_this.f6, _1313_v13, _this.f6);
        _1444_v102 = _1444_v102;
        let _1445_v103;
        _1445_v103 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_this.f6);
        _1445_v103 = (_1445_v103).update(!(!(_1311_v11) || (_1311_v11)), (_this).f5);
        let _1446_v104;
        let _init33 = ((_1447_v11) => function (_1448_i10) {
          return ((_1447_v11) ? ((_this).f7) : (_1447_v11));
        })(_1311_v11);
        let _nw291 = Array((new BigNumber(18)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw291.length); _i0_33++) {
          _nw291[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1446_v104 = _nw291;
        let _index265 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1446_v104).length));
        (_1446_v104)[_index265] = (_this).f7;
        let _1449_v105;
        _1449_v105 = _dafny.Seq.of(_this.f6);
        let _1450_v106;
        let _nw292 = new _module.C1();
        _nw292.__ctor(new BigNumber((_1305_v7).length), (((_this).f7) ? ((_this).f5) : (new BigNumber((_1449_v105).length))));
        _1450_v106 = _nw292;
        let _index266 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1446_v104).length));
        let _rhs250 = _1311_v11;
        let _rhs251 = _1450_v106;
        let _rhs252 = _1311_v11;
        let _lhs155 = _1446_v104;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1446_v104).length));
        _lhs155[_lhs156] = _rhs250;
        _1450_v106 = _rhs251;
        _1311_v11 = _rhs252;
      } else {
        let _1451_v107;
        let _init34 = ((_1452_v11) => function (_1453_i11) {
          return _1452_v11;
        })(_1311_v11);
        let _nw293 = Array((new BigNumber(5)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw293.length); _i0_34++) {
          _nw293[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1451_v107 = _nw293;
        let _1454_v108;
        let _nw294 = Array((new BigNumber(7)).toNumber());
        _nw294[(_dafny.ZERO).toNumber()] = (((_this).f7) ? (_1451_v107) : (_1451_v107));
        _nw294[(_dafny.ONE).toNumber()] = _1451_v107;
        _nw294[(new BigNumber(2)).toNumber()] = _1451_v107;
        _nw294[(new BigNumber(3)).toNumber()] = _1451_v107;
        _nw294[(new BigNumber(4)).toNumber()] = _1451_v107;
        _nw294[(new BigNumber(5)).toNumber()] = _1451_v107;
        _nw294[(new BigNumber(6)).toNumber()] = _1451_v107;
        _1454_v108 = _nw294;
        let _index267 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_1454_v108).length));
        (_1454_v108)[_index267] = _1451_v107;
        let _index268 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_1454_v108).length));
        let _nw295 = Array((new BigNumber(11)).toNumber()).fill(false);
        (_1454_v108)[_index268] = _nw295;
        let _1455_v109;
        let _1456_v110;
        let _1457_v111;
        let _1458_v112;
        let _out62;
        let _out63;
        let _out64;
        let _out65;
        let _outcollector18 = _module.__default.m0(_1313_v13, globalState);
        _out62 = _outcollector18[0];
        _out63 = _outcollector18[1];
        _out64 = _outcollector18[2];
        _out65 = _outcollector18[3];
        _1455_v109 = _out62;
        _1456_v110 = _out63;
        _1457_v111 = _out64;
        _1458_v112 = _out65;
        let _1459_v113;
        let _1460_v114;
        let _1461_v115;
        let _1462_v116;
        let _out66;
        let _out67;
        let _out68;
        let _out69;
        let _outcollector19 = _module.__default.m0(new _dafny.CodePoint('y'.codePointAt(0)), globalState);
        _out66 = _outcollector19[0];
        _out67 = _outcollector19[1];
        _out68 = _outcollector19[2];
        _out69 = _outcollector19[3];
        _1459_v113 = _out66;
        _1460_v114 = _out67;
        _1461_v115 = _out68;
        _1462_v116 = _out69;
        (globalState).f2 = (_this).fm2(_1458_v112, _1302_v4, globalState);
        let _1463_v117;
        let _nw296 = new _module.C0();
        _nw296.__ctor((_this).f5);
        _1463_v117 = _nw296;
        let _1464_v118;
        _1464_v118 = _dafny.MultiSet.fromElements(_1463_v117);
        let _1465_v119;
        _1465_v119 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("vwwebb"), _1313_v13),_1464_v118);
        let _1466_v120;
        _1466_v120 = _dafny.MultiSet.fromElements(new BigNumber(326), (_this).f5);
        let _1467_v121;
        _1467_v121 = _dafny.Seq.of(_1466_v120);
        _1458_v112 = new BigNumber(((((_1465_v119).contains((_1466_v120).IsSubsetOf((_1467_v121)[_module.__default.safeIndex((_this).f5, new BigNumber((_1467_v121).length))]))) ? ((_1465_v119).get((_1466_v120).IsSubsetOf((_1467_v121)[_module.__default.safeIndex((_this).f5, new BigNumber((_1467_v121).length))]))) : (_1464_v118))).cardinality());
      }
      let _nw297 = new _module.C2();
      _nw297.__ctor((_this.f6).plus(new BigNumber(159)), (_dafny.ZERO).minus((_this).f5));
      r0 = _nw297;
      return r0;
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
