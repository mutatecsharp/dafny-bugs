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
    static fm0(p0, globalState) {
      if (false) {
        return true;
      } else {
        return !(_dafny.Set.fromElements(!(false))).contains(false);
      }
    };
    static fm1(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), function (_0_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      });
    };
    static fm3(p0, p1, globalState) {
      return _module.D3.create_DC5(_module.D3.create_DC5(_module.D3.create_DC5(_module.D3.create_DC4(new BigNumber(-120), true, false))));
    };
    static fm4(p0, globalState) {
      return (function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Set.fromElements(new BigNumber(311), new BigNumber(461), new BigNumber((_dafny.Seq.of(new BigNumber((function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(512))).length))).Elements) {
            let _1_v0 = _compr_1;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(512))).length))).contains(_1_v0)) {
              _coll1.push([_module.__default.safeDivisionInt(_1_v0, new BigNumber((_dafny.Seq.UnicodeFromString("cpviudyfg")).length)),new BigNumber(446)]);
            }
          }
          return _coll1;
        }()).length))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-299),false)).length)))).Elements) {
          let _2_v1 = _compr_0;
          if ((_dafny.Set.fromElements(new BigNumber(311), new BigNumber(461), new BigNumber((_dafny.Seq.of(new BigNumber((function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(512))).length))).Elements) {
              let _1_v0 = _compr_2;
              if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(512))).length))).contains(_1_v0)) {
                _coll2.push([_module.__default.safeDivisionInt(_1_v0, new BigNumber((_dafny.Seq.UnicodeFromString("cpviudyfg")).length)),new BigNumber(446)]);
              }
            }
            return _coll2;
          }()).length))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-299),false)).length)))).contains(_2_v1)) {
            _coll0.add((_2_v1).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, false)).length))).length)));
          }
        }
        return _coll0;
      }()).Intersect((_dafny.Set.fromElements(new BigNumber(633), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),true)).length))).Intersect(function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(381), new BigNumber(334))) {
          let _3_v2 = _compr_3;
          if (((new BigNumber(381)).isLessThanOrEqualTo(_3_v2)) && ((_3_v2).isLessThan(new BigNumber(334)))) {
            _coll3.add((_3_v2).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(398),false),false)).length)));
          }
        }
        return _coll3;
      }()));
    };
    static fm5(globalState) {
      return (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true, !(false)))).length))).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(978))).length))));
    };
    static fm6(p0, p1, p2, globalState) {
      if ((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(15)), function (_4_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })).length)).isLessThan(new BigNumber(865))) {
        return _dafny.MultiSet.fromElements(new BigNumber(233));
      } else {
        return _dafny.MultiSet.fromElements(new BigNumber(289));
      }
    };
    static fm7(globalState) {
      return _dafny.Seq.UnicodeFromString("hkhqfcu");
    };
    static fm8(globalState) {
      return new _dafny.CodePoint('x'.codePointAt(0));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(821)), function (_5_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }))).Union(_dafny.MultiSet.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0))));
    };
    static fm10(globalState) {
      let _source0 = ((!(false)) ? (_module.D7.create_DC15(new BigNumber(407), false, new BigNumber((_dafny.Seq.of(false, false)).length))) : (_module.D7.create_DC15((_dafny.ZERO).minus(new BigNumber(-917)), false, new BigNumber(757))));
      if (_source0.is_DC14) {
        let _6___mcc_h0 = (_source0).cf21;
        let _7___mcc_h1 = (_source0).cf22;
        let _8___mcc_h2 = (_source0).cf23;
        let _9___mcc_h3 = (_source0).cf24;
        let _10___mcc_h4 = (_source0).cf25;
        let _11_cf25 = _10___mcc_h4;
        let _12_cf24 = _9___mcc_h3;
        let _13_cf23 = _8___mcc_h2;
        let _14_cf22 = _7___mcc_h1;
        let _15_cf21 = _6___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ksuka"),true);
      } else if (_source0.is_DC15) {
        let _16___mcc_h5 = (_source0).cf26;
        let _17___mcc_h6 = (_source0).cf27;
        let _18___mcc_h7 = (_source0).cf28;
        let _19_cf28 = _18___mcc_h7;
        let _20_cf27 = _17___mcc_h6;
        let _21_cf26 = _16___mcc_h5;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(183)), function (_22_i0) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }),_20_cf27)).Merge(function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("enjjy"),new _dafny.CodePoint('e'.codePointAt(0)))).Keys.Elements) {
            let _23_v0 = _compr_4;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("enjjy"),new _dafny.CodePoint('e'.codePointAt(0)))).contains(_23_v0)) {
              _coll4.push([_23_v0,_20_cf27]);
            }
          }
          return _coll4;
        }());
      } else {
        let _24___mcc_h8 = (_source0).cf20;
        let _25_cf20 = _24___mcc_h8;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("bxdetlja"),false);
      }
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return (true) && (false);
    };
    static fm12(p0, p1, p2, globalState) {
      return _dafny.Seq.of(!(!(true)) || (false), true);
    };
    static fm13(globalState) {
      let _source1 = ((false) ? (_module.D12.create_DC27(false, new BigNumber(154))) : (_module.D12.create_DC27(false, new BigNumber(85))));
      if (_source1.is_DC27) {
        let _26___mcc_h0 = (_source1).cf45;
        let _27___mcc_h1 = (_source1).cf46;
        let _28_cf46 = _27___mcc_h1;
        let _29_cf45 = _26___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_28_cf46), _dafny.Seq.of(_28_cf46));
      } else {
        let _30___mcc_h2 = (_source1).cf44;
        let _31_cf44 = _30___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality())), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-351)), function (_32_i0) {
          return new BigNumber(-346);
        }));
      }
    };
    static fm14(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(91)).multipliedBy(new BigNumber(417)),(new BigNumber((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(341),new BigNumber(-240))).length), new BigNumber((_dafny.Seq.UnicodeFromString("fwkthi")).length))).Elements) {
          let _33_v0 = _compr_5;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(341),new BigNumber(-240))).length), new BigNumber((_dafny.Seq.UnicodeFromString("fwkthi")).length)), _33_v0)) {
            _coll5.push([(_33_v0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-881)), function (_34_i0) {
              return new BigNumber(87);
            })).length)),new BigNumber(520)]);
          }
        }
        return _coll5;
      }()).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)))).length))).length)));
    };
    static fm15(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-59))).cardinality()))).Merge((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),new BigNumber(568))).Keys.Elements) {
          let _35_v0 = _compr_6;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),new BigNumber(568))).contains(_35_v0)) {
            _coll6.push([_35_v0,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))]);
          }
        }
        return _coll6;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),(_module.D8.create_DC17(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new BigNumber((_dafny.Seq.of(true)).length))).dtor_cf32)));
    };
    static fm16(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("hfqueonrf"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(258)), function (_36_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("yayviyc"))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(616)), function (_37_i1) {
        return _dafny.Seq.UnicodeFromString("ocnyb");
      }));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),(_dafny.ZERO).minus(new BigNumber((function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(594), new BigNumber(946))) {
          let _38_v0 = _compr_7;
          if (((new BigNumber(594)).isLessThanOrEqualTo(_38_v0)) && ((_38_v0).isLessThan(new BigNumber(946)))) {
            _coll7.push([_module.__default.safeModuloInt(_38_v0, new BigNumber(-345)),new BigNumber(331)]);
          }
        }
        return _coll7;
      }()).length))));
    };
    static fm18(p0, p1, p2, globalState) {
      return (_module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber(93))).length), new BigNumber(393))).multipliedBy(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_module.D11.create_DC25(new BigNumber((_dafny.Seq.of(false)).length), !(!(!(false))))).dtor_cf43,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).length));
    };
    static fm19(globalState) {
      let _source2 = _module.D8.create_DC18(_module.D8.create_DC18(_module.D8.create_DC17(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("ibkkhb")).length))));
      if (_source2.is_DC17) {
        let _39___mcc_h0 = (_source2).cf30;
        let _40___mcc_h1 = (_source2).cf31;
        let _41___mcc_h2 = (_source2).cf32;
        let _42_cf32 = _41___mcc_h2;
        let _43_cf31 = _40___mcc_h1;
        let _44_cf30 = _39___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(true), (_module.D16.create_DC34(_dafny.Seq.of(false))).dtor_cf57);
      } else if (_source2.is_DC16) {
        let _45___mcc_h3 = (_source2).cf29;
        let _46_cf29 = _45___mcc_h3;
        return _dafny.Seq.Concat(_dafny.Seq.of(false, true), _dafny.Seq.of((_module.D4.create_DC8(new _dafny.CodePoint('d'.codePointAt(0)), true, false)).dtor_cf14, false, !(true)));
      } else {
        let _47___mcc_h4 = (_source2).cf33;
        let _48_cf33 = _47___mcc_h4;
        return _dafny.Seq.of((_module.D7.create_DC15(new BigNumber(763), false, new BigNumber(866))).dtor_cf27, false, !((_module.D16.create_DC35(true, new BigNumber(348), new BigNumber((_dafny.Seq.UnicodeFromString("vbrr")).length))).dtor_cf58), false);
      }
    };
    static fm20(p0, globalState) {
      let _source3 = _module.D11.create_DC25((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(720), new BigNumber(-331), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), function (_49_i0) {
  return new _dafny.CodePoint('f'.codePointAt(0));
})).length), (_dafny.ZERO).minus(new BigNumber(-308)))).length)), true);
      if (_source3.is_DC25) {
        let _50___mcc_h0 = (_source3).cf42;
        let _51___mcc_h1 = (_source3).cf43;
        let _52_cf43 = _51___mcc_h1;
        let _53_cf42 = _50___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-190)), ((_54_cf42) => function (_55_i1) {
          return _54_cf42;
        })(_53_cf42)), _dafny.Seq.of(_53_cf42, _53_cf42, _53_cf42), _dafny.Seq.of(_53_cf42, _53_cf42, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_53_cf42)).length))), _53_cf42, _53_cf42))).length));
      } else {
        let _56___mcc_h2 = (_source3).cf41;
        let _57_cf41 = _56___mcc_h2;
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),new BigNumber(628));
      }
    };
    static fm21(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(true, false, true)).Intersect(_dafny.Set.fromElements(false, false))).Intersect((_dafny.Set.fromElements(false, true)).Union(_dafny.Set.fromElements(true)));
    };
    static fm22(p0, globalState) {
      return new _dafny.CodePoint('r'.codePointAt(0));
    };
    static fm23(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ybjet")).length)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)), new BigNumber((_dafny.Seq.UnicodeFromString("odkdqes")).length), new BigNumber(-633), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("ymjr")).length)))).length));
    };
    static fm24(p0, globalState) {
      if (!(_dafny.Map.Empty.slice().updateUnsafe(!(!(!(false))),new BigNumber(420))).contains(false)) {
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ggcg")).length))), _dafny.Set.fromElements(new BigNumber(836), new BigNumber(188))), _dafny.Seq.of(function () {
          let _coll8 = new _dafny.Set();
          for (const _compr_8 of _dafny.IntegerRange(new BigNumber(393), new BigNumber(306))) {
            let _58_v0 = _compr_8;
            if (((new BigNumber(393)).isLessThanOrEqualTo(_58_v0)) && ((_58_v0).isLessThan(new BigNumber(306)))) {
              _coll8.add((_58_v0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(716)), function (_59_i0) {
                return (_dafny.ZERO).minus(new BigNumber(-150));
              })).length)));
            }
          }
          return _coll8;
        }()));
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-234))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), function (_60_i1) {
          return _dafny.Set.fromElements(new BigNumber(512));
        }));
      }
    };
    static fm25(globalState) {
      return _module.D3.create_DC3(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("koggtmo")));
    };
    static fm26(globalState) {
      if (!(_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).equals(_dafny.MultiSet.fromElements(!(true)))) {
        if (true) {
          return _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll9 = new _dafny.Set();
            for (const _compr_9 of (function () {
              let _coll10 = new _dafny.Map();
              for (const _compr_10 of _dafny.IntegerRange(new BigNumber(571), new BigNumber(-350))) {
                let _61_v0 = _compr_10;
                if (((new BigNumber(571)).isLessThanOrEqualTo(_61_v0)) && ((_61_v0).isLessThan(new BigNumber(-350)))) {
                  _coll10.push([(_61_v0).multipliedBy(new BigNumber(278)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).length)]);
                }
              }
              return _coll10;
            }()).Keys.Elements) {
              let _62_v1 = _compr_9;
              if ((function () {
                let _coll11 = new _dafny.Map();
                for (const _compr_11 of _dafny.IntegerRange(new BigNumber(571), new BigNumber(-350))) {
                  let _61_v0 = _compr_11;
                  if (((new BigNumber(571)).isLessThanOrEqualTo(_61_v0)) && ((_61_v0).isLessThan(new BigNumber(-350)))) {
                    _coll11.push([(_61_v0).multipliedBy(new BigNumber(278)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).length)]);
                  }
                }
                return _coll11;
              }()).contains(_62_v1)) {
                _coll9.add(_module.__default.safeDivisionInt(_62_v1, new BigNumber((_dafny.Seq.of(new BigNumber(335))).length)));
              }
            }
            return _coll9;
          }(),_dafny.Seq.UnicodeFromString("esyx"));
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())),_dafny.Seq.Create(_module.__default.abs(new BigNumber(329)), function (_63_i0) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          }));
        }
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(142)),_dafny.Seq.UnicodeFromString("kraaoyc"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(957)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), function (_64_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        })));
      }
    };
    static fm27(p0, globalState) {
      return _module.D8.create_DC17(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), (_dafny.ZERO).minus(new BigNumber(-416)));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(false,!(true) || (true));
    };
    static fm29(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rknh"),new BigNumber(259));
    };
    static m0(globalState) {
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _65_v0;
      _65_v0 = new BigNumber(772);
      let _hi0 = new BigNumber(310);
      for (let _66_i0 = _65_v0; _66_i0.isLessThan(_hi0); _66_i0 = _66_i0.plus(_dafny.ONE)) {
        let _67_v1;
        _67_v1 = _dafny.Seq.UnicodeFromString("cviinusi");
        let _68_v2;
        _68_v2 = _dafny.MultiSet.fromElements(_65_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_65_v0,new BigNumber(374))).length), new BigNumber((_67_v1).length), _66_i0, (_dafny.ZERO).minus(_66_i0));
        _68_v2 = _68_v2;
        let _69_v3;
        _69_v3 = false;
        let _70_v4;
        _70_v4 = _69_v3;
        let _71_v5;
        let _nw0 = new _module.C0();
        _nw0.__ctor(_69_v3);
        _71_v5 = _nw0;
        let _72_v6;
        _72_v6 = _dafny.Map.Empty.slice().updateUnsafe(_71_v5,_69_v3);
        let _73_v7;
        _73_v7 = _dafny.MultiSet.fromElements(_71_v5);
        let _rhs0 = _70_v4;
        let _rhs1 = ((_dafny.ZERO).minus((_65_v0).multipliedBy(new BigNumber((_72_v6).length)))).minus(_66_i0);
        let _rhs2 = _69_v3;
        let _rhs3 = _69_v3;
        let _rhs4 = (((_73_v7).contains(_71_v5)) ? ((_73_v7).get(_71_v5)) : (_module.__default.safeDivisionInt(_66_i0, _66_i0)));
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        let _lhs2 = globalState;
        _70_v4 = _rhs0;
        r0 = _rhs1;
        _lhs0.f5 = _rhs2;
        _lhs1.f5 = _rhs3;
        _lhs2.f8 = _rhs4;
        r1 = _69_v3;
        (globalState).f8 = _65_v0;
      }
      let _74_v8;
      _74_v8 = true;
      (globalState).f17 = _74_v8;
      let _75_v9;
      let _nw1 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _75_v9 = _nw1;
      let _index0 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_75_v9).length));
      (_75_v9)[_index0] = _65_v0;
      let _76_v10;
      _76_v10 = _dafny.Seq.of(_74_v8, false);
      let _77_v11;
      _77_v11 = _dafny.MultiSet.fromElements(_74_v8, _74_v8, _module.__default.fm0(_76_v10, globalState), !(_74_v8), _74_v8);
      let _78_v12;
      let _nw2 = new _module.C0();
      _nw2.__ctor(_74_v8);
      _78_v12 = _nw2;
      let _79_v13;
      _79_v13 = new _dafny.CodePoint('c'.codePointAt(0));
      let _80_v14;
      _80_v14 = _dafny.Map.Empty.slice().updateUnsafe(_78_v12,_79_v13);
      let _index1 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_75_v9).length));
      (_75_v9)[_index1] = ((((_77_v11).contains(_74_v8)) ? ((_77_v11).get(_74_v8)) : (_65_v0))).plus(new BigNumber((_80_v14).length));
      let _81_v15;
      let _nw3 = Array((new BigNumber(10)).toNumber()).fill(false);
      _81_v15 = _nw3;
      let _82_v16;
      _82_v16 = _dafny.Seq.UnicodeFromString("xnm");
      let _83_v17;
      _83_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_65_v0),_82_v16);
      let _84_v18;
      let _nw4 = new _module.C4();
      _nw4.__ctor(_65_v0);
      _84_v18 = _nw4;
      let _85_v19;
      let _nw5 = new _module.C1();
      _nw5.__ctor(_83_v17, _84_v18);
      _85_v19 = _nw5;
      let _86_v20;
      _86_v20 = _dafny.Set.fromElements(_74_v8);
      let _rhs5 = _81_v15;
      let _rhs6 = _85_v19;
      let _rhs7 = _module.__default.safeModuloInt((new BigNumber((_86_v20).length)).minus(new BigNumber((_module.__default.fm13(globalState)).length)), _module.__default.safeDivisionInt(_65_v0, (_75_v9)[_module.__default.safeIndex(new BigNumber(482), new BigNumber((_75_v9).length))]));
      let _rhs8 = ((new BigNumber(818)).multipliedBy(_65_v0)).plus(_65_v0);
      let _lhs3 = globalState;
      _81_v15 = _rhs5;
      _85_v19 = _rhs6;
      _lhs3.f8 = _rhs7;
      _65_v0 = _rhs8;
      let _87_v21;
      let _nw6 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _87_v21 = _nw6;
      let _index2 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_87_v21).length));
      (_87_v21)[_index2] = _module.__default.fm1(globalState);
      let _index3 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_87_v21).length));
      (_87_v21)[_index3] = _82_v16;
      let _88_v22;
      let _nw7 = new _module.C3();
      _nw7.__ctor(_75_v9);
      _88_v22 = _nw7;
      r0 = (_65_v0).multipliedBy(new BigNumber(-651));
      let _89_v23;
      _89_v23 = _module.D8.create_DC16((_87_v21)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_87_v21).length))]);
      r1 = _dafny.Seq.IsPrefixOf((_89_v23).dtor_cf29, _82_v16);
      r2 = true;
      r3 = (_65_v0).minus((_75_v9)[_module.__default.safeIndex(new BigNumber(482), new BigNumber((_75_v9).length))]);
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _90_v0;
      _90_v0 = _dafny.Seq.UnicodeFromString("wifcawqh");
      let _91_v1;
      let _init0 = ((_92_v0) => function (_93_i0) {
        return _module.__default.safeDivisionInt(_93_i0, new BigNumber((_92_v0).length));
      })(_90_v0);
      let _nw8 = Array((new BigNumber(16)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw8.length); _i0_0++) {
        _nw8[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _91_v1 = _nw8;
      let _94_v2;
      let _nw9 = Array((new BigNumber(3)).toNumber());
      _nw9[(_dafny.ZERO).toNumber()] = _91_v1;
      _nw9[(_dafny.ONE).toNumber()] = _91_v1;
      _nw9[(new BigNumber(2)).toNumber()] = _91_v1;
      _94_v2 = _nw9;
      let _95_v3;
      _95_v3 = false;
      let _96_v4;
      _96_v4 = _dafny.Map.Empty.slice().updateUnsafe(_95_v3,false);
      let _97_v5;
      _97_v5 = _dafny.Set.fromElements(false);
      let _98_v6;
      let _nw10 = Array((_dafny.ONE).toNumber());
      _nw10[(_dafny.ZERO).toNumber()] = _97_v5;
      _98_v6 = _nw10;
      let _99_v7;
      _99_v7 = new BigNumber(637);
      let _100_v8;
      _100_v8 = _95_v3;
      let _101_v9;
      _101_v9 = _dafny.Map.Empty.slice().updateUnsafe(_99_v7,(_100_v8));
      let _102_v10;
      let _nw11 = Array((new BigNumber(4)).toNumber());
      _nw11[(_dafny.ZERO).toNumber()] = _90_v0;
      _nw11[(_dafny.ONE).toNumber()] = _90_v0;
      _nw11[(new BigNumber(2)).toNumber()] = _90_v0;
      _nw11[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(499)), function (_103_i1) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      });
      _102_v10 = _nw11;
      let _104_v11;
      _104_v11 = _dafny.Map.Empty.slice().updateUnsafe((_100_v8),_dafny.Seq.UnicodeFromString("dwmjjp"));
      let _105_globalState;
      let _nw12 = new _module.GlobalState();
      _nw12.__ctor(_90_v0, _94_v2, (_96_v4).Merge(_96_v4), new _dafny.CodePoint('i'.codePointAt(0)), new BigNumber(992), false, new BigNumber(-915), _98_v6, new BigNumber(504), _101_v9, _90_v0, true, _102_v10, _94_v2, new BigNumber(34), new BigNumber(176), _104_v11, true, new BigNumber(649), true, false, false, true, _91_v1, _dafny.Seq.UnicodeFromString("l"));
      _105_globalState = _nw12;
      let _106_v12;
      _106_v12 = _dafny.Seq.of(_99_v7);
      let _107_v13;
      _107_v13 = _dafny.Seq.of(_95_v3);
      let _108_v14;
      _108_v14 = _dafny.Map.Empty.slice().updateUnsafe(_106_v12,_module.__default.fm0(_107_v13, _105_globalState));
      _108_v14 = (_108_v14).Merge((_108_v14).Merge(_108_v14));
      let _109_v15;
      _109_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(40)), function (_110_i3) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }),_95_v3);
      let _111_i2;
      _111_i2 = _dafny.ZERO;
      L0: {
        while ((_109_v15).contains(_dafny.Seq.Concat(_90_v0, _90_v0))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_111_i2)) {
              break L0;
            }
            _111_i2 = (_111_i2).plus(_dafny.ONE);
            (_105_globalState).f22 = !(_95_v3);
            _99_v7 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("puigmlj"), _dafny.Seq.UnicodeFromString("e"))).length);
            _109_v15 = (_109_v15).update(_90_v0, _95_v3);
            let _112_v16;
            _112_v16 = _dafny.Map.Empty.slice().updateUnsafe(_100_v8,_95_v3);
            let _113_v17;
            _113_v17 = _dafny.Map.Empty.slice().updateUnsafe(_91_v1,_112_v16);
            _112_v16 = (((_113_v17).contains(_91_v1)) ? ((_113_v17).get(_91_v1)) : (_112_v16));
          }
        }
      }
      let _114_v18;
      let _115_v19;
      let _116_v20;
      let _117_v21;
      let _out0;
      let _out1;
      let _out2;
      let _out3;
      let _outcollector0 = _module.__default.m0(_105_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _out3 = _outcollector0[3];
      _114_v18 = _out0;
      _115_v19 = _out1;
      _116_v20 = _out2;
      _117_v21 = _out3;
      let _118_v22;
      let _nw13 = new _module.C0();
      _nw13.__ctor(_95_v3);
      _118_v22 = _nw13;
      _95_v3 = _115_v19;
      let _119_v23;
      let _nw14 = new _module.C3();
      _nw14.__ctor(_91_v1);
      _119_v23 = _nw14;
      let _120_v24;
      _120_v24 = _module.D13.create_DC29(_119_v23);
      let _121_v25;
      _121_v25 = _module.D13.create_DC30(_120_v24);
      let _source4 = _121_v25;
      if (_source4.is_DC29) {
        let _122___mcc_h0 = (_source4).cf48;
        let _123_cf48 = _122___mcc_h0;
        let _124_v26;
        let _125_v27;
        let _126_v28;
        let _127_v29;
        let _out4;
        let _out5;
        let _out6;
        let _out7;
        let _outcollector1 = _module.__default.m0(_105_globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _out6 = _outcollector1[2];
        _out7 = _outcollector1[3];
        _124_v26 = _out4;
        _125_v27 = _out5;
        _126_v28 = _out6;
        _127_v29 = _out7;
        let _128_v30;
        _128_v30 = _dafny.Set.fromElements(_114_v18);
        _117_v21 = new BigNumber((_128_v30).length);
        let _129_v31;
        let _nw15 = new _module.C0();
        _nw15.__ctor(!(_116_v20));
        _129_v31 = _nw15;
        let _130_v32;
        let _nw16 = new _module.C1();
        _nw16.__ctor(_module.__default.fm26(_105_globalState), _129_v31);
        _130_v32 = _nw16;
        let _nw17 = new _module.C1();
        _nw17.__ctor((_130_v32).f28, (_130_v32).f29);
        _130_v32 = _nw17;
        let _131_v33;
        _131_v33 = _module.D6.create_DC10(_91_v1);
        let _132_v34;
        let _nw18 = new _module.C3();
        _nw18.__ctor((_131_v33).dtor_cf17);
        _132_v34 = _nw18;
      } else if (_source4.is_DC28) {
        let _133___mcc_h1 = (_source4).cf47;
        let _134_cf47 = _133___mcc_h1;
        (_105_globalState).f17 = _116_v20;
        let _135_v35;
        _135_v35 = _dafny.MultiSet.fromElements(_116_v20);
        let _136_v37;
        _136_v37 = _dafny.Map.Empty.slice().updateUnsafe(_100_v8,_116_v20);
        if ((function () {
          let _coll12 = new _dafny.Set();
          for (const _compr_12 of (_135_v35).Elements) {
            let _137_v36 = _compr_12;
            if ((_135_v35).contains(_137_v36)) {
              _coll12.add(_137_v36);
            }
          }
          return _coll12;
        }()).IsDisjointFrom(function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of (_136_v37).Keys.Elements) {
            let _138_v38 = _compr_13;
            if ((_136_v37).contains(_138_v38)) {
              _coll13.add(_138_v38);
            }
          }
          return _coll13;
        }())) {
          let _139_v39;
          _139_v39 = _dafny.MultiSet.fromElements(false, (_118_v22).f27);
          let _140_v40;
          _140_v40 = _dafny.Map.Empty.slice().updateUnsafe(_91_v1,_module.__default.fm18(_139_v39, _module.__default.fm27(_114_v18, _105_globalState), (_118_v22).f27, _105_globalState));
          let _141_v41;
          _141_v41 = _dafny.Seq.of(_119_v23, _119_v23);
          let _rhs9 = _140_v40;
          let _rhs10 = new BigNumber((_dafny.Seq.Concat(_141_v41, _141_v41)).length);
          let _rhs11 = _117_v21;
          let _lhs4 = _105_globalState;
          let _lhs5 = _105_globalState;
          _140_v40 = _rhs9;
          _lhs4.f8 = _rhs10;
          _lhs5.f8 = _rhs11;
          let _rhs12 = _115_v19;
          let _rhs13 = _116_v20;
          let _lhs6 = _105_globalState;
          let _lhs7 = _105_globalState;
          _lhs6.f19 = _rhs12;
          _lhs7.f22 = _rhs13;
          (_119_v23).f26 = _91_v1;
          _117_v21 = (_99_v7).multipliedBy(_99_v7);
          let _arr0 = _119_v23.f26;
          let _index4 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_119_v23.f26).length));
          _arr0[_index4] = _module.__default.safeDivisionInt(_117_v21, _117_v21);
          let _arr1 = _119_v23.f26;
          let _index5 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_119_v23.f26).length));
          let _rhs14 = _114_v18;
          let _rhs15 = (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(296)), function (_142_i4) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })).length)).plus(_117_v21));
          let _lhs8 = _119_v23.f26;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_119_v23.f26).length));
          _114_v18 = _rhs14;
          _lhs8[_lhs9] = _rhs15;
        } else {
          _99_v7 = _114_v18;
          (_105_globalState).f22 = true;
          let _143_v42;
          _143_v42 = _dafny.MultiSet.fromElements((_118_v22).f27);
          _99_v7 = _module.__default.safeDivisionInt(new BigNumber((_143_v42).cardinality()), _117_v21);
          let _arr2 = _119_v23.f26;
          let _index6 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_119_v23.f26).length));
          _arr2[_index6] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_99_v7), _114_v18);
          let _144_v43;
          _144_v43 = _module.D11.create_DC24(_dafny.MultiSet.FromArray(_107_v13));
          let _145_v44;
          _145_v44 = _dafny.Map.Empty.slice().updateUnsafe(_115_v19,_114_v18);
          let _146_v45;
          _146_v45 = _dafny.Map.Empty.slice().updateUnsafe(_144_v43,_145_v44);
          let _arr3 = _119_v23.f26;
          let _index7 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_119_v23.f26).length));
          _arr3[_index7] = new BigNumber(((((_146_v45).contains(_144_v43)) ? ((_146_v45).get(_144_v43)) : ((_145_v44).Merge(_145_v44)))).length);
          let _147_v46;
          _147_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_90_v0).length),_91_v1);
          let _148_v47;
          _148_v47 = _dafny.Seq.of(_91_v1);
          (_105_globalState).f23 = (((_147_v46).contains(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_118_v22).f27,(_118_v22).f27)).length)), _module.__default.safeIndex(new BigNumber((_107_v13).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_118_v22).f27,(_118_v22).f27)).length))).length)), _99_v7)).length))) ? ((_147_v46).get(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_118_v22).f27,(_118_v22).f27)).length)), _module.__default.safeIndex(new BigNumber((_107_v13).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_118_v22).f27,(_118_v22).f27)).length))).length)), _99_v7)).length))) : ((_148_v47)[_module.__default.safeIndex(_117_v21, new BigNumber((_148_v47).length))]));
        }
        (_105_globalState).f17 = !(_116_v20);
        (_119_v23).m2(_105_globalState);
      } else {
        let _149___mcc_h2 = (_source4).cf49;
        let _150_cf49 = _149___mcc_h2;
        let _151_v48;
        let _nw19 = new _module.C0();
        _nw19.__ctor((_99_v7).isLessThanOrEqualTo(_117_v21));
        _151_v48 = _nw19;
        _151_v48 = (((_118_v22).f27) ? (_151_v48) : (_151_v48));
        _90_v0 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), function (_152_i5) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }), _90_v0);
        (_119_v23).m1(_105_globalState);
        let _153_v49;
        let _nw20 = Array((new BigNumber(16)).toNumber()).fill([]);
        _153_v49 = _nw20;
        let _index8 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_153_v49).length));
        (_153_v49)[_index8] = _94_v2;
        let _index9 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_153_v49).length));
        (_153_v49)[_index9] = _94_v2;
      }
      let _154_v50;
      let _nw21 = new _module.C0();
      _nw21.__ctor((_107_v13)[_module.__default.safeIndex((_dafny.ZERO).minus(_99_v7), new BigNumber((_107_v13).length))]);
      _154_v50 = _nw21;
      let _155_v51;
      _155_v51 = _module.D11.create_DC24(_dafny.MultiSet.fromElements(_115_v19));
      let _156_v52;
      _156_v52 = _dafny.Map.Empty.slice().updateUnsafe(_99_v7,_155_v51);
      let _157_v53;
      _157_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_156_v52).length),new BigNumber(375));
      _117_v21 = (_117_v21).plus(((((_157_v53).contains(_117_v21)) ? ((_157_v53).get(_117_v21)) : (_99_v7))).plus(_114_v18));
      let _158_v54;
      let _nw22 = new _module.C4();
      _nw22.__ctor(new BigNumber(849));
      _158_v54 = _nw22;
      _158_v54 = _158_v54;
      let _159_v55;
      let _nw23 = new _module.C0();
      _nw23.__ctor(((_154_v50).f27) === (_116_v20));
      _159_v55 = _nw23;
      let _160_v56;
      _160_v56 = _dafny.Map.Empty.slice().updateUnsafe(_107_v13,false);
      let _161_v58;
      _161_v58 = _dafny.Seq.of(_107_v13, _107_v13, _107_v13);
      _160_v56 = (function () {
        let _coll14 = new _dafny.Map();
        for (const _compr_14 of (_161_v58).Elements) {
          let _162_v57 = _compr_14;
          if (_dafny.Seq.contains(_161_v58, _162_v57)) {
            _coll14.push([_162_v57,!((_154_v50).f27)]);
          }
        }
        return _coll14;
      }()).update(_dafny.Seq.Concat(_107_v13, _107_v13), (_118_v22).f27);
      (_158_v54).m2(_105_globalState);
      let _163_v59;
      let _init1 = function (_164_i6) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      };
      let _nw24 = Array((new BigNumber(18)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw24.length); _i0_1++) {
        _nw24[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _163_v59 = _nw24;
      let _165_v60;
      _165_v60 = _module.D8.create_DC16(_90_v0);
      let _166_v61;
      _166_v61 = _module.D8.create_DC18(_165_v60);
      let _167_v62;
      _167_v62 = _dafny.Seq.of(_163_v59);
      let _pat_let_tv0 = _100_v8;
      let _pat_let_tv1 = _100_v8;
      let _pat_let_tv2 = _154_v50;
      let _rhs16 = function (_source5) {
        if (_source5.is_DC17) {
          let _168___mcc_h3 = (_source5).cf30;
          let _169___mcc_h4 = (_source5).cf31;
          let _170___mcc_h5 = (_source5).cf32;
          let _171_cf32 = _170___mcc_h5;
          let _172_cf31 = _169___mcc_h4;
          let _173_cf30 = _168___mcc_h3;
          return _pat_let_tv0;
        } else if (_source5.is_DC16) {
          let _174___mcc_h6 = (_source5).cf29;
          let _175_cf29 = _174___mcc_h6;
          return _pat_let_tv1;
        } else {
          let _176___mcc_h7 = (_source5).cf33;
          let _177_cf33 = _176___mcc_h7;
          return (_pat_let_tv2).f27;
        }
      }(_166_v61);
      let _rhs17 = _99_v7;
      let _rhs18 = (_167_v62)[_module.__default.safeIndex(_module.__default.safeModuloInt((_158_v54).f25, new BigNumber(-610)), new BigNumber((_167_v62).length))];
      let _lhs10 = _105_globalState;
      _100_v8 = _rhs16;
      _lhs10.f8 = _rhs17;
      _163_v59 = _rhs18;
      if ((_154_v50).f27) {
        let _178_v63;
        _178_v63 = new _dafny.CodePoint('a'.codePointAt(0));
        let _index10 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_163_v59).length));
        (_163_v59)[_index10] = _178_v63;
        let _index11 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_163_v59).length));
        let _rhs19 = (_dafny.ZERO).minus((_117_v21).minus((_158_v54).f25));
        let _rhs20 = _178_v63;
        let _rhs21 = (_106_v12)[_module.__default.safeIndex((_158_v54).f25, new BigNumber((_106_v12).length))];
        let _lhs11 = _105_globalState;
        let _lhs12 = _163_v59;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_163_v59).length));
        _lhs11.f8 = _rhs19;
        _lhs12[_lhs13] = _rhs20;
        _99_v7 = _rhs21;
        let _179_v64;
        _179_v64 = _dafny.MultiSet.fromElements(!((_118_v22).f27));
        let _180_v65;
        _180_v65 = _module.D8.create_DC17((_163_v59)[_module.__default.safeIndex(new BigNumber(255), new BigNumber((_163_v59).length))], (_163_v59)[_module.__default.safeIndex(new BigNumber(255), new BigNumber((_163_v59).length))], (_158_v54).f25);
        _114_v18 = _module.__default.safeModuloInt((_module.__default.fm18(_179_v64, _180_v65, _95_v3, _105_globalState)).plus((_158_v54).f25), (_dafny.ZERO).minus(_114_v18));
        let _index12 = _module.__default.safeIndex(new BigNumber(255), new BigNumber((_163_v59).length));
        (_163_v59)[_index12] = _module.__default.fm8(_105_globalState);
        let _181_v66;
        let _init2 = function (_182_i7) {
          return false;
        };
        let _nw25 = Array((new BigNumber(8)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw25.length); _i0_2++) {
          _nw25[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _181_v66 = _nw25;
        let _index13 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_181_v66).length));
        (_181_v66)[_index13] = _95_v3;
        let _index14 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_181_v66).length));
        (_181_v66)[_index14] = !(_157_v53).contains((_158_v54).f25);
        _95_v3 = _115_v19;
      } else {
        let _183_v67;
        let _init3 = ((_184_v55) => function (_185_i8) {
          return (_184_v55).f27;
        })(_159_v55);
        let _nw26 = Array((new BigNumber(19)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw26.length); _i0_3++) {
          _nw26[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _183_v67 = _nw26;
        let _186_v68;
        _186_v68 = _dafny.MultiSet.fromElements(_114_v18);
        let _index15 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_183_v67).length));
        (_183_v67)[_index15] = (_186_v68).IsDisjointFrom(_186_v68);
        let _187_v69;
        _187_v69 = _dafny.Seq.of(_90_v0);
        let _index16 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_183_v67).length));
        (_183_v67)[_index16] = _dafny.Seq.IsProperPrefixOf((((_118_v22).f27) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(8)), function (_188_i9) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        })) : ((_187_v69)[_module.__default.safeIndex((_158_v54).f25, new BigNumber((_187_v69).length))])), _90_v0);
        _106_v12 = _106_v12;
        let _189_v70;
        _189_v70 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_157_v53).length)));
        _189_v70 = _189_v70;
        let _190_v71;
        let _nw27 = new _module.C0();
        _nw27.__ctor((_dafny.Map.Empty.slice().updateUnsafe(_117_v21,true)).contains(_module.__default.fm5(_105_globalState)));
        _190_v71 = _nw27;
        (_105_globalState).f5 = !(!(((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), function (_191_i10) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        })).length)).multipliedBy(_117_v21)).isLessThan((_dafny.ZERO).minus(((_158_v54).f25).plus(_99_v7)))));
      }
      let _192_v72;
      let _nw28 = Array((new BigNumber(6)).toNumber());
      _nw28[(_dafny.ZERO).toNumber()] = !(_116_v20) || ((_154_v50).f27);
      _nw28[(_dafny.ONE).toNumber()] = !(_module.__default.fm0(_107_v13, _105_globalState));
      _nw28[(new BigNumber(2)).toNumber()] = false;
      _nw28[(new BigNumber(3)).toNumber()] = true;
      _nw28[(new BigNumber(4)).toNumber()] = _115_v19;
      _nw28[(new BigNumber(5)).toNumber()] = _95_v3;
      _192_v72 = _nw28;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_192_v72).length))) {
        let _193_i11 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_193_i11)) && ((_193_i11).isLessThan(new BigNumber((_192_v72).length))))) {
          (_192_v72)[(_193_i11)] = (_118_v22).f27;
        }
      }
      (_105_globalState).f22 = (_159_v55).f27;
      process.stdout.write((_90_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[_dafny.ONE])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_94_v2)[new BigNumber(2)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_95_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_v5).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_98_v6)[_dafny.ZERO]).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_99_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v8)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(637),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_102_v10)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_102_v10)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_102_v10)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_102_v10)[new BigNumber(3)], _dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_104_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("dwmjjp")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_105_globalState).f0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[_dafny.ONE])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f1)[new BigNumber(2)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f2).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f7)[_dafny.ZERO]).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(637),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_105_globalState).f10).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_105_globalState).f12)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_105_globalState).f12)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_105_globalState).f12)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_105_globalState).f12)[new BigNumber(3)], _dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[_dafny.ONE])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_105_globalState.f13)[new BigNumber(2)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f16).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("dwmjjp")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_globalState.f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_globalState.f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_globalState.f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f23)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f23)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f23)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f23)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState.f23)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_105_globalState).f24).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_106_v12, _dafny.Seq.of(new BigNumber(637)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_107_v13, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_108_v14).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(637)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_109_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0))),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_111_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_114_v18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_115_v19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_116_v20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_117_v21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_v22).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v23.f26)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_120_v24).dtor_cf48.f26)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_121_v25).dtor_cf49).dtor_cf48.f26)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v50).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_155_v51).dtor_cf41).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v52).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(637),_module.D11.create_DC24(_dafny.MultiSet.fromElements(true))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_157_v53).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(375)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v54).f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v55).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v56).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),true).updateUnsafe(_dafny.Seq.of(false, false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_161_v58, _dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(false), _dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_163_v59)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_165_v60).dtor_cf29).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_166_v61).dtor_cf33).dtor_cf29).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_167_v62).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v72)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v72)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v72)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v72)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v72)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_192_v72)[new BigNumber(5)]));
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
    static create_DC1(cf1) {
      let $dt = new D1(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
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
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2(cf2) {
      let $dt = new D2(0);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf2 === other.cf2;
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
    static create_DC4(cf4, cf5, cf6) {
      let $dt = new D3(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D3(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC5(cf7) {
      let $dt = new D3(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC4" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5 && this.cf6 === other.cf6;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC4(_dafny.ZERO, false, false);
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
    static create_DC7(cf9, cf10, cf11, cf12) {
      let $dt = new D4(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC8(cf13, cf14, cf15) {
      let $dt = new D4(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC6(cf8) {
      let $dt = new D4(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC6" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14 && this.cf15 === other.cf15;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC7(_dafny.ZERO, null, false, _dafny.ZERO);
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
    static create_DC9(cf16) {
      let $dt = new D5(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC9" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf16 === other.cf16;
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf18) {
      let $dt = new D6(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC10(cf17) {
      let $dt = new D6(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC12(cf19) {
      let $dt = new D6(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC10" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf18 === other.cf18;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf17 === other.cf17;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC11(false);
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
    static create_DC14(cf21, cf22, cf23, cf24, cf25) {
      let $dt = new D7(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC15(cf26, cf27, cf28) {
      let $dt = new D7(1);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC13(cf20) {
      let $dt = new D7(2);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC13" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21) && this.cf22 === other.cf22 && this.cf23 === other.cf23 && this.cf24 === other.cf24 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf26, other.cf26) && this.cf27 === other.cf27 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC14(_dafny.ZERO, false, false, false, _dafny.ZERO);
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
    static create_DC17(cf30, cf31, cf32) {
      let $dt = new D8(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC16(cf29) {
      let $dt = new D8(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC18(cf33) {
      let $dt = new D8(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC16" + "(" + this.cf29.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC17(new _dafny.CodePoint('D'.codePointAt(0)), new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC20(cf35, cf36) {
      let $dt = new D9(0);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC19(cf34) {
      let $dt = new D9(1);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf35 === other.cf35 && this.cf36 === other.cf36;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC20(false, false);
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
    static create_DC22(cf38, cf39) {
      let $dt = new D10(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC23(cf40) {
      let $dt = new D10(1);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC21(cf37) {
      let $dt = new D10(2);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC22" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38) && this.cf39 === other.cf39;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC22(_dafny.ZERO, false);
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
    static create_DC25(cf42, cf43) {
      let $dt = new D11(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC24(cf41) {
      let $dt = new D11(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42) && this.cf43 === other.cf43;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC25(_dafny.ZERO, false);
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
    static create_DC27(cf45, cf46) {
      let $dt = new D12(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC26(cf44) {
      let $dt = new D12(1);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC26" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC27(false, _dafny.ZERO);
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
    static create_DC29(cf48) {
      let $dt = new D13(0);
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC28(cf47) {
      let $dt = new D13(1);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC30(cf49) {
      let $dt = new D13(2);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC28" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf48 === other.cf48;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC29(null);
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
    static create_DC31(cf50) {
      let $dt = new D14(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50);
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
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC33(cf52, cf53, cf54, cf55, cf56) {
      let $dt = new D15(0);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC32(cf51) {
      let $dt = new D15(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC33" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC32" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf52 === other.cf52 && this.cf53 === other.cf53 && _dafny.areEqual(this.cf54, other.cf54) && this.cf55 === other.cf55 && this.cf56 === other.cf56;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC33(false, false, _dafny.ZERO, null, false);
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
    static create_DC35(cf58, cf59, cf60) {
      let $dt = new D16(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC34(cf57) {
      let $dt = new D16(1);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC35" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC34" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59) && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC35(false, _dafny.ZERO, _dafny.ZERO);
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
      this.f1 = [];
      this.f2 = _dafny.Map.Empty;
      this.f5 = false;
      this.f7 = [];
      this.f8 = _dafny.ZERO;
      this.f13 = [];
      this.f16 = _dafny.Map.Empty;
      this.f17 = false;
      this.f19 = false;
      this.f22 = false;
      this.f23 = [];
      this._f0 = _dafny.Seq.UnicodeFromString("");
      this._f3 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f4 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f9 = _dafny.Map.Empty;
      this._f10 = _dafny.Seq.UnicodeFromString("");
      this._f11 = false;
      this._f12 = [];
      this._f14 = _dafny.ZERO;
      this._f15 = _dafny.ZERO;
      this._f18 = _dafny.ZERO;
      this._f20 = false;
      this._f21 = false;
      this._f24 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this).f16 = f16;
      (_this).f17 = f17;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this).f22 = f22;
      (_this).f23 = f23;
      (_this)._f24 = f24;
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
    get f6() {
      let _this = this;
      return _this._f6;
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
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f27 = false;
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor(f27) {
      let _this = this;
      (_this)._f27 = f27;
      return;
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f28 = _dafny.Map.Empty;
      this._f29 = undefined;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f28, f29) {
      let _this = this;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    m1(globalState) {
      let _this = this;
      let _194_v0;
      _194_v0 = new BigNumber(141);
      (globalState).f8 = _194_v0;
      let _195_v1;
      _195_v1 = false;
      let _196_v2;
      _196_v2 = _dafny.Seq.of(_195_v1);
      let _197_v3;
      let _nw29 = Array((new BigNumber(2)).toNumber()).fill(false);
      _197_v3 = _nw29;
      let _198_v4;
      _198_v4 = _dafny.Map.Empty.slice().updateUnsafe(_194_v0,_197_v3);
      let _199_v5;
      _199_v5 = _dafny.MultiSet.fromElements(_195_v1, (_196_v2)[_module.__default.safeIndex(new BigNumber((_198_v4).length), new BigNumber((_196_v2).length))]);
      let _200_v6;
      _200_v6 = new _dafny.CodePoint('f'.codePointAt(0));
      let _201_v7;
      _201_v7 = _module.D8.create_DC17(_200_v6, _200_v6, _194_v0);
      let _hi1 = new BigNumber(937);
      for (let _202_i0 = _module.__default.fm18(_199_v5, _201_v7, _195_v1, globalState); _202_i0.isLessThan(_hi1); _202_i0 = _202_i0.plus(_dafny.ONE)) {
        let _203_v8;
        _203_v8 = _202_i0;
        let _204_v9;
        _204_v9 = _dafny.Map.Empty.slice().updateUnsafe(_194_v0,_203_v8);
        let _205_v10;
        _205_v10 = _dafny.Seq.of(_194_v0, _202_i0, new BigNumber((_204_v9).length), _202_i0, _194_v0);
        let _206_v11;
        let _nw30 = Array((new BigNumber(8)).toNumber());
        _nw30[(_dafny.ZERO).toNumber()] = _202_i0;
        _nw30[(_dafny.ONE).toNumber()] = _194_v0;
        _nw30[(new BigNumber(2)).toNumber()] = new BigNumber((_205_v10).length);
        _nw30[(new BigNumber(3)).toNumber()] = _202_i0;
        _nw30[(new BigNumber(4)).toNumber()] = _202_i0;
        _nw30[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_195_v1,!(_module.__default.fm0(_module.__default.fm19(globalState), globalState)))).length);
        _nw30[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_194_v0);
        _nw30[(new BigNumber(7)).toNumber()] = _194_v0;
        _206_v11 = _nw30;
        let _207_v12;
        _207_v12 = _dafny.Map.Empty.slice().updateUnsafe(!(_195_v1) || (_195_v1),_206_v11);
        _207_v12 = (_207_v12).update(_195_v1, _206_v11);
        (globalState).f8 = _module.__default.safeModuloInt(_194_v0, _194_v0);
        (globalState).f8 = ((_202_i0).minus(_194_v0)).minus((_202_i0).multipliedBy(_202_i0));
        (globalState).f19 = _195_v1;
      }
      _194_v0 = _194_v0;
      (globalState).f8 = _194_v0;
      let _208_v13;
      _208_v13 = _dafny.Map.Empty.slice().updateUnsafe(_194_v0,_194_v0);
      _208_v13 = (_208_v13).update(((_dafny.ZERO).minus(_194_v0)).multipliedBy(_194_v0), _194_v0);
      let _209_v14;
      _209_v14 = _dafny.Map.Empty.slice().updateUnsafe(_194_v0,_196_v2);
      _209_v14 = _209_v14;
      return;
    }
    m2(globalState) {
      let _this = this;
      let _210_v0;
      _210_v0 = false;
      let _211_v1;
      _211_v1 = _dafny.Seq.of(_210_v0);
      let _212_v2;
      _212_v2 = _dafny.Seq.of(_211_v1, _dafny.Seq.of(_210_v0), _211_v1);
      if (_dafny.Seq.IsProperPrefixOf(_212_v2, _dafny.Seq.Concat(_212_v2, _212_v2))) {
        let _213_v3;
        _213_v3 = new _dafny.CodePoint('a'.codePointAt(0));
        let _214_v4;
        _214_v4 = new BigNumber(-243);
        let _215_v5;
        _215_v5 = _dafny.Seq.of(_214_v4);
        let _216_v6;
        _216_v6 = _dafny.Map.Empty.slice().updateUnsafe(_213_v3,(_215_v5)[_module.__default.safeIndex(_214_v4, new BigNumber((_215_v5).length))]);
        _216_v6 = _module.__default.fm20(_214_v4, globalState);
        let _217_v7;
        let _init4 = ((_218_v4) => function (_219_i0) {
          return _module.__default.safeModuloInt(_219_i0, _218_v4);
        })(_214_v4);
        let _nw31 = Array((new BigNumber(29)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw31.length); _i0_4++) {
          _nw31[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _217_v7 = _nw31;
        let _index17 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_217_v7).length));
        (_217_v7)[_index17] = _214_v4;
        let _220_v8;
        _220_v8 = _dafny.Map.Empty.slice().updateUnsafe(_214_v4,_214_v4);
        let _221_v9;
        _221_v9 = _module.D7.create_DC14(_214_v4, _210_v0, _210_v0, _210_v0, _214_v4);
        let _index18 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_217_v7).length));
        let _rhs22 = ((((_220_v8).contains(_214_v4)) ? ((_220_v8).get(_214_v4)) : ((_221_v9).dtor_cf25))).multipliedBy(_214_v4);
        let _rhs23 = _214_v4;
        let _rhs24 = _217_v7;
        let _lhs14 = _217_v7;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_217_v7).length));
        let _lhs16 = globalState;
        _214_v4 = _rhs22;
        _lhs14[_lhs15] = _rhs23;
        _lhs16.f23 = _rhs24;
        let _222_v10;
        _222_v10 = _dafny.MultiSet.fromElements(_213_v3, _213_v3, _213_v3);
        (globalState).f17 = !(_222_v10).equals((_222_v10).Union(_dafny.MultiSet.fromElements(_213_v3)));
        let _223_v11;
        _223_v11 = _module.D4.create_DC8(_213_v3, _210_v0, _210_v0);
        let _224_v12;
        let _nw32 = Array((new BigNumber(2)).toNumber());
        _nw32[(_dafny.ZERO).toNumber()] = _210_v0;
        _nw32[(_dafny.ONE).toNumber()] = _210_v0;
        _224_v12 = _nw32;
        let _225_v13;
        _225_v13 = _224_v12;
        let _226_v14;
        _226_v14 = _dafny.Seq.of(_224_v12);
        let _227_v15;
        _227_v15 = _dafny.Map.Empty.slice().updateUnsafe(_210_v0,(_226_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_226_v14).length))]);
        let _228_v16;
        let _nw33 = Array((new BigNumber(20)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = _224_v12;
        _nw33[(_dafny.ONE).toNumber()] = _224_v12;
        _nw33[(new BigNumber(2)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(3)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(4)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(5)).toNumber()] = (_225_v13);
        _nw33[(new BigNumber(6)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(7)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(8)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(9)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(10)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(11)).toNumber()] = (((_227_v15).contains(_210_v0)) ? ((_227_v15).get(_210_v0)) : (_224_v12));
        _nw33[(new BigNumber(12)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(13)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(14)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(15)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(16)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(17)).toNumber()] = (_226_v14)[_module.__default.safeIndex(_214_v4, new BigNumber((_226_v14).length))];
        _nw33[(new BigNumber(18)).toNumber()] = _224_v12;
        _nw33[(new BigNumber(19)).toNumber()] = _224_v12;
        _228_v16 = _nw33;
        let _229_v17;
        _229_v17 = _228_v16;
        let _source6 = (((_223_v11).dtor_cf14) ? (_229_v17) : (_229_v17));
        let _230___mcc_h0 = _source6;
        let _231_cf16 = _230___mcc_h0;
        (globalState).f17 = _210_v0;
        let _index19 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_224_v12).length));
        (_224_v12)[_index19] = _210_v0;
        let _index20 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_224_v12).length));
        (_224_v12)[_index20] = _210_v0;
        let _232_v18;
        _232_v18 = _dafny.Seq.UnicodeFromString("rd");
        _232_v18 = _dafny.Seq.Concat(_232_v18, _232_v18);
        _211_v1 = _211_v1;
        let _233_v19;
        _233_v19 = _dafny.Map.Empty.slice().updateUnsafe(_210_v0,new BigNumber((_215_v5).length));
        let _234_v20;
        _234_v20 = _dafny.MultiSet.fromElements(_210_v0);
        let _235_v21;
        _235_v21 = _dafny.Map.Empty.slice().updateUnsafe(_210_v0,_234_v20);
        let _236_v22;
        _236_v22 = _module.D8.create_DC17(_213_v3, _213_v3, (_217_v7)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_217_v7).length))]);
        let _237_v23;
        _237_v23 = _dafny.Map.Empty.slice().updateUnsafe((_233_v19).Merge(_233_v19),((((_235_v21).contains(_210_v0)) ? ((_235_v21).get(_210_v0)) : (_234_v20))).IsDisjointFrom((_234_v20).update(_210_v0, _module.__default.abs(_module.__default.fm18(_dafny.MultiSet.FromArray(_211_v1), _236_v22, _210_v0, globalState)))));
        _237_v23 = (_237_v23).update(((!(_210_v0)) ? (_233_v19) : (_233_v19)), _210_v0);
      } else {
        let _238_v25;
        let _init5 = function (_239_i1) {
          return _module.__default.safeModuloInt(_239_i1, new BigNumber((function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(673)), function (_240_i2) {
              return new BigNumber((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(149)), function (_241_i3) {
                return new _dafny.CodePoint('i'.codePointAt(0));
              }), _dafny.Seq.UnicodeFromString("buof"))).length);
            })).Elements) {
              let _242_v24 = _compr_15;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(673)), function (_240_i2) {
                return new BigNumber((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(149)), function (_241_i3) {
                  return new _dafny.CodePoint('i'.codePointAt(0));
                }), _dafny.Seq.UnicodeFromString("buof"))).length);
              }), _242_v24)) {
                _coll15.push([(_242_v24).multipliedBy(new BigNumber(-364)),_dafny.Seq.UnicodeFromString("u")]);
              }
            }
            return _coll15;
          }()).length));
        };
        let _nw34 = Array((new BigNumber(13)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw34.length); _i0_5++) {
          _nw34[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _238_v25 = _nw34;
        let _243_v26;
        _243_v26 = new BigNumber(581);
        let _index21 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_238_v25).length));
        (_238_v25)[_index21] = _243_v26;
        let _index22 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_238_v25).length));
        (_238_v25)[_index22] = (_243_v26).plus((_dafny.ZERO).minus(_243_v26));
        (globalState).f19 = _210_v0;
        let _244_v27;
        _244_v27 = _dafny.Seq.UnicodeFromString("wxrqlhdfx");
        let _245_v28;
        _245_v28 = _dafny.Map.Empty.slice().updateUnsafe((_238_v25)[_module.__default.safeIndex(new BigNumber(503), new BigNumber((_238_v25).length))],_244_v27);
        _245_v28 = (_245_v28).update(_243_v26, _244_v27);
        let _246_v29;
        _246_v29 = _module.D3.create_DC4(_243_v26, _210_v0, _210_v0);
        let _247_v30;
        _247_v30 = _dafny.MultiSet.fromElements(_246_v29);
        (globalState).f17 = !(_module.__default.fm0(_211_v1, globalState)) || (!((_247_v30).IsSubsetOf(_247_v30)));
        (globalState).f19 = _210_v0;
      }
      let _248_v31;
      let _nw35 = new _module.C0();
      _nw35.__ctor(_210_v0);
      _248_v31 = _nw35;
      let _249_v32;
      _249_v32 = _dafny.MultiSet.fromElements(_248_v31, _248_v31);
      let _250_v33;
      let _nw36 = new _module.C0();
      _nw36.__ctor(!(_249_v32).contains(_248_v31));
      _250_v33 = _nw36;
      let _251_v34;
      _251_v34 = new BigNumber(169);
      let _252_v35;
      _252_v35 = _dafny.Set.fromElements(_251_v34);
      let _253_v36;
      _253_v36 = _dafny.Map.Empty.slice().updateUnsafe((_252_v35).Intersect(_252_v35),new BigNumber(-526));
      _253_v36 = ((_253_v36).Merge(function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_253_v36).Keys.Elements) {
          let _254_v37 = _compr_16;
          if ((_253_v36).contains(_254_v37)) {
            _coll16.push([_254_v37,(_dafny.ZERO).minus(_251_v34)]);
          }
        }
        return _coll16;
      }())).Merge(_253_v36);
      let _255_v38;
      _255_v38 = _dafny.Set.fromElements(!((_248_v31).f27));
      let _256_v39;
      _256_v39 = _dafny.Seq.UnicodeFromString("nfyl");
      let _257_v40;
      _257_v40 = _dafny.Seq.of(_256_v39);
      let _258_v41;
      _258_v41 = _dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC3(_257_v40));
      let _259_v42;
      _259_v42 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_251_v34),_210_v0);
      let _260_v44;
      _260_v44 = _dafny.MultiSet.fromElements(new BigNumber((_256_v39).length));
      let _261_v45;
      _261_v45 = _dafny.Map.Empty.slice().updateUnsafe(_251_v34,new BigNumber(812));
      let _262_v46;
      _262_v46 = _dafny.MultiSet.fromElements((_250_v33).f27);
      let _263_v47;
      _263_v47 = _dafny.Seq.of(new BigNumber((_262_v46).cardinality()));
      let _264_v48;
      let _nw37 = Array((new BigNumber(27)).toNumber());
      _nw37[(_dafny.ZERO).toNumber()] = false;
      _nw37[(_dafny.ONE).toNumber()] = !(_dafny.Set.fromElements((_250_v33).f27, (_250_v33).f27, (_250_v33).f27, (_250_v33).f27)).equals(_255_v38);
      _nw37[(new BigNumber(2)).toNumber()] = !(_210_v0);
      _nw37[(new BigNumber(3)).toNumber()] = (_248_v31).f27;
      _nw37[(new BigNumber(4)).toNumber()] = _module.__default.fm0(_211_v1, globalState);
      _nw37[(new BigNumber(5)).toNumber()] = (_250_v33).f27;
      _nw37[(new BigNumber(6)).toNumber()] = (new BigNumber((_258_v41).length)).isLessThan(_251_v34);
      _nw37[(new BigNumber(7)).toNumber()] = _210_v0;
      _nw37[(new BigNumber(8)).toNumber()] = (_259_v42).equals(_259_v42);
      _nw37[(new BigNumber(9)).toNumber()] = _210_v0;
      _nw37[(new BigNumber(10)).toNumber()] = _210_v0;
      _nw37[(new BigNumber(11)).toNumber()] = (function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_260_v44).Elements) {
          let _265_v43 = _compr_17;
          if ((_260_v44).contains(_265_v43)) {
            _coll17.push([_module.__default.safeModuloInt(_265_v43, _251_v34),_251_v34]);
          }
        }
        return _coll17;
      }()).equals(_261_v45);
      _nw37[(new BigNumber(12)).toNumber()] = (_211_v1)[_module.__default.safeIndex(new BigNumber((_263_v47).length), new BigNumber((_211_v1).length))];
      _nw37[(new BigNumber(13)).toNumber()] = !((_248_v31).f27);
      _nw37[(new BigNumber(14)).toNumber()] = (_248_v31).f27;
      _nw37[(new BigNumber(15)).toNumber()] = (_250_v33).f27;
      _nw37[(new BigNumber(16)).toNumber()] = (_250_v33).f27;
      _nw37[(new BigNumber(17)).toNumber()] = (_248_v31).f27;
      _nw37[(new BigNumber(18)).toNumber()] = (_250_v33).f27;
      _nw37[(new BigNumber(19)).toNumber()] = _210_v0;
      _nw37[(new BigNumber(20)).toNumber()] = (_250_v33).f27;
      _nw37[(new BigNumber(21)).toNumber()] = _210_v0;
      _nw37[(new BigNumber(22)).toNumber()] = (_248_v31).f27;
      _nw37[(new BigNumber(23)).toNumber()] = (_250_v33).f27;
      _nw37[(new BigNumber(24)).toNumber()] = (((_248_v31).f27) ? ((_248_v31).f27) : ((_248_v31).f27));
      _nw37[(new BigNumber(25)).toNumber()] = (_248_v31).f27;
      _nw37[(new BigNumber(26)).toNumber()] = _210_v0;
      _264_v48 = _nw37;
      let _index23 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_264_v48).length));
      (_264_v48)[_index23] = (_250_v33).f27;
      let _266_v49;
      _266_v49 = new _dafny.CodePoint('o'.codePointAt(0));
      let _267_v50;
      _267_v50 = _module.D8.create_DC17(_266_v49, _266_v49, new BigNumber((_256_v39).length));
      let _268_v51;
      _268_v51 = _dafny.Map.Empty.slice().updateUnsafe((_248_v31).f27,_module.__default.fm0(_dafny.Seq.of(!(false), _210_v0, (_250_v33).f27, (_250_v33).f27, (_248_v31).f27), globalState));
      let _269_v52;
      _269_v52 = _dafny.Map.Empty.slice().updateUnsafe((((_259_v42).contains((((_260_v44).contains(_251_v34)) ? ((_260_v44).get(_251_v34)) : (_module.__default.fm18(_262_v46, _267_v50, (((_268_v51).contains(_210_v0)) ? ((_268_v51).get(_210_v0)) : ((_248_v31).f27)), globalState))))) ? ((_259_v42).get((((_260_v44).contains(_251_v34)) ? ((_260_v44).get(_251_v34)) : (_module.__default.fm18(_262_v46, _267_v50, (((_268_v51).contains(_210_v0)) ? ((_268_v51).get(_210_v0)) : ((_248_v31).f27)), globalState))))) : ((_250_v33).f27)),_251_v34);
      let _270_v53;
      _270_v53 = _module.D7.create_DC13(_269_v52);
      let _pat_let_tv3 = _250_v33;
      let _pat_let_tv4 = _210_v0;
      let _pat_let_tv5 = _210_v0;
      let _index24 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_264_v48).length));
      (_264_v48)[_index24] = function (_source7) {
        if (_source7.is_DC14) {
          let _271___mcc_h1 = (_source7).cf21;
          let _272___mcc_h2 = (_source7).cf22;
          let _273___mcc_h3 = (_source7).cf23;
          let _274___mcc_h4 = (_source7).cf24;
          let _275___mcc_h5 = (_source7).cf25;
          let _276_cf25 = _275___mcc_h5;
          let _277_cf24 = _274___mcc_h4;
          let _278_cf23 = _273___mcc_h3;
          let _279_cf22 = _272___mcc_h2;
          let _280_cf21 = _271___mcc_h1;
          return (_pat_let_tv3).f27;
        } else if (_source7.is_DC15) {
          let _281___mcc_h6 = (_source7).cf26;
          let _282___mcc_h7 = (_source7).cf27;
          let _283___mcc_h8 = (_source7).cf28;
          let _284_cf28 = _283___mcc_h8;
          let _285_cf27 = _282___mcc_h7;
          let _286_cf26 = _281___mcc_h6;
          return _pat_let_tv4;
        } else {
          let _287___mcc_h9 = (_source7).cf20;
          let _288_cf20 = _287___mcc_h9;
          return _pat_let_tv5;
        }
      }(_270_v53);
      let _source8 = _267_v50;
      if (_source8.is_DC17) {
        let _289___mcc_h10 = (_source8).cf30;
        let _290___mcc_h11 = (_source8).cf31;
        let _291___mcc_h12 = (_source8).cf32;
        let _292_cf32 = _291___mcc_h12;
        let _293_cf31 = _290___mcc_h11;
        let _294_cf30 = _289___mcc_h10;
        let _295_v54;
        let _296_v55;
        let _297_v56;
        let _298_v57;
        let _out8;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector2 = _module.__default.m0(globalState);
        _out8 = _outcollector2[0];
        _out9 = _outcollector2[1];
        _out10 = _outcollector2[2];
        _out11 = _outcollector2[3];
        _295_v54 = _out8;
        _296_v55 = _out9;
        _297_v56 = _out10;
        _298_v57 = _out11;
        (globalState).f8 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_295_v54, _298_v57), _292_cf32);
        if ((_250_v33).f27) {
          _292_cf32 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_295_v54));
          let _299_v58;
          _299_v58 = _dafny.Set.fromElements(_module.__default.fm21(_module.__default.fm22((_250_v33).f27, globalState), _295_v54, globalState));
          let _300_v59;
          _300_v59 = _dafny.Seq.of(_299_v58);
          let _301_v60;
          let _nw38 = new _module.C0();
          _nw38.__ctor(!(((_300_v59)[_module.__default.safeIndex(_292_cf32, new BigNumber((_300_v59).length))]).IsProperSubsetOf(_299_v58)));
          _301_v60 = _nw38;
          let _302_v61;
          let _303_v62;
          let _304_v63;
          let _305_v64;
          let _out12;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = _module.__default.m0(globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _out15 = _outcollector3[3];
          _302_v61 = _out12;
          _303_v62 = _out13;
          _304_v63 = _out14;
          _305_v64 = _out15;
          let _306_v65;
          _306_v65 = _module.D7.create_DC15(_305_v64, _296_v55, _305_v64);
          let _307_v66;
          let _nw39 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _307_v66 = _nw39;
          let _308_v67;
          _308_v67 = _dafny.Seq.of(_307_v66);
          let _309_v68;
          _309_v68 = _dafny.Seq.of(_308_v67);
          let _310_v69;
          let _nw40 = Array((new BigNumber(22)).toNumber());
          _nw40[(_dafny.ZERO).toNumber()] = _305_v64;
          _nw40[(_dafny.ONE).toNumber()] = new BigNumber(((((_306_v65).dtor_cf27) ? (_259_v42) : (_259_v42))).length);
          _nw40[(new BigNumber(2)).toNumber()] = _302_v61;
          _nw40[(new BigNumber(3)).toNumber()] = _298_v57;
          _nw40[(new BigNumber(4)).toNumber()] = _251_v34;
          _nw40[(new BigNumber(5)).toNumber()] = _305_v64;
          _nw40[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(_305_v64, _295_v54);
          _nw40[(new BigNumber(7)).toNumber()] = _302_v61;
          _nw40[(new BigNumber(8)).toNumber()] = _302_v61;
          _nw40[(new BigNumber(9)).toNumber()] = _292_cf32;
          _nw40[(new BigNumber(10)).toNumber()] = (_295_v54).minus(_305_v64);
          _nw40[(new BigNumber(11)).toNumber()] = _305_v64;
          _nw40[(new BigNumber(12)).toNumber()] = new BigNumber((_261_v45).length);
          _nw40[(new BigNumber(13)).toNumber()] = _298_v57;
          _nw40[(new BigNumber(14)).toNumber()] = (((_250_v33).f27) ? (new BigNumber((_256_v39).length)) : (new BigNumber(909)));
          _nw40[(new BigNumber(15)).toNumber()] = _295_v54;
          _nw40[(new BigNumber(16)).toNumber()] = new BigNumber(((_309_v68)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_309_v68).length))]).length);
          _nw40[(new BigNumber(17)).toNumber()] = ((_210_v0) ? (_302_v61) : (_module.__default.fm18(_262_v46, _module.D8.create_DC17(_294_cf30, _266_v49, new BigNumber(435)), true, globalState)));
          _nw40[(new BigNumber(18)).toNumber()] = _292_cf32;
          _nw40[(new BigNumber(19)).toNumber()] = _302_v61;
          _nw40[(new BigNumber(20)).toNumber()] = new BigNumber(226);
          _nw40[(new BigNumber(21)).toNumber()] = _295_v54;
          _310_v69 = _nw40;
          let _index25 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_307_v66).length));
          (_307_v66)[_index25] = _302_v61;
          let _index26 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_307_v66).length));
          let _rhs25 = (_302_v61).isLessThanOrEqualTo(_305_v64);
          let _rhs26 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_268_v51).length), (((_261_v45).contains(_292_cf32)) ? ((_261_v45).get(_292_cf32)) : (_298_v57))), ((_dafny.ZERO).minus(_251_v34)).multipliedBy(_module.__default.fm18(_262_v46, _267_v50, (_248_v31).f27, globalState)));
          let _lhs17 = globalState;
          let _lhs18 = _307_v66;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_307_v66).length));
          _lhs17.f22 = _rhs25;
          _lhs18[_lhs19] = _rhs26;
          _302_v61 = (_dafny.ZERO).minus(_295_v54);
        } else {
          (globalState).f5 = !((_dafny.MultiSet.fromElements((_248_v31).f27, (_211_v1)[_module.__default.safeIndex(_298_v57, new BigNumber((_211_v1).length))], (_250_v33).f27, (_264_v48)[_module.__default.safeIndex(new BigNumber(380), new BigNumber((_264_v48).length))], (_248_v31).f27)).IsDisjointFrom((_262_v46).Intersect(_262_v46)));
          let _311_v70;
          _311_v70 = _dafny.MultiSet.fromElements(_269_v52);
          _251_v34 = (_dafny.ZERO).minus((((_248_v31).f27) ? (new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.of(_269_v52, _269_v52, _269_v52, _269_v52, _269_v52))).Union(_311_v70)).cardinality())) : (_292_cf32)));
          _251_v34 = _module.__default.safeModuloInt(new BigNumber(-796), (_251_v34).plus(_295_v54));
          _269_v52 = (_269_v52).update((_250_v33).f27, _module.__default.safeModuloInt(_295_v54, new BigNumber((_269_v52).length)));
          let _312_v71;
          let _nw41 = Array((new BigNumber(2)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _264_v48;
          _nw41[(_dafny.ONE).toNumber()] = _264_v48;
          _312_v71 = _nw41;
          let _index27 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_312_v71).length));
          (_312_v71)[_index27] = _264_v48;
          let _index28 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_264_v48).length));
          let _index29 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_312_v71).length));
          let _rhs27 = (_211_v1)[_module.__default.safeIndex(_251_v34, new BigNumber((_211_v1).length))];
          let _rhs28 = _264_v48;
          let _rhs29 = !((_248_v31).f27);
          let _rhs30 = _296_v55;
          let _lhs20 = _264_v48;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_264_v48).length));
          let _lhs22 = _312_v71;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_312_v71).length));
          let _lhs24 = globalState;
          _lhs20[_lhs21] = _rhs27;
          _lhs22[_lhs23] = _rhs28;
          _lhs24.f22 = _rhs29;
          _210_v0 = _rhs30;
        }
        let _313_v72;
        let _nw42 = new _module.C0();
        _nw42.__ctor((_248_v31).f27);
        _313_v72 = _nw42;
      } else if (_source8.is_DC16) {
        let _314___mcc_h13 = (_source8).cf29;
        let _315_cf29 = _314___mcc_h13;
        (globalState).f5 = _210_v0;
        if (!((((_260_v44).contains(new BigNumber((_211_v1).length))) ? ((_260_v44).get(new BigNumber((_211_v1).length))) : (new BigNumber((_module.__default.fm23((_248_v31).f27, _251_v34, _210_v0, globalState)).length)))).isEqualTo(_251_v34)) {
          (globalState).f8 = (((_261_v45).contains((_251_v34).minus(_251_v34))) ? ((_261_v45).get((_251_v34).minus(_251_v34))) : (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_250_v33).f27,new BigNumber((_dafny.Seq.UnicodeFromString("rclqacbxk")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!((_264_v48)[_module.__default.safeIndex(new BigNumber(380), new BigNumber((_264_v48).length))]),_251_v34))).length)));
          _256_v39 = _315_cf29;
          let _316_v73;
          _316_v73 = _dafny.MultiSet.fromElements(_264_v48);
          _316_v73 = _316_v73;
          let _317_v74;
          let _nw43 = Array((new BigNumber(19)).toNumber()).fill(_module.D8.Default());
          _317_v74 = _nw43;
          let _318_v75;
          _318_v75 = _dafny.Seq.of(_317_v74, _317_v74);
          let _319_v76;
          _319_v76 = _dafny.MultiSet.fromElements((_318_v75)[_module.__default.safeIndex(_251_v34, new BigNumber((_318_v75).length))]);
          let _320_v77;
          _320_v77 = _dafny.Map.Empty.slice().updateUnsafe((_319_v76).Union(_319_v76),(_dafny.ZERO).minus(_module.__default.safeModuloInt(_251_v34, _251_v34)));
          _320_v77 = (_320_v77).update((_319_v76).Intersect(_dafny.MultiSet.fromElements(_317_v74, _317_v74, _317_v74)), (((_262_v46).contains((_248_v31).f27)) ? ((_262_v46).get((_248_v31).f27)) : (new BigNumber(((_dafny.MultiSet.FromArray(_module.__default.fm24((_250_v33).f27, globalState))).update(function () {
            let _coll18 = new _dafny.Set();
            for (const _compr_18 of _dafny.IntegerRange(new BigNumber(905), new BigNumber(5))) {
              let _321_v78 = _compr_18;
              if (((new BigNumber(905)).isLessThanOrEqualTo(_321_v78)) && ((_321_v78).isLessThan(new BigNumber(5)))) {
                _coll18.add((_321_v78).plus(_251_v34));
              }
            }
            return _coll18;
          }(), _module.__default.abs(_251_v34))).cardinality()))));
          let _322_v79;
          let _nw44 = Array((new BigNumber(26)).toNumber());
          _322_v79 = _nw44;
          let _index30 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_322_v79).length));
          (_322_v79)[_index30] = _250_v33;
          let _323_v80;
          let _init6 = ((_324_v34) => function (_325_i4) {
            return _module.__default.safeModuloInt(_325_i4, _324_v34);
          })(_251_v34);
          let _nw45 = Array((new BigNumber(26)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw45.length); _i0_6++) {
            _nw45[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _323_v80 = _nw45;
          let _326_v81;
          _326_v81 = _dafny.Seq.of(_323_v80);
          let _index31 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_322_v79).length));
          let _nw46 = new _module.C0();
          _nw46.__ctor(_dafny.areEqual(_326_v81, _326_v81));
          (_322_v79)[_index31] = _nw46;
        } else {
          let _327_v82;
          _327_v82 = _module.D4.create_DC8(_266_v49, (_248_v31).f27, (_248_v31).f27);
          _266_v49 = (((((_250_v33).f27) ? ((_250_v33).f27) : (false))) ? ((_327_v82).dtor_cf13) : (_266_v49));
          (globalState).f5 = (_248_v31).f27;
          let _328_v83;
          _328_v83 = _dafny.Set.fromElements(_263_v47, _263_v47);
          _268_v51 = (_268_v51).update((_250_v33).f27, (_dafny.Set.fromElements(_263_v47)).IsProperSubsetOf(_328_v83));
          let _329_v84;
          _329_v84 = _dafny.Map.Empty.slice().updateUnsafe(_211_v1,(_250_v33).f27);
          _329_v84 = (_329_v84).update(_211_v1, !(!(_262_v46).contains((_248_v31).f27)));
          let _330_v85;
          _330_v85 = _module.D3.create_DC3(_dafny.Seq.of(_315_cf29));
          let _331_v86;
          _331_v86 = _module.D3.create_DC5(((true) ? (_module.__default.fm25(globalState)) : (_330_v85)));
          let _pat_let_tv6 = _330_v85;
          _331_v86 = function (_pat_let0_0) {
            return function (_332_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_333_dt__update_hcf7_h0) {
                  return _module.D3.create_DC5(_333_dt__update_hcf7_h0);
                }(_pat_let1_0);
              }(_pat_let_tv6);
            }(_pat_let0_0);
          }(_331_v86);
        }
        (globalState).f17 = (_211_v1)[_module.__default.safeIndex(_251_v34, new BigNumber((_211_v1).length))];
        let _rhs31 = _211_v1;
        let _rhs32 = (_dafny.ZERO).minus(new BigNumber((_268_v51).length));
        let _lhs25 = globalState;
        _211_v1 = _rhs31;
        _lhs25.f8 = _rhs32;
      } else {
        let _334___mcc_h14 = (_source8).cf33;
        let _335_cf33 = _334___mcc_h14;
        _264_v48 = _264_v48;
        let _336_v87;
        let _nw47 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _336_v87 = _nw47;
        let _337_v88;
        _337_v88 = _dafny.Map.Empty.slice().updateUnsafe(_248_v31,(_this).f29);
        let _338_v89;
        let _nw48 = Array((new BigNumber(15)).toNumber());
        _nw48[(_dafny.ZERO).toNumber()] = _251_v34;
        _nw48[(_dafny.ONE).toNumber()] = new BigNumber(64);
        _nw48[(new BigNumber(2)).toNumber()] = _251_v34;
        _nw48[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_251_v34);
        _nw48[(new BigNumber(4)).toNumber()] = _251_v34;
        _nw48[(new BigNumber(5)).toNumber()] = _251_v34;
        _nw48[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_211_v1, _211_v1)).length));
        _nw48[(new BigNumber(7)).toNumber()] = _251_v34;
        _nw48[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_251_v34, _251_v34));
        _nw48[(new BigNumber(9)).toNumber()] = (new BigNumber(-983)).multipliedBy(new BigNumber((_337_v88).length));
        _nw48[(new BigNumber(10)).toNumber()] = _251_v34;
        _nw48[(new BigNumber(11)).toNumber()] = _251_v34;
        _nw48[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(_251_v34, _251_v34);
        _nw48[(new BigNumber(13)).toNumber()] = _251_v34;
        _nw48[(new BigNumber(14)).toNumber()] = new BigNumber((_261_v45).length);
        _338_v89 = _nw48;
        let _rhs33 = _251_v34;
        let _rhs34 = _338_v89;
        let _rhs35 = _336_v87;
        let _lhs26 = globalState;
        let _lhs27 = globalState;
        _lhs26.f8 = _rhs33;
        _lhs27.f23 = _rhs34;
        _336_v87 = _rhs35;
        let _index32 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_338_v89).length));
        (_338_v89)[_index32] = (_251_v34).multipliedBy(_251_v34);
        let _index33 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_338_v89).length));
        (_338_v89)[_index33] = (_module.__default.fm18(_262_v46, _267_v50, _210_v0, globalState)).plus(new BigNumber(-785));
        _256_v39 = _256_v39;
      }
      let _index34 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_264_v48).length));
      (_264_v48)[_index34] = !(_module.__default.fm0(_211_v1, globalState));
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

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m1(globalState) {
      let _this = this;
      let _339_v0;
      _339_v0 = new BigNumber(420);
      let _340_v1;
      _340_v1 = _dafny.Seq.of(_339_v0);
      let _341_v2;
      _341_v2 = true;
      let _342_v3;
      _342_v3 = _module.D3.create_DC4(new BigNumber((_340_v1).length), _341_v2, _341_v2);
      if (((_dafny.ZERO).minus(_339_v0)).isLessThanOrEqualTo((_342_v3).dtor_cf4)) {
        let _343_v4;
        _343_v4 = _dafny.Seq.UnicodeFromString("t");
        let _344_v5;
        _344_v5 = _dafny.Seq.of(_343_v4);
        let _345_v6;
        _345_v6 = _module.D3.create_DC3(_344_v5);
        let _346_v7;
        _346_v7 = _module.D3.create_DC5(_345_v6);
        let _347_v8;
        _347_v8 = _module.D3.create_DC5(_345_v6);
        let _348_v9;
        _348_v9 = _dafny.Map.Empty.slice().updateUnsafe(_341_v2,_339_v0);
        let _349_v10;
        _349_v10 = _dafny.Map.Empty.slice().updateUnsafe(_347_v8,((_348_v9).update(_341_v2, _339_v0)).equals(_dafny.Map.Empty.slice().updateUnsafe(_341_v2,_339_v0)));
        let _350_v11;
        _350_v11 = new _dafny.CodePoint('k'.codePointAt(0));
        let _351_v12;
        _351_v12 = _dafny.Map.Empty.slice().updateUnsafe(_339_v0,_341_v2);
        _349_v10 = (_349_v10).update(_module.__default.fm3(_350_v11, _dafny.Seq.of(new BigNumber((_351_v12).length)), globalState), _341_v2);
        let _352_v13;
        _352_v13 = _dafny.MultiSet.fromElements(_341_v2, _341_v2, _341_v2);
        let _353_v14;
        _353_v14 = _dafny.Seq.of(_352_v13, _352_v13, (_352_v13).Union(_352_v13));
        _353_v14 = _353_v14;
        let _354_v15;
        _354_v15 = _dafny.Map.Empty.slice().updateUnsafe(_341_v2,_343_v4);
        let _355_v16;
        _355_v16 = _dafny.Map.Empty.slice().updateUnsafe(_341_v2,_352_v13);
        let _356_v17;
        let _nw49 = Array((new BigNumber(4)).toNumber());
        _nw49[(_dafny.ZERO).toNumber()] = new BigNumber(-502);
        _nw49[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_339_v0, new BigNumber(((((_354_v15).contains(true)) ? ((_354_v15).get(true)) : (_343_v4))).length))));
        _nw49[(new BigNumber(2)).toNumber()] = new BigNumber(124);
        _nw49[(new BigNumber(3)).toNumber()] = new BigNumber((((_341_v2) ? (_dafny.Map.Empty.slice().updateUnsafe(_341_v2,_dafny.MultiSet.fromElements(_341_v2))) : (_355_v16))).length);
        _356_v17 = _nw49;
        let _index35 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_356_v17).length));
        (_356_v17)[_index35] = _339_v0;
        let _index36 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_356_v17).length));
        (_356_v17)[_index36] = _module.__default.safeDivisionInt(_339_v0, _339_v0);
        let _357_v18;
        let _nw50 = new _module.C0();
        _nw50.__ctor(_341_v2);
        _357_v18 = _nw50;
        _357_v18 = _357_v18;
        _339_v0 = (_356_v17)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_356_v17).length))];
      } else {
        (globalState).f5 = ((_339_v0).minus(new BigNumber((_module.__default.fm4(_341_v2, globalState)).length))).isLessThanOrEqualTo(_339_v0);
        let _358_v19;
        _358_v19 = _module.D3.create_DC4(_339_v0, _341_v2, _341_v2);
        let _359_v20;
        _359_v20 = _module.D3.create_DC5(_358_v19);
        let _360_v21;
        _360_v21 = _dafny.Map.Empty.slice().updateUnsafe(_359_v20,_341_v2);
        _360_v21 = (_360_v21).update(_module.D3.create_DC5(_358_v19), _dafny.Seq.IsProperPrefixOf(_340_v1, _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(728)), ((_361_v0) => function (_362_i0) {
          return _361_v0;
        })(_339_v0))).length), _339_v0)));
        let _363_v22;
        _363_v22 = _dafny.Set.fromElements(_339_v0);
        let _364_v23;
        _364_v23 = _module.D4.create_DC6(_363_v22);
        let _365_v24;
        let _nw51 = new _module.C0();
        _nw51.__ctor(_341_v2);
        _365_v24 = _nw51;
        let _366_v25;
        _366_v25 = _dafny.Seq.UnicodeFromString("qky");
        let _367_v26;
        let _nw52 = Array((new BigNumber(5)).toNumber());
        _nw52[(_dafny.ZERO).toNumber()] = _339_v0;
        _nw52[(_dafny.ONE).toNumber()] = (_module.D4.create_DC7(_339_v0, _365_v24, _341_v2, _339_v0)).dtor_cf9;
        _nw52[(new BigNumber(2)).toNumber()] = new BigNumber(895);
        _nw52[(new BigNumber(3)).toNumber()] = new BigNumber((_366_v25).length);
        _nw52[(new BigNumber(4)).toNumber()] = _339_v0;
        _367_v26 = _nw52;
        let _368_v27;
        _368_v27 = _dafny.Map.Empty.slice().updateUnsafe(_367_v26,(_365_v24).f27);
        let _369_v29;
        _369_v29 = _dafny.Set.fromElements(true);
        let _370_v30;
        _370_v30 = _dafny.Map.Empty.slice().updateUnsafe((_365_v24).f27,(_365_v24).f27);
        let _371_v32;
        _371_v32 = _dafny.MultiSet.fromElements(!(_341_v2), (_365_v24).f27);
        let _372_v33;
        let _nw53 = Array((new BigNumber(25)).toNumber());
        _nw53[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_339_v0);
        _nw53[(_dafny.ONE).toNumber()] = _363_v22;
        _nw53[(new BigNumber(2)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(3)).toNumber()] = (_364_v23).dtor_cf8;
        _nw53[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_368_v27).length), _339_v0, _339_v0);
        _nw53[(new BigNumber(5)).toNumber()] = function () {
          let _coll19 = new _dafny.Set();
          for (const _compr_19 of _dafny.IntegerRange(new BigNumber(983), new BigNumber(683))) {
            let _373_v28 = _compr_19;
            if (((new BigNumber(983)).isLessThanOrEqualTo(_373_v28)) && ((_373_v28).isLessThan(new BigNumber(683)))) {
              _coll19.add((_373_v28).multipliedBy(_339_v0));
            }
          }
          return _coll19;
        }();
        _nw53[(new BigNumber(6)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), function (_374_i1) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length), _339_v0);
        _nw53[(new BigNumber(8)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(9)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(10)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(11)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(12)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(13)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(14)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(15)).toNumber()] = _dafny.Set.fromElements(new BigNumber(267), new BigNumber((_dafny.MultiSet.FromArray(_340_v1)).cardinality()), _module.__default.fm5(globalState), new BigNumber((_369_v29).length), _339_v0);
        _nw53[(new BigNumber(16)).toNumber()] = function () {
          let _coll20 = new _dafny.Set();
          for (const _compr_20 of (_module.__default.fm6(_340_v1, new BigNumber((_370_v30).length), new BigNumber(655), globalState)).Elements) {
            let _375_v31 = _compr_20;
            if ((_module.__default.fm6(_340_v1, new BigNumber((_370_v30).length), new BigNumber(655), globalState)).contains(_375_v31)) {
              _coll20.add(_module.__default.safeDivisionInt(_375_v31, new BigNumber((_dafny.Seq.of(true)).length)));
            }
          }
          return _coll20;
        }();
        _nw53[(new BigNumber(17)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(18)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(19)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(20)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_371_v32).cardinality()));
        _nw53[(new BigNumber(21)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(22)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(23)).toNumber()] = _363_v22;
        _nw53[(new BigNumber(24)).toNumber()] = _dafny.Set.fromElements(_339_v0);
        _372_v33 = _nw53;
        let _376_v34;
        let _nw54 = Array((new BigNumber(7)).toNumber());
        _376_v34 = _nw54;
        let _377_v35;
        _377_v35 = _dafny.Map.Empty.slice().updateUnsafe(_372_v33,_376_v34);
        let _378_v36;
        _378_v36 = _dafny.Seq.of(_376_v34, _376_v34, _376_v34);
        _377_v35 = (_377_v35).update(_372_v33, (_378_v36)[_module.__default.safeIndex(_339_v0, new BigNumber((_378_v36).length))]);
        if (!((_365_v24).f27) || (_341_v2)) {
          let _379_v37;
          _379_v37 = new _dafny.CodePoint('n'.codePointAt(0));
          _379_v37 = _379_v37;
          let _380_v38;
          _380_v38 = _dafny.Seq.of(_341_v2);
          let _381_v39;
          let _nw55 = Array((new BigNumber(13)).toNumber());
          _nw55[(_dafny.ZERO).toNumber()] = (_365_v24).f27;
          _nw55[(_dafny.ONE).toNumber()] = (_365_v24).f27;
          _nw55[(new BigNumber(2)).toNumber()] = _module.__default.fm0(_380_v38, globalState);
          _nw55[(new BigNumber(3)).toNumber()] = _module.__default.fm0(_380_v38, globalState);
          _nw55[(new BigNumber(4)).toNumber()] = _341_v2;
          _nw55[(new BigNumber(5)).toNumber()] = (_365_v24).f27;
          _nw55[(new BigNumber(6)).toNumber()] = _341_v2;
          _nw55[(new BigNumber(7)).toNumber()] = ((_365_v24).f27) || ((_365_v24).f27);
          _nw55[(new BigNumber(8)).toNumber()] = false;
          _nw55[(new BigNumber(9)).toNumber()] = ((_365_v24).f27);
          _nw55[(new BigNumber(10)).toNumber()] = (_365_v24).f27;
          _nw55[(new BigNumber(11)).toNumber()] = !_dafny.areEqual(_380_v38, _380_v38);
          _nw55[(new BigNumber(12)).toNumber()] = (_365_v24).f27;
          _381_v39 = _nw55;
          let _index37 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_381_v39).length));
          (_381_v39)[_index37] = false;
          let _382_v40;
          _382_v40 = _dafny.Map.Empty.slice().updateUnsafe(_339_v0,false);
          let _383_v41;
          _383_v41 = _dafny.Map.Empty.slice().updateUnsafe((((_382_v40).contains(_339_v0)) ? ((_382_v40).get(_339_v0)) : ((_365_v24).f27)),_366_v25);
          let _index38 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_381_v39).length));
          let _rhs36 = (_module.D4.create_DC8(new _dafny.CodePoint('m'.codePointAt(0)), (_365_v24).f27, _341_v2)).dtor_cf14;
          let _rhs37 = !(_dafny.Seq.IsProperPrefixOf(_366_v25, (((_383_v41).contains(!(_341_v2))) ? ((_383_v41).get(!(_341_v2))) : (_366_v25))));
          let _lhs28 = _381_v39;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_381_v39).length));
          let _lhs30 = globalState;
          _lhs28[_lhs29] = _rhs36;
          _lhs30.f19 = _rhs37;
          let _384_v42;
          let _385_v43;
          let _386_v44;
          let _387_v45;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector4 = _module.__default.m0(globalState);
          _out16 = _outcollector4[0];
          _out17 = _outcollector4[1];
          _out18 = _outcollector4[2];
          _out19 = _outcollector4[3];
          _384_v42 = _out16;
          _385_v43 = _out17;
          _386_v44 = _out18;
          _387_v45 = _out19;
          let _388_v46;
          _388_v46 = _dafny.Map.Empty.slice().updateUnsafe(_379_v37,_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("jpftdff"), _module.__default.fm7(globalState)));
          _388_v46 = (_388_v46).update(_module.__default.fm8(globalState), _341_v2);
          let _389_v47;
          let _nw56 = new _module.C0();
          _nw56.__ctor((_381_v39)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_381_v39).length))]);
          _389_v47 = _nw56;
        } else {
          let _390_v48;
          _390_v48 = new _dafny.CodePoint('o'.codePointAt(0));
          let _391_v49;
          _391_v49 = _dafny.MultiSet.fromElements(_390_v48, _390_v48, new _dafny.CodePoint('t'.codePointAt(0)));
          _391_v49 = _module.__default.fm9(!(_341_v2) || ((_365_v24).f27), _339_v0, _340_v1, (_module.__default.fm10(globalState)).update(_dafny.Seq.UnicodeFromString("cb"), (_365_v24).f27), globalState);
          _390_v48 = _390_v48;
          let _392_v50;
          let _nw57 = Array((new BigNumber(18)).toNumber());
          _nw57[(_dafny.ZERO).toNumber()] = (_365_v24).f27;
          _nw57[(_dafny.ONE).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(2)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(3)).toNumber()] = _341_v2;
          _nw57[(new BigNumber(4)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(5)).toNumber()] = _module.__default.fm0(_dafny.Seq.of((_365_v24).f27, !(_341_v2)), globalState);
          _nw57[(new BigNumber(6)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(7)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(8)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(9)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(10)).toNumber()] = (_342_v3).dtor_cf5;
          _nw57[(new BigNumber(11)).toNumber()] = _341_v2;
          _nw57[(new BigNumber(12)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(13)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(14)).toNumber()] = _341_v2;
          _nw57[(new BigNumber(15)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(16)).toNumber()] = (_365_v24).f27;
          _nw57[(new BigNumber(17)).toNumber()] = (_365_v24).f27;
          _392_v50 = _nw57;
          let _393_v51;
          _393_v51 = _392_v50;
          let _394_v52;
          let _nw58 = Array((new BigNumber(14)).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = _392_v50;
          _nw58[(_dafny.ONE).toNumber()] = _392_v50;
          _nw58[(new BigNumber(2)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(3)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(4)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(5)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(6)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(7)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(8)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(9)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(10)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(11)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(12)).toNumber()] = _392_v50;
          _nw58[(new BigNumber(13)).toNumber()] = _392_v50;
          _394_v52 = _nw58;
          let _395_v53;
          _395_v53 = _dafny.Map.Empty.slice().updateUnsafe((_394_v52),_339_v0);
          let _396_v54;
          _396_v54 = _dafny.Map.Empty.slice().updateUnsafe(_367_v26,(((_395_v53).contains(_394_v52)) ? ((_395_v53).get(_394_v52)) : (_339_v0)));
          let _397_v55;
          _397_v55 = _module.D4.create_DC7((((_396_v54).contains(_367_v26)) ? ((_396_v54).get(_367_v26)) : (new BigNumber(251))), _365_v24, (_module.__default.fm11(_339_v0, new BigNumber((_module.__default.fm12(_341_v2, new BigNumber((_340_v1).length), (_365_v24).f27, globalState)).length), (_365_v24).f27, _339_v0, globalState)), _339_v0);
          let _rhs38 = _393_v51;
          let _rhs39 = _397_v55;
          let _rhs40 = (((_365_v24).f27) ? (_342_v3) : (_342_v3));
          let _rhs41 = _392_v50;
          _393_v51 = _rhs38;
          _397_v55 = _rhs39;
          _342_v3 = _rhs40;
          _392_v50 = _rhs41;
          let _398_v56;
          _398_v56 = _module.D6.create_DC10(_367_v26);
          let _399_v57;
          let _nw59 = Array((new BigNumber(27)).toNumber());
          _nw59[(_dafny.ZERO).toNumber()] = _367_v26;
          _nw59[(_dafny.ONE).toNumber()] = (_398_v56).dtor_cf17;
          _nw59[(new BigNumber(2)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(3)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(4)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(5)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(6)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(7)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(8)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(9)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(10)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(11)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(12)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(13)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(14)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(15)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(16)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(17)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(18)).toNumber()] = (((_365_v24).f27) ? (_367_v26) : (_367_v26));
          _nw59[(new BigNumber(19)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(20)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(21)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(22)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(23)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(24)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(25)).toNumber()] = _367_v26;
          _nw59[(new BigNumber(26)).toNumber()] = _367_v26;
          _399_v57 = _nw59;
          let _index39 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_399_v57).length));
          (_399_v57)[_index39] = _367_v26;
          let _index40 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_399_v57).length));
          let _rhs42 = _390_v48;
          let _rhs43 = _367_v26;
          let _rhs44 = _dafny.Seq.update(_366_v25, _module.__default.safeIndex(_339_v0, new BigNumber((_366_v25).length)), _module.__default.fm8(globalState));
          let _rhs45 = _365_v24;
          let _lhs31 = _399_v57;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_399_v57).length));
          _390_v48 = _rhs42;
          _lhs31[_lhs32] = _rhs43;
          _366_v25 = _rhs44;
          _365_v24 = _rhs45;
          let _400_v58;
          _400_v58 = _dafny.Map.Empty.slice().updateUnsafe((_365_v24).f27,new BigNumber(-320));
          let _401_v59;
          _401_v59 = _dafny.Seq.of(_341_v2);
          let _402_v60;
          _402_v60 = _module.D7.create_DC13(_400_v58);
          _400_v58 = ((_400_v58).update(!(_341_v2), new BigNumber((_401_v59).length))).Merge((_402_v60).dtor_cf20);
        }
        let _403_v61;
        let _init7 = ((_404_v24, _405_v2) => function (_406_i2) {
          return !((_404_v24).f27) || (_405_v2);
        })(_365_v24, _341_v2);
        let _nw60 = Array((new BigNumber(23)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw60.length); _i0_7++) {
          _nw60[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _403_v61 = _nw60;
        let _index41 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_403_v61).length));
        (_403_v61)[_index41] = true;
        let _index42 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_403_v61).length));
        (_403_v61)[_index42] = _module.__default.fm0(_dafny.Seq.of(_341_v2), globalState);
      }
      let _407_v62;
      let _nw61 = Array((new BigNumber(14)).toNumber()).fill([]);
      _407_v62 = _nw61;
      let _408_v63;
      let _nw62 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _408_v63 = _nw62;
      let _409_v64;
      _409_v64 = _module.D6.create_DC10(_408_v63);
      let _index43 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length));
      (_407_v62)[_index43] = (_409_v64).dtor_cf17;
      let _index44 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length));
      (_407_v62)[_index44] = _408_v63;
      let _410_v65;
      let _init8 = ((_411_v2) => function (_412_i4) {
        return _411_v2;
      })(_341_v2);
      let _nw63 = Array((new BigNumber(28)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw63.length); _i0_8++) {
        _nw63[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _410_v65 = _nw63;
      let _413_v66;
      _413_v66 = _dafny.Map.Empty.slice().updateUnsafe(_339_v0,_410_v65);
      let _hi2 = (_339_v0).multipliedBy(_339_v0);
      for (let _414_i3 = ((_341_v2) ? (_339_v0) : ((_dafny.ZERO).minus(new BigNumber((_413_v66).length)))); _414_i3.isLessThan(_hi2); _414_i3 = _414_i3.plus(_dafny.ONE)) {
        (globalState).f5 = false;
        (globalState).f8 = _339_v0;
        let _415_v68;
        _415_v68 = _dafny.Set.fromElements(_414_i3);
        let _416_v69;
        let _nw64 = new _module.C0();
        _nw64.__ctor((_415_v68).IsProperSubsetOf(function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of _dafny.IntegerRange(new BigNumber(692), new BigNumber(714))) {
            let _417_v67 = _compr_21;
            if (((new BigNumber(692)).isLessThanOrEqualTo(_417_v67)) && ((_417_v67).isLessThan(new BigNumber(714)))) {
              _coll21.add((_417_v67).plus(_339_v0));
            }
          }
          return _coll21;
        }()));
        _416_v69 = _nw64;
        (globalState).f8 = _module.__default.fm5(globalState);
      }
      if (_341_v2) {
        let _418_v70;
        _418_v70 = new _dafny.CodePoint('p'.codePointAt(0));
        let _419_v71;
        _419_v71 = _dafny.Map.Empty.slice().updateUnsafe(_418_v70,_339_v0);
        _419_v71 = (_419_v71).update(_418_v70, _339_v0);
        let _arr4 = (_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))];
        let _index45 = _module.__default.safeIndex(new BigNumber(603), new BigNumber(((_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))]).length));
        _arr4[_index45] = _339_v0;
        let _arr5 = (_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))];
        let _index46 = _module.__default.safeIndex(new BigNumber(603), new BigNumber(((_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))]).length));
        _arr5[_index46] = _339_v0;
        (globalState).f5 = _341_v2;
        let _arr6 = (_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))];
        let _index47 = _module.__default.safeIndex(new BigNumber(603), new BigNumber(((_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))]).length));
        _arr6[_index47] = ((_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))])[_module.__default.safeIndex(new BigNumber(603), new BigNumber(((_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))]).length))];
        let _420_v72;
        _420_v72 = _dafny.Seq.UnicodeFromString("xm");
        let _421_v73;
        _421_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm13(globalState)).length),_dafny.Seq.Concat(_420_v72, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("nbgsu"), _module.__default.safeIndex(((_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))])[_module.__default.safeIndex(new BigNumber(603), new BigNumber(((_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))]).length))], new BigNumber((_dafny.Seq.UnicodeFromString("nbgsu")).length)), new _dafny.CodePoint('t'.codePointAt(0)))));
        _421_v73 = (_421_v73).update(((_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))])[_module.__default.safeIndex(new BigNumber(603), new BigNumber(((_407_v62)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_407_v62).length))]).length))], _dafny.Seq.Concat(_420_v72, _420_v72));
      } else {
        (globalState).f22 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("buaexsrxf"), _dafny.Seq.UnicodeFromString("pwblopusp"));
        let _422_v74;
        _422_v74 = _dafny.Map.Empty.slice().updateUnsafe(_339_v0,_341_v2);
        let _423_v75;
        _423_v75 = _dafny.Seq.UnicodeFromString("pprdiuoua");
        let _424_v76;
        _424_v76 = _dafny.Seq.of(_422_v74, _module.__default.fm14(_341_v2, _423_v75, globalState), _422_v74);
        let _425_v78;
        _425_v78 = _dafny.Set.fromElements((_dafny.ZERO).minus(_339_v0));
        let _426_v79;
        _426_v79 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of (_425_v78).Elements) {
            let _427_v77 = _compr_22;
            if ((_425_v78).contains(_427_v77)) {
              _coll22.push([_module.__default.safeDivisionInt(_427_v77, _339_v0),(_module.D4.create_DC8(new _dafny.CodePoint('x'.codePointAt(0)), _341_v2, _341_v2)).dtor_cf14]);
            }
          }
          return _coll22;
        }()));
        let _428_v80;
        _428_v80 = _dafny.MultiSet.fromElements(true, _341_v2);
        let _rhs46 = _dafny.Seq.Concat(_424_v76, _dafny.Seq.Concat((((_426_v79).contains((((_422_v74).contains(new BigNumber((_423_v75).length))) ? ((_422_v74).get(new BigNumber((_423_v75).length))) : (_341_v2)))) ? ((_426_v79).get((((_422_v74).contains(new BigNumber((_423_v75).length))) ? ((_422_v74).get(new BigNumber((_423_v75).length))) : (_341_v2)))) : (_dafny.Seq.update(_424_v76, _module.__default.safeIndex(new BigNumber((_428_v80).cardinality()), new BigNumber((_424_v76).length)), _422_v74))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_423_v75).length),_341_v2))));
        let _rhs47 = _341_v2;
        let _rhs48 = (new BigNumber((_423_v75).length)).isLessThan(_339_v0);
        let _rhs49 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), ((_429_v0) => function (_430_i5) {
          return _429_v0;
        })(_339_v0))).length);
        let _lhs33 = globalState;
        let _lhs34 = globalState;
        _424_v76 = _rhs46;
        _lhs33.f22 = _rhs47;
        _lhs34.f19 = _rhs48;
        _339_v0 = _rhs49;
        let _431_v81;
        _431_v81 = _dafny.Seq.of(_341_v2, _341_v2, _341_v2);
        let _432_v82;
        _432_v82 = _dafny.MultiSet.fromElements(_339_v0);
        let _433_v83;
        _433_v83 = _module.D8.create_DC16(_423_v75);
        let _434_v84;
        _434_v84 = _dafny.MultiSet.fromElements(new BigNumber(165), (_339_v0).multipliedBy(_339_v0), new BigNumber((_431_v81).length), (((_432_v82).contains(new BigNumber(375))) ? ((_432_v82).get(new BigNumber(375))) : (new BigNumber(((_433_v83).dtor_cf29).length))));
        _432_v82 = ((_434_v84).Union(_432_v82)).Union(_432_v82);
        (globalState).f17 = _341_v2;
        _422_v74 = _422_v74;
      }
      let _435_v86;
      _435_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_341_v2, true)).cardinality()),_module.__default.fm15(new BigNumber(-198), globalState));
      let _436_v87;
      _436_v87 = new _dafny.CodePoint('j'.codePointAt(0));
      let _437_v88;
      _437_v88 = _dafny.Map.Empty.slice().updateUnsafe(_436_v87,_339_v0);
      (globalState).f22 = !(_339_v0).isEqualTo((_dafny.ZERO).minus((_339_v0).plus(new BigNumber((function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of ((((_435_v86).contains(_339_v0)) ? ((_435_v86).get(_339_v0)) : (_437_v88))).Keys.Elements) {
          let _438_v85 = _compr_23;
          if (((((_435_v86).contains(_339_v0)) ? ((_435_v86).get(_339_v0)) : (_437_v88))).contains(_438_v85)) {
            _coll23.push([_438_v85,_339_v0]);
          }
        }
        return _coll23;
      }()).length))));
      (globalState).f5 = _341_v2;
      return;
    }
    m2(globalState) {
      let _this = this;
      let _439_v0;
      _439_v0 = _dafny.Seq.UnicodeFromString("nmnolf");
      let _440_v1;
      _440_v1 = _module.D8.create_DC16(_439_v0);
      let _441_v2;
      _441_v2 = new BigNumber(346);
      let _pat_let_tv7 = _439_v0;
      let _pat_let_tv8 = _441_v2;
      let _pat_let_tv9 = _439_v0;
      _440_v1 = function (_pat_let2_0) {
        return function (_442_dt__update__tmp_h0) {
          return function (_pat_let3_0) {
            return function (_443_dt__update_hcf29_h0) {
              return _module.D8.create_DC16(_443_dt__update_hcf29_h0);
            }(_pat_let3_0);
          }(_dafny.Seq.update(_pat_let_tv7, _module.__default.safeIndex((_dafny.ZERO).minus(_pat_let_tv8), new BigNumber((_pat_let_tv9).length)), new _dafny.CodePoint('k'.codePointAt(0))));
        }(_pat_let2_0);
      }(_440_v1);
      let _444_v3;
      _444_v3 = false;
      let _445_v4;
      _445_v4 = _dafny.MultiSet.fromElements(_444_v3);
      let _446_v5;
      let _nw65 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _446_v5 = _nw65;
      let _447_v6;
      _447_v6 = new _dafny.CodePoint('y'.codePointAt(0));
      let _448_v7;
      _448_v7 = _dafny.Seq.of(false);
      let _449_v8;
      _449_v8 = _dafny.MultiSet.fromElements(new BigNumber((_448_v7).length));
      let _450_v9;
      _450_v9 = _dafny.Seq.of(_441_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_441_v2,_444_v3)).length));
      let _451_v10;
      let _nw66 = Array((new BigNumber(29)).toNumber());
      _nw66[(_dafny.ZERO).toNumber()] = _441_v2;
      _nw66[(_dafny.ONE).toNumber()] = new BigNumber((_445_v4).cardinality());
      _nw66[(new BigNumber(2)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(3)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(4)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(5)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(6)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(7)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(8)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(9)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(10)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.of(_446_v5)).length);
      _nw66[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dauxjhvs")).length));
      _nw66[(new BigNumber(13)).toNumber()] = new BigNumber(103);
      _nw66[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_447_v6)).length);
      _nw66[(new BigNumber(15)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(16)).toNumber()] = new BigNumber(294);
      _nw66[(new BigNumber(17)).toNumber()] = new BigNumber(958);
      _nw66[(new BigNumber(18)).toNumber()] = new BigNumber(990);
      _nw66[(new BigNumber(19)).toNumber()] = new BigNumber((_449_v8).cardinality());
      _nw66[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(_441_v2);
      _nw66[(new BigNumber(21)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(22)).toNumber()] = _module.__default.fm5(globalState);
      _nw66[(new BigNumber(23)).toNumber()] = new BigNumber((_450_v9).length);
      _nw66[(new BigNumber(24)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(25)).toNumber()] = new BigNumber((_439_v0).length);
      _nw66[(new BigNumber(26)).toNumber()] = new BigNumber((_dafny.Seq.of(_441_v2, (_450_v9)[_module.__default.safeIndex(_441_v2, new BigNumber((_450_v9).length))], new BigNumber((_439_v0).length), _441_v2)).length);
      _nw66[(new BigNumber(27)).toNumber()] = _441_v2;
      _nw66[(new BigNumber(28)).toNumber()] = _441_v2;
      _451_v10 = _nw66;
      let _452_v11;
      _452_v11 = _dafny.Seq.of(_446_v5);
      let _453_v12;
      _453_v12 = _441_v2;
      let _454_v13;
      _454_v13 = _dafny.Map.Empty.slice().updateUnsafe(_453_v12,_447_v6);
      let _455_i0;
      _455_i0 = _dafny.ZERO;
      L1: {
        while (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(_451_v10, (_452_v11)[_module.__default.safeIndex(new BigNumber((_454_v13).length), new BigNumber((_452_v11).length))], _446_v5, _451_v10), _452_v11), _dafny.Seq.Concat(_452_v11, _452_v11))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_455_i0)) {
              break L1;
            }
            _455_i0 = (_455_i0).plus(_dafny.ONE);
            if (_module.__default.fm0(_448_v7, globalState)) {
              let _456_v14;
              let _nw67 = Array((new BigNumber(23)).toNumber()).fill(false);
              _456_v14 = _nw67;
              let _index48 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_456_v14).length));
              (_456_v14)[_index48] = _444_v3;
              let _index49 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_456_v14).length));
              (_456_v14)[_index49] = !(false);
              (globalState).f8 = _module.__default.fm5(globalState);
              let _index50 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_451_v10).length));
              (_451_v10)[_index50] = _module.__default.safeDivisionInt(_441_v2, _441_v2);
              let _index51 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_451_v10).length));
              let _rhs50 = _450_v9;
              let _rhs51 = _441_v2;
              let _lhs35 = _451_v10;
              let _lhs36 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_451_v10).length));
              _450_v9 = _rhs50;
              _lhs35[_lhs36] = _rhs51;
              _445_v4 = (_445_v4).Difference(_445_v4);
              let _index52 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_451_v10).length));
              (_451_v10)[_index52] = _module.__default.safeModuloInt((_451_v10)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_451_v10).length))], ((_451_v10)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_451_v10).length))]).plus(new BigNumber((_439_v0).length)));
            } else {
              let _457_v15;
              _457_v15 = _dafny.Map.Empty.slice().updateUnsafe(_444_v3,_444_v3);
              (globalState).f17 = ((_441_v2).minus((_dafny.ZERO).minus(new BigNumber((_457_v15).length)))).isEqualTo(_441_v2);
              let _458_v16;
              let _nw68 = new _module.C0();
              _nw68.__ctor(_444_v3);
              _458_v16 = _nw68;
              let _459_v17;
              let _460_v18;
              let _461_v19;
              let _462_v20;
              let _out20;
              let _out21;
              let _out22;
              let _out23;
              let _outcollector5 = _module.__default.m0(globalState);
              _out20 = _outcollector5[0];
              _out21 = _outcollector5[1];
              _out22 = _outcollector5[2];
              _out23 = _outcollector5[3];
              _459_v17 = _out20;
              _460_v18 = _out21;
              _461_v19 = _out22;
              _462_v20 = _out23;
              _447_v6 = _447_v6;
              let _463_v21;
              let _init9 = ((_464_v6) => function (_465_i1) {
                return ((_module.D9.create_DC19(_dafny.MultiSet.fromElements(_464_v6))).dtor_cf34).Union(_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), _464_v6));
              })(_447_v6);
              let _nw69 = Array((new BigNumber(27)).toNumber());
              for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw69.length); _i0_9++) {
                _nw69[_i0_9] = _init9(new BigNumber(_i0_9));
              }
              _463_v21 = _nw69;
              let _466_v22;
              _466_v22 = _dafny.MultiSet.fromElements(_447_v6);
              let _467_v23;
              _467_v23 = _dafny.Seq.of(_439_v0);
              let _index53 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_463_v21).length));
              (_463_v21)[_index53] = ((_466_v22).update(_447_v6, _module.__default.abs(new BigNumber((_467_v23).length)))).update(_447_v6, _module.__default.abs(_459_v17));
              let _index54 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_463_v21).length));
              (_463_v21)[_index54] = _466_v22;
            }
            let _468_v24;
            _468_v24 = _dafny.Set.fromElements(_441_v2, _441_v2, _441_v2);
            let _469_v26;
            _469_v26 = _dafny.MultiSet.fromElements(_468_v24, _468_v24, function () {
              let _coll24 = new _dafny.Set();
              for (const _compr_24 of _dafny.IntegerRange(new BigNumber(197), new BigNumber(452))) {
                let _470_v25 = _compr_24;
                if (((new BigNumber(197)).isLessThanOrEqualTo(_470_v25)) && ((_470_v25).isLessThan(new BigNumber(452)))) {
                  _coll24.add((_470_v25).multipliedBy(_441_v2));
                }
              }
              return _coll24;
            }(), (_468_v24).Difference(_468_v24));
            let _471_v27;
            let _nw70 = Array((new BigNumber(13)).toNumber()).fill(false);
            _471_v27 = _nw70;
            let _index55 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_471_v27).length));
            (_471_v27)[_index55] = !(_444_v3);
            let _472_v28;
            let _nw71 = new _module.C0();
            _nw71.__ctor(!(false));
            _472_v28 = _nw71;
            let _473_v29;
            _473_v29 = _module.D4.create_DC7(_441_v2, _472_v28, true, _441_v2);
            let _index56 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_471_v27).length));
            let _rhs52 = (_469_v26).Intersect(_469_v26);
            let _rhs53 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(744)), ((_474_v2) => function (_475_i2) {
              return (_474_v2).minus(_474_v2);
            })(_441_v2))).length);
            let _rhs54 = _447_v6;
            let _rhs55 = (_473_v29).dtor_cf11;
            let _lhs37 = _471_v27;
            let _lhs38 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_471_v27).length));
            _469_v26 = _rhs52;
            _441_v2 = _rhs53;
            _447_v6 = _rhs54;
            _lhs37[_lhs38] = _rhs55;
            _441_v2 = _441_v2;
            (globalState).f22 = !(_dafny.areEqual(_450_v9, _dafny.Seq.Concat(_450_v9, _450_v9)));
          }
        }
      }
      let _source9 = _440_v1;
      if (_source9.is_DC17) {
        let _476___mcc_h0 = (_source9).cf30;
        let _477___mcc_h1 = (_source9).cf31;
        let _478___mcc_h2 = (_source9).cf32;
        let _479_cf32 = _478___mcc_h2;
        let _480_cf31 = _477___mcc_h1;
        let _481_cf30 = _476___mcc_h0;
        let _482_v30;
        _482_v30 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("wmd"));
        let _483_v31;
        let _nw72 = Array((new BigNumber(16)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_482_v30, _482_v30);
        _nw72[(_dafny.ONE).toNumber()] = _482_v30;
        _nw72[(new BigNumber(2)).toNumber()] = _482_v30;
        _nw72[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_482_v30, _module.__default.fm16(_480_cf31, _441_v2, new BigNumber(526), _479_cf32, globalState));
        _nw72[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), ((_484_v0) => function (_485_i3) {
          return _484_v0;
        })(_439_v0));
        _nw72[(new BigNumber(5)).toNumber()] = _482_v30;
        _nw72[(new BigNumber(6)).toNumber()] = _482_v30;
        _nw72[(new BigNumber(7)).toNumber()] = _482_v30;
        _nw72[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("sf"), _439_v0), _482_v30);
        _nw72[(new BigNumber(9)).toNumber()] = _482_v30;
        _nw72[(new BigNumber(10)).toNumber()] = _482_v30;
        _nw72[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_482_v30, _482_v30);
        _nw72[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_439_v0), _482_v30);
        _nw72[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(102)), ((_486_v0) => function (_487_i4) {
          return _486_v0;
        })(_439_v0));
        _nw72[(new BigNumber(14)).toNumber()] = _482_v30;
        _nw72[(new BigNumber(15)).toNumber()] = _482_v30;
        _483_v31 = _nw72;
        let _index57 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_483_v31).length));
        (_483_v31)[_index57] = _482_v30;
        let _index58 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_483_v31).length));
        (_483_v31)[_index58] = _dafny.Seq.update(_dafny.Seq.Concat(((_444_v3) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(231)), ((_488_v0) => function (_489_i5) {
          return _488_v0;
        })(_439_v0))) : (_482_v30)), _dafny.Seq.Concat(_482_v30, _dafny.Seq.of(_439_v0, _439_v0, _439_v0, _439_v0))), _module.__default.safeIndex(_module.__default.safeModuloInt((_dafny.ZERO).minus(_479_cf32), _479_cf32), new BigNumber((_dafny.Seq.Concat(((_444_v3) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(231)), ((_490_v0) => function (_491_i5) {
          return _490_v0;
        })(_439_v0))) : (_482_v30)), _dafny.Seq.Concat(_482_v30, _dafny.Seq.of(_439_v0, _439_v0, _439_v0, _439_v0)))).length)), _dafny.Seq.Concat(_439_v0, _439_v0));
        let _492_v33;
        _492_v33 = _dafny.Set.fromElements(_479_cf32);
        let _493_v34;
        _493_v34 = _dafny.Set.fromElements(_441_v2, new BigNumber((_492_v33).length));
        let _494_v35;
        _494_v35 = _dafny.Map.Empty.slice().updateUnsafe(true,_479_cf32);
        let _495_v36;
        _495_v36 = _dafny.Map.Empty.slice().updateUnsafe(_494_v35,_444_v3);
        let _496_v37;
        _496_v37 = _dafny.Set.fromElements(new BigNumber((_492_v33).length), new BigNumber((_495_v36).length), _479_cf32);
        let _497_v38;
        _497_v38 = _dafny.Map.Empty.slice().updateUnsafe(_444_v3,_dafny.Seq.update(_dafny.Seq.of(_451_v10), _module.__default.safeIndex(_441_v2, new BigNumber((_dafny.Seq.of(_451_v10)).length)), _446_v5));
        let _498_v39;
        _498_v39 = _dafny.Map.Empty.slice().updateUnsafe(!(function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of _dafny.IntegerRange(new BigNumber(731), new BigNumber(-903))) {
            let _499_v32 = _compr_25;
            if (((new BigNumber(731)).isLessThanOrEqualTo(_499_v32)) && ((_499_v32).isLessThan(new BigNumber(-903)))) {
              _coll25.add((_499_v32).minus(new BigNumber((_dafny.Seq.UnicodeFromString("mxwfxlfth")).length)));
            }
          }
          return _coll25;
        }()).equals(_492_v33),new BigNumber(((((_497_v38).contains(_444_v3)) ? ((_497_v38).get(_444_v3)) : (_dafny.Seq.of(_446_v5, _446_v5)))).length));
        _494_v35 = (_494_v35).update(_444_v3, _479_cf32);
        let _500_v40;
        let _nw73 = Array((new BigNumber(4)).toNumber()).fill(false);
        _500_v40 = _nw73;
        let _501_v41;
        _501_v41 = _dafny.Seq.of(_500_v40, _500_v40);
        let _502_v42;
        _502_v42 = _dafny.Map.Empty.slice().updateUnsafe((_441_v2).plus(new BigNumber((_439_v0).length)),_501_v41);
        _502_v42 = (_502_v42).update(_441_v2, _dafny.Seq.Concat(_501_v41, _dafny.Seq.of(_500_v40, _500_v40, _500_v40, _500_v40)));
        _441_v2 = new BigNumber(289);
      } else if (_source9.is_DC16) {
        let _503___mcc_h3 = (_source9).cf29;
        let _504_cf29 = _503___mcc_h3;
        let _505_v43;
        let _nw74 = Array((new BigNumber(2)).toNumber());
        _nw74[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(428)), ((_506_v6) => function (_507_i6) {
          return _506_v6;
        })(_447_v6));
        _nw74[(_dafny.ONE).toNumber()] = ((_444_v3) ? (_504_cf29) : (_439_v0));
        _505_v43 = _nw74;
        let _index59 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_505_v43).length));
        (_505_v43)[_index59] = _439_v0;
        let _index60 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_505_v43).length));
        let _rhs56 = (_441_v2).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(936)), ((_508_cf29, _509_v2) => function (_510_i7) {
          return (_508_cf29)[_module.__default.safeIndex(_509_v2, new BigNumber((_508_cf29).length))];
        })(_504_cf29, _441_v2))).length));
        let _rhs57 = _module.__default.safeModuloInt(_441_v2, _module.__default.fm5(globalState));
        let _rhs58 = !(_444_v3);
        let _rhs59 = _dafny.Seq.Concat(_dafny.Seq.Concat(_504_cf29, _504_cf29), _439_v0);
        let _lhs39 = globalState;
        let _lhs40 = _505_v43;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_505_v43).length));
        _441_v2 = _rhs56;
        _441_v2 = _rhs57;
        _lhs39.f17 = _rhs58;
        _lhs40[_lhs41] = _rhs59;
        let _511_v44;
        _511_v44 = _dafny.Map.Empty.slice().updateUnsafe(_445_v4,_444_v3);
        _511_v44 = (_511_v44).Merge(_511_v44);
        _447_v6 = _447_v6;
        let _512_v45;
        _512_v45 = _dafny.Map.Empty.slice().updateUnsafe(_444_v3,new BigNumber(479));
        _512_v45 = ((_module.__default.fm17(_444_v3, _441_v2, _444_v3, _444_v3, globalState)).Merge(_512_v45)).Merge(_512_v45);
      } else {
        let _513___mcc_h4 = (_source9).cf33;
        let _514_cf33 = _513___mcc_h4;
        let _index61 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_446_v5).length));
        (_446_v5)[_index61] = (_dafny.ZERO).minus(new BigNumber((_448_v7).length));
        let _index62 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_446_v5).length));
        (_446_v5)[_index62] = _441_v2;
        let _515_v46;
        let _nw75 = Array((new BigNumber(26)).toNumber()).fill([]);
        _515_v46 = _nw75;
        _515_v46 = _515_v46;
        let _516_v47;
        _516_v47 = _dafny.Map.Empty.slice().updateUnsafe(_439_v0,_444_v3);
        _516_v47 = ((_516_v47).Merge(_516_v47)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_439_v0,_444_v3));
        (globalState).f22 = _module.__default.fm0(_dafny.Seq.Concat(_448_v7, _dafny.Seq.of(_444_v3, !(_444_v3), !(_444_v3))), globalState);
      }
      let _hi3 = new BigNumber((_448_v7).length);
      for (let _517_i8 = (_450_v9)[_module.__default.safeIndex(_441_v2, new BigNumber((_450_v9).length))]; _517_i8.isLessThan(_hi3); _517_i8 = _517_i8.plus(_dafny.ONE)) {
        let _518_v48;
        let _nw76 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
        _518_v48 = _nw76;
        let _index63 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_518_v48).length));
        (_518_v48)[_index63] = _448_v7;
        let _index64 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_518_v48).length));
        (_518_v48)[_index64] = _448_v7;
        let _519_v49;
        _519_v49 = _dafny.Set.fromElements(new BigNumber((_439_v0).length));
        let _520_v50;
        _520_v50 = _dafny.Map.Empty.slice().updateUnsafe(_519_v49,_439_v0);
        let _521_v51;
        let _nw77 = new _module.C0();
        _nw77.__ctor(_444_v3);
        _521_v51 = _nw77;
        let _522_v52;
        let _nw78 = new _module.C1();
        _nw78.__ctor(_520_v50, _521_v51);
        _522_v52 = _nw78;
        let _523_v53;
        _523_v53 = _dafny.Map.Empty.slice().updateUnsafe(_522_v52,_441_v2);
        let _rhs60 = _523_v53;
        let _rhs61 = (_module.__default.safeDivisionInt(new BigNumber(329), _441_v2)).plus(_517_i8);
        let _rhs62 = _517_i8;
        let _lhs42 = globalState;
        let _lhs43 = globalState;
        _523_v53 = _rhs60;
        _lhs42.f8 = _rhs61;
        _lhs43.f8 = _rhs62;
        let _524_v54;
        _524_v54 = _module.D7.create_DC15(new BigNumber((_445_v4).cardinality()), _444_v3, _441_v2);
        let _pat_let_tv10 = _439_v0;
        let _source10 = function (_pat_let4_0) {
          return function (_525_dt__update__tmp_h1) {
            return function (_pat_let5_0) {
              return function (_526_dt__update_hcf28_h0) {
                return _module.D7.create_DC15((_525_dt__update__tmp_h1).dtor_cf26, (_525_dt__update__tmp_h1).dtor_cf27, _526_dt__update_hcf28_h0);
              }(_pat_let5_0);
            }(_module.__default.safeModuloInt(new BigNumber((_pat_let_tv10).length), new BigNumber(422)));
          }(_pat_let4_0);
        }(_524_v54);
        if (_source10.is_DC14) {
          let _527___mcc_h5 = (_source10).cf21;
          let _528___mcc_h6 = (_source10).cf22;
          let _529___mcc_h7 = (_source10).cf23;
          let _530___mcc_h8 = (_source10).cf24;
          let _531___mcc_h9 = (_source10).cf25;
          let _532_cf25 = _531___mcc_h9;
          let _533_cf24 = _530___mcc_h8;
          let _534_cf23 = _529___mcc_h7;
          let _535_cf22 = _528___mcc_h6;
          let _536_cf21 = _527___mcc_h5;
          _519_v49 = _519_v49;
          let _537_v55;
          let _init10 = ((_538_v6, _539_cf22) => function (_540_i9) {
            return (_dafny.Map.Empty.slice().updateUnsafe(false,_538_v6)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_539_cf22,_538_v6));
          })(_447_v6, _535_cf22);
          let _nw79 = Array((new BigNumber(11)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw79.length); _i0_10++) {
            _nw79[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _537_v55 = _nw79;
          let _rhs63 = _537_v55;
          let _rhs64 = _446_v5;
          let _lhs44 = globalState;
          _537_v55 = _rhs63;
          _lhs44.f23 = _rhs64;
          (globalState).f5 = _534_cf23;
          let _541_v56;
          let _nw80 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _541_v56 = _nw80;
          let _index65 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_541_v56).length));
          (_541_v56)[_index65] = _447_v6;
          let _542_v57;
          _542_v57 = _dafny.Set.fromElements(_447_v6, _447_v6, new _dafny.CodePoint('l'.codePointAt(0)));
          let _index66 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_541_v56).length));
          let _rhs65 = _447_v6;
          let _rhs66 = _dafny.Seq.Concat(_439_v0, _dafny.Seq.UnicodeFromString("o"));
          let _rhs67 = (((_542_v57).IsProperSubsetOf(_542_v57)) ? ((_dafny.ZERO).minus(_517_i8)) : (_517_i8));
          let _lhs45 = _541_v56;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_541_v56).length));
          let _lhs47 = globalState;
          _lhs45[_lhs46] = _rhs65;
          _439_v0 = _rhs66;
          _lhs47.f8 = _rhs67;
        } else if (_source10.is_DC15) {
          let _543___mcc_h10 = (_source10).cf26;
          let _544___mcc_h11 = (_source10).cf27;
          let _545___mcc_h12 = (_source10).cf28;
          let _546_cf28 = _545___mcc_h12;
          let _547_cf27 = _544___mcc_h11;
          let _548_cf26 = _543___mcc_h10;
          let _index67 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_451_v10).length));
          (_451_v10)[_index67] = _517_i8;
          let _index68 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_451_v10).length));
          (_451_v10)[_index68] = _441_v2;
          let _549_v58;
          _549_v58 = _module.D6.create_DC11(_444_v3);
          let _550_v59;
          let _nw81 = Array((new BigNumber(20)).toNumber());
          _nw81[(_dafny.ZERO).toNumber()] = _444_v3;
          _nw81[(_dafny.ONE).toNumber()] = (new BigNumber(383)).isEqualTo((_dafny.ZERO).minus(_546_cf28));
          _nw81[(new BigNumber(2)).toNumber()] = (true) || (_444_v3);
          _nw81[(new BigNumber(3)).toNumber()] = _444_v3;
          _nw81[(new BigNumber(4)).toNumber()] = _444_v3;
          _nw81[(new BigNumber(5)).toNumber()] = false;
          _nw81[(new BigNumber(6)).toNumber()] = (_549_v58).dtor_cf18;
          _nw81[(new BigNumber(7)).toNumber()] = _444_v3;
          _nw81[(new BigNumber(8)).toNumber()] = !((false) && (_444_v3));
          _nw81[(new BigNumber(9)).toNumber()] = _547_cf27;
          _nw81[(new BigNumber(10)).toNumber()] = !(_547_cf27);
          _nw81[(new BigNumber(11)).toNumber()] = (_548_cf26).isLessThanOrEqualTo(_module.__default.fm5(globalState));
          _nw81[(new BigNumber(12)).toNumber()] = (_module.D6.create_DC11(_547_cf27)).dtor_cf18;
          _nw81[(new BigNumber(13)).toNumber()] = _444_v3;
          _nw81[(new BigNumber(14)).toNumber()] = _547_cf27;
          _nw81[(new BigNumber(15)).toNumber()] = (_449_v8).IsDisjointFrom((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_548_cf26), _546_cf28)).update((_451_v10)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_451_v10).length))], _module.__default.abs(_517_i8)));
          _nw81[(new BigNumber(16)).toNumber()] = _547_cf27;
          _nw81[(new BigNumber(17)).toNumber()] = _444_v3;
          _nw81[(new BigNumber(18)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber(168))).IsDisjointFrom(_449_v8);
          _nw81[(new BigNumber(19)).toNumber()] = _444_v3;
          _550_v59 = _nw81;
          let _index69 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_550_v59).length));
          (_550_v59)[_index69] = _444_v3;
          let _index70 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_550_v59).length));
          (_550_v59)[_index70] = _547_cf27;
          let _551_v60;
          _551_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_450_v9, (_450_v9)[_module.__default.safeIndex((_451_v10)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_451_v10).length))], new BigNumber((_450_v9).length))]),(_519_v49).Difference(_519_v49));
          let _552_v61;
          _552_v61 = _dafny.Seq.of(_519_v49);
          _551_v60 = (_551_v60).update(!_dafny.Seq.contains(_439_v0, _447_v6), ((_552_v61)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update((_518_v48)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_518_v48).length))], _module.__default.safeIndex(_517_i8, new BigNumber(((_518_v48)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_518_v48).length))]).length)), (_550_v59)[_module.__default.safeIndex(new BigNumber(610), new BigNumber((_550_v59).length))])).length), new BigNumber((_552_v61).length))]).Difference(function () {
            let _coll26 = new _dafny.Set();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(-289), new BigNumber(653))) {
              let _553_v62 = _compr_26;
              if (((new BigNumber(-289)).isLessThanOrEqualTo(_553_v62)) && ((_553_v62).isLessThan(new BigNumber(653)))) {
                _coll26.add(_module.__default.safeModuloInt(_553_v62, _441_v2));
              }
            }
            return _coll26;
          }()));
          let _index71 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_451_v10).length));
          (_451_v10)[_index71] = (_dafny.ZERO).minus(_517_i8);
        } else {
          let _554___mcc_h13 = (_source10).cf20;
          let _555_cf20 = _554___mcc_h13;
          (globalState).f17 = !((!(_444_v3)) && (_dafny.areEqual(_439_v0, _439_v0)));
          let _556_v63;
          let _nw82 = Array((new BigNumber(24)).toNumber()).fill(false);
          _556_v63 = _nw82;
          _556_v63 = _556_v63;
          _439_v0 = _dafny.Seq.Concat(_439_v0, _439_v0);
          let _557_v64;
          let _nw83 = new _module.C1();
          _nw83.__ctor((_520_v50).Merge(_520_v50), _521_v51);
          _557_v64 = _nw83;
        }
        let _558_v65;
        let _nw84 = new _module.C1();
        _nw84.__ctor(_module.__default.fm26(globalState), _521_v51);
        _558_v65 = _nw84;
        let _559_v66;
        let _nw85 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
        _559_v66 = _nw85;
        let _560_v67;
        _560_v67 = _dafny.Map.Empty.slice().updateUnsafe(_517_i8,new BigNumber(670));
        let _index72 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_559_v66).length));
        (_559_v66)[_index72] = _560_v67;
        let _index73 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_559_v66).length));
        let _rhs68 = _558_v65;
        let _rhs69 = (_560_v67).update(_module.__default.fm5(globalState), _517_i8);
        let _lhs48 = _559_v66;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_559_v66).length));
        _558_v65 = _rhs68;
        _lhs48[_lhs49] = _rhs69;
      }
      let _source11 = _module.D4.create_DC8(_447_v6, _444_v3, _444_v3);
      if (_source11.is_DC7) {
        let _561___mcc_h14 = (_source11).cf9;
        let _562___mcc_h15 = (_source11).cf10;
        let _563___mcc_h16 = (_source11).cf11;
        let _564___mcc_h17 = (_source11).cf12;
        let _565_cf12 = _564___mcc_h17;
        let _566_cf11 = _563___mcc_h16;
        let _567_cf10 = _562___mcc_h15;
        let _568_cf9 = _561___mcc_h14;
        let _569_v68;
        _569_v68 = _dafny.Map.Empty.slice().updateUnsafe(!(_444_v3),_568_cf9);
        let _570_v69;
        _570_v69 = _module.D7.create_DC13(_569_v68);
        let _571_v70;
        _571_v70 = _dafny.Seq.of(_570_v69, _570_v69, _module.D7.create_DC13(_569_v68));
        (globalState).f17 = !_dafny.areEqual(_571_v70, _dafny.Seq.update(_dafny.Seq.Concat(_571_v70, _571_v70), _module.__default.safeIndex(_565_cf12, new BigNumber((_dafny.Seq.Concat(_571_v70, _571_v70)).length)), _570_v69));
        let _572_v71;
        _572_v71 = (((_567_cf10).f27) ? (_444_v3) : (_566_cf11));
        let _573_v72;
        _573_v72 = _dafny.Map.Empty.slice().updateUnsafe(false,_566_cf11);
        let _574_v73;
        _574_v73 = _dafny.Set.fromElements(new BigNumber(-908), _565_cf12, new BigNumber((_573_v72).length), new BigNumber(315), _441_v2);
        _572_v71 = (_574_v73).IsProperSubsetOf(_dafny.Set.fromElements(_565_cf12));
        _439_v0 = _module.__default.fm7(globalState);
        let _575_v74;
        _575_v74 = _module.D9.create_DC20((_567_cf10).f27, false);
        let _index74 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_451_v10).length));
        (_451_v10)[_index74] = new BigNumber((_dafny.Seq.of((_575_v74).dtor_cf35, _444_v3)).length);
        let _576_v75;
        _576_v75 = _dafny.Map.Empty.slice().updateUnsafe(_451_v10,_574_v73);
        let _577_v76;
        _577_v76 = _dafny.Set.fromElements((_567_cf10).f27, false);
        let _578_v77;
        _578_v77 = _dafny.Map.Empty.slice().updateUnsafe(_576_v75,new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(_568_cf9), (_dafny.ZERO).minus(new BigNumber((_439_v0).length)), new BigNumber((_577_v76).length), _565_cf12), _dafny.Seq.Create(_module.__default.abs(new BigNumber(687)), ((_579_cf9) => function (_580_i10) {
          return (_dafny.ZERO).minus(_579_cf9);
        })(_568_cf9)))).length));
        let _index75 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_451_v10).length));
        (_451_v10)[_index75] = (((_578_v77).contains(_576_v75)) ? ((_578_v77).get(_576_v75)) : (_441_v2));
      } else if (_source11.is_DC8) {
        let _581___mcc_h18 = (_source11).cf13;
        let _582___mcc_h19 = (_source11).cf14;
        let _583___mcc_h20 = (_source11).cf15;
        let _584_cf15 = _583___mcc_h20;
        let _585_cf14 = _582___mcc_h19;
        let _586_cf13 = _581___mcc_h18;
        let _587_v78;
        let _nw86 = Array((new BigNumber(4)).toNumber()).fill(false);
        _587_v78 = _nw86;
        let _index76 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_587_v78).length));
        (_587_v78)[_index76] = _444_v3;
        let _index77 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_587_v78).length));
        (_587_v78)[_index77] = _444_v3;
        (globalState).f19 = (_587_v78)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_587_v78).length))];
        let _588_v79;
        _588_v79 = _module.D6.create_DC11(!_dafny.Seq.contains(_439_v0, _447_v6));
        let _source12 = _588_v79;
        if (_source12.is_DC11) {
          let _589___mcc_h22 = (_source12).cf18;
          let _590_cf18 = _589___mcc_h22;
          _587_v78 = _587_v78;
          let _591_v80;
          _591_v80 = _dafny.MultiSet.fromElements(_439_v0);
          let _rhs70 = (new BigNumber((_439_v0).length)).plus(_441_v2);
          let _rhs71 = new BigNumber((_591_v80).cardinality());
          let _lhs50 = globalState;
          let _lhs51 = globalState;
          _lhs50.f8 = _rhs70;
          _lhs51.f8 = _rhs71;
          (globalState).f8 = new BigNumber(-443);
          (globalState).f5 = _444_v3;
        } else if (_source12.is_DC10) {
          let _592___mcc_h23 = (_source12).cf17;
          let _593_cf17 = _592___mcc_h23;
          _441_v2 = _441_v2;
          let _rhs72 = (_587_v78)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_587_v78).length))];
          let _rhs73 = _441_v2;
          let _lhs52 = globalState;
          _585_cf14 = _rhs72;
          _lhs52.f8 = _rhs73;
          _450_v9 = _450_v9;
          let _594_v81;
          _594_v81 = _dafny.Map.Empty.slice().updateUnsafe(true,_584_cf15);
          let _595_v82;
          _595_v82 = _dafny.Seq.of(_594_v81, _dafny.Map.Empty.slice().updateUnsafe(_444_v3,_585_cf14));
          let _rhs74 = _439_v0;
          let _rhs75 = _595_v82;
          _439_v0 = _rhs74;
          _595_v82 = _rhs75;
        } else {
          let _596___mcc_h24 = (_source12).cf19;
          let _597_cf19 = _596___mcc_h24;
          (globalState).f8 = (new BigNumber(-467));
          _584_cf15 = (_585_cf14) === (!(!((_587_v78)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_587_v78).length))])));
          let _598_v83;
          let _nw87 = Array((new BigNumber(18)).toNumber()).fill([]);
          _598_v83 = _nw87;
          let _599_v84;
          _599_v84 = _dafny.Map.Empty.slice().updateUnsafe(_584_cf15,false);
          let _600_v85;
          _600_v85 = _dafny.Map.Empty.slice().updateUnsafe(_598_v83,new BigNumber((_599_v84).length));
          _600_v85 = (_600_v85).update(_598_v83, _441_v2);
          let _601_v86;
          let _602_v87;
          let _603_v88;
          let _out24;
          let _out25;
          let _out26;
          let _outcollector6 = (_this).m5(_444_v3, globalState);
          _out24 = _outcollector6[0];
          _out25 = _outcollector6[1];
          _out26 = _outcollector6[2];
          _601_v86 = _out24;
          _602_v87 = _out25;
          _603_v88 = _out26;
        }
        _585_cf14 = _444_v3;
      } else {
        let _604___mcc_h21 = (_source11).cf8;
        let _605_cf8 = _604___mcc_h21;
        let _606_v89;
        _606_v89 = _dafny.Set.fromElements(_444_v3, _444_v3, _444_v3, _444_v3);
        if ((_606_v89).IsSubsetOf(_606_v89)) {
          let _607_v90;
          let _init11 = ((_608_v3, _609_v2) => function (_610_i11) {
            return _module.D7.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_608_v3,_609_v2));
          })(_444_v3, _441_v2);
          let _nw88 = Array((new BigNumber(26)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw88.length); _i0_11++) {
            _nw88[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _607_v90 = _nw88;
          let _611_v91;
          _611_v91 = _dafny.Seq.of(_607_v90, _607_v90, _607_v90);
          let _612_v92;
          let _nw89 = Array((new BigNumber(29)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _607_v90;
          _nw89[(_dafny.ONE).toNumber()] = _607_v90;
          _nw89[(new BigNumber(2)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(3)).toNumber()] = (_611_v91)[_module.__default.safeIndex(_441_v2, new BigNumber((_611_v91).length))];
          _nw89[(new BigNumber(4)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(5)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(6)).toNumber()] = ((_444_v3) ? (_607_v90) : (_607_v90));
          _nw89[(new BigNumber(7)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(8)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(9)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(10)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(11)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(12)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(13)).toNumber()] = (_611_v91)[_module.__default.safeIndex(_441_v2, new BigNumber((_611_v91).length))];
          _nw89[(new BigNumber(14)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(15)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(16)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(17)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(18)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(19)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(20)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(21)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(22)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(23)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(24)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(25)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(26)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(27)).toNumber()] = _607_v90;
          _nw89[(new BigNumber(28)).toNumber()] = _607_v90;
          _612_v92 = _nw89;
          let _index78 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_612_v92).length));
          (_612_v92)[_index78] = _607_v90;
          let _index79 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_612_v92).length));
          (_612_v92)[_index79] = _607_v90;
          let _613_v93;
          _613_v93 = _dafny.Seq.of(_439_v0);
          let _614_v94;
          _614_v94 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_439_v0).length)),(_613_v93)[_module.__default.safeIndex(_441_v2, new BigNumber((_613_v93).length))]);
          let _615_v95;
          let _nw90 = new _module.C0();
          _nw90.__ctor(true);
          _615_v95 = _nw90;
          let _616_v96;
          let _nw91 = new _module.C1();
          _nw91.__ctor(_614_v94, _615_v95);
          _616_v96 = _nw91;
          _616_v96 = _616_v96;
          let _617_v97;
          _617_v97 = _dafny.Map.Empty.slice().updateUnsafe(!((_module.__default.fm0(_448_v7, globalState)) || (_444_v3)),(_450_v9)[_module.__default.safeIndex(_441_v2, new BigNumber((_450_v9).length))]);
          let _rhs76 = _444_v3;
          let _rhs77 = _617_v97;
          let _rhs78 = (_449_v8).IsSubsetOf((_449_v8).Intersect(_449_v8));
          let _lhs53 = globalState;
          let _lhs54 = globalState;
          _lhs53.f17 = _rhs76;
          _617_v97 = _rhs77;
          _lhs54.f5 = _rhs78;
          (globalState).f8 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_452_v11, _452_v11), ((_444_v3) ? (_dafny.Seq.update(_452_v11, _module.__default.safeIndex(new BigNumber((_445_v4).cardinality()), new BigNumber((_452_v11).length)), _451_v10)) : (_452_v11)))).length);
          _450_v9 = _450_v9;
        } else {
          let _618_v98;
          _618_v98 = _dafny.MultiSet.fromElements(_447_v6, _module.__default.fm8(globalState));
          _441_v2 = (((_618_v98).contains(_447_v6)) ? ((_618_v98).get(_447_v6)) : (new BigNumber((_606_v89).length)));
          (globalState).f5 = _module.__default.fm0(_448_v7, globalState);
          _439_v0 = _439_v0;
          let _619_v99;
          let _nw92 = new _module.C0();
          _nw92.__ctor(_444_v3);
          _619_v99 = _nw92;
          let _620_v100;
          _620_v100 = _dafny.Set.fromElements(_619_v99);
          (globalState).f5 = (_620_v100).IsProperSubsetOf(_620_v100);
          let _621_v101;
          _621_v101 = _dafny.Map.Empty.slice().updateUnsafe(_441_v2,_module.__default.fm0(_448_v7, globalState));
          let _622_v102;
          _622_v102 = _dafny.Set.fromElements(_449_v8, _dafny.MultiSet.fromElements(_module.__default.fm18(_dafny.MultiSet.FromArray(_448_v7), _module.__default.fm27(_441_v2, globalState), (_619_v99).f27, globalState), new BigNumber((_dafny.Seq.UnicodeFromString("lgd")).length)), _449_v8, _449_v8);
          (globalState).f22 = (((_621_v101).contains(new BigNumber((_622_v102).length))) ? ((_621_v101).get(new BigNumber((_622_v102).length))) : (!(_444_v3)));
        }
        let _623_v103;
        let _nw93 = Array((new BigNumber(21)).toNumber()).fill(false);
        _623_v103 = _nw93;
        let _index80 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_623_v103).length));
        (_623_v103)[_index80] = false;
        let _index81 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_623_v103).length));
        (_623_v103)[_index81] = _444_v3;
        let _624_v104;
        let _nw94 = Array((new BigNumber(22)).toNumber()).fill([]);
        _624_v104 = _nw94;
        let _625_v105;
        let _nw95 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _625_v105 = _nw95;
        let _index82 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_624_v104).length));
        (_624_v104)[_index82] = _625_v105;
        let _index83 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_624_v104).length));
        (_624_v104)[_index83] = _625_v105;
        let _626_v106;
        _626_v106 = _dafny.Map.Empty.slice().updateUnsafe(_605_cf8,_dafny.Seq.UnicodeFromString("krramqlq"));
        let _627_v107;
        let _nw96 = new _module.C0();
        _nw96.__ctor((_623_v103)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_623_v103).length))]);
        _627_v107 = _nw96;
        let _628_v108;
        let _nw97 = new _module.C1();
        _nw97.__ctor(_626_v106, _627_v107);
        _628_v108 = _nw97;
      }
      let _629_v109;
      _629_v109 = _dafny.Map.Empty.slice().updateUnsafe(_439_v0,_module.__default.fm5(globalState));
      let _rhs79 = _441_v2;
      let _rhs80 = (((_629_v109).contains(_dafny.Seq.Concat(_439_v0, _439_v0))) ? ((_629_v109).get(_dafny.Seq.Concat(_439_v0, _439_v0))) : (_441_v2));
      _441_v2 = _rhs79;
      _441_v2 = _rhs80;
      return;
    }
    m5(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = false;
      let _630_v0;
      _630_v0 = _dafny.Seq.of(p0, p0);
      (globalState).f5 = _module.__default.fm0(_630_v0, globalState);
      let _631_v1;
      let _init12 = ((_632_p0) => function (_633_i1) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rgdrpsk"),new BigNumber(107))).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(124)), function (_634_i2) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(941),_632_p0)).length))));
      })(p0);
      let _nw98 = Array((new BigNumber(12)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw98.length); _i0_12++) {
        _nw98[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _631_v1 = _nw98;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_631_v1).length))) {
        let _635_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_635_i0)) && ((_635_i0).isLessThan(new BigNumber((_631_v1).length))))) {
          (_631_v1)[(_635_i0)] = p0;
        }
      }
      let _636_v2;
      _636_v2 = new _dafny.CodePoint('g'.codePointAt(0));
      let _637_v3;
      let _nw99 = Array((new BigNumber(9)).toNumber());
      _nw99[(_dafny.ZERO).toNumber()] = _636_v2;
      _nw99[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
      _nw99[(new BigNumber(2)).toNumber()] = _636_v2;
      _nw99[(new BigNumber(3)).toNumber()] = _636_v2;
      _nw99[(new BigNumber(4)).toNumber()] = _636_v2;
      _nw99[(new BigNumber(5)).toNumber()] = _636_v2;
      _nw99[(new BigNumber(6)).toNumber()] = _636_v2;
      _nw99[(new BigNumber(7)).toNumber()] = _636_v2;
      _nw99[(new BigNumber(8)).toNumber()] = _636_v2;
      _637_v3 = _nw99;
      let _index84 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_637_v3).length));
      (_637_v3)[_index84] = _636_v2;
      let _index85 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_637_v3).length));
      (_637_v3)[_index85] = _636_v2;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_631_v1).length))) {
        let _638_i3 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_638_i3)) && ((_638_i3).isLessThan(new BigNumber((_631_v1).length))))) {
          (_631_v1)[(_638_i3)] = p0;
        }
      }
      let _639_v4;
      _639_v4 = new BigNumber(790);
      let _hi4 = (new BigNumber((_dafny.Seq.of((_637_v3)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_637_v3).length))])).length)).multipliedBy(_639_v4);
      for (let _640_i4 = _639_v4; _640_i4.isLessThan(_hi4); _640_i4 = _640_i4.plus(_dafny.ONE)) {
        let _641_v5;
        _641_v5 = _dafny.Set.fromElements(_636_v2);
        _641_v5 = _641_v5;
        let _642_v6;
        _642_v6 = _module.D7.create_DC15((_dafny.ZERO).minus(_module.__default.safeModuloInt(_639_v4, _639_v4)), p0, _640_i4);
        let _source13 = _642_v6;
        if (_source13.is_DC14) {
          let _643___mcc_h0 = (_source13).cf21;
          let _644___mcc_h1 = (_source13).cf22;
          let _645___mcc_h2 = (_source13).cf23;
          let _646___mcc_h3 = (_source13).cf24;
          let _647___mcc_h4 = (_source13).cf25;
          let _648_cf25 = _647___mcc_h4;
          let _649_cf24 = _646___mcc_h3;
          let _650_cf23 = _645___mcc_h2;
          let _651_cf22 = _644___mcc_h1;
          let _652_cf21 = _643___mcc_h0;
          let _653_v7;
          _653_v7 = _dafny.MultiSet.fromElements(_651_cf22);
          let _654_v8;
          let _nw100 = new _module.C0();
          _nw100.__ctor(_649_cf24);
          _654_v8 = _nw100;
          let _655_v9;
          _655_v9 = _module.D4.create_DC7(new BigNumber((_653_v7).cardinality()), _654_v8, _649_cf24, _640_i4);
          _655_v9 = _655_v9;
          (globalState).f19 = _650_cf23;
          let _656_v10;
          _656_v10 = _dafny.Map.Empty.slice().updateUnsafe(_651_cf22,!(_649_cf24));
          let _657_v11;
          _657_v11 = _dafny.Seq.of(_656_v10);
          let _658_v12;
          _658_v12 = _module.D10.create_DC21(_656_v10);
          let _659_v13;
          _659_v13 = _dafny.Map.Empty.slice().updateUnsafe((_637_v3)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_637_v3).length))],_650_cf23);
          let _660_v14;
          _660_v14 = _dafny.MultiSet.fromElements(new BigNumber(906));
          let _661_v15;
          _661_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_660_v14).cardinality()),_656_v10);
          let _pat_let_tv11 = p0;
          let _pat_let_tv12 = _659_v13;
          let _pat_let_tv13 = _637_v3;
          let _pat_let_tv14 = _637_v3;
          let _pat_let_tv15 = _654_v8;
          let _pat_let_tv16 = _659_v13;
          let _pat_let_tv17 = _637_v3;
          let _pat_let_tv18 = _637_v3;
          let _662_v16;
          let _nw101 = Array((new BigNumber(13)).toNumber());
          _nw101[(_dafny.ZERO).toNumber()] = (_657_v11)[_module.__default.safeIndex(_648_cf25, new BigNumber((_657_v11).length))];
          _nw101[(_dafny.ONE).toNumber()] = _656_v10;
          _nw101[(new BigNumber(2)).toNumber()] = _656_v10;
          _nw101[(new BigNumber(3)).toNumber()] = ((function (_pat_let6_0) {
            return function (_663_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_664_dt__update_hcf37_h0) {
                  return _module.D10.create_DC21(_664_dt__update_hcf37_h0);
                }(_pat_let7_0);
              }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv11,(((_pat_let_tv16).contains((_pat_let_tv18)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_pat_let_tv17).length))])) ? ((_pat_let_tv12).get((_pat_let_tv14)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_pat_let_tv13).length))])) : ((_pat_let_tv15).f27))));
            }(_pat_let6_0);
          }(_658_v12)).dtor_cf37).update(p0, _650_cf23);
          _nw101[(new BigNumber(4)).toNumber()] = _656_v10;
          _nw101[(new BigNumber(5)).toNumber()] = _656_v10;
          _nw101[(new BigNumber(6)).toNumber()] = _656_v10;
          _nw101[(new BigNumber(7)).toNumber()] = (((_661_v15).contains(_652_cf21)) ? ((_661_v15).get(_652_cf21)) : (_656_v10));
          _nw101[(new BigNumber(8)).toNumber()] = _656_v10;
          _nw101[(new BigNumber(9)).toNumber()] = _656_v10;
          _nw101[(new BigNumber(10)).toNumber()] = _656_v10;
          _nw101[(new BigNumber(11)).toNumber()] = _656_v10;
          _nw101[(new BigNumber(12)).toNumber()] = ((_module.__default.fm0(_630_v0, globalState)) ? ((_658_v12).dtor_cf37) : (_656_v10));
          _662_v16 = _nw101;
          let _index86 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_662_v16).length));
          (_662_v16)[_index86] = _dafny.Map.Empty.slice().updateUnsafe(true,_650_cf23);
          let _index87 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_662_v16).length));
          (_662_v16)[_index87] = (((_656_v10).update(_649_cf24, _651_cf22)).Merge(_656_v10)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_650_cf23,_649_cf24)).update(p0, _649_cf24));
          let _665_v17;
          _665_v17 = _dafny.Seq.of((_dafny.ZERO).minus(_652_cf21), _640_i4, _639_v4);
          let _666_v18;
          _666_v18 = _dafny.Seq.of(_665_v17, _665_v17);
          let _667_v19;
          let _nw102 = Array((new BigNumber(26)).toNumber());
          _nw102[(_dafny.ZERO).toNumber()] = _665_v17;
          _nw102[(_dafny.ONE).toNumber()] = _665_v17;
          _nw102[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm13(globalState), _dafny.Seq.of(_648_cf25));
          _nw102[(new BigNumber(3)).toNumber()] = (((_654_v8).f27) ? (_665_v17) : (_dafny.Seq.of(_652_cf21, _652_cf21)));
          _nw102[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_640_i4, _648_cf25);
          _nw102[(new BigNumber(5)).toNumber()] = (_666_v18)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(317)), ((_668_v2) => function (_669_i5) {
            return _668_v2;
          })(_636_v2))).length), new BigNumber((_666_v18).length))];
          _nw102[(new BigNumber(6)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_665_v17, _module.__default.safeIndex(_640_i4, new BigNumber((_665_v17).length)), new BigNumber(-505));
          _nw102[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(802)), ((_670_cf25) => function (_671_i6) {
            return _670_cf25;
          })(_648_cf25)), _665_v17);
          _nw102[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(127)), ((_672_cf24) => function (_673_i7) {
            return new BigNumber((_dafny.Set.fromElements(_672_cf24)).length);
          })(_649_cf24)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(457)), ((_674_cf21) => function (_675_i8) {
            return _674_cf21;
          })(_652_cf21)));
          _nw102[(new BigNumber(10)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_665_v17, _665_v17);
          _nw102[(new BigNumber(12)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_665_v17, _665_v17);
          _nw102[(new BigNumber(14)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(15)).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_665_v17).length)));
          _nw102[(new BigNumber(16)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(17)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(18)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(19)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_665_v17, _dafny.Seq.of(new BigNumber(604), _640_i4));
          _nw102[(new BigNumber(21)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_665_v17, _665_v17);
          _nw102[(new BigNumber(23)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(24)).toNumber()] = _665_v17;
          _nw102[(new BigNumber(25)).toNumber()] = _665_v17;
          _667_v19 = _nw102;
          let _676_v20;
          _676_v20 = _dafny.Map.Empty.slice().updateUnsafe((_654_v8).f27,_665_v17);
          let _nw103 = Array((new BigNumber(17)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = _665_v17;
          _nw103[(_dafny.ONE).toNumber()] = _665_v17;
          _nw103[(new BigNumber(2)).toNumber()] = (((_654_v8).f27) ? (_665_v17) : (_665_v17));
          _nw103[(new BigNumber(3)).toNumber()] = _665_v17;
          _nw103[(new BigNumber(4)).toNumber()] = _665_v17;
          _nw103[(new BigNumber(5)).toNumber()] = _665_v17;
          _nw103[(new BigNumber(6)).toNumber()] = _665_v17;
          _nw103[(new BigNumber(7)).toNumber()] = _665_v17;
          _nw103[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(184)), ((_677_cf25) => function (_678_i9) {
            return _677_cf25;
          })(_648_cf25));
          _nw103[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), function (_679_i10) {
            return new BigNumber(702);
          });
          _nw103[(new BigNumber(10)).toNumber()] = (((_676_v20).contains(true)) ? ((_676_v20).get(true)) : (_dafny.Seq.update(_665_v17, _module.__default.safeIndex(_648_cf25, new BigNumber((_665_v17).length)), _652_cf21)));
          _nw103[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), ((_680_i4) => function (_681_i11) {
            return _680_i4;
          })(_640_i4));
          _nw103[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_665_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(292)), ((_682_v8, _683_v7, _684_cf21) => function (_685_i12) {
            return (((_683_v7).contains((_682_v8).f27)) ? ((_683_v7).get((_682_v8).f27)) : (_684_cf21));
          })(_654_v8, _653_v7, _652_cf21)));
          _nw103[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_665_v17, _665_v17);
          _nw103[(new BigNumber(14)).toNumber()] = _665_v17;
          _nw103[(new BigNumber(15)).toNumber()] = _665_v17;
          _nw103[(new BigNumber(16)).toNumber()] = _665_v17;
          _667_v19 = _nw103;
        } else if (_source13.is_DC15) {
          let _686___mcc_h5 = (_source13).cf26;
          let _687___mcc_h6 = (_source13).cf27;
          let _688___mcc_h7 = (_source13).cf28;
          let _689_cf28 = _688___mcc_h7;
          let _690_cf27 = _687___mcc_h6;
          let _691_cf26 = _686___mcc_h5;
          let _692_v21;
          _692_v21 = _dafny.Map.Empty.slice().updateUnsafe(!(false),p0);
          let _693_v22;
          _693_v22 = _dafny.Seq.UnicodeFromString("y");
          let _694_v23;
          _694_v23 = _dafny.Map.Empty.slice().updateUnsafe(_640_i4,_636_v2);
          let _695_v24;
          _695_v24 = _dafny.Set.fromElements(_690_cf27, _690_cf27);
          let _696_v25;
          _696_v25 = _dafny.MultiSet.fromElements(p0, _690_cf27);
          let _697_v26;
          _697_v26 = _module.D11.create_DC24(_696_v25);
          let _698_v27;
          _698_v27 = _module.D8.create_DC17((_637_v3)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_637_v3).length))], _636_v2, new BigNumber((_693_v22).length));
          let _699_v28;
          _699_v28 = _dafny.MultiSet.fromElements((_637_v3)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_637_v3).length))], (_637_v3)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_637_v3).length))]);
          let _700_v29;
          _700_v29 = _dafny.Seq.of(_689_cf28, _module.__default.fm18(_dafny.MultiSet.fromElements(!(_690_cf27), !(_690_cf27)), _698_v27, !(p0), globalState), _691_cf26, _691_cf26);
          let _701_v30;
          let _nw104 = Array((new BigNumber(25)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_630_v0, globalState),!(p0))).Merge(_692_v21)).length);
          _nw104[(_dafny.ONE).toNumber()] = _689_cf28;
          _nw104[(new BigNumber(2)).toNumber()] = (new BigNumber((_693_v22).length)).multipliedBy(_640_i4);
          _nw104[(new BigNumber(3)).toNumber()] = (_module.__default.fm18(_dafny.MultiSet.fromElements(p0), _module.D8.create_DC17((_637_v3)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_637_v3).length))], (((_694_v23).contains(new BigNumber((_695_v24).length))) ? ((_694_v23).get(new BigNumber((_695_v24).length))) : (_636_v2)), _691_cf26), _690_cf27, globalState)).multipliedBy(_691_cf26);
          _nw104[(new BigNumber(4)).toNumber()] = _module.__default.fm18((_697_v26).dtor_cf41, _698_v27, false, globalState);
          _nw104[(new BigNumber(5)).toNumber()] = _module.__default.fm5(globalState);
          _nw104[(new BigNumber(6)).toNumber()] = new BigNumber((_699_v28).cardinality());
          _nw104[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_640_i4);
          _nw104[(new BigNumber(8)).toNumber()] = _689_cf28;
          _nw104[(new BigNumber(9)).toNumber()] = new BigNumber(517);
          _nw104[(new BigNumber(10)).toNumber()] = new BigNumber((_696_v25).cardinality());
          _nw104[(new BigNumber(11)).toNumber()] = (new BigNumber((_630_v0).length)).plus(_640_i4);
          _nw104[(new BigNumber(12)).toNumber()] = ((_690_cf27) ? (new BigNumber((_630_v0).length)) : (_691_cf26));
          _nw104[(new BigNumber(13)).toNumber()] = _689_cf28;
          _nw104[(new BigNumber(14)).toNumber()] = _module.__default.fm18(_696_v25, _698_v27, _690_cf27, globalState);
          _nw104[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_639_v4));
          _nw104[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(_639_v4, (_dafny.ZERO).minus(_639_v4));
          _nw104[(new BigNumber(17)).toNumber()] = (_640_i4).multipliedBy(new BigNumber(235));
          _nw104[(new BigNumber(18)).toNumber()] = (_640_i4).multipliedBy(_691_cf26);
          _nw104[(new BigNumber(19)).toNumber()] = (new BigNumber((_700_v29).length)).minus(_639_v4);
          _nw104[(new BigNumber(20)).toNumber()] = (new BigNumber((_693_v22).length)).minus(_691_cf26);
          _nw104[(new BigNumber(21)).toNumber()] = new BigNumber(690);
          _nw104[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(_691_cf26);
          _nw104[(new BigNumber(23)).toNumber()] = ((_690_cf27) ? ((_dafny.ZERO).minus(new BigNumber((_630_v0).length))) : (_640_i4));
          _nw104[(new BigNumber(24)).toNumber()] = _639_v4;
          _701_v30 = _nw104;
          let _index88 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_701_v30).length));
          (_701_v30)[_index88] = _691_cf26;
          let _index89 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_701_v30).length));
          (_701_v30)[_index89] = _module.__default.safeModuloInt(new BigNumber(((_696_v25).update(p0, _module.__default.abs((_700_v29)[_module.__default.safeIndex(_639_v4, new BigNumber((_700_v29).length))]))).cardinality()), new BigNumber(-466));
          let _702_v31;
          _702_v31 = _dafny.Map.Empty.slice().updateUnsafe(_636_v2,p0);
          let _703_v32;
          _703_v32 = _dafny.Map.Empty.slice().updateUnsafe(_702_v31,(_630_v0)[_module.__default.safeIndex(_689_cf28, new BigNumber((_630_v0).length))]);
          _703_v32 = (_703_v32).update(_702_v31, p0);
          _691_cf26 = _691_cf26;
          _689_cf28 = ((_700_v29)[_module.__default.safeIndex((_701_v30)[_module.__default.safeIndex(new BigNumber(714), new BigNumber((_701_v30).length))], new BigNumber((_700_v29).length))]).minus(_689_cf28);
        } else {
          let _704___mcc_h8 = (_source13).cf20;
          let _705_cf20 = _704___mcc_h8;
          _636_v2 = _636_v2;
          let _706_v33;
          _706_v33 = _dafny.Seq.UnicodeFromString("nc");
          let _707_v34;
          _707_v34 = _dafny.Map.Empty.slice().updateUnsafe(_706_v33,!(p0));
          let _708_v35;
          _708_v35 = _module.D3.create_DC4(new BigNumber((_707_v34).length), p0, p0);
          let _709_v36;
          _709_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-165),_637_v3);
          let _710_v37;
          _710_v37 = _dafny.Map.Empty.slice().updateUnsafe(_708_v35,new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_639_v4), new BigNumber((_709_v36).length))).cardinality()));
          let _index90 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_631_v1).length));
          (_631_v1)[_index90] = p0;
          let _711_v38;
          _711_v38 = _dafny.Seq.of(_640_i4, _640_i4);
          let _index91 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_631_v1).length));
          let _rhs81 = _710_v37;
          let _rhs82 = !((p0) || (p0));
          let _rhs83 = (_639_v4).multipliedBy(_640_i4);
          let _rhs84 = ((p0) ? (_630_v0) : (_630_v0));
          let _rhs85 = _module.__default.safeDivisionInt((_640_i4).multipliedBy(new BigNumber((_711_v38).length)), new BigNumber((_706_v33).length));
          let _lhs55 = _631_v1;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_631_v1).length));
          let _lhs57 = globalState;
          _710_v37 = _rhs81;
          _lhs55[_lhs56] = _rhs82;
          _639_v4 = _rhs83;
          _630_v0 = _rhs84;
          _lhs57.f8 = _rhs85;
          let _index92 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_631_v1).length));
          let _rhs86 = (_631_v1)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_631_v1).length))];
          let _rhs87 = _module.__default.safeDivisionInt(_640_i4, _640_i4);
          let _lhs58 = _631_v1;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_631_v1).length));
          let _lhs60 = globalState;
          _lhs58[_lhs59] = _rhs86;
          _lhs60.f8 = _rhs87;
          (globalState).f8 = (_dafny.ZERO).minus(_639_v4);
        }
        let _712_v39;
        _712_v39 = _dafny.Seq.of(_630_v0, _630_v0);
        let _713_v40;
        _713_v40 = _dafny.Map.Empty.slice().updateUnsafe(_640_i4,_639_v4);
        let _714_v41;
        let _nw105 = new _module.C0();
        _nw105.__ctor(p0);
        _714_v41 = _nw105;
        let _715_v42;
        _715_v42 = _module.D4.create_DC7(_639_v4, _714_v41, true, _639_v4);
        let _716_v43;
        _716_v43 = _dafny.MultiSet.fromElements(_640_i4, (_715_v42).dtor_cf12, new BigNumber(557));
        let _717_v44;
        let _nw106 = new _module.C0();
        _nw106.__ctor((_714_v41).f27);
        _717_v44 = _nw106;
        let _718_v45;
        _718_v45 = _dafny.Map.Empty.slice().updateUnsafe(_717_v44,new BigNumber(468));
        let _719_v46;
        _719_v46 = _dafny.Seq.of(_639_v4);
        let _720_v47;
        let _nw107 = Array((new BigNumber(23)).toNumber());
        _nw107[(_dafny.ZERO).toNumber()] = _639_v4;
        _nw107[(_dafny.ONE).toNumber()] = _640_i4;
        _nw107[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_640_i4);
        _nw107[(new BigNumber(3)).toNumber()] = _640_i4;
        _nw107[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_640_i4);
        _nw107[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_640_i4);
        _nw107[(new BigNumber(6)).toNumber()] = _640_i4;
        _nw107[(new BigNumber(7)).toNumber()] = _640_i4;
        _nw107[(new BigNumber(8)).toNumber()] = _639_v4;
        _nw107[(new BigNumber(9)).toNumber()] = _639_v4;
        _nw107[(new BigNumber(10)).toNumber()] = _module.__default.fm5(globalState);
        _nw107[(new BigNumber(11)).toNumber()] = new BigNumber((_712_v39).length);
        _nw107[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(p0)).length),!(p0))).length);
        _nw107[(new BigNumber(13)).toNumber()] = _639_v4;
        _nw107[(new BigNumber(14)).toNumber()] = _639_v4;
        _nw107[(new BigNumber(15)).toNumber()] = new BigNumber((_713_v40).length);
        _nw107[(new BigNumber(16)).toNumber()] = new BigNumber(((_716_v43).update(new BigNumber((_dafny.Seq.of((_630_v0)[_module.__default.safeIndex(_639_v4, new BigNumber((_630_v0).length))])).length), _module.__default.abs(_639_v4))).cardinality());
        _nw107[(new BigNumber(17)).toNumber()] = new BigNumber((_718_v45).length);
        _nw107[(new BigNumber(18)).toNumber()] = new BigNumber((_719_v46).length);
        _nw107[(new BigNumber(19)).toNumber()] = _640_i4;
        _nw107[(new BigNumber(20)).toNumber()] = new BigNumber((_716_v43).cardinality());
        _nw107[(new BigNumber(21)).toNumber()] = _640_i4;
        _nw107[(new BigNumber(22)).toNumber()] = new BigNumber((_630_v0).length);
        _720_v47 = _nw107;
        let _721_v48;
        _721_v48 = _dafny.Set.fromElements(_639_v4);
        let _722_v49;
        _722_v49 = _module.D3.create_DC4(_639_v4, (_714_v41).f27, p0);
        let _723_v50;
        _723_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_721_v48).length),_dafny.Seq.of((_722_v49).dtor_cf4, _640_i4));
        let _724_v51;
        _724_v51 = _dafny.Map.Empty.slice().updateUnsafe(_720_v47,(_module.D12.create_DC26(_723_v50)).dtor_cf44);
        _724_v51 = (_724_v51).update(_720_v47, function () {
          let _coll27 = new _dafny.Map();
          for (const _compr_27 of _dafny.IntegerRange(new BigNumber(-908), new BigNumber(148))) {
            let _725_v52 = _compr_27;
            if (((new BigNumber(-908)).isLessThanOrEqualTo(_725_v52)) && ((_725_v52).isLessThan(new BigNumber(148)))) {
              _coll27.push([_module.__default.safeModuloInt(_725_v52, new BigNumber(-723)),_719_v46]);
            }
          }
          return _coll27;
        }());
        let _726_v53;
        _726_v53 = _dafny.Seq.UnicodeFromString("yq");
        _726_v53 = _726_v53;
      }
      (globalState).f8 = (_dafny.ZERO).minus((_639_v4).multipliedBy(_639_v4));
      let _727_v54;
      _727_v54 = _dafny.Seq.UnicodeFromString("uigynbjvj");
      r0 = (new BigNumber(252)).plus(new BigNumber((_727_v54).length));
      let _728_v55;
      let _nw108 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
      _728_v55 = _nw108;
      r1 = _728_v55;
      let _729_v56;
      _729_v56 = _module.D7.create_DC14(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(962)), ((_730_v3) => function (_731_i13) {
  return (_730_v3)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_730_v3).length))];
})(_637_v3))).length), false, true, !(p0), _639_v4);
      let _pat_let_tv19 = _639_v4;
      let _pat_let_tv20 = p0;
      let _pat_let_tv21 = _637_v3;
      let _pat_let_tv22 = _637_v3;
      let _pat_let_tv23 = _639_v4;
      r2 = function (_source14) {
        if (_source14.is_DC14) {
          let _732___mcc_h9 = (_source14).cf21;
          let _733___mcc_h10 = (_source14).cf22;
          let _734___mcc_h11 = (_source14).cf23;
          let _735___mcc_h12 = (_source14).cf24;
          let _736___mcc_h13 = (_source14).cf25;
          let _737_cf25 = _736___mcc_h13;
          let _738_cf24 = _735___mcc_h12;
          let _739_cf23 = _734___mcc_h11;
          let _740_cf22 = _733___mcc_h10;
          let _741_cf21 = _732___mcc_h9;
          return _740_cf22;
        } else if (_source14.is_DC15) {
          let _742___mcc_h14 = (_source14).cf26;
          let _743___mcc_h15 = (_source14).cf27;
          let _744___mcc_h16 = (_source14).cf28;
          let _745_cf28 = _744___mcc_h16;
          let _746_cf27 = _743___mcc_h15;
          let _747_cf26 = _742___mcc_h14;
          return !(false);
        } else {
          let _748___mcc_h17 = (_source14).cf20;
          let _749_cf20 = _748___mcc_h17;
          return (_module.D7.create_DC15(_pat_let_tv19, _pat_let_tv20, new BigNumber((function () {
  let _coll28 = new _dafny.Map();
  for (const _compr_28 of _dafny.IntegerRange(new BigNumber(688), new BigNumber(852))) {
    let _750_v57 = _compr_28;
    if (((new BigNumber(688)).isLessThanOrEqualTo(_750_v57)) && ((_750_v57).isLessThan(new BigNumber(852)))) {
      _coll28.push([_module.__default.safeModuloInt(_750_v57, _pat_let_tv23),new BigNumber((_dafny.Set.fromElements((_pat_let_tv22)[_module.__default.safeIndex(new BigNumber(58), new BigNumber((_pat_let_tv21).length))])).length)]);
    }
  }
  return _coll28;
}()).length))).dtor_cf27;
        }
      }(_729_v56);
      return [r0, r1, r2];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this.f26 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f26) {
      let _this = this;
      (_this).f26 = f26;
      return;
    }
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(613)), function (_751_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("glht"), _dafny.Seq.UnicodeFromString("wx"), _dafny.Seq.UnicodeFromString("rlclyu"), _dafny.Seq.UnicodeFromString("llfeqpt"), _dafny.Seq.UnicodeFromString("ao"))), _dafny.Seq.Concat((_module.D3.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), function (_752_i1) {
  return _dafny.Seq.Create(_module.__default.abs(new BigNumber(95)), function (_753_i2) {
    return new _dafny.CodePoint('y'.codePointAt(0));
  });
}))).dtor_cf3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(613)), function (_754_i3) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(291)), function (_755_i4) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        });
      })));
    };
    m1(globalState) {
      let _this = this;
      let _756_v0;
      _756_v0 = new BigNumber(-373);
      if (!(_756_v0).isEqualTo(_756_v0)) {
        let _757_v1;
        _757_v1 = true;
        (globalState).f17 = _757_v1;
        (globalState).f5 = _757_v1;
        let _758_v2;
        let _nw109 = new _module.C0();
        _nw109.__ctor(_757_v1);
        _758_v2 = _nw109;
        let _759_v3;
        let _nw110 = Array((new BigNumber(16)).toNumber());
        _nw110[(_dafny.ZERO).toNumber()] = false;
        _nw110[(_dafny.ONE).toNumber()] = _757_v1;
        _nw110[(new BigNumber(2)).toNumber()] = (_758_v2).f27;
        _nw110[(new BigNumber(3)).toNumber()] = (_758_v2).f27;
        _nw110[(new BigNumber(4)).toNumber()] = _757_v1;
        _nw110[(new BigNumber(5)).toNumber()] = (_758_v2).f27;
        _nw110[(new BigNumber(6)).toNumber()] = _757_v1;
        _nw110[(new BigNumber(7)).toNumber()] = false;
        _nw110[(new BigNumber(8)).toNumber()] = (_758_v2).f27;
        _nw110[(new BigNumber(9)).toNumber()] = _757_v1;
        _nw110[(new BigNumber(10)).toNumber()] = _757_v1;
        _nw110[(new BigNumber(11)).toNumber()] = _757_v1;
        _nw110[(new BigNumber(12)).toNumber()] = _757_v1;
        _nw110[(new BigNumber(13)).toNumber()] = !((_758_v2).f27);
        _nw110[(new BigNumber(14)).toNumber()] = _757_v1;
        _nw110[(new BigNumber(15)).toNumber()] = !((_758_v2).f27);
        _759_v3 = _nw110;
        let _760_v4;
        _760_v4 = _dafny.Set.fromElements(_759_v3);
        let _761_v5;
        _761_v5 = _dafny.Seq.of(_760_v4);
        let _762_v6;
        _762_v6 = _dafny.Seq.of(_756_v0, _756_v0);
        let _arr7 = _this.f26;
        let _index93 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_this.f26).length));
        _arr7[_index93] = new BigNumber(((_761_v5)[_module.__default.safeIndex((_762_v6)[_module.__default.safeIndex(_756_v0, new BigNumber((_762_v6).length))], new BigNumber((_761_v5).length))]).length);
        let _763_v7;
        _763_v7 = _dafny.Map.Empty.slice().updateUnsafe(_756_v0,(_module.D4.create_DC7(new BigNumber((_dafny.Seq.of(_this.f26, _this.f26)).length), _758_v2, !(false), _756_v0)).dtor_cf11);
        let _764_v8;
        _764_v8 = _dafny.Seq.of(_757_v1, _757_v1, (_758_v2).f27);
        let _index94 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_759_v3).length));
        (_759_v3)[_index94] = (((_763_v7).contains(new BigNumber((_764_v8).length))) ? ((_763_v7).get(new BigNumber((_764_v8).length))) : ((_758_v2).f27));
        let _765_v9;
        _765_v9 = _dafny.Seq.UnicodeFromString("bycndgxtw");
        let _766_v10;
        _766_v10 = _dafny.MultiSet.fromElements(new BigNumber(348), new BigNumber((_765_v9).length));
        let _767_v11;
        _767_v11 = _dafny.Set.fromElements(_757_v1, (_764_v8)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_764_v8).length))]);
        let _arr8 = _this.f26;
        let _index95 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_this.f26).length));
        let _index96 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_759_v3).length));
        let _rhs88 = _756_v0;
        let _rhs89 = _757_v1;
        let _rhs90 = _module.__default.safeDivisionInt(_756_v0, _756_v0);
        let _rhs91 = _module.__default.fm28((_756_v0).isLessThan((_dafny.ZERO).minus(new BigNumber((_766_v10).cardinality()))), (_767_v11).IsDisjointFrom(_767_v11), (_758_v2).f27, _756_v0, globalState);
        let _lhs61 = _this.f26;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_this.f26).length));
        let _lhs63 = _759_v3;
        let _lhs64 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_759_v3).length));
        let _lhs65 = globalState;
        let _lhs66 = globalState;
        _lhs61[_lhs62] = _rhs88;
        _lhs63[_lhs64] = _rhs89;
        _lhs65.f8 = _rhs90;
        _lhs66.f2 = _rhs91;
        let _768_v12;
        let _nw111 = Array((new BigNumber(17)).toNumber()).fill(_module.D6.Default());
        _768_v12 = _nw111;
        let _769_v13;
        _769_v13 = _dafny.MultiSet.fromElements(false, _757_v1);
        let _770_v14;
        _770_v14 = new _dafny.CodePoint('w'.codePointAt(0));
        let _771_v15;
        _771_v15 = _module.D8.create_DC17(_770_v14, _770_v14, new BigNumber(891));
        let _772_v16;
        let _nw112 = Array((new BigNumber(16)).toNumber());
        _nw112[(_dafny.ZERO).toNumber()] = new BigNumber(536);
        _nw112[(_dafny.ONE).toNumber()] = (_this.f26)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_this.f26).length))];
        _nw112[(new BigNumber(2)).toNumber()] = new BigNumber(195);
        _nw112[(new BigNumber(3)).toNumber()] = (_this.f26)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_this.f26).length))];
        _nw112[(new BigNumber(4)).toNumber()] = _756_v0;
        _nw112[(new BigNumber(5)).toNumber()] = new BigNumber(740);
        _nw112[(new BigNumber(6)).toNumber()] = _756_v0;
        _nw112[(new BigNumber(7)).toNumber()] = _module.__default.fm18(_769_v13, _771_v15, (_758_v2).f27, globalState);
        _nw112[(new BigNumber(8)).toNumber()] = (_this.f26)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_this.f26).length))];
        _nw112[(new BigNumber(9)).toNumber()] = _756_v0;
        _nw112[(new BigNumber(10)).toNumber()] = _756_v0;
        _nw112[(new BigNumber(11)).toNumber()] = new BigNumber((_769_v13).cardinality());
        _nw112[(new BigNumber(12)).toNumber()] = _756_v0;
        _nw112[(new BigNumber(13)).toNumber()] = _756_v0;
        _nw112[(new BigNumber(14)).toNumber()] = _756_v0;
        _nw112[(new BigNumber(15)).toNumber()] = _756_v0;
        _772_v16 = _nw112;
        let _index97 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_768_v12).length));
        (_768_v12)[_index97] = _module.D6.create_DC10(_772_v16);
        let _773_v17;
        _773_v17 = _module.D6.create_DC10((((_759_v3)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_759_v3).length))]) ? (_772_v16) : (_772_v16)));
        let _index98 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_768_v12).length));
        (_768_v12)[_index98] = _773_v17;
      } else {
        let _774_v18;
        let _nw113 = new _module.C2();
        _nw113.__ctor();
        _774_v18 = _nw113;
        let _775_v19;
        _775_v19 = _dafny.Seq.UnicodeFromString("qujdwm");
        _775_v19 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-237)), function (_776_i0) {
          return ((false) ? (new _dafny.CodePoint('d'.codePointAt(0))) : (new _dafny.CodePoint('g'.codePointAt(0))));
        });
        (globalState).f8 = _756_v0;
        let _777_v20;
        _777_v20 = false;
        let _778_v21;
        let _nw114 = new _module.C0();
        _nw114.__ctor(_777_v20);
        _778_v21 = _nw114;
        let _779_v23;
        _779_v23 = new _dafny.CodePoint('n'.codePointAt(0));
        let _780_v24;
        _780_v24 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("qd"), _module.__default.safeIndex(_756_v0, new BigNumber((_dafny.Seq.UnicodeFromString("qd")).length)), _779_v23), _dafny.Seq.UnicodeFromString("sir"));
        let _781_v25;
        _781_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm7(globalState)).length),(_module.D4.create_DC7(_756_v0, _778_v21, _777_v20, new BigNumber((function () {
  let _coll29 = new _dafny.Map();
  for (const _compr_29 of (_780_v24).Elements) {
    let _782_v22 = _compr_29;
    if ((_780_v24).contains(_782_v22)) {
      _coll29.push([_782_v22,false]);
    }
  }
  return _coll29;
}()).length))).dtor_cf11);
        _781_v25 = (_781_v25).update(_756_v0, _777_v20);
        let _783_v26;
        _783_v26 = _dafny.Seq.of(_777_v20, _777_v20, (((_781_v25).contains(_756_v0)) ? ((_781_v25).get(_756_v0)) : (false)));
        (globalState).f17 = _module.__default.fm0(_783_v26, globalState);
      }
      let _784_v27;
      _784_v27 = true;
      let _785_v28;
      _785_v28 = _dafny.Map.Empty.slice().updateUnsafe(((_784_v27) ? (_756_v0) : (new BigNumber((_module.__default.fm7(globalState)).length))),_module.__default.fm5(globalState));
      let _786_v29;
      _786_v29 = _dafny.Seq.of(_784_v27, _784_v27);
      let _787_v30;
      _787_v30 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-615)), ((_788_v0) => function (_789_i1) {
        return _788_v0;
      })(_756_v0))));
      let _790_v31;
      _790_v31 = _dafny.Seq.of(new BigNumber((_786_v29).length), new BigNumber(((_787_v30)[_module.__default.safeIndex(_756_v0, new BigNumber((_787_v30).length))]).cardinality()), _module.__default.fm5(globalState), new BigNumber(173), _756_v0);
      let _791_v32;
      _791_v32 = _dafny.Map.Empty.slice().updateUnsafe(_790_v31,_dafny.Map.Empty.slice().updateUnsafe(_756_v0,_784_v27));
      let _792_v33;
      _792_v33 = _dafny.Map.Empty.slice().updateUnsafe(_756_v0,_784_v27);
      let _793_v35;
      _793_v35 = _dafny.MultiSet.fromElements((((_791_v32).contains(_790_v31)) ? ((_791_v32).get(_790_v31)) : (_792_v33)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-59),_784_v27), function () {
        let _coll30 = new _dafny.Map();
        for (const _compr_30 of (_dafny.Set.fromElements(_756_v0)).Elements) {
          let _794_v34 = _compr_30;
          if ((_dafny.Set.fromElements(_756_v0)).contains(_794_v34)) {
            _coll30.push([_module.__default.safeModuloInt(_794_v34, _756_v0),_784_v27]);
          }
        }
        return _coll30;
      }());
      _785_v28 = (_785_v28).update((((_793_v35).contains(_792_v33)) ? ((_793_v35).get(_792_v33)) : (_756_v0)), _756_v0);
      (globalState).f8 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(((_784_v27) ? (_756_v0) : (_756_v0)))).minus((_module.D10.create_DC22(_756_v0, _784_v27)).dtor_cf38));
      let _795_v36;
      let _nw115 = new _module.C2();
      _nw115.__ctor();
      _795_v36 = _nw115;
      let _hi5 = _756_v0;
      for (let _796_i2 = _756_v0; _796_i2.isLessThan(_hi5); _796_i2 = _796_i2.plus(_dafny.ONE)) {
        let _797_v37;
        _797_v37 = _dafny.Seq.UnicodeFromString("kxl");
        let _798_v38;
        _798_v38 = _dafny.Map.Empty.slice().updateUnsafe(_785_v28,_797_v37);
        _792_v33 = (_792_v33).update(((_784_v27) ? (new BigNumber((_798_v38).length)) : (_796_i2)), false);
        let _799_v39;
        let _800_v40;
        let _801_v41;
        let _802_v42;
        let _out27;
        let _out28;
        let _out29;
        let _out30;
        let _outcollector7 = _module.__default.m0(globalState);
        _out27 = _outcollector7[0];
        _out28 = _outcollector7[1];
        _out29 = _outcollector7[2];
        _out30 = _outcollector7[3];
        _799_v39 = _out27;
        _800_v40 = _out28;
        _801_v41 = _out29;
        _802_v42 = _out30;
        let _803_v43;
        let _init13 = function (_804_i3) {
          return false;
        };
        let _nw116 = Array((new BigNumber(5)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw116.length); _i0_13++) {
          _nw116[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _803_v43 = _nw116;
        _803_v43 = _803_v43;
        let _805_v44;
        let _nw117 = new _module.C0();
        _nw117.__ctor(_module.__default.fm0(_786_v29, globalState));
        _805_v44 = _nw117;
      }
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_this.f26).length))) {
        let _806_i4 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_806_i4)) && ((_806_i4).isLessThan(new BigNumber((_this.f26).length))))) {
          let _arr9 = _this.f26;
          _arr9[(_806_i4)] = _module.__default.safeModuloInt(_806_i4, _756_v0);
        }
      }
      return;
    }
    m2(globalState) {
      let _this = this;
      let _807_v0;
      let _init14 = function (_808_i0) {
        return false;
      };
      let _nw118 = Array((new BigNumber(4)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw118.length); _i0_14++) {
        _nw118[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _807_v0 = _nw118;
      let _809_v1;
      _809_v1 = false;
      let _810_v2;
      let _nw119 = Array((new BigNumber(28)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = _807_v0;
      _nw119[(_dafny.ONE).toNumber()] = _807_v0;
      _nw119[(new BigNumber(2)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(3)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(4)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(5)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(6)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(7)).toNumber()] = ((_809_v1) ? (_807_v0) : (_807_v0));
      _nw119[(new BigNumber(8)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(9)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(10)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(11)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(12)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(13)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(14)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(15)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(16)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(17)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(18)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(19)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(20)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(21)).toNumber()] = ((false) ? (_807_v0) : (_807_v0));
      _nw119[(new BigNumber(22)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(23)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(24)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(25)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(26)).toNumber()] = _807_v0;
      _nw119[(new BigNumber(27)).toNumber()] = _807_v0;
      _810_v2 = _nw119;
      _810_v2 = _810_v2;
      (globalState).f5 = ((_809_v1) ? (_809_v1) : ((_809_v1) && (_809_v1)));
      let _811_v3;
      _811_v3 = new BigNumber(487);
      let _812_v4;
      _812_v4 = _dafny.Seq.of(_811_v3, _811_v3);
      let _813_v5;
      _813_v5 = _dafny.MultiSet.fromElements(_811_v3, new BigNumber((_dafny.Seq.of(_809_v1)).length), (_812_v4)[_module.__default.safeIndex(_811_v3, new BigNumber((_812_v4).length))]);
      let _814_v6;
      _814_v6 = new _dafny.CodePoint('e'.codePointAt(0));
      let _815_v7;
      _815_v7 = _dafny.Map.Empty.slice().updateUnsafe(_809_v1,_813_v5);
      let _816_v8;
      _816_v8 = _dafny.Map.Empty.slice().updateUnsafe(_811_v3,_814_v6);
      let _817_v9;
      _817_v9 = _dafny.MultiSet.fromElements(true);
      let _818_v10;
      _818_v10 = _module.D8.create_DC17(_814_v6, _814_v6, _811_v3);
      let _819_v11;
      let _nw120 = new _module.C0();
      _nw120.__ctor(!(_809_v1));
      _819_v11 = _nw120;
      let _820_v12;
      _820_v12 = _dafny.Map.Empty.slice().updateUnsafe(_811_v3,_819_v11);
      let _821_v13;
      _821_v13 = _dafny.Map.Empty.slice().updateUnsafe((((_820_v12).contains(new BigNumber(-651))) ? ((_820_v12).get(new BigNumber(-651))) : (_819_v11)),_811_v3);
      let _pat_let_tv24 = _811_v3;
      let _822_v14;
      _822_v14 = _dafny.Seq.of((_813_v5).Intersect(_813_v5), _dafny.MultiSet.fromElements(_811_v3, _811_v3, new BigNumber(487), new BigNumber((_module.__default.fm21(_814_v6, _811_v3, globalState)).length)), (((_815_v7).contains(_809_v1)) ? ((_815_v7).get(_809_v1)) : (_module.__default.fm6(_812_v4, new BigNumber((_816_v8).length), _811_v3, globalState))), _dafny.MultiSet.fromElements(_811_v3), _module.__default.fm6(_dafny.Seq.of(_module.__default.fm18(_817_v9, function (_pat_let8_0) {
        return function (_823_dt__update__tmp_h0) {
          return function (_pat_let9_0) {
            return function (_824_dt__update_hcf32_h0) {
              return _module.D8.create_DC17((_823_dt__update__tmp_h0).dtor_cf30, (_823_dt__update__tmp_h0).dtor_cf31, _824_dt__update_hcf32_h0);
            }(_pat_let9_0);
          }(_pat_let_tv24);
        }(_pat_let8_0);
      }(_818_v10), _809_v1, globalState)), (((_821_v13).contains(_819_v11)) ? ((_821_v13).get(_819_v11)) : (_811_v3)), _811_v3, globalState));
      _822_v14 = _822_v14;
      let _825_i1;
      _825_i1 = _dafny.ZERO;
      L2: {
        while (_809_v1) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_825_i1)) {
              break L2;
            }
            _825_i1 = (_825_i1).plus(_dafny.ONE);
            let _826_v15;
            _826_v15 = _dafny.Seq.of(_this.f26, _this.f26, _this.f26, _this.f26, _this.f26);
            let _arr10 = _this.f26;
            let _index99 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_this.f26).length));
            _arr10[_index99] = _module.__default.safeModuloInt(_811_v3, _811_v3);
            let _arr11 = _this.f26;
            let _index100 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_this.f26).length));
            let _rhs92 = _826_v15;
            let _rhs93 = _811_v3;
            let _lhs67 = _this.f26;
            let _lhs68 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_this.f26).length));
            _826_v15 = _rhs92;
            _lhs67[_lhs68] = _rhs93;
            let _827_v16;
            _827_v16 = _dafny.Set.fromElements((_this.f26)[_module.__default.safeIndex(new BigNumber(751), new BigNumber((_this.f26).length))], _811_v3);
            let _828_v17;
            _828_v17 = _dafny.Seq.UnicodeFromString("jqlnl");
            let _829_v18;
            _829_v18 = _dafny.Map.Empty.slice().updateUnsafe(_827_v16,_828_v17);
            let _830_v19;
            let _nw121 = new _module.C1();
            _nw121.__ctor(_829_v18, _819_v11);
            _830_v19 = _nw121;
            _807_v0 = _807_v0;
            _827_v16 = (_827_v16).Difference(_827_v16);
          }
        }
      }
      let _831_v20;
      let _nw122 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
      _831_v20 = _nw122;
      _831_v20 = _831_v20;
      (globalState).f8 = _811_v3;
      return;
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f25 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor(f25) {
      let _this = this;
      (_this)._f25 = f25;
      return;
    }
    m1(globalState) {
      let _this = this;
      let _832_v0;
      _832_v0 = _dafny.Seq.of((_this).f25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(241)), function (_833_i1) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length), (_this).f25, (_this).f25, (_this).f25);
      let _834_v1;
      _834_v1 = new BigNumber((_832_v0).length);
      let _hi6 = (_this).f25;
      for (let _835_i0 = (_834_v1); _835_i0.isLessThan(_hi6); _835_i0 = _835_i0.plus(_dafny.ONE)) {
        let _836_v2;
        _836_v2 = true;
        let _rhs94 = _836_v2;
        let _rhs95 = _module.__default.safeDivisionInt(new BigNumber(-854), _835_i0);
        let _lhs69 = globalState;
        let _lhs70 = globalState;
        _lhs69.f17 = _rhs94;
        _lhs70.f8 = _rhs95;
        let _837_v3;
        let _nw123 = Array((new BigNumber(24)).toNumber()).fill(false);
        _837_v3 = _nw123;
        let _838_v4;
        _838_v4 = _dafny.Seq.of(_837_v3, _837_v3);
        let _839_v5;
        let _nw124 = Array((new BigNumber(18)).toNumber());
        _nw124[(_dafny.ZERO).toNumber()] = ((_836_v2) ? (_837_v3) : (_837_v3));
        _nw124[(_dafny.ONE).toNumber()] = _837_v3;
        _nw124[(new BigNumber(2)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(3)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(4)).toNumber()] = ((_836_v2) ? (_837_v3) : (_837_v3));
        _nw124[(new BigNumber(5)).toNumber()] = ((_836_v2) ? (_837_v3) : (_837_v3));
        _nw124[(new BigNumber(6)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(7)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(8)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(9)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(10)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(11)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(12)).toNumber()] = ((_836_v2) ? ((_838_v4)[_module.__default.safeIndex((_this).f25, new BigNumber((_838_v4).length))]) : (_837_v3));
        _nw124[(new BigNumber(13)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(14)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(15)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(16)).toNumber()] = _837_v3;
        _nw124[(new BigNumber(17)).toNumber()] = _837_v3;
        _839_v5 = _nw124;
        let _index101 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_839_v5).length));
        (_839_v5)[_index101] = _837_v3;
        let _840_v6;
        _840_v6 = _dafny.Seq.UnicodeFromString("upnyjowbd");
        let _841_v7;
        _841_v7 = new _dafny.CodePoint('t'.codePointAt(0));
        let _842_v8;
        let _nw125 = Array((new BigNumber(9)).toNumber());
        _nw125[(_dafny.ZERO).toNumber()] = _module.__default.fm1(globalState);
        _nw125[(_dafny.ONE).toNumber()] = _840_v6;
        _nw125[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_840_v6, _module.__default.safeIndex(_835_i0, new BigNumber((_840_v6).length)), _841_v7);
        _nw125[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_840_v6, _840_v6);
        _nw125[(new BigNumber(4)).toNumber()] = _840_v6;
        _nw125[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-241)), ((_843_v7) => function (_844_i2) {
          return _843_v7;
        })(_841_v7));
        _nw125[(new BigNumber(6)).toNumber()] = _840_v6;
        _nw125[(new BigNumber(7)).toNumber()] = _840_v6;
        _nw125[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_840_v6, _840_v6);
        _842_v8 = _nw125;
        let _index102 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_842_v8).length));
        (_842_v8)[_index102] = _840_v6;
        let _845_v9;
        _845_v9 = _dafny.Seq.of(!(_836_v2));
        let _index103 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_839_v5).length));
        let _index104 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_842_v8).length));
        let _rhs96 = _837_v3;
        let _rhs97 = _840_v6;
        let _rhs98 = _module.__default.fm0(_845_v9, globalState);
        let _lhs71 = _839_v5;
        let _lhs72 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_839_v5).length));
        let _lhs73 = _842_v8;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_842_v8).length));
        let _lhs75 = globalState;
        _lhs71[_lhs72] = _rhs96;
        _lhs73[_lhs74] = _rhs97;
        _lhs75.f22 = _rhs98;
        _836_v2 = _836_v2;
        let _846_v10;
        _846_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_835_i0);
        _846_v10 = (_846_v10).update(new BigNumber(66), new BigNumber((_832_v0).length));
      }
      let _847_v11;
      _847_v11 = false;
      let _848_v12;
      _848_v12 = _dafny.Set.fromElements(true, false, _847_v11);
      let _849_v13;
      _849_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_847_v11);
      let _850_v14;
      _850_v14 = _dafny.Seq.of(_847_v11, _847_v11);
      let _851_v15;
      let _nw126 = Array((new BigNumber(17)).toNumber());
      _nw126[(_dafny.ZERO).toNumber()] = _847_v11;
      _nw126[(_dafny.ONE).toNumber()] = _847_v11;
      _nw126[(new BigNumber(2)).toNumber()] = _847_v11;
      _nw126[(new BigNumber(3)).toNumber()] = _847_v11;
      _nw126[(new BigNumber(4)).toNumber()] = _847_v11;
      _nw126[(new BigNumber(5)).toNumber()] = true;
      _nw126[(new BigNumber(6)).toNumber()] = _847_v11;
      _nw126[(new BigNumber(7)).toNumber()] = (((_849_v13).contains((_dafny.ZERO).minus((_this).f25))) ? ((_849_v13).get((_dafny.ZERO).minus((_this).f25))) : (_847_v11));
      _nw126[(new BigNumber(8)).toNumber()] = _847_v11;
      _nw126[(new BigNumber(9)).toNumber()] = _847_v11;
      _nw126[(new BigNumber(10)).toNumber()] = _module.__default.fm0(_850_v14, globalState);
      _nw126[(new BigNumber(11)).toNumber()] = _847_v11;
      _nw126[(new BigNumber(12)).toNumber()] = _847_v11;
      _nw126[(new BigNumber(13)).toNumber()] = _847_v11;
      _nw126[(new BigNumber(14)).toNumber()] = false;
      _nw126[(new BigNumber(15)).toNumber()] = false;
      _nw126[(new BigNumber(16)).toNumber()] = _847_v11;
      _851_v15 = _nw126;
      let _852_v16;
      _852_v16 = _851_v15;
      let _853_v17;
      _853_v17 = _dafny.Seq.UnicodeFromString("na");
      let _854_v18;
      let _855_v19;
      let _856_v20;
      let _out31;
      let _out32;
      let _out33;
      let _outcollector8 = (_this).m3(_848_v12, _847_v11, (_851_v15), _dafny.Seq.Concat(_853_v17, _853_v17), globalState);
      _out31 = _outcollector8[0];
      _out32 = _outcollector8[1];
      _out33 = _outcollector8[2];
      _854_v18 = _out31;
      _855_v19 = _out32;
      _856_v20 = _out33;
      let _857_v21;
      _857_v21 = _dafny.MultiSet.fromElements((_this).f25);
      let _858_v22;
      _858_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_857_v21).Difference(_857_v21));
      _858_v22 = (_858_v22).Merge(_858_v22);
      (globalState).f19 = _847_v11;
      let _859_i3;
      _859_i3 = _dafny.ZERO;
      L3: {
        while (_854_v18) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_859_i3)) {
              break L3;
            }
            _859_i3 = (_859_i3).plus(_dafny.ONE);
            let _860_v24;
            _860_v24 = _dafny.Map.Empty.slice().updateUnsafe(function () {
              let _coll31 = new _dafny.Set();
              for (const _compr_31 of _dafny.IntegerRange(new BigNumber(138), new BigNumber(58))) {
                let _861_v23 = _compr_31;
                if (((new BigNumber(138)).isLessThanOrEqualTo(_861_v23)) && ((_861_v23).isLessThan(new BigNumber(58)))) {
                  _coll31.add((_861_v23).multipliedBy(new BigNumber((_850_v14).length)));
                }
              }
              return _coll31;
            }(),_dafny.Seq.UnicodeFromString("nttpwyv"));
            let _862_v25;
            let _nw127 = new _module.C0();
            _nw127.__ctor(_855_v19);
            _862_v25 = _nw127;
            let _863_v26;
            let _nw128 = new _module.C1();
            _nw128.__ctor((_860_v24).Merge(_860_v24), _862_v25);
            _863_v26 = _nw128;
            _863_v26 = _863_v26;
            let _index105 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_851_v15).length));
            (_851_v15)[_index105] = true;
            let _864_v27;
            _864_v27 = new _dafny.CodePoint('v'.codePointAt(0));
            let _865_v28;
            let _nw129 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
            _865_v28 = _nw129;
            let _866_v29;
            _866_v29 = _dafny.MultiSet.fromElements(_865_v28);
            let _867_v30;
            _867_v30 = _module.D8.create_DC17(_864_v27, _module.__default.fm22(false, globalState), ((_this).f25).minus(new BigNumber((_866_v29).cardinality())));
            let _868_v31;
            let _nw130 = new _module.C2();
            _nw130.__ctor();
            _868_v31 = _nw130;
            let _869_v32;
            _869_v32 = _dafny.MultiSet.fromElements(_868_v31, _868_v31, _868_v31);
            let _index106 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_851_v15).length));
            let _rhs99 = true;
            let _rhs100 = _module.__default.fm0(_dafny.Seq.Concat(_module.__default.fm12(_856_v20, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(939)), function (_870_i4) {
              return (_this).f25;
            })).length)), !(_854_v18), globalState), _850_v14), globalState);
            let _rhs101 = (((_869_v32).IsProperSubsetOf(_869_v32)) ? (_867_v30) : (_867_v30));
            let _rhs102 = _847_v11;
            let _lhs76 = _851_v15;
            let _lhs77 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_851_v15).length));
            let _lhs78 = globalState;
            _lhs76[_lhs77] = _rhs99;
            _lhs78.f17 = _rhs100;
            _867_v30 = _rhs101;
            _856_v20 = _rhs102;
            if (_854_v18) {
              let _index107 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length));
              (_865_v28)[_index107] = (_this).f25;
              let _871_v33;
              _871_v33 = _dafny.Seq.of(_832_v0);
              let _872_v34;
              _872_v34 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f25));
              let _873_v35;
              _873_v35 = _dafny.MultiSet.fromElements((_dafny.Set.fromElements(true, true)).IsSubsetOf(_848_v12), _847_v11, _855_v19, !(_872_v34).contains((_module.D13.create_DC28((_871_v33)[_module.__default.safeIndex(new BigNumber((_850_v14).length), new BigNumber((_871_v33).length))])).dtor_cf47));
              let _index108 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_851_v15).length));
              let _index109 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length));
              let _rhs103 = new BigNumber((_853_v17).length);
              let _rhs104 = _854_v18;
              let _rhs105 = (((_873_v35).contains(false)) ? ((_873_v35).get(false)) : (_module.__default.fm18(_873_v35, _867_v30, _856_v20, globalState)));
              let _rhs106 = ((_this).f25).multipliedBy((_this).f25);
              let _lhs79 = globalState;
              let _lhs80 = _851_v15;
              let _lhs81 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_851_v15).length));
              let _lhs82 = _865_v28;
              let _lhs83 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length));
              let _lhs84 = globalState;
              _lhs79.f8 = _rhs103;
              _lhs80[_lhs81] = _rhs104;
              _lhs82[_lhs83] = _rhs105;
              _lhs84.f8 = _rhs106;
              let _874_v36;
              let _init15 = function (_875_i5) {
                return (_875_i5).minus(new BigNumber((_dafny.Seq.UnicodeFromString("trutx")).length));
              };
              let _nw131 = Array((new BigNumber(26)).toNumber());
              for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw131.length); _i0_15++) {
                _nw131[_i0_15] = _init15(new BigNumber(_i0_15));
              }
              _874_v36 = _nw131;
              let _876_v37;
              _876_v37 = _dafny.Set.fromElements(_874_v36, _874_v36, _874_v36);
              let _877_v38;
              _877_v38 = _876_v37;
              let _878_v39;
              _878_v39 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),_876_v37);
              let _879_v40;
              _879_v40 = _dafny.Map.Empty.slice().updateUnsafe((_865_v28)[_module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length))],(((_878_v39).contains(new _dafny.CodePoint('j'.codePointAt(0)))) ? ((_878_v39).get(new _dafny.CodePoint('j'.codePointAt(0)))) : (_876_v37)));
              let _880_v41;
              let _nw132 = Array((new BigNumber(25)).toNumber());
              _nw132[(_dafny.ZERO).toNumber()] = (_876_v37).Intersect(_876_v37);
              _nw132[(_dafny.ONE).toNumber()] = (_876_v37).Intersect((_877_v38));
              _nw132[(new BigNumber(2)).toNumber()] = (_dafny.Set.fromElements(_874_v36, _874_v36)).Intersect(_876_v37);
              _nw132[(new BigNumber(3)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(4)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(5)).toNumber()] = (_876_v37).Intersect(_dafny.Set.fromElements(_874_v36));
              _nw132[(new BigNumber(6)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(7)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_874_v36, _874_v36, _874_v36);
              _nw132[(new BigNumber(9)).toNumber()] = ((((_879_v40).contains((_865_v28)[_module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length))])) ? ((_879_v40).get((_865_v28)[_module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length))])) : (_876_v37))).Intersect(_dafny.Set.fromElements(_874_v36, _874_v36, _874_v36, _874_v36, _874_v36));
              _nw132[(new BigNumber(10)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(11)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(12)).toNumber()] = ((_854_v18) ? (_876_v37) : (_876_v37));
              _nw132[(new BigNumber(13)).toNumber()] = (_dafny.Set.fromElements(_874_v36)).Intersect(_876_v37);
              _nw132[(new BigNumber(14)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(15)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(16)).toNumber()] = (_876_v37).Difference(_876_v37);
              _nw132[(new BigNumber(17)).toNumber()] = (_876_v37).Union(_876_v37);
              _nw132[(new BigNumber(18)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(19)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(20)).toNumber()] = _dafny.Set.fromElements(_874_v36, _874_v36);
              _nw132[(new BigNumber(21)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(22)).toNumber()] = _876_v37;
              _nw132[(new BigNumber(23)).toNumber()] = (_876_v37).Intersect(_876_v37);
              _nw132[(new BigNumber(24)).toNumber()] = ((_854_v18) ? (_876_v37) : (_876_v37));
              _880_v41 = _nw132;
              let _index110 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_880_v41).length));
              (_880_v41)[_index110] = _876_v37;
              let _881_v42;
              let _init16 = ((_882_v19) => function (_883_i6) {
                return _882_v19;
              })(_855_v19);
              let _nw133 = Array((new BigNumber(21)).toNumber());
              for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw133.length); _i0_16++) {
                _nw133[_i0_16] = _init16(new BigNumber(_i0_16));
              }
              _881_v42 = _nw133;
              let _884_v43;
              _884_v43 = _dafny.Map.Empty.slice().updateUnsafe(_881_v42,_847_v11);
              let _885_v44;
              let _nw134 = Array((new BigNumber(8)).toNumber());
              _nw134[(_dafny.ZERO).toNumber()] = ((!(!(_module.__default.fm0(_850_v14, globalState)))) ? (_884_v43) : ((_884_v43).update(_881_v42, _854_v18)));
              _nw134[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_881_v42,_856_v20);
              _nw134[(new BigNumber(2)).toNumber()] = _884_v43;
              _nw134[(new BigNumber(3)).toNumber()] = (_884_v43).Merge(_dafny.Map.Empty.slice().updateUnsafe(_881_v42,_854_v18));
              _nw134[(new BigNumber(4)).toNumber()] = _884_v43;
              _nw134[(new BigNumber(5)).toNumber()] = _884_v43;
              _nw134[(new BigNumber(6)).toNumber()] = _884_v43;
              _nw134[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_881_v42,_855_v19);
              _885_v44 = _nw134;
              let _index111 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_885_v44).length));
              (_885_v44)[_index111] = _884_v43;
              let _index112 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_880_v41).length));
              let _index113 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_885_v44).length));
              let _rhs107 = ((_856_v20) ? (_876_v37) : (_876_v37));
              let _rhs108 = (_884_v43).Merge(_884_v43);
              let _lhs85 = _880_v41;
              let _lhs86 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_880_v41).length));
              let _lhs87 = _885_v44;
              let _lhs88 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_885_v44).length));
              _lhs85[_lhs86] = _rhs107;
              _lhs87[_lhs88] = _rhs108;
              let _886_v45;
              _886_v45 = _dafny.Map.Empty.slice().updateUnsafe(_874_v36,(_865_v28)[_module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length))]);
              let _index114 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length));
              let _rhs109 = ((((_857_v21).contains((_this).f25)) ? ((_857_v21).get((_this).f25)) : ((_865_v28)[_module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length))]))).minus((_865_v28)[_module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length))]);
              let _rhs110 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_853_v17).length), new BigNumber((_886_v45).length)), ((_855_v19) ? ((_865_v28)[_module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length))]) : (new BigNumber(180)))));
              let _lhs89 = globalState;
              let _lhs90 = _865_v28;
              let _lhs91 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length));
              _lhs89.f8 = _rhs109;
              _lhs90[_lhs91] = _rhs110;
              _864_v27 = ((_847_v11) ? (_864_v27) : (_864_v27));
              let _index115 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length));
              (_865_v28)[_index115] = _module.__default.safeModuloInt(((_865_v28)[_module.__default.safeIndex(new BigNumber(155), new BigNumber((_865_v28).length))]).multipliedBy((_this).f25), (_this).f25);
            } else {
              let _887_v46;
              _887_v46 = _module.D6.create_DC10(_865_v28);
              (globalState).f23 = (_887_v46).dtor_cf17;
              let _888_v47;
              _888_v47 = _dafny.MultiSet.fromElements(true, ((_dafny.ZERO).minus((_this).f25)).isEqualTo((_this).f25));
              _888_v47 = _888_v47;
              let _889_v48;
              _889_v48 = _module.D3.create_DC4(new BigNumber(372), _855_v19, _856_v20);
              let _890_v49;
              _890_v49 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm18(_dafny.MultiSet.fromElements((_851_v15)[_module.__default.safeIndex(new BigNumber(274), new BigNumber((_851_v15).length))], _855_v19, _854_v18, !(_854_v18), _847_v11), _867_v30, true, globalState),_889_v48);
              let _891_v50;
              _891_v50 = _dafny.Seq.of(_890_v49, _890_v49, _890_v49);
              let _892_v51;
              _892_v51 = _dafny.Map.Empty.slice().updateUnsafe(_834_v1,(_this).f25);
              let _index116 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_851_v15).length));
              let _rhs111 = (((_851_v15)[_module.__default.safeIndex(new BigNumber(274), new BigNumber((_851_v15).length))]) ? (_module.__default.fm22(_854_v18, globalState)) : (_864_v27));
              let _rhs112 = true;
              let _rhs113 = _module.__default.safeDivisionInt(new BigNumber((_832_v0).length), _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f25), (_this).f25));
              let _rhs114 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_891_v50, _module.__default.safeIndex((((_892_v51).contains(_834_v1)) ? ((_892_v51).get(_834_v1)) : (new BigNumber(563))), new BigNumber((_891_v50).length)), _890_v49), _891_v50);
              let _lhs92 = globalState;
              let _lhs93 = globalState;
              let _lhs94 = _851_v15;
              let _lhs95 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_851_v15).length));
              _864_v27 = _rhs111;
              _lhs92.f17 = _rhs112;
              _lhs93.f8 = _rhs113;
              _lhs94[_lhs95] = _rhs114;
              (globalState).f19 = !(_854_v18);
              let _893_v52;
              let _nw135 = Array((new BigNumber(17)).toNumber()).fill(false);
              _893_v52 = _nw135;
              let _894_v53;
              let _nw136 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _894_v53 = _nw136;
              let _895_v54;
              let _nw137 = new _module.C1();
              _nw137.__ctor(_860_v24, _862_v25);
              _895_v54 = _nw137;
              let _896_v55;
              _896_v55 = _dafny.Seq.of(_895_v54);
              let _897_v56;
              _897_v56 = _dafny.Map.Empty.slice().updateUnsafe(_854_v18,_896_v55);
              let _rhs115 = !(_847_v11);
              let _rhs116 = new BigNumber((_dafny.Set.fromElements(_894_v53)).length);
              let _rhs117 = _865_v28;
              let _rhs118 = _893_v52;
              let _rhs119 = ((_856_v20) ? (new BigNumber((_897_v56).length)) : ((_this).f25));
              let _lhs96 = globalState;
              let _lhs97 = globalState;
              let _lhs98 = globalState;
              _lhs96.f17 = _rhs115;
              _lhs97.f8 = _rhs116;
              _865_v28 = _rhs117;
              _893_v52 = _rhs118;
              _lhs98.f8 = _rhs119;
            }
            let _898_v57;
            _898_v57 = _dafny.MultiSet.fromElements(_module.__default.fm0(_850_v14, globalState), _854_v18);
            let _899_v58;
            _899_v58 = _dafny.Set.fromElements((((_857_v21).contains((_this).f25)) ? ((_857_v21).get((_this).f25)) : (new BigNumber(635))));
            (globalState).f8 = (((_898_v57).contains((_899_v58).IsDisjointFrom(_899_v58))) ? ((_898_v57).get((_899_v58).IsDisjointFrom(_899_v58))) : ((_this).f25));
          }
        }
      }
      let _hi7 = (_this).f25;
      for (let _900_i7 = (_this).f25; _900_i7.isLessThan(_hi7); _900_i7 = _900_i7.plus(_dafny.ONE)) {
        (globalState).f22 = _856_v20;
        let _901_v59;
        let _nw138 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _901_v59 = _nw138;
        let _index117 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_901_v59).length));
        (_901_v59)[_index117] = (_this).f25;
        let _index118 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_901_v59).length));
        (_901_v59)[_index118] = (_this).f25;
        let _902_v60;
        _902_v60 = _dafny.Map.Empty.slice().updateUnsafe(_854_v18,_856_v20);
        let _903_v61;
        let _904_v62;
        let _905_v63;
        let _906_v64;
        let _out34;
        let _out35;
        let _out36;
        let _out37;
        let _outcollector9 = (_this).m4(_dafny.Seq.of((_this).f25), _902_v60, _847_v11, globalState);
        _out34 = _outcollector9[0];
        _out35 = _outcollector9[1];
        _out36 = _outcollector9[2];
        _out37 = _outcollector9[3];
        _903_v61 = _out34;
        _904_v62 = _out35;
        _905_v63 = _out36;
        _906_v64 = _out37;
        _856_v20 = _905_v63;
      }
      return;
    }
    m2(globalState) {
      let _this = this;
      let _907_v0;
      _907_v0 = true;
      let _908_v1;
      _908_v1 = _dafny.Map.Empty.slice().updateUnsafe(_907_v0,!(_907_v0));
      let _909_v2;
      _909_v2 = _dafny.Map.Empty.slice().updateUnsafe(_907_v0,!(_907_v0) || ((((_908_v1).contains(_907_v0)) ? ((_908_v1).get(_907_v0)) : (_907_v0))));
      _908_v1 = (_908_v1).update(_907_v0, !(_907_v0));
      let _910_v3;
      _910_v3 = _dafny.Seq.UnicodeFromString("ymupw");
      let _911_v4;
      _911_v4 = _dafny.MultiSet.fromElements(_907_v0, _907_v0);
      let _912_v5;
      _912_v5 = _dafny.Seq.of((_this).f25, (_this).f25, new BigNumber((_911_v4).cardinality()));
      let _913_v6;
      _913_v6 = _dafny.Map.Empty.slice().updateUnsafe(_910_v3,_912_v5);
      let _914_v7;
      _914_v7 = new _dafny.CodePoint('x'.codePointAt(0));
      _913_v6 = (_913_v6).update(_dafny.Seq.Concat(_910_v3, _dafny.Seq.update(_910_v3, _module.__default.safeIndex((_this).f25, new BigNumber((_910_v3).length)), _914_v7)), _912_v5);
      let _915_v8;
      _915_v8 = _dafny.Seq.of(_910_v3);
      let _916_v9;
      _916_v9 = _module.D3.create_DC3(_915_v8);
      let _917_v10;
      _917_v10 = _module.D3.create_DC5(_916_v9);
      let _918_i0;
      _918_i0 = _dafny.ZERO;
      L4: {
        let _pat_let_tv25 = _907_v0;
        while (function (_source15) {
          if (_source15.is_DC4) {
            let _940___mcc_h0 = (_source15).cf4;
            let _941___mcc_h1 = (_source15).cf5;
            let _942___mcc_h2 = (_source15).cf6;
            let _943_cf6 = _942___mcc_h2;
            let _944_cf5 = _941___mcc_h1;
            let _945_cf4 = _940___mcc_h0;
            return !(_943_cf6) || (_943_cf6);
          } else if (_source15.is_DC3) {
            let _946___mcc_h3 = (_source15).cf3;
            let _947_cf3 = _946___mcc_h3;
            return !(_pat_let_tv25);
          } else {
            let _948___mcc_h4 = (_source15).cf7;
            let _949_cf7 = _948___mcc_h4;
            return ((_this).f25).isLessThanOrEqualTo((_this).f25);
          }
        }(_917_v10)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_918_i0)) {
              break L4;
            }
            _918_i0 = (_918_i0).plus(_dafny.ONE);
            _909_v2 = (_909_v2).update(!(_907_v0), !(_dafny.Seq.IsPrefixOf(_dafny.Seq.of((_this).f25, (_this).f25), _912_v5)));
            let _919_v11;
            let _nw139 = Array((new BigNumber(12)).toNumber());
            _nw139[(_dafny.ZERO).toNumber()] = _module.__default.fm1(globalState);
            _nw139[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(288)), ((_920_v7) => function (_921_i1) {
              return _920_v7;
            })(_914_v7));
            _nw139[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("fvis");
            _nw139[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_910_v3, _910_v3);
            _nw139[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(506)), ((_922_v7) => function (_923_i2) {
              return _922_v7;
            })(_914_v7));
            _nw139[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_910_v3, _module.__default.safeIndex((_this).f25, new BigNumber((_910_v3).length)), _914_v7);
            _nw139[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("qknenuyap");
            _nw139[(new BigNumber(7)).toNumber()] = _910_v3;
            _nw139[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-977)), ((_924_v7) => function (_925_i3) {
              return _924_v7;
            })(_914_v7));
            _nw139[(new BigNumber(9)).toNumber()] = _910_v3;
            _nw139[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(98)), ((_926_v7) => function (_927_i4) {
              return _926_v7;
            })(_914_v7));
            _nw139[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("s");
            _919_v11 = _nw139;
            let _index119 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_919_v11).length));
            (_919_v11)[_index119] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(837)), ((_928_v7) => function (_929_i5) {
              return _928_v7;
            })(_914_v7)), _910_v3);
            let _index120 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_919_v11).length));
            (_919_v11)[_index120] = _dafny.Seq.Concat(_910_v3, _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-410)), function (_930_i6) {
              return new _dafny.CodePoint('l'.codePointAt(0));
            }), _module.__default.safeIndex((_this).f25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-410)), function (_931_i6) {
              return new _dafny.CodePoint('l'.codePointAt(0));
            })).length)), _914_v7), _910_v3));
            let _932_v12;
            let _init17 = function (_933_i7) {
              return (_dafny.Set.fromElements((_this).f25)).Union(_dafny.Set.fromElements((_this).f25, (_this).f25));
            };
            let _nw140 = Array((new BigNumber(21)).toNumber());
            for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw140.length); _i0_17++) {
              _nw140[_i0_17] = _init17(new BigNumber(_i0_17));
            }
            _932_v12 = _nw140;
            let _934_v13;
            _934_v13 = _dafny.Set.fromElements((_this).f25, (_this).f25, (_this).f25);
            let _935_v14;
            _935_v14 = _dafny.Seq.of(_934_v13);
            let _index121 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_932_v12).length));
            (_932_v12)[_index121] = ((_935_v14)[_module.__default.safeIndex(new BigNumber(403), new BigNumber((_935_v14).length))]).Union(_dafny.Set.fromElements((_this).f25));
            let _index122 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_932_v12).length));
            (_932_v12)[_index122] = _934_v13;
            let _936_v15;
            let _init18 = ((_937_v0) => function (_938_i8) {
              return (_938_i8).multipliedBy(new BigNumber((_dafny.Seq.of(_937_v0, _937_v0)).length));
            })(_907_v0);
            let _nw141 = Array((new BigNumber(16)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw141.length); _i0_18++) {
              _nw141[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _936_v15 = _nw141;
            let _939_v16;
            let _nw142 = new _module.C3();
            _nw142.__ctor(((!(_907_v0)) ? (_936_v15) : (_936_v15)));
            _939_v16 = _nw142;
          }
        }
      }
      (globalState).f17 = _dafny.Seq.contains(_912_v5, _module.__default.safeModuloInt((_this).f25, (_this).f25));
      let _950_v17;
      let _init19 = function (_951_i9) {
        return _module.__default.safeModuloInt(_951_i9, (_this).f25);
      };
      let _nw143 = Array((new BigNumber(5)).toNumber());
      for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw143.length); _i0_19++) {
        _nw143[_i0_19] = _init19(new BigNumber(_i0_19));
      }
      _950_v17 = _nw143;
      (globalState).f23 = _950_v17;
      let _952_v18;
      _952_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
      let _953_v19;
      _953_v19 = _dafny.Seq.of(_907_v0);
      let _954_v20;
      _954_v20 = _dafny.Map.Empty.slice().updateUnsafe(_952_v18,_953_v19);
      let _955_v21;
      let _nw144 = Array((new BigNumber(5)).toNumber());
      _nw144[(_dafny.ZERO).toNumber()] = (true) && (true);
      _nw144[(_dafny.ONE).toNumber()] = _module.__default.fm0((((_954_v20).contains(_952_v18)) ? ((_954_v20).get(_952_v18)) : (_953_v19)), globalState);
      _nw144[(new BigNumber(2)).toNumber()] = _907_v0;
      _nw144[(new BigNumber(3)).toNumber()] = _907_v0;
      _nw144[(new BigNumber(4)).toNumber()] = !(_907_v0);
      _955_v21 = _nw144;
      let _index123 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_955_v21).length));
      (_955_v21)[_index123] = !(_907_v0);
      let _index124 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_955_v21).length));
      (_955_v21)[_index124] = _module.__default.fm0(_dafny.Seq.Concat(_953_v19, _953_v19), globalState);
      return;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let _956_v0;
      _956_v0 = _dafny.Seq.of(true);
      let _957_v1;
      _957_v1 = _dafny.Seq.of(_module.__default.fm0(_956_v0, globalState));
      (globalState).f8 = new BigNumber((_dafny.Seq.Concat(_957_v1, _dafny.Seq.Concat(_dafny.Seq.of(p1, _module.__default.fm0(_956_v0, globalState)), _956_v0))).length);
      let _hi8 = (_this).f25;
      for (let _958_i0 = (_this).f25; _958_i0.isLessThan(_hi8); _958_i0 = _958_i0.plus(_dafny.ONE)) {
        let _959_v2;
        let _nw145 = new _module.C0();
        _nw145.__ctor(p1);
        _959_v2 = _nw145;
        let _960_v3;
        _960_v3 = _dafny.MultiSet.fromElements(p1, p1, (_959_v2).f27, (_959_v2).f27);
        _960_v3 = (_960_v3).Union(_960_v3);
        let _961_v5;
        let _init20 = ((_962_i0) => function (_963_i1) {
          return (_dafny.Set.fromElements(_dafny.Seq.of(_962_i0, (_dafny.ZERO).minus((_this).f25)))).Intersect(function () {
            let _coll32 = new _dafny.Set();
            for (const _compr_32 of (_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(798)), ((_964_i0) => function (_965_i2) {
              return (_dafny.ZERO).minus(_964_i0);
            })(_962_i0)), _dafny.Seq.of(_962_i0, (_this).f25))).Elements) {
              let _966_v4 = _compr_32;
              if ((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(798)), ((_967_i0) => function (_965_i2) {
                return (_dafny.ZERO).minus(_967_i0);
              })(_962_i0)), _dafny.Seq.of(_962_i0, (_this).f25))).contains(_966_v4)) {
                _coll32.add(_966_v4);
              }
            }
            return _coll32;
          }());
        })(_958_i0);
        let _nw146 = Array((new BigNumber(14)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw146.length); _i0_20++) {
          _nw146[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _961_v5 = _nw146;
        let _968_v6;
        _968_v6 = _dafny.Seq.of(_958_i0);
        let _969_v7;
        _969_v7 = _module.D13.create_DC28(_968_v6);
        let _970_v8;
        _970_v8 = _dafny.Set.fromElements((_969_v7).dtor_cf47, _dafny.Seq.of((_this).f25, _958_i0), _dafny.Seq.update(_968_v6, _module.__default.safeIndex(_958_i0, new BigNumber((_968_v6).length)), _958_i0), _968_v6);
        let _index125 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_961_v5).length));
        (_961_v5)[_index125] = (_970_v8).Difference(_dafny.Set.fromElements(_968_v6, _968_v6));
        let _971_v9;
        _971_v9 = _dafny.Seq.of(_968_v6);
        let _index126 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_961_v5).length));
        (_961_v5)[_index126] = function () {
          let _coll33 = new _dafny.Set();
          for (const _compr_33 of (_971_v9).Elements) {
            let _972_v10 = _compr_33;
            if (_dafny.Seq.contains(_971_v9, _972_v10)) {
              _coll33.add(_972_v10);
            }
          }
          return _coll33;
        }();
        let _973_v11;
        _973_v11 = new _dafny.CodePoint('k'.codePointAt(0));
        let _974_v12;
        _974_v12 = _module.D8.create_DC17(new _dafny.CodePoint('c'.codePointAt(0)), _973_v11, new BigNumber((p3).length));
        let _975_v13;
        _975_v13 = _module.D12.create_DC27((_959_v2).f27, (_this).f25);
        (globalState).f8 = _module.__default.fm18(_960_v3, _974_v12, !((_975_v13).dtor_cf46).isEqualTo((_this).f25), globalState);
      }
      let _976_v14;
      _976_v14 = _dafny.Set.fromElements((_this).f25);
      let _977_v15;
      _977_v15 = new _dafny.CodePoint('j'.codePointAt(0));
      let _index127 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((p2).length));
      (p2)[_index127] = p1;
      let _978_v16;
      _978_v16 = _dafny.Seq.of(_976_v14, _976_v14);
      let _979_v17;
      let _nw147 = new _module.C0();
      _nw147.__ctor(p1);
      _979_v17 = _nw147;
      let _980_v18;
      _980_v18 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f25),_979_v17);
      let _981_v19;
      _981_v19 = _dafny.Set.fromElements((((_980_v18).contains((_this).f25)) ? ((_980_v18).get((_this).f25)) : (_979_v17)));
      let _index128 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((p2).length));
      let _rhs120 = (_978_v16)[_module.__default.safeIndex(new BigNumber(((_dafny.Set.fromElements(_979_v17, _979_v17)).Union(_981_v19)).length), new BigNumber((_978_v16).length))];
      let _rhs121 = _module.__default.fm8(globalState);
      let _rhs122 = (new BigNumber(980)).isLessThan((_this).f25);
      let _rhs123 = !((_979_v17).f27);
      let _lhs99 = globalState;
      let _lhs100 = p2;
      let _lhs101 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((p2).length));
      _976_v14 = _rhs120;
      _977_v15 = _rhs121;
      _lhs99.f19 = _rhs122;
      _lhs100[_lhs101] = _rhs123;
      let _982_v20;
      let _init21 = function (_983_i3) {
        return (_983_i3).plus(new BigNumber((_dafny.Seq.of((_this).f25)).length));
      };
      let _nw148 = Array((new BigNumber(23)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw148.length); _i0_21++) {
        _nw148[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _982_v20 = _nw148;
      let _index129 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_982_v20).length));
      (_982_v20)[_index129] = new BigNumber((_956_v0).length);
      let _index130 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_982_v20).length));
      (_982_v20)[_index130] = ((_this).f25).plus(_module.__default.fm5(globalState));
      let _984_v21;
      _984_v21 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_982_v20)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_982_v20).length))]);
      let _hi9 = new BigNumber((_984_v21).length);
      for (let _985_i4 = (_this).f25; _985_i4.isLessThan(_hi9); _985_i4 = _985_i4.plus(_dafny.ONE)) {
        let _986_v22;
        _986_v22 = _module.D13.create_DC28(_dafny.Seq.of((_this).f25, (_982_v20)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_982_v20).length))], new BigNumber((p3).length)));
        _986_v22 = _986_v22;
        (globalState).f8 = (_982_v20)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_982_v20).length))];
        let _987_v23;
        let _nw149 = new _module.C0();
        _nw149.__ctor(true);
        _987_v23 = _nw149;
        let _988_v24;
        _988_v24 = _dafny.Map.Empty.slice().updateUnsafe(_976_v14,_dafny.Seq.UnicodeFromString("p"));
        let _989_v25;
        let _nw150 = new _module.C0();
        _nw150.__ctor(p1);
        _989_v25 = _nw150;
        let _990_v26;
        let _nw151 = new _module.C1();
        _nw151.__ctor(_988_v24, _989_v25);
        _990_v26 = _nw151;
        let _991_v27;
        _991_v27 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm0(_956_v0, globalState));
        let _992_v28;
        _992_v28 = _dafny.Map.Empty.slice().updateUnsafe(_990_v26,(((_979_v17).f27) ? (_dafny.Map.Empty.slice().updateUnsafe(!((p2)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((p2).length))]),(p2)[_module.__default.safeIndex(new BigNumber(145), new BigNumber((p2).length))])) : (_991_v27)));
        _992_v28 = (_992_v28).update(_990_v26, _991_v27);
      }
      let _993_v29;
      let _994_v30;
      let _995_v31;
      let _996_v32;
      let _out38;
      let _out39;
      let _out40;
      let _out41;
      let _outcollector10 = _module.__default.m0(globalState);
      _out38 = _outcollector10[0];
      _out39 = _outcollector10[1];
      _out40 = _outcollector10[2];
      _out41 = _outcollector10[3];
      _993_v29 = _out38;
      _994_v30 = _out39;
      _995_v31 = _out40;
      _996_v32 = _out41;
      r0 = _module.__default.fm0(_956_v0, globalState);
      r1 = _995_v31;
      r2 = p1;
      return [r0, r1, r2];
    }
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _997_v0;
      _997_v0 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("oapapxtbs"),p2);
      (globalState).f19 = ((p2) ? ((((_997_v0).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(918)), function (_999_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }))) ? ((_997_v0).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(918)), function (_998_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }))) : (p2))) : (false));
      let _1000_v1;
      _1000_v1 = _dafny.Seq.UnicodeFromString("ahhoqv");
      let _rhs124 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(25)), function (_1001_i1) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _1000_v1);
      let _rhs125 = !(true) || (p2);
      _1000_v1 = _rhs124;
      r2 = _rhs125;
      let _1002_v2;
      _1002_v2 = new _dafny.CodePoint('b'.codePointAt(0));
      let _1003_v3;
      _1003_v3 = _dafny.Seq.of(p2);
      let _1004_v4;
      _1004_v4 = _module.D4.create_DC8(((!(p2)) ? (_1002_v2) : (_1002_v2)), _module.__default.fm0(_1003_v3, globalState), p2);
      let _source16 = _1004_v4;
      if (_source16.is_DC7) {
        let _1005___mcc_h0 = (_source16).cf9;
        let _1006___mcc_h1 = (_source16).cf10;
        let _1007___mcc_h2 = (_source16).cf11;
        let _1008___mcc_h3 = (_source16).cf12;
        let _1009_cf12 = _1008___mcc_h3;
        let _1010_cf11 = _1007___mcc_h2;
        let _1011_cf10 = _1006___mcc_h1;
        let _1012_cf9 = _1005___mcc_h0;
        (globalState).f19 = _1010_cf11;
        let _1013_v5;
        let _nw152 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _1013_v5 = _nw152;
        let _index131 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1013_v5).length));
        (_1013_v5)[_index131] = new BigNumber((_1000_v1).length);
        let _1014_v6;
        _1014_v6 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1002_v2);
        let _index132 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_1013_v5).length));
        (_1013_v5)[_index132] = (new BigNumber(((_1014_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,_1002_v2))).length)).multipliedBy(_1009_cf12);
        let _1015_v7;
        _1015_v7 = (_1013_v5)[_module.__default.safeIndex(new BigNumber(456), new BigNumber((_1013_v5).length))];
        let _1016_v8;
        _1016_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1010_cf11,_1015_v7);
        (globalState).f17 = !(_1016_v8).contains(p2);
        let _1017_v9;
        _1017_v9 = _module.D6.create_DC10(_1013_v5);
        let _1018_v10;
        _1018_v10 = _module.D6.create_DC12(_module.D6.create_DC12(_1017_v9));
        let _1019_v11;
        _1019_v11 = _module.D6.create_DC12(_1017_v9);
        let _1020_v12;
        _1020_v12 = _module.D6.create_DC12(_module.D6.create_DC12(_1017_v9));
        let _1021_v13;
        let _nw153 = new _module.C0();
        _nw153.__ctor((_1011_cf10).f27);
        _1021_v13 = _nw153;
        let _1022_v14;
        let _nw154 = Array((new BigNumber(28)).toNumber()).fill(false);
        _1022_v14 = _nw154;
        let _1023_v15;
        let _nw155 = Array((new BigNumber(2)).toNumber());
        _nw155[(_dafny.ZERO).toNumber()] = _1022_v14;
        _nw155[(_dafny.ONE).toNumber()] = _1022_v14;
        _1023_v15 = _nw155;
        let _index133 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1023_v15).length));
        (_1023_v15)[_index133] = _1022_v14;
        let _index134 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1023_v15).length));
        let _rhs126 = _module.D6.create_DC12(_1018_v10);
        let _rhs127 = _1021_v13;
        let _rhs128 = _1013_v5;
        let _rhs129 = _1022_v14;
        let _rhs130 = _dafny.Seq.Concat(_1000_v1, _1000_v1);
        let _lhs102 = _1023_v15;
        let _lhs103 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1023_v15).length));
        _1020_v12 = _rhs126;
        _1021_v13 = _rhs127;
        _1013_v5 = _rhs128;
        _lhs102[_lhs103] = _rhs129;
        _1000_v1 = _rhs130;
      } else if (_source16.is_DC8) {
        let _1024___mcc_h4 = (_source16).cf13;
        let _1025___mcc_h5 = (_source16).cf14;
        let _1026___mcc_h6 = (_source16).cf15;
        let _1027_cf15 = _1026___mcc_h6;
        let _1028_cf14 = _1025___mcc_h5;
        let _1029_cf13 = _1024___mcc_h4;
        let _1030_v16;
        _1030_v16 = _module.D7.create_DC15(new BigNumber(747), _1027_cf15, (_this).f25);
        let _1031_v17;
        _1031_v17 = _dafny.Set.fromElements(p2);
        let _1032_v18;
        _1032_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1031_v17).length),_1028_cf14);
        let _1033_v19;
        _1033_v19 = _module.D12.create_DC27(p2, new BigNumber(-214));
        let _1034_v20;
        _1034_v20 = _dafny.MultiSet.fromElements(_1029_cf13, _1002_v2);
        let _1035_v21;
        let _nw156 = Array((new BigNumber(27)).toNumber());
        _nw156[(_dafny.ZERO).toNumber()] = true;
        _nw156[(_dafny.ONE).toNumber()] = ((_1027_cf15) ? (p2) : (_1027_cf15));
        _nw156[(new BigNumber(2)).toNumber()] = true;
        _nw156[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_1000_v1, _1000_v1);
        _nw156[(new BigNumber(4)).toNumber()] = ((_this).f25).isLessThanOrEqualTo((_this).f25);
        _nw156[(new BigNumber(5)).toNumber()] = true;
        _nw156[(new BigNumber(6)).toNumber()] = (_1028_cf14) && (_1027_cf15);
        _nw156[(new BigNumber(7)).toNumber()] = (_1030_v16).dtor_cf27;
        _nw156[(new BigNumber(8)).toNumber()] = (new BigNumber((_1032_v18).length)).isLessThan((_this).f25);
        _nw156[(new BigNumber(9)).toNumber()] = true;
        _nw156[(new BigNumber(10)).toNumber()] = (new BigNumber((_1032_v18).length)).isLessThan((_this).f25);
        _nw156[(new BigNumber(11)).toNumber()] = ((_1027_cf15) ? (_1028_cf14) : (_1028_cf14));
        _nw156[(new BigNumber(12)).toNumber()] = !(_dafny.Seq.IsProperPrefixOf(_1003_v3, _1003_v3));
        _nw156[(new BigNumber(13)).toNumber()] = _1027_cf15;
        _nw156[(new BigNumber(14)).toNumber()] = (_1033_v19).dtor_cf45;
        _nw156[(new BigNumber(15)).toNumber()] = _1028_cf14;
        _nw156[(new BigNumber(16)).toNumber()] = _1027_cf15;
        _nw156[(new BigNumber(17)).toNumber()] = _1027_cf15;
        _nw156[(new BigNumber(18)).toNumber()] = true;
        _nw156[(new BigNumber(19)).toNumber()] = false;
        _nw156[(new BigNumber(20)).toNumber()] = p2;
        _nw156[(new BigNumber(21)).toNumber()] = p2;
        _nw156[(new BigNumber(22)).toNumber()] = _1027_cf15;
        _nw156[(new BigNumber(23)).toNumber()] = (_dafny.MultiSet.fromElements(_1002_v2, _1029_cf13)).IsDisjointFrom(_1034_v20);
        _nw156[(new BigNumber(24)).toNumber()] = _1028_cf14;
        _nw156[(new BigNumber(25)).toNumber()] = _1028_cf14;
        _nw156[(new BigNumber(26)).toNumber()] = false;
        _1035_v21 = _nw156;
        r1 = _1035_v21;
        _1030_v16 = _1030_v16;
        r3 = (_this).f25;
        let _1036_v23;
        let _init22 = ((_1037_p0, _1038_v1) => function (_1039_i2) {
          return (function () {
            let _coll34 = new _dafny.Map();
            for (const _compr_34 of (_dafny.Seq.of(_1038_v1)).Elements) {
              let _1040_v22 = _compr_34;
              if (_dafny.Seq.contains(_dafny.Seq.of(_1038_v1), _1040_v22)) {
                _coll34.push([_1040_v22,(_this).f25]);
              }
            }
            return _coll34;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("yv"),new BigNumber((_1037_p0).length)));
        })(p0, _1000_v1);
        let _nw157 = Array((new BigNumber(16)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw157.length); _i0_22++) {
          _nw157[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _1036_v23 = _nw157;
        let _index135 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_1036_v23).length));
        (_1036_v23)[_index135] = _module.__default.fm29(false, p2, globalState);
        let _1041_v24;
        _1041_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1000_v1,(_this).f25);
        let _1042_v25;
        _1042_v25 = _dafny.Seq.of(_module.__default.fm1(globalState), _1000_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-864)), ((_1043_v2) => function (_1044_i3) {
          return _1043_v2;
        })(_1002_v2)));
        let _index136 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_1036_v23).length));
        let _rhs131 = _1041_v24;
        let _rhs132 = _dafny.Seq.of(p2);
        let _rhs133 = (_this).f25;
        let _rhs134 = new BigNumber((((_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("k"), (_1042_v25)[_module.__default.safeIndex((_this).f25, new BigNumber((_1042_v25).length))])) ? (p1) : (p1))).length);
        let _lhs104 = _1036_v23;
        let _lhs105 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_1036_v23).length));
        let _lhs106 = globalState;
        let _lhs107 = globalState;
        _lhs104[_lhs105] = _rhs131;
        _1003_v3 = _rhs132;
        _lhs106.f8 = _rhs133;
        _lhs107.f8 = _rhs134;
      } else {
        let _1045___mcc_h7 = (_source16).cf8;
        let _1046_cf8 = _1045___mcc_h7;
        let _1047_v26;
        _1047_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,((_this).f25).minus(new BigNumber((_1046_cf8).length)));
        _1047_v26 = (_1047_v26).update((_this).f25, new BigNumber((_dafny.Seq.Concat(_1000_v1, _1000_v1)).length));
        r3 = (_this).f25;
        r3 = (_this).f25;
        let _1048_v27;
        let _nw158 = new _module.C0();
        _nw158.__ctor(p2);
        _1048_v27 = _nw158;
        let _1049_v28;
        let _nw159 = new _module.C1();
        _nw159.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_1046_cf8,_1000_v1), _1048_v27);
        _1049_v28 = _nw159;
        let _1050_v29;
        _1050_v29 = _dafny.MultiSet.fromElements(_1049_v28, _1049_v28, _1049_v28);
        let _1051_v30;
        _1051_v30 = _module.D15.create_DC32(_1050_v29);
        _1050_v29 = (_1050_v29).Difference((_1051_v30).dtor_cf51);
      }
      let _1052_i4;
      _1052_i4 = _dafny.ZERO;
      L5: {
        while (p2) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1052_i4)) {
              break L5;
            }
            _1052_i4 = (_1052_i4).plus(_dafny.ONE);
            (globalState).f19 = p2;
            (globalState).f8 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f25, ((_this).f25).multipliedBy((_this).f25)));
            let _1053_v31;
            let _nw160 = Array((new BigNumber(6)).toNumber()).fill(false);
            _1053_v31 = _nw160;
            let _index137 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1053_v31).length));
            (_1053_v31)[_index137] = p2;
            let _1054_v32;
            let _nw161 = Array((new BigNumber(3)).toNumber()).fill([]);
            _1054_v32 = _nw161;
            let _1055_v33;
            let _nw162 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
            _1055_v33 = _nw162;
            let _index138 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1053_v31).length));
            let _rhs135 = p2;
            let _rhs136 = _dafny.Seq.contains(_dafny.Seq.of(_1055_v33), _1055_v33);
            let _rhs137 = p2;
            let _rhs138 = _1054_v32;
            let _lhs108 = _1053_v31;
            let _lhs109 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1053_v31).length));
            let _lhs110 = globalState;
            _lhs108[_lhs109] = _rhs135;
            _lhs110.f22 = _rhs136;
            r2 = _rhs137;
            _1054_v32 = _rhs138;
            _997_v0 = ((_997_v0).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1000_v1,_module.__default.fm0(_dafny.Seq.update(_1003_v3, _module.__default.safeIndex((_this).f25, new BigNumber((_1003_v3).length)), p2), globalState)))).Merge(_997_v0);
          }
        }
      }
      let _1056_i5;
      _1056_i5 = _dafny.ZERO;
      L6: {
        while (p2) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1056_i5)) {
              break L6;
            }
            _1056_i5 = (_1056_i5).plus(_dafny.ONE);
            let _1057_v34;
            _1057_v34 = _dafny.MultiSet.fromElements(((_this).f25).isLessThanOrEqualTo(new BigNumber((_1000_v1).length)));
            let _1058_v35;
            _1058_v35 = _module.D11.create_DC24(_1057_v34);
            _1057_v34 = (_1058_v35).dtor_cf41;
            (globalState).f19 = ((p2) ? (p2) : (p2));
            let _1059_v36;
            _1059_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1003_v3);
            (globalState).f19 = _module.__default.fm0(((p2) ? (_dafny.Seq.of(p2)) : ((((_1059_v36).contains(p2)) ? ((_1059_v36).get(p2)) : (_1003_v3)))), globalState);
            (globalState).f8 = ((_this).f25).minus(new BigNumber(884));
          }
        }
      }
      let _1060_v37;
      _1060_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
      let _1061_v38;
      let _nw163 = Array((new BigNumber(8)).toNumber());
      _nw163[(_dafny.ZERO).toNumber()] = (_this).f25;
      _nw163[(_dafny.ONE).toNumber()] = (_this).f25;
      _nw163[(new BigNumber(2)).toNumber()] = new BigNumber(-867);
      _nw163[(new BigNumber(3)).toNumber()] = (_this).f25;
      _nw163[(new BigNumber(4)).toNumber()] = (((_1060_v37).contains((_this).f25)) ? ((_1060_v37).get((_this).f25)) : ((_this).f25));
      _nw163[(new BigNumber(5)).toNumber()] = ((_this).f25).minus(new BigNumber(172));
      _nw163[(new BigNumber(6)).toNumber()] = ((_this).f25).minus((_this).f25);
      _nw163[(new BigNumber(7)).toNumber()] = (_this).f25;
      _1061_v38 = _nw163;
      let _index139 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1061_v38).length));
      (_1061_v38)[_index139] = (_this).f25;
      let _index140 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_1061_v38).length));
      (_1061_v38)[_index140] = (_this).f25;
      let _index141 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_1061_v38).length));
      (_1061_v38)[_index141] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.fm5(globalState)).multipliedBy((_this).f25)));
      let _1062_v39;
      _1062_v39 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("bjxoc"), _1000_v1);
      let _index142 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1061_v38).length));
      let _index143 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_1061_v38).length));
      let _index144 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_1061_v38).length));
      let _rhs139 = new BigNumber((_dafny.Seq.Concat(((p2) ? (_1062_v39) : (_1062_v39)), _dafny.Seq.Concat(_1062_v39, _1062_v39))).length);
      let _rhs140 = new BigNumber((_1000_v1).length);
      let _rhs141 = p2;
      let _rhs142 = (((_this).f25).multipliedBy((_this).f25)).multipliedBy(((p2) ? ((_this).f25) : ((_this).f25)));
      let _lhs111 = _1061_v38;
      let _lhs112 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1061_v38).length));
      let _lhs113 = _1061_v38;
      let _lhs114 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_1061_v38).length));
      let _lhs115 = _1061_v38;
      let _lhs116 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_1061_v38).length));
      _lhs111[_lhs112] = _rhs139;
      _lhs113[_lhs114] = _rhs140;
      r0 = _rhs141;
      _lhs115[_lhs116] = _rhs142;
      r0 = !(p2) || (_dafny.Seq.IsPrefixOf(_1000_v1, _dafny.Seq.UnicodeFromString("sjbp")));
      let _1063_v40;
      let _nw164 = Array((new BigNumber(27)).toNumber());
      _nw164[(_dafny.ZERO).toNumber()] = p2;
      _nw164[(_dafny.ONE).toNumber()] = true;
      _nw164[(new BigNumber(2)).toNumber()] = p2;
      _nw164[(new BigNumber(3)).toNumber()] = p2;
      _nw164[(new BigNumber(4)).toNumber()] = p2;
      _nw164[(new BigNumber(5)).toNumber()] = p2;
      _nw164[(new BigNumber(6)).toNumber()] = p2;
      _nw164[(new BigNumber(7)).toNumber()] = !(p2);
      _nw164[(new BigNumber(8)).toNumber()] = p2;
      _nw164[(new BigNumber(9)).toNumber()] = p2;
      _nw164[(new BigNumber(10)).toNumber()] = p2;
      _nw164[(new BigNumber(11)).toNumber()] = p2;
      _nw164[(new BigNumber(12)).toNumber()] = p2;
      _nw164[(new BigNumber(13)).toNumber()] = p2;
      _nw164[(new BigNumber(14)).toNumber()] = p2;
      _nw164[(new BigNumber(15)).toNumber()] = p2;
      _nw164[(new BigNumber(16)).toNumber()] = p2;
      _nw164[(new BigNumber(17)).toNumber()] = p2;
      _nw164[(new BigNumber(18)).toNumber()] = p2;
      _nw164[(new BigNumber(19)).toNumber()] = p2;
      _nw164[(new BigNumber(20)).toNumber()] = !(p2);
      _nw164[(new BigNumber(21)).toNumber()] = p2;
      _nw164[(new BigNumber(22)).toNumber()] = p2;
      _nw164[(new BigNumber(23)).toNumber()] = p2;
      _nw164[(new BigNumber(24)).toNumber()] = true;
      _nw164[(new BigNumber(25)).toNumber()] = false;
      _nw164[(new BigNumber(26)).toNumber()] = p2;
      _1063_v40 = _nw164;
      let _1064_v41;
      _1064_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1000_v1,_1063_v40);
      let _1065_v42;
      _1065_v42 = _dafny.Seq.of(_dafny.Seq.of(p2, p2), _1003_v3, _1003_v3);
      let _nw165 = Array((new BigNumber(5)).toNumber());
      _nw165[(_dafny.ZERO).toNumber()] = !((_1064_v41).update(_1000_v1, _1063_v40)).equals(_1064_v41);
      _nw165[(_dafny.ONE).toNumber()] = !(p2);
      _nw165[(new BigNumber(2)).toNumber()] = _module.__default.fm0((_1065_v42)[_module.__default.safeIndex(new BigNumber((_1000_v1).length), new BigNumber((_1065_v42).length))], globalState);
      _nw165[(new BigNumber(3)).toNumber()] = p2;
      _nw165[(new BigNumber(4)).toNumber()] = p2;
      r1 = _nw165;
      let _1066_v43;
      _1066_v43 = _dafny.Set.fromElements(new BigNumber(-603));
      let _1067_v44;
      let _nw166 = new _module.C0();
      _nw166.__ctor(p2);
      _1067_v44 = _nw166;
      let _1068_v45;
      let _nw167 = new _module.C1();
      _nw167.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_1066_v43,_1000_v1), _1067_v44);
      _1068_v45 = _nw167;
      let _1069_v46;
      _1069_v46 = _module.D15.create_DC33(false, p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(171)), function (_1070_i6) {
  return new _dafny.CodePoint('q'.codePointAt(0));
})).length), _1068_v45, p2);
      let _1071_v47;
      _1071_v47 = _module.D8.create_DC17(_1002_v2, _1002_v2, (_1069_v46).dtor_cf54);
      r2 = (_module.__default.fm18(_dafny.MultiSet.fromElements(!(p2), p2, p2), _1071_v47, p2, globalState)).isLessThanOrEqualTo((_1061_v38)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_1061_v38).length))]);
      r3 = (_1061_v38)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_1061_v38).length))];
      return [r0, r1, r2, r3];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
