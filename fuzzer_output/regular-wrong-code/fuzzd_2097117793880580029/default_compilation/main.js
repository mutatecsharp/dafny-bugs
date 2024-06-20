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
    static fm2(p0, p1, globalState) {
      let _source0 = _module.D2.create_DC6(new BigNumber((_dafny.Seq.of(true, true, false, true)).length));
      if (_source0.is_DC5) {
        let _0___mcc_h0 = (_source0).cf6;
        let _1___mcc_h1 = (_source0).cf7;
        let _2___mcc_h2 = (_source0).cf8;
        let _3_cf8 = _2___mcc_h2;
        let _4_cf7 = _1___mcc_h1;
        let _5_cf6 = _0___mcc_h0;
        return _dafny.Seq.of(false);
      } else if (_source0.is_DC6) {
        let _6___mcc_h3 = (_source0).cf9;
        let _7_cf9 = _6___mcc_h3;
        return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true));
      } else if (_source0.is_DC7) {
        let _8___mcc_h4 = (_source0).cf10;
        let _9___mcc_h5 = (_source0).cf11;
        let _10_cf11 = _9___mcc_h5;
        let _11_cf10 = _8___mcc_h4;
        return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true));
      } else {
        let _12___mcc_h6 = (_source0).cf5;
        let _13_cf5 = _12___mcc_h6;
        return _dafny.Seq.of(true, false, true);
      }
    };
    static fm5(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(-960), new BigNumber(274))).length), new BigNumber(714)))).Difference(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("rbbbp")).length))))).Union((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(311), new BigNumber((_dafny.Seq.UnicodeFromString("yckvpg")).length), new BigNumber((_dafny.Set.fromElements(_module.D6.create_DC19(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("eshn")).length))).length)),new BigNumber(869))).length), new BigNumber(101), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-108))).length), new BigNumber((_dafny.Seq.UnicodeFromString("i")).length)), _module.D6.create_DC19(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(950))).length), new BigNumber(121), new BigNumber((_dafny.Set.fromElements(new BigNumber((function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of _dafny.IntegerRange(new BigNumber(385), new BigNumber(991))) {
    let _14_v0 = _compr_0;
    if (((new BigNumber(385)).isLessThanOrEqualTo(_14_v0)) && ((_14_v0).isLessThan(new BigNumber(991)))) {
      _coll0.push([(_14_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("vlcha")).length)),new BigNumber(726)]);
    }
  }
  return _coll0;
}()).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(365),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hnjtc")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),false)).length))).length))).length))).length))).length)), _module.D6.create_DC19(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber((function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of _dafny.IntegerRange(new BigNumber(476), new BigNumber(250))) {
    let _15_v1 = _compr_1;
    if (((new BigNumber(476)).isLessThanOrEqualTo(_15_v1)) && ((_15_v1).isLessThan(new BigNumber(250)))) {
      _coll1.push([(_15_v1).plus(new BigNumber((_dafny.Seq.UnicodeFromString("ge")).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length))]);
    }
  }
  return _coll1;
}()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, false)).length))).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-390), new BigNumber(638))) {
    let _16_v2 = _compr_2;
    if (((new BigNumber(-390)).isLessThanOrEqualTo(_16_v2)) && ((_16_v2).isLessThan(new BigNumber(638)))) {
      _coll2.push([(_16_v2).multipliedBy(new BigNumber(581)),new _dafny.CodePoint('u'.codePointAt(0))]);
    }
  }
  return _coll2;
}()).length),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber(903), new BigNumber((_dafny.Seq.UnicodeFromString("mi")).length), new BigNumber(-912))).cardinality()), new BigNumber(272))).cardinality()))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length))).length), new BigNumber(948)))).length), new BigNumber(863)))).Intersect(function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-420)))))).Elements) {
          let _17_v3 = _compr_3;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-420)))))).contains(_17_v3)) {
            _coll3.add(_17_v3);
          }
        }
        return _coll3;
      }()));
    };
    static fm6(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("rf")).length))).cardinality()))).length),true);
    };
    static fm7(p0, p1, globalState) {
      return (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.safeDivisionInt(new BigNumber(951), new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-444)), function (_18_i0) {
          return new BigNumber((_dafny.Seq.of(true)).length);
        })).Elements) {
          let _19_v0 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-444)), function (_18_i0) {
            return new BigNumber((_dafny.Seq.of(true)).length);
          }), _19_v0)) {
            _coll4.push([(_19_v0).plus(new BigNumber((_dafny.Set.fromElements(true)).length)),new BigNumber(390)]);
          }
        }
        return _coll4;
      }()).length))).length))).multipliedBy(((_module.D5.create_DC16(new BigNumber((_dafny.Seq.UnicodeFromString("bmlkxdin")).length), true, true)).dtor_cf22).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length),new _dafny.CodePoint('f'.codePointAt(0)))).length)))));
    };
    static fm8(globalState) {
      let _source1 = _module.D5.create_DC15(true, false);
      if (_source1.is_DC15) {
        let _20___mcc_h0 = (_source1).cf20;
        let _21___mcc_h1 = (_source1).cf21;
        let _22_cf21 = _21___mcc_h1;
        let _23_cf20 = _20___mcc_h0;
        return _module.D1.create_DC2(new BigNumber((_dafny.Seq.UnicodeFromString("cpt")).length), _23_cf20);
      } else if (_source1.is_DC16) {
        let _24___mcc_h2 = (_source1).cf22;
        let _25___mcc_h3 = (_source1).cf23;
        let _26___mcc_h4 = (_source1).cf24;
        let _27_cf24 = _26___mcc_h4;
        let _28_cf23 = _25___mcc_h3;
        let _29_cf22 = _24___mcc_h2;
        return _module.D1.create_DC2(new BigNumber((_dafny.Set.fromElements(_27_cf24, _28_cf23)).length), _27_cf24);
      } else {
        let _30___mcc_h5 = (_source1).cf19;
        let _31_cf19 = _30___mcc_h5;
        return _module.D1.create_DC2(new BigNumber(-957), true);
      }
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ggyklnkw")).length), new BigNumber(133));
    };
    static fm10(p0, globalState) {
      return (_dafny.Set.fromElements(false)).Difference((_dafny.Set.fromElements(!(!(!(false))))).Intersect(_dafny.Set.fromElements(false, true, !(!(false)), true, true)));
    };
    static fm11(p0, globalState) {
      if (!(!(((false) ? (true) : (false))))) {
        return !(true);
      } else {
        return false;
      }
    };
    static fm12(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("suuupmma"), _dafny.Seq.UnicodeFromString("ogyyc"))).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("bgrjnfjl")))).Intersect(function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("hjljvpj"), _dafny.Seq.UnicodeFromString("imkj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(208)), function (_32_i0) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }))).Elements) {
          let _33_v0 = _compr_5;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("hjljvpj"), _dafny.Seq.UnicodeFromString("imkj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(208)), function (_32_i0) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }))).contains(_33_v0)) {
            _coll5.add(_33_v0);
          }
        }
        return _coll5;
      }());
    };
    static fm13(p0, p1, globalState) {
      return _dafny.Seq.of((new BigNumber(762)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gdm")).length)));
    };
    static fm14(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-952)), function (_34_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      });
    };
    static fm15(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-981)).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D5.create_DC16(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
  let _coll6 = new _dafny.Set();
  for (const _compr_6 of (function () {
    let _coll7 = new _dafny.Set();
    for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D2.create_DC5(true, false, _dafny.Set.fromElements(new BigNumber(-109), new BigNumber(518)))).dtor_cf6))).cardinality()),(_dafny.ZERO).minus(new BigNumber(-434)))).length)))).length))).Keys.Elements) {
      let _35_v0 = _compr_7;
      if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D2.create_DC5(true, false, _dafny.Set.fromElements(new BigNumber(-109), new BigNumber(518)))).dtor_cf6))).cardinality()),(_dafny.ZERO).minus(new BigNumber(-434)))).length)))).length))).contains(_35_v0)) {
        _coll7.add(_35_v0);
      }
    }
    return _coll7;
  }()).Elements) {
    let _36_v1 = _compr_6;
    if ((function () {
      let _coll8 = new _dafny.Set();
      for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D2.create_DC5(true, false, _dafny.Set.fromElements(new BigNumber(-109), new BigNumber(518)))).dtor_cf6))).cardinality()),(_dafny.ZERO).minus(new BigNumber(-434)))).length)))).length))).Keys.Elements) {
        let _37_v0 = _compr_8;
        if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D2.create_DC5(true, false, _dafny.Set.fromElements(new BigNumber(-109), new BigNumber(518)))).dtor_cf6))).cardinality()),(_dafny.ZERO).minus(new BigNumber(-434)))).length)))).length))).contains(_37_v0)) {
          _coll8.add(_37_v0);
        }
      }
      return _coll8;
    }()).contains(_36_v1)) {
      _coll6.add(_36_v1);
    }
  }
  return _coll6;
}()).length))).length), true, false)).dtor_cf22,true)).length)),(_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(35), new BigNumber(349))))).IsProperSubsetOf(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber(837), new BigNumber((_dafny.Seq.UnicodeFromString("gcdsylhn")).length)))));
    };
    static m0(p0, globalState) {
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let _38_v0;
      _38_v0 = new BigNumber(353);
      let _39_v1;
      _39_v1 = new _dafny.CodePoint('d'.codePointAt(0));
      let _40_v2;
      _40_v2 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm7(new BigNumber(136), _38_v0, globalState),_39_v1);
      _40_v2 = (_40_v2).update(_38_v0, _39_v1);
      let _41_v3;
      _41_v3 = _module.D2.create_DC6(_38_v0);
      let _42_v4;
      let _nw0 = new _module.C0();
      _nw0.__ctor((_41_v3).dtor_cf9);
      _42_v4 = _nw0;
      let _43_v5;
      _43_v5 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm11(new BigNumber(641), globalState),_42_v4);
      let _44_v6;
      let _nw1 = Array((new BigNumber(10)).toNumber()).fill([]);
      _44_v6 = _nw1;
      let _45_v8;
      let _init0 = ((_46_v1) => function (_47_i0) {
        return function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of _dafny.IntegerRange(new BigNumber(826), new BigNumber(649))) {
            let _48_v7 = _compr_9;
            if (((new BigNumber(826)).isLessThanOrEqualTo(_48_v7)) && ((_48_v7).isLessThan(new BigNumber(649)))) {
              _coll9.push([_module.__default.safeModuloInt(_48_v7, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_46_v1)).length)),false]);
            }
          }
          return _coll9;
        }();
      })(_39_v1);
      let _nw2 = Array((new BigNumber(16)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
        _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _45_v8 = _nw2;
      let _index0 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_44_v6).length));
      (_44_v6)[_index0] = _45_v8;
      let _49_v9;
      _49_v9 = true;
      let _50_v10;
      _50_v10 = _dafny.Seq.UnicodeFromString("kngri");
      let _51_v11;
      _51_v11 = _module.D4.create_DC10(_45_v8);
      let _pat_let_tv0 = _45_v8;
      let _index1 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_44_v6).length));
      let _rhs0 = (new BigNumber((_dafny.Set.fromElements(!(_49_v9), _49_v9, _49_v9)).length)).minus((_38_v0).minus(_module.__default.fm7((_dafny.ZERO).minus((_42_v4).f22), (_dafny.ZERO).minus(_38_v0), globalState)));
      let _rhs1 = _43_v5;
      let _rhs2 = _module.__default.safeModuloInt(_module.__default.fm7((_42_v4).f22, _38_v0, globalState), new BigNumber((_dafny.Seq.Concat(_50_v10, _50_v10)).length));
      let _rhs3 = (function (_pat_let0_0) {
        return function (_52_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_53_dt__update_hcf14_h0) {
              return _module.D4.create_DC10(_53_dt__update_hcf14_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_51_v11)).dtor_cf14;
      let _rhs4 = _49_v9;
      let _lhs0 = globalState;
      let _lhs1 = _44_v6;
      let _lhs2 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_44_v6).length));
      let _lhs3 = globalState;
      _38_v0 = _rhs0;
      _43_v5 = _rhs1;
      _lhs0.f3 = _rhs2;
      _lhs1[_lhs2] = _rhs3;
      _lhs3.f14 = _rhs4;
      _50_v10 = _50_v10;
      let _54_v12;
      _54_v12 = _dafny.Map.Empty.slice().updateUnsafe(_38_v0,_49_v9);
      let _55_v13;
      _55_v13 = _54_v12;
      let _56_v14;
      _56_v14 = _dafny.Map.Empty.slice().updateUnsafe(_49_v9,new BigNumber(((_55_v13)).length));
      let _57_v15;
      _57_v15 = _dafny.Set.fromElements(new BigNumber((_module.__default.fm12((_42_v4).f22, _38_v0, _50_v10, globalState)).length), (_42_v4).f22, (((_56_v14).contains(_49_v9)) ? ((_56_v14).get(_49_v9)) : ((_42_v4).f22)), _38_v0);
      let _58_i1;
      _58_i1 = _dafny.ZERO;
      L0: {
        while (!((_57_v15).IsProperSubsetOf(_dafny.Set.fromElements(_38_v0)))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_58_i1)) {
              break L0;
            }
            _58_i1 = (_58_i1).plus(_dafny.ONE);
            let _59_v16;
            _59_v16 = _dafny.Seq.of(_43_v5, _43_v5);
            (globalState).f14 = ((_42_v4).f22).isLessThanOrEqualTo(new BigNumber((((_59_v16)[_module.__default.safeIndex(new BigNumber(607), new BigNumber((_59_v16).length))]).Merge((_43_v5).update(false, _42_v4))).length));
            let _60_v17;
            _60_v17 = _dafny.Seq.of(_49_v9, _49_v9);
            let _61_v18;
            let _nw3 = new _module.C1();
            _nw3.__ctor(new BigNumber(929), _60_v17);
            _61_v18 = _nw3;
            let _62_v19;
            _62_v19 = _dafny.Set.fromElements(_61_v18, _61_v18);
            _62_v19 = _62_v19;
            let _63_v20;
            _63_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm13(_39_v1, _49_v9, globalState),_module.__default.safeDivisionInt(new BigNumber(552), (_42_v4).f22));
            let _64_v21;
            _64_v21 = _dafny.Seq.of((_42_v4).f22);
            _63_v20 = (_63_v20).update(_64_v21, (_42_v4).f22);
            let _65_v22;
            let _nw4 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _65_v22 = _nw4;
            let _index2 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_65_v22).length));
            (_65_v22)[_index2] = _50_v10;
            let _index3 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_65_v22).length));
            (_65_v22)[_index3] = _50_v10;
          }
        }
      }
      if (!(_49_v9)) {
        if (_49_v9) {
          let _66_v23;
          _66_v23 = _dafny.Seq.of(_49_v9);
          let _67_v24;
          let _nw5 = new _module.C1();
          _nw5.__ctor((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), function (_68_i2) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length)).plus((_42_v4).f22), _66_v23);
          _67_v24 = _nw5;
          let _rhs5 = ((_49_v9) ? ((_42_v4).f22) : ((_42_v4).f22));
          let _rhs6 = (_42_v4).f22;
          let _lhs4 = globalState;
          _lhs4.f3 = _rhs5;
          _38_v0 = _rhs6;
          let _69_v25;
          _69_v25 = _module.D3.create_DC8((_module.D3.create_DC8(_67_v24)).dtor_cf12);
          let _70_v26;
          _70_v26 = _dafny.MultiSet.fromElements(_69_v25);
          let _71_v27;
          let _nw6 = Array((new BigNumber(10)).toNumber()).fill([]);
          _71_v27 = _nw6;
          let _72_v28;
          let _init1 = ((_73_v9) => function (_74_i3) {
            return _73_v9;
          })(_49_v9);
          let _nw7 = Array((new BigNumber(12)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw7.length); _i0_1++) {
            _nw7[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _72_v28 = _nw7;
          let _index4 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_71_v27).length));
          (_71_v27)[_index4] = ((_49_v9) ? (_72_v28) : (_72_v28));
          let _index5 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_71_v27).length));
          let _rhs7 = _dafny.MultiSet.fromElements(_69_v25);
          let _rhs8 = _72_v28;
          let _rhs9 = !(_49_v9) || (!(_49_v9));
          let _rhs10 = !_dafny.Seq.contains(_50_v10, _39_v1);
          let _lhs5 = _71_v27;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_71_v27).length));
          let _lhs7 = globalState;
          let _lhs8 = globalState;
          _70_v26 = _rhs7;
          _lhs5[_lhs6] = _rhs8;
          _lhs7.f14 = _rhs9;
          _lhs8.f14 = _rhs10;
          (globalState).f7 = _49_v9;
          let _75_v30;
          _75_v30 = _module.D2.create_DC7((_dafny.ZERO).minus((_42_v4).f22), new BigNumber((function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of _dafny.IntegerRange(new BigNumber(667), new BigNumber(551))) {
    let _76_v29 = _compr_10;
    if (((new BigNumber(667)).isLessThanOrEqualTo(_76_v29)) && ((_76_v29).isLessThan(new BigNumber(551)))) {
      _coll10.push([_module.__default.safeDivisionInt(_76_v29, (_42_v4).f22),_49_v9]);
    }
  }
  return _coll10;
}()).length));
          let _77_v31;
          _77_v31 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_49_v9,_39_v1)).update(_49_v9, _39_v1),_66_v23);
          let _78_v32;
          _78_v32 = _dafny.Set.fromElements(_49_v9);
          let _79_v33;
          _79_v33 = _dafny.Map.Empty.slice().updateUnsafe(_49_v9,_78_v32);
          let _80_v34;
          let _nw8 = Array((new BigNumber(5)).toNumber());
          _nw8[(_dafny.ZERO).toNumber()] = (_75_v30).dtor_cf10;
          _nw8[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_module.__default.fm14(globalState), _50_v10)).length);
          _nw8[(new BigNumber(2)).toNumber()] = ((_42_v4).f22).multipliedBy((_42_v4).f22);
          _nw8[(new BigNumber(3)).toNumber()] = ((_67_v24).fm1(_77_v31, globalState)).multipliedBy(new BigNumber(351));
          _nw8[(new BigNumber(4)).toNumber()] = new BigNumber((_79_v33).length);
          _80_v34 = _nw8;
          let _nw9 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _80_v34 = _nw9;
        } else {
          let _81_v35;
          _81_v35 = _dafny.Seq.of(_49_v9);
          let _82_v36;
          let _nw10 = new _module.C1();
          _nw10.__ctor((_42_v4).f22, _81_v35);
          _82_v36 = _nw10;
          let _83_v37;
          _83_v37 = _dafny.Seq.of(_82_v36);
          let _84_v38;
          _84_v38 = _dafny.Seq.of((_82_v36).f20, (_42_v4).f22, (_42_v4).f22, (_42_v4).f22);
          let _85_v39;
          _85_v39 = _dafny.Seq.of(_84_v38, _84_v38);
          let _86_v40;
          _86_v40 = _dafny.Seq.of(_50_v10);
          let _87_v41;
          let _nw11 = Array((new BigNumber(15)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = new BigNumber((_50_v10).length);
          _nw11[(_dafny.ONE).toNumber()] = _38_v0;
          _nw11[(new BigNumber(2)).toNumber()] = new BigNumber(-302);
          _nw11[(new BigNumber(3)).toNumber()] = new BigNumber((_83_v37).length);
          _nw11[(new BigNumber(4)).toNumber()] = (_82_v36).f20;
          _nw11[(new BigNumber(5)).toNumber()] = (_42_v4).f22;
          _nw11[(new BigNumber(6)).toNumber()] = (_42_v4).f22;
          _nw11[(new BigNumber(7)).toNumber()] = (_42_v4).f22;
          _nw11[(new BigNumber(8)).toNumber()] = ((_42_v4).f22).minus(new BigNumber(531));
          _nw11[(new BigNumber(9)).toNumber()] = _38_v0;
          _nw11[(new BigNumber(10)).toNumber()] = _38_v0;
          _nw11[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_81_v35).length));
          _nw11[(new BigNumber(12)).toNumber()] = (_42_v4).f22;
          _nw11[(new BigNumber(13)).toNumber()] = (new BigNumber((_85_v39).length)).minus(new BigNumber((_86_v40).length));
          _nw11[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(((_dafny.ZERO).minus((_42_v4).f22)).minus(new BigNumber(366)));
          _87_v41 = _nw11;
          let _index6 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_87_v41).length));
          (_87_v41)[_index6] = new BigNumber((_50_v10).length);
          let _index7 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_87_v41).length));
          (_87_v41)[_index7] = (_dafny.ZERO).minus(_module.__default.fm7(_module.__default.safeDivisionInt((_82_v36).f20, (_42_v4).f22), (_42_v4).f22, globalState));
          let _88_v42;
          let _nw12 = Array((new BigNumber(5)).toNumber()).fill(false);
          _88_v42 = _nw12;
          let _index8 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_88_v42).length));
          (_88_v42)[_index8] = _49_v9;
          let _index9 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_88_v42).length));
          (_88_v42)[_index9] = (_82_v36).fm0(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lsnskptul"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(893)), ((_89_v1) => function (_90_i4) {
            return _89_v1;
          })(_39_v1))), (_42_v4).f22, (_42_v4).f22, new BigNumber(994), globalState);
          let _index10 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_88_v42).length));
          (_88_v42)[_index10] = !(_49_v9);
          let _91_v43;
          _91_v43 = _dafny.Map.Empty.slice().updateUnsafe(_49_v9,(_88_v42)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((_88_v42).length))]);
          let _index11 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_87_v41).length));
          (_87_v41)[_index11] = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_41_v3).dtor_cf9, (_82_v36).f20), _module.__default.safeDivisionInt((_82_v36).f20, new BigNumber((_91_v43).length)));
          (globalState).f11 = (_81_v35)[_module.__default.safeIndex(_38_v0, new BigNumber((_81_v35).length))];
        }
        let _92_v44;
        let _nw13 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _92_v44 = _nw13;
        let _index12 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_92_v44).length));
        (_92_v44)[_index12] = _38_v0;
        let _93_v45;
        _93_v45 = _dafny.MultiSet.fromElements(_49_v9);
        let _index13 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_92_v44).length));
        (_92_v44)[_index13] = ((_49_v9) ? ((_42_v4).f22) : (_module.__default.safeDivisionInt(new BigNumber((_93_v45).cardinality()), _38_v0)));
        (globalState).f14 = _49_v9;
        (globalState).f3 = (_92_v44)[_module.__default.safeIndex(new BigNumber(577), new BigNumber((_92_v44).length))];
        let _index14 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_92_v44).length));
        (_92_v44)[_index14] = ((_42_v4).f22).plus((_92_v44)[_module.__default.safeIndex(new BigNumber(577), new BigNumber((_92_v44).length))]);
      } else {
        let _94_v46;
        _94_v46 = _module.D1.create_DC2((_42_v4).f22, _49_v9);
        (globalState).f5 = ((_49_v9) ? ((_94_v46).dtor_cf3) : (_49_v9));
        (globalState).f14 = _module.__default.fm11((_42_v4).f22, globalState);
        let _95_v47;
        _95_v47 = _dafny.Seq.of(true, _49_v9);
        let _96_v48;
        let _nw14 = new _module.C1();
        _nw14.__ctor(_38_v0, _95_v47);
        _96_v48 = _nw14;
        let _97_v49;
        _97_v49 = _dafny.Map.Empty.slice().updateUnsafe(_96_v48,_dafny.Seq.of((_96_v48).f20, (_42_v4).f22));
        let _98_v50;
        _98_v50 = _dafny.Seq.of(new BigNumber(-91), (_42_v4).f22);
        if (((!_dafny.Seq.contains((((_97_v49).contains(_96_v48)) ? ((_97_v49).get(_96_v48)) : (_98_v50)), _38_v0)) ? ((new BigNumber((_dafny.Seq.UnicodeFromString("ihliafrys")).length)).isLessThan((_96_v48).f20)) : (_49_v9))) {
          let _99_v51;
          _99_v51 = _dafny.Map.Empty.slice().updateUnsafe(_42_v4,(_96_v48).f20);
          _38_v0 = new BigNumber((_99_v51).length);
          let _100_v52;
          _100_v52 = _dafny.Seq.of(_50_v10, _50_v10);
          let _101_v53;
          _101_v53 = _dafny.MultiSet.fromElements(_39_v1);
          let _102_v54;
          _102_v54 = _module.D6.create_DC17(_39_v1);
          let _rhs11 = (_dafny.ZERO).minus(((_42_v4).f22).plus(new BigNumber(830)));
          let _rhs12 = _dafny.Seq.Concat((_module.D5.create_DC14((_100_v52)[_module.__default.safeIndex(new BigNumber((_101_v53).cardinality()), new BigNumber((_100_v52).length))])).dtor_cf19, _dafny.Seq.update(_50_v10, _module.__default.safeIndex(_38_v0, new BigNumber((_50_v10).length)), (_102_v54).dtor_cf25));
          let _rhs13 = (_42_v4).f22;
          let _lhs9 = globalState;
          let _lhs10 = globalState;
          _lhs9.f3 = _rhs11;
          _50_v10 = _rhs12;
          _lhs10.f3 = _rhs13;
          let _103_v55;
          _103_v55 = _dafny.MultiSet.fromElements(_42_v4, _42_v4);
          let _104_v56;
          _104_v56 = _dafny.Map.Empty.slice().updateUnsafe(_49_v9,true);
          let _105_v58;
          _105_v58 = _dafny.Seq.of(_95_v47);
          let _106_v59;
          _106_v59 = _dafny.Map.Empty.slice().updateUnsafe(_104_v56,function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_105_v58).Elements) {
              let _107_v57 = _compr_11;
              if (_dafny.Seq.contains(_105_v58, _107_v57)) {
                _coll11.push([_107_v57,(_96_v48).f20]);
              }
            }
            return _coll11;
          }());
          let _108_v60;
          _108_v60 = _dafny.Map.Empty.slice().updateUnsafe((_96_v48).f21,new BigNumber(-797));
          let _109_v61;
          _109_v61 = _dafny.MultiSet.fromElements((((_106_v59).contains(_module.__default.fm15((_42_v4).f22, _50_v10, false, (_42_v4).f22, globalState))) ? ((_106_v59).get(_module.__default.fm15((_42_v4).f22, _50_v10, false, (_42_v4).f22, globalState))) : (_108_v60)), _108_v60);
          let _110_v62;
          let _nw15 = Array((new BigNumber(26)).toNumber());
          _nw15[(_dafny.ZERO).toNumber()] = (_96_v48).f20;
          _nw15[(_dafny.ONE).toNumber()] = new BigNumber((_50_v10).length);
          _nw15[(new BigNumber(2)).toNumber()] = new BigNumber((_103_v55).cardinality());
          _nw15[(new BigNumber(3)).toNumber()] = new BigNumber((_50_v10).length);
          _nw15[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_96_v48).f20);
          _nw15[(new BigNumber(5)).toNumber()] = new BigNumber(-971);
          _nw15[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_38_v0);
          _nw15[(new BigNumber(7)).toNumber()] = (_42_v4).f22;
          _nw15[(new BigNumber(8)).toNumber()] = (_96_v48).f20;
          _nw15[(new BigNumber(9)).toNumber()] = (_42_v4).f22;
          _nw15[(new BigNumber(10)).toNumber()] = new BigNumber(570);
          _nw15[(new BigNumber(11)).toNumber()] = new BigNumber(-802);
          _nw15[(new BigNumber(12)).toNumber()] = (_96_v48).f20;
          _nw15[(new BigNumber(13)).toNumber()] = (_42_v4).f22;
          _nw15[(new BigNumber(14)).toNumber()] = (_96_v48).f20;
          _nw15[(new BigNumber(15)).toNumber()] = (_42_v4).f22;
          _nw15[(new BigNumber(16)).toNumber()] = (_42_v4).f22;
          _nw15[(new BigNumber(17)).toNumber()] = (((_109_v61).contains(_dafny.Map.Empty.slice().updateUnsafe(_95_v47,new BigNumber(980)))) ? ((_109_v61).get(_dafny.Map.Empty.slice().updateUnsafe(_95_v47,new BigNumber(980)))) : ((_42_v4).f22));
          _nw15[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((_42_v4).f22);
          _nw15[(new BigNumber(19)).toNumber()] = (_42_v4).f22;
          _nw15[(new BigNumber(20)).toNumber()] = (_96_v48).f20;
          _nw15[(new BigNumber(21)).toNumber()] = (_96_v48).f20;
          _nw15[(new BigNumber(22)).toNumber()] = new BigNumber((_56_v14).length);
          _nw15[(new BigNumber(23)).toNumber()] = (_42_v4).f22;
          _nw15[(new BigNumber(24)).toNumber()] = (_42_v4).f22;
          _nw15[(new BigNumber(25)).toNumber()] = (_dafny.ZERO).minus((((_56_v14).contains(_49_v9)) ? ((_56_v14).get(_49_v9)) : ((_96_v48).f20)));
          _110_v62 = _nw15;
          let _111_v63;
          _111_v63 = _dafny.Map.Empty.slice().updateUnsafe(_49_v9,_110_v62);
          _111_v63 = (_111_v63).update(_49_v9, _110_v62);
          let _112_v64;
          _112_v64 = _dafny.MultiSet.fromElements((_96_v48).f20);
          _112_v64 = _dafny.MultiSet.fromElements(_38_v0, (_42_v4).f22, (_96_v48).f20, (_96_v48).f20, _38_v0);
          let _113_v65;
          _113_v65 = _dafny.Seq.of(_42_v4, _42_v4, _42_v4, _42_v4);
          let _114_v66;
          _114_v66 = _42_v4;
          let _115_v67;
          let _nw16 = Array((new BigNumber(23)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = (_113_v65)[_module.__default.safeIndex(new BigNumber((_50_v10).length), new BigNumber((_113_v65).length))];
          _nw16[(_dafny.ONE).toNumber()] = (_114_v66);
          _nw16[(new BigNumber(2)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(3)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(4)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(5)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(6)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(7)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(8)).toNumber()] = ((_49_v9) ? (_42_v4) : (_42_v4));
          _nw16[(new BigNumber(9)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(10)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(11)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(12)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(13)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(14)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(15)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(16)).toNumber()] = (_113_v65)[_module.__default.safeIndex((_42_v4).f22, new BigNumber((_113_v65).length))];
          _nw16[(new BigNumber(17)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(18)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(19)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(20)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(21)).toNumber()] = _42_v4;
          _nw16[(new BigNumber(22)).toNumber()] = _42_v4;
          _115_v67 = _nw16;
          let _index15 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_115_v67).length));
          (_115_v67)[_index15] = _42_v4;
          let _index16 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_115_v67).length));
          (_115_v67)[_index16] = _42_v4;
        } else {
          let _116_v68;
          let _nw17 = Array((new BigNumber(13)).toNumber()).fill(false);
          _116_v68 = _nw17;
          let _index17 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_116_v68).length));
          (_116_v68)[_index17] = _49_v9;
          let _index18 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_116_v68).length));
          (_116_v68)[_index18] = (_38_v0).isLessThan((_dafny.ZERO).minus((_96_v48).f20));
          let _117_v69;
          let _nw18 = Array((_dafny.ONE).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-659));
          _117_v69 = _nw18;
          let _118_v70;
          _118_v70 = _dafny.MultiSet.fromElements((_116_v68)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((_116_v68).length))], (_116_v68)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((_116_v68).length))]);
          let _119_v71;
          _119_v71 = _dafny.Map.Empty.slice().updateUnsafe(_49_v9,(_42_v4).fm4(new BigNumber((_118_v70).cardinality()), globalState));
          let _120_v72;
          _120_v72 = _dafny.Map.Empty.slice().updateUnsafe(_119_v71,(_96_v48).f21);
          let _index19 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_117_v69).length));
          (_117_v69)[_index19] = (_96_v48).fm1(_120_v72, globalState);
          let _index20 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_117_v69).length));
          (_117_v69)[_index20] = (_96_v48).f20;
          (globalState).f3 = _module.__default.safeDivisionInt(((_49_v9) ? (new BigNumber(-49)) : (new BigNumber(290))), (_42_v4).f22);
          let _121_v73;
          _121_v73 = _module.D5.create_DC14(_50_v10);
          let _122_v74;
          _122_v74 = _dafny.Map.Empty.slice().updateUnsafe((_116_v68)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((_116_v68).length))],_dafny.Seq.UnicodeFromString("btkawurs"));
          let _123_v75;
          let _nw19 = Array((new BigNumber(19)).toNumber());
          _nw19[(_dafny.ZERO).toNumber()] = _50_v10;
          _nw19[(_dafny.ONE).toNumber()] = _50_v10;
          _nw19[(new BigNumber(2)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat((_121_v73).dtor_cf19, _50_v10);
          _nw19[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_50_v10, _50_v10);
          _nw19[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_50_v10, _50_v10), _module.__default.safeIndex((_98_v50)[_module.__default.safeIndex(new BigNumber(988), new BigNumber((_98_v50).length))], new BigNumber((_dafny.Seq.Concat(_50_v10, _50_v10)).length)), _39_v1);
          _nw19[(new BigNumber(6)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(7)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(8)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(9)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kspjtwcah"), _module.__default.fm14(globalState));
          _nw19[(new BigNumber(11)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_50_v10, _module.__default.safeIndex((_42_v4).f22, new BigNumber((_50_v10).length)), new _dafny.CodePoint('j'.codePointAt(0)));
          _nw19[(new BigNumber(13)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(14)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(15)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(16)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(17)).toNumber()] = _50_v10;
          _nw19[(new BigNumber(18)).toNumber()] = (((_122_v74).contains(_49_v9)) ? ((_122_v74).get(_49_v9)) : (_50_v10));
          _123_v75 = _nw19;
          let _index21 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_123_v75).length));
          (_123_v75)[_index21] = _50_v10;
          let _index22 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_123_v75).length));
          (_123_v75)[_index22] = ((false) ? (_50_v10) : (_50_v10));
          let _124_v76;
          _124_v76 = _dafny.Map.Empty.slice().updateUnsafe(_117_v69,_49_v9);
          let _125_v77;
          _125_v77 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_117_v69)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_117_v69).length))], (_42_v4).f22),_124_v76);
          _125_v77 = (_125_v77).update((_96_v48).f20, _124_v76);
        }
        let _126_v78;
        _126_v78 = _module.D8.create_DC21((_96_v48).f21);
        let _127_v79;
        _127_v79 = _dafny.Map.Empty.slice().updateUnsafe(!(_49_v9),_49_v9);
        let _128_v80;
        _128_v80 = _dafny.Set.fromElements(_42_v4);
        let _129_v81;
        _129_v81 = _dafny.Map.Empty.slice().updateUnsafe((((_127_v79).contains(_module.__default.fm11(new BigNumber((_128_v80).length), globalState))) ? ((_127_v79).get(_module.__default.fm11(new BigNumber((_128_v80).length), globalState))) : (!(_49_v9))),_57_v15);
        let _130_v82;
        _130_v82 = _dafny.Map.Empty.slice().updateUnsafe((_126_v78).dtor_cf34,(((_129_v81).contains(_49_v9)) ? ((_129_v81).get(_49_v9)) : (_dafny.Set.fromElements(new BigNumber((_95_v47).length)))));
        _130_v82 = (_130_v82).update(_dafny.Seq.update(_95_v47, _module.__default.safeIndex((_42_v4).f22, new BigNumber((_95_v47).length)), (((_127_v79).contains(_49_v9)) ? ((_127_v79).get(_49_v9)) : (_49_v9))), _57_v15);
        let _131_v83;
        let _init2 = ((_132_v15) => function (_133_i5) {
          return _module.__default.safeModuloInt(_133_i5, new BigNumber((_132_v15).length));
        })(_57_v15);
        let _nw20 = Array((new BigNumber(3)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
          _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _131_v83 = _nw20;
        let _index23 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_131_v83).length));
        (_131_v83)[_index23] = (_42_v4).f22;
        let _index24 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_131_v83).length));
        let _rhs14 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_42_v4).f22, (_42_v4).f22), new BigNumber(5));
        let _rhs15 = (_96_v48).f20;
        let _lhs11 = globalState;
        let _lhs12 = _131_v83;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_131_v83).length));
        _lhs11.f3 = _rhs14;
        _lhs12[_lhs13] = _rhs15;
      }
      let _134_v84;
      _134_v84 = _dafny.MultiSet.fromElements(_50_v10, _50_v10);
      let _hi0 = (_42_v4).f22;
      for (let _135_i6 = new BigNumber((_134_v84).cardinality()); _135_i6.isLessThan(_hi0); _135_i6 = _135_i6.plus(_dafny.ONE)) {
        let _136_v85;
        let _nw21 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _136_v85 = _nw21;
        let _index25 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_136_v85).length));
        (_136_v85)[_index25] = (_42_v4).f22;
        let _index26 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_136_v85).length));
        (_136_v85)[_index26] = _module.__default.safeModuloInt(_135_i6, (_42_v4).f22);
        let _137_v86;
        _137_v86 = _dafny.Seq.of(_135_i6);
        let _138_v87;
        _138_v87 = _dafny.Map.Empty.slice().updateUnsafe(_49_v9,_137_v86);
        let _139_v88;
        let _nw22 = new _module.C0();
        _nw22.__ctor(_module.__default.safeModuloInt(new BigNumber((_138_v87).length), _135_i6));
        _139_v88 = _nw22;
        let _140_v89;
        let _init3 = function (_141_i7) {
          return true;
        };
        let _nw23 = Array((new BigNumber(25)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw23.length); _i0_3++) {
          _nw23[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _140_v89 = _nw23;
        let _index27 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_140_v89).length));
        (_140_v89)[_index27] = _49_v9;
        let _index28 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_140_v89).length));
        (_140_v89)[_index28] = (_49_v9) || (_49_v9);
        let _142_v90;
        let _nw24 = new _module.C1();
        _nw24.__ctor(_135_i6, _dafny.Seq.of(true));
        _142_v90 = _nw24;
        let _143_v91;
        _143_v91 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_142_v90),true);
        let _144_v92;
        _144_v92 = _module.D3.create_DC8(_142_v90);
        _143_v91 = (_143_v91).update(_144_v92, (((_140_v89)[_module.__default.safeIndex(new BigNumber(64), new BigNumber((_140_v89).length))]) ? (true) : ((_140_v89)[_module.__default.safeIndex(new BigNumber(64), new BigNumber((_140_v89).length))])));
      }
      let _145_v93;
      _145_v93 = _dafny.Set.fromElements(!(_49_v9), _49_v9);
      r0 = _145_v93;
      let _146_v94;
      let _nw25 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
      _146_v94 = _nw25;
      let _147_v95;
      _147_v95 = _dafny.Seq.of(new BigNumber(580));
      let _148_v96;
      let _nw26 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
      _148_v96 = _nw26;
      let _149_v97;
      _149_v97 = _module.D8.create_DC22(_146_v94, _147_v95, _38_v0, _148_v96, _49_v9);
      r1 = (((_149_v97).dtor_cf39) ? (_49_v9) : (_49_v9));
      r2 = _54_v12;
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _150_v0;
      _150_v0 = false;
      let _151_v1;
      let _nw27 = Array((new BigNumber(17)).toNumber()).fill([]);
      _151_v1 = _nw27;
      let _152_v2;
      _152_v2 = new BigNumber(204);
      let _153_v3;
      let _nw28 = Array((new BigNumber(22)).toNumber()).fill(false);
      _153_v3 = _nw28;
      let _154_v4;
      _154_v4 = _dafny.Map.Empty.slice().updateUnsafe(_152_v2,_153_v3);
      let _155_v5;
      _155_v5 = _dafny.Seq.of(true);
      let _156_v6;
      _156_v6 = _dafny.Map.Empty.slice().updateUnsafe(_150_v0,_152_v2);
      let _157_v7;
      _157_v7 = _dafny.Seq.of(_150_v0, (_155_v5)[_module.__default.safeIndex(new BigNumber((_156_v6).length), new BigNumber((_155_v5).length))]);
      let _158_v8;
      _158_v8 = _dafny.Set.fromElements(_152_v2);
      let _159_v9;
      _159_v9 = _dafny.Set.fromElements(_150_v0);
      let _160_v10;
      let _nw29 = Array((new BigNumber(16)).toNumber()).fill([]);
      _160_v10 = _nw29;
      let _161_globalState;
      let _nw30 = new _module.GlobalState();
      _nw30.__ctor(((_150_v0) ? (_151_v1) : (_151_v1)), new BigNumber(872), false, new BigNumber(91), new BigNumber(751), true, new BigNumber(-494), true, new BigNumber(204), _154_v4, false, true, _157_v7, _158_v8, false, new BigNumber(-442), _159_v9, new BigNumber(341), new BigNumber(699), _160_v10);
      _161_globalState = _nw30;
      (_161_globalState).f3 = new BigNumber(((((_150_v0) ? (_159_v9) : (_159_v9))).Intersect(_159_v9)).length);
      let _162_v11;
      let _nw31 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _162_v11 = _nw31;
      let _163_v12;
      _163_v12 = _dafny.Seq.UnicodeFromString("skhwrlhb");
      let _index29 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length));
      (_162_v11)[_index29] = _dafny.Seq.Concat(_163_v12, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-478)), ((_164_v12, _165_v2) => function (_166_i0) {
        return (_164_v12)[_module.__default.safeIndex(_165_v2, new BigNumber((_164_v12).length))];
      })(_163_v12, _152_v2)));
      let _index30 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length));
      (_162_v11)[_index30] = _163_v12;
      if (_150_v0) {
        let _167_v13;
        let _init4 = function (_168_i1) {
          return (_168_i1).multipliedBy(new BigNumber(553));
        };
        let _nw32 = Array((new BigNumber(21)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw32.length); _i0_4++) {
          _nw32[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _167_v13 = _nw32;
        let _169_v14;
        _169_v14 = _dafny.Seq.of(_163_v12);
        let _rhs16 = _167_v13;
        let _rhs17 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_169_v14).length), (_152_v2).plus(_152_v2)));
        let _rhs18 = !(true);
        let _lhs14 = _161_globalState;
        _167_v13 = _rhs16;
        _lhs14.f3 = _rhs17;
        _150_v0 = _rhs18;
        let _170_v15;
        _170_v15 = _dafny.MultiSet.fromElements(_152_v2);
        let _171_v16;
        let _172_v17;
        let _173_v18;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = _module.__default.m0(_170_v15, _161_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _171_v16 = _out0;
        _172_v17 = _out1;
        _173_v18 = _out2;
        let _174_v19;
        _174_v19 = _dafny.Seq.of(_152_v2);
        let _175_v20;
        _175_v20 = _dafny.Map.Empty.slice().updateUnsafe(_174_v19,_152_v2);
        let _176_v21;
        _176_v21 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(32)).plus(new BigNumber(-306)),(((_175_v20).contains(_dafny.Seq.of(_152_v2))) ? ((_175_v20).get(_dafny.Seq.of(_152_v2))) : (_152_v2)));
        _176_v21 = (_176_v21).update((_dafny.ZERO).minus(_152_v2), new BigNumber(((_162_v11)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length))]).length));
        let _177_v22;
        _177_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(_150_v0),_150_v0);
        _156_v6 = (_156_v6).update((_152_v2).isLessThan(new BigNumber((_177_v22).length)), _152_v2);
        let _178_v23;
        let _nw33 = new _module.C1();
        _nw33.__ctor(new BigNumber(362), _157_v7);
        _178_v23 = _nw33;
      } else {
        _152_v2 = _152_v2;
        let _179_v24;
        let _nw34 = new _module.C1();
        _nw34.__ctor(_152_v2, _dafny.Seq.of(_150_v0));
        _179_v24 = _nw34;
        let _180_v25;
        _180_v25 = _dafny.Map.Empty.slice().updateUnsafe(_179_v24,new BigNumber(-39));
        let _181_v26;
        _181_v26 = _dafny.MultiSet.fromElements(new BigNumber((_180_v25).length), new BigNumber((_159_v9).length), (_179_v24).f20, _152_v2);
        let _182_v27;
        let _183_v28;
        let _184_v29;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector1 = _module.__default.m0(_181_v26, _161_globalState);
        _out3 = _outcollector1[0];
        _out4 = _outcollector1[1];
        _out5 = _outcollector1[2];
        _182_v27 = _out3;
        _183_v28 = _out4;
        _184_v29 = _out5;
        _152_v2 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_179_v24).f20, (_179_v24).f20), new BigNumber((_159_v9).length));
        (_161_globalState).f11 = !(_module.__default.safeDivisionInt(_152_v2, (_179_v24).f20)).isEqualTo((_179_v24).f20);
        let _185_v30;
        let _init5 = function (_186_i2) {
          return (_186_i2).multipliedBy(new BigNumber(-99));
        };
        let _nw35 = Array((new BigNumber(21)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw35.length); _i0_5++) {
          _nw35[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _185_v30 = _nw35;
        (_161_globalState).f7 = (_179_v24).fm0(_163_v12, (_179_v24).f20, _module.__default.safeDivisionInt(_152_v2, new BigNumber((_dafny.Set.fromElements(_185_v30, _185_v30, _185_v30, _185_v30, _185_v30)).length)), _152_v2, _161_globalState);
      }
      let _187_v31;
      _187_v31 = _dafny.MultiSet.fromElements(_150_v0, _150_v0, _150_v0);
      let _188_i3;
      _188_i3 = _dafny.ZERO;
      L1: {
        while ((_155_v5)[_module.__default.safeIndex((((_187_v31).contains(_150_v0)) ? ((_187_v31).get(_150_v0)) : (_152_v2)), new BigNumber((_155_v5).length))]) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_188_i3)) {
              break L1;
            }
            _188_i3 = (_188_i3).plus(_dafny.ONE);
            let _189_v33;
            _189_v33 = _dafny.Seq.of(_152_v2, _152_v2);
            let _190_v34;
            _190_v34 = function () {
              let _coll12 = new _dafny.Map();
              for (const _compr_12 of (_189_v33).Elements) {
                let _191_v32 = _compr_12;
                if (_dafny.Seq.contains(_189_v33, _191_v32)) {
                  _coll12.push([(_191_v32).plus(_152_v2),_150_v0]);
                }
              }
              return _coll12;
            }();
            _190_v34 = _module.__default.fm6((_dafny.ZERO).minus(_module.__default.fm7((((_187_v31).contains(_150_v0)) ? ((_187_v31).get(_150_v0)) : (_152_v2)), _152_v2, _161_globalState)), _161_globalState);
            let _192_v35;
            _192_v35 = _dafny.Map.Empty.slice().updateUnsafe(_190_v34,_dafny.Seq.UnicodeFromString("syahspiq"));
            _192_v35 = (_192_v35).update(_190_v34, _163_v12);
            let _193_v36;
            let _nw36 = Array((new BigNumber(2)).toNumber()).fill([]);
            _193_v36 = _nw36;
            _193_v36 = _193_v36;
            (_161_globalState).f14 = (_module.__default.safeModuloInt(_152_v2, (_dafny.ZERO).minus(_152_v2))).isEqualTo((_module.__default.fm8(_161_globalState)).dtor_cf2);
          }
        }
      }
      let _194_v37;
      let _nw37 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
      _194_v37 = _nw37;
      let _index31 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length));
      (_194_v37)[_index31] = _152_v2;
      let _195_v38;
      _195_v38 = _dafny.Seq.of(_163_v12, _dafny.Seq.UnicodeFromString("qsf"), _dafny.Seq.Concat((_162_v11)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length))], _163_v12), _163_v12, _163_v12);
      let _index32 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length));
      let _index33 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length));
      let _rhs19 = (_195_v38)[_module.__default.safeIndex((_dafny.ZERO).minus(_152_v2), new BigNumber((_195_v38).length))];
      let _rhs20 = _module.__default.safeDivisionInt(_152_v2, _152_v2);
      let _lhs15 = _162_v11;
      let _lhs16 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length));
      let _lhs17 = _194_v37;
      let _lhs18 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length));
      _lhs15[_lhs16] = _rhs19;
      _lhs17[_lhs18] = _rhs20;
      let _196_v39;
      _196_v39 = _module.D1.create_DC2(_152_v2, _150_v0);
      let _pat_let_tv1 = _150_v0;
      let _pat_let_tv2 = _194_v37;
      let _pat_let_tv3 = _194_v37;
      let _pat_let_tv4 = _152_v2;
      let _pat_let_tv5 = _194_v37;
      let _pat_let_tv6 = _194_v37;
      let _pat_let_tv7 = _152_v2;
      (_161_globalState).f7 = function (_source2) {
        if (_source2.is_DC2) {
          let _197___mcc_h0 = (_source2).cf2;
          let _198___mcc_h1 = (_source2).cf3;
          let _199_cf3 = _198___mcc_h1;
          let _200_cf2 = _197___mcc_h0;
          return !(_pat_let_tv1) || (_199_cf3);
        } else if (_source2.is_DC1) {
          let _201___mcc_h2 = (_source2).cf1;
          let _202_cf1 = _201___mcc_h2;
          return ((_pat_let_tv3)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_pat_let_tv2).length))]).isLessThanOrEqualTo(_pat_let_tv4);
        } else {
          let _203___mcc_h3 = (_source2).cf4;
          let _204_cf4 = _203___mcc_h3;
          return ((_pat_let_tv6)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_pat_let_tv5).length))]).isLessThanOrEqualTo(_pat_let_tv7);
        }
      }(_196_v39);
      let _205_v40;
      _205_v40 = _dafny.Map.Empty.slice().updateUnsafe((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))],_150_v0);
      let _206_v41;
      _206_v41 = _205_v40;
      let _pat_let_tv8 = _152_v2;
      (_161_globalState).f3 = function (_source3) {
        let _207___mcc_h4 = _source3;
        let _208_cf0 = _207___mcc_h4;
        return _pat_let_tv8;
      }(_206_v41);
      let _209_v42;
      _209_v42 = new _dafny.CodePoint('d'.codePointAt(0));
      let _210_v43;
      _210_v43 = _dafny.Set.fromElements(_209_v42, _209_v42);
      let _211_v44;
      _211_v44 = _dafny.Map.Empty.slice().updateUnsafe(_152_v2,_210_v43);
      let _212_v45;
      _212_v45 = _dafny.MultiSet.fromElements((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))], new BigNumber((_159_v9).length), (_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))], new BigNumber(26), new BigNumber(((_162_v11)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length))]).length));
      (_161_globalState).f3 = (_module.__default.safeModuloInt(_152_v2, (_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))])).multipliedBy(_module.__default.fm7(new BigNumber(((((_211_v44).contains((((_187_v31).contains(_150_v0)) ? ((_187_v31).get(_150_v0)) : ((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))])))) ? ((_211_v44).get((((_187_v31).contains(_150_v0)) ? ((_187_v31).get(_150_v0)) : ((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))])))) : (_210_v43))).length), (((_212_v45).contains(new BigNumber(((_162_v11)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length))]).length))) ? ((_212_v45).get(new BigNumber(((_162_v11)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length))]).length))) : ((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))])), _161_globalState));
      let _213_v46;
      _213_v46 = _dafny.Map.Empty.slice().updateUnsafe(_150_v0,_212_v45);
      let _214_v47;
      _214_v47 = _dafny.Seq.of(_152_v2);
      let _215_v48;
      _215_v48 = _module.D2.create_DC4(_dafny.MultiSet.FromArray(_214_v47));
      let _pat_let_tv9 = _152_v2;
      let _216_v49;
      _216_v49 = _dafny.Seq.of((function (_pat_let2_0) {
        return function (_217_dt__update__tmp_h0) {
          return function (_pat_let3_0) {
            return function (_218_dt__update_hcf5_h0) {
              return _module.D2.create_DC4(_218_dt__update_hcf5_h0);
            }(_pat_let3_0);
          }(_dafny.MultiSet.fromElements(_pat_let_tv9, new BigNumber(765)));
        }(_pat_let2_0);
      }(_215_v48)).dtor_cf5, _212_v45);
      let _219_v50;
      _219_v50 = _dafny.Map.Empty.slice().updateUnsafe(false,_163_v12);
      (_161_globalState).f14 = !((_212_v45).Union(_212_v45)).equals((((_213_v46).contains(_150_v0)) ? ((_213_v46).get(_150_v0)) : ((_216_v49)[_module.__default.safeIndex(new BigNumber(((((_219_v50).contains(_150_v0)) ? ((_219_v50).get(_150_v0)) : ((_162_v11)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length))]))).length), new BigNumber((_216_v49).length))])));
      let _hi1 = (((_156_v6).contains(_150_v0)) ? ((_156_v6).get(_150_v0)) : (_152_v2));
      for (let _220_i4 = (_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))]; _220_i4.isLessThan(_hi1); _220_i4 = _220_i4.plus(_dafny.ONE)) {
        let _221_v51;
        _221_v51 = _dafny.Map.Empty.slice().updateUnsafe(_205_v40,(_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))]);
        let _222_v52;
        _222_v52 = _dafny.Seq.of(_221_v51);
        let _223_v53;
        _223_v53 = _dafny.Map.Empty.slice().updateUnsafe(_152_v2,_152_v2);
        let _224_v54;
        let _225_v55;
        let _226_v56;
        let _out6;
        let _out7;
        let _out8;
        let _outcollector2 = _module.__default.m0(_module.__default.fm9(false, (_163_v12)[_module.__default.safeIndex((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))], new BigNumber((_163_v12).length))], (_222_v52)[_module.__default.safeIndex(_152_v2, new BigNumber((_222_v52).length))], _223_v53, _161_globalState), _161_globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _out8 = _outcollector2[2];
        _224_v54 = _out6;
        _225_v55 = _out7;
        _226_v56 = _out8;
        (_161_globalState).f3 = (_dafny.ZERO).minus(_module.__default.fm7(_152_v2, new BigNumber(((_162_v11)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length))]).length), _161_globalState));
        (_161_globalState).f7 = !(_150_v0);
        let _227_v57;
        _227_v57 = _dafny.Map.Empty.slice().updateUnsafe(true,!(new BigNumber((_dafny.Seq.UnicodeFromString("skjwiqufa")).length)).isEqualTo((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))]));
        _227_v57 = (_227_v57).update(_150_v0, _150_v0);
      }
      let _228_v58;
      _228_v58 = _module.D1.create_DC2(_module.__default.fm7(_152_v2, _152_v2, _161_globalState), (_155_v5)[_module.__default.safeIndex((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))], new BigNumber((_155_v5).length))]);
      let _229_v59;
      _229_v59 = _module.D1.create_DC3(_228_v58);
      let _230_v60;
      _230_v60 = _dafny.Map.Empty.slice().updateUnsafe(_229_v59,_209_v42);
      let _231_v61;
      _231_v61 = _dafny.Map.Empty.slice().updateUnsafe(_230_v60,(_150_v0) || (true));
      _231_v61 = (_231_v61).update(_230_v60, ((_150_v0) ? (_150_v0) : (_150_v0)));
      let _232_v62;
      _232_v62 = _dafny.Map.Empty.slice().updateUnsafe(_209_v42,_194_v37);
      let _hi2 = new BigNumber((_232_v62).length);
      for (let _233_i5 = _152_v2; _233_i5.isLessThan(_hi2); _233_i5 = _233_i5.plus(_dafny.ONE)) {
        let _234_v63;
        let _nw38 = new _module.C1();
        _nw38.__ctor(_152_v2, _dafny.Seq.Concat(_157_v7, _157_v7));
        _234_v63 = _nw38;
        (_161_globalState).f3 = (_233_i5).multipliedBy((_234_v63).f20);
        (_161_globalState).f11 = _150_v0;
        let _235_v64;
        _235_v64 = _module.D1.create_DC1(_233_i5);
        let _source4 = _235_v64;
        if (_source4.is_DC2) {
          let _236___mcc_h5 = (_source4).cf2;
          let _237___mcc_h6 = (_source4).cf3;
          let _238_cf3 = _237___mcc_h6;
          let _239_cf2 = _236___mcc_h5;
          let _240_v65;
          _240_v65 = _module.D2.create_DC6((_234_v63).f20);
          let _241_v66;
          _241_v66 = _dafny.Map.Empty.slice().updateUnsafe(_234_v63,(_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))]);
          let _242_v67;
          _242_v67 = _dafny.Seq.of(_241_v66, _241_v66);
          let _243_v68;
          _243_v68 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_152_v2),_209_v42);
          let _244_v69;
          _244_v69 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_242_v67).length),_209_v42), _243_v68, _243_v68, _dafny.Map.Empty.slice().updateUnsafe((_234_v63).f20,_209_v42));
          let _pat_let_tv10 = _244_v69;
          let _245_v70;
          let _nw39 = Array((new BigNumber(18)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _159_v9;
          _nw39[(_dafny.ONE).toNumber()] = _159_v9;
          _nw39[(new BigNumber(2)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(3)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(4)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(5)).toNumber()] = _module.__default.fm10(_240_v65, _161_globalState);
          _nw39[(new BigNumber(6)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(7)).toNumber()] = (_159_v9).Difference(_159_v9);
          _nw39[(new BigNumber(8)).toNumber()] = _module.__default.fm10(function (_pat_let4_0) {
            return function (_246_dt__update__tmp_h1) {
              return function (_pat_let5_0) {
                return function (_247_dt__update_hcf9_h0) {
                  return _module.D2.create_DC6(_247_dt__update_hcf9_h0);
                }(_pat_let5_0);
              }(new BigNumber((_dafny.MultiSet.FromArray(_pat_let_tv10)).cardinality()));
            }(_pat_let4_0);
          }(_240_v65), _161_globalState);
          _nw39[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(_150_v0, _150_v0, _150_v0);
          _nw39[(new BigNumber(10)).toNumber()] = (_159_v9).Difference(_159_v9);
          _nw39[(new BigNumber(11)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(12)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(13)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(14)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(15)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(16)).toNumber()] = _159_v9;
          _nw39[(new BigNumber(17)).toNumber()] = _module.__default.fm10(_240_v65, _161_globalState);
          _245_v70 = _nw39;
          let _index34 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_245_v70).length));
          (_245_v70)[_index34] = (_dafny.Set.fromElements(!(_238_cf3))).Difference(_159_v9);
          let _248_v71;
          let _nw40 = Array((new BigNumber(18)).toNumber()).fill([]);
          _248_v71 = _nw40;
          let _index35 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_248_v71).length));
          (_248_v71)[_index35] = _162_v11;
          let _index36 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length));
          let _index37 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_245_v70).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_248_v71).length));
          let _rhs21 = (_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))];
          let _rhs22 = ((_dafny.ZERO).minus((_234_v63).f20)).plus(new BigNumber(107));
          let _rhs23 = _159_v9;
          let _rhs24 = _162_v11;
          let _lhs19 = _161_globalState;
          let _lhs20 = _194_v37;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length));
          let _lhs22 = _245_v70;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_245_v70).length));
          let _lhs24 = _248_v71;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_248_v71).length));
          _lhs19.f3 = _rhs21;
          _lhs20[_lhs21] = _rhs22;
          _lhs22[_lhs23] = _rhs23;
          _lhs24[_lhs25] = _rhs24;
          _156_v6 = ((_156_v6).update(_150_v0, _239_cf2)).Merge(_156_v6);
          (_161_globalState).f7 = _150_v0;
          (_161_globalState).f14 = (_234_v63).fm0((_162_v11)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length))], (new BigNumber(17)).multipliedBy(_152_v2), _152_v2, (new BigNumber((_212_v45).cardinality())).plus((_234_v63).f20), _161_globalState);
        } else if (_source4.is_DC1) {
          let _249___mcc_h7 = (_source4).cf1;
          let _250_cf1 = _249___mcc_h7;
          _215_v48 = _215_v48;
          let _251_v72;
          _251_v72 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC1(new BigNumber((_dafny.Seq.of(_153_v3, _153_v3)).length)),_214_v47);
          _251_v72 = (_251_v72).update(_235_v64, _dafny.Seq.Concat(_214_v47, _214_v47));
          (_161_globalState).f11 = !((_159_v9).Intersect(_159_v9)).equals((_159_v9).Union(_159_v9));
          let _index39 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length));
          (_162_v11)[_index39] = (_162_v11)[_module.__default.safeIndex(new BigNumber(804), new BigNumber((_162_v11).length))];
        } else {
          let _252___mcc_h8 = (_source4).cf4;
          let _253_cf4 = _252___mcc_h8;
          _205_v40 = _205_v40;
          let _index40 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length));
          let _rhs25 = _150_v0;
          let _rhs26 = _233_i5;
          let _rhs27 = (_150_v0) && (!_dafny.Seq.contains(_163_v12, _209_v42));
          let _rhs28 = (_dafny.ZERO).minus((new BigNumber((_205_v40).length)).multipliedBy(new BigNumber((_163_v12).length)));
          let _lhs26 = _161_globalState;
          let _lhs27 = _194_v37;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length));
          let _lhs29 = _161_globalState;
          _lhs26.f7 = _rhs25;
          _lhs27[_lhs28] = _rhs26;
          _150_v0 = _rhs27;
          _lhs29.f3 = _rhs28;
          _157_v7 = _155_v5;
          _157_v7 = _157_v7;
        }
      }
      let _254_v73;
      let _nw41 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _254_v73 = _nw41;
      _254_v73 = _254_v73;
      (_161_globalState).f3 = new BigNumber(((_187_v31).Intersect((_187_v31).Difference(_187_v31))).cardinality());
      let _255_v74;
      let _256_v75;
      let _257_v76;
      let _out9;
      let _out10;
      let _out11;
      let _outcollector3 = _module.__default.m0(_dafny.MultiSet.fromElements(new BigNumber(857), (_dafny.ZERO).minus((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))]), (_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))], new BigNumber((_195_v38).length), _152_v2), _161_globalState);
      _out9 = _outcollector3[0];
      _out10 = _outcollector3[1];
      _out11 = _outcollector3[2];
      _255_v74 = _out9;
      _256_v75 = _out10;
      _257_v76 = _out11;
      let _258_v77;
      let _nw42 = new _module.C1();
      _nw42.__ctor(_module.__default.fm7((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))], new BigNumber(-177), _161_globalState), _155_v5);
      _258_v77 = _nw42;
      let _259_v78;
      _259_v78 = _module.D3.create_DC8(_258_v77);
      let _260_v79;
      _260_v79 = _dafny.Set.fromElements((_259_v78).dtor_cf12);
      let _rhs29 = true;
      let _rhs30 = ((_194_v37)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_194_v37).length))]).plus(_152_v2);
      let _rhs31 = (_dafny.ZERO).minus((((_187_v31).contains(_150_v0)) ? ((_187_v31).get(_150_v0)) : (new BigNumber((_260_v79).length))));
      let _lhs30 = _161_globalState;
      _150_v0 = _rhs29;
      _152_v2 = _rhs30;
      _lhs30.f3 = _rhs31;
      process.stdout.write(_dafny.toString(_150_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_152_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_154_v4).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_155_v5, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(204)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_157_v7, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v8).equals(_dafny.Set.fromElements(new BigNumber(204)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v9).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_161_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_161_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_161_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_161_globalState).f9).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_161_globalState.f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_161_globalState).f12, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_161_globalState).f13).equals(_dafny.Set.fromElements(new BigNumber(204)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_161_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_161_globalState).f16).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_162_v11)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_163_v12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v31).equals(_dafny.MultiSet.fromElements(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_188_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_194_v37)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_195_v38, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("skhwrlhb"), _dafny.Seq.UnicodeFromString("qsf"), _dafny.Seq.UnicodeFromString("skhwrlhbskhwrlhb"), _dafny.Seq.UnicodeFromString("skhwrlhb"), _dafny.Seq.UnicodeFromString("skhwrlhb")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v39).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v39).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v40).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_206_v41)).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_209_v42));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_210_v43).equals(_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_211_v44).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v45).equals(_dafny.MultiSet.fromElements(_dafny.ZERO, _dafny.ZERO, _dafny.ONE, new BigNumber(26), new BigNumber(8)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_213_v46).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.MultiSet.fromElements(_dafny.ZERO, _dafny.ZERO, _dafny.ONE, new BigNumber(26), new BigNumber(8))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_214_v47, _dafny.Seq.of(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_215_v48).dtor_cf5).equals(_dafny.MultiSet.fromElements(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_216_v49, _dafny.Seq.of(_dafny.MultiSet.fromElements(_dafny.ZERO, new BigNumber(765)), _dafny.MultiSet.fromElements(_dafny.ZERO, _dafny.ZERO, _dafny.ONE, new BigNumber(26), new BigNumber(8))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_v50).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("skhwrlhb")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_v58).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_v58).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_229_v59).dtor_cf4).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_229_v59).dtor_cf4).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v60).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(_module.D1.create_DC2(new BigNumber(7608), true)),new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_231_v61).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(_module.D1.create_DC2(new BigNumber(7608), true)),new _dafny.CodePoint('d'.codePointAt(0))),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_232_v62).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v74).equals(_dafny.Set.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_256_v75));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_257_v76).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(7257),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v77).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_258_v77).f21, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v78).dtor_cf12).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_259_v78).dtor_cf12).f21, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_260_v79).length)));
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
      return _dafny.Map.Empty;
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
    static create_DC2(cf2, cf3) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, false);
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
    static create_DC5(cf6, cf7, cf8) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC6(cf9) {
      let $dt = new D2(1);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC7(cf10, cf11) {
      let $dt = new D2(2);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC4(cf5) {
      let $dt = new D2(3);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get is_DC4() { return this.$tag === 3; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf6 === other.cf6 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(false, false, _dafny.Set.Empty);
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
    static create_DC9(cf13) {
      let $dt = new D3(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC8(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf12 === other.cf12;
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
    static create_DC11(cf15) {
      let $dt = new D4(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC12(cf16, cf17) {
      let $dt = new D4(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC10(cf14) {
      let $dt = new D4(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC13(cf18) {
      let $dt = new D4(3);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get is_DC13() { return this.$tag === 3; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf18) + ")";
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
        return other.$tag === 2 && this.cf14 === other.cf14;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_dafny.Seq.of());
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
    static create_DC15(cf20, cf21) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC16(cf22, cf23, cf24) {
      let $dt = new D5(1);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC14(cf19) {
      let $dt = new D5(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC14" + "(" + this.cf19.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23 && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(false, false);
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
    static create_DC18(cf26, cf27) {
      let $dt = new D6(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC19(cf28, cf29, cf30, cf31, cf32) {
      let $dt = new D6(1);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC17(cf25) {
      let $dt = new D6(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC18(false, _dafny.ZERO);
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
    static create_DC20(cf33) {
      let $dt = new D7(0);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf33 === other.cf33;
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf35, cf36, cf37, cf38, cf39) {
      let $dt = new D8(0);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC23(cf40, cf41) {
      let $dt = new D8(1);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC24(cf42, cf43, cf44) {
      let $dt = new D8(2);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC21(cf34) {
      let $dt = new D8(3);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC25(cf45) {
      let $dt = new D8(4);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get is_DC21() { return this.$tag === 3; }
    get is_DC25() { return this.$tag === 4; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 4) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37) && this.cf38 === other.cf38 && this.cf39 === other.cf39;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf40 === other.cf40 && this.cf41 === other.cf41;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf42 === other.cf42 && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC22([], _dafny.Seq.of(), _dafny.ZERO, [], false);
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
      this.f3 = _dafny.ZERO;
      this.f5 = false;
      this.f7 = false;
      this.f11 = false;
      this.f14 = false;
      this._f0 = [];
      this._f1 = _dafny.ZERO;
      this._f2 = false;
      this._f4 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this._f9 = _dafny.Map.Empty;
      this._f10 = false;
      this._f12 = _dafny.Seq.of();
      this._f13 = _dafny.Set.Empty;
      this._f15 = _dafny.ZERO;
      this._f16 = _dafny.Set.Empty;
      this._f17 = _dafny.ZERO;
      this._f18 = _dafny.ZERO;
      this._f19 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
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
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f6() {
      let _this = this;
      return _this._f6;
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
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f22 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f22) {
      let _this = this;
      (_this)._f22 = f22;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_dafny.MultiSet.fromElements(true, !(!(true)))).contains(false));
    };
    fm4(p0, globalState) {
      let _this = this;
      return new _dafny.CodePoint('g'.codePointAt(0));
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f20 = _dafny.ZERO;
      this._f21 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f20, f21) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(false) || (((_this).f21)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), function (_261_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length), new BigNumber(((_this).f21).length))]);
    };
    fm1(p0, globalState) {
      let _this = this;
      return ((_this).f20).minus((_this).f20);
    };
    m1(globalState) {
      let _this = this;
      let _262_v0;
      _262_v0 = false;
      let _263_v1;
      _263_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_262_v0);
      let _264_v3;
      _264_v3 = _dafny.Seq.of(new BigNumber(50));
      let _265_v5;
      _265_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f21).length),(_this).f20);
      let _266_v6;
      _266_v6 = function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_265_v5).Keys.Elements) {
          let _267_v4 = _compr_13;
          if ((_265_v5).contains(_267_v4)) {
            _coll13.push([_module.__default.safeDivisionInt(_267_v4, (_this).f20),_262_v0]);
          }
        }
        return _coll13;
      }();
      let _268_v7;
      _268_v7 = (_266_v6);
      let _269_v9;
      _269_v9 = _dafny.Seq.of(function () {
        let _coll14 = new _dafny.Map();
        for (const _compr_14 of _dafny.IntegerRange(new BigNumber(207), new BigNumber(401))) {
          let _270_v8 = _compr_14;
          if (((new BigNumber(207)).isLessThanOrEqualTo(_270_v8)) && ((_270_v8).isLessThan(new BigNumber(401)))) {
            _coll14.push([_module.__default.safeDivisionInt(_270_v8, (_this).f20),_262_v0]);
          }
        }
        return _coll14;
      }());
      let _271_v10;
      let _nw43 = Array((new BigNumber(20)).toNumber());
      _nw43[(_dafny.ZERO).toNumber()] = (_263_v1).Merge(_263_v1);
      _nw43[(_dafny.ONE).toNumber()] = (_263_v1).Merge(_263_v1);
      _nw43[(new BigNumber(2)).toNumber()] = function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of (_264_v3).Elements) {
          let _272_v2 = _compr_15;
          if (_dafny.Seq.contains(_264_v3, _272_v2)) {
            _coll15.push([(_272_v2).plus((_this).f20),_262_v0]);
          }
        }
        return _coll15;
      }();
      _nw43[(new BigNumber(3)).toNumber()] = (_263_v1).Merge((_263_v1).update((_this).f20, _262_v0));
      _nw43[(new BigNumber(4)).toNumber()] = _263_v1;
      _nw43[(new BigNumber(5)).toNumber()] = _263_v1;
      _nw43[(new BigNumber(6)).toNumber()] = _263_v1;
      _nw43[(new BigNumber(7)).toNumber()] = (_263_v1).update((_this).f20, true);
      _nw43[(new BigNumber(8)).toNumber()] = _263_v1;
      _nw43[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(69),_262_v0);
      _nw43[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_262_v0);
      _nw43[(new BigNumber(11)).toNumber()] = _263_v1;
      _nw43[(new BigNumber(12)).toNumber()] = _263_v1;
      _nw43[(new BigNumber(13)).toNumber()] = (_268_v7);
      _nw43[(new BigNumber(14)).toNumber()] = (_269_v9)[_module.__default.safeIndex((_this).f20, new BigNumber((_269_v9).length))];
      _nw43[(new BigNumber(15)).toNumber()] = _263_v1;
      _nw43[(new BigNumber(16)).toNumber()] = _263_v1;
      _nw43[(new BigNumber(17)).toNumber()] = _263_v1;
      _nw43[(new BigNumber(18)).toNumber()] = (_263_v1).Merge(_263_v1);
      _nw43[(new BigNumber(19)).toNumber()] = _263_v1;
      _271_v10 = _nw43;
      let _273_v11;
      _273_v11 = _dafny.Seq.UnicodeFromString("mjg");
      let _rhs32 = _271_v10;
      let _rhs33 = (_this).f20;
      let _rhs34 = !((_this).fm0(_273_v11, (_this).f20, (_this).f20, (_this).f20, globalState));
      let _lhs31 = globalState;
      let _lhs32 = globalState;
      _271_v10 = _rhs32;
      _lhs31.f3 = _rhs33;
      _lhs32.f5 = _rhs34;
      let _274_v12;
      let _init6 = ((_275_v0) => function (_276_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(_275_v0,_275_v0);
      })(_262_v0);
      let _nw44 = Array((new BigNumber(29)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw44.length); _i0_6++) {
        _nw44[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _274_v12 = _nw44;
      let _277_v13;
      _277_v13 = _dafny.Seq.of(_274_v12);
      let _278_v14;
      _278_v14 = _dafny.Map.Empty.slice().updateUnsafe(_262_v0,new BigNumber(468));
      let _279_v15;
      _279_v15 = _dafny.Map.Empty.slice().updateUnsafe((_277_v13)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_262_v0,_262_v0)).length), new BigNumber((_277_v13).length))],(_278_v14).Merge(_278_v14));
      _279_v15 = (_279_v15).update(_274_v12, _278_v14);
      let _280_i1;
      _280_i1 = _dafny.ZERO;
      L2: {
        while (_262_v0) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_280_i1)) {
              break L2;
            }
            _280_i1 = (_280_i1).plus(_dafny.ONE);
            (globalState).f3 = (_dafny.ZERO).minus((_this).f20);
            (globalState).f7 = _262_v0;
            let _281_v16;
            _281_v16 = new _dafny.CodePoint('w'.codePointAt(0));
            (globalState).f5 = !(!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("axeqcyxa"), _281_v16)) || (_262_v0);
            let _282_v17;
            _282_v17 = _dafny.MultiSet.fromElements(_262_v0, true);
            (globalState).f3 = ((_262_v0) ? ((_dafny.ZERO).minus(new BigNumber((_273_v11).length))) : ((((_282_v17).contains(true)) ? ((_282_v17).get(true)) : ((_this).f20))));
          }
        }
      }
      let _283_v18;
      _283_v18 = _dafny.MultiSet.fromElements((_this).f20);
      let _284_v19;
      let _285_v20;
      let _286_v21;
      let _out12;
      let _out13;
      let _out14;
      let _outcollector4 = _module.__default.m0(_283_v18, globalState);
      _out12 = _outcollector4[0];
      _out13 = _outcollector4[1];
      _out14 = _outcollector4[2];
      _284_v19 = _out12;
      _285_v20 = _out13;
      _286_v21 = _out14;
      if (_285_v20) {
        let _287_v22;
        _287_v22 = _dafny.Map.Empty.slice().updateUnsafe(_285_v20,_262_v0);
        let _index41 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_274_v12).length));
        (_274_v12)[_index41] = _287_v22;
        let _index42 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_274_v12).length));
        (_274_v12)[_index42] = ((_287_v22).update(_285_v20, _285_v20)).Merge(_287_v22);
        _268_v7 = _266_v6;
        let _288_v23;
        _288_v23 = new _dafny.CodePoint('s'.codePointAt(0));
        let _289_v24;
        _289_v24 = _dafny.Map.Empty.slice().updateUnsafe(_262_v0,_288_v23);
        (globalState).f3 = (_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(_289_v24,_module.__default.fm2(_262_v0, new BigNumber((_278_v14).length), globalState)), globalState);
        let _source5 = _266_v6;
        let _290___mcc_h0 = _source5;
        let _291_cf0 = _290___mcc_h0;
        _273_v11 = _273_v11;
        let _292_v25;
        let _nw45 = new _module.C0();
        _nw45.__ctor((_this).f20);
        _292_v25 = _nw45;
        let _293_v26;
        let _nw46 = Array((new BigNumber(25)).toNumber()).fill(false);
        _293_v26 = _nw46;
        let _index43 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_293_v26).length));
        (_293_v26)[_index43] = _285_v20;
        let _index44 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_293_v26).length));
        (_293_v26)[_index44] = !((new BigNumber(-390)).isLessThan(((_292_v25).f22).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), function (_294_i2) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        })).length))));
        let _295_v27;
        _295_v27 = _dafny.Seq.of(_292_v25, _292_v25, _292_v25);
        let _rhs35 = (_295_v27)[_module.__default.safeIndex(((_292_v25).f22).minus((_292_v25).f22), new BigNumber((_295_v27).length))];
        let _rhs36 = !((_this).fm0(_dafny.Seq.UnicodeFromString("rdmtla"), ((_292_v25).f22).plus((_dafny.ZERO).minus((_292_v25).f22)), (_this).f20, new BigNumber(-502), globalState));
        let _rhs37 = true;
        let _rhs38 = _262_v0;
        let _lhs33 = globalState;
        let _lhs34 = globalState;
        _292_v25 = _rhs35;
        _lhs33.f7 = _rhs36;
        _lhs34.f5 = _rhs37;
        _285_v20 = _rhs38;
        _286_v21 = (function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of _dafny.IntegerRange(new BigNumber(99), new BigNumber(-193))) {
            let _296_v28 = _compr_16;
            if (((new BigNumber(99)).isLessThanOrEqualTo(_296_v28)) && ((_296_v28).isLessThan(new BigNumber(-193)))) {
              _coll16.push([(_296_v28).plus((_this).f20),_262_v0]);
            }
          }
          return _coll16;
        }()).Merge((_269_v9)[_module.__default.safeIndex((_this).f20, new BigNumber((_269_v9).length))]);
      } else {
        _285_v20 = _262_v0;
        if (!(_285_v20) || (((_285_v20) ? (true) : (_285_v20)))) {
          let _297_v29;
          let _nw47 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _297_v29 = _nw47;
          let _nw48 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _297_v29 = _nw48;
          (globalState).f11 = !(_262_v0) || (((_this).f21)[_module.__default.safeIndex(new BigNumber((_module.__default.fm5((_this).f20, _285_v20, (_this).f20, globalState)).length), new BigNumber(((_this).f21).length))]);
          _297_v29 = _297_v29;
          (globalState).f3 = (_this).f20;
          let _298_v30;
          _298_v30 = new _dafny.CodePoint('f'.codePointAt(0));
          let _299_v31;
          _299_v31 = _dafny.Map.Empty.slice().updateUnsafe(_298_v30,_285_v20);
          (globalState).f11 = (((_299_v31).contains(_298_v30)) ? ((_299_v31).get(_298_v30)) : ((_283_v18).IsDisjointFrom(_283_v18)));
        } else {
          (globalState).f3 = (_this).f20;
          let _300_v32;
          let _nw49 = new _module.C0();
          _nw49.__ctor((_this).f20);
          _300_v32 = _nw49;
          (globalState).f14 = ((_284_v19).Intersect(_284_v19)).IsSubsetOf(_284_v19);
          let _301_v33;
          _301_v33 = new _dafny.CodePoint('s'.codePointAt(0));
          _301_v33 = _301_v33;
          _273_v11 = _dafny.Seq.Concat(_dafny.Seq.update(_273_v11, _module.__default.safeIndex((_this).f20, new BigNumber((_273_v11).length)), _301_v33), _273_v11);
        }
        (globalState).f5 = _285_v20;
        let _302_v34;
        _302_v34 = _dafny.Seq.of(_266_v6);
        let _rhs39 = _285_v20;
        let _rhs40 = (_dafny.Seq.contains(_302_v34, _268_v7)) && (_285_v20);
        let _lhs35 = globalState;
        _262_v0 = _rhs39;
        _lhs35.f7 = _rhs40;
        (globalState).f3 = _module.__default.safeModuloInt((_this).f20, (_this).f20);
      }
      let _source6 = _263_v1;
      let _303___mcc_h1 = _source6;
      let _304_cf0 = _303___mcc_h1;
      let _305_v35;
      let _nw50 = new _module.C0();
      _nw50.__ctor((_dafny.ZERO).minus((_this).f20));
      _305_v35 = _nw50;
      let _306_v36;
      let _init7 = function (_307_i3) {
        return _module.__default.safeDivisionInt(_307_i3, new BigNumber(-681));
      };
      let _nw51 = Array((new BigNumber(17)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw51.length); _i0_7++) {
        _nw51[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _306_v36 = _nw51;
      let _308_v37;
      _308_v37 = _dafny.Seq.of(_306_v36);
      let _309_v38;
      _309_v38 = _dafny.Map.Empty.slice().updateUnsafe(_283_v18,(_308_v37)[_module.__default.safeIndex((_this).f20, new BigNumber((_308_v37).length))]);
      let _310_v39;
      let _nw52 = Array((new BigNumber(8)).toNumber());
      _nw52[(_dafny.ZERO).toNumber()] = _306_v36;
      _nw52[(_dafny.ONE).toNumber()] = (((_309_v38).contains(_dafny.MultiSet.FromArray(_264_v3))) ? ((_309_v38).get(_dafny.MultiSet.FromArray(_264_v3))) : (_306_v36));
      _nw52[(new BigNumber(2)).toNumber()] = _306_v36;
      _nw52[(new BigNumber(3)).toNumber()] = _306_v36;
      _nw52[(new BigNumber(4)).toNumber()] = _306_v36;
      _nw52[(new BigNumber(5)).toNumber()] = _306_v36;
      _nw52[(new BigNumber(6)).toNumber()] = _306_v36;
      _nw52[(new BigNumber(7)).toNumber()] = (_308_v37)[_module.__default.safeIndex((_305_v35).f22, new BigNumber((_308_v37).length))];
      _310_v39 = _nw52;
      let _index45 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_310_v39).length));
      (_310_v39)[_index45] = _306_v36;
      let _index46 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_310_v39).length));
      let _rhs41 = _306_v36;
      let _rhs42 = _268_v7;
      let _rhs43 = (_dafny.ZERO).minus(new BigNumber((_264_v3).length));
      let _lhs36 = _310_v39;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_310_v39).length));
      let _lhs38 = globalState;
      _lhs36[_lhs37] = _rhs41;
      _266_v6 = _rhs42;
      _lhs38.f3 = _rhs43;
      let _311_v40;
      let _nw53 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
      _311_v40 = _nw53;
      let _index47 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_311_v40).length));
      (_311_v40)[_index47] = _266_v6;
      let _index48 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_311_v40).length));
      (_311_v40)[_index48] = _266_v6;
      let _arr0 = (_310_v39)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_310_v39).length))];
      let _index49 = _module.__default.safeIndex(new BigNumber(929), new BigNumber(((_310_v39)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_310_v39).length))]).length));
      _arr0[_index49] = (_this).f20;
      let _312_v41;
      _312_v41 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_262_v0)).length));
      let _313_v42;
      _313_v42 = _dafny.Map.Empty.slice().updateUnsafe(_266_v6,_312_v41);
      let _arr1 = (_310_v39)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_310_v39).length))];
      let _index50 = _module.__default.safeIndex(new BigNumber(929), new BigNumber(((_310_v39)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_310_v39).length))]).length));
      let _rhs44 = (_this).f20;
      let _rhs45 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(((_this).f20).multipliedBy((_this).f20), (_305_v35).f22)));
      let _rhs46 = _313_v42;
      let _rhs47 = _312_v41;
      let _lhs39 = globalState;
      let _lhs40 = (_310_v39)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_310_v39).length))];
      let _lhs41 = _module.__default.safeIndex(new BigNumber(929), new BigNumber(((_310_v39)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_310_v39).length))]).length));
      _lhs39.f3 = _rhs44;
      _lhs40[_lhs41] = _rhs45;
      _313_v42 = _rhs46;
      _312_v41 = _rhs47;
      return;
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
