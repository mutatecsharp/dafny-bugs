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
      if ((new BigNumber((_dafny.Seq.of(!(!(false)))).length)).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-917)), function (_0_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length))) {
        return (new _dafny.CodePoint('h'.codePointAt(0)));
      } else {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }
    };
    static fm1(p0, p1, p2, globalState) {
      return true;
    };
    static fm2(p0, globalState) {
      return ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),true)).length)).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(490)))).cardinality()))).length))).length))).multipliedBy(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("wtycevgkj"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),(_dafny.Seq.UnicodeFromString("lqmpa"))))).length));
    };
    static fm3(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(229),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(838)), function (_1_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(294)), function (_2_i1) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(720)), function (_3_i2) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("auo")).length);
      })), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-51)), function (_4_i3) {
        return new BigNumber(318);
      })).length), new BigNumber((_dafny.Seq.of(true, !(!(!(true))), true)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(497)), function (_5_i4) {
        return new BigNumber(245);
      })));
    };
    static fm7(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("p"), _dafny.Seq.UnicodeFromString("rjwgaoud"));
    };
    static fm8(p0, p1, p2, globalState) {
      let _source0 = _module.D7.create_DC14(_module.D7.create_DC14(_module.D7.create_DC14(_module.D7.create_DC14(_module.D7.create_DC12(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(989), new BigNumber((function () {
  let _coll0 = new _dafny.Set();
  for (const _compr_0 of (function () {
    let _coll1 = new _dafny.Map();
    for (const _compr_1 of _dafny.IntegerRange(new BigNumber(46), new BigNumber(483))) {
      let _6_v0 = _compr_1;
      if (((new BigNumber(46)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(483)))) {
        _coll1.push([_module.__default.safeModuloInt(_6_v0, new BigNumber((function () {
          let _coll2 = new _dafny.Map();
          for (const _compr_2 of (_dafny.MultiSet.fromElements(new BigNumber(-312), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(220))).cardinality()))).Elements) {
            let _7_v1 = _compr_2;
            if ((_dafny.MultiSet.fromElements(new BigNumber(-312), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(220))).cardinality()))).contains(_7_v1)) {
              _coll2.push([_module.__default.safeDivisionInt(_7_v1, new BigNumber(-651)),new BigNumber(555)]);
            }
          }
          return _coll2;
        }()).length)),new BigNumber(899)]);
      }
    }
    return _coll1;
  }()).Keys.Elements) {
    let _8_v2 = _compr_0;
    if ((function () {
      let _coll3 = new _dafny.Map();
      for (const _compr_3 of _dafny.IntegerRange(new BigNumber(46), new BigNumber(483))) {
        let _6_v0 = _compr_3;
        if (((new BigNumber(46)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(483)))) {
          _coll3.push([_module.__default.safeModuloInt(_6_v0, new BigNumber((function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_dafny.MultiSet.fromElements(new BigNumber(-312), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(220))).cardinality()))).Elements) {
              let _7_v1 = _compr_4;
              if ((_dafny.MultiSet.fromElements(new BigNumber(-312), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(220))).cardinality()))).contains(_7_v1)) {
                _coll4.push([_module.__default.safeDivisionInt(_7_v1, new BigNumber(-651)),new BigNumber(555)]);
              }
            }
            return _coll4;
          }()).length)),new BigNumber(899)]);
        }
      }
      return _coll3;
    }()).contains(_8_v2)) {
      _coll0.add((_8_v2).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Set.fromElements(new BigNumber(899), new BigNumber(703), new BigNumber((_dafny.Seq.UnicodeFromString("llm")).length))).length))).length),new BigNumber(-706))).length)));
    }
  }
  return _coll0;
}()).length))))))));
      if (_source0.is_DC13) {
        let _9___mcc_h0 = (_source0).cf17;
        let _10_cf17 = _9___mcc_h0;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(37)), function (_11_i0) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        });
      } else if (_source0.is_DC12) {
        let _12___mcc_h1 = (_source0).cf16;
        let _13_cf16 = _12___mcc_h1;
        return _dafny.Seq.UnicodeFromString("y");
      } else {
        let _14___mcc_h2 = (_source0).cf18;
        let _15_cf18 = _14___mcc_h2;
        return _dafny.Seq.UnicodeFromString("hvccdvgrp");
      }
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(916),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-611))).length),false)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(991),!(true)));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      let _source1 = ((!(true)) ? (_module.D6.create_DC11(new BigNumber(-673))) : (_module.D6.create_DC11(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("gqh")).length))).length))));
      if (_source1.is_DC11) {
        let _16___mcc_h0 = (_source1).cf15;
        let _17_cf15 = _16___mcc_h0;
        return new _dafny.CodePoint('g'.codePointAt(0));
      } else {
        let _18___mcc_h1 = (_source1).cf14;
        let _19_cf14 = _18___mcc_h1;
        return new _dafny.CodePoint('r'.codePointAt(0));
      }
    };
    static fm11(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),new BigNumber(-514))).Keys.Elements) {
          let _20_v0 = _compr_5;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),new BigNumber(-514))).contains(_20_v0)) {
            _coll5.add(_20_v0);
          }
        }
        return _coll5;
      }()));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return (((false) ? (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(134)), function (_21_i0) {
        return new BigNumber(943);
      }))) : (_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(7), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(520),!(true))).length)), new BigNumber(496), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-47)))))))).Union(_dafny.MultiSet.FromArray(((!(true)) ? (_dafny.Seq.of(_dafny.Seq.of(new BigNumber(975)))) : (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(907)), function (_22_i1) {
        return (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-786)));
      }), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(811), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), function (_23_i2) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length), new BigNumber(948))).cardinality()), new BigNumber((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(481), new BigNumber(593))) {
            let _24_v1 = _compr_7;
            if (((new BigNumber(481)).isLessThanOrEqualTo(_24_v1)) && ((_24_v1).isLessThan(new BigNumber(593)))) {
              _coll7.push([_module.__default.safeModuloInt(_24_v1, new BigNumber(843)),new BigNumber((_dafny.Seq.of(false)).length)]);
            }
          }
          return _coll7;
        }()).Keys.Elements) {
          let _25_v0 = _compr_6;
          if ((function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of _dafny.IntegerRange(new BigNumber(481), new BigNumber(593))) {
              let _24_v1 = _compr_8;
              if (((new BigNumber(481)).isLessThanOrEqualTo(_24_v1)) && ((_24_v1).isLessThan(new BigNumber(593)))) {
                _coll8.push([_module.__default.safeModuloInt(_24_v1, new BigNumber(843)),new BigNumber((_dafny.Seq.of(false)).length)]);
              }
            }
            return _coll8;
          }()).contains(_25_v0)) {
            _coll6.push([_module.__default.safeDivisionInt(_25_v0, new BigNumber(-620)),_dafny.Seq.UnicodeFromString("vyeoofmh")]);
          }
        }
        return _coll6;
      }()).length), new BigNumber((_dafny.Seq.UnicodeFromString("lmjx")).length)))))));
    };
    static fm13(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(false, false));
    };
    static fm14(p0, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-173)), function (_26_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(!(false), true)).length));
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("vetmckar")).length))).length))).length))), _dafny.Seq.of(new BigNumber((function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.fromElements(true))).Elements) {
          let _27_v0 = _compr_9;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.fromElements(true)), _27_v0)) {
            _coll9.push([_27_v0,new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)))).length)]);
          }
        }
        return _coll9;
      }()).length), new BigNumber((_dafny.Set.fromElements(!(!(!(true))))).length), new BigNumber(670)), _dafny.Seq.of(new BigNumber(818), new BigNumber(((_module.D10.create_DC21(_dafny.Seq.of(!(true)))).dtor_cf33).length), new BigNumber(699)));
    };
    static fm15(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-782),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-698))).cardinality()))).cardinality()),!(true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-381),!(false)));
    };
    static fm16(globalState) {
      return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false, true))).length));
    };
    static fm17(p0, p1, p2, globalState) {
      if ((_dafny.Set.fromElements(_dafny.Seq.of(false))).IsDisjointFrom(_dafny.Set.fromElements(_dafny.Seq.of(false, true)))) {
        if (true) {
          return _module.D1.create_DC3(_dafny.MultiSet.fromElements(!(false), false, false), false, true);
        } else {
          return _module.D1.create_DC3(_dafny.MultiSet.fromElements(true), false, false);
        }
      } else {
        return _module.D1.create_DC3(_dafny.MultiSet.fromElements(!(!(true)), true), !((_module.D7.create_DC13(true)).dtor_cf17), true);
      }
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Set.Empty;
      let _28_v0;
      _28_v0 = new BigNumber(232);
      let _29_v1;
      _29_v1 = _dafny.Seq.of(_28_v0);
      let _30_v2;
      _30_v2 = _module.D8.create_DC15(_29_v1);
      _29_v1 = (_30_v2).dtor_cf19;
      let _31_v3;
      _31_v3 = _module.D7.create_DC13(p1);
      let _32_v4;
      _32_v4 = _module.D7.create_DC14(_31_v3);
      let _33_v5;
      _33_v5 = _dafny.Map.Empty.slice().updateUnsafe(_32_v4,p1);
      _33_v5 = (_33_v5).update(_32_v4, false);
      let _34_v6;
      _34_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,_28_v0);
      let _35_v7;
      _35_v7 = _dafny.Set.fromElements(_34_v6, _34_v6);
      let _36_i0;
      _36_i0 = _dafny.ZERO;
      L0: {
        while ((_35_v7).IsProperSubsetOf(function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of (_35_v7).Elements) {
            let _53_v8 = _compr_10;
            if ((_35_v7).contains(_53_v8)) {
              _coll10.add(_53_v8);
            }
          }
          return _coll10;
        }())) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_36_i0)) {
              break L0;
            }
            _36_i0 = (_36_i0).plus(_dafny.ONE);
            let _37_v9;
            let _init0 = ((_38_v0) => function (_39_i1) {
              return (_39_i1).multipliedBy(_38_v0);
            })(_28_v0);
            let _nw0 = Array((new BigNumber(8)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
              _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _37_v9 = _nw0;
            _37_v9 = (_module.D1.create_DC1(_37_v9)).dtor_cf1;
            let _40_v10;
            let _init1 = ((_41_v1, _42_v0, _43_p1) => function (_44_i2) {
              return (_dafny.MultiSet.fromElements(_41_v1)).Difference(_dafny.MultiSet.fromElements(_41_v1, _dafny.Seq.update(_dafny.Seq.of(_42_v0, new BigNumber((_41_v1).length)), _module.__default.safeIndex(_42_v0, new BigNumber((_dafny.Seq.of(_42_v0, new BigNumber((_41_v1).length))).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_43_p1)).length)))));
            })(_29_v1, _28_v0, p1);
            let _nw1 = Array((new BigNumber(18)).toNumber());
            for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw1.length); _i0_1++) {
              _nw1[_i0_1] = _init1(new BigNumber(_i0_1));
            }
            _40_v10 = _nw1;
            let _45_v11;
            _45_v11 = _dafny.Seq.of(_29_v1);
            let _index0 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_40_v10).length));
            (_40_v10)[_index0] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(_28_v0)), _45_v11));
            let _46_v12;
            let _nw2 = new _module.C0();
            _nw2.__ctor(p1);
            _46_v12 = _nw2;
            let _47_v13;
            _47_v13 = _dafny.Map.Empty.slice().updateUnsafe(_46_v12,_29_v1);
            let _48_v14;
            _48_v14 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_28_v0, _28_v0), (((_47_v13).contains(_46_v12)) ? ((_47_v13).get(_46_v12)) : (_29_v1)));
            let _index1 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_40_v10).length));
            (_40_v10)[_index1] = (_module.D7.create_DC12(_48_v14)).dtor_cf16;
            let _49_v15;
            _49_v15 = false;
            let _50_v16;
            _50_v16 = new _dafny.CodePoint('r'.codePointAt(0));
            let _51_v17;
            _51_v17 = _50_v16;
            let _52_v18;
            _52_v18 = _dafny.Seq.of((_51_v17));
            let _rhs0 = !_dafny.areEqual(_module.__default.fm7(new BigNumber((_52_v18).length), _28_v0, globalState), ((p1) ? (_dafny.Seq.UnicodeFromString("wrtno")) : (_52_v18)));
            let _rhs1 = _46_v12;
            _49_v15 = _rhs0;
            _46_v12 = _rhs1;
            _29_v1 = _dafny.Seq.Concat(_dafny.Seq.of(_28_v0), _29_v1);
          }
        }
      }
      let _54_v19;
      _54_v19 = true;
      _54_v19 = p1;
      let _55_i3;
      _55_i3 = _dafny.ZERO;
      L1: {
        while (p1) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_55_i3)) {
              break L1;
            }
            _55_i3 = (_55_i3).plus(_dafny.ONE);
            _54_v19 = p1;
            r0 = _28_v0;
            let _56_v20;
            _56_v20 = _dafny.Seq.of(false);
            let _57_v21;
            _57_v21 = _dafny.Seq.of(_32_v4, _module.D7.create_DC14(_31_v3));
            let _58_v22;
            _58_v22 = _dafny.Map.Empty.slice().updateUnsafe(_28_v0,_28_v0);
            let _59_v23;
            _59_v23 = _module.D8.create_DC16(_54_v19, _dafny.MultiSet.FromArray(_56_v20), _57_v21, _58_v22);
            let _60_v24;
            _60_v24 = _module.D8.create_DC18(_59_v23);
            let _61_v25;
            _61_v25 = _module.D8.create_DC18(_59_v23);
            let _62_v26;
            _62_v26 = _module.D8.create_DC18(_59_v23);
            let _63_v27;
            _63_v27 = _module.D8.create_DC18(_62_v26);
            let _64_v28;
            _64_v28 = _module.D8.create_DC18(_62_v26);
            let _source2 = _64_v28;
            if (_source2.is_DC16) {
              let _65___mcc_h0 = (_source2).cf20;
              let _66___mcc_h1 = (_source2).cf21;
              let _67___mcc_h2 = (_source2).cf22;
              let _68___mcc_h3 = (_source2).cf23;
              let _69_cf23 = _68___mcc_h3;
              let _70_cf22 = _67___mcc_h2;
              let _71_cf21 = _66___mcc_h1;
              let _72_cf20 = _65___mcc_h0;
              let _73_v29;
              let _nw3 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
              _73_v29 = _nw3;
              let _index2 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_73_v29).length));
              (_73_v29)[_index2] = _module.__default.fm15(globalState);
              let _74_v30;
              _74_v30 = _dafny.Map.Empty.slice().updateUnsafe(_28_v0,_72_cf20);
              let _75_v31;
              _75_v31 = _dafny.Seq.of(_74_v30, _74_v30);
              let _76_v32;
              _76_v32 = _dafny.Seq.UnicodeFromString("nixsur");
              let _77_v33;
              _77_v33 = _dafny.Map.Empty.slice().updateUnsafe(_76_v32,_28_v0);
              let _78_v34;
              _78_v34 = _dafny.Seq.of(_77_v33, _77_v33);
              let _79_v35;
              _79_v35 = _dafny.Set.fromElements(_54_v19, _54_v19, p1);
              let _index3 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_73_v29).length));
              (_73_v29)[_index3] = ((_75_v31)[_module.__default.safeIndex(new BigNumber(((_78_v34)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_28_v0, _28_v0)).length), new BigNumber((_78_v34).length))]).length), new BigNumber((_75_v31).length))]).Merge(((_74_v30).update((_dafny.ZERO).minus(_28_v0), p1)).update(_28_v0, _module.__default.fm1(_dafny.Seq.UnicodeFromString("hmgxqq"), _28_v0, _79_v35, globalState)));
              let _80_v36;
              let _nw4 = new _module.C0();
              _nw4.__ctor(p1);
              _80_v36 = _nw4;
              _80_v36 = _80_v36;
              let _81_v37;
              _81_v37 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_74_v30).update(_28_v0, _72_cf20),_54_v19)).length)));
              let _82_v38;
              let _nw5 = Array((new BigNumber(4)).toNumber());
              _nw5[(_dafny.ZERO).toNumber()] = _81_v37;
              _nw5[(_dafny.ONE).toNumber()] = _81_v37;
              _nw5[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_28_v0, _28_v0);
              _nw5[(new BigNumber(3)).toNumber()] = _81_v37;
              _82_v38 = _nw5;
              let _index4 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_82_v38).length));
              (_82_v38)[_index4] = _81_v37;
              let _83_v39;
              let _init2 = ((_84_v0) => function (_85_i4) {
                return _module.__default.safeDivisionInt(_85_i4, _84_v0);
              })(_28_v0);
              let _nw6 = Array((new BigNumber(22)).toNumber());
              for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw6.length); _i0_2++) {
                _nw6[_i0_2] = _init2(new BigNumber(_i0_2));
              }
              _83_v39 = _nw6;
              let _index5 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_83_v39).length));
              (_83_v39)[_index5] = _module.__default.safeModuloInt(_28_v0, new BigNumber((_58_v22).length));
              let _index6 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_83_v39).length));
              (_83_v39)[_index6] = new BigNumber((_29_v1).length);
              let _86_v40;
              let _nw7 = new _module.C0();
              _nw7.__ctor(p1);
              _86_v40 = _nw7;
              let _87_v41;
              _87_v41 = _dafny.Seq.of(_86_v40);
              let _88_v42;
              _88_v42 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_87_v41)[_module.__default.safeIndex(_28_v0, new BigNumber((_87_v41).length))]), _module.__default.safeIndex(new BigNumber((_87_v41).length), new BigNumber((_dafny.Seq.of((_87_v41)[_module.__default.safeIndex(_28_v0, new BigNumber((_87_v41).length))])).length)), _86_v40), _87_v41);
              let _index7 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_82_v38).length));
              let _index8 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_83_v39).length));
              let _index9 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_83_v39).length));
              let _rhs2 = ((_81_v37).Difference(_module.__default.fm16(globalState))).Intersect(_81_v37);
              let _rhs3 = _76_v32;
              let _rhs4 = (_29_v1)[_module.__default.safeIndex(_28_v0, new BigNumber((_29_v1).length))];
              let _rhs5 = (_28_v0).multipliedBy(new BigNumber((_76_v32).length));
              let _rhs6 = _88_v42;
              let _lhs0 = _82_v38;
              let _lhs1 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_82_v38).length));
              let _lhs2 = _83_v39;
              let _lhs3 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_83_v39).length));
              let _lhs4 = _83_v39;
              let _lhs5 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_83_v39).length));
              _lhs0[_lhs1] = _rhs2;
              _76_v32 = _rhs3;
              _lhs2[_lhs3] = _rhs4;
              _lhs4[_lhs5] = _rhs5;
              _88_v42 = _rhs6;
              let _89_v43;
              _89_v43 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_28_v0, (_83_v39)[_module.__default.safeIndex(new BigNumber(833), new BigNumber((_83_v39).length))]), _29_v1, _29_v1, _29_v1);
              let _90_v44;
              _90_v44 = _module.D7.create_DC12(_89_v43);
              let _pat_let_tv0 = _89_v43;
              let _rhs7 = ((_80_v36.f0) ? (_72_cf20) : (_80_v36.f0));
              let _rhs8 = (_83_v39)[_module.__default.safeIndex(new BigNumber(833), new BigNumber((_83_v39).length))];
              let _rhs9 = function (_pat_let0_0) {
                return function (_91_dt__update__tmp_h0) {
                  return function (_pat_let1_0) {
                    return function (_92_dt__update_hcf16_h0) {
                      return _module.D7.create_DC12(_92_dt__update_hcf16_h0);
                    }(_pat_let1_0);
                  }(_pat_let_tv0);
                }(_pat_let0_0);
              }(_90_v44);
              _72_cf20 = _rhs7;
              r2 = _rhs8;
              _90_v44 = _rhs9;
            } else if (_source2.is_DC17) {
              let _93___mcc_h4 = (_source2).cf24;
              let _94___mcc_h5 = (_source2).cf25;
              let _95___mcc_h6 = (_source2).cf26;
              let _96___mcc_h7 = (_source2).cf27;
              let _97_cf27 = _96___mcc_h7;
              let _98_cf26 = _95___mcc_h6;
              let _99_cf25 = _94___mcc_h5;
              let _100_cf24 = _93___mcc_h4;
              let _101_v45;
              _101_v45 = new _dafny.CodePoint('f'.codePointAt(0));
              let _102_v46;
              _102_v46 = _module.D3.create_DC6(_28_v0, _98_cf26, p1);
              let _103_v47;
              _103_v47 = _dafny.Set.fromElements((_102_v46).dtor_cf10, p1, p1);
              let _104_v48;
              _104_v48 = _dafny.Map.Empty.slice().updateUnsafe((_29_v1)[_module.__default.safeIndex(_99_cf25, new BigNumber((_29_v1).length))],_module.__default.fm1(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("efgeywy"), _module.__default.safeIndex(_99_cf25, new BigNumber((_dafny.Seq.UnicodeFromString("efgeywy")).length)), _101_v45), _28_v0, _103_v47, globalState));
              _104_v48 = (_104_v48).update(_97_cf27, _54_v19);
              _54_v19 = (_99_cf25).isLessThanOrEqualTo((_29_v1)[_module.__default.safeIndex((_dafny.ZERO).minus((_module.D6.create_DC11(_97_cf27)).dtor_cf15), new BigNumber((_29_v1).length))]);
              let _105_v49;
              let _nw8 = new _module.C0();
              _nw8.__ctor(_54_v19);
              _105_v49 = _nw8;
              _105_v49 = _105_v49;
              let _106_v50;
              let _nw9 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
              _106_v50 = _nw9;
              _106_v50 = _106_v50;
            } else if (_source2.is_DC15) {
              let _107___mcc_h8 = (_source2).cf19;
              let _108_cf19 = _107___mcc_h8;
              let _109_v51;
              _109_v51 = _28_v0;
              let _110_v52;
              _110_v52 = _dafny.Map.Empty.slice().updateUnsafe(_54_v19,_109_v51);
              let _111_v53;
              _111_v53 = _dafny.Map.Empty.slice().updateUnsafe(p1,_54_v19);
              let _112_v54;
              _112_v54 = _dafny.Map.Empty.slice().updateUnsafe(_110_v52,(((_111_v53).contains(_54_v19)) ? ((_111_v53).get(_54_v19)) : (p1)));
              _112_v54 = (_112_v54).update(_110_v52, p1);
              _54_v19 = (((_111_v53).contains(_54_v19)) ? ((_111_v53).get(_54_v19)) : (!(!(p1)) || (p1)));
              _54_v19 = p1;
              let _113_v55;
              _113_v55 = _dafny.Set.fromElements(_54_v19, _54_v19);
              _54_v19 = (_113_v55).IsProperSubsetOf((_113_v55).Intersect(_113_v55));
            } else {
              let _114___mcc_h9 = (_source2).cf28;
              let _115_cf28 = _114___mcc_h9;
              let _116_v56;
              _116_v56 = _dafny.MultiSet.fromElements((_56_v20)[_module.__default.safeIndex(_28_v0, new BigNumber((_56_v20).length))], p1);
              let _117_v57;
              _117_v57 = _dafny.Set.fromElements(p1, p1);
              _54_v19 = !(_module.__default.fm1(_module.__default.fm7((((_116_v56).contains(_54_v19)) ? ((_116_v56).get(_54_v19)) : (_28_v0)), _28_v0, globalState), _28_v0, _117_v57, globalState));
              _29_v1 = _dafny.Seq.update(((p1) ? (_29_v1) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(907)), ((_118_v0) => function (_119_i5) {
                return _118_v0;
              })(_28_v0)))), _module.__default.safeIndex(_28_v0, new BigNumber((((p1) ? (_29_v1) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(907)), ((_120_v0) => function (_121_i5) {
                return _120_v0;
              })(_28_v0))))).length)), _28_v0);
              let _122_v58;
              let _nw10 = new _module.C0();
              _nw10.__ctor(p1);
              _122_v58 = _nw10;
              let _123_v59;
              _123_v59 = _dafny.Map.Empty.slice().updateUnsafe(_122_v58,p1);
              let _124_v60;
              _124_v60 = _dafny.Map.Empty.slice().updateUnsafe(_123_v59,p1);
              let _125_v61;
              let _nw11 = new _module.C0();
              _nw11.__ctor((((_124_v60).contains(_123_v59)) ? ((_124_v60).get(_123_v59)) : (p1)));
              _125_v61 = _nw11;
              let _126_v62;
              _126_v62 = _dafny.Seq.UnicodeFromString("owkgj");
              let _127_v63;
              _127_v63 = _dafny.Map.Empty.slice().updateUnsafe(_28_v0,_126_v62);
              let _128_v64;
              _128_v64 = new _dafny.CodePoint('d'.codePointAt(0));
              let _129_v65;
              _129_v65 = _dafny.Map.Empty.slice().updateUnsafe(_54_v19,_128_v64);
              let _130_v66;
              _130_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((((_54_v19) ? ((((_127_v63).contains(_28_v0)) ? ((_127_v63).get(_28_v0)) : (_126_v62))) : (_dafny.Seq.update(_126_v62, _module.__default.safeIndex(_28_v0, new BigNumber((_126_v62).length)), (((_129_v65).contains(p1)) ? ((_129_v65).get(p1)) : (_128_v64)))))).length),_125_v61);
              let _131_v67;
              _131_v67 = _module.D6.create_DC10(_125_v61);
              _125_v61 = (((_130_v66).contains(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_56_v20, _module.__default.safeIndex(new BigNumber(302), new BigNumber((_56_v20).length)), _54_v19), _dafny.Seq.of(_54_v19, p1, true, _54_v19))).length))) ? ((_130_v66).get(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_56_v20, _module.__default.safeIndex(new BigNumber(302), new BigNumber((_56_v20).length)), _54_v19), _dafny.Seq.of(_54_v19, p1, true, _54_v19))).length))) : ((_131_v67).dtor_cf14));
              _54_v19 = p1;
            }
            let _132_v68;
            _132_v68 = _dafny.Seq.UnicodeFromString("emnf");
            if (_module.__default.fm1(_132_v68, _28_v0, _module.__default.fm13(_54_v19, _54_v19, p1, globalState), globalState)) {
              _54_v19 = _54_v19;
              let _133_v69;
              _133_v69 = _dafny.Set.fromElements(_54_v19, _54_v19);
              let _134_v70;
              let _nw12 = new _module.C0();
              _nw12.__ctor((p1) || (_module.__default.fm1(_132_v68, new BigNumber((_dafny.Seq.UnicodeFromString("ertqblss")).length), _133_v69, globalState)));
              _134_v70 = _nw12;
              let _135_v71;
              let _nw13 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
              _135_v71 = _nw13;
              let _index10 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_135_v71).length));
              (_135_v71)[_index10] = _28_v0;
              let _index11 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_135_v71).length));
              let _rhs10 = _dafny.Seq.IsPrefixOf(_29_v1, _29_v1);
              let _rhs11 = _132_v68;
              let _rhs12 = _28_v0;
              let _lhs6 = _135_v71;
              let _lhs7 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_135_v71).length));
              _54_v19 = _rhs10;
              _132_v68 = _rhs11;
              _lhs6[_lhs7] = _rhs12;
              let _136_v72;
              let _init3 = function (_137_i6) {
                return false;
              };
              let _nw14 = Array((new BigNumber(22)).toNumber());
              for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw14.length); _i0_3++) {
                _nw14[_i0_3] = _init3(new BigNumber(_i0_3));
              }
              _136_v72 = _nw14;
              let _index12 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_136_v72).length));
              (_136_v72)[_index12] = p1;
              let _index13 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_136_v72).length));
              (_136_v72)[_index13] = p1;
              let _138_v73;
              let _nw15 = new _module.C0();
              _nw15.__ctor((_module.__default.fm17(_28_v0, new BigNumber(755), p1, globalState)).dtor_cf5);
              _138_v73 = _nw15;
            } else {
              let _139_v74;
              let _nw16 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
              _139_v74 = _nw16;
              let _index14 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_139_v74).length));
              (_139_v74)[_index14] = _28_v0;
              let _140_v75;
              _140_v75 = new _dafny.CodePoint('d'.codePointAt(0));
              let _141_v76;
              _141_v76 = _dafny.Seq.of(_132_v68);
              let _142_v77;
              _142_v77 = _28_v0;
              let _143_v78;
              _143_v78 = _dafny.Map.Empty.slice().updateUnsafe(_28_v0,_141_v76);
              let _144_v79;
              _144_v79 = _dafny.Seq.of((((_143_v78).contains(new BigNumber(-272))) ? ((_143_v78).get(new BigNumber(-272))) : (_dafny.Seq.update(_141_v76, _module.__default.safeIndex(_28_v0, new BigNumber((_141_v76).length)), _132_v68))));
              let _index15 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_139_v74).length));
              let _rhs13 = (_142_v77);
              let _rhs14 = _140_v75;
              let _rhs15 = (_132_v68)[_module.__default.safeIndex(_28_v0, new BigNumber((_132_v68).length))];
              let _rhs16 = _dafny.Seq.Concat((_144_v79)[_module.__default.safeIndex(_module.__default.fm2(p1, globalState), new BigNumber((_144_v79).length))], _141_v76);
              let _rhs17 = _module.__default.safeDivisionInt(_28_v0, (_29_v1)[_module.__default.safeIndex(_28_v0, new BigNumber((_29_v1).length))]);
              let _lhs8 = _139_v74;
              let _lhs9 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_139_v74).length));
              _lhs8[_lhs9] = _rhs13;
              _140_v75 = _rhs14;
              _140_v75 = _rhs15;
              _141_v76 = _rhs16;
              _28_v0 = _rhs17;
              let _145_v80;
              let _nw17 = Array((new BigNumber(5)).toNumber());
              _145_v80 = _nw17;
              let _nw18 = Array((new BigNumber(7)).toNumber());
              _145_v80 = _nw18;
              let _146_v82;
              _146_v82 = _module.D9.create_DC19(function () {
  let _coll11 = new _dafny.Map();
  for (const _compr_11 of _dafny.IntegerRange(new BigNumber(-359), new BigNumber(203))) {
    let _147_v81 = _compr_11;
    if (((new BigNumber(-359)).isLessThanOrEqualTo(_147_v81)) && ((_147_v81).isLessThan(new BigNumber(203)))) {
      _coll11.push([_module.__default.safeModuloInt(_147_v81, _28_v0),_56_v20]);
    }
  }
  return _coll11;
}());
              let _index16 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_139_v74).length));
              (_139_v74)[_index16] = new BigNumber(((_146_v82).dtor_cf29).length);
              _54_v19 = _54_v19;
              let _148_v83;
              let _nw19 = new _module.C0();
              _nw19.__ctor(_54_v19);
              _148_v83 = _nw19;
            }
          }
        }
      }
      let _149_i7;
      _149_i7 = _dafny.ZERO;
      L2: {
        while (!(!(_54_v19))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_149_i7)) {
              break L2;
            }
            _149_i7 = (_149_i7).plus(_dafny.ONE);
            let _150_v84;
            _150_v84 = _dafny.Seq.of(!(p1), _54_v19);
            let _151_v85;
            _151_v85 = _dafny.MultiSet.fromElements((_150_v84)[_module.__default.safeIndex(_28_v0, new BigNumber((_150_v84).length))]);
            let _152_v86;
            _152_v86 = _module.D1.create_DC3(_151_v85, _54_v19, _54_v19);
            let _153_v87;
            _153_v87 = _dafny.Set.fromElements(p1, _54_v19, p1, (_152_v86).dtor_cf4);
            _54_v19 = _module.__default.fm1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(253)), function (_154_i8) {
              return new _dafny.CodePoint('l'.codePointAt(0));
            }), _28_v0, _153_v87, globalState);
            if ((new BigNumber(448)).isEqualTo(_28_v0)) {
              let _155_v88;
              _155_v88 = _dafny.Seq.UnicodeFromString("nlwntcgy");
              let _156_v89;
              let _nw20 = new _module.C0();
              _nw20.__ctor(p1);
              _156_v89 = _nw20;
              let _157_v90;
              _157_v90 = new _dafny.CodePoint('m'.codePointAt(0));
              _155_v88 = _dafny.Seq.update(_155_v88, _module.__default.safeIndex((_28_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements(_156_v89)).length)), new BigNumber((_155_v88).length)), ((p1) ? (_157_v90) : (new _dafny.CodePoint('j'.codePointAt(0)))));
              let _158_v91;
              let _init4 = ((_159_v2) => function (_160_i9) {
                return _159_v2;
              })(_30_v2);
              let _nw21 = Array((new BigNumber(19)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw21.length); _i0_4++) {
                _nw21[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _158_v91 = _nw21;
              let _161_v92;
              _161_v92 = _dafny.Set.fromElements(_158_v91);
              _54_v19 = !(((_54_v19) ? (new BigNumber(521)) : (_28_v0))).isEqualTo(new BigNumber(((_161_v92).Intersect(_161_v92)).length));
              _29_v1 = _29_v1;
              let _162_v93;
              let _nw22 = Array((new BigNumber(26)).toNumber()).fill(false);
              _162_v93 = _nw22;
              let _index17 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_162_v93).length));
              (_162_v93)[_index17] = _54_v19;
              let _index18 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_162_v93).length));
              (_162_v93)[_index18] = false;
              _155_v88 = _module.__default.fm7(_28_v0, _28_v0, globalState);
            } else {
              let _163_v94;
              _163_v94 = _dafny.Map.Empty.slice().updateUnsafe(false,_54_v19);
              _163_v94 = (_163_v94).update(_54_v19, !(p1) || (p1));
              r1 = _28_v0;
              let _164_v95;
              let _nw23 = Array((new BigNumber(20)).toNumber());
              _164_v95 = _nw23;
              let _165_v96;
              let _nw24 = new _module.C0();
              _nw24.__ctor(!(p1));
              _165_v96 = _nw24;
              let _index19 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_164_v95).length));
              (_164_v95)[_index19] = _165_v96;
              let _index20 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_164_v95).length));
              (_164_v95)[_index20] = _165_v96;
              let _166_v97;
              let _nw25 = new _module.C0();
              _nw25.__ctor(p1);
              _166_v97 = _nw25;
              let _167_v98;
              _167_v98 = _dafny.Map.Empty.slice().updateUnsafe(_28_v0,!(p1));
              let _168_v99;
              _168_v99 = _dafny.Set.fromElements(_29_v1, _29_v1);
              let _169_v100;
              _169_v100 = _module.D1.create_DC2(_168_v99);
              let _170_v101;
              _170_v101 = new _dafny.CodePoint('w'.codePointAt(0));
              let _171_v102;
              _171_v102 = _dafny.Seq.of(_163_v94);
              let _172_v103;
              _172_v103 = _dafny.MultiSet.fromElements(_170_v101, (((_166_v97).fm6(_28_v0, _dafny.MultiSet.FromArray(_171_v102), _28_v0, globalState)) ? (_170_v101) : (_170_v101)), _170_v101, _170_v101, _170_v101);
              let _rhs18 = (_167_v98).contains(_28_v0);
              let _rhs19 = (function () {
                let _coll12 = new _dafny.Map();
                for (const _compr_12 of (_module.__default.fm3(globalState)).Elements) {
                  let _173_v104 = _compr_12;
                  if (_dafny.Seq.contains(_module.__default.fm3(globalState), _173_v104)) {
                    _coll12.push([(_173_v104).minus((_dafny.ZERO).minus(_28_v0)),false]);
                  }
                }
                return _coll12;
              }()).Merge((_167_v98).update(new BigNumber(556), p1));
              let _rhs20 = _169_v100;
              let _rhs21 = _164_v95;
              let _rhs22 = (_172_v103).Intersect(_172_v103);
              _54_v19 = _rhs18;
              _167_v98 = _rhs19;
              _169_v100 = _rhs20;
              _164_v95 = _rhs21;
              _172_v103 = _rhs22;
            }
            let _174_v105;
            _174_v105 = _dafny.Seq.UnicodeFromString("hr");
            _28_v0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((new BigNumber(20)).plus(new BigNumber(655)), _module.__default.safeModuloInt(_28_v0, new BigNumber((_174_v105).length))));
            _54_v19 = (p1) && (((_54_v19) ? (_54_v19) : (false)));
          }
        }
      }
      r0 = _28_v0;
      r1 = _28_v0;
      r2 = _28_v0;
      let _175_v106;
      _175_v106 = _dafny.Set.fromElements(true);
      let _176_v107;
      _176_v107 = _dafny.Set.fromElements(_175_v106, _175_v106, _175_v106);
      r3 = _176_v107;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _177_globalState;
      let _nw26 = new _module.GlobalState();
      _nw26.__ctor();
      _177_globalState = _nw26;
      let _178_v0;
      let _nw27 = Array((new BigNumber(24)).toNumber()).fill(false);
      _178_v0 = _nw27;
      let _179_v1;
      _179_v1 = new BigNumber(365);
      let _180_v2;
      _180_v2 = true;
      let _181_v3;
      let _182_v4;
      let _183_v5;
      let _184_v6;
      let _out0;
      let _out1;
      let _out2;
      let _out3;
      let _outcollector0 = _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe(_178_v0,_179_v1), _180_v2, _177_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _out3 = _outcollector0[3];
      _181_v3 = _out0;
      _182_v4 = _out1;
      _183_v5 = _out2;
      _184_v6 = _out3;
      let _185_v7;
      _185_v7 = _dafny.Seq.UnicodeFromString("gbehkj");
      _182_v4 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_dafny.ZERO).minus(_182_v4), new BigNumber((_185_v7).length)), new BigNumber(813));
      let _186_v8;
      _186_v8 = new _dafny.CodePoint('k'.codePointAt(0));
      let _187_v9;
      _187_v9 = _dafny.Map.Empty.slice().updateUnsafe(_180_v2,_179_v1);
      let _188_v10;
      _188_v10 = new _dafny.CodePoint('u'.codePointAt(0));
      let _189_v11;
      let _nw28 = Array((new BigNumber(20)).toNumber());
      _nw28[(_dafny.ZERO).toNumber()] = _186_v8;
      _nw28[(_dafny.ONE).toNumber()] = _186_v8;
      _nw28[(new BigNumber(2)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(3)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(4)).toNumber()] = _module.__default.fm0(_185_v7, (_dafny.ZERO).minus(_183_v5), _183_v5, _177_globalState);
      _nw28[(new BigNumber(5)).toNumber()] = _module.__default.fm0(_185_v7, new BigNumber((_dafny.Seq.of(new BigNumber((_187_v9).length), _183_v5, _179_v1)).length), _179_v1, _177_globalState);
      _nw28[(new BigNumber(6)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(7)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(8)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(9)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
      _nw28[(new BigNumber(11)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(12)).toNumber()] = (_188_v10);
      _nw28[(new BigNumber(13)).toNumber()] = _module.__default.fm0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(441)), ((_190_v8) => function (_191_i0) {
        return _190_v8;
      })(_186_v8)), _183_v5, _181_v3, _177_globalState);
      _nw28[(new BigNumber(14)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(15)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(16)).toNumber()] = (_185_v7)[_module.__default.safeIndex(_181_v3, new BigNumber((_185_v7).length))];
      _nw28[(new BigNumber(17)).toNumber()] = _186_v8;
      _nw28[(new BigNumber(18)).toNumber()] = _module.__default.fm0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(916)), function (_192_i1) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _181_v3, _181_v3, _177_globalState);
      _nw28[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
      _189_v11 = _nw28;
      _189_v11 = _189_v11;
      let _193_v12;
      _193_v12 = _dafny.Map.Empty.slice().updateUnsafe(_182_v4,_182_v4);
      if (!((((_193_v12).contains(_182_v4)) ? ((_193_v12).get(_182_v4)) : (_182_v4))).isEqualTo(_183_v5)) {
        if (true) {
          let _194_v13;
          let _nw29 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _194_v13 = _nw29;
          _194_v13 = _194_v13;
          let _195_v14;
          _195_v14 = _dafny.Seq.of(_182_v4, _182_v4);
          _187_v9 = (_187_v9).update(_180_v2, new BigNumber((_195_v14).length));
          let _196_v15;
          _196_v15 = _module.D1.create_DC1(_194_v13);
          _194_v13 = ((!(_180_v2)) ? (_194_v13) : ((_196_v15).dtor_cf1));
          let _index21 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_178_v0).length));
          (_178_v0)[_index21] = _180_v2;
          let _index22 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_178_v0).length));
          (_178_v0)[_index22] = _180_v2;
          let _index23 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_178_v0).length));
          (_178_v0)[_index23] = (_178_v0)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((_178_v0).length))];
        } else {
          let _197_v16;
          _197_v16 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,_182_v4);
          let _198_v17;
          let _199_v18;
          let _200_v19;
          let _201_v20;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = _module.__default.m0((_197_v16).Merge(_197_v16), true, _177_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _198_v17 = _out4;
          _199_v18 = _out5;
          _200_v19 = _out6;
          _201_v20 = _out7;
          _180_v2 = _180_v2;
          let _202_v21;
          _202_v21 = _dafny.Seq.of(_180_v2, _180_v2);
          let _203_v22;
          _203_v22 = _dafny.Set.fromElements(new BigNumber((_202_v21).length));
          let _204_v23;
          let _205_v24;
          let _206_v25;
          let _207_v26;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = _module.__default.m0(_197_v16, (_203_v22).equals(_203_v22), _177_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _204_v23 = _out8;
          _205_v24 = _out9;
          _206_v25 = _out10;
          _207_v26 = _out11;
          _204_v23 = _module.__default.safeDivisionInt(_206_v25, _206_v25);
          let _208_v27;
          _208_v27 = _module.D1.create_DC3(_dafny.MultiSet.FromArray(_dafny.Seq.of(_180_v2, _180_v2)), _180_v2, _180_v2);
          let _pat_let_tv1 = _180_v2;
          _180_v2 = ((function (_pat_let2_0) {
            return function (_209_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_210_dt__update_hcf5_h0) {
                  return _module.D1.create_DC3((_209_dt__update__tmp_h0).dtor_cf3, (_209_dt__update__tmp_h0).dtor_cf4, _210_dt__update_hcf5_h0);
                }(_pat_let3_0);
              }(_pat_let_tv1);
            }(_pat_let2_0);
          }(_208_v27)).dtor_cf4) && (_180_v2);
        }
        let _211_v28;
        _211_v28 = _dafny.MultiSet.fromElements(_183_v5, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(448)), function (_212_i2) {
          return _dafny.MultiSet.fromElements(new BigNumber(31));
        })).length), _182_v4, new BigNumber(-248));
        _180_v2 = (_211_v28).IsDisjointFrom((_211_v28).Union(_211_v28));
        let _213_v29;
        _213_v29 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(259)), ((_214_v8) => function (_215_i3) {
          return _214_v8;
        })(_186_v8)));
        let _216_v30;
        _216_v30 = _dafny.Seq.of(_180_v2);
        let _217_v31;
        _217_v31 = new BigNumber((_216_v30).length);
        let _218_v32;
        _218_v32 = _dafny.Set.fromElements(_180_v2, _180_v2);
        let _219_v33;
        _219_v33 = _dafny.Set.fromElements(!(_180_v2), _180_v2, _module.__default.fm1(_185_v7, (_217_v31), _218_v32, _177_globalState), !(!(_180_v2)));
        let _rhs23 = _180_v2;
        let _rhs24 = !((_dafny.MultiSet.fromElements(_185_v7)).IsSubsetOf((_213_v29).Intersect(_213_v29)));
        let _rhs25 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(232)), ((_220_v8) => function (_221_i4) {
          return _220_v8;
        })(_186_v8));
        let _rhs26 = _186_v8;
        let _rhs27 = new BigNumber((_218_v32).length);
        _180_v2 = _rhs23;
        _180_v2 = _rhs24;
        _185_v7 = _rhs25;
        _186_v8 = _rhs26;
        _181_v3 = _rhs27;
        let _222_v34;
        _222_v34 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,_182_v4);
        let _223_v35;
        _223_v35 = _module.D3.create_DC5(_222_v34);
        let _224_v36;
        let _225_v37;
        let _226_v38;
        let _227_v39;
        let _out12;
        let _out13;
        let _out14;
        let _out15;
        let _outcollector3 = _module.__default.m0((_223_v35).dtor_cf7, !(true), _177_globalState);
        _out12 = _outcollector3[0];
        _out13 = _outcollector3[1];
        _out14 = _outcollector3[2];
        _out15 = _outcollector3[3];
        _224_v36 = _out12;
        _225_v37 = _out13;
        _226_v38 = _out14;
        _227_v39 = _out15;
        let _228_v40;
        let _nw30 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
        _228_v40 = _nw30;
        let _229_v41;
        _229_v41 = _dafny.Seq.of(_185_v7, _185_v7, _185_v7);
        let _index24 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_228_v40).length));
        (_228_v40)[_index24] = _dafny.Seq.Concat(_229_v41, _229_v41);
        let _index25 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_228_v40).length));
        (_228_v40)[_index25] = _dafny.Seq.Concat(_dafny.Seq.of(_185_v7, _185_v7), _dafny.Seq.of(_185_v7));
      } else {
        _179_v1 = _179_v1;
        _180_v2 = false;
        _193_v12 = (_193_v12).update(_181_v3, (new BigNumber(265)).plus(_module.__default.fm2(_180_v2, _177_globalState)));
        let _rhs28 = _180_v2;
        let _rhs29 = _186_v8;
        let _rhs30 = _181_v3;
        _180_v2 = _rhs28;
        _186_v8 = _rhs29;
        _182_v4 = _rhs30;
        let _index26 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_178_v0).length));
        (_178_v0)[_index26] = _180_v2;
        let _index27 = _module.__default.safeIndex(new BigNumber(603), new BigNumber((_178_v0).length));
        (_178_v0)[_index27] = ((_182_v4).isLessThan(_183_v5)) === (_180_v2);
      }
      let _230_v42;
      _230_v42 = _dafny.Set.fromElements(_180_v2);
      if (((true) ? (_module.__default.fm1(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("pkkborlrx"), _module.__default.safeIndex(_183_v5, new BigNumber((_dafny.Seq.UnicodeFromString("pkkborlrx")).length)), _186_v8), _module.__default.fm2(_180_v2, _177_globalState), _230_v42, _177_globalState)) : ((_230_v42).equals(_230_v42)))) {
        let _231_v43;
        _231_v43 = _dafny.MultiSet.fromElements(_180_v2);
        let _232_v44;
        _232_v44 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,new BigNumber((_231_v43).cardinality()));
        let _233_v45;
        _233_v45 = _module.D3.create_DC5(_232_v44);
        let _234_v46;
        let _235_v47;
        let _236_v48;
        let _237_v49;
        let _out16;
        let _out17;
        let _out18;
        let _out19;
        let _outcollector4 = _module.__default.m0(((_180_v2) ? (_232_v44) : ((_233_v45).dtor_cf7)), _180_v2, _177_globalState);
        _out16 = _outcollector4[0];
        _out17 = _outcollector4[1];
        _out18 = _outcollector4[2];
        _out19 = _outcollector4[3];
        _234_v46 = _out16;
        _235_v47 = _out17;
        _236_v48 = _out18;
        _237_v49 = _out19;
        let _238_v50;
        let _239_v51;
        let _240_v52;
        let _241_v53;
        let _out20;
        let _out21;
        let _out22;
        let _out23;
        let _outcollector5 = _module.__default.m0(_232_v44, _180_v2, _177_globalState);
        _out20 = _outcollector5[0];
        _out21 = _outcollector5[1];
        _out22 = _outcollector5[2];
        _out23 = _outcollector5[3];
        _238_v50 = _out20;
        _239_v51 = _out21;
        _240_v52 = _out22;
        _241_v53 = _out23;
        let _242_v54;
        let _243_v55;
        let _244_v56;
        let _245_v57;
        let _out24;
        let _out25;
        let _out26;
        let _out27;
        let _outcollector6 = _module.__default.m0((_232_v44).Merge(_232_v44), _180_v2, _177_globalState);
        _out24 = _outcollector6[0];
        _out25 = _outcollector6[1];
        _out26 = _outcollector6[2];
        _out27 = _outcollector6[3];
        _242_v54 = _out24;
        _243_v55 = _out25;
        _244_v56 = _out26;
        _245_v57 = _out27;
        let _246_v58;
        let _247_v59;
        let _248_v60;
        let _249_v61;
        let _out28;
        let _out29;
        let _out30;
        let _out31;
        let _outcollector7 = _module.__default.m0((_232_v44).Merge(_232_v44), _180_v2, _177_globalState);
        _out28 = _outcollector7[0];
        _out29 = _outcollector7[1];
        _out30 = _outcollector7[2];
        _out31 = _outcollector7[3];
        _246_v58 = _out28;
        _247_v59 = _out29;
        _248_v60 = _out30;
        _249_v61 = _out31;
        let _rhs31 = !((_180_v2) || (_180_v2));
        let _rhs32 = _180_v2;
        let _rhs33 = _180_v2;
        let _rhs34 = ((_180_v2) ? (_189_v11) : (_189_v11));
        let _rhs35 = _180_v2;
        _180_v2 = _rhs31;
        _180_v2 = _rhs32;
        _180_v2 = _rhs33;
        _189_v11 = _rhs34;
        _180_v2 = _rhs35;
      } else {
        let _250_v62;
        _250_v62 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,_179_v1);
        let _251_v63;
        let _252_v64;
        let _253_v65;
        let _254_v66;
        let _out32;
        let _out33;
        let _out34;
        let _out35;
        let _outcollector8 = _module.__default.m0(_250_v62, _180_v2, _177_globalState);
        _out32 = _outcollector8[0];
        _out33 = _outcollector8[1];
        _out34 = _outcollector8[2];
        _out35 = _outcollector8[3];
        _251_v63 = _out32;
        _252_v64 = _out33;
        _253_v65 = _out34;
        _254_v66 = _out35;
        _253_v65 = (_181_v3).multipliedBy(new BigNumber((_187_v9).length));
        _180_v2 = !(!_dafny.Seq.contains(_module.__default.fm3(_177_globalState), ((_180_v2) ? (_183_v5) : ((((_187_v9).contains(_180_v2)) ? ((_187_v9).get(_180_v2)) : (new BigNumber((_dafny.Seq.of(_180_v2, _180_v2)).length)))))));
        let _255_v67;
        _255_v67 = _dafny.Seq.of(_185_v7, _185_v7, _185_v7);
        let _256_v68;
        let _nw31 = new _module.C0();
        _nw31.__ctor(_module.__default.fm1((_255_v67)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_181_v3)).length), new BigNumber((_255_v67).length))], _179_v1, _230_v42, _177_globalState));
        _256_v68 = _nw31;
        let _257_v69;
        let _nw32 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
        _257_v69 = _nw32;
        let _258_v70;
        _258_v70 = _dafny.Seq.of(_module.__default.fm2(_180_v2, _177_globalState));
        let _index28 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_257_v69).length));
        (_257_v69)[_index28] = _258_v70;
        let _index29 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_257_v69).length));
        (_257_v69)[_index29] = _258_v70;
      }
      let _259_v71;
      let _nw33 = new _module.C0();
      _nw33.__ctor(_180_v2);
      _259_v71 = _nw33;
      let _260_v72;
      _260_v72 = _dafny.Seq.of(true);
      let _hi0 = (_dafny.ZERO).minus(((false) ? (new BigNumber((_dafny.Seq.of(_259_v71)).length)) : ((_dafny.ZERO).minus(new BigNumber((_260_v72).length)))));
      for (let _261_i5 = new BigNumber(796); _261_i5.isLessThan(_hi0); _261_i5 = _261_i5.plus(_dafny.ONE)) {
        let _262_v73;
        _262_v73 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,_182_v4);
        let _263_v74;
        let _264_v75;
        let _265_v76;
        let _266_v77;
        let _out36;
        let _out37;
        let _out38;
        let _out39;
        let _outcollector9 = _module.__default.m0((_262_v73).Merge(_262_v73), _module.__default.fm1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-795)), ((_267_v8) => function (_268_i6) {
          return _267_v8;
        })(_186_v8)), _181_v3, _230_v42, _177_globalState), _177_globalState);
        _out36 = _outcollector9[0];
        _out37 = _outcollector9[1];
        _out38 = _outcollector9[2];
        _out39 = _outcollector9[3];
        _263_v74 = _out36;
        _264_v75 = _out37;
        _265_v76 = _out38;
        _266_v77 = _out39;
        let _269_v78;
        let _270_v79;
        let _271_v80;
        let _272_v81;
        let _out40;
        let _out41;
        let _out42;
        let _out43;
        let _outcollector10 = _module.__default.m0(((_180_v2) ? (_262_v73) : (_262_v73)), (_230_v42).IsDisjointFrom(_dafny.Set.fromElements(_180_v2, _180_v2, false, _180_v2, _180_v2)), _177_globalState);
        _out40 = _outcollector10[0];
        _out41 = _outcollector10[1];
        _out42 = _outcollector10[2];
        _out43 = _outcollector10[3];
        _269_v78 = _out40;
        _270_v79 = _out41;
        _271_v80 = _out42;
        _272_v81 = _out43;
        _180_v2 = !(false);
        let _273_v82;
        _273_v82 = _dafny.Seq.of(_259_v71, _259_v71);
        let _274_v83;
        let _nw34 = Array((new BigNumber(3)).toNumber());
        _nw34[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_273_v82, _module.__default.safeIndex(_270_v79, new BigNumber((_273_v82).length)), _259_v71);
        _nw34[(_dafny.ONE).toNumber()] = _273_v82;
        _nw34[(new BigNumber(2)).toNumber()] = _273_v82;
        _274_v83 = _nw34;
        let _index30 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_274_v83).length));
        (_274_v83)[_index30] = _273_v82;
        let _275_v84;
        _275_v84 = _dafny.Map.Empty.slice().updateUnsafe(_180_v2,_180_v2);
        let _276_v85;
        _276_v85 = _module.D1.create_DC3(_dafny.MultiSet.fromElements(_180_v2, (((_275_v84).contains(_180_v2)) ? ((_275_v84).get(_180_v2)) : (_180_v2))), true, _180_v2);
        let _277_v86;
        _277_v86 = _dafny.Set.fromElements(new BigNumber(38));
        let _278_v87;
        let _nw35 = Array((new BigNumber(7)).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = new BigNumber(826);
        _nw35[(_dafny.ONE).toNumber()] = _181_v3;
        _nw35[(new BigNumber(2)).toNumber()] = (((_187_v9).contains((_276_v85).dtor_cf4)) ? ((_187_v9).get((_276_v85).dtor_cf4)) : (_265_v76));
        _nw35[(new BigNumber(3)).toNumber()] = new BigNumber((_277_v86).length);
        _nw35[(new BigNumber(4)).toNumber()] = _181_v3;
        _nw35[(new BigNumber(5)).toNumber()] = _265_v76;
        _nw35[(new BigNumber(6)).toNumber()] = _271_v80;
        _278_v87 = _nw35;
        let _279_v88;
        _279_v88 = _dafny.Map.Empty.slice().updateUnsafe(_278_v87,_273_v82);
        let _280_v89;
        _280_v89 = _dafny.Seq.of(_278_v87);
        let _index31 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_274_v83).length));
        (_274_v83)[_index31] = (((_279_v88).contains((_280_v89)[_module.__default.safeIndex(new BigNumber((_185_v7).length), new BigNumber((_280_v89).length))])) ? ((_279_v88).get((_280_v89)[_module.__default.safeIndex(new BigNumber((_185_v7).length), new BigNumber((_280_v89).length))])) : (_dafny.Seq.Concat(_dafny.Seq.of(_259_v71), _273_v82)));
      }
      let _281_v90;
      let _nw36 = new _module.C0();
      _nw36.__ctor((new BigNumber((_260_v72).length)).isLessThan(new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_260_v72)[_module.__default.safeIndex(new BigNumber(-989), new BigNumber((_260_v72).length))]), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_182_v4, _179_v1)).length), new BigNumber((_dafny.Seq.of((_260_v72)[_module.__default.safeIndex(new BigNumber(-989), new BigNumber((_260_v72).length))])).length)), _180_v2)).length)));
      _281_v90 = _nw36;
      let _hi1 = _183_v5;
      for (let _282_i7 = _module.__default.fm2(_180_v2, _177_globalState); _282_i7.isLessThan(_hi1); _282_i7 = _282_i7.plus(_dafny.ONE)) {
        let _283_v91;
        _283_v91 = _dafny.Seq.of(new BigNumber((_185_v7).length));
        let _284_v92;
        _284_v92 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,new BigNumber((_283_v91).length));
        let _285_v93;
        let _286_v94;
        let _287_v95;
        let _288_v96;
        let _out44;
        let _out45;
        let _out46;
        let _out47;
        let _outcollector11 = _module.__default.m0((_284_v92).Merge(_284_v92), _180_v2, _177_globalState);
        _out44 = _outcollector11[0];
        _out45 = _outcollector11[1];
        _out46 = _outcollector11[2];
        _out47 = _outcollector11[3];
        _285_v93 = _out44;
        _286_v94 = _out45;
        _287_v95 = _out46;
        _288_v96 = _out47;
        _180_v2 = _180_v2;
        let _289_v97;
        let _nw37 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.of());
        _289_v97 = _nw37;
        let _290_v98;
        _290_v98 = _dafny.Seq.of(_259_v71, _281_v90);
        let _index32 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_289_v97).length));
        (_289_v97)[_index32] = _dafny.Seq.Concat(_290_v98, _290_v98);
        let _291_v99;
        _291_v99 = _dafny.Seq.of(_259_v71);
        let _index33 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_289_v97).length));
        (_289_v97)[_index33] = _dafny.Seq.Concat(_290_v98, (_291_v99));
        _259_v71 = _259_v71;
      }
      let _pat_let_tv2 = _183_v5;
      if (function (_source3) {
        let _292___mcc_h0 = _source3;
        let _293_cf0 = _292___mcc_h0;
        return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(45),_293_cf0)).length)).isLessThanOrEqualTo(_pat_let_tv2);
      }(((_180_v2) ? (_186_v8) : (_186_v8)))) {
        let _294_v100;
        let _nw38 = new _module.C0();
        _nw38.__ctor(_180_v2);
        _294_v100 = _nw38;
        let _295_v101;
        _295_v101 = _dafny.Map.Empty.slice().updateUnsafe(false,_294_v100);
        let _296_v102;
        let _nw39 = Array((new BigNumber(18)).toNumber());
        _nw39[(_dafny.ZERO).toNumber()] = _294_v100;
        _nw39[(_dafny.ONE).toNumber()] = _294_v100;
        _nw39[(new BigNumber(2)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(3)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(4)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(5)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(6)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(7)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(8)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(9)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(10)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(11)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(12)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(13)).toNumber()] = ((true) ? (_294_v100) : (_294_v100));
        _nw39[(new BigNumber(14)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(15)).toNumber()] = (((_295_v101).contains(!(_294_v100.f0))) ? ((_295_v101).get(!(_294_v100.f0))) : (_294_v100));
        _nw39[(new BigNumber(16)).toNumber()] = _294_v100;
        _nw39[(new BigNumber(17)).toNumber()] = _294_v100;
        _296_v102 = _nw39;
        let _rhs36 = _296_v102;
        let _rhs37 = ((((_294_v100.f0) ? (!(_294_v100.f0)) : (_294_v100.f0))) ? (_294_v100.f0) : (_294_v100.f0));
        _296_v102 = _rhs36;
        _180_v2 = _rhs37;
        _182_v4 = _182_v4;
        let _297_v103;
        _297_v103 = _dafny.Seq.of(_185_v7, _185_v7, _185_v7, _dafny.Seq.UnicodeFromString("hucmr"));
        let _298_v104;
        _298_v104 = _dafny.Seq.of(new BigNumber((_260_v72).length), new BigNumber(479));
        let _299_v105;
        _299_v105 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_297_v103, _dafny.Seq.update(_297_v103, _module.__default.safeIndex((_298_v104)[_module.__default.safeIndex(_182_v4, new BigNumber((_298_v104).length))], new BigNumber((_297_v103).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(5)), ((_300_v8) => function (_301_i8) {
          return _300_v8;
        })(_186_v8)))),_186_v8);
        _299_v105 = (_299_v105).update(_dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm7(_181_v3, _182_v4, _177_globalState), _185_v7, _185_v7), _dafny.Seq.of(_185_v7)), new _dafny.CodePoint('e'.codePointAt(0)));
        let _302_v106;
        let _init5 = function (_303_i9) {
          return (_303_i9).multipliedBy(new BigNumber(439));
        };
        let _nw40 = Array((new BigNumber(7)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw40.length); _i0_5++) {
          _nw40[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _302_v106 = _nw40;
        let _index34 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_302_v106).length));
        (_302_v106)[_index34] = _182_v4;
        let _304_v107;
        _304_v107 = _181_v3;
        let _305_v108;
        _305_v108 = _dafny.Set.fromElements(((_294_v100.f0) ? (_179_v1) : (_304_v107)), _304_v107);
        let _index35 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_302_v106).length));
        (_302_v106)[_index35] = new BigNumber((_305_v108).length);
        let _306_v109;
        _306_v109 = _185_v7;
        _185_v7 = (((false) ? (_306_v109) : (_module.__default.fm8(_180_v2, !(_180_v2), _179_v1, _177_globalState))));
      } else {
        let _307_v110;
        let _nw41 = new _module.C0();
        _nw41.__ctor((_179_v1).isLessThan(_182_v4));
        _307_v110 = _nw41;
        let _308_v111;
        _308_v111 = _module.D1.create_DC3(_dafny.MultiSet.fromElements(_180_v2, _180_v2), _180_v2, _180_v2);
        let _309_v112;
        _309_v112 = _module.D6.create_DC10(_281_v90);
        let _310_v113;
        _310_v113 = _dafny.Seq.of(_259_v71, _281_v90, ((_180_v2) ? (_259_v71) : (_259_v71)), _281_v90, (((_308_v111).dtor_cf4) ? (_307_v110) : ((_309_v112).dtor_cf14)));
        let _311_v114;
        let _nw42 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
        _311_v114 = _nw42;
        let _312_v115;
        _312_v115 = _dafny.Map.Empty.slice().updateUnsafe(_179_v1,_180_v2);
        let _313_v116;
        _313_v116 = _dafny.Set.fromElements(_312_v115);
        let _index36 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_311_v114).length));
        (_311_v114)[_index36] = _313_v116;
        let _314_v117;
        let _nw43 = new _module.C0();
        _nw43.__ctor(_180_v2);
        _314_v117 = _nw43;
        let _315_v118;
        _315_v118 = _dafny.Map.Empty.slice().updateUnsafe(_185_v7,_314_v117);
        let _index37 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_311_v114).length));
        let _rhs38 = new BigNumber(721);
        let _rhs39 = _dafny.Seq.update(_dafny.Seq.Concat(_310_v113, _dafny.Seq.Concat(_310_v113, _310_v113)), _module.__default.safeIndex(new BigNumber((_315_v118).length), new BigNumber((_dafny.Seq.Concat(_310_v113, _dafny.Seq.Concat(_310_v113, _310_v113))).length)), _307_v110);
        let _rhs40 = _module.__default.fm9(new BigNumber((_dafny.Seq.Concat(_185_v7, _dafny.Seq.UnicodeFromString("icsjbibk"))).length), new BigNumber((_dafny.Seq.update(_260_v72, _module.__default.safeIndex(_182_v4, new BigNumber((_260_v72).length)), _314_v117.f0)).length), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_314_v117.f0,new BigNumber(-841))).length)).plus((_314_v117).fm5(_181_v3, true, _177_globalState)), _186_v8, _177_globalState);
        let _lhs10 = _311_v114;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_311_v114).length));
        _182_v4 = _rhs38;
        _310_v113 = _rhs39;
        _lhs10[_lhs11] = _rhs40;
        let _316_v119;
        _316_v119 = _dafny.MultiSet.fromElements(_180_v2);
        let _317_v120;
        _317_v120 = _dafny.MultiSet.fromElements(_316_v119);
        _182_v4 = new BigNumber((_317_v120).cardinality());
        _179_v1 = _module.__default.fm2((_314_v117.f0) && (!(false)), _177_globalState);
        if (_180_v2) {
          let _318_v121;
          _318_v121 = _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-218), _182_v4, new BigNumber(-520), (_module.D6.create_DC11(_179_v1)).dtor_cf15, _179_v1));
          (_314_v117).f0 = _dafny.areEqual(_318_v121, _318_v121);
          let _319_v122;
          _319_v122 = _dafny.Seq.of(_181_v3, (_dafny.ZERO).minus(_181_v3));
          _319_v122 = _319_v122;
          _186_v8 = new _dafny.CodePoint('o'.codePointAt(0));
          let _320_v123;
          let _nw44 = new _module.C0();
          _nw44.__ctor(_180_v2);
          _320_v123 = _nw44;
          _180_v2 = _180_v2;
        } else {
          let _321_v124;
          let _nw45 = new _module.C0();
          _nw45.__ctor(_314_v117.f0);
          _321_v124 = _nw45;
          let _322_v125;
          let _nw46 = new _module.C0();
          _nw46.__ctor(_180_v2);
          _322_v125 = _nw46;
          let _323_v126;
          let _nw47 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _323_v126 = _nw47;
          let _index38 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_323_v126).length));
          (_323_v126)[_index38] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(347)), ((_324_v8) => function (_325_i10) {
            return _324_v8;
          })(_186_v8));
          let _index39 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_323_v126).length));
          (_323_v126)[_index39] = _185_v7;
          (_314_v117).f0 = _314_v117.f0;
          let _326_v127;
          let _nw48 = new _module.C0();
          _nw48.__ctor(_180_v2);
          _326_v127 = _nw48;
        }
      }
      let _327_v128;
      let _nw49 = new _module.C0();
      _nw49.__ctor(_180_v2);
      _327_v128 = _nw49;
      if ((_179_v1).isLessThan((_182_v4).minus(_183_v5))) {
        let _328_v129;
        _328_v129 = _dafny.Map.Empty.slice().updateUnsafe(_180_v2,_180_v2);
        let _329_v130;
        _329_v130 = _dafny.MultiSet.fromElements(_328_v129);
        _180_v2 = (_259_v71).fm6((_dafny.ZERO).minus(_183_v5), (_329_v130).Union(_329_v130), (_dafny.ZERO).minus((_dafny.ZERO).minus(_182_v4)), _177_globalState);
        let _rhs41 = (_188_v10);
        let _rhs42 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-426)), ((_330_v8) => function (_331_i11) {
          return _330_v8;
        })(_186_v8));
        let _rhs43 = new BigNumber(406);
        _186_v8 = _rhs41;
        _185_v7 = _rhs42;
        _182_v4 = _rhs43;
        let _332_v131;
        _332_v131 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_185_v7, _181_v3, new BigNumber((_187_v9).length), _177_globalState),new BigNumber((_185_v7).length));
        let _333_v132;
        _333_v132 = _dafny.Set.fromElements(_332_v131);
        _333_v132 = _333_v132;
        _180_v2 = _180_v2;
        let _334_v133;
        _334_v133 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,new BigNumber(940));
        let _335_v134;
        _335_v134 = _dafny.MultiSet.fromElements(_180_v2, _180_v2);
        let _pat_let_tv3 = _335_v134;
        let _336_v135;
        let _337_v136;
        let _338_v137;
        let _339_v138;
        let _out48;
        let _out49;
        let _out50;
        let _out51;
        let _outcollector12 = _module.__default.m0(_334_v133, (function (_pat_let4_0) {
          return function (_340_dt__update__tmp_h1) {
            return function (_pat_let5_0) {
              return function (_341_dt__update_hcf3_h0) {
                return _module.D1.create_DC3(_341_dt__update_hcf3_h0, (_340_dt__update__tmp_h1).dtor_cf4, (_340_dt__update__tmp_h1).dtor_cf5);
              }(_pat_let5_0);
            }(_pat_let_tv3);
          }(_pat_let4_0);
        }(_module.D1.create_DC3(_335_v134, !(_180_v2), _180_v2))).dtor_cf5, _177_globalState);
        _out48 = _outcollector12[0];
        _out49 = _outcollector12[1];
        _out50 = _outcollector12[2];
        _out51 = _outcollector12[3];
        _336_v135 = _out48;
        _337_v136 = _out49;
        _338_v137 = _out50;
        _339_v138 = _out51;
      } else {
        let _342_v139;
        let _nw50 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
        _342_v139 = _nw50;
        let _init6 = ((_343_v2) => function (_344_i12) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_343_v2,_343_v2)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_343_v2,_343_v2));
        })(_180_v2);
        let _nw51 = Array((new BigNumber(2)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw51.length); _i0_6++) {
          _nw51[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _342_v139 = _nw51;
        let _345_v140;
        _345_v140 = _dafny.Seq.of(_230_v42);
        let _346_v141;
        _346_v141 = _dafny.Map.Empty.slice().updateUnsafe(_183_v5,_345_v140);
        _346_v141 = (_346_v141).update(_module.__default.safeModuloInt(_182_v4, _182_v4), _dafny.Seq.Concat(_345_v140, _345_v140));
        let _347_v142;
        _347_v142 = _dafny.Map.Empty.slice().updateUnsafe(_181_v3,_185_v7);
        let _348_v143;
        _348_v143 = _dafny.Seq.of(_185_v7, _185_v7, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ffchstvvi"), _185_v7));
        let _rhs44 = _260_v72;
        let _rhs45 = _347_v142;
        let _rhs46 = (_327_v128).fm5(_183_v5, false, _177_globalState);
        let _rhs47 = _185_v7;
        let _rhs48 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-882)), ((_349_v7) => function (_350_i13) {
          return _349_v7;
        })(_185_v7)), _348_v143);
        _260_v72 = _rhs44;
        _347_v142 = _rhs45;
        _182_v4 = _rhs46;
        _185_v7 = _rhs47;
        _348_v143 = _rhs48;
        _281_v90 = _281_v90;
        let _351_v144;
        let _init7 = ((_352_v12) => function (_353_i14) {
          return _module.__default.safeDivisionInt(_353_i14, new BigNumber((_352_v12).length));
        })(_193_v12);
        let _nw52 = Array((new BigNumber(4)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw52.length); _i0_7++) {
          _nw52[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _351_v144 = _nw52;
        let _index40 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_351_v144).length));
        (_351_v144)[_index40] = ((false) ? ((_dafny.ZERO).minus(_182_v4)) : (_183_v5));
        let _354_v145;
        _354_v145 = _dafny.MultiSet.fromElements(_185_v7, (_dafny.Seq.UnicodeFromString("stnwqxv")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(61)), ((_355_v8) => function (_356_i15) {
          return _355_v8;
        })(_186_v8)));
        let _index41 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_351_v144).length));
        let _rhs49 = (_354_v145).IsSubsetOf((_354_v145).Difference(_354_v145));
        let _rhs50 = (_183_v5).multipliedBy(_181_v3);
        let _rhs51 = (_281_v90).fm5(_182_v4, !(_180_v2), _177_globalState);
        let _lhs12 = _351_v144;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_351_v144).length));
        _180_v2 = _rhs49;
        _lhs12[_lhs13] = _rhs50;
        _183_v5 = _rhs51;
      }
      let _357_i16;
      _357_i16 = _dafny.ZERO;
      L3: {
        while ((_260_v72)[_module.__default.safeIndex(_179_v1, new BigNumber((_260_v72).length))]) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_357_i16)) {
              break L3;
            }
            _357_i16 = (_357_i16).plus(_dafny.ONE);
            _327_v128 = _281_v90;
            let _358_v146;
            _358_v146 = _dafny.Map.Empty.slice().updateUnsafe((_182_v4).multipliedBy((_327_v128).fm5(new BigNumber((_230_v42).length), _180_v2, _177_globalState)),_180_v2);
            _358_v146 = _358_v146;
            let _359_v147;
            _359_v147 = _dafny.MultiSet.fromElements(_182_v4, new BigNumber(122));
            let _360_v148;
            _360_v148 = _dafny.Seq.of(new BigNumber(176), _183_v5, new BigNumber(-170), _module.__default.fm2(_module.__default.fm1(_185_v7, new BigNumber(943), _dafny.Set.fromElements(_180_v2), _177_globalState), _177_globalState));
            let _rhs52 = (_dafny.MultiSet.FromArray(_360_v148)).IsSubsetOf(_359_v147);
            let _rhs53 = _186_v8;
            let _rhs54 = (((!(false)) || (_180_v2)) ? (_module.__default.safeDivisionInt(_183_v5, _183_v5)) : (new BigNumber((_185_v7).length)));
            _180_v2 = _rhs52;
            _186_v8 = _rhs53;
            _182_v4 = _rhs54;
            _183_v5 = new BigNumber(((_185_v7)).length);
          }
        }
      }
      _183_v5 = (_dafny.ZERO).minus(_179_v1);
      let _361_v149;
      _361_v149 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_180_v2,_180_v2));
      let _362_v150;
      _362_v150 = _dafny.Map.Empty.slice().updateUnsafe(_180_v2,_180_v2);
      if ((_327_v128).fm6(new BigNumber((_185_v7).length), (_361_v149).Union((_361_v149).update(_362_v150, _module.__default.abs(_182_v4))), new BigNumber(342), _177_globalState)) {
        let _rhs55 = _178_v0;
        let _rhs56 = (_dafny.ZERO).minus((_183_v5).plus(_183_v5));
        let _rhs57 = _180_v2;
        _178_v0 = _rhs55;
        _181_v3 = _rhs56;
        _180_v2 = _rhs57;
        _185_v7 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(522)), ((_363_v8) => function (_364_i17) {
          return _363_v8;
        })(_186_v8));
        if (_180_v2) {
          _178_v0 = _178_v0;
          _181_v3 = _module.__default.safeDivisionInt(new BigNumber(-845), _182_v4);
          _185_v7 = _185_v7;
          let _365_v151;
          let _nw53 = new _module.C0();
          _nw53.__ctor(!(_180_v2));
          _365_v151 = _nw53;
          let _366_v152;
          let _init8 = ((_367_v10) => function (_368_i18) {
            return _367_v10;
          })(_188_v10);
          let _nw54 = Array((new BigNumber(16)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw54.length); _i0_8++) {
            _nw54[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _366_v152 = _nw54;
          let _index42 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_366_v152).length));
          (_366_v152)[_index42] = _188_v10;
          let _369_v153;
          _369_v153 = _dafny.Seq.of(_183_v5);
          let _370_v154;
          _370_v154 = _dafny.Seq.of(_369_v153);
          let _index43 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_366_v152).length));
          (_366_v152)[_index43] = _186_v8;
        } else {
          let _index44 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length));
          (_178_v0)[_index44] = _180_v2;
          let _index45 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length));
          (_178_v0)[_index45] = _180_v2;
          let _371_v155;
          let _nw55 = new _module.C0();
          _nw55.__ctor(!((_178_v0)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length))]) || (_180_v2));
          _371_v155 = _nw55;
          let _372_v156;
          _372_v156 = _dafny.Set.fromElements(_186_v8);
          let _373_v157;
          _373_v157 = _dafny.Map.Empty.slice().updateUnsafe(((_180_v2) ? ((_260_v72)[_module.__default.safeIndex(_179_v1, new BigNumber((_260_v72).length))]) : (_180_v2)),(_dafny.Set.fromElements(_186_v8, _186_v8, _186_v8)).Union(_372_v156));
          let _374_v158;
          _374_v158 = _dafny.Set.fromElements(_183_v5);
          let _375_v159;
          _375_v159 = _dafny.MultiSet.fromElements((_178_v0)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length))]);
          let _376_v160;
          _376_v160 = _module.D1.create_DC3(_375_v159, _180_v2, _180_v2);
          let _index46 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length));
          let _rhs58 = _module.__default.safeDivisionInt(_179_v1, (new BigNumber(-967)).minus(_179_v1));
          let _rhs59 = ((_180_v2) ? (_180_v2) : ((_178_v0)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length))]));
          let _rhs60 = _module.__default.safeModuloInt(_182_v4, (new BigNumber((_374_v158).length)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_185_v7,(_376_v160).dtor_cf4)).length)));
          let _rhs61 = (_373_v157).Merge((_module.__default.fm11(_177_globalState)).Merge(_373_v157));
          let _lhs14 = _178_v0;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length));
          _182_v4 = _rhs58;
          _lhs14[_lhs15] = _rhs59;
          _183_v5 = _rhs60;
          _373_v157 = _rhs61;
          let _377_v161;
          let _nw56 = Array((new BigNumber(15)).toNumber()).fill(false);
          _377_v161 = _nw56;
          let _378_v162;
          _378_v162 = _dafny.Map.Empty.slice().updateUnsafe(_377_v161,(_371_v155).fm5(_183_v5, (_178_v0)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length))], _177_globalState));
          let _379_v163;
          _379_v163 = _dafny.Map.Empty.slice().updateUnsafe(_327_v128,_378_v162);
          let _380_v164;
          let _381_v165;
          let _382_v166;
          let _383_v167;
          let _out52;
          let _out53;
          let _out54;
          let _out55;
          let _outcollector13 = _module.__default.m0((((_379_v163).contains(_281_v90)) ? ((_379_v163).get(_281_v90)) : (_378_v162)), _180_v2, _177_globalState);
          _out52 = _outcollector13[0];
          _out53 = _outcollector13[1];
          _out54 = _outcollector13[2];
          _out55 = _outcollector13[3];
          _380_v164 = _out52;
          _381_v165 = _out53;
          _382_v166 = _out54;
          _383_v167 = _out55;
          let _384_v168;
          let _init9 = ((_385_v72) => function (_386_i19) {
            return _385_v72;
          })(_260_v72);
          let _nw57 = Array((new BigNumber(3)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw57.length); _i0_9++) {
            _nw57[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _384_v168 = _nw57;
          let _387_v169;
          _387_v169 = _dafny.MultiSet.fromElements(_327_v128);
          let _index47 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length));
          let _rhs62 = ((_dafny.MultiSet.fromElements(_327_v128, _327_v128)).Difference(_387_v169)).IsSubsetOf(_387_v169);
          let _rhs63 = ((!((_178_v0)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length))])) ? (_384_v168) : (_384_v168));
          let _rhs64 = _182_v4;
          let _lhs16 = _178_v0;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_178_v0).length));
          _lhs16[_lhs17] = _rhs62;
          _384_v168 = _rhs63;
          _181_v3 = _rhs64;
        }
        if ((_327_v128).fm6(new BigNumber((_230_v42).length), _361_v149, _182_v4, _177_globalState)) {
          _183_v5 = new BigNumber(826);
          let _388_v170;
          _388_v170 = _dafny.MultiSet.fromElements(_182_v4);
          let _389_v171;
          _389_v171 = _dafny.Map.Empty.slice().updateUnsafe(_180_v2,_178_v0);
          let _390_v172;
          let _nw58 = Array((new BigNumber(18)).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = (_183_v5).minus(_179_v1);
          _nw58[(_dafny.ONE).toNumber()] = new BigNumber((_193_v12).length);
          _nw58[(new BigNumber(2)).toNumber()] = new BigNumber(((_dafny.MultiSet.fromElements(_182_v4)).update(_181_v3, _module.__default.abs(_182_v4))).cardinality());
          _nw58[(new BigNumber(3)).toNumber()] = _183_v5;
          _nw58[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_185_v7).length), new BigNumber((_230_v42).length));
          _nw58[(new BigNumber(5)).toNumber()] = _182_v4;
          _nw58[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_181_v3).multipliedBy(_183_v5));
          _nw58[(new BigNumber(7)).toNumber()] = ((_180_v2) ? ((_dafny.ZERO).minus(new BigNumber((_388_v170).cardinality()))) : (_183_v5));
          _nw58[(new BigNumber(8)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_180_v2,_178_v0)).Merge(_389_v171)).length);
          _nw58[(new BigNumber(9)).toNumber()] = _181_v3;
          _nw58[(new BigNumber(10)).toNumber()] = _182_v4;
          _nw58[(new BigNumber(11)).toNumber()] = _182_v4;
          _nw58[(new BigNumber(12)).toNumber()] = _182_v4;
          _nw58[(new BigNumber(13)).toNumber()] = new BigNumber(-876);
          _nw58[(new BigNumber(14)).toNumber()] = _183_v5;
          _nw58[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_180_v2, _180_v2, _180_v2)).length);
          _nw58[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus((_281_v90).fm5(new BigNumber(947), _180_v2, _177_globalState));
          _nw58[(new BigNumber(17)).toNumber()] = _182_v4;
          _390_v172 = _nw58;
          let _index48 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_390_v172).length));
          (_390_v172)[_index48] = _183_v5;
          let _index49 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_390_v172).length));
          (_390_v172)[_index49] = new BigNumber(-73);
          let _391_v173;
          let _392_v174;
          let _393_v175;
          let _394_v176;
          let _out56;
          let _out57;
          let _out58;
          let _out59;
          let _outcollector14 = _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe(_178_v0,_183_v5), _180_v2, _177_globalState);
          _out56 = _outcollector14[0];
          _out57 = _outcollector14[1];
          _out58 = _outcollector14[2];
          _out59 = _outcollector14[3];
          _391_v173 = _out56;
          _392_v174 = _out57;
          _393_v175 = _out58;
          _394_v176 = _out59;
          let _395_v177;
          _395_v177 = _dafny.MultiSet.fromElements(_180_v2);
          _181_v3 = (_392_v174).multipliedBy(new BigNumber((_395_v177).cardinality()));
          _179_v1 = (_327_v128).fm5(((_390_v172)[_module.__default.safeIndex(new BigNumber(493), new BigNumber((_390_v172).length))]).multipliedBy(_393_v175), false, _177_globalState);
        } else {
          _183_v5 = (_183_v5).plus(_182_v4);
          let _396_v178;
          _396_v178 = _module.D3.create_DC6((_dafny.ZERO).minus(new BigNumber((_260_v72).length)), _182_v4, _180_v2);
          _180_v2 = (_396_v178).dtor_cf10;
          let _397_v179;
          _397_v179 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,_181_v3);
          let _398_v180;
          let _399_v181;
          let _400_v182;
          let _401_v183;
          let _out60;
          let _out61;
          let _out62;
          let _out63;
          let _outcollector15 = _module.__default.m0((_397_v179).Merge(_397_v179), (((_362_v150).contains(_180_v2)) ? ((_362_v150).get(_180_v2)) : (_180_v2)), _177_globalState);
          _out60 = _outcollector15[0];
          _out61 = _outcollector15[1];
          _out62 = _outcollector15[2];
          _out63 = _outcollector15[3];
          _398_v180 = _out60;
          _399_v181 = _out61;
          _400_v182 = _out62;
          _401_v183 = _out63;
          let _402_v185;
          _402_v185 = _185_v7;
          let _403_v186;
          _403_v186 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(417),_402_v185);
          let _404_v187;
          _404_v187 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_260_v72).length),_186_v8);
          let _405_v188;
          _405_v188 = _dafny.Seq.of(function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of (_403_v186).Keys.Elements) {
              let _406_v184 = _compr_13;
              if ((_403_v186).contains(_406_v184)) {
                _coll13.push([(_406_v184).multipliedBy(_399_v181),_186_v8]);
              }
            }
            return _coll13;
          }(), (_404_v187).Merge(_404_v187), (_404_v187).Merge(_404_v187), _404_v187);
          let _407_v189;
          let _nw59 = Array((new BigNumber(27)).toNumber()).fill([]);
          _407_v189 = _nw59;
          let _408_v190;
          let _init10 = ((_409_v1) => function (_410_i20) {
            return _module.__default.safeModuloInt(_410_i20, _409_v1);
          })(_179_v1);
          let _nw60 = Array((new BigNumber(17)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw60.length); _i0_10++) {
            _nw60[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _408_v190 = _nw60;
          let _index50 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_407_v189).length));
          (_407_v189)[_index50] = _408_v190;
          let _411_v191;
          _411_v191 = _dafny.Map.Empty.slice().updateUnsafe(_230_v42,_180_v2);
          let _index51 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_407_v189).length));
          let _rhs65 = ((false) ? (_405_v188) : (_405_v188));
          let _rhs66 = _281_v90;
          let _rhs67 = _408_v190;
          let _rhs68 = (_399_v181).plus((_dafny.ZERO).minus(new BigNumber((_185_v7).length)));
          let _rhs69 = (_411_v191).equals(_411_v191);
          let _lhs18 = _407_v189;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_407_v189).length));
          _405_v188 = _rhs65;
          _259_v71 = _rhs66;
          _lhs18[_lhs19] = _rhs67;
          _400_v182 = _rhs68;
          _180_v2 = _rhs69;
          let _rhs70 = ((_dafny.ZERO).minus(_181_v3)).plus(new BigNumber(229));
          let _rhs71 = (_396_v178).dtor_cf8;
          let _rhs72 = _186_v8;
          let _rhs73 = _178_v0;
          _399_v181 = _rhs70;
          _398_v180 = _rhs71;
          _186_v8 = _rhs72;
          _178_v0 = _rhs73;
        }
        let _412_v192;
        _412_v192 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,new BigNumber((_260_v72).length));
        let _413_v193;
        let _414_v194;
        let _415_v195;
        let _416_v196;
        let _out64;
        let _out65;
        let _out66;
        let _out67;
        let _outcollector16 = _module.__default.m0(_412_v192, _180_v2, _177_globalState);
        _out64 = _outcollector16[0];
        _out65 = _outcollector16[1];
        _out66 = _outcollector16[2];
        _out67 = _outcollector16[3];
        _413_v193 = _out64;
        _414_v194 = _out65;
        _415_v195 = _out66;
        _416_v196 = _out67;
      } else {
        _180_v2 = _180_v2;
        let _417_v197;
        let _nw61 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _417_v197 = _nw61;
        let _index52 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_417_v197).length));
        (_417_v197)[_index52] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("odeh"), _185_v7)).length);
        let _index53 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_417_v197).length));
        (_417_v197)[_index53] = _183_v5;
        let _418_v198;
        _418_v198 = _module.D7.create_DC12(_module.__default.fm12(true, _186_v8, _180_v2, _179_v1, _177_globalState));
        let _419_v199;
        _419_v199 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,new BigNumber(((_418_v198).dtor_cf16).cardinality()));
        let _420_v200;
        let _421_v201;
        let _422_v202;
        let _423_v203;
        let _out68;
        let _out69;
        let _out70;
        let _out71;
        let _outcollector17 = _module.__default.m0(_419_v199, _180_v2, _177_globalState);
        _out68 = _outcollector17[0];
        _out69 = _outcollector17[1];
        _out70 = _outcollector17[2];
        _out71 = _outcollector17[3];
        _420_v200 = _out68;
        _421_v201 = _out69;
        _422_v202 = _out70;
        _423_v203 = _out71;
        _180_v2 = !(_180_v2);
        let _424_v204;
        let _nw62 = Array((new BigNumber(2)).toNumber()).fill([]);
        _424_v204 = _nw62;
        let _rhs74 = _179_v1;
        let _rhs75 = _424_v204;
        let _rhs76 = !(_180_v2) || (_180_v2);
        let _rhs77 = _180_v2;
        _179_v1 = _rhs74;
        _424_v204 = _rhs75;
        _180_v2 = _rhs76;
        _180_v2 = _rhs77;
      }
      let _index54 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length));
      (_178_v0)[_index54] = _180_v2;
      let _index55 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length));
      (_178_v0)[_index55] = _module.__default.fm1(_185_v7, new BigNumber((_185_v7).length), _230_v42, _177_globalState);
      let _425_v205;
      _425_v205 = _dafny.Seq.of(new BigNumber((_185_v7).length), new BigNumber((_260_v72).length));
      let _426_v206;
      _426_v206 = _dafny.MultiSet.fromElements(_425_v205);
      let _427_v207;
      _427_v207 = _module.D7.create_DC12(_426_v206);
      let _428_v208;
      _428_v208 = _module.D7.create_DC12((_427_v207).dtor_cf16);
      let _429_v209;
      _429_v209 = _module.D7.create_DC14(_428_v208);
      let _pat_let_tv4 = _428_v208;
      let _source4 = function (_pat_let6_0) {
        return function (_430_dt__update__tmp_h4) {
          return function (_pat_let7_0) {
            return function (_431_dt__update_hcf18_h0) {
              return _module.D7.create_DC14(_431_dt__update_hcf18_h0);
            }(_pat_let7_0);
          }(_pat_let_tv4);
        }(_pat_let6_0);
      }(_429_v209);
      if (_source4.is_DC13) {
        let _432___mcc_h1 = (_source4).cf17;
        let _433_cf17 = _432___mcc_h1;
        _182_v4 = (_dafny.ZERO).minus(_181_v3);
        _433_cf17 = _433_cf17;
        let _434_v210;
        _434_v210 = _dafny.Map.Empty.slice().updateUnsafe(_433_cf17,_259_v71);
        let _435_v211;
        let _nw63 = Array((new BigNumber(18)).toNumber());
        _nw63[(_dafny.ZERO).toNumber()] = _259_v71;
        _nw63[(_dafny.ONE).toNumber()] = _281_v90;
        _nw63[(new BigNumber(2)).toNumber()] = _281_v90;
        _nw63[(new BigNumber(3)).toNumber()] = _281_v90;
        _nw63[(new BigNumber(4)).toNumber()] = _327_v128;
        _nw63[(new BigNumber(5)).toNumber()] = _259_v71;
        _nw63[(new BigNumber(6)).toNumber()] = _259_v71;
        _nw63[(new BigNumber(7)).toNumber()] = _327_v128;
        _nw63[(new BigNumber(8)).toNumber()] = _259_v71;
        _nw63[(new BigNumber(9)).toNumber()] = _259_v71;
        _nw63[(new BigNumber(10)).toNumber()] = _327_v128;
        _nw63[(new BigNumber(11)).toNumber()] = (((_434_v210).contains(_180_v2)) ? ((_434_v210).get(_180_v2)) : (_327_v128));
        _nw63[(new BigNumber(12)).toNumber()] = _327_v128;
        _nw63[(new BigNumber(13)).toNumber()] = _327_v128;
        _nw63[(new BigNumber(14)).toNumber()] = _281_v90;
        _nw63[(new BigNumber(15)).toNumber()] = _281_v90;
        _nw63[(new BigNumber(16)).toNumber()] = _281_v90;
        _nw63[(new BigNumber(17)).toNumber()] = _259_v71;
        _435_v211 = _nw63;
        let _index56 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_435_v211).length));
        (_435_v211)[_index56] = _281_v90;
        let _index57 = _module.__default.safeIndex(new BigNumber(917), new BigNumber((_435_v211).length));
        (_435_v211)[_index57] = _327_v128;
        _179_v1 = (_module.__default.safeDivisionInt(_181_v3, _183_v5)).multipliedBy((new BigNumber((_260_v72).length)).minus(new BigNumber(283)));
      } else if (_source4.is_DC12) {
        let _436___mcc_h2 = (_source4).cf16;
        let _437_cf16 = _436___mcc_h2;
        let _index58 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length));
        let _rhs78 = _183_v5;
        let _rhs79 = _182_v4;
        let _rhs80 = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xya"), _dafny.Seq.UnicodeFromString("kbshdbhf")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-774)), ((_438_v8) => function (_439_i21) {
          return _438_v8;
        })(_186_v8)));
        let _rhs81 = _186_v8;
        let _rhs82 = _180_v2;
        let _lhs20 = _178_v0;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length));
        _182_v4 = _rhs78;
        _181_v3 = _rhs79;
        _180_v2 = _rhs80;
        _186_v8 = _rhs81;
        _lhs20[_lhs21] = _rhs82;
        let _440_v212;
        _440_v212 = _dafny.Set.fromElements(_181_v3);
        let _index59 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length));
        (_178_v0)[_index59] = (_259_v71).fm6(_module.__default.safeDivisionInt(_183_v5, new BigNumber((_185_v7).length)), _361_v149, (_dafny.ZERO).minus((_183_v5).minus(new BigNumber((_440_v212).length))), _177_globalState);
        let _441_v213;
        _441_v213 = _dafny.MultiSet.fromElements(_180_v2);
        let _442_v214;
        _442_v214 = _module.D1.create_DC3(_441_v213, _180_v2, (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))]);
        let _source5 = ((true) ? (_442_v214) : (_442_v214));
        if (_source5.is_DC2) {
          let _443___mcc_h4 = (_source5).cf2;
          let _444_cf2 = _443___mcc_h4;
          _183_v5 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_182_v4));
          let _445_v215;
          let _446_v216;
          let _447_v217;
          let _448_v218;
          let _out72;
          let _out73;
          let _out74;
          let _out75;
          let _outcollector18 = _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe(_178_v0,_181_v3), _180_v2, _177_globalState);
          _out72 = _outcollector18[0];
          _out73 = _outcollector18[1];
          _out74 = _outcollector18[2];
          _out75 = _outcollector18[3];
          _445_v215 = _out72;
          _446_v216 = _out73;
          _447_v217 = _out74;
          _448_v218 = _out75;
          let _449_v219;
          _449_v219 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,new BigNumber((_425_v205).length));
          let _450_v220;
          let _451_v221;
          let _452_v222;
          let _453_v223;
          let _out76;
          let _out77;
          let _out78;
          let _out79;
          let _outcollector19 = _module.__default.m0((_449_v219).Merge(_449_v219), _180_v2, _177_globalState);
          _out76 = _outcollector19[0];
          _out77 = _outcollector19[1];
          _out78 = _outcollector19[2];
          _out79 = _outcollector19[3];
          _450_v220 = _out76;
          _451_v221 = _out77;
          _452_v222 = _out78;
          _453_v223 = _out79;
          let _454_v224;
          _454_v224 = _module.D3.create_DC6(_module.__default.safeDivisionInt(new BigNumber((_441_v213).cardinality()), _183_v5), (new BigNumber(225)).minus((_dafny.ZERO).minus(_182_v4)), (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))]);
          let _pat_let_tv5 = _185_v7;
          let _pat_let_tv6 = _185_v7;
          _454_v224 = function (_pat_let8_0) {
            return function (_455_dt__update__tmp_h5) {
              return function (_pat_let9_0) {
                return function (_456_dt__update_hcf9_h0) {
                  return _module.D3.create_DC6((_455_dt__update__tmp_h5).dtor_cf8, _456_dt__update_hcf9_h0, (_455_dt__update__tmp_h5).dtor_cf10);
                }(_pat_let9_0);
              }(new BigNumber((_dafny.Seq.Concat(_pat_let_tv5, _pat_let_tv6)).length));
            }(_pat_let8_0);
          }(_454_v224);
        } else if (_source5.is_DC3) {
          let _457___mcc_h5 = (_source5).cf3;
          let _458___mcc_h6 = (_source5).cf4;
          let _459___mcc_h7 = (_source5).cf5;
          let _460_cf5 = _459___mcc_h7;
          let _461_cf4 = _458___mcc_h6;
          let _462_cf3 = _457___mcc_h5;
          _260_v72 = _dafny.Seq.Concat(_260_v72, _260_v72);
          let _463_v225;
          _463_v225 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,new BigNumber(691));
          let _464_v226;
          _464_v226 = _dafny.Seq.of(_259_v71);
          let _465_v227;
          _465_v227 = _dafny.Map.Empty.slice().updateUnsafe(_464_v226,(_259_v71).fm6(_183_v5, _361_v149, _179_v1, _177_globalState));
          let _466_v228;
          _466_v228 = _dafny.Map.Empty.slice().updateUnsafe(_465_v227,false);
          let _467_v229;
          let _468_v230;
          let _469_v231;
          let _470_v232;
          let _out80;
          let _out81;
          let _out82;
          let _out83;
          let _outcollector20 = _module.__default.m0(_463_v225, (((_466_v228).contains(_465_v227)) ? ((_466_v228).get(_465_v227)) : ((_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))])), _177_globalState);
          _out80 = _outcollector20[0];
          _out81 = _outcollector20[1];
          _out82 = _outcollector20[2];
          _out83 = _outcollector20[3];
          _467_v229 = _out80;
          _468_v230 = _out81;
          _469_v231 = _out82;
          _470_v232 = _out83;
          _185_v7 = _dafny.Seq.Concat(_module.__default.fm7(_179_v1, _467_v229, _177_globalState), _185_v7);
          _182_v4 = (_dafny.ZERO).minus(_182_v4);
        } else {
          let _471___mcc_h8 = (_source5).cf1;
          let _472_cf1 = _471___mcc_h8;
          let _473_v233;
          _473_v233 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,new BigNumber((_185_v7).length));
          let _474_v234;
          _474_v234 = _dafny.Map.Empty.slice().updateUnsafe(_281_v90,_182_v4);
          let _475_v235;
          let _476_v236;
          let _477_v237;
          let _478_v238;
          let _out84;
          let _out85;
          let _out86;
          let _out87;
          let _outcollector21 = _module.__default.m0((_473_v233).Merge((_473_v233).update(_178_v0, _182_v4)), (_179_v1).isLessThan(new BigNumber((_474_v234).length)), _177_globalState);
          _out84 = _outcollector21[0];
          _out85 = _outcollector21[1];
          _out86 = _outcollector21[2];
          _out87 = _outcollector21[3];
          _475_v235 = _out84;
          _476_v236 = _out85;
          _477_v237 = _out86;
          _478_v238 = _out87;
          let _479_v239;
          let _480_v240;
          let _481_v241;
          let _482_v242;
          let _out88;
          let _out89;
          let _out90;
          let _out91;
          let _outcollector22 = _module.__default.m0(_473_v233, (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))], _177_globalState);
          _out88 = _outcollector22[0];
          _out89 = _outcollector22[1];
          _out90 = _outcollector22[2];
          _out91 = _outcollector22[3];
          _479_v239 = _out88;
          _480_v240 = _out89;
          _481_v241 = _out90;
          _482_v242 = _out91;
          _230_v42 = _module.__default.fm13(_dafny.Seq.IsPrefixOf(_185_v7, _185_v7), !((_441_v213).contains(_180_v2)), false, _177_globalState);
          let _483_v243;
          _483_v243 = _module.D6.create_DC11(_481_v241);
          let _484_v244;
          _484_v244 = _dafny.Set.fromElements(_483_v243, _module.D6.create_DC11(_480_v240), _483_v243);
          let _485_v245;
          let _nw64 = new _module.C0();
          _nw64.__ctor(_180_v2);
          _485_v245 = _nw64;
          let _486_v246;
          _486_v246 = _dafny.Seq.of(_485_v245, _485_v245);
          let _487_v247;
          let _nw65 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Set.Empty);
          _487_v247 = _nw65;
          let _rhs83 = _484_v244;
          let _rhs84 = _dafny.Seq.Concat(_dafny.Seq.of(_485_v245, _485_v245), _dafny.Seq.Concat(_dafny.Seq.of(_485_v245), _486_v246));
          let _rhs85 = _487_v247;
          _484_v244 = _rhs83;
          _486_v246 = _rhs84;
          _487_v247 = _rhs85;
        }
        let _488_v248;
        _488_v248 = _dafny.Map.Empty.slice().updateUnsafe(_180_v2,_178_v0);
        _488_v248 = (_488_v248).update(_dafny.Seq.contains(_185_v7, new _dafny.CodePoint('m'.codePointAt(0))), _178_v0);
      } else {
        let _489___mcc_h3 = (_source4).cf18;
        let _490_cf18 = _489___mcc_h3;
        if ((_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))]) {
          _182_v4 = _183_v5;
          let _491_v249;
          _491_v249 = _dafny.Set.fromElements(_425_v205, _425_v205, _425_v205, _425_v205);
          let _index60 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length));
          (_178_v0)[_index60] = !((_491_v249).IsSubsetOf(_module.__default.fm14(new BigNumber((_187_v9).length), _177_globalState)));
          _186_v8 = _186_v8;
          let _492_v250;
          _492_v250 = _dafny.Map.Empty.slice().updateUnsafe(_178_v0,_183_v5);
          let _493_v251;
          let _494_v252;
          let _495_v253;
          let _496_v254;
          let _out92;
          let _out93;
          let _out94;
          let _out95;
          let _outcollector23 = _module.__default.m0(_492_v250, (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))], _177_globalState);
          _out92 = _outcollector23[0];
          _out93 = _outcollector23[1];
          _out94 = _outcollector23[2];
          _out95 = _outcollector23[3];
          _493_v251 = _out92;
          _494_v252 = _out93;
          _495_v253 = _out94;
          _496_v254 = _out95;
          let _497_v255;
          _497_v255 = _dafny.Set.fromElements((((_187_v9).contains(!(!(_180_v2)))) ? ((_187_v9).get(!(!(_180_v2)))) : ((_dafny.ZERO).minus(new BigNumber((_185_v7).length)))));
          let _498_v256;
          _498_v256 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!((_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))]), (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))], _180_v2),new BigNumber((_497_v255).length));
          let _499_v257;
          _499_v257 = _dafny.Seq.of(_260_v72);
          _498_v256 = (_498_v256).update(_dafny.Seq.Concat((_499_v257)[_module.__default.safeIndex(_493_v251, new BigNumber((_499_v257).length))], _260_v72), _183_v5);
        } else {
          let _index61 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length));
          (_178_v0)[_index61] = _180_v2;
          let _500_v258;
          _500_v258 = _dafny.MultiSet.fromElements(_185_v7, _185_v7);
          let _501_v259;
          _501_v259 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(669)), ((_502_v1) => function (_503_i22) {
            return _502_v1;
          })(_179_v1)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(23)), ((_504_v8) => function (_505_i23) {
            return _504_v8;
          })(_186_v8)));
          let _nw66 = Array((new BigNumber(19)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = (((_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))]) ? (_180_v2) : ((_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))]));
          _nw66[(_dafny.ONE).toNumber()] = _180_v2;
          _nw66[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsPrefixOf(_260_v72, _dafny.Seq.of(_180_v2));
          _nw66[(new BigNumber(3)).toNumber()] = (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))];
          _nw66[(new BigNumber(4)).toNumber()] = (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))];
          _nw66[(new BigNumber(5)).toNumber()] = _180_v2;
          _nw66[(new BigNumber(6)).toNumber()] = (_260_v72)[_module.__default.safeIndex(new BigNumber((_500_v258).cardinality()), new BigNumber((_260_v72).length))];
          _nw66[(new BigNumber(7)).toNumber()] = (true) === ((_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))]);
          _nw66[(new BigNumber(8)).toNumber()] = _180_v2;
          _nw66[(new BigNumber(9)).toNumber()] = _180_v2;
          _nw66[(new BigNumber(10)).toNumber()] = (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))];
          _nw66[(new BigNumber(11)).toNumber()] = (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))];
          _nw66[(new BigNumber(12)).toNumber()] = _180_v2;
          _nw66[(new BigNumber(13)).toNumber()] = _180_v2;
          _nw66[(new BigNumber(14)).toNumber()] = (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))];
          _nw66[(new BigNumber(15)).toNumber()] = _module.__default.fm1((((_501_v259).contains(_425_v205)) ? ((_501_v259).get(_425_v205)) : (_185_v7)), _181_v3, _230_v42, _177_globalState);
          _nw66[(new BigNumber(16)).toNumber()] = (_module.D7.create_DC13((_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))])).dtor_cf17;
          _nw66[(new BigNumber(17)).toNumber()] = (_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))];
          _nw66[(new BigNumber(18)).toNumber()] = !((_178_v0)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_178_v0).length))]);
          _178_v0 = _nw66;
          _180_v2 = !(!(_dafny.Seq.IsProperPrefixOf(_185_v7, _185_v7)));
          _186_v8 = _186_v8;
          let _506_v260;
          _506_v260 = _dafny.Set.fromElements(new BigNumber(324), _179_v1);
          let _507_v261;
          let _nw67 = new _module.C0();
          _nw67.__ctor((_506_v260).IsDisjointFrom(_506_v260));
          _507_v261 = _nw67;
        }
        let _init11 = function (_508_i24) {
          return false;
        };
        let _nw68 = Array((new BigNumber(15)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw68.length); _i0_11++) {
          _nw68[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _178_v0 = _nw68;
        let _509_v262;
        let _nw69 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
        _509_v262 = _nw69;
        let _index62 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_509_v262).length));
        (_509_v262)[_index62] = (_362_v150).Merge(_362_v150);
        let _index63 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_509_v262).length));
        (_509_v262)[_index63] = (_362_v150).Merge(_362_v150);
        let _510_v263;
        _510_v263 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_425_v205, _425_v205),_183_v5);
        _510_v263 = _510_v263;
      }
      process.stdout.write(_dafny.toString((_178_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_179_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_180_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_181_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_183_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v6).equals(_dafny.Set.fromElements(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_185_v7, _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_186_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(365)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v10)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v11)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_193_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,new BigNumber(265)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v42).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_260_v72, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_357_i16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_361_v149).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_362_v150).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_425_v205, _dafny.Seq.of(new BigNumber(426), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_426_v206).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(426), _dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_427_v207).dtor_cf16).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(426), _dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_428_v208).dtor_cf16).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(426), _dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_429_v209).dtor_cf18).dtor_cf16).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(426), _dafny.ONE)))));
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
      return new _dafny.CodePoint('D'.codePointAt(0));
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
    static create_DC3(cf3, cf4, cf5) {
      let $dt = new D1(1);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
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
    get dtor_cf5() { return this.cf5; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4 && this.cf5 === other.cf5;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf1 === other.cf1;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.Set.Empty);
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
    static create_DC4(cf6) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf8, cf9, cf10) {
      let $dt = new D3(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC5(cf7) {
      let $dt = new D3(1);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC7(cf11) {
      let $dt = new D3(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC6(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC8(cf12) {
      let $dt = new D4(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12);
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf13) {
      let $dt = new D5(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC9" + "(" + this.cf13.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13);
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
    static create_DC11(cf15) {
      let $dt = new D6(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC10(cf14) {
      let $dt = new D6(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC11" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC10" + "(" + _dafny.toString(this.cf14) + ")";
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
        return other.$tag === 1 && this.cf14 === other.cf14;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC11(_dafny.ZERO);
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
    static create_DC13(cf17) {
      let $dt = new D7(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC12(cf16) {
      let $dt = new D7(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC14(cf18) {
      let $dt = new D7(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC13" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC12" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf17 === other.cf17;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC13(false);
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
    static create_DC16(cf20, cf21, cf22, cf23) {
      let $dt = new D8(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC17(cf24, cf25, cf26, cf27) {
      let $dt = new D8(1);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC15(cf19) {
      let $dt = new D8(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC18(cf28) {
      let $dt = new D8(3);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get is_DC18() { return this.$tag === 3; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC16" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC17" + "(" + this.cf24.toVerbatimString(true) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC15" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC16(false, _dafny.MultiSet.Empty, _dafny.Seq.of(), _dafny.Map.Empty);
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
    static create_DC20(cf30, cf31, cf32) {
      let $dt = new D9(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC19(cf29) {
      let $dt = new D9(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30 && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC20([], false, _dafny.ZERO);
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
    static create_DC22(cf34, cf35, cf36, cf37) {
      let $dt = new D10(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC21(cf33) {
      let $dt = new D10(1);
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC23(cf38) {
      let $dt = new D10(2);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC22" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35) && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC22(false, _dafny.ZERO, false, _dafny.ZERO);
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

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f0 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    set f0(value) {
      let _this = this;
      _this._f0 = value;
    };
    __ctor(f0) {
      let _this = this;
      (_this)._f0 = f0;
      return;
    }
    fm4(globalState) {
      let _this = this;
      let _source6 = _module.D1.create_DC2(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(840),_this.f0)).length))));
      if (_source6.is_DC2) {
        let _511___mcc_h0 = (_source6).cf2;
        let _512_cf2 = _511___mcc_h0;
        return new BigNumber(-299);
      } else if (_source6.is_DC3) {
        let _513___mcc_h1 = (_source6).cf3;
        let _514___mcc_h2 = (_source6).cf4;
        let _515___mcc_h3 = (_source6).cf5;
        let _516_cf5 = _515___mcc_h3;
        let _517_cf4 = _514___mcc_h2;
        let _518_cf3 = _513___mcc_h1;
        return new BigNumber((_dafny.Set.fromElements(_517_cf4)).length);
      } else {
        let _519___mcc_h4 = (_source6).cf1;
        let _520_cf1 = _519___mcc_h4;
        return (_dafny.ZERO).minus(new BigNumber((function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(617), new BigNumber(586))) {
            let _521_v0 = _compr_14;
            if (((new BigNumber(617)).isLessThanOrEqualTo(_521_v0)) && ((_521_v0).isLessThan(new BigNumber(586)))) {
              _coll14.push([(_521_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("vo")).length)),new BigNumber(129)]);
            }
          }
          return _coll14;
        }()).length));
      }
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((new BigNumber(39)).multipliedBy((new BigNumber(452))));
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(222)), function (_522_i0) {
        return new BigNumber(-394);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-784)), function (_523_i1) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("hpmca")).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("hs")).length), new BigNumber((_dafny.Seq.UnicodeFromString("gjhdka")).length), new BigNumber(-765)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(511)), function (_524_i2) {
        return new BigNumber(240);
      }), _dafny.Seq.of(new BigNumber(815), new BigNumber(633), (new BigNumber((_dafny.Set.fromElements(new BigNumber(-347))).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length), new BigNumber(390)))).IsSubsetOf(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(583))));
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
