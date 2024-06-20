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
    static fm2(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false, true));
    };
    static fm6(globalState) {
      return _dafny.Seq.UnicodeFromString("jfyngxuaa");
    };
    static fm7(p0, globalState) {
      return new BigNumber((_dafny.Seq.of(_module.D6.create_DC24(_module.D6.create_DC21(_dafny.MultiSet.fromElements(new BigNumber(198)))), _module.D6.create_DC24(_module.D6.create_DC21(_dafny.MultiSet.fromElements(new BigNumber((function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(155)), function (_0_i0) {
    return _dafny.Seq.UnicodeFromString("crncpbwhy");
  })).Elements) {
    let _1_v0 = _compr_0;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(155)), function (_0_i0) {
      return _dafny.Seq.UnicodeFromString("crncpbwhy");
    }), _1_v0)) {
      _coll0.push([_1_v0,new BigNumber(275)]);
    }
  }
  return _coll0;
}()).length)))), _module.D6.create_DC24(_module.D6.create_DC22(new BigNumber(-556))), _module.D6.create_DC24(_module.D6.create_DC21(_dafny.MultiSet.fromElements(new BigNumber(245), new BigNumber((_dafny.Seq.of(function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(754), new BigNumber(-590), new BigNumber(382), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('t'.codePointAt(0)))).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).length))).length))).Elements) {
    let _2_v1 = _compr_1;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(754), new BigNumber(-590), new BigNumber(382), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('t'.codePointAt(0)))).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).length))).length)), _2_v1)) {
      _coll1.push([_module.__default.safeDivisionInt(_2_v1, new BigNumber(-199)),false]);
    }
  }
  return _coll1;
}())).length)))))).length);
    };
    static fm8(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, true)).length),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(!(false), false)).cardinality()))).length),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Set.fromElements(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))).Elements) {
          let _3_v0 = _compr_2;
          if ((_dafny.Set.fromElements(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))).contains(_3_v0)) {
            _coll2.add(_3_v0);
          }
        }
        return _coll2;
      }()).length)),true));
    };
    static fm9(globalState) {
      if ((!(true)) || (true)) {
        return _dafny.Seq.of(new BigNumber(276));
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(142)), function (_4_i0) {
          return new BigNumber((_dafny.Seq.of(false, false)).length);
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), function (_5_i1) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber(44))).length))).length);
        }));
      }
    };
    static fm10(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(false))));
    };
    static fm11(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(!(false)))).Intersect(_dafny.Set.fromElements(false, true, true));
    };
    static fm12(p0, p1, globalState) {
      return (_module.D10.create_DC34(new _dafny.CodePoint('h'.codePointAt(0)), false, false, true, false)).dtor_cf55;
    };
    static fm14(p0, p1, globalState) {
      return _dafny.Seq.of(new BigNumber(547));
    };
    static fm15(p0, globalState) {
      return false;
    };
    static fm18(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(291)), function (_6_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("luposje")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(624)), function (_7_i1) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("yiwbhd")));
    };
    static fm19(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("thuvag"),new BigNumber(-296))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("njhdgt"),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("sgehlao")).length)))).Merge(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(990)), function (_8_i0) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("em"))).Elements) {
          let _9_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(990)), function (_8_i0) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("em")), _9_v0)) {
            _coll3.push([_9_v0,new BigNumber(322)]);
          }
        }
        return _coll3;
      }()));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      let _source0 = ((true) ? (_module.D6.create_DC23(true, new BigNumber(-860), !(true))) : (_module.D6.create_DC23(true, new BigNumber(663), true)));
      if (_source0.is_DC22) {
        let _10___mcc_h0 = (_source0).cf36;
        let _11_cf36 = _10___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(_11_cf36,true)).Merge(function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of _dafny.IntegerRange(new BigNumber(180), new BigNumber(854))) {
            let _12_v0 = _compr_4;
            if (((new BigNumber(180)).isLessThanOrEqualTo(_12_v0)) && ((_12_v0).isLessThan(new BigNumber(854)))) {
              _coll4.push([_module.__default.safeModuloInt(_12_v0, _11_cf36),false]);
            }
          }
          return _coll4;
        }());
      } else if (_source0.is_DC23) {
        let _13___mcc_h1 = (_source0).cf37;
        let _14___mcc_h2 = (_source0).cf38;
        let _15___mcc_h3 = (_source0).cf39;
        let _16_cf39 = _15___mcc_h3;
        let _17_cf38 = _14___mcc_h2;
        let _18_cf37 = _13___mcc_h1;
        return (function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(_17_cf38,_18_cf37)).Keys.Elements) {
            let _19_v1 = _compr_5;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_17_cf38,_18_cf37)).contains(_19_v1)) {
              _coll5.push([(_19_v1).multipliedBy(_17_cf38),_18_cf37]);
            }
          }
          return _coll5;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_17_cf38,_18_cf37));
      } else if (_source0.is_DC21) {
        let _20___mcc_h4 = (_source0).cf35;
        let _21_cf35 = _20___mcc_h4;
        return function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (function () {
            let _coll7 = new _dafny.Set();
            for (const _compr_7 of (_21_cf35).Elements) {
              let _22_v3 = _compr_7;
              if ((_21_cf35).contains(_22_v3)) {
                _coll7.add((_22_v3).minus(new BigNumber((_dafny.MultiSet.fromElements(true, true, true)).cardinality())));
              }
            }
            return _coll7;
          }()).Elements) {
            let _23_v2 = _compr_6;
            if ((function () {
              let _coll8 = new _dafny.Set();
              for (const _compr_8 of (_21_cf35).Elements) {
                let _24_v3 = _compr_8;
                if ((_21_cf35).contains(_24_v3)) {
                  _coll8.add((_24_v3).minus(new BigNumber((_dafny.MultiSet.fromElements(true, true, true)).cardinality())));
                }
              }
              return _coll8;
            }()).contains(_23_v2)) {
              _coll6.push([_module.__default.safeModuloInt(_23_v2, new BigNumber((_dafny.Seq.UnicodeFromString("udnqqbuso")).length)),!(!(false))]);
            }
          }
          return _coll6;
        }();
      } else {
        let _25___mcc_h5 = (_source0).cf40;
        let _26_cf40 = _25___mcc_h5;
        return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)),true));
      }
    };
    static fm22(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('k'.codePointAt(0));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xunsdlrj")).length)))).cardinality()))).length))).Intersect(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("e")).length))))).Union((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("pqwo")).length),false)).length))).Elements) {
          let _27_v0 = _compr_9;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("pqwo")).length),false)).length))).contains(_27_v0)) {
            _coll9.add(_module.__default.safeDivisionInt(_27_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(296),new BigNumber((_dafny.Set.fromElements(new BigNumber(276), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(157), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("pvxtdfslh")).length),new BigNumber((_dafny.Seq.UnicodeFromString("mlkshnmc")).length))).length), new BigNumber(968), new BigNumber(660), new BigNumber(940)))).cardinality()))).length))).length)));
          }
        }
        return _coll9;
      }()).Union(_dafny.Set.fromElements(new BigNumber(723))));
    };
    static fm24(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(614)), function (_28_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })).length))), _dafny.Seq.of(new BigNumber(743), new BigNumber(349), new BigNumber(228), new BigNumber(-108))), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("gw")).length), new BigNumber(-30))).cardinality())));
    };
    static fm25(p0, p1, globalState) {
      return (function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC10(function () {
  let _coll11 = new _dafny.Map();
  for (const _compr_11 of (function () {
    let _coll12 = new _dafny.Map();
    for (const _compr_12 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(67),new BigNumber(98))).length)), new BigNumber(587))).Elements) {
      let _29_v1 = _compr_12;
      if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(67),new BigNumber(98))).length)), new BigNumber(587)), _29_v1)) {
        _coll12.push([_module.__default.safeDivisionInt(_29_v1, new BigNumber(890)),true]);
      }
    }
    return _coll12;
  }()).Keys.Elements) {
    let _30_v0 = _compr_11;
    if ((function () {
      let _coll13 = new _dafny.Map();
      for (const _compr_13 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(67),new BigNumber(98))).length)), new BigNumber(587))).Elements) {
        let _29_v1 = _compr_13;
        if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(67),new BigNumber(98))).length)), new BigNumber(587)), _29_v1)) {
          _coll13.push([_module.__default.safeDivisionInt(_29_v1, new BigNumber(890)),true]);
        }
      }
      return _coll13;
    }()).contains(_30_v0)) {
      _coll11.push([_module.__default.safeModuloInt(_30_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),false)).length))))).length)),new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, true)).length)))).length)]);
    }
  }
  return _coll11;
}()),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(992)), function (_31_i0) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        })).length))).Keys.Elements) {
          let _32_v2 = _compr_10;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC10(function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of (function () {
    let _coll15 = new _dafny.Map();
    for (const _compr_15 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(67),new BigNumber(98))).length)), new BigNumber(587))).Elements) {
      let _29_v1 = _compr_15;
      if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(67),new BigNumber(98))).length)), new BigNumber(587)), _29_v1)) {
        _coll15.push([_module.__default.safeDivisionInt(_29_v1, new BigNumber(890)),true]);
      }
    }
    return _coll15;
  }()).Keys.Elements) {
    let _30_v0 = _compr_14;
    if ((function () {
      let _coll16 = new _dafny.Map();
      for (const _compr_16 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(67),new BigNumber(98))).length)), new BigNumber(587))).Elements) {
        let _29_v1 = _compr_16;
        if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(67),new BigNumber(98))).length)), new BigNumber(587)), _29_v1)) {
          _coll16.push([_module.__default.safeDivisionInt(_29_v1, new BigNumber(890)),true]);
        }
      }
      return _coll16;
    }()).contains(_30_v0)) {
      _coll14.push([_module.__default.safeModuloInt(_30_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),false)).length))))).length)),new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, true)).length)))).length)]);
    }
  }
  return _coll14;
}()),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(992)), function (_31_i0) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          })).length))).contains(_32_v2)) {
            _coll10.add(_32_v2);
          }
        }
        return _coll10;
      }()).Union(_dafny.Set.fromElements(_module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(575),new BigNumber(899)))));
    };
    static fm26(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(466), new BigNumber(203)), _dafny.Seq.of(new BigNumber(107), new BigNumber((_dafny.Seq.UnicodeFromString("udohgqy")).length)), ((false) ? (_dafny.Seq.of(new BigNumber(591))) : (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(326)), function (_33_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length), new BigNumber((_dafny.Seq.of(true)).length), (_module.D6.create_DC23(false, new BigNumber(174), false)).dtor_cf38, new BigNumber(942)))), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(239)), function (_34_i1) {
        return new BigNumber(-517);
      })).length), new BigNumber((_dafny.Set.fromElements(false)).length)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(685)), function (_35_i2) {
        return new BigNumber(855);
      }), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("mogyijfrl")).length)))));
    };
    static fm27(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(553))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), function (_36_i0) {
        return _dafny.Seq.UnicodeFromString("sf");
      })).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(-295))));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D5.create_DC17(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("eql"),(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, true, false, true, !(true))).cardinality()))));
      if (_source1.is_DC18) {
        let _37___mcc_h0 = (_source1).cf30;
        let _38___mcc_h1 = (_source1).cf31;
        let _39_cf31 = _38___mcc_h1;
        let _40_cf30 = _37___mcc_h0;
        return function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of _dafny.IntegerRange(new BigNumber(936), new BigNumber(244))) {
            let _41_v0 = _compr_17;
            if (((new BigNumber(936)).isLessThanOrEqualTo(_41_v0)) && ((_41_v0).isLessThan(new BigNumber(244)))) {
              _coll17.push([(_41_v0).plus((_dafny.ZERO).minus(_40_cf30)),_40_cf30]);
            }
          }
          return _coll17;
        }();
      } else if (_source1.is_DC19) {
        let _42___mcc_h2 = (_source1).cf32;
        let _43___mcc_h3 = (_source1).cf33;
        let _44_cf33 = _43___mcc_h3;
        let _45_cf32 = _42___mcc_h2;
        return _45_cf32;
      } else if (_source1.is_DC17) {
        let _46___mcc_h4 = (_source1).cf29;
        let _47_cf29 = _46___mcc_h4;
        return (_module.D5.create_DC19(function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of _dafny.IntegerRange(new BigNumber(213), new BigNumber(935))) {
    let _48_v1 = _compr_18;
    if (((new BigNumber(213)).isLessThanOrEqualTo(_48_v1)) && ((_48_v1).isLessThan(new BigNumber(935)))) {
      _coll18.push([(_48_v1).minus(new BigNumber(-704)),new BigNumber(766)]);
    }
  }
  return _coll18;
}(), !(false))).dtor_cf32;
      } else {
        let _49___mcc_h5 = (_source1).cf34;
        let _50_cf34 = _49___mcc_h5;
        return function () {
          let _coll19 = new _dafny.Map();
          for (const _compr_19 of _dafny.IntegerRange(new BigNumber(-975), new BigNumber(332))) {
            let _51_v2 = _compr_19;
            if (((new BigNumber(-975)).isLessThanOrEqualTo(_51_v2)) && ((_51_v2).isLessThan(new BigNumber(332)))) {
              _coll19.push([_module.__default.safeDivisionInt(_51_v2, new BigNumber(-974)),new BigNumber(-753)]);
            }
          }
          return _coll19;
        }();
      }
    };
    static fm29(p0, p1, p2, p3, globalState) {
      return (((true) ? (_dafny.MultiSet.fromElements(new BigNumber(626))) : (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false)).length)))))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-947))));
    };
    static fm30(globalState) {
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(287)), function (_52_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })) : (_dafny.Seq.UnicodeFromString("wuqb"))), _dafny.Seq.UnicodeFromString("nfp"));
    };
    static fm31(p0, globalState) {
      if (true) {
        return _module.D4.create_DC15(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_module.D12.create_DC42(new BigNumber((_dafny.Seq.UnicodeFromString("nekfj")).length))).dtor_cf69)).length), new BigNumber(237));
      } else {
        return _module.D4.create_DC15((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cpwd")).length),true)).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(!(true))).length))).length));
      }
    };
    static fm32(globalState) {
      return _module.D7.create_DC25((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("yhjvyhhbk")).length)), _dafny.Seq.of(new BigNumber(270)), _dafny.Seq.of(new BigNumber(281), new BigNumber((_dafny.Seq.UnicodeFromString("ujoxgmcix")).length)), _dafny.Seq.of(new BigNumber(240)), _dafny.Seq.of(new BigNumber(-985)))).Union(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(815), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(439))).cardinality())), _dafny.Seq.of(new BigNumber(381), new BigNumber((_dafny.Seq.of(!(false))).length)))));
    };
    static fm33(globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('c'.codePointAt(0)));
    };
    static fm36(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(53)), function (_53_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("qsupnh"));
    };
    static fm37(p0, p1, p2, globalState) {
      return !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("ujybiveij"), _dafny.Seq.UnicodeFromString("vvlkj"))) || (((true) ? (true) : (true)));
    };
    static fm38(p0, p1, globalState) {
      return new _dafny.CodePoint('x'.codePointAt(0));
    };
    static fm39(p0, p1, globalState) {
      return (_module.D19.create_DC59(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.FromArray(_dafny.Seq.of(false))))).dtor_cf89;
    };
    static fm40(p0, p1, p2, globalState) {
      if ((new BigNumber((_dafny.Seq.of(false)).length)).isLessThan(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-647)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(442)), function (_54_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-834),false)).length), new BigNumber(258))).cardinality()))) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(337),true);
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-765),true);
      }
    };
    static fm41(p0, p1, p2, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.of(false));
    };
    static fm42(globalState) {
      return _module.D5.create_DC17(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("icyaksiiv"),new BigNumber((_dafny.Seq.UnicodeFromString("hf")).length)));
    };
    static fm43(p0, p1, p2, p3, globalState) {
      return _module.D11.create_DC36(((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(876)), function (_55_i0) {
  return _dafny.MultiSet.fromElements(new BigNumber(-360));
})) : (_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(-909))))));
    };
    static fm44(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(78),!(true))).length),new _dafny.CodePoint('y'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("nsq")).length)),new _dafny.CodePoint('c'.codePointAt(0))))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-871)),new _dafny.CodePoint('d'.codePointAt(0))));
    };
    static fm45(p0, p1, globalState) {
      return _module.D2.create_DC6((new BigNumber(424)).multipliedBy(new BigNumber(-232)));
    };
    static fm46(globalState) {
      if ((false) && (false)) {
        return _dafny.MultiSet.fromElements(_module.D6.create_DC22(new BigNumber(-148)));
      } else {
        return _dafny.MultiSet.fromElements(_module.D6.create_DC22(new BigNumber((_dafny.Seq.UnicodeFromString("mko")).length)), _module.D6.create_DC22(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(858)), function (_56_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length)), _module.D6.create_DC22(new BigNumber((_dafny.MultiSet.fromElements(false, false, true)).cardinality())), _module.D6.create_DC22(new BigNumber((_dafny.Seq.UnicodeFromString("xh")).length)), _module.D6.create_DC22(new BigNumber(936)));
      }
    };
    static fm47(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(21), new BigNumber(456))).length),new BigNumber((_dafny.Seq.UnicodeFromString("aridp")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,function () {
        let _coll20 = new _dafny.Map();
        for (const _compr_20 of _dafny.IntegerRange(new BigNumber(329), new BigNumber(-646))) {
          let _57_v0 = _compr_20;
          if (((new BigNumber(329)).isLessThanOrEqualTo(_57_v0)) && ((_57_v0).isLessThan(new BigNumber(-646)))) {
            _coll20.push([_module.__default.safeDivisionInt(_57_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).length)),new BigNumber(-100)]);
          }
        }
        return _coll20;
      }()));
    };
    static fm48(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(true, false, true, false)), _dafny.Seq.of(_dafny.Seq.of((_module.D2.create_DC7(_dafny.Seq.of((_module.D2.create_DC6((_dafny.ZERO).minus(new BigNumber(-604)))).dtor_cf9, new BigNumber(((_module.D20.create_DC61(function () {
  let _coll21 = new _dafny.Map();
  for (const _compr_21 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("me"), _dafny.Seq.UnicodeFromString("uwx"))).Elements) {
    let _58_v0 = _compr_21;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("me"), _dafny.Seq.UnicodeFromString("uwx")), _58_v0)) {
      _coll21.push([_58_v0,true]);
    }
  }
  return _coll21;
}())).dtor_cf93).length)), !(false), !(!(!(!(false)))), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(true, !(false))).length), new BigNumber(-983), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length), new BigNumber(480), new BigNumber((function () {
  let _coll22 = new _dafny.Map();
  for (const _compr_22 of _dafny.IntegerRange(new BigNumber(-807), new BigNumber(704))) {
    let _59_v1 = _compr_22;
    if (((new BigNumber(-807)).isLessThanOrEqualTo(_59_v1)) && ((_59_v1).isLessThan(new BigNumber(704)))) {
      _coll22.push([_module.__default.safeDivisionInt(_59_v1, new BigNumber(861)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(3)), function (_60_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length)]);
    }
  }
  return _coll22;
}()).length))).cardinality()))).dtor_cf11), _dafny.Seq.of(!(false)), _dafny.Seq.of(true), _dafny.Seq.of(!(false)), _dafny.Seq.of(true)));
    };
    static fm49(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),true);
    };
    static fm50(p0, globalState) {
      return _dafny.MultiSet.FromArray(function (_source2) {
        if (_source2.is_DC13) {
          let _61___mcc_h0 = (_source2).cf22;
          let _62_cf22 = _61___mcc_h0;
          return _dafny.Seq.of(_dafny.Seq.of(true));
        } else if (_source2.is_DC14) {
          let _63___mcc_h1 = (_source2).cf23;
          let _64___mcc_h2 = (_source2).cf24;
          let _65___mcc_h3 = (_source2).cf25;
          let _66_cf25 = _65___mcc_h3;
          let _67_cf24 = _64___mcc_h2;
          let _68_cf23 = _63___mcc_h1;
          return _dafny.Seq.of(_dafny.Seq.of(!(false)), _dafny.Seq.of(true, true), _dafny.Seq.of(false), _dafny.Seq.of(false));
        } else if (_source2.is_DC15) {
          let _69___mcc_h4 = (_source2).cf26;
          let _70___mcc_h5 = (_source2).cf27;
          let _71_cf27 = _70___mcc_h5;
          let _72_cf26 = _69___mcc_h4;
          return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(false, (_module.D15.create_DC51((_module.D6.create_DC23(false, _71_cf27, true)).dtor_cf39, !(false))).dtor_cf83), _dafny.Seq.of(false), _dafny.Seq.of(true, true), _dafny.Seq.of(false, true)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(565)), function (_73_i0) {
            return _dafny.Seq.of(false);
          }));
        } else if (_source2.is_DC12) {
          let _74___mcc_h6 = (_source2).cf21;
          let _75_cf21 = _74___mcc_h6;
          return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(!(false), false, !(false), true, true), _dafny.Seq.of(true, !(true)), _dafny.Seq.of(true)), _dafny.Seq.of(_dafny.Seq.of(!(false))));
        } else {
          let _76___mcc_h7 = (_source2).cf28;
          let _77_cf28 = _76___mcc_h7;
          return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(false)), _dafny.Seq.of(_dafny.Seq.of(true, true, false, false, false)));
        }
      }(_module.D4.create_DC15(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(473), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-197))).length))).length))).length)))).cardinality()), new BigNumber((_dafny.Seq.of(false)).length))).length), new BigNumber(50))));
    };
    static fm51(globalState) {
      if ((!(!(true))) || (false)) {
        return _module.D4.create_DC12(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(933)))));
      } else {
        return _module.D4.create_DC12(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-648),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber(132))));
      }
    };
    static fm52(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat((_module.D13.create_DC44(_dafny.Seq.of(true))).dtor_cf71, _dafny.Seq.of(true)), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true)));
    };
    static fm53(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), function (_78_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(296)), function (_79_i1) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        });
      });
    };
    static fm54(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), function (_80_i0) {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vslhotx")).length));
      })) : (_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ctbyue")).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(554)), function (_81_i1) {
        return new BigNumber(767);
      })))), ((!(false)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(754)), function (_82_i2) {
        return _dafny.Seq.of(new BigNumber(-507));
      })) : (_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("rvii")).length)), _dafny.Seq.of(new BigNumber(-259), new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(803)), function (_83_i3) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length)))).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("rqnhvr")).length))))));
    };
    static fm55(p0, p1, p2, globalState) {
      if ((function () {
        let _coll23 = new _dafny.Set();
        for (const _compr_23 of _dafny.IntegerRange(new BigNumber(73), new BigNumber(703))) {
          let _84_v0 = _compr_23;
          if (((new BigNumber(73)).isLessThanOrEqualTo(_84_v0)) && ((_84_v0).isLessThan(new BigNumber(703)))) {
            _coll23.add((_84_v0).multipliedBy(new BigNumber(697)));
          }
        }
        return _coll23;
      }()).IsSubsetOf(_dafny.Set.fromElements(new BigNumber(952), new BigNumber(130), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)))) {
        return _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("lnfrahgri"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-769)), function (_85_i0) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("as"), _dafny.Seq.UnicodeFromString("jkaix"), _dafny.Seq.UnicodeFromString("ujnwkx"));
      } else {
        return _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(175)), function (_86_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), function (_87_i2) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        }));
      }
    };
    static fm56(p0, globalState) {
      return _module.D0.create_DC1(_module.__default.safeDivisionInt(new BigNumber(-605), new BigNumber(-71)), true, _dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("sg")).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xxkfrh")).length)))), new BigNumber((_dafny.Seq.of(true, false)).length), new _dafny.CodePoint('k'.codePointAt(0)));
    };
    static Main(__noArgsParameter) {
      let _88_v0;
      _88_v0 = _dafny.MultiSet.fromElements(false);
      let _89_v1;
      _89_v1 = _dafny.MultiSet.fromElements(_88_v0, _88_v0);
      let _90_v2;
      _90_v2 = true;
      let _91_v3;
      _91_v3 = _dafny.Seq.of(_90_v2, _90_v2);
      let _92_v4;
      _92_v4 = new BigNumber(166);
      let _93_v5;
      _93_v5 = _dafny.Seq.UnicodeFromString("c");
      let _94_v6;
      _94_v6 = _dafny.MultiSet.fromElements(new BigNumber((_93_v5).length), new BigNumber(476));
      let _95_v7;
      _95_v7 = _dafny.MultiSet.fromElements(new BigNumber((_94_v6).cardinality()));
      let _96_v8;
      _96_v8 = _dafny.Seq.of(_95_v7);
      let _97_v9;
      _97_v9 = _dafny.Seq.of((_96_v8)[_module.__default.safeIndex(_92_v4, new BigNumber((_96_v8).length))]);
      let _98_v10;
      let _init0 = function (_99_i0) {
        return true;
      };
      let _nw0 = Array((new BigNumber(10)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _98_v10 = _nw0;
      let _100_v11;
      let _nw1 = Array((new BigNumber(7)).toNumber());
      _nw1[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_92_v4));
      _nw1[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(656)), function (_101_i1) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })).length);
      _nw1[(new BigNumber(2)).toNumber()] = _92_v4;
      _nw1[(new BigNumber(3)).toNumber()] = _92_v4;
      _nw1[(new BigNumber(4)).toNumber()] = _92_v4;
      _nw1[(new BigNumber(5)).toNumber()] = _92_v4;
      _nw1[(new BigNumber(6)).toNumber()] = _92_v4;
      _100_v11 = _nw1;
      let _102_globalState;
      let _nw2 = new _module.GlobalState();
      _nw2.__ctor(true, _89_v1, _dafny.Seq.of(_90_v2), _dafny.Seq.Concat(_91_v3, _91_v3), false, _dafny.Seq.of(_90_v2), new BigNumber(301), true, new BigNumber(794), new BigNumber(616), false, new BigNumber(333), new BigNumber(349), _dafny.Set.fromElements(_92_v4), new BigNumber(578), (_97_v9)[_module.__default.safeIndex(_92_v4, new BigNumber((_97_v9).length))], new BigNumber(982), new BigNumber(-565), _98_v10, _100_v11, _95_v7, _98_v10);
      _102_globalState = _nw2;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_98_v10).length))) {
        let _103_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_103_i2)) && ((_103_i2).isLessThan(new BigNumber((_98_v10).length))))) {
          (_98_v10)[(_103_i2)] = _90_v2;
        }
      }
      let _104_v12;
      _104_v12 = _dafny.Seq.of(_98_v10, _98_v10);
      let _105_v13;
      let _nw3 = new _module.C4();
      _nw3.__ctor(_90_v2, _104_v12);
      _105_v13 = _nw3;
      let _106_v14;
      _106_v14 = _module.D2.create_DC8(_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("ufigqlgw")), _100_v11, _90_v2);
      let _107_v15;
      _107_v15 = new _dafny.CodePoint('m'.codePointAt(0));
      let _108_v16;
      _108_v16 = _dafny.Map.Empty.slice().updateUnsafe(_107_v15,_90_v2);
      let _109_v17;
      _109_v17 = _dafny.Set.fromElements((_106_v14).dtor_cf16, (((_108_v16).contains(_107_v15)) ? ((_108_v16).get(_107_v15)) : (_90_v2)), _90_v2);
      let _110_v18;
      _110_v18 = _dafny.Seq.of(new BigNumber(-999), _92_v4);
      let _111_v19;
      _111_v19 = _dafny.Map.Empty.slice().updateUnsafe((_109_v17).contains(_105_v13.f28),(_105_v13).fm13(_dafny.Seq.update(_110_v18, _module.__default.safeIndex(_92_v4, new BigNumber((_110_v18).length)), new BigNumber(731)), _90_v2, new BigNumber((_110_v18).length), _102_globalState));
      let _112_v20;
      let _nw4 = Array((new BigNumber(13)).toNumber());
      _112_v20 = _nw4;
      let _113_v21;
      _113_v21 = _module.D11.create_DC37(_112_v20);
      let _114_v22;
      _114_v22 = _dafny.Map.Empty.slice().updateUnsafe(_113_v21,_90_v2);
      _111_v19 = (_111_v19).update((((_114_v22).contains(_113_v21)) ? ((_114_v22).get(_113_v21)) : (_90_v2)), !((_105_v13.f28) && (_90_v2)));
      let _115_v23;
      _115_v23 = _dafny.Set.fromElements(new BigNumber((_module.__default.fm6(_102_globalState)).length), _92_v4, _92_v4, _92_v4, _92_v4);
      let _116_v24;
      _116_v24 = _module.D17.create_DC55(_115_v23);
      let _pat_let_tv0 = _92_v4;
      let _pat_let_tv1 = _105_v13;
      let _pat_let_tv2 = _93_v5;
      let _pat_let_tv3 = _92_v4;
      let _pat_let_tv4 = _92_v4;
      let _pat_let_tv5 = _92_v4;
      let _source3 = function (_source4) {
        if (_source4.is_DC56) {
          let _117___mcc_h6 = (_source4).cf87;
          let _118_cf87 = _117___mcc_h6;
          return _module.D5.create_DC20(_module.D5.create_DC18(_pat_let_tv0, new BigNumber(280)));
        } else {
          let _119___mcc_h7 = (_source4).cf86;
          let _120_cf86 = _119___mcc_h7;
          if (_pat_let_tv1.f28) {
            return _module.D5.create_DC20(_module.D5.create_DC17(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv2,_pat_let_tv3)));
          } else {
            return _module.D5.create_DC20(_module.D5.create_DC18(_pat_let_tv4, _pat_let_tv5));
          }
        }
      }(_116_v24);
      if (_source3.is_DC18) {
        let _121___mcc_h0 = (_source3).cf30;
        let _122___mcc_h1 = (_source3).cf31;
        let _123_cf31 = _122___mcc_h1;
        let _124_cf30 = _121___mcc_h0;
        let _125_v25;
        _125_v25 = _dafny.Map.Empty.slice().updateUnsafe(_90_v2,_124_cf30);
        _125_v25 = _125_v25;
        let _126_v26;
        let _nw5 = new _module.C4();
        _nw5.__ctor(_90_v2, _104_v12);
        _126_v26 = _nw5;
        let _127_v27;
        let _nw6 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
        _127_v27 = _nw6;
        let _128_v28;
        _128_v28 = _dafny.Map.Empty.slice().updateUnsafe(_105_v13.f28,_127_v27);
        let _129_v29;
        _129_v29 = _module.D10.create_DC34(_107_v15, true, _105_v13.f28, !(_105_v13.f28), _126_v26.f28);
        let _130_v30;
        let _nw7 = Array((new BigNumber(10)).toNumber());
        _nw7[(_dafny.ZERO).toNumber()] = (((_128_v28).contains((_129_v29).dtor_cf57)) ? ((_128_v28).get((_129_v29).dtor_cf57)) : (_127_v27));
        _nw7[(_dafny.ONE).toNumber()] = _127_v27;
        _nw7[(new BigNumber(2)).toNumber()] = _127_v27;
        _nw7[(new BigNumber(3)).toNumber()] = _127_v27;
        _nw7[(new BigNumber(4)).toNumber()] = _127_v27;
        _nw7[(new BigNumber(5)).toNumber()] = _127_v27;
        _nw7[(new BigNumber(6)).toNumber()] = ((_105_v13.f28) ? (_127_v27) : (_127_v27));
        _nw7[(new BigNumber(7)).toNumber()] = _127_v27;
        _nw7[(new BigNumber(8)).toNumber()] = _127_v27;
        _nw7[(new BigNumber(9)).toNumber()] = ((_90_v2) ? (_127_v27) : (_127_v27));
        _130_v30 = _nw7;
        let _index0 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_130_v30).length));
        (_130_v30)[_index0] = _127_v27;
        let _index1 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_130_v30).length));
        (_130_v30)[_index1] = _127_v27;
        let _131_v31;
        _131_v31 = _module.D6.create_DC22(new BigNumber(590));
        let _132_v32;
        _132_v32 = _module.D6.create_DC24(_131_v31);
        let _133_v33;
        _133_v33 = _dafny.Map.Empty.slice().updateUnsafe(_100_v11,_132_v32);
        _133_v33 = _133_v33;
      } else if (_source3.is_DC19) {
        let _134___mcc_h2 = (_source3).cf32;
        let _135___mcc_h3 = (_source3).cf33;
        let _136_cf33 = _135___mcc_h3;
        let _137_cf32 = _134___mcc_h2;
        (_102_globalState).f11 = _92_v4;
        let _138_v34;
        let _nw8 = new _module.C5();
        _nw8.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_92_v4,_105_v13.f28), _92_v4);
        _138_v34 = _nw8;
        let _139_v35;
        _139_v35 = _dafny.Seq.of(_115_v23);
        (_102_globalState).f11 = (_92_v4).minus((_dafny.ZERO).minus((_105_v13).fm0(_95_v7, new BigNumber((_dafny.Set.fromElements((_138_v34).f26)).length), _139_v35, _105_v13.f28, _102_globalState)));
        let _140_v36;
        let _init1 = ((_141_v13, _142_v34) => function (_143_i3) {
          return _dafny.Map.Empty.slice().updateUnsafe(_141_v13.f28,(_142_v34).f27);
        })(_105_v13, _138_v34);
        let _nw9 = Array((new BigNumber(10)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw9.length); _i0_1++) {
          _nw9[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _140_v36 = _nw9;
        _140_v36 = _140_v36;
      } else if (_source3.is_DC17) {
        let _144___mcc_h4 = (_source3).cf29;
        let _145_cf29 = _144___mcc_h4;
        let _146_v37;
        let _nw10 = Array((new BigNumber(27)).toNumber());
        _146_v37 = _nw10;
        (_102_globalState).f11 = (((_105_v13.f28) ? (_92_v4) : (_92_v4))).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_105_v13.f28,_146_v37)).length));
        let _147_v38;
        let _nw11 = Array((_dafny.ONE).toNumber());
        _nw11[(_dafny.ZERO).toNumber()] = _93_v5;
        _147_v38 = _nw11;
        let _index2 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_147_v38).length));
        (_147_v38)[_index2] = _93_v5;
        let _index3 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_147_v38).length));
        (_147_v38)[_index3] = _93_v5;
        (_105_v13).f28 = _105_v13.f28;
        let _148_v39;
        let _149_v40;
        let _150_v41;
        let _151_v42;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = (_105_v13).m0(_105_v13.f28, _102_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _148_v39 = _out0;
        _149_v40 = _out1;
        _150_v41 = _out2;
        _151_v42 = _out3;
      } else {
        let _152___mcc_h5 = (_source3).cf34;
        let _153_cf34 = _152___mcc_h5;
        let _154_v43;
        _154_v43 = _dafny.Map.Empty.slice().updateUnsafe(_98_v10,_92_v4);
        let _155_v44;
        _155_v44 = _module.D14.create_DC48(_92_v4, _154_v43, _105_v13.f28);
        let _source5 = _155_v44;
        if (_source5.is_DC48) {
          let _156___mcc_h8 = (_source5).cf78;
          let _157___mcc_h9 = (_source5).cf79;
          let _158___mcc_h10 = (_source5).cf80;
          let _159_cf80 = _158___mcc_h10;
          let _160_cf79 = _157___mcc_h9;
          let _161_cf78 = _156___mcc_h8;
          (_105_v13).f28 = !((_105_v13).fm13(_dafny.Seq.Concat(_110_v18, _110_v18), _105_v13.f28, new BigNumber(201), _102_globalState));
          let _rhs0 = _105_v13.f28;
          let _rhs1 = _105_v13.f28;
          let _rhs2 = (_161_cf78).isEqualTo(((_159_cf80) ? ((_dafny.ZERO).minus(new BigNumber((_93_v5).length))) : (_161_cf78)));
          let _lhs0 = _105_v13;
          let _lhs1 = _105_v13;
          let _lhs2 = _102_globalState;
          _lhs0.f28 = _rhs0;
          _lhs1.f28 = _rhs1;
          _lhs2.f10 = _rhs2;
          _93_v5 = _dafny.Seq.Concat(_93_v5, _dafny.Seq.update(_93_v5, _module.__default.safeIndex(new BigNumber((_115_v23).length), new BigNumber((_93_v5).length)), _107_v15));
          let _162_v45;
          let _163_v46;
          let _164_v47;
          let _165_v48;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = (_105_v13).m0(_105_v13.f28, _102_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _162_v45 = _out4;
          _163_v46 = _out5;
          _164_v47 = _out6;
          _165_v48 = _out7;
        } else {
          let _166___mcc_h11 = (_source5).cf77;
          let _167_cf77 = _166___mcc_h11;
          let _168_v49;
          _168_v49 = _dafny.Map.Empty.slice().updateUnsafe((_module.D3.create_DC11((_91_v3)[_module.__default.safeIndex(_92_v4, new BigNumber((_91_v3).length))], _105_v13.f28)).dtor_cf19,(_dafny.ZERO).minus(_92_v4));
          let _169_v50;
          _169_v50 = _dafny.Seq.of(_115_v23);
          let _170_v51;
          _170_v51 = _dafny.Seq.of(_169_v50);
          _168_v49 = (_168_v49).update(true, (_105_v13).fm0(_95_v7, _92_v4, (_170_v51)[_module.__default.safeIndex(new BigNumber(616), new BigNumber((_170_v51).length))], _105_v13.f28, _102_globalState));
          let _171_v52;
          let _nw12 = new _module.C5();
          _nw12.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_92_v4,_105_v13.f28), _92_v4);
          _171_v52 = _nw12;
          let _172_v53;
          let _nw13 = new _module.C2();
          _nw13.__ctor((_module.D12.create_DC42(_92_v4)).dtor_cf69, _104_v12);
          _172_v53 = _nw13;
          let _173_v54;
          _173_v54 = _dafny.Set.fromElements(_93_v5, _93_v5);
          let _174_v56;
          let _nw14 = new _module.C7();
          _nw14.__ctor(_115_v23, _dafny.Seq.of(_98_v10, _98_v10, _98_v10, _98_v10, (_104_v12)[_module.__default.safeIndex(new BigNumber((_167_cf77).cardinality()), new BigNumber((_104_v12).length))]));
          _174_v56 = _nw14;
          let _175_v57;
          _175_v57 = _dafny.Map.Empty.slice().updateUnsafe(_174_v56,_173_v54);
          let _176_v58;
          _176_v58 = _module.D13.create_DC46((_171_v52).fm5(!(_105_v13.f28), _102_globalState), _174_v56, _92_v4, _dafny.Set.fromElements(_100_v11, _100_v11));
          let _177_v59;
          _177_v59 = _dafny.Map.Empty.slice().updateUnsafe(_93_v5,_105_v13.f28);
          let _178_v62;
          _178_v62 = _dafny.Seq.of(_173_v54, _173_v54);
          let _179_v63;
          let _nw15 = Array((new BigNumber(15)).toNumber());
          _nw15[(_dafny.ZERO).toNumber()] = _173_v54;
          _nw15[(_dafny.ONE).toNumber()] = _173_v54;
          _nw15[(new BigNumber(2)).toNumber()] = (function () {
            let _coll24 = new _dafny.Set();
            for (const _compr_24 of (_173_v54).Elements) {
              let _180_v55 = _compr_24;
              if ((_173_v54).contains(_180_v55)) {
                _coll24.add(_180_v55);
              }
            }
            return _coll24;
          }()).Difference((((_175_v57).contains((_176_v58).dtor_cf74)) ? ((_175_v57).get((_176_v58).dtor_cf74)) : (_173_v54)));
          _nw15[(new BigNumber(3)).toNumber()] = _173_v54;
          _nw15[(new BigNumber(4)).toNumber()] = _173_v54;
          _nw15[(new BigNumber(5)).toNumber()] = function () {
            let _coll25 = new _dafny.Set();
            for (const _compr_25 of (_177_v59).Keys.Elements) {
              let _181_v60 = _compr_25;
              if ((_177_v59).contains(_181_v60)) {
                _coll25.add(_181_v60);
              }
            }
            return _coll25;
          }();
          _nw15[(new BigNumber(6)).toNumber()] = _173_v54;
          _nw15[(new BigNumber(7)).toNumber()] = (_module.__default.fm55(_105_v13.f28, new BigNumber((_91_v3).length), _107_v15, _102_globalState)).Union(_173_v54);
          _nw15[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_93_v5, _93_v5, _93_v5);
          _nw15[(new BigNumber(9)).toNumber()] = _173_v54;
          _nw15[(new BigNumber(10)).toNumber()] = _173_v54;
          _nw15[(new BigNumber(11)).toNumber()] = (_173_v54).Intersect(_173_v54);
          _nw15[(new BigNumber(12)).toNumber()] = _module.__default.fm55(_105_v13.f28, new BigNumber((_93_v5).length), _107_v15, _102_globalState);
          _nw15[(new BigNumber(13)).toNumber()] = (function () {
            let _coll26 = new _dafny.Set();
            for (const _compr_26 of (_173_v54).Elements) {
              let _182_v61 = _compr_26;
              if ((_173_v54).contains(_182_v61)) {
                _coll26.add(_182_v61);
              }
            }
            return _coll26;
          }()).Union((_178_v62)[_module.__default.safeIndex((_172_v53).f29, new BigNumber((_178_v62).length))]);
          _nw15[(new BigNumber(14)).toNumber()] = _173_v54;
          _179_v63 = _nw15;
          let _index4 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_179_v63).length));
          (_179_v63)[_index4] = (_173_v54).Union(_173_v54);
          let _index5 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_179_v63).length));
          (_179_v63)[_index5] = (_173_v54).Difference(_173_v54);
        }
        let _183_v64;
        let _nw16 = Array((new BigNumber(12)).toNumber());
        _nw16[(_dafny.ZERO).toNumber()] = _module.__default.fm38(_90_v2, _92_v4, _102_globalState);
        _nw16[(_dafny.ONE).toNumber()] = _107_v15;
        _nw16[(new BigNumber(2)).toNumber()] = _107_v15;
        _nw16[(new BigNumber(3)).toNumber()] = _107_v15;
        _nw16[(new BigNumber(4)).toNumber()] = _107_v15;
        _nw16[(new BigNumber(5)).toNumber()] = _107_v15;
        _nw16[(new BigNumber(6)).toNumber()] = _107_v15;
        _nw16[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
        _nw16[(new BigNumber(8)).toNumber()] = _107_v15;
        _nw16[(new BigNumber(9)).toNumber()] = _107_v15;
        _nw16[(new BigNumber(10)).toNumber()] = _107_v15;
        _nw16[(new BigNumber(11)).toNumber()] = _107_v15;
        _183_v64 = _nw16;
        let _184_v65;
        _184_v65 = _dafny.Seq.of(_183_v64);
        let _index6 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_98_v10).length));
        (_98_v10)[_index6] = !_dafny.areEqual(_184_v65, _184_v65);
        let _index7 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_98_v10).length));
        let _rhs3 = (new BigNumber((_93_v5).length)).isEqualTo((_92_v4).plus(new BigNumber((_93_v5).length)));
        let _rhs4 = _93_v5;
        let _rhs5 = (_109_v17).equals((_109_v17).Union(_109_v17));
        let _rhs6 = _92_v4;
        let _rhs7 = _107_v15;
        let _lhs3 = _102_globalState;
        let _lhs4 = _98_v10;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_98_v10).length));
        let _lhs6 = _102_globalState;
        _lhs3.f0 = _rhs3;
        _93_v5 = _rhs4;
        _lhs4[_lhs5] = _rhs5;
        _lhs6.f16 = _rhs6;
        _107_v15 = _rhs7;
        let _185_v66;
        let _nw17 = new _module.C3();
        _nw17.__ctor(_104_v12);
        _185_v66 = _nw17;
        _185_v66 = (((_98_v10)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_98_v10).length))]) ? (_185_v66) : (_185_v66));
        let _186_v67;
        let _init2 = function (_187_i4) {
          return _module.D15.create_DC52();
        };
        let _nw18 = Array((new BigNumber(3)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw18.length); _i0_2++) {
          _nw18[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _186_v67 = _nw18;
        let _188_v68;
        _188_v68 = _module.D15.create_DC52();
        let _index8 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_186_v67).length));
        (_186_v67)[_index8] = _188_v68;
        let _189_v69;
        let _init3 = ((_190_v13) => function (_191_i5) {
          return _190_v13.f28;
        })(_105_v13);
        let _nw19 = Array((new BigNumber(27)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw19.length); _i0_3++) {
          _nw19[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _189_v69 = _nw19;
        let _192_v70;
        let _nw20 = new _module.C2();
        _nw20.__ctor((_dafny.ZERO).minus(_92_v4), _dafny.Seq.of(_189_v69, _189_v69));
        _192_v70 = _nw20;
        let _index9 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_186_v67).length));
        let _rhs8 = ((_192_v70).f29).isLessThanOrEqualTo(new BigNumber((_88_v0).cardinality()));
        let _rhs9 = _188_v68;
        let _rhs10 = _192_v70;
        let _lhs7 = _105_v13;
        let _lhs8 = _186_v67;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_186_v67).length));
        _lhs7.f28 = _rhs8;
        _lhs8[_lhs9] = _rhs9;
        _192_v70 = _rhs10;
      }
      let _193_v71;
      _193_v71 = _module.D18.create_DC57(_105_v13);
      let _194_v72;
      _194_v72 = _dafny.Map.Empty.slice().updateUnsafe(_90_v2,_105_v13);
      let _195_v73;
      let _nw21 = Array((new BigNumber(13)).toNumber());
      _nw21[(_dafny.ZERO).toNumber()] = (_193_v71).dtor_cf88;
      _nw21[(_dafny.ONE).toNumber()] = _105_v13;
      _nw21[(new BigNumber(2)).toNumber()] = (((_194_v72).contains(_90_v2)) ? ((_194_v72).get(_90_v2)) : (_105_v13));
      _nw21[(new BigNumber(3)).toNumber()] = _105_v13;
      _nw21[(new BigNumber(4)).toNumber()] = _105_v13;
      _nw21[(new BigNumber(5)).toNumber()] = _105_v13;
      _nw21[(new BigNumber(6)).toNumber()] = _105_v13;
      _nw21[(new BigNumber(7)).toNumber()] = _105_v13;
      _nw21[(new BigNumber(8)).toNumber()] = _105_v13;
      _nw21[(new BigNumber(9)).toNumber()] = (_193_v71).dtor_cf88;
      _nw21[(new BigNumber(10)).toNumber()] = _105_v13;
      _nw21[(new BigNumber(11)).toNumber()] = _105_v13;
      _nw21[(new BigNumber(12)).toNumber()] = _105_v13;
      _195_v73 = _nw21;
      let _index10 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_195_v73).length));
      (_195_v73)[_index10] = _105_v13;
      let _index11 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_195_v73).length));
      (_195_v73)[_index11] = _105_v13;
      if ((_92_v4).isEqualTo(_92_v4)) {
        let _196_v74;
        _196_v74 = _dafny.Map.Empty.slice().updateUnsafe(_105_v13.f28,_92_v4);
        let _197_v75;
        _197_v75 = _dafny.Seq.of(_93_v5);
        _196_v74 = (_196_v74).update(_90_v2, (new BigNumber(((_197_v75)[_module.__default.safeIndex(_92_v4, new BigNumber((_197_v75).length))]).length)).plus(_92_v4));
        let _198_v76;
        let _nw22 = new _module.C3();
        _nw22.__ctor(_dafny.Seq.of(_98_v10, _98_v10, _98_v10, _98_v10, _98_v10));
        _198_v76 = _nw22;
        let _199_v77;
        _199_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_115_v23).length),_98_v10);
        let _200_v78;
        _200_v78 = _dafny.Map.Empty.slice().updateUnsafe((((_199_v77).contains(_92_v4)) ? ((_199_v77).get(_92_v4)) : (_98_v10)),_100_v11);
        _200_v78 = (_200_v78).update(_98_v10, _100_v11);
        (_102_globalState).f16 = _module.__default.safeDivisionInt((new BigNumber((_93_v5).length)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_92_v4,_90_v2)).length)), _92_v4);
        let _201_v79;
        _201_v79 = _dafny.Seq.of(_109_v17, _109_v17, _dafny.Set.fromElements(_105_v13.f28, _105_v13.f28));
        let _202_v80;
        _202_v80 = _module.D12.create_DC42(new BigNumber((_115_v23).length));
        let _203_v81;
        _203_v81 = _dafny.Map.Empty.slice().updateUnsafe(_202_v80,_109_v17);
        let _204_v82;
        let _nw23 = Array((_dafny.ONE).toNumber());
        _nw23[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_201_v79, _dafny.Seq.of((((_203_v81).contains(_module.D12.create_DC42(_92_v4))) ? ((_203_v81).get(_module.D12.create_DC42(_92_v4))) : (_109_v17)), _109_v17));
        _204_v82 = _nw23;
        let _index12 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_204_v82).length));
        (_204_v82)[_index12] = _201_v79;
        let _index13 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_204_v82).length));
        (_204_v82)[_index13] = _201_v79;
      } else {
        let _205_v83;
        let _nw24 = new _module.C0();
        _nw24.__ctor();
        _205_v83 = _nw24;
        let _206_v84;
        _206_v84 = _dafny.MultiSet.fromElements(_205_v83);
        let _207_v85;
        _207_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(113),_module.D12.create_DC41(_206_v84));
        _207_v85 = _207_v85;
        let _208_v86;
        let _nw25 = new _module.C3();
        _nw25.__ctor(_104_v12);
        _208_v86 = _nw25;
        let _rhs11 = (_module.__default.safeDivisionInt(_92_v4, new BigNumber((_dafny.Seq.update(_93_v5, _module.__default.safeIndex(new BigNumber(-282), new BigNumber((_93_v5).length)), _107_v15)).length))).multipliedBy(new BigNumber(208));
        let _rhs12 = _208_v86;
        _92_v4 = _rhs11;
        _208_v86 = _rhs12;
        _93_v5 = _93_v5;
        let _209_v87;
        _209_v87 = _dafny.Map.Empty.slice().updateUnsafe(_91_v3,_92_v4);
        let _210_v88;
        _210_v88 = _dafny.Seq.of(_209_v87, _209_v87, _209_v87);
        (_102_globalState).f8 = (new BigNumber(287)).multipliedBy(new BigNumber(((_210_v88)[_module.__default.safeIndex(_92_v4, new BigNumber((_210_v88).length))]).length));
        let _index14 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_98_v10).length));
        (_98_v10)[_index14] = _module.__default.fm15(_90_v2, _102_globalState);
        let _index15 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_98_v10).length));
        (_98_v10)[_index15] = _105_v13.f28;
      }
      let _211_v89;
      let _212_v90;
      let _213_v91;
      let _214_v92;
      let _out8;
      let _out9;
      let _out10;
      let _out11;
      let _outcollector2 = (_105_v13).m0(_90_v2, _102_globalState);
      _out8 = _outcollector2[0];
      _out9 = _outcollector2[1];
      _out10 = _outcollector2[2];
      _out11 = _outcollector2[3];
      _211_v89 = _out8;
      _212_v90 = _out9;
      _213_v91 = _out10;
      _214_v92 = _out11;
      let _hi0 = new BigNumber(465);
      for (let _215_i6 = (_dafny.ZERO).minus(_92_v4); _215_i6.isLessThan(_hi0); _215_i6 = _215_i6.plus(_dafny.ONE)) {
        let _216_v93;
        _216_v93 = _module.D12.create_DC42(_92_v4);
        _216_v93 = function (_pat_let0_0) {
          return function (_217_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_218_dt__update_hcf69_h0) {
                return _module.D12.create_DC42(_218_dt__update_hcf69_h0);
              }(_pat_let1_0);
            }(_215_i6);
          }(_pat_let0_0);
        }(_216_v93);
        if (true) {
          let _219_v94;
          let _nw26 = new _module.C1();
          _nw26.__ctor(_214_v92, _104_v12);
          _219_v94 = _nw26;
          let _220_v95;
          _220_v95 = _dafny.Seq.of(_219_v94, _219_v94);
          _211_v89 = ((_dafny.ZERO).minus(_92_v4)).isLessThanOrEqualTo(_module.__default.safeDivisionInt(new BigNumber((_220_v95).length), _92_v4));
          let _221_v96;
          let _nw27 = new _module.C4();
          _nw27.__ctor(_211_v89, _dafny.Seq.Concat(_104_v12, _104_v12));
          _221_v96 = _nw27;
          _88_v0 = _88_v0;
          let _222_v97;
          let _nw28 = Array((new BigNumber(3)).toNumber());
          _222_v97 = _nw28;
          let _223_v98;
          _223_v98 = _dafny.Map.Empty.slice().updateUnsafe(_215_i6,true);
          let _224_v99;
          let _nw29 = new _module.C1();
          _nw29.__ctor(new BigNumber((_223_v98).length), _104_v12);
          _224_v99 = _nw29;
          let _index16 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_222_v97).length));
          (_222_v97)[_index16] = _224_v99;
          let _index17 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_222_v97).length));
          (_222_v97)[_index17] = _224_v99;
          let _225_v100;
          let _nw30 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _225_v100 = _nw30;
          let _226_v101;
          _226_v101 = _dafny.MultiSet.fromElements(_225_v100);
          let _index18 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_98_v10).length));
          (_98_v10)[_index18] = (_226_v101).IsProperSubsetOf(_dafny.MultiSet.fromElements(_225_v100, _225_v100));
          let _index19 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_98_v10).length));
          (_98_v10)[_index19] = !(_105_v13.f28);
        } else {
          (_102_globalState).f0 = _105_v13.f28;
          _110_v18 = _110_v18;
          let _227_v102;
          _227_v102 = _dafny.Map.Empty.slice().updateUnsafe(_214_v92,_215_i6);
          let _228_v103;
          _228_v103 = _module.D5.create_DC19(_227_v102, _90_v2);
          let _229_v104;
          let _230_v105;
          let _231_v106;
          let _232_v107;
          let _out12;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = (_105_v13).m0(((_90_v2) ? (!(true)) : ((_228_v103).dtor_cf33)), _102_globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _out15 = _outcollector3[3];
          _229_v104 = _out12;
          _230_v105 = _out13;
          _231_v106 = _out14;
          _232_v107 = _out15;
          let _233_v108;
          _233_v108 = _dafny.Map.Empty.slice().updateUnsafe(_232_v107,_90_v2);
          let _234_v109;
          let _nw31 = new _module.C5();
          _nw31.__ctor(_233_v108, _92_v4);
          _234_v109 = _nw31;
          let _235_v110;
          let _236_v111;
          let _237_v112;
          let _238_v113;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector4 = (_105_v13).m0((_module.__default.fm56(_232_v107, _102_globalState)).dtor_cf1, _102_globalState);
          _out16 = _outcollector4[0];
          _out17 = _outcollector4[1];
          _out18 = _outcollector4[2];
          _out19 = _outcollector4[3];
          _235_v110 = _out16;
          _236_v111 = _out17;
          _237_v112 = _out18;
          _238_v113 = _out19;
        }
        let _index20 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_100_v11).length));
        (_100_v11)[_index20] = _215_i6;
        let _index21 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_100_v11).length));
        (_100_v11)[_index21] = _module.__default.safeDivisionInt(new BigNumber(628), (_92_v4).plus(_215_i6));
        let _239_v114;
        let _nw32 = new _module.C7();
        _nw32.__ctor(_115_v23, _dafny.Seq.of(_98_v10));
        _239_v114 = _nw32;
      }
      if (!((_91_v3)[_module.__default.safeIndex(new BigNumber(254), new BigNumber((_91_v3).length))])) {
        _93_v5 = _dafny.Seq.update(_93_v5, _module.__default.safeIndex((_214_v92).multipliedBy(new BigNumber(-876)), new BigNumber((_93_v5).length)), _107_v15);
        let _240_v115;
        let _241_v116;
        let _242_v117;
        let _243_v118;
        let _out20;
        let _out21;
        let _out22;
        let _out23;
        let _outcollector5 = (_105_v13).m0(_211_v89, _102_globalState);
        _out20 = _outcollector5[0];
        _out21 = _outcollector5[1];
        _out22 = _outcollector5[2];
        _out23 = _outcollector5[3];
        _240_v115 = _out20;
        _241_v116 = _out21;
        _242_v117 = _out22;
        _243_v118 = _out23;
        let _244_v119;
        let _245_v120;
        let _246_v121;
        let _247_v122;
        let _out24;
        let _out25;
        let _out26;
        let _out27;
        let _outcollector6 = (_105_v13).m0((_95_v7).IsProperSubsetOf(_95_v7), _102_globalState);
        _out24 = _outcollector6[0];
        _out25 = _outcollector6[1];
        _out26 = _outcollector6[2];
        _out27 = _outcollector6[3];
        _244_v119 = _out24;
        _245_v120 = _out25;
        _246_v121 = _out26;
        _247_v122 = _out27;
        let _248_v123;
        _248_v123 = _dafny.Set.fromElements(_91_v3, _91_v3, _91_v3);
        _240_v115 = !((_248_v123).Union(_dafny.Set.fromElements(_91_v3, _91_v3))).equals(_248_v123);
        let _249_v124;
        _249_v124 = _module.D17.create_DC56(_90_v2);
        (_102_globalState).f0 = (((_249_v124).dtor_cf87) || (_90_v2)) || (_105_v13.f28);
      } else {
        let _250_v125;
        _250_v125 = _dafny.Set.fromElements(_111_v19, (_111_v19).Merge(_111_v19));
        _250_v125 = _250_v125;
        let _251_v126;
        _251_v126 = _dafny.Map.Empty.slice().updateUnsafe(((!(_105_v13.f28)) ? (_105_v13.f28) : (_211_v89)),_module.__default.fm18(_90_v2, _102_globalState));
        _251_v126 = _251_v126;
        let _252_v127;
        _252_v127 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(475),_214_v92);
        let _253_v128;
        _253_v128 = _dafny.Map.Empty.slice().updateUnsafe(_214_v92,_252_v127);
        let _index22 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_100_v11).length));
        (_100_v11)[_index22] = new BigNumber(((_252_v127).Merge((((_253_v128).contains(_214_v92)) ? ((_253_v128).get(_214_v92)) : (_252_v127)))).length);
        let _index23 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_100_v11).length));
        (_100_v11)[_index23] = new BigNumber(170);
        let _254_v129;
        let _nw33 = new _module.C4();
        _nw33.__ctor(_211_v89, _dafny.Seq.Concat(_104_v12, _104_v12));
        _254_v129 = _nw33;
        (_102_globalState).f8 = _module.__default.fm7(new BigNumber((_115_v23).length), _102_globalState);
      }
      let _255_v130;
      let _256_v131;
      let _257_v132;
      let _258_v133;
      let _out28;
      let _out29;
      let _out30;
      let _out31;
      let _outcollector7 = (_105_v13).m0(_211_v89, _102_globalState);
      _out28 = _outcollector7[0];
      _out29 = _outcollector7[1];
      _out30 = _outcollector7[2];
      _out31 = _outcollector7[3];
      _255_v130 = _out28;
      _256_v131 = _out29;
      _257_v132 = _out30;
      _258_v133 = _out31;
      _93_v5 = _93_v5;
      let _259_v134;
      let _nw34 = new _module.C3();
      _nw34.__ctor(_104_v12);
      _259_v134 = _nw34;
      _90_v2 = ((_214_v92).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_259_v134)).cardinality()))).isLessThanOrEqualTo(_92_v4);
      let _260_i7;
      _260_i7 = _dafny.ZERO;
      L0: {
        while (!(_105_v13.f28)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_260_i7)) {
              break L0;
            }
            _260_i7 = (_260_i7).plus(_dafny.ONE);
            let _index24 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_100_v11).length));
            (_100_v11)[_index24] = _92_v4;
            let _index25 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_100_v11).length));
            (_100_v11)[_index25] = (new BigNumber((_dafny.Seq.UnicodeFromString("xxxok")).length)).multipliedBy(_module.__default.safeModuloInt(_258_v133, _258_v133));
            (_102_globalState).f11 = (_100_v11)[_module.__default.safeIndex(new BigNumber(845), new BigNumber((_100_v11).length))];
            let _261_v135;
            _261_v135 = _dafny.Seq.of(_115_v23, _115_v23);
            (_102_globalState).f0 = (new BigNumber(318)).isEqualTo((_105_v13).fm0(_94_v6, (_100_v11)[_module.__default.safeIndex(new BigNumber(845), new BigNumber((_100_v11).length))], _261_v135, _105_v13.f28, _102_globalState));
            let _index26 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_195_v73).length));
            (_195_v73)[_index26] = (_195_v73)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_195_v73).length))];
          }
        }
      }
      let _262_v136;
      let _263_v137;
      let _out32;
      let _out33;
      let _outcollector8 = (_259_v134).m4(_102_globalState);
      _out32 = _outcollector8[0];
      _out33 = _outcollector8[1];
      _262_v136 = _out32;
      _263_v137 = _out33;
      let _264_v138;
      _264_v138 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(188)), ((_265_v15) => function (_266_i8) {
        return _265_v15;
      })(_107_v15))).length),_211_v89);
      let _267_v139;
      let _nw35 = new _module.C5();
      _nw35.__ctor(_264_v138, (_92_v4).multipliedBy((_dafny.ZERO).minus(_258_v133)));
      _267_v139 = _nw35;
      let _268_v140;
      _268_v140 = _dafny.MultiSet.fromElements(_107_v15, new _dafny.CodePoint('d'.codePointAt(0)));
      let _269_v141;
      _269_v141 = _dafny.Map.Empty.slice().updateUnsafe(_268_v140,_214_v92);
      _269_v141 = (_269_v141).update((_268_v140).update(_107_v15, _module.__default.abs((_267_v139).f27)), _module.__default.safeDivisionInt(new BigNumber((_110_v18).length), _263_v137));
      process.stdout.write(_dafny.toString((_88_v0).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_89_v1).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_90_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_91_v3, _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_92_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_93_v5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v6).equals(_dafny.MultiSet.fromElements(_dafny.ONE, new BigNumber(476)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_95_v7).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_96_v8, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_97_v9, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v10)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_102_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_globalState.f1).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_102_globalState).f2, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_102_globalState).f3, _dafny.Seq.of(true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_102_globalState.f5, _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_102_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_102_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_102_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_102_globalState.f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_globalState.f13).equals(_dafny.Set.fromElements(new BigNumber(166)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f15).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_102_globalState.f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f18)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f19)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f19)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f19)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f19)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f19)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f19)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f19)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f20).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_102_globalState).f21)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_104_v12).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_105_v13.f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_106_v14).dtor_cf14).dtor_cf6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_v14).dtor_cf15)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_v14).dtor_cf15)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_v14).dtor_cf15)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_v14).dtor_cf15)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_v14).dtor_cf15)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_v14).dtor_cf15)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_v14).dtor_cf15)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v14).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_107_v15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_108_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_109_v17).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_110_v18, _dafny.Seq.of(new BigNumber(-999), new BigNumber(166)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_111_v19).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_114_v22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_115_v23).equals(_dafny.Set.fromElements(new BigNumber(9), new BigNumber(166)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_v24).dtor_cf86).equals(_dafny.Set.fromElements(new BigNumber(9), new BigNumber(166)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_193_v71).dtor_cf88.f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_193_v71).dtor_cf88).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_194_v72).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[_dafny.ZERO].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[_dafny.ZERO]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[_dafny.ONE].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[_dafny.ONE]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(2)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(2)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(3)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(3)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(4)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(4)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(5)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(5)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(6)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(6)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(7)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(7)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(8)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(8)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(9)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(9)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(10)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(10)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(11)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(11)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v73)[new BigNumber(12)].f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((((_195_v73)[new BigNumber(12)]).f22).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_211_v89));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_213_v91).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(509), new BigNumber(508), new BigNumber(508), new BigNumber(508), new BigNumber(509)),new _dafny.CodePoint('l'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_214_v92));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_255_v130));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v131)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_257_v132).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(509), new BigNumber(508), new BigNumber(508), new BigNumber(508), new BigNumber(509)),new _dafny.CodePoint('l'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_258_v133));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_260_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_262_v136)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_262_v136)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_262_v136)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_262_v136)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_262_v136)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_262_v136)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_262_v136)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_263_v137));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_264_v138).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(188),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v139).f26).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(188),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_267_v139).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_268_v140).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v141).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))),new BigNumber(4)).updateUnsafe(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))),_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC1(cf0, cf1, cf2, cf3, cf4) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0) && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0();
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
    static create_DC3(cf6) {
      let $dt = new D1(0);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC2(cf5) {
      let $dt = new D1(1);
      $dt.cf5 = cf5;
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
    get dtor_cf6() { return this.cf6; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + this.cf6.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + this.cf5.toVerbatimString(true) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.Seq.UnicodeFromString(""));
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
    static create_DC6(cf9) {
      let $dt = new D2(0);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC7(cf10, cf11, cf12, cf13) {
      let $dt = new D2(1);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC8(cf14, cf15, cf16) {
      let $dt = new D2(2);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC5(cf8) {
      let $dt = new D2(3);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC9(cf17) {
      let $dt = new D2(4);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get is_DC9() { return this.$tag === 4; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 4) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11 && this.cf12 === other.cf12 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15 && this.cf16 === other.cf16;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf8 === other.cf8;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.ZERO);
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
    static create_DC11(cf19, cf20) {
      let $dt = new D3(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC10(cf18) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf19 === other.cf19 && this.cf20 === other.cf20;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11(false, false);
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
    static create_DC13(cf22) {
      let $dt = new D4(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC14(cf23, cf24, cf25) {
      let $dt = new D4(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC15(cf26, cf27) {
      let $dt = new D4(2);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC12(cf21) {
      let $dt = new D4(3);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC16(cf28) {
      let $dt = new D4(4);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get is_DC16() { return this.$tag === 4; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + this.cf25.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 4) {
        return "D4.DC16" + "(" + _dafny.toString(this.cf28) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(_module.D0.Default());
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
    static create_DC18(cf30, cf31) {
      let $dt = new D5(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC19(cf32, cf33) {
      let $dt = new D5(1);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC17(cf29) {
      let $dt = new D5(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC20(cf34) {
      let $dt = new D5(3);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get is_DC20() { return this.$tag === 3; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC19" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC20" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf32, other.cf32) && this.cf33 === other.cf33;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC18(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC22(cf36) {
      let $dt = new D6(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC23(cf37, cf38, cf39) {
      let $dt = new D6(1);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC21(cf35) {
      let $dt = new D6(2);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC24(cf40) {
      let $dt = new D6(3);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC22" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC23" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC21" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC24" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38) && this.cf39 === other.cf39;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC22(_dafny.ZERO);
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
    static create_DC26(cf42, cf43, cf44) {
      let $dt = new D7(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC25(cf41) {
      let $dt = new D7(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC27(cf45) {
      let $dt = new D7(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC26" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC25" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC27" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42) && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC26(_dafny.Seq.of(), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC29(cf47, cf48, cf49) {
      let $dt = new D8(0);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC28(cf46) {
      let $dt = new D8(1);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC29" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC28" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47 && this.cf48 === other.cf48 && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC29(false, false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC31(cf51, cf52) {
      let $dt = new D9(0);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC30(cf50) {
      let $dt = new D9(1);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC32(cf53) {
      let $dt = new D9(2);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC32() { return this.$tag === 2; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC31" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC30" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC32" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf53, other.cf53);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC31(_dafny.ZERO, []);
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
    static create_DC34(cf55, cf56, cf57, cf58, cf59) {
      let $dt = new D10(0);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC35(cf60, cf61) {
      let $dt = new D10(1);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC33(cf54) {
      let $dt = new D10(2);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC34" + "(" + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC35" + "(" + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC33" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf55, other.cf55) && this.cf56 === other.cf56 && this.cf57 === other.cf57 && this.cf58 === other.cf58 && this.cf59 === other.cf59;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf60, other.cf60) && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC34(new _dafny.CodePoint('D'.codePointAt(0)), false, false, false, false);
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
    static create_DC37(cf63) {
      let $dt = new D11(0);
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC38(cf64, cf65) {
      let $dt = new D11(1);
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC39(cf66) {
      let $dt = new D11(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC36(cf62) {
      let $dt = new D11(3);
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC40(cf67) {
      let $dt = new D11(4);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC39() { return this.$tag === 2; }
    get is_DC36() { return this.$tag === 3; }
    get is_DC40() { return this.$tag === 4; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC37" + "(" + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC38" + "(" + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC39" + "(" + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC36" + "(" + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 4) {
        return "D11.DC40" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf63 === other.cf63;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf64, other.cf64) && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf66 === other.cf66;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf67, other.cf67);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC37([]);
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
    static create_DC42(cf69) {
      let $dt = new D12(0);
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC41(cf68) {
      let $dt = new D12(1);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC43(cf70) {
      let $dt = new D12(2);
      $dt.cf70 = cf70;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get is_DC43() { return this.$tag === 2; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf70() { return this.cf70; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC42" + "(" + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC41" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC43" + "(" + _dafny.toString(this.cf70) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf69, other.cf69);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf70, other.cf70);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC42(_dafny.ZERO);
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
    static create_DC45(cf72) {
      let $dt = new D13(0);
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC46(cf73, cf74, cf75, cf76) {
      let $dt = new D13(1);
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC44(cf71) {
      let $dt = new D13(2);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC45" + "(" + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC46" + "(" + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC44" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf72, other.cf72);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf73 === other.cf73 && this.cf74 === other.cf74 && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf71, other.cf71);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC45(_dafny.ZERO);
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
    static create_DC48(cf78, cf79, cf80) {
      let $dt = new D14(0);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      return $dt;
    }
    static create_DC47(cf77) {
      let $dt = new D14(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC48" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC47" + "(" + _dafny.toString(this.cf77) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79) && this.cf80 === other.cf80;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC48(_dafny.ZERO, _dafny.Map.Empty, false);
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
    static create_DC50() {
      let $dt = new D15(0);
      return $dt;
    }
    static create_DC51(cf82, cf83) {
      let $dt = new D15(1);
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC52() {
      let $dt = new D15(2);
      return $dt;
    }
    static create_DC49(cf81) {
      let $dt = new D15(3);
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC53(cf84) {
      let $dt = new D15(4);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get is_DC51() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get is_DC49() { return this.$tag === 3; }
    get is_DC53() { return this.$tag === 4; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC50";
      } else if (this.$tag === 1) {
        return "D15.DC51" + "(" + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC52";
      } else if (this.$tag === 3) {
        return "D15.DC49" + "(" + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 4) {
        return "D15.DC53" + "(" + _dafny.toString(this.cf84) + ")";
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
        return other.$tag === 1 && this.cf82 === other.cf82 && this.cf83 === other.cf83;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf84, other.cf84);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC50();
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
    static create_DC54(cf85) {
      let $dt = new D16(0);
      $dt.cf85 = cf85;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get dtor_cf85() { return this.cf85; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC54" + "(" + _dafny.toString(this.cf85) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf85, other.cf85);
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
    static create_DC56(cf87) {
      let $dt = new D17(0);
      $dt.cf87 = cf87;
      return $dt;
    }
    static create_DC55(cf86) {
      let $dt = new D17(1);
      $dt.cf86 = cf86;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC55() { return this.$tag === 1; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf86() { return this.cf86; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC56" + "(" + _dafny.toString(this.cf87) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC55" + "(" + _dafny.toString(this.cf86) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf87 === other.cf87;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf86, other.cf86);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC56(false);
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
    static create_DC58() {
      let $dt = new D18(0);
      return $dt;
    }
    static create_DC57(cf88) {
      let $dt = new D18(1);
      $dt.cf88 = cf88;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get dtor_cf88() { return this.cf88; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC58";
      } else if (this.$tag === 1) {
        return "D18.DC57" + "(" + _dafny.toString(this.cf88) + ")";
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
        return other.$tag === 1 && this.cf88 === other.cf88;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC58();
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
    static create_DC60(cf90, cf91, cf92) {
      let $dt = new D19(0);
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      return $dt;
    }
    static create_DC59(cf89) {
      let $dt = new D19(1);
      $dt.cf89 = cf89;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get is_DC59() { return this.$tag === 1; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf89() { return this.cf89; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC60" + "(" + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC59" + "(" + _dafny.toString(this.cf89) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf90, other.cf90) && _dafny.areEqual(this.cf91, other.cf91) && _dafny.areEqual(this.cf92, other.cf92);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf89, other.cf89);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC60(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC62(cf94, cf95, cf96) {
      let $dt = new D20(0);
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC61(cf93) {
      let $dt = new D20(1);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC61() { return this.$tag === 1; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC62" + "(" + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC61" + "(" + _dafny.toString(this.cf93) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf94, other.cf94) && this.cf95 === other.cf95 && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf93, other.cf93);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC62(_dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = false;
      this.f1 = _dafny.MultiSet.Empty;
      this.f5 = _dafny.Seq.of();
      this.f6 = _dafny.ZERO;
      this.f8 = _dafny.ZERO;
      this.f10 = false;
      this.f11 = _dafny.ZERO;
      this.f13 = _dafny.Set.Empty;
      this.f16 = _dafny.ZERO;
      this._f2 = _dafny.Seq.of();
      this._f3 = _dafny.Seq.of();
      this._f4 = false;
      this._f7 = false;
      this._f9 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
      this._f14 = _dafny.ZERO;
      this._f15 = _dafny.MultiSet.Empty;
      this._f17 = _dafny.ZERO;
      this._f18 = [];
      this._f19 = [];
      this._f20 = _dafny.MultiSet.Empty;
      this._f21 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this).f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
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
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f9() {
      let _this = this;
      return _this._f9;
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
    get f21() {
      let _this = this;
      return _this._f21;
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
    fm16(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(!(false), false), _dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(false)));
    };
    fm17(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(712)), function (_270_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("kdpncg")), _dafny.Seq.UnicodeFromString("ujpfdfbpa"));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f22 = _dafny.Seq.of();
      this._f30 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
    __ctor(f30, f22) {
      let _this = this;
      (_this)._f30 = f30;
      (_this)._f22 = f22;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f30;
    };
    fm34(p0, globalState) {
      let _this = this;
      return (new BigNumber(259)).multipliedBy((_this).f30);
    };
    fm35(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),(_this).f30)).Merge((function () {
        let _coll27 = new _dafny.Map();
        for (const _compr_27 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), function (_271_i0) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).Elements) {
          let _272_v0 = _compr_27;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), function (_271_i0) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          }), _272_v0)) {
            _coll27.push([_272_v0,(_this).f30]);
          }
        }
        return _coll27;
      }()).Merge(function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of (_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))).Elements) {
          let _273_v1 = _compr_28;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0))), _273_v1)) {
            _coll28.push([_273_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true, true, true))).cardinality()),false)).length)]);
          }
        }
        return _coll28;
      }()))).length);
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.ZERO;
      (globalState).f11 = new BigNumber(664);
      let _274_v0;
      let _nw36 = Array((new BigNumber(20)).toNumber()).fill(false);
      _274_v0 = _nw36;
      let _index27 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_274_v0).length));
      (_274_v0)[_index27] = true;
      let _275_v1;
      _275_v1 = _module.D3.create_DC11(p0, p0);
      let _276_v2;
      _276_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,p0)).length));
      let _277_v3;
      _277_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,_276_v2);
      let _278_v4;
      _278_v4 = new _dafny.CodePoint('o'.codePointAt(0));
      let _279_v5;
      _279_v5 = _module.D0.create_DC1((_this).f30, p0, _277_v3, (_this).f30, _278_v4);
      let _280_v6;
      _280_v6 = _dafny.Seq.of(p0, p0);
      let _pat_let_tv6 = p0;
      let _index28 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_274_v0).length));
      let _rhs13 = !(p0);
      let _rhs14 = function (_source6) {
        if (_source6.is_DC11) {
          let _281___mcc_h0 = (_source6).cf19;
          let _282___mcc_h1 = (_source6).cf20;
          let _283_cf20 = _282___mcc_h1;
          let _284_cf19 = _281___mcc_h0;
          return _283_cf20;
        } else {
          let _285___mcc_h2 = (_source6).cf18;
          let _286_cf18 = _285___mcc_h2;
          return _pat_let_tv6;
        }
      }(_275_v1);
      let _rhs15 = !(((p0) ? ((_this).f30) : ((_279_v5).dtor_cf3))).isEqualTo(new BigNumber((_280_v6).length));
      let _lhs10 = globalState;
      let _lhs11 = globalState;
      let _lhs12 = _274_v0;
      let _lhs13 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_274_v0).length));
      _lhs10.f0 = _rhs13;
      _lhs11.f0 = _rhs14;
      _lhs12[_lhs13] = _rhs15;
      let _index29 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_274_v0).length));
      (_274_v0)[_index29] = ((p0) ? (false) : (((_274_v0)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_274_v0).length))]) && (p0)));
      let _287_v7;
      _287_v7 = _module.D8.create_DC29((new BigNumber((_dafny.Seq.UnicodeFromString("jn")).length)).isLessThanOrEqualTo(new BigNumber((_276_v2).length)), ((_this).f30).isLessThanOrEqualTo((_this).f30), _278_v4);
      _287_v7 = (((_274_v0)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_274_v0).length))]) ? (_module.D8.create_DC29(p0, (_274_v0)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_274_v0).length))], _278_v4)) : (_287_v7));
      (globalState).f16 = (_module.__default.safeModuloInt((_this).f30, (_this).f30)).multipliedBy((_this).f30);
      let _288_v8;
      _288_v8 = _dafny.Set.fromElements((_this).f30);
      let _289_v9;
      _289_v9 = _dafny.MultiSet.fromElements(_274_v0);
      (globalState).f11 = ((new BigNumber((_288_v8).length)).plus((_this).f30)).multipliedBy(new BigNumber(((_289_v9).Difference(_289_v9)).cardinality()));
      r0 = (_274_v0)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_274_v0).length))];
      r1 = _274_v0;
      let _290_v10;
      _290_v10 = _dafny.Seq.of((_this).f30);
      let _291_v11;
      _291_v11 = _dafny.Seq.of(_280_v6);
      let _292_v12;
      _292_v12 = _dafny.Seq.of(new BigNumber((_290_v10).length), (_this).f30, new BigNumber((_291_v11).length));
      let _293_v13;
      _293_v13 = _dafny.Seq.UnicodeFromString("aipcn");
      let _294_v14;
      _294_v14 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_292_v12, _module.__default.safeIndex(new BigNumber((_293_v13).length), new BigNumber((_292_v12).length)), (_this).f30),(((_274_v0)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_274_v0).length))]) ? (_278_v4) : (new _dafny.CodePoint('y'.codePointAt(0)))));
      r2 = _294_v14;
      let _295_v15;
      _295_v15 = _module.D6.create_DC22((_dafny.ZERO).minus((_this).f30));
      r3 = (_this).fm35((_this).f30, (_dafny.ZERO).minus((_this).f30), p0, (_dafny.ZERO).minus((_295_v15).dtor_cf36), globalState);
      return [r0, r1, r2, r3];
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      (globalState).f0 = ((true) ? (p1) : (p1));
      let _296_v0;
      _296_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f30);
      let _297_v1;
      let _nw37 = Array((new BigNumber(20)).toNumber()).fill(false);
      _297_v1 = _nw37;
      let _298_v2;
      _298_v2 = _dafny.MultiSet.fromElements(_297_v1, _297_v1);
      let _299_v3;
      _299_v3 = _dafny.Seq.of(p1, false, p1, p1, p1);
      _296_v0 = (_296_v0).update(_module.__default.safeDivisionInt(new BigNumber((_298_v2).cardinality()), new BigNumber((_dafny.MultiSet.FromArray(_299_v3)).cardinality())), (p2).minus((_this).fm35(new BigNumber(220), p2, false, (_dafny.ZERO).minus(p0), globalState)));
      let _300_v4;
      _300_v4 = _dafny.MultiSet.fromElements(p2);
      (globalState).f0 = !(((false) ? ((_300_v4).IsSubsetOf(_300_v4)) : (!(((p1) ? (p1) : (!(p1)))))));
      let _index30 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_297_v1).length));
      (_297_v1)[_index30] = p1;
      let _index31 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_297_v1).length));
      (_297_v1)[_index31] = p1;
      let _301_v5;
      _301_v5 = _dafny.Map.Empty.slice().updateUnsafe(_297_v1,false);
      let _302_v6;
      _302_v6 = _dafny.Set.fromElements(_299_v3, _dafny.Seq.of(p1, !(p1), !((((_301_v5).contains(_297_v1)) ? ((_301_v5).get(_297_v1)) : ((_297_v1)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_297_v1).length))])))), _299_v3, _299_v3);
      _302_v6 = ((_302_v6).Union(_302_v6)).Difference((_302_v6).Union(_dafny.Set.fromElements(_299_v3, _dafny.Seq.update(_299_v3, _module.__default.safeIndex((_this).fm0(_300_v4, p3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(773)), ((_303_p0) => function (_304_i0) {
        return _dafny.Set.fromElements(_303_p0);
      })(p0)), p1, globalState), new BigNumber((_299_v3).length)), p1), _299_v3)));
      let _305_v7;
      _305_v7 = _dafny.Map.Empty.slice().updateUnsafe((_297_v1)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_297_v1).length))],(_this).f30);
      _305_v7 = (_module.D9.create_DC30(_305_v7)).dtor_cf50;
      r0 = new BigNumber(-425);
      let _306_v8;
      _306_v8 = _dafny.Set.fromElements((_this).f30, p0, p0);
      r1 = _306_v8;
      r2 = p0;
      r3 = p2;
      return [r0, r1, r2, r3];
    }
    m7(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.Set.Empty;
      (globalState).f8 = (_this).f30;
      r2 = p0;
      let _307_v0;
      _307_v0 = _dafny.MultiSet.fromElements(new BigNumber(214), (_this).f30);
      let _308_v1;
      _308_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_307_v0);
      let _309_v2;
      _309_v2 = _module.D4.create_DC12(_308_v1);
      let _310_v3;
      _310_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,_309_v2);
      let _311_v4;
      _311_v4 = _dafny.Seq.of(new BigNumber(112), new BigNumber((_310_v3).length), (_this).f30);
      let _312_v5;
      _312_v5 = _dafny.Seq.of(_311_v4, _dafny.Seq.Concat(_311_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), function (_313_i0) {
        return (_this).f30;
      })), _311_v4);
      _311_v4 = (_312_v5)[_module.__default.safeIndex((_this).f30, new BigNumber((_312_v5).length))];
      let _314_v6;
      _314_v6 = new _dafny.CodePoint('i'.codePointAt(0));
      let _315_v7;
      _315_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,_314_v6);
      let _316_i1;
      _316_i1 = _dafny.ZERO;
      L1: {
        while ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length)).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_315_v7).contains(p0)) ? ((_315_v7).get(p0)) : (_314_v6)),(_this).f30)).length))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_316_i1)) {
              break L1;
            }
            _316_i1 = (_316_i1).plus(_dafny.ONE);
            if (p0) {
              let _317_v8;
              let _nw38 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _317_v8 = _nw38;
              let _318_v9;
              _318_v9 = _dafny.Map.Empty.slice().updateUnsafe(_314_v6,_317_v8);
              _318_v9 = (_318_v9).update(_314_v6, _317_v8);
              _314_v6 = _314_v6;
              r0 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_311_v4, _311_v4)).length))).isLessThanOrEqualTo((_this).f30);
              let _319_v10;
              _319_v10 = _dafny.Seq.UnicodeFromString("frwriukby");
              _319_v10 = _dafny.Seq.Concat(_319_v10, _module.__default.fm36(p0, _dafny.Set.fromElements((_this).f30, (_this).f30), (_this).f30, (_this).f30, globalState));
              let _320_v11;
              let _nw39 = Array((new BigNumber(2)).toNumber()).fill(false);
              _320_v11 = _nw39;
              let _rhs16 = _319_v10;
              let _rhs17 = _314_v6;
              let _rhs18 = _320_v11;
              let _rhs19 = _319_v10;
              _319_v10 = _rhs16;
              _314_v6 = _rhs17;
              _320_v11 = _rhs18;
              _319_v10 = _rhs19;
            } else {
              (globalState).f6 = (_this).f30;
              let _321_v12;
              _321_v12 = _module.D2.create_DC7(_dafny.Seq.update(_311_v4, _module.__default.safeIndex((_this).f30, new BigNumber((_311_v4).length)), (_this).f30), p0, p0, new BigNumber(27));
              (globalState).f0 = (_321_v12).dtor_cf11;
              _314_v6 = _314_v6;
              (globalState).f8 = ((_this).f30).multipliedBy((_this).f30);
              let _322_v13;
              _322_v13 = _dafny.Seq.of(p0);
              let _323_v14;
              _323_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,p0);
              let _324_v15;
              _324_v15 = _dafny.Seq.UnicodeFromString("xw");
              (globalState).f10 = _module.__default.fm37(_322_v13, _323_v14, _dafny.Seq.Concat(_324_v15, _324_v15), globalState);
            }
            (globalState).f16 = (_311_v4)[_module.__default.safeIndex((_this).f30, new BigNumber((_311_v4).length))];
            let _325_v16;
            _325_v16 = _dafny.Seq.UnicodeFromString("kdib");
            _325_v16 = _325_v16;
            r2 = p0;
          }
        }
      }
      let _326_v17;
      _326_v17 = _module.D3.create_DC11(p0, true);
      let _327_i2;
      _327_i2 = _dafny.ZERO;
      L2: {
        while (!((((p0) && (p0)) ? (!(!(p0)) || (p0)) : ((_326_v17).dtor_cf20)))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_327_i2)) {
              break L2;
            }
            _327_i2 = (_327_i2).plus(_dafny.ONE);
            (globalState).f6 = _module.__default.safeModuloInt((new BigNumber((_dafny.Seq.UnicodeFromString("pcye")).length)).minus((_dafny.ZERO).minus((_this).f30)), ((_this).f30).minus((_this).f30));
            let _328_v18;
            let _nw40 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _328_v18 = _nw40;
            let _index32 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_328_v18).length));
            (_328_v18)[_index32] = _module.__default.fm38(p0, (_this).f30, globalState);
            let _index33 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_328_v18).length));
            (_328_v18)[_index33] = _314_v6;
            let _329_v19;
            _329_v19 = _dafny.Seq.UnicodeFromString("whqyh");
            _329_v19 = _329_v19;
            let _330_v20;
            _330_v20 = _dafny.MultiSet.fromElements(p0);
            let _331_v21;
            _331_v21 = _dafny.Set.fromElements(new BigNumber((_330_v20).cardinality()));
            let _332_v22;
            _332_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_329_v19).length));
            let _333_v23;
            let _nw41 = Array((new BigNumber(24)).toNumber());
            _nw41[(_dafny.ZERO).toNumber()] = _329_v19;
            _nw41[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("mx");
            _nw41[(new BigNumber(2)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(3)).toNumber()] = ((p0) ? (_329_v19) : (_329_v19));
            _nw41[(new BigNumber(4)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(5)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_329_v19, _dafny.Seq.UnicodeFromString("foa"));
            _nw41[(new BigNumber(7)).toNumber()] = _module.__default.fm36(p0, _331_v21, new BigNumber((_329_v19).length), (_this).f30, globalState);
            _nw41[(new BigNumber(8)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(9)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(10)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("k"), _module.__default.safeIndex((_this).f30, new BigNumber((_dafny.Seq.UnicodeFromString("k")).length)), (_328_v18)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_328_v18).length))]);
            _nw41[(new BigNumber(12)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(13)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(14)).toNumber()] = _module.__default.fm36(p0, _331_v21, (((_332_v22).contains(p0)) ? ((_332_v22).get(p0)) : ((_this).f30)), (_this).f30, globalState);
            _nw41[(new BigNumber(15)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(16)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(17)).toNumber()] = ((p0) ? (_dafny.Seq.UnicodeFromString("ohl")) : (_329_v19));
            _nw41[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_329_v19, _329_v19);
            _nw41[(new BigNumber(19)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_329_v19, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("wdespalc"), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f30), new BigNumber((_dafny.Seq.UnicodeFromString("wdespalc")).length)), (_328_v18)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_328_v18).length))]));
            _nw41[(new BigNumber(21)).toNumber()] = _329_v19;
            _nw41[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("n"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(433)), ((_334_v6) => function (_335_i3) {
              return _334_v6;
            })(_314_v6)));
            _nw41[(new BigNumber(23)).toNumber()] = _329_v19;
            _333_v23 = _nw41;
            let _index34 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_333_v23).length));
            (_333_v23)[_index34] = _329_v19;
            let _index35 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_333_v23).length));
            (_333_v23)[_index35] = _dafny.Seq.Concat(_329_v19, _329_v19);
          }
        }
      }
      (globalState).f16 = (_this).f30;
      let _336_v24;
      _336_v24 = _dafny.Seq.of(p0, false);
      let _337_v25;
      _337_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,p0);
      let _338_v26;
      _338_v26 = _dafny.Seq.UnicodeFromString("gewwn");
      r0 = !(_module.__default.fm37(_336_v24, _337_v25, _dafny.Seq.update(_338_v26, _module.__default.safeIndex(new BigNumber((_336_v24).length), new BigNumber((_338_v26).length)), _314_v6), globalState)) || (!(p0));
      r1 = (_this).f30;
      r2 = !(p0);
      r3 = (_dafny.Set.fromElements(new BigNumber((_338_v26).length), new BigNumber(169))).Intersect(_dafny.Set.fromElements((_this).f30));
      return [r0, r1, r2, r3];
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f22 = _dafny.Seq.of();
      this._f29 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
    __ctor(f29, f22) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f22 = f22;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f29;
    };
    fm21(p0, p1, globalState) {
      let _this = this;
      return !_dafny.areEqual(_dafny.Seq.UnicodeFromString("qomukech"), _dafny.Seq.UnicodeFromString("eevpwlb"));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.ZERO;
      let _339_v0;
      _339_v0 = _module.D3.create_DC11(p0, ((!(p0)) ? (p0) : (p0)));
      let _source7 = _339_v0;
      if (_source7.is_DC11) {
        let _340___mcc_h0 = (_source7).cf19;
        let _341___mcc_h1 = (_source7).cf20;
        let _342_cf20 = _341___mcc_h1;
        let _343_cf19 = _340___mcc_h0;
        let _344_v1;
        let _nw42 = new _module.C0();
        _nw42.__ctor();
        _344_v1 = _nw42;
        let _345_v2;
        _345_v2 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
        let _346_v3;
        _346_v3 = _dafny.Map.Empty.slice().updateUnsafe(!((((_345_v2).contains(_342_cf20)) ? ((_345_v2).get(_342_cf20)) : (_343_cf19))),p0);
        let _347_v4;
        _347_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f29));
        let _348_v5;
        _348_v5 = new _dafny.CodePoint('h'.codePointAt(0));
        let _349_v6;
        _349_v6 = _module.D0.create_DC1(new BigNumber((_345_v2).length), _343_cf19, _347_v4, (_this).f29, _348_v5);
        let _350_v7;
        _350_v7 = _module.D4.create_DC13(_349_v6);
        let _source8 = _350_v7;
        if (_source8.is_DC13) {
          let _351___mcc_h3 = (_source8).cf22;
          let _352_cf22 = _351___mcc_h3;
          let _353_v8;
          _353_v8 = _dafny.Seq.UnicodeFromString("bgkie");
          let _354_v9;
          _354_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_342_cf20);
          let _355_v10;
          _355_v10 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm22(_342_cf20, !(_342_cf20), (_this).f29, new BigNumber((_module.__default.fm23((_this).f29, new BigNumber((_353_v8).length), !(false), !(_343_cf19), globalState)).length), globalState),((((_354_v9).contains((_dafny.ZERO).minus((_this).f29))) ? ((_354_v9).get((_dafny.ZERO).minus((_this).f29))) : (false))) && (true));
          let _356_v11;
          _356_v11 = _dafny.Set.fromElements(_342_cf20);
          let _rhs20 = (_dafny.Set.fromElements(false)).IsDisjointFrom((_dafny.Set.fromElements(_342_cf20)).Intersect(_356_v11));
          let _rhs21 = _355_v10;
          let _rhs22 = (_this).f29;
          let _lhs14 = globalState;
          _343_cf19 = _rhs20;
          _355_v10 = _rhs21;
          _lhs14.f6 = _rhs22;
          let _357_v12;
          let _nw43 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _357_v12 = _nw43;
          let _358_v13;
          _358_v13 = _dafny.Seq.of((_this).f29);
          let _359_v14;
          _359_v14 = _dafny.MultiSet.fromElements(_342_cf20, _343_cf19);
          let _360_v15;
          _360_v15 = _dafny.Seq.of(false, p0, p0, _343_cf19);
          let _361_v16;
          _361_v16 = _dafny.Seq.of(_358_v13, _module.__default.fm24(_359_v14, new BigNumber(948), (_this).f29, (_360_v15)[_module.__default.safeIndex((_this).f29, new BigNumber((_360_v15).length))], globalState), _358_v13);
          let _index36 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_357_v12).length));
          (_357_v12)[_index36] = new BigNumber((_361_v16).length);
          let _362_v17;
          let _init4 = function (_363_i0) {
            return false;
          };
          let _nw44 = Array((new BigNumber(3)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw44.length); _i0_4++) {
            _nw44[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _362_v17 = _nw44;
          let _index37 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_362_v17).length));
          (_362_v17)[_index37] = !(_342_cf20);
          let _364_v18;
          _364_v18 = _module.D3.create_DC10(_dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f29));
          let _365_v20;
          _365_v20 = _dafny.Set.fromElements(_364_v18, _364_v18, _module.D3.create_DC10(function () {
  let _coll29 = new _dafny.Map();
  for (const _compr_29 of _dafny.IntegerRange(new BigNumber(5), new BigNumber(124))) {
    let _366_v19 = _compr_29;
    if (((new BigNumber(5)).isLessThanOrEqualTo(_366_v19)) && ((_366_v19).isLessThan(new BigNumber(124)))) {
      _coll29.push([_module.__default.safeModuloInt(_366_v19, (_this).f29),(_this).f29]);
    }
  }
  return _coll29;
}()));
          let _367_v21;
          _367_v21 = _dafny.MultiSet.fromElements((_this).f29);
          let _368_v22;
          _368_v22 = _module.D6.create_DC21(_367_v21);
          let _369_v23;
          _369_v23 = _module.D1.create_DC3(_dafny.Seq.UnicodeFromString("hk"));
          let _370_v24;
          _370_v24 = _module.D1.create_DC4(_369_v23);
          let _371_v25;
          _371_v25 = _dafny.Map.Empty.slice().updateUnsafe(_370_v24,(_this).f29);
          let _372_v26;
          _372_v26 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("aiu"),(((_371_v25).contains(_module.D1.create_DC4(_369_v23))) ? ((_371_v25).get(_module.D1.create_DC4(_369_v23))) : ((_dafny.ZERO).minus((_this).f29))));
          let _373_v27;
          _373_v27 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),_module.D5.create_DC17(_372_v26));
          let _374_v28;
          _374_v28 = _dafny.Set.fromElements(_349_v6);
          let _375_v29;
          _375_v29 = _dafny.Set.fromElements(_358_v13, _358_v13, _358_v13);
          let _376_v30;
          _376_v30 = _module.D7.create_DC25(_375_v29);
          let _index38 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_357_v12).length));
          let _index39 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_362_v17).length));
          let _rhs23 = ((_358_v13)[_module.__default.safeIndex((_this).f29, new BigNumber((_358_v13).length))]).minus((((_359_v14).contains(_342_cf20)) ? ((_359_v14).get(_342_cf20)) : ((_this).f29)));
          let _rhs24 = (_module.__default.fm25((_this).f29, (_this).f29, globalState)).IsSubsetOf(_365_v20);
          let _rhs25 = ((_module.__default.fm26(new BigNumber(((_368_v22).dtor_cf35).cardinality()), _373_v27, p0, _343_cf19, globalState)).Difference(_module.__default.fm26((_this).f29, _373_v27, _343_cf19, (_this).fm21(_374_v28, (_this).f29, globalState), globalState))).IsDisjointFrom((_376_v30).dtor_cf41);
          let _rhs26 = new BigNumber(860);
          let _lhs15 = _357_v12;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_357_v12).length));
          let _lhs17 = _362_v17;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_362_v17).length));
          let _lhs19 = globalState;
          _lhs15[_lhs16] = _rhs23;
          _343_cf19 = _rhs24;
          _lhs17[_lhs18] = _rhs25;
          _lhs19.f6 = _rhs26;
          let _377_v31;
          _377_v31 = _module.D1.create_DC3(_353_v8);
          let _pat_let_tv7 = _353_v8;
          let _pat_let_tv8 = _353_v8;
          let _378_v32;
          let _nw45 = Array((new BigNumber(26)).toNumber());
          _nw45[(_dafny.ZERO).toNumber()] = ((p0) ? (_377_v31) : (_377_v31));
          _nw45[(_dafny.ONE).toNumber()] = _377_v31;
          _nw45[(new BigNumber(2)).toNumber()] = _module.D1.create_DC3(_353_v8);
          _nw45[(new BigNumber(3)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(4)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(5)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(6)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(7)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(8)).toNumber()] = _module.D1.create_DC3(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-589)), ((_379_v5) => function (_380_i1) {
  return _379_v5;
})(_348_v5)), _module.__default.safeIndex((_this).f29, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-589)), ((_381_v5) => function (_382_i1) {
  return _381_v5;
})(_348_v5))).length)), _348_v5));
          _nw45[(new BigNumber(9)).toNumber()] = _module.D1.create_DC3((_344_v1).fm17(_353_v8, new BigNumber((_dafny.Seq.UnicodeFromString("uwe")).length), _343_cf19, new BigNumber((_module.__default.fm27(new BigNumber((_dafny.Seq.UnicodeFromString("gu")).length), globalState)).length), globalState));
          _nw45[(new BigNumber(10)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(11)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(12)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(13)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(14)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(15)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(16)).toNumber()] = function (_pat_let2_0) {
            return function (_383_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_384_dt__update_hcf6_h0) {
                  return _module.D1.create_DC3(_384_dt__update_hcf6_h0);
                }(_pat_let3_0);
              }(_pat_let_tv7);
            }(_pat_let2_0);
          }(_377_v31);
          _nw45[(new BigNumber(17)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(18)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(19)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(20)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(21)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(22)).toNumber()] = function (_pat_let4_0) {
            return function (_385_dt__update__tmp_h1) {
              return function (_pat_let5_0) {
                return function (_386_dt__update_hcf6_h1) {
                  return _module.D1.create_DC3(_386_dt__update_hcf6_h1);
                }(_pat_let5_0);
              }(_pat_let_tv8);
            }(_pat_let4_0);
          }(_377_v31);
          _nw45[(new BigNumber(23)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(24)).toNumber()] = _377_v31;
          _nw45[(new BigNumber(25)).toNumber()] = _377_v31;
          _378_v32 = _nw45;
          let _index40 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_378_v32).length));
          (_378_v32)[_index40] = _377_v31;
          let _index41 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_378_v32).length));
          (_378_v32)[_index41] = _377_v31;
          let _387_v33;
          let _nw46 = new _module.C0();
          _nw46.__ctor();
          _387_v33 = _nw46;
        } else if (_source8.is_DC14) {
          let _388___mcc_h4 = (_source8).cf23;
          let _389___mcc_h5 = (_source8).cf24;
          let _390___mcc_h6 = (_source8).cf25;
          let _391_cf25 = _390___mcc_h6;
          let _392_cf24 = _389___mcc_h5;
          let _393_cf23 = _388___mcc_h4;
          _346_v3 = (_346_v3).update(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(975)), function (_394_i2) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }), _391_cf25), p0);
          let _395_v34;
          _395_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), ((_396_cf23) => function (_397_i3) {
            return _396_cf23;
          })(_393_cf23)), _dafny.Seq.UnicodeFromString("xeuvryom")));
          let _398_v35;
          _398_v35 = _dafny.MultiSet.fromElements(!(p0), p0);
          _395_v34 = (_395_v34).update(((false) ? ((_this).f29) : ((_this).f29)), (_398_v35).IsProperSubsetOf(_398_v35));
          let _399_v36;
          _399_v36 = _dafny.MultiSet.fromElements((_this).f29, (_this).f29, (_this).f29, (_this).f29);
          let _400_v37;
          _400_v37 = _dafny.Set.fromElements((_this).f29);
          let _401_v38;
          _401_v38 = _dafny.Seq.of(_400_v37, _400_v37, _400_v37, _400_v37);
          let _402_v39;
          let _nw47 = Array((new BigNumber(7)).toNumber());
          _nw47[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(991)), ((_403_cf23) => function (_404_i4) {
            return _403_cf23;
          })(_393_cf23)), _391_cf25)).length);
          _nw47[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt((_this).fm0(_399_v36, (_this).f29, _401_v38, p0, globalState), (_dafny.ZERO).minus((_this).f29));
          _nw47[(new BigNumber(2)).toNumber()] = (_this).f29;
          _nw47[(new BigNumber(3)).toNumber()] = new BigNumber(-527);
          _nw47[(new BigNumber(4)).toNumber()] = ((_342_cf20) ? ((_this).f29) : ((_this).f29));
          _nw47[(new BigNumber(5)).toNumber()] = new BigNumber((_399_v36).cardinality());
          _nw47[(new BigNumber(6)).toNumber()] = (_this).f29;
          _402_v39 = _nw47;
          let _405_v40;
          _405_v40 = _dafny.Map.Empty.slice().updateUnsafe(_342_cf20,(_this).f29);
          let _index42 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_402_v39).length));
          (_402_v39)[_index42] = (((_405_v40).contains(_343_cf19)) ? ((_405_v40).get(_343_cf19)) : (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(953)), ((_406_cf25) => function (_407_i5) {
            return _406_cf25;
          })(_391_cf25)), _module.__default.safeIndex((_this).f29, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(953)), ((_408_cf25) => function (_409_i5) {
            return _408_cf25;
          })(_391_cf25))).length)), _391_cf25))).cardinality())));
          let _index43 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_402_v39).length));
          let _rhs27 = _module.__default.safeDivisionInt((_this).f29, (_dafny.ZERO).minus((_this).f29));
          let _rhs28 = _342_cf20;
          let _rhs29 = !(_343_cf19) || (p0);
          let _rhs30 = (_this).f29;
          let _rhs31 = _392_cf24;
          let _lhs20 = globalState;
          let _lhs21 = globalState;
          let _lhs22 = globalState;
          let _lhs23 = _402_v39;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_402_v39).length));
          _lhs20.f11 = _rhs27;
          _lhs21.f10 = _rhs28;
          _lhs22.f10 = _rhs29;
          _lhs23[_lhs24] = _rhs30;
          _392_cf24 = _rhs31;
          let _410_v41;
          let _411_v42;
          let _412_v43;
          let _413_v44;
          let _out34;
          let _out35;
          let _out36;
          let _out37;
          let _outcollector9 = (_this).m5((_dafny.ZERO).minus((_this).f29), (_this).f29, globalState);
          _out34 = _outcollector9[0];
          _out35 = _outcollector9[1];
          _out36 = _outcollector9[2];
          _out37 = _outcollector9[3];
          _410_v41 = _out34;
          _411_v42 = _out35;
          _412_v43 = _out36;
          _413_v44 = _out37;
        } else if (_source8.is_DC15) {
          let _414___mcc_h7 = (_source8).cf26;
          let _415___mcc_h8 = (_source8).cf27;
          let _416_cf27 = _415___mcc_h8;
          let _417_cf26 = _414___mcc_h7;
          let _418_v45;
          _418_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f29);
          let _419_v46;
          _419_v46 = _dafny.Map.Empty.slice().updateUnsafe(_416_cf27,_module.D6.create_DC23(_342_cf20, (((_418_v45).contains(new BigNumber(42))) ? ((_418_v45).get(new BigNumber(42))) : ((_this).f29)), p0));
          let _420_v47;
          _420_v47 = _module.D6.create_DC23(p0, _416_cf27, false);
          _419_v46 = (_419_v46).update(new BigNumber(144), _420_v47);
          (globalState).f16 = _417_cf26;
          (globalState).f6 = _module.__default.safeDivisionInt(_417_cf26, _module.__default.safeModuloInt(_417_cf26, _417_cf26));
          let _421_v48;
          _421_v48 = _dafny.MultiSet.fromElements(_417_cf26);
          let _422_v49;
          let _nw48 = Array((new BigNumber(22)).toNumber()).fill(false);
          _422_v49 = _nw48;
          let _423_v50;
          _423_v50 = _dafny.Map.Empty.slice().updateUnsafe(_343_cf19,(_this).f29);
          let _424_v51;
          _424_v51 = _module.D4.create_DC14(_348_v5, _422_v49, _dafny.Seq.update((_344_v1).fm17(_dafny.Seq.Create(_module.__default.abs(new BigNumber(897)), ((_425_v5) => function (_426_i6) {
  return _425_v5;
})(_348_v5)), new BigNumber(496), _342_cf20, _417_cf26, globalState), _module.__default.safeIndex(new BigNumber((_423_v50).length), new BigNumber(((_344_v1).fm17(_dafny.Seq.Create(_module.__default.abs(new BigNumber(897)), ((_427_v5) => function (_428_i6) {
  return _427_v5;
})(_348_v5)), new BigNumber(496), _342_cf20, _417_cf26, globalState)).length)), _348_v5));
          let _429_v52;
          _429_v52 = _module.D4.create_DC16(_424_v51);
          let _430_v53;
          _430_v53 = _dafny.Map.Empty.slice().updateUnsafe(false,_429_v52);
          let _431_v55;
          _431_v55 = _dafny.Set.fromElements(new BigNumber(124));
          (globalState).f11 = (_this).fm0(_421_v48, new BigNumber((_430_v53).length), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(576)), ((_432_cf26) => function (_433_i7) {
            return function () {
              let _coll30 = new _dafny.Set();
              for (const _compr_30 of _dafny.IntegerRange(new BigNumber(742), new BigNumber(-148))) {
                let _434_v54 = _compr_30;
                if (((new BigNumber(742)).isLessThanOrEqualTo(_434_v54)) && ((_434_v54).isLessThan(new BigNumber(-148)))) {
                  _coll30.add(_module.__default.safeModuloInt(_434_v54, _432_cf26));
                }
              }
              return _coll30;
            }();
          })(_417_cf26)), _dafny.Seq.of(_431_v55, _431_v55)), p0, globalState);
        } else if (_source8.is_DC12) {
          let _435___mcc_h9 = (_source8).cf21;
          let _436_cf21 = _435___mcc_h9;
          (globalState).f10 = _343_cf19;
          let _437_v56;
          _437_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(960)), function (_438_i8) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }),(_this).f29);
          let _439_v57;
          _439_v57 = _dafny.Seq.UnicodeFromString("delahctx");
          let _440_v58;
          _440_v58 = _module.D5.create_DC17((_437_v56).update(_439_v57, (_this).f29));
          let _441_v59;
          _441_v59 = _module.D5.create_DC20(_440_v58);
          _441_v59 = _441_v59;
          let _442_v60;
          let _nw49 = Array((new BigNumber(12)).toNumber()).fill(false);
          _442_v60 = _nw49;
          let _index44 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_442_v60).length));
          (_442_v60)[_index44] = false;
          let _443_v61;
          _443_v61 = _dafny.MultiSet.fromElements((_this).f29, (_this).f29);
          let _index45 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_442_v60).length));
          (_442_v60)[_index45] = (_443_v61).IsDisjointFrom(_443_v61);
          let _444_v62;
          let _nw50 = Array((new BigNumber(12)).toNumber()).fill(_module.D3.Default());
          _444_v62 = _nw50;
          let _445_v63;
          _445_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f29);
          let _446_v64;
          _446_v64 = _module.D3.create_DC10(_445_v63);
          let _index46 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_444_v62).length));
          (_444_v62)[_index46] = _446_v64;
          let _447_v65;
          _447_v65 = _dafny.Set.fromElements(_349_v6);
          let _448_v66;
          _448_v66 = _dafny.Map.Empty.slice().updateUnsafe(_439_v57,(_442_v60)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_442_v60).length))]);
          let _449_v67;
          _449_v67 = _dafny.Set.fromElements(new BigNumber((_448_v66).length));
          let _450_v68;
          _450_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f29);
          let _451_v69;
          _451_v69 = _dafny.Seq.of(p0, true);
          let _index47 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_444_v62).length));
          let _rhs32 = _module.D3.create_DC10(_module.__default.fm28(_342_cf20, _343_cf19, !((_this).fm21(_447_v65, new BigNumber((_449_v67).length), globalState)), _450_v68, globalState));
          let _rhs33 = p0;
          let _rhs34 = !((_451_v69)[_module.__default.safeIndex((_this).f29, new BigNumber((_451_v69).length))]);
          let _lhs25 = _444_v62;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_444_v62).length));
          let _lhs27 = globalState;
          _lhs25[_lhs26] = _rhs32;
          _lhs27.f0 = _rhs33;
          _343_cf19 = _rhs34;
        } else {
          let _452___mcc_h10 = (_source8).cf28;
          let _453_cf28 = _452___mcc_h10;
          let _454_v70;
          let _init5 = function (_455_i9) {
            return (_455_i9).minus(new BigNumber((_dafny.Seq.of((_this).f29)).length));
          };
          let _nw51 = Array((new BigNumber(2)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw51.length); _i0_5++) {
            _nw51[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _454_v70 = _nw51;
          _454_v70 = ((!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("roqfhujdu"), _348_v5)) ? (_454_v70) : (_454_v70));
          let _456_v71;
          let _nw52 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
          _456_v71 = _nw52;
          let _457_v72;
          _457_v72 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vsdcwur"),_456_v71);
          _457_v72 = (_457_v72).update(_dafny.Seq.UnicodeFromString("tlmwfrvcm"), _456_v71);
          (globalState).f8 = (_this).f29;
          let _458_v73;
          _458_v73 = _dafny.Seq.UnicodeFromString("frayqgv");
          let _rhs35 = ((_this).f29).isLessThan((_this).f29);
          let _rhs36 = _454_v70;
          let _rhs37 = _342_cf20;
          let _rhs38 = (_344_v1).fm17(_458_v73, (_this).f29, _dafny.Seq.IsPrefixOf(_458_v73, _458_v73), (_this).f29, globalState);
          r0 = _rhs35;
          _454_v70 = _rhs36;
          _343_cf19 = _rhs37;
          _458_v73 = _rhs38;
        }
        let _459_v74;
        let _nw53 = Array((new BigNumber(13)).toNumber()).fill([]);
        _459_v74 = _nw53;
        let _460_v75;
        _460_v75 = _dafny.Seq.of(p0);
        let _461_v76;
        _461_v76 = _dafny.Map.Empty.slice().updateUnsafe(_459_v74,!(!((_460_v75)[_module.__default.safeIndex((_this).f29, new BigNumber((_460_v75).length))]) || (p0)));
        _461_v76 = (_461_v76).update(_459_v74, p0);
        let _462_v77;
        let _nw54 = Array((new BigNumber(26)).toNumber()).fill(false);
        _462_v77 = _nw54;
        let _index48 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_462_v77).length));
        (_462_v77)[_index48] = ((p0) ? (_342_cf20) : (false));
        let _463_v78;
        _463_v78 = _dafny.Seq.UnicodeFromString("xljlu");
        let _index49 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_462_v77).length));
        let _rhs39 = new BigNumber((_463_v78).length);
        let _rhs40 = _343_cf19;
        let _lhs28 = globalState;
        let _lhs29 = _462_v77;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_462_v77).length));
        _lhs28.f11 = _rhs39;
        _lhs29[_lhs30] = _rhs40;
      } else {
        let _464___mcc_h2 = (_source7).cf18;
        let _465_cf18 = _464___mcc_h2;
        let _466_v79;
        _466_v79 = _module.D2.create_DC6((_this).f29);
        let _467_v80;
        _467_v80 = _module.D2.create_DC9(_466_v79);
        let _468_v81;
        _468_v81 = new _dafny.CodePoint('j'.codePointAt(0));
        let _469_v82;
        let _nw55 = Array((new BigNumber(18)).toNumber()).fill(_module.D2.Default());
        _469_v82 = _nw55;
        let _470_v83;
        _470_v83 = _dafny.MultiSet.fromElements(_469_v82);
        let _471_v84;
        _471_v84 = _dafny.Map.Empty.slice().updateUnsafe(_468_v81,_470_v83);
        let _472_v85;
        _472_v85 = _dafny.Map.Empty.slice().updateUnsafe(_467_v80,(((_471_v84).contains(_468_v81)) ? ((_471_v84).get(_468_v81)) : ((_470_v83).update(_469_v82, _module.__default.abs((_this).f29)))));
        let _473_v86;
        _473_v86 = _dafny.Seq.of(_470_v83);
        _472_v85 = (_472_v85).update(_467_v80, ((_473_v86)[_module.__default.safeIndex((_this).f29, new BigNumber((_473_v86).length))]).Intersect(_470_v83));
        let _474_v87;
        _474_v87 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,p0);
        let _475_v88;
        _475_v88 = _dafny.Seq.of((((_474_v87).contains((_this).f29)) ? ((_474_v87).get((_this).f29)) : (p0)));
        let _476_v89;
        _476_v89 = _dafny.Seq.UnicodeFromString("lklrnmhq");
        let _477_v90;
        let _nw56 = Array((new BigNumber(14)).toNumber());
        _nw56[(_dafny.ZERO).toNumber()] = p0;
        _nw56[(_dafny.ONE).toNumber()] = (_475_v88)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_476_v89).length))).cardinality()), new BigNumber((_475_v88).length))];
        _nw56[(new BigNumber(2)).toNumber()] = !(p0);
        _nw56[(new BigNumber(3)).toNumber()] = true;
        _nw56[(new BigNumber(4)).toNumber()] = p0;
        _nw56[(new BigNumber(5)).toNumber()] = (_475_v88)[_module.__default.safeIndex((_this).f29, new BigNumber((_475_v88).length))];
        _nw56[(new BigNumber(6)).toNumber()] = p0;
        _nw56[(new BigNumber(7)).toNumber()] = p0;
        _nw56[(new BigNumber(8)).toNumber()] = true;
        _nw56[(new BigNumber(9)).toNumber()] = p0;
        _nw56[(new BigNumber(10)).toNumber()] = p0;
        _nw56[(new BigNumber(11)).toNumber()] = p0;
        _nw56[(new BigNumber(12)).toNumber()] = p0;
        _nw56[(new BigNumber(13)).toNumber()] = p0;
        _477_v90 = _nw56;
        let _478_v91;
        let _nw57 = Array((new BigNumber(6)).toNumber());
        _nw57[(_dafny.ZERO).toNumber()] = _477_v90;
        _nw57[(_dafny.ONE).toNumber()] = _477_v90;
        _nw57[(new BigNumber(2)).toNumber()] = _477_v90;
        _nw57[(new BigNumber(3)).toNumber()] = ((p0) ? (_477_v90) : (_477_v90));
        _nw57[(new BigNumber(4)).toNumber()] = _477_v90;
        _nw57[(new BigNumber(5)).toNumber()] = _477_v90;
        _478_v91 = _nw57;
        let _index50 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_478_v91).length));
        (_478_v91)[_index50] = _477_v90;
        let _index51 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_478_v91).length));
        let _init6 = ((_479_p0) => function (_480_i10) {
          return _479_p0;
        })(p0);
        let _nw58 = Array((new BigNumber(14)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw58.length); _i0_6++) {
          _nw58[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        (_478_v91)[_index51] = _nw58;
        let _481_v92;
        _481_v92 = _dafny.MultiSet.fromElements(p0);
        let _482_v93;
        let _483_v94;
        let _484_v95;
        let _485_v96;
        let _out38;
        let _out39;
        let _out40;
        let _out41;
        let _outcollector10 = (_this).m5((_dafny.ZERO).minus((_this).f29), (_dafny.ZERO).minus(_module.__default.safeModuloInt((((_465_cf18).contains(new BigNumber((_481_v92).cardinality()))) ? ((_465_cf18).get(new BigNumber((_481_v92).cardinality()))) : (new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()))), (_this).f29)), globalState);
        _out38 = _outcollector10[0];
        _out39 = _outcollector10[1];
        _out40 = _outcollector10[2];
        _out41 = _outcollector10[3];
        _482_v93 = _out38;
        _483_v94 = _out39;
        _484_v95 = _out40;
        _485_v96 = _out41;
        let _486_v97;
        let _init7 = ((_487_v96) => function (_488_i11) {
          return (_488_i11).multipliedBy(_487_v96);
        })(_485_v96);
        let _nw59 = Array((new BigNumber(23)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw59.length); _i0_7++) {
          _nw59[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _486_v97 = _nw59;
        let _489_v98;
        _489_v98 = _dafny.Seq.of(_486_v97, _486_v97, _486_v97);
        let _490_v99;
        let _nw60 = new _module.C0();
        _nw60.__ctor();
        _490_v99 = _nw60;
        _486_v97 = (_489_v98)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_482_v93,_490_v99)).length), new BigNumber((_489_v98).length))];
      }
      let _491_v100;
      let _nw61 = Array((new BigNumber(17)).toNumber()).fill([]);
      _491_v100 = _nw61;
      let _492_v101;
      let _nw62 = Array((new BigNumber(16)).toNumber()).fill(false);
      _492_v101 = _nw62;
      let _nw63 = Array((new BigNumber(7)).toNumber());
      _nw63[(_dafny.ZERO).toNumber()] = _492_v101;
      _nw63[(_dafny.ONE).toNumber()] = _492_v101;
      _nw63[(new BigNumber(2)).toNumber()] = _492_v101;
      _nw63[(new BigNumber(3)).toNumber()] = _492_v101;
      _nw63[(new BigNumber(4)).toNumber()] = ((_this).f22)[_module.__default.safeIndex((_this).f29, new BigNumber(((_this).f22).length))];
      _nw63[(new BigNumber(5)).toNumber()] = _492_v101;
      _nw63[(new BigNumber(6)).toNumber()] = _492_v101;
      _491_v100 = _nw63;
      r3 = (_module.__default.safeDivisionInt((_this).f29, new BigNumber((_module.__default.fm29(p0, new BigNumber(277), p0, (_this).f29, globalState)).cardinality()))).minus(new BigNumber(638));
      let _hi1 = new BigNumber(388);
      for (let _493_i12 = ((p0) ? ((_this).f29) : ((_this).f29)); _493_i12.isLessThan(_hi1); _493_i12 = _493_i12.plus(_dafny.ONE)) {
        let _494_v102;
        _494_v102 = new _dafny.CodePoint('u'.codePointAt(0));
        let _495_v103;
        _495_v103 = _dafny.Seq.of(_494_v102);
        _495_v103 = _dafny.Seq.Concat(_495_v103, _495_v103);
        (globalState).f16 = _493_i12;
        (globalState).f6 = _493_i12;
        let _496_v104;
        _496_v104 = _dafny.MultiSet.fromElements((_this).f29, new BigNumber(680));
        let _497_v106;
        _497_v106 = _dafny.Seq.of(p0, p0);
        let _498_v107;
        _498_v107 = _dafny.Seq.of(new BigNumber((_497_v106).length), (_this).f29);
        let _499_v108;
        _499_v108 = _module.D4.create_DC14(_494_v102, _492_v101, _dafny.Seq.UnicodeFromString("gxovnuk"));
        let _500_v109;
        _500_v109 = _dafny.MultiSet.fromElements(_499_v108);
        let _501_v110;
        _501_v110 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_500_v109).cardinality()));
        let _502_v111;
        _502_v111 = _dafny.Set.fromElements(new BigNumber(12), _493_i12, (_498_v107)[_module.__default.safeIndex(new BigNumber((_498_v107).length), new BigNumber((_498_v107).length))]);
        let _503_v112;
        _503_v112 = _dafny.Map.Empty.slice().updateUnsafe(p0,_502_v111);
        let _504_v113;
        _504_v113 = _dafny.Seq.of((((_503_v112).contains(p0)) ? ((_503_v112).get(p0)) : (_502_v111)), _502_v111);
        let _505_v114;
        _505_v114 = _dafny.Set.fromElements((_this).fm0(_dafny.MultiSet.FromArray(_498_v107), new BigNumber((_501_v110).length), _504_v113, false, globalState));
        let _506_v115;
        _506_v115 = _dafny.Seq.of(function () {
          let _coll31 = new _dafny.Set();
          for (const _compr_31 of _dafny.IntegerRange(new BigNumber(984), new BigNumber(889))) {
            let _507_v105 = _compr_31;
            if (((new BigNumber(984)).isLessThanOrEqualTo(_507_v105)) && ((_507_v105).isLessThan(new BigNumber(889)))) {
              _coll31.add((_507_v105).plus(_493_i12));
            }
          }
          return _coll31;
        }(), _505_v114);
        let _508_v116;
        _508_v116 = _module.D6.create_DC22((_this).fm0(_496_v104, new BigNumber((_dafny.Seq.UnicodeFromString("ekh")).length), _504_v113, p0, globalState));
        let _509_v117;
        _509_v117 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xxitwh")).length),(_this).f29);
        let _510_v118;
        _510_v118 = _dafny.Map.Empty.slice().updateUnsafe(p0,_509_v117);
        let _511_v119;
        _511_v119 = _module.D0.create_DC1((_508_v116).dtor_cf36, p0, _510_v118, _493_i12, new _dafny.CodePoint('h'.codePointAt(0)));
        let _512_v120;
        _512_v120 = _module.D4.create_DC13(_511_v119);
        _512_v120 = _512_v120;
      }
      let _513_v121;
      _513_v121 = _dafny.Set.fromElements(p0);
      if ((_513_v121).IsSubsetOf(_513_v121)) {
        let _514_v122;
        _514_v122 = _dafny.Seq.UnicodeFromString("bknlragrj");
        let _515_v123;
        _515_v123 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_514_v122).length));
        (globalState).f6 = _module.__default.safeDivisionInt(new BigNumber(69), new BigNumber((_515_v123).length));
        let _516_v124;
        _516_v124 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm30(globalState),false);
        (globalState).f0 = ((((_515_v123).contains(p0)) ? ((_515_v123).get(p0)) : ((_this).f29))).isLessThan((_dafny.ZERO).minus(new BigNumber(((_516_v124).Merge(_516_v124)).length)));
        let _517_v125;
        let _nw64 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _517_v125 = _nw64;
        let _518_v127;
        _518_v127 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_dafny.Set.fromElements((_this).f29))).contains(function () {
          let _coll32 = new _dafny.Set();
          for (const _compr_32 of _dafny.IntegerRange(new BigNumber(-266), new BigNumber(398))) {
            let _519_v126 = _compr_32;
            if (((new BigNumber(-266)).isLessThanOrEqualTo(_519_v126)) && ((_519_v126).isLessThan(new BigNumber(398)))) {
              _coll32.add(_module.__default.safeModuloInt(_519_v126, (_dafny.ZERO).minus((_this).f29)));
            }
          }
          return _coll32;
        }()),_517_v125);
        _517_v125 = (((_518_v127).contains(p0)) ? ((_518_v127).get(p0)) : (_517_v125));
        let _520_v128;
        _520_v128 = _module.D2.create_DC5(_517_v125);
        let _pat_let_tv9 = _517_v125;
        let _521_v129;
        _521_v129 = _dafny.Map.Empty.slice().updateUnsafe(_492_v101,(function (_pat_let6_0) {
          return function (_522_dt__update__tmp_h2) {
            return function (_pat_let7_0) {
              return function (_523_dt__update_hcf8_h0) {
                return _module.D2.create_DC5(_523_dt__update_hcf8_h0);
              }(_pat_let7_0);
            }(_pat_let_tv9);
          }(_pat_let6_0);
        }(_520_v128)).dtor_cf8);
        (globalState).f6 = (_dafny.ZERO).minus(new BigNumber(((_521_v129).Merge((_dafny.Map.Empty.slice().updateUnsafe(_492_v101,_517_v125)).update(_492_v101, _517_v125))).length));
        if (!(p0) || (p0)) {
          let _524_v130;
          _524_v130 = _dafny.MultiSet.fromElements((_this).f29);
          (globalState).f10 = ((_524_v130).IsDisjointFrom(_524_v130)) === (p0);
          let _525_v131;
          _525_v131 = _dafny.Set.fromElements(_492_v101, _492_v101, _492_v101);
          _525_v131 = ((p0) ? (_525_v131) : ((_dafny.Set.fromElements(_492_v101)).Intersect(_525_v131)));
          let _index52 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_517_v125).length));
          (_517_v125)[_index52] = (_this).f29;
          let _index53 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_492_v101).length));
          (_492_v101)[_index53] = p0;
          let _index54 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_517_v125).length));
          (_517_v125)[_index54] = (_this).f29;
          let _526_v132;
          _526_v132 = _dafny.Seq.of(p0);
          let _527_v133;
          _527_v133 = _dafny.Set.fromElements(new BigNumber(686), new BigNumber(316));
          let _528_v134;
          _528_v134 = _dafny.Seq.of((_this).fm0(_524_v130, new BigNumber((_526_v132).length), _dafny.Seq.of(_527_v133), p0, globalState));
          let _529_v135;
          _529_v135 = _module.D6.create_DC22((_this).f29);
          let _index55 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_517_v125).length));
          let _index56 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_492_v101).length));
          let _index57 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_517_v125).length));
          let _rhs41 = new BigNumber((_528_v134).length);
          let _rhs42 = (_module.__default.fm31(false, globalState)).dtor_cf27;
          let _rhs43 = p0;
          let _rhs44 = (((p0) ? (_529_v135) : (_529_v135))).dtor_cf36;
          let _lhs31 = _517_v125;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_517_v125).length));
          let _lhs33 = globalState;
          let _lhs34 = _492_v101;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_492_v101).length));
          let _lhs36 = _517_v125;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_517_v125).length));
          _lhs31[_lhs32] = _rhs41;
          _lhs33.f11 = _rhs42;
          _lhs34[_lhs35] = _rhs43;
          _lhs36[_lhs37] = _rhs44;
          let _530_v136;
          _530_v136 = _dafny.Seq.of((_527_v133).Union(_527_v133));
          _530_v136 = _dafny.Seq.Concat(_530_v136, _530_v136);
          let _531_v137;
          let _nw65 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _531_v137 = _nw65;
          let _532_v138;
          _532_v138 = _dafny.Seq.of(_531_v137);
          let _533_v139;
          _533_v139 = _module.D8.create_DC28(_532_v138);
          let _534_v140;
          _534_v140 = _module.D8.create_DC28((_533_v139).dtor_cf46);
          let _535_v141;
          let _nw66 = Array((new BigNumber(21)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = _532_v138;
          _nw66[(_dafny.ONE).toNumber()] = _532_v138;
          _nw66[(new BigNumber(2)).toNumber()] = ((p0) ? (_dafny.Seq.of(_531_v137)) : (_dafny.Seq.of(_531_v137)));
          _nw66[(new BigNumber(3)).toNumber()] = _532_v138;
          _nw66[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_531_v137, _531_v137, _531_v137);
          _nw66[(new BigNumber(5)).toNumber()] = _532_v138;
          _nw66[(new BigNumber(6)).toNumber()] = _532_v138;
          _nw66[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_532_v138, _532_v138);
          _nw66[(new BigNumber(8)).toNumber()] = _532_v138;
          _nw66[(new BigNumber(9)).toNumber()] = (_534_v140).dtor_cf46;
          _nw66[(new BigNumber(10)).toNumber()] = _532_v138;
          _nw66[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_532_v138, _532_v138);
          _nw66[(new BigNumber(12)).toNumber()] = _532_v138;
          _nw66[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(_531_v137);
          _nw66[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_532_v138, _dafny.Seq.of(_531_v137, _531_v137));
          _nw66[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_532_v138, _532_v138);
          _nw66[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_532_v138, _532_v138);
          _nw66[(new BigNumber(17)).toNumber()] = _532_v138;
          _nw66[(new BigNumber(18)).toNumber()] = _532_v138;
          _nw66[(new BigNumber(19)).toNumber()] = _532_v138;
          _nw66[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_532_v138, _532_v138);
          _535_v141 = _nw66;
          let _index58 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_535_v141).length));
          (_535_v141)[_index58] = _532_v138;
          let _index59 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_535_v141).length));
          (_535_v141)[_index59] = _dafny.Seq.Concat(_dafny.Seq.Concat(_532_v138, _dafny.Seq.of(_531_v137, _531_v137, _531_v137, _531_v137, _531_v137)), _dafny.Seq.of(_531_v137, _531_v137, _531_v137, _531_v137, (((_518_v127).contains(p0)) ? ((_518_v127).get(p0)) : (_531_v137))));
        } else {
          (globalState).f8 = (_this).f29;
          let _536_v142;
          let _nw67 = Array((new BigNumber(29)).toNumber()).fill([]);
          _536_v142 = _nw67;
          let _537_v143;
          _537_v143 = _dafny.Seq.of((_this).f29, (_this).f29);
          let _538_v144;
          _538_v144 = _dafny.Set.fromElements(_537_v143);
          let _539_v145;
          _539_v145 = _module.D7.create_DC25(_538_v144);
          let _pat_let_tv10 = _537_v143;
          let _pat_let_tv11 = _537_v143;
          let _540_v146;
          let _nw68 = Array((new BigNumber(24)).toNumber());
          _nw68[(_dafny.ZERO).toNumber()] = _539_v145;
          _nw68[(_dafny.ONE).toNumber()] = _539_v145;
          _nw68[(new BigNumber(2)).toNumber()] = _module.__default.fm32(globalState);
          _nw68[(new BigNumber(3)).toNumber()] = _module.D7.create_DC25(_dafny.Set.fromElements(_537_v143));
          _nw68[(new BigNumber(4)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(5)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(6)).toNumber()] = _module.D7.create_DC25(_538_v144);
          _nw68[(new BigNumber(7)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(8)).toNumber()] = _module.D7.create_DC25(_538_v144);
          _nw68[(new BigNumber(9)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(10)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(11)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(12)).toNumber()] = _module.D7.create_DC25(_538_v144);
          _nw68[(new BigNumber(13)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(14)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(15)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(16)).toNumber()] = function (_pat_let8_0) {
            return function (_541_dt__update__tmp_h3) {
              return function (_pat_let9_0) {
                return function (_542_dt__update_hcf41_h0) {
                  return _module.D7.create_DC25(_542_dt__update_hcf41_h0);
                }(_pat_let9_0);
              }(_dafny.Set.fromElements(_pat_let_tv10, _pat_let_tv11));
            }(_pat_let8_0);
          }(_539_v145);
          _nw68[(new BigNumber(17)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(18)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(19)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(20)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(21)).toNumber()] = _539_v145;
          _nw68[(new BigNumber(22)).toNumber()] = _module.D7.create_DC25(_538_v144);
          _nw68[(new BigNumber(23)).toNumber()] = _539_v145;
          _540_v146 = _nw68;
          let _index60 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_536_v142).length));
          (_536_v142)[_index60] = _540_v146;
          let _543_v147;
          _543_v147 = _dafny.Seq.of(_540_v146, _540_v146);
          let _index61 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_536_v142).length));
          (_536_v142)[_index61] = (_543_v147)[_module.__default.safeIndex((((_515_v123).contains(p0)) ? ((_515_v123).get(p0)) : (new BigNumber(-113))), new BigNumber((_543_v147).length))];
          let _544_v148;
          _544_v148 = _module.D2.create_DC6((_this).f29);
          let _545_v149;
          _545_v149 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f29).multipliedBy(new BigNumber(-50)),((_544_v148).dtor_cf9).multipliedBy((_this).f29));
          _545_v149 = (_545_v149).update(_module.__default.safeModuloInt((_this).f29, new BigNumber((_514_v122).length)), (_this).f29);
          let _index62 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_492_v101).length));
          (_492_v101)[_index62] = p0;
          let _546_v150;
          _546_v150 = _dafny.MultiSet.fromElements((_this).f29, (_this).f29);
          let _547_v151;
          _547_v151 = _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_514_v122).length)));
          let _index63 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_492_v101).length));
          (_492_v101)[_index63] = (((_this).fm0(_546_v150, new BigNumber((_514_v122).length), _dafny.Seq.update(_547_v151, _module.__default.safeIndex((_this).f29, new BigNumber((_547_v151).length)), _dafny.Set.fromElements(new BigNumber((_module.__default.fm33(globalState)).length))), !(p0), globalState)).isLessThanOrEqualTo(new BigNumber((_514_v122).length))) === (true);
          let _548_v152;
          let _nw69 = Array((new BigNumber(27)).toNumber()).fill(false);
          _548_v152 = _nw69;
          let _549_v153;
          let _nw70 = new _module.C1();
          _nw70.__ctor((_dafny.ZERO).minus((_this).f29), _dafny.Seq.of(_548_v152));
          _549_v153 = _nw70;
          let _550_v154;
          let _nw71 = Array((new BigNumber(16)).toNumber());
          _nw71[(_dafny.ZERO).toNumber()] = _549_v153;
          _nw71[(_dafny.ONE).toNumber()] = _549_v153;
          _nw71[(new BigNumber(2)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(3)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(4)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(5)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(6)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(7)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(8)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(9)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(10)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(11)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(12)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(13)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(14)).toNumber()] = _549_v153;
          _nw71[(new BigNumber(15)).toNumber()] = _549_v153;
          _550_v154 = _nw71;
          let _index64 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_550_v154).length));
          (_550_v154)[_index64] = _549_v153;
          let _index65 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_550_v154).length));
          (_550_v154)[_index65] = _549_v153;
        }
      } else {
        let _551_v155;
        let _nw72 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _551_v155 = _nw72;
        let _552_v156;
        _552_v156 = _dafny.Seq.UnicodeFromString("rmsqxsyv");
        let _553_v157;
        _553_v157 = _dafny.Seq.of(_552_v156, _552_v156);
        let _index66 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_551_v155).length));
        (_551_v155)[_index66] = (_dafny.ZERO).minus(new BigNumber(((_553_v157)[_module.__default.safeIndex((_this).f29, new BigNumber((_553_v157).length))]).length));
        let _554_v158;
        _554_v158 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (!(p0)) : (p0)),(_this).f29);
        let _555_v159;
        _555_v159 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,(_this).f29);
        let _556_v160;
        _556_v160 = _dafny.Set.fromElements((_this).f29);
        let _557_v161;
        _557_v161 = _dafny.Seq.of(new BigNumber((_556_v160).length));
        let _index67 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_551_v155).length));
        (_551_v155)[_index67] = (((_554_v158).contains(p0)) ? ((_554_v158).get(p0)) : ((_dafny.ZERO).minus(_module.__default.safeDivisionInt((((_555_v159).contains((_this).f29)) ? ((_555_v159).get((_this).f29)) : (new BigNumber((_557_v161).length))), (_this).f29))));
        let _558_v162;
        _558_v162 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.update(_557_v161, _module.__default.safeIndex((_551_v155)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_551_v155).length))], new BigNumber((_557_v161).length)), (_551_v155)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_551_v155).length))])),_557_v161);
        let _559_v163;
        _559_v163 = _dafny.MultiSet.fromElements(false);
        let _560_v164;
        _560_v164 = _dafny.Seq.of(_module.__default.fm24(_559_v163, (_551_v155)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_551_v155).length))], new BigNumber((_dafny.Seq.of((_551_v155)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_551_v155).length))])).length), p0, globalState), _557_v161, _dafny.Seq.of((_this).f29), _557_v161, _557_v161);
        let _561_v165;
        _561_v165 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_560_v164)[_module.__default.safeIndex((_551_v155)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_551_v155).length))], new BigNumber((_560_v164).length))]);
        let _562_v166;
        _562_v166 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_561_v165);
        _557_v161 = (((_558_v162).contains(((((_562_v166).contains((_this).f29)) ? ((_562_v166).get((_this).f29)) : (_561_v165))).Merge(_561_v165))) ? ((_558_v162).get(((((_562_v166).contains((_this).f29)) ? ((_562_v166).get((_this).f29)) : (_561_v165))).Merge(_561_v165))) : (_dafny.Seq.Concat(_557_v161, _dafny.Seq.Create(_module.__default.abs(new BigNumber(244)), function (_563_i13) {
          return (_this).f29;
        }))));
        let _564_v167;
        let _init8 = ((_565_v159) => function (_566_i14) {
          return _module.D3.create_DC10(_565_v159);
        })(_555_v159);
        let _nw73 = Array((new BigNumber(7)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw73.length); _i0_8++) {
          _nw73[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _564_v167 = _nw73;
        let _567_v168;
        _567_v168 = _module.D3.create_DC10(_555_v159);
        let _index68 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_564_v167).length));
        (_564_v167)[_index68] = _567_v168;
        let _568_v170;
        _568_v170 = _dafny.MultiSet.fromElements(new BigNumber((_552_v156).length));
        let _index69 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_564_v167).length));
        (_564_v167)[_index69] = _module.D3.create_DC10((_555_v159).Merge(function () {
  let _coll33 = new _dafny.Map();
  for (const _compr_33 of (_568_v170).Elements) {
    let _569_v169 = _compr_33;
    if ((_568_v170).contains(_569_v169)) {
      _coll33.push([_module.__default.safeDivisionInt(_569_v169, (_551_v155)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_551_v155).length))]),(_this).f29]);
    }
  }
  return _coll33;
}()));
        let _570_v171;
        _570_v171 = new _dafny.CodePoint('s'.codePointAt(0));
        let _571_v172;
        _571_v172 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-994)), ((_572_v171) => function (_573_i15) {
          return _572_v171;
        })(_570_v171)));
        let _574_v173;
        let _nw74 = Array((new BigNumber(6)).toNumber());
        _nw74[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_552_v156, _module.__default.safeIndex(new BigNumber((_552_v156).length), new BigNumber((_552_v156).length)), _570_v171);
        _nw74[(_dafny.ONE).toNumber()] = _552_v156;
        _nw74[(new BigNumber(2)).toNumber()] = (((_571_v172).contains(p0)) ? ((_571_v172).get(p0)) : (_552_v156));
        _nw74[(new BigNumber(3)).toNumber()] = _552_v156;
        _nw74[(new BigNumber(4)).toNumber()] = _552_v156;
        _nw74[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_552_v156, _dafny.Seq.UnicodeFromString("jwtvlaj"));
        _574_v173 = _nw74;
        let _index70 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_574_v173).length));
        (_574_v173)[_index70] = _552_v156;
        let _575_v174;
        _575_v174 = _dafny.Map.Empty.slice().updateUnsafe(p0,_555_v159);
        let _576_v175;
        _576_v175 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,p0);
        let _577_v176;
        _577_v176 = _module.D0.create_DC1((_551_v155)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_551_v155).length))], p0, _575_v174, new BigNumber((_576_v175).length), _570_v171);
        let _578_v177;
        _578_v177 = _module.D4.create_DC13(_577_v176);
        let _579_v178;
        _579_v178 = _module.D4.create_DC16(_578_v177);
        let _580_v179;
        _580_v179 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
        let _581_v180;
        _581_v180 = _dafny.Map.Empty.slice().updateUnsafe(_579_v178,(((_580_v179).contains(false)) ? ((_580_v179).get(false)) : (p0)));
        let _582_v181;
        _582_v181 = _module.D10.create_DC33(_581_v180);
        let _index71 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_574_v173).length));
        (_574_v173)[_index71] = _dafny.Seq.update(_module.__default.fm30(globalState), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_582_v181).dtor_cf54).length)), new BigNumber((_module.__default.fm30(globalState)).length)), _570_v171);
        r3 = ((new BigNumber((_dafny.Set.fromElements(_551_v155)).length)).minus((_this).f29)).plus((_dafny.ZERO).minus(((_this).f29).minus((_this).f29)));
      }
      let _583_v182;
      _583_v182 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,p0);
      let _584_v183;
      _584_v183 = _dafny.Seq.UnicodeFromString("kdkvyfx");
      let _585_v184;
      _585_v184 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm37(_dafny.Seq.of(p0), _583_v182, _584_v183, globalState),new BigNumber((_584_v183).length));
      let _586_v185;
      _586_v185 = _dafny.Seq.of(_585_v184, _585_v184);
      let _587_v186;
      _587_v186 = _module.D9.create_DC30((_586_v185)[_module.__default.safeIndex((_this).f29, new BigNumber((_586_v185).length))]);
      let _588_v187;
      _588_v187 = _module.D9.create_DC32(_587_v186);
      let _source9 = _588_v187;
      if (_source9.is_DC31) {
        let _589___mcc_h11 = (_source9).cf51;
        let _590___mcc_h12 = (_source9).cf52;
        let _591_cf52 = _590___mcc_h12;
        let _592_cf51 = _589___mcc_h11;
        let _593_v188;
        _593_v188 = _dafny.MultiSet.fromElements(new BigNumber((_513_v121).length));
        let _594_v189;
        _594_v189 = _dafny.Seq.of(_592_cf51);
        let _595_v190;
        _595_v190 = _dafny.Set.fromElements(new BigNumber((_594_v189).length));
        let _596_v191;
        _596_v191 = _dafny.Map.Empty.slice().updateUnsafe(p0,_595_v190);
        let _597_v193;
        _597_v193 = _dafny.Seq.of(_595_v190, _595_v190, _595_v190, _595_v190, (((_596_v191).contains(p0)) ? ((_596_v191).get(p0)) : (function () {
          let _coll34 = new _dafny.Set();
          for (const _compr_34 of (_595_v190).Elements) {
            let _598_v192 = _compr_34;
            if ((_595_v190).contains(_598_v192)) {
              _coll34.add((_598_v192).plus(new BigNumber(492)));
            }
          }
          return _coll34;
        }())));
        let _599_v194;
        _599_v194 = _dafny.Seq.of(_592_cf51, (_this).fm0(_593_v188, (_this).f29, _597_v193, p0, globalState));
        let _600_v195;
        _600_v195 = _module.D2.create_DC6(new BigNumber((_599_v194).length));
        let _601_v196;
        _601_v196 = _module.D2.create_DC9(_600_v195);
        let _602_v197;
        _602_v197 = _dafny.Set.fromElements(_601_v196, _601_v196);
        _602_v197 = _602_v197;
        let _603_v198;
        _603_v198 = _dafny.Map.Empty.slice().updateUnsafe(_513_v121,new _dafny.CodePoint('u'.codePointAt(0)));
        let _604_v199;
        _604_v199 = new _dafny.CodePoint('t'.codePointAt(0));
        _603_v198 = (_603_v198).update(_513_v121, _604_v199);
        let _605_v200;
        _605_v200 = _dafny.Map.Empty.slice().updateUnsafe(_513_v121,p0);
        let _606_v201;
        _606_v201 = _module.D8.create_DC29(p0, (((_605_v200).contains(_513_v121)) ? ((_605_v200).get(_513_v121)) : (p0)), (_584_v183)[_module.__default.safeIndex((_this).f29, new BigNumber((_584_v183).length))]);
        _606_v201 = _606_v201;
        _599_v194 = _599_v194;
      } else if (_source9.is_DC30) {
        let _607___mcc_h13 = (_source9).cf50;
        let _608_cf50 = _607___mcc_h13;
        r0 = p0;
        let _609_v202;
        let _610_v203;
        let _611_v204;
        let _612_v205;
        let _out42;
        let _out43;
        let _out44;
        let _out45;
        let _outcollector11 = (_this).m5((_this).f29, (_this).f29, globalState);
        _out42 = _outcollector11[0];
        _out43 = _outcollector11[1];
        _out44 = _outcollector11[2];
        _out45 = _outcollector11[3];
        _609_v202 = _out42;
        _610_v203 = _out43;
        _611_v204 = _out44;
        _612_v205 = _out45;
        let _613_v206;
        _613_v206 = _dafny.MultiSet.fromElements(_610_v203);
        let _614_v207;
        _614_v207 = _dafny.Map.Empty.slice().updateUnsafe(_610_v203,_613_v206);
        let _615_v208;
        _615_v208 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(993),_614_v207);
        let _616_v209;
        _616_v209 = _dafny.Seq.of(_610_v203);
        let _index72 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length));
        (_492_v101)[_index72] = !((((_615_v208).contains(_609_v202)) ? ((_615_v208).get(_609_v202)) : (_614_v207))).equals(_module.__default.fm39((_this).f29, _616_v209, globalState));
        let _index73 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length));
        let _rhs45 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), ((_617_p0) => function (_618_i16) {
          return ((_617_p0) ? (new _dafny.CodePoint('t'.codePointAt(0))) : (new _dafny.CodePoint('l'.codePointAt(0))));
        })(p0));
        let _rhs46 = new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber(737), new BigNumber(722), (_this).f29, new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)))).length))).update(_609_v202, _module.__default.abs(new BigNumber(-639)))).cardinality());
        let _rhs47 = _513_v121;
        let _rhs48 = !(_611_v204);
        let _rhs49 = _584_v183;
        let _lhs38 = globalState;
        let _lhs39 = _492_v101;
        let _lhs40 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length));
        _584_v183 = _rhs45;
        _lhs38.f11 = _rhs46;
        _513_v121 = _rhs47;
        _lhs39[_lhs40] = _rhs48;
        _584_v183 = _rhs49;
        if (_module.__default.fm37(_616_v209, _583_v182, _584_v183, globalState)) {
          let _index74 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length));
          (_492_v101)[_index74] = p0;
          let _619_v210;
          _619_v210 = _dafny.Set.fromElements(new BigNumber(-516), (_this).f29);
          let _620_v212;
          _620_v212 = _dafny.Seq.of(new BigNumber((_584_v183).length));
          let _621_v213;
          let _nw75 = Array((new BigNumber(22)).toNumber());
          _nw75[(_dafny.ZERO).toNumber()] = _619_v210;
          _nw75[(_dafny.ONE).toNumber()] = (_619_v210).Difference(_dafny.Set.fromElements(new BigNumber((_584_v183).length)));
          _nw75[(new BigNumber(2)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(3)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements((_this).f29, _612_v205, new BigNumber(60), new BigNumber((_584_v183).length), _609_v202);
          _nw75[(new BigNumber(5)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(_609_v202);
          _nw75[(new BigNumber(7)).toNumber()] = function () {
            let _coll35 = new _dafny.Set();
            for (const _compr_35 of _dafny.IntegerRange(new BigNumber(-907), new BigNumber(-883))) {
              let _622_v211 = _compr_35;
              if (((new BigNumber(-907)).isLessThanOrEqualTo(_622_v211)) && ((_622_v211).isLessThan(new BigNumber(-883)))) {
                _coll35.add((_622_v211).multipliedBy(new BigNumber((_513_v121).length)));
              }
            }
            return _coll35;
          }();
          _nw75[(new BigNumber(8)).toNumber()] = (_619_v210).Difference(_619_v210);
          _nw75[(new BigNumber(9)).toNumber()] = (_dafny.Set.fromElements(_612_v205)).Union(_619_v210);
          _nw75[(new BigNumber(10)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(11)).toNumber()] = (_619_v210).Difference(_619_v210);
          _nw75[(new BigNumber(12)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(13)).toNumber()] = _module.__default.fm23(_609_v202, _609_v202, p0, (_492_v101)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length))], globalState);
          _nw75[(new BigNumber(14)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(15)).toNumber()] = (_619_v210).Difference(_619_v210);
          _nw75[(new BigNumber(16)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(17)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(18)).toNumber()] = (_module.__default.fm23(new BigNumber((_619_v210).length), (_620_v212)[_module.__default.safeIndex((_this).f29, new BigNumber((_620_v212).length))], true, _611_v204, globalState)).Intersect(_619_v210);
          _nw75[(new BigNumber(19)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(20)).toNumber()] = _619_v210;
          _nw75[(new BigNumber(21)).toNumber()] = _619_v210;
          _621_v213 = _nw75;
          let _623_v214;
          _623_v214 = _dafny.Map.Empty.slice().updateUnsafe(_584_v183,(_this).f29);
          let _624_v215;
          _624_v215 = _dafny.Map.Empty.slice().updateUnsafe(_609_v202,_623_v214);
          let _625_v216;
          _625_v216 = _module.D5.create_DC17((((_624_v215).contains(_612_v205)) ? ((_624_v215).get(_612_v205)) : (_dafny.Map.Empty.slice().updateUnsafe(_584_v183,(_this).f29))));
          let _626_v217;
          _626_v217 = _module.D5.create_DC20(_625_v216);
          let _627_v218;
          _627_v218 = new _dafny.CodePoint('h'.codePointAt(0));
          let _628_v219;
          _628_v219 = _module.D10.create_DC34(_627_v218, (_492_v101)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length))], _611_v204, p0, false);
          let _index75 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length));
          let _rhs50 = _621_v213;
          let _rhs51 = _module.__default.safeModuloInt(_612_v205, _609_v202);
          let _rhs52 = _626_v217;
          let _rhs53 = (_628_v219).dtor_cf57;
          let _lhs41 = globalState;
          let _lhs42 = _492_v101;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length));
          _621_v213 = _rhs50;
          _lhs41.f6 = _rhs51;
          _626_v217 = _rhs52;
          _lhs42[_lhs43] = _rhs53;
          let _index76 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length));
          (_492_v101)[_index76] = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("f"), _584_v183);
          let _629_v220;
          _629_v220 = _dafny.Map.Empty.slice().updateUnsafe(_611_v204,(_492_v101)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length))]);
          (globalState).f10 = (_609_v202).isLessThan(new BigNumber((_629_v220).length));
          (globalState).f8 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber((_620_v212).length), _609_v202), (((_608_cf50).contains(p0)) ? ((_608_cf50).get(p0)) : ((_this).f29)));
        } else {
          _584_v183 = _584_v183;
          _610_v203 = _611_v204;
          (globalState).f0 = !(true) || ((_492_v101)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_492_v101).length))]);
          let _630_v221;
          let _nw76 = new _module.C0();
          _nw76.__ctor();
          _630_v221 = _nw76;
          let _631_v222;
          let _nw77 = Array((new BigNumber(15)).toNumber()).fill(false);
          _631_v222 = _nw77;
          let _632_v223;
          _632_v223 = _dafny.Map.Empty.slice().updateUnsafe(_609_v202,_dafny.Seq.of(_631_v222));
          let _633_v224;
          let _nw78 = new _module.C1();
          _nw78.__ctor((_this).f29, _dafny.Seq.Concat((_this).f22, (((_632_v223).contains(_612_v205)) ? ((_632_v223).get(_612_v205)) : (_dafny.Seq.of(_631_v222, _631_v222, _631_v222, _631_v222)))));
          _633_v224 = _nw78;
        }
      } else {
        let _634___mcc_h14 = (_source9).cf53;
        let _635_cf53 = _634___mcc_h14;
        (globalState).f16 = ((!(p0)) ? ((_this).f29) : (new BigNumber((_dafny.MultiSet.fromElements(!(p0))).cardinality())));
        (globalState).f11 = (_this).f29;
        let _636_v225;
        _636_v225 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
        let _637_v226;
        _637_v226 = _dafny.Map.Empty.slice().updateUnsafe((_636_v225).Merge(_636_v225),p0);
        _637_v226 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_636_v225) : (_636_v225)),(new BigNumber((_584_v183).length)).isLessThanOrEqualTo((_this).f29));
        let _638_v227;
        _638_v227 = _module.D9.create_DC31((_this).f29, _492_v101);
        let _639_v228;
        _639_v228 = _dafny.Map.Empty.slice().updateUnsafe((((_636_v225).contains(p0)) ? ((_636_v225).get(p0)) : (p0)),_638_v227);
        let _640_v229;
        _640_v229 = _dafny.MultiSet.fromElements(new BigNumber((_639_v228).length));
        let _641_v230;
        _641_v230 = _dafny.Map.Empty.slice().updateUnsafe(_492_v101,_640_v229);
        let _642_v231;
        _642_v231 = new _dafny.CodePoint('q'.codePointAt(0));
        let _643_v232;
        _643_v232 = _dafny.Seq.of(_640_v229);
        let _644_v233;
        _644_v233 = _module.D11.create_DC36(_643_v232);
        let _645_v234;
        _645_v234 = _dafny.Seq.of(p0, p0);
        let _rhs54 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat((_644_v233).dtor_cf62, _643_v232), _643_v232)).length);
        let _rhs55 = (_641_v230).update(_492_v101, (_640_v229).update((_this).f29, _module.__default.abs((_this).f29)));
        let _rhs56 = p0;
        let _rhs57 = _module.__default.fm22(p0, (_645_v234)[_module.__default.safeIndex((_this).f29, new BigNumber((_645_v234).length))], new BigNumber(667), (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f29, new BigNumber((_dafny.Seq.of(p0, p0)).length))), globalState);
        let _rhs58 = !(!(p0)) || (p0);
        let _lhs44 = globalState;
        let _lhs45 = globalState;
        let _lhs46 = globalState;
        _lhs44.f6 = _rhs54;
        _641_v230 = _rhs55;
        _lhs45.f0 = _rhs56;
        _642_v231 = _rhs57;
        _lhs46.f10 = _rhs58;
      }
      r0 = p0;
      r1 = ((!(p0)) ? (_492_v101) : (_492_v101));
      let _646_v235;
      let _nw79 = new _module.C1();
      _nw79.__ctor((_this).f29, (_this).f22);
      _646_v235 = _nw79;
      let _647_v236;
      _647_v236 = _dafny.Map.Empty.slice().updateUnsafe(p0,_646_v235);
      let _648_v237;
      let _nw80 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
      _648_v237 = _nw80;
      let _649_v238;
      _649_v238 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_648_v237);
      let _650_v239;
      _650_v239 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f29), new BigNumber((_647_v236).length), new BigNumber((_649_v238).length), (_this).f29);
      let _651_v240;
      _651_v240 = new _dafny.CodePoint('y'.codePointAt(0));
      let _652_v241;
      _652_v241 = _dafny.Map.Empty.slice().updateUnsafe(_650_v239,_651_v240);
      r2 = _652_v241;
      let _653_v242;
      _653_v242 = _dafny.MultiSet.fromElements((_this).f29, (_this).f29, new BigNumber(-818));
      let _654_v243;
      _654_v243 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _655_v244;
      _655_v244 = _dafny.Set.fromElements((_this).f29, (_this).f29);
      let _656_v245;
      _656_v245 = _dafny.Seq.of(_655_v244);
      r3 = _module.__default.safeDivisionInt((_this).f29, (_646_v235).fm0(_653_v242, new BigNumber((_654_v243).length), _656_v245, p0, globalState));
      return [r0, r1, r2, r3];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _657_v0;
      _657_v0 = true;
      let _658_v1;
      _658_v1 = _dafny.MultiSet.fromElements(_657_v0);
      let _659_v2;
      _659_v2 = _dafny.Seq.of(_658_v1);
      r3 = new BigNumber(((_659_v2)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("xrquxsrl")).length), p0), new BigNumber((_659_v2).length))]).cardinality());
      let _660_v3;
      _660_v3 = _dafny.MultiSet.fromElements(p1);
      let _661_v4;
      _661_v4 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(804))).length), p1);
      let _662_v5;
      _662_v5 = _dafny.Seq.of(_661_v4, _dafny.Set.fromElements((_this).f29), _661_v4);
      let _663_v6;
      _663_v6 = _dafny.Seq.of((_this).fm0(_660_v3, (_this).f29, _662_v5, _657_v0, globalState));
      let _664_v7;
      _664_v7 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(p1), _dafny.MultiSet.FromArray(_663_v6));
      let _665_v8;
      let _nw81 = Array((new BigNumber(6)).toNumber()).fill(false);
      _665_v8 = _nw81;
      let _666_v9;
      _666_v9 = _module.D9.create_DC31(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_664_v7).cardinality()),new BigNumber(927))).length), _665_v8);
      let _667_v10;
      _667_v10 = _module.D9.create_DC32(_666_v9);
      _667_v10 = _667_v10;
      let _668_v11;
      _668_v11 = _dafny.Seq.UnicodeFromString("mcwffodmy");
      let _rhs59 = _668_v11;
      let _rhs60 = _660_v3;
      _668_v11 = _rhs59;
      _660_v3 = _rhs60;
      if (_657_v0) {
        let _669_v13;
        _669_v13 = _dafny.Seq.of(_657_v0, true);
        let _670_v14;
        _670_v14 = _dafny.Seq.of(_669_v13, _dafny.Seq.update(_dafny.Seq.of(_657_v0, _657_v0, _657_v0, _657_v0, _657_v0), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.of(_657_v0, _657_v0, _657_v0, _657_v0, _657_v0)).length)), _657_v0), _669_v13, _669_v13);
        let _671_v15;
        _671_v15 = _dafny.Map.Empty.slice().updateUnsafe(_660_v3,_657_v0);
        let _672_v16;
        let _nw82 = Array((new BigNumber(20)).toNumber());
        _nw82[(_dafny.ZERO).toNumber()] = new BigNumber((function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of (_661_v4).Elements) {
            let _673_v12 = _compr_36;
            if ((_661_v4).contains(_673_v12)) {
              _coll36.push([(_673_v12).multipliedBy(p0),(_668_v11)[_module.__default.safeIndex((_this).f29, new BigNumber((_668_v11).length))]]);
            }
          }
          return _coll36;
        }()).length);
        _nw82[(_dafny.ONE).toNumber()] = (new BigNumber(275)).plus(p0);
        _nw82[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_670_v14, _dafny.Seq.of(_669_v13))).length);
        _nw82[(new BigNumber(3)).toNumber()] = (((_658_v1).contains(_657_v0)) ? ((_658_v1).get(_657_v0)) : (new BigNumber((_658_v1).cardinality())));
        _nw82[(new BigNumber(4)).toNumber()] = (p1).plus(p1);
        _nw82[(new BigNumber(5)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("wudmpyivk")).length)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(269)), function (_674_i0) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).length)));
        _nw82[(new BigNumber(6)).toNumber()] = p1;
        _nw82[(new BigNumber(7)).toNumber()] = p1;
        _nw82[(new BigNumber(8)).toNumber()] = p0;
        _nw82[(new BigNumber(9)).toNumber()] = ((_657_v0) ? (p1) : (p0));
        _nw82[(new BigNumber(10)).toNumber()] = (_this).f29;
        _nw82[(new BigNumber(11)).toNumber()] = new BigNumber((_660_v3).cardinality());
        _nw82[(new BigNumber(12)).toNumber()] = p0;
        _nw82[(new BigNumber(13)).toNumber()] = p1;
        _nw82[(new BigNumber(14)).toNumber()] = new BigNumber(983);
        _nw82[(new BigNumber(15)).toNumber()] = p1;
        _nw82[(new BigNumber(16)).toNumber()] = p1;
        _nw82[(new BigNumber(17)).toNumber()] = new BigNumber((_671_v15).length);
        _nw82[(new BigNumber(18)).toNumber()] = (_this).f29;
        _nw82[(new BigNumber(19)).toNumber()] = (new BigNumber(100)).minus(p1);
        _672_v16 = _nw82;
        let _index77 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_672_v16).length));
        (_672_v16)[_index77] = (_this).f29;
        let _index78 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_672_v16).length));
        (_672_v16)[_index78] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(350)), ((_675_v0, _676_v11) => function (_677_i1) {
          return ((!(_675_v0)) ? ((_676_v11)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_676_v11).length))]) : (new _dafny.CodePoint('h'.codePointAt(0))));
        })(_657_v0, _668_v11))).length);
        let _678_v17;
        let _nw83 = new _module.C1();
        _nw83.__ctor((_672_v16)[_module.__default.safeIndex(new BigNumber(898), new BigNumber((_672_v16).length))], (_this).f22);
        _678_v17 = _nw83;
        let _679_v18;
        _679_v18 = _dafny.MultiSet.fromElements(_678_v17);
        _679_v18 = _679_v18;
        let _680_v19;
        _680_v19 = _dafny.MultiSet.fromElements(_665_v8);
        _680_v19 = ((_680_v19).Union(_680_v19)).Intersect(_dafny.MultiSet.fromElements(_665_v8));
        (globalState).f6 = new BigNumber((_dafny.Seq.Concat(((_657_v0) ? (_668_v11) : (_668_v11)), _668_v11)).length);
        _668_v11 = _668_v11;
      } else {
        let _681_v20;
        _681_v20 = _dafny.Seq.of(_657_v0, !(_657_v0), _657_v0);
        let _682_v21;
        _682_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.safeModuloInt(new BigNumber((_661_v4).length), new BigNumber((_681_v20).length)));
        let _index79 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_665_v8).length));
        (_665_v8)[_index79] = _657_v0;
        let _683_v23;
        _683_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_660_v3).cardinality()),_657_v0);
        let _index80 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_665_v8).length));
        let _rhs61 = new BigNumber(499);
        let _rhs62 = function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of _dafny.IntegerRange(new BigNumber(-244), new BigNumber(-210))) {
            let _684_v22 = _compr_37;
            if (((new BigNumber(-244)).isLessThanOrEqualTo(_684_v22)) && ((_684_v22).isLessThan(new BigNumber(-210)))) {
              _coll37.push([_module.__default.safeModuloInt(_684_v22, (_this).f29),new BigNumber((_682_v21).length)]);
            }
          }
          return _coll37;
        }();
        let _rhs63 = ((_657_v0) ? ((_681_v20)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_681_v20).length))]) : ((new BigNumber((_681_v20).length)).isLessThan(new BigNumber((_661_v4).length))));
        let _rhs64 = (((_683_v23).contains((_this).f29)) ? ((_683_v23).get((_this).f29)) : (!(_657_v0)));
        let _lhs47 = globalState;
        let _lhs48 = _665_v8;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_665_v8).length));
        _lhs47.f6 = _rhs61;
        _682_v21 = _rhs62;
        r1 = _rhs63;
        _lhs48[_lhs49] = _rhs64;
        (globalState).f16 = new BigNumber(614);
        if ((((!(false)) ? (_module.__default.fm40((_this).f29, p0, (_668_v11)[_module.__default.safeIndex(p0, new BigNumber((_668_v11).length))], globalState)) : (_683_v23))).contains((_this).f29)) {
          (globalState).f0 = _657_v0;
          let _685_v24;
          _685_v24 = _dafny.MultiSet.fromElements(_663_v6, _dafny.Seq.of((_this).f29));
          let _686_v25;
          _686_v25 = _dafny.Set.fromElements((_665_v8)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_665_v8).length))], _657_v0);
          let _rhs65 = _685_v24;
          let _rhs66 = _658_v1;
          let _rhs67 = _module.__default.fm41(p0, _668_v11, _686_v25, globalState);
          _685_v24 = _rhs65;
          _658_v1 = _rhs66;
          _658_v1 = _rhs67;
          let _687_v26;
          let _nw84 = Array((new BigNumber(4)).toNumber());
          _nw84[(_dafny.ZERO).toNumber()] = p1;
          _nw84[(_dafny.ONE).toNumber()] = p0;
          _nw84[(new BigNumber(2)).toNumber()] = p1;
          _nw84[(new BigNumber(3)).toNumber()] = p0;
          _687_v26 = _nw84;
          let _init9 = ((_688_p1) => function (_689_i2) {
            return (_689_i2).minus(_688_p1);
          })(p1);
          let _nw85 = Array((new BigNumber(16)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw85.length); _i0_9++) {
            _nw85[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _687_v26 = _nw85;
          let _690_v27;
          let _nw86 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
          _690_v27 = _nw86;
          let _rhs68 = _690_v27;
          let _rhs69 = (_665_v8)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_665_v8).length))];
          let _lhs50 = globalState;
          _690_v27 = _rhs68;
          _lhs50.f10 = _rhs69;
          let _index81 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_665_v8).length));
          (_665_v8)[_index81] = _657_v0;
        } else {
          let _691_v28;
          let _init10 = ((_692_v0) => function (_693_i3) {
            return _692_v0;
          })(_657_v0);
          let _nw87 = Array((new BigNumber(23)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw87.length); _i0_10++) {
            _nw87[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _691_v28 = _nw87;
          let _694_v29;
          _694_v29 = _dafny.Map.Empty.slice().updateUnsafe(_691_v28,p1);
          let _695_v30;
          _695_v30 = _dafny.Map.Empty.slice().updateUnsafe((_663_v6)[_module.__default.safeIndex(new BigNumber((_681_v20).length), new BigNumber((_663_v6).length))],_694_v29);
          _694_v29 = (((_695_v30).contains((p1).multipliedBy(new BigNumber(613)))) ? ((_695_v30).get((p1).multipliedBy(new BigNumber(613)))) : (_dafny.Map.Empty.slice().updateUnsafe(_691_v28,p0)));
          (globalState).f16 = ((_this).f29).plus((_this).f29);
          let _696_v31;
          let _nw88 = new _module.C1();
          _nw88.__ctor(_module.__default.safeModuloInt(p0, (_this).f29), _dafny.Seq.of(_691_v28, _691_v28, _691_v28));
          _696_v31 = _nw88;
          let _697_v32;
          let _nw89 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _697_v32 = _nw89;
          _697_v32 = _697_v32;
          let _698_v33;
          let _init11 = ((_699_p1) => function (_700_i4) {
            return (_700_i4).minus(_699_p1);
          })(p1);
          let _nw90 = Array((new BigNumber(3)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw90.length); _i0_11++) {
            _nw90[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _698_v33 = _nw90;
          let _index82 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_698_v33).length));
          (_698_v33)[_index82] = (p0).minus(p0);
          let _index83 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_698_v33).length));
          (_698_v33)[_index83] = new BigNumber(917);
        }
        let _701_v34;
        let _nw91 = new _module.C0();
        _nw91.__ctor();
        _701_v34 = _nw91;
        let _702_v35;
        _702_v35 = _dafny.MultiSet.fromElements(_701_v34);
        let _703_v36;
        _703_v36 = _module.D12.create_DC41(_702_v35);
        let _704_v37;
        let _nw92 = new _module.C1();
        _nw92.__ctor((((_665_v8)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_665_v8).length))]) ? ((_this).fm0(_660_v3, new BigNumber(((_703_v36).dtor_cf68).cardinality()), _662_v5, false, globalState)) : (p1)), (_this).f22);
        _704_v37 = _nw92;
        _657_v0 = (((_683_v23).contains((p0).minus((_this).f29))) ? ((_683_v23).get((p0).minus((_this).f29))) : ((_665_v8)[_module.__default.safeIndex(new BigNumber(591), new BigNumber((_665_v8).length))]));
      }
      let _705_i5;
      _705_i5 = _dafny.ZERO;
      L3: {
        while (!(_module.__default.fm37(_dafny.Seq.of(_657_v0), _dafny.Map.Empty.slice().updateUnsafe(p1,!(_657_v0)), _668_v11, globalState))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_705_i5)) {
              break L3;
            }
            _705_i5 = (_705_i5).plus(_dafny.ONE);
            let _706_v38;
            _706_v38 = _module.D12.create_DC42((_this).f29);
            let _707_v39;
            _707_v39 = _dafny.Seq.of(!(true), _657_v0, _657_v0);
            let _708_v40;
            _708_v40 = _dafny.Set.fromElements(_706_v38, _706_v38, (((_707_v39)[_module.__default.safeIndex(new BigNumber((_668_v11).length), new BigNumber((_707_v39).length))]) ? (_706_v38) : (_module.D12.create_DC42(new BigNumber((_dafny.Seq.UnicodeFromString("agnqcr")).length)))));
            _708_v40 = _708_v40;
            (globalState).f0 = !((_707_v39)[_module.__default.safeIndex(p1, new BigNumber((_707_v39).length))]) || (!(((_657_v0) ? (_657_v0) : (_657_v0))));
            let _709_v41;
            _709_v41 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_710_i6) {
              return new _dafny.CodePoint('m'.codePointAt(0));
            }),p0);
            let _711_v42;
            _711_v42 = _module.D5.create_DC17((_709_v41).Merge(_709_v41));
            let _index84 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_665_v8).length));
            (_665_v8)[_index84] = _657_v0;
            let _712_v43;
            _712_v43 = _dafny.Map.Empty.slice().updateUnsafe(p1,_657_v0);
            let _713_v44;
            _713_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_712_v43).length),_657_v0);
            let _index85 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_665_v8).length));
            let _rhs70 = _module.__default.fm42(globalState);
            let _rhs71 = !(((_dafny.ZERO).minus(p1)).isLessThan(p1)) || (_module.__default.fm37(_dafny.Seq.of(_657_v0), _713_v44, _668_v11, globalState));
            let _rhs72 = _dafny.Seq.Concat(_663_v6, _dafny.Seq.of(new BigNumber(696)));
            let _rhs73 = !(_657_v0);
            let _lhs51 = globalState;
            let _lhs52 = _665_v8;
            let _lhs53 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_665_v8).length));
            _711_v42 = _rhs70;
            _lhs51.f10 = _rhs71;
            _663_v6 = _rhs72;
            _lhs52[_lhs53] = _rhs73;
            (globalState).f11 = (new BigNumber(-412)).minus(new BigNumber(621));
          }
        }
      }
      let _714_v45;
      _714_v45 = _dafny.Seq.of(_657_v0);
      (globalState).f5 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_714_v45, _module.__default.safeIndex((_this).f29, new BigNumber((_714_v45).length)), _657_v0), _714_v45), _714_v45);
      let _715_v46;
      _715_v46 = new _dafny.CodePoint('c'.codePointAt(0));
      let _716_v47;
      _716_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,_715_v46);
      r0 = ((_663_v6)[_module.__default.safeIndex(p1, new BigNumber((_663_v6).length))]).multipliedBy((new BigNumber((_716_v47).length)).minus(p0));
      r1 = _657_v0;
      r2 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(499)), ((_717_v1) => function (_718_i7) {
        return _717_v1;
      })(_658_v1)), (_658_v1).Intersect(_658_v1));
      r3 = ((_657_v0) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))) : ((_this).f29));
      return [r0, r1, r2, r3];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f22 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
    __ctor(f22) {
      let _this = this;
      (_this)._f22 = f22;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.__default.safeModuloInt(new BigNumber(-883), new BigNumber(290))).minus(new BigNumber(-935));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.ZERO;
      let _719_v0;
      _719_v0 = new BigNumber(-487);
      let _720_v1;
      _720_v1 = _dafny.Seq.UnicodeFromString("f");
      let _721_v2;
      _721_v2 = _dafny.Seq.of(_719_v0, new BigNumber((_720_v1).length));
      let _722_v3;
      _722_v3 = _dafny.MultiSet.fromElements(_module.__default.fm14(new BigNumber(-870), _719_v0, globalState), _721_v2);
      let _723_v4;
      _723_v4 = _dafny.Seq.of((_722_v3).IsProperSubsetOf(_722_v3), p0);
      let _724_v5;
      _724_v5 = _dafny.Set.fromElements(_719_v0, _719_v0);
      let _725_v6;
      _725_v6 = _dafny.Seq.of(_724_v5, _724_v5, _724_v5);
      (globalState).f10 = (_723_v4)[_module.__default.safeIndex((_this).fm0((_dafny.MultiSet.fromElements(_719_v0)).update(_719_v0, _module.__default.abs(_719_v0)), _719_v0, _725_v6, p0, globalState), new BigNumber((_723_v4).length))];
      let _726_i0;
      _726_i0 = _dafny.ZERO;
      L4: {
        while (_module.__default.fm15((_719_v0).isLessThanOrEqualTo(_719_v0), globalState)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_726_i0)) {
              break L4;
            }
            _726_i0 = (_726_i0).plus(_dafny.ONE);
            r0 = ((new BigNumber(433)).plus(new BigNumber(-409))).isLessThanOrEqualTo(_719_v0);
            (globalState).f10 = !((_723_v4)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_723_v4).length))]);
            let _727_v7;
            _727_v7 = _dafny.MultiSet.fromElements(new BigNumber(100));
            (globalState).f11 = ((new BigNumber((_727_v7).cardinality())).minus(_719_v0)).plus(_719_v0);
            let _728_v8;
            _728_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
            _728_v8 = _728_v8;
          }
        }
      }
      let _729_v9;
      _729_v9 = _dafny.Set.fromElements(p0);
      (globalState).f0 = (_729_v9).IsProperSubsetOf((_729_v9).Intersect(_729_v9));
      let _730_i1;
      _730_i1 = _dafny.ZERO;
      L5: {
        while (!((_723_v4)[_module.__default.safeIndex(new BigNumber(443), new BigNumber((_723_v4).length))])) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_730_i1)) {
              break L5;
            }
            _730_i1 = (_730_i1).plus(_dafny.ONE);
            let _731_v10;
            _731_v10 = _dafny.Map.Empty.slice().updateUnsafe((_721_v2)[_module.__default.safeIndex(_719_v0, new BigNumber((_721_v2).length))],_723_v4);
            (globalState).f11 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_719_v0,_723_v4)).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(65),_723_v4)).Merge(_731_v10))).length);
            let _732_v11;
            _732_v11 = new _dafny.CodePoint('s'.codePointAt(0));
            let _733_v12;
            _733_v12 = _module.D4.create_DC14(_732_v11, ((_this).f22)[_module.__default.safeIndex(_719_v0, new BigNumber(((_this).f22).length))], _720_v1);
            let _734_v13;
            _734_v13 = _module.D4.create_DC16(_733_v12);
            let _735_v14;
            _735_v14 = _module.D4.create_DC16(_733_v12);
            _735_v14 = _735_v14;
            (globalState).f10 = p0;
            if ((((p0) ? (p0) : (p0))) && ((_723_v4)[_module.__default.safeIndex(new BigNumber(120), new BigNumber((_723_v4).length))])) {
              _720_v1 = _dafny.Seq.UnicodeFromString("hdmeaxjx");
              r3 = new BigNumber(284);
              let _736_v15;
              let _nw93 = Array((new BigNumber(21)).toNumber()).fill(_module.D2.Default());
              _736_v15 = _nw93;
              let _737_v16;
              _737_v16 = _dafny.Seq.of(_736_v15);
              _736_v15 = (_737_v16)[_module.__default.safeIndex((_dafny.ZERO).minus(_719_v0), new BigNumber((_737_v16).length))];
              let _738_v17;
              let _nw94 = Array((new BigNumber(3)).toNumber());
              _nw94[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("swkonbn"), _720_v1);
              _nw94[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("ldxmpoy");
              _nw94[(new BigNumber(2)).toNumber()] = _720_v1;
              _738_v17 = _nw94;
              let _index86 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_738_v17).length));
              (_738_v17)[_index86] = _720_v1;
              let _index87 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_738_v17).length));
              (_738_v17)[_index87] = _720_v1;
              let _739_v18;
              _739_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_723_v4).length),p0);
              _739_v18 = (_739_v18).update(_719_v0, !(p0));
            } else {
              r0 = _dafny.Seq.contains(_dafny.Seq.Concat(_723_v4, _723_v4), p0);
              let _740_v19;
              let _nw95 = new _module.C0();
              _nw95.__ctor();
              _740_v19 = _nw95;
              let _741_v20;
              let _init12 = function (_742_i2) {
                return _module.__default.safeDivisionInt(_742_i2, new BigNumber(144));
              };
              let _nw96 = Array((new BigNumber(26)).toNumber());
              for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw96.length); _i0_12++) {
                _nw96[_i0_12] = _init12(new BigNumber(_i0_12));
              }
              _741_v20 = _nw96;
              let _743_v21;
              let _nw97 = Array((new BigNumber(6)).toNumber()).fill(false);
              _743_v21 = _nw97;
              let _744_v22;
              _744_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              let _index88 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_743_v21).length));
              (_743_v21)[_index88] = (((_744_v22).contains(p0)) ? ((_744_v22).get(p0)) : (false));
              let _745_v23;
              let _nw98 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
              _745_v23 = _nw98;
              let _746_v24;
              _746_v24 = _dafny.MultiSet.fromElements(_720_v1, _720_v1);
              let _index89 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_745_v23).length));
              (_745_v23)[_index89] = _746_v24;
              let _747_v25;
              _747_v25 = _dafny.Map.Empty.slice().updateUnsafe(false,_741_v20);
              let _748_v26;
              _748_v26 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(348)).plus(_719_v0),(((_747_v25).contains(p0)) ? ((_747_v25).get(p0)) : (_741_v20)));
              let _749_v27;
              _749_v27 = _dafny.Map.Empty.slice().updateUnsafe(_719_v0,_719_v0);
              let _750_v28;
              _750_v28 = _dafny.Map.Empty.slice().updateUnsafe(_749_v27,(_724_v5).IsSubsetOf(_724_v5));
              let _index90 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_743_v21).length));
              let _index91 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_745_v23).length));
              let _rhs74 = (((_748_v26).contains(_module.__default.safeModuloInt(_719_v0, new BigNumber((_721_v2).length)))) ? ((_748_v26).get(_module.__default.safeModuloInt(_719_v0, new BigNumber((_721_v2).length)))) : (_741_v20));
              let _rhs75 = (((_750_v28).contains(_dafny.Map.Empty.slice().updateUnsafe(_719_v0,_719_v0))) ? ((_750_v28).get(_dafny.Map.Empty.slice().updateUnsafe(_719_v0,_719_v0))) : (p0));
              let _rhs76 = _746_v24;
              let _lhs54 = _743_v21;
              let _lhs55 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_743_v21).length));
              let _lhs56 = _745_v23;
              let _lhs57 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_745_v23).length));
              _741_v20 = _rhs74;
              _lhs54[_lhs55] = _rhs75;
              _lhs56[_lhs57] = _rhs76;
              let _751_v29;
              _751_v29 = _dafny.MultiSet.fromElements(p0, p0);
              let _752_v30;
              _752_v30 = _dafny.Seq.of(_751_v29);
              let _753_v31;
              _753_v31 = _module.D4.create_DC14(_732_v11, _743_v21, _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(736)), ((_754_v11) => function (_755_i3) {
  return _754_v11;
})(_732_v11)), _module.__default.safeIndex(_719_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(736)), ((_756_v11) => function (_757_i3) {
  return _756_v11;
})(_732_v11))).length)), _732_v11), _module.__default.safeIndex(new BigNumber(688), new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(736)), ((_758_v11) => function (_759_i3) {
  return _758_v11;
})(_732_v11)), _module.__default.safeIndex(_719_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(736)), ((_760_v11) => function (_761_i3) {
  return _760_v11;
})(_732_v11))).length)), _732_v11)).length)), _732_v11));
              let _762_v32;
              _762_v32 = _dafny.Map.Empty.slice().updateUnsafe(_753_v31,(_dafny.ZERO).minus(_719_v0));
              let _index92 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_743_v21).length));
              (_743_v21)[_index92] = (((!(_module.__default.fm15(p0, globalState))) ? (_dafny.MultiSet.fromElements(p0, (_743_v21)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_743_v21).length))], (_743_v21)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_743_v21).length))])) : (_751_v29))).IsDisjointFrom((_752_v30)[_module.__default.safeIndex(new BigNumber((_762_v32).length), new BigNumber((_752_v30).length))]);
              let _763_v33;
              _763_v33 = _dafny.Map.Empty.slice().updateUnsafe(_740_v19,(_743_v21)[_module.__default.safeIndex(new BigNumber(541), new BigNumber((_743_v21).length))]);
              _763_v33 = (_763_v33).Merge(_763_v33);
            }
          }
        }
      }
      let _764_v34;
      _764_v34 = _module.D1.create_DC3(_720_v1);
      let _765_v35;
      let _nw99 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _765_v35 = _nw99;
      let _766_v36;
      _766_v36 = _module.D2.create_DC8(_764_v34, _765_v35, p0);
      if ((((p0) ? (_module.D2.create_DC8(_764_v34, _765_v35, p0)) : (_766_v36))).dtor_cf16) {
        let _index93 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length));
        (_765_v35)[_index93] = _719_v0;
        let _index94 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length));
        (_765_v35)[_index94] = _719_v0;
        let _767_v37;
        _767_v37 = new _dafny.CodePoint('q'.codePointAt(0));
        let _768_v38;
        _768_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_720_v1).length),false);
        let _index95 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length));
        let _rhs77 = _719_v0;
        let _rhs78 = ((p0) ? (_767_v37) : (_767_v37));
        let _rhs79 = ((_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))]).isLessThanOrEqualTo(new BigNumber((_768_v38).length));
        let _rhs80 = _719_v0;
        let _lhs58 = _765_v35;
        let _lhs59 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length));
        let _lhs60 = globalState;
        let _lhs61 = globalState;
        _lhs58[_lhs59] = _rhs77;
        _767_v37 = _rhs78;
        _lhs60.f0 = _rhs79;
        _lhs61.f11 = _rhs80;
        let _769_v39;
        _769_v39 = _module.D2.create_DC6(_module.__default.safeModuloInt(_719_v0, _719_v0));
        _769_v39 = _769_v39;
        if ((_723_v4)[_module.__default.safeIndex(new BigNumber((_721_v2).length), new BigNumber((_723_v4).length))]) {
          let _rhs81 = !(_724_v5).contains((_dafny.ZERO).minus(_719_v0));
          let _rhs82 = _module.__default.fm18(p0, globalState);
          let _rhs83 = _module.__default.fm15((_719_v0).isLessThan((_dafny.ZERO).minus((_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))])), globalState);
          let _rhs84 = p0;
          let _lhs62 = globalState;
          let _lhs63 = globalState;
          _lhs62.f10 = _rhs81;
          _720_v1 = _rhs82;
          _lhs63.f10 = _rhs83;
          r0 = _rhs84;
          (globalState).f16 = (_dafny.ZERO).minus((_719_v0).minus((_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))]));
          let _770_v40;
          let _nw100 = Array((new BigNumber(15)).toNumber());
          _nw100[(_dafny.ZERO).toNumber()] = _767_v37;
          _nw100[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
          _nw100[(new BigNumber(2)).toNumber()] = _767_v37;
          _nw100[(new BigNumber(3)).toNumber()] = _767_v37;
          _nw100[(new BigNumber(4)).toNumber()] = _767_v37;
          _nw100[(new BigNumber(5)).toNumber()] = _767_v37;
          _nw100[(new BigNumber(6)).toNumber()] = ((true) ? (_767_v37) : (_767_v37));
          _nw100[(new BigNumber(7)).toNumber()] = ((p0) ? (_767_v37) : (_767_v37));
          _nw100[(new BigNumber(8)).toNumber()] = _767_v37;
          _nw100[(new BigNumber(9)).toNumber()] = (_720_v1)[_module.__default.safeIndex(_719_v0, new BigNumber((_720_v1).length))];
          _nw100[(new BigNumber(10)).toNumber()] = _767_v37;
          _nw100[(new BigNumber(11)).toNumber()] = _767_v37;
          _nw100[(new BigNumber(12)).toNumber()] = _767_v37;
          _nw100[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
          _nw100[(new BigNumber(14)).toNumber()] = _767_v37;
          _770_v40 = _nw100;
          let _index96 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_770_v40).length));
          (_770_v40)[_index96] = _767_v37;
          let _index97 = _module.__default.safeIndex(new BigNumber(455), new BigNumber((_770_v40).length));
          (_770_v40)[_index97] = _767_v37;
          let _771_v41;
          _771_v41 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p0);
          r3 = (new BigNumber((_771_v41).length)).multipliedBy(_719_v0);
          let _772_v42;
          _772_v42 = _dafny.Map.Empty.slice().updateUnsafe(_720_v1,_dafny.Seq.IsPrefixOf(_723_v4, _723_v4));
          _772_v42 = (_772_v42).update(_720_v1, _dafny.Seq.IsPrefixOf(_723_v4, _723_v4));
        } else {
          _767_v37 = new _dafny.CodePoint('l'.codePointAt(0));
          let _index98 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length));
          (_765_v35)[_index98] = new BigNumber((_module.__default.fm18(p0, globalState)).length);
          let _773_v43;
          _773_v43 = _dafny.Map.Empty.slice().updateUnsafe(_720_v1,(_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))]);
          let _774_v45;
          _774_v45 = _dafny.Set.fromElements(_720_v1);
          let _775_v46;
          _775_v46 = _module.D5.create_DC17(_773_v43);
          let _776_v47;
          let _nw101 = Array((new BigNumber(20)).toNumber());
          _nw101[(_dafny.ZERO).toNumber()] = _773_v43;
          _nw101[(_dafny.ONE).toNumber()] = _773_v43;
          _nw101[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_720_v1,_719_v0);
          _nw101[(new BigNumber(3)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ucufv"), _module.__default.safeIndex((_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))], new BigNumber((_dafny.Seq.UnicodeFromString("ucufv")).length)), _767_v37),_719_v0);
          _nw101[(new BigNumber(5)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(6)).toNumber()] = (((_723_v4)[_module.__default.safeIndex(_719_v0, new BigNumber((_723_v4).length))]) ? ((function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of (_774_v45).Elements) {
              let _777_v44 = _compr_38;
              if ((_774_v45).contains(_777_v44)) {
                _coll38.push([_777_v44,(_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))]]);
              }
            }
            return _coll38;
          }()).update(_720_v1, new BigNumber((_721_v2).length))) : (_773_v43));
          _nw101[(new BigNumber(7)).toNumber()] = _module.__default.fm19(p0, _719_v0, (_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))], true, globalState);
          _nw101[(new BigNumber(8)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(9)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(10)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(11)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(12)).toNumber()] = (_773_v43).Merge(_dafny.Map.Empty.slice().updateUnsafe(_720_v1,new BigNumber((_723_v4).length)));
          _nw101[(new BigNumber(13)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(14)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_720_v1,_719_v0);
          _nw101[(new BigNumber(16)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(17)).toNumber()] = _773_v43;
          _nw101[(new BigNumber(18)).toNumber()] = (_775_v46).dtor_cf29;
          _nw101[(new BigNumber(19)).toNumber()] = ((p0) ? (_773_v43) : (_dafny.Map.Empty.slice().updateUnsafe(_720_v1,_719_v0)));
          _776_v47 = _nw101;
          let _index99 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_776_v47).length));
          (_776_v47)[_index99] = _773_v43;
          let _778_v48;
          _778_v48 = _dafny.MultiSet.fromElements(p0, false);
          let _779_v50;
          _779_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of _dafny.IntegerRange(new BigNumber(733), new BigNumber(-844))) {
              let _780_v49 = _compr_39;
              if (((new BigNumber(733)).isLessThanOrEqualTo(_780_v49)) && ((_780_v49).isLessThan(new BigNumber(-844)))) {
                _coll39.push([(_780_v49).multipliedBy((_721_v2)[_module.__default.safeIndex(_719_v0, new BigNumber((_721_v2).length))]),new BigNumber(859)]);
              }
            }
            return _coll39;
          }());
          let _781_v51;
          _781_v51 = _module.D0.create_DC1((_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))], p0, _779_v50, _719_v0, _767_v37);
          let _index100 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_776_v47).length));
          let _index101 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length));
          let _rhs85 = _module.__default.fm19(((((_778_v48).contains(p0)) ? ((_778_v48).get(p0)) : ((_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))]))).isLessThan((_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))]), (_719_v0).multipliedBy(_719_v0), new BigNumber(829), (_781_v51).dtor_cf1, globalState);
          let _rhs86 = _module.__default.safeDivisionInt((_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))], _module.__default.safeDivisionInt(new BigNumber(557), (_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))]));
          let _rhs87 = new BigNumber(-230);
          let _lhs64 = _776_v47;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_776_v47).length));
          let _lhs66 = _765_v35;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length));
          _lhs64[_lhs65] = _rhs85;
          _lhs66[_lhs67] = _rhs86;
          r3 = _rhs87;
          let _782_v52;
          let _nw102 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _782_v52 = _nw102;
          let _783_v53;
          let _nw103 = Array((new BigNumber(9)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = (_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))];
          _nw103[(_dafny.ONE).toNumber()] = (_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))];
          _nw103[(new BigNumber(2)).toNumber()] = _719_v0;
          _nw103[(new BigNumber(3)).toNumber()] = new BigNumber((_module.__default.fm20(p0, _767_v37, _719_v0, p0, globalState)).length);
          _nw103[(new BigNumber(4)).toNumber()] = _719_v0;
          _nw103[(new BigNumber(5)).toNumber()] = _719_v0;
          _nw103[(new BigNumber(6)).toNumber()] = _719_v0;
          _nw103[(new BigNumber(7)).toNumber()] = (_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))];
          _nw103[(new BigNumber(8)).toNumber()] = (_765_v35)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length))];
          _783_v53 = _nw103;
          let _784_v54;
          _784_v54 = _dafny.Map.Empty.slice().updateUnsafe(_783_v53,p0);
          let _index102 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_782_v52).length));
          (_782_v52)[_index102] = _784_v54;
          let _785_v55;
          let _nw104 = new _module.C0();
          _nw104.__ctor();
          _785_v55 = _nw104;
          let _index103 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_782_v52).length));
          let _rhs88 = _784_v54;
          let _rhs89 = _785_v55;
          let _lhs68 = _782_v52;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_782_v52).length));
          _lhs68[_lhs69] = _rhs88;
          _785_v55 = _rhs89;
          let _786_v56;
          _786_v56 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          (globalState).f0 = (_723_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_786_v56).length), (_dafny.ZERO).minus(_719_v0))), new BigNumber((_723_v4).length))];
        }
        let _index104 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length));
        let _rhs90 = _719_v0;
        let _rhs91 = p0;
        let _lhs70 = _765_v35;
        let _lhs71 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_765_v35).length));
        let _lhs72 = globalState;
        _lhs70[_lhs71] = _rhs90;
        _lhs72.f0 = _rhs91;
      } else {
        (globalState).f10 = _module.__default.fm15(p0, globalState);
        (globalState).f0 = p0;
        let _787_v57;
        let _init13 = ((_788_p0) => function (_789_i4) {
          return _788_p0;
        })(p0);
        let _nw105 = Array((new BigNumber(26)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw105.length); _i0_13++) {
          _nw105[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _787_v57 = _nw105;
        let _index105 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_787_v57).length));
        (_787_v57)[_index105] = false;
        let _790_v58;
        _790_v58 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_719_v0,p0));
        let _791_v59;
        _791_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(205),_module.__default.fm15(false, globalState));
        let _index106 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_787_v57).length));
        let _rhs92 = (_dafny.Set.fromElements(_791_v59)).IsProperSubsetOf(_790_v58);
        let _rhs93 = !((_724_v5).contains(_719_v0)) || (_module.__default.fm15(p0, globalState));
        let _lhs73 = _787_v57;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_787_v57).length));
        r0 = _rhs92;
        _lhs73[_lhs74] = _rhs93;
        let _792_v60;
        let _nw106 = new _module.C0();
        _nw106.__ctor();
        _792_v60 = _nw106;
        let _793_v61;
        _793_v61 = _module.D2.create_DC9(_module.D2.create_DC6(_719_v0));
        let _794_v62;
        _794_v62 = _module.D2.create_DC9(_793_v61);
        let _795_v63;
        _795_v63 = _dafny.Set.fromElements(_794_v62);
        let _796_v64;
        _796_v64 = _dafny.MultiSet.fromElements(_787_v57);
        let _797_v65;
        _797_v65 = _dafny.MultiSet.fromElements(new BigNumber((_796_v64).cardinality()), new BigNumber(715));
        let _798_v66;
        _798_v66 = _dafny.Map.Empty.slice().updateUnsafe(_719_v0,_797_v65);
        let _799_v67;
        _799_v67 = _module.D4.create_DC12(_798_v66);
        let _800_v68;
        _800_v68 = _module.D4.create_DC16(_799_v67);
        let _801_v69;
        _801_v69 = _module.D4.create_DC16((_800_v68).dtor_cf28);
        let _802_v70;
        _802_v70 = _module.D4.create_DC16(_799_v67);
        let _803_v71;
        _803_v71 = _dafny.Seq.of(_800_v68, _800_v68, _800_v68);
        let _804_v72;
        _804_v72 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_794_v62)).Intersect(_795_v63),_803_v71);
        _804_v72 = (_804_v72).update(_795_v63, _803_v71);
      }
      let _805_v73;
      _805_v73 = new _dafny.CodePoint('o'.codePointAt(0));
      let _806_v74;
      let _nw107 = Array((new BigNumber(9)).toNumber()).fill(false);
      _806_v74 = _nw107;
      let _807_v75;
      _807_v75 = _module.D4.create_DC14(_805_v73, _806_v74, _720_v1);
      let _source10 = _807_v75;
      if (_source10.is_DC13) {
        let _808___mcc_h0 = (_source10).cf22;
        let _809_cf22 = _808___mcc_h0;
        if (p0) {
          (globalState).f0 = (_723_v4)[_module.__default.safeIndex(new BigNumber((_729_v9).length), new BigNumber((_723_v4).length))];
          let _810_v76;
          let _init14 = ((_811_v2) => function (_812_i5) {
            return _811_v2;
          })(_721_v2);
          let _nw108 = Array((new BigNumber(16)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw108.length); _i0_14++) {
            _nw108[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _810_v76 = _nw108;
          let _index107 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_810_v76).length));
          (_810_v76)[_index107] = _721_v2;
          let _index108 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_765_v35).length));
          (_765_v35)[_index108] = _719_v0;
          let _index109 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_810_v76).length));
          let _index110 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_765_v35).length));
          let _rhs94 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), ((_813_v0) => function (_814_i6) {
            return (_dafny.ZERO).minus(_813_v0);
          })(_719_v0)), _dafny.Seq.Concat(_721_v2, _721_v2)), _module.__default.safeIndex((_dafny.ZERO).minus(_719_v0), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), ((_815_v0) => function (_816_i6) {
            return (_dafny.ZERO).minus(_815_v0);
          })(_719_v0)), _dafny.Seq.Concat(_721_v2, _721_v2))).length)), _719_v0);
          let _rhs95 = (_719_v0).plus((_719_v0).minus(_719_v0));
          let _lhs75 = _810_v76;
          let _lhs76 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_810_v76).length));
          let _lhs77 = _765_v35;
          let _lhs78 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_765_v35).length));
          _lhs75[_lhs76] = _rhs94;
          _lhs77[_lhs78] = _rhs95;
          let _817_v77;
          let _818_v78;
          let _out46;
          let _out47;
          let _outcollector12 = (_this).m4(globalState);
          _out46 = _outcollector12[0];
          _out47 = _outcollector12[1];
          _817_v77 = _out46;
          _818_v78 = _out47;
          (globalState).f5 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, p0), _723_v4), _723_v4);
          let _index111 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_806_v74).length));
          (_806_v74)[_index111] = !(!(((_765_v35)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_765_v35).length))]).isLessThan(_719_v0)));
          let _index112 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_806_v74).length));
          (_806_v74)[_index112] = _module.__default.fm15(true, globalState);
        } else {
          let _819_v79;
          _819_v79 = _dafny.Map.Empty.slice().updateUnsafe(p0,_806_v74);
          (globalState).f6 = new BigNumber((_819_v79).length);
          _765_v35 = _765_v35;
          _805_v73 = _805_v73;
          let _820_v80;
          let _nw109 = new _module.C0();
          _nw109.__ctor();
          _820_v80 = _nw109;
          let _rhs96 = _765_v35;
          let _rhs97 = _719_v0;
          let _lhs79 = globalState;
          _765_v35 = _rhs96;
          _lhs79.f11 = _rhs97;
        }
        let _index113 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length));
        (_765_v35)[_index113] = _719_v0;
        let _821_v81;
        _821_v81 = _dafny.MultiSet.fromElements(_719_v0);
        let _822_v82;
        _822_v82 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("lgi"),p0);
        let _index114 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length));
        (_765_v35)[_index114] = (_this).fm0(_821_v81, (_dafny.ZERO).minus(new BigNumber((_822_v82).length)), _dafny.Seq.of(function () {
          let _coll40 = new _dafny.Set();
          for (const _compr_40 of _dafny.IntegerRange(new BigNumber(623), new BigNumber(366))) {
            let _823_v83 = _compr_40;
            if (((new BigNumber(623)).isLessThanOrEqualTo(_823_v83)) && ((_823_v83).isLessThan(new BigNumber(366)))) {
              _coll40.add(_module.__default.safeDivisionInt(_823_v83, _719_v0));
            }
          }
          return _coll40;
        }()), p0, globalState);
        let _824_v84;
        _824_v84 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p0);
        let _825_v85;
        _825_v85 = _dafny.Seq.of(_824_v84);
        (globalState).f6 = (_dafny.ZERO).minus(new BigNumber((_825_v85).length));
        if (p0) {
          (globalState).f10 = ((!_dafny.Seq.contains(_720_v1, _805_v73)) ? (p0) : (p0));
          (globalState).f0 = (new BigNumber(-252)).isLessThan(_719_v0);
          let _826_v86;
          let _nw110 = new _module.C1();
          _nw110.__ctor(_719_v0, (_this).f22);
          _826_v86 = _nw110;
          let _827_v87;
          _827_v87 = _dafny.Seq.of(_826_v86);
          let _828_v88;
          _828_v88 = _dafny.Set.fromElements(_827_v87, _dafny.Seq.of(_826_v86));
          let _829_v89;
          _829_v89 = _dafny.Map.Empty.slice().updateUnsafe((_723_v4)[_module.__default.safeIndex((_this).fm0(_dafny.MultiSet.fromElements((_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))], _719_v0, _719_v0), (_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))], _725_v6, p0, globalState), new BigNumber((_723_v4).length))],_828_v88);
          let _rhs98 = (new BigNumber(((_828_v88).Difference((((_829_v89).contains(p0)) ? ((_829_v89).get(p0)) : (_dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.of(_826_v86), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))],(_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))])).length), new BigNumber((_dafny.Seq.of(_826_v86)).length)), _826_v86)))))).length)).plus((_719_v0).multipliedBy((_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))]));
          let _rhs99 = _module.__default.safeDivisionInt(_719_v0, _719_v0);
          let _lhs80 = globalState;
          _lhs80.f11 = _rhs98;
          r3 = _rhs99;
          let _830_v90;
          _830_v90 = _dafny.MultiSet.fromElements(p0);
          let _831_v91;
          _831_v91 = _dafny.Map.Empty.slice().updateUnsafe((_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))],p0);
          let _832_v92;
          _832_v92 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? ((_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))]) : ((((_830_v90).contains((((_831_v91).contains(_719_v0)) ? ((_831_v91).get(_719_v0)) : (p0)))) ? ((_830_v90).get((((_831_v91).contains(_719_v0)) ? ((_831_v91).get(_719_v0)) : (p0)))) : (_719_v0)))),_module.__default.safeDivisionInt((_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))], (_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))]));
          let _rhs100 = _805_v73;
          let _rhs101 = (((_832_v92).contains(new BigNumber((_720_v1).length))) ? ((_832_v92).get(new BigNumber((_720_v1).length))) : (new BigNumber((_dafny.MultiSet.fromElements(_720_v1)).cardinality())));
          let _rhs102 = p0;
          let _rhs103 = p0;
          let _rhs104 = _826_v86;
          let _lhs81 = globalState;
          let _lhs82 = globalState;
          let _lhs83 = globalState;
          _805_v73 = _rhs100;
          _lhs81.f6 = _rhs101;
          _lhs82.f0 = _rhs102;
          _lhs83.f10 = _rhs103;
          _826_v86 = _rhs104;
          let _833_v93;
          _833_v93 = _dafny.MultiSet.fromElements((_766_v36).dtor_cf15);
          let _834_v94;
          let _nw111 = new _module.C1();
          _nw111.__ctor(new BigNumber((_833_v93).cardinality()), (_826_v86).f22);
          _834_v94 = _nw111;
        } else {
          let _835_v95;
          _835_v95 = _module.D1.create_DC2(_720_v1);
          let _836_v96;
          _836_v96 = _dafny.Set.fromElements(_835_v95, _835_v95);
          (globalState).f10 = !((_dafny.Set.fromElements(_835_v95, _835_v95)).IsDisjointFrom(_836_v96));
          let _837_v98;
          _837_v98 = _dafny.Map.Empty.slice().updateUnsafe((_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))],p0);
          let _838_v100;
          _838_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,(function () {
            let _coll41 = new _dafny.Map();
            for (const _compr_41 of (_837_v98).Keys.Elements) {
              let _839_v97 = _compr_41;
              if ((_837_v98).contains(_839_v97)) {
                _coll41.push([_module.__default.safeDivisionInt(_839_v97, (_765_v35)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length))]),false]);
              }
            }
            return _coll41;
          }()).Merge((function () {
            let _coll42 = new _dafny.Map();
            for (const _compr_42 of (_837_v98).Keys.Elements) {
              let _840_v99 = _compr_42;
              if ((_837_v98).contains(_840_v99)) {
                _coll42.push([(_840_v99).minus(_719_v0),p0]);
              }
            }
            return _coll42;
          }()).update(_719_v0, true)));
          let _index115 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_765_v35).length));
          (_765_v35)[_index115] = new BigNumber(((((_838_v100).contains(!((true) && (p0)))) ? ((_838_v100).get(!((true) && (p0)))) : ((_837_v98).Merge(_837_v98)))).length);
          _837_v98 = ((_837_v98).Merge(_837_v98)).Merge(_837_v98);
          let _841_v101;
          let _init15 = ((_842_v1, _843_v95) => function (_844_i7) {
            return _dafny.Seq.Concat(_842_v1, (_843_v95).dtor_cf5);
          })(_720_v1, _835_v95);
          let _nw112 = Array((new BigNumber(20)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw112.length); _i0_15++) {
            _nw112[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _841_v101 = _nw112;
          let _index116 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_841_v101).length));
          (_841_v101)[_index116] = _720_v1;
          let _845_v102;
          _845_v102 = _module.D12.create_DC42(_719_v0);
          let _846_v103;
          _846_v103 = _dafny.Map.Empty.slice().updateUnsafe(p0,new _dafny.CodePoint('f'.codePointAt(0)));
          let _index117 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_841_v101).length));
          let _rhs105 = p0;
          let _rhs106 = (_dafny.ZERO).minus((_845_v102).dtor_cf69);
          let _rhs107 = _dafny.Seq.Concat(_720_v1, _720_v1);
          let _rhs108 = _dafny.Seq.update(_720_v1, _module.__default.safeIndex((_dafny.ZERO).minus(_719_v0), new BigNumber((_720_v1).length)), (((_846_v103).contains(p0)) ? ((_846_v103).get(p0)) : (_805_v73)));
          let _rhs109 = p0;
          let _lhs84 = globalState;
          let _lhs85 = globalState;
          let _lhs86 = _841_v101;
          let _lhs87 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_841_v101).length));
          let _lhs88 = globalState;
          _lhs84.f0 = _rhs105;
          _lhs85.f11 = _rhs106;
          _720_v1 = _rhs107;
          _lhs86[_lhs87] = _rhs108;
          _lhs88.f10 = _rhs109;
          let _847_v104;
          let _nw113 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _847_v104 = _nw113;
          let _848_v105;
          let _nw114 = Array((new BigNumber(21)).toNumber());
          _nw114[(_dafny.ZERO).toNumber()] = _847_v104;
          _nw114[(_dafny.ONE).toNumber()] = _847_v104;
          _nw114[(new BigNumber(2)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(3)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(4)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(5)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(6)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(7)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(8)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(9)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(10)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(11)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(12)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(13)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(14)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(15)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(16)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(17)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(18)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(19)).toNumber()] = _847_v104;
          _nw114[(new BigNumber(20)).toNumber()] = _847_v104;
          _848_v105 = _nw114;
          _848_v105 = _848_v105;
        }
      } else if (_source10.is_DC14) {
        let _849___mcc_h1 = (_source10).cf23;
        let _850___mcc_h2 = (_source10).cf24;
        let _851___mcc_h3 = (_source10).cf25;
        let _852_cf25 = _851___mcc_h3;
        let _853_cf24 = _850___mcc_h2;
        let _854_cf23 = _849___mcc_h1;
        let _855_v106;
        let _856_v107;
        let _out48;
        let _out49;
        let _outcollector13 = (_this).m4(globalState);
        _out48 = _outcollector13[0];
        _out49 = _outcollector13[1];
        _855_v106 = _out48;
        _856_v107 = _out49;
        let _857_v109;
        _857_v109 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll43 = new _dafny.Map();
          for (const _compr_43 of (_720_v1).Elements) {
            let _858_v108 = _compr_43;
            if (_dafny.Seq.contains(_720_v1, _858_v108)) {
              _coll43.push([_858_v108,new BigNumber((_724_v5).length)]);
            }
          }
          return _coll43;
        }(),_856_v107);
        let _859_v110;
        _859_v110 = _dafny.Map.Empty.slice().updateUnsafe(_854_cf23,new BigNumber((_724_v5).length));
        (globalState).f11 = (_dafny.ZERO).minus(new BigNumber((((_857_v109).Merge((_857_v109).update(_859_v110, _719_v0))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_859_v110,_719_v0))).length));
        if (!((_719_v0).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_719_v0, _856_v107)))) {
          (globalState).f0 = p0;
          _720_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_852_cf25, _720_v1), _852_cf25);
          let _860_v111;
          _860_v111 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(960),p0);
          let _861_v112;
          _861_v112 = _dafny.Map.Empty.slice().updateUnsafe(p0,_719_v0);
          let _862_v113;
          _862_v113 = _module.D9.create_DC30(_861_v112);
          let _863_v114;
          _863_v114 = _module.D9.create_DC32(_862_v113);
          let _864_v115;
          _864_v115 = _module.D9.create_DC32((((((_860_v111).contains(_719_v0)) ? ((_860_v111).get(_719_v0)) : (p0))) ? (_862_v113) : (_863_v114)));
          let _pat_let_tv12 = _862_v113;
          _864_v115 = function (_pat_let10_0) {
            return function (_865_dt__update__tmp_h0) {
              return function (_pat_let11_0) {
                return function (_866_dt__update_hcf53_h0) {
                  return _module.D9.create_DC32(_866_dt__update_hcf53_h0);
                }(_pat_let11_0);
              }(_pat_let_tv12);
            }(_pat_let10_0);
          }(_864_v115);
          r0 = false;
          let _867_v116;
          _867_v116 = _dafny.Seq.of(_module.__default.fm29(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), ((_868_v73) => function (_869_i8) {
            return _868_v73;
          })(_805_v73))).length), p0, _719_v0, globalState));
          let _870_v117;
          _870_v117 = _module.D11.create_DC36(_867_v116);
          let _index118 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_806_v74).length));
          (_806_v74)[_index118] = p0;
          let _871_v118;
          _871_v118 = _dafny.MultiSet.fromElements(_856_v107, _856_v107, _856_v107);
          let _872_v119;
          _872_v119 = _dafny.Map.Empty.slice().updateUnsafe(false,_871_v118);
          let _873_v120;
          _873_v120 = _dafny.Map.Empty.slice().updateUnsafe(p0,_852_cf25);
          let _874_v121;
          _874_v121 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-940)), function (_875_i9) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })).length),new BigNumber((_873_v120).length));
          let _876_v122;
          _876_v122 = _dafny.Map.Empty.slice().updateUnsafe(p0,_874_v121);
          let _index119 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_806_v74).length));
          let _rhs110 = (_this).fm0((((_872_v119).contains(p0)) ? ((_872_v119).get(p0)) : (_871_v118)), new BigNumber((_dafny.MultiSet.fromElements(!(p0), p0, p0)).cardinality()), _dafny.Seq.of(_dafny.Set.fromElements(_856_v107), _dafny.Set.fromElements(_719_v0, _856_v107)), (_723_v4)[_module.__default.safeIndex(_719_v0, new BigNumber((_723_v4).length))], globalState);
          let _rhs111 = _module.__default.fm43(new _dafny.CodePoint('v'.codePointAt(0)), _module.D0.create_DC1(_719_v0, p0, _876_v122, _719_v0, _854_cf23), p0, (_dafny.MultiSet.fromElements(p0)).update(p0, _module.__default.abs(_856_v107)), globalState);
          let _rhs112 = p0;
          let _lhs89 = globalState;
          let _lhs90 = _806_v74;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_806_v74).length));
          _lhs89.f6 = _rhs110;
          _870_v117 = _rhs111;
          _lhs90[_lhs91] = _rhs112;
        } else {
          (globalState).f0 = ((_dafny.Seq.IsProperPrefixOf(_723_v4, _dafny.Seq.of(p0))) ? (p0) : (p0));
          let _877_v123;
          _877_v123 = _module.D9.create_DC31(_719_v0, _855_v106);
          let _878_v124;
          _878_v124 = _dafny.Map.Empty.slice().updateUnsafe(p0,_719_v0);
          let _879_v125;
          _879_v125 = _dafny.Map.Empty.slice().updateUnsafe(_719_v0,_878_v124);
          let _880_v126;
          let _nw115 = Array((new BigNumber(8)).toNumber());
          _nw115[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_877_v123).dtor_cf52);
          _nw115[(_dafny.ONE).toNumber()] = (_this).f22;
          _nw115[(new BigNumber(2)).toNumber()] = (_this).f22;
          _nw115[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_853_cf24);
          _nw115[(new BigNumber(4)).toNumber()] = (_this).f22;
          _nw115[(new BigNumber(5)).toNumber()] = (_this).f22;
          _nw115[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((_this).f22, (_this).f22);
          _nw115[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update((_this).f22, _module.__default.safeIndex(new BigNumber((_879_v125).length), new BigNumber(((_this).f22).length)), _853_cf24), (_this).f22);
          _880_v126 = _nw115;
          let _index120 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_880_v126).length));
          (_880_v126)[_index120] = _dafny.Seq.of(_855_v106);
          let _index121 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_880_v126).length));
          (_880_v126)[_index121] = (_this).f22;
          let _881_v127;
          let _nw116 = new _module.C1();
          _nw116.__ctor((_dafny.ZERO).minus(_719_v0), _dafny.Seq.Concat((_this).f22, (_880_v126)[_module.__default.safeIndex(new BigNumber(20), new BigNumber((_880_v126).length))]));
          _881_v127 = _nw116;
          let _882_v128;
          let _nw117 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _882_v128 = _nw117;
          _882_v128 = ((p0) ? (_882_v128) : (_882_v128));
          (globalState).f8 = _module.__default.safeModuloInt(_719_v0, _719_v0);
        }
        let _883_v129;
        let _nw118 = Array((new BigNumber(9)).toNumber()).fill([]);
        _883_v129 = _nw118;
        let _rhs113 = _883_v129;
        let _rhs114 = _856_v107;
        let _lhs92 = globalState;
        _883_v129 = _rhs113;
        _lhs92.f6 = _rhs114;
      } else if (_source10.is_DC15) {
        let _884___mcc_h4 = (_source10).cf26;
        let _885___mcc_h5 = (_source10).cf27;
        let _886_cf27 = _885___mcc_h5;
        let _887_cf26 = _884___mcc_h4;
        r1 = _806_v74;
        (globalState).f16 = new BigNumber((_720_v1).length);
        (globalState).f0 = (_719_v0).isLessThan(_719_v0);
        _805_v73 = (_module.D10.create_DC34(new _dafny.CodePoint('b'.codePointAt(0)), p0, p0, p0, p0)).dtor_cf55;
      } else if (_source10.is_DC12) {
        let _888___mcc_h6 = (_source10).cf21;
        let _889_cf21 = _888___mcc_h6;
        (globalState).f5 = _723_v4;
        let _890_v130;
        _890_v130 = _dafny.MultiSet.fromElements(true, !(p0), p0);
        _890_v130 = _dafny.MultiSet.fromElements((_723_v4)[_module.__default.safeIndex(_719_v0, new BigNumber((_723_v4).length))], p0);
        let _891_v131;
        _891_v131 = _module.D7.create_DC26(_721_v2, _719_v0, _719_v0);
        let _pat_let_tv13 = _721_v2;
        let _source11 = function (_pat_let12_0) {
          return function (_892_dt__update__tmp_h1) {
            return function (_pat_let13_0) {
              return function (_893_dt__update_hcf42_h0) {
                return _module.D7.create_DC26(_893_dt__update_hcf42_h0, (_892_dt__update__tmp_h1).dtor_cf43, (_892_dt__update__tmp_h1).dtor_cf44);
              }(_pat_let13_0);
            }(_pat_let_tv13);
          }(_pat_let12_0);
        }(_891_v131);
        if (_source11.is_DC26) {
          let _894___mcc_h8 = (_source11).cf42;
          let _895___mcc_h9 = (_source11).cf43;
          let _896___mcc_h10 = (_source11).cf44;
          let _897_cf44 = _896___mcc_h10;
          let _898_cf43 = _895___mcc_h9;
          let _899_cf42 = _894___mcc_h8;
          (globalState).f6 = _719_v0;
          let _900_v132;
          let _nw119 = new _module.C0();
          _nw119.__ctor();
          _900_v132 = _nw119;
          let _rhs115 = !(p0);
          let _rhs116 = _dafny.areEqual(_723_v4, _723_v4);
          let _lhs93 = globalState;
          let _lhs94 = globalState;
          _lhs93.f10 = _rhs115;
          _lhs94.f10 = _rhs116;
          (globalState).f10 = _module.__default.fm15(!(p0), globalState);
        } else if (_source11.is_DC25) {
          let _901___mcc_h11 = (_source11).cf41;
          let _902_cf41 = _901___mcc_h11;
          (globalState).f6 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(_719_v0)), _719_v0);
          let _903_v133;
          _903_v133 = _module.D9.create_DC31(_module.__default.safeDivisionInt(_719_v0, _719_v0), _806_v74);
          _903_v133 = _903_v133;
          (globalState).f11 = _719_v0;
          (globalState).f6 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_729_v9).length)), _module.__default.safeModuloInt(_719_v0, _719_v0));
        } else {
          let _904___mcc_h12 = (_source11).cf45;
          let _905_cf45 = _904___mcc_h12;
          let _index122 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_806_v74).length));
          (_806_v74)[_index122] = p0;
          let _906_v134;
          _906_v134 = _dafny.MultiSet.fromElements(_806_v74, _806_v74);
          let _index123 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_806_v74).length));
          (_806_v74)[_index123] = !(((_906_v134).Intersect(_906_v134)).IsSubsetOf((_906_v134).Union(_906_v134)));
          (globalState).f11 = _719_v0;
          let _907_v135;
          let _nw120 = new _module.C2();
          _nw120.__ctor(_719_v0, (_this).f22);
          _907_v135 = _nw120;
          let _908_v136;
          _908_v136 = _dafny.Map.Empty.slice().updateUnsafe((_907_v135).f29,_719_v0);
          let _909_v137;
          _909_v137 = _module.D0.create_DC1((_907_v135).f29, (_806_v74)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_806_v74).length))], _dafny.Map.Empty.slice().updateUnsafe(!((_806_v74)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_806_v74).length))]),_908_v136), _719_v0, _805_v73);
          let _910_v138;
          _910_v138 = _dafny.Set.fromElements(_909_v137);
          let _index124 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_806_v74).length));
          let _rhs117 = (new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_805_v73, _805_v73), _720_v1), _module.__default.safeIndex(_719_v0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_805_v73, _805_v73), _720_v1)).length)), _805_v73)).length)).plus(_719_v0);
          let _rhs118 = (_723_v4)[_module.__default.safeIndex((_907_v135).f29, new BigNumber((_723_v4).length))];
          let _rhs119 = _907_v135;
          let _rhs120 = (_907_v135).fm21(_910_v138, _719_v0, globalState);
          let _rhs121 = p0;
          let _lhs95 = globalState;
          let _lhs96 = globalState;
          let _lhs97 = _806_v74;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_806_v74).length));
          let _lhs99 = globalState;
          _lhs95.f11 = _rhs117;
          _lhs96.f10 = _rhs118;
          _907_v135 = _rhs119;
          _lhs97[_lhs98] = _rhs120;
          _lhs99.f10 = _rhs121;
          let _index125 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_806_v74).length));
          (_806_v74)[_index125] = p0;
        }
        let _911_v139;
        _911_v139 = _dafny.Map.Empty.slice().updateUnsafe(_719_v0,_805_v73);
        let _912_v140;
        _912_v140 = _dafny.Map.Empty.slice().updateUnsafe(_765_v35,_719_v0);
        let _913_v141;
        _913_v141 = _dafny.Map.Empty.slice().updateUnsafe(_912_v140,_911_v139);
        let _914_v142;
        _914_v142 = _dafny.Seq.of(_765_v35);
        let _915_v146;
        _915_v146 = _dafny.Map.Empty.slice().updateUnsafe(false,_911_v139);
        let _916_v147;
        let _nw121 = Array((new BigNumber(26)).toNumber());
        _nw121[(_dafny.ZERO).toNumber()] = _911_v139;
        _nw121[(_dafny.ONE).toNumber()] = _911_v139;
        _nw121[(new BigNumber(2)).toNumber()] = (((_913_v141).contains(_912_v140)) ? ((_913_v141).get(_912_v140)) : (_911_v139));
        _nw121[(new BigNumber(3)).toNumber()] = _module.__default.fm44(p0, !(!(true)), p0, _720_v1, globalState);
        _nw121[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_914_v142).length),new _dafny.CodePoint('i'.codePointAt(0)))).Merge(_911_v139);
        _nw121[(new BigNumber(5)).toNumber()] = _911_v139;
        _nw121[(new BigNumber(6)).toNumber()] = function () {
          let _coll44 = new _dafny.Map();
          for (const _compr_44 of _dafny.IntegerRange(new BigNumber(47), new BigNumber(90))) {
            let _917_v143 = _compr_44;
            if (((new BigNumber(47)).isLessThanOrEqualTo(_917_v143)) && ((_917_v143).isLessThan(new BigNumber(90)))) {
              _coll44.push([(_917_v143).multipliedBy(_719_v0),_805_v73]);
            }
          }
          return _coll44;
        }();
        _nw121[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(175),_805_v73);
        _nw121[(new BigNumber(8)).toNumber()] = _911_v139;
        _nw121[(new BigNumber(9)).toNumber()] = _911_v139;
        _nw121[(new BigNumber(10)).toNumber()] = (function () {
          let _coll45 = new _dafny.Map();
          for (const _compr_45 of _dafny.IntegerRange(new BigNumber(-152), new BigNumber(201))) {
            let _918_v144 = _compr_45;
            if (((new BigNumber(-152)).isLessThanOrEqualTo(_918_v144)) && ((_918_v144).isLessThan(new BigNumber(201)))) {
              _coll45.push([(_918_v144).plus(_719_v0),(((_911_v139).contains(_719_v0)) ? ((_911_v139).get(_719_v0)) : (_805_v73))]);
            }
          }
          return _coll45;
        }()).Merge((_911_v139).update((_dafny.ZERO).minus(_719_v0), _805_v73));
        _nw121[(new BigNumber(11)).toNumber()] = (function () {
          let _coll46 = new _dafny.Map();
          for (const _compr_46 of (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(978)))).cardinality()), _719_v0, new BigNumber(314))).Elements) {
            let _919_v145 = _compr_46;
            if ((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(978)))).cardinality()), _719_v0, new BigNumber(314))).contains(_919_v145)) {
              _coll46.push([_module.__default.safeModuloInt(_919_v145, _719_v0),_805_v73]);
            }
          }
          return _coll46;
        }()).Merge(_911_v139);
        _nw121[(new BigNumber(12)).toNumber()] = _911_v139;
        _nw121[(new BigNumber(13)).toNumber()] = _911_v139;
        _nw121[(new BigNumber(14)).toNumber()] = _911_v139;
        _nw121[(new BigNumber(15)).toNumber()] = (_911_v139).Merge(_911_v139);
        _nw121[(new BigNumber(16)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_719_v0,_805_v73)).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_719_v0,new BigNumber(242))).length), _805_v73);
        _nw121[(new BigNumber(17)).toNumber()] = (_911_v139).update(new BigNumber(780), _805_v73);
        _nw121[(new BigNumber(18)).toNumber()] = (_911_v139).Merge(_911_v139);
        _nw121[(new BigNumber(19)).toNumber()] = _911_v139;
        _nw121[(new BigNumber(20)).toNumber()] = (_911_v139).Merge(_911_v139);
        _nw121[(new BigNumber(21)).toNumber()] = _911_v139;
        _nw121[(new BigNumber(22)).toNumber()] = (((_915_v146).contains(true)) ? ((_915_v146).get(true)) : ((_911_v139).update(_719_v0, _805_v73)));
        _nw121[(new BigNumber(23)).toNumber()] = (_911_v139).Merge(_911_v139);
        _nw121[(new BigNumber(24)).toNumber()] = (_911_v139).update(_719_v0, _805_v73);
        _nw121[(new BigNumber(25)).toNumber()] = _911_v139;
        _916_v147 = _nw121;
        let _index126 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_916_v147).length));
        (_916_v147)[_index126] = _dafny.Map.Empty.slice().updateUnsafe(_719_v0,_805_v73);
        let _index127 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_916_v147).length));
        (_916_v147)[_index127] = (_911_v139).Merge(_911_v139);
      } else {
        let _920___mcc_h7 = (_source10).cf28;
        let _921_cf28 = _920___mcc_h7;
        let _index128 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_765_v35).length));
        (_765_v35)[_index128] = _719_v0;
        let _index129 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_765_v35).length));
        (_765_v35)[_index129] = (_719_v0).plus(((p0) ? (_719_v0) : (_719_v0)));
        (globalState).f0 = p0;
        let _922_v148;
        _922_v148 = _dafny.MultiSet.fromElements(_719_v0, _module.__default.safeDivisionInt(_719_v0, _719_v0));
        _922_v148 = _922_v148;
        let _923_v149;
        let _nw122 = Array((new BigNumber(14)).toNumber()).fill(_module.D12.Default());
        _923_v149 = _nw122;
        let _924_v150;
        let _nw123 = new _module.C0();
        _nw123.__ctor();
        _924_v150 = _nw123;
        let _925_v151;
        _925_v151 = _dafny.MultiSet.fromElements(_924_v150, _924_v150);
        let _926_v152;
        _926_v152 = _module.D12.create_DC41(_925_v151);
        let _index130 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_923_v149).length));
        (_923_v149)[_index130] = _926_v152;
        let _index131 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_765_v35).length));
        let _index132 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_923_v149).length));
        let _index133 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_765_v35).length));
        let _rhs122 = (_765_v35)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_765_v35).length))];
        let _rhs123 = _926_v152;
        let _rhs124 = p0;
        let _rhs125 = p0;
        let _rhs126 = (_this).fm0((_922_v148).Intersect(_922_v148), new BigNumber((_dafny.Seq.Concat(_720_v1, _720_v1)).length), _725_v6, p0, globalState);
        let _lhs100 = _765_v35;
        let _lhs101 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_765_v35).length));
        let _lhs102 = _923_v149;
        let _lhs103 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_923_v149).length));
        let _lhs104 = globalState;
        let _lhs105 = globalState;
        let _lhs106 = _765_v35;
        let _lhs107 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_765_v35).length));
        _lhs100[_lhs101] = _rhs122;
        _lhs102[_lhs103] = _rhs123;
        _lhs104.f0 = _rhs124;
        _lhs105.f0 = _rhs125;
        _lhs106[_lhs107] = _rhs126;
      }
      r0 = !(((!(!(p0))) ? (p0) : (false)));
      let _init16 = ((_927_p0) => function (_928_i10) {
        return (_927_p0) === (_927_p0);
      })(p0);
      let _nw124 = Array((new BigNumber(25)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw124.length); _i0_16++) {
        _nw124[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      r1 = _nw124;
      let _929_v153;
      _929_v153 = _dafny.Map.Empty.slice().updateUnsafe(_719_v0,_719_v0);
      let _930_v154;
      _930_v154 = _module.D5.create_DC19(_929_v153, p0);
      let _pat_let_tv14 = _721_v2;
      let _pat_let_tv15 = _805_v73;
      let _pat_let_tv16 = _721_v2;
      let _pat_let_tv17 = _805_v73;
      let _pat_let_tv18 = _719_v0;
      let _pat_let_tv19 = p0;
      let _pat_let_tv20 = _805_v73;
      let _pat_let_tv21 = _721_v2;
      let _pat_let_tv22 = _805_v73;
      let _pat_let_tv23 = _721_v2;
      let _pat_let_tv24 = _805_v73;
      let _pat_let_tv25 = _721_v2;
      let _pat_let_tv26 = _805_v73;
      let _pat_let_tv27 = _721_v2;
      let _pat_let_tv28 = _805_v73;
      let _pat_let_tv29 = _721_v2;
      let _pat_let_tv30 = _805_v73;
      let _pat_let_tv31 = _721_v2;
      let _pat_let_tv32 = _719_v0;
      let _pat_let_tv33 = _721_v2;
      let _pat_let_tv34 = _719_v0;
      let _pat_let_tv35 = _805_v73;
      r2 = function (_source12) {
        if (_source12.is_DC18) {
          let _931___mcc_h13 = (_source12).cf30;
          let _932___mcc_h14 = (_source12).cf31;
          let _933_cf31 = _932___mcc_h14;
          let _934_cf30 = _931___mcc_h13;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv14,_pat_let_tv15)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv16,_pat_let_tv17)).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), ((_935_v0) => function (_936_i11) {
            return _935_v0;
          })(_pat_let_tv18)), new _dafny.CodePoint('t'.codePointAt(0))));
        } else if (_source12.is_DC19) {
          let _937___mcc_h15 = (_source12).cf32;
          let _938___mcc_h16 = (_source12).cf33;
          let _939_cf33 = _938___mcc_h16;
          let _940_cf32 = _937___mcc_h15;
          if (_pat_let_tv19) {
            return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), function (_941_i12) {
              return new BigNumber((_dafny.Seq.UnicodeFromString("k")).length);
            }),_pat_let_tv20);
          } else {
            return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv21,_pat_let_tv22);
          }
        } else if (_source12.is_DC17) {
          let _942___mcc_h17 = (_source12).cf29;
          let _943_cf29 = _942___mcc_h17;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv23,_pat_let_tv24)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv25,_pat_let_tv26));
        } else {
          let _944___mcc_h18 = (_source12).cf34;
          let _945_cf34 = _944___mcc_h18;
          return ((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv27,_pat_let_tv28)).update(_pat_let_tv29, _pat_let_tv30)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_pat_let_tv31, _module.__default.safeIndex(_pat_let_tv32, new BigNumber((_pat_let_tv33).length)), _pat_let_tv34),_pat_let_tv35));
        }
      }(_930_v154);
      r3 = _719_v0;
      return [r0, r1, r2, r3];
    }
    m4(globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _946_v0;
      _946_v0 = true;
      let _947_v1;
      _947_v1 = _dafny.Set.fromElements(_module.__default.fm15(_946_v0, globalState));
      let _948_v2;
      _948_v2 = new BigNumber(54);
      let _949_v3;
      _949_v3 = _dafny.MultiSet.fromElements(_948_v2, _948_v2, _948_v2);
      let _950_v4;
      _950_v4 = _dafny.MultiSet.fromElements(_949_v3);
      let _951_v5;
      _951_v5 = _dafny.Seq.UnicodeFromString("ftnyfkc");
      let _952_v6;
      _952_v6 = _dafny.Set.fromElements(new BigNumber((_951_v5).length));
      let _953_v7;
      _953_v7 = _dafny.Map.Empty.slice().updateUnsafe(_946_v0,_948_v2);
      let _954_v8;
      _954_v8 = _dafny.Seq.of(_946_v0);
      let _955_v9;
      _955_v9 = _dafny.Seq.of(_952_v6);
      let _956_v10;
      let _init17 = ((_957_v0) => function (_958_i0) {
        return _957_v0;
      })(_946_v0);
      let _nw125 = Array((new BigNumber(7)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw125.length); _i0_17++) {
        _nw125[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _956_v10 = _nw125;
      let _959_v11;
      _959_v11 = _dafny.Map.Empty.slice().updateUnsafe(_946_v0,_956_v10);
      let _960_v12;
      let _nw126 = Array((new BigNumber(12)).toNumber());
      _nw126[(_dafny.ZERO).toNumber()] = new BigNumber(((_dafny.MultiSet.fromElements(_949_v3)).Intersect(_950_v4)).cardinality());
      _nw126[(_dafny.ONE).toNumber()] = new BigNumber((_952_v6).length);
      _nw126[(new BigNumber(2)).toNumber()] = (((_953_v7).contains(_946_v0)) ? ((_953_v7).get(_946_v0)) : (new BigNumber((_954_v8).length)));
      _nw126[(new BigNumber(3)).toNumber()] = ((!(_946_v0)) ? (new BigNumber((_954_v8).length)) : (_948_v2));
      _nw126[(new BigNumber(4)).toNumber()] = (_this).fm0(_949_v3, new BigNumber((_951_v5).length), _955_v9, _946_v0, globalState);
      _nw126[(new BigNumber(5)).toNumber()] = (_this).fm0(_dafny.MultiSet.fromElements(new BigNumber((_954_v8).length)), new BigNumber(-673), _955_v9, false, globalState);
      _nw126[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(_948_v2, _948_v2);
      _nw126[(new BigNumber(7)).toNumber()] = new BigNumber((((_946_v0) ? (_959_v11) : (_959_v11))).length);
      _nw126[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(_948_v2, _948_v2);
      _nw126[(new BigNumber(9)).toNumber()] = new BigNumber((_951_v5).length);
      _nw126[(new BigNumber(10)).toNumber()] = _948_v2;
      _nw126[(new BigNumber(11)).toNumber()] = _948_v2;
      _960_v12 = _nw126;
      let _index134 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
      (_960_v12)[_index134] = (_948_v2).multipliedBy(_948_v2);
      let _961_v13;
      _961_v13 = _dafny.MultiSet.fromElements(_946_v0);
      let _index135 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
      (_956_v10)[_index135] = (_961_v13).IsProperSubsetOf(_961_v13);
      let _962_v14;
      _962_v14 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("dkfdq"), _951_v5),_946_v0);
      let _963_v15;
      _963_v15 = _dafny.Seq.of(_951_v5);
      let _964_v16;
      _964_v16 = new _dafny.CodePoint('i'.codePointAt(0));
      let _965_v17;
      _965_v17 = _module.D8.create_DC29(_946_v0, _946_v0, _964_v16);
      let _966_v18;
      _966_v18 = _module.D13.create_DC44(_954_v8);
      let _967_v19;
      _967_v19 = _dafny.Map.Empty.slice().updateUnsafe(_948_v2,_946_v0);
      let _pat_let_tv36 = _946_v0;
      let _index136 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
      let _index137 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
      let _rhs127 = _dafny.Set.fromElements((((_962_v14).contains(_963_v15)) ? ((_962_v14).get(_963_v15)) : (_946_v0)));
      let _rhs128 = new BigNumber(971);
      let _rhs129 = _dafny.Seq.contains(_954_v8, (_948_v2).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(_965_v17, function (_pat_let14_0) {
        return function (_968_dt__update__tmp_h0) {
          return function (_pat_let15_0) {
            return function (_969_dt__update_hcf47_h0) {
              return _module.D8.create_DC29(_969_dt__update_hcf47_h0, (_968_dt__update__tmp_h0).dtor_cf48, (_968_dt__update__tmp_h0).dtor_cf49);
            }(_pat_let15_0);
          }(_pat_let_tv36);
        }(_pat_let14_0);
      }(_965_v17), _965_v17)).length)));
      let _rhs130 = _dafny.areEqual(_951_v5, _module.__default.fm18(_946_v0, globalState));
      let _rhs131 = _module.__default.fm37((_966_v18).dtor_cf71, (_967_v19).Merge(_module.__default.fm40(_948_v2, (_dafny.ZERO).minus(_948_v2), _module.__default.fm38(true, _948_v2, globalState), globalState)), _951_v5, globalState);
      let _lhs108 = _960_v12;
      let _lhs109 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
      let _lhs110 = _956_v10;
      let _lhs111 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
      let _lhs112 = globalState;
      _947_v1 = _rhs127;
      _lhs108[_lhs109] = _rhs128;
      _lhs110[_lhs111] = _rhs129;
      _lhs112.f0 = _rhs130;
      _946_v0 = _rhs131;
      let _970_v20;
      _970_v20 = _dafny.MultiSet.fromElements(_965_v17);
      let _971_v21;
      _971_v21 = _module.D11.create_DC36(_dafny.Seq.of(_949_v3));
      let _972_v22;
      _972_v22 = _dafny.Seq.of((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
      let _973_v23;
      _973_v23 = _module.D2.create_DC7(_972_v22, (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], _948_v2);
      let _pat_let_tv37 = _946_v0;
      let _pat_let_tv38 = _956_v10;
      let _pat_let_tv39 = _956_v10;
      let _pat_let_tv40 = _946_v0;
      let _pat_let_tv41 = _956_v10;
      let _pat_let_tv42 = _956_v10;
      let _pat_let_tv43 = _946_v0;
      let _pat_let_tv44 = _956_v10;
      let _pat_let_tv45 = _956_v10;
      let _pat_let_tv46 = _956_v10;
      let _pat_let_tv47 = _956_v10;
      let _index138 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
      let _rhs132 = !(function (_source13) {
        if (_source13.is_DC37) {
          let _974___mcc_h0 = (_source13).cf63;
          let _975_cf63 = _974___mcc_h0;
          return (true) || (_pat_let_tv37);
        } else if (_source13.is_DC38) {
          let _976___mcc_h1 = (_source13).cf64;
          let _977___mcc_h2 = (_source13).cf65;
          let _978_cf65 = _977___mcc_h2;
          let _979_cf64 = _976___mcc_h1;
          if ((_pat_let_tv39)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_pat_let_tv38).length))]) {
            return _pat_let_tv40;
          } else {
            return (_pat_let_tv42)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_pat_let_tv41).length))];
          }
        } else if (_source13.is_DC39) {
          let _980___mcc_h3 = (_source13).cf66;
          let _981_cf66 = _980___mcc_h3;
          return _pat_let_tv43;
        } else if (_source13.is_DC36) {
          let _982___mcc_h4 = (_source13).cf62;
          let _983_cf62 = _982___mcc_h4;
          return !(false) || ((_pat_let_tv45)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_pat_let_tv44).length))]);
        } else {
          let _984___mcc_h5 = (_source13).cf67;
          let _985_cf67 = _984___mcc_h5;
          return (_pat_let_tv47)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_pat_let_tv46).length))];
        }
      }(_971_v21));
      let _rhs133 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
      let _rhs134 = _module.__default.fm22((_948_v2).isLessThan((_973_v23).dtor_cf13), (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], _948_v2, _948_v2, globalState);
      let _rhs135 = _970_v20;
      let _lhs113 = _956_v10;
      let _lhs114 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
      _lhs113[_lhs114] = _rhs132;
      _948_v2 = _rhs133;
      _964_v16 = _rhs134;
      _970_v20 = _rhs135;
      let _986_v24;
      _986_v24 = _module.D9.create_DC31(new BigNumber(618), _956_v10);
      let _987_v25;
      _987_v25 = _module.D9.create_DC32(_986_v24);
      let _source14 = _987_v25;
      if (_source14.is_DC31) {
        let _988___mcc_h6 = (_source14).cf51;
        let _989___mcc_h7 = (_source14).cf52;
        let _990_cf52 = _989___mcc_h7;
        let _991_cf51 = _988___mcc_h6;
        let _rhs136 = !(_946_v0);
        let _rhs137 = (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))];
        let _rhs138 = new BigNumber(837);
        let _lhs115 = globalState;
        let _lhs116 = globalState;
        _lhs115.f0 = _rhs136;
        _lhs116.f0 = _rhs137;
        r1 = _rhs138;
        let _index139 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
        (_956_v10)[_index139] = (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))];
        let _992_v26;
        _992_v26 = _dafny.Map.Empty.slice().updateUnsafe(_953_v7,_946_v0);
        _992_v26 = (_992_v26).update(_953_v7, _946_v0);
        let _993_v27;
        _993_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("y"),(_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]);
        if ((((_993_v27).contains(_951_v5)) ? ((_993_v27).get(_951_v5)) : ((((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]) ? (!(!((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]))) : (_946_v0))))) {
          _954_v8 = _954_v8;
          let _994_v28;
          _994_v28 = _dafny.Map.Empty.slice().updateUnsafe((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))],(_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
          let _995_v29;
          _995_v29 = _module.D5.create_DC19(_994_v28, _946_v0);
          let _996_v30;
          _996_v30 = _dafny.Seq.of(_995_v29, _995_v29, _995_v29);
          let _997_v31;
          _997_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_946_v0, (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]),_996_v30);
          let _rhs139 = _960_v12;
          let _rhs140 = _956_v10;
          let _rhs141 = ((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]) || (_946_v0);
          let _rhs142 = (_dafny.Map.Empty.slice().updateUnsafe(_947_v1,_996_v30)).Merge(_997_v31);
          _960_v12 = _rhs139;
          r0 = _rhs140;
          _946_v0 = _rhs141;
          _997_v31 = _rhs142;
          let _index140 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          (_956_v10)[_index140] = _946_v0;
          let _998_v32;
          let _nw127 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _998_v32 = _nw127;
          let _999_v33;
          _999_v33 = _dafny.Map.Empty.slice().updateUnsafe(_946_v0,_951_v5);
          let _index141 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_998_v32).length));
          (_998_v32)[_index141] = (((_999_v33).contains(_946_v0)) ? ((_999_v33).get(_946_v0)) : (_951_v5));
          let _index142 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_998_v32).length));
          (_998_v32)[_index142] = _dafny.Seq.Concat(_dafny.Seq.Concat(_951_v5, _951_v5), _dafny.Seq.update(_951_v5, _module.__default.safeIndex(_948_v2, new BigNumber((_951_v5).length)), _964_v16));
          let _index143 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_998_v32).length));
          (_998_v32)[_index143] = _module.__default.fm18(_946_v0, globalState);
        } else {
          (globalState).f0 = (_946_v0) || (true);
          (globalState).f10 = _946_v0;
          let _1000_v34;
          _1000_v34 = _dafny.Map.Empty.slice().updateUnsafe(_972_v22,_991_cf51);
          let _1001_v35;
          _1001_v35 = _dafny.Map.Empty.slice().updateUnsafe(_946_v0,_967_v19);
          let _1002_v36;
          _1002_v36 = _dafny.Map.Empty.slice().updateUnsafe(_951_v5,(_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
          let _index144 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          let _index145 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          let _rhs143 = (_1001_v35).contains((_dafny.Map.Empty.slice().updateUnsafe(_972_v22,new BigNumber((_dafny.Seq.update(_951_v5, _module.__default.safeIndex((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], new BigNumber((_951_v5).length)), new _dafny.CodePoint('b'.codePointAt(0)))).length))).equals(_1000_v34));
          let _rhs144 = _946_v0;
          let _rhs145 = _module.__default.safeModuloInt((((_1002_v36).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), ((_1005_v16) => function (_1006_i1) {
            return _1005_v16;
          })(_964_v16)))) ? ((_1002_v36).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), ((_1003_v16) => function (_1004_i1) {
            return _1003_v16;
          })(_964_v16)))) : ((_dafny.ZERO).minus(_948_v2))), ((_946_v0) ? (_991_cf51) : (_991_cf51)));
          let _lhs117 = _956_v10;
          let _lhs118 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          let _lhs119 = _960_v12;
          let _lhs120 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          _946_v0 = _rhs143;
          _lhs117[_lhs118] = _rhs144;
          _lhs119[_lhs120] = _rhs145;
          let _1007_v37;
          _1007_v37 = _module.D1.create_DC3(_951_v5);
          let _1008_v38;
          _1008_v38 = _module.D2.create_DC8(_1007_v37, _960_v12, (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]);
          let _1009_v39;
          _1009_v39 = _dafny.Seq.of(_960_v12, _960_v12, _960_v12, _960_v12);
          let _1010_v40;
          let _nw128 = Array((new BigNumber(20)).toNumber());
          _nw128[(_dafny.ZERO).toNumber()] = _960_v12;
          _nw128[(_dafny.ONE).toNumber()] = _960_v12;
          _nw128[(new BigNumber(2)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(3)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(4)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(5)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(6)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(7)).toNumber()] = (_1008_v38).dtor_cf15;
          _nw128[(new BigNumber(8)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(9)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(10)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(11)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(12)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(13)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(14)).toNumber()] = (_1009_v39)[_module.__default.safeIndex((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], new BigNumber((_1009_v39).length))];
          _nw128[(new BigNumber(15)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(16)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(17)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(18)).toNumber()] = _960_v12;
          _nw128[(new BigNumber(19)).toNumber()] = _960_v12;
          _1010_v40 = _nw128;
          let _index146 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1010_v40).length));
          (_1010_v40)[_index146] = (_1008_v38).dtor_cf15;
          let _1011_v41;
          let _nw129 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
          _1011_v41 = _nw129;
          let _1012_v42;
          _1012_v42 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_991_cf51),_960_v12);
          let _index147 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1011_v41).length));
          (_1011_v41)[_index147] = _1012_v42;
          let _1013_v43;
          _1013_v43 = _dafny.Map.Empty.slice().updateUnsafe(_946_v0,_dafny.Map.Empty.slice().updateUnsafe(_991_cf51,(_1009_v39)[_module.__default.safeIndex(_991_cf51, new BigNumber((_1009_v39).length))]));
          let _1014_v44;
          _1014_v44 = _dafny.Map.Empty.slice().updateUnsafe(_946_v0,(_954_v8)[_module.__default.safeIndex(_991_cf51, new BigNumber((_954_v8).length))]);
          let _index148 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1010_v40).length));
          let _index149 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1011_v41).length));
          let _rhs146 = _960_v12;
          let _rhs147 = _946_v0;
          let _rhs148 = (_1012_v42).Merge(((((_1013_v43).contains((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))])) ? ((_1013_v43).get((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))])) : (_dafny.Map.Empty.slice().updateUnsafe(_991_cf51,_960_v12)))).Merge(_1012_v42));
          let _rhs149 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_991_cf51, new BigNumber(47)), new BigNumber((_1014_v44).length))));
          let _lhs121 = _1010_v40;
          let _lhs122 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1010_v40).length));
          let _lhs123 = globalState;
          let _lhs124 = _1011_v41;
          let _lhs125 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_1011_v41).length));
          _lhs121[_lhs122] = _rhs146;
          _lhs123.f0 = _rhs147;
          _lhs124[_lhs125] = _rhs148;
          _991_cf51 = _rhs149;
          (globalState).f8 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
        }
      } else if (_source14.is_DC30) {
        let _1015___mcc_h8 = (_source14).cf50;
        let _1016_cf50 = _1015___mcc_h8;
        r1 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
        (globalState).f0 = (_954_v8)[_module.__default.safeIndex(_948_v2, new BigNumber((_954_v8).length))];
        let _1017_v45;
        _1017_v45 = _dafny.Map.Empty.slice().updateUnsafe(_948_v2,_949_v3);
        let _1018_v46;
        _1018_v46 = _module.D4.create_DC12(_1017_v45);
        _1018_v46 = _1018_v46;
        let _1019_v47;
        let _nw130 = Array((new BigNumber(15)).toNumber()).fill([]);
        _1019_v47 = _nw130;
        let _index150 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_1019_v47).length));
        (_1019_v47)[_index150] = _956_v10;
        let _index151 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_1019_v47).length));
        (_1019_v47)[_index151] = _956_v10;
      } else {
        let _1020___mcc_h9 = (_source14).cf53;
        let _1021_cf53 = _1020___mcc_h9;
        _948_v2 = _948_v2;
        let _index152 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
        let _rhs150 = (new BigNumber((_952_v6).length)).minus(new BigNumber((_972_v22).length));
        let _rhs151 = _946_v0;
        let _lhs126 = globalState;
        let _lhs127 = _956_v10;
        let _lhs128 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
        _lhs126.f16 = _rhs150;
        _lhs127[_lhs128] = _rhs151;
        let _index153 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
        let _rhs152 = _946_v0;
        let _rhs153 = ((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).multipliedBy((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
        let _lhs129 = globalState;
        let _lhs130 = _960_v12;
        let _lhs131 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
        _lhs129.f10 = _rhs152;
        _lhs130[_lhs131] = _rhs153;
        _951_v5 = _dafny.Seq.UnicodeFromString("ia");
      }
      let _1022_v48;
      _1022_v48 = _dafny.Map.Empty.slice().updateUnsafe((_948_v2).multipliedBy(_948_v2),_module.__default.safeModuloInt((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], _948_v2));
      _1022_v48 = _dafny.Map.Empty.slice().updateUnsafe(_948_v2,(_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
      if ((((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).plus(new BigNumber(-544))).isLessThanOrEqualTo(_948_v2)) {
        let _1023_v49;
        let _nw131 = Array((new BigNumber(6)).toNumber());
        _1023_v49 = _nw131;
        let _1024_v50;
        _1024_v50 = _module.D11.create_DC37(_1023_v49);
        let _source15 = _1024_v50;
        if (_source15.is_DC37) {
          let _1025___mcc_h10 = (_source15).cf63;
          let _1026_cf63 = _1025___mcc_h10;
          let _rhs154 = (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))];
          let _rhs155 = _946_v0;
          let _rhs156 = (new BigNumber(212)).multipliedBy(_948_v2);
          let _rhs157 = (((_949_v3).contains((_this).fm0(_949_v3, new BigNumber(467), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-839)), ((_1029_v6) => function (_1030_i2) {
            return _1029_v6;
          })(_952_v6)), !(_946_v0), globalState))) ? ((_949_v3).get((_this).fm0(_949_v3, new BigNumber(467), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-839)), ((_1027_v6) => function (_1028_i2) {
            return _1027_v6;
          })(_952_v6)), !(_946_v0), globalState))) : (new BigNumber(750)));
          let _rhs158 = (_948_v2).minus((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
          let _lhs132 = globalState;
          let _lhs133 = globalState;
          let _lhs134 = globalState;
          let _lhs135 = globalState;
          let _lhs136 = globalState;
          _lhs132.f0 = _rhs154;
          _lhs133.f0 = _rhs155;
          _lhs134.f6 = _rhs156;
          _lhs135.f6 = _rhs157;
          _lhs136.f11 = _rhs158;
          let _index154 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          let _rhs159 = _948_v2;
          let _rhs160 = ((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).isEqualTo(_948_v2);
          let _rhs161 = ((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).multipliedBy((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
          let _lhs137 = _956_v10;
          let _lhs138 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          let _lhs139 = globalState;
          r1 = _rhs159;
          _lhs137[_lhs138] = _rhs160;
          _lhs139.f16 = _rhs161;
          let _index155 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          (_956_v10)[_index155] = _module.__default.fm37(_954_v8, (_967_v19).update(_948_v2, _946_v0), _951_v5, globalState);
          let _index156 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          let _index157 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          let _rhs162 = (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))];
          let _rhs163 = _948_v2;
          let _rhs164 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
          let _lhs140 = globalState;
          let _lhs141 = _960_v12;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          let _lhs143 = _960_v12;
          let _lhs144 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          _lhs140.f10 = _rhs162;
          _lhs141[_lhs142] = _rhs163;
          _lhs143[_lhs144] = _rhs164;
        } else if (_source15.is_DC38) {
          let _1031___mcc_h11 = (_source15).cf64;
          let _1032___mcc_h12 = (_source15).cf65;
          let _1033_cf65 = _1032___mcc_h12;
          let _1034_cf64 = _1031___mcc_h11;
          _967_v19 = function () {
            let _coll47 = new _dafny.Map();
            for (const _compr_47 of (_dafny.Seq.of(new BigNumber(925), _1033_cf65)).Elements) {
              let _1035_v51 = _compr_47;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(925), _1033_cf65), _1035_v51)) {
                _coll47.push([(_1035_v51).plus((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]),(_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]]);
              }
            }
            return _coll47;
          }();
          let _index158 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          (_956_v10)[_index158] = _946_v0;
          _959_v11 = (_959_v11).update((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], _956_v10);
          let _index159 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          let _rhs165 = (_1034_cf64).isLessThan(_948_v2);
          let _rhs166 = _1033_cf65;
          let _rhs167 = (((_1022_v48).contains(_1034_cf64)) ? ((_1022_v48).get(_1034_cf64)) : (((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).minus(_1034_cf64)));
          let _rhs168 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
          let _lhs145 = globalState;
          let _lhs146 = globalState;
          let _lhs147 = _960_v12;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          let _lhs149 = globalState;
          _lhs145.f10 = _rhs165;
          _lhs146.f11 = _rhs166;
          _lhs147[_lhs148] = _rhs167;
          _lhs149.f6 = _rhs168;
        } else if (_source15.is_DC39) {
          let _1036___mcc_h13 = (_source15).cf66;
          let _1037_cf66 = _1036___mcc_h13;
          let _1038_v52;
          _1038_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1037_cf66,_946_v0);
          let _1039_v53;
          _1039_v53 = _dafny.Seq.of(_954_v8);
          let _1040_v54;
          _1040_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1038_v52,!_dafny.areEqual(_dafny.Seq.of(_954_v8), _1039_v53));
          let _1041_v55;
          _1041_v55 = _dafny.Seq.of(_module.D6.create_DC24(_module.D6.create_DC21(_dafny.MultiSet.fromElements((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], new BigNumber((_954_v8).length), (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]))));
          let _1042_v56;
          let _nw132 = new _module.C2();
          _nw132.__ctor(new BigNumber((_951_v5).length), (_this).f22);
          _1042_v56 = _nw132;
          let _1043_v57;
          _1043_v57 = _module.D6.create_DC23(_1037_cf66, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(35),_1042_v56)).length), _946_v0);
          let _1044_v58;
          _1044_v58 = _module.D6.create_DC24(_1043_v57);
          let _rhs169 = _dafny.areEqual(_1041_v55, _dafny.Seq.update(_dafny.Seq.Concat(_1041_v55, _1041_v55), _module.__default.safeIndex((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], new BigNumber((_dafny.Seq.Concat(_1041_v55, _1041_v55)).length)), _1044_v58));
          let _rhs170 = (((((_967_v19).contains(_948_v2)) ? ((_967_v19).get(_948_v2)) : ((_954_v8)[_module.__default.safeIndex(_948_v2, new BigNumber((_954_v8).length))]))) ? ((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]) : (_948_v2));
          let _rhs171 = _972_v22;
          let _rhs172 = function () {
            let _coll48 = new _dafny.Map();
            for (const _compr_48 of (_1040_v54).Keys.Elements) {
              let _1045_v59 = _compr_48;
              if ((_1040_v54).contains(_1045_v59)) {
                _coll48.push([_1045_v59,_946_v0]);
              }
            }
            return _coll48;
          }();
          let _rhs173 = _951_v5;
          let _lhs150 = globalState;
          let _lhs151 = globalState;
          _lhs150.f0 = _rhs169;
          _lhs151.f8 = _rhs170;
          _972_v22 = _rhs171;
          _1040_v54 = _rhs172;
          _951_v5 = _rhs173;
          _1038_v52 = (_1038_v52).update((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], !((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]) || ((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]));
          let _1046_v60;
          let _nw133 = Array((new BigNumber(28)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1046_v60 = _nw133;
          let _1047_v61;
          _1047_v61 = _dafny.Seq.of(_1046_v60, _1046_v60, _1046_v60);
          _1046_v60 = (_1047_v61)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_948_v2, (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]), new BigNumber((_1047_v61).length))];
          let _1048_v62;
          _1048_v62 = _module.D5.create_DC18(_948_v2, (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
          let _1049_v63;
          _1049_v63 = _dafny.Seq.of(_dafny.Seq.of(_1048_v62, _1048_v62, _1048_v62));
          (globalState).f0 = !(_dafny.Seq.IsPrefixOf(_1049_v63, _1049_v63));
        } else if (_source15.is_DC36) {
          let _1050___mcc_h14 = (_source15).cf62;
          let _1051_cf62 = _1050___mcc_h14;
          r1 = ((_946_v0) ? ((_dafny.ZERO).minus(_948_v2)) : ((_this).fm0(_dafny.MultiSet.fromElements((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], _948_v2), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_946_v0, false)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-26)), ((_1052_v6) => function (_1053_i3) {
            return _1052_v6;
          })(_952_v6)), _946_v0, globalState)));
          (globalState).f6 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_948_v2, _948_v2));
          r0 = _956_v10;
          let _index160 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          (_960_v12)[_index160] = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
        } else {
          let _1054___mcc_h15 = (_source15).cf67;
          let _1055_cf67 = _1054___mcc_h15;
          let _index161 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          (_956_v10)[_index161] = (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))];
          (globalState).f0 = _946_v0;
          let _index162 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          (_960_v12)[_index162] = new BigNumber(411);
          let _index163 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          let _rhs174 = (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))];
          let _rhs175 = _948_v2;
          let _rhs176 = (((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]) ? (_947_v1) : (((_946_v0) ? (_dafny.Set.fromElements((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))])) : (_947_v1))));
          let _rhs177 = _dafny.Seq.of(((_946_v0) ? ((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]) : (_948_v2)), (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
          let _rhs178 = (_954_v8)[_module.__default.safeIndex(new BigNumber(-182), new BigNumber((_954_v8).length))];
          let _lhs152 = _956_v10;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          let _lhs154 = globalState;
          let _lhs155 = globalState;
          _lhs152[_lhs153] = _rhs174;
          _lhs154.f16 = _rhs175;
          _947_v1 = _rhs176;
          _972_v22 = _rhs177;
          _lhs155.f0 = _rhs178;
        }
        if ((_948_v2).isLessThanOrEqualTo((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))])) {
          (globalState).f11 = (_948_v2).minus(((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).plus((_this).fm0(_949_v3, new BigNumber((_954_v8).length), _955_v9, _module.__default.fm15(_946_v0, globalState), globalState)));
          let _1056_v64;
          let _nw134 = new _module.C2();
          _nw134.__ctor((_dafny.ZERO).minus((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]), _dafny.Seq.Concat((_this).f22, _dafny.Seq.of(_956_v10)));
          _1056_v64 = _nw134;
          _951_v5 = _951_v5;
          let _1057_v65;
          let _nw135 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1057_v65 = _nw135;
          _1057_v65 = _1057_v65;
          let _1058_v66;
          let _nw136 = new _module.C2();
          _nw136.__ctor(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), ((_1059_v16) => function (_1060_i4) {
            return _1059_v16;
          })(_964_v16))).length), (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]), (_this).f22);
          _1058_v66 = _nw136;
        } else {
          _951_v5 = _951_v5;
          (globalState).f11 = _948_v2;
          (globalState).f16 = new BigNumber(-93);
          (globalState).f0 = (_dafny.Set.fromElements(new BigNumber(606))).IsSubsetOf(_952_v6);
          let _1061_v67;
          _1061_v67 = _dafny.Map.Empty.slice().updateUnsafe(_946_v0,_947_v1);
          _1061_v67 = _1061_v67;
        }
        (globalState).f0 = _946_v0;
        _948_v2 = new BigNumber(484);
        _951_v5 = _dafny.Seq.Concat(_951_v5, (_963_v15)[_module.__default.safeIndex((_dafny.ZERO).minus(_948_v2), new BigNumber((_963_v15).length))]);
      } else {
        if (((((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).isLessThan(new BigNumber((_951_v5).length))) ? (((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).isLessThan(_948_v2)) : ((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]))) {
          (globalState).f6 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
          let _index164 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          (_960_v12)[_index164] = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
          let _index165 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
          (_960_v12)[_index165] = _module.__default.safeDivisionInt((_this).fm0(_949_v3, _948_v2, _955_v9, true, globalState), new BigNumber(511));
          let _1062_v68;
          let _init18 = ((_1063_v16) => function (_1064_i5) {
            return _1063_v16;
          })(_964_v16);
          let _nw137 = Array((new BigNumber(21)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw137.length); _i0_18++) {
            _nw137[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _1062_v68 = _nw137;
          let _1065_v69;
          _1065_v69 = _dafny.MultiSet.fromElements(_1062_v68, _1062_v68);
          let _rhs179 = _946_v0;
          let _rhs180 = _946_v0;
          let _rhs181 = (_dafny.ZERO).minus((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
          let _rhs182 = ((((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]) ? (_1065_v69) : (((_1065_v69).update(_1062_v68, _module.__default.abs(new BigNumber(990)))).update(_1062_v68, _module.__default.abs((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]))))).IsDisjointFrom((_1065_v69).update(_1062_v68, _module.__default.abs((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))])));
          let _lhs156 = globalState;
          let _lhs157 = globalState;
          let _lhs158 = globalState;
          let _lhs159 = globalState;
          _lhs156.f10 = _rhs179;
          _lhs157.f0 = _rhs180;
          _lhs158.f8 = _rhs181;
          _lhs159.f10 = _rhs182;
          _967_v19 = (_967_v19).update(new BigNumber((_949_v3).cardinality()), _946_v0);
        } else {
          let _1066_v70;
          _1066_v70 = _module.D2.create_DC6(_948_v2);
          let _1067_v71;
          _1067_v71 = _dafny.Map.Empty.slice().updateUnsafe((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))],_dafny.MultiSet.fromElements((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], new BigNumber(928)));
          let _1068_v72;
          _1068_v72 = _dafny.Map.Empty.slice().updateUnsafe((_961_v13).update((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], _module.__default.abs((_this).fm0((((_1067_v71).contains(_946_v0)) ? ((_1067_v71).get(_946_v0)) : (_dafny.MultiSet.FromArray(_972_v22))), _948_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(818)), ((_1069_v6) => function (_1070_i6) {
            return _1069_v6;
          })(_952_v6)), (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], globalState))),true);
          let _pat_let_tv48 = _972_v22;
          let _rhs183 = new BigNumber(130);
          let _rhs184 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
          let _rhs185 = new BigNumber((_1068_v72).length);
          let _rhs186 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
          let _rhs187 = function (_pat_let16_0) {
            return function (_1071_dt__update__tmp_h1) {
              return function (_pat_let17_0) {
                return function (_1072_dt__update_hcf9_h0) {
                  return _module.D2.create_DC6(_1072_dt__update_hcf9_h0);
                }(_pat_let17_0);
              }(new BigNumber((_pat_let_tv48).length));
            }(_pat_let16_0);
          }(_module.__default.fm45((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], globalState));
          let _lhs160 = globalState;
          let _lhs161 = globalState;
          let _lhs162 = globalState;
          let _lhs163 = globalState;
          _lhs160.f11 = _rhs183;
          _lhs161.f16 = _rhs184;
          _lhs162.f11 = _rhs185;
          _lhs163.f6 = _rhs186;
          _1066_v70 = _rhs187;
          _961_v13 = (_dafny.MultiSet.fromElements(true)).Intersect(_961_v13);
          let _1073_v73;
          let _nw138 = new _module.C1();
          _nw138.__ctor((new BigNumber((_1022_v48).length)).plus(new BigNumber((_951_v5).length)), (_this).f22);
          _1073_v73 = _nw138;
          let _1074_v74;
          let _nw139 = new _module.C0();
          _nw139.__ctor();
          _1074_v74 = _nw139;
          let _1075_v75;
          let _nw140 = new _module.C0();
          _nw140.__ctor();
          _1075_v75 = _nw140;
        }
        let _1076_v76;
        _1076_v76 = _module.D5.create_DC17(_dafny.Map.Empty.slice().updateUnsafe(_951_v5,(_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]));
        let _source16 = _1076_v76;
        if (_source16.is_DC18) {
          let _1077___mcc_h16 = (_source16).cf30;
          let _1078___mcc_h17 = (_source16).cf31;
          let _1079_cf31 = _1078___mcc_h17;
          let _1080_cf30 = _1077___mcc_h16;
          (globalState).f8 = _1080_cf30;
          let _1081_v77;
          let _nw141 = new _module.C2();
          _nw141.__ctor(_948_v2, (_this).f22);
          _1081_v77 = _nw141;
          (globalState).f16 = _948_v2;
          let _1082_v78;
          _1082_v78 = _dafny.Map.Empty.slice().updateUnsafe(_960_v12,!(true) || (_946_v0));
          _1082_v78 = (_1082_v78).update(_960_v12, ((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).isLessThan(new BigNumber(417)));
        } else if (_source16.is_DC19) {
          let _1083___mcc_h18 = (_source16).cf32;
          let _1084___mcc_h19 = (_source16).cf33;
          let _1085_cf33 = _1084___mcc_h19;
          let _1086_cf32 = _1083___mcc_h18;
          _947_v1 = _947_v1;
          let _1087_v79;
          let _nw142 = new _module.C0();
          _nw142.__ctor();
          _1087_v79 = _nw142;
          (globalState).f10 = _1085_cf33;
          _948_v2 = (new BigNumber((_dafny.Seq.Concat(_951_v5, _951_v5)).length)).minus((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
        } else if (_source16.is_DC17) {
          let _1088___mcc_h20 = (_source16).cf29;
          let _1089_cf29 = _1088___mcc_h20;
          (globalState).f0 = (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))];
          _946_v0 = false;
          (globalState).f0 = _946_v0;
          let _1090_v80;
          _1090_v80 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("xbtvyab"));
          let _index166 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          (_956_v10)[_index166] = (_1090_v80).IsProperSubsetOf(_1090_v80);
        } else {
          let _1091___mcc_h21 = (_source16).cf34;
          let _1092_cf34 = _1091___mcc_h21;
          (globalState).f11 = (_dafny.ZERO).minus((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
          let _1093_v81;
          _1093_v81 = _module.D10.create_DC34(_964_v16, _946_v0, _946_v0, false, _946_v0);
          let _1094_v82;
          _1094_v82 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(164)), ((_1095_v16) => function (_1096_i7) {
            return _1095_v16;
          })(_964_v16)),(_1093_v81).dtor_cf56);
          (globalState).f6 = _module.__default.safeModuloInt(new BigNumber((_1094_v82).length), (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
          _951_v5 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), ((_1097_v16) => function (_1098_i8) {
            return _1097_v16;
          })(_964_v16)), _module.__default.safeIndex(((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).plus((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), ((_1099_v16) => function (_1100_i8) {
            return _1099_v16;
          })(_964_v16))).length)), _964_v16), _module.__default.safeIndex((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), ((_1101_v16) => function (_1102_i8) {
            return _1101_v16;
          })(_964_v16)), _module.__default.safeIndex(((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).plus((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), ((_1103_v16) => function (_1104_i8) {
            return _1103_v16;
          })(_964_v16))).length)), _964_v16)).length)), _964_v16);
          let _index167 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          let _rhs188 = _960_v12;
          let _rhs189 = !(!((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))])) || (_946_v0);
          let _lhs164 = _956_v10;
          let _lhs165 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length));
          _960_v12 = _rhs188;
          _lhs164[_lhs165] = _rhs189;
        }
        let _1105_v83;
        _1105_v83 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_954_v8).length),_953_v7);
        (globalState).f8 = new BigNumber((((_1105_v83).update(new BigNumber((_951_v5).length), _953_v7)).Merge(_1105_v83)).length);
        let _1106_v84;
        _1106_v84 = _module.D6.create_DC23(!((_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]), (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]);
        _1106_v84 = _1106_v84;
        let _1107_v85;
        _1107_v85 = _dafny.Seq.of(_960_v12, _960_v12);
        let _1108_v86;
        _1108_v86 = _module.D8.create_DC28(_dafny.Seq.Concat(_1107_v85, _dafny.Seq.of(_960_v12)));
        let _index168 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
        let _rhs190 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
        let _rhs191 = (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))];
        let _rhs192 = _1108_v86;
        let _rhs193 = ((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).multipliedBy((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]);
        let _lhs166 = globalState;
        let _lhs167 = globalState;
        let _lhs168 = _960_v12;
        let _lhs169 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length));
        _lhs166.f16 = _rhs190;
        _lhs167.f16 = _rhs191;
        _1108_v86 = _rhs192;
        _lhs168[_lhs169] = _rhs193;
      }
      let _1109_i9;
      _1109_i9 = _dafny.ZERO;
      L6: {
        while (_module.__default.fm37(_954_v8, _967_v19, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(964)), function (_1113_i10) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }), _module.__default.safeIndex(new BigNumber(168), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(964)), function (_1114_i10) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length)), _964_v16), globalState)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1109_i9)) {
              break L6;
            }
            _1109_i9 = (_1109_i9).plus(_dafny.ONE);
            let _1110_v87;
            let _nw143 = Array((new BigNumber(2)).toNumber());
            _nw143[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_951_v5, _951_v5);
            _nw143[(_dafny.ONE).toNumber()] = _951_v5;
            _1110_v87 = _nw143;
            _1110_v87 = _1110_v87;
            (globalState).f0 = true;
            (globalState).f10 = (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))];
            let _1111_v88;
            _1111_v88 = _dafny.Map.Empty.slice().updateUnsafe(_946_v0,(_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))]);
            let _1112_v89;
            _1112_v89 = _module.D12.create_DC42(_948_v2);
            _1111_v88 = _dafny.Map.Empty.slice().updateUnsafe(_946_v0,((_this).fm0(_dafny.MultiSet.fromElements((_1112_v89).dtor_cf69), (_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))], _955_v9, (_956_v10)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_956_v10).length))], globalState)).isLessThan((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]));
          }
        }
      }
      r0 = _956_v10;
      r1 = (((_960_v12)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_960_v12).length))]).minus(new BigNumber(65))).minus(new BigNumber((_dafny.Seq.UnicodeFromString("nlauvedc")).length));
      return [r0, r1];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f22 = _dafny.Seq.of();
      this.f28 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
    __ctor(f28, f22) {
      let _this = this;
      (_this).f28 = f28;
      (_this)._f22 = f22;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(955);
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return false;
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.ZERO;
      let _1115_v0;
      _1115_v0 = _dafny.Seq.of(!(p0));
      let _1116_v1;
      _1116_v1 = new BigNumber(508);
      let _1117_v2;
      _1117_v2 = _dafny.MultiSet.fromElements(p0);
      (globalState).f5 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1115_v0, _dafny.Seq.update(_1115_v0, _module.__default.safeIndex(_1116_v1, new BigNumber((_1115_v0).length)), (_this).fm13(_dafny.Seq.Create(_module.__default.abs(new BigNumber(170)), ((_1118_v1) => function (_1119_i0) {
        return _1118_v1;
      })(_1116_v1)), p0, new BigNumber((_1117_v2).cardinality()), globalState))), ((_this.f28) ? (_1115_v0) : (_dafny.Seq.of(_this.f28))));
      let _1120_v3;
      let _init19 = function (_1121_i2) {
        return !(false);
      };
      let _nw144 = Array((new BigNumber(21)).toNumber());
      for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw144.length); _i0_19++) {
        _nw144[_i0_19] = _init19(new BigNumber(_i0_19));
      }
      _1120_v3 = _nw144;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1120_v3).length))) {
        let _1122_i1 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1122_i1)) && ((_1122_i1).isLessThan(new BigNumber((_1120_v3).length))))) {
          (_1120_v3)[(_1122_i1)] = p0;
        }
      }
      _1117_v2 = _1117_v2;
      let _1123_v4;
      let _nw145 = new _module.C1();
      _nw145.__ctor((new BigNumber((_1115_v0).length)).minus((_dafny.ZERO).minus(_1116_v1)), (_this).f22);
      _1123_v4 = _nw145;
      let _1124_v5;
      _1124_v5 = _dafny.Seq.UnicodeFromString("kqfjidb");
      _1124_v5 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-549)), function (_1125_i3) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _1124_v5);
      _1120_v3 = ((p0) ? (_1120_v3) : (_1120_v3));
      r0 = _this.f28;
      r1 = _1120_v3;
      let _1126_v6;
      _1126_v6 = _dafny.Seq.of(_1116_v1, _1116_v1, (_1123_v4).f30);
      let _1127_v7;
      _1127_v7 = _dafny.Seq.of((_1123_v4).f30, new BigNumber((_1126_v6).length), (_1123_v4).f30, (_1123_v4).f30, _1116_v1);
      let _1128_v8;
      _1128_v8 = new _dafny.CodePoint('l'.codePointAt(0));
      let _1129_v9;
      _1129_v9 = _module.D10.create_DC34(_1128_v8, _this.f28, p0, _this.f28, p0);
      let _1130_v10;
      _1130_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of((_1123_v4).f30, _1116_v1), _1126_v6),(_1129_v9).dtor_cf55);
      r2 = _1130_v10;
      r3 = (_1123_v4).fm35((_1123_v4).f30, (_1123_v4).f30, p0, (_1123_v4).f30, globalState);
      return [r0, r1, r2, r3];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f26 = _dafny.Map.Empty;
      this._f27 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f26, f27) {
      let _this = this;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(true, true), _dafny.Seq.of(!(false)));
    };
    fm5(p0, globalState) {
      let _this = this;
      return !((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(511)))).contains((_this).f27)) || (!((_this).f27).isEqualTo(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements((_this).f27)).length), (_this).f27, (_this).f27, (_this).f27, new BigNumber(((_this).f26).length))).length)));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let _1131_v0;
      _1131_v0 = _module.D0.create_DC0();
      let _source17 = _1131_v0;
      if (_source17.is_DC0) {
        let _1132_v1;
        _1132_v1 = true;
        if ((_this).fm5(!(!(true) || (_1132_v1)), globalState)) {
          let _1133_v2;
          _1133_v2 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1134_v3;
          _1134_v3 = _dafny.Seq.of(_1132_v1, _1132_v1, _1132_v1);
          let _1135_v4;
          _1135_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1133_v2,_1134_v3);
          let _rhs194 = ((_1134_v3)[_module.__default.safeIndex(p0, new BigNumber((_1134_v3).length))]) || (_1132_v1);
          let _rhs195 = (_dafny.Map.Empty.slice().updateUnsafe(_1133_v2,_1134_v3)).Merge(_1135_v4);
          let _rhs196 = p2;
          let _lhs170 = globalState;
          let _lhs171 = globalState;
          _lhs170.f0 = _rhs194;
          _1135_v4 = _rhs195;
          _lhs171.f8 = _rhs196;
          let _1136_v5;
          let _nw146 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1136_v5 = _nw146;
          let _1137_v6;
          _1137_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1136_v5,new _dafny.CodePoint('q'.codePointAt(0)));
          let _1138_v7;
          let _nw147 = Array((new BigNumber(17)).toNumber());
          _nw147[(_dafny.ZERO).toNumber()] = _1137_v6;
          _nw147[(_dafny.ONE).toNumber()] = _1137_v6;
          _nw147[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1136_v5,_1133_v2);
          _nw147[(new BigNumber(3)).toNumber()] = _1137_v6;
          _nw147[(new BigNumber(4)).toNumber()] = (_1137_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1136_v5,_1133_v2));
          _nw147[(new BigNumber(5)).toNumber()] = (_1137_v6).Merge(_1137_v6);
          _nw147[(new BigNumber(6)).toNumber()] = (_1137_v6).Merge(_1137_v6);
          _nw147[(new BigNumber(7)).toNumber()] = _1137_v6;
          _nw147[(new BigNumber(8)).toNumber()] = _1137_v6;
          _nw147[(new BigNumber(9)).toNumber()] = _1137_v6;
          _nw147[(new BigNumber(10)).toNumber()] = _1137_v6;
          _nw147[(new BigNumber(11)).toNumber()] = (_1137_v6).update(_1136_v5, _1133_v2);
          _nw147[(new BigNumber(12)).toNumber()] = _1137_v6;
          _nw147[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1136_v5,_1133_v2);
          _nw147[(new BigNumber(14)).toNumber()] = (_1137_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1136_v5,_1133_v2));
          _nw147[(new BigNumber(15)).toNumber()] = (_1137_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1136_v5,_1133_v2));
          _nw147[(new BigNumber(16)).toNumber()] = _1137_v6;
          _1138_v7 = _nw147;
          let _index169 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_1138_v7).length));
          (_1138_v7)[_index169] = _1137_v6;
          let _1139_v8;
          _1139_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber(636));
          let _1140_v9;
          _1140_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1132_v1,_1139_v8);
          let _1141_v10;
          _1141_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.of(new BigNumber(988)));
          let _1142_v11;
          _1142_v11 = _module.D0.create_DC1((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(496)), ((_1143_v2) => function (_1144_i0) {
  return _1143_v2;
})(_1133_v2))).length), p2)), ((_1132_v1) ? (_1132_v1) : (_1132_v1)), (_1140_v9).Merge(_1140_v9), new BigNumber((_1141_v10).length), _1133_v2);
          let _1145_v12;
          let _nw148 = Array((new BigNumber(6)).toNumber()).fill([]);
          _1145_v12 = _nw148;
          let _1146_v13;
          let _nw149 = Array((new BigNumber(7)).toNumber()).fill([]);
          _1146_v13 = _nw149;
          let _index170 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_1145_v12).length));
          (_1145_v12)[_index170] = _1146_v13;
          let _pat_let_tv49 = p2;
          let _pat_let_tv50 = _1132_v1;
          let _pat_let_tv51 = p0;
          let _pat_let_tv52 = _1132_v1;
          let _pat_let_tv53 = _1139_v8;
          let _index171 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_1138_v7).length));
          let _index172 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_1145_v12).length));
          let _rhs197 = !(_module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Seq.UnicodeFromString("e")).length))).isEqualTo(p2);
          let _rhs198 = _1137_v6;
          let _rhs199 = function (_pat_let18_0) {
            return function (_1147_dt__update__tmp_h0) {
              return function (_pat_let19_0) {
                return function (_1148_dt__update_hcf3_h0) {
                  return function (_pat_let20_0) {
                    return function (_1149_dt__update_hcf0_h0) {
                      return function (_pat_let21_0) {
                        return function (_1150_dt__update_hcf2_h0) {
                          return _module.D0.create_DC1(_1149_dt__update_hcf0_h0, (_1147_dt__update__tmp_h0).dtor_cf1, _1150_dt__update_hcf2_h0, _1148_dt__update_hcf3_h0, (_1147_dt__update__tmp_h0).dtor_cf4);
                        }(_pat_let21_0);
                      }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv52,_pat_let_tv53));
                    }(_pat_let20_0);
                  }(_pat_let_tv51);
                }(_pat_let19_0);
              }(((_pat_let_tv50) ? (_pat_let_tv49) : (new BigNumber(-302))));
            }(_pat_let18_0);
          }(_1142_v11);
          let _rhs200 = _1146_v13;
          let _lhs172 = globalState;
          let _lhs173 = _1138_v7;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_1138_v7).length));
          let _lhs175 = _1145_v12;
          let _lhs176 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_1145_v12).length));
          _lhs172.f10 = _rhs197;
          _lhs173[_lhs174] = _rhs198;
          _1142_v11 = _rhs199;
          _lhs175[_lhs176] = _rhs200;
          let _1151_v14;
          _1151_v14 = _dafny.Seq.UnicodeFromString("hcvvsgggf");
          let _1152_v15;
          _1152_v15 = _dafny.MultiSet.fromElements((_this).f27, p0);
          let _rhs201 = _1151_v14;
          let _rhs202 = (_module.__default.safeModuloInt(new BigNumber((_1152_v15).cardinality()), p0)).multipliedBy(p2);
          let _rhs203 = _1132_v1;
          let _lhs177 = globalState;
          let _lhs178 = globalState;
          _1151_v14 = _rhs201;
          _lhs177.f8 = _rhs202;
          _lhs178.f0 = _rhs203;
          let _1153_v16;
          let _1154_v17;
          let _out50;
          let _out51;
          let _outcollector14 = (_this).m3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(882)), ((_1155_v2) => function (_1156_i1) {
            return _1155_v2;
          })(_1133_v2)), !(!(_1132_v1) || (_1132_v1)), globalState);
          _out50 = _outcollector14[0];
          _out51 = _outcollector14[1];
          _1153_v16 = _out50;
          _1154_v17 = _out51;
          let _1157_v18;
          _1157_v18 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_1132_v1)) || (true),_1136_v5);
          _1157_v18 = (_1157_v18).update(_1132_v1, _1136_v5);
        } else {
          let _1158_v19;
          _1158_v19 = _dafny.Seq.UnicodeFromString("w");
          _1158_v19 = ((false) ? (_dafny.Seq.UnicodeFromString("tggp")) : (_dafny.Seq.Concat(p1, p1)));
          _1132_v1 = _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hkfl"), _1158_v19), _module.__default.fm6(globalState));
          let _1159_v20;
          _1159_v20 = new _dafny.CodePoint('e'.codePointAt(0));
          let _1160_v21;
          _1160_v21 = _dafny.Set.fromElements(p1, _dafny.Seq.update(_1158_v19, _module.__default.safeIndex(p2, new BigNumber((_1158_v19).length)), _1159_v20));
          let _1161_v22;
          _1161_v22 = _dafny.Seq.of(_1158_v19);
          let _rhs204 = (function () {
            let _coll49 = new _dafny.Set();
            for (const _compr_49 of (_1161_v22).Elements) {
              let _1162_v23 = _compr_49;
              if (_dafny.Seq.contains(_1161_v22, _1162_v23)) {
                _coll49.add(_1162_v23);
              }
            }
            return _coll49;
          }()).Union(_1160_v21);
          let _rhs205 = ((_this).f27).multipliedBy(new BigNumber(893));
          let _rhs206 = _1132_v1;
          let _rhs207 = ((_this).f27).minus((_dafny.ZERO).minus(p0));
          let _lhs179 = globalState;
          let _lhs180 = globalState;
          _1160_v21 = _rhs204;
          _lhs179.f11 = _rhs205;
          _1132_v1 = _rhs206;
          _lhs180.f11 = _rhs207;
          let _1163_v24;
          _1163_v24 = _dafny.Seq.of(_1132_v1, true);
          let _1164_v25;
          _1164_v25 = _dafny.MultiSet.fromElements(_1132_v1, _1132_v1);
          let _1165_v26;
          _1165_v26 = _dafny.Map.Empty.slice().updateUnsafe((((_1164_v25).contains(_1132_v1)) ? ((_1164_v25).get(_1132_v1)) : (new BigNumber((_1158_v19).length))),p0);
          let _1166_v27;
          _1166_v27 = _dafny.Map.Empty.slice().updateUnsafe(!((_1163_v24)[_module.__default.safeIndex(p0, new BigNumber((_1163_v24).length))]),_1165_v26);
          let _1167_v28;
          _1167_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1158_v19,_1132_v1);
          let _1168_v29;
          _1168_v29 = _module.D0.create_DC1(new BigNumber(((_this).f26).length), false, _1166_v27, new BigNumber((_1167_v28).length), _1159_v20);
          (globalState).f0 = (((_1168_v29).dtor_cf1) ? (!((_this).fm5(_1132_v1, globalState))) : (_1132_v1));
          _1132_v1 = !((_1164_v25).IsSubsetOf(_1164_v25));
        }
        if (_1132_v1) {
          _1131_v0 = _1131_v0;
          let _1169_v30;
          let _nw150 = Array((new BigNumber(20)).toNumber()).fill([]);
          _1169_v30 = _nw150;
          let _1170_v31;
          let _init20 = ((_1171_v1) => function (_1172_i2) {
            return _1171_v1;
          })(_1132_v1);
          let _nw151 = Array((new BigNumber(7)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw151.length); _i0_20++) {
            _nw151[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _1170_v31 = _nw151;
          let _index173 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_1169_v30).length));
          (_1169_v30)[_index173] = _1170_v31;
          let _index174 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_1169_v30).length));
          let _nw152 = Array((new BigNumber(7)).toNumber());
          _nw152[(_dafny.ZERO).toNumber()] = _1132_v1;
          _nw152[(_dafny.ONE).toNumber()] = false;
          _nw152[(new BigNumber(2)).toNumber()] = (_module.__default.fm7(p0, globalState)).isEqualTo((_this).f27);
          _nw152[(new BigNumber(3)).toNumber()] = _1132_v1;
          _nw152[(new BigNumber(4)).toNumber()] = _1132_v1;
          _nw152[(new BigNumber(5)).toNumber()] = _1132_v1;
          _nw152[(new BigNumber(6)).toNumber()] = _1132_v1;
          (_1169_v30)[_index174] = _nw152;
          let _1173_v32;
          _1173_v32 = _dafny.Seq.of((_this).f26, (_this).f26, (_this).f26);
          let _1174_v35;
          _1174_v35 = _dafny.Seq.of(new BigNumber(738));
          let _1175_v36;
          let _nw153 = Array((new BigNumber(20)).toNumber());
          _nw153[(_dafny.ZERO).toNumber()] = (_this).f26;
          _nw153[(_dafny.ONE).toNumber()] = (_this).f26;
          _nw153[(new BigNumber(2)).toNumber()] = (_1173_v32)[_module.__default.safeIndex(p2, new BigNumber((_1173_v32).length))];
          _nw153[(new BigNumber(3)).toNumber()] = (_this).f26;
          _nw153[(new BigNumber(4)).toNumber()] = (((_this).f26).update(p0, false)).Merge((_this).f26);
          _nw153[(new BigNumber(5)).toNumber()] = (_this).f26;
          _nw153[(new BigNumber(6)).toNumber()] = (_this).f26;
          _nw153[(new BigNumber(7)).toNumber()] = ((_this).f26).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_1132_v1, !(_1132_v1))).cardinality()),_1132_v1));
          _nw153[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1132_v1);
          _nw153[(new BigNumber(9)).toNumber()] = (_this).f26;
          _nw153[(new BigNumber(10)).toNumber()] = (_this).f26;
          _nw153[(new BigNumber(11)).toNumber()] = ((_1132_v1) ? ((_1173_v32)[_module.__default.safeIndex(p0, new BigNumber((_1173_v32).length))]) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1132_v1)));
          _nw153[(new BigNumber(12)).toNumber()] = ((_this).f26).Merge((_this).f26);
          _nw153[(new BigNumber(13)).toNumber()] = (_this).f26;
          _nw153[(new BigNumber(14)).toNumber()] = (_this).f26;
          _nw153[(new BigNumber(15)).toNumber()] = function () {
            let _coll50 = new _dafny.Map();
            for (const _compr_50 of _dafny.IntegerRange(new BigNumber(243), new BigNumber(987))) {
              let _1176_v33 = _compr_50;
              if (((new BigNumber(243)).isLessThanOrEqualTo(_1176_v33)) && ((_1176_v33).isLessThan(new BigNumber(987)))) {
                _coll50.push([(_1176_v33).minus(new BigNumber(-492)),_1132_v1]);
              }
            }
            return _coll50;
          }();
          _nw153[(new BigNumber(16)).toNumber()] = (_this).f26;
          _nw153[(new BigNumber(17)).toNumber()] = (function () {
            let _coll51 = new _dafny.Map();
            for (const _compr_51 of (_1174_v35).Elements) {
              let _1177_v34 = _compr_51;
              if (_dafny.Seq.contains(_1174_v35, _1177_v34)) {
                _coll51.push([(_1177_v34).plus((_this).f27),_1132_v1]);
              }
            }
            return _coll51;
          }()).Merge(_module.__default.fm8(p1, _1132_v1, p2, globalState));
          _nw153[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(653),_1132_v1);
          _nw153[(new BigNumber(19)).toNumber()] = (_this).f26;
          _1175_v36 = _nw153;
          _1175_v36 = _1175_v36;
          let _1178_v37;
          _1178_v37 = _dafny.Set.fromElements(_1132_v1);
          let _1179_v38;
          _1179_v38 = _dafny.MultiSet.fromElements(_1178_v37);
          (globalState).f0 = (_dafny.areEqual(_1174_v35, _1174_v35)) === ((_1179_v38).IsProperSubsetOf(_1179_v38));
          let _1180_v39;
          let _nw154 = Array((new BigNumber(2)).toNumber()).fill([]);
          _1180_v39 = _nw154;
          let _1181_v40;
          _1181_v40 = _dafny.MultiSet.fromElements(new BigNumber((p1).length));
          let _1182_v41;
          let _nw155 = Array((new BigNumber(8)).toNumber());
          _nw155[(_dafny.ZERO).toNumber()] = new BigNumber(740);
          _nw155[(_dafny.ONE).toNumber()] = p0;
          _nw155[(new BigNumber(2)).toNumber()] = p2;
          _nw155[(new BigNumber(3)).toNumber()] = (_this).f27;
          _nw155[(new BigNumber(4)).toNumber()] = p0;
          _nw155[(new BigNumber(5)).toNumber()] = new BigNumber(23);
          _nw155[(new BigNumber(6)).toNumber()] = new BigNumber((p1).length);
          _nw155[(new BigNumber(7)).toNumber()] = new BigNumber((_1181_v40).cardinality());
          _1182_v41 = _nw155;
          let _index175 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((_1180_v39).length));
          (_1180_v39)[_index175] = _1182_v41;
          let _index176 = _module.__default.safeIndex(new BigNumber(853), new BigNumber((_1180_v39).length));
          (_1180_v39)[_index176] = _1182_v41;
        } else {
          let _1183_v42;
          let _init21 = ((_1184_p2) => function (_1185_i3) {
            return (_1185_i3).plus(_1184_p2);
          })(p2);
          let _nw156 = Array((_dafny.ONE).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw156.length); _i0_21++) {
            _nw156[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _1183_v42 = _nw156;
          _1183_v42 = _1183_v42;
          let _1186_v43;
          _1186_v43 = _dafny.Set.fromElements(_1132_v1);
          let _1187_v44;
          _1187_v44 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(-518));
          let _1188_v45;
          _1188_v45 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1189_v46;
          _1189_v46 = _module.D0.create_DC1((new BigNumber((_1186_v43).length)).plus((_this).f27), _1132_v1, _dafny.Map.Empty.slice().updateUnsafe(!(_1132_v1),_1187_v44), (_this).f27, ((_1132_v1) ? (_1188_v45) : (_1188_v45)));
          let _1190_v47;
          _1190_v47 = _dafny.Seq.of(_module.__default.fm7((_this).f27, globalState), new BigNumber(-995), new BigNumber((p1).length));
          let _1191_v48;
          _1191_v48 = _dafny.MultiSet.fromElements(p0, new BigNumber((_dafny.MultiSet.fromElements((_this).f26)).cardinality()), p2, new BigNumber((_1190_v47).length), (_this).f27);
          let _rhs208 = _1189_v46;
          let _rhs209 = (_dafny.MultiSet.fromElements(p0)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(p2)));
          let _rhs210 = (p2).plus(p2);
          let _lhs181 = globalState;
          _1189_v46 = _rhs208;
          _1191_v48 = _rhs209;
          _lhs181.f8 = _rhs210;
          _1188_v45 = _1188_v45;
          (globalState).f0 = _1132_v1;
          let _1192_v49;
          _1192_v49 = _dafny.Seq.UnicodeFromString("lx");
          let _1193_v50;
          _1193_v50 = _module.D1.create_DC3(p1);
          _1192_v49 = _dafny.Seq.Concat(_dafny.Seq.Concat(p1, (_1193_v50).dtor_cf6), p1);
        }
        let _1194_v51;
        _1194_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1132_v1,_dafny.Map.Empty.slice().updateUnsafe(p2,p0));
        let _1195_v52;
        _1195_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1132_v1,_1132_v1);
        let _1196_v53;
        _1196_v53 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1197_v54;
        _1197_v54 = _module.D0.create_DC1(((_this).f27).plus((_this).f27), _1132_v1, (_1194_v51).Merge(_1194_v51), _module.__default.fm7(new BigNumber((_1195_v52).length), globalState), _1196_v53);
        let _source18 = _1197_v54;
        if (_source18.is_DC0) {
          (globalState).f0 = !_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-337)), ((_1198_v53) => function (_1199_i4) {
            return _1198_v53;
          })(_1196_v53)), _dafny.Seq.UnicodeFromString("qewbkgcot")), _1196_v53);
          (globalState).f11 = _module.__default.fm7(new BigNumber(93), globalState);
          let _1200_v55;
          _1200_v55 = _dafny.Seq.of(_1132_v1);
          let _1201_v56;
          _1201_v56 = _dafny.Seq.of(new BigNumber((p1).length), p2);
          let _1202_v57;
          _1202_v57 = _dafny.Seq.of(_1200_v55, _dafny.Seq.Concat(_1200_v55, (_this).fm4(_1201_v56, p1, _dafny.Set.fromElements(!(_1132_v1)), p2, globalState)));
          _1202_v57 = _1202_v57;
          let _1203_v58;
          let _init22 = ((_1204_v1) => function (_1205_i5) {
            return ((_1204_v1) ? (_1204_v1) : (_1204_v1));
          })(_1132_v1);
          let _nw157 = Array((new BigNumber(18)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw157.length); _i0_22++) {
            _nw157[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _1203_v58 = _nw157;
          _1203_v58 = _1203_v58;
        } else {
          let _1206___mcc_h5 = (_source18).cf0;
          let _1207___mcc_h6 = (_source18).cf1;
          let _1208___mcc_h7 = (_source18).cf2;
          let _1209___mcc_h8 = (_source18).cf3;
          let _1210___mcc_h9 = (_source18).cf4;
          let _1211_cf4 = _1210___mcc_h9;
          let _1212_cf3 = _1209___mcc_h8;
          let _1213_cf2 = _1208___mcc_h7;
          let _1214_cf1 = _1207___mcc_h6;
          let _1215_cf0 = _1206___mcc_h5;
          let _1216_v59;
          let _1217_v60;
          let _out52;
          let _out53;
          let _outcollector15 = (_this).m3(_dafny.Seq.UnicodeFromString("bubrnhav"), _1132_v1, globalState);
          _out52 = _outcollector15[0];
          _out53 = _outcollector15[1];
          _1216_v59 = _out52;
          _1217_v60 = _out53;
          let _1218_v61;
          let _init23 = ((_1219_cf0) => function (_1220_i6) {
            return (_1220_i6).minus(_1219_cf0);
          })(_1215_cf0);
          let _nw158 = Array((new BigNumber(21)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw158.length); _i0_23++) {
            _nw158[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1218_v61 = _nw158;
          _1218_v61 = _1218_v61;
          let _1221_v62;
          _1221_v62 = _dafny.Seq.UnicodeFromString("kgfnnee");
          _1221_v62 = _dafny.Seq.update(p1, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_module.__default.fm9(globalState))).cardinality()), new BigNumber((p1).length)), _1211_cf4);
          let _1222_v63;
          _1222_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1214_cf1,_1211_cf4);
          let _1223_v64;
          _1223_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_1222_v63);
          let _1224_v65;
          let _nw159 = Array((new BigNumber(8)).toNumber());
          _nw159[(_dafny.ZERO).toNumber()] = (_1222_v63).Merge(_1222_v63);
          _nw159[(_dafny.ONE).toNumber()] = (((_1223_v64).contains(_1215_cf0)) ? ((_1223_v64).get(_1215_cf0)) : (_1222_v63));
          _nw159[(new BigNumber(2)).toNumber()] = _1222_v63;
          _nw159[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1214_cf1,_1196_v53);
          _nw159[(new BigNumber(4)).toNumber()] = (_1222_v63).update(_1214_cf1, _1196_v53);
          _nw159[(new BigNumber(5)).toNumber()] = _1222_v63;
          _nw159[(new BigNumber(6)).toNumber()] = (_1222_v63).Merge(_1222_v63);
          _nw159[(new BigNumber(7)).toNumber()] = (_1222_v63).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1214_cf1,_1211_cf4));
          _1224_v65 = _nw159;
          let _index177 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_1224_v65).length));
          (_1224_v65)[_index177] = (_1222_v63).update(_1132_v1, new _dafny.CodePoint('e'.codePointAt(0)));
          let _1225_v66;
          let _nw160 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
          _1225_v66 = _nw160;
          let _1226_v67;
          _1226_v67 = _dafny.Seq.of(_1132_v1);
          let _index178 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_1225_v66).length));
          (_1225_v66)[_index178] = _1226_v67;
          let _1227_v68;
          _1227_v68 = _dafny.Seq.of((_this).f27);
          let _1228_v69;
          _1228_v69 = _dafny.MultiSet.fromElements(new BigNumber((_1227_v68).length), _1215_cf0);
          let _index179 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_1224_v65).length));
          let _index180 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_1225_v66).length));
          let _rhs211 = _1132_v1;
          let _rhs212 = _1132_v1;
          let _rhs213 = (_1222_v63).Merge(_1222_v63);
          let _rhs214 = _dafny.Seq.Concat(_dafny.Seq.of(_1132_v1), _dafny.Seq.of((_module.D0.create_DC1((((_1228_v69).contains(_1212_cf3)) ? ((_1228_v69).get(_1212_cf3)) : (new BigNumber((_module.__default.fm10(new BigNumber((_1221_v62).length), _1132_v1, _1212_cf3, globalState)).length))), false, _1213_cf2, _module.__default.fm7(p0, globalState), _1196_v53)).dtor_cf1, _1214_cf1, _1132_v1));
          let _lhs182 = globalState;
          let _lhs183 = _1224_v65;
          let _lhs184 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_1224_v65).length));
          let _lhs185 = _1225_v66;
          let _lhs186 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_1225_v66).length));
          _lhs182.f10 = _rhs211;
          _1214_cf1 = _rhs212;
          _lhs183[_lhs184] = _rhs213;
          _lhs185[_lhs186] = _rhs214;
        }
        _1196_v53 = _1196_v53;
      } else {
        let _1229___mcc_h0 = (_source17).cf0;
        let _1230___mcc_h1 = (_source17).cf1;
        let _1231___mcc_h2 = (_source17).cf2;
        let _1232___mcc_h3 = (_source17).cf3;
        let _1233___mcc_h4 = (_source17).cf4;
        let _1234_cf4 = _1233___mcc_h4;
        let _1235_cf3 = _1232___mcc_h3;
        let _1236_cf2 = _1231___mcc_h2;
        let _1237_cf1 = _1230___mcc_h1;
        let _1238_cf0 = _1229___mcc_h0;
        let _1239_v70;
        _1239_v70 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_module.__default.fm11(_1238_cf0, _1234_cf4, _1237_cf1, globalState)).length));
        let _1240_v71;
        _1240_v71 = _dafny.Seq.of(_1238_cf0, new BigNumber(-250), (((_1239_v70).contains((_this).f27)) ? ((_1239_v70).get((_this).f27)) : ((_this).f27)), (_dafny.ZERO).minus((_this).f27), p2);
        let _1241_v72;
        _1241_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1237_cf1,_1238_cf0);
        let _1242_v73;
        _1242_v73 = _dafny.Set.fromElements((_this).f27, (((_1241_v72).contains(_1237_cf1)) ? ((_1241_v72).get(_1237_cf1)) : (_1235_cf3)));
        let _1243_v74;
        _1243_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1237_cf1,_1242_v73);
        let _1244_v75;
        _1244_v75 = _module.D0.create_DC1(new BigNumber((_1240_v71).length), _1237_cf1, _dafny.Map.Empty.slice().updateUnsafe(_1237_cf1,_1239_v70), p2, _1234_cf4);
        let _1245_v76;
        _1245_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v73,_1238_cf0);
        let _1246_v77;
        let _nw161 = Array((new BigNumber(25)).toNumber());
        _nw161[(_dafny.ZERO).toNumber()] = new BigNumber((_1240_v71).length);
        _nw161[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1243_v74).length));
        _nw161[(new BigNumber(2)).toNumber()] = new BigNumber(-82);
        _nw161[(new BigNumber(3)).toNumber()] = _1235_cf3;
        _nw161[(new BigNumber(4)).toNumber()] = new BigNumber(191);
        _nw161[(new BigNumber(5)).toNumber()] = _1235_cf3;
        _nw161[(new BigNumber(6)).toNumber()] = _1235_cf3;
        _nw161[(new BigNumber(7)).toNumber()] = p2;
        _nw161[(new BigNumber(8)).toNumber()] = p0;
        _nw161[(new BigNumber(9)).toNumber()] = p2;
        _nw161[(new BigNumber(10)).toNumber()] = (_1240_v71)[_module.__default.safeIndex(new BigNumber(322), new BigNumber((_1240_v71).length))];
        _nw161[(new BigNumber(11)).toNumber()] = _1238_cf0;
        _nw161[(new BigNumber(12)).toNumber()] = p0;
        _nw161[(new BigNumber(13)).toNumber()] = p0;
        _nw161[(new BigNumber(14)).toNumber()] = _1238_cf0;
        _nw161[(new BigNumber(15)).toNumber()] = (_this).f27;
        _nw161[(new BigNumber(16)).toNumber()] = (_this).f27;
        _nw161[(new BigNumber(17)).toNumber()] = new BigNumber(65);
        _nw161[(new BigNumber(18)).toNumber()] = _1235_cf3;
        _nw161[(new BigNumber(19)).toNumber()] = new BigNumber(750);
        _nw161[(new BigNumber(20)).toNumber()] = (_1244_v75).dtor_cf3;
        _nw161[(new BigNumber(21)).toNumber()] = (_this).f27;
        _nw161[(new BigNumber(22)).toNumber()] = (((_1245_v76).contains(_1242_v73)) ? ((_1245_v76).get(_1242_v73)) : (p2));
        _nw161[(new BigNumber(23)).toNumber()] = new BigNumber((_1241_v72).length);
        _nw161[(new BigNumber(24)).toNumber()] = p0;
        _1246_v77 = _nw161;
        let _1247_v78;
        _1247_v78 = _module.D2.create_DC5(_1246_v77);
        let _1248_v79;
        let _nw162 = Array((new BigNumber(14)).toNumber());
        _nw162[(_dafny.ZERO).toNumber()] = _1246_v77;
        _nw162[(_dafny.ONE).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(2)).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(3)).toNumber()] = ((_1237_cf1) ? (_1246_v77) : (_1246_v77));
        _nw162[(new BigNumber(4)).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(5)).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(6)).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(7)).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(8)).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(9)).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(10)).toNumber()] = ((_1237_cf1) ? (_1246_v77) : ((_1247_v78).dtor_cf8));
        _nw162[(new BigNumber(11)).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(12)).toNumber()] = _1246_v77;
        _nw162[(new BigNumber(13)).toNumber()] = _1246_v77;
        _1248_v79 = _nw162;
        let _index181 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_1248_v79).length));
        (_1248_v79)[_index181] = _1246_v77;
        let _index182 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_1248_v79).length));
        let _rhs215 = _1246_v77;
        let _rhs216 = (new BigNumber((_module.__default.fm6(globalState)).length)).plus((_dafny.ZERO).minus(_1235_cf3));
        let _rhs217 = _1234_cf4;
        let _rhs218 = (_this).f27;
        let _lhs187 = _1248_v79;
        let _lhs188 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_1248_v79).length));
        let _lhs189 = globalState;
        let _lhs190 = globalState;
        _lhs187[_lhs188] = _rhs215;
        _lhs189.f16 = _rhs216;
        _1234_cf4 = _rhs217;
        _lhs190.f16 = _rhs218;
        let _1249_v80;
        let _init24 = ((_1250_v75) => function (_1251_i7) {
          return (_1250_v75).dtor_cf1;
        })(_1244_v75);
        let _nw163 = Array((new BigNumber(12)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw163.length); _i0_24++) {
          _nw163[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _1249_v80 = _nw163;
        _1249_v80 = _1249_v80;
        (globalState).f10 = false;
        _1237_cf1 = false;
      }
      let _hi2 = p0;
      for (let _1252_i8 = (_dafny.ZERO).minus(new BigNumber((p1).length)); _1252_i8.isLessThan(_hi2); _1252_i8 = _1252_i8.plus(_dafny.ONE)) {
        let _1253_v81;
        _1253_v81 = _dafny.Seq.of(_module.__default.safeDivisionInt(new BigNumber(617), _module.__default.fm7(_1252_i8, globalState)));
        let _1254_v82;
        _1254_v82 = _dafny.Map.Empty.slice().updateUnsafe((p2).plus((_this).f27),_dafny.Seq.Concat(_dafny.Seq.of(p0, (_dafny.ZERO).minus((_this).f27), _1252_i8, (_this).f27), _1253_v81));
        _1253_v81 = _dafny.Seq.update((((_1254_v82).contains((_dafny.ZERO).minus(p2))) ? ((_1254_v82).get((_dafny.ZERO).minus(p2))) : (_1253_v81)), _module.__default.safeIndex(p2, new BigNumber(((((_1254_v82).contains((_dafny.ZERO).minus(p2))) ? ((_1254_v82).get((_dafny.ZERO).minus(p2))) : (_1253_v81))).length)), (p0).minus(p2));
        let _1255_v83;
        let _nw164 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
        _1255_v83 = _nw164;
        _1255_v83 = _1255_v83;
        let _1256_v84;
        _1256_v84 = true;
        if (((_this).fm5(_1256_v84, globalState)) === (_1256_v84)) {
          (globalState).f10 = (_1256_v84) || (_1256_v84);
          let _1257_v85;
          let _nw165 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1257_v85 = _nw165;
          let _nw166 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1257_v85 = _nw166;
          (globalState).f0 = _1256_v84;
          let _1258_v86;
          let _nw167 = Array((new BigNumber(16)).toNumber()).fill(_module.D2.Default());
          _1258_v86 = _nw167;
          let _index183 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1258_v86).length));
          (_1258_v86)[_index183] = _module.D2.create_DC5(_1257_v85);
          let _pat_let_tv54 = _1257_v85;
          let _index184 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1258_v86).length));
          (_1258_v86)[_index184] = function (_pat_let22_0) {
            return function (_1259_dt__update__tmp_h1) {
              return function (_pat_let23_0) {
                return function (_1260_dt__update_hcf8_h0) {
                  return _module.D2.create_DC5(_1260_dt__update_hcf8_h0);
                }(_pat_let23_0);
              }(_pat_let_tv54);
            }(_pat_let22_0);
          }(_module.D2.create_DC5(_1257_v85));
          (globalState).f11 = _module.__default.fm7((_this).f27, globalState);
        } else {
          let _1261_v87;
          _1261_v87 = new _dafny.CodePoint('h'.codePointAt(0));
          _1261_v87 = _module.__default.fm12((_dafny.Map.Empty.slice().updateUnsafe(!(_1256_v84),(_1253_v81)[_module.__default.safeIndex(p0, new BigNumber((_1253_v81).length))])).contains(_1256_v84), new BigNumber(287), globalState);
          let _1262_v88;
          let _1263_v89;
          let _out54;
          let _out55;
          let _outcollector16 = (_this).m3(p1, _1256_v84, globalState);
          _out54 = _outcollector16[0];
          _out55 = _outcollector16[1];
          _1262_v88 = _out54;
          _1263_v89 = _out55;
          let _1264_v90;
          _1264_v90 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm6(globalState),new BigNumber((p1).length));
          _1264_v90 = (_1264_v90).update(p1, p2);
          let _1265_v91;
          let _init25 = ((_1266_p1, _1267_p2) => function (_1268_i9) {
            return (new BigNumber((_dafny.Seq.update(_1266_p1, _module.__default.safeIndex(_1267_p2, new BigNumber((_1266_p1).length)), new _dafny.CodePoint('u'.codePointAt(0)))).length)).isEqualTo(_1267_p2);
          })(p1, p2);
          let _nw168 = Array((new BigNumber(17)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw168.length); _i0_25++) {
            _nw168[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1265_v91 = _nw168;
          let _index185 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_1265_v91).length));
          (_1265_v91)[_index185] = _1256_v84;
          let _index186 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_1265_v91).length));
          (_1265_v91)[_index186] = _1256_v84;
          let _1269_v92;
          _1269_v92 = _dafny.Seq.UnicodeFromString("a");
          let _1270_v93;
          _1270_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1256_v84,_1256_v84);
          let _1271_v94;
          _1271_v94 = _module.D2.create_DC7(_1253_v81, _1256_v84, _1256_v84, new BigNumber((_1270_v93).length));
          let _1272_v95;
          _1272_v95 = _dafny.MultiSet.fromElements(_1271_v94);
          let _1273_v96;
          _1273_v96 = _dafny.Map.Empty.slice().updateUnsafe(_1252_i8,(_this).f27);
          let _1274_v97;
          _1274_v97 = _module.D3.create_DC10(_1273_v96);
          let _1275_v98;
          _1275_v98 = _dafny.Map.Empty.slice().updateUnsafe(_1256_v84,(_1274_v97).dtor_cf18);
          let _1276_v99;
          _1276_v99 = _module.D0.create_DC1(p0, _1256_v84, _1275_v98, p0, _1261_v87);
          let _index187 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_1265_v91).length));
          let _index188 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_1265_v91).length));
          let _rhs219 = (_1272_v95).IsProperSubsetOf(_1272_v95);
          let _rhs220 = (_1276_v99).dtor_cf1;
          let _rhs221 = new BigNumber(29);
          let _rhs222 = (p1)[_module.__default.safeIndex((_this).f27, new BigNumber((p1).length))];
          let _rhs223 = p1;
          let _lhs191 = _1265_v91;
          let _lhs192 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_1265_v91).length));
          let _lhs193 = _1265_v91;
          let _lhs194 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_1265_v91).length));
          let _lhs195 = globalState;
          _lhs191[_lhs192] = _rhs219;
          _lhs193[_lhs194] = _rhs220;
          _lhs195.f11 = _rhs221;
          _1261_v87 = _rhs222;
          _1269_v92 = _rhs223;
          let _1277_v100;
          _1277_v100 = _dafny.Map.Empty.slice().updateUnsafe(_1261_v87,_1262_v88);
          _1277_v100 = (_1277_v100).update(_module.__default.fm12((_1265_v91)[_module.__default.safeIndex(new BigNumber(651), new BigNumber((_1265_v91).length))], new BigNumber((_1269_v92).length), globalState), _1252_i8);
        }
        let _1278_v101;
        _1278_v101 = _dafny.Map.Empty.slice().updateUnsafe(_1252_i8,p0);
        let _1279_v102;
        _1279_v102 = _module.D3.create_DC10(_1278_v101);
        let _source19 = _1279_v102;
        if (_source19.is_DC11) {
          let _1280___mcc_h10 = (_source19).cf19;
          let _1281___mcc_h11 = (_source19).cf20;
          let _1282_cf20 = _1281___mcc_h11;
          let _1283_cf19 = _1280___mcc_h10;
          _1283_cf19 = _1256_v84;
          (globalState).f0 = _1256_v84;
          (globalState).f0 = !(_1283_cf19);
          let _1284_v103;
          _1284_v103 = _dafny.Seq.UnicodeFromString("bkhdpyee");
          let _1285_v104;
          _1285_v104 = _dafny.Seq.of(p1);
          let _1286_v105;
          _1286_v105 = _dafny.Map.Empty.slice().updateUnsafe(_1283_cf19,_1278_v101);
          let _1287_v106;
          _1287_v106 = _module.D0.create_DC1(_1252_i8, _1283_cf19, _1286_v105, (_dafny.ZERO).minus(p2), new _dafny.CodePoint('e'.codePointAt(0)));
          _1284_v103 = _dafny.Seq.Concat((_1285_v104)[_module.__default.safeIndex((_1287_v106).dtor_cf0, new BigNumber((_1285_v104).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(-122)), function (_1288_i10) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }));
        } else {
          let _1289___mcc_h12 = (_source19).cf18;
          let _1290_cf18 = _1289___mcc_h12;
          let _1291_v107;
          _1291_v107 = _dafny.Seq.of(_1256_v84);
          let _1292_v108;
          _1292_v108 = _module.D3.create_DC11(_1256_v84, (_1291_v107)[_module.__default.safeIndex((_this).f27, new BigNumber((_1291_v107).length))]);
          (globalState).f10 = !((_1292_v108).dtor_cf19);
          let _1293_v109;
          _1293_v109 = new _dafny.CodePoint('h'.codePointAt(0));
          let _1294_v110;
          _1294_v110 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.of(_1293_v109));
          let _1295_v111;
          let _nw169 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _1295_v111 = _nw169;
          let _1296_v112;
          _1296_v112 = _dafny.MultiSet.fromElements(_1295_v111, _1295_v111);
          let _1297_v113;
          _1297_v113 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(((((_1294_v110).contains(_1256_v84)) ? ((_1294_v110).get(_1256_v84)) : (p1))).length), new BigNumber((_dafny.Set.fromElements((((_1296_v112).contains(_1295_v111)) ? ((_1296_v112).get(_1295_v111)) : (_1252_i8)))).length))),true);
          _1297_v113 = (_this).f26;
          _1295_v111 = _1295_v111;
          let _1298_v114;
          _1298_v114 = _dafny.Set.fromElements(_1293_v109, _1293_v109);
          let _1299_v115;
          _1299_v115 = _dafny.Seq.of(_1297_v113, (_1297_v113).Merge(_1297_v113), (_this).f26, (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_1256_v84)).update(new BigNumber((_1298_v114).length), _1256_v84));
          let _1300_v116;
          _1300_v116 = _dafny.MultiSet.fromElements(_module.__default.fm7((_this).f27, globalState), new BigNumber(122), new BigNumber((_1291_v107).length), new BigNumber(107));
          let _1301_v117;
          _1301_v117 = _dafny.Map.Empty.slice().updateUnsafe((_1300_v116).IsProperSubsetOf(_1300_v116),!(!(_1256_v84)));
          let _rhs224 = (p2).isLessThan(p0);
          let _rhs225 = _1256_v84;
          let _rhs226 = _1252_i8;
          let _rhs227 = _1299_v115;
          let _rhs228 = _1301_v117;
          let _lhs196 = globalState;
          let _lhs197 = globalState;
          let _lhs198 = globalState;
          _lhs196.f0 = _rhs224;
          _lhs197.f0 = _rhs225;
          _lhs198.f8 = _rhs226;
          _1299_v115 = _rhs227;
          _1301_v117 = _rhs228;
        }
      }
      let _1302_v119;
      _1302_v119 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,p2);
      let _1303_v120;
      _1303_v120 = false;
      let _1304_v121;
      _1304_v121 = _dafny.Map.Empty.slice().updateUnsafe(_1303_v120,p2);
      let _1305_v122;
      _1305_v122 = _dafny.MultiSet.fromElements(new BigNumber((_1304_v121).length), (_this).f27);
      let _1306_v123;
      _1306_v123 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),_1305_v122);
      let _1307_v124;
      _1307_v124 = _dafny.Map.Empty.slice().updateUnsafe((function () {
        let _coll52 = new _dafny.Map();
        for (const _compr_52 of ((_1302_v119).update((_this).f27, p0)).Keys.Elements) {
          let _1308_v118 = _compr_52;
          if (((_1302_v119).update((_this).f27, p0)).contains(_1308_v118)) {
            _coll52.push([(_1308_v118).multipliedBy((_this).f27),_dafny.MultiSet.fromElements((_this).f27)]);
          }
        }
        return _coll52;
      }()).Merge((_module.D4.create_DC12(_1306_v123)).dtor_cf21),_1303_v120);
      let _1309_v125;
      _1309_v125 = _dafny.Seq.of(_1306_v123);
      _1307_v124 = (_1307_v124).update((_1309_v125)[_module.__default.safeIndex(p0, new BigNumber((_1309_v125).length))], _1303_v120);
      if (_1303_v120) {
        (globalState).f6 = new BigNumber((p1).length);
        let _1310_v126;
        _1310_v126 = _dafny.Map.Empty.slice().updateUnsafe(_1303_v120,_1303_v120);
        let _1311_v127;
        let _init26 = ((_1312_v120) => function (_1313_i11) {
          return _1312_v120;
        })(_1303_v120);
        let _nw170 = Array((new BigNumber(28)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw170.length); _i0_26++) {
          _nw170[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1311_v127 = _nw170;
        let _1314_v128;
        _1314_v128 = _dafny.Seq.of(_1311_v127);
        let _1315_v129;
        let _nw171 = new _module.C2();
        _nw171.__ctor(new BigNumber((((_1303_v120) ? (_1310_v126) : (_1310_v126))).length), _dafny.Seq.Concat(_dafny.Seq.of(_1311_v127, _1311_v127), _1314_v128));
        _1315_v129 = _nw171;
        let _1316_v130;
        _1316_v130 = _dafny.Seq.of(_1305_v122);
        let _1317_v131;
        _1317_v131 = _module.D11.create_DC36(_1316_v130);
        let _rhs229 = _1303_v120;
        let _rhs230 = _1317_v131;
        let _lhs199 = globalState;
        _lhs199.f10 = _rhs229;
        _1317_v131 = _rhs230;
        _1304_v121 = (_1304_v121).update(_1303_v120, new BigNumber((p1).length));
        let _source20 = _1131_v0;
        if (_source20.is_DC0) {
          let _1318_v132;
          let _nw172 = new _module.C3();
          _nw172.__ctor(_dafny.Seq.Concat(_1314_v128, _1314_v128));
          _1318_v132 = _nw172;
          (globalState).f8 = (_this).f27;
          let _1319_v133;
          _1319_v133 = _dafny.MultiSet.fromElements(_1303_v120, _1303_v120);
          let _1320_v134;
          _1320_v134 = _dafny.MultiSet.fromElements(_1319_v133, _1319_v133, _1319_v133);
          (globalState).f1 = (_1320_v134).Difference(_1320_v134);
          let _index189 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_1311_v127).length));
          (_1311_v127)[_index189] = false;
          let _index190 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_1311_v127).length));
          (_1311_v127)[_index190] = !(!(_1303_v120) || (true));
        } else {
          let _1321___mcc_h13 = (_source20).cf0;
          let _1322___mcc_h14 = (_source20).cf1;
          let _1323___mcc_h15 = (_source20).cf2;
          let _1324___mcc_h16 = (_source20).cf3;
          let _1325___mcc_h17 = (_source20).cf4;
          let _1326_cf4 = _1325___mcc_h17;
          let _1327_cf3 = _1324___mcc_h16;
          let _1328_cf2 = _1323___mcc_h15;
          let _1329_cf1 = _1322___mcc_h14;
          let _1330_cf0 = _1321___mcc_h13;
          let _1331_v135;
          _1331_v135 = _dafny.Seq.UnicodeFromString("riagx");
          let _rhs231 = p1;
          let _rhs232 = _1327_cf3;
          _1331_v135 = _rhs231;
          _1327_cf3 = _rhs232;
          let _1332_v136;
          let _nw173 = new _module.C0();
          _nw173.__ctor();
          _1332_v136 = _nw173;
          let _1333_v137;
          let _nw174 = new _module.C4();
          _nw174.__ctor(!(!(_1303_v120)), _1314_v128);
          _1333_v137 = _nw174;
          let _1334_v138;
          _1334_v138 = _dafny.Set.fromElements(_1333_v137);
          (globalState).f8 = ((p2).multipliedBy(new BigNumber((_1334_v138).length))).minus(new BigNumber((p1).length));
          let _1335_v139;
          _1335_v139 = _dafny.Seq.of((_1315_v129).f29, (_1315_v129).f29);
          let _1336_v140;
          _1336_v140 = _module.D2.create_DC6(new BigNumber((_1335_v139).length));
          (globalState).f0 = ((((_this).f26).contains((_1330_cf0).plus(new BigNumber((_module.__default.fm14(p2, (_1336_v140).dtor_cf9, globalState)).length)))) ? (((_this).f26).get((_1330_cf0).plus(new BigNumber((_module.__default.fm14(p2, (_1336_v140).dtor_cf9, globalState)).length)))) : (((_this).f27).isLessThanOrEqualTo((_1315_v129).f29)));
        }
      } else {
        let _1337_v141;
        _1337_v141 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _1338_v142;
        _1338_v142 = _dafny.Seq.of((_this).f27);
        _1337_v141 = (_1337_v141).update(_dafny.Seq.update(_dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("kceepx")), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("kceepx"))).length)), new _dafny.CodePoint('j'.codePointAt(0))), new BigNumber((_dafny.MultiSet.fromElements(p2, new BigNumber((_1338_v142).length), p0)).cardinality()));
        let _1339_v143;
        _1339_v143 = _module.D4.create_DC15(p2, new BigNumber((_1304_v121).length));
        (globalState).f10 = !((_module.__default.fm7((_1339_v143).dtor_cf26, globalState)).isEqualTo(p2));
        let _1340_v144;
        _1340_v144 = _dafny.Seq.UnicodeFromString("bxtwiau");
        _1340_v144 = _1340_v144;
        let _1341_v145;
        _1341_v145 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1340_v144);
        let _1342_v146;
        _1342_v146 = _dafny.Set.fromElements(_1303_v120);
        let _1343_v147;
        _1343_v147 = _dafny.Set.fromElements(new BigNumber((_1342_v146).length), p0);
        let _1344_v148;
        _1344_v148 = _dafny.Map.Empty.slice().updateUnsafe(_1343_v147,_1303_v120);
        _1341_v145 = (_1341_v145).update((p2).plus(new BigNumber((_1344_v148).length)), ((!(_1303_v120)) ? (p1) : (p1)));
        (globalState).f0 = _1303_v120;
      }
      let _1345_v149;
      _1345_v149 = new _dafny.CodePoint('q'.codePointAt(0));
      let _1346_v150;
      let _nw175 = Array((new BigNumber(25)).toNumber()).fill(false);
      _1346_v150 = _nw175;
      let _pat_let_tv55 = _1345_v149;
      let _1347_v151;
      let _nw176 = new _module.C2();
      _nw176.__ctor(p2, _dafny.Seq.of((function (_pat_let24_0) {
        return function (_1348_dt__update__tmp_h2) {
          return function (_pat_let25_0) {
            return function (_1349_dt__update_hcf23_h0) {
              return _module.D4.create_DC14(_1349_dt__update_hcf23_h0, (_1348_dt__update__tmp_h2).dtor_cf24, (_1348_dt__update__tmp_h2).dtor_cf25);
            }(_pat_let25_0);
          }(_pat_let_tv55);
        }(_pat_let24_0);
      }(_module.D4.create_DC14(_1345_v149, _1346_v150, _dafny.Seq.UnicodeFromString("jixw")))).dtor_cf24));
      _1347_v151 = _nw176;
      let _1350_v152;
      _1350_v152 = _dafny.MultiSet.fromElements(false, _1303_v120);
      let _1351_v153;
      _1351_v153 = _dafny.Seq.of(_1303_v120);
      (globalState).f11 = new BigNumber((_dafny.Seq.update(_module.__default.fm24((_1350_v152).Difference(_1350_v152), p2, (new BigNumber((_1351_v153).length)).minus(p2), !_dafny.areEqual(_dafny.Seq.UnicodeFromString("jrtm"), p1), globalState), _module.__default.safeIndex((_dafny.ZERO).minus((_1347_v151).f29), new BigNumber((_module.__default.fm24((_1350_v152).Difference(_1350_v152), p2, (new BigNumber((_1351_v153).length)).minus(p2), !_dafny.areEqual(_dafny.Seq.UnicodeFromString("jrtm"), p1), globalState)).length)), p2)).length);
      return;
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _module.D0.Default();
      let _1352_v0;
      let _nw177 = Array((new BigNumber(16)).toNumber()).fill(false);
      _1352_v0 = _nw177;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1352_v0).length))) {
        let _1353_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1353_i0)) && ((_1353_i0).isLessThan(new BigNumber((_1352_v0).length))))) {
          (_1352_v0)[(_1353_i0)] = !(!(((_this).f27).isLessThan(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("mtnfgu")).length), (_this).f27))));
        }
      }
      let _1354_v1;
      _1354_v1 = _dafny.Seq.of((_this).f27);
      let _1355_v2;
      _1355_v2 = _dafny.Seq.of(_1354_v1);
      let _1356_v4;
      _1356_v4 = _module.D7.create_DC25(function () {
  let _coll53 = new _dafny.Set();
  for (const _compr_53 of (_1355_v2).Elements) {
    let _1357_v3 = _compr_53;
    if (_dafny.Seq.contains(_1355_v2, _1357_v3)) {
      _coll53.add(_1357_v3);
    }
  }
  return _coll53;
}());
      let _1358_v5;
      _1358_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1356_v4,(_this).f27);
      let _1359_v7;
      _1359_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,function () {
        let _coll54 = new _dafny.Map();
        for (const _compr_54 of _dafny.IntegerRange(new BigNumber(-409), new BigNumber(24))) {
          let _1360_v6 = _compr_54;
          if (((new BigNumber(-409)).isLessThanOrEqualTo(_1360_v6)) && ((_1360_v6).isLessThan(new BigNumber(24)))) {
            _coll54.push([(_1360_v6).multipliedBy((_this).f27),new BigNumber((p0).length)]);
          }
        }
        return _coll54;
      }());
      let _1361_v8;
      _1361_v8 = _dafny.Seq.of(p1, p1);
      let _1362_v9;
      _1362_v9 = new _dafny.CodePoint('u'.codePointAt(0));
      let _1363_v10;
      _1363_v10 = _module.D0.create_DC1(new BigNumber((_1358_v5).length), p1, _1359_v7, new BigNumber((_dafny.Seq.of(new BigNumber((_1361_v8).length))).length), _1362_v9);
      let _1364_v11;
      _1364_v11 = _dafny.MultiSet.fromElements(_1363_v10, _1363_v10);
      let _1365_v12;
      let _nw178 = Array((_dafny.ONE).toNumber());
      _nw178[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.fromElements(_1363_v10)).Difference(_1364_v11);
      _1365_v12 = _nw178;
      let _index191 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1365_v12).length));
      (_1365_v12)[_index191] = _1364_v11;
      let _index192 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1365_v12).length));
      (_1365_v12)[_index192] = _1364_v11;
      let _1366_v13;
      _1366_v13 = _module.D6.create_DC22((_this).f27);
      let _hi3 = _module.__default.safeDivisionInt((_this).f27, (_1366_v13).dtor_cf36);
      for (let _1367_i1 = (_this).f27; _1367_i1.isLessThan(_hi3); _1367_i1 = _1367_i1.plus(_dafny.ONE)) {
        let _1368_v14;
        let _nw179 = new _module.C0();
        _nw179.__ctor();
        _1368_v14 = _nw179;
        let _1369_v15;
        _1369_v15 = _dafny.Seq.of(_1368_v14, _1368_v14);
        let _1370_v16;
        let _nw180 = Array((new BigNumber(21)).toNumber());
        _nw180[(_dafny.ZERO).toNumber()] = _1368_v14;
        _nw180[(_dafny.ONE).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(2)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(3)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(4)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(5)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(6)).toNumber()] = (_1369_v15)[_module.__default.safeIndex(_1367_i1, new BigNumber((_1369_v15).length))];
        _nw180[(new BigNumber(7)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(8)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(9)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(10)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(11)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(12)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(13)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(14)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(15)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(16)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(17)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(18)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(19)).toNumber()] = _1368_v14;
        _nw180[(new BigNumber(20)).toNumber()] = _1368_v14;
        _1370_v16 = _nw180;
        let _1371_v17;
        let _nw181 = Array((new BigNumber(4)).toNumber());
        _nw181[(_dafny.ZERO).toNumber()] = _1370_v16;
        _nw181[(_dafny.ONE).toNumber()] = _1370_v16;
        _nw181[(new BigNumber(2)).toNumber()] = _1370_v16;
        _nw181[(new BigNumber(3)).toNumber()] = _1370_v16;
        _1371_v17 = _nw181;
        _1371_v17 = _1371_v17;
        let _1372_v18;
        _1372_v18 = _dafny.Seq.of(_1352_v0, _1352_v0, _1352_v0, _1352_v0);
        let _1373_v19;
        let _nw182 = new _module.C4();
        _nw182.__ctor(p1, _1372_v18);
        _1373_v19 = _nw182;
        (_1373_v19).f28 = p1;
        let _1374_v20;
        _1374_v20 = _dafny.Seq.UnicodeFromString("myxqtmm");
        _1374_v20 = _1374_v20;
      }
      let _1375_v21;
      let _nw183 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _1375_v21 = _nw183;
      let _index193 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_1375_v21).length));
      (_1375_v21)[_index193] = new BigNumber(361);
      let _index194 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_1375_v21).length));
      (_1375_v21)[_index194] = (_this).f27;
      let _1376_v22;
      _1376_v22 = _module.D2.create_DC6((_1375_v21)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_1375_v21).length))]);
      let _1377_v23;
      _1377_v23 = _dafny.Map.Empty.slice().updateUnsafe((_1376_v22).dtor_cf9,(_this).f27);
      _1377_v23 = (_1377_v23).update(_module.__default.safeDivisionInt((_this).f27, (_this).f27), (_this).f27);
      let _rhs233 = _module.D6.create_DC22((_this).f27);
      let _rhs234 = p1;
      let _lhs200 = globalState;
      _1366_v13 = _rhs233;
      _lhs200.f10 = _rhs234;
      r0 = _module.__default.safeDivisionInt(new BigNumber((_1361_v8).length), (_dafny.ZERO).minus(((p1) ? (new BigNumber((p0).length)) : ((_1375_v21)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_1375_v21).length))]))));
      r1 = _module.D0.create_DC0();
      return [r0, r1];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f22 = _dafny.Seq.of();
      this._f24 = false;
      this._f25 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
    __ctor(f24, f25, f22) {
      let _this = this;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      (_this)._f22 = f22;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.__default.safeModuloInt(new BigNumber(378), new BigNumber(-343))).plus(new BigNumber(-858));
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f24;
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.ZERO;
      let _1378_v0;
      let _nw184 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1378_v0 = _nw184;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1378_v0).length))) {
        let _1379_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1379_i0)) && ((_1379_i0).isLessThan(new BigNumber((_1378_v0).length))))) {
          (_1378_v0)[(_1379_i0)] = ((p0) ? (_dafny.Seq.UnicodeFromString("tri")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(268)), function (_1380_i1) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })));
        }
      }
      (globalState).f0 = (_this).f24;
      let _1381_v1;
      _1381_v1 = _module.D0.create_DC0();
      let _source21 = _1381_v1;
      if (_source21.is_DC0) {
        let _1382_v2;
        _1382_v2 = new BigNumber(468);
        (globalState).f16 = _1382_v2;
        let _1383_v3;
        let _nw185 = new _module.C3();
        _nw185.__ctor(_dafny.Seq.Concat((_this).f22, (_this).f22));
        _1383_v3 = _nw185;
        if (p0) {
          let _1384_v4;
          let _nw186 = new _module.C2();
          _nw186.__ctor(_1382_v2, _dafny.Seq.Concat((_this).f22, (_this).f22));
          _1384_v4 = _nw186;
          let _1385_v5;
          let _init27 = function (_1386_i2) {
            return _module.D11.create_DC39((_this).f24);
          };
          let _nw187 = Array((new BigNumber(17)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw187.length); _i0_27++) {
            _nw187[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1385_v5 = _nw187;
          let _1387_v6;
          _1387_v6 = _module.D11.create_DC39((_this).f24);
          let _index195 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_1385_v5).length));
          (_1385_v5)[_index195] = _1387_v6;
          let _index196 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_1385_v5).length));
          (_1385_v5)[_index196] = _1387_v6;
          let _1388_v7;
          _1388_v7 = _dafny.Seq.UnicodeFromString("kfs");
          let _1389_v8;
          let _init28 = ((_1390_v4) => function (_1391_i3) {
            return _module.__default.safeDivisionInt(_1391_i3, (_1390_v4).f29);
          })(_1384_v4);
          let _nw188 = Array((new BigNumber(17)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw188.length); _i0_28++) {
            _nw188[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1389_v8 = _nw188;
          let _1392_v9;
          _1392_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1388_v7).length),_1389_v8);
          let _1393_v10;
          _1393_v10 = _dafny.Seq.of(new BigNumber((_1388_v7).length), (_1384_v4).f29);
          let _1394_v11;
          _1394_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1384_v4).f29,new BigNumber((_1393_v10).length));
          let _1395_v12;
          _1395_v12 = _dafny.Seq.of((((_1394_v11).contains(new BigNumber((_1393_v10).length))) ? ((_1394_v11).get(new BigNumber((_1393_v10).length))) : (_module.__default.fm7(_1382_v2, globalState))));
          let _1396_v13;
          _1396_v13 = _module.D2.create_DC5(_1389_v8);
          _1392_v9 = (_1392_v9).update(new BigNumber((_dafny.Set.fromElements(_1382_v2, _1382_v2, (_1395_v12)[_module.__default.safeIndex(_1382_v2, new BigNumber((_1395_v12).length))], new BigNumber((_module.__default.fm33(globalState)).length))).length), (_1396_v13).dtor_cf8);
          let _1397_v14;
          _1397_v14 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1398_v15;
          _1398_v15 = _dafny.MultiSet.fromElements((_this).f24);
          let _1399_v16;
          _1399_v16 = _dafny.Map.Empty.slice().updateUnsafe(!(p0) || ((_this).f24),(_dafny.MultiSet.fromElements(new BigNumber(502), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1397_v14,(_1384_v4).f29)).length))).update(new BigNumber((_1398_v15).cardinality()), _module.__default.abs(_1382_v2)));
          _1399_v16 = (_1399_v16).update((_this).f24, (_dafny.MultiSet.FromArray(_1393_v10)).Intersect(_dafny.MultiSet.fromElements((_1384_v4).f29)));
          let _1400_v17;
          _1400_v17 = _dafny.MultiSet.fromElements((_1384_v4).f29, _1382_v2, new BigNumber((_1388_v7).length), _1382_v2);
          let _1401_v21;
          _1401_v21 = _dafny.Seq.of(function () {
            let _coll55 = new _dafny.Set();
            for (const _compr_55 of (function () {
              let _coll56 = new _dafny.Map();
              for (const _compr_56 of _dafny.IntegerRange(new BigNumber(325), new BigNumber(-867))) {
                let _1402_v18 = _compr_56;
                if (((new BigNumber(325)).isLessThanOrEqualTo(_1402_v18)) && ((_1402_v18).isLessThan(new BigNumber(-867)))) {
                  _coll56.push([(_1402_v18).multipliedBy(_1382_v2),p0]);
                }
              }
              return _coll56;
            }()).Keys.Elements) {
              let _1403_v19 = _compr_55;
              if ((function () {
                let _coll57 = new _dafny.Map();
                for (const _compr_57 of _dafny.IntegerRange(new BigNumber(325), new BigNumber(-867))) {
                  let _1402_v18 = _compr_57;
                  if (((new BigNumber(325)).isLessThanOrEqualTo(_1402_v18)) && ((_1402_v18).isLessThan(new BigNumber(-867)))) {
                    _coll57.push([(_1402_v18).multipliedBy(_1382_v2),p0]);
                  }
                }
                return _coll57;
              }()).contains(_1403_v19)) {
                _coll55.add((_1403_v19).multipliedBy(new BigNumber((function () {
                  let _coll58 = new _dafny.Map();
                  for (const _compr_58 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(219),new BigNumber((_dafny.Set.fromElements(!(false), false, false, true, false)).length))).Keys.Elements) {
                    let _1404_v20 = _compr_58;
                    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(219),new BigNumber((_dafny.Set.fromElements(!(false), false, false, true, false)).length))).contains(_1404_v20)) {
                      _coll58.push([_module.__default.safeDivisionInt(_1404_v20, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false), true))).cardinality()))).length)),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber(-324)))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kqmipstkb")).length),new _dafny.CodePoint('g'.codePointAt(0)))).length))]);
                    }
                  }
                  return _coll58;
                }()).length)));
              }
            }
            return _coll55;
          }());
          let _1405_v22;
          _1405_v22 = _dafny.Set.fromElements(_1382_v2);
          (globalState).f11 = (_1384_v4).fm0(_1400_v17, new BigNumber(574), _dafny.Seq.update(_dafny.Seq.update(_1401_v21, _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1401_v21).length)), _dafny.Set.fromElements((_1384_v4).f29, new BigNumber((_1395_v12).length), _1382_v2, (_1384_v4).f29)), _module.__default.safeIndex((_dafny.ZERO).minus((_1384_v4).f29), new BigNumber((_dafny.Seq.update(_1401_v21, _module.__default.safeIndex(new BigNumber(816), new BigNumber((_1401_v21).length)), _dafny.Set.fromElements((_1384_v4).f29, new BigNumber((_1395_v12).length), _1382_v2, (_1384_v4).f29))).length)), _1405_v22), p0, globalState);
        } else {
          let _1406_v23;
          _1406_v23 = _dafny.Seq.of(_1382_v2);
          let _1407_v24;
          _1407_v24 = _dafny.Seq.of(_1406_v23);
          let _1408_v25;
          _1408_v25 = _dafny.Set.fromElements(_1382_v2, _1382_v2, _1382_v2, _1382_v2);
          r0 = !(_dafny.areEqual(_1406_v23, (_1407_v24)[_module.__default.safeIndex(new BigNumber((_1408_v25).length), new BigNumber((_1407_v24).length))]));
          _1378_v0 = _1378_v0;
          (globalState).f6 = _1382_v2;
          r3 = ((p0) ? (_1382_v2) : (_1382_v2));
          (globalState).f0 = (_this).f24;
        }
        let _1409_v26;
        _1409_v26 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),(_this).f24);
        let _1410_v27;
        _1410_v27 = new _dafny.CodePoint('m'.codePointAt(0));
        let _1411_v28;
        let _nw189 = Array((new BigNumber(9)).toNumber());
        _nw189[(_dafny.ZERO).toNumber()] = (_this).f24;
        _nw189[(_dafny.ONE).toNumber()] = p0;
        _nw189[(new BigNumber(2)).toNumber()] = !((((_1409_v26).contains(_1410_v27)) ? ((_1409_v26).get(_1410_v27)) : ((_this).f24)));
        _nw189[(new BigNumber(3)).toNumber()] = p0;
        _nw189[(new BigNumber(4)).toNumber()] = (_this).f24;
        _nw189[(new BigNumber(5)).toNumber()] = p0;
        _nw189[(new BigNumber(6)).toNumber()] = (_this).f24;
        _nw189[(new BigNumber(7)).toNumber()] = p0;
        _nw189[(new BigNumber(8)).toNumber()] = p0;
        _1411_v28 = _nw189;
        let _1412_v29;
        let _nw190 = new _module.C3();
        _nw190.__ctor(_dafny.Seq.of(_1411_v28, _1411_v28, _1411_v28, _1411_v28, _1411_v28));
        _1412_v29 = _nw190;
      } else {
        let _1413___mcc_h0 = (_source21).cf0;
        let _1414___mcc_h1 = (_source21).cf1;
        let _1415___mcc_h2 = (_source21).cf2;
        let _1416___mcc_h3 = (_source21).cf3;
        let _1417___mcc_h4 = (_source21).cf4;
        let _1418_cf4 = _1417___mcc_h4;
        let _1419_cf3 = _1416___mcc_h3;
        let _1420_cf2 = _1415___mcc_h2;
        let _1421_cf1 = _1414___mcc_h1;
        let _1422_cf0 = _1413___mcc_h0;
        let _1423_v30;
        let _1424_v31;
        let _out56;
        let _out57;
        let _outcollector17 = (_this).m1(globalState);
        _out56 = _outcollector17[0];
        _out57 = _outcollector17[1];
        _1423_v30 = _out56;
        _1424_v31 = _out57;
        let _1425_v32;
        let _init29 = ((_1426_cf0) => function (_1427_i4) {
          return (_1427_i4).multipliedBy(_1426_cf0);
        })(_1422_cf0);
        let _nw191 = Array((_dafny.ONE).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw191.length); _i0_29++) {
          _nw191[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1425_v32 = _nw191;
        let _index197 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_1425_v32).length));
        (_1425_v32)[_index197] = _1419_cf3;
        let _index198 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_1425_v32).length));
        (_1425_v32)[_index198] = ((_dafny.ZERO).minus(_1422_cf0)).minus(_1424_v31);
        if (_1423_v30) {
          let _1428_v33;
          let _nw192 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _1428_v33 = _nw192;
          _1428_v33 = _1428_v33;
          let _1429_v34;
          let _nw193 = new _module.C0();
          _nw193.__ctor();
          _1429_v34 = _nw193;
          (globalState).f6 = _module.__default.safeModuloInt(_1422_cf0, _1422_cf0);
          let _1430_v35;
          let _init30 = ((_1431_cf0) => function (_1432_i5) {
            return _dafny.Set.fromElements(_1431_cf0, _1431_cf0);
          })(_1422_cf0);
          let _nw194 = Array((new BigNumber(28)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw194.length); _i0_30++) {
            _nw194[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1430_v35 = _nw194;
          let _1433_v36;
          _1433_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1419_cf3,_1424_v31);
          let _1434_v37;
          _1434_v37 = _dafny.Seq.of(_1433_v36, _1433_v36);
          let _1435_v38;
          _1435_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1424_v31);
          let _1436_v40;
          _1436_v40 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_1434_v37, _module.__default.safeIndex(new BigNumber((_1435_v38).length), new BigNumber((_1434_v37).length)), function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of _dafny.IntegerRange(new BigNumber(-383), new BigNumber(-369))) {
              let _1437_v39 = _compr_59;
              if (((new BigNumber(-383)).isLessThanOrEqualTo(_1437_v39)) && ((_1437_v39).isLessThan(new BigNumber(-369)))) {
                _coll59.push([_module.__default.safeModuloInt(_1437_v39, (_1425_v32)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_1425_v32).length))]),_1422_cf0]);
              }
            }
            return _coll59;
          }())).length), (_1425_v32)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_1425_v32).length))], _1422_cf0, (_dafny.ZERO).minus(_module.__default.fm7(_1422_cf0, globalState)), _1419_cf3);
          let _index199 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1430_v35).length));
          (_1430_v35)[_index199] = _1436_v40;
          let _index200 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_1430_v35).length));
          (_1430_v35)[_index200] = (function () {
            let _coll60 = new _dafny.Set();
            for (const _compr_60 of _dafny.IntegerRange(new BigNumber(460), new BigNumber(-440))) {
              let _1438_v41 = _compr_60;
              if (((new BigNumber(460)).isLessThanOrEqualTo(_1438_v41)) && ((_1438_v41).isLessThan(new BigNumber(-440)))) {
                _coll60.add((_1438_v41).multipliedBy((_1425_v32)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_1425_v32).length))]));
              }
            }
            return _coll60;
          }()).Union(_1436_v40);
          (globalState).f13 = (_1430_v35)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_1430_v35).length))];
        } else {
          (globalState).f6 = (_dafny.ZERO).minus(_1422_cf0);
          let _1439_v42;
          _1439_v42 = _dafny.Set.fromElements(_1418_cf4, new _dafny.CodePoint('w'.codePointAt(0)), _1418_cf4, _1418_cf4);
          (globalState).f10 = (_dafny.Set.fromElements(_1418_cf4)).IsProperSubsetOf(_1439_v42);
          let _1440_v43;
          _1440_v43 = _dafny.Seq.UnicodeFromString("mkn");
          (globalState).f11 = new BigNumber((((!((_this).f24)) ? (_1440_v43) : (_dafny.Seq.Concat(_1440_v43, _1440_v43)))).length);
          let _1441_v44;
          _1441_v44 = _dafny.Seq.of((_this).f24);
          let _1442_v45;
          let _nw195 = Array((new BigNumber(4)).toNumber());
          _nw195[(_dafny.ZERO).toNumber()] = _module.__default.fm37(_1441_v44, _dafny.Map.Empty.slice().updateUnsafe(_1422_cf0,_1423_v30), _dafny.Seq.UnicodeFromString("xqe"), globalState);
          _nw195[(_dafny.ONE).toNumber()] = p0;
          _nw195[(new BigNumber(2)).toNumber()] = p0;
          _nw195[(new BigNumber(3)).toNumber()] = p0;
          _1442_v45 = _nw195;
          let _1443_v46;
          let _nw196 = new _module.C4();
          _nw196.__ctor(_1421_cf1, _dafny.Seq.of(_1442_v45, _1442_v45, _1442_v45));
          _1443_v46 = _nw196;
          let _1444_v47;
          _1444_v47 = _dafny.Seq.of(_1378_v0);
          let _1445_v48;
          let _nw197 = Array((new BigNumber(20)).toNumber());
          _nw197[(_dafny.ZERO).toNumber()] = _1378_v0;
          _nw197[(_dafny.ONE).toNumber()] = (_1444_v47)[_module.__default.safeIndex(_1419_cf3, new BigNumber((_1444_v47).length))];
          _nw197[(new BigNumber(2)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(3)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(4)).toNumber()] = (_1444_v47)[_module.__default.safeIndex((_1425_v32)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_1425_v32).length))], new BigNumber((_1444_v47).length))];
          _nw197[(new BigNumber(5)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(6)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(7)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(8)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(9)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(10)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(11)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(12)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(13)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(14)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(15)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(16)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(17)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(18)).toNumber()] = _1378_v0;
          _nw197[(new BigNumber(19)).toNumber()] = _1378_v0;
          _1445_v48 = _nw197;
          let _index201 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1445_v48).length));
          (_1445_v48)[_index201] = _1378_v0;
          let _index202 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1445_v48).length));
          (_1445_v48)[_index202] = ((((_1423_v30) ? ((_this).f24) : (_1423_v30))) ? (_1378_v0) : (_1378_v0));
        }
        let _1446_v49;
        _1446_v49 = _dafny.Seq.UnicodeFromString("n");
        _1418_cf4 = (_1446_v49)[_module.__default.safeIndex(_1419_cf3, new BigNumber((_1446_v49).length))];
      }
      let _1447_v50;
      _1447_v50 = new BigNumber(-250);
      if (_module.__default.fm15((_dafny.MultiSet.fromElements(_1447_v50)).IsSubsetOf(_dafny.MultiSet.fromElements(_1447_v50)), globalState)) {
        let _1448_v51;
        let _init31 = ((_1449_p0) => function (_1450_i6) {
          return _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_this).f24, (_this).f24), _dafny.MultiSet.fromElements(_1449_p0), _dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f24)));
        })(p0);
        let _nw198 = Array((new BigNumber(8)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw198.length); _i0_31++) {
          _nw198[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1448_v51 = _nw198;
        let _1451_v52;
        _1451_v52 = _dafny.Seq.UnicodeFromString("cnkxhr");
        let _1452_v53;
        _1452_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(_1451_v52, (_this).f24, false, globalState),(_this).f24);
        let _1453_v54;
        _1453_v54 = _dafny.Set.fromElements((((_1452_v53).contains(false)) ? ((_1452_v53).get(false)) : ((_this).f24)), p0, (_this).f24, (_this).f24, (_this).f24);
        let _1454_v55;
        _1454_v55 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24);
        let _1455_v56;
        _1455_v56 = _dafny.MultiSet.fromElements(_module.__default.fm41(_1447_v50, _1451_v52, _1453_v54, globalState), _1454_v55, _1454_v55);
        let _index203 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1448_v51).length));
        (_1448_v51)[_index203] = _1455_v56;
        let _index204 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1448_v51).length));
        (_1448_v51)[_index204] = (_1455_v56).Difference(_1455_v56);
        let _1456_v57;
        let _init32 = ((_1457_p0) => function (_1458_i7) {
          return _1457_p0;
        })(p0);
        let _nw199 = Array((new BigNumber(3)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw199.length); _i0_32++) {
          _nw199[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1456_v57 = _nw199;
        let _index205 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length));
        (_1456_v57)[_index205] = !(p0);
        let _1459_v58;
        _1459_v58 = new _dafny.CodePoint('i'.codePointAt(0));
        let _1460_v59;
        let _nw200 = new _module.C1();
        _nw200.__ctor(_1447_v50, (_this).f22);
        _1460_v59 = _nw200;
        let _1461_v60;
        _1461_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1460_v59);
        let _1462_v61;
        _1462_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1459_v58,_1461_v60);
        let _index206 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length));
        (_1456_v57)[_index206] = ((((false) ? (false) : ((_this).f24))) ? (!((_this).f24) || (false)) : ((_1447_v50).isLessThan(new BigNumber((_1462_v61).length))));
        let _1463_v62;
        _1463_v62 = _module.D4.create_DC15(((_1460_v59).f30).multipliedBy(new BigNumber((_1451_v52).length)), (_1460_v59).f30);
        let _source22 = _1463_v62;
        if (_source22.is_DC13) {
          let _1464___mcc_h5 = (_source22).cf22;
          let _1465_cf22 = _1464___mcc_h5;
          let _index207 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length));
          (_1456_v57)[_index207] = (_this).f24;
          _1459_v58 = _1459_v58;
          _1460_v59 = _1460_v59;
          let _1466_v63;
          let _nw201 = Array((new BigNumber(6)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1466_v63 = _nw201;
          let _1467_v64;
          _1467_v64 = _dafny.MultiSet.fromElements(_module.D6.create_DC22((_1460_v59).f30));
          let _index208 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1466_v63).length));
          (_1466_v63)[_index208] = _1467_v64;
          let _index209 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1466_v63).length));
          (_1466_v63)[_index209] = _module.__default.fm46(globalState);
        } else if (_source22.is_DC14) {
          let _1468___mcc_h6 = (_source22).cf23;
          let _1469___mcc_h7 = (_source22).cf24;
          let _1470___mcc_h8 = (_source22).cf25;
          let _1471_cf25 = _1470___mcc_h8;
          let _1472_cf24 = _1469___mcc_h7;
          let _1473_cf23 = _1468___mcc_h6;
          _1473_cf23 = _1473_cf23;
          let _index210 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length));
          (_1456_v57)[_index210] = p0;
          let _1474_v65;
          let _nw202 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1474_v65 = _nw202;
          let _rhs235 = _1474_v65;
          let _rhs236 = (_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))];
          let _lhs201 = globalState;
          _1474_v65 = _rhs235;
          _lhs201.f10 = _rhs236;
          let _index211 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1474_v65).length));
          (_1474_v65)[_index211] = _module.__default.safeModuloInt((_1460_v59).f30, (_dafny.ZERO).minus((_1460_v59).f30));
          let _index212 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1474_v65).length));
          (_1474_v65)[_index212] = (_1460_v59).f30;
        } else if (_source22.is_DC15) {
          let _1475___mcc_h9 = (_source22).cf26;
          let _1476___mcc_h10 = (_source22).cf27;
          let _1477_cf27 = _1476___mcc_h10;
          let _1478_cf26 = _1475___mcc_h9;
          let _1479_v66;
          let _1480_v67;
          let _out58;
          let _out59;
          let _outcollector18 = (_this).m1(globalState);
          _out58 = _outcollector18[0];
          _out59 = _outcollector18[1];
          _1479_v66 = _out58;
          _1480_v67 = _out59;
          let _1481_v68;
          let _nw203 = new _module.C2();
          _nw203.__ctor((_1460_v59).f30, (_this).f22);
          _1481_v68 = _nw203;
          (globalState).f6 = _1447_v50;
          let _1482_v69;
          let _init33 = ((_1483_v59) => function (_1484_i8) {
            return (_1484_i8).minus((_1483_v59).f30);
          })(_1460_v59);
          let _nw204 = Array((new BigNumber(14)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw204.length); _i0_33++) {
            _nw204[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1482_v69 = _nw204;
          let _1485_v70;
          _1485_v70 = _dafny.Seq.of(new BigNumber(-441));
          let _index213 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_1482_v69).length));
          (_1482_v69)[_index213] = _module.__default.safeDivisionInt(_1447_v50, (_1485_v70)[_module.__default.safeIndex((_1481_v68).f29, new BigNumber((_1485_v70).length))]);
          let _index214 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_1482_v69).length));
          (_1482_v69)[_index214] = (_1485_v70)[_module.__default.safeIndex((_1460_v59).f30, new BigNumber((_1485_v70).length))];
        } else if (_source22.is_DC12) {
          let _1486___mcc_h11 = (_source22).cf21;
          let _1487_cf21 = _1486___mcc_h11;
          (globalState).f16 = (_1460_v59).f30;
          let _1488_v71;
          let _nw205 = new _module.C3();
          _nw205.__ctor((_this).f22);
          _1488_v71 = _nw205;
          let _1489_v72;
          let _nw206 = Array((new BigNumber(29)).toNumber()).fill([]);
          _1489_v72 = _nw206;
          let _1490_v73;
          let _nw207 = Array((new BigNumber(23)).toNumber()).fill([]);
          _1490_v73 = _nw207;
          let _index215 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_1489_v72).length));
          (_1489_v72)[_index215] = _1490_v73;
          let _index216 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_1489_v72).length));
          (_1489_v72)[_index216] = _1490_v73;
          let _1491_v74;
          _1491_v74 = _module.D4.create_DC14(_1459_v58, _1456_v57, _dafny.Seq.Create(_module.__default.abs(new BigNumber(778)), ((_1492_v58) => function (_1493_i9) {
  return _1492_v58;
})(_1459_v58)));
          let _pat_let_tv56 = _1451_v52;
          let _rhs237 = _1447_v50;
          let _rhs238 = (function (_pat_let26_0) {
            return function (_1494_dt__update__tmp_h0) {
              return function (_pat_let27_0) {
                return function (_1495_dt__update_hcf25_h0) {
                  return _module.D4.create_DC14((_1494_dt__update__tmp_h0).dtor_cf23, (_1494_dt__update__tmp_h0).dtor_cf24, _1495_dt__update_hcf25_h0);
                }(_pat_let27_0);
              }(_pat_let_tv56);
            }(_pat_let26_0);
          }(_1491_v74)).dtor_cf24;
          let _rhs239 = (_1460_v59).f30;
          let _lhs202 = globalState;
          r3 = _rhs237;
          r1 = _rhs238;
          _lhs202.f11 = _rhs239;
        } else {
          let _1496___mcc_h12 = (_source22).cf28;
          let _1497_cf28 = _1496___mcc_h12;
          (globalState).f6 = (_dafny.ZERO).minus((_1460_v59).f30);
          let _1498_v75;
          _1498_v75 = _module.D14.create_DC47(_1454_v55);
          (globalState).f0 = ((_1498_v75).dtor_cf77).IsSubsetOf((_1454_v55).Difference(_1454_v55));
          let _1499_v76;
          let _nw208 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1499_v76 = _nw208;
          let _1500_v77;
          _1500_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1499_v76,(_1460_v59).f30);
          let _1501_v78;
          _1501_v78 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1447_v50);
          let _1502_v79;
          _1502_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v50,!((_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))]));
          let _1503_v80;
          _1503_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1502_v79,!((_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))]));
          let _1504_v81;
          _1504_v81 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1454_v55);
          let _1505_v82;
          _1505_v82 = _dafny.Seq.of(_1447_v50, (_1460_v59).f30, (_1460_v59).f30, (_1460_v59).f30, new BigNumber((_1504_v81).length));
          let _1506_v83;
          _1506_v83 = _dafny.Seq.of(_1502_v79);
          let _1507_v84;
          let _nw209 = Array((new BigNumber(24)).toNumber());
          _nw209[(_dafny.ZERO).toNumber()] = (_1460_v59).f30;
          _nw209[(_dafny.ONE).toNumber()] = (_1460_v59).fm35((_dafny.ZERO).minus((_1460_v59).f30), new BigNumber((_1500_v77).length), p0, new BigNumber(665), globalState);
          _nw209[(new BigNumber(2)).toNumber()] = new BigNumber((_1501_v78).length);
          _nw209[(new BigNumber(3)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(4)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(5)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(6)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(7)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(8)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(9)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(10)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(11)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(12)).toNumber()] = _1447_v50;
          _nw209[(new BigNumber(13)).toNumber()] = new BigNumber(-563);
          _nw209[(new BigNumber(14)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(15)).toNumber()] = new BigNumber((_1503_v80).length);
          _nw209[(new BigNumber(16)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(17)).toNumber()] = (_1505_v82)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_1460_v59).f30)), new BigNumber((_1505_v82).length))];
          _nw209[(new BigNumber(18)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(19)).toNumber()] = (_1460_v59).f30;
          _nw209[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_1505_v82)).cardinality());
          _nw209[(new BigNumber(21)).toNumber()] = new BigNumber((_1506_v83).length);
          _nw209[(new BigNumber(22)).toNumber()] = new BigNumber(624);
          _nw209[(new BigNumber(23)).toNumber()] = (_1460_v59).f30;
          _1507_v84 = _nw209;
          let _1508_v85;
          _1508_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1507_v84,(_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))]);
          let _1509_v86;
          _1509_v86 = _module.D10.create_DC34(_1459_v58, (_this).f24, (_module.D8.create_DC29(!(p0), p0, new _dafny.CodePoint('l'.codePointAt(0)))).dtor_cf48, (_this).f24, true);
          let _1510_v87;
          let _nw210 = new _module.C4();
          _nw210.__ctor((_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))], (_this).f22);
          _1510_v87 = _nw210;
          let _1511_v88;
          _1511_v88 = _dafny.Set.fromElements(_1499_v76);
          let _1512_v89;
          _1512_v89 = _module.D13.create_DC46((_this).f24, _1510_v87, _1447_v50, _1511_v88);
          let _1513_v90;
          _1513_v90 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1507_v84,(_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))]), _1508_v85);
          let _1514_v91;
          let _nw211 = Array((new BigNumber(17)).toNumber());
          _nw211[(_dafny.ZERO).toNumber()] = _1508_v85;
          _nw211[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1499_v76,false);
          _nw211[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1499_v76,_module.__default.fm15((_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))], globalState))).Merge(_1508_v85);
          _nw211[(new BigNumber(3)).toNumber()] = _1508_v85;
          _nw211[(new BigNumber(4)).toNumber()] = _1508_v85;
          _nw211[(new BigNumber(5)).toNumber()] = _1508_v85;
          _nw211[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1499_v76,(_1509_v86).dtor_cf57)).Merge(_1508_v85);
          _nw211[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1499_v76,(_this).f24);
          _nw211[(new BigNumber(8)).toNumber()] = _1508_v85;
          _nw211[(new BigNumber(9)).toNumber()] = _1508_v85;
          _nw211[(new BigNumber(10)).toNumber()] = _1508_v85;
          _nw211[(new BigNumber(11)).toNumber()] = (((_1512_v89).dtor_cf73) ? (_dafny.Map.Empty.slice().updateUnsafe(_1507_v84,(_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))])) : ((_1513_v90)[_module.__default.safeIndex(new BigNumber((_1505_v82).length), new BigNumber((_1513_v90).length))]));
          _nw211[(new BigNumber(12)).toNumber()] = (_1508_v85).Merge(_1508_v85);
          _nw211[(new BigNumber(13)).toNumber()] = _1508_v85;
          _nw211[(new BigNumber(14)).toNumber()] = _1508_v85;
          _nw211[(new BigNumber(15)).toNumber()] = _1508_v85;
          _nw211[(new BigNumber(16)).toNumber()] = (_1508_v85).update(_1499_v76, p0);
          _1514_v91 = _nw211;
          let _index217 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1514_v91).length));
          (_1514_v91)[_index217] = _dafny.Map.Empty.slice().updateUnsafe(_1507_v84,(_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))]);
          let _index218 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1514_v91).length));
          (_1514_v91)[_index218] = _1508_v85;
          let _1515_v92;
          _1515_v92 = _dafny.Seq.of(_1499_v76);
          _1499_v76 = (_1515_v92)[_module.__default.safeIndex(_1447_v50, new BigNumber((_1515_v92).length))];
        }
        let _1516_v93;
        _1516_v93 = _dafny.Seq.of((_1460_v59).f30);
        _1452_v53 = (_1452_v53).update(!(((_1516_v93)[_module.__default.safeIndex(new BigNumber((_1516_v93).length), new BigNumber((_1516_v93).length))]).isLessThanOrEqualTo((_1460_v59).f30)), (_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))]);
        let _1517_v94;
        _1517_v94 = _dafny.Map.Empty.slice().updateUnsafe((_1456_v57)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_1456_v57).length))],(_1460_v59).f30);
        let _1518_v95;
        _1518_v95 = _module.D9.create_DC30((_module.D9.create_DC30(_1517_v94)).dtor_cf50);
        _1518_v95 = _1518_v95;
      } else {
        let _1519_v96;
        _1519_v96 = _dafny.Seq.of(!(_1447_v50).isEqualTo(_1447_v50));
        if ((_1519_v96)[_module.__default.safeIndex(_1447_v50, new BigNumber((_1519_v96).length))]) {
          let _init34 = function (_1520_i10) {
            return (_this).f24;
          };
          let _nw212 = Array((new BigNumber(16)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw212.length); _i0_34++) {
            _nw212[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          r1 = _nw212;
          let _1521_v97;
          let _nw213 = new _module.C1();
          _nw213.__ctor(_1447_v50, (_this).f22);
          _1521_v97 = _nw213;
          let _1522_v98;
          let _init35 = ((_1523_v97) => function (_1524_i11) {
            return (_1524_i11).plus((_1523_v97).f30);
          })(_1521_v97);
          let _nw214 = Array((new BigNumber(24)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw214.length); _i0_35++) {
            _nw214[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1522_v98 = _nw214;
          let _index219 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1522_v98).length));
          (_1522_v98)[_index219] = (_1521_v97).f30;
          let _index220 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1522_v98).length));
          (_1522_v98)[_index220] = _module.__default.fm7((_1521_v97).f30, globalState);
          let _1525_v99;
          _1525_v99 = new _dafny.CodePoint('y'.codePointAt(0));
          let _index221 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1522_v98).length));
          let _index222 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1522_v98).length));
          let _rhs240 = (_1521_v97).f30;
          let _rhs241 = _module.__default.fm38(true, _1447_v50, globalState);
          let _rhs242 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(_1447_v50, new BigNumber((_1519_v96).length)), (_1447_v50).multipliedBy((_1521_v97).f30));
          let _rhs243 = ((_1521_v97).f30).multipliedBy(_1447_v50);
          let _rhs244 = (_1522_v98)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_1522_v98).length))];
          let _lhs203 = _1522_v98;
          let _lhs204 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1522_v98).length));
          let _lhs205 = globalState;
          let _lhs206 = _1522_v98;
          let _lhs207 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1522_v98).length));
          _lhs203[_lhs204] = _rhs240;
          _1525_v99 = _rhs241;
          _lhs205.f6 = _rhs242;
          r3 = _rhs243;
          _lhs206[_lhs207] = _rhs244;
          let _index223 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1522_v98).length));
          (_1522_v98)[_index223] = _1447_v50;
        } else {
          let _1526_v100;
          _1526_v100 = _module.D2.create_DC7(_dafny.Seq.of(_1447_v50), (_this).f24, !((_this).f24), _1447_v50);
          let _1527_v101;
          _1527_v101 = _dafny.Set.fromElements(_1526_v100, _1526_v100, _1526_v100);
          (globalState).f10 = (_1527_v101).equals(_1527_v101);
          let _1528_v102;
          let _init36 = ((_1529_p0) => function (_1530_i12) {
            return _1529_p0;
          })(p0);
          let _nw215 = Array((new BigNumber(7)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw215.length); _i0_36++) {
            _nw215[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1528_v102 = _nw215;
          let _1531_v103;
          let _nw216 = new _module.C2();
          _nw216.__ctor(_1447_v50, _dafny.Seq.of(_1528_v102));
          _1531_v103 = _nw216;
          let _1532_v104;
          _1532_v104 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(584)), ((_1533_v50) => function (_1534_i13) {
            return _1533_v50;
          })(_1447_v50))).length));
          let _1535_v105;
          let _nw217 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _1535_v105 = _nw217;
          let _1536_v106;
          _1536_v106 = _module.D13.create_DC46(p0, _1531_v103, new BigNumber((_1532_v104).length), _dafny.Set.fromElements(_1535_v105));
          let _1537_v107;
          _1537_v107 = _module.D13.create_DC46((_this).f24, (_1536_v106).dtor_cf74, _1447_v50, _dafny.Set.fromElements(_1535_v105, _1535_v105, _1535_v105, _1535_v105));
          let _1538_v108;
          _1538_v108 = _dafny.MultiSet.fromElements(_1447_v50);
          let _1539_v111;
          _1539_v111 = _dafny.Set.fromElements(function () {
            let _coll61 = new _dafny.Set();
            for (const _compr_61 of (function () {
              let _coll62 = new _dafny.Set();
              for (const _compr_62 of (_1538_v108).Elements) {
                let _1540_v109 = _compr_62;
                if ((_1538_v108).contains(_1540_v109)) {
                  _coll62.add((_1540_v109).minus(new BigNumber((_dafny.Seq.UnicodeFromString("eelqwwl")).length)));
                }
              }
              return _coll62;
            }()).Elements) {
              let _1541_v110 = _compr_61;
              if ((function () {
                let _coll63 = new _dafny.Set();
                for (const _compr_63 of (_1538_v108).Elements) {
                  let _1542_v109 = _compr_63;
                  if ((_1538_v108).contains(_1542_v109)) {
                    _coll63.add((_1542_v109).minus(new BigNumber((_dafny.Seq.UnicodeFromString("eelqwwl")).length)));
                  }
                }
                return _coll63;
              }()).contains(_1541_v110)) {
                _coll61.add((_1541_v110).multipliedBy(new BigNumber(274)));
              }
            }
            return _coll61;
          }());
          let _1543_v112;
          _1543_v112 = _dafny.Set.fromElements(_1535_v105, _1535_v105, _1535_v105);
          let _1544_v113;
          _1544_v113 = _dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC46(p0, _1531_v103, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1447_v50)).length), _1543_v112),new BigNumber(718));
          let _1545_v114;
          _1545_v114 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1537_v107,new BigNumber((_1539_v111).length)), (_1544_v113).update(_1537_v107, new BigNumber(-263)));
          (globalState).f11 = new BigNumber((_1545_v114).length);
          let _1546_v115;
          _1546_v115 = _dafny.Seq.UnicodeFromString("phpv");
          _1546_v115 = _1546_v115;
          let _1547_v116;
          _1547_v116 = new _dafny.CodePoint('y'.codePointAt(0));
          let _1548_v117;
          _1548_v117 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v50,_1547_v116);
          let _1549_v119;
          _1549_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v50,!(_1548_v117).equals((function () {
            let _coll64 = new _dafny.Map();
            for (const _compr_64 of _dafny.IntegerRange(new BigNumber(667), new BigNumber(816))) {
              let _1550_v118 = _compr_64;
              if (((new BigNumber(667)).isLessThanOrEqualTo(_1550_v118)) && ((_1550_v118).isLessThan(new BigNumber(816)))) {
                _coll64.push([(_1550_v118).multipliedBy(new BigNumber((_1532_v104).length)),_1547_v116]);
              }
            }
            return _coll64;
          }()).update(_1447_v50, new _dafny.CodePoint('n'.codePointAt(0)))));
          let _1551_v120;
          _1551_v120 = _dafny.Seq.of(_1447_v50);
          let _1552_v121;
          _1552_v121 = _module.D13.create_DC45(new BigNumber((_1551_v120).length));
          let _pat_let_tv57 = _1447_v50;
          _1549_v119 = (_1549_v119).update(_1447_v50, (_1447_v50).isEqualTo((function (_pat_let28_0) {
            return function (_1553_dt__update__tmp_h1) {
              return function (_pat_let29_0) {
                return function (_1554_dt__update_hcf72_h0) {
                  return _module.D13.create_DC45(_1554_dt__update_hcf72_h0);
                }(_pat_let29_0);
              }(_pat_let_tv57);
            }(_pat_let28_0);
          }(_1552_v121)).dtor_cf72));
          (globalState).f16 = _1447_v50;
        }
        let _1555_v122;
        _1555_v122 = _module.D10.create_DC35((_dafny.ZERO).minus(_1447_v50), _1447_v50);
        let _source23 = _1555_v122;
        if (_source23.is_DC34) {
          let _1556___mcc_h13 = (_source23).cf55;
          let _1557___mcc_h14 = (_source23).cf56;
          let _1558___mcc_h15 = (_source23).cf57;
          let _1559___mcc_h16 = (_source23).cf58;
          let _1560___mcc_h17 = (_source23).cf59;
          let _1561_cf59 = _1560___mcc_h17;
          let _1562_cf58 = _1559___mcc_h16;
          let _1563_cf57 = _1558___mcc_h15;
          let _1564_cf56 = _1557___mcc_h14;
          let _1565_cf55 = _1556___mcc_h13;
          let _1566_v123;
          _1566_v123 = _dafny.Seq.UnicodeFromString("n");
          let _1567_v124;
          _1567_v124 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_dafny.Seq.update(_1519_v96, _module.__default.safeIndex(_1447_v50, new BigNumber((_1519_v96).length)), _1561_cf59), _1519_v96),(_1447_v50).plus((_dafny.ZERO).minus(new BigNumber((_1566_v123).length))));
          let _rhs245 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_1447_v50), _1447_v50))));
          let _rhs246 = !(_dafny.Seq.IsPrefixOf(_1566_v123, _1566_v123));
          let _rhs247 = ((_1567_v124).Merge(_1567_v124)).Merge(_1567_v124);
          let _lhs208 = globalState;
          _lhs208.f8 = _rhs245;
          _1562_cf58 = _rhs246;
          _1567_v124 = _rhs247;
          let _1568_v125;
          let _init37 = ((_1569_v50) => function (_1570_i14) {
            return _module.__default.safeModuloInt(_1570_i14, _1569_v50);
          })(_1447_v50);
          let _nw218 = Array((new BigNumber(7)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw218.length); _i0_37++) {
            _nw218[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1568_v125 = _nw218;
          let _nw219 = Array((_dafny.ONE).toNumber());
          _nw219[(_dafny.ZERO).toNumber()] = new BigNumber(-204);
          _1568_v125 = _nw219;
          (globalState).f16 = _1447_v50;
          r3 = _1447_v50;
        } else if (_source23.is_DC35) {
          let _1571___mcc_h18 = (_source23).cf60;
          let _1572___mcc_h19 = (_source23).cf61;
          let _1573_cf61 = _1572___mcc_h19;
          let _1574_cf60 = _1571___mcc_h18;
          let _1575_v126;
          _1575_v126 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1576_v127;
          _1576_v127 = _dafny.Seq.of(_1447_v50);
          let _1577_v128;
          _1577_v128 = _module.D0.create_DC1(_1573_cf61, p0, _module.__default.fm47(_1575_v126, false, globalState), new BigNumber((_1576_v127).length), _1575_v126);
          let _1578_v129;
          _1578_v129 = _module.D4.create_DC13(_1577_v128);
          let _1579_v130;
          _1579_v130 = _dafny.Seq.of(_1578_v129);
          let _1580_v131;
          _1580_v131 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1579_v130).length),p0);
          let _1581_v132;
          _1581_v132 = _dafny.Seq.UnicodeFromString("kmopkrvet");
          let _rhs248 = (!(p0)) || (_module.__default.fm15((_this).f24, globalState));
          let _rhs249 = _module.__default.fm37(_dafny.Seq.of(p0, p0, p0), _1580_v131, _1581_v132, globalState);
          let _rhs250 = _1573_cf61;
          let _lhs209 = globalState;
          let _lhs210 = globalState;
          r0 = _rhs248;
          _lhs209.f0 = _rhs249;
          _lhs210.f16 = _rhs250;
          let _1582_v133;
          _1582_v133 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1581_v132).length));
          let _1583_v134;
          _1583_v134 = _module.D9.create_DC30(_1582_v133);
          let _1584_v135;
          _1584_v135 = _dafny.Seq.of(_1582_v133, _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_1576_v127).length)));
          let _pat_let_tv58 = _1584_v135;
          let _pat_let_tv59 = _1573_cf61;
          let _pat_let_tv60 = _1584_v135;
          let _1585_v136;
          _1585_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1581_v132,function (_pat_let30_0) {
            return function (_1586_dt__update__tmp_h2) {
              return function (_pat_let31_0) {
                return function (_1587_dt__update_hcf50_h0) {
                  return _module.D9.create_DC30(_1587_dt__update_hcf50_h0);
                }(_pat_let31_0);
              }((_pat_let_tv58)[_module.__default.safeIndex(_pat_let_tv59, new BigNumber((_pat_let_tv60).length))]);
            }(_pat_let30_0);
          }(_1583_v134));
          _1585_v136 = _1585_v136;
          let _1588_v137;
          let _nw220 = Array((new BigNumber(11)).toNumber()).fill(false);
          _1588_v137 = _nw220;
          r1 = _1588_v137;
          let _1589_v138;
          let _init38 = ((_1590_cf61) => function (_1591_i15) {
            return (_1591_i15).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1590_cf61,new BigNumber(43))).length)));
          })(_1573_cf61);
          let _nw221 = Array((new BigNumber(3)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw221.length); _i0_38++) {
            _nw221[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1589_v138 = _nw221;
          let _index224 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_1589_v138).length));
          (_1589_v138)[_index224] = new BigNumber(-623);
          let _1592_v139;
          _1592_v139 = _dafny.MultiSet.fromElements(_1447_v50);
          let _1593_v140;
          _1593_v140 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24, (_this).f24, p0, p0);
          let _1594_v141;
          _1594_v141 = _dafny.Set.fromElements(new BigNumber(726), new BigNumber(-786));
          let _1595_v142;
          _1595_v142 = _dafny.Seq.of(_module.__default.fm23(_1447_v50, _1447_v50, p0, p0, globalState), _1594_v141);
          let _index225 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_1589_v138).length));
          let _rhs251 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1581_v132, _dafny.Seq.UnicodeFromString("q")), _dafny.Seq.Concat(_1581_v132, _1581_v132));
          let _rhs252 = (_this).f24;
          let _rhs253 = (_this).fm0(_1592_v139, ((p0) ? (_1447_v50) : ((_dafny.ZERO).minus(new BigNumber((_1593_v140).cardinality())))), _1595_v142, (_this).f24, globalState);
          let _rhs254 = (_1576_v127)[_module.__default.safeIndex((_dafny.ZERO).minus(_1573_cf61), new BigNumber((_1576_v127).length))];
          let _lhs211 = globalState;
          let _lhs212 = _1589_v138;
          let _lhs213 = _module.__default.safeIndex(new BigNumber(516), new BigNumber((_1589_v138).length));
          _1581_v132 = _rhs251;
          _lhs211.f10 = _rhs252;
          r3 = _rhs253;
          _lhs212[_lhs213] = _rhs254;
        } else {
          let _1596___mcc_h20 = (_source23).cf54;
          let _1597_cf54 = _1596___mcc_h20;
          let _1598_v143;
          _1598_v143 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1599_v144;
          _1599_v144 = _dafny.Map.Empty.slice().updateUnsafe(_1381_v1,_1598_v143);
          let _1600_v145;
          _1600_v145 = _dafny.Seq.of(_1599_v144);
          let _1601_v146;
          _1601_v146 = _dafny.Seq.of(_1447_v50, new BigNumber(306));
          _1599_v144 = (_1600_v145)[_module.__default.safeIndex((_1601_v146)[_module.__default.safeIndex((_dafny.ZERO).minus(_1447_v50), new BigNumber((_1601_v146).length))], new BigNumber((_1600_v145).length))];
          (globalState).f11 = _1447_v50;
          let _1602_v147;
          _1602_v147 = _dafny.MultiSet.fromElements(_1447_v50);
          (globalState).f0 = (_dafny.MultiSet.fromElements(_1447_v50)).IsProperSubsetOf(_1602_v147);
          let _1603_v148;
          let _nw222 = Array((new BigNumber(2)).toNumber()).fill(false);
          _1603_v148 = _nw222;
          let _1604_v149;
          _1604_v149 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1605_v150;
          _1605_v150 = _dafny.Map.Empty.slice().updateUnsafe(_1603_v148,_1604_v149);
          _1605_v150 = (_1605_v150).update(_1603_v148, _1604_v149);
        }
        let _1606_v151;
        let _nw223 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
        _1606_v151 = _nw223;
        let _1607_v152;
        let _nw224 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
        _1607_v152 = _nw224;
        let _1608_v153;
        _1608_v153 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v50,p0);
        let _1609_v154;
        _1609_v154 = _dafny.Seq.UnicodeFromString("lntn");
        let _1610_v155;
        _1610_v155 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v50,_module.__default.fm37(_1519_v96, _1608_v153, _1609_v154, globalState));
        let _1611_v156;
        _1611_v156 = _dafny.Seq.of(_1610_v155);
        let _rhs255 = !(p0);
        let _rhs256 = _1606_v151;
        let _rhs257 = _1607_v152;
        let _rhs258 = _1611_v156;
        let _rhs259 = p0;
        let _lhs214 = globalState;
        let _lhs215 = globalState;
        _lhs214.f10 = _rhs255;
        _1606_v151 = _rhs256;
        _1607_v152 = _rhs257;
        _1611_v156 = _rhs258;
        _lhs215.f10 = _rhs259;
        let _1612_v157;
        _1612_v157 = _dafny.MultiSet.fromElements(_1609_v154);
        let _1613_v158;
        _1613_v158 = _module.D10.create_DC34(_module.__default.fm38(p0, _1447_v50, globalState), p0, p0, !((_1612_v157).IsSubsetOf(_1612_v157)), p0);
        let _source24 = _1613_v158;
        if (_source24.is_DC34) {
          let _1614___mcc_h21 = (_source24).cf55;
          let _1615___mcc_h22 = (_source24).cf56;
          let _1616___mcc_h23 = (_source24).cf57;
          let _1617___mcc_h24 = (_source24).cf58;
          let _1618___mcc_h25 = (_source24).cf59;
          let _1619_cf59 = _1618___mcc_h25;
          let _1620_cf58 = _1617___mcc_h24;
          let _1621_cf57 = _1616___mcc_h23;
          let _1622_cf56 = _1615___mcc_h22;
          let _1623_cf55 = _1614___mcc_h21;
          let _1624_v159;
          let _nw225 = new _module.C0();
          _nw225.__ctor();
          _1624_v159 = _nw225;
          let _1625_v160;
          let _nw226 = new _module.C3();
          _nw226.__ctor((_this).f22);
          _1625_v160 = _nw226;
          let _1626_v161;
          _1626_v161 = _dafny.MultiSet.fromElements(_1625_v160);
          let _1627_v162;
          _1627_v162 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v50,_1626_v161);
          let _1628_v163;
          let _nw227 = new _module.C1();
          _nw227.__ctor(new BigNumber((_1627_v162).length), _dafny.Seq.Concat((_this).f22, (_this).f22));
          _1628_v163 = _nw227;
          let _1629_v164;
          _1629_v164 = _dafny.Map.Empty.slice().updateUnsafe(_1619_cf59,new BigNumber(-861));
          let _1630_v165;
          _1630_v165 = _dafny.Map.Empty.slice().updateUnsafe(_1619_cf59,_1625_v160);
          let _1631_v166;
          _1631_v166 = _dafny.Set.fromElements(_1630_v165, _1630_v165, _1630_v165, (_1630_v165).update(p0, _1625_v160), (_1630_v165).update(_1619_cf59, _1625_v160));
          let _1632_v167;
          _1632_v167 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_1619_cf59));
          let _rhs260 = (_dafny.ZERO).minus(_1447_v50);
          let _rhs261 = (_1628_v163).f30;
          let _rhs262 = _1629_v164;
          let _rhs263 = _dafny.Set.fromElements((_1630_v165).Merge(_1630_v165), _1630_v165, _1630_v165);
          let _rhs264 = _1632_v167;
          let _lhs216 = globalState;
          _1447_v50 = _rhs260;
          _lhs216.f16 = _rhs261;
          _1629_v164 = _rhs262;
          _1631_v166 = _rhs263;
          _1632_v167 = _rhs264;
          (globalState).f16 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1628_v163).fm34(true, globalState)));
        } else if (_source24.is_DC35) {
          let _1633___mcc_h26 = (_source24).cf60;
          let _1634___mcc_h27 = (_source24).cf61;
          let _1635_cf61 = _1634___mcc_h27;
          let _1636_cf60 = _1633___mcc_h26;
          let _1637_v168;
          let _nw228 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _1637_v168 = _nw228;
          let _1638_v169;
          let _nw229 = Array((new BigNumber(17)).toNumber()).fill(false);
          _1638_v169 = _nw229;
          let _index226 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1638_v169).length));
          (_1638_v169)[_index226] = (_this).f24;
          let _1639_v170;
          _1639_v170 = _dafny.Set.fromElements(p0);
          let _1640_v171;
          _1640_v171 = _dafny.MultiSet.fromElements(_1639_v170, _1639_v170);
          let _index227 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1638_v169).length));
          let _rhs265 = _1637_v168;
          let _rhs266 = _1447_v50;
          let _rhs267 = (_1640_v171).IsSubsetOf(_1640_v171);
          let _rhs268 = !((_this).f24);
          let _lhs217 = globalState;
          let _lhs218 = globalState;
          let _lhs219 = _1638_v169;
          let _lhs220 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1638_v169).length));
          _1637_v168 = _rhs265;
          _lhs217.f16 = _rhs266;
          _lhs218.f10 = _rhs267;
          _lhs219[_lhs220] = _rhs268;
          (globalState).f5 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1519_v96, _1519_v96), _dafny.Seq.Concat(_dafny.Seq.of((_1638_v169)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_1638_v169).length))], p0, (_1638_v169)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_1638_v169).length))], (_this).f24, (_this).f24), _1519_v96));
          let _1641_v172;
          _1641_v172 = _dafny.Seq.of(_dafny.Seq.of(_module.__default.fm15((_this).f24, globalState)));
          let _1642_v173;
          _1642_v173 = _dafny.MultiSet.fromElements(p0);
          _1641_v172 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_1641_v172, _module.__default.fm48(globalState)), _dafny.Seq.of(_1519_v96)), _module.__default.safeIndex((_1447_v50).multipliedBy(new BigNumber((_1642_v173).cardinality())), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1641_v172, _module.__default.fm48(globalState)), _dafny.Seq.of(_1519_v96))).length)), _1519_v96);
          let _1643_v174;
          _1643_v174 = _dafny.Seq.of(_1635_cf61);
          _1643_v174 = _1643_v174;
        } else {
          let _1644___mcc_h28 = (_source24).cf54;
          let _1645_cf54 = _1644___mcc_h28;
          let _rhs269 = !(((p0) ? (p0) : (p0))) || ((_this).fm3(_1609_v154, p0, (_1519_v96)[_module.__default.safeIndex(new BigNumber(124), new BigNumber((_1519_v96).length))], globalState));
          let _rhs270 = _dafny.Seq.IsPrefixOf(_1519_v96, _1519_v96);
          let _lhs221 = globalState;
          r0 = _rhs269;
          _lhs221.f10 = _rhs270;
          let _1646_v176;
          let _init39 = ((_1647_v50) => function (_1648_i16) {
            return (_module.D5.create_DC19(function () {
  let _coll65 = new _dafny.Map();
  for (const _compr_65 of _dafny.IntegerRange(new BigNumber(205), new BigNumber(829))) {
    let _1649_v175 = _compr_65;
    if (((new BigNumber(205)).isLessThanOrEqualTo(_1649_v175)) && ((_1649_v175).isLessThan(new BigNumber(829)))) {
      _coll65.push([(_1649_v175).multipliedBy(_1647_v50),_1647_v50]);
    }
  }
  return _coll65;
}(), true)).dtor_cf33;
          })(_1447_v50);
          let _nw230 = Array((new BigNumber(17)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw230.length); _i0_39++) {
            _nw230[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1646_v176 = _nw230;
          let _1650_v177;
          _1650_v177 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v50,_1646_v176);
          let _1651_v178;
          _1651_v178 = _dafny.Seq.of(new BigNumber((_1650_v177).length));
          let _1652_v179;
          _1652_v179 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1447_v50);
          let _1653_v180;
          _1653_v180 = _module.D2.create_DC7(_1651_v178, (_this).f24, true, new BigNumber((_1652_v179).length));
          let _index228 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1646_v176).length));
          (_1646_v176)[_index228] = (_1653_v180).dtor_cf11;
          let _index229 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1646_v176).length));
          (_1646_v176)[_index229] = (_1519_v96)[_module.__default.safeIndex(new BigNumber((((false) ? (_dafny.Seq.UnicodeFromString("on")) : (_1609_v154))).length), new BigNumber((_1519_v96).length))];
          _1609_v154 = _dafny.Seq.UnicodeFromString("hrcsc");
          let _1654_v181;
          _1654_v181 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1609_v154).length),new BigNumber(-637));
          let _1655_v182;
          let _nw231 = Array((new BigNumber(20)).toNumber());
          _nw231[(_dafny.ZERO).toNumber()] = _1447_v50;
          _nw231[(_dafny.ONE).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(2)).toNumber()] = new BigNumber((_1609_v154).length);
          _nw231[(new BigNumber(3)).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(4)).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(5)).toNumber()] = new BigNumber((_1609_v154).length);
          _nw231[(new BigNumber(6)).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(7)).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(8)).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(9)).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(10)).toNumber()] = new BigNumber(605);
          _nw231[(new BigNumber(11)).toNumber()] = new BigNumber(413);
          _nw231[(new BigNumber(12)).toNumber()] = (((_1654_v181).contains(_module.__default.fm7(_1447_v50, globalState))) ? ((_1654_v181).get(_module.__default.fm7(_1447_v50, globalState))) : (_1447_v50));
          _nw231[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("qbq")).length);
          _nw231[(new BigNumber(14)).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(15)).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(16)).toNumber()] = _1447_v50;
          _nw231[(new BigNumber(17)).toNumber()] = new BigNumber((_1609_v154).length);
          _nw231[(new BigNumber(18)).toNumber()] = new BigNumber(402);
          _nw231[(new BigNumber(19)).toNumber()] = new BigNumber(-262);
          _1655_v182 = _nw231;
          let _1656_v183;
          _1656_v183 = _dafny.Seq.of(_1655_v182);
          let _1657_v184;
          let _nw232 = Array((new BigNumber(29)).toNumber());
          _nw232[(_dafny.ZERO).toNumber()] = (_1656_v183)[_module.__default.safeIndex((_dafny.ZERO).minus(_1447_v50), new BigNumber((_1656_v183).length))];
          _nw232[(_dafny.ONE).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(2)).toNumber()] = (((_1646_v176)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1646_v176).length))]) ? (_1655_v182) : (_1655_v182));
          _nw232[(new BigNumber(3)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(4)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(5)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(6)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(7)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(8)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(9)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(10)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(11)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(12)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(13)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(14)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(15)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(16)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(17)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(18)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(19)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(20)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(21)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(22)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(23)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(24)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(25)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(26)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(27)).toNumber()] = _1655_v182;
          _nw232[(new BigNumber(28)).toNumber()] = _1655_v182;
          _1657_v184 = _nw232;
          let _index230 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_1657_v184).length));
          (_1657_v184)[_index230] = _1655_v182;
          let _index231 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_1657_v184).length));
          (_1657_v184)[_index231] = _1655_v182;
        }
        let _1658_v185;
        _1658_v185 = new _dafny.CodePoint('h'.codePointAt(0));
        let _1659_v186;
        _1659_v186 = _dafny.Map.Empty.slice().updateUnsafe(_1658_v185,p0);
        let _rhs271 = ((_module.__default.fm49(p0, globalState)).update(_1658_v185, false)).Merge(_1659_v186);
        let _rhs272 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1519_v96, _1519_v96), _1519_v96);
        let _lhs222 = globalState;
        _1659_v186 = _rhs271;
        _lhs222.f0 = _rhs272;
      }
      let _1660_v187;
      _1660_v187 = _dafny.Seq.UnicodeFromString("emlcy");
      _1660_v187 = _dafny.Seq.UnicodeFromString("usnjcg");
      if (!(!(!_dafny.areEqual(_1660_v187, _1660_v187)))) {
        _1660_v187 = _module.__default.fm30(globalState);
        let _1661_v188;
        _1661_v188 = _dafny.MultiSet.fromElements(new BigNumber((_1660_v187).length));
        _1661_v188 = (_dafny.MultiSet.fromElements(_1447_v50, _1447_v50, (_dafny.ZERO).minus(_module.__default.fm7(_1447_v50, globalState)), _1447_v50, _1447_v50)).Difference((_1661_v188).Intersect(_1661_v188));
        let _1662_v189;
        let _nw233 = new _module.C4();
        _nw233.__ctor(p0, (_this).f22);
        _1662_v189 = _nw233;
        r0 = _dafny.Seq.IsProperPrefixOf(_1660_v187, _dafny.Seq.Concat(_1660_v187, _1660_v187));
        let _1663_v190;
        _1663_v190 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)))).length));
        let _1664_v191;
        _1664_v191 = _dafny.Set.fromElements(!(p0));
        let _1665_v192;
        _1665_v192 = _dafny.Set.fromElements(_1447_v50, new BigNumber((_1663_v190).length), (_dafny.ZERO).minus(new BigNumber((_1664_v191).length)));
        let _1666_v194;
        _1666_v194 = _dafny.MultiSet.fromElements(((p0) ? (_1665_v192) : (function () {
          let _coll66 = new _dafny.Set();
          for (const _compr_66 of _dafny.IntegerRange(new BigNumber(965), new BigNumber(848))) {
            let _1667_v193 = _compr_66;
            if (((new BigNumber(965)).isLessThanOrEqualTo(_1667_v193)) && ((_1667_v193).isLessThan(new BigNumber(848)))) {
              _coll66.add(_module.__default.safeModuloInt(_1667_v193, _1447_v50));
            }
          }
          return _coll66;
        }())), _1665_v192, _1665_v192, ((p0) ? (_1665_v192) : (_1665_v192)));
        let _index232 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1378_v0).length));
        (_1378_v0)[_index232] = _1660_v187;
        let _1668_v196;
        _1668_v196 = _dafny.Map.Empty.slice().updateUnsafe(_1660_v187,new BigNumber((function () {
          let _coll67 = new _dafny.Set();
          for (const _compr_67 of _dafny.IntegerRange(new BigNumber(377), new BigNumber(-477))) {
            let _1669_v195 = _compr_67;
            if (((new BigNumber(377)).isLessThanOrEqualTo(_1669_v195)) && ((_1669_v195).isLessThan(new BigNumber(-477)))) {
              _coll67.add(_module.__default.safeDivisionInt(_1669_v195, _1447_v50));
            }
          }
          return _coll67;
        }()).length));
        let _index233 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1378_v0).length));
        let _rhs273 = _1666_v194;
        let _rhs274 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(610)), function (_1670_i17) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        });
        let _rhs275 = _1662_v189.f28;
        let _rhs276 = _1660_v187;
        let _rhs277 = (((_this).f24) ? (new BigNumber(((_1668_v196).Merge(_1668_v196)).length)) : ((_1447_v50).multipliedBy(new BigNumber(841))));
        let _lhs223 = _1378_v0;
        let _lhs224 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1378_v0).length));
        let _lhs225 = globalState;
        let _lhs226 = globalState;
        _1666_v194 = _rhs273;
        _lhs223[_lhs224] = _rhs274;
        _lhs225.f0 = _rhs275;
        _1660_v187 = _rhs276;
        _lhs226.f16 = _rhs277;
      } else {
        let _1671_v197;
        _1671_v197 = new _dafny.CodePoint('d'.codePointAt(0));
        if (!(_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("emuwsan"), _1671_v197))) {
          let _index234 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_1378_v0).length));
          (_1378_v0)[_index234] = _1660_v187;
          let _index235 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_1378_v0).length));
          (_1378_v0)[_index235] = _1660_v187;
          let _1672_v198;
          let _init40 = ((_1673_v50) => function (_1674_i18) {
            return _module.__default.safeModuloInt(_1674_i18, _1673_v50);
          })(_1447_v50);
          let _nw234 = Array((new BigNumber(3)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw234.length); _i0_40++) {
            _nw234[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1672_v198 = _nw234;
          let _1675_v199;
          _1675_v199 = _dafny.Set.fromElements(new BigNumber(-974));
          let _1676_v200;
          _1676_v200 = _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_1675_v199).length)));
          let _index236 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_1672_v198).length));
          (_1672_v198)[_index236] = new BigNumber(((_1676_v200).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-823)))).length);
          let _1677_v201;
          _1677_v201 = _dafny.Set.fromElements(p0, true);
          let _1678_v202;
          _1678_v202 = _dafny.Seq.of(_1447_v50, new BigNumber(((_1677_v201).Difference(_1677_v201)).length), _1447_v50);
          let _1679_v203;
          let _init41 = ((_1680_v50) => function (_1681_i19) {
            return (_1680_v50).isLessThanOrEqualTo(new BigNumber(417));
          })(_1447_v50);
          let _nw235 = Array((new BigNumber(26)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw235.length); _i0_41++) {
            _nw235[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1679_v203 = _nw235;
          let _index237 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_1672_v198).length));
          let _rhs278 = (_dafny.ZERO).minus(new BigNumber((_1678_v202).length));
          let _rhs279 = _1679_v203;
          let _rhs280 = _dafny.Seq.IsPrefixOf(_1660_v187, _1660_v187);
          let _lhs227 = _1672_v198;
          let _lhs228 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_1672_v198).length));
          let _lhs229 = globalState;
          _lhs227[_lhs228] = _rhs278;
          r1 = _rhs279;
          _lhs229.f0 = _rhs280;
          let _1682_v204;
          _1682_v204 = _dafny.Seq.of(p0, (_this).f24);
          let _1683_v205;
          _1683_v205 = _dafny.Seq.of(_dafny.Seq.Concat(_1682_v204, _1682_v204));
          _1683_v205 = _1683_v205;
          let _1684_v206;
          let _nw236 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1684_v206 = _nw236;
          let _1685_v207;
          _1685_v207 = _dafny.MultiSet.fromElements((_1672_v198)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_1672_v198).length))]);
          let _index238 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1684_v206).length));
          (_1684_v206)[_index238] = _1685_v207;
          let _index239 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_1684_v206).length));
          (_1684_v206)[_index239] = _1685_v207;
          let _1686_v208;
          _1686_v208 = _module.D2.create_DC6((_1672_v198)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_1672_v198).length))]);
          let _1687_v209;
          _1687_v209 = _module.D2.create_DC9(_module.D2.create_DC9(_1686_v208));
          let _1688_v210;
          _1688_v210 = _dafny.MultiSet.fromElements(_1687_v209);
          let _1689_v211;
          let _nw237 = new _module.C4();
          _nw237.__ctor(!(_1688_v210).equals(_1688_v210), (_this).f22);
          _1689_v211 = _nw237;
        } else {
          let _1690_v212;
          _1690_v212 = _module.D11.create_DC39((_this).f24);
          let _1691_v213;
          let _nw238 = Array((_dafny.ONE).toNumber());
          _nw238[(_dafny.ZERO).toNumber()] = _1690_v212;
          _1691_v213 = _nw238;
          let _index240 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1691_v213).length));
          (_1691_v213)[_index240] = _module.D11.create_DC39((_this).f24);
          let _1692_v214;
          _1692_v214 = _module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-352)), ((_1693_v197) => function (_1694_i20) {
  return _1693_v197;
})(_1671_v197)));
          let _1695_v215;
          _1695_v215 = _dafny.Seq.of(_1692_v214, function (_pat_let32_0) {
            return function (_1696_dt__update__tmp_h3) {
              return function (_pat_let33_0) {
                return function (_1697_dt__update_hcf6_h0) {
                  return _module.D1.create_DC3(_1697_dt__update_hcf6_h0);
                }(_pat_let33_0);
              }(_dafny.Seq.UnicodeFromString("ujsmsmjoj"));
            }(_pat_let32_0);
          }(_module.D1.create_DC3(_1660_v187)));
          let _1698_v216;
          let _nw239 = Array((new BigNumber(4)).toNumber());
          _nw239[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
          _nw239[(_dafny.ONE).toNumber()] = _1671_v197;
          _nw239[(new BigNumber(2)).toNumber()] = (_1660_v187)[_module.__default.safeIndex(_1447_v50, new BigNumber((_1660_v187).length))];
          _nw239[(new BigNumber(3)).toNumber()] = _1671_v197;
          _1698_v216 = _nw239;
          let _1699_v217;
          _1699_v217 = _dafny.Seq.of(p0);
          let _index241 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1698_v216).length));
          (_1698_v216)[_index241] = (_1660_v187)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1699_v217).length)), new BigNumber((_1660_v187).length))];
          let _1700_v218;
          _1700_v218 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1447_v50);
          let _1701_v219;
          _1701_v219 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1700_v218).length),new BigNumber(4));
          let _1702_v220;
          _1702_v220 = _module.D0.create_DC1(_1447_v50, (p0) || (p0), _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1701_v219), ((p0) ? (_1447_v50) : (new BigNumber(186))), _1671_v197);
          let _index242 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1691_v213).length));
          let _index243 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1698_v216).length));
          let _rhs281 = _1447_v50;
          let _rhs282 = _1690_v212;
          let _rhs283 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(90)), ((_1703_v214) => function (_1704_i21) {
            return _1703_v214;
          })(_1692_v214));
          let _rhs284 = _1671_v197;
          let _rhs285 = _1702_v220;
          let _lhs230 = _1691_v213;
          let _lhs231 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1691_v213).length));
          let _lhs232 = _1698_v216;
          let _lhs233 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_1698_v216).length));
          r3 = _rhs281;
          _lhs230[_lhs231] = _rhs282;
          _1695_v215 = _rhs283;
          _lhs232[_lhs233] = _rhs284;
          _1702_v220 = _rhs285;
          let _rhs286 = _1447_v50;
          let _rhs287 = (_1698_v216)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1698_v216).length))];
          let _lhs234 = globalState;
          _lhs234.f11 = _rhs286;
          _1671_v197 = _rhs287;
          let _1705_v221;
          let _nw240 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
          _1705_v221 = _nw240;
          let _1706_v222;
          let _nw241 = new _module.C1();
          _nw241.__ctor(_1447_v50, (_this).f22);
          _1706_v222 = _nw241;
          let _1707_v223;
          _1707_v223 = _dafny.MultiSet.fromElements(p0);
          let _1708_v224;
          _1708_v224 = _module.D14.create_DC47(_1707_v223);
          let _1709_v225;
          _1709_v225 = _dafny.Map.Empty.slice().updateUnsafe(_1706_v222,(_1708_v224).dtor_cf77);
          let _index244 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_1705_v221).length));
          (_1705_v221)[_index244] = _1709_v225;
          let _index245 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_1705_v221).length));
          (_1705_v221)[_index245] = ((_1709_v225).Merge((_1709_v225).update(_1706_v222, _1707_v223))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1706_v222,_1707_v223)).Merge(_1709_v225));
          let _1710_v226;
          let _nw242 = Array((new BigNumber(7)).toNumber());
          _1710_v226 = _nw242;
          let _1711_v227;
          _1711_v227 = _module.D11.create_DC37(_1710_v226);
          let _1712_v228;
          let _nw243 = Array((new BigNumber(10)).toNumber());
          _nw243[(_dafny.ZERO).toNumber()] = _1447_v50;
          _nw243[(_dafny.ONE).toNumber()] = new BigNumber((_1660_v187).length);
          _nw243[(new BigNumber(2)).toNumber()] = _1447_v50;
          _nw243[(new BigNumber(3)).toNumber()] = _1447_v50;
          _nw243[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.update(_1660_v187, _module.__default.safeIndex(new BigNumber((_1660_v187).length), new BigNumber((_1660_v187).length)), (_1698_v216)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_1698_v216).length))])).length);
          _nw243[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.update(_1699_v217, _module.__default.safeIndex(new BigNumber((_1707_v223).cardinality()), new BigNumber((_1699_v217).length)), (_this).f24)).length);
          _nw243[(new BigNumber(6)).toNumber()] = _1447_v50;
          _nw243[(new BigNumber(7)).toNumber()] = _1447_v50;
          _nw243[(new BigNumber(8)).toNumber()] = new BigNumber(380);
          _nw243[(new BigNumber(9)).toNumber()] = _1447_v50;
          _1712_v228 = _nw243;
          let _1713_v229;
          _1713_v229 = _module.D2.create_DC8(_1692_v214, _1712_v228, p0);
          let _1714_v230;
          _1714_v230 = _dafny.Map.Empty.slice().updateUnsafe(_1711_v227,_1713_v229);
          let _1715_v231;
          _1715_v231 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v50,_1710_v226);
          let _1716_v232;
          _1716_v232 = _dafny.Seq.of(_1712_v228);
          let _pat_let_tv61 = _1715_v231;
          let _pat_let_tv62 = _1716_v232;
          let _pat_let_tv63 = _1710_v226;
          let _pat_let_tv64 = _1715_v231;
          let _pat_let_tv65 = _1716_v232;
          _1714_v230 = (_1714_v230).update(function (_pat_let34_0) {
            return function (_1717_dt__update__tmp_h4) {
              return function (_pat_let35_0) {
                return function (_1718_dt__update_hcf63_h0) {
                  return _module.D11.create_DC37(_1718_dt__update_hcf63_h0);
                }(_pat_let35_0);
              }((((_pat_let_tv64).contains(new BigNumber((_pat_let_tv65).length))) ? ((_pat_let_tv61).get(new BigNumber((_pat_let_tv62).length))) : (_pat_let_tv63)));
            }(_pat_let34_0);
          }(_1711_v227), _1713_v229);
          let _1719_v233;
          let _nw244 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1719_v233 = _nw244;
          let _1720_v234;
          _1720_v234 = _dafny.MultiSet.fromElements(_1719_v233, _1719_v233, _1719_v233);
          _1720_v234 = ((_1720_v234).Difference(_dafny.MultiSet.fromElements(_1719_v233, _1719_v233))).Intersect(_1720_v234);
        }
        let _1721_v235;
        let _nw245 = Array((new BigNumber(27)).toNumber()).fill(false);
        _1721_v235 = _nw245;
        let _index246 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1721_v235).length));
        (_1721_v235)[_index246] = p0;
        let _index247 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1721_v235).length));
        (_1721_v235)[_index247] = p0;
        let _1722_v236;
        let _nw246 = new _module.C4();
        _nw246.__ctor((_1721_v235)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1721_v235).length))], (_this).f22);
        _1722_v236 = _nw246;
        let _1723_v237;
        _1723_v237 = _dafny.Set.fromElements(_1447_v50);
        let _1724_v238;
        _1724_v238 = _dafny.Seq.of(_1723_v237);
        let _1725_v239;
        _1725_v239 = _dafny.Set.fromElements(true, (_this).f24, p0);
        let _1726_v240;
        let _nw247 = Array((new BigNumber(9)).toNumber());
        _nw247[(_dafny.ZERO).toNumber()] = new BigNumber(-876);
        _nw247[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1660_v187).length));
        _nw247[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_1447_v50);
        _nw247[(new BigNumber(3)).toNumber()] = _1447_v50;
        _nw247[(new BigNumber(4)).toNumber()] = _1447_v50;
        _nw247[(new BigNumber(5)).toNumber()] = _1447_v50;
        _nw247[(new BigNumber(6)).toNumber()] = _1447_v50;
        _nw247[(new BigNumber(7)).toNumber()] = (_1722_v236).fm0(((_this).f25)[_module.__default.safeIndex(_1447_v50, new BigNumber(((_this).f25).length))], _1447_v50, _1724_v238, (_this).f24, globalState);
        _nw247[(new BigNumber(8)).toNumber()] = new BigNumber((_1725_v239).length);
        _1726_v240 = _nw247;
        let _1727_v241;
        _1727_v241 = _dafny.Set.fromElements(_1726_v240);
        (globalState).f10 = (_module.D13.create_DC46((_this).f24, _1722_v236, _1447_v50, _1727_v241)).dtor_cf73;
        let _1728_v242;
        _1728_v242 = _dafny.Seq.of((_1721_v235)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1721_v235).length))]);
        let _1729_v243;
        _1729_v243 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_1728_v242, _dafny.Seq.of((_this).f24)));
        let _1730_v244;
        _1730_v244 = _dafny.MultiSet.fromElements((_1721_v235)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1721_v235).length))], (_this).f24);
        _1729_v243 = _module.__default.fm50((new BigNumber((_1730_v244).cardinality())).multipliedBy(_1447_v50), globalState);
        (globalState).f0 = (_1721_v235)[_module.__default.safeIndex(new BigNumber(46), new BigNumber((_1721_v235).length))];
      }
      let _1731_v245;
      _1731_v245 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1732_v246;
      _1732_v246 = _dafny.Set.fromElements(_1731_v245);
      let _1733_v247;
      _1733_v247 = _dafny.Map.Empty.slice().updateUnsafe(_1731_v245,_dafny.Set.fromElements(_1732_v246, _module.__default.fm33(globalState), _1732_v246, _1732_v246));
      let _1734_v248;
      _1734_v248 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_1731_v245),new BigNumber((_dafny.Seq.of(_1447_v50)).length));
      r0 = (_module.__default.fm7(_1447_v50, globalState)).isLessThan(new BigNumber(((((_1733_v247).contains(_1731_v245)) ? ((_1733_v247).get(_1731_v245)) : (function () {
        let _coll68 = new _dafny.Set();
        for (const _compr_68 of (_1734_v248).Keys.Elements) {
          let _1735_v249 = _compr_68;
          if ((_1734_v248).contains(_1735_v249)) {
            _coll68.add(_1735_v249);
          }
        }
        return _coll68;
      }()))).length));
      let _nw248 = Array((new BigNumber(21)).toNumber()).fill(false);
      r1 = _nw248;
      r2 = function () {
        let _coll69 = new _dafny.Map();
        for (const _compr_69 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(915)), ((_1736_v50, _1737_v245) => function (_1738_i22) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), ((_1739_v50, _1740_v245) => function (_1741_i23) {
            return new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("jfkcni"), _module.__default.safeIndex(_1739_v50, new BigNumber((_dafny.Seq.UnicodeFromString("jfkcni")).length)), _1740_v245)).length);
          })(_1736_v50, _1737_v245));
        })(_1447_v50, _1731_v245))).Elements) {
          let _1742_v250 = _compr_69;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(915)), ((_1743_v50, _1744_v245) => function (_1738_i22) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), ((_1745_v50, _1746_v245) => function (_1741_i23) {
              return new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("jfkcni"), _module.__default.safeIndex(_1745_v50, new BigNumber((_dafny.Seq.UnicodeFromString("jfkcni")).length)), _1746_v245)).length);
            })(_1743_v50, _1744_v245));
          })(_1447_v50, _1731_v245)), _1742_v250)) {
            _coll69.push([_1742_v250,_1731_v245]);
          }
        }
        return _coll69;
      }();
      r3 = _module.__default.safeModuloInt(_1447_v50, _1447_v50);
      return [r0, r1, r2, r3];
    }
    m1(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _1747_v0;
      _1747_v0 = new BigNumber(217);
      r1 = _1747_v0;
      if ((_this).f24) {
        let _1748_v1;
        _1748_v1 = _module.D15.create_DC49(_dafny.Set.fromElements((_this).f24));
        let _1749_v2;
        _1749_v2 = _dafny.Set.fromElements((_this).f24);
        (globalState).f0 = (((_1748_v1).dtor_cf81).Difference(_1749_v2)).equals(_1749_v2);
        let _1750_v3;
        let _init42 = ((_1751_v0) => function (_1752_i0) {
          return _module.__default.safeModuloInt(_1752_i0, new BigNumber((_dafny.Set.fromElements(_1751_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1751_v0)).length))).length));
        })(_1747_v0);
        let _nw249 = Array((new BigNumber(23)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw249.length); _i0_42++) {
          _nw249[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1750_v3 = _nw249;
        let _1753_v4;
        _1753_v4 = _dafny.Set.fromElements(_1750_v3, _1750_v3, _1750_v3, _1750_v3);
        let _1754_v5;
        _1754_v5 = _dafny.Seq.UnicodeFromString("uxel");
        let _1755_v6;
        _1755_v6 = _module.D15.create_DC51((_this).f24, (_this).f24);
        let _1756_v7;
        _1756_v7 = _dafny.Seq.of((_this).f24, (_this).f24, (_this).f24, (_this).f24);
        let _1757_v8;
        _1757_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1747_v0,!((_this).f24));
        let _1758_v9;
        let _nw250 = Array((new BigNumber(24)).toNumber());
        _nw250[(_dafny.ZERO).toNumber()] = (_this).f24;
        _nw250[(_dafny.ONE).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(2)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(3)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(4)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(5)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(6)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(7)).toNumber()] = !((_this).f24);
        _nw250[(new BigNumber(8)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(9)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(10)).toNumber()] = _module.__default.fm15((_this).f24, globalState);
        _nw250[(new BigNumber(11)).toNumber()] = !((_this).f24);
        _nw250[(new BigNumber(12)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(13)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(14)).toNumber()] = (_1755_v6).dtor_cf82;
        _nw250[(new BigNumber(15)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(16)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(17)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(18)).toNumber()] = _module.__default.fm37(_1756_v7, _1757_v8, _1754_v5, globalState);
        _nw250[(new BigNumber(19)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(20)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(21)).toNumber()] = (_this).f24;
        _nw250[(new BigNumber(22)).toNumber()] = _module.__default.fm15((_this).f24, globalState);
        _nw250[(new BigNumber(23)).toNumber()] = true;
        _1758_v9 = _nw250;
        let _1759_v10;
        _1759_v10 = _module.D9.create_DC31((new BigNumber((_1753_v4).length)).plus(new BigNumber((_1754_v5).length)), _1758_v9);
        let _index248 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1750_v3).length));
        (_1750_v3)[_index248] = _1747_v0;
        let _index249 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1750_v3).length));
        let _rhs288 = _1759_v10;
        let _rhs289 = _1750_v3;
        let _rhs290 = _1747_v0;
        let _rhs291 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_1756_v7, _module.__default.safeIndex(_1747_v0, new BigNumber((_1756_v7).length)), (_this).f24), _module.__default.safeIndex((_dafny.ZERO).minus(_1747_v0), new BigNumber((_dafny.Seq.update(_1756_v7, _module.__default.safeIndex(_1747_v0, new BigNumber((_1756_v7).length)), (_this).f24)).length)), (_this).f24), _dafny.Seq.of((_this).f24, (_this).f24, (_this).f24));
        let _rhs292 = _1747_v0;
        let _lhs235 = _1750_v3;
        let _lhs236 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1750_v3).length));
        let _lhs237 = globalState;
        _1759_v10 = _rhs288;
        _1750_v3 = _rhs289;
        _lhs235[_lhs236] = _rhs290;
        _1756_v7 = _rhs291;
        _lhs237.f8 = _rhs292;
        let _1760_v11;
        _1760_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1750_v3)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1750_v3).length))],_module.__default.fm7((_dafny.ZERO).minus(new BigNumber((_1749_v2).length)), globalState));
        let _1761_v12;
        _1761_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
        _1760_v11 = (_1760_v11).update(_module.__default.safeDivisionInt(_1747_v0, (_1750_v3)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1750_v3).length))]), (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.of(_1761_v12, _1761_v12)).length)).minus((_1750_v3)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1750_v3).length))])));
        let _1762_v13;
        _1762_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1754_v5,(_this).f24);
        let _1763_v14;
        _1763_v14 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f24),_1762_v13);
        let _rhs293 = (_1762_v13).Merge((((_1763_v14).contains((_this).f24)) ? ((_1763_v14).get((_this).f24)) : (_1762_v13)));
        let _rhs294 = (new BigNumber((_1754_v5).length)).isEqualTo((_1750_v3)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1750_v3).length))]);
        let _lhs238 = globalState;
        _1762_v13 = _rhs293;
        _lhs238.f0 = _rhs294;
        _1747_v0 = _1747_v0;
      } else {
        let _1764_v15;
        let _init43 = function (_1765_i1) {
          return (_this).f24;
        };
        let _nw251 = Array((new BigNumber(2)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw251.length); _i0_43++) {
          _nw251[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1764_v15 = _nw251;
        let _1766_v16;
        let _nw252 = new _module.C1();
        _nw252.__ctor(_module.__default.safeDivisionInt(_1747_v0, new BigNumber(817)), _dafny.Seq.Concat(_dafny.Seq.of(_1764_v15, _1764_v15, _1764_v15, _1764_v15, _1764_v15), (_this).f22));
        _1766_v16 = _nw252;
        if (_module.__default.fm15(true, globalState)) {
          let _1767_v17;
          let _1768_v18;
          let _1769_v19;
          let _1770_v20;
          let _out60;
          let _out61;
          let _out62;
          let _out63;
          let _outcollector19 = (_1766_v16).m6(_1747_v0, ((_1766_v16).f30).isEqualTo((_1766_v16).f30), _module.__default.safeModuloInt(_1747_v0, _1747_v0), (_1766_v16).f30, globalState);
          _out60 = _outcollector19[0];
          _out61 = _outcollector19[1];
          _out62 = _outcollector19[2];
          _out63 = _outcollector19[3];
          _1767_v17 = _out60;
          _1768_v18 = _out61;
          _1769_v19 = _out62;
          _1770_v20 = _out63;
          let _1771_v21;
          _1771_v21 = new _dafny.CodePoint('q'.codePointAt(0));
          _1771_v21 = _1771_v21;
          let _1772_v22;
          _1772_v22 = _dafny.Set.fromElements((_this).f24, (_this).f24, false);
          let _1773_v23;
          _1773_v23 = _dafny.Seq.of((_this).f24);
          let _1774_v24;
          let _nw253 = Array((new BigNumber(23)).toNumber());
          _nw253[(_dafny.ZERO).toNumber()] = (_1772_v22).Intersect(_1772_v22);
          _nw253[(_dafny.ONE).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(2)).toNumber()] = (((_this).f24) ? (_1772_v22) : (_1772_v22));
          _nw253[(new BigNumber(3)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(4)).toNumber()] = (_1772_v22).Intersect(_1772_v22);
          _nw253[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements((_this).f24);
          _nw253[(new BigNumber(6)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(7)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(8)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(9)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(10)).toNumber()] = (_1772_v22).Intersect(_1772_v22);
          _nw253[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements((_this).f24, (_this).f24, (_this).f24);
          _nw253[(new BigNumber(12)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(13)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(14)).toNumber()] = (_dafny.Set.fromElements((_1773_v23)[_module.__default.safeIndex((_1766_v16).f30, new BigNumber((_1773_v23).length))])).Difference(_1772_v22);
          _nw253[(new BigNumber(15)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(16)).toNumber()] = (_1772_v22).Union(_1772_v22);
          _nw253[(new BigNumber(17)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(18)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(19)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(20)).toNumber()] = _1772_v22;
          _nw253[(new BigNumber(21)).toNumber()] = (_module.__default.fm11(_1770_v20, _1771_v21, (_this).f24, globalState)).Union(_1772_v22);
          _nw253[(new BigNumber(22)).toNumber()] = _dafny.Set.fromElements(!((_this).f24), false);
          _1774_v24 = _nw253;
          let _index250 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1774_v24).length));
          (_1774_v24)[_index250] = _1772_v22;
          let _index251 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_1774_v24).length));
          (_1774_v24)[_index251] = _1772_v22;
          let _1775_v25;
          _1775_v25 = _module.D13.create_DC45(_1769_v19);
          let _pat_let_tv66 = _1767_v17;
          _1775_v25 = function (_pat_let36_0) {
            return function (_1776_dt__update__tmp_h0) {
              return function (_pat_let37_0) {
                return function (_1777_dt__update_hcf72_h0) {
                  return _module.D13.create_DC45(_1777_dt__update_hcf72_h0);
                }(_pat_let37_0);
              }(_pat_let_tv66);
            }(_pat_let36_0);
          }(_1775_v25);
          let _1778_v26;
          let _nw254 = new _module.C2();
          _nw254.__ctor(_1770_v20, _dafny.Seq.of(_1764_v15, _1764_v15, _1764_v15, _1764_v15));
          _1778_v26 = _nw254;
          let _1779_v27;
          _1779_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(296),_1767_v17);
          let _1780_v28;
          let _nw255 = Array((new BigNumber(11)).toNumber());
          _nw255[(_dafny.ZERO).toNumber()] = (_1766_v16).f30;
          _nw255[(_dafny.ONE).toNumber()] = new BigNumber(-755);
          _nw255[(new BigNumber(2)).toNumber()] = new BigNumber((_1779_v27).length);
          _nw255[(new BigNumber(3)).toNumber()] = new BigNumber(978);
          _nw255[(new BigNumber(4)).toNumber()] = (_1766_v16).f30;
          _nw255[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(692)), ((_1781_v21) => function (_1782_i2) {
            return _1781_v21;
          })(_1771_v21))).length);
          _nw255[(new BigNumber(6)).toNumber()] = _1769_v19;
          _nw255[(new BigNumber(7)).toNumber()] = (_1778_v26).f29;
          _nw255[(new BigNumber(8)).toNumber()] = _1769_v19;
          _nw255[(new BigNumber(9)).toNumber()] = (_1766_v16).f30;
          _nw255[(new BigNumber(10)).toNumber()] = _1747_v0;
          _1780_v28 = _nw255;
          let _1783_v29;
          _1783_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Map.Empty.slice().updateUnsafe(_1778_v26,_1780_v28));
          let _1784_v30;
          _1784_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1778_v26,_1780_v28);
          _1783_v29 = (_1783_v29).update(!((_this).f24), (_1784_v30).Merge(_1784_v30));
        } else {
          let _1785_v31;
          _1785_v31 = _dafny.Seq.UnicodeFromString("iyldgbb");
          let _1786_v32;
          _1786_v32 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1787_v33;
          _1787_v33 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1785_v31, _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1785_v31).length)), _1786_v32), _1785_v31)).length));
          let _1788_v34;
          _1788_v34 = _module.D2.create_DC7(_1787_v33, (_this).f24, (_this).f24, (_1766_v16).f30);
          let _1789_v35;
          _1789_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-395),(_1788_v34).dtor_cf10);
          _1787_v33 = _dafny.Seq.Concat(_dafny.Seq.update(_1787_v33, _module.__default.safeIndex((_1766_v16).fm34(false, globalState), new BigNumber((_1787_v33).length)), new BigNumber((_1785_v31).length)), (((_1789_v35).contains((_1766_v16).f30)) ? ((_1789_v35).get((_1766_v16).f30)) : (_1787_v33)));
          let _1790_v36;
          let _nw256 = new _module.C1();
          _nw256.__ctor((_1766_v16).f30, _dafny.Seq.Concat(_dafny.Seq.update((_this).f22, _module.__default.safeIndex((_1766_v16).f30, new BigNumber(((_this).f22).length)), _1764_v15), _dafny.Seq.of(_1764_v15)));
          _1790_v36 = _nw256;
          let _1791_v37;
          _1791_v37 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24, (_this).f24, (_this).f24);
          let _1792_v38;
          _1792_v38 = _dafny.MultiSet.fromElements(_1747_v0);
          let _1793_v39;
          _1793_v39 = _dafny.Set.fromElements((_1766_v16).f30);
          let _1794_v40;
          _1794_v40 = _dafny.Seq.of(_1793_v39);
          let _1795_v41;
          _1795_v41 = _dafny.Map.Empty.slice().updateUnsafe((_1766_v16).fm0(_1792_v38, _1747_v0, _1794_v40, (_this).f24, globalState),(_this).f24);
          (globalState).f16 = ((((_1791_v37).contains((_this).f24)) ? ((_1791_v37).get((_this).f24)) : (new BigNumber((_1795_v41).length)))).multipliedBy(((_1790_v36).f30).plus((_1790_v36).f30));
          r0 = (_this).fm3(_1785_v31, (((_this).f24) ? ((_this).f24) : ((_this).f24)), (_this).f24, globalState);
          let _1796_v42;
          _1796_v42 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vcllqkg"),_1787_v33);
          let _1797_v44;
          _1797_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1785_v31,(_1790_v36).f30);
          _1796_v42 = function () {
            let _coll70 = new _dafny.Map();
            for (const _compr_70 of (_1797_v44).Keys.Elements) {
              let _1798_v43 = _compr_70;
              if ((_1797_v44).contains(_1798_v43)) {
                _coll70.push([_1798_v43,_dafny.Seq.Concat(_1787_v33, _1787_v33)]);
              }
            }
            return _coll70;
          }();
        }
        r1 = _1747_v0;
        let _1799_v45;
        _1799_v45 = _dafny.Seq.of(_module.__default.fm51(globalState));
        let _1800_v46;
        let _nw257 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _1800_v46 = _nw257;
        let _1801_v47;
        _1801_v47 = _dafny.Seq.of(_1800_v46);
        let _1802_v48;
        _1802_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(573),_1800_v46);
        let _1803_v49;
        _1803_v49 = _dafny.Set.fromElements((_1766_v16).f30);
        let _1804_v50;
        let _nw258 = Array((new BigNumber(19)).toNumber());
        _nw258[(_dafny.ZERO).toNumber()] = _1800_v46;
        _nw258[(_dafny.ONE).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(2)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(3)).toNumber()] = (_1801_v47)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_1801_v47).length))];
        _nw258[(new BigNumber(4)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(5)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(6)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(7)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(8)).toNumber()] = (((_this).f24) ? (_1800_v46) : (_1800_v46));
        _nw258[(new BigNumber(9)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(10)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(11)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(12)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(13)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(14)).toNumber()] = (((_this).f24) ? (_1800_v46) : (_1800_v46));
        _nw258[(new BigNumber(15)).toNumber()] = _1800_v46;
        _nw258[(new BigNumber(16)).toNumber()] = (((_1802_v48).contains((_1766_v16).f30)) ? ((_1802_v48).get((_1766_v16).f30)) : (_1800_v46));
        _nw258[(new BigNumber(17)).toNumber()] = (((_this).f24) ? ((((_1802_v48).contains(new BigNumber((_1803_v49).length))) ? ((_1802_v48).get(new BigNumber((_1803_v49).length))) : (_1800_v46))) : (_1800_v46));
        _nw258[(new BigNumber(18)).toNumber()] = _1800_v46;
        _1804_v50 = _nw258;
        let _index252 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1804_v50).length));
        (_1804_v50)[_index252] = _1800_v46;
        let _1805_v51;
        _1805_v51 = _module.D4.create_DC15(new BigNumber(238), _1747_v0);
        let _index253 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1804_v50).length));
        let _rhs295 = _dafny.Seq.Concat(_1799_v45, _1799_v45);
        let _rhs296 = _1747_v0;
        let _rhs297 = _1800_v46;
        let _rhs298 = ((_1747_v0).isLessThanOrEqualTo((_1805_v51).dtor_cf26)) === ((_this).f24);
        let _lhs239 = globalState;
        let _lhs240 = _1804_v50;
        let _lhs241 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1804_v50).length));
        let _lhs242 = globalState;
        _1799_v45 = _rhs295;
        _lhs239.f6 = _rhs296;
        _lhs240[_lhs241] = _rhs297;
        _lhs242.f10 = _rhs298;
        let _1806_v52;
        let _nw259 = new _module.C3();
        _nw259.__ctor(_dafny.Seq.Concat((_this).f22, (_this).f22));
        _1806_v52 = _nw259;
      }
      let _1807_i3;
      _1807_i3 = _dafny.ZERO;
      L7: {
        while ((_this).f24) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1807_i3)) {
              break L7;
            }
            _1807_i3 = (_1807_i3).plus(_dafny.ONE);
            let _1808_v53;
            let _init44 = function (_1809_i4) {
              return !(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24)).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24));
            };
            let _nw260 = Array((new BigNumber(4)).toNumber());
            for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw260.length); _i0_44++) {
              _nw260[_i0_44] = _init44(new BigNumber(_i0_44));
            }
            _1808_v53 = _nw260;
            let _index254 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1808_v53).length));
            (_1808_v53)[_index254] = !((_this).f24);
            let _1810_v54;
            _1810_v54 = _dafny.Seq.of(_1747_v0);
            let _1811_v55;
            _1811_v55 = _dafny.Seq.of(_1810_v54, _dafny.Seq.update(_1810_v54, _module.__default.safeIndex(_1747_v0, new BigNumber((_1810_v54).length)), _1747_v0), _1810_v54);
            let _index255 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1808_v53).length));
            (_1808_v53)[_index255] = !_dafny.Seq.contains(_1811_v55, _1810_v54);
            let _1812_v56;
            _1812_v56 = _dafny.Seq.of((_1808_v53)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1808_v53).length))]);
            let _1813_v57;
            _1813_v57 = _dafny.MultiSet.fromElements(_1747_v0, _1747_v0, _1747_v0, new BigNumber((_1812_v56).length));
            (globalState).f0 = !((_1813_v57).Intersect(_1813_v57)).contains(_1747_v0);
            let _1814_v58;
            _1814_v58 = _dafny.Set.fromElements((_1808_v53)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1808_v53).length))], (_1808_v53)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1808_v53).length))], (_this).f24);
            let _1815_v59;
            _1815_v59 = _dafny.Set.fromElements(new BigNumber(526), new BigNumber((_1814_v58).length), new BigNumber(272), _1747_v0);
            (globalState).f13 = _1815_v59;
            let _1816_v60;
            _1816_v60 = _dafny.Map.Empty.slice().updateUnsafe((_1808_v53)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1808_v53).length))],_1747_v0);
            let _1817_v61;
            _1817_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1747_v0,new BigNumber((_1816_v60).length));
            let _1818_v62;
            _1818_v62 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f24),_1817_v61);
            _1818_v62 = (_1818_v62).update((_1808_v53)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_1808_v53).length))], (_1817_v61).update(_1747_v0, _1747_v0));
          }
        }
      }
      if (!((_this).f24)) {
        let _1819_v63;
        let _nw261 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1819_v63 = _nw261;
        let _1820_v64;
        _1820_v64 = _dafny.Seq.UnicodeFromString("xlluhdhw");
        let _index256 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_1819_v63).length));
        (_1819_v63)[_index256] = _1820_v64;
        let _1821_v65;
        _1821_v65 = _dafny.Set.fromElements(_1747_v0);
        let _1822_v67;
        _1822_v67 = _dafny.MultiSet.fromElements(true);
        let _index257 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_1819_v63).length));
        (_1819_v63)[_index257] = _module.__default.fm36((_this).f24, (_1821_v65).Difference(function () {
          let _coll71 = new _dafny.Set();
          for (const _compr_71 of _dafny.IntegerRange(new BigNumber(9), new BigNumber(495))) {
            let _1823_v66 = _compr_71;
            if (((new BigNumber(9)).isLessThanOrEqualTo(_1823_v66)) && ((_1823_v66).isLessThan(new BigNumber(495)))) {
              _coll71.add(_module.__default.safeModuloInt(_1823_v66, _1747_v0));
            }
          }
          return _coll71;
        }()), _1747_v0, new BigNumber((_1822_v67).cardinality()), globalState);
        let _1824_v68;
        let _init45 = function (_1825_i5) {
          return (_this).f24;
        };
        let _nw262 = Array((_dafny.ONE).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw262.length); _i0_45++) {
          _nw262[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1824_v68 = _nw262;
        let _1826_v69;
        _1826_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1824_v68,(_this).f24);
        _1826_v69 = ((_1826_v69).update(_1824_v68, (_this).f24)).Merge(_1826_v69);
        let _1827_v70;
        _1827_v70 = new _dafny.CodePoint('b'.codePointAt(0));
        let _1828_v71;
        _1828_v71 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_this).f24),_1827_v70);
        let _1829_v72;
        _1829_v72 = _dafny.Seq.of(false);
        let _1830_v73;
        _1830_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1747_v0,(_this).f24);
        let _1831_v74;
        _1831_v74 = _dafny.MultiSet.fromElements(new BigNumber(618), new BigNumber((_1830_v73).length), (_dafny.ZERO).minus(new BigNumber(((_1819_v63)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_1819_v63).length))]).length)));
        let _1832_v75;
        _1832_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
        let _1833_v76;
        let _nw263 = Array((new BigNumber(13)).toNumber());
        _nw263[(_dafny.ZERO).toNumber()] = new BigNumber((_1828_v71).length);
        _nw263[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1829_v72, _dafny.Seq.of((_this).f24, (_this).f24, (_this).f24, false))).length);
        _nw263[(new BigNumber(2)).toNumber()] = (new BigNumber((_1821_v65).length)).minus(new BigNumber(983));
        _nw263[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_1747_v0);
        _nw263[(new BigNumber(4)).toNumber()] = _1747_v0;
        _nw263[(new BigNumber(5)).toNumber()] = new BigNumber((_1820_v64).length);
        _nw263[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(_1747_v0, _1747_v0);
        _nw263[(new BigNumber(7)).toNumber()] = (_1747_v0).multipliedBy((((_1831_v74).contains(new BigNumber((_dafny.Seq.UnicodeFromString("qrhauoaod")).length))) ? ((_1831_v74).get(new BigNumber((_dafny.Seq.UnicodeFromString("qrhauoaod")).length))) : (_1747_v0)));
        _nw263[(new BigNumber(8)).toNumber()] = _1747_v0;
        _nw263[(new BigNumber(9)).toNumber()] = new BigNumber(-508);
        _nw263[(new BigNumber(10)).toNumber()] = new BigNumber(((_1832_v75).Merge(_1832_v75)).length);
        _nw263[(new BigNumber(11)).toNumber()] = _1747_v0;
        _nw263[(new BigNumber(12)).toNumber()] = _1747_v0;
        _1833_v76 = _nw263;
        let _1834_v77;
        _1834_v77 = _module.D2.create_DC6(_1747_v0);
        let _index258 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_1833_v76).length));
        (_1833_v76)[_index258] = (_1834_v77).dtor_cf9;
        let _index259 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_1833_v76).length));
        let _rhs299 = (_1747_v0).plus(new BigNumber(((_1819_v63)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_1819_v63).length))]).length));
        let _rhs300 = ((_this).f24) || (!((_this).f24));
        let _lhs243 = _1833_v76;
        let _lhs244 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_1833_v76).length));
        _lhs243[_lhs244] = _rhs299;
        r0 = _rhs300;
        let _index260 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_1824_v68).length));
        (_1824_v68)[_index260] = (!((_this).f24)) || ((_this).f24);
        let _index261 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_1824_v68).length));
        let _rhs301 = _dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm52(globalState), _module.__default.safeIndex(new BigNumber((_1830_v73).length), new BigNumber((_module.__default.fm52(globalState)).length)), (_this).f24), _1829_v72), (_1829_v72)[_module.__default.safeIndex(_1747_v0, new BigNumber((_1829_v72).length))]);
        let _rhs302 = ((_1833_v76)[_module.__default.safeIndex(new BigNumber(711), new BigNumber((_1833_v76).length))]).minus(_1747_v0);
        let _rhs303 = _1747_v0;
        let _lhs245 = _1824_v68;
        let _lhs246 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_1824_v68).length));
        let _lhs247 = globalState;
        _lhs245[_lhs246] = _rhs301;
        r1 = _rhs302;
        _lhs247.f8 = _rhs303;
        _1824_v68 = _1824_v68;
      } else {
        r0 = (_this).f24;
        if ((_this).f24) {
          let _1835_v78;
          let _nw264 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1835_v78 = _nw264;
          let _1836_v79;
          _1836_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1835_v78,(_1747_v0).isEqualTo(_1747_v0));
          _1836_v79 = (_dafny.Map.Empty.slice().updateUnsafe(_1835_v78,false)).Merge(_1836_v79);
          (globalState).f10 = (_this).f24;
          let _1837_v80;
          _1837_v80 = _dafny.Seq.UnicodeFromString("dxmowc");
          _1837_v80 = _dafny.Seq.Concat(_1837_v80, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("w"), _module.__default.fm30(globalState)));
          let _1838_v81;
          let _nw265 = new _module.C5();
          _nw265.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_1747_v0,(_this).f24), new BigNumber(895));
          _1838_v81 = _nw265;
          let _1839_v82;
          let _nw266 = new _module.C2();
          _nw266.__ctor(_module.__default.safeDivisionInt(_1747_v0, _1747_v0), (_this).f22);
          _1839_v82 = _nw266;
        } else {
          let _1840_v83;
          let _nw267 = Array((new BigNumber(28)).toNumber()).fill([]);
          _1840_v83 = _nw267;
          let _1841_v85;
          let _init46 = ((_1842_v0) => function (_1843_i6) {
            return function () {
              let _coll72 = new _dafny.Map();
              for (const _compr_72 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1842_v0,(_this).f24)).length), _1842_v0, (_dafny.ZERO).minus(new BigNumber(-811)), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_1842_v0))).length), _1842_v0)).Elements) {
                let _1844_v84 = _compr_72;
                if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1842_v0,(_this).f24)).length), _1842_v0, (_dafny.ZERO).minus(new BigNumber(-811)), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_1842_v0))).length), _1842_v0), _1844_v84)) {
                  _coll72.push([_module.__default.safeModuloInt(_1844_v84, _1842_v0),_1842_v0]);
                }
              }
              return _coll72;
            }();
          })(_1747_v0);
          let _nw268 = Array((new BigNumber(18)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw268.length); _i0_46++) {
            _nw268[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1841_v85 = _nw268;
          let _index262 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_1840_v83).length));
          (_1840_v83)[_index262] = _1841_v85;
          let _1845_v86;
          _1845_v86 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1846_v87;
          let _nw269 = Array((new BigNumber(28)).toNumber()).fill(false);
          _1846_v87 = _nw269;
          let _index263 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1846_v87).length));
          (_1846_v87)[_index263] = (_this).f24;
          let _index264 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_1840_v83).length));
          let _index265 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1846_v87).length));
          let _rhs304 = _1841_v85;
          let _rhs305 = _module.__default.fm22((_this).f24, !((_this).f24), new BigNumber((_dafny.Seq.Concat(_module.__default.fm52(globalState), _module.__default.fm52(globalState))).length), (new BigNumber(684)).multipliedBy(_1747_v0), globalState);
          let _rhs306 = (_this).f24;
          let _lhs248 = _1840_v83;
          let _lhs249 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_1840_v83).length));
          let _lhs250 = _1846_v87;
          let _lhs251 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_1846_v87).length));
          _lhs248[_lhs249] = _rhs304;
          _1845_v86 = _rhs305;
          _lhs250[_lhs251] = _rhs306;
          let _1847_v88;
          _1847_v88 = _dafny.Set.fromElements(_1747_v0, _1747_v0);
          let _1848_v89;
          _1848_v89 = _dafny.MultiSet.fromElements(_1747_v0, _1747_v0);
          let _1849_v90;
          _1849_v90 = _module.D6.create_DC23((_this).f24, _1747_v0, (_1846_v87)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_1846_v87).length))]);
          let _1850_v91;
          _1850_v91 = _dafny.Seq.of(_1847_v88);
          let _1851_v92;
          _1851_v92 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm36((_this).f24, _1847_v88, _1747_v0, (_dafny.ZERO).minus((_this).fm0(_1848_v89, (_1849_v90).dtor_cf38, _1850_v91, (_1846_v87)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_1846_v87).length))], globalState)), globalState),(_1747_v0).plus(_1747_v0));
          let _1852_v93;
          _1852_v93 = _dafny.Seq.UnicodeFromString("uji");
          _1851_v92 = (_1851_v92).update(_1852_v93, (_dafny.ZERO).minus(_1747_v0));
          let _1853_v94;
          let _init47 = ((_1854_v87) => function (_1855_i7) {
            return _module.__default.safeModuloInt(_1855_i7, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1854_v87)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_1854_v87).length))],true)).length));
          })(_1846_v87);
          let _nw270 = Array((new BigNumber(16)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw270.length); _i0_47++) {
            _nw270[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1853_v94 = _nw270;
          let _index266 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1853_v94).length));
          (_1853_v94)[_index266] = (new BigNumber(713)).plus(new BigNumber((_module.__default.fm27(_1747_v0, globalState)).length));
          let _index267 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1853_v94).length));
          (_1853_v94)[_index267] = (_1747_v0).multipliedBy(_1747_v0);
          _1853_v94 = _1853_v94;
          let _1856_v95;
          _1856_v95 = _dafny.Map.Empty.slice().updateUnsafe((_1846_v87)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_1846_v87).length))],new BigNumber((_1852_v93).length));
          let _1857_v96;
          _1857_v96 = _dafny.Set.fromElements(_module.__default.fm15(false, globalState));
          let _1858_v97;
          _1858_v97 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1856_v95).length),_1857_v96);
          let _rhs307 = !((_this).f24);
          let _rhs308 = _dafny.Seq.IsPrefixOf(_1852_v93, _1852_v93);
          let _rhs309 = ((_1858_v97).Merge(_1858_v97)).update((_1853_v94)[_module.__default.safeIndex(new BigNumber(528), new BigNumber((_1853_v94).length))], _1857_v96);
          let _rhs310 = _dafny.Seq.Concat(_1852_v93, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-748)), ((_1859_v86) => function (_1860_i8) {
            return _1859_v86;
          })(_1845_v86)));
          let _lhs252 = globalState;
          r0 = _rhs307;
          _lhs252.f10 = _rhs308;
          _1858_v97 = _rhs309;
          _1852_v93 = _rhs310;
        }
        (globalState).f0 = (_this).f24;
        let _1861_v98;
        let _init48 = function (_1862_i9) {
          return _dafny.Set.fromElements(!((_this).f24), !(!((_this).f24)));
        };
        let _nw271 = Array((new BigNumber(26)).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw271.length); _i0_48++) {
          _nw271[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _1861_v98 = _nw271;
        let _1863_v99;
        _1863_v99 = _dafny.Set.fromElements((_this).f24);
        let _1864_v100;
        _1864_v100 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1863_v99);
        let _index268 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_1861_v98).length));
        (_1861_v98)[_index268] = (((_1864_v100).contains((_this).f24)) ? ((_1864_v100).get((_this).f24)) : (_1863_v99));
        let _index269 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_1861_v98).length));
        (_1861_v98)[_index269] = _1863_v99;
        let _1865_v101;
        _1865_v101 = _dafny.MultiSet.fromElements(_1747_v0);
        let _1866_v102;
        _1866_v102 = _dafny.Set.fromElements(_1865_v101);
        let _rhs311 = _1747_v0;
        let _rhs312 = ((_this).f24) || ((_1866_v102).IsDisjointFrom(_1866_v102));
        let _lhs253 = globalState;
        let _lhs254 = globalState;
        _lhs253.f6 = _rhs311;
        _lhs254.f10 = _rhs312;
      }
      let _1867_v103;
      _1867_v103 = _dafny.Set.fromElements(((_this).f24) === ((_this).f24), (_this).f24, (_this).f24);
      _1867_v103 = _1867_v103;
      let _1868_v104;
      _1868_v104 = _module.D15.create_DC51((_this).f24, (_this).f24);
      let _source25 = (((_this).f24) ? (_1868_v104) : (_1868_v104));
      if (_source25.is_DC50) {
        (globalState).f6 = _1747_v0;
        let _1869_v105;
        _1869_v105 = _dafny.Seq.of((_this).f24);
        let _1870_v106;
        _1870_v106 = _dafny.MultiSet.fromElements(false, (_this).f24, (_this).f24, (_this).f24, (_this).f24);
        let _1871_v107;
        _1871_v107 = new _dafny.CodePoint('x'.codePointAt(0));
        let _1872_v108;
        _1872_v108 = _module.D10.create_DC34(_1871_v107, (_this).f24, (_this).f24, (_this).f24, (_this).f24);
        let _1873_v109;
        _1873_v109 = _module.D11.create_DC39((_this).f24);
        let _1874_v110;
        _1874_v110 = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_1869_v105)).cardinality()));
        let _1875_v111;
        _1875_v111 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1874_v110).length),(_this).f24);
        let _1876_v112;
        _1876_v112 = _1875_v111;
        let _1877_v113;
        _1877_v113 = _dafny.Seq.UnicodeFromString("lsxnhwfp");
        let _1878_v114;
        let _nw272 = Array((new BigNumber(22)).toNumber());
        _nw272[(_dafny.ZERO).toNumber()] = (_1869_v105)[_module.__default.safeIndex((((_1870_v106).contains((_this).f24)) ? ((_1870_v106).get((_this).f24)) : (_1747_v0)), new BigNumber((_1869_v105).length))];
        _nw272[(_dafny.ONE).toNumber()] = (_this).f24;
        _nw272[(new BigNumber(2)).toNumber()] = (_this).f24;
        _nw272[(new BigNumber(3)).toNumber()] = (_this).f24;
        _nw272[(new BigNumber(4)).toNumber()] = (((_this).f24) ? ((_this).f24) : ((_this).f24));
        _nw272[(new BigNumber(5)).toNumber()] = (((_this).f24) ? ((_this).f24) : ((_this).f24));
        _nw272[(new BigNumber(6)).toNumber()] = true;
        _nw272[(new BigNumber(7)).toNumber()] = (_this).f24;
        _nw272[(new BigNumber(8)).toNumber()] = (_this).f24;
        _nw272[(new BigNumber(9)).toNumber()] = (_this).f24;
        _nw272[(new BigNumber(10)).toNumber()] = (_this).f24;
        _nw272[(new BigNumber(11)).toNumber()] = (_this).f24;
        _nw272[(new BigNumber(12)).toNumber()] = !((_1872_v108).dtor_cf57);
        _nw272[(new BigNumber(13)).toNumber()] = (_this).f24;
        _nw272[(new BigNumber(14)).toNumber()] = (_1747_v0).isLessThan(_1747_v0);
        _nw272[(new BigNumber(15)).toNumber()] = (_1873_v109).dtor_cf66;
        _nw272[(new BigNumber(16)).toNumber()] = !((_this).f24);
        _nw272[(new BigNumber(17)).toNumber()] = (_1747_v0).isLessThanOrEqualTo(new BigNumber(-636));
        _nw272[(new BigNumber(18)).toNumber()] = (((_this).f24) ? ((_this).f24) : ((_this).f24));
        _nw272[(new BigNumber(19)).toNumber()] = true;
        _nw272[(new BigNumber(20)).toNumber()] = _module.__default.fm37(_1869_v105, (_1876_v112), _1877_v113, globalState);
        _nw272[(new BigNumber(21)).toNumber()] = true;
        _1878_v114 = _nw272;
        let _index270 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_1878_v114).length));
        (_1878_v114)[_index270] = !((_this).f24);
        let _index271 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_1878_v114).length));
        (_1878_v114)[_index271] = (_1874_v110).IsSubsetOf(_1874_v110);
        let _1879_v115;
        let _nw273 = new _module.C3();
        _nw273.__ctor((_this).f22);
        _1879_v115 = _nw273;
        _1879_v115 = _1879_v115;
        _1747_v0 = new BigNumber(861);
      } else if (_source25.is_DC51) {
        let _1880___mcc_h0 = (_source25).cf82;
        let _1881___mcc_h1 = (_source25).cf83;
        let _1882_cf83 = _1881___mcc_h1;
        let _1883_cf82 = _1880___mcc_h0;
        let _1884_v116;
        _1884_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1882_cf83,false);
        let _1885_v117;
        let _init49 = function (_1886_i10) {
          return (_1886_i10).minus(new BigNumber((_dafny.Seq.UnicodeFromString("pvvl")).length));
        };
        let _nw274 = Array((new BigNumber(4)).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw274.length); _i0_49++) {
          _nw274[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _1885_v117 = _nw274;
        let _1887_v118;
        _1887_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1885_v117,_1882_cf83);
        _1884_v116 = (_1884_v116).update((((_1887_v118).contains(_1885_v117)) ? ((_1887_v118).get(_1885_v117)) : ((_this).f24)), _1882_cf83);
        let _1888_v119;
        let _nw275 = new _module.C1();
        _nw275.__ctor(new BigNumber(779), (_this).f22);
        _1888_v119 = _nw275;
        let _1889_v120;
        _1889_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1747_v0,_1888_v119);
        let _1890_v121;
        _1890_v121 = _dafny.Seq.of(true, !((_this).f24), false);
        let _1891_v122;
        let _nw276 = Array((new BigNumber(23)).toNumber());
        _nw276[(_dafny.ZERO).toNumber()] = _1888_v119;
        _nw276[(_dafny.ONE).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(2)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(3)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(4)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(5)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(6)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(7)).toNumber()] = (((_1889_v120).contains(new BigNumber((_1890_v121).length))) ? ((_1889_v120).get(new BigNumber((_1890_v121).length))) : (_1888_v119));
        _nw276[(new BigNumber(8)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(9)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(10)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(11)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(12)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(13)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(14)).toNumber()] = (((_this).f24) ? (_1888_v119) : (_1888_v119));
        _nw276[(new BigNumber(15)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(16)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(17)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(18)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(19)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(20)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(21)).toNumber()] = _1888_v119;
        _nw276[(new BigNumber(22)).toNumber()] = _1888_v119;
        _1891_v122 = _nw276;
        let _index272 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1891_v122).length));
        (_1891_v122)[_index272] = _1888_v119;
        let _index273 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1891_v122).length));
        (_1891_v122)[_index273] = _1888_v119;
        _1883_cf82 = !(_1882_cf83) || (false);
        let _1892_v123;
        let _nw277 = new _module.C2();
        _nw277.__ctor(_1747_v0, (_this).f22);
        _1892_v123 = _nw277;
        _1892_v123 = _1892_v123;
      } else if (_source25.is_DC52) {
        let _1893_v124;
        let _nw278 = Array((new BigNumber(15)).toNumber()).fill(false);
        _1893_v124 = _nw278;
        let _index274 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length));
        (_1893_v124)[_index274] = true;
        let _index275 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length));
        (_1893_v124)[_index275] = (new BigNumber(312)).isLessThan(_1747_v0);
        let _1894_v125;
        _1894_v125 = _dafny.Seq.UnicodeFromString("elp");
        let _1895_v126;
        _1895_v126 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1896_v127;
        _1896_v127 = _dafny.Map.Empty.slice().updateUnsafe((_1893_v124)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length))],_1895_v126);
        let _1897_v128;
        _1897_v128 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(486), (_dafny.ZERO).minus(new BigNumber((_1894_v125).length))),((false) ? (_1896_v127) : (_1896_v127)));
        let _1898_v129;
        _1898_v129 = _dafny.MultiSet.fromElements(_1867_v103);
        let _1899_v130;
        _1899_v130 = _dafny.Seq.of(new BigNumber((_1898_v129).cardinality()), (_dafny.ZERO).minus(_1747_v0));
        let _1900_v131;
        _1900_v131 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!((_1893_v124)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length))]),_1895_v126));
        let _1901_v132;
        _1901_v132 = _dafny.MultiSet.fromElements(_1747_v0, _1747_v0);
        _1897_v128 = (_1897_v128).update(_1899_v130, (_1900_v131)[_module.__default.safeIndex(new BigNumber((_1901_v132).cardinality()), new BigNumber((_1900_v131).length))]);
        if ((_this).f24) {
          let _1902_v133;
          let _nw279 = new _module.C3();
          _nw279.__ctor(_dafny.Seq.update((_this).f22, _module.__default.safeIndex(_1747_v0, new BigNumber(((_this).f22).length)), _1893_v124));
          _1902_v133 = _nw279;
          let _1903_v134;
          let _nw280 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _1903_v134 = _nw280;
          let _index276 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1903_v134).length));
          (_1903_v134)[_index276] = (_1747_v0).minus(_1747_v0);
          let _1904_v135;
          _1904_v135 = _dafny.Seq.of((_this).f24);
          let _1905_v136;
          _1905_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1747_v0,(_1893_v124)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length))]);
          let _index277 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1903_v134).length));
          let _rhs313 = _dafny.Seq.of(((false) ? ((_1893_v124)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length))]) : (_module.__default.fm37(_1904_v135, _1905_v136, _1894_v125, globalState))), ((_this).f24) || ((_this).f24));
          let _rhs314 = (_1747_v0).minus(_1747_v0);
          let _rhs315 = _1747_v0;
          let _rhs316 = _1893_v124;
          let _rhs317 = _1747_v0;
          let _lhs255 = globalState;
          let _lhs256 = _1903_v134;
          let _lhs257 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_1903_v134).length));
          let _lhs258 = globalState;
          _lhs255.f5 = _rhs313;
          _lhs256[_lhs257] = _rhs314;
          _lhs258.f8 = _rhs315;
          _1893_v124 = _rhs316;
          _1747_v0 = _rhs317;
          _1895_v126 = _1895_v126;
          r1 = (_1899_v130)[_module.__default.safeIndex((_1903_v134)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_1903_v134).length))], new BigNumber((_1899_v130).length))];
          _1895_v126 = _1895_v126;
        } else {
          let _1906_v137;
          _1906_v137 = _dafny.Seq.of((_1893_v124)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length))]);
          let _1907_v138;
          _1907_v138 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((new BigNumber(-419)).minus((_dafny.ZERO).minus(_1747_v0))),_dafny.Seq.Concat(_1906_v137, _1906_v137));
          let _1908_v139;
          let _nw281 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _1908_v139 = _nw281;
          let _index278 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1908_v139).length));
          (_1908_v139)[_index278] = _1747_v0;
          let _1909_v140;
          _1909_v140 = _dafny.Map.Empty.slice().updateUnsafe(_1895_v126,_1747_v0);
          let _1910_v141;
          _1910_v141 = _dafny.Seq.of(_1909_v140);
          let _1911_v142;
          _1911_v142 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(367),(_1893_v124)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length))]);
          let _index279 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1908_v139).length));
          let _rhs318 = _1906_v137;
          let _rhs319 = (_1907_v138).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1747_v0,_1906_v137));
          let _rhs320 = _module.__default.safeModuloInt(new BigNumber(((_1910_v141)[_module.__default.safeIndex(_1747_v0, new BigNumber((_1910_v141).length))]).length), new BigNumber((_1911_v142).length));
          let _rhs321 = !(new BigNumber((_1901_v132).cardinality())).isEqualTo(_module.__default.safeDivisionInt(new BigNumber((_1894_v125).length), _1747_v0));
          let _lhs259 = globalState;
          let _lhs260 = _1908_v139;
          let _lhs261 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1908_v139).length));
          _lhs259.f5 = _rhs318;
          _1907_v138 = _rhs319;
          _lhs260[_lhs261] = _rhs320;
          r0 = _rhs321;
          let _1912_v143;
          _1912_v143 = _dafny.Seq.of(_1894_v125, _1894_v125);
          let _1913_v144;
          _1913_v144 = _dafny.Map.Empty.slice().updateUnsafe((_1893_v124)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length))],(_this).f24);
          let _1914_v145;
          _1914_v145 = _module.D9.create_DC31(new BigNumber((_1913_v144).length), _1893_v124);
          let _1915_v146;
          _1915_v146 = _module.D9.create_DC32(_1914_v145);
          let _1916_v147;
          _1916_v147 = _dafny.Map.Empty.slice().updateUnsafe(_1912_v143,_module.D9.create_DC32(_1915_v146));
          let _1917_v148;
          _1917_v148 = _module.D9.create_DC32((((_1916_v147).contains(_1912_v143)) ? ((_1916_v147).get(_1912_v143)) : (_1915_v146)));
          let _1918_v149;
          _1918_v149 = _module.D9.create_DC32(_1914_v145);
          let _1919_v150;
          _1919_v150 = _dafny.Map.Empty.slice().updateUnsafe(_1918_v149,(_this).f24);
          let _1920_v151;
          _1920_v151 = _dafny.Seq.of(_1919_v150, _dafny.Map.Empty.slice().updateUnsafe(_1918_v149,(_this).f24), (_1919_v150).Merge(_1919_v150), _1919_v150);
          let _pat_let_tv67 = _1915_v146;
          let _rhs322 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(function (_pat_let38_0) {
            return function (_1921_dt__update__tmp_h1) {
              return function (_pat_let39_0) {
                return function (_1922_dt__update_hcf53_h0) {
                  return _module.D9.create_DC32(_1922_dt__update_hcf53_h0);
                }(_pat_let39_0);
              }(_pat_let_tv67);
            }(_pat_let38_0);
          }(_module.D9.create_DC32(_1915_v146)),true)).Merge(_1919_v150));
          let _rhs323 = (_this).f24;
          let _rhs324 = new _dafny.CodePoint('m'.codePointAt(0));
          let _lhs262 = globalState;
          _1920_v151 = _rhs322;
          _lhs262.f10 = _rhs323;
          _1895_v126 = _rhs324;
          let _1923_v152;
          _1923_v152 = _dafny.MultiSet.fromElements(true);
          _1895_v126 = (_1894_v125)[_module.__default.safeIndex((((_1923_v152).contains((_1893_v124)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length))])) ? ((_1923_v152).get((_1893_v124)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_1893_v124).length))])) : ((_1908_v139)[_module.__default.safeIndex(new BigNumber(802), new BigNumber((_1908_v139).length))])), new BigNumber((_1894_v125).length))];
          _1894_v125 = _1894_v125;
          let _1924_v153;
          _1924_v153 = _dafny.Set.fromElements(_1747_v0);
          (globalState).f13 = _1924_v153;
        }
        let _1925_v154;
        let _init50 = function (_1926_i11) {
          return _dafny.Seq.of((_this).f24, (_this).f24);
        };
        let _nw282 = Array((new BigNumber(16)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw282.length); _i0_50++) {
          _nw282[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _1925_v154 = _nw282;
        let _1927_v155;
        _1927_v155 = _dafny.Map.Empty.slice().updateUnsafe(_1899_v130,_1925_v154);
        let _1928_v156;
        let _nw283 = Array((new BigNumber(26)).toNumber());
        _nw283[(_dafny.ZERO).toNumber()] = _1925_v154;
        _nw283[(_dafny.ONE).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(2)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(3)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(4)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(5)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(6)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(7)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(8)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(9)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(10)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(11)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(12)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(13)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(14)).toNumber()] = (((_this).f24) ? (_1925_v154) : ((((_1927_v155).contains(_1899_v130)) ? ((_1927_v155).get(_1899_v130)) : (_1925_v154))));
        _nw283[(new BigNumber(15)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(16)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(17)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(18)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(19)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(20)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(21)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(22)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(23)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(24)).toNumber()] = _1925_v154;
        _nw283[(new BigNumber(25)).toNumber()] = _1925_v154;
        _1928_v156 = _nw283;
        let _index280 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1928_v156).length));
        (_1928_v156)[_index280] = _1925_v154;
        let _index281 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_1928_v156).length));
        (_1928_v156)[_index281] = _1925_v154;
      } else if (_source25.is_DC49) {
        let _1929___mcc_h2 = (_source25).cf81;
        let _1930_cf81 = _1929___mcc_h2;
        let _1931_v157;
        let _nw284 = new _module.C3();
        _nw284.__ctor(_dafny.Seq.Concat((_this).f22, (_this).f22));
        _1931_v157 = _nw284;
        let _1932_v158;
        let _nw285 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
        _1932_v158 = _nw285;
        _1932_v158 = _1932_v158;
        let _rhs325 = (_this).f24;
        let _rhs326 = ((new BigNumber(910)).minus(new BigNumber(218))).plus(_1747_v0);
        let _lhs263 = globalState;
        let _lhs264 = globalState;
        _lhs263.f0 = _rhs325;
        _lhs264.f16 = _rhs326;
        let _1933_v159;
        _1933_v159 = _dafny.MultiSet.fromElements(_1747_v0, _1747_v0, _1747_v0, _1747_v0);
        let _1934_v160;
        _1934_v160 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(974)), function (_1935_i13) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }));
        (globalState).f6 = (_module.__default.safeDivisionInt(_1747_v0, (_1931_v157).fm0(_1933_v159, _1747_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(861)), ((_1936_v0) => function (_1937_i12) {
          return _dafny.Set.fromElements(_1936_v0, _1936_v0);
        })(_1747_v0)), (_this).f24, globalState))).minus(new BigNumber((_1934_v160).length));
      } else {
        let _1938___mcc_h3 = (_source25).cf84;
        let _1939_cf84 = _1938___mcc_h3;
        let _1940_v161;
        let _nw286 = new _module.C1();
        _nw286.__ctor(_1747_v0, (_this).f22);
        _1940_v161 = _nw286;
        let _1941_v162;
        _1941_v162 = _dafny.Seq.UnicodeFromString("wrcpc");
        let _1942_v163;
        _1942_v163 = _module.D1.create_DC3(_1941_v162);
        let _source26 = _1942_v163;
        if (_source26.is_DC3) {
          let _1943___mcc_h4 = (_source26).cf6;
          let _1944_cf6 = _1943___mcc_h4;
          let _1945_v164;
          _1945_v164 = _module.D4.create_DC15(new BigNumber((_1941_v162).length), (_dafny.ZERO).minus((_1940_v161).fm34((_this).f24, globalState)));
          _1945_v164 = _1945_v164;
          _1941_v162 = _1944_cf6;
          (globalState).f0 = (_this).f24;
          let _1946_v165;
          _1946_v165 = _dafny.Map.Empty.slice().updateUnsafe(_1944_cf6,true);
          _1946_v165 = (_1946_v165).update(_1944_cf6, (_this).f24);
        } else if (_source26.is_DC2) {
          let _1947___mcc_h5 = (_source26).cf5;
          let _1948_cf5 = _1947___mcc_h5;
          _1948_cf5 = _dafny.Seq.UnicodeFromString("nqyii");
          let _1949_v166;
          let _nw287 = new _module.C3();
          _nw287.__ctor((_this).f22);
          _1949_v166 = _nw287;
          let _1950_v167;
          let _nw288 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _1950_v167 = _nw288;
          let _1951_v168;
          _1951_v168 = _dafny.Set.fromElements(_1950_v167);
          let _1952_v169;
          _1952_v169 = _module.D13.create_DC46((_this).f24, _1949_v166, (_1940_v161).f30, _1951_v168);
          let _1953_v170;
          _1953_v170 = _dafny.Map.Empty.slice().updateUnsafe(_1948_cf5,(_1940_v161).f30);
          let _1954_v171;
          _1954_v171 = _dafny.Map.Empty.slice().updateUnsafe(_1747_v0,(_1940_v161).f30);
          let _1955_v172;
          _1955_v172 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1954_v171).length)), (_1940_v161).f30, (_1940_v161).f30);
          let _1956_v173;
          let _nw289 = Array((new BigNumber(10)).toNumber());
          _nw289[(_dafny.ZERO).toNumber()] = (_this).f24;
          _nw289[(_dafny.ONE).toNumber()] = (_1952_v169).dtor_cf73;
          _nw289[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of((((_1953_v170).contains(_dafny.Seq.UnicodeFromString("os"))) ? ((_1953_v170).get(_dafny.Seq.UnicodeFromString("os"))) : ((_1940_v161).f30))), _1955_v172);
          _nw289[(new BigNumber(3)).toNumber()] = true;
          _nw289[(new BigNumber(4)).toNumber()] = false;
          _nw289[(new BigNumber(5)).toNumber()] = (_this).f24;
          _nw289[(new BigNumber(6)).toNumber()] = (_this).f24;
          _nw289[(new BigNumber(7)).toNumber()] = (_this).f24;
          _nw289[(new BigNumber(8)).toNumber()] = !((_this).f24) || ((_this).f24);
          _nw289[(new BigNumber(9)).toNumber()] = _module.__default.fm15((_this).f24, globalState);
          _1956_v173 = _nw289;
          let _1957_v174;
          _1957_v174 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,true);
          let _1958_v175;
          _1958_v175 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1957_v174).length),(_this).f24);
          let _index282 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1956_v173).length));
          (_1956_v173)[_index282] = !(_module.__default.fm37(_dafny.Seq.of((_this).f24, false, (_this).f24), _1958_v175, _dafny.Seq.UnicodeFromString("a"), globalState)) || ((_this).f24);
          let _index283 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1956_v173).length));
          (_1956_v173)[_index283] = (_this).f24;
          let _1959_v176;
          _1959_v176 = _dafny.Map.Empty.slice().updateUnsafe(_1958_v175,(((_this).f24) ? ((_1956_v173)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_1956_v173).length))]) : (_module.__default.fm15(false, globalState))));
          _1959_v176 = (_1959_v176).update(_1958_v175, (_1956_v173)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_1956_v173).length))]);
          let _index284 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1950_v167).length));
          (_1950_v167)[_index284] = (_1940_v161).f30;
          let _1960_v177;
          _1960_v177 = _dafny.MultiSet.fromElements(_1747_v0);
          let _1961_v178;
          _1961_v178 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-361)), ((_1962_v161) => function (_1963_i14) {
            return (_1962_v161).f30;
          })(_1940_v161)));
          let _1964_v179;
          _1964_v179 = _dafny.Set.fromElements((_1940_v161).f30);
          let _1965_v180;
          _1965_v180 = _dafny.Seq.of(_1964_v179);
          let _1966_v181;
          _1966_v181 = new _dafny.CodePoint('k'.codePointAt(0));
          let _index285 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1950_v167).length));
          let _index286 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1956_v173).length));
          let _rhs327 = ((_this).f24) || ((_1956_v173)[_module.__default.safeIndex(new BigNumber(697), new BigNumber((_1956_v173).length))]);
          let _rhs328 = (_dafny.ZERO).minus((_1949_v166).fm0(_1960_v177, new BigNumber((_1961_v178).length), _1965_v180, (_this).f24, globalState));
          let _rhs329 = _1747_v0;
          let _rhs330 = _1747_v0;
          let _rhs331 = _dafny.Seq.contains(_1948_cf5, _1966_v181);
          let _lhs265 = globalState;
          let _lhs266 = _1950_v167;
          let _lhs267 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_1950_v167).length));
          let _lhs268 = globalState;
          let _lhs269 = globalState;
          let _lhs270 = _1956_v173;
          let _lhs271 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1956_v173).length));
          _lhs265.f0 = _rhs327;
          _lhs266[_lhs267] = _rhs328;
          _lhs268.f11 = _rhs329;
          _lhs269.f6 = _rhs330;
          _lhs270[_lhs271] = _rhs331;
        } else {
          let _1967___mcc_h6 = (_source26).cf7;
          let _1968_cf7 = _1967___mcc_h6;
          (globalState).f8 = _1747_v0;
          let _1969_v182;
          _1969_v182 = _module.D1.create_DC3(_dafny.Seq.UnicodeFromString("cokvtw"));
          let _1970_v183;
          _1970_v183 = _module.D1.create_DC4(_1969_v182);
          let _1971_v184;
          _1971_v184 = _module.D1.create_DC4(_1969_v182);
          let _1972_v185;
          _1972_v185 = _module.D1.create_DC4(_1970_v183);
          let _1973_v186;
          let _nw290 = Array((_dafny.ONE).toNumber());
          _nw290[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1941_v162, _1941_v162);
          _1973_v186 = _nw290;
          let _1974_v187;
          _1974_v187 = new _dafny.CodePoint('y'.codePointAt(0));
          let _index287 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_1973_v186).length));
          (_1973_v186)[_index287] = _dafny.Seq.update(_dafny.Seq.update(_1941_v162, _module.__default.safeIndex((_1940_v161).f30, new BigNumber((_1941_v162).length)), _1974_v187), _module.__default.safeIndex(_1747_v0, new BigNumber((_dafny.Seq.update(_1941_v162, _module.__default.safeIndex((_1940_v161).f30, new BigNumber((_1941_v162).length)), _1974_v187)).length)), _1974_v187);
          let _1975_v188;
          let _init51 = ((_1976_v161) => function (_1977_i15) {
            return (_1977_i15).minus((_1976_v161).f30);
          })(_1940_v161);
          let _nw291 = Array((new BigNumber(16)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw291.length); _i0_51++) {
            _nw291[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1975_v188 = _nw291;
          let _index288 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_1975_v188).length));
          (_1975_v188)[_index288] = new BigNumber((_dafny.Seq.UnicodeFromString("krpsp")).length);
          let _1978_v189;
          _1978_v189 = _dafny.MultiSet.fromElements(new BigNumber((_1941_v162).length), new BigNumber(999), _1747_v0);
          let _1979_v190;
          _1979_v190 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(391)), ((_1980_v187) => function (_1981_i16) {
            return _1980_v187;
          })(_1974_v187))).length));
          let _1982_v191;
          _1982_v191 = _dafny.Map.Empty.slice().updateUnsafe((((_1978_v189).contains(new BigNumber(232))) ? ((_1978_v189).get(new BigNumber(232))) : ((_1940_v161).f30)),_1979_v190);
          let _index289 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_1973_v186).length));
          let _index290 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_1975_v188).length));
          let _rhs332 = _module.__default.safeDivisionInt((_1940_v161).f30, (_1940_v161).f30);
          let _rhs333 = _module.D1.create_DC4(_1970_v183);
          let _rhs334 = _dafny.Seq.UnicodeFromString("cbjw");
          let _rhs335 = (_1940_v161).f30;
          let _rhs336 = function () {
            let _coll73 = new _dafny.Map();
            for (const _compr_73 of _dafny.IntegerRange(new BigNumber(249), new BigNumber(61))) {
              let _1983_v192 = _compr_73;
              if (((new BigNumber(249)).isLessThanOrEqualTo(_1983_v192)) && ((_1983_v192).isLessThan(new BigNumber(61)))) {
                _coll73.push([_module.__default.safeDivisionInt(_1983_v192, (_1940_v161).f30),((_1979_v190).update((_this).f24, _1747_v0)).Merge(_1979_v190)]);
              }
            }
            return _coll73;
          }();
          let _lhs272 = globalState;
          let _lhs273 = _1973_v186;
          let _lhs274 = _module.__default.safeIndex(new BigNumber(293), new BigNumber((_1973_v186).length));
          let _lhs275 = _1975_v188;
          let _lhs276 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_1975_v188).length));
          _lhs272.f16 = _rhs332;
          _1972_v185 = _rhs333;
          _lhs273[_lhs274] = _rhs334;
          _lhs275[_lhs276] = _rhs335;
          _1982_v191 = _rhs336;
          let _1984_v193;
          _1984_v193 = _module.D11.create_DC36((_this).f25);
          let _1985_v194;
          _1985_v194 = _module.D11.create_DC40(_1984_v193);
          let _1986_v195;
          _1986_v195 = _dafny.Seq.of(_1985_v194);
          let _1987_v196;
          _1987_v196 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1986_v195, _1986_v195),_1747_v0);
          _1987_v196 = (_1987_v196).update(_dafny.Seq.Concat(_1986_v195, _dafny.Seq.of(_1985_v194)), (_1940_v161).f30);
          let _1988_v197;
          let _1989_v198;
          let _1990_v199;
          let _1991_v200;
          let _out64;
          let _out65;
          let _out66;
          let _out67;
          let _outcollector20 = (_1940_v161).m6(new BigNumber(104), (_this).f24, ((_1975_v188)[_module.__default.safeIndex(new BigNumber(345), new BigNumber((_1975_v188).length))]).multipliedBy((_1940_v161).f30), (_1940_v161).f30, globalState);
          _out64 = _outcollector20[0];
          _out65 = _outcollector20[1];
          _out66 = _outcollector20[2];
          _out67 = _outcollector20[3];
          _1988_v197 = _out64;
          _1989_v198 = _out65;
          _1990_v199 = _out66;
          _1991_v200 = _out67;
        }
        let _1992_v201;
        let _nw292 = Array((new BigNumber(2)).toNumber());
        _1992_v201 = _nw292;
        let _1993_v202;
        _1993_v202 = _module.D11.create_DC37(_1992_v201);
        let _1994_v203;
        _1994_v203 = _module.D11.create_DC40(_1993_v202);
        let _pat_let_tv68 = _1993_v202;
        let _1995_v204;
        _1995_v204 = _dafny.Seq.of(_1994_v203, function (_pat_let40_0) {
          return function (_1996_dt__update__tmp_h2) {
            return function (_pat_let41_0) {
              return function (_1997_dt__update_hcf67_h0) {
                return _module.D11.create_DC40(_1997_dt__update_hcf67_h0);
              }(_pat_let41_0);
            }(_pat_let_tv68);
          }(_pat_let40_0);
        }(_1994_v203), _1994_v203, _1994_v203, function (_pat_let42_0) {
          return function (_1998_dt__update__tmp_h3) {
            return function (_pat_let43_0) {
              return function (_1999_dt__update_hcf67_h1) {
                return _module.D11.create_DC40(_1999_dt__update_hcf67_h1);
              }(_pat_let43_0);
            }(_module.D11.create_DC39(!((_this).f24)));
          }(_pat_let42_0);
        }(_module.D11.create_DC40(_1993_v202)));
        _1995_v204 = _dafny.Seq.Concat(_1995_v204, _1995_v204);
        let _2000_v205;
        _2000_v205 = _dafny.Map.Empty.slice().updateUnsafe(_module.D15.create_DC50(),(_this).f24);
        let _2001_v206;
        _2001_v206 = _module.D15.create_DC50();
        _2000_v205 = (_2000_v205).update(_2001_v206, (_this).f24);
      }
      r0 = (_this).f24;
      r1 = _1747_v0;
      return [r0, r1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f22 = _dafny.Seq.of();
      this._f23 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
    __ctor(f23, f22) {
      let _this = this;
      (_this)._f23 = f23;
      (_this)._f22 = f22;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(((new BigNumber(((_this).f23).length)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(new BigNumber(44)))).length))).multipliedBy(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(((_this).f23).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(134),false)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-324)), function (_2002_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).length))).length))), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(313)), function (_2003_i1) {
        return new BigNumber(176);
      })).length)))).length)));
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return new _dafny.CodePoint('q'.codePointAt(0));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.ZERO;
      let _2004_v0;
      _2004_v0 = _dafny.Seq.of(true, p0);
      let _2005_v1;
      _2005_v1 = new BigNumber(840);
      let _2006_v2;
      _2006_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2004_v0,_dafny.Set.fromElements(_2005_v1));
      if (!((((_2006_v2).contains(_module.__default.fm2(globalState))) ? ((_2006_v2).get(_module.__default.fm2(globalState))) : (_dafny.Set.fromElements(_2005_v1, _2005_v1)))).equals((_this).f23)) {
        let _2007_v3;
        _2007_v3 = _dafny.MultiSet.fromElements(_2005_v1);
        let _2008_v4;
        _2008_v4 = _dafny.Seq.of(_2007_v3, _2007_v3, _dafny.MultiSet.fromElements(_2005_v1, new BigNumber(285)), (_2007_v3).update(_2005_v1, _module.__default.abs(_2005_v1)));
        let _2009_v5;
        let _nw293 = new _module.C6();
        _nw293.__ctor(p0, _2008_v4, (_this).f22);
        _2009_v5 = _nw293;
        let _2010_v6;
        _2010_v6 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-446)), function (_2011_i0) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }));
        let _2012_v7;
        _2012_v7 = _dafny.Seq.UnicodeFromString("wp");
        let _2013_v8;
        let _init52 = ((_2014_p0) => function (_2015_i1) {
          return _2014_p0;
        })(p0);
        let _nw294 = Array((_dafny.ONE).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw294.length); _i0_52++) {
          _nw294[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _2013_v8 = _nw294;
        let _2016_v9;
        _2016_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2005_v1,_2013_v8);
        let _2017_v10;
        let _nw295 = Array((new BigNumber(13)).toNumber());
        _nw295[(_dafny.ZERO).toNumber()] = _2005_v1;
        _nw295[(_dafny.ONE).toNumber()] = _2005_v1;
        _nw295[(new BigNumber(2)).toNumber()] = _2005_v1;
        _nw295[(new BigNumber(3)).toNumber()] = (((_2010_v6).contains(_2012_v7)) ? ((_2010_v6).get(_2012_v7)) : (_2005_v1));
        _nw295[(new BigNumber(4)).toNumber()] = _2005_v1;
        _nw295[(new BigNumber(5)).toNumber()] = _2005_v1;
        _nw295[(new BigNumber(6)).toNumber()] = _module.__default.fm7(_2005_v1, globalState);
        _nw295[(new BigNumber(7)).toNumber()] = new BigNumber(((_2016_v9).update(_2005_v1, _2013_v8)).length);
        _nw295[(new BigNumber(8)).toNumber()] = new BigNumber((_2012_v7).length);
        _nw295[(new BigNumber(9)).toNumber()] = new BigNumber(484);
        _nw295[(new BigNumber(10)).toNumber()] = _2005_v1;
        _nw295[(new BigNumber(11)).toNumber()] = _2005_v1;
        _nw295[(new BigNumber(12)).toNumber()] = new BigNumber(571);
        _2017_v10 = _nw295;
        let _2018_v11;
        _2018_v11 = _dafny.Set.fromElements(_2017_v10, _2017_v10, _2017_v10);
        _2018_v11 = ((_2018_v11).Intersect(_dafny.Set.fromElements(_2017_v10, _2017_v10, _2017_v10))).Intersect(_2018_v11);
        let _2019_v12;
        _2019_v12 = _dafny.Seq.of(_2005_v1);
        let _2020_v13;
        _2020_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_2009_v5).f24, _module.__default.fm15(p0, globalState)),_2019_v12);
        let _2021_v14;
        _2021_v14 = _module.D10.create_DC34((_this).fm1(p0, _2020_v13, globalState), (_2009_v5).f24, false, !(p0), (_2009_v5).f24);
        let _2022_v15;
        _2022_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2005_v1,(_dafny.ZERO).minus(_2005_v1));
        let _2023_v16;
        _2023_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2021_v14,(((_2022_v15).contains(_2005_v1)) ? ((_2022_v15).get(_2005_v1)) : (new BigNumber(490))));
        _2023_v16 = (_2023_v16).update(_2021_v14, _2005_v1);
        let _2024_v17;
        _2024_v17 = _dafny.Seq.of(_2012_v7);
        let _2025_v18;
        _2025_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2005_v1,_2024_v17);
        let _2026_v19;
        _2026_v19 = _module.D12.create_DC42(new BigNumber((_2019_v12).length));
        let _2027_v20;
        _2027_v20 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-56)), ((_2028_v7) => function (_2029_i2) {
          return _2028_v7;
        })(_2012_v7)), _2024_v17, (((_2025_v18).contains((_2026_v19).dtor_cf69)) ? ((_2025_v18).get((_2026_v19).dtor_cf69)) : (_2024_v17)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(42)), ((_2030_v7) => function (_2031_i3) {
          return _2030_v7;
        })(_2012_v7)), _2024_v17);
        let _2032_v21;
        let _nw296 = Array((new BigNumber(8)).toNumber());
        _nw296[(_dafny.ZERO).toNumber()] = _2024_v17;
        _nw296[(_dafny.ONE).toNumber()] = _2024_v17;
        _nw296[(new BigNumber(2)).toNumber()] = (((_2025_v18).contains(_2005_v1)) ? ((_2025_v18).get(_2005_v1)) : (_2024_v17));
        _nw296[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat((_2027_v20)[_module.__default.safeIndex(new BigNumber(((_this).f23).length), new BigNumber((_2027_v20).length))], _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), function (_2033_i4) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), _2012_v7));
        _nw296[(new BigNumber(4)).toNumber()] = _2024_v17;
        _nw296[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_2024_v17, _2024_v17);
        _nw296[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_2024_v17, _2024_v17);
        _nw296[(new BigNumber(7)).toNumber()] = _2024_v17;
        _2032_v21 = _nw296;
        let _index291 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2032_v21).length));
        (_2032_v21)[_index291] = _2024_v17;
        let _index292 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2032_v21).length));
        let _rhs337 = ((p0) ? (_2019_v12) : (_2019_v12));
        let _rhs338 = _module.__default.fm53(globalState);
        let _lhs277 = _2032_v21;
        let _lhs278 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2032_v21).length));
        _2019_v12 = _rhs337;
        _lhs277[_lhs278] = _rhs338;
        (globalState).f16 = (_dafny.ZERO).minus((((_2022_v15).contains(_2005_v1)) ? ((_2022_v15).get(_2005_v1)) : ((_dafny.ZERO).minus(_module.__default.fm7(new BigNumber((_dafny.MultiSet.fromElements((_2009_v5).f24)).cardinality()), globalState)))));
      } else {
        (globalState).f10 = !(p0);
        let _2034_v23;
        _2034_v23 = _dafny.Seq.of(new BigNumber((function () {
          let _coll74 = new _dafny.Set();
          for (const _compr_74 of _dafny.IntegerRange(new BigNumber(-127), new BigNumber(193))) {
            let _2035_v22 = _compr_74;
            if (((new BigNumber(-127)).isLessThanOrEqualTo(_2035_v22)) && ((_2035_v22).isLessThan(new BigNumber(193)))) {
              _coll74.add(_module.__default.safeModuloInt(_2035_v22, (_dafny.ZERO).minus(_2005_v1)));
            }
          }
          return _coll74;
        }()).length));
        let _2036_v24;
        _2036_v24 = _dafny.Seq.of(_2034_v23);
        (globalState).f16 = (_2005_v1).minus(new BigNumber((_2036_v24).length));
        (globalState).f16 = _module.__default.safeModuloInt(_2005_v1, new BigNumber(471));
        (globalState).f11 = _2005_v1;
        let _2037_v25;
        let _init53 = function (_2038_i5) {
          return true;
        };
        let _nw297 = Array((new BigNumber(16)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw297.length); _i0_53++) {
          _nw297[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _2037_v25 = _nw297;
        let _index293 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_2037_v25).length));
        (_2037_v25)[_index293] = true;
        let _index294 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_2037_v25).length));
        (_2037_v25)[_index294] = (_2004_v0)[_module.__default.safeIndex(_2005_v1, new BigNumber((_2004_v0).length))];
      }
      let _2039_v26;
      _2039_v26 = _module.D17.create_DC55((_this).f23);
      let _rhs339 = new BigNumber((((_2039_v26).dtor_cf86).Union(((_this).f23).Intersect((_this).f23))).length);
      let _rhs340 = (_this).f23;
      let _rhs341 = _2005_v1;
      let _lhs279 = globalState;
      let _lhs280 = globalState;
      _2005_v1 = _rhs339;
      _lhs279.f13 = _rhs340;
      _lhs280.f8 = _rhs341;
      let _2040_v27;
      _2040_v27 = _dafny.Seq.UnicodeFromString("cfcba");
      let _2041_v28;
      _2041_v28 = _dafny.Seq.of(new BigNumber((_2040_v27).length));
      let _2042_v29;
      _2042_v29 = _module.D2.create_DC7(_2041_v28, p0, p0, (_dafny.ZERO).minus(_2005_v1));
      let _hi4 = (_2005_v1).plus((_2042_v29).dtor_cf13);
      for (let _2043_i6 = _2005_v1; _2043_i6.isLessThan(_hi4); _2043_i6 = _2043_i6.plus(_dafny.ONE)) {
        if (p0) {
          let _2044_v30;
          let _nw298 = Array((new BigNumber(20)).toNumber()).fill(false);
          _2044_v30 = _nw298;
          let _2045_v31;
          _2045_v31 = _dafny.MultiSet.fromElements(p0, p0);
          let _index295 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_2044_v30).length));
          (_2044_v30)[_index295] = (_2045_v31).IsProperSubsetOf(_2045_v31);
          let _2046_v33;
          _2046_v33 = new _dafny.CodePoint('c'.codePointAt(0));
          let _2047_v34;
          _2047_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2046_v33,_2043_i6);
          let _2048_v35;
          _2048_v35 = _dafny.MultiSet.fromElements(new BigNumber((_2047_v34).length), _2043_i6, (_dafny.ZERO).minus(_2043_i6), _2005_v1);
          let _index296 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_2044_v30).length));
          let _rhs342 = (false) === (p0);
          let _rhs343 = ((_this).f23).IsSubsetOf(function () {
            let _coll75 = new _dafny.Set();
            for (const _compr_75 of (_dafny.Seq.update(_2041_v28, _module.__default.safeIndex(_2005_v1, new BigNumber((_2041_v28).length)), _2005_v1)).Elements) {
              let _2049_v32 = _compr_75;
              if (_dafny.Seq.contains(_dafny.Seq.update(_2041_v28, _module.__default.safeIndex(_2005_v1, new BigNumber((_2041_v28).length)), _2005_v1), _2049_v32)) {
                _coll75.add(_module.__default.safeModuloInt(_2049_v32, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("kkgbxwxpm")).length))).length))));
              }
            }
            return _coll75;
          }());
          let _rhs344 = (((_2048_v35).IsSubsetOf(_dafny.MultiSet.fromElements(_2043_i6))) ? (new BigNumber(601)) : (_2005_v1));
          let _rhs345 = _dafny.Seq.Concat(_2041_v28, _dafny.Seq.Concat(_2041_v28, _2041_v28));
          let _lhs281 = globalState;
          let _lhs282 = _2044_v30;
          let _lhs283 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_2044_v30).length));
          let _lhs284 = globalState;
          _lhs281.f0 = _rhs342;
          _lhs282[_lhs283] = _rhs343;
          _lhs284.f6 = _rhs344;
          _2041_v28 = _rhs345;
          let _2050_v36;
          _2050_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_2004_v0).length));
          let _2051_v37;
          let _nw299 = new _module.C1();
          _nw299.__ctor((new BigNumber((_2050_v36).length)).multipliedBy(_2005_v1), (_this).f22);
          _2051_v37 = _nw299;
          _2051_v37 = _2051_v37;
          let _2052_v38;
          _2052_v38 = _dafny.Map.Empty.slice().updateUnsafe((_2044_v30)[_module.__default.safeIndex(new BigNumber(976), new BigNumber((_2044_v30).length))],(_2004_v0)[_module.__default.safeIndex((_2051_v37).f30, new BigNumber((_2004_v0).length))]);
          _2052_v38 = _2052_v38;
          let _2053_v39;
          _2053_v39 = _dafny.MultiSet.fromElements(_2044_v30);
          let _2054_v40;
          let _init54 = ((_2055_v33) => function (_2056_i7) {
            return (_dafny.MultiSet.fromElements(_2055_v33)).Intersect(_dafny.MultiSet.fromElements(_2055_v33));
          })(_2046_v33);
          let _nw300 = Array((new BigNumber(24)).toNumber());
          for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw300.length); _i0_54++) {
            _nw300[_i0_54] = _init54(new BigNumber(_i0_54));
          }
          _2054_v40 = _nw300;
          let _index297 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_2054_v40).length));
          (_2054_v40)[_index297] = _dafny.MultiSet.fromElements(_2046_v33, _2046_v33);
          let _2057_v41;
          let _nw301 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _2057_v41 = _nw301;
          let _index298 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_2057_v41).length));
          (_2057_v41)[_index298] = new BigNumber((_dafny.Seq.UnicodeFromString("r")).length);
          let _2058_v42;
          let _nw302 = new _module.C2();
          _nw302.__ctor(new BigNumber((_2048_v35).cardinality()), (_this).f22);
          _2058_v42 = _nw302;
          let _2059_v43;
          _2059_v43 = _dafny.Seq.of(_2041_v28, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-755)), ((_2060_v37) => function (_2061_i8) {
            return (_2060_v37).f30;
          })(_2051_v37)));
          let _2062_v44;
          _2062_v44 = _dafny.Set.fromElements((_2044_v30)[_module.__default.safeIndex(new BigNumber(976), new BigNumber((_2044_v30).length))]);
          let _2063_v45;
          _2063_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2058_v42,_module.__default.safeModuloInt(_2005_v1, new BigNumber((_dafny.Seq.update((_2059_v43)[_module.__default.safeIndex(new BigNumber((_2062_v44).length), new BigNumber((_2059_v43).length))], _module.__default.safeIndex(_2043_i6, new BigNumber(((_2059_v43)[_module.__default.safeIndex(new BigNumber((_2062_v44).length), new BigNumber((_2059_v43).length))]).length)), (_2051_v37).f30)).length)));
          let _2064_v46;
          _2064_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm23(new BigNumber((_2041_v28).length), _2005_v1, p0, (_2044_v30)[_module.__default.safeIndex(new BigNumber(976), new BigNumber((_2044_v30).length))], globalState)).length),(_2044_v30)[_module.__default.safeIndex(new BigNumber(976), new BigNumber((_2044_v30).length))]);
          let _2065_v47;
          _2065_v47 = _dafny.MultiSet.fromElements(_2046_v33, new _dafny.CodePoint('d'.codePointAt(0)), _2046_v33);
          let _2066_v48;
          _2066_v48 = _module.D14.create_DC48(_2043_i6, _dafny.Map.Empty.slice().updateUnsafe(_2044_v30,new BigNumber(-950)), p0);
          let _index299 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_2054_v40).length));
          let _index300 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_2057_v41).length));
          let _rhs346 = (_2053_v39).update(_2044_v30, _module.__default.abs((_dafny.ZERO).minus((_2051_v37).fm34((((_2064_v46).contains(_2005_v1)) ? ((_2064_v46).get(_2005_v1)) : (p0)), globalState))));
          let _rhs347 = _2065_v47;
          let _rhs348 = (_dafny.ZERO).minus(_2005_v1);
          let _rhs349 = ((((_2044_v30)[_module.__default.safeIndex(new BigNumber(976), new BigNumber((_2044_v30).length))]) ? (_2063_v45) : (_2063_v45))).Merge(_2063_v45);
          let _rhs350 = (((_2045_v31).contains((_2045_v31).IsProperSubsetOf(_2045_v31))) ? ((_2045_v31).get((_2045_v31).IsProperSubsetOf(_2045_v31))) : (_module.__default.safeDivisionInt((_2066_v48).dtor_cf78, (_dafny.ZERO).minus(_2005_v1))));
          let _lhs285 = _2054_v40;
          let _lhs286 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_2054_v40).length));
          let _lhs287 = _2057_v41;
          let _lhs288 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_2057_v41).length));
          let _lhs289 = globalState;
          _2053_v39 = _rhs346;
          _lhs285[_lhs286] = _rhs347;
          _lhs287[_lhs288] = _rhs348;
          _2063_v45 = _rhs349;
          _lhs289.f16 = _rhs350;
          let _2067_v49;
          let _nw303 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2067_v49 = _nw303;
          let _index301 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_2067_v49).length));
          (_2067_v49)[_index301] = _2040_v27;
          let _index302 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_2067_v49).length));
          (_2067_v49)[_index302] = _2040_v27;
        } else {
          (globalState).f16 = _module.__default.safeModuloInt(_2043_i6, new BigNumber(737));
          (globalState).f8 = (_2043_i6).minus(_2005_v1);
          let _2068_v50;
          let _nw304 = new _module.C0();
          _nw304.__ctor();
          _2068_v50 = _nw304;
          (globalState).f5 = ((p0) ? (_2004_v0) : (_2004_v0));
          let _2069_v51;
          let _nw305 = new _module.C4();
          _nw305.__ctor(p0, (_this).f22);
          _2069_v51 = _nw305;
        }
        let _2070_v52;
        _2070_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0),_dafny.Map.Empty.slice().updateUnsafe(_2043_i6,p0));
        let _2071_v53;
        _2071_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2005_v1,false);
        let _2072_v54;
        _2072_v54 = _dafny.Seq.of((((_2070_v52).contains(_2004_v0)) ? ((_2070_v52).get(_2004_v0)) : (_2071_v53)));
        _2072_v54 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(25)), ((_2073_v53) => function (_2074_i9) {
          return _2073_v53;
        })(_2071_v53)), _2072_v54);
        if (p0) {
          (globalState).f10 = p0;
          let _2075_v55;
          _2075_v55 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_2040_v27).length)), _2043_i6));
          _2075_v55 = (_2075_v55).update(p0, new BigNumber((_2040_v27).length));
          (globalState).f11 = (_2041_v28)[_module.__default.safeIndex(_2005_v1, new BigNumber((_2041_v28).length))];
          let _2076_v56;
          let _nw306 = Array((new BigNumber(3)).toNumber());
          _nw306[(_dafny.ZERO).toNumber()] = p0;
          _nw306[(_dafny.ONE).toNumber()] = (p0) && (p0);
          _nw306[(new BigNumber(2)).toNumber()] = p0;
          _2076_v56 = _nw306;
          let _index303 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_2076_v56).length));
          (_2076_v56)[_index303] = true;
          let _2077_v57;
          let _nw307 = new _module.C3();
          _nw307.__ctor((_this).f22);
          _2077_v57 = _nw307;
          let _2078_v58;
          _2078_v58 = _dafny.Map.Empty.slice().updateUnsafe((_2005_v1).isEqualTo(new BigNumber(419)),p0);
          let _index304 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_2076_v56).length));
          let _rhs351 = (_module.__default.fm7(_2005_v1, globalState)).isLessThan(new BigNumber((_dafny.MultiSet.fromElements(_2076_v56)).cardinality()));
          let _rhs352 = (((_2078_v58).contains(p0)) ? ((_2078_v58).get(p0)) : (p0));
          let _rhs353 = _2077_v57;
          let _rhs354 = !(p0) || (p0);
          let _rhs355 = _dafny.Seq.Concat(_2004_v0, _2004_v0);
          let _lhs290 = globalState;
          let _lhs291 = _2076_v56;
          let _lhs292 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_2076_v56).length));
          let _lhs293 = globalState;
          _lhs290.f0 = _rhs351;
          _lhs291[_lhs292] = _rhs352;
          _2077_v57 = _rhs353;
          r0 = _rhs354;
          _lhs293.f5 = _rhs355;
          let _2079_v59;
          let _nw308 = new _module.C4();
          _nw308.__ctor(p0, (_this).f22);
          _2079_v59 = _nw308;
        } else {
          let _2080_v60;
          _2080_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus((_dafny.ZERO).minus(_2043_i6)));
          let _2081_v61;
          let _nw309 = Array((new BigNumber(6)).toNumber()).fill(false);
          _2081_v61 = _nw309;
          let _2082_v62;
          let _nw310 = new _module.C4();
          _nw310.__ctor(!_dafny.areEqual(_2041_v28, _dafny.Seq.of(new BigNumber((_2080_v60).length), _2043_i6)), _dafny.Seq.Concat(_dafny.Seq.of(_2081_v61), _dafny.Seq.update((_this).f22, _module.__default.safeIndex(_2005_v1, new BigNumber(((_this).f22).length)), _2081_v61)));
          _2082_v62 = _nw310;
          _2082_v62 = _2082_v62;
          let _2083_v63;
          _2083_v63 = _module.D10.create_DC35(_2043_i6, _2043_i6);
          (globalState).f8 = (function (_pat_let44_0) {
            return function (_2084_dt__update__tmp_h0) {
              return function (_pat_let45_0) {
                return function (_2085_dt__update_hcf61_h0) {
                  return _module.D10.create_DC35((_2084_dt__update__tmp_h0).dtor_cf60, _2085_dt__update_hcf61_h0);
                }(_pat_let45_0);
              }(_2043_i6);
            }(_pat_let44_0);
          }(_2083_v63)).dtor_cf60;
          (globalState).f10 = p0;
          let _2086_v64;
          let _init55 = ((_2087_p0) => function (_2088_i10) {
            return _dafny.Map.Empty.slice().updateUnsafe(_2087_p0,false);
          })(p0);
          let _nw311 = Array((new BigNumber(3)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw311.length); _i0_55++) {
            _nw311[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _2086_v64 = _nw311;
          let _2089_v65;
          let _nw312 = Array((new BigNumber(3)).toNumber());
          _nw312[(_dafny.ZERO).toNumber()] = _2086_v64;
          _nw312[(_dafny.ONE).toNumber()] = _2086_v64;
          _nw312[(new BigNumber(2)).toNumber()] = _2086_v64;
          _2089_v65 = _nw312;
          let _index305 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2089_v65).length));
          (_2089_v65)[_index305] = ((_2082_v62.f28) ? (_2086_v64) : (_2086_v64));
          let _2090_v66;
          _2090_v66 = _dafny.MultiSet.fromElements(_2082_v62.f28, !(_2082_v62.f28), p0);
          let _index306 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2089_v65).length));
          let _rhs356 = (new BigNumber(802)).multipliedBy(_module.__default.safeModuloInt(_2005_v1, new BigNumber((_2090_v66).cardinality())));
          let _rhs357 = new BigNumber((_2090_v66).cardinality());
          let _rhs358 = ((_2082_v62.f28) ? (_2086_v64) : (_2086_v64));
          let _rhs359 = p0;
          let _rhs360 = p0;
          let _lhs294 = globalState;
          let _lhs295 = globalState;
          let _lhs296 = _2089_v65;
          let _lhs297 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_2089_v65).length));
          let _lhs298 = globalState;
          let _lhs299 = globalState;
          _lhs294.f16 = _rhs356;
          _lhs295.f16 = _rhs357;
          _lhs296[_lhs297] = _rhs358;
          _lhs298.f10 = _rhs359;
          _lhs299.f0 = _rhs360;
          _2041_v28 = _dafny.Seq.Concat(_module.__default.fm14(_2043_i6, _2005_v1, globalState), _2041_v28);
        }
        (globalState).f0 = p0;
      }
      let _2091_v67;
      _2091_v67 = _dafny.MultiSet.fromElements(p0, p0, p0);
      let _hi5 = ((_dafny.ZERO).minus(new BigNumber(-607))).plus((((_2091_v67).contains(p0)) ? ((_2091_v67).get(p0)) : (_2005_v1)));
      for (let _2092_i11 = new BigNumber(621); _2092_i11.isLessThan(_hi5); _2092_i11 = _2092_i11.plus(_dafny.ONE)) {
        let _2093_v68;
        _2093_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_2040_v27).length));
        let _2094_v69;
        _2094_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2092_i11,_2005_v1);
        (globalState).f8 = (((_2093_v68).contains(p0)) ? ((_2093_v68).get(p0)) : ((((_2094_v69).contains(_2092_i11)) ? ((_2094_v69).get(_2092_i11)) : (new BigNumber((_2041_v28).length)))));
        let _2095_v70;
        let _init56 = ((_2096_v1, _2097_p0, _2098_i11) => function (_2099_i12) {
          return _dafny.Seq.of(_2096_v1, _2096_v1, new BigNumber((_dafny.Seq.of(_2097_p0, _2097_p0)).length), new BigNumber((_dafny.Set.fromElements(_2098_i11, _2098_i11, _2096_v1)).length));
        })(_2005_v1, p0, _2092_i11);
        let _nw313 = Array((new BigNumber(10)).toNumber());
        for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw313.length); _i0_56++) {
          _nw313[_i0_56] = _init56(new BigNumber(_i0_56));
        }
        _2095_v70 = _nw313;
        let _index307 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_2095_v70).length));
        (_2095_v70)[_index307] = _2041_v28;
        let _index308 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_2095_v70).length));
        (_2095_v70)[_index308] = _dafny.Seq.update(_dafny.Seq.update(_2041_v28, _module.__default.safeIndex(_module.__default.fm7(_2005_v1, globalState), new BigNumber((_2041_v28).length)), _2092_i11), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(781)), ((_2100_v69, _2101_v1) => function (_2102_i13) {
          return (((_2100_v69).contains(new BigNumber((_2100_v69).length))) ? ((_2100_v69).get(new BigNumber((_2100_v69).length))) : (_2101_v1));
        })(_2094_v69, _2005_v1))).length), new BigNumber((_dafny.Seq.update(_2041_v28, _module.__default.safeIndex(_module.__default.fm7(_2005_v1, globalState), new BigNumber((_2041_v28).length)), _2092_i11)).length)), (_dafny.ZERO).minus(_2005_v1));
        if (p0) {
          r3 = _2005_v1;
          let _index309 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_2095_v70).length));
          (_2095_v70)[_index309] = _2041_v28;
          (globalState).f11 = _2092_i11;
          r0 = true;
          (globalState).f10 = p0;
        } else {
          let _2103_v71;
          _2103_v71 = _dafny.Map.Empty.slice().updateUnsafe(_2092_i11,p0);
          let _2104_v72;
          let _nw314 = Array((new BigNumber(26)).toNumber());
          _nw314[(_dafny.ZERO).toNumber()] = p0;
          _nw314[(_dafny.ONE).toNumber()] = p0;
          _nw314[(new BigNumber(2)).toNumber()] = p0;
          _nw314[(new BigNumber(3)).toNumber()] = p0;
          _nw314[(new BigNumber(4)).toNumber()] = p0;
          _nw314[(new BigNumber(5)).toNumber()] = _module.__default.fm37(_2004_v0, _2103_v71, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(928)), function (_2105_i14) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }), _module.__default.safeIndex(_2092_i11, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(928)), function (_2106_i14) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          })).length)), new _dafny.CodePoint('n'.codePointAt(0))), globalState);
          _nw314[(new BigNumber(6)).toNumber()] = (!(p0)) === (false);
          _nw314[(new BigNumber(7)).toNumber()] = p0;
          _nw314[(new BigNumber(8)).toNumber()] = (_2005_v1).isLessThanOrEqualTo(_2005_v1);
          _nw314[(new BigNumber(9)).toNumber()] = (_2005_v1).isEqualTo(_2005_v1);
          _nw314[(new BigNumber(10)).toNumber()] = p0;
          _nw314[(new BigNumber(11)).toNumber()] = p0;
          _nw314[(new BigNumber(12)).toNumber()] = p0;
          _nw314[(new BigNumber(13)).toNumber()] = (((_2004_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("awkmh")).length), new BigNumber((_2004_v0).length))]) ? (p0) : (true));
          _nw314[(new BigNumber(14)).toNumber()] = (p0) || (p0);
          _nw314[(new BigNumber(15)).toNumber()] = (_2005_v1).isEqualTo(new BigNumber((_2004_v0).length));
          _nw314[(new BigNumber(16)).toNumber()] = ((p0) ? (p0) : (p0));
          _nw314[(new BigNumber(17)).toNumber()] = (p0) && (p0);
          _nw314[(new BigNumber(18)).toNumber()] = p0;
          _nw314[(new BigNumber(19)).toNumber()] = !(!(p0));
          _nw314[(new BigNumber(20)).toNumber()] = p0;
          _nw314[(new BigNumber(21)).toNumber()] = (_2092_i11).isLessThanOrEqualTo(_2005_v1);
          _nw314[(new BigNumber(22)).toNumber()] = (_2092_i11).isLessThan(_2005_v1);
          _nw314[(new BigNumber(23)).toNumber()] = p0;
          _nw314[(new BigNumber(24)).toNumber()] = p0;
          _nw314[(new BigNumber(25)).toNumber()] = p0;
          _2104_v72 = _nw314;
          let _index310 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_2104_v72).length));
          (_2104_v72)[_index310] = !(p0) || (p0);
          let _2107_v73;
          _2107_v73 = _dafny.Map.Empty.slice().updateUnsafe(_2091_v67,_2004_v0);
          let _2108_v75;
          _2108_v75 = _dafny.Set.fromElements(_2091_v67);
          let _index311 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_2104_v72).length));
          (_2104_v72)[_index311] = (function () {
            let _coll76 = new _dafny.Set();
            for (const _compr_76 of (_2107_v73).Keys.Elements) {
              let _2109_v74 = _compr_76;
              if ((_2107_v73).contains(_2109_v74)) {
                _coll76.add(_2109_v74);
              }
            }
            return _coll76;
          }()).IsDisjointFrom((_2108_v75).Union(_2108_v75));
          (globalState).f8 = _2092_i11;
          let _2110_v76;
          let _nw315 = new _module.C4();
          _nw315.__ctor(p0, (_this).f22);
          _2110_v76 = _nw315;
          let _2111_v77;
          _2111_v77 = new _dafny.CodePoint('u'.codePointAt(0));
          _2111_v77 = _2111_v77;
          (globalState).f16 = _2092_i11;
        }
        let _2112_v78;
        let _nw316 = Array((new BigNumber(13)).toNumber()).fill(false);
        _2112_v78 = _nw316;
        let _index312 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2112_v78).length));
        (_2112_v78)[_index312] = p0;
        let _index313 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2112_v78).length));
        (_2112_v78)[_index313] = p0;
      }
      let _2113_v79;
      _2113_v79 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2040_v27);
      let _hi6 = (new BigNumber((_2004_v0).length)).multipliedBy(new BigNumber((_2113_v79).length));
      for (let _2114_i15 = _2005_v1; _2114_i15.isLessThan(_hi6); _2114_i15 = _2114_i15.plus(_dafny.ONE)) {
        (globalState).f10 = p0;
        let _2115_v80;
        let _nw317 = Array((new BigNumber(8)).toNumber()).fill(false);
        _2115_v80 = _nw317;
        let _index314 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_2115_v80).length));
        (_2115_v80)[_index314] = (!(p0)) || (false);
        let _2116_v81;
        _2116_v81 = _dafny.MultiSet.fromElements((_2114_i15).minus(_2114_i15));
        let _2117_v82;
        _2117_v82 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2005_v1);
        let _index315 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_2115_v80).length));
        let _rhs361 = !(true);
        let _rhs362 = true;
        let _rhs363 = (_dafny.ZERO).minus(new BigNumber((_2116_v81).cardinality()));
        let _rhs364 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_2004_v0, _2004_v0, _2004_v0, _2004_v0, _dafny.Seq.update(_2004_v0, _module.__default.safeIndex(new BigNumber((_2117_v82).length), new BigNumber((_2004_v0).length)), p0)), _module.__default.safeIndex(_2005_v1, new BigNumber((_dafny.Seq.of(_2004_v0, _2004_v0, _2004_v0, _2004_v0, _dafny.Seq.update(_2004_v0, _module.__default.safeIndex(new BigNumber((_2117_v82).length), new BigNumber((_2004_v0).length)), p0))).length)), _2004_v0)).length));
        let _lhs300 = globalState;
        let _lhs301 = _2115_v80;
        let _lhs302 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_2115_v80).length));
        let _lhs303 = globalState;
        let _lhs304 = globalState;
        _lhs300.f0 = _rhs361;
        _lhs301[_lhs302] = _rhs362;
        _lhs303.f11 = _rhs363;
        _lhs304.f6 = _rhs364;
        if (false) {
          let _2118_v83;
          _2118_v83 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          _2118_v83 = (_2118_v83).update((_2115_v80)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_2115_v80).length))], (_2115_v80)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_2115_v80).length))]);
          r1 = _2115_v80;
          let _2119_v84;
          _2119_v84 = _module.D9.create_DC31(_2114_i15, _2115_v80);
          (globalState).f6 = (_2119_v84).dtor_cf51;
          let _index316 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_2115_v80).length));
          (_2115_v80)[_index316] = p0;
          let _2120_v85;
          _2120_v85 = _dafny.Map.Empty.slice().updateUnsafe(_2114_i15,false);
          let _2121_v86;
          let _nw318 = new _module.C4();
          _nw318.__ctor(_module.__default.fm37(_2004_v0, _2120_v85, _2040_v27, globalState), (_this).f22);
          _2121_v86 = _nw318;
        } else {
          let _2122_v87;
          let _nw319 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _2122_v87 = _nw319;
          let _index317 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2122_v87).length));
          (_2122_v87)[_index317] = _module.__default.safeModuloInt(_2005_v1, new BigNumber((_2117_v82).length));
          let _index318 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2122_v87).length));
          (_2122_v87)[_index318] = _2005_v1;
          _2122_v87 = _2122_v87;
          let _2123_v88;
          let _nw320 = new _module.C1();
          _nw320.__ctor((_2122_v87)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2122_v87).length))], (_this).f22);
          _2123_v88 = _nw320;
          let _2124_v89;
          let _nw321 = Array((new BigNumber(27)).toNumber());
          _nw321[(_dafny.ZERO).toNumber()] = _2123_v88;
          _nw321[(_dafny.ONE).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(2)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(3)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(4)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(5)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(6)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(7)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(8)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(9)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(10)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(11)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(12)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(13)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(14)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(15)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(16)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(17)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(18)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(19)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(20)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(21)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(22)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(23)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(24)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(25)).toNumber()] = _2123_v88;
          _nw321[(new BigNumber(26)).toNumber()] = _2123_v88;
          _2124_v89 = _nw321;
          let _2125_v90;
          _2125_v90 = _module.D11.create_DC37(_2124_v89);
          let _2126_v91;
          _2126_v91 = _module.D11.create_DC40(_2125_v90);
          let _2127_v92;
          _2127_v92 = _dafny.Seq.of(_2125_v90);
          let _2128_v93;
          _2128_v93 = _dafny.Map.Empty.slice().updateUnsafe((_2122_v87)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_2122_v87).length))],(_2115_v80)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_2115_v80).length))]);
          let _2129_v94;
          let _nw322 = new _module.C5();
          _nw322.__ctor(_2128_v93, _2005_v1);
          _2129_v94 = _nw322;
          let _pat_let_tv69 = _2125_v90;
          let _pat_let_tv70 = _2125_v90;
          let _pat_let_tv71 = _2125_v90;
          let _2130_v95;
          let _nw323 = Array((new BigNumber(19)).toNumber());
          _nw323[(_dafny.ZERO).toNumber()] = function (_pat_let46_0) {
            return function (_2131_dt__update__tmp_h1) {
              return function (_pat_let47_0) {
                return function (_2132_dt__update_hcf67_h0) {
                  return _module.D11.create_DC40(_2132_dt__update_hcf67_h0);
                }(_pat_let47_0);
              }(_pat_let_tv69);
            }(_pat_let46_0);
          }(_2126_v91);
          _nw323[(_dafny.ONE).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(2)).toNumber()] = _module.D11.create_DC40(_2125_v90);
          _nw323[(new BigNumber(3)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(4)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(5)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(6)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(7)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(8)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(9)).toNumber()] = function (_pat_let48_0) {
            return function (_2133_dt__update__tmp_h2) {
              return function (_pat_let49_0) {
                return function (_2134_dt__update_hcf67_h1) {
                  return _module.D11.create_DC40(_2134_dt__update_hcf67_h1);
                }(_pat_let49_0);
              }(_pat_let_tv70);
            }(_pat_let48_0);
          }(_2126_v91);
          _nw323[(new BigNumber(10)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(11)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(12)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(13)).toNumber()] = function (_pat_let50_0) {
            return function (_2135_dt__update__tmp_h3) {
              return function (_pat_let51_0) {
                return function (_2136_dt__update_hcf67_h2) {
                  return _module.D11.create_DC40(_2136_dt__update_hcf67_h2);
                }(_pat_let51_0);
              }(_pat_let_tv71);
            }(_pat_let50_0);
          }(_2126_v91);
          _nw323[(new BigNumber(14)).toNumber()] = (((_2115_v80)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_2115_v80).length))]) ? (_2126_v91) : (_2126_v91));
          _nw323[(new BigNumber(15)).toNumber()] = _module.D11.create_DC40((_2127_v92)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_2129_v94)).length), new BigNumber((_2127_v92).length))]);
          _nw323[(new BigNumber(16)).toNumber()] = _2126_v91;
          _nw323[(new BigNumber(17)).toNumber()] = _module.D11.create_DC40(_2125_v90);
          _nw323[(new BigNumber(18)).toNumber()] = _2126_v91;
          _2130_v95 = _nw323;
          let _index319 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_2130_v95).length));
          (_2130_v95)[_index319] = _2126_v91;
          let _2137_v96;
          _2137_v96 = _dafny.Set.fromElements(p0);
          let _index320 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_2130_v95).length));
          let _rhs365 = !((_2115_v80)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_2115_v80).length))]) || ((_2137_v96).IsProperSubsetOf(_2137_v96));
          let _rhs366 = _2126_v91;
          let _rhs367 = _2115_v80;
          let _lhs305 = globalState;
          let _lhs306 = _2130_v95;
          let _lhs307 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_2130_v95).length));
          _lhs305.f10 = _rhs365;
          _lhs306[_lhs307] = _rhs366;
          r1 = _rhs367;
          let _index321 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_2122_v87).length));
          (_2122_v87)[_index321] = new BigNumber(625);
          (globalState).f16 = (_2129_v94).f27;
        }
        let _2138_v97;
        _2138_v97 = new _dafny.CodePoint('j'.codePointAt(0));
        _2138_v97 = _2138_v97;
      }
      let _2139_v98;
      let _init57 = ((_2140_v28) => function (_2141_i16) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(93)), ((_2142_v28) => function (_2143_i17) {
          return _2142_v28;
        })(_2140_v28));
      })(_2041_v28);
      let _nw324 = Array((new BigNumber(7)).toNumber());
      for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw324.length); _i0_57++) {
        _nw324[_i0_57] = _init57(new BigNumber(_i0_57));
      }
      _2139_v98 = _nw324;
      let _2144_v99;
      _2144_v99 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2005_v1);
      let _2145_v100;
      _2145_v100 = _dafny.Seq.of(_dafny.Seq.of(_2005_v1, new BigNumber((_2144_v99).length)));
      let _index322 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_2139_v98).length));
      (_2139_v98)[_index322] = _dafny.Seq.Concat(_2145_v100, _2145_v100);
      let _2146_v101;
      _2146_v101 = _module.D15.create_DC51(p0, false);
      let _2147_v102;
      _2147_v102 = new _dafny.CodePoint('q'.codePointAt(0));
      let _2148_v104;
      _2148_v104 = _dafny.Map.Empty.slice().updateUnsafe(p0,function () {
        let _coll77 = new _dafny.Map();
        for (const _compr_77 of (_dafny.Seq.of(new BigNumber(-952))).Elements) {
          let _2149_v103 = _compr_77;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(-952)), _2149_v103)) {
            _coll77.push([(_2149_v103).multipliedBy(_2005_v1),_2005_v1]);
          }
        }
        return _coll77;
      }());
      let _2150_v105;
      _2150_v105 = _module.D0.create_DC1(_2005_v1, p0, _2148_v104, new BigNumber((_2091_v67).cardinality()), _2147_v102);
      let _2151_v106;
      _2151_v106 = _dafny.Seq.of(_2150_v105, _2150_v105);
      let _index323 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_2139_v98).length));
      (_2139_v98)[_index323] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm54(_2146_v101, _2005_v1, _2147_v102, new BigNumber((_2151_v106).length), globalState), _dafny.Seq.of(_2041_v28)), _module.__default.safeIndex(new BigNumber(64), new BigNumber((_dafny.Seq.Concat(_module.__default.fm54(_2146_v101, _2005_v1, _2147_v102, new BigNumber((_2151_v106).length), globalState), _dafny.Seq.of(_2041_v28))).length)), _2041_v28), _2145_v100);
      r0 = p0;
      let _init58 = ((_2152_p0) => function (_2153_i18) {
        return _2152_p0;
      })(p0);
      let _nw325 = Array((new BigNumber(4)).toNumber());
      for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw325.length); _i0_58++) {
        _nw325[_i0_58] = _init58(new BigNumber(_i0_58));
      }
      r1 = _nw325;
      let _2154_v107;
      _2154_v107 = _dafny.Map.Empty.slice().updateUnsafe(_2041_v28,_2147_v102);
      r2 = (_2154_v107).Merge(_2154_v107);
      r3 = _2005_v1;
      return [r0, r1, r2, r3];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
