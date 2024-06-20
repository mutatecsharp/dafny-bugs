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
      return new BigNumber(74);
    };
    static fm1(p0, p1, globalState) {
      return (new BigNumber(-621)).isLessThan((new BigNumber(151)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("awbiybj"), _dafny.Seq.UnicodeFromString("d"));
    };
    static fm3(p0, p1, globalState) {
      return false;
    };
    static fm6(globalState) {
      return new _dafny.CodePoint('f'.codePointAt(0));
    };
    static fm7(p0, globalState) {
      return ((_dafny.Set.fromElements(!(true))).Intersect(_dafny.Set.fromElements(true))).Union(_dafny.Set.fromElements(true, !(true), !(true)));
    };
    static fm8(globalState) {
      return (_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))).Difference((_dafny.Set.fromElements(new _dafny.CodePoint('o'.codePointAt(0)))).Difference(_dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))));
    };
    static fm9(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false));
    };
    static m0(p0, p1, globalState) {
      let r0 = false;
      let r1 = _dafny.Set.Empty;
      let _0_v0;
      _0_v0 = false;
      let _1_i0;
      _1_i0 = _dafny.ZERO;
      L0: {
        while (_0_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1_i0)) {
              break L0;
            }
            _1_i0 = (_1_i0).plus(_dafny.ONE);
            let _2_v1;
            _2_v1 = new _dafny.CodePoint('q'.codePointAt(0));
            let _3_v2;
            _3_v2 = _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), _2_v1, _2_v1);
            let _4_v3;
            _4_v3 = _dafny.Map.Empty.slice().updateUnsafe(_0_v0,p0);
            let _5_v4;
            _5_v4 = _dafny.Set.fromElements(p1, p0, p0, new BigNumber((_4_v3).length), p0);
            let _6_v5;
            let _nw0 = new _module.C0();
            _nw0.__ctor(_3_v2, (p0).plus(p1), _3_v2, (_5_v4).contains(p1));
            _6_v5 = _nw0;
            _6_v5 = _6_v5;
            let _7_v6;
            _7_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hpidlcdvp"), _3_v2));
            let _8_v7;
            _8_v7 = _dafny.Set.fromElements(_0_v0);
            _3_v2 = (((_7_v6).contains(new BigNumber((((_0_v0) ? (_8_v7) : (_module.__default.fm7(p0, globalState)))).length))) ? ((_7_v6).get(new BigNumber((((_0_v0) ? (_8_v7) : (_module.__default.fm7(p0, globalState)))).length))) : (_dafny.Seq.Concat((_6_v5).f7, _3_v2)));
            let _9_v8;
            let _nw1 = new _module.C0();
            _nw1.__ctor(_3_v2, p1, (_6_v5).f7, _0_v0);
            _9_v8 = _nw1;
            (globalState).f1 = p0;
          }
        }
      }
      let _10_v9;
      _10_v9 = new _dafny.CodePoint('p'.codePointAt(0));
      let _11_v10;
      _11_v10 = _dafny.Seq.UnicodeFromString("u");
      let _12_v11;
      let _nw2 = new _module.C0();
      _nw2.__ctor(_dafny.Seq.of(_10_v9), new BigNumber((_11_v10).length), _11_v10, _0_v0);
      _12_v11 = _nw2;
      let _13_v12;
      _13_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(795)), function (_14_i1) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }),(_12_v11).f8);
      let _15_v13;
      _15_v13 = _module.D6.create_DC15(p0, _12_v11, _13_v12, _0_v0);
      if ((_15_v13).dtor_cf33) {
        let _16_v14;
        _16_v14 = _dafny.MultiSet.fromElements(_0_v0, _0_v0, _0_v0);
        (globalState).f1 = new BigNumber(((_16_v14).Difference((_16_v14).Union(_16_v14))).cardinality());
        (globalState).f1 = _module.__default.safeDivisionInt(new BigNumber(-28), p1);
        let _17_v15;
        _17_v15 = _dafny.Map.Empty.slice().updateUnsafe(_0_v0,_10_v9);
        let _pat_let_tv0 = _13_v12;
        let _pat_let_tv1 = _0_v0;
        let _pat_let_tv2 = _12_v11;
        let _18_v16;
        _18_v16 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let0_0) {
          return function (_19_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_20_dt__update_hcf32_h0) {
                return function (_pat_let2_0) {
                  return function (_21_dt__update_hcf33_h0) {
                    return function (_pat_let3_0) {
                      return function (_22_dt__update_hcf31_h0) {
                        return _module.D6.create_DC15((_19_dt__update__tmp_h0).dtor_cf30, _22_dt__update_hcf31_h0, _20_dt__update_hcf32_h0, _21_dt__update_hcf33_h0);
                      }(_pat_let3_0);
                    }(_pat_let_tv2);
                  }(_pat_let2_0);
                }(_pat_let_tv1);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_15_v13)).dtor_cf33,p1);
        let _23_v17;
        _23_v17 = _dafny.MultiSet.fromElements(_10_v9, _10_v9, _10_v9);
        let _24_v18;
        _24_v18 = _dafny.Map.Empty.slice().updateUnsafe(_12_v11,_0_v0);
        let _25_v19;
        _25_v19 = _dafny.Seq.of(_0_v0);
        let _26_v20;
        let _nw3 = Array((new BigNumber(28)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = ((_12_v11).f8).multipliedBy((_12_v11).f8);
        _nw3[(_dafny.ONE).toNumber()] = p0;
        _nw3[(new BigNumber(2)).toNumber()] = (p0).multipliedBy((_dafny.ZERO).minus(p1));
        _nw3[(new BigNumber(3)).toNumber()] = (p0).plus(p1);
        _nw3[(new BigNumber(4)).toNumber()] = new BigNumber(-691);
        _nw3[(new BigNumber(5)).toNumber()] = (_12_v11).f8;
        _nw3[(new BigNumber(6)).toNumber()] = p1;
        _nw3[(new BigNumber(7)).toNumber()] = (_12_v11).f8;
        _nw3[(new BigNumber(8)).toNumber()] = new BigNumber((_17_v15).length);
        _nw3[(new BigNumber(9)).toNumber()] = (((_18_v16).contains(_0_v0)) ? ((_18_v16).get(_0_v0)) : (new BigNumber((_11_v10).length)));
        _nw3[(new BigNumber(10)).toNumber()] = (_12_v11).f8;
        _nw3[(new BigNumber(11)).toNumber()] = (_12_v11).f8;
        _nw3[(new BigNumber(12)).toNumber()] = p0;
        _nw3[(new BigNumber(13)).toNumber()] = p1;
        _nw3[(new BigNumber(14)).toNumber()] = p1;
        _nw3[(new BigNumber(15)).toNumber()] = (_12_v11).f8;
        _nw3[(new BigNumber(16)).toNumber()] = p1;
        _nw3[(new BigNumber(17)).toNumber()] = ((_12_v11).f8).multipliedBy((_15_v13).dtor_cf30);
        _nw3[(new BigNumber(18)).toNumber()] = p0;
        _nw3[(new BigNumber(19)).toNumber()] = (_12_v11).f8;
        _nw3[(new BigNumber(20)).toNumber()] = (_12_v11).f8;
        _nw3[(new BigNumber(21)).toNumber()] = (((_23_v17).contains(_10_v9)) ? ((_23_v17).get(_10_v9)) : ((_12_v11).f8));
        _nw3[(new BigNumber(22)).toNumber()] = (_12_v11).f8;
        _nw3[(new BigNumber(23)).toNumber()] = (_12_v11).f8;
        _nw3[(new BigNumber(24)).toNumber()] = p0;
        _nw3[(new BigNumber(25)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, new BigNumber((_24_v18).length)));
        _nw3[(new BigNumber(26)).toNumber()] = new BigNumber((_25_v19).length);
        _nw3[(new BigNumber(27)).toNumber()] = (_12_v11).f8;
        _26_v20 = _nw3;
        let _index0 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_26_v20).length));
        (_26_v20)[_index0] = _module.__default.safeDivisionInt((((_16_v14).contains(_0_v0)) ? ((_16_v14).get(_0_v0)) : ((_12_v11).f8)), p1);
        let _index1 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_26_v20).length));
        (_26_v20)[_index1] = (_12_v11).f8;
        _25_v19 = _25_v19;
        (globalState).f0 = _module.__default.fm1(new BigNumber(740), _0_v0, globalState);
      } else {
        let _27_v21;
        _27_v21 = _0_v0;
        let _28_v22;
        let _nw4 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _28_v22 = _nw4;
        let _29_v23;
        _29_v23 = _dafny.Map.Empty.slice().updateUnsafe(_27_v21,_28_v22);
        _29_v23 = ((_29_v23).Merge(_29_v23)).Merge(_29_v23);
        let _30_v24;
        let _nw5 = new _module.C0();
        _nw5.__ctor(_11_v10, (_12_v11).f8, (_12_v11).f7, (_0_v0) && (_0_v0));
        _30_v24 = _nw5;
        let _31_v25;
        _31_v25 = _dafny.Seq.of(true);
        let _32_v26;
        let _nw6 = new _module.C0();
        _nw6.__ctor(_dafny.Seq.Concat((_30_v24).f7, _dafny.Seq.of(_10_v9, _module.__default.fm6(globalState), _10_v9)), new BigNumber((_dafny.Seq.UnicodeFromString("qf")).length), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), ((_33_v9) => function (_34_i2) {
          return _33_v9;
        })(_10_v9)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), ((_35_v9) => function (_36_i2) {
          return _35_v9;
        })(_10_v9))).length)), _10_v9), !((_31_v25)[_module.__default.safeIndex((_30_v24).f8, new BigNumber((_31_v25).length))]) || (_0_v0));
        _32_v26 = _nw6;
        let _37_v27;
        _37_v27 = _dafny.Map.Empty.slice().updateUnsafe(_10_v9,(_30_v24).f8);
        let _38_v28;
        _38_v28 = _dafny.Set.fromElements(_10_v9);
        let _39_v29;
        _39_v29 = _dafny.Map.Empty.slice().updateUnsafe(!(_37_v27).equals(_37_v27),((_0_v0) ? (_38_v28) : (_module.__default.fm8(globalState))));
        _39_v29 = (_39_v29).update(_0_v0, _38_v28);
        let _index2 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_28_v22).length));
        (_28_v22)[_index2] = (_32_v26).f8;
        let _40_v30;
        _40_v30 = _dafny.Map.Empty.slice().updateUnsafe(_28_v22,_0_v0);
        let _41_v31;
        _41_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_40_v30).length),new BigNumber(754));
        let _42_v32;
        _42_v32 = _dafny.Map.Empty.slice().updateUnsafe(_0_v0,_0_v0);
        let _43_v33;
        _43_v33 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_41_v31).length)), new BigNumber((_42_v32).length));
        let _44_v34;
        _44_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(_0_v0),(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(415)), function (_45_i3) {
          return new BigNumber(537);
        })).length)).plus(new BigNumber((_43_v33).length)));
        let _46_v35;
        _46_v35 = _dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus((_32_v26).f8)));
        let _index3 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_28_v22).length));
        (_28_v22)[_index3] = (((_44_v34).contains(!_dafny.areEqual(_31_v25, _31_v25))) ? ((_44_v34).get(!_dafny.areEqual(_31_v25, _31_v25))) : (new BigNumber(((_dafny.Set.fromElements(p0)).Union(_46_v35)).length)));
      }
      let _hi0 = p1;
      for (let _47_i4 = (_12_v11).f8; _47_i4.isLessThan(_hi0); _47_i4 = _47_i4.plus(_dafny.ONE)) {
        let _48_v36;
        _48_v36 = _dafny.Seq.of(p1, p1);
        if (_module.__default.fm1(((_0_v0) ? (p1) : ((_48_v36)[_module.__default.safeIndex(p1, new BigNumber((_48_v36).length))])), _dafny.Seq.IsProperPrefixOf(_11_v10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(324)), ((_49_v9) => function (_50_i5) {
          return _49_v9;
        })(_10_v9))), globalState)) {
          _11_v10 = (_12_v11).f7;
          (globalState).f0 = !((_0_v0) === (_0_v0));
          (globalState).f1 = ((((false) ? (_0_v0) : (_0_v0))) ? ((_12_v11).f8) : ((_dafny.ZERO).minus(p1)));
          let _51_v37;
          let _nw7 = new _module.C0();
          _nw7.__ctor(_11_v10, (_12_v11).f8, (_12_v11).f7, true);
          _51_v37 = _nw7;
          let _52_v38;
          _52_v38 = _dafny.Set.fromElements(_51_v37, _51_v37);
          let _53_v39;
          _53_v39 = _dafny.MultiSet.fromElements(_0_v0, _0_v0);
          let _54_v40;
          _54_v40 = _dafny.Map.Empty.slice().updateUnsafe(_52_v38,_53_v39);
          let _rhs0 = new BigNumber(((_12_v11).f7).length);
          let _rhs1 = _module.__default.fm6(globalState);
          let _rhs2 = _module.__default.safeDivisionInt(new BigNumber(((_54_v40).Merge(_54_v40)).length), (_dafny.ZERO).minus((_12_v11).f8));
          let _lhs0 = globalState;
          let _lhs1 = globalState;
          _lhs0.f1 = _rhs0;
          _10_v9 = _rhs1;
          _lhs1.f1 = _rhs2;
          let _55_v41;
          _55_v41 = _dafny.Map.Empty.slice().updateUnsafe(_0_v0,(_12_v11).f8);
          _55_v41 = (_55_v41).update(_0_v0, _module.__default.safeModuloInt(p0, p0));
        } else {
          _10_v9 = new _dafny.CodePoint('i'.codePointAt(0));
          let _56_v42;
          _56_v42 = _dafny.Seq.of(_12_v11, _12_v11, _12_v11, _12_v11, _12_v11);
          _56_v42 = _dafny.Seq.Concat(_56_v42, _56_v42);
          let _57_v43;
          _57_v43 = _dafny.Seq.of(_0_v0, false, _0_v0);
          let _58_v44;
          let _init0 = ((_59_v0) => function (_60_i6) {
            return _59_v0;
          })(_0_v0);
          let _nw8 = Array((new BigNumber(24)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw8.length); _i0_0++) {
            _nw8[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _58_v44 = _nw8;
          let _61_v45;
          _61_v45 = _dafny.Seq.of(_58_v44, _58_v44);
          let _62_v46;
          _62_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_61_v45).length),_dafny.Seq.contains((_12_v11).f7, _10_v9));
          let _rhs3 = (_12_v11).fm4((_12_v11).f8, new BigNumber(843), _dafny.Seq.Concat(_57_v43, _module.__default.fm9(_dafny.Seq.UnicodeFromString("klg"), p0, globalState)), globalState);
          let _rhs4 = new BigNumber((_62_v46).length);
          let _lhs2 = globalState;
          let _lhs3 = globalState;
          _lhs2.f1 = _rhs3;
          _lhs3.f1 = _rhs4;
          _12_v11 = _12_v11;
          _12_v11 = (_15_v13).dtor_cf31;
        }
        _13_v12 = (_13_v12).update(_11_v10, (_12_v11).f8);
        let _63_v47;
        let _nw9 = new _module.C0();
        _nw9.__ctor((_12_v11).f7, (_12_v11).f8, _dafny.Seq.UnicodeFromString("cwoxfwq"), _0_v0);
        _63_v47 = _nw9;
        let _64_v48;
        _64_v48 = _0_v0;
        let _65_v49;
        _65_v49 = _dafny.Map.Empty.slice().updateUnsafe(_0_v0,_0_v0);
        let _66_v50;
        let _nw10 = Array((new BigNumber(29)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _0_v0;
        _nw10[(_dafny.ONE).toNumber()] = true;
        _nw10[(new BigNumber(2)).toNumber()] = false;
        _nw10[(new BigNumber(3)).toNumber()] = _0_v0;
        _nw10[(new BigNumber(4)).toNumber()] = _0_v0;
        _nw10[(new BigNumber(5)).toNumber()] = _0_v0;
        _nw10[(new BigNumber(6)).toNumber()] = _0_v0;
        _nw10[(new BigNumber(7)).toNumber()] = (_64_v48);
        _nw10[(new BigNumber(8)).toNumber()] = _0_v0;
        _nw10[(new BigNumber(9)).toNumber()] = true;
        _nw10[(new BigNumber(10)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(11)).toNumber()] = _0_v0;
        _nw10[(new BigNumber(12)).toNumber()] = true;
        _nw10[(new BigNumber(13)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(14)).toNumber()] = _0_v0;
        _nw10[(new BigNumber(15)).toNumber()] = _0_v0;
        _nw10[(new BigNumber(16)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(17)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(18)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(19)).toNumber()] = !(_0_v0);
        _nw10[(new BigNumber(20)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(21)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(22)).toNumber()] = _0_v0;
        _nw10[(new BigNumber(23)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(24)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(25)).toNumber()] = false;
        _nw10[(new BigNumber(26)).toNumber()] = _63_v47.f6;
        _nw10[(new BigNumber(27)).toNumber()] = (((_65_v49).contains(_63_v47.f6)) ? ((_65_v49).get(_63_v47.f6)) : (_0_v0));
        _nw10[(new BigNumber(28)).toNumber()] = _63_v47.f6;
        _66_v50 = _nw10;
        let _67_v51;
        _67_v51 = _dafny.Map.Empty.slice().updateUnsafe(_63_v47,_66_v50);
        _67_v51 = (_67_v51).update(_63_v47, _66_v50);
        if ((p0).isLessThan(_47_i4)) {
          let _68_v52;
          let _nw11 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _68_v52 = _nw11;
          let _index4 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_68_v52).length));
          (_68_v52)[_index4] = (_48_v36)[_module.__default.safeIndex((_12_v11).f8, new BigNumber((_48_v36).length))];
          let _index5 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_68_v52).length));
          (_68_v52)[_index5] = (_12_v11).f8;
          let _index6 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_68_v52).length));
          (_68_v52)[_index6] = new BigNumber(749);
          let _69_v53;
          let _nw12 = new _module.C0();
          _nw12.__ctor((_12_v11).f7, new BigNumber(617), (_63_v47).f5, _63_v47.f6);
          _69_v53 = _nw12;
          let _index7 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_66_v50).length));
          (_66_v50)[_index7] = _63_v47.f6;
          let _70_v54;
          _70_v54 = _module.D6.create_DC13(_63_v47.f6, !(_0_v0));
          let _71_v55;
          let _init1 = ((_72_p0) => function (_73_i7) {
            return _dafny.Set.fromElements(_72_p0);
          })(p0);
          let _nw13 = Array((new BigNumber(19)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw13.length); _i0_1++) {
            _nw13[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _71_v55 = _nw13;
          let _74_v56;
          _74_v56 = _dafny.Set.fromElements((_12_v11).f8, (_68_v52)[_module.__default.safeIndex(new BigNumber(507), new BigNumber((_68_v52).length))], p1, (_69_v53).f8);
          let _index8 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_71_v55).length));
          (_71_v55)[_index8] = (_74_v56).Union(_74_v56);
          let _75_v57;
          let _nw14 = Array((new BigNumber(2)).toNumber());
          _nw14[(_dafny.ZERO).toNumber()] = ((_63_v47.f6) ? (_12_v11) : (_12_v11));
          _nw14[(_dafny.ONE).toNumber()] = _12_v11;
          _75_v57 = _nw14;
          let _76_v58;
          _76_v58 = _dafny.Map.Empty.slice().updateUnsafe(_12_v11,_65_v49);
          let _77_v59;
          _77_v59 = _dafny.Seq.of(_74_v56);
          let _pat_let_tv3 = _0_v0;
          let _index9 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_66_v50).length));
          let _index10 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_71_v55).length));
          let _rhs5 = (_dafny.Map.Empty.slice().updateUnsafe(_0_v0,_63_v47.f6)).Merge(((((_76_v58).contains(_12_v11)) ? ((_76_v58).get(_12_v11)) : (_65_v49))).update(_0_v0, _0_v0));
          let _rhs6 = _63_v47.f6;
          let _rhs7 = function (_pat_let4_0) {
            return function (_78_dt__update__tmp_h1) {
              return function (_pat_let5_0) {
                return function (_79_dt__update_hcf26_h0) {
                  return _module.D6.create_DC13(_79_dt__update_hcf26_h0, (_78_dt__update__tmp_h1).dtor_cf27);
                }(_pat_let5_0);
              }(_pat_let_tv3);
            }(_pat_let4_0);
          }(_70_v54);
          let _rhs8 = (_77_v59)[_module.__default.safeIndex((_69_v53).f8, new BigNumber((_77_v59).length))];
          let _rhs9 = _75_v57;
          let _lhs4 = _66_v50;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_66_v50).length));
          let _lhs6 = _71_v55;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_71_v55).length));
          _65_v49 = _rhs5;
          _lhs4[_lhs5] = _rhs6;
          _70_v54 = _rhs7;
          _lhs6[_lhs7] = _rhs8;
          _75_v57 = _rhs9;
          (globalState).f1 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))).plus((_12_v11).f8);
        } else {
          let _80_v60;
          let _nw15 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _80_v60 = _nw15;
          let _81_v61;
          let _nw16 = new _module.C0();
          _nw16.__ctor(_dafny.Seq.of(_module.__default.fm6(globalState)), (_dafny.ZERO).minus((_12_v11).f8), _dafny.Seq.Create(_module.__default.abs(new BigNumber(264)), ((_82_v9) => function (_83_i8) {
            return _82_v9;
          })(_10_v9)), _63_v47.f6);
          _81_v61 = _nw16;
          let _84_v62;
          _84_v62 = _dafny.Seq.of(_81_v61);
          let _85_v63;
          _85_v63 = _dafny.Seq.of(_84_v62, _84_v62, _84_v62, _84_v62, _84_v62);
          let _index11 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_80_v60).length));
          (_80_v60)[_index11] = (_85_v63)[_module.__default.safeIndex(_47_i4, new BigNumber((_85_v63).length))];
          let _index12 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_80_v60).length));
          (_80_v60)[_index12] = _84_v62;
          let _index13 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_66_v50).length));
          (_66_v50)[_index13] = _63_v47.f6;
          let _index14 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_66_v50).length));
          (_66_v50)[_index14] = _63_v47.f6;
          let _86_v64;
          _86_v64 = _dafny.Map.Empty.slice().updateUnsafe(_63_v47,(_12_v11).f8);
          _86_v64 = (_86_v64).update(_63_v47, new BigNumber((_dafny.Seq.Concat(_48_v36, _48_v36)).length));
          (_63_v47).f6 = (_66_v50)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_66_v50).length))];
          let _87_v65;
          let _nw17 = new _module.C0();
          _nw17.__ctor((_12_v11).f7, (_12_v11).f8, (_81_v61).f5, !(_63_v47.f6) || (true));
          _87_v65 = _nw17;
        }
      }
      _11_v10 = _dafny.Seq.Concat(_11_v10, (_12_v11).f7);
      let _88_v66;
      _88_v66 = _dafny.Set.fromElements(_0_v0);
      let _89_v67;
      _89_v67 = _dafny.Seq.of(_0_v0, !(_0_v0));
      let _hi1 = new BigNumber((_dafny.MultiSet.FromArray(_89_v67)).cardinality());
      for (let _90_i9 = new BigNumber((_88_v66).length); _90_i9.isLessThan(_hi1); _90_i9 = _90_i9.plus(_dafny.ONE)) {
        let _91_v68;
        let _nw18 = Array((new BigNumber(20)).toNumber()).fill([]);
        _91_v68 = _nw18;
        let _92_v69;
        _92_v69 = _dafny.Set.fromElements(p1);
        let _93_v70;
        _93_v70 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),false);
        let _94_v71;
        _94_v71 = _dafny.MultiSet.fromElements(p0);
        let _95_v72;
        _95_v72 = _dafny.Seq.of(new BigNumber((_94_v71).cardinality()), p1);
        let _96_v73;
        let _nw19 = Array((new BigNumber(11)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = new BigNumber((_92_v69).length);
        _nw19[(_dafny.ONE).toNumber()] = _90_i9;
        _nw19[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_93_v70).length));
        _nw19[(new BigNumber(3)).toNumber()] = new BigNumber(-43);
        _nw19[(new BigNumber(4)).toNumber()] = new BigNumber((_95_v72).length);
        _nw19[(new BigNumber(5)).toNumber()] = _90_i9;
        _nw19[(new BigNumber(6)).toNumber()] = new BigNumber((_94_v71).cardinality());
        _nw19[(new BigNumber(7)).toNumber()] = _90_i9;
        _nw19[(new BigNumber(8)).toNumber()] = (_12_v11).f8;
        _nw19[(new BigNumber(9)).toNumber()] = new BigNumber(-71);
        _nw19[(new BigNumber(10)).toNumber()] = (_95_v72)[_module.__default.safeIndex((_12_v11).f8, new BigNumber((_95_v72).length))];
        _96_v73 = _nw19;
        let _index15 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_91_v68).length));
        (_91_v68)[_index15] = _96_v73;
        let _index16 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_91_v68).length));
        (_91_v68)[_index16] = ((_0_v0) ? ((_96_v73)) : (_96_v73));
        let _97_v74;
        let _nw20 = new _module.C0();
        _nw20.__ctor((_12_v11).f7, p0, (_12_v11).f7, _0_v0);
        _97_v74 = _nw20;
        let _98_v75;
        let _nw21 = Array((new BigNumber(4)).toNumber());
        _98_v75 = _nw21;
        let _index17 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_98_v75).length));
        (_98_v75)[_index17] = _12_v11;
        let _99_v76;
        let _nw22 = Array((new BigNumber(14)).toNumber()).fill(false);
        _99_v76 = _nw22;
        let _index18 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_99_v76).length));
        (_99_v76)[_index18] = (_0_v0) && (_0_v0);
        let _index19 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_99_v76).length));
        (_99_v76)[_index19] = !(_88_v66).contains(!(_0_v0));
        let _index20 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_96_v73).length));
        (_96_v73)[_index20] = (_12_v11).f8;
        let _100_v77;
        _100_v77 = _dafny.Seq.of(_97_v74, _12_v11, _12_v11);
        let _index21 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_98_v75).length));
        let _index22 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_99_v76).length));
        let _index23 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_99_v76).length));
        let _index24 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_96_v73).length));
        let _rhs10 = (_100_v77)[_module.__default.safeIndex((_12_v11).f8, new BigNumber((_100_v77).length))];
        let _rhs11 = !(!(_0_v0));
        let _rhs12 = (_97_v74).f7;
        let _rhs13 = _0_v0;
        let _rhs14 = (_12_v11).f8;
        let _lhs8 = _98_v75;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_98_v75).length));
        let _lhs10 = _99_v76;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_99_v76).length));
        let _lhs12 = _99_v76;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_99_v76).length));
        let _lhs14 = _96_v73;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_96_v73).length));
        _lhs8[_lhs9] = _rhs10;
        _lhs10[_lhs11] = _rhs11;
        _11_v10 = _rhs12;
        _lhs12[_lhs13] = _rhs13;
        _lhs14[_lhs15] = _rhs14;
        let _101_v78;
        _101_v78 = _dafny.Map.Empty.slice().updateUnsafe((((_99_v76)[_module.__default.safeIndex(new BigNumber(771), new BigNumber((_99_v76).length))]) ? (p0) : ((_12_v11).f8)),(_97_v74).f8);
        let _102_v79;
        _102_v79 = _module.D6.create_DC13(_0_v0, _0_v0);
        let _103_v80;
        _103_v80 = _module.D6.create_DC16(_102_v79);
        let _104_v81;
        _104_v81 = _dafny.Map.Empty.slice().updateUnsafe(_103_v80,true);
        let _105_v82;
        _105_v82 = _dafny.Map.Empty.slice().updateUnsafe(_0_v0,_104_v81);
        let _index25 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_96_v73).length));
        (_96_v73)[_index25] = (((_101_v78).contains((_97_v74).f8)) ? ((_101_v78).get((_97_v74).f8)) : ((_dafny.ZERO).minus(new BigNumber((_105_v82).length))));
      }
      _13_v12 = (_13_v12).update(_11_v10, ((_12_v11).f8).multipliedBy(p1));
      r0 = _0_v0;
      let _106_v83;
      _106_v83 = _dafny.Set.fromElements((_12_v11).f8);
      r1 = _106_v83;
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _107_globalState;
      let _nw23 = new _module.GlobalState();
      _nw23.__ctor(false, new BigNumber(475), new _dafny.CodePoint('r'.codePointAt(0)), new BigNumber(449), false);
      _107_globalState = _nw23;
      let _108_v0;
      _108_v0 = new BigNumber(-316);
      (_107_globalState).f1 = _108_v0;
      let _109_v1;
      let _init2 = function (_110_i0) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eas"), _dafny.Seq.UnicodeFromString("opxmqoxt"));
      };
      let _nw24 = Array((new BigNumber(18)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw24.length); _i0_2++) {
        _nw24[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _109_v1 = _nw24;
      let _111_v2;
      _111_v2 = _dafny.Seq.UnicodeFromString("eedslel");
      let _112_v3;
      _112_v3 = _dafny.Map.Empty.slice().updateUnsafe(true,_111_v2);
      let _index26 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
      (_109_v1)[_index26] = _dafny.Seq.Concat((((_112_v3).contains(false)) ? ((_112_v3).get(false)) : (_111_v2)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-891)), function (_113_i1) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }));
      let _index27 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
      (_109_v1)[_index27] = _dafny.Seq.UnicodeFromString("osndtkw");
      let _114_v4;
      let _115_v5;
      let _out0;
      let _out1;
      let _outcollector0 = _module.__default.m0(_108_v0, (_dafny.ZERO).minus(_108_v0), _107_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _114_v4 = _out0;
      _115_v5 = _out1;
      let _116_v6;
      let _nw25 = Array((new BigNumber(18)).toNumber()).fill(false);
      _116_v6 = _nw25;
      let _117_v7;
      _117_v7 = _dafny.Set.fromElements(_116_v6, _116_v6);
      let _rhs15 = (_dafny.Set.fromElements(_116_v6)).Difference(_117_v7);
      let _rhs16 = _108_v0;
      let _rhs17 = _114_v4;
      let _lhs16 = _107_globalState;
      _117_v7 = _rhs15;
      _lhs16.f1 = _rhs16;
      _114_v4 = _rhs17;
      if (((_117_v7).Difference(_117_v7)).IsProperSubsetOf(_dafny.Set.fromElements(_116_v6))) {
        let _118_v8;
        _118_v8 = _dafny.Seq.of(_114_v4);
        let _119_v9;
        _119_v9 = _dafny.MultiSet.fromElements(_108_v0, _108_v0);
        let _120_v10;
        _120_v10 = _dafny.MultiSet.fromElements(_114_v4);
        let _121_v11;
        _121_v11 = new _dafny.CodePoint('p'.codePointAt(0));
        let _rhs18 = _dafny.Seq.contains(_118_v8, _114_v4);
        let _rhs19 = _114_v4;
        let _rhs20 = _module.__default.fm0((_dafny.MultiSet.fromElements(new BigNumber((_120_v10).cardinality()))).IsProperSubsetOf(_119_v9), _111_v2, _114_v4, _107_globalState);
        let _rhs21 = !(_114_v4) || (_dafny.Seq.contains((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))], _121_v11));
        let _lhs17 = _107_globalState;
        let _lhs18 = _107_globalState;
        let _lhs19 = _107_globalState;
        _lhs17.f0 = _rhs18;
        _114_v4 = _rhs19;
        _lhs18.f1 = _rhs20;
        _lhs19.f0 = _rhs21;
        if (!(_114_v4)) {
          (_107_globalState).f1 = new BigNumber((_118_v8).length);
          let _122_v12;
          _122_v12 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_dafny.ZERO).minus(_108_v0), _114_v4, _107_globalState),_114_v4);
          let _123_v13;
          _123_v13 = _dafny.Set.fromElements(_114_v4);
          let _rhs22 = (((_122_v12).contains(_114_v4)) ? ((_122_v12).get(_114_v4)) : ((_dafny.Set.fromElements(_114_v4, !(_114_v4))).IsProperSubsetOf(_123_v13)));
          let _rhs23 = _module.__default.fm1(_module.__default.safeModuloInt(new BigNumber((_111_v2).length), _108_v0), _114_v4, _107_globalState);
          _114_v4 = _rhs22;
          _114_v4 = _rhs23;
          let _124_v14;
          let _125_v15;
          let _out2;
          let _out3;
          let _outcollector1 = _module.__default.m0(new BigNumber(((_115_v5).Difference(_115_v5)).length), new BigNumber(416), _107_globalState);
          _out2 = _outcollector1[0];
          _out3 = _outcollector1[1];
          _124_v14 = _out2;
          _125_v15 = _out3;
          (_107_globalState).f0 = _124_v14;
          let _126_v16;
          let _nw26 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.of());
          _126_v16 = _nw26;
          let _index28 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_126_v16).length));
          (_126_v16)[_index28] = _118_v8;
          let _index29 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_126_v16).length));
          (_126_v16)[_index29] = _dafny.Seq.of(_114_v4);
        } else {
          _121_v11 = new _dafny.CodePoint('h'.codePointAt(0));
          let _127_v17;
          _127_v17 = _114_v4;
          let _128_v18;
          _128_v18 = _dafny.Map.Empty.slice().updateUnsafe((_127_v17),(_dafny.ZERO).minus(_108_v0));
          _128_v18 = (_128_v18).update(_114_v4, (_108_v0).plus(_108_v0));
          (_107_globalState).f0 = _114_v4;
          _108_v0 = _108_v0;
          let _129_v19;
          _129_v19 = _dafny.MultiSet.fromElements(_119_v9);
          (_107_globalState).f0 = !(_129_v19).contains(_dafny.MultiSet.fromElements(new BigNumber(765)));
        }
        let _130_v21;
        _130_v21 = _module.D1.create_DC3(false, new BigNumber(((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))]).length), _114_v4, _121_v11);
        let _rhs24 = (_module.D1.create_DC3(_114_v4, new BigNumber((function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of (_119_v9).Elements) {
    let _131_v20 = _compr_0;
    if ((_119_v9).contains(_131_v20)) {
      _coll0.push([_module.__default.safeDivisionInt(_131_v20, new BigNumber(((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))]).length)),_114_v4]);
    }
  }
  return _coll0;
}()).length), (_130_v21).dtor_cf9, _121_v11)).dtor_cf8;
        let _rhs25 = _114_v4;
        let _lhs20 = _107_globalState;
        _108_v0 = _rhs24;
        _lhs20.f0 = _rhs25;
        let _index30 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
        (_109_v1)[_index30] = _111_v2;
        let _132_v22;
        let _133_v23;
        let _out4;
        let _out5;
        let _outcollector2 = _module.__default.m0((_dafny.ZERO).minus(_108_v0), _108_v0, _107_globalState);
        _out4 = _outcollector2[0];
        _out5 = _outcollector2[1];
        _132_v22 = _out4;
        _133_v23 = _out5;
      } else {
        let _134_v24;
        let _nw27 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
        _134_v24 = _nw27;
        let _135_v25;
        _135_v25 = _dafny.Map.Empty.slice().updateUnsafe(_108_v0,true);
        let _index31 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_134_v24).length));
        (_134_v24)[_index31] = _135_v25;
        let _index32 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_134_v24).length));
        (_134_v24)[_index32] = ((_module.__default.fm1(_108_v0, _114_v4, _107_globalState)) ? (_135_v25) : (_dafny.Map.Empty.slice().updateUnsafe(_108_v0,_114_v4)));
        let _136_v26;
        _136_v26 = _dafny.Seq.of(_108_v0, (_dafny.ZERO).minus(new BigNumber(((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))]).length)));
        let _137_v27;
        _137_v27 = _dafny.MultiSet.fromElements(_114_v4);
        if (((_dafny.ZERO).minus(((((_137_v27).contains(_114_v4)) ? ((_137_v27).get(_114_v4)) : (_108_v0))).minus(_108_v0))).isLessThan(((_module.__default.fm1((_dafny.ZERO).minus(_108_v0), _114_v4, _107_globalState)) ? (_108_v0) : (new BigNumber((_136_v26).length))))) {
          let _138_v28;
          let _nw28 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
          _138_v28 = _nw28;
          let _139_v29;
          _139_v29 = _dafny.Seq.of(_114_v4, _114_v4);
          let _index33 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_138_v28).length));
          (_138_v28)[_index33] = _dafny.Seq.Concat(_139_v29, _139_v29);
          let _index34 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_138_v28).length));
          (_138_v28)[_index34] = _139_v29;
          (_107_globalState).f1 = new BigNumber(((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))]).length);
          let _140_v30;
          _140_v30 = _module.D1.create_DC1(_108_v0);
          _140_v30 = _140_v30;
          let _141_v31;
          let _142_v32;
          let _out6;
          let _out7;
          let _outcollector3 = _module.__default.m0(_108_v0, _module.__default.safeModuloInt(_108_v0, _108_v0), _107_globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _141_v31 = _out6;
          _142_v32 = _out7;
          let _nw29 = Array((new BigNumber(8)).toNumber()).fill(false);
          _116_v6 = _nw29;
        } else {
          let _143_v33;
          _143_v33 = new _dafny.CodePoint('s'.codePointAt(0));
          _143_v33 = _143_v33;
          let _144_v34;
          _144_v34 = _module.D1.create_DC2(_dafny.Seq.Concat(_dafny.Seq.of(_108_v0), _136_v26), _108_v0, _108_v0, ((_114_v4) ? (_111_v2) : (_module.__default.fm2(true, _114_v4, _108_v0, _114_v4, _107_globalState))), _115_v5);
          _144_v34 = _144_v34;
          let _index35 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_116_v6).length));
          (_116_v6)[_index35] = _114_v4;
          let _index36 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_116_v6).length));
          (_116_v6)[_index36] = !(((_114_v4) ? (_114_v4) : (!(_114_v4)))) || (_114_v4);
          let _145_v35;
          _145_v35 = _dafny.Seq.of(_114_v4, _114_v4);
          let _146_v36;
          let _147_v37;
          let _out8;
          let _out9;
          let _outcollector4 = _module.__default.m0(new BigNumber((_145_v35).length), (new BigNumber((_111_v2).length)).minus(_108_v0), _107_globalState);
          _out8 = _outcollector4[0];
          _out9 = _outcollector4[1];
          _146_v36 = _out8;
          _147_v37 = _out9;
          let _index37 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_116_v6).length));
          let _rhs26 = !(_108_v0).isEqualTo(_108_v0);
          let _rhs27 = _114_v4;
          let _lhs21 = _116_v6;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_116_v6).length));
          _114_v4 = _rhs26;
          _lhs21[_lhs22] = _rhs27;
        }
        let _index38 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
        (_109_v1)[_index38] = (_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))];
        let _148_v38;
        let _nw30 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
        _148_v38 = _nw30;
        let _149_v39;
        _149_v39 = _dafny.Map.Empty.slice().updateUnsafe(_114_v4,_148_v38);
        _149_v39 = (_149_v39).update(false, _148_v38);
        let _150_v40;
        _150_v40 = _dafny.Seq.of(_module.__default.fm1((_136_v26)[_module.__default.safeIndex(_108_v0, new BigNumber((_136_v26).length))], false, _107_globalState));
        (_107_globalState).f0 = !((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_114_v4, true), _150_v40))).IsDisjointFrom(_137_v27));
      }
      let _151_v41;
      let _init3 = ((_152_v0) => function (_153_i2) {
        return _module.D1.create_DC4(_module.D1.create_DC1(_152_v0));
      })(_108_v0);
      let _nw31 = Array((new BigNumber(9)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw31.length); _i0_3++) {
        _nw31[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _151_v41 = _nw31;
      let _154_v42;
      _154_v42 = _dafny.MultiSet.fromElements(_108_v0);
      let _155_v43;
      _155_v43 = _dafny.Map.Empty.slice().updateUnsafe(_114_v4,_154_v42);
      let _156_v44;
      let _nw32 = Array((new BigNumber(14)).toNumber());
      _nw32[(_dafny.ZERO).toNumber()] = new BigNumber((_111_v2).length);
      _nw32[(_dafny.ONE).toNumber()] = _108_v0;
      _nw32[(new BigNumber(2)).toNumber()] = _108_v0;
      _nw32[(new BigNumber(3)).toNumber()] = new BigNumber((_155_v43).length);
      _nw32[(new BigNumber(4)).toNumber()] = _108_v0;
      _nw32[(new BigNumber(5)).toNumber()] = _108_v0;
      _nw32[(new BigNumber(6)).toNumber()] = new BigNumber(99);
      _nw32[(new BigNumber(7)).toNumber()] = _108_v0;
      _nw32[(new BigNumber(8)).toNumber()] = _108_v0;
      _nw32[(new BigNumber(9)).toNumber()] = new BigNumber(861);
      _nw32[(new BigNumber(10)).toNumber()] = _108_v0;
      _nw32[(new BigNumber(11)).toNumber()] = _108_v0;
      _nw32[(new BigNumber(12)).toNumber()] = _108_v0;
      _nw32[(new BigNumber(13)).toNumber()] = new BigNumber(623);
      _156_v44 = _nw32;
      let _157_v45;
      _157_v45 = _dafny.Seq.of(_156_v44, _156_v44);
      let _158_v46;
      _158_v46 = _dafny.Map.Empty.slice().updateUnsafe((_157_v45)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_157_v45).length))],_114_v4);
      let _159_v47;
      _159_v47 = _module.D1.create_DC1(new BigNumber((_158_v46).length));
      let _160_v48;
      _160_v48 = _module.D1.create_DC4(_159_v47);
      let _index39 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_151_v41).length));
      (_151_v41)[_index39] = _160_v48;
      let _161_v49;
      let _nw33 = Array((new BigNumber(23)).toNumber()).fill([]);
      _161_v49 = _nw33;
      let _162_v50;
      _162_v50 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), function (_163_i3) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length));
      let _164_v51;
      _164_v51 = _dafny.Set.fromElements(_114_v4);
      let _pat_let_tv4 = _159_v47;
      let _index40 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_151_v41).length));
      let _rhs28 = _108_v0;
      let _rhs29 = function (_pat_let6_0) {
        return function (_165_dt__update__tmp_h0) {
          return function (_pat_let7_0) {
            return function (_166_dt__update_hcf11_h0) {
              return _module.D1.create_DC4(_166_dt__update_hcf11_h0);
            }(_pat_let7_0);
          }(_pat_let_tv4);
        }(_pat_let6_0);
      }(_160_v48);
      let _rhs30 = _module.__default.fm1(((((_162_v50).contains(_114_v4)) ? ((_162_v50).get(_114_v4)) : (_108_v0))).multipliedBy(_108_v0), ((_114_v4) ? (_114_v4) : (true)), _107_globalState);
      let _rhs31 = _161_v49;
      let _rhs32 = (_dafny.ZERO).minus((((_154_v42).contains(new BigNumber((_164_v51).length))) ? ((_154_v42).get(new BigNumber((_164_v51).length))) : ((_dafny.ZERO).minus(_108_v0))));
      let _lhs23 = _151_v41;
      let _lhs24 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_151_v41).length));
      let _lhs25 = _107_globalState;
      let _lhs26 = _107_globalState;
      _108_v0 = _rhs28;
      _lhs23[_lhs24] = _rhs29;
      _lhs25.f0 = _rhs30;
      _161_v49 = _rhs31;
      _lhs26.f1 = _rhs32;
      _108_v0 = _108_v0;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_116_v6).length))) {
        let _167_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_167_i4)) && ((_167_i4).isLessThan(new BigNumber((_116_v6).length))))) {
          (_116_v6)[(_167_i4)] = _114_v4;
        }
      }
      let _168_v52;
      _168_v52 = new _dafny.CodePoint('f'.codePointAt(0));
      let _169_v53;
      _169_v53 = _module.D1.create_DC3(false, _108_v0, true, _168_v52);
      let _170_v54;
      _170_v54 = _dafny.Set.fromElements(_169_v53, _169_v53);
      _170_v54 = _170_v54;
      let _171_v55;
      _171_v55 = _dafny.Seq.of(_108_v0);
      _111_v2 = _dafny.Seq.update(_111_v2, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((((_162_v50).update(_114_v4, new BigNumber((_171_v55).length))).Merge(_162_v50)).length)), new BigNumber((_111_v2).length)), _168_v52);
      let _source0 = _169_v53;
      if (_source0.is_DC2) {
        let _172___mcc_h0 = (_source0).cf2;
        let _173___mcc_h1 = (_source0).cf3;
        let _174___mcc_h2 = (_source0).cf4;
        let _175___mcc_h3 = (_source0).cf5;
        let _176___mcc_h4 = (_source0).cf6;
        let _177_cf6 = _176___mcc_h4;
        let _178_cf5 = _175___mcc_h3;
        let _179_cf4 = _174___mcc_h2;
        let _180_cf3 = _173___mcc_h1;
        let _181_cf2 = _172___mcc_h0;
        _114_v4 = !(false);
        let _182_v56;
        _182_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(595),(_dafny.ZERO).minus(_180_cf3));
        _182_v56 = (_182_v56).update(new BigNumber(526), _108_v0);
        _154_v42 = ((_dafny.MultiSet.fromElements(_180_cf3, _108_v0)).Intersect(_154_v42)).Difference(_154_v42);
        let _183_v57;
        _183_v57 = _114_v4;
        let _source1 = _183_v57;
        let _184___mcc_h11 = _source1;
        let _185_cf0 = _184___mcc_h11;
        let _index41 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_116_v6).length));
        (_116_v6)[_index41] = _114_v4;
        let _186_v58;
        _186_v58 = _dafny.Map.Empty.slice().updateUnsafe(_114_v4,_185_cf0);
        let _187_v59;
        _187_v59 = _dafny.Map.Empty.slice().updateUnsafe((_169_v53).dtor_cf8,_168_v52);
        let _188_v60;
        _188_v60 = _dafny.MultiSet.fromElements((((_187_v59).contains(_180_cf3)) ? ((_187_v59).get(_180_cf3)) : (_168_v52)));
        let _189_v61;
        _189_v61 = _dafny.Map.Empty.slice().updateUnsafe(_179_cf4,_188_v60);
        let _index42 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_116_v6).length));
        (_116_v6)[_index42] = (((_186_v58).contains((_dafny.MultiSet.FromArray(_111_v2)).IsDisjointFrom((((_189_v61).contains(_108_v0)) ? ((_189_v61).get(_108_v0)) : (_188_v60))))) ? ((_186_v58).get((_dafny.MultiSet.FromArray(_111_v2)).IsDisjointFrom((((_189_v61).contains(_108_v0)) ? ((_189_v61).get(_108_v0)) : (_188_v60))))) : (_114_v4));
        _178_cf5 = _178_cf5;
        (_107_globalState).f0 = _114_v4;
        _164_v51 = _164_v51;
      } else if (_source0.is_DC3) {
        let _190___mcc_h5 = (_source0).cf7;
        let _191___mcc_h6 = (_source0).cf8;
        let _192___mcc_h7 = (_source0).cf9;
        let _193___mcc_h8 = (_source0).cf10;
        let _194_cf10 = _193___mcc_h8;
        let _195_cf9 = _192___mcc_h7;
        let _196_cf8 = _191___mcc_h6;
        let _197_cf7 = _190___mcc_h5;
        _115_v5 = (_115_v5).Difference(_115_v5);
        let _198_v62;
        _198_v62 = _197_cf7;
        let _199_v63;
        _199_v63 = _dafny.Map.Empty.slice().updateUnsafe(_108_v0,_198_v62);
        let _200_v64;
        _200_v64 = _dafny.Map.Empty.slice().updateUnsafe(_116_v6,_108_v0);
        let _201_v65;
        let _202_v66;
        let _out10;
        let _out11;
        let _outcollector5 = _module.__default.m0(_module.__default.safeModuloInt(new BigNumber((_199_v63).length), _196_cf8), new BigNumber(((_200_v64).Merge(_200_v64)).length), _107_globalState);
        _out10 = _outcollector5[0];
        _out11 = _outcollector5[1];
        _201_v65 = _out10;
        _202_v66 = _out11;
        if (!(_197_cf7)) {
          _156_v44 = _156_v44;
          _154_v42 = _154_v42;
          (_107_globalState).f1 = new BigNumber((_111_v2).length);
          let _index43 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_156_v44).length));
          (_156_v44)[_index43] = ((_197_cf7) ? ((_dafny.ZERO).minus(_196_cf8)) : (new BigNumber(-94)));
          let _index44 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_156_v44).length));
          (_156_v44)[_index44] = new BigNumber(532);
          let _203_v67;
          _203_v67 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm3((_156_v44)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_156_v44).length))], _108_v0, _107_globalState)),_201_v65);
          _203_v67 = _203_v67;
        } else {
          let _204_v68;
          _204_v68 = _dafny.Seq.of(_201_v65);
          let _205_v69;
          let _206_v70;
          let _out12;
          let _out13;
          let _outcollector6 = _module.__default.m0((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_204_v68, _204_v68)).length)), (_dafny.ZERO).minus(_108_v0), _107_globalState);
          _out12 = _outcollector6[0];
          _out13 = _outcollector6[1];
          _205_v69 = _out12;
          _206_v70 = _out13;
          let _207_v71;
          let _208_v72;
          let _out14;
          let _out15;
          let _outcollector7 = _module.__default.m0(new BigNumber(320), (_196_cf8).multipliedBy(_196_cf8), _107_globalState);
          _out14 = _outcollector7[0];
          _out15 = _outcollector7[1];
          _207_v71 = _out14;
          _208_v72 = _out15;
          let _209_v73;
          _209_v73 = _156_v44;
          let _210_v74;
          _210_v74 = _dafny.MultiSet.fromElements(_156_v44, _156_v44, (_209_v73));
          let _rhs33 = _197_cf7;
          let _rhs34 = new BigNumber((_dafny.Seq.Concat(_111_v2, _111_v2)).length);
          let _rhs35 = (_210_v74).IsSubsetOf((_210_v74).Difference(_210_v74));
          _195_cf9 = _rhs33;
          _196_cf8 = _rhs34;
          _197_cf7 = _rhs35;
          let _211_v75;
          let _nw34 = new _module.C0();
          _nw34.__ctor(_dafny.Seq.update(_111_v2, _module.__default.safeIndex(_196_cf8, new BigNumber((_111_v2).length)), _168_v52), (_108_v0).multipliedBy(_108_v0), _dafny.Seq.UnicodeFromString("homss"), _197_cf7);
          _211_v75 = _nw34;
          let _212_v76;
          let _213_v77;
          let _out16;
          let _out17;
          let _outcollector8 = _module.__default.m0(_108_v0, new BigNumber(90), _107_globalState);
          _out16 = _outcollector8[0];
          _out17 = _outcollector8[1];
          _212_v76 = _out16;
          _213_v77 = _out17;
        }
        _195_cf9 = _197_cf7;
      } else if (_source0.is_DC1) {
        let _214___mcc_h9 = (_source0).cf1;
        let _215_cf1 = _214___mcc_h9;
        let _216_v78;
        _216_v78 = _dafny.Map.Empty.slice().updateUnsafe(_116_v6,_114_v4);
        let _217_v79;
        _217_v79 = _dafny.MultiSet.fromElements(_114_v4, !((((_216_v78).contains(_116_v6)) ? ((_216_v78).get(_116_v6)) : (_114_v4))), _114_v4, _114_v4, _114_v4);
        (_107_globalState).f0 = ((_114_v4) === (_114_v4)) === ((_217_v79).equals(_217_v79));
        let _218_v81;
        _218_v81 = _dafny.Map.Empty.slice().updateUnsafe(_215_cf1,_215_cf1);
        let _219_v82;
        _219_v82 = _dafny.Seq.of(_218_v81);
        (_107_globalState).f1 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber((function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_219_v82).Elements) {
            let _220_v80 = _compr_1;
            if (_dafny.Seq.contains(_219_v82, _220_v80)) {
              _coll1.push([_220_v80,_114_v4]);
            }
          }
          return _coll1;
        }()).length), _215_cf1), (new BigNumber(105)).plus(_108_v0));
        let _index45 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_156_v44).length));
        (_156_v44)[_index45] = _108_v0;
        let _221_v83;
        _221_v83 = _dafny.Map.Empty.slice().updateUnsafe(_111_v2,_114_v4);
        let _index46 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
        let _index47 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_156_v44).length));
        let _rhs36 = _dafny.Seq.UnicodeFromString("s");
        let _rhs37 = _108_v0;
        let _rhs38 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_221_v83).length), _108_v0), (_108_v0).minus(_215_cf1));
        let _rhs39 = ((_114_v4) ? (new BigNumber(((_164_v51).Union(_164_v51)).length)) : (_215_cf1));
        let _rhs40 = _108_v0;
        let _lhs27 = _109_v1;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
        let _lhs29 = _156_v44;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_156_v44).length));
        let _lhs31 = _107_globalState;
        _lhs27[_lhs28] = _rhs36;
        _lhs29[_lhs30] = _rhs37;
        _lhs31.f1 = _rhs38;
        _108_v0 = _rhs39;
        _215_cf1 = _rhs40;
        if ((_114_v4) === (((_114_v4) ? (_114_v4) : (_114_v4)))) {
          let _index48 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
          (_109_v1)[_index48] = (_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))];
          let _index49 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
          (_109_v1)[_index49] = _111_v2;
          (_107_globalState).f0 = false;
          (_107_globalState).f0 = _114_v4;
          let _222_v84;
          _222_v84 = _dafny.Map.Empty.slice().updateUnsafe((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))],_168_v52);
          _222_v84 = (_222_v84).update((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))], _168_v52);
        } else {
          let _index50 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_116_v6).length));
          (_116_v6)[_index50] = _114_v4;
          let _index51 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_116_v6).length));
          (_116_v6)[_index51] = _114_v4;
          (_107_globalState).f0 = _114_v4;
          let _223_v85;
          let _nw35 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _223_v85 = _nw35;
          let _224_v86;
          _224_v86 = _dafny.Seq.of(true, (_116_v6)[_module.__default.safeIndex(new BigNumber(230), new BigNumber((_116_v6).length))]);
          let _225_v87;
          _225_v87 = _dafny.Map.Empty.slice().updateUnsafe(false,_224_v86);
          let _226_v88;
          _226_v88 = _225_v87;
          let _index52 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_223_v85).length));
          (_223_v85)[_index52] = _226_v88;
          let _227_v89;
          _227_v89 = (_116_v6)[_module.__default.safeIndex(new BigNumber(230), new BigNumber((_116_v6).length))];
          let _228_v90;
          let _nw36 = new _module.C0();
          _nw36.__ctor(_dafny.Seq.of(_168_v52), _module.__default.fm0(_114_v4, (_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))], (_227_v89), _107_globalState), (_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))], (_116_v6)[_module.__default.safeIndex(new BigNumber(230), new BigNumber((_116_v6).length))]);
          _228_v90 = _nw36;
          let _index53 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_223_v85).length));
          let _rhs41 = _226_v88;
          let _rhs42 = _228_v90;
          let _lhs32 = _223_v85;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_223_v85).length));
          _lhs32[_lhs33] = _rhs41;
          _228_v90 = _rhs42;
          _168_v52 = _module.__default.fm6(_107_globalState);
          let _229_v91;
          let _nw37 = new _module.C0();
          _nw37.__ctor((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))], (_156_v44)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_156_v44).length))], _dafny.Seq.UnicodeFromString("lf"), (_116_v6)[_module.__default.safeIndex(new BigNumber(230), new BigNumber((_116_v6).length))]);
          _229_v91 = _nw37;
          let _230_v92;
          _230_v92 = _dafny.Map.Empty.slice().updateUnsafe(_114_v4,_229_v91);
          let _231_v93;
          let _nw38 = Array((new BigNumber(26)).toNumber());
          _nw38[(_dafny.ZERO).toNumber()] = _229_v91;
          _nw38[(_dafny.ONE).toNumber()] = _229_v91;
          _nw38[(new BigNumber(2)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(3)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(4)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(5)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(6)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(7)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(8)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(9)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(10)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(11)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(12)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(13)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(14)).toNumber()] = (((_230_v92).contains((_116_v6)[_module.__default.safeIndex(new BigNumber(230), new BigNumber((_116_v6).length))])) ? ((_230_v92).get((_116_v6)[_module.__default.safeIndex(new BigNumber(230), new BigNumber((_116_v6).length))])) : (_229_v91));
          _nw38[(new BigNumber(15)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(16)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(17)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(18)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(19)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(20)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(21)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(22)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(23)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(24)).toNumber()] = _229_v91;
          _nw38[(new BigNumber(25)).toNumber()] = _229_v91;
          _231_v93 = _nw38;
          let _232_v94;
          let _nw39 = Array((new BigNumber(7)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _231_v93;
          _nw39[(_dafny.ONE).toNumber()] = _231_v93;
          _nw39[(new BigNumber(2)).toNumber()] = _231_v93;
          _nw39[(new BigNumber(3)).toNumber()] = _231_v93;
          _nw39[(new BigNumber(4)).toNumber()] = _231_v93;
          _nw39[(new BigNumber(5)).toNumber()] = ((_114_v4) ? (_231_v93) : (_231_v93));
          _nw39[(new BigNumber(6)).toNumber()] = _231_v93;
          _232_v94 = _nw39;
          let _index54 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_232_v94).length));
          (_232_v94)[_index54] = _231_v93;
          let _index55 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_232_v94).length));
          (_232_v94)[_index55] = _231_v93;
        }
      } else {
        let _233___mcc_h10 = (_source0).cf11;
        let _234_cf11 = _233___mcc_h10;
        let _235_v95;
        let _nw40 = new _module.C0();
        _nw40.__ctor((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))], _108_v0, _111_v2, !(!(_114_v4)));
        _235_v95 = _nw40;
        let _236_v96;
        _236_v96 = _dafny.Map.Empty.slice().updateUnsafe(_235_v95,_116_v6);
        _236_v96 = (_236_v96).update(_235_v95, _116_v6);
        let _237_v97;
        _237_v97 = _114_v4;
        (_107_globalState).f0 = (((_237_v97)) ? (_114_v4) : (_114_v4));
        let _238_v98;
        _238_v98 = _dafny.Seq.of(!(_114_v4), _114_v4);
        let _239_v99;
        let _nw41 = new _module.C0();
        _nw41.__ctor((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))], (_235_v95).fm4((_235_v95).f8, _108_v0, _238_v98, _107_globalState), _dafny.Seq.UnicodeFromString("xalwfr"), _114_v4);
        _239_v99 = _nw41;
        _239_v99 = _239_v99;
        let _240_v100;
        _240_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat((_239_v99).f5, _dafny.Seq.UnicodeFromString("le"))).length),_114_v4);
        _240_v100 = (_240_v100).update(_module.__default.safeModuloInt((_235_v95).f8, (_171_v55)[_module.__default.safeIndex((_235_v95).f8, new BigNumber((_171_v55).length))]), _114_v4);
      }
      let _index56 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length));
      (_116_v6)[_index56] = !(_114_v4);
      let _index57 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length));
      (_116_v6)[_index57] = !(false);
      _111_v2 = _dafny.Seq.update(_module.__default.fm2(((_114_v4) ? ((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]) : ((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))])), _114_v4, _108_v0, _114_v4, _107_globalState), _module.__default.safeIndex((_108_v0).plus(_module.__default.fm0((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))], _111_v2, _module.__default.fm1(_108_v0, false, _107_globalState), _107_globalState)), new BigNumber((_module.__default.fm2(((_114_v4) ? ((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]) : ((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))])), _114_v4, _108_v0, _114_v4, _107_globalState)).length)), _168_v52);
      let _index58 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_156_v44).length));
      (_156_v44)[_index58] = new BigNumber((_111_v2).length);
      let _index59 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_156_v44).length));
      (_156_v44)[_index59] = new BigNumber(82);
      let _241_v101;
      _241_v101 = _dafny.Set.fromElements(_156_v44, _156_v44);
      let _242_v102;
      _242_v102 = _dafny.Map.Empty.slice().updateUnsafe(_241_v101,(true) || ((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]));
      _242_v102 = (_242_v102).update(_dafny.Set.fromElements(_156_v44), (_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]);
      let _hi2 = (_dafny.ZERO).minus(_108_v0);
      for (let _243_i5 = _108_v0; _243_i5.isLessThan(_hi2); _243_i5 = _243_i5.plus(_dafny.ONE)) {
        let _244_v103;
        let _nw42 = new _module.C0();
        _nw42.__ctor(_111_v2, _module.__default.safeDivisionInt(_108_v0, (_156_v44)[_module.__default.safeIndex(new BigNumber(579), new BigNumber((_156_v44).length))]), _dafny.Seq.Concat((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))], (((_112_v3).contains(!((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]))) ? ((_112_v3).get(!((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(902)), ((_245_v52) => function (_246_i6) {
          return _245_v52;
        })(_168_v52))))), (_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]);
        _244_v103 = _nw42;
        if ((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]) {
          let _247_v105;
          _247_v105 = _dafny.MultiSet.fromElements((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))], _114_v4);
          let _248_v106;
          _248_v106 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of _dafny.IntegerRange(new BigNumber(79), new BigNumber(-6))) {
              let _249_v104 = _compr_2;
              if (((new BigNumber(79)).isLessThanOrEqualTo(_249_v104)) && ((_249_v104).isLessThan(new BigNumber(-6)))) {
                _coll2.push([(_249_v104).plus(new BigNumber(((_244_v103).f7).length)),(_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]]);
              }
            }
            return _coll2;
          }()).length),_247_v105);
          let _index60 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
          let _rhs43 = (_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))];
          let _rhs44 = _dafny.Seq.UnicodeFromString("tkbtcsqb");
          let _rhs45 = _108_v0;
          let _rhs46 = !(!((((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]) ? (_248_v106) : (_248_v106))).contains(new BigNumber(2)));
          let _rhs47 = !((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]);
          let _lhs34 = _109_v1;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
          let _lhs36 = _107_globalState;
          let _lhs37 = _107_globalState;
          _114_v4 = _rhs43;
          _lhs34[_lhs35] = _rhs44;
          _lhs36.f1 = _rhs45;
          _lhs37.f0 = _rhs46;
          _114_v4 = _rhs47;
          let _index61 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
          let _rhs48 = (_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))];
          let _rhs49 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ilf"), ((true) ? (_111_v2) : ((_109_v1)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length))])));
          let _rhs50 = (_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))];
          let _lhs38 = _109_v1;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_109_v1).length));
          let _lhs40 = _107_globalState;
          _111_v2 = _rhs48;
          _lhs38[_lhs39] = _rhs49;
          _lhs40.f0 = _rhs50;
          let _250_v107;
          _250_v107 = _dafny.Map.Empty.slice().updateUnsafe(_114_v4,_244_v103);
          let _251_v108;
          _251_v108 = _dafny.Seq.of(_244_v103);
          let _252_v110;
          _252_v110 = _dafny.Seq.of(_115_v5, _115_v5, function () {
            let _coll3 = new _dafny.Set();
            for (const _compr_3 of _dafny.IntegerRange(new BigNumber(172), new BigNumber(-314))) {
              let _253_v109 = _compr_3;
              if (((new BigNumber(172)).isLessThanOrEqualTo(_253_v109)) && ((_253_v109).isLessThan(new BigNumber(-314)))) {
                _coll3.add(_module.__default.safeDivisionInt(_253_v109, _108_v0));
              }
            }
            return _coll3;
          }(), _115_v5, _115_v5);
          let _254_v111;
          _254_v111 = _dafny.Map.Empty.slice().updateUnsafe(_243_i5,(_115_v5).IsProperSubsetOf((_252_v110)[_module.__default.safeIndex(_108_v0, new BigNumber((_252_v110).length))]));
          let _255_v112;
          _255_v112 = _dafny.Seq.of(_114_v4, _114_v4, (_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))], !(_114_v4));
          let _rhs51 = _dafny.Map.Empty.slice().updateUnsafe((_114_v4) || (_114_v4),_244_v103);
          let _rhs52 = ((_244_v103).f7)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_171_v55)[_module.__default.safeIndex(new BigNumber((_251_v108).length), new BigNumber((_171_v55).length))], new BigNumber(((_244_v103).f7).length), (((_247_v105).contains((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))])) ? ((_247_v105).get((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))])) : (new BigNumber((_module.__default.fm7((_244_v103).f8, _107_globalState)).length))), _243_i5)).length), new BigNumber(((_244_v103).f7).length))];
          let _rhs53 = (((_254_v111).contains(new BigNumber((_dafny.Seq.Concat(_255_v112, _255_v112)).length))) ? ((_254_v111).get(new BigNumber((_dafny.Seq.Concat(_255_v112, _255_v112)).length))) : (_114_v4));
          let _lhs41 = _107_globalState;
          _250_v107 = _rhs51;
          _168_v52 = _rhs52;
          _lhs41.f0 = _rhs53;
          let _256_v113;
          let _nw43 = Array((new BigNumber(19)).toNumber());
          _nw43[(_dafny.ZERO).toNumber()] = _116_v6;
          _nw43[(_dafny.ONE).toNumber()] = _116_v6;
          _nw43[(new BigNumber(2)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(3)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(4)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(5)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(6)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(7)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(8)).toNumber()] = (_module.D4.create_DC7(_116_v6)).dtor_cf14;
          _nw43[(new BigNumber(9)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(10)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(11)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(12)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(13)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(14)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(15)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(16)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(17)).toNumber()] = _116_v6;
          _nw43[(new BigNumber(18)).toNumber()] = _116_v6;
          _256_v113 = _nw43;
          _256_v113 = _256_v113;
          let _index62 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_156_v44).length));
          (_156_v44)[_index62] = (_244_v103).f8;
        } else {
          let _257_v114;
          _257_v114 = _157_v45;
          let _index63 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_156_v44).length));
          let _rhs54 = _156_v44;
          let _rhs55 = _244_v103;
          let _rhs56 = !((_115_v5).equals(_115_v5));
          let _rhs57 = _114_v4;
          let _rhs58 = (new BigNumber((_dafny.Seq.UnicodeFromString("mxtsyoj")).length)).minus((new BigNumber((_171_v55).length)).plus(new BigNumber(((_257_v114)).length)));
          let _lhs42 = _107_globalState;
          let _lhs43 = _107_globalState;
          let _lhs44 = _156_v44;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_156_v44).length));
          _156_v44 = _rhs54;
          _244_v103 = _rhs55;
          _lhs42.f0 = _rhs56;
          _lhs43.f0 = _rhs57;
          _lhs44[_lhs45] = _rhs58;
          let _258_v115;
          let _nw44 = new _module.C0();
          _nw44.__ctor((_244_v103).f7, _243_i5, _111_v2, (_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]);
          _258_v115 = _nw44;
          let _259_v116;
          _259_v116 = _module.D6.create_DC12(_154_v42);
          let _260_v117;
          _260_v117 = _dafny.MultiSet.fromElements(_114_v4, ((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]) && (_114_v4), (((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]) ? ((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]) : (_114_v4)), ((_259_v116).dtor_cf25).IsDisjointFrom(_154_v42), !(_dafny.Map.Empty.slice().updateUnsafe((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))],(_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))])).contains((_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))]));
          let _index64 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_156_v44).length));
          (_156_v44)[_index64] = new BigNumber((_260_v117).cardinality());
          let _261_v118;
          let _nw45 = Array((new BigNumber(8)).toNumber());
          _nw45[(_dafny.ZERO).toNumber()] = _116_v6;
          _nw45[(_dafny.ONE).toNumber()] = _116_v6;
          _nw45[(new BigNumber(2)).toNumber()] = _116_v6;
          _nw45[(new BigNumber(3)).toNumber()] = _116_v6;
          _nw45[(new BigNumber(4)).toNumber()] = _116_v6;
          _nw45[(new BigNumber(5)).toNumber()] = _116_v6;
          _nw45[(new BigNumber(6)).toNumber()] = _116_v6;
          _nw45[(new BigNumber(7)).toNumber()] = ((_114_v4) ? (_116_v6) : (_116_v6));
          _261_v118 = _nw45;
          let _index65 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_261_v118).length));
          (_261_v118)[_index65] = _116_v6;
          let _index66 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_261_v118).length));
          (_261_v118)[_index66] = _116_v6;
          let _262_v119;
          let _nw46 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
          _262_v119 = _nw46;
          let _263_v121;
          _263_v121 = _dafny.MultiSet.fromElements(_111_v2, (_244_v103).f7);
          let _264_v122;
          _264_v122 = _module.D6.create_DC15((_244_v103).f8, _258_v115, function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of (_263_v121).Elements) {
    let _265_v120 = _compr_4;
    if ((_263_v121).contains(_265_v120)) {
      _coll4.push([_265_v120,_108_v0]);
    }
  }
  return _coll4;
}(), _module.__default.fm1(new BigNumber((_260_v117).cardinality()), (_116_v6)[_module.__default.safeIndex(new BigNumber(75), new BigNumber((_116_v6).length))], _107_globalState));
          let _266_v123;
          _266_v123 = _module.D6.create_DC16(_264_v122);
          let _267_v124;
          _267_v124 = _dafny.Map.Empty.slice().updateUnsafe(_266_v123,new BigNumber(((_244_v103).f7).length));
          let _index67 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_262_v119).length));
          (_262_v119)[_index67] = (_267_v124).Merge(_267_v124);
          let _index68 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_262_v119).length));
          (_262_v119)[_index68] = (_267_v124).Merge((_267_v124).Merge(_267_v124));
        }
        let _index69 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_156_v44).length));
        (_156_v44)[_index69] = _243_i5;
        (_107_globalState).f0 = _114_v4;
      }
      process.stdout.write(_dafny.toString(_107_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_107_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_107_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_107_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_107_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_108_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(15)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_109_v1)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_111_v2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_112_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("eedslel")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_114_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_115_v5).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_v6)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_117_v7).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v41)[_dafny.ZERO]).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v41)[_dafny.ONE]).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v41)[new BigNumber(2)]).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v41)[new BigNumber(3)]).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v41)[new BigNumber(4)]).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v41)[new BigNumber(5)]).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v41)[new BigNumber(6)]).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v41)[new BigNumber(7)]).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v41)[new BigNumber(8)]).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v42).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_155_v43).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v44)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_157_v45).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_158_v46).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v47).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_160_v48).dtor_cf11).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_v50).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(524)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v51).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_168_v52));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v53).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v53).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v53).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v53).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_v54).equals(_dafny.Set.fromElements(_module.D1.create_DC3(false, _dafny.ONE, true, new _dafny.CodePoint('f'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_171_v55, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_241_v101).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_242_v102).length)));
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
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return false;
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
    static create_DC2(cf2, cf3, cf4, cf5, cf6) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC3(cf7, cf8, cf9, cf10) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(2);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC4(cf11) {
      let $dt = new D1(3);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC1() { return this.$tag === 2; }
    get is_DC4() { return this.$tag === 3; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + this.cf5.toVerbatimString(true) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.Seq.of(), _dafny.ZERO, _dafny.ZERO, _dafny.Seq.UnicodeFromString(""), _dafny.Set.Empty);
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
    static create_DC5(cf12) {
      let $dt = new D2(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12;
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
    static create_DC6(cf13) {
      let $dt = new D3(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf13) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC8(cf15, cf16, cf17, cf18, cf19) {
      let $dt = new D4(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC9(cf20, cf21, cf22) {
      let $dt = new D4(1);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC7(cf14) {
      let $dt = new D4(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC10(cf23) {
      let $dt = new D4(3);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get is_DC10() { return this.$tag === 3; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16 && this.cf17 === other.cf17 && this.cf18 === other.cf18 && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && this.cf22 === other.cf22;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf14 === other.cf14;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC8(_dafny.ZERO, null, false, false, null);
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
    static create_DC11(cf24) {
      let $dt = new D5(0);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC13(cf26, cf27) {
      let $dt = new D6(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC14(cf28, cf29) {
      let $dt = new D6(1);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC15(cf30, cf31, cf32, cf33) {
      let $dt = new D6(2);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC12(cf25) {
      let $dt = new D6(3);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC16(cf34) {
      let $dt = new D6(4);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get is_DC16() { return this.$tag === 4; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 4) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf28 === other.cf28 && this.cf29 === other.cf29;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32) && this.cf33 === other.cf33;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC13(false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
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
      this.f1 = _dafny.ZERO;
      this._f2 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f3 = _dafny.ZERO;
      this._f4 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f6 = false;
      this._f5 = _dafny.Seq.UnicodeFromString("");
      this._f7 = _dafny.Seq.UnicodeFromString("");
      this._f8 = _dafny.ZERO;
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
    __ctor(f7, f8, f5, f6) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_this).f8, ((_this.f6) ? ((_this).f8) : ((_this).f8)));
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_this.f6,_dafny.Seq.of(_this.f6)));
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
