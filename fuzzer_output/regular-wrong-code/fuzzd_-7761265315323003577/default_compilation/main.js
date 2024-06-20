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
      let _source0 = _module.D1.create_DC3(_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0))));
      if (_source0.is_DC4) {
        let _0___mcc_h0 = (_source0).cf7;
        let _1___mcc_h1 = (_source0).cf8;
        let _2___mcc_h2 = (_source0).cf9;
        let _3___mcc_h3 = (_source0).cf10;
        let _4___mcc_h4 = (_source0).cf11;
        let _5_cf11 = _4___mcc_h4;
        let _6_cf10 = _3___mcc_h3;
        let _7_cf9 = _2___mcc_h2;
        let _8_cf8 = _1___mcc_h1;
        let _9_cf7 = _0___mcc_h0;
        return (_7_cf9)[_module.__default.safeIndex(new BigNumber((_7_cf9).length), new BigNumber((_7_cf9).length))];
      } else {
        let _10___mcc_h5 = (_source0).cf6;
        let _11_cf6 = _10___mcc_h5;
        return new _dafny.CodePoint('n'.codePointAt(0));
      }
    };
    static fm1(globalState) {
      return _dafny.Set.fromElements((_dafny.ZERO).minus(((false) ? (new BigNumber((_dafny.MultiSet.fromElements(false, !(!(false)), false)).cardinality())) : (new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.MultiSet.fromElements(new BigNumber(-912))).Elements) {
          let _12_v0 = _compr_0;
          if ((_dafny.MultiSet.fromElements(new BigNumber(-912))).contains(_12_v0)) {
            _coll0.push([_module.__default.safeModuloInt(_12_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)),new BigNumber((function () {
              let _coll1 = new _dafny.Map();
              for (const _compr_1 of _dafny.IntegerRange(new BigNumber(676), new BigNumber(222))) {
                let _13_v1 = _compr_1;
                if (((new BigNumber(676)).isLessThanOrEqualTo(_13_v1)) && ((_13_v1).isLessThan(new BigNumber(222)))) {
                  _coll1.push([(_13_v1).plus(new BigNumber(-910)),true]);
                }
              }
              return _coll1;
            }()).length)]);
          }
        }
        return _coll0;
      }()).length))).length))).length)))), (_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),false)).length)).multipliedBy(new BigNumber(-907))), ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("lk")).length))).plus(new BigNumber(802)));
    };
    static fm2(p0, globalState) {
      return ((new BigNumber(-235)).minus(new BigNumber(-236))).multipliedBy(new BigNumber(663));
    };
    static fm3(globalState) {
      return false;
    };
    static fm4(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(false, false)), _dafny.Seq.of(false));
    };
    static fm9(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D11.create_DC33(new BigNumber(-562), !(false), new BigNumber(369), false, _dafny.Set.fromElements(true))).dtor_cf65, false))).cardinality()),new BigNumber(412))).length),new BigNumber(706))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true, false)).length))).length),new BigNumber(906)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(697)), function (_14_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length))).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(294), new BigNumber(803))) {
          let _15_v0 = _compr_2;
          if (((new BigNumber(294)).isLessThanOrEqualTo(_15_v0)) && ((_15_v0).isLessThan(new BigNumber(803)))) {
            _coll2.push([(_15_v0).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber((_dafny.Seq.of(true)).length))).length)))).cardinality())),new BigNumber(781)]);
          }
        }
        return _coll2;
      }()));
    };
    static fm12(p0, p1, p2, globalState) {
      return _module.D2.create_DC8(false, !((_dafny.Set.fromElements(new BigNumber(90))).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber(894), new BigNumber((_dafny.MultiSet.fromElements(false, true, false)).cardinality())))));
    };
    static fm17(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), function (_16_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("fqjluaqn")), _dafny.Seq.UnicodeFromString("pfpcdepv"));
    };
    static fm18(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(902)), function (_17_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("go")), _dafny.Seq.UnicodeFromString("ywfd"));
    };
    static fm19(p0, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(400), new BigNumber((_dafny.MultiSet.fromElements(false, false, !((_module.D0.create_DC1(!(true), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-171)), function (_18_i0) {
  return _dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(!(true))).length)), new BigNumber((_dafny.Seq.UnicodeFromString("qxcxo")).length)));
}))).cardinality()), new _dafny.CodePoint('h'.codePointAt(0)))).dtor_cf2), !(false))).cardinality()))).cardinality()), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), new BigNumber((_dafny.MultiSet.fromElements(false, true, false, true)).cardinality()))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(538)), function (_19_i1) {
        return new BigNumber(-397);
      })));
    };
    static fm20(p0, p1, p2, globalState) {
      return _module.D0.create_DC1((new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of (_dafny.Seq.of(function () {
    let _coll4 = new _dafny.Set();
    for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-648), new BigNumber(965))) {
      let _20_v1 = _compr_4;
      if (((new BigNumber(-648)).isLessThanOrEqualTo(_20_v1)) && ((_20_v1).isLessThan(new BigNumber(965)))) {
        _coll4.add(_module.__default.safeDivisionInt(_20_v1, new BigNumber(422)));
      }
    }
    return _coll4;
  }())).Elements) {
    let _21_v0 = _compr_3;
    if (_dafny.Seq.contains(_dafny.Seq.of(function () {
      let _coll5 = new _dafny.Set();
      for (const _compr_5 of _dafny.IntegerRange(new BigNumber(-648), new BigNumber(965))) {
        let _22_v1 = _compr_5;
        if (((new BigNumber(-648)).isLessThanOrEqualTo(_22_v1)) && ((_22_v1).isLessThan(new BigNumber(965)))) {
          _coll5.add(_module.__default.safeDivisionInt(_22_v1, new BigNumber(422)));
        }
      }
      return _coll5;
    }()), _21_v0)) {
      _coll3.push([_21_v0,true]);
    }
  }
  return _coll3;
}()).length)))).length)).isLessThanOrEqualTo(new BigNumber(-119)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qnlh"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), function (_23_i0) {
  return new _dafny.CodePoint('r'.codePointAt(0));
}))).length), new _dafny.CodePoint('r'.codePointAt(0)));
    };
    static fm25(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(761)), function (_24_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      });
    };
    static fm26(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-716), new BigNumber((_dafny.Set.fromElements(true)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(993)), function (_25_i0) {
        return (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length)));
      })), ((true) ? (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), function (_26_i1) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length))) : (_dafny.Seq.of(new BigNumber(90), new BigNumber(979)))));
    };
    static fm27(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("tqkyhgesq"),new BigNumber(-595))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("bk"),new BigNumber(753)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mrd"),new BigNumber(673)));
    };
    static fm29(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("yhclj")), (_module.D21.create_DC49(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("jcnirnr"), _dafny.Seq.UnicodeFromString("ttxfurlkd")))).dtor_cf88);
    };
    static fm30(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("xty");
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,(_module.D9.create_DC25(new BigNumber(-77), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("tussinc")).length), new BigNumber((_dafny.Seq.UnicodeFromString("tk")).length))).length)), new BigNumber((_dafny.Seq.of(false, true)).length), new _dafny.CodePoint('n'.codePointAt(0)))).dtor_cf47)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber(10))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-344)));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      if (!((false) || (false))) {
        return _module.D0.create_DC0(_dafny.Seq.of(!(false), false), new BigNumber(468));
      } else {
        return _module.D0.create_DC0(_dafny.Seq.of(true), new BigNumber(501));
      }
    };
    static fm33(p0, p1, p2, p3, globalState) {
      return (((_module.D22.create_DC51(function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(593),new BigNumber(-876))).Keys.Elements) {
    let _27_v0 = _compr_6;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(593),new BigNumber(-876))).contains(_27_v0)) {
      _coll6.push([(_27_v0).minus(new BigNumber(505)),true]);
    }
  }
  return _coll6;
}())).dtor_cf93).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(482),!(!(true))))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(381)), function (_28_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()));
      })).length)),true));
    };
    static fm34(p0, globalState) {
      return (_dafny.Set.fromElements(true)).Intersect((_dafny.Set.fromElements(!(false))).Union(_dafny.Set.fromElements(false, !(false), false, false, true)));
    };
    static fm35(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber(778), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(584),_dafny.Seq.of(true))).length)).minus(new BigNumber(-262)), new BigNumber(812), new BigNumber(881), (new BigNumber(222)).multipliedBy(new BigNumber(-684)));
    };
    static fm36(globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ekwnkte")).length), new BigNumber(380), new BigNumber(628), new BigNumber(-978), new BigNumber(812)), _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-503))), _dafny.MultiSet.fromElements((_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("xuckykvh")).length), new _dafny.CodePoint('f'.codePointAt(0)))).dtor_cf3, new BigNumber(815)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(377), new BigNumber(315)), _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-684))).length))), new BigNumber(826)), _dafny.MultiSet.fromElements(new BigNumber(150), new BigNumber(489)))))).Difference((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(326),!(false))).length)))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(868)), function (_29_i0) {
        return _dafny.Seq.UnicodeFromString("ijuq");
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(true))).length))))));
    };
    static fm37(p0, globalState) {
      return _module.D6.create_DC19(new BigNumber(543), _module.__default.safeDivisionInt(new BigNumber(691), new BigNumber(460)), (new BigNumber(95)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-772)))).length))).length)));
    };
    static fm38(globalState) {
      if (((!(true)) ? (true) : (!(!(false))))) {
        return _module.D5.create_DC16(_dafny.Set.fromElements(!(false), false, true));
      } else {
        return _module.D5.create_DC16(_dafny.Set.fromElements(true));
      }
    };
    static fm39(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(767)), function (_30_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })));
    };
    static fm40(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(12), new BigNumber(391))) {
          let _31_v0 = _compr_7;
          if (((new BigNumber(12)).isLessThanOrEqualTo(_31_v0)) && ((_31_v0).isLessThan(new BigNumber(391)))) {
            _coll7.push([(_31_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).length)),false]);
          }
        }
        return _coll7;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(670),false))).length)), new BigNumber(-354)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),!(!(false)))).length))).length)));
    };
    static fm41(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),true)));
    };
    static fm42(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("emmw")).length),(new BigNumber(749)).plus(new BigNumber(-156)));
    };
    static fm43(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-894)), function (_32_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(344)), function (_33_i1) {
          return new BigNumber(-599);
        });
      }), _dafny.Seq.of(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-910))), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, !(true))).length)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), function (_34_i2) {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(28)), function (_35_i3) {
          return _dafny.Seq.of(new BigNumber(375), new BigNumber(729));
        })).length));
      }));
    };
    static fm44(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).length),false)).length)), new BigNumber((_dafny.Set.fromElements(new BigNumber(499), new BigNumber(373), new BigNumber((_dafny.Set.fromElements(true)).length))).length)),_module.D9.create_DC25(new BigNumber(788), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("r")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), function (_36_i0) {
  return new _dafny.CodePoint('b'.codePointAt(0));
})).length),!(false))).length))).length), new BigNumber((_dafny.Set.fromElements(false, false)).length), new _dafny.CodePoint('u'.codePointAt(0))));
    };
    static fm45(p0, p1, globalState) {
      let _source1 = _module.D4.create_DC15(true, !(false), _dafny.Seq.UnicodeFromString("vypeoo"));
      if (_source1.is_DC13) {
        let _37___mcc_h0 = (_source1).cf25;
        let _38___mcc_h1 = (_source1).cf26;
        let _39___mcc_h2 = (_source1).cf27;
        let _40_cf27 = _39___mcc_h2;
        let _41_cf26 = _38___mcc_h1;
        let _42_cf25 = _37___mcc_h0;
        return _module.D4.create_DC13(_dafny.Seq.of(_40_cf27), _41_cf26, _40_cf27);
      } else if (_source1.is_DC14) {
        let _43___mcc_h3 = (_source1).cf28;
        let _44___mcc_h4 = (_source1).cf29;
        let _45___mcc_h5 = (_source1).cf30;
        let _46___mcc_h6 = (_source1).cf31;
        let _47_cf31 = _46___mcc_h6;
        let _48_cf30 = _45___mcc_h5;
        let _49_cf29 = _44___mcc_h4;
        let _50_cf28 = _43___mcc_h3;
        return _module.D4.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(43)), function (_51_i0) {
  return new BigNumber(652);
}), _50_cf28, _48_cf30);
      } else if (_source1.is_DC15) {
        let _52___mcc_h7 = (_source1).cf32;
        let _53___mcc_h8 = (_source1).cf33;
        let _54___mcc_h9 = (_source1).cf34;
        let _55_cf34 = _54___mcc_h9;
        let _56_cf33 = _53___mcc_h8;
        let _57_cf32 = _52___mcc_h7;
        return _module.D4.create_DC13(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_56_cf33,new BigNumber(34))).length), (_module.D1.create_DC4(_module.D0.create_DC0(_dafny.Seq.of(_57_cf32, _57_cf32, _56_cf33, true), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_56_cf33))).cardinality())), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_56_cf33,new BigNumber(-314))).length), _55_cf34, _56_cf33, new BigNumber(222))).dtor_cf11), _57_cf32, (_dafny.ZERO).minus(new BigNumber(-758)));
      } else {
        let _58___mcc_h10 = (_source1).cf24;
        let _59_cf24 = _58___mcc_h10;
        return _module.D4.create_DC13(_dafny.Seq.of(new BigNumber(368)), false, new BigNumber(-642));
      }
    };
    static fm46(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,!(false))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,!(true))));
    };
    static fm47(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(719)), function (_60_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }));
    };
    static fm48(globalState) {
      return (((false) ? (_dafny.MultiSet.fromElements(false)) : (_dafny.MultiSet.fromElements(false, true)))).Union((_dafny.MultiSet.fromElements(false)).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false, false))));
    };
    static fm49(p0, p1, p2, p3, globalState) {
      if ((_module.D4.create_DC13(_dafny.Seq.of(new BigNumber(-20), new BigNumber(680)), !(!(!(false))), new BigNumber(-244))).dtor_cf26) {
        return _module.D2.create_DC8(!(true), true);
      } else {
        return _module.D2.create_DC8(true, true);
      }
    };
    static fm50(p0, globalState) {
      return _module.D3.create_DC11(_module.D3.create_DC11(_module.D3.create_DC9(_dafny.Seq.of(new BigNumber(-215)))));
    };
    static fm51(globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), function (_61_i0) {
        return new BigNumber(-201);
      }), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(491))).length))), _dafny.Seq.of(new BigNumber(986)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("dyeiakrtw")).length), new BigNumber(-582), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(755), new BigNumber(-79), (_dafny.ZERO).minus(new BigNumber(-688)))).cardinality()))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(122)), function (_62_i1) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(23))).length);
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), function (_63_i2) {
        return (_module.D11.create_DC31(new BigNumber(((_module.D4.create_DC14(false, _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.ZERO, _dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).dtor_cf29).length), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).dtor_cf59;
      }));
    };
    static fm52(p0, globalState) {
      return _module.D14.create_DC37((_dafny.Set.fromElements(false, true)).IsSubsetOf(_dafny.Set.fromElements(false)), _dafny.Seq.UnicodeFromString("k"));
    };
    static fm53(p0, globalState) {
      return (_dafny.MultiSet.fromElements(_module.D1.create_DC3(_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-950)), function (_64_i0) {
        return _module.D1.create_DC3(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0))));
      }), _dafny.Seq.of(_module.D1.create_DC3(_dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))), _module.D1.create_DC3(_dafny.MultiSet.fromElements(new _dafny.CodePoint('l'.codePointAt(0))))))));
    };
    static fm54(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)));
    };
    static fm55(globalState) {
      return _module.D4.create_DC15((new BigNumber((_dafny.Seq.UnicodeFromString("iptf")).length)).isLessThanOrEqualTo(new BigNumber(-985)), !((false) && (true)), (_module.D4.create_DC15(true, false, _dafny.Seq.UnicodeFromString("uj"))).dtor_cf34);
    };
    static fm56(p0, p1, p2, globalState) {
      return _module.D11.create_DC31(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(804), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, !(true)))).cardinality())), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(316)), function (_65_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length), new BigNumber(827)))).length), (new BigNumber(518)).plus(new BigNumber(449)));
    };
    static Main(__noArgsParameter) {
      let _66_v0;
      _66_v0 = new _dafny.CodePoint('s'.codePointAt(0));
      let _67_v1;
      _67_v1 = _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)), _66_v0);
      let _68_v2;
      _68_v2 = new BigNumber(291);
      let _69_v3;
      _69_v3 = _dafny.Seq.of(_68_v2);
      let _70_v4;
      _70_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_69_v3).length),new BigNumber((_dafny.Seq.UnicodeFromString("vyifsmwe")).length));
      let _71_globalState;
      let _nw0 = new _module.GlobalState();
      _nw0.__ctor(true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(215)), function (_72_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }), _67_v1, true, _70_v4, new BigNumber(583));
      _71_globalState = _nw0;
      let _73_v5;
      _73_v5 = false;
      let _74_v6;
      _74_v6 = _module.D0.create_DC1(_73_v5, _module.__default.safeModuloInt(_68_v2, _68_v2), _module.__default.fm0(_68_v2, _module.__default.fm1(_71_globalState), _73_v5, _71_globalState));
      let _source2 = _74_v6;
      if (_source2.is_DC0) {
        let _75___mcc_h0 = (_source2).cf0;
        let _76___mcc_h1 = (_source2).cf1;
        let _77_cf1 = _76___mcc_h1;
        let _78_cf0 = _75___mcc_h0;
        let _79_v7;
        _79_v7 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_68_v2, _71_globalState),_66_v0);
        _79_v7 = _79_v7;
        let _80_v8;
        let _nw1 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _80_v8 = _nw1;
        let _index0 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length));
        (_80_v8)[_index0] = new BigNumber(-792);
        let _index1 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length));
        (_80_v8)[_index1] = (_77_cf1).minus(_module.__default.safeModuloInt(_77_cf1, new BigNumber((_78_cf0).length)));
        let _index2 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length));
        (_80_v8)[_index2] = (_80_v8)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length))];
        let _81_v9;
        _81_v9 = _module.D0.create_DC2(_module.D0.create_DC1(_73_v5, _68_v2, _66_v0));
        let _82_v10;
        _82_v10 = _dafny.Seq.of(_67_v1, _67_v1, _67_v1, _67_v1, _67_v1);
        let _83_v11;
        _83_v11 = _module.D0.create_DC0(_module.__default.fm4(_68_v2, new BigNumber((_82_v10).length), (_80_v8)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length))], _71_globalState), _77_cf1);
        let _84_v12;
        _84_v12 = _module.D0.create_DC2(_83_v11);
        let _85_v13;
        _85_v13 = _module.D0.create_DC2(_83_v11);
        let _pat_let_tv0 = _83_v11;
        let _source3 = ((_module.__default.fm3(_71_globalState)) ? (_81_v9) : (function (_pat_let0_0) {
          return function (_86_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_87_dt__update_hcf5_h0) {
                return _module.D0.create_DC2(_87_dt__update_hcf5_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_81_v9)));
        if (_source3.is_DC0) {
          let _88___mcc_h6 = (_source3).cf0;
          let _89___mcc_h7 = (_source3).cf1;
          let _90_cf1 = _89___mcc_h7;
          let _91_cf0 = _88___mcc_h6;
          let _92_v14;
          let _nw2 = new _module.C1();
          _nw2.__ctor(_68_v2);
          _92_v14 = _nw2;
          _73_v5 = (_91_cf0)[_module.__default.safeIndex((_dafny.ZERO).minus(_90_cf1), new BigNumber((_91_cf0).length))];
          _90_cf1 = (_dafny.ZERO).minus((_80_v8)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length))]);
          let _93_v15;
          let _nw3 = new _module.C5();
          _nw3.__ctor(_90_cf1, _90_cf1);
          _93_v15 = _nw3;
          let _94_v16;
          _94_v16 = _dafny.MultiSet.fromElements(_93_v15);
          let _95_v17;
          _95_v17 = _dafny.Map.Empty.slice().updateUnsafe(_68_v2,_73_v5);
          let _96_v18;
          _96_v18 = _dafny.Map.Empty.slice().updateUnsafe((((_95_v17).contains(_90_cf1)) ? ((_95_v17).get(_90_cf1)) : (!(_73_v5))),(_80_v8)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length))]);
          let _97_v19;
          let _init0 = ((_98_v5) => function (_99_i1) {
            return _98_v5;
          })(_73_v5);
          let _nw4 = Array((new BigNumber(26)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw4.length); _i0_0++) {
            _nw4[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _97_v19 = _nw4;
          let _100_v20;
          _100_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(547),_67_v1);
          let _101_v21;
          _101_v21 = _dafny.MultiSet.fromElements(_68_v2);
          let _102_v22;
          _102_v22 = _dafny.MultiSet.fromElements(_101_v21);
          let _103_v23;
          let _nw5 = new _module.C4();
          _nw5.__ctor((((_94_v16).contains(_93_v15)) ? ((_94_v16).get(_93_v15)) : (_68_v2)), new BigNumber(((_96_v18).update(_73_v5, new BigNumber((_67_v1).length))).length), _97_v19, _100_v20, _68_v2, _102_v22);
          _103_v23 = _nw5;
          let _104_v24;
          _104_v24 = _module.D6.create_DC18(_103_v23);
          let _105_v25;
          _105_v25 = _dafny.Map.Empty.slice().updateUnsafe(_80_v8,_69_v3);
          let _106_v26;
          _106_v26 = _dafny.Map.Empty.slice().updateUnsafe(_104_v24,_105_v25);
          (_71_globalState).f5 = new BigNumber((((((_106_v26).contains(_104_v24)) ? ((_106_v26).get(_104_v24)) : (_105_v25))).update(_80_v8, _69_v3)).length);
        } else if (_source3.is_DC1) {
          let _107___mcc_h8 = (_source3).cf2;
          let _108___mcc_h9 = (_source3).cf3;
          let _109___mcc_h10 = (_source3).cf4;
          let _110_cf4 = _109___mcc_h10;
          let _111_cf3 = _108___mcc_h9;
          let _112_cf2 = _107___mcc_h8;
          let _113_v27;
          let _nw6 = Array((new BigNumber(17)).toNumber()).fill([]);
          _113_v27 = _nw6;
          let _nw7 = Array((new BigNumber(10)).toNumber()).fill([]);
          _113_v27 = _nw7;
          let _114_v28;
          _114_v28 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_111_cf3),_112_cf2);
          _114_v28 = (_dafny.Map.Empty.slice().updateUnsafe((_80_v8)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length))],!(_112_cf2))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_77_cf1,_112_cf2));
          let _115_v29;
          _115_v29 = _module.D6.create_DC19(_68_v2, (_80_v8)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length))], new BigNumber(121));
          let _pat_let_tv1 = _80_v8;
          let _pat_let_tv2 = _80_v8;
          let _index3 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length));
          (_80_v8)[_index3] = _module.__default.fm2((function (_pat_let2_0) {
            return function (_116_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_117_dt__update_hcf41_h0) {
                  return _module.D6.create_DC19((_116_dt__update__tmp_h1).dtor_cf39, (_116_dt__update__tmp_h1).dtor_cf40, _117_dt__update_hcf41_h0);
                }(_pat_let3_0);
              }((_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_pat_let_tv1).length))]);
            }(_pat_let2_0);
          }(_115_v29)).dtor_cf41, _71_globalState);
          let _118_v30;
          let _nw8 = new _module.C6();
          _nw8.__ctor(_68_v2, _73_v5);
          _118_v30 = _nw8;
          let _119_v31;
          _119_v31 = _dafny.Set.fromElements(_118_v30, _118_v30, _118_v30, _118_v30, _118_v30);
          let _120_v32;
          _120_v32 = _dafny.Map.Empty.slice().updateUnsafe(_119_v31,_78_cf0);
          let _121_v33;
          _121_v33 = _dafny.Seq.of(_119_v31, _dafny.Set.fromElements(_118_v30, _118_v30, _118_v30));
          _120_v32 = (_120_v32).update((_121_v33)[_module.__default.safeIndex(_68_v2, new BigNumber((_121_v33).length))], _78_cf0);
        } else {
          let _122___mcc_h11 = (_source3).cf5;
          let _123_cf5 = _122___mcc_h11;
          let _124_v34;
          let _nw9 = Array((new BigNumber(27)).toNumber()).fill(false);
          _124_v34 = _nw9;
          let _index4 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_124_v34).length));
          (_124_v34)[_index4] = _73_v5;
          let _125_v35;
          _125_v35 = _dafny.Map.Empty.slice().updateUnsafe(_73_v5,_73_v5);
          let _index5 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_124_v34).length));
          (_124_v34)[_index5] = (((_125_v35).contains(_73_v5)) ? ((_125_v35).get(_73_v5)) : (_73_v5));
          let _126_v36;
          let _nw10 = new _module.C9();
          _nw10.__ctor((_80_v8)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length))]);
          _126_v36 = _nw10;
          let _127_v37;
          _127_v37 = _dafny.MultiSet.fromElements((_124_v34)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_124_v34).length))]);
          let _128_v38;
          _128_v38 = _dafny.Map.Empty.slice().updateUnsafe(_73_v5,(_dafny.ZERO).minus(new BigNumber((_127_v37).cardinality())));
          _68_v2 = (_68_v2).multipliedBy(_module.__default.safeModuloInt(_68_v2, (((_128_v38).contains(true)) ? ((_128_v38).get(true)) : (new BigNumber(858)))));
          let _129_v39;
          let _nw11 = Array((new BigNumber(25)).toNumber()).fill([]);
          _129_v39 = _nw11;
          let _130_v40;
          let _nw12 = Array((new BigNumber(6)).toNumber());
          _nw12[(_dafny.ZERO).toNumber()] = _129_v39;
          _nw12[(_dafny.ONE).toNumber()] = _129_v39;
          _nw12[(new BigNumber(2)).toNumber()] = _129_v39;
          _nw12[(new BigNumber(3)).toNumber()] = _129_v39;
          _nw12[(new BigNumber(4)).toNumber()] = ((_73_v5) ? (_129_v39) : (_129_v39));
          _nw12[(new BigNumber(5)).toNumber()] = _129_v39;
          _130_v40 = _nw12;
          let _index6 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_130_v40).length));
          (_130_v40)[_index6] = _129_v39;
          let _131_v41;
          let _nw13 = new _module.C10();
          _nw13.__ctor(_77_cf1);
          _131_v41 = _nw13;
          let _132_v42;
          _132_v42 = _dafny.Map.Empty.slice().updateUnsafe(_131_v41,new BigNumber(692));
          let _133_v43;
          _133_v43 = _dafny.Seq.of(_125_v35);
          let _index7 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_130_v40).length));
          let _rhs0 = _129_v39;
          let _rhs1 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.Concat(_dafny.Seq.update(_78_cf0, _module.__default.safeIndex(new BigNumber((_127_v37).cardinality()), new BigNumber((_78_cf0).length)), (_124_v34)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_124_v34).length))]), _module.__default.fm4((_80_v8)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length))], _68_v2, new BigNumber((_132_v42).length), _71_globalState))));
          let _rhs2 = _dafny.areEqual(_dafny.Seq.Concat(_133_v43, _module.__default.fm46(_71_globalState)), _133_v43);
          let _rhs3 = ((_80_v8)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_80_v8).length))]).plus((new BigNumber((_67_v1).length)).plus(_68_v2));
          let _lhs0 = _130_v40;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_130_v40).length));
          let _lhs2 = _71_globalState;
          let _lhs3 = _71_globalState;
          _lhs0[_lhs1] = _rhs0;
          _127_v37 = _rhs1;
          _lhs2.f0 = _rhs2;
          _lhs3.f5 = _rhs3;
        }
      } else if (_source2.is_DC1) {
        let _134___mcc_h2 = (_source2).cf2;
        let _135___mcc_h3 = (_source2).cf3;
        let _136___mcc_h4 = (_source2).cf4;
        let _137_cf4 = _136___mcc_h4;
        let _138_cf3 = _135___mcc_h3;
        let _139_cf2 = _134___mcc_h2;
        _73_v5 = _73_v5;
        let _140_v44;
        let _nw14 = Array((_dafny.ONE).toNumber());
        _140_v44 = _nw14;
        let _141_v45;
        _141_v45 = _module.D20.create_DC46(_140_v44);
        let _142_v46;
        _142_v46 = _dafny.Map.Empty.slice().updateUnsafe((_141_v45).dtor_cf84,_138_cf3);
        _142_v46 = (_142_v46).update(_140_v44, _68_v2);
        (_71_globalState).f0 = _139_cf2;
        (_71_globalState).f5 = _138_cf3;
      } else {
        let _143___mcc_h5 = (_source2).cf5;
        let _144_cf5 = _143___mcc_h5;
        let _145_v47;
        let _nw15 = Array((new BigNumber(14)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = new BigNumber(563);
        _nw15[(_dafny.ONE).toNumber()] = _68_v2;
        _nw15[(new BigNumber(2)).toNumber()] = _68_v2;
        _nw15[(new BigNumber(3)).toNumber()] = _68_v2;
        _nw15[(new BigNumber(4)).toNumber()] = _68_v2;
        _nw15[(new BigNumber(5)).toNumber()] = new BigNumber((_70_v4).length);
        _nw15[(new BigNumber(6)).toNumber()] = new BigNumber(-185);
        _nw15[(new BigNumber(7)).toNumber()] = _68_v2;
        _nw15[(new BigNumber(8)).toNumber()] = _68_v2;
        _nw15[(new BigNumber(9)).toNumber()] = _68_v2;
        _nw15[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(_68_v2);
        _nw15[(new BigNumber(11)).toNumber()] = new BigNumber(-959);
        _nw15[(new BigNumber(12)).toNumber()] = new BigNumber((_69_v3).length);
        _nw15[(new BigNumber(13)).toNumber()] = new BigNumber((_67_v1).length);
        _145_v47 = _nw15;
        let _146_v48;
        let _init1 = function (_147_i2) {
          return true;
        };
        let _nw16 = Array((new BigNumber(5)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw16.length); _i0_1++) {
          _nw16[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _146_v48 = _nw16;
        let _148_v49;
        _148_v49 = _dafny.Seq.of(false);
        let _149_v50;
        _149_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_148_v49).length),_dafny.Seq.UnicodeFromString("vuxo"));
        let _150_v51;
        let _nw17 = new _module.C7();
        _nw17.__ctor(_145_v47, _145_v47, _module.__default.fm2(_68_v2, _71_globalState), _146_v48, _149_v50);
        _150_v51 = _nw17;
        let _rhs4 = _68_v2;
        let _rhs5 = (_68_v2).minus(_68_v2);
        let _lhs4 = _71_globalState;
        let _lhs5 = _71_globalState;
        _lhs4.f5 = _rhs4;
        _lhs5.f5 = _rhs5;
        let _151_v52;
        let _nw18 = new _module.C0();
        _nw18.__ctor(_146_v48, (new BigNumber((_69_v3).length)).isLessThanOrEqualTo((_150_v51).fm14(_71_globalState)));
        _151_v52 = _nw18;
        _151_v52 = _151_v52;
        let _152_v53;
        let _nw19 = new _module.C8();
        _nw19.__ctor();
        _152_v53 = _nw19;
      }
      (_71_globalState).f5 = _module.__default.safeDivisionInt(_68_v2, _68_v2);
      let _153_v54;
      _153_v54 = _dafny.Map.Empty.slice().updateUnsafe(_73_v5,_module.__default.fm2(_68_v2, _71_globalState));
      let _154_v55;
      _154_v55 = _dafny.Map.Empty.slice().updateUnsafe(_153_v54,_dafny.MultiSet.fromElements(!(_73_v5)));
      _154_v55 = (_154_v55).Merge(_154_v55);
      _73_v5 = ((_73_v5) ? (_73_v5) : (_73_v5));
      let _155_v56;
      _155_v56 = _dafny.Seq.of(_73_v5, _73_v5);
      let _156_v57;
      _156_v57 = _dafny.Set.fromElements((_dafny.ZERO).minus(_68_v2), _module.__default.fm2(_68_v2, _71_globalState));
      let _157_v58;
      _157_v58 = _dafny.Set.fromElements(_73_v5, _73_v5);
      let _158_v59;
      let _nw20 = Array((new BigNumber(7)).toNumber()).fill(false);
      _158_v59 = _nw20;
      let _159_v60;
      _159_v60 = _dafny.Map.Empty.slice().updateUnsafe(_68_v2,_67_v1);
      let _160_v61;
      let _nw21 = new _module.C3();
      _nw21.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_68_v2,(_dafny.ZERO).minus(_68_v2)), _158_v59, _159_v60, _68_v2);
      _160_v61 = _nw21;
      let _161_v62;
      _161_v62 = _dafny.Map.Empty.slice().updateUnsafe(_69_v3,_160_v61);
      let _162_v63;
      _162_v63 = _dafny.Set.fromElements(_66_v0);
      let _163_v64;
      _163_v64 = _dafny.MultiSet.fromElements(_157_v58);
      let _164_v65;
      let _nw22 = Array((new BigNumber(25)).toNumber());
      _nw22[(_dafny.ZERO).toNumber()] = _module.__default.fm2(_68_v2, _71_globalState);
      _nw22[(_dafny.ONE).toNumber()] = new BigNumber((_155_v56).length);
      _nw22[(new BigNumber(2)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(3)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(4)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(5)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(6)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(7)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(8)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(9)).toNumber()] = new BigNumber(-22);
      _nw22[(new BigNumber(10)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(11)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(12)).toNumber()] = new BigNumber((_156_v57).length);
      _nw22[(new BigNumber(13)).toNumber()] = new BigNumber((_157_v58).length);
      _nw22[(new BigNumber(14)).toNumber()] = new BigNumber(-993);
      _nw22[(new BigNumber(15)).toNumber()] = new BigNumber((_161_v62).length);
      _nw22[(new BigNumber(16)).toNumber()] = new BigNumber((_70_v4).length);
      _nw22[(new BigNumber(17)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(18)).toNumber()] = _68_v2;
      _nw22[(new BigNumber(19)).toNumber()] = new BigNumber(253);
      _nw22[(new BigNumber(20)).toNumber()] = (_160_v61).f6;
      _nw22[(new BigNumber(21)).toNumber()] = new BigNumber((_162_v63).length);
      _nw22[(new BigNumber(22)).toNumber()] = (_160_v61).f6;
      _nw22[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ibypqjp")).length);
      _nw22[(new BigNumber(24)).toNumber()] = (((_163_v64).contains(_157_v58)) ? ((_163_v64).get(_157_v58)) : ((_160_v61).f6));
      _164_v65 = _nw22;
      let _165_v66;
      _165_v66 = _dafny.Map.Empty.slice().updateUnsafe(_164_v65,!(_73_v5));
      let _166_v67;
      _166_v67 = _module.D10.create_DC28(_165_v66);
      _166_v67 = _166_v67;
      let _167_v68;
      _167_v68 = _dafny.Map.Empty.slice().updateUnsafe(_73_v5,_160_v61.f9);
      let _168_v69;
      let _169_v70;
      let _out0;
      let _out1;
      let _outcollector0 = (_160_v61).m6(_73_v5, _dafny.Seq.Concat(_dafny.Seq.of(_73_v5, _73_v5), _module.__default.fm4((_160_v61).f6, _68_v2, new BigNumber((_167_v68).length), _71_globalState)), _68_v2, _71_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _168_v69 = _out0;
      _169_v70 = _out1;
      (_71_globalState).f2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_67_v1, _67_v1), _67_v1);
      let _170_v71;
      let _171_v72;
      let _out2;
      let _out3;
      let _outcollector1 = (_160_v61).m6(_73_v5, _module.__default.fm4((_160_v61).f6, _68_v2, (_160_v61).f6, _71_globalState), new BigNumber((_dafny.Seq.of(_73_v5, _169_v70, _73_v5, _73_v5)).length), _71_globalState);
      _out2 = _outcollector1[0];
      _out3 = _outcollector1[1];
      _170_v71 = _out2;
      _171_v72 = _out3;
      let _172_i3;
      _172_i3 = _dafny.ZERO;
      L0: {
        while (!(_171_v72)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_172_i3)) {
              break L0;
            }
            _172_i3 = (_172_i3).plus(_dafny.ONE);
            let _173_v73;
            let _nw23 = new _module.C6();
            _nw23.__ctor((_160_v61).f6, _169_v70);
            _173_v73 = _nw23;
            _173_v73 = _173_v73;
            let _174_v74;
            let _175_v75;
            let _out4;
            let _out5;
            let _outcollector2 = (_160_v61).m6(((_dafny.ZERO).minus((_160_v61).f6)).isLessThanOrEqualTo((_160_v61).f6), _dafny.Seq.Concat(_155_v56, _155_v56), _module.__default.safeDivisionInt(_68_v2, (_dafny.ZERO).minus((_160_v61).f6)), _71_globalState);
            _out4 = _outcollector2[0];
            _out5 = _outcollector2[1];
            _174_v74 = _out4;
            _175_v75 = _out5;
            let _176_v76;
            _176_v76 = _dafny.MultiSet.fromElements((_156_v57).IsSubsetOf(_156_v57), true, _73_v5, _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), function (_177_i4) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            }), _67_v1), _171_v72);
            _176_v76 = (_176_v76).Intersect(_176_v76);
            _69_v3 = _module.__default.fm40(_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.FromArray(_155_v56)).cardinality()), (_160_v61).f6), _157_v58, _71_globalState);
          }
        }
      }
      _73_v5 = _module.__default.fm3(_71_globalState);
      let _hi0 = _module.__default.safeModuloInt((_160_v61).f6, new BigNumber((_156_v57).length));
      for (let _178_i5 = _68_v2; _178_i5.isLessThan(_hi0); _178_i5 = _178_i5.plus(_dafny.ONE)) {
        let _179_v77;
        let _180_v78;
        let _out6;
        let _out7;
        let _outcollector3 = (_160_v61).m6(_171_v72, _dafny.Seq.Concat(_dafny.Seq.of(_169_v70), _155_v56), new BigNumber(481), _71_globalState);
        _out6 = _outcollector3[0];
        _out7 = _outcollector3[1];
        _179_v77 = _out6;
        _180_v78 = _out7;
        if (!(_169_v70)) {
          let _181_v79;
          _181_v79 = _dafny.Map.Empty.slice().updateUnsafe(false,_180_v78);
          let _182_v80;
          _182_v80 = _dafny.Set.fromElements(_181_v79);
          let _183_v81;
          let _nw24 = Array((new BigNumber(8)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(175)), ((_184_v0) => function (_185_i6) {
            return _184_v0;
          })(_66_v0));
          _nw24[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("kosnitju");
          _nw24[(new BigNumber(2)).toNumber()] = _module.__default.fm17((_160_v61).f6, new BigNumber((_67_v1).length), _71_globalState);
          _nw24[(new BigNumber(3)).toNumber()] = _67_v1;
          _nw24[(new BigNumber(4)).toNumber()] = _67_v1;
          _nw24[(new BigNumber(5)).toNumber()] = _67_v1;
          _nw24[(new BigNumber(6)).toNumber()] = _67_v1;
          _nw24[(new BigNumber(7)).toNumber()] = _67_v1;
          _183_v81 = _nw24;
          let _186_v82;
          _186_v82 = _module.D10.create_DC29((_160_v61).f6, _183_v81, _180_v78, _178_i5);
          let _rhs6 = _164_v65;
          let _rhs7 = _182_v80;
          let _rhs8 = (((_181_v79).contains(_169_v70)) ? ((_181_v79).get(_169_v70)) : ((_186_v82).dtor_cf55));
          let _lhs6 = _71_globalState;
          _164_v65 = _rhs6;
          _182_v80 = _rhs7;
          _lhs6.f0 = _rhs8;
          let _187_v83;
          _187_v83 = _160_v61.f9;
          let _188_v84;
          let _nw25 = new _module.C7();
          _nw25.__ctor(_164_v65, _164_v65, _module.__default.safeDivisionInt((_160_v61).f6, new BigNumber(((_168_v69)[_module.__default.safeIndex((_160_v61).f6, new BigNumber((_168_v69).length))]).length)), (_187_v83), _dafny.Map.Empty.slice().updateUnsafe(_68_v2,_67_v1));
          _188_v84 = _nw25;
          let _189_v85;
          let _nw26 = new _module.C1();
          _nw26.__ctor((new BigNumber(684)).minus((_dafny.ZERO).minus(new BigNumber((_67_v1).length))));
          _189_v85 = _nw26;
          _189_v85 = _189_v85;
          let _190_v86;
          _190_v86 = _module.D11.create_DC31((_160_v61).f6, (_dafny.ZERO).minus((_160_v61).f6));
          _190_v86 = _module.__default.fm56(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-459)), ((_191_v0) => function (_192_i7) {
            return _191_v0;
          })(_66_v0)), _73_v5, _67_v1, _71_globalState);
          let _193_v87;
          let _nw27 = new _module.C1();
          _nw27.__ctor(new BigNumber((_dafny.Seq.update(_69_v3, _module.__default.safeIndex((_160_v61).f6, new BigNumber((_69_v3).length)), _178_i5)).length));
          _193_v87 = _nw27;
          let _194_v88;
          _194_v88 = _dafny.Seq.of(_193_v87, _193_v87, _193_v87, _193_v87);
          let _195_v89;
          let _nw28 = Array((new BigNumber(22)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = _193_v87;
          _nw28[(_dafny.ONE).toNumber()] = _193_v87;
          _nw28[(new BigNumber(2)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(3)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(4)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(5)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(6)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(7)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(8)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(9)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(10)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(11)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(12)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(13)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(14)).toNumber()] = (_194_v88)[_module.__default.safeIndex(_module.__default.fm2((_193_v87).f6, _71_globalState), new BigNumber((_194_v88).length))];
          _nw28[(new BigNumber(15)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(16)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(17)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(18)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(19)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(20)).toNumber()] = _193_v87;
          _nw28[(new BigNumber(21)).toNumber()] = _193_v87;
          _195_v89 = _nw28;
          let _index8 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_195_v89).length));
          (_195_v89)[_index8] = _193_v87;
          let _196_v90;
          _196_v90 = _module.D2.create_DC6((_160_v61).f6, _188_v84.f12);
          let _197_v91;
          _197_v91 = _dafny.MultiSet.fromElements((_196_v90).dtor_cf14, _188_v84.f12);
          let _198_v92;
          _198_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_70_v4).update((_160_v61).f6, (_160_v61).f6)).length),!(_73_v5));
          let _index9 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_195_v89).length));
          let _rhs9 = _193_v87;
          let _rhs10 = ((new BigNumber((_198_v92).length)).plus(_module.__default.fm2((_160_v61).f6, _71_globalState))).isLessThanOrEqualTo((_160_v61).f6);
          let _rhs11 = _155_v56;
          let _rhs12 = (_67_v1)[_module.__default.safeIndex(_68_v2, new BigNumber((_67_v1).length))];
          let _rhs13 = _197_v91;
          let _lhs7 = _195_v89;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_195_v89).length));
          _lhs7[_lhs8] = _rhs9;
          _171_v72 = _rhs10;
          _155_v56 = _rhs11;
          _66_v0 = _rhs12;
          _197_v91 = _rhs13;
        } else {
          (_160_v61).f9 = _160_v61.f9;
          let _199_v93;
          _199_v93 = _module.D4.create_DC13(_69_v3, !(_169_v70), (_160_v61).f6);
          let _200_v94;
          _200_v94 = _module.D14.create_DC37(_73_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(743)), function (_201_i9) {
  return new _dafny.CodePoint('o'.codePointAt(0));
}));
          _199_v93 = _module.D4.create_DC13(((_169_v70) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(569)), ((_202_v70) => function (_203_i8) {
  return new BigNumber((_dafny.MultiSet.fromElements(_202_v70)).cardinality());
})(_169_v70))) : (_69_v3)), (_200_v94).dtor_cf70, ((_dafny.ZERO).minus((_160_v61).f6)).multipliedBy((_160_v61).f6));
          let _204_v95;
          let _nw29 = Array((new BigNumber(21)).toNumber()).fill(_module.D2.Default());
          _204_v95 = _nw29;
          let _205_v96;
          _205_v96 = _dafny.Map.Empty.slice().updateUnsafe(_178_i5,_180_v78);
          let _206_v97;
          _206_v97 = _dafny.Map.Empty.slice().updateUnsafe(_205_v96,_169_v70);
          let _207_v98;
          _207_v98 = _module.D2.create_DC6(new BigNumber((_206_v97).length), _164_v65);
          let _index10 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_204_v95).length));
          (_204_v95)[_index10] = _207_v98;
          let _index11 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_204_v95).length));
          (_204_v95)[_index11] = _module.D2.create_DC6(_178_i5, _164_v65);
          _155_v56 = _dafny.Seq.Concat(_155_v56, _155_v56);
          _66_v0 = _66_v0;
        }
        let _208_v99;
        let _init2 = ((_209_v0) => function (_210_i10) {
          return _209_v0;
        })(_66_v0);
        let _nw30 = Array((new BigNumber(24)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw30.length); _i0_2++) {
          _nw30[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _208_v99 = _nw30;
        let _index12 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_208_v99).length));
        (_208_v99)[_index12] = _66_v0;
        let _index13 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_208_v99).length));
        (_208_v99)[_index13] = _66_v0;
        let _index14 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_164_v65).length));
        (_164_v65)[_index14] = (_160_v61).f6;
        let _211_v100;
        let _nw31 = Array((new BigNumber(11)).toNumber()).fill(_module.D1.Default());
        _211_v100 = _nw31;
        let _212_v101;
        _212_v101 = _module.D1.create_DC4(function (_pat_let4_0) {
  return function (_213_dt__update__tmp_h2) {
    return function (_pat_let5_0) {
      return function (_214_dt__update_hcf1_h0) {
        return _module.D0.create_DC0((_213_dt__update__tmp_h2).dtor_cf0, _214_dt__update_hcf1_h0);
      }(_pat_let5_0);
    }(new BigNumber(-120));
  }(_pat_let4_0);
}(_module.D0.create_DC0(_155_v56, _178_i5)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_171_v72,_171_v72)).length), _67_v1, _169_v70, _68_v2);
        let _index15 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_211_v100).length));
        (_211_v100)[_index15] = _212_v101;
        let _index16 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_164_v65).length));
        (_164_v65)[_index16] = _178_i5;
        let _215_v102;
        _215_v102 = _dafny.Map.Empty.slice().updateUnsafe((_208_v99)[_module.__default.safeIndex(new BigNumber(630), new BigNumber((_208_v99).length))],_73_v5);
        let _index17 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_164_v65).length));
        let _index18 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_211_v100).length));
        let _index19 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_164_v65).length));
        let _rhs14 = _module.__default.fm2(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(101)), ((_216_v0) => function (_217_i11) {
          return _216_v0;
        })(_66_v0)), _dafny.Seq.UnicodeFromString("r"))).length), _71_globalState);
        let _rhs15 = _module.__default.fm2(_178_i5, _71_globalState);
        let _rhs16 = _212_v101;
        let _rhs17 = _module.__default.safeDivisionInt(new BigNumber((_215_v102).length), _68_v2);
        let _rhs18 = _178_i5;
        let _lhs9 = _164_v65;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_164_v65).length));
        let _lhs11 = _211_v100;
        let _lhs12 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_211_v100).length));
        let _lhs13 = _164_v65;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_164_v65).length));
        _lhs9[_lhs10] = _rhs14;
        _68_v2 = _rhs15;
        _lhs11[_lhs12] = _rhs16;
        _lhs13[_lhs14] = _rhs17;
        _68_v2 = _rhs18;
      }
      let _index20 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_164_v65).length));
      (_164_v65)[_index20] = _module.__default.safeDivisionInt(_68_v2, new BigNumber((_dafny.Set.fromElements(_171_v72)).length));
      let _index21 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_164_v65).length));
      (_164_v65)[_index21] = (_160_v61).f6;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_164_v65).length))) {
        let _218_i12 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_218_i12)) && ((_218_i12).isLessThan(new BigNumber((_164_v65).length))))) {
          (_164_v65)[(_218_i12)] = _module.__default.safeModuloInt(_218_i12, new BigNumber((_153_v54).length));
        }
      }
      let _arr0 = _160_v61.f9;
      let _index22 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_160_v61.f9).length));
      _arr0[_index22] = !(_171_v72) || (_171_v72);
      let _arr1 = _160_v61.f9;
      let _index23 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_160_v61.f9).length));
      let _rhs19 = (new BigNumber(309)).isLessThan((_164_v65)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_164_v65).length))]);
      let _rhs20 = _73_v5;
      let _lhs15 = _160_v61.f9;
      let _lhs16 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_160_v61.f9).length));
      _171_v72 = _rhs19;
      _lhs15[_lhs16] = _rhs20;
      let _219_v103;
      _219_v103 = _dafny.MultiSet.fromElements(_68_v2, (_160_v61).f6);
      let _220_v104;
      let _221_v105;
      let _222_v106;
      let _223_v107;
      let _out8;
      let _out9;
      let _out10;
      let _out11;
      let _outcollector4 = (_160_v61).m7((_219_v103).Union(_module.__default.fm19(new BigNumber(534), _71_globalState)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_169_v70, _169_v70), _155_v56)).length), _71_globalState);
      _out8 = _outcollector4[0];
      _out9 = _outcollector4[1];
      _out10 = _outcollector4[2];
      _out11 = _outcollector4[3];
      _220_v104 = _out8;
      _221_v105 = _out9;
      _222_v106 = _out10;
      _223_v107 = _out11;
      let _224_i13;
      _224_i13 = _dafny.ZERO;
      L1: {
        while (_module.__default.fm3(_71_globalState)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_224_i13)) {
              break L1;
            }
            _224_i13 = (_224_i13).plus(_dafny.ONE);
            let _225_v108;
            let _nw32 = Array((new BigNumber(17)).toNumber());
            _225_v108 = _nw32;
            let _226_v109;
            let _nw33 = new _module.C0();
            _nw33.__ctor(_160_v61.f9, _73_v5);
            _226_v109 = _nw33;
            let _index24 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_225_v108).length));
            (_225_v108)[_index24] = _226_v109;
            let _index25 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_225_v108).length));
            (_225_v108)[_index25] = ((_169_v70) ? (_226_v109) : (_226_v109));
            _68_v2 = (_dafny.ZERO).minus((_160_v61).f6);
            _169_v70 = _73_v5;
            if (_171_v72) {
              let _227_v110;
              let _228_v111;
              let _229_v112;
              let _230_v113;
              let _out12;
              let _out13;
              let _out14;
              let _out15;
              let _outcollector5 = (_160_v61).m7(_219_v103, ((_171_v72) ? ((_164_v65)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_164_v65).length))]) : ((_164_v65)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_164_v65).length))])), _71_globalState);
              _out12 = _outcollector5[0];
              _out13 = _outcollector5[1];
              _out14 = _outcollector5[2];
              _out15 = _outcollector5[3];
              _227_v110 = _out12;
              _228_v111 = _out13;
              _229_v112 = _out14;
              _230_v113 = _out15;
              let _231_v114;
              _231_v114 = _dafny.Seq.of(_164_v65, _164_v65);
              _231_v114 = _231_v114;
              let _232_v115;
              _232_v115 = _dafny.MultiSet.fromElements(false, (_155_v56)[_module.__default.safeIndex(_230_v113, new BigNumber((_155_v56).length))]);
              let _233_v116;
              _233_v116 = _dafny.Map.Empty.slice().updateUnsafe(_73_v5,_226_v109.f20);
              _223_v107 = ((new BigNumber((_232_v115).cardinality())).multipliedBy(_228_v111)).minus((new BigNumber((_67_v1).length)).plus(new BigNumber((_233_v116).length)));
              let _234_v117;
              _234_v117 = _dafny.Map.Empty.slice().updateUnsafe(_74_v6,true);
              _234_v117 = (_234_v117).update(_74_v6, (_160_v61.f9)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_160_v61.f9).length))]);
              _220_v104 = _module.__default.fm0(_module.__default.fm2((_160_v61).f6, _71_globalState), (_module.__default.fm1(_71_globalState)).Union(_156_v57), _dafny.areEqual(_67_v1, _dafny.Seq.UnicodeFromString("eiyyqrevp")), _71_globalState);
            } else {
              _169_v70 = !(_68_v2).isEqualTo(_module.__default.fm2(new BigNumber(-764), _71_globalState));
              _223_v107 = (_dafny.ZERO).minus((_160_v61).f6);
              let _rhs21 = (_226_v109.f20) || (_171_v72);
              let _rhs22 = _160_v61;
              let _rhs23 = ((_164_v65)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_164_v65).length))]).isLessThanOrEqualTo((_68_v2).minus((_69_v3)[_module.__default.safeIndex((_164_v65)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_164_v65).length))], new BigNumber((_69_v3).length))]));
              let _rhs24 = _223_v107;
              let _rhs25 = _156_v57;
              let _lhs17 = _71_globalState;
              let _lhs18 = _71_globalState;
              _73_v5 = _rhs21;
              _160_v61 = _rhs22;
              _lhs17.f0 = _rhs23;
              _lhs18.f5 = _rhs24;
              _156_v57 = _rhs25;
              let _index26 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_164_v65).length));
              (_164_v65)[_index26] = _68_v2;
              _155_v56 = _155_v56;
            }
          }
        }
      }
      process.stdout.write(_dafny.toString(_66_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_67_v1, _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_68_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_69_v3, _dafny.Seq.of(new BigNumber(-2), new BigNumber(-354), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_70_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(8)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_71_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_71_globalState).f1, _dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_71_globalState.f2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(8)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_71_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_73_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_74_v6).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_74_v6).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_74_v6).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v54).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(663)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v55).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(663)),_dafny.MultiSet.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_155_v56, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v57).equals(_dafny.Set.fromElements(new BigNumber(-291), new BigNumber(663)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_v58).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v59)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v59)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v59)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v59)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v60).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(291),_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v61.f9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v61.f9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v61.f9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v61.f9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_160_v61).f10).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(291),_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v61).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_161_v62).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_v63).equals(_dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v64).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v65)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_165_v66).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_166_v67).dtor_cf52).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_167_v68).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_168_v69, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("yhclj"), _dafny.Seq.UnicodeFromString("jcnirnr"), _dafny.Seq.UnicodeFromString("ttxfurlkd")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_169_v70));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_170_v71, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("yhclj"), _dafny.Seq.UnicodeFromString("jcnirnr"), _dafny.Seq.UnicodeFromString("ttxfurlkd")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_171_v72));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_172_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_v103).equals(_dafny.MultiSet.fromElements(new BigNumber(291), new BigNumber(291)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_220_v104));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_221_v105));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v106)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v106)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v106)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v106)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v106)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v106)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v106)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_223_v107));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_224_i13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0, cf1) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC1(cf2, cf3, cf4) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC2(cf5) {
      let $dt = new D0(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0) && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(_dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC4(cf7, cf8, cf9, cf10, cf11) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + this.cf9.toVerbatimString(true) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_module.D0.Default(), _dafny.ZERO, _dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO);
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
    static create_DC6(cf13, cf14) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC7(cf15, cf16, cf17, cf18) {
      let $dt = new D2(1);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC8(cf19, cf20) {
      let $dt = new D2(2);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC5(cf12) {
      let $dt = new D2(3);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf19 === other.cf19 && this.cf20 === other.cf20;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.ZERO, []);
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
    static create_DC10(cf22) {
      let $dt = new D3(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC9(cf21) {
      let $dt = new D3(1);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC11(cf23) {
      let $dt = new D3(2);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.Map.Empty);
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
    static create_DC13(cf25, cf26, cf27) {
      let $dt = new D4(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC14(cf28, cf29, cf30, cf31) {
      let $dt = new D4(1);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC15(cf32, cf33, cf34) {
      let $dt = new D4(2);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC12(cf24) {
      let $dt = new D4(3);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + this.cf34.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25) && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf32 === other.cf32 && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(_dafny.Seq.of(), false, _dafny.ZERO);
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
    static create_DC17(cf36, cf37) {
      let $dt = new D5(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC16(cf35) {
      let $dt = new D5(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC17(_dafny.Seq.of(), false);
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
    static create_DC19(cf39, cf40, cf41) {
      let $dt = new D6(0);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC18(cf38) {
      let $dt = new D6(1);
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC20(cf42) {
      let $dt = new D6(2);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf39, other.cf39) && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf38 === other.cf38;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC19(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC22() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC21(cf43) {
      let $dt = new D7(1);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC22";
      } else if (this.$tag === 1) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf43) + ")";
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
        return other.$tag === 1 && this.cf43 === other.cf43;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC22();
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
    static create_DC23(cf44) {
      let $dt = new D8(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf44 === other.cf44;
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC25(cf46, cf47, cf48, cf49) {
      let $dt = new D9(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC26(cf50) {
      let $dt = new D9(1);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC24(cf45) {
      let $dt = new D9(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC27(cf51) {
      let $dt = new D9(3);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get is_DC27() { return this.$tag === 3; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48) && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf50 === other.cf50;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC25(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC29(cf53, cf54, cf55, cf56) {
      let $dt = new D10(0);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
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
    get dtor_cf56() { return this.cf56; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf53, other.cf53) && this.cf54 === other.cf54 && this.cf55 === other.cf55 && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29(_dafny.ZERO, [], false, _dafny.ZERO);
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
    static create_DC31(cf58, cf59) {
      let $dt = new D11(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC32(cf60, cf61) {
      let $dt = new D11(1);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC33(cf62, cf63, cf64, cf65, cf66) {
      let $dt = new D11(2);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC30(cf57) {
      let $dt = new D11(3);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get is_DC30() { return this.$tag === 3; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf60) + ", " + this.cf61.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf60, other.cf60) && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf62, other.cf62) && this.cf63 === other.cf63 && _dafny.areEqual(this.cf64, other.cf64) && this.cf65 === other.cf65 && _dafny.areEqual(this.cf66, other.cf66);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC31(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC34(cf67) {
      let $dt = new D12(0);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf67, other.cf67);
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
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC35(cf68) {
      let $dt = new D13(0);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68);
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
    static create_DC37(cf70, cf71) {
      let $dt = new D14(0);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC36(cf69) {
      let $dt = new D14(1);
      $dt.cf69 = cf69;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf69() { return this.cf69; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf70) + ", " + this.cf71.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf69) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf70 === other.cf70 && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf69, other.cf69);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC37(false, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC38(cf72) {
      let $dt = new D15(0);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf72 === other.cf72;
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
    static create_DC39(cf73) {
      let $dt = new D16(0);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf73, other.cf73);
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
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf75, cf76) {
      let $dt = new D17(0);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC40(cf74) {
      let $dt = new D17(1);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC41" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC40" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf75 === other.cf75 && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf74 === other.cf74;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC41(false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC43(cf78) {
      let $dt = new D18(0);
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC44(cf79, cf80, cf81, cf82) {
      let $dt = new D18(1);
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC42(cf77) {
      let $dt = new D18(2);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get is_DC42() { return this.$tag === 2; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC43" + "(" + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC44" + "(" + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC42" + "(" + _dafny.toString(this.cf77) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf79 === other.cf79 && this.cf80 === other.cf80 && _dafny.areEqual(this.cf81, other.cf81) && this.cf82 === other.cf82;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf77, other.cf77);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC43(_dafny.ZERO);
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
    static create_DC45(cf83) {
      let $dt = new D19(0);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC45" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf83 === other.cf83;
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
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC47(cf85, cf86) {
      let $dt = new D20(0);
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC48(cf87) {
      let $dt = new D20(1);
      $dt.cf87 = cf87;
      return $dt;
    }
    static create_DC46(cf84) {
      let $dt = new D20(2);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get is_DC46() { return this.$tag === 2; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC47" + "(" + _dafny.toString(this.cf85) + ", " + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC48" + "(" + _dafny.toString(this.cf87) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC46" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf85 === other.cf85 && _dafny.areEqual(this.cf86, other.cf86);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf87, other.cf87);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf84 === other.cf84;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC47(false, _dafny.ZERO);
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
    static create_DC50(cf89, cf90, cf91, cf92) {
      let $dt = new D21(0);
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      return $dt;
    }
    static create_DC49(cf88) {
      let $dt = new D21(1);
      $dt.cf88 = cf88;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf88() { return this.cf88; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC50" + "(" + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC49" + "(" + _dafny.toString(this.cf88) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf89 === other.cf89 && _dafny.areEqual(this.cf90, other.cf90) && this.cf91 === other.cf91 && this.cf92 === other.cf92;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf88, other.cf88);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC50(false, _dafny.ZERO, false, false);
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
    static create_DC52(cf94) {
      let $dt = new D22(0);
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC53(cf95, cf96, cf97, cf98) {
      let $dt = new D22(1);
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC51(cf93) {
      let $dt = new D22(2);
      $dt.cf93 = cf93;
      return $dt;
    }
    static create_DC54(cf99) {
      let $dt = new D22(3);
      $dt.cf99 = cf99;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC53() { return this.$tag === 1; }
    get is_DC51() { return this.$tag === 2; }
    get is_DC54() { return this.$tag === 3; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf99() { return this.cf99; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC52" + "(" + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC53" + "(" + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ", " + _dafny.toString(this.cf97) + ", " + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC51" + "(" + _dafny.toString(this.cf93) + ")";
      } else if (this.$tag === 3) {
        return "D22.DC54" + "(" + _dafny.toString(this.cf99) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf94, other.cf94);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf95 === other.cf95 && _dafny.areEqual(this.cf96, other.cf96) && this.cf97 === other.cf97 && this.cf98 === other.cf98;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf93, other.cf93);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf99, other.cf99);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC52(_dafny.Map.Empty);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
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
      this.f0 = false;
      this.f2 = _dafny.Seq.UnicodeFromString("");
      this.f5 = _dafny.ZERO;
      this._f1 = _dafny.Seq.UnicodeFromString("");
      this._f3 = false;
      this._f4 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      return;
    }
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
      this.f19 = [];
      this.f20 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f19, f20) {
      let _this = this;
      (_this).f19 = f19;
      (_this).f20 = f20;
      return;
    }
    fm23(p0, globalState) {
      let _this = this;
      return new BigNumber(638);
    };
    fm24(globalState) {
      let _this = this;
      return _dafny.Seq.of(!((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f20,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),_this.f20)).length),_this.f20)).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-605)), function (_235_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }))).length))).length)).isLessThan(new BigNumber(779))), _this.f20, _this.f20, _this.f20);
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      let _source4 = ((false) ? (_module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,false),false))) : (_module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(214),false),false))));
      if (_source4.is_DC10) {
        let _236___mcc_h0 = (_source4).cf22;
        let _237_cf22 = _236___mcc_h0;
        return _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0))));
      } else if (_source4.is_DC9) {
        let _238___mcc_h1 = (_source4).cf21;
        let _239_cf21 = _238___mcc_h1;
        return (_module.D4.create_DC12(_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)))))).dtor_cf24;
      } else {
        let _240___mcc_h2 = (_source4).cf23;
        let _241_cf23 = _240___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))));
      }
    };
    fm6(p0, globalState) {
      let _this = this;
      return function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of ((_dafny.MultiSet.fromElements((_this).f6, (_this).f6, new BigNumber((_dafny.Seq.UnicodeFromString("pxn")).length), (_this).f6)).Union(_dafny.MultiSet.fromElements((_this).f6))).Elements) {
          let _242_v0 = _compr_8;
          if (((_dafny.MultiSet.fromElements((_this).f6, (_this).f6, new BigNumber((_dafny.Seq.UnicodeFromString("pxn")).length), (_this).f6)).Union(_dafny.MultiSet.fromElements((_this).f6))).contains(_242_v0)) {
            _coll8.push([(_242_v0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(272),(_this).f6)).length)),_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(256)), function (_243_i0) {
              return (_this).f6;
            }), _dafny.Seq.of((_this).f6))]);
          }
        }
        return _coll8;
      }();
    };
    fm28(globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f6), _dafny.Seq.of((_this).f6, (_this).f6))).length))).isLessThanOrEqualTo(new BigNumber(618));
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f16 = _dafny.MultiSet.Empty;
      this._f9 = [];
      this._f6 = _dafny.ZERO;
      this._f10 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0, _module.T2, _module.T1];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    set f9(value) {
      let _this = this;
      _this._f9 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f6, f16, f9, f10) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f16 = f16;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      let _source5 = _module.D2.create_DC8(!(!(false)), false);
      if (_source5.is_DC6) {
        let _244___mcc_h0 = (_source5).cf13;
        let _245___mcc_h1 = (_source5).cf14;
        let _246_cf14 = _245___mcc_h1;
        let _247_cf13 = _244___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))));
      } else if (_source5.is_DC7) {
        let _248___mcc_h2 = (_source5).cf15;
        let _249___mcc_h3 = (_source5).cf16;
        let _250___mcc_h4 = (_source5).cf17;
        let _251___mcc_h5 = (_source5).cf18;
        let _252_cf18 = _251___mcc_h5;
        let _253_cf17 = _250___mcc_h4;
        let _254_cf16 = _249___mcc_h3;
        let _255_cf15 = _248___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))));
      } else if (_source5.is_DC8) {
        let _256___mcc_h6 = (_source5).cf19;
        let _257___mcc_h7 = (_source5).cf20;
        let _258_cf20 = _257___mcc_h7;
        let _259_cf19 = _256___mcc_h6;
        return _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0))));
      } else {
        let _260___mcc_h8 = (_source5).cf12;
        let _261_cf12 = _260___mcc_h8;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(523)), function (_262_i0) {
          return _dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)));
        }));
      }
    };
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.of((_this).f6));
    };
    fm21(p0, globalState) {
      let _this = this;
      return ((_this).f6).minus((_this).f6);
    };
    fm22(p0, p1, p2, globalState) {
      let _this = this;
      if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))), new _dafny.CodePoint('d'.codePointAt(0)))) {
        return true;
      } else {
        return ((_dafny.ZERO).minus((_this).f6)).isLessThanOrEqualTo(new BigNumber(248));
      }
    };
    fm13(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(!(!(true)))))),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-918)), function (_263_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.Create(_module.__default.abs(new BigNumber(2)), function (_264_i1) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("rxhrrtyr"))));
    };
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _265_v0;
      _265_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,!(p0));
      let _266_v1;
      _266_v1 = _dafny.Set.fromElements((_265_v0).Merge(_265_v0));
      let _267_v2;
      _267_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(91),(_this).f6);
      let _268_v3;
      _268_v3 = _module.D2.create_DC5(_267_v2);
      let _269_v4;
      _269_v4 = _dafny.Set.fromElements(p0);
      let _270_v5;
      _270_v5 = _dafny.MultiSet.fromElements(p0);
      let _271_v6;
      _271_v6 = _dafny.MultiSet.fromElements((_this).f6, (_this).f6);
      let _272_v7;
      _272_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_271_v6).update((_this).f6, _module.__default.abs(p2)));
      let _rhs26 = ((_this).f6).isLessThanOrEqualTo((new BigNumber((_269_v4).length)).minus(new BigNumber(648)));
      let _rhs27 = ((((_this).fm22((_270_v5).update(p0, _module.__default.abs((_this).f6)), p0, (_this).f6, globalState)) ? (_266_v1) : (_266_v1))).Difference(((p0) ? (_266_v1) : (_266_v1)));
      let _rhs28 = _module.D2.create_DC5((_267_v2).Merge(_267_v2));
      let _rhs29 = _module.__default.safeModuloInt((_this).f6, new BigNumber(((((_272_v7).contains(p0)) ? ((_272_v7).get(p0)) : (_271_v6))).cardinality()));
      let _rhs30 = p0;
      let _lhs19 = globalState;
      let _lhs20 = globalState;
      _lhs19.f0 = _rhs26;
      _266_v1 = _rhs27;
      _268_v3 = _rhs28;
      _lhs20.f5 = _rhs29;
      r1 = _rhs30;
      let _273_v8;
      let _nw34 = Array((new BigNumber(3)).toNumber());
      _nw34[(_dafny.ZERO).toNumber()] = (_this).f6;
      _nw34[(_dafny.ONE).toNumber()] = (_this).f6;
      _nw34[(new BigNumber(2)).toNumber()] = (_this).f6;
      _273_v8 = _nw34;
      let _274_v9;
      _274_v9 = _dafny.Set.fromElements((_this).f6);
      let _275_v10;
      _275_v10 = _dafny.Seq.of(_274_v9, _274_v9, _274_v9, _274_v9);
      let _index27 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_273_v8).length));
      (_273_v8)[_index27] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_275_v10, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(p2))).cardinality()), new BigNumber((_275_v10).length)), _274_v9), _275_v10)).length);
      let _index28 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_273_v8).length));
      let _rhs31 = (_module.__default.safeModuloInt((_this).f6, (_this).f6)).isLessThanOrEqualTo((((_this).fm22(_270_v5, p0, (_this).f6, globalState)) ? (p2) : ((_this).f6)));
      let _rhs32 = true;
      let _rhs33 = (((_271_v6).IsDisjointFrom(_271_v6)) ? (new BigNumber((_269_v4).length)) : (((_this).f6).plus((_this).f6)));
      let _rhs34 = p0;
      let _lhs21 = globalState;
      let _lhs22 = globalState;
      let _lhs23 = _273_v8;
      let _lhs24 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_273_v8).length));
      let _lhs25 = globalState;
      _lhs21.f0 = _rhs31;
      _lhs22.f0 = _rhs32;
      _lhs23[_lhs24] = _rhs33;
      _lhs25.f0 = _rhs34;
      let _276_v11;
      let _nw35 = new _module.C1();
      _nw35.__ctor(new BigNumber(-380));
      _276_v11 = _nw35;
      (globalState).f0 = !(((_273_v8)[_module.__default.safeIndex(new BigNumber(48), new BigNumber((_273_v8).length))]).isLessThanOrEqualTo(p2)) || (p0);
      let _277_v12;
      let _nw36 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
      _277_v12 = _nw36;
      let _278_v13;
      _278_v13 = _module.D2.create_DC7(p1, _277_v12, (_this).f6, new BigNumber((_269_v4).length));
      let _279_v14;
      _279_v14 = _dafny.Map.Empty.slice().updateUnsafe(_278_v13,_dafny.Seq.UnicodeFromString("gtvimpmrf"));
      let _280_v15;
      _280_v15 = _dafny.Seq.UnicodeFromString("ohcxxhxh");
      let _281_v16;
      _281_v16 = new _dafny.CodePoint('t'.codePointAt(0));
      (globalState).f5 = ((_273_v8)[_module.__default.safeIndex(new BigNumber(48), new BigNumber((_273_v8).length))]).multipliedBy((new BigNumber((_dafny.Seq.update((((_279_v14).contains(_278_v13)) ? ((_279_v14).get(_278_v13)) : (_280_v15)), _module.__default.safeIndex((_this).f6, new BigNumber(((((_279_v14).contains(_278_v13)) ? ((_279_v14).get(_278_v13)) : (_280_v15))).length)), _281_v16)).length)).plus(new BigNumber((_280_v15).length)));
      let _arr2 = _this.f9;
      let _index29 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_this.f9).length));
      _arr2[_index29] = p0;
      let _arr3 = _this.f9;
      let _index30 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_this.f9).length));
      _arr3[_index30] = !(p0);
      r0 = _dafny.Seq.update(_module.__default.fm29(_274_v9, globalState), _module.__default.safeIndex(((p0) ? ((_this).f6) : ((_273_v8)[_module.__default.safeIndex(new BigNumber(48), new BigNumber((_273_v8).length))])), new BigNumber((_module.__default.fm29(_274_v9, globalState)).length)), _dafny.Seq.Concat(_280_v15, _280_v15));
      let _282_v17;
      _282_v17 = _dafny.Seq.of(new BigNumber((_267_v2).length));
      r1 = ((new BigNumber((_282_v17).length)).multipliedBy((_273_v8)[_module.__default.safeIndex(new BigNumber(48), new BigNumber((_273_v8).length))])).isLessThanOrEqualTo(p2);
      return [r0, r1];
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = [];
      let r3 = _dafny.ZERO;
      let _283_v0;
      _283_v0 = true;
      if (!(_283_v0) || (!(!(_283_v0)))) {
        let _284_v1;
        _284_v1 = _dafny.Set.fromElements(_283_v0);
        let _285_v2;
        _285_v2 = _module.D5.create_DC16(_284_v1);
        let _286_v3;
        _286_v3 = _dafny.Seq.of((_285_v2).dtor_cf35, (_284_v1).Difference(_284_v1), (_284_v1).Difference(_284_v1));
        let _287_v4;
        let _nw37 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Set.Empty);
        _287_v4 = _nw37;
        let _index31 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_287_v4).length));
        (_287_v4)[_index31] = (_dafny.Set.fromElements(_283_v0)).Difference((_285_v2).dtor_cf35);
        let _288_v5;
        _288_v5 = new _dafny.CodePoint('v'.codePointAt(0));
        let _289_v6;
        _289_v6 = _dafny.Map.Empty.slice().updateUnsafe(!(_283_v0),(_this).f6);
        let _index32 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_287_v4).length));
        let _rhs35 = _288_v5;
        let _rhs36 = ((_283_v0) ? (_286_v3) : (_286_v3));
        let _rhs37 = (_this).fm21(new BigNumber(((_289_v6).Merge(_289_v6)).length), globalState);
        let _rhs38 = p1;
        let _rhs39 = ((_284_v1).Intersect(_284_v1)).Intersect(_dafny.Set.fromElements(_283_v0, true, _283_v0));
        let _lhs26 = globalState;
        let _lhs27 = _287_v4;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_287_v4).length));
        r0 = _rhs35;
        _286_v3 = _rhs36;
        _lhs26.f5 = _rhs37;
        r3 = _rhs38;
        _lhs27[_lhs28] = _rhs39;
        let _290_v7;
        _290_v7 = _dafny.Map.Empty.slice().updateUnsafe(_283_v0,!(_283_v0));
        let _291_v8;
        _291_v8 = _module.D4.create_DC14(_283_v0, _290_v7, (_this).f6, _dafny.Map.Empty.slice().updateUnsafe(_283_v0,_283_v0));
        let _292_v9;
        _292_v9 = _dafny.Seq.of(false, _283_v0);
        let _293_v10;
        _293_v10 = _dafny.Map.Empty.slice().updateUnsafe(_291_v8,_292_v9);
        _293_v10 = (_293_v10).update(_291_v8, _dafny.Seq.update(_dafny.Seq.of(true), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.of(true)).length)), !((_292_v9)[_module.__default.safeIndex(p1, new BigNumber((_292_v9).length))])));
        r3 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_283_v0,_283_v0)).Merge((_290_v7).Merge(_290_v7))).length);
        let _294_v11;
        _294_v11 = _dafny.Seq.UnicodeFromString("cidjjdvpa");
        (globalState).f2 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ik"), _dafny.Seq.Concat(_294_v11, _294_v11)), _module.__default.safeIndex(_module.__default.safeModuloInt((_this).f6, p1), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ik"), _dafny.Seq.Concat(_294_v11, _294_v11))).length)), _288_v5);
        (globalState).f0 = _module.__default.fm3(globalState);
      } else {
        (globalState).f0 = _283_v0;
        if (!(_283_v0)) {
          (globalState).f0 = _283_v0;
          let _295_v12;
          let _nw38 = Array((new BigNumber(5)).toNumber()).fill([]);
          _295_v12 = _nw38;
          let _296_v13;
          let _nw39 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
          _296_v13 = _nw39;
          let _index33 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_295_v12).length));
          (_295_v12)[_index33] = _296_v13;
          let _index34 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_295_v12).length));
          (_295_v12)[_index34] = _296_v13;
          let _297_v14;
          _297_v14 = _dafny.Seq.UnicodeFromString("rdc");
          _283_v0 = _dafny.areEqual(_297_v14, _297_v14);
          let _298_v15;
          let _nw40 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _298_v15 = _nw40;
          let _index35 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_298_v15).length));
          (_298_v15)[_index35] = _297_v14;
          let _299_v16;
          let _nw41 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _299_v16 = _nw41;
          let _300_v17;
          _300_v17 = _dafny.Seq.of(p1);
          let _index36 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_299_v16).length));
          (_299_v16)[_index36] = (_dafny.ZERO).minus(((_283_v0) ? (p1) : (new BigNumber((_300_v17).length))));
          let _301_v18;
          _301_v18 = _dafny.Seq.of(true);
          let _302_v19;
          _302_v19 = _module.D4.create_DC13(_dafny.Seq.of(new BigNumber(-406)), _283_v0, p1);
          let _303_v20;
          _303_v20 = _dafny.Map.Empty.slice().updateUnsafe(_302_v19,new BigNumber((_300_v17).length));
          let _304_v21;
          _304_v21 = _dafny.Map.Empty.slice().updateUnsafe((((_303_v20).contains(_302_v19)) ? ((_303_v20).get(_302_v19)) : (p1)),(_this).f6);
          let _305_v22;
          _305_v22 = new _dafny.CodePoint('c'.codePointAt(0));
          let _index37 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_298_v15).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_299_v16).length));
          let _rhs40 = new BigNumber((_300_v17).length);
          let _rhs41 = _module.__default.safeModuloInt((_this).f6, _module.__default.safeDivisionInt(p1, new BigNumber((_301_v18).length)));
          let _rhs42 = _dafny.Seq.update(_297_v14, _module.__default.safeIndex((((_304_v21).contains((_dafny.ZERO).minus(new BigNumber((_304_v21).length)))) ? ((_304_v21).get((_dafny.ZERO).minus(new BigNumber((_304_v21).length)))) : (p1)), new BigNumber((_297_v14).length)), _305_v22);
          let _rhs43 = (p1).multipliedBy(p1);
          let _rhs44 = p1;
          let _lhs29 = globalState;
          let _lhs30 = globalState;
          let _lhs31 = _298_v15;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_298_v15).length));
          let _lhs33 = _299_v16;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_299_v16).length));
          let _lhs35 = globalState;
          _lhs29.f5 = _rhs40;
          _lhs30.f5 = _rhs41;
          _lhs31[_lhs32] = _rhs42;
          _lhs33[_lhs34] = _rhs43;
          _lhs35.f5 = _rhs44;
          let _index39 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_299_v16).length));
          (_299_v16)[_index39] = p1;
        } else {
          let _306_v23;
          _306_v23 = new _dafny.CodePoint('j'.codePointAt(0));
          let _307_v24;
          _307_v24 = _dafny.Seq.UnicodeFromString("tgruuabrh");
          let _308_v25;
          _308_v25 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yjxblq"), _dafny.Seq.UnicodeFromString("kfxtlw")), ((_283_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), function (_309_i0) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })) : (_dafny.Seq.UnicodeFromString("yiaoosnib"))), _module.__default.fm30(true, _306_v23, globalState), _307_v24);
          (globalState).f2 = (_308_v25)[_module.__default.safeIndex((_this).f6, new BigNumber((_308_v25).length))];
          (globalState).f0 = !(!(_283_v0));
          _283_v0 = !((_this).f6).isEqualTo((_this).f6);
          (globalState).f2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_307_v24, _307_v24), _307_v24);
          (globalState).f0 = _283_v0;
        }
        let _310_v26;
        _310_v26 = _dafny.Map.Empty.slice().updateUnsafe(_283_v0,(_this).f6);
        let _311_v27;
        _311_v27 = _dafny.Seq.of((_310_v26).Merge(_310_v26));
        let _rhs45 = (_310_v26).Merge(_module.__default.fm31(_283_v0, (_this).f6, _module.__default.fm2(p1, globalState), _dafny.Seq.of(_283_v0), globalState));
        let _rhs46 = _dafny.Seq.of(_310_v26, (_310_v26).Merge(_310_v26), _310_v26, _310_v26);
        _310_v26 = _rhs45;
        _311_v27 = _rhs46;
        let _312_v28;
        _312_v28 = _dafny.Seq.of(_283_v0);
        let _313_v29;
        _313_v29 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_312_v28, _312_v28),_283_v0);
        _313_v29 = _313_v29;
        let _314_v30;
        _314_v30 = _dafny.Map.Empty.slice().updateUnsafe(_283_v0,_283_v0);
        let _315_v31;
        _315_v31 = new _dafny.CodePoint('i'.codePointAt(0));
        let _316_v32;
        _316_v32 = _dafny.MultiSet.fromElements(_315_v31, _315_v31, _315_v31, _315_v31, new _dafny.CodePoint('m'.codePointAt(0)));
        let _317_v33;
        _317_v33 = _dafny.MultiSet.fromElements(!(_283_v0), !(_283_v0));
        let _318_v34;
        _318_v34 = _module.D4.create_DC14(true, _dafny.Map.Empty.slice().updateUnsafe(_283_v0,true), new BigNumber(-428), _314_v30);
        let _319_v35;
        _319_v35 = _dafny.Set.fromElements(true, !(_283_v0), (_318_v34).dtor_cf28, _283_v0, true);
        let _320_v36;
        let _nw42 = Array((new BigNumber(14)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = new BigNumber((_314_v30).length);
        _nw42[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_315_v31,_283_v0)).length);
        _nw42[(new BigNumber(2)).toNumber()] = ((_this).f6).minus(p1);
        _nw42[(new BigNumber(3)).toNumber()] = (((_316_v32).contains(_315_v31)) ? ((_316_v32).get(_315_v31)) : (p1));
        _nw42[(new BigNumber(4)).toNumber()] = new BigNumber((_312_v28).length);
        _nw42[(new BigNumber(5)).toNumber()] = _module.__default.fm2((_this).f6, globalState);
        _nw42[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw42[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_283_v0)).length);
        _nw42[(new BigNumber(8)).toNumber()] = (_module.__default.fm2(p1, globalState)).minus(new BigNumber((_317_v33).cardinality()));
        _nw42[(new BigNumber(9)).toNumber()] = p1;
        _nw42[(new BigNumber(10)).toNumber()] = (_this).f6;
        _nw42[(new BigNumber(11)).toNumber()] = (_this).fm21(p1, globalState);
        _nw42[(new BigNumber(12)).toNumber()] = p1;
        _nw42[(new BigNumber(13)).toNumber()] = ((_283_v0) ? (new BigNumber((_319_v35).length)) : ((_this).f6));
        _320_v36 = _nw42;
        let _index40 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_320_v36).length));
        (_320_v36)[_index40] = new BigNumber(-446);
        let _321_v37;
        _321_v37 = _module.D0.create_DC0(_312_v28, new BigNumber(593));
        let _322_v38;
        _322_v38 = _module.D1.create_DC4(_321_v37, (_this).fm21(p1, globalState), _dafny.Seq.UnicodeFromString("rdeikwgw"), _283_v0, p1);
        let _323_v39;
        _323_v39 = _dafny.Seq.UnicodeFromString("s");
        let _324_v40;
        _324_v40 = _dafny.Map.Empty.slice().updateUnsafe(_283_v0,(_323_v39)[_module.__default.safeIndex(new BigNumber((_323_v39).length), new BigNumber((_323_v39).length))]);
        let _pat_let_tv3 = _324_v40;
        let _pat_let_tv4 = _323_v39;
        let _325_v42;
        _325_v42 = _module.D1.create_DC4(_321_v37, (_this).f6, (function (_pat_let6_0) {
  return function (_326_dt__update__tmp_h0) {
    return function (_pat_let7_0) {
      return function (_327_dt__update_hcf8_h0) {
        return function (_pat_let8_0) {
          return function (_328_dt__update_hcf9_h0) {
            return _module.D1.create_DC4((_326_dt__update__tmp_h0).dtor_cf7, _327_dt__update_hcf8_h0, _328_dt__update_hcf9_h0, (_326_dt__update__tmp_h0).dtor_cf10, (_326_dt__update__tmp_h0).dtor_cf11);
          }(_pat_let8_0);
        }(_pat_let_tv4);
      }(_pat_let7_0);
    }(new BigNumber((_pat_let_tv3).length));
  }(_pat_let6_0);
}(_322_v38)).dtor_cf9, _283_v0, new BigNumber((function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of _dafny.IntegerRange(new BigNumber(401), new BigNumber(516))) {
    let _329_v41 = _compr_9;
    if (((new BigNumber(401)).isLessThanOrEqualTo(_329_v41)) && ((_329_v41).isLessThan(new BigNumber(516)))) {
      _coll9.push([(_329_v41).plus(p1),p1]);
    }
  }
  return _coll9;
}()).length));
        let _330_v43;
        _330_v43 = _dafny.Seq.of(_325_v42, _325_v42);
        let _index41 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_320_v36).length));
        (_320_v36)[_index41] = p1;
        let _331_v44;
        _331_v44 = _dafny.Set.fromElements(_320_v36);
        let _332_v45;
        _332_v45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_283_v0),_283_v0);
        let _333_v46;
        _333_v46 = _module.D3.create_DC10(_332_v45);
        let _334_v47;
        _334_v47 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(_333_v46, _333_v46, _module.D3.create_DC10(_332_v45)));
        let _index42 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_320_v36).length));
        let _index43 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_320_v36).length));
        let _rhs47 = (_this).f6;
        let _rhs48 = _dafny.Seq.of(_module.D1.create_DC4(_module.D0.create_DC0(_312_v28, new BigNumber(-852)), (_this).f6, _323_v39, _283_v0, (_this).f6));
        let _rhs49 = new BigNumber((((_334_v47).Difference(_334_v47)).Union(_334_v47)).length);
        let _rhs50 = _dafny.Set.fromElements(_320_v36, _320_v36, _320_v36, _320_v36);
        let _lhs36 = _320_v36;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_320_v36).length));
        let _lhs38 = _320_v36;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_320_v36).length));
        _lhs36[_lhs37] = _rhs47;
        _330_v43 = _rhs48;
        _lhs38[_lhs39] = _rhs49;
        _331_v44 = _rhs50;
      }
      let _335_v48;
      let _init3 = function (_336_i1) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(651)), function (_337_i2) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        });
      };
      let _nw43 = Array((new BigNumber(11)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw43.length); _i0_3++) {
        _nw43[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _335_v48 = _nw43;
      let _338_v49;
      _338_v49 = _dafny.Seq.UnicodeFromString("aabonynb");
      let _index44 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_335_v48).length));
      (_335_v48)[_index44] = _dafny.Seq.Concat(_338_v49, _338_v49);
      let _339_v50;
      _339_v50 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f6, (_this).f6),new BigNumber((_338_v49).length));
      let _340_v51;
      _340_v51 = _dafny.Map.Empty.slice().updateUnsafe(_283_v0,true);
      let _index45 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_335_v48).length));
      let _rhs51 = (((_339_v50).contains((new BigNumber(731)).minus((_this).f6))) ? ((_339_v50).get((new BigNumber(731)).minus((_this).f6))) : ((_this).f6));
      let _rhs52 = _283_v0;
      let _rhs53 = (p1).isLessThanOrEqualTo(new BigNumber((_340_v51).length));
      let _rhs54 = _dafny.Seq.Concat(_338_v49, _338_v49);
      let _lhs40 = _335_v48;
      let _lhs41 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_335_v48).length));
      r1 = _rhs51;
      _283_v0 = _rhs52;
      _283_v0 = _rhs53;
      _lhs40[_lhs41] = _rhs54;
      let _index46 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_335_v48).length));
      (_335_v48)[_index46] = _338_v49;
      let _hi1 = _module.__default.safeModuloInt((_this).f6, p1);
      for (let _341_i3 = p1; _341_i3.isLessThan(_hi1); _341_i3 = _341_i3.plus(_dafny.ONE)) {
        let _342_v52;
        let _nw44 = new _module.C0();
        _nw44.__ctor(_this.f9, ((_283_v0) ? (_283_v0) : (_283_v0)));
        _342_v52 = _nw44;
        let _343_v53;
        _343_v53 = new _dafny.CodePoint('t'.codePointAt(0));
        let _344_v54;
        _344_v54 = _dafny.Map.Empty.slice().updateUnsafe(_343_v53,_342_v52.f20);
        (_342_v52).f20 = !_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_343_v53,_283_v0)), (_344_v54).Merge(_344_v54));
        r3 = new BigNumber(15);
        let _345_v55;
        _345_v55 = _dafny.MultiSet.fromElements(_342_v52.f20);
        let _346_v56;
        _346_v56 = _module.D1.create_DC4(function (_pat_let9_0) {
  return function (_347_dt__update__tmp_h1) {
    return function (_pat_let10_0) {
      return function (_348_dt__update_hcf1_h0) {
        return _module.D0.create_DC0((_347_dt__update__tmp_h1).dtor_cf0, _348_dt__update_hcf1_h0);
      }(_pat_let10_0);
    }(new BigNumber(811));
  }(_pat_let9_0);
}(_module.__default.fm32(_283_v0, _342_v52.f20, _341_i3, _283_v0, globalState)), (((_345_v55).contains(_342_v52.f20)) ? ((_345_v55).get(_342_v52.f20)) : ((_this).f6)), _338_v49, _283_v0, (_dafny.ZERO).minus(p1));
        let _349_v57;
        _349_v57 = _dafny.Set.fromElements(!(_283_v0));
        let _350_v58;
        _350_v58 = _module.D5.create_DC16(_349_v57);
        let _pat_let_tv5 = _350_v58;
        let _351_v59;
        _351_v59 = _dafny.Map.Empty.slice().updateUnsafe(_342_v52.f20,function (_pat_let11_0) {
          return function (_352_dt__update__tmp_h2) {
            return function (_pat_let12_0) {
              return function (_353_dt__update_hcf8_h1) {
                return _module.D1.create_DC4((_352_dt__update__tmp_h2).dtor_cf7, _353_dt__update_hcf8_h1, (_352_dt__update__tmp_h2).dtor_cf9, (_352_dt__update__tmp_h2).dtor_cf10, (_352_dt__update__tmp_h2).dtor_cf11);
              }(_pat_let12_0);
            }(new BigNumber(((_pat_let_tv5).dtor_cf35).length));
          }(_pat_let11_0);
        }(_346_v56));
        _351_v59 = (_351_v59).update(!(_283_v0), _346_v56);
      }
      let _nw45 = Array((new BigNumber(3)).toNumber()).fill(false);
      (_this).f9 = _nw45;
      let _354_v60;
      _354_v60 = _dafny.Seq.of((_this).f6);
      let _355_v61;
      _355_v61 = _module.D3.create_DC9(_dafny.Seq.of((_this).f6, (_this).f6));
      let _356_v62;
      _356_v62 = _dafny.Seq.of(_this.f9, ((_283_v0) ? (_this.f9) : (_this.f9)), _this.f9, _this.f9);
      let _357_v64;
      _357_v64 = _module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of _dafny.IntegerRange(new BigNumber(620), new BigNumber(163))) {
    let _358_v63 = _compr_10;
    if (((new BigNumber(620)).isLessThanOrEqualTo(_358_v63)) && ((_358_v63).isLessThan(new BigNumber(163)))) {
      _coll10.push([(_358_v63).multipliedBy(new BigNumber((_354_v60).length)),_283_v0]);
    }
  }
  return _coll10;
}(),_283_v0));
      let _pat_let_tv6 = p1;
      let _pat_let_tv7 = p1;
      let _pat_let_tv8 = p1;
      let _pat_let_tv9 = _283_v0;
      let _pat_let_tv10 = _283_v0;
      let _rhs55 = (new BigNumber((_dafny.Seq.Concat(_338_v49, _338_v49)).length)).minus(_module.__default.safeDivisionInt((_this).f6, new BigNumber((_354_v60).length)));
      let _rhs56 = function (_source6) {
        if (_source6.is_DC10) {
          let _359___mcc_h0 = (_source6).cf22;
          let _360_cf22 = _359___mcc_h0;
          return !(!(!((_dafny.ZERO).minus(_pat_let_tv6)).isEqualTo(_pat_let_tv7)));
        } else if (_source6.is_DC9) {
          let _361___mcc_h1 = (_source6).cf21;
          let _362_cf21 = _361___mcc_h1;
          return !((_dafny.ZERO).minus(_pat_let_tv8)).isEqualTo((_this).f6);
        } else {
          let _363___mcc_h2 = (_source6).cf23;
          let _364_cf23 = _363___mcc_h2;
          return !(_pat_let_tv9) || (_pat_let_tv10);
        }
      }(_module.D3.create_DC11(_module.D3.create_DC11(_module.D3.create_DC11(_module.D3.create_DC11(_355_v61)))));
      let _rhs57 = (_356_v62)[_module.__default.safeIndex((_dafny.ZERO).minus(((_this).f6).multipliedBy((_this).f6)), new BigNumber((_356_v62).length))];
      let _rhs58 = !(_283_v0);
      let _rhs59 = _module.__default.safeModuloInt((_this).f6, (new BigNumber((_module.__default.fm33(_357_v64, new BigNumber(969), _283_v0, _283_v0, globalState)).length)).plus((_this).f6));
      let _lhs42 = _this;
      let _lhs43 = globalState;
      let _lhs44 = globalState;
      r3 = _rhs55;
      _283_v0 = _rhs56;
      _lhs42.f9 = _rhs57;
      _lhs43.f0 = _rhs58;
      _lhs44.f5 = _rhs59;
      r0 = new _dafny.CodePoint('j'.codePointAt(0));
      r1 = _module.__default.fm2(p1, globalState);
      r2 = _this.f9;
      r3 = (((_this).f6).plus(new BigNumber((_338_v49).length))).minus((_this).f6);
      return [r0, r1, r2, r3];
    }
    m15(p0, p1, p2, globalState) {
      let _this = this;
      (globalState).f0 = !(p0);
      let _hi2 = (new BigNumber(442)).minus(new BigNumber(553));
      for (let _365_i0 = new BigNumber(832); _365_i0.isLessThan(_hi2); _365_i0 = _365_i0.plus(_dafny.ONE)) {
        let _366_v0;
        let _nw46 = new _module.C1();
        _nw46.__ctor((_this).f6);
        _366_v0 = _nw46;
        let _367_v1;
        let _nw47 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _367_v1 = _nw47;
        let _index47 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length));
        (_367_v1)[_index47] = (_365_i0).plus(_365_i0);
        let _index48 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_367_v1).length));
        (_367_v1)[_index48] = _365_i0;
        let _368_v2;
        let _init4 = function (_369_i1) {
          return (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of((_this).f6)).length))).Difference(_dafny.Set.fromElements((_this).f6));
        };
        let _nw48 = Array((new BigNumber(26)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw48.length); _i0_4++) {
          _nw48[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _368_v2 = _nw48;
        let _370_v3;
        _370_v3 = _dafny.Set.fromElements(_365_i0);
        let _index49 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_368_v2).length));
        (_368_v2)[_index49] = (_370_v3).Intersect(_370_v3);
        let _371_v4;
        _371_v4 = _dafny.Seq.of(p1);
        let _index50 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length));
        let _index51 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_367_v1).length));
        let _index52 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_368_v2).length));
        let _rhs60 = (_this).f6;
        let _rhs61 = ((_dafny.areEqual(_371_v4, _371_v4)) ? (_module.__default.safeDivisionInt(new BigNumber(682), new BigNumber(-15))) : (((_this).f6).minus((_this).f6)));
        let _rhs62 = _370_v3;
        let _lhs45 = _367_v1;
        let _lhs46 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length));
        let _lhs47 = _367_v1;
        let _lhs48 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_367_v1).length));
        let _lhs49 = _368_v2;
        let _lhs50 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_368_v2).length));
        _lhs45[_lhs46] = _rhs60;
        _lhs47[_lhs48] = _rhs61;
        _lhs49[_lhs50] = _rhs62;
        let _372_v5;
        _372_v5 = _dafny.Seq.UnicodeFromString("fwds");
        if (!(true) || ((_371_v4)[_module.__default.safeIndex(new BigNumber((_372_v5).length), new BigNumber((_371_v4).length))])) {
          let _373_v6;
          let _init5 = ((_374_v5) => function (_375_i2) {
            return _374_v5;
          })(_372_v5);
          let _nw49 = Array((new BigNumber(23)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw49.length); _i0_5++) {
            _nw49[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _373_v6 = _nw49;
          let _index53 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_373_v6).length));
          (_373_v6)[_index53] = _dafny.Seq.Concat(_372_v5, _372_v5);
          let _index54 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_373_v6).length));
          (_373_v6)[_index54] = _372_v5;
          let _376_v7;
          _376_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
          let _377_v8;
          _377_v8 = _dafny.Seq.of(new BigNumber((_376_v7).length), _365_i0);
          let _378_v9;
          _378_v9 = _module.D4.create_DC13(_377_v8, p1, (_this).f6);
          let _379_v10;
          _379_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm21(_365_i0, globalState),(_378_v9).dtor_cf26);
          let _380_v11;
          _380_v11 = _dafny.Set.fromElements(p0, false, p1);
          let _381_v12;
          _381_v12 = _dafny.Set.fromElements(_380_v11, _380_v11);
          _379_v10 = (_379_v10).update((_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))], !((_381_v12).contains(_380_v11)));
          (globalState).f5 = _365_i0;
          let _index55 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length));
          (_367_v1)[_index55] = (_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))];
          let _382_v13;
          _382_v13 = _dafny.Map.Empty.slice().updateUnsafe((_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))],(_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))]);
          _382_v13 = (_382_v13).update(_module.__default.safeModuloInt((_this).f6, new BigNumber(-29)), (_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))]);
        } else {
          let _383_v14;
          _383_v14 = new _dafny.CodePoint('p'.codePointAt(0));
          let _384_v15;
          _384_v15 = _dafny.Map.Empty.slice().updateUnsafe((_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))],p1);
          let _385_v16;
          _385_v16 = _dafny.Seq.of((_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))], new BigNumber((_384_v15).length));
          let _386_v17;
          _386_v17 = _dafny.Map.Empty.slice().updateUnsafe((_module.D4.create_DC13(_385_v16, p1, _365_i0)).dtor_cf27,p0);
          let _387_v18;
          _387_v18 = _dafny.Map.Empty.slice().updateUnsafe(_383_v14,new BigNumber(((_386_v17).update((_this).f6, p0)).length));
          _387_v18 = (_387_v18).update(_383_v14, new BigNumber(43));
          (globalState).f0 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(32)), function (_388_i3) {
            return (_this).f6;
          }), (_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))]);
          (globalState).f5 = (_dafny.ZERO).minus((new BigNumber((_371_v4).length)).multipliedBy((_dafny.ZERO).minus((_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))])));
          let _index56 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length));
          (_367_v1)[_index56] = (_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))];
          (globalState).f0 = p0;
        }
        let _389_v19;
        _389_v19 = _dafny.Map.Empty.slice().updateUnsafe(_365_i0,p0);
        let _390_v20;
        _390_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,((!((((_389_v19).contains(new BigNumber(947))) ? ((_389_v19).get(new BigNumber(947))) : (p1)))) ? ((_367_v1)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_367_v1).length))]) : (new BigNumber(229))));
        _390_v20 = _390_v20;
      }
      let _391_v21;
      _391_v21 = _dafny.Set.fromElements(_this.f9);
      let _392_v22;
      _392_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_391_v21).length));
      let _393_v23;
      _393_v23 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length));
      let _hi3 = (new BigNumber(((_392_v22).update(p0, (_this).f6)).length)).minus((_393_v23)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f6), new BigNumber((_393_v23).length))]);
      for (let _394_i4 = (_this).f6; _394_i4.isLessThan(_hi3); _394_i4 = _394_i4.plus(_dafny.ONE)) {
        let _395_v24;
        let _nw50 = new _module.C0();
        _nw50.__ctor(_this.f9, p1);
        _395_v24 = _nw50;
        let _396_v25;
        let _nw51 = new _module.C1();
        _nw51.__ctor((_this).f6);
        _396_v25 = _nw51;
        let _397_v26;
        _397_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_395_v24.f19);
        let _398_v27;
        _398_v27 = new _dafny.CodePoint('c'.codePointAt(0));
        let _399_v28;
        _399_v28 = _dafny.Map.Empty.slice().updateUnsafe(_398_v27,(_this).f6);
        let _400_v29;
        _400_v29 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_394_i4)).length)));
        (_395_v24).f19 = (((_397_v26).contains(_module.__default.safeModuloInt(new BigNumber((_399_v28).length), new BigNumber((_400_v29).length)))) ? ((_397_v26).get(_module.__default.safeModuloInt(new BigNumber((_399_v28).length), new BigNumber((_400_v29).length)))) : (_this.f9));
        let _401_v30;
        let _nw52 = new _module.C1();
        _nw52.__ctor((_this).f6);
        _401_v30 = _nw52;
      }
      let _402_v31;
      _402_v31 = _dafny.MultiSet.fromElements(_393_v23);
      let _403_v32;
      _403_v32 = _dafny.Seq.UnicodeFromString("waawxeny");
      let _404_v33;
      _404_v33 = _dafny.Map.Empty.slice().updateUnsafe(_403_v32,_dafny.Seq.of(_module.__default.fm2(new BigNumber((_403_v32).length), globalState)));
      let _405_v34;
      _405_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      let _406_v35;
      _406_v35 = _module.D4.create_DC15(p1, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(53)), function (_407_i6) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}));
      let _408_v36;
      _408_v36 = _dafny.MultiSet.fromElements(_module.D4.create_DC15(p1, p0, _403_v32), _406_v35);
      let _409_v37;
      _409_v37 = _dafny.Seq.of(_dafny.Seq.of((_this).f6));
      let _410_v38;
      let _nw53 = Array((new BigNumber(11)).toNumber());
      _nw53[(_dafny.ZERO).toNumber()] = (((_402_v31).contains((((_404_v33).contains(_403_v32)) ? ((_404_v33).get(_403_v32)) : (_393_v23)))) ? ((_402_v31).get((((_404_v33).contains(_403_v32)) ? ((_404_v33).get(_403_v32)) : (_393_v23)))) : (new BigNumber((_405_v34).length)));
      _nw53[(_dafny.ONE).toNumber()] = (((_392_v22).contains(p0)) ? ((_392_v22).get(p0)) : ((_this).f6));
      _nw53[(new BigNumber(2)).toNumber()] = (new BigNumber((_408_v36).cardinality())).multipliedBy(new BigNumber(383));
      _nw53[(new BigNumber(3)).toNumber()] = (_this).f6;
      _nw53[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt((_this).f6, new BigNumber((_403_v32).length));
      _nw53[(new BigNumber(5)).toNumber()] = new BigNumber((_393_v23).length);
      _nw53[(new BigNumber(6)).toNumber()] = ((p1) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6)).length)) : ((_this).f6));
      _nw53[(new BigNumber(7)).toNumber()] = (_this).f6;
      _nw53[(new BigNumber(8)).toNumber()] = (_this).f6;
      _nw53[(new BigNumber(9)).toNumber()] = ((_this).f6).multipliedBy(new BigNumber((_module.__default.fm34(new BigNumber(((_409_v37)[_module.__default.safeIndex((_this).f6, new BigNumber((_409_v37).length))]).length), globalState)).length));
      _nw53[(new BigNumber(10)).toNumber()] = (_393_v23)[_module.__default.safeIndex(new BigNumber(39), new BigNumber((_393_v23).length))];
      _410_v38 = _nw53;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_410_v38).length))) {
        let _411_i5 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_411_i5)) && ((_411_i5).isLessThan(new BigNumber((_410_v38).length))))) {
          (_410_v38)[(_411_i5)] = (_411_i5).multipliedBy(new BigNumber(404));
        }
      }
      let _hi4 = (_this).f6;
      for (let _412_i7 = (_dafny.ZERO).minus((_this).f6); _412_i7.isLessThan(_hi4); _412_i7 = _412_i7.plus(_dafny.ONE)) {
        (globalState).f0 = p0;
        let _413_v39;
        _413_v39 = _dafny.Seq.of(p1);
        _413_v39 = _413_v39;
        let _414_v40;
        let _nw54 = Array((new BigNumber(9)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(p1);
        _nw54[(_dafny.ONE).toNumber()] = _413_v39;
        _nw54[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_413_v39, _413_v39);
        _nw54[(new BigNumber(3)).toNumber()] = _413_v39;
        _nw54[(new BigNumber(4)).toNumber()] = _413_v39;
        _nw54[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_413_v39, _413_v39);
        _nw54[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm3(globalState), p1, p0, p0, p1), _dafny.Seq.of(p0, p1));
        _nw54[(new BigNumber(7)).toNumber()] = _413_v39;
        _nw54[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_413_v39, _dafny.Seq.of(true, p1));
        _414_v40 = _nw54;
        let _index57 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_414_v40).length));
        (_414_v40)[_index57] = _413_v39;
        let _index58 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_414_v40).length));
        (_414_v40)[_index58] = _413_v39;
        let _415_v41;
        let _nw55 = Array((_dafny.ONE).toNumber()).fill([]);
        _415_v41 = _nw55;
        let _416_v42;
        let _nw56 = Array((new BigNumber(4)).toNumber());
        _nw56[(_dafny.ZERO).toNumber()] = _this.f9;
        _nw56[(_dafny.ONE).toNumber()] = _this.f9;
        _nw56[(new BigNumber(2)).toNumber()] = _this.f9;
        _nw56[(new BigNumber(3)).toNumber()] = _this.f9;
        _416_v42 = _nw56;
        let _index59 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_415_v41).length));
        (_415_v41)[_index59] = _416_v42;
        let _index60 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_415_v41).length));
        (_415_v41)[_index60] = _416_v42;
      }
      (globalState).f0 = p0;
      return;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f9 = [];
      this._f6 = _dafny.ZERO;
      this._f10 = _dafny.Map.Empty;
      this._f21 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    set f9(value) {
      let _this = this;
      _this._f9 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f21, f9, f10, f6) {
      let _this = this;
      (_this)._f21 = f21;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f6 = f6;
      return;
    }
    fm13(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("o"), _dafny.Seq.UnicodeFromString("qhkakwoox")));
    };
    fm5(p0, globalState) {
      let _this = this;
      return _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))))), (_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0))))).Intersect(_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0))));
    };
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.of(new BigNumber(-371), (_this).f6, (_this).f6, (_this).f6, (_dafny.ZERO).minus((_this).f6)));
    };
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _417_v0;
      _417_v0 = _dafny.Seq.of((_this).f6, new BigNumber(-821));
      let _418_v1;
      _418_v1 = _module.D3.create_DC9(_417_v0);
      let _pat_let_tv11 = _417_v0;
      (globalState).f5 = (_dafny.ZERO).minus(new BigNumber(((function (_pat_let13_0) {
        return function (_419_dt__update__tmp_h0) {
          return function (_pat_let14_0) {
            return function (_420_dt__update_hcf21_h0) {
              return _module.D3.create_DC9(_420_dt__update_hcf21_h0);
            }(_pat_let14_0);
          }(_pat_let_tv11);
        }(_pat_let13_0);
      }(_418_v1)).dtor_cf21).length));
      let _421_v2;
      let _nw57 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _421_v2 = _nw57;
      let _index61 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length));
      (_421_v2)[_index61] = p2;
      let _index62 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length));
      (_421_v2)[_index62] = (_this).f6;
      (globalState).f2 = _dafny.Seq.UnicodeFromString("djjjt");
      _417_v0 = _417_v0;
      let _422_v3;
      let _init6 = ((_423_v0) => function (_424_i0) {
        return _dafny.MultiSet.FromArray(_423_v0);
      })(_417_v0);
      let _nw58 = Array((new BigNumber(8)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw58.length); _i0_6++) {
        _nw58[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _422_v3 = _nw58;
      let _425_v4;
      _425_v4 = _module.D2.create_DC7(p1, _422_v3, ((_this).f6).minus((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]), ((_this).f6).minus(new BigNumber((_dafny.Seq.UnicodeFromString("olpetvk")).length)));
      let _source7 = _425_v4;
      if (_source7.is_DC6) {
        let _426___mcc_h0 = (_source7).cf13;
        let _427___mcc_h1 = (_source7).cf14;
        let _428_cf14 = _427___mcc_h1;
        let _429_cf13 = _426___mcc_h0;
        let _arr4 = _this.f9;
        let _index63 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_this.f9).length));
        _arr4[_index63] = p0;
        let _430_v5;
        _430_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,p0);
        let _431_v6;
        _431_v6 = _module.D0.create_DC0(p1, new BigNumber((p1).length));
        let _432_v7;
        _432_v7 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("hjgjnji"));
        let _pat_let_tv12 = _421_v2;
        let _pat_let_tv13 = _421_v2;
        let _433_v8;
        _433_v8 = _module.D1.create_DC4(function (_pat_let15_0) {
  return function (_434_dt__update__tmp_h1) {
    return function (_pat_let16_0) {
      return function (_435_dt__update_hcf1_h0) {
        return _module.D0.create_DC0((_434_dt__update__tmp_h1).dtor_cf0, _435_dt__update_hcf1_h0);
      }(_pat_let16_0);
    }((_pat_let_tv13)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_pat_let_tv12).length))]);
  }(_pat_let15_0);
}(_431_v6), _429_cf13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), function (_436_i1) {
  return new _dafny.CodePoint('t'.codePointAt(0));
}), true, new BigNumber((_432_v7).length));
        let _arr5 = _this.f9;
        let _index64 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_this.f9).length));
        let _rhs63 = new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe(_this.f9,p0)).Merge(_430_v5)).Merge(_430_v5)).length);
        let _rhs64 = (_433_v8).dtor_cf10;
        let _rhs65 = p0;
        let _lhs51 = globalState;
        let _lhs52 = globalState;
        let _lhs53 = _this.f9;
        let _lhs54 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_this.f9).length));
        _lhs51.f5 = _rhs63;
        _lhs52.f0 = _rhs64;
        _lhs53[_lhs54] = _rhs65;
        let _index65 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length));
        (_421_v2)[_index65] = _429_cf13;
        let _437_v9;
        _437_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_429_cf13, globalState),p0);
        let _438_v10;
        _438_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(((_437_v9).contains((_this).f6)) ? ((_437_v9).get((_this).f6)) : ((_this.f9)[_module.__default.safeIndex(new BigNumber(633), new BigNumber((_this.f9).length))])));
        _437_v9 = (_437_v9).update(_429_cf13, p0);
        let _arr6 = _this.f9;
        let _index66 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_this.f9).length));
        _arr6[_index66] = p0;
      } else if (_source7.is_DC7) {
        let _439___mcc_h2 = (_source7).cf15;
        let _440___mcc_h3 = (_source7).cf16;
        let _441___mcc_h4 = (_source7).cf17;
        let _442___mcc_h5 = (_source7).cf18;
        let _443_cf18 = _442___mcc_h5;
        let _444_cf17 = _441___mcc_h4;
        let _445_cf16 = _440___mcc_h3;
        let _446_cf15 = _439___mcc_h2;
        let _447_v11;
        _447_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,p0);
        _447_v11 = (_447_v11).update(_this.f9, p0);
        let _448_v12;
        _448_v12 = _dafny.MultiSet.fromElements(p0, p0, p0, !(true));
        let _449_v13;
        _449_v13 = _dafny.Seq.UnicodeFromString("o");
        let _rhs66 = _module.__default.safeDivisionInt((p2).plus(new BigNumber(((_448_v12).update(p0, _module.__default.abs(_443_cf18))).cardinality())), p2);
        let _rhs67 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lcenstrw"), _449_v13);
        let _rhs68 = (_this).f6;
        let _rhs69 = p0;
        let _lhs55 = globalState;
        let _lhs56 = globalState;
        let _lhs57 = globalState;
        let _lhs58 = globalState;
        _lhs55.f5 = _rhs66;
        _lhs56.f2 = _rhs67;
        _lhs57.f5 = _rhs68;
        _lhs58.f0 = _rhs69;
        let _450_v14;
        _450_v14 = _dafny.Set.fromElements(false, p0);
        let _451_v15;
        _451_v15 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_450_v14);
        _451_v15 = (_451_v15).update(p0, _450_v14);
        let _452_v16;
        let _nw59 = new _module.C1();
        _nw59.__ctor(_module.__default.safeDivisionInt(_444_cf17, new BigNumber(67)));
        _452_v16 = _nw59;
        let _453_v17;
        _453_v17 = new _dafny.CodePoint('o'.codePointAt(0));
        let _454_v18;
        _454_v18 = _dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC1(p0, (_452_v16).f6, _453_v17)).dtor_cf2,((p0) ? (_dafny.MultiSet.fromElements(p0, p0)) : (_dafny.MultiSet.FromArray(_446_cf15))));
        let _rhs70 = _452_v16;
        let _rhs71 = ((p0) ? (_454_v18) : (_454_v18));
        _452_v16 = _rhs70;
        _454_v18 = _rhs71;
      } else if (_source7.is_DC8) {
        let _455___mcc_h6 = (_source7).cf19;
        let _456___mcc_h7 = (_source7).cf20;
        let _457_cf20 = _456___mcc_h7;
        let _458_cf19 = _455___mcc_h6;
        if (((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]).isEqualTo(new BigNumber((function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(619)), ((_459_p1, _460_cf20) => function (_461_i2) {
            return _dafny.Seq.update(_459_p1, _module.__default.safeIndex(new BigNumber(339), new BigNumber((_459_p1).length)), _460_cf20);
          })(p1, _457_cf20))).Elements) {
            let _462_v19 = _compr_11;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(619)), ((_463_p1, _464_cf20) => function (_461_i2) {
              return _dafny.Seq.update(_463_p1, _module.__default.safeIndex(new BigNumber(339), new BigNumber((_463_p1).length)), _464_cf20);
            })(p1, _457_cf20)), _462_v19)) {
              _coll11.add(_462_v19);
            }
          }
          return _coll11;
        }()).length))) {
          let _465_v20;
          _465_v20 = _dafny.Seq.UnicodeFromString("hgq");
          (globalState).f2 = _465_v20;
          let _466_v21;
          let _467_v22;
          let _468_v23;
          let _out16;
          let _out17;
          let _out18;
          let _outcollector6 = (_this).m14(_module.__default.fm3(globalState), true, globalState);
          _out16 = _outcollector6[0];
          _out17 = _outcollector6[1];
          _out18 = _outcollector6[2];
          _466_v21 = _out16;
          _467_v22 = _out17;
          _468_v23 = _out18;
          (globalState).f5 = (p2).minus((_this).f6);
          _458_cf19 = !((_468_v23) === (p0));
          let _469_v24;
          let _nw60 = new _module.C0();
          _nw60.__ctor(_this.f9, p0);
          _469_v24 = _nw60;
        } else {
          let _470_v26;
          _470_v26 = new _dafny.CodePoint('r'.codePointAt(0));
          let _471_v27;
          _471_v27 = _dafny.Map.Empty.slice().updateUnsafe(_470_v26,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-134)), ((_472_v26) => function (_473_i3) {
            return _472_v26;
          })(_470_v26))).length));
          let _474_v28;
          _474_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_457_cf20);
          let _475_v29;
          _475_v29 = _dafny.MultiSet.fromElements(new BigNumber((_474_v28).length), p2, (_this).f6);
          let _476_v30;
          _476_v30 = _dafny.Map.Empty.slice().updateUnsafe((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))],p0);
          (globalState).f0 = !((_475_v29).update(new BigNumber((_476_v30).length), _module.__default.abs((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]))).contains(new BigNumber((function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of ((_471_v27).update(_470_v26, (_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))])).Keys.Elements) {
              let _477_v25 = _compr_12;
              if (((_471_v27).update(_470_v26, (_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))])).contains(_477_v25)) {
                _coll12.push([_477_v25,p0]);
              }
            }
            return _coll12;
          }()).length));
          let _478_v31;
          let _nw61 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _478_v31 = _nw61;
          _478_v31 = _478_v31;
          let _479_v32;
          _479_v32 = _dafny.Map.Empty.slice().updateUnsafe(_458_cf19,_this.f9);
          let _480_v33;
          _480_v33 = _dafny.Set.fromElements(new BigNumber(790));
          let _481_v34;
          _481_v34 = _dafny.MultiSet.fromElements(_458_cf19, _458_cf19);
          let _482_v35;
          _482_v35 = _dafny.MultiSet.fromElements(_module.__default.fm35(p2, new BigNumber((_480_v33).length), _458_cf19, new BigNumber((p1).length), globalState), _dafny.MultiSet.fromElements(new BigNumber(((_481_v34).update(false, _module.__default.abs((_this).f6))).cardinality())), _475_v29);
          let _483_v36;
          let _nw62 = new _module.C2();
          _nw62.__ctor(p2, _482_v35, _this.f9, (_this).f10);
          _483_v36 = _nw62;
          let _484_v37;
          let _init7 = ((_485_p1) => function (_486_i4) {
            return _485_p1;
          })(p1);
          let _nw63 = Array((new BigNumber(5)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw63.length); _i0_7++) {
            _nw63[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _484_v37 = _nw63;
          let _index67 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_484_v37).length));
          (_484_v37)[_index67] = p1;
          let _index68 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_484_v37).length));
          let _rhs72 = _479_v32;
          let _rhs73 = (_module.D6.create_DC18(_483_v36)).dtor_cf38;
          let _rhs74 = p1;
          let _lhs59 = _484_v37;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_484_v37).length));
          _479_v32 = _rhs72;
          _483_v36 = _rhs73;
          _lhs59[_lhs60] = _rhs74;
          let _index69 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length));
          (_421_v2)[_index69] = (_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))];
          let _487_v38;
          let _nw64 = new _module.C2();
          _nw64.__ctor((_483_v36).f6, (_482_v35).Union(_482_v35), _this.f9, (_this).f10);
          _487_v38 = _nw64;
        }
        _457_cf20 = ((_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(p2, (_this).f6), _417_v0)) ? (_458_cf19) : (_457_cf20));
        let _488_v39;
        _488_v39 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
        let _489_v40;
        _489_v40 = _module.D4.create_DC13(_417_v0, false, new BigNumber((_488_v39).length));
        _489_v40 = _489_v40;
        let _490_v41;
        _490_v41 = _module.D4.create_DC14(_457_cf20, _488_v39, (_this).f6, _dafny.Map.Empty.slice().updateUnsafe(true,p0));
        let _491_v42;
        _491_v42 = _dafny.Map.Empty.slice().updateUnsafe(_490_v41,((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_457_cf20,_457_cf20)).length)));
        _491_v42 = (_491_v42).update(_490_v41, _457_cf20);
      } else {
        let _492___mcc_h8 = (_source7).cf12;
        let _493_cf12 = _492___mcc_h8;
        let _494_v43;
        _494_v43 = _module.D2.create_DC8(p0, p0);
        let _495_v44;
        _495_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_417_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-880)), ((_496_v2) => function (_497_i5) {
          return (_496_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_496_v2).length))];
        })(_421_v2)))).length),(_494_v43).dtor_cf20);
        let _498_v45;
        _498_v45 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex((_this).f6, new BigNumber((p1).length))],p0);
        _495_v44 = (_495_v44).update(new BigNumber((_498_v45).length), true);
        let _499_v46;
        let _init8 = function (_500_i6) {
          return _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-810)), function (_501_i7) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }), _module.__default.safeIndex(new BigNumber(238), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-810)), function (_502_i7) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })).length)), new _dafny.CodePoint('l'.codePointAt(0))));
        };
        let _nw65 = Array((new BigNumber(27)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw65.length); _i0_8++) {
          _nw65[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _499_v46 = _nw65;
        let _503_v47;
        _503_v47 = new _dafny.CodePoint('b'.codePointAt(0));
        let _504_v48;
        _504_v48 = _dafny.Seq.of(_503_v47, _503_v47);
        let _505_v49;
        _505_v49 = _dafny.Set.fromElements(_504_v48);
        let _index70 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_499_v46).length));
        (_499_v46)[_index70] = _505_v49;
        let _index71 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_499_v46).length));
        (_499_v46)[_index71] = _505_v49;
        if (false) {
          _503_v47 = _503_v47;
          let _506_v50;
          _506_v50 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('l'.codePointAt(0)), _503_v47);
          let _507_v51;
          _507_v51 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(619), (_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]));
          let _508_v52;
          let _nw66 = new _module.C2();
          _nw66.__ctor((_dafny.ZERO).minus(new BigNumber((_506_v50).cardinality())), (_module.__default.fm36(globalState)).Difference(_507_v51), _this.f9, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_504_v48).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(670)), ((_509_v47) => function (_510_i8) {
            return _509_v47;
          })(_503_v47))));
          _508_v52 = _nw66;
          let _511_v53;
          _511_v53 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p2, (_this).f6),_421_v2);
          _511_v53 = _511_v53;
          let _arr7 = _this.f9;
          let _index72 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length));
          _arr7[_index72] = p0;
          let _512_v54;
          _512_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,_504_v48);
          let _513_v55;
          _513_v55 = _dafny.MultiSet.fromElements(p0);
          let _514_v56;
          _514_v56 = _dafny.MultiSet.fromElements((_this).f6);
          let _515_v57;
          _515_v57 = _dafny.Set.fromElements(new BigNumber(739), (_dafny.ZERO).minus((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]), (_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]);
          let _pat_let_tv14 = p1;
          let _516_v58;
          let _nw67 = Array((new BigNumber(16)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _512_v54;
          _nw67[(_dafny.ONE).toNumber()] = (_512_v54).update(p0, _504_v48);
          _nw67[(new BigNumber(2)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_504_v48);
          _nw67[(new BigNumber(4)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(5)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(6)).toNumber()] = (_this).fm13(new _dafny.CodePoint('s'.codePointAt(0)), _513_v55, p0, new BigNumber((_514_v56).cardinality()), globalState);
          _nw67[(new BigNumber(7)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(8)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_504_v48);
          _nw67[(new BigNumber(10)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(11)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(12)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(13)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(14)).toNumber()] = _512_v54;
          _nw67[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,(_module.D1.create_DC4(function (_pat_let17_0) {
  return function (_517_dt__update__tmp_h2) {
    return function (_pat_let18_0) {
      return function (_518_dt__update_hcf0_h0) {
        return _module.D0.create_DC0(_518_dt__update_hcf0_h0, (_517_dt__update__tmp_h2).dtor_cf1);
      }(_pat_let18_0);
    }(_pat_let_tv14);
  }(_pat_let17_0);
}(_module.D0.create_DC0(p1, (_this).f6)), (_this).f6, _504_v48, p0, new BigNumber((_515_v57).length))).dtor_cf9);
          _516_v58 = _nw67;
          let _index73 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_516_v58).length));
          (_516_v58)[_index73] = _512_v54;
          let _519_v59;
          _519_v59 = _module.D0.create_DC1(p0, (_this).f6, _503_v47);
          let _pat_let_tv15 = p0;
          let _arr8 = _this.f9;
          let _index74 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length));
          let _index75 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_516_v58).length));
          let _index76 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length));
          let _rhs75 = p0;
          let _rhs76 = _512_v54;
          let _rhs77 = new BigNumber((_dafny.Seq.update(_504_v48, _module.__default.safeIndex((p2).multipliedBy(p2), new BigNumber((_504_v48).length)), _503_v47)).length);
          let _rhs78 = function (_pat_let19_0) {
            return function (_520_dt__update__tmp_h3) {
              return function (_pat_let20_0) {
                return function (_521_dt__update_hcf2_h0) {
                  return _module.D0.create_DC1(_521_dt__update_hcf2_h0, (_520_dt__update__tmp_h3).dtor_cf3, (_520_dt__update__tmp_h3).dtor_cf4);
                }(_pat_let20_0);
              }(_pat_let_tv15);
            }(_pat_let19_0);
          }(_519_v59);
          let _rhs79 = _dafny.Seq.contains(p1, p0);
          let _lhs61 = _this.f9;
          let _lhs62 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length));
          let _lhs63 = _516_v58;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_516_v58).length));
          let _lhs65 = _421_v2;
          let _lhs66 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length));
          _lhs61[_lhs62] = _rhs75;
          _lhs63[_lhs64] = _rhs76;
          _lhs65[_lhs66] = _rhs77;
          _519_v59 = _rhs78;
          r1 = _rhs79;
          let _522_v60;
          _522_v60 = _dafny.Set.fromElements(p0, (_this.f9)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length))]);
          let _523_v61;
          _523_v61 = _module.D0.create_DC0(_dafny.Seq.update(p1, _module.__default.safeIndex(new BigNumber((_522_v60).length), new BigNumber((p1).length)), p0), new BigNumber(-92));
          let _524_v62;
          _524_v62 = _module.D1.create_DC4(_523_v61, (_this).f6, _504_v48, p0, p2);
          let _525_v63;
          _525_v63 = _dafny.Set.fromElements((((_495_v44).contains((_this).f6)) ? ((_495_v44).get((_this).f6)) : ((_this.f9)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length))])), (_524_v62).dtor_cf10);
          let _526_v64;
          _526_v64 = _dafny.Set.fromElements(_522_v60, _525_v63, _dafny.Set.fromElements((((_498_v45).contains(false)) ? ((_498_v45).get(false)) : ((_this.f9)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length))]))));
          let _527_v65;
          _527_v65 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_525_v63, _522_v60));
          let _528_v67;
          _528_v67 = _dafny.Map.Empty.slice().updateUnsafe(_525_v63,(_this).f6);
          let _529_v69;
          _529_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(p0, (_this.f9)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length))]),_522_v60);
          let _530_v71;
          _530_v71 = _dafny.Seq.of(_526_v64);
          let _531_v72;
          _531_v72 = _dafny.Map.Empty.slice().updateUnsafe(_525_v63,(_this.f9)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length))]);
          let _532_v74;
          let _nw68 = Array((new BigNumber(27)).toNumber());
          _nw68[(_dafny.ZERO).toNumber()] = _526_v64;
          _nw68[(_dafny.ONE).toNumber()] = _526_v64;
          _nw68[(new BigNumber(2)).toNumber()] = _526_v64;
          _nw68[(new BigNumber(3)).toNumber()] = _526_v64;
          _nw68[(new BigNumber(4)).toNumber()] = _526_v64;
          _nw68[(new BigNumber(5)).toNumber()] = (_526_v64).Intersect(_526_v64);
          _nw68[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(_525_v63, _525_v63, _522_v60);
          _nw68[(new BigNumber(7)).toNumber()] = _526_v64;
          _nw68[(new BigNumber(8)).toNumber()] = ((p0) ? (_526_v64) : (_526_v64));
          _nw68[(new BigNumber(9)).toNumber()] = function () {
            let _coll13 = new _dafny.Set();
            for (const _compr_13 of ((_527_v65)[_module.__default.safeIndex(new BigNumber(735), new BigNumber((_527_v65).length))]).Elements) {
              let _533_v66 = _compr_13;
              if (((_527_v65)[_module.__default.safeIndex(new BigNumber(735), new BigNumber((_527_v65).length))]).contains(_533_v66)) {
                _coll13.add(_533_v66);
              }
            }
            return _coll13;
          }();
          _nw68[(new BigNumber(10)).toNumber()] = (_526_v64).Union(_526_v64);
          _nw68[(new BigNumber(11)).toNumber()] = function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of (_528_v67).Keys.Elements) {
              let _534_v68 = _compr_14;
              if ((_528_v67).contains(_534_v68)) {
                _coll14.add(_534_v68);
              }
            }
            return _coll14;
          }();
          _nw68[(new BigNumber(12)).toNumber()] = _526_v64;
          _nw68[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(_dafny.Set.fromElements((p1)[_module.__default.safeIndex((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))], new BigNumber((p1).length))], (_this.f9)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length))]), _522_v60);
          _nw68[(new BigNumber(14)).toNumber()] = _526_v64;
          _nw68[(new BigNumber(15)).toNumber()] = (_526_v64).Difference(function () {
            let _coll15 = new _dafny.Set();
            for (const _compr_15 of (_529_v69).Keys.Elements) {
              let _535_v70 = _compr_15;
              if ((_529_v69).contains(_535_v70)) {
                _coll15.add(_535_v70);
              }
            }
            return _coll15;
          }());
          _nw68[(new BigNumber(16)).toNumber()] = (_530_v71)[_module.__default.safeIndex(new BigNumber((_513_v55).cardinality()), new BigNumber((_530_v71).length))];
          _nw68[(new BigNumber(17)).toNumber()] = function () {
            let _coll16 = new _dafny.Set();
            for (const _compr_16 of (_531_v72).Keys.Elements) {
              let _536_v73 = _compr_16;
              if ((_531_v72).contains(_536_v73)) {
                _coll16.add(_536_v73);
              }
            }
            return _coll16;
          }();
          _nw68[(new BigNumber(18)).toNumber()] = (_530_v71)[_module.__default.safeIndex((_this).f6, new BigNumber((_530_v71).length))];
          _nw68[(new BigNumber(19)).toNumber()] = (_526_v64).Intersect(_dafny.Set.fromElements(_522_v60, _522_v60));
          _nw68[(new BigNumber(20)).toNumber()] = _dafny.Set.fromElements(_525_v63);
          _nw68[(new BigNumber(21)).toNumber()] = (_dafny.Set.fromElements(_525_v63, _dafny.Set.fromElements(p0), _525_v63, _522_v60, _522_v60)).Intersect(_526_v64);
          _nw68[(new BigNumber(22)).toNumber()] = _526_v64;
          _nw68[(new BigNumber(23)).toNumber()] = _526_v64;
          _nw68[(new BigNumber(24)).toNumber()] = _526_v64;
          _nw68[(new BigNumber(25)).toNumber()] = ((p0) ? (_526_v64) : (_526_v64));
          _nw68[(new BigNumber(26)).toNumber()] = _526_v64;
          _532_v74 = _nw68;
          let _index77 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_532_v74).length));
          (_532_v74)[_index77] = _dafny.Set.fromElements(_525_v63);
          let _index78 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_532_v74).length));
          let _rhs80 = (p0) && ((_this.f9)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length))]);
          let _rhs81 = (_this.f9)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length))];
          let _rhs82 = (p0) && ((_this.f9)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_this.f9).length))]);
          let _rhs83 = _module.__default.fm2(((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]).minus(p2), globalState);
          let _rhs84 = _526_v64;
          let _lhs67 = globalState;
          let _lhs68 = globalState;
          let _lhs69 = globalState;
          let _lhs70 = globalState;
          let _lhs71 = _532_v74;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_532_v74).length));
          _lhs67.f0 = _rhs80;
          _lhs68.f0 = _rhs81;
          _lhs69.f0 = _rhs82;
          _lhs70.f5 = _rhs83;
          _lhs71[_lhs72] = _rhs84;
        } else {
          let _index79 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length));
          (_421_v2)[_index79] = (_dafny.ZERO).minus(((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f6))));
          let _537_v75;
          _537_v75 = _dafny.MultiSet.fromElements((_this).f6);
          let _538_v76;
          _538_v76 = _dafny.MultiSet.fromElements(_537_v75);
          let _539_v77;
          let _nw69 = new _module.C2();
          _nw69.__ctor((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))], (_538_v76).Difference(_538_v76), _this.f9, _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_504_v48));
          _539_v77 = _nw69;
          let _540_v78;
          _540_v78 = _dafny.Set.fromElements(p0);
          let _541_v79;
          _541_v79 = _module.D5.create_DC16(_540_v78);
          _541_v79 = _541_v79;
          (globalState).f5 = (_539_v77).fm21(new BigNumber((p1).length), globalState);
          _421_v2 = _421_v2;
        }
        (globalState).f5 = (_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))];
      }
      let _542_v80;
      _542_v80 = _module.D5.create_DC17(_417_v0, p0);
      let _source8 = _542_v80;
      if (_source8.is_DC17) {
        let _543___mcc_h9 = (_source8).cf36;
        let _544___mcc_h10 = (_source8).cf37;
        let _545_cf37 = _544___mcc_h10;
        let _546_cf36 = _543___mcc_h9;
        let _547_v81;
        _547_v81 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]);
        let _548_v82;
        _548_v82 = _dafny.Map.Empty.slice().updateUnsafe(_545_cf37,_545_cf37);
        let _549_v83;
        _549_v83 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm37(_545_cf37, globalState)).dtor_cf41,(((_548_v82).contains(p0)) ? ((_548_v82).get(p0)) : (_545_cf37)));
        _547_v81 = (_547_v81).update(new BigNumber((_dafny.Seq.Concat(p1, p1)).length), new BigNumber((_549_v83).length));
        let _550_v84;
        let _nw70 = new _module.C1();
        _nw70.__ctor(p2);
        _550_v84 = _nw70;
        let _551_v85;
        _551_v85 = _dafny.Seq.of(_550_v84, _550_v84);
        let _552_v86;
        _552_v86 = _dafny.Map.Empty.slice().updateUnsafe((_551_v85)[_module.__default.safeIndex(new BigNumber((p1).length), new BigNumber((_551_v85).length))],_545_cf37);
        let _arr9 = _this.f9;
        let _index80 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_this.f9).length));
        _arr9[_index80] = (((_552_v86).contains(_550_v84)) ? ((_552_v86).get(_550_v84)) : (_545_cf37));
        let _553_v87;
        _553_v87 = new _dafny.CodePoint('r'.codePointAt(0));
        let _554_v88;
        _554_v88 = _dafny.MultiSet.fromElements(_553_v87);
        let _555_v89;
        _555_v89 = _dafny.Map.Empty.slice().updateUnsafe(_545_cf37,_554_v88);
        let _556_v90;
        _556_v90 = _dafny.Seq.of(_554_v88, _dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), _553_v87, _553_v87)), _554_v88, (((_555_v89).contains(_545_cf37)) ? ((_555_v89).get(_545_cf37)) : ((_554_v88).update(new _dafny.CodePoint('c'.codePointAt(0)), _module.__default.abs(p2)))));
        let _557_v91;
        _557_v91 = _module.D4.create_DC12(_556_v90);
        let _arr10 = _this.f9;
        let _index81 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_this.f9).length));
        let _rhs85 = false;
        let _rhs86 = ((new BigNumber(182)).minus(_module.__default.fm2(new BigNumber(565), globalState))).isLessThan(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(p1), _module.__default.safeIndex(new BigNumber(237), new BigNumber((_dafny.Seq.of(p1)).length)), _dafny.Seq.of(p0, !(_545_cf37), true, p0)), _dafny.Seq.of(p1))).length));
        let _rhs87 = _557_v91;
        let _lhs73 = _this.f9;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_this.f9).length));
        _lhs73[_lhs74] = _rhs85;
        _545_cf37 = _rhs86;
        _557_v91 = _rhs87;
        let _558_v92;
        _558_v92 = _dafny.MultiSet.fromElements(new BigNumber(-930));
        let _559_v93;
        let _nw71 = Array((new BigNumber(10)).toNumber());
        _559_v93 = _nw71;
        let _560_v94;
        _560_v94 = _module.D0.create_DC1(_545_cf37, (_this).f6, _553_v87);
        let _561_v95;
        _561_v95 = _dafny.Seq.UnicodeFromString("vst");
        let _562_v96;
        let _init9 = ((_563_p1) => function (_564_i11) {
          return _module.D0.create_DC0(_563_p1, (_this).f6);
        })(p1);
        let _nw72 = Array((new BigNumber(15)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw72.length); _i0_9++) {
          _nw72[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _562_v96 = _nw72;
        let _565_v97;
        _565_v97 = _module.D7.create_DC21(_562_v96);
        let _566_v98;
        _566_v98 = _dafny.Set.fromElements((_565_v97).dtor_cf43);
        let _rhs88 = ((_545_cf37) ? ((_560_v94).dtor_cf2) : (((_this.f9)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_this.f9).length))]) === (_545_cf37)));
        let _rhs89 = _module.__default.fm35((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))], ((_this).f6).multipliedBy((_dafny.ZERO).minus((_this).f6)), false, p2, globalState);
        let _rhs90 = _421_v2;
        let _rhs91 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(928)), ((_567_v87) => function (_568_i9) {
          return _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), function (_569_i10) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }), _module.__default.safeIndex((_this).f6, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), function (_570_i10) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          })).length)), _567_v87);
        })(_553_v87)), _module.__default.safeIndex(new BigNumber(-96), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(928)), ((_571_v87) => function (_572_i9) {
          return _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), function (_573_i10) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }), _module.__default.safeIndex((_this).f6, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), function (_574_i10) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          })).length)), _571_v87);
        })(_553_v87))).length)), _561_v95)).length);
        let _rhs92 = ((!(_566_v98).equals(_566_v98)) ? (_559_v93) : (_559_v93));
        let _lhs75 = globalState;
        _545_cf37 = _rhs88;
        _558_v92 = _rhs89;
        _421_v2 = _rhs90;
        _lhs75.f5 = _rhs91;
        _559_v93 = _rhs92;
        let _arr11 = _this.f9;
        let _index82 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_this.f9).length));
        _arr11[_index82] = p0;
      } else {
        let _575___mcc_h11 = (_source8).cf35;
        let _576_cf35 = _575___mcc_h11;
        let _577_v99;
        _577_v99 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(p1)).cardinality()),_dafny.Seq.Create(_module.__default.abs(new BigNumber(416)), function (_578_i14) {
          return (_this).f6;
        }));
        let _579_v100;
        _579_v100 = _dafny.MultiSet.fromElements(p0, true);
        let _580_v101;
        _580_v101 = new _dafny.CodePoint('h'.codePointAt(0));
        let _581_v102;
        _581_v102 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), _580_v101);
        let _582_v103;
        _582_v103 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_dafny.Seq.update(_417_v0, _module.__default.safeIndex(new BigNumber(782), new BigNumber((_417_v0).length)), (((_579_v100).contains(p0)) ? ((_579_v100).get(p0)) : (new BigNumber((_581_v102).cardinality())))));
        let _583_v104;
        _583_v104 = _dafny.MultiSet.fromElements((_this).f6, (_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]);
        let _584_v105;
        _584_v105 = _dafny.Seq.of(_417_v0);
        let _585_v106;
        _585_v106 = _dafny.Seq.UnicodeFromString("nhw");
        let _586_v107;
        _586_v107 = _dafny.Seq.of(_417_v0, _417_v0, (_584_v105)[_module.__default.safeIndex(_module.__default.fm2(new BigNumber((_585_v106).length), globalState), new BigNumber((_584_v105).length))]);
        let _587_v108;
        let _nw73 = Array((new BigNumber(14)).toNumber());
        _nw73[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(49)), ((_588_v2) => function (_589_i12) {
          return (_588_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_588_v2).length))];
        })(_421_v2)), _417_v0);
        _nw73[(_dafny.ONE).toNumber()] = _417_v0;
        _nw73[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_417_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(332)), ((_590_p2) => function (_591_i13) {
          return _590_p2;
        })(p2)));
        _nw73[(new BigNumber(3)).toNumber()] = (((_577_v99).contains(new BigNumber((_417_v0).length))) ? ((_577_v99).get(new BigNumber((_417_v0).length))) : (_dafny.Seq.of(p2, (_this).f6)));
        _nw73[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_417_v0, _module.__default.safeIndex((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))], new BigNumber((_417_v0).length)), _module.__default.fm2((_this).f6, globalState));
        _nw73[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_417_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(640)), function (_592_i15) {
          return (_this).f6;
        }));
        _nw73[(new BigNumber(6)).toNumber()] = _417_v0;
        _nw73[(new BigNumber(7)).toNumber()] = _417_v0;
        _nw73[(new BigNumber(8)).toNumber()] = (((_582_v103).contains(p0)) ? ((_582_v103).get(p0)) : (_dafny.Seq.update(_417_v0, _module.__default.safeIndex(new BigNumber((_583_v104).cardinality()), new BigNumber((_417_v0).length)), (_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))])));
        _nw73[(new BigNumber(9)).toNumber()] = _417_v0;
        _nw73[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_417_v0, _417_v0);
        _nw73[(new BigNumber(11)).toNumber()] = (_584_v105)[_module.__default.safeIndex((_this).f6, new BigNumber((_584_v105).length))];
        _nw73[(new BigNumber(12)).toNumber()] = ((!(!(!(p0)))) ? (_417_v0) : (_417_v0));
        _nw73[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))], (_this).f6, (_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))], new BigNumber(953), p2), _dafny.Seq.of((_421_v2)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_421_v2).length))]));
        _587_v108 = _nw73;
        let _init10 = ((_593_v0, _594_p1) => function (_595_i16) {
          return _dafny.Seq.Concat(_593_v0, _dafny.Seq.update(_593_v0, _module.__default.safeIndex((_this).f6, new BigNumber((_593_v0).length)), new BigNumber((_594_p1).length)));
        })(_417_v0, p1);
        let _nw74 = Array((new BigNumber(28)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw74.length); _i0_10++) {
          _nw74[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _587_v108 = _nw74;
        let _596_v109;
        _596_v109 = _module.D5.create_DC16(_576_cf35);
        let _597_v110;
        _597_v110 = _module.D5.create_DC16((_596_v109).dtor_cf35);
        let _pat_let_tv16 = globalState;
        let _pat_let_tv17 = _576_cf35;
        let _rhs93 = p2;
        let _rhs94 = function (_pat_let21_0) {
          return function (_598_dt__update__tmp_h4) {
            return function (_pat_let22_0) {
              return function (_599_dt__update_hcf35_h0) {
                return _module.D5.create_DC16(_599_dt__update_hcf35_h0);
              }(_pat_let22_0);
            }(((_module.__default.fm38(_pat_let_tv16)).dtor_cf35).Union(_pat_let_tv17));
          }(_pat_let21_0);
        }(_597_v110);
        let _lhs76 = globalState;
        _lhs76.f5 = _rhs93;
        _596_v109 = _rhs94;
        (globalState).f5 = ((p0) ? ((_this).f6) : (p2));
        let _600_v111;
        _600_v111 = _dafny.Seq.of(_585_v106, _585_v106);
        (globalState).f5 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_600_v111, _dafny.Seq.of(_585_v106, _dafny.Seq.UnicodeFromString("f")))).length), new BigNumber((_585_v106).length));
      }
      let _601_v112;
      _601_v112 = _dafny.Set.fromElements(new BigNumber(-87));
      r0 = _module.__default.fm29(_601_v112, globalState);
      r1 = p0;
      return [r0, r1];
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = [];
      let r3 = _dafny.ZERO;
      let _602_v0;
      _602_v0 = new _dafny.CodePoint('j'.codePointAt(0));
      r0 = _602_v0;
      (_this).f9 = ((_this.f9));
      let _603_v1;
      let _init11 = function (_604_i0) {
        return (_604_i0).plus((_this).f6);
      };
      let _nw75 = Array((new BigNumber(9)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw75.length); _i0_11++) {
        _nw75[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _603_v1 = _nw75;
      let _index83 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length));
      (_603_v1)[_index83] = (_this).f6;
      let _605_v2;
      _605_v2 = false;
      let _index84 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length));
      let _rhs95 = new BigNumber(-180);
      let _rhs96 = new BigNumber(84);
      let _rhs97 = _605_v2;
      let _lhs77 = _603_v1;
      let _lhs78 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length));
      let _lhs79 = globalState;
      r3 = _rhs95;
      _lhs77[_lhs78] = _rhs96;
      _lhs79.f0 = _rhs97;
      (globalState).f0 = _605_v2;
      let _606_v3;
      _606_v3 = _dafny.Map.Empty.slice().updateUnsafe(_605_v2,p1);
      let _607_v5;
      _607_v5 = _dafny.Seq.of((_603_v1)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length))], p1, (_603_v1)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length))]);
      if ((new BigNumber((_module.__default.fm1(globalState)).length)).isLessThanOrEqualTo(new BigNumber(((_dafny.Set.fromElements(_606_v3, _dafny.Map.Empty.slice().updateUnsafe(!(_605_v2),new BigNumber((function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_dafny.Map.Empty.slice().updateUnsafe(_607_v5,_605_v2)).Keys.Elements) {
          let _608_v4 = _compr_17;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_607_v5,_605_v2)).contains(_608_v4)) {
            _coll17.push([_608_v4,_605_v2]);
          }
        }
        return _coll17;
      }()).length)))).Union(_dafny.Set.fromElements(_606_v3, _606_v3))).length))) {
        let _609_v6;
        let _nw76 = new _module.C0();
        _nw76.__ctor(_this.f9, _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("gxtawrg"), _dafny.Seq.UnicodeFromString("qi")));
        _609_v6 = _nw76;
        (globalState).f0 = _module.__default.fm3(globalState);
        let _610_v7;
        _610_v7 = _dafny.Seq.of(_609_v6.f20);
        let _611_v8;
        _611_v8 = _module.D0.create_DC0(_610_v7, (_603_v1)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length))]);
        let _612_v9;
        _612_v9 = _dafny.Seq.UnicodeFromString("fsmbn");
        let _613_v10;
        _613_v10 = _module.D1.create_DC4(_611_v8, p1, _612_v9, _605_v2, (_this).f6);
        let _614_v11;
        let _615_v12;
        let _616_v13;
        let _out19;
        let _out20;
        let _out21;
        let _outcollector7 = (_this).m14((_613_v10).dtor_cf10, _609_v6.f20, globalState);
        _out19 = _outcollector7[0];
        _out20 = _outcollector7[1];
        _out21 = _outcollector7[2];
        _614_v11 = _out19;
        _615_v12 = _out20;
        _616_v13 = _out21;
        let _617_v14;
        let _nw77 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
        _617_v14 = _nw77;
        let _618_v15;
        _618_v15 = _dafny.Map.Empty.slice().updateUnsafe(_617_v14,_605_v2);
        _618_v15 = (_618_v15).update(_617_v14, _module.__default.fm3(globalState));
        let _619_v16;
        _619_v16 = _dafny.MultiSet.fromElements(_603_v1);
        (_609_v6).f20 = (_619_v16).IsProperSubsetOf((_dafny.MultiSet.fromElements(_603_v1)).Intersect(_619_v16));
      } else {
        let _620_v17;
        let _nw78 = Array((new BigNumber(13)).toNumber()).fill(_module.D2.Default());
        _620_v17 = _nw78;
        let _nw79 = Array((new BigNumber(25)).toNumber()).fill(_module.D2.Default());
        _620_v17 = _nw79;
        let _621_v18;
        _621_v18 = _dafny.MultiSet.fromElements((p0).update((_603_v1)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length))], _module.__default.abs((_this).f6)));
        let _622_v19;
        let _nw80 = new _module.C2();
        _nw80.__ctor((_603_v1)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length))], (_621_v18).update(p0, _module.__default.abs((_603_v1)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length))])), _this.f9, _module.__default.fm39(_602_v0, globalState));
        _622_v19 = _nw80;
        let _623_v20;
        _623_v20 = _dafny.Set.fromElements((_603_v1)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length))], (_this).f6, (p1).plus(_module.__default.fm2((_this).f6, globalState)));
        _623_v20 = _623_v20;
        let _624_v21;
        let _nw81 = Array((new BigNumber(24)).toNumber()).fill([]);
        _624_v21 = _nw81;
        let _625_v22;
        let _nw82 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Set.Empty);
        _625_v22 = _nw82;
        let _index85 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_624_v21).length));
        (_624_v21)[_index85] = _625_v22;
        let _index86 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_624_v21).length));
        (_624_v21)[_index86] = _625_v22;
        if (!(_605_v2) || (_605_v2)) {
          let _626_v23;
          let _627_v24;
          let _628_v25;
          let _out22;
          let _out23;
          let _out24;
          let _outcollector8 = (_this).m14(true, _605_v2, globalState);
          _out22 = _outcollector8[0];
          _out23 = _outcollector8[1];
          _out24 = _outcollector8[2];
          _626_v23 = _out22;
          _627_v24 = _out23;
          _628_v25 = _out24;
          let _629_v26;
          _629_v26 = _this.f9;
          _629_v26 = _629_v26;
          let _630_v27;
          _630_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          _605_v2 = !(_626_v23).isEqualTo((((_630_v27).contains(p0)) ? ((_630_v27).get(p0)) : ((_this).f6)));
          let _631_v28;
          _631_v28 = _dafny.Seq.of(_628_v25);
          let _632_v29;
          _632_v29 = _dafny.Map.Empty.slice().updateUnsafe(_628_v25,_628_v25);
          let _633_v30;
          _633_v30 = _dafny.Map.Empty.slice().updateUnsafe(!((_631_v28)[_module.__default.safeIndex(new BigNumber(255), new BigNumber((_631_v28).length))]),(((_632_v29).contains(_628_v25)) ? ((_632_v29).get(_628_v25)) : (true)));
          r1 = ((_dafny.ZERO).minus((((_606_v3).contains(_605_v2)) ? ((_606_v3).get(_605_v2)) : (new BigNumber((_633_v30).length))))).plus(new BigNumber(116));
          _623_v20 = _623_v20;
        } else {
          let _634_v31;
          _634_v31 = _module.D9.create_DC24(_621_v18);
          let _635_v32;
          let _nw83 = new _module.C2();
          _nw83.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("ytxvws")).length), (_621_v18).Union((_634_v31).dtor_cf45), _this.f9, (_this).f10);
          _635_v32 = _nw83;
          let _636_v33;
          let _nw84 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _636_v33 = _nw84;
          _636_v33 = _636_v33;
          _603_v1 = _603_v1;
          let _637_v34;
          _637_v34 = _dafny.Set.fromElements(_605_v2, _605_v2, true);
          (globalState).f5 = new BigNumber((_637_v34).length);
          let _638_v35;
          let _nw85 = new _module.C0();
          _nw85.__ctor(_this.f9, !(_605_v2));
          _638_v35 = _nw85;
        }
      }
      let _639_v36;
      let _nw86 = Array((new BigNumber(6)).toNumber());
      _nw86[(_dafny.ZERO).toNumber()] = (_this).f10;
      _nw86[(_dafny.ONE).toNumber()] = (_this).f10;
      _nw86[(new BigNumber(2)).toNumber()] = ((_this).f10).Merge((_this).f10);
      _nw86[(new BigNumber(3)).toNumber()] = (_this).f10;
      _nw86[(new BigNumber(4)).toNumber()] = (_this).f10;
      _nw86[(new BigNumber(5)).toNumber()] = (_this).f10;
      _639_v36 = _nw86;
      let _640_v37;
      _640_v37 = _dafny.Seq.UnicodeFromString("rtg");
      let _index87 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_639_v36).length));
      (_639_v36)[_index87] = _dafny.Map.Empty.slice().updateUnsafe((_603_v1)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length))],_640_v37);
      let _index88 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_639_v36).length));
      (_639_v36)[_index88] = (_this).f10;
      r0 = _602_v0;
      let _641_v38;
      _641_v38 = _dafny.Seq.of(_605_v2, _module.__default.fm3(globalState));
      r1 = ((_603_v1)[_module.__default.safeIndex(new BigNumber(211), new BigNumber((_603_v1).length))]).minus(new BigNumber((_641_v38).length));
      let _nw87 = Array((new BigNumber(7)).toNumber());
      _nw87[(_dafny.ZERO).toNumber()] = _605_v2;
      _nw87[(_dafny.ONE).toNumber()] = false;
      _nw87[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsPrefixOf(_641_v38, _dafny.Seq.update(_641_v38, _module.__default.safeIndex((_this).f6, new BigNumber((_641_v38).length)), _module.__default.fm3(globalState)));
      _nw87[(new BigNumber(3)).toNumber()] = (new BigNumber(-757)).isEqualTo(new BigNumber((_607_v5).length));
      _nw87[(new BigNumber(4)).toNumber()] = ((_605_v2) ? (_605_v2) : (_605_v2));
      _nw87[(new BigNumber(5)).toNumber()] = _605_v2;
      _nw87[(new BigNumber(6)).toNumber()] = !(_605_v2) || (true);
      r2 = _nw87;
      r3 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_640_v37).length), (_this).f6));
      return [r0, r1, r2, r3];
    }
    m14(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = new _dafny.CodePoint('D'.codePointAt(0));
      let r2 = false;
      if (p1) {
        r0 = (_this).f6;
        (globalState).f5 = (_this).f6;
        let _642_v0;
        let _nw88 = new _module.C0();
        _nw88.__ctor(_this.f9, p0);
        _642_v0 = _nw88;
        let _643_v1;
        _643_v1 = _dafny.Seq.of(((_this).f6).minus((_this).f6), (_this).f6);
        _643_v1 = _dafny.Seq.Concat(_643_v1, _643_v1);
        let _644_v2;
        _644_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p0);
        _644_v2 = (_644_v2).update(new BigNumber(-864), _642_v0.f20);
      } else {
        let _rhs98 = p0;
        let _rhs99 = (_this).f6;
        let _lhs80 = globalState;
        let _lhs81 = globalState;
        _lhs80.f0 = _rhs98;
        _lhs81.f5 = _rhs99;
        let _arr12 = _this.f9;
        let _index89 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length));
        _arr12[_index89] = p0;
        let _arr13 = _this.f9;
        let _index90 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length));
        _arr13[_index90] = p0;
        (globalState).f5 = ((p1) ? ((_this).f6) : ((_this).f6));
        if (!(p0)) {
          let _645_v3;
          _645_v3 = _dafny.Seq.of(p1);
          let _646_v4;
          _646_v4 = _module.D0.create_DC0(_645_v3, (_this).f6);
          let _647_v5;
          _647_v5 = new _dafny.CodePoint('v'.codePointAt(0));
          let _648_v6;
          _648_v6 = _module.D1.create_DC4(_646_v4, (_this).f6, _module.__default.fm30(p0, _647_v5, globalState), (_this.f9)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length))], (_this).f6);
          (globalState).f2 = (_648_v6).dtor_cf9;
          let _649_v7;
          _649_v7 = _dafny.Set.fromElements(_647_v5, _647_v5, _647_v5);
          let _650_v8;
          _650_v8 = _dafny.MultiSet.fromElements(_649_v7, _649_v7, _649_v7);
          let _651_v9;
          _651_v9 = _dafny.Set.fromElements((_this.f9)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length))], p0);
          let _652_v10;
          _652_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_651_v9).length),p0);
          let _653_v11;
          _653_v11 = _dafny.Map.Empty.slice().updateUnsafe(_652_v10,p1);
          let _654_v12;
          _654_v12 = _module.D3.create_DC10(_653_v11);
          let _655_v13;
          _655_v13 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_module.__default.fm33(_654_v12, (_this).f6, false, p1, globalState)).length))).length));
          let _656_v14;
          _656_v14 = _dafny.Seq.UnicodeFromString("b");
          let _657_v15;
          let _nw89 = Array((new BigNumber(22)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = (_this).f6;
          _nw89[(_dafny.ONE).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_this).f6);
          _nw89[(new BigNumber(3)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(4)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(5)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(6)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(7)).toNumber()] = ((_this).f6).minus((_this).f6);
          _nw89[(new BigNumber(8)).toNumber()] = new BigNumber(247);
          _nw89[(new BigNumber(9)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(10)).toNumber()] = ((_this).f6).multipliedBy(new BigNumber(558));
          _nw89[(new BigNumber(11)).toNumber()] = ((_this).f6).minus(_module.__default.fm2((_dafny.ZERO).minus((_this).f6), globalState));
          _nw89[(new BigNumber(12)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(13)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(14)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(15)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(16)).toNumber()] = new BigNumber((_650_v8).cardinality());
          _nw89[(new BigNumber(17)).toNumber()] = ((_this).f6).plus(new BigNumber((_655_v13).length));
          _nw89[(new BigNumber(18)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_656_v14, _656_v14)).length);
          _nw89[(new BigNumber(20)).toNumber()] = (_this).f6;
          _nw89[(new BigNumber(21)).toNumber()] = (_this).f6;
          _657_v15 = _nw89;
          let _index91 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_657_v15).length));
          (_657_v15)[_index91] = (_this).f6;
          let _index92 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_657_v15).length));
          (_657_v15)[_index92] = new BigNumber((_dafny.Seq.of(_647_v5)).length);
          let _658_v16;
          let _nw90 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _658_v16 = _nw90;
          let _659_v17;
          _659_v17 = _dafny.Seq.of(_658_v16, _658_v16);
          let _660_v18;
          _660_v18 = _module.D9.create_DC26((_659_v17)[_module.__default.safeIndex((_this).f6, new BigNumber((_659_v17).length))]);
          let _index93 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_657_v15).length));
          let _arr14 = _this.f9;
          let _index94 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length));
          let _rhs100 = _660_v18;
          let _rhs101 = _module.__default.safeModuloInt((_657_v15)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_657_v15).length))], new BigNumber((_dafny.Seq.Concat(_655_v13, _655_v13)).length));
          let _rhs102 = p0;
          let _lhs82 = _657_v15;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_657_v15).length));
          let _lhs84 = _this.f9;
          let _lhs85 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length));
          _660_v18 = _rhs100;
          _lhs82[_lhs83] = _rhs101;
          _lhs84[_lhs85] = _rhs102;
          let _661_v19;
          _661_v19 = _dafny.Set.fromElements(_module.__default.safeDivisionInt((_657_v15)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_657_v15).length))], new BigNumber((_656_v14).length)));
          _661_v19 = function () {
            let _coll18 = new _dafny.Set();
            for (const _compr_18 of (_655_v13).Elements) {
              let _662_v20 = _compr_18;
              if (_dafny.Seq.contains(_655_v13, _662_v20)) {
                _coll18.add(_module.__default.safeDivisionInt(_662_v20, (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(-986)),_module.D4.create_DC12(_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0))))))).length)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(!(!(!(false))))).length), new BigNumber((_dafny.Seq.UnicodeFromString("ehssnh")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-954))).length)))).cardinality()))).cardinality()))));
              }
            }
            return _coll18;
          }();
          let _663_v21;
          let _nw91 = Array((new BigNumber(28)).toNumber());
          _nw91[(_dafny.ZERO).toNumber()] = _647_v5;
          _nw91[(_dafny.ONE).toNumber()] = _647_v5;
          _nw91[(new BigNumber(2)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(3)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
          _nw91[(new BigNumber(5)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(6)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(7)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(8)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(9)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
          _nw91[(new BigNumber(11)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(12)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(13)).toNumber()] = (_module.D0.create_DC1(p1, (_this).f6, _647_v5)).dtor_cf4;
          _nw91[(new BigNumber(14)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(15)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(16)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(17)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(18)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(19)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(20)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(21)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(22)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(23)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(24)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(25)).toNumber()] = _647_v5;
          _nw91[(new BigNumber(26)).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
          _nw91[(new BigNumber(27)).toNumber()] = _647_v5;
          _663_v21 = _nw91;
          let _index95 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_663_v21).length));
          (_663_v21)[_index95] = _647_v5;
          let _664_v22;
          _664_v22 = _dafny.Map.Empty.slice().updateUnsafe(!((_this.f9)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length))]),p1);
          let _index96 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_663_v21).length));
          (_663_v21)[_index96] = _module.__default.fm0(new BigNumber((_645_v3).length), _dafny.Set.fromElements((_this).f6, new BigNumber((_645_v3).length)), (((_664_v22).contains(p0)) ? ((_664_v22).get(p0)) : ((_this.f9)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length))])), globalState);
        } else {
          let _665_v23;
          _665_v23 = _dafny.MultiSet.fromElements((_this).f6, new BigNumber((_dafny.Seq.UnicodeFromString("idv")).length));
          (globalState).f5 = (((_this).f6).minus(new BigNumber((_665_v23).cardinality()))).plus(_module.__default.safeDivisionInt((_this).f6, (_this).f6));
          let _666_v24;
          _666_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6);
          _666_v24 = (_666_v24).update(p0, (_this).f6);
          let _667_v25;
          let _init12 = function (_668_i0) {
            return _dafny.Seq.of((_this).f6, (_this).f6);
          };
          let _nw92 = Array((new BigNumber(13)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw92.length); _i0_12++) {
            _nw92[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _667_v25 = _nw92;
          _667_v25 = _667_v25;
          let _669_v26;
          let _nw93 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _669_v26 = _nw93;
          let _index97 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_669_v26).length));
          (_669_v26)[_index97] = (_this).f6;
          let _index98 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_669_v26).length));
          let _rhs103 = p1;
          let _rhs104 = (_this).f6;
          let _lhs86 = globalState;
          let _lhs87 = _669_v26;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_669_v26).length));
          _lhs86.f0 = _rhs103;
          _lhs87[_lhs88] = _rhs104;
          let _670_v27;
          _670_v27 = _dafny.Seq.of((_this.f9)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length))], _module.__default.fm3(globalState), p1);
          let _671_v28;
          let _nw94 = Array((new BigNumber(17)).toNumber()).fill(false);
          _671_v28 = _nw94;
          let _672_v29;
          _672_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_671_v28);
          let _673_v30;
          _673_v30 = _dafny.Set.fromElements(_module.__default.safeDivisionInt((_669_v26)[_module.__default.safeIndex(new BigNumber(339), new BigNumber((_669_v26).length))], (_669_v26)[_module.__default.safeIndex(new BigNumber(339), new BigNumber((_669_v26).length))]), (new BigNumber(263)).multipliedBy(new BigNumber((_672_v29).length)));
          let _674_v31;
          _674_v31 = _dafny.Seq.of((_this).f6);
          let _675_v32;
          _675_v32 = _module.D2.create_DC6((_669_v26)[_module.__default.safeIndex(new BigNumber(339), new BigNumber((_669_v26).length))], _669_v26);
          let _rhs105 = _669_v26;
          let _rhs106 = new BigNumber((_dafny.Seq.Concat(_674_v31, _dafny.Seq.Concat(_674_v31, _674_v31))).length);
          let _rhs107 = _dafny.Seq.Concat(_670_v27, _dafny.Seq.update(_670_v27, _module.__default.safeIndex((_this).f6, new BigNumber((_670_v27).length)), (_this.f9)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length))]));
          let _rhs108 = (_673_v30).Intersect(_dafny.Set.fromElements(new BigNumber(((_this).f21).length)));
          let _rhs109 = (_675_v32).dtor_cf13;
          let _lhs89 = globalState;
          _669_v26 = _rhs105;
          r0 = _rhs106;
          _670_v27 = _rhs107;
          _673_v30 = _rhs108;
          _lhs89.f5 = _rhs109;
        }
        let _676_v33;
        _676_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,true);
        let _677_v34;
        _677_v34 = _dafny.Seq.of((_this).f6, (_this).f6, new BigNumber((_676_v33).length));
        let _678_v35;
        _678_v35 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f6), (_this).f6, (_dafny.ZERO).minus((_this).f6));
        let _679_v36;
        _679_v36 = _dafny.Seq.UnicodeFromString("jfb");
        let _680_v37;
        _680_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(globalState),p0);
        let _681_v38;
        _681_v38 = _module.D4.create_DC14(p1, _dafny.Map.Empty.slice().updateUnsafe(p1,(_this.f9)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length))]), _module.__default.fm2((_this).f6, globalState), _680_v37);
        let _682_v39;
        _682_v39 = new _dafny.CodePoint('i'.codePointAt(0));
        let _683_v40;
        _683_v40 = _dafny.MultiSet.fromElements((_this).f6);
        let _684_v41;
        _684_v41 = _dafny.MultiSet.fromElements(new BigNumber((_683_v40).cardinality()), new BigNumber(877));
        let _685_v42;
        _685_v42 = _module.D0.create_DC1((_this.f9)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_this.f9).length))], (_this).f6, _682_v39);
        let _686_v43;
        _686_v43 = _dafny.Seq.of(_679_v36);
        let _687_v44;
        let _nw95 = Array((new BigNumber(24)).toNumber());
        _nw95[(_dafny.ZERO).toNumber()] = new BigNumber(-382);
        _nw95[(_dafny.ONE).toNumber()] = (_this).f6;
        _nw95[(new BigNumber(2)).toNumber()] = (_this).f6;
        _nw95[(new BigNumber(3)).toNumber()] = (_this).f6;
        _nw95[(new BigNumber(4)).toNumber()] = new BigNumber((_679_v36).length);
        _nw95[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_module.__default.fm0((_this).f6, _678_v35, p0, globalState), _682_v39)).cardinality()))));
        _nw95[(new BigNumber(6)).toNumber()] = new BigNumber((_679_v36).length);
        _nw95[(new BigNumber(7)).toNumber()] = (_this).f6;
        _nw95[(new BigNumber(8)).toNumber()] = (_this).f6;
        _nw95[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_682_v39), _module.__default.safeIndex((_this).f6, new BigNumber((_dafny.Seq.of(_682_v39)).length)), new _dafny.CodePoint('m'.codePointAt(0)))).length);
        _nw95[(new BigNumber(10)).toNumber()] = new BigNumber((_677_v34).length);
        _nw95[(new BigNumber(11)).toNumber()] = new BigNumber(850);
        _nw95[(new BigNumber(12)).toNumber()] = new BigNumber(216);
        _nw95[(new BigNumber(13)).toNumber()] = new BigNumber((_683_v40).cardinality());
        _nw95[(new BigNumber(14)).toNumber()] = new BigNumber(-913);
        _nw95[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_this).f6);
        _nw95[(new BigNumber(16)).toNumber()] = (_this).f6;
        _nw95[(new BigNumber(17)).toNumber()] = (_this).f6;
        _nw95[(new BigNumber(18)).toNumber()] = (_685_v42).dtor_cf3;
        _nw95[(new BigNumber(19)).toNumber()] = (_this).f6;
        _nw95[(new BigNumber(20)).toNumber()] = new BigNumber(227);
        _nw95[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus((_this).f6);
        _nw95[(new BigNumber(22)).toNumber()] = (((_684_v41).contains((_this).f6)) ? ((_684_v41).get((_this).f6)) : (new BigNumber(((_686_v43)[_module.__default.safeIndex((_this).f6, new BigNumber((_686_v43).length))]).length)));
        _nw95[(new BigNumber(23)).toNumber()] = (_this).f6;
        _687_v44 = _nw95;
        let _688_v45;
        _688_v45 = _module.D2.create_DC6(new BigNumber(181), _687_v44);
        let _689_v46;
        let _nw96 = Array((new BigNumber(27)).toNumber());
        _nw96[(_dafny.ZERO).toNumber()] = ((_this).f6).plus((_this).f6);
        _nw96[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-882));
        _nw96[(new BigNumber(2)).toNumber()] = new BigNumber(-131);
        _nw96[(new BigNumber(3)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(4)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(5)).toNumber()] = ((_this).f6).plus(new BigNumber(-829));
        _nw96[(new BigNumber(6)).toNumber()] = new BigNumber(21);
        _nw96[(new BigNumber(7)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_677_v34).length), new BigNumber((_678_v35).length));
        _nw96[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_679_v36).length));
        _nw96[(new BigNumber(10)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(11)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(12)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_681_v38).dtor_cf30);
        _nw96[(new BigNumber(14)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(15)).toNumber()] = new BigNumber((_679_v36).length);
        _nw96[(new BigNumber(16)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(17)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(18)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(19)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(20)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(21)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(22)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(23)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(24)).toNumber()] = new BigNumber((_dafny.Seq.update(_module.__default.fm30(p1, _682_v39, globalState), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), ((_690_v39) => function (_691_i1) {
          return _690_v39;
        })(_682_v39))).length), new BigNumber((_module.__default.fm30(p1, _682_v39, globalState)).length)), _682_v39)).length);
        _nw96[(new BigNumber(25)).toNumber()] = (_this).f6;
        _nw96[(new BigNumber(26)).toNumber()] = (_688_v45).dtor_cf13;
        _689_v46 = _nw96;
        let _692_v47;
        _692_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Set.fromElements(_module.__default.fm2(new BigNumber((_679_v36).length), globalState), (_this).f6, (_this).f6));
        let _693_v48;
        let _nw97 = Array((new BigNumber(23)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = _module.D2.create_DC6((_this).f6, _687_v44);
        _nw97[(_dafny.ONE).toNumber()] = _688_v45;
        _nw97[(new BigNumber(2)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(3)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(4)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(5)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(6)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(7)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(8)).toNumber()] = _module.D2.create_DC6((_this).f6, _687_v44);
        _nw97[(new BigNumber(9)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(10)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(11)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(12)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(13)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(14)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(15)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(16)).toNumber()] = _module.D2.create_DC6(new BigNumber(((((_692_v47).contains((_this).f6)) ? ((_692_v47).get((_this).f6)) : (_678_v35))).length), _689_v46);
        _nw97[(new BigNumber(17)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(18)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(19)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(20)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(21)).toNumber()] = _688_v45;
        _nw97[(new BigNumber(22)).toNumber()] = _688_v45;
        _693_v48 = _nw97;
        let _index99 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_693_v48).length));
        (_693_v48)[_index99] = _688_v45;
        let _index100 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_693_v48).length));
        let _rhs110 = ((p0) ? ((_dafny.ZERO).minus((_this).f6)) : ((_this).f6));
        let _rhs111 = _dafny.Seq.UnicodeFromString("i");
        let _rhs112 = (_this).f6;
        let _rhs113 = _687_v44;
        let _rhs114 = _688_v45;
        let _lhs90 = globalState;
        let _lhs91 = globalState;
        let _lhs92 = globalState;
        let _lhs93 = _693_v48;
        let _lhs94 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_693_v48).length));
        _lhs90.f5 = _rhs110;
        _lhs91.f2 = _rhs111;
        _lhs92.f5 = _rhs112;
        _689_v46 = _rhs113;
        _lhs93[_lhs94] = _rhs114;
      }
      let _694_v49;
      let _nw98 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
      _694_v49 = _nw98;
      let _695_v50;
      _695_v50 = _dafny.Map.Empty.slice().updateUnsafe(_694_v49,p1);
      let _696_v51;
      _696_v51 = _module.D10.create_DC28(_695_v50);
      _695_v50 = (_696_v51).dtor_cf52;
      let _697_v52;
      _697_v52 = _dafny.Set.fromElements((_this).f6, (_this).f6);
      let _index101 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_694_v49).length));
      (_694_v49)[_index101] = new BigNumber((_697_v52).length);
      let _index102 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_694_v49).length));
      (_694_v49)[_index102] = (_this).f6;
      let _698_v53;
      _698_v53 = _dafny.Seq.of(new BigNumber(409));
      let _699_v54;
      _699_v54 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(665), new BigNumber(647))).cardinality()), (_this).f6);
      let _700_v55;
      _700_v55 = _dafny.MultiSet.fromElements(_699_v54);
      let _701_v56;
      let _nw99 = new _module.C2();
      _nw99.__ctor((_dafny.ZERO).minus((_698_v53)[_module.__default.safeIndex((_this).f6, new BigNumber((_698_v53).length))]), _700_v55, _this.f9, (_this).f10);
      _701_v56 = _nw99;
      let _702_v57;
      _702_v57 = _dafny.Map.Empty.slice().updateUnsafe(_701_v56,p0);
      let _703_v58;
      _703_v58 = _dafny.Map.Empty.slice().updateUnsafe(_702_v57,_698_v53);
      let _704_v59;
      _704_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_703_v58);
      _704_v59 = (_704_v59).update((_694_v49)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_694_v49).length))], _703_v58);
      let _705_i2;
      _705_i2 = _dafny.ZERO;
      L2: {
        while (p0) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_705_i2)) {
              break L2;
            }
            _705_i2 = (_705_i2).plus(_dafny.ONE);
            r2 = p1;
            r0 = new BigNumber(839);
            let _arr15 = _this.f9;
            let _index103 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_this.f9).length));
            _arr15[_index103] = p1;
            let _arr16 = _this.f9;
            let _index104 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_this.f9).length));
            _arr16[_index104] = (((p0) || (p0)) ? (p1) : (true));
            let _706_v60;
            let _nw100 = Array((new BigNumber(6)).toNumber());
            _nw100[(_dafny.ZERO).toNumber()] = _694_v49;
            _nw100[(_dafny.ONE).toNumber()] = _694_v49;
            _nw100[(new BigNumber(2)).toNumber()] = _694_v49;
            _nw100[(new BigNumber(3)).toNumber()] = _694_v49;
            _nw100[(new BigNumber(4)).toNumber()] = _694_v49;
            _nw100[(new BigNumber(5)).toNumber()] = _694_v49;
            _706_v60 = _nw100;
            let _index105 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_706_v60).length));
            (_706_v60)[_index105] = _694_v49;
            let _707_v61;
            _707_v61 = _dafny.Seq.of(_698_v53, _dafny.Seq.Create(_module.__default.abs(new BigNumber(640)), ((_708_v56) => function (_709_i3) {
              return (_708_v56).f6;
            })(_701_v56)), _698_v53);
            let _710_v62;
            let _init13 = ((_711_p0) => function (_712_i4) {
              return _711_p0;
            })(p0);
            let _nw101 = Array((new BigNumber(9)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw101.length); _i0_13++) {
              _nw101[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _710_v62 = _nw101;
            let _713_v63;
            _713_v63 = _dafny.Map.Empty.slice().updateUnsafe((_707_v61)[_module.__default.safeIndex((_this).f6, new BigNumber((_707_v61).length))],_710_v62);
            let _714_v64;
            _714_v64 = _dafny.Set.fromElements((_this.f9)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_this.f9).length))], true);
            let _index106 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_694_v49).length));
            let _index107 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_706_v60).length));
            let _rhs115 = (_module.__default.safeModuloInt((_this).f6, new BigNumber((_module.__default.fm40((_this).f6, _714_v64, globalState)).length))).multipliedBy((_dafny.ZERO).minus((_701_v56).f6));
            let _rhs116 = (_694_v49)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_694_v49).length))];
            let _rhs117 = _694_v49;
            let _rhs118 = ((_713_v63).Merge(_713_v63)).Merge(_713_v63);
            let _rhs119 = p1;
            let _lhs95 = _694_v49;
            let _lhs96 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_694_v49).length));
            let _lhs97 = _706_v60;
            let _lhs98 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_706_v60).length));
            let _lhs99 = globalState;
            r0 = _rhs115;
            _lhs95[_lhs96] = _rhs116;
            _lhs97[_lhs98] = _rhs117;
            _713_v63 = _rhs118;
            _lhs99.f0 = _rhs119;
          }
        }
      }
      let _rhs120 = p0;
      let _rhs121 = ((p0) ? (_694_v49) : (_694_v49));
      r2 = _rhs120;
      _694_v49 = _rhs121;
      r0 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("wwrq")).length), (_694_v49)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((_694_v49).length))]);
      let _715_v65;
      _715_v65 = new _dafny.CodePoint('y'.codePointAt(0));
      r1 = _715_v65;
      let _716_v66;
      _716_v66 = _dafny.Set.fromElements(p0, p1, true);
      r2 = (_716_v66).IsProperSubsetOf(_dafny.Set.fromElements(p0));
      return [r0, r1, r2];
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f9 = [];
      this._f16 = _dafny.MultiSet.Empty;
      this._f6 = _dafny.ZERO;
      this._f10 = _dafny.Map.Empty;
      this._f17 = _dafny.ZERO;
      this._f18 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0, _module.T2];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    set f9(value) {
      let _this = this;
      _this._f9 = value;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f17, f18, f9, f10, f6, f16) {
      let _this = this;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f6 = f6;
      (_this)._f16 = f16;
      return;
    }
    fm13(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source9 = _module.D1.create_DC3(_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0))));
      if (_source9.is_DC4) {
        let _717___mcc_h0 = (_source9).cf7;
        let _718___mcc_h1 = (_source9).cf8;
        let _719___mcc_h2 = (_source9).cf9;
        let _720___mcc_h3 = (_source9).cf10;
        let _721___mcc_h4 = (_source9).cf11;
        let _722_cf11 = _721___mcc_h4;
        let _723_cf10 = _720___mcc_h3;
        let _724_cf9 = _719___mcc_h2;
        let _725_cf8 = _718___mcc_h1;
        let _726_cf7 = _717___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(false,_724_cf9)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_723_cf10,_724_cf9));
      } else {
        let _727___mcc_h5 = (_source9).cf6;
        let _728_cf6 = _727___mcc_h5;
        return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("kg"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("xmpjtp")));
      }
    };
    fm5(p0, globalState) {
      let _this = this;
      let _source10 = _module.D0.create_DC0(_dafny.Seq.of(false, true), (_this).f6);
      if (_source10.is_DC0) {
        let _729___mcc_h0 = (_source10).cf0;
        let _730___mcc_h1 = (_source10).cf1;
        let _731_cf1 = _730___mcc_h1;
        let _732_cf0 = _729___mcc_h0;
        return _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))));
      } else if (_source10.is_DC1) {
        let _733___mcc_h2 = (_source10).cf2;
        let _734___mcc_h3 = (_source10).cf3;
        let _735___mcc_h4 = (_source10).cf4;
        let _736_cf4 = _735___mcc_h4;
        let _737_cf3 = _734___mcc_h3;
        let _738_cf2 = _733___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(_736_cf4), _dafny.MultiSet.fromElements(_736_cf4)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(448)), ((_739_cf4) => function (_740_i0) {
          return _dafny.MultiSet.fromElements(_739_cf4);
        })(_736_cf4)));
      } else {
        let _741___mcc_h5 = (_source10).cf5;
        let _742_cf5 = _741___mcc_h5;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), function (_743_i1) {
          return _dafny.MultiSet.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(613)), function (_744_i2) {
          return _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(237)), function (_745_i3) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }));
        }));
      }
    };
    fm6(p0, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe((_this).f17,_dafny.Seq.of((_this).f18))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(669)), function (_746_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length),_dafny.Seq.of((_this).f18))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, true)).length),_dafny.Seq.of((_this).f18))));
    };
    fm21(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_this).f18, ((_this).f17).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).length)));
    };
    fm22(p0, p1, p2, globalState) {
      let _this = this;
      let _source11 = ((!(false)) ? (_module.D0.create_DC1(false, (_this).f6, new _dafny.CodePoint('y'.codePointAt(0)))) : (_module.D0.create_DC1(true, (_this).f18, new _dafny.CodePoint('n'.codePointAt(0)))));
      if (_source11.is_DC0) {
        let _747___mcc_h0 = (_source11).cf0;
        let _748___mcc_h1 = (_source11).cf1;
        let _749_cf1 = _748___mcc_h1;
        let _750_cf0 = _747___mcc_h0;
        return !(true);
      } else if (_source11.is_DC1) {
        let _751___mcc_h2 = (_source11).cf2;
        let _752___mcc_h3 = (_source11).cf3;
        let _753___mcc_h4 = (_source11).cf4;
        let _754_cf4 = _753___mcc_h4;
        let _755_cf3 = _752___mcc_h3;
        let _756_cf2 = _751___mcc_h2;
        return _756_cf2;
      } else {
        let _757___mcc_h5 = (_source11).cf5;
        let _758_cf5 = _757___mcc_h5;
        return (_dafny.Set.fromElements(false)).IsProperSubsetOf(_dafny.Set.fromElements(true));
      }
    };
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _759_v0;
      _759_v0 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(18),p0);
      let _hi5 = (p2).minus(p2);
      for (let _760_i0 = new BigNumber(((_759_v0).Merge(_759_v0)).length); _760_i0.isLessThan(_hi5); _760_i0 = _760_i0.plus(_dafny.ONE)) {
        (globalState).f5 = p2;
        let _761_v1;
        _761_v1 = _module.D0.create_DC0(p1, (_this).f18);
        let _762_v2;
        _762_v2 = _dafny.Seq.UnicodeFromString("jnbsg");
        let _763_v3;
        _763_v3 = _module.D1.create_DC4(_761_v1, (_this).f6, ((p0) ? (_762_v2) : (_762_v2)), p0, (_this).f18);
        _763_v3 = _module.D1.create_DC4(_761_v1, new BigNumber(403), _762_v2, p0, (_this).f18);
        _762_v2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_762_v2, _762_v2), _dafny.Seq.Concat(_762_v2, _762_v2));
        let _764_v4;
        let _init14 = function (_765_i1) {
          return _module.D1.create_DC3(_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0))));
        };
        let _nw102 = Array((new BigNumber(21)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw102.length); _i0_14++) {
          _nw102[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _764_v4 = _nw102;
        let _766_v5;
        let _nw103 = Array((new BigNumber(26)).toNumber()).fill(_dafny.MultiSet.Empty);
        _766_v5 = _nw103;
        let _767_v6;
        _767_v6 = _dafny.Map.Empty.slice().updateUnsafe(_764_v4,_766_v5);
        let _768_v7;
        _768_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p1)[_module.__default.safeIndex(new BigNumber(-465), new BigNumber((p1).length))]);
        let _769_v8;
        _769_v8 = _dafny.MultiSet.fromElements(p0, (((_768_v7).contains(!(p0))) ? ((_768_v7).get(!(p0))) : (p0)));
        let _770_v9;
        _770_v9 = _module.D2.create_DC7(_dafny.Seq.Concat(p1, _module.__default.fm4(_760_i0, (_dafny.ZERO).minus(_760_i0), new BigNumber((_dafny.Seq.UnicodeFromString("m")).length), globalState)), (((_767_v6).contains(_764_v4)) ? ((_767_v6).get(_764_v4)) : (_766_v5)), _module.__default.fm2((_this).f17, globalState), ((_this).f18).plus(new BigNumber((_769_v8).cardinality())));
        _770_v9 = _770_v9;
      }
      (globalState).f5 = (_this).f18;
      let _771_i2;
      _771_i2 = _dafny.ZERO;
      L3: {
        while (p0) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_771_i2)) {
              break L3;
            }
            _771_i2 = (_771_i2).plus(_dafny.ONE);
            let _772_v10;
            let _nw104 = new _module.C0();
            _nw104.__ctor(_this.f9, _module.__default.fm3(globalState));
            _772_v10 = _nw104;
            (globalState).f5 = (_this).f18;
            let _773_v11;
            _773_v11 = _module.D0.create_DC0(_dafny.Seq.of(p0), (_this).f6);
            (globalState).f5 = _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm25(p2, _772_v10.f20, globalState)).length), (_773_v11).dtor_cf1);
            let _774_v12;
            let _init15 = function (_775_i3) {
              return _module.__default.safeDivisionInt(_775_i3, (_this).f18);
            };
            let _nw105 = Array((new BigNumber(9)).toNumber());
            for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw105.length); _i0_15++) {
              _nw105[_i0_15] = _init15(new BigNumber(_i0_15));
            }
            _774_v12 = _nw105;
            let _index108 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_774_v12).length));
            (_774_v12)[_index108] = (_this).f6;
            let _776_v13;
            _776_v13 = _dafny.Seq.of(new BigNumber(8), (_this).f17);
            let _777_v14;
            _777_v14 = _module.D2.create_DC8(_772_v10.f20, p0);
            let _778_v15;
            _778_v15 = _dafny.MultiSet.fromElements((_this).f6);
            let _779_v16;
            _779_v16 = new _dafny.CodePoint('l'.codePointAt(0));
            let _780_v17;
            _780_v17 = _dafny.Map.Empty.slice().updateUnsafe(_779_v16,(_this).f6);
            let _index109 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_774_v12).length));
            let _rhs122 = _772_v10.f20;
            let _rhs123 = (_this).f6;
            let _rhs124 = _dafny.Seq.IsPrefixOf(_776_v13, _776_v13);
            let _rhs125 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_772_v10.f20,_777_v14)).length);
            let _rhs126 = _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm4((_dafny.ZERO).minus(new BigNumber(((_778_v15).update((_this).f6, _module.__default.abs((_dafny.ZERO).minus((_this).f18)))).cardinality())), new BigNumber(621), (_this).f18, globalState)).length), _module.__default.safeModuloInt((((_780_v17).contains(new _dafny.CodePoint('g'.codePointAt(0)))) ? ((_780_v17).get(new _dafny.CodePoint('g'.codePointAt(0)))) : ((_this).f6)), (_this).f6));
            let _lhs100 = _774_v12;
            let _lhs101 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_774_v12).length));
            let _lhs102 = globalState;
            let _lhs103 = globalState;
            let _lhs104 = globalState;
            r1 = _rhs122;
            _lhs100[_lhs101] = _rhs123;
            _lhs102.f0 = _rhs124;
            _lhs103.f5 = _rhs125;
            _lhs104.f5 = _rhs126;
          }
        }
      }
      if (p0) {
        let _781_v18;
        _781_v18 = _module.D0.create_DC0(p1, p2);
        let _782_v19;
        _782_v19 = _module.D0.create_DC2(_781_v18);
        let _783_v20;
        _783_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_782_v19);
        let _784_v21;
        _784_v21 = _dafny.Set.fromElements(_783_v20);
        (globalState).f0 = !(((function () {
          let _coll19 = new _dafny.Set();
          for (const _compr_19 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-885)), ((_785_v20) => function (_786_i4) {
            return _785_v20;
          })(_783_v20))).Elements) {
            let _787_v22 = _compr_19;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-885)), ((_788_v20) => function (_786_i4) {
              return _788_v20;
            })(_783_v20)), _787_v22)) {
              _coll19.add(_787_v22);
            }
          }
          return _coll19;
        }()).Union(_784_v21)).IsProperSubsetOf(_784_v21));
        let _789_v23;
        _789_v23 = new _dafny.CodePoint('f'.codePointAt(0));
        _789_v23 = _789_v23;
        let _790_v24;
        let _nw106 = new _module.C0();
        _nw106.__ctor(_this.f9, p0);
        _790_v24 = _nw106;
        _790_v24 = _790_v24;
        let _791_v25;
        let _nw107 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _791_v25 = _nw107;
        let _index110 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_791_v25).length));
        (_791_v25)[_index110] = _789_v23;
        let _index111 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_791_v25).length));
        (_791_v25)[_index111] = _789_v23;
        let _792_v26;
        _792_v26 = _dafny.Seq.UnicodeFromString("wjkiaxif");
        let _793_v27;
        _793_v27 = _dafny.MultiSet.fromElements(new BigNumber((_792_v26).length));
        (globalState).f0 = ((_this).fm21((_dafny.ZERO).minus(new BigNumber((_759_v0).length)), globalState)).isLessThan(_module.__default.safeModuloInt((_this).f17, new BigNumber((_793_v27).cardinality())));
      } else {
        if (!((_this).f18).isEqualTo(p2)) {
          let _794_v28;
          let _init16 = function (_795_i5) {
            return (_795_i5).minus((_this).f6);
          };
          let _nw108 = Array((new BigNumber(17)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw108.length); _i0_16++) {
            _nw108[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _794_v28 = _nw108;
          let _796_v29;
          _796_v29 = _dafny.Map.Empty.slice().updateUnsafe(_794_v28,new BigNumber((_dafny.Seq.Concat(p1, p1)).length));
          _796_v29 = _796_v29;
          let _797_v30;
          let _nw109 = new _module.C0();
          _nw109.__ctor(_this.f9, p0);
          _797_v30 = _nw109;
          let _798_v31;
          let _nw110 = new _module.C0();
          _nw110.__ctor(_this.f9, p0);
          _798_v31 = _nw110;
          let _799_v32;
          _799_v32 = _dafny.Seq.of(_798_v31.f20);
          let _rhs127 = p1;
          let _rhs128 = true;
          let _lhs105 = _798_v31;
          _799_v32 = _rhs127;
          _lhs105.f20 = _rhs128;
          let _800_v33;
          let _nw111 = new _module.C0();
          _nw111.__ctor(_this.f9, p0);
          _800_v33 = _nw111;
        } else {
          (globalState).f0 = (_dafny.Set.fromElements(p0)).IsSubsetOf((_dafny.Set.fromElements(p0, p0, false)).Difference(_dafny.Set.fromElements(p0, p0)));
          (globalState).f5 = (_this).f17;
          (globalState).f0 = p0;
          let _801_v34;
          let _nw112 = new _module.C0();
          _nw112.__ctor(((p0) ? (_this.f9) : (_this.f9)), p0);
          _801_v34 = _nw112;
          let _802_v35;
          let _nw113 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _802_v35 = _nw113;
          let _803_v36;
          _803_v36 = _dafny.Seq.UnicodeFromString("koabsqd");
          let _804_v37;
          _804_v37 = new _dafny.CodePoint('v'.codePointAt(0));
          let _index112 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_802_v35).length));
          (_802_v35)[_index112] = (new BigNumber((_803_v36).length)).multipliedBy(new BigNumber((_module.__default.fm26((_this).f6, _804_v37, _801_v34.f20, _804_v37, globalState)).length));
          let _index113 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_802_v35).length));
          (_802_v35)[_index113] = (_this).f6;
        }
        let _805_v38;
        _805_v38 = _dafny.Seq.of((_this).f17);
        let _806_v39;
        let _nw114 = new _module.C0();
        _nw114.__ctor(_this.f9, ((_805_v38)[_module.__default.safeIndex((_this).f18, new BigNumber((_805_v38).length))]).isLessThan((_this).f17));
        _806_v39 = _nw114;
        let _807_v40;
        _807_v40 = _dafny.Seq.UnicodeFromString("snhg");
        let _808_v41;
        _808_v41 = _dafny.Set.fromElements(_806_v39.f20);
        let _809_v42;
        _809_v42 = new _dafny.CodePoint('b'.codePointAt(0));
        let _810_v43;
        _810_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_808_v41).length))).cardinality()),_809_v42);
        let _811_v45;
        _811_v45 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(18)), ((_812_v42) => function (_813_i6) {
          return _812_v42;
        })(_809_v42)));
        let _814_v46;
        _814_v46 = _dafny.Map.Empty.slice().updateUnsafe(_807_v40,new BigNumber(-128));
        let _815_v48;
        let _nw115 = Array((new BigNumber(13)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = ((p0) ? (_dafny.Map.Empty.slice().updateUnsafe(_807_v40,new BigNumber((_dafny.Seq.of((_807_v40)[_module.__default.safeIndex(p2, new BigNumber((_807_v40).length))], (((_810_v43).contains((_this).f17)) ? ((_810_v43).get((_this).f17)) : ((_807_v40)[_module.__default.safeIndex(p2, new BigNumber((_807_v40).length))])))).length))) : (function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of (_811_v45).Elements) {
            let _816_v44 = _compr_20;
            if ((_811_v45).contains(_816_v44)) {
              _coll20.push([_816_v44,new BigNumber(313)]);
            }
          }
          return _coll20;
        }()));
        _nw115[(_dafny.ONE).toNumber()] = _814_v46;
        _nw115[(new BigNumber(2)).toNumber()] = _814_v46;
        _nw115[(new BigNumber(3)).toNumber()] = _814_v46;
        _nw115[(new BigNumber(4)).toNumber()] = _814_v46;
        _nw115[(new BigNumber(5)).toNumber()] = _module.__default.fm27(globalState);
        _nw115[(new BigNumber(6)).toNumber()] = _814_v46;
        _nw115[(new BigNumber(7)).toNumber()] = _814_v46;
        _nw115[(new BigNumber(8)).toNumber()] = _814_v46;
        _nw115[(new BigNumber(9)).toNumber()] = _814_v46;
        _nw115[(new BigNumber(10)).toNumber()] = (_814_v46).Merge(_dafny.Map.Empty.slice().updateUnsafe(_807_v40,(_dafny.ZERO).minus(new BigNumber(-528))));
        _nw115[(new BigNumber(11)).toNumber()] = _814_v46;
        _nw115[(new BigNumber(12)).toNumber()] = (function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of ((_dafny.MultiSet.fromElements(_807_v40)).update(_807_v40, _module.__default.abs((_this).f17))).Elements) {
            let _817_v47 = _compr_21;
            if (((_dafny.MultiSet.fromElements(_807_v40)).update(_807_v40, _module.__default.abs((_this).f17))).contains(_817_v47)) {
              _coll21.push([_817_v47,(_this).f17]);
            }
          }
          return _coll21;
        }()).Merge(_814_v46);
        _815_v48 = _nw115;
        let _index114 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_815_v48).length));
        (_815_v48)[_index114] = _814_v46;
        let _index115 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_815_v48).length));
        (_815_v48)[_index115] = _814_v46;
        let _818_v49;
        _818_v49 = _dafny.Set.fromElements((_this).f17, (_dafny.ZERO).minus(p2), (_dafny.ZERO).minus((_this).f6));
        let _819_v50;
        _819_v50 = _dafny.MultiSet.fromElements((_this).f18, p2, new BigNumber((_818_v49).length));
        let _rhs129 = ((!((p0) === (p0))) ? (_819_v50) : (_819_v50));
        let _rhs130 = _806_v39.f20;
        let _lhs106 = _806_v39;
        _819_v50 = _rhs129;
        _lhs106.f20 = _rhs130;
        let _arr17 = _806_v39.f19;
        let _index116 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_806_v39.f19).length));
        _arr17[_index116] = p0;
        let _arr18 = _806_v39.f19;
        let _index117 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_806_v39.f19).length));
        _arr18[_index117] = (_819_v50).IsSubsetOf((_819_v50).Union(_819_v50));
      }
      let _820_i7;
      _820_i7 = _dafny.ZERO;
      L4: {
        while (p0) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_820_i7)) {
              break L4;
            }
            _820_i7 = (_820_i7).plus(_dafny.ONE);
            let _821_v51;
            _821_v51 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f9,(_this).f18)).length));
            _759_v0 = (_759_v0).update(_module.__default.safeModuloInt(p2, new BigNumber((_821_v51).cardinality())), true);
            let _arr19 = _this.f9;
            let _index118 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
            _arr19[_index118] = p0;
            let _822_v52;
            let _init17 = function (_823_i8) {
              return (_823_i8).minus((_this).f17);
            };
            let _nw116 = Array((new BigNumber(8)).toNumber());
            for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw116.length); _i0_17++) {
              _nw116[_i0_17] = _init17(new BigNumber(_i0_17));
            }
            _822_v52 = _nw116;
            let _index119 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_822_v52).length));
            (_822_v52)[_index119] = p2;
            let _824_v53;
            _824_v53 = _dafny.Seq.UnicodeFromString("wmae");
            let _arr20 = _this.f9;
            let _index120 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
            let _index121 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_822_v52).length));
            let _rhs131 = p0;
            let _rhs132 = (_this).f6;
            let _rhs133 = p0;
            let _rhs134 = (new BigNumber((_824_v53).length)).plus((_this).f17);
            let _lhs107 = globalState;
            let _lhs108 = globalState;
            let _lhs109 = _this.f9;
            let _lhs110 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
            let _lhs111 = _822_v52;
            let _lhs112 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_822_v52).length));
            _lhs107.f0 = _rhs131;
            _lhs108.f5 = _rhs132;
            _lhs109[_lhs110] = _rhs133;
            _lhs111[_lhs112] = _rhs134;
            (globalState).f5 = p2;
            let _825_v54;
            _825_v54 = _dafny.MultiSet.fromElements(_822_v52);
            let _rhs135 = (((_this.f9)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length))]) ? (((p0) ? (_822_v52) : (_822_v52))) : (_822_v52));
            let _rhs136 = _825_v54;
            let _rhs137 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), ((_826_v53) => function (_827_i9) {
              return (_826_v53)[_module.__default.safeIndex((_this).f17, new BigNumber((_826_v53).length))];
            })(_824_v53));
            let _rhs138 = _module.__default.fm3(globalState);
            let _lhs113 = globalState;
            let _lhs114 = globalState;
            _822_v52 = _rhs135;
            _825_v54 = _rhs136;
            _lhs113.f2 = _rhs137;
            _lhs114.f0 = _rhs138;
          }
        }
      }
      let _828_v55;
      let _nw117 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
      _828_v55 = _nw117;
      let _index122 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_828_v55).length));
      (_828_v55)[_index122] = _dafny.Set.fromElements((_this).f6, (_this).f17);
      let _829_v56;
      _829_v56 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),(_this).f17);
      let _830_v57;
      _830_v57 = _module.D2.create_DC5(_829_v56);
      let _831_v58;
      let _nw118 = new _module.C3();
      _nw118.__ctor((_830_v57).dtor_cf12, _this.f9, (_this).f10, new BigNumber(13));
      _831_v58 = _nw118;
      let _832_v59;
      _832_v59 = _dafny.Seq.of(_831_v58);
      let _833_v60;
      _833_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _834_v61;
      _834_v61 = _dafny.Seq.of(_833_v60, _833_v60);
      let _835_v62;
      _835_v62 = _dafny.Set.fromElements(p2, new BigNumber((_dafny.MultiSet.FromArray(_832_v59)).cardinality()), new BigNumber((_834_v61).length));
      let _836_v63;
      _836_v63 = _module.D11.create_DC30(_835_v62);
      let _837_v64;
      _837_v64 = _dafny.Seq.of((_835_v62).Difference((_836_v63).dtor_cf57));
      let _index123 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_828_v55).length));
      let _rhs139 = ((_this).f18).minus((_this).f18);
      let _rhs140 = _this.f9;
      let _rhs141 = (_837_v64)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,false)).length), new BigNumber((_837_v64).length))];
      let _lhs115 = globalState;
      let _lhs116 = _this;
      let _lhs117 = _828_v55;
      let _lhs118 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_828_v55).length));
      _lhs115.f5 = _rhs139;
      _lhs116.f9 = _rhs140;
      _lhs117[_lhs118] = _rhs141;
      r0 = _module.__default.fm29(((p0) ? ((_828_v55)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_828_v55).length))]) : ((_828_v55)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_828_v55).length))])), globalState);
      r1 = p0;
      return [r0, r1];
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = [];
      let r3 = _dafny.ZERO;
      let _838_v0;
      _838_v0 = new _dafny.CodePoint('v'.codePointAt(0));
      let _839_v1;
      _839_v1 = _dafny.MultiSet.fromElements(false);
      let _840_v2;
      _840_v2 = _dafny.Seq.of((_this).f6);
      let _841_v3;
      _841_v3 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f17, new BigNumber((_839_v1).cardinality())), _840_v2);
      let _842_v5;
      _842_v5 = _dafny.Map.Empty.slice().updateUnsafe(_838_v0,new BigNumber((function () {
        let _coll22 = new _dafny.Set();
        for (const _compr_22 of (_841_v3).Elements) {
          let _843_v4 = _compr_22;
          if ((_841_v3).contains(_843_v4)) {
            _coll22.add(_843_v4);
          }
        }
        return _coll22;
      }()).length));
      let _hi6 = (((_842_v5).contains(_838_v0)) ? ((_842_v5).get(_838_v0)) : (new BigNumber((function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of _dafny.IntegerRange(new BigNumber(35), new BigNumber(958))) {
          let _844_v6 = _compr_23;
          if (((new BigNumber(35)).isLessThanOrEqualTo(_844_v6)) && ((_844_v6).isLessThan(new BigNumber(958)))) {
            _coll23.push([(_844_v6).minus((_this).f18),true]);
          }
        }
        return _coll23;
      }()).length)));
      for (let _845_i0 = (_this).f6; _845_i0.isLessThan(_hi6); _845_i0 = _845_i0.plus(_dafny.ONE)) {
        let _846_v7;
        _846_v7 = _dafny.Seq.UnicodeFromString("rrvwuy");
        let _847_v8;
        _847_v8 = false;
        let _848_v9;
        _848_v9 = _dafny.Map.Empty.slice().updateUnsafe(_847_v8,p1);
        let _849_v10;
        let _nw119 = new _module.C3();
        _nw119.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_845_i0,(_this).f6), _this.f9, (_this).f10, _845_i0);
        _849_v10 = _nw119;
        let _850_v11;
        _850_v11 = _dafny.Map.Empty.slice().updateUnsafe(_849_v10,(_this).f17);
        let _851_v12;
        let _nw120 = Array((new BigNumber(10)).toNumber());
        _nw120[(_dafny.ZERO).toNumber()] = p1;
        _nw120[(_dafny.ONE).toNumber()] = new BigNumber(-173);
        _nw120[(new BigNumber(2)).toNumber()] = new BigNumber((_846_v7).length);
        _nw120[(new BigNumber(3)).toNumber()] = new BigNumber((_848_v9).length);
        _nw120[(new BigNumber(4)).toNumber()] = new BigNumber((_850_v11).length);
        _nw120[(new BigNumber(5)).toNumber()] = new BigNumber(-505);
        _nw120[(new BigNumber(6)).toNumber()] = (_this).f6;
        _nw120[(new BigNumber(7)).toNumber()] = p1;
        _nw120[(new BigNumber(8)).toNumber()] = (_this).f18;
        _nw120[(new BigNumber(9)).toNumber()] = new BigNumber(-558);
        _851_v12 = _nw120;
        let _852_v13;
        _852_v13 = _module.D5.create_DC17(_840_v2, !(_847_v8));
        let _853_v14;
        _853_v14 = _dafny.Map.Empty.slice().updateUnsafe(_851_v12,((_847_v8) ? (_847_v8) : ((_852_v13).dtor_cf37)));
        _853_v14 = (_853_v14).update(_851_v12, _847_v8);
        if (_847_v8) {
          let _854_v15;
          _854_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(199)), function (_855_i1) {
            return (_this).f18;
          }),!(_847_v8) || (_847_v8));
          _854_v15 = _854_v15;
          (globalState).f0 = _847_v8;
          let _856_v16;
          _856_v16 = _module.D4.create_DC14(_847_v8, _module.__default.fm41(globalState), _845_i0, _dafny.Map.Empty.slice().updateUnsafe(_847_v8,_847_v8));
          let _857_v17;
          _857_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_856_v16).dtor_cf28);
          _857_v17 = (_857_v17).update(new BigNumber(((_857_v17).Merge(_857_v17)).length), ((_this).f18).isLessThan(new BigNumber(-185)));
          let _858_v18;
          let _init18 = ((_859_v8) => function (_860_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe(_859_v8,_859_v8);
          })(_847_v8);
          let _nw121 = Array((new BigNumber(28)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw121.length); _i0_18++) {
            _nw121[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _858_v18 = _nw121;
          let _861_v19;
          _861_v19 = _dafny.Set.fromElements(_module.D9.create_DC26(_858_v18));
          let _862_v20;
          _862_v20 = _dafny.Map.Empty.slice().updateUnsafe(_847_v8,_861_v19);
          let _863_v21;
          let _nw122 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
          _863_v21 = _nw122;
          let _864_v22;
          _864_v22 = _module.D10.create_DC28(_853_v14);
          let _rhs142 = ((_862_v20).Merge(_862_v20)).Merge(_862_v20);
          let _rhs143 = _863_v21;
          let _rhs144 = _864_v22;
          let _rhs145 = _847_v8;
          let _lhs119 = globalState;
          _862_v20 = _rhs142;
          _863_v21 = _rhs143;
          _864_v22 = _rhs144;
          _lhs119.f0 = _rhs145;
          _838_v0 = _838_v0;
        } else {
          let _arr21 = _this.f9;
          let _index124 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
          _arr21[_index124] = _847_v8;
          let _865_v23;
          _865_v23 = _dafny.Map.Empty.slice().updateUnsafe(!(_847_v8),_847_v8);
          let _866_v24;
          _866_v24 = _module.D4.create_DC14(_847_v8, _865_v23, p1, _865_v23);
          let _arr22 = _this.f9;
          let _index125 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
          _arr22[_index125] = (!(_847_v8)) && ((_866_v24).dtor_cf28);
          let _867_v25;
          let _init19 = ((_868_v10) => function (_869_i3) {
            return ((_868_v10).f21).Merge((_868_v10).f21);
          })(_849_v10);
          let _nw123 = Array((new BigNumber(7)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw123.length); _i0_19++) {
            _nw123[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _867_v25 = _nw123;
          _867_v25 = _867_v25;
          let _870_v26;
          _870_v26 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_838_v0));
          let _871_v27;
          _871_v27 = _module.D4.create_DC12(_870_v26);
          let _872_v28;
          let _nw124 = Array((new BigNumber(21)).toNumber()).fill(_dafny.MultiSet.Empty);
          _872_v28 = _nw124;
          let _index126 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_872_v28).length));
          (_872_v28)[_index126] = (_dafny.MultiSet.fromElements(_840_v2)).update(_dafny.Seq.of(new BigNumber(766)), _module.__default.abs(p1));
          let _873_v29;
          let _nw125 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
          _873_v29 = _nw125;
          let _index127 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_872_v28).length));
          let _rhs146 = _871_v27;
          let _rhs147 = _841_v3;
          let _rhs148 = _873_v29;
          let _lhs120 = _872_v28;
          let _lhs121 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_872_v28).length));
          _871_v27 = _rhs146;
          _lhs120[_lhs121] = _rhs147;
          _873_v29 = _rhs148;
          r3 = new BigNumber(((((_847_v8) ? (_865_v23) : (_865_v23))).Merge(_865_v23)).length);
          let _arr23 = _this.f9;
          let _index128 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
          let _arr24 = _this.f9;
          let _index129 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
          let _arr25 = _this.f9;
          let _index130 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
          let _rhs149 = _847_v8;
          let _rhs150 = !((_this.f9)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length))]);
          let _rhs151 = (_module.D0.create_DC1(_847_v8, p1, _838_v0)).dtor_cf4;
          let _rhs152 = _847_v8;
          let _lhs122 = _this.f9;
          let _lhs123 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
          let _lhs124 = _this.f9;
          let _lhs125 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
          let _lhs126 = _this.f9;
          let _lhs127 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_this.f9).length));
          _lhs122[_lhs123] = _rhs149;
          _lhs124[_lhs125] = _rhs150;
          _838_v0 = _rhs151;
          _lhs126[_lhs127] = _rhs152;
        }
        let _874_v30;
        let _nw126 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
        _874_v30 = _nw126;
        let _index131 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_874_v30).length));
        (_874_v30)[_index131] = _840_v2;
        let _875_v31;
        _875_v31 = _dafny.Map.Empty.slice().updateUnsafe(false,_847_v8);
        let _876_v32;
        _876_v32 = _module.D4.create_DC14(_847_v8, _875_v31, (_this).f17, _875_v31);
        let _877_v33;
        _877_v33 = _dafny.Seq.of((_this).fm22((_dafny.MultiSet.fromElements(_847_v8)).update((_876_v32).dtor_cf28, _module.__default.abs((_this).f6)), false, (_this).f6, globalState));
        let _878_v34;
        _878_v34 = _dafny.Set.fromElements(p1, new BigNumber((_877_v33).length), new BigNumber((_839_v1).cardinality()), (_this).f6);
        let _index132 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_874_v30).length));
        (_874_v30)[_index132] = _dafny.Seq.Concat(_dafny.Seq.Concat(_840_v2, _840_v2), _dafny.Seq.Concat(_840_v2, _dafny.Seq.of(new BigNumber((_878_v34).length), (_this).f17, (_this).f18)));
        r3 = ((p1).multipliedBy((_this).f18)).plus((_this).f6);
      }
      let _879_v37;
      let _init20 = ((_880_v1) => function (_881_i4) {
        return function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of (function () {
            let _coll25 = new _dafny.Set();
            for (const _compr_25 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("d"),true)).Keys.Elements) {
              let _882_v36 = _compr_25;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("d"),true)).contains(_882_v36)) {
                _coll25.add(_882_v36);
              }
            }
            return _coll25;
          }()).Elements) {
            let _883_v35 = _compr_24;
            if ((function () {
              let _coll26 = new _dafny.Set();
              for (const _compr_26 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("d"),true)).Keys.Elements) {
                let _884_v36 = _compr_26;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("d"),true)).contains(_884_v36)) {
                  _coll26.add(_884_v36);
                }
              }
              return _coll26;
            }()).contains(_883_v35)) {
              _coll24.push([_883_v35,new BigNumber((_880_v1).cardinality())]);
            }
          }
          return _coll24;
        }();
      })(_839_v1);
      let _nw127 = Array((new BigNumber(3)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw127.length); _i0_20++) {
        _nw127[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _879_v37 = _nw127;
      let _885_v38;
      _885_v38 = _dafny.Seq.UnicodeFromString("jgfwo");
      let _886_v39;
      _886_v39 = _dafny.Map.Empty.slice().updateUnsafe(_885_v38,(_this).f17);
      let _index133 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_879_v37).length));
      (_879_v37)[_index133] = _886_v39;
      let _index134 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_879_v37).length));
      (_879_v37)[_index134] = _886_v39;
      _839_v1 = _839_v1;
      let _887_v40;
      _887_v40 = true;
      let _888_v41;
      _888_v41 = _dafny.Set.fromElements(_887_v40);
      _888_v41 = ((_888_v41).Difference(_dafny.Set.fromElements(false))).Difference((_888_v41).Difference(_dafny.Set.fromElements(_887_v40)));
      let _hi7 = new BigNumber((_885_v38).length);
      for (let _889_i5 = (_this).f17; _889_i5.isLessThan(_hi7); _889_i5 = _889_i5.plus(_dafny.ONE)) {
        let _890_v42;
        _890_v42 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(115));
        let _891_v43;
        _891_v43 = _dafny.Seq.of(_this.f9, _this.f9);
        let _892_v44;
        _892_v44 = _dafny.Seq.of(_this.f9, (_891_v43)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_891_v43).length))], _this.f9, _this.f9);
        let _893_v45;
        _893_v45 = _module.D4.create_DC13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-484)), function (_894_i6) {
  return (_this).f6;
}), _887_v40, (_this).f18);
        let _895_v46;
        _895_v46 = _dafny.Map.Empty.slice().updateUnsafe(_887_v40,(_893_v45).dtor_cf26);
        let _896_v47;
        let _nw128 = new _module.C3();
        _nw128.__ctor(_890_v42, (_892_v44)[_module.__default.safeIndex(new BigNumber((_895_v46).length), new BigNumber((_892_v44).length))], (_this).f10, new BigNumber((_885_v38).length));
        _896_v47 = _nw128;
        r1 = (_this).f6;
        let _897_v48;
        let _898_v49;
        let _899_v50;
        let _out25;
        let _out26;
        let _out27;
        let _outcollector9 = (_this).m12(_887_v40, (_dafny.Set.fromElements(true)).IsSubsetOf(_888_v41), globalState);
        _out25 = _outcollector9[0];
        _out26 = _outcollector9[1];
        _out27 = _outcollector9[2];
        _897_v48 = _out25;
        _898_v49 = _out26;
        _899_v50 = _out27;
        let _900_v51;
        _900_v51 = _dafny.Set.fromElements(new BigNumber(8), (_dafny.ZERO).minus((_dafny.ZERO).minus(p1)));
        if (!(_900_v51).equals(_900_v51)) {
          let _901_v52;
          _901_v52 = _this.f9;
          let _902_v53;
          let _nw129 = new _module.C0();
          _nw129.__ctor((_901_v52), _dafny.Seq.IsPrefixOf(_module.__default.fm30(_887_v40, _838_v0, globalState), _885_v38));
          _902_v53 = _nw129;
          let _rhs153 = p1;
          let _rhs154 = new _dafny.CodePoint('o'.codePointAt(0));
          r3 = _rhs153;
          _838_v0 = _rhs154;
          let _903_v55;
          let _nw130 = new _module.C3();
          _nw130.__ctor(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f17), _902_v53.f19, ((_this).f10).Merge(function () {
            let _coll27 = new _dafny.Map();
            for (const _compr_27 of _dafny.IntegerRange(new BigNumber(978), new BigNumber(814))) {
              let _904_v54 = _compr_27;
              if (((new BigNumber(978)).isLessThanOrEqualTo(_904_v54)) && ((_904_v54).isLessThan(new BigNumber(814)))) {
                _coll27.push([(_904_v54).multipliedBy(new BigNumber(456)),_885_v38]);
              }
            }
            return _coll27;
          }()), (_dafny.ZERO).minus((_this).f6));
          _903_v55 = _nw130;
          let _905_v56;
          _905_v56 = _dafny.Map.Empty.slice().updateUnsafe(_903_v55,new BigNumber((_885_v38).length));
          _890_v42 = (_890_v42).update(new BigNumber((_905_v56).length), (_this).f6);
          let _906_v57;
          _906_v57 = _dafny.Map.Empty.slice().updateUnsafe(!(_897_v48),_899_v50);
          let _907_v58;
          _907_v58 = _dafny.Map.Empty.slice().updateUnsafe(_906_v57,(_this).f18);
          let _908_v59;
          _908_v59 = _dafny.Map.Empty.slice().updateUnsafe(p1,_906_v57);
          _907_v58 = (_907_v58).update((((_908_v59).contains((_this).f6)) ? ((_908_v59).get((_this).f6)) : (_906_v57)), ((_this).f6).minus((_840_v2)[_module.__default.safeIndex((_this).f6, new BigNumber((_840_v2).length))]));
        } else {
          let _index135 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_899_v50).length));
          (_899_v50)[_index135] = new BigNumber((_839_v1).cardinality());
          let _index136 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_899_v50).length));
          (_899_v50)[_index136] = (_this).f17;
          let _909_v60;
          _909_v60 = _dafny.Set.fromElements(_838_v0);
          let _910_v61;
          _910_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_dafny.Set.fromElements(_838_v0, _838_v0, _838_v0)).Union(_909_v60));
          _910_v61 = function () {
            let _coll28 = new _dafny.Map();
            for (const _compr_28 of (((_896_v47).f21).Merge(_module.__default.fm42(_898_v49, p0, (_dafny.ZERO).minus((_this).f6), _838_v0, globalState))).Keys.Elements) {
              let _911_v62 = _compr_28;
              if ((((_896_v47).f21).Merge(_module.__default.fm42(_898_v49, p0, (_dafny.ZERO).minus((_this).f6), _838_v0, globalState))).contains(_911_v62)) {
                _coll28.push([(_911_v62).multipliedBy(new BigNumber((function () {
                  let _coll29 = new _dafny.Map();
                  for (const _compr_29 of (_dafny.MultiSet.fromElements(_module.D0.create_DC1(_898_v49, new BigNumber(265), _838_v0), _module.D0.create_DC1(_897_v48, new BigNumber((_885_v38).length), _838_v0), _module.D0.create_DC1(_887_v40, (_this).f18, new _dafny.CodePoint('g'.codePointAt(0))), _module.D0.create_DC1(_897_v48, new BigNumber(376), _838_v0), _module.D0.create_DC1(_898_v49, (_this).f6, _838_v0))).Elements) {
                    let _912_v63 = _compr_29;
                    if ((_dafny.MultiSet.fromElements(_module.D0.create_DC1(_898_v49, new BigNumber(265), _838_v0), _module.D0.create_DC1(_897_v48, new BigNumber((_885_v38).length), _838_v0), _module.D0.create_DC1(_887_v40, (_this).f18, new _dafny.CodePoint('g'.codePointAt(0))), _module.D0.create_DC1(_897_v48, new BigNumber(376), _838_v0), _module.D0.create_DC1(_898_v49, (_this).f6, _838_v0))).contains(_912_v63)) {
                      _coll29.push([_912_v63,(_dafny.ZERO).minus((_899_v50)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_899_v50).length))])]);
                    }
                  }
                  return _coll29;
                }()).length)),_dafny.Set.fromElements(_838_v0, _838_v0)]);
              }
            }
            return _coll28;
          }();
          let _913_v64;
          let _nw131 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _913_v64 = _nw131;
          let _index137 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_913_v64).length));
          (_913_v64)[_index137] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(948)), ((_914_v38, _915_v50) => function (_916_i7) {
            return (_914_v38)[_module.__default.safeIndex((_915_v50)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_915_v50).length))], new BigNumber((_914_v38).length))];
          })(_885_v38, _899_v50));
          let _index138 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_913_v64).length));
          (_913_v64)[_index138] = _885_v38;
          let _917_v65;
          let _nw132 = new _module.C0();
          _nw132.__ctor(_this.f9, _887_v40);
          _917_v65 = _nw132;
          let _918_v66;
          let _nw133 = Array((new BigNumber(8)).toNumber()).fill([]);
          _918_v66 = _nw133;
          let _919_v67;
          let _nw134 = Array((new BigNumber(2)).toNumber());
          _nw134[(_dafny.ZERO).toNumber()] = (_this).f18;
          _nw134[(_dafny.ONE).toNumber()] = (_this).f6;
          _919_v67 = _nw134;
          let _index139 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_918_v66).length));
          (_918_v66)[_index139] = _919_v67;
          let _920_v68;
          let _init21 = ((_921_v65, _922_v49) => function (_923_i8) {
            return _dafny.Map.Empty.slice().updateUnsafe(_921_v65.f20,_922_v49);
          })(_917_v65, _898_v49);
          let _nw135 = Array((new BigNumber(4)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw135.length); _i0_21++) {
            _nw135[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _920_v68 = _nw135;
          let _924_v69;
          _924_v69 = _module.D9.create_DC26(_920_v68);
          let _925_v70;
          _925_v70 = _dafny.Seq.of(_917_v65.f20, _898_v49);
          let _926_v71;
          _926_v71 = _dafny.Map.Empty.slice().updateUnsafe(_924_v69,_dafny.Seq.update(_925_v70, _module.__default.safeIndex((_this).f18, new BigNumber((_925_v70).length)), _887_v40));
          let _index140 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_918_v66).length));
          let _nw136 = Array((new BigNumber(12)).toNumber());
          _nw136[(_dafny.ZERO).toNumber()] = p1;
          _nw136[(_dafny.ONE).toNumber()] = new BigNumber((_926_v71).length);
          _nw136[(new BigNumber(2)).toNumber()] = (_this).f18;
          _nw136[(new BigNumber(3)).toNumber()] = new BigNumber(566);
          _nw136[(new BigNumber(4)).toNumber()] = ((_897_v48) ? ((_this).f6) : ((_dafny.ZERO).minus((_this).f17)));
          _nw136[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(114), _889_i5);
          _nw136[(new BigNumber(6)).toNumber()] = new BigNumber(-915);
          _nw136[(new BigNumber(7)).toNumber()] = (_this).f17;
          _nw136[(new BigNumber(8)).toNumber()] = (_this).f6;
          _nw136[(new BigNumber(9)).toNumber()] = p1;
          _nw136[(new BigNumber(10)).toNumber()] = (_899_v50)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_899_v50).length))];
          _nw136[(new BigNumber(11)).toNumber()] = (_this).f6;
          (_918_v66)[_index140] = _nw136;
        }
      }
      let _927_v72;
      _927_v72 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,_885_v38);
      let _928_i9;
      _928_i9 = _dafny.ZERO;
      L5: {
        while (!(_927_v72).equals(_927_v72)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_928_i9)) {
              break L5;
            }
            _928_i9 = (_928_i9).plus(_dafny.ONE);
            let _929_v73;
            _929_v73 = _dafny.Seq.of(_887_v40);
            let _930_v74;
            _930_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_885_v38).length),_929_v73);
            let _931_v75;
            _931_v75 = _dafny.Map.Empty.slice().updateUnsafe(p1,_887_v40);
            let _932_v76;
            _932_v76 = _module.D5.create_DC17(_840_v2, _887_v40);
            let _933_v77;
            _933_v77 = _dafny.Seq.of(_module.D5.create_DC17(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(868)), ((_934_v40) => function (_935_i10) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_934_v40,(_this).f18)).length);
})(_887_v40)), _module.__default.safeIndex(new BigNumber((_930_v74).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(868)), ((_936_v40) => function (_937_i10) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_936_v40,(_this).f18)).length);
})(_887_v40))).length)), new BigNumber((_931_v75).length)), !(_887_v40)), _module.D5.create_DC17(_840_v2, _887_v40), _932_v76);
            let _938_v78;
            _938_v78 = _dafny.Seq.of(_module.__default.fm35((_this).f17, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_this.f9)).length), _887_v40, new BigNumber((_888_v41).length), globalState));
            let _939_v79;
            _939_v79 = _dafny.Map.Empty.slice().updateUnsafe(_933_v77,((_887_v40) ? ((_this).f17) : (new BigNumber((((_938_v78)[_module.__default.safeIndex((_this).f17, new BigNumber((_938_v78).length))]).update((_this).f17, _module.__default.abs((_this).f6))).cardinality()))));
            _939_v79 = _939_v79;
            let _940_v80;
            _940_v80 = _dafny.Set.fromElements((_this).f6);
            let _941_v81;
            _941_v81 = _dafny.Seq.of(_dafny.Seq.update(_885_v38, _module.__default.safeIndex(p1, new BigNumber((_885_v38).length)), _module.__default.fm0((_this).f6, _940_v80, (_929_v73)[_module.__default.safeIndex((_this).f18, new BigNumber((_929_v73).length))], globalState)));
            let _942_v82;
            let _nw137 = Array((new BigNumber(17)).toNumber());
            _nw137[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_885_v38, _dafny.Seq.UnicodeFromString("dpxwvvn"));
            _nw137[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("oplylgwjk");
            _nw137[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(439)), ((_943_v0) => function (_944_i11) {
              return _943_v0;
            })(_838_v0));
            _nw137[(new BigNumber(3)).toNumber()] = _885_v38;
            _nw137[(new BigNumber(4)).toNumber()] = _885_v38;
            _nw137[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("huf");
            _nw137[(new BigNumber(6)).toNumber()] = _module.__default.fm30(_887_v40, new _dafny.CodePoint('n'.codePointAt(0)), globalState);
            _nw137[(new BigNumber(7)).toNumber()] = _885_v38;
            _nw137[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("a"), _885_v38);
            _nw137[(new BigNumber(9)).toNumber()] = _885_v38;
            _nw137[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("xkaku");
            _nw137[(new BigNumber(11)).toNumber()] = _885_v38;
            _nw137[(new BigNumber(12)).toNumber()] = _dafny.Seq.update((_941_v81)[_module.__default.safeIndex(new BigNumber((_929_v73).length), new BigNumber((_941_v81).length))], _module.__default.safeIndex((_this).f18, new BigNumber(((_941_v81)[_module.__default.safeIndex(new BigNumber((_929_v73).length), new BigNumber((_941_v81).length))]).length)), _838_v0);
            _nw137[(new BigNumber(13)).toNumber()] = _module.__default.fm25((_this).f17, _887_v40, globalState);
            _nw137[(new BigNumber(14)).toNumber()] = _module.__default.fm30(_887_v40, _838_v0, globalState);
            _nw137[(new BigNumber(15)).toNumber()] = _885_v38;
            _nw137[(new BigNumber(16)).toNumber()] = _885_v38;
            _942_v82 = _nw137;
            let _945_v83;
            _945_v83 = _dafny.Seq.of(_942_v82, _942_v82, _942_v82, _942_v82);
            _942_v82 = (_945_v83)[_module.__default.safeIndex(new BigNumber(91), new BigNumber((_945_v83).length))];
            _887_v40 = (new BigNumber(-609)).isLessThanOrEqualTo((_this).f18);
            if (!(_887_v40)) {
              let _946_v84;
              _946_v84 = _dafny.Seq.of(_module.D5.create_DC16(_888_v41));
              let _947_v85;
              _947_v85 = _dafny.Set.fromElements(_946_v84);
              let _948_v86;
              _948_v86 = _dafny.Map.Empty.slice().updateUnsafe(_946_v84,(_this).f6);
              _947_v85 = function () {
                let _coll30 = new _dafny.Set();
                for (const _compr_30 of (_948_v86).Keys.Elements) {
                  let _949_v87 = _compr_30;
                  if ((_948_v86).contains(_949_v87)) {
                    _coll30.add(_949_v87);
                  }
                }
                return _coll30;
              }();
              let _950_v88;
              _950_v88 = _module.D9.create_DC25(((_887_v40) ? (p1) : ((_this).f6)), (((_839_v1).contains(_887_v40)) ? ((_839_v1).get(_887_v40)) : (new BigNumber(139))), ((_887_v40) ? (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-255)), function (_951_i12) {
  return (_this).f18;
})).length)) : ((_dafny.ZERO).minus((_this).f18))), _838_v0);
              _950_v88 = _950_v88;
              let _952_v89;
              _952_v89 = _module.D0.create_DC0(_929_v73, (_this).f18);
              let _rhs155 = _dafny.Seq.Concat(_885_v38, _885_v38);
              let _rhs156 = (((_887_v40) ? ((_952_v89).dtor_cf1) : ((_this).f18))).plus(((_this).f17).multipliedBy((_this).f17));
              let _lhs128 = globalState;
              _lhs128.f2 = _rhs155;
              r3 = _rhs156;
              let _953_v90;
              let _nw138 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
              _953_v90 = _nw138;
              let _index141 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_953_v90).length));
              (_953_v90)[_index141] = (_this).f18;
              let _index142 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_953_v90).length));
              let _rhs157 = _887_v40;
              let _rhs158 = _module.__default.safeModuloInt(new BigNumber(-824), p1);
              let _rhs159 = (_dafny.Set.fromElements(_887_v40)).Difference(_888_v41);
              let _rhs160 = (_module.D4.create_DC13(_840_v2, _887_v40, (_dafny.ZERO).minus((_this).f17))).dtor_cf27;
              let _lhs129 = globalState;
              let _lhs130 = _953_v90;
              let _lhs131 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_953_v90).length));
              _lhs129.f0 = _rhs157;
              r3 = _rhs158;
              _888_v41 = _rhs159;
              _lhs130[_lhs131] = _rhs160;
              r3 = new BigNumber((_dafny.Seq.Concat(_885_v38, _885_v38)).length);
            } else {
              (globalState).f0 = _887_v40;
              let _arr26 = _this.f9;
              let _index143 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_this.f9).length));
              _arr26[_index143] = _887_v40;
              let _954_v91;
              _954_v91 = _dafny.Seq.of(_840_v2, _840_v2);
              let _arr27 = _this.f9;
              let _index144 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_this.f9).length));
              let _rhs161 = (_module.__default.safeDivisionInt((_this).f18, (_dafny.ZERO).minus(p1))).isLessThan(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("afv"), _885_v38)).length));
              let _rhs162 = _954_v91;
              let _rhs163 = (_this).f6;
              let _lhs132 = _this.f9;
              let _lhs133 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_this.f9).length));
              let _lhs134 = globalState;
              _lhs132[_lhs133] = _rhs161;
              _954_v91 = _rhs162;
              _lhs134.f5 = _rhs163;
              let _arr28 = _this.f9;
              let _index145 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_this.f9).length));
              _arr28[_index145] = (_this.f9)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_this.f9).length))];
              let _955_v92;
              _955_v92 = _839_v1;
              let _956_v93;
              _956_v93 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_955_v92));
              _839_v1 = (_839_v1).Union(((_839_v1).update(false, _module.__default.abs((_this).f17))).Difference((((_956_v93).contains((_this).fm21(new BigNumber((_929_v73).length), globalState))) ? ((_956_v93).get((_this).fm21(new BigNumber((_929_v73).length), globalState))) : (_dafny.MultiSet.fromElements(_887_v40, (_this.f9)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_this.f9).length))])))));
              (globalState).f5 = (_dafny.ZERO).minus((_this).f6);
            }
          }
        }
      }
      let _957_v94;
      _957_v94 = _module.D7.create_DC22();
      let _pat_let_tv18 = _838_v0;
      r0 = function (_source12) {
        if (_source12.is_DC22) {
          return _pat_let_tv18;
        } else {
          let _958___mcc_h0 = (_source12).cf43;
          let _959_cf43 = _958___mcc_h0;
          return new _dafny.CodePoint('r'.codePointAt(0));
        }
      }(((_887_v40) ? (_957_v94) : (_957_v94)));
      let _960_v95;
      _960_v95 = _dafny.MultiSet.fromElements(_this.f9);
      r1 = (((_960_v95).contains(_this.f9)) ? ((_960_v95).get(_this.f9)) : ((_this).f18));
      r2 = _this.f9;
      r3 = (_this).f17;
      return [r0, r1, r2, r3];
    }
    m12(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = [];
      let _961_v0;
      _961_v0 = _dafny.Seq.of((_this).f6);
      let _962_v1;
      _962_v1 = _dafny.Seq.of(((_this).f18).minus((_this).f17), (_961_v0)[_module.__default.safeIndex((_this).f18, new BigNumber((_961_v0).length))], (_this).f18, _module.__default.safeDivisionInt(new BigNumber(880), (_this).f6), ((_this).f6).plus((_this).f18));
      let _963_v2;
      _963_v2 = _dafny.Seq.UnicodeFromString("xxnkqk");
      let _rhs164 = _963_v2;
      let _rhs165 = _961_v0;
      let _rhs166 = (new BigNumber(899)).minus((_this).f17);
      let _lhs135 = globalState;
      let _lhs136 = globalState;
      _lhs135.f2 = _rhs164;
      _961_v0 = _rhs165;
      _lhs136.f5 = _rhs166;
      let _arr29 = _this.f9;
      let _index146 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_this.f9).length));
      _arr29[_index146] = p0;
      let _964_v3;
      _964_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(240),p0);
      let _965_v4;
      _965_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm22((_dafny.MultiSet.fromElements(!(true))), (((_964_v3).contains((_this).f17)) ? ((_964_v3).get((_this).f17)) : (p1)), new BigNumber((_dafny.Seq.of((_this).f17, (_this).f17, (_this).f17)).length), globalState),!(p1) || (p1));
      let _966_v5;
      _966_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f6);
      let _967_v6;
      _967_v6 = _dafny.MultiSet.fromElements((_this).f18);
      let _968_v8;
      _968_v8 = _dafny.Set.fromElements((_this).f6);
      let _969_v9;
      let _nw139 = new _module.C2();
      _nw139.__ctor(new BigNumber(((_966_v5).Merge(_966_v5)).length), _dafny.MultiSet.fromElements(_967_v6), _this.f9, (function () {
        let _coll31 = new _dafny.Map();
        for (const _compr_31 of (_968_v8).Elements) {
          let _970_v7 = _compr_31;
          if ((_968_v8).contains(_970_v7)) {
            _coll31.push([_module.__default.safeModuloInt(_970_v7, new BigNumber(-615)),_dafny.Seq.UnicodeFromString("m")]);
          }
        }
        return _coll31;
      }()).update((_this).f17, _963_v2));
      _969_v9 = _nw139;
      let _971_v10;
      _971_v10 = _dafny.Seq.of(_964_v3);
      let _972_v11;
      _972_v11 = _dafny.Seq.of(true);
      let _973_v12;
      _973_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,_dafny.MultiSet.FromArray(_dafny.Seq.update(_972_v11, _module.__default.safeIndex(new BigNumber((_963_v2).length), new BigNumber((_972_v11).length)), p1)));
      let _974_v13;
      _974_v13 = _dafny.Seq.of(((_971_v10)[_module.__default.safeIndex((_this).f6, new BigNumber((_971_v10).length))]).update(new BigNumber((_973_v12).length), p1), _964_v3);
      let _975_v14;
      _975_v14 = _dafny.Set.fromElements(!(false), ((_this).f17).isLessThan((_this).f6), _dafny.areEqual(_971_v10, _974_v13));
      let _976_v15;
      _976_v15 = new _dafny.CodePoint('y'.codePointAt(0));
      let _977_v16;
      _977_v16 = _dafny.Map.Empty.slice().updateUnsafe(_976_v15,!((_972_v11)[_module.__default.safeIndex((_this).f6, new BigNumber((_972_v11).length))]));
      let _arr30 = _this.f9;
      let _index147 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_this.f9).length));
      let _rhs167 = (((_977_v16).contains(_976_v15)) ? ((_977_v16).get(_976_v15)) : (_module.__default.fm3(globalState)));
      let _rhs168 = _965_v4;
      let _rhs169 = _969_v9;
      let _rhs170 = _975_v14;
      let _rhs171 = p1;
      let _lhs137 = _this.f9;
      let _lhs138 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_this.f9).length));
      _lhs137[_lhs138] = _rhs167;
      _965_v4 = _rhs168;
      _969_v9 = _rhs169;
      _975_v14 = _rhs170;
      r1 = _rhs171;
      let _hi8 = ((!(p1)) ? (new BigNumber((_972_v11).length)) : ((_this).f6));
      for (let _978_i0 = (_this).f18; _978_i0.isLessThan(_hi8); _978_i0 = _978_i0.plus(_dafny.ONE)) {
        if (_dafny.Seq.contains(_972_v11, !((_dafny.ZERO).minus((_969_v9).fm21(new BigNumber((_967_v6).cardinality()), globalState))).isEqualTo((_dafny.ZERO).minus((_this).f17)))) {
          (globalState).f5 = ((p0) ? (_978_i0) : (_978_i0));
          (globalState).f0 = p0;
          let _979_v17;
          let _nw140 = Array((new BigNumber(28)).toNumber());
          _979_v17 = _nw140;
          let _980_v18;
          _980_v18 = _dafny.Map.Empty.slice().updateUnsafe(_979_v17,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f17)).length),(_this).f17));
          _980_v18 = _980_v18;
          (globalState).f0 = !(!(p1));
          let _981_v19;
          _981_v19 = _dafny.Seq.of(_961_v0, _961_v0, _961_v0, _dafny.Seq.of((_this).f6));
          let _982_v20;
          _982_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_981_v19, _module.__default.fm43(_976_v15, (_this).f17, globalState)),(_this).f6);
          _982_v20 = (_982_v20).update(_981_v19, (_this).f17);
        } else {
          let _983_v21;
          _983_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this.f9)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_this.f9).length))],(_this).f18);
          (globalState).f5 = (_module.D0.create_DC1(p1, new BigNumber((_dafny.Set.fromElements((_this).f6, new BigNumber((_983_v21).length))).length), _976_v15)).dtor_cf3;
          let _984_v22;
          let _nw141 = Array((new BigNumber(23)).toNumber()).fill(false);
          _984_v22 = _nw141;
          let _985_v23;
          let _nw142 = new _module.C2();
          _nw142.__ctor((_dafny.ZERO).minus((_this).f18), (_this.f16).Intersect(_this.f16), _984_v22, _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Seq.update(_963_v2, _module.__default.safeIndex((_this).f18, new BigNumber((_963_v2).length)), _976_v15)));
          _985_v23 = _nw142;
          let _986_v24;
          _986_v24 = _module.D3.create_DC9(_dafny.Seq.update(_962_v1, _module.__default.safeIndex((_this).f6, new BigNumber((_962_v1).length)), _module.__default.fm2(_978_i0, globalState)));
          let _987_v25;
          let _nw143 = new _module.C2();
          _nw143.__ctor((_this).f18, (_this.f16).update(_dafny.MultiSet.FromArray((_986_v24).dtor_cf21), _module.__default.abs(new BigNumber((_963_v2).length))), (_984_v22), (_this).f10);
          _987_v25 = _nw143;
          (globalState).f0 = p1;
          _976_v15 = (((_this.f9)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_this.f9).length))]) ? (_976_v15) : (_976_v15));
        }
        (globalState).f5 = _978_i0;
        let _988_v26;
        _988_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,_978_i0);
        _988_v26 = (_988_v26).update(((_this).f18).isLessThan((_this).f18), _978_i0);
        let _989_v27;
        let _990_v28;
        let _991_v29;
        let _out28;
        let _out29;
        let _out30;
        let _outcollector10 = (_this).m13((_this).f18, globalState);
        _out28 = _outcollector10[0];
        _out29 = _outcollector10[1];
        _out30 = _outcollector10[2];
        _989_v27 = _out28;
        _990_v28 = _out29;
        _991_v29 = _out30;
      }
      let _992_v30;
      let _nw144 = Array((new BigNumber(5)).toNumber()).fill(false);
      _992_v30 = _nw144;
      let _993_v31;
      let _nw145 = new _module.C0();
      _nw145.__ctor(_992_v30, (_this.f9)[_module.__default.safeIndex(new BigNumber(202), new BigNumber((_this.f9).length))]);
      _993_v31 = _nw145;
      let _994_v32;
      _994_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17);
      let _995_v33;
      _995_v33 = _dafny.MultiSet.fromElements(p1);
      _967_v6 = _module.__default.fm35(new BigNumber((_994_v32).length), (_this).f17, ((_this).f18).isEqualTo(new BigNumber((_967_v6).cardinality())), (((_995_v33).contains(p1)) ? ((_995_v33).get(p1)) : ((_this).f17)), globalState);
      let _arr31 = _this.f9;
      let _index148 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_this.f9).length));
      let _rhs172 = false;
      let _rhs173 = _976_v15;
      let _lhs139 = _this.f9;
      let _lhs140 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_this.f9).length));
      _lhs139[_lhs140] = _rhs172;
      _976_v15 = _rhs173;
      r0 = p1;
      r1 = p0;
      let _996_v34;
      let _init22 = ((_997_v33) => function (_998_i1) {
        return _module.__default.safeDivisionInt(_998_i1, new BigNumber((_997_v33).cardinality()));
      })(_995_v33);
      let _nw146 = Array((new BigNumber(24)).toNumber());
      for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw146.length); _i0_22++) {
        _nw146[_i0_22] = _init22(new BigNumber(_i0_22));
      }
      _996_v34 = _nw146;
      r2 = _996_v34;
      return [r0, r1, r2];
    }
    m13(p0, globalState) {
      let _this = this;
      let r0 = _module.D2.Default();
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.Seq.UnicodeFromString("");
      if (((_this).f17).isEqualTo(_module.__default.safeDivisionInt(p0, p0))) {
        let _999_v0;
        _999_v0 = false;
        let _1000_v1;
        _1000_v1 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1001_v2;
        _1001_v2 = _dafny.Seq.of((_this).f18);
        let _1002_v3;
        _1002_v3 = _module.D4.create_DC13(_1001_v2, _999_v0, (_this).f18);
        let _1003_v4;
        _1003_v4 = _dafny.MultiSet.fromElements((_1002_v3).dtor_cf26);
        let _1004_v5;
        _1004_v5 = _dafny.Set.fromElements(_999_v0, true, _999_v0);
        let _1005_v6;
        _1005_v6 = _dafny.Map.Empty.slice().updateUnsafe(((_999_v0) ? (_1000_v1) : (new _dafny.CodePoint('a'.codePointAt(0)))),_module.__default.fm44(p0, new BigNumber((_1003_v4).cardinality()), _999_v0, new BigNumber((_1004_v5).length), globalState));
        let _1006_v7;
        _1006_v7 = _module.D9.create_DC25((_dafny.ZERO).minus((_this).f17), new BigNumber((_1001_v2).length), (_this).f17, _1000_v1);
        _1005_v6 = (_1005_v6).update(_1000_v1, _dafny.Map.Empty.slice().updateUnsafe(p0,_1006_v7));
        if (false) {
          let _1007_v8;
          let _nw147 = new _module.C2();
          _nw147.__ctor(new BigNumber(-796), _this.f16, _this.f9, (_this).f10);
          _1007_v8 = _nw147;
          let _1008_v9;
          _1008_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1007_v8,_999_v0);
          let _1009_v10;
          _1009_v10 = _dafny.Seq.of(!(_999_v0));
          let _1010_v11;
          _1010_v11 = _1008_v9;
          let _1011_v12;
          let _nw148 = Array((new BigNumber(28)).toNumber());
          _nw148[(_dafny.ZERO).toNumber()] = _1008_v9;
          _nw148[(_dafny.ONE).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(2)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(3)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(4)).toNumber()] = ((_999_v0) ? (_dafny.Map.Empty.slice().updateUnsafe(_1007_v8,_999_v0)) : (_1008_v9));
          _nw148[(new BigNumber(5)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(6)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(7)).toNumber()] = (_1008_v9).Merge(_1008_v9);
          _nw148[(new BigNumber(8)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(9)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(10)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(11)).toNumber()] = ((true) ? (_1008_v9) : (_1008_v9));
          _nw148[(new BigNumber(12)).toNumber()] = (_1008_v9).update(_1007_v8, _999_v0);
          _nw148[(new BigNumber(13)).toNumber()] = (_1008_v9).update(_1007_v8, true);
          _nw148[(new BigNumber(14)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1007_v8,_999_v0)).Merge(_1008_v9);
          _nw148[(new BigNumber(16)).toNumber()] = (_1008_v9).Merge((_1008_v9).update(_1007_v8, (_1009_v10)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("tk")).length), new BigNumber((_1009_v10).length))]));
          _nw148[(new BigNumber(17)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1007_v8,_999_v0)).Merge(_1008_v9);
          _nw148[(new BigNumber(18)).toNumber()] = (_1010_v11);
          _nw148[(new BigNumber(19)).toNumber()] = (_1008_v9).update(_1007_v8, _999_v0);
          _nw148[(new BigNumber(20)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(21)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(22)).toNumber()] = (_1008_v9).Merge(_1008_v9);
          _nw148[(new BigNumber(23)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(24)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(25)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(26)).toNumber()] = _1008_v9;
          _nw148[(new BigNumber(27)).toNumber()] = _1008_v9;
          _1011_v12 = _nw148;
          let _index149 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1011_v12).length));
          (_1011_v12)[_index149] = (_1008_v9).Merge(_1008_v9);
          let _1012_v13;
          _1012_v13 = _dafny.Map.Empty.slice().updateUnsafe(_999_v0,_dafny.Map.Empty.slice().updateUnsafe(_999_v0,true));
          let _1013_v14;
          _1013_v14 = _dafny.Map.Empty.slice().updateUnsafe(_999_v0,false);
          let _index150 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1011_v12).length));
          let _rhs174 = _1008_v9;
          let _rhs175 = _module.__default.fm45((new BigNumber((_1001_v2).length)).minus((_this).f17), (_this).f6, globalState);
          let _rhs176 = _dafny.Seq.Concat(_1009_v10, _dafny.Seq.Concat(_1009_v10, _module.__default.fm4((_this).f18, p0, new BigNumber(((((_1012_v13).contains(_999_v0)) ? ((_1012_v13).get(_999_v0)) : (_1013_v14))).length), globalState)));
          let _rhs177 = _999_v0;
          let _rhs178 = _999_v0;
          let _lhs141 = _1011_v12;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1011_v12).length));
          let _lhs143 = globalState;
          let _lhs144 = globalState;
          _lhs141[_lhs142] = _rhs174;
          _1002_v3 = _rhs175;
          _1009_v10 = _rhs176;
          _lhs143.f0 = _rhs177;
          _lhs144.f0 = _rhs178;
          let _1014_v15;
          let _nw149 = new _module.C1();
          _nw149.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("uemcv")).length));
          _1014_v15 = _nw149;
          let _1015_v16;
          let _init23 = function (_1016_i0) {
            return (_1016_i0).multipliedBy((_dafny.ZERO).minus((_this).f17));
          };
          let _nw150 = Array((new BigNumber(23)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw150.length); _i0_23++) {
            _nw150[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1015_v16 = _nw150;
          let _1017_v17;
          _1017_v17 = _module.D2.create_DC6((_this).f17, _1015_v16);
          _1015_v16 = (_1017_v17).dtor_cf14;
          let _1018_v18;
          let _nw151 = new _module.C1();
          _nw151.__ctor((_this).f18);
          _1018_v18 = _nw151;
          (_this).f9 = _this.f9;
        } else {
          let _1019_v19;
          let _nw152 = new _module.C2();
          _nw152.__ctor((_this).f17, _this.f16, _this.f9, _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.Seq.UnicodeFromString("nbgkv")));
          _1019_v19 = _nw152;
          let _1020_v20;
          _1020_v20 = _module.D6.create_DC18(_1019_v19);
          let _1021_v21;
          _1021_v21 = _module.D6.create_DC20(_1020_v20);
          _1021_v21 = _module.D6.create_DC20(_1020_v20);
          let _1022_v22;
          _1022_v22 = _dafny.Map.Empty.slice().updateUnsafe(_999_v0,_999_v0);
          let _1023_v23;
          _1023_v23 = _dafny.Seq.of(_999_v0, (((_1022_v22).contains(false)) ? ((_1022_v22).get(false)) : (true)));
          let _1024_v24;
          _1024_v24 = _module.D6.create_DC19(new BigNumber((_1023_v23).length), new BigNumber((_1001_v2).length), (_1019_v19).f6);
          let _1025_v25;
          _1025_v25 = _dafny.MultiSet.fromElements(_1024_v24, _1024_v24, _1024_v24, _1024_v24);
          (globalState).f5 = (((_1025_v25).contains(_1024_v24)) ? ((_1025_v25).get(_1024_v24)) : ((_this).f6));
          let _1026_v26;
          let _nw153 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1026_v26 = _nw153;
          let _1027_v27;
          _1027_v27 = _dafny.Map.Empty.slice().updateUnsafe(!(_999_v0) || (_999_v0),_1026_v26);
          let _1028_v28;
          _1028_v28 = _dafny.Seq.of(_1026_v26, _1026_v26);
          _1027_v27 = (_1027_v27).update(_999_v0, (((_1027_v27).contains(_999_v0)) ? ((_1027_v27).get(_999_v0)) : ((_1028_v28)[_module.__default.safeIndex((_this).f17, new BigNumber((_1028_v28).length))])));
          let _1029_v29;
          let _nw154 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1029_v29 = _nw154;
          let _1030_v30;
          _1030_v30 = _dafny.Map.Empty.slice().updateUnsafe(_999_v0,p0);
          let _index151 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_1029_v29).length));
          (_1029_v29)[_index151] = _module.__default.safeDivisionInt(new BigNumber(264), new BigNumber((_1030_v30).length));
          let _1031_v31;
          _1031_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(914),_999_v0);
          let _index152 = _module.__default.safeIndex(new BigNumber(768), new BigNumber((_1029_v29).length));
          (_1029_v29)[_index152] = _module.__default.fm2((((_1030_v30).contains(_999_v0)) ? ((_1030_v30).get(_999_v0)) : (new BigNumber((_1031_v31).length))), globalState);
          let _1032_v32;
          _1032_v32 = _dafny.Seq.UnicodeFromString("lnnj");
          let _1033_v33;
          _1033_v33 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hieeu"), _1032_v32), _1032_v32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(130)), ((_1034_v1) => function (_1035_i1) {
            return _1034_v1;
          })(_1000_v1)));
          let _rhs179 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(704)), function (_1036_i2) {
            return _dafny.Seq.UnicodeFromString("xrfic");
          });
          let _rhs180 = (_this).f18;
          let _lhs145 = globalState;
          _1033_v33 = _rhs179;
          _lhs145.f5 = _rhs180;
        }
        let _1037_v34;
        _1037_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm22(_dafny.MultiSet.fromElements(_999_v0, !(_999_v0)), _999_v0, new BigNumber(66), globalState),(_this).f6);
        let _1038_v35;
        let _init24 = function (_1039_i3) {
          return _module.__default.safeDivisionInt(_1039_i3, new BigNumber(-228));
        };
        let _nw155 = Array((new BigNumber(12)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw155.length); _i0_24++) {
          _nw155[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _1038_v35 = _nw155;
        let _1040_v36;
        _1040_v36 = _dafny.MultiSet.fromElements((_this).f17);
        let _index153 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_1038_v35).length));
        (_1038_v35)[_index153] = ((_999_v0) ? (p0) : (new BigNumber((_1040_v36).cardinality())));
        let _1041_v37;
        _1041_v37 = _dafny.Seq.UnicodeFromString("w");
        let _index154 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_1038_v35).length));
        let _rhs181 = new BigNumber((_dafny.Seq.Concat(_1041_v37, _1041_v37)).length);
        let _rhs182 = _1037_v34;
        let _rhs183 = (_dafny.ZERO).minus(new BigNumber(-547));
        let _rhs184 = (_this).f18;
        let _rhs185 = ((_999_v0) ? (_1038_v35) : (((_999_v0) ? (_1038_v35) : (_1038_v35))));
        let _lhs146 = globalState;
        let _lhs147 = _1038_v35;
        let _lhs148 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_1038_v35).length));
        let _lhs149 = globalState;
        _lhs146.f5 = _rhs181;
        _1037_v34 = _rhs182;
        _lhs147[_lhs148] = _rhs183;
        _lhs149.f5 = _rhs184;
        _1038_v35 = _rhs185;
        if (_999_v0) {
          (globalState).f5 = _module.__default.safeDivisionInt((_this).f6, p0);
          let _1042_v38;
          let _nw156 = new _module.C2();
          _nw156.__ctor((p0).plus(new BigNumber((_1041_v37).length)), _this.f16, _this.f9, ((_this).f10).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f17,_dafny.Seq.update(_1041_v37, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(939)), ((_1043_v1) => function (_1044_i4) {
            return _1043_v1;
          })(_1000_v1))).length), new BigNumber((_1041_v37).length)), _1000_v1))));
          _1042_v38 = _nw156;
          let _1045_v39;
          let _nw157 = new _module.C1();
          _nw157.__ctor(_module.__default.safeDivisionInt((_this).f17, (_this).f17));
          _1045_v39 = _nw157;
          (globalState).f0 = false;
          _1037_v34 = (_1037_v34).update(((_this).f18).isLessThanOrEqualTo(new BigNumber(-227)), new BigNumber(68));
        } else {
          let _1046_v40;
          _1046_v40 = _dafny.Map.Empty.slice().updateUnsafe(_999_v0,_999_v0);
          _1046_v40 = (_1046_v40).update(_999_v0, _999_v0);
          let _1047_v41;
          let _nw158 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1047_v41 = _nw158;
          let _1048_v42;
          let _nw159 = Array((new BigNumber(6)).toNumber());
          _nw159[(_dafny.ZERO).toNumber()] = _1047_v41;
          _nw159[(_dafny.ONE).toNumber()] = _1047_v41;
          _nw159[(new BigNumber(2)).toNumber()] = _1047_v41;
          _nw159[(new BigNumber(3)).toNumber()] = _1047_v41;
          _nw159[(new BigNumber(4)).toNumber()] = _1047_v41;
          _nw159[(new BigNumber(5)).toNumber()] = _1047_v41;
          _1048_v42 = _nw159;
          let _index155 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_1048_v42).length));
          (_1048_v42)[_index155] = _1047_v41;
          let _index156 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_1048_v42).length));
          let _rhs186 = _1047_v41;
          let _rhs187 = (_this).f17;
          let _rhs188 = !(_999_v0);
          let _lhs150 = _1048_v42;
          let _lhs151 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_1048_v42).length));
          let _lhs152 = globalState;
          let _lhs153 = globalState;
          _lhs150[_lhs151] = _rhs186;
          _lhs152.f5 = _rhs187;
          _lhs153.f0 = _rhs188;
          let _1049_v43;
          _1049_v43 = _dafny.Seq.of(_999_v0);
          r1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_999_v0), _1049_v43), _1049_v43);
          (globalState).f5 = (_this).f17;
          let _arr32 = _this.f9;
          let _index157 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_this.f9).length));
          _arr32[_index157] = _999_v0;
          let _arr33 = _this.f9;
          let _index158 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_this.f9).length));
          _arr33[_index158] = (!(_999_v0) || (_999_v0)) === (_999_v0);
        }
        let _1050_v44;
        let _init25 = ((_1051_v0, _1052_v5) => function (_1053_i5) {
          return ((_1051_v0) ? (_module.D5.create_DC16(_dafny.Set.fromElements(_1051_v0))) : (_module.D5.create_DC16(_1052_v5)));
        })(_999_v0, _1004_v5);
        let _nw160 = Array((new BigNumber(10)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw160.length); _i0_25++) {
          _nw160[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1050_v44 = _nw160;
        let _1054_v45;
        _1054_v45 = _module.D5.create_DC16(_1004_v5);
        let _index159 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1050_v44).length));
        (_1050_v44)[_index159] = _1054_v45;
        let _index160 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1050_v44).length));
        (_1050_v44)[_index160] = _1054_v45;
      } else {
        let _1055_v46;
        _1055_v46 = false;
        let _1056_v47;
        _1056_v47 = _dafny.Seq.of(p0);
        let _1057_v48;
        _1057_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1055_v46,_1056_v47);
        let _1058_v49;
        _1058_v49 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1059_v50;
        _1059_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1055_v46);
        let _1060_v51;
        _1060_v51 = _dafny.MultiSet.fromElements(_1059_v50);
        _1057_v48 = (_1057_v48).update((p0).isLessThan(p0), _dafny.Seq.Concat(_module.__default.fm26((_this).f6, _1058_v49, !(_1055_v46), _1058_v49, globalState), _dafny.Seq.of(_module.__default.fm2(new BigNumber((_1060_v51).cardinality()), globalState), new BigNumber(904))));
        let _1061_v52;
        _1061_v52 = _module.D6.create_DC19(new BigNumber(313), p0, (_this).f17);
        let _1062_v53;
        _1062_v53 = _module.D6.create_DC20(_1061_v52);
        let _source13 = _1062_v53;
        if (_source13.is_DC19) {
          let _1063___mcc_h0 = (_source13).cf39;
          let _1064___mcc_h1 = (_source13).cf40;
          let _1065___mcc_h2 = (_source13).cf41;
          let _1066_cf41 = _1065___mcc_h2;
          let _1067_cf40 = _1064___mcc_h1;
          let _1068_cf39 = _1063___mcc_h0;
          let _arr34 = _this.f9;
          let _index161 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_this.f9).length));
          _arr34[_index161] = _1055_v46;
          let _arr35 = _this.f9;
          let _index162 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_this.f9).length));
          _arr35[_index162] = _1055_v46;
          let _1069_v54;
          let _nw161 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1069_v54 = _nw161;
          let _index163 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1069_v54).length));
          (_1069_v54)[_index163] = (_this).f6;
          let _1070_v55;
          _1070_v55 = _dafny.Seq.of(_1069_v54, _1069_v54, _1069_v54, _1069_v54, _1069_v54);
          let _1071_v56;
          _1071_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1058_v49,(_this).f10);
          let _1072_v57;
          _1072_v57 = _module.D3.create_DC9(_1056_v47);
          let _1073_v58;
          _1073_v58 = _dafny.MultiSet.fromElements((_1072_v57).dtor_cf21);
          let _1074_v59;
          _1074_v59 = _module.D0.create_DC0(_module.__default.fm4(p0, _1067_cf40, new BigNumber((_1073_v58).cardinality()), globalState), new BigNumber(644));
          let _arr36 = _this.f9;
          let _index164 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_this.f9).length));
          let _index165 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1069_v54).length));
          let _rhs189 = ((((_1071_v56).contains(_1058_v49)) ? ((_1071_v56).get(_1058_v49)) : ((_this).f10))).contains((_dafny.ZERO).minus(p0));
          let _rhs190 = (_1074_v59).dtor_cf1;
          let _rhs191 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_1068_cf39), new BigNumber(481));
          let _rhs192 = _1070_v55;
          let _lhs154 = _this.f9;
          let _lhs155 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_this.f9).length));
          let _lhs156 = _1069_v54;
          let _lhs157 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1069_v54).length));
          let _lhs158 = globalState;
          _lhs154[_lhs155] = _rhs189;
          _lhs156[_lhs157] = _rhs190;
          _lhs158.f5 = _rhs191;
          _1070_v55 = _rhs192;
          let _1075_v60;
          let _init26 = ((_1076_v46) => function (_1077_i6) {
            return _dafny.Map.Empty.slice().updateUnsafe(false,_1076_v46);
          })(_1055_v46);
          let _nw162 = Array((new BigNumber(6)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw162.length); _i0_26++) {
            _nw162[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1075_v60 = _nw162;
          let _1078_v61;
          _1078_v61 = _module.D9.create_DC26(_1075_v60);
          _1078_v61 = _1078_v61;
          (globalState).f0 = (_this.f9)[_module.__default.safeIndex(new BigNumber(900), new BigNumber((_this.f9).length))];
        } else if (_source13.is_DC18) {
          let _1079___mcc_h3 = (_source13).cf38;
          let _1080_cf38 = _1079___mcc_h3;
          let _1081_v62;
          _1081_v62 = _dafny.Seq.of(_1056_v47, _1056_v47, _1056_v47);
          _1081_v62 = _1081_v62;
          (globalState).f5 = (_this).f17;
          let _1082_v63;
          _1082_v63 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(59)), function (_1083_i7) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          }));
          let _1084_v64;
          _1084_v64 = _dafny.Seq.UnicodeFromString("oaqqtxvm");
          let _1085_v65;
          _1085_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1055_v46,_1084_v64);
          _1082_v63 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_1085_v65).contains(_1055_v46)) ? ((_1085_v65).get(_1055_v46)) : (_1084_v64)));
          let _1086_v66;
          let _nw163 = Array((new BigNumber(24)).toNumber()).fill([]);
          _1086_v66 = _nw163;
          let _1087_v67;
          let _init27 = function (_1088_i8) {
            return _module.__default.safeDivisionInt(_1088_i8, (_this).f6);
          };
          let _nw164 = Array((new BigNumber(7)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw164.length); _i0_27++) {
            _nw164[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1087_v67 = _nw164;
          let _index166 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1086_v66).length));
          (_1086_v66)[_index166] = ((_1055_v46) ? (_1087_v67) : (_1087_v67));
          let _1089_v68;
          _1089_v68 = _dafny.Seq.of(_1055_v46, _1055_v46, (p0).isLessThanOrEqualTo(p0));
          let _index167 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1086_v66).length));
          let _rhs193 = (_1056_v47)[_module.__default.safeIndex((_this).f17, new BigNumber((_1056_v47).length))];
          let _rhs194 = (_this).f18;
          let _rhs195 = (_1089_v68)[_module.__default.safeIndex((_this).f6, new BigNumber((_1089_v68).length))];
          let _rhs196 = _1087_v67;
          let _lhs159 = globalState;
          let _lhs160 = globalState;
          let _lhs161 = _1086_v66;
          let _lhs162 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1086_v66).length));
          _lhs159.f5 = _rhs193;
          _lhs160.f5 = _rhs194;
          _1055_v46 = _rhs195;
          _lhs161[_lhs162] = _rhs196;
        } else {
          let _1090___mcc_h4 = (_source13).cf42;
          let _1091_cf42 = _1090___mcc_h4;
          let _1092_v69;
          _1092_v69 = _dafny.Seq.of(_1055_v46);
          let _1093_v70;
          _1093_v70 = _dafny.Map.Empty.slice().updateUnsafe((_1092_v69)[_module.__default.safeIndex((_this).f17, new BigNumber((_1092_v69).length))],!(_1055_v46));
          let _1094_v71;
          _1094_v71 = _dafny.Seq.of(_1093_v70);
          _1094_v71 = ((_1055_v46) ? (_dafny.Seq.Concat(_dafny.Seq.of(_1093_v70, _1093_v70), _module.__default.fm46(globalState))) : (_dafny.Seq.Concat(_1094_v71, _1094_v71)));
          (globalState).f0 = false;
          let _1095_v72;
          _1095_v72 = _dafny.MultiSet.fromElements((_this).f17);
          let _1096_v74;
          let _init28 = ((_1097_p0) => function (_1098_i9) {
            return (_1098_i9).plus(_1097_p0);
          })(p0);
          let _nw165 = Array((new BigNumber(19)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw165.length); _i0_28++) {
            _nw165[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1096_v74 = _nw165;
          let _1099_v75;
          _1099_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1055_v46,_1096_v74);
          let _pat_let_tv19 = _1095_v72;
          let _pat_let_tv20 = _1095_v72;
          let _1100_v76;
          _1100_v76 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let23_0) {
            return function (_1101_dt__update__tmp_h0) {
              return function (_pat_let24_0) {
                return function (_1103_dt__update_hcf57_h0) {
                  return _module.D11.create_DC30(_1103_dt__update_hcf57_h0);
                }(_pat_let24_0);
              }(function () {
                let _coll32 = new _dafny.Set();
                for (const _compr_32 of (_pat_let_tv19).Elements) {
                  let _1102_v73 = _compr_32;
                  if ((_pat_let_tv20).contains(_1102_v73)) {
                    _coll32.add(_module.__default.safeDivisionInt(_1102_v73, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality())));
                  }
                }
                return _coll32;
              }());
            }(_pat_let23_0);
          }(_module.D11.create_DC30(_dafny.Set.fromElements((_this).f17))),_1099_v75);
          let _1104_v77;
          _1104_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,new BigNumber(931));
          let _1105_v78;
          let _nw166 = new _module.C3();
          _nw166.__ctor(_1104_v77, _this.f9, (_this).f10, p0);
          _1105_v78 = _nw166;
          let _1106_v79;
          _1106_v79 = _dafny.Seq.of(_1105_v78, _1105_v78);
          let _1107_v80;
          _1107_v80 = _dafny.Map.Empty.slice().updateUnsafe((_1106_v79)[_module.__default.safeIndex((_this).f6, new BigNumber((_1106_v79).length))],false);
          let _1108_v81;
          _1108_v81 = _dafny.Set.fromElements(new BigNumber((_1107_v80).length));
          let _1109_v82;
          _1109_v82 = _module.D11.create_DC30(_1108_v81);
          _1100_v76 = (_1100_v76).update(_1109_v82, _1099_v75);
          let _rhs197 = (_1093_v70).update(_1055_v46, (p0).isLessThanOrEqualTo(p0));
          let _rhs198 = _1055_v46;
          let _rhs199 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_this).f18, (_this).f6), (_this).f6);
          let _lhs163 = globalState;
          let _lhs164 = globalState;
          _1093_v70 = _rhs197;
          _lhs163.f0 = _rhs198;
          _lhs164.f5 = _rhs199;
        }
        let _arr37 = _this.f9;
        let _index168 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_this.f9).length));
        _arr37[_index168] = _1055_v46;
        let _1110_v83;
        _1110_v83 = _dafny.Set.fromElements(_1055_v46);
        let _arr38 = _this.f9;
        let _index169 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_this.f9).length));
        _arr38[_index169] = (_1110_v83).IsProperSubsetOf(_1110_v83);
        let _arr39 = _this.f9;
        let _index170 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_this.f9).length));
        _arr39[_index170] = (_this.f9)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_this.f9).length))];
        let _arr40 = _this.f9;
        let _index171 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_this.f9).length));
        _arr40[_index171] = (_this.f9)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_this.f9).length))];
      }
      let _1111_v84;
      let _init29 = function (_1112_i10) {
        return _dafny.Seq.UnicodeFromString("rkntbsman");
      };
      let _nw167 = Array((new BigNumber(4)).toNumber());
      for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw167.length); _i0_29++) {
        _nw167[_i0_29] = _init29(new BigNumber(_i0_29));
      }
      _1111_v84 = _nw167;
      let _1113_v85;
      _1113_v85 = true;
      let _1114_v86;
      _1114_v86 = _module.D10.create_DC29(p0, _1111_v84, _1113_v85, (_this).f18);
      let _source14 = _1114_v86;
      if (_source14.is_DC29) {
        let _1115___mcc_h5 = (_source14).cf53;
        let _1116___mcc_h6 = (_source14).cf54;
        let _1117___mcc_h7 = (_source14).cf55;
        let _1118___mcc_h8 = (_source14).cf56;
        let _1119_cf56 = _1118___mcc_h8;
        let _1120_cf55 = _1117___mcc_h7;
        let _1121_cf54 = _1116___mcc_h6;
        let _1122_cf53 = _1115___mcc_h5;
        (globalState).f2 = _dafny.Seq.UnicodeFromString("na");
        let _arr41 = _this.f9;
        let _index172 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_this.f9).length));
        _arr41[_index172] = _1113_v85;
        let _1123_v87;
        _1123_v87 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1124_v88;
        _1124_v88 = _dafny.Set.fromElements(new BigNumber(754), (_this).f18);
        let _1125_v89;
        let _nw168 = Array((new BigNumber(21)).toNumber());
        _nw168[(_dafny.ZERO).toNumber()] = _1123_v87;
        _nw168[(_dafny.ONE).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(2)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(3)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(4)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(5)).toNumber()] = _module.__default.fm0((_this).f18, _1124_v88, _1113_v85, globalState);
        _nw168[(new BigNumber(6)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(7)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(8)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(9)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(10)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(11)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(12)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(13)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(14)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(15)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(16)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(17)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(18)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(19)).toNumber()] = _1123_v87;
        _nw168[(new BigNumber(20)).toNumber()] = _1123_v87;
        _1125_v89 = _nw168;
        let _1126_v90;
        _1126_v90 = _dafny.Seq.of(_1125_v89, _1125_v89, _1125_v89);
        let _arr42 = _this.f9;
        let _index173 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_this.f9).length));
        _arr42[_index173] = !(_dafny.Map.Empty.slice().updateUnsafe(_1125_v89,_1123_v87)).contains((_1126_v90)[_module.__default.safeIndex((_this).f18, new BigNumber((_1126_v90).length))]);
        let _1127_v91;
        _1127_v91 = _dafny.Seq.of((_this.f9)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_this.f9).length))]);
        let _1128_v92;
        _1128_v92 = _dafny.Seq.of(_1127_v91, _1127_v91);
        let _1129_v93;
        _1129_v93 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1128_v92, _1128_v92)).length)),(_this).f17);
        _1129_v93 = (_1129_v93).update(new BigNumber((_1124_v88).length), p0);
        let _1130_v94;
        _1130_v94 = _module.D0.create_DC0(_1127_v91, _1122_cf53);
        let _1131_v95;
        _1131_v95 = _module.D1.create_DC4(_1130_v94, _module.__default.safeModuloInt((_this).f6, (_this).f6), _dafny.Seq.UnicodeFromString("pceere"), _1120_cf55, p0);
        let _source15 = _1131_v95;
        if (_source15.is_DC4) {
          let _1132___mcc_h10 = (_source15).cf7;
          let _1133___mcc_h11 = (_source15).cf8;
          let _1134___mcc_h12 = (_source15).cf9;
          let _1135___mcc_h13 = (_source15).cf10;
          let _1136___mcc_h14 = (_source15).cf11;
          let _1137_cf11 = _1136___mcc_h14;
          let _1138_cf10 = _1135___mcc_h13;
          let _1139_cf9 = _1134___mcc_h12;
          let _1140_cf8 = _1133___mcc_h11;
          let _1141_cf7 = _1132___mcc_h10;
          let _1142_v96;
          _1142_v96 = _dafny.Map.Empty.slice().updateUnsafe(false,_1111_v84);
          let _1143_v97;
          let _nw169 = Array((new BigNumber(21)).toNumber());
          _nw169[(_dafny.ZERO).toNumber()] = _1121_cf54;
          _nw169[(_dafny.ONE).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(2)).toNumber()] = (((_1142_v96).contains((_this.f9)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_this.f9).length))])) ? ((_1142_v96).get((_this.f9)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_this.f9).length))])) : (_1111_v84));
          _nw169[(new BigNumber(3)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(4)).toNumber()] = _1121_cf54;
          _nw169[(new BigNumber(5)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(6)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(7)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(8)).toNumber()] = _1121_cf54;
          _nw169[(new BigNumber(9)).toNumber()] = _1121_cf54;
          _nw169[(new BigNumber(10)).toNumber()] = _1121_cf54;
          _nw169[(new BigNumber(11)).toNumber()] = _1121_cf54;
          _nw169[(new BigNumber(12)).toNumber()] = (_module.D10.create_DC29((_this).f6, _1111_v84, _1113_v85, _module.__default.fm2(new BigNumber(534), globalState))).dtor_cf54;
          _nw169[(new BigNumber(13)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(14)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(15)).toNumber()] = _1121_cf54;
          _nw169[(new BigNumber(16)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(17)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(18)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(19)).toNumber()] = _1111_v84;
          _nw169[(new BigNumber(20)).toNumber()] = _1111_v84;
          _1143_v97 = _nw169;
          let _index174 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_1143_v97).length));
          (_1143_v97)[_index174] = _1121_cf54;
          let _index175 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_1143_v97).length));
          (_1143_v97)[_index175] = _1121_cf54;
          let _arr43 = _this.f9;
          let _index176 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_this.f9).length));
          _arr43[_index176] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_1127_v91, _dafny.Seq.of(_1138_cf10, _1120_cf55)), _1128_v92);
          _1122_cf53 = _module.__default.fm2(new BigNumber(615), globalState);
          let _1144_v98;
          let _nw170 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Set.Empty);
          _1144_v98 = _nw170;
          let _1145_v99;
          let _nw171 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _1145_v99 = _nw171;
          let _rhs200 = _1144_v98;
          let _rhs201 = _1145_v99;
          let _rhs202 = _module.__default.fm25(_1137_cf11, !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-665)), ((_1146_cf11) => function (_1147_i11) {
            return _1146_cf11;
          })(_1137_cf11)), (_this).f6), globalState);
          _1144_v98 = _rhs200;
          _1145_v99 = _rhs201;
          r2 = _rhs202;
        } else {
          let _1148___mcc_h15 = (_source15).cf6;
          let _1149_cf6 = _1148___mcc_h15;
          let _1150_v100;
          let _nw172 = Array((new BigNumber(8)).toNumber()).fill(false);
          _1150_v100 = _nw172;
          let _1151_v101;
          let _nw173 = new _module.C0();
          _nw173.__ctor(_1150_v100, !(!((_this.f9)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_this.f9).length))])));
          _1151_v101 = _nw173;
          let _1152_v102;
          let _nw174 = new _module.C1();
          _nw174.__ctor((_this).f6);
          _1152_v102 = _nw174;
          let _1153_v103;
          _1153_v103 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(187),_1152_v102);
          let _1154_v104;
          _1154_v104 = _dafny.MultiSet.fromElements(_1119_cf56, new BigNumber((_1153_v103).length));
          _1154_v104 = (_1154_v104).update((_this).f18, _module.__default.abs((_this).f18));
          let _1155_v105;
          _1155_v105 = _dafny.Seq.UnicodeFromString("uyva");
          let _1156_v106;
          _1156_v106 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jxfruork"), _1155_v105, _1155_v105);
          let _1157_v108;
          _1157_v108 = _module.D0.create_DC1(_1151_v101.f20, _1119_cf56, _1123_v87);
          let _1158_v109;
          _1158_v109 = _dafny.Set.fromElements(_1113_v85, _1120_cf55);
          _1156_v106 = ((_module.__default.fm47(_1123_v87, function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(355), new BigNumber(690))) {
              let _1159_v107 = _compr_33;
              if (((new BigNumber(355)).isLessThanOrEqualTo(_1159_v107)) && ((_1159_v107).isLessThan(new BigNumber(690)))) {
                _coll33.push([_module.__default.safeModuloInt(_1159_v107, (_this).f6),_1113_v85]);
              }
            }
            return _coll33;
          }(), _1120_cf55, globalState)).Difference(_dafny.MultiSet.fromElements(_1155_v105))).Difference((_1156_v106).update(_module.__default.fm30(_1113_v85, (_1157_v108).dtor_cf4, globalState), _module.__default.abs(new BigNumber((_1158_v109).length))));
          _1123_v87 = _1123_v87;
        }
      } else {
        let _1160___mcc_h9 = (_source14).cf52;
        let _1161_cf52 = _1160___mcc_h9;
        let _1162_v110;
        let _nw175 = new _module.C0();
        _nw175.__ctor(_this.f9, _1113_v85);
        _1162_v110 = _nw175;
        (globalState).f2 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-170)), function (_1163_i12) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        });
        let _1164_v111;
        _1164_v111 = _dafny.Seq.of(_1162_v110.f20);
        let _1165_v112;
        _1165_v112 = _dafny.MultiSet.fromElements((_this).f18, _module.__default.fm2(new BigNumber((_1164_v111).length), globalState), (_this).f17, (_this).f18);
        let _1166_v113;
        _1166_v113 = _dafny.Map.Empty.slice().updateUnsafe(_1113_v85,(_1165_v112).Intersect(_1165_v112));
        let _1167_v114;
        _1167_v114 = _dafny.Seq.of((_this).f6, p0, (_this).f17, (_this).f17);
        _1166_v113 = (_1166_v113).update(_1162_v110.f20, _dafny.MultiSet.FromArray(_dafny.Seq.update(_1167_v114, _module.__default.safeIndex((_this).f18, new BigNumber((_1167_v114).length)), (_this).f17)));
        let _arr44 = _1162_v110.f19;
        let _index177 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1162_v110.f19).length));
        _arr44[_index177] = ((_1162_v110.f20) ? (_1113_v85) : (_1113_v85));
        let _1168_v115;
        let _init30 = function (_1169_i13) {
          return _module.__default.safeModuloInt(_1169_i13, (_this).f17);
        };
        let _nw176 = Array((_dafny.ONE).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw176.length); _i0_30++) {
          _nw176[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1168_v115 = _nw176;
        let _index178 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_1168_v115).length));
        (_1168_v115)[_index178] = new BigNumber(134);
        let _1170_v116;
        _1170_v116 = _dafny.Seq.UnicodeFromString("tl");
        let _1171_v117;
        _1171_v117 = _dafny.Seq.of(_1170_v116);
        let _arr45 = _1162_v110.f19;
        let _index179 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1162_v110.f19).length));
        let _index180 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_1168_v115).length));
        let _rhs203 = !(_1113_v85);
        let _rhs204 = _module.__default.safeDivisionInt(p0, new BigNumber(367));
        let _rhs205 = p0;
        let _rhs206 = (_module.D4.create_DC13(_1167_v114, false, p0)).dtor_cf26;
        let _rhs207 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f17, new BigNumber(((_1171_v117)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_1171_v117).length))]).length), (_this).f17, (_this).f17), _1167_v114)).length);
        let _lhs165 = _1162_v110.f19;
        let _lhs166 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1162_v110.f19).length));
        let _lhs167 = globalState;
        let _lhs168 = _1168_v115;
        let _lhs169 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_1168_v115).length));
        let _lhs170 = globalState;
        let _lhs171 = globalState;
        _lhs165[_lhs166] = _rhs203;
        _lhs167.f5 = _rhs204;
        _lhs168[_lhs169] = _rhs205;
        _lhs170.f0 = _rhs206;
        _lhs171.f5 = _rhs207;
      }
      let _1172_v118;
      _1172_v118 = new _dafny.CodePoint('i'.codePointAt(0));
      let _1173_v119;
      _1173_v119 = _dafny.Seq.of(_1172_v118);
      let _1174_i14;
      _1174_i14 = _dafny.ZERO;
      L6: {
        while (!(((_this).f6).isLessThan(_module.__default.safeDivisionInt(new BigNumber(981), _module.__default.fm2(new BigNumber((_1173_v119).length), globalState))))) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1174_i14)) {
              break L6;
            }
            _1174_i14 = (_1174_i14).plus(_dafny.ONE);
            (globalState).f5 = _module.__default.safeDivisionInt((_this).f17, _module.__default.safeDivisionInt(new BigNumber((_1173_v119).length), (_this).f18));
            let _1175_v121;
            _1175_v121 = _dafny.Seq.of(p0);
            let _1176_v122;
            _1176_v122 = _dafny.Set.fromElements(new BigNumber((_1175_v121).length), (_this).f6);
            let _1177_v123;
            let _nw177 = new _module.C2();
            _nw177.__ctor(p0, _this.f16, _this.f9, function () {
              let _coll34 = new _dafny.Map();
              for (const _compr_34 of (_1176_v122).Elements) {
                let _1178_v120 = _compr_34;
                if ((_1176_v122).contains(_1178_v120)) {
                  _coll34.push([_module.__default.safeModuloInt(_1178_v120, new BigNumber(138)),_1173_v119]);
                }
              }
              return _coll34;
            }());
            _1177_v123 = _nw177;
            _1176_v122 = ((!(_1113_v85)) ? ((_1176_v122).Difference(_1176_v122)) : ((_1176_v122).Union(_1176_v122)));
            let _1179_v124;
            _1179_v124 = _dafny.Map.Empty.slice().updateUnsafe(_1113_v85,_1113_v85);
            let _1180_v125;
            _1180_v125 = _dafny.MultiSet.fromElements(_1113_v85);
            let _1181_v126;
            _1181_v126 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
            let _rhs208 = ((_this).f6).isLessThan((p0).minus(p0));
            let _rhs209 = _dafny.Map.Empty.slice().updateUnsafe(_1113_v85,(_1177_v123).fm22(_1180_v125, _1113_v85, new BigNumber((_1181_v126).length), globalState));
            let _lhs172 = globalState;
            _lhs172.f0 = _rhs208;
            _1179_v124 = _rhs209;
          }
        }
      }
      let _1182_v127;
      _1182_v127 = _dafny.MultiSet.fromElements(_1172_v118);
      let _1183_v128;
      _1183_v128 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1182_v127);
      let _1184_v129;
      let _nw178 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _1184_v129 = _nw178;
      let _1185_v130;
      _1185_v130 = _dafny.MultiSet.fromElements(_1184_v129, _1184_v129, _1184_v129, _1184_v129);
      let _1186_v131;
      _1186_v131 = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(p0, new BigNumber((_1185_v130).cardinality()))).cardinality()));
      _1183_v128 = (_1183_v128).update(new BigNumber((_1186_v131).length), _1182_v127);
      let _1187_v132;
      let _nw179 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
      _1187_v132 = _nw179;
      let _1188_v133;
      _1188_v133 = _dafny.Seq.of(_1113_v85);
      let _1189_v134;
      _1189_v134 = _dafny.Seq.of(_1188_v133, _1188_v133, _dafny.Seq.of(_1113_v85));
      let _index181 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1187_v132).length));
      (_1187_v132)[_index181] = _1189_v134;
      let _index182 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1187_v132).length));
      (_1187_v132)[_index182] = _1189_v134;
      let _1190_v135;
      _1190_v135 = _module.D0.create_DC1(_1113_v85, p0, (_1173_v119)[_module.__default.safeIndex((_this).f6, new BigNumber((_1173_v119).length))]);
      let _1191_v136;
      _1191_v136 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,_1113_v85);
      let _1192_v137;
      _1192_v137 = _dafny.Map.Empty.slice().updateUnsafe(_1113_v85,(((_1190_v135).dtor_cf2) ? (_1191_v136) : (_1191_v136)));
      _1192_v137 = (_1192_v137).update(!(_1113_v85) || (_1113_v85), _1191_v136);
      r0 = _module.D2.create_DC8(true, !(((_this).f6).isLessThan(_module.__default.fm2(new BigNumber((_1173_v119).length), globalState))));
      r1 = _dafny.Seq.of(!(_module.__default.fm3(globalState)) || (false));
      r2 = _dafny.Seq.Concat(_1173_v119, _1173_v119);
      return [r0, r1, r2];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f6 = _dafny.ZERO;
      this._f15 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f15, f6) {
      let _this = this;
      (_this)._f15 = f15;
      (_this)._f6 = f6;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))))), _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0))))));
    };
    fm6(p0, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,_dafny.Seq.of((_this).f6))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC0(_dafny.Seq.of(false, (_module.D2.create_DC8(!(false), false)).dtor_cf20), new BigNumber((_dafny.Seq.UnicodeFromString("ig")).length))).dtor_cf1,_dafny.Seq.Create(_module.__default.abs(new BigNumber(174)), function (_1193_i0) {
        return (_this).f15;
      })))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,_dafny.Seq.of((_module.D1.create_DC4(_module.D0.create_DC0(_dafny.Seq.of(!(true), true), (_this).f6), (_this).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), function (_1194_i1) {
  return new _dafny.CodePoint('p'.codePointAt(0));
}), false, (_this).f15)).dtor_cf11, (_this).f6))).Merge(function () {
        let _coll35 = new _dafny.Map();
        for (const _compr_35 of _dafny.IntegerRange(new BigNumber(94), new BigNumber(657))) {
          let _1195_v0 = _compr_35;
          if (((new BigNumber(94)).isLessThanOrEqualTo(_1195_v0)) && ((_1195_v0).isLessThan(new BigNumber(657)))) {
            _coll35.push([(_1195_v0).plus((_this).f6),_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))]);
          }
        }
        return _coll35;
      }()));
    };
    m10(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _1196_v0;
      let _nw180 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
      _1196_v0 = _nw180;
      let _1197_v1;
      _1197_v1 = _module.D2.create_DC6(p1, _1196_v0);
      let _source16 = ((p0) ? (_module.D2.create_DC6((_1197_v1).dtor_cf13, _1196_v0)) : (_1197_v1));
      if (_source16.is_DC6) {
        let _1198___mcc_h0 = (_source16).cf13;
        let _1199___mcc_h1 = (_source16).cf14;
        let _1200_cf14 = _1199___mcc_h1;
        let _1201_cf13 = _1198___mcc_h0;
        let _1202_v2;
        _1202_v2 = _dafny.Seq.of(!(p1).isEqualTo(p1), false, p0);
        _1202_v2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(p0), _1202_v2), _1202_v2);
        (globalState).f0 = p0;
        r1 = p0;
        let _1203_v3;
        _1203_v3 = new _dafny.CodePoint('m'.codePointAt(0));
        let _1204_v4;
        _1204_v4 = _dafny.MultiSet.fromElements(_1203_v3, _1203_v3);
        let _1205_v5;
        _1205_v5 = _module.D1.create_DC3(_1204_v4);
        let _1206_v6;
        _1206_v6 = _dafny.Set.fromElements(_1205_v5);
        let _1207_v7;
        _1207_v7 = _dafny.Set.fromElements(_1206_v6, _1206_v6, _1206_v6, _1206_v6, _1206_v6);
        let _1208_v8;
        _1208_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1207_v7);
        _1208_v8 = (_1208_v8).update(p0, _1207_v7);
      } else if (_source16.is_DC7) {
        let _1209___mcc_h2 = (_source16).cf15;
        let _1210___mcc_h3 = (_source16).cf16;
        let _1211___mcc_h4 = (_source16).cf17;
        let _1212___mcc_h5 = (_source16).cf18;
        let _1213_cf18 = _1212___mcc_h5;
        let _1214_cf17 = _1211___mcc_h4;
        let _1215_cf16 = _1210___mcc_h3;
        let _1216_cf15 = _1209___mcc_h2;
        let _1217_v9;
        _1217_v9 = _dafny.Seq.UnicodeFromString("y");
        let _1218_v10;
        _1218_v10 = _dafny.MultiSet.fromElements(!_dafny.areEqual(_1217_v9, _1217_v9));
        _1218_v10 = (_1218_v10).Union(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(p0), _1216_cf15)));
        _1217_v9 = _dafny.Seq.Concat(_1217_v9, _1217_v9);
        let _1219_v11;
        _1219_v11 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1217_v9);
        let _1220_v12;
        _1220_v12 = new _dafny.CodePoint('i'.codePointAt(0));
        let _1221_v13;
        _1221_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1217_v9);
        _1219_v11 = (_1219_v11).update(_dafny.Seq.contains(_module.__default.fm18(p0, globalState), _1220_v12), (((_1221_v13).contains((_this).f6)) ? ((_1221_v13).get((_this).f6)) : (_1217_v9)));
        if (p0) {
          let _1222_v14;
          _1222_v14 = _dafny.MultiSet.fromElements(_1213_cf18, (_this).f15);
          let _1223_v15;
          _1223_v15 = _dafny.Map.Empty.slice().updateUnsafe((_1222_v14).IsProperSubsetOf(_module.__default.fm19(p1, globalState)),(_this).f15);
          _1223_v15 = (_1223_v15).update(_dafny.Seq.IsPrefixOf(_1217_v9, _dafny.Seq.update(_1217_v9, _module.__default.safeIndex(new BigNumber((_1217_v9).length), new BigNumber((_1217_v9).length)), new _dafny.CodePoint('x'.codePointAt(0)))), (_dafny.ZERO).minus((_this).f15));
          let _1224_v16;
          _1224_v16 = _dafny.MultiSet.fromElements(_1220_v12);
          let _1225_v17;
          _1225_v17 = _dafny.Map.Empty.slice().updateUnsafe((((_1224_v16).contains(_1220_v12)) ? ((_1224_v16).get(_1220_v12)) : ((_this).f6)),_dafny.areEqual(_dafny.Seq.UnicodeFromString("dgs"), _1217_v9));
          let _1226_v18;
          _1226_v18 = _dafny.Set.fromElements(p0, p0);
          let _index183 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_1196_v0).length));
          (_1196_v0)[_index183] = new BigNumber(((_1226_v18).Union(_1226_v18)).length);
          let _1227_v19;
          _1227_v19 = _dafny.Map.Empty.slice().updateUnsafe(false,_1220_v12);
          let _1228_v20;
          _1228_v20 = _dafny.Set.fromElements(new BigNumber(58), (_this).f6);
          let _1229_v21;
          _1229_v21 = _module.D0.create_DC1(!(!_dafny.Seq.contains(_1217_v9, (((_1227_v19).contains(p0)) ? ((_1227_v19).get(p0)) : (_module.__default.fm0((_this).f15, _1228_v20, !(p0), globalState))))), (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_1213_cf18))), _1220_v12);
          let _1230_v22;
          let _nw181 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1230_v22 = _nw181;
          let _index184 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_1196_v0).length));
          let _rhs210 = _1225_v17;
          let _rhs211 = p1;
          let _rhs212 = _module.__default.fm20(p0, (p0) || (p0), p0, globalState);
          let _rhs213 = _1230_v22;
          let _lhs173 = _1196_v0;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_1196_v0).length));
          _1225_v17 = _rhs210;
          _lhs173[_lhs174] = _rhs211;
          _1229_v21 = _rhs212;
          _1230_v22 = _rhs213;
          r1 = _module.__default.fm3(globalState);
          _1213_cf18 = (((_1196_v0)[_module.__default.safeIndex(new BigNumber(799), new BigNumber((_1196_v0).length))]).multipliedBy((_this).f6)).multipliedBy(_1213_cf18);
          let _1231_v23;
          let _nw182 = Array((new BigNumber(11)).toNumber());
          _nw182[(_dafny.ZERO).toNumber()] = _1216_cf15;
          _nw182[(_dafny.ONE).toNumber()] = _1216_cf15;
          _nw182[(new BigNumber(2)).toNumber()] = _1216_cf15;
          _nw182[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(p0);
          _nw182[(new BigNumber(4)).toNumber()] = _module.__default.fm4((_this).f6, p1, _1213_cf18, globalState);
          _nw182[(new BigNumber(5)).toNumber()] = _1216_cf15;
          _nw182[(new BigNumber(6)).toNumber()] = _1216_cf15;
          _nw182[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(p0);
          _nw182[(new BigNumber(8)).toNumber()] = _1216_cf15;
          _nw182[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(p0, p0);
          _nw182[(new BigNumber(10)).toNumber()] = _1216_cf15;
          _1231_v23 = _nw182;
          let _1232_v24;
          _1232_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_1228_v20).length), new BigNumber((_1217_v9).length)),_1231_v23);
          _1232_v24 = (_1232_v24).update((_dafny.ZERO).minus((_this).f6), _1231_v23);
        } else {
          let _1233_v25;
          let _nw183 = new _module.C1();
          _nw183.__ctor(new BigNumber(-736));
          _1233_v25 = _nw183;
          let _index185 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1196_v0).length));
          (_1196_v0)[_index185] = _1214_cf17;
          let _index186 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1196_v0).length));
          (_1196_v0)[_index186] = _1214_cf17;
          _1220_v12 = _1220_v12;
          (globalState).f0 = p0;
          r1 = false;
        }
      } else if (_source16.is_DC8) {
        let _1234___mcc_h6 = (_source16).cf19;
        let _1235___mcc_h7 = (_source16).cf20;
        let _1236_cf20 = _1235___mcc_h7;
        let _1237_cf19 = _1234___mcc_h6;
        let _1238_v26;
        let _nw184 = Array((new BigNumber(5)).toNumber());
        _nw184[(_dafny.ZERO).toNumber()] = _1237_cf19;
        _nw184[(_dafny.ONE).toNumber()] = false;
        _nw184[(new BigNumber(2)).toNumber()] = _1237_cf19;
        _nw184[(new BigNumber(3)).toNumber()] = false;
        _nw184[(new BigNumber(4)).toNumber()] = p0;
        _1238_v26 = _nw184;
        let _1239_v27;
        _1239_v27 = _dafny.Seq.UnicodeFromString("dusfea");
        let _1240_v28;
        _1240_v28 = _dafny.Seq.of(_1239_v27, _1239_v27);
        let _1241_v29;
        _1241_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(622),(_1240_v28)[_module.__default.safeIndex(p1, new BigNumber((_1240_v28).length))]);
        let _1242_v30;
        let _nw185 = new _module.C3();
        _nw185.__ctor(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wbwsr")).length),p1), ((_1236_cf20) ? (_1238_v26) : (_1238_v26)), _1241_v29, (_this).f6);
        _1242_v30 = _nw185;
        let _1243_v31;
        let _nw186 = new _module.C1();
        _nw186.__ctor(p1);
        _1243_v31 = _nw186;
        let _1244_v32;
        _1244_v32 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(427));
        let _1245_v33;
        _1245_v33 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1246_v34;
        _1246_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1245_v33,p0);
        r0 = (_1244_v32).update(p0, new BigNumber((_1246_v34).length));
        let _1247_v35;
        _1247_v35 = _dafny.Seq.of(_1237_cf19, _1237_cf19);
        let _1248_v36;
        _1248_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1247_v35);
        let _1249_v37;
        _1249_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_1236_cf20);
        let _1250_v38;
        _1250_v38 = _dafny.Seq.of(_1249_v37);
        let _1251_v39;
        _1251_v39 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(((((_1248_v36).contains(!(true))) ? ((_1248_v36).get(!(true))) : (_dafny.Seq.update(_1247_v35, _module.__default.safeIndex((_this).f15, new BigNumber((_1247_v35).length)), p0)))).length)).isEqualTo(new BigNumber(87)),!_dafny.areEqual(_1250_v38, _1250_v38));
        let _rhs214 = _1251_v39;
        let _rhs215 = _1196_v0;
        let _rhs216 = p1;
        let _lhs175 = globalState;
        _1251_v39 = _rhs214;
        _1196_v0 = _rhs215;
        _lhs175.f5 = _rhs216;
      } else {
        let _1252___mcc_h8 = (_source16).cf12;
        let _1253_cf12 = _1252___mcc_h8;
        let _1254_v40;
        _1254_v40 = _dafny.Set.fromElements(p0);
        _1254_v40 = _1254_v40;
        let _1255_v41;
        _1255_v41 = _dafny.MultiSet.fromElements(p1);
        let _1256_v42;
        _1256_v42 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1257_v43;
        _1257_v43 = _dafny.Set.fromElements(_module.__default.fm42(p0, _1255_v41, p1, _1256_v42, globalState));
        let _1258_v44;
        _1258_v44 = _dafny.Seq.UnicodeFromString("s");
        let _1259_v45;
        _1259_v45 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f15, new BigNumber((_1258_v44).length)));
        let _1260_v46;
        _1260_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1257_v43,(((_1259_v45).contains(_dafny.Seq.update(_dafny.Seq.of(_module.__default.fm2(new BigNumber(-138), globalState)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.of(_module.__default.fm2(new BigNumber(-138), globalState))).length)), (_this).f6))) ? ((_1259_v45).get(_dafny.Seq.update(_dafny.Seq.of(_module.__default.fm2(new BigNumber(-138), globalState)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.of(_module.__default.fm2(new BigNumber(-138), globalState))).length)), (_this).f6))) : (new BigNumber((_1258_v44).length))));
        _1260_v46 = (_1260_v46).update((_1257_v43).Intersect(function () {
          let _coll36 = new _dafny.Set();
          for (const _compr_36 of (_dafny.Map.Empty.slice().updateUnsafe(_1253_cf12,p1)).Keys.Elements) {
            let _1261_v47 = _compr_36;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_1253_cf12,p1)).contains(_1261_v47)) {
              _coll36.add(_1261_v47);
            }
          }
          return _coll36;
        }()), (_this).f6);
        let _1262_v48;
        _1262_v48 = _dafny.Seq.of(p0);
        let _1263_v49;
        _1263_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1262_v48,p0);
        let _1264_v50;
        _1264_v50 = _dafny.Set.fromElements(_1196_v0);
        let _1265_v51;
        _1265_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1263_v49,_1264_v50);
        _1265_v51 = (_1265_v51).update(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0),_module.__default.fm3(globalState)), _1264_v50);
        let _1266_v52;
        _1266_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,((_this).f15).isEqualTo(p1));
        _1266_v52 = (_1266_v52).update(p0, p0);
      }
      let _1267_v53;
      let _nw187 = Array((new BigNumber(2)).toNumber()).fill(false);
      _1267_v53 = _nw187;
      let _1268_v54;
      _1268_v54 = _1267_v53;
      let _1269_v55;
      _1269_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1268_v54,p0);
      let _index187 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_1196_v0).length));
      (_1196_v0)[_index187] = new BigNumber((_1269_v55).length);
      let _index188 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_1196_v0).length));
      (_1196_v0)[_index188] = p1;
      let _1270_v56;
      _1270_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1196_v0,((_1196_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_1196_v0).length))]).multipliedBy((_1196_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_1196_v0).length))]));
      let _1271_v57;
      _1271_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1196_v0);
      _1270_v56 = (_1270_v56).update((((_1271_v57).contains(false)) ? ((_1271_v57).get(false)) : (_1196_v0)), (_this).f15);
      let _1272_v58;
      _1272_v58 = new _dafny.CodePoint('n'.codePointAt(0));
      let _1273_v59;
      _1273_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1272_v58,_1267_v53);
      _1273_v59 = (_1273_v59).update(_1272_v58, _1267_v53);
      let _1274_v60;
      _1274_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f15);
      let _1275_v61;
      _1275_v61 = _dafny.Seq.of(new BigNumber((_1274_v60).length), (_this).f15);
      let _1276_v62;
      _1276_v62 = _dafny.Seq.of(p0);
      let _1277_v63;
      _1277_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_1276_v62)[_module.__default.safeIndex(new BigNumber((_module.__default.fm1(globalState)).length), new BigNumber((_1276_v62).length))]);
      let _1278_v64;
      _1278_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1277_v63,p0);
      r1 = ((p0) ? ((new BigNumber((_1275_v61).length)).isLessThan(new BigNumber((_module.__default.fm33(_module.D3.create_DC10(_1278_v64), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), function (_1279_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length), !(p0), (((_1277_v63).contains((_1196_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_1196_v0).length))])) ? ((_1277_v63).get((_1196_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_1196_v0).length))])) : (false)), globalState)).length))) : (p0));
      let _1280_v65;
      _1280_v65 = _dafny.Set.fromElements(false, p0);
      let _1281_v66;
      _1281_v66 = _module.D11.create_DC33((_this).f15, p0, (_this).f15, p0, _1280_v65);
      let _1282_v67;
      _1282_v67 = _dafny.MultiSet.fromElements((_1281_v66).dtor_cf64, (_this).f15, new BigNumber(112), (_this).f6, new BigNumber(477));
      let _index189 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_1196_v0).length));
      (_1196_v0)[_index189] = ((p1).minus(new BigNumber((_dafny.MultiSet.fromElements((_1196_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_1196_v0).length))])).cardinality()))).multipliedBy(new BigNumber((_module.__default.fm42(p0, _1282_v67, new BigNumber((_1282_v67).cardinality()), _1272_v58, globalState)).length));
      r0 = (_module.D14.create_DC36(_1274_v60)).dtor_cf69;
      r1 = p0;
      return [r0, r1];
    }
    m11(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _1283_v0;
      _1283_v0 = _dafny.Seq.UnicodeFromString("sc");
      let _1284_v1;
      _1284_v1 = _dafny.Seq.of((_this).f15);
      let _1285_v2;
      _1285_v2 = new _dafny.CodePoint('o'.codePointAt(0));
      let _1286_v3;
      _1286_v3 = _module.D9.create_DC25((_this).f6, (_this).f6, (_this).f15, _1285_v2);
      let _1287_v4;
      _1287_v4 = _module.D9.create_DC27(_1286_v3);
      let _1288_v5;
      _1288_v5 = _module.D9.create_DC27(_1287_v4);
      let _1289_v6;
      let _nw188 = Array((new BigNumber(6)).toNumber());
      _nw188[(_dafny.ZERO).toNumber()] = _module.D9.create_DC27(_module.D9.create_DC25(new BigNumber((_1283_v0).length), new BigNumber((_1284_v1).length), (_this).f6, _1285_v2));
      _nw188[(_dafny.ONE).toNumber()] = _1288_v5;
      _nw188[(new BigNumber(2)).toNumber()] = _1288_v5;
      _nw188[(new BigNumber(3)).toNumber()] = _1288_v5;
      _nw188[(new BigNumber(4)).toNumber()] = _1288_v5;
      _nw188[(new BigNumber(5)).toNumber()] = _1288_v5;
      _1289_v6 = _nw188;
      let _index190 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1289_v6).length));
      (_1289_v6)[_index190] = _1288_v5;
      let _1290_v7;
      _1290_v7 = _dafny.MultiSet.fromElements(true);
      let _1291_v8;
      _1291_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1290_v7);
      let _1292_v9;
      _1292_v9 = _dafny.Seq.of(((p1) ? (_1290_v7) : ((((_1291_v8).contains(p1)) ? ((_1291_v8).get(p1)) : (_1290_v7)))));
      let _index191 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1289_v6).length));
      let _rhs217 = ((_this).f6).plus(new BigNumber(10));
      let _rhs218 = _1288_v5;
      let _rhs219 = (_1292_v9)[_module.__default.safeIndex(((_this).f15).plus((_1284_v1)[_module.__default.safeIndex((_this).f6, new BigNumber((_1284_v1).length))]), new BigNumber((_1292_v9).length))];
      let _lhs176 = _1289_v6;
      let _lhs177 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1289_v6).length));
      r3 = _rhs217;
      _lhs176[_lhs177] = _rhs218;
      _1290_v7 = _rhs219;
      let _hi9 = (_1284_v1)[_module.__default.safeIndex((_this).f15, new BigNumber((_1284_v1).length))];
      for (let _1293_i0 = new BigNumber(417); _1293_i0.isLessThan(_hi9); _1293_i0 = _1293_i0.plus(_dafny.ONE)) {
        r0 = ((_dafny.ZERO).minus((_this).f15)).minus((((_1290_v7).contains(p0)) ? ((_1290_v7).get(p0)) : ((_this).f6)));
        let _1294_v10;
        _1294_v10 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(863)).isEqualTo((_dafny.ZERO).minus((_this).f15)),p1);
        _1294_v10 = (_1294_v10).update(p1, _module.__default.fm3(globalState));
        let _1295_v11;
        _1295_v11 = _dafny.MultiSet.fromElements(new BigNumber((_1283_v0).length));
        let _1296_v12;
        _1296_v12 = _dafny.MultiSet.fromElements(_1295_v11, _1295_v11);
        let _1297_v13;
        _1297_v13 = _module.D4.create_DC13(_1284_v1, p0, (_1284_v1)[_module.__default.safeIndex((_this).f15, new BigNumber((_1284_v1).length))]);
        let _1298_v14;
        let _nw189 = Array((new BigNumber(18)).toNumber());
        _nw189[(_dafny.ZERO).toNumber()] = p1;
        _nw189[(_dafny.ONE).toNumber()] = p1;
        _nw189[(new BigNumber(2)).toNumber()] = p0;
        _nw189[(new BigNumber(3)).toNumber()] = p0;
        _nw189[(new BigNumber(4)).toNumber()] = p0;
        _nw189[(new BigNumber(5)).toNumber()] = p0;
        _nw189[(new BigNumber(6)).toNumber()] = p0;
        _nw189[(new BigNumber(7)).toNumber()] = (_1297_v13).dtor_cf26;
        _nw189[(new BigNumber(8)).toNumber()] = p1;
        _nw189[(new BigNumber(9)).toNumber()] = false;
        _nw189[(new BigNumber(10)).toNumber()] = p1;
        _nw189[(new BigNumber(11)).toNumber()] = p0;
        _nw189[(new BigNumber(12)).toNumber()] = true;
        _nw189[(new BigNumber(13)).toNumber()] = p1;
        _nw189[(new BigNumber(14)).toNumber()] = p1;
        _nw189[(new BigNumber(15)).toNumber()] = p1;
        _nw189[(new BigNumber(16)).toNumber()] = p1;
        _nw189[(new BigNumber(17)).toNumber()] = !(p0);
        _1298_v14 = _nw189;
        let _1299_v15;
        _1299_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1285_v2,_1298_v14);
        let _1300_v16;
        let _nw190 = new _module.C2();
        _nw190.__ctor((_this).f6, _1296_v12, (((_1299_v15).contains(_1285_v2)) ? ((_1299_v15).get(_1285_v2)) : (_1298_v14)), _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_module.__default.fm18(p1, globalState)));
        _1300_v16 = _nw190;
        (globalState).f0 = false;
      }
      let _source17 = _module.D5.create_DC17(_1284_v1, p0);
      if (_source17.is_DC17) {
        let _1301___mcc_h0 = (_source17).cf36;
        let _1302___mcc_h1 = (_source17).cf37;
        let _1303_cf37 = _1302___mcc_h1;
        let _1304_cf36 = _1301___mcc_h0;
        let _1305_v17;
        let _init31 = ((_1306_p1) => function (_1307_i1) {
          return ((_1306_p1) ? (true) : (_1306_p1));
        })(p1);
        let _nw191 = Array((new BigNumber(21)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw191.length); _i0_31++) {
          _nw191[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1305_v17 = _nw191;
        let _index192 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1305_v17).length));
        (_1305_v17)[_index192] = !(p0);
        let _index193 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1305_v17).length));
        (_1305_v17)[_index193] = p0;
        let _1308_v18;
        _1308_v18 = _dafny.Seq.of(_1303_cf37);
        let _1309_v19;
        _1309_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_1305_v17)[_module.__default.safeIndex(new BigNumber(166), new BigNumber((_1305_v17).length))]);
        let _1310_v20;
        _1310_v20 = _dafny.Set.fromElements(new BigNumber((_1309_v19).length), new BigNumber(667), (_this).f6);
        let _1311_v21;
        _1311_v21 = _dafny.Seq.of(_1310_v20, _1310_v20);
        let _1312_v22;
        let _nw192 = Array((new BigNumber(22)).toNumber());
        _nw192[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1303_cf37,true)).length);
        _nw192[(_dafny.ONE).toNumber()] = (_this).f15;
        _nw192[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.update(_1308_v18, _module.__default.safeIndex(new BigNumber((_1304_cf36).length), new BigNumber((_1308_v18).length)), p0)).length);
        _nw192[(new BigNumber(3)).toNumber()] = (_this).f15;
        _nw192[(new BigNumber(4)).toNumber()] = new BigNumber(((_1311_v21)[_module.__default.safeIndex((_this).f6, new BigNumber((_1311_v21).length))]).length);
        _nw192[(new BigNumber(5)).toNumber()] = new BigNumber((_1283_v0).length);
        _nw192[(new BigNumber(6)).toNumber()] = (_this).f6;
        _nw192[(new BigNumber(7)).toNumber()] = _module.__default.fm2((_this).f6, globalState);
        _nw192[(new BigNumber(8)).toNumber()] = (_this).f15;
        _nw192[(new BigNumber(9)).toNumber()] = (_this).f6;
        _nw192[(new BigNumber(10)).toNumber()] = (_this).f6;
        _nw192[(new BigNumber(11)).toNumber()] = (_this).f15;
        _nw192[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_this).f15);
        _nw192[(new BigNumber(13)).toNumber()] = (_this).f6;
        _nw192[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_this).f15);
        _nw192[(new BigNumber(15)).toNumber()] = (_1284_v1)[_module.__default.safeIndex((_this).f15, new BigNumber((_1284_v1).length))];
        _nw192[(new BigNumber(16)).toNumber()] = _module.__default.fm2((_this).f6, globalState);
        _nw192[(new BigNumber(17)).toNumber()] = new BigNumber(697);
        _nw192[(new BigNumber(18)).toNumber()] = (_this).f15;
        _nw192[(new BigNumber(19)).toNumber()] = (_this).f15;
        _nw192[(new BigNumber(20)).toNumber()] = new BigNumber(576);
        _nw192[(new BigNumber(21)).toNumber()] = (_this).f15;
        _1312_v22 = _nw192;
        let _1313_v23;
        _1313_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1312_v22,(_1305_v17)[_module.__default.safeIndex(new BigNumber(166), new BigNumber((_1305_v17).length))]);
        _1303_cf37 = !((((_1313_v23).contains(_1312_v22)) ? ((_1313_v23).get(_1312_v22)) : (p0)));
        r2 = _dafny.Seq.IsPrefixOf(_1283_v0, _1283_v0);
        _1309_v19 = (_1309_v19).update((_this).f6, (_1305_v17)[_module.__default.safeIndex(new BigNumber(166), new BigNumber((_1305_v17).length))]);
      } else {
        let _1314___mcc_h2 = (_source17).cf35;
        let _1315_cf35 = _1314___mcc_h2;
        let _1316_v24;
        let _nw193 = Array((new BigNumber(17)).toNumber()).fill(false);
        _1316_v24 = _nw193;
        _1316_v24 = _1316_v24;
        (globalState).f5 = (_this).f15;
        let _1317_v25;
        _1317_v25 = _dafny.Seq.of(p1);
        let _1318_v26;
        _1318_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1317_v25,(_this).f15);
        _1318_v26 = (_1318_v26).update(_1317_v25, new BigNumber(-172));
        r2 = !(!(p1)) || (((_this).f15).isEqualTo((_this).f15));
      }
      let _1319_i2;
      _1319_i2 = _dafny.ZERO;
      L7: {
        while (p0) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1319_i2)) {
              break L7;
            }
            _1319_i2 = (_1319_i2).plus(_dafny.ONE);
            let _1320_v27;
            _1320_v27 = _dafny.Seq.of(((p0) ? (p1) : (p1)));
            let _1321_v28;
            _1321_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_1320_v27);
            _1320_v27 = _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm3(globalState)), (((_1321_v28).contains((_this).f6)) ? ((_1321_v28).get((_this).f6)) : (_1320_v27)));
            let _1322_v29;
            let _init32 = function (_1323_i3) {
              return _module.__default.safeModuloInt(_1323_i3, (_this).f15);
            };
            let _nw194 = Array((new BigNumber(14)).toNumber());
            for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw194.length); _i0_32++) {
              _nw194[_i0_32] = _init32(new BigNumber(_i0_32));
            }
            _1322_v29 = _nw194;
            let _1324_v30;
            _1324_v30 = _module.D2.create_DC6((_this).f6, _1322_v29);
            let _1325_v31;
            _1325_v31 = _dafny.Seq.of(_1322_v29, _1322_v29);
            let _pat_let_tv21 = _1322_v29;
            let _pat_let_tv22 = _1322_v29;
            let _1326_v32;
            let _nw195 = Array((new BigNumber(11)).toNumber());
            _nw195[(_dafny.ZERO).toNumber()] = _1322_v29;
            _nw195[(_dafny.ONE).toNumber()] = _1322_v29;
            _nw195[(new BigNumber(2)).toNumber()] = _1322_v29;
            _nw195[(new BigNumber(3)).toNumber()] = _1322_v29;
            _nw195[(new BigNumber(4)).toNumber()] = _1322_v29;
            _nw195[(new BigNumber(5)).toNumber()] = (function (_pat_let25_0) {
              return function (_1329_dt__update__tmp_h1) {
                return function (_pat_let28_0) {
                  return function (_1330_dt__update_hcf14_h1) {
                    return _module.D2.create_DC6((_1329_dt__update__tmp_h1).dtor_cf13, _1330_dt__update_hcf14_h1);
                  }(_pat_let28_0);
                }(_pat_let_tv22);
              }(_pat_let25_0);
            }(function (_pat_let26_0) {
              return function (_1327_dt__update__tmp_h0) {
                return function (_pat_let27_0) {
                  return function (_1328_dt__update_hcf14_h0) {
                    return _module.D2.create_DC6((_1327_dt__update__tmp_h0).dtor_cf13, _1328_dt__update_hcf14_h0);
                  }(_pat_let27_0);
                }(_pat_let_tv21);
              }(_pat_let26_0);
            }(_1324_v30))).dtor_cf14;
            _nw195[(new BigNumber(6)).toNumber()] = (_1325_v31)[_module.__default.safeIndex(new BigNumber(((_1290_v7).update(p0, _module.__default.abs((_this).f15))).cardinality()), new BigNumber((_1325_v31).length))];
            _nw195[(new BigNumber(7)).toNumber()] = _1322_v29;
            _nw195[(new BigNumber(8)).toNumber()] = _1322_v29;
            _nw195[(new BigNumber(9)).toNumber()] = _1322_v29;
            _nw195[(new BigNumber(10)).toNumber()] = _1322_v29;
            _1326_v32 = _nw195;
            let _index194 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1326_v32).length));
            (_1326_v32)[_index194] = _1322_v29;
            let _index195 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1326_v32).length));
            (_1326_v32)[_index195] = _1322_v29;
            (globalState).f5 = (_this).f15;
            let _1331_v33;
            let _nw196 = new _module.C1();
            _nw196.__ctor((_this).f15);
            _1331_v33 = _nw196;
          }
        }
      }
      let _1332_v34;
      _1332_v34 = _dafny.Seq.of(p0);
      let _1333_v35;
      _1333_v35 = _dafny.Seq.of((_1332_v34)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(p1)).length)), new BigNumber((_1332_v34).length))], false);
      let _1334_v37;
      _1334_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll37 = new _dafny.Map();
        for (const _compr_37 of (_1284_v1).Elements) {
          let _1335_v36 = _compr_37;
          if (_dafny.Seq.contains(_1284_v1, _1335_v36)) {
            _coll37.push([(_1335_v36).multipliedBy((_dafny.ZERO).minus((_this).f15)),(_this).f6]);
          }
        }
        return _coll37;
      }()).length),p0);
      let _1336_v38;
      _1336_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1334_v37,p0);
      let _1337_v39;
      _1337_v39 = _module.D3.create_DC10(_1336_v38);
      let _pat_let_tv23 = _1336_v38;
      let _1338_v40;
      let _nw197 = Array((new BigNumber(3)).toNumber());
      _nw197[(_dafny.ZERO).toNumber()] = (((_1332_v34)[_module.__default.safeIndex((_this).f15, new BigNumber((_1332_v34).length))]) ? (_1337_v39) : (function (_pat_let29_0) {
        return function (_1339_dt__update__tmp_h2) {
          return function (_pat_let30_0) {
            return function (_1340_dt__update_hcf22_h0) {
              return _module.D3.create_DC10(_1340_dt__update_hcf22_h0);
            }(_pat_let30_0);
          }(_pat_let_tv23);
        }(_pat_let29_0);
      }(_1337_v39)));
      _nw197[(_dafny.ONE).toNumber()] = _module.D3.create_DC10(_1336_v38);
      _nw197[(new BigNumber(2)).toNumber()] = _1337_v39;
      _1338_v40 = _nw197;
      let _index196 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1338_v40).length));
      (_1338_v40)[_index196] = _1337_v39;
      let _index197 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1338_v40).length));
      (_1338_v40)[_index197] = (((p1) === (p1)) ? (_1337_v39) : (_1337_v39));
      if (p1) {
        if (true) {
          (globalState).f2 = _module.__default.fm30(p1, _1285_v2, globalState);
          let _1341_v41;
          let _init33 = ((_1342_p1) => function (_1343_i4) {
            return !(_1342_p1);
          })(p1);
          let _nw198 = Array((new BigNumber(10)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw198.length); _i0_33++) {
            _nw198[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1341_v41 = _nw198;
          let _1344_v42;
          _1344_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_dafny.Seq.UnicodeFromString("amxb"));
          let _1345_v43;
          _1345_v43 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.update(_1284_v1, _module.__default.safeIndex((_this).f6, new BigNumber((_1284_v1).length)), new BigNumber(680))));
          let _1346_v44;
          let _nw199 = new _module.C4();
          _nw199.__ctor(_module.__default.fm2(_module.__default.fm2((_this).f6, globalState), globalState), (_this).f15, ((p0) ? (_1341_v41) : (_1341_v41)), (_1344_v42).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), ((_1347_v2) => function (_1348_i5) {
            return _1347_v2;
          })(_1285_v2))).length),_1283_v0)).update(_module.__default.fm2(new BigNumber((_1290_v7).cardinality()), globalState), _dafny.Seq.UnicodeFromString("coc"))), (_this).f6, _1345_v43);
          _1346_v44 = _nw199;
          (globalState).f5 = _module.__default.safeModuloInt(new BigNumber((_1284_v1).length), (_this).f6);
          let _1349_v45;
          _1349_v45 = _module.D0.create_DC1(p1, (_1346_v44).f18, _1285_v2);
          let _1350_v46;
          _1350_v46 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
          let _1351_v47;
          _1351_v47 = _module.D4.create_DC14(p1, _1350_v46, (_this).f15, _1350_v46);
          let _1352_v48;
          _1352_v48 = _dafny.Seq.of(_1351_v47, _module.D4.create_DC14(p1, _1350_v46, (_this).f6, _1350_v46));
          let _1353_v49;
          _1353_v49 = _dafny.Seq.of(_1352_v48, _dafny.Seq.Create(_module.__default.abs(new BigNumber(710)), ((_1354_v47) => function (_1355_i6) {
            return _1354_v47;
          })(_1351_v47)), _1352_v48, _dafny.Seq.update(_dafny.Seq.update(_1352_v48, _module.__default.safeIndex((_this).f6, new BigNumber((_1352_v48).length)), _1351_v47), _module.__default.safeIndex((_this).f6, new BigNumber((_dafny.Seq.update(_1352_v48, _module.__default.safeIndex((_this).f6, new BigNumber((_1352_v48).length)), _1351_v47)).length)), _module.D4.create_DC14(p1, _1350_v46, (_dafny.ZERO).minus((_1346_v44).f17), _dafny.Map.Empty.slice().updateUnsafe(p0,p0))), _1352_v48);
          let _rhs220 = (_1349_v45).dtor_cf2;
          let _rhs221 = (new BigNumber((_1353_v49).length)).multipliedBy((_this).f6);
          let _rhs222 = (!_dafny.areEqual(_dafny.Seq.UnicodeFromString("kx"), _1283_v0)) || (p0);
          let _rhs223 = new _dafny.CodePoint('t'.codePointAt(0));
          let _lhs178 = globalState;
          let _lhs179 = globalState;
          let _lhs180 = globalState;
          _lhs178.f0 = _rhs220;
          _lhs179.f5 = _rhs221;
          _lhs180.f0 = _rhs222;
          _1285_v2 = _rhs223;
          let _1356_v50;
          _1356_v50 = _dafny.Set.fromElements(p0);
          let _1357_v51;
          let _nw200 = new _module.C1();
          _nw200.__ctor(_module.__default.safeModuloInt(new BigNumber(548), (_1346_v44).fm21(new BigNumber((_1356_v50).length), globalState)));
          _1357_v51 = _nw200;
        } else {
          (globalState).f0 = p0;
          let _1358_v52;
          _1358_v52 = _dafny.MultiSet.fromElements(new BigNumber(28), (_this).f15);
          let _1359_v53;
          _1359_v53 = _dafny.MultiSet.fromElements(_1358_v52);
          let _1360_v54;
          let _init34 = function (_1361_i7) {
            return true;
          };
          let _nw201 = Array((new BigNumber(28)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw201.length); _i0_34++) {
            _nw201[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1360_v54 = _nw201;
          let _1362_v55;
          _1362_v55 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f6),_1283_v0);
          let _1363_v57;
          let _nw202 = new _module.C2();
          _nw202.__ctor(new BigNumber(-676), (_1359_v53).Union(_dafny.MultiSet.fromElements(_1358_v52, _1358_v52)), _1360_v54, (_1362_v55).Merge(function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of _dafny.IntegerRange(new BigNumber(-58), new BigNumber(555))) {
              let _1364_v56 = _compr_38;
              if (((new BigNumber(-58)).isLessThanOrEqualTo(_1364_v56)) && ((_1364_v56).isLessThan(new BigNumber(555)))) {
                _coll38.push([_module.__default.safeModuloInt(_1364_v56, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_1358_v52).cardinality()))).length)),_1283_v0]);
              }
            }
            return _coll38;
          }()));
          _1363_v57 = _nw202;
          let _1365_v59;
          let _nw203 = Array((new BigNumber(26)).toNumber());
          _nw203[(_dafny.ZERO).toNumber()] = new BigNumber((_1290_v7).cardinality());
          _nw203[(_dafny.ONE).toNumber()] = (_this).f6;
          _nw203[(new BigNumber(2)).toNumber()] = (_this).f15;
          _nw203[(new BigNumber(3)).toNumber()] = (_this).f6;
          _nw203[(new BigNumber(4)).toNumber()] = (_this).f15;
          _nw203[(new BigNumber(5)).toNumber()] = (_this).f15;
          _nw203[(new BigNumber(6)).toNumber()] = (_this).f15;
          _nw203[(new BigNumber(7)).toNumber()] = (_this).f6;
          _nw203[(new BigNumber(8)).toNumber()] = new BigNumber((_1283_v0).length);
          _nw203[(new BigNumber(9)).toNumber()] = (_this).f6;
          _nw203[(new BigNumber(10)).toNumber()] = (_this).f15;
          _nw203[(new BigNumber(11)).toNumber()] = (_this).f6;
          _nw203[(new BigNumber(12)).toNumber()] = new BigNumber((_1358_v52).cardinality());
          _nw203[(new BigNumber(13)).toNumber()] = new BigNumber(99);
          _nw203[(new BigNumber(14)).toNumber()] = (_this).f15;
          _nw203[(new BigNumber(15)).toNumber()] = new BigNumber(-832);
          _nw203[(new BigNumber(16)).toNumber()] = (_this).f6;
          _nw203[(new BigNumber(17)).toNumber()] = (_this).f6;
          _nw203[(new BigNumber(18)).toNumber()] = (_1363_v57).fm21((_this).f15, globalState);
          _nw203[(new BigNumber(19)).toNumber()] = new BigNumber((p2).length);
          _nw203[(new BigNumber(20)).toNumber()] = (_this).f6;
          _nw203[(new BigNumber(21)).toNumber()] = (_this).f15;
          _nw203[(new BigNumber(22)).toNumber()] = (_this).f15;
          _nw203[(new BigNumber(23)).toNumber()] = (_this).f15;
          _nw203[(new BigNumber(24)).toNumber()] = (_this).f6;
          _nw203[(new BigNumber(25)).toNumber()] = (_this).f15;
          _1365_v59 = _nw203;
          let _1366_v60;
          _1366_v60 = _module.D2.create_DC6(new BigNumber((_dafny.Seq.update(_1283_v0, _module.__default.safeIndex((_this).f15, new BigNumber((_1283_v0).length)), _1285_v2)).length), _1365_v59);
          let _1367_v61;
          let _nw204 = new _module.C2();
          _nw204.__ctor((_this).f15, _1359_v53, _1360_v54, function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of (_module.__default.fm35((_1366_v60).dtor_cf13, (_this).f6, p1, (_this).f6, globalState)).Elements) {
              let _1368_v58 = _compr_39;
              if ((_module.__default.fm35((_1366_v60).dtor_cf13, (_this).f6, p1, (_this).f6, globalState)).contains(_1368_v58)) {
                _coll39.push([(_1368_v58).plus(new BigNumber((_1332_v34).length)),_1283_v0]);
              }
            }
            return _coll39;
          }());
          _1367_v61 = _nw204;
          let _1369_v62;
          _1369_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(651),_1365_v59);
          _1369_v62 = (_1369_v62).update((_this).f6, _1365_v59);
          let _1370_v63;
          _1370_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1285_v2,p0);
          let _1371_v64;
          _1371_v64 = _dafny.Set.fromElements(p0, p1);
          _1370_v63 = (_1370_v63).update(_1285_v2, !_dafny.areEqual(_dafny.Seq.update(_dafny.Seq.of((_1367_v61).fm21((_this).f6, globalState), new BigNumber((_module.__default.fm35((_this).f6, (_this).f6, p1, (_this).f6, globalState)).cardinality()), (_this).f6, new BigNumber(991), (_this).f15), _module.__default.safeIndex(new BigNumber((_module.__default.fm40((_this).f6, _1371_v64, globalState)).length), new BigNumber((_dafny.Seq.of((_1367_v61).fm21((_this).f6, globalState), new BigNumber((_module.__default.fm35((_this).f6, (_this).f6, p1, (_this).f6, globalState)).cardinality()), (_this).f6, new BigNumber(991), (_this).f15)).length)), new BigNumber(409)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(761)), function (_1372_i8) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("tvotwjin")).length);
          })));
        }
        let _1373_v65;
        _1373_v65 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-371)), ((_1374_v2) => function (_1375_i9) {
          return _1374_v2;
        })(_1285_v2)), _1283_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(714)), ((_1376_v2) => function (_1377_i10) {
          return _1376_v2;
        })(_1285_v2)));
        _1373_v65 = (((p1) || (p1)) ? (_1373_v65) : (_1373_v65));
        let _1378_v66;
        let _nw205 = Array((new BigNumber(6)).toNumber()).fill(false);
        _1378_v66 = _nw205;
        let _index198 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_1378_v66).length));
        (_1378_v66)[_index198] = (_1290_v7).IsSubsetOf(_module.__default.fm48(globalState));
        let _1379_v67;
        let _nw206 = Array((new BigNumber(11)).toNumber()).fill(_module.D1.Default());
        _1379_v67 = _nw206;
        let _1380_v68;
        _1380_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1379_v67,p0);
        let _index199 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_1378_v66).length));
        let _rhs224 = p1;
        let _rhs225 = ((_this).f15).isLessThan((_this).f15);
        let _rhs226 = _1380_v68;
        let _rhs227 = (_this).f15;
        let _rhs228 = (_1333_v35)[_module.__default.safeIndex((_this).f6, new BigNumber((_1333_v35).length))];
        let _lhs181 = _1378_v66;
        let _lhs182 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_1378_v66).length));
        let _lhs183 = globalState;
        _lhs181[_lhs182] = _rhs224;
        r2 = _rhs225;
        _1380_v68 = _rhs226;
        r3 = _rhs227;
        _lhs183.f0 = _rhs228;
        let _1381_v69;
        _1381_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f15);
        let _1382_v70;
        _1382_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1378_v66,(_module.__default.fm2((_this).f15, globalState)).isLessThan((((_1381_v69).contains(true)) ? ((_1381_v69).get(true)) : ((_this).f6))));
        _1382_v70 = (_dafny.Map.Empty.slice().updateUnsafe(_1378_v66,(_1378_v66)[_module.__default.safeIndex(new BigNumber(861), new BigNumber((_1378_v66).length))])).Merge(_1382_v70);
        let _1383_v71;
        _1383_v71 = _dafny.MultiSet.fromElements((_this).f15, (_this).f15);
        let _1384_v72;
        _1384_v72 = _dafny.MultiSet.fromElements(_1383_v71, _dafny.MultiSet.FromArray(_1284_v1));
        let _1385_v73;
        _1385_v73 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("sigwl")).length)),_1283_v0);
        let _1386_v74;
        _1386_v74 = _dafny.Seq.of(_1385_v73);
        let _1387_v75;
        let _nw207 = new _module.C2();
        _nw207.__ctor((_this).f15, _1384_v72, _1378_v66, (_1386_v74)[_module.__default.safeIndex((_this).f15, new BigNumber((_1386_v74).length))]);
        _1387_v75 = _nw207;
      } else {
        let _1388_v76;
        _1388_v76 = _dafny.Set.fromElements(_1285_v2);
        let _1389_v78;
        _1389_v78 = _dafny.Seq.of(_1388_v76, function () {
          let _coll40 = new _dafny.Set();
          for (const _compr_40 of (_dafny.Map.Empty.slice().updateUnsafe(_1285_v2,p0)).Keys.Elements) {
            let _1390_v77 = _compr_40;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_1285_v2,p0)).contains(_1390_v77)) {
              _coll40.add(_1390_v77);
            }
          }
          return _coll40;
        }());
        (globalState).f5 = _module.__default.safeDivisionInt(new BigNumber((_1389_v78).length), ((p0) ? (new BigNumber(988)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(699)), function (_1391_i11) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length))));
        let _1392_v79;
        _1392_v79 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(554)), function (_1393_i12) {
          return new BigNumber(-145);
        }), _1284_v1),p0);
        _1392_v79 = (_1392_v79).update(p0, p0);
        (globalState).f0 = p1;
        _1290_v7 = (_1290_v7).Union(_dafny.MultiSet.fromElements(p0));
        let _1394_v80;
        _1394_v80 = _dafny.Set.fromElements((_this).f15);
        let _1395_v81;
        let _init35 = function (_1396_i13) {
          return (_1396_i13).minus((_dafny.ZERO).minus((_this).f15));
        };
        let _nw208 = Array((new BigNumber(5)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw208.length); _i0_35++) {
          _nw208[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1395_v81 = _nw208;
        let _1397_v82;
        _1397_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_1395_v81);
        let _1398_v83;
        _1398_v83 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f6));
        let _1399_v84;
        _1399_v84 = _dafny.MultiSet.fromElements((_1284_v1)[_module.__default.safeIndex(_module.__default.fm2(new BigNumber((_1398_v83).cardinality()), globalState), new BigNumber((_1284_v1).length))]);
        let _1400_v85;
        _1400_v85 = _dafny.Map.Empty.slice().updateUnsafe(p1,new _dafny.CodePoint('a'.codePointAt(0)));
        let _1401_v86;
        let _nw209 = Array((new BigNumber(16)).toNumber());
        _nw209[(_dafny.ZERO).toNumber()] = (_this).f15;
        _nw209[(_dafny.ONE).toNumber()] = new BigNumber((_1394_v80).length);
        _nw209[(new BigNumber(2)).toNumber()] = (_this).f15;
        _nw209[(new BigNumber(3)).toNumber()] = (_this).f6;
        _nw209[(new BigNumber(4)).toNumber()] = new BigNumber((_1397_v82).length);
        _nw209[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("loqmtebnn")).length);
        _nw209[(new BigNumber(6)).toNumber()] = (_this).f15;
        _nw209[(new BigNumber(7)).toNumber()] = (_this).f6;
        _nw209[(new BigNumber(8)).toNumber()] = (_this).f15;
        _nw209[(new BigNumber(9)).toNumber()] = (_this).f6;
        _nw209[(new BigNumber(10)).toNumber()] = new BigNumber((_1283_v0).length);
        _nw209[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p1)).length);
        _nw209[(new BigNumber(12)).toNumber()] = (_this).f6;
        _nw209[(new BigNumber(13)).toNumber()] = (_this).f15;
        _nw209[(new BigNumber(14)).toNumber()] = (((_1399_v84).contains(new BigNumber((_1400_v85).length))) ? ((_1399_v84).get(new BigNumber((_1400_v85).length))) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(p0)).length))));
        _nw209[(new BigNumber(15)).toNumber()] = (_this).f15;
        _1401_v86 = _nw209;
        let _1402_v87;
        _1402_v87 = _module.D2.create_DC6((new BigNumber((_dafny.Seq.of(_1333_v35)).length)).plus((_this).f15), _1401_v86);
        let _source18 = _1402_v87;
        if (_source18.is_DC6) {
          let _1403___mcc_h3 = (_source18).cf13;
          let _1404___mcc_h4 = (_source18).cf14;
          let _1405_cf14 = _1404___mcc_h4;
          let _1406_cf13 = _1403___mcc_h3;
          (globalState).f5 = (_module.__default.safeModuloInt((_this).f6, new BigNumber((_1283_v0).length))).multipliedBy((_this).f6);
          r3 = ((_this).f6).plus((_dafny.ZERO).minus(((_this).f15).plus(new BigNumber(337))));
          (globalState).f2 = _1283_v0;
          (globalState).f0 = ((p0) ? (p0) : (((_this).f15).isLessThanOrEqualTo((_this).f15)));
        } else if (_source18.is_DC7) {
          let _1407___mcc_h5 = (_source18).cf15;
          let _1408___mcc_h6 = (_source18).cf16;
          let _1409___mcc_h7 = (_source18).cf17;
          let _1410___mcc_h8 = (_source18).cf18;
          let _1411_cf18 = _1410___mcc_h8;
          let _1412_cf17 = _1409___mcc_h7;
          let _1413_cf16 = _1408___mcc_h6;
          let _1414_cf15 = _1407___mcc_h5;
          (globalState).f0 = ((((p0) ? (p0) : (p0))) ? (p1) : ((_1414_cf15)[_module.__default.safeIndex((_this).f15, new BigNumber((_1414_cf15).length))]));
          let _1415_v88;
          _1415_v88 = _module.D0.create_DC1(!(p0), new BigNumber(88), new _dafny.CodePoint('f'.codePointAt(0)));
          let _1416_v89;
          _1416_v89 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1290_v7).cardinality()));
          let _pat_let_tv24 = _1285_v2;
          let _1417_v90;
          let _nw210 = Array((new BigNumber(26)).toNumber());
          _nw210[(_dafny.ZERO).toNumber()] = _1415_v88;
          _nw210[(_dafny.ONE).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(2)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(3)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(4)).toNumber()] = ((p1) ? (_1415_v88) : (_module.D0.create_DC1(p1, _1411_cf18, _1285_v2)));
          _nw210[(new BigNumber(5)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(6)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(7)).toNumber()] = _module.__default.fm20(p1, !(p0), true, globalState);
          _nw210[(new BigNumber(8)).toNumber()] = function (_pat_let31_0) {
            return function (_1418_dt__update__tmp_h3) {
              return function (_pat_let32_0) {
                return function (_1419_dt__update_hcf4_h0) {
                  return _module.D0.create_DC1((_1418_dt__update__tmp_h3).dtor_cf2, (_1418_dt__update__tmp_h3).dtor_cf3, _1419_dt__update_hcf4_h0);
                }(_pat_let32_0);
              }(_pat_let_tv24);
            }(_pat_let31_0);
          }(_1415_v88);
          _nw210[(new BigNumber(9)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(10)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(11)).toNumber()] = _module.D0.create_DC1((_1333_v35)[_module.__default.safeIndex((((_1416_v89).contains(p1)) ? ((_1416_v89).get(p1)) : ((_this).f6)), new BigNumber((_1333_v35).length))], (_this).f15, _1285_v2);
          _nw210[(new BigNumber(12)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(13)).toNumber()] = ((p1) ? (_1415_v88) : (_1415_v88));
          _nw210[(new BigNumber(14)).toNumber()] = _module.__default.fm20(p0, p1, p1, globalState);
          _nw210[(new BigNumber(15)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(16)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(17)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(18)).toNumber()] = _module.D0.create_DC1(p1, _1412_cf17, _1285_v2);
          _nw210[(new BigNumber(19)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(20)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(21)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(22)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(23)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(24)).toNumber()] = _1415_v88;
          _nw210[(new BigNumber(25)).toNumber()] = _1415_v88;
          _1417_v90 = _nw210;
          _1417_v90 = _1417_v90;
          let _1420_v91;
          let _nw211 = Array((new BigNumber(8)).toNumber());
          _nw211[(_dafny.ZERO).toNumber()] = true;
          _nw211[(_dafny.ONE).toNumber()] = !(p1);
          _nw211[(new BigNumber(2)).toNumber()] = p0;
          _nw211[(new BigNumber(3)).toNumber()] = p0;
          _nw211[(new BigNumber(4)).toNumber()] = !(false);
          _nw211[(new BigNumber(5)).toNumber()] = p1;
          _nw211[(new BigNumber(6)).toNumber()] = p1;
          _nw211[(new BigNumber(7)).toNumber()] = p1;
          _1420_v91 = _nw211;
          let _1421_v92;
          let _nw212 = new _module.C0();
          _nw212.__ctor(_1420_v91, true);
          _1421_v92 = _nw212;
          (globalState).f2 = _dafny.Seq.Concat(_1283_v0, _dafny.Seq.Concat(_1283_v0, _1283_v0));
        } else if (_source18.is_DC8) {
          let _1422___mcc_h9 = (_source18).cf19;
          let _1423___mcc_h10 = (_source18).cf20;
          let _1424_cf20 = _1423___mcc_h10;
          let _1425_cf19 = _1422___mcc_h9;
          _1285_v2 = new _dafny.CodePoint('l'.codePointAt(0));
          (globalState).f5 = (_this).f15;
          let _1426_v93;
          _1426_v93 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f6, new BigNumber((_1333_v35).length)));
          let _1427_v94;
          _1427_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1283_v0,_1286_v3);
          let _pat_let_tv25 = _1427_v94;
          let _pat_let_tv26 = _1283_v0;
          let _pat_let_tv27 = _1286_v3;
          let _pat_let_tv28 = _1427_v94;
          let _pat_let_tv29 = _1283_v0;
          let _index200 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1289_v6).length));
          let _rhs229 = (_1426_v93)[_module.__default.safeIndex((_this).f6, new BigNumber((_1426_v93).length))];
          let _rhs230 = function (_pat_let33_0) {
            return function (_1428_dt__update__tmp_h4) {
              return function (_pat_let34_0) {
                return function (_1429_dt__update_hcf51_h0) {
                  return _module.D9.create_DC27(_1429_dt__update_hcf51_h0);
                }(_pat_let34_0);
              }((((_pat_let_tv28).contains(_pat_let_tv29)) ? ((_pat_let_tv25).get(_pat_let_tv26)) : (_module.D9.create_DC27(_pat_let_tv27))));
            }(_pat_let33_0);
          }(_1288_v5);
          let _rhs231 = p0;
          let _lhs184 = _1289_v6;
          let _lhs185 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_1289_v6).length));
          let _lhs186 = globalState;
          _1394_v80 = _rhs229;
          _lhs184[_lhs185] = _rhs230;
          _lhs186.f0 = _rhs231;
          let _1430_v95;
          _1430_v95 = _module.D9.create_DC25((_this).f15, (_this).f15, (_this).f6, _1285_v2);
          let _1431_v96;
          _1431_v96 = _dafny.Map.Empty.slice().updateUnsafe((_1402_v87).dtor_cf14,_1430_v95);
          _1431_v96 = (_1431_v96).update(_1401_v86, _1430_v95);
        } else {
          let _1432___mcc_h11 = (_source18).cf12;
          let _1433_cf12 = _1432___mcc_h11;
          let _1434_v97;
          _1434_v97 = _dafny.Set.fromElements(_1433_cf12);
          _1434_v97 = _1434_v97;
          let _1435_v98;
          let _nw213 = new _module.C1();
          _nw213.__ctor(((_this).f6).minus((_this).f15));
          _1435_v98 = _nw213;
          let _rhs232 = p0;
          let _rhs233 = true;
          let _lhs187 = globalState;
          r2 = _rhs232;
          _lhs187.f0 = _rhs233;
          _1395_v81 = _1401_v86;
        }
      }
      r0 = (_module.__default.safeModuloInt(new BigNumber((_1283_v0).length), new BigNumber(352))).multipliedBy((_this).f15);
      r1 = _1284_v1;
      r2 = !(p0) || ((new BigNumber((_dafny.Seq.of((_this).f6, new BigNumber(895), (_this).f6, (_this).f6, (_this).f6)).length)).isLessThan((_this).f6));
      let _1436_v99;
      let _nw214 = Array((new BigNumber(19)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1436_v99 = _nw214;
      let _1437_v100;
      _1437_v100 = _module.D2.create_DC7(_1332_v34, _1436_v99, new BigNumber((_1283_v0).length), (_dafny.ZERO).minus((_this).f6));
      r3 = (_1437_v100).dtor_cf18;
      return [r0, r1, r2, r3];
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f13 = _dafny.ZERO;
      this._f14 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f13, f14) {
      let _this = this;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      return;
    }
    fm15(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(903)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-992)), function (_1438_i0) {
        return (_this).f13;
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), function (_1439_i1) {
        return new BigNumber(741);
      })));
    };
    fm16(p0, p1, globalState) {
      let _this = this;
      return ((_this).f13).multipliedBy(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f14))).length));
    };
    m8(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _module.D2.Default();
      let r2 = false;
      let _1440_v0;
      _1440_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
      let _1441_v1;
      _1441_v1 = _dafny.Map.Empty.slice().updateUnsafe((((_1440_v0).contains((_this).f13)) ? ((_1440_v0).get((_this).f13)) : (new BigNumber(956))),(_this).f14);
      _1441_v1 = (_1441_v1).update((_this).f13, (_this).f14);
      (globalState).f0 = _module.__default.fm3(globalState);
      let _1442_i0;
      _1442_i0 = _dafny.ZERO;
      L8: {
        while ((_this).f14) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1442_i0)) {
              break L8;
            }
            _1442_i0 = (_1442_i0).plus(_dafny.ONE);
            let _1443_v2;
            let _nw215 = Array((new BigNumber(25)).toNumber()).fill(false);
            _1443_v2 = _nw215;
            let _1444_v3;
            _1444_v3 = _dafny.MultiSet.fromElements((_this).f13);
            let _1445_v4;
            _1445_v4 = _dafny.Seq.of(new BigNumber((_1444_v3).cardinality()), (_this).f13);
            let _1446_v5;
            _1446_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1443_v2,_1445_v4);
            _1446_v5 = (_1446_v5).update(_1443_v2, _1445_v4);
            let _1447_v6;
            _1447_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_dafny.ZERO).minus(new BigNumber(-212)));
            _1447_v6 = (_1447_v6).update((_this).f14, (_this).f13);
            r2 = (_this).f14;
            if ((_this).f14) {
              (globalState).f0 = !(false);
              (globalState).f5 = (_this).f13;
              let _1448_v7;
              _1448_v7 = _dafny.Seq.UnicodeFromString("cfytvh");
              let _1449_v8;
              _1449_v8 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("cevep"), (((_this).f14) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(447)), function (_1450_i1) {
                return new _dafny.CodePoint('h'.codePointAt(0));
              })) : (_dafny.Seq.UnicodeFromString("o"))), _dafny.Seq.Concat(_1448_v7, _1448_v7));
              (globalState).f2 = (_1449_v8)[_module.__default.safeIndex(((_dafny.ZERO).minus((_this).f13)).plus(new BigNumber(97)), new BigNumber((_1449_v8).length))];
              let _1451_v9;
              _1451_v9 = new _dafny.CodePoint('e'.codePointAt(0));
              _1451_v9 = _1451_v9;
              (globalState).f5 = (_this).f13;
            } else {
              let _1452_v10;
              _1452_v10 = _dafny.Set.fromElements(_1443_v2, _1443_v2, _1443_v2, _1443_v2, _1443_v2);
              _1452_v10 = _1452_v10;
              let _1453_v11;
              _1453_v11 = _dafny.Seq.of((_this).f14, (_this).f14);
              let _1454_v12;
              _1454_v12 = _dafny.Seq.UnicodeFromString("snuxhfgl");
              let _1455_v13;
              _1455_v13 = _dafny.Set.fromElements(_1454_v12);
              let _1456_v14;
              _1456_v14 = _module.D0.create_DC0(_1453_v11, new BigNumber((_1455_v13).length));
              let _1457_v15;
              _1457_v15 = _module.D1.create_DC4(_1456_v14, new BigNumber(923), _module.__default.fm17((_this).f13, (_this).f13, globalState), (_this).f14, (_this).f13);
              (globalState).f2 = (_1457_v15).dtor_cf9;
              let _1458_v16;
              let _nw216 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
              _1458_v16 = _nw216;
              let _1459_v17;
              _1459_v17 = _dafny.MultiSet.fromElements(_1458_v16, _1458_v16, _1458_v16, _1458_v16);
              let _1460_v18;
              _1460_v18 = _dafny.Seq.of(_1445_v4);
              let _1461_v19;
              _1461_v19 = new _dafny.CodePoint('f'.codePointAt(0));
              let _1462_v20;
              _1462_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1461_v19,_1458_v16);
              let _1463_v21;
              _1463_v21 = _dafny.Seq.of((_dafny.MultiSet.fromElements(_1458_v16)).update((((_1462_v20).contains(new _dafny.CodePoint('e'.codePointAt(0)))) ? ((_1462_v20).get(new _dafny.CodePoint('e'.codePointAt(0)))) : (_1458_v16)), _module.__default.abs(new BigNumber((_dafny.Seq.of(function (_pat_let35_0) {
                return function (_1464_dt__update__tmp_h0) {
                  return function (_pat_let36_0) {
                    return function (_1465_dt__update_hcf1_h0) {
                      return _module.D0.create_DC0((_1464_dt__update__tmp_h0).dtor_cf0, _1465_dt__update_hcf1_h0);
                    }(_pat_let36_0);
                  }((_this).f13);
                }(_pat_let35_0);
              }(_1456_v14), _1456_v14, _1456_v14, _1456_v14)).length))), _1459_v17);
              _1459_v17 = ((((_1459_v17).update(_1458_v16, _module.__default.abs(new BigNumber(774)))).update(_1458_v16, _module.__default.abs(new BigNumber(((_1460_v18)[_module.__default.safeIndex((_this).f13, new BigNumber((_1460_v18).length))]).length)))).Union(_1459_v17)).Union((_1459_v17).Difference((_1463_v21)[_module.__default.safeIndex((_this).f13, new BigNumber((_1463_v21).length))]));
              let _index201 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_1458_v16).length));
              (_1458_v16)[_index201] = new BigNumber((_1444_v3).cardinality());
              let _1466_v22;
              _1466_v22 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("vp"), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.UnicodeFromString("vp")).length)), new _dafny.CodePoint('q'.codePointAt(0)))).length));
              let _index202 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_1458_v16).length));
              let _rhs234 = _module.__default.safeDivisionInt((_this).f13, new BigNumber((_1466_v22).length));
              let _rhs235 = (_dafny.ZERO).minus(((_this).f13).multipliedBy((((_1444_v3).contains((_this).f13)) ? ((_1444_v3).get((_this).f13)) : ((_this).f13))));
              let _lhs188 = _1458_v16;
              let _lhs189 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_1458_v16).length));
              let _lhs190 = globalState;
              _lhs188[_lhs189] = _rhs234;
              _lhs190.f5 = _rhs235;
              let _index203 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_1458_v16).length));
              (_1458_v16)[_index203] = (_this).f13;
            }
          }
        }
      }
      let _hi10 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f14)).length);
      for (let _1467_i2 = (_this).f13; _1467_i2.isLessThan(_hi10); _1467_i2 = _1467_i2.plus(_dafny.ONE)) {
        let _1468_v23;
        _1468_v23 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f14),(_this).f13);
        let _1469_v24;
        _1469_v24 = _module.D2.create_DC5(_dafny.Map.Empty.slice().updateUnsafe((((_1468_v23).contains((_this).f14)) ? ((_1468_v23).get((_this).f14)) : ((_dafny.ZERO).minus(_1467_i2))),(_dafny.ZERO).minus(_1467_i2)));
        let _source19 = _1469_v24;
        if (_source19.is_DC6) {
          let _1470___mcc_h0 = (_source19).cf13;
          let _1471___mcc_h1 = (_source19).cf14;
          let _1472_cf14 = _1471___mcc_h1;
          let _1473_cf13 = _1470___mcc_h0;
          let _1474_v25;
          _1474_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f14);
          _1474_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_module.__default.fm3(globalState));
          let _1475_v26;
          _1475_v26 = _dafny.Seq.of(_1468_v23, _1468_v23, _dafny.Map.Empty.slice().updateUnsafe(!((_this).f14),_1473_cf13));
          let _1476_v27;
          _1476_v27 = _dafny.Seq.UnicodeFromString("qk");
          let _1477_v28;
          _1477_v28 = _dafny.Seq.of((_this).f14);
          r2 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_1475_v26, _module.__default.safeIndex(new BigNumber((_1476_v27).length), new BigNumber((_1475_v26).length)), ((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((_1477_v28).length))).update((_this).f14, new BigNumber(83))).update((_this).f14, _1473_cf13)), _1475_v26);
          _1474_v25 = (_1474_v25).update(true, (((_1474_v25).contains((_this).f14)) ? ((_1474_v25).get((_this).f14)) : (!((_this).f14))));
          (globalState).f0 = (_this).f14;
        } else if (_source19.is_DC7) {
          let _1478___mcc_h2 = (_source19).cf15;
          let _1479___mcc_h3 = (_source19).cf16;
          let _1480___mcc_h4 = (_source19).cf17;
          let _1481___mcc_h5 = (_source19).cf18;
          let _1482_cf18 = _1481___mcc_h5;
          let _1483_cf17 = _1480___mcc_h4;
          let _1484_cf16 = _1479___mcc_h3;
          let _1485_cf15 = _1478___mcc_h2;
          let _1486_v29;
          let _nw217 = new _module.C1();
          _nw217.__ctor((_this).f13);
          _1486_v29 = _nw217;
          let _1487_v30;
          _1487_v30 = _dafny.Seq.of(_1485_cf15, _1485_cf15, _1485_cf15, _1485_cf15, _1485_cf15);
          _1487_v30 = _1487_v30;
          (globalState).f5 = (_this).f13;
          (globalState).f5 = (_this).f13;
        } else if (_source19.is_DC8) {
          let _1488___mcc_h6 = (_source19).cf19;
          let _1489___mcc_h7 = (_source19).cf20;
          let _1490_cf20 = _1489___mcc_h7;
          let _1491_cf19 = _1488___mcc_h6;
          (globalState).f5 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_this).f13, (_this).f13), (_this).f13);
          let _1492_v31;
          let _nw218 = Array((new BigNumber(13)).toNumber()).fill([]);
          _1492_v31 = _nw218;
          let _1493_v32;
          let _nw219 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1493_v32 = _nw219;
          let _1494_v33;
          _1494_v33 = _dafny.MultiSet.fromElements(_1493_v32, _1493_v32);
          let _1495_v34;
          _1495_v34 = new _dafny.CodePoint('g'.codePointAt(0));
          let _1496_v35;
          _1496_v35 = _dafny.MultiSet.fromElements(_1495_v34);
          let _1497_v36;
          _1497_v36 = _dafny.Seq.of((_this).f14, _module.__default.fm3(globalState));
          let _1498_v37;
          let _nw220 = Array((new BigNumber(9)).toNumber());
          _nw220[(_dafny.ZERO).toNumber()] = _1468_v23;
          _nw220[(_dafny.ONE).toNumber()] = _1468_v23;
          _nw220[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1491_cf19,_1467_i2);
          _nw220[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1490_cf20,new BigNumber((_1494_v33).cardinality()));
          _nw220[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((_1496_v35).cardinality()));
          _nw220[(new BigNumber(5)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1491_cf19,(_this).f13)).update(_1491_cf19, (_dafny.ZERO).minus(_1467_i2));
          _nw220[(new BigNumber(6)).toNumber()] = _1468_v23;
          _nw220[(new BigNumber(7)).toNumber()] = _module.__default.fm31(_1491_cf19, (_this).f13, _1467_i2, _1497_v36, globalState);
          _nw220[(new BigNumber(8)).toNumber()] = _1468_v23;
          _1498_v37 = _nw220;
          let _index204 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1492_v31).length));
          (_1492_v31)[_index204] = _1498_v37;
          let _index205 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1492_v31).length));
          (_1492_v31)[_index205] = _1498_v37;
          (globalState).f5 = (_dafny.ZERO).minus(((_this).f13).minus(((_this).f13).minus((_this).f13)));
          let _1499_v38;
          let _init36 = function (_1500_i3) {
            return _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("tsfvbtav"), new _dafny.CodePoint('c'.codePointAt(0)));
          };
          let _nw221 = Array((new BigNumber(9)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw221.length); _i0_36++) {
            _nw221[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1499_v38 = _nw221;
          let _1501_v39;
          _1501_v39 = _dafny.Seq.of(_1499_v38);
          _1499_v38 = (_1501_v39)[_module.__default.safeIndex((new BigNumber(-871)).minus(new BigNumber((_module.__default.fm4(_1467_i2, (_this).f13, (_this).f13, globalState)).length)), new BigNumber((_1501_v39).length))];
        } else {
          let _1502___mcc_h8 = (_source19).cf12;
          let _1503_cf12 = _1502___mcc_h8;
          let _1504_v40;
          let _init37 = ((_1505_i2) => function (_1506_i4) {
            return _module.__default.safeDivisionInt(_1506_i4, (_dafny.ZERO).minus(_1505_i2));
          })(_1467_i2);
          let _nw222 = Array((new BigNumber(22)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw222.length); _i0_37++) {
            _nw222[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1504_v40 = _nw222;
          let _1507_v41;
          _1507_v41 = _dafny.MultiSet.fromElements(_1467_i2);
          let _1508_v42;
          _1508_v42 = _dafny.MultiSet.fromElements(_1507_v41);
          let _index206 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1504_v40).length));
          (_1504_v40)[_index206] = _module.__default.safeModuloInt(new BigNumber((_1508_v42).cardinality()), _1467_i2);
          let _index207 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1504_v40).length));
          (_1504_v40)[_index207] = (_this).f13;
          let _1509_v43;
          _1509_v43 = _dafny.Seq.UnicodeFromString("ecuurfubj");
          let _index208 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1504_v40).length));
          (_1504_v40)[_index208] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gyxgc"), _1509_v43)).length));
          let _1510_v44;
          _1510_v44 = new _dafny.CodePoint('t'.codePointAt(0));
          let _1511_v45;
          _1511_v45 = _dafny.Set.fromElements(_dafny.Seq.update(_1509_v43, _module.__default.safeIndex((_1504_v40)[_module.__default.safeIndex(new BigNumber(298), new BigNumber((_1504_v40).length))], new BigNumber((_1509_v43).length)), _1510_v44), _1509_v43, _1509_v43);
          let _1512_v46;
          _1512_v46 = _dafny.Seq.of((_1511_v45).Intersect(_1511_v45));
          _1511_v45 = (_1512_v46)[_module.__default.safeIndex(_1467_i2, new BigNumber((_1512_v46).length))];
          let _1513_v47;
          let _init38 = function (_1514_i5) {
            return (_this).f14;
          };
          let _nw223 = Array((new BigNumber(12)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw223.length); _i0_38++) {
            _nw223[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1513_v47 = _nw223;
          let _1515_v48;
          _1515_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_1509_v43);
          let _1516_v49;
          let _nw224 = new _module.C3();
          _nw224.__ctor(_1440_v0, _1513_v47, _1515_v48, (_this).f13);
          _1516_v49 = _nw224;
          let _1517_v50;
          _1517_v50 = _dafny.Seq.of(_1516_v49);
          let _1518_v51;
          _1518_v51 = _dafny.Set.fromElements((_1517_v50)[_module.__default.safeIndex((_this).f13, new BigNumber((_1517_v50).length))]);
          let _rhs236 = (_this).f14;
          let _rhs237 = (_1518_v51).Union((_dafny.Set.fromElements(_1516_v49, _1516_v49)).Difference(_1518_v51));
          let _lhs191 = globalState;
          _lhs191.f0 = _rhs236;
          _1518_v51 = _rhs237;
        }
        let _1519_v52;
        _1519_v52 = _dafny.Seq.of(_1467_i2, _1467_i2);
        let _1520_v53;
        _1520_v53 = _dafny.Set.fromElements(_dafny.Seq.IsProperPrefixOf(_1519_v52, _1519_v52), ((_dafny.ZERO).minus((_this).f13)).isLessThanOrEqualTo((_this).f13), (_this).f14, (_this).f14, (_this).f14);
        _1520_v53 = (_1520_v53).Intersect(_1520_v53);
        let _1521_v54;
        _1521_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,!((_this).f14));
        let _1522_v55;
        _1522_v55 = _module.D4.create_DC14((_this).f14, _1521_v54, (_this).f13, _1521_v54);
        let _1523_v56;
        _1523_v56 = _dafny.Seq.of((_1522_v55).dtor_cf28, (_this).f14);
        let _1524_v57;
        _1524_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1467_i2,_1523_v56);
        let _1525_v58;
        _1525_v58 = _dafny.Seq.UnicodeFromString("wdgqi");
        let _1526_v59;
        _1526_v59 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_this).f14, (_this).f14),(_dafny.Map.Empty.slice().updateUnsafe(_1467_i2,_1467_i2)).update(new BigNumber((_1525_v58).length), _1467_i2));
        _1524_v57 = (_1524_v57).update(new BigNumber((_1526_v59).length), _1523_v56);
        (globalState).f5 = _1467_i2;
      }
      let _1527_v60;
      let _nw225 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _1527_v60 = _nw225;
      let _1528_v61;
      _1528_v61 = _module.D2.create_DC6((_this).f13, _1527_v60);
      _1528_v61 = _1528_v61;
      let _1529_v62;
      _1529_v62 = _module.D2.create_DC8(false, (_this).f14);
      let _1530_v63;
      _1530_v63 = new _dafny.CodePoint('h'.codePointAt(0));
      let _1531_v64;
      _1531_v64 = _module.D0.create_DC1((_this).f14, (_this).f13, _1530_v63);
      let _hi11 = _module.__default.safeModuloInt((_this).f13, (_this).fm16(_module.__default.fm49((_this).f13, (_this).fm16(_1529_v62, _1531_v64, globalState), (_this).f14, (_this).f13, globalState), _1531_v64, globalState));
      for (let _1532_i6 = (_this).f13; _1532_i6.isLessThan(_hi11); _1532_i6 = _1532_i6.plus(_dafny.ONE)) {
        (globalState).f0 = (_this).f14;
        _1440_v0 = (_1440_v0).update((_this).f13, (new BigNumber(-670)).plus((_this).f13));
        let _1533_v65;
        _1533_v65 = _dafny.MultiSet.fromElements((_this).f14, false);
        r0 = (new BigNumber((_1533_v65).cardinality())).multipliedBy(_module.__default.fm2(new BigNumber((_dafny.Set.fromElements(_1532_i6)).length), globalState));
        let _1534_v66;
        _1534_v66 = _dafny.Seq.of((_this).f13);
        (globalState).f0 = (new BigNumber((_dafny.Seq.Concat(_1534_v66, _1534_v66)).length)).isLessThanOrEqualTo(new BigNumber(-581));
      }
      r0 = (_this).f13;
      let _1535_v67;
      _1535_v67 = _module.D2.create_DC5(_1440_v0);
      r1 = _1535_v67;
      r2 = (_this).f14;
      return [r0, r1, r2];
    }
    m9(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.MultiSet.Empty;
      let _1536_v0;
      _1536_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f13);
      let _1537_v1;
      _1537_v1 = _dafny.MultiSet.fromElements((_this).f13);
      let _1538_v2;
      _1538_v2 = _dafny.MultiSet.fromElements(_1537_v1);
      let _1539_v3;
      let _init39 = function (_1540_i0) {
        return (_this).f14;
      };
      let _nw226 = Array((new BigNumber(22)).toNumber());
      for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw226.length); _i0_39++) {
        _nw226[_i0_39] = _init39(new BigNumber(_i0_39));
      }
      _1539_v3 = _nw226;
      let _1541_v4;
      _1541_v4 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_this).f13, globalState),p0);
      let _1542_v5;
      let _nw227 = new _module.C2();
      _nw227.__ctor((_this).f13, _1538_v2, _1539_v3, _1541_v4);
      _1542_v5 = _nw227;
      let _1543_v6;
      _1543_v6 = _dafny.Seq.of((_this).f14, (_this).f14);
      let _1544_v7;
      _1544_v7 = _dafny.Seq.of(new BigNumber((_1543_v6).length));
      let _1545_v8;
      _1545_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1542_v5,_1544_v7);
      _1536_v0 = (_1536_v0).update(_dafny.Seq.contains((((_1545_v8).contains(_1542_v5)) ? ((_1545_v8).get(_1542_v5)) : (_1544_v7)), (_this).f13), (_this).f13);
      let _1546_i1;
      _1546_i1 = _dafny.ZERO;
      L9: {
        while ((_this).f14) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1546_i1)) {
              break L9;
            }
            _1546_i1 = (_1546_i1).plus(_dafny.ONE);
            r1 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f13), (_this).f13);
            let _1547_v9;
            _1547_v9 = _dafny.Set.fromElements(_1539_v3, _1539_v3, _1539_v3, _1539_v3);
            _1547_v9 = _1547_v9;
            (globalState).f5 = (_this).f13;
            let _1548_v11;
            _1548_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f14);
            let _1549_v12;
            let _nw228 = new _module.C2();
            _nw228.__ctor((_this).f13, _1538_v2, _1539_v3, (function () {
              let _coll41 = new _dafny.Map();
              for (const _compr_41 of (_1548_v11).Keys.Elements) {
                let _1550_v10 = _compr_41;
                if ((_1548_v11).contains(_1550_v10)) {
                  _coll41.push([(_1550_v10).plus((_this).f13),p1]);
                }
              }
              return _coll41;
            }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-554),p1)));
            _1549_v12 = _nw228;
          }
        }
      }
      let _1551_v13;
      _1551_v13 = _module.D2.create_DC8((_this).f14, (_this).f14);
      let _1552_v14;
      _1552_v14 = new _dafny.CodePoint('d'.codePointAt(0));
      let _1553_v15;
      _1553_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_dafny.Set.fromElements((_this).f13));
      let _1554_v16;
      _1554_v16 = _dafny.Set.fromElements(new BigNumber((_1536_v0).length), (_this).f13, (_this).f13);
      let _rhs238 = (((false) ? (new BigNumber(449)) : ((_this).f13))).plus(((_this).f13).plus((_this).fm16(function (_pat_let37_0) {
        return function (_1555_dt__update__tmp_h0) {
          return function (_pat_let38_0) {
            return function (_1556_dt__update_hcf19_h0) {
              return _module.D2.create_DC8(_1556_dt__update_hcf19_h0, (_1555_dt__update__tmp_h0).dtor_cf20);
            }(_pat_let38_0);
          }((_this).f14);
        }(_pat_let37_0);
      }(_1551_v13), _module.D0.create_DC1(true, (_this).f13, _1552_v14), globalState)));
      let _rhs239 = _1539_v3;
      let _rhs240 = _1543_v6;
      let _rhs241 = ((_this).f13).isLessThan((_dafny.ZERO).minus((_this).f13));
      let _rhs242 = (_1554_v16).IsSubsetOf((((_1553_v15).contains((_this).f13)) ? ((_1553_v15).get((_this).f13)) : (_1554_v16)));
      let _lhs192 = globalState;
      let _lhs193 = globalState;
      r1 = _rhs238;
      _1539_v3 = _rhs239;
      _1543_v6 = _rhs240;
      _lhs192.f0 = _rhs241;
      _lhs193.f0 = _rhs242;
      if (((((_this).f13).isEqualTo((_this).f13)) ? ((_this).f14) : (_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("q"), _1552_v14)))) {
        let _1557_v17;
        _1557_v17 = _dafny.Set.fromElements((_this).f14, (_this).f14, (_this).f14);
        (globalState).f0 = (_1557_v17).IsDisjointFrom(_1557_v17);
        let _1558_v18;
        _1558_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
        let _1559_v19;
        _1559_v19 = _module.D2.create_DC5(_1558_v18);
        let _rhs243 = _module.D2.create_DC5(_1558_v18);
        let _rhs244 = (_1543_v6)[_module.__default.safeIndex(new BigNumber(-56), new BigNumber((_1543_v6).length))];
        let _rhs245 = ((_this).f13).minus((_1544_v7)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_1544_v7).length))]);
        let _lhs194 = globalState;
        let _lhs195 = globalState;
        _1559_v19 = _rhs243;
        _lhs194.f0 = _rhs244;
        _lhs195.f5 = _rhs245;
        r1 = ((_this).f13).plus((_this).f13);
        let _1560_v20;
        _1560_v20 = _dafny.Map.Empty.slice().updateUnsafe((_1543_v6)[_module.__default.safeIndex((_this).f13, new BigNumber((_1543_v6).length))],false);
        (globalState).f0 = (((_1560_v20).contains(((_this).f14) || (false))) ? ((_1560_v20).get(((_this).f14) || (false))) : (!((_1554_v16).IsDisjointFrom(_1554_v16))));
        let _1561_v21;
        let _init40 = function (_1562_i2) {
          return _dafny.Seq.UnicodeFromString("jfhhroih");
        };
        let _nw229 = Array((new BigNumber(28)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw229.length); _i0_40++) {
          _nw229[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1561_v21 = _nw229;
        let _index209 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1561_v21).length));
        (_1561_v21)[_index209] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pli"), p0);
        let _index210 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1561_v21).length));
        (_1561_v21)[_index210] = p0;
      } else {
        _1554_v16 = _1554_v16;
        let _1563_v22;
        let _nw230 = Array((new BigNumber(11)).toNumber());
        _nw230[(_dafny.ZERO).toNumber()] = _1539_v3;
        _nw230[(_dafny.ONE).toNumber()] = (((_this).f14) ? (_1539_v3) : (_1539_v3));
        _nw230[(new BigNumber(2)).toNumber()] = _1539_v3;
        _nw230[(new BigNumber(3)).toNumber()] = _1539_v3;
        _nw230[(new BigNumber(4)).toNumber()] = _1539_v3;
        _nw230[(new BigNumber(5)).toNumber()] = _1539_v3;
        _nw230[(new BigNumber(6)).toNumber()] = _1539_v3;
        _nw230[(new BigNumber(7)).toNumber()] = _1539_v3;
        _nw230[(new BigNumber(8)).toNumber()] = _1539_v3;
        _nw230[(new BigNumber(9)).toNumber()] = _1539_v3;
        _nw230[(new BigNumber(10)).toNumber()] = _1539_v3;
        _1563_v22 = _nw230;
        let _index211 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1563_v22).length));
        (_1563_v22)[_index211] = _1539_v3;
        let _index212 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1563_v22).length));
        (_1563_v22)[_index212] = _1539_v3;
        let _1564_v23;
        _1564_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1544_v7).length),(_this).f14);
        (globalState).f5 = (new BigNumber((_1564_v23).length)).plus(((_this).f13).multipliedBy((_this).f13));
        let _index213 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1539_v3).length));
        (_1539_v3)[_index213] = (_this).f14;
        let _index214 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1539_v3).length));
        (_1539_v3)[_index214] = false;
        _1544_v7 = _1544_v7;
      }
      let _1565_v24;
      _1565_v24 = _dafny.Set.fromElements(true);
      (globalState).f5 = new BigNumber((_1565_v24).length);
      let _1566_v25;
      _1566_v25 = _module.D1.create_DC4(_module.D0.create_DC0(_dafny.Seq.of((_this).f14, (_this).f14), (_this).f13), (_dafny.ZERO).minus((_this).f13), p1, (_this).f14, (_this).f13);
      let _1567_i3;
      _1567_i3 = _dafny.ZERO;
      L10: {
        while (!(!((_this).f13).isEqualTo(new BigNumber(623))) || (((_1566_v25).dtor_cf8).isLessThanOrEqualTo((_this).f13))) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1567_i3)) {
              break L10;
            }
            _1567_i3 = (_1567_i3).plus(_dafny.ONE);
            let _1568_v26;
            let _nw231 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
            _1568_v26 = _nw231;
            let _index215 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1568_v26).length));
            (_1568_v26)[_index215] = _dafny.Seq.Concat(_1544_v7, _dafny.Seq.update(_1544_v7, _module.__default.safeIndex((_this).f13, new BigNumber((_1544_v7).length)), new BigNumber((_dafny.Seq.of(!((_this).f14))).length)));
            let _1569_v27;
            _1569_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
            let _1570_v28;
            _1570_v28 = _module.D0.create_DC0(_module.__default.fm4((_this).f13, (_this).f13, (_this).f13, globalState), (_this).f13);
            let _1571_v29;
            _1571_v29 = _dafny.Set.fromElements(_1539_v3);
            let _pat_let_tv30 = _1570_v28;
            let _pat_let_tv31 = globalState;
            let _1572_v30;
            let _nw232 = Array((new BigNumber(5)).toNumber());
            _nw232[(_dafny.ZERO).toNumber()] = (new BigNumber((_1543_v6).length)).multipliedBy((_this).f13);
            _nw232[(_dafny.ONE).toNumber()] = (_this).f13;
            _nw232[(new BigNumber(2)).toNumber()] = new BigNumber(((_1569_v27).Merge(_1569_v27)).length);
            _nw232[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((function (_pat_let39_0) {
              return function (_1573_dt__update__tmp_h1) {
                return function (_pat_let40_0) {
                  return function (_1574_dt__update_hcf7_h0) {
                    return function (_pat_let41_0) {
                      return function (_1575_dt__update_hcf9_h0) {
                        return function (_pat_let42_0) {
                          return function (_1576_dt__update_hcf10_h0) {
                            return _module.D1.create_DC4(_1574_dt__update_hcf7_h0, (_1573_dt__update__tmp_h1).dtor_cf8, _1575_dt__update_hcf9_h0, _1576_dt__update_hcf10_h0, (_1573_dt__update__tmp_h1).dtor_cf11);
                          }(_pat_let42_0);
                        }((_this).f14);
                      }(_pat_let41_0);
                    }(_module.__default.fm18((_this).f14, _pat_let_tv31));
                  }(_pat_let40_0);
                }(_pat_let_tv30);
              }(_pat_let39_0);
            }(_1566_v25)).dtor_cf11);
            _nw232[(new BigNumber(4)).toNumber()] = new BigNumber((_1571_v29).length);
            _1572_v30 = _nw232;
            let _index216 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
            (_1572_v30)[_index216] = _module.__default.fm2((_this).f13, globalState);
            let _1577_v31;
            _1577_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,p0);
            let _index217 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1568_v26).length));
            let _index218 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
            let _rhs246 = _dafny.Seq.update(_module.__default.fm26(_module.__default.safeModuloInt(new BigNumber(((_1577_v31).update((_this).f14, p1)).length), new BigNumber((_1569_v27).length)), _1552_v14, (_this).f14, _1552_v14, globalState), _module.__default.safeIndex((_this).f13, new BigNumber((_module.__default.fm26(_module.__default.safeModuloInt(new BigNumber(((_1577_v31).update((_this).f14, p1)).length), new BigNumber((_1569_v27).length)), _1552_v14, (_this).f14, _1552_v14, globalState)).length)), (_this).f13);
            let _rhs247 = ((_module.__default.fm34(new BigNumber((p0).length), globalState)).Intersect(_1565_v24)).IsProperSubsetOf(_1565_v24);
            let _rhs248 = _1544_v7;
            let _rhs249 = (_this).f13;
            let _rhs250 = _module.__default.fm3(globalState);
            let _lhs196 = globalState;
            let _lhs197 = _1568_v26;
            let _lhs198 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1568_v26).length));
            let _lhs199 = _1572_v30;
            let _lhs200 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
            let _lhs201 = globalState;
            _1544_v7 = _rhs246;
            _lhs196.f0 = _rhs247;
            _lhs197[_lhs198] = _rhs248;
            _lhs199[_lhs200] = _rhs249;
            _lhs201.f0 = _rhs250;
            r1 = (_this).f13;
            if ((_this).f14) {
              let _1578_v32;
              let _nw233 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _1578_v32 = _nw233;
              _1578_v32 = _1578_v32;
              let _1579_v33;
              _1579_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1568_v26)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_1568_v26).length))],(_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))]);
              let _1580_v34;
              _1580_v34 = _dafny.Map.Empty.slice().updateUnsafe((_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))],_1569_v27);
              let _index219 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
              (_1572_v30)[_index219] = (((_1579_v33).contains(_dafny.Seq.Concat((_1568_v26)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_1568_v26).length))], _1544_v7))) ? ((_1579_v33).get(_dafny.Seq.Concat((_1568_v26)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_1568_v26).length))], _1544_v7))) : (new BigNumber((_1580_v34).length)));
              let _index220 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1568_v26).length));
              (_1568_v26)[_index220] = _dafny.Seq.Concat(_module.__default.fm40((_this).f13, _dafny.Set.fromElements((_this).f14, (_this).f14), globalState), (_1568_v26)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_1568_v26).length))]);
              let _index221 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
              let _rhs251 = (_dafny.ZERO).minus((_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))]);
              let _rhs252 = _module.__default.fm2(new BigNumber((_dafny.Seq.of(_1565_v24, _1565_v24)).length), globalState);
              let _lhs202 = _1572_v30;
              let _lhs203 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
              let _lhs204 = globalState;
              _lhs202[_lhs203] = _rhs251;
              _lhs204.f5 = _rhs252;
              let _1581_v35;
              _1581_v35 = _dafny.MultiSet.fromElements(_module.__default.fm3(globalState));
              (globalState).f0 = (_1542_v5).fm22(_1581_v35, ((_this).f13).isLessThanOrEqualTo((_this).f13), (_this).f13, globalState);
            } else {
              let _index222 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1539_v3).length));
              (_1539_v3)[_index222] = (_this).f14;
              let _1582_v36;
              _1582_v36 = _dafny.Map.Empty.slice().updateUnsafe((_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))],_1543_v6);
              let _1583_v38;
              _1583_v38 = _dafny.Map.Empty.slice().updateUnsafe((_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))],(_this).f14);
              let _1584_v40;
              _1584_v40 = _module.D11.create_DC30(function () {
  let _coll42 = new _dafny.Set();
  for (const _compr_42 of (_1583_v38).Keys.Elements) {
    let _1585_v39 = _compr_42;
    if ((_1583_v38).contains(_1585_v39)) {
      _coll42.add((_1585_v39).multipliedBy(new BigNumber(342)));
    }
  }
  return _coll42;
}());
              let _index223 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1539_v3).length));
              (_1539_v3)[_index223] = (((_1584_v40).dtor_cf57).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-74)), ((_1586_v14) => function (_1587_i4) {
                return _1586_v14;
              })(_1552_v14))).length), (_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))]))).IsProperSubsetOf((_1554_v16).Intersect(function () {
                let _coll43 = new _dafny.Set();
                for (const _compr_43 of (_1582_v36).Keys.Elements) {
                  let _1588_v37 = _compr_43;
                  if ((_1582_v36).contains(_1588_v37)) {
                    _coll43.add((_1588_v37).plus(new BigNumber(744)));
                  }
                }
                return _coll43;
              }()));
              let _index224 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1568_v26).length));
              (_1568_v26)[_index224] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(401)), ((_1589_v16) => function (_1590_i5) {
                return (_dafny.ZERO).minus(new BigNumber((_1589_v16).length));
              })(_1554_v16)), _1544_v7), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(401)), ((_1591_v16) => function (_1592_i5) {
                return (_dafny.ZERO).minus(new BigNumber((_1591_v16).length));
              })(_1554_v16)), _1544_v7)).length)), (new BigNumber(771)).multipliedBy((_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))]));
              (globalState).f5 = (_this).f13;
              let _1593_v41;
              _1593_v41 = _module.D11.create_DC33(((_this).f13).minus(new BigNumber((_1569_v27).length)), ((_this).f13).isLessThanOrEqualTo((_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))]), (_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))], false, _1565_v24);
              let _1594_v43;
              _1594_v43 = _module.D0.create_DC1((_this).f14, new BigNumber((function () {
  let _coll44 = new _dafny.Map();
  for (const _compr_44 of (p1).Elements) {
    let _1595_v42 = _compr_44;
    if (_dafny.Seq.contains(p1, _1595_v42)) {
      _coll44.push([_1595_v42,_1543_v6]);
    }
  }
  return _coll44;
}()).length), _1552_v14);
              let _index225 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
              let _index226 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
              let _index227 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1539_v3).length));
              let _rhs253 = _1593_v41;
              let _rhs254 = (_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))];
              let _rhs255 = (_this).fm16(_1551_v13, _1594_v43, globalState);
              let _rhs256 = ((_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))]).isLessThanOrEqualTo((_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))]);
              let _rhs257 = _1571_v29;
              let _lhs205 = _1572_v30;
              let _lhs206 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
              let _lhs207 = _1572_v30;
              let _lhs208 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length));
              let _lhs209 = _1539_v3;
              let _lhs210 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_1539_v3).length));
              _1593_v41 = _rhs253;
              _lhs205[_lhs206] = _rhs254;
              _lhs207[_lhs208] = _rhs255;
              _lhs209[_lhs210] = _rhs256;
              _1571_v29 = _rhs257;
              let _1596_v44;
              _1596_v44 = _dafny.MultiSet.fromElements((_1539_v3)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1539_v3).length))], (_this).f14, (_1539_v3)[_module.__default.safeIndex(new BigNumber(981), new BigNumber((_1539_v3).length))]);
              let _1597_v45;
              _1597_v45 = _module.D5.create_DC17(_1544_v7, !((_this).f14));
              let _rhs258 = (_1542_v5).fm22(_1596_v44, (_1577_v31).contains((_this).f14), ((_this).f13).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("uuiikup")).length)), globalState);
              let _rhs259 = ((_1554_v16).Intersect(_dafny.Set.fromElements(new BigNumber((_1543_v6).length), (_this).f13, (_this).f13, (_this).f13))).IsProperSubsetOf(_1554_v16);
              let _rhs260 = (((_1583_v38).contains(new BigNumber((p1).length))) ? ((_1583_v38).get(new BigNumber((p1).length))) : ((_this).f14));
              let _rhs261 = (((_this).f14) ? ((_dafny.ZERO).minus((_1572_v30)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1572_v30).length))])) : (_module.__default.safeDivisionInt((_this).f13, (_this).f13)));
              let _rhs262 = (_1597_v45).dtor_cf37;
              let _lhs211 = globalState;
              let _lhs212 = globalState;
              let _lhs213 = globalState;
              let _lhs214 = globalState;
              let _lhs215 = globalState;
              _lhs211.f0 = _rhs258;
              _lhs212.f0 = _rhs259;
              _lhs213.f0 = _rhs260;
              _lhs214.f5 = _rhs261;
              _lhs215.f0 = _rhs262;
            }
            (globalState).f2 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xpbyxlev"), _module.__default.fm18((_this).f14, globalState));
          }
        }
      }
      let _1598_v46;
      let _init41 = function (_1599_i6) {
        return (_1599_i6).multipliedBy((_this).f13);
      };
      let _nw234 = Array((new BigNumber(2)).toNumber());
      for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw234.length); _i0_41++) {
        _nw234[_i0_41] = _init41(new BigNumber(_i0_41));
      }
      _1598_v46 = _nw234;
      let _1600_v47;
      _1600_v47 = _dafny.Set.fromElements(_1598_v46, _1598_v46, _1598_v46, _1598_v46);
      r0 = _1600_v47;
      r1 = ((_this).f13).plus(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f13), (_this).f13));
      r2 = (_1537_v1).Intersect((_1537_v1).Difference(_1537_v1));
      return [r0, r1, r2];
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f9 = [];
      this._f6 = _dafny.ZERO;
      this._f10 = _dafny.Map.Empty;
      this.f11 = [];
      this.f12 = [];
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    set f9(value) {
      let _this = this;
      _this._f9 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f11, f12, f6, f9, f10) {
      let _this = this;
      (_this).f11 = f11;
      (_this).f12 = f12;
      (_this)._f6 = f6;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-609)), function (_1601_i0) {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(596)), function (_1602_i1) {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), (_module.D0.create_DC1(false, new BigNumber(711), new _dafny.CodePoint('g'.codePointAt(0)))).dtor_cf4, new _dafny.CodePoint('k'.codePointAt(0)));
      })), _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)))));
    };
    fm6(p0, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f6),_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f6)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f6, (_this).f6)).length), (_this).f6, (_this).f6, (_this).f6)));
    };
    fm13(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (((false) ? (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("j"))) : (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("ggjslaqas"))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("kdhe"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.Create(_module.__default.abs(new BigNumber(298)), function (_1603_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }))));
    };
    fm14(globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(((_module.D0.create_DC0(_dafny.Seq.of(false, false), (_this).f6)).dtor_cf1).minus(((_this).f6).plus((_this).f6)));
    };
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let _hi12 = (_this).fm14(globalState);
      for (let _1604_i0 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p0, p0), p1)).length); _1604_i0.isLessThan(_hi12); _1604_i0 = _1604_i0.plus(_dafny.ONE)) {
        (globalState).f5 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(((p0) ? (p2) : (_1604_i0)), _1604_i0));
        (globalState).f2 = _dafny.Seq.UnicodeFromString("bksq");
        let _1605_v0;
        _1605_v0 = _dafny.Seq.UnicodeFromString("rwrd");
        let _1606_v1;
        _1606_v1 = _module.D0.create_DC0(_module.__default.fm4(_1604_i0, _1604_i0, (_this).f6, globalState), new BigNumber((_dafny.Seq.Concat(_1605_v0, _1605_v0)).length));
        _1606_v1 = _1606_v1;
        let _1607_v2;
        _1607_v2 = _dafny.MultiSet.fromElements(_this.f9);
        if ((_1607_v2).IsProperSubsetOf(_1607_v2)) {
          let _1608_v3;
          let _nw235 = new _module.C0();
          _nw235.__ctor(_this.f9, p0);
          _1608_v3 = _nw235;
          (globalState).f5 = (_this).f6;
          (globalState).f5 = ((p2).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vsvuef")).length))).plus(p2);
          let _1609_v4;
          let _init42 = ((_1610_v3) => function (_1611_i1) {
            return _dafny.Seq.of(_1610_v3.f20);
          })(_1608_v3);
          let _nw236 = Array((new BigNumber(16)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw236.length); _i0_42++) {
            _nw236[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1609_v4 = _nw236;
          let _index228 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_1609_v4).length));
          (_1609_v4)[_index228] = _dafny.Seq.Concat(_dafny.Seq.of(_1608_v3.f20, _1608_v3.f20), p1);
          let _1612_v5;
          _1612_v5 = _dafny.Seq.of(p0);
          let _1613_v6;
          _1613_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(551)), function (_1614_i2) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })).length),_dafny.Seq.of(!(_1608_v3.f20), _1608_v3.f20));
          let _index229 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_1609_v4).length));
          let _rhs263 = _module.__default.fm3(globalState);
          let _rhs264 = (((_1613_v6).contains((_this).f6)) ? ((_1613_v6).get((_this).f6)) : (_1612_v5));
          let _rhs265 = p1;
          let _rhs266 = _this.f12;
          let _rhs267 = _1604_i0;
          let _lhs216 = globalState;
          let _lhs217 = _1609_v4;
          let _lhs218 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_1609_v4).length));
          let _lhs219 = _this;
          let _lhs220 = globalState;
          _lhs216.f0 = _rhs263;
          _lhs217[_lhs218] = _rhs264;
          _1612_v5 = _rhs265;
          _lhs219.f11 = _rhs266;
          _lhs220.f5 = _rhs267;
          let _1615_v8;
          _1615_v8 = _dafny.Set.fromElements(false, _1608_v3.f20, _1608_v3.f20);
          let _1616_v9;
          _1616_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(930),(_module.__default.fm2(_1604_i0, globalState)).isEqualTo(new BigNumber((function () {
            let _coll45 = new _dafny.Map();
            for (const _compr_45 of (_dafny.Seq.of(_1615_v8, _1615_v8)).Elements) {
              let _1617_v7 = _compr_45;
              if (_dafny.Seq.contains(_dafny.Seq.of(_1615_v8, _1615_v8), _1617_v7)) {
                _coll45.push([_1617_v7,_1608_v3.f20]);
              }
            }
            return _coll45;
          }()).length)));
          (_1608_v3).f20 = (((_1616_v9).contains(_module.__default.safeModuloInt(p2, new BigNumber((_1605_v0).length)))) ? ((_1616_v9).get(_module.__default.safeModuloInt(p2, new BigNumber((_1605_v0).length)))) : (p0));
        } else {
          let _1618_v10;
          _1618_v10 = _dafny.MultiSet.fromElements(p2);
          let _1619_v11;
          _1619_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1618_v10,p2);
          _1619_v11 = _1619_v11;
          let _1620_v12;
          _1620_v12 = _module.D2.create_DC6(new BigNumber(835), _this.f11);
          let _1621_v13;
          let _init43 = function (_1622_i3) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(957)), function (_1623_i4) {
              return (_this).f6;
            });
          };
          let _nw237 = Array((new BigNumber(28)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw237.length); _i0_43++) {
            _nw237[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1621_v13 = _nw237;
          let _1624_v14;
          _1624_v14 = _dafny.Seq.of(new BigNumber((_1618_v10).cardinality()));
          let _index230 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_1621_v13).length));
          (_1621_v13)[_index230] = _1624_v14;
          let _1625_v15;
          let _nw238 = Array((new BigNumber(20)).toNumber()).fill(_module.D7.Default());
          _1625_v15 = _nw238;
          let _pat_let_tv32 = p1;
          let _1626_v16;
          let _nw239 = Array((new BigNumber(16)).toNumber());
          _nw239[(_dafny.ZERO).toNumber()] = _1606_v1;
          _nw239[(_dafny.ONE).toNumber()] = _module.D0.create_DC0(p1, new BigNumber((_1605_v0).length));
          _nw239[(new BigNumber(2)).toNumber()] = _1606_v1;
          _nw239[(new BigNumber(3)).toNumber()] = _1606_v1;
          _nw239[(new BigNumber(4)).toNumber()] = _1606_v1;
          _nw239[(new BigNumber(5)).toNumber()] = _module.__default.fm32(p0, p0, p2, p0, globalState);
          _nw239[(new BigNumber(6)).toNumber()] = _1606_v1;
          _nw239[(new BigNumber(7)).toNumber()] = _1606_v1;
          _nw239[(new BigNumber(8)).toNumber()] = _1606_v1;
          _nw239[(new BigNumber(9)).toNumber()] = _module.__default.fm32(p0, p0, _1604_i0, p0, globalState);
          _nw239[(new BigNumber(10)).toNumber()] = _1606_v1;
          _nw239[(new BigNumber(11)).toNumber()] = _module.D0.create_DC0(_dafny.Seq.of(p0), _1604_i0);
          _nw239[(new BigNumber(12)).toNumber()] = _1606_v1;
          _nw239[(new BigNumber(13)).toNumber()] = function (_pat_let43_0) {
            return function (_1627_dt__update__tmp_h0) {
              return function (_pat_let44_0) {
                return function (_1628_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_1628_dt__update_hcf0_h0, (_1627_dt__update__tmp_h0).dtor_cf1);
                }(_pat_let44_0);
              }(_pat_let_tv32);
            }(_pat_let43_0);
          }(_1606_v1);
          _nw239[(new BigNumber(14)).toNumber()] = _1606_v1;
          _nw239[(new BigNumber(15)).toNumber()] = _1606_v1;
          _1626_v16 = _nw239;
          let _1629_v17;
          _1629_v17 = _module.D7.create_DC21(_1626_v16);
          let _index231 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1625_v15).length));
          (_1625_v15)[_index231] = _1629_v17;
          let _1630_v18;
          _1630_v18 = _dafny.Seq.of(_this.f12, _this.f11);
          let _index232 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_1621_v13).length));
          let _index233 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1625_v15).length));
          let _rhs268 = _1620_v12;
          let _rhs269 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length)), _module.__default.safeIndex(new BigNumber(65), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length))).length)), new BigNumber(848)), _dafny.Seq.Concat(_1624_v14, _1624_v14));
          let _rhs270 = _1629_v17;
          let _rhs271 = p0;
          let _rhs272 = (_1630_v18)[_module.__default.safeIndex(p2, new BigNumber((_1630_v18).length))];
          let _lhs221 = _1621_v13;
          let _lhs222 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_1621_v13).length));
          let _lhs223 = _1625_v15;
          let _lhs224 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1625_v15).length));
          let _lhs225 = globalState;
          let _lhs226 = _this;
          _1620_v12 = _rhs268;
          _lhs221[_lhs222] = _rhs269;
          _lhs223[_lhs224] = _rhs270;
          _lhs225.f0 = _rhs271;
          _lhs226.f11 = _rhs272;
          let _1631_v19;
          _1631_v19 = _dafny.MultiSet.fromElements(_1618_v10);
          let _1632_v20;
          let _nw240 = new _module.C2();
          _nw240.__ctor(((_this).f6).plus(_1604_i0), ((!(p0)) ? (_1631_v19) : (_1631_v19)), _this.f9, (_this).f10);
          _1632_v20 = _nw240;
          let _1633_v21;
          _1633_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1605_v0);
          _1633_v21 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.UnicodeFromString("dkcllx"));
          (globalState).f0 = ((p0) ? (p0) : (p0));
        }
      }
      let _arr46 = _this.f9;
      let _index234 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length));
      _arr46[_index234] = p0;
      let _1634_v22;
      _1634_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6);
      let _arr47 = _this.f9;
      let _index235 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length));
      _arr47[_index235] = ((p0) ? (!((new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,p0), _dafny.Map.Empty.slice().updateUnsafe(p0,p0))).length)).isEqualTo(new BigNumber((_1634_v22).length)))) : (((p0) ? (p0) : (p0))));
      if (p0) {
        let _1635_v23;
        _1635_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,!(p0));
        let _1636_v24;
        _1636_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1635_v23,(_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
        let _1637_v25;
        _1637_v25 = _module.D3.create_DC10(_1636_v24);
        let _1638_v26;
        _1638_v26 = _module.D3.create_DC11(_1637_v25);
        let _1639_v27;
        _1639_v27 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
        _1638_v26 = _module.__default.fm50((((_1639_v27).contains((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])) ? ((_1639_v27).get((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])) : (p1)), globalState);
        (globalState).f5 = (_dafny.ZERO).minus((_this).f6);
        let _1640_v28;
        _1640_v28 = _dafny.Seq.UnicodeFromString("bvftkwuvg");
        (globalState).f5 = _module.__default.safeModuloInt((new BigNumber((_1640_v28).length)).plus((_this).f6), p2);
        if ((((_this).f6).multipliedBy((_this).f6)).isLessThan(p2)) {
          let _1641_v29;
          _1641_v29 = _dafny.Seq.of(_this.f11);
          let _1642_v30;
          _1642_v30 = _dafny.Map.Empty.slice().updateUnsafe(!((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]),p0);
          let _1643_v31;
          _1643_v31 = _module.D4.create_DC14((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], _1642_v30, (_this).f6, _1642_v30);
          let _1644_v32;
          _1644_v32 = _dafny.MultiSet.fromElements((_1643_v31).dtor_cf28);
          let _1645_v33;
          _1645_v33 = _dafny.Seq.of((((_1644_v32).contains((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])) ? ((_1644_v32).get((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])) : ((_this).f6)), (_this).f6);
          let _rhs273 = _1641_v29;
          let _rhs274 = ((_1645_v33)[_module.__default.safeIndex((_this).f6, new BigNumber((_1645_v33).length))]).isLessThan(_module.__default.fm2(p2, globalState));
          let _lhs227 = globalState;
          _1641_v29 = _rhs273;
          _lhs227.f0 = _rhs274;
          let _1646_v34;
          let _nw241 = Array((new BigNumber(15)).toNumber()).fill(false);
          _1646_v34 = _nw241;
          let _1647_v35;
          _1647_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1646_v34,(_this).f6);
          (globalState).f5 = (((_1647_v35).contains(_1646_v34)) ? ((_1647_v35).get(_1646_v34)) : (p2));
          _1645_v33 = _dafny.Seq.update(_dafny.Seq.Concat(_1645_v33, _dafny.Seq.of(_module.__default.fm2(p2, globalState))), _module.__default.safeIndex(_module.__default.fm2(new BigNumber((p1).length), globalState), new BigNumber((_dafny.Seq.Concat(_1645_v33, _dafny.Seq.of(_module.__default.fm2(p2, globalState)))).length)), (new BigNumber(745)).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(94)), function (_1648_i5) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          })).length)));
          let _arr48 = _this.f9;
          let _index236 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length));
          _arr48[_index236] = p0;
          let _1649_v36;
          _1649_v36 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f6));
          let _1650_v37;
          _1650_v37 = _dafny.MultiSet.fromElements(_1649_v36, _1649_v36, _1649_v36);
          let _1651_v38;
          let _nw242 = new _module.C2();
          _nw242.__ctor((_dafny.ZERO).minus(new BigNumber((_1635_v23).length)), (_1650_v37).Difference(_1650_v37), _1646_v34, (_dafny.Map.Empty.slice().updateUnsafe(p2,((((_this).f10).contains(new BigNumber((_1640_v28).length))) ? (((_this).f10).get(new BigNumber((_1640_v28).length))) : (_1640_v28)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6)).length),_1640_v28)));
          _1651_v38 = _nw242;
        } else {
          let _1652_v39;
          let _nw243 = new _module.C6();
          _nw243.__ctor(_module.__default.safeModuloInt((_dafny.ZERO).minus(p2), (_this).f6), (((_1635_v23).contains(p2)) ? ((_1635_v23).get(p2)) : ((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])));
          _1652_v39 = _nw243;
          _1652_v39 = _1652_v39;
          let _1653_v40;
          _1653_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber(-125));
          let _1654_v41;
          _1654_v41 = _dafny.Seq.of((_1652_v39).f13);
          let _1655_v42;
          let _nw244 = Array((new BigNumber(16)).toNumber());
          _nw244[(_dafny.ZERO).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw244[(_dafny.ONE).toNumber()] = p0;
          _nw244[(new BigNumber(2)).toNumber()] = p0;
          _nw244[(new BigNumber(3)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw244[(new BigNumber(4)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber((_1654_v41).length), new BigNumber((p1).length))];
          _nw244[(new BigNumber(5)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw244[(new BigNumber(6)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw244[(new BigNumber(7)).toNumber()] = p0;
          _nw244[(new BigNumber(8)).toNumber()] = p0;
          _nw244[(new BigNumber(9)).toNumber()] = p0;
          _nw244[(new BigNumber(10)).toNumber()] = p0;
          _nw244[(new BigNumber(11)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw244[(new BigNumber(12)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw244[(new BigNumber(13)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw244[(new BigNumber(14)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw244[(new BigNumber(15)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _1655_v42 = _nw244;
          let _1656_v43;
          let _nw245 = new _module.C3();
          _nw245.__ctor(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber((_1653_v40).length)), _1655_v42, (_this).f10, (_1654_v41)[_module.__default.safeIndex((_1652_v39).f13, new BigNumber((_1654_v41).length))]);
          _1656_v43 = _nw245;
          let _1657_v44;
          _1657_v44 = _dafny.Set.fromElements((_1652_v39).f13, (_this).f6);
          let _1658_v45;
          _1658_v45 = _module.D11.create_DC30(_1657_v44);
          _1658_v45 = _1658_v45;
          let _1659_v46;
          let _nw246 = new _module.C5();
          _nw246.__ctor(((_this).f6).multipliedBy((_1652_v39).f13), new BigNumber(690));
          _1659_v46 = _nw246;
          let _1660_v47;
          let _nw247 = Array((new BigNumber(23)).toNumber()).fill(_module.D3.Default());
          _1660_v47 = _nw247;
          let _pat_let_tv33 = _1637_v25;
          let _index237 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_1660_v47).length));
          (_1660_v47)[_index237] = function (_pat_let45_0) {
            return function (_1661_dt__update__tmp_h1) {
              return function (_pat_let46_0) {
                return function (_1662_dt__update_hcf23_h0) {
                  return _module.D3.create_DC11(_1662_dt__update_hcf23_h0);
                }(_pat_let46_0);
              }(_pat_let_tv33);
            }(_pat_let45_0);
          }(_1638_v26);
          let _index238 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_1660_v47).length));
          (_1660_v47)[_index238] = _1638_v26;
        }
        (globalState).f5 = (_module.D1.create_DC4(_module.__default.fm32(p0, p0, (_this).f6, (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], globalState), p2, _1640_v28, (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], p2)).dtor_cf8;
      } else {
        let _1663_v48;
        _1663_v48 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1664_v49;
        _1664_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1663_v48,!(p0));
        _1664_v49 = (_1664_v49).update(_1663_v48, p0);
        let _1665_v50;
        _1665_v50 = _dafny.MultiSet.fromElements(p0);
        let _1666_v51;
        _1666_v51 = _dafny.Set.fromElements((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
        (globalState).f5 = (_dafny.ZERO).minus(((((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]) ? (p2) : ((((_1665_v50).contains((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])) ? ((_1665_v50).get((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])) : ((_this).f6))))).minus(((true) ? ((_this).f6) : (new BigNumber((_1666_v51).length)))));
        let _1667_v52;
        _1667_v52 = _module.D0.create_DC1((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], p2, _1663_v48);
        let _rhs275 = (_1667_v52).dtor_cf4;
        let _rhs276 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(p2)));
        let _lhs228 = globalState;
        _1663_v48 = _rhs275;
        _lhs228.f5 = _rhs276;
        let _rhs277 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(p2, (_this).fm14(globalState)), (_this).f6);
        let _rhs278 = ((_this).f6).minus(((_this).f6).multipliedBy(p2));
        let _lhs229 = globalState;
        let _lhs230 = globalState;
        _lhs229.f5 = _rhs277;
        _lhs230.f5 = _rhs278;
        let _1668_v53;
        _1668_v53 = _dafny.MultiSet.fromElements(_1663_v48);
        let _1669_v54;
        _1669_v54 = _dafny.Seq.of(_1668_v53, _1668_v53, _1668_v53);
        let _1670_v55;
        _1670_v55 = _module.D4.create_DC12(_1669_v54);
        let _1671_v56;
        _1671_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1670_v55,p2);
        let _1672_v57;
        _1672_v57 = _dafny.Seq.UnicodeFromString("ogh");
        let _arr49 = _this.f9;
        let _index239 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length));
        let _rhs279 = (((new BigNumber((_1671_v56).length)).isLessThanOrEqualTo(_module.__default.fm2(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f6, (_this).f6))).cardinality()), globalState))) ? ((!(!(p0))) === (_module.__default.fm3(globalState))) : (!(!(!((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])))));
        let _rhs280 = _1672_v57;
        let _lhs231 = _this.f9;
        let _lhs232 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length));
        let _lhs233 = globalState;
        _lhs231[_lhs232] = _rhs279;
        _lhs233.f2 = _rhs280;
      }
      let _1673_v58;
      _1673_v58 = _dafny.MultiSet.fromElements((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
      let _1674_v59;
      _1674_v59 = _module.D6.create_DC19((_this).f6, (_this).f6, new BigNumber(428));
      let _1675_v60;
      _1675_v60 = _dafny.Map.Empty.slice().updateUnsafe((((_1673_v58).contains((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])) ? ((_1673_v58).get((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])) : ((_dafny.ZERO).minus(p2))),(_1674_v59).dtor_cf41);
      _1675_v60 = (function () {
        let _coll46 = new _dafny.Map();
        for (const _compr_46 of (_dafny.Seq.update(_dafny.Seq.of((_this).f6), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.of((_this).f6)).length)), new BigNumber((_1675_v60).length))).Elements) {
          let _1676_v61 = _compr_46;
          if (_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.of((_this).f6), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.of((_this).f6)).length)), new BigNumber((_1675_v60).length)), _1676_v61)) {
            _coll46.push([_module.__default.safeDivisionInt(_1676_v61, (_this).f6),p2]);
          }
        }
        return _coll46;
      }()).update(p2, (_this).f6);
      let _1677_v62;
      _1677_v62 = new _dafny.CodePoint('y'.codePointAt(0));
      (globalState).f5 = (new BigNumber((_dafny.Seq.of(_1677_v62)).length)).plus((_this).f6);
      let _1678_v63;
      let _init44 = ((_1679_p0) => function (_1680_i6) {
        return _dafny.Map.Empty.slice().updateUnsafe(_1679_p0,(_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
      })(p0);
      let _nw248 = Array((new BigNumber(19)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw248.length); _i0_44++) {
        _nw248[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1678_v63 = _nw248;
      let _1681_v64;
      _1681_v64 = _module.D9.create_DC26(_1678_v63);
      let _source20 = _1681_v64;
      if (_source20.is_DC25) {
        let _1682___mcc_h0 = (_source20).cf46;
        let _1683___mcc_h1 = (_source20).cf47;
        let _1684___mcc_h2 = (_source20).cf48;
        let _1685___mcc_h3 = (_source20).cf49;
        let _1686_cf49 = _1685___mcc_h3;
        let _1687_cf48 = _1684___mcc_h2;
        let _1688_cf47 = _1683___mcc_h1;
        let _1689_cf46 = _1682___mcc_h0;
        let _1690_v65;
        _1690_v65 = _dafny.Seq.of(_1689_cf46);
        let _1691_v66;
        _1691_v66 = _dafny.MultiSet.fromElements(new BigNumber((_1690_v65).length), _1688_cf47);
        let _1692_v67;
        _1692_v67 = _dafny.Seq.of(_1686_cf49);
        let _1693_v68;
        _1693_v68 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1691_v66).cardinality())).plus(_1688_cf47),_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(50)), ((_1694_cf49) => function (_1695_i7) {
          return _1694_cf49;
        })(_1686_cf49)), _1692_v67));
        _1693_v68 = (_1693_v68).update((_1690_v65)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-709)), new BigNumber((_1690_v65).length))], _1692_v67);
        if (p0) {
          let _1696_v69;
          _1696_v69 = _module.D2.create_DC6(_1687_cf48, _this.f11);
          let _1697_v70;
          _1697_v70 = _dafny.Seq.of((_1696_v69).dtor_cf14, _this.f11);
          _1697_v70 = _dafny.Seq.of(_this.f12, _this.f12, _this.f11);
          (globalState).f0 = p0;
          let _1698_v71;
          _1698_v71 = _dafny.Seq.of(_1690_v65);
          let _1699_v72;
          _1699_v72 = _dafny.Set.fromElements((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
          let _1700_v73;
          let _nw249 = Array((new BigNumber(28)).toNumber());
          _nw249[(_dafny.ZERO).toNumber()] = _1690_v65;
          _nw249[(_dafny.ONE).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(2)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(3)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-787)), function (_1701_i8) {
            return (_this).f6;
          });
          _nw249[(new BigNumber(5)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(6)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_1690_v65, _module.__default.safeIndex(_module.__default.fm2(p2, globalState), new BigNumber((_1690_v65).length)), (_this).f6);
          _nw249[(new BigNumber(8)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(9)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-670)), function (_1702_i9) {
            return new BigNumber(-939);
          });
          _nw249[(new BigNumber(11)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(12)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(13)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(14)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(15)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(16)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(17)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(new BigNumber(-584));
          _nw249[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(_1688_cf47, p2);
          _nw249[(new BigNumber(20)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(21)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(22)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(23)).toNumber()] = _dafny.Seq.update((_1698_v71)[_module.__default.safeIndex(p2, new BigNumber((_1698_v71).length))], _module.__default.safeIndex(_1688_cf47, new BigNumber(((_1698_v71)[_module.__default.safeIndex(p2, new BigNumber((_1698_v71).length))]).length)), _1688_cf47);
          _nw249[(new BigNumber(24)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(25)).toNumber()] = _1690_v65;
          _nw249[(new BigNumber(26)).toNumber()] = _dafny.Seq.update(_1690_v65, _module.__default.safeIndex(_1689_cf46, new BigNumber((_1690_v65).length)), new BigNumber((_1699_v72).length));
          _nw249[(new BigNumber(27)).toNumber()] = _dafny.Seq.of((_this).f6, p2);
          _1700_v73 = _nw249;
          let _1703_v74;
          _1703_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1700_v73,_1692_v67);
          _1703_v74 = (_1703_v74).update(_1700_v73, _1692_v67);
          let _1704_v75;
          _1704_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1677_v62,new BigNumber(-779));
          _1704_v75 = (_1704_v75).Merge(_1704_v75);
          (globalState).f0 = p0;
        } else {
          let _1705_v76;
          _1705_v76 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _1706_v77;
          _1706_v77 = _module.D4.create_DC15((((_1705_v76).contains((_dafny.ZERO).minus(_1687_cf48))) ? ((_1705_v76).get((_dafny.ZERO).minus(_1687_cf48))) : (p0)), p0, _1692_v67);
          r1 = !((_1706_v77).dtor_cf33);
          let _1707_v78;
          let _nw250 = Array((new BigNumber(15)).toNumber());
          _nw250[(_dafny.ZERO).toNumber()] = true;
          _nw250[(_dafny.ONE).toNumber()] = false;
          _nw250[(new BigNumber(2)).toNumber()] = p0;
          _nw250[(new BigNumber(3)).toNumber()] = !(p0);
          _nw250[(new BigNumber(4)).toNumber()] = !((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
          _nw250[(new BigNumber(5)).toNumber()] = _module.__default.fm3(globalState);
          _nw250[(new BigNumber(6)).toNumber()] = p0;
          _nw250[(new BigNumber(7)).toNumber()] = (_1706_v77).dtor_cf33;
          _nw250[(new BigNumber(8)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw250[(new BigNumber(9)).toNumber()] = p0;
          _nw250[(new BigNumber(10)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _nw250[(new BigNumber(11)).toNumber()] = p0;
          _nw250[(new BigNumber(12)).toNumber()] = p0;
          _nw250[(new BigNumber(13)).toNumber()] = p0;
          _nw250[(new BigNumber(14)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
          _1707_v78 = _nw250;
          let _1708_v79;
          let _nw251 = new _module.C3();
          _nw251.__ctor((_1675_v60).Merge(_1675_v60), _1707_v78, (_1693_v68).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("thytt"), _module.__default.safeIndex(_1689_cf46, new BigNumber((_dafny.Seq.UnicodeFromString("thytt")).length)), _1686_cf49))), _1687_cf48);
          _1708_v79 = _nw251;
          (globalState).f5 = _1688_cf47;
          let _1709_v80;
          _1709_v80 = _dafny.Map.Empty.slice().updateUnsafe((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))],_1686_cf49);
          let _1710_v81;
          _1710_v81 = _dafny.Set.fromElements(_1688_cf47, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(752)), ((_1711_cf46) => function (_1712_i10) {
            return _1711_cf46;
          })(_1689_cf46))).length)));
          _1686_cf49 = _module.__default.fm0(new BigNumber((_1709_v80).length), _1710_v81, (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], globalState);
          let _1713_v82;
          _1713_v82 = _dafny.MultiSet.fromElements((_1691_v66).update(new BigNumber(772), _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), ((_1714_cf49) => function (_1715_i11) {
            return _1714_cf49;
          })(_1686_cf49))).length)))), _1691_v66);
          let _1716_v83;
          let _nw252 = new _module.C2();
          _nw252.__ctor(_1689_cf46, _1713_v82, _1707_v78, _1693_v68);
          _1716_v83 = _nw252;
          let _1717_v84;
          _1717_v84 = _dafny.Seq.of(_1716_v83, _1716_v83, _1716_v83);
          let _1718_v85;
          let _nw253 = new _module.C3();
          _nw253.__ctor((_1675_v60).Merge(_1675_v60), ((p0) ? (_1707_v78) : (_1707_v78)), ((false) ? ((_this).f10) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1692_v67).length),_dafny.Seq.update(_1692_v67, _module.__default.safeIndex(new BigNumber((_1717_v84).length), new BigNumber((_1692_v67).length)), new _dafny.CodePoint('j'.codePointAt(0)))))), (_dafny.ZERO).minus(new BigNumber((_1673_v58).cardinality())));
          _1718_v85 = _nw253;
          let _nw254 = new _module.C2();
          _nw254.__ctor((_1716_v83).f6, _1713_v82, _1718_v85.f9, ((_this).f10).Merge((_1718_v85).f10));
          _1718_v85 = _nw254;
        }
        let _arr50 = _this.f9;
        let _index240 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length));
        _arr50[_index240] = !(((_module.__default.fm3(globalState)) ? (new BigNumber(93)) : (p2))).isEqualTo(_1688_cf47);
        let _arr51 = _this.f9;
        let _index241 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length));
        _arr51[_index241] = !(!(((new BigNumber((p1).length)).minus(_1687_cf48)).isEqualTo(new BigNumber((_1673_v58).cardinality()))));
      } else if (_source20.is_DC26) {
        let _1719___mcc_h4 = (_source20).cf50;
        let _1720_cf50 = _1719___mcc_h4;
        let _1721_v86;
        _1721_v86 = _dafny.Seq.of(_1677_v62);
        let _1722_v87;
        _1722_v87 = _dafny.Set.fromElements(p0, true);
        let _1723_v88;
        _1723_v88 = _module.D11.create_DC33(p2, p0, new BigNumber((_1721_v86).length), (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], _1722_v87);
        if ((_1723_v88).dtor_cf65) {
          (globalState).f5 = (_this).f6;
          let _1724_v89;
          _1724_v89 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1721_v86).length),_1677_v62);
          let _1725_v90;
          _1725_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1724_v89).length),(_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
          let _1726_v91;
          _1726_v91 = _dafny.Map.Empty.slice().updateUnsafe((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))],(_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
          _1725_v90 = (_1725_v90).update(p2, (_1726_v91).contains((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]));
          (globalState).f0 = (p0) && (!_dafny.Seq.contains(p1, (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]));
          let _rhs281 = p2;
          let _rhs282 = ((_module.D11.create_DC31(new BigNumber(-330), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])).length))).dtor_cf59).multipliedBy((_dafny.ZERO).minus(new BigNumber((_1722_v87).length)));
          let _lhs234 = globalState;
          let _lhs235 = globalState;
          _lhs234.f5 = _rhs281;
          _lhs235.f5 = _rhs282;
          (globalState).f5 = (_this).f6;
        } else {
          let _1727_v92;
          let _nw255 = Array((new BigNumber(26)).toNumber()).fill(_module.D2.Default());
          _1727_v92 = _nw255;
          let _1728_v93;
          _1728_v93 = _module.D2.create_DC8((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
          let _index242 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1727_v92).length));
          (_1727_v92)[_index242] = _1728_v93;
          let _index243 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1727_v92).length));
          (_1727_v92)[_index243] = ((p0) ? (_1728_v93) : (_1728_v93));
          (globalState).f5 = (p2).plus((_this).f6);
          let _1729_v94;
          let _nw256 = new _module.C1();
          _nw256.__ctor(new BigNumber((_1721_v86).length));
          _1729_v94 = _nw256;
          let _1730_v95;
          let _nw257 = Array((new BigNumber(9)).toNumber()).fill([]);
          _1730_v95 = _nw257;
          let _arr52 = _this.f11;
          let _index244 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_this.f11).length));
          _arr52[_index244] = (_dafny.ZERO).minus(p2);
          let _1731_v96;
          _1731_v96 = _dafny.MultiSet.fromElements(p1, _dafny.Seq.update(p1, _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1721_v86, _module.__default.safeIndex(p2, new BigNumber((_1721_v86).length)), _1677_v62)).length), new BigNumber((p1).length)), (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]), p1, p1);
          let _arr53 = _this.f11;
          let _index245 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_this.f11).length));
          let _rhs283 = _1730_v95;
          let _rhs284 = new BigNumber((((_module.__default.fm3(globalState)) ? ((_1731_v96).update(p1, _module.__default.abs(new BigNumber(569)))) : (_1731_v96))).cardinality());
          let _lhs236 = _this.f11;
          let _lhs237 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_this.f11).length));
          _1730_v95 = _rhs283;
          _lhs236[_lhs237] = _rhs284;
          let _1732_v97;
          _1732_v97 = _dafny.Seq.of(_this.f12);
          let _1733_v98;
          _1733_v98 = _dafny.Seq.of((_this).f6, p2);
          let _1734_v99;
          let _nw258 = Array((new BigNumber(9)).toNumber());
          _nw258[(_dafny.ZERO).toNumber()] = _this.f12;
          _nw258[(_dafny.ONE).toNumber()] = _this.f12;
          _nw258[(new BigNumber(2)).toNumber()] = _this.f12;
          _nw258[(new BigNumber(3)).toNumber()] = _this.f12;
          _nw258[(new BigNumber(4)).toNumber()] = _this.f12;
          _nw258[(new BigNumber(5)).toNumber()] = _this.f12;
          _nw258[(new BigNumber(6)).toNumber()] = _this.f12;
          _nw258[(new BigNumber(7)).toNumber()] = _this.f12;
          _nw258[(new BigNumber(8)).toNumber()] = (_1732_v97)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1733_v98).length)))), new BigNumber((_1732_v97).length))];
          _1734_v99 = _nw258;
          let _1735_v101;
          _1735_v101 = _dafny.Set.fromElements((_this.f11)[_module.__default.safeIndex(new BigNumber(471), new BigNumber((_this.f11).length))]);
          let _1736_v102;
          _1736_v102 = _dafny.Seq.of((function () {
            let _coll47 = new _dafny.Set();
            for (const _compr_47 of _dafny.IntegerRange(new BigNumber(-932), new BigNumber(-282))) {
              let _1737_v100 = _compr_47;
              if (((new BigNumber(-932)).isLessThanOrEqualTo(_1737_v100)) && ((_1737_v100).isLessThan(new BigNumber(-282)))) {
                _coll47.add(_module.__default.safeModuloInt(_1737_v100, p2));
              }
            }
            return _coll47;
          }()).IsProperSubsetOf(_1735_v101), (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], !((_this).f6).isEqualTo(new BigNumber(-21)));
          let _1738_v103;
          let _nw259 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _1738_v103 = _nw259;
          let _arr54 = _this.f11;
          let _index246 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_this.f11).length));
          let _rhs285 = new BigNumber((_dafny.Seq.UnicodeFromString("psadyloiw")).length);
          let _rhs286 = _1734_v99;
          let _rhs287 = _module.__default.fm4(p2, (new BigNumber(-846)).multipliedBy(new BigNumber(17)), ((p0) ? ((_this.f11)[_module.__default.safeIndex(new BigNumber(471), new BigNumber((_this.f11).length))]) : ((_this).f6)), globalState);
          let _rhs288 = _1738_v103;
          let _lhs238 = _this.f11;
          let _lhs239 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_this.f11).length));
          _lhs238[_lhs239] = _rhs285;
          _1734_v99 = _rhs286;
          _1736_v102 = _rhs287;
          _1738_v103 = _rhs288;
        }
        let _1739_v104;
        _1739_v104 = _dafny.MultiSet.fromElements(new BigNumber(105));
        let _1740_v105;
        _1740_v105 = _dafny.MultiSet.fromElements(_1739_v104, _1739_v104);
        let _1741_v106;
        let _nw260 = Array((new BigNumber(8)).toNumber());
        _nw260[(_dafny.ZERO).toNumber()] = p0;
        _nw260[(_dafny.ONE).toNumber()] = p0;
        _nw260[(new BigNumber(2)).toNumber()] = p0;
        _nw260[(new BigNumber(3)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
        _nw260[(new BigNumber(4)).toNumber()] = false;
        _nw260[(new BigNumber(5)).toNumber()] = p0;
        _nw260[(new BigNumber(6)).toNumber()] = p0;
        _nw260[(new BigNumber(7)).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
        _1741_v106 = _nw260;
        let _1742_v107;
        let _nw261 = new _module.C2();
        _nw261.__ctor((_this).f6, _1740_v105, _1741_v106, (_this).f10);
        _1742_v107 = _nw261;
        let _1743_v108;
        _1743_v108 = _dafny.Set.fromElements(_1742_v107, _1742_v107, _1742_v107);
        let _1744_v109;
        _1744_v109 = _dafny.Map.Empty.slice().updateUnsafe(_1743_v108,new _dafny.CodePoint('g'.codePointAt(0)));
        _1744_v109 = ((_module.__default.fm3(globalState)) ? (_1744_v109) : (_1744_v109));
        (globalState).f5 = p2;
        let _arr55 = _this.f9;
        let _index247 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length));
        _arr55[_index247] = !(p0);
      } else if (_source20.is_DC24) {
        let _1745___mcc_h5 = (_source20).cf45;
        let _1746_cf45 = _1745___mcc_h5;
        let _arr56 = _this.f9;
        let _index248 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length));
        _arr56[_index248] = p0;
        let _1747_v110;
        _1747_v110 = _dafny.Set.fromElements(p0, p0, p0);
        let _1748_v111;
        let _nw262 = Array((new BigNumber(6)).toNumber());
        _nw262[(_dafny.ZERO).toNumber()] = false;
        _nw262[(_dafny.ONE).toNumber()] = (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))];
        _nw262[(new BigNumber(2)).toNumber()] = p0;
        _nw262[(new BigNumber(3)).toNumber()] = !((_1747_v110).IsProperSubsetOf(_1747_v110));
        _nw262[(new BigNumber(4)).toNumber()] = false;
        _nw262[(new BigNumber(5)).toNumber()] = (new BigNumber((p1).length)).isEqualTo((_this).f6);
        _1748_v111 = _nw262;
        let _1749_v112;
        _1749_v112 = _dafny.Seq.of(((p0) ? (_1748_v111) : (_1748_v111)));
        _1748_v111 = (_1749_v112)[_module.__default.safeIndex((_this).f6, new BigNumber((_1749_v112).length))];
        (globalState).f5 = (_this).f6;
        if (((_this).f6).isLessThanOrEqualTo((_this).f6)) {
          let _1750_v113;
          let _nw263 = new _module.C0();
          _nw263.__ctor(_1748_v111, (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
          _1750_v113 = _nw263;
          let _1751_v114;
          _1751_v114 = _dafny.Seq.UnicodeFromString("hqae");
          let _1752_v115;
          _1752_v115 = _module.D9.create_DC25(p2, _module.__default.safeDivisionInt(new BigNumber((_1751_v114).length), p2), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(125)), ((_1753_p2) => function (_1754_i12) {
  return _module.D3.create_DC9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(266)), ((_1755_p2) => function (_1756_i13) {
  return _1755_p2;
})(_1753_p2)));
})(p2))).length)), _1677_v62);
          _1752_v115 = _1752_v115;
          let _1757_v116;
          _1757_v116 = _dafny.Seq.of(_dafny.Set.fromElements(p0));
          let _1758_v117;
          _1758_v117 = _dafny.Map.Empty.slice().updateUnsafe(((_1757_v116)[_module.__default.safeIndex(new BigNumber((_1751_v114).length), new BigNumber((_1757_v116).length))]).Union(_1747_v110),!((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]));
          _1758_v117 = (_1758_v117).update(_1747_v110, _1750_v113.f20);
          let _1759_v118;
          _1759_v118 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          _1759_v118 = (_1759_v118).Merge((_1759_v118).Merge(function () {
            let _coll48 = new _dafny.Map();
            for (const _compr_48 of (_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1675_v60).length))).Keys.Elements) {
              let _1760_v119 = _compr_48;
              if ((_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1675_v60).length))).contains(_1760_v119)) {
                _coll48.push([_module.__default.safeDivisionInt(_1760_v119, p2),p0]);
              }
            }
            return _coll48;
          }()));
          let _1761_v120;
          _1761_v120 = _dafny.Seq.of(p2);
          let _1762_v121;
          _1762_v121 = _dafny.Map.Empty.slice().updateUnsafe(_1750_v113.f20,_1750_v113.f20);
          (globalState).f5 = (((_1761_v120)[_module.__default.safeIndex(new BigNumber((_1762_v121).length), new BigNumber((_1761_v120).length))]).multipliedBy((_this).f6)).multipliedBy(p2);
        } else {
          let _1763_v122;
          _1763_v122 = _dafny.MultiSet.fromElements(p2, p2);
          let _arr57 = _this.f12;
          let _index249 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_this.f12).length));
          _arr57[_index249] = ((((_1763_v122).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(998)), function (_1765_i14) {
            return (_this).f6;
          })).length))) ? ((_1763_v122).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(998)), function (_1764_i14) {
            return (_this).f6;
          })).length))) : (p2))).minus(p2);
          let _arr58 = _this.f12;
          let _index250 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_this.f12).length));
          _arr58[_index250] = (_this).f6;
          let _1766_v123;
          let _nw264 = new _module.C6();
          _nw264.__ctor((_this).f6, (((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]) ? ((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]) : ((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])));
          _1766_v123 = _nw264;
          let _1767_v124;
          let _nw265 = Array((new BigNumber(26)).toNumber()).fill(_module.D6.Default());
          _1767_v124 = _nw265;
          let _1768_v125;
          _1768_v125 = _1767_v124;
          _1767_v124 = (_1768_v125);
          let _1769_v126;
          _1769_v126 = (_this).f10;
          let _1770_v127;
          let _nw266 = new _module.C3();
          _nw266.__ctor(_1675_v60, _1748_v111, ((_1769_v126)).Merge((_this).f10), new BigNumber((((((_this).f10).contains((_1766_v123).f13)) ? (((_this).f10).get((_1766_v123).f13)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(646)), ((_1771_v62) => function (_1772_i15) {
            return _1771_v62;
          })(_1677_v62))))).length));
          _1770_v127 = _nw266;
          let _1773_v128;
          let _nw267 = new _module.C5();
          _nw267.__ctor((_this.f12)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_this.f12).length))], (_1766_v123).f13);
          _1773_v128 = _nw267;
        }
      } else {
        let _1774___mcc_h6 = (_source20).cf51;
        let _1775_cf51 = _1774___mcc_h6;
        let _1776_v129;
        let _init45 = ((_1777_p1, _1778_p2) => function (_1779_i16) {
          return _module.D0.create_DC0(_1777_p1, _1778_p2);
        })(p1, p2);
        let _nw268 = Array((new BigNumber(16)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw268.length); _i0_45++) {
          _nw268[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1776_v129 = _nw268;
        let _1780_v130;
        _1780_v130 = _module.D7.create_DC21(_1776_v129);
        _1780_v130 = _1780_v130;
        let _1781_v131;
        _1781_v131 = _dafny.Seq.UnicodeFromString("dcbeliug");
        let _1782_v132;
        let _nw269 = new _module.C5();
        _nw269.__ctor((_dafny.ZERO).minus(new BigNumber((_1781_v131).length)), new BigNumber(510));
        _1782_v132 = _nw269;
        let _1783_v133;
        _1783_v133 = _dafny.Map.Empty.slice().updateUnsafe(_1782_v132,_this.f12);
        let _rhs289 = (p2).multipliedBy((_this).f6);
        let _rhs290 = (_this).f6;
        let _rhs291 = (_1783_v133).Merge(_1783_v133);
        let _lhs240 = globalState;
        let _lhs241 = globalState;
        _lhs240.f5 = _rhs289;
        _lhs241.f5 = _rhs290;
        _1783_v133 = _rhs291;
        let _1784_v134;
        _1784_v134 = _dafny.Map.Empty.slice().updateUnsafe(p2,(((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]) ? (!((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))])) : (p0)));
        _1784_v134 = (_1784_v134).update((_1782_v132).f15, (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
        let _arr59 = _this.f11;
        let _index251 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_this.f11).length));
        _arr59[_index251] = (_1782_v132).f15;
        let _1785_v135;
        _1785_v135 = _dafny.Seq.of(_this.f12, _this.f11);
        let _1786_v136;
        let _nw270 = Array((new BigNumber(14)).toNumber()).fill(false);
        _1786_v136 = _nw270;
        let _1787_v137;
        _1787_v137 = _dafny.Seq.of((_1782_v132).f15, (_this).f6, (_dafny.ZERO).minus((_1782_v132).f15));
        let _1788_v138;
        _1788_v138 = _dafny.MultiSet.fromElements(_1787_v137, _1787_v137);
        let _1789_v139;
        _1789_v139 = _dafny.Seq.of(_1787_v137);
        let _arr60 = _this.f11;
        let _index252 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_this.f11).length));
        let _rhs292 = (_dafny.ZERO).minus((_1782_v132).f15);
        let _rhs293 = ((p0) ? (_1781_v131) : (_dafny.Seq.update(_1781_v131, _module.__default.safeIndex((_this).fm14(globalState), new BigNumber((_1781_v131).length)), _1677_v62)));
        let _rhs294 = (((_dafny.MultiSet.FromArray(_1789_v139)).IsSubsetOf(_1788_v138)) ? (_1785_v135) : (_1785_v135));
        let _rhs295 = (_1782_v132).f15;
        let _rhs296 = _1786_v136;
        let _lhs242 = _this.f11;
        let _lhs243 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_this.f11).length));
        let _lhs244 = globalState;
        let _lhs245 = globalState;
        _lhs242[_lhs243] = _rhs292;
        _lhs244.f2 = _rhs293;
        _1785_v135 = _rhs294;
        _lhs245.f5 = _rhs295;
        _1786_v136 = _rhs296;
      }
      r0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(780)), function (_1790_i17) {
        return _dafny.Seq.UnicodeFromString("qvp");
      });
      let _1791_v140;
      _1791_v140 = _dafny.Set.fromElements((_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))]);
      let _1792_v141;
      _1792_v141 = _dafny.MultiSet.fromElements(new BigNumber(976), new BigNumber((_1791_v140).length));
      let _1793_v142;
      _1793_v142 = _module.D11.create_DC31((_this).f6, (_this).f6);
      let _1794_v143;
      _1794_v143 = _dafny.Map.Empty.slice().updateUnsafe(_1792_v141,(p1)[_module.__default.safeIndex((_1793_v142).dtor_cf58, new BigNumber((p1).length))]);
      let _1795_v144;
      _1795_v144 = _module.D5.create_DC16(_1791_v140);
      let _1796_v145;
      _1796_v145 = _module.D11.create_DC33(new BigNumber((_1794_v143).length), p0, (_this).f6, (_this.f9)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_this.f9).length))], (_1795_v144).dtor_cf35);
      r1 = (_1796_v145).dtor_cf63;
      return [r0, r1];
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = [];
      let r3 = _dafny.ZERO;
      (globalState).f5 = (_this).f6;
      let _1797_v0;
      _1797_v0 = false;
      let _1798_v1;
      _1798_v1 = _dafny.Seq.UnicodeFromString("ldvfwsx");
      (globalState).f0 = ((_1797_v0) ? (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("gsjgptuvl"), _1798_v1)) : (_1797_v0));
      (globalState).f0 = _module.__default.fm3(globalState);
      _1797_v0 = !(_1797_v0);
      let _1799_v2;
      _1799_v2 = _module.D7.create_DC22();
      let _1800_v3;
      _1800_v3 = _dafny.Seq.of(_1797_v0);
      let _1801_v4;
      _1801_v4 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1797_v0),_1797_v0);
      let _1802_v5;
      _1802_v5 = _module.D1.create_DC4(_module.D0.create_DC0(_1800_v3, new BigNumber((_1801_v4).length)), _module.__default.fm2((_this).f6, globalState), _module.__default.fm18(_1797_v0, globalState), _1797_v0, (_dafny.ZERO).minus((_this).f6));
      let _rhs297 = (_1802_v5).dtor_cf10;
      let _rhs298 = _1799_v2;
      let _rhs299 = _1797_v0;
      let _rhs300 = _module.__default.fm3(globalState);
      let _lhs246 = globalState;
      let _lhs247 = globalState;
      let _lhs248 = globalState;
      _lhs246.f0 = _rhs297;
      _1799_v2 = _rhs298;
      _lhs247.f0 = _rhs299;
      _lhs248.f0 = _rhs300;
      let _1803_v6;
      _1803_v6 = new _dafny.CodePoint('l'.codePointAt(0));
      let _1804_v7;
      _1804_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(287),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ljwkjfnhn"), _module.__default.safeIndex(new BigNumber(458), new BigNumber((_dafny.Seq.UnicodeFromString("ljwkjfnhn")).length)), _1803_v6)).length)));
      let _1805_v9;
      _1805_v9 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Set.fromElements(_1803_v6));
      let _1806_v10;
      let _nw271 = new _module.C3();
      _nw271.__ctor(_1804_v7, _this.f9, function () {
        let _coll49 = new _dafny.Map();
        for (const _compr_49 of (_1805_v9).Keys.Elements) {
          let _1807_v8 = _compr_49;
          if ((_1805_v9).contains(_1807_v8)) {
            _coll49.push([(_1807_v8).plus(p1),_dafny.Seq.UnicodeFromString("f")]);
          }
        }
        return _coll49;
      }(), p1);
      _1806_v10 = _nw271;
      r0 = _1803_v6;
      let _1808_v11;
      _1808_v11 = _dafny.MultiSet.fromElements(_1797_v0);
      r1 = ((_1797_v0) ? (_module.__default.safeModuloInt((_this).f6, (_this).f6)) : (((_this).f6).plus((_dafny.ZERO).minus((((_1808_v11).contains(false)) ? ((_1808_v11).get(false)) : (new BigNumber(788)))))));
      let _nw272 = Array((new BigNumber(12)).toNumber()).fill(false);
      r2 = _nw272;
      r3 = p1;
      return [r0, r1, r2, r3];
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm10(p0, p1, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(new BigNumber(697), new BigNumber(243), new BigNumber(884), new BigNumber(-12))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(835)));
    };
    fm11(globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((new BigNumber((_dafny.Seq.of(_module.D1.create_DC4(_module.D0.create_DC0(_dafny.Seq.of(!(true), false), new BigNumber((_dafny.Seq.UnicodeFromString("skr")).length)), new BigNumber(887), _dafny.Seq.UnicodeFromString("lqerk"), false, new BigNumber(739)))).length)).minus(new BigNumber(-290)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ctlfngak"), _dafny.Seq.UnicodeFromString("snnrnqplc"))).length));
    };
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.Map.Empty;
      let r2 = _dafny.ZERO;
      let _1809_v0;
      _1809_v0 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
      let _1810_v1;
      let _nw273 = Array((new BigNumber(5)).toNumber());
      _nw273[(_dafny.ZERO).toNumber()] = p1;
      _nw273[(_dafny.ONE).toNumber()] = _module.__default.fm3(globalState);
      _nw273[(new BigNumber(2)).toNumber()] = (p0) && (p3);
      _nw273[(new BigNumber(3)).toNumber()] = p0;
      _nw273[(new BigNumber(4)).toNumber()] = ((!((((_1809_v0).contains(p3)) ? ((_1809_v0).get(p3)) : (_module.__default.fm3(globalState))))) ? ((((_1809_v0).contains(p0)) ? ((_1809_v0).get(p0)) : (p0))) : (p3));
      _1810_v1 = _nw273;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1810_v1).length))) {
        let _1811_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1811_i0)) && ((_1811_i0).isLessThan(new BigNumber((_1810_v1).length))))) {
          (_1810_v1)[(_1811_i0)] = ((p3) ? (true) : ((p0) === (true)));
        }
      }
      let _1812_v2;
      _1812_v2 = _dafny.Seq.of(p2);
      let _1813_v3;
      _1813_v3 = _dafny.MultiSet.fromElements(_1812_v2, _1812_v2);
      if ((_1813_v3).IsDisjointFrom(_1813_v3)) {
        let _1814_v4;
        _1814_v4 = new _dafny.CodePoint('x'.codePointAt(0));
        let _1815_v5;
        _1815_v5 = _module.D0.create_DC1(!(!(p0)), new BigNumber(172), _1814_v4);
        let _1816_v6;
        _1816_v6 = _dafny.Set.fromElements((_this).fm11(globalState), p2, p2);
        let _1817_v8;
        _1817_v8 = _dafny.Seq.UnicodeFromString("mjcroura");
        let _1818_v9;
        _1818_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1814_v4,new BigNumber((_1817_v8).length));
        let _1819_v10;
        _1819_v10 = _dafny.Seq.of(_1818_v9);
        let _rhs301 = (_1816_v6).IsDisjointFrom(function () {
          let _coll50 = new _dafny.Set();
          for (const _compr_50 of _dafny.IntegerRange(new BigNumber(-492), new BigNumber(679))) {
            let _1820_v7 = _compr_50;
            if (((new BigNumber(-492)).isLessThanOrEqualTo(_1820_v7)) && ((_1820_v7).isLessThan(new BigNumber(679)))) {
              _coll50.add((_1820_v7).minus(new BigNumber((_dafny.Seq.of((_1815_v5).dtor_cf2, p1)).length)));
            }
          }
          return _coll50;
        }());
        let _rhs302 = _module.D0.create_DC1(p3, p2, (_1815_v5).dtor_cf4);
        let _rhs303 = ((_1819_v10)[_module.__default.safeIndex(p2, new BigNumber((_1819_v10).length))]).contains(new _dafny.CodePoint('k'.codePointAt(0)));
        let _rhs304 = p0;
        let _lhs249 = globalState;
        let _lhs250 = globalState;
        let _lhs251 = globalState;
        _lhs249.f0 = _rhs301;
        _1815_v5 = _rhs302;
        _lhs250.f0 = _rhs303;
        _lhs251.f0 = _rhs304;
        let _1821_v11;
        let _nw274 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _1821_v11 = _nw274;
        let _index253 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1821_v11).length));
        (_1821_v11)[_index253] = p2;
        let _index254 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1821_v11).length));
        (_1821_v11)[_index254] = (p2).multipliedBy(new BigNumber(240));
        let _pat_let_tv34 = p3;
        _1814_v4 = (function (_pat_let47_0) {
          return function (_1822_dt__update__tmp_h0) {
            return function (_pat_let48_0) {
              return function (_1823_dt__update_hcf2_h0) {
                return _module.D0.create_DC1(_1823_dt__update_hcf2_h0, (_1822_dt__update__tmp_h0).dtor_cf3, (_1822_dt__update__tmp_h0).dtor_cf4);
              }(_pat_let48_0);
            }(_pat_let_tv34);
          }(_pat_let47_0);
        }(_module.D0.create_DC1(p1, (_1821_v11)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_1821_v11).length))], _1814_v4))).dtor_cf4;
        (globalState).f5 = (_1821_v11)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_1821_v11).length))];
        let _index255 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_1821_v11).length));
        (_1821_v11)[_index255] = (_this).fm11(globalState);
      } else {
        let _1824_v12;
        let _init46 = ((_1825_v2) => function (_1826_i1) {
          return _1825_v2;
        })(_1812_v2);
        let _nw275 = Array((new BigNumber(6)).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw275.length); _i0_46++) {
          _nw275[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _1824_v12 = _nw275;
        _1824_v12 = ((!(p0)) ? (_1824_v12) : (_1824_v12));
        let _1827_v13;
        let _nw276 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _1827_v13 = _nw276;
        _1827_v13 = _1827_v13;
        let _1828_v14;
        _1828_v14 = _module.D2.create_DC6(new BigNumber(11), _1827_v13);
        _1828_v14 = _1828_v14;
        _1812_v2 = _1812_v2;
        let _index256 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_1810_v1).length));
        (_1810_v1)[_index256] = p1;
        let _index257 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_1810_v1).length));
        (_1810_v1)[_index257] = p0;
      }
      (globalState).f2 = _dafny.Seq.UnicodeFromString("hrnahy");
      let _1829_v15;
      _1829_v15 = _module.D2.create_DC8(p0, !(false));
      let _pat_let_tv35 = _1809_v0;
      let _pat_let_tv36 = _1809_v0;
      let _pat_let_tv37 = p1;
      let _pat_let_tv38 = _1809_v0;
      _1809_v0 = function (_source21) {
        if (_source21.is_DC6) {
          let _1830___mcc_h0 = (_source21).cf13;
          let _1831___mcc_h1 = (_source21).cf14;
          let _1832_cf14 = _1831___mcc_h1;
          let _1833_cf13 = _1830___mcc_h0;
          return _pat_let_tv35;
        } else if (_source21.is_DC7) {
          let _1834___mcc_h2 = (_source21).cf15;
          let _1835___mcc_h3 = (_source21).cf16;
          let _1836___mcc_h4 = (_source21).cf17;
          let _1837___mcc_h5 = (_source21).cf18;
          let _1838_cf18 = _1837___mcc_h5;
          let _1839_cf17 = _1836___mcc_h4;
          let _1840_cf16 = _1835___mcc_h3;
          let _1841_cf15 = _1834___mcc_h2;
          return _pat_let_tv36;
        } else if (_source21.is_DC8) {
          let _1842___mcc_h6 = (_source21).cf19;
          let _1843___mcc_h7 = (_source21).cf20;
          let _1844_cf20 = _1843___mcc_h7;
          let _1845_cf19 = _1842___mcc_h6;
          return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv37,_1844_cf20);
        } else {
          let _1846___mcc_h8 = (_source21).cf12;
          let _1847_cf12 = _1846___mcc_h8;
          return _pat_let_tv38;
        }
      }(_1829_v15);
      let _1848_v16;
      _1848_v16 = new _dafny.CodePoint('i'.codePointAt(0));
      let _1849_v17;
      _1849_v17 = _dafny.Seq.UnicodeFromString("m");
      let _source22 = _module.__default.fm12(p2, !(_dafny.Seq.contains(_1849_v17, _1848_v16)), !(!(!(p0))), globalState);
      if (_source22.is_DC6) {
        let _1850___mcc_h9 = (_source22).cf13;
        let _1851___mcc_h10 = (_source22).cf14;
        let _1852_cf14 = _1851___mcc_h10;
        let _1853_cf13 = _1850___mcc_h9;
        let _1854_v18;
        _1854_v18 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        let _1855_v19;
        _1855_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1848_v16,false);
        let _1856_v21;
        _1856_v21 = _dafny.MultiSet.fromElements(_1848_v16);
        let _1857_v22;
        _1857_v22 = _dafny.Seq.of(_1855_v19, function () {
          let _coll51 = new _dafny.Map();
          for (const _compr_51 of (_1856_v21).Elements) {
            let _1858_v20 = _compr_51;
            if ((_1856_v21).contains(_1858_v20)) {
              _coll51.push([_1858_v20,!(p1)]);
            }
          }
          return _coll51;
        }());
        _1854_v18 = (_1854_v18).update(p0, new BigNumber((_1857_v22).length));
        _1853_cf13 = (_module.__default.safeDivisionInt(p2, new BigNumber(650))).multipliedBy(new BigNumber(91));
        let _1859_v23;
        let _nw277 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
        _1859_v23 = _nw277;
        let _1860_v24;
        _1860_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1849_v17).length),p1);
        let _1861_v25;
        _1861_v25 = _dafny.Seq.of(_1860_v24);
        let _index258 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_1859_v23).length));
        (_1859_v23)[_index258] = (_1861_v25)[_module.__default.safeIndex(p2, new BigNumber((_1861_v25).length))];
        let _index259 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_1859_v23).length));
        (_1859_v23)[_index259] = _1860_v24;
        let _1862_v26;
        _1862_v26 = _dafny.Set.fromElements(true, p1);
        if ((_dafny.Set.fromElements(p1, !(p3))).IsSubsetOf(_1862_v26)) {
          let _1863_v27;
          _1863_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1853_cf13,_1848_v16);
          _1863_v27 = (_1863_v27).update(_1853_cf13, _1848_v16);
          let _1864_v28;
          let _init47 = ((_1865_v2, _1866_p2) => function (_1867_i2) {
            return _dafny.Seq.Concat(_1865_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-622)), ((_1868_p2) => function (_1869_i3) {
              return _1868_p2;
            })(_1866_p2)));
          })(_1812_v2, p2);
          let _nw278 = Array((new BigNumber(7)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw278.length); _i0_47++) {
            _nw278[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1864_v28 = _nw278;
          let _rhs305 = ((_1853_cf13).multipliedBy(_1853_cf13)).isLessThan(p2);
          let _rhs306 = (_1853_cf13).minus(new BigNumber((_1849_v17).length));
          let _rhs307 = _1864_v28;
          let _lhs252 = globalState;
          _lhs252.f0 = _rhs305;
          r2 = _rhs306;
          _1864_v28 = _rhs307;
          (globalState).f5 = (_1853_cf13).multipliedBy((_dafny.ZERO).minus(_1853_cf13));
          (globalState).f0 = p3;
          let _1870_v29;
          let _nw279 = new _module.C0();
          _nw279.__ctor(_1810_v1, p1);
          _1870_v29 = _nw279;
        } else {
          (globalState).f2 = _1849_v17;
          (globalState).f0 = p1;
          (globalState).f0 = p3;
          (globalState).f0 = ((p3) ? (p1) : (p0));
          let _1871_v30;
          let _nw280 = new _module.C6();
          _nw280.__ctor(((false) ? (p2) : (p2)), p0);
          _1871_v30 = _nw280;
        }
      } else if (_source22.is_DC7) {
        let _1872___mcc_h11 = (_source22).cf15;
        let _1873___mcc_h12 = (_source22).cf16;
        let _1874___mcc_h13 = (_source22).cf17;
        let _1875___mcc_h14 = (_source22).cf18;
        let _1876_cf18 = _1875___mcc_h14;
        let _1877_cf17 = _1874___mcc_h13;
        let _1878_cf16 = _1873___mcc_h12;
        let _1879_cf15 = _1872___mcc_h11;
        if (p0) {
          let _1880_v31;
          _1880_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(537),p2);
          let _1881_v32;
          let _nw281 = new _module.C3();
          _nw281.__ctor(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1849_v17).length),p2), _1810_v1, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(426),_1849_v17), new BigNumber(-162));
          _1881_v32 = _nw281;
          let _1882_v33;
          _1882_v33 = _dafny.Seq.of(_1881_v32, _1881_v32);
          let _1883_v34;
          _1883_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1880_v31).update(new BigNumber((_1882_v33).length), new BigNumber(-92))).length),p3);
          let _1884_v35;
          _1884_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("dppy"),_1876_cf18);
          let _index260 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_1810_v1).length));
          (_1810_v1)[_index260] = (((_1883_v34).contains(new BigNumber((_1884_v35).length))) ? ((_1883_v34).get(new BigNumber((_1884_v35).length))) : (!(p1)));
          let _1885_v36;
          _1885_v36 = _dafny.Seq.of(_module.__default.fm51(globalState));
          let _1886_v37;
          _1886_v37 = _dafny.Set.fromElements(new BigNumber((_1809_v0).length));
          let _index261 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_1810_v1).length));
          let _rhs308 = (_1885_v36)[_module.__default.safeIndex(p2, new BigNumber((_1885_v36).length))];
          let _rhs309 = _1876_cf18;
          let _rhs310 = (_1886_v37).IsSubsetOf((_1886_v37).Difference(_dafny.Set.fromElements(_1877_cf17)));
          let _lhs253 = globalState;
          let _lhs254 = _1810_v1;
          let _lhs255 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_1810_v1).length));
          _1813_v3 = _rhs308;
          _lhs253.f5 = _rhs309;
          _lhs254[_lhs255] = _rhs310;
          let _1887_v38;
          let _nw282 = new _module.C5();
          _nw282.__ctor(new BigNumber((_1879_cf15).length), _1876_cf18);
          _1887_v38 = _nw282;
          let _index262 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_1810_v1).length));
          (_1810_v1)[_index262] = p0;
          let _1888_v39;
          _1888_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((((p0) ? (_1812_v2) : (_dafny.Seq.of(_1876_cf18)))).length),_module.__default.fm0((_1887_v38).f15, _dafny.Set.fromElements(_1877_cf17), (_1810_v1)[_module.__default.safeIndex(new BigNumber(479), new BigNumber((_1810_v1).length))], globalState));
          _1888_v39 = (_1888_v39).update(((p1) ? ((_1887_v38).f15) : ((_1887_v38).f15)), _1848_v16);
          let _1889_v40;
          let _nw283 = new _module.C5();
          _nw283.__ctor(_1876_cf18, (_dafny.ZERO).minus((_this).fm11(globalState)));
          _1889_v40 = _nw283;
        } else {
          let _index263 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length));
          (_1810_v1)[_index263] = p1;
          let _index264 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_1810_v1).length));
          (_1810_v1)[_index264] = p0;
          let _index265 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length));
          let _index266 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_1810_v1).length));
          let _rhs311 = (p0) && (p1);
          let _rhs312 = p1;
          let _rhs313 = !(p1);
          let _rhs314 = p1;
          let _lhs256 = globalState;
          let _lhs257 = _1810_v1;
          let _lhs258 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length));
          let _lhs259 = globalState;
          let _lhs260 = _1810_v1;
          let _lhs261 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_1810_v1).length));
          _lhs256.f0 = _rhs311;
          _lhs257[_lhs258] = _rhs312;
          _lhs259.f0 = _rhs313;
          _lhs260[_lhs261] = _rhs314;
          let _1890_v41;
          _1890_v41 = _dafny.MultiSet.fromElements(p3, (_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))], p3);
          let _1891_v42;
          let _nw284 = new _module.C5();
          _nw284.__ctor((_this).fm11(globalState), (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_1890_v41).cardinality())).plus(new BigNumber(-241)))));
          _1891_v42 = _nw284;
          let _1892_v43;
          let _nw285 = Array((new BigNumber(28)).toNumber());
          _nw285[(_dafny.ZERO).toNumber()] = p0;
          _nw285[(_dafny.ONE).toNumber()] = !(!(p0));
          _nw285[(new BigNumber(2)).toNumber()] = p1;
          _nw285[(new BigNumber(3)).toNumber()] = (_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))];
          _nw285[(new BigNumber(4)).toNumber()] = _module.__default.fm3(globalState);
          _nw285[(new BigNumber(5)).toNumber()] = (_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))];
          _nw285[(new BigNumber(6)).toNumber()] = (new BigNumber(-985)).isLessThanOrEqualTo(_1877_cf17);
          _nw285[(new BigNumber(7)).toNumber()] = p3;
          _nw285[(new BigNumber(8)).toNumber()] = !(true);
          _nw285[(new BigNumber(9)).toNumber()] = p0;
          _nw285[(new BigNumber(10)).toNumber()] = p0;
          _nw285[(new BigNumber(11)).toNumber()] = p3;
          _nw285[(new BigNumber(12)).toNumber()] = true;
          _nw285[(new BigNumber(13)).toNumber()] = p3;
          _nw285[(new BigNumber(14)).toNumber()] = p0;
          _nw285[(new BigNumber(15)).toNumber()] = !(p0);
          _nw285[(new BigNumber(16)).toNumber()] = true;
          _nw285[(new BigNumber(17)).toNumber()] = !(p3) || (p3);
          _nw285[(new BigNumber(18)).toNumber()] = (_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))];
          _nw285[(new BigNumber(19)).toNumber()] = (_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))];
          _nw285[(new BigNumber(20)).toNumber()] = (_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))];
          _nw285[(new BigNumber(21)).toNumber()] = p3;
          _nw285[(new BigNumber(22)).toNumber()] = p3;
          _nw285[(new BigNumber(23)).toNumber()] = !(p0) || (p1);
          _nw285[(new BigNumber(24)).toNumber()] = (_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))];
          _nw285[(new BigNumber(25)).toNumber()] = true;
          _nw285[(new BigNumber(26)).toNumber()] = !_dafny.areEqual(_1879_cf15, _1879_cf15);
          _nw285[(new BigNumber(27)).toNumber()] = _module.__default.fm3(globalState);
          _1892_v43 = _nw285;
          let _1893_v44;
          let _init48 = ((_1894_p2) => function (_1895_i4) {
            return (_1895_i4).multipliedBy(_1894_p2);
          })(p2);
          let _nw286 = Array((new BigNumber(27)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw286.length); _i0_48++) {
            _nw286[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1893_v44 = _nw286;
          let _index267 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length));
          let _rhs315 = !(p3);
          let _rhs316 = (_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))];
          let _rhs317 = _1892_v43;
          let _rhs318 = _1893_v44;
          let _lhs262 = _1810_v1;
          let _lhs263 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length));
          let _lhs264 = globalState;
          _lhs262[_lhs263] = _rhs315;
          _lhs264.f0 = _rhs316;
          _1892_v43 = _rhs317;
          _1893_v44 = _rhs318;
          let _1896_v45;
          let _init49 = ((_1897_v2) => function (_1898_i5) {
            return _1897_v2;
          })(_1812_v2);
          let _nw287 = Array((new BigNumber(23)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw287.length); _i0_49++) {
            _nw287[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1896_v45 = _nw287;
          let _index268 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1896_v45).length));
          (_1896_v45)[_index268] = _1812_v2;
          let _1899_v46;
          _1899_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1812_v2).length),_1877_cf17);
          let _1900_v47;
          _1900_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1899_v46,_dafny.Seq.UnicodeFromString("dgimpf"));
          let _1901_v48;
          _1901_v48 = _module.D14.create_DC37(_module.__default.fm3(globalState), (((_1900_v47).contains(_1899_v46)) ? ((_1900_v47).get(_1899_v46)) : (_1849_v17)));
          let _1902_v49;
          _1902_v49 = _module.D5.create_DC17(_1812_v2, p3);
          let _pat_let_tv39 = p0;
          let _pat_let_tv40 = _1810_v1;
          let _pat_let_tv41 = _1810_v1;
          let _pat_let_tv42 = _1849_v17;
          let _1903_v50;
          let _nw288 = Array((new BigNumber(29)).toNumber());
          _nw288[(_dafny.ZERO).toNumber()] = _1901_v48;
          _nw288[(_dafny.ONE).toNumber()] = _module.D14.create_DC37((((_1809_v0).contains(p1)) ? ((_1809_v0).get(p1)) : ((_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))])), _1849_v17);
          _nw288[(new BigNumber(2)).toNumber()] = _module.D14.create_DC37(_module.__default.fm3(globalState), _dafny.Seq.UnicodeFromString("im"));
          _nw288[(new BigNumber(3)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(4)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(5)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(6)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(7)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(8)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(9)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(10)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(11)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(12)).toNumber()] = ((true) ? (_module.D14.create_DC37(p3, _1849_v17)) : (_1901_v48));
          _nw288[(new BigNumber(13)).toNumber()] = function (_pat_let49_0) {
            return function (_1904_dt__update__tmp_h1) {
              return function (_pat_let50_0) {
                return function (_1905_dt__update_hcf70_h0) {
                  return _module.D14.create_DC37(_1905_dt__update_hcf70_h0, (_1904_dt__update__tmp_h1).dtor_cf71);
                }(_pat_let50_0);
              }(_pat_let_tv39);
            }(_pat_let49_0);
          }(_1901_v48);
          _nw288[(new BigNumber(14)).toNumber()] = _module.D14.create_DC37((_1810_v1)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_1810_v1).length))], _1849_v17);
          _nw288[(new BigNumber(15)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(16)).toNumber()] = _module.D14.create_DC37(p3, _1849_v17);
          _nw288[(new BigNumber(17)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(18)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(19)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(20)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(21)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(22)).toNumber()] = function (_pat_let51_0) {
            return function (_1906_dt__update__tmp_h2) {
              return function (_pat_let52_0) {
                return function (_1907_dt__update_hcf70_h1) {
                  return _module.D14.create_DC37(_1907_dt__update_hcf70_h1, (_1906_dt__update__tmp_h2).dtor_cf71);
                }(_pat_let52_0);
              }((_pat_let_tv41)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_pat_let_tv40).length))]);
            }(_pat_let51_0);
          }(_1901_v48);
          _nw288[(new BigNumber(23)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(24)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(25)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(26)).toNumber()] = (((_1902_v49).dtor_cf37) ? (_1901_v48) : (_1901_v48));
          _nw288[(new BigNumber(27)).toNumber()] = _1901_v48;
          _nw288[(new BigNumber(28)).toNumber()] = function (_pat_let53_0) {
            return function (_1908_dt__update__tmp_h3) {
              return function (_pat_let54_0) {
                return function (_1909_dt__update_hcf71_h0) {
                  return _module.D14.create_DC37((_1908_dt__update__tmp_h3).dtor_cf70, _1909_dt__update_hcf71_h0);
                }(_pat_let54_0);
              }(_pat_let_tv42);
            }(_pat_let53_0);
          }(_module.__default.fm52(p1, globalState));
          _1903_v50 = _nw288;
          let _index269 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_1903_v50).length));
          (_1903_v50)[_index269] = _module.D14.create_DC37(p0, _1849_v17);
          let _index270 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1896_v45).length));
          let _index271 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_1903_v50).length));
          let _rhs319 = _1848_v16;
          let _rhs320 = _dafny.Seq.update(_1812_v2, _module.__default.safeIndex(_module.__default.safeModuloInt((_this).fm11(globalState), _1876_cf18), new BigNumber((_1812_v2).length)), _1876_cf18);
          let _rhs321 = _1878_cf16;
          let _rhs322 = p0;
          let _rhs323 = _1901_v48;
          let _lhs265 = _1896_v45;
          let _lhs266 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1896_v45).length));
          let _lhs267 = globalState;
          let _lhs268 = _1903_v50;
          let _lhs269 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_1903_v50).length));
          _1848_v16 = _rhs319;
          _lhs265[_lhs266] = _rhs320;
          _1878_cf16 = _rhs321;
          _lhs267.f0 = _rhs322;
          _lhs268[_lhs269] = _rhs323;
          let _1910_v51;
          _1910_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1877_cf17,p1);
          _1910_v51 = (_1910_v51).update(p2, p1);
        }
        let _source23 = _1829_v15;
        if (_source23.is_DC6) {
          let _1911___mcc_h18 = (_source23).cf13;
          let _1912___mcc_h19 = (_source23).cf14;
          let _1913_cf14 = _1912___mcc_h19;
          let _1914_cf13 = _1911___mcc_h18;
          let _1915_v52;
          _1915_v52 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), _1848_v16);
          let _1916_v53;
          _1916_v53 = _module.D1.create_DC3(_1915_v52);
          let _1917_v54;
          _1917_v54 = _dafny.MultiSet.fromElements(_1916_v53, _1916_v53, _1916_v53, _1916_v53, _1916_v53);
          let _rhs324 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(61)), ((_1918_v16) => function (_1919_i6) {
            return _1918_v16;
          })(_1848_v16));
          let _rhs325 = (_module.__default.fm53(_1848_v16, globalState)).IsDisjointFrom(_1917_v54);
          let _lhs270 = globalState;
          let _lhs271 = globalState;
          _lhs270.f2 = _rhs324;
          _lhs271.f0 = _rhs325;
          (globalState).f5 = _1877_cf17;
          (globalState).f0 = p0;
          let _1920_v55;
          let _nw289 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1920_v55 = _nw289;
          let _1921_v56;
          _1921_v56 = _module.D10.create_DC29(new BigNumber(534), _1920_v55, true, p2);
          let _1922_v57;
          _1922_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(525));
          let _pat_let_tv43 = _1920_v55;
          let _pat_let_tv44 = _1914_cf13;
          let _pat_let_tv45 = p3;
          let _1923_v58;
          let _nw290 = Array((new BigNumber(23)).toNumber());
          _nw290[(_dafny.ZERO).toNumber()] = _1921_v56;
          _nw290[(_dafny.ONE).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(2)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(3)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(4)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(5)).toNumber()] = ((p3) ? (function (_pat_let55_0) {
            return function (_1924_dt__update__tmp_h4) {
              return function (_pat_let56_0) {
                return function (_1925_dt__update_hcf54_h0) {
                  return function (_pat_let57_0) {
                    return function (_1926_dt__update_hcf56_h0) {
                      return function (_pat_let58_0) {
                        return function (_1927_dt__update_hcf55_h0) {
                          return _module.D10.create_DC29((_1924_dt__update__tmp_h4).dtor_cf53, _1925_dt__update_hcf54_h0, _1927_dt__update_hcf55_h0, _1926_dt__update_hcf56_h0);
                        }(_pat_let58_0);
                      }(_pat_let_tv45);
                    }(_pat_let57_0);
                  }(_pat_let_tv44);
                }(_pat_let56_0);
              }(_pat_let_tv43);
            }(_pat_let55_0);
          }(_module.D10.create_DC29(new BigNumber((_1849_v17).length), _1920_v55, p1, new BigNumber(-766)))) : (_1921_v56));
          _nw290[(new BigNumber(6)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(7)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(8)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(9)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(10)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(11)).toNumber()] = _module.D10.create_DC29(_1876_cf18, _1920_v55, p0, _1877_cf17);
          _nw290[(new BigNumber(12)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(13)).toNumber()] = _module.D10.create_DC29(_1877_cf17, _1920_v55, p3, _1914_cf13);
          _nw290[(new BigNumber(14)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(15)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(16)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(17)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(18)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(19)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(20)).toNumber()] = _module.D10.create_DC29(new BigNumber(((_1922_v57).update(p0, _1876_cf18)).length), _1920_v55, false, p2);
          _nw290[(new BigNumber(21)).toNumber()] = _1921_v56;
          _nw290[(new BigNumber(22)).toNumber()] = _1921_v56;
          _1923_v58 = _nw290;
          _1923_v58 = _1923_v58;
        } else if (_source23.is_DC7) {
          let _1928___mcc_h20 = (_source23).cf15;
          let _1929___mcc_h21 = (_source23).cf16;
          let _1930___mcc_h22 = (_source23).cf17;
          let _1931___mcc_h23 = (_source23).cf18;
          let _1932_cf18 = _1931___mcc_h23;
          let _1933_cf17 = _1930___mcc_h22;
          let _1934_cf16 = _1929___mcc_h21;
          let _1935_cf15 = _1928___mcc_h20;
          let _1936_v59;
          _1936_v59 = _dafny.MultiSet.fromElements(new BigNumber((_1812_v2).length), p2, p2, _1933_cf17, new BigNumber(534));
          let _1937_v60;
          _1937_v60 = _dafny.Seq.of(_1812_v2, _1812_v2);
          let _1938_v61;
          _1938_v61 = _dafny.Set.fromElements(new BigNumber(((_1937_v60)[_module.__default.safeIndex(_1933_cf17, new BigNumber((_1937_v60).length))]).length));
          _1848_v16 = _module.__default.fm0(new BigNumber((_1936_v59).cardinality()), _1938_v61, p0, globalState);
          let _1939_v62;
          _1939_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1934_cf16,_1933_cf17);
          _1939_v62 = (_1939_v62).update(_1934_cf16, _1932_cf18);
          let _rhs326 = !(p3) || ((_dafny.Map.Empty.slice().updateUnsafe(p3,_1932_cf18)).contains(p1));
          let _rhs327 = _1932_cf18;
          let _lhs272 = globalState;
          _lhs272.f0 = _rhs326;
          r2 = _rhs327;
          let _1940_v63;
          _1940_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1933_cf17,_module.__default.fm0(_1933_cf17, _1938_v61, p0, globalState));
          let _1941_v64;
          _1941_v64 = _dafny.MultiSet.fromElements(_1940_v63, _1940_v63);
          let _index272 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_1810_v1).length));
          (_1810_v1)[_index272] = (_1941_v64).IsDisjointFrom(_1941_v64);
          let _index273 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_1810_v1).length));
          (_1810_v1)[_index273] = p0;
        } else if (_source23.is_DC8) {
          let _1942___mcc_h24 = (_source23).cf19;
          let _1943___mcc_h25 = (_source23).cf20;
          let _1944_cf20 = _1943___mcc_h25;
          let _1945_cf19 = _1942___mcc_h24;
          let _1946_v65;
          let _nw291 = new _module.C1();
          _nw291.__ctor(_1877_cf17);
          _1946_v65 = _nw291;
          let _1947_v66;
          _1947_v66 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _1948_v67;
          _1948_v67 = _dafny.Set.fromElements(_1944_cf20);
          let _1949_v68;
          let _nw292 = Array((new BigNumber(9)).toNumber());
          _nw292[(_dafny.ZERO).toNumber()] = p2;
          _nw292[(_dafny.ONE).toNumber()] = new BigNumber((_1947_v66).length);
          _nw292[(new BigNumber(2)).toNumber()] = _1877_cf17;
          _nw292[(new BigNumber(3)).toNumber()] = p2;
          _nw292[(new BigNumber(4)).toNumber()] = new BigNumber((_1948_v67).length);
          _nw292[(new BigNumber(5)).toNumber()] = _1876_cf18;
          _nw292[(new BigNumber(6)).toNumber()] = p2;
          _nw292[(new BigNumber(7)).toNumber()] = _1876_cf18;
          _nw292[(new BigNumber(8)).toNumber()] = new BigNumber((_1849_v17).length);
          _1949_v68 = _nw292;
          let _1950_v69;
          _1950_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1949_v68,_1944_cf20);
          let _1951_v70;
          _1951_v70 = _dafny.Set.fromElements(new BigNumber((_1950_v69).length));
          let _1952_v71;
          _1952_v71 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1951_v70).length),p1);
          let _1953_v72;
          _1953_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1877_cf17,_module.__default.fm4(new BigNumber((_1947_v66).length), p2, _1877_cf17, globalState));
          let _1954_v73;
          let _nw293 = Array((new BigNumber(11)).toNumber());
          _nw293[(_dafny.ZERO).toNumber()] = p2;
          _nw293[(_dafny.ONE).toNumber()] = _1876_cf18;
          _nw293[(new BigNumber(2)).toNumber()] = _1877_cf17;
          _nw293[(new BigNumber(3)).toNumber()] = _1876_cf18;
          _nw293[(new BigNumber(4)).toNumber()] = _1877_cf17;
          _nw293[(new BigNumber(5)).toNumber()] = (new BigNumber((_1953_v72).length)).plus(p2);
          _nw293[(new BigNumber(6)).toNumber()] = _1877_cf17;
          _nw293[(new BigNumber(7)).toNumber()] = _1877_cf17;
          _nw293[(new BigNumber(8)).toNumber()] = new BigNumber(902);
          _nw293[(new BigNumber(9)).toNumber()] = _1877_cf17;
          _nw293[(new BigNumber(10)).toNumber()] = new BigNumber(905);
          _1954_v73 = _nw293;
          let _index274 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1954_v73).length));
          (_1954_v73)[_index274] = _1876_cf18;
          let _1955_v74;
          _1955_v74 = _dafny.MultiSet.fromElements(p0);
          let _1956_v75;
          _1956_v75 = _1955_v74;
          let _index275 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1954_v73).length));
          let _rhs328 = (_dafny.ZERO).minus((_1876_cf18).minus(_1877_cf17));
          let _rhs329 = p2;
          let _rhs330 = p2;
          let _rhs331 = ((_1955_v74).Intersect((_1956_v75))).Intersect((_1955_v74).Difference(_1955_v74));
          let _lhs273 = _1954_v73;
          let _lhs274 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1954_v73).length));
          let _lhs275 = globalState;
          _1876_cf18 = _rhs328;
          _lhs273[_lhs274] = _rhs329;
          _lhs275.f5 = _rhs330;
          _1955_v74 = _rhs331;
          let _1957_v76;
          _1957_v76 = _dafny.Set.fromElements(_1951_v70, _1951_v70);
          let _1958_v77;
          _1958_v77 = _dafny.Map.Empty.slice().updateUnsafe(((_1945_cf19) ? (_1877_cf17) : (_1876_cf18)),(_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1957_v76).length), (_dafny.ZERO).minus((_1954_v73)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1954_v73).length))]))));
          let _1959_v78;
          _1959_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1810_v1,p1);
          let _1960_v79;
          _1960_v79 = _module.D14.create_DC37(_1944_cf20, _1849_v17);
          let _1961_v80;
          _1961_v80 = _module.D0.create_DC1(true, _1876_cf18, _module.__default.fm0((_1954_v73)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1954_v73).length))], _1951_v70, _1945_cf19, globalState));
          let _pat_let_tv46 = p3;
          let _pat_let_tv47 = _1848_v16;
          let _rhs332 = (_1958_v77).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1879_cf15).length),(function (_pat_let59_0) {
            return function (_1962_dt__update__tmp_h5) {
              return function (_pat_let60_0) {
                return function (_1963_dt__update_hcf2_h1) {
                  return function (_pat_let61_0) {
                    return function (_1964_dt__update_hcf4_h0) {
                      return _module.D0.create_DC1(_1963_dt__update_hcf2_h1, (_1962_dt__update__tmp_h5).dtor_cf3, _1964_dt__update_hcf4_h0);
                    }(_pat_let61_0);
                  }(_pat_let_tv47);
                }(_pat_let60_0);
              }(_pat_let_tv46);
            }(_pat_let59_0);
          }(_1961_v80)).dtor_cf3));
          let _rhs333 = ((_1944_cf20) ? (_1959_v78) : (_1959_v78));
          let _rhs334 = _1960_v79;
          _1958_v77 = _rhs332;
          _1959_v78 = _rhs333;
          _1960_v79 = _rhs334;
          let _1965_v82;
          let _nw294 = new _module.C3();
          _nw294.__ctor(_1958_v77, _1810_v1, function () {
            let _coll52 = new _dafny.Map();
            for (const _compr_52 of (_1947_v66).Keys.Elements) {
              let _1966_v81 = _compr_52;
              if ((_1947_v66).contains(_1966_v81)) {
                _coll52.push([(_1966_v81).multipliedBy(_1876_cf18),_dafny.Seq.Create(_module.__default.abs(new BigNumber(517)), ((_1967_v16) => function (_1968_i7) {
                  return _1967_v16;
                })(_1848_v16))]);
              }
            }
            return _coll52;
          }(), _1876_cf18);
          _1965_v82 = _nw294;
          let _1969_v83;
          _1969_v83 = _dafny.Seq.of(_1965_v82);
          let _1970_v84;
          _1970_v84 = _module.D17.create_DC40(_1965_v82);
          let _1971_v85;
          let _nw295 = Array((new BigNumber(20)).toNumber());
          _nw295[(_dafny.ZERO).toNumber()] = _1965_v82;
          _nw295[(_dafny.ONE).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(2)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(3)).toNumber()] = ((p0) ? (_1965_v82) : (_1965_v82));
          _nw295[(new BigNumber(4)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(5)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(6)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(7)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(8)).toNumber()] = (_1969_v83)[_module.__default.safeIndex(_1876_cf18, new BigNumber((_1969_v83).length))];
          _nw295[(new BigNumber(9)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(10)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(11)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(12)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(13)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(14)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(15)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(16)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(17)).toNumber()] = (_1970_v84).dtor_cf74;
          _nw295[(new BigNumber(18)).toNumber()] = _1965_v82;
          _nw295[(new BigNumber(19)).toNumber()] = _1965_v82;
          _1971_v85 = _nw295;
          let _index276 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1971_v85).length));
          (_1971_v85)[_index276] = _1965_v82;
          let _1972_v86;
          _1972_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1877_cf17,_1965_v82);
          let _index277 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1971_v85).length));
          (_1971_v85)[_index277] = (((_1972_v86).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(((_1954_v73)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1954_v73).length))]).plus((_1954_v73)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1954_v73).length))]))))) ? ((_1972_v86).get((_dafny.ZERO).minus((_dafny.ZERO).minus(((_1954_v73)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1954_v73).length))]).plus((_1954_v73)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1954_v73).length))]))))) : (_1965_v82));
        } else {
          let _1973___mcc_h26 = (_source23).cf12;
          let _1974_cf12 = _1973___mcc_h26;
          let _1975_v87;
          let _nw296 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _1975_v87 = _nw296;
          let _1976_v88;
          _1976_v88 = _dafny.MultiSet.fromElements(_1975_v87);
          let _1977_v89;
          _1977_v89 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(403)), ((_1978_cf18) => function (_1979_i8) {
            return _1978_cf18;
          })(_1876_cf18)));
          let _1980_v90;
          _1980_v90 = _dafny.MultiSet.fromElements(p1);
          let _1981_v91;
          _1981_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(284),_dafny.Seq.Concat(_dafny.Seq.update(_1849_v17, _module.__default.safeIndex(new BigNumber(850), new BigNumber((_1849_v17).length)), _1848_v16), _dafny.Seq.Create(_module.__default.abs(new BigNumber(396)), ((_1982_v16) => function (_1983_i9) {
            return _1982_v16;
          })(_1848_v16))));
          let _rhs335 = (_1976_v88).IsDisjointFrom(_1976_v88);
          let _rhs336 = !(_1877_cf17).isEqualTo(_1877_cf17);
          let _rhs337 = (_1876_cf18).isLessThanOrEqualTo(_module.__default.safeDivisionInt(new BigNumber(((((_1977_v89).contains(p0)) ? ((_1977_v89).get(p0)) : (_1812_v2))).length), new BigNumber((_1980_v90).cardinality())));
          let _rhs338 = (((_1981_v91).contains(_1876_cf18)) ? ((_1981_v91).get(_1876_cf18)) : (_1849_v17));
          let _lhs276 = globalState;
          let _lhs277 = globalState;
          let _lhs278 = globalState;
          _lhs276.f0 = _rhs335;
          _lhs277.f0 = _rhs336;
          _lhs278.f0 = _rhs337;
          _1849_v17 = _rhs338;
          _1810_v1 = _1810_v1;
          _1848_v16 = _1848_v16;
          let _index278 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_1975_v87).length));
          (_1975_v87)[_index278] = _1877_cf17;
          let _1984_v92;
          _1984_v92 = _dafny.Seq.of(_1975_v87);
          let _1985_v93;
          _1985_v93 = _module.D4.create_DC15(p1, p0, _dafny.Seq.update(_1849_v17, _module.__default.safeIndex(_1877_cf17, new BigNumber((_1849_v17).length)), _1848_v16));
          let _1986_v94;
          _1986_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1985_v93,_1975_v87);
          let _1987_v95;
          _1987_v95 = _dafny.Map.Empty.slice().updateUnsafe(p2,(((_1986_v94).contains(_1985_v93)) ? ((_1986_v94).get(_1985_v93)) : (_1975_v87)));
          let _1988_v96;
          _1988_v96 = _module.D2.create_DC6(p2, _1975_v87);
          let _1989_v97;
          let _nw297 = Array((new BigNumber(28)).toNumber());
          _nw297[(_dafny.ZERO).toNumber()] = _1975_v87;
          _nw297[(_dafny.ONE).toNumber()] = (_1984_v92)[_module.__default.safeIndex(new BigNumber((_1987_v95).length), new BigNumber((_1984_v92).length))];
          _nw297[(new BigNumber(2)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(3)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(4)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(5)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(6)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(7)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(8)).toNumber()] = ((p1) ? (_1975_v87) : (_1975_v87));
          _nw297[(new BigNumber(9)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(10)).toNumber()] = ((p0) ? (_1975_v87) : ((_1984_v92)[_module.__default.safeIndex(_1877_cf17, new BigNumber((_1984_v92).length))]));
          _nw297[(new BigNumber(11)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(12)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(13)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(14)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(15)).toNumber()] = (_1988_v96).dtor_cf14;
          _nw297[(new BigNumber(16)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(17)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(18)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(19)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(20)).toNumber()] = ((p3) ? (_1975_v87) : (_1975_v87));
          _nw297[(new BigNumber(21)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(22)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(23)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(24)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(25)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(26)).toNumber()] = _1975_v87;
          _nw297[(new BigNumber(27)).toNumber()] = _1975_v87;
          _1989_v97 = _nw297;
          let _index279 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1989_v97).length));
          (_1989_v97)[_index279] = _1975_v87;
          let _1990_v98;
          _1990_v98 = _dafny.Set.fromElements((_1849_v17)[_module.__default.safeIndex(_1876_cf18, new BigNumber((_1849_v17).length))], _1848_v16, _1848_v16);
          let _1991_v99;
          _1991_v99 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_1876_cf18, globalState),_1812_v2);
          let _index280 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_1975_v87).length));
          let _index281 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1989_v97).length));
          let _rhs339 = _1877_cf17;
          let _rhs340 = _1975_v87;
          let _rhs341 = (_1990_v98).Intersect(_dafny.Set.fromElements(_1848_v16, new _dafny.CodePoint('y'.codePointAt(0))));
          let _rhs342 = (_dafny.ZERO).minus(new BigNumber(((((_1991_v99).contains(new BigNumber((_1879_cf15).length))) ? ((_1991_v99).get(new BigNumber((_1879_cf15).length))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(836)), ((_1992_p2) => function (_1993_i10) {
            return _1992_p2;
          })(p2))))).length));
          let _lhs279 = _1975_v87;
          let _lhs280 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_1975_v87).length));
          let _lhs281 = _1989_v97;
          let _lhs282 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1989_v97).length));
          let _lhs283 = globalState;
          _lhs279[_lhs280] = _rhs339;
          _lhs281[_lhs282] = _rhs340;
          _1990_v98 = _rhs341;
          _lhs283.f5 = _rhs342;
        }
        let _1994_v100;
        let _nw298 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.of());
        _1994_v100 = _nw298;
        let _index282 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_1994_v100).length));
        (_1994_v100)[_index282] = _1812_v2;
        let _index283 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_1994_v100).length));
        (_1994_v100)[_index283] = _1812_v2;
        (globalState).f5 = (_1877_cf17).plus(_1877_cf17);
      } else if (_source22.is_DC8) {
        let _1995___mcc_h15 = (_source22).cf19;
        let _1996___mcc_h16 = (_source22).cf20;
        let _1997_cf20 = _1996___mcc_h16;
        let _1998_cf19 = _1995___mcc_h15;
        let _1999_v101;
        let _nw299 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _1999_v101 = _nw299;
        let _2000_v102;
        _2000_v102 = _dafny.Map.Empty.slice().updateUnsafe(_1999_v101,_1997_cf20);
        let _2001_v103;
        _2001_v103 = _module.D10.create_DC28(_2000_v102);
        let _source24 = _2001_v103;
        if (_source24.is_DC29) {
          let _2002___mcc_h27 = (_source24).cf53;
          let _2003___mcc_h28 = (_source24).cf54;
          let _2004___mcc_h29 = (_source24).cf55;
          let _2005___mcc_h30 = (_source24).cf56;
          let _2006_cf56 = _2005___mcc_h30;
          let _2007_cf55 = _2004___mcc_h29;
          let _2008_cf54 = _2003___mcc_h28;
          let _2009_cf53 = _2002___mcc_h27;
          let _2010_v104;
          _2010_v104 = _dafny.Map.Empty.slice().updateUnsafe(_2009_cf53,_dafny.Seq.Create(_module.__default.abs(new BigNumber(958)), ((_2011_v16) => function (_2012_i11) {
            return _2011_v16;
          })(_1848_v16)));
          let _2013_v105;
          let _nw300 = new _module.C7();
          _nw300.__ctor(_1999_v101, _1999_v101, _2009_cf53, _1810_v1, _2010_v104);
          _2013_v105 = _nw300;
          let _2014_v106;
          _2014_v106 = _dafny.Set.fromElements(p3);
          let _2015_v107;
          _2015_v107 = _module.D11.create_DC33(_2009_cf53, _2007_cf55, new BigNumber(655), false, _2014_v106);
          let _2016_v108;
          _2016_v108 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_2015_v107).dtor_cf65, _1998_cf19, _1998_cf19, p1, _1998_cf19),_2010_v104);
          let _2017_v109;
          _2017_v109 = _dafny.Seq.of(p0, _2007_cf55);
          let _nw301 = new _module.C7();
          _nw301.__ctor(_2013_v105.f11, _1999_v101, _2009_cf53, _1810_v1, (((_2016_v108).contains(_2017_v109)) ? ((_2016_v108).get(_2017_v109)) : (_2010_v104)));
          _2013_v105 = _nw301;
          _1849_v17 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1849_v17, _1849_v17), _1849_v17);
          let _index284 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_2008_cf54).length));
          (_2008_cf54)[_index284] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("btcbtiuf"), _1849_v17);
          let _index285 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_2008_cf54).length));
          let _rhs343 = p2;
          let _rhs344 = _dafny.Seq.Concat(_1849_v17, _1849_v17);
          let _lhs284 = _2008_cf54;
          let _lhs285 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_2008_cf54).length));
          _2009_cf53 = _rhs343;
          _lhs284[_lhs285] = _rhs344;
          (globalState).f5 = (p2).plus(_2009_cf53);
        } else {
          let _2018___mcc_h31 = (_source24).cf52;
          let _2019_cf52 = _2018___mcc_h31;
          let _2020_v110;
          _2020_v110 = _dafny.MultiSet.fromElements(_module.__default.fm0(new BigNumber(192), _dafny.Set.fromElements(new BigNumber(-474)), _1998_cf19, globalState), _1848_v16);
          let _2021_v111;
          _2021_v111 = _module.D1.create_DC3(_2020_v110);
          let _pat_let_tv48 = _2020_v110;
          let _2022_v112;
          _2022_v112 = _dafny.Map.Empty.slice().updateUnsafe(p2,function (_pat_let62_0) {
            return function (_2023_dt__update__tmp_h6) {
              return function (_pat_let63_0) {
                return function (_2024_dt__update_hcf6_h0) {
                  return _module.D1.create_DC3(_2024_dt__update_hcf6_h0);
                }(_pat_let63_0);
              }(_pat_let_tv48);
            }(_pat_let62_0);
          }(_2021_v111));
          _2022_v112 = (_2022_v112).update(p2, _2021_v111);
          let _2025_v113;
          _2025_v113 = _dafny.MultiSet.fromElements(p2, new BigNumber((_1849_v17).length));
          let _2026_v114;
          _2026_v114 = _dafny.Seq.of(_2025_v113, _2025_v113, _2025_v113);
          let _2027_v115;
          _2027_v115 = _module.D11.create_DC31(p2, p2);
          let _pat_let_tv49 = p2;
          let _pat_let_tv50 = p2;
          let _pat_let_tv51 = _1849_v17;
          let _2028_v116;
          _2028_v116 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2026_v114).length),(function (_pat_let64_0) {
            return function (_2033_dt__update__tmp_h9) {
              return function (_pat_let69_0) {
                return function (_2034_dt__update_hcf58_h0) {
                  return _module.D11.create_DC31(_2034_dt__update_hcf58_h0, (_2033_dt__update__tmp_h9).dtor_cf59);
                }(_pat_let69_0);
              }(new BigNumber((_pat_let_tv51).length));
            }(_pat_let64_0);
          }(function (_pat_let65_0) {
            return function (_2031_dt__update__tmp_h8) {
              return function (_pat_let68_0) {
                return function (_2032_dt__update_hcf59_h1) {
                  return _module.D11.create_DC31((_2031_dt__update__tmp_h8).dtor_cf58, _2032_dt__update_hcf59_h1);
                }(_pat_let68_0);
              }(_pat_let_tv50);
            }(_pat_let65_0);
          }(function (_pat_let66_0) {
            return function (_2029_dt__update__tmp_h7) {
              return function (_pat_let67_0) {
                return function (_2030_dt__update_hcf59_h0) {
                  return _module.D11.create_DC31((_2029_dt__update__tmp_h7).dtor_cf58, _2030_dt__update_hcf59_h0);
                }(_pat_let67_0);
              }(_pat_let_tv49);
            }(_pat_let66_0);
          }(_2027_v115)))).dtor_cf59);
          _2028_v116 = ((_dafny.Map.Empty.slice().updateUnsafe(p2,p2)).Merge(_2028_v116)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,p2));
          _1998_cf19 = _1997_cf20;
          (globalState).f5 = (_this).fm11(globalState);
        }
        let _2035_v117;
        let _nw302 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2035_v117 = _nw302;
        let _index286 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_2035_v117).length));
        (_2035_v117)[_index286] = _1849_v17;
        let _2036_v118;
        _2036_v118 = _dafny.Seq.of(_1812_v2);
        let _2037_v119;
        _2037_v119 = _dafny.Seq.of(_dafny.Seq.Concat(_1849_v17, _1849_v17));
        let _2038_v120;
        _2038_v120 = _dafny.Set.fromElements(_1999_v101, _1999_v101);
        let _2039_v121;
        _2039_v121 = _dafny.MultiSet.fromElements(new BigNumber((_2038_v120).length));
        let _index287 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_2035_v117).length));
        let _rhs345 = (_2036_v118)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(291)), ((_2040_v17) => function (_2041_i12) {
          return _2040_v17;
        })(_1849_v17))).length), new BigNumber((_2036_v118).length))];
        let _rhs346 = (_2037_v119)[_module.__default.safeIndex((((_2039_v121).contains(new BigNumber(974))) ? ((_2039_v121).get(new BigNumber(974))) : (p2)), new BigNumber((_2037_v119).length))];
        let _lhs286 = _2035_v117;
        let _lhs287 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_2035_v117).length));
        _1812_v2 = _rhs345;
        _lhs286[_lhs287] = _rhs346;
        (globalState).f5 = _module.__default.safeModuloInt(p2, new BigNumber(-255));
        let _2042_v122;
        _2042_v122 = _dafny.MultiSet.fromElements(_1997_cf20);
        let _2043_v123;
        _2043_v123 = _dafny.Map.Empty.slice().updateUnsafe(_2042_v122,p1);
        _2043_v123 = (_2043_v123).update(_dafny.MultiSet.fromElements(_1997_cf20, _1997_cf20, (((_1809_v0).contains(p3)) ? ((_1809_v0).get(p3)) : (p0))), p3);
      } else {
        let _2044___mcc_h17 = (_source22).cf12;
        let _2045_cf12 = _2044___mcc_h17;
        let _2046_v124;
        let _init50 = ((_2047_p2) => function (_2048_i13) {
          return (_2048_i13).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_2047_p2, _2047_p2, _2047_p2)).length)));
        })(p2);
        let _nw303 = Array((new BigNumber(18)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw303.length); _i0_50++) {
          _nw303[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _2046_v124 = _nw303;
        let _index288 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length));
        (_2046_v124)[_index288] = p2;
        let _index289 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length));
        (_2046_v124)[_index289] = p2;
        let _2049_v125;
        _2049_v125 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.__default.fm2(p2, globalState));
        let _2050_v126;
        _2050_v126 = _dafny.Map.Empty.slice().updateUnsafe(_2049_v125,_module.__default.safeModuloInt((_2046_v124)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length))], p2));
        let _index290 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length));
        (_2046_v124)[_index290] = (((_2050_v126).contains(_2049_v125)) ? ((_2050_v126).get(_2049_v125)) : (_module.__default.safeDivisionInt((_2046_v124)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length))], (_2046_v124)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length))])));
        let _2051_v127;
        _2051_v127 = _module.D4.create_DC13(_dafny.Seq.of(new BigNumber((_1849_v17).length), (_2046_v124)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length))], new BigNumber((_module.__default.fm17((_dafny.ZERO).minus(p2), p2, globalState)).length), (_dafny.ZERO).minus((_2046_v124)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length))]), p2), false, new BigNumber((_1812_v2).length));
        _2051_v127 = _2051_v127;
        let _2052_v128;
        let _nw304 = Array((new BigNumber(2)).toNumber());
        _nw304[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_1849_v17, _module.__default.safeIndex((_dafny.ZERO).minus((_2046_v124)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length))]), new BigNumber((_1849_v17).length)), _1848_v16);
        _nw304[(_dafny.ONE).toNumber()] = _1849_v17;
        _2052_v128 = _nw304;
        let _2053_v129;
        _2053_v129 = _module.D10.create_DC29(p2, _2052_v128, p3, p2);
        let _2054_v130;
        _2054_v130 = _dafny.Map.Empty.slice().updateUnsafe(!(p3),_2053_v129);
        let _index291 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_1810_v1).length));
        (_1810_v1)[_index291] = ((_2046_v124)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length))]).isLessThan(p2);
        let _2055_v131;
        _2055_v131 = _dafny.Map.Empty.slice().updateUnsafe(_1848_v16,((p1) ? (p1) : (p0)));
        let _2056_v132;
        _2056_v132 = _dafny.Seq.of(_2054_v130);
        let _2057_v133;
        _2057_v133 = _dafny.Map.Empty.slice().updateUnsafe((_2046_v124)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length))],_2054_v130);
        let _2058_v134;
        _2058_v134 = _dafny.Map.Empty.slice().updateUnsafe(p2,true);
        let _index292 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_1810_v1).length));
        let _rhs347 = new BigNumber(326);
        let _rhs348 = _2046_v124;
        let _rhs349 = (((p3) ? (((_2056_v132)[_module.__default.safeIndex((_2046_v124)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_2046_v124).length))], new BigNumber((_2056_v132).length))]).update(p0, _2053_v129)) : (_2054_v130))).Merge((((_2057_v133).contains(p2)) ? ((_2057_v133).get(p2)) : (_dafny.Map.Empty.slice().updateUnsafe(p0,_2053_v129))));
        let _rhs350 = (_module.D2.create_DC8((((_2058_v134).contains(new BigNumber(-284))) ? ((_2058_v134).get(new BigNumber(-284))) : (p1)), p3)).dtor_cf19;
        let _rhs351 = _dafny.Map.Empty.slice().updateUnsafe(_1848_v16,p0);
        let _lhs288 = _1810_v1;
        let _lhs289 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_1810_v1).length));
        r2 = _rhs347;
        _2046_v124 = _rhs348;
        _2054_v130 = _rhs349;
        _lhs288[_lhs289] = _rhs350;
        _2055_v131 = _rhs351;
      }
      let _2059_v135;
      _2059_v135 = _dafny.Set.fromElements(!(p0), p3);
      let _2060_v136;
      _2060_v136 = _module.D6.create_DC19(p2, (_this).fm11(globalState), new BigNumber(((_2059_v135).Union(_2059_v135)).length));
      let _2061_v137;
      let _nw305 = Array((new BigNumber(21)).toNumber());
      _2061_v137 = _nw305;
      let _2062_v138;
      _2062_v138 = _dafny.MultiSet.fromElements(_2061_v137, _2061_v137, _2061_v137);
      let _rhs352 = _2060_v136;
      let _rhs353 = new BigNumber(((_2062_v138).Union(_2062_v138)).cardinality());
      let _lhs290 = globalState;
      _2060_v136 = _rhs352;
      _lhs290.f5 = _rhs353;
      let _2063_v139;
      let _nw306 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _2063_v139 = _nw306;
      let _2064_v140;
      _2064_v140 = _module.D2.create_DC6(p2, _2063_v139);
      let _2065_v141;
      _2065_v141 = _dafny.Map.Empty.slice().updateUnsafe(_2064_v140,p0);
      let _2066_v142;
      _2066_v142 = _dafny.Seq.of(_2065_v141);
      r0 = ((_2065_v141).Merge(_2065_v141)).Merge(((false) ? (_2065_v141) : ((_2066_v142)[_module.__default.safeIndex(p2, new BigNumber((_2066_v142).length))])));
      r1 = _1809_v0;
      r2 = (_1812_v2)[_module.__default.safeIndex(p2, new BigNumber((_1812_v2).length))];
      return [r0, r1, r2];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _2067_v0;
      _2067_v0 = new BigNumber(372);
      let _2068_v1;
      _2068_v1 = _dafny.Map.Empty.slice().updateUnsafe(_2067_v0,new BigNumber(100));
      let _2069_v2;
      _2069_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2067_v0,_2068_v1);
      r2 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_2067_v0,_2068_v1)).Merge(_2069_v2)).length);
      _2067_v0 = _2067_v0;
      let _2070_v3;
      _2070_v3 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)));
      let _2071_v4;
      _2071_v4 = _module.D1.create_DC3(_2070_v3);
      let _2072_v5;
      _2072_v5 = new _dafny.CodePoint('o'.codePointAt(0));
      let _2073_v6;
      let _nw307 = Array((new BigNumber(23)).toNumber());
      _nw307[(_dafny.ZERO).toNumber()] = _2071_v4;
      _nw307[(_dafny.ONE).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(2)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(3)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(4)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(5)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(6)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(7)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(8)).toNumber()] = _module.D1.create_DC3(_2070_v3);
      _nw307[(new BigNumber(9)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(10)).toNumber()] = _module.D1.create_DC3(_2070_v3);
      _nw307[(new BigNumber(11)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(12)).toNumber()] = ((p0) ? (_2071_v4) : (_2071_v4));
      _nw307[(new BigNumber(13)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(14)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(15)).toNumber()] = _module.D1.create_DC3(_dafny.MultiSet.fromElements(_2072_v5));
      _nw307[(new BigNumber(16)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(17)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(18)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(19)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(20)).toNumber()] = ((p0) ? (_2071_v4) : (_2071_v4));
      _nw307[(new BigNumber(21)).toNumber()] = _2071_v4;
      _nw307[(new BigNumber(22)).toNumber()] = _2071_v4;
      _2073_v6 = _nw307;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2073_v6).length))) {
        let _2074_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2074_i0)) && ((_2074_i0).isLessThan(new BigNumber((_2073_v6).length))))) {
          (_2073_v6)[(_2074_i0)] = ((p0) ? (_2071_v4) : (_2071_v4));
        }
      }
      let _2075_v7;
      let _init51 = function (_2076_i1) {
        return _dafny.Seq.UnicodeFromString("sggnxrdx");
      };
      let _nw308 = Array((new BigNumber(10)).toNumber());
      for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw308.length); _i0_51++) {
        _nw308[_i0_51] = _init51(new BigNumber(_i0_51));
      }
      _2075_v7 = _nw308;
      let _2077_v8;
      _2077_v8 = _module.D10.create_DC29(_2067_v0, _2075_v7, p0, _2067_v0);
      let _2078_v9;
      let _nw309 = Array((_dafny.ONE).toNumber());
      _nw309[(_dafny.ZERO).toNumber()] = (_2077_v8).dtor_cf53;
      _2078_v9 = _nw309;
      let _2079_v10;
      _2079_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2078_v9,p0);
      _2079_v10 = (_2079_v10).update(_2078_v9, p0);
      let _2080_v11;
      _2080_v11 = _dafny.Set.fromElements(p1, p1, (p0) || (!(false)));
      _2080_v11 = (_module.D11.create_DC33(new BigNumber(-74), p0, (_this).fm11(globalState), p0, _2080_v11)).dtor_cf66;
      let _hi13 = _2067_v0;
      for (let _2081_i2 = (_dafny.ZERO).minus(_2067_v0); _2081_i2.isLessThan(_hi13); _2081_i2 = _2081_i2.plus(_dafny.ONE)) {
        let _2082_v12;
        _2082_v12 = _dafny.Seq.of(new BigNumber(428), _2067_v0);
        let _2083_v13;
        _2083_v13 = _module.D11.create_DC31(new BigNumber(67), (_2082_v12)[_module.__default.safeIndex(_2081_i2, new BigNumber((_2082_v12).length))]);
        let _pat_let_tv52 = _2068_v1;
        let _rhs354 = function (_pat_let70_0) {
          return function (_2084_dt__update__tmp_h0) {
            return function (_pat_let71_0) {
              return function (_2085_dt__update_hcf59_h0) {
                return _module.D11.create_DC31((_2084_dt__update__tmp_h0).dtor_cf58, _2085_dt__update_hcf59_h0);
              }(_pat_let71_0);
            }((_2081_i2).plus(new BigNumber((_pat_let_tv52).length)));
          }(_pat_let70_0);
        }(_2083_v13);
        let _rhs355 = (_2067_v0).multipliedBy(_2067_v0);
        _2083_v13 = _rhs354;
        r2 = _rhs355;
        let _2086_v14;
        _2086_v14 = _dafny.Seq.of(_2070_v3, _2070_v3);
        let _2087_v15;
        _2087_v15 = _dafny.Seq.UnicodeFromString("iguyj");
        let _2088_v16;
        _2088_v16 = _module.D4.create_DC12(_dafny.Seq.update(_2086_v14, _module.__default.safeIndex(_2081_i2, new BigNumber((_2086_v14).length)), (_2070_v3).update(_2072_v5, _module.__default.abs(new BigNumber((_2087_v15).length)))));
        _2088_v16 = _2088_v16;
        let _2089_v17;
        let _nw310 = Array((new BigNumber(23)).toNumber()).fill(false);
        _2089_v17 = _nw310;
        let _2090_v18;
        _2090_v18 = _2089_v17;
        let _2091_v19;
        _2091_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2067_v0,_dafny.Seq.UnicodeFromString("by"));
        let _2092_v20;
        let _nw311 = new _module.C3();
        _nw311.__ctor(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(87),_2067_v0), (_2090_v18), _2091_v19, _2081_i2);
        _2092_v20 = _nw311;
        _2067_v0 = _2067_v0;
      }
      r0 = p1;
      let _2093_v21;
      _2093_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2067_v0,p1);
      let _2094_v22;
      _2094_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2093_v21,p1);
      let _2095_v23;
      _2095_v23 = _module.D3.create_DC10(_2094_v22);
      let _pat_let_tv53 = p0;
      let _pat_let_tv54 = p1;
      let _pat_let_tv55 = p0;
      r1 = !(function (_source25) {
        if (_source25.is_DC10) {
          let _2096___mcc_h0 = (_source25).cf22;
          let _2097_cf22 = _2096___mcc_h0;
          return _pat_let_tv53;
        } else if (_source25.is_DC9) {
          let _2098___mcc_h1 = (_source25).cf21;
          let _2099_cf21 = _2098___mcc_h1;
          return _pat_let_tv54;
        } else {
          let _2100___mcc_h2 = (_source25).cf23;
          let _2101_cf23 = _2100___mcc_h2;
          return _pat_let_tv55;
        }
      }(_2095_v23));
      r2 = (_dafny.ZERO).minus(_2067_v0);
      return [r0, r1, r2];
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))), _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(26)), function (_2102_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })), _dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))), _dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))), (_module.D1.create_DC3(_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0))))).dtor_cf6, _dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0))))));
    };
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.of((_this).f6));
    };
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _2103_v0;
      let _nw312 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
      _2103_v0 = _nw312;
      let _2104_v1;
      _2104_v1 = _dafny.Seq.UnicodeFromString("hbqce");
      let _2105_v2;
      _2105_v2 = _dafny.Set.fromElements((_this).f6, new BigNumber((_2104_v1).length));
      let _2106_v3;
      _2106_v3 = _dafny.Set.fromElements(false);
      let _2107_v4;
      _2107_v4 = _dafny.Seq.of(new BigNumber((_2106_v3).length));
      let _index293 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_2103_v0).length));
      (_2103_v0)[_index293] = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_2105_v2).length)), _dafny.Seq.update(_2107_v4, _module.__default.safeIndex((_this).f6, new BigNumber((_2107_v4).length)), (_this).f6));
      let _index294 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_2103_v0).length));
      (_2103_v0)[_index294] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_2107_v4, _2107_v4), _2107_v4), _module.__default.safeIndex((_this).f6, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2107_v4, _2107_v4), _2107_v4)).length)), new BigNumber((_2104_v1).length));
      let _2108_v5;
      _2108_v5 = _dafny.Seq.of(p2, p2);
      (globalState).f0 = !((((p0) ? (new BigNumber((_2108_v5).length)) : ((_this).f6))).isLessThanOrEqualTo((_this).f6));
      let _2109_v6;
      _2109_v6 = _module.D0.create_DC0(_2108_v5, (_this).f6);
      let _source26 = _2109_v6;
      if (_source26.is_DC0) {
        let _2110___mcc_h0 = (_source26).cf0;
        let _2111___mcc_h1 = (_source26).cf1;
        let _2112_cf1 = _2111___mcc_h1;
        let _2113_cf0 = _2110___mcc_h0;
        let _2114_v7;
        _2114_v7 = new _dafny.CodePoint('e'.codePointAt(0));
        let _2115_v8;
        _2115_v8 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_2112_cf1) : ((_this).f6)),(_this).f6);
        let _2116_v9;
        _2116_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,new _dafny.CodePoint('v'.codePointAt(0)));
        let _rhs356 = true;
        let _rhs357 = p0;
        let _rhs358 = (((_2116_v9).contains(_dafny.Seq.IsPrefixOf(_2104_v1, _2104_v1))) ? ((_2116_v9).get(_dafny.Seq.IsPrefixOf(_2104_v1, _2104_v1))) : (_2114_v7));
        let _rhs359 = (_2115_v8).update(_module.__default.fm2(_2112_cf1, globalState), (_this).f6);
        let _lhs291 = globalState;
        let _lhs292 = globalState;
        _lhs291.f0 = _rhs356;
        _lhs292.f0 = _rhs357;
        _2114_v7 = _rhs358;
        _2115_v8 = _rhs359;
        (globalState).f0 = false;
        if (p0) {
          let _2117_v10;
          _2117_v10 = _dafny.Seq.of(_2104_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-641)), ((_2118_v7) => function (_2119_i1) {
            return _2118_v7;
          })(_2114_v7)), _dafny.Seq.update(_2104_v1, _module.__default.safeIndex(_2112_cf1, new BigNumber((_2104_v1).length)), _2114_v7));
          _2104_v1 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(885)), ((_2120_v7) => function (_2121_i0) {
            return _2120_v7;
          })(_2114_v7)), (_2117_v10)[_module.__default.safeIndex(new BigNumber((function () {
            let _coll53 = new _dafny.Map();
            for (const _compr_53 of (function () {
              let _coll54 = new _dafny.Map();
              for (const _compr_54 of _dafny.IntegerRange(new BigNumber(381), new BigNumber(-819))) {
                let _2122_v12 = _compr_54;
                if (((new BigNumber(381)).isLessThanOrEqualTo(_2122_v12)) && ((_2122_v12).isLessThan(new BigNumber(-819)))) {
                  _coll54.push([(_2122_v12).minus((_this).f6),new BigNumber((_dafny.MultiSet.fromElements(!(p0), p2)).cardinality())]);
                }
              }
              return _coll54;
            }()).Keys.Elements) {
              let _2123_v11 = _compr_53;
              if ((function () {
                let _coll55 = new _dafny.Map();
                for (const _compr_55 of _dafny.IntegerRange(new BigNumber(381), new BigNumber(-819))) {
                  let _2122_v12 = _compr_55;
                  if (((new BigNumber(381)).isLessThanOrEqualTo(_2122_v12)) && ((_2122_v12).isLessThan(new BigNumber(-819)))) {
                    _coll55.push([(_2122_v12).minus((_this).f6),new BigNumber((_dafny.MultiSet.fromElements(!(p0), p2)).cardinality())]);
                  }
                }
                return _coll55;
              }()).contains(_2123_v11)) {
                _coll53.push([(_2123_v11).minus(_2112_cf1),_2112_cf1]);
              }
            }
            return _coll53;
          }()).length), new BigNumber((_2117_v10).length))]);
          (globalState).f0 = p2;
          let _2124_v13;
          _2124_v13 = _dafny.MultiSet.fromElements(p1, p1);
          let _2125_v14;
          _2125_v14 = _dafny.MultiSet.fromElements(_2104_v1);
          let _2126_v15;
          _2126_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2125_v14);
          let _2127_v16;
          _2127_v16 = _dafny.Map.Empty.slice().updateUnsafe((_2124_v13).IsProperSubsetOf(_2124_v13),(((_2126_v15).contains(p0)) ? ((_2126_v15).get(p0)) : (_dafny.MultiSet.fromElements(_2104_v1, _2104_v1, _2104_v1))));
          _2127_v16 = (_2127_v16).update((_module.__default.fm3(globalState)) === (p0), _2125_v14);
          (globalState).f0 = (p0) || (false);
          let _index295 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((p1).length));
          (p1)[_index295] = (_this).f6;
          let _index296 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((p1).length));
          (p1)[_index296] = _module.__default.fm2(_module.__default.fm2(_2112_cf1, globalState), globalState);
        } else {
          let _2128_v17;
          _2128_v17 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_2112_cf1, (_this).f6)).length), (_this).f6, (_this).f6, _2112_cf1);
          (globalState).f0 = (_dafny.MultiSet.fromElements((_this).f6, (_this).f6)).IsDisjointFrom(_2128_v17);
          _2109_v6 = _2109_v6;
          _2115_v8 = (_2115_v8).update((_this).f6, (_this).f6);
          let _2129_v18;
          let _init52 = function (_2130_i2) {
            return false;
          };
          let _nw313 = Array((_dafny.ONE).toNumber());
          for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw313.length); _i0_52++) {
            _nw313[_i0_52] = _init52(new BigNumber(_i0_52));
          }
          _2129_v18 = _nw313;
          let _index297 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_2129_v18).length));
          (_2129_v18)[_index297] = !(p0);
          let _2131_v19;
          _2131_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2112_cf1,false);
          let _index298 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_2129_v18).length));
          (_2129_v18)[_index298] = ((_2131_v19).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2112_cf1,p0))).equals((_2131_v19).Merge(_2131_v19));
          let _2132_v20;
          _2132_v20 = _dafny.MultiSet.fromElements(p2, (((_2131_v19).contains(new BigNumber(158))) ? ((_2131_v19).get(new BigNumber(158))) : (p0)), p2);
          let _2133_v21;
          _2133_v21 = _module.D1.create_DC4(_2109_v6, new BigNumber((_2132_v20).cardinality()), _2104_v1, (_2108_v5)[_module.__default.safeIndex(_2112_cf1, new BigNumber((_2108_v5).length))], _2112_cf1);
          let _2134_v22;
          _2134_v22 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.fm2(new BigNumber(971), globalState)),_2133_v21);
          _2134_v22 = _2134_v22;
        }
        let _2135_v23;
        let _nw314 = Array((new BigNumber(2)).toNumber()).fill([]);
        _2135_v23 = _nw314;
        let _2136_v24;
        let _nw315 = Array((new BigNumber(6)).toNumber()).fill(_module.D0.Default());
        _2136_v24 = _nw315;
        let _index299 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_2135_v23).length));
        (_2135_v23)[_index299] = _2136_v24;
        let _index300 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_2135_v23).length));
        let _init53 = ((_2137_v6) => function (_2138_i3) {
          return _2137_v6;
        })(_2109_v6);
        let _nw316 = Array((new BigNumber(6)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw316.length); _i0_53++) {
          _nw316[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        (_2135_v23)[_index300] = _nw316;
      } else if (_source26.is_DC1) {
        let _2139___mcc_h2 = (_source26).cf2;
        let _2140___mcc_h3 = (_source26).cf3;
        let _2141___mcc_h4 = (_source26).cf4;
        let _2142_cf4 = _2141___mcc_h4;
        let _2143_cf3 = _2140___mcc_h3;
        let _2144_cf2 = _2139___mcc_h2;
        _2143_cf3 = (_this).f6;
        r0 = !(p0) || (p0);
        let _2145_v25;
        let _init54 = ((_2146_v1) => function (_2147_i4) {
          return _2146_v1;
        })(_2104_v1);
        let _nw317 = Array((new BigNumber(10)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw317.length); _i0_54++) {
          _nw317[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _2145_v25 = _nw317;
        let _2148_v26;
        _2148_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2145_v25,true);
        _2148_v26 = (_2148_v26).update(_2145_v25, p0);
        _2143_cf3 = ((p0) ? (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(p2, p2)).length), _2143_cf3)) : (_2143_cf3));
      } else {
        let _2149___mcc_h5 = (_source26).cf5;
        let _2150_cf5 = _2149___mcc_h5;
        (globalState).f5 = (_this).f6;
        (globalState).f5 = (_this).f6;
        let _2151_v27;
        _2151_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
        let _2152_v31;
        let _nw318 = Array((new BigNumber(16)).toNumber());
        _nw318[(_dafny.ZERO).toNumber()] = _2151_v27;
        _nw318[(_dafny.ONE).toNumber()] = (_2151_v27).update((_this).f6, new BigNumber(-510));
        _nw318[(new BigNumber(2)).toNumber()] = _2151_v27;
        _nw318[(new BigNumber(3)).toNumber()] = _2151_v27;
        _nw318[(new BigNumber(4)).toNumber()] = (_2151_v27).Merge(_2151_v27);
        _nw318[(new BigNumber(5)).toNumber()] = _2151_v27;
        _nw318[(new BigNumber(6)).toNumber()] = ((p2) ? (_2151_v27) : (_2151_v27));
        _nw318[(new BigNumber(7)).toNumber()] = function () {
          let _coll56 = new _dafny.Map();
          for (const _compr_56 of _dafny.IntegerRange(new BigNumber(-1), new BigNumber(458))) {
            let _2153_v28 = _compr_56;
            if (((new BigNumber(-1)).isLessThanOrEqualTo(_2153_v28)) && ((_2153_v28).isLessThan(new BigNumber(458)))) {
              _coll56.push([_module.__default.safeModuloInt(_2153_v28, new BigNumber(76)),(_this).f6]);
            }
          }
          return _coll56;
        }();
        _nw318[(new BigNumber(8)).toNumber()] = function () {
          let _coll57 = new _dafny.Map();
          for (const _compr_57 of _dafny.IntegerRange(new BigNumber(52), new BigNumber(438))) {
            let _2154_v29 = _compr_57;
            if (((new BigNumber(52)).isLessThanOrEqualTo(_2154_v29)) && ((_2154_v29).isLessThan(new BigNumber(438)))) {
              _coll57.push([(_2154_v29).plus((_this).f6),(_this).f6]);
            }
          }
          return _coll57;
        }();
        _nw318[(new BigNumber(9)).toNumber()] = _module.__default.fm9(globalState);
        _nw318[(new BigNumber(10)).toNumber()] = (_2151_v27).Merge((_2151_v27).update(new BigNumber((_2108_v5).length), (_this).f6));
        _nw318[(new BigNumber(11)).toNumber()] = function () {
          let _coll58 = new _dafny.Map();
          for (const _compr_58 of (_2107_v4).Elements) {
            let _2155_v30 = _compr_58;
            if (_dafny.Seq.contains(_2107_v4, _2155_v30)) {
              _coll58.push([(_2155_v30).plus(new BigNumber((_dafny.Seq.update(_dafny.Seq.update((_2103_v0)[_module.__default.safeIndex(new BigNumber(841), new BigNumber((_2103_v0).length))], _module.__default.safeIndex((_this).f6, new BigNumber(((_2103_v0)[_module.__default.safeIndex(new BigNumber(841), new BigNumber((_2103_v0).length))]).length)), new BigNumber((_dafny.Seq.of(p2, p2)).length)), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,p0)).length), new BigNumber((_dafny.Seq.update((_2103_v0)[_module.__default.safeIndex(new BigNumber(841), new BigNumber((_2103_v0).length))], _module.__default.safeIndex((_this).f6, new BigNumber(((_2103_v0)[_module.__default.safeIndex(new BigNumber(841), new BigNumber((_2103_v0).length))]).length)), new BigNumber((_dafny.Seq.of(p2, p2)).length))).length)), (_this).f6)).length)),(_this).f6]);
            }
          }
          return _coll58;
        }();
        _nw318[(new BigNumber(12)).toNumber()] = _2151_v27;
        _nw318[(new BigNumber(13)).toNumber()] = (_module.D2.create_DC5(_2151_v27)).dtor_cf12;
        _nw318[(new BigNumber(14)).toNumber()] = _2151_v27;
        _nw318[(new BigNumber(15)).toNumber()] = (_2151_v27).Merge(_2151_v27);
        _2152_v31 = _nw318;
        _2152_v31 = (((p2) || (!(p2))) ? (_2152_v31) : (_2152_v31));
        let _rhs360 = (_2107_v4)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f6), new BigNumber((_2107_v4).length))];
        let _rhs361 = p2;
        let _rhs362 = _2104_v1;
        let _rhs363 = (_this).f6;
        let _lhs293 = globalState;
        let _lhs294 = globalState;
        _lhs293.f5 = _rhs360;
        r0 = _rhs361;
        _2104_v1 = _rhs362;
        _lhs294.f5 = _rhs363;
      }
      let _2156_v32;
      _2156_v32 = new _dafny.CodePoint('f'.codePointAt(0));
      let _2157_v33;
      _2157_v33 = _module.D0.create_DC1(p0, (_this).f6, _2156_v32);
      let _source27 = ((p0) ? (_2157_v33) : (_2157_v33));
      if (_source27.is_DC0) {
        let _2158___mcc_h6 = (_source27).cf0;
        let _2159___mcc_h7 = (_source27).cf1;
        let _2160_cf1 = _2159___mcc_h7;
        let _2161_cf0 = _2158___mcc_h6;
        _2160_cf1 = (_this).f6;
        let _2162_v34;
        _2162_v34 = _dafny.MultiSet.fromElements(_2156_v32);
        let _2163_v35;
        let _nw319 = Array((new BigNumber(28)).toNumber());
        _nw319[(_dafny.ZERO).toNumber()] = p2;
        _nw319[(_dafny.ONE).toNumber()] = !(p2);
        _nw319[(new BigNumber(2)).toNumber()] = p2;
        _nw319[(new BigNumber(3)).toNumber()] = p0;
        _nw319[(new BigNumber(4)).toNumber()] = p2;
        _nw319[(new BigNumber(5)).toNumber()] = p0;
        _nw319[(new BigNumber(6)).toNumber()] = p2;
        _nw319[(new BigNumber(7)).toNumber()] = p2;
        _nw319[(new BigNumber(8)).toNumber()] = false;
        _nw319[(new BigNumber(9)).toNumber()] = p2;
        _nw319[(new BigNumber(10)).toNumber()] = _module.__default.fm3(globalState);
        _nw319[(new BigNumber(11)).toNumber()] = _module.__default.fm3(globalState);
        _nw319[(new BigNumber(12)).toNumber()] = p2;
        _nw319[(new BigNumber(13)).toNumber()] = false;
        _nw319[(new BigNumber(14)).toNumber()] = p0;
        _nw319[(new BigNumber(15)).toNumber()] = p2;
        _nw319[(new BigNumber(16)).toNumber()] = p0;
        _nw319[(new BigNumber(17)).toNumber()] = p2;
        _nw319[(new BigNumber(18)).toNumber()] = p2;
        _nw319[(new BigNumber(19)).toNumber()] = p0;
        _nw319[(new BigNumber(20)).toNumber()] = p0;
        _nw319[(new BigNumber(21)).toNumber()] = p0;
        _nw319[(new BigNumber(22)).toNumber()] = p2;
        _nw319[(new BigNumber(23)).toNumber()] = p0;
        _nw319[(new BigNumber(24)).toNumber()] = p2;
        _nw319[(new BigNumber(25)).toNumber()] = p2;
        _nw319[(new BigNumber(26)).toNumber()] = p0;
        _nw319[(new BigNumber(27)).toNumber()] = p2;
        _2163_v35 = _nw319;
        let _2164_v36;
        _2164_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2104_v1);
        let _2165_v37;
        _2165_v37 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_2107_v4));
        let _2166_v38;
        let _nw320 = new _module.C4();
        _nw320.__ctor(_2160_cf1, (((_2162_v34).contains(_2156_v32)) ? ((_2162_v34).get(_2156_v32)) : (new BigNumber((_2104_v1).length))), ((p0) ? (_2163_v35) : (_2163_v35)), (_2164_v36).Merge(_2164_v36), _2160_cf1, _2165_v37);
        _2166_v38 = _nw320;
        let _2167_v39;
        _2167_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p2);
        let _2168_v40;
        _2168_v40 = _module.D1.create_DC4(_2109_v6, _module.__default.fm2(new BigNumber((_2167_v39).length), globalState), _2104_v1, false, (_2166_v38).f18);
        let _2169_v41;
        _2169_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2168_v40,_2163_v35);
        (globalState).f5 = (_dafny.ZERO).minus(new BigNumber((_2169_v41).length));
        (globalState).f0 = (p2) === (p2);
      } else if (_source27.is_DC1) {
        let _2170___mcc_h8 = (_source27).cf2;
        let _2171___mcc_h9 = (_source27).cf3;
        let _2172___mcc_h10 = (_source27).cf4;
        let _2173_cf4 = _2172___mcc_h10;
        let _2174_cf3 = _2171___mcc_h9;
        let _2175_cf2 = _2170___mcc_h8;
        let _2176_v42;
        _2176_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2157_v33,!(_2175_cf2));
        (globalState).f5 = ((_dafny.ZERO).minus((_this).f6)).multipliedBy(new BigNumber((_2176_v42).length));
        _2174_cf3 = (((_2108_v5)[_module.__default.safeIndex((_this).f6, new BigNumber((_2108_v5).length))]) ? (((_this).f6).minus(new BigNumber(169))) : ((_this).f6));
        let _2177_v43;
        let _nw321 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _2177_v43 = _nw321;
        _2177_v43 = p1;
        let _2178_v44;
        let _init55 = ((_2179_p2, _2180_p0) => function (_2181_i5) {
          return _dafny.Map.Empty.slice().updateUnsafe(!(!(_2179_p2)),_2180_p0);
        })(p2, p0);
        let _nw322 = Array((new BigNumber(21)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw322.length); _i0_55++) {
          _nw322[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _2178_v44 = _nw322;
        let _2182_v45;
        _2182_v45 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(globalState),p2);
        let _index301 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_2178_v44).length));
        (_2178_v44)[_index301] = _2182_v45;
        let _index302 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_2178_v44).length));
        (_2178_v44)[_index302] = _2182_v45;
      } else {
        let _2183___mcc_h11 = (_source27).cf5;
        let _2184_cf5 = _2183___mcc_h11;
        (globalState).f5 = ((_this).f6).plus((_this).f6);
        let _2185_v46;
        _2185_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber((_dafny.Seq.update(_2104_v1, _module.__default.safeIndex((_this).f6, new BigNumber((_2104_v1).length)), _2156_v32)).length));
        _2185_v46 = (_2185_v46).update((_this).f6, (_this).f6);
        let _2186_v47;
        _2186_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(789),!(p0));
        let _rhs364 = (((((_2186_v47).contains((_this).f6)) ? ((_2186_v47).get((_this).f6)) : (p2))) ? (_dafny.Set.fromElements((_this).f6, (_this).f6)) : (_dafny.Set.fromElements((_this).f6)));
        let _rhs365 = ((_this).f6).isLessThanOrEqualTo((_this).f6);
        _2105_v2 = _rhs364;
        r0 = _rhs365;
        (globalState).f0 = p2;
      }
      let _2187_v48;
      _2187_v48 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_dafny.Seq.UnicodeFromString("icjgt")).length));
      let _2188_v49;
      _2188_v49 = _module.D14.create_DC36(_2187_v48);
      let _2189_v50;
      let _nw323 = Array((new BigNumber(2)).toNumber());
      _nw323[(_dafny.ZERO).toNumber()] = ((!(p2)) ? (_2188_v49) : (_2188_v49));
      _nw323[(_dafny.ONE).toNumber()] = _2188_v49;
      _2189_v50 = _nw323;
      let _2190_v51;
      _2190_v51 = _dafny.Seq.of(_module.__default.fm41(globalState), _dafny.Map.Empty.slice().updateUnsafe(p0,p0));
      let _2191_v52;
      _2191_v52 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
      let _2192_v53;
      _2192_v53 = _dafny.MultiSet.fromElements(_2191_v52, _2191_v52);
      let _rhs366 = new BigNumber(6);
      let _rhs367 = (_2192_v53).IsSubsetOf((_dafny.MultiSet.FromArray(_2190_v51)).Difference(_2192_v53));
      let _rhs368 = _2189_v50;
      let _lhs295 = globalState;
      let _lhs296 = globalState;
      _lhs295.f5 = _rhs366;
      _lhs296.f0 = _rhs367;
      _2189_v50 = _rhs368;
      let _2193_v54;
      _2193_v54 = _dafny.MultiSet.fromElements(p0, _module.__default.fm3(globalState));
      let _2194_v55;
      _2194_v55 = _dafny.Map.Empty.slice().updateUnsafe(_2193_v54,new _dafny.CodePoint('a'.codePointAt(0)));
      let _2195_v56;
      _2195_v56 = _dafny.Seq.of(_2194_v55);
      if ((_module.__default.safeModuloInt(new BigNumber(((_2195_v56)[_module.__default.safeIndex((_this).f6, new BigNumber((_2195_v56).length))]).length), (_this).f6)).isEqualTo((_this).f6)) {
        let _2196_v57;
        let _nw324 = new _module.C6();
        _nw324.__ctor((_this).f6, true);
        _2196_v57 = _nw324;
        let _2197_v58;
        _2197_v58 = _module.D6.create_DC19((_this).f6, (_this).f6, new BigNumber((_2105_v2).length));
        let _pat_let_tv56 = _2196_v57;
        _2197_v58 = function (_pat_let72_0) {
          return function (_2198_dt__update__tmp_h0) {
            return function (_pat_let73_0) {
              return function (_2199_dt__update_hcf41_h0) {
                return _module.D6.create_DC19((_2198_dt__update__tmp_h0).dtor_cf39, (_2198_dt__update__tmp_h0).dtor_cf40, _2199_dt__update_hcf41_h0);
              }(_pat_let73_0);
            }((_pat_let_tv56).f13);
          }(_pat_let72_0);
        }(_module.D6.create_DC19(new BigNumber((_2104_v1).length), (_2196_v57).f13, (_2196_v57).f13));
        let _2200_v59;
        _2200_v59 = _dafny.Map.Empty.slice().updateUnsafe(((_2196_v57).f14) && (p2),(_dafny.Map.Empty.slice().updateUnsafe(!(!(p2)),new BigNumber(-942))).update(p2, (_this).f6));
        _2200_v59 = (_2200_v59).update(!((_2196_v57).f14), ((_2187_v48).update(true, (_2196_v57).f13)).update(false, (_this).f6));
        let _2201_v60;
        let _nw325 = Array((new BigNumber(2)).toNumber()).fill([]);
        _2201_v60 = _nw325;
        _2201_v60 = _2201_v60;
        let _2202_v61;
        let _init56 = ((_2203_v6) => function (_2204_i6) {
          return _2203_v6;
        })(_2109_v6);
        let _nw326 = Array((new BigNumber(22)).toNumber());
        for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw326.length); _i0_56++) {
          _nw326[_i0_56] = _init56(new BigNumber(_i0_56));
        }
        _2202_v61 = _nw326;
        let _2205_v62;
        _2205_v62 = _module.D7.create_DC21(_2202_v61);
        let _source28 = _2205_v62;
        if (_source28.is_DC22) {
          let _index303 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((p1).length));
          (p1)[_index303] = ((_2196_v57).f13).minus((_this).f6);
          let _index304 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((p1).length));
          (p1)[_index304] = (_this).f6;
          let _rhs369 = p2;
          let _rhs370 = (_dafny.ZERO).minus((_2107_v4)[_module.__default.safeIndex((p1)[_module.__default.safeIndex(new BigNumber(11), new BigNumber((p1).length))], new BigNumber((_2107_v4).length))]);
          let _rhs371 = _2104_v1;
          let _rhs372 = (_this).f6;
          let _lhs297 = globalState;
          let _lhs298 = globalState;
          let _lhs299 = globalState;
          let _lhs300 = globalState;
          _lhs297.f0 = _rhs369;
          _lhs298.f5 = _rhs370;
          _lhs299.f2 = _rhs371;
          _lhs300.f5 = _rhs372;
          let _2206_v63;
          let _nw327 = new _module.C1();
          _nw327.__ctor((_this).f6);
          _2206_v63 = _nw327;
          (globalState).f0 = p0;
        } else {
          let _2207___mcc_h12 = (_source28).cf43;
          let _2208_cf43 = _2207___mcc_h12;
          (globalState).f0 = (_2196_v57).f14;
          let _2209_v64;
          let _nw328 = new _module.C1();
          _nw328.__ctor(new BigNumber(469));
          _2209_v64 = _nw328;
          let _2210_v65;
          _2210_v65 = _dafny.Map.Empty.slice().updateUnsafe(_2187_v48,_2196_v57);
          let _2211_v66;
          _2211_v66 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2196_v57);
          let _2212_v67;
          let _nw329 = Array((new BigNumber(14)).toNumber());
          _nw329[(_dafny.ZERO).toNumber()] = (((_2210_v65).contains(_2187_v48)) ? ((_2210_v65).get(_2187_v48)) : (_2196_v57));
          _nw329[(_dafny.ONE).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(2)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(3)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(4)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(5)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(6)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(7)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(8)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(9)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(10)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(11)).toNumber()] = _2196_v57;
          _nw329[(new BigNumber(12)).toNumber()] = (((_2211_v66).contains((_2209_v64).fm28(globalState))) ? ((_2211_v66).get((_2209_v64).fm28(globalState))) : (_2196_v57));
          _nw329[(new BigNumber(13)).toNumber()] = _2196_v57;
          _2212_v67 = _nw329;
          _2212_v67 = _2212_v67;
          _2191_v52 = (_2191_v52).update((_2196_v57).f14, ((_2196_v57).f13).isLessThanOrEqualTo((_2196_v57).f13));
        }
      } else {
        let _2213_v68;
        _2213_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-195),_2191_v52);
        let _2214_v69;
        _2214_v69 = _dafny.MultiSet.fromElements(p1, p1);
        let _2215_v71;
        _2215_v71 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2191_v52), _2213_v68);
        let _2216_v72;
        let _nw330 = Array((new BigNumber(29)).toNumber());
        _nw330[(_dafny.ZERO).toNumber()] = _2213_v68;
        _nw330[(_dafny.ONE).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(2)).toNumber()] = (_2213_v68).Merge(_2213_v68);
        _nw330[(new BigNumber(3)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(4)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(5)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(6)).toNumber()] = (_2213_v68).update(new BigNumber((_2214_v69).cardinality()), _2191_v52);
        _nw330[(new BigNumber(7)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(8)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(9)).toNumber()] = (_2213_v68).Merge(_2213_v68);
        _nw330[(new BigNumber(10)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(11)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(12)).toNumber()] = function () {
          let _coll59 = new _dafny.Map();
          for (const _compr_59 of _dafny.IntegerRange(new BigNumber(604), new BigNumber(-336))) {
            let _2217_v70 = _compr_59;
            if (((new BigNumber(604)).isLessThanOrEqualTo(_2217_v70)) && ((_2217_v70).isLessThan(new BigNumber(-336)))) {
              _coll59.push([(_2217_v70).plus((_this).f6),_2191_v52]);
            }
          }
          return _coll59;
        }();
        _nw330[(new BigNumber(13)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(14)).toNumber()] = (_2213_v68).update(new BigNumber((_2108_v5).length), _2191_v52);
        _nw330[(new BigNumber(15)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(16)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(17)).toNumber()] = (_2213_v68).Merge(_2213_v68);
        _nw330[(new BigNumber(18)).toNumber()] = (_2213_v68).Merge((_2215_v71)[_module.__default.safeIndex(new BigNumber(((_2103_v0)[_module.__default.safeIndex(new BigNumber(841), new BigNumber((_2103_v0).length))]).length), new BigNumber((_2215_v71).length))]);
        _nw330[(new BigNumber(19)).toNumber()] = (_2213_v68).Merge(_2213_v68);
        _nw330[(new BigNumber(20)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(21)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(22)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(854),_dafny.Map.Empty.slice().updateUnsafe(p0,p2));
        _nw330[(new BigNumber(24)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(25)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(26)).toNumber()] = _2213_v68;
        _nw330[(new BigNumber(27)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_2190_v51)[_module.__default.safeIndex(new BigNumber((_2104_v1).length), new BigNumber((_2190_v51).length))]);
        _nw330[(new BigNumber(28)).toNumber()] = _2213_v68;
        _2216_v72 = _nw330;
        let _index305 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2216_v72).length));
        (_2216_v72)[_index305] = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Map.Empty.slice().updateUnsafe(p0,p0));
        let _2218_v73;
        _2218_v73 = _module.D4.create_DC14(false, _2191_v52, (_this).f6, _2191_v52);
        let _index306 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_2216_v72).length));
        (_2216_v72)[_index306] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_2218_v73).dtor_cf29)).Merge(_2213_v68);
        (globalState).f5 = (_this).f6;
        if ((p0) || (p0)) {
          let _2219_v74;
          _2219_v74 = _dafny.Seq.of(_2108_v5, _2108_v5, _dafny.Seq.of(p2), _2108_v5);
          _2219_v74 = _2219_v74;
          (globalState).f0 = p0;
          let _2220_v75;
          let _nw331 = new _module.C5();
          _nw331.__ctor((_this).f6, (_this).f6);
          _2220_v75 = _nw331;
          let _2221_v76;
          _2221_v76 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains(_module.__default.fm26((_this).f6, _2156_v32, (_2108_v5)[_module.__default.safeIndex((_this).f6, new BigNumber((_2108_v5).length))], _2156_v32, globalState), (_this).f6),_2104_v1);
          _2221_v76 = (_2221_v76).update(p0, _2104_v1);
          let _2222_v77;
          _2222_v77 = _dafny.Seq.of(_2193_v54, _2193_v54, _2193_v54, _dafny.MultiSet.FromArray(_2108_v5), (_2193_v54).update(p2, _module.__default.abs((_dafny.ZERO).minus((_this).f6))));
          let _2223_v78;
          _2223_v78 = _dafny.Set.fromElements(_2107_v4, _2107_v4);
          let _2224_v79;
          _2224_v79 = _dafny.Map.Empty.slice().updateUnsafe(_2223_v78,_2222_v77);
          _2222_v77 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_dafny.MultiSet.fromElements(p0)).update(p0, _module.__default.abs(new BigNumber(683)))), _2222_v77), (((_2224_v79).contains(_2223_v78)) ? ((_2224_v79).get(_2223_v78)) : (_2222_v77)));
        } else {
          let _2225_v80;
          let _nw332 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2225_v80 = _nw332;
          let _2226_v81;
          _2226_v81 = _dafny.Map.Empty.slice().updateUnsafe(_2156_v32,_2225_v80);
          _2226_v81 = (_2226_v81).update(_module.__default.fm0((_this).f6, _2105_v2, p2, globalState), _2225_v80);
          let _2227_v82;
          _2227_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p2);
          let _2228_v83;
          _2228_v83 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(p2, true, (((_2227_v82).contains((_this).f6)) ? ((_2227_v82).get((_this).f6)) : (!(p0))))).cardinality()), (new BigNumber(843)).minus(new BigNumber((_2104_v1).length)), ((_this).f6).multipliedBy((_this).f6), (_this).f6, ((_this).f6).plus(new BigNumber(758)));
          let _2229_v84;
          _2229_v84 = _dafny.Seq.of(_2104_v1, _2104_v1);
          let _rhs373 = p2;
          let _rhs374 = ((_this).f6).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber(-848), (_this).f6));
          let _rhs375 = ((_this).f6).plus((_this).f6);
          let _rhs376 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_2229_v84, _2229_v84), _2229_v84);
          let _rhs377 = _2228_v83;
          let _lhs301 = globalState;
          let _lhs302 = globalState;
          let _lhs303 = globalState;
          let _lhs304 = globalState;
          _lhs301.f0 = _rhs373;
          _lhs302.f0 = _rhs374;
          _lhs303.f5 = _rhs375;
          _lhs304.f0 = _rhs376;
          _2228_v83 = _rhs377;
          _2227_v82 = (_2227_v82).update((_this).f6, ((_this).f6).isLessThan((_this).f6));
          (globalState).f0 = p2;
          let _index307 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((p1).length));
          (p1)[_index307] = (_this).f6;
          let _2230_v85;
          _2230_v85 = _dafny.Seq.of(_2157_v33);
          let _index308 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((p1).length));
          (p1)[_index308] = new BigNumber((_dafny.Seq.Concat(_2230_v85, _2230_v85)).length);
        }
        let _2231_v86;
        _2231_v86 = _dafny.Seq.of(_dafny.Seq.of(p2));
        let _2232_v87;
        _2232_v87 = _dafny.Seq.of(_2231_v86, _2231_v86, _2231_v86, _2231_v86);
        let _2233_v88;
        _2233_v88 = _dafny.Map.Empty.slice().updateUnsafe((_2232_v87)[_module.__default.safeIndex((_this).f6, new BigNumber((_2232_v87).length))],p0);
        _2233_v88 = (_2233_v88).update(_dafny.Seq.of(_2108_v5), _module.__default.fm3(globalState));
        let _2234_v89;
        let _nw333 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2234_v89 = _nw333;
        let _index309 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_2234_v89).length));
        (_2234_v89)[_index309] = _dafny.Seq.Concat(_2104_v1, _2104_v1);
        let _index310 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_2234_v89).length));
        (_2234_v89)[_index310] = _2104_v1;
      }
      r0 = p0;
      return r0;
    }
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('o'.codePointAt(0))))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(771)), function (_2235_i0) {
        return _dafny.MultiSet.FromArray((_module.D1.create_DC4(_module.D0.create_DC0(_dafny.Seq.of(true, !(true)), new BigNumber(796)), (_this).f6, _dafny.Seq.UnicodeFromString("lxfutcy"), false, (_this).f6)).dtor_cf9);
      })), ((false) ? (_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0))))) : (_dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0))), _dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0))))))));
    };
    fm6(p0, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.of(new BigNumber(471), new BigNumber((_dafny.Seq.of(false, false, false, false, false)).length)))).Merge(function () {
        let _coll60 = new _dafny.Map();
        for (const _compr_60 of _dafny.IntegerRange(new BigNumber(858), new BigNumber(166))) {
          let _2236_v0 = _compr_60;
          if (((new BigNumber(858)).isLessThanOrEqualTo(_2236_v0)) && ((_2236_v0).isLessThan(new BigNumber(166)))) {
            _coll60.push([_module.__default.safeModuloInt(_2236_v0, (_this).f6),_dafny.Seq.of((_this).f6, (_this).f6)]);
          }
        }
        return _coll60;
      }());
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _2237_v0;
      _2237_v0 = _dafny.Seq.of(new BigNumber((p0).length), (_this).f6);
      let _2238_v1;
      _2238_v1 = false;
      let _2239_v2;
      _2239_v2 = _dafny.Seq.of(_2238_v1, _2238_v1);
      let _2240_v3;
      _2240_v3 = _module.D0.create_DC0(_2239_v2, (_this).f6);
      let _2241_v4;
      _2241_v4 = _dafny.Seq.of(((_this).f6).multipliedBy((_2237_v0)[_module.__default.safeIndex((_2240_v3).dtor_cf1, new BigNumber((_2237_v0).length))]));
      (globalState).f5 = (_2237_v0)[_module.__default.safeIndex((_this).f6, new BigNumber((_2237_v0).length))];
      let _2242_v5;
      let _init57 = ((_2243_v1) => function (_2244_i0) {
        return _2243_v1;
      })(_2238_v1);
      let _nw334 = Array((new BigNumber(9)).toNumber());
      for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw334.length); _i0_57++) {
        _nw334[_i0_57] = _init57(new BigNumber(_i0_57));
      }
      _2242_v5 = _nw334;
      _2242_v5 = _2242_v5;
      let _2245_i1;
      _2245_i1 = _dafny.ZERO;
      L11: {
        while (_2238_v1) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2245_i1)) {
              break L11;
            }
            _2245_i1 = (_2245_i1).plus(_dafny.ONE);
            let _2246_v6;
            _2246_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(7),_2238_v1);
            let _2247_v7;
            let _nw335 = new _module.C6();
            _nw335.__ctor(new BigNumber((_2246_v6).length), _2238_v1);
            _2247_v7 = _nw335;
            let _2248_v8;
            _2248_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2247_v7);
            let _2249_v9;
            _2249_v9 = _module.D18.create_DC42(_2248_v8);
            let _2250_v10;
            _2250_v10 = _dafny.MultiSet.fromElements(new BigNumber(((_2249_v9).dtor_cf77).length));
            let _2251_v11;
            _2251_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_dafny.ZERO).minus((((_2250_v10).contains((_this).f6)) ? ((_2250_v10).get((_this).f6)) : (new BigNumber(-417)))));
            let _2252_v12;
            _2252_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p0);
            let _2253_v13;
            let _nw336 = new _module.C3();
            _nw336.__ctor(_2251_v11, _2242_v5, (_2252_v12).Merge(_2252_v12), (_2247_v7).f13);
            _2253_v13 = _nw336;
            let _2254_v14;
            _2254_v14 = _dafny.Map.Empty.slice().updateUnsafe((_2247_v7).f14,(_this).f6);
            let _2255_v15;
            _2255_v15 = _dafny.Seq.of(_2241_v4);
            let _2256_v16;
            _2256_v16 = _dafny.Set.fromElements((_2247_v7).f14);
            let _2257_v17;
            let _nw337 = Array((new BigNumber(28)).toNumber());
            _nw337[(_dafny.ZERO).toNumber()] = new BigNumber(-254);
            _nw337[(_dafny.ONE).toNumber()] = (_this).f6;
            _nw337[(new BigNumber(2)).toNumber()] = (((_2250_v10).contains(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f6, new BigNumber(953), new BigNumber(72)))).cardinality()))) ? ((_2250_v10).get(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f6, new BigNumber(953), new BigNumber(72)))).cardinality()))) : ((_this).f6));
            _nw337[(new BigNumber(3)).toNumber()] = (_this).f6;
            _nw337[(new BigNumber(4)).toNumber()] = (_2247_v7).f13;
            _nw337[(new BigNumber(5)).toNumber()] = (_2247_v7).f13;
            _nw337[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(((((_2250_v10).contains((_2241_v4)[_module.__default.safeIndex((_2247_v7).f13, new BigNumber((_2241_v4).length))])) ? ((_2250_v10).get((_2241_v4)[_module.__default.safeIndex((_2247_v7).f13, new BigNumber((_2241_v4).length))])) : ((_dafny.ZERO).minus(new BigNumber((_2239_v2).length))))).plus(new BigNumber((_2254_v14).length))));
            _nw337[(new BigNumber(7)).toNumber()] = (_this).f6;
            _nw337[(new BigNumber(8)).toNumber()] = (_this).f6;
            _nw337[(new BigNumber(9)).toNumber()] = (((_2247_v7).f14) ? ((_this).f6) : ((_dafny.ZERO).minus((_this).f6)));
            _nw337[(new BigNumber(10)).toNumber()] = (_this).f6;
            _nw337[(new BigNumber(11)).toNumber()] = (_2247_v7).f13;
            _nw337[(new BigNumber(12)).toNumber()] = (_2247_v7).f13;
            _nw337[(new BigNumber(13)).toNumber()] = (_2247_v7).f13;
            _nw337[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2237_v0, (_2255_v15)[_module.__default.safeIndex(new BigNumber((_2256_v16).length), new BigNumber((_2255_v15).length))])).length);
            _nw337[(new BigNumber(15)).toNumber()] = new BigNumber(175);
            _nw337[(new BigNumber(16)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(509), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(352)), function (_2258_i2) {
              return new _dafny.CodePoint('w'.codePointAt(0));
            })).length));
            _nw337[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt((_this).f6, new BigNumber((_dafny.Seq.UnicodeFromString("yl")).length));
            _nw337[(new BigNumber(18)).toNumber()] = (_this).f6;
            _nw337[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f6));
            _nw337[(new BigNumber(20)).toNumber()] = new BigNumber((p0).length);
            _nw337[(new BigNumber(21)).toNumber()] = ((_2247_v7).f13).plus(new BigNumber((_module.__default.fm19((_2247_v7).f13, globalState)).cardinality()));
            _nw337[(new BigNumber(22)).toNumber()] = (_2247_v7).f13;
            _nw337[(new BigNumber(23)).toNumber()] = _module.__default.safeDivisionInt((_2247_v7).f13, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_2247_v7).f13)).length)));
            _nw337[(new BigNumber(24)).toNumber()] = (_2247_v7).f13;
            _nw337[(new BigNumber(25)).toNumber()] = new BigNumber(118);
            _nw337[(new BigNumber(26)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((p0).length)), new BigNumber(((_module.D5.create_DC17(_2241_v4, _2238_v1)).dtor_cf36).length)));
            _nw337[(new BigNumber(27)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f6), (_2247_v7).f13);
            _2257_v17 = _nw337;
            let _2259_v18;
            _2259_v18 = _module.D2.create_DC8(_2238_v1, false);
            let _2260_v19;
            _2260_v19 = new _dafny.CodePoint('f'.codePointAt(0));
            let _index311 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_2257_v17).length));
            (_2257_v17)[_index311] = (_2247_v7).fm16(_2259_v18, _module.D0.create_DC1(_2238_v1, (_this).f6, _2260_v19), globalState);
            let _index312 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_2257_v17).length));
            (_2257_v17)[_index312] = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_this).f6, new BigNumber(-30)), (_2247_v7).f13);
            let _2261_v20;
            _2261_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(!(_2238_v1)));
            let _2262_v21;
            _2262_v21 = _dafny.Seq.of(p0);
            let _2263_v22;
            let _nw338 = new _module.C8();
            _nw338.__ctor();
            _2263_v22 = _nw338;
            let _2264_v23;
            _2264_v23 = _dafny.Seq.of(_2263_v22);
            _2261_v20 = (_2261_v20).update(_dafny.Seq.Concat((_2262_v21)[_module.__default.safeIndex(new BigNumber((_2264_v23).length), new BigNumber((_2262_v21).length))], p0), (_2247_v7).f14);
            let _2265_v24;
            _2265_v24 = _module.D9.create_DC25(new BigNumber(765), (((_2247_v7).f14) ? ((_2247_v7).f13) : ((_this).f6)), (_2257_v17)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_2257_v17).length))], _2260_v19);
            let _2266_v25;
            _2266_v25 = _dafny.MultiSet.fromElements(_2250_v10);
            let _2267_v26;
            _2267_v26 = _dafny.Seq.of(_module.__default.fm4((_this).f6, (_2257_v17)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_2257_v17).length))], (_2247_v7).f13, globalState));
            let _2268_v27;
            _2268_v27 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.update(_2237_v0, _module.__default.safeIndex((_2257_v17)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_2257_v17).length))], new BigNumber((_2237_v0).length)), (_2257_v17)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_2257_v17).length))])), _2250_v10);
            let _rhs378 = ((_2239_v2)[_module.__default.safeIndex((_this).f6, new BigNumber((_2239_v2).length))]) && ((_2239_v2)[_module.__default.safeIndex((_2265_v24).dtor_cf48, new BigNumber((_2239_v2).length))]);
            let _rhs379 = _2242_v5;
            let _rhs380 = _2265_v24;
            let _rhs381 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_2239_v2), _2267_v26);
            let _rhs382 = ((_2266_v25).Difference(_dafny.MultiSet.FromArray(_2268_v27))).Union(_2266_v25);
            let _lhs305 = globalState;
            let _lhs306 = globalState;
            _lhs305.f0 = _rhs378;
            _2242_v5 = _rhs379;
            _2265_v24 = _rhs380;
            _lhs306.f0 = _rhs381;
            _2266_v25 = _rhs382;
          }
        }
      }
      let _2269_v28;
      _2269_v28 = new _dafny.CodePoint('q'.codePointAt(0));
      _2269_v28 = _2269_v28;
      (globalState).f2 = _dafny.Seq.UnicodeFromString("henupki");
      let _2270_v29;
      _2270_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      let _2271_v30;
      _2271_v30 = _module.D6.create_DC19(new BigNumber((_2270_v29).length), new BigNumber((_dafny.Seq.of((_this).f6, (_this).f6, (_this).f6)).length), (_this).f6);
      let _2272_v31;
      _2272_v31 = _module.D6.create_DC20(_2271_v30);
      let _pat_let_tv57 = _2271_v30;
      let _2273_v32;
      let _nw339 = Array((new BigNumber(29)).toNumber());
      _nw339[(_dafny.ZERO).toNumber()] = _module.D6.create_DC20(_2271_v30);
      _nw339[(_dafny.ONE).toNumber()] = _module.D6.create_DC20(_2271_v30);
      _nw339[(new BigNumber(2)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(3)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(4)).toNumber()] = function (_pat_let74_0) {
        return function (_2274_dt__update__tmp_h0) {
          return function (_pat_let75_0) {
            return function (_2275_dt__update_hcf42_h0) {
              return _module.D6.create_DC20(_2275_dt__update_hcf42_h0);
            }(_pat_let75_0);
          }(_pat_let_tv57);
        }(_pat_let74_0);
      }(_2272_v31);
      _nw339[(new BigNumber(5)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(6)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(7)).toNumber()] = _module.D6.create_DC20(_module.D6.create_DC20(_2271_v30));
      _nw339[(new BigNumber(8)).toNumber()] = _module.D6.create_DC20(_2271_v30);
      _nw339[(new BigNumber(9)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(10)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(11)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(12)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(13)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(14)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(15)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(16)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(17)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(18)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(19)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(20)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(21)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(22)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(23)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(24)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(25)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(26)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(27)).toNumber()] = _2272_v31;
      _nw339[(new BigNumber(28)).toNumber()] = _2272_v31;
      _2273_v32 = _nw339;
      let _2276_v33;
      _2276_v33 = _2273_v32;
      let _source29 = _2276_v33;
      let _2277___mcc_h0 = _source29;
      let _2278_cf72 = _2277___mcc_h0;
      let _2279_v34;
      _2279_v34 = _module.D1.create_DC4(_2240_v3, (_this).f6, p0, _2238_v1, _module.__default.fm2((_this).f6, globalState));
      let _source30 = _2279_v34;
      if (_source30.is_DC4) {
        let _2280___mcc_h1 = (_source30).cf7;
        let _2281___mcc_h2 = (_source30).cf8;
        let _2282___mcc_h3 = (_source30).cf9;
        let _2283___mcc_h4 = (_source30).cf10;
        let _2284___mcc_h5 = (_source30).cf11;
        let _2285_cf11 = _2284___mcc_h5;
        let _2286_cf10 = _2283___mcc_h4;
        let _2287_cf9 = _2282___mcc_h3;
        let _2288_cf8 = _2281___mcc_h2;
        let _2289_cf7 = _2280___mcc_h1;
        (globalState).f5 = _2288_cf8;
        let _2290_v35;
        _2290_v35 = _dafny.MultiSet.fromElements(_2285_cf11);
        let _2291_v36;
        _2291_v36 = _dafny.Seq.of(_2290_v35);
        let _2292_v37;
        _2292_v37 = _dafny.MultiSet.fromElements(_2238_v1);
        let _2293_v38;
        _2293_v38 = _dafny.Seq.of(_2290_v35, _2290_v35, _2290_v35, _2290_v35, (_2291_v36)[_module.__default.safeIndex(new BigNumber((_2292_v37).cardinality()), new BigNumber((_2291_v36).length))]);
        let _2294_v39;
        let _nw340 = Array((new BigNumber(15)).toNumber());
        _nw340[(_dafny.ZERO).toNumber()] = _module.__default.fm2(_2285_cf11, globalState);
        _nw340[(_dafny.ONE).toNumber()] = new BigNumber(-393);
        _nw340[(new BigNumber(2)).toNumber()] = _2285_cf11;
        _nw340[(new BigNumber(3)).toNumber()] = new BigNumber(-989);
        _nw340[(new BigNumber(4)).toNumber()] = ((_this).f6).plus(new BigNumber(999));
        _nw340[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(_2288_cf8, _2285_cf11);
        _nw340[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(((((_2270_v29).contains(new BigNumber((_2287_cf9).length))) ? ((_2270_v29).get(new BigNumber((_2287_cf9).length))) : (new BigNumber((_2290_v35).cardinality())))).multipliedBy(new BigNumber((_2291_v36).length)));
        _nw340[(new BigNumber(7)).toNumber()] = (new BigNumber((p0).length)).multipliedBy(_2285_cf11);
        _nw340[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2288_cf8,(_2239_v2)[_module.__default.safeIndex(_2285_cf11, new BigNumber((_2239_v2).length))])).length);
        _nw340[(new BigNumber(9)).toNumber()] = (_this).f6;
        _nw340[(new BigNumber(10)).toNumber()] = (((_2290_v35).contains((_this).f6)) ? ((_2290_v35).get((_this).f6)) : (_2285_cf11));
        _nw340[(new BigNumber(11)).toNumber()] = (_2288_cf8).minus((_dafny.ZERO).minus(_2285_cf11));
        _nw340[(new BigNumber(12)).toNumber()] = new BigNumber(664);
        _nw340[(new BigNumber(13)).toNumber()] = new BigNumber((_2292_v37).cardinality());
        _nw340[(new BigNumber(14)).toNumber()] = _2288_cf8;
        _2294_v39 = _nw340;
        let _index313 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_2294_v39).length));
        (_2294_v39)[_index313] = _2288_cf8;
        let _2295_v40;
        _2295_v40 = _module.D2.create_DC6(_2285_cf11, _2294_v39);
        let _2296_v41;
        _2296_v41 = _dafny.MultiSet.fromElements((_2295_v40).dtor_cf14);
        let _index314 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_2294_v39).length));
        (_2294_v39)[_index314] = new BigNumber(((_2296_v41).Union(_dafny.MultiSet.fromElements(_2294_v39, _2294_v39))).cardinality());
        (globalState).f0 = _2286_cf10;
        let _2297_v42;
        _2297_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2285_cf11,p0);
        let _2298_v43;
        let _nw341 = new _module.C7();
        _nw341.__ctor(_2294_v39, _2294_v39, _module.__default.fm2((_this).f6, globalState), _2242_v5, _2297_v42);
        _2298_v43 = _nw341;
        let _rhs383 = _2298_v43;
        let _rhs384 = _2285_cf11;
        _2298_v43 = _rhs383;
        _2288_cf8 = _rhs384;
      } else {
        let _2299___mcc_h6 = (_source30).cf6;
        let _2300_cf6 = _2299___mcc_h6;
        (globalState).f0 = !(!_dafny.Seq.contains(p0, _2269_v28));
        _2240_v3 = _2240_v3;
        (globalState).f0 = (new BigNumber(322)).isLessThan((_this).f6);
        let _2301_v44;
        let _nw342 = new _module.C1();
        _nw342.__ctor((_this).f6);
        _2301_v44 = _nw342;
      }
      let _index315 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2242_v5).length));
      (_2242_v5)[_index315] = _2238_v1;
      let _index316 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2242_v5).length));
      (_2242_v5)[_index316] = (new BigNumber((p0).length)).isLessThan((_this).f6);
      let _2302_v45;
      _2302_v45 = _dafny.Set.fromElements((_this).f6);
      let _2303_v46;
      _2303_v46 = _dafny.MultiSet.fromElements(new BigNumber((_2302_v45).length), new BigNumber(951));
      let _2304_v47;
      _2304_v47 = _dafny.MultiSet.fromElements(_2303_v46);
      let _2305_v48;
      let _nw343 = Array((new BigNumber(23)).toNumber()).fill(false);
      _2305_v48 = _nw343;
      let _2306_v50;
      let _nw344 = new _module.C2();
      _nw344.__ctor((_this).f6, _2304_v47, _2305_v48, function () {
        let _coll61 = new _dafny.Map();
        for (const _compr_61 of _dafny.IntegerRange(new BigNumber(993), new BigNumber(-151))) {
          let _2307_v49 = _compr_61;
          if (((new BigNumber(993)).isLessThanOrEqualTo(_2307_v49)) && ((_2307_v49).isLessThan(new BigNumber(-151)))) {
            _coll61.push([(_2307_v49).minus((_this).f6),p0]);
          }
        }
        return _coll61;
      }());
      _2306_v50 = _nw344;
      let _2308_v51;
      _2308_v51 = _dafny.MultiSet.fromElements(_2306_v50);
      let _2309_v52;
      _2309_v52 = _dafny.Map.Empty.slice().updateUnsafe((_2306_v50).f6,p0);
      let _2310_v53;
      _2310_v53 = _2309_v52;
      let _2311_v54;
      _2311_v54 = _dafny.MultiSet.fromElements(_2310_v53);
      let _2312_v55;
      _2312_v55 = _dafny.MultiSet.fromElements(!(_2238_v1));
      let _2313_v56;
      let _nw345 = Array((new BigNumber(17)).toNumber());
      _nw345[(_dafny.ZERO).toNumber()] = (_this).f6;
      _nw345[(_dafny.ONE).toNumber()] = ((_this).f6).plus((_dafny.ZERO).minus((_this).f6));
      _nw345[(new BigNumber(2)).toNumber()] = (_this).f6;
      _nw345[(new BigNumber(3)).toNumber()] = (_this).f6;
      _nw345[(new BigNumber(4)).toNumber()] = ((_this).f6).plus((_this).f6);
      _nw345[(new BigNumber(5)).toNumber()] = (_this).f6;
      _nw345[(new BigNumber(6)).toNumber()] = (_this).f6;
      _nw345[(new BigNumber(7)).toNumber()] = new BigNumber((_2308_v51).cardinality());
      _nw345[(new BigNumber(8)).toNumber()] = (_2306_v50).f6;
      _nw345[(new BigNumber(9)).toNumber()] = ((_2306_v50).f6).minus(_module.__default.fm2((_this).f6, globalState));
      _nw345[(new BigNumber(10)).toNumber()] = (_this).f6;
      _nw345[(new BigNumber(11)).toNumber()] = ((_2306_v50).f6).minus((_2306_v50).f6);
      _nw345[(new BigNumber(12)).toNumber()] = ((_2306_v50).f6).minus((_this).f6);
      _nw345[(new BigNumber(13)).toNumber()] = (_this).f6;
      _nw345[(new BigNumber(14)).toNumber()] = (_this).f6;
      _nw345[(new BigNumber(15)).toNumber()] = (_2306_v50).f6;
      _nw345[(new BigNumber(16)).toNumber()] = new BigNumber(((((_2311_v54).update(_2310_v53, _module.__default.abs(new BigNumber((_2312_v55).cardinality())))).update(_2310_v53, _module.__default.abs((_this).f6))).Intersect(_2311_v54)).cardinality());
      _2313_v56 = _nw345;
      let _index317 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length));
      (_2313_v56)[_index317] = (_2306_v50).f6;
      let _index318 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length));
      (_2313_v56)[_index318] = _module.__default.fm2((_2306_v50).f6, globalState);
      if (_module.__default.fm3(globalState)) {
        let _2314_v57;
        _2314_v57 = _dafny.Seq.of(_2306_v50);
        let _2315_v58;
        _2315_v58 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_2239_v2, _module.__default.safeIndex(new BigNumber((_2312_v55).cardinality()), new BigNumber((_2239_v2).length)), false),new BigNumber((_2314_v57).length));
        _2315_v58 = (_2315_v58).update(_2239_v2, ((_2313_v56)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length))]).minus(new BigNumber(804)));
        let _2316_v59;
        let _out31;
        _out31 = (_this).m2(globalState);
        _2316_v59 = _out31;
        let _2317_v60;
        _2317_v60 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f6), _2241_v4);
        let _2318_v61;
        _2318_v61 = _dafny.Map.Empty.slice().updateUnsafe((_2313_v56)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length))],_2317_v60);
        _2237_v0 = _dafny.Seq.of(((_2306_v50).f6).plus(new BigNumber((function () {
          let _coll62 = new _dafny.Set();
          for (const _compr_62 of ((((_2318_v61).contains((_2306_v50).f6)) ? ((_2318_v61).get((_2306_v50).f6)) : (_2317_v60))).Elements) {
            let _2319_v62 = _compr_62;
            if (((((_2318_v61).contains((_2306_v50).f6)) ? ((_2318_v61).get((_2306_v50).f6)) : (_2317_v60))).contains(_2319_v62)) {
              _coll62.add(_2319_v62);
            }
          }
          return _coll62;
        }()).length)));
        let _2320_v63;
        let _nw346 = Array((new BigNumber(18)).toNumber()).fill(_module.D11.Default());
        _2320_v63 = _nw346;
        let _index319 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2242_v5).length));
        let _index320 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length));
        let _rhs385 = _2320_v63;
        let _rhs386 = !(_module.__default.fm3(globalState));
        let _rhs387 = _2238_v1;
        let _rhs388 = _module.__default.safeDivisionInt((_2306_v50).f6, new BigNumber((p0).length));
        let _rhs389 = _2313_v56;
        let _lhs307 = _2242_v5;
        let _lhs308 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2242_v5).length));
        let _lhs309 = _2313_v56;
        let _lhs310 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length));
        _2320_v63 = _rhs385;
        _2316_v59 = _rhs386;
        _lhs307[_lhs308] = _rhs387;
        _lhs309[_lhs310] = _rhs388;
        _2313_v56 = _rhs389;
        (globalState).f5 = (_dafny.ZERO).minus((((_2306_v50).f6).multipliedBy((_2306_v50).f6)).multipliedBy((_2237_v0)[_module.__default.safeIndex((_2313_v56)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length))], new BigNumber((_2237_v0).length))]));
      } else {
        let _index321 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length));
        (_2313_v56)[_index321] = (_this).f6;
        let _2321_v64;
        let _nw347 = new _module.C9();
        _nw347.__ctor(((_2306_v50).f6).plus((_2313_v56)[_module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length))]));
        _2321_v64 = _nw347;
        let _index322 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length));
        let _rhs390 = (_2239_v2)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_2306_v50).f6, (_dafny.ZERO).minus((_2306_v50).f6))).length), new BigNumber((_2239_v2).length))];
        let _rhs391 = _2321_v64;
        let _rhs392 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_2306_v50).f6, (_2306_v50).f6));
        let _rhs393 = (_dafny.ZERO).minus((_this).f6);
        let _lhs311 = _2313_v56;
        let _lhs312 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_2313_v56).length));
        let _lhs313 = globalState;
        _2238_v1 = _rhs390;
        _2321_v64 = _rhs391;
        _lhs311[_lhs312] = _rhs392;
        _lhs313.f5 = _rhs393;
        let _2322_v65;
        _2322_v65 = _dafny.Map.Empty.slice().updateUnsafe((_2306_v50).f6,(_2242_v5)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_2242_v5).length))]);
        (globalState).f2 = _dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_2322_v65).length), new BigNumber((p0).length)), _2269_v28);
        (globalState).f5 = (_2306_v50).f6;
        let _2323_v66;
        let _out32;
        _out32 = (_this).m2(globalState);
        _2323_v66 = _out32;
      }
      r0 = ((_this).f6).isEqualTo((_this).f6);
      return r0;
    }
    m2(globalState) {
      let _this = this;
      let r0 = false;
      let _2324_v0;
      _2324_v0 = false;
      (globalState).f2 = _module.__default.fm18(_2324_v0, globalState);
      let _2325_v1;
      let _nw348 = Array((new BigNumber(5)).toNumber());
      _nw348[(_dafny.ZERO).toNumber()] = (false) || (_2324_v0);
      _nw348[(_dafny.ONE).toNumber()] = _2324_v0;
      _nw348[(new BigNumber(2)).toNumber()] = _2324_v0;
      _nw348[(new BigNumber(3)).toNumber()] = (_2324_v0) || (_2324_v0);
      _nw348[(new BigNumber(4)).toNumber()] = _2324_v0;
      _2325_v1 = _nw348;
      let _index323 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length));
      (_2325_v1)[_index323] = _2324_v0;
      let _2326_v2;
      _2326_v2 = _dafny.Seq.of(_2324_v0);
      let _2327_v3;
      _2327_v3 = _dafny.MultiSet.fromElements(_2326_v2, _2326_v2);
      let _2328_v4;
      _2328_v4 = _dafny.Seq.UnicodeFromString("rpavxdbs");
      let _2329_v5;
      _2329_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2328_v4);
      let _2330_v6;
      _2330_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      let _2331_v7;
      _2331_v7 = _dafny.MultiSet.fromElements(new BigNumber(264));
      let _2332_v8;
      _2332_v8 = _dafny.MultiSet.fromElements(new BigNumber((_2331_v7).cardinality()));
      let _2333_v9;
      _2333_v9 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_this).f6), _2332_v8, _2332_v8, _2332_v8, _2331_v7);
      let _2334_v10;
      let _nw349 = new _module.C4();
      _nw349.__ctor(new BigNumber((_2328_v4).length), (_this).f6, _2325_v1, _2329_v5, (((_2330_v6).contains(new BigNumber((_dafny.Set.fromElements((_this).f6)).length))) ? ((_2330_v6).get(new BigNumber((_dafny.Set.fromElements((_this).f6)).length))) : ((_this).f6)), _2333_v9);
      _2334_v10 = _nw349;
      let _2335_v11;
      _2335_v11 = _module.D18.create_DC44(_2324_v0, _2324_v0, _2327_v3, _2334_v10);
      let _index324 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length));
      let _rhs394 = (_2335_v11).dtor_cf80;
      let _rhs395 = (_2334_v10).f17;
      let _rhs396 = _2324_v0;
      let _rhs397 = _2324_v0;
      let _rhs398 = _2325_v1;
      let _lhs314 = globalState;
      let _lhs315 = _2325_v1;
      let _lhs316 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length));
      let _lhs317 = globalState;
      r0 = _rhs394;
      _lhs314.f5 = _rhs395;
      _lhs315[_lhs316] = _rhs396;
      _lhs317.f0 = _rhs397;
      _2325_v1 = _rhs398;
      if (false) {
        (globalState).f5 = (_2334_v10).fm21((_2334_v10).f17, globalState);
        let _2336_v12;
        _2336_v12 = new _dafny.CodePoint('f'.codePointAt(0));
        let _2337_v13;
        _2337_v13 = _dafny.Set.fromElements((_2334_v10).f17, (_dafny.ZERO).minus((_2334_v10).f18));
        let _2338_v14;
        _2338_v14 = _module.D18.create_DC43(new BigNumber(266));
        let _2339_v15;
        _2339_v15 = _dafny.Seq.of((_2334_v10).f17, new BigNumber((_dafny.Seq.UnicodeFromString("xsvnce")).length));
        let _2340_v17;
        _2340_v17 = _dafny.Set.fromElements(((false) ? (_2336_v12) : ((_2328_v4)[_module.__default.safeIndex(new BigNumber((_2337_v13).length), new BigNumber((_2328_v4).length))])), _2336_v12, _module.__default.fm0((((_2330_v6).contains((_2338_v14).dtor_cf78)) ? ((_2330_v6).get((_2338_v14).dtor_cf78)) : (new BigNumber((_2339_v15).length))), function () {
          let _coll63 = new _dafny.Set();
          for (const _compr_63 of (_2339_v15).Elements) {
            let _2341_v16 = _compr_63;
            if (_dafny.Seq.contains(_2339_v15, _2341_v16)) {
              _coll63.add((_2341_v16).multipliedBy(new BigNumber(2)));
            }
          }
          return _coll63;
        }(), (_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))], globalState), _2336_v12);
        let _2342_v18;
        let _nw350 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _2342_v18 = _nw350;
        let _2343_v19;
        _2343_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2342_v18,_2340_v17);
        let _2344_v20;
        _2344_v20 = _dafny.Seq.of(_2326_v2, _2326_v2, _2326_v2, _2326_v2);
        let _2345_v21;
        _2345_v21 = _module.D9.create_DC25(new BigNumber(((_2344_v20)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("q")).length), new BigNumber((_2344_v20).length))]).length), (_2334_v10).f18, (_2334_v10).f18, (_2328_v4)[_module.__default.safeIndex((_2334_v10).f18, new BigNumber((_2328_v4).length))]);
        let _2346_v22;
        _2346_v22 = _dafny.Seq.of(_module.D9.create_DC25(new BigNumber((_dafny.Seq.UnicodeFromString("qqakq")).length), (_2334_v10).f18, new BigNumber((_module.__default.fm9(globalState)).length), _2336_v12), _2345_v21);
        _2340_v17 = (((_2343_v19).contains(_2342_v18)) ? ((_2343_v19).get(_2342_v18)) : (_module.__default.fm54((_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))], _module.__default.fm0((_2334_v10).f17, _2337_v13, (_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))], globalState), _2346_v22, globalState)));
        _2336_v12 = _2336_v12;
        let _2347_v23;
        let _nw351 = Array((new BigNumber(7)).toNumber()).fill([]);
        _2347_v23 = _nw351;
        let _2348_v24;
        let _nw352 = Array((new BigNumber(22)).toNumber());
        _2348_v24 = _nw352;
        let _index325 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_2347_v23).length));
        (_2347_v23)[_index325] = _2348_v24;
        let _2349_v25;
        let _nw353 = new _module.C9();
        _nw353.__ctor((_2334_v10).f18);
        _2349_v25 = _nw353;
        let _2350_v26;
        _2350_v26 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_2349_v25).f6),_2349_v25);
        let _2351_v27;
        _2351_v27 = _2349_v25;
        let _index326 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_2347_v23).length));
        let _nw354 = Array((new BigNumber(20)).toNumber());
        _nw354[(_dafny.ZERO).toNumber()] = _2349_v25;
        _nw354[(_dafny.ONE).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(2)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(3)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(4)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(5)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(6)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(7)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(8)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(9)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(10)).toNumber()] = (((_2350_v26).contains((_2334_v10).f17)) ? ((_2350_v26).get((_2334_v10).f17)) : (_2349_v25));
        _nw354[(new BigNumber(11)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(12)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(13)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(14)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(15)).toNumber()] = (_2351_v27);
        _nw354[(new BigNumber(16)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(17)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(18)).toNumber()] = _2349_v25;
        _nw354[(new BigNumber(19)).toNumber()] = (_2349_v25);
        (_2347_v23)[_index326] = _nw354;
        let _2352_v28;
        _2352_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2324_v0,(_2334_v10).f17);
        let _2353_v29;
        _2353_v29 = _dafny.Map.Empty.slice().updateUnsafe(!((_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))]),new BigNumber(((_2352_v28).Merge(_2352_v28)).length));
        _2352_v28 = (_2352_v28).update(_dafny.Seq.IsProperPrefixOf(_2328_v4, _dafny.Seq.UnicodeFromString("ywxdsc")), new BigNumber((_dafny.Seq.Concat(_2339_v15, _2339_v15)).length));
      } else {
        let _2354_v30;
        _2354_v30 = _module.D0.create_DC0(_dafny.Seq.of(_2324_v0, false, (_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))], (_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))]), (_2334_v10).f18);
        let _2355_v31;
        let _init58 = function (_2356_i0) {
          return _module.__default.safeDivisionInt(_2356_i0, (_this).f6);
        };
        let _nw355 = Array((new BigNumber(19)).toNumber());
        for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw355.length); _i0_58++) {
          _nw355[_i0_58] = _init58(new BigNumber(_i0_58));
        }
        _2355_v31 = _nw355;
        let _2357_v32;
        _2357_v32 = _dafny.MultiSet.fromElements(_2355_v31, _2355_v31);
        let _2358_v33;
        _2358_v33 = _dafny.Map.Empty.slice().updateUnsafe(!((_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))]),_2355_v31);
        let _2359_v34;
        _2359_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2354_v30,(_2357_v32).update((((_2358_v33).contains((_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))])) ? ((_2358_v33).get((_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))])) : (_2355_v31)), _module.__default.abs((_this).f6)));
        _2359_v34 = (_2359_v34).update(_2354_v30, (_2357_v32).Difference(_2357_v32));
        let _2360_v35;
        _2360_v35 = _dafny.Seq.of((_2334_v10).f17, (_2334_v10).f18, new BigNumber(-387), (_2334_v10).f18);
        let _2361_v36;
        _2361_v36 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_2360_v35, _module.__default.safeIndex(new BigNumber(-334), new BigNumber((_2360_v35).length)), (_2334_v10).f18),_dafny.Seq.of(_module.__default.fm55(globalState)));
        _2361_v36 = _2361_v36;
        (globalState).f0 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_2360_v35, _dafny.Seq.Create(_module.__default.abs(new BigNumber(249)), ((_2362_v10) => function (_2363_i1) {
          return (_2362_v10).f17;
        })(_2334_v10))), _2360_v35);
        let _2364_v37;
        _2364_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2324_v0,new BigNumber(788));
        let _2365_v38;
        _2365_v38 = _dafny.Seq.of(_2364_v37);
        (globalState).f5 = (new BigNumber((((_2365_v38)[_module.__default.safeIndex(new BigNumber((_2360_v35).length), new BigNumber((_2365_v38).length))]).Merge(_2364_v37)).length)).multipliedBy(((_dafny.ZERO).minus((_2334_v10).f17)).plus((_2334_v10).f18));
        let _2366_v39;
        let _nw356 = new _module.C4();
        _nw356.__ctor((_2334_v10).f18, new BigNumber(82), _2325_v1, (_2329_v5).update((_2334_v10).f18, _2328_v4), (_this).f6, _2333_v9);
        _2366_v39 = _nw356;
        let _2367_v40;
        _2367_v40 = _dafny.Map.Empty.slice().updateUnsafe((_2334_v10).f17,_2366_v39);
        let _2368_v41;
        _2368_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2367_v40,((_2334_v10).f18).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("m")).length)));
        _2368_v41 = (_2368_v41).update(_2367_v40, _2324_v0);
      }
      let _2369_v42;
      let _nw357 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _2369_v42 = _nw357;
      let _2370_v43;
      _2370_v43 = new _dafny.CodePoint('t'.codePointAt(0));
      let _index327 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_2369_v42).length));
      (_2369_v42)[_index327] = _2370_v43;
      let _2371_v44;
      _2371_v44 = _module.D14.create_DC37((_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))], _2328_v4);
      let _pat_let_tv58 = _2334_v10;
      let _pat_let_tv59 = _2324_v0;
      let _pat_let_tv60 = _2334_v10;
      let _pat_let_tv61 = _2334_v10;
      let _pat_let_tv62 = _2370_v43;
      let _pat_let_tv63 = _2325_v1;
      let _pat_let_tv64 = _2325_v1;
      let _pat_let_tv65 = _2325_v1;
      let _pat_let_tv66 = _2325_v1;
      let _pat_let_tv67 = _2370_v43;
      let _pat_let_tv68 = _2370_v43;
      let _index328 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_2369_v42).length));
      (_2369_v42)[_index328] = function (_source31) {
        if (_source31.is_DC37) {
          let _2372___mcc_h0 = (_source31).cf70;
          let _2373___mcc_h1 = (_source31).cf71;
          let _2374_cf71 = _2373___mcc_h1;
          let _2375_cf70 = _2372___mcc_h0;
          return (_module.D9.create_DC25((_pat_let_tv58).f17, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv59,new BigNumber((_dafny.Seq.of((_pat_let_tv60).f17)).length))).length), (_pat_let_tv61).f18)).length), new BigNumber(-684), _pat_let_tv62)).dtor_cf49;
        } else {
          let _2376___mcc_h2 = (_source31).cf69;
          let _2377_cf69 = _2376___mcc_h2;
          if ((_pat_let_tv64)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_pat_let_tv63).length))]) {
            return (_module.D0.create_DC1((_pat_let_tv66)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_pat_let_tv65).length))], (_this).f6, _pat_let_tv67)).dtor_cf4;
          } else {
            return _pat_let_tv68;
          }
        }
      }(_2371_v44);
      let _2378_v45;
      _2378_v45 = _dafny.Seq.of(new BigNumber(701));
      let _rhs399 = (_2378_v45)[_module.__default.safeIndex((_2334_v10).f18, new BigNumber((_2378_v45).length))];
      let _rhs400 = (_2378_v45)[_module.__default.safeIndex(((_module.__default.fm56(_2328_v4, _2324_v0, _2328_v4, globalState)).dtor_cf59).minus((_this).f6), new BigNumber((_2378_v45).length))];
      let _rhs401 = (_2334_v10).f17;
      let _lhs318 = globalState;
      let _lhs319 = globalState;
      let _lhs320 = globalState;
      _lhs318.f5 = _rhs399;
      _lhs319.f5 = _rhs400;
      _lhs320.f5 = _rhs401;
      let _2379_v46;
      _2379_v46 = _dafny.Map.Empty.slice().updateUnsafe(false,_2324_v0);
      let _2380_v47;
      _2380_v47 = _dafny.Set.fromElements(_2379_v46, (_2379_v46).update((_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))], true), _2379_v46);
      (globalState).f5 = new BigNumber((_2380_v47).length);
      r0 = (_2325_v1)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((_2325_v1).length))];
      return r0;
    }
  };

  $module.C11 = class C11 {
    constructor () {
      this._tname = "_module.C11";
      this._f6 = _dafny.ZERO;
      this.f8 = false;
      this._f7 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f7, f8, f6) {
      let _this = this;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f6 = f6;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      let _source32 = ((_this.f8) ? (_module.D0.create_DC2(_module.D0.create_DC1(_this.f8, (_this).f6, new _dafny.CodePoint('l'.codePointAt(0))))) : (_module.D0.create_DC2(_module.D0.create_DC0(_dafny.Seq.of(_this.f8), (_this).f6))));
      if (_source32.is_DC0) {
        let _2381___mcc_h0 = (_source32).cf0;
        let _2382___mcc_h1 = (_source32).cf1;
        let _2383_cf1 = _2382___mcc_h1;
        let _2384_cf0 = _2381___mcc_h0;
        return _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)))));
      } else if (_source32.is_DC1) {
        let _2385___mcc_h2 = (_source32).cf2;
        let _2386___mcc_h3 = (_source32).cf3;
        let _2387___mcc_h4 = (_source32).cf4;
        let _2388_cf4 = _2387___mcc_h4;
        let _2389_cf3 = _2386___mcc_h3;
        let _2390_cf2 = _2385___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(_2388_cf4)), _dafny.Seq.of(_dafny.MultiSet.fromElements(_2388_cf4, _2388_cf4)));
      } else {
        let _2391___mcc_h5 = (_source32).cf5;
        let _2392_cf5 = _2391___mcc_h5;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(538)), function (_2393_i0) {
          return _dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))));
        }), _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0))), (_module.D1.create_DC3(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0))))).dtor_cf6, _dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))));
      }
    };
    fm6(p0, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.of((_this).f6))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_this.f8, _this.f8)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_this.f8)).length))).length), (_this).f6, (_this).f6)));
    };
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f8;
    };
    fm8(globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_this).f6,((_this.f8) ? (_module.D0.create_DC0(_dafny.Seq.of(false), (_this).f6)) : (_module.D0.create_DC0(_dafny.Seq.of(_this.f8), (_dafny.ZERO).minus((_this).f6)))));
    };
    m0(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _2394_v0;
      _2394_v0 = _dafny.Seq.of((_this).f6);
      let _2395_v1;
      _2395_v1 = _dafny.Set.fromElements(_this.f8, _this.f8);
      let _2396_v2;
      _2396_v2 = _dafny.Seq.of(p1, _this.f8, false, _this.f8);
      let _2397_v3;
      _2397_v3 = _dafny.Seq.UnicodeFromString("tbon");
      let _2398_v4;
      _2398_v4 = _module.D0.create_DC1(_this.f8, (_this).f6, new _dafny.CodePoint('p'.codePointAt(0)));
      let _2399_v5;
      _2399_v5 = _dafny.MultiSet.fromElements((_this).f6);
      let _2400_v6;
      _2400_v6 = _dafny.MultiSet.fromElements(_2399_v5);
      let _2401_v7;
      _2401_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.UnicodeFromString("ogtymm"));
      let _2402_v8;
      let _nw358 = new _module.C2();
      _nw358.__ctor(new BigNumber((_2397_v3).length), _2400_v6, p2, _2401_v7);
      _2402_v8 = _nw358;
      let _2403_v9;
      _2403_v9 = _dafny.Seq.of(_2402_v8);
      let _2404_v10;
      _2404_v10 = _module.D9.create_DC24(_2400_v6);
      let _2405_v11;
      _2405_v11 = _dafny.Seq.of(_2404_v10, _2404_v10, _2404_v10);
      let _2406_v12;
      let _nw359 = Array((new BigNumber(20)).toNumber());
      _nw359[(_dafny.ZERO).toNumber()] = !(p1) || (_this.f8);
      _nw359[(_dafny.ONE).toNumber()] = _this.f8;
      _nw359[(new BigNumber(2)).toNumber()] = ((_this).f6).isLessThanOrEqualTo(new BigNumber(365));
      _nw359[(new BigNumber(3)).toNumber()] = _this.f8;
      _nw359[(new BigNumber(4)).toNumber()] = true;
      _nw359[(new BigNumber(5)).toNumber()] = _dafny.areEqual(_2394_v0, _2394_v0);
      _nw359[(new BigNumber(6)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_2395_v1).length))).isEqualTo((_dafny.ZERO).minus((_this).f6));
      _nw359[(new BigNumber(7)).toNumber()] = !_dafny.Seq.contains(_2396_v2, p1);
      _nw359[(new BigNumber(8)).toNumber()] = p1;
      _nw359[(new BigNumber(9)).toNumber()] = false;
      _nw359[(new BigNumber(10)).toNumber()] = !((new BigNumber((_dafny.Seq.update(_2397_v3, _module.__default.safeIndex((_2398_v4).dtor_cf3, new BigNumber((_2397_v3).length)), p0)).length)).isLessThanOrEqualTo((_this).f6));
      _nw359[(new BigNumber(11)).toNumber()] = p1;
      _nw359[(new BigNumber(12)).toNumber()] = _dafny.Seq.IsPrefixOf(_2403_v9, _2403_v9);
      _nw359[(new BigNumber(13)).toNumber()] = p1;
      _nw359[(new BigNumber(14)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_2405_v11, _2405_v11);
      _nw359[(new BigNumber(15)).toNumber()] = (_2396_v2)[_module.__default.safeIndex((_this).f6, new BigNumber((_2396_v2).length))];
      _nw359[(new BigNumber(16)).toNumber()] = ((_2402_v8).f6).isLessThan(new BigNumber(746));
      _nw359[(new BigNumber(17)).toNumber()] = !(_this.f8);
      _nw359[(new BigNumber(18)).toNumber()] = p1;
      _nw359[(new BigNumber(19)).toNumber()] = ((_2402_v8).f6).isLessThanOrEqualTo(new BigNumber((_2397_v3).length));
      _2406_v12 = _nw359;
      _2406_v12 = _2406_v12;
      (globalState).f5 = (_2402_v8).f6;
      (globalState).f5 = (_2402_v8).f6;
      let _2407_v13;
      let _nw360 = new _module.C4();
      _nw360.__ctor((_2402_v8).f6, (_this).f6, p2, _2401_v7, (_this).f6, _2400_v6);
      _2407_v13 = _nw360;
      let _2408_v14;
      _2408_v14 = _dafny.MultiSet.fromElements(_2407_v13, _2407_v13);
      let _hi14 = (_dafny.ZERO).minus(((p1) ? (new BigNumber((_2395_v1).length)) : ((_2402_v8).f6)));
      for (let _2409_i0 = (((_2408_v14).contains(_2407_v13)) ? ((_2408_v14).get(_2407_v13)) : ((_2407_v13).f17)); _2409_i0.isLessThan(_hi14); _2409_i0 = _2409_i0.plus(_dafny.ONE)) {
        let _2410_v15;
        _2410_v15 = _dafny.Map.Empty.slice().updateUnsafe(((_2407_v13).f18).plus((_2407_v13).f17),(_dafny.ZERO).minus((_2407_v13).f18));
        _2410_v15 = (function () {
          let _coll64 = new _dafny.Map();
          for (const _compr_64 of _dafny.IntegerRange(new BigNumber(-162), new BigNumber(475))) {
            let _2411_v16 = _compr_64;
            if (((new BigNumber(-162)).isLessThanOrEqualTo(_2411_v16)) && ((_2411_v16).isLessThan(new BigNumber(475)))) {
              _coll64.push([(_2411_v16).plus(new BigNumber((_dafny.MultiSet.fromElements((_2407_v13).f18, (_2402_v8).f6)).cardinality())),(_this).f6]);
            }
          }
          return _coll64;
        }()).Merge(_2410_v15);
        if (_this.f8) {
          let _2412_v17;
          let _init59 = ((_2413_v3) => function (_2414_i1) {
            return (_2414_i1).plus(new BigNumber((_2413_v3).length));
          })(_2397_v3);
          let _nw361 = Array((new BigNumber(29)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw361.length); _i0_59++) {
            _nw361[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _2412_v17 = _nw361;
          let _index329 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_2412_v17).length));
          (_2412_v17)[_index329] = (new BigNumber(291)).multipliedBy((_2407_v13).f17);
          let _index330 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_2412_v17).length));
          let _rhs402 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2397_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(750)), function (_2415_i2) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })), _module.__default.safeIndex((_2402_v8).f6, new BigNumber((_dafny.Seq.Concat(_2397_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(750)), function (_2416_i2) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }))).length)), p0)).length);
          let _rhs403 = _this.f8;
          let _rhs404 = _2397_v3;
          let _rhs405 = p1;
          let _lhs321 = _2412_v17;
          let _lhs322 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_2412_v17).length));
          let _lhs323 = globalState;
          let _lhs324 = globalState;
          _lhs321[_lhs322] = _rhs402;
          _lhs323.f0 = _rhs403;
          _2397_v3 = _rhs404;
          _lhs324.f0 = _rhs405;
          let _2417_v18;
          _2417_v18 = _module.D14.create_DC37(_this.f8, _2397_v3);
          (globalState).f0 = (_2417_v18).dtor_cf70;
          (globalState).f5 = (_2394_v0)[_module.__default.safeIndex((_2407_v13).f18, new BigNumber((_2394_v0).length))];
          let _2418_v19;
          _2418_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2402_v8,_2406_v12);
          _2406_v12 = (((_2418_v19).contains(_2402_v8)) ? ((_2418_v19).get(_2402_v8)) : (p2));
          let _2419_v20;
          let _2420_v21;
          let _2421_v22;
          let _2422_v23;
          let _out33;
          let _out34;
          let _out35;
          let _out36;
          let _outcollector11 = (_2407_v13).m7((_2399_v5).Union(_2399_v5), (_2407_v13).f18, globalState);
          _out33 = _outcollector11[0];
          _out34 = _outcollector11[1];
          _out35 = _outcollector11[2];
          _out36 = _outcollector11[3];
          _2419_v20 = _out33;
          _2420_v21 = _out34;
          _2421_v22 = _out35;
          _2422_v23 = _out36;
        } else {
          let _2423_v24;
          let _nw362 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _2423_v24 = _nw362;
          let _index331 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_2423_v24).length));
          (_2423_v24)[_index331] = (new BigNumber((_2396_v2).length)).multipliedBy((_2407_v13).f17);
          let _index332 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_2423_v24).length));
          let _rhs406 = (_2407_v13).f18;
          let _rhs407 = (_2407_v13).f17;
          let _lhs325 = globalState;
          let _lhs326 = _2423_v24;
          let _lhs327 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_2423_v24).length));
          _lhs325.f5 = _rhs406;
          _lhs326[_lhs327] = _rhs407;
          let _rhs408 = ((_2399_v5).update((_2402_v8).f6, _module.__default.abs(_module.__default.fm2((_2402_v8).f6, globalState)))).Difference(_2399_v5);
          let _rhs409 = (_2410_v15).Merge(_2410_v15);
          let _rhs410 = !((_module.__default.safeModuloInt((_2407_v13).f18, new BigNumber(747))).isLessThanOrEqualTo((_2402_v8).f6));
          let _rhs411 = (_2402_v8).f6;
          let _rhs412 = _2406_v12;
          let _lhs328 = globalState;
          let _lhs329 = globalState;
          _2399_v5 = _rhs408;
          _2410_v15 = _rhs409;
          _lhs328.f0 = _rhs410;
          _lhs329.f5 = _rhs411;
          _2406_v12 = _rhs412;
          let _2424_v25;
          let _nw363 = Array((new BigNumber(23)).toNumber());
          _nw363[(_dafny.ZERO).toNumber()] = _2402_v8;
          _nw363[(_dafny.ONE).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(2)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(3)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(4)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(5)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(6)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(7)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(8)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(9)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(10)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(11)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(12)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(13)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(14)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(15)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(16)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(17)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(18)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(19)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(20)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(21)).toNumber()] = _2402_v8;
          _nw363[(new BigNumber(22)).toNumber()] = _2402_v8;
          _2424_v25 = _nw363;
          let _2425_v26;
          _2425_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2424_v25,_2423_v24);
          _2425_v26 = (_2425_v26).update(_2424_v25, _2423_v24);
          let _2426_v27;
          let _2427_v28;
          let _2428_v29;
          let _out37;
          let _out38;
          let _out39;
          let _outcollector12 = (_2407_v13).m13((_2407_v13).f18, globalState);
          _out37 = _outcollector12[0];
          _out38 = _outcollector12[1];
          _out39 = _outcollector12[2];
          _2426_v27 = _out37;
          _2427_v28 = _out38;
          _2428_v29 = _out39;
          let _2429_v30;
          _2429_v30 = _dafny.MultiSet.fromElements(_2396_v2);
          let _2430_v31;
          _2430_v31 = _module.D18.create_DC44(true, p1, (_2429_v30).Intersect(_2429_v30), _2407_v13);
          _2430_v31 = _2430_v31;
        }
        let _2431_v32;
        let _nw364 = new _module.C5();
        _nw364.__ctor((_2407_v13).f18, _module.__default.safeModuloInt((_this).f6, (_2402_v8).f6));
        _2431_v32 = _nw364;
        let _2432_v33;
        _2432_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,new _dafny.CodePoint('n'.codePointAt(0)));
        let _rhs413 = (_2407_v13).f17;
        let _rhs414 = new BigNumber((((_this.f8) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_2397_v3, _module.__default.safeIndex((_2431_v32).f15, new BigNumber((_2397_v3).length)), p0), _module.__default.safeIndex((_2402_v8).f6, new BigNumber((_dafny.Seq.update(_2397_v3, _module.__default.safeIndex((_2431_v32).f15, new BigNumber((_2397_v3).length)), p0)).length)), p0)).length),p0)) : (_2432_v33))).length);
        let _rhs415 = (_2407_v13).f18;
        let _rhs416 = _2431_v32;
        let _lhs330 = globalState;
        let _lhs331 = globalState;
        let _lhs332 = globalState;
        _lhs330.f5 = _rhs413;
        _lhs331.f5 = _rhs414;
        _lhs332.f5 = _rhs415;
        _2431_v32 = _rhs416;
        (globalState).f0 = p1;
      }
      let _2433_v34;
      let _init60 = function (_2434_i3) {
        return (_2434_i3).multipliedBy((_this).f6);
      };
      let _nw365 = Array((new BigNumber(2)).toNumber());
      for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw365.length); _i0_60++) {
        _nw365[_i0_60] = _init60(new BigNumber(_i0_60));
      }
      _2433_v34 = _nw365;
      let _2435_v35;
      let _nw366 = new _module.C7();
      _nw366.__ctor(_2433_v34, _2433_v34, (_2402_v8).f6, _2406_v12, _2401_v7);
      _2435_v35 = _nw366;
      let _2436_v36;
      let _nw367 = new _module.C2();
      _nw367.__ctor(new BigNumber((_2396_v2).length), _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_2407_v13).f17, (_2402_v8).f6), _2399_v5), p2, _2401_v7);
      _2436_v36 = _nw367;
      let _2437_v37;
      _2437_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2436_v36,true);
      let _source33 = _dafny.Map.Empty.slice().updateUnsafe(_2436_v36,_module.__default.fm3(globalState));
      let _2438___mcc_h0 = _source33;
      let _2439_cf68 = _2438___mcc_h0;
      let _2440_v38;
      _2440_v38 = new _dafny.CodePoint('a'.codePointAt(0));
      let _rhs417 = _2396_v2;
      let _rhs418 = p0;
      _2396_v2 = _rhs417;
      _2440_v38 = _rhs418;
      let _2441_v39;
      let _nw368 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _2441_v39 = _nw368;
      let _2442_v40;
      let _nw369 = Array((new BigNumber(27)).toNumber());
      _nw369[(_dafny.ZERO).toNumber()] = _2441_v39;
      _nw369[(_dafny.ONE).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(2)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(3)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(4)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(5)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(6)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(7)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(8)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(9)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(10)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(11)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(12)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(13)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(14)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(15)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(16)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(17)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(18)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(19)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(20)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(21)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(22)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(23)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(24)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(25)).toNumber()] = _2441_v39;
      _nw369[(new BigNumber(26)).toNumber()] = _2441_v39;
      _2442_v40 = _nw369;
      let _index333 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_2442_v40).length));
      (_2442_v40)[_index333] = _2441_v39;
      let _index334 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_2442_v40).length));
      let _rhs419 = _2441_v39;
      let _rhs420 = (_module.__default.safeModuloInt(new BigNumber(977), (_2402_v8).f6)).isLessThan((_2407_v13).f18);
      let _lhs333 = _2442_v40;
      let _lhs334 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_2442_v40).length));
      let _lhs335 = _this;
      _lhs333[_lhs334] = _rhs419;
      _lhs335.f8 = _rhs420;
      (_this).f8 = true;
      let _2443_v41;
      let _nw370 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2443_v41 = _nw370;
      let _index335 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_2443_v41).length));
      (_2443_v41)[_index335] = _2397_v3;
      let _index336 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_2443_v41).length));
      (_2443_v41)[_index336] = _dafny.Seq.Concat(_2397_v3, _dafny.Seq.Concat(_2397_v3, _2397_v3));
      r0 = _2397_v3;
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
