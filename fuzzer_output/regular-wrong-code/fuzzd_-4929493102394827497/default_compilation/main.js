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
      if ((false) && (false)) {
        return !(true) || (!(true));
      } else {
        return !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(402)), function (_0_i0) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("pdxgcrjxl"));
      }
    };
    static fm1(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(!(!(false)))).Intersect((_dafny.MultiSet.fromElements(!(false), true)).Intersect(_dafny.MultiSet.fromElements(true, false)));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return (((false) ? (new BigNumber((_dafny.Seq.of(new BigNumber(240), new BigNumber(-231), new BigNumber((_dafny.Seq.UnicodeFromString("rljhwsp")).length), new BigNumber(206))).length)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xe")).length),new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Seq.of(new BigNumber(223))).Elements) {
          let _1_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(223)), _1_v0)) {
            _coll0.add((_1_v0).plus(new BigNumber(741)));
          }
        }
        return _coll0;
      }()).length))).length))).length)))).plus((new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)))).length)).multipliedBy(new BigNumber(-387)));
    };
    static fm7(p0, globalState) {
      if (true) {
        return _dafny.Seq.UnicodeFromString("cs");
      } else {
        return _dafny.Seq.UnicodeFromString("huytibwmr");
      }
    };
    static fm10(globalState) {
      let _source0 = _module.D2.create_DC6(_module.D2.create_DC6(_module.D2.create_DC5(new BigNumber(-139), true, true, true)));
      if (_source0.is_DC4) {
        let _2___mcc_h0 = (_source0).cf4;
        let _3___mcc_h1 = (_source0).cf5;
        let _4_cf5 = _3___mcc_h1;
        let _5_cf4 = _2___mcc_h0;
        return new _dafny.CodePoint('t'.codePointAt(0));
      } else if (_source0.is_DC5) {
        let _6___mcc_h2 = (_source0).cf6;
        let _7___mcc_h3 = (_source0).cf7;
        let _8___mcc_h4 = (_source0).cf8;
        let _9___mcc_h5 = (_source0).cf9;
        let _10_cf9 = _9___mcc_h5;
        let _11_cf8 = _8___mcc_h4;
        let _12_cf7 = _7___mcc_h3;
        let _13_cf6 = _6___mcc_h2;
        return new _dafny.CodePoint('v'.codePointAt(0));
      } else if (_source0.is_DC3) {
        let _14___mcc_h6 = (_source0).cf3;
        let _15_cf3 = _14___mcc_h6;
        if (false) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('n'.codePointAt(0));
        }
      } else {
        let _16___mcc_h7 = (_source0).cf10;
        let _17_cf10 = _16___mcc_h7;
        return new _dafny.CodePoint('o'.codePointAt(0));
      }
    };
    static fm11(globalState) {
      return _module.D2.create_DC4((new BigNumber(855)).plus(new BigNumber(195)), (_dafny.Set.fromElements(new BigNumber(188))).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(-694), new BigNumber(297))));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of((_dafny.ZERO).minus((new BigNumber(354)).minus(new BigNumber(954))));
    };
    static fm13(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(!(false))))),new BigNumber(430))).Merge(((true) ? (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false, !(true)), _dafny.Seq.of(false), _dafny.Seq.of(true))).length))) : (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(false)).length)))));
    };
    static fm16(p0, globalState) {
      return new _dafny.CodePoint('a'.codePointAt(0));
    };
    static fm17(p0, p1, globalState) {
      return _dafny.Seq.of(!(!(true)), false, false, (_dafny.Set.fromElements(new BigNumber(360))).IsSubsetOf(_dafny.Set.fromElements(new BigNumber(35))));
    };
    static fm18(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ckp"), ((!(false)) ? (_dafny.Seq.UnicodeFromString("dy")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(360)), function (_18_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }))));
    };
    static fm19(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("kimkmjliq")).length)))).length),new _dafny.CodePoint('x'.codePointAt(0)))).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Seq.of(new BigNumber(-691), new BigNumber(373))).Elements) {
          let _19_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-691), new BigNumber(373)), _19_v0)) {
            _coll1.push([(_19_v0).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(310))).length)),new _dafny.CodePoint('x'.codePointAt(0))]);
          }
        }
        return _coll1;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Seq.of(_module.D3.create_DC9(new BigNumber((_dafny.Seq.of(false)).length), new _dafny.CodePoint('c'.codePointAt(0)), new BigNumber(512)))).Elements) {
          let _20_v1 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D3.create_DC9(new BigNumber((_dafny.Seq.of(false)).length), new _dafny.CodePoint('c'.codePointAt(0)), new BigNumber(512))), _20_v1)) {
            _coll2.push([_20_v1,false]);
          }
        }
        return _coll2;
      }()).length),new _dafny.CodePoint('d'.codePointAt(0))));
    };
    static fm20(p0, globalState) {
      return _module.D2.create_DC6(_module.D2.create_DC5(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)))).cardinality()), !(!(true)), !(false), false));
    };
    static fm21(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(640)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-152)), function (_21_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("moh")).length);
      })).length), new BigNumber(532)))).length),_module.__default.safeDivisionInt(new BigNumber(798), new BigNumber((_dafny.Seq.UnicodeFromString("hnvcppjf")).length)));
    };
    static fm22(p0, p1, p2, p3, globalState) {
      if (!(false)) {
        return _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)));
      } else {
        return _dafny.Set.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)));
      }
    };
    static fm23(p0, globalState) {
      return _module.D7.create_DC21(new BigNumber(-405), (_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false)));
    };
    static fm24(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(!(!(false)))).length))).Union(_dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("bdxocc")).length)))))).Union(_dafny.Set.fromElements(new BigNumber(449)));
    };
    static fm25(globalState) {
      if (false) {
        return _module.D3.create_DC7(_dafny.Map.Empty.slice().updateUnsafe((_module.D7.create_DC22(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(693),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).length),new BigNumber(42))).length), new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))).length), false)).dtor_cf40,true));
      } else {
        return _module.D3.create_DC7(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(319),true));
      }
    };
    static fm26(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())), new BigNumber((function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(385)), function (_22_i0) {
          return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-738)), function (_23_i1) {
            return new BigNumber((_dafny.Seq.of(true)).length);
          })).length);
        })).Elements) {
          let _24_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(385)), function (_22_i0) {
            return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-738)), function (_23_i1) {
              return new BigNumber((_dafny.Seq.of(true)).length);
            })).length);
          }), _24_v0)) {
            _coll3.add((_24_v0).plus(new BigNumber(686)));
          }
        }
        return _coll3;
      }()).length))).length),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-473),true));
    };
    static fm27(globalState) {
      let _source1 = _module.D3.create_DC8(false);
      if (_source1.is_DC8) {
        let _25___mcc_h0 = (_source1).cf12;
        let _26_cf12 = _25___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_26_cf12,!(!(_26_cf12)));
      } else if (_source1.is_DC9) {
        let _27___mcc_h1 = (_source1).cf13;
        let _28___mcc_h2 = (_source1).cf14;
        let _29___mcc_h3 = (_source1).cf15;
        let _30_cf15 = _29___mcc_h3;
        let _31_cf14 = _28___mcc_h2;
        let _32_cf13 = _27___mcc_h1;
        return _dafny.Map.Empty.slice().updateUnsafe(false,true);
      } else if (_source1.is_DC7) {
        let _33___mcc_h4 = (_source1).cf11;
        let _34_cf11 = _33___mcc_h4;
        return (_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),!(!(false)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false));
      } else {
        let _35___mcc_h5 = (_source1).cf16;
        let _36_cf16 = _35___mcc_h5;
        return (_module.D7.create_DC21(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), _dafny.Map.Empty.slice().updateUnsafe(false,true))).dtor_cf39;
      }
    };
    static fm28(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).Union(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(192), new BigNumber(255))).cardinality())))),function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(235), new BigNumber(-647))) {
          let _37_v0 = _compr_4;
          if (((new BigNumber(235)).isLessThanOrEqualTo(_37_v0)) && ((_37_v0).isLessThan(new BigNumber(-647)))) {
            _coll4.push([(_37_v0).minus(new BigNumber(54)),true]);
          }
        }
        return _coll4;
      }());
    };
    static m0(globalState) {
      let r0 = _dafny.ZERO;
      let _38_v0;
      _38_v0 = new BigNumber(142);
      let _39_v1;
      _39_v1 = _dafny.Seq.of(_38_v0);
      let _40_v2;
      _40_v2 = _dafny.Seq.UnicodeFromString("rfkqvu");
      let _41_v3;
      _41_v3 = _dafny.Set.fromElements(_38_v0, (_39_v1)[_module.__default.safeIndex(new BigNumber((_40_v2).length), new BigNumber((_39_v1).length))]);
      let _hi0 = _module.__default.safeDivisionInt(_38_v0, new BigNumber((_41_v3).length));
      for (let _42_i0 = new BigNumber(927); _42_i0.isLessThan(_hi0); _42_i0 = _42_i0.plus(_dafny.ONE)) {
        let _43_v4;
        _43_v4 = false;
        (globalState).f2 = _43_v4;
        let _44_v5;
        _44_v5 = _module.D4.create_DC11(_40_v2);
        _43_v4 = _dafny.Seq.IsProperPrefixOf(_40_v2, (_44_v5).dtor_cf17);
        let _45_v6;
        let _nw0 = new _module.C0();
        _nw0.__ctor();
        _45_v6 = _nw0;
        let _46_v7;
        _46_v7 = new _dafny.CodePoint('m'.codePointAt(0));
        let _47_v8;
        _47_v8 = _dafny.Set.fromElements(_46_v7, _46_v7, _46_v7, _46_v7, _46_v7);
        let _48_v9;
        _48_v9 = _dafny.MultiSet.fromElements(_47_v8);
        let _rhs0 = _42_i0;
        let _rhs1 = _module.__default.safeDivisionInt(_38_v0, _42_i0);
        let _rhs2 = (_48_v9).Intersect(_48_v9);
        let _rhs3 = _38_v0;
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        let _lhs2 = globalState;
        _lhs0.f18 = _rhs0;
        _lhs1.f21 = _rhs1;
        _48_v9 = _rhs2;
        _lhs2.f5 = _rhs3;
      }
      if (false) {
        let _49_v10;
        _49_v10 = false;
        (globalState).f2 = !(_dafny.Seq.IsPrefixOf(_40_v2, ((_49_v10) ? (_40_v2) : (_40_v2))));
        let _50_v11;
        _50_v11 = _dafny.Set.fromElements(_49_v10, _49_v10, _49_v10);
        if (!(_module.__default.fm0(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_50_v11).length), _38_v0)).cardinality()), _module.__default.safeDivisionInt(_38_v0, new BigNumber(775)), _49_v10, globalState))) {
          let _51_v12;
          _51_v12 = new _dafny.CodePoint('u'.codePointAt(0));
          let _52_v13;
          let _nw1 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
          _52_v13 = _nw1;
          let _53_v14;
          let _nw2 = new _module.C2();
          _nw2.__ctor(_51_v12, _52_v13, true);
          _53_v14 = _nw2;
          let _54_v15;
          let _nw3 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _54_v15 = _nw3;
          let _55_v16;
          _55_v16 = _module.D4.create_DC12(true, _54_v15, _38_v0);
          (globalState).f2 = (_55_v16).dtor_cf18;
          let _56_v17;
          _56_v17 = _dafny.Seq.of(_49_v10);
          (globalState).f2 = (_38_v0).isLessThanOrEqualTo(new BigNumber((_56_v17).length));
          _54_v15 = _54_v15;
          let _57_v18;
          let _nw4 = new _module.C0();
          _nw4.__ctor();
          _57_v18 = _nw4;
        } else {
          (globalState).f2 = _49_v10;
          let _58_v19;
          let _init0 = ((_59_v0) => function (_60_i1) {
            return _module.__default.safeModuloInt(_60_i1, _59_v0);
          })(_38_v0);
          let _nw5 = Array((new BigNumber(23)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw5.length); _i0_0++) {
            _nw5[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _58_v19 = _nw5;
          let _index0 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_58_v19).length));
          (_58_v19)[_index0] = ((_49_v10) ? (new BigNumber((_dafny.Seq.update(_39_v1, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ec")).length), new BigNumber((_39_v1).length)), _38_v0)).length)) : (new BigNumber((_41_v3).length)));
          let _index1 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_58_v19).length));
          (_58_v19)[_index1] = (_39_v1)[_module.__default.safeIndex(_38_v0, new BigNumber((_39_v1).length))];
          _40_v2 = _40_v2;
          (globalState).f21 = (_58_v19)[_module.__default.safeIndex(new BigNumber(47), new BigNumber((_58_v19).length))];
          let _61_v20;
          let _nw6 = Array((new BigNumber(7)).toNumber());
          _nw6[(_dafny.ZERO).toNumber()] = true;
          _nw6[(_dafny.ONE).toNumber()] = _49_v10;
          _nw6[(new BigNumber(2)).toNumber()] = !(_49_v10);
          _nw6[(new BigNumber(3)).toNumber()] = _49_v10;
          _nw6[(new BigNumber(4)).toNumber()] = _49_v10;
          _nw6[(new BigNumber(5)).toNumber()] = _49_v10;
          _nw6[(new BigNumber(6)).toNumber()] = _49_v10;
          _61_v20 = _nw6;
          let _62_v21;
          _62_v21 = _dafny.MultiSet.fromElements(_61_v20);
          let _index2 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_58_v19).length));
          (_58_v19)[_index2] = new BigNumber((_62_v21).cardinality());
        }
        (globalState).f5 = _module.__default.safeModuloInt(_38_v0, new BigNumber(860));
        let _63_v22;
        let _nw7 = new _module.C0();
        _nw7.__ctor();
        _63_v22 = _nw7;
        let _64_v23;
        _64_v23 = _dafny.Seq.of(!(_49_v10), _49_v10);
        let _65_v24;
        let _nw8 = Array((new BigNumber(12)).toNumber());
        _nw8[(_dafny.ZERO).toNumber()] = _38_v0;
        _nw8[(_dafny.ONE).toNumber()] = _38_v0;
        _nw8[(new BigNumber(2)).toNumber()] = new BigNumber((_64_v23).length);
        _nw8[(new BigNumber(3)).toNumber()] = _38_v0;
        _nw8[(new BigNumber(4)).toNumber()] = ((_dafny.ZERO).minus(_38_v0)).plus((_dafny.ZERO).minus(_38_v0));
        _nw8[(new BigNumber(5)).toNumber()] = _38_v0;
        _nw8[(new BigNumber(6)).toNumber()] = _38_v0;
        _nw8[(new BigNumber(7)).toNumber()] = _38_v0;
        _nw8[(new BigNumber(8)).toNumber()] = _38_v0;
        _nw8[(new BigNumber(9)).toNumber()] = _38_v0;
        _nw8[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((_38_v0).multipliedBy(new BigNumber(250)));
        _nw8[(new BigNumber(11)).toNumber()] = _38_v0;
        _65_v24 = _nw8;
        let _66_v25;
        _66_v25 = _dafny.Map.Empty.slice().updateUnsafe(_38_v0,_49_v10);
        let _index3 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_65_v24).length));
        (_65_v24)[_index3] = _module.__default.safeDivisionInt(_38_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_38_v0,_66_v25)).length));
        let _index4 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_65_v24).length));
        (_65_v24)[_index4] = (_63_v22).fm5(globalState);
      } else {
        let _67_v26;
        _67_v26 = new _dafny.CodePoint('p'.codePointAt(0));
        let _68_v27;
        _68_v27 = _dafny.MultiSet.fromElements(_67_v26);
        let _69_v28;
        let _nw9 = Array((new BigNumber(21)).toNumber()).fill([]);
        _69_v28 = _nw9;
        let _70_v29;
        _70_v29 = false;
        let _71_v30;
        _71_v30 = _module.D2.create_DC4(_38_v0, _70_v29);
        let _pat_let_tv0 = _70_v29;
        let _72_v31;
        let _nw10 = Array((new BigNumber(26)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _71_v30;
        _nw10[(_dafny.ONE).toNumber()] = _71_v30;
        _nw10[(new BigNumber(2)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(3)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(4)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(5)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(6)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(7)).toNumber()] = _module.D2.create_DC4((_dafny.ZERO).minus(_38_v0), _70_v29);
        _nw10[(new BigNumber(8)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(9)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(10)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(11)).toNumber()] = _module.D2.create_DC4(_38_v0, _70_v29);
        _nw10[(new BigNumber(12)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(13)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(14)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(15)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(16)).toNumber()] = function (_pat_let0_0) {
          return function (_73_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_74_dt__update_hcf5_h0) {
                return _module.D2.create_DC4((_73_dt__update__tmp_h0).dtor_cf4, _74_dt__update_hcf5_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_module.D2.create_DC4(_38_v0, _70_v29));
        _nw10[(new BigNumber(17)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(18)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(19)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(20)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(21)).toNumber()] = _module.D2.create_DC4(new BigNumber((_40_v2).length), _70_v29);
        _nw10[(new BigNumber(22)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(23)).toNumber()] = _71_v30;
        _nw10[(new BigNumber(24)).toNumber()] = _module.D2.create_DC4((_dafny.ZERO).minus(_38_v0), false);
        _nw10[(new BigNumber(25)).toNumber()] = _71_v30;
        _72_v31 = _nw10;
        let _index5 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_69_v28).length));
        (_69_v28)[_index5] = _72_v31;
        let _75_v32;
        _75_v32 = _dafny.Set.fromElements(_70_v29, _70_v29);
        let _76_v33;
        _76_v33 = _dafny.Seq.of(_72_v31, _72_v31, ((_70_v29) ? (_72_v31) : (_72_v31)));
        let _index6 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_69_v28).length));
        let _rhs4 = ((_68_v27).update(_67_v26, _module.__default.abs(new BigNumber((_75_v32).length)))).Difference(_68_v27);
        let _rhs5 = (_76_v33)[_module.__default.safeIndex(_38_v0, new BigNumber((_76_v33).length))];
        let _rhs6 = (_dafny.ZERO).minus(_38_v0);
        let _rhs7 = _41_v3;
        let _rhs8 = _module.__default.fm2((_38_v0).isEqualTo(new BigNumber((_dafny.MultiSet.FromArray(_39_v1)).cardinality())), (_dafny.ZERO).minus(_module.__default.fm2(_70_v29, (_dafny.ZERO).minus(_38_v0), _70_v29, _67_v26, globalState)), (_38_v0).isLessThanOrEqualTo(_38_v0), new _dafny.CodePoint('d'.codePointAt(0)), globalState);
        let _lhs3 = _69_v28;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_69_v28).length));
        let _lhs5 = globalState;
        let _lhs6 = globalState;
        _68_v27 = _rhs4;
        _lhs3[_lhs4] = _rhs5;
        _lhs5.f18 = _rhs6;
        _41_v3 = _rhs7;
        _lhs6.f5 = _rhs8;
        (globalState).f2 = _70_v29;
        let _77_v34;
        let _nw11 = Array((new BigNumber(18)).toNumber());
        _77_v34 = _nw11;
        let _78_v35;
        let _init1 = ((_79_v0) => function (_80_i2) {
          return _dafny.MultiSet.fromElements(_79_v0, _79_v0);
        })(_38_v0);
        let _nw12 = Array((new BigNumber(12)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw12.length); _i0_1++) {
          _nw12[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _78_v35 = _nw12;
        let _81_v36;
        let _nw13 = new _module.C2();
        _nw13.__ctor(_67_v26, _78_v35, _70_v29);
        _81_v36 = _nw13;
        let _index7 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_77_v34).length));
        (_77_v34)[_index7] = _81_v36;
        let _index8 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_77_v34).length));
        (_77_v34)[_index8] = _81_v36;
        let _82_v37;
        let _init2 = ((_83_v2) => function (_84_i3) {
          return !_dafny.Seq.contains(_83_v2, new _dafny.CodePoint('n'.codePointAt(0)));
        })(_40_v2);
        let _nw14 = Array((new BigNumber(4)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw14.length); _i0_2++) {
          _nw14[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _82_v37 = _nw14;
        let _85_v38;
        _85_v38 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-407)), ((_86_v0) => function (_87_i4) {
          return _86_v0;
        })(_38_v0)));
        let _index9 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_82_v37).length));
        (_82_v37)[_index9] = !(new BigNumber(((_85_v38)[_module.__default.safeIndex(_38_v0, new BigNumber((_85_v38).length))]).length)).isEqualTo(_38_v0);
        let _index10 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_82_v37).length));
        (_82_v37)[_index10] = ((_dafny.Seq.IsProperPrefixOf(_39_v1, _dafny.Seq.of(_38_v0, new BigNumber((_39_v1).length)))) ? (_70_v29) : (_70_v29));
        (globalState).f19 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-27)), ((_88_v0) => function (_89_i5) {
          return _88_v0;
        })(_38_v0));
      }
      let _90_v39;
      _90_v39 = false;
      let _91_v40;
      _91_v40 = new _dafny.CodePoint('p'.codePointAt(0));
      let _92_v41;
      let _nw15 = new _module.C1();
      _nw15.__ctor(_40_v2, _38_v0, _90_v39);
      _92_v41 = _nw15;
      let _93_v42;
      _93_v42 = _dafny.Seq.of(_92_v41);
      let _94_v43;
      _94_v43 = _dafny.Seq.of((_93_v42)[_module.__default.safeIndex(_38_v0, new BigNumber((_93_v42).length))]);
      let _95_v44;
      _95_v44 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_90_v39, _38_v0, _90_v39, _91_v40, globalState),new BigNumber((_94_v43).length));
      (globalState).f2 = !_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_40_v2).length),_38_v0)), (_95_v44).Merge(function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_39_v1).Elements) {
          let _96_v45 = _compr_5;
          if (_dafny.Seq.contains(_39_v1, _96_v45)) {
            _coll5.push([_module.__default.safeDivisionInt(_96_v45, _38_v0),new BigNumber(2)]);
          }
        }
        return _coll5;
      }()));
      let _97_v46;
      let _nw16 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
      _97_v46 = _nw16;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_97_v46).length))) {
        let _98_i6 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_98_i6)) && ((_98_i6).isLessThan(new BigNumber((_97_v46).length))))) {
          (_97_v46)[(_98_i6)] = _dafny.MultiSet.fromElements((_module.D2.create_DC5(_38_v0, (_92_v41).f25, (_92_v41).f25, (_92_v41).f25)).dtor_cf6, new BigNumber(843));
        }
      }
      let _99_v47;
      _99_v47 = _dafny.MultiSet.fromElements(_38_v0);
      let _100_v49;
      _100_v49 = _dafny.Map.Empty.slice().updateUnsafe(_38_v0,_module.__default.fm0(new BigNumber(-145), (_dafny.ZERO).minus(_38_v0), true, globalState));
      let _101_v50;
      _101_v50 = _dafny.Set.fromElements((((_100_v49).contains(_38_v0)) ? ((_100_v49).get(_38_v0)) : (_90_v39)), _90_v39);
      let _102_v51;
      _102_v51 = _dafny.Map.Empty.slice().updateUnsafe((_92_v41).f25,_91_v40);
      let _103_v52;
      _103_v52 = _dafny.Map.Empty.slice().updateUnsafe(_38_v0,_102_v51);
      let _104_v53;
      _104_v53 = _dafny.Map.Empty.slice().updateUnsafe((_92_v41).f25,_90_v39);
      let _105_v54;
      _105_v54 = _module.D7.create_DC21(_38_v0, _module.__default.fm27(globalState));
      let _106_v55;
      let _nw17 = new _module.C3();
      _nw17.__ctor(_91_v40, _97_v46, (_92_v41).f25);
      _106_v55 = _nw17;
      let _107_v56;
      let _nw18 = Array((new BigNumber(28)).toNumber());
      _nw18[(_dafny.ZERO).toNumber()] = new BigNumber((_39_v1).length);
      _nw18[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_38_v0);
      _nw18[(new BigNumber(2)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(3)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(_module.__default.fm2(_module.__default.fm0(_38_v0, _38_v0, _90_v39, globalState), new BigNumber((_99_v47).cardinality()), _90_v39, _91_v40, globalState), _38_v0);
      _nw18[(new BigNumber(5)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(616)), ((_108_v40) => function (_109_i7) {
        return _108_v40;
      })(_91_v40))).length);
      _nw18[(new BigNumber(7)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_38_v0).multipliedBy(_38_v0)));
      _nw18[(new BigNumber(9)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(10)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(11)).toNumber()] = (new BigNumber((function () {
        let _coll6 = new _dafny.Set();
        for (const _compr_6 of _dafny.IntegerRange(new BigNumber(117), new BigNumber(380))) {
          let _110_v48 = _compr_6;
          if (((new BigNumber(117)).isLessThanOrEqualTo(_110_v48)) && ((_110_v48).isLessThan(new BigNumber(380)))) {
            _coll6.add(_module.__default.safeDivisionInt(_110_v48, _38_v0));
          }
        }
        return _coll6;
      }()).length)).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_38_v0)));
      _nw18[(new BigNumber(12)).toNumber()] = (new BigNumber((_101_v50).length)).plus(new BigNumber(4));
      _nw18[(new BigNumber(13)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(14)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(15)).toNumber()] = new BigNumber(111);
      _nw18[(new BigNumber(16)).toNumber()] = (new BigNumber((_103_v52).length)).multipliedBy(_38_v0);
      _nw18[(new BigNumber(17)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(18)).toNumber()] = _module.__default.fm2(_90_v39, new BigNumber(191), _90_v39, _91_v40, globalState);
      _nw18[(new BigNumber(19)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(20)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(21)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(22)).toNumber()] = new BigNumber((((_104_v53).update(_90_v39, (_92_v41).f25)).Merge((_105_v54).dtor_cf39)).length);
      _nw18[(new BigNumber(23)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(24)).toNumber()] = _module.__default.safeModuloInt(_38_v0, (_dafny.ZERO).minus(_38_v0));
      _nw18[(new BigNumber(25)).toNumber()] = new BigNumber(77);
      _nw18[(new BigNumber(26)).toNumber()] = _38_v0;
      _nw18[(new BigNumber(27)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_106_v55,_38_v0)).length);
      _107_v56 = _nw18;
      let _index11 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length));
      (_107_v56)[_index11] = (_dafny.ZERO).minus((_38_v0).minus(new BigNumber((_40_v2).length)));
      let _111_v57;
      let _nw19 = new _module.C1();
      _nw19.__ctor(_dafny.Seq.UnicodeFromString("s"), new BigNumber((_40_v2).length), (_92_v41).f25);
      _111_v57 = _nw19;
      let _112_v58;
      _112_v58 = _module.D1.create_DC1(_90_v39);
      let _113_v59;
      _113_v59 = _dafny.Map.Empty.slice().updateUnsafe(_111_v57,_112_v58);
      let _index12 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length));
      (_107_v56)[_index12] = (_dafny.ZERO).minus(new BigNumber((((_113_v59).update(_111_v57, _112_v58)).Merge(_113_v59)).length));
      if (!(true) || (_90_v39)) {
        let _114_v60;
        _114_v60 = _dafny.Map.Empty.slice().updateUnsafe(_111_v57.f29,_dafny.Set.fromElements(_90_v39, (_92_v41).f25, _90_v39, (_92_v41).f25));
        let _115_v61;
        _115_v61 = _dafny.Seq.of((_92_v41).f25);
        let _116_v62;
        _116_v62 = _dafny.Map.Empty.slice().updateUnsafe((((_114_v60).contains(_111_v57.f29)) ? ((_114_v60).get(_111_v57.f29)) : (_101_v50)),_115_v61);
        let _117_v63;
        _117_v63 = _dafny.Seq.of(_116_v62);
        let _118_v64;
        _118_v64 = _dafny.Seq.of(_99_v47);
        let _119_v66;
        _119_v66 = _dafny.Set.fromElements(_99_v47);
        let _120_v67;
        let _nw20 = Array((new BigNumber(23)).toNumber());
        _120_v67 = _nw20;
        let _121_v68;
        _121_v68 = _module.D4.create_DC13(_120_v67, _90_v39, (_92_v41).f25);
        let _122_v69;
        _122_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(38), new BigNumber((_39_v1).length)),(_92_v41).f25);
        let _123_v70;
        let _nw21 = Array((new BigNumber(29)).toNumber());
        _nw21[(_dafny.ZERO).toNumber()] = _90_v39;
        _nw21[(_dafny.ONE).toNumber()] = _dafny.Seq.IsPrefixOf(_40_v2, _40_v2);
        _nw21[(new BigNumber(2)).toNumber()] = (_dafny.Set.fromElements((_92_v41).f25, (_92_v41).f25)).IsSubsetOf(_101_v50);
        _nw21[(new BigNumber(3)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(4)).toNumber()] = true;
        _nw21[(new BigNumber(5)).toNumber()] = false;
        _nw21[(new BigNumber(6)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(7)).toNumber()] = ((_117_v63)[_module.__default.safeIndex(_111_v57.f29, new BigNumber((_117_v63).length))]).contains(_101_v50);
        _nw21[(new BigNumber(8)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(9)).toNumber()] = _90_v39;
        _nw21[(new BigNumber(10)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(11)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(12)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(13)).toNumber()] = !(_95_v44).contains((_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))]);
        _nw21[(new BigNumber(14)).toNumber()] = true;
        _nw21[(new BigNumber(15)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(16)).toNumber()] = _90_v39;
        _nw21[(new BigNumber(17)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(18)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(19)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(20)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(21)).toNumber()] = (_92_v41).f25;
        _nw21[(new BigNumber(22)).toNumber()] = (function () {
          let _coll7 = new _dafny.Set();
          for (const _compr_7 of (_118_v64).Elements) {
            let _124_v65 = _compr_7;
            if (_dafny.Seq.contains(_118_v64, _124_v65)) {
              _coll7.add(_124_v65);
            }
          }
          return _coll7;
        }()).IsDisjointFrom(_119_v66);
        _nw21[(new BigNumber(23)).toNumber()] = (_121_v68).dtor_cf22;
        _nw21[(new BigNumber(24)).toNumber()] = (_dafny.MultiSet.FromArray(_39_v1)).IsSubsetOf(_99_v47);
        _nw21[(new BigNumber(25)).toNumber()] = !((new BigNumber((_122_v69).length)).isLessThanOrEqualTo(new BigNumber(91)));
        _nw21[(new BigNumber(26)).toNumber()] = _dafny.Seq.IsPrefixOf(_111_v57.f28, _40_v2);
        _nw21[(new BigNumber(27)).toNumber()] = ((_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))]).isLessThan(_111_v57.f29);
        _nw21[(new BigNumber(28)).toNumber()] = ((_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))]).isLessThanOrEqualTo(_38_v0);
        _123_v70 = _nw21;
        let _125_v71;
        _125_v71 = _dafny.Map.Empty.slice().updateUnsafe((_99_v47).Intersect(_99_v47),_106_v55);
        let _rhs9 = _123_v70;
        let _rhs10 = (((_125_v71).contains(_99_v47)) ? ((_125_v71).get(_99_v47)) : (_106_v55));
        _123_v70 = _rhs9;
        _106_v55 = _rhs10;
        let _126_v72;
        _126_v72 = _module.D5.create_DC15((_92_v41).f25, _123_v70, _111_v57.f28, _90_v39);
        let _127_v73;
        _127_v73 = _module.D5.create_DC16(_126_v72);
        let _128_v74;
        _128_v74 = _dafny.Map.Empty.slice().updateUnsafe((_38_v0).isLessThan(_111_v57.f29),_module.D5.create_DC16(_126_v72));
        let _129_v75;
        _129_v75 = _module.D5.create_DC16(_126_v72);
        _128_v74 = (_128_v74).update(_90_v39, _129_v75);
        let _130_v76;
        let _nw22 = new _module.C3();
        _nw22.__ctor(_91_v40, _97_v46, (_92_v41).f25);
        _130_v76 = _nw22;
        let _131_v77;
        _131_v77 = _dafny.Seq.of(_130_v76);
        _131_v77 = _dafny.Seq.Concat(_131_v77, _dafny.Seq.Concat(_131_v77, _131_v77));
        let _132_v78;
        let _nw23 = new _module.C0();
        _nw23.__ctor();
        _132_v78 = _nw23;
        (globalState).f18 = _111_v57.f29;
      } else {
        (globalState).f22 = (_dafny.ZERO).minus(_111_v57.f29);
        let _133_v79;
        _133_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-974),_92_v41);
        if ((_133_v79).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-283),_92_v41))) {
          let _134_v80;
          let _nw24 = Array((new BigNumber(13)).toNumber()).fill(false);
          _134_v80 = _nw24;
          let _index13 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_134_v80).length));
          (_134_v80)[_index13] = _dafny.Seq.IsPrefixOf(_111_v57.f28, _40_v2);
          let _index14 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_134_v80).length));
          (_134_v80)[_index14] = _90_v39;
          let _index15 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_134_v80).length));
          let _rhs11 = !_dafny.areEqual(_111_v57.f28, _111_v57.f28);
          let _rhs12 = _106_v55;
          let _lhs7 = _134_v80;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_134_v80).length));
          _lhs7[_lhs8] = _rhs11;
          _106_v55 = _rhs12;
          (globalState).f2 = (((_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))]).plus(new BigNumber(739))).isEqualTo(_111_v57.f29);
          _99_v47 = (_dafny.MultiSet.fromElements(_111_v57.f29)).Intersect((_99_v47).Difference((_99_v47).update((_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))], _module.__default.abs((_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))]))));
          (globalState).f2 = (_92_v41).f25;
        } else {
          let _index16 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length));
          let _rhs13 = (_module.__default.safeModuloInt((_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))], _111_v57.f29)).multipliedBy(_111_v57.f29);
          let _rhs14 = (_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))];
          let _lhs9 = _107_v56;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length));
          let _lhs11 = globalState;
          _lhs9[_lhs10] = _rhs13;
          _lhs11.f18 = _rhs14;
          (globalState).f2 = !((_92_v41).f25);
          let _135_v81;
          let _nw25 = new _module.C2();
          _nw25.__ctor(_91_v40, _97_v46, !((_92_v41).f25));
          _135_v81 = _nw25;
          let _136_v82;
          _136_v82 = _dafny.Seq.of(_135_v81);
          let _137_v83;
          _137_v83 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm18(new BigNumber((_136_v82).length), _111_v57.f29, globalState),_91_v40);
          let _138_v84;
          let _nw26 = new _module.C3();
          _nw26.__ctor((((_137_v83).contains(_111_v57.f28)) ? ((_137_v83).get(_111_v57.f28)) : (_91_v40)), _97_v46, !(true));
          _138_v84 = _nw26;
          let _139_v85;
          _139_v85 = _dafny.Map.Empty.slice().updateUnsafe(_107_v56,_138_v84);
          let _140_v86;
          _140_v86 = _dafny.Map.Empty.slice().updateUnsafe(false,(((_139_v85).contains(_107_v56)) ? ((_139_v85).get(_107_v56)) : (_138_v84)));
          _138_v84 = (((_140_v86).contains(true)) ? ((_140_v86).get(true)) : (_138_v84));
          let _141_v87;
          let _nw27 = Array((new BigNumber(17)).toNumber()).fill(false);
          _141_v87 = _nw27;
          let _index17 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_141_v87).length));
          (_141_v87)[_index17] = (_92_v41).f25;
          let _index18 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_141_v87).length));
          (_141_v87)[_index18] = (_138_v84).f25;
          let _142_v88;
          _142_v88 = _dafny.Map.Empty.slice().updateUnsafe(_41_v3,_100_v49);
          let _143_v90;
          _143_v90 = _dafny.Map.Empty.slice().updateUnsafe((_138_v84).f26,_111_v57.f29);
          let _144_v91;
          _144_v91 = _dafny.MultiSet.fromElements((_92_v41).f25, (_141_v87)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_141_v87).length))], (_92_v41).f25, (_92_v41).f25, (_92_v41).f25);
          _142_v88 = (_dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll8 = new _dafny.Set();
            for (const _compr_8 of (_41_v3).Elements) {
              let _145_v89 = _compr_8;
              if ((_41_v3).contains(_145_v89)) {
                _coll8.add((_145_v89).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(313))).cardinality())));
              }
            }
            return _coll8;
          }(),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-325),(_141_v87)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_141_v87).length))]))).Merge((_module.__default.fm28((_106_v55).fm3((_141_v87)[_module.__default.safeIndex(new BigNumber(169), new BigNumber((_141_v87).length))], new BigNumber(819), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_111_v57.f29,(_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))])).length), _143_v90, globalState), (((_144_v91).contains((_92_v41).f25)) ? ((_144_v91).get((_92_v41).f25)) : (_111_v57.f29)), new BigNumber(342), globalState)).update(_41_v3, _100_v49));
        }
        (globalState).f2 = (_92_v41).f25;
        (globalState).f22 = _38_v0;
        let _146_v92;
        _146_v92 = _dafny.Seq.of((_92_v41).f25);
        r0 = _module.__default.safeModuloInt((new BigNumber((_146_v92).length)).plus((_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))]), _111_v57.f29);
      }
      let _147_v93;
      _147_v93 = _dafny.MultiSet.fromElements(_90_v39);
      r0 = (((_107_v56)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_107_v56).length))]).minus(new BigNumber(((_147_v93).update((_92_v41).f25, _module.__default.abs(new BigNumber(629)))).cardinality()))).multipliedBy(new BigNumber((_41_v3).length));
      return r0;
    }
    static Main(__noArgsParameter) {
      let _148_v0;
      _148_v0 = _dafny.Seq.UnicodeFromString("qyvk");
      let _149_v1;
      let _init3 = ((_150_v0) => function (_151_i0) {
        return _dafny.Set.fromElements(new BigNumber((_150_v0).length), new BigNumber((_150_v0).length));
      })(_148_v0);
      let _nw28 = Array((new BigNumber(11)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw28.length); _i0_3++) {
        _nw28[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _149_v1 = _nw28;
      let _152_v2;
      let _nw29 = Array((new BigNumber(19)).toNumber()).fill([]);
      _152_v2 = _nw29;
      let _153_v3;
      _153_v3 = true;
      let _154_v4;
      _154_v4 = _dafny.Set.fromElements(_153_v3);
      let _155_v5;
      _155_v5 = _dafny.Seq.of(new BigNumber((_154_v4).length));
      let _156_v6;
      let _nw30 = Array((new BigNumber(20)).toNumber()).fill(false);
      _156_v6 = _nw30;
      let _157_globalState;
      let _nw31 = new _module.GlobalState();
      _nw31.__ctor(new BigNumber(460), new BigNumber(238), true, _148_v0, _149_v1, new BigNumber(698), false, new BigNumber(939), false, _152_v2, new BigNumber(659), false, new _dafny.CodePoint('w'.codePointAt(0)), false, false, new BigNumber(856), new BigNumber(966), false, new BigNumber(806), _155_v5, new BigNumber(962), new BigNumber(561), new BigNumber(182), _156_v6, new _dafny.CodePoint('w'.codePointAt(0)));
      _157_globalState = _nw31;
      let _158_v7;
      let _out0;
      _out0 = _module.__default.m0(_157_globalState);
      _158_v7 = _out0;
      let _159_v8;
      _159_v8 = _dafny.MultiSet.fromElements(_148_v0);
      let _160_v9;
      _160_v9 = _dafny.MultiSet.fromElements(_153_v3, !(!(_159_v8).equals(_159_v8)), !(!(_153_v3)) || (_module.__default.fm0(_158_v7, _158_v7, _153_v3, _157_globalState)), _153_v3, _153_v3);
      let _161_v10;
      _161_v10 = _dafny.Seq.of(_153_v3, _153_v3);
      _160_v9 = _module.__default.fm1(_158_v7, new BigNumber((_dafny.Seq.Concat(_161_v10, _161_v10)).length), _157_globalState);
      let _rhs15 = (_158_v7).multipliedBy(_158_v7);
      let _rhs16 = !(_153_v3);
      let _rhs17 = _156_v6;
      let _lhs12 = _157_globalState;
      let _lhs13 = _157_globalState;
      _lhs12.f22 = _rhs15;
      _lhs13.f2 = _rhs16;
      _156_v6 = _rhs17;
      let _162_v11;
      _162_v11 = new _dafny.CodePoint('k'.codePointAt(0));
      _162_v11 = _162_v11;
      let _163_v12;
      _163_v12 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_153_v3, _158_v7, _153_v3, _162_v11, _157_globalState),_153_v3);
      let _164_v13;
      _164_v13 = _dafny.Seq.of(_160_v9);
      let _165_v14;
      let _nw32 = Array((new BigNumber(21)).toNumber());
      _nw32[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.fromElements((((_163_v12).contains(new BigNumber(926))) ? ((_163_v12).get(new BigNumber(926))) : (_153_v3)))).Union(_160_v9);
      _nw32[(_dafny.ONE).toNumber()] = _160_v9;
      _nw32[(new BigNumber(2)).toNumber()] = _160_v9;
      _nw32[(new BigNumber(3)).toNumber()] = _160_v9;
      _nw32[(new BigNumber(4)).toNumber()] = _module.__default.fm1(_158_v7, _158_v7, _157_globalState);
      _nw32[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(true);
      _nw32[(new BigNumber(6)).toNumber()] = _160_v9;
      _nw32[(new BigNumber(7)).toNumber()] = (_164_v13)[_module.__default.safeIndex(_158_v7, new BigNumber((_164_v13).length))];
      _nw32[(new BigNumber(8)).toNumber()] = _module.__default.fm1(_158_v7, _158_v7, _157_globalState);
      _nw32[(new BigNumber(9)).toNumber()] = _160_v9;
      _nw32[(new BigNumber(10)).toNumber()] = (_160_v9).Intersect(_module.__default.fm1(_158_v7, _158_v7, _157_globalState));
      _nw32[(new BigNumber(11)).toNumber()] = _160_v9;
      _nw32[(new BigNumber(12)).toNumber()] = _module.__default.fm1(new BigNumber(-362), _158_v7, _157_globalState);
      _nw32[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.FromArray(_161_v10);
      _nw32[(new BigNumber(14)).toNumber()] = _160_v9;
      _nw32[(new BigNumber(15)).toNumber()] = _160_v9;
      _nw32[(new BigNumber(16)).toNumber()] = (_dafny.MultiSet.fromElements(_153_v3)).Intersect(_160_v9);
      _nw32[(new BigNumber(17)).toNumber()] = _160_v9;
      _nw32[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_153_v3), _module.__default.safeIndex(_158_v7, new BigNumber((_dafny.Seq.of(_153_v3)).length)), _153_v3), _module.__default.safeIndex(new BigNumber(317), new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_153_v3), _module.__default.safeIndex(_158_v7, new BigNumber((_dafny.Seq.of(_153_v3)).length)), _153_v3)).length)), _153_v3));
      _nw32[(new BigNumber(19)).toNumber()] = _160_v9;
      _nw32[(new BigNumber(20)).toNumber()] = (_160_v9).Intersect(_160_v9);
      _165_v14 = _nw32;
      let _index19 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_165_v14).length));
      (_165_v14)[_index19] = _160_v9;
      let _index20 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length));
      (_156_v6)[_index20] = true;
      let _166_v15;
      _166_v15 = _dafny.Map.Empty.slice().updateUnsafe(!((new BigNumber((_148_v0).length)).isLessThan(_module.__default.fm2(_153_v3, _158_v7, !(_153_v3), _162_v11, _157_globalState))),_dafny.MultiSet.FromArray(_161_v10));
      let _index21 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_165_v14).length));
      let _index22 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length));
      let _rhs18 = _158_v7;
      let _rhs19 = false;
      let _rhs20 = (((_166_v15).contains(_153_v3)) ? ((_166_v15).get(_153_v3)) : (_160_v9));
      let _rhs21 = (new BigNumber((_dafny.Seq.of(_module.__default.fm2(_153_v3, new BigNumber((_dafny.Seq.update(_148_v0, _module.__default.safeIndex(_158_v7, new BigNumber((_148_v0).length)), _162_v11)).length), _153_v3, _162_v11, _157_globalState))).length)).isLessThanOrEqualTo(_158_v7);
      let _rhs22 = _158_v7;
      let _lhs14 = _165_v14;
      let _lhs15 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_165_v14).length));
      let _lhs16 = _156_v6;
      let _lhs17 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length));
      _158_v7 = _rhs18;
      _153_v3 = _rhs19;
      _lhs14[_lhs15] = _rhs20;
      _lhs16[_lhs17] = _rhs21;
      _158_v7 = _rhs22;
      let _167_v16;
      let _out1;
      _out1 = _module.__default.m0(_157_globalState);
      _167_v16 = _out1;
      let _168_v17;
      _168_v17 = _dafny.Map.Empty.slice().updateUnsafe((_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))],(_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))]);
      let _169_v18;
      _169_v18 = _dafny.MultiSet.fromElements(new BigNumber((_168_v17).length));
      let _170_v19;
      _170_v19 = _dafny.Set.fromElements(_167_v16, _158_v7);
      let _171_v20;
      let _nw33 = Array((new BigNumber(19)).toNumber());
      _nw33[(_dafny.ZERO).toNumber()] = _167_v16;
      _nw33[(_dafny.ONE).toNumber()] = _167_v16;
      _nw33[(new BigNumber(2)).toNumber()] = _167_v16;
      _nw33[(new BigNumber(3)).toNumber()] = _167_v16;
      _nw33[(new BigNumber(4)).toNumber()] = new BigNumber((_148_v0).length);
      _nw33[(new BigNumber(5)).toNumber()] = _167_v16;
      _nw33[(new BigNumber(6)).toNumber()] = new BigNumber(757);
      _nw33[(new BigNumber(7)).toNumber()] = _167_v16;
      _nw33[(new BigNumber(8)).toNumber()] = _167_v16;
      _nw33[(new BigNumber(9)).toNumber()] = _module.__default.fm2((_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))], _158_v7, _153_v3, _162_v11, _157_globalState);
      _nw33[(new BigNumber(10)).toNumber()] = new BigNumber((_169_v18).cardinality());
      _nw33[(new BigNumber(11)).toNumber()] = _167_v16;
      _nw33[(new BigNumber(12)).toNumber()] = new BigNumber((_170_v19).length);
      _nw33[(new BigNumber(13)).toNumber()] = _158_v7;
      _nw33[(new BigNumber(14)).toNumber()] = _module.__default.fm2(_153_v3, _167_v16, _153_v3, new _dafny.CodePoint('w'.codePointAt(0)), _157_globalState);
      _nw33[(new BigNumber(15)).toNumber()] = _167_v16;
      _nw33[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_158_v7);
      _nw33[(new BigNumber(17)).toNumber()] = _158_v7;
      _nw33[(new BigNumber(18)).toNumber()] = _module.__default.fm2((_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))], new BigNumber((_154_v4).length), _153_v3, _162_v11, _157_globalState);
      _171_v20 = _nw33;
      let _hi1 = _167_v16;
      for (let _172_i1 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_155_v5,_171_v20)).update(_155_v5, _171_v20)).length); _172_i1.isLessThan(_hi1); _172_i1 = _172_i1.plus(_dafny.ONE)) {
        let _173_v21;
        let _nw34 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
        _173_v21 = _nw34;
        let _index23 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_173_v21).length));
        (_173_v21)[_index23] = _163_v12;
        let _index24 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_165_v14).length));
        let _index25 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_173_v21).length));
        let _rhs23 = _dafny.MultiSet.fromElements((_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))], (_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))]);
        let _rhs24 = _163_v12;
        let _lhs18 = _165_v14;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_165_v14).length));
        let _lhs20 = _173_v21;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_173_v21).length));
        _lhs18[_lhs19] = _rhs23;
        _lhs20[_lhs21] = _rhs24;
        _148_v0 = _dafny.Seq.update((((_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))]) ? (_148_v0) : (_148_v0)), _module.__default.safeIndex(_module.__default.safeDivisionInt(_172_i1, _167_v16), new BigNumber(((((_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))]) ? (_148_v0) : (_148_v0))).length)), new _dafny.CodePoint('n'.codePointAt(0)));
        let _174_v22;
        let _out2;
        _out2 = _module.__default.m0(_157_globalState);
        _174_v22 = _out2;
        let _175_v23;
        let _nw35 = new _module.C0();
        _nw35.__ctor();
        _175_v23 = _nw35;
      }
      let _176_v24;
      let _out3;
      _out3 = _module.__default.m0(_157_globalState);
      _176_v24 = _out3;
      _176_v24 = _167_v16;
      (_157_globalState).f2 = (_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))];
      let _index26 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_171_v20).length));
      (_171_v20)[_index26] = _158_v7;
      let _index27 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_171_v20).length));
      (_171_v20)[_index27] = new BigNumber(967);
      (_157_globalState).f5 = _module.__default.safeDivisionInt(_module.__default.fm2((((_168_v17).contains(true)) ? ((_168_v17).get(true)) : (!(_153_v3))), new BigNumber(317), (_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))], _162_v11, _157_globalState), (_dafny.ZERO).minus(new BigNumber(((function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(670), new BigNumber(941))) {
          let _177_v25 = _compr_9;
          if (((new BigNumber(670)).isLessThanOrEqualTo(_177_v25)) && ((_177_v25).isLessThan(new BigNumber(941)))) {
            _coll9.push([_module.__default.safeDivisionInt(_177_v25, new BigNumber(913)),_153_v3]);
          }
        }
        return _coll9;
      }()).Merge(_module.__default.fm26((_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))], (((_160_v9).contains(_153_v3)) ? ((_160_v9).get(_153_v3)) : (_167_v16)), new BigNumber((_148_v0).length), _157_globalState))).length)));
      let _178_v26;
      let _init4 = ((_179_v18) => function (_180_i2) {
        return _179_v18;
      })(_169_v18);
      let _nw36 = Array((new BigNumber(14)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw36.length); _i0_4++) {
        _nw36[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _178_v26 = _nw36;
      let _181_v27;
      let _nw37 = new _module.C3();
      _nw37.__ctor(_162_v11, _178_v26, ((_171_v20)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_171_v20).length))]).isEqualTo(_158_v7));
      _181_v27 = _nw37;
      let _182_v28;
      _182_v28 = _module.D7.create_DC23(_176_v24, new BigNumber(294), _153_v3, _158_v7);
      let _183_v29;
      _183_v29 = _module.D7.create_DC24(_182_v28);
      let _184_i3;
      _184_i3 = _dafny.ZERO;
      L0: {
        let _pat_let_tv1 = _153_v3;
        let _pat_let_tv2 = _153_v3;
        let _pat_let_tv3 = _162_v11;
        let _pat_let_tv4 = _156_v6;
        let _pat_let_tv5 = _156_v6;
        let _pat_let_tv6 = _148_v0;
        while (function (_source2) {
          if (_source2.is_DC21) {
            let _189___mcc_h0 = (_source2).cf38;
            let _190___mcc_h1 = (_source2).cf39;
            let _191_cf39 = _190___mcc_h1;
            let _192_cf38 = _189___mcc_h0;
            return _pat_let_tv1;
          } else if (_source2.is_DC22) {
            let _193___mcc_h2 = (_source2).cf40;
            let _194___mcc_h3 = (_source2).cf41;
            let _195___mcc_h4 = (_source2).cf42;
            let _196_cf42 = _195___mcc_h4;
            let _197_cf41 = _194___mcc_h3;
            let _198_cf40 = _193___mcc_h2;
            return _pat_let_tv2;
          } else if (_source2.is_DC23) {
            let _199___mcc_h5 = (_source2).cf43;
            let _200___mcc_h6 = (_source2).cf44;
            let _201___mcc_h7 = (_source2).cf45;
            let _202___mcc_h8 = (_source2).cf46;
            let _203_cf46 = _202___mcc_h8;
            let _204_cf45 = _201___mcc_h7;
            let _205_cf44 = _200___mcc_h6;
            let _206_cf43 = _199___mcc_h5;
            return _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("tdneecvnd"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-460)), ((_207_v11) => function (_208_i4) {
              return _207_v11;
            })(_pat_let_tv3)));
          } else if (_source2.is_DC20) {
            let _209___mcc_h9 = (_source2).cf37;
            let _210_cf37 = _209___mcc_h9;
            return (_pat_let_tv5)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_pat_let_tv4).length))];
          } else {
            let _211___mcc_h10 = (_source2).cf47;
            let _212_cf47 = _211___mcc_h10;
            return _dafny.Seq.IsPrefixOf(_pat_let_tv6, _dafny.Seq.UnicodeFromString("qarhghwh"));
          }
        }(_183_v29)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_184_i3)) {
              break L0;
            }
            _184_i3 = (_184_i3).plus(_dafny.ONE);
            let _index28 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length));
            (_156_v6)[_index28] = _153_v3;
            let _185_v30;
            let _out4;
            _out4 = _module.__default.m0(_157_globalState);
            _185_v30 = _out4;
            _148_v0 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mukobu"), _148_v0);
            let _186_v31;
            _186_v31 = _dafny.Set.fromElements(_162_v11);
            (_181_v27).m2((((_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))]) ? (_186_v31) : (_dafny.Set.fromElements(_162_v11, _162_v11, new _dafny.CodePoint('o'.codePointAt(0)), _162_v11, _162_v11))), _153_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-712)), ((_187_v0) => function (_188_i5) {
              return _187_v0;
            })(_148_v0))).length), _157_globalState);
          }
        }
      }
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_156_v6).length))) {
        let _213_i6 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_213_i6)) && ((_213_i6).isLessThan(new BigNumber((_156_v6).length))))) {
          (_156_v6)[(_213_i6)] = _153_v3;
        }
      }
      let _214_v32;
      _214_v32 = _dafny.Seq.of(_156_v6);
      let _215_v33;
      _215_v33 = _module.D3.create_DC9(new BigNumber((_214_v32).length), _162_v11, new BigNumber(326));
      let _216_v34;
      _216_v34 = _dafny.Map.Empty.slice().updateUnsafe((_215_v33).dtor_cf14,_167_v16);
      if ((_181_v27).fm3(false, _176_v24, _module.__default.fm2(_153_v3, _167_v16, (_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))], _162_v11, _157_globalState), (_216_v34).Merge(_216_v34), _157_globalState)) {
        (_157_globalState).f5 = (_171_v20)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_171_v20).length))];
        (_157_globalState).f22 = (((_169_v18).contains(_176_v24)) ? ((_169_v18).get(_176_v24)) : (_176_v24));
        _214_v32 = _214_v32;
        let _217_v35;
        let _nw38 = Array((new BigNumber(24)).toNumber()).fill(_module.D3.Default());
        _217_v35 = _nw38;
        let _218_v36;
        _218_v36 = _dafny.Map.Empty.slice().updateUnsafe(_217_v35,_dafny.Seq.Concat(_161_v10, _161_v10));
        let _rhs25 = _167_v16;
        let _rhs26 = _dafny.Map.Empty.slice().updateUnsafe(_217_v35,_dafny.Seq.Concat(_161_v10, _dafny.Seq.of(_153_v3)));
        let _rhs27 = (_169_v18).Intersect((_169_v18).Difference(_dafny.MultiSet.FromArray(_155_v5)));
        _167_v16 = _rhs25;
        _218_v36 = _rhs26;
        _169_v18 = _rhs27;
        (_157_globalState).f2 = (_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))];
      } else {
        (_157_globalState).f2 = !(_169_v18).equals(_169_v18);
        (_157_globalState).f2 = !((_154_v4).IsDisjointFrom(_154_v4));
        _153_v3 = (_181_v27).fm3(_153_v3, new BigNumber(971), (_171_v20)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_171_v20).length))], _216_v34, _157_globalState);
        let _219_v37;
        _219_v37 = _module.D8.create_DC26(_171_v20, (_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))], (_148_v0)[_module.__default.safeIndex((_171_v20)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_171_v20).length))], new BigNumber((_148_v0).length))], (_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))], _153_v3);
        let _220_v38;
        _220_v38 = _dafny.Seq.of(_219_v37);
        _220_v38 = ((_153_v3) ? (_220_v38) : ((((_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))]) ? (_220_v38) : (_220_v38))));
        let _221_v39;
        _221_v39 = _dafny.Map.Empty.slice().updateUnsafe(_158_v7,_161_v10);
        _221_v39 = (_221_v39).update(_176_v24, _dafny.Seq.of(_module.__default.fm0(_158_v7, new BigNumber(512), (_156_v6)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_156_v6).length))], _157_globalState)));
      }
      process.stdout.write((_148_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[_dafny.ZERO]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[_dafny.ONE]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[new BigNumber(2)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[new BigNumber(3)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[new BigNumber(4)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[new BigNumber(5)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[new BigNumber(6)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[new BigNumber(7)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[new BigNumber(8)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[new BigNumber(9)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_149_v1)[new BigNumber(10)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_153_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v4).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_155_v5, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_157_globalState).f3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[_dafny.ZERO]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[_dafny.ONE]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[new BigNumber(2)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[new BigNumber(3)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[new BigNumber(4)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[new BigNumber(5)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[new BigNumber(6)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[new BigNumber(7)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[new BigNumber(8)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[new BigNumber(9)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_157_globalState).f4)[new BigNumber(10)]).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_157_globalState.f19, _dafny.Seq.of(new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142), new BigNumber(142)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_globalState.f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_157_globalState).f23)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_globalState).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_158_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v8).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("qyvk")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v9).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_161_v10, _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_162_v11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-386),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_164_v13, _dafny.Seq.of(_dafny.MultiSet.fromElements()))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(13)]).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(14)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(15)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(16)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(17)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(18)]).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(19)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_165_v14)[new BigNumber(20)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_167_v16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v17).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v18).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_v19).equals(_dafny.Set.fromElements(new BigNumber(-599)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_171_v20)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_176_v24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v26)[new BigNumber(13)]).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_v28).dtor_cf43));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_v28).dtor_cf44));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_v28).dtor_cf45));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_v28).dtor_cf46));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_183_v29).dtor_cf47).dtor_cf43));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_183_v29).dtor_cf47).dtor_cf44));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_183_v29).dtor_cf47).dtor_cf45));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_183_v29).dtor_cf47).dtor_cf46));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_184_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_214_v32).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_215_v33).dtor_cf13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_215_v33).dtor_cf14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_215_v33).dtor_cf15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_216_v34).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(-599)))));
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
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
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
        return other.$tag === 1 && this.cf1 === other.cf1;
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
    static create_DC4(cf4, cf5) {
      let $dt = new D2(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC5(cf6, cf7, cf8, cf9) {
      let $dt = new D2(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D2(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D2(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7 && this.cf8 === other.cf8 && this.cf9 === other.cf9;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(_dafny.ZERO, false);
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
    static create_DC8(cf12) {
      let $dt = new D3(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC9(cf13, cf14, cf15) {
      let $dt = new D3(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC7(cf11) {
      let $dt = new D3(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC10(cf16) {
      let $dt = new D3(3);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get is_DC10() { return this.$tag === 3; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(false);
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
    static create_DC12(cf18, cf19, cf20) {
      let $dt = new D4(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC13(cf21, cf22, cf23) {
      let $dt = new D4(1);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC11(cf17) {
      let $dt = new D4(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + this.cf17.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf18 === other.cf18 && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf21 === other.cf21 && this.cf22 === other.cf22 && this.cf23 === other.cf23;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(false, [], _dafny.ZERO);
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
    static create_DC15(cf25, cf26, cf27, cf28) {
      let $dt = new D5(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC14(cf24) {
      let $dt = new D5(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC16(cf29) {
      let $dt = new D5(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + this.cf27.toVerbatimString(true) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(false, [], _dafny.Seq.UnicodeFromString(""), false);
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
    static create_DC18(cf31, cf32, cf33, cf34, cf35) {
      let $dt = new D6(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC17(cf30) {
      let $dt = new D6(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC19(cf36) {
      let $dt = new D6(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && this.cf33 === other.cf33 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC18(_dafny.ZERO, false, false, false, _dafny.ZERO);
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
    static create_DC21(cf38, cf39) {
      let $dt = new D7(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC22(cf40, cf41, cf42) {
      let $dt = new D7(1);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC23(cf43, cf44, cf45, cf46) {
      let $dt = new D7(2);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC20(cf37) {
      let $dt = new D7(3);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC24(cf47) {
      let $dt = new D7(4);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get is_DC20() { return this.$tag === 3; }
    get is_DC24() { return this.$tag === 4; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC23" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 4) {
        return "D7.DC24" + "(" + _dafny.toString(this.cf47) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41) && this.cf42 === other.cf42;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44) && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf47, other.cf47);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC21(_dafny.ZERO, _dafny.Map.Empty);
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
    static create_DC26(cf49, cf50, cf51, cf52, cf53) {
      let $dt = new D8(0);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC25(cf48) {
      let $dt = new D8(1);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC26" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf49 === other.cf49 && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52 && this.cf53 === other.cf53;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC26([], false, new _dafny.CodePoint('D'.codePointAt(0)), false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
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
      this.f2 = false;
      this.f5 = _dafny.ZERO;
      this.f18 = _dafny.ZERO;
      this.f19 = _dafny.Seq.of();
      this.f21 = _dafny.ZERO;
      this.f22 = _dafny.ZERO;
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f3 = _dafny.Seq.UnicodeFromString("");
      this._f4 = [];
      this._f6 = false;
      this._f7 = _dafny.ZERO;
      this._f8 = false;
      this._f9 = [];
      this._f10 = _dafny.ZERO;
      this._f11 = false;
      this._f12 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f13 = false;
      this._f14 = false;
      this._f15 = _dafny.ZERO;
      this._f16 = _dafny.ZERO;
      this._f17 = false;
      this._f20 = _dafny.ZERO;
      this._f23 = [];
      this._f24 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this).f18 = f18;
      (_this).f19 = f19;
      (_this)._f20 = f20;
      (_this).f21 = f21;
      (_this).f22 = f22;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
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
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((((true) ? (new BigNumber(-144)) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(610),true)).length))))).plus(new BigNumber(-131)));
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.of(false)).length)).plus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("roa"), _dafny.Seq.UnicodeFromString("w"))).length));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f25 = false;
      this.f28 = _dafny.Seq.UnicodeFromString("");
      this.f29 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f28, f29, f25) {
      let _this = this;
      (_this).f28 = f28;
      (_this).f29 = f29;
      (_this)._f25 = f25;
      return;
    }
    fm14(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.fromElements((_this).f25, (_this).f25, (_this).f25, false)).Union(_dafny.MultiSet.fromElements((_this).f25))).cardinality())), _this.f29);
    };
    fm15(globalState) {
      let _this = this;
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), function (_222_i0) {
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_this.f29), _dafny.Seq.of(new BigNumber((function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of (_dafny.MultiSet.fromElements(_this.f29)).Elements) {
            let _223_v0 = _compr_10;
            if ((_dafny.MultiSet.fromElements(_this.f29)).contains(_223_v0)) {
              _coll10.add(_module.__default.safeModuloInt(_223_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-186),new BigNumber((function () {
                let _coll11 = new _dafny.Map();
                for (const _compr_11 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),new _dafny.CodePoint('j'.codePointAt(0)))).Keys.Elements) {
                  let _224_v1 = _compr_11;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),new _dafny.CodePoint('j'.codePointAt(0)))).contains(_224_v1)) {
                    _coll11.push([_224_v1,new BigNumber(873)]);
                  }
                }
                return _coll11;
              }()).length))).length)));
            }
          }
          return _coll10;
        }()).length)))).length);
      });
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _225_v0;
      _225_v0 = _module.D3.create_DC7(_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f25));
      let _226_v2;
      _226_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f25);
      let _227_v3;
      _227_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,false);
      let _228_v4;
      let _nw39 = Array((new BigNumber(14)).toNumber());
      _nw39[(_dafny.ZERO).toNumber()] = (_this).f25;
      _nw39[(_dafny.ONE).toNumber()] = (_this).f25;
      _nw39[(new BigNumber(2)).toNumber()] = (_this).f25;
      _nw39[(new BigNumber(3)).toNumber()] = (_this).f25;
      _nw39[(new BigNumber(4)).toNumber()] = (_this).f25;
      _nw39[(new BigNumber(5)).toNumber()] = false;
      _nw39[(new BigNumber(6)).toNumber()] = (_this).f25;
      _nw39[(new BigNumber(7)).toNumber()] = (_this).f25;
      _nw39[(new BigNumber(8)).toNumber()] = (_this).f25;
      _nw39[(new BigNumber(9)).toNumber()] = (p2)[_module.__default.safeIndex(p1, new BigNumber((p2).length))];
      _nw39[(new BigNumber(10)).toNumber()] = (((_227_v3).contains((_this).f25)) ? ((_227_v3).get((_this).f25)) : ((_this).f25));
      _nw39[(new BigNumber(11)).toNumber()] = (_this).f25;
      _nw39[(new BigNumber(12)).toNumber()] = (_this).f25;
      _nw39[(new BigNumber(13)).toNumber()] = (_this).f25;
      _228_v4 = _nw39;
      let _229_v5;
      _229_v5 = _dafny.Map.Empty.slice().updateUnsafe(_228_v4,_226_v2);
      let _230_v7;
      _230_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(270),p1);
      let _231_v8;
      _231_v8 = _dafny.Seq.of(function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_230_v7).Keys.Elements) {
          let _232_v6 = _compr_12;
          if ((_230_v7).contains(_232_v6)) {
            _coll12.push([_module.__default.safeDivisionInt(_232_v6, _this.f29),(_this).f25]);
          }
        }
        return _coll12;
      }());
      let _233_v9;
      let _nw40 = Array((new BigNumber(21)).toNumber());
      _nw40[(_dafny.ZERO).toNumber()] = (_225_v0).dtor_cf11;
      _nw40[(_dafny.ONE).toNumber()] = function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of _dafny.IntegerRange(new BigNumber(148), new BigNumber(557))) {
          let _234_v1 = _compr_13;
          if (((new BigNumber(148)).isLessThanOrEqualTo(_234_v1)) && ((_234_v1).isLessThan(new BigNumber(557)))) {
            _coll13.push([(_234_v1).minus(_this.f29),(_this).f25]);
          }
        }
        return _coll13;
      }();
      _nw40[(new BigNumber(2)).toNumber()] = (_226_v2).Merge((((_229_v5).contains(_228_v4)) ? ((_229_v5).get(_228_v4)) : (_226_v2)));
      _nw40[(new BigNumber(3)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(4)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(5)).toNumber()] = ((true) ? (_226_v2) : (_226_v2));
      _nw40[(new BigNumber(6)).toNumber()] = (_231_v8)[_module.__default.safeIndex(new BigNumber(-530), new BigNumber((_231_v8).length))];
      _nw40[(new BigNumber(7)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(8)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(9)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(10)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(11)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_227_v3).length),(_this).f25);
      _nw40[(new BigNumber(13)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(14)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(15)).toNumber()] = (_226_v2).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).f25));
      _nw40[(new BigNumber(16)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(17)).toNumber()] = (_226_v2).Merge(_226_v2);
      _nw40[(new BigNumber(18)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(19)).toNumber()] = _226_v2;
      _nw40[(new BigNumber(20)).toNumber()] = _226_v2;
      _233_v9 = _nw40;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_233_v9).length))) {
        let _235_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_235_i0)) && ((_235_i0).isLessThan(new BigNumber((_233_v9).length))))) {
          (_233_v9)[(_235_i0)] = _226_v2;
        }
      }
      if ((_this).f25) {
        let _236_v10;
        _236_v10 = _dafny.Seq.of((new BigNumber((_this.f28).length)).plus(_this.f29));
        let _237_v11;
        _237_v11 = new _dafny.CodePoint('f'.codePointAt(0));
        let _238_v12;
        _238_v12 = _dafny.MultiSet.fromElements(_237_v11, _237_v11);
        (_this).f29 = (_236_v10)[_module.__default.safeIndex(new BigNumber((_238_v12).cardinality()), new BigNumber((_236_v10).length))];
        (_this).f28 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(107)), ((_239_v11) => function (_240_i1) {
          return _239_v11;
        })(_237_v11)), _dafny.Seq.of(_237_v11)), _module.__default.safeIndex(_module.__default.safeDivisionInt(p1, _this.f29), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(107)), ((_241_v11) => function (_242_i1) {
          return _241_v11;
        })(_237_v11)), _dafny.Seq.of(_237_v11))).length)), new _dafny.CodePoint('c'.codePointAt(0)));
        (globalState).f2 = (p2)[_module.__default.safeIndex(p1, new BigNumber((p2).length))];
        let _243_v13;
        let _nw41 = new _module.C0();
        _nw41.__ctor();
        _243_v13 = _nw41;
        _237_v11 = _module.__default.fm16(_this.f29, globalState);
      } else {
        (_this).f28 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(530)), function (_244_i2) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        });
        let _245_v14;
        _245_v14 = new _dafny.CodePoint('y'.codePointAt(0));
        let _246_v15;
        _246_v15 = _dafny.Set.fromElements(_this.f29, new BigNumber((_230_v7).length));
        let _247_v16;
        _247_v16 = _dafny.Seq.of(_246_v15);
        let _248_v17;
        _248_v17 = _dafny.Seq.of(_module.__default.fm2((_this).f25, new BigNumber((_dafny.Seq.of((_this).f25)).length), (((_226_v2).contains(new BigNumber((_this.f28).length))) ? ((_226_v2).get(new BigNumber((_this.f28).length))) : ((_this).f25)), _245_v14, globalState), new BigNumber(((_247_v16)[_module.__default.safeIndex(p1, new BigNumber((_247_v16).length))]).length), _this.f29);
        (globalState).f19 = _248_v17;
        let _index29 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_228_v4).length));
        (_228_v4)[_index29] = (_this).f25;
        let _index30 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_228_v4).length));
        (_228_v4)[_index30] = (_this).f25;
        let _249_v18;
        let _nw42 = new _module.C0();
        _nw42.__ctor();
        _249_v18 = _nw42;
        let _250_v19;
        _250_v19 = _module.D3.create_DC9(p1, _245_v14, _this.f29);
        let _251_v20;
        _251_v20 = _dafny.Map.Empty.slice().updateUnsafe(_250_v19,(_this).f25);
        _251_v20 = (_251_v20).update(_250_v19, (_228_v4)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_228_v4).length))]);
      }
      let _252_v21;
      let _init5 = function (_253_i3) {
        return _module.__default.safeModuloInt(_253_i3, new BigNumber((_dafny.MultiSet.fromElements((_this).f25, (_this).f25)).cardinality()));
      };
      let _nw43 = Array((new BigNumber(7)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw43.length); _i0_5++) {
        _nw43[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _252_v21 = _nw43;
      let _index31 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_252_v21).length));
      (_252_v21)[_index31] = _module.__default.safeDivisionInt(new BigNumber(218), p1);
      let _index32 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_252_v21).length));
      (_252_v21)[_index32] = (p1).plus(new BigNumber(-798));
      let _index33 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_252_v21).length));
      (_252_v21)[_index33] = _this.f29;
      let _254_v22;
      let _nw44 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _254_v22 = _nw44;
      _254_v22 = _254_v22;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_228_v4).length))) {
        let _255_i4 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_255_i4)) && ((_255_i4).isLessThan(new BigNumber((_228_v4).length))))) {
          (_228_v4)[(_255_i4)] = (_this).f25;
        }
      }
      r0 = (_this).f25;
      r1 = _this.f29;
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let _256_v0;
      let _init6 = function (_257_i0) {
        return (_257_i0).minus(_this.f29);
      };
      let _nw45 = Array((new BigNumber(4)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw45.length); _i0_6++) {
        _nw45[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _256_v0 = _nw45;
      let _index34 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length));
      (_256_v0)[_index34] = _this.f29;
      let _258_v1;
      _258_v1 = _dafny.Seq.of((_this).f25, (_this).f25);
      let _index35 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length));
      (_256_v0)[_index35] = _module.__default.safeDivisionInt(_this.f29, (new BigNumber((_258_v1).length)).plus(_this.f29));
      let _259_v2;
      let _nw46 = new _module.C0();
      _nw46.__ctor();
      _259_v2 = _nw46;
      let _260_v3;
      _260_v3 = _dafny.MultiSet.fromElements(_this.f28);
      let _261_v4;
      _261_v4 = _dafny.Map.Empty.slice().updateUnsafe((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))],(_this).f25);
      let _262_v5;
      let _nw47 = Array((new BigNumber(19)).toNumber());
      _nw47[(_dafny.ZERO).toNumber()] = (_this).f25;
      _nw47[(_dafny.ONE).toNumber()] = (_this).f25;
      _nw47[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements(_this.f28, _this.f28)).IsSubsetOf(_260_v3);
      _nw47[(new BigNumber(3)).toNumber()] = (new BigNumber(43)).isLessThan(_this.f29);
      _nw47[(new BigNumber(4)).toNumber()] = _module.__default.fm0(_this.f29, _this.f29, (_this).f25, globalState);
      _nw47[(new BigNumber(5)).toNumber()] = true;
      _nw47[(new BigNumber(6)).toNumber()] = _module.__default.fm0((_dafny.ZERO).minus((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]), new BigNumber((_258_v1).length), !((_this).f25), globalState);
      _nw47[(new BigNumber(7)).toNumber()] = (_this).f25;
      _nw47[(new BigNumber(8)).toNumber()] = (_this).f25;
      _nw47[(new BigNumber(9)).toNumber()] = (_this).f25;
      _nw47[(new BigNumber(10)).toNumber()] = true;
      _nw47[(new BigNumber(11)).toNumber()] = _dafny.areEqual(_this.f28, _this.f28);
      _nw47[(new BigNumber(12)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of((_this).f25), _258_v1);
      _nw47[(new BigNumber(13)).toNumber()] = (_258_v1)[_module.__default.safeIndex((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], new BigNumber((_258_v1).length))];
      _nw47[(new BigNumber(14)).toNumber()] = (_this).f25;
      _nw47[(new BigNumber(15)).toNumber()] = (_this).f25;
      _nw47[(new BigNumber(16)).toNumber()] = (((_261_v4).contains(new BigNumber(299))) ? ((_261_v4).get(new BigNumber(299))) : (false));
      _nw47[(new BigNumber(17)).toNumber()] = (_this).f25;
      _nw47[(new BigNumber(18)).toNumber()] = _module.__default.fm0(new BigNumber(-373), new BigNumber((_this.f28).length), !((_this).f25), globalState);
      _262_v5 = _nw47;
      let _index36 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length));
      (_262_v5)[_index36] = _module.__default.fm0(_this.f29, _this.f29, (_this).f25, globalState);
      let _263_v6;
      _263_v6 = new _dafny.CodePoint('i'.codePointAt(0));
      let _264_v7;
      _264_v7 = _dafny.Map.Empty.slice().updateUnsafe((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))],_263_v6);
      let _265_v8;
      _265_v8 = _dafny.Seq.of(new BigNumber((_264_v7).length));
      let _266_v9;
      _266_v9 = _module.D3.create_DC9(_this.f29, _263_v6, new BigNumber((_265_v8).length));
      let _pat_let_tv7 = _256_v0;
      let _pat_let_tv8 = _256_v0;
      let _index37 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length));
      (_262_v5)[_index37] = function (_source3) {
        if (_source3.is_DC8) {
          let _267___mcc_h0 = (_source3).cf12;
          let _268_cf12 = _267___mcc_h0;
          return (_this).f25;
        } else if (_source3.is_DC9) {
          let _269___mcc_h1 = (_source3).cf13;
          let _270___mcc_h2 = (_source3).cf14;
          let _271___mcc_h3 = (_source3).cf15;
          let _272_cf15 = _271___mcc_h3;
          let _273_cf14 = _270___mcc_h2;
          let _274_cf13 = _269___mcc_h1;
          return ((_pat_let_tv8)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_pat_let_tv7).length))]).isLessThanOrEqualTo(_272_cf15);
        } else if (_source3.is_DC7) {
          let _275___mcc_h4 = (_source3).cf11;
          let _276_cf11 = _275___mcc_h4;
          return ((_this).f25) && ((_this).f25);
        } else {
          let _277___mcc_h5 = (_source3).cf16;
          let _278_cf16 = _277___mcc_h5;
          return (_this).f25;
        }
      }(_266_v9);
      (globalState).f21 = new BigNumber((_module.__default.fm17((_this).f25, ((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]).minus(new BigNumber(298)), globalState)).length);
      let _279_v10;
      _279_v10 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(778));
      let _280_v11;
      _280_v11 = _dafny.Seq.of(_279_v10, ((_279_v10).update((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))], (_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))])).update(true, new BigNumber(48)));
      let _hi2 = new BigNumber((_280_v11).length);
      for (let _281_i1 = (_259_v2).fm6(_265_v8, (_this).f25, globalState); _281_i1.isLessThan(_hi2); _281_i1 = _281_i1.plus(_dafny.ONE)) {
        let _282_v12;
        _282_v12 = _dafny.Seq.of(_259_v2, _259_v2, _259_v2, _259_v2);
        let _283_v13;
        let _nw48 = Array((new BigNumber(17)).toNumber());
        _nw48[(_dafny.ZERO).toNumber()] = _259_v2;
        _nw48[(_dafny.ONE).toNumber()] = (((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))]) ? (_259_v2) : (_259_v2));
        _nw48[(new BigNumber(2)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(3)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(4)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(5)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(6)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(7)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(8)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(9)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(10)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(11)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(12)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(13)).toNumber()] = (_282_v12)[_module.__default.safeIndex(_this.f29, new BigNumber((_282_v12).length))];
        _nw48[(new BigNumber(14)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(15)).toNumber()] = _259_v2;
        _nw48[(new BigNumber(16)).toNumber()] = _259_v2;
        _283_v13 = _nw48;
        let _index38 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_283_v13).length));
        (_283_v13)[_index38] = (_282_v12)[_module.__default.safeIndex((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], new BigNumber((_282_v12).length))];
        let _284_v14;
        let _nw49 = Array((new BigNumber(9)).toNumber()).fill([]);
        _284_v14 = _nw49;
        let _index39 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_284_v14).length));
        (_284_v14)[_index39] = _262_v5;
        let _285_v15;
        _285_v15 = _module.D1.create_DC1((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))]);
        let _286_v16;
        _286_v16 = _dafny.Seq.of(_285_v15);
        let _287_v17;
        _287_v17 = _module.D4.create_DC11(_module.__default.fm18(new BigNumber((_286_v16).length), new BigNumber((_dafny.Seq.of((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))], false, (_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))], (_this).f25, (_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))])).length), globalState));
        let _index40 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_283_v13).length));
        let _index41 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_284_v14).length));
        let _rhs28 = (_287_v17).dtor_cf17;
        let _rhs29 = (_279_v10).Merge(_279_v10);
        let _rhs30 = _259_v2;
        let _rhs31 = _262_v5;
        let _lhs22 = _this;
        let _lhs23 = _283_v13;
        let _lhs24 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_283_v13).length));
        let _lhs25 = _284_v14;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_284_v14).length));
        _lhs22.f28 = _rhs28;
        _279_v10 = _rhs29;
        _lhs23[_lhs24] = _rhs30;
        _lhs25[_lhs26] = _rhs31;
        let _288_v18;
        let _init7 = function (_289_i2) {
          return _module.D3.create_DC8((_this).f25);
        };
        let _nw50 = Array((new BigNumber(26)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw50.length); _i0_7++) {
          _nw50[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _288_v18 = _nw50;
        let _290_v19;
        _290_v19 = _module.D3.create_DC8((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))]);
        let _index42 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_288_v18).length));
        (_288_v18)[_index42] = _290_v19;
        let _index43 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_288_v18).length));
        let _rhs32 = (new BigNumber((_this.f28).length)).multipliedBy(_this.f29);
        let _rhs33 = (((_this).f25) ? (_256_v0) : ((((_this).f25) ? (_256_v0) : (_256_v0))));
        let _rhs34 = _263_v6;
        let _rhs35 = _290_v19;
        let _lhs27 = _this;
        let _lhs28 = _288_v18;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_288_v18).length));
        _lhs27.f29 = _rhs32;
        _256_v0 = _rhs33;
        _263_v6 = _rhs34;
        _lhs28[_lhs29] = _rhs35;
        (globalState).f2 = true;
        let _index44 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length));
        let _index45 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length));
        let _rhs36 = true;
        let _rhs37 = (_281_i1).multipliedBy(_281_i1);
        let _rhs38 = _266_v9;
        let _rhs39 = (_281_i1).plus((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]);
        let _rhs40 = (((_this).f25) ? (_this.f29) : ((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]));
        let _lhs30 = globalState;
        let _lhs31 = _256_v0;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length));
        let _lhs33 = _256_v0;
        let _lhs34 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length));
        let _lhs35 = globalState;
        _lhs30.f2 = _rhs36;
        _lhs31[_lhs32] = _rhs37;
        _266_v9 = _rhs38;
        _lhs33[_lhs34] = _rhs39;
        _lhs35.f5 = _rhs40;
      }
      if (_module.__default.fm0((_dafny.ZERO).minus((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]), (_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], (_this).f25, globalState)) {
        _263_v6 = new _dafny.CodePoint('w'.codePointAt(0));
        if ((((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))]) ? ((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))]) : (false))) {
          let _291_v20;
          _291_v20 = _dafny.MultiSet.fromElements((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))], true);
          let _292_v21;
          _292_v21 = _dafny.Map.Empty.slice().updateUnsafe((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))],_291_v20);
          _292_v21 = _292_v21;
          let _index46 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length));
          (_256_v0)[_index46] = (_dafny.ZERO).minus(_this.f29);
          let _index47 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length));
          (_256_v0)[_index47] = new BigNumber(-356);
          let _index48 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length));
          (_262_v5)[_index48] = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_this.f28, _module.__default.safeIndex(_this.f29, new BigNumber((_this.f28).length)), _263_v6), _this.f28);
          (globalState).f2 = (_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))];
        } else {
          let _293_v22;
          _293_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], (_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], (_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))], globalState),(function (_pat_let2_0) {
            return function (_294_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_295_dt__update_hcf17_h0) {
                  return _module.D4.create_DC11(_295_dt__update_hcf17_h0);
                }(_pat_let3_0);
              }(_this.f28);
            }(_pat_let2_0);
          }(_module.D4.create_DC11(_this.f28))).dtor_cf17);
          let _296_v23;
          _296_v23 = _dafny.Seq.of(_this.f28, _dafny.Seq.update(_this.f28, _module.__default.safeIndex((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], new BigNumber((_this.f28).length)), _263_v6), _this.f28, (((_293_v22).contains((_this).f25)) ? ((_293_v22).get((_this).f25)) : (_dafny.Seq.UnicodeFromString("idiijp"))));
          let _297_v24;
          let _nw51 = Array((new BigNumber(23)).toNumber()).fill([]);
          _297_v24 = _nw51;
          let _298_v25;
          _298_v25 = _dafny.Map.Empty.slice().updateUnsafe((_296_v23)[_module.__default.safeIndex(_this.f29, new BigNumber((_296_v23).length))],_297_v24);
          _298_v25 = (_298_v25).update(_this.f28, _297_v24);
          (globalState).f18 = (new BigNumber(-494)).multipliedBy(((!(false)) ? ((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]) : ((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))])));
          (globalState).f2 = ((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]).isLessThanOrEqualTo(_this.f29);
          let _index49 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length));
          (_256_v0)[_index49] = (_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))];
          let _299_v26;
          let _nw52 = new _module.C0();
          _nw52.__ctor();
          _299_v26 = _nw52;
        }
        (globalState).f2 = ((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))]) || ((_this).f25);
        (_this).f28 = _this.f28;
        (globalState).f5 = _this.f29;
      } else {
        let _300_v27;
        _300_v27 = _module.D4.create_DC11(_dafny.Seq.UnicodeFromString("qptks"));
        let _301_v28;
        _301_v28 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_300_v27);
        _301_v28 = (_301_v28).update((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], _300_v27);
        let _302_v29;
        _302_v29 = _dafny.Set.fromElements(((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]).multipliedBy((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]), _module.__default.safeDivisionInt(_this.f29, (_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]));
        _302_v29 = (_302_v29).Union(_302_v29);
        let _303_v30;
        _303_v30 = _module.D2.create_DC4(_this.f29, (_this).f25);
        let _source4 = _303_v30;
        if (_source4.is_DC4) {
          let _304___mcc_h6 = (_source4).cf4;
          let _305___mcc_h7 = (_source4).cf5;
          let _306_cf5 = _305___mcc_h7;
          let _307_cf4 = _304___mcc_h6;
          let _308_v31;
          _308_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-706)), ((_309_v0) => function (_310_i3) {
            return (_309_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_309_v0).length))];
          })(_256_v0)), _265_v8),_306_cf5);
          _308_v31 = (_308_v31).update(_265_v8, _306_cf5);
          let _311_v32;
          _311_v32 = _dafny.Map.Empty.slice().updateUnsafe((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))],((true) ? (_260_v3) : (_260_v3)));
          _311_v32 = (_311_v32).update(new BigNumber((_265_v8).length), _260_v3);
          let _312_v33;
          _312_v33 = _dafny.Set.fromElements(true, (_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))]);
          let _313_v34;
          _313_v34 = _dafny.Map.Empty.slice().updateUnsafe((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))],new BigNumber((_312_v33).length));
          let _314_v35;
          _314_v35 = _dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)));
          let _315_v36;
          _315_v36 = _dafny.MultiSet.fromElements(new BigNumber((_314_v35).length), new BigNumber(305));
          let _316_v37;
          _316_v37 = _dafny.Map.Empty.slice().updateUnsafe((((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))]) ? (_302_v29) : (_302_v29)),_dafny.Seq.Concat(_module.__default.fm18((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], (((_313_v34).contains(_this.f29)) ? ((_313_v34).get(_this.f29)) : ((((_315_v36).contains((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))])) ? ((_315_v36).get((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))])) : (_307_cf4)))), globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-110)), ((_317_v6) => function (_318_i4) {
            return _317_v6;
          })(_263_v6))));
          _316_v37 = _316_v37;
          let _rhs41 = _307_cf4;
          let _rhs42 = (_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))];
          let _rhs43 = (_dafny.ZERO).minus(_this.f29);
          let _rhs44 = _this.f29;
          let _rhs45 = _262_v5;
          let _lhs36 = globalState;
          let _lhs37 = globalState;
          let _lhs38 = globalState;
          _lhs36.f5 = _rhs41;
          _306_cf5 = _rhs42;
          _lhs37.f18 = _rhs43;
          _lhs38.f18 = _rhs44;
          _262_v5 = _rhs45;
        } else if (_source4.is_DC5) {
          let _319___mcc_h8 = (_source4).cf6;
          let _320___mcc_h9 = (_source4).cf7;
          let _321___mcc_h10 = (_source4).cf8;
          let _322___mcc_h11 = (_source4).cf9;
          let _323_cf9 = _322___mcc_h11;
          let _324_cf8 = _321___mcc_h10;
          let _325_cf7 = _320___mcc_h9;
          let _326_cf6 = _319___mcc_h8;
          let _index50 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length));
          (_262_v5)[_index50] = (new BigNumber(275)).isLessThan(_326_cf6);
          (globalState).f22 = new BigNumber((_module.__default.fm1((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], _326_cf6, globalState)).cardinality());
          let _327_v38;
          _327_v38 = _dafny.MultiSet.fromElements(_this.f29, _this.f29);
          let _328_v39;
          _328_v39 = new BigNumber((_327_v38).cardinality());
          let _329_v42;
          _329_v42 = _dafny.Seq.of(_264_v7);
          let _330_v43;
          let _nw53 = Array((new BigNumber(26)).toNumber());
          _nw53[(_dafny.ZERO).toNumber()] = _264_v7;
          _nw53[(_dafny.ONE).toNumber()] = (_module.__default.fm19(new BigNumber((_this.f28).length), true, globalState)).Merge(_264_v7);
          _nw53[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_258_v1).length),new _dafny.CodePoint('i'.codePointAt(0)));
          _nw53[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this.f28)[_module.__default.safeIndex(_this.f29, new BigNumber((_this.f28).length))]);
          _nw53[(new BigNumber(4)).toNumber()] = (_264_v7).update((_328_v39), new _dafny.CodePoint('e'.codePointAt(0)));
          _nw53[(new BigNumber(5)).toNumber()] = (_264_v7).Merge(_264_v7);
          _nw53[(new BigNumber(6)).toNumber()] = _264_v7;
          _nw53[(new BigNumber(7)).toNumber()] = _264_v7;
          _nw53[(new BigNumber(8)).toNumber()] = _264_v7;
          _nw53[(new BigNumber(9)).toNumber()] = (_264_v7).update(_this.f29, _263_v6);
          _nw53[(new BigNumber(10)).toNumber()] = _module.__default.fm19(_this.f29, _325_cf7, globalState);
          _nw53[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))],(_this.f28)[_module.__default.safeIndex((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], new BigNumber((_this.f28).length))]);
          _nw53[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_263_v6);
          _nw53[(new BigNumber(13)).toNumber()] = (_264_v7).Merge(_264_v7);
          _nw53[(new BigNumber(14)).toNumber()] = (_264_v7).Merge(_264_v7);
          _nw53[(new BigNumber(15)).toNumber()] = (_264_v7).Merge(_264_v7);
          _nw53[(new BigNumber(16)).toNumber()] = _264_v7;
          _nw53[(new BigNumber(17)).toNumber()] = ((_module.D5.create_DC14(_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_263_v6))).dtor_cf24).Merge(_264_v7);
          _nw53[(new BigNumber(18)).toNumber()] = function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of (_265_v8).Elements) {
              let _331_v40 = _compr_14;
              if (_dafny.Seq.contains(_265_v8, _331_v40)) {
                _coll14.push([(_331_v40).multipliedBy(_this.f29),_263_v6]);
              }
            }
            return _coll14;
          }();
          _nw53[(new BigNumber(19)).toNumber()] = _264_v7;
          _nw53[(new BigNumber(20)).toNumber()] = _264_v7;
          _nw53[(new BigNumber(21)).toNumber()] = (function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of _dafny.IntegerRange(new BigNumber(297), new BigNumber(675))) {
              let _332_v41 = _compr_15;
              if (((new BigNumber(297)).isLessThanOrEqualTo(_332_v41)) && ((_332_v41).isLessThan(new BigNumber(675)))) {
                _coll15.push([_module.__default.safeDivisionInt(_332_v41, _this.f29),new _dafny.CodePoint('s'.codePointAt(0))]);
              }
            }
            return _coll15;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(402),_263_v6));
          _nw53[(new BigNumber(22)).toNumber()] = _264_v7;
          _nw53[(new BigNumber(23)).toNumber()] = (_329_v42)[_module.__default.safeIndex((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], new BigNumber((_329_v42).length))];
          _nw53[(new BigNumber(24)).toNumber()] = _264_v7;
          _nw53[(new BigNumber(25)).toNumber()] = _264_v7;
          _330_v43 = _nw53;
          let _index51 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_330_v43).length));
          (_330_v43)[_index51] = _264_v7;
          let _index52 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_330_v43).length));
          (_330_v43)[_index52] = ((_264_v7).Merge(_264_v7)).Merge((_329_v42)[_module.__default.safeIndex(_326_cf6, new BigNumber((_329_v42).length))]);
          (globalState).f22 = (_module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(_326_cf6)).length), _this.f29)).multipliedBy((new BigNumber(331)).plus(_this.f29));
        } else if (_source4.is_DC3) {
          let _333___mcc_h12 = (_source4).cf3;
          let _334_cf3 = _333___mcc_h12;
          (globalState).f18 = _this.f29;
          (_this).f28 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ccueaym"), _this.f28), _dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), ((_335_v6) => function (_336_i5) {
            return _335_v6;
          })(_263_v6)));
          let _337_v44;
          _337_v44 = _module.D2.create_DC4(_this.f29, (_this).f25);
          let _338_v45;
          _338_v45 = _module.D2.create_DC6(_337_v44);
          _338_v45 = _module.__default.fm20((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], globalState);
          (globalState).f18 = _this.f29;
        } else {
          let _339___mcc_h13 = (_source4).cf10;
          let _340_cf10 = _339___mcc_h13;
          let _341_v46;
          _341_v46 = _dafny.MultiSet.fromElements((_262_v5)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_262_v5).length))]);
          (globalState).f2 = _module.__default.fm0(new BigNumber((_341_v46).cardinality()), (_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], (_this).f25, globalState);
          _256_v0 = ((!((_this).f25)) ? (_256_v0) : (_256_v0));
          (_this).f29 = (_259_v2).fm6(_dafny.Seq.Concat(_265_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), function (_342_i6) {
            return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(882))).cardinality());
          })), (_this).f25, globalState);
          _265_v8 = (_this).fm15(globalState);
        }
        (globalState).f2 = (_this).f25;
        let _343_v47;
        _343_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_258_v1).length),_this.f29);
        let _344_v50;
        _344_v50 = _module.D6.create_DC17(_module.__default.fm21(globalState));
        let _345_v52;
        _345_v52 = _dafny.Seq.of(_module.__default.fm21(globalState), _343_v47, _343_v47);
        let _346_v53;
        let _nw54 = Array((new BigNumber(24)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = (_343_v47).Merge(_343_v47);
        _nw54[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,new BigNumber(159));
        _nw54[(new BigNumber(2)).toNumber()] = ((function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of _dafny.IntegerRange(new BigNumber(155), new BigNumber(340))) {
            let _347_v48 = _compr_16;
            if (((new BigNumber(155)).isLessThanOrEqualTo(_347_v48)) && ((_347_v48).isLessThan(new BigNumber(340)))) {
              _coll16.push([(_347_v48).plus((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_258_v1)[_module.__default.safeIndex(_this.f29, new BigNumber((_258_v1).length))],true)).length)]);
            }
          }
          return _coll16;
        }()).update(_this.f29, _this.f29)).Merge(_343_v47);
        _nw54[(new BigNumber(3)).toNumber()] = (function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of _dafny.IntegerRange(new BigNumber(222), new BigNumber(365))) {
            let _348_v49 = _compr_17;
            if (((new BigNumber(222)).isLessThanOrEqualTo(_348_v49)) && ((_348_v49).isLessThan(new BigNumber(365)))) {
              _coll17.push([_module.__default.safeModuloInt(_348_v49, (_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]),(_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]]);
            }
          }
          return _coll17;
        }()).Merge(_343_v47);
        _nw54[(new BigNumber(4)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(5)).toNumber()] = _module.__default.fm21(globalState);
        _nw54[(new BigNumber(6)).toNumber()] = (_343_v47).update((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], (_259_v2).fm5(globalState));
        _nw54[(new BigNumber(7)).toNumber()] = (_343_v47).update((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], (_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]);
        _nw54[(new BigNumber(8)).toNumber()] = (_343_v47).update(new BigNumber((_module.__default.fm18((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], (_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))], globalState)).length), _this.f29);
        _nw54[(new BigNumber(9)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(10)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(11)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(12)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(13)).toNumber()] = ((_344_v50).dtor_cf30).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f29),(_dafny.ZERO).minus(_this.f29)));
        _nw54[(new BigNumber(14)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(15)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(16)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(17)).toNumber()] = function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of _dafny.IntegerRange(new BigNumber(463), new BigNumber(475))) {
            let _349_v51 = _compr_18;
            if (((new BigNumber(463)).isLessThanOrEqualTo(_349_v51)) && ((_349_v51).isLessThan(new BigNumber(475)))) {
              _coll18.push([(_349_v51).multipliedBy((_256_v0)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_256_v0).length))]),new BigNumber(685)]);
            }
          }
          return _coll18;
        }();
        _nw54[(new BigNumber(18)).toNumber()] = (_345_v52)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f29), new BigNumber((_345_v52).length))];
        _nw54[(new BigNumber(19)).toNumber()] = (_343_v47).Merge(_343_v47);
        _nw54[(new BigNumber(20)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(21)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(22)).toNumber()] = _343_v47;
        _nw54[(new BigNumber(23)).toNumber()] = _343_v47;
        _346_v53 = _nw54;
        _346_v53 = _346_v53;
      }
      return;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f26 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f27 = [];
      this._f25 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f26, f27, f25) {
      let _this = this;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      (_this)._f25 = f25;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(new BigNumber(195))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Map.Empty.slice().updateUnsafe((_this).f25,new BigNumber(602)))).length)));
    };
    fm8(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("w"), _dafny.Seq.UnicodeFromString("goffav")), _dafny.Seq.UnicodeFromString("atkuqwah"));
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(681), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,true)).length)));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let _350_v0;
      _350_v0 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
      (globalState).f5 = new BigNumber((_350_v0).length);
      let _351_v1;
      _351_v1 = _dafny.Seq.UnicodeFromString("w");
      let _352_v2;
      _352_v2 = _dafny.Seq.of(new BigNumber((_351_v1).length));
      let _hi3 = p2;
      for (let _353_i0 = new BigNumber((_dafny.Seq.Concat(_352_v2, _352_v2)).length); _353_i0.isLessThan(_hi3); _353_i0 = _353_i0.plus(_dafny.ONE)) {
        (globalState).f2 = !(true);
        let _354_v3;
        _354_v3 = _dafny.Set.fromElements(p1);
        let _355_v4;
        _355_v4 = _module.D2.create_DC4(new BigNumber((_354_v3).length), p1);
        let _356_v5;
        _356_v5 = _module.D2.create_DC6(_355_v4);
        let _source5 = _356_v5;
        if (_source5.is_DC4) {
          let _357___mcc_h0 = (_source5).cf4;
          let _358___mcc_h1 = (_source5).cf5;
          let _359_cf5 = _358___mcc_h1;
          let _360_cf4 = _357___mcc_h0;
          let _361_v6;
          let _nw55 = Array((new BigNumber(26)).toNumber()).fill([]);
          _361_v6 = _nw55;
          let _362_v7;
          _362_v7 = _dafny.Map.Empty.slice().updateUnsafe(_359_cf5,_361_v6);
          _362_v7 = (_362_v7).update(_359_cf5, _361_v6);
          let _363_v8;
          _363_v8 = _dafny.Map.Empty.slice().updateUnsafe(_359_cf5,(_this).f25);
          let _364_v9;
          _364_v9 = _dafny.Map.Empty.slice().updateUnsafe(_363_v8,p1);
          _364_v9 = _364_v9;
          _359_cf5 = false;
          _363_v8 = _363_v8;
        } else if (_source5.is_DC5) {
          let _365___mcc_h2 = (_source5).cf6;
          let _366___mcc_h3 = (_source5).cf7;
          let _367___mcc_h4 = (_source5).cf8;
          let _368___mcc_h5 = (_source5).cf9;
          let _369_cf9 = _368___mcc_h5;
          let _370_cf8 = _367___mcc_h4;
          let _371_cf7 = _366___mcc_h3;
          let _372_cf6 = _365___mcc_h2;
          (globalState).f2 = ((_dafny.Seq.contains(_351_v1, _module.__default.fm10(globalState))) ? (_369_cf9) : ((_this).f25));
          let _373_v10;
          _373_v10 = (p2).multipliedBy(p2);
          _373_v10 = _373_v10;
          (globalState).f19 = _dafny.Seq.Concat(_352_v2, _352_v2);
          let _374_v11;
          let _init8 = function (_375_i1) {
            return (_this).f25;
          };
          let _nw56 = Array((new BigNumber(2)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw56.length); _i0_8++) {
            _nw56[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _374_v11 = _nw56;
          let _376_v12;
          _376_v12 = _dafny.Seq.of(p1, _371_cf7);
          let _377_v13;
          _377_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_351_v1).length));
          let _378_v14;
          _378_v14 = _dafny.MultiSet.fromElements(new BigNumber(-568), _372_cf6, _353_i0, new BigNumber((_377_v13).length));
          let _index53 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_374_v11).length));
          (_374_v11)[_index53] = (!((_376_v12)[_module.__default.safeIndex(new BigNumber((function () {
            let _coll19 = new _dafny.Set();
            for (const _compr_19 of (_378_v14).Elements) {
              let _379_v15 = _compr_19;
              if ((_378_v14).contains(_379_v15)) {
                _coll19.add(_module.__default.safeModuloInt(_379_v15, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((function () {
                  let _coll20 = new _dafny.Set();
                  for (const _compr_20 of _dafny.IntegerRange(new BigNumber(308), new BigNumber(305))) {
                    let _380_v16 = _compr_20;
                    if (((new BigNumber(308)).isLessThanOrEqualTo(_380_v16)) && ((_380_v16).isLessThan(new BigNumber(305)))) {
                      _coll20.add(_module.__default.safeModuloInt(_380_v16, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-982))).length))).length)));
                    }
                  }
                  return _coll20;
                }()).length))).cardinality())));
              }
            }
            return _coll19;
          }()).length), new BigNumber((_376_v12).length))])) === ((_this).f25);
          let _381_v17;
          let _init9 = ((_382_v3) => function (_383_i2) {
            return _382_v3;
          })(_354_v3);
          let _nw57 = Array((new BigNumber(21)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw57.length); _i0_9++) {
            _nw57[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _381_v17 = _nw57;
          let _index54 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_374_v11).length));
          let _rhs46 = p1;
          let _rhs47 = true;
          let _rhs48 = _381_v17;
          let _lhs39 = _374_v11;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_374_v11).length));
          let _lhs41 = globalState;
          _lhs39[_lhs40] = _rhs46;
          _lhs41.f2 = _rhs47;
          _381_v17 = _rhs48;
        } else if (_source5.is_DC3) {
          let _384___mcc_h6 = (_source5).cf3;
          let _385_cf3 = _384___mcc_h6;
          let _386_v18;
          let _nw58 = new _module.C0();
          _nw58.__ctor();
          _386_v18 = _nw58;
          _386_v18 = ((p1) ? (_386_v18) : (_386_v18));
          (globalState).f22 = ((p1) ? (_353_i0) : (_353_i0));
          let _387_v19;
          let _out5;
          _out5 = _module.__default.m0(globalState);
          _387_v19 = _out5;
          (globalState).f2 = !((_this).f25);
        } else {
          let _388___mcc_h7 = (_source5).cf10;
          let _389_cf10 = _388___mcc_h7;
          let _390_v20;
          _390_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(685)), ((_391_v0) => function (_392_i3) {
            return _391_v0;
          })(_350_v0)),((p1) ? ((_this).f25) : ((_this).f25)));
          let _393_v21;
          _393_v21 = _dafny.Seq.of(_350_v0, _350_v0);
          let _394_v22;
          _394_v22 = _module.D2.create_DC5(_353_i0, p1, (_this).f25, (_this).f25);
          _390_v20 = (_390_v20).update(_393_v21, (_394_v22).dtor_cf9);
          (globalState).f22 = _353_i0;
          (globalState).f2 = p1;
          let _395_v23;
          let _init10 = ((_396_v3) => function (_397_i4) {
            return (new BigNumber((_396_v3).length)).isLessThan(new BigNumber(444));
          })(_354_v3);
          let _nw59 = Array((new BigNumber(6)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw59.length); _i0_10++) {
            _nw59[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _395_v23 = _nw59;
          let _index55 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_395_v23).length));
          (_395_v23)[_index55] = (_this).f25;
          let _index56 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_395_v23).length));
          (_395_v23)[_index56] = (_this).f25;
        }
        let _398_v24;
        _398_v24 = new _dafny.CodePoint('y'.codePointAt(0));
        _398_v24 = _module.__default.fm10(globalState);
        let _399_v25;
        let _nw60 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _399_v25 = _nw60;
        let _nw61 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _399_v25 = _nw61;
      }
      let _400_i5;
      _400_i5 = _dafny.ZERO;
      L1: {
        while ((_this).f25) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_400_i5)) {
              break L1;
            }
            _400_i5 = (_400_i5).plus(_dafny.ONE);
            let _401_v26;
            _401_v26 = _module.D2.create_DC4(p2, p1);
            let _402_v27;
            _402_v27 = _dafny.Seq.of(_401_v26, _module.__default.fm11(globalState), _401_v26);
            let _403_v28;
            _403_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,p2);
            let _404_v29;
            let _nw62 = Array((new BigNumber(20)).toNumber());
            _nw62[(_dafny.ZERO).toNumber()] = p1;
            _nw62[(_dafny.ONE).toNumber()] = (_this).f25;
            _nw62[(new BigNumber(2)).toNumber()] = p1;
            _nw62[(new BigNumber(3)).toNumber()] = (_this).f25;
            _nw62[(new BigNumber(4)).toNumber()] = !(p1);
            _nw62[(new BigNumber(5)).toNumber()] = p1;
            _nw62[(new BigNumber(6)).toNumber()] = (_this).f25;
            _nw62[(new BigNumber(7)).toNumber()] = !(p1) || ((_this).f25);
            _nw62[(new BigNumber(8)).toNumber()] = !(p1) || (p1);
            _nw62[(new BigNumber(9)).toNumber()] = (_this).f25;
            _nw62[(new BigNumber(10)).toNumber()] = (_this).fm3((_this).f25, p2, new BigNumber((_351_v1).length), _403_v28, globalState);
            _nw62[(new BigNumber(11)).toNumber()] = !(false);
            _nw62[(new BigNumber(12)).toNumber()] = (_this).fm3(p1, p2, p2, _403_v28, globalState);
            _nw62[(new BigNumber(13)).toNumber()] = (_this).f25;
            _nw62[(new BigNumber(14)).toNumber()] = (_this).f25;
            _nw62[(new BigNumber(15)).toNumber()] = p1;
            _nw62[(new BigNumber(16)).toNumber()] = !(true);
            _nw62[(new BigNumber(17)).toNumber()] = !((_this).f25);
            _nw62[(new BigNumber(18)).toNumber()] = (_this).f25;
            _nw62[(new BigNumber(19)).toNumber()] = true;
            _404_v29 = _nw62;
            let _index57 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_404_v29).length));
            (_404_v29)[_index57] = p1;
            let _405_v30;
            _405_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
            let _406_v31;
            _406_v31 = _dafny.Set.fromElements(p1);
            let _407_v32;
            _407_v32 = _dafny.MultiSet.fromElements(p2, p2, _module.__default.safeModuloInt(new BigNumber((_406_v31).length), new BigNumber(502)), _module.__default.safeDivisionInt(p2, p2));
            let _index58 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_404_v29).length));
            let _rhs49 = _dafny.Seq.Concat(_402_v27, _dafny.Seq.Concat(_402_v27, _402_v27));
            let _rhs50 = new BigNumber(387);
            let _rhs51 = new BigNumber((_407_v32).cardinality());
            let _rhs52 = ((true) ? ((_this).f25) : (false));
            let _rhs53 = ((_405_v30).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,p2))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,p2)).Merge(_405_v30));
            let _lhs42 = globalState;
            let _lhs43 = globalState;
            let _lhs44 = _404_v29;
            let _lhs45 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_404_v29).length));
            _402_v27 = _rhs49;
            _lhs42.f22 = _rhs50;
            _lhs43.f18 = _rhs51;
            _lhs44[_lhs45] = _rhs52;
            _405_v30 = _rhs53;
            let _408_v33;
            let _nw63 = Array((new BigNumber(2)).toNumber()).fill([]);
            _408_v33 = _nw63;
            _408_v33 = _408_v33;
            let _rhs54 = ((p1) ? (_404_v29) : (_404_v29));
            let _rhs55 = !(_module.__default.fm0(_module.__default.fm2((_404_v29)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_404_v29).length))], new BigNumber(994), !(p1), (_this).f26, globalState), _module.__default.safeDivisionInt(new BigNumber(-823), p2), _dafny.Seq.IsPrefixOf(_351_v1, _351_v1), globalState));
            let _lhs46 = globalState;
            _404_v29 = _rhs54;
            _lhs46.f2 = _rhs55;
            if (p1) {
              (globalState).f2 = (_this).f25;
              let _409_v34;
              let _init11 = function (_410_i6) {
                return (_410_i6).multipliedBy(new BigNumber(-782));
              };
              let _nw64 = Array((new BigNumber(14)).toNumber());
              for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw64.length); _i0_11++) {
                _nw64[_i0_11] = _init11(new BigNumber(_i0_11));
              }
              _409_v34 = _nw64;
              let _index59 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_409_v34).length));
              (_409_v34)[_index59] = p2;
              let _index60 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_409_v34).length));
              (_409_v34)[_index60] = (_module.__default.safeModuloInt(p2, p2)).multipliedBy(p2);
              let _411_v35;
              _411_v35 = _dafny.Seq.of((_404_v29)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_404_v29).length))]);
              let _index61 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_409_v34).length));
              (_409_v34)[_index61] = ((_409_v34)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_409_v34).length))]).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p1, p1, (_404_v29)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_404_v29).length))]), _dafny.Seq.of((_411_v35)[_module.__default.safeIndex(new BigNumber((_411_v35).length), new BigNumber((_411_v35).length))]))).length))));
              let _index62 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_404_v29).length));
              (_404_v29)[_index62] = false;
              let _index63 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_404_v29).length));
              (_404_v29)[_index63] = (_this).f25;
            } else {
              (globalState).f2 = (new BigNumber(-36)).isLessThanOrEqualTo(p2);
              let _412_v36;
              let _nw65 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
              _412_v36 = _nw65;
              let _index64 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_412_v36).length));
              (_412_v36)[_index64] = p2;
              let _413_v37;
              _413_v37 = _dafny.Seq.of(_404_v29, _404_v29);
              let _414_v38;
              _414_v38 = _dafny.MultiSet.fromElements((_this).f25);
              let _415_v39;
              _415_v39 = _dafny.MultiSet.fromElements(_404_v29, (((_this).f25) ? (_404_v29) : (_404_v29)), _404_v29, (_413_v37)[_module.__default.safeIndex(new BigNumber((_414_v38).cardinality()), new BigNumber((_413_v37).length))]);
              let _index65 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_412_v36).length));
              let _rhs56 = new BigNumber((_415_v39).cardinality());
              let _rhs57 = (_module.__default.safeModuloInt(new BigNumber(811), new BigNumber(345))).minus(_module.__default.safeDivisionInt(new BigNumber((_414_v38).cardinality()), new BigNumber((_351_v1).length)));
              let _rhs58 = _module.__default.safeDivisionInt(new BigNumber(884), _module.__default.safeModuloInt(p2, p2));
              let _lhs47 = globalState;
              let _lhs48 = globalState;
              let _lhs49 = _412_v36;
              let _lhs50 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_412_v36).length));
              _lhs47.f18 = _rhs56;
              _lhs48.f5 = _rhs57;
              _lhs49[_lhs50] = _rhs58;
              (globalState).f22 = p2;
              (globalState).f18 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p2), p2));
              let _index66 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_404_v29).length));
              (_404_v29)[_index66] = (_this).f25;
            }
          }
        }
      }
      let _416_v40;
      let _nw66 = Array((new BigNumber(18)).toNumber()).fill(false);
      _416_v40 = _nw66;
      let _index67 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length));
      (_416_v40)[_index67] = !((((_350_v0).contains(p2)) ? ((_350_v0).get(p2)) : (!(false))));
      let _index68 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length));
      (_416_v40)[_index68] = (_this).f25;
      let _417_v41;
      let _init12 = ((_418_p1, _419_p2) => function (_420_i7) {
        return _dafny.Map.Empty.slice().updateUnsafe(_418_p1,_dafny.MultiSet.fromElements(_419_p2, _419_p2));
      })(p1, p2);
      let _nw67 = Array((new BigNumber(11)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw67.length); _i0_12++) {
        _nw67[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _417_v41 = _nw67;
      let _421_v42;
      _421_v42 = _dafny.MultiSet.fromElements((_this).f25);
      let _422_v43;
      _422_v43 = _dafny.MultiSet.fromElements(new BigNumber((_421_v42).cardinality()));
      let _423_v44;
      _423_v44 = _dafny.Map.Empty.slice().updateUnsafe(p1,_422_v43);
      let _index69 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_417_v41).length));
      (_417_v41)[_index69] = _423_v44;
      let _424_v45;
      _424_v45 = _module.D2.create_DC4(p2, !((_this).f25));
      let _425_v46;
      let _nw68 = new _module.C0();
      _nw68.__ctor();
      _425_v46 = _nw68;
      let _426_v47;
      _426_v47 = _dafny.Seq.of(_425_v46, _425_v46);
      let _427_v48;
      _427_v48 = _dafny.Set.fromElements((_426_v47)[_module.__default.safeIndex(p2, new BigNumber((_426_v47).length))]);
      let _index70 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_417_v41).length));
      let _rhs59 = _dafny.Seq.Concat(_dafny.Seq.of(p2, p2, new BigNumber((_427_v48).length)), _dafny.Seq.Concat(_352_v2, _352_v2));
      let _rhs60 = (_423_v44).Merge(_423_v44);
      let _rhs61 = _module.D2.create_DC4((p2).minus(new BigNumber(-845)), p1);
      let _lhs51 = _417_v41;
      let _lhs52 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_417_v41).length));
      _352_v2 = _rhs59;
      _lhs51[_lhs52] = _rhs60;
      _424_v45 = _rhs61;
      let _428_v49;
      _428_v49 = _dafny.Seq.of(_352_v2);
      let _hi4 = new BigNumber((_421_v42).cardinality());
      for (let _429_i8 = new BigNumber((_428_v49).length); _429_i8.isLessThan(_hi4); _429_i8 = _429_i8.plus(_dafny.ONE)) {
        _351_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_351_v1, _351_v1), _dafny.Seq.Concat(_351_v1, _351_v1));
        (globalState).f5 = _429_i8;
        let _430_v50;
        _430_v50 = _dafny.Map.Empty.slice().updateUnsafe(_416_v40,_429_i8);
        let _431_v51;
        _431_v51 = _dafny.Set.fromElements(new BigNumber((_430_v50).length));
        if (!(_431_v51).contains(_429_i8)) {
          let _432_v52;
          _432_v52 = _module.D1.create_DC2(p1);
          let _rhs62 = _416_v40;
          let _rhs63 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((new BigNumber(572)).multipliedBy(new BigNumber(896)), _429_i8));
          let _rhs64 = (_416_v40)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length))];
          let _rhs65 = _432_v52;
          let _rhs66 = (_425_v46).fm5(globalState);
          let _lhs53 = globalState;
          let _lhs54 = globalState;
          let _lhs55 = globalState;
          _416_v40 = _rhs62;
          _lhs53.f5 = _rhs63;
          _lhs54.f2 = _rhs64;
          _432_v52 = _rhs65;
          _lhs55.f22 = _rhs66;
          (globalState).f18 = new BigNumber((_dafny.Seq.update(_352_v2, _module.__default.safeIndex(p2, new BigNumber((_352_v2).length)), new BigNumber((_352_v2).length))).length);
          let _433_v53;
          _433_v53 = _dafny.Map.Empty.slice().updateUnsafe(_416_v40,(_416_v40)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length))]);
          let _434_v54;
          let _nw69 = Array((new BigNumber(3)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = _433_v53;
          _nw69[(_dafny.ONE).toNumber()] = (_433_v53).update(_416_v40, (_416_v40)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length))]);
          _nw69[(new BigNumber(2)).toNumber()] = _433_v53;
          _434_v54 = _nw69;
          let _index71 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_434_v54).length));
          (_434_v54)[_index71] = _dafny.Map.Empty.slice().updateUnsafe(_416_v40,true);
          let _435_v55;
          _435_v55 = _dafny.Map.Empty.slice().updateUnsafe((_416_v40)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length))],_433_v53);
          let _index72 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_434_v54).length));
          (_434_v54)[_index72] = ((((_435_v55).contains((_416_v40)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length))])) ? ((_435_v55).get((_416_v40)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length))])) : (_433_v53))).Merge(_433_v53);
          let _436_v56;
          _436_v56 = _dafny.Map.Empty.slice().updateUnsafe(_422_v43,(p2).plus(new BigNumber((_351_v1).length)));
          _436_v56 = ((p1) ? (_436_v56) : (_436_v56));
          let _437_v57;
          let _nw70 = new _module.C0();
          _nw70.__ctor();
          _437_v57 = _nw70;
        } else {
          (globalState).f21 = new BigNumber((_351_v1).length);
          (globalState).f5 = new BigNumber(-32);
          let _438_v58;
          let _init13 = ((_439_v1) => function (_440_i9) {
            return (_dafny.MultiSet.fromElements(_439_v1, _439_v1)).Union(_dafny.MultiSet.fromElements(_439_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(441)), function (_441_i10) {
              return (_this).f26;
            })));
          })(_351_v1);
          let _nw71 = Array((new BigNumber(20)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw71.length); _i0_13++) {
            _nw71[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _438_v58 = _nw71;
          let _442_v59;
          _442_v59 = _dafny.MultiSet.fromElements(_351_v1, _351_v1);
          let _index73 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_438_v58).length));
          (_438_v58)[_index73] = _442_v59;
          let _index74 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_438_v58).length));
          (_438_v58)[_index74] = (_442_v59).update(_351_v1, _module.__default.abs(p2));
          let _443_v60;
          let _nw72 = Array((new BigNumber(23)).toNumber());
          _nw72[(_dafny.ZERO).toNumber()] = _352_v2;
          _nw72[(_dafny.ONE).toNumber()] = _352_v2;
          _nw72[(new BigNumber(2)).toNumber()] = _352_v2;
          _nw72[(new BigNumber(3)).toNumber()] = _352_v2;
          _nw72[(new BigNumber(4)).toNumber()] = _352_v2;
          _nw72[(new BigNumber(5)).toNumber()] = _352_v2;
          _nw72[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_352_v2, _352_v2);
          _nw72[(new BigNumber(7)).toNumber()] = ((_module.__default.fm0(_429_i8, _429_i8, p1, globalState)) ? (_352_v2) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-463)), ((_444_p2) => function (_445_i11) {
            return _444_p2;
          })(p2))));
          _nw72[(new BigNumber(8)).toNumber()] = (_428_v49)[_module.__default.safeIndex(new BigNumber((_351_v1).length), new BigNumber((_428_v49).length))];
          _nw72[(new BigNumber(9)).toNumber()] = _module.__default.fm12(p2, (_this).f25, (_350_v0).update(_429_i8, (_416_v40)[_module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length))]), p2, globalState);
          _nw72[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_352_v2, _352_v2);
          _nw72[(new BigNumber(11)).toNumber()] = _352_v2;
          _nw72[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_429_i8, p2), _352_v2);
          _nw72[(new BigNumber(13)).toNumber()] = _352_v2;
          _nw72[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_428_v49)).cardinality()), p2, (_dafny.ZERO).minus(p2));
          _nw72[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(p2), _352_v2);
          _nw72[(new BigNumber(16)).toNumber()] = _352_v2;
          _nw72[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(new BigNumber((_module.__default.fm13(p1, p2, globalState)).length), new BigNumber(100), p2, p2);
          _nw72[(new BigNumber(18)).toNumber()] = _352_v2;
          _nw72[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_352_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-853)), ((_446_i8) => function (_447_i12) {
            return _446_i8;
          })(_429_i8)));
          _nw72[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(237)), ((_448_p2) => function (_449_i13) {
            return _448_p2;
          })(p2));
          _nw72[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_352_v2, _352_v2);
          _nw72[(new BigNumber(22)).toNumber()] = _module.__default.fm12(p2, !((_this).f25), _350_v0, p2, globalState);
          _443_v60 = _nw72;
          let _450_v61;
          _450_v61 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.of(p2));
          let _index75 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_443_v60).length));
          (_443_v60)[_index75] = (((_this).f25) ? ((((_450_v61).contains(_429_i8)) ? ((_450_v61).get(_429_i8)) : (_352_v2))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(3)), ((_451_i8) => function (_452_i14) {
            return _451_i8;
          })(_429_i8))));
          let _453_v62;
          _453_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_421_v42).cardinality()),new BigNumber(669));
          let _454_v63;
          _454_v63 = _dafny.Seq.of(_453_v62, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(41),p2));
          let _455_v64;
          _455_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_429_i8);
          let _456_v65;
          _456_v65 = _module.D2.create_DC5(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(_module.D2.create_DC3(_454_v63)),p2)).length), (_this).fm3(p1, _429_i8, p2, _455_v64, globalState), p1, (_this).f25);
          let _index76 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_443_v60).length));
          let _index77 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length));
          let _rhs67 = _dafny.Seq.Concat(_dafny.Seq.Concat(_352_v2, _352_v2), _dafny.Seq.of(new BigNumber((_351_v1).length), _429_i8));
          let _rhs68 = _module.__default.fm2(p1, ((_456_v65).dtor_cf6).minus(p2), !((_this).f25), (_this).f26, globalState);
          let _rhs69 = _429_i8;
          let _rhs70 = (((_this).f25) ? ((_this).f25) : (p1));
          let _lhs56 = _443_v60;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_443_v60).length));
          let _lhs58 = globalState;
          let _lhs59 = globalState;
          let _lhs60 = _416_v40;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length));
          _lhs56[_lhs57] = _rhs67;
          _lhs58.f5 = _rhs68;
          _lhs59.f22 = _rhs69;
          _lhs60[_lhs61] = _rhs70;
          let _index78 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_416_v40).length));
          (_416_v40)[_index78] = (p1) === ((_this).f25);
        }
        let _457_v66;
        let _out6;
        _out6 = _module.__default.m0(globalState);
        _457_v66 = _out6;
      }
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      if (!(!(false))) {
        (globalState).f21 = p1;
        let _458_v0;
        let _nw73 = new _module.C0();
        _nw73.__ctor();
        _458_v0 = _nw73;
        let _459_v1;
        let _out7;
        _out7 = _module.__default.m0(globalState);
        _459_v1 = _out7;
        let _rhs71 = (_this).f25;
        let _rhs72 = _module.__default.safeModuloInt((_458_v0).fm5(globalState), _module.__default.safeModuloInt(p1, _459_v1));
        let _lhs62 = globalState;
        let _lhs63 = globalState;
        _lhs62.f2 = _rhs71;
        _lhs63.f5 = _rhs72;
        if ((_this).f25) {
          let _460_v2;
          _460_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _461_v3;
          _461_v3 = _dafny.Seq.of(_460_v2, _460_v2);
          let _462_v4;
          _462_v4 = _module.D2.create_DC3(_461_v3);
          let _463_v5;
          _463_v5 = _dafny.Map.Empty.slice().updateUnsafe(_462_v4,p1);
          _463_v5 = ((_463_v5).Merge(_463_v5)).Merge(_463_v5);
          let _464_v6;
          _464_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_459_v1);
          let _465_v7;
          let _nw74 = Array((new BigNumber(16)).toNumber());
          _nw74[(_dafny.ZERO).toNumber()] = true;
          _nw74[(_dafny.ONE).toNumber()] = true;
          _nw74[(new BigNumber(2)).toNumber()] = false;
          _nw74[(new BigNumber(3)).toNumber()] = !((_this).f25);
          _nw74[(new BigNumber(4)).toNumber()] = (_this).f25;
          _nw74[(new BigNumber(5)).toNumber()] = (_this).f25;
          _nw74[(new BigNumber(6)).toNumber()] = (_this).fm3((_this).f25, p1, p1, _464_v6, globalState);
          _nw74[(new BigNumber(7)).toNumber()] = (_this).f25;
          _nw74[(new BigNumber(8)).toNumber()] = (_this).f25;
          _nw74[(new BigNumber(9)).toNumber()] = (_this).f25;
          _nw74[(new BigNumber(10)).toNumber()] = (_this).f25;
          _nw74[(new BigNumber(11)).toNumber()] = (_this).f25;
          _nw74[(new BigNumber(12)).toNumber()] = !((_this).f25);
          _nw74[(new BigNumber(13)).toNumber()] = false;
          _nw74[(new BigNumber(14)).toNumber()] = (_this).f25;
          _nw74[(new BigNumber(15)).toNumber()] = (_this).f25;
          _465_v7 = _nw74;
          let _466_v8;
          _466_v8 = _dafny.MultiSet.fromElements(_465_v7, _465_v7, _465_v7, _465_v7, _465_v7);
          let _467_v9;
          _467_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_466_v8);
          r0 = !((((_467_v9).contains(!((_this).f25))) ? ((_467_v9).get(!((_this).f25))) : (_dafny.MultiSet.fromElements(_465_v7, _465_v7, _465_v7, _465_v7, _465_v7)))).contains(_465_v7);
          (globalState).f22 = _module.__default.fm2((_this).f25, p1, (_this).f25, (_this).f26, globalState);
          let _468_v10;
          _468_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p1);
          let _469_v11;
          _469_v11 = _dafny.Set.fromElements(_module.__default.fm0(p1, p1, (_this).f25, globalState));
          let _470_v12;
          _470_v12 = _dafny.Seq.of(_459_v1);
          let _471_v13;
          _471_v13 = p1;
          let _472_v14;
          let _nw75 = Array((new BigNumber(26)).toNumber());
          _nw75[(_dafny.ZERO).toNumber()] = _459_v1;
          _nw75[(_dafny.ONE).toNumber()] = ((((_468_v10).contains((_this).f25)) ? ((_468_v10).get((_this).f25)) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f25)).length))))).plus(_459_v1);
          _nw75[(new BigNumber(2)).toNumber()] = new BigNumber((p0).length);
          _nw75[(new BigNumber(3)).toNumber()] = (_458_v0).fm5(globalState);
          _nw75[(new BigNumber(4)).toNumber()] = new BigNumber(187);
          _nw75[(new BigNumber(5)).toNumber()] = _459_v1;
          _nw75[(new BigNumber(6)).toNumber()] = new BigNumber(184);
          _nw75[(new BigNumber(7)).toNumber()] = _459_v1;
          _nw75[(new BigNumber(8)).toNumber()] = _459_v1;
          _nw75[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(_459_v1, p1);
          _nw75[(new BigNumber(10)).toNumber()] = (((_this).f25) ? (_459_v1) : (_459_v1));
          _nw75[(new BigNumber(11)).toNumber()] = (p1).plus(new BigNumber((p2).length));
          _nw75[(new BigNumber(12)).toNumber()] = new BigNumber((_469_v11).length);
          _nw75[(new BigNumber(13)).toNumber()] = new BigNumber(940);
          _nw75[(new BigNumber(14)).toNumber()] = p1;
          _nw75[(new BigNumber(15)).toNumber()] = p1;
          _nw75[(new BigNumber(16)).toNumber()] = _459_v1;
          _nw75[(new BigNumber(17)).toNumber()] = (_458_v0).fm5(globalState);
          _nw75[(new BigNumber(18)).toNumber()] = _459_v1;
          _nw75[(new BigNumber(19)).toNumber()] = _459_v1;
          _nw75[(new BigNumber(20)).toNumber()] = _459_v1;
          _nw75[(new BigNumber(21)).toNumber()] = new BigNumber(676);
          _nw75[(new BigNumber(22)).toNumber()] = (_470_v12)[_module.__default.safeIndex(new BigNumber(-626), new BigNumber((_470_v12).length))];
          _nw75[(new BigNumber(23)).toNumber()] = new BigNumber(419);
          _nw75[(new BigNumber(24)).toNumber()] = new BigNumber((_module.__default.fm12(_459_v1, (_this).f25, _dafny.Map.Empty.slice().updateUnsafe(_459_v1,(_this).f25), p1, globalState)).length);
          _nw75[(new BigNumber(25)).toNumber()] = (_471_v13);
          _472_v14 = _nw75;
          let _index79 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_472_v14).length));
          (_472_v14)[_index79] = _459_v1;
          let _473_v15;
          _473_v15 = _dafny.MultiSet.fromElements((_this).fm3((_this).f25, p1, _459_v1, _464_v6, globalState));
          let _index80 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_472_v14).length));
          (_472_v14)[_index80] = _module.__default.safeDivisionInt((((_473_v15).contains((_this).f25)) ? ((_473_v15).get((_this).f25)) : (p1)), (_459_v1).minus(new BigNumber((_468_v10).length)));
          let _474_v16;
          let _nw76 = new _module.C0();
          _nw76.__ctor();
          _474_v16 = _nw76;
        } else {
          (globalState).f2 = (_459_v1).isLessThan(p1);
          let _475_v17;
          let _nw77 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _475_v17 = _nw77;
          _475_v17 = _475_v17;
          let _476_v18;
          _476_v18 = _dafny.Set.fromElements(p1, p1);
          let _477_v19;
          _477_v19 = _dafny.MultiSet.fromElements(new BigNumber((_476_v18).length));
          let _478_v20;
          _478_v20 = _dafny.Seq.of((_this).f25, (_this).f25, (p1).isLessThanOrEqualTo(new BigNumber((_477_v19).cardinality())));
          _478_v20 = _478_v20;
          let _479_v21;
          let _nw78 = Array((new BigNumber(26)).toNumber()).fill([]);
          _479_v21 = _nw78;
          let _index81 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_479_v21).length));
          (_479_v21)[_index81] = _475_v17;
          let _480_v22;
          _480_v22 = _module.D7.create_DC20(_dafny.Seq.of((_this).f25, (_this).f25, (_this).f25, !((_this).f25)));
          let _481_v23;
          let _nw79 = new _module.C1();
          _nw79.__ctor(_dafny.Seq.Concat(p0, p0), new BigNumber(((_480_v22).dtor_cf37).length), false);
          _481_v23 = _nw79;
          let _482_v24;
          _482_v24 = _dafny.Seq.of(_458_v0, _458_v0);
          let _483_v25;
          let _nw80 = Array((new BigNumber(20)).toNumber());
          _nw80[(_dafny.ZERO).toNumber()] = _482_v24;
          _nw80[(_dafny.ONE).toNumber()] = _482_v24;
          _nw80[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_482_v24, _module.__default.safeIndex(_459_v1, new BigNumber((_482_v24).length)), _458_v0);
          _nw80[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_458_v0, (_482_v24)[_module.__default.safeIndex(p1, new BigNumber((_482_v24).length))]), _482_v24);
          _nw80[(new BigNumber(4)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(5)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(6)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(7)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(8)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_458_v0);
          _nw80[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(_458_v0, _458_v0);
          _nw80[(new BigNumber(11)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(12)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(13)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(14)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_458_v0, _458_v0);
          _nw80[(new BigNumber(16)).toNumber()] = _482_v24;
          _nw80[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(_458_v0);
          _nw80[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_482_v24, _dafny.Seq.of(_458_v0));
          _nw80[(new BigNumber(19)).toNumber()] = _482_v24;
          _483_v25 = _nw80;
          let _index82 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_483_v25).length));
          (_483_v25)[_index82] = _dafny.Seq.of(_458_v0, _458_v0);
          let _484_v26;
          _484_v26 = _dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)));
          let _index83 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_479_v21).length));
          let _index84 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_483_v25).length));
          let _rhs73 = _475_v17;
          let _rhs74 = _481_v23;
          let _rhs75 = _dafny.Seq.update(_482_v24, _module.__default.safeIndex(new BigNumber((_484_v26).length), new BigNumber((_482_v24).length)), _458_v0);
          let _rhs76 = (_481_v23).f25;
          let _lhs64 = _479_v21;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_479_v21).length));
          let _lhs66 = _483_v25;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_483_v25).length));
          let _lhs68 = globalState;
          _lhs64[_lhs65] = _rhs73;
          _481_v23 = _rhs74;
          _lhs66[_lhs67] = _rhs75;
          _lhs68.f2 = _rhs76;
          let _485_v27;
          let _nw81 = Array((new BigNumber(11)).toNumber());
          _nw81[(_dafny.ZERO).toNumber()] = (_459_v1).isLessThanOrEqualTo(p1);
          _nw81[(_dafny.ONE).toNumber()] = (_481_v23).f25;
          _nw81[(new BigNumber(2)).toNumber()] = (_module.__default.fm22(_459_v1, _459_v1, p1, (_this).f25, globalState)).IsSubsetOf(_484_v26);
          _nw81[(new BigNumber(3)).toNumber()] = !((_481_v23).f25);
          _nw81[(new BigNumber(4)).toNumber()] = (_this).f25;
          _nw81[(new BigNumber(5)).toNumber()] = !((_this).f25);
          _nw81[(new BigNumber(6)).toNumber()] = (_481_v23).f25;
          _nw81[(new BigNumber(7)).toNumber()] = (p2)[_module.__default.safeIndex((_dafny.ZERO).minus(_459_v1), new BigNumber((p2).length))];
          _nw81[(new BigNumber(8)).toNumber()] = (_481_v23).f25;
          _nw81[(new BigNumber(9)).toNumber()] = ((_this).f25) === ((_this).f25);
          _nw81[(new BigNumber(10)).toNumber()] = (_this).f25;
          _485_v27 = _nw81;
          let _index85 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_485_v27).length));
          (_485_v27)[_index85] = (_481_v23).f25;
          let _index86 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_485_v27).length));
          (_485_v27)[_index86] = (_this).f25;
        }
      } else {
        (globalState).f5 = (((_this).f25) ? (p1) : (new BigNumber((_dafny.MultiSet.fromElements(!((_this).f25))).cardinality())));
        let _486_v28;
        let _nw82 = new _module.C1();
        _nw82.__ctor(_dafny.Seq.UnicodeFromString("g"), p1, (_this).f25);
        _486_v28 = _nw82;
        let _487_v29;
        let _init14 = function (_488_i0) {
          return (_488_i0).minus(new BigNumber(586));
        };
        let _nw83 = Array((new BigNumber(12)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw83.length); _i0_14++) {
          _nw83[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _487_v29 = _nw83;
        let _index87 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_487_v29).length));
        (_487_v29)[_index87] = _486_v28.f29;
        let _index88 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_487_v29).length));
        (_487_v29)[_index88] = p1;
        let _489_v30;
        _489_v30 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f25),(_this).f25);
        let _490_v31;
        _490_v31 = _module.D7.create_DC21((_487_v29)[_module.__default.safeIndex(new BigNumber(144), new BigNumber((_487_v29).length))], (_489_v30).Merge((_489_v30).update(!(true), (_this).f25)));
        _490_v31 = _module.__default.fm23(!(!((_this).f25)), globalState);
        let _491_v32;
        let _nw84 = new _module.C0();
        _nw84.__ctor();
        _491_v32 = _nw84;
      }
      let _492_v33;
      _492_v33 = p1;
      let _source6 = _module.D7.create_DC22(p1, _492_v33, (_this).f25);
      if (_source6.is_DC21) {
        let _493___mcc_h0 = (_source6).cf38;
        let _494___mcc_h1 = (_source6).cf39;
        let _495_cf39 = _494___mcc_h1;
        let _496_cf38 = _493___mcc_h0;
        let _497_v34;
        _497_v34 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),(_this).f26);
        let _498_v35;
        _498_v35 = _module.D5.create_DC14(_497_v34);
        let _499_v36;
        _499_v36 = _module.D5.create_DC16(_498_v35);
        let _500_v37;
        _500_v37 = _dafny.Set.fromElements((_this).f25, (_this).f25);
        _499_v36 = ((!((_500_v37).IsProperSubsetOf(_500_v37))) ? (_499_v36) : (_499_v36));
        (globalState).f21 = (p1).plus(p1);
        let _501_v38;
        let _init15 = function (_502_i1) {
          return (_this).f25;
        };
        let _nw85 = Array((new BigNumber(17)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw85.length); _i0_15++) {
          _nw85[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _501_v38 = _nw85;
        let _503_v39;
        _503_v39 = _dafny.Set.fromElements(p1);
        let _index89 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_501_v38).length));
        (_501_v38)[_index89] = (_503_v39).IsProperSubsetOf(_503_v39);
        let _504_v40;
        _504_v40 = _dafny.Map.Empty.slice().updateUnsafe(_496_cf38,(_this).f25);
        let _index90 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_501_v38).length));
        (_501_v38)[_index90] = !(((((_504_v40).contains(_496_cf38)) ? ((_504_v40).get(_496_cf38)) : ((_this).f25))) || (false)) || ((_this).f25);
        let _505_v41;
        _505_v41 = _dafny.Seq.of(new BigNumber((p2).length), p1);
        let _506_v42;
        _506_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,!((_501_v38)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_501_v38).length))]));
        let _507_v43;
        let _nw86 = Array((new BigNumber(12)).toNumber());
        _nw86[(_dafny.ZERO).toNumber()] = (((_this).f25) ? (p1) : (p1));
        _nw86[(_dafny.ONE).toNumber()] = _496_cf38;
        _nw86[(new BigNumber(2)).toNumber()] = _496_cf38;
        _nw86[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements((_501_v38)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_501_v38).length))], (p2)[_module.__default.safeIndex((_dafny.ZERO).minus(_496_cf38), new BigNumber((p2).length))])).length), p1);
        _nw86[(new BigNumber(4)).toNumber()] = _496_cf38;
        _nw86[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(97), _496_cf38);
        _nw86[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((((_this).f25) ? ((_505_v41)[_module.__default.safeIndex(_496_cf38, new BigNumber((_505_v41).length))]) : (p1)));
        _nw86[(new BigNumber(7)).toNumber()] = _496_cf38;
        _nw86[(new BigNumber(8)).toNumber()] = p1;
        _nw86[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_506_v42).Merge(_506_v42)).length)));
        _nw86[(new BigNumber(10)).toNumber()] = (_module.D3.create_DC9(p1, (_this).f26, new BigNumber((p0).length))).dtor_cf15;
        _nw86[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(p1, p1);
        _507_v43 = _nw86;
        let _index91 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_507_v43).length));
        (_507_v43)[_index91] = (_505_v41)[_module.__default.safeIndex(p1, new BigNumber((_505_v41).length))];
        let _508_v44;
        _508_v44 = _dafny.MultiSet.fromElements((_501_v38)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_501_v38).length))], (_this).f25);
        let _509_v45;
        _509_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,new BigNumber((_508_v44).cardinality()));
        let _index92 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_507_v43).length));
        let _rhs77 = _module.__default.safeModuloInt(new BigNumber((_509_v45).length), _module.__default.safeModuloInt(new BigNumber(-466), _496_cf38));
        let _rhs78 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p1, _496_cf38),(_this).f25);
        let _lhs69 = _507_v43;
        let _lhs70 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_507_v43).length));
        _lhs69[_lhs70] = _rhs77;
        _504_v40 = _rhs78;
      } else if (_source6.is_DC22) {
        let _510___mcc_h2 = (_source6).cf40;
        let _511___mcc_h3 = (_source6).cf41;
        let _512___mcc_h4 = (_source6).cf42;
        let _513_cf42 = _512___mcc_h4;
        let _514_cf41 = _511___mcc_h3;
        let _515_cf40 = _510___mcc_h2;
        let _516_v46;
        let _nw87 = new _module.C0();
        _nw87.__ctor();
        _516_v46 = _nw87;
        let _517_v47;
        let _out8;
        _out8 = _module.__default.m0(globalState);
        _517_v47 = _out8;
        let _518_v48;
        let _init16 = ((_519_p0) => function (_520_i2) {
          return _519_p0;
        })(p0);
        let _nw88 = Array((new BigNumber(13)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw88.length); _i0_16++) {
          _nw88[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _518_v48 = _nw88;
        let _521_v49;
        _521_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_518_v48);
        let _522_v50;
        _522_v50 = _dafny.Map.Empty.slice().updateUnsafe(_513_cf42,(_this).f26);
        let _523_v51;
        _523_v51 = _dafny.Seq.of(_522_v50, _522_v50);
        _521_v49 = (_521_v49).update((new BigNumber(-937)).isEqualTo(new BigNumber((_523_v51).length)), _518_v48);
        let _524_v52;
        _524_v52 = _module.D3.create_DC8(!((_this).f25));
        (globalState).f2 = !((_524_v52).dtor_cf12) || (!(p1).isEqualTo(_517_v47));
      } else if (_source6.is_DC23) {
        let _525___mcc_h5 = (_source6).cf43;
        let _526___mcc_h6 = (_source6).cf44;
        let _527___mcc_h7 = (_source6).cf45;
        let _528___mcc_h8 = (_source6).cf46;
        let _529_cf46 = _528___mcc_h8;
        let _530_cf45 = _527___mcc_h7;
        let _531_cf44 = _526___mcc_h6;
        let _532_cf43 = _525___mcc_h5;
        let _533_v53;
        _533_v53 = _dafny.MultiSet.fromElements((_this).f25);
        _530_cf45 = (_533_v53).IsProperSubsetOf(_533_v53);
        if (!(_530_cf45)) {
          (globalState).f18 = ((_530_cf45) ? (_531_cf44) : ((_dafny.ZERO).minus(p1)));
          let _534_v54;
          _534_v54 = _module.D2.create_DC5(new BigNumber(-739), (_this).f25, _530_cf45, _module.__default.fm0(_531_cf44, _532_cf43, true, globalState));
          let _535_v55;
          _535_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_dafny.ZERO).minus(_531_cf44));
          let _536_v56;
          _536_v56 = _dafny.Map.Empty.slice().updateUnsafe(_530_cf45,(_this).f25);
          let _pat_let_tv9 = _530_cf45;
          let _537_v57;
          let _nw89 = Array((new BigNumber(18)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _534_v54;
          _nw89[(_dafny.ONE).toNumber()] = _534_v54;
          _nw89[(new BigNumber(2)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(3)).toNumber()] = _module.D2.create_DC5(new BigNumber((_535_v55).length), (_this).f25, (((_536_v56).contains((_this).f25)) ? ((_536_v56).get((_this).f25)) : ((_this).f25)), _530_cf45);
          _nw89[(new BigNumber(4)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(5)).toNumber()] = function (_pat_let4_0) {
            return function (_538_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_539_dt__update_hcf7_h0) {
                  return function (_pat_let6_0) {
                    return function (_540_dt__update_hcf8_h0) {
                      return _module.D2.create_DC5((_538_dt__update__tmp_h0).dtor_cf6, _539_dt__update_hcf7_h0, _540_dt__update_hcf8_h0, (_538_dt__update__tmp_h0).dtor_cf9);
                    }(_pat_let6_0);
                  }((_this).f25);
                }(_pat_let5_0);
              }(_pat_let_tv9);
            }(_pat_let4_0);
          }(_534_v54);
          _nw89[(new BigNumber(6)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(7)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(8)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(9)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(10)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(11)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(12)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(13)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(14)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(15)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(16)).toNumber()] = _534_v54;
          _nw89[(new BigNumber(17)).toNumber()] = _534_v54;
          _537_v57 = _nw89;
          let _index93 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_537_v57).length));
          (_537_v57)[_index93] = _534_v54;
          let _index94 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_537_v57).length));
          (_537_v57)[_index94] = _534_v54;
          r0 = (_this).f25;
          let _541_v58;
          let _nw90 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Set.Empty);
          _541_v58 = _nw90;
          let _542_v59;
          let _nw91 = new _module.C0();
          _nw91.__ctor();
          _542_v59 = _nw91;
          let _543_v60;
          _543_v60 = _dafny.Set.fromElements(_542_v59, _542_v59);
          let _index95 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_541_v58).length));
          (_541_v58)[_index95] = _543_v60;
          let _index96 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_541_v58).length));
          (_541_v58)[_index96] = _dafny.Set.fromElements(_542_v59);
          let _544_v61;
          let _nw92 = new _module.C0();
          _nw92.__ctor();
          _544_v61 = _nw92;
        } else {
          let _545_v62;
          _545_v62 = _dafny.Seq.of(_529_cf46, _529_cf46);
          let _546_v63;
          _546_v63 = _dafny.Map.Empty.slice().updateUnsafe(_529_cf46,_module.__default.safeModuloInt((_545_v62)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_545_v62).length))], p1));
          _546_v63 = (_546_v63).update((_529_cf46).minus(_module.__default.fm2(false, _531_cf44, _530_cf45, (_this).f26, globalState)), _529_cf46);
          let _547_v64;
          let _nw93 = new _module.C1();
          _nw93.__ctor(_dafny.Seq.UnicodeFromString("awhlibqi"), (_529_cf46).minus(p1), _530_cf45);
          _547_v64 = _nw93;
          _530_cf45 = _530_cf45;
          (globalState).f22 = (_module.__default.safeDivisionInt(_529_cf46, new BigNumber(160))).multipliedBy(_547_v64.f29);
          let _548_v66;
          _548_v66 = _dafny.Map.Empty.slice().updateUnsafe(_531_cf44,(_this).f25);
          let _549_v67;
          _549_v67 = _module.D6.create_DC17(function () {
  let _coll21 = new _dafny.Map();
  for (const _compr_21 of (_548_v66).Keys.Elements) {
    let _550_v65 = _compr_21;
    if ((_548_v66).contains(_550_v65)) {
      _coll21.push([(_550_v65).minus(_531_cf44),_531_cf44]);
    }
  }
  return _coll21;
}());
          let _551_v68;
          _551_v68 = _module.D6.create_DC19(_549_v67);
          _551_v68 = _551_v68;
        }
        let _552_v69;
        let _init17 = function (_553_i3) {
          return (_this).f26;
        };
        let _nw94 = Array((new BigNumber(17)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw94.length); _i0_17++) {
          _nw94[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _552_v69 = _nw94;
        let _index97 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_552_v69).length));
        (_552_v69)[_index97] = (_this).f26;
        let _index98 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_552_v69).length));
        (_552_v69)[_index98] = (_this).f26;
        let _554_v70;
        _554_v70 = _dafny.Set.fromElements(!(_530_cf45));
        let _index99 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_552_v69).length));
        let _index100 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_552_v69).length));
        let _rhs79 = (new BigNumber(169)).plus(_532_cf43);
        let _rhs80 = (_this).f26;
        let _rhs81 = new _dafny.CodePoint('n'.codePointAt(0));
        let _rhs82 = new BigNumber(((_554_v70).Intersect(_554_v70)).length);
        let _rhs83 = p1;
        let _lhs71 = globalState;
        let _lhs72 = _552_v69;
        let _lhs73 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_552_v69).length));
        let _lhs74 = _552_v69;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_552_v69).length));
        _lhs71.f22 = _rhs79;
        _lhs72[_lhs73] = _rhs80;
        _lhs74[_lhs75] = _rhs81;
        r1 = _rhs82;
        _531_cf44 = _rhs83;
        (globalState).f18 = (p1).plus(_529_cf46);
      } else if (_source6.is_DC20) {
        let _555___mcc_h9 = (_source6).cf37;
        let _556_cf37 = _555___mcc_h9;
        (globalState).f2 = (_this).f25;
        if (!(true)) {
          let _557_v71;
          _557_v71 = _dafny.MultiSet.fromElements(p1, p1, new BigNumber(988), p1, p1);
          (globalState).f2 = _module.__default.fm0((new BigNumber(610)).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(763)), ((_558_p1) => function (_559_i4) {
            return _558_p1;
          })(p1))).length)), _module.__default.fm2((_this).f25, (_dafny.ZERO).minus(new BigNumber(((_557_v71).update(new BigNumber(909), _module.__default.abs(p1))).cardinality())), !((_this).f25), (_this).f26, globalState), !(false), globalState);
          (globalState).f21 = p1;
          (globalState).f2 = (_this).f25;
          let _560_v72;
          let _init18 = function (_561_i5) {
            return (_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25));
          };
          let _nw95 = Array((new BigNumber(9)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw95.length); _i0_18++) {
            _nw95[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _560_v72 = _nw95;
          _560_v72 = _560_v72;
          let _562_v73;
          let _nw96 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _562_v73 = _nw96;
          let _index101 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_562_v73).length));
          (_562_v73)[_index101] = p1;
          let _index102 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_562_v73).length));
          (_562_v73)[_index102] = p1;
        } else {
          let _563_v74;
          let _init19 = function (_564_i6) {
            return !((_this).f25) || ((_this).f25);
          };
          let _nw97 = Array((new BigNumber(5)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw97.length); _i0_19++) {
            _nw97[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _563_v74 = _nw97;
          let _index103 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_563_v74).length));
          (_563_v74)[_index103] = (_this).f25;
          let _index104 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_563_v74).length));
          (_563_v74)[_index104] = (_this).f25;
          let _index105 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_563_v74).length));
          (_563_v74)[_index105] = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("hqccefcs"), _module.__default.fm16(p1, globalState));
          let _565_v75;
          _565_v75 = _dafny.Seq.of(p2, p2, _556_cf37, _dafny.Seq.Concat(_dafny.Seq.of((_this).f25), p2), p2);
          let _566_v76;
          _566_v76 = _module.D7.create_DC20(_556_cf37);
          let _567_v77;
          _567_v77 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_563_v74)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_563_v74).length))]);
          let _rhs84 = (_563_v74)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_563_v74).length))];
          let _rhs85 = (new BigNumber((p0).length)).minus((_dafny.ZERO).minus((p1).plus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_566_v76)).length))).length))));
          let _rhs86 = (((_this).f25) ? (new BigNumber((p0).length)) : (p1));
          let _rhs87 = _565_v75;
          let _rhs88 = new BigNumber((_567_v77).length);
          let _lhs76 = globalState;
          let _lhs77 = globalState;
          let _lhs78 = globalState;
          _lhs76.f2 = _rhs84;
          r1 = _rhs85;
          _lhs77.f22 = _rhs86;
          _565_v75 = _rhs87;
          _lhs78.f22 = _rhs88;
          _556_cf37 = p2;
          let _568_v78;
          _568_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p1);
          let _569_v79;
          _569_v79 = _dafny.MultiSet.fromElements((_563_v74)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_563_v74).length))], (_563_v74)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_563_v74).length))], (_563_v74)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_563_v74).length))]);
          let _570_v80;
          _570_v80 = _dafny.MultiSet.fromElements((_568_v78).Merge(_dafny.Map.Empty.slice().updateUnsafe((_563_v74)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_563_v74).length))],new BigNumber((_569_v79).cardinality()))));
          _570_v80 = (_570_v80).Union(_570_v80);
        }
        let _571_v81;
        _571_v81 = _dafny.Seq.of(p1);
        let _572_v82;
        _572_v82 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((p0).length)));
        (globalState).f2 = (_572_v82).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_571_v81, _571_v81)));
        let _573_v83;
        _573_v83 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(612),(_this).f26);
        let _574_v84;
        _574_v84 = _module.D5.create_DC16(_module.D5.create_DC14(_573_v83));
        let _575_v85;
        _575_v85 = _dafny.MultiSet.fromElements(p2, _556_cf37, _dafny.Seq.of((_this).f25));
        let _rhs89 = (((_this).f25) ? (_574_v84) : (_574_v84));
        let _rhs90 = (p1).isLessThanOrEqualTo(new BigNumber(((_575_v85).update(p2, _module.__default.abs(p1))).cardinality()));
        let _rhs91 = _492_v33;
        _574_v84 = _rhs89;
        r0 = _rhs90;
        _492_v33 = _rhs91;
      } else {
        let _576___mcc_h10 = (_source6).cf47;
        let _577_cf47 = _576___mcc_h10;
        let _578_v86;
        _578_v86 = _module.D2.create_DC5(p1, (_this).f25, !((_this).f25), (_this).f25);
        let _579_v87;
        _579_v87 = _module.D2.create_DC6(_578_v86);
        let _580_v88;
        _580_v88 = _module.D2.create_DC6(_579_v87);
        let _pat_let_tv10 = _579_v87;
        _580_v88 = function (_pat_let7_0) {
          return function (_581_dt__update__tmp_h1) {
            return function (_pat_let8_0) {
              return function (_582_dt__update_hcf10_h0) {
                return _module.D2.create_DC6(_582_dt__update_hcf10_h0);
              }(_pat_let8_0);
            }(_pat_let_tv10);
          }(_pat_let7_0);
        }(_580_v88);
        let _583_v89;
        let _init20 = function (_584_i7) {
          return (_584_i7).multipliedBy(new BigNumber(741));
        };
        let _nw98 = Array((new BigNumber(18)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw98.length); _i0_20++) {
          _nw98[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _583_v89 = _nw98;
        let _index106 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_583_v89).length));
        (_583_v89)[_index106] = p1;
        let _585_v90;
        _585_v90 = _module.D7.create_DC22(new BigNumber((p0).length), _492_v33, (_this).f25);
        let _index107 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_583_v89).length));
        (_583_v89)[_index107] = _module.__default.safeDivisionInt((_585_v90).dtor_cf40, p1);
        let _586_v91;
        _586_v91 = _dafny.MultiSet.fromElements(p1, p1, p1, p1);
        let _587_v92;
        _587_v92 = _dafny.Seq.of(p1, new BigNumber((_module.__default.fm24((_this).f25, p1, p1, globalState)).length));
        let _index108 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_583_v89).length));
        let _index109 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_583_v89).length));
        let _rhs92 = p1;
        let _rhs93 = p1;
        let _rhs94 = (_dafny.MultiSet.fromElements(new BigNumber((_587_v92).length), p1, p1)).IsProperSubsetOf(_586_v91);
        let _lhs79 = _583_v89;
        let _lhs80 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_583_v89).length));
        let _lhs81 = _583_v89;
        let _lhs82 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_583_v89).length));
        let _lhs83 = globalState;
        _lhs79[_lhs80] = _rhs92;
        _lhs81[_lhs82] = _rhs93;
        _lhs83.f2 = _rhs94;
        (globalState).f2 = (_this).f25;
        let _588_v93;
        _588_v93 = _dafny.Set.fromElements(!((_this).f25));
        let _589_v94;
        _589_v94 = _dafny.Map.Empty.slice().updateUnsafe(_588_v93,(_this).f25);
        _589_v94 = (_589_v94).update(_588_v93, !(p1).isEqualTo(p1));
      }
      let _590_v95;
      _590_v95 = _dafny.Seq.of(p1, p1, p1);
      let _hi5 = new BigNumber((_590_v95).length);
      for (let _591_i8 = p1; _591_i8.isLessThan(_hi5); _591_i8 = _591_i8.plus(_dafny.ONE)) {
        let _592_v96;
        let _nw99 = new _module.C0();
        _nw99.__ctor();
        _592_v96 = _nw99;
        let _593_v97;
        _593_v97 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,true);
        let _594_v98;
        _594_v98 = _dafny.Map.Empty.slice().updateUnsafe(!((((_593_v97).contains(false)) ? ((_593_v97).get(false)) : ((_this).f25))),(_this).f25);
        _594_v98 = (_594_v98).update(true, (_this).f25);
        let _595_v99;
        _595_v99 = _dafny.Seq.of(_590_v95);
        (globalState).f19 = _dafny.Seq.Concat(_590_v95, (_595_v99)[_module.__default.safeIndex(_591_i8, new BigNumber((_595_v99).length))]);
        let _596_v100;
        let _nw100 = Array((new BigNumber(7)).toNumber()).fill(false);
        _596_v100 = _nw100;
        let _index110 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_596_v100).length));
        (_596_v100)[_index110] = (p1).isEqualTo(p1);
        let _index111 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_596_v100).length));
        (_596_v100)[_index111] = true;
        let _index112 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_596_v100).length));
        let _index113 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_596_v100).length));
        let _rhs95 = (_this).f25;
        let _rhs96 = (_this).f25;
        let _rhs97 = (p2)[_module.__default.safeIndex(p1, new BigNumber((p2).length))];
        let _rhs98 = p1;
        let _rhs99 = _591_i8;
        let _lhs84 = _596_v100;
        let _lhs85 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_596_v100).length));
        let _lhs86 = _596_v100;
        let _lhs87 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_596_v100).length));
        let _lhs88 = globalState;
        let _lhs89 = globalState;
        r0 = _rhs95;
        _lhs84[_lhs85] = _rhs96;
        _lhs86[_lhs87] = _rhs97;
        _lhs88.f22 = _rhs98;
        _lhs89.f22 = _rhs99;
      }
      let _597_v101;
      _597_v101 = _dafny.MultiSet.fromElements(p1);
      let _598_v102;
      _598_v102 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
      let _599_v103;
      _599_v103 = _module.D7.create_DC21((_dafny.ZERO).minus(p1), _598_v102);
      let _600_v104;
      let _nw101 = Array((new BigNumber(24)).toNumber());
      _nw101[(_dafny.ZERO).toNumber()] = (_this).f25;
      _nw101[(_dafny.ONE).toNumber()] = !(p1).isEqualTo(new BigNumber(-887));
      _nw101[(new BigNumber(2)).toNumber()] = (_597_v101).IsSubsetOf((_597_v101).update(p1, _module.__default.abs((_599_v103).dtor_cf38)));
      _nw101[(new BigNumber(3)).toNumber()] = !((_this).f25);
      _nw101[(new BigNumber(4)).toNumber()] = ((_dafny.ZERO).minus(p1)).isLessThan(p1);
      _nw101[(new BigNumber(5)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(6)).toNumber()] = (!((p2)[_module.__default.safeIndex(p1, new BigNumber((p2).length))])) && ((_this).f25);
      _nw101[(new BigNumber(7)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(8)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(9)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(10)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(11)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(12)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(13)).toNumber()] = (p1).isEqualTo(p1);
      _nw101[(new BigNumber(14)).toNumber()] = !_dafny.Seq.contains(p0, (_this).f26);
      _nw101[(new BigNumber(15)).toNumber()] = (_597_v101).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-374)), function (_601_i9) {
        return new BigNumber(254);
      })));
      _nw101[(new BigNumber(16)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(17)).toNumber()] = (p2)[_module.__default.safeIndex(p1, new BigNumber((p2).length))];
      _nw101[(new BigNumber(18)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(19)).toNumber()] = false;
      _nw101[(new BigNumber(20)).toNumber()] = (((_this).f25) ? ((_this).f25) : ((_this).f25));
      _nw101[(new BigNumber(21)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(22)).toNumber()] = (_this).f25;
      _nw101[(new BigNumber(23)).toNumber()] = (_this).f25;
      _600_v104 = _nw101;
      let _index114 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length));
      (_600_v104)[_index114] = (_this).f25;
      let _602_v105;
      _602_v105 = _dafny.Set.fromElements((_this).f26, (_this).f26, _module.__default.fm16(p1, globalState));
      let _index115 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length));
      (_600_v104)[_index115] = !(_module.__default.safeDivisionInt(new BigNumber(-284), p1)).isEqualTo(new BigNumber(((_602_v105).Difference(_602_v105)).length));
      let _603_i10;
      _603_i10 = _dafny.ZERO;
      L2: {
        while ((_this).f25) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_603_i10)) {
              break L2;
            }
            _603_i10 = (_603_i10).plus(_dafny.ONE);
            let _604_v106;
            let _init21 = ((_605_p1) => function (_606_i11) {
              return _module.__default.safeModuloInt(_606_i11, _605_p1);
            })(p1);
            let _nw102 = Array((new BigNumber(19)).toNumber());
            for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw102.length); _i0_21++) {
              _nw102[_i0_21] = _init21(new BigNumber(_i0_21));
            }
            _604_v106 = _nw102;
            let _607_v107;
            _607_v107 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(505)), ((_608_p0, _609_v105) => function (_610_i12) {
              return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_608_p0,new BigNumber((_609_v105).length))).length),(_this).f25)).length);
            })(p0, _602_v105)));
            let _611_v108;
            _611_v108 = _module.D7.create_DC23(new BigNumber(934), p1, (_600_v104)[_module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length))], p1);
            let _612_v109;
            _612_v109 = _dafny.Map.Empty.slice().updateUnsafe(_590_v95,_611_v108);
            let _613_v111;
            _613_v111 = _dafny.Seq.of(_600_v104);
            let _index116 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length));
            let _rhs100 = (function () {
              let _coll22 = new _dafny.Set();
              for (const _compr_22 of (_612_v109).Keys.Elements) {
                let _614_v110 = _compr_22;
                if ((_612_v109).contains(_614_v110)) {
                  _coll22.add(_614_v110);
                }
              }
              return _coll22;
            }()).IsSubsetOf(_607_v107);
            let _rhs101 = (_600_v104)[_module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length))];
            let _rhs102 = (_613_v111)[_module.__default.safeIndex(p1, new BigNumber((_613_v111).length))];
            let _rhs103 = (_this).f25;
            let _rhs104 = _604_v106;
            let _lhs90 = _600_v104;
            let _lhs91 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length));
            _lhs90[_lhs91] = _rhs100;
            r0 = _rhs101;
            _600_v104 = _rhs102;
            r0 = _rhs103;
            _604_v106 = _rhs104;
            let _615_v112;
            _615_v112 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f25);
            let _616_v113;
            _616_v113 = _dafny.Set.fromElements(_module.__default.fm2((_this).f25, new BigNumber(-283), (_600_v104)[_module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length))], (_this).f26, globalState));
            let _617_v114;
            let _nw103 = Array((new BigNumber(6)).toNumber());
            _nw103[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(p1, _module.__default.fm2((_600_v104)[_module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length))], new BigNumber((p2).length), (_600_v104)[_module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length))], new _dafny.CodePoint('m'.codePointAt(0)), globalState), p1);
            _nw103[(_dafny.ONE).toNumber()] = _module.__default.fm24((_600_v104)[_module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length))], new BigNumber((_615_v112).length), p1, globalState);
            _nw103[(new BigNumber(2)).toNumber()] = _616_v113;
            _nw103[(new BigNumber(3)).toNumber()] = (_616_v113).Union(_dafny.Set.fromElements(p1, (_dafny.ZERO).minus(p1)));
            _nw103[(new BigNumber(4)).toNumber()] = _616_v113;
            _nw103[(new BigNumber(5)).toNumber()] = (_616_v113).Union(_dafny.Set.fromElements(_module.__default.fm2((_this).f25, p1, (_this).f25, (_this).f26, globalState), new BigNumber((p2).length)));
            _617_v114 = _nw103;
            _617_v114 = _617_v114;
            (globalState).f18 = ((true) ? (((_dafny.ZERO).minus(p1)).minus(p1)) : ((p1).minus(new BigNumber(-111))));
            let _index117 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_600_v104).length));
            (_600_v104)[_index117] = true;
          }
        }
      }
      let _618_v115;
      _618_v115 = _module.D3.create_DC9(p1, (_this).f26, p1);
      let _619_v116;
      _619_v116 = _dafny.MultiSet.fromElements(_dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_590_v95).length), new BigNumber((p0).length)), (_618_v115).dtor_cf14), p0);
      r0 = ((_619_v116).update(_dafny.Seq.UnicodeFromString("v"), _module.__default.abs(new BigNumber(764)))).IsProperSubsetOf(_619_v116);
      r0 = (_this).f25;
      r1 = p1;
      return [r0, r1];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f26 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f27 = [];
      this._f25 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f26, f27, f25) {
      let _this = this;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      (_this)._f25 = f25;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f25;
    };
    fm4(p0, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(((_dafny.ZERO).minus(new BigNumber(-836))).multipliedBy(new BigNumber((_dafny.Seq.of((_this).f25, (_this).f25)).length)), new BigNumber(314));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let _620_v0;
      let _nw104 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
      _620_v0 = _nw104;
      let _index118 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length));
      (_620_v0)[_index118] = _module.__default.safeModuloInt(p2, p2);
      let _index119 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length));
      (_620_v0)[_index119] = _module.__default.safeDivisionInt(p2, _module.__default.safeDivisionInt(_module.__default.fm2((_this).f25, _module.__default.fm2((_this).f25, p2, false, new _dafny.CodePoint('e'.codePointAt(0)), globalState), p1, (_this).f26, globalState), new BigNumber((function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(94)), ((_621_p0) => function (_622_i0) {
          return _dafny.Seq.of(new BigNumber((_621_p0).length));
        })(p0))).Elements) {
          let _623_v1 = _compr_23;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(94)), ((_624_p0) => function (_622_i0) {
            return _dafny.Seq.of(new BigNumber((_624_p0).length));
          })(p0)), _623_v1)) {
            _coll23.push([_623_v1,new BigNumber((_dafny.Seq.UnicodeFromString("cela")).length)]);
          }
        }
        return _coll23;
      }()).length)));
      let _625_v2;
      _625_v2 = _dafny.Seq.UnicodeFromString("pgjv");
      _625_v2 = _625_v2;
      let _626_v3;
      let _nw105 = new _module.C0();
      _nw105.__ctor();
      _626_v3 = _nw105;
      let _627_v4;
      let _nw106 = new _module.C0();
      _nw106.__ctor();
      _627_v4 = _nw106;
      (globalState).f2 = ((_this).f25) === (p1);
      if (p1) {
        (globalState).f18 = _module.__default.safeDivisionInt(p2, p2);
        let _628_v5;
        _628_v5 = _dafny.Set.fromElements((_this).f25);
        (globalState).f5 = new BigNumber((_628_v5).length);
        (globalState).f22 = ((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))]).plus((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))]);
        if (p1) {
          let _629_v6;
          let _nw107 = Array((new BigNumber(21)).toNumber());
          _629_v6 = _nw107;
          let _index120 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_629_v6).length));
          (_629_v6)[_index120] = _626_v3;
          let _index121 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_629_v6).length));
          (_629_v6)[_index121] = _626_v3;
          let _630_v7;
          let _nw108 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _630_v7 = _nw108;
          let _631_v8;
          _631_v8 = _dafny.Seq.of(_625_v2);
          let _index122 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_630_v7).length));
          (_630_v7)[_index122] = _631_v8;
          let _index123 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_630_v7).length));
          (_630_v7)[_index123] = _631_v8;
          (globalState).f2 = !(p1);
          _625_v2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_625_v2, _625_v2), _625_v2);
          _625_v2 = _dafny.Seq.Concat(_625_v2, _dafny.Seq.UnicodeFromString("rr"));
        } else {
          (globalState).f21 = (_dafny.ZERO).minus((_module.__default.safeModuloInt((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))], p2)).minus(new BigNumber(495)));
          (globalState).f18 = p2;
          let _632_v9;
          _632_v9 = _dafny.Map.Empty.slice().updateUnsafe((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))],p2);
          _632_v9 = (_632_v9).update(p2, (p2).minus(new BigNumber(319)));
          let _633_v10;
          let _nw109 = Array((new BigNumber(9)).toNumber()).fill([]);
          _633_v10 = _nw109;
          let _634_v11;
          let _init22 = ((_635_v2) => function (_636_i1) {
            return _635_v2;
          })(_625_v2);
          let _nw110 = Array((new BigNumber(6)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw110.length); _i0_22++) {
            _nw110[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _634_v11 = _nw110;
          let _index124 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_633_v10).length));
          (_633_v10)[_index124] = _634_v11;
          let _index125 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_633_v10).length));
          (_633_v10)[_index125] = _634_v11;
          let _637_v12;
          let _nw111 = Array((new BigNumber(4)).toNumber());
          _637_v12 = _nw111;
          let _638_v13;
          _638_v13 = _dafny.Seq.of(_637_v12);
          _637_v12 = (_638_v13)[_module.__default.safeIndex((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))], new BigNumber((_638_v13).length))];
        }
        let _639_v14;
        _639_v14 = p2;
        let _640_v15;
        _640_v15 = _dafny.Seq.of(p2, (_639_v14));
        let _641_v16;
        _641_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-107),p1);
        let _642_v17;
        _642_v17 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber((_640_v15).length)).isEqualTo((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))]),_641_v16);
        _642_v17 = (_642_v17).update(true, _641_v16);
      } else {
        let _643_v18;
        _643_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f26);
        (globalState).f2 = _module.__default.fm0((_dafny.ZERO).minus(new BigNumber((_643_v18).length)), (_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))], (_this).f25, globalState);
        (globalState).f18 = (_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))];
        (globalState).f2 = (_this).f25;
        let _index126 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length));
        (_620_v0)[_index126] = p2;
        let _644_v19;
        _644_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f25);
        let _645_v20;
        _645_v20 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _646_v21;
        _646_v21 = _dafny.Set.fromElements(p1, p1, (((_645_v20).contains(!(!((_this).f25)))) ? ((_645_v20).get(!(!((_this).f25)))) : (p1)), false, p1);
        let _647_v22;
        _647_v22 = _module.D1.create_DC2(p1);
        let _648_v23;
        _648_v23 = _module.D1.create_DC1((_this).f25);
        let _649_v24;
        _649_v24 = _dafny.Seq.of((_648_v23).dtor_cf1);
        let _650_v25;
        _650_v25 = _dafny.Set.fromElements(_dafny.Seq.update(_649_v24, _module.__default.safeIndex(p2, new BigNumber((_649_v24).length)), p1));
        let _651_v26;
        let _nw112 = Array((new BigNumber(16)).toNumber());
        _nw112[(_dafny.ZERO).toNumber()] = p1;
        _nw112[(_dafny.ONE).toNumber()] = (_this).f25;
        _nw112[(new BigNumber(2)).toNumber()] = (_this).f25;
        _nw112[(new BigNumber(3)).toNumber()] = !((_this).f25);
        _nw112[(new BigNumber(4)).toNumber()] = ((_dafny.ZERO).minus((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))])).isLessThanOrEqualTo((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))]);
        _nw112[(new BigNumber(5)).toNumber()] = false;
        _nw112[(new BigNumber(6)).toNumber()] = (p2).isEqualTo(p2);
        _nw112[(new BigNumber(7)).toNumber()] = (((_644_v19).contains((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))])) ? ((_644_v19).get((_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))])) : ((_this).f25));
        _nw112[(new BigNumber(8)).toNumber()] = (_this).fm3((_this).f25, new BigNumber((_646_v21).length), _module.__default.fm2((_this).f25, p2, (_this).fm3(p1, (_620_v0)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_620_v0).length))], new BigNumber(-711), _dafny.Map.Empty.slice().updateUnsafe((_this).f26,p2), globalState), (_this).f26, globalState), _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber(926)), globalState);
        _nw112[(new BigNumber(9)).toNumber()] = (_this).f25;
        _nw112[(new BigNumber(10)).toNumber()] = (_this).f25;
        _nw112[(new BigNumber(11)).toNumber()] = (_647_v22).dtor_cf2;
        _nw112[(new BigNumber(12)).toNumber()] = (_this).f25;
        _nw112[(new BigNumber(13)).toNumber()] = p1;
        _nw112[(new BigNumber(14)).toNumber()] = (_650_v25).IsProperSubsetOf(_650_v25);
        _nw112[(new BigNumber(15)).toNumber()] = p1;
        _651_v26 = _nw112;
        _651_v26 = _651_v26;
      }
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _652_v0;
      _652_v0 = _dafny.MultiSet.fromElements(p1);
      let _653_v1;
      _653_v1 = _dafny.MultiSet.fromElements(p1, new BigNumber(149), new BigNumber((_652_v0).cardinality()), p1);
      let _654_i0;
      _654_i0 = _dafny.ZERO;
      L3: {
        while (((_652_v0).Union(_652_v0)).IsProperSubsetOf(_653_v1)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_654_i0)) {
              break L3;
            }
            _654_i0 = (_654_i0).plus(_dafny.ONE);
            let _655_v2;
            let _nw113 = Array((new BigNumber(14)).toNumber());
            _nw113[(_dafny.ZERO).toNumber()] = p0;
            _nw113[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), function (_656_i1) {
              return (_this).f26;
            }));
            _nw113[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-96)), function (_657_i2) {
              return (_this).f26;
            });
            _nw113[(new BigNumber(3)).toNumber()] = p0;
            _nw113[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(p0, _module.__default.safeIndex(p1, new BigNumber((p0).length)), (_this).f26);
            _nw113[(new BigNumber(5)).toNumber()] = p0;
            _nw113[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(p0, _dafny.Seq.UnicodeFromString("dhnjxv"));
            _nw113[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(469)), function (_658_i3) {
              return (_this).f26;
            });
            _nw113[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("nyxcrvbu");
            _nw113[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(992)), function (_659_i4) {
              return (_this).f26;
            }), _dafny.Seq.UnicodeFromString("imrgbtikj"));
            _nw113[(new BigNumber(10)).toNumber()] = p0;
            _nw113[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dewtel"), _dafny.Seq.UnicodeFromString("fcxwj"));
            _nw113[(new BigNumber(12)).toNumber()] = (((_this).f25) ? (p0) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(304)), function (_660_i5) {
              return (_this).f26;
            })));
            _nw113[(new BigNumber(13)).toNumber()] = p0;
            _655_v2 = _nw113;
            let _index127 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length));
            (_655_v2)[_index127] = p0;
            let _index128 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length));
            (_655_v2)[_index128] = _dafny.Seq.update(p0, _module.__default.safeIndex(p1, new BigNumber((p0).length)), (_this).f26);
            let _661_v3;
            _661_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_655_v2)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length))]).length),p1);
            let _662_v4;
            _662_v4 = _dafny.Seq.of(_661_v3);
            let _663_v5;
            _663_v5 = _module.D2.create_DC3(_662_v4);
            let _664_v6;
            _664_v6 = _dafny.MultiSet.fromElements((_this).f25);
            let _665_v7;
            _665_v7 = _dafny.Map.Empty.slice().updateUnsafe(_664_v6,_662_v4);
            (globalState).f21 = new BigNumber((_dafny.Seq.Concat((_663_v5).dtor_cf3, (((_665_v7).contains(_664_v6)) ? ((_665_v7).get(_664_v6)) : (_dafny.Seq.of(_661_v3))))).length);
            let _666_v8;
            _666_v8 = _dafny.Seq.of(p0, _dafny.Seq.update(_dafny.Seq.Concat((_655_v2)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length))], (_655_v2)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length))]), _module.__default.safeIndex(new BigNumber(5), new BigNumber((_dafny.Seq.Concat((_655_v2)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length))], (_655_v2)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length))])).length)), (_this).f26), (_655_v2)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length))], _dafny.Seq.Concat(_module.__default.fm7((_this).f26, globalState), _dafny.Seq.UnicodeFromString("h")), (_655_v2)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length))]);
            let _index129 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_655_v2).length));
            (_655_v2)[_index129] = (_666_v8)[_module.__default.safeIndex(p1, new BigNumber((_666_v8).length))];
            let _667_v9;
            _667_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_module.__default.fm2((_this).f25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(887)), function (_668_i6) {
              return (_this).f26;
            })).length), (_this).f25, new _dafny.CodePoint('g'.codePointAt(0)), globalState));
            let _669_v10;
            let _init23 = function (_670_i7) {
              return true;
            };
            let _nw114 = Array((new BigNumber(4)).toNumber());
            for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw114.length); _i0_23++) {
              _nw114[_i0_23] = _init23(new BigNumber(_i0_23));
            }
            _669_v10 = _nw114;
            let _index130 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_669_v10).length));
            (_669_v10)[_index130] = (_this).f25;
            let _671_v11;
            _671_v11 = _module.D1.create_DC2((_this).f25);
            let _index131 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_669_v10).length));
            (_669_v10)[_index131] = (_671_v11).dtor_cf2;
            let _index132 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_669_v10).length));
            let _index133 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_669_v10).length));
            let _rhs105 = _667_v9;
            let _rhs106 = (_this).f25;
            let _rhs107 = _module.__default.fm0((p1).plus(p1), (p1).minus(p1), (_this).f25, globalState);
            let _rhs108 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(p1))).minus(p1);
            let _rhs109 = true;
            let _lhs92 = globalState;
            let _lhs93 = _669_v10;
            let _lhs94 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_669_v10).length));
            let _lhs95 = _669_v10;
            let _lhs96 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_669_v10).length));
            _667_v9 = _rhs105;
            _lhs92.f2 = _rhs106;
            _lhs93[_lhs94] = _rhs107;
            r1 = _rhs108;
            _lhs95[_lhs96] = _rhs109;
          }
        }
      }
      let _672_v12;
      let _init24 = ((_673_p1) => function (_674_i8) {
        return _module.__default.safeDivisionInt(_674_i8, _673_p1);
      })(p1);
      let _nw115 = Array((new BigNumber(3)).toNumber());
      for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw115.length); _i0_24++) {
        _nw115[_i0_24] = _init24(new BigNumber(_i0_24));
      }
      _672_v12 = _nw115;
      let _index134 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length));
      (_672_v12)[_index134] = new BigNumber(249);
      let _index135 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length));
      (_672_v12)[_index135] = p1;
      if ((_this).f25) {
        r0 = (_this).f25;
        let _675_v13;
        _675_v13 = new _dafny.CodePoint('w'.codePointAt(0));
        let _676_v14;
        _676_v14 = _dafny.Set.fromElements(p0, _dafny.Seq.Concat(p0, p0), p0);
        let _677_v15;
        _677_v15 = _dafny.Seq.of((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]);
        let _678_v16;
        _678_v16 = new BigNumber((_dafny.Seq.of((_this).f25)).length);
        let _679_v17;
        _679_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_678_v16);
        let _rhs110 = !(_module.__default.fm0((_dafny.ZERO).minus(_module.__default.safeModuloInt((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))], (_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))])), (_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))], (_this).f25, globalState));
        let _rhs111 = (_this).f26;
        let _rhs112 = ((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]).isEqualTo((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]);
        let _rhs113 = _module.__default.safeModuloInt(new BigNumber((_677_v15).length), _module.__default.safeModuloInt((_677_v15)[_module.__default.safeIndex(new BigNumber((_679_v17).length), new BigNumber((_677_v15).length))], new BigNumber((_dafny.Seq.update(p0, _module.__default.safeIndex(p1, new BigNumber((p0).length)), (_this).f26)).length)));
        let _rhs114 = _676_v14;
        let _lhs97 = globalState;
        let _lhs98 = globalState;
        _lhs97.f2 = _rhs110;
        _675_v13 = _rhs111;
        r0 = _rhs112;
        _lhs98.f21 = _rhs113;
        _676_v14 = _rhs114;
        let _index136 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length));
        (_672_v12)[_index136] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(609),(_this).f25)).length)), new BigNumber((p0).length)));
        let _680_v18;
        let _nw116 = Array((new BigNumber(3)).toNumber()).fill([]);
        _680_v18 = _nw116;
        let _681_v19;
        let _nw117 = Array((new BigNumber(18)).toNumber());
        _nw117[(_dafny.ZERO).toNumber()] = _680_v18;
        _nw117[(_dafny.ONE).toNumber()] = _680_v18;
        _nw117[(new BigNumber(2)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(3)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(4)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(5)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(6)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(7)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(8)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(9)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(10)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(11)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(12)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(13)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(14)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(15)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(16)).toNumber()] = _680_v18;
        _nw117[(new BigNumber(17)).toNumber()] = (((_this).f25) ? (_680_v18) : (_680_v18));
        _681_v19 = _nw117;
        let _index137 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_681_v19).length));
        (_681_v19)[_index137] = _680_v18;
        let _index138 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_681_v19).length));
        (_681_v19)[_index138] = _680_v18;
        (globalState).f5 = (p1).plus(p1);
      } else {
        let _682_v20;
        _682_v20 = _dafny.Set.fromElements(new BigNumber(586), p1);
        let _683_v21;
        _683_v21 = _dafny.Seq.of(new BigNumber((_682_v20).length));
        if (((!((new BigNumber((_683_v21).length)).isLessThan((_dafny.ZERO).minus(new BigNumber((p2).length))))) ? (!((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]).isEqualTo((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))])) : (!((_this).f25) || (false)))) {
          let _index139 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length));
          (_672_v12)[_index139] = (_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))];
          r0 = (new BigNumber(496)).isLessThanOrEqualTo((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]);
          (globalState).f18 = new BigNumber((p0).length);
          (globalState).f2 = false;
          let _684_v22;
          _684_v22 = new _dafny.CodePoint('t'.codePointAt(0));
          let _rhs115 = _684_v22;
          let _rhs116 = (_this).f25;
          let _lhs99 = globalState;
          _684_v22 = _rhs115;
          _lhs99.f2 = _rhs116;
        } else {
          let _index140 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length));
          (_672_v12)[_index140] = p1;
          (globalState).f2 = (_this).f25;
          let _685_v23;
          let _nw118 = new _module.C2();
          _nw118.__ctor((_this).f26, (_this).f27, _module.__default.fm0((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))], (_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))], (_this).f25, globalState));
          _685_v23 = _nw118;
          let _686_v24;
          let _init25 = function (_687_i9) {
            return (_this).f25;
          };
          let _nw119 = Array((new BigNumber(3)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw119.length); _i0_25++) {
            _nw119[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _686_v24 = _nw119;
          (globalState).f5 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_685_v23,_686_v24)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_685_v23,_686_v24))).length);
          let _688_v25;
          _688_v25 = _dafny.MultiSet.fromElements(true, (_this).f25);
          let _index141 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length));
          (_672_v12)[_index141] = (((_685_v23).f25) ? (((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]).plus((((_688_v25).contains((_this).f25)) ? ((_688_v25).get((_this).f25)) : ((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))])))) : (p1));
          let _689_v26;
          let _nw120 = Array((new BigNumber(25)).toNumber());
          _689_v26 = _nw120;
          let _index142 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_689_v26).length));
          (_689_v26)[_index142] = _685_v23;
          let _index143 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_689_v26).length));
          let _rhs117 = (_this).f25;
          let _rhs118 = _685_v23;
          let _lhs100 = globalState;
          let _lhs101 = _689_v26;
          let _lhs102 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_689_v26).length));
          _lhs100.f2 = _rhs117;
          _lhs101[_lhs102] = _rhs118;
        }
        let _690_v27;
        _690_v27 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(435)), function (_691_i10) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f25)).length);
        }), _dafny.Seq.of((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]), _683_v21, _683_v21);
        if ((_690_v27).IsDisjointFrom(_690_v27)) {
          (globalState).f5 = p1;
          (globalState).f18 = ((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]).plus((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]);
          let _692_v28;
          _692_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f25);
          let _693_v29;
          _693_v29 = _module.D3.create_DC7(_692_v28);
          let _694_v30;
          let _nw121 = Array((new BigNumber(7)).toNumber());
          _nw121[(_dafny.ZERO).toNumber()] = _module.__default.fm25(globalState);
          _nw121[(_dafny.ONE).toNumber()] = _module.__default.fm25(globalState);
          _nw121[(new BigNumber(2)).toNumber()] = _module.D3.create_DC7(_692_v28);
          _nw121[(new BigNumber(3)).toNumber()] = _module.__default.fm25(globalState);
          _nw121[(new BigNumber(4)).toNumber()] = _693_v29;
          _nw121[(new BigNumber(5)).toNumber()] = _693_v29;
          _nw121[(new BigNumber(6)).toNumber()] = _693_v29;
          _694_v30 = _nw121;
          let _index144 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_694_v30).length));
          (_694_v30)[_index144] = _693_v29;
          let _pat_let_tv11 = _692_v28;
          let _index145 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_694_v30).length));
          (_694_v30)[_index145] = function (_pat_let9_0) {
            return function (_695_dt__update__tmp_h1) {
              return function (_pat_let10_0) {
                return function (_696_dt__update_hcf11_h0) {
                  return _module.D3.create_DC7(_696_dt__update_hcf11_h0);
                }(_pat_let10_0);
              }(_pat_let_tv11);
            }(_pat_let9_0);
          }(_693_v29);
          let _697_v31;
          _697_v31 = _dafny.Map.Empty.slice().updateUnsafe(_683_v21,p0);
          _697_v31 = (_697_v31).update(_683_v21, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yjmlwg"), _dafny.Seq.UnicodeFromString("x")));
          let _698_v32;
          let _nw122 = new _module.C2();
          _nw122.__ctor((_this).f26, (_this).f27, (_this).f25);
          _698_v32 = _nw122;
        } else {
          (globalState).f5 = new BigNumber(807);
          (globalState).f5 = p1;
          let _699_v33;
          _699_v33 = _dafny.Set.fromElements((_this).f25);
          (globalState).f18 = new BigNumber((_699_v33).length);
          let _700_v34;
          let _nw123 = new _module.C2();
          _nw123.__ctor((_this).f26, (_this).f27, (_this).f25);
          _700_v34 = _nw123;
          let _rhs119 = (_this).f25;
          let _rhs120 = (((_this).f25) ? (_700_v34) : (_700_v34));
          let _lhs103 = globalState;
          _lhs103.f2 = _rhs119;
          _700_v34 = _rhs120;
          (globalState).f2 = (_this).f25;
        }
        (globalState).f2 = (_this).f25;
        let _701_v35;
        let _init26 = function (_702_i11) {
          return false;
        };
        let _nw124 = Array((new BigNumber(29)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw124.length); _i0_26++) {
          _nw124[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _701_v35 = _nw124;
        let _index146 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_701_v35).length));
        (_701_v35)[_index146] = false;
        let _index147 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_701_v35).length));
        (_701_v35)[_index147] = ((_this).f25) || (!((_this).f25));
        let _703_v36;
        _703_v36 = _module.D5.create_DC15((_this).f25, _701_v35, p0, false);
        let _source7 = _703_v36;
        if (_source7.is_DC15) {
          let _704___mcc_h0 = (_source7).cf25;
          let _705___mcc_h1 = (_source7).cf26;
          let _706___mcc_h2 = (_source7).cf27;
          let _707___mcc_h3 = (_source7).cf28;
          let _708_cf28 = _707___mcc_h3;
          let _709_cf27 = _706___mcc_h2;
          let _710_cf26 = _705___mcc_h1;
          let _711_cf25 = _704___mcc_h0;
          (globalState).f18 = p1;
          (globalState).f22 = _module.__default.fm2((_701_v35)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_701_v35).length))], (_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))], (_this).f25, new _dafny.CodePoint('f'.codePointAt(0)), globalState);
          let _712_v37;
          let _nw125 = new _module.C2();
          _nw125.__ctor((_this).f26, (_this).f27, _711_cf25);
          _712_v37 = _nw125;
          let _713_v38;
          _713_v38 = _module.D3.create_DC9(p1, new _dafny.CodePoint('r'.codePointAt(0)), p1);
          let _714_v39;
          _714_v39 = _module.D3.create_DC10(_713_v38);
          let _715_v40;
          _715_v40 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f25);
          let _pat_let_tv12 = _715_v40;
          _714_v39 = function (_pat_let11_0) {
            return function (_716_dt__update__tmp_h2) {
              return function (_pat_let12_0) {
                return function (_717_dt__update_hcf16_h0) {
                  return _module.D3.create_DC10(_717_dt__update_hcf16_h0);
                }(_pat_let12_0);
              }(_module.D3.create_DC7(_pat_let_tv12));
            }(_pat_let11_0);
          }(_714_v39);
        } else if (_source7.is_DC14) {
          let _718___mcc_h4 = (_source7).cf24;
          let _719_cf24 = _718___mcc_h4;
          let _720_v41;
          _720_v41 = _module.D2.create_DC5(new BigNumber(204), true, (_701_v35)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_701_v35).length))], (_this).f25);
          let _721_v42;
          _721_v42 = _dafny.Map.Empty.slice().updateUnsafe((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))],_701_v35);
          let _722_v43;
          _722_v43 = _dafny.Map.Empty.slice().updateUnsafe((_720_v41).dtor_cf6,_721_v42);
          let _index148 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length));
          (_672_v12)[_index148] = (((_701_v35)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_701_v35).length))]) ? ((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]) : (new BigNumber(((((_722_v43).contains(p1)) ? ((_722_v43).get(p1)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(171),_701_v35)))).length)));
          (globalState).f18 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((new BigNumber(335)).plus((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]), ((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]).minus((_dafny.ZERO).minus((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]))));
          let _723_v44;
          let _nw126 = new _module.C2();
          _nw126.__ctor((_this).f26, (_this).f27, (_this).f25);
          _723_v44 = _nw126;
          let _724_v45;
          _724_v45 = _dafny.Map.Empty.slice().updateUnsafe(_723_v44,_dafny.Seq.Concat(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), function (_725_i12) {
            return (_this).f26;
          })));
          _724_v45 = _724_v45;
          let _726_v46;
          _726_v46 = _dafny.Seq.UnicodeFromString("if");
          _726_v46 = _dafny.Seq.Concat(_dafny.Seq.Concat(p0, _dafny.Seq.UnicodeFromString("xdalhwjvr")), _726_v46);
        } else {
          let _727___mcc_h5 = (_source7).cf29;
          let _728_cf29 = _727___mcc_h5;
          (globalState).f2 = !((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]).isEqualTo((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]);
          (globalState).f5 = (_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))];
          let _729_v47;
          _729_v47 = _dafny.Map.Empty.slice().updateUnsafe((_701_v35)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_701_v35).length))],p0);
          (globalState).f5 = (_dafny.ZERO).minus((_683_v21)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_729_v47).length)), new BigNumber((_683_v21).length))]);
          let _index149 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_701_v35).length));
          (_701_v35)[_index149] = !((_652_v0).update(p1, _module.__default.abs((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]))).contains(p1);
        }
      }
      r1 = (((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]).multipliedBy((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))])).plus(_module.__default.safeDivisionInt(p1, new BigNumber((p0).length)));
      r1 = (_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))];
      r0 = (_this).f25;
      let _730_v49;
      _730_v49 = _module.D8.create_DC25(_653_v1);
      let _731_v50;
      _731_v50 = _module.D6.create_DC17(function () {
  let _coll24 = new _dafny.Map();
  for (const _compr_24 of (((_730_v49).dtor_cf48).update(p1, _module.__default.abs((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]))).Elements) {
    let _732_v48 = _compr_24;
    if ((((_730_v49).dtor_cf48).update(p1, _module.__default.abs((_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))]))).contains(_732_v48)) {
      _coll24.push([_module.__default.safeModuloInt(_732_v48, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(66)), function (_733_i13) {
        return (_this).f26;
      })).length)),p1]);
    }
  }
  return _coll24;
}());
      r0 = function (_source8) {
        if (_source8.is_DC18) {
          let _734___mcc_h6 = (_source8).cf31;
          let _735___mcc_h7 = (_source8).cf32;
          let _736___mcc_h8 = (_source8).cf33;
          let _737___mcc_h9 = (_source8).cf34;
          let _738___mcc_h10 = (_source8).cf35;
          let _739_cf35 = _738___mcc_h10;
          let _740_cf34 = _737___mcc_h9;
          let _741_cf33 = _736___mcc_h8;
          let _742_cf32 = _735___mcc_h7;
          let _743_cf31 = _734___mcc_h6;
          return (_741_cf33) || (_740_cf34);
        } else if (_source8.is_DC17) {
          let _744___mcc_h11 = (_source8).cf30;
          let _745_cf30 = _744___mcc_h11;
          return true;
        } else {
          let _746___mcc_h12 = (_source8).cf36;
          let _747_cf36 = _746___mcc_h12;
          return (_this).f25;
        }
      }(_731_v50);
      r1 = (_672_v12)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_672_v12).length))];
      return [r0, r1];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
