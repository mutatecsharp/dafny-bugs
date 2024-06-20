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
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(874),true), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, false)).length)),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(773),!(true))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(431),false), function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(160), new BigNumber(497))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(160)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(497)))) {
            _coll0.push([(_0_v0).plus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false, false, true, true))).cardinality())),false]);
          }
        }
        return _coll0;
      }()), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(416),true), function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(926)), function (_1_i0) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        })).length),!(true))).Keys.Elements) {
          let _2_v1 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(926)), function (_1_i0) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          })).length),!(true))).contains(_2_v1)) {
            _coll1.push([_module.__default.safeModuloInt(_2_v1, new BigNumber((_dafny.Seq.of(new BigNumber(334), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(-137), new BigNumber(864), new BigNumber(242))).length)),true]);
          }
        }
        return _coll1;
      }())));
    };
    static fm1(p0, p1, globalState) {
      return (new BigNumber(((_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(!(true)))).length)).plus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(583))).length),false)).length)).multipliedBy(new BigNumber(-101)));
    };
    static fm8(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wlhgm")).length),new BigNumber(865))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, false)).length),new BigNumber((_dafny.Set.fromElements(!(false))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-747),new BigNumber((_dafny.Seq.UnicodeFromString("sfpaxux")).length)));
    };
    static fm9(p0, globalState) {
      return !(_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xtwhavygd"),new BigNumber(606))).length),!(false)))).equals(_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(794),!(true))));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).Merge(((!(false)) ? (_dafny.Map.Empty.slice().updateUnsafe(false,true)) : (_dafny.Map.Empty.slice().updateUnsafe(true,true))));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC4(!(false), (new BigNumber(912)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber(-641))));
    };
    static fm12(p0, globalState) {
      return new _dafny.CodePoint('a'.codePointAt(0));
    };
    static fm18(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("yathjppk"), _dafny.Seq.UnicodeFromString("g")), (_dafny.Set.fromElements(!(false))).IsSubsetOf(_dafny.Set.fromElements(false, false)));
    };
    static fm20(p0, globalState) {
      return _dafny.Seq.Concat((_module.D3.create_DC8(_dafny.Seq.UnicodeFromString("gdigkgpmo"))).dtor_cf14, _dafny.Seq.UnicodeFromString("ybuwlsavw"));
    };
    static fm21(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(true)).Union(_dafny.MultiSet.fromElements(true, false, !(false)))).Union(_dafny.MultiSet.fromElements(true, true));
    };
    static fm22(p0, p1, p2, p3, globalState) {
      if ((false) === (!(!(true)))) {
        return _dafny.MultiSet.fromElements(false);
      } else {
        return (_dafny.MultiSet.fromElements(true)).Difference(_dafny.MultiSet.fromElements(true, false, true, !(false)));
      }
    };
    static fm23(p0, globalState) {
      return ((_dafny.Set.fromElements(_module.D0.create_DC2(_module.D0.create_DC0(true)), _module.D0.create_DC2(_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.of(true, true, false, true, !(!(!(false))))).length))), _module.D0.create_DC2(_module.D0.create_DC1(!(true), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dqkjudxne")).length))).length))), _module.D0.create_DC2(_module.D0.create_DC1(false, new BigNumber(478))))).Union(_dafny.Set.fromElements(_module.D0.create_DC2(_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("lcy")).length))), _module.D0.create_DC2(_module.D0.create_DC1(!(true), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true))).cardinality())))))).Union(_dafny.Set.fromElements(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC0(false))))), _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1(!(true), new BigNumber((_dafny.Seq.of(new BigNumber(650))).length))))));
    };
    static fm24(p0, p1, p2, globalState) {
      return _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1(true, (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())))));
    };
    static fm25(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hjukjem"), _dafny.Seq.UnicodeFromString("wuyw")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("epjwidlg"), _dafny.Seq.UnicodeFromString("ieovtcopc")));
    };
    static fm26(p0, p1, globalState) {
      let _source0 = _module.D9.create_DC27(_module.D9.create_DC26());
      if (_source0.is_DC26) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      } else if (_source0.is_DC25) {
        let _3___mcc_h0 = (_source0).cf50;
        let _4_cf50 = _3___mcc_h0;
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else {
        let _5___mcc_h1 = (_source0).cf51;
        let _6_cf51 = _5___mcc_h1;
        if (false) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('i'.codePointAt(0));
        }
      }
    };
    static fm27(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(false,false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(true,true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(!(true),false))));
    };
    static fm28(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,_module.D1.create_DC5(!(false), _dafny.Map.Empty.slice().updateUnsafe(false,!(true)), new BigNumber((_dafny.Seq.of(new BigNumber(57))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).length), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(859))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(true),_module.D1.create_DC5(true, _dafny.Map.Empty.slice().updateUnsafe(false,true), new BigNumber((_dafny.Seq.UnicodeFromString("np")).length), new BigNumber(692), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(!(false))).cardinality()))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D1.create_DC5(true, _dafny.Map.Empty.slice().updateUnsafe(false,true), new BigNumber((function () {
  let _coll2 = new _dafny.Set();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(354), new BigNumber(741))) {
    let _7_v0 = _compr_2;
    if (((new BigNumber(354)).isLessThanOrEqualTo(_7_v0)) && ((_7_v0).isLessThan(new BigNumber(741)))) {
      _coll2.add(_module.__default.safeModuloInt(_7_v0, new BigNumber(168)));
    }
  }
  return _coll2;
}()).length), new BigNumber(-782), _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(652))))));
    };
    static fm29(p0, p1, p2, globalState) {
      return !(((_module.D15.create_DC37(false, _dafny.Map.Empty.slice().updateUnsafe(true,true), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-626)), function (_8_i0) {
  return new BigNumber(218);
})).length))).dtor_cf67).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("eaasr")).length))) || (false);
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.Concat(_dafny.Seq.of(true, true), _dafny.Seq.of(false)));
    };
    static fm31(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(437), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-733)), function (_9_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })).length)), (_module.D10.create_DC28(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber((_dafny.Seq.of(true)).length)))).dtor_cf52);
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(true, !(_dafny.Set.fromElements(true, false)).equals(_dafny.Set.fromElements(true, false, true)));
    };
    static fm33(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(261)), function (_10_i0) {
        return _module.D1.create_DC4(false, false);
      });
    };
    static fm34(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(938),true)).Merge((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).length)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(495),false)));
    };
    static fm35(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(false, true)).Intersect(_dafny.Set.fromElements((_module.D3.create_DC10(true, _dafny.Map.Empty.slice().updateUnsafe(false,false))).dtor_cf16, true, true)),(new BigNumber(-239)).isLessThanOrEqualTo(new BigNumber(536)));
    };
    static fm36(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(650)).isLessThan(new BigNumber(30)),_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Set.fromElements(false, !(true))).length)));
    };
    static fm37(p0, p1, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(false))).length))), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(676),true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("crt")).length))))).Union((_dafny.MultiSet.fromElements(new BigNumber(-349), new BigNumber(-137), new BigNumber(200))).Union(_dafny.MultiSet.fromElements(new BigNumber(633))));
    };
    static fm38(p0, p1, globalState) {
      return (_dafny.Set.fromElements(_dafny.Set.fromElements(false, true))).Intersect((_dafny.Set.fromElements(_dafny.Set.fromElements(!(true)))).Intersect(_dafny.Set.fromElements(_dafny.Set.fromElements(false, true, true))));
    };
    static fm39(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-876)), function (_11_i0) {
  return new _dafny.CodePoint('g'.codePointAt(0));
}));
    };
    static fm40(p0, globalState) {
      return function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.of(_dafny.Seq.of(_module.D6.create_DC16(_dafny.Seq.of(false, true))))).Elements) {
          let _12_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(_module.D6.create_DC16(_dafny.Seq.of(false, true)))), _12_v0)) {
            _coll3.push([_12_v0,(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-294), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(238), new BigNumber(653), new BigNumber(98))).cardinality()), new BigNumber((_dafny.Seq.of(false)).length))).length)))).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber(-670)))]);
          }
        }
        return _coll3;
      }();
    };
    static fm41(globalState) {
      return (_dafny.Set.fromElements(new BigNumber(((_module.D7.create_DC19(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false)))).dtor_cf34).cardinality()), new BigNumber(-751), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(146)), function (_13_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_14_i1) {
        return new BigNumber(((_module.D10.create_DC28(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(254),false)).length)))).dtor_cf52).length);
      }))).cardinality()), new BigNumber(192))).Union(((_module.D8.create_DC23(new BigNumber(319), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("hm")).length)), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("tcggd")).length))).cardinality()), new BigNumber(581), new BigNumber(-312))).dtor_cf41).Union(function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-842), new BigNumber(321))) {
          let _15_v0 = _compr_4;
          if (((new BigNumber(-842)).isLessThanOrEqualTo(_15_v0)) && ((_15_v0).isLessThan(new BigNumber(321)))) {
            _coll4.add(_module.__default.safeModuloInt(_15_v0, new BigNumber(646)));
          }
        }
        return _coll4;
      }()));
    };
    static Main(__noArgsParameter) {
      let _16_v0;
      _16_v0 = true;
      let _17_v1;
      _17_v1 = _dafny.Map.Empty.slice().updateUnsafe(_16_v0,(_dafny.ZERO).minus(new BigNumber(-313)));
      let _18_v2;
      _18_v2 = _dafny.Seq.of(_dafny.Set.fromElements(_16_v0, _16_v0));
      let _19_v3;
      let _nw0 = Array((new BigNumber(26)).toNumber()).fill(false);
      _19_v3 = _nw0;
      let _20_v4;
      _20_v4 = new BigNumber(333);
      let _21_v5;
      _21_v5 = _dafny.Map.Empty.slice().updateUnsafe(_20_v4,_20_v4);
      let _22_v6;
      _22_v6 = _dafny.Map.Empty.slice().updateUnsafe(_19_v3,_21_v5);
      let _23_v7;
      _23_v7 = _dafny.Seq.UnicodeFromString("hllu");
      let _24_v8;
      let _init0 = ((_25_v4, _26_v0) => function (_27_i0) {
        return _dafny.MultiSet.fromElements(_25_v4, (_dafny.ZERO).minus(_25_v4), new BigNumber((_dafny.MultiSet.fromElements(_26_v0, _26_v0)).cardinality()), new BigNumber(837));
      })(_20_v4, _16_v0);
      let _nw1 = Array((new BigNumber(12)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
        _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _24_v8 = _nw1;
      let _28_globalState;
      let _nw2 = new _module.GlobalState();
      _nw2.__ctor(_17_v1, true, new BigNumber(-406), false, _18_v2, new BigNumber(145), false, new BigNumber(873), new BigNumber(506), true, _22_v6, _23_v7, new BigNumber(895), new BigNumber(755), new BigNumber(670), true, new BigNumber(550), new BigNumber(75), _24_v8, true, false, true, new BigNumber(-889));
      _28_globalState = _nw2;
      let _29_v9;
      _29_v9 = _dafny.Seq.of(new BigNumber((_23_v7).length));
      let _hi0 = ((_16_v0) ? ((_29_v9)[_module.__default.safeIndex(_20_v4, new BigNumber((_29_v9).length))]) : (_20_v4));
      for (let _30_i1 = (_20_v4).minus(new BigNumber((_23_v7).length)); _30_i1.isLessThan(_hi0); _30_i1 = _30_i1.plus(_dafny.ONE)) {
        let _31_v10;
        _31_v10 = new _dafny.CodePoint('v'.codePointAt(0));
        let _32_v11;
        _32_v11 = _dafny.Map.Empty.slice().updateUnsafe(_31_v10,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-829),_16_v0));
        let _33_v13;
        _33_v13 = _dafny.Map.Empty.slice().updateUnsafe(_20_v4,_16_v0);
        let _34_v14;
        _34_v14 = _dafny.Seq.of(function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of (_33_v13).Keys.Elements) {
            let _35_v12 = _compr_5;
            if ((_33_v13).contains(_35_v12)) {
              _coll5.push([_module.__default.safeDivisionInt(_35_v12, _20_v4),_16_v0]);
            }
          }
          return _coll5;
        }());
        let _36_v16;
        _36_v16 = _dafny.Seq.of((((_32_v11).contains(_31_v10)) ? ((_32_v11).get(_31_v10)) : ((_34_v14)[_module.__default.safeIndex(new BigNumber((_29_v9).length), new BigNumber((_34_v14).length))])), function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of _dafny.IntegerRange(new BigNumber(901), new BigNumber(437))) {
            let _37_v15 = _compr_6;
            if (((new BigNumber(901)).isLessThanOrEqualTo(_37_v15)) && ((_37_v15).isLessThan(new BigNumber(437)))) {
              _coll6.push([(_37_v15).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_16_v0, _16_v0)).length))),_16_v0]);
            }
          }
          return _coll6;
        }());
        let _38_v17;
        _38_v17 = _dafny.Set.fromElements(!(_16_v0));
        _36_v16 = _module.__default.fm0(((_16_v0) ? ((_dafny.ZERO).minus(new BigNumber(-98))) : (_module.__default.fm1(_module.__default.fm1(_20_v4, _16_v0, _28_globalState), !(_16_v0), _28_globalState))), _20_v4, _23_v7, new BigNumber((_38_v17).length), _28_globalState);
        (_28_globalState).f5 = _20_v4;
        let _39_v18;
        _39_v18 = _dafny.Set.fromElements(_33_v13);
        _39_v18 = (_39_v18).Difference((_39_v18).Intersect(_39_v18));
        let _rhs0 = _16_v0;
        let _rhs1 = _31_v10;
        let _lhs0 = _28_globalState;
        _lhs0.f9 = _rhs0;
        _31_v10 = _rhs1;
      }
      (_28_globalState).f9 = _16_v0;
      if (_16_v0) {
        let _40_v19;
        _40_v19 = _dafny.Set.fromElements(_16_v0, _16_v0);
        let _41_v20;
        _41_v20 = _dafny.Map.Empty.slice().updateUnsafe(_16_v0,(_40_v19).IsProperSubsetOf(_40_v19));
        let _42_v21;
        _42_v21 = _dafny.Seq.of(_16_v0);
        (_28_globalState).f6 = (((_41_v20).contains((_42_v21)[_module.__default.safeIndex(_20_v4, new BigNumber((_42_v21).length))])) ? ((_41_v20).get((_42_v21)[_module.__default.safeIndex(_20_v4, new BigNumber((_42_v21).length))])) : (!(_16_v0)));
        (_28_globalState).f14 = _20_v4;
        let _43_v22;
        _43_v22 = _dafny.Set.fromElements(_23_v7);
        if (!((_43_v22).IsSubsetOf((_43_v22).Difference(_43_v22)))) {
          let _44_v23;
          let _init1 = ((_45_v0) => function (_46_i2) {
            return _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false, _45_v0));
          })(_16_v0);
          let _nw3 = Array((new BigNumber(8)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw3.length); _i0_1++) {
            _nw3[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _44_v23 = _nw3;
          let _47_v24;
          _47_v24 = _dafny.MultiSet.fromElements(_16_v0);
          let _48_v25;
          _48_v25 = _dafny.MultiSet.fromElements(_47_v24);
          let _index0 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_44_v23).length));
          (_44_v23)[_index0] = _48_v25;
          let _49_v26;
          _49_v26 = _dafny.Map.Empty.slice().updateUnsafe(_16_v0,_48_v25);
          let _50_v27;
          let _nw4 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _50_v27 = _nw4;
          let _51_v28;
          _51_v28 = _dafny.Map.Empty.slice().updateUnsafe(_50_v27,_16_v0);
          let _52_v30;
          _52_v30 = new _dafny.CodePoint('t'.codePointAt(0));
          let _53_v31;
          _53_v31 = _dafny.MultiSet.fromElements(_52_v30);
          let _54_v32;
          _54_v32 = _dafny.Set.fromElements(function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of (_53_v31).Elements) {
              let _55_v29 = _compr_7;
              if ((_53_v31).contains(_55_v29)) {
                _coll7.push([_55_v29,_16_v0]);
              }
            }
            return _coll7;
          }());
          let _index1 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_44_v23).length));
          let _rhs2 = (_48_v25).Difference((((_49_v26).contains((((_51_v28).contains(_50_v27)) ? ((_51_v28).get(_50_v27)) : (_16_v0)))) ? ((_49_v26).get((((_51_v28).contains(_50_v27)) ? ((_51_v28).get(_50_v27)) : (_16_v0)))) : (_48_v25)));
          let _rhs3 = _17_v1;
          let _rhs4 = new BigNumber((((_54_v32).Intersect(_54_v32)).Union(_54_v32)).length);
          let _lhs1 = _44_v23;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_44_v23).length));
          let _lhs3 = _28_globalState;
          _lhs1[_lhs2] = _rhs2;
          _17_v1 = _rhs3;
          _lhs3.f5 = _rhs4;
          let _56_v33;
          let _nw5 = Array((new BigNumber(5)).toNumber()).fill([]);
          _56_v33 = _nw5;
          let _rhs5 = (((_42_v21)[_module.__default.safeIndex(new BigNumber(125), new BigNumber((_42_v21).length))]) ? (_56_v33) : (_56_v33));
          let _rhs6 = (_20_v4).minus(_20_v4);
          let _lhs4 = _28_globalState;
          _56_v33 = _rhs5;
          _lhs4.f17 = _rhs6;
          let _57_v34;
          let _nw6 = new _module.C2();
          _nw6.__ctor(_20_v4);
          _57_v34 = _nw6;
          let _index2 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_56_v33).length));
          (_56_v33)[_index2] = ((_16_v0) ? (_50_v27) : (_50_v27));
          let _index3 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_56_v33).length));
          let _rhs7 = !(_16_v0);
          let _rhs8 = _16_v0;
          let _rhs9 = _17_v1;
          let _rhs10 = _50_v27;
          let _lhs5 = _28_globalState;
          let _lhs6 = _28_globalState;
          let _lhs7 = _28_globalState;
          let _lhs8 = _56_v33;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_56_v33).length));
          _lhs5.f1 = _rhs7;
          _lhs6.f1 = _rhs8;
          _lhs7.f0 = _rhs9;
          _lhs8[_lhs9] = _rhs10;
          let _58_v35;
          let _nw7 = Array((new BigNumber(27)).toNumber());
          _58_v35 = _nw7;
          let _59_v36;
          _59_v36 = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_58_v35)).cardinality()), new BigNumber((_21_v5).length));
          let _60_v37;
          let _out0;
          _out0 = (_57_v34).m7(_19_v3, _16_v0, _16_v0, _module.__default.fm29(new BigNumber((_59_v36).length), _20_v4, _16_v0, _28_globalState), _28_globalState);
          _60_v37 = _out0;
        } else {
          let _61_v38;
          _61_v38 = new _dafny.CodePoint('j'.codePointAt(0));
          let _rhs11 = _61_v38;
          let _rhs12 = _20_v4;
          let _lhs10 = _28_globalState;
          _61_v38 = _rhs11;
          _lhs10.f14 = _rhs12;
          let _62_v39;
          _62_v39 = _dafny.Set.fromElements(_19_v3);
          let _63_v40;
          let _nw8 = new _module.C0();
          _nw8.__ctor((_62_v39).Union(_62_v39), _20_v4, _20_v4);
          _63_v40 = _nw8;
          let _64_v41;
          let _nw9 = new _module.C8();
          _nw9.__ctor(_20_v4, _16_v0, _16_v0);
          _64_v41 = _nw9;
          _64_v41 = _64_v41;
          let _65_v42;
          let _nw10 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
          _65_v42 = _nw10;
          let _66_v43;
          let _nw11 = new _module.C4();
          _nw11.__ctor(_16_v0, _16_v0, _16_v0);
          _66_v43 = _nw11;
          let _67_v44;
          _67_v44 = _dafny.Map.Empty.slice().updateUnsafe(_66_v43,new BigNumber((_42_v21).length));
          let _68_v45;
          _68_v45 = _dafny.Map.Empty.slice().updateUnsafe(_63_v40,_67_v44);
          let _index4 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_65_v42).length));
          (_65_v42)[_index4] = ((_16_v0) ? (_68_v45) : (_68_v45));
          let _index5 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_65_v42).length));
          (_65_v42)[_index5] = (_68_v45).Merge((_68_v45).Merge((_68_v45).update(_63_v40, _67_v44)));
          let _69_v46;
          let _nw12 = Array((new BigNumber(19)).toNumber());
          _69_v46 = _nw12;
          let _70_v47;
          _70_v47 = _dafny.Set.fromElements(_69_v46);
          let _71_v48;
          let _nw13 = Array((new BigNumber(17)).toNumber());
          _71_v48 = _nw13;
          let _72_v49;
          let _nw14 = new _module.C4();
          _nw14.__ctor(_66_v43.f24, _66_v43.f24, _16_v0);
          _72_v49 = _nw14;
          let _index6 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_71_v48).length));
          (_71_v48)[_index6] = _72_v49;
          let _index7 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_71_v48).length));
          let _rhs13 = ((((_66_v43).f25) ? (new BigNumber((_dafny.Set.fromElements(_63_v40.f31)).length)) : (_63_v40.f31))).isLessThan(_20_v4);
          let _rhs14 = _70_v47;
          let _rhs15 = _63_v40.f31;
          let _rhs16 = _72_v49;
          let _lhs11 = _28_globalState;
          let _lhs12 = _63_v40;
          let _lhs13 = _71_v48;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_71_v48).length));
          _lhs11.f15 = _rhs13;
          _70_v47 = _rhs14;
          _lhs12.f31 = _rhs15;
          _lhs13[_lhs14] = _rhs16;
        }
        let _73_v50;
        _73_v50 = _dafny.Set.fromElements(_19_v3, _19_v3, _19_v3);
        let _74_v51;
        let _nw15 = new _module.C0();
        _nw15.__ctor(_73_v50, _20_v4, _module.__default.safeModuloInt(_20_v4, _20_v4));
        _74_v51 = _nw15;
        let _75_v52;
        _75_v52 = _dafny.Map.Empty.slice().updateUnsafe((_23_v7)[_module.__default.safeIndex(_74_v51.f31, new BigNumber((_23_v7).length))],_dafny.Seq.of(false, _16_v0));
        (_28_globalState).f8 = (_dafny.ZERO).minus((_74_v51.f31).multipliedBy(((_dafny.ZERO).minus(new BigNumber((_75_v52).length))).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("ja")).length))));
      } else {
        let _76_v53;
        _76_v53 = _module.D3.create_DC10(_16_v0, _dafny.Map.Empty.slice().updateUnsafe(_16_v0,_16_v0));
        let _77_v54;
        _77_v54 = _dafny.Seq.of(!((_20_v4).isLessThanOrEqualTo(_20_v4)), (_76_v53).dtor_cf16, _16_v0);
        _77_v54 = _dafny.Seq.Concat(_dafny.Seq.update(_77_v54, _module.__default.safeIndex(_20_v4, new BigNumber((_77_v54).length)), _16_v0), _77_v54);
        let _78_v55;
        _78_v55 = _dafny.Set.fromElements((new BigNumber((_23_v7).length)).isEqualTo(_20_v4), !(_16_v0));
        let _79_v56;
        _79_v56 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(626)), function (_80_i4) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }));
        let _rhs17 = _19_v3;
        let _rhs18 = _78_v55;
        let _rhs19 = new BigNumber(((((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(689)), function (_81_i3) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }))).update(_23_v7, _module.__default.abs(_20_v4))).Union(_79_v56)).Difference(_79_v56)).cardinality());
        _19_v3 = _rhs17;
        _78_v55 = _rhs18;
        _20_v4 = _rhs19;
        let _82_v57;
        let _nw16 = new _module.C3();
        _nw16.__ctor(_16_v0, ((true) ? (new BigNumber(-487)) : (_20_v4)), !(!(_16_v0)), !((new BigNumber(-148)).isLessThan(_20_v4)), _20_v4);
        _82_v57 = _nw16;
        let _83_v58;
        _83_v58 = _dafny.Set.fromElements(_19_v3);
        let _84_v59;
        _84_v59 = _dafny.MultiSet.fromElements(!(_16_v0), _16_v0);
        let _85_v60;
        let _nw17 = new _module.C0();
        _nw17.__ctor((_dafny.Set.fromElements(_19_v3)).Union(_83_v58), _20_v4, (_20_v4).plus(new BigNumber((_84_v59).cardinality())));
        _85_v60 = _nw17;
        let _86_v61;
        let _nw18 = new _module.C0();
        _nw18.__ctor(_dafny.Set.fromElements(_19_v3, _19_v3), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_19_v3,(_82_v57).f28)).length), (_82_v57).f29);
        _86_v61 = _nw18;
      }
      let _87_v62;
      let _nw19 = new _module.C4();
      _nw19.__ctor(_16_v0, _16_v0, _16_v0);
      _87_v62 = _nw19;
      let _88_v63;
      let _nw20 = Array((new BigNumber(24)).toNumber()).fill(_module.D6.Default());
      _88_v63 = _nw20;
      let _89_v64;
      _89_v64 = _dafny.Map.Empty.slice().updateUnsafe(_87_v62,_88_v63);
      let _90_v65;
      _90_v65 = _dafny.Map.Empty.slice().updateUnsafe((((_89_v64).contains(_87_v62)) ? ((_89_v64).get(_87_v62)) : (_88_v63)),(new BigNumber(22)).isEqualTo(_20_v4));
      _90_v65 = (_90_v65).update(_88_v63, _16_v0);
      let _index8 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
      (_19_v3)[_index8] = _16_v0;
      let _index9 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
      (_19_v3)[_index9] = _87_v62.f24;
      (_28_globalState).f1 = !(_module.__default.fm34(_20_v4, _28_globalState)).contains((_20_v4).minus(new BigNumber(660)));
      let _91_v66;
      let _nw21 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _91_v66 = _nw21;
      let _index10 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length));
      (_91_v66)[_index10] = _20_v4;
      let _index11 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length));
      let _rhs20 = _20_v4;
      let _rhs21 = !((_87_v62).f25);
      let _rhs22 = _16_v0;
      let _lhs15 = _91_v66;
      let _lhs16 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length));
      let _lhs17 = _28_globalState;
      let _lhs18 = _28_globalState;
      _lhs15[_lhs16] = _rhs20;
      _lhs17.f9 = _rhs21;
      _lhs18.f1 = _rhs22;
      let _92_v67;
      let _nw22 = Array((new BigNumber(14)).toNumber()).fill([]);
      _92_v67 = _nw22;
      let _93_v68;
      let _nw23 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
      _93_v68 = _nw23;
      let _index12 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_92_v67).length));
      (_92_v67)[_index12] = _93_v68;
      let _index13 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_92_v67).length));
      (_92_v67)[_index13] = _93_v68;
      let _94_v69;
      _94_v69 = _dafny.Set.fromElements(_19_v3, _19_v3, _19_v3);
      let _95_v70;
      let _nw24 = new _module.C0();
      _nw24.__ctor(_94_v69, (_20_v4).multipliedBy((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))]), _20_v4);
      _95_v70 = _nw24;
      let _96_v71;
      _96_v71 = _dafny.MultiSet.fromElements(_87_v62.f24);
      let _97_v72;
      _97_v72 = _dafny.Map.Empty.slice().updateUnsafe(_96_v71,_19_v3);
      let _rhs23 = _95_v70;
      let _rhs24 = (_dafny.ZERO).minus((_95_v70).f23);
      let _rhs25 = _97_v72;
      let _lhs19 = _28_globalState;
      _95_v70 = _rhs23;
      _lhs19.f8 = _rhs24;
      _97_v72 = _rhs25;
      let _98_v73;
      _98_v73 = _dafny.Map.Empty.slice().updateUnsafe((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))],_16_v0);
      _98_v73 = (_98_v73).update((_95_v70).f23, (((_87_v62).f25) ? (!(_87_v62.f24)) : ((_87_v62).f25)));
      let _99_v74;
      _99_v74 = _dafny.Seq.of(_87_v62.f24);
      let _100_v75;
      _100_v75 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_99_v74).length),_23_v7);
      _23_v7 = _dafny.Seq.Concat((((_100_v75).contains(new BigNumber(456))) ? ((_100_v75).get(new BigNumber(456))) : (_dafny.Seq.UnicodeFromString("gstyka"))), _dafny.Seq.Concat(_23_v7, _23_v7));
      let _101_v76;
      _101_v76 = _dafny.Map.Empty.slice().updateUnsafe((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))],(_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))]);
      let _102_v77;
      _102_v77 = _dafny.Map.Empty.slice().updateUnsafe(_101_v76,new BigNumber((_21_v5).length));
      let _103_v78;
      _103_v78 = _dafny.Map.Empty.slice().updateUnsafe(_102_v77,(_87_v62).f25);
      let _104_v79;
      let _nw25 = new _module.C1();
      _nw25.__ctor(true, (((_103_v78).contains(_102_v77)) ? ((_103_v78).get(_102_v77)) : ((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))])));
      _104_v79 = _nw25;
      let _105_i5;
      _105_i5 = _dafny.ZERO;
      L0: {
        while ((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))]) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_105_i5)) {
              break L0;
            }
            _105_i5 = (_105_i5).plus(_dafny.ONE);
            let _106_v80;
            let _nw26 = new _module.C0();
            _nw26.__ctor(_94_v69, (_95_v70).f23, (_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))]);
            _106_v80 = _nw26;
            let _107_v81;
            _107_v81 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_95_v70).f23, new BigNumber((_dafny.MultiSet.FromArray(_99_v74)).cardinality())),_106_v80);
            _107_v81 = (_107_v81).update(_module.__default.safeDivisionInt((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))], _20_v4), _106_v80);
            let _index14 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
            (_19_v3)[_index14] = false;
            let _108_v82;
            _108_v82 = _dafny.MultiSet.fromElements((_95_v70).f23, _106_v80.f31, new BigNumber((_23_v7).length));
            let _109_v83;
            let _nw27 = new _module.C7();
            _nw27.__ctor();
            _109_v83 = _nw27;
            let _110_v84;
            _110_v84 = _dafny.Map.Empty.slice().updateUnsafe(_109_v83,(_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))]);
            (_28_globalState).f9 = !(_108_v82).equals(_dafny.MultiSet.fromElements((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))], new BigNumber((_dafny.Seq.update(_29_v9, _module.__default.safeIndex((((_21_v5).contains(_106_v80.f31)) ? ((_21_v5).get(_106_v80.f31)) : ((((_21_v5).contains((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))])) ? ((_21_v5).get((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))])) : (_106_v80.f31)))), new BigNumber((_29_v9).length)), new BigNumber((_110_v84).length))).length)));
            if (((_95_v70).f23).isLessThanOrEqualTo((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))])) {
              let _index15 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length));
              (_91_v66)[_index15] = ((((_87_v62).f25) ? (_20_v4) : ((_95_v70).f23))).minus((_95_v70).f23);
              let _111_v85;
              _111_v85 = new _dafny.CodePoint('a'.codePointAt(0));
              _111_v85 = _111_v85;
              (_28_globalState).f17 = new BigNumber(297);
              let _112_v86;
              _112_v86 = _dafny.Map.Empty.slice().updateUnsafe(_101_v76,_87_v62);
              let _113_v87;
              _113_v87 = _dafny.Map.Empty.slice().updateUnsafe((_87_v62).f25,_87_v62);
              let _114_v88;
              let _nw28 = Array((new BigNumber(23)).toNumber());
              _nw28[(_dafny.ZERO).toNumber()] = ((!((_87_v62).f25)) ? (_87_v62) : (_87_v62));
              _nw28[(_dafny.ONE).toNumber()] = _87_v62;
              _nw28[(new BigNumber(2)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(3)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(4)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(5)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(6)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(7)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(8)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(9)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(10)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(11)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(12)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(13)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(14)).toNumber()] = (((_87_v62).f25) ? (_87_v62) : ((((_112_v86).contains(_101_v76)) ? ((_112_v86).get(_101_v76)) : ((((_113_v87).contains(_16_v0)) ? ((_113_v87).get(_16_v0)) : (_87_v62))))));
              _nw28[(new BigNumber(15)).toNumber()] = ((_87_v62.f24) ? (_87_v62) : (_87_v62));
              _nw28[(new BigNumber(16)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(17)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(18)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(19)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(20)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(21)).toNumber()] = _87_v62;
              _nw28[(new BigNumber(22)).toNumber()] = _87_v62;
              _114_v88 = _nw28;
              _114_v88 = _114_v88;
              let _115_v91;
              _115_v91 = _dafny.MultiSet.fromElements(_98_v73, (function () {
                let _coll8 = new _dafny.Map();
                for (const _compr_8 of _dafny.IntegerRange(new BigNumber(391), new BigNumber(463))) {
                  let _116_v89 = _compr_8;
                  if (((new BigNumber(391)).isLessThanOrEqualTo(_116_v89)) && ((_116_v89).isLessThan(new BigNumber(463)))) {
                    _coll8.push([_module.__default.safeDivisionInt(_116_v89, (_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))]),(_87_v62).f25]);
                  }
                }
                return _coll8;
              }()).Merge(function () {
                let _coll9 = new _dafny.Map();
                for (const _compr_9 of (_98_v73).Keys.Elements) {
                  let _117_v90 = _compr_9;
                  if ((_98_v73).contains(_117_v90)) {
                    _coll9.push([_module.__default.safeDivisionInt(_117_v90, _20_v4),true]);
                  }
                }
                return _coll9;
              }()), _98_v73, _98_v73, _98_v73);
              let _rhs26 = _115_v91;
              let _rhs27 = _16_v0;
              let _lhs20 = _28_globalState;
              _115_v91 = _rhs26;
              _lhs20.f1 = _rhs27;
            } else {
              let _118_v92;
              let _nw29 = new _module.C3();
              _nw29.__ctor((_87_v62).f25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(288)), ((_119_v62) => function (_120_i6) {
                return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),(_119_v62).f25);
              })(_87_v62))).length), (_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))], _87_v62.f24, _106_v80.f31);
              _118_v92 = _nw29;
              let _121_v93;
              _121_v93 = _dafny.Seq.of(_118_v92, _118_v92);
              _121_v93 = _dafny.Seq.Concat(_121_v93, _dafny.Seq.Concat(_121_v93, _121_v93));
              let _122_v94;
              _122_v94 = _dafny.Map.Empty.slice().updateUnsafe(_106_v80,_96_v71);
              let _rhs28 = _87_v62.f24;
              let _rhs29 = ((_122_v94).Merge(_122_v94)).Merge(_122_v94);
              let _lhs21 = _87_v62;
              _lhs21.f24 = _rhs28;
              _122_v94 = _rhs29;
              let _123_v95;
              _123_v95 = _118_v92;
              let _index16 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
              let _rhs30 = (_87_v62).f25;
              let _rhs31 = false;
              let _rhs32 = (_123_v95);
              let _rhs33 = _19_v3;
              let _lhs22 = _19_v3;
              let _lhs23 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
              let _lhs24 = _28_globalState;
              _lhs22[_lhs23] = _rhs30;
              _lhs24.f6 = _rhs31;
              _118_v92 = _rhs32;
              _19_v3 = _rhs33;
              let _124_v96;
              let _125_v97;
              let _126_v98;
              let _127_v99;
              let _out1;
              let _out2;
              let _out3;
              let _out4;
              let _outcollector0 = (_104_v79).m1(_28_globalState);
              _out1 = _outcollector0[0];
              _out2 = _outcollector0[1];
              _out3 = _outcollector0[2];
              _out4 = _outcollector0[3];
              _124_v96 = _out1;
              _125_v97 = _out2;
              _126_v98 = _out3;
              _127_v99 = _out4;
              _19_v3 = _19_v3;
            }
          }
        }
      }
      let _128_v100;
      let _nw30 = Array((new BigNumber(9)).toNumber());
      _128_v100 = _nw30;
      let _129_v101;
      _129_v101 = _dafny.Seq.of(_128_v100);
      let _130_v102;
      _130_v102 = _dafny.Seq.of((_129_v101)[_module.__default.safeIndex((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))], new BigNumber((_129_v101).length))]);
      let _131_v103;
      _131_v103 = _module.D3.create_DC10((_87_v62).f25, _101_v76);
      let _pat_let_tv0 = _101_v76;
      let _132_v104;
      _132_v104 = _dafny.Seq.of(function (_pat_let0_0) {
        return function (_133_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_134_dt__update_hcf17_h0) {
              return _module.D3.create_DC10((_133_dt__update__tmp_h0).dtor_cf16, _134_dt__update_hcf17_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_module.D3.create_DC10((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))], _101_v76)), _131_v103);
      let _135_v105;
      _135_v105 = _dafny.Seq.of(_87_v62, _87_v62, _87_v62);
      let _136_v106;
      _136_v106 = _dafny.Set.fromElements(_135_v105);
      let _index17 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
      let _index18 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length));
      let _index19 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
      let _rhs34 = (new BigNumber((_132_v104).length)).isLessThan((_95_v70).f23);
      let _rhs35 = new BigNumber(((((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))]) ? (_136_v106) : (_136_v106))).length);
      let _rhs36 = _16_v0;
      let _rhs37 = _20_v4;
      let _rhs38 = _dafny.Seq.Concat(_130_v102, _129_v101);
      let _lhs25 = _19_v3;
      let _lhs26 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
      let _lhs27 = _91_v66;
      let _lhs28 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length));
      let _lhs29 = _19_v3;
      let _lhs30 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
      let _lhs31 = _28_globalState;
      _lhs25[_lhs26] = _rhs34;
      _lhs27[_lhs28] = _rhs35;
      _lhs29[_lhs30] = _rhs36;
      _lhs31.f8 = _rhs37;
      _129_v101 = _rhs38;
      let _137_v107;
      _137_v107 = new _dafny.CodePoint('d'.codePointAt(0));
      let _138_v108;
      _138_v108 = _dafny.Map.Empty.slice().updateUnsafe((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))],_137_v107);
      let _139_v109;
      _139_v109 = _dafny.Map.Empty.slice().updateUnsafe(_138_v108,_16_v0);
      (_104_v79).m8(!((((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))]) ? ((_87_v62).f25) : (!(_16_v0)))), (_95_v70).fm2(_28_globalState), _139_v109, false, _28_globalState);
      if (_16_v0) {
        if (!_dafny.Seq.contains(_23_v7, _137_v107)) {
          let _140_v110;
          let _141_v111;
          let _142_v112;
          let _143_v113;
          let _out5;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector1 = (_87_v62).m1(_28_globalState);
          _out5 = _outcollector1[0];
          _out6 = _outcollector1[1];
          _out7 = _outcollector1[2];
          _out8 = _outcollector1[3];
          _140_v110 = _out5;
          _141_v111 = _out6;
          _142_v112 = _out7;
          _143_v113 = _out8;
          let _144_v114;
          let _nw31 = new _module.C5();
          _nw31.__ctor(_dafny.areEqual(_23_v7, _dafny.Seq.update(_module.__default.fm25((_87_v62).f25, (_87_v62).f25, _137_v107, _28_globalState), _module.__default.safeIndex((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))], new BigNumber((_module.__default.fm25((_87_v62).f25, (_87_v62).f25, _137_v107, _28_globalState)).length)), new _dafny.CodePoint('t'.codePointAt(0)))), _140_v110);
          _144_v114 = _nw31;
          let _145_v115;
          let _nw32 = new _module.C0();
          _nw32.__ctor(_94_v69, (_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))], (_95_v70).f23);
          _145_v115 = _nw32;
          _145_v115 = _145_v115;
          let _146_v116;
          _146_v116 = _dafny.Seq.of(_91_v66);
          _91_v66 = (_146_v116)[_module.__default.safeIndex((_20_v4).multipliedBy((_29_v9)[_module.__default.safeIndex((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))], new BigNumber((_29_v9).length))]), new BigNumber((_146_v116).length))];
          let _147_v117;
          _147_v117 = _dafny.Seq.of(_135_v105);
          let _148_v118;
          _148_v118 = _dafny.MultiSet.fromElements((_141_v111).f23, _145_v115.f31);
          let _149_v119;
          _149_v119 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_143_v113, _20_v4),(_147_v117)[_module.__default.safeIndex(new BigNumber((_148_v118).cardinality()), new BigNumber((_147_v117).length))]);
          _149_v119 = (_149_v119).update(_20_v4, _dafny.Seq.update(_135_v105, _module.__default.safeIndex((_141_v111).f23, new BigNumber((_135_v105).length)), _87_v62));
        } else {
          _23_v7 = _dafny.Seq.Concat(_dafny.Seq.Concat(_23_v7, _23_v7), _23_v7);
          let _index20 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
          (_19_v3)[_index20] = (_87_v62).f25;
          let _150_v120;
          let _nw33 = new _module.C0();
          _nw33.__ctor(_94_v69, (_95_v70).f23, (_95_v70).f23);
          _150_v120 = _nw33;
          let _rhs39 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_95_v70).f23, (_95_v70).f23), (_95_v70).f23);
          let _rhs40 = _150_v120;
          let _lhs32 = _28_globalState;
          _lhs32.f8 = _rhs39;
          _150_v120 = _rhs40;
          let _151_v121;
          _151_v121 = _dafny.Set.fromElements(!(_87_v62.f24), ((_87_v62.f24) ? ((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))]) : (!(!(!((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))]))))));
          _151_v121 = _151_v121;
          let _152_v122;
          _152_v122 = _dafny.MultiSet.fromElements(_23_v7, _dafny.Seq.UnicodeFromString("guaqmlf"));
          _16_v0 = !((((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))]).plus(new BigNumber((_23_v7).length))).isEqualTo((((_152_v122).contains(_23_v7)) ? ((_152_v122).get(_23_v7)) : ((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))]))));
        }
        _104_v79 = _104_v79;
        let _153_v123;
        _153_v123 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC10((_87_v62).f25, _101_v76),_17_v1);
        _153_v123 = ((_87_v62.f24) ? ((_153_v123).Merge(_153_v123)) : (_153_v123));
        let _154_v124;
        _154_v124 = _dafny.Set.fromElements((_87_v62).f25, (_87_v62).f25);
        let _index21 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length));
        (_91_v66)[_index21] = (_dafny.ZERO).minus(new BigNumber((_154_v124).length));
        _16_v0 = (_154_v124).IsProperSubsetOf(_154_v124);
      } else {
        (_28_globalState).f5 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))]), new BigNumber((_23_v7).length));
        let _155_v125;
        let _nw34 = Array((new BigNumber(21)).toNumber()).fill([]);
        _155_v125 = _nw34;
        let _index22 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_155_v125).length));
        (_155_v125)[_index22] = _19_v3;
        let _156_v126;
        _156_v126 = _dafny.Seq.of(_19_v3);
        let _index23 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_155_v125).length));
        (_155_v125)[_index23] = ((true) ? (((_16_v0) ? (_19_v3) : (_19_v3))) : ((_156_v126)[_module.__default.safeIndex((_95_v70).f23, new BigNumber((_156_v126).length))]));
        let _157_v127;
        let _nw35 = new _module.C7();
        _nw35.__ctor();
        _157_v127 = _nw35;
        let _158_v128;
        _158_v128 = _dafny.Seq.of(_157_v127);
        _157_v127 = (_158_v128)[_module.__default.safeIndex((_91_v66)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_91_v66).length))], new BigNumber((_158_v128).length))];
        let _159_v129;
        _159_v129 = _dafny.Map.Empty.slice().updateUnsafe((_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))],_94_v69);
        let _160_v130;
        let _nw36 = new _module.C0();
        _nw36.__ctor((((_159_v129).contains((_87_v62).f25)) ? ((_159_v129).get((_87_v62).f25)) : (_94_v69)), (_95_v70).fm2(_28_globalState), (_95_v70).f23);
        _160_v130 = _nw36;
        _160_v130 = _160_v130;
        let _161_v131;
        _161_v131 = _dafny.Seq.of(_91_v66, _91_v66);
        let _index24 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
        let _rhs41 = (_module.__default.safeModuloInt((_95_v70).f23, _160_v130.f31)).plus((_95_v70).f23);
        let _rhs42 = (_19_v3)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length))];
        let _rhs43 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_161_v131, _dafny.Seq.of(_91_v66)), _161_v131);
        let _lhs33 = _28_globalState;
        let _lhs34 = _19_v3;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_19_v3).length));
        let _lhs36 = _28_globalState;
        _lhs33.f8 = _rhs41;
        _lhs34[_lhs35] = _rhs42;
        _lhs36.f9 = _rhs43;
      }
      process.stdout.write(_dafny.toString(_16_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v1).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(313)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_18_v2, _dafny.Seq.of(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v3)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_19_v3)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_20_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_21_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(333),new BigNumber(333)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_22_v6).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_23_v7).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_24_v8)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState.f0).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(313)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_28_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_28_globalState.f4, _dafny.Seq.of(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_28_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_28_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_28_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_28_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_28_globalState).f10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_28_globalState).f11).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_28_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_28_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_28_globalState.f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_28_globalState).f18)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(new BigNumber(333), new BigNumber(-333), new BigNumber(2), new BigNumber(837)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_28_globalState).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_29_v9, _dafny.Seq.of(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_87_v62.f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_87_v62).f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_89_v64).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_90_v65).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v66)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_94_v69).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_95_v70).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_v71).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_97_v72).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v73).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(333),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_99_v74, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v75).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.Seq.UnicodeFromString("hllu")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v76).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v77).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,true),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_v78).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,true),_dafny.ONE),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_129_v101).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_130_v102).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v103).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_v103).dtor_cf17).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_132_v104, _dafny.Seq.of(_module.D3.create_DC10(false, _dafny.Map.Empty.slice().updateUnsafe(true,true)), _module.D3.create_DC10(true, _dafny.Map.Empty.slice().updateUnsafe(true,true))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_135_v105).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_136_v106).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_v107));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v108).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v109).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('d'.codePointAt(0))),true))));
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
        return other.$tag === 1 && this.cf0 === other.cf0;
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
    static create_DC4(cf5, cf6) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf7, cf8, cf9, cf10, cf11) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC6(cf12) {
      let $dt = new D1(3);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf5 === other.cf5 && this.cf6 === other.cf6;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false, false);
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
    static create_DC7(cf13) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13;
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
    static create_DC9(cf15) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC10(cf16, cf17) {
      let $dt = new D3(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC8(cf14) {
      let $dt = new D3(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + this.cf14.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
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
    static create_DC12(cf19, cf20) {
      let $dt = new D4(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC11(cf18) {
      let $dt = new D4(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC14(cf22, cf23, cf24) {
      let $dt = new D5(0);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf25, cf26, cf27, cf28) {
      let $dt = new D5(1);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC13(cf21) {
      let $dt = new D5(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(_dafny.ZERO, false, _dafny.Set.Empty);
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
    static create_DC17(cf30) {
      let $dt = new D6(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC18(cf31, cf32, cf33) {
      let $dt = new D6(1);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC16(cf29) {
      let $dt = new D6(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf31 === other.cf31 && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false);
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
    static create_DC20(cf35, cf36, cf37) {
      let $dt = new D7(0);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC19(cf34) {
      let $dt = new D7(1);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC21(cf38) {
      let $dt = new D7(2);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC20(false, _dafny.Map.Empty, []);
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
    static create_DC23(cf40, cf41, cf42, cf43, cf44) {
      let $dt = new D8(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC24(cf45, cf46, cf47, cf48, cf49) {
      let $dt = new D8(1);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC22(cf39) {
      let $dt = new D8(2);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42) && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf45, other.cf45) && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && this.cf48 === other.cf48 && this.cf49 === other.cf49;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf39 === other.cf39;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(_dafny.ZERO, _dafny.Set.Empty, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC26() {
      let $dt = new D9(0);
      return $dt;
    }
    static create_DC25(cf50) {
      let $dt = new D9(1);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC27(cf51) {
      let $dt = new D9(2);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26";
      } else if (this.$tag === 1) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf51) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC26();
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
    static create_DC29(cf53, cf54, cf55) {
      let $dt = new D10(0);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC28(cf52) {
      let $dt = new D10(1);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf53 === other.cf53 && this.cf54 === other.cf54 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29(false, false, _dafny.ZERO);
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
    static create_DC31(cf57, cf58, cf59) {
      let $dt = new D11(0);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC30(cf56) {
      let $dt = new D11(1);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58) && this.cf59 === other.cf59;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf56 === other.cf56;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC31([], _dafny.ZERO, false);
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
    static create_DC33(cf61) {
      let $dt = new D12(0);
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC32(cf60) {
      let $dt = new D12(1);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC33(_dafny.ZERO);
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
    static create_DC34(cf62) {
      let $dt = new D13(0);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf62, other.cf62);
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
    static create_DC35(cf63) {
      let $dt = new D14(0);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf63) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC37(cf65, cf66, cf67) {
      let $dt = new D15(0);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC38(cf68) {
      let $dt = new D15(1);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC36(cf64) {
      let $dt = new D15(2);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC37" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC36" + "(" + _dafny.toString(this.cf64) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf65 === other.cf65 && _dafny.areEqual(this.cf66, other.cf66) && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf68 === other.cf68;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC37(false, _dafny.Map.Empty, _dafny.ZERO);
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
    static create_DC40(cf70, cf71, cf72) {
      let $dt = new D16(0);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC39(cf69) {
      let $dt = new D16(1);
      $dt.cf69 = cf69;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf69() { return this.cf69; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC40" + "(" + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf69) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf70, other.cf70) && this.cf71 === other.cf71 && _dafny.areEqual(this.cf72, other.cf72);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf69 === other.cf69;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC40(_dafny.ZERO, [], _dafny.Map.Empty);
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
    static create_DC42() {
      let $dt = new D17(0);
      return $dt;
    }
    static create_DC41(cf73) {
      let $dt = new D17(1);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC42";
      } else if (this.$tag === 1) {
        return "D17.DC41" + "(" + _dafny.toString(this.cf73) + ")";
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
        return other.$tag === 1 && this.cf73 === other.cf73;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC42();
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
    static create_DC43(cf74) {
      let $dt = new D18(0);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC43" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf74 === other.cf74;
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
          return D18.Default();
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
      this.f0 = _dafny.Map.Empty;
      this.f1 = false;
      this.f4 = _dafny.Seq.of();
      this.f5 = _dafny.ZERO;
      this.f6 = false;
      this.f8 = _dafny.ZERO;
      this.f9 = false;
      this.f14 = _dafny.ZERO;
      this.f15 = false;
      this.f17 = _dafny.ZERO;
      this._f2 = _dafny.ZERO;
      this._f3 = false;
      this._f7 = _dafny.ZERO;
      this._f10 = _dafny.Map.Empty;
      this._f11 = _dafny.Seq.UnicodeFromString("");
      this._f12 = _dafny.ZERO;
      this._f13 = _dafny.ZERO;
      this._f16 = _dafny.ZERO;
      this._f18 = [];
      this._f19 = false;
      this._f20 = false;
      this._f21 = false;
      this._f22 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this).f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this).f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      return;
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f7() {
      let _this = this;
      return _this._f7;
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
    get f16() {
      let _this = this;
      return _this._f16;
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
      this._f23 = _dafny.ZERO;
      this.f31 = _dafny.ZERO;
      this._f30 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    __ctor(f30, f31, f23) {
      let _this = this;
      (_this)._f30 = f30;
      (_this).f31 = f31;
      (_this)._f23 = f23;
      return;
    }
    fm2(globalState) {
      let _this = this;
      return (new BigNumber(867)).minus(_this.f31);
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source1 = _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1(true, new BigNumber((_dafny.Seq.UnicodeFromString("jwdbhn")).length))));
      if (_source1.is_DC1) {
        let _162___mcc_h0 = (_source1).cf1;
        let _163___mcc_h1 = (_source1).cf2;
        let _164_cf2 = _163___mcc_h1;
        let _165_cf1 = _162___mcc_h0;
        return !(_165_cf1);
      } else if (_source1.is_DC0) {
        let _166___mcc_h2 = (_source1).cf0;
        let _167_cf0 = _166___mcc_h2;
        return _167_cf0;
      } else {
        let _168___mcc_h3 = (_source1).cf3;
        let _169_cf3 = _168___mcc_h3;
        return false;
      }
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = false;
      r2 = p0;
      let _170_v0;
      _170_v0 = new _dafny.CodePoint('h'.codePointAt(0));
      let _171_v1;
      _171_v1 = _dafny.Seq.of(_170_v0, _170_v0);
      _171_v1 = _171_v1;
      let _172_v2;
      let _nw37 = Array((new BigNumber(6)).toNumber()).fill([]);
      _172_v2 = _nw37;
      let _173_v3;
      let _nw38 = Array((new BigNumber(10)).toNumber()).fill(false);
      _173_v3 = _nw38;
      let _174_v4;
      _174_v4 = _173_v3;
      let _index25 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_172_v2).length));
      (_172_v2)[_index25] = _173_v3;
      let _index26 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_172_v2).length));
      (_172_v2)[_index26] = _174_v4;
      let _175_v5;
      _175_v5 = _module.D0.create_DC2(_module.D0.create_DC0(p1));
      let _176_v6;
      _176_v6 = _dafny.Set.fromElements(_module.__default.fm24((_dafny.ZERO).minus((_this).f23), p0, p0, globalState), _175_v5);
      (globalState).f1 = ((!(!(p1))) ? (!(_module.__default.fm23(_this.f31, globalState)).equals(_176_v6)) : (true));
      let _index27 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_173_v3).length));
      (_173_v3)[_index27] = p1;
      let _index28 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_173_v3).length));
      (_173_v3)[_index28] = p1;
      let _177_v7;
      _177_v7 = _dafny.MultiSet.fromElements(p1, p1);
      let _178_v8;
      _178_v8 = _dafny.Seq.of(new BigNumber(44), _this.f31, (_this).f23, _this.f31);
      let _179_v9;
      _179_v9 = _dafny.Map.Empty.slice().updateUnsafe((_173_v3)[_module.__default.safeIndex(new BigNumber(21), new BigNumber((_173_v3).length))],_171_v1);
      let _180_v10;
      _180_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(_171_v1, p0, _this.f31, _178_v8, globalState),(_this).f23);
      let _181_v11;
      _181_v11 = _dafny.Seq.of(p0);
      let _182_v12;
      _182_v12 = _module.D4.create_DC12((_this).f23, _this.f31);
      let _183_v13;
      let _nw39 = Array((new BigNumber(26)).toNumber());
      _nw39[(_dafny.ZERO).toNumber()] = (_this).f23;
      _nw39[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f31));
      _nw39[(new BigNumber(2)).toNumber()] = _this.f31;
      _nw39[(new BigNumber(3)).toNumber()] = _this.f31;
      _nw39[(new BigNumber(4)).toNumber()] = (_this).f23;
      _nw39[(new BigNumber(5)).toNumber()] = (_this).f23;
      _nw39[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_178_v8).length));
      _nw39[(new BigNumber(7)).toNumber()] = new BigNumber((_179_v9).length);
      _nw39[(new BigNumber(8)).toNumber()] = (_this).f23;
      _nw39[(new BigNumber(9)).toNumber()] = (_178_v8)[_module.__default.safeIndex((_this).f23, new BigNumber((_178_v8).length))];
      _nw39[(new BigNumber(10)).toNumber()] = (_this).f23;
      _nw39[(new BigNumber(11)).toNumber()] = (_this).f23;
      _nw39[(new BigNumber(12)).toNumber()] = new BigNumber(656);
      _nw39[(new BigNumber(13)).toNumber()] = (_this).f23;
      _nw39[(new BigNumber(14)).toNumber()] = (((_180_v10).contains(p0)) ? ((_180_v10).get(p0)) : (_this.f31));
      _nw39[(new BigNumber(15)).toNumber()] = _module.__default.fm1(new BigNumber((_171_v1).length), p0, globalState);
      _nw39[(new BigNumber(16)).toNumber()] = new BigNumber(678);
      _nw39[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ft")).length);
      _nw39[(new BigNumber(18)).toNumber()] = _this.f31;
      _nw39[(new BigNumber(19)).toNumber()] = (_this).f23;
      _nw39[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(_this.f31);
      _nw39[(new BigNumber(21)).toNumber()] = _this.f31;
      _nw39[(new BigNumber(22)).toNumber()] = _this.f31;
      _nw39[(new BigNumber(23)).toNumber()] = new BigNumber((_181_v11).length);
      _nw39[(new BigNumber(24)).toNumber()] = new BigNumber(970);
      _nw39[(new BigNumber(25)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_182_v12).dtor_cf19,new BigNumber((_177_v7).cardinality()))).length);
      _183_v13 = _nw39;
      let _184_v14;
      _184_v14 = _dafny.Seq.of(_183_v13, _183_v13);
      let _rhs44 = (_184_v14)[_module.__default.safeIndex(_this.f31, new BigNumber((_184_v14).length))];
      let _rhs45 = (_177_v7).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(p1), _dafny.Seq.of((_173_v3)[_module.__default.safeIndex(new BigNumber(21), new BigNumber((_173_v3).length))]))));
      let _rhs46 = (new BigNumber((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(61), new BigNumber(671))) {
          let _185_v15 = _compr_10;
          if (((new BigNumber(61)).isLessThanOrEqualTo(_185_v15)) && ((_185_v15).isLessThan(new BigNumber(671)))) {
            _coll10.push([(_185_v15).minus(new BigNumber(758)),_dafny.MultiSet.fromElements(_170_v0, _170_v0)]);
          }
        }
        return _coll10;
      }()).length)).minus((_this.f31).plus((_this).f23));
      let _rhs47 = _module.__default.safeModuloInt(new BigNumber((((p1) ? (_180_v10) : (_180_v10))).length), (_this.f31).multipliedBy(_this.f31));
      let _rhs48 = (_this).f23;
      let _lhs37 = globalState;
      let _lhs38 = globalState;
      let _lhs39 = globalState;
      r1 = _rhs44;
      _177_v7 = _rhs45;
      _lhs37.f17 = _rhs46;
      _lhs38.f14 = _rhs47;
      _lhs39.f14 = _rhs48;
      r0 = ((_dafny.ZERO).minus(_this.f31)).plus((_this).f23);
      r1 = _183_v13;
      let _186_v16;
      _186_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f31,new BigNumber(151));
      let _187_v17;
      _187_v17 = _dafny.Map.Empty.slice().updateUnsafe(_173_v3,(_this).fm3(_171_v1, true, (_this).f23, _dafny.Seq.of(new BigNumber((_186_v16).length), _this.f31), globalState));
      r2 = !((_187_v17).Merge(_187_v17)).contains(_173_v3);
      return [r0, r1, r2];
    }
    m10(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let r3 = false;
      r2 = p2;
      let _188_v0;
      _188_v0 = _dafny.Seq.UnicodeFromString("mevk");
      _188_v0 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(717)), function (_189_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), _dafny.Seq.Concat(_188_v0, _188_v0));
      let _190_v1;
      _190_v1 = new _dafny.CodePoint('w'.codePointAt(0));
      let _191_v2;
      _191_v2 = _dafny.Seq.of(p0);
      let _192_v3;
      _192_v3 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
      let _193_v4;
      _193_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _194_v5;
      _194_v5 = _module.D3.create_DC10(p0, _193_v4);
      let _195_v6;
      let _nw40 = Array((new BigNumber(26)).toNumber());
      _nw40[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_188_v0, _188_v0);
      _nw40[(_dafny.ONE).toNumber()] = _188_v0;
      _nw40[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_188_v0, _module.__default.safeIndex(_this.f31, new BigNumber((_188_v0).length)), _190_v1);
      _nw40[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("rkd");
      _nw40[(new BigNumber(4)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("aai");
      _nw40[(new BigNumber(6)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(7)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(8)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(9)).toNumber()] = _module.__default.fm25(p0, p0, _190_v1, globalState);
      _nw40[(new BigNumber(10)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(58)), ((_196_v1) => function (_197_i1) {
        return _196_v1;
      })(_190_v1));
      _nw40[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_188_v0, _188_v0);
      _nw40[(new BigNumber(13)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(14)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(285)), ((_198_v1) => function (_199_i2) {
        return _198_v1;
      })(_190_v1)), _188_v0);
      _nw40[(new BigNumber(16)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(17)).toNumber()] = _module.__default.fm25((_191_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_191_v2).length))], p0, _module.__default.fm26(new BigNumber((_192_v3).length), p0, globalState), globalState);
      _nw40[(new BigNumber(18)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(19)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(20)).toNumber()] = _module.__default.fm25((((_193_v4).contains(p0)) ? ((_193_v4).get(p0)) : (!(p0))), !((_194_v5).dtor_cf16), _190_v1, globalState);
      _nw40[(new BigNumber(21)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(22)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(23)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(24)).toNumber()] = _188_v0;
      _nw40[(new BigNumber(25)).toNumber()] = _188_v0;
      _195_v6 = _nw40;
      let _rhs49 = !(!((_this).f23).isEqualTo(p2));
      let _rhs50 = _195_v6;
      let _rhs51 = p0;
      let _lhs40 = globalState;
      _lhs40.f1 = _rhs49;
      _195_v6 = _rhs50;
      r3 = _rhs51;
      let _200_v7;
      _200_v7 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_188_v0, _188_v0),_193_v4);
      let _201_v8;
      _201_v8 = _dafny.MultiSet.fromElements(p0, p0, p0);
      let _202_v9;
      _202_v9 = _dafny.MultiSet.fromElements(_module.__default.fm1(p2, p0, globalState));
      _200_v7 = (_module.__default.fm27(_this.f31, _201_v8, new BigNumber((_202_v9).cardinality()), p2, globalState)).Merge(_200_v7);
      let _203_v10;
      _203_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f23);
      let _204_v11;
      _204_v11 = _dafny.Seq.of(_this.f31, p3);
      let _hi1 = new BigNumber((_204_v11).length);
      for (let _205_i3 = (new BigNumber((_dafny.Seq.UnicodeFromString("hk")).length)).plus((((_203_v10).contains((((_193_v4).contains(false)) ? ((_193_v4).get(false)) : (!(p0))))) ? ((_203_v10).get((((_193_v4).contains(false)) ? ((_193_v4).get(false)) : (!(p0))))) : ((_this).f23))); _205_i3.isLessThan(_hi1); _205_i3 = _205_i3.plus(_dafny.ONE)) {
        let _206_v12;
        let _nw41 = Array((new BigNumber(5)).toNumber()).fill(false);
        _206_v12 = _nw41;
        let _index29 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length));
        (_206_v12)[_index29] = p0;
        let _207_v13;
        let _nw42 = Array((new BigNumber(6)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = _190_v1;
        _nw42[(_dafny.ONE).toNumber()] = _190_v1;
        _nw42[(new BigNumber(2)).toNumber()] = ((p0) ? (_190_v1) : (_190_v1));
        _nw42[(new BigNumber(3)).toNumber()] = _190_v1;
        _nw42[(new BigNumber(4)).toNumber()] = _190_v1;
        _nw42[(new BigNumber(5)).toNumber()] = _190_v1;
        _207_v13 = _nw42;
        let _index30 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length));
        let _rhs52 = _191_v2;
        let _rhs53 = p0;
        let _rhs54 = _dafny.Seq.Concat(_188_v0, ((p0) ? (_188_v0) : (_188_v0)));
        let _rhs55 = _207_v13;
        let _lhs41 = _206_v12;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length));
        _191_v2 = _rhs52;
        _lhs41[_lhs42] = _rhs53;
        _188_v0 = _rhs54;
        _207_v13 = _rhs55;
        let _index31 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length));
        let _index32 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length));
        let _rhs56 = !(!_dafny.Seq.contains(_191_v2, (_191_v2)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ahkl")).length), new BigNumber((_191_v2).length))]));
        let _rhs57 = p0;
        let _rhs58 = (_206_v12)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length))];
        let _rhs59 = (!((_206_v12)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length))]) || (false)) && ((_206_v12)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length))]);
        let _lhs43 = globalState;
        let _lhs44 = _206_v12;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length));
        let _lhs46 = globalState;
        let _lhs47 = _206_v12;
        let _lhs48 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_206_v12).length));
        _lhs43.f15 = _rhs56;
        _lhs44[_lhs45] = _rhs57;
        _lhs46.f1 = _rhs58;
        _lhs47[_lhs48] = _rhs59;
        (globalState).f8 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_this.f31, _205_i3));
        (globalState).f17 = (_this).fm2(globalState);
      }
      (_this).f31 = _module.__default.safeDivisionInt((_this.f31).minus((_dafny.ZERO).minus((_this).f23)), new BigNumber((_201_v8).cardinality()));
      r0 = _202_v9;
      let _208_v14;
      let _nw43 = Array((new BigNumber(3)).toNumber());
      _nw43[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).contains(p0);
      _nw43[(_dafny.ONE).toNumber()] = p0;
      _nw43[(new BigNumber(2)).toNumber()] = p0;
      _208_v14 = _nw43;
      r1 = _208_v14;
      let _209_v15;
      _209_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_188_v0).length),new BigNumber(497));
      r2 = (((_202_v9).contains((_dafny.ZERO).minus(p3))) ? ((_202_v9).get((_dafny.ZERO).minus(p3))) : ((((_209_v15).contains((_this).fm2(globalState))) ? ((_209_v15).get((_this).fm2(globalState))) : ((_dafny.ZERO).minus(_this.f31)))));
      r3 = p0;
      return [r0, r1, r2, r3];
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f24 = false;
      this._f25 = false;
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    set f24(value) {
      let _this = this;
      _this._f24 = value;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f24, f25) {
      let _this = this;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_this.f24))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_dafny.Map.Empty.slice().updateUnsafe((_this).f25,true))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Map.Empty.slice().updateUnsafe(_this.f24,(_this).f25))));
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return ((new BigNumber((_dafny.Seq.of(true, (_this).f25)).length)).multipliedBy(new BigNumber(953))).plus(((_this.f24) ? (new BigNumber(954)) : (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(36), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("lidbxgdo")).length)))).cardinality()))));
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = undefined;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _210_v0;
      _210_v0 = _dafny.Seq.UnicodeFromString("pqbcphfi");
      let _211_v1;
      _211_v1 = _module.D3.create_DC8(_210_v0);
      let _source2 = _211_v1;
      if (_source2.is_DC9) {
        let _212___mcc_h0 = (_source2).cf15;
        let _213_cf15 = _212___mcc_h0;
        let _214_v2;
        let _nw44 = Array((new BigNumber(26)).toNumber()).fill(false);
        _214_v2 = _nw44;
        let _index33 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_214_v2).length));
        (_214_v2)[_index33] = (_this).f25;
        let _215_v3;
        _215_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_this.f24);
        let _216_v4;
        _216_v4 = _dafny.Seq.of(_213_cf15, new BigNumber((_215_v3).length), _213_cf15);
        let _217_v5;
        _217_v5 = _dafny.Set.fromElements(_this.f24);
        let _218_v6;
        _218_v6 = new _dafny.CodePoint('l'.codePointAt(0));
        let _index34 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_214_v2).length));
        let _rhs60 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), ((_219_cf15) => function (_220_i0) {
          return _219_cf15;
        })(_213_cf15)), _216_v4), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_216_v4).length), _213_cf15, _213_cf15), _216_v4));
        let _rhs61 = ((_217_v5).Union(_217_v5)).IsProperSubsetOf(_217_v5);
        let _rhs62 = _dafny.Seq.update(_210_v0, _module.__default.safeIndex(_213_cf15, new BigNumber((_210_v0).length)), _218_v6);
        let _rhs63 = !((_this).f25) || (_this.f24);
        let _rhs64 = _210_v0;
        let _lhs49 = globalState;
        let _lhs50 = _214_v2;
        let _lhs51 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_214_v2).length));
        let _lhs52 = globalState;
        _lhs49.f6 = _rhs60;
        _lhs50[_lhs51] = _rhs61;
        _210_v0 = _rhs62;
        _lhs52.f1 = _rhs63;
        _210_v0 = _rhs64;
        let _221_v7;
        let _nw45 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
        _221_v7 = _nw45;
        let _index35 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_221_v7).length));
        (_221_v7)[_index35] = _dafny.MultiSet.fromElements(_213_cf15);
        let _222_v8;
        _222_v8 = _dafny.MultiSet.fromElements(_213_cf15);
        let _223_v9;
        _223_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_213_cf15);
        let _index36 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_221_v7).length));
        (_221_v7)[_index36] = ((_222_v8).Union(_222_v8)).Difference((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_223_v9).length),_this.f24)).length))).Union(_222_v8));
        if (false) {
          (globalState).f8 = _213_cf15;
          let _224_v10;
          let _nw46 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
          _224_v10 = _nw46;
          let _index37 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_224_v10).length));
          (_224_v10)[_index37] = _213_cf15;
          let _index38 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_224_v10).length));
          (_224_v10)[_index38] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_213_cf15));
          (globalState).f17 = _module.__default.fm1(((_224_v10)[_module.__default.safeIndex(new BigNumber(521), new BigNumber((_224_v10).length))]).minus(new BigNumber(582)), ((_this).f25) === ((_this).f25), globalState);
          let _225_v11;
          _225_v11 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_216_v4, _216_v4));
          _225_v11 = _225_v11;
          let _226_v12;
          _226_v12 = _dafny.MultiSet.fromElements((_214_v2)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_214_v2).length))], _this.f24);
          r2 = ((_226_v12).Intersect(_module.__default.fm22(_213_cf15, (_214_v2)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_214_v2).length))], _this.f24, (_224_v10)[_module.__default.safeIndex(new BigNumber(521), new BigNumber((_224_v10).length))], globalState))).IsDisjointFrom(_dafny.MultiSet.fromElements((_214_v2)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_214_v2).length))], _this.f24, (_this).f25, true));
        } else {
          (globalState).f1 = _this.f24;
          (_this).f24 = ((_213_cf15).plus(_213_cf15)).isLessThan(_module.__default.safeModuloInt(new BigNumber(615), _213_cf15));
          r2 = (_214_v2)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_214_v2).length))];
          _210_v0 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(517)), ((_227_v6) => function (_228_i1) {
            return _227_v6;
          })(_218_v6)), _210_v0);
          let _index39 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_214_v2).length));
          (_214_v2)[_index39] = (false) || ((_this).f25);
        }
        (globalState).f5 = (_dafny.ZERO).minus(_213_cf15);
      } else if (_source2.is_DC10) {
        let _229___mcc_h1 = (_source2).cf16;
        let _230___mcc_h2 = (_source2).cf17;
        let _231_cf17 = _230___mcc_h2;
        let _232_cf16 = _229___mcc_h1;
        let _233_v13;
        let _nw47 = Array((new BigNumber(12)).toNumber()).fill(false);
        _233_v13 = _nw47;
        let _234_v14;
        _234_v14 = _233_v13;
        _234_v14 = _234_v14;
        let _235_v16;
        let _init2 = ((_236_cf16) => function (_237_i2) {
          return function () {
            let _coll11 = new _dafny.Set();
            for (const _compr_11 of _dafny.IntegerRange(new BigNumber(380), new BigNumber(719))) {
              let _238_v15 = _compr_11;
              if (((new BigNumber(380)).isLessThanOrEqualTo(_238_v15)) && ((_238_v15).isLessThan(new BigNumber(719)))) {
                _coll11.add(_module.__default.safeDivisionInt(_238_v15, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_236_cf16,_dafny.MultiSet.fromElements(_this.f24))).length)));
              }
            }
            return _coll11;
          }();
        })(_232_cf16);
        let _nw48 = Array((new BigNumber(9)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw48.length); _i0_2++) {
          _nw48[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _235_v16 = _nw48;
        let _239_v17;
        _239_v17 = new BigNumber(721);
        let _240_v18;
        _240_v18 = _dafny.Map.Empty.slice().updateUnsafe(_232_cf16,_239_v17);
        let _241_v19;
        _241_v19 = _dafny.Set.fromElements(new BigNumber((_240_v18).length), (_dafny.ZERO).minus(_239_v17), new BigNumber((_dafny.Set.fromElements(_232_cf16)).length), new BigNumber(63));
        let _index40 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_235_v16).length));
        (_235_v16)[_index40] = _241_v19;
        let _index41 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_235_v16).length));
        (_235_v16)[_index41] = _241_v19;
        let _242_v20;
        let _nw49 = Array((new BigNumber(22)).toNumber()).fill(_dafny.MultiSet.Empty);
        _242_v20 = _nw49;
        let _243_v21;
        _243_v21 = _dafny.Seq.of(_242_v20);
        let _244_v23;
        _244_v23 = _dafny.Seq.of(_242_v20, _242_v20, (_243_v21)[_module.__default.safeIndex(new BigNumber((function () {
          let _coll12 = new _dafny.Map();
          for (const _compr_12 of _dafny.IntegerRange(new BigNumber(628), new BigNumber(146))) {
            let _245_v22 = _compr_12;
            if (((new BigNumber(628)).isLessThanOrEqualTo(_245_v22)) && ((_245_v22).isLessThan(new BigNumber(146)))) {
              _coll12.push([_module.__default.safeDivisionInt(_245_v22, new BigNumber((_dafny.Seq.of(_232_cf16)).length)),_232_cf16]);
            }
          }
          return _coll12;
        }()).length), new BigNumber((_243_v21).length))]);
        _242_v20 = (_244_v23)[_module.__default.safeIndex((_239_v17).plus(new BigNumber(155)), new BigNumber((_244_v23).length))];
        let _246_v24;
        _246_v24 = _dafny.MultiSet.fromElements((_this).f25);
        let _rhs65 = _239_v17;
        let _rhs66 = false;
        let _rhs67 = (_dafny.MultiSet.fromElements(_232_cf16, _232_cf16)).IsSubsetOf((_module.D7.create_DC19(_246_v24)).dtor_cf34);
        let _lhs53 = globalState;
        let _lhs54 = globalState;
        _lhs53.f5 = _rhs65;
        _lhs54.f15 = _rhs66;
        r2 = _rhs67;
      } else {
        let _247___mcc_h3 = (_source2).cf14;
        let _248_cf14 = _247___mcc_h3;
        let _249_v25;
        _249_v25 = new _dafny.CodePoint('l'.codePointAt(0));
        _249_v25 = _249_v25;
        let _250_v26;
        let _init3 = function (_251_i3) {
          return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(552)), _dafny.Seq.of(new BigNumber(-777), new BigNumber((_dafny.Seq.of(_this.f24, _this.f24)).length)));
        };
        let _nw50 = Array((new BigNumber(10)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw50.length); _i0_3++) {
          _nw50[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _250_v26 = _nw50;
        let _index42 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_250_v26).length));
        (_250_v26)[_index42] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(396)), function (_252_i4) {
          return new BigNumber(-996);
        });
        let _253_v27;
        _253_v27 = new BigNumber(40);
        let _254_v28;
        _254_v28 = _dafny.Seq.of(_253_v27);
        let _index43 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_250_v26).length));
        (_250_v26)[_index43] = _254_v28;
        let _255_v29;
        _255_v29 = _dafny.Seq.of(_this.f24, _this.f24, !((_this).f25));
        let _256_v30;
        _256_v30 = _dafny.MultiSet.fromElements(_this.f24, (_this).f25);
        let _257_v31;
        _257_v31 = _dafny.Set.fromElements(_253_v27);
        (_this).m9(_249_v25, !((_this).f25) || (!(_this.f24)), (_255_v29)[_module.__default.safeIndex(_253_v27, new BigNumber((_255_v29).length))], _module.__default.safeModuloInt(new BigNumber((_256_v30).cardinality()), new BigNumber((_257_v31).length)), globalState);
        let _258_v32;
        let _nw51 = Array((new BigNumber(25)).toNumber());
        _nw51[(_dafny.ZERO).toNumber()] = _this.f24;
        _nw51[(_dafny.ONE).toNumber()] = true;
        _nw51[(new BigNumber(2)).toNumber()] = _this.f24;
        _nw51[(new BigNumber(3)).toNumber()] = _this.f24;
        _nw51[(new BigNumber(4)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(5)).toNumber()] = _this.f24;
        _nw51[(new BigNumber(6)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(7)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(8)).toNumber()] = _this.f24;
        _nw51[(new BigNumber(9)).toNumber()] = !(_this.f24);
        _nw51[(new BigNumber(10)).toNumber()] = !((_this).f25);
        _nw51[(new BigNumber(11)).toNumber()] = _this.f24;
        _nw51[(new BigNumber(12)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(13)).toNumber()] = true;
        _nw51[(new BigNumber(14)).toNumber()] = (_255_v29)[_module.__default.safeIndex(_253_v27, new BigNumber((_255_v29).length))];
        _nw51[(new BigNumber(15)).toNumber()] = _this.f24;
        _nw51[(new BigNumber(16)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(17)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(18)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(19)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(20)).toNumber()] = _this.f24;
        _nw51[(new BigNumber(21)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(22)).toNumber()] = (_this).f25;
        _nw51[(new BigNumber(23)).toNumber()] = _this.f24;
        _nw51[(new BigNumber(24)).toNumber()] = (_this).f25;
        _258_v32 = _nw51;
        let _259_v33;
        _259_v33 = _dafny.Set.fromElements(_258_v32);
        let _260_v34;
        let _nw52 = new _module.C0();
        _nw52.__ctor((_259_v33).Intersect(_259_v33), _module.__default.fm1(_253_v27, (_this).f25, globalState), _253_v27);
        _260_v34 = _nw52;
        r1 = _260_v34;
      }
      let _261_v35;
      _261_v35 = _dafny.Seq.of((_this).f25);
      let _262_v36;
      _262_v36 = new BigNumber(-850);
      let _263_v37;
      _263_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
      let _264_v38;
      _264_v38 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_262_v36);
      let _265_v39;
      _265_v39 = _module.D1.create_DC5((_261_v35)[_module.__default.safeIndex(_262_v36, new BigNumber((_261_v35).length))], _263_v37, _262_v36, _262_v36, _264_v38);
      let _266_v40;
      _266_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_265_v39);
      let _267_v41;
      _267_v41 = _dafny.MultiSet.fromElements(_module.__default.fm28(_262_v36, _210_v0, _262_v36, globalState), _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_265_v39), _266_v40);
      let _268_i5;
      _268_i5 = _dafny.ZERO;
      L1: {
        while (((_dafny.MultiSet.fromElements(_266_v40, _266_v40)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_265_v39))))).IsProperSubsetOf((_dafny.MultiSet.fromElements(_266_v40)).Difference(_267_v41))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_268_i5)) {
              break L1;
            }
            _268_i5 = (_268_i5).plus(_dafny.ONE);
            let _269_v42;
            let _init4 = function (_270_i6) {
              return _this.f24;
            };
            let _nw53 = Array((new BigNumber(29)).toNumber());
            for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw53.length); _i0_4++) {
              _nw53[_i0_4] = _init4(new BigNumber(_i0_4));
            }
            _269_v42 = _nw53;
            let _271_v43;
            _271_v43 = _dafny.Set.fromElements(_269_v42, _269_v42);
            let _272_v44;
            _272_v44 = _dafny.Map.Empty.slice().updateUnsafe(_262_v36,_dafny.Seq.UnicodeFromString("wvrqm"));
            let _273_v45;
            let _nw54 = new _module.C0();
            _nw54.__ctor(_271_v43, _262_v36, _module.__default.safeDivisionInt(_262_v36, new BigNumber((_272_v44).length)));
            _273_v45 = _nw54;
            let _274_v46;
            _274_v46 = _dafny.Seq.of(_262_v36);
            let _275_v47;
            _275_v47 = _dafny.Set.fromElements(new BigNumber((_274_v46).length));
            let _276_v48;
            _276_v48 = _dafny.MultiSet.fromElements(_275_v47);
            let _277_v49;
            _277_v49 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_262_v36))).Difference(_276_v48),(_273_v45.f31).minus((_this).fm5(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(292),_269_v42)).length), _273_v45.f31, globalState)));
            let _278_v50;
            _278_v50 = _dafny.Seq.of(_275_v47, _275_v47, _275_v47);
            let _279_v51;
            _279_v51 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f24),_269_v42);
            let _280_v52;
            _280_v52 = _dafny.MultiSet.fromElements(_269_v42, (((_279_v51).contains(_this.f24)) ? ((_279_v51).get(_this.f24)) : (_269_v42)), _269_v42, _269_v42);
            _277_v49 = (_277_v49).update(_dafny.MultiSet.FromArray(_278_v50), (((_280_v52).contains(_269_v42)) ? ((_280_v52).get(_269_v42)) : ((_dafny.ZERO).minus((_this).fm5(_273_v45.f31, _273_v45.f31, globalState)))));
            r3 = _273_v45.f31;
            let _281_v53;
            _281_v53 = _module.D0.create_DC1(_this.f24, _262_v36);
            let _282_v54;
            _282_v54 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_274_v46, _module.__default.safeIndex((_281_v53).dtor_cf2, new BigNumber((_274_v46).length)), _273_v45.f31),(_this).f25);
            let _283_v55;
            _283_v55 = _dafny.MultiSet.fromElements(_module.__default.fm22(_273_v45.f31, (_this).f25, !(true), _273_v45.f31, globalState), _module.__default.fm22(_273_v45.f31, false, (_this).f25, new BigNumber((_282_v54).length), globalState));
            _283_v55 = (_283_v55).Intersect((_283_v55).Union(_283_v55));
          }
        }
      }
      let _284_v56;
      let _nw55 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
      _284_v56 = _nw55;
      let _285_v57;
      _285_v57 = _dafny.Map.Empty.slice().updateUnsafe(_262_v36,(_this).f25);
      let _index44 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_284_v56).length));
      (_284_v56)[_index44] = _285_v57;
      let _index45 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_284_v56).length));
      (_284_v56)[_index45] = _285_v57;
      if (_this.f24) {
        let _286_v58;
        let _init5 = function (_287_i7) {
          return (((_this).f25) ? (new _dafny.CodePoint('g'.codePointAt(0))) : (new _dafny.CodePoint('c'.codePointAt(0))));
        };
        let _nw56 = Array((new BigNumber(21)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw56.length); _i0_5++) {
          _nw56[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _286_v58 = _nw56;
        let _288_v59;
        _288_v59 = new _dafny.CodePoint('a'.codePointAt(0));
        let _index46 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_286_v58).length));
        (_286_v58)[_index46] = (((_this).f25) ? (_288_v59) : (_288_v59));
        let _289_v60;
        _289_v60 = _dafny.Seq.of(new BigNumber(973));
        let _290_v61;
        _290_v61 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_289_v60, _289_v60));
        let _index47 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_286_v58).length));
        let _index48 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_284_v56).length));
        let _rhs68 = (((_290_v61).contains(_dafny.Seq.Concat(_289_v60, _289_v60))) ? ((_290_v61).get(_dafny.Seq.Concat(_289_v60, _289_v60))) : (_262_v36));
        let _rhs69 = _288_v59;
        let _rhs70 = (_284_v56)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_284_v56).length))];
        let _lhs55 = globalState;
        let _lhs56 = _286_v58;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_286_v58).length));
        let _lhs58 = _284_v56;
        let _lhs59 = _module.__default.safeIndex(new BigNumber(621), new BigNumber((_284_v56).length));
        _lhs55.f8 = _rhs68;
        _lhs56[_lhs57] = _rhs69;
        _lhs58[_lhs59] = _rhs70;
        let _291_v62;
        let _nw57 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _291_v62 = _nw57;
        let _rhs71 = (new BigNumber((_261_v35).length)).multipliedBy((_262_v36).plus(_262_v36));
        let _rhs72 = _291_v62;
        let _rhs73 = (_261_v35)[_module.__default.safeIndex(_module.__default.fm1(_262_v36, (_this).f25, globalState), new BigNumber((_261_v35).length))];
        let _lhs60 = globalState;
        let _lhs61 = globalState;
        _lhs60.f14 = _rhs71;
        _291_v62 = _rhs72;
        _lhs61.f9 = _rhs73;
        (globalState).f8 = _262_v36;
        let _292_v63;
        _292_v63 = _module.D1.create_DC4((_this).f25, _this.f24);
        let _293_v64;
        let _nw58 = Array((new BigNumber(4)).toNumber());
        _nw58[(_dafny.ZERO).toNumber()] = (_292_v63).dtor_cf5;
        _nw58[(_dafny.ONE).toNumber()] = _this.f24;
        _nw58[(new BigNumber(2)).toNumber()] = _this.f24;
        _nw58[(new BigNumber(3)).toNumber()] = !((_this).f25) || ((_this).f25);
        _293_v64 = _nw58;
        let _index49 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_293_v64).length));
        (_293_v64)[_index49] = _this.f24;
        let _294_v65;
        _294_v65 = _dafny.Set.fromElements(_this.f24, (_this).f25);
        let _index50 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_293_v64).length));
        (_293_v64)[_index50] = (_294_v65).IsSubsetOf(_294_v65);
        (globalState).f15 = (_this).f25;
      } else {
        let _295_v66;
        _295_v66 = new _dafny.CodePoint('p'.codePointAt(0));
        _295_v66 = _295_v66;
        let _296_v67;
        let _nw59 = Array((new BigNumber(25)).toNumber());
        _nw59[(_dafny.ZERO).toNumber()] = _this.f24;
        _nw59[(_dafny.ONE).toNumber()] = !(false);
        _nw59[(new BigNumber(2)).toNumber()] = true;
        _nw59[(new BigNumber(3)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(4)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(5)).toNumber()] = true;
        _nw59[(new BigNumber(6)).toNumber()] = true;
        _nw59[(new BigNumber(7)).toNumber()] = _this.f24;
        _nw59[(new BigNumber(8)).toNumber()] = true;
        _nw59[(new BigNumber(9)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(10)).toNumber()] = false;
        _nw59[(new BigNumber(11)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(12)).toNumber()] = _module.__default.fm29(_262_v36, _262_v36, (_this).f25, globalState);
        _nw59[(new BigNumber(13)).toNumber()] = _this.f24;
        _nw59[(new BigNumber(14)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(15)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(16)).toNumber()] = _this.f24;
        _nw59[(new BigNumber(17)).toNumber()] = _this.f24;
        _nw59[(new BigNumber(18)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(19)).toNumber()] = _this.f24;
        _nw59[(new BigNumber(20)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(21)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(22)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(23)).toNumber()] = (_this).f25;
        _nw59[(new BigNumber(24)).toNumber()] = !(_this.f24);
        _296_v67 = _nw59;
        let _297_v68;
        _297_v68 = _dafny.Set.fromElements(_296_v67, _296_v67);
        let _298_v69;
        let _nw60 = new _module.C0();
        _nw60.__ctor(_297_v68, new BigNumber((_261_v35).length), (_262_v36).multipliedBy(_262_v36));
        _298_v69 = _nw60;
        let _299_v70;
        let _init6 = function (_300_i8) {
          return _module.D0.create_DC2(_module.D0.create_DC0((_this).f25));
        };
        let _nw61 = Array((new BigNumber(29)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw61.length); _i0_6++) {
          _nw61[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _299_v70 = _nw61;
        let _301_v71;
        _301_v71 = _dafny.Map.Empty.slice().updateUnsafe(_299_v70,_295_v66);
        let _302_v72;
        let _nw62 = new _module.C0();
        _nw62.__ctor((_298_v69).f30, _262_v36, _298_v69.f31);
        _302_v72 = _nw62;
        let _303_v73;
        _303_v73 = _module.D3.create_DC9(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_302_v72,_298_v69.f31)).length));
        let _304_v74;
        _304_v74 = _dafny.Map.Empty.slice().updateUnsafe(_262_v36,_210_v0);
        let _305_v75;
        _305_v75 = _dafny.MultiSet.fromElements(_262_v36, new BigNumber(-199));
        let _306_v76;
        _306_v76 = _dafny.MultiSet.fromElements(_this.f24);
        let _307_v77;
        _307_v77 = _dafny.Seq.of(_262_v36, _262_v36);
        let _308_v79;
        _308_v79 = _dafny.Set.fromElements(new BigNumber(978));
        let _309_v80;
        let _nw63 = Array((new BigNumber(25)).toNumber());
        _nw63[(_dafny.ZERO).toNumber()] = (_302_v72).f23;
        _nw63[(_dafny.ONE).toNumber()] = _298_v69.f31;
        _nw63[(new BigNumber(2)).toNumber()] = new BigNumber(966);
        _nw63[(new BigNumber(3)).toNumber()] = _298_v69.f31;
        _nw63[(new BigNumber(4)).toNumber()] = (_302_v72).f23;
        _nw63[(new BigNumber(5)).toNumber()] = new BigNumber((_261_v35).length);
        _nw63[(new BigNumber(6)).toNumber()] = _262_v36;
        _nw63[(new BigNumber(7)).toNumber()] = _262_v36;
        _nw63[(new BigNumber(8)).toNumber()] = _298_v69.f31;
        _nw63[(new BigNumber(9)).toNumber()] = new BigNumber((_304_v74).length);
        _nw63[(new BigNumber(10)).toNumber()] = _262_v36;
        _nw63[(new BigNumber(11)).toNumber()] = new BigNumber((_305_v75).cardinality());
        _nw63[(new BigNumber(12)).toNumber()] = (_302_v72).f23;
        _nw63[(new BigNumber(13)).toNumber()] = _262_v36;
        _nw63[(new BigNumber(14)).toNumber()] = (((_306_v76).contains((_this).f25)) ? ((_306_v76).get((_this).f25)) : (new BigNumber((_307_v77).length)));
        _nw63[(new BigNumber(15)).toNumber()] = (_302_v72).f23;
        _nw63[(new BigNumber(16)).toNumber()] = new BigNumber((_307_v77).length);
        _nw63[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of (_307_v77).Elements) {
            let _310_v78 = _compr_13;
            if (_dafny.Seq.contains(_307_v77, _310_v78)) {
              _coll13.add((_310_v78).minus(new BigNumber(421)));
            }
          }
          return _coll13;
        }(), _308_v79)).cardinality()),_this.f24)).length);
        _nw63[(new BigNumber(18)).toNumber()] = (_307_v77)[_module.__default.safeIndex(_262_v36, new BigNumber((_307_v77).length))];
        _nw63[(new BigNumber(19)).toNumber()] = _262_v36;
        _nw63[(new BigNumber(20)).toNumber()] = _298_v69.f31;
        _nw63[(new BigNumber(21)).toNumber()] = new BigNumber((_210_v0).length);
        _nw63[(new BigNumber(22)).toNumber()] = _262_v36;
        _nw63[(new BigNumber(23)).toNumber()] = _262_v36;
        _nw63[(new BigNumber(24)).toNumber()] = _262_v36;
        _309_v80 = _nw63;
        let _311_v81;
        _311_v81 = _module.D8.create_DC22(_309_v80);
        let _312_v82;
        _312_v82 = _dafny.Map.Empty.slice().updateUnsafe(_303_v73,(_311_v81).dtor_cf39);
        let _313_v83;
        _313_v83 = _module.D5.create_DC15((_302_v72).f23, _295_v66, _dafny.Map.Empty.slice().updateUnsafe(_299_v70,_295_v66), _312_v82);
        let _314_v84;
        _314_v84 = _module.D5.create_DC15(new BigNumber(589), _295_v66, ((_301_v71).update(_299_v70, _295_v66)).Merge(_301_v71), (_dafny.Map.Empty.slice().updateUnsafe(_303_v73,_309_v80)).Merge((_313_v83).dtor_cf28));
        _314_v84 = _314_v84;
        if ((_module.__default.fm22((_302_v72).f23, (_this).f25, !((_this).f25), (_302_v72).f23, globalState)).IsProperSubsetOf(_306_v76)) {
          let _315_v85;
          let _init7 = ((_316_v76) => function (_317_i9) {
            return _316_v76;
          })(_306_v76);
          let _nw64 = Array((_dafny.ONE).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw64.length); _i0_7++) {
            _nw64[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _315_v85 = _nw64;
          let _318_v86;
          _318_v86 = _dafny.Seq.of(_315_v85);
          _315_v85 = (_318_v86)[_module.__default.safeIndex(new BigNumber(-367), new BigNumber((_318_v86).length))];
          (globalState).f5 = (_302_v72).f23;
          let _319_v87;
          _319_v87 = _dafny.Map.Empty.slice().updateUnsafe(_298_v69.f31,new BigNumber((_307_v77).length));
          (globalState).f1 = (_308_v79).IsDisjointFrom(_dafny.Set.fromElements((((_319_v87).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_319_v87,(_302_v72).f23)).length))) ? ((_319_v87).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_319_v87,(_302_v72).f23)).length))) : ((_302_v72).f23)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("pdjrt")).length),(_this).f25)).length), _298_v69.f31));
          let _320_v88;
          let _nw65 = Array((new BigNumber(9)).toNumber()).fill(_module.D5.Default());
          _320_v88 = _nw65;
          let _index51 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_320_v88).length));
          (_320_v88)[_index51] = ((_this.f24) ? (_module.D5.create_DC15(_262_v36, _295_v66, _301_v71, ((_314_v84).dtor_cf28).update(_module.D3.create_DC9(_262_v36), _309_v80))) : (_314_v84));
          let _index52 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_320_v88).length));
          (_320_v88)[_index52] = _313_v83;
          _319_v87 = (_319_v87).update(new BigNumber(446), _module.__default.safeModuloInt(new BigNumber((_261_v35).length), (_302_v72).f23));
        } else {
          let _index53 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_309_v80).length));
          (_309_v80)[_index53] = _298_v69.f31;
          let _index54 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_309_v80).length));
          (_309_v80)[_index54] = (_dafny.ZERO).minus(new BigNumber(((_298_v69).f30).length));
          let _321_v89;
          let _nw66 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _321_v89 = _nw66;
          let _322_v90;
          _322_v90 = _dafny.Seq.of(_321_v89);
          let _323_v91;
          let _nw67 = Array((new BigNumber(9)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _321_v89;
          _nw67[(_dafny.ONE).toNumber()] = _321_v89;
          _nw67[(new BigNumber(2)).toNumber()] = _321_v89;
          _nw67[(new BigNumber(3)).toNumber()] = _321_v89;
          _nw67[(new BigNumber(4)).toNumber()] = _321_v89;
          _nw67[(new BigNumber(5)).toNumber()] = _321_v89;
          _nw67[(new BigNumber(6)).toNumber()] = (_322_v90)[_module.__default.safeIndex((_307_v77)[_module.__default.safeIndex((_302_v72).f23, new BigNumber((_307_v77).length))], new BigNumber((_322_v90).length))];
          _nw67[(new BigNumber(7)).toNumber()] = _321_v89;
          _nw67[(new BigNumber(8)).toNumber()] = _321_v89;
          _323_v91 = _nw67;
          let _index55 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_323_v91).length));
          (_323_v91)[_index55] = _321_v89;
          let _index56 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_323_v91).length));
          (_323_v91)[_index56] = _321_v89;
          r3 = (_dafny.ZERO).minus(_298_v69.f31);
          let _324_v92;
          let _nw68 = new _module.C0();
          _nw68.__ctor(_297_v68, (_302_v72).f23, (_302_v72).f23);
          _324_v92 = _nw68;
          (globalState).f6 = _this.f24;
        }
        let _325_v93;
        let _init8 = function (_326_i10) {
          return _dafny.Seq.UnicodeFromString("uqbkxbwhd");
        };
        let _nw69 = Array((new BigNumber(24)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw69.length); _i0_8++) {
          _nw69[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _325_v93 = _nw69;
        let _index57 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_325_v93).length));
        (_325_v93)[_index57] = _210_v0;
        let _index58 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_325_v93).length));
        (_325_v93)[_index58] = _dafny.Seq.UnicodeFromString("easmd");
      }
      r2 = (_this).f25;
      let _327_v94;
      _327_v94 = new _dafny.CodePoint('x'.codePointAt(0));
      _327_v94 = _327_v94;
      r0 = _this.f24;
      let _328_v95;
      let _nw70 = Array((new BigNumber(18)).toNumber()).fill(false);
      _328_v95 = _nw70;
      let _329_v96;
      _329_v96 = _dafny.Set.fromElements(_328_v95, _328_v95, _328_v95);
      let _330_v97;
      let _nw71 = new _module.C0();
      _nw71.__ctor((_dafny.Set.fromElements(_328_v95, _328_v95)).Union(_329_v96), _262_v36, new BigNumber((_263_v37).length));
      _330_v97 = _nw71;
      r1 = _330_v97;
      r2 = !(_this.f24);
      let _331_v98;
      _331_v98 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-174),_module.__default.safeDivisionInt(_262_v36, new BigNumber(-599)));
      let _332_v99;
      _332_v99 = _dafny.Map.Empty.slice().updateUnsafe(_261_v35,_dafny.Set.fromElements((_this).f25, (_this).f25));
      r3 = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_331_v98).contains(_262_v36)) ? ((_331_v98).get(_262_v36)) : (_module.__default.safeModuloInt(new BigNumber((_332_v99).length), _262_v36)))));
      return [r0, r1, r2, r3];
    }
    m8(p0, p1, p2, p3, globalState) {
      let _this = this;
      (globalState).f8 = p1;
      (globalState).f6 = _module.__default.fm29(new BigNumber(856), p1, (_this).f25, globalState);
      let _333_v0;
      let _nw72 = Array((new BigNumber(6)).toNumber());
      _nw72[(_dafny.ZERO).toNumber()] = p3;
      _nw72[(_dafny.ONE).toNumber()] = _this.f24;
      _nw72[(new BigNumber(2)).toNumber()] = p3;
      _nw72[(new BigNumber(3)).toNumber()] = _this.f24;
      _nw72[(new BigNumber(4)).toNumber()] = (_this).f25;
      _nw72[(new BigNumber(5)).toNumber()] = (_this).f25;
      _333_v0 = _nw72;
      let _334_v1;
      _334_v1 = _dafny.Set.fromElements(_333_v0, _333_v0);
      let _335_v2;
      let _nw73 = new _module.C0();
      _nw73.__ctor(_334_v1, (_dafny.ZERO).minus((_dafny.ZERO).minus(p1)), new BigNumber(372));
      _335_v2 = _nw73;
      let _336_v3;
      _336_v3 = _dafny.Set.fromElements(p0);
      let _337_i0;
      _337_i0 = _dafny.ZERO;
      L2: {
        while ((_336_v3).IsProperSubsetOf((_dafny.Set.fromElements((_this).f25)).Union(_336_v3))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_337_i0)) {
              break L2;
            }
            _337_i0 = (_337_i0).plus(_dafny.ONE);
            let _338_v4;
            _338_v4 = _dafny.Seq.UnicodeFromString("helwwm");
            let _339_v5;
            _339_v5 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-671)), function (_340_i1) {
              return new _dafny.CodePoint('n'.codePointAt(0));
            }));
            let _341_v6;
            _341_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,p1);
            let _342_v7;
            _342_v7 = _dafny.Seq.of(new BigNumber((_341_v6).length));
            let _index59 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_333_v0).length));
            (_333_v0)[_index59] = _dafny.areEqual(_338_v4, (_339_v5)[_module.__default.safeIndex(new BigNumber((_342_v7).length), new BigNumber((_339_v5).length))]);
            let _index60 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_333_v0).length));
            (_333_v0)[_index60] = true;
            (globalState).f9 = !(!(_this.f24));
            (globalState).f15 = (_333_v0)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((_333_v0).length))];
            let _343_v8;
            let _init9 = ((_344_p0) => function (_345_i2) {
              return !(_344_p0);
            })(p0);
            let _nw74 = Array((new BigNumber(16)).toNumber());
            for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw74.length); _i0_9++) {
              _nw74[_i0_9] = _init9(new BigNumber(_i0_9));
            }
            _343_v8 = _nw74;
            let _346_v9;
            _346_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pfq"),_343_v8);
            _346_v9 = ((_346_v9).Merge(_346_v9)).Merge(_346_v9);
          }
        }
      }
      (globalState).f9 = p3;
      let _347_v10;
      _347_v10 = _dafny.Map.Empty.slice().updateUnsafe(p3,_335_v2.f31);
      let _source3 = _module.D0.create_DC1(!((_347_v10).equals(_347_v10)), p1);
      if (_source3.is_DC1) {
        let _348___mcc_h0 = (_source3).cf1;
        let _349___mcc_h1 = (_source3).cf2;
        let _350_cf2 = _349___mcc_h1;
        let _351_cf1 = _348___mcc_h0;
        (globalState).f1 = p3;
        let _352_v11;
        _352_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_335_v2.f31, _335_v2.f31, new BigNumber(483), _350_cf2),false);
        _352_v11 = _352_v11;
        let _353_v12;
        _353_v12 = _dafny.Seq.UnicodeFromString("ovboiepi");
        (globalState).f5 = (_dafny.ZERO).minus(((_this).fm5(_350_cf2, p1, globalState)).multipliedBy(_module.__default.fm1(new BigNumber((_353_v12).length), (_this).f25, globalState)));
        (_335_v2).f31 = (_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("gqn")).length), _335_v2.f31)).plus(_335_v2.f31);
      } else if (_source3.is_DC0) {
        let _354___mcc_h2 = (_source3).cf0;
        let _355_cf0 = _354___mcc_h2;
        let _356_v13;
        _356_v13 = _module.D4.create_DC12(new BigNumber((_dafny.Seq.of((_this).f25)).length), p1);
        let _357_v14;
        let _nw75 = Array((new BigNumber(3)).toNumber());
        _nw75[(_dafny.ZERO).toNumber()] = new BigNumber(985);
        _nw75[(_dafny.ONE).toNumber()] = (_this).fm5((_356_v13).dtor_cf19, (((_347_v10).contains(_this.f24)) ? ((_347_v10).get(_this.f24)) : (_335_v2.f31)), globalState);
        _nw75[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.of(p3, !((_this).f25))).length);
        _357_v14 = _nw75;
        let _358_v15;
        _358_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(304),_357_v14);
        let _359_v16;
        _359_v16 = _dafny.Seq.of(_357_v14, _357_v14, _357_v14, _357_v14, (((_358_v15).contains(p1)) ? ((_358_v15).get(p1)) : (_357_v14)));
        let _index61 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_333_v0).length));
        (_333_v0)[_index61] = _module.__default.fm29(_335_v2.f31, new BigNumber((_347_v10).length), !(_355_cf0), globalState);
        let _360_v17;
        _360_v17 = _dafny.Set.fromElements(_335_v2.f31, _335_v2.f31, (_dafny.ZERO).minus(p1));
        let _361_v18;
        _361_v18 = _dafny.Map.Empty.slice().updateUnsafe(_335_v2.f31,new BigNumber((_360_v17).length));
        let _362_v19;
        _362_v19 = _dafny.Seq.of((_this).f25);
        let _index62 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_333_v0).length));
        let _rhs74 = _dafny.Seq.Concat(_359_v16, _359_v16);
        let _rhs75 = _this.f24;
        let _rhs76 = _module.__default.safeDivisionInt(p1, _335_v2.f31);
        let _rhs77 = (((_this.f24) ? (p0) : (false))) === (((((_361_v18).contains(new BigNumber((_362_v19).length))) ? ((_361_v18).get(new BigNumber((_362_v19).length))) : (new BigNumber(-467)))).isLessThanOrEqualTo(new BigNumber(335)));
        let _rhs78 = _335_v2.f31;
        let _lhs62 = globalState;
        let _lhs63 = globalState;
        let _lhs64 = _333_v0;
        let _lhs65 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_333_v0).length));
        let _lhs66 = globalState;
        _359_v16 = _rhs74;
        _lhs62.f6 = _rhs75;
        _lhs63.f14 = _rhs76;
        _lhs64[_lhs65] = _rhs77;
        _lhs66.f8 = _rhs78;
        let _363_v20;
        let _nw76 = Array((new BigNumber(11)).toNumber()).fill(false);
        _363_v20 = _nw76;
        let _364_v21;
        _364_v21 = new _dafny.CodePoint('d'.codePointAt(0));
        let _365_v22;
        let _nw77 = new _module.C0();
        _nw77.__ctor(_dafny.Set.fromElements(_363_v20), new BigNumber((_module.__default.fm25((_362_v19)[_module.__default.safeIndex(new BigNumber((_361_v18).length), new BigNumber((_362_v19).length))], _this.f24, _364_v21, globalState)).length), p1);
        _365_v22 = _nw77;
        (globalState).f6 = (_362_v19)[_module.__default.safeIndex((p1).minus(_365_v22.f31), new BigNumber((_362_v19).length))];
        let _366_v24;
        let _init10 = ((_367_v21, _368_v22) => function (_369_i3) {
          return (_module.D9.create_DC25(_dafny.Seq.of(function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of (_dafny.Map.Empty.slice().updateUnsafe(_367_v21,_367_v21)).Keys.Elements) {
    let _370_v23 = _compr_14;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_367_v21,_367_v21)).contains(_370_v23)) {
      _coll14.push([_370_v23,_368_v22.f31]);
    }
  }
  return _coll14;
}()))).dtor_cf50;
        })(_364_v21, _365_v22);
        let _nw78 = Array((new BigNumber(5)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw78.length); _i0_10++) {
          _nw78[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _366_v24 = _nw78;
        let _371_v25;
        _371_v25 = _dafny.Map.Empty.slice().updateUnsafe(_364_v21,(_dafny.ZERO).minus(_365_v22.f31));
        let _372_v27;
        _372_v27 = _dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)));
        let _373_v28;
        _373_v28 = _dafny.Seq.of(_371_v25, _371_v25, _371_v25, function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of (_372_v27).Elements) {
            let _374_v26 = _compr_15;
            if (_dafny.Seq.contains(_372_v27, _374_v26)) {
              _coll15.push([_374_v26,_335_v2.f31]);
            }
          }
          return _coll15;
        }());
        let _index63 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_366_v24).length));
        (_366_v24)[_index63] = _373_v28;
        let _index64 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_366_v24).length));
        (_366_v24)[_index64] = _373_v28;
      } else {
        let _375___mcc_h3 = (_source3).cf3;
        let _376_cf3 = _375___mcc_h3;
        let _377_v29;
        _377_v29 = _dafny.Seq.of(_335_v2.f31);
        let _378_v30;
        let _nw79 = new _module.C0();
        _nw79.__ctor((_335_v2).f30, _335_v2.f31, new BigNumber((_377_v29).length));
        _378_v30 = _nw79;
        let _379_v31;
        _379_v31 = _dafny.MultiSet.fromElements(p0, p0, p3);
        let _380_v32;
        _380_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_379_v31);
        let _381_v33;
        _381_v33 = _dafny.Seq.UnicodeFromString("vb");
        _380_v32 = (_380_v32).update((_335_v2).fm3(_381_v33, (_378_v30).fm3(_381_v33, (_this).f25, p1, _dafny.Seq.of(new BigNumber((_381_v33).length)), globalState), _378_v30.f31, _377_v29, globalState), _379_v31);
        let _382_v34;
        _382_v34 = _dafny.MultiSet.fromElements(_336_v3);
        let _383_v35;
        _383_v35 = _dafny.MultiSet.fromElements(_378_v30.f31);
        let _384_v36;
        _384_v36 = _dafny.MultiSet.fromElements(new BigNumber((_382_v34).cardinality()), (((_383_v35).contains(new BigNumber((_377_v29).length))) ? ((_383_v35).get(new BigNumber((_377_v29).length))) : (_335_v2.f31)));
        let _385_v37;
        _385_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_381_v33).length),p0);
        let _386_v38;
        let _nw80 = Array((new BigNumber(10)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = _335_v2.f31;
        _nw80[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_335_v2.f31);
        _nw80[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(_335_v2.f31, _335_v2.f31);
        _nw80[(new BigNumber(3)).toNumber()] = (((_383_v35).contains(_378_v30.f31)) ? ((_383_v35).get(_378_v30.f31)) : ((_dafny.ZERO).minus(_378_v30.f31)));
        _nw80[(new BigNumber(4)).toNumber()] = (_378_v30.f31).plus(_335_v2.f31);
        _nw80[(new BigNumber(5)).toNumber()] = (_335_v2.f31).plus((_dafny.ZERO).minus(_335_v2.f31));
        _nw80[(new BigNumber(6)).toNumber()] = (p1).multipliedBy(p1);
        _nw80[(new BigNumber(7)).toNumber()] = (_335_v2).fm2(globalState);
        _nw80[(new BigNumber(8)).toNumber()] = new BigNumber((_336_v3).length);
        _nw80[(new BigNumber(9)).toNumber()] = new BigNumber(((_385_v37).Merge(_385_v37)).length);
        _386_v38 = _nw80;
        _386_v38 = _386_v38;
        _381_v33 = _381_v33;
      }
      return;
    }
    m9(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _387_v0;
      _387_v0 = _dafny.Seq.UnicodeFromString("nyt");
      let _388_v1;
      _388_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,p3);
      let _389_v2;
      _389_v2 = _module.D3.create_DC9(new BigNumber((_388_v1).length));
      let _hi2 = (_389_v2).dtor_cf15;
      for (let _390_i0 = new BigNumber((_dafny.Seq.Concat(_387_v0, _387_v0)).length); _390_i0.isLessThan(_hi2); _390_i0 = _390_i0.plus(_dafny.ONE)) {
        let _391_v3;
        _391_v3 = _dafny.Seq.of((_this).f25);
        (globalState).f9 = _module.__default.fm29(new BigNumber(444), _390_i0, (_391_v3)[_module.__default.safeIndex(_390_i0, new BigNumber((_391_v3).length))], globalState);
        let _392_v4;
        _392_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _393_v5;
        _393_v5 = _dafny.Map.Empty.slice().updateUnsafe(_392_v4,(_this).f25);
        (_this).m8(_this.f24, p3, (_393_v5).update(_392_v4, _this.f24), _module.__default.fm29(_390_i0, (_dafny.ZERO).minus(_390_i0), p2, globalState), globalState);
        let _394_v6;
        _394_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_this.f24);
        let _395_v7;
        _395_v7 = _dafny.Map.Empty.slice().updateUnsafe(_394_v6,p3);
        (_this).f24 = _module.__default.fm29((_dafny.ZERO).minus(_390_i0), (((_395_v7).contains(_394_v6)) ? ((_395_v7).get(_394_v6)) : (_390_i0)), (_this).f25, globalState);
        let _396_v8;
        let _nw81 = Array((new BigNumber(10)).toNumber()).fill(false);
        _396_v8 = _nw81;
        let _397_v9;
        _397_v9 = _dafny.Set.fromElements(_396_v8);
        let _398_v10;
        let _nw82 = new _module.C0();
        _nw82.__ctor(((p1) ? (_397_v9) : (_397_v9)), _390_i0, (new BigNumber((_387_v0).length)).minus(new BigNumber((_387_v0).length)));
        _398_v10 = _nw82;
      }
      let _399_v11;
      _399_v11 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      let _400_v12;
      let _nw83 = Array((new BigNumber(10)).toNumber());
      _nw83[(_dafny.ZERO).toNumber()] = p3;
      _nw83[(_dafny.ONE).toNumber()] = p3;
      _nw83[(new BigNumber(2)).toNumber()] = new BigNumber((_387_v0).length);
      _nw83[(new BigNumber(3)).toNumber()] = new BigNumber(-13);
      _nw83[(new BigNumber(4)).toNumber()] = p3;
      _nw83[(new BigNumber(5)).toNumber()] = p3;
      _nw83[(new BigNumber(6)).toNumber()] = p3;
      _nw83[(new BigNumber(7)).toNumber()] = p3;
      _nw83[(new BigNumber(8)).toNumber()] = new BigNumber((_387_v0).length);
      _nw83[(new BigNumber(9)).toNumber()] = (new BigNumber((_399_v11).length)).multipliedBy(p3);
      _400_v12 = _nw83;
      let _index65 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_400_v12).length));
      (_400_v12)[_index65] = (_dafny.ZERO).minus(new BigNumber((_387_v0).length));
      let _401_v13;
      _401_v13 = _dafny.Seq.of(_this.f24, p2);
      let _index66 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_400_v12).length));
      (_400_v12)[_index66] = (p3).minus(new BigNumber((_401_v13).length));
      (globalState).f14 = (_dafny.ZERO).minus(p3);
      let _index67 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_400_v12).length));
      (_400_v12)[_index67] = (_400_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_400_v12).length))];
      let _402_v14;
      _402_v14 = _dafny.Seq.of(p3);
      let _403_v15;
      _403_v15 = _dafny.MultiSet.fromElements(p3, (_400_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_400_v12).length))], (_400_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_400_v12).length))]);
      let _404_v16;
      let _nw84 = Array((new BigNumber(23)).toNumber());
      _nw84[(_dafny.ZERO).toNumber()] = p2;
      _nw84[(_dafny.ONE).toNumber()] = false;
      _nw84[(new BigNumber(2)).toNumber()] = p1;
      _nw84[(new BigNumber(3)).toNumber()] = _this.f24;
      _nw84[(new BigNumber(4)).toNumber()] = (_this).f25;
      _nw84[(new BigNumber(5)).toNumber()] = !(p1);
      _nw84[(new BigNumber(6)).toNumber()] = (_this).f25;
      _nw84[(new BigNumber(7)).toNumber()] = p2;
      _nw84[(new BigNumber(8)).toNumber()] = _this.f24;
      _nw84[(new BigNumber(9)).toNumber()] = (_this).f25;
      _nw84[(new BigNumber(10)).toNumber()] = (_this).f25;
      _nw84[(new BigNumber(11)).toNumber()] = p1;
      _nw84[(new BigNumber(12)).toNumber()] = p1;
      _nw84[(new BigNumber(13)).toNumber()] = p2;
      _nw84[(new BigNumber(14)).toNumber()] = (_this).f25;
      _nw84[(new BigNumber(15)).toNumber()] = (_this).f25;
      _nw84[(new BigNumber(16)).toNumber()] = _this.f24;
      _nw84[(new BigNumber(17)).toNumber()] = (_this).f25;
      _nw84[(new BigNumber(18)).toNumber()] = p2;
      _nw84[(new BigNumber(19)).toNumber()] = p2;
      _nw84[(new BigNumber(20)).toNumber()] = _this.f24;
      _nw84[(new BigNumber(21)).toNumber()] = p1;
      _nw84[(new BigNumber(22)).toNumber()] = (_this).f25;
      _404_v16 = _nw84;
      let _405_v17;
      _405_v17 = _dafny.Seq.of(_404_v16, _404_v16);
      let _406_v18;
      _406_v18 = _dafny.MultiSet.fromElements(p2);
      let _407_v19;
      _407_v19 = _dafny.Seq.of(new BigNumber((_402_v14).length), (((_403_v15).contains(new BigNumber((_405_v17).length))) ? ((_403_v15).get(new BigNumber((_405_v17).length))) : (new BigNumber((_406_v18).cardinality()))));
      let _408_v20;
      _408_v20 = _dafny.Seq.of((_402_v14)[_module.__default.safeIndex(p3, new BigNumber((_402_v14).length))], (_400_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_400_v12).length))]);
      let _rhs79 = (_dafny.ZERO).minus(p3);
      let _rhs80 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_400_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_400_v12).length))]), p3);
      let _rhs81 = p1;
      let _rhs82 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(812)), ((_409_p0) => function (_410_i1) {
        return _409_p0;
      })(p0));
      let _rhs83 = _dafny.Seq.IsPrefixOf(((!(_module.__default.fm29(new BigNumber(724), p3, (_this).f25, globalState))) ? (_408_v20) : (_dafny.Seq.update((_module.D10.create_DC28(_dafny.Seq.Create(_module.__default.abs(new BigNumber(820)), function (_411_i2) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("a")).length);
}))).dtor_cf52, _module.__default.safeIndex(p3, new BigNumber(((_module.D10.create_DC28(_dafny.Seq.Create(_module.__default.abs(new BigNumber(820)), function (_412_i2) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("a")).length);
}))).dtor_cf52).length)), p3))), _402_v14);
      let _lhs67 = globalState;
      let _lhs68 = globalState;
      let _lhs69 = _this;
      let _lhs70 = globalState;
      _lhs67.f14 = _rhs79;
      _lhs68.f8 = _rhs80;
      _lhs69.f24 = _rhs81;
      _387_v0 = _rhs82;
      _lhs70.f15 = _rhs83;
      (globalState).f15 = _module.__default.fm29(p3, (p3).multipliedBy((_400_v12)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_400_v12).length))]), (_this).f25, globalState);
      return;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f23 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    __ctor(f23) {
      let _this = this;
      (_this)._f23 = f23;
      return;
    }
    fm2(globalState) {
      let _this = this;
      return (_this).f23;
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(!(false), false), _dafny.Seq.of(!(false)))) || ((_dafny.MultiSet.fromElements((_this).f23)).IsProperSubsetOf(_dafny.MultiSet.fromElements((_this).f23, (_this).f23)));
    };
    fm19(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true, true), _dafny.Seq.of(false, true), _dafny.Seq.of(true))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(false, false, true), _dafny.Seq.of(false)));
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = false;
      r0 = (_this).f23;
      if ((p0) || (p0)) {
        let _413_v0;
        let _init11 = ((_414_p1) => function (_415_i0) {
          return ((_414_p1) ? (_414_p1) : (!(_414_p1)));
        })(p1);
        let _nw85 = Array((new BigNumber(5)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw85.length); _i0_11++) {
          _nw85[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _413_v0 = _nw85;
        let _index68 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_413_v0).length));
        (_413_v0)[_index68] = p0;
        let _416_v1;
        _416_v1 = _dafny.Seq.UnicodeFromString("qbhsiypk");
        let _index69 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_413_v0).length));
        let _rhs84 = _module.__default.safeDivisionInt((_this).f23, new BigNumber(982));
        let _rhs85 = p0;
        let _rhs86 = false;
        let _rhs87 = _416_v1;
        let _lhs71 = globalState;
        let _lhs72 = _413_v0;
        let _lhs73 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_413_v0).length));
        let _lhs74 = globalState;
        _lhs71.f8 = _rhs84;
        _lhs72[_lhs73] = _rhs85;
        _lhs74.f1 = _rhs86;
        _416_v1 = _rhs87;
        let _417_v2;
        _417_v2 = _module.D3.create_DC8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-517)), function (_418_i1) {
  return (_module.D1.create_DC3(new _dafny.CodePoint('m'.codePointAt(0)))).dtor_cf4;
}));
        let _419_v3;
        _419_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_413_v0)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_413_v0).length))]);
        let _420_v4;
        _420_v4 = _dafny.Set.fromElements(p0, p0);
        let _rhs88 = (new BigNumber(((_419_v3).Merge(_419_v3)).length)).minus((_this).f23);
        let _rhs89 = _417_v2;
        let _rhs90 = ((_this).f23).minus((_this).f23);
        let _rhs91 = (p1) && (p0);
        let _rhs92 = _module.__default.safeDivisionInt(new BigNumber(((_420_v4).Difference(_dafny.Set.fromElements(false))).length), (_this).f23);
        let _lhs75 = globalState;
        let _lhs76 = globalState;
        let _lhs77 = globalState;
        let _lhs78 = globalState;
        _lhs75.f17 = _rhs88;
        _417_v2 = _rhs89;
        _lhs76.f14 = _rhs90;
        _lhs77.f6 = _rhs91;
        _lhs78.f17 = _rhs92;
        (globalState).f14 = (_this).f23;
        let _421_v5;
        _421_v5 = _dafny.Map.Empty.slice().updateUnsafe(_416_v1,(new BigNumber((_dafny.Seq.UnicodeFromString("oqfaosxj")).length)).isLessThan((_this).f23));
        let _422_v6;
        _422_v6 = _dafny.Set.fromElements((_this).f23);
        let _423_v7;
        _423_v7 = _dafny.Seq.of((_413_v0)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_413_v0).length))], (_413_v0)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_413_v0).length))]);
        let _rhs93 = _421_v5;
        let _rhs94 = _module.__default.safeDivisionInt((_this).f23, ((_this).f23).plus((_this).f23));
        let _rhs95 = (((_422_v6).IsProperSubsetOf(_422_v6)) ? ((_this).f23) : (new BigNumber((_423_v7).length)));
        let _rhs96 = (_this).f23;
        let _lhs79 = globalState;
        let _lhs80 = globalState;
        let _lhs81 = globalState;
        _421_v5 = _rhs93;
        _lhs79.f17 = _rhs94;
        _lhs80.f14 = _rhs95;
        _lhs81.f17 = _rhs96;
        let _424_v8;
        let _nw86 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
        _424_v8 = _nw86;
        let _425_v9;
        _425_v9 = _dafny.Map.Empty.slice().updateUnsafe(_413_v0,p1);
        let _index70 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_424_v8).length));
        (_424_v8)[_index70] = _425_v9;
        let _index71 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_424_v8).length));
        (_424_v8)[_index71] = (_module.D4.create_DC11(_425_v9)).dtor_cf18;
      } else {
        let _426_v10;
        let _nw87 = Array((new BigNumber(23)).toNumber());
        _nw87[(_dafny.ZERO).toNumber()] = !(p1);
        _nw87[(_dafny.ONE).toNumber()] = p1;
        _nw87[(new BigNumber(2)).toNumber()] = p0;
        _nw87[(new BigNumber(3)).toNumber()] = p1;
        _nw87[(new BigNumber(4)).toNumber()] = p1;
        _nw87[(new BigNumber(5)).toNumber()] = false;
        _nw87[(new BigNumber(6)).toNumber()] = p0;
        _nw87[(new BigNumber(7)).toNumber()] = !(p1);
        _nw87[(new BigNumber(8)).toNumber()] = p0;
        _nw87[(new BigNumber(9)).toNumber()] = p0;
        _nw87[(new BigNumber(10)).toNumber()] = !(false);
        _nw87[(new BigNumber(11)).toNumber()] = p0;
        _nw87[(new BigNumber(12)).toNumber()] = p0;
        _nw87[(new BigNumber(13)).toNumber()] = p0;
        _nw87[(new BigNumber(14)).toNumber()] = p1;
        _nw87[(new BigNumber(15)).toNumber()] = p1;
        _nw87[(new BigNumber(16)).toNumber()] = p1;
        _nw87[(new BigNumber(17)).toNumber()] = p1;
        _nw87[(new BigNumber(18)).toNumber()] = p1;
        _nw87[(new BigNumber(19)).toNumber()] = p0;
        _nw87[(new BigNumber(20)).toNumber()] = p1;
        _nw87[(new BigNumber(21)).toNumber()] = p0;
        _nw87[(new BigNumber(22)).toNumber()] = p0;
        _426_v10 = _nw87;
        let _427_v11;
        _427_v11 = _dafny.Seq.UnicodeFromString("wt");
        let _428_v12;
        _428_v12 = _dafny.Map.Empty.slice().updateUnsafe(p0,_427_v11);
        let _429_v13;
        _429_v13 = _dafny.MultiSet.fromElements(_427_v11, (((_428_v12).contains(p1)) ? ((_428_v12).get(p1)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-505)), function (_430_i2) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        }))), _427_v11);
        let _431_v14;
        _431_v14 = _dafny.Seq.of(_429_v13, _429_v13, _429_v13);
        let _432_v15;
        _432_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(341),new BigNumber(((_431_v14)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(861)), function (_433_i3) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).length), new BigNumber((_431_v14).length))]).cardinality()));
        let _434_v16;
        _434_v16 = _dafny.Map.Empty.slice().updateUnsafe(_426_v10,_432_v15);
        _434_v16 = (_434_v16).Merge(_434_v16);
        (globalState).f1 = false;
        let _435_v17;
        let _nw88 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _435_v17 = _nw88;
        let _index72 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_435_v17).length));
        (_435_v17)[_index72] = ((_this).f23).minus((_this).f23);
        let _index73 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_435_v17).length));
        (_435_v17)[_index73] = _module.__default.safeModuloInt((_this).f23, ((_this).f23).multipliedBy((_this).f23));
        (globalState).f15 = p1;
        let _436_v18;
        _436_v18 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
        let _437_v19;
        _437_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(431));
        let _438_v20;
        _438_v20 = _module.D1.create_DC5(p0, _436_v18, (_this).f23, (((_432_v15).contains((_this).f23)) ? ((_432_v15).get((_this).f23)) : ((_435_v17)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_435_v17).length))])), _437_v19);
        let _439_v21;
        _439_v21 = _dafny.Map.Empty.slice().updateUnsafe((_435_v17)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_435_v17).length))],_438_v20);
        _439_v21 = (_439_v21).update((_435_v17)[_module.__default.safeIndex(new BigNumber(406), new BigNumber((_435_v17).length))], _438_v20);
      }
      let _440_v22;
      _440_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(!(p0)),p0);
      let _441_v23;
      _441_v23 = _module.D3.create_DC10(!(p1), (_440_v22).Merge(_440_v22));
      let _442_v24;
      _442_v24 = new _dafny.CodePoint('g'.codePointAt(0));
      let _443_v26;
      _443_v26 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of _dafny.IntegerRange(new BigNumber(240), new BigNumber(994))) {
          let _444_v25 = _compr_16;
          if (((new BigNumber(240)).isLessThanOrEqualTo(_444_v25)) && ((_444_v25).isLessThan(new BigNumber(994)))) {
            _coll16.push([(_444_v25).minus((_this).f23),(_this).f23]);
          }
        }
        return _coll16;
      }(),(_this).f23);
      let _445_v27;
      _445_v27 = _module.D5.create_DC13(_443_v26);
      let _446_v28;
      _446_v28 = _dafny.Seq.of(false);
      let _447_v29;
      _447_v29 = _dafny.Seq.UnicodeFromString("ngg");
      let _rhs97 = new BigNumber(((_445_v27).dtor_cf21).length);
      let _rhs98 = ((_446_v28)[_module.__default.safeIndex((_this).f23, new BigNumber((_446_v28).length))]) && (!(_dafny.Seq.IsProperPrefixOf(_447_v29, _447_v29)));
      let _rhs99 = _441_v23;
      let _rhs100 = _442_v24;
      let _lhs82 = globalState;
      let _lhs83 = globalState;
      _lhs82.f8 = _rhs97;
      _lhs83.f15 = _rhs98;
      _441_v23 = _rhs99;
      _442_v24 = _rhs100;
      _447_v29 = _dafny.Seq.Concat(_447_v29, _module.__default.fm20(p0, globalState));
      let _448_v30;
      _448_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f23);
      let _449_v31;
      _449_v31 = _module.D1.create_DC5(true, _440_v22, (_this).f23, (_this).f23, _448_v30);
      let _450_v32;
      _450_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((((_449_v31).dtor_cf8).Merge(_440_v22)).length));
      let _451_v33;
      _451_v33 = _dafny.Seq.of((_this).f23, (_this).f23, (_this).f23);
      _450_v32 = (_450_v32).update((_this).fm3(_447_v29, p1, (_this).f23, _451_v33, globalState), ((_this).f23).plus(_module.__default.fm1((_this).f23, p1, globalState)));
      let _452_v34;
      _452_v34 = _dafny.MultiSet.fromElements(p0);
      let _453_v35;
      _453_v35 = _dafny.Seq.of(_451_v33);
      let _454_v36;
      _454_v36 = _dafny.MultiSet.fromElements(new BigNumber(-143));
      let _455_v37;
      _455_v37 = _dafny.Set.fromElements(new BigNumber((_454_v36).cardinality()));
      let _rhs101 = (_this).fm3(((p1) ? (_447_v29) : (_447_v29)), true, (_this).f23, (_453_v35)[_module.__default.safeIndex((_this).f23, new BigNumber((_453_v35).length))], globalState);
      let _rhs102 = (_455_v37).IsSubsetOf(_455_v37);
      let _rhs103 = (_dafny.MultiSet.fromElements(p0, p1)).Union((_dafny.MultiSet.fromElements(p1, p0)).update(!(p0), _module.__default.abs(new BigNumber((_dafny.Seq.of(p0)).length))));
      let _lhs84 = globalState;
      r2 = _rhs101;
      _lhs84.f6 = _rhs102;
      _452_v34 = _rhs103;
      r0 = _module.__default.safeDivisionInt(new BigNumber(30), (_this).f23);
      let _456_v38;
      let _nw89 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _456_v38 = _nw89;
      r1 = _456_v38;
      r2 = (new BigNumber(384)).isLessThan((_this).f23);
      return [r0, r1, r2];
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      (globalState).f9 = p2;
      let _457_v0;
      _457_v0 = _dafny.Seq.UnicodeFromString("wqyjmxlhd");
      let _hi3 = (_this).fm2(globalState);
      for (let _458_i0 = ((_dafny.ZERO).minus(new BigNumber((_457_v0).length))).plus(new BigNumber((_dafny.MultiSet.fromElements(_module.D3.create_DC9(p3))).cardinality())); _458_i0.isLessThan(_hi3); _458_i0 = _458_i0.plus(_dafny.ONE)) {
        let _459_v1;
        _459_v1 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_458_i0,p3)).length));
        let _460_v2;
        _460_v2 = _module.D1.create_DC5(false, _dafny.Map.Empty.slice().updateUnsafe(p0,p2), _458_i0, p3, _459_v1);
        let _461_v3;
        _461_v3 = _dafny.Map.Empty.slice().updateUnsafe((_460_v2).dtor_cf7,p1);
        let _462_v4;
        _462_v4 = _dafny.Map.Empty.slice().updateUnsafe(_460_v2,_458_i0);
        let _463_v5;
        _463_v5 = _dafny.Seq.of(_462_v4, _462_v4, _462_v4);
        let _464_v6;
        _464_v6 = _dafny.Seq.of(_463_v5, _463_v5, _463_v5);
        let _465_v7;
        _465_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_461_v3).length),_dafny.Seq.Concat((_464_v6)[_module.__default.safeIndex(_458_i0, new BigNumber((_464_v6).length))], (_464_v6)[_module.__default.safeIndex(_458_i0, new BigNumber((_464_v6).length))]));
        _465_v7 = (_465_v7).update((_this).f23, _463_v5);
        _459_v1 = (_459_v1).update(!(p2), _458_i0);
        let _466_v8;
        _466_v8 = _module.D3.create_DC8(_457_v0);
        let _467_v9;
        _467_v9 = new _dafny.CodePoint('i'.codePointAt(0));
        let _468_v10;
        _468_v10 = _dafny.Seq.of(p2, p0);
        let _469_v11;
        _469_v11 = _dafny.Set.fromElements(p0);
        let _470_v12;
        _470_v12 = _dafny.MultiSet.fromElements(_458_i0);
        let _471_v13;
        let _nw90 = Array((new BigNumber(29)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = p2;
        _nw90[(_dafny.ONE).toNumber()] = (_this).fm3((_466_v8).dtor_cf14, p0, _458_i0, p1, globalState);
        _nw90[(new BigNumber(2)).toNumber()] = p2;
        _nw90[(new BigNumber(3)).toNumber()] = p2;
        _nw90[(new BigNumber(4)).toNumber()] = _dafny.Seq.IsPrefixOf(_457_v0, _dafny.Seq.update(_457_v0, _module.__default.safeIndex(_458_i0, new BigNumber((_457_v0).length)), _467_v9));
        _nw90[(new BigNumber(5)).toNumber()] = p0;
        _nw90[(new BigNumber(6)).toNumber()] = p0;
        _nw90[(new BigNumber(7)).toNumber()] = p0;
        _nw90[(new BigNumber(8)).toNumber()] = p0;
        _nw90[(new BigNumber(9)).toNumber()] = !(p0);
        _nw90[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_468_v10, _dafny.Seq.of((_468_v10)[_module.__default.safeIndex(_module.__default.fm1((_this).f23, !(p0), globalState), new BigNumber((_468_v10).length))], p2));
        _nw90[(new BigNumber(11)).toNumber()] = p0;
        _nw90[(new BigNumber(12)).toNumber()] = !(p2);
        _nw90[(new BigNumber(13)).toNumber()] = !((p3).isLessThan((_this).f23));
        _nw90[(new BigNumber(14)).toNumber()] = (_469_v11).IsSubsetOf(_469_v11);
        _nw90[(new BigNumber(15)).toNumber()] = ((p2) ? (!((_468_v10)[_module.__default.safeIndex(_458_i0, new BigNumber((_468_v10).length))])) : (!(!(p0))));
        _nw90[(new BigNumber(16)).toNumber()] = p2;
        _nw90[(new BigNumber(17)).toNumber()] = p0;
        _nw90[(new BigNumber(18)).toNumber()] = p2;
        _nw90[(new BigNumber(19)).toNumber()] = p2;
        _nw90[(new BigNumber(20)).toNumber()] = (_this).fm3(_457_v0, p0, _458_i0, p1, globalState);
        _nw90[(new BigNumber(21)).toNumber()] = p0;
        _nw90[(new BigNumber(22)).toNumber()] = true;
        _nw90[(new BigNumber(23)).toNumber()] = p2;
        _nw90[(new BigNumber(24)).toNumber()] = p0;
        _nw90[(new BigNumber(25)).toNumber()] = p0;
        _nw90[(new BigNumber(26)).toNumber()] = p0;
        _nw90[(new BigNumber(27)).toNumber()] = (_470_v12).IsDisjointFrom(_470_v12);
        _nw90[(new BigNumber(28)).toNumber()] = p0;
        _471_v13 = _nw90;
        let _index74 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_471_v13).length));
        (_471_v13)[_index74] = true;
        let _index75 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_471_v13).length));
        (_471_v13)[_index75] = p2;
        let _472_v14;
        _472_v14 = _dafny.Set.fromElements(new BigNumber(-582), _458_i0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(366)), ((_473_i0) => function (_474_i1) {
          return _473_i0;
        })(_458_i0))).length), p3);
        let _475_v15;
        _475_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(_dafny.Set.fromElements((_this).f23, p3)).equals(_472_v14));
        if ((((_475_v15).contains((new BigNumber((_457_v0).length)).isLessThanOrEqualTo(new BigNumber((_457_v0).length)))) ? ((_475_v15).get((new BigNumber((_457_v0).length)).isLessThanOrEqualTo(new BigNumber((_457_v0).length)))) : (p0))) {
          let _476_v16;
          _476_v16 = _dafny.Map.Empty.slice().updateUnsafe(_467_v9,(_this).f23);
          let _477_v17;
          _477_v17 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_467_v9,_module.__default.fm1((_this).f23, true, globalState)), _476_v16);
          _477_v17 = _dafny.Seq.Concat(_477_v17, _477_v17);
          let _478_v18;
          let _nw91 = Array((new BigNumber(14)).toNumber()).fill(_module.D0.Default());
          _478_v18 = _nw91;
          let _479_v19;
          _479_v19 = _dafny.Map.Empty.slice().updateUnsafe(_478_v18,_467_v9);
          let _480_v20;
          _480_v20 = _module.D3.create_DC9(_458_i0);
          let _481_v21;
          let _nw92 = Array((new BigNumber(29)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = new BigNumber(-441);
          _nw92[(_dafny.ONE).toNumber()] = (_this).fm2(globalState);
          _nw92[(new BigNumber(2)).toNumber()] = p3;
          _nw92[(new BigNumber(3)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(4)).toNumber()] = p3;
          _nw92[(new BigNumber(5)).toNumber()] = _458_i0;
          _nw92[(new BigNumber(6)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(7)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(8)).toNumber()] = new BigNumber((p1).length);
          _nw92[(new BigNumber(9)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(10)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(11)).toNumber()] = new BigNumber(-676);
          _nw92[(new BigNumber(12)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(13)).toNumber()] = _458_i0;
          _nw92[(new BigNumber(14)).toNumber()] = _458_i0;
          _nw92[(new BigNumber(15)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(16)).toNumber()] = p3;
          _nw92[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_458_i0);
          _nw92[(new BigNumber(18)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(19)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(20)).toNumber()] = new BigNumber((_469_v11).length);
          _nw92[(new BigNumber(21)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(22)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(23)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(24)).toNumber()] = (_this).f23;
          _nw92[(new BigNumber(25)).toNumber()] = p3;
          _nw92[(new BigNumber(26)).toNumber()] = p3;
          _nw92[(new BigNumber(27)).toNumber()] = p3;
          _nw92[(new BigNumber(28)).toNumber()] = (_this).f23;
          _481_v21 = _nw92;
          let _482_v22;
          _482_v22 = _dafny.Map.Empty.slice().updateUnsafe(_480_v20,_481_v21);
          let _483_v23;
          _483_v23 = _module.D5.create_DC15(new BigNumber(292), (((_471_v13)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_471_v13).length))]) ? ((_457_v0)[_module.__default.safeIndex((_this).f23, new BigNumber((_457_v0).length))]) : (_467_v9)), _479_v19, (_482_v22).Merge(_482_v22));
          let _484_v24;
          let _init12 = ((_485_p0) => function (_486_i2) {
            return _dafny.Seq.of(_485_p0);
          })(p0);
          let _nw93 = Array((new BigNumber(21)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw93.length); _i0_12++) {
            _nw93[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _484_v24 = _nw93;
          let _487_v25;
          _487_v25 = _module.D6.create_DC16(_468_v10);
          let _index76 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_484_v24).length));
          (_484_v24)[_index76] = (_487_v25).dtor_cf29;
          let _488_v26;
          _488_v26 = _dafny.Map.Empty.slice().updateUnsafe(((p2) ? (new BigNumber(-978)) : ((p1)[_module.__default.safeIndex(new BigNumber(-663), new BigNumber((p1).length))])),_467_v9);
          let _pat_let_tv1 = _479_v19;
          let _pat_let_tv2 = _479_v19;
          let _pat_let_tv3 = p3;
          let _index77 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_484_v24).length));
          let _rhs104 = function (_pat_let2_0) {
            return function (_489_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_490_dt__update_hcf27_h0) {
                  return function (_pat_let4_0) {
                    return function (_491_dt__update_hcf25_h0) {
                      return _module.D5.create_DC15(_491_dt__update_hcf25_h0, (_489_dt__update__tmp_h0).dtor_cf26, _490_dt__update_hcf27_h0, (_489_dt__update__tmp_h0).dtor_cf28);
                    }(_pat_let4_0);
                  }((_pat_let_tv3).minus((_this).f23));
                }(_pat_let3_0);
              }((_pat_let_tv1).Merge(_pat_let_tv2));
            }(_pat_let2_0);
          }(_483_v23);
          let _rhs105 = _471_v13;
          let _rhs106 = _471_v13;
          let _rhs107 = _468_v10;
          let _rhs108 = (_488_v26).Merge(_488_v26);
          let _lhs85 = _484_v24;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_484_v24).length));
          _483_v23 = _rhs104;
          _471_v13 = _rhs105;
          _471_v13 = _rhs106;
          _lhs85[_lhs86] = _rhs107;
          _488_v26 = _rhs108;
          (globalState).f5 = _module.__default.fm1(p3, true, globalState);
          (globalState).f1 = (_469_v11).IsSubsetOf(_469_v11);
          (globalState).f9 = !((_471_v13)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_471_v13).length))]);
        } else {
          let _492_v27;
          _492_v27 = _module.D6.create_DC16(_468_v10);
          _492_v27 = _492_v27;
          let _index78 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_471_v13).length));
          (_471_v13)[_index78] = p2;
          let _493_v28;
          _493_v28 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.MultiSet.fromElements((_468_v10)[_module.__default.safeIndex((_dafny.ZERO).minus(p3), new BigNumber((_468_v10).length))]));
          _493_v28 = (_493_v28).update((_471_v13)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_471_v13).length))], _module.__default.fm21(p2, true, (_this).f23, globalState));
          let _494_v29;
          _494_v29 = _dafny.Map.Empty.slice().updateUnsafe((_458_i0).minus((_this).f23),(((_471_v13)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_471_v13).length))]) ? (p3) : (_458_i0)));
          _494_v29 = (_494_v29).update(p3, (_this).f23);
          (globalState).f6 = (!(p0)) === (false);
        }
      }
      let _495_v30;
      _495_v30 = _dafny.Map.Empty.slice().updateUnsafe(p2,(p3).minus(p3));
      (globalState).f0 = _495_v30;
      let _496_v31;
      let _nw94 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Set.Empty);
      _496_v31 = _nw94;
      let _497_v32;
      _497_v32 = _dafny.Set.fromElements(!(p0));
      let _index79 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_496_v31).length));
      (_496_v31)[_index79] = _497_v32;
      let _index80 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_496_v31).length));
      (_496_v31)[_index80] = _497_v32;
      _495_v30 = (_495_v30).Merge(_495_v30);
      (globalState).f17 = (_this).f23;
      let _498_v33;
      _498_v33 = _module.D0.create_DC1(p2, p3);
      r0 = _498_v33;
      r1 = p3;
      let _499_v34;
      _499_v34 = _module.D0.create_DC0(p0);
      let _pat_let_tv4 = _457_v0;
      r2 = function (_source4) {
        if (_source4.is_DC1) {
          let _500___mcc_h0 = (_source4).cf1;
          let _501___mcc_h1 = (_source4).cf2;
          let _502_cf2 = _501___mcc_h1;
          let _503_cf1 = _500___mcc_h0;
          return new BigNumber((_pat_let_tv4).length);
        } else if (_source4.is_DC0) {
          let _504___mcc_h2 = (_source4).cf0;
          let _505_cf0 = _504___mcc_h2;
          return new BigNumber((_dafny.Seq.of(new BigNumber(-577))).length);
        } else {
          let _506___mcc_h3 = (_source4).cf3;
          let _507_cf3 = _506___mcc_h3;
          return (_this).f23;
        }
      }(_module.D0.create_DC2(_499_v34));
      return [r0, r1, r2];
    }
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      (globalState).f14 = (_this).f23;
      let _index81 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length));
      (p0)[_index81] = p2;
      let _508_v0;
      _508_v0 = _dafny.Seq.UnicodeFromString("qiu");
      let _index82 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length));
      (p0)[_index82] = ((true) ? (!_dafny.areEqual(_dafny.Seq.UnicodeFromString("mseqsel"), _508_v0)) : (p1));
      (globalState).f17 = (_this).f23;
      let _509_v1;
      _509_v1 = _dafny.MultiSet.fromElements(p1, p1, p1, p2);
      let _510_v2;
      _510_v2 = _dafny.MultiSet.fromElements((_this).f23);
      let _511_v3;
      _511_v3 = _dafny.Seq.of(p2, p2, p1);
      let _512_v4;
      _512_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_510_v2).cardinality()),p2);
      let _513_v5;
      _513_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_512_v4).length),_508_v0);
      let _514_v6;
      _514_v6 = _dafny.Map.Empty.slice().updateUnsafe((((_513_v5).contains((_this).f23)) ? ((_513_v5).get((_this).f23)) : (_508_v0)),_511_v3);
      let _515_v7;
      _515_v7 = new _dafny.CodePoint('j'.codePointAt(0));
      let _516_v8;
      _516_v8 = _dafny.Set.fromElements(_515_v7);
      let _517_v9;
      let _init13 = ((_518_p1) => function (_519_i1) {
        return _module.D0.create_DC2(_module.D0.create_DC1(_518_p1, (_this).f23));
      })(p1);
      let _nw95 = Array((new BigNumber(20)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw95.length); _i0_13++) {
        _nw95[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _517_v9 = _nw95;
      let _520_v10;
      _520_v10 = _dafny.Map.Empty.slice().updateUnsafe(_517_v9,_515_v7);
      let _521_v11;
      _521_v11 = _module.D3.create_DC9((_this).f23);
      let _522_v12;
      let _init14 = function (_523_i2) {
        return _module.__default.safeDivisionInt(_523_i2, (_this).f23);
      };
      let _nw96 = Array((new BigNumber(9)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw96.length); _i0_14++) {
        _nw96[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _522_v12 = _nw96;
      let _524_v13;
      _524_v13 = _module.D5.create_DC15(new BigNumber((_dafny.Seq.update(_508_v0, _module.__default.safeIndex(new BigNumber((_516_v8).length), new BigNumber((_508_v0).length)), _515_v7)).length), _515_v7, _520_v10, _dafny.Map.Empty.slice().updateUnsafe(_521_v11,_522_v12));
      let _525_v14;
      let _nw97 = Array((new BigNumber(15)).toNumber());
      _nw97[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(223)), function (_526_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length);
      _nw97[(_dafny.ONE).toNumber()] = (_this).f23;
      _nw97[(new BigNumber(2)).toNumber()] = (_this).f23;
      _nw97[(new BigNumber(3)).toNumber()] = (((_509_v1).contains(p2)) ? ((_509_v1).get(p2)) : ((_this).f23));
      _nw97[(new BigNumber(4)).toNumber()] = new BigNumber((_510_v2).cardinality());
      _nw97[(new BigNumber(5)).toNumber()] = new BigNumber((_module.__default.fm21(true, p2, (_this).f23, globalState)).cardinality());
      _nw97[(new BigNumber(6)).toNumber()] = (_this).f23;
      _nw97[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm20((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))], globalState)).length));
      _nw97[(new BigNumber(8)).toNumber()] = (new BigNumber((_508_v0).length)).multipliedBy((_this).f23);
      _nw97[(new BigNumber(9)).toNumber()] = (_this).f23;
      _nw97[(new BigNumber(10)).toNumber()] = (_this).f23;
      _nw97[(new BigNumber(11)).toNumber()] = (_this).fm2(globalState);
      _nw97[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_511_v3, (((_514_v6).contains(_508_v0)) ? ((_514_v6).get(_508_v0)) : (_511_v3)))).length);
      _nw97[(new BigNumber(13)).toNumber()] = (_524_v13).dtor_cf25;
      _nw97[(new BigNumber(14)).toNumber()] = (_this).f23;
      _525_v14 = _nw97;
      let _index83 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length));
      (_522_v12)[_index83] = _module.__default.safeDivisionInt((_this).f23, (_this).f23);
      let _527_v15;
      _527_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2(globalState),p0);
      let _528_v16;
      _528_v16 = _dafny.Set.fromElements((((_527_v15).contains((_this).f23)) ? ((_527_v15).get((_this).f23)) : (p0)));
      let _index84 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length));
      (_522_v12)[_index84] = (_dafny.ZERO).minus((new BigNumber((((p3) ? (_dafny.Set.fromElements(p0)) : (_528_v16))).length)).minus((_this).f23));
      let _529_i3;
      _529_i3 = _dafny.ZERO;
      L3: {
        while (!_dafny.Seq.contains(_508_v0, _515_v7)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_529_i3)) {
              break L3;
            }
            _529_i3 = (_529_i3).plus(_dafny.ONE);
            if (!(p2) || (p2)) {
              let _530_v17;
              let _nw98 = new _module.C1();
              _nw98.__ctor(p3, false);
              _530_v17 = _nw98;
              let _531_v18;
              let _nw99 = new _module.C0();
              _nw99.__ctor((_528_v16).Difference(_528_v16), (_522_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length))], new BigNumber(262));
              _531_v18 = _nw99;
              let _532_v19;
              _532_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,p3);
              (globalState).f1 = (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm29(_531_v18.f31, (_this).f23, (p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))], globalState),p3)).Merge(_532_v19)).length)).isLessThanOrEqualTo((_522_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length))]);
              let _533_v20;
              let _nw100 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _533_v20 = _nw100;
              let _index85 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_533_v20).length));
              (_533_v20)[_index85] = _dafny.Seq.Concat(_508_v0, _508_v0);
              let _index86 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_533_v20).length));
              (_533_v20)[_index86] = _dafny.Seq.Concat(_508_v0, _module.__default.fm20((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))], globalState));
              let _534_v21;
              _534_v21 = _dafny.Seq.of(_531_v18.f31, new BigNumber(715));
              let _535_v22;
              _535_v22 = _dafny.Set.fromElements((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))]);
              let _536_v23;
              _536_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_534_v21, _module.__default.safeIndex((_this).f23, new BigNumber((_534_v21).length)), (_522_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length))]),(_535_v22).IsDisjointFrom(_535_v22));
              let _537_v25;
              _537_v25 = _dafny.Seq.of(_534_v21, _534_v21);
              _536_v23 = ((function () {
                let _coll17 = new _dafny.Map();
                for (const _compr_17 of (_537_v25).Elements) {
                  let _538_v24 = _compr_17;
                  if (_dafny.Seq.contains(_537_v25, _538_v24)) {
                    _coll17.push([_538_v24,(p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))]]);
                  }
                }
                return _coll17;
              }()).Merge(function () {
                let _coll18 = new _dafny.Map();
                for (const _compr_18 of (_536_v23).Keys.Elements) {
                  let _539_v26 = _compr_18;
                  if ((_536_v23).contains(_539_v26)) {
                    _coll18.push([_539_v26,p1]);
                  }
                }
                return _coll18;
              }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(_534_v21,p1));
            } else {
              let _540_v27;
              _540_v27 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt((_522_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length))], new BigNumber(-120))),new BigNumber((_dafny.MultiSet.fromElements((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))], p3, p1, true, p3)).cardinality()));
              _540_v27 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_510_v2).contains((_this).f23)) ? ((_510_v2).get((_this).f23)) : (new BigNumber(-265))),new _dafny.CodePoint('p'.codePointAt(0)))).length),_module.__default.fm1(_module.__default.fm1(new BigNumber((_dafny.MultiSet.FromArray(_511_v3)).cardinality()), p2, globalState), p2, globalState))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23));
              r0 = !(!_dafny.Seq.contains(_dafny.Seq.Concat(_508_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(511)), function (_541_i4) {
                return new _dafny.CodePoint('m'.codePointAt(0));
              })), _515_v7));
              (globalState).f17 = _module.__default.safeDivisionInt(((_this).f23).plus((_522_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length))]), (_this).f23);
              (globalState).f17 = new BigNumber(368);
              let _542_v30;
              let _init15 = ((_543_v12, _544_v7, _545_v4) => function (_546_i5) {
                return function () {
                  let _coll19 = new _dafny.Map();
                  for (const _compr_19 of (function () {
                    let _coll20 = new _dafny.Set();
                    for (const _compr_20 of (_545_v4).Keys.Elements) {
                      let _547_v29 = _compr_20;
                      if ((_545_v4).contains(_547_v29)) {
                        _coll20.add(_module.__default.safeModuloInt(_547_v29, new BigNumber((_dafny.Seq.UnicodeFromString("qnqjr")).length)));
                      }
                    }
                    return _coll20;
                  }()).Elements) {
                    let _548_v28 = _compr_19;
                    if ((function () {
                      let _coll21 = new _dafny.Set();
                      for (const _compr_21 of (_545_v4).Keys.Elements) {
                        let _549_v29 = _compr_21;
                        if ((_545_v4).contains(_549_v29)) {
                          _coll21.add(_module.__default.safeModuloInt(_549_v29, new BigNumber((_dafny.Seq.UnicodeFromString("qnqjr")).length)));
                        }
                      }
                      return _coll21;
                    }()).contains(_548_v28)) {
                      _coll19.push([_module.__default.safeDivisionInt(_548_v28, (_543_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_543_v12).length))]),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_544_v7,_544_v7)).length)]);
                    }
                  }
                  return _coll19;
                }();
              })(_522_v12, _515_v7, _512_v4);
              let _nw101 = Array((new BigNumber(18)).toNumber());
              for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw101.length); _i0_15++) {
                _nw101[_i0_15] = _init15(new BigNumber(_i0_15));
              }
              _542_v30 = _nw101;
              let _init16 = ((_550_v27) => function (_551_i6) {
                return _550_v27;
              })(_540_v27);
              let _nw102 = Array((new BigNumber(24)).toNumber());
              for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw102.length); _i0_16++) {
                _nw102[_i0_16] = _init16(new BigNumber(_i0_16));
              }
              _542_v30 = _nw102;
            }
            let _552_v31;
            _552_v31 = _dafny.Map.Empty.slice().updateUnsafe(p3,(p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))]);
            let _553_v32;
            _553_v32 = _dafny.Set.fromElements((_this).f23);
            let _554_v33;
            _554_v33 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_553_v32).length));
            let _555_v34;
            _555_v34 = _module.D1.create_DC5((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))], _552_v31, new BigNumber(-897), new BigNumber((_dafny.Seq.update(_module.__default.fm30((((_509_v1).contains(p2)) ? ((_509_v1).get(p2)) : (new BigNumber(-629))), _508_v0, p3, (((_509_v1).contains(true)) ? ((_509_v1).get(true)) : ((((_554_v33).contains((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))])) ? ((_554_v33).get((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))])) : (new BigNumber((_508_v0).length))))), globalState), _module.__default.safeIndex((_522_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length))], new BigNumber((_module.__default.fm30((((_509_v1).contains(p2)) ? ((_509_v1).get(p2)) : (new BigNumber(-629))), _508_v0, p3, (((_509_v1).contains(true)) ? ((_509_v1).get(true)) : ((((_554_v33).contains((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))])) ? ((_554_v33).get((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))])) : (new BigNumber((_508_v0).length))))), globalState)).length)), (p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))])).length), (_554_v33).Merge(_554_v33));
            _555_v34 = _555_v34;
            (globalState).f8 = (_522_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length))];
            (globalState).f8 = (_this).f23;
          }
        }
      }
      let _556_v35;
      _556_v35 = _dafny.Seq.of(_510_v2);
      let _557_v36;
      _557_v36 = _dafny.Seq.of(_556_v35, _556_v35, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-187)), ((_558_v2) => function (_559_i7) {
        return _558_v2;
      })(_510_v2)));
      (globalState).f9 = _dafny.Seq.IsProperPrefixOf(_556_v35, (_557_v36)[_module.__default.safeIndex((_522_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length))], new BigNumber((_557_v36).length))]);
      let _560_v37;
      let _nw103 = new _module.C0();
      _nw103.__ctor(_528_v16, (_this).f23, (_dafny.ZERO).minus((_522_v12)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_522_v12).length))]));
      _560_v37 = _nw103;
      let _561_v38;
      _561_v38 = _dafny.Seq.of(_560_v37, _560_v37);
      r0 = ((_dafny.Map.Empty.slice().updateUnsafe(p3,_561_v38)).update((p0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((p0).length))], _561_v38)).contains(p3);
      return r0;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f24 = false;
      this._f23 = _dafny.ZERO;
      this._f25 = false;
      this._f28 = false;
      this._f29 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    set f24(value) {
      let _this = this;
      _this._f24 = value;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f28, f29, f24, f25, f23) {
      let _this = this;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      (_this)._f23 = f23;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f28)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28)));
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((_this).f29)).multipliedBy((_this).f29);
    };
    fm2(globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).length)).plus((_this).f29);
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((((_this).f28) ? (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("njpng"))) : (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("arqurnxwi"))))).IsProperSubsetOf((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("pfvvyx"), _dafny.Seq.UnicodeFromString("htve"))).Difference(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("exnwakoad"), _dafny.Seq.UnicodeFromString("jfhvlkukr"))));
    };
    fm16(p0, p1, p2, globalState) {
      let _this = this;
      return !(!(!(!((_this).f29).isEqualTo((_this).f23))) || ((new BigNumber((_dafny.Seq.of((_this).f25)).length)).isLessThanOrEqualTo((_this).f23)));
    };
    fm17(p0, globalState) {
      let _this = this;
      return (_this).f25;
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = undefined;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _562_v0;
      let _nw104 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
      _562_v0 = _nw104;
      let _563_v1;
      _563_v1 = new _dafny.CodePoint('d'.codePointAt(0));
      let _564_v2;
      _564_v2 = _module.D1.create_DC3(_563_v1);
      let _565_v3;
      _565_v3 = _module.D1.create_DC6(_564_v2);
      let _566_v4;
      let _nw105 = Array((new BigNumber(6)).toNumber());
      _nw105[(_dafny.ZERO).toNumber()] = (_this).f25;
      _nw105[(_dafny.ONE).toNumber()] = _this.f24;
      _nw105[(new BigNumber(2)).toNumber()] = (_this).f28;
      _nw105[(new BigNumber(3)).toNumber()] = (_this).fm17(_565_v3, globalState);
      _nw105[(new BigNumber(4)).toNumber()] = (_this).f25;
      _nw105[(new BigNumber(5)).toNumber()] = _this.f24;
      _566_v4 = _nw105;
      let _567_v5;
      _567_v5 = _dafny.Set.fromElements(_566_v4);
      let _index87 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_562_v0).length));
      (_562_v0)[_index87] = _567_v5;
      let _568_v6;
      _568_v6 = _dafny.Seq.of(_567_v5);
      let _569_v7;
      _569_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_567_v5);
      let _index88 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_562_v0).length));
      (_562_v0)[_index88] = ((_568_v6)[_module.__default.safeIndex((_this).f23, new BigNumber((_568_v6).length))]).Difference((((_569_v7).contains(new BigNumber(20))) ? ((_569_v7).get(new BigNumber(20))) : (_567_v5)));
      let _570_v8;
      _570_v8 = _dafny.MultiSet.fromElements((_this).f25, ((_this).f28) || (true));
      _570_v8 = _570_v8;
      (globalState).f15 = (_this).f28;
      let _571_v9;
      _571_v9 = _dafny.Seq.of((_this).f28);
      let _572_v10;
      _572_v10 = _dafny.Seq.of((_this).f23);
      let _573_v11;
      _573_v11 = _dafny.Seq.of(_572_v10);
      let _574_v12;
      _574_v12 = _dafny.MultiSet.fromElements(new BigNumber((_573_v11).length), (_this).f29);
      let _575_v13;
      _575_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(((_574_v12).contains(new BigNumber(197))) ? ((_574_v12).get(new BigNumber(197))) : (new BigNumber(410))));
      let _576_v14;
      _576_v14 = _module.D1.create_DC5(_this.f24, _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_this.f24), new BigNumber((_dafny.Seq.UnicodeFromString("lhhqehb")).length), new BigNumber((_571_v9).length), _575_v13);
      let _577_v15;
      _577_v15 = _dafny.Seq.of((_576_v14).dtor_cf10, (_this).f23);
      (globalState).f8 = (_572_v10)[_module.__default.safeIndex((_dafny.ZERO).minus(((_this).f29).multipliedBy((_this).f29)), new BigNumber((_572_v10).length))];
      let _578_v16;
      _578_v16 = _dafny.Seq.of(_576_v14);
      r3 = (((_this).f25) ? (((true) ? (new BigNumber((_578_v16).length)) : ((_this).f23))) : ((_this).f23));
      let _579_v17;
      let _nw106 = Array((new BigNumber(15)).toNumber()).fill([]);
      _579_v17 = _nw106;
      let _index89 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_579_v17).length));
      (_579_v17)[_index89] = _566_v4;
      let _index90 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_566_v4).length));
      (_566_v4)[_index90] = ((_this).f25) || ((_571_v9)[_module.__default.safeIndex((_this).f23, new BigNumber((_571_v9).length))]);
      let _580_v18;
      _580_v18 = _module.D3.create_DC9((_this).f23);
      let _581_v19;
      let _nw107 = Array((new BigNumber(9)).toNumber());
      _nw107[(_dafny.ZERO).toNumber()] = new BigNumber((_module.__default.fm18((_this).f23, (_this).f29, _571_v9, _580_v18, globalState)).length);
      _nw107[(_dafny.ONE).toNumber()] = ((_this).f23).minus((_this).f23);
      _nw107[(new BigNumber(2)).toNumber()] = (_this).f29;
      _nw107[(new BigNumber(3)).toNumber()] = (_this).f23;
      _nw107[(new BigNumber(4)).toNumber()] = (_module.__default.fm1((_this).f29, true, globalState)).minus((_this).f29);
      _nw107[(new BigNumber(5)).toNumber()] = (_this).f29;
      _nw107[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.of((_this).f28, (_this).f25)).length);
      _nw107[(new BigNumber(7)).toNumber()] = (_this).f29;
      _nw107[(new BigNumber(8)).toNumber()] = (_this).f29;
      _581_v19 = _nw107;
      let _index91 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_581_v19).length));
      (_581_v19)[_index91] = (_this).f23;
      let _582_v20;
      _582_v20 = _dafny.Set.fromElements(new BigNumber(181), (_this).f23);
      let _583_v21;
      _583_v21 = _dafny.Set.fromElements(_582_v20);
      let _index92 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_579_v17).length));
      let _index93 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_566_v4).length));
      let _index94 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_581_v19).length));
      let _rhs109 = _566_v4;
      let _rhs110 = (_this).fm16((_this).f25, (_this).f28, _this.f24, globalState);
      let _rhs111 = _module.__default.safeModuloInt(_module.__default.fm1((_dafny.ZERO).minus((_this).f23), _this.f24, globalState), ((_this).f29).minus(new BigNumber((_583_v21).length)));
      let _rhs112 = (_dafny.ZERO).minus((_this).f29);
      let _lhs87 = _579_v17;
      let _lhs88 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_579_v17).length));
      let _lhs89 = _566_v4;
      let _lhs90 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_566_v4).length));
      let _lhs91 = _581_v19;
      let _lhs92 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_581_v19).length));
      let _lhs93 = globalState;
      _lhs87[_lhs88] = _rhs109;
      _lhs89[_lhs90] = _rhs110;
      _lhs91[_lhs92] = _rhs111;
      _lhs93.f8 = _rhs112;
      let _584_v22;
      _584_v22 = _dafny.Map.Empty.slice().updateUnsafe(_563_v1,new _dafny.CodePoint('h'.codePointAt(0)));
      let _pat_let_tv5 = _571_v9;
      let _pat_let_tv6 = _581_v19;
      let _pat_let_tv7 = _581_v19;
      let _pat_let_tv8 = _571_v9;
      let _pat_let_tv9 = _566_v4;
      let _pat_let_tv10 = _566_v4;
      r0 = function (_source5) {
        if (_source5.is_DC4) {
          let _585___mcc_h0 = (_source5).cf5;
          let _586___mcc_h1 = (_source5).cf6;
          let _587_cf6 = _586___mcc_h1;
          let _588_cf5 = _585___mcc_h0;
          return _587_cf6;
        } else if (_source5.is_DC5) {
          let _589___mcc_h2 = (_source5).cf7;
          let _590___mcc_h3 = (_source5).cf8;
          let _591___mcc_h4 = (_source5).cf9;
          let _592___mcc_h5 = (_source5).cf10;
          let _593___mcc_h6 = (_source5).cf11;
          let _594_cf11 = _593___mcc_h6;
          let _595_cf10 = _592___mcc_h5;
          let _596_cf9 = _591___mcc_h4;
          let _597_cf8 = _590___mcc_h3;
          let _598_cf7 = _589___mcc_h2;
          return _598_cf7;
        } else if (_source5.is_DC3) {
          let _599___mcc_h7 = (_source5).cf4;
          let _600_cf4 = _599___mcc_h7;
          return (_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv5)[_module.__default.safeIndex((_pat_let_tv7)[_module.__default.safeIndex(new BigNumber(52), new BigNumber((_pat_let_tv6).length))], new BigNumber((_pat_let_tv8).length))],(_this).f28)).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_pat_let_tv10)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_pat_let_tv9).length))]));
        } else {
          let _601___mcc_h8 = (_source5).cf12;
          let _602_cf12 = _601___mcc_h8;
          return (_this).f28;
        }
      }(_module.D1.create_DC3((((_584_v22).contains(_563_v1)) ? ((_584_v22).get(_563_v1)) : (new _dafny.CodePoint('k'.codePointAt(0))))));
      let _603_v23;
      let _nw108 = new _module.C2();
      _nw108.__ctor((_dafny.ZERO).minus((_this).f29));
      _603_v23 = _nw108;
      r1 = _603_v23;
      r2 = !((_this).f25) || ((_this).f28);
      r3 = (_this).f23;
      return [r0, r1, r2, r3];
    }
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = false;
      let _604_v0;
      _604_v0 = _dafny.Set.fromElements((_this).f29);
      let _605_v1;
      _605_v1 = _module.D8.create_DC23((_this).f23, _604_v0, (_this).f23, (_dafny.ZERO).minus(_module.__default.fm1((_this).f29, p1, globalState)), new BigNumber(850));
      let _606_v2;
      _606_v2 = _dafny.Map.Empty.slice().updateUnsafe(_605_v1,_dafny.Seq.UnicodeFromString("h"));
      let _607_v3;
      _607_v3 = _dafny.Seq.UnicodeFromString("gnilurti");
      _606_v2 = (_606_v2).update(((p0) ? (_605_v1) : (_605_v1)), _607_v3);
      let _608_v4;
      _608_v4 = _dafny.Seq.of(_module.__default.fm31(p0, (_this).f28, globalState));
      let _609_v5;
      _609_v5 = _dafny.Seq.of(new BigNumber(692));
      let _610_i0;
      _610_i0 = _dafny.ZERO;
      L4: {
        while (_dafny.areEqual(_dafny.Seq.update((_608_v4)[_module.__default.safeIndex((_this).f29, new BigNumber((_608_v4).length))], _module.__default.safeIndex((_dafny.ZERO).minus((_this).f29), new BigNumber(((_608_v4)[_module.__default.safeIndex((_this).f29, new BigNumber((_608_v4).length))]).length)), (_dafny.ZERO).minus((_this).f29)), _609_v5)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_610_i0)) {
              break L4;
            }
            _610_i0 = (_610_i0).plus(_dafny.ONE);
            let _611_v6;
            let _nw109 = Array((new BigNumber(24)).toNumber()).fill(false);
            _611_v6 = _nw109;
            let _612_v7;
            _612_v7 = _dafny.Set.fromElements(_611_v6, _611_v6, _611_v6, _611_v6);
            let _613_v8;
            let _nw110 = new _module.C0();
            _nw110.__ctor(_612_v7, (_this).f23, _module.__default.safeDivisionInt((_this).f23, (_dafny.ZERO).minus((_this).f23)));
            _613_v8 = _nw110;
            let _614_v9;
            let _nw111 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
            _614_v9 = _nw111;
            let _615_v10;
            _615_v10 = _dafny.Set.fromElements(_614_v9);
            let _616_v11;
            _616_v11 = _dafny.Set.fromElements((_this).f28);
            let _617_v12;
            _617_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_this).f23, (_this).f29, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(845)), function (_618_i1) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            })).length), new BigNumber((_615_v10).length), new BigNumber((_616_v11).length))).length),(_this).f23);
            let _619_v13;
            let _nw112 = new _module.C1();
            _nw112.__ctor(false, (_617_v12).equals(_617_v12));
            _619_v13 = _nw112;
            let _620_v14;
            _620_v14 = _dafny.Seq.of(_this.f24, p1);
            let _index95 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_611_v6).length));
            (_611_v6)[_index95] = _dafny.Seq.IsPrefixOf(_620_v14, _620_v14);
            let _index96 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_611_v6).length));
            (_611_v6)[_index96] = true;
            let _621_v15;
            _621_v15 = _dafny.Map.Empty.slice().updateUnsafe(_607_v3,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(848)), function (_622_i2) {
              return new _dafny.CodePoint('h'.codePointAt(0));
            })).length));
            (globalState).f6 = !(!(_621_v15).contains(_607_v3));
          }
        }
      }
      let _623_v16;
      _623_v16 = _dafny.MultiSet.fromElements((_this).f23);
      let _624_v17;
      _624_v17 = _dafny.MultiSet.fromElements(!(p1));
      let _625_v18;
      _625_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_624_v17).cardinality()),_623_v16);
      let _626_v19;
      _626_v19 = _module.D4.create_DC12(new BigNumber(((_623_v16).Difference((((_625_v18).contains((_this).f29)) ? ((_625_v18).get((_this).f29)) : (_623_v16)))).cardinality()), (((_this).f28) ? ((_this).f23) : ((_this).f23)));
      _626_v19 = _module.D4.create_DC12((_this).f29, (_dafny.ZERO).minus((_this).f23));
      let _hi4 = new BigNumber(-402);
      for (let _627_i3 = (_this).f29; _627_i3.isLessThan(_hi4); _627_i3 = _627_i3.plus(_dafny.ONE)) {
        let _628_v20;
        _628_v20 = _dafny.Set.fromElements(p0);
        let _629_v21;
        _629_v21 = _dafny.Seq.of(p1);
        let _630_v22;
        _630_v22 = _dafny.Seq.of(_629_v21, _629_v21, _629_v21, _629_v21);
        let _631_v23;
        _631_v23 = _module.D1.create_DC4(p1, true);
        let _632_v24;
        _632_v24 = _dafny.Seq.of(_631_v23, _631_v23, _631_v23, _631_v23, _631_v23);
        let _633_v26;
        let _init17 = ((_634_p1) => function (_635_i4) {
          return _634_p1;
        })(p1);
        let _nw113 = Array((new BigNumber(15)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw113.length); _i0_17++) {
          _nw113[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _633_v26 = _nw113;
        let _636_v27;
        _636_v27 = _dafny.Map.Empty.slice().updateUnsafe(_633_v26,(_this).f28);
        let _637_v28;
        _637_v28 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_633_v26);
        let _638_v29;
        _638_v29 = _dafny.Set.fromElements(_631_v23);
        let _639_v31;
        _639_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-446),_628_v20);
        let _640_v32;
        _640_v32 = _dafny.Seq.of(_628_v20, _628_v20, _628_v20, _628_v20, _628_v20);
        let _641_v33;
        _641_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,p1);
        let _642_v34;
        _642_v34 = _module.D1.create_DC6(_module.D1.create_DC5((_this).f28, _641_v33, _627_i3, (_this).f23, _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(_627_i3))));
        let _643_v35;
        let _nw114 = Array((new BigNumber(22)).toNumber());
        _nw114[(_dafny.ZERO).toNumber()] = _628_v20;
        _nw114[(_dafny.ONE).toNumber()] = _628_v20;
        _nw114[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(true);
        _nw114[(new BigNumber(3)).toNumber()] = _628_v20;
        _nw114[(new BigNumber(4)).toNumber()] = _628_v20;
        _nw114[(new BigNumber(5)).toNumber()] = _module.__default.fm32(_this.f24, _609_v5, (_630_v22)[_module.__default.safeIndex((_this).f29, new BigNumber((_630_v22).length))], function () {
          let _coll22 = new _dafny.Set();
          for (const _compr_22 of (_632_v24).Elements) {
            let _644_v25 = _compr_22;
            if (_dafny.Seq.contains(_632_v24, _644_v25)) {
              _coll22.add(_644_v25);
            }
          }
          return _coll22;
        }(), globalState);
        _nw114[(new BigNumber(6)).toNumber()] = _628_v20;
        _nw114[(new BigNumber(7)).toNumber()] = _628_v20;
        _nw114[(new BigNumber(8)).toNumber()] = _module.__default.fm32((((_636_v27).contains((((_637_v28).contains(!(p0))) ? ((_637_v28).get(!(p0))) : (_633_v26)))) ? ((_636_v27).get((((_637_v28).contains(!(p0))) ? ((_637_v28).get(!(p0))) : (_633_v26)))) : (p1)), _609_v5, _629_v21, _638_v29, globalState);
        _nw114[(new BigNumber(9)).toNumber()] = _module.__default.fm32((_this).f28, _609_v5, _629_v21, function () {
          let _coll23 = new _dafny.Set();
          for (const _compr_23 of (_module.__default.fm33((_this).f28, _607_v3, new BigNumber(846), (_this).f28, globalState)).Elements) {
            let _645_v30 = _compr_23;
            if (_dafny.Seq.contains(_module.__default.fm33((_this).f28, _607_v3, new BigNumber(846), (_this).f28, globalState), _645_v30)) {
              _coll23.add(_645_v30);
            }
          }
          return _coll23;
        }(), globalState);
        _nw114[(new BigNumber(10)).toNumber()] = _628_v20;
        _nw114[(new BigNumber(11)).toNumber()] = (((_639_v31).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_629_v21)[_module.__default.safeIndex(new BigNumber((_607_v3).length), new BigNumber((_629_v21).length))],(_this).f28)).length))) ? ((_639_v31).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_629_v21)[_module.__default.safeIndex(new BigNumber((_607_v3).length), new BigNumber((_629_v21).length))],(_this).f28)).length))) : (_628_v20));
        _nw114[(new BigNumber(12)).toNumber()] = ((_640_v32)[_module.__default.safeIndex(new BigNumber((_641_v33).length), new BigNumber((_640_v32).length))]).Intersect(_628_v20);
        _nw114[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(_this.f24, _this.f24, (_629_v21)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f29), new BigNumber((_629_v21).length))]);
        _nw114[(new BigNumber(14)).toNumber()] = ((true) ? (_628_v20) : (_dafny.Set.fromElements(_this.f24)));
        _nw114[(new BigNumber(15)).toNumber()] = _dafny.Set.fromElements(!(p1));
        _nw114[(new BigNumber(16)).toNumber()] = _628_v20;
        _nw114[(new BigNumber(17)).toNumber()] = _628_v20;
        _nw114[(new BigNumber(18)).toNumber()] = _628_v20;
        _nw114[(new BigNumber(19)).toNumber()] = _dafny.Set.fromElements(!((_this).fm17(_642_v34, globalState)), (_this).f25, !(p0));
        _nw114[(new BigNumber(20)).toNumber()] = _628_v20;
        _nw114[(new BigNumber(21)).toNumber()] = _module.__default.fm32(_this.f24, _609_v5, _629_v21, _638_v29, globalState);
        _643_v35 = _nw114;
        let _index97 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_643_v35).length));
        (_643_v35)[_index97] = _628_v20;
        let _index98 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_633_v26).length));
        (_633_v26)[_index98] = _this.f24;
        let _index99 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_643_v35).length));
        let _index100 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_633_v26).length));
        let _rhs113 = (_628_v20).Union(_628_v20);
        let _rhs114 = (_module.__default.safeDivisionInt(new BigNumber((_607_v3).length), (_this).f23)).isLessThanOrEqualTo(new BigNumber(386));
        let _rhs115 = (_dafny.ZERO).minus((_this).f23);
        let _rhs116 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f23, (_this).f23)));
        let _lhs94 = _643_v35;
        let _lhs95 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_643_v35).length));
        let _lhs96 = _633_v26;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_633_v26).length));
        let _lhs98 = globalState;
        let _lhs99 = globalState;
        _lhs94[_lhs95] = _rhs113;
        _lhs96[_lhs97] = _rhs114;
        _lhs98.f17 = _rhs115;
        _lhs99.f14 = _rhs116;
        (globalState).f5 = (((_624_v17).contains(((_633_v26)[_module.__default.safeIndex(new BigNumber(820), new BigNumber((_633_v26).length))]) === (_this.f24))) ? ((_624_v17).get(((_633_v26)[_module.__default.safeIndex(new BigNumber(820), new BigNumber((_633_v26).length))]) === (_this.f24))) : (((_this).f29).minus(new BigNumber((_607_v3).length))));
        let _646_v36;
        _646_v36 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(553)), function (_647_i5) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }));
        let _648_v37;
        let _nw115 = new _module.C1();
        _nw115.__ctor((_this).f28, (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("riafjk"))).IsDisjointFrom(_646_v36));
        _648_v37 = _nw115;
        (globalState).f14 = (_this).f29;
      }
      if (p0) {
        if ((_this).fm16(true, (_this).f28, _this.f24, globalState)) {
          let _649_v38;
          let _nw116 = new _module.C2();
          _nw116.__ctor((_module.D10.create_DC29(_this.f24, _this.f24, (_this).f23)).dtor_cf55);
          _649_v38 = _nw116;
          (globalState).f15 = ((_this).f25) && (_this.f24);
          let _650_v39;
          _650_v39 = new _dafny.CodePoint('g'.codePointAt(0));
          _607_v3 = _dafny.Seq.Concat(_607_v3, _dafny.Seq.update((_module.D3.create_DC8(_dafny.Seq.update(_607_v3, _module.__default.safeIndex((_this).f29, new BigNumber((_607_v3).length)), _650_v39))).dtor_cf14, _module.__default.safeIndex((_this).f23, new BigNumber(((_module.D3.create_DC8(_dafny.Seq.update(_607_v3, _module.__default.safeIndex((_this).f29, new BigNumber((_607_v3).length)), _650_v39))).dtor_cf14).length)), _650_v39));
          let _651_v40;
          let _nw117 = new _module.C1();
          _nw117.__ctor((_this).f25, p1);
          _651_v40 = _nw117;
          let _652_v41;
          let _nw118 = Array((new BigNumber(5)).toNumber());
          _nw118[(_dafny.ZERO).toNumber()] = _651_v40;
          _nw118[(_dafny.ONE).toNumber()] = _651_v40;
          _nw118[(new BigNumber(2)).toNumber()] = _651_v40;
          _nw118[(new BigNumber(3)).toNumber()] = _651_v40;
          _nw118[(new BigNumber(4)).toNumber()] = _651_v40;
          _652_v41 = _nw118;
          let _653_v42;
          _653_v42 = _dafny.MultiSet.fromElements(_652_v41);
          let _654_v43;
          let _nw119 = Array((new BigNumber(28)).toNumber()).fill(false);
          _654_v43 = _nw119;
          let _655_v44;
          _655_v44 = _dafny.Map.Empty.slice().updateUnsafe(_654_v43,(((_624_v17).contains((_this).f28)) ? ((_624_v17).get((_this).f28)) : ((_this).f29)));
          let _rhs117 = ((_653_v42).update(_652_v41, _module.__default.abs((_this).f29))).IsDisjointFrom(_653_v42);
          let _rhs118 = _module.__default.safeDivisionInt(new BigNumber(-779), new BigNumber(((_655_v44).Merge(_655_v44)).length));
          let _lhs100 = globalState;
          let _lhs101 = globalState;
          _lhs100.f1 = _rhs117;
          _lhs101.f17 = _rhs118;
          let _656_v45;
          _656_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,true);
          let _657_v46;
          _657_v46 = _dafny.Map.Empty.slice().updateUnsafe(_624_v17,p0);
          let _658_v47;
          _658_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(65));
          let _659_v48;
          _659_v48 = _module.D1.create_DC5(p1, _656_v45, new BigNumber(-776), new BigNumber((_657_v46).length), _658_v47);
          let _660_v49;
          _660_v49 = _dafny.Seq.of((_this).f28, _this.f24);
          let _661_v50;
          let _nw120 = Array((new BigNumber(8)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = (_this).f23;
          _nw120[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(156)), function (_662_i6) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(858),true)).length);
          })).length);
          _nw120[(new BigNumber(2)).toNumber()] = (_this).f23;
          _nw120[(new BigNumber(3)).toNumber()] = (_659_v48).dtor_cf10;
          _nw120[(new BigNumber(4)).toNumber()] = (_this).f23;
          _nw120[(new BigNumber(5)).toNumber()] = (_this).f29;
          _nw120[(new BigNumber(6)).toNumber()] = new BigNumber((_660_v49).length);
          _nw120[(new BigNumber(7)).toNumber()] = (_this).f23;
          _661_v50 = _nw120;
          let _663_v51;
          let _nw121 = Array((new BigNumber(14)).toNumber());
          _nw121[(_dafny.ZERO).toNumber()] = _661_v50;
          _nw121[(_dafny.ONE).toNumber()] = _661_v50;
          _nw121[(new BigNumber(2)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(3)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(4)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(5)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(6)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(7)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(8)).toNumber()] = ((p1) ? (_661_v50) : (_661_v50));
          _nw121[(new BigNumber(9)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(10)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(11)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(12)).toNumber()] = _661_v50;
          _nw121[(new BigNumber(13)).toNumber()] = _661_v50;
          _663_v51 = _nw121;
          let _rhs119 = (_this).f29;
          let _rhs120 = (_this).f23;
          let _rhs121 = p0;
          let _rhs122 = new _dafny.CodePoint('t'.codePointAt(0));
          let _rhs123 = _663_v51;
          let _lhs102 = globalState;
          let _lhs103 = globalState;
          let _lhs104 = globalState;
          _lhs102.f8 = _rhs119;
          _lhs103.f8 = _rhs120;
          _lhs104.f9 = _rhs121;
          _650_v39 = _rhs122;
          _663_v51 = _rhs123;
        } else {
          let _664_v52;
          _664_v52 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,true);
          let _665_v53;
          _665_v53 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-899)), function (_666_i7) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })).length)).plus(new BigNumber((_664_v52).length)),(_this).f25);
          _665_v53 = (_665_v53).update((_this).f29, p1);
          (globalState).f5 = _module.__default.safeDivisionInt((_this).f29, (_this).f29);
          let _667_v54;
          _667_v54 = new _dafny.CodePoint('m'.codePointAt(0));
          _667_v54 = _667_v54;
          let _668_v55;
          _668_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f23);
          (globalState).f8 = (((_this).f29).multipliedBy(new BigNumber((_668_v55).length))).multipliedBy(((_this).f29).minus((_this).f29));
          (globalState).f5 = (_this).fm2(globalState);
        }
        let _669_v56;
        _669_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_604_v0).length),p0);
        _669_v56 = _669_v56;
        let _670_v57;
        let _nw122 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _670_v57 = _nw122;
        r1 = _670_v57;
        let _671_v58;
        _671_v58 = new _dafny.CodePoint('n'.codePointAt(0));
        _671_v58 = _671_v58;
        (globalState).f6 = false;
      } else {
        (globalState).f1 = (_624_v17).IsDisjointFrom(_624_v17);
        let _672_v59;
        let _init18 = function (_673_i8) {
          return (_this).f25;
        };
        let _nw123 = Array((new BigNumber(5)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw123.length); _i0_18++) {
          _nw123[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _672_v59 = _nw123;
        _672_v59 = _672_v59;
        let _rhs124 = _dafny.Seq.Concat(_607_v3, _607_v3);
        let _rhs125 = (_this).f29;
        let _lhs105 = globalState;
        _607_v3 = _rhs124;
        _lhs105.f5 = _rhs125;
        let _index101 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_672_v59).length));
        (_672_v59)[_index101] = (_this).f28;
        let _index102 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_672_v59).length));
        (_672_v59)[_index102] = _this.f24;
        let _674_v60;
        let _nw124 = new _module.C1();
        _nw124.__ctor(p0, (_this).f28);
        _674_v60 = _nw124;
        let _675_v61;
        _675_v61 = _dafny.Set.fromElements(_674_v60);
        (_this).f24 = (p0) === (!((_675_v61).equals(_675_v61)));
      }
      let _676_v62;
      let _init19 = function (_677_i9) {
        return (_677_i9).plus((_this).f29);
      };
      let _nw125 = Array((new BigNumber(27)).toNumber());
      for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw125.length); _i0_19++) {
        _nw125[_i0_19] = _init19(new BigNumber(_i0_19));
      }
      _676_v62 = _nw125;
      let _index103 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_676_v62).length));
      (_676_v62)[_index103] = (_this).f23;
      let _index104 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_676_v62).length));
      (_676_v62)[_index104] = (_this).f23;
      r0 = (_this).f29;
      r1 = _676_v62;
      r2 = !(_this.f24);
      return [r0, r1, r2];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let _678_v0;
      let _init20 = function (_679_i1) {
        return _this.f24;
      };
      let _nw126 = Array((new BigNumber(15)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw126.length); _i0_20++) {
        _nw126[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _678_v0 = _nw126;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_678_v0).length))) {
        let _680_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_680_i0)) && ((_680_i0).isLessThan(new BigNumber((_678_v0).length))))) {
          (_678_v0)[(_680_i0)] = !_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus((_this).f23)), (_this).f23);
        }
      }
      let _681_v1;
      let _init21 = function (_682_i2) {
        return (_682_i2).multipliedBy((_this).f29);
      };
      let _nw127 = Array((new BigNumber(24)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw127.length); _i0_21++) {
        _nw127[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _681_v1 = _nw127;
      let _683_v2;
      _683_v2 = _dafny.Map.Empty.slice().updateUnsafe(_681_v1,_678_v0);
      _683_v2 = (_683_v2).update(_681_v1, _678_v0);
      let _684_i3;
      _684_i3 = _dafny.ZERO;
      L5: {
        while (((_this).f28) && ((_this).f25)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_684_i3)) {
              break L5;
            }
            _684_i3 = (_684_i3).plus(_dafny.ONE);
            let _685_v3;
            _685_v3 = _dafny.MultiSet.fromElements(!(p0));
            let _686_v4;
            _686_v4 = _module.D10.create_DC29(_this.f24, (_685_v3).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f25, (_this).f25, (_this).f25))), p1);
            let _687_v5;
            _687_v5 = _dafny.Seq.of(_this.f24, p0);
            let _index105 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_678_v0).length));
            (_678_v0)[_index105] = !_dafny.areEqual(_687_v5, _dafny.Seq.of(_this.f24, (_this).f28, _this.f24, (_this).f25, _this.f24));
            let _index106 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_678_v0).length));
            let _rhs126 = _686_v4;
            let _rhs127 = _this.f24;
            let _rhs128 = (_this).f29;
            let _lhs106 = _678_v0;
            let _lhs107 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_678_v0).length));
            let _lhs108 = globalState;
            _686_v4 = _rhs126;
            _lhs106[_lhs107] = _rhs127;
            _lhs108.f14 = _rhs128;
            let _688_v6;
            _688_v6 = _dafny.Seq.of((_this).f29);
            let _689_v7;
            _689_v7 = _module.D10.create_DC28(_688_v6);
            let _690_v8;
            _690_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f23);
            let _691_v9;
            _691_v9 = _dafny.Map.Empty.slice().updateUnsafe((_689_v7).dtor_cf52,new BigNumber((_690_v8).length));
            let _692_v10;
            _692_v10 = _dafny.Map.Empty.slice().updateUnsafe(_691_v9,(_this).f29);
            let _693_v11;
            _693_v11 = _dafny.Seq.of((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f29, (_dafny.ZERO).minus((_this).f29))), new BigNumber(68), new BigNumber((_688_v6).length), new BigNumber(800), (((_692_v10).contains(_691_v9)) ? ((_692_v10).get(_691_v9)) : ((_this).f29)));
            _693_v11 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(581)), function (_694_i4) {
              return (_this).f23;
            });
            let _695_v12;
            let _nw128 = new _module.C1();
            _nw128.__ctor((_this.f24) || (_this.f24), p0);
            _695_v12 = _nw128;
            let _696_v13;
            _696_v13 = _dafny.Seq.of(_681_v1, _681_v1);
            let _697_v14;
            _697_v14 = _dafny.Map.Empty.slice().updateUnsafe(true,_681_v1);
            let _698_v15;
            _698_v15 = _dafny.Seq.of((((_678_v0)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_678_v0).length))]) ? ((_696_v13)[_module.__default.safeIndex((_this).f29, new BigNumber((_696_v13).length))]) : (_681_v1)), (((_this).f25) ? (_681_v1) : ((((_697_v14).contains(_this.f24)) ? ((_697_v14).get(_this.f24)) : (_681_v1)))), _681_v1, _681_v1);
            let _699_v16;
            _699_v16 = _module.D0.create_DC0(p0);
            let _700_v17;
            _700_v17 = new _dafny.CodePoint('e'.codePointAt(0));
            let _701_v18;
            _701_v18 = _dafny.Set.fromElements(_this.f24, (_678_v0)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_678_v0).length))], (_this).f28);
            let _rhs129 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_module.__default.fm25((_699_v16).dtor_cf0, (_this).f28, _700_v17, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_702_v17) => function (_703_i5) {
              return _702_v17;
            })(_700_v17)))).length));
            let _rhs130 = (_701_v18).IsProperSubsetOf(_701_v18);
            let _rhs131 = (_this).f29;
            let _rhs132 = _696_v13;
            let _lhs109 = globalState;
            let _lhs110 = globalState;
            let _lhs111 = globalState;
            _lhs109.f17 = _rhs129;
            _lhs110.f6 = _rhs130;
            _lhs111.f8 = _rhs131;
            _698_v15 = _rhs132;
          }
        }
      }
      let _704_v19;
      _704_v19 = new _dafny.CodePoint('v'.codePointAt(0));
      let _705_v20;
      _705_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_704_v19);
      let _706_v21;
      _706_v21 = _dafny.Seq.UnicodeFromString("vvejejg");
      (globalState).f6 = !(new BigNumber(((_705_v20).Merge(_705_v20)).length)).isEqualTo(new BigNumber((_706_v21).length));
      let _707_v22;
      _707_v22 = _dafny.Map.Empty.slice().updateUnsafe(p1,(new BigNumber(643)).minus((_this).f23));
      _707_v22 = (_707_v22).update(p1, (_this).f29);
      let _708_v23;
      _708_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
      let _709_v25;
      _709_v25 = _dafny.Seq.of((_this).f23);
      let _710_v26;
      let _nw129 = Array((new BigNumber(8)).toNumber());
      _nw129[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f23,_this.f24)).Merge(_708_v23);
      _nw129[(_dafny.ONE).toNumber()] = (_708_v23).Merge(_module.__default.fm34((_dafny.ZERO).minus(new BigNumber((_706_v21).length)), globalState));
      _nw129[(new BigNumber(2)).toNumber()] = _708_v23;
      _nw129[(new BigNumber(3)).toNumber()] = _708_v23;
      _nw129[(new BigNumber(4)).toNumber()] = _708_v23;
      _nw129[(new BigNumber(5)).toNumber()] = function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of (_709_v25).Elements) {
          let _711_v24 = _compr_24;
          if (_dafny.Seq.contains(_709_v25, _711_v24)) {
            _coll24.push([(_711_v24).plus((_this).f29),(_this).f25]);
          }
        }
        return _coll24;
      }();
      _nw129[(new BigNumber(6)).toNumber()] = _708_v23;
      _nw129[(new BigNumber(7)).toNumber()] = _708_v23;
      _710_v26 = _nw129;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_710_v26).length))) {
        let _712_i6 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_712_i6)) && ((_712_i6).isLessThan(new BigNumber((_710_v26).length))))) {
          (_710_v26)[(_712_i6)] = (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-59)),_this.f24)).Merge(_708_v23);
        }
      }
      return;
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f24 = false;
      this._f25 = false;
      this._f27 = false;
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    set f24(value) {
      let _this = this;
      _this._f24 = value;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f27, f24, f25) {
      let _this = this;
      (_this)._f27 = f27;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_this.f24))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(_this.f24,(_module.D0.create_DC1(!((_this).f25), new BigNumber((_dafny.Seq.UnicodeFromString("aeynwfee")).length))).dtor_cf1)))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f27))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_this.f24),_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_this.f24))));
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-771)),_dafny.MultiSet.fromElements(new BigNumber(233)))).Merge(function () {
        let _coll25 = new _dafny.Map();
        for (const _compr_25 of _dafny.IntegerRange(new BigNumber(923), new BigNumber(571))) {
          let _713_v0 = _compr_25;
          if (((new BigNumber(923)).isLessThanOrEqualTo(_713_v0)) && ((_713_v0).isLessThan(new BigNumber(571)))) {
            _coll25.push([(_713_v0).minus(new BigNumber(248)),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("mpuhcuieu")).length), new BigNumber((function () {
              let _coll26 = new _dafny.Set();
              for (const _compr_26 of (_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(362))).length))).Elements) {
                let _714_v1 = _compr_26;
                if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(362))).length)), _714_v1)) {
                  _coll26.add(_module.__default.safeDivisionInt(_714_v1, new BigNumber((_dafny.Seq.UnicodeFromString("t")).length)));
                }
              }
              return _coll26;
            }()).length), new BigNumber(220))]);
          }
        }
        return _coll25;
      }())).length);
    };
    fm14(globalState) {
      let _this = this;
      return (_this).f27;
    };
    fm15(p0, p1, p2, globalState) {
      let _this = this;
      let _source6 = _module.D1.create_DC4((_this).f27, _this.f24);
      if (_source6.is_DC4) {
        let _715___mcc_h0 = (_source6).cf5;
        let _716___mcc_h1 = (_source6).cf6;
        let _717_cf6 = _716___mcc_h1;
        let _718_cf5 = _715___mcc_h0;
        return new BigNumber(981);
      } else if (_source6.is_DC5) {
        let _719___mcc_h2 = (_source6).cf7;
        let _720___mcc_h3 = (_source6).cf8;
        let _721___mcc_h4 = (_source6).cf9;
        let _722___mcc_h5 = (_source6).cf10;
        let _723___mcc_h6 = (_source6).cf11;
        let _724_cf11 = _723___mcc_h6;
        let _725_cf10 = _722___mcc_h5;
        let _726_cf9 = _721___mcc_h4;
        let _727_cf8 = _720___mcc_h3;
        let _728_cf7 = _719___mcc_h2;
        return _725_cf10;
      } else if (_source6.is_DC3) {
        let _729___mcc_h7 = (_source6).cf4;
        let _730_cf4 = _729___mcc_h7;
        return new BigNumber(810);
      } else {
        let _731___mcc_h8 = (_source6).cf12;
        let _732_cf12 = _731___mcc_h8;
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("caeqeee"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_733_i0) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        }))).length);
      }
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = undefined;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _734_v0;
      _734_v0 = new BigNumber(428);
      let _hi5 = _734_v0;
      for (let _735_i0 = (_734_v0).multipliedBy(_734_v0); _735_i0.isLessThan(_hi5); _735_i0 = _735_i0.plus(_dafny.ONE)) {
        (globalState).f15 = !((_this).f27) || (_this.f24);
        let _736_v1;
        let _nw130 = Array((new BigNumber(27)).toNumber()).fill(false);
        _736_v1 = _nw130;
        let _index107 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_736_v1).length));
        (_736_v1)[_index107] = (_this).f27;
        let _737_v2;
        _737_v2 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f24),(_this).f27);
        let _738_v3;
        _738_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber((_737_v2).length));
        let _739_v4;
        _739_v4 = _module.D1.create_DC5((_this).f27, _737_v2, _735_i0, new BigNumber(996), _738_v3);
        let _740_v5;
        _740_v5 = _module.D1.create_DC6(_739_v4);
        let _pat_let_tv11 = _739_v4;
        let _741_v6;
        _741_v6 = _dafny.MultiSet.fromElements(function (_pat_let5_0) {
          return function (_742_dt__update__tmp_h0) {
            return function (_pat_let6_0) {
              return function (_743_dt__update_hcf12_h0) {
                return _module.D1.create_DC6(_743_dt__update_hcf12_h0);
              }(_pat_let6_0);
            }(_pat_let_tv11);
          }(_pat_let5_0);
        }(_740_v5), _740_v5, _740_v5);
        let _index108 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_736_v1).length));
        (_736_v1)[_index108] = ((_741_v6).IsProperSubsetOf(_741_v6)) || (false);
        let _744_v7;
        _744_v7 = _module.D1.create_DC5(false, (_737_v2).Merge(_737_v2), _734_v0, _735_i0, _738_v3);
        _744_v7 = _744_v7;
        let _745_v8;
        let _nw131 = Array((_dafny.ONE).toNumber());
        _nw131[(_dafny.ZERO).toNumber()] = _736_v1;
        _745_v8 = _nw131;
        let _index109 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_745_v8).length));
        (_745_v8)[_index109] = _736_v1;
        let _746_v9;
        let _nw132 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
        _746_v9 = _nw132;
        let _747_v10;
        let _init22 = ((_748_i0) => function (_749_i1) {
          return (_749_i1).minus(_748_i0);
        })(_735_i0);
        let _nw133 = Array((new BigNumber(29)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw133.length); _i0_22++) {
          _nw133[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _747_v10 = _nw133;
        let _750_v11;
        _750_v11 = _dafny.Map.Empty.slice().updateUnsafe(_734_v0,_747_v10);
        let _index110 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_746_v9).length));
        (_746_v9)[_index110] = _750_v11;
        let _751_v12;
        _751_v12 = _736_v1;
        let _index111 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_745_v8).length));
        let _index112 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_746_v9).length));
        let _rhs133 = (_736_v1);
        let _rhs134 = _735_i0;
        let _rhs135 = _750_v11;
        let _rhs136 = _734_v0;
        let _lhs112 = _745_v8;
        let _lhs113 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_745_v8).length));
        let _lhs114 = globalState;
        let _lhs115 = _746_v9;
        let _lhs116 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_746_v9).length));
        let _lhs117 = globalState;
        _lhs112[_lhs113] = _rhs133;
        _lhs114.f14 = _rhs134;
        _lhs115[_lhs116] = _rhs135;
        _lhs117.f14 = _rhs136;
      }
      let _752_v13;
      let _init23 = function (_753_i2) {
        return _this.f24;
      };
      let _nw134 = Array((new BigNumber(26)).toNumber());
      for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw134.length); _i0_23++) {
        _nw134[_i0_23] = _init23(new BigNumber(_i0_23));
      }
      _752_v13 = _nw134;
      let _index113 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length));
      (_752_v13)[_index113] = (_this).f25;
      let _754_v14;
      _754_v14 = _dafny.Map.Empty.slice().updateUnsafe(_734_v0,_734_v0);
      let _755_v16;
      _755_v16 = _dafny.MultiSet.fromElements(_754_v14, function () {
        let _coll27 = new _dafny.Map();
        for (const _compr_27 of _dafny.IntegerRange(new BigNumber(-595), new BigNumber(441))) {
          let _756_v15 = _compr_27;
          if (((new BigNumber(-595)).isLessThanOrEqualTo(_756_v15)) && ((_756_v15).isLessThan(new BigNumber(441)))) {
            _coll27.push([_module.__default.safeModuloInt(_756_v15, _734_v0),_734_v0]);
          }
        }
        return _coll27;
      }(), _754_v14);
      let _757_v17;
      _757_v17 = _dafny.Map.Empty.slice().updateUnsafe(_755_v16,_734_v0);
      let _index114 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length));
      let _rhs137 = !((_734_v0).isLessThan(_734_v0)) || (_this.f24);
      let _rhs138 = _this.f24;
      let _rhs139 = _this.f24;
      let _rhs140 = (((_757_v17).contains(_755_v16)) ? ((_757_v17).get(_755_v16)) : ((_734_v0).minus(_734_v0)));
      let _lhs118 = globalState;
      let _lhs119 = _this;
      let _lhs120 = _752_v13;
      let _lhs121 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length));
      let _lhs122 = globalState;
      _lhs118.f9 = _rhs137;
      _lhs119.f24 = _rhs138;
      _lhs120[_lhs121] = _rhs139;
      _lhs122.f5 = _rhs140;
      let _758_v18;
      _758_v18 = _dafny.Seq.of((_752_v13)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length))], false);
      let _759_v19;
      _759_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_this.f24);
      let _760_v20;
      _760_v20 = _dafny.MultiSet.fromElements((((_759_v19).contains(_this.f24)) ? ((_759_v19).get(_this.f24)) : ((_this).f25)));
      let _761_v21;
      _761_v21 = _dafny.Seq.UnicodeFromString("lwbojnlo");
      let _762_v22;
      _762_v22 = _dafny.MultiSet.fromElements(_734_v0, _734_v0);
      let _763_v23;
      _763_v23 = _dafny.Set.fromElements(_734_v0);
      let _764_v24;
      let _nw135 = Array((new BigNumber(19)).toNumber());
      _nw135[(_dafny.ZERO).toNumber()] = _734_v0;
      _nw135[(_dafny.ONE).toNumber()] = _734_v0;
      _nw135[(new BigNumber(2)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(3)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(4)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(5)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.update(_758_v18, _module.__default.safeIndex(_734_v0, new BigNumber((_758_v18).length)), _this.f24)).length);
      _nw135[(new BigNumber(7)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(8)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(9)).toNumber()] = new BigNumber((_760_v20).cardinality());
      _nw135[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("x")).length);
      _nw135[(new BigNumber(11)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(12)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(13)).toNumber()] = new BigNumber((_761_v21).length);
      _nw135[(new BigNumber(14)).toNumber()] = (((_762_v22).contains(_734_v0)) ? ((_762_v22).get(_734_v0)) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(_734_v0))));
      _nw135[(new BigNumber(15)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(16)).toNumber()] = new BigNumber((_763_v23).length);
      _nw135[(new BigNumber(17)).toNumber()] = _734_v0;
      _nw135[(new BigNumber(18)).toNumber()] = _734_v0;
      _764_v24 = _nw135;
      let _765_v25;
      _765_v25 = _dafny.Seq.of(_764_v24);
      let _766_v26;
      _766_v26 = _dafny.Seq.of(new BigNumber(662), _734_v0, new BigNumber((_761_v21).length));
      let _767_v27;
      _767_v27 = _dafny.Seq.of((_766_v26)[_module.__default.safeIndex(_734_v0, new BigNumber((_766_v26).length))]);
      let _hi6 = (new BigNumber((_dafny.Seq.update(_765_v25, _module.__default.safeIndex(new BigNumber((_766_v26).length), new BigNumber((_765_v25).length)), (_765_v25)[_module.__default.safeIndex(_734_v0, new BigNumber((_765_v25).length))])).length)).minus(new BigNumber(845));
      for (let _768_i3 = _734_v0; _768_i3.isLessThan(_hi6); _768_i3 = _768_i3.plus(_dafny.ONE)) {
        if ((_this).f25) {
          _761_v21 = _761_v21;
          _758_v18 = _758_v18;
          let _769_v28;
          let _init24 = ((_770_i3, _771_v21) => function (_772_i4) {
            return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),_770_i3)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber(((_module.D3.create_DC8(_771_v21)).dtor_cf14).length)));
          })(_768_i3, _761_v21);
          let _nw136 = Array((new BigNumber(26)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw136.length); _i0_24++) {
            _nw136[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _769_v28 = _nw136;
          let _773_v29;
          _773_v29 = new _dafny.CodePoint('n'.codePointAt(0));
          let _774_v30;
          _774_v30 = _dafny.Map.Empty.slice().updateUnsafe(_773_v29,_768_i3);
          let _index115 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_769_v28).length));
          (_769_v28)[_index115] = _774_v30;
          let _index116 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_769_v28).length));
          (_769_v28)[_index116] = (_dafny.Map.Empty.slice().updateUnsafe(_773_v29,_734_v0)).update(_773_v29, _768_i3);
          _773_v29 = (((_768_i3).isLessThan(_734_v0)) ? ((_761_v21)[_module.__default.safeIndex(_734_v0, new BigNumber((_761_v21).length))]) : (_773_v29));
          let _775_v31;
          _775_v31 = _module.D0.create_DC1((_752_v13)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length))], _734_v0);
          let _776_v32;
          _776_v32 = _module.D0.create_DC2(_775_v31);
          let _777_v33;
          _777_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_776_v32);
          _777_v33 = (_777_v33).update(_this.f24, _776_v32);
        } else {
          let _778_v34;
          _778_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm14(globalState),_752_v13);
          _778_v34 = _778_v34;
          _759_v19 = (_759_v19).Merge((_759_v19).Merge(_759_v19));
          (globalState).f6 = (_module.D0.create_DC0(_this.f24)).dtor_cf0;
          (globalState).f9 = (_734_v0).isLessThanOrEqualTo((_768_i3).minus(new BigNumber(468)));
          let _779_v35;
          _779_v35 = _dafny.MultiSet.fromElements(_762_v22);
          _779_v35 = _779_v35;
        }
        let _780_v36;
        let _init25 = ((_781_v18) => function (_782_i5) {
          return _781_v18;
        })(_758_v18);
        let _nw137 = Array((new BigNumber(15)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw137.length); _i0_25++) {
          _nw137[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _780_v36 = _nw137;
        let _783_v37;
        _783_v37 = _module.D3.create_DC9(new BigNumber((_761_v21).length));
        let _784_v38;
        _784_v38 = _dafny.Map.Empty.slice().updateUnsafe(_783_v37,new BigNumber(666));
        let _785_v39;
        _785_v39 = _dafny.Map.Empty.slice().updateUnsafe(_784_v38,_780_v36);
        let _786_v40;
        let _nw138 = Array((new BigNumber(19)).toNumber());
        _nw138[(_dafny.ZERO).toNumber()] = _780_v36;
        _nw138[(_dafny.ONE).toNumber()] = _780_v36;
        _nw138[(new BigNumber(2)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(3)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(4)).toNumber()] = (((_785_v39).contains(_784_v38)) ? ((_785_v39).get(_784_v38)) : (_780_v36));
        _nw138[(new BigNumber(5)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(6)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(7)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(8)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(9)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(10)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(11)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(12)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(13)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(14)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(15)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(16)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(17)).toNumber()] = _780_v36;
        _nw138[(new BigNumber(18)).toNumber()] = _780_v36;
        _786_v40 = _nw138;
        let _index117 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_786_v40).length));
        (_786_v40)[_index117] = _780_v36;
        let _787_v41;
        _787_v41 = _dafny.Set.fromElements(_766_v26);
        let _index118 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_786_v40).length));
        let _rhs141 = (_this).fm15((new BigNumber((_787_v41).length)).plus(_734_v0), new BigNumber(375), (_this).fm14(globalState), globalState);
        let _rhs142 = _780_v36;
        let _lhs123 = globalState;
        let _lhs124 = _786_v40;
        let _lhs125 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_786_v40).length));
        _lhs123.f14 = _rhs141;
        _lhs124[_lhs125] = _rhs142;
        if ((_this).f27) {
          r0 = (_this).f25;
          let _788_v42;
          _788_v42 = new _dafny.CodePoint('o'.codePointAt(0));
          let _789_v43;
          _789_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_788_v42);
          let _index119 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_764_v24).length));
          (_764_v24)[_index119] = new BigNumber((_789_v43).length);
          let _index120 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_764_v24).length));
          (_764_v24)[_index120] = _768_i3;
          let _790_v44;
          _790_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_761_v21).length),_752_v13);
          _790_v44 = (_790_v44).update(new BigNumber(94), _752_v13);
          let _791_v45;
          _791_v45 = _dafny.Map.Empty.slice().updateUnsafe(_752_v13,_762_v22);
          let _792_v46;
          _792_v46 = _dafny.Seq.of(_762_v22, _dafny.MultiSet.FromArray(_767_v27));
          _791_v45 = (_791_v45).update(_752_v13, (_792_v46)[_module.__default.safeIndex(_768_i3, new BigNumber((_792_v46).length))]);
          (globalState).f5 = _734_v0;
        } else {
          (globalState).f8 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_768_i3, (_dafny.ZERO).minus((_768_i3).minus(_768_i3))));
          let _index121 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length));
          (_752_v13)[_index121] = (_this).fm14(globalState);
          r0 = (_this).f27;
          let _793_v47;
          let _nw139 = new _module.C1();
          _nw139.__ctor(_this.f24, true);
          _793_v47 = _nw139;
          let _794_v48;
          _794_v48 = new _dafny.CodePoint('c'.codePointAt(0));
          let _795_v49;
          let _nw140 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _795_v49 = _nw140;
          let _796_v50;
          let _nw141 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
          _796_v50 = _nw141;
          let _rhs143 = (_761_v21)[_module.__default.safeIndex(_734_v0, new BigNumber((_761_v21).length))];
          let _rhs144 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_768_i3, new BigNumber((_758_v18).length)));
          let _rhs145 = _795_v49;
          let _rhs146 = _761_v21;
          let _rhs147 = (((_752_v13)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length))]) ? (_796_v50) : (_796_v50));
          let _lhs126 = globalState;
          _794_v48 = _rhs143;
          _lhs126.f8 = _rhs144;
          _795_v49 = _rhs145;
          _761_v21 = _rhs146;
          _796_v50 = _rhs147;
        }
        let _797_v51;
        let _nw142 = new _module.C3();
        _nw142.__ctor(!((_762_v22).IsDisjointFrom(_dafny.MultiSet.fromElements(_734_v0))), (_dafny.ZERO).minus(_768_i3), (_this).f27, _this.f24, _734_v0);
        _797_v51 = _nw142;
      }
      let _798_v52;
      _798_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_764_v24);
      _798_v52 = _798_v52;
      (globalState).f5 = _734_v0;
      let _799_i6;
      _799_i6 = _dafny.ZERO;
      L6: {
        while (_this.f24) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_799_i6)) {
              break L6;
            }
            _799_i6 = (_799_i6).plus(_dafny.ONE);
            let _800_v53;
            _800_v53 = new _dafny.CodePoint('d'.codePointAt(0));
            let _801_v54;
            _801_v54 = _dafny.Seq.of(_761_v21, _dafny.Seq.update(_761_v21, _module.__default.safeIndex(_734_v0, new BigNumber((_761_v21).length)), _800_v53));
            _801_v54 = _801_v54;
            let _nw143 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
            _764_v24 = _nw143;
            let _802_v55;
            _802_v55 = _dafny.Set.fromElements(!(_this.f24), true);
            let _803_v56;
            _803_v56 = _dafny.Map.Empty.slice().updateUnsafe(_802_v55,(_752_v13)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length))]);
            let _804_v58;
            _804_v58 = _dafny.Seq.of(_803_v56);
            let _805_v59;
            _805_v59 = _dafny.Map.Empty.slice().updateUnsafe(_734_v0,_803_v56);
            let _806_v61;
            let _nw144 = Array((new BigNumber(14)).toNumber());
            _nw144[(_dafny.ZERO).toNumber()] = _803_v56;
            _nw144[(_dafny.ONE).toNumber()] = _803_v56;
            _nw144[(new BigNumber(2)).toNumber()] = (((_this).f27) ? (function () {
              let _coll28 = new _dafny.Map();
              for (const _compr_28 of (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), ((_807_v55) => function (_808_i7) {
                return _807_v55;
              })(_802_v55)))).Elements) {
                let _809_v57 = _compr_28;
                if ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), ((_810_v55) => function (_808_i7) {
                  return _810_v55;
                })(_802_v55)))).contains(_809_v57)) {
                  _coll28.push([_809_v57,(_this).f27]);
                }
              }
              return _coll28;
            }()) : (_803_v56));
            _nw144[(new BigNumber(3)).toNumber()] = (_803_v56).Merge((_804_v58)[_module.__default.safeIndex(_734_v0, new BigNumber((_804_v58).length))]);
            _nw144[(new BigNumber(4)).toNumber()] = (((_805_v59).contains(_734_v0)) ? ((_805_v59).get(_734_v0)) : (_module.__default.fm35(globalState)));
            _nw144[(new BigNumber(5)).toNumber()] = _803_v56;
            _nw144[(new BigNumber(6)).toNumber()] = (_803_v56).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_752_v13)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length))], (_this).f25),_this.f24));
            _nw144[(new BigNumber(7)).toNumber()] = _803_v56;
            _nw144[(new BigNumber(8)).toNumber()] = (_803_v56).Merge(_803_v56);
            _nw144[(new BigNumber(9)).toNumber()] = _803_v56;
            _nw144[(new BigNumber(10)).toNumber()] = _803_v56;
            _nw144[(new BigNumber(11)).toNumber()] = (((_752_v13)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_752_v13).length))]) ? (function () {
              let _coll29 = new _dafny.Map();
              for (const _compr_29 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), ((_811_v55) => function (_812_i8) {
                return _811_v55;
              })(_802_v55))).Elements) {
                let _813_v60 = _compr_29;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), ((_814_v55) => function (_812_i8) {
                  return _814_v55;
                })(_802_v55)), _813_v60)) {
                  _coll29.push([_813_v60,(_this).f27]);
                }
              }
              return _coll29;
            }()) : (_803_v56));
            _nw144[(new BigNumber(12)).toNumber()] = _803_v56;
            _nw144[(new BigNumber(13)).toNumber()] = _803_v56;
            _806_v61 = _nw144;
            let _index122 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_806_v61).length));
            (_806_v61)[_index122] = _803_v56;
            let _index123 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_806_v61).length));
            (_806_v61)[_index123] = _803_v56;
            (globalState).f1 = _this.f24;
          }
        }
      }
      r0 = ((((_this).f25) === ((_this).f25)) ? (!((_this).f27)) : ((new BigNumber(235)).isEqualTo(_734_v0)));
      let _815_v62;
      let _nw145 = new _module.C2();
      _nw145.__ctor(new BigNumber((_dafny.Seq.of(_this.f24, (_this).f27)).length));
      _815_v62 = _nw145;
      let _816_v63;
      _816_v63 = _module.D11.create_DC30(_815_v62);
      r1 = (_816_v63).dtor_cf56;
      r2 = (_this.f24) === (_this.f24);
      let _817_v65;
      _817_v65 = _dafny.MultiSet.fromElements(_758_v18);
      r3 = (((_this).f25) ? (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_758_v18,false)).Merge(function () {
        let _coll30 = new _dafny.Map();
        for (const _compr_30 of (_817_v65).Elements) {
          let _818_v64 = _compr_30;
          if ((_817_v65).contains(_818_v64)) {
            _coll30.push([_818_v64,_this.f24]);
          }
        }
        return _coll30;
      }())).length)) : (_734_v0));
      return [r0, r1, r2, r3];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f24 = false;
      this._f25 = false;
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    set f24(value) {
      let _this = this;
      _this._f24 = value;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f24, f25) {
      let _this = this;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      let _source7 = _module.D0.create_DC0((_this).f25);
      if (_source7.is_DC1) {
        let _819___mcc_h0 = (_source7).cf1;
        let _820___mcc_h1 = (_source7).cf2;
        let _821_cf2 = _820___mcc_h1;
        let _822_cf1 = _819___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_dafny.Map.Empty.slice().updateUnsafe(_822_cf1,!(false)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_822_cf1,_dafny.Map.Empty.slice().updateUnsafe(true,_822_cf1)));
      } else if (_source7.is_DC0) {
        let _823___mcc_h2 = (_source7).cf0;
        let _824_cf0 = _823___mcc_h2;
        return _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25));
      } else {
        let _825___mcc_h3 = (_source7).cf3;
        let _826_cf3 = _825___mcc_h3;
        return (_dafny.Map.Empty.slice().updateUnsafe(!(_this.f24),_dafny.Map.Empty.slice().updateUnsafe(true,_this.f24))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25)));
      }
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(5)), function (_827_i0) {
        return _dafny.Set.fromElements(new BigNumber(86), new BigNumber(469), new BigNumber((_dafny.Seq.of(_this.f24)).length));
      }), _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(224))))).length);
    };
    fm13(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((new BigNumber((_dafny.Seq.of(new BigNumber(939))).length)).plus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(634))).length), new BigNumber((_dafny.Seq.UnicodeFromString("i")).length))).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(680)), function (_828_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      })).length));
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = undefined;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _829_v0;
      _829_v0 = new BigNumber(82);
      let _830_v1;
      _830_v1 = new _dafny.CodePoint('v'.codePointAt(0));
      let _831_v2;
      _831_v2 = _dafny.Seq.of(_830_v1, _830_v1);
      let _hi7 = new BigNumber((_dafny.Seq.Concat(_831_v2, _dafny.Seq.of(_830_v1, _830_v1))).length);
      for (let _832_i0 = _829_v0; _832_i0.isLessThan(_hi7); _832_i0 = _832_i0.plus(_dafny.ONE)) {
        let _833_v3;
        let _nw146 = new _module.C2();
        _nw146.__ctor(_829_v0);
        _833_v3 = _nw146;
        if (!(_this.f24)) {
          let _834_v4;
          _834_v4 = _dafny.Map.Empty.slice().updateUnsafe(_832_i0,new BigNumber(134));
          _834_v4 = (_834_v4).update(_832_i0, _832_i0);
          (globalState).f1 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("radkjif"), _831_v2);
          let _835_v5;
          let _init26 = function (_836_i1) {
            return (_this).f25;
          };
          let _nw147 = Array((new BigNumber(10)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw147.length); _i0_26++) {
            _nw147[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _835_v5 = _nw147;
          let _837_v6;
          let _nw148 = new _module.C0();
          _nw148.__ctor(_dafny.Set.fromElements(_835_v5, _835_v5, _835_v5, _835_v5), (_this).fm13(new BigNumber((_dafny.Seq.UnicodeFromString("arwqolr")).length), (_this).f25, globalState), new BigNumber(-476));
          _837_v6 = _nw148;
          let _838_v7;
          _838_v7 = _dafny.Map.Empty.slice().updateUnsafe(_837_v6,_this.f24);
          _838_v7 = (_838_v7).update(_837_v6, (_this).f25);
          let _839_v8;
          _839_v8 = _dafny.Map.Empty.slice().updateUnsafe(_835_v5,_837_v6.f31);
          let _840_v9;
          _840_v9 = _dafny.Set.fromElements(_this.f24, false, false);
          let _841_v10;
          let _nw149 = Array((new BigNumber(2)).toNumber());
          _nw149[(_dafny.ZERO).toNumber()] = new BigNumber((_840_v9).length);
          _nw149[(_dafny.ONE).toNumber()] = _829_v0;
          _841_v10 = _nw149;
          let _842_v11;
          _842_v11 = _module.D11.create_DC31(_841_v10, _832_i0, _this.f24);
          let _843_v12;
          let _nw150 = Array((new BigNumber(8)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = _830_v1;
          _nw150[(_dafny.ONE).toNumber()] = _module.__default.fm26((_842_v11).dtor_cf58, (_this).f25, globalState);
          _nw150[(new BigNumber(2)).toNumber()] = _830_v1;
          _nw150[(new BigNumber(3)).toNumber()] = _830_v1;
          _nw150[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _nw150[(new BigNumber(5)).toNumber()] = _830_v1;
          _nw150[(new BigNumber(6)).toNumber()] = _830_v1;
          _nw150[(new BigNumber(7)).toNumber()] = _830_v1;
          _843_v12 = _nw150;
          let _index124 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_843_v12).length));
          (_843_v12)[_index124] = _830_v1;
          let _844_v13;
          let _nw151 = Array((new BigNumber(24)).toNumber()).fill([]);
          _844_v13 = _nw151;
          let _index125 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_844_v13).length));
          (_844_v13)[_index125] = _843_v12;
          let _index126 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_835_v5).length));
          (_835_v5)[_index126] = (_this).f25;
          let _845_v14;
          _845_v14 = _dafny.Seq.of((_this).f25, _this.f24);
          let _846_v15;
          _846_v15 = _dafny.MultiSet.fromElements((_this).f25, _this.f24, ((_this).f25) && (!(_this.f24)), (_845_v14)[_module.__default.safeIndex((_833_v3).fm2(globalState), new BigNumber((_845_v14).length))], (_this).f25);
          let _index127 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_843_v12).length));
          let _index128 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_844_v13).length));
          let _index129 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_835_v5).length));
          let _rhs148 = _839_v8;
          let _rhs149 = _830_v1;
          let _rhs150 = (((_846_v15).contains(_dafny.Seq.contains(_module.__default.fm25(true, _this.f24, _830_v1, globalState), new _dafny.CodePoint('n'.codePointAt(0))))) ? ((_846_v15).get(_dafny.Seq.contains(_module.__default.fm25(true, _this.f24, _830_v1, globalState), new _dafny.CodePoint('n'.codePointAt(0))))) : (new BigNumber(191)));
          let _rhs151 = _843_v12;
          let _rhs152 = _module.__default.fm29(new BigNumber((_831_v2).length), _837_v6.f31, _module.__default.fm29(_837_v6.f31, _837_v6.f31, _this.f24, globalState), globalState);
          let _lhs127 = _843_v12;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_843_v12).length));
          let _lhs129 = globalState;
          let _lhs130 = _844_v13;
          let _lhs131 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_844_v13).length));
          let _lhs132 = _835_v5;
          let _lhs133 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_835_v5).length));
          _839_v8 = _rhs148;
          _lhs127[_lhs128] = _rhs149;
          _lhs129.f17 = _rhs150;
          _lhs130[_lhs131] = _rhs151;
          _lhs132[_lhs133] = _rhs152;
          r0 = (_835_v5)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_835_v5).length))];
        } else {
          let _847_v16;
          _847_v16 = _dafny.MultiSet.fromElements(!(_this.f24));
          let _848_v17;
          _848_v17 = _dafny.Seq.of((_847_v16).IsDisjointFrom(_847_v16), (_this).f25);
          let _849_v18;
          _849_v18 = _dafny.Map.Empty.slice().updateUnsafe(_830_v1,_829_v0);
          (globalState).f15 = (_848_v17)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_849_v18).length), _829_v0), new BigNumber((_848_v17).length))];
          (globalState).f14 = _829_v0;
          let _850_v19;
          let _nw152 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _850_v19 = _nw152;
          let _index130 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_850_v19).length));
          (_850_v19)[_index130] = (((_848_v17)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_848_v17).length))]) ? (_831_v2) : (_831_v2));
          let _index131 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_850_v19).length));
          (_850_v19)[_index131] = _dafny.Seq.Concat(_dafny.Seq.Concat(_831_v2, _831_v2), _831_v2);
          let _851_v20;
          _851_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_848_v17).length),new BigNumber(((_850_v19)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_850_v19).length))]).length));
          let _852_v21;
          let _nw153 = new _module.C2();
          _nw153.__ctor(new BigNumber((_851_v20).length));
          _852_v21 = _nw153;
          let _853_v22;
          _853_v22 = _module.D11.create_DC30(_852_v21);
          _853_v22 = _853_v22;
          let _854_v23;
          let _nw154 = Array((_dafny.ONE).toNumber());
          _nw154[(_dafny.ZERO).toNumber()] = (_this).f25;
          _854_v23 = _nw154;
          let _855_v24;
          let _nw155 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _855_v24 = _nw155;
          let _856_v25;
          _856_v25 = _module.D11.create_DC31(_855_v24, _829_v0, (_848_v17)[_module.__default.safeIndex((_852_v21).f23, new BigNumber((_848_v17).length))]);
          let _index132 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_854_v23).length));
          (_854_v23)[_index132] = (_856_v25).dtor_cf59;
          let _index133 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_854_v23).length));
          (_854_v23)[_index133] = (_this).f25;
        }
        if ((_this).f25) {
          let _857_v26;
          let _nw156 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _857_v26 = _nw156;
          let _index134 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_857_v26).length));
          (_857_v26)[_index134] = _829_v0;
          let _858_v27;
          _858_v27 = _dafny.Seq.of(_857_v26);
          let _index135 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_857_v26).length));
          let _rhs153 = _829_v0;
          let _rhs154 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(-158), new BigNumber(526)))).isLessThanOrEqualTo(_832_i0);
          let _rhs155 = new BigNumber((_858_v27).length);
          let _rhs156 = (_this).f25;
          let _rhs157 = _832_i0;
          let _lhs134 = globalState;
          let _lhs135 = globalState;
          let _lhs136 = _857_v26;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_857_v26).length));
          let _lhs138 = globalState;
          let _lhs139 = globalState;
          _lhs134.f17 = _rhs153;
          _lhs135.f9 = _rhs154;
          _lhs136[_lhs137] = _rhs155;
          _lhs138.f1 = _rhs156;
          _lhs139.f14 = _rhs157;
          let _859_v28;
          _859_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_830_v1);
          let _860_v29;
          _860_v29 = _dafny.MultiSet.fromElements(_859_v28, _859_v28);
          (globalState).f15 = (_860_v29).IsProperSubsetOf(_860_v29);
          let _index136 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_857_v26).length));
          let _rhs158 = _858_v27;
          let _rhs159 = _832_i0;
          let _rhs160 = _829_v0;
          let _lhs140 = _857_v26;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_857_v26).length));
          _858_v27 = _rhs158;
          _lhs140[_lhs141] = _rhs159;
          _829_v0 = _rhs160;
          (globalState).f8 = (_829_v0).plus(_832_i0);
          let _rhs161 = true;
          let _rhs162 = !((_this).f25);
          let _rhs163 = (_857_v26)[_module.__default.safeIndex(new BigNumber(787), new BigNumber((_857_v26).length))];
          let _rhs164 = _832_i0;
          let _lhs142 = globalState;
          let _lhs143 = globalState;
          let _lhs144 = globalState;
          let _lhs145 = globalState;
          _lhs142.f1 = _rhs161;
          _lhs143.f1 = _rhs162;
          _lhs144.f14 = _rhs163;
          _lhs145.f17 = _rhs164;
        } else {
          (globalState).f1 = _this.f24;
          (globalState).f17 = _829_v0;
          let _861_v30;
          let _nw157 = Array((new BigNumber(26)).toNumber()).fill(false);
          _861_v30 = _nw157;
          let _862_v31;
          _862_v31 = _dafny.Seq.of(_861_v30, _861_v30, (((_this).f25) ? (_861_v30) : (_861_v30)));
          _862_v31 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_862_v31, _862_v31), _module.__default.safeIndex(_832_i0, new BigNumber((_dafny.Seq.Concat(_862_v31, _862_v31)).length)), _861_v30), _862_v31);
          (globalState).f8 = (_module.__default.safeDivisionInt(_829_v0, (_dafny.ZERO).minus(_832_i0))).minus(_832_i0);
          let _863_v32;
          _863_v32 = _dafny.Set.fromElements(!(_this.f24));
          let _864_v33;
          let _nw158 = Array((new BigNumber(3)).toNumber());
          _nw158[(_dafny.ZERO).toNumber()] = _863_v32;
          _nw158[(_dafny.ONE).toNumber()] = _863_v32;
          _nw158[(new BigNumber(2)).toNumber()] = _863_v32;
          _864_v33 = _nw158;
          let _index137 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_864_v33).length));
          (_864_v33)[_index137] = _863_v32;
          let _865_v34;
          _865_v34 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_829_v0);
          let _866_v35;
          let _nw159 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _866_v35 = _nw159;
          let _867_v36;
          _867_v36 = _module.D7.create_DC20((_this).f25, _865_v34, _866_v35);
          let _868_v37;
          _868_v37 = _module.D7.create_DC21(_867_v36);
          let _869_v38;
          _869_v38 = _module.D7.create_DC21(_868_v37);
          let _pat_let_tv12 = _865_v34;
          let _pat_let_tv13 = _866_v35;
          let _870_v39;
          _870_v39 = _dafny.MultiSet.fromElements(function (_pat_let7_0) {
            return function (_871_dt__update__tmp_h0) {
              return function (_pat_let8_0) {
                return function (_872_dt__update_hcf38_h0) {
                  return _module.D7.create_DC21(_872_dt__update_hcf38_h0);
                }(_pat_let8_0);
              }(_module.D7.create_DC20((_this).f25, _pat_let_tv12, _pat_let_tv13));
            }(_pat_let7_0);
          }(_869_v38), _869_v38);
          let _873_v40;
          _873_v40 = _dafny.Seq.of(new BigNumber((_870_v39).cardinality()));
          let _874_v41;
          _874_v41 = _dafny.Seq.of((_this).f25, _this.f24);
          let _875_v42;
          _875_v42 = _module.D1.create_DC4((_this).f25, _this.f24);
          let _876_v43;
          _876_v43 = _dafny.Set.fromElements(_875_v42);
          let _index138 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_864_v33).length));
          (_864_v33)[_index138] = (_863_v32).Intersect((_module.__default.fm32((_this).f25, _873_v40, _874_v41, _876_v43, globalState)).Union(_863_v32));
        }
        let _877_v44;
        let _nw160 = Array((new BigNumber(25)).toNumber()).fill(false);
        _877_v44 = _nw160;
        let _878_v45;
        _878_v45 = _dafny.Map.Empty.slice().updateUnsafe(false,false);
        let _879_v46;
        _879_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_832_i0);
        let _880_v47;
        _880_v47 = _module.D1.create_DC5((_this).f25, _878_v45, _832_i0, _832_i0, _879_v46);
        let _index139 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_877_v44).length));
        (_877_v44)[_index139] = !((_880_v47).dtor_cf7);
        let _index140 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_877_v44).length));
        (_877_v44)[_index140] = !(((_832_i0).minus(_829_v0)).isEqualTo(_832_i0));
      }
      let _881_v48;
      _881_v48 = _dafny.Seq.of((_this).f25);
      let _882_v49;
      _882_v49 = _dafny.MultiSet.fromElements(false);
      let _883_v50;
      _883_v50 = _dafny.Seq.of(_829_v0, _829_v0, (((_882_v49).contains(_this.f24)) ? ((_882_v49).get(_this.f24)) : (_829_v0)));
      r0 = (_881_v48)[_module.__default.safeIndex((_883_v50)[_module.__default.safeIndex(_829_v0, new BigNumber((_883_v50).length))], new BigNumber((_881_v48).length))];
      let _884_v51;
      _884_v51 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_829_v0);
      let _885_v52;
      _885_v52 = _dafny.Map.Empty.slice().updateUnsafe(_829_v0,_dafny.Map.Empty.slice().updateUnsafe(_829_v0,_829_v0));
      let _886_v53;
      let _nw161 = Array((new BigNumber(18)).toNumber());
      _nw161[(_dafny.ZERO).toNumber()] = (_829_v0).minus((_dafny.ZERO).minus(_829_v0));
      _nw161[(_dafny.ONE).toNumber()] = new BigNumber((_884_v51).length);
      _nw161[(new BigNumber(2)).toNumber()] = _829_v0;
      _nw161[(new BigNumber(3)).toNumber()] = _829_v0;
      _nw161[(new BigNumber(4)).toNumber()] = _829_v0;
      _nw161[(new BigNumber(5)).toNumber()] = new BigNumber((_885_v52).length);
      _nw161[(new BigNumber(6)).toNumber()] = _829_v0;
      _nw161[(new BigNumber(7)).toNumber()] = _829_v0;
      _nw161[(new BigNumber(8)).toNumber()] = _829_v0;
      _nw161[(new BigNumber(9)).toNumber()] = new BigNumber((_831_v2).length);
      _nw161[(new BigNumber(10)).toNumber()] = _829_v0;
      _nw161[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_829_v0);
      _nw161[(new BigNumber(12)).toNumber()] = _829_v0;
      _nw161[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt(_829_v0, _829_v0);
      _nw161[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_829_v0);
      _nw161[(new BigNumber(15)).toNumber()] = _829_v0;
      _nw161[(new BigNumber(16)).toNumber()] = new BigNumber((_883_v50).length);
      _nw161[(new BigNumber(17)).toNumber()] = ((_this.f24) ? (_829_v0) : (_829_v0));
      _886_v53 = _nw161;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_886_v53).length))) {
        let _887_i2 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_887_i2)) && ((_887_i2).isLessThan(new BigNumber((_886_v53).length))))) {
          (_886_v53)[(_887_i2)] = (_887_i2).plus((((_this).f25) ? (_829_v0) : (_829_v0)));
        }
      }
      (globalState).f5 = _829_v0;
      let _888_v54;
      _888_v54 = _dafny.Map.Empty.slice().updateUnsafe(_884_v51,(_this).f25);
      let _889_v57;
      _889_v57 = _dafny.Map.Empty.slice().updateUnsafe(_829_v0,_829_v0);
      let _890_v58;
      let _nw162 = Array((new BigNumber(28)).toNumber());
      _nw162[(_dafny.ZERO).toNumber()] = (_this).f25;
      _nw162[(_dafny.ONE).toNumber()] = _this.f24;
      _nw162[(new BigNumber(2)).toNumber()] = ((_881_v48)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_883_v50)[_module.__default.safeIndex(_829_v0, new BigNumber((_883_v50).length))])), new BigNumber((_881_v48).length))]) === (!(_this.f24));
      _nw162[(new BigNumber(3)).toNumber()] = _this.f24;
      _nw162[(new BigNumber(4)).toNumber()] = (_888_v54).equals(_888_v54);
      _nw162[(new BigNumber(5)).toNumber()] = ((_this).f25) || (true);
      _nw162[(new BigNumber(6)).toNumber()] = _this.f24;
      _nw162[(new BigNumber(7)).toNumber()] = (_this).f25;
      _nw162[(new BigNumber(8)).toNumber()] = _this.f24;
      _nw162[(new BigNumber(9)).toNumber()] = (_this).f25;
      _nw162[(new BigNumber(10)).toNumber()] = true;
      _nw162[(new BigNumber(11)).toNumber()] = (_this).f25;
      _nw162[(new BigNumber(12)).toNumber()] = (_this).f25;
      _nw162[(new BigNumber(13)).toNumber()] = (_this).f25;
      _nw162[(new BigNumber(14)).toNumber()] = _module.__default.fm29(new BigNumber((function () {
        let _coll31 = new _dafny.Map();
        for (const _compr_31 of _dafny.IntegerRange(new BigNumber(322), new BigNumber(256))) {
          let _891_v55 = _compr_31;
          if (((new BigNumber(322)).isLessThanOrEqualTo(_891_v55)) && ((_891_v55).isLessThan(new BigNumber(256)))) {
            _coll31.push([(_891_v55).plus(_829_v0),_this.f24]);
          }
        }
        return _coll31;
      }()).length), new BigNumber((function () {
        let _coll32 = new _dafny.Map();
        for (const _compr_32 of (_dafny.Seq.update(_883_v50, _module.__default.safeIndex(new BigNumber(674), new BigNumber((_883_v50).length)), new BigNumber((_882_v49).cardinality()))).Elements) {
          let _892_v56 = _compr_32;
          if (_dafny.Seq.contains(_dafny.Seq.update(_883_v50, _module.__default.safeIndex(new BigNumber(674), new BigNumber((_883_v50).length)), new BigNumber((_882_v49).cardinality())), _892_v56)) {
            _coll32.push([(_892_v56).minus(_829_v0),(_this).f25]);
          }
        }
        return _coll32;
      }()).length), (_this).f25, globalState);
      _nw162[(new BigNumber(15)).toNumber()] = _this.f24;
      _nw162[(new BigNumber(16)).toNumber()] = _this.f24;
      _nw162[(new BigNumber(17)).toNumber()] = !((((_889_v57).contains(_829_v0)) ? ((_889_v57).get(_829_v0)) : ((_dafny.ZERO).minus(_829_v0)))).isEqualTo((_dafny.ZERO).minus(_829_v0));
      _nw162[(new BigNumber(18)).toNumber()] = (_882_v49).IsProperSubsetOf(_882_v49);
      _nw162[(new BigNumber(19)).toNumber()] = _this.f24;
      _nw162[(new BigNumber(20)).toNumber()] = !((_this).f25);
      _nw162[(new BigNumber(21)).toNumber()] = (_this).f25;
      _nw162[(new BigNumber(22)).toNumber()] = _module.__default.fm29(_829_v0, _829_v0, _this.f24, globalState);
      _nw162[(new BigNumber(23)).toNumber()] = !(_this.f24);
      _nw162[(new BigNumber(24)).toNumber()] = !((_829_v0).isLessThanOrEqualTo(_829_v0));
      _nw162[(new BigNumber(25)).toNumber()] = (new BigNumber((_883_v50).length)).isLessThan(_829_v0);
      _nw162[(new BigNumber(26)).toNumber()] = false;
      _nw162[(new BigNumber(27)).toNumber()] = (_this).f25;
      _890_v58 = _nw162;
      let _index141 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_890_v58).length));
      (_890_v58)[_index141] = true;
      let _index142 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_890_v58).length));
      (_890_v58)[_index142] = _this.f24;
      let _893_v59;
      let _nw163 = new _module.C1();
      _nw163.__ctor(_this.f24, true);
      _893_v59 = _nw163;
      let _894_v60;
      _894_v60 = _dafny.Map.Empty.slice().updateUnsafe(_829_v0,_893_v59);
      let _rhs165 = _dafny.MultiSet.fromElements((_890_v58)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_890_v58).length))], (_this).f25, (_881_v48)[_module.__default.safeIndex(_829_v0, new BigNumber((_881_v48).length))]);
      let _rhs166 = _module.__default.fm29(_829_v0, _829_v0, (_829_v0).isEqualTo(new BigNumber((_894_v60).length)), globalState);
      let _lhs146 = globalState;
      _882_v49 = _rhs165;
      _lhs146.f9 = _rhs166;
      r0 = ((_dafny.ZERO).minus(_829_v0)).isLessThanOrEqualTo(_829_v0);
      let _nw164 = new _module.C3();
      _nw164.__ctor(_this.f24, _829_v0, (_890_v58)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_890_v58).length))], (_890_v58)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_890_v58).length))], (_dafny.ZERO).minus((_dafny.ZERO).minus(_829_v0)));
      r1 = _nw164;
      r2 = _this.f24;
      r3 = _829_v0;
      return [r0, r1, r2, r3];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f24 = false;
      this._f23 = _dafny.ZERO;
      this._f25 = false;
      this.f26 = _module.D0.Default();
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    set f24(value) {
      let _this = this;
      _this._f24 = value;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f26, f24, f25, f23) {
      let _this = this;
      (_this).f26 = f26;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      (_this)._f23 = f23;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_this.f24)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_dafny.Map.Empty.slice().updateUnsafe(false,true))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_this.f24))));
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return (_this).f23;
    };
    fm2(globalState) {
      let _this = this;
      return (_this).f23;
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f25;
    };
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = undefined;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _895_v0;
      _895_v0 = _dafny.Seq.UnicodeFromString("koa");
      let _896_v1;
      _896_v1 = _dafny.Map.Empty.slice().updateUnsafe(_895_v0,_this.f24);
      _896_v1 = (_896_v1).update(_dafny.Seq.UnicodeFromString("xx"), !(_this.f24));
      let _897_v2;
      let _nw165 = Array((new BigNumber(13)).toNumber()).fill([]);
      _897_v2 = _nw165;
      let _898_v3;
      _898_v3 = new _dafny.CodePoint('r'.codePointAt(0));
      let _899_v4;
      let _out9;
      _out9 = (_this).m4(_897_v2, (_this).f23, !(!(_this.f24)), _898_v3, globalState);
      _899_v4 = _out9;
      let _900_v5;
      _900_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f23);
      let _901_i0;
      _901_i0 = _dafny.ZERO;
      L7: {
        while (!((((_module.D1.create_DC5(false, _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25), (_this).f23, (_this).f23, _900_v5)).dtor_cf7) ? (((_this).f23).isEqualTo(new BigNumber(-781))) : ((_this).f25)))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_901_i0)) {
              break L7;
            }
            _901_i0 = (_901_i0).plus(_dafny.ONE);
            r2 = (_this).f25;
            let _902_v6;
            _902_v6 = _dafny.Seq.of((_this).f25);
            let _903_v7;
            let _init27 = function (_904_i1) {
              return (_904_i1).minus((_this).f23);
            };
            let _nw166 = Array((new BigNumber(11)).toNumber());
            for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw166.length); _i0_27++) {
              _nw166[_i0_27] = _init27(new BigNumber(_i0_27));
            }
            _903_v7 = _nw166;
            let _905_v8;
            _905_v8 = _dafny.Set.fromElements(_903_v7);
            let _906_v9;
            _906_v9 = _dafny.Seq.of((_this).f23);
            let _907_v10;
            _907_v10 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,!((_this).f25));
            let _908_v11;
            let _nw167 = Array((new BigNumber(26)).toNumber());
            _nw167[(_dafny.ZERO).toNumber()] = (_this).f23;
            _nw167[(_dafny.ONE).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(2)).toNumber()] = (((_this).f25) ? ((_this).f23) : (new BigNumber((_902_v6).length)));
            _nw167[(new BigNumber(3)).toNumber()] = new BigNumber((_902_v6).length);
            _nw167[(new BigNumber(4)).toNumber()] = new BigNumber((_895_v0).length);
            _nw167[(new BigNumber(5)).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(6)).toNumber()] = new BigNumber(142);
            _nw167[(new BigNumber(7)).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(8)).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_this).f23);
            _nw167[(new BigNumber(10)).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(11)).toNumber()] = ((_this).f23).minus((_this).f23);
            _nw167[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.of(_module.__default.fm11((_this).f23, _this.f24, _this.f24, new _dafny.CodePoint('a'.codePointAt(0)), globalState))).length);
            _nw167[(new BigNumber(13)).toNumber()] = new BigNumber(89);
            _nw167[(new BigNumber(14)).toNumber()] = ((_this.f24) ? ((_this).f23) : ((_this).f23));
            _nw167[(new BigNumber(15)).toNumber()] = new BigNumber(-763);
            _nw167[(new BigNumber(16)).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(17)).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(18)).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(((_this).f23).plus(new BigNumber((_905_v8).length)));
            _nw167[(new BigNumber(20)).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus((_this).f23);
            _nw167[(new BigNumber(22)).toNumber()] = (_this).f23;
            _nw167[(new BigNumber(23)).toNumber()] = (new BigNumber((_906_v9).length)).plus((_this).f23);
            _nw167[(new BigNumber(24)).toNumber()] = _module.__default.fm1((_module.D1.create_DC5(!(false), _907_v10, (_this).f23, (_this).f23, _900_v5)).dtor_cf10, (_this).f25, globalState);
            _nw167[(new BigNumber(25)).toNumber()] = (_this).fm2(globalState);
            _908_v11 = _nw167;
            let _909_v12;
            _909_v12 = _dafny.MultiSet.fromElements(_908_v11, _908_v11, _903_v7, _908_v11);
            let _index143 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_903_v7).length));
            (_903_v7)[_index143] = (((_909_v12).contains(_908_v11)) ? ((_909_v12).get(_908_v11)) : ((_this).f23));
            let _910_v13;
            _910_v13 = _module.D1.create_DC5(_this.f24, _907_v10, new BigNumber((_906_v9).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f25)).length), _900_v5);
            let _index144 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_903_v7).length));
            (_903_v7)[_index144] = (_910_v13).dtor_cf9;
            _898_v3 = _module.__default.fm12((_this).f25, globalState);
            let _911_v14;
            let _nw168 = new _module.C5();
            _nw168.__ctor(_this.f24, (_this).f25);
            _911_v14 = _nw168;
            let _912_v15;
            let _out10;
            _out10 = (_this).m4(_897_v2, (_this).f23, (_this).fm3(_895_v0, _this.f24, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_911_v14,_907_v10)).length), _dafny.Seq.Create(_module.__default.abs(new BigNumber(930)), function (_913_i2) {
              return (_this).f23;
            }), globalState), new _dafny.CodePoint('a'.codePointAt(0)), globalState);
            _912_v15 = _out10;
          }
        }
      }
      let _index145 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length));
      (_899_v4)[_index145] = _this.f24;
      let _index146 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length));
      (_899_v4)[_index146] = false;
      let _914_v16;
      _914_v16 = _dafny.Seq.of((_this).f23);
      let _915_v18;
      _915_v18 = _dafny.Set.fromElements((_this).f23, (_dafny.ZERO).minus((_this).f23), (_this).f23, (_this).f23, (_this).f23);
      let _916_v19;
      _916_v19 = _module.D0.create_DC0((_this).f25);
      let _917_v20;
      _917_v20 = _module.D0.create_DC2(_module.D0.create_DC2(_916_v19));
      let _pat_let_tv14 = _917_v20;
      let _918_v21;
      let _nw169 = Array((new BigNumber(21)).toNumber());
      _nw169[(_dafny.ZERO).toNumber()] = _this.f26;
      _nw169[(_dafny.ONE).toNumber()] = _this.f26;
      _nw169[(new BigNumber(2)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(3)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(4)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(5)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(6)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(7)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(8)).toNumber()] = _module.D0.create_DC2(_916_v19);
      _nw169[(new BigNumber(9)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(10)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(11)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(12)).toNumber()] = function (_pat_let9_0) {
        return function (_919_dt__update__tmp_h0) {
          return function (_pat_let10_0) {
            return function (_920_dt__update_hcf3_h0) {
              return _module.D0.create_DC2(_920_dt__update_hcf3_h0);
            }(_pat_let10_0);
          }(_pat_let_tv14);
        }(_pat_let9_0);
      }(_this.f26);
      _nw169[(new BigNumber(13)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(14)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(15)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(16)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(17)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(18)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(19)).toNumber()] = _this.f26;
      _nw169[(new BigNumber(20)).toNumber()] = _this.f26;
      _918_v21 = _nw169;
      let _921_v22;
      let _init28 = function (_922_i3) {
        return _module.__default.safeModuloInt(_922_i3, (_this).f23);
      };
      let _nw170 = Array((new BigNumber(8)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw170.length); _i0_28++) {
        _nw170[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _921_v22 = _nw170;
      let _923_v23;
      _923_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC9((_this).f23),_921_v22);
      let _924_v24;
      let _nw171 = Array((new BigNumber(13)).toNumber());
      _nw171[(_dafny.ZERO).toNumber()] = new BigNumber((_914_v16).length);
      _nw171[(_dafny.ONE).toNumber()] = ((_this).f23).plus(new BigNumber(968));
      _nw171[(new BigNumber(2)).toNumber()] = _module.__default.fm1((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of (_915_v18).Elements) {
          let _925_v17 = _compr_33;
          if ((_915_v18).contains(_925_v17)) {
            _coll33.push([_module.__default.safeModuloInt(_925_v17, new BigNumber(529)),_this.f24]);
          }
        }
        return _coll33;
      }()).length)), (_this).f25, globalState);
      _nw171[(new BigNumber(3)).toNumber()] = (_this).f23;
      _nw171[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(((_this).f23).plus((_this).f23));
      _nw171[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt((_this).f23, (_module.D5.create_DC15((_this).f23, _898_v3, _dafny.Map.Empty.slice().updateUnsafe(_918_v21,_898_v3), _923_v23)).dtor_cf25);
      _nw171[(new BigNumber(6)).toNumber()] = (_this).f23;
      _nw171[(new BigNumber(7)).toNumber()] = (_this).f23;
      _nw171[(new BigNumber(8)).toNumber()] = (_this).f23;
      _nw171[(new BigNumber(9)).toNumber()] = (_this).f23;
      _nw171[(new BigNumber(10)).toNumber()] = ((_this).f23).minus((_this).f23);
      _nw171[(new BigNumber(11)).toNumber()] = (_this).f23;
      _nw171[(new BigNumber(12)).toNumber()] = (_this).f23;
      _924_v24 = _nw171;
      let _926_v25;
      _926_v25 = _dafny.Map.Empty.slice().updateUnsafe(_924_v24,_this.f24);
      let _index147 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length));
      let _rhs167 = (_899_v4)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length))];
      let _rhs168 = ((!(new BigNumber(140)).isEqualTo(new BigNumber(926))) ? (_924_v24) : (_921_v22));
      let _rhs169 = _926_v25;
      let _lhs147 = _899_v4;
      let _lhs148 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length));
      _lhs147[_lhs148] = _rhs167;
      _921_v22 = _rhs168;
      _926_v25 = _rhs169;
      let _hi8 = (_this).f23;
      for (let _927_i4 = new BigNumber(825); _927_i4.isLessThan(_hi8); _927_i4 = _927_i4.plus(_dafny.ONE)) {
        _900_v5 = ((_module.__default.fm36(globalState)).update((_this).f25, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_this.f24)).length))).update((_this).f25, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f25),new BigNumber(285))).Merge(_900_v5)).length));
        let _928_v26;
        let _nw172 = new _module.C4();
        _nw172.__ctor(!((_this).f25), _this.f24, !((_this).f25));
        _928_v26 = _nw172;
        r3 = (_this).f23;
        let _929_v27;
        let _nw173 = new _module.C4();
        _nw173.__ctor(((_this).f25) && (!((_this).f25)), (_this).f25, !((_899_v4)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length))]) || (_this.f24));
        _929_v27 = _nw173;
        let _930_v28;
        _930_v28 = _dafny.Seq.of((_929_v27).f27, (_928_v26).f27, (_899_v4)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length))]);
        let _rhs170 = _this.f24;
        let _rhs171 = (_930_v28)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_927_i4, (_this).f23), new BigNumber((_930_v28).length))];
        let _rhs172 = _928_v26;
        let _rhs173 = (_929_v27).f27;
        let _rhs174 = _899_v4;
        let _lhs149 = globalState;
        let _lhs150 = _this;
        r0 = _rhs170;
        _lhs149.f6 = _rhs171;
        _929_v27 = _rhs172;
        _lhs150.f24 = _rhs173;
        _899_v4 = _rhs174;
      }
      let _931_v29;
      _931_v29 = _dafny.Seq.of((_899_v4)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length))], (_this).f25, _this.f24);
      r0 = _dafny.Seq.contains(_931_v29, (_899_v4)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length))]);
      let _932_v30;
      let _nw174 = new _module.C2();
      _nw174.__ctor((_this).f23);
      _932_v30 = _nw174;
      r1 = _932_v30;
      r2 = (_this).f25;
      let _933_v31;
      _933_v31 = _dafny.Map.Empty.slice().updateUnsafe(false,(_899_v4)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_899_v4).length))]);
      let _934_v32;
      _934_v32 = _module.D1.create_DC5(_this.f24, _933_v31, (_this).f23, (_this).f23, _900_v5);
      let _935_v33;
      _935_v33 = _dafny.Map.Empty.slice().updateUnsafe((_932_v30).f23,(_934_v32).dtor_cf9);
      let _936_v34;
      _936_v34 = _dafny.Map.Empty.slice().updateUnsafe((_932_v30).f23,new BigNumber((_935_v33).length));
      r3 = _module.__default.safeDivisionInt((_932_v30).f23, new BigNumber(((_936_v34).Merge(_dafny.Map.Empty.slice().updateUnsafe((_932_v30).f23,(_this).f23))).length));
      return [r0, r1, r2, r3];
    }
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = false;
      let _937_v0;
      _937_v0 = _dafny.Seq.of((_this).f23);
      let _938_v1;
      _938_v1 = _module.D10.create_DC28(_dafny.Seq.update(_dafny.Seq.of((_this).f23, (_this).f23), _module.__default.safeIndex((_this).f23, new BigNumber((_dafny.Seq.of((_this).f23, (_this).f23)).length)), new BigNumber(952)));
      let _939_v2;
      _939_v2 = _dafny.Seq.UnicodeFromString("itfta");
      let _940_v3;
      let _nw175 = Array((new BigNumber(18)).toNumber());
      _nw175[(_dafny.ZERO).toNumber()] = _937_v0;
      _nw175[(_dafny.ONE).toNumber()] = _937_v0;
      _nw175[(new BigNumber(2)).toNumber()] = _937_v0;
      _nw175[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(new BigNumber(-698));
      _nw175[(new BigNumber(4)).toNumber()] = _937_v0;
      _nw175[(new BigNumber(5)).toNumber()] = _937_v0;
      _nw175[(new BigNumber(6)).toNumber()] = (_938_v1).dtor_cf52;
      _nw175[(new BigNumber(7)).toNumber()] = _937_v0;
      _nw175[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_937_v0, _module.__default.safeIndex(new BigNumber((_939_v2).length), new BigNumber((_937_v0).length)), (_this).f23);
      _nw175[(new BigNumber(9)).toNumber()] = _module.__default.fm31(p1, true, globalState);
      _nw175[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(180)), function (_941_i1) {
        return (_this).f23;
      });
      _nw175[(new BigNumber(11)).toNumber()] = _937_v0;
      _nw175[(new BigNumber(12)).toNumber()] = _937_v0;
      _nw175[(new BigNumber(13)).toNumber()] = _937_v0;
      _nw175[(new BigNumber(14)).toNumber()] = _937_v0;
      _nw175[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_937_v0, _937_v0);
      _nw175[(new BigNumber(16)).toNumber()] = _937_v0;
      _nw175[(new BigNumber(17)).toNumber()] = _dafny.Seq.of((_this).f23);
      _940_v3 = _nw175;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_940_v3).length))) {
        let _942_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_942_i0)) && ((_942_i0).isLessThan(new BigNumber((_940_v3).length))))) {
          (_940_v3)[(_942_i0)] = _dafny.Seq.update(_dafny.Seq.Concat(_937_v0, _dafny.Seq.update(_937_v0, _module.__default.safeIndex((_this).f23, new BigNumber((_937_v0).length)), (_this).f23)), _module.__default.safeIndex((_this).f23, new BigNumber((_dafny.Seq.Concat(_937_v0, _dafny.Seq.update(_937_v0, _module.__default.safeIndex((_this).f23, new BigNumber((_937_v0).length)), (_this).f23))).length)), (_this).f23);
        }
      }
      let _943_v4;
      let _init29 = function (_944_i2) {
        return true;
      };
      let _nw176 = Array((new BigNumber(8)).toNumber());
      for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw176.length); _i0_29++) {
        _nw176[_i0_29] = _init29(new BigNumber(_i0_29));
      }
      _943_v4 = _nw176;
      _943_v4 = _943_v4;
      (globalState).f15 = ((_this).f23).isLessThanOrEqualTo(((_this).f23).plus((_this).f23));
      _937_v0 = _937_v0;
      let _945_v5;
      _945_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,p1);
      let _946_v6;
      _946_v6 = _dafny.Set.fromElements((_this).f25);
      let _947_v7;
      _947_v7 = _dafny.Seq.of(_946_v6);
      let _948_v8;
      _948_v8 = _dafny.Seq.of(_939_v2);
      let _949_v9;
      _949_v9 = _dafny.Seq.of((_this).f25);
      let _950_v10;
      let _init30 = function (_951_i4) {
        return _module.__default.safeModuloInt(_951_i4, new BigNumber((_dafny.Seq.UnicodeFromString("sl")).length));
      };
      let _nw177 = Array((new BigNumber(22)).toNumber());
      for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw177.length); _i0_30++) {
        _nw177[_i0_30] = _init30(new BigNumber(_i0_30));
      }
      _950_v10 = _nw177;
      let _952_v11;
      _952_v11 = _dafny.MultiSet.fromElements(_950_v10);
      let _953_v12;
      _953_v12 = _dafny.MultiSet.fromElements(_dafny.Seq.of(false, _this.f24, (_this).f25, _this.f24, (_this).f25));
      let _954_v13;
      _954_v13 = _dafny.MultiSet.fromElements(_this.f24);
      let _955_v14;
      _955_v14 = _dafny.Set.fromElements(_954_v13);
      let _956_v15;
      let _nw178 = Array((new BigNumber(21)).toNumber());
      _nw178[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_937_v0)[_module.__default.safeIndex(new BigNumber((_939_v2).length), new BigNumber((_937_v0).length))]);
      _nw178[(_dafny.ONE).toNumber()] = (_this).f23;
      _nw178[(new BigNumber(2)).toNumber()] = (_this).f23;
      _nw178[(new BigNumber(3)).toNumber()] = (_this).f23;
      _nw178[(new BigNumber(4)).toNumber()] = (_this).f23;
      _nw178[(new BigNumber(5)).toNumber()] = (_this).f23;
      _nw178[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_937_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_945_v5)).length), new BigNumber((_937_v0).length))]);
      _nw178[(new BigNumber(7)).toNumber()] = (_this).f23;
      _nw178[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_947_v7).length));
      _nw178[(new BigNumber(9)).toNumber()] = (_this).f23;
      _nw178[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt((_this).f23, (_dafny.ZERO).minus(new BigNumber(((_948_v8)[_module.__default.safeIndex((_this).f23, new BigNumber((_948_v8).length))]).length)));
      _nw178[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt((_this).f23, (_this).f23);
      _nw178[(new BigNumber(12)).toNumber()] = _module.__default.fm1((_this).f23, (_this).f25, globalState);
      _nw178[(new BigNumber(13)).toNumber()] = new BigNumber(((((_949_v9)[_module.__default.safeIndex((_this).f23, new BigNumber((_949_v9).length))]) ? (_952_v11) : (_952_v11))).cardinality());
      _nw178[(new BigNumber(14)).toNumber()] = ((_this).f23).minus((_this).f23);
      _nw178[(new BigNumber(15)).toNumber()] = (_this).f23;
      _nw178[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_939_v2).length), new BigNumber(996));
      _nw178[(new BigNumber(17)).toNumber()] = (((_953_v12).contains(_949_v9)) ? ((_953_v12).get(_949_v9)) : ((_this).f23));
      _nw178[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt((_this).f23, (_this).f23);
      _nw178[(new BigNumber(19)).toNumber()] = new BigNumber((_955_v14).length);
      _nw178[(new BigNumber(20)).toNumber()] = (_this).f23;
      _956_v15 = _nw178;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_950_v10).length))) {
        let _957_i3 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_957_i3)) && ((_957_i3).isLessThan(new BigNumber((_950_v10).length))))) {
          (_950_v10)[(_957_i3)] = (_957_i3).multipliedBy(((_this).f23).plus(new BigNumber((_937_v0).length)));
        }
      }
      let _hi9 = (_937_v0)[_module.__default.safeIndex((_this).f23, new BigNumber((_937_v0).length))];
      for (let _958_i5 = (new BigNumber((_946_v6).length)).plus((_this).f23); _958_i5.isLessThan(_hi9); _958_i5 = _958_i5.plus(_dafny.ONE)) {
        (globalState).f14 = (_module.__default.safeDivisionInt(_958_i5, _958_i5)).plus((_this).f23);
        let _959_v16;
        _959_v16 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),p0);
        let _960_v17;
        _960_v17 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(757)), function (_961_i6) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length), new BigNumber((_module.__default.fm37((_this).f25, new BigNumber((_959_v16).length), globalState)).cardinality()), _958_i5, _958_i5, (_this).f23);
        let _962_v18;
        _962_v18 = _dafny.Set.fromElements((_960_v17).update((_this).f23, _module.__default.abs((_this).f23)));
        let _963_v19;
        _963_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements(false, !(p1)));
        let _964_v20;
        let _nw179 = Array((new BigNumber(20)).toNumber());
        _nw179[(_dafny.ZERO).toNumber()] = (_946_v6).Difference(_946_v6);
        _nw179[(_dafny.ONE).toNumber()] = _946_v6;
        _nw179[(new BigNumber(2)).toNumber()] = _946_v6;
        _nw179[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(p1, p1, (_this).f25, p1);
        _nw179[(new BigNumber(4)).toNumber()] = (_dafny.Set.fromElements(_this.f24)).Difference(_946_v6);
        _nw179[(new BigNumber(5)).toNumber()] = ((_this.f24) ? (_946_v6) : (_946_v6));
        _nw179[(new BigNumber(6)).toNumber()] = (_946_v6).Union(_946_v6);
        _nw179[(new BigNumber(7)).toNumber()] = (_946_v6).Union(_dafny.Set.fromElements(_this.f24));
        _nw179[(new BigNumber(8)).toNumber()] = _946_v6;
        _nw179[(new BigNumber(9)).toNumber()] = _946_v6;
        _nw179[(new BigNumber(10)).toNumber()] = _946_v6;
        _nw179[(new BigNumber(11)).toNumber()] = _946_v6;
        _nw179[(new BigNumber(12)).toNumber()] = _946_v6;
        _nw179[(new BigNumber(13)).toNumber()] = (_946_v6).Intersect(_946_v6);
        _nw179[(new BigNumber(14)).toNumber()] = _946_v6;
        _nw179[(new BigNumber(15)).toNumber()] = (_946_v6).Union(_946_v6);
        _nw179[(new BigNumber(16)).toNumber()] = (_dafny.Set.fromElements(p1, (_this).f25, _this.f24, _this.f24)).Intersect((((_963_v19).contains((_this).f25)) ? ((_963_v19).get((_this).f25)) : (_946_v6)));
        _nw179[(new BigNumber(17)).toNumber()] = _dafny.Set.fromElements(p0, p0, p0, _this.f24, (_this).fm3(_939_v2, (_this).f25, _958_i5, _937_v0, globalState));
        _nw179[(new BigNumber(18)).toNumber()] = _946_v6;
        _nw179[(new BigNumber(19)).toNumber()] = _946_v6;
        _964_v20 = _nw179;
        let _index148 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_964_v20).length));
        (_964_v20)[_index148] = _946_v6;
        let _index149 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_964_v20).length));
        let _rhs175 = _962_v18;
        let _rhs176 = _module.__default.safeDivisionInt(_958_i5, (_dafny.ZERO).minus(_module.__default.fm1((_dafny.ZERO).minus(_958_i5), p0, globalState)));
        let _rhs177 = _dafny.Set.fromElements(p1, (_958_i5).isLessThan(_958_i5), !(false));
        let _lhs151 = globalState;
        let _lhs152 = _964_v20;
        let _lhs153 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_964_v20).length));
        _962_v18 = _rhs175;
        _lhs151.f8 = _rhs176;
        _lhs152[_lhs153] = _rhs177;
        if (p0) {
          let _965_v21;
          _965_v21 = new _dafny.CodePoint('y'.codePointAt(0));
          let _966_v22;
          _966_v22 = _dafny.MultiSet.fromElements(_960_v17);
          let _967_v23;
          _967_v23 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.of(_965_v21)).length)).multipliedBy((_this).f23),_966_v22);
          _967_v23 = (_967_v23).update(((((_960_v17).contains(new BigNumber((_949_v9).length))) ? ((_960_v17).get(new BigNumber((_949_v9).length))) : ((_this).f23))).plus(new BigNumber(895)), _966_v22);
          let _968_v24;
          _968_v24 = _dafny.Set.fromElements(_943_v4);
          let _969_v25;
          let _nw180 = new _module.C0();
          _nw180.__ctor((_968_v24).Difference(_968_v24), (_this).f23, _958_i5);
          _969_v25 = _nw180;
          (globalState).f17 = new BigNumber((_dafny.Seq.of(_dafny.Seq.of(p0, _this.f24), _dafny.Seq.Concat(_949_v9, _949_v9), _949_v9)).length);
          let _970_v26;
          let _nw181 = Array((new BigNumber(22)).toNumber()).fill([]);
          _970_v26 = _nw181;
          let _index150 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_970_v26).length));
          (_970_v26)[_index150] = _950_v10;
          let _index151 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_970_v26).length));
          let _init31 = ((_971_v25) => function (_972_i7) {
            return (_972_i7).multipliedBy((_dafny.ZERO).minus(_971_v25.f31));
          })(_969_v25);
          let _nw182 = Array((new BigNumber(28)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw182.length); _i0_31++) {
            _nw182[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          (_970_v26)[_index151] = _nw182;
          _943_v4 = _943_v4;
        } else {
          _939_v2 = _dafny.Seq.Concat(_939_v2, _939_v2);
          let _973_v27;
          _973_v27 = _dafny.Set.fromElements(_946_v6, _946_v6);
          (globalState).f9 = (((p1) ? (_dafny.Set.fromElements((_964_v20)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_964_v20).length))], (_964_v20)[_module.__default.safeIndex(new BigNumber(234), new BigNumber((_964_v20).length))])) : (_module.__default.fm38((_this).f23, (_dafny.ZERO).minus(_958_i5), globalState)))).IsDisjointFrom(_973_v27);
          (globalState).f6 = (_949_v9)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_949_v9).length))];
          (globalState).f15 = (_this).fm3(_dafny.Seq.UnicodeFromString("gvwh"), p0, _module.__default.safeDivisionInt(new BigNumber((_937_v0).length), _958_i5), _937_v0, globalState);
          let _974_v28;
          _974_v28 = _dafny.Map.Empty.slice().updateUnsafe(_949_v9,_958_i5);
          let _975_v29;
          let _nw183 = new _module.C3();
          _nw183.__ctor(p0, new BigNumber((_974_v28).length), p0, (_958_i5).isLessThan(_958_i5), _958_i5);
          _975_v29 = _nw183;
        }
        (globalState).f14 = new BigNumber((_dafny.Seq.Concat(_949_v9, ((_this.f24) ? (_949_v9) : (_dafny.Seq.of(_this.f24, _this.f24))))).length);
      }
      r0 = (_this).f23;
      r1 = _956_v15;
      r2 = (_this).f25;
      return [r0, r1, r2];
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let _976_v1;
      _976_v1 = _dafny.Seq.of(p3, new _dafny.CodePoint('u'.codePointAt(0)));
      let _rhs178 = (_this).fm2(globalState);
      let _rhs179 = (new BigNumber((_dafny.Seq.UnicodeFromString("u")).length)).isLessThanOrEqualTo((new BigNumber((function () {
        let _coll34 = new _dafny.Map();
        for (const _compr_34 of _dafny.IntegerRange(new BigNumber(-489), new BigNumber(-153))) {
          let _977_v0 = _compr_34;
          if (((new BigNumber(-489)).isLessThanOrEqualTo(_977_v0)) && ((_977_v0).isLessThan(new BigNumber(-153)))) {
            _coll34.push([(_977_v0).plus(new BigNumber((_dafny.Seq.of((_this).f23, p1, p1)).length)),(_this).f25]);
          }
        }
        return _coll34;
      }()).length)).minus(new BigNumber((_dafny.Seq.update(_976_v1, _module.__default.safeIndex((_this).f23, new BigNumber((_976_v1).length)), new _dafny.CodePoint('q'.codePointAt(0)))).length)));
      let _lhs154 = globalState;
      let _lhs155 = globalState;
      _lhs154.f8 = _rhs178;
      _lhs155.f9 = _rhs179;
      let _978_v2;
      _978_v2 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("bhkyjg")).length), (_this).f23),(_this).f23);
      let _979_v3;
      _979_v3 = _dafny.MultiSet.fromElements((_this).f23, p1);
      let _980_v4;
      _980_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f23);
      let _981_v5;
      let _nw184 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _981_v5 = _nw184;
      let _982_v6;
      _982_v6 = _module.D7.create_DC21(_module.D7.create_DC20(_this.f24, _980_v4, _981_v5));
      let _983_v7;
      _983_v7 = _module.D7.create_DC21(_982_v6);
      let _984_v8;
      _984_v8 = _dafny.Seq.of(p1);
      let _985_v9;
      _985_v9 = _dafny.Map.Empty.slice().updateUnsafe(_983_v7,(_984_v8)[_module.__default.safeIndex(p1, new BigNumber((_984_v8).length))]);
      let _986_v10;
      _986_v10 = _dafny.Seq.of((((_985_v9).contains(_983_v7)) ? ((_985_v9).get(_983_v7)) : ((_this).f23)));
      (globalState).f1 = (_this).fm3(_dafny.Seq.Concat(_976_v1, _dafny.Seq.update(_976_v1, _module.__default.safeIndex((_this).fm5(p1, new BigNumber((_978_v2).length), globalState), new BigNumber((_976_v1).length)), p3)), (_this).f25, new BigNumber((_979_v3).cardinality()), _986_v10, globalState);
      if ((new BigNumber(754)).isEqualTo((_this).f23)) {
        let _987_v11;
        let _nw185 = new _module.C5();
        _nw185.__ctor((_this).f25, (_this).f25);
        _987_v11 = _nw185;
        _987_v11 = _987_v11;
        let _988_v12;
        _988_v12 = _dafny.Set.fromElements((_this).f25);
        let _source8 = _module.__default.fm39(p2, (p1).minus((_this).f23), true, (_dafny.Set.fromElements(p2, _this.f24, (_this).f25, p2, true)).Difference(_988_v12), globalState);
        if (_source8.is_DC9) {
          let _989___mcc_h0 = (_source8).cf15;
          let _990_cf15 = _989___mcc_h0;
          (globalState).f8 = (new BigNumber(((_dafny.MultiSet.fromElements((_this).f23, _990_cf15)).Intersect(_979_v3)).cardinality())).multipliedBy(_990_cf15);
          let _991_v13;
          _991_v13 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),(_dafny.ZERO).minus((_this).f23));
          let _992_v14;
          let _nw186 = Array((new BigNumber(28)).toNumber());
          _nw186[(_dafny.ZERO).toNumber()] = p3;
          _nw186[(_dafny.ONE).toNumber()] = p3;
          _nw186[(new BigNumber(2)).toNumber()] = p3;
          _nw186[(new BigNumber(3)).toNumber()] = p3;
          _nw186[(new BigNumber(4)).toNumber()] = p3;
          _nw186[(new BigNumber(5)).toNumber()] = p3;
          _nw186[(new BigNumber(6)).toNumber()] = p3;
          _nw186[(new BigNumber(7)).toNumber()] = p3;
          _nw186[(new BigNumber(8)).toNumber()] = p3;
          _nw186[(new BigNumber(9)).toNumber()] = p3;
          _nw186[(new BigNumber(10)).toNumber()] = p3;
          _nw186[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
          _nw186[(new BigNumber(12)).toNumber()] = p3;
          _nw186[(new BigNumber(13)).toNumber()] = p3;
          _nw186[(new BigNumber(14)).toNumber()] = p3;
          _nw186[(new BigNumber(15)).toNumber()] = p3;
          _nw186[(new BigNumber(16)).toNumber()] = p3;
          _nw186[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _nw186[(new BigNumber(18)).toNumber()] = p3;
          _nw186[(new BigNumber(19)).toNumber()] = p3;
          _nw186[(new BigNumber(20)).toNumber()] = p3;
          _nw186[(new BigNumber(21)).toNumber()] = p3;
          _nw186[(new BigNumber(22)).toNumber()] = p3;
          _nw186[(new BigNumber(23)).toNumber()] = p3;
          _nw186[(new BigNumber(24)).toNumber()] = p3;
          _nw186[(new BigNumber(25)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
          _nw186[(new BigNumber(26)).toNumber()] = p3;
          _nw186[(new BigNumber(27)).toNumber()] = p3;
          _992_v14 = _nw186;
          let _993_v15;
          _993_v15 = _dafny.Set.fromElements(_992_v14);
          let _994_v16;
          _994_v16 = _module.D5.create_DC14((_this).f23, p2, _993_v15);
          let _995_v17;
          _995_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p2);
          let _996_v18;
          let _nw187 = Array((new BigNumber(16)).toNumber());
          _nw187[(_dafny.ZERO).toNumber()] = _this.f24;
          _nw187[(_dafny.ONE).toNumber()] = (_994_v16).dtor_cf23;
          _nw187[(new BigNumber(2)).toNumber()] = (_this).f25;
          _nw187[(new BigNumber(3)).toNumber()] = (_this).fm3(_976_v1, _this.f24, new BigNumber((_980_v4).length), _984_v8, globalState);
          _nw187[(new BigNumber(4)).toNumber()] = false;
          _nw187[(new BigNumber(5)).toNumber()] = (((_995_v17).contains((_this).f25)) ? ((_995_v17).get((_this).f25)) : (_this.f24));
          _nw187[(new BigNumber(6)).toNumber()] = true;
          _nw187[(new BigNumber(7)).toNumber()] = p2;
          _nw187[(new BigNumber(8)).toNumber()] = (_this).f25;
          _nw187[(new BigNumber(9)).toNumber()] = p2;
          _nw187[(new BigNumber(10)).toNumber()] = _this.f24;
          _nw187[(new BigNumber(11)).toNumber()] = _this.f24;
          _nw187[(new BigNumber(12)).toNumber()] = (_this).fm3(_976_v1, _this.f24, (_this).f23, _986_v10, globalState);
          _nw187[(new BigNumber(13)).toNumber()] = _this.f24;
          _nw187[(new BigNumber(14)).toNumber()] = (_this).f25;
          _nw187[(new BigNumber(15)).toNumber()] = (_this).f25;
          _996_v18 = _nw187;
          let _997_v19;
          let _nw188 = Array((new BigNumber(26)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = (_this).f23;
          _nw188[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(p1)).cardinality());
          _nw188[(new BigNumber(2)).toNumber()] = _990_cf15;
          _nw188[(new BigNumber(3)).toNumber()] = _990_cf15;
          _nw188[(new BigNumber(4)).toNumber()] = p1;
          _nw188[(new BigNumber(5)).toNumber()] = new BigNumber(283);
          _nw188[(new BigNumber(6)).toNumber()] = (_this).f23;
          _nw188[(new BigNumber(7)).toNumber()] = p1;
          _nw188[(new BigNumber(8)).toNumber()] = (_this).f23;
          _nw188[(new BigNumber(9)).toNumber()] = _990_cf15;
          _nw188[(new BigNumber(10)).toNumber()] = (_this).f23;
          _nw188[(new BigNumber(11)).toNumber()] = new BigNumber(406);
          _nw188[(new BigNumber(12)).toNumber()] = (_this).f23;
          _nw188[(new BigNumber(13)).toNumber()] = p1;
          _nw188[(new BigNumber(14)).toNumber()] = new BigNumber((_991_v13).length);
          _nw188[(new BigNumber(15)).toNumber()] = _990_cf15;
          _nw188[(new BigNumber(16)).toNumber()] = p1;
          _nw188[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f24,p2)).length);
          _nw188[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("huhwbkh")).length);
          _nw188[(new BigNumber(19)).toNumber()] = new BigNumber((_976_v1).length);
          _nw188[(new BigNumber(20)).toNumber()] = _990_cf15;
          _nw188[(new BigNumber(21)).toNumber()] = new BigNumber((_986_v10).length);
          _nw188[(new BigNumber(22)).toNumber()] = (_this).f23;
          _nw188[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.of(_996_v18)).length);
          _nw188[(new BigNumber(24)).toNumber()] = (_this).f23;
          _nw188[(new BigNumber(25)).toNumber()] = p1;
          _997_v19 = _nw188;
          let _998_v20;
          _998_v20 = _dafny.MultiSet.fromElements(_997_v19);
          let _999_v21;
          _999_v21 = _dafny.Map.Empty.slice().updateUnsafe(_998_v20,(_987_v11).fm5((_this).f23, _990_cf15, globalState));
          _999_v21 = (_999_v21).update((_module.D12.create_DC32(_998_v20)).dtor_cf60, ((_this).f23).plus(_990_cf15));
          (globalState).f9 = true;
          (globalState).f1 = !(_990_cf15).isEqualTo((_this).f23);
        } else if (_source8.is_DC10) {
          let _1000___mcc_h1 = (_source8).cf16;
          let _1001___mcc_h2 = (_source8).cf17;
          let _1002_cf17 = _1001___mcc_h2;
          let _1003_cf16 = _1000___mcc_h1;
          (globalState).f8 = _module.__default.safeDivisionInt((_this).f23, (_this).f23);
          (globalState).f14 = p1;
          let _rhs180 = p1;
          let _rhs181 = (_this).fm2(globalState);
          let _rhs182 = (false) || (_this.f24);
          let _rhs183 = p1;
          let _rhs184 = p2;
          let _lhs156 = globalState;
          let _lhs157 = globalState;
          let _lhs158 = _this;
          let _lhs159 = globalState;
          let _lhs160 = globalState;
          _lhs156.f5 = _rhs180;
          _lhs157.f5 = _rhs181;
          _lhs158.f24 = _rhs182;
          _lhs159.f17 = _rhs183;
          _lhs160.f15 = _rhs184;
          let _1004_v22;
          let _nw189 = new _module.C1();
          _nw189.__ctor(p2, false);
          _1004_v22 = _nw189;
        } else {
          let _1005___mcc_h3 = (_source8).cf14;
          let _1006_cf14 = _1005___mcc_h3;
          let _1007_v23;
          let _init32 = function (_1008_i0) {
            return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f24))).Intersect(_dafny.MultiSet.fromElements((_this).f25));
          };
          let _nw190 = Array((new BigNumber(24)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw190.length); _i0_32++) {
            _nw190[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1007_v23 = _nw190;
          let _1009_v24;
          _1009_v24 = _dafny.MultiSet.fromElements(p2);
          let _index152 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_1007_v23).length));
          (_1007_v23)[_index152] = _1009_v24;
          let _1010_v25;
          _1010_v25 = _module.D7.create_DC19(_1009_v24);
          let _index153 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_1007_v23).length));
          (_1007_v23)[_index153] = (_1010_v25).dtor_cf34;
          (globalState).f8 = new BigNumber((_dafny.MultiSet.fromElements(p2, p2)).cardinality());
          let _1011_v26;
          _1011_v26 = _dafny.Seq.of(_981_v5, _981_v5, _981_v5, _981_v5);
          _981_v5 = (_1011_v26)[_module.__default.safeIndex((_this).f23, new BigNumber((_1011_v26).length))];
          (globalState).f17 = p1;
        }
        let _1012_v27;
        _1012_v27 = _dafny.Seq.of(_979_v3);
        let _rhs185 = ((_1012_v27)[_module.__default.safeIndex((_this).f23, new BigNumber((_1012_v27).length))]).Difference(_979_v3);
        let _rhs186 = p2;
        let _rhs187 = p2;
        let _lhs161 = globalState;
        let _lhs162 = globalState;
        _979_v3 = _rhs185;
        _lhs161.f1 = _rhs186;
        _lhs162.f1 = _rhs187;
        (globalState).f17 = (_this).f23;
        let _1013_v28;
        _1013_v28 = new _dafny.CodePoint('m'.codePointAt(0));
        _1013_v28 = _1013_v28;
      } else {
        (globalState).f15 = p2;
        (globalState).f17 = (_dafny.ZERO).minus(p1);
        (globalState).f8 = (_this).f23;
        (globalState).f8 = new BigNumber(129);
        (globalState).f8 = new BigNumber(((_dafny.MultiSet.fromElements((_this).f25, (_this).f25, !(true))).update((_this).f25, _module.__default.abs(_module.__default.safeDivisionInt((_this).f23, (_this).f23)))).cardinality());
      }
      (globalState).f15 = _this.f24;
      (globalState).f9 = ((false) ? (_this.f24) : ((((_this).f25) ? ((_this).fm3(_976_v1, p2, (_this).f23, _984_v8, globalState)) : (p2))));
      let _1014_v29;
      let _nw191 = Array((new BigNumber(16)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1014_v29 = _nw191;
      let _1015_v30;
      _1015_v30 = _dafny.Set.fromElements(_1014_v29, _1014_v29);
      let _1016_v31;
      _1016_v31 = _module.D5.create_DC14((_this).f23, _this.f24, _1015_v30);
      let _1017_v32;
      _1017_v32 = _dafny.MultiSet.fromElements(_1016_v31, _1016_v31, _1016_v31, _1016_v31, _1016_v31);
      (globalState).f14 = (((_this).f25) ? (new BigNumber((_1017_v32).cardinality())) : (p1));
      let _1018_v33;
      let _nw192 = Array((new BigNumber(2)).toNumber());
      _nw192[(_dafny.ZERO).toNumber()] = _this.f24;
      _nw192[(_dafny.ONE).toNumber()] = _this.f24;
      _1018_v33 = _nw192;
      r0 = _1018_v33;
      return r0;
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm7(globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.FromArray(_dafny.Seq.of(false, !(true), false, !(false), !(true))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(false, false, true, true)));
    };
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _module.D0.Default();
      r0 = p1;
      let _1019_v0;
      _1019_v0 = _dafny.Seq.UnicodeFromString("hrd");
      let _1020_v1;
      _1020_v1 = _module.D0.create_DC1(p1, p0);
      let _pat_let_tv15 = _1020_v1;
      let _pat_let_tv16 = p0;
      let _pat_let_tv17 = p0;
      let _rhs188 = (new BigNumber((_1019_v0).length)).minus(p0);
      let _rhs189 = (_dafny.ZERO).minus(function (_source9) {
        if (_source9.is_DC1) {
          let _1021___mcc_h0 = (_source9).cf1;
          let _1022___mcc_h1 = (_source9).cf2;
          let _1023_cf2 = _1022___mcc_h1;
          let _1024_cf1 = _1021___mcc_h0;
          return (_1023_cf2).minus(_1023_cf2);
        } else if (_source9.is_DC0) {
          let _1025___mcc_h2 = (_source9).cf0;
          let _1026_cf0 = _1025___mcc_h2;
          return (_pat_let_tv15).dtor_cf2;
        } else {
          let _1027___mcc_h3 = (_source9).cf3;
          let _1028_cf3 = _1027___mcc_h3;
          return _module.__default.safeDivisionInt(_pat_let_tv16, _pat_let_tv17);
        }
      }(_1020_v1));
      let _lhs163 = globalState;
      let _lhs164 = globalState;
      _lhs163.f8 = _rhs188;
      _lhs164.f8 = _rhs189;
      let _1029_v2;
      let _nw193 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
      _1029_v2 = _nw193;
      let _1030_v3;
      let _nw194 = Array((new BigNumber(5)).toNumber()).fill(false);
      _1030_v3 = _nw194;
      let _1031_v4;
      _1031_v4 = _dafny.Seq.of(_1030_v3, _1030_v3, _1030_v3);
      let _index154 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_1029_v2).length));
      (_1029_v2)[_index154] = _dafny.Seq.Concat(_1031_v4, _1031_v4);
      let _index155 = _module.__default.safeIndex(new BigNumber(399), new BigNumber((_1029_v2).length));
      (_1029_v2)[_index155] = _dafny.Seq.Concat(_1031_v4, _1031_v4);
      let _1032_v5;
      _1032_v5 = _dafny.Seq.of(p0);
      (globalState).f8 = _module.__default.fm1(new BigNumber((_module.__default.fm8(new BigNumber((_1032_v5).length), globalState)).length), (new BigNumber((_1019_v0).length)).isEqualTo(new BigNumber(639)), globalState);
      let _1033_v7;
      let _init33 = ((_1034_v5, _1035_p1) => function (_1036_i0) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_1034_v5,_dafny.Seq.of(_1035_p1, _1035_p1))).Merge(function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of (_dafny.MultiSet.fromElements(_1034_v5, _1034_v5)).Elements) {
            let _1037_v6 = _compr_35;
            if ((_dafny.MultiSet.fromElements(_1034_v5, _1034_v5)).contains(_1037_v6)) {
              _coll35.push([_1037_v6,_dafny.Seq.of(_1035_p1, _1035_p1)]);
            }
          }
          return _coll35;
        }());
      })(_1032_v5, p1);
      let _nw195 = Array((new BigNumber(10)).toNumber());
      for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw195.length); _i0_33++) {
        _nw195[_i0_33] = _init33(new BigNumber(_i0_33));
      }
      _1033_v7 = _nw195;
      let _1038_v8;
      _1038_v8 = _dafny.Seq.of(p1);
      let _1039_v9;
      _1039_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0, p0, p0),_1038_v8);
      let _index156 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1033_v7).length));
      (_1033_v7)[_index156] = (_1039_v9).Merge(_1039_v9);
      let _index157 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1033_v7).length));
      (_1033_v7)[_index157] = _1039_v9;
      let _index158 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1030_v3).length));
      (_1030_v3)[_index158] = p1;
      let _index159 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1030_v3).length));
      (_1030_v3)[_index159] = ((false) ? (p1) : (!(_module.__default.fm9(p1, globalState)) || (_module.__default.fm9(true, globalState))));
      r0 = p1;
      r1 = _module.__default.fm9(p1, globalState);
      r2 = _1020_v1;
      return [r0, r1, r2];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = [];
      let _1040_v0;
      _1040_v0 = false;
      (globalState).f15 = _1040_v0;
      let _1041_v1;
      _1041_v1 = new _dafny.CodePoint('l'.codePointAt(0));
      let _1042_v2;
      _1042_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1040_v0,new _dafny.CodePoint('i'.codePointAt(0)));
      let _1043_v3;
      _1043_v3 = _module.D1.create_DC3(_1041_v1);
      let _1044_v4;
      let _nw196 = Array((new BigNumber(14)).toNumber());
      _nw196[(_dafny.ZERO).toNumber()] = _1041_v1;
      _nw196[(_dafny.ONE).toNumber()] = _1041_v1;
      _nw196[(new BigNumber(2)).toNumber()] = (((_1042_v2).contains(_1040_v0)) ? ((_1042_v2).get(_1040_v0)) : (_1041_v1));
      _nw196[(new BigNumber(3)).toNumber()] = _1041_v1;
      _nw196[(new BigNumber(4)).toNumber()] = _1041_v1;
      _nw196[(new BigNumber(5)).toNumber()] = _1041_v1;
      _nw196[(new BigNumber(6)).toNumber()] = _1041_v1;
      _nw196[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
      _nw196[(new BigNumber(8)).toNumber()] = (_1043_v3).dtor_cf4;
      _nw196[(new BigNumber(9)).toNumber()] = _1041_v1;
      _nw196[(new BigNumber(10)).toNumber()] = ((_1040_v0) ? (_1041_v1) : (_1041_v1));
      _nw196[(new BigNumber(11)).toNumber()] = (((_1042_v2).contains(_1040_v0)) ? ((_1042_v2).get(_1040_v0)) : (_1041_v1));
      _nw196[(new BigNumber(12)).toNumber()] = _1041_v1;
      _nw196[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
      _1044_v4 = _nw196;
      let _index160 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length));
      (_1044_v4)[_index160] = _1041_v1;
      let _1045_v5;
      _1045_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1040_v0,_1040_v0);
      let _pat_let_tv18 = _1040_v0;
      let _pat_let_tv19 = _1040_v0;
      let _pat_let_tv20 = _1040_v0;
      let _pat_let_tv21 = _1040_v0;
      let _index161 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length));
      let _rhs190 = _1041_v1;
      let _rhs191 = function (_source10) {
        if (_source10.is_DC4) {
          let _1046___mcc_h0 = (_source10).cf5;
          let _1047___mcc_h1 = (_source10).cf6;
          let _1048_cf6 = _1047___mcc_h1;
          let _1049_cf5 = _1046___mcc_h0;
          return _pat_let_tv18;
        } else if (_source10.is_DC5) {
          let _1050___mcc_h2 = (_source10).cf7;
          let _1051___mcc_h3 = (_source10).cf8;
          let _1052___mcc_h4 = (_source10).cf9;
          let _1053___mcc_h5 = (_source10).cf10;
          let _1054___mcc_h6 = (_source10).cf11;
          let _1055_cf11 = _1054___mcc_h6;
          let _1056_cf10 = _1053___mcc_h5;
          let _1057_cf9 = _1052___mcc_h4;
          let _1058_cf8 = _1051___mcc_h3;
          let _1059_cf7 = _1050___mcc_h2;
          return _pat_let_tv19;
        } else if (_source10.is_DC3) {
          let _1060___mcc_h7 = (_source10).cf4;
          let _1061_cf4 = _1060___mcc_h7;
          return _pat_let_tv20;
        } else {
          let _1062___mcc_h8 = (_source10).cf12;
          let _1063_cf12 = _1062___mcc_h8;
          return _pat_let_tv21;
        }
      }(_module.D1.create_DC3(_1041_v1));
      let _rhs192 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1040_v0,_1040_v0)).Merge(_1045_v5)).length);
      let _rhs193 = false;
      let _lhs165 = _1044_v4;
      let _lhs166 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length));
      let _lhs167 = globalState;
      let _lhs168 = globalState;
      let _lhs169 = globalState;
      _lhs165[_lhs166] = _rhs190;
      _lhs167.f15 = _rhs191;
      _lhs168.f8 = _rhs192;
      _lhs169.f1 = _rhs193;
      let _1064_i0;
      _1064_i0 = _dafny.ZERO;
      L8: {
        while (_1040_v0) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1064_i0)) {
              break L8;
            }
            _1064_i0 = (_1064_i0).plus(_dafny.ONE);
            let _1065_v6;
            _1065_v6 = _dafny.Seq.UnicodeFromString("asg");
            _1065_v6 = _1065_v6;
            (globalState).f8 = _module.__default.safeDivisionInt(p2, p0);
            let _1066_v7;
            _1066_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1040_v0,p0);
            let _1067_v8;
            _1067_v8 = _module.D1.create_DC5(_1040_v0, (_dafny.Map.Empty.slice().updateUnsafe(_1040_v0,_1040_v0)).Merge(_1045_v5), (p0).plus(p1), new BigNumber(357), _1066_v7);
            let _source11 = _1067_v8;
            if (_source11.is_DC4) {
              let _1068___mcc_h9 = (_source11).cf5;
              let _1069___mcc_h10 = (_source11).cf6;
              let _1070_cf6 = _1069___mcc_h10;
              let _1071_cf5 = _1068___mcc_h9;
              let _1072_v9;
              _1072_v9 = _dafny.Seq.of(_1071_cf5);
              let _1073_v10;
              _1073_v10 = _dafny.Seq.of(_module.__default.fm1(p0, _1040_v0, globalState), p0);
              let _1074_v11;
              _1074_v11 = _dafny.Seq.of(new BigNumber((_1073_v10).length), p1, new BigNumber((_1072_v9).length), p1);
              let _1075_v12;
              _1075_v12 = _dafny.Seq.of((_1072_v9)[_module.__default.safeIndex(new BigNumber((_1074_v11).length), new BigNumber((_1072_v9).length))], _1071_cf5);
              _1075_v12 = _1075_v12;
              _1041_v1 = ((!((_1071_cf5) && (_1040_v0))) ? (_1041_v1) : (new _dafny.CodePoint('b'.codePointAt(0))));
              let _1076_v13;
              _1076_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
              _1076_v13 = (_1076_v13).update(new BigNumber((_1065_v6).length), p1);
              (globalState).f14 = ((_1070_cf6) ? (new BigNumber((_module.__default.fm10(_1040_v0, new BigNumber((_1075_v12).length), _1041_v1, (_1044_v4)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length))], globalState)).length)) : (p1));
            } else if (_source11.is_DC5) {
              let _1077___mcc_h11 = (_source11).cf7;
              let _1078___mcc_h12 = (_source11).cf8;
              let _1079___mcc_h13 = (_source11).cf9;
              let _1080___mcc_h14 = (_source11).cf10;
              let _1081___mcc_h15 = (_source11).cf11;
              let _1082_cf11 = _1081___mcc_h15;
              let _1083_cf10 = _1080___mcc_h14;
              let _1084_cf9 = _1079___mcc_h13;
              let _1085_cf8 = _1078___mcc_h12;
              let _1086_cf7 = _1077___mcc_h11;
              let _1087_v14;
              let _init34 = ((_1088_v0) => function (_1089_i1) {
                return _1088_v0;
              })(_1040_v0);
              let _nw197 = Array((new BigNumber(9)).toNumber());
              for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw197.length); _i0_34++) {
                _nw197[_i0_34] = _init34(new BigNumber(_i0_34));
              }
              _1087_v14 = _nw197;
              let _1090_v15;
              _1090_v15 = _dafny.Set.fromElements(_1087_v14);
              let _1091_v16;
              _1091_v16 = _1090_v15;
              let _1092_v17;
              let _nw198 = new _module.C0();
              _nw198.__ctor((_1090_v15).Intersect((_1091_v16)), new BigNumber(675), (_dafny.ZERO).minus(_1084_cf9));
              _1092_v17 = _nw198;
              let _1093_v18;
              _1093_v18 = _dafny.Set.fromElements(_1084_cf9);
              let _1094_v19;
              _1094_v19 = _dafny.Seq.of(_1093_v18);
              let _rhs194 = new BigNumber((_1094_v19).length);
              let _rhs195 = _1040_v0;
              let _rhs196 = !(_1086_cf7);
              let _lhs170 = globalState;
              let _lhs171 = globalState;
              let _lhs172 = globalState;
              _lhs170.f8 = _rhs194;
              _lhs171.f1 = _rhs195;
              _lhs172.f6 = _rhs196;
              let _rhs197 = _module.__default.safeDivisionInt((new BigNumber(-321)).plus(p2), _module.__default.safeModuloInt(_1084_cf9, new BigNumber(319)));
              let _rhs198 = _1041_v1;
              let _lhs173 = globalState;
              _lhs173.f14 = _rhs197;
              _1041_v1 = _rhs198;
              let _1095_v20;
              _1095_v20 = _dafny.Seq.of(_1084_cf9);
              let _1096_v21;
              let _nw199 = new _module.C5();
              _nw199.__ctor((_1092_v17).fm3(_1065_v6, _1040_v0, p1, _1095_v20, globalState), !(_1040_v0));
              _1096_v21 = _nw199;
              let _1097_v22;
              _1097_v22 = _dafny.MultiSet.fromElements(_1096_v21);
              let _1098_v23;
              _1098_v23 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((p1).plus(_1083_cf10)),new BigNumber(((_1097_v22).Difference(_1097_v22)).cardinality()));
              let _1099_v25;
              _1099_v25 = _dafny.MultiSet.fromElements(p2, _1092_v17.f31);
              _1098_v23 = function () {
                let _coll36 = new _dafny.Map();
                for (const _compr_36 of ((_1099_v25).update((_1096_v21).fm13(_1083_cf10, _1040_v0, globalState), _module.__default.abs(new BigNumber((_1065_v6).length)))).Elements) {
                  let _1100_v24 = _compr_36;
                  if (((_1099_v25).update((_1096_v21).fm13(_1083_cf10, _1040_v0, globalState), _module.__default.abs(new BigNumber((_1065_v6).length)))).contains(_1100_v24)) {
                    _coll36.push([(_1100_v24).plus(_1084_cf9),_module.__default.safeModuloInt(p0, p0)]);
                  }
                }
                return _coll36;
              }();
            } else if (_source11.is_DC3) {
              let _1101___mcc_h16 = (_source11).cf4;
              let _1102_cf4 = _1101___mcc_h16;
              (globalState).f17 = _module.__default.fm1(p1, true, globalState);
              let _1103_v26;
              _1103_v26 = _dafny.Set.fromElements(p0, p0);
              let _1104_v27;
              _1104_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1103_v26).length),_module.__default.fm9(_1040_v0, globalState));
              let _1105_v28;
              _1105_v28 = _dafny.MultiSet.fromElements(_module.__default.fm1(new BigNumber((_1103_v26).length), _1040_v0, globalState));
              let _1106_v29;
              _1106_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1104_v27,(_1105_v28).IsSubsetOf(_1105_v28));
              _1106_v29 = ((_1106_v29).Merge(_1106_v29)).Merge(_1106_v29);
              (globalState).f8 = (_dafny.ZERO).minus(p0);
              _1045_v5 = (_1045_v5).Merge(_module.__default.fm10(_1040_v0, p2, _1102_cf4, _1041_v1, globalState));
            } else {
              let _1107___mcc_h17 = (_source11).cf12;
              let _1108_cf12 = _1107___mcc_h17;
              let _1109_v30;
              let _init35 = ((_1110_p1) => function (_1111_i2) {
                return (_1111_i2).minus(_1110_p1);
              })(p1);
              let _nw200 = Array((new BigNumber(27)).toNumber());
              for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw200.length); _i0_35++) {
                _nw200[_i0_35] = _init35(new BigNumber(_i0_35));
              }
              _1109_v30 = _nw200;
              _1109_v30 = p3;
              (globalState).f17 = p0;
              let _1112_v31;
              _1112_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1040_v0,_1065_v6);
              let _1113_v32;
              _1113_v32 = _dafny.Seq.of((((_1112_v31).contains(_1040_v0)) ? ((_1112_v31).get(_1040_v0)) : (_dafny.Seq.UnicodeFromString("buqwaquu"))), _1065_v6);
              let _1114_v34;
              _1114_v34 = _dafny.Seq.of(_1040_v0, _1040_v0, _1040_v0);
              let _1115_v35;
              _1115_v35 = _dafny.Seq.of(p1);
              let _1116_v36;
              _1116_v36 = _dafny.Set.fromElements(_1115_v35);
              let _1117_v37;
              let _nw201 = Array((new BigNumber(25)).toNumber());
              _nw201[(_dafny.ZERO).toNumber()] = _1040_v0;
              _nw201[(_dafny.ONE).toNumber()] = (function () {
                let _coll37 = new _dafny.Set();
                for (const _compr_37 of (_dafny.MultiSet.FromArray(_1113_v32)).Elements) {
                  let _1118_v33 = _compr_37;
                  if ((_dafny.MultiSet.FromArray(_1113_v32)).contains(_1118_v33)) {
                    _coll37.add(_1118_v33);
                  }
                }
                return _coll37;
              }()).IsDisjointFrom(_dafny.Set.fromElements(_1065_v6));
              _nw201[(new BigNumber(2)).toNumber()] = !(true);
              _nw201[(new BigNumber(3)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(4)).toNumber()] = true;
              _nw201[(new BigNumber(5)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(6)).toNumber()] = (_1040_v0) || (true);
              _nw201[(new BigNumber(7)).toNumber()] = false;
              _nw201[(new BigNumber(8)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(9)).toNumber()] = (_1114_v34)[_module.__default.safeIndex(p1, new BigNumber((_1114_v34).length))];
              _nw201[(new BigNumber(10)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(11)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(12)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(13)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(14)).toNumber()] = true;
              _nw201[(new BigNumber(15)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(16)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(17)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(18)).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_1065_v6, _module.__default.safeIndex(p2, new BigNumber((_1065_v6).length)), (_1044_v4)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length))])).length),_1040_v0)).equals(_dafny.Map.Empty.slice().updateUnsafe(p2,_1040_v0));
              _nw201[(new BigNumber(19)).toNumber()] = !(_1040_v0);
              _nw201[(new BigNumber(20)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(21)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(22)).toNumber()] = _1040_v0;
              _nw201[(new BigNumber(23)).toNumber()] = !(_1040_v0);
              _nw201[(new BigNumber(24)).toNumber()] = !(_1116_v36).contains(_1115_v35);
              _1117_v37 = _nw201;
              let _index162 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1117_v37).length));
              (_1117_v37)[_index162] = _1040_v0;
              let _index163 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1117_v37).length));
              (_1117_v37)[_index163] = !(_module.__default.safeModuloInt(new BigNumber((_1114_v34).length), p2)).isEqualTo((p1).minus(p0));
              let _1119_v38;
              _1119_v38 = _module.D0.create_DC2(_module.D0.create_DC1(_1040_v0, p2));
              let _1120_v39;
              _1120_v39 = _module.D0.create_DC2(_1119_v38);
              let _1121_v40;
              _1121_v40 = _module.D0.create_DC2(_1119_v38);
              let _1122_v41;
              _1122_v41 = _module.D0.create_DC2(_1119_v38);
              let _1123_v42;
              _1123_v42 = _module.D0.create_DC2((_1122_v41).dtor_cf3);
              let _1124_v43;
              _1124_v43 = _module.D0.create_DC2(_1119_v38);
              let _1125_v44;
              _1125_v44 = _module.D0.create_DC2(_1121_v40);
              let _1126_v45;
              _1126_v45 = _dafny.Seq.of(_1120_v39);
              let _1127_v46;
              _1127_v46 = _module.D0.create_DC2((_1126_v45)[_module.__default.safeIndex(p1, new BigNumber((_1126_v45).length))]);
              let _1128_v47;
              _1128_v47 = _module.D0.create_DC2(_1125_v44);
              let _1129_v48;
              let _nw202 = new _module.C6();
              _nw202.__ctor(_1128_v47, (_1117_v37)[_module.__default.safeIndex(new BigNumber(741), new BigNumber((_1117_v37).length))], (_1117_v37)[_module.__default.safeIndex(new BigNumber(741), new BigNumber((_1117_v37).length))], _module.__default.safeModuloInt((_dafny.ZERO).minus(p0), p1));
              _1129_v48 = _nw202;
            }
            let _1130_v49;
            let _nw203 = Array((new BigNumber(28)).toNumber()).fill(false);
            _1130_v49 = _nw203;
            let _1131_v50;
            _1131_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1040_v0,_1130_v49);
            let _1132_v51;
            _1132_v51 = _dafny.Seq.of(_1131_v50, _1131_v50);
            _1131_v50 = (((_1132_v51)[_module.__default.safeIndex(p0, new BigNumber((_1132_v51).length))]).Merge(_1131_v50)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1040_v0,_1130_v49));
          }
        }
      }
      let _1133_v52;
      let _init36 = ((_1134_p1) => function (_1135_i4) {
        return (_1135_i4).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1134_p1)).length))))).length))).length));
      })(p1);
      let _nw204 = Array((new BigNumber(19)).toNumber());
      for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw204.length); _i0_36++) {
        _nw204[_i0_36] = _init36(new BigNumber(_i0_36));
      }
      _1133_v52 = _nw204;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1133_v52).length))) {
        let _1136_i3 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1136_i3)) && ((_1136_i3).isLessThan(new BigNumber((_1133_v52).length))))) {
          (_1133_v52)[(_1136_i3)] = _module.__default.safeModuloInt(_1136_i3, (p2).minus(p2));
        }
      }
      let _1137_v53;
      _1137_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1040_v0,p3);
      let _1138_v54;
      _1138_v54 = _dafny.Seq.UnicodeFromString("etgsrnnx");
      _1137_v53 = (_1137_v53).update(_dafny.Seq.IsProperPrefixOf(_1138_v54, _1138_v54), _1133_v52);
      let _1139_v56;
      _1139_v56 = _dafny.Seq.of(p2);
      let _1140_v57;
      _1140_v57 = _module.D10.create_DC28(_1139_v56);
      if (!(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(778))).equals((function () {
        let _coll38 = new _dafny.Map();
        for (const _compr_38 of ((_1140_v57).dtor_cf52).Elements) {
          let _1141_v55 = _compr_38;
          if (_dafny.Seq.contains((_1140_v57).dtor_cf52, _1141_v55)) {
            _coll38.push([_module.__default.safeModuloInt(_1141_v55, p2),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,_1040_v0)).length)]);
          }
        }
        return _coll38;
      }()))) {
        let _1142_v58;
        let _nw205 = new _module.C2();
        _nw205.__ctor(p1);
        _1142_v58 = _nw205;
        (globalState).f17 = ((_dafny.ZERO).minus(p2)).minus((p1).minus(new BigNumber(518)));
        let _1143_v59;
        _1143_v59 = _dafny.MultiSet.fromElements(p1);
        if ((_1143_v59).IsProperSubsetOf(_1143_v59)) {
          (globalState).f6 = _1040_v0;
          (globalState).f14 = new BigNumber(-370);
          (globalState).f9 = _1040_v0;
          let _1144_v60;
          _1144_v60 = _dafny.Seq.of((_1143_v59).Union(_1143_v59), (_1143_v59).Intersect(_1143_v59));
          _1144_v60 = _dafny.Seq.of(_1143_v59);
          let _1145_v61;
          let _1146_v62;
          let _1147_v63;
          let _out11;
          let _out12;
          let _out13;
          let _outcollector2 = (_1142_v58).m0(_1040_v0, _1040_v0, globalState);
          _out11 = _outcollector2[0];
          _out12 = _outcollector2[1];
          _out13 = _outcollector2[2];
          _1145_v61 = _out11;
          _1146_v62 = _out12;
          _1147_v63 = _out13;
        } else {
          let _1148_v64;
          _1148_v64 = _dafny.MultiSet.fromElements(_1040_v0, (_1040_v0) === (_1040_v0), !(false), !(_1045_v5).equals(_1045_v5));
          _1148_v64 = _1148_v64;
          (globalState).f15 = _1040_v0;
          let _1149_v65;
          _1149_v65 = _dafny.Seq.of(_1040_v0);
          let _1150_v66;
          _1150_v66 = _module.D6.create_DC16(_1149_v65);
          let _1151_v67;
          _1151_v67 = _dafny.Seq.of(_1150_v66, _module.D6.create_DC16(_1149_v65));
          let _1152_v68;
          _1152_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1151_v67,_1040_v0);
          _1152_v68 = _module.__default.fm40(_1040_v0, globalState);
          let _1153_v69;
          let _nw206 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _1153_v69 = _nw206;
          let _1154_v70;
          _1154_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1153_v69,(_1040_v0) || (_1040_v0));
          _1154_v70 = (_1154_v70).update(_1153_v69, _1040_v0);
          let _1155_v71;
          _1155_v71 = _dafny.Map.Empty.slice().updateUnsafe((p0).plus(p0),_1153_v69);
          _1155_v71 = (_1155_v71).update((p2).minus(p2), _1153_v69);
        }
        (globalState).f8 = p1;
        let _rhs199 = !(_1040_v0);
        let _rhs200 = (_1040_v0) === ((_1040_v0) || (_1040_v0));
        let _rhs201 = _1040_v0;
        let _rhs202 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), ((_1156_v1) => function (_1157_i5) {
          return _1156_v1;
        })(_1041_v1));
        let _rhs203 = new BigNumber(264);
        let _lhs174 = globalState;
        let _lhs175 = globalState;
        let _lhs176 = globalState;
        let _lhs177 = globalState;
        _lhs174.f1 = _rhs199;
        _lhs175.f6 = _rhs200;
        _lhs176.f6 = _rhs201;
        _1138_v54 = _rhs202;
        _lhs177.f5 = _rhs203;
      } else {
        if (true) {
          _1138_v54 = _dafny.Seq.of(_1041_v1);
          let _1158_v72;
          let _nw207 = new _module.C4();
          _nw207.__ctor(!(_1040_v0), !(true), (new BigNumber(456)).isLessThan(p2));
          _1158_v72 = _nw207;
          _1158_v72 = _1158_v72;
          _1138_v54 = _dafny.Seq.Concat(_1138_v54, _1138_v54);
          let _1159_v73;
          _1159_v73 = _dafny.Map.Empty.slice().updateUnsafe((_1158_v72).f27,p1);
          _1159_v73 = (_1159_v73).update(!(!((_1158_v72).f27)), p0);
          let _1160_v74;
          _1160_v74 = _dafny.Seq.of(_1040_v0);
          let _1161_v75;
          let _nw208 = Array((new BigNumber(8)).toNumber());
          _nw208[(_dafny.ZERO).toNumber()] = _1139_v56;
          _nw208[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(428)), ((_1162_p2) => function (_1163_i6) {
            return _1162_p2;
          })(p2)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(787)), ((_1164_p1) => function (_1165_i7) {
            return _1164_p1;
          })(p1)));
          _nw208[(new BigNumber(2)).toNumber()] = _1139_v56;
          _nw208[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_1160_v74).length), p1), _1139_v56);
          _nw208[(new BigNumber(4)).toNumber()] = (_1140_v57).dtor_cf52;
          _nw208[(new BigNumber(5)).toNumber()] = _1139_v56;
          _nw208[(new BigNumber(6)).toNumber()] = _1139_v56;
          _nw208[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1139_v56, _1139_v56);
          _1161_v75 = _nw208;
          let _index164 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1161_v75).length));
          (_1161_v75)[_index164] = _1139_v56;
          let _1166_v76;
          _1166_v76 = _dafny.Set.fromElements(new BigNumber((_1139_v56).length), p2, p1, new BigNumber((_1160_v74).length), new BigNumber((_dafny.Seq.UnicodeFromString("luqkew")).length));
          let _1167_v77;
          _1167_v77 = _dafny.MultiSet.fromElements(p1);
          let _1168_v80;
          _1168_v80 = _module.D8.create_DC23(p0, _1166_v76, p2, p0, p2);
          let _1169_v81;
          _1169_v81 = _module.D0.create_DC1(_1040_v0, new BigNumber(244));
          let _1170_v82;
          _1170_v82 = _dafny.MultiSet.fromElements(_1169_v81, _1169_v81);
          let _1171_v83;
          let _nw209 = Array((new BigNumber(23)).toNumber()).fill(false);
          _1171_v83 = _nw209;
          let _1172_v84;
          _1172_v84 = _dafny.Set.fromElements(_1171_v83);
          let _1173_v85;
          let _nw210 = new _module.C0();
          _nw210.__ctor(_1172_v84, new BigNumber((_1138_v54).length), p0);
          _1173_v85 = _nw210;
          let _1174_v86;
          _1174_v86 = _module.D8.create_DC24(p2, (_1044_v4)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length))], (_1170_v82).update(_1169_v81, _module.__default.abs(new BigNumber((_module.__default.fm32(_1040_v0, _dafny.Seq.of(new BigNumber((_1042_v2).length)), _1160_v74, _dafny.Set.fromElements(_module.D1.create_DC4((_1158_v72).f27, (_1158_v72).f27)), globalState)).length))), _1040_v0, _1173_v85);
          let _1175_v87;
          _1175_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1173_v85.f31,_1173_v85);
          let _1176_v90;
          let _nw211 = Array((new BigNumber(21)).toNumber());
          _nw211[(_dafny.ZERO).toNumber()] = _1166_v76;
          _nw211[(_dafny.ONE).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(2)).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(3)).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(4)).toNumber()] = function () {
            let _coll39 = new _dafny.Set();
            for (const _compr_39 of (_1167_v77).Elements) {
              let _1177_v78 = _compr_39;
              if ((_1167_v77).contains(_1177_v78)) {
                _coll39.add((_1177_v78).plus(new BigNumber((_dafny.Seq.of(new BigNumber(292))).length)));
              }
            }
            return _coll39;
          }();
          _nw211[(new BigNumber(5)).toNumber()] = (_dafny.Set.fromElements(new BigNumber(-874), new BigNumber((_1045_v5).length))).Difference(_1166_v76);
          _nw211[(new BigNumber(6)).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(7)).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(8)).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(9)).toNumber()] = (_1166_v76).Union(function () {
            let _coll40 = new _dafny.Set();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(628), new BigNumber(436))) {
              let _1178_v79 = _compr_40;
              if (((new BigNumber(628)).isLessThanOrEqualTo(_1178_v79)) && ((_1178_v79).isLessThan(new BigNumber(436)))) {
                _coll40.add((_1178_v79).plus(p1));
              }
            }
            return _coll40;
          }());
          _nw211[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(p2, p1, new BigNumber((_1045_v5).length));
          _nw211[(new BigNumber(11)).toNumber()] = (_1168_v80).dtor_cf41;
          _nw211[(new BigNumber(12)).toNumber()] = (_1166_v76).Difference(_1166_v76);
          _nw211[(new BigNumber(13)).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(p2, p2, (_1174_v86).dtor_cf45, new BigNumber((_1175_v87).length));
          _nw211[(new BigNumber(15)).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(16)).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(17)).toNumber()] = (function () {
            let _coll41 = new _dafny.Set();
            for (const _compr_41 of _dafny.IntegerRange(new BigNumber(51), new BigNumber(172))) {
              let _1179_v88 = _compr_41;
              if (((new BigNumber(51)).isLessThanOrEqualTo(_1179_v88)) && ((_1179_v88).isLessThan(new BigNumber(172)))) {
                _coll41.add(_module.__default.safeDivisionInt(_1179_v88, new BigNumber((_1138_v54).length)));
              }
            }
            return _coll41;
          }()).Union(_1166_v76);
          _nw211[(new BigNumber(18)).toNumber()] = _module.__default.fm41(globalState);
          _nw211[(new BigNumber(19)).toNumber()] = _1166_v76;
          _nw211[(new BigNumber(20)).toNumber()] = function () {
            let _coll42 = new _dafny.Set();
            for (const _compr_42 of _dafny.IntegerRange(new BigNumber(806), new BigNumber(-107))) {
              let _1180_v89 = _compr_42;
              if (((new BigNumber(806)).isLessThanOrEqualTo(_1180_v89)) && ((_1180_v89).isLessThan(new BigNumber(-107)))) {
                _coll42.add((_1180_v89).minus(_1173_v85.f31));
              }
            }
            return _coll42;
          }();
          _1176_v90 = _nw211;
          let _index165 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1161_v75).length));
          let _rhs204 = _1139_v56;
          let _rhs205 = _1176_v90;
          let _rhs206 = (_1158_v72).f27;
          let _lhs178 = _1161_v75;
          let _lhs179 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1161_v75).length));
          let _lhs180 = globalState;
          _lhs178[_lhs179] = _rhs204;
          _1176_v90 = _rhs205;
          _lhs180.f1 = _rhs206;
        } else {
          let _1181_v91;
          _1181_v91 = _dafny.MultiSet.fromElements(_1040_v0, _1040_v0);
          _1040_v0 = ((_1181_v91).Difference(_1181_v91)).contains(false);
          let _1182_v92;
          let _nw212 = new _module.C5();
          _nw212.__ctor(_module.__default.fm29(new BigNumber((_1181_v91).cardinality()), p2, _1040_v0, globalState), _1040_v0);
          _1182_v92 = _nw212;
          let _nw213 = new _module.C5();
          _nw213.__ctor(_1040_v0, _1040_v0);
          _1182_v92 = _nw213;
          let _index166 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length));
          (_1044_v4)[_index166] = _1041_v1;
          let _rhs207 = ((false) ? (_module.__default.safeModuloInt(new BigNumber(698), p2)) : (new BigNumber((_1138_v54).length)));
          let _rhs208 = _1138_v54;
          let _lhs181 = globalState;
          _lhs181.f17 = _rhs207;
          _1138_v54 = _rhs208;
          let _index167 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length));
          (_1044_v4)[_index167] = _1041_v1;
        }
        (globalState).f14 = _module.__default.fm1(p0, true, globalState);
        let _1183_v93;
        _1183_v93 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1040_v0);
        let _1184_v94;
        _1184_v94 = _module.D0.create_DC0((((_1183_v93).contains(p0)) ? ((_1183_v93).get(p0)) : (_1040_v0)));
        let _1185_v95;
        _1185_v95 = _module.D0.create_DC2(_1184_v94);
        let _source12 = _1185_v95;
        if (_source12.is_DC1) {
          let _1186___mcc_h18 = (_source12).cf1;
          let _1187___mcc_h19 = (_source12).cf2;
          let _1188_cf2 = _1187___mcc_h19;
          let _1189_cf1 = _1186___mcc_h18;
          let _1190_v96;
          let _nw214 = new _module.C6();
          _nw214.__ctor(_1185_v95, _1189_cf1, _1189_cf1, _module.__default.fm1(new BigNumber((_dafny.Seq.UnicodeFromString("ohhj")).length), _1189_cf1, globalState));
          _1190_v96 = _nw214;
          let _1191_v97;
          _1191_v97 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("urgll"),_1190_v96);
          _1191_v97 = (_1191_v97).update(_1138_v54, _1190_v96);
          let _1192_v98;
          let _nw215 = new _module.C2();
          _nw215.__ctor((_1190_v96).fm5(p0, new BigNumber((_module.__default.fm31(_1189_cf1, _1040_v0, globalState)).length), globalState));
          _1192_v98 = _nw215;
          let _index168 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1133_v52).length));
          (_1133_v52)[_index168] = _1188_cf2;
          let _1193_v99;
          _1193_v99 = _dafny.Seq.of(_1189_cf1, _1189_cf1);
          let _1194_v100;
          _1194_v100 = _dafny.Map.Empty.slice().updateUnsafe(_1193_v99,p2);
          let _1195_v101;
          _1195_v101 = _dafny.Map.Empty.slice().updateUnsafe((_1190_v96).fm2(globalState),_1192_v98);
          let _1196_v102;
          _1196_v102 = _dafny.Seq.of((_1195_v101).update(p1, _1192_v98), _1195_v101, _1195_v101);
          let _1197_v103;
          _1197_v103 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1196_v102).length),p1);
          let _1198_v104;
          _1198_v104 = _dafny.MultiSet.fromElements(_1041_v1, (_1044_v4)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length))], _1041_v1);
          let _1199_v105;
          _1199_v105 = _module.D0.create_DC0((_1198_v104).IsProperSubsetOf(_1198_v104));
          let _1200_v106;
          _1200_v106 = _module.D15.create_DC36(_1194_v100);
          let _index169 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1133_v52).length));
          let _rhs209 = _1188_cf2;
          let _rhs210 = (_1194_v100).Merge((_1200_v106).dtor_cf64);
          let _rhs211 = _1138_v54;
          let _rhs212 = (_module.__default.fm8(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1188_cf2,new BigNumber(((_1183_v93).update(new BigNumber((function () {
            let _coll43 = new _dafny.Map();
            for (const _compr_43 of _dafny.IntegerRange(new BigNumber(109), new BigNumber(642))) {
              let _1201_v107 = _compr_43;
              if (((new BigNumber(109)).isLessThanOrEqualTo(_1201_v107)) && ((_1201_v107).isLessThan(new BigNumber(642)))) {
                _coll43.push([(_1201_v107).multipliedBy(p1),_1189_cf1]);
              }
            }
            return _coll43;
          }()).length), _1189_cf1)).length))).length), globalState)).Merge((_1197_v103).Merge(_1197_v103));
          let _rhs213 = _1199_v105;
          let _lhs182 = _1133_v52;
          let _lhs183 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_1133_v52).length));
          _lhs182[_lhs183] = _rhs209;
          _1194_v100 = _rhs210;
          _1138_v54 = _rhs211;
          _1197_v103 = _rhs212;
          _1199_v105 = _rhs213;
          _1188_cf2 = (_1133_v52)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_1133_v52).length))];
        } else if (_source12.is_DC0) {
          let _1202___mcc_h20 = (_source12).cf0;
          let _1203_cf0 = _1202___mcc_h20;
          let _index170 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((p3).length));
          (p3)[_index170] = p0;
          let _index171 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((p3).length));
          (p3)[_index171] = _module.__default.fm1(p0, _1203_cf0, globalState);
          (globalState).f9 = false;
          let _1204_v108;
          let _nw216 = new _module.C3();
          _nw216.__ctor(_1040_v0, p1, _1040_v0, _1203_cf0, p1);
          _1204_v108 = _nw216;
          let _1205_v109;
          let _nw217 = Array((new BigNumber(18)).toNumber());
          _nw217[(_dafny.ZERO).toNumber()] = _1204_v108;
          _nw217[(_dafny.ONE).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(2)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(3)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(4)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(5)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(6)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(7)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(8)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(9)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(10)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(11)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(12)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(13)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(14)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(15)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(16)).toNumber()] = _1204_v108;
          _nw217[(new BigNumber(17)).toNumber()] = _1204_v108;
          _1205_v109 = _nw217;
          let _1206_v110;
          _1206_v110 = _module.D16.create_DC39(_1205_v109);
          _1205_v109 = (_1206_v110).dtor_cf69;
          let _1207_v111;
          let _nw218 = new _module.C3();
          _nw218.__ctor(_1203_cf0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_1203_cf0)).length), _1203_cf0, (_1204_v108).f28, new BigNumber(150));
          _1207_v111 = _nw218;
          _1207_v111 = _1207_v111;
        } else {
          let _1208___mcc_h21 = (_source12).cf3;
          let _1209_cf3 = _1208___mcc_h21;
          let _1210_v112;
          let _nw219 = Array((new BigNumber(6)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1210_v112 = _nw219;
          let _1211_v113;
          _1211_v113 = _dafny.MultiSet.fromElements(_1040_v0);
          let _index172 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_1210_v112).length));
          (_1210_v112)[_index172] = _1211_v113;
          let _index173 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_1210_v112).length));
          (_1210_v112)[_index173] = _1211_v113;
          let _1212_v114;
          _1212_v114 = _dafny.Seq.of(_1040_v0);
          let _1213_v115;
          _1213_v115 = _dafny.Set.fromElements((_1212_v114)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1138_v54, _module.__default.safeIndex(p2, new BigNumber((_1138_v54).length)), (_1044_v4)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length))])).length), new BigNumber((_1212_v114).length))]);
          let _1214_v116;
          _1214_v116 = _dafny.MultiSet.fromElements(_1213_v115, _1213_v115, _1213_v115, _dafny.Set.fromElements(_1040_v0, _1040_v0, _1040_v0, !(_1040_v0)));
          _1214_v116 = _dafny.MultiSet.fromElements(_1213_v115);
          let _1215_v117;
          _1215_v117 = _dafny.Map.Empty.slice().updateUnsafe(_1212_v114,_1212_v114);
          _1215_v117 = (_1215_v117).update(_1212_v114, _1212_v114);
          _1138_v54 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_1138_v54, _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_1138_v54, _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))))).length)), _1041_v1), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1138_v54, _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_1138_v54, _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))))).length)), _1041_v1)).length)), _1041_v1), _dafny.Seq.Concat(_1138_v54, _1138_v54));
        }
        let _1216_v118;
        let _nw220 = new _module.C5();
        _nw220.__ctor(!(_1040_v0), _1040_v0);
        _1216_v118 = _nw220;
        let _1217_v119;
        _1217_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1216_v118,(((_1183_v93).contains(p1)) ? ((_1183_v93).get(p1)) : ((_1216_v118).f25)));
        if (_module.__default.fm29(new BigNumber((_1217_v119).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(p2)), !((_1216_v118).f25) || (_1216_v118.f24), globalState)) {
          let _1218_v120;
          _1218_v120 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          (globalState).f5 = (((_1218_v120).contains(p2)) ? ((_1218_v120).get(p2)) : ((_dafny.ZERO).minus(p1)));
          let _1219_v121;
          _1219_v121 = _module.D3.create_DC8(_1138_v54);
          let _1220_v122;
          _1220_v122 = _module.D8.create_DC22(p3);
          let _1221_v123;
          _1221_v123 = _dafny.MultiSet.fromElements(_1133_v52, (_1220_v122).dtor_cf39);
          let _1222_v124;
          _1222_v124 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_module.__default.fm25(_1040_v0, (_1216_v118).f25, new _dafny.CodePoint('j'.codePointAt(0)), globalState), (_1219_v121).dtor_cf14),_1221_v123);
          _1222_v124 = (_1222_v124).update(_1138_v54, (_1221_v123).update(_1133_v52, _module.__default.abs(p0)));
          (globalState).f5 = _module.__default.safeDivisionInt(new BigNumber((_1138_v54).length), new BigNumber((_1218_v120).length));
          _1041_v1 = (_1044_v4)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length))];
          _1138_v54 = _dafny.Seq.update(_dafny.Seq.Concat(_1138_v54, _dafny.Seq.Create(_module.__default.abs(new BigNumber(151)), ((_1223_v1) => function (_1224_i8) {
            return _1223_v1;
          })(_1041_v1))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_1138_v54, _dafny.Seq.Create(_module.__default.abs(new BigNumber(151)), ((_1225_v1) => function (_1226_i8) {
            return _1225_v1;
          })(_1041_v1)))).length)), (_1044_v4)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_1044_v4).length))]);
        } else {
          let _1227_v125;
          _1227_v125 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-822)).plus(p0),_1133_v52);
          let _1228_v126;
          _1228_v126 = _module.D15.create_DC37((_1216_v118).f25, _dafny.Map.Empty.slice().updateUnsafe(_1040_v0,(_1216_v118).f25), p2);
          let _1229_v127;
          _1229_v127 = _dafny.Map.Empty.slice().updateUnsafe((_1228_v126).dtor_cf65,p0);
          let _1230_v128;
          _1230_v128 = _dafny.MultiSet.fromElements(_1229_v127);
          _1227_v125 = (_1227_v125).update(_module.__default.safeDivisionInt(new BigNumber((_1230_v128).cardinality()), p2), p3);
          let _index174 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_1133_v52).length));
          (_1133_v52)[_index174] = new BigNumber((_1138_v54).length);
          let _1231_v129;
          let _nw221 = new _module.C6();
          _nw221.__ctor(_1185_v95, (_1216_v118).f25, (_1216_v118).f25, p0);
          _1231_v129 = _nw221;
          let _1232_v130;
          _1232_v130 = _dafny.Map.Empty.slice().updateUnsafe(_1231_v129,(p2).isLessThan(new BigNumber(151)));
          let _1233_v131;
          let _init37 = ((_1234_v118) => function (_1235_i9) {
            return _1234_v118.f24;
          })(_1216_v118);
          let _nw222 = Array((new BigNumber(24)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw222.length); _i0_37++) {
            _nw222[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1233_v131 = _nw222;
          let _index175 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_1133_v52).length));
          let _rhs214 = (((_1232_v130).contains(_1231_v129)) ? ((_1232_v130).get(_1231_v129)) : ((_1216_v118).f25));
          let _rhs215 = _1040_v0;
          let _rhs216 = (p0).multipliedBy((_1139_v56)[_module.__default.safeIndex(p0, new BigNumber((_1139_v56).length))]);
          let _rhs217 = _1233_v131;
          let _lhs184 = globalState;
          let _lhs185 = _1133_v52;
          let _lhs186 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_1133_v52).length));
          _1040_v0 = _rhs214;
          _lhs184.f6 = _rhs215;
          _lhs185[_lhs186] = _rhs216;
          r1 = _rhs217;
          (globalState).f15 = _1040_v0;
          let _1236_v132;
          _1236_v132 = _1233_v131;
          let _1237_v133;
          _1237_v133 = _dafny.Map.Empty.slice().updateUnsafe(_1040_v0,(_1236_v132));
          let _1238_v134;
          let _nw223 = Array((new BigNumber(28)).toNumber());
          _nw223[(_dafny.ZERO).toNumber()] = _1233_v131;
          _nw223[(_dafny.ONE).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(2)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(3)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(4)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(5)).toNumber()] = (_1236_v132);
          _nw223[(new BigNumber(6)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(7)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(8)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(9)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(10)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(11)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(12)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(13)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(14)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(15)).toNumber()] = (((_1237_v133).contains(_1216_v118.f24)) ? ((_1237_v133).get(_1216_v118.f24)) : (_1233_v131));
          _nw223[(new BigNumber(16)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(17)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(18)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(19)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(20)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(21)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(22)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(23)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(24)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(25)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(26)).toNumber()] = _1233_v131;
          _nw223[(new BigNumber(27)).toNumber()] = _1233_v131;
          _1238_v134 = _nw223;
          let _index176 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1238_v134).length));
          (_1238_v134)[_index176] = (((_1237_v133).contains((_1216_v118).f25)) ? ((_1237_v133).get((_1216_v118).f25)) : (_1233_v131));
          let _index177 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1238_v134).length));
          let _rhs218 = _1233_v131;
          let _rhs219 = (_1133_v52)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_1133_v52).length))];
          let _rhs220 = !(!(!(_module.__default.fm29((_1133_v52)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_1133_v52).length))], p2, (_1216_v118).f25, globalState))));
          let _lhs187 = _1238_v134;
          let _lhs188 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1238_v134).length));
          let _lhs189 = globalState;
          let _lhs190 = globalState;
          _lhs187[_lhs188] = _rhs218;
          _lhs189.f17 = _rhs219;
          _lhs190.f15 = _rhs220;
          let _1239_v135;
          let _init38 = ((_1240_v93) => function (_1241_i10) {
            return (_1240_v93).Merge(_1240_v93);
          })(_1183_v93);
          let _nw224 = Array((new BigNumber(8)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw224.length); _i0_38++) {
            _nw224[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1239_v135 = _nw224;
          let _1242_v136;
          _1242_v136 = _dafny.MultiSet.fromElements(p2);
          let _index178 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_1133_v52).length));
          let _rhs221 = new BigNumber(481);
          let _rhs222 = _1239_v135;
          let _rhs223 = (_1216_v118).f25;
          let _rhs224 = ((_1133_v52)[_module.__default.safeIndex(new BigNumber(884), new BigNumber((_1133_v52).length))]).multipliedBy(new BigNumber(((_1242_v136).Intersect(_1242_v136)).cardinality()));
          let _lhs191 = globalState;
          let _lhs192 = globalState;
          let _lhs193 = _1133_v52;
          let _lhs194 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_1133_v52).length));
          _lhs191.f5 = _rhs221;
          _1239_v135 = _rhs222;
          _lhs192.f1 = _rhs223;
          _lhs193[_lhs194] = _rhs224;
        }
        let _1243_v137;
        _1243_v137 = _dafny.Map.Empty.slice().updateUnsafe(_1138_v54,p0);
        _1243_v137 = (_1243_v137).update(_1138_v54, p2);
      }
      let _1244_v138;
      let _nw225 = new _module.C3();
      _nw225.__ctor(_1040_v0, new BigNumber((_1139_v56).length), _1040_v0, _1040_v0, p2);
      _1244_v138 = _nw225;
      let _1245_v139;
      _1245_v139 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1138_v54).length),_1244_v138);
      let _1246_v140;
      _1246_v140 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1245_v139);
      let _1247_v141;
      _1247_v141 = _dafny.Seq.of(_module.__default.fm9(_1244_v138.f24, globalState), _module.__default.fm29(p1, p2, _1040_v0, globalState));
      r0 = (((_1246_v140).contains(new BigNumber((_1247_v141).length))) ? ((_1246_v140).get(new BigNumber((_1247_v141).length))) : ((_1245_v139).Merge(_1245_v139)));
      let _1248_v142;
      let _nw226 = Array((new BigNumber(26)).toNumber()).fill(false);
      _1248_v142 = _nw226;
      r1 = _1248_v142;
      return [r0, r1];
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f24 = false;
      this._f23 = _dafny.ZERO;
      this._f25 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    set f24(value) {
      let _this = this;
      _this._f24 = value;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    __ctor(f23, f24, f25) {
      let _this = this;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm2(globalState) {
      let _this = this;
      return (_this).f23;
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("s"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("motvhqyi"), _dafny.Seq.UnicodeFromString("hcsjvtutw")));
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_this.f24))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f25)))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_dafny.Map.Empty.slice().updateUnsafe(_this.f24,(_this).f25)));
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return (_this).f23;
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_this).f25);
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = false;
      (_this).f24 = false;
      let _1249_i0;
      _1249_i0 = _dafny.ZERO;
      L9: {
        while (!(!(!((_this).f25) || (p1)) || (_this.f24))) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1249_i0)) {
              break L9;
            }
            _1249_i0 = (_1249_i0).plus(_dafny.ONE);
            let _1250_v0;
            _1250_v0 = _dafny.Seq.UnicodeFromString("luxyg");
            let _1251_v1;
            _1251_v1 = _dafny.Seq.of(new BigNumber((_1250_v0).length));
            let _1252_v2;
            _1252_v2 = _module.D0.create_DC0(p0);
            let _1253_v3;
            let _nw227 = Array((new BigNumber(22)).toNumber());
            _nw227[(_dafny.ZERO).toNumber()] = (_this).fm6(new BigNumber((_dafny.Seq.of((_this).f23)).length), _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), ((_1254_p1) => function (_1255_i2) {
              return new BigNumber((_dafny.Seq.of(_1254_p1, true)).length);
            })(p1)), _1251_v1), (_this).f25, (_this).f23, globalState);
            _nw227[(_dafny.ONE).toNumber()] = (_this).f25;
            _nw227[(new BigNumber(2)).toNumber()] = (_this).f25;
            _nw227[(new BigNumber(3)).toNumber()] = true;
            _nw227[(new BigNumber(4)).toNumber()] = p1;
            _nw227[(new BigNumber(5)).toNumber()] = (_this).f25;
            _nw227[(new BigNumber(6)).toNumber()] = true;
            _nw227[(new BigNumber(7)).toNumber()] = false;
            _nw227[(new BigNumber(8)).toNumber()] = (_this).f25;
            _nw227[(new BigNumber(9)).toNumber()] = p1;
            _nw227[(new BigNumber(10)).toNumber()] = (_1252_v2).dtor_cf0;
            _nw227[(new BigNumber(11)).toNumber()] = false;
            _nw227[(new BigNumber(12)).toNumber()] = _this.f24;
            _nw227[(new BigNumber(13)).toNumber()] = (_this).f25;
            _nw227[(new BigNumber(14)).toNumber()] = (_1252_v2).dtor_cf0;
            _nw227[(new BigNumber(15)).toNumber()] = true;
            _nw227[(new BigNumber(16)).toNumber()] = p0;
            _nw227[(new BigNumber(17)).toNumber()] = p0;
            _nw227[(new BigNumber(18)).toNumber()] = p1;
            _nw227[(new BigNumber(19)).toNumber()] = p0;
            _nw227[(new BigNumber(20)).toNumber()] = (_this).f25;
            _nw227[(new BigNumber(21)).toNumber()] = true;
            _1253_v3 = _nw227;
            let _1256_v4;
            _1256_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1253_v3,(_this).f23);
            (globalState).f8 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(600)), function (_1257_i1) {
              return new _dafny.CodePoint('y'.codePointAt(0));
            })).length))).minus(new BigNumber((_1256_v4).length));
            let _1258_v5;
            let _nw228 = new _module.C6();
            _nw228.__ctor(_module.__default.fm24((_this).f23, (_this).f25, (_this).f25, globalState), _this.f24, _dafny.Seq.IsPrefixOf(_1250_v0, _dafny.Seq.UnicodeFromString("s")), _module.__default.fm1((_this).f23, true, globalState));
            _1258_v5 = _nw228;
            let _1259_v6;
            _1259_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_this.f24);
            _1259_v6 = (_1259_v6).update((_this).f23, p0);
            let _1260_v7;
            let _nw229 = new _module.C7();
            _nw229.__ctor();
            _1260_v7 = _nw229;
          }
        }
      }
      (globalState).f14 = new BigNumber(846);
      let _1261_v8;
      let _nw230 = new _module.C5();
      _nw230.__ctor(false, (_this).f25);
      _1261_v8 = _nw230;
      let _1262_v9;
      _1262_v9 = _dafny.Seq.UnicodeFromString("xwvnhkl");
      let _1263_v10;
      _1263_v10 = _dafny.Set.fromElements(new BigNumber((_1262_v9).length), (_this).f23);
      let _1264_v11;
      _1264_v11 = _dafny.Set.fromElements(_1263_v10, _1263_v10, _dafny.Set.fromElements((_this).f23));
      (globalState).f5 = new BigNumber((_1264_v11).length);
      let _1265_v12;
      _1265_v12 = _module.D0.create_DC0(_this.f24);
      let _1266_v13;
      _1266_v13 = _module.D0.create_DC2(_1265_v12);
      _1266_v13 = _1266_v13;
      r0 = new BigNumber((_dafny.Seq.UnicodeFromString("iasboupi")).length);
      let _1267_v14;
      let _nw231 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _1267_v14 = _nw231;
      r1 = (((_this).f25) ? (_1267_v14) : (_1267_v14));
      r2 = p1;
      return [r0, r1, r2];
    }
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = undefined;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _1268_v0;
      let _nw232 = Array((new BigNumber(15)).toNumber()).fill(false);
      _1268_v0 = _nw232;
      let _1269_v1;
      _1269_v1 = _dafny.Set.fromElements(_1268_v0);
      let _1270_v2;
      _1270_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(122),new BigNumber(-910));
      let _1271_v3;
      let _nw233 = new _module.C0();
      _nw233.__ctor(_1269_v1, ((_this).f23).plus(new BigNumber((_1270_v2).length)), (_this).f23);
      _1271_v3 = _nw233;
      let _1272_i0;
      _1272_i0 = _dafny.ZERO;
      L10: {
        while (!(_this.f24)) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1272_i0)) {
              break L10;
            }
            _1272_i0 = (_1272_i0).plus(_dafny.ONE);
            let _1273_v4;
            let _nw234 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
            _1273_v4 = _nw234;
            _1273_v4 = _1273_v4;
            let _1274_v5;
            let _nw235 = new _module.C7();
            _nw235.__ctor();
            _1274_v5 = _nw235;
            let _1275_v6;
            _1275_v6 = _dafny.Seq.of(_1271_v3.f31);
            let _1276_v7;
            _1276_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1274_v5,_1275_v6);
            let _1277_v8;
            _1277_v8 = _dafny.MultiSet.fromElements((_this).f25);
            (globalState).f14 = (((_this).f25) ? (new BigNumber((_dafny.Seq.Concat((((_1276_v7).contains(_1274_v5)) ? ((_1276_v7).get(_1274_v5)) : (_1275_v6)), _1275_v6)).length)) : (new BigNumber((_1277_v8).cardinality())));
            let _1278_v9;
            _1278_v9 = _dafny.Seq.of((_this).f25);
            let _1279_v10;
            let _nw236 = new _module.C5();
            _nw236.__ctor(!((_this).f25), (_1278_v9)[_module.__default.safeIndex(_1271_v3.f31, new BigNumber((_1278_v9).length))]);
            _1279_v10 = _nw236;
            _1279_v10 = _1279_v10;
            (globalState).f8 = _1271_v3.f31;
          }
        }
      }
      let _1280_v11;
      _1280_v11 = _dafny.MultiSet.fromElements((_this).f23);
      let _1281_v12;
      _1281_v12 = _dafny.Seq.of((_this).f25, (_this).f25, _this.f24);
      let _1282_v13;
      _1282_v13 = _dafny.Seq.of(true, (_1280_v11).equals(_1280_v11), (_1281_v12)[_module.__default.safeIndex(_1271_v3.f31, new BigNumber((_1281_v12).length))], !(_module.__default.fm29(_1271_v3.f31, (((_1280_v11).contains((_this).f23)) ? ((_1280_v11).get((_this).f23)) : ((_this).f23)), (_this).f25, globalState)));
      let _rhs225 = (_this).f23;
      let _rhs226 = (_1282_v13)[_module.__default.safeIndex(_module.__default.safeModuloInt((_this).f23, _1271_v3.f31), new BigNumber((_1282_v13).length))];
      let _lhs195 = globalState;
      let _lhs196 = _this;
      _lhs195.f5 = _rhs225;
      _lhs196.f24 = _rhs226;
      let _1283_v14;
      let _nw237 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _1283_v14 = _nw237;
      let _1284_v15;
      _1284_v15 = _module.D8.create_DC22(_1283_v14);
      let _source13 = _1284_v15;
      if (_source13.is_DC23) {
        let _1285___mcc_h0 = (_source13).cf40;
        let _1286___mcc_h1 = (_source13).cf41;
        let _1287___mcc_h2 = (_source13).cf42;
        let _1288___mcc_h3 = (_source13).cf43;
        let _1289___mcc_h4 = (_source13).cf44;
        let _1290_cf44 = _1289___mcc_h4;
        let _1291_cf43 = _1288___mcc_h3;
        let _1292_cf42 = _1287___mcc_h2;
        let _1293_cf41 = _1286___mcc_h1;
        let _1294_cf40 = _1285___mcc_h0;
        let _1295_v16;
        _1295_v16 = (_1271_v3).f30;
        let _1296_v17;
        _1296_v17 = _dafny.Seq.of(_1295_v16);
        (globalState).f5 = (new BigNumber(904)).minus(new BigNumber((_1296_v17).length));
        (globalState).f1 = _this.f24;
        if ((_this.f24) || (_this.f24)) {
          _1281_v12 = _1281_v12;
          let _1297_v19;
          _1297_v19 = _dafny.Seq.of(_1292_cf42);
          (globalState).f1 = (_module.D10.create_DC29((_this).f25, (_this).f25, new BigNumber((function () {
  let _coll44 = new _dafny.Map();
  for (const _compr_44 of (_1297_v19).Elements) {
    let _1298_v18 = _compr_44;
    if (_dafny.Seq.contains(_1297_v19, _1298_v18)) {
      _coll44.push([_module.__default.safeDivisionInt(_1298_v18, _1290_cf44),(_this).f25]);
    }
  }
  return _coll44;
}()).length))).dtor_cf54;
          let _1299_v20;
          _1299_v20 = _dafny.Seq.UnicodeFromString("wrebo");
          let _1300_v21;
          let _nw238 = new _module.C4();
          _nw238.__ctor(!(_this.f24), !(_this.f24), (_this).f25);
          _1300_v21 = _nw238;
          let _1301_v22;
          _1301_v22 = _module.D17.create_DC41(_1300_v21);
          let _1302_v23;
          _1302_v23 = _dafny.MultiSet.fromElements(_this.f24, _this.f24);
          let _rhs227 = _1293_cf41;
          let _rhs228 = _1299_v20;
          let _rhs229 = (((_this).f25) ? (_1300_v21) : ((_1301_v22).dtor_cf73));
          let _rhs230 = (_1302_v23).IsSubsetOf(_1302_v23);
          _1293_cf41 = _rhs227;
          _1299_v20 = _rhs228;
          _1300_v21 = _rhs229;
          r0 = _rhs230;
          let _1303_v24;
          _1303_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1280_v11,!((_this).f25));
          let _1304_v25;
          _1304_v25 = _dafny.Seq.of(_1297_v19);
          (globalState).f15 = (((((_1300_v21).f27) ? (_this.f24) : ((_this).f25))) ? ((_1300_v21).f27) : ((((_1303_v24).contains(((_dafny.MultiSet.FromArray((_1304_v25)[_module.__default.safeIndex(_1290_cf44, new BigNumber((_1304_v25).length))])).update(_1292_cf42, _module.__default.abs((_this).f23))).update(_1294_cf40, _module.__default.abs(_1271_v3.f31)))) ? ((_1303_v24).get(((_dafny.MultiSet.FromArray((_1304_v25)[_module.__default.safeIndex(_1290_cf44, new BigNumber((_1304_v25).length))])).update(_1292_cf42, _module.__default.abs((_this).f23))).update(_1294_cf40, _module.__default.abs(_1271_v3.f31)))) : (_this.f24))));
          let _1305_v26;
          let _nw239 = Array((new BigNumber(25)).toNumber());
          _1305_v26 = _nw239;
          let _1306_v27;
          let _nw240 = new _module.C4();
          _nw240.__ctor((_1300_v21).f27, (_this).f25, (_1300_v21).f27);
          _1306_v27 = _nw240;
          let _1307_v28;
          _1307_v28 = _dafny.Seq.of(_1306_v27);
          let _nw241 = Array((new BigNumber(21)).toNumber());
          _nw241[(_dafny.ZERO).toNumber()] = _1306_v27;
          _nw241[(_dafny.ONE).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(2)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(3)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(4)).toNumber()] = (((_this).fm3(_1299_v20, (_1306_v27).f25, _1271_v3.f31, _1297_v19, globalState)) ? (_1306_v27) : (_1306_v27));
          _nw241[(new BigNumber(5)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(6)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(7)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(8)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(9)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(10)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(11)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(12)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(13)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(14)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(15)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(16)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(17)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(18)).toNumber()] = _1306_v27;
          _nw241[(new BigNumber(19)).toNumber()] = (_1307_v28)[_module.__default.safeIndex(new BigNumber((_1293_cf41).length), new BigNumber((_1307_v28).length))];
          _nw241[(new BigNumber(20)).toNumber()] = _1306_v27;
          _1305_v26 = _nw241;
        } else {
          let _1308_v29;
          _1308_v29 = _dafny.Seq.UnicodeFromString("gigjdhnxs");
          let _1309_v30;
          _1309_v30 = new _dafny.CodePoint('p'.codePointAt(0));
          _1308_v29 = _dafny.Seq.Concat(_1308_v29, _dafny.Seq.update(_dafny.Seq.Concat(_1308_v29, _1308_v29), _module.__default.safeIndex(new BigNumber((_1281_v12).length), new BigNumber((_dafny.Seq.Concat(_1308_v29, _1308_v29)).length)), _1309_v30));
          _1293_cf41 = (function () {
            let _coll45 = new _dafny.Set();
            for (const _compr_45 of _dafny.IntegerRange(new BigNumber(-720), new BigNumber(674))) {
              let _1310_v31 = _compr_45;
              if (((new BigNumber(-720)).isLessThanOrEqualTo(_1310_v31)) && ((_1310_v31).isLessThan(new BigNumber(674)))) {
                _coll45.add((_1310_v31).plus(_1271_v3.f31));
              }
            }
            return _coll45;
          }()).Difference((_1293_cf41).Intersect(_1293_cf41));
          let _index179 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_1268_v0).length));
          (_1268_v0)[_index179] = (_this).f25;
          let _index180 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_1268_v0).length));
          (_1268_v0)[_index180] = !((_this).f25);
          let _1311_v32;
          _1311_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1271_v3.f31,(_this).f25);
          let _1312_v33;
          _1312_v33 = _dafny.Seq.of(_1271_v3.f31, new BigNumber(702), _1294_cf40, _1294_cf40);
          let _index181 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_1268_v0).length));
          let _index182 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_1268_v0).length));
          let _rhs231 = _this.f24;
          let _rhs232 = !(true);
          let _rhs233 = (_1293_cf41).IsSubsetOf(_1293_cf41);
          let _rhs234 = (_this).fm3(_1308_v29, !(_this.f24), new BigNumber((_1311_v32).length), _1312_v33, globalState);
          let _lhs197 = _1268_v0;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_1268_v0).length));
          let _lhs199 = _1268_v0;
          let _lhs200 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_1268_v0).length));
          let _lhs201 = globalState;
          let _lhs202 = globalState;
          _lhs197[_lhs198] = _rhs231;
          _lhs199[_lhs200] = _rhs232;
          _lhs201.f15 = _rhs233;
          _lhs202.f15 = _rhs234;
          let _1313_v34;
          let _init39 = ((_1314_v29) => function (_1315_i1) {
            return _1314_v29;
          })(_1308_v29);
          let _nw242 = Array((new BigNumber(15)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw242.length); _i0_39++) {
            _nw242[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1313_v34 = _nw242;
          let _index183 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1313_v34).length));
          (_1313_v34)[_index183] = _dafny.Seq.Concat(_1308_v29, _1308_v29);
          let _index184 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1313_v34).length));
          (_1313_v34)[_index184] = _1308_v29;
          let _1316_v35;
          _1316_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1283_v14,(_1313_v34)[_module.__default.safeIndex(new BigNumber(899), new BigNumber((_1313_v34).length))]);
          (globalState).f9 = (_1316_v35).contains(_1283_v14);
        }
        let _1317_v36;
        let _init40 = function (_1318_i2) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        };
        let _nw243 = Array((new BigNumber(11)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw243.length); _i0_40++) {
          _nw243[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1317_v36 = _nw243;
        let _index185 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_1317_v36).length));
        (_1317_v36)[_index185] = new _dafny.CodePoint('l'.codePointAt(0));
        let _1319_v37;
        _1319_v37 = new _dafny.CodePoint('m'.codePointAt(0));
        let _index186 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_1317_v36).length));
        (_1317_v36)[_index186] = ((false) ? (_1319_v37) : (_1319_v37));
      } else if (_source13.is_DC24) {
        let _1320___mcc_h5 = (_source13).cf45;
        let _1321___mcc_h6 = (_source13).cf46;
        let _1322___mcc_h7 = (_source13).cf47;
        let _1323___mcc_h8 = (_source13).cf48;
        let _1324___mcc_h9 = (_source13).cf49;
        let _1325_cf49 = _1324___mcc_h9;
        let _1326_cf48 = _1323___mcc_h8;
        let _1327_cf47 = _1322___mcc_h7;
        let _1328_cf46 = _1321___mcc_h6;
        let _1329_cf45 = _1320___mcc_h5;
        let _1330_v38;
        _1330_v38 = _dafny.Seq.UnicodeFromString("n");
        _1330_v38 = _1330_v38;
        let _1331_v39;
        let _init41 = function (_1332_i3) {
          return _module.D17.create_DC42();
        };
        let _nw244 = Array((new BigNumber(26)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw244.length); _i0_41++) {
          _nw244[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1331_v39 = _nw244;
        _1331_v39 = _1331_v39;
        let _1333_v40;
        _1333_v40 = _dafny.Seq.of((_this).f23);
        let _1334_v41;
        _1334_v41 = _module.D11.create_DC31(_1283_v14, _1325_cf49.f31, (_this).f25);
        let _rhs235 = _1333_v40;
        let _rhs236 = _1281_v12;
        let _rhs237 = _module.__default.fm25(!(_1326_cf48), ((_dafny.ZERO).minus(_1329_cf45)).isLessThan(_1325_cf49.f31), _1328_cf46, globalState);
        let _rhs238 = _dafny.Seq.IsPrefixOf(_module.__default.fm31((_this).f25, (_1334_v41).dtor_cf59, globalState), ((_1326_cf48) ? (_dafny.Seq.update(_1333_v40, _module.__default.safeIndex(_1329_cf45, new BigNumber((_1333_v40).length)), _1271_v3.f31)) : (_1333_v40)));
        let _lhs203 = globalState;
        _1333_v40 = _rhs235;
        _1281_v12 = _rhs236;
        _1330_v38 = _rhs237;
        _lhs203.f1 = _rhs238;
        (globalState).f8 = _1325_cf49.f31;
      } else {
        let _1335___mcc_h10 = (_source13).cf39;
        let _1336_cf39 = _1335___mcc_h10;
        if (_this.f24) {
          let _1337_v42;
          _1337_v42 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,_1280_v11);
          let _1338_v43;
          _1338_v43 = _dafny.Set.fromElements((_this).f25);
          let _1339_v44;
          _1339_v44 = _dafny.Seq.of(_1271_v3.f31);
          let _1340_v45;
          _1340_v45 = _module.D1.create_DC4(_this.f24, _this.f24);
          let _1341_v46;
          _1341_v46 = _dafny.Set.fromElements(_1340_v45, _1340_v45, _1340_v45);
          let _1342_v47;
          _1342_v47 = _dafny.Map.Empty.slice().updateUnsafe((_1338_v43).Union(_module.__default.fm32(_this.f24, _1339_v44, _1281_v12, _1341_v46, globalState)),_1280_v11);
          let _1343_v48;
          _1343_v48 = _dafny.Seq.of(_1338_v43);
          let _1344_v49;
          _1344_v49 = _dafny.Map.Empty.slice().updateUnsafe(_this.f24,new BigNumber(-195));
          let _rhs239 = (((_1342_v47).contains((_1343_v48)[_module.__default.safeIndex(_1271_v3.f31, new BigNumber((_1343_v48).length))])) ? ((_1342_v47).get((_1343_v48)[_module.__default.safeIndex(_1271_v3.f31, new BigNumber((_1343_v48).length))])) : (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_1344_v49).length)))));
          let _rhs240 = (_1339_v44)[_module.__default.safeIndex(_1271_v3.f31, new BigNumber((_1339_v44).length))];
          let _rhs241 = (_1271_v3.f31).minus(_1271_v3.f31);
          let _rhs242 = _1337_v42;
          let _lhs204 = globalState;
          _1280_v11 = _rhs239;
          _lhs204.f5 = _rhs240;
          r3 = _rhs241;
          _1337_v42 = _rhs242;
          let _1345_v50;
          _1345_v50 = new _dafny.CodePoint('v'.codePointAt(0));
          let _1346_v51;
          _1346_v51 = _module.D0.create_DC1(_this.f24, _1271_v3.f31);
          let _1347_v52;
          _1347_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1345_v50,!((_1346_v51).dtor_cf1));
          let _1348_v53;
          _1348_v53 = _module.D16.create_DC40((_this).f23, _1336_cf39, _1347_v52);
          let _1349_v54;
          _1349_v54 = _module.D0.create_DC2(_module.D0.create_DC0((_this).f25));
          let _1350_v55;
          _1350_v55 = _module.D0.create_DC2(_1349_v54);
          let _1351_v56;
          let _nw245 = new _module.C6();
          _nw245.__ctor(_1350_v55, true, _this.f24, _1271_v3.f31);
          _1351_v56 = _nw245;
          let _1352_v57;
          _1352_v57 = _dafny.MultiSet.fromElements(_1351_v56, _1351_v56, _1351_v56, _1351_v56);
          let _1353_v58;
          _1353_v58 = _dafny.Map.Empty.slice().updateUnsafe((_1348_v53).dtor_cf71,new BigNumber((_dafny.Seq.of(_1352_v57, _1352_v57, _1352_v57, _1352_v57)).length));
          let _1354_v59;
          _1354_v59 = _dafny.Seq.UnicodeFromString("qqj");
          (globalState).f5 = _module.__default.safeModuloInt(new BigNumber(((_1270_v2).update(_1271_v3.f31, (_this).f23)).length), (((_1353_v58).contains(_1336_cf39)) ? ((_1353_v58).get(_1336_cf39)) : (new BigNumber((_1354_v59).length))));
          let _1355_v60;
          let _nw246 = new _module.C1();
          _nw246.__ctor(_this.f24, _this.f24);
          _1355_v60 = _nw246;
          let _1356_v61;
          _1356_v61 = _dafny.Seq.of(_1268_v0);
          let _1357_v62;
          _1357_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1356_v61).length),(_this).f25);
          let _1358_v63;
          _1358_v63 = _dafny.Seq.of(_1357_v62, _1357_v62);
          let _1359_v64;
          _1359_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1355_v60,new BigNumber((_1358_v63).length));
          let _1360_v65;
          _1360_v65 = _dafny.Set.fromElements(new BigNumber((_1359_v64).length));
          let _1361_v66;
          _1361_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1360_v65,(_dafny.ZERO).minus((_this).f23));
          let _index187 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1268_v0).length));
          (_1268_v0)[_index187] = ((((_1361_v66).contains(_1360_v65)) ? ((_1361_v66).get(_1360_v65)) : ((_this).f23))).isLessThan(_1271_v3.f31);
          let _index188 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1268_v0).length));
          (_1268_v0)[_index188] = _this.f24;
          _1345_v50 = _1345_v50;
          (globalState).f15 = !(!((_this).f25));
        } else {
          let _index189 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1336_cf39).length));
          (_1336_cf39)[_index189] = (_1271_v3.f31).multipliedBy(new BigNumber(272));
          let _1362_v67;
          _1362_v67 = new _dafny.CodePoint('g'.codePointAt(0));
          let _1363_v68;
          _1363_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
          let _index190 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1336_cf39).length));
          (_1336_cf39)[_index190] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_this.f24,false)).Merge((_module.__default.fm10(_this.f24, _1271_v3.f31, _module.__default.fm12(false, globalState), _1362_v67, globalState)).Merge(_1363_v68))).length);
          let _1364_v69;
          let _nw247 = new _module.C5();
          _nw247.__ctor((new BigNumber(6)).isEqualTo(_1271_v3.f31), (_this).f25);
          _1364_v69 = _nw247;
          let _1365_v70;
          _1365_v70 = _dafny.Set.fromElements((_1281_v12)[_module.__default.safeIndex(_module.__default.fm1(_1271_v3.f31, (_this).f25, globalState), new BigNumber((_1281_v12).length))], !(_this.f24), _this.f24, (_this).f25);
          _1365_v70 = (_1365_v70).Intersect(_dafny.Set.fromElements((_this).f25, _this.f24, !(_this.f24)));
          let _1366_v71;
          _1366_v71 = _dafny.Set.fromElements(_1271_v3.f31, (_this).f23);
          (globalState).f15 = !(!((_1366_v71).Intersect(_1366_v71)).equals(_1366_v71));
          let _1367_v72;
          let _nw248 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
          _1367_v72 = _nw248;
          let _nw249 = Array((new BigNumber(4)).toNumber());
          _nw249[(_dafny.ZERO).toNumber()] = _1366_v71;
          _nw249[(_dafny.ONE).toNumber()] = function () {
            let _coll46 = new _dafny.Set();
            for (const _compr_46 of _dafny.IntegerRange(new BigNumber(988), new BigNumber(324))) {
              let _1368_v73 = _compr_46;
              if (((new BigNumber(988)).isLessThanOrEqualTo(_1368_v73)) && ((_1368_v73).isLessThan(new BigNumber(324)))) {
                _coll46.add(_module.__default.safeDivisionInt(_1368_v73, new BigNumber((_1365_v70).length)));
              }
            }
            return _coll46;
          }();
          _nw249[(new BigNumber(2)).toNumber()] = (_1366_v71).Difference(_1366_v71);
          _nw249[(new BigNumber(3)).toNumber()] = (_1366_v71).Union(_1366_v71);
          _1367_v72 = _nw249;
        }
        let _1369_v74;
        let _nw250 = new _module.C3();
        _nw250.__ctor((_this).f25, _1271_v3.f31, (_this).f25, _this.f24, _1271_v3.f31);
        _1369_v74 = _nw250;
        let _1370_v75;
        _1370_v75 = _dafny.Seq.of(_1369_v74, _1369_v74);
        let _1371_v76;
        _1371_v76 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.__default.fm29(new BigNumber((_1370_v75).length), _1271_v3.f31, _1369_v74.f24, globalState)),_1282_v13);
        let _index191 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_1268_v0).length));
        (_1268_v0)[_index191] = _dafny.Seq.IsPrefixOf((((_1371_v76).contains(_1281_v12)) ? ((_1371_v76).get(_1281_v12)) : (_1281_v12)), _1281_v12);
        let _index192 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_1268_v0).length));
        (_1268_v0)[_index192] = (_this).f25;
        (globalState).f6 = _this.f24;
        if ((_1369_v74).f25) {
          let _1372_v77;
          let _init42 = ((_1373_v74, _1374_v3) => function (_1375_i4) {
            return (((_1373_v74).f25) ? (_module.D10.create_DC29((_this).f25, !((_1373_v74).f25), _1374_v3.f31)) : (_module.D10.create_DC29(_1373_v74.f24, _1373_v74.f24, (_this).f23)));
          })(_1369_v74, _1271_v3);
          let _nw251 = Array((new BigNumber(28)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw251.length); _i0_42++) {
            _nw251[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1372_v77 = _nw251;
          let _1376_v78;
          _1376_v78 = _dafny.Seq.of(_1271_v3.f31);
          let _1377_v79;
          _1377_v79 = _dafny.Seq.UnicodeFromString("aawnfu");
          let _1378_v80;
          _1378_v80 = _dafny.MultiSet.fromElements(_1376_v78, _dafny.Seq.of(new BigNumber((_1377_v79).length)));
          let _1379_v81;
          _1379_v81 = _dafny.Set.fromElements(true, true, !((_this).f25));
          let _1380_v82;
          _1380_v82 = _module.D10.create_DC29((_this).fm6(_1271_v3.f31, _1378_v80, _this.f24, new BigNumber((_1379_v81).length), globalState), (_1369_v74).f25, (_this).f23);
          let _index193 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_1372_v77).length));
          (_1372_v77)[_index193] = _1380_v82;
          let _index194 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_1268_v0).length));
          let _index195 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_1372_v77).length));
          let _rhs243 = (_1380_v82).dtor_cf54;
          let _rhs244 = _1380_v82;
          let _lhs205 = _1268_v0;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_1268_v0).length));
          let _lhs207 = _1372_v77;
          let _lhs208 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_1372_v77).length));
          _lhs205[_lhs206] = _rhs243;
          _lhs207[_lhs208] = _rhs244;
          let _1381_v83;
          _1381_v83 = _dafny.MultiSet.fromElements((_1369_v74).f25);
          let _index196 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1283_v14).length));
          (_1283_v14)[_index196] = new BigNumber((_1381_v83).cardinality());
          let _index197 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1283_v14).length));
          let _rhs245 = _1377_v79;
          let _rhs246 = ((_this).f23).isEqualTo((_this).f23);
          let _rhs247 = (((_1381_v83).contains(!((_1369_v74).f25) || ((_this).f25))) ? ((_1381_v83).get(!((_1369_v74).f25) || ((_this).f25))) : (_1271_v3.f31));
          let _rhs248 = _1271_v3.f31;
          let _lhs209 = globalState;
          let _lhs210 = globalState;
          let _lhs211 = _1283_v14;
          let _lhs212 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1283_v14).length));
          _1377_v79 = _rhs245;
          _lhs209.f1 = _rhs246;
          _lhs210.f5 = _rhs247;
          _lhs211[_lhs212] = _rhs248;
          let _1382_v84;
          let _nw252 = new _module.C4();
          _nw252.__ctor((_1268_v0)[_module.__default.safeIndex(new BigNumber(66), new BigNumber((_1268_v0).length))], _1369_v74.f24, _this.f24);
          _1382_v84 = _nw252;
          let _rhs249 = ((!((_1381_v83).IsSubsetOf(_1381_v83))) ? ((_1280_v11).IsSubsetOf(_1280_v11)) : ((_1268_v0)[_module.__default.safeIndex(new BigNumber(66), new BigNumber((_1268_v0).length))]));
          let _rhs250 = (_1382_v84).f27;
          let _rhs251 = _dafny.Seq.Concat(_1281_v12, _1281_v12);
          let _rhs252 = _1271_v3.f31;
          let _lhs213 = globalState;
          let _lhs214 = globalState;
          let _lhs215 = globalState;
          _lhs213.f15 = _rhs249;
          _lhs214.f1 = _rhs250;
          _1281_v12 = _rhs251;
          _lhs215.f8 = _rhs252;
          let _rhs253 = _1336_cf39;
          let _rhs254 = ((_this).f25) || ((_this).f25);
          let _lhs216 = globalState;
          _1336_cf39 = _rhs253;
          _lhs216.f1 = _rhs254;
        } else {
          let _1383_v85;
          let _init43 = ((_1384_v11, _1385_v74) => function (_1386_i5) {
            return (_dafny.MultiSet.fromElements(_1384_v11, _1384_v11)).Difference(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1385_v74.f24)).length), new BigNumber(981))));
          })(_1280_v11, _1369_v74);
          let _nw253 = Array((new BigNumber(23)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw253.length); _i0_43++) {
            _nw253[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1383_v85 = _nw253;
          let _1387_v86;
          _1387_v86 = _dafny.MultiSet.fromElements(_1280_v11, _1280_v11);
          let _index198 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1383_v85).length));
          (_1383_v85)[_index198] = ((_1369_v74.f24) ? (_1387_v86) : (_1387_v86));
          let _index199 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_1383_v85).length));
          (_1383_v85)[_index199] = _1387_v86;
          _1270_v2 = _1270_v2;
          let _1388_v87;
          _1388_v87 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1389_v88;
          _1389_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1388_v87,(_this).f23);
          let _1390_v89;
          _1390_v89 = _dafny.Seq.of(_1389_v88);
          let _1391_v90;
          _1391_v90 = _module.D9.create_DC27(_module.D9.create_DC25(_1390_v89));
          let _1392_v91;
          _1392_v91 = _module.D9.create_DC25(_1390_v89);
          _1391_v90 = _module.D9.create_DC27(_1392_v91);
          let _1393_v92;
          let _nw254 = Array((new BigNumber(24)).toNumber());
          _1393_v92 = _nw254;
          _1393_v92 = _1393_v92;
          let _1394_v93;
          let _nw255 = Array((new BigNumber(2)).toNumber()).fill(_module.D6.Default());
          _1394_v93 = _nw255;
          let _init44 = ((_1395_v74) => function (_1396_i6) {
            return ((!(false)) ? (_module.D6.create_DC17((_1395_v74).f25)) : (_module.D6.create_DC17(_1395_v74.f24)));
          })(_1369_v74);
          let _nw256 = Array((new BigNumber(21)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw256.length); _i0_44++) {
            _nw256[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1394_v93 = _nw256;
        }
      }
      let _1397_v94;
      _1397_v94 = _dafny.Seq.UnicodeFromString("erevcsyu");
      _1397_v94 = _dafny.Seq.UnicodeFromString("s");
      if (_this.f24) {
        let _1398_v95;
        _1398_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(83),_1268_v0);
        let _1399_v96;
        let _nw257 = new _module.C3();
        _nw257.__ctor(_this.f24, new BigNumber((_1397_v94).length), _this.f24, (_this).f25, new BigNumber((_1398_v95).length));
        _1399_v96 = _nw257;
        let _1400_v97;
        _1400_v97 = _dafny.Map.Empty.slice().updateUnsafe((_1399_v96).f28,_1399_v96);
        let _rhs255 = _this.f24;
        let _rhs256 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_this.f24,_1399_v96)).Merge(_1400_v97)).length);
        let _rhs257 = !(_this.f24);
        let _lhs217 = globalState;
        let _lhs218 = globalState;
        let _lhs219 = globalState;
        _lhs217.f9 = _rhs255;
        _lhs218.f8 = _rhs256;
        _lhs219.f1 = _rhs257;
        let _1401_v98;
        _1401_v98 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1402_v99;
        _1402_v99 = _module.D0.create_DC1((_1399_v96).f28, new BigNumber((_module.__default.fm20((_1399_v96).f28, globalState)).length));
        let _1403_v100;
        _1403_v100 = _module.D0.create_DC2(_1402_v99);
        let _1404_v101;
        let _nw258 = new _module.C6();
        _nw258.__ctor(_1403_v100, (_1399_v96).f28, true, new BigNumber(219));
        _1404_v101 = _nw258;
        let _1405_v102;
        _1405_v102 = _dafny.MultiSet.fromElements(_1404_v101);
        let _pat_let_tv22 = _1402_v99;
        let _1406_v103;
        let _nw259 = Array((new BigNumber(6)).toNumber());
        _nw259[(_dafny.ZERO).toNumber()] = _module.D0.create_DC2(_1402_v99);
        _nw259[(_dafny.ONE).toNumber()] = _1403_v100;
        _nw259[(new BigNumber(2)).toNumber()] = _1403_v100;
        _nw259[(new BigNumber(3)).toNumber()] = _1403_v100;
        _nw259[(new BigNumber(4)).toNumber()] = _1403_v100;
        _nw259[(new BigNumber(5)).toNumber()] = function (_pat_let11_0) {
          return function (_1407_dt__update__tmp_h0) {
            return function (_pat_let12_0) {
              return function (_1408_dt__update_hcf3_h0) {
                return _module.D0.create_DC2(_1408_dt__update_hcf3_h0);
              }(_pat_let12_0);
            }(_pat_let_tv22);
          }(_pat_let11_0);
        }(_1404_v101.f26);
        _1406_v103 = _nw259;
        let _1409_v104;
        _1409_v104 = _dafny.Map.Empty.slice().updateUnsafe(_1406_v103,_1401_v98);
        let _1410_v105;
        _1410_v105 = _module.D3.create_DC9(new BigNumber(101));
        let _1411_v106;
        _1411_v106 = _dafny.Map.Empty.slice().updateUnsafe(_1410_v105,_1283_v14);
        let _1412_v107;
        _1412_v107 = _module.D5.create_DC15((((_1270_v2).contains(new BigNumber((_1405_v102).cardinality()))) ? ((_1270_v2).get(new BigNumber((_1405_v102).cardinality()))) : ((_this).f23)), _1401_v98, _1409_v104, _1411_v106);
        let _1413_v108;
        let _nw260 = Array((new BigNumber(18)).toNumber());
        _nw260[(_dafny.ZERO).toNumber()] = _1401_v98;
        _nw260[(_dafny.ONE).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(2)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(3)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(4)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(5)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(6)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(7)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
        _nw260[(new BigNumber(9)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(10)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(11)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
        _nw260[(new BigNumber(13)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(14)).toNumber()] = (_1412_v107).dtor_cf26;
        _nw260[(new BigNumber(15)).toNumber()] = _1401_v98;
        _nw260[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
        _nw260[(new BigNumber(17)).toNumber()] = _1401_v98;
        _1413_v108 = _nw260;
        let _1414_v109;
        _1414_v109 = _dafny.Set.fromElements(_1413_v108, _1413_v108);
        let _1415_v110;
        _1415_v110 = _module.D5.create_DC14(_1271_v3.f31, ((_this).f23).isLessThanOrEqualTo(new BigNumber((_1397_v94).length)), _1414_v109);
        let _1416_v111;
        _1416_v111 = _dafny.Seq.of((_this).f23, (_this).f23);
        let _pat_let_tv23 = _1416_v111;
        let _pat_let_tv24 = _1271_v3;
        let _pat_let_tv25 = _1416_v111;
        _1415_v110 = function (_pat_let13_0) {
          return function (_1417_dt__update__tmp_h1) {
            return function (_pat_let14_0) {
              return function (_1418_dt__update_hcf23_h0) {
                return function (_pat_let15_0) {
                  return function (_1419_dt__update_hcf22_h0) {
                    return _module.D5.create_DC14(_1419_dt__update_hcf22_h0, _1418_dt__update_hcf23_h0, (_1417_dt__update__tmp_h1).dtor_cf24);
                  }(_pat_let15_0);
                }((_pat_let_tv23)[_module.__default.safeIndex(_pat_let_tv24.f31, new BigNumber((_pat_let_tv25).length))]);
              }(_pat_let14_0);
            }(_this.f24);
          }(_pat_let13_0);
        }(_1415_v110);
        let _1420_v113;
        _1420_v113 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(999),(_1399_v96).f28);
        let _1421_v114;
        _1421_v114 = _dafny.Seq.of(_1270_v2, function () {
          let _coll47 = new _dafny.Map();
          for (const _compr_47 of (_1420_v113).Keys.Elements) {
            let _1422_v112 = _compr_47;
            if ((_1420_v113).contains(_1422_v112)) {
              _coll47.push([_module.__default.safeModuloInt(_1422_v112, (_1399_v96).f29),_1271_v3.f31]);
            }
          }
          return _coll47;
        }(), _1270_v2, _dafny.Map.Empty.slice().updateUnsafe(_1271_v3.f31,(_1399_v96).f29), _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_1271_v3.f31));
        _1421_v114 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_1270_v2), _module.__default.safeIndex((_1399_v96).f29, new BigNumber((_dafny.Seq.of(_1270_v2)).length)), _1270_v2), _dafny.Seq.Concat(_1421_v114, _1421_v114));
        let _1423_v115;
        let _nw261 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
        _1423_v115 = _nw261;
        let _index200 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_1423_v115).length));
        (_1423_v115)[_index200] = _1281_v12;
        let _index201 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_1423_v115).length));
        (_1423_v115)[_index201] = _dafny.Seq.of((_this).f25, !((_1281_v12)[_module.__default.safeIndex((_1399_v96).f29, new BigNumber((_1281_v12).length))]), (((_1399_v96).f28) ? ((_1399_v96).f28) : ((_1399_v96).f28)), (_1399_v96).f28);
        let _1424_v116;
        let _nw262 = Array((new BigNumber(3)).toNumber()).fill([]);
        _1424_v116 = _nw262;
        let _1425_v117;
        let _nw263 = new _module.C1();
        _nw263.__ctor((_1399_v96).f28, true);
        _1425_v117 = _nw263;
        let _1426_v118;
        _1426_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1425_v117,_1401_v98);
        let _index202 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_1283_v14).length));
        (_1283_v14)[_index202] = ((_1399_v96).f29).multipliedBy(_1271_v3.f31);
        let _index203 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_1283_v14).length));
        let _rhs258 = (_1399_v96).fm5(_1271_v3.f31, (_this).f23, globalState);
        let _rhs259 = _1424_v116;
        let _rhs260 = _1271_v3.f31;
        let _rhs261 = _1426_v118;
        let _rhs262 = _1271_v3.f31;
        let _lhs220 = globalState;
        let _lhs221 = globalState;
        let _lhs222 = _1283_v14;
        let _lhs223 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_1283_v14).length));
        _lhs220.f14 = _rhs258;
        _1424_v116 = _rhs259;
        _lhs221.f8 = _rhs260;
        _1426_v118 = _rhs261;
        _lhs222[_lhs223] = _rhs262;
      } else {
        let _1427_v119;
        _1427_v119 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1428_v120;
        _1428_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1271_v3.f31,_module.__default.fm11(new BigNumber((_dafny.Seq.UnicodeFromString("fdnmeq")).length), _this.f24, _this.f24, _1427_v119, globalState));
        let _1429_v121;
        _1429_v121 = _module.D1.create_DC4((_this).f25, (_this).f25);
        (_this).f24 = !((_1428_v120).update(_1271_v3.f31, _1429_v121)).equals(_1428_v120);
        let _1430_v122;
        _1430_v122 = _dafny.Seq.of(_1397_v94);
        let _1431_v123;
        _1431_v123 = _dafny.Seq.of(_1430_v122);
        let _rhs263 = _dafny.Seq.Concat(_1430_v122, (_1431_v123)[_module.__default.safeIndex(_1271_v3.f31, new BigNumber((_1431_v123).length))]);
        let _rhs264 = _1427_v119;
        _1430_v122 = _rhs263;
        _1427_v119 = _rhs264;
        let _1432_v124;
        let _nw264 = new _module.C7();
        _nw264.__ctor();
        _1432_v124 = _nw264;
        let _1433_v125;
        _1433_v125 = _dafny.MultiSet.fromElements(_1432_v124, _1432_v124);
        let _1434_v126;
        _1434_v126 = _dafny.Seq.of((_this).f23, (_1271_v3.f31).multipliedBy(new BigNumber((_1433_v125).cardinality())), ((!(_this.f24)) ? (_1271_v3.f31) : (_1271_v3.f31)), _1271_v3.f31, (_1271_v3.f31).plus((_this).f23));
        _1434_v126 = _dafny.Seq.of((_this).f23);
        if ((_1271_v3.f31).isLessThan(new BigNumber((_1280_v11).cardinality()))) {
          let _1435_v127;
          let _1436_v128;
          let _1437_v129;
          let _out14;
          let _out15;
          let _out16;
          let _outcollector3 = (_1432_v124).m2((_this).f23, ((_1280_v11).update(new BigNumber(745), _module.__default.abs(new BigNumber((_1270_v2).length)))).IsProperSubsetOf(_1280_v11), globalState);
          _out14 = _outcollector3[0];
          _out15 = _outcollector3[1];
          _out16 = _outcollector3[2];
          _1435_v127 = _out14;
          _1436_v128 = _out15;
          _1437_v129 = _out16;
          let _1438_v130;
          let _nw265 = new _module.C7();
          _nw265.__ctor();
          _1438_v130 = _nw265;
          (globalState).f1 = _this.f24;
          _1435_v127 = _this.f24;
          _1436_v128 = _this.f24;
        } else {
          let _index204 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_1283_v14).length));
          (_1283_v14)[_index204] = new BigNumber(((function () {
            let _coll48 = new _dafny.Map();
            for (const _compr_48 of _dafny.IntegerRange(new BigNumber(421), new BigNumber(182))) {
              let _1439_v131 = _compr_48;
              if (((new BigNumber(421)).isLessThanOrEqualTo(_1439_v131)) && ((_1439_v131).isLessThan(new BigNumber(182)))) {
                _coll48.push([_module.__default.safeDivisionInt(_1439_v131, new BigNumber((_1397_v94).length)),_this.f24]);
              }
            }
            return _coll48;
          }()).update(_1271_v3.f31, (_this).f25)).length);
          let _index205 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_1283_v14).length));
          (_1283_v14)[_index205] = _module.__default.safeModuloInt((_this).f23, (_this).f23);
          let _index206 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1268_v0).length));
          (_1268_v0)[_index206] = (_this).f25;
          let _index207 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1268_v0).length));
          (_1268_v0)[_index207] = _this.f24;
          let _1440_v132;
          let _nw266 = new _module.C3();
          _nw266.__ctor((_1268_v0)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_1268_v0).length))], _1271_v3.f31, (_1268_v0)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_1268_v0).length))], (_1268_v0)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_1268_v0).length))], _1271_v3.f31);
          _1440_v132 = _nw266;
          let _1441_v133;
          _1441_v133 = _dafny.Seq.of(_1440_v132);
          let _1442_v134;
          _1442_v134 = _dafny.Seq.of(_1441_v133, _1441_v133);
          let _1443_v135;
          let _nw267 = Array((new BigNumber(12)).toNumber());
          _nw267[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1441_v133, _1441_v133);
          _nw267[(_dafny.ONE).toNumber()] = _1441_v133;
          _nw267[(new BigNumber(2)).toNumber()] = (_1442_v134)[_module.__default.safeIndex((_this).f23, new BigNumber((_1442_v134).length))];
          _nw267[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1441_v133, _dafny.Seq.of(_1440_v132));
          _nw267[(new BigNumber(4)).toNumber()] = _1441_v133;
          _nw267[(new BigNumber(5)).toNumber()] = _1441_v133;
          _nw267[(new BigNumber(6)).toNumber()] = _1441_v133;
          _nw267[(new BigNumber(7)).toNumber()] = (((_1440_v132).f28) ? (_1441_v133) : (_dafny.Seq.of(_1440_v132)));
          _nw267[(new BigNumber(8)).toNumber()] = _1441_v133;
          _nw267[(new BigNumber(9)).toNumber()] = _1441_v133;
          _nw267[(new BigNumber(10)).toNumber()] = _1441_v133;
          _nw267[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_1441_v133, _module.__default.safeIndex(_1271_v3.f31, new BigNumber((_1441_v133).length)), _1440_v132);
          _1443_v135 = _nw267;
          let _1444_v136;
          let _nw268 = new _module.C1();
          _nw268.__ctor((_1440_v132).f28, !((_1440_v132).f28));
          _1444_v136 = _nw268;
          let _rhs265 = _1271_v3.f31;
          let _rhs266 = _1443_v135;
          let _rhs267 = ((_1271_v3.f31).minus(_1271_v3.f31)).multipliedBy(_1271_v3.f31);
          let _rhs268 = _1444_v136;
          let _lhs224 = globalState;
          let _lhs225 = globalState;
          _lhs224.f5 = _rhs265;
          _1443_v135 = _rhs266;
          _lhs225.f5 = _rhs267;
          _1444_v136 = _rhs268;
          let _1445_v137;
          let _nw269 = new _module.C5();
          _nw269.__ctor((_1268_v0)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_1268_v0).length))], (_1268_v0)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_1268_v0).length))]);
          _1445_v137 = _nw269;
          let _1446_v138;
          _1446_v138 = _dafny.Map.Empty.slice().updateUnsafe(_1445_v137,_1445_v137.f24);
          let _1447_v139;
          _1447_v139 = _dafny.Map.Empty.slice().updateUnsafe((_1268_v0)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_1268_v0).length))],_1446_v138);
          let _1448_v140;
          let _nw270 = Array((new BigNumber(3)).toNumber()).fill(_module.D0.Default());
          _1448_v140 = _nw270;
          let _1449_v141;
          _1449_v141 = _dafny.Map.Empty.slice().updateUnsafe(_1448_v140,_1427_v119);
          let _1450_v142;
          _1450_v142 = _module.D3.create_DC9((_1283_v14)[_module.__default.safeIndex(new BigNumber(461), new BigNumber((_1283_v14).length))]);
          let _1451_v143;
          let _nw271 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _1451_v143 = _nw271;
          let _1452_v144;
          _1452_v144 = _dafny.Map.Empty.slice().updateUnsafe(_1450_v142,_1451_v143);
          let _1453_v145;
          _1453_v145 = _module.D5.create_DC15(new BigNumber(((_1447_v139).update(false, _1446_v138)).length), _1427_v119, _1449_v141, (_1452_v144).Merge(_1452_v144));
          let _1454_v146;
          _1454_v146 = _dafny.Set.fromElements(_1427_v119, _1427_v119);
          _1453_v145 = (((_1281_v12)[_module.__default.safeIndex(new BigNumber((_1454_v146).length), new BigNumber((_1281_v12).length))]) ? (_1453_v145) : (_1453_v145));
          let _1455_v147;
          _1455_v147 = _module.D0.create_DC0((_1440_v132).f28);
          let _1456_v148;
          _1456_v148 = _module.D0.create_DC2(_1455_v147);
          let _1457_v149;
          _1457_v149 = _dafny.Map.Empty.slice().updateUnsafe((_1440_v132).f29,_this.f24);
          let _1458_v150;
          let _nw272 = new _module.C6();
          _nw272.__ctor((((_1268_v0)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_1268_v0).length))]) ? (_1456_v148) : (_module.D0.create_DC2(_module.D0.create_DC0((((_1457_v149).contains((_1440_v132).f29)) ? ((_1457_v149).get((_1440_v132).f29)) : (_this.f24)))))), (_1440_v132).f28, (_1445_v137).f25, _1271_v3.f31);
          _1458_v150 = _nw272;
        }
        (globalState).f15 = (_this).f25;
      }
      let _1459_v151;
      _1459_v151 = _dafny.Seq.of(new BigNumber(171), (_dafny.ZERO).minus((_dafny.ZERO).minus(_1271_v3.f31)), _1271_v3.f31, _1271_v3.f31, new BigNumber(443));
      r0 = (_1271_v3).fm3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-291)), function (_1460_i7) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }), _this.f24, new BigNumber((_1459_v151).length), ((!((_this).f25)) ? (_1459_v151) : (_dafny.Seq.update(_1459_v151, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-625)), new BigNumber((_1459_v151).length)), (_this).f23))), globalState);
      let _1461_v152;
      let _nw273 = new _module.C2();
      _nw273.__ctor((_dafny.ZERO).minus(_1271_v3.f31));
      _1461_v152 = _nw273;
      r1 = _1461_v152;
      r2 = (_this).f25;
      r3 = _module.__default.fm1((_this).fm2(globalState), false, globalState);
      return [r0, r1, r2, r3];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
