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
    static fm0(globalState) {
      return false;
    };
    static fm1(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bur"), _dafny.Seq.UnicodeFromString("noxdr"));
    };
    static fm2(p0, p1, globalState) {
      return _module.__default.safeModuloInt((new BigNumber((_dafny.Seq.UnicodeFromString("iu")).length)).multipliedBy(new BigNumber(694)), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)))).length));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(90))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-787))));
    };
    static fm11(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(true))));
    };
    static fm12(p0, p1, globalState) {
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.of(true)) : (_dafny.Seq.of(false))), _dafny.Seq.of(false, true));
    };
    static fm15(p0, globalState) {
      return ((true) ? (_dafny.Seq.UnicodeFromString("kgdqdqqcq")) : (_dafny.Seq.UnicodeFromString("qcmmbgxd")));
    };
    static fm16(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rchljvs"),true)).Merge(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("qxsrnkgpo"))).Elements) {
          let _0_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("qxsrnkgpo")), _0_v0)) {
            _coll0.push([_0_v0,false]);
          }
        }
        return _coll0;
      }())).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(283)), function (_1_i0) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("ukufam"))).Elements) {
          let _2_v1 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(283)), function (_1_i0) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("ukufam")), _2_v1)) {
            _coll1.push([_2_v1,!(!(false))]);
          }
        }
        return _coll1;
      }());
    };
    static fm17(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D2.create_DC3(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("djvdjhw"), _dafny.Seq.UnicodeFromString("vqfwcrmm")));
      if (_source0.is_DC4) {
        let _3___mcc_h0 = (_source0).cf6;
        let _4___mcc_h1 = (_source0).cf7;
        let _5___mcc_h2 = (_source0).cf8;
        let _6_cf8 = _5___mcc_h2;
        let _7_cf7 = _4___mcc_h1;
        let _8_cf6 = _3___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_6_cf8)).length)), _dafny.Seq.of(_6_cf8, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(true)).length))).length),new BigNumber((_dafny.Seq.UnicodeFromString("cfcbcpfd")).length)),false)).length))));
      } else if (_source0.is_DC5) {
        return _dafny.Seq.of(new BigNumber((function () {
          let _coll2 = new _dafny.Map();
          for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(119),(_module.D8.create_DC16(true, (_module.D16.create_DC39(new BigNumber((_dafny.Seq.of(true, false)).length), true, _dafny.Map.Empty.slice().updateUnsafe(true,true), new _dafny.CodePoint('s'.codePointAt(0)))).dtor_cf62)).dtor_cf24)).Keys.Elements) {
            let _9_v0 = _compr_2;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(119),(_module.D8.create_DC16(true, (_module.D16.create_DC39(new BigNumber((_dafny.Seq.of(true, false)).length), true, _dafny.Map.Empty.slice().updateUnsafe(true,true), new _dafny.CodePoint('s'.codePointAt(0)))).dtor_cf62)).dtor_cf24)).contains(_9_v0)) {
              _coll2.push([_module.__default.safeDivisionInt(_9_v0, new BigNumber((_dafny.Seq.of(new BigNumber(333))).length)),_dafny.MultiSet.fromElements(new BigNumber(807))]);
            }
          }
          return _coll2;
        }()).length));
      } else {
        let _10___mcc_h3 = (_source0).cf5;
        let _11_cf5 = _10___mcc_h3;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_12_i0) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(952)), function (_13_i1) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })).length));
        });
      }
    };
    static fm18(p0, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(950))).Intersect(_dafny.Set.fromElements(new BigNumber((function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), function (_14_i0) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        })).length))).Elements) {
          let _15_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), function (_14_i0) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          })).length)), _15_v0)) {
            _coll3.add((_15_v0).plus(new BigNumber(443)));
          }
        }
        return _coll3;
      }()).length)))).Union((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))).length))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(!(false))).length), new BigNumber(778))));
    };
    static fm19(p0, p1, globalState) {
      return _dafny.Seq.of(false);
    };
    static fm20(p0, p1, p2, globalState) {
      let _source1 = _module.D22.create_DC51(_dafny.Set.fromElements(new BigNumber(103)));
      if (_source1.is_DC52) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      } else if (_source1.is_DC53) {
        let _16___mcc_h0 = (_source1).cf87;
        let _17___mcc_h1 = (_source1).cf88;
        let _18___mcc_h2 = (_source1).cf89;
        let _19___mcc_h3 = (_source1).cf90;
        let _20_cf90 = _19___mcc_h3;
        let _21_cf89 = _18___mcc_h2;
        let _22_cf88 = _17___mcc_h1;
        let _23_cf87 = _16___mcc_h0;
        return new _dafny.CodePoint('h'.codePointAt(0));
      } else {
        let _24___mcc_h4 = (_source1).cf86;
        let _25_cf86 = _24___mcc_h4;
        return new _dafny.CodePoint('j'.codePointAt(0));
      }
    };
    static fm22(p0, p1, globalState) {
      return ((_module.D6.create_DC9(function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of _dafny.IntegerRange(new BigNumber(40), new BigNumber(-988))) {
    let _26_v0 = _compr_4;
    if (((new BigNumber(40)).isLessThanOrEqualTo(_26_v0)) && ((_26_v0).isLessThan(new BigNumber(-988)))) {
      _coll4.push([(_26_v0).multipliedBy(new BigNumber(144)),true]);
    }
  }
  return _coll4;
}())).dtor_cf12).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(58),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),false)).length),false)));
    };
    static fm23(p0, p1, globalState) {
      return _dafny.Seq.of(_dafny.Seq.of(true, false), _dafny.Seq.of(!(false), true), _dafny.Seq.of(!(!(true)), false, true));
    };
    static fm24(p0, globalState) {
      if (false) {
        return _dafny.MultiSet.fromElements(false, false);
      } else {
        return _dafny.MultiSet.fromElements(true);
      }
    };
    static fm26(globalState) {
      return _dafny.Set.fromElements(new BigNumber(373));
    };
    static fm27(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(12)), function (_27_i0) {
        return _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(_module.D13.create_DC28(_dafny.Seq.UnicodeFromString("livdrujf"), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D6.create_DC9(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll5 = new _dafny.Set();
  for (const _compr_5 of _dafny.IntegerRange(new BigNumber(951), new BigNumber(703))) {
    let _28_v0 = _compr_5;
    if (((new BigNumber(951)).isLessThanOrEqualTo(_28_v0)) && ((_28_v0).isLessThan(new BigNumber(703)))) {
      _coll5.add(_module.__default.safeModuloInt(_28_v0, new BigNumber(631)));
    }
  }
  return _coll5;
}()).length)),false)))).length), true))).length), (_dafny.ZERO).minus(new BigNumber(-719)));
      });
    };
    static fm28(p0, p1, globalState) {
      return _dafny.Set.fromElements(((false) ? (_dafny.Set.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)))) : (_dafny.Set.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0))))), (function () {
        let _coll6 = new _dafny.Set();
        for (const _compr_6 of (_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)))).Elements) {
          let _29_v0 = _compr_6;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))), _29_v0)) {
            _coll6.add(_29_v0);
          }
        }
        return _coll6;
      }()).Difference(_dafny.Set.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)))), (_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))).Difference(_dafny.Set.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))), function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),false)).Keys.Elements) {
          let _30_v1 = _compr_7;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),false)).contains(_30_v1)) {
            _coll7.add(_30_v1);
          }
        }
        return _coll7;
      }());
    };
    static fm29(globalState) {
      return ((_dafny.Set.fromElements(_module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
  let _coll8 = new _dafny.Map();
  for (const _compr_8 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), function (_31_i0) {
    return new BigNumber(-482);
  })).Elements) {
    let _32_v0 = _compr_8;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), function (_31_i0) {
      return new BigNumber(-482);
    }), _32_v0)) {
      _coll8.push([_module.__default.safeDivisionInt(_32_v0, new BigNumber((_dafny.Seq.UnicodeFromString("wuy")).length)),true]);
    }
  }
  return _coll8;
}()).length),false)).length),true)).length))))).Union(_dafny.Set.fromElements(_module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-716)))))).Difference((_dafny.Set.fromElements(_module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.UnicodeFromString("hucnnesk")).length))), _module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(759))), _module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(true)).length))), _module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber((function () {
  let _coll9 = new _dafny.Set();
  for (const _compr_9 of (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(473), new BigNumber(328)), _dafny.Set.fromElements(new BigNumber(464)), _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("v")).length), new BigNumber((_dafny.Seq.of(true, false, true, false, true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("tloyglr")).length))).length)))).Elements) {
    let _33_v1 = _compr_9;
    if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(473), new BigNumber(328)), _dafny.Set.fromElements(new BigNumber(464)), _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("v")).length), new BigNumber((_dafny.Seq.of(true, false, true, false, true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("tloyglr")).length))).length)))).contains(_33_v1)) {
      _coll9.add(_33_v1);
    }
  }
  return _coll9;
}()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber(772))).length))), _module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(832))))).Difference(_dafny.Set.fromElements(_module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("m")).length))))));
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-526),_dafny.Map.Empty.slice().updateUnsafe(!(false),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(144),_dafny.Map.Empty.slice().updateUnsafe(false,true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(354),_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm31(p0, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(368)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-154)), function (_34_i0) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-107)), function (_35_i1) {
          return new BigNumber(741);
        })).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pbacfmeic"),new BigNumber(500))).length), (_dafny.ZERO).minus(new BigNumber(-351)), new BigNumber(576), new BigNumber(622)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), function (_36_i2) {
        return new BigNumber(901);
      }));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(313), new BigNumber(538))).length))).length)))).Merge(function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, false)).length)))))).Elements) {
          let _37_v0 = _compr_10;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, false)).length)))))).contains(_37_v0)) {
            _coll10.push([(_37_v0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)),(_module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(841)))).dtor_cf26]);
          }
        }
        return _coll10;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(703),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(669))).length))));
    };
    static fm33(p0, globalState) {
      if (!(false)) {
        return _module.D6.create_DC10((_module.D11.create_DC24(new BigNumber(413), _dafny.Seq.of(true), new _dafny.CodePoint('f'.codePointAt(0)))).dtor_cf35);
      } else {
        return _module.D6.create_DC10(_dafny.Seq.of(true));
      }
    };
    static fm34(p0, p1, globalState) {
      let _source2 = _module.D22.create_DC52();
      if (_source2.is_DC52) {
        return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("re")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),true)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), function (_38_i0) {
          return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dppgedr")).length))).cardinality());
        }));
      } else if (_source2.is_DC53) {
        let _39___mcc_h0 = (_source2).cf87;
        let _40___mcc_h1 = (_source2).cf88;
        let _41___mcc_h2 = (_source2).cf89;
        let _42___mcc_h3 = (_source2).cf90;
        let _43_cf90 = _42___mcc_h3;
        let _44_cf89 = _41___mcc_h2;
        let _45_cf88 = _40___mcc_h1;
        let _46_cf87 = _39___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_46_cf87), _dafny.Seq.Create(_module.__default.abs(new BigNumber(140)), function (_47_i1) {
          return new BigNumber(921);
        }));
      } else {
        let _48___mcc_h4 = (_source2).cf86;
        let _49_cf86 = _48___mcc_h4;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-42)), function (_50_i2) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-538),true)).length);
        });
      }
    };
    static fm35(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false)), _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(true)));
    };
    static fm36(p0, p1, p2, globalState) {
      let _source3 = _module.D6.create_DC9(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(19),!(true)));
      if (_source3.is_DC10) {
        let _51___mcc_h0 = (_source3).cf13;
        let _52_cf13 = _51___mcc_h0;
        return _module.D7.create_DC13(_dafny.Set.fromElements(false), new BigNumber((_dafny.Seq.UnicodeFromString("xcxj")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(628))).length),new BigNumber(-166))).length));
      } else if (_source3.is_DC11) {
        let _53___mcc_h1 = (_source3).cf14;
        let _54___mcc_h2 = (_source3).cf15;
        let _55___mcc_h3 = (_source3).cf16;
        let _56_cf16 = _55___mcc_h3;
        let _57_cf15 = _54___mcc_h2;
        let _58_cf14 = _53___mcc_h1;
        return _module.D7.create_DC13(_dafny.Set.fromElements(true, true), _57_cf15, _56_cf16);
      } else {
        let _59___mcc_h4 = (_source3).cf12;
        let _60_cf12 = _59___mcc_h4;
        return _module.D7.create_DC13(_dafny.Set.fromElements(false), new BigNumber(461), new BigNumber(-56));
      }
    };
    static fm37(globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber((_dafny.Set.fromElements(new BigNumber(863))).length)).isEqualTo(new BigNumber((_dafny.Set.fromElements(false, true)).length)), _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("dexnes"), _dafny.Seq.UnicodeFromString("sijt")), !(!(new BigNumber(369)).isEqualTo(new BigNumber((_dafny.Set.fromElements(false)).length))), false);
    };
    static fm38(p0, p1, p2, globalState) {
      return _module.D0.create_DC0(_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(278)), function (_61_i0) {
  return new _dafny.CodePoint('r'.codePointAt(0));
})).length), new BigNumber((_dafny.Seq.UnicodeFromString("jrpag")).length))).cardinality()), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber(566), new BigNumber((_dafny.Seq.of(true, true, true, true)).length))).length)));
    };
    static fm39(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('a'.codePointAt(0));
    };
    static fm40(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(_module.D11.create_DC23(false, new BigNumber(575), new BigNumber(896)), _module.D11.create_DC23(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(169))).length), new BigNumber(427)));
    };
    static fm41(globalState) {
      let _source4 = _module.D15.create_DC36(new BigNumber(4), false, true, true, new BigNumber(388));
      if (_source4.is_DC34) {
        return _dafny.Set.fromElements(false, false, !((_module.D20.create_DC45(true, true)).dtor_cf76), !(!(false)), true);
      } else if (_source4.is_DC35) {
        let _62___mcc_h0 = (_source4).cf50;
        let _63___mcc_h1 = (_source4).cf51;
        let _64___mcc_h2 = (_source4).cf52;
        let _65___mcc_h3 = (_source4).cf53;
        let _66___mcc_h4 = (_source4).cf54;
        let _67_cf54 = _66___mcc_h4;
        let _68_cf53 = _65___mcc_h3;
        let _69_cf52 = _64___mcc_h2;
        let _70_cf51 = _63___mcc_h1;
        let _71_cf50 = _62___mcc_h0;
        return ((_70_cf51).f12).Union((_70_cf51).f12);
      } else if (_source4.is_DC36) {
        let _72___mcc_h5 = (_source4).cf55;
        let _73___mcc_h6 = (_source4).cf56;
        let _74___mcc_h7 = (_source4).cf57;
        let _75___mcc_h8 = (_source4).cf58;
        let _76___mcc_h9 = (_source4).cf59;
        let _77_cf59 = _76___mcc_h9;
        let _78_cf58 = _75___mcc_h8;
        let _79_cf57 = _74___mcc_h7;
        let _80_cf56 = _73___mcc_h6;
        let _81_cf55 = _72___mcc_h5;
        return (_dafny.Set.fromElements(_78_cf58, _79_cf57)).Difference(_dafny.Set.fromElements(_78_cf58));
      } else if (_source4.is_DC33) {
        let _82___mcc_h10 = (_source4).cf49;
        let _83_cf49 = _82___mcc_h10;
        return _dafny.Set.fromElements(true);
      } else {
        let _84___mcc_h11 = (_source4).cf60;
        let _85_cf60 = _84___mcc_h11;
        return _dafny.Set.fromElements(!(!(true)), !(true));
      }
    };
    static fm42(p0, globalState) {
      return _module.D9.create_DC19();
    };
    static fm43(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("bha"),_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(703))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("hpxjw"),_dafny.MultiSet.fromElements(new BigNumber(-382))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(529)), function (_86_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(697)), function (_87_i1) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      })).length), new BigNumber((function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of (_dafny.Seq.of(new BigNumber(347), new BigNumber((function () {
          let _coll12 = new _dafny.Map();
          for (const _compr_12 of _dafny.IntegerRange(new BigNumber(633), new BigNumber(-258))) {
            let _88_v0 = _compr_12;
            if (((new BigNumber(633)).isLessThanOrEqualTo(_88_v0)) && ((_88_v0).isLessThan(new BigNumber(-258)))) {
              _coll12.push([_module.__default.safeModuloInt(_88_v0, new BigNumber((_dafny.Seq.of(true)).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)]);
            }
          }
          return _coll12;
        }()).length))).Elements) {
          let _89_v1 = _compr_11;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(347), new BigNumber((function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of _dafny.IntegerRange(new BigNumber(633), new BigNumber(-258))) {
              let _88_v0 = _compr_13;
              if (((new BigNumber(633)).isLessThanOrEqualTo(_88_v0)) && ((_88_v0).isLessThan(new BigNumber(-258)))) {
                _coll13.push([_module.__default.safeModuloInt(_88_v0, new BigNumber((_dafny.Seq.of(true)).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)]);
              }
            }
            return _coll13;
          }()).length)), _89_v1)) {
            _coll11.add((_89_v1).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(634), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(993)), function (_90_i2) {
              return new _dafny.CodePoint('l'.codePointAt(0));
            })).length))).length), new BigNumber((function () {
              let _coll14 = new _dafny.Map();
              for (const _compr_14 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wra"),new BigNumber(574))).Keys.Elements) {
                let _91_v2 = _compr_14;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wra"),new BigNumber(574))).contains(_91_v2)) {
                  _coll14.push([_91_v2,false]);
                }
              }
              return _coll14;
            }()).length))).cardinality())));
          }
        }
        return _coll11;
      }()).length))));
    };
    static fm44(p0, p1, globalState) {
      return _dafny.Seq.of((function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of (_dafny.MultiSet.fromElements(_module.D8.create_DC17(_module.D8.create_DC16(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('p'.codePointAt(0)))).length))))).Elements) {
          let _92_v0 = _compr_15;
          if ((_dafny.MultiSet.fromElements(_module.D8.create_DC17(_module.D8.create_DC16(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('p'.codePointAt(0)))).length))))).contains(_92_v0)) {
            _coll15.push([_92_v0,(_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(!((_module.D21.create_DC48(false)).dtor_cf78), !(false))).length))]);
          }
        }
        return _coll15;
      }()).Merge(function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(727)), function (_93_i0) {
          return _module.D8.create_DC17(_module.D8.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
        })).Elements) {
          let _94_v1 = _compr_16;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(727)), function (_93_i0) {
            return _module.D8.create_DC17(_module.D8.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
          }), _94_v1)) {
            _coll16.push([_94_v1,new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(true)).length))))).cardinality())]);
          }
        }
        return _coll16;
      }()), function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_dafny.Seq.of(_module.D8.create_DC17(_module.D8.create_DC17(_module.D8.create_DC16(false, new BigNumber((_dafny.Seq.of(new BigNumber(456))).length)))))).Elements) {
          let _95_v2 = _compr_17;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D8.create_DC17(_module.D8.create_DC17(_module.D8.create_DC16(false, new BigNumber((_dafny.Seq.of(new BigNumber(456))).length))))), _95_v2)) {
            _coll17.push([_95_v2,new BigNumber(988)]);
          }
        }
        return _coll17;
      }(), (_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC17(_module.D8.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(false,true))),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(819)), function (_96_i1) {
        return new BigNumber((_dafny.Set.fromElements(!(true), true, true, false)).length);
      })).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC17(_module.D8.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(true,!(true)))),new BigNumber(-665))));
    };
    static fm45(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber(-462))).Union((_dafny.MultiSet.fromElements(new BigNumber(-710))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("mxu")).length))));
    };
    static fm46(globalState) {
      return (((_module.D16.create_DC38(function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of (function () {
    let _coll19 = new _dafny.Map();
    for (const _compr_19 of _dafny.IntegerRange(new BigNumber(945), new BigNumber(197))) {
      let _97_v1 = _compr_19;
      if (((new BigNumber(945)).isLessThanOrEqualTo(_97_v1)) && ((_97_v1).isLessThan(new BigNumber(197)))) {
        _coll19.push([(_97_v1).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)]);
      }
    }
    return _coll19;
  }()).Keys.Elements) {
    let _98_v0 = _compr_18;
    if ((function () {
      let _coll20 = new _dafny.Map();
      for (const _compr_20 of _dafny.IntegerRange(new BigNumber(945), new BigNumber(197))) {
        let _97_v1 = _compr_20;
        if (((new BigNumber(945)).isLessThanOrEqualTo(_97_v1)) && ((_97_v1).isLessThan(new BigNumber(197)))) {
          _coll20.push([(_97_v1).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)]);
        }
      }
      return _coll20;
    }()).contains(_98_v0)) {
      _coll18.push([(_98_v0).minus(new BigNumber(-664)),new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())]);
    }
  }
  return _coll18;
}())).dtor_cf61).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(173),new BigNumber(261))).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(666))).length),new BigNumber((_dafny.Seq.UnicodeFromString("vvgkqkf")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).length),new BigNumber((_dafny.Seq.UnicodeFromString("w")).length))));
    };
    static fm47(p0, p1, p2, p3, globalState) {
      return _module.D13.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC17(_module.D8.create_DC16(false, new BigNumber((function () {
  let _coll21 = new _dafny.Map();
  for (const _compr_21 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(!(true))).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).Keys.Elements) {
    let _99_v0 = _compr_21;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(!(true))).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).contains(_99_v0)) {
      _coll21.push([_99_v0,new BigNumber((_dafny.Set.fromElements(true)).length)]);
    }
  }
  return _coll21;
}()).length))),new BigNumber(777)));
    };
    static fm48(p0, globalState) {
      if (false) {
        return _module.D8.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(true,!(true)));
      } else {
        return _module.D8.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(!(false),false));
      }
    };
    static fm49(p0, globalState) {
      if ((new BigNumber((_dafny.Set.fromElements(false, false, false, !(true), true)).length)).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), function (_100_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length))) {
        return _module.D0.create_DC1(true, new BigNumber(817), true);
      } else {
        return _module.D0.create_DC1(!(true), new BigNumber((_dafny.Seq.of(true)).length), false);
      }
    };
    static fm50(p0, p1, p2, p3, globalState) {
      return _module.D11.create_DC23(((false) ? (true) : (!(false))), ((true) ? (new BigNumber(-90)) : (new BigNumber(147))), new BigNumber(((_dafny.Set.fromElements(true)).Difference(_dafny.Set.fromElements(false))).length));
    };
    static fm51(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D14.create_DC32(_module.D14.create_DC31(true, new BigNumber(177), new BigNumber(233)))), _dafny.Seq.of(_module.D14.create_DC32(_module.D14.create_DC32(_module.D14.create_DC32(_module.D14.create_DC31(false, new BigNumber(71), new BigNumber(802))))), _module.D14.create_DC32(_module.D14.create_DC32(_module.D14.create_DC32(_module.D14.create_DC31(false, new BigNumber(-926), new BigNumber(123))))), _module.D14.create_DC32(_module.D14.create_DC32(_module.D14.create_DC31(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), (_dafny.ZERO).minus(new BigNumber((function () {
  let _coll22 = new _dafny.Set();
  for (const _compr_22 of (_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)))).Elements) {
    let _101_v0 = _compr_22;
    if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))), _101_v0)) {
      _coll22.add(_101_v0);
    }
  }
  return _coll22;
}()).length))))))), _dafny.Seq.of(_module.D14.create_DC32(_module.D14.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.UnicodeFromString("immtempc")))), _module.D14.create_DC32(_module.D14.create_DC31(true, new BigNumber(903), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-189))).length))).cardinality())))));
    };
    static fm52(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(!(false), new BigNumber(949), false),false)).Merge((function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of (function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of (_dafny.Seq.of(_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("nkuwoejgm")).length), false), _module.D0.create_DC1(false, new BigNumber(45), !(false)), _module.D0.create_DC1(!(false), new BigNumber((_dafny.Seq.of(new BigNumber(568))).length), !(!(false))))).Elements) {
            let _102_v1 = _compr_24;
            if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("nkuwoejgm")).length), false), _module.D0.create_DC1(false, new BigNumber(45), !(false)), _module.D0.create_DC1(!(false), new BigNumber((_dafny.Seq.of(new BigNumber(568))).length), !(!(false)))), _102_v1)) {
              _coll24.push([_102_v1,_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))]);
            }
          }
          return _coll24;
        }()).Keys.Elements) {
          let _103_v0 = _compr_23;
          if ((function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of (_dafny.Seq.of(_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("nkuwoejgm")).length), false), _module.D0.create_DC1(false, new BigNumber(45), !(false)), _module.D0.create_DC1(!(false), new BigNumber((_dafny.Seq.of(new BigNumber(568))).length), !(!(false))))).Elements) {
              let _102_v1 = _compr_25;
              if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("nkuwoejgm")).length), false), _module.D0.create_DC1(false, new BigNumber(45), !(false)), _module.D0.create_DC1(!(false), new BigNumber((_dafny.Seq.of(new BigNumber(568))).length), !(!(false)))), _102_v1)) {
                _coll25.push([_102_v1,_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))]);
              }
            }
            return _coll25;
          }()).contains(_103_v0)) {
            _coll23.push([_103_v0,true]);
          }
        }
        return _coll23;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(true, (_dafny.ZERO).minus(new BigNumber(-656)), true),false)));
    };
    static fm53(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Set.fromElements(false), _dafny.Set.fromElements(!(!(true)), !(!(false))));
    };
    static fm54(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),new BigNumber((function () {
        let _coll26 = new _dafny.Map();
        for (const _compr_26 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new _dafny.CodePoint('h'.codePointAt(0))),_module.D6.create_DC9(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), function (_104_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})).length),true)))).Keys.Elements) {
          let _105_v0 = _compr_26;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new _dafny.CodePoint('h'.codePointAt(0))),_module.D6.create_DC9(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), function (_104_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})).length),true)))).contains(_105_v0)) {
            _coll26.push([_105_v0,false]);
          }
        }
        return _coll26;
      }()).length));
    };
    static fm55(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(!(true), false)).length),new BigNumber(294))).length)),_dafny.Map.Empty.slice().updateUnsafe(false,false))).Merge(function () {
        let _coll27 = new _dafny.Map();
        for (const _compr_27 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(464)), function (_106_i0) {
          return _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(-889));
        })).Elements) {
          let _107_v0 = _compr_27;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(464)), function (_106_i0) {
            return _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(-889));
          }), _107_v0)) {
            _coll27.push([_107_v0,_dafny.Map.Empty.slice().updateUnsafe(true,false)]);
          }
        }
        return _coll27;
      }());
    };
    static fm56(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(501)), function (_108_i0) {
        return _dafny.Seq.of(new BigNumber(125), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)))).cardinality()), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('s'.codePointAt(0))))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,true), _dafny.Map.Empty.slice().updateUnsafe(true,(_module.D15.create_DC36(new BigNumber((_dafny.Set.fromElements(new BigNumber(878))).length), true, true, true, new BigNumber(714))).dtor_cf57))).length)), new BigNumber(750));
      }), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(-861)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("dgbtjar")).length), new BigNumber(24)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(97)), function (_109_i1) {
        return new BigNumber(63);
      })));
    };
    static Main(__noArgsParameter) {
      let _110_v0;
      _110_v0 = false;
      let _111_v1;
      _111_v1 = new BigNumber(754);
      let _112_v2;
      _112_v2 = _dafny.Seq.of(true, _110_v0, _110_v0);
      let _113_v3;
      _113_v3 = _dafny.Map.Empty.slice().updateUnsafe(_111_v1,new BigNumber((_112_v2).length));
      let _114_v4;
      _114_v4 = _dafny.Map.Empty.slice().updateUnsafe(_110_v0,(((_113_v3).contains(_111_v1)) ? ((_113_v3).get(_111_v1)) : (_111_v1)));
      let _115_v5;
      let _nw0 = Array((new BigNumber(8)).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = _114_v4;
      _nw0[(_dafny.ONE).toNumber()] = _114_v4;
      _nw0[(new BigNumber(2)).toNumber()] = _114_v4;
      _nw0[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_110_v0,_111_v1);
      _nw0[(new BigNumber(4)).toNumber()] = _114_v4;
      _nw0[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_110_v0,_111_v1);
      _nw0[(new BigNumber(6)).toNumber()] = _114_v4;
      _nw0[(new BigNumber(7)).toNumber()] = _114_v4;
      _115_v5 = _nw0;
      let _116_globalState;
      let _nw1 = new _module.GlobalState();
      _nw1.__ctor(false, new BigNumber(92), new BigNumber(431), new BigNumber(-645), false, _115_v5, true, true, false, new BigNumber(-357), true);
      _116_globalState = _nw1;
      let _117_i0;
      _117_i0 = _dafny.ZERO;
      L0: {
        while (_110_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_117_i0)) {
              break L0;
            }
            _117_i0 = (_117_i0).plus(_dafny.ONE);
            (_116_globalState).f6 = _110_v0;
            if (!(_110_v0)) {
              let _118_v6;
              _118_v6 = _module.D0.create_DC1(_110_v0, _111_v1, _110_v0);
              let _119_v7;
              _119_v7 = _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_110_v0, _110_v0)).length));
              let _120_v8;
              _120_v8 = _dafny.MultiSet.fromElements(_111_v1, _111_v1, _111_v1, (_dafny.ZERO).minus(_111_v1), _111_v1);
              let _121_v9;
              _121_v9 = _dafny.Set.fromElements((((_120_v8).contains(_111_v1)) ? ((_120_v8).get(_111_v1)) : (new BigNumber(664))));
              let _122_v10;
              _122_v10 = _dafny.Seq.of(_111_v1, (_118_v6).dtor_cf2, _111_v1, new BigNumber((_119_v7).length), new BigNumber((_121_v9).length));
              let _rhs0 = !(_dafny.areEqual(_dafny.Seq.of(_111_v1, _111_v1), _119_v7));
              let _rhs1 = _110_v0;
              let _rhs2 = !(false) || (_110_v0);
              let _lhs0 = _116_globalState;
              let _lhs1 = _116_globalState;
              _lhs0.f7 = _rhs0;
              _lhs1.f7 = _rhs1;
              _110_v0 = _rhs2;
              let _123_v11;
              let _nw2 = Array((new BigNumber(26)).toNumber()).fill(false);
              _123_v11 = _nw2;
              _123_v11 = _123_v11;
              let _124_v12;
              _124_v12 = _dafny.Seq.of(_123_v11);
              let _125_v13;
              _125_v13 = _dafny.Set.fromElements(_module.__default.fm0(_116_globalState));
              let _126_v14;
              _126_v14 = _dafny.Map.Empty.slice().updateUnsafe(_111_v1,_125_v13);
              let _127_v15;
              let _nw3 = Array((new BigNumber(17)).toNumber());
              _nw3[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_119_v7).length)).multipliedBy(new BigNumber((_113_v3).length)));
              _nw3[(_dafny.ONE).toNumber()] = _111_v1;
              _nw3[(new BigNumber(2)).toNumber()] = _111_v1;
              _nw3[(new BigNumber(3)).toNumber()] = _111_v1;
              _nw3[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(_111_v1, _111_v1);
              _nw3[(new BigNumber(5)).toNumber()] = new BigNumber((_124_v12).length);
              _nw3[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_126_v14).length));
              _nw3[(new BigNumber(7)).toNumber()] = _111_v1;
              _nw3[(new BigNumber(8)).toNumber()] = _111_v1;
              _nw3[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(753)), function (_128_i1) {
                return new _dafny.CodePoint('s'.codePointAt(0));
              })).length);
              _nw3[(new BigNumber(10)).toNumber()] = _111_v1;
              _nw3[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_112_v2, _112_v2)).length);
              _nw3[(new BigNumber(12)).toNumber()] = ((_dafny.ZERO).minus(_111_v1)).multipliedBy(new BigNumber((_113_v3).length));
              _nw3[(new BigNumber(13)).toNumber()] = _111_v1;
              _nw3[(new BigNumber(14)).toNumber()] = _111_v1;
              _nw3[(new BigNumber(15)).toNumber()] = ((false) ? (_111_v1) : (_111_v1));
              _nw3[(new BigNumber(16)).toNumber()] = _111_v1;
              _127_v15 = _nw3;
              let _index0 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_127_v15).length));
              (_127_v15)[_index0] = _111_v1;
              let _129_v16;
              _129_v16 = _module.D0.create_DC0(_111_v1);
              let _pat_let_tv0 = _111_v1;
              let _130_v17;
              let _nw4 = Array((new BigNumber(12)).toNumber());
              _nw4[(_dafny.ZERO).toNumber()] = _module.D0.create_DC0(_111_v1);
              _nw4[(_dafny.ONE).toNumber()] = _129_v16;
              _nw4[(new BigNumber(2)).toNumber()] = _129_v16;
              _nw4[(new BigNumber(3)).toNumber()] = _129_v16;
              _nw4[(new BigNumber(4)).toNumber()] = _129_v16;
              _nw4[(new BigNumber(5)).toNumber()] = _129_v16;
              _nw4[(new BigNumber(6)).toNumber()] = _module.D0.create_DC0(_111_v1);
              _nw4[(new BigNumber(7)).toNumber()] = _129_v16;
              _nw4[(new BigNumber(8)).toNumber()] = _129_v16;
              _nw4[(new BigNumber(9)).toNumber()] = _129_v16;
              _nw4[(new BigNumber(10)).toNumber()] = _129_v16;
              _nw4[(new BigNumber(11)).toNumber()] = function (_pat_let0_0) {
                return function (_131_dt__update__tmp_h0) {
                  return function (_pat_let1_0) {
                    return function (_132_dt__update_hcf0_h0) {
                      return _module.D0.create_DC0(_132_dt__update_hcf0_h0);
                    }(_pat_let1_0);
                  }(_pat_let_tv0);
                }(_pat_let0_0);
              }(_129_v16);
              _130_v17 = _nw4;
              let _133_v18;
              _133_v18 = _dafny.MultiSet.fromElements(_130_v17, _130_v17);
              let _index1 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_127_v15).length));
              let _rhs3 = _110_v0;
              let _rhs4 = (_133_v18).IsProperSubsetOf(((_133_v18).update(_130_v17, _module.__default.abs(_111_v1))).Difference(_133_v18));
              let _rhs5 = (_dafny.ZERO).minus((_111_v1).minus(_111_v1));
              let _lhs2 = _116_globalState;
              let _lhs3 = _116_globalState;
              let _lhs4 = _127_v15;
              let _lhs5 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_127_v15).length));
              _lhs2.f6 = _rhs3;
              _lhs3.f6 = _rhs4;
              _lhs4[_lhs5] = _rhs5;
              let _rhs6 = _127_v15;
              let _rhs7 = (_111_v1).isLessThan(new BigNumber((_119_v7).length));
              let _rhs8 = ((_127_v15)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_127_v15).length))]).multipliedBy(_111_v1);
              let _lhs6 = _116_globalState;
              _127_v15 = _rhs6;
              _lhs6.f6 = _rhs7;
              _111_v1 = _rhs8;
              let _index2 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_127_v15).length));
              (_127_v15)[_index2] = _111_v1;
            } else {
              let _134_v19;
              _134_v19 = _dafny.Seq.UnicodeFromString("prcfca");
              let _135_v20;
              _135_v20 = _dafny.Map.Empty.slice().updateUnsafe(_110_v0,_114_v4);
              let _136_v21;
              _136_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_module.__default.fm1(_116_globalState), _134_v19),_module.__default.fm2(new BigNumber((_dafny.Seq.UnicodeFromString("xngayfcy")).length), new BigNumber((_135_v20).length), _116_globalState));
              let _137_v22;
              _137_v22 = _dafny.Seq.of(_111_v1, _111_v1);
              _136_v21 = (_136_v21).update(((_110_v0) ? (_134_v19) : (_134_v19)), (_137_v22)[_module.__default.safeIndex((_dafny.ZERO).minus(_111_v1), new BigNumber((_137_v22).length))]);
              let _138_v23;
              _138_v23 = _dafny.MultiSet.fromElements(_module.__default.fm2(_111_v1, _111_v1, _116_globalState));
              let _139_v24;
              _139_v24 = _dafny.Map.Empty.slice().updateUnsafe(_138_v23,_110_v0);
              _139_v24 = (_139_v24).update(_dafny.MultiSet.FromArray(_137_v22), _110_v0);
              let _140_v25;
              _140_v25 = _dafny.Map.Empty.slice().updateUnsafe(false,_110_v0);
              let _141_v26;
              _141_v26 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_110_v0,_110_v0));
              let _142_v27;
              _142_v27 = _dafny.Map.Empty.slice().updateUnsafe(_111_v1,!(_110_v0));
              let _143_v28;
              _143_v28 = _module.D0.create_DC1(!(_110_v0), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_111_v1,_110_v0)).length), _110_v0);
              let _144_v29;
              let _nw5 = Array((new BigNumber(24)).toNumber());
              _nw5[(_dafny.ZERO).toNumber()] = _110_v0;
              _nw5[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(_140_v25, _dafny.Map.Empty.slice().updateUnsafe(_110_v0,_110_v0))).IsSubsetOf(_141_v26);
              _nw5[(new BigNumber(2)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(3)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(4)).toNumber()] = !(_110_v0);
              _nw5[(new BigNumber(5)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(6)).toNumber()] = _dafny.Seq.IsPrefixOf(_137_v22, _137_v22);
              _nw5[(new BigNumber(7)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(8)).toNumber()] = !(!(true));
              _nw5[(new BigNumber(9)).toNumber()] = (_138_v23).IsSubsetOf(_138_v23);
              _nw5[(new BigNumber(10)).toNumber()] = (true) || (_110_v0);
              _nw5[(new BigNumber(11)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(12)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(13)).toNumber()] = (_111_v1).isLessThan(new BigNumber(135));
              _nw5[(new BigNumber(14)).toNumber()] = true;
              _nw5[(new BigNumber(15)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(16)).toNumber()] = _module.__default.fm0(_116_globalState);
              _nw5[(new BigNumber(17)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(18)).toNumber()] = (new BigNumber(930)).isEqualTo(_111_v1);
              _nw5[(new BigNumber(19)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(20)).toNumber()] = (((_142_v27).contains(_111_v1)) ? ((_142_v27).get(_111_v1)) : (_110_v0));
              _nw5[(new BigNumber(21)).toNumber()] = (_143_v28).dtor_cf1;
              _nw5[(new BigNumber(22)).toNumber()] = _110_v0;
              _nw5[(new BigNumber(23)).toNumber()] = !(_142_v27).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_111_v1,(_112_v2)[_module.__default.safeIndex(_111_v1, new BigNumber((_112_v2).length))])).length),!(_module.__default.fm0(_116_globalState))));
              _144_v29 = _nw5;
              let _index3 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_144_v29).length));
              (_144_v29)[_index3] = _110_v0;
              let _index4 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_144_v29).length));
              (_144_v29)[_index4] = _110_v0;
              let _145_v30;
              _145_v30 = new _dafny.CodePoint('v'.codePointAt(0));
              let _146_v31;
              let _nw6 = new _module.C2();
              _nw6.__ctor(_145_v30, _111_v1);
              _146_v31 = _nw6;
              let _147_v32;
              let _148_v33;
              let _149_v34;
              let _150_v35;
              let _out0;
              let _out1;
              let _out2;
              let _out3;
              let _outcollector0 = (_146_v31).m8(_116_globalState);
              _out0 = _outcollector0[0];
              _out1 = _outcollector0[1];
              _out2 = _outcollector0[2];
              _out3 = _outcollector0[3];
              _147_v32 = _out0;
              _148_v33 = _out1;
              _149_v34 = _out2;
              _150_v35 = _out3;
            }
            if (_110_v0) {
              let _151_v36;
              _151_v36 = _dafny.Set.fromElements(_110_v0);
              let _152_v37;
              _152_v37 = _dafny.Set.fromElements(_110_v0, _110_v0, _110_v0, (_112_v2)[_module.__default.safeIndex(new BigNumber((_151_v36).length), new BigNumber((_112_v2).length))], _110_v0);
              let _153_v38;
              _153_v38 = _dafny.Seq.of(_151_v36);
              let _154_v39;
              _154_v39 = _dafny.Seq.UnicodeFromString("eua");
              let _155_v40;
              _155_v40 = _module.D7.create_DC13((_153_v38)[_module.__default.safeIndex(new BigNumber((_154_v39).length), new BigNumber((_153_v38).length))], _111_v1, _111_v1);
              let _pat_let_tv1 = _111_v1;
              let _pat_let_tv2 = _111_v1;
              let _pat_let_tv3 = _116_globalState;
              _111_v1 = (_111_v1).multipliedBy((function (_pat_let2_0) {
                return function (_156_dt__update__tmp_h1) {
                  return function (_pat_let3_0) {
                    return function (_157_dt__update_hcf19_h0) {
                      return _module.D7.create_DC13((_156_dt__update__tmp_h1).dtor_cf18, _157_dt__update_hcf19_h0, (_156_dt__update__tmp_h1).dtor_cf20);
                    }(_pat_let3_0);
                  }(_module.__default.fm2(_pat_let_tv1, _pat_let_tv2, _pat_let_tv3));
                }(_pat_let2_0);
              }(_155_v40)).dtor_cf19);
              _154_v39 = _module.__default.fm1(_116_globalState);
              let _158_v41;
              _158_v41 = new _dafny.CodePoint('n'.codePointAt(0));
              let _159_v42;
              let _init0 = ((_160_v41) => function (_161_i2) {
                return _160_v41;
              })(_158_v41);
              let _nw7 = Array((new BigNumber(8)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw7.length); _i0_0++) {
                _nw7[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _159_v42 = _nw7;
              let _162_v43;
              _162_v43 = _module.D21.create_DC50(_159_v42, (_dafny.ZERO).minus(_111_v1), _110_v0);
              let _163_v44;
              _163_v44 = _dafny.Seq.of(_154_v39);
              let _164_v45;
              _164_v45 = _dafny.Seq.of(_111_v1, _111_v1, _111_v1, _111_v1, new BigNumber((_163_v44).length));
              let _165_v46;
              let _nw8 = Array((new BigNumber(24)).toNumber());
              _nw8[(_dafny.ZERO).toNumber()] = true;
              _nw8[(_dafny.ONE).toNumber()] = !_dafny.Seq.contains(_154_v39, _158_v41);
              _nw8[(new BigNumber(2)).toNumber()] = _module.__default.fm0(_116_globalState);
              _nw8[(new BigNumber(3)).toNumber()] = _110_v0;
              _nw8[(new BigNumber(4)).toNumber()] = _module.__default.fm0(_116_globalState);
              _nw8[(new BigNumber(5)).toNumber()] = _110_v0;
              _nw8[(new BigNumber(6)).toNumber()] = (_162_v43).dtor_cf85;
              _nw8[(new BigNumber(7)).toNumber()] = (_112_v2)[_module.__default.safeIndex(_111_v1, new BigNumber((_112_v2).length))];
              _nw8[(new BigNumber(8)).toNumber()] = _110_v0;
              _nw8[(new BigNumber(9)).toNumber()] = !(_110_v0);
              _nw8[(new BigNumber(10)).toNumber()] = ((_164_v45)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_110_v0)).cardinality()), new BigNumber((_164_v45).length))]).isLessThan(_111_v1);
              _nw8[(new BigNumber(11)).toNumber()] = _110_v0;
              _nw8[(new BigNumber(12)).toNumber()] = _110_v0;
              _nw8[(new BigNumber(13)).toNumber()] = _module.__default.fm0(_116_globalState);
              _nw8[(new BigNumber(14)).toNumber()] = _110_v0;
              _nw8[(new BigNumber(15)).toNumber()] = (_111_v1).isEqualTo(new BigNumber(305));
              _nw8[(new BigNumber(16)).toNumber()] = !(_110_v0);
              _nw8[(new BigNumber(17)).toNumber()] = _110_v0;
              _nw8[(new BigNumber(18)).toNumber()] = true;
              _nw8[(new BigNumber(19)).toNumber()] = (!(_110_v0)) && (_110_v0);
              _nw8[(new BigNumber(20)).toNumber()] = true;
              _nw8[(new BigNumber(21)).toNumber()] = _module.__default.fm0(_116_globalState);
              _nw8[(new BigNumber(22)).toNumber()] = _110_v0;
              _nw8[(new BigNumber(23)).toNumber()] = _110_v0;
              _165_v46 = _nw8;
              _165_v46 = _165_v46;
              let _166_v47;
              _166_v47 = _dafny.MultiSet.fromElements((_111_v1).isLessThanOrEqualTo(_111_v1));
              let _167_v48;
              _167_v48 = _dafny.Map.Empty.slice().updateUnsafe(_110_v0,_110_v0);
              let _168_v49;
              _168_v49 = _dafny.Map.Empty.slice().updateUnsafe(_110_v0,_167_v48);
              let _169_v50;
              _169_v50 = _module.D16.create_DC39(_111_v1, _110_v0, (((_168_v49).contains(_110_v0)) ? ((_168_v49).get(_110_v0)) : (_167_v48)), _158_v41);
              let _170_v51;
              _170_v51 = _module.D6.create_DC10(_112_v2);
              let _171_v52;
              _171_v52 = _module.D19.create_DC43(_110_v0, (_169_v50).dtor_cf65, _110_v0, _170_v51, _154_v39);
              let _rhs9 = (_171_v52).dtor_cf70;
              let _rhs10 = _110_v0;
              let _rhs11 = _166_v47;
              _158_v41 = _rhs9;
              _110_v0 = _rhs10;
              _166_v47 = _rhs11;
              let _172_v53;
              let _init1 = ((_173_v1) => function (_174_i3) {
                return (_174_i3).multipliedBy(_173_v1);
              })(_111_v1);
              let _nw9 = Array((new BigNumber(24)).toNumber());
              for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw9.length); _i0_1++) {
                _nw9[_i0_1] = _init1(new BigNumber(_i0_1));
              }
              _172_v53 = _nw9;
              let _175_v54;
              let _nw10 = Array((new BigNumber(29)).toNumber());
              _nw10[(_dafny.ZERO).toNumber()] = _172_v53;
              _nw10[(_dafny.ONE).toNumber()] = _172_v53;
              _nw10[(new BigNumber(2)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(3)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(4)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(5)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(6)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(7)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(8)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(9)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(10)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(11)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(12)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(13)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(14)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(15)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(16)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(17)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(18)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(19)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(20)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(21)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(22)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(23)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(24)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(25)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(26)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(27)).toNumber()] = _172_v53;
              _nw10[(new BigNumber(28)).toNumber()] = _172_v53;
              _175_v54 = _nw10;
              let _176_v55;
              _176_v55 = _175_v54;
              let _177_v56;
              let _nw11 = Array((new BigNumber(23)).toNumber());
              _nw11[(_dafny.ZERO).toNumber()] = _175_v54;
              _nw11[(_dafny.ONE).toNumber()] = _176_v55;
              _nw11[(new BigNumber(2)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(3)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(4)).toNumber()] = _175_v54;
              _nw11[(new BigNumber(5)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(6)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(7)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(8)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(9)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(10)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(11)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(12)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(13)).toNumber()] = _175_v54;
              _nw11[(new BigNumber(14)).toNumber()] = _175_v54;
              _nw11[(new BigNumber(15)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(16)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(17)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(18)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(19)).toNumber()] = _175_v54;
              _nw11[(new BigNumber(20)).toNumber()] = _175_v54;
              _nw11[(new BigNumber(21)).toNumber()] = _176_v55;
              _nw11[(new BigNumber(22)).toNumber()] = _176_v55;
              _177_v56 = _nw11;
              let _178_v57;
              _178_v57 = _dafny.Map.Empty.slice().updateUnsafe(_177_v56,_172_v53);
              let _179_v58;
              _179_v58 = _module.D11.create_DC22(_178_v57);
              _179_v58 = _179_v58;
            } else {
              let _180_v59;
              _180_v59 = new _dafny.CodePoint('k'.codePointAt(0));
              let _181_v60;
              _181_v60 = _dafny.Seq.of(_180_v59, _180_v59, _180_v59, _180_v59, _180_v59);
              (_116_globalState).f1 = new BigNumber((_dafny.MultiSet.FromArray(_181_v60)).cardinality());
              _110_v0 = _110_v0;
              let _182_v61;
              _182_v61 = _dafny.MultiSet.fromElements(_112_v2);
              let _183_v62;
              let _nw12 = Array((new BigNumber(16)).toNumber());
              _nw12[(_dafny.ZERO).toNumber()] = (_182_v61).Union(_182_v61);
              _nw12[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements(_112_v2, _112_v2, _112_v2)).Difference(_dafny.MultiSet.fromElements(_112_v2, _112_v2, _112_v2, _112_v2, _dafny.Seq.of(!(_110_v0))));
              _nw12[(new BigNumber(2)).toNumber()] = _182_v61;
              _nw12[(new BigNumber(3)).toNumber()] = (_182_v61).update(_112_v2, _module.__default.abs(_module.__default.fm2(_111_v1, _111_v1, _116_globalState)));
              _nw12[(new BigNumber(4)).toNumber()] = _182_v61;
              _nw12[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(_112_v2);
              _nw12[(new BigNumber(6)).toNumber()] = (_182_v61).Difference(_182_v61);
              _nw12[(new BigNumber(7)).toNumber()] = _182_v61;
              _nw12[(new BigNumber(8)).toNumber()] = _182_v61;
              _nw12[(new BigNumber(9)).toNumber()] = _182_v61;
              _nw12[(new BigNumber(10)).toNumber()] = (_182_v61).Intersect(_dafny.MultiSet.fromElements(_112_v2, _dafny.Seq.update(_112_v2, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_110_v0)).length), new BigNumber((_112_v2).length)), true)));
              _nw12[(new BigNumber(11)).toNumber()] = _182_v61;
              _nw12[(new BigNumber(12)).toNumber()] = _182_v61;
              _nw12[(new BigNumber(13)).toNumber()] = _182_v61;
              _nw12[(new BigNumber(14)).toNumber()] = _182_v61;
              _nw12[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.fromElements(_dafny.Seq.of(_110_v0))).Union(_182_v61);
              _183_v62 = _nw12;
              let _index5 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_183_v62).length));
              (_183_v62)[_index5] = (_dafny.MultiSet.fromElements(_112_v2, _112_v2)).Intersect(_182_v61);
              let _184_v63;
              _184_v63 = _dafny.Seq.of(_dafny.Seq.of(_110_v0, _110_v0, true, _110_v0, _110_v0));
              let _185_v64;
              _185_v64 = _dafny.MultiSet.fromElements(false);
              let _186_v65;
              _186_v65 = _dafny.MultiSet.fromElements(_111_v1, _111_v1, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ymwi"), _module.__default.safeIndex(_111_v1, new BigNumber((_dafny.Seq.UnicodeFromString("ymwi")).length)), _180_v59)).length));
              let _187_v66;
              _187_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_186_v65).cardinality()),_112_v2);
              let _index6 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_183_v62).length));
              (_183_v62)[_index6] = (_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.Concat(_184_v63, _dafny.Seq.of(_112_v2)), _module.__default.safeIndex(new BigNumber((_185_v64).cardinality()), new BigNumber((_dafny.Seq.Concat(_184_v63, _dafny.Seq.of(_112_v2))).length)), (((_187_v66).contains(_111_v1)) ? ((_187_v66).get(_111_v1)) : (_112_v2))))).update(_112_v2, _module.__default.abs(_module.__default.fm2(_111_v1, _111_v1, _116_globalState)));
              let _188_v67;
              let _nw13 = Array((new BigNumber(7)).toNumber());
              _nw13[(_dafny.ZERO).toNumber()] = _110_v0;
              _nw13[(_dafny.ONE).toNumber()] = !_dafny.Seq.contains(_112_v2, _110_v0);
              _nw13[(new BigNumber(2)).toNumber()] = !(_110_v0);
              _nw13[(new BigNumber(3)).toNumber()] = ((_110_v0) ? (_110_v0) : (_110_v0));
              _nw13[(new BigNumber(4)).toNumber()] = _110_v0;
              _nw13[(new BigNumber(5)).toNumber()] = _110_v0;
              _nw13[(new BigNumber(6)).toNumber()] = _110_v0;
              _188_v67 = _nw13;
              let _index7 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_188_v67).length));
              (_188_v67)[_index7] = _110_v0;
              let _189_v68;
              _189_v68 = _dafny.Set.fromElements(_110_v0, _110_v0);
              let _190_v69;
              let _nw14 = new _module.C8();
              _nw14.__ctor(_110_v0, _189_v68, new BigNumber((_186_v65).cardinality()));
              _190_v69 = _nw14;
              let _191_v70;
              _191_v70 = _dafny.Set.fromElements(_190_v69);
              let _index8 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_188_v67).length));
              let _rhs12 = _module.__default.safeDivisionInt(new BigNumber((_112_v2).length), (_dafny.ZERO).minus((new BigNumber((_191_v70).length)).minus(_111_v1)));
              let _rhs13 = (((_113_v3).contains((new BigNumber((_dafny.MultiSet.fromElements((_190_v69).f11)).cardinality())).minus((_190_v69).f11))) ? ((_113_v3).get((new BigNumber((_dafny.MultiSet.fromElements((_190_v69).f11)).cardinality())).minus((_190_v69).f11))) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_110_v0)).length))));
              let _rhs14 = (_111_v1).isLessThan(_module.__default.safeDivisionInt(_111_v1, _111_v1));
              let _lhs7 = _116_globalState;
              let _lhs8 = _116_globalState;
              let _lhs9 = _188_v67;
              let _lhs10 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_188_v67).length));
              _lhs7.f1 = _rhs12;
              _lhs8.f1 = _rhs13;
              _lhs9[_lhs10] = _rhs14;
              let _192_v71;
              _192_v71 = _module.D0.create_DC1(false, (_190_v69).f11, _110_v0);
              let _193_v72;
              let _nw15 = Array((_dafny.ONE).toNumber());
              _nw15[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_111_v1, (_192_v71).dtor_cf2);
              _193_v72 = _nw15;
              let _index9 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_193_v72).length));
              (_193_v72)[_index9] = _module.__default.safeModuloInt((_190_v69).f11, new BigNumber((_dafny.Seq.UnicodeFromString("day")).length));
              let _index10 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_193_v72).length));
              let _rhs15 = _module.__default.safeModuloInt((_190_v69).f11, _111_v1);
              let _rhs16 = _180_v59;
              let _rhs17 = (_190_v69).f11;
              let _lhs11 = _193_v72;
              let _lhs12 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_193_v72).length));
              _lhs11[_lhs12] = _rhs15;
              _180_v59 = _rhs16;
              _111_v1 = _rhs17;
            }
            let _194_v73;
            _194_v73 = _dafny.Seq.of(_111_v1);
            let _195_v74;
            _195_v74 = _dafny.MultiSet.fromElements(_111_v1);
            let _196_v75;
            let _nw16 = Array((new BigNumber(4)).toNumber());
            _nw16[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.FromArray(_194_v73)).update(_111_v1, _module.__default.abs(_111_v1));
            _nw16[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(_111_v1, _111_v1);
            _nw16[(new BigNumber(2)).toNumber()] = _195_v74;
            _nw16[(new BigNumber(3)).toNumber()] = _195_v74;
            _196_v75 = _nw16;
            let _197_v76;
            _197_v76 = _196_v75;
            _197_v76 = _197_v76;
          }
        }
      }
      let _198_v77;
      _198_v77 = _dafny.Seq.UnicodeFromString("lyaxksq");
      let _199_v78;
      _199_v78 = _module.D13.create_DC28(_dafny.Seq.Concat(_198_v77, _198_v77), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_111_v1, new BigNumber(668))), _110_v0);
      let _pat_let_tv4 = _110_v0;
      let _pat_let_tv5 = _198_v77;
      _199_v78 = function (_pat_let4_0) {
        return function (_200_dt__update__tmp_h3) {
          return function (_pat_let5_0) {
            return function (_201_dt__update_hcf42_h0) {
              return function (_pat_let6_0) {
                return function (_202_dt__update_hcf40_h0) {
                  return _module.D13.create_DC28(_202_dt__update_hcf40_h0, (_200_dt__update__tmp_h3).dtor_cf41, _201_dt__update_hcf42_h0);
                }(_pat_let6_0);
              }(_pat_let_tv5);
            }(_pat_let5_0);
          }(_pat_let_tv4);
        }(_pat_let4_0);
      }(_199_v78);
      if (_110_v0) {
        let _203_v79;
        let _init2 = ((_204_v1) => function (_205_i4) {
          return (_205_i4).plus(_204_v1);
        })(_111_v1);
        let _nw17 = Array((new BigNumber(2)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw17.length); _i0_2++) {
          _nw17[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _203_v79 = _nw17;
        let _206_v80;
        _206_v80 = _module.D19.create_DC42(_203_v79);
        _203_v79 = (_206_v80).dtor_cf68;
        let _207_v81;
        let _nw18 = Array((new BigNumber(17)).toNumber());
        _207_v81 = _nw18;
        let _208_v82;
        _208_v82 = _dafny.MultiSet.fromElements(_111_v1, _111_v1);
        let _209_v83;
        _209_v83 = _dafny.Set.fromElements(_110_v0, _110_v0);
        let _210_v84;
        let _nw19 = new _module.C4();
        _nw19.__ctor(_208_v82, _110_v0, _209_v83, new BigNumber((_113_v3).length));
        _210_v84 = _nw19;
        let _211_v85;
        _211_v85 = _dafny.Map.Empty.slice().updateUnsafe(_203_v79,_210_v84);
        let _index11 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_207_v81).length));
        (_207_v81)[_index11] = (((_211_v85).contains(_203_v79)) ? ((_211_v85).get(_203_v79)) : (_210_v84));
        let _212_v86;
        _212_v86 = _dafny.Seq.of(new BigNumber((_112_v2).length));
        let _index12 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_207_v81).length));
        let _rhs18 = _210_v84;
        let _rhs19 = (_210_v84).f17;
        let _rhs20 = _dafny.Seq.Concat(_212_v86, _212_v86);
        let _lhs13 = _207_v81;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_207_v81).length));
        let _lhs15 = _116_globalState;
        _lhs13[_lhs14] = _rhs18;
        _lhs15.f7 = _rhs19;
        _212_v86 = _rhs20;
        let _213_v87;
        let _nw20 = new _module.C1();
        _nw20.__ctor(_110_v0, (_210_v84).f17);
        _213_v87 = _nw20;
        let _214_v88;
        _214_v88 = _dafny.Set.fromElements(_213_v87, _213_v87, _213_v87);
        if ((_module.__default.safeModuloInt(new BigNumber((_214_v88).length), _111_v1)).isEqualTo(_111_v1)) {
          (_116_globalState).f6 = (_210_v84).f17;
          (_116_globalState).f1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_210_v84).f17), _112_v2)).length);
          let _215_v89;
          let _nw21 = Array((new BigNumber(11)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = _198_v77;
          _nw21[(_dafny.ONE).toNumber()] = _module.__default.fm1(_116_globalState);
          _nw21[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(722)), function (_216_i5) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          });
          _nw21[(new BigNumber(3)).toNumber()] = _198_v77;
          _nw21[(new BigNumber(4)).toNumber()] = _198_v77;
          _nw21[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(181)), function (_217_i6) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          });
          _nw21[(new BigNumber(6)).toNumber()] = _198_v77;
          _nw21[(new BigNumber(7)).toNumber()] = _module.__default.fm1(_116_globalState);
          _nw21[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("ltloi");
          _nw21[(new BigNumber(9)).toNumber()] = _198_v77;
          _nw21[(new BigNumber(10)).toNumber()] = _198_v77;
          _215_v89 = _nw21;
          let _index13 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_215_v89).length));
          (_215_v89)[_index13] = _dafny.Seq.UnicodeFromString("brktfp");
          let _218_v90;
          _218_v90 = _dafny.MultiSet.fromElements(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("tk"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(121)), function (_219_i7) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          })));
          let _220_v91;
          _220_v91 = _module.D24.create_DC58((_210_v84).f16);
          let _221_v92;
          _221_v92 = _dafny.Map.Empty.slice().updateUnsafe((_220_v91).dtor_cf96,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_212_v86).length),_module.__default.fm2(_111_v1, new BigNumber((_112_v2).length), _116_globalState)));
          let _index14 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_215_v89).length));
          let _rhs21 = _dafny.Seq.Concat(_198_v77, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-425)), function (_222_i8) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }), _198_v77));
          let _rhs22 = _dafny.Seq.Concat(_198_v77, _dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), function (_223_i9) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          }));
          let _rhs23 = (_210_v84).fm4((_111_v1).multipliedBy(_111_v1), _221_v92, _111_v1, _110_v0, _116_globalState);
          let _rhs24 = _111_v1;
          let _rhs25 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of((_213_v87).f18), _112_v2));
          let _lhs16 = _215_v89;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_215_v89).length));
          let _lhs18 = _116_globalState;
          let _lhs19 = _116_globalState;
          _198_v77 = _rhs21;
          _lhs16[_lhs17] = _rhs22;
          _lhs18.f7 = _rhs23;
          _lhs19.f1 = _rhs24;
          _218_v90 = _rhs25;
          let _224_v93;
          _224_v93 = new _dafny.CodePoint('a'.codePointAt(0));
          let _225_v94;
          _225_v94 = _dafny.Map.Empty.slice().updateUnsafe(_111_v1,(_210_v84).f16);
          let _226_v95;
          _226_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_225_v94).length),(_213_v87).f19);
          let _227_v96;
          _227_v96 = _dafny.Map.Empty.slice().updateUnsafe(_224_v93,_226_v95);
          let _228_v97;
          let _229_v98;
          let _out4;
          let _out5;
          let _outcollector1 = (_210_v84).m1(new BigNumber(((((_227_v96).contains(_224_v93)) ? ((_227_v96).get(_224_v93)) : (_226_v95))).length), _116_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _228_v97 = _out4;
          _229_v98 = _out5;
          let _rhs26 = new BigNumber((_198_v77).length);
          let _rhs27 = (_111_v1).plus(_111_v1);
          let _rhs28 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_111_v1, (_dafny.ZERO).minus(_111_v1)));
          let _rhs29 = new BigNumber(271);
          let _rhs30 = (_213_v87).f18;
          let _lhs20 = _116_globalState;
          let _lhs21 = _116_globalState;
          let _lhs22 = _116_globalState;
          let _lhs23 = _116_globalState;
          let _lhs24 = _116_globalState;
          _lhs20.f1 = _rhs26;
          _lhs21.f1 = _rhs27;
          _lhs22.f1 = _rhs28;
          _lhs23.f1 = _rhs29;
          _lhs24.f7 = _rhs30;
        } else {
          let _230_v99;
          _230_v99 = _dafny.Map.Empty.slice().updateUnsafe((_213_v87).f18,(_213_v87).f18);
          _230_v99 = (_230_v99).update(true, (_213_v87).f18);
          let _231_v100;
          let _232_v101;
          let _out6;
          let _out7;
          let _outcollector2 = (_210_v84).m1(_111_v1, _116_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _231_v100 = _out6;
          _232_v101 = _out7;
          let _index15 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_203_v79).length));
          (_203_v79)[_index15] = _module.__default.safeDivisionInt(_111_v1, _111_v1);
          let _233_v102;
          _233_v102 = _dafny.Set.fromElements(_111_v1);
          let _index16 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_203_v79).length));
          (_203_v79)[_index16] = (_210_v84).fm3(_111_v1, (_module.__default.fm2(new BigNumber((_dafny.Seq.UnicodeFromString("a")).length), _111_v1, _116_globalState)).isLessThan((_dafny.ZERO).minus(new BigNumber((_233_v102).length))), (_213_v87).f19, _116_globalState);
          let _234_v103;
          _234_v103 = new _dafny.CodePoint('p'.codePointAt(0));
          let _235_v104;
          let _nw22 = new _module.C2();
          _nw22.__ctor(_234_v103, _111_v1);
          _235_v104 = _nw22;
          let _236_v105;
          let _init3 = ((_237_v87) => function (_238_i10) {
            return (_237_v87).f19;
          })(_213_v87);
          let _nw23 = Array((new BigNumber(14)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw23.length); _i0_3++) {
            _nw23[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _236_v105 = _nw23;
          let _239_v106;
          _239_v106 = _dafny.Map.Empty.slice().updateUnsafe(_236_v105,new BigNumber((_212_v86).length));
          let _nw24 = new _module.C2();
          _nw24.__ctor(_234_v103, (((_239_v106).contains(_236_v105)) ? ((_239_v106).get(_236_v105)) : (_111_v1)));
          _235_v104 = _nw24;
          let _index17 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_203_v79).length));
          (_203_v79)[_index17] = new BigNumber(146);
        }
        let _240_v107;
        let _241_v108;
        let _out8;
        let _out9;
        let _outcollector3 = (_210_v84).m1(_111_v1, _116_globalState);
        _out8 = _outcollector3[0];
        _out9 = _outcollector3[1];
        _240_v107 = _out8;
        _241_v108 = _out9;
        let _index18 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_203_v79).length));
        (_203_v79)[_index18] = (_210_v84).fm6(_110_v0, _111_v1, _116_globalState);
        let _index19 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_203_v79).length));
        let _rhs31 = (_111_v1).minus(_module.__default.safeDivisionInt(new BigNumber((_198_v77).length), _111_v1));
        let _rhs32 = _111_v1;
        let _rhs33 = (_111_v1).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("s")).length));
        let _lhs25 = _203_v79;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_203_v79).length));
        let _lhs27 = _116_globalState;
        _lhs25[_lhs26] = _rhs31;
        _111_v1 = _rhs32;
        _lhs27.f1 = _rhs33;
      } else {
        let _242_v109;
        _242_v109 = _dafny.MultiSet.fromElements(_111_v1);
        let _243_v110;
        _243_v110 = _dafny.Set.fromElements(_110_v0);
        let _244_v111;
        let _nw25 = new _module.C3();
        _nw25.__ctor(((_110_v0) ? (_111_v1) : (new BigNumber((_114_v4).length))), _111_v1, _module.__default.safeModuloInt(new BigNumber((_242_v109).cardinality()), _111_v1), _243_v110);
        _244_v111 = _nw25;
        _244_v111 = _244_v111;
        let _245_v112;
        let _nw26 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _245_v112 = _nw26;
        let _246_v113;
        _246_v113 = new _dafny.CodePoint('r'.codePointAt(0));
        let _247_v114;
        _247_v114 = _dafny.Map.Empty.slice().updateUnsafe(_245_v112,_246_v113);
        let _index20 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length));
        (_245_v112)[_index20] = (_244_v111.f23).minus((((_113_v3).contains(new BigNumber((_247_v114).length))) ? ((_113_v3).get(new BigNumber((_247_v114).length))) : ((_244_v111).f22)));
        let _index21 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length));
        (_245_v112)[_index21] = (_244_v111).f22;
        _111_v1 = _module.__default.safeDivisionInt(new BigNumber(-630), ((_110_v0) ? ((_244_v111).f22) : ((_dafny.ZERO).minus((_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))]))));
        let _248_v115;
        let _249_v116;
        let _out10;
        let _out11;
        let _outcollector4 = (_244_v111).m1((_244_v111).f22, _116_globalState);
        _out10 = _outcollector4[0];
        _out11 = _outcollector4[1];
        _248_v115 = _out10;
        _249_v116 = _out11;
        let _250_v117;
        _250_v117 = _dafny.Seq.of(_111_v1);
        let _source5 = _module.D7.create_DC13((_243_v110).Union(_243_v110), (_244_v111).f22, new BigNumber((_250_v117).length));
        if (_source5.is_DC13) {
          let _251___mcc_h0 = (_source5).cf18;
          let _252___mcc_h1 = (_source5).cf19;
          let _253___mcc_h2 = (_source5).cf20;
          let _254_cf20 = _253___mcc_h2;
          let _255_cf19 = _252___mcc_h1;
          let _256_cf18 = _251___mcc_h0;
          let _index22 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length));
          (_245_v112)[_index22] = (_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))];
          let _257_v118;
          _257_v118 = _module.D22.create_DC52();
          _257_v118 = _257_v118;
          let _258_v119;
          _258_v119 = _dafny.Map.Empty.slice().updateUnsafe(_110_v0,true);
          let _259_v120;
          _259_v120 = _dafny.Set.fromElements((((_242_v109).contains(_module.__default.fm2((_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))], (_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))], _116_globalState))) ? ((_242_v109).get(_module.__default.fm2((_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))], (_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))], _116_globalState))) : (_254_cf20)), _244_v111.f23, (new BigNumber(866)).plus(new BigNumber(72)), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), ((_260_v1) => function (_261_i11) {
            return _260_v1;
          })(_111_v1)), _module.__default.safeIndex(new BigNumber((_258_v119).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), ((_262_v1) => function (_263_i11) {
            return _262_v1;
          })(_111_v1))).length)), (_dafny.ZERO).minus(_244_v111.f23))).length), _111_v1), new BigNumber((_258_v119).length));
          _259_v120 = _dafny.Set.fromElements(_111_v1);
          let _264_v122;
          _264_v122 = _dafny.Seq.of(_dafny.Set.fromElements((((_242_v109).contains((_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))])) ? ((_242_v109).get((_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))])) : ((_244_v111).f22))), _dafny.Set.fromElements(new BigNumber(261)), function () {
            let _coll28 = new _dafny.Set();
            for (const _compr_28 of _dafny.IntegerRange(new BigNumber(113), new BigNumber(-958))) {
              let _265_v121 = _compr_28;
              if (((new BigNumber(113)).isLessThanOrEqualTo(_265_v121)) && ((_265_v121).isLessThan(new BigNumber(-958)))) {
                _coll28.add((_265_v121).multipliedBy(_111_v1));
              }
            }
            return _coll28;
          }());
          _264_v122 = _264_v122;
        } else if (_source5.is_DC12) {
          let _266___mcc_h3 = (_source5).cf17;
          let _267_cf17 = _266___mcc_h3;
          let _268_v123;
          _268_v123 = _dafny.Map.Empty.slice().updateUnsafe(_110_v0,_dafny.Seq.Concat(_198_v77, _198_v77));
          _268_v123 = (_268_v123).update(false, _dafny.Seq.Concat(_198_v77, _198_v77));
          (_116_globalState).f6 = !((_module.__default.fm2(new BigNumber((_198_v77).length), _244_v111.f23, _116_globalState)).isLessThan((_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))]));
          let _269_v124;
          let _nw27 = Array((new BigNumber(26)).toNumber()).fill(false);
          _269_v124 = _nw27;
          let _index23 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_269_v124).length));
          (_269_v124)[_index23] = _110_v0;
          let _index24 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_269_v124).length));
          (_269_v124)[_index24] = _module.__default.fm0(_116_globalState);
          let _270_v125;
          let _nw28 = new _module.C4();
          _nw28.__ctor(_242_v109, false, _243_v110, (_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))]);
          _270_v125 = _nw28;
          let _271_v126;
          _271_v126 = _module.D25.create_DC61(_270_v125);
          _270_v125 = (_271_v126).dtor_cf98;
        } else {
          let _272___mcc_h4 = (_source5).cf21;
          let _273_cf21 = _272___mcc_h4;
          let _274_v127;
          _274_v127 = _module.D6.create_DC10(_dafny.Seq.of(false, _110_v0));
          let _275_v128;
          _275_v128 = _module.D19.create_DC43(_110_v0, _246_v113, true, _274_v127, _198_v77);
          let _276_v129;
          let _nw29 = Array((new BigNumber(10)).toNumber());
          _nw29[(_dafny.ZERO).toNumber()] = true;
          _nw29[(_dafny.ONE).toNumber()] = _110_v0;
          _nw29[(new BigNumber(2)).toNumber()] = !((_244_v111).fm5(_module.D0.create_DC0(_244_v111.f23), _112_v2, _111_v1, (_275_v128).dtor_cf70, _116_globalState));
          _nw29[(new BigNumber(3)).toNumber()] = _110_v0;
          _nw29[(new BigNumber(4)).toNumber()] = _110_v0;
          _nw29[(new BigNumber(5)).toNumber()] = true;
          _nw29[(new BigNumber(6)).toNumber()] = _110_v0;
          _nw29[(new BigNumber(7)).toNumber()] = _110_v0;
          _nw29[(new BigNumber(8)).toNumber()] = false;
          _nw29[(new BigNumber(9)).toNumber()] = false;
          _276_v129 = _nw29;
          let _277_v130;
          _277_v130 = _dafny.Set.fromElements(_276_v129, _276_v129, _276_v129);
          _277_v130 = (_dafny.Set.fromElements(_276_v129, _276_v129, _276_v129)).Union(_277_v130);
          _111_v1 = _111_v1;
          (_116_globalState).f6 = ((_110_v0) ? (_110_v0) : (!(function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of (_module.__default.fm27(_110_v0, _110_v0, _116_globalState)).Elements) {
              let _278_v131 = _compr_29;
              if (_dafny.Seq.contains(_module.__default.fm27(_110_v0, _110_v0, _116_globalState), _278_v131)) {
                _coll29.add(_module.__default.safeDivisionInt(_278_v131, new BigNumber(-307)));
              }
            }
            return _coll29;
          }()).contains(_244_v111.f23)));
          let _279_v132;
          let _nw30 = new _module.C7();
          _nw30.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), ((_280_v111, _281_v1) => function (_282_i12) {
            return _dafny.Seq.of(_280_v111.f23, _281_v1);
          })(_244_v111, _111_v1)), _243_v110, (_245_v112)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_245_v112).length))]);
          _279_v132 = _nw30;
        }
      }
      let _283_v133;
      _283_v133 = new _dafny.CodePoint('o'.codePointAt(0));
      let _284_v134;
      _284_v134 = _module.D6.create_DC10(_dafny.Seq.of(_110_v0, _110_v0));
      (_116_globalState).f7 = (_module.D19.create_DC43(false, _283_v133, _110_v0, _284_v134, _198_v77)).dtor_cf71;
      let _285_i13;
      _285_i13 = _dafny.ZERO;
      L1: {
        while (_110_v0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_285_i13)) {
              break L1;
            }
            _285_i13 = (_285_i13).plus(_dafny.ONE);
            let _286_v135;
            let _nw31 = new _module.C5();
            _nw31.__ctor();
            _286_v135 = _nw31;
            let _rhs34 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(_110_v0, true, _110_v0), _dafny.Seq.of(_110_v0)), _112_v2);
            let _rhs35 = _111_v1;
            let _rhs36 = _111_v1;
            let _rhs37 = _286_v135;
            let _rhs38 = _module.__default.safeDivisionInt(((true) ? (_111_v1) : (_111_v1)), _111_v1);
            let _lhs28 = _116_globalState;
            let _lhs29 = _116_globalState;
            let _lhs30 = _116_globalState;
            let _lhs31 = _116_globalState;
            _lhs28.f6 = _rhs34;
            _lhs29.f1 = _rhs35;
            _lhs30.f1 = _rhs36;
            _286_v135 = _rhs37;
            _lhs31.f1 = _rhs38;
            let _287_v136;
            _287_v136 = _dafny.Set.fromElements(_110_v0, _110_v0, _110_v0, _110_v0);
            let _288_v137;
            let _nw32 = new _module.C7();
            _nw32.__ctor(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(269)), ((_289_v0) => function (_290_i14) {
              return _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_289_v0)).length));
            })(_110_v0)), _module.__default.fm56(_111_v1, _111_v1, _116_globalState)), _287_v136, (_111_v1).plus(_111_v1));
            _288_v137 = _nw32;
            _288_v137 = _288_v137;
            let _291_v138;
            let _nw33 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
            _291_v138 = _nw33;
            let _index25 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_291_v138).length));
            (_291_v138)[_index25] = (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_111_v1)).length));
            let _292_v139;
            let _nw34 = new _module.C0();
            _nw34.__ctor(_112_v2, _111_v1);
            _292_v139 = _nw34;
            let _293_v140;
            _293_v140 = _dafny.MultiSet.fromElements(_292_v139, _292_v139);
            let _index26 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_291_v138).length));
            (_291_v138)[_index26] = _module.__default.safeDivisionInt(_111_v1, new BigNumber(((_293_v140).Intersect(_293_v140)).cardinality()));
            (_116_globalState).f7 = _110_v0;
          }
        }
      }
      (_116_globalState).f7 = true;
      (_116_globalState).f1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(812)), ((_294_v133) => function (_295_i15) {
        return _294_v133;
      })(_283_v133)), _dafny.Seq.update(_198_v77, _module.__default.safeIndex(_111_v1, new BigNumber((_198_v77).length)), _283_v133))).length);
      let _296_v141;
      let _nw35 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
      _296_v141 = _nw35;
      let _index27 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length));
      (_296_v141)[_index27] = new BigNumber((_112_v2).length);
      let _297_v142;
      _297_v142 = _dafny.Seq.of(_111_v1);
      let _298_v143;
      _298_v143 = _dafny.Seq.of(_111_v1, (_dafny.ZERO).minus(_111_v1), new BigNumber((_297_v142).length), _111_v1);
      let _index28 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length));
      let _rhs39 = _111_v1;
      let _rhs40 = !(!(!_dafny.areEqual(_297_v142, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-502)), ((_299_v1) => function (_300_i16) {
        return _299_v1;
      })(_111_v1)))));
      let _lhs32 = _296_v141;
      let _lhs33 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length));
      let _lhs34 = _116_globalState;
      _lhs32[_lhs33] = _rhs39;
      _lhs34.f7 = _rhs40;
      let _301_v145;
      _301_v145 = _module.D8.create_DC16(true, _111_v1);
      let _302_v146;
      _302_v146 = _module.D8.create_DC17(_301_v145);
      let _303_v147;
      _303_v147 = _dafny.Seq.of(_302_v146, _302_v146);
      let _304_v148;
      _304_v148 = _dafny.Map.Empty.slice().updateUnsafe(_283_v133,function () {
        let _coll30 = new _dafny.Map();
        for (const _compr_30 of (_dafny.MultiSet.FromArray(_dafny.Seq.update(_303_v147, _module.__default.safeIndex((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))], new BigNumber((_303_v147).length)), _302_v146))).Elements) {
          let _305_v144 = _compr_30;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.update(_303_v147, _module.__default.safeIndex((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))], new BigNumber((_303_v147).length)), _302_v146))).contains(_305_v144)) {
            _coll30.push([_305_v144,(_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))]]);
          }
        }
        return _coll30;
      }());
      let _306_v149;
      _306_v149 = _dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC17(_301_v145),_module.__default.fm2((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))], new BigNumber((_298_v143).length), _116_globalState));
      let _307_v150;
      _307_v150 = _module.D13.create_DC27((((_304_v148).contains(_283_v133)) ? ((_304_v148).get(_283_v133)) : (_306_v149)));
      let _308_v151;
      _308_v151 = _dafny.Map.Empty.slice().updateUnsafe((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))],_307_v150);
      _308_v151 = (_308_v151).update((_dafny.ZERO).minus((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))]), _307_v150);
      let _309_v152;
      _309_v152 = _dafny.MultiSet.fromElements(_111_v1);
      let _310_v153;
      let _nw36 = new _module.C4();
      _nw36.__ctor(_309_v152, !(_110_v0), _dafny.Set.fromElements(_110_v0), ((_110_v0) ? (_111_v1) : ((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))])));
      _310_v153 = _nw36;
      let _311_v154;
      let _nw37 = Array((new BigNumber(11)).toNumber()).fill(false);
      _311_v154 = _nw37;
      let _index29 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length));
      (_311_v154)[_index29] = _110_v0;
      let _index30 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length));
      let _rhs41 = _110_v0;
      let _rhs42 = (((_310_v153).f17) ? (_310_v153) : (_310_v153));
      let _rhs43 = (_310_v153).f17;
      let _rhs44 = (_310_v153).f17;
      let _lhs35 = _116_globalState;
      let _lhs36 = _116_globalState;
      let _lhs37 = _311_v154;
      let _lhs38 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length));
      _lhs35.f7 = _rhs41;
      _310_v153 = _rhs42;
      _lhs36.f7 = _rhs43;
      _lhs37[_lhs38] = _rhs44;
      (_116_globalState).f1 = _module.__default.safeModuloInt(_111_v1, new BigNumber(223));
      let _312_v155;
      let _nw38 = new _module.C5();
      _nw38.__ctor();
      _312_v155 = _nw38;
      let _313_v156;
      _313_v156 = _dafny.Seq.of(_312_v155);
      let _314_v157;
      _314_v157 = _module.D20.create_DC44((_313_v156)[_module.__default.safeIndex(new BigNumber((_198_v77).length), new BigNumber((_313_v156).length))]);
      _314_v157 = _314_v157;
      let _315_v159;
      _315_v159 = _dafny.Map.Empty.slice().updateUnsafe(true,true);
      let _316_v160;
      _316_v160 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_315_v159).length), _111_v1));
      let _317_v161;
      _317_v161 = _dafny.MultiSet.fromElements((_310_v153).fm4(_111_v1, function () {
        let _coll31 = new _dafny.Map();
        for (const _compr_31 of (_316_v160).Elements) {
          let _318_v158 = _compr_31;
          if ((_316_v160).contains(_318_v158)) {
            _coll31.push([_318_v158,_113_v3]);
          }
        }
        return _coll31;
      }(), _111_v1, (_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))], _116_globalState), (_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))]);
      let _319_v163;
      _319_v163 = _module.D22.create_DC51(function () {
  let _coll32 = new _dafny.Set();
  for (const _compr_32 of (_dafny.Map.Empty.slice().updateUnsafe((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))],(_310_v153).f17)).Keys.Elements) {
    let _320_v162 = _compr_32;
    if ((_dafny.Map.Empty.slice().updateUnsafe((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))],(_310_v153).f17)).contains(_320_v162)) {
      _coll32.add(_module.__default.safeModuloInt(_320_v162, new BigNumber(-545)));
    }
  }
  return _coll32;
}());
      let _321_v164;
      _321_v164 = _dafny.Seq.of(_319_v163);
      (_116_globalState).f6 = (_module.__default.safeModuloInt((((_317_v161).contains((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))])) ? ((_317_v161).get((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))])) : ((_dafny.ZERO).minus(_111_v1))), new BigNumber((_dafny.MultiSet.FromArray(_321_v164)).cardinality()))).isLessThan(new BigNumber((_297_v142).length));
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_296_v141).length))) {
        let _322_i17 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_322_i17)) && ((_322_i17).isLessThan(new BigNumber((_296_v141).length))))) {
          (_296_v141)[(_322_i17)] = (_322_i17).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(new BigNumber((function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of (_dafny.MultiSet.fromElements(_283_v133)).Elements) {
              let _323_v165 = _compr_33;
              if ((_dafny.MultiSet.fromElements(_283_v133)).contains(_323_v165)) {
                _coll33.add(_323_v165);
              }
            }
            return _coll33;
          }()).length), _111_v1)).length), new BigNumber((_112_v2).length)));
        }
      }
      let _324_v166;
      _324_v166 = _module.D24.create_DC58((_310_v153).f16);
      let _325_i18;
      _325_i18 = _dafny.ZERO;
      L2: {
        let _pat_let_tv6 = _298_v143;
        let _pat_let_tv7 = _298_v143;
        let _pat_let_tv8 = _111_v1;
        let _pat_let_tv9 = _311_v154;
        let _pat_let_tv10 = _311_v154;
        let _pat_let_tv11 = _298_v143;
        let _pat_let_tv12 = _111_v1;
        let _pat_let_tv13 = _298_v143;
        let _pat_let_tv14 = _296_v141;
        let _pat_let_tv15 = _296_v141;
        let _pat_let_tv16 = _296_v141;
        let _pat_let_tv17 = _296_v141;
        let _pat_let_tv18 = _111_v1;
        let _pat_let_tv19 = _296_v141;
        let _pat_let_tv20 = _296_v141;
        while (function (_source6) {
          if (_source6.is_DC59) {
            let _351___mcc_h5 = (_source6).cf97;
            let _352_cf97 = _351___mcc_h5;
            return !_dafny.areEqual(_dafny.Seq.of(_pat_let_tv6, _pat_let_tv7), _dafny.Seq.of(_dafny.Seq.of(_pat_let_tv8)));
          } else if (_source6.is_DC60) {
            return (_pat_let_tv10)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_pat_let_tv9).length))];
          } else {
            let _353___mcc_h6 = (_source6).cf96;
            let _354_cf96 = _353___mcc_h6;
            return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_pat_let_tv11, _module.__default.safeIndex(_pat_let_tv12, new BigNumber((_pat_let_tv13).length)), (_pat_let_tv15)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_pat_let_tv14).length))]), _dafny.Seq.of((_pat_let_tv17)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_pat_let_tv16).length))], _pat_let_tv18, (_pat_let_tv20)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_pat_let_tv19).length))]));
          }
        }((((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))]) ? (_module.D24.create_DC58((_310_v153).f16)) : (_324_v166)))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_325_i18)) {
              break L2;
            }
            _325_i18 = (_325_i18).plus(_dafny.ONE);
            let _326_v167;
            let _nw39 = new _module.C0();
            _nw39.__ctor(_dafny.Seq.Concat(_112_v2, _112_v2), _module.__default.safeModuloInt((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))], new BigNumber(((_310_v153).f16).cardinality())));
            _326_v167 = _nw39;
            let _327_v168;
            let _nw40 = Array((new BigNumber(5)).toNumber()).fill([]);
            _327_v168 = _nw40;
            let _328_v169;
            let _nw41 = Array((new BigNumber(27)).toNumber());
            _nw41[(_dafny.ZERO).toNumber()] = _198_v77;
            _nw41[(_dafny.ONE).toNumber()] = _198_v77;
            _nw41[(new BigNumber(2)).toNumber()] = _module.__default.fm1(_116_globalState);
            _nw41[(new BigNumber(3)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(4)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(5)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(6)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(7)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(8)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("smkfbwtdu");
            _nw41[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(667)), ((_329_v133) => function (_330_i19) {
              return _329_v133;
            })(_283_v133));
            _nw41[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("mfcvhqiw");
            _nw41[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(849)), ((_331_v133) => function (_332_i20) {
              return _331_v133;
            })(_283_v133));
            _nw41[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("kjwo");
            _nw41[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(996)), function (_333_i21) {
              return new _dafny.CodePoint('y'.codePointAt(0));
            });
            _nw41[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("gnhdyskyq");
            _nw41[(new BigNumber(16)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-3)), ((_334_v133) => function (_335_i22) {
              return _334_v133;
            })(_283_v133));
            _nw41[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("ykantl");
            _nw41[(new BigNumber(19)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(20)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_198_v77, _module.__default.safeIndex((_326_v167).f21, new BigNumber((_198_v77).length)), _283_v133);
            _nw41[(new BigNumber(22)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(780)), ((_336_v133) => function (_337_i23) {
              return _336_v133;
            })(_283_v133));
            _nw41[(new BigNumber(23)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(24)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(25)).toNumber()] = _198_v77;
            _nw41[(new BigNumber(26)).toNumber()] = _198_v77;
            _328_v169 = _nw41;
            let _index31 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_327_v168).length));
            (_327_v168)[_index31] = _328_v169;
            let _index32 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_327_v168).length));
            (_327_v168)[_index32] = _328_v169;
            _198_v77 = _dafny.Seq.Concat(_dafny.Seq.Concat(_198_v77, _198_v77), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bpbqxjl"), _198_v77));
            if ((_310_v153).f17) {
              let _338_v170;
              _338_v170 = _dafny.Seq.of(_317_v161, _317_v161, _317_v161, _317_v161);
              let _339_v171;
              _339_v171 = _dafny.Set.fromElements((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))]);
              let _340_v172;
              let _nw42 = Array((new BigNumber(25)).toNumber());
              _nw42[(_dafny.ZERO).toNumber()] = _317_v161;
              _nw42[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))], (_310_v153).f17);
              _nw42[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements(!((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))]))).Difference(_317_v161);
              _nw42[(new BigNumber(3)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(4)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_112_v2);
              _nw42[(new BigNumber(6)).toNumber()] = (_module.__default.fm37(_116_globalState)).Union(_317_v161);
              _nw42[(new BigNumber(7)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(8)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(9)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements((_326_v167).fm21(true, _111_v1, _116_globalState), (_310_v153).f17, (_310_v153).f17);
              _nw42[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements((_310_v153).f17);
              _nw42[(new BigNumber(12)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(13)).toNumber()] = (_317_v161).Difference(_317_v161);
              _nw42[(new BigNumber(14)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))], (_310_v153).f17, !(_110_v0));
              _nw42[(new BigNumber(16)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(17)).toNumber()] = ((!((_310_v153).f17)) ? (_317_v161) : (_dafny.MultiSet.FromArray(_112_v2)));
              _nw42[(new BigNumber(18)).toNumber()] = (_338_v170)[_module.__default.safeIndex(_111_v1, new BigNumber((_338_v170).length))];
              _nw42[(new BigNumber(19)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(20)).toNumber()] = (_317_v161).Union(_317_v161);
              _nw42[(new BigNumber(21)).toNumber()] = (_317_v161).Difference(_module.__default.fm37(_116_globalState));
              _nw42[(new BigNumber(22)).toNumber()] = (_dafny.MultiSet.fromElements(_110_v0, !((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))]))).Difference(_317_v161);
              _nw42[(new BigNumber(23)).toNumber()] = _317_v161;
              _nw42[(new BigNumber(24)).toNumber()] = (_317_v161).update((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))], _module.__default.abs(new BigNumber((_339_v171).length)));
              _340_v172 = _nw42;
              let _index33 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_340_v172).length));
              (_340_v172)[_index33] = _317_v161;
              let _index34 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_340_v172).length));
              (_340_v172)[_index34] = (_317_v161).Union(_dafny.MultiSet.fromElements((_310_v153).f17));
              let _341_v173;
              _341_v173 = _dafny.Map.Empty.slice().updateUnsafe(_114_v4,!((_310_v153).f17) || ((_310_v153).f17));
              let _rhs45 = (_module.__default.safeModuloInt((_298_v143)[_module.__default.safeIndex((_326_v167).f21, new BigNumber((_298_v143).length))], (_326_v167).f21)).plus(_module.__default.safeModuloInt((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))], (_326_v167).f21));
              let _rhs46 = (_310_v153).f17;
              let _rhs47 = (((_310_v153).f17) ? ((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))]) : (!((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))])));
              let _rhs48 = ((_326_v167).f21).minus(new BigNumber((_198_v77).length));
              let _rhs49 = _341_v173;
              let _lhs39 = _116_globalState;
              let _lhs40 = _116_globalState;
              let _lhs41 = _116_globalState;
              _lhs39.f1 = _rhs45;
              _110_v0 = _rhs46;
              _lhs40.f6 = _rhs47;
              _lhs41.f1 = _rhs48;
              _341_v173 = _rhs49;
              let _342_v174;
              let _nw43 = Array((new BigNumber(29)).toNumber());
              _342_v174 = _nw43;
              let _343_v175;
              let _nw44 = new _module.C8();
              _nw44.__ctor(!((_310_v153).f17), _339_v171, new BigNumber((_198_v77).length));
              _343_v175 = _nw44;
              let _index35 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_342_v174).length));
              (_342_v174)[_index35] = _343_v175;
              let _index36 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_342_v174).length));
              (_342_v174)[_index36] = _343_v175;
              let _344_v176;
              let _nw45 = new _module.C3();
              _nw45.__ctor(_111_v1, new BigNumber(282), (new BigNumber((_198_v77).length)).multipliedBy((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))]), _339_v171);
              _344_v176 = _nw45;
              let _345_v177;
              let _346_v178;
              let _out12;
              let _out13;
              let _outcollector5 = (_343_v175).m1((_344_v176.f23).minus((_343_v175).fm3((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))], (_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))], (_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))], _116_globalState)), _116_globalState);
              _out12 = _outcollector5[0];
              _out13 = _outcollector5[1];
              _345_v177 = _out12;
              _346_v178 = _out13;
            } else {
              (_116_globalState).f7 = !(_110_v0) || ((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))]);
              let _347_v179;
              let _nw46 = Array((new BigNumber(9)).toNumber()).fill([]);
              _347_v179 = _nw46;
              let _348_v180;
              let _nw47 = Array((new BigNumber(6)).toNumber());
              _348_v180 = _nw47;
              let _index37 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_347_v179).length));
              (_347_v179)[_index37] = _348_v180;
              let _349_v181;
              _349_v181 = _dafny.Seq.of(_348_v180);
              let _350_v182;
              _350_v182 = _dafny.Map.Empty.slice().updateUnsafe((_310_v153).f17,_297_v142);
              let _index38 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_347_v179).length));
              let _index39 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length));
              let _rhs50 = (_349_v181)[_module.__default.safeIndex((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))], new BigNumber((_349_v181).length))];
              let _rhs51 = (_dafny.ZERO).minus((new BigNumber((_350_v182).length)).minus(((_326_v167).f21).plus((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))])));
              let _rhs52 = _111_v1;
              let _lhs42 = _347_v179;
              let _lhs43 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_347_v179).length));
              let _lhs44 = _116_globalState;
              let _lhs45 = _296_v141;
              let _lhs46 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length));
              _lhs42[_lhs43] = _rhs50;
              _lhs44.f1 = _rhs51;
              _lhs45[_lhs46] = _rhs52;
              let _index40 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length));
              let _rhs53 = (_312_v155).fm9((_326_v167).f21, _module.__default.fm2(_111_v1, new BigNumber(325), _116_globalState), (_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))], _116_globalState);
              let _rhs54 = _296_v141;
              let _rhs55 = (_326_v167).f21;
              let _rhs56 = (_310_v153).fm4(_module.__default.safeDivisionInt(_111_v1, _111_v1), _dafny.Map.Empty.slice().updateUnsafe((_310_v153).f16,_113_v3), _111_v1, !((_310_v153).f17) || ((_310_v153).f17), _116_globalState);
              let _lhs47 = _116_globalState;
              let _lhs48 = _296_v141;
              let _lhs49 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length));
              let _lhs50 = _116_globalState;
              _lhs47.f1 = _rhs53;
              _296_v141 = _rhs54;
              _lhs48[_lhs49] = _rhs55;
              _lhs50.f6 = _rhs56;
              (_116_globalState).f1 = (_dafny.ZERO).minus((((_310_v153).f17) ? ((((_310_v153).f17) ? ((_326_v167).f21) : ((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))]))) : (_111_v1)));
              let _rhs57 = (_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))];
              let _rhs58 = (((_114_v4).contains(_110_v0)) ? ((_114_v4).get(_110_v0)) : (_111_v1));
              let _rhs59 = (_310_v153).f17;
              let _rhs60 = !((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))]).isEqualTo((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))]);
              let _lhs51 = _116_globalState;
              let _lhs52 = _116_globalState;
              let _lhs53 = _116_globalState;
              _lhs51.f1 = _rhs57;
              _lhs52.f1 = _rhs58;
              _lhs53.f7 = _rhs59;
              _110_v0 = _rhs60;
            }
          }
        }
      }
      let _355_i24;
      _355_i24 = _dafny.ZERO;
      L3: {
        while ((_311_v154)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_311_v154).length))]) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_355_i24)) {
              break L3;
            }
            _355_i24 = (_355_i24).plus(_dafny.ONE);
            let _356_v183;
            let _out14;
            _out14 = (_310_v153).m0(_283_v133, _116_globalState);
            _356_v183 = _out14;
            _312_v155 = _312_v155;
            _296_v141 = _296_v141;
            let _index41 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length));
            (_296_v141)[_index41] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt((_296_v141)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_296_v141).length))], ((true) ? (_111_v1) : (_111_v1))))));
          }
        }
      }
      process.stdout.write(_dafny.toString(_110_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_111_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_112_v2, _dafny.Seq.of(true, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_113_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(754),new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_115_v5)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_115_v5)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_115_v5)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_115_v5)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(754)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_115_v5)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_115_v5)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(754)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_115_v5)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_115_v5)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_116_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_globalState.f5)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_globalState.f5)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_globalState.f5)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_globalState.f5)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(754)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_globalState.f5)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_globalState.f5)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(754)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_globalState.f5)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_116_globalState.f5)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_116_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_116_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_116_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_117_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_198_v77).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_199_v78).dtor_cf40).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_v78).dtor_cf41));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_199_v78).dtor_cf42));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_283_v133));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_284_v134).dtor_cf13, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_285_i13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_296_v141)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_297_v142, _dafny.Seq.of(new BigNumber(630)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_298_v143, _dafny.Seq.of(new BigNumber(630), new BigNumber(-630), _dafny.ONE, new BigNumber(630)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_301_v145).dtor_cf23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_301_v145).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_302_v146).dtor_cf25).dtor_cf23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_302_v146).dtor_cf25).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_303_v147, _dafny.Seq.of(_module.D8.create_DC17(_module.D8.create_DC16(true, new BigNumber(630))), _module.D8.create_DC17(_module.D8.create_DC16(true, new BigNumber(630)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_304_v148).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC17(_module.D8.create_DC16(true, new BigNumber(630))),new BigNumber(630))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v149).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC17(_module.D8.create_DC16(true, new BigNumber(630))),_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_307_v150).dtor_cf39).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC17(_module.D8.create_DC16(true, new BigNumber(630))),new BigNumber(630)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_308_v151).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(630),_module.D13.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC17(_module.D8.create_DC16(true, new BigNumber(630))),new BigNumber(630)))).updateUnsafe(new BigNumber(-630),_module.D13.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC17(_module.D8.create_DC16(true, new BigNumber(630))),new BigNumber(630)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_v152).equals(_dafny.MultiSet.fromElements(new BigNumber(630)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_310_v153).f16).equals(_dafny.MultiSet.fromElements(new BigNumber(630)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v153).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_311_v154)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_313_v156).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_315_v159).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_316_v160).equals(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_dafny.ONE, new BigNumber(630))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_317_v161).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_319_v163).dtor_cf86).equals(_dafny.Set.fromElements(new BigNumber(85)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_321_v164, _dafny.Seq.of(_module.D22.create_DC51(_dafny.Set.fromElements(new BigNumber(85)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_324_v166).dtor_cf96).equals(_dafny.MultiSet.fromElements(new BigNumber(630)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_325_i18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_355_i24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
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
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, _dafny.ZERO, false);
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
    static create_DC2(cf4) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + this.cf4.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4);
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
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf6, cf7, cf8) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC5() {
      let $dt = new D2(1);
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D2(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5";
      } else if (this.$tag === 2) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC6(cf9) {
      let $dt = new D3(0);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf9 === other.cf9;
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf10) {
      let $dt = new D4(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10);
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
    static create_DC8(cf11) {
      let $dt = new D5(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC8" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf13) {
      let $dt = new D6(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC11(cf14, cf15, cf16) {
      let $dt = new D6(1);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf12) {
      let $dt = new D6(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC10" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC11" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC9" + "(" + _dafny.toString(this.cf12) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC10(_dafny.Seq.of());
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
    static create_DC13(cf18, cf19, cf20) {
      let $dt = new D7(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC12(cf17) {
      let $dt = new D7(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC14(cf21) {
      let $dt = new D7(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC13" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC12" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC13(_dafny.Set.Empty, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC16(cf23, cf24) {
      let $dt = new D8(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf22) {
      let $dt = new D8(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC17(cf25) {
      let $dt = new D8(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC16" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC15" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC16(false, _dafny.ZERO);
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
    static create_DC19() {
      let $dt = new D9(0);
      return $dt;
    }
    static create_DC18(cf26) {
      let $dt = new D9(1);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC19";
      } else if (this.$tag === 1) {
        return "D9.DC18" + "(" + _dafny.toString(this.cf26) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC19();
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
    static create_DC21(cf28, cf29) {
      let $dt = new D10(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC20(cf27) {
      let $dt = new D10(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC20" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28 && this.cf29 === other.cf29;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC21(false, false);
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
    static create_DC23(cf31, cf32, cf33) {
      let $dt = new D11(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC24(cf34, cf35, cf36) {
      let $dt = new D11(1);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC22(cf30) {
      let $dt = new D11(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC23" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC22" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC23(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC26(cf38) {
      let $dt = new D12(0);
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC25(cf37) {
      let $dt = new D12(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC26" + "(" + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC25" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf37 === other.cf37;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC26(_dafny.Map.Empty);
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
    static create_DC28(cf40, cf41, cf42) {
      let $dt = new D13(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC27(cf39) {
      let $dt = new D13(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC29(cf43) {
      let $dt = new D13(2);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC28" + "(" + this.cf40.toVerbatimString(true) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC27" + "(" + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41) && this.cf42 === other.cf42;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC28(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, false);
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
    static create_DC31(cf45, cf46, cf47) {
      let $dt = new D14(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC30(cf44) {
      let $dt = new D14(1);
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC32(cf48) {
      let $dt = new D14(2);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC32() { return this.$tag === 2; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC30" + "(" + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC32" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC31(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC34() {
      let $dt = new D15(0);
      return $dt;
    }
    static create_DC35(cf50, cf51, cf52, cf53, cf54) {
      let $dt = new D15(1);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC36(cf55, cf56, cf57, cf58, cf59) {
      let $dt = new D15(2);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC33(cf49) {
      let $dt = new D15(3);
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC37(cf60) {
      let $dt = new D15(4);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get is_DC37() { return this.$tag === 4; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC34";
      } else if (this.$tag === 1) {
        return "D15.DC35" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC36" + "(" + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC33" + "(" + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 4) {
        return "D15.DC37" + "(" + _dafny.toString(this.cf60) + ")";
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
        return other.$tag === 1 && this.cf50 === other.cf50 && this.cf51 === other.cf51 && this.cf52 === other.cf52 && _dafny.areEqual(this.cf53, other.cf53) && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55) && this.cf56 === other.cf56 && this.cf57 === other.cf57 && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC34();
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
    static create_DC39(cf62, cf63, cf64, cf65) {
      let $dt = new D16(0);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC38(cf61) {
      let $dt = new D16(1);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC38" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf62, other.cf62) && this.cf63 === other.cf63 && _dafny.areEqual(this.cf64, other.cf64) && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC39(_dafny.ZERO, false, _dafny.Map.Empty, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC40(cf66) {
      let $dt = new D17(0);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC40" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf66 === other.cf66;
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
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf67) {
      let $dt = new D18(0);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC41" + "(" + _dafny.toString(this.cf67) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC43(cf69, cf70, cf71, cf72, cf73) {
      let $dt = new D19(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC42(cf68) {
      let $dt = new D19(1);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC43" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ", " + this.cf73.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC42" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf69 === other.cf69 && _dafny.areEqual(this.cf70, other.cf70) && this.cf71 === other.cf71 && _dafny.areEqual(this.cf72, other.cf72) && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf68 === other.cf68;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC43(false, new _dafny.CodePoint('D'.codePointAt(0)), false, _module.D6.Default(), _dafny.Seq.UnicodeFromString(""));
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
    static create_DC45(cf75, cf76) {
      let $dt = new D20(0);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC46() {
      let $dt = new D20(1);
      return $dt;
    }
    static create_DC44(cf74) {
      let $dt = new D20(2);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC45" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC46";
      } else if (this.$tag === 2) {
        return "D20.DC44" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf75 === other.cf75 && this.cf76 === other.cf76;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf74 === other.cf74;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC45(false, false);
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
    static create_DC48(cf78) {
      let $dt = new D21(0);
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC49(cf79, cf80, cf81, cf82) {
      let $dt = new D21(1);
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC50(cf83, cf84, cf85) {
      let $dt = new D21(2);
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC47(cf77) {
      let $dt = new D21(3);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get is_DC50() { return this.$tag === 2; }
    get is_DC47() { return this.$tag === 3; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC48" + "(" + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC49" + "(" + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC50" + "(" + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC47" + "(" + _dafny.toString(this.cf77) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf78 === other.cf78;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf79 === other.cf79 && this.cf80 === other.cf80 && _dafny.areEqual(this.cf81, other.cf81) && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf83 === other.cf83 && _dafny.areEqual(this.cf84, other.cf84) && this.cf85 === other.cf85;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf77, other.cf77);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC48(false);
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
    static create_DC52() {
      let $dt = new D22(0);
      return $dt;
    }
    static create_DC53(cf87, cf88, cf89, cf90) {
      let $dt = new D22(1);
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC51(cf86) {
      let $dt = new D22(2);
      $dt.cf86 = cf86;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC53() { return this.$tag === 1; }
    get is_DC51() { return this.$tag === 2; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf86() { return this.cf86; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC52";
      } else if (this.$tag === 1) {
        return "D22.DC53" + "(" + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + this.cf89.toVerbatimString(true) + ", " + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC51" + "(" + _dafny.toString(this.cf86) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf87, other.cf87) && this.cf88 === other.cf88 && _dafny.areEqual(this.cf89, other.cf89) && _dafny.areEqual(this.cf90, other.cf90);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf86, other.cf86);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC52();
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
    static create_DC55(cf92, cf93, cf94) {
      let $dt = new D23(0);
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC56() {
      let $dt = new D23(1);
      return $dt;
    }
    static create_DC54(cf91) {
      let $dt = new D23(2);
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC57(cf95) {
      let $dt = new D23(3);
      $dt.cf95 = cf95;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get is_DC54() { return this.$tag === 2; }
    get is_DC57() { return this.$tag === 3; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf95() { return this.cf95; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC55" + "(" + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC56";
      } else if (this.$tag === 2) {
        return "D23.DC54" + "(" + _dafny.toString(this.cf91) + ")";
      } else if (this.$tag === 3) {
        return "D23.DC57" + "(" + _dafny.toString(this.cf95) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf92, other.cf92) && this.cf93 === other.cf93 && _dafny.areEqual(this.cf94, other.cf94);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf91 === other.cf91;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf95, other.cf95);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC55(_dafny.Seq.of(), [], _dafny.ZERO);
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
    static create_DC59(cf97) {
      let $dt = new D24(0);
      $dt.cf97 = cf97;
      return $dt;
    }
    static create_DC60() {
      let $dt = new D24(1);
      return $dt;
    }
    static create_DC58(cf96) {
      let $dt = new D24(2);
      $dt.cf96 = cf96;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get is_DC60() { return this.$tag === 1; }
    get is_DC58() { return this.$tag === 2; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf96() { return this.cf96; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC59" + "(" + _dafny.toString(this.cf97) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC60";
      } else if (this.$tag === 2) {
        return "D24.DC58" + "(" + _dafny.toString(this.cf96) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf97, other.cf97);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf96, other.cf96);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC59(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC62(cf99) {
      let $dt = new D25(0);
      $dt.cf99 = cf99;
      return $dt;
    }
    static create_DC61(cf98) {
      let $dt = new D25(1);
      $dt.cf98 = cf98;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC61() { return this.$tag === 1; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf98() { return this.cf98; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC62" + "(" + _dafny.toString(this.cf99) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC61" + "(" + _dafny.toString(this.cf98) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf99 === other.cf99;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf98 === other.cf98;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC62(false);
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
    static create_DC63(cf100) {
      let $dt = new D26(0);
      $dt.cf100 = cf100;
      return $dt;
    }
    get is_DC63() { return this.$tag === 0; }
    get dtor_cf100() { return this.cf100; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC63" + "(" + _dafny.toString(this.cf100) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf100, other.cf100);
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
          return D26.Default();
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
      this.f1 = _dafny.ZERO;
      this.f5 = [];
      this.f6 = false;
      this.f7 = false;
      this._f0 = false;
      this._f2 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f4 = false;
      this._f8 = false;
      this._f9 = _dafny.ZERO;
      this._f10 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this).f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f20 = _dafny.Seq.of();
      this._f21 = _dafny.ZERO;
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
    fm21(p0, p1, globalState) {
      let _this = this;
      if (!(_dafny.MultiSet.fromElements(false, true)).equals(_dafny.MultiSet.fromElements(true))) {
        return (!(!(true))) && (false);
      } else {
        return false;
      }
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

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f18 = false;
      this._f19 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f18, f19) {
      let _this = this;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm13(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f18;
    };
    fm14(p0, globalState) {
      let _this = this;
      return (_this).f18;
    };
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let _357_v0;
      _357_v0 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18, (_this).f18);
      let _358_v1;
      _358_v1 = _dafny.Set.fromElements(_module.__default.fm15(p2, globalState));
      let _359_v3;
      _359_v3 = _dafny.Set.fromElements(_358_v1, _358_v1, function () {
        let _coll34 = new _dafny.Set();
        for (const _compr_34 of (_module.__default.fm16(p2, new BigNumber((p3).length), globalState)).Keys.Elements) {
          let _360_v2 = _compr_34;
          if ((_module.__default.fm16(p2, new BigNumber((p3).length), globalState)).contains(_360_v2)) {
            _coll34.add(_360_v2);
          }
        }
        return _coll34;
      }());
      let _361_v4;
      _361_v4 = _dafny.Seq.of(p1);
      let _362_v5;
      _362_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,p1);
      let _363_v6;
      _363_v6 = _dafny.Map.Empty.slice().updateUnsafe(_362_v5,(_this).f19);
      let _364_v8;
      _364_v8 = _module.D2.create_DC5();
      let _365_v9;
      _365_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus(p1));
      let _366_v10;
      _366_v10 = _dafny.Seq.of(_dafny.Seq.update(_module.__default.fm17(_363_v6, p1, _module.D2.create_DC5(), (_this).f18, globalState), _module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_module.__default.fm17(_363_v6, p1, _module.D2.create_DC5(), (_this).f18, globalState)).length)), new BigNumber(((_365_v9).update(new BigNumber(-171), p1)).length)), _361_v4, _module.__default.fm17(_363_v6, new BigNumber(415), _364_v8, (_this).f19, globalState));
      let _367_v11;
      _367_v11 = _dafny.Seq.of(_361_v4, _module.__default.fm17(_363_v6, new BigNumber((function () {
        let _coll35 = new _dafny.Map();
        for (const _compr_35 of _dafny.IntegerRange(new BigNumber(-770), new BigNumber(-776))) {
          let _368_v7 = _compr_35;
          if (((new BigNumber(-770)).isLessThanOrEqualTo(_368_v7)) && ((_368_v7).isLessThan(new BigNumber(-776)))) {
            _coll35.push([(_368_v7).multipliedBy(p2),(_this).f18]);
          }
        }
        return _coll35;
      }()).length), _364_v8, (_this).f18, globalState), _361_v4, (_366_v10)[_module.__default.safeIndex(p2, new BigNumber((_366_v10).length))]);
      let _369_v12;
      let _out15;
      _out15 = (_this).m7(_357_v0, _359_v3, (_367_v11)[_module.__default.safeIndex(_module.__default.fm2(p2, new BigNumber(450), globalState), new BigNumber((_367_v11).length))], globalState);
      _369_v12 = _out15;
      let _370_v13;
      _370_v13 = _dafny.Seq.of((_this).f19, (_this).f18, (p2).isLessThan(p2), (_this).f18);
      if ((_370_v13)[_module.__default.safeIndex(p2, new BigNumber((_370_v13).length))]) {
        let _371_v14;
        let _nw48 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
        _371_v14 = _nw48;
        _371_v14 = _371_v14;
        let _372_v15;
        let _nw49 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
        _372_v15 = _nw49;
        let _373_v16;
        _373_v16 = new _dafny.CodePoint('k'.codePointAt(0));
        let _374_v17;
        let _nw50 = Array((new BigNumber(14)).toNumber()).fill(false);
        _374_v17 = _nw50;
        let _375_v18;
        _375_v18 = _dafny.Set.fromElements(_374_v17);
        let _376_v19;
        let _nw51 = Array((new BigNumber(21)).toNumber());
        _nw51[(_dafny.ZERO).toNumber()] = _module.__default.fm1(globalState);
        _nw51[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(149)), function (_377_i0) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        });
        _nw51[(new BigNumber(2)).toNumber()] = p0;
        _nw51[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(p0, _module.__default.safeIndex(p2, new BigNumber((p0).length)), _373_v16);
        _nw51[(new BigNumber(4)).toNumber()] = p0;
        _nw51[(new BigNumber(5)).toNumber()] = p0;
        _nw51[(new BigNumber(6)).toNumber()] = p0;
        _nw51[(new BigNumber(7)).toNumber()] = p0;
        _nw51[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("vpilsnhk");
        _nw51[(new BigNumber(9)).toNumber()] = p0;
        _nw51[(new BigNumber(10)).toNumber()] = p0;
        _nw51[(new BigNumber(11)).toNumber()] = p0;
        _nw51[(new BigNumber(12)).toNumber()] = p0;
        _nw51[(new BigNumber(13)).toNumber()] = p0;
        _nw51[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("gn");
        _nw51[(new BigNumber(15)).toNumber()] = p0;
        _nw51[(new BigNumber(16)).toNumber()] = p0;
        _nw51[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(p0, _module.__default.safeIndex((_361_v4)[_module.__default.safeIndex(new BigNumber((_375_v18).length), new BigNumber((_361_v4).length))], new BigNumber((p0).length)), new _dafny.CodePoint('a'.codePointAt(0)));
        _nw51[(new BigNumber(18)).toNumber()] = p0;
        _nw51[(new BigNumber(19)).toNumber()] = p0;
        _nw51[(new BigNumber(20)).toNumber()] = p0;
        _376_v19 = _nw51;
        let _378_v20;
        _378_v20 = _dafny.Map.Empty.slice().updateUnsafe(_372_v15,_376_v19);
        _378_v20 = (_378_v20).update(_372_v15, (((_370_v13)[_module.__default.safeIndex(new BigNumber((_365_v9).length), new BigNumber((_370_v13).length))]) ? (_376_v19) : (_376_v19)));
        let _379_v21;
        let _nw52 = Array((new BigNumber(13)).toNumber()).fill([]);
        _379_v21 = _nw52;
        let _380_v22;
        _380_v22 = _dafny.Set.fromElements(p1);
        let _381_v23;
        let _nw53 = Array((new BigNumber(3)).toNumber());
        _nw53[(_dafny.ZERO).toNumber()] = _380_v22;
        _nw53[(_dafny.ONE).toNumber()] = _380_v22;
        _nw53[(new BigNumber(2)).toNumber()] = _module.__default.fm18(new BigNumber((_module.__default.fm19(true, (_this).f18, globalState)).length), globalState);
        _381_v23 = _nw53;
        let _index42 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_379_v21).length));
        (_379_v21)[_index42] = _381_v23;
        let _index43 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_379_v21).length));
        let _nw54 = Array((new BigNumber(3)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = _380_v22;
        _nw54[(_dafny.ONE).toNumber()] = _380_v22;
        _nw54[(new BigNumber(2)).toNumber()] = _380_v22;
        (_379_v21)[_index43] = _nw54;
        let _382_v24;
        _382_v24 = _dafny.Set.fromElements((_this).f19);
        let _383_v25;
        _383_v25 = _dafny.Map.Empty.slice().updateUnsafe(_382_v24,(_this).f19);
        (globalState).f1 = new BigNumber((_383_v25).length);
        (globalState).f7 = (_this).f18;
      } else {
        let _384_v26;
        let _nw55 = Array((new BigNumber(13)).toNumber()).fill([]);
        _384_v26 = _nw55;
        let _385_v27;
        let _init4 = function (_386_i1) {
          return false;
        };
        let _nw56 = Array((new BigNumber(8)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw56.length); _i0_4++) {
          _nw56[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _385_v27 = _nw56;
        let _index44 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_384_v26).length));
        (_384_v26)[_index44] = _385_v27;
        let _index45 = _module.__default.safeIndex(new BigNumber(759), new BigNumber((_384_v26).length));
        (_384_v26)[_index45] = _385_v27;
        (globalState).f6 = _module.__default.fm0(globalState);
        let _387_v28;
        _387_v28 = new _dafny.CodePoint('f'.codePointAt(0));
        let _388_v29;
        _388_v29 = _dafny.Set.fromElements(p2);
        let _389_v30;
        let _nw57 = Array((new BigNumber(5)).toNumber());
        _nw57[(_dafny.ZERO).toNumber()] = _387_v28;
        _nw57[(_dafny.ONE).toNumber()] = _387_v28;
        _nw57[(new BigNumber(2)).toNumber()] = _387_v28;
        _nw57[(new BigNumber(3)).toNumber()] = _387_v28;
        _nw57[(new BigNumber(4)).toNumber()] = _387_v28;
        _389_v30 = _nw57;
        let _index46 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_389_v30).length));
        (_389_v30)[_index46] = _387_v28;
        let _index47 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v30).length));
        (_389_v30)[_index47] = _387_v28;
        let _390_v31;
        _390_v31 = _dafny.MultiSet.fromElements(_module.D2.create_DC4(new BigNumber((p0).length), new BigNumber((p0).length), p2));
        let _391_v32;
        _391_v32 = _dafny.Map.Empty.slice().updateUnsafe(_390_v31,_388_v29);
        let _index48 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_389_v30).length));
        let _index49 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v30).length));
        let _rhs61 = _387_v28;
        let _rhs62 = (((_391_v32).contains(_dafny.MultiSet.fromElements(_369_v12, _369_v12, _369_v12))) ? ((_391_v32).get(_dafny.MultiSet.fromElements(_369_v12, _369_v12, _369_v12))) : (_388_v29));
        let _rhs63 = _module.__default.fm20(_370_v13, (p1).minus(new BigNumber(936)), (_this).f18, globalState);
        let _rhs64 = _387_v28;
        let _lhs54 = _389_v30;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_389_v30).length));
        let _lhs56 = _389_v30;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v30).length));
        _387_v28 = _rhs61;
        _388_v29 = _rhs62;
        _lhs54[_lhs55] = _rhs63;
        _lhs56[_lhs57] = _rhs64;
        _385_v27 = _385_v27;
        let _392_v33;
        _392_v33 = _370_v13;
        let _393_v34;
        let _nw58 = new _module.C0();
        _nw58.__ctor((_392_v33), (_module.__default.fm2(p1, p2, globalState)).multipliedBy(new BigNumber(353)));
        _393_v34 = _nw58;
      }
      (globalState).f6 = (_this).f19;
      let _394_v35;
      _394_v35 = _module.D0.create_DC0(p1);
      let _source7 = _394_v35;
      if (_source7.is_DC1) {
        let _395___mcc_h0 = (_source7).cf1;
        let _396___mcc_h1 = (_source7).cf2;
        let _397___mcc_h2 = (_source7).cf3;
        let _398_cf3 = _397___mcc_h2;
        let _399_cf2 = _396___mcc_h1;
        let _400_cf1 = _395___mcc_h0;
        (globalState).f7 = (_370_v13)[_module.__default.safeIndex(new BigNumber(-413), new BigNumber((_370_v13).length))];
        let _401_v36;
        _401_v36 = _dafny.MultiSet.fromElements(_399_cf2, (_dafny.ZERO).minus(p1));
        _398_cf3 = ((((_401_v36).contains((_dafny.ZERO).minus(new BigNumber((p3).length)))) ? ((_401_v36).get((_dafny.ZERO).minus(new BigNumber((p3).length)))) : (p2))).isLessThan(_399_cf2);
        let _402_v37;
        let _nw59 = new _module.C0();
        _nw59.__ctor(_dafny.Seq.Concat(_370_v13, _370_v13), _module.__default.safeDivisionInt(new BigNumber(906), _399_cf2));
        _402_v37 = _nw59;
        let _403_v38;
        _403_v38 = _dafny.Map.Empty.slice().updateUnsafe(_364_v8,new BigNumber(((_module.__default.fm22(p0, true, globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f19))).length));
        let _404_v39;
        _404_v39 = _dafny.Seq.UnicodeFromString("jr");
        let _405_v40;
        let _init5 = ((_406_v37) => function (_407_i2) {
          return (_407_i2).minus((_406_v37).f21);
        })(_402_v37);
        let _nw60 = Array((new BigNumber(5)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw60.length); _i0_5++) {
          _nw60[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _405_v40 = _nw60;
        let _index50 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_405_v40).length));
        (_405_v40)[_index50] = (p2).plus(_399_cf2);
        let _408_v41;
        _408_v41 = new _dafny.CodePoint('v'.codePointAt(0));
        let _index51 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_405_v40).length));
        let _rhs65 = _403_v38;
        let _rhs66 = _dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm1(globalState), _404_v39), _module.__default.safeIndex((_dafny.ZERO).minus((((_this).f18) ? (_399_cf2) : ((_dafny.ZERO).minus((_369_v12).dtor_cf8)))), new BigNumber((_dafny.Seq.Concat(_module.__default.fm1(globalState), _404_v39)).length)), _408_v41);
        let _rhs67 = _369_v12;
        let _rhs68 = p1;
        let _rhs69 = p1;
        let _lhs58 = _405_v40;
        let _lhs59 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_405_v40).length));
        _403_v38 = _rhs65;
        _404_v39 = _rhs66;
        _369_v12 = _rhs67;
        _lhs58[_lhs59] = _rhs68;
        r0 = _rhs69;
      } else {
        let _409___mcc_h3 = (_source7).cf0;
        let _410_cf0 = _409___mcc_h3;
        let _411_v42;
        let _init6 = ((_412_v0, _413_p1) => function (_414_i3) {
          return (_414_i3).minus((((_412_v0).contains((_this).f19)) ? ((_412_v0).get((_this).f19)) : (_413_p1)));
        })(_357_v0, p1);
        let _nw61 = Array((new BigNumber(29)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw61.length); _i0_6++) {
          _nw61[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _411_v42 = _nw61;
        let _415_v43;
        _415_v43 = _dafny.MultiSet.fromElements(_411_v42);
        let _416_v44;
        _416_v44 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), function (_417_i4) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }));
        let _418_v45;
        _418_v45 = _dafny.Map.Empty.slice().updateUnsafe((_415_v43).Intersect(_415_v43),_416_v44);
        _418_v45 = (_418_v45).update(_415_v43, _416_v44);
        (globalState).f7 = (((_361_v4)[_module.__default.safeIndex(new BigNumber((_362_v5).length), new BigNumber((_361_v4).length))]).plus(_410_cf0)).isLessThanOrEqualTo(_410_cf0);
        let _419_v46;
        _419_v46 = _dafny.MultiSet.fromElements(_361_v4);
        let _420_v47;
        let _nw62 = Array((new BigNumber(17)).toNumber()).fill(false);
        _420_v47 = _nw62;
        let _421_v48;
        _421_v48 = _dafny.MultiSet.fromElements(_420_v47, _420_v47, _420_v47);
        let _422_v49;
        _422_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f18);
        let _423_v50;
        let _nw63 = Array((new BigNumber(24)).toNumber());
        _nw63[(_dafny.ZERO).toNumber()] = (_this).f19;
        _nw63[(_dafny.ONE).toNumber()] = (_419_v46).IsDisjointFrom(_419_v46);
        _nw63[(new BigNumber(2)).toNumber()] = (_this).f18;
        _nw63[(new BigNumber(3)).toNumber()] = !((_this).f18) || ((_this).f19);
        _nw63[(new BigNumber(4)).toNumber()] = !((_this).f18);
        _nw63[(new BigNumber(5)).toNumber()] = (_this).f18;
        _nw63[(new BigNumber(6)).toNumber()] = (_this).f19;
        _nw63[(new BigNumber(7)).toNumber()] = (_this).f18;
        _nw63[(new BigNumber(8)).toNumber()] = (_this).f18;
        _nw63[(new BigNumber(9)).toNumber()] = (p3).equals(p3);
        _nw63[(new BigNumber(10)).toNumber()] = (_this).f18;
        _nw63[(new BigNumber(11)).toNumber()] = (_this).f19;
        _nw63[(new BigNumber(12)).toNumber()] = false;
        _nw63[(new BigNumber(13)).toNumber()] = (_this).f19;
        _nw63[(new BigNumber(14)).toNumber()] = (_421_v48).IsProperSubsetOf(_dafny.MultiSet.fromElements(_420_v47));
        _nw63[(new BigNumber(15)).toNumber()] = (_this).f18;
        _nw63[(new BigNumber(16)).toNumber()] = (_this).f19;
        _nw63[(new BigNumber(17)).toNumber()] = _dafny.Seq.IsPrefixOf(_module.__default.fm1(globalState), _dafny.Seq.UnicodeFromString("kppg"));
        _nw63[(new BigNumber(18)).toNumber()] = (_this).f18;
        _nw63[(new BigNumber(19)).toNumber()] = (_this).f19;
        _nw63[(new BigNumber(20)).toNumber()] = (_this).f18;
        _nw63[(new BigNumber(21)).toNumber()] = (_this).f18;
        _nw63[(new BigNumber(22)).toNumber()] = !((p1).isLessThanOrEqualTo(new BigNumber((_422_v49).length)));
        _nw63[(new BigNumber(23)).toNumber()] = (_this).f18;
        _423_v50 = _nw63;
        let _index52 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_420_v47).length));
        (_420_v47)[_index52] = (_this).f19;
        let _424_v51;
        _424_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_357_v0);
        let _425_v52;
        _425_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm18(p2, globalState)).length),_dafny.MultiSet.FromArray(_370_v13));
        let _426_v53;
        _426_v53 = (((_424_v51).contains((_this).f19)) ? ((_424_v51).get((_this).f19)) : (_357_v0));
        let _427_v54;
        let _nw64 = Array((new BigNumber(28)).toNumber());
        _nw64[(_dafny.ZERO).toNumber()] = _357_v0;
        _nw64[(_dafny.ONE).toNumber()] = _357_v0;
        _nw64[(new BigNumber(2)).toNumber()] = (((_this).f19) ? (_357_v0) : (_dafny.MultiSet.FromArray(_370_v13)));
        _nw64[(new BigNumber(3)).toNumber()] = (_357_v0).Difference(_dafny.MultiSet.fromElements((_this).f18));
        _nw64[(new BigNumber(4)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(5)).toNumber()] = (_357_v0).Union(_357_v0);
        _nw64[(new BigNumber(6)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(7)).toNumber()] = (((_424_v51).contains(!((_this).f18))) ? ((_424_v51).get(!((_this).f18))) : (_357_v0));
        _nw64[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements((_this).f19);
        _nw64[(new BigNumber(9)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements((_this).f19, !((_this).f19), (_this).f18, (_this).f19);
        _nw64[(new BigNumber(11)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(12)).toNumber()] = ((false) ? (_dafny.MultiSet.fromElements((_this).f19, false)) : (_357_v0));
        _nw64[(new BigNumber(13)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(14)).toNumber()] = (((_425_v52).contains(_410_cf0)) ? ((_425_v52).get(_410_cf0)) : (_dafny.MultiSet.FromArray(_370_v13)));
        _nw64[(new BigNumber(15)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(16)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(17)).toNumber()] = (_357_v0).Difference((_426_v53));
        _nw64[(new BigNumber(18)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(19)).toNumber()] = (_357_v0).Union((_dafny.MultiSet.fromElements((_this).f18)).update((_this).f19, _module.__default.abs(new BigNumber(727))));
        _nw64[(new BigNumber(20)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(21)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(22)).toNumber()] = (_dafny.MultiSet.fromElements((_this).f19, (_this).f19)).Intersect((_357_v0).update(!((_this).f18), _module.__default.abs(p2)));
        _nw64[(new BigNumber(23)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(24)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(25)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(26)).toNumber()] = _357_v0;
        _nw64[(new BigNumber(27)).toNumber()] = _357_v0;
        _427_v54 = _nw64;
        let _index53 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_427_v54).length));
        (_427_v54)[_index53] = (_357_v0).update((_this).f18, _module.__default.abs(p2));
        let _428_v55;
        _428_v55 = _module.D0.create_DC1(!((_this).f18), p2, (_this).f19);
        let _429_v56;
        _429_v56 = new _dafny.CodePoint('t'.codePointAt(0));
        let _430_v57;
        let _nw65 = Array((new BigNumber(11)).toNumber());
        _nw65[(_dafny.ZERO).toNumber()] = (((_428_v55).dtor_cf1) ? (_429_v56) : (_module.__default.fm20(_370_v13, _410_cf0, (_this).f19, globalState)));
        _nw65[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
        _nw65[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _nw65[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
        _nw65[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
        _nw65[(new BigNumber(5)).toNumber()] = _429_v56;
        _nw65[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
        _nw65[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
        _nw65[(new BigNumber(8)).toNumber()] = _429_v56;
        _nw65[(new BigNumber(9)).toNumber()] = _429_v56;
        _nw65[(new BigNumber(10)).toNumber()] = _429_v56;
        _430_v57 = _nw65;
        let _index54 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_430_v57).length));
        (_430_v57)[_index54] = _429_v56;
        let _431_v58;
        _431_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,p0);
        let _index55 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_420_v47).length));
        let _index56 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_427_v54).length));
        let _index57 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_430_v57).length));
        let _rhs70 = _410_cf0;
        let _rhs71 = !(((_410_cf0).plus(_410_cf0)).isLessThanOrEqualTo(new BigNumber(705)));
        let _rhs72 = !_dafny.areEqual(_dafny.Seq.Concat(p0, p0), (((_431_v58).contains((_this).f18)) ? ((_431_v58).get((_this).f18)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(91)), ((_432_v56) => function (_433_i5) {
          return _432_v56;
        })(_429_v56)))));
        let _rhs73 = _357_v0;
        let _rhs74 = _429_v56;
        let _lhs60 = globalState;
        let _lhs61 = _420_v47;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_420_v47).length));
        let _lhs63 = globalState;
        let _lhs64 = _427_v54;
        let _lhs65 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_427_v54).length));
        let _lhs66 = _430_v57;
        let _lhs67 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_430_v57).length));
        _lhs60.f1 = _rhs70;
        _lhs61[_lhs62] = _rhs71;
        _lhs63.f6 = _rhs72;
        _lhs64[_lhs65] = _rhs73;
        _lhs66[_lhs67] = _rhs74;
        (globalState).f6 = (_420_v47)[_module.__default.safeIndex(new BigNumber(839), new BigNumber((_420_v47).length))];
      }
      (globalState).f1 = p2;
      let _434_v59;
      _434_v59 = _module.D0.create_DC1(((_this).f18) === ((_this).f19), p1, !((_this).f18) || ((_this).f18));
      let _pat_let_tv21 = p1;
      _434_v59 = function (_pat_let7_0) {
        return function (_435_dt__update__tmp_h0) {
          return function (_pat_let8_0) {
            return function (_436_dt__update_hcf1_h0) {
              return function (_pat_let9_0) {
                return function (_437_dt__update_hcf2_h0) {
                  return _module.D0.create_DC1(_436_dt__update_hcf1_h0, _437_dt__update_hcf2_h0, (_435_dt__update__tmp_h0).dtor_cf3);
                }(_pat_let9_0);
              }(_module.__default.safeModuloInt(new BigNumber(473), _pat_let_tv21));
            }(_pat_let8_0);
          }((_this).f18);
        }(_pat_let7_0);
      }(_434_v59);
      r0 = p1;
      let _438_v60;
      _438_v60 = _module.D6.create_DC9(p3);
      let _pat_let_tv22 = p3;
      let _439_v61;
      let _nw66 = Array((new BigNumber(10)).toNumber());
      _nw66[(_dafny.ZERO).toNumber()] = p3;
      _nw66[(_dafny.ONE).toNumber()] = (_438_v60).dtor_cf12;
      _nw66[(new BigNumber(2)).toNumber()] = (p3).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f19));
      _nw66[(new BigNumber(3)).toNumber()] = p3;
      _nw66[(new BigNumber(4)).toNumber()] = p3;
      _nw66[(new BigNumber(5)).toNumber()] = (p3).Merge(p3);
      _nw66[(new BigNumber(6)).toNumber()] = p3;
      _nw66[(new BigNumber(7)).toNumber()] = ((function (_pat_let10_0) {
        return function (_440_dt__update__tmp_h1) {
          return function (_pat_let11_0) {
            return function (_441_dt__update_hcf12_h0) {
              return _module.D6.create_DC9(_441_dt__update_hcf12_h0);
            }(_pat_let11_0);
          }(_pat_let_tv22);
        }(_pat_let10_0);
      }(_438_v60)).dtor_cf12).Merge((p3).update(p2, (_this).f19));
      _nw66[(new BigNumber(8)).toNumber()] = (p3).Merge(p3);
      _nw66[(new BigNumber(9)).toNumber()] = p3;
      _439_v61 = _nw66;
      r1 = _439_v61;
      return [r0, r1];
    }
    m7(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _module.D2.Default();
      let _442_v0;
      _442_v0 = _dafny.Seq.UnicodeFromString("rpix");
      let _443_i0;
      _443_i0 = _dafny.ZERO;
      L4: {
        while (_dafny.Seq.IsPrefixOf(_module.__default.fm1(globalState), _dafny.Seq.Concat(_442_v0, _442_v0))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_443_i0)) {
              break L4;
            }
            _443_i0 = (_443_i0).plus(_dafny.ONE);
            let _444_v1;
            _444_v1 = new BigNumber(239);
            (globalState).f6 = ((((_this).f19) === ((_this).f19)) ? ((_this).f18) : ((_444_v1).isLessThan(_444_v1)));
            let _445_v2;
            _445_v2 = new _dafny.CodePoint('s'.codePointAt(0));
            _445_v2 = _445_v2;
            let _446_v3;
            _446_v3 = _dafny.Map.Empty.slice().updateUnsafe(false,_445_v2);
            _445_v2 = (((_446_v3).contains((_this).f19)) ? ((_446_v3).get((_this).f19)) : (_445_v2));
            let _447_v4;
            _447_v4 = _module.D0.create_DC1((_this).f19, _444_v1, (p0).contains((_this).f18));
            let _448_v5;
            let _init7 = ((_449_v1) => function (_450_i1) {
              return _module.__default.safeDivisionInt(_450_i1, _449_v1);
            })(_444_v1);
            let _nw67 = Array((new BigNumber(24)).toNumber());
            for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw67.length); _i0_7++) {
              _nw67[_i0_7] = _init7(new BigNumber(_i0_7));
            }
            _448_v5 = _nw67;
            let _index58 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_448_v5).length));
            (_448_v5)[_index58] = _444_v1;
            let _451_v6;
            _451_v6 = _dafny.MultiSet.fromElements(new BigNumber(794));
            let _452_v7;
            _452_v7 = _dafny.Map.Empty.slice().updateUnsafe(_451_v6,_444_v1);
            let _453_v8;
            _453_v8 = _dafny.Seq.of((_this).f19);
            let _454_v9;
            _454_v9 = _dafny.Seq.of(_442_v0, _442_v0, _442_v0, _dafny.Seq.UnicodeFromString("efknn"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(735)), ((_455_v2) => function (_456_i2) {
              return _455_v2;
            })(_445_v2)));
            let _index59 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_448_v5).length));
            let _rhs75 = _module.D0.create_DC1((_this).f19, _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm23(_444_v1, _444_v1, globalState)).length), _444_v1), (_this).f19);
            let _rhs76 = (_dafny.ZERO).minus(_444_v1);
            let _rhs77 = (_dafny.ZERO).minus((((_452_v7).contains(_451_v6)) ? ((_452_v7).get(_451_v6)) : (new BigNumber((_453_v8).length))));
            let _rhs78 = new BigNumber(((((_this).f19) ? (_442_v0) : ((_454_v9)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_454_v9).length))]))).length);
            let _lhs68 = globalState;
            let _lhs69 = _448_v5;
            let _lhs70 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_448_v5).length));
            _447_v4 = _rhs75;
            _lhs68.f1 = _rhs76;
            _lhs69[_lhs70] = _rhs77;
            _444_v1 = _rhs78;
          }
        }
      }
      let _457_v10;
      _457_v10 = new BigNumber(907);
      let _458_v12;
      let _nw68 = Array((new BigNumber(5)).toNumber());
      _nw68[(_dafny.ZERO).toNumber()] = _module.__default.fm2(_457_v10, _457_v10, globalState);
      _nw68[(_dafny.ONE).toNumber()] = _457_v10;
      _nw68[(new BigNumber(2)).toNumber()] = new BigNumber((_442_v0).length);
      _nw68[(new BigNumber(3)).toNumber()] = new BigNumber((function () {
        let _coll36 = new _dafny.Map();
        for (const _compr_36 of _dafny.IntegerRange(new BigNumber(180), new BigNumber(711))) {
          let _459_v11 = _compr_36;
          if (((new BigNumber(180)).isLessThanOrEqualTo(_459_v11)) && ((_459_v11).isLessThan(new BigNumber(711)))) {
            _coll36.push([(_459_v11).plus(new BigNumber((_442_v0).length)),new BigNumber((_dafny.Set.fromElements(_457_v10)).length)]);
          }
        }
        return _coll36;
      }()).length);
      _nw68[(new BigNumber(4)).toNumber()] = _457_v10;
      _458_v12 = _nw68;
      let _460_v13;
      _460_v13 = _dafny.Set.fromElements(_458_v12);
      _460_v13 = _460_v13;
      let _461_v14;
      _461_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_457_v10);
      let _462_v15;
      _462_v15 = _dafny.Set.fromElements(!((_this).f19), (_this).f18);
      _457_v10 = _module.__default.safeModuloInt((((_461_v14).contains((_this).f18)) ? ((_461_v14).get((_this).f18)) : (_457_v10)), _module.__default.safeModuloInt(new BigNumber((_462_v15).length), _457_v10));
      (globalState).f6 = (_this).f19;
      let _463_i3;
      _463_i3 = _dafny.ZERO;
      L5: {
        while ((_this).f18) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_463_i3)) {
              break L5;
            }
            _463_i3 = (_463_i3).plus(_dafny.ONE);
            let _464_v16;
            _464_v16 = _dafny.Set.fromElements(new BigNumber(789));
            let _465_v17;
            _465_v17 = _module.D7.create_DC12(p2);
            _457_v10 = _module.__default.fm2(new BigNumber(((_dafny.Set.fromElements(_457_v10, _457_v10, _457_v10, _457_v10, _457_v10)).Union(_464_v16)).length), (new BigNumber(((_465_v17).dtor_cf17).length)).minus(_457_v10), globalState);
            let _466_v18;
            _466_v18 = _module.D2.create_DC4(_457_v10, _457_v10, _457_v10);
            (globalState).f1 = _module.__default.fm2(_module.__default.fm2(_457_v10, _457_v10, globalState), (_457_v10).multipliedBy((_466_v18).dtor_cf6), globalState);
            let _467_v19;
            let _init8 = ((_468_v16) => function (_469_i4) {
              return _468_v16;
            })(_464_v16);
            let _nw69 = Array((new BigNumber(3)).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw69.length); _i0_8++) {
              _nw69[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _467_v19 = _nw69;
            let _index60 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_467_v19).length));
            (_467_v19)[_index60] = _464_v16;
            let _470_v20;
            _470_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_module.__default.fm18(new BigNumber(519), globalState));
            let _index61 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_467_v19).length));
            (_467_v19)[_index61] = ((((_470_v20).contains((_this).f18)) ? ((_470_v20).get((_this).f18)) : (_464_v16))).Union(_464_v16);
            (globalState).f6 = (_this).f19;
          }
        }
      }
      let _471_v21;
      let _nw70 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
      _471_v21 = _nw70;
      _471_v21 = _471_v21;
      let _472_v22;
      _472_v22 = _module.D0.create_DC1((_this).f19, _457_v10, false);
      let _pat_let_tv23 = _457_v10;
      let _pat_let_tv24 = _457_v10;
      let _pat_let_tv25 = _457_v10;
      r0 = function (_source8) {
        let _473___mcc_h0 = _source8;
        let _474_cf11 = _473___mcc_h0;
        return _module.D2.create_DC4(_pat_let_tv23, _pat_let_tv24, _pat_let_tv25);
      }(_module.__default.fm24((_472_v22).dtor_cf3, globalState));
      return r0;
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f11 = _dafny.ZERO;
      this._f24 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f24, f11) {
      let _this = this;
      (_this)._f24 = f24;
      (_this)._f11 = f11;
      return;
    }
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this).f11).plus((_this).f11);
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(337)), function (_475_i0) {
        return (_this).f24;
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(428)), function (_476_i1) {
        return (_this).f24;
      })), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pykmo"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-42)), function (_477_i2) {
        return (_this).f24;
      })));
    };
    m8(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _478_v0;
      _478_v0 = false;
      let _479_v1;
      _479_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-775),(_this).fm3((_this).f11, !(_478_v0), _478_v0, globalState));
      let _480_v2;
      _480_v2 = _module.D8.create_DC16(true, (((_479_v1).contains((_this).f11)) ? ((_479_v1).get((_this).f11)) : ((_this).f11)));
      (globalState).f1 = ((_480_v2).dtor_cf24).minus(((_this).f11).plus((_this).f11));
      let _481_v3;
      _481_v3 = _dafny.Seq.of(_478_v0, _478_v0);
      let _482_v4;
      let _nw71 = new _module.C0();
      _nw71.__ctor(_dafny.Seq.Concat(_481_v3, _481_v3), ((_this).f11).multipliedBy(new BigNumber(687)));
      _482_v4 = _nw71;
      let _483_v5;
      _483_v5 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_482_v4).f21),!(_478_v0));
      let _484_v6;
      _484_v6 = _module.D6.create_DC9((_483_v5).Merge((_483_v5).update((_482_v4).f21, _478_v0)));
      let _source9 = _484_v6;
      if (_source9.is_DC10) {
        let _485___mcc_h0 = (_source9).cf13;
        let _486_cf13 = _485___mcc_h0;
        let _487_v7;
        _487_v7 = _dafny.Seq.UnicodeFromString("qpp");
        let _488_v8;
        _488_v8 = _dafny.Seq.of(new BigNumber((_487_v7).length));
        let _489_v9;
        _489_v9 = _module.D7.create_DC12(_488_v8);
        let _490_v10;
        _490_v10 = _dafny.Map.Empty.slice().updateUnsafe(false,_478_v0);
        let _491_v11;
        _491_v11 = _dafny.Map.Empty.slice().updateUnsafe(_490_v10,_478_v0);
        let _492_v12;
        _492_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_491_v11).length),_488_v8);
        let _493_v13;
        let _nw72 = Array((new BigNumber(7)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = _488_v8;
        _nw72[(_dafny.ONE).toNumber()] = _488_v8;
        _nw72[(new BigNumber(2)).toNumber()] = _488_v8;
        _nw72[(new BigNumber(3)).toNumber()] = (_489_v9).dtor_cf17;
        _nw72[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(364)), ((_494_v4) => function (_495_i0) {
          return (_494_v4).f21;
        })(_482_v4));
        _nw72[(new BigNumber(5)).toNumber()] = _488_v8;
        _nw72[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_488_v8, (((_492_v12).contains(new BigNumber(((_482_v4).f20).length))) ? ((_492_v12).get(new BigNumber(((_482_v4).f20).length))) : (_488_v8)));
        _493_v13 = _nw72;
        let _rhs79 = ((_this).f11).isLessThan((_dafny.ZERO).minus((_dafny.ZERO).minus((_482_v4).f21)));
        let _rhs80 = ((_478_v0) ? (_493_v13) : (_493_v13));
        let _rhs81 = _478_v0;
        let _lhs71 = globalState;
        let _lhs72 = globalState;
        _lhs71.f6 = _rhs79;
        _493_v13 = _rhs80;
        _lhs72.f6 = _rhs81;
        let _496_v14;
        _496_v14 = _dafny.MultiSet.fromElements(_478_v0);
        if (!(_496_v14).contains(_478_v0)) {
          let _497_v15;
          _497_v15 = _dafny.Map.Empty.slice().updateUnsafe(_478_v0,(_482_v4).f21);
          r3 = _module.__default.safeDivisionInt((_this).fm3(new BigNumber(((_482_v4).f20).length), _478_v0, _478_v0, globalState), (((_497_v15).contains(_478_v0)) ? ((_497_v15).get(_478_v0)) : ((_this).f11)));
          let _498_v16;
          _498_v16 = _dafny.Set.fromElements((_this).f24);
          let _499_v17;
          _499_v17 = _dafny.Set.fromElements(_498_v16);
          let _500_v18;
          _500_v18 = _dafny.Map.Empty.slice().updateUnsafe(_487_v7,_499_v17);
          _500_v18 = (_500_v18).update(_487_v7, _module.__default.fm28(_478_v0, new BigNumber(-68), globalState));
          let _501_v19;
          let _nw73 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _501_v19 = _nw73;
          let _502_v20;
          let _nw74 = Array((new BigNumber(13)).toNumber());
          _nw74[(_dafny.ZERO).toNumber()] = _501_v19;
          _nw74[(_dafny.ONE).toNumber()] = _501_v19;
          _nw74[(new BigNumber(2)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(3)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(4)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(5)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(6)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(7)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(8)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(9)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(10)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(11)).toNumber()] = _501_v19;
          _nw74[(new BigNumber(12)).toNumber()] = ((false) ? (_501_v19) : (_501_v19));
          _502_v20 = _nw74;
          let _index62 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_502_v20).length));
          (_502_v20)[_index62] = _501_v19;
          let _503_v21;
          _503_v21 = _dafny.MultiSet.fromElements((_this).f11, (_482_v4).f21);
          let _index63 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_502_v20).length));
          let _rhs82 = (_dafny.ZERO).minus((_482_v4).f21);
          let _rhs83 = ((_this).f11).plus(_module.__default.safeDivisionInt((((_503_v21).contains((_482_v4).f21)) ? ((_503_v21).get((_482_v4).f21)) : ((((_479_v1).contains((_482_v4).f21)) ? ((_479_v1).get((_482_v4).f21)) : ((_482_v4).f21)))), (_this).f11));
          let _rhs84 = _501_v19;
          let _lhs73 = globalState;
          let _lhs74 = globalState;
          let _lhs75 = _502_v20;
          let _lhs76 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_502_v20).length));
          _lhs73.f1 = _rhs82;
          _lhs74.f1 = _rhs83;
          _lhs75[_lhs76] = _rhs84;
          (globalState).f7 = _478_v0;
          let _504_v22;
          _504_v22 = _dafny.Set.fromElements(((_478_v0) ? (_478_v0) : (false)));
          let _505_v23;
          _505_v23 = new _dafny.CodePoint('v'.codePointAt(0));
          let _rhs85 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(249)), ((_506_v7, _507_v0, _508_v1, _509_v4) => function (_510_i1) {
            return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_506_v7).length),_507_v0)).length)).minus((_dafny.ZERO).minus((((_508_v1).contains(new BigNumber(968))) ? ((_508_v1).get(new BigNumber(968))) : ((_509_v4).f21))));
          })(_487_v7, _478_v0, _479_v1, _482_v4));
          let _rhs86 = _504_v22;
          let _rhs87 = ((_478_v0) ? ((_this).f24) : (_505_v23));
          let _rhs88 = (_482_v4).f21;
          let _lhs77 = globalState;
          _488_v8 = _rhs85;
          _504_v22 = _rhs86;
          _505_v23 = _rhs87;
          _lhs77.f1 = _rhs88;
        } else {
          _483_v5 = (_483_v5).update(((_478_v0) ? (_module.__default.fm2((_dafny.ZERO).minus((_482_v4).f21), (_dafny.ZERO).minus((_this).f11), globalState)) : (new BigNumber(-92))), _478_v0);
          let _511_v24;
          _511_v24 = _dafny.Set.fromElements(((_478_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(260)), ((_512_v4) => function (_513_i2) {
            return (_512_v4).f21;
          })(_482_v4))) : (_488_v8)));
          _511_v24 = _511_v24;
          let _514_v25;
          _514_v25 = new _dafny.CodePoint('i'.codePointAt(0));
          let _515_v26;
          _515_v26 = _dafny.Set.fromElements(_module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(_478_v0,(_482_v4).f21)));
          let _516_v27;
          let _nw75 = Array((new BigNumber(9)).toNumber());
          _nw75[(_dafny.ZERO).toNumber()] = _515_v26;
          _nw75[(_dafny.ONE).toNumber()] = _515_v26;
          _nw75[(new BigNumber(2)).toNumber()] = _515_v26;
          _nw75[(new BigNumber(3)).toNumber()] = _515_v26;
          _nw75[(new BigNumber(4)).toNumber()] = (_515_v26).Difference(_515_v26);
          _nw75[(new BigNumber(5)).toNumber()] = _515_v26;
          _nw75[(new BigNumber(6)).toNumber()] = _module.__default.fm29(globalState);
          _nw75[(new BigNumber(7)).toNumber()] = _515_v26;
          _nw75[(new BigNumber(8)).toNumber()] = _515_v26;
          _516_v27 = _nw75;
          let _index64 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_516_v27).length));
          (_516_v27)[_index64] = _515_v26;
          let _index65 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_516_v27).length));
          let _rhs89 = (_this).f24;
          let _rhs90 = _515_v26;
          let _lhs78 = _516_v27;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_516_v27).length));
          _514_v25 = _rhs89;
          _lhs78[_lhs79] = _rhs90;
          let _517_v28;
          _517_v28 = _module.D8.create_DC15(_490_v10);
          let _518_v29;
          _518_v29 = _module.D8.create_DC17(_517_v28);
          _518_v29 = _518_v29;
          let _519_v30;
          let _nw76 = new _module.C0();
          _nw76.__ctor(_dafny.Seq.of(_478_v0), _module.__default.safeDivisionInt((_482_v4).f21, (_482_v4).f21));
          _519_v30 = _nw76;
        }
        let _520_v31;
        let _nw77 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _520_v31 = _nw77;
        let _index66 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_520_v31).length));
        (_520_v31)[_index66] = (_482_v4).f21;
        let _index67 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_520_v31).length));
        (_520_v31)[_index67] = _module.__default.safeModuloInt((_482_v4).f21, (_480_v2).dtor_cf24);
        r3 = (_module.__default.safeModuloInt(new BigNumber(359), (_482_v4).f21)).multipliedBy((_this).f11);
      } else if (_source9.is_DC11) {
        let _521___mcc_h1 = (_source9).cf14;
        let _522___mcc_h2 = (_source9).cf15;
        let _523___mcc_h3 = (_source9).cf16;
        let _524_cf16 = _523___mcc_h3;
        let _525_cf15 = _522___mcc_h2;
        let _526_cf14 = _521___mcc_h1;
        let _527_v32;
        _527_v32 = _dafny.Seq.of((_this).f11, _526_cf14);
        let _528_v33;
        let _nw78 = Array((new BigNumber(8)).toNumber());
        _nw78[(_dafny.ZERO).toNumber()] = (_this).f11;
        _nw78[(_dafny.ONE).toNumber()] = (_482_v4).f21;
        _nw78[(new BigNumber(2)).toNumber()] = new BigNumber(552);
        _nw78[(new BigNumber(3)).toNumber()] = new BigNumber((_527_v32).length);
        _nw78[(new BigNumber(4)).toNumber()] = _524_cf16;
        _nw78[(new BigNumber(5)).toNumber()] = new BigNumber(683);
        _nw78[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_this).f11);
        _nw78[(new BigNumber(7)).toNumber()] = _526_cf14;
        _528_v33 = _nw78;
        let _529_v34;
        _529_v34 = _dafny.Map.Empty.slice().updateUnsafe(((_478_v0) ? (_528_v33) : (_528_v33)),_478_v0);
        _529_v34 = (_529_v34).update(_528_v33, (_482_v4).fm21(_478_v0, new BigNumber(777), globalState));
        let _530_v35;
        _530_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_479_v1).length),_528_v33);
        _530_v35 = _530_v35;
        let _531_v36;
        _531_v36 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_527_v32).length)).multipliedBy((_482_v4).f21),(_this).f24);
        let _532_v37;
        _532_v37 = _module.D7.create_DC12(_527_v32);
        let _533_v39;
        _533_v39 = _dafny.MultiSet.fromElements(new BigNumber(557), _524_cf16, new BigNumber((_dafny.Seq.UnicodeFromString("iwljmxlt")).length));
        let _534_v40;
        _534_v40 = _dafny.Map.Empty.slice().updateUnsafe(_478_v0,_482_v4);
        let _rhs91 = (function () {
          let _coll37 = new _dafny.Set();
          for (const _compr_37 of (_dafny.Seq.of((_532_v37).dtor_cf17, _dafny.Seq.of(new BigNumber(144)))).Elements) {
            let _535_v38 = _compr_37;
            if (_dafny.Seq.contains(_dafny.Seq.of((_532_v37).dtor_cf17, _dafny.Seq.of(new BigNumber(144))), _535_v38)) {
              _coll37.add(_535_v38);
            }
          }
          return _coll37;
        }()).contains(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_482_v4).f21), _module.__default.safeIndex((_dafny.ZERO).minus((_482_v4).f21), new BigNumber((_dafny.Seq.of((_482_v4).f21)).length)), (_this).f11), _dafny.Seq.of((_482_v4).f21)));
        let _rhs92 = _module.__default.safeModuloInt((_482_v4).f21, (new BigNumber((_533_v39).cardinality())).plus((_dafny.ZERO).minus((_482_v4).f21)));
        let _rhs93 = _531_v36;
        let _rhs94 = !((_534_v40).Merge(_534_v40)).contains(_478_v0);
        let _rhs95 = ((_482_v4).f20)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-724)), new BigNumber(((_482_v4).f20).length))];
        let _lhs80 = globalState;
        let _lhs81 = globalState;
        let _lhs82 = globalState;
        _lhs80.f7 = _rhs91;
        _lhs81.f1 = _rhs92;
        _531_v36 = _rhs93;
        _lhs82.f6 = _rhs94;
        r0 = _rhs95;
        _526_cf14 = (_482_v4).f21;
      } else {
        let _536___mcc_h4 = (_source9).cf12;
        let _537_cf12 = _536___mcc_h4;
        if ((_478_v0) || (((_this).f11).isLessThanOrEqualTo((_482_v4).f21))) {
          (globalState).f1 = (_482_v4).f21;
          (globalState).f7 = false;
          let _538_v41;
          let _init9 = function (_539_i3) {
            return (_539_i3).multipliedBy((_this).f11);
          };
          let _nw79 = Array((new BigNumber(2)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw79.length); _i0_9++) {
            _nw79[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _538_v41 = _nw79;
          let _540_v42;
          _540_v42 = _dafny.Seq.of(_538_v41, _538_v41, _538_v41);
          let _541_v43;
          _541_v43 = _dafny.Seq.of((_540_v42)[_module.__default.safeIndex((_482_v4).f21, new BigNumber((_540_v42).length))], _538_v41, _538_v41);
          let _542_v44;
          _542_v44 = _dafny.Seq.UnicodeFromString("miujifbj");
          let _543_v45;
          _543_v45 = _dafny.Seq.of((_482_v4).f21);
          let _544_v46;
          _544_v46 = _dafny.Seq.of(_543_v45);
          let _545_v47;
          _545_v47 = _module.D9.create_DC19();
          let _546_v48;
          _546_v48 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_482_v4).f21),_479_v1);
          let _547_v50;
          _547_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qfr")).length),_dafny.Set.fromElements((_this).f11));
          let _548_v52;
          _548_v52 = _dafny.Set.fromElements((_482_v4).f21, (_this).f11, (_482_v4).f21);
          let _549_v53;
          let _nw80 = Array((new BigNumber(20)).toNumber());
          _nw80[(_dafny.ZERO).toNumber()] = ((!(_478_v0)) ? (_478_v0) : (_478_v0));
          _nw80[(_dafny.ONE).toNumber()] = _478_v0;
          _nw80[(new BigNumber(2)).toNumber()] = _dafny.areEqual(_544_v46, _544_v46);
          _nw80[(new BigNumber(3)).toNumber()] = _478_v0;
          _nw80[(new BigNumber(4)).toNumber()] = (_482_v4).fm21(_478_v0, (_482_v4).f21, globalState);
          _nw80[(new BigNumber(5)).toNumber()] = _478_v0;
          _nw80[(new BigNumber(6)).toNumber()] = (_module.__default.fm30((_482_v4).f21, false, (_482_v4).f21, new BigNumber(930), globalState)).contains((_482_v4).f21);
          _nw80[(new BigNumber(7)).toNumber()] = true;
          _nw80[(new BigNumber(8)).toNumber()] = _478_v0;
          _nw80[(new BigNumber(9)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.of(_545_v47), _545_v47);
          _nw80[(new BigNumber(10)).toNumber()] = _478_v0;
          _nw80[(new BigNumber(11)).toNumber()] = (_this).fm4((_this).f11, _546_v48, new BigNumber((function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of (_module.__default.fm31((_this).f24, globalState)).Elements) {
              let _550_v49 = _compr_38;
              if ((_module.__default.fm31((_this).f24, globalState)).contains(_550_v49)) {
                _coll38.push([_550_v49,_478_v0]);
              }
            }
            return _coll38;
          }()).length), false, globalState);
          _nw80[(new BigNumber(12)).toNumber()] = ((_478_v0) ? ((_481_v3)[_module.__default.safeIndex((_this).f11, new BigNumber((_481_v3).length))]) : (_478_v0));
          _nw80[(new BigNumber(13)).toNumber()] = false;
          _nw80[(new BigNumber(14)).toNumber()] = !((_548_v52).IsSubsetOf((((_547_v50).contains((_482_v4).f21)) ? ((_547_v50).get((_482_v4).f21)) : (function () {
            let _coll39 = new _dafny.Set();
            for (const _compr_39 of (_dafny.Seq.of((_482_v4).f21)).Elements) {
              let _551_v51 = _compr_39;
              if (_dafny.Seq.contains(_dafny.Seq.of((_482_v4).f21), _551_v51)) {
                _coll39.add((_551_v51).minus(new BigNumber((_dafny.Seq.UnicodeFromString("knsv")).length)));
              }
            }
            return _coll39;
          }()))));
          _nw80[(new BigNumber(15)).toNumber()] = _478_v0;
          _nw80[(new BigNumber(16)).toNumber()] = false;
          _nw80[(new BigNumber(17)).toNumber()] = _478_v0;
          _nw80[(new BigNumber(18)).toNumber()] = ((_482_v4).f20)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f11), new BigNumber(((_482_v4).f20).length))];
          _nw80[(new BigNumber(19)).toNumber()] = _478_v0;
          _549_v53 = _nw80;
          let _index68 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length));
          (_549_v53)[_index68] = !(_478_v0);
          let _index69 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length));
          let _rhs96 = _dafny.Seq.Concat(_dafny.Seq.Concat(_540_v42, _541_v43), _541_v43);
          let _rhs97 = _dafny.Seq.Concat(_module.__default.fm1(globalState), ((!(true)) ? (_542_v44) : (_dafny.Seq.UnicodeFromString("vtq"))));
          let _rhs98 = true;
          let _rhs99 = (_482_v4).f21;
          let _lhs83 = _549_v53;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length));
          let _lhs85 = globalState;
          _541_v43 = _rhs96;
          _542_v44 = _rhs97;
          _lhs83[_lhs84] = _rhs98;
          _lhs85.f1 = _rhs99;
          let _552_v54;
          _552_v54 = _dafny.Map.Empty.slice().updateUnsafe((_549_v53)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length))],new BigNumber((_479_v1).length));
          let _553_v55;
          _553_v55 = _dafny.Map.Empty.slice().updateUnsafe(((_482_v4).f21).multipliedBy(new BigNumber((_542_v44).length)),_552_v54);
          let _554_v57;
          _554_v57 = _dafny.MultiSet.fromElements((_482_v4).f21);
          let _index70 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length));
          let _index71 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length));
          let _rhs100 = (((_this).fm4((_482_v4).f21, _546_v48, (_this).f11, (_this).fm4(new BigNumber(635), function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of (_dafny.Set.fromElements(_554_v57)).Elements) {
              let _555_v56 = _compr_40;
              if ((_dafny.Set.fromElements(_554_v57)).contains(_555_v56)) {
                _coll40.push([_555_v56,_479_v1]);
              }
            }
            return _coll40;
          }(), (_482_v4).f21, (_549_v53)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length))], globalState), globalState)) ? ((_549_v53)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length))]) : ((_549_v53)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length))]));
          let _rhs101 = !(!(_483_v5).equals(_537_cf12)) || (((_478_v0) ? (_478_v0) : ((_549_v53)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length))])));
          let _rhs102 = _478_v0;
          let _rhs103 = (_module.__default.fm32((_482_v4).f21, (_482_v4).f21, (_482_v4).f20, (_482_v4).f21, globalState)).Merge(_553_v55);
          let _rhs104 = (_549_v53)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length))];
          let _lhs86 = _549_v53;
          let _lhs87 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length));
          let _lhs88 = _549_v53;
          let _lhs89 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_549_v53).length));
          r2 = _rhs100;
          r0 = _rhs101;
          _lhs86[_lhs87] = _rhs102;
          _553_v55 = _rhs103;
          _lhs88[_lhs89] = _rhs104;
          let _556_v58;
          let _nw81 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
          _556_v58 = _nw81;
          let _557_v59;
          _557_v59 = _dafny.Map.Empty.slice().updateUnsafe(_556_v58,_549_v53);
          _557_v59 = (_557_v59).update(_556_v58, _549_v53);
        } else {
          (globalState).f6 = (!(_478_v0)) && ((_dafny.MultiSet.fromElements((_482_v4).f21)).contains((_482_v4).f21));
          let _558_v60;
          _558_v60 = _dafny.Map.Empty.slice().updateUnsafe(_478_v0,_dafny.Seq.of((_482_v4).f21, (_482_v4).f21));
          let _559_v61;
          _559_v61 = _dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)));
          let _560_v62;
          _560_v62 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(442)), function (_561_i4) {
            return (_this).f24;
          }));
          let _562_v63;
          _562_v63 = _module.D7.create_DC12(_dafny.Seq.of(new BigNumber((_559_v61).length), new BigNumber(((_560_v62)[_module.__default.safeIndex((_this).f11, new BigNumber((_560_v62).length))]).length)));
          let _563_v64;
          _563_v64 = _dafny.Seq.of((_this).f11, new BigNumber((_559_v61).length), (_482_v4).f21);
          let _pat_let_tv26 = _563_v64;
          _558_v60 = (_558_v60).update(_module.__default.fm0(globalState), _dafny.Seq.Concat((function (_pat_let12_0) {
            return function (_564_dt__update__tmp_h0) {
              return function (_pat_let13_0) {
                return function (_565_dt__update_hcf17_h0) {
                  return _module.D7.create_DC12(_565_dt__update_hcf17_h0);
                }(_pat_let13_0);
              }(_pat_let_tv26);
            }(_pat_let12_0);
          }(_562_v63)).dtor_cf17, _563_v64));
          let _566_v66;
          _566_v66 = _dafny.MultiSet.fromElements(function () {
            let _coll41 = new _dafny.Set();
            for (const _compr_41 of _dafny.IntegerRange(new BigNumber(891), new BigNumber(318))) {
              let _567_v65 = _compr_41;
              if (((new BigNumber(891)).isLessThanOrEqualTo(_567_v65)) && ((_567_v65).isLessThan(new BigNumber(318)))) {
                _coll41.add((_567_v65).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(338)), ((_568_v4) => function (_569_i5) {
                  return (_568_v4).f21;
                })(_482_v4))).length)));
              }
            }
            return _coll41;
          }());
          _566_v66 = (_566_v66).Difference(_566_v66);
          let _570_v67;
          _570_v67 = _dafny.MultiSet.fromElements((_482_v4).f21, (_482_v4).f21, (_482_v4).f21);
          _570_v67 = _570_v67;
          (globalState).f1 = ((_482_v4).f21).minus((_482_v4).f21);
        }
        let _571_v68;
        _571_v68 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_478_v0)),(_this).f11);
        let _572_v69;
        _572_v69 = _dafny.Set.fromElements((_this).fm3(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("sbmbrcsg"), _module.__default.safeIndex((_this).f11, new BigNumber((_dafny.Seq.UnicodeFromString("sbmbrcsg")).length)), (_this).f24)).length), _478_v0, _478_v0, globalState), (_482_v4).f21, _module.__default.fm2((_482_v4).f21, (_this).f11, globalState), new BigNumber((_571_v68).length));
        _572_v69 = _572_v69;
        let _573_v70;
        _573_v70 = new _dafny.CodePoint('d'.codePointAt(0));
        _573_v70 = ((_478_v0) ? ((_this).f24) : (_573_v70));
        let _574_v71;
        _574_v71 = _module.D9.create_DC19();
        _574_v71 = _574_v71;
      }
      let _575_v72;
      _575_v72 = _dafny.Seq.UnicodeFromString("e");
      let _576_v73;
      _576_v73 = _dafny.Map.Empty.slice().updateUnsafe(_575_v72,_478_v0);
      r2 = (_576_v73).equals(_576_v73);
      let _577_v74;
      _577_v74 = _dafny.MultiSet.fromElements((_this).f11, (_482_v4).f21);
      let _578_v76;
      _578_v76 = _dafny.Map.Empty.slice().updateUnsafe(_577_v74,function () {
        let _coll42 = new _dafny.Map();
        for (const _compr_42 of _dafny.IntegerRange(new BigNumber(679), new BigNumber(599))) {
          let _579_v75 = _compr_42;
          if (((new BigNumber(679)).isLessThanOrEqualTo(_579_v75)) && ((_579_v75).isLessThan(new BigNumber(599)))) {
            _coll42.push([_module.__default.safeDivisionInt(_579_v75, (_module.D7.create_DC13(_dafny.Set.fromElements(_478_v0), (_this).f11, (_482_v4).f21)).dtor_cf20),new BigNumber((_575_v72).length)]);
          }
        }
        return _coll42;
      }());
      let _580_v77;
      _580_v77 = _module.D0.create_DC1(_478_v0, (_this).f11, _478_v0);
      r2 = !((_this).fm4(new BigNumber(-452), _578_v76, (_482_v4).f21, (_580_v77).dtor_cf1, globalState));
      let _581_v78;
      _581_v78 = _dafny.Set.fromElements(false, _478_v0);
      let _582_v79;
      _582_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_581_v78);
      let _583_v80;
      _583_v80 = _module.D7.create_DC13((((_582_v79).contains((_this).f24)) ? ((_582_v79).get((_this).f24)) : (_581_v78)), (_482_v4).f21, (_this).f11);
      let _584_v81;
      _584_v81 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_478_v0);
      _583_v80 = _module.D7.create_DC13((_dafny.Set.fromElements(_478_v0, (((_584_v81).contains((_this).f24)) ? ((_584_v81).get((_this).f24)) : (_478_v0)))).Union(_581_v78), (_482_v4).f21, new BigNumber(71));
      let _585_v82;
      _585_v82 = _dafny.Seq.of((_this).f11);
      let _586_v83;
      _586_v83 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_585_v82,_478_v0)).length));
      let _587_v84;
      let _nw82 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _587_v84 = _nw82;
      let _588_v85;
      _588_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_587_v84);
      r0 = (_482_v4).fm21((_dafny.MultiSet.fromElements((_this).f11)).IsProperSubsetOf(_577_v74), _module.__default.safeModuloInt(new BigNumber((_585_v82).length), new BigNumber((_588_v85).length)), globalState);
      r1 = (((_576_v73).contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xbqkmva"), _dafny.Seq.UnicodeFromString("bnu")))) ? ((_576_v73).get(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xbqkmva"), _dafny.Seq.UnicodeFromString("bnu")))) : (_478_v0));
      let _589_v86;
      _589_v86 = _dafny.Seq.of(_481_v3);
      let _590_v87;
      _590_v87 = _dafny.Seq.of(_589_v86, _dafny.Seq.Create(_module.__default.abs(new BigNumber(308)), ((_591_v3) => function (_592_i6) {
        return _591_v3;
      })(_481_v3)), _589_v86, _589_v86);
      r2 = (((_this).f11).multipliedBy((_482_v4).f21)).isLessThan(new BigNumber(((_590_v87)[_module.__default.safeIndex((_482_v4).f21, new BigNumber((_590_v87).length))]).length));
      r3 = ((_482_v4).f21).minus((_this).f11);
      return [r0, r1, r2, r3];
    }
    m9(p0, p1, globalState) {
      let _this = this;
      let r0 = _module.D7.Default();
      let r1 = [];
      let _593_i0;
      _593_i0 = _dafny.ZERO;
      L6: {
        while (p0) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_593_i0)) {
              break L6;
            }
            _593_i0 = (_593_i0).plus(_dafny.ONE);
            let _594_v0;
            _594_v0 = _module.D0.create_DC1(p0, (_this).f11, _module.__default.fm0(globalState));
            let _pat_let_tv27 = p0;
            let _pat_let_tv28 = p0;
            let _pat_let_tv29 = p0;
            let _source10 = function (_pat_let14_0) {
              return function (_595_dt__update__tmp_h0) {
                return function (_pat_let15_0) {
                  return function (_596_dt__update_hcf1_h0) {
                    return _module.D0.create_DC1(_596_dt__update_hcf1_h0, (_595_dt__update__tmp_h0).dtor_cf2, (_595_dt__update__tmp_h0).dtor_cf3);
                  }(_pat_let15_0);
                }(((_pat_let_tv29) ? (_pat_let_tv27) : (_pat_let_tv28)));
              }(_pat_let14_0);
            }(_594_v0);
            if (_source10.is_DC1) {
              let _597___mcc_h0 = (_source10).cf1;
              let _598___mcc_h1 = (_source10).cf2;
              let _599___mcc_h2 = (_source10).cf3;
              let _600_cf3 = _599___mcc_h2;
              let _601_cf2 = _598___mcc_h1;
              let _602_cf1 = _597___mcc_h0;
              (globalState).f6 = false;
              let _603_v1;
              let _604_v2;
              let _605_v3;
              let _606_v4;
              let _out16;
              let _out17;
              let _out18;
              let _out19;
              let _outcollector6 = (_this).m8(globalState);
              _out16 = _outcollector6[0];
              _out17 = _outcollector6[1];
              _out18 = _outcollector6[2];
              _out19 = _outcollector6[3];
              _603_v1 = _out16;
              _604_v2 = _out17;
              _605_v3 = _out18;
              _606_v4 = _out19;
              let _607_v5;
              let _nw83 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
              _607_v5 = _nw83;
              _607_v5 = _607_v5;
              let _608_v6;
              _608_v6 = _module.D9.create_DC19();
              _608_v6 = _608_v6;
            } else {
              let _609___mcc_h3 = (_source10).cf0;
              let _610_cf0 = _609___mcc_h3;
              let _611_v7;
              _611_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              (globalState).f6 = (_610_cf0).isLessThan(new BigNumber(((_611_v7).Merge(_611_v7)).length));
              let _612_v8;
              _612_v8 = _dafny.Map.Empty.slice().updateUnsafe(_610_cf0,_610_cf0);
              let _613_v9;
              _613_v9 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-762)), ((_614_v8) => function (_615_i1) {
                return _614_v8;
              })(_612_v8)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-221)), ((_616_v8) => function (_617_i2) {
                return _616_v8;
              })(_612_v8)));
              let _618_v10;
              let _nw84 = Array((new BigNumber(11)).toNumber());
              _nw84[(_dafny.ZERO).toNumber()] = _610_cf0;
              _nw84[(_dafny.ONE).toNumber()] = (_this).f11;
              _nw84[(new BigNumber(2)).toNumber()] = (_this).f11;
              _nw84[(new BigNumber(3)).toNumber()] = (new BigNumber((_612_v8).length)).plus(new BigNumber(((_613_v9)[_module.__default.safeIndex((_this).f11, new BigNumber((_613_v9).length))]).length));
              _nw84[(new BigNumber(4)).toNumber()] = _610_cf0;
              _nw84[(new BigNumber(5)).toNumber()] = (_this).f11;
              _nw84[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt((_this).f11, (_this).f11);
              _nw84[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_610_cf0);
              _nw84[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Set.fromElements(!(p0))).length);
              _nw84[(new BigNumber(9)).toNumber()] = _610_cf0;
              _nw84[(new BigNumber(10)).toNumber()] = (_this).f11;
              _618_v10 = _nw84;
              let _619_v11;
              _619_v11 = _dafny.Set.fromElements(_610_cf0);
              let _620_v12;
              _620_v12 = _dafny.Map.Empty.slice().updateUnsafe(_610_cf0,_619_v11);
              let _rhs105 = _618_v10;
              let _rhs106 = ((_dafny.ZERO).minus(_610_cf0)).minus((_this).f11);
              let _rhs107 = p0;
              let _rhs108 = !(new BigNumber(((((_620_v12).contains(_610_cf0)) ? ((_620_v12).get(_610_cf0)) : (_dafny.Set.fromElements(_610_cf0, (_this).f11)))).length)).isEqualTo((_this).f11);
              let _lhs90 = globalState;
              let _lhs91 = globalState;
              let _lhs92 = globalState;
              _618_v10 = _rhs105;
              _lhs90.f1 = _rhs106;
              _lhs91.f7 = _rhs107;
              _lhs92.f7 = _rhs108;
              let _621_v13;
              _621_v13 = _dafny.MultiSet.fromElements(_610_cf0);
              let _622_v14;
              _622_v14 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements((_this).f11)).Union(_621_v13),p0);
              (globalState).f7 = (((_622_v14).contains((_621_v13).Difference(_dafny.MultiSet.fromElements((_this).f11)))) ? ((_622_v14).get((_621_v13).Difference(_dafny.MultiSet.fromElements((_this).f11)))) : (p0));
              let _623_v15;
              _623_v15 = _dafny.Seq.of(p0, p0);
              let _624_v16;
              _624_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,false);
              let _625_v17;
              let _nw85 = new _module.C0();
              _nw85.__ctor(_623_v15, _module.__default.fm2(_610_cf0, new BigNumber((_624_v16).length), globalState));
              _625_v17 = _nw85;
              let _626_v18;
              _626_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,_625_v17);
              let _627_v19;
              _627_v19 = _dafny.Map.Empty.slice().updateUnsafe((((_626_v18).contains(p0)) ? ((_626_v18).get(p0)) : (_625_v17)),_610_cf0);
              (globalState).f1 = (((_627_v19).contains(_625_v17)) ? ((_627_v19).get(_625_v17)) : (new BigNumber(893)));
            }
            let _628_v20;
            let _nw86 = Array((new BigNumber(13)).toNumber()).fill([]);
            _628_v20 = _nw86;
            let _629_v21;
            let _nw87 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
            _629_v21 = _nw87;
            let _index72 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_628_v20).length));
            (_628_v20)[_index72] = _629_v21;
            let _630_v22;
            _630_v22 = _dafny.Seq.of(_629_v21, _629_v21);
            let _index73 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_628_v20).length));
            (_628_v20)[_index73] = (_630_v22)[_module.__default.safeIndex((_this).f11, new BigNumber((_630_v22).length))];
            let _631_v23;
            _631_v23 = _dafny.Seq.UnicodeFromString("vb");
            let _632_v24;
            _632_v24 = _dafny.Seq.of(p0);
            let _633_v25;
            _633_v25 = _dafny.MultiSet.fromElements(((p0) ? (_632_v24) : (_632_v24)), _632_v24, _dafny.Seq.Concat(_632_v24, _632_v24), _632_v24);
            let _634_v26;
            _634_v26 = _dafny.MultiSet.fromElements(p0);
            let _635_v27;
            _635_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_634_v26);
            let _636_v29;
            _636_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f11);
            let _637_v30;
            _637_v30 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((function () {
              let _coll43 = new _dafny.Set();
              for (const _compr_43 of (_dafny.Seq.update(_631_v23, _module.__default.safeIndex((_this).f11, new BigNumber((_631_v23).length)), (_this).f24)).Elements) {
                let _638_v28 = _compr_43;
                if (_dafny.Seq.contains(_dafny.Seq.update(_631_v23, _module.__default.safeIndex((_this).f11, new BigNumber((_631_v23).length)), (_this).f24), _638_v28)) {
                  _coll43.add(_638_v28);
                }
              }
              return _coll43;
            }()).length)),_636_v29);
            let _rhs109 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_631_v23, _module.__default.safeIndex(new BigNumber(((_635_v27).update((_this).f11, _dafny.MultiSet.fromElements(p0, p0))).length), new BigNumber((_631_v23).length)), (_this).f24), _dafny.Seq.UnicodeFromString("vjdgtkgo")), _dafny.Seq.Concat(_631_v23, _631_v23));
            let _rhs110 = !(true) || (_module.__default.fm0(globalState));
            let _rhs111 = (((p0) && (p0)) ? (_633_v25) : ((_633_v25).Difference(_dafny.MultiSet.fromElements(_632_v24, _632_v24, _632_v24, _632_v24, _632_v24))));
            let _rhs112 = (((_634_v26).contains((_this).fm4((_this).f11, _637_v30, (_this).f11, p0, globalState))) ? ((_634_v26).get((_this).fm4((_this).f11, _637_v30, (_this).f11, p0, globalState))) : ((_this).f11));
            let _lhs93 = globalState;
            let _lhs94 = globalState;
            _631_v23 = _rhs109;
            _lhs93.f7 = _rhs110;
            _633_v25 = _rhs111;
            _lhs94.f1 = _rhs112;
            (globalState).f7 = p0;
          }
        }
      }
      let _639_v31;
      _639_v31 = _dafny.MultiSet.fromElements((_this).f11);
      let _640_v32;
      _640_v32 = _dafny.Seq.of((_this).f11, (_this).f11);
      let _hi0 = new BigNumber(770);
      for (let _641_i3 = new BigNumber(((_639_v31).Difference(_dafny.MultiSet.FromArray(_640_v32))).cardinality()); _641_i3.isLessThan(_hi0); _641_i3 = _641_i3.plus(_dafny.ONE)) {
        (globalState).f6 = p0;
        let _642_v33;
        _642_v33 = _module.D8.create_DC16(p0, _641_i3);
        let _source11 = _642_v33;
        if (_source11.is_DC16) {
          let _643___mcc_h4 = (_source11).cf23;
          let _644___mcc_h5 = (_source11).cf24;
          let _645_cf24 = _644___mcc_h5;
          let _646_cf23 = _643___mcc_h4;
          let _647_v34;
          let _init10 = function (_648_i4) {
            return _module.__default.safeModuloInt(_648_i4, (_this).f11);
          };
          let _nw88 = Array((new BigNumber(4)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw88.length); _i0_10++) {
            _nw88[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _647_v34 = _nw88;
          let _649_v35;
          _649_v35 = _dafny.Map.Empty.slice().updateUnsafe(_647_v34,p0);
          let _650_v36;
          _650_v36 = _dafny.Seq.of(p0, _646_cf23);
          let _651_v37;
          _651_v37 = _dafny.MultiSet.fromElements((_650_v36)[_module.__default.safeIndex(_645_cf24, new BigNumber((_650_v36).length))], !(p0));
          let _652_v38;
          _652_v38 = _651_v37;
          let _653_v39;
          _653_v39 = _dafny.Set.fromElements(p0, p0);
          let _654_v40;
          _654_v40 = _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.Set.fromElements(_646_cf23, _module.__default.fm0(globalState), _646_cf23, p0, p0)).IsProperSubsetOf(_653_v39));
          let _655_v41;
          _655_v41 = _module.D10.create_DC20(_649_v35);
          let _rhs113 = (((_654_v40).contains(_646_cf23)) ? ((_654_v40).get(_646_cf23)) : (false));
          let _rhs114 = (_649_v35).Merge((_655_v41).dtor_cf27);
          let _rhs115 = (_this).f11;
          let _rhs116 = _652_v38;
          let _lhs95 = globalState;
          let _lhs96 = globalState;
          _lhs95.f6 = _rhs113;
          _649_v35 = _rhs114;
          _lhs96.f1 = _rhs115;
          _652_v38 = _rhs116;
          let _656_v42;
          let _nw89 = Array((_dafny.ONE).toNumber()).fill([]);
          _656_v42 = _nw89;
          let _657_v43;
          _657_v43 = _656_v42;
          let _658_v44;
          let _nw90 = Array((new BigNumber(12)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _656_v42;
          _nw90[(_dafny.ONE).toNumber()] = _657_v43;
          _nw90[(new BigNumber(2)).toNumber()] = _657_v43;
          _nw90[(new BigNumber(3)).toNumber()] = _657_v43;
          _nw90[(new BigNumber(4)).toNumber()] = _657_v43;
          _nw90[(new BigNumber(5)).toNumber()] = _657_v43;
          _nw90[(new BigNumber(6)).toNumber()] = _657_v43;
          _nw90[(new BigNumber(7)).toNumber()] = _657_v43;
          _nw90[(new BigNumber(8)).toNumber()] = _657_v43;
          _nw90[(new BigNumber(9)).toNumber()] = _656_v42;
          _nw90[(new BigNumber(10)).toNumber()] = _657_v43;
          _nw90[(new BigNumber(11)).toNumber()] = _657_v43;
          _658_v44 = _nw90;
          let _659_v45;
          _659_v45 = _dafny.Map.Empty.slice().updateUnsafe(_658_v44,_647_v34);
          let _660_v46;
          _660_v46 = _dafny.Map.Empty.slice().updateUnsafe(_659_v45,p0);
          let _661_v47;
          _661_v47 = _module.D11.create_DC22((_659_v45).update(_658_v44, _647_v34));
          _660_v46 = (_660_v46).update((_659_v45).Merge((_661_v47).dtor_cf30), p0);
          let _662_v48;
          _662_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_641_i3).minus(_645_cf24));
          let _index74 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_647_v34).length));
          (_647_v34)[_index74] = new BigNumber(-915);
          let _index75 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_647_v34).length));
          let _rhs117 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber(196));
          let _rhs118 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(414)), function (_663_i5) {
            return (_this).f24;
          })).length));
          let _rhs119 = (_dafny.ZERO).minus(((_this).f11).multipliedBy(new BigNumber(630)));
          let _lhs97 = _647_v34;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_647_v34).length));
          _662_v48 = _rhs117;
          _lhs97[_lhs98] = _rhs118;
          _645_cf24 = _rhs119;
          let _664_v49;
          _664_v49 = _dafny.Map.Empty.slice().updateUnsafe(_656_v42,_645_cf24);
          _664_v49 = (_664_v49).update(((_646_cf23) ? (_657_v43) : (_657_v43)), (_dafny.ZERO).minus(_module.__default.safeModuloInt(_645_cf24, (_this).fm3((_this).f11, _646_cf23, p0, globalState))));
        } else if (_source11.is_DC15) {
          let _665___mcc_h6 = (_source11).cf22;
          let _666_cf22 = _665___mcc_h6;
          _640_v32 = _640_v32;
          let _667_v50;
          _667_v50 = _dafny.Seq.of(p0, p0, p0, p0);
          let _668_v51;
          _668_v51 = _module.D6.create_DC10(_dafny.Seq.Concat(_667_v50, _667_v50));
          _668_v51 = _668_v51;
          let _669_v52;
          _669_v52 = _module.D11.create_DC23(false, _641_i3, new BigNumber(662));
          let _670_v53;
          _670_v53 = _dafny.Seq.UnicodeFromString("o");
          let _671_v54;
          _671_v54 = _dafny.MultiSet.fromElements(p0, !(p0));
          let _672_v55;
          let _nw91 = Array((new BigNumber(27)).toNumber());
          _nw91[(_dafny.ZERO).toNumber()] = (_639_v31).Union(_639_v31);
          _nw91[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements((_this).f11);
          _nw91[(new BigNumber(2)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(3)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(4)).toNumber()] = (_639_v31).Intersect(_639_v31);
          _nw91[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(_641_i3, new BigNumber((_dafny.Seq.of(_641_i3)).length), _641_i3, (_this).f11, (_this).f11);
          _nw91[(new BigNumber(6)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.fromElements((_669_v52).dtor_cf33)).Difference(_639_v31);
          _nw91[(new BigNumber(8)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(9)).toNumber()] = ((p0) ? (_639_v31) : (_639_v31));
          _nw91[(new BigNumber(10)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(11)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(12)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(13)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(14)).toNumber()] = (_639_v31).Intersect(_639_v31);
          _nw91[(new BigNumber(15)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(16)).toNumber()] = (_639_v31).update(_module.__default.fm2(new BigNumber(888), new BigNumber((_670_v53).length), globalState), _module.__default.abs((_this).f11));
          _nw91[(new BigNumber(17)).toNumber()] = ((_dafny.MultiSet.fromElements(new BigNumber(58), (_this).f11)).update(_641_i3, _module.__default.abs(_641_i3))).Intersect(_639_v31);
          _nw91[(new BigNumber(18)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(19)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(20)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(21)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(22)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(23)).toNumber()] = ((_639_v31).update((_this).f11, _module.__default.abs(new BigNumber((_671_v54).cardinality())))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(436)), ((_673_i3) => function (_674_i6) {
            return _673_i3;
          })(_641_i3))));
          _nw91[(new BigNumber(24)).toNumber()] = _639_v31;
          _nw91[(new BigNumber(25)).toNumber()] = _dafny.MultiSet.fromElements(_641_i3, _641_i3);
          _nw91[(new BigNumber(26)).toNumber()] = _dafny.MultiSet.fromElements((_this).f11);
          _672_v55 = _nw91;
          _672_v55 = _672_v55;
          let _675_v56;
          let _nw92 = new _module.C0();
          _nw92.__ctor(_667_v50, ((_this).f11).minus(new BigNumber(-174)));
          _675_v56 = _nw92;
        } else {
          let _676___mcc_h7 = (_source11).cf25;
          let _677_cf25 = _676___mcc_h7;
          let _678_v57;
          let _nw93 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
          _678_v57 = _nw93;
          let _679_v58;
          _679_v58 = _dafny.MultiSet.fromElements((_this).f24);
          let _680_v59;
          let _nw94 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _680_v59 = _nw94;
          let _681_v60;
          _681_v60 = _dafny.Seq.of(_680_v59);
          let _682_v61;
          _682_v61 = _dafny.Map.Empty.slice().updateUnsafe(_679_v58,(_681_v60)[_module.__default.safeIndex((_this).f11, new BigNumber((_681_v60).length))]);
          let _index76 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_678_v57).length));
          (_678_v57)[_index76] = _682_v61;
          let _index77 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_678_v57).length));
          let _rhs120 = ((_682_v61).Merge(_682_v61)).Merge((_682_v61).update(_679_v58, _680_v59));
          let _rhs121 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_681_v60, _681_v60), _681_v60)).length);
          let _rhs122 = p0;
          let _rhs123 = _dafny.Seq.Concat(_640_v32, _dafny.Seq.Concat(_dafny.Seq.of((_this).f11), _640_v32));
          let _lhs99 = _678_v57;
          let _lhs100 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_678_v57).length));
          let _lhs101 = globalState;
          let _lhs102 = globalState;
          _lhs99[_lhs100] = _rhs120;
          _lhs101.f1 = _rhs121;
          _lhs102.f7 = _rhs122;
          _640_v32 = _rhs123;
          let _683_v62;
          let _nw95 = new _module.C1();
          _nw95.__ctor(p0, !((_641_i3).isLessThan(_module.__default.fm2((_this).f11, _641_i3, globalState))));
          _683_v62 = _nw95;
          let _684_v63;
          _684_v63 = _dafny.Seq.UnicodeFromString("fgsyppx");
          let _685_v64;
          _685_v64 = _dafny.MultiSet.fromElements(_684_v63, _dafny.Seq.UnicodeFromString("gg"));
          let _686_v65;
          _686_v65 = _module.D2.create_DC3(_685_v64);
          let _687_v66;
          _687_v66 = _dafny.Seq.of((_683_v62).f19, (_683_v62).fm13(_686_v65, _module.__default.fm0(globalState), (_683_v62).f18, (_this).f11, globalState));
          let _688_v67;
          _688_v67 = _module.D6.create_DC10(_687_v66);
          let _pat_let_tv30 = _687_v66;
          let _pat_let_tv31 = _683_v62;
          let _pat_let_tv32 = _683_v62;
          let _689_v68;
          let _nw96 = Array((new BigNumber(25)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = _688_v67;
          _nw96[(_dafny.ONE).toNumber()] = _module.D6.create_DC10(_687_v66);
          _nw96[(new BigNumber(2)).toNumber()] = _module.D6.create_DC10((_688_v67).dtor_cf13);
          _nw96[(new BigNumber(3)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(4)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(5)).toNumber()] = _module.D6.create_DC10(_687_v66);
          _nw96[(new BigNumber(6)).toNumber()] = _module.__default.fm33(p0, globalState);
          _nw96[(new BigNumber(7)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(8)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(9)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(10)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(11)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(12)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(13)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(14)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(15)).toNumber()] = function (_pat_let16_0) {
            return function (_690_dt__update__tmp_h1) {
              return function (_pat_let17_0) {
                return function (_691_dt__update_hcf13_h0) {
                  return _module.D6.create_DC10(_691_dt__update_hcf13_h0);
                }(_pat_let17_0);
              }(_pat_let_tv30);
            }(_pat_let16_0);
          }(_688_v67);
          _nw96[(new BigNumber(16)).toNumber()] = _module.D6.create_DC10(_687_v66);
          _nw96[(new BigNumber(17)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(18)).toNumber()] = _module.D6.create_DC10(_687_v66);
          _nw96[(new BigNumber(19)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(20)).toNumber()] = _module.D6.create_DC10(_dafny.Seq.of((_683_v62).f19, (_683_v62).f19));
          _nw96[(new BigNumber(21)).toNumber()] = function (_pat_let18_0) {
            return function (_692_dt__update__tmp_h2) {
              return function (_pat_let19_0) {
                return function (_693_dt__update_hcf13_h1) {
                  return _module.D6.create_DC10(_693_dt__update_hcf13_h1);
                }(_pat_let19_0);
              }(_dafny.Seq.of((_pat_let_tv31).f19, !((_pat_let_tv32).f18)));
            }(_pat_let18_0);
          }(_688_v67);
          _nw96[(new BigNumber(22)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(23)).toNumber()] = _688_v67;
          _nw96[(new BigNumber(24)).toNumber()] = _688_v67;
          _689_v68 = _nw96;
          let _index78 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_689_v68).length));
          (_689_v68)[_index78] = _module.D6.create_DC10(_687_v66);
          let _index79 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_689_v68).length));
          (_689_v68)[_index79] = _688_v67;
          let _694_v69;
          _694_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,_680_v59);
          let _695_v70;
          _695_v70 = _dafny.Map.Empty.slice().updateUnsafe((_694_v69).contains(!(!((_683_v62).f18))),(p0) && ((_683_v62).f19));
          let _696_v71;
          _696_v71 = _dafny.Set.fromElements((_683_v62).f19);
          _695_v70 = (_695_v70).update((new BigNumber(429)).isLessThan((_this).f11), (_696_v71).IsProperSubsetOf(_696_v71));
        }
        let _697_v72;
        _697_v72 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),_dafny.Seq.Concat(_640_v32, _640_v32));
        (globalState).f1 = new BigNumber(((((_697_v72).contains(p0)) ? ((_697_v72).get(p0)) : (_640_v32))).length);
        let _698_v73;
        let _nw97 = new _module.C1();
        _nw97.__ctor(p0, p0);
        _698_v73 = _nw97;
      }
      let _699_v74;
      _699_v74 = _dafny.Seq.of(p0);
      let _700_v75;
      let _nw98 = new _module.C0();
      _nw98.__ctor(_dafny.Seq.Concat(_699_v74, _699_v74), ((_this).f11).minus((_this).f11));
      _700_v75 = _nw98;
      let _701_v76;
      _701_v76 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,p0);
      let _702_v77;
      _702_v77 = _module.D8.create_DC16((((_701_v76).contains((_700_v75).f21)) ? ((_701_v76).get((_700_v75).f21)) : (p0)), (_700_v75).f21);
      let _pat_let_tv33 = p0;
      let _pat_let_tv34 = p0;
      if (function (_source12) {
        if (_source12.is_DC16) {
          let _703___mcc_h17 = (_source12).cf23;
          let _704___mcc_h18 = (_source12).cf24;
          let _705_cf24 = _704___mcc_h18;
          let _706_cf23 = _703___mcc_h17;
          return true;
        } else if (_source12.is_DC15) {
          let _707___mcc_h19 = (_source12).cf22;
          let _708_cf22 = _707___mcc_h19;
          return true;
        } else {
          let _709___mcc_h20 = (_source12).cf25;
          let _710_cf25 = _709___mcc_h20;
          return (_pat_let_tv33) === (false);
        }
      }(function (_pat_let20_0) {
        return function (_711_dt__update__tmp_h3) {
          return function (_pat_let21_0) {
            return function (_712_dt__update_hcf23_h0) {
              return _module.D8.create_DC16(_712_dt__update_hcf23_h0, (_711_dt__update__tmp_h3).dtor_cf24);
            }(_pat_let21_0);
          }(_pat_let_tv34);
        }(_pat_let20_0);
      }(_702_v77))) {
        let _713_v78;
        _713_v78 = _module.D2.create_DC5();
        let _source13 = _713_v78;
        if (_source13.is_DC4) {
          let _714___mcc_h8 = (_source13).cf6;
          let _715___mcc_h9 = (_source13).cf7;
          let _716___mcc_h10 = (_source13).cf8;
          let _717_cf8 = _716___mcc_h10;
          let _718_cf7 = _715___mcc_h9;
          let _719_cf6 = _714___mcc_h8;
          let _720_v79;
          let _init11 = ((_721_v77) => function (_722_i7) {
            return _721_v77;
          })(_702_v77);
          let _nw99 = Array((new BigNumber(26)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw99.length); _i0_11++) {
            _nw99[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _720_v79 = _nw99;
          let _723_v80;
          _723_v80 = _module.D12.create_DC25(_720_v79);
          let _pat_let_tv35 = _720_v79;
          let _rhs124 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_700_v75).f21, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length)), (_700_v75).f21);
          let _rhs125 = (function (_pat_let22_0) {
            return function (_724_dt__update__tmp_h4) {
              return function (_pat_let23_0) {
                return function (_725_dt__update_hcf37_h0) {
                  return _module.D12.create_DC25(_725_dt__update_hcf37_h0);
                }(_pat_let23_0);
              }(_pat_let_tv35);
            }(_pat_let22_0);
          }(_723_v80)).dtor_cf37;
          let _lhs103 = globalState;
          _lhs103.f1 = _rhs124;
          _720_v79 = _rhs125;
          let _726_v81;
          let _nw100 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
          _726_v81 = _nw100;
          let _727_v82;
          _727_v82 = _dafny.MultiSet.fromElements(p0, p0);
          let _index80 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_726_v81).length));
          (_726_v81)[_index80] = new BigNumber((_727_v82).cardinality());
          let _index81 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_726_v81).length));
          (_726_v81)[_index81] = (_this).f11;
          r1 = p1;
          let _728_v83;
          let _nw101 = Array((new BigNumber(10)).toNumber());
          _nw101[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_700_v75).f21, (_700_v75).f21, new BigNumber(737));
          _nw101[(_dafny.ONE).toNumber()] = _640_v32;
          _nw101[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_719_cf6);
          _nw101[(new BigNumber(3)).toNumber()] = _640_v32;
          _nw101[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_700_v75).f21), _dafny.Seq.of(_718_cf7));
          _nw101[(new BigNumber(5)).toNumber()] = _640_v32;
          _nw101[(new BigNumber(6)).toNumber()] = _640_v32;
          _nw101[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(296)), function (_729_i8) {
            return new BigNumber((_dafny.Seq.of((_this).f24)).length);
          });
          _nw101[(new BigNumber(8)).toNumber()] = _640_v32;
          _nw101[(new BigNumber(9)).toNumber()] = _640_v32;
          _728_v83 = _nw101;
          let _index82 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_728_v83).length));
          (_728_v83)[_index82] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(47)), ((_730_cf6) => function (_731_i9) {
            return _730_cf6;
          })(_719_cf6)), _module.__default.fm34(p0, _module.D6.create_DC11(new BigNumber(100), _718_cf7, (_700_v75).f21), globalState));
          let _index83 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_728_v83).length));
          (_728_v83)[_index83] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus((_726_v81)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_726_v81).length))])), _640_v32), _640_v32);
        } else if (_source13.is_DC5) {
          let _732_v84;
          _732_v84 = _dafny.Set.fromElements(p0, false);
          let _733_v85;
          let _nw102 = Array((new BigNumber(3)).toNumber());
          _nw102[(_dafny.ZERO).toNumber()] = _732_v84;
          _nw102[(_dafny.ONE).toNumber()] = (_732_v84).Difference(_dafny.Set.fromElements(p0));
          _nw102[(new BigNumber(2)).toNumber()] = (_732_v84).Union(_732_v84);
          _733_v85 = _nw102;
          let _index84 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_733_v85).length));
          (_733_v85)[_index84] = _732_v84;
          let _734_v86;
          _734_v86 = _dafny.Seq.UnicodeFromString("cym");
          let _735_v87;
          _735_v87 = _module.D6.create_DC11((_this).f11, (_700_v75).f21, (_dafny.ZERO).minus(new BigNumber((_734_v86).length)));
          let _index85 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_733_v85).length));
          let _rhs126 = _module.__default.fm34((new BigNumber((_dafny.Seq.UnicodeFromString("c")).length)).isLessThan((_700_v75).f21), _735_v87, globalState);
          let _rhs127 = _732_v84;
          let _lhs104 = _733_v85;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(418), new BigNumber((_733_v85).length));
          _640_v32 = _rhs126;
          _lhs104[_lhs105] = _rhs127;
          let _index86 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length));
          (p1)[_index86] = p0;
          let _index87 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length));
          (p1)[_index87] = !(p0);
          (globalState).f1 = _module.__default.safeDivisionInt((new BigNumber((_640_v32).length)).multipliedBy((_dafny.ZERO).minus((_this).f11)), (new BigNumber(48)).multipliedBy((_700_v75).f21));
          let _736_v88;
          _736_v88 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))],p0);
          let _737_v89;
          let _nw103 = Array((new BigNumber(23)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
          _nw103[(_dafny.ONE).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
          _nw103[(new BigNumber(2)).toNumber()] = false;
          _nw103[(new BigNumber(3)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
          _nw103[(new BigNumber(4)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
          _nw103[(new BigNumber(5)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
          _nw103[(new BigNumber(6)).toNumber()] = p0;
          _nw103[(new BigNumber(7)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
          _nw103[(new BigNumber(8)).toNumber()] = p0;
          _nw103[(new BigNumber(9)).toNumber()] = !(p0) || (p0);
          _nw103[(new BigNumber(10)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
          _nw103[(new BigNumber(11)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
          _nw103[(new BigNumber(12)).toNumber()] = (((p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))]) ? (p0) : (!(p0)));
          _nw103[(new BigNumber(13)).toNumber()] = (_700_v75).fm21((p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))], (_700_v75).f21, globalState);
          _nw103[(new BigNumber(14)).toNumber()] = (p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))];
          _nw103[(new BigNumber(15)).toNumber()] = false;
          _nw103[(new BigNumber(16)).toNumber()] = false;
          _nw103[(new BigNumber(17)).toNumber()] = (p0) && ((p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))]);
          _nw103[(new BigNumber(18)).toNumber()] = ((p0) ? (true) : ((p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))]));
          _nw103[(new BigNumber(19)).toNumber()] = p0;
          _nw103[(new BigNumber(20)).toNumber()] = (_700_v75).fm21(p0, (_735_v87).dtor_cf14, globalState);
          _nw103[(new BigNumber(21)).toNumber()] = (((_736_v88).contains(p0)) ? ((_736_v88).get(p0)) : (!((p1)[_module.__default.safeIndex(new BigNumber(182), new BigNumber((p1).length))])));
          _nw103[(new BigNumber(22)).toNumber()] = p0;
          _737_v89 = _nw103;
          r1 = _737_v89;
        } else {
          let _738___mcc_h11 = (_source13).cf5;
          let _739_cf5 = _738___mcc_h11;
          let _740_v90;
          _740_v90 = _dafny.Seq.of(_dafny.Seq.of(p0, true), _dafny.Seq.of(p0));
          let _741_v91;
          _741_v91 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
          let _742_v92;
          let _nw104 = new _module.C0();
          _nw104.__ctor(_dafny.Seq.update((_740_v90)[_module.__default.safeIndex((_this).f11, new BigNumber((_740_v90).length))], _module.__default.safeIndex(new BigNumber((_741_v91).length), new BigNumber(((_740_v90)[_module.__default.safeIndex((_this).f11, new BigNumber((_740_v90).length))]).length)), p0), (_this).f11);
          _742_v92 = _nw104;
          let _743_v93;
          _743_v93 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm1(globalState));
          let _744_v94;
          _744_v94 = _dafny.Seq.UnicodeFromString("qxgefovt");
          let _745_v95;
          _745_v95 = _dafny.Map.Empty.slice().updateUnsafe((((_743_v93).contains(p0)) ? ((_743_v93).get(p0)) : (_744_v94)),(_700_v75).f21);
          let _746_v96;
          _746_v96 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_700_v75).f21);
          _745_v95 = (_745_v95).update(_dafny.Seq.Concat(_744_v94, _744_v94), new BigNumber(((_746_v96).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,(_700_v75).f21))).length));
          let _747_v97;
          _747_v97 = _dafny.MultiSet.fromElements(p0, p0);
          (globalState).f1 = (_module.__default.safeDivisionInt((_700_v75).f21, (_this).f11)).multipliedBy(((((_747_v97).contains(p0)) ? ((_747_v97).get(p0)) : ((_700_v75).f21))).multipliedBy((_742_v92).f21));
          let _748_v98;
          let _nw105 = new _module.C0();
          _nw105.__ctor(_dafny.Seq.Concat(_dafny.Seq.of(p0), _dafny.Seq.update((_700_v75).f20, _module.__default.safeIndex((_700_v75).f21, new BigNumber(((_700_v75).f20).length)), p0)), new BigNumber(((_746_v96).Merge(_746_v96)).length));
          _748_v98 = _nw105;
        }
        (globalState).f1 = (_dafny.ZERO).minus((_700_v75).f21);
        let _749_v99;
        _749_v99 = _dafny.Seq.UnicodeFromString("uesixv");
        let _rhs128 = (_700_v75).f21;
        let _rhs129 = _749_v99;
        let _lhs106 = globalState;
        _lhs106.f1 = _rhs128;
        _749_v99 = _rhs129;
        (globalState).f1 = ((_this).f11).plus((_this).f11);
        let _rhs130 = ((_700_v75).f21).isLessThanOrEqualTo(((_700_v75).f21).multipliedBy((_700_v75).f21));
        let _rhs131 = !(p0);
        let _lhs107 = globalState;
        let _lhs108 = globalState;
        _lhs107.f7 = _rhs130;
        _lhs108.f6 = _rhs131;
      } else {
        (globalState).f7 = p0;
        let _750_v100;
        let _nw106 = Array((new BigNumber(18)).toNumber()).fill([]);
        _750_v100 = _nw106;
        let _751_v101;
        let _nw107 = Array((new BigNumber(20)).toNumber()).fill([]);
        _751_v101 = _nw107;
        let _752_v102;
        _752_v102 = _751_v101;
        let _index88 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_750_v100).length));
        (_750_v100)[_index88] = _752_v102;
        let _index89 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_750_v100).length));
        (_750_v100)[_index89] = _752_v102;
        let _753_v103;
        let _nw108 = new _module.C0();
        _nw108.__ctor(_dafny.Seq.Concat(_699_v74, _699_v74), (_700_v75).f21);
        _753_v103 = _nw108;
        let _754_v104;
        _754_v104 = _module.D7.create_DC12(_dafny.Seq.update(_640_v32, _module.__default.safeIndex((_753_v103).f21, new BigNumber((_640_v32).length)), (_753_v103).f21));
        let _source14 = _754_v104;
        if (_source14.is_DC13) {
          let _755___mcc_h12 = (_source14).cf18;
          let _756___mcc_h13 = (_source14).cf19;
          let _757___mcc_h14 = (_source14).cf20;
          let _758_cf20 = _757___mcc_h14;
          let _759_cf19 = _756___mcc_h13;
          let _760_cf18 = _755___mcc_h12;
          let _761_v105;
          let _762_v106;
          let _763_v107;
          let _764_v108;
          let _out20;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector7 = (_this).m8(globalState);
          _out20 = _outcollector7[0];
          _out21 = _outcollector7[1];
          _out22 = _outcollector7[2];
          _out23 = _outcollector7[3];
          _761_v105 = _out20;
          _762_v106 = _out21;
          _763_v107 = _out22;
          _764_v108 = _out23;
          let _765_v109;
          _765_v109 = _dafny.Seq.UnicodeFromString("mbcbg");
          let _rhs132 = _dafny.Seq.of(_dafny.Seq.IsPrefixOf(_765_v109, _765_v109));
          let _rhs133 = p1;
          _699_v74 = _rhs132;
          r1 = _rhs133;
          (globalState).f6 = p0;
          _763_v107 = !(_761_v105);
        } else if (_source14.is_DC12) {
          let _766___mcc_h15 = (_source14).cf17;
          let _767_cf17 = _766___mcc_h15;
          (globalState).f7 = !(p0);
          let _768_v110;
          let _nw109 = new _module.C0();
          _nw109.__ctor((_753_v103).f20, (_700_v75).f21);
          _768_v110 = _nw109;
          let _769_v111;
          _769_v111 = _dafny.Seq.UnicodeFromString("ckt");
          _769_v111 = _769_v111;
          let _770_v112;
          _770_v112 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f24));
          let _771_v113;
          _771_v113 = _769_v111;
          let _772_v114;
          _772_v114 = _dafny.Map.Empty.slice().updateUnsafe((_770_v112)[_module.__default.safeIndex(_module.__default.fm2((_753_v103).f21, new BigNumber((_769_v111).length), globalState), new BigNumber((_770_v112).length))],_771_v113);
          let _773_v115;
          _773_v115 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f24);
          _772_v114 = (_772_v114).update(_773_v115, _771_v113);
        } else {
          let _774___mcc_h16 = (_source14).cf21;
          let _775_cf21 = _774___mcc_h16;
          let _776_v116;
          _776_v116 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          (globalState).f7 = (_753_v103).fm21((((_776_v116).contains(p0)) ? ((_776_v116).get(p0)) : (p0)), (_700_v75).f21, globalState);
          (globalState).f7 = p0;
          let _777_v117;
          _777_v117 = _dafny.Set.fromElements(new BigNumber(324), (_700_v75).f21, ((_700_v75).f21).plus((_753_v103).f21));
          let _rhs134 = new BigNumber((_777_v117).length);
          let _rhs135 = _dafny.areEqual(_699_v74, _module.__default.fm35((_753_v103).f21, (_700_v75).f21, _module.__default.fm36(p0, p0, (_this).f11, globalState), new BigNumber(60), globalState));
          let _lhs109 = globalState;
          let _lhs110 = globalState;
          _lhs109.f1 = _rhs134;
          _lhs110.f7 = _rhs135;
          _640_v32 = _640_v32;
        }
        if (((p0) ? (false) : (p0))) {
          _701_v76 = (_701_v76).update(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(535)), ((_778_v103) => function (_779_i10) {
            return (_778_v103).f21;
          })(_753_v103)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(85)), ((_780_v75) => function (_781_i11) {
            return new BigNumber((function () {
              let _coll44 = new _dafny.Map();
              for (const _compr_44 of _dafny.IntegerRange(new BigNumber(-190), new BigNumber(507))) {
                let _782_v118 = _compr_44;
                if (((new BigNumber(-190)).isLessThanOrEqualTo(_782_v118)) && ((_782_v118).isLessThan(new BigNumber(507)))) {
                  _coll44.push([(_782_v118).minus((_780_v75).f21),(_dafny.ZERO).minus((_780_v75).f21)]);
                }
              }
              return _coll44;
            }()).length);
          })(_700_v75)))).length), p0);
          (globalState).f1 = (_753_v103).f21;
          let _783_v119;
          _783_v119 = _dafny.Seq.UnicodeFromString("v");
          let _784_v120;
          _784_v120 = _dafny.Seq.of(_783_v119);
          (globalState).f7 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_783_v119, _dafny.Seq.UnicodeFromString("lviqvc")), _dafny.Seq.Concat(_783_v119, (_784_v120)[_module.__default.safeIndex((_700_v75).f21, new BigNumber((_784_v120).length))]));
          (globalState).f1 = new BigNumber((_module.__default.fm1(globalState)).length);
          let _785_v121;
          _785_v121 = _dafny.MultiSet.fromElements(_dafny.Seq.IsProperPrefixOf(_783_v119, _783_v119), !_dafny.Seq.contains(_783_v119, (_this).f24));
          let _rhs136 = ((_this).fm3(new BigNumber(983), true, p0, globalState)).minus((_700_v75).f21);
          let _rhs137 = (((p0) ? (_module.__default.fm37(globalState)) : (_dafny.MultiSet.FromArray((_700_v75).f20)))).update(((false) ? (p0) : (p0)), _module.__default.abs((_753_v103).f21));
          let _rhs138 = !(p0);
          let _lhs111 = globalState;
          let _lhs112 = globalState;
          _lhs111.f1 = _rhs136;
          _785_v121 = _rhs137;
          _lhs112.f7 = _rhs138;
        } else {
          (globalState).f6 = p0;
          (globalState).f1 = ((p0) ? ((_dafny.ZERO).minus((_this).f11)) : ((_753_v103).f21));
          let _786_v122;
          let _nw110 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _786_v122 = _nw110;
          _786_v122 = _786_v122;
          let _787_v123;
          let _init12 = function (_788_i12) {
            return _dafny.Seq.UnicodeFromString("pswwmeq");
          };
          let _nw111 = Array((new BigNumber(7)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw111.length); _i0_12++) {
            _nw111[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _787_v123 = _nw111;
          let _rhs139 = _787_v123;
          let _rhs140 = (_753_v103).f21;
          let _lhs113 = globalState;
          _787_v123 = _rhs139;
          _lhs113.f1 = _rhs140;
          let _789_v124;
          _789_v124 = _dafny.MultiSet.fromElements(_699_v74);
          let _index90 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_786_v122).length));
          (_786_v122)[_index90] = new BigNumber((((p0) ? (_789_v124) : (_789_v124))).cardinality());
          let _index91 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_786_v122).length));
          (_786_v122)[_index91] = _module.__default.fm2((_dafny.ZERO).minus((_this).fm3((_700_v75).f21, !(p0), p0, globalState)), (_700_v75).f21, globalState);
        }
      }
      let _790_v125;
      _790_v125 = _module.D6.create_DC11(new BigNumber(85), (_700_v75).f21, (_700_v75).f21);
      if (_dafny.areEqual(_module.__default.fm34(p0, _790_v125, globalState), _dafny.Seq.of((_700_v75).f21, (_this).f11))) {
        (globalState).f1 = (_700_v75).f21;
        if (false) {
          let _791_v126;
          _791_v126 = _dafny.Seq.UnicodeFromString("jkdpcyeue");
          let _792_v127;
          _792_v127 = _dafny.Map.Empty.slice().updateUnsafe(p0,_791_v126);
          (globalState).f1 = (_module.__default.fm38(_792_v127, p0, false, globalState)).dtor_cf0;
          let _793_v128;
          let _init13 = ((_794_v126, _795_p0) => function (_796_i13) {
            return (_796_i13).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_794_v126).length),_795_p0)).length));
          })(_791_v126, p0);
          let _nw112 = Array((new BigNumber(2)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw112.length); _i0_13++) {
            _nw112[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _793_v128 = _nw112;
          let _index92 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_793_v128).length));
          (_793_v128)[_index92] = (_700_v75).f21;
          let _index93 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_793_v128).length));
          (_793_v128)[_index93] = (_700_v75).f21;
          let _index94 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_793_v128).length));
          let _index95 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_793_v128).length));
          let _rhs141 = (_700_v75).f21;
          let _rhs142 = ((_700_v75).f21).multipliedBy((_700_v75).f21);
          let _lhs114 = _793_v128;
          let _lhs115 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_793_v128).length));
          let _lhs116 = _793_v128;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_793_v128).length));
          _lhs114[_lhs115] = _rhs141;
          _lhs116[_lhs117] = _rhs142;
          let _rhs143 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_793_v128)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_793_v128).length))]));
          let _rhs144 = p0;
          let _rhs145 = _dafny.Seq.IsPrefixOf(_640_v32, _640_v32);
          let _rhs146 = ((!(false)) ? (p0) : (true));
          let _lhs118 = globalState;
          let _lhs119 = globalState;
          let _lhs120 = globalState;
          let _lhs121 = globalState;
          _lhs118.f1 = _rhs143;
          _lhs119.f7 = _rhs144;
          _lhs120.f7 = _rhs145;
          _lhs121.f6 = _rhs146;
          (globalState).f6 = p0;
          (globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_700_v75).f21, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(35)), function (_797_i14) {
            return (_this).f24;
          }), _dafny.Seq.UnicodeFromString("w"))).length)));
        } else {
          let _798_v129;
          let _799_v130;
          let _800_v131;
          let _801_v132;
          let _out24;
          let _out25;
          let _out26;
          let _out27;
          let _outcollector8 = (_this).m8(globalState);
          _out24 = _outcollector8[0];
          _out25 = _outcollector8[1];
          _out26 = _outcollector8[2];
          _out27 = _outcollector8[3];
          _798_v129 = _out24;
          _799_v130 = _out25;
          _800_v131 = _out26;
          _801_v132 = _out27;
          (globalState).f6 = !((_700_v75).fm21(false, _module.__default.safeDivisionInt(_801_v132, (_700_v75).f21), globalState));
          r1 = p1;
          let _802_v133;
          _802_v133 = _module.D11.create_DC23(_799_v130, (_this).f11, (_700_v75).f21);
          (globalState).f7 = (_802_v133).dtor_cf31;
          (globalState).f1 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt((_700_v75).f21, (_dafny.ZERO).minus((_700_v75).f21))).plus((_this).f11));
        }
        let _803_v134;
        _803_v134 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24);
        let _804_v135;
        _804_v135 = _dafny.Seq.UnicodeFromString("vds");
        _701_v76 = (_701_v76).update((((_803_v134).contains((_this).f24)) ? ((_803_v134).get((_this).f24)) : ((_this).fm3((_700_v75).f21, p0, p0, globalState))), !_dafny.areEqual(_804_v135, _804_v135));
        let _805_v136;
        let _nw113 = new _module.C0();
        _nw113.__ctor(_699_v74, (_700_v75).f21);
        _805_v136 = _nw113;
        let _806_v137;
        let _nw114 = new _module.C0();
        _nw114.__ctor(((p0) ? ((_700_v75).f20) : (_699_v74)), (_805_v136).f21);
        _806_v137 = _nw114;
      } else {
        let _807_v138;
        let _nw115 = new _module.C0();
        _nw115.__ctor((_700_v75).f20, _module.__default.safeDivisionInt(_module.__default.fm2((_this).f11, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(451)), ((_808_v75) => function (_809_i15) {
          return (_808_v75).f21;
        })(_700_v75))).length), globalState), (((_639_v31).contains((_this).f11)) ? ((_639_v31).get((_this).f11)) : ((_700_v75).f21))));
        _807_v138 = _nw115;
        (globalState).f6 = ((p0) ? ((_700_v75).fm21(p0, (_700_v75).f21, globalState)) : (p0));
        let _810_v139;
        let _nw116 = new _module.C0();
        _nw116.__ctor((_700_v75).f20, (_700_v75).f21);
        _810_v139 = _nw116;
        (globalState).f1 = (_807_v138).f21;
        let _811_v140;
        _811_v140 = _dafny.Map.Empty.slice().updateUnsafe(((_807_v138).f20)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).fm3((_dafny.ZERO).minus((_810_v139).f21), p0, p0, globalState)), new BigNumber(((_807_v138).f20).length))],p0);
        let _812_v141;
        _812_v141 = _dafny.MultiSet.fromElements(p0);
        let _813_v142;
        _813_v142 = _dafny.Seq.UnicodeFromString("pag");
        let _814_v143;
        let _nw117 = Array((new BigNumber(23)).toNumber());
        _nw117[(_dafny.ZERO).toNumber()] = (_807_v138).f21;
        _nw117[(_dafny.ONE).toNumber()] = (_810_v139).f21;
        _nw117[(new BigNumber(2)).toNumber()] = new BigNumber(((_807_v138).f20).length);
        _nw117[(new BigNumber(3)).toNumber()] = (_700_v75).f21;
        _nw117[(new BigNumber(4)).toNumber()] = new BigNumber(-913);
        _nw117[(new BigNumber(5)).toNumber()] = (_700_v75).f21;
        _nw117[(new BigNumber(6)).toNumber()] = (_810_v139).f21;
        _nw117[(new BigNumber(7)).toNumber()] = (_700_v75).f21;
        _nw117[(new BigNumber(8)).toNumber()] = (_807_v138).f21;
        _nw117[(new BigNumber(9)).toNumber()] = (_807_v138).f21;
        _nw117[(new BigNumber(10)).toNumber()] = (_700_v75).f21;
        _nw117[(new BigNumber(11)).toNumber()] = new BigNumber((_639_v31).cardinality());
        _nw117[(new BigNumber(12)).toNumber()] = (_807_v138).f21;
        _nw117[(new BigNumber(13)).toNumber()] = (_this).f11;
        _nw117[(new BigNumber(14)).toNumber()] = (_807_v138).f21;
        _nw117[(new BigNumber(15)).toNumber()] = (_700_v75).f21;
        _nw117[(new BigNumber(16)).toNumber()] = new BigNumber((_813_v142).length);
        _nw117[(new BigNumber(17)).toNumber()] = new BigNumber((_640_v32).length);
        _nw117[(new BigNumber(18)).toNumber()] = new BigNumber((_701_v76).length);
        _nw117[(new BigNumber(19)).toNumber()] = new BigNumber((_813_v142).length);
        _nw117[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("sowrbha")).length);
        _nw117[(new BigNumber(21)).toNumber()] = (_700_v75).f21;
        _nw117[(new BigNumber(22)).toNumber()] = (_this).f11;
        _814_v143 = _nw117;
        let _815_v144;
        _815_v144 = _dafny.Seq.of(_814_v143);
        let _816_v145;
        _816_v145 = _dafny.Map.Empty.slice().updateUnsafe(_640_v32,p0);
        let _817_v146;
        _817_v146 = _dafny.Map.Empty.slice().updateUnsafe((((_639_v31).contains(new BigNumber(26))) ? ((_639_v31).get(new BigNumber(26))) : ((_700_v75).f21)),_module.__default.fm39(new BigNumber((_813_v142).length), p0, new BigNumber((_640_v32).length), globalState));
        let _818_v147;
        let _nw118 = Array((new BigNumber(27)).toNumber());
        _nw118[(_dafny.ZERO).toNumber()] = new BigNumber(660);
        _nw118[(_dafny.ONE).toNumber()] = ((true) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_811_v140)).length)) : ((_810_v139).f21));
        _nw118[(new BigNumber(2)).toNumber()] = new BigNumber((_699_v74).length);
        _nw118[(new BigNumber(3)).toNumber()] = (_700_v75).f21;
        _nw118[(new BigNumber(4)).toNumber()] = (_807_v138).f21;
        _nw118[(new BigNumber(5)).toNumber()] = (_810_v139).f21;
        _nw118[(new BigNumber(6)).toNumber()] = ((_807_v138).f21).multipliedBy((_700_v75).f21);
        _nw118[(new BigNumber(7)).toNumber()] = (_700_v75).f21;
        _nw118[(new BigNumber(8)).toNumber()] = _module.__default.fm2((_this).f11, (_700_v75).f21, globalState);
        _nw118[(new BigNumber(9)).toNumber()] = ((_807_v138).f21).minus(new BigNumber((_812_v141).cardinality()));
        _nw118[(new BigNumber(10)).toNumber()] = new BigNumber((_815_v144).length);
        _nw118[(new BigNumber(11)).toNumber()] = (_this).f11;
        _nw118[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_700_v75).f21);
        _nw118[(new BigNumber(13)).toNumber()] = (_this).f11;
        _nw118[(new BigNumber(14)).toNumber()] = new BigNumber(98);
        _nw118[(new BigNumber(15)).toNumber()] = (_700_v75).f21;
        _nw118[(new BigNumber(16)).toNumber()] = new BigNumber((_816_v145).length);
        _nw118[(new BigNumber(17)).toNumber()] = ((_700_v75).f21).plus((_700_v75).f21);
        _nw118[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt((_807_v138).f21, new BigNumber(((_dafny.MultiSet.fromElements(p0)).update(p0, _module.__default.abs((_700_v75).f21))).cardinality()));
        _nw118[(new BigNumber(19)).toNumber()] = (_this).f11;
        _nw118[(new BigNumber(20)).toNumber()] = (_810_v139).f21;
        _nw118[(new BigNumber(21)).toNumber()] = new BigNumber((_817_v146).length);
        _nw118[(new BigNumber(22)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(983), new BigNumber((_813_v142).length));
        _nw118[(new BigNumber(23)).toNumber()] = (_807_v138).f21;
        _nw118[(new BigNumber(24)).toNumber()] = (_810_v139).f21;
        _nw118[(new BigNumber(25)).toNumber()] = (_this).f11;
        _nw118[(new BigNumber(26)).toNumber()] = ((_807_v138).f21).multipliedBy((_807_v138).f21);
        _818_v147 = _nw118;
        let _819_v148;
        _819_v148 = _dafny.Map.Empty.slice().updateUnsafe(p1,_814_v143);
        let _820_v149;
        _820_v149 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm40((_807_v138).f20, (_this).f24, (_this).f11, globalState)).length),_818_v147);
        _814_v143 = (((_819_v148).contains(p1)) ? ((_819_v148).get(p1)) : ((((_820_v149).contains((_810_v139).f21)) ? ((_820_v149).get((_810_v139).f21)) : (_814_v143))));
      }
      let _821_v150;
      _821_v150 = _dafny.Seq.UnicodeFromString("tbi");
      let _822_v151;
      let _nw119 = new _module.C1();
      _nw119.__ctor(_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-580)), function (_823_i16) {
        return (_this).f24;
      }), _821_v150), (p0) && (p0));
      _822_v151 = _nw119;
      let _824_v152;
      _824_v152 = _dafny.Set.fromElements(p0);
      let _825_v153;
      _825_v153 = _module.D7.create_DC13((_824_v152).Union(_dafny.Set.fromElements(p0)), new BigNumber((_639_v31).cardinality()), new BigNumber(263));
      r0 = _825_v153;
      r1 = p1;
      return [r0, r1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f12 = _dafny.Set.Empty;
      this._f11 = _dafny.ZERO;
      this.f23 = _dafny.ZERO;
      this._f22 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f22, f23, f11, f12) {
      let _this = this;
      (_this)._f22 = f22;
      (_this).f23 = f23;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      return;
    }
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this).f11).plus((_this).f11);
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(!(!(!((_this.f23).isLessThanOrEqualTo(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(true, true, false), _dafny.Seq.of(!(!(true)), true, false, true)))).cardinality()))))));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source15 = _module.D7.create_DC14(_module.D7.create_DC14(_module.D7.create_DC12(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-233)), function (_826_i0) {
  return (_this).f11;
}))));
      if (_source15.is_DC13) {
        let _827___mcc_h0 = (_source15).cf18;
        let _828___mcc_h1 = (_source15).cf19;
        let _829___mcc_h2 = (_source15).cf20;
        let _830_cf20 = _829___mcc_h2;
        let _831_cf19 = _828___mcc_h1;
        let _832_cf18 = _827___mcc_h0;
        return !(true);
      } else if (_source15.is_DC12) {
        let _833___mcc_h3 = (_source15).cf17;
        let _834_cf17 = _833___mcc_h3;
        return !((_this).f11).isEqualTo((_this).f11);
      } else {
        let _835___mcc_h4 = (_source15).cf21;
        let _836_cf21 = _835___mcc_h4;
        return (true) && (false);
      }
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(694);
    };
    fm25(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(!((_module.D0.create_DC1(false, new BigNumber(444), true)).dtor_cf1),!(false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge((_module.D8.create_DC15(_dafny.Map.Empty.slice().updateUnsafe(false,true))).dtor_cf22));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let _837_v0;
      _837_v0 = _dafny.Seq.UnicodeFromString("njxqrukhg");
      (globalState).f6 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xihlij"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(598)), function (_838_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })), _837_v0);
      let _839_v1;
      _839_v1 = false;
      if (_839_v1) {
        let _840_v2;
        _840_v2 = _dafny.Seq.of((_this).f11);
        let _rhs147 = !(_839_v1);
        let _rhs148 = ((false) ? ((_840_v2)[_module.__default.safeIndex(new BigNumber((_837_v0).length), new BigNumber((_840_v2).length))]) : ((_this).f11));
        let _rhs149 = ((_this).f11).isLessThanOrEqualTo(((_this).f22).minus(new BigNumber((_837_v0).length)));
        let _lhs122 = globalState;
        let _lhs123 = _this;
        _lhs122.f7 = _rhs147;
        _lhs123.f23 = _rhs148;
        _839_v1 = _rhs149;
        let _841_v3;
        _841_v3 = _dafny.MultiSet.fromElements(_839_v1, _839_v1);
        let _842_v4;
        _842_v4 = _dafny.Seq.of(true, false, _839_v1);
        let _843_v5;
        _843_v5 = _module.D0.create_DC0((_this).f11);
        let _844_v6;
        _844_v6 = _dafny.Seq.of(_840_v2);
        let _845_v7;
        let _nw120 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _845_v7 = _nw120;
        let _846_v8;
        _846_v8 = _dafny.Map.Empty.slice().updateUnsafe(_845_v7,_839_v1);
        let _847_v9;
        let _nw121 = Array((new BigNumber(21)).toNumber());
        _nw121[(_dafny.ZERO).toNumber()] = (_841_v3).IsSubsetOf(_841_v3);
        _nw121[(_dafny.ONE).toNumber()] = _839_v1;
        _nw121[(new BigNumber(2)).toNumber()] = _839_v1;
        _nw121[(new BigNumber(3)).toNumber()] = (_this).fm5(_module.D0.create_DC0(new BigNumber(-65)), _842_v4, new BigNumber(533), p0, globalState);
        _nw121[(new BigNumber(4)).toNumber()] = _839_v1;
        _nw121[(new BigNumber(5)).toNumber()] = !((_this).fm6(_839_v1, _this.f23, globalState)).isEqualTo((_this).f22);
        _nw121[(new BigNumber(6)).toNumber()] = ((_839_v1) ? (false) : (!(_839_v1)));
        _nw121[(new BigNumber(7)).toNumber()] = _839_v1;
        _nw121[(new BigNumber(8)).toNumber()] = _839_v1;
        _nw121[(new BigNumber(9)).toNumber()] = _dafny.areEqual(_dafny.Seq.UnicodeFromString("qkcgskmp"), _dafny.Seq.UnicodeFromString("ju"));
        _nw121[(new BigNumber(10)).toNumber()] = (_this).fm5(_843_v5, _842_v4, (_this).f11, new _dafny.CodePoint('o'.codePointAt(0)), globalState);
        _nw121[(new BigNumber(11)).toNumber()] = _dafny.Seq.IsPrefixOf(_844_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), ((_848_v2) => function (_849_i1) {
          return _848_v2;
        })(_840_v2)));
        _nw121[(new BigNumber(12)).toNumber()] = (_this.f23).isLessThanOrEqualTo((_this).f11);
        _nw121[(new BigNumber(13)).toNumber()] = (_839_v1) && (_839_v1);
        _nw121[(new BigNumber(14)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_840_v2, _dafny.Seq.update(_840_v2, _module.__default.safeIndex((_this).f11, new BigNumber((_840_v2).length)), new BigNumber(((_844_v6)[_module.__default.safeIndex(_this.f23, new BigNumber((_844_v6).length))]).length)));
        _nw121[(new BigNumber(15)).toNumber()] = !((((_846_v8).contains(_845_v7)) ? ((_846_v8).get(_845_v7)) : (_839_v1))) || (_839_v1);
        _nw121[(new BigNumber(16)).toNumber()] = (_842_v4)[_module.__default.safeIndex((_this).f22, new BigNumber((_842_v4).length))];
        _nw121[(new BigNumber(17)).toNumber()] = !(false) || (true);
        _nw121[(new BigNumber(18)).toNumber()] = (_this.f23).isLessThanOrEqualTo((_this).f22);
        _nw121[(new BigNumber(19)).toNumber()] = _839_v1;
        _nw121[(new BigNumber(20)).toNumber()] = _839_v1;
        _847_v9 = _nw121;
        let _850_v10;
        _850_v10 = _dafny.MultiSet.fromElements(new BigNumber(145), new BigNumber(((_this).f12).length));
        let _851_v11;
        _851_v11 = _dafny.Map.Empty.slice().updateUnsafe(_839_v1,(_this).f11);
        let _852_v12;
        _852_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_851_v11).length),new BigNumber((_837_v0).length));
        let _853_v13;
        _853_v13 = _dafny.Map.Empty.slice().updateUnsafe(_850_v10,_852_v12);
        let _index96 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_847_v9).length));
        (_847_v9)[_index96] = (_this).fm4(new BigNumber((_dafny.Seq.UnicodeFromString("oxsuipcwe")).length), _853_v13, _this.f23, _839_v1, globalState);
        let _index97 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_847_v9).length));
        (_847_v9)[_index97] = _839_v1;
        let _854_v14;
        let _nw122 = new _module.C0();
        _nw122.__ctor(_842_v4, (((_851_v11).contains((_847_v9)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_847_v9).length))])) ? ((_851_v11).get((_847_v9)[_module.__default.safeIndex(new BigNumber(939), new BigNumber((_847_v9).length))])) : ((_dafny.ZERO).minus((_this).f22))));
        _854_v14 = _nw122;
        _854_v14 = _854_v14;
        let _855_v15;
        _855_v15 = _837_v0;
        let _rhs150 = _839_v1;
        let _rhs151 = _dafny.Seq.UnicodeFromString("nb");
        let _lhs124 = globalState;
        _lhs124.f7 = _rhs150;
        _855_v15 = _rhs151;
        let _856_v16;
        _856_v16 = _dafny.Map.Empty.slice().updateUnsafe(_841_v3,_module.__default.fm26(globalState));
        _851_v11 = (_851_v11).update((new BigNumber((_851_v11).length)).isLessThan((_this).f11), new BigNumber((_856_v16).length));
      } else {
        let _857_v17;
        _857_v17 = _dafny.Seq.of(_this.f23);
        let _858_v18;
        _858_v18 = _module.D8.create_DC16(_839_v1, (_dafny.ZERO).minus(_this.f23));
        let _859_v19;
        _859_v19 = _module.D8.create_DC17(_858_v18);
        let _860_v20;
        _860_v20 = _dafny.Map.Empty.slice().updateUnsafe(_857_v17,_858_v18);
        let _861_v21;
        _861_v21 = _module.D8.create_DC17((((_860_v20).contains(_dafny.Seq.of(_this.f23))) ? ((_860_v20).get(_dafny.Seq.of(_this.f23))) : (_858_v18)));
        _861_v21 = _861_v21;
        let _862_v22;
        _862_v22 = _module.D0.create_DC0(_this.f23);
        let _863_v23;
        _863_v23 = _dafny.Seq.of(_839_v1);
        let _864_v24;
        _864_v24 = _dafny.MultiSet.fromElements((_this).f11);
        let _865_v25;
        _865_v25 = _dafny.MultiSet.fromElements(_864_v24);
        let _866_v26;
        _866_v26 = _dafny.Seq.of(_863_v23);
        let _867_v27;
        let _nw123 = Array((new BigNumber(18)).toNumber());
        _nw123[(_dafny.ZERO).toNumber()] = !(_839_v1) || (_839_v1);
        _nw123[(_dafny.ONE).toNumber()] = !(_839_v1);
        _nw123[(new BigNumber(2)).toNumber()] = _839_v1;
        _nw123[(new BigNumber(3)).toNumber()] = true;
        _nw123[(new BigNumber(4)).toNumber()] = _839_v1;
        _nw123[(new BigNumber(5)).toNumber()] = (_this).fm5(_862_v22, _863_v23, _this.f23, p0, globalState);
        _nw123[(new BigNumber(6)).toNumber()] = (_this.f23).isLessThan(_this.f23);
        _nw123[(new BigNumber(7)).toNumber()] = !(!(_839_v1));
        _nw123[(new BigNumber(8)).toNumber()] = _839_v1;
        _nw123[(new BigNumber(9)).toNumber()] = _839_v1;
        _nw123[(new BigNumber(10)).toNumber()] = !((_dafny.Set.fromElements(new BigNumber((_864_v24).cardinality()))).IsProperSubsetOf(_dafny.Set.fromElements((_this).f11)));
        _nw123[(new BigNumber(11)).toNumber()] = (_865_v25).IsSubsetOf(_865_v25);
        _nw123[(new BigNumber(12)).toNumber()] = _839_v1;
        _nw123[(new BigNumber(13)).toNumber()] = !(_839_v1);
        _nw123[(new BigNumber(14)).toNumber()] = _839_v1;
        _nw123[(new BigNumber(15)).toNumber()] = _839_v1;
        _nw123[(new BigNumber(16)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber(((_866_v26)[_module.__default.safeIndex(_this.f23, new BigNumber((_866_v26).length))]).length))).isEqualTo(_this.f23);
        _nw123[(new BigNumber(17)).toNumber()] = _839_v1;
        _867_v27 = _nw123;
        let _index98 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_867_v27).length));
        (_867_v27)[_index98] = _839_v1;
        let _index99 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_867_v27).length));
        (_867_v27)[_index99] = _839_v1;
        let _868_v28;
        _868_v28 = _dafny.Map.Empty.slice().updateUnsafe(_867_v27,(_dafny.ZERO).minus((_this.f23).plus(new BigNumber(732))));
        _868_v28 = (_868_v28).update(_867_v27, (new BigNumber((_837_v0).length)).minus((_this).f22));
        _839_v1 = true;
        let _869_v29;
        let _nw124 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _869_v29 = _nw124;
        let _870_v30;
        _870_v30 = _dafny.Set.fromElements((_this).f22, _this.f23, new BigNumber((_dafny.Set.fromElements((_this).f11)).length), _this.f23);
        let _871_v31;
        _871_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(491),new BigNumber((_870_v30).length));
        let _872_v32;
        _872_v32 = _dafny.Map.Empty.slice().updateUnsafe(_869_v29,_871_v31);
        _872_v32 = (_872_v32).update((((_867_v27)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_867_v27).length))]) ? (_869_v29) : (_869_v29)), _871_v31);
      }
      let _873_i2;
      _873_i2 = _dafny.ZERO;
      L7: {
        while (false) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_873_i2)) {
              break L7;
            }
            _873_i2 = (_873_i2).plus(_dafny.ONE);
            let _874_v33;
            let _nw125 = new _module.C1();
            _nw125.__ctor(_839_v1, !_dafny.areEqual(_dafny.Seq.UnicodeFromString("kdveoar"), _dafny.Seq.UnicodeFromString("csqsqpon")));
            _874_v33 = _nw125;
            (globalState).f6 = false;
            (globalState).f1 = (_this).f22;
            (_this).f23 = (_dafny.ZERO).minus((_this).f11);
          }
        }
      }
      let _875_v34;
      _875_v34 = _dafny.Seq.of(true, _839_v1);
      let _876_v35;
      let _nw126 = new _module.C0();
      _nw126.__ctor(_dafny.Seq.Concat(_875_v34, _875_v34), (new BigNumber(461)).minus((_this).f11));
      _876_v35 = _nw126;
      let _877_v36;
      _877_v36 = _dafny.Set.fromElements(new BigNumber(342));
      let _878_v37;
      let _nw127 = Array((new BigNumber(16)).toNumber());
      _nw127[(_dafny.ZERO).toNumber()] = !((new BigNumber((_837_v0).length)).isEqualTo(_this.f23));
      _nw127[(_dafny.ONE).toNumber()] = _839_v1;
      _nw127[(new BigNumber(2)).toNumber()] = (_dafny.Set.fromElements((_this).f22)).IsProperSubsetOf(_877_v36);
      _nw127[(new BigNumber(3)).toNumber()] = false;
      _nw127[(new BigNumber(4)).toNumber()] = !(_839_v1);
      _nw127[(new BigNumber(5)).toNumber()] = (new BigNumber(-62)).isLessThan(new BigNumber((_dafny.Seq.of(_839_v1)).length));
      _nw127[(new BigNumber(6)).toNumber()] = _839_v1;
      _nw127[(new BigNumber(7)).toNumber()] = ((_839_v1) ? (_839_v1) : (_839_v1));
      _nw127[(new BigNumber(8)).toNumber()] = false;
      _nw127[(new BigNumber(9)).toNumber()] = _839_v1;
      _nw127[(new BigNumber(10)).toNumber()] = _839_v1;
      _nw127[(new BigNumber(11)).toNumber()] = (_839_v1) && (_839_v1);
      _nw127[(new BigNumber(12)).toNumber()] = _839_v1;
      _nw127[(new BigNumber(13)).toNumber()] = _839_v1;
      _nw127[(new BigNumber(14)).toNumber()] = _839_v1;
      _nw127[(new BigNumber(15)).toNumber()] = _839_v1;
      _878_v37 = _nw127;
      let _index100 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length));
      (_878_v37)[_index100] = _839_v1;
      let _879_v38;
      _879_v38 = _dafny.Map.Empty.slice().updateUnsafe(_839_v1,(_876_v35).f21);
      let _880_v39;
      _880_v39 = _dafny.MultiSet.fromElements(_879_v38);
      let _881_v40;
      _881_v40 = _dafny.Map.Empty.slice().updateUnsafe((_876_v35).f21,_880_v39);
      let _882_v41;
      _882_v41 = _module.D9.create_DC18(_879_v38);
      let _index101 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length));
      (_878_v37)[_index101] = ((_880_v39).Difference(_880_v39)).IsSubsetOf((((_881_v40).contains(_this.f23)) ? ((_881_v40).get(_this.f23)) : (_dafny.MultiSet.fromElements(_879_v38, _879_v38, (_882_v41).dtor_cf26))));
      let _883_v42;
      _883_v42 = _dafny.Map.Empty.slice().updateUnsafe((_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))],(_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))]);
      let _884_i3;
      _884_i3 = _dafny.ZERO;
      L8: {
        while ((_883_v42).contains((_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))])) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_884_i3)) {
              break L8;
            }
            _884_i3 = (_884_i3).plus(_dafny.ONE);
            if (!(_839_v1)) {
              let _885_v43;
              _885_v43 = _dafny.MultiSet.fromElements((_this).f22, (_876_v35).f21, (_this).fm3((_dafny.ZERO).minus((_this).f11), _839_v1, _839_v1, globalState));
              let _886_v44;
              _886_v44 = _dafny.Map.Empty.slice().updateUnsafe((_876_v35).f21,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-99)), ((_887_p0) => function (_888_i4) {
                return _887_p0;
              })(p0)));
              (_this).f23 = (((_this).f22).minus((_876_v35).f21)).minus(_module.__default.safeModuloInt(new BigNumber((_885_v43).cardinality()), (_this).fm6((_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))], new BigNumber((_886_v44).length), globalState)));
              (globalState).f7 = _839_v1;
              _839_v1 = (_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))];
              (globalState).f1 = _module.__default.safeModuloInt((_876_v35).f21, _this.f23);
              let _889_v45;
              _889_v45 = _dafny.MultiSet.fromElements((_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))], _839_v1);
              let _890_v46;
              _890_v46 = _dafny.Map.Empty.slice().updateUnsafe(_889_v45,_878_v37);
              _890_v46 = _890_v46;
            } else {
              r0 = _dafny.MultiSet.fromElements(_839_v1, (_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))], (((_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))]) ? (_839_v1) : (_839_v1)));
              let _891_v47;
              _891_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm26(globalState)).length),_839_v1);
              let _892_v48;
              _892_v48 = new _dafny.CodePoint('d'.codePointAt(0));
              let _893_v49;
              _893_v49 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_module.__default.fm27(_839_v1, _839_v1, globalState)),false);
              let _rhs152 = (_893_v49).equals((_893_v49).Merge(_893_v49));
              let _rhs153 = ((_this).f11).minus(_this.f23);
              let _rhs154 = _dafny.Seq.Concat(_837_v0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qsx"), _dafny.Seq.UnicodeFromString("mhrhniii")));
              let _rhs155 = (_891_v47).Merge(_891_v47);
              let _rhs156 = p0;
              let _lhs125 = globalState;
              let _lhs126 = globalState;
              _lhs125.f6 = _rhs152;
              _lhs126.f1 = _rhs153;
              _837_v0 = _rhs154;
              _891_v47 = _rhs155;
              _892_v48 = _rhs156;
              let _894_v50;
              _894_v50 = _dafny.MultiSet.fromElements(false);
              let _895_v51;
              _895_v51 = _dafny.Seq.of((_dafny.ZERO).minus((((_894_v50).contains(_839_v1)) ? ((_894_v50).get(_839_v1)) : (_this.f23))));
              let _896_v52;
              _896_v52 = _dafny.MultiSet.fromElements((_895_v51)[_module.__default.safeIndex((_this).f22, new BigNumber((_895_v51).length))], _this.f23);
              _896_v52 = _896_v52;
              let _897_v53;
              _897_v53 = _dafny.Map.Empty.slice().updateUnsafe(_839_v1,p0);
              let _898_v54;
              let _nw128 = Array((new BigNumber(19)).toNumber());
              _nw128[(_dafny.ZERO).toNumber()] = (_this).f11;
              _nw128[(_dafny.ONE).toNumber()] = new BigNumber((_897_v53).length);
              _nw128[(new BigNumber(2)).toNumber()] = _this.f23;
              _nw128[(new BigNumber(3)).toNumber()] = new BigNumber(522);
              _nw128[(new BigNumber(4)).toNumber()] = new BigNumber((_895_v51).length);
              _nw128[(new BigNumber(5)).toNumber()] = new BigNumber(845);
              _nw128[(new BigNumber(6)).toNumber()] = (_this).f11;
              _nw128[(new BigNumber(7)).toNumber()] = (_876_v35).f21;
              _nw128[(new BigNumber(8)).toNumber()] = (_876_v35).f21;
              _nw128[(new BigNumber(9)).toNumber()] = (_this).f22;
              _nw128[(new BigNumber(10)).toNumber()] = (_876_v35).f21;
              _nw128[(new BigNumber(11)).toNumber()] = new BigNumber((_883_v42).length);
              _nw128[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements((_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))], _839_v1)).cardinality());
              _nw128[(new BigNumber(13)).toNumber()] = _this.f23;
              _nw128[(new BigNumber(14)).toNumber()] = new BigNumber((_837_v0).length);
              _nw128[(new BigNumber(15)).toNumber()] = (_this).f22;
              _nw128[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus((_this).f22);
              _nw128[(new BigNumber(17)).toNumber()] = new BigNumber(138);
              _nw128[(new BigNumber(18)).toNumber()] = (_876_v35).f21;
              _898_v54 = _nw128;
              let _899_v55;
              _899_v55 = _dafny.Set.fromElements(_898_v54);
              _899_v55 = (_899_v55).Intersect((_899_v55).Union(_899_v55));
              let _900_v56;
              _900_v56 = _837_v0;
              let _901_v57;
              _901_v57 = _dafny.Seq.of(_837_v0);
              (globalState).f7 = _dafny.Seq.contains(_dafny.Seq.update(_901_v57, _module.__default.safeIndex(new BigNumber(62), new BigNumber((_901_v57).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(537)), ((_902_v48) => function (_903_i5) {
                return _902_v48;
              })(_892_v48))), (_900_v56));
            }
            (globalState).f6 = (_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))];
            let _904_v58;
            _904_v58 = new _dafny.CodePoint('w'.codePointAt(0));
            _904_v58 = p0;
            let _905_v59;
            _905_v59 = _dafny.Seq.of((_876_v35).f21);
            let _906_v60;
            let _nw129 = new _module.C2();
            _nw129.__ctor(_904_v58, new BigNumber((_905_v59).length));
            _906_v60 = _nw129;
            let _907_v61;
            _907_v61 = _dafny.Map.Empty.slice().updateUnsafe(_906_v60,_878_v37);
            let _908_v62;
            _908_v62 = _module.D0.create_DC0(new BigNumber((_907_v61).length));
            (globalState).f6 = (_this).fm5(_908_v62, _dafny.Seq.Concat(_dafny.Seq.of((_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))]), (_876_v35).f20), ((_this).fm6((_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))], (_this).f11, globalState)).multipliedBy(_this.f23), new _dafny.CodePoint('e'.codePointAt(0)), globalState);
          }
        }
      }
      let _909_v63;
      _909_v63 = _dafny.MultiSet.fromElements(_839_v1);
      r0 = ((_909_v63).Intersect(_dafny.MultiSet.fromElements((_878_v37)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_878_v37).length))], _839_v1, _839_v1, _839_v1))).update((_839_v1) && (_839_v1), _module.__default.abs((_this).f11));
      return r0;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.Map.Empty;
      let _910_v0;
      _910_v0 = true;
      if (_910_v0) {
        let _911_v1;
        _911_v1 = new _dafny.CodePoint('v'.codePointAt(0));
        let _912_v2;
        let _nw130 = Array((_dafny.ONE).toNumber());
        _nw130[(_dafny.ZERO).toNumber()] = (_this).f22;
        _912_v2 = _nw130;
        let _913_v3;
        let _nw131 = new _module.C2();
        _nw131.__ctor(_911_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_912_v2,(_this).f11)).length));
        _913_v3 = _nw131;
        (globalState).f6 = _910_v0;
        let _914_v4;
        _914_v4 = _dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)));
        (globalState).f1 = (_dafny.ZERO).minus((_this).fm3(new BigNumber((_914_v4).length), !(_910_v0), _910_v0, globalState));
        let _915_v5;
        _915_v5 = _dafny.Set.fromElements((_this).f22, (new BigNumber(283)).multipliedBy(p0), ((_this).f22).minus(new BigNumber(565)), p0);
        (globalState).f1 = new BigNumber((_915_v5).length);
        let _916_v6;
        _916_v6 = _dafny.Set.fromElements(_910_v0);
        _916_v6 = _module.__default.fm41(globalState);
      } else {
        let _917_v7;
        _917_v7 = _dafny.Seq.of(_910_v0);
        let _918_v8;
        _918_v8 = _dafny.Seq.of(p0, p0);
        let _919_v9;
        _919_v9 = _dafny.Set.fromElements(_918_v8);
        let _920_v10;
        _920_v10 = _dafny.Map.Empty.slice().updateUnsafe(_910_v0,_919_v9);
        let _921_v11;
        _921_v11 = _dafny.Map.Empty.slice().updateUnsafe(_910_v0,new BigNumber(((((_920_v10).contains(_910_v0)) ? ((_920_v10).get(_910_v0)) : (_919_v9))).length));
        let _922_v12;
        _922_v12 = _dafny.Map.Empty.slice().updateUnsafe((((_921_v11).contains(_910_v0)) ? ((_921_v11).get(_910_v0)) : ((_this).f11)),_917_v7);
        let _923_v13;
        let _nw132 = new _module.C1();
        _nw132.__ctor(!_dafny.areEqual(_917_v7, (((_922_v12).contains((_this).f22)) ? ((_922_v12).get((_this).f22)) : (_917_v7))), true);
        _923_v13 = _nw132;
        (globalState).f7 = _910_v0;
        let _924_v14;
        _924_v14 = _dafny.Seq.UnicodeFromString("vea");
        let _925_v15;
        _925_v15 = _module.D7.create_DC13((_this).f12, (_this).f11, p0);
        let _926_v16;
        let _nw133 = new _module.C0();
        _nw133.__ctor(_dafny.Seq.Concat(_917_v7, _module.__default.fm35(new BigNumber((_924_v14).length), p0, _925_v15, (_this).f11, globalState)), p0);
        _926_v16 = _nw133;
        let _927_v17;
        let _init14 = ((_928_v0) => function (_929_i0) {
          return _928_v0;
        })(_910_v0);
        let _nw134 = Array((new BigNumber(7)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw134.length); _i0_14++) {
          _nw134[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _927_v17 = _nw134;
        let _rhs157 = _926_v16;
        let _rhs158 = _927_v17;
        let _rhs159 = _910_v0;
        let _rhs160 = ((!(_910_v0)) ? (_910_v0) : ((((_923_v13).f19) ? ((_923_v13).f18) : ((_923_v13).f18))));
        let _rhs161 = !(_module.__default.fm0(globalState)) || ((_923_v13).f19);
        let _lhs127 = globalState;
        let _lhs128 = globalState;
        _926_v16 = _rhs157;
        _927_v17 = _rhs158;
        _910_v0 = _rhs159;
        _lhs127.f7 = _rhs160;
        _lhs128.f7 = _rhs161;
        (globalState).f1 = _this.f23;
        let _index102 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_927_v17).length));
        (_927_v17)[_index102] = _910_v0;
        let _index103 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_927_v17).length));
        (_927_v17)[_index103] = _910_v0;
      }
      let _930_v18;
      _930_v18 = _dafny.Seq.of(_910_v0);
      let _931_v19;
      _931_v19 = _dafny.Seq.of(_910_v0, (_930_v18)[_module.__default.safeIndex(_this.f23, new BigNumber((_930_v18).length))], _910_v0, !(_910_v0), _910_v0);
      let _932_i1;
      _932_i1 = _dafny.ZERO;
      L9: {
        while (_dafny.areEqual(_931_v19, _930_v18)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_932_i1)) {
              break L9;
            }
            _932_i1 = (_932_i1).plus(_dafny.ONE);
            let _933_v20;
            _933_v20 = new _dafny.CodePoint('u'.codePointAt(0));
            _933_v20 = _933_v20;
            let _934_v21;
            _934_v21 = _module.D0.create_DC0((_this).f11);
            let _935_v22;
            _935_v22 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("bu"));
            let _936_v23;
            _936_v23 = _dafny.MultiSet.fromElements((_this).fm5(_934_v21, _931_v19, new BigNumber((_935_v22).cardinality()), _933_v20, globalState), _module.__default.fm0(globalState), _910_v0, _910_v0);
            let _937_v24;
            _937_v24 = _936_v23;
            let _source16 = _937_v24;
            let _938___mcc_h0 = _source16;
            let _939_cf11 = _938___mcc_h0;
            let _rhs162 = _910_v0;
            let _rhs163 = !(((_910_v0) ? (_910_v0) : (!(_910_v0) || (false))));
            let _lhs129 = globalState;
            let _lhs130 = globalState;
            _lhs129.f7 = _rhs162;
            _lhs130.f6 = _rhs163;
            (globalState).f6 = _910_v0;
            let _940_v25;
            let _nw135 = Array((new BigNumber(8)).toNumber());
            _nw135[(_dafny.ZERO).toNumber()] = _937_v24;
            _nw135[(_dafny.ONE).toNumber()] = _937_v24;
            _nw135[(new BigNumber(2)).toNumber()] = _937_v24;
            _nw135[(new BigNumber(3)).toNumber()] = _936_v23;
            _nw135[(new BigNumber(4)).toNumber()] = _937_v24;
            _nw135[(new BigNumber(5)).toNumber()] = _937_v24;
            _nw135[(new BigNumber(6)).toNumber()] = _937_v24;
            _nw135[(new BigNumber(7)).toNumber()] = _939_cf11;
            _940_v25 = _nw135;
            let _index104 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_940_v25).length));
            (_940_v25)[_index104] = _937_v24;
            let _941_v26;
            _941_v26 = _dafny.Seq.of((_this).f22);
            let _942_v27;
            _942_v27 = _dafny.Map.Empty.slice().updateUnsafe((_941_v26)[_module.__default.safeIndex(new BigNumber(392), new BigNumber((_941_v26).length))],false);
            let _943_v28;
            _943_v28 = _dafny.Seq.of(_931_v19, _930_v18);
            let _944_v29;
            _944_v29 = _dafny.Set.fromElements(_dafny.Seq.of(_910_v0), (_943_v28)[_module.__default.safeIndex((_this).f11, new BigNumber((_943_v28).length))]);
            let _945_v31;
            _945_v31 = _dafny.MultiSet.fromElements(p0, new BigNumber(496));
            let _946_v33;
            _946_v33 = _dafny.Map.Empty.slice().updateUnsafe(_945_v31,function () {
              let _coll45 = new _dafny.Map();
              for (const _compr_45 of _dafny.IntegerRange(new BigNumber(381), new BigNumber(912))) {
                let _947_v32 = _compr_45;
                if (((new BigNumber(381)).isLessThanOrEqualTo(_947_v32)) && ((_947_v32).isLessThan(new BigNumber(912)))) {
                  _coll45.push([(_947_v32).minus((_this).f11),(_this).f11]);
                }
              }
              return _coll45;
            }());
            let _948_v34;
            let _nw136 = Array((new BigNumber(10)).toNumber());
            _nw136[(_dafny.ZERO).toNumber()] = _910_v0;
            _nw136[(_dafny.ONE).toNumber()] = _910_v0;
            _nw136[(new BigNumber(2)).toNumber()] = _910_v0;
            _nw136[(new BigNumber(3)).toNumber()] = _910_v0;
            _nw136[(new BigNumber(4)).toNumber()] = _910_v0;
            _nw136[(new BigNumber(5)).toNumber()] = _910_v0;
            _nw136[(new BigNumber(6)).toNumber()] = _910_v0;
            _nw136[(new BigNumber(7)).toNumber()] = _910_v0;
            _nw136[(new BigNumber(8)).toNumber()] = (((_942_v27).contains(new BigNumber((function () {
              let _coll47 = new _dafny.Set();
              for (const _compr_47 of (_944_v29).Elements) {
                let _950_v30 = _compr_47;
                if ((_944_v29).contains(_950_v30)) {
                  _coll47.add(_950_v30);
                }
              }
              return _coll47;
            }()).length))) ? ((_942_v27).get(new BigNumber((function () {
              let _coll46 = new _dafny.Set();
              for (const _compr_46 of (_944_v29).Elements) {
                let _949_v30 = _compr_46;
                if ((_944_v29).contains(_949_v30)) {
                  _coll46.add(_949_v30);
                }
              }
              return _coll46;
            }()).length))) : ((_this).fm4(new BigNumber(970), _946_v33, new BigNumber((_module.__default.fm26(globalState)).length), _910_v0, globalState)));
            _nw136[(new BigNumber(9)).toNumber()] = (_945_v31).IsProperSubsetOf(_945_v31);
            _948_v34 = _nw136;
            let _index105 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_948_v34).length));
            (_948_v34)[_index105] = _910_v0;
            let _951_v35;
            let _nw137 = new _module.C0();
            _nw137.__ctor(_931_v19, (_this).f22);
            _951_v35 = _nw137;
            let _index106 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_940_v25).length));
            let _index107 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_948_v34).length));
            let _rhs164 = ((true) ? (_937_v24) : (_939_cf11));
            let _rhs165 = _910_v0;
            let _rhs166 = _951_v35;
            let _lhs131 = _940_v25;
            let _lhs132 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_940_v25).length));
            let _lhs133 = _948_v34;
            let _lhs134 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_948_v34).length));
            _lhs131[_lhs132] = _rhs164;
            _lhs133[_lhs134] = _rhs165;
            _951_v35 = _rhs166;
            let _index108 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_948_v34).length));
            (_948_v34)[_index108] = true;
            let _952_v36;
            _952_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm26(globalState));
            let _rhs167 = true;
            let _rhs168 = (_dafny.ZERO).minus(((_this).f11).minus(p0));
            let _rhs169 = !(_952_v36).equals(_952_v36);
            let _lhs135 = _this;
            let _lhs136 = globalState;
            _910_v0 = _rhs167;
            _lhs135.f23 = _rhs168;
            _lhs136.f6 = _rhs169;
            let _953_v38;
            _953_v38 = _dafny.MultiSet.fromElements(function () {
              let _coll48 = new _dafny.Map();
              for (const _compr_48 of (_dafny.Set.fromElements(p0)).Elements) {
                let _954_v37 = _compr_48;
                if ((_dafny.Set.fromElements(p0)).contains(_954_v37)) {
                  _coll48.push([(_954_v37).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_910_v0,new BigNumber((_dafny.Seq.of(_this.f23, (_dafny.ZERO).minus((_this).f11))).length))).length)),_933_v20]);
                }
              }
              return _coll48;
            }());
            let _955_v39;
            _955_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_933_v20);
            let _956_v40;
            _956_v40 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_955_v39, _955_v39), _953_v38);
            let _957_v41;
            _957_v41 = _dafny.Set.fromElements(_930_v18);
            if (!((((_910_v0) ? ((_956_v40)[_module.__default.safeIndex(new BigNumber((_957_v41).length), new BigNumber((_956_v40).length))]) : (_953_v38))).IsSubsetOf(_953_v38))) {
              let _958_v42;
              let _nw138 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
              _958_v42 = _nw138;
              let _959_v43;
              _959_v43 = _dafny.Map.Empty.slice().updateUnsafe(_958_v42,_910_v0);
              _959_v43 = _959_v43;
              let _960_v44;
              let _init15 = ((_961_v0) => function (_962_i2) {
                return _961_v0;
              })(_910_v0);
              let _nw139 = Array((new BigNumber(5)).toNumber());
              for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw139.length); _i0_15++) {
                _nw139[_i0_15] = _init15(new BigNumber(_i0_15));
              }
              _960_v44 = _nw139;
              _960_v44 = ((_910_v0) ? (_960_v44) : (_960_v44));
              (globalState).f7 = !((_dafny.ZERO).minus((_this).f22)).isEqualTo(_this.f23);
              let _963_v45;
              _963_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_910_v0);
              let _964_v46;
              _964_v46 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f22),_963_v45);
              let _965_v47;
              _965_v47 = _module.D7.create_DC13((_this).f12, new BigNumber((_936_v23).cardinality()), _this.f23);
              let _966_v48;
              _966_v48 = _dafny.Seq.UnicodeFromString("vxka");
              let _967_v49;
              _967_v49 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,(_this).f11);
              let _968_v50;
              let _nw140 = new _module.C1();
              _nw140.__ctor(_910_v0, _910_v0);
              _968_v50 = _nw140;
              let _969_v51;
              _969_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_968_v50);
              let _rhs170 = (_964_v46).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_963_v45));
              let _rhs171 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm35((_this).f22, (((_967_v49).contains(p0)) ? ((_967_v49).get(p0)) : ((_dafny.ZERO).minus(_this.f23))), _module.D7.create_DC13((_this).f12, p0, p0), _this.f23, globalState), _module.__default.safeIndex(new BigNumber((_969_v51).length), new BigNumber((_module.__default.fm35((_this).f22, (((_967_v49).contains(p0)) ? ((_967_v49).get(p0)) : ((_dafny.ZERO).minus(_this.f23))), _module.D7.create_DC13((_this).f12, p0, p0), _this.f23, globalState)).length)), (_968_v50).f18), _930_v18), _module.__default.safeIndex((_this).f22, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm35((_this).f22, (((_967_v49).contains(p0)) ? ((_967_v49).get(p0)) : ((_dafny.ZERO).minus(_this.f23))), _module.D7.create_DC13((_this).f12, p0, p0), _this.f23, globalState), _module.__default.safeIndex(new BigNumber((_969_v51).length), new BigNumber((_module.__default.fm35((_this).f22, (((_967_v49).contains(p0)) ? ((_967_v49).get(p0)) : ((_dafny.ZERO).minus(_this.f23))), _module.D7.create_DC13((_this).f12, p0, p0), _this.f23, globalState)).length)), (_968_v50).f18), _930_v18)).length)), _910_v0), _module.__default.safeIndex(_this.f23, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm35((_this).f22, (((_967_v49).contains(p0)) ? ((_967_v49).get(p0)) : ((_dafny.ZERO).minus(_this.f23))), _module.D7.create_DC13((_this).f12, p0, p0), _this.f23, globalState), _module.__default.safeIndex(new BigNumber((_969_v51).length), new BigNumber((_module.__default.fm35((_this).f22, (((_967_v49).contains(p0)) ? ((_967_v49).get(p0)) : ((_dafny.ZERO).minus(_this.f23))), _module.D7.create_DC13((_this).f12, p0, p0), _this.f23, globalState)).length)), (_968_v50).f18), _930_v18), _module.__default.safeIndex((_this).f22, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm35((_this).f22, (((_967_v49).contains(p0)) ? ((_967_v49).get(p0)) : ((_dafny.ZERO).minus(_this.f23))), _module.D7.create_DC13((_this).f12, p0, p0), _this.f23, globalState), _module.__default.safeIndex(new BigNumber((_969_v51).length), new BigNumber((_module.__default.fm35((_this).f22, (((_967_v49).contains(p0)) ? ((_967_v49).get(p0)) : ((_dafny.ZERO).minus(_this.f23))), _module.D7.create_DC13((_this).f12, p0, p0), _this.f23, globalState)).length)), (_968_v50).f18), _930_v18)).length)), _910_v0)).length)), (_968_v50).f18);
              let _rhs172 = (((_968_v50).f19) ? (_965_v47) : (_965_v47));
              let _rhs173 = _dafny.Seq.Concat(_966_v48, _dafny.Seq.UnicodeFromString("eppqimqu"));
              _964_v46 = _rhs170;
              _930_v18 = _rhs171;
              _965_v47 = _rhs172;
              _966_v48 = _rhs173;
              let _970_v52;
              _970_v52 = _dafny.Seq.of(new BigNumber(-590), _this.f23);
              let _971_v53;
              _971_v53 = _dafny.MultiSet.fromElements((_this).f22);
              let _972_v54;
              _972_v54 = _dafny.Seq.of(_970_v52, _970_v52, _970_v52, _dafny.Seq.Concat(_dafny.Seq.of(p0, new BigNumber((_971_v53).cardinality()), new BigNumber(894)), _970_v52), _dafny.Seq.of((_this).f11));
              _970_v52 = (_972_v54)[_module.__default.safeIndex((new BigNumber((_970_v52).length)).multipliedBy((_this).f11), new BigNumber((_972_v54).length))];
            } else {
              let _973_v55;
              let _nw141 = Array((new BigNumber(22)).toNumber()).fill(false);
              _973_v55 = _nw141;
              let _index109 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_973_v55).length));
              (_973_v55)[_index109] = _910_v0;
              let _974_v56;
              _974_v56 = _dafny.Seq.UnicodeFromString("iq");
              let _975_v57;
              _975_v57 = _module.D2.create_DC3(_935_v22);
              let _index110 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_973_v55).length));
              (_973_v55)[_index110] = ((_975_v57).dtor_cf5).IsProperSubsetOf((_dafny.MultiSet.fromElements(_974_v56)).Union(_935_v22));
              _936_v23 = (_936_v23).Union((_dafny.MultiSet.fromElements((_973_v55)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_973_v55).length))], (_973_v55)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_973_v55).length))], (_973_v55)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_973_v55).length))], (_973_v55)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_973_v55).length))], _910_v0)).Intersect(_936_v23));
              let _976_v58;
              let _nw142 = new _module.C2();
              _nw142.__ctor(_module.__default.fm39(new BigNumber(175), (_973_v55)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_973_v55).length))], (_this).f11, globalState), p0);
              _976_v58 = _nw142;
              let _977_v59;
              _977_v59 = _dafny.Map.Empty.slice().updateUnsafe(_910_v0,(_dafny.ZERO).minus((_this.f23).multipliedBy((_this).f22)));
              let _978_v60;
              _978_v60 = _dafny.MultiSet.fromElements(new BigNumber(204));
              let _979_v61;
              _979_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f11);
              let _980_v62;
              _980_v62 = _dafny.Map.Empty.slice().updateUnsafe(_978_v60,_979_v61);
              _977_v59 = (_977_v59).update(!((_this).fm4(_this.f23, _980_v62, new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_974_v56, _module.__default.safeIndex((_this).f11, new BigNumber((_974_v56).length)), _933_v20), _module.__default.safeIndex(new BigNumber(117), new BigNumber((_dafny.Seq.update(_974_v56, _module.__default.safeIndex((_this).f11, new BigNumber((_974_v56).length)), _933_v20)).length)), _933_v20)).length), true, globalState)), (_dafny.ZERO).minus((_this).f22));
              let _index111 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_973_v55).length));
              (_973_v55)[_index111] = !(!(_910_v0));
            }
          }
        }
      }
      let _981_v63;
      _981_v63 = _dafny.Set.fromElements((_this).f11);
      let _rhs174 = _910_v0;
      let _rhs175 = _module.__default.safeDivisionInt((_this).f22, (_this).f11);
      let _rhs176 = p0;
      let _rhs177 = !(_981_v63).contains((_this).f22);
      let _lhs137 = _this;
      let _lhs138 = _this;
      let _lhs139 = globalState;
      _910_v0 = _rhs174;
      _lhs137.f23 = _rhs175;
      _lhs138.f23 = _rhs176;
      _lhs139.f6 = _rhs177;
      let _982_v64;
      _982_v64 = _dafny.Seq.UnicodeFromString("qnfp");
      let _983_v65;
      _983_v65 = _dafny.Seq.of(new BigNumber((_982_v64).length));
      let _984_v66;
      _984_v66 = _dafny.MultiSet.fromElements((_983_v65)[_module.__default.safeIndex((_this).f11, new BigNumber((_983_v65).length))], (_this).f11, (_this).f11);
      let _985_v67;
      _985_v67 = _dafny.Map.Empty.slice().updateUnsafe(_982_v64,_984_v66);
      r1 = _985_v67;
      let _986_v68;
      _986_v68 = _dafny.Seq.of(_930_v18, _930_v18, _930_v18);
      let _987_i3;
      _987_i3 = _dafny.ZERO;
      L10: {
        while (!(true) || (_dafny.Seq.contains(_986_v68, _931_v19))) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_987_i3)) {
              break L10;
            }
            _987_i3 = (_987_i3).plus(_dafny.ONE);
            if (_910_v0) {
              let _988_v69;
              _988_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(785),_982_v64);
              let _rhs178 = (_910_v0) === ((_dafny.MultiSet.fromElements((((_988_v69).contains((_this).f22)) ? ((_988_v69).get((_this).f22)) : (_982_v64)))).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(761)), ((_989_v64) => function (_990_i4) {
                return _989_v64;
              })(_982_v64)))));
              let _rhs179 = p0;
              let _rhs180 = new BigNumber(598);
              let _lhs140 = globalState;
              let _lhs141 = globalState;
              let _lhs142 = globalState;
              _lhs140.f6 = _rhs178;
              _lhs141.f1 = _rhs179;
              _lhs142.f1 = _rhs180;
              (globalState).f1 = (_this).f22;
              let _991_v70;
              _991_v70 = new _dafny.CodePoint('g'.codePointAt(0));
              let _992_v72;
              let _nw143 = new _module.C0();
              _nw143.__ctor(_931_v19, (_this).f22);
              _992_v72 = _nw143;
              let _993_v73;
              _993_v73 = _dafny.Seq.of(_992_v72);
              let _994_v74;
              _994_v74 = _module.D0.create_DC0(new BigNumber((_993_v73).length));
              let _995_v75;
              _995_v75 = _dafny.MultiSet.fromElements(_994_v74, _994_v74, _994_v74, _994_v74, _994_v74);
              _991_v70 = _module.__default.fm39(new BigNumber(599), _module.__default.fm0(globalState), new BigNumber((function () {
                let _coll49 = new _dafny.Map();
                for (const _compr_49 of (_995_v75).Elements) {
                  let _996_v71 = _compr_49;
                  if ((_995_v75).contains(_996_v71)) {
                    _coll49.push([_996_v71,(_this).f11]);
                  }
                }
                return _coll49;
              }()).length), globalState);
              (globalState).f7 = _dafny.areEqual(_930_v18, _dafny.Seq.Concat((_992_v72).f20, _930_v18));
              let _997_v76;
              let _init16 = ((_998_v0, _999_v66) => function (_1000_i5) {
                return ((_998_v0) ? (_999_v66) : (_999_v66));
              })(_910_v0, _984_v66);
              let _nw144 = Array((_dafny.ONE).toNumber());
              for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw144.length); _i0_16++) {
                _nw144[_i0_16] = _init16(new BigNumber(_i0_16));
              }
              _997_v76 = _nw144;
              let _index112 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_997_v76).length));
              (_997_v76)[_index112] = _984_v66;
              let _index113 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_997_v76).length));
              let _rhs181 = _module.__default.safeDivisionInt(p0, (_this).f22);
              let _rhs182 = _dafny.Seq.Concat(_982_v64, _dafny.Seq.update(_dafny.Seq.Concat(_982_v64, _982_v64), _module.__default.safeIndex((_983_v65)[_module.__default.safeIndex(new BigNumber((_983_v65).length), new BigNumber((_983_v65).length))], new BigNumber((_dafny.Seq.Concat(_982_v64, _982_v64)).length)), _991_v70));
              let _rhs183 = _this.f23;
              let _rhs184 = _984_v66;
              let _lhs143 = globalState;
              let _lhs144 = _this;
              let _lhs145 = _997_v76;
              let _lhs146 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_997_v76).length));
              _lhs143.f1 = _rhs181;
              _982_v64 = _rhs182;
              _lhs144.f23 = _rhs183;
              _lhs145[_lhs146] = _rhs184;
            } else {
              (globalState).f1 = (_this).f22;
              let _1001_v77;
              _1001_v77 = _dafny.Seq.of(_983_v65);
              let _1002_v78;
              _1002_v78 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_1001_v77, _dafny.Seq.of(_this.f23, new BigNumber(-896))),_this.f23);
              _1002_v78 = _1002_v78;
              let _1003_v79;
              _1003_v79 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-274)), function (_1004_i6) {
                return new _dafny.CodePoint('q'.codePointAt(0));
              }));
              let _1005_v80;
              _1005_v80 = new _dafny.CodePoint('n'.codePointAt(0));
              let _rhs185 = ((new BigNumber((_981_v63).length)).minus((_this).fm3(p0, false, _910_v0, globalState))).multipliedBy(((_910_v0) ? ((_this).f11) : ((_this).f11)));
              let _rhs186 = !(!((_931_v19)[_module.__default.safeIndex(_this.f23, new BigNumber((_931_v19).length))]));
              let _rhs187 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm1(globalState), _dafny.Seq.of(_1005_v80)), _982_v64);
              let _rhs188 = _1003_v79;
              let _rhs189 = true;
              let _lhs147 = globalState;
              let _lhs148 = globalState;
              _lhs147.f1 = _rhs185;
              _lhs148.f6 = _rhs186;
              _982_v64 = _rhs187;
              _1003_v79 = _rhs188;
              _910_v0 = _rhs189;
              (globalState).f1 = (_this).f11;
              let _1006_v81;
              let _nw145 = new _module.C1();
              _nw145.__ctor(_910_v0, (_984_v66).IsDisjointFrom(_984_v66));
              _1006_v81 = _nw145;
            }
            let _1007_v82;
            _1007_v82 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_930_v18);
            let _1008_v83;
            let _nw146 = new _module.C0();
            _nw146.__ctor((((_1007_v82).contains((_983_v65)[_module.__default.safeIndex(new BigNumber((_983_v65).length), new BigNumber((_983_v65).length))])) ? ((_1007_v82).get((_983_v65)[_module.__default.safeIndex(new BigNumber((_983_v65).length), new BigNumber((_983_v65).length))])) : (_930_v18)), _this.f23);
            _1008_v83 = _nw146;
            let _1009_v84;
            _1009_v84 = _dafny.Seq.of((_this).f12, (_this).f12, (_this).f12, (_this).f12);
            _1009_v84 = _1009_v84;
            _983_v65 = _dafny.Seq.of((_this).f22, p0, (_1008_v83).f21, _this.f23, (_this).f22);
          }
        }
      }
      let _1010_v85;
      _1010_v85 = _dafny.Set.fromElements(_983_v65);
      let _1011_v86;
      _1011_v86 = _module.D9.create_DC19();
      let _1012_v87;
      let _nw147 = Array((new BigNumber(18)).toNumber());
      _nw147[(_dafny.ZERO).toNumber()] = _module.__default.fm42(_1010_v85, globalState);
      _nw147[(_dafny.ONE).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(2)).toNumber()] = ((_910_v0) ? (_1011_v86) : (_1011_v86));
      _nw147[(new BigNumber(3)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(4)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(5)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(6)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(7)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(8)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(9)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(10)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(11)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(12)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(13)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(14)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(15)).toNumber()] = _1011_v86;
      _nw147[(new BigNumber(16)).toNumber()] = _module.__default.fm42(_1010_v85, globalState);
      _nw147[(new BigNumber(17)).toNumber()] = _1011_v86;
      _1012_v87 = _nw147;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1012_v87).length))) {
        let _1013_i7 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1013_i7)) && ((_1013_i7).isLessThan(new BigNumber((_1012_v87).length))))) {
          (_1012_v87)[(_1013_i7)] = (((_this.f23).isLessThan(_this.f23)) ? (((_910_v0) ? (_1011_v86) : (_1011_v86))) : (_1011_v86));
        }
      }
      let _1014_v88;
      let _init17 = ((_1015_v0) => function (_1016_i8) {
        return _dafny.MultiSet.fromElements(_1015_v0);
      })(_910_v0);
      let _nw148 = Array((new BigNumber(8)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw148.length); _i0_17++) {
        _nw148[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _1014_v88 = _nw148;
      let _1017_v89;
      _1017_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_1014_v88);
      r0 = (((_1017_v89).contains((_dafny.ZERO).minus(new BigNumber(-6)))) ? ((_1017_v89).get((_dafny.ZERO).minus(new BigNumber(-6)))) : (_1014_v88));
      r1 = (_985_v67).Merge(_module.__default.fm43(_910_v0, globalState));
      return [r0, r1];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f12 = _dafny.Set.Empty;
      this._f11 = _dafny.ZERO;
      this._f16 = _dafny.MultiSet.Empty;
      this._f17 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f16, f17, f12, f11) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f12 = f12;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f17;
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((((_this).f16).Intersect((_this).f16)).cardinality());
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("njlnvedvb"), _dafny.Seq.UnicodeFromString("jjlvfx")), _dafny.Seq.UnicodeFromString("uocufamr"))).length);
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f17;
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      (globalState).f7 = (_this).f17;
      let _1018_i0;
      _1018_i0 = _dafny.ZERO;
      L11: {
        while ((_this).f17) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1018_i0)) {
              break L11;
            }
            _1018_i0 = (_1018_i0).plus(_dafny.ONE);
            let _1019_v0;
            _1019_v0 = _dafny.MultiSet.fromElements((_this).f17, (_this).f17, (_this).f17);
            r0 = ((((_this).f17) ? (_1019_v0) : (_dafny.MultiSet.fromElements((_this).f17)))).Difference(_1019_v0);
            let _1020_v1;
            let _nw149 = Array((new BigNumber(8)).toNumber()).fill(false);
            _1020_v1 = _nw149;
            let _index114 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1020_v1).length));
            (_1020_v1)[_index114] = (_this).f17;
            let _1021_v2;
            let _nw150 = Array((new BigNumber(4)).toNumber()).fill(_module.D2.Default());
            _1021_v2 = _nw150;
            let _1022_v3;
            _1022_v3 = _dafny.MultiSet.fromElements(_1021_v2);
            let _index115 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1020_v1).length));
            (_1020_v1)[_index115] = ((_dafny.MultiSet.fromElements(_1021_v2, _1021_v2)).Difference(_1022_v3)).IsDisjointFrom(_1022_v3);
            (globalState).f7 = false;
            let _1023_v4;
            let _nw151 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Set.Empty);
            _1023_v4 = _nw151;
            let _1024_v5;
            _1024_v5 = _dafny.Seq.of((_this).f17, (_this).f17);
            let _1025_v6;
            _1025_v6 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f17),(_this).f11);
            let _1026_v7;
            _1026_v7 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_1020_v1)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_1020_v1).length))],new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1024_v5,(_this).fm3((_this).f11, (_1020_v1)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_1020_v1).length))], (_1020_v1)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_1020_v1).length))], globalState))).length)), _1025_v6);
            let _index116 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1023_v4).length));
            (_1023_v4)[_index116] = (_1026_v7).Difference(_1026_v7);
            let _1027_v8;
            let _nw152 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
            _1027_v8 = _nw152;
            let _1028_v9;
            _1028_v9 = _dafny.MultiSet.fromElements(_1027_v8);
            let _1029_v10;
            _1029_v10 = _module.D2.create_DC4((_dafny.ZERO).minus((_this).f11), (((_1028_v9).contains(_1027_v8)) ? ((_1028_v9).get(_1027_v8)) : ((_this).f11)), (_this).f11);
            let _1030_v11;
            _1030_v11 = _dafny.Seq.of((_this).f11, (_this).f11);
            let _1031_v12;
            _1031_v12 = _dafny.Seq.UnicodeFromString("j");
            let _1032_v13;
            _1032_v13 = _dafny.Seq.of(new BigNumber((_1030_v11).length), new BigNumber((_1024_v5).length), new BigNumber((_1031_v12).length));
            let _pat_let_tv36 = _1030_v11;
            let _1033_v14;
            _1033_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,function (_pat_let24_0) {
              return function (_1034_dt__update__tmp_h0) {
                return function (_pat_let25_0) {
                  return function (_1035_dt__update_hcf6_h0) {
                    return function (_pat_let26_0) {
                      return function (_1036_dt__update_hcf8_h0) {
                        return _module.D2.create_DC4(_1035_dt__update_hcf6_h0, (_1034_dt__update__tmp_h0).dtor_cf7, _1036_dt__update_hcf8_h0);
                      }(_pat_let26_0);
                    }(new BigNumber((_pat_let_tv36).length));
                  }(_pat_let25_0);
                }((_this).f11);
              }(_pat_let24_0);
            }(_1029_v10));
            let _index117 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_1023_v4).length));
            (_1023_v4)[_index117] = _module.__default.fm10((_this).f11, _dafny.Seq.update(((!((_1020_v1)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_1020_v1).length))])) ? (_dafny.Seq.UnicodeFromString("smwrdj")) : (_dafny.Seq.UnicodeFromString("tctiechk"))), _module.__default.safeIndex((_this).f11, new BigNumber((((!((_1020_v1)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_1020_v1).length))])) ? (_dafny.Seq.UnicodeFromString("smwrdj")) : (_dafny.Seq.UnicodeFromString("tctiechk")))).length)), new _dafny.CodePoint('q'.codePointAt(0))), new BigNumber((_1033_v14).length), (_1020_v1)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_1020_v1).length))], globalState);
          }
        }
      }
      let _1037_v15;
      let _nw153 = Array((new BigNumber(5)).toNumber()).fill(false);
      _1037_v15 = _nw153;
      let _1038_v16;
      _1038_v16 = _dafny.Seq.UnicodeFromString("oc");
      let _rhs190 = (_this).fm6((_this).f17, _module.__default.safeModuloInt((_this).f11, (_dafny.ZERO).minus((_this).f11)), globalState);
      let _rhs191 = (_this).f11;
      let _rhs192 = _1037_v15;
      let _rhs193 = _1037_v15;
      let _rhs194 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bjar"), _1038_v16)).length)).plus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f17)).Merge(_module.__default.fm11(globalState))).length));
      let _lhs149 = globalState;
      let _lhs150 = globalState;
      let _lhs151 = globalState;
      _lhs149.f1 = _rhs190;
      _lhs150.f1 = _rhs191;
      _1037_v15 = _rhs192;
      _1037_v15 = _rhs193;
      _lhs151.f1 = _rhs194;
      let _index118 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length));
      (_1037_v15)[_index118] = (_this).f17;
      let _index119 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length));
      (_1037_v15)[_index119] = (_this).f17;
      (globalState).f6 = (_this).f17;
      let _1039_v17;
      _1039_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]),(_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]);
      if ((_1039_v17).contains(_module.__default.fm12((_dafny.ZERO).minus((_this).f11), _1038_v16, globalState))) {
        let _1040_v18;
        _1040_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f11);
        let _rhs195 = new BigNumber((_1040_v18).length);
        let _rhs196 = (_this).f17;
        let _lhs152 = globalState;
        let _lhs153 = globalState;
        _lhs152.f1 = _rhs195;
        _lhs153.f7 = _rhs196;
        let _1041_v19;
        _1041_v19 = _dafny.Map.Empty.slice().updateUnsafe((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))],(_this).f17);
        let _1042_v20;
        _1042_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1038_v16).length),(_this).f17);
        let _1043_v21;
        _1043_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_1042_v20).contains((_this).f11)) ? ((_1042_v20).get((_this).f11)) : (false)));
        let _1044_v22;
        _1044_v22 = _dafny.Seq.of((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f11, (_this).f11)), (((((_1041_v19).contains((_this).f17)) ? ((_1041_v19).get((_this).f17)) : ((_this).f17))) ? ((_this).f11) : ((_this).f11)), (_this).f11, (_dafny.ZERO).minus(new BigNumber(((((_1043_v21).update(p0, (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))])).update(new _dafny.CodePoint('w'.codePointAt(0)), (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))])).Merge(_1043_v21)).length)), _module.__default.safeDivisionInt((_this).f11, new BigNumber(923)));
        let _1045_v23;
        let _nw154 = Array((new BigNumber(5)).toNumber()).fill([]);
        _1045_v23 = _nw154;
        let _1046_v24;
        let _nw155 = Array((new BigNumber(14)).toNumber()).fill([]);
        _1046_v24 = _nw155;
        let _index120 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_1045_v23).length));
        (_1045_v23)[_index120] = _1046_v24;
        let _1047_v26;
        _1047_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]);
        let _1048_v28;
        _1048_v28 = _1046_v24;
        let _index121 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_1045_v23).length));
        let _rhs197 = _dafny.Seq.Concat(_1044_v22, _1044_v22);
        let _rhs198 = (_1044_v22)[_module.__default.safeIndex((_this).f11, new BigNumber((_1044_v22).length))];
        let _rhs199 = (((((_this).f17) ? ((_this).fm4((_this).f11, function () {
          let _coll50 = new _dafny.Map();
          for (const _compr_50 of (_1047_v26).Keys.Elements) {
            let _1049_v25 = _compr_50;
            if ((_1047_v26).contains(_1049_v25)) {
              _coll50.push([_1049_v25,function () {
                let _coll51 = new _dafny.Map();
                for (const _compr_51 of _dafny.IntegerRange(new BigNumber(257), new BigNumber(278))) {
                  let _1050_v27 = _compr_51;
                  if (((new BigNumber(257)).isLessThanOrEqualTo(_1050_v27)) && ((_1050_v27).isLessThan(new BigNumber(278)))) {
                    _coll51.push([(_1050_v27).multipliedBy(new BigNumber(((_this).f12).length)),new BigNumber((_dafny.Seq.of((_this).f17, (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))])).length)]);
                  }
                }
                return _coll51;
              }()]);
            }
          }
          return _coll50;
        }(), (_this).f11, (_this).f17, globalState)) : ((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]))) ? (_1046_v24) : ((_1048_v28)));
        let _rhs200 = (((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]) ? (_1037_v15) : (_1037_v15));
        let _lhs154 = globalState;
        let _lhs155 = _1045_v23;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_1045_v23).length));
        _1044_v22 = _rhs197;
        _lhs154.f1 = _rhs198;
        _lhs155[_lhs156] = _rhs199;
        _1037_v15 = _rhs200;
        if (!((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))])) {
          let _1051_v29;
          let _1052_v30;
          let _1053_v31;
          let _out28;
          let _out29;
          let _out30;
          let _outcollector9 = (_this).m5((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))], !((_this).f17) || (true), (_this).f11, globalState);
          _out28 = _outcollector9[0];
          _out29 = _outcollector9[1];
          _out30 = _outcollector9[2];
          _1051_v29 = _out28;
          _1052_v30 = _out29;
          _1053_v31 = _out30;
          _1038_v16 = _module.__default.fm1(globalState);
          let _index122 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length));
          (_1037_v15)[_index122] = (_this).f17;
          let _1054_v32;
          _1054_v32 = _dafny.Seq.of((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))], (_this).f17, (_this).f17);
          let _1055_v33;
          let _nw156 = new _module.C0();
          _nw156.__ctor(_1054_v32, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(648)), ((_1056_v20) => function (_1057_i1) {
            return (_dafny.ZERO).minus(new BigNumber((_1056_v20).length));
          })(_1042_v20))).length));
          _1055_v33 = _nw156;
          let _1058_v34;
          let _nw157 = new _module.C3();
          _nw157.__ctor(new BigNumber((_1038_v16).length), new BigNumber(729), _1053_v31, (_this).f12);
          _1058_v34 = _nw157;
          let _1059_v35;
          _1059_v35 = _dafny.Seq.of(_1058_v34, _1058_v34, _1058_v34, _1058_v34);
          let _index123 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length));
          let _rhs201 = _module.__default.safeDivisionInt((_this).f11, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_1059_v35).length))).cardinality()));
          let _rhs202 = ((((_this).f16).contains((_this).f11)) ? (((_this).f16).get((_this).f11)) : ((_1053_v31).multipliedBy(_1051_v29)));
          let _rhs203 = !(!((((_1042_v20).contains((_1058_v34).f11)) ? ((_1042_v20).get((_1058_v34).f11)) : (_1052_v30))));
          let _rhs204 = _1052_v30;
          let _lhs157 = globalState;
          let _lhs158 = globalState;
          let _lhs159 = _1037_v15;
          let _lhs160 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length));
          _1053_v31 = _rhs201;
          _lhs157.f1 = _rhs202;
          _lhs158.f6 = _rhs203;
          _lhs159[_lhs160] = _rhs204;
        } else {
          let _1060_v36;
          let _init18 = ((_1061_v16) => function (_1062_i2) {
            return (_1061_v16);
          })(_1038_v16);
          let _nw158 = Array((new BigNumber(25)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw158.length); _i0_18++) {
            _nw158[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _1060_v36 = _nw158;
          _1060_v36 = _1060_v36;
          let _1063_v37;
          _1063_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1044_v22,(_this).f11);
          let _1064_v38;
          _1064_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1063_v37,(_this).f11);
          (globalState).f1 = ((_dafny.ZERO).minus((_this).f11)).multipliedBy((((_1064_v38).contains(_1063_v37)) ? ((_1064_v38).get(_1063_v37)) : ((_this).f11)));
          (globalState).f6 = (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))];
          let _1065_v39;
          let _nw159 = new _module.C2();
          _nw159.__ctor(new _dafny.CodePoint('d'.codePointAt(0)), _module.__default.safeDivisionInt((_this).f11, (_this).f11));
          _1065_v39 = _nw159;
          (globalState).f6 = ((_this).f16).contains((_this).f11);
        }
        let _1066_v41;
        _1066_v41 = _dafny.Set.fromElements(new BigNumber((_1044_v22).length), new BigNumber((function () {
          let _coll52 = new _dafny.Map();
          for (const _compr_52 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-494)), function (_1067_i3) {
            return (_this).f11;
          })).Elements) {
            let _1068_v40 = _compr_52;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-494)), function (_1067_i3) {
              return (_this).f11;
            }), _1068_v40)) {
              _coll52.push([(_1068_v40).minus(new BigNumber((_dafny.MultiSet.fromElements((_this).f11)).cardinality())),true]);
            }
          }
          return _coll52;
        }()).length));
        (globalState).f1 = new BigNumber(((_1066_v41).Union(_1066_v41)).length);
        let _1069_v42;
        let _nw160 = new _module.C3();
        _nw160.__ctor(new BigNumber((_1038_v16).length), (_this).f11, (_this).f11, (_this).f12);
        _1069_v42 = _nw160;
        let _1070_v43;
        _1070_v43 = _dafny.Seq.of(_1069_v42, _1069_v42);
        (globalState).f1 = new BigNumber((_1070_v43).length);
      } else {
        let _1071_v44;
        let _nw161 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1071_v44 = _nw161;
        let _1072_v45;
        _1072_v45 = _dafny.Seq.of((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]);
        let _1073_v46;
        _1073_v46 = _module.D11.create_DC24((_this).f11, _1072_v45, p0);
        let _index124 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length));
        let _rhs205 = (_1073_v46).dtor_cf34;
        let _rhs206 = (_this).f11;
        let _rhs207 = (_this).f17;
        let _rhs208 = (_this).f11;
        let _rhs209 = _1071_v44;
        let _lhs161 = globalState;
        let _lhs162 = globalState;
        let _lhs163 = _1037_v15;
        let _lhs164 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length));
        let _lhs165 = globalState;
        _lhs161.f1 = _rhs205;
        _lhs162.f1 = _rhs206;
        _lhs163[_lhs164] = _rhs207;
        _lhs165.f1 = _rhs208;
        _1071_v44 = _rhs209;
        let _1074_v47;
        _1074_v47 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f11).plus((_this).f11),((_this).f11).plus((_this).f11));
        _1074_v47 = (_1074_v47).update((_this).f11, (_this).f11);
        let _1075_v48;
        let _nw162 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1075_v48 = _nw162;
        let _1076_v49;
        _1076_v49 = _module.D6.create_DC11((_this).f11, (_this).f11, (_dafny.ZERO).minus((_this).f11));
        let _1077_v50;
        _1077_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1076_v49,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("go"), _module.__default.safeIndex((_this).f11, new BigNumber((_dafny.Seq.UnicodeFromString("go")).length)), p0));
        let _1078_v51;
        _1078_v51 = _dafny.MultiSet.fromElements((((_1077_v50).contains(_1076_v49)) ? ((_1077_v50).get(_1076_v49)) : (_1038_v16)));
        let _index125 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_1075_v48).length));
        (_1075_v48)[_index125] = new BigNumber((_1078_v51).cardinality());
        let _index126 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_1075_v48).length));
        (_1075_v48)[_index126] = (new BigNumber(-940)).multipliedBy((_dafny.ZERO).minus((_this).f11));
        let _1079_v52;
        _1079_v52 = _dafny.Set.fromElements((_this).f11, (_this).f11);
        if ((new BigNumber((_1038_v16).length)).isLessThan(((_this).f11).minus(new BigNumber((_1079_v52).length)))) {
          (globalState).f1 = new BigNumber(-492);
          let _1080_v53;
          _1080_v53 = _dafny.MultiSet.fromElements((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]);
          let _1081_v54;
          let _nw163 = new _module.C1();
          _nw163.__ctor((_dafny.MultiSet.FromArray(_1072_v45)).IsSubsetOf(_1080_v53), (_this).f17);
          _1081_v54 = _nw163;
          let _1082_v55;
          _1082_v55 = _dafny.Seq.of((_1075_v48)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_1075_v48).length))], new BigNumber(-48), (_1075_v48)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_1075_v48).length))]);
          (globalState).f7 = (_module.__default.fm0(globalState)) && (((_this).f11).isLessThan((_1082_v55)[_module.__default.safeIndex((_this).f11, new BigNumber((_1082_v55).length))]));
          let _index127 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length));
          (_1037_v15)[_index127] = (_this).f17;
          (globalState).f7 = (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))];
        } else {
          let _1083_v56;
          _1083_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17);
          let _1084_v57;
          _1084_v57 = _module.D8.create_DC15(_1083_v56);
          let _1085_v58;
          _1085_v58 = _module.D8.create_DC17(_1084_v57);
          let _1086_v59;
          _1086_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1085_v58,(_1075_v48)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_1075_v48).length))]);
          let _1087_v60;
          _1087_v60 = _dafny.Seq.of(_1086_v59, _1086_v59, _1086_v59, _1086_v59);
          let _1088_v61;
          _1088_v61 = _dafny.MultiSet.fromElements(_module.__default.fm0(globalState));
          let _1089_v62;
          _1089_v62 = _module.D13.create_DC27(_1086_v59);
          let _1090_v63;
          _1090_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1087_v60);
          let _1091_v64;
          _1091_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(11),(_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]);
          let _1092_v65;
          let _nw164 = Array((new BigNumber(15)).toNumber());
          _nw164[(_dafny.ZERO).toNumber()] = _1087_v60;
          _nw164[(_dafny.ONE).toNumber()] = _1087_v60;
          _nw164[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_1087_v60, _module.__default.safeIndex((((_1088_v61).contains(true)) ? ((_1088_v61).get(true)) : ((_this).f11)), new BigNumber((_1087_v60).length)), _1086_v59), _1087_v60);
          _nw164[(new BigNumber(3)).toNumber()] = _1087_v60;
          _nw164[(new BigNumber(4)).toNumber()] = _1087_v60;
          _nw164[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_1086_v59, _1086_v59, _dafny.Map.Empty.slice().updateUnsafe(_1085_v58,(_this).f11));
          _nw164[(new BigNumber(6)).toNumber()] = _1087_v60;
          _nw164[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of((_1089_v62).dtor_cf39, _1086_v59, (_1089_v62).dtor_cf39), _dafny.Seq.of(_1086_v59)), _module.__default.safeIndex((_this).f11, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_1089_v62).dtor_cf39, _1086_v59, (_1089_v62).dtor_cf39), _dafny.Seq.of(_1086_v59))).length)), _1086_v59);
          _nw164[(new BigNumber(8)).toNumber()] = (((_1090_v63).contains((((_1091_v64).contains((_this).f11)) ? ((_1091_v64).get((_this).f11)) : ((_this).f17)))) ? ((_1090_v63).get((((_1091_v64).contains((_this).f11)) ? ((_1091_v64).get((_this).f11)) : ((_this).f17)))) : (_module.__default.fm44(!((_this).f17), (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))], globalState)));
          _nw164[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm44((_this).f17, (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))], globalState), _module.__default.safeIndex(new BigNumber(960), new BigNumber((_module.__default.fm44((_this).f17, (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))], globalState)).length)), _1086_v59), _1087_v60);
          _nw164[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1087_v60, _1087_v60);
          _nw164[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_1086_v59, _1086_v59);
          _nw164[(new BigNumber(12)).toNumber()] = _module.__default.fm44(!((_this).f17), (_this).f17, globalState);
          _nw164[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1087_v60, _1087_v60);
          _nw164[(new BigNumber(14)).toNumber()] = _dafny.Seq.of((_1086_v59).update(_module.D8.create_DC17(_1084_v57), (_dafny.ZERO).minus(new BigNumber(-8))));
          _1092_v65 = _nw164;
          let _index128 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1092_v65).length));
          (_1092_v65)[_index128] = (((_1090_v63).contains(!((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]))) ? ((_1090_v63).get(!((_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-255)), ((_1093_v59) => function (_1094_i4) {
            return _1093_v59;
          })(_1086_v59))));
          let _index129 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1092_v65).length));
          (_1092_v65)[_index129] = _1087_v60;
          let _1095_v66;
          let _init19 = ((_1096_v58) => function (_1097_i5) {
            return _1096_v58;
          })(_1085_v58);
          let _nw165 = Array((new BigNumber(23)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw165.length); _i0_19++) {
            _nw165[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _1095_v66 = _nw165;
          let _index130 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_1095_v66).length));
          (_1095_v66)[_index130] = _1085_v58;
          let _index131 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_1095_v66).length));
          (_1095_v66)[_index131] = _1085_v58;
          _1037_v15 = _1037_v15;
          (globalState).f6 = !(((_1075_v48)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_1075_v48).length))]).isLessThanOrEqualTo((_this).f11));
          let _1098_v67;
          let _nw166 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1098_v67 = _nw166;
          let _1099_v68;
          _1099_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1098_v67);
          (globalState).f1 = new BigNumber(((_1099_v68).update(((_this).f17) || ((_this).f17), _1098_v67)).length);
        }
        (globalState).f7 = (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))];
      }
      let _1100_v70;
      _1100_v70 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f16).update((_this).f11, _module.__default.abs((_this).f11)),(_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))]);
      let _1101_v72;
      _1101_v72 = _dafny.MultiSet.fromElements((_this).fm4((_this).f11, function () {
        let _coll53 = new _dafny.Map();
        for (const _compr_53 of (_1100_v70).Keys.Elements) {
          let _1102_v69 = _compr_53;
          if ((_1100_v70).contains(_1102_v69)) {
            _coll53.push([_1102_v69,function () {
              let _coll54 = new _dafny.Map();
              for (const _compr_54 of _dafny.IntegerRange(new BigNumber(958), new BigNumber(898))) {
                let _1103_v71 = _compr_54;
                if (((new BigNumber(958)).isLessThanOrEqualTo(_1103_v71)) && ((_1103_v71).isLessThan(new BigNumber(898)))) {
                  _coll54.push([_module.__default.safeModuloInt(_1103_v71, (_this).f11),new BigNumber(-113)]);
                }
              }
              return _coll54;
            }()]);
          }
        }
        return _coll53;
      }(), new BigNumber(893), (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))], globalState));
      let _1104_v73;
      _1104_v73 = _module.D0.create_DC0((_this).f11);
      let _1105_v74;
      _1105_v74 = _dafny.Seq.of(true, (_1037_v15)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1037_v15).length))], (_this).f17, (_this).f17);
      r0 = (_1101_v72).update((_this).fm5(_1104_v73, _1105_v74, (_this).f11, new _dafny.CodePoint('x'.codePointAt(0)), globalState), _module.__default.abs(((_this).f11).multipliedBy(new BigNumber(((_this).f12).length))));
      return r0;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.Map.Empty;
      let _1106_i0;
      _1106_i0 = _dafny.ZERO;
      L12: {
        while ((_this).f17) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1106_i0)) {
              break L12;
            }
            _1106_i0 = (_1106_i0).plus(_dafny.ONE);
            let _1107_v0;
            _1107_v0 = new _dafny.CodePoint('c'.codePointAt(0));
            let _1108_v1;
            let _nw167 = new _module.C2();
            _nw167.__ctor(_1107_v0, new BigNumber(813));
            _1108_v1 = _nw167;
            (globalState).f1 = (_this).f11;
            let _1109_v2;
            _1109_v2 = _dafny.Seq.UnicodeFromString("ulwlw");
            let _1110_v3;
            _1110_v3 = _dafny.Seq.of(true);
            let _1111_v4;
            let _nw168 = Array((new BigNumber(9)).toNumber());
            _nw168[(_dafny.ZERO).toNumber()] = _module.__default.fm12((_this).f11, _1109_v2, globalState);
            _nw168[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_1110_v3, _module.__default.safeIndex((_this).f11, new BigNumber((_1110_v3).length)), (_this).f17);
            _nw168[(new BigNumber(2)).toNumber()] = _1110_v3;
            _nw168[(new BigNumber(3)).toNumber()] = _1110_v3;
            _nw168[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_1110_v3, _module.__default.safeIndex((_this).f11, new BigNumber((_1110_v3).length)), (_this).f17);
            _nw168[(new BigNumber(5)).toNumber()] = _1110_v3;
            _nw168[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1110_v3, _1110_v3);
            _nw168[(new BigNumber(7)).toNumber()] = _1110_v3;
            _nw168[(new BigNumber(8)).toNumber()] = _1110_v3;
            _1111_v4 = _nw168;
            let _index132 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1111_v4).length));
            (_1111_v4)[_index132] = _1110_v3;
            let _index133 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1111_v4).length));
            (_1111_v4)[_index133] = _dafny.Seq.Concat(_1110_v3, _1110_v3);
            let _1112_v5;
            _1112_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f17);
            let _1113_v6;
            _1113_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17);
            let _1114_v7;
            let _nw169 = Array((new BigNumber(9)).toNumber());
            _nw169[(_dafny.ZERO).toNumber()] = new BigNumber((_1112_v5).length);
            _nw169[(_dafny.ONE).toNumber()] = p0;
            _nw169[(new BigNumber(2)).toNumber()] = (_this).f11;
            _nw169[(new BigNumber(3)).toNumber()] = p0;
            _nw169[(new BigNumber(4)).toNumber()] = new BigNumber((_1113_v6).length);
            _nw169[(new BigNumber(5)).toNumber()] = new BigNumber(815);
            _nw169[(new BigNumber(6)).toNumber()] = (_this).f11;
            _nw169[(new BigNumber(7)).toNumber()] = new BigNumber(((_this).f12).length);
            _nw169[(new BigNumber(8)).toNumber()] = new BigNumber(745);
            _1114_v7 = _nw169;
            let _1115_v8;
            _1115_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1114_v7,!((new BigNumber(((_1111_v4)[_module.__default.safeIndex(new BigNumber(346), new BigNumber((_1111_v4).length))]).length)).isLessThanOrEqualTo((_this).f11)));
            _1115_v8 = (_1115_v8).update(_1114_v7, (_this).f17);
          }
        }
      }
      let _1116_v9;
      _1116_v9 = _dafny.Seq.UnicodeFromString("g");
      let _1117_v10;
      let _nw170 = new _module.C3();
      _nw170.__ctor((_module.D13.create_DC28(_1116_v9, (_this).f11, (_this).f17)).dtor_cf41, (p0).plus(new BigNumber((_1116_v9).length)), _module.__default.fm2(p0, p0, globalState), (_this).f12);
      _1117_v10 = _nw170;
      if ((false) || (((_this).f17) || ((_this).f17))) {
        let _1118_v11;
        _1118_v11 = _dafny.Seq.of((_this).f16, (_this).f16);
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(185)), function (_1119_i1) {
          return (_this).f16;
        }), _1118_v11)) {
          let _1120_v13;
          _1120_v13 = _dafny.Seq.of((_dafny.ZERO).minus((_1117_v10).f22));
          let _1121_v14;
          _1121_v14 = _dafny.Set.fromElements((_this).f11, (_1117_v10).f22);
          let _1122_v15;
          _1122_v15 = _dafny.Seq.of((_this).f17);
          let _1123_v16;
          _1123_v16 = _module.D0.create_DC0(new BigNumber(-405));
          let _1124_v17;
          _1124_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(811),(_this).f11);
          let _1125_v18;
          _1125_v18 = new _dafny.CodePoint('u'.codePointAt(0));
          let _1126_v19;
          let _nw171 = Array((new BigNumber(28)).toNumber());
          _nw171[(_dafny.ZERO).toNumber()] = (_this).f17;
          _nw171[(_dafny.ONE).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(2)).toNumber()] = !((_this).f17);
          _nw171[(new BigNumber(3)).toNumber()] = ((_dafny.ZERO).minus(_1117_v10.f23)).isLessThanOrEqualTo(_1117_v10.f23);
          _nw171[(new BigNumber(4)).toNumber()] = !((_this).f11).isEqualTo(new BigNumber((function () {
            let _coll55 = new _dafny.Set();
            for (const _compr_55 of (_1118_v11).Elements) {
              let _1127_v12 = _compr_55;
              if (_dafny.Seq.contains(_1118_v11, _1127_v12)) {
                _coll55.add(_1127_v12);
              }
            }
            return _coll55;
          }()).length));
          _nw171[(new BigNumber(5)).toNumber()] = _dafny.Seq.IsPrefixOf(_1116_v9, _1116_v9);
          _nw171[(new BigNumber(6)).toNumber()] = !((_this).f17) || (!(!(!((_this).f17))));
          _nw171[(new BigNumber(7)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1120_v13, _1120_v13);
          _nw171[(new BigNumber(8)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(9)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(10)).toNumber()] = (_dafny.Set.fromElements(p0)).IsDisjointFrom(_1121_v14);
          _nw171[(new BigNumber(11)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(12)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(13)).toNumber()] = true;
          _nw171[(new BigNumber(14)).toNumber()] = !_dafny.Seq.contains(_1122_v15, (_this).f17);
          _nw171[(new BigNumber(15)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(16)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(17)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(18)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(19)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(20)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(21)).toNumber()] = !(true);
          _nw171[(new BigNumber(22)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(23)).toNumber()] = (_1117_v10).fm5(_1123_v16, _1122_v15, (_1120_v13)[_module.__default.safeIndex(new BigNumber((_1124_v17).length), new BigNumber((_1120_v13).length))], _1125_v18, globalState);
          _nw171[(new BigNumber(24)).toNumber()] = !((_this).f17) || (true);
          _nw171[(new BigNumber(25)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(26)).toNumber()] = (_this).f17;
          _nw171[(new BigNumber(27)).toNumber()] = (_1117_v10.f23).isLessThan(new BigNumber((_1116_v9).length));
          _1126_v19 = _nw171;
          let _index134 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1126_v19).length));
          (_1126_v19)[_index134] = (_this).fm4((_1117_v10).fm6((_this).f17, _1117_v10.f23, globalState), _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_1124_v17), (_1117_v10).f22, true, globalState);
          let _index135 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1126_v19).length));
          (_1126_v19)[_index135] = !(_module.__default.safeModuloInt(new BigNumber((_1116_v9).length), _1117_v10.f23)).isEqualTo(_1117_v10.f23);
          _1125_v18 = _1125_v18;
          let _1128_v20;
          let _nw172 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1128_v20 = _nw172;
          _1128_v20 = _1128_v20;
          let _1129_v21;
          _1129_v21 = _1116_v9;
          _1129_v21 = _1129_v21;
          let _index136 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1126_v19).length));
          (_1126_v19)[_index136] = (_1117_v10).fm5(_1123_v16, _1122_v15, p0, _module.__default.fm39(p0, false, p0, globalState), globalState);
        } else {
          let _1130_v22;
          _1130_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1117_v10.f23,(_1117_v10).f22);
          let _1131_v23;
          _1131_v23 = _dafny.Seq.of(new BigNumber((_1130_v22).length));
          let _1132_v24;
          _1132_v24 = _dafny.Seq.of(((_1117_v10).f22).plus(new BigNumber((_1131_v23).length)));
          let _1133_v25;
          _1133_v25 = _module.D10.create_DC21((_this).f17, (_this).f17);
          let _1134_v28;
          _1134_v28 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f17);
          let _rhs210 = (_1117_v10).f22;
          let _rhs211 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(78)), ((_1135_v10) => function (_1136_i2) {
            return new BigNumber((function () {
              let _coll56 = new _dafny.Map();
              for (const _compr_56 of (function () {
                let _coll57 = new _dafny.Set();
                for (const _compr_57 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17))).Elements) {
                  let _1137_v27 = _compr_57;
                  if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17))).contains(_1137_v27)) {
                    _coll57.add(_1137_v27);
                  }
                }
                return _coll57;
              }()).Elements) {
                let _1138_v26 = _compr_56;
                if ((function () {
                  let _coll58 = new _dafny.Set();
                  for (const _compr_58 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17))).Elements) {
                    let _1139_v27 = _compr_58;
                    if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f17))).contains(_1139_v27)) {
                      _coll58.add(_1139_v27);
                    }
                  }
                  return _coll58;
                }()).contains(_1138_v26)) {
                  _coll56.push([_1138_v26,_1135_v10.f23]);
                }
              }
              return _coll56;
            }()).length);
          })(_1117_v10)), _1132_v24), _dafny.Seq.update(_dafny.Seq.of((_1117_v10).f22, (_1117_v10).f22, new BigNumber((_dafny.Seq.of(_1117_v10.f23, _1117_v10.f23, _1117_v10.f23)).length), new BigNumber(-313)), _module.__default.safeIndex(new BigNumber((_1134_v28).length), new BigNumber((_dafny.Seq.of((_1117_v10).f22, (_1117_v10).f22, new BigNumber((_dafny.Seq.of(_1117_v10.f23, _1117_v10.f23, _1117_v10.f23)).length), new BigNumber(-313))).length)), (_1117_v10).f22));
          let _rhs212 = _1133_v25;
          let _lhs166 = _1117_v10;
          _lhs166.f23 = _rhs210;
          _1132_v24 = _rhs211;
          _1133_v25 = _rhs212;
          let _1140_v29;
          let _nw173 = new _module.C3();
          _nw173.__ctor((_dafny.ZERO).minus((_this).fm3((_this).f11, (_this).f17, true, globalState)), (_this).f11, _module.__default.safeDivisionInt(new BigNumber((_1116_v9).length), (_dafny.ZERO).minus(new BigNumber((_1134_v28).length))), (_this).f12);
          _1140_v29 = _nw173;
          let _1141_v30;
          let _nw174 = Array((new BigNumber(25)).toNumber());
          _1141_v30 = _nw174;
          let _1142_v31;
          _1142_v31 = _dafny.MultiSet.fromElements((_this).f17);
          let _1143_v32;
          let _nw175 = new _module.C3();
          _nw175.__ctor(new BigNumber((_1142_v31).cardinality()), (_1140_v29).f22, new BigNumber((_1116_v9).length), (_this).f12);
          _1143_v32 = _nw175;
          let _index137 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_1141_v30).length));
          (_1141_v30)[_index137] = _1143_v32;
          let _1144_v33;
          _1144_v33 = _dafny.Seq.of(_1116_v9);
          let _1145_v34;
          _1145_v34 = _dafny.Seq.of(_1131_v23, _1131_v23);
          let _1146_v35;
          let _nw176 = Array((new BigNumber(7)).toNumber());
          _nw176[(_dafny.ZERO).toNumber()] = _1140_v29.f23;
          _nw176[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_1117_v10.f23);
          _nw176[(new BigNumber(2)).toNumber()] = (p0).plus((_1117_v10).f22);
          _nw176[(new BigNumber(3)).toNumber()] = (_1117_v10).f22;
          _nw176[(new BigNumber(4)).toNumber()] = new BigNumber((_1144_v33).length);
          _nw176[(new BigNumber(5)).toNumber()] = (((_this).f17) ? ((_1117_v10).f22) : (_1117_v10.f23));
          _nw176[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt((_1143_v32).f11, new BigNumber((_1145_v34).length));
          _1146_v35 = _nw176;
          let _1147_v36;
          _1147_v36 = new _dafny.CodePoint('o'.codePointAt(0));
          let _1148_v37;
          _1148_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1147_v36,(_this).f11);
          let _index138 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1146_v35).length));
          (_1146_v35)[_index138] = new BigNumber(((_1148_v37).Merge(_1148_v37)).length);
          let _1149_v39;
          _1149_v39 = _dafny.Set.fromElements((_1116_v9)[_module.__default.safeIndex((_1117_v10).f22, new BigNumber((_1116_v9).length))], _1147_v36, _1147_v36);
          let _1150_v40;
          _1150_v40 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (_1149_v39).Elements) {
              let _1151_v38 = _compr_59;
              if ((_1149_v39).contains(_1151_v38)) {
                _coll59.push([_1151_v38,(_this).f17]);
              }
            }
            return _coll59;
          }(),(_this).f17);
          let _1152_v41;
          _1152_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1147_v36,!((_this).f17));
          let _index139 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_1141_v30).length));
          let _index140 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1146_v35).length));
          let _rhs213 = _1143_v32;
          let _rhs214 = (_dafny.ZERO).minus((_1143_v32).f11);
          let _rhs215 = new BigNumber((_1130_v22).length);
          let _rhs216 = (_1150_v40).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1152_v41,(_this).f17)).update((_1152_v41).update(_1147_v36, (_this).f17), (_this).f17));
          let _lhs167 = _1141_v30;
          let _lhs168 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_1141_v30).length));
          let _lhs169 = _1117_v10;
          let _lhs170 = _1146_v35;
          let _lhs171 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1146_v35).length));
          _lhs167[_lhs168] = _rhs213;
          _lhs169.f23 = _rhs214;
          _lhs170[_lhs171] = _rhs215;
          _1150_v40 = _rhs216;
          let _1153_v42;
          let _nw177 = Array((new BigNumber(4)).toNumber()).fill(false);
          _1153_v42 = _nw177;
          let _1154_v43;
          _1154_v43 = _module.D0.create_DC0((_1146_v35)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_1146_v35).length))]);
          let _1155_v44;
          _1155_v44 = _dafny.Set.fromElements((_1146_v35)[_module.__default.safeIndex(new BigNumber(95), new BigNumber((_1146_v35).length))], (_1154_v43).dtor_cf0);
          let _index141 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1153_v42).length));
          (_1153_v42)[_index141] = (_1117_v10.f23).isEqualTo(new BigNumber((_1155_v44).length));
          let _index142 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1153_v42).length));
          (_1153_v42)[_index142] = (_this).f17;
          let _1156_v45;
          _1156_v45 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f11),(((_1153_v42)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_1153_v42).length))]) ? ((_this).f17) : ((_this).f17)));
          _1156_v45 = (_1156_v45).update(new BigNumber((_1116_v9).length), _module.__default.fm0(globalState));
        }
        let _1157_v46;
        let _nw178 = new _module.C2();
        _nw178.__ctor(new _dafny.CodePoint('g'.codePointAt(0)), (_this).f11);
        _1157_v46 = _nw178;
        let _1158_v47;
        _1158_v47 = _dafny.Seq.of(_1117_v10.f23, new BigNumber((_1116_v9).length), _1117_v10.f23);
        (globalState).f6 = (((_module.__default.fm45(false, (_this).f17, (_1117_v10).f22, new BigNumber((_1116_v9).length), globalState)).IsProperSubsetOf(_dafny.MultiSet.FromArray(_1158_v47))) ? ((_this).f17) : ((_this).f17));
        let _1159_v48;
        _1159_v48 = _module.D13.create_DC28(_1116_v9, new BigNumber(-728), (_this).f17);
        let _pat_let_tv37 = globalState;
        let _pat_let_tv38 = _1116_v9;
        _1159_v48 = function (_pat_let27_0) {
          return function (_1160_dt__update__tmp_h0) {
            return function (_pat_let28_0) {
              return function (_1161_dt__update_hcf40_h0) {
                return _module.D13.create_DC28(_1161_dt__update_hcf40_h0, (_1160_dt__update__tmp_h0).dtor_cf41, (_1160_dt__update__tmp_h0).dtor_cf42);
              }(_pat_let28_0);
            }(_dafny.Seq.Concat(_module.__default.fm1(_pat_let_tv37), _pat_let_tv38));
          }(_pat_let27_0);
        }(_module.D13.create_DC28(_dafny.Seq.UnicodeFromString("bsb"), (_this).f11, (_this).f17));
        (globalState).f6 = !(!(((_this).f17) || (true)));
      } else {
        (globalState).f1 = _module.__default.safeDivisionInt((_1117_v10).f22, p0);
        let _1162_v49;
        _1162_v49 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1163_v50;
        _1163_v50 = _dafny.Seq.of((_this).f17, (_this).f17);
        let _1164_v51;
        _1164_v51 = _dafny.Seq.of(_1116_v9, _1116_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(730)), ((_1165_v49) => function (_1166_i3) {
          return _1165_v49;
        })(_1162_v49)));
        let _1167_v52;
        _1167_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,_1117_v10.f23);
        let _1168_v53;
        _1168_v53 = _dafny.Seq.of(_1117_v10.f23, (_1117_v10).f22);
        _1162_v49 = _module.__default.fm20(_dafny.Seq.Concat(_1163_v50, _1163_v50), _module.__default.safeModuloInt((_1117_v10).f22, _1117_v10.f23), _dafny.areEqual(_dafny.Seq.of(new BigNumber(((_1164_v51)[_module.__default.safeIndex(new BigNumber((_1167_v52).length), new BigNumber((_1164_v51).length))]).length)), _1168_v53), globalState);
        let _1169_v54;
        let _init20 = function (_1170_i4) {
          return _dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f11));
        };
        let _nw179 = Array((new BigNumber(13)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw179.length); _i0_20++) {
          _nw179[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _1169_v54 = _nw179;
        _1169_v54 = _1169_v54;
        let _1171_v55;
        _1171_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f17);
        let _1172_v56;
        _1172_v56 = _dafny.Seq.of(_1171_v55);
        let _1173_v57;
        let _nw180 = new _module.C3();
        _nw180.__ctor(new BigNumber(((_1172_v56)[_module.__default.safeIndex((_1117_v10).f22, new BigNumber((_1172_v56).length))]).length), (_this).f11, _1117_v10.f23, (_this).f12);
        _1173_v57 = _nw180;
        let _rhs217 = new BigNumber(24);
        let _rhs218 = !(((_this).f17) && (!((_this).f17)));
        let _lhs172 = _1117_v10;
        let _lhs173 = globalState;
        _lhs172.f23 = _rhs217;
        _lhs173.f7 = _rhs218;
      }
      let _1174_v58;
      _1174_v58 = _dafny.Seq.of((_this).f17, (_this).f17);
      let _1175_v59;
      _1175_v59 = _dafny.Seq.of(new BigNumber((_1174_v58).length), _1117_v10.f23, (_this).f11);
      let _1176_v60;
      _1176_v60 = _dafny.MultiSet.fromElements((_this).f17);
      let _1177_v61;
      _1177_v61 = _1176_v60;
      let _1178_v62;
      let _nw181 = Array((new BigNumber(16)).toNumber());
      _nw181[(_dafny.ZERO).toNumber()] = (_this).f11;
      _nw181[(_dafny.ONE).toNumber()] = (_1175_v59)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(p0)).length), new BigNumber((_1175_v59).length))];
      _nw181[(new BigNumber(2)).toNumber()] = p0;
      _nw181[(new BigNumber(3)).toNumber()] = ((_1117_v10).f22).minus((_this).f11);
      _nw181[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1174_v58).length));
      _nw181[(new BigNumber(5)).toNumber()] = (_this).f11;
      _nw181[(new BigNumber(6)).toNumber()] = (_this).f11;
      _nw181[(new BigNumber(7)).toNumber()] = ((_this).f11).plus(p0);
      _nw181[(new BigNumber(8)).toNumber()] = new BigNumber(((_1177_v61)).cardinality());
      _nw181[(new BigNumber(9)).toNumber()] = (_this).f11;
      _nw181[(new BigNumber(10)).toNumber()] = _1117_v10.f23;
      _nw181[(new BigNumber(11)).toNumber()] = new BigNumber((_module.__default.fm11(globalState)).length);
      _nw181[(new BigNumber(12)).toNumber()] = (_1117_v10).f22;
      _nw181[(new BigNumber(13)).toNumber()] = (_1117_v10).fm3(p0, (_this).f17, (_this).f17, globalState);
      _nw181[(new BigNumber(14)).toNumber()] = p0;
      _nw181[(new BigNumber(15)).toNumber()] = (_this).f11;
      _1178_v62 = _nw181;
      let _index143 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1178_v62).length));
      (_1178_v62)[_index143] = new BigNumber(255);
      let _index144 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1178_v62).length));
      let _rhs219 = (_dafny.ZERO).minus((((_1176_v60).contains((_this).f17)) ? ((_1176_v60).get((_this).f17)) : (_1117_v10.f23)));
      let _rhs220 = (_1117_v10).f22;
      let _rhs221 = new BigNumber(297);
      let _lhs174 = _1178_v62;
      let _lhs175 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1178_v62).length));
      let _lhs176 = _1117_v10;
      let _lhs177 = _1117_v10;
      _lhs174[_lhs175] = _rhs219;
      _lhs176.f23 = _rhs220;
      _lhs177.f23 = _rhs221;
      let _index145 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_1178_v62).length));
      (_1178_v62)[_index145] = new BigNumber(600);
      (_1117_v10).f23 = (_1117_v10).f22;
      let _1179_v63;
      let _nw182 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1179_v63 = _nw182;
      r0 = _1179_v63;
      let _1180_v64;
      _1180_v64 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(825)), function (_1181_i5) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }),(_this).f16);
      r1 = _1180_v64;
      return [r0, r1];
    }
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      if (p1) {
        let _1182_v0;
        let _init21 = ((_1183_p1) => function (_1184_i0) {
          return _1183_p1;
        })(p1);
        let _nw183 = Array((new BigNumber(15)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw183.length); _i0_21++) {
          _nw183[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _1182_v0 = _nw183;
        let _1185_v1;
        _1185_v1 = _dafny.Set.fromElements(_module.__default.fm46(globalState));
        let _1186_v2;
        _1186_v2 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_1182_v0),_1185_v1);
        let _1187_v3;
        _1187_v3 = _dafny.MultiSet.fromElements(_1182_v0, _1182_v0, _1182_v0);
        _1186_v2 = (_1186_v2).update(_1187_v3, (_1185_v1).Union(_1185_v1));
        let _1188_v5;
        _1188_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _1189_v6;
        _1189_v6 = _module.D8.create_DC15(_1188_v5);
        let _1190_v7;
        _1190_v7 = _module.D8.create_DC17(_1189_v6);
        let _1191_v8;
        _1191_v8 = _module.D13.create_DC27(function () {
  let _coll60 = new _dafny.Map();
  for (const _compr_60 of (_dafny.Seq.of(_1190_v7)).Elements) {
    let _1192_v4 = _compr_60;
    if (_dafny.Seq.contains(_dafny.Seq.of(_1190_v7), _1192_v4)) {
      _coll60.push([_1192_v4,(_this).f11]);
    }
  }
  return _coll60;
}());
        let _1193_v9;
        _1193_v9 = _module.D13.create_DC29(_1191_v8);
        let _1194_v10;
        _1194_v10 = _module.D13.create_DC29(_1191_v8);
        let _1195_v11;
        _1195_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1190_v7,new BigNumber(992));
        _1194_v10 = _module.D13.create_DC29(_module.D13.create_DC27(_1195_v11));
        (globalState).f7 = p1;
        let _1196_v12;
        _1196_v12 = _dafny.Seq.of(p1, p0, p1);
        let _1197_v13;
        _1197_v13 = _dafny.Seq.of(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(p1), _1196_v12)), p0, ((p0) ? (!(p0)) : (p0)), !(((_this).f16).IsSubsetOf((_this).f16)), p0);
        let _1198_v14;
        _1198_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1188_v5);
        let _1199_v15;
        _1199_v15 = _dafny.Seq.UnicodeFromString("aw");
        _1196_v12 = _module.__default.fm12(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f11), new BigNumber((_1198_v14).length)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-390)), function (_1200_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }), _1199_v15), globalState);
        let _1201_v16;
        let _nw184 = new _module.C3();
        _nw184.__ctor((_this).f11, (_dafny.ZERO).minus(((_this).f11).minus(new BigNumber(267))), p2, (_this).f12);
        _1201_v16 = _nw184;
      } else {
        let _1202_v17;
        let _nw185 = new _module.C3();
        _nw185.__ctor(p2, new BigNumber(-116), ((((_this).f16).contains(p2)) ? (((_this).f16).get(p2)) : (p2)), (_this).f12);
        _1202_v17 = _nw185;
        let _1203_v18;
        let _init22 = ((_1204_p1, _1205_p0) => function (_1206_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe(_module.D11.create_DC24(new BigNumber(2), _dafny.Seq.of(_1204_p1, _1205_p0), new _dafny.CodePoint('n'.codePointAt(0))),_1204_p1);
        })(p1, p0);
        let _nw186 = Array((new BigNumber(24)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw186.length); _i0_22++) {
          _nw186[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _1203_v18 = _nw186;
        let _1207_v19;
        _1207_v19 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1208_v20;
        _1208_v20 = _module.D11.create_DC24(p2, _dafny.Seq.of(p1, p1), _1207_v19);
        let _1209_v21;
        _1209_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1208_v20,p1);
        let _index146 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1203_v18).length));
        (_1203_v18)[_index146] = _1209_v21;
        let _index147 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1203_v18).length));
        (_1203_v18)[_index147] = (_dafny.Map.Empty.slice().updateUnsafe(_1208_v20,p1)).Merge((_1209_v21).Merge(_1209_v21));
        let _1210_v22;
        _1210_v22 = _module.D0.create_DC1((_this).f17, _1202_v17.f23, !(p1));
        let _1211_v23;
        _1211_v23 = _dafny.Seq.of((_1202_v17).f22);
        let _1212_v24;
        _1212_v24 = _module.D7.create_DC12(_1211_v23);
        let _1213_v25;
        _1213_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm39((_1202_v17).f22, !(p1), (_1210_v22).dtor_cf2, globalState),_1212_v24);
        _1213_v25 = (_1213_v25).update(_1207_v19, _1212_v24);
        (_1202_v17).f23 = (_dafny.ZERO).minus((_1202_v17).f22);
        let _1214_v26;
        let _nw187 = new _module.C2();
        _nw187.__ctor(_1207_v19, new BigNumber(-968));
        _1214_v26 = _nw187;
      }
      let _1215_v27;
      _1215_v27 = _dafny.Seq.of(p1);
      let _1216_v28;
      _1216_v28 = _module.D6.create_DC10(_1215_v27);
      let _pat_let_tv39 = p2;
      let _pat_let_tv40 = p2;
      let _pat_let_tv41 = p2;
      r0 = (_dafny.ZERO).minus(function (_source17) {
        if (_source17.is_DC10) {
          let _1217___mcc_h0 = (_source17).cf13;
          let _1218_cf13 = _1217___mcc_h0;
          return _pat_let_tv39;
        } else if (_source17.is_DC11) {
          let _1219___mcc_h1 = (_source17).cf14;
          let _1220___mcc_h2 = (_source17).cf15;
          let _1221___mcc_h3 = (_source17).cf16;
          let _1222_cf16 = _1221___mcc_h3;
          let _1223_cf15 = _1220___mcc_h2;
          let _1224_cf14 = _1219___mcc_h1;
          return _module.__default.safeDivisionInt(new BigNumber(437), (_this).f11);
        } else {
          let _1225___mcc_h4 = (_source17).cf12;
          let _1226_cf12 = _1225___mcc_h4;
          return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f17,_pat_let_tv40)).length)).multipliedBy(_pat_let_tv41);
        }
      }(_1216_v28));
      r0 = p2;
      let _1227_v29;
      let _init23 = function (_1228_i3) {
        return (_this).f17;
      };
      let _nw188 = Array((new BigNumber(9)).toNumber());
      for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw188.length); _i0_23++) {
        _nw188[_i0_23] = _init23(new BigNumber(_i0_23));
      }
      _1227_v29 = _nw188;
      let _nw189 = Array((new BigNumber(5)).toNumber()).fill(false);
      _1227_v29 = _nw189;
      let _1229_v30;
      let _nw190 = Array((new BigNumber(14)).toNumber()).fill([]);
      _1229_v30 = _nw190;
      let _1230_v31;
      let _init24 = function (_1231_i4) {
        return _module.D2.create_DC5();
      };
      let _nw191 = Array((new BigNumber(6)).toNumber());
      for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw191.length); _i0_24++) {
        _nw191[_i0_24] = _init24(new BigNumber(_i0_24));
      }
      _1230_v31 = _nw191;
      let _index148 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1229_v30).length));
      (_1229_v30)[_index148] = _1230_v31;
      let _index149 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1229_v30).length));
      let _rhs222 = _1230_v31;
      let _rhs223 = p1;
      let _rhs224 = (_this).f17;
      let _rhs225 = p2;
      let _rhs226 = new BigNumber(93);
      let _lhs178 = _1229_v30;
      let _lhs179 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1229_v30).length));
      let _lhs180 = globalState;
      let _lhs181 = globalState;
      _lhs178[_lhs179] = _rhs222;
      r1 = _rhs223;
      _lhs180.f6 = _rhs224;
      r2 = _rhs225;
      _lhs181.f1 = _rhs226;
      let _1232_i5;
      _1232_i5 = _dafny.ZERO;
      L13: {
        while ((_this).f17) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1232_i5)) {
              break L13;
            }
            _1232_i5 = (_1232_i5).plus(_dafny.ONE);
            (globalState).f1 = (_this).f11;
            let _1233_v32;
            let _nw192 = new _module.C0();
            _nw192.__ctor((_1216_v28).dtor_cf13, (_this).f11);
            _1233_v32 = _nw192;
            let _1234_v33;
            _1234_v33 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1233_v32);
            let _1235_v34;
            let _nw193 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
            _1235_v34 = _nw193;
            let _1236_v35;
            _1236_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1234_v33).length),_1235_v34);
            _1236_v35 = (_1236_v35).update(p2, _1235_v34);
            let _1237_v36;
            _1237_v36 = new _dafny.CodePoint('b'.codePointAt(0));
            _1237_v36 = _1237_v36;
            (globalState).f1 = ((_dafny.ZERO).minus((new BigNumber((_dafny.Set.fromElements(new BigNumber(593), p2)).length)).minus(p2))).plus(new BigNumber(((_1233_v32).f20).length));
          }
        }
      }
      let _1238_v37;
      _1238_v37 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f17), _dafny.Seq.of(p1, (_this).f17), _1215_v27, _1215_v27);
      r0 = (_dafny.ZERO).minus(new BigNumber(((_1238_v37).update(_dafny.Seq.of((_this).f17, p0, true), _module.__default.abs((_dafny.ZERO).minus((_this).f11)))).cardinality()));
      r1 = !(((_this).f12).Difference(_module.__default.fm41(globalState))).contains(p0);
      let _1239_v38;
      _1239_v38 = _dafny.MultiSet.fromElements(p0, true, false, p1);
      let _1240_v39;
      _1240_v39 = _dafny.MultiSet.fromElements(_1239_v38);
      r2 = ((p1) ? ((((_1240_v39).contains(((_1239_v38).update(true, _module.__default.abs(p2))).update(false, _module.__default.abs(p2)))) ? ((_1240_v39).get(((_1239_v38).update(true, _module.__default.abs(p2))).update(false, _module.__default.abs(p2)))) : (p2))) : (new BigNumber(14)));
      return [r0, r1, r2];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(((new BigNumber(539)).plus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("usdxyfb")).length))))).minus(new BigNumber(933)));
    };
    m4(p0, p1, globalState) {
      let _this = this;
      let _1241_v0;
      _1241_v0 = new BigNumber(826);
      (globalState).f1 = _1241_v0;
      let _1242_v1;
      _1242_v1 = true;
      let _1243_v2;
      _1243_v2 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1244_v3;
      _1244_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v1,(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1241_v0,_1243_v2)).length), _1241_v0)).IsProperSubsetOf(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1241_v0))));
      if ((((_1244_v3).contains(_1242_v1)) ? ((_1244_v3).get(_1242_v1)) : (_1242_v1))) {
        let _1245_v4;
        _1245_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1241_v0,new BigNumber(821));
        _1245_v4 = (_1245_v4).update(_1241_v0, (_dafny.ZERO).minus(_1241_v0));
        let _1246_v5;
        _1246_v5 = _dafny.Seq.of(_1242_v1);
        (globalState).f6 = (_1242_v1) && ((_1246_v5)[_module.__default.safeIndex(_1241_v0, new BigNumber((_1246_v5).length))]);
        let _1247_v6;
        _1247_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1243_v2,_1241_v0);
        let _1248_v7;
        let _nw194 = new _module.C0();
        _nw194.__ctor(_1246_v5, (((_1247_v6).contains(_1243_v2)) ? ((_1247_v6).get(_1243_v2)) : (_1241_v0)));
        _1248_v7 = _nw194;
        let _1249_v8;
        let _nw195 = Array((new BigNumber(28)).toNumber());
        _1249_v8 = _nw195;
        let _1250_v9;
        _1250_v9 = _dafny.Set.fromElements(false);
        let _1251_v10;
        let _nw196 = new _module.C3();
        _nw196.__ctor(new BigNumber(611), (_dafny.ZERO).minus(_1241_v0), _1241_v0, _1250_v9);
        _1251_v10 = _nw196;
        let _index150 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1249_v8).length));
        (_1249_v8)[_index150] = _1251_v10;
        let _index151 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1249_v8).length));
        (_1249_v8)[_index151] = _1251_v10;
        (globalState).f1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(p0, p0), p0)).length);
      } else {
        (globalState).f1 = _1241_v0;
        let _1252_v11;
        _1252_v11 = _dafny.Seq.of(_module.__default.safeModuloInt(_1241_v0, _1241_v0), _1241_v0);
        _1252_v11 = _dafny.Seq.of(_1241_v0);
        let _1253_v12;
        let _nw197 = new _module.C3();
        _nw197.__ctor(_1241_v0, _1241_v0, _1241_v0, _dafny.Set.fromElements(_1242_v1));
        _1253_v12 = _nw197;
        let _1254_v13;
        let _nw198 = Array((new BigNumber(17)).toNumber());
        _nw198[(_dafny.ZERO).toNumber()] = _1253_v12;
        _nw198[(_dafny.ONE).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(2)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(3)).toNumber()] = ((_1242_v1) ? (_1253_v12) : (_1253_v12));
        _nw198[(new BigNumber(4)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(5)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(6)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(7)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(8)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(9)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(10)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(11)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(12)).toNumber()] = ((!(!(!(true)))) ? (_1253_v12) : (_1253_v12));
        _nw198[(new BigNumber(13)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(14)).toNumber()] = ((_1242_v1) ? (_1253_v12) : (_1253_v12));
        _nw198[(new BigNumber(15)).toNumber()] = _1253_v12;
        _nw198[(new BigNumber(16)).toNumber()] = _1253_v12;
        _1254_v13 = _nw198;
        let _index152 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1254_v13).length));
        (_1254_v13)[_index152] = _1253_v12;
        let _index153 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1254_v13).length));
        (_1254_v13)[_index153] = ((_module.__default.fm0(globalState)) ? (_1253_v12) : (_1253_v12));
        let _1255_v14;
        _1255_v14 = _module.D0.create_DC0(_1253_v12.f23);
        let _pat_let_tv42 = _1241_v0;
        _1255_v14 = function (_pat_let29_0) {
          return function (_1256_dt__update__tmp_h0) {
            return function (_pat_let30_0) {
              return function (_1257_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_1257_dt__update_hcf0_h0);
              }(_pat_let30_0);
            }(_pat_let_tv42);
          }(_pat_let29_0);
        }(_1255_v14);
        let _1258_v15;
        _1258_v15 = _dafny.Seq.of(_1242_v1, _1242_v1);
        let _1259_v16;
        _1259_v16 = _dafny.Seq.of(_1258_v15);
        (globalState).f1 = (new BigNumber((_1259_v16).length)).multipliedBy((_dafny.ZERO).minus(_1241_v0));
      }
      let _1260_v17;
      _1260_v17 = _module.D13.create_DC28(p0, _1241_v0, _1242_v1);
      if ((_1260_v17).dtor_cf42) {
        let _1261_v18;
        let _init25 = ((_1262_v3) => function (_1263_i0) {
          return (_1262_v3).Merge(_1262_v3);
        })(_1244_v3);
        let _nw199 = Array((new BigNumber(16)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw199.length); _i0_25++) {
          _nw199[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1261_v18 = _nw199;
        _1261_v18 = _1261_v18;
        (globalState).f1 = new BigNumber(498);
        let _1264_v19;
        _1264_v19 = _dafny.MultiSet.fromElements(_module.__default.fm0(globalState), _1242_v1);
        (globalState).f7 = !((_1264_v19).IsDisjointFrom(_1264_v19)) || (_1242_v1);
        (globalState).f6 = _1242_v1;
        let _1265_v20;
        let _nw200 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _1265_v20 = _nw200;
        let _index154 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1265_v20).length));
        (_1265_v20)[_index154] = _1241_v0;
        let _1266_v21;
        _1266_v21 = _dafny.Seq.of(_1242_v1);
        let _1267_v22;
        _1267_v22 = _dafny.MultiSet.fromElements(_1266_v21, _1266_v21, _1266_v21, _dafny.Seq.Concat(_dafny.Seq.update(_1266_v21, _module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_1266_v21).length)), _1242_v1), _dafny.Seq.update(_1266_v21, _module.__default.safeIndex(_1241_v0, new BigNumber((_1266_v21).length)), _1242_v1)));
        let _index155 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1265_v20).length));
        (_1265_v20)[_index155] = (((_1267_v22).contains(((!(_1242_v1)) ? (_1266_v21) : (_dafny.Seq.of(_1242_v1))))) ? ((_1267_v22).get(((!(_1242_v1)) ? (_1266_v21) : (_dafny.Seq.of(_1242_v1))))) : (_1241_v0));
      } else {
        let _1268_v23;
        let _init26 = ((_1269_v0) => function (_1270_i1) {
          return _module.__default.safeDivisionInt(_1270_i1, _1269_v0);
        })(_1241_v0);
        let _nw201 = Array((new BigNumber(6)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw201.length); _i0_26++) {
          _nw201[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1268_v23 = _nw201;
        let _1271_v24;
        _1271_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1268_v23,true);
        let _1272_v25;
        _1272_v25 = _module.D10.create_DC20(_1271_v24);
        let _source18 = _1272_v25;
        if (_source18.is_DC21) {
          let _1273___mcc_h0 = (_source18).cf28;
          let _1274___mcc_h1 = (_source18).cf29;
          let _1275_cf29 = _1274___mcc_h1;
          let _1276_cf28 = _1273___mcc_h0;
          let _1277_v26;
          _1277_v26 = _dafny.MultiSet.fromElements(_module.__default.fm2(_1241_v0, new BigNumber((p0).length), globalState));
          _1277_v26 = _1277_v26;
          let _1278_v27;
          _1278_v27 = _dafny.Seq.of(_1242_v1);
          let _1279_v28;
          _1279_v28 = _module.D6.create_DC10(_1278_v27);
          _1279_v28 = (((((_1244_v3).contains(_1275_cf29)) ? ((_1244_v3).get(_1275_cf29)) : (_module.__default.fm0(globalState)))) ? (_1279_v28) : (_1279_v28));
          let _1280_v29;
          let _init27 = ((_1281_cf28, _1282_p0, _1283_v1) => function (_1284_i2) {
            return ((_module.D14.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(_1281_cf28,_1282_p0))).dtor_cf44).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1283_v1,_1282_p0));
          })(_1276_cf28, p0, _1242_v1);
          let _nw202 = Array((new BigNumber(29)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw202.length); _i0_27++) {
            _nw202[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1280_v29 = _nw202;
          let _1285_v30;
          _1285_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v1,p0);
          let _index156 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1280_v29).length));
          (_1280_v29)[_index156] = _1285_v30;
          let _1286_v31;
          _1286_v31 = _dafny.Seq.of(_1285_v30);
          let _index157 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_1280_v29).length));
          (_1280_v29)[_index157] = (_1286_v31)[_module.__default.safeIndex(_1241_v0, new BigNumber((_1286_v31).length))];
          let _1287_v32;
          _1287_v32 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Concat(p0, p0)).length));
          let _1288_v33;
          _1288_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1241_v0,_1241_v0);
          let _1289_v34;
          _1289_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1243_v2,_module.__default.fm0(globalState));
          let _1290_v36;
          _1290_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1243_v2,_1241_v0);
          let _1291_v37;
          let _nw203 = Array((new BigNumber(29)).toNumber());
          _nw203[(_dafny.ZERO).toNumber()] = !(_1276_cf28) || (_1276_cf28);
          _nw203[(_dafny.ONE).toNumber()] = _1242_v1;
          _nw203[(new BigNumber(2)).toNumber()] = ((_this).fm9((_dafny.ZERO).minus(new BigNumber((_1288_v33).length)), new BigNumber(-717), _1242_v1, globalState)).isLessThan(_1241_v0);
          _nw203[(new BigNumber(3)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.of(function () {
            let _coll61 = new _dafny.Map();
            for (const _compr_61 of (_1290_v36).Keys.Elements) {
              let _1292_v35 = _compr_61;
              if ((_1290_v36).contains(_1292_v35)) {
                _coll61.push([_1292_v35,_1275_cf29]);
              }
            }
            return _coll61;
          }(), _1289_v34), _1289_v34);
          _nw203[(new BigNumber(4)).toNumber()] = !((_1278_v27)[_module.__default.safeIndex(_1241_v0, new BigNumber((_1278_v27).length))]);
          _nw203[(new BigNumber(5)).toNumber()] = _1276_cf28;
          _nw203[(new BigNumber(6)).toNumber()] = ((!(_1275_cf29)) ? (_1275_cf29) : ((_1278_v27)[_module.__default.safeIndex(_1241_v0, new BigNumber((_1278_v27).length))]));
          _nw203[(new BigNumber(7)).toNumber()] = true;
          _nw203[(new BigNumber(8)).toNumber()] = false;
          _nw203[(new BigNumber(9)).toNumber()] = ((_1242_v1) ? (_1276_cf28) : (_1242_v1));
          _nw203[(new BigNumber(10)).toNumber()] = !(_1276_cf28);
          _nw203[(new BigNumber(11)).toNumber()] = !(_1275_cf29) || (_module.__default.fm0(globalState));
          _nw203[(new BigNumber(12)).toNumber()] = !(_dafny.Seq.IsProperPrefixOf(p0, p0));
          _nw203[(new BigNumber(13)).toNumber()] = _1242_v1;
          _nw203[(new BigNumber(14)).toNumber()] = false;
          _nw203[(new BigNumber(15)).toNumber()] = _1242_v1;
          _nw203[(new BigNumber(16)).toNumber()] = (_1242_v1) && (_1276_cf28);
          _nw203[(new BigNumber(17)).toNumber()] = _1242_v1;
          _nw203[(new BigNumber(18)).toNumber()] = _1276_cf28;
          _nw203[(new BigNumber(19)).toNumber()] = ((_1242_v1) ? (_1275_cf29) : (_1242_v1));
          _nw203[(new BigNumber(20)).toNumber()] = true;
          _nw203[(new BigNumber(21)).toNumber()] = (((_1244_v3).contains(_1242_v1)) ? ((_1244_v3).get(_1242_v1)) : (_1242_v1));
          _nw203[(new BigNumber(22)).toNumber()] = _dafny.Seq.IsPrefixOf(p0, _dafny.Seq.UnicodeFromString("ley"));
          _nw203[(new BigNumber(23)).toNumber()] = _1276_cf28;
          _nw203[(new BigNumber(24)).toNumber()] = _1276_cf28;
          _nw203[(new BigNumber(25)).toNumber()] = !(_1242_v1);
          _nw203[(new BigNumber(26)).toNumber()] = (_1242_v1) && (_1276_cf28);
          _nw203[(new BigNumber(27)).toNumber()] = ((_1276_cf28) ? (_1276_cf28) : (true));
          _nw203[(new BigNumber(28)).toNumber()] = _1275_cf29;
          _1291_v37 = _nw203;
          let _index158 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1291_v37).length));
          (_1291_v37)[_index158] = _1242_v1;
          let _index159 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1291_v37).length));
          let _rhs227 = (_1287_v32).Difference(_1287_v32);
          let _rhs228 = _module.__default.fm0(globalState);
          let _lhs182 = _1291_v37;
          let _lhs183 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1291_v37).length));
          _1287_v32 = _rhs227;
          _lhs182[_lhs183] = _rhs228;
        } else {
          let _1293___mcc_h2 = (_source18).cf27;
          let _1294_cf27 = _1293___mcc_h2;
          let _1295_v38;
          _1295_v38 = _dafny.Seq.of(_1242_v1);
          let _1296_v39;
          let _nw204 = new _module.C1();
          _nw204.__ctor(_1242_v1, (_1295_v38)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_1295_v38).length))]);
          _1296_v39 = _nw204;
          (globalState).f7 = (_1296_v39).f19;
          let _1297_v40;
          _1297_v40 = _module.D8.create_DC16(_1242_v1, (_this).fm9(_1241_v0, _1241_v0, _1242_v1, globalState));
          let _1298_v41;
          _1298_v41 = _module.D8.create_DC17(_1297_v40);
          let _1299_v42;
          _1299_v42 = _module.D8.create_DC17(_1297_v40);
          let _1300_v43;
          _1300_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1299_v42,_1241_v0);
          let _1301_v44;
          _1301_v44 = _module.D13.create_DC27(_1300_v43);
          _1301_v44 = _1301_v44;
          (globalState).f7 = true;
        }
        (globalState).f1 = _1241_v0;
        let _1302_v45;
        _1302_v45 = _dafny.Seq.of((_dafny.ZERO).minus(_1241_v0));
        _1242_v1 = _dafny.Seq.IsProperPrefixOf(_1302_v45, _dafny.Seq.of(_1241_v0));
        let _1303_v46;
        _1303_v46 = _dafny.Seq.of(_1242_v1, !(_1242_v1), _1242_v1, _1242_v1);
        let _1304_v47;
        _1304_v47 = _1303_v46;
        let _source19 = _1303_v46;
        let _1305___mcc_h3 = _source19;
        let _1306_cf10 = _1305___mcc_h3;
        let _index160 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1268_v23).length));
        (_1268_v23)[_index160] = _1241_v0;
        let _index161 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1268_v23).length));
        let _rhs229 = _1241_v0;
        let _rhs230 = _1241_v0;
        let _rhs231 = (_dafny.ZERO).minus(new BigNumber((_1302_v45).length));
        let _lhs184 = globalState;
        let _lhs185 = globalState;
        let _lhs186 = _1268_v23;
        let _lhs187 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1268_v23).length));
        _lhs184.f1 = _rhs229;
        _lhs185.f1 = _rhs230;
        _lhs186[_lhs187] = _rhs231;
        let _1307_v48;
        _1307_v48 = _module.D8.create_DC16(_1242_v1, _1241_v0);
        _1307_v48 = _1307_v48;
        let _1308_v49;
        let _nw205 = new _module.C1();
        _nw205.__ctor(_1242_v1, _1242_v1);
        _1308_v49 = _nw205;
        let _1309_v50;
        _1309_v50 = _dafny.Map.Empty.slice().updateUnsafe((_1268_v23)[_module.__default.safeIndex(new BigNumber(251), new BigNumber((_1268_v23).length))],_1241_v0);
        let _1310_v51;
        let _nw206 = new _module.C1();
        _nw206.__ctor(_dafny.Seq.IsPrefixOf(p0, p0), !(_1309_v50).equals(_1309_v50));
        _1310_v51 = _nw206;
        let _1311_v52;
        _1311_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1241_v0,_1268_v23);
        _1311_v52 = (_1311_v52).update(_1241_v0, _1268_v23);
      }
      let _1312_i3;
      _1312_i3 = _dafny.ZERO;
      L14: {
        while (!(!(_1241_v0).isEqualTo(_1241_v0))) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1312_i3)) {
              break L14;
            }
            _1312_i3 = (_1312_i3).plus(_dafny.ONE);
            let _1313_v53;
            let _nw207 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
            _1313_v53 = _nw207;
            let _index162 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1313_v53).length));
            (_1313_v53)[_index162] = _1244_v3;
            let _index163 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1313_v53).length));
            (_1313_v53)[_index163] = _module.__default.fm11(globalState);
            if (!(_1242_v1) || ((_1242_v1) === (_1242_v1))) {
              let _1314_v54;
              let _nw208 = new _module.C2();
              _nw208.__ctor(_1243_v2, _1241_v0);
              _1314_v54 = _nw208;
              let _1315_v55;
              _1315_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v1,_1314_v54);
              let _1316_v56;
              let _init28 = ((_1317_v1) => function (_1318_i4) {
                return _1317_v1;
              })(_1242_v1);
              let _nw209 = Array((new BigNumber(14)).toNumber());
              for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw209.length); _i0_28++) {
                _nw209[_i0_28] = _init28(new BigNumber(_i0_28));
              }
              _1316_v56 = _nw209;
              let _1319_v57;
              let _nw210 = Array((new BigNumber(13)).toNumber());
              _nw210[(_dafny.ZERO).toNumber()] = _1316_v56;
              _nw210[(_dafny.ONE).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(2)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(3)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(4)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(5)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(6)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(7)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(8)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(9)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(10)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(11)).toNumber()] = _1316_v56;
              _nw210[(new BigNumber(12)).toNumber()] = _1316_v56;
              _1319_v57 = _nw210;
              let _index164 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_1316_v56).length));
              (_1316_v56)[_index164] = _1242_v1;
              let _index165 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_1316_v56).length));
              let _rhs232 = _1315_v55;
              let _rhs233 = _1319_v57;
              let _rhs234 = _1241_v0;
              let _rhs235 = _1242_v1;
              let _lhs188 = _1316_v56;
              let _lhs189 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_1316_v56).length));
              _1315_v55 = _rhs232;
              _1319_v57 = _rhs233;
              _1241_v0 = _rhs234;
              _lhs188[_lhs189] = _rhs235;
              let _1320_v58;
              _1320_v58 = _dafny.Seq.UnicodeFromString("hnmuw");
              _1320_v58 = p0;
              (globalState).f7 = (_1316_v56)[_module.__default.safeIndex(new BigNumber(45), new BigNumber((_1316_v56).length))];
              let _1321_v59;
              _1321_v59 = _dafny.Seq.of(_1320_v58, p0, p0, p0);
              (globalState).f1 = (_1314_v54).fm3(new BigNumber((_1321_v59).length), (_1241_v0).isLessThan(_1241_v0), false, globalState);
              let _1322_v60;
              _1322_v60 = _dafny.MultiSet.fromElements(_1241_v0);
              let _1323_v61;
              _1323_v61 = _dafny.Set.fromElements(_1242_v1);
              let _1324_v62;
              _1324_v62 = _module.D7.create_DC13(_1323_v61, _1241_v0, _1241_v0);
              let _1325_v63;
              let _nw211 = new _module.C4();
              _nw211.__ctor(_1322_v60, false, (_1324_v62).dtor_cf18, (_1241_v0).minus(_1241_v0));
              _1325_v63 = _nw211;
            } else {
              let _1326_v64;
              _1326_v64 = _dafny.MultiSet.fromElements(_module.__default.fm2(_1241_v0, _1241_v0, globalState));
              _1241_v0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_1241_v0, _module.__default.safeModuloInt(_1241_v0, (((_1326_v64).contains(_1241_v0)) ? ((_1326_v64).get(_1241_v0)) : (new BigNumber(-984))))));
              let _1327_v65;
              _1327_v65 = _module.D9.create_DC19();
              _1327_v65 = _1327_v65;
              _1242_v1 = _1242_v1;
              let _1328_v66;
              _1328_v66 = _dafny.Seq.UnicodeFromString("rxlji");
              _1328_v66 = _dafny.Seq.update(_dafny.Seq.Concat(_1328_v66, p0), _module.__default.safeIndex(_1241_v0, new BigNumber((_dafny.Seq.Concat(_1328_v66, p0)).length)), _1243_v2);
              let _1329_v67;
              let _nw212 = Array((new BigNumber(2)).toNumber()).fill(false);
              _1329_v67 = _nw212;
              let _1330_v68;
              _1330_v68 = _dafny.Set.fromElements(true, _1242_v1, _1242_v1);
              let _index166 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_1329_v67).length));
              (_1329_v67)[_index166] = (_1330_v68).IsDisjointFrom(_dafny.Set.fromElements(_1242_v1));
              let _1331_v69;
              _1331_v69 = _dafny.Seq.of(_1241_v0);
              let _1332_v70;
              _1332_v70 = _dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_1331_v69)).cardinality()))), _1241_v0, _1241_v0, _1241_v0);
              let _1333_v71;
              let _nw213 = Array((new BigNumber(5)).toNumber());
              _nw213[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_1241_v0);
              _nw213[(_dafny.ONE).toNumber()] = _1241_v0;
              _nw213[(new BigNumber(2)).toNumber()] = (_1241_v0).multipliedBy(_1241_v0);
              _nw213[(new BigNumber(3)).toNumber()] = _1241_v0;
              _nw213[(new BigNumber(4)).toNumber()] = (_1241_v0).plus(new BigNumber((_1332_v70).length));
              _1333_v71 = _nw213;
              let _index167 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1333_v71).length));
              (_1333_v71)[_index167] = _1241_v0;
              let _1334_v72;
              _1334_v72 = _dafny.Seq.of(!(_1242_v1), _1242_v1);
              let _1335_v73;
              _1335_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1241_v0,_1334_v72);
              let _1336_v74;
              _1336_v74 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).fm9(new BigNumber((_1335_v73).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1330_v68).length))), _1242_v1, globalState)),(_dafny.ZERO).minus(_1241_v0));
              let _1337_v75;
              _1337_v75 = _dafny.Seq.of(_1336_v74);
              let _1338_v76;
              _1338_v76 = _dafny.MultiSet.fromElements(_module.__default.fm39(_1241_v0, _1242_v1, _1241_v0, globalState));
              let _index168 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_1329_v67).length));
              let _index169 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1333_v71).length));
              let _rhs236 = (_1338_v76).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1328_v66, _module.__default.fm1(globalState))));
              let _rhs237 = !(_1242_v1) || (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_1334_v72, _module.__default.safeIndex(_1241_v0, new BigNumber((_1334_v72).length)), _1242_v1), (((_1335_v73).contains(_1241_v0)) ? ((_1335_v73).get(_1241_v0)) : (_1334_v72))));
              let _rhs238 = (_dafny.ZERO).minus(_1241_v0);
              let _rhs239 = (_module.__default.safeModuloInt(_1241_v0, new BigNumber(-527))).multipliedBy((new BigNumber(338)).minus(new BigNumber((p0).length)));
              let _rhs240 = _dafny.Seq.update(_1337_v75, _module.__default.safeIndex(_1241_v0, new BigNumber((_1337_v75).length)), (_1336_v74).Merge(_1336_v74));
              let _lhs190 = _1329_v67;
              let _lhs191 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_1329_v67).length));
              let _lhs192 = _1333_v71;
              let _lhs193 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1333_v71).length));
              let _lhs194 = globalState;
              _1242_v1 = _rhs236;
              _lhs190[_lhs191] = _rhs237;
              _lhs192[_lhs193] = _rhs238;
              _lhs194.f1 = _rhs239;
              _1337_v75 = _rhs240;
            }
            (globalState).f7 = !(_1242_v1);
            if (_1242_v1) {
              let _1339_v77;
              _1339_v77 = _dafny.Seq.of(_1242_v1, _1242_v1);
              let _1340_v78;
              _1340_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1241_v0,_1241_v0);
              let _1341_v79;
              _1341_v79 = p1;
              let _1342_v80;
              _1342_v80 = _dafny.Set.fromElements(_1243_v2);
              let _1343_v81;
              _1343_v81 = _dafny.Map.Empty.slice().updateUnsafe(_1341_v79,_1342_v80);
              let _1344_v82;
              let _nw214 = Array((new BigNumber(10)).toNumber());
              _nw214[(_dafny.ZERO).toNumber()] = _1241_v0;
              _nw214[(_dafny.ONE).toNumber()] = new BigNumber(-279);
              _nw214[(new BigNumber(2)).toNumber()] = _1241_v0;
              _nw214[(new BigNumber(3)).toNumber()] = _1241_v0;
              _nw214[(new BigNumber(4)).toNumber()] = _1241_v0;
              _nw214[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_1241_v0);
              _nw214[(new BigNumber(6)).toNumber()] = _1241_v0;
              _nw214[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.update(_1339_v77, _module.__default.safeIndex(_module.__default.fm2(_1241_v0, new BigNumber((_1340_v78).length), globalState), new BigNumber((_1339_v77).length)), !(_1242_v1))).length);
              _nw214[(new BigNumber(8)).toNumber()] = _1241_v0;
              _nw214[(new BigNumber(9)).toNumber()] = new BigNumber(((((_1343_v81).contains(_1341_v79)) ? ((_1343_v81).get(_1341_v79)) : (_dafny.Set.fromElements(_1243_v2)))).length);
              _1344_v82 = _nw214;
              let _index170 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length));
              (p1)[_index170] = _1344_v82;
              let _index171 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length));
              (p1)[_index171] = _1344_v82;
              let _1345_v83;
              let _nw215 = Array((new BigNumber(14)).toNumber()).fill(false);
              _1345_v83 = _nw215;
              let _1346_v84;
              _1346_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v1,_1345_v83);
              let _1347_v86;
              _1347_v86 = _module.D15.create_DC33(_1346_v84);
              let _1348_v87;
              _1348_v87 = _dafny.Map.Empty.slice().updateUnsafe(function () {
                let _coll62 = new _dafny.Map();
                for (const _compr_62 of _dafny.IntegerRange(new BigNumber(-81), new BigNumber(373))) {
                  let _1349_v85 = _compr_62;
                  if (((new BigNumber(-81)).isLessThanOrEqualTo(_1349_v85)) && ((_1349_v85).isLessThan(new BigNumber(373)))) {
                    _coll62.push([_module.__default.safeModuloInt(_1349_v85, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1242_v1,_1241_v0)).length)),new BigNumber((_dafny.MultiSet.fromElements(_1241_v0)).cardinality())]);
                  }
                }
                return _coll62;
              }(),(_1347_v86).dtor_cf49);
              let _1350_v88;
              _1350_v88 = _module.D16.create_DC38(_1340_v78);
              _1346_v84 = (((_1348_v87).contains((_1350_v88).dtor_cf61)) ? ((_1348_v87).get((_1350_v88).dtor_cf61)) : (_1346_v84));
              let _1351_v89;
              _1351_v89 = _dafny.Seq.of((_dafny.ZERO).minus(_1241_v0), new BigNumber(-243));
              (globalState).f7 = (new BigNumber((_1351_v89).length)).isLessThan(_1241_v0);
              let _arr0 = (p1)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length))];
              let _index172 = _module.__default.safeIndex(new BigNumber(724), new BigNumber(((p1)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length))]).length));
              _arr0[_index172] = new BigNumber(161);
              let _arr1 = (p1)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length))];
              let _index173 = _module.__default.safeIndex(new BigNumber(724), new BigNumber(((p1)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length))]).length));
              let _rhs241 = (_1241_v0).multipliedBy(_1241_v0);
              let _rhs242 = _1243_v2;
              let _lhs195 = (p1)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length))];
              let _lhs196 = _module.__default.safeIndex(new BigNumber(724), new BigNumber(((p1)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((p1).length))]).length));
              _lhs195[_lhs196] = _rhs241;
              _1243_v2 = _rhs242;
              let _1352_v90;
              _1352_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1244_v3).Merge((_1313_v53)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_1313_v53).length))])).length),_1351_v89);
              _1352_v90 = _1352_v90;
            } else {
              (globalState).f1 = _1241_v0;
              let _1353_v91;
              let _nw216 = Array((new BigNumber(7)).toNumber()).fill(false);
              _1353_v91 = _nw216;
              let _index174 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1353_v91).length));
              (_1353_v91)[_index174] = _1242_v1;
              let _1354_v92;
              _1354_v92 = _dafny.Map.Empty.slice().updateUnsafe(_1241_v0,!(false));
              let _1355_v93;
              _1355_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v1,_1241_v0);
              let _1356_v94;
              _1356_v94 = _dafny.Seq.of(_1242_v1);
              let _1357_v95;
              _1357_v95 = _dafny.Seq.of(_1241_v0);
              let _1358_v96;
              let _nw217 = Array((new BigNumber(17)).toNumber());
              _nw217[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_1241_v0).plus(new BigNumber(755)));
              _nw217[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1354_v92).length), _1241_v0);
              _nw217[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_1241_v0);
              _nw217[(new BigNumber(3)).toNumber()] = _1241_v0;
              _nw217[(new BigNumber(4)).toNumber()] = (new BigNumber((p0).length)).multipliedBy(_1241_v0);
              _nw217[(new BigNumber(5)).toNumber()] = (((_1355_v93).contains(_1242_v1)) ? ((_1355_v93).get(_1242_v1)) : ((_dafny.ZERO).minus(_1241_v0)));
              _nw217[(new BigNumber(6)).toNumber()] = _1241_v0;
              _nw217[(new BigNumber(7)).toNumber()] = new BigNumber(784);
              _nw217[(new BigNumber(8)).toNumber()] = new BigNumber((_1356_v94).length);
              _nw217[(new BigNumber(9)).toNumber()] = _1241_v0;
              _nw217[(new BigNumber(10)).toNumber()] = _1241_v0;
              _nw217[(new BigNumber(11)).toNumber()] = _1241_v0;
              _nw217[(new BigNumber(12)).toNumber()] = (_1241_v0).plus(new BigNumber((_1356_v94).length));
              _nw217[(new BigNumber(13)).toNumber()] = (_1357_v95)[_module.__default.safeIndex(_1241_v0, new BigNumber((_1357_v95).length))];
              _nw217[(new BigNumber(14)).toNumber()] = new BigNumber((_1356_v94).length);
              _nw217[(new BigNumber(15)).toNumber()] = new BigNumber((p0).length);
              _nw217[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_1241_v0);
              _1358_v96 = _nw217;
              let _1359_v97;
              let _nw218 = Array((new BigNumber(14)).toNumber()).fill([]);
              _1359_v97 = _nw218;
              let _index175 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1359_v97).length));
              (_1359_v97)[_index175] = _1353_v91;
              let _1360_v98;
              _1360_v98 = _dafny.Seq.UnicodeFromString("isyb");
              let _1361_v99;
              _1361_v99 = _dafny.Set.fromElements(_1241_v0, _1241_v0);
              let _index176 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1353_v91).length));
              let _index177 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1359_v97).length));
              let _rhs243 = false;
              let _rhs244 = ((!(_1242_v1)) ? (_1358_v96) : (_1358_v96));
              let _rhs245 = _1353_v91;
              let _rhs246 = _dafny.Seq.update(_module.__default.fm1(globalState), _module.__default.safeIndex(new BigNumber(((_1361_v99).Difference(_1361_v99)).length), new BigNumber((_module.__default.fm1(globalState)).length)), _1243_v2);
              let _lhs197 = _1353_v91;
              let _lhs198 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1353_v91).length));
              let _lhs199 = _1359_v97;
              let _lhs200 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1359_v97).length));
              _lhs197[_lhs198] = _rhs243;
              _1358_v96 = _rhs244;
              _lhs199[_lhs200] = _rhs245;
              _1360_v98 = _rhs246;
              let _1362_v100;
              _1362_v100 = _dafny.Set.fromElements((_1359_v97)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_1359_v97).length))], (_1359_v97)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_1359_v97).length))]);
              let _1363_v101;
              _1363_v101 = _dafny.Map.Empty.slice().updateUnsafe((_1362_v100).Union(_1362_v100),_1241_v0);
              _1363_v101 = (_1363_v101).update(_1362_v100, new BigNumber((_1355_v93).length));
              _1358_v96 = _1358_v96;
              let _1364_v102;
              let _nw219 = new _module.C0();
              _nw219.__ctor(_1356_v94, (_dafny.ZERO).minus(((!(_1242_v1)) ? (_1241_v0) : (_1241_v0))));
              _1364_v102 = _nw219;
            }
          }
        }
      }
      let _1365_v103;
      _1365_v103 = _dafny.Map.Empty.slice().updateUnsafe(false,_1241_v0);
      let _1366_v104;
      _1366_v104 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1365_v103).length)), _1241_v0, _1241_v0);
      let _1367_v105;
      _1367_v105 = _dafny.MultiSet.fromElements(_1241_v0);
      (globalState).f1 = (((_1242_v1) ? (_1241_v0) : (new BigNumber((_1366_v104).length)))).plus((((_1367_v105).contains(_1241_v0)) ? ((_1367_v105).get(_1241_v0)) : (_1241_v0)));
      let _1368_v106;
      _1368_v106 = _dafny.Seq.of(_1242_v1, _1242_v1);
      let _1369_v107;
      let _nw220 = new _module.C2();
      _nw220.__ctor(_1243_v2, _1241_v0);
      _1369_v107 = _nw220;
      let _1370_v108;
      _1370_v108 = _dafny.MultiSet.fromElements(_1369_v107);
      let _1371_v109;
      let _nw221 = Array((new BigNumber(14)).toNumber());
      _nw221[(_dafny.ZERO).toNumber()] = _1241_v0;
      _nw221[(_dafny.ONE).toNumber()] = new BigNumber((_1365_v103).length);
      _nw221[(new BigNumber(2)).toNumber()] = new BigNumber((_1368_v106).length);
      _nw221[(new BigNumber(3)).toNumber()] = new BigNumber((_1365_v103).length);
      _nw221[(new BigNumber(4)).toNumber()] = new BigNumber(83);
      _nw221[(new BigNumber(5)).toNumber()] = new BigNumber(347);
      _nw221[(new BigNumber(6)).toNumber()] = _1241_v0;
      _nw221[(new BigNumber(7)).toNumber()] = _1241_v0;
      _nw221[(new BigNumber(8)).toNumber()] = _1241_v0;
      _nw221[(new BigNumber(9)).toNumber()] = _1241_v0;
      _nw221[(new BigNumber(10)).toNumber()] = new BigNumber(956);
      _nw221[(new BigNumber(11)).toNumber()] = _1241_v0;
      _nw221[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(_1241_v0);
      _nw221[(new BigNumber(13)).toNumber()] = new BigNumber((_1370_v108).cardinality());
      _1371_v109 = _nw221;
      let _1372_v110;
      let _nw222 = Array((new BigNumber(9)).toNumber()).fill([]);
      _1372_v110 = _nw222;
      let _1373_v111;
      _1373_v111 = _dafny.Map.Empty.slice().updateUnsafe(_1371_v109,_1372_v110);
      (globalState).f7 = (!(_1242_v1)) && ((_1373_v111).contains(_1371_v109));
      return;
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f12 = _dafny.Set.Empty;
      this._f11 = _dafny.ZERO;
      this._f15 = _module.D2.Default();
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f15, f12, f11) {
      let _this = this;
      (_this)._f15 = f15;
      (_this)._f12 = f12;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (false) {
        return (new BigNumber((_dafny.Seq.UnicodeFromString("uoywkyi")).length)).isLessThanOrEqualTo((_this).f11);
      } else {
        return !(!(true));
      }
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((_this).f11, new BigNumber((((!(!(true))) ? (_dafny.Map.Empty.slice().updateUnsafe(false,true)) : (_dafny.Map.Empty.slice().updateUnsafe(false,false)))).length));
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return (new BigNumber((function () {
        let _coll63 = new _dafny.Map();
        for (const _compr_63 of (_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)))).Elements) {
          let _1374_v0 = _compr_63;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0))), _1374_v0)) {
            _coll63.push([_1374_v0,_dafny.Seq.UnicodeFromString("mhdueatrb")]);
          }
        }
        return _coll63;
      }()).length)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-288)), function (_1375_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(941)), function (_1376_i1) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }))).length)));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source20 = (_this).f15;
      if (_source20.is_DC4) {
        let _1377___mcc_h0 = (_source20).cf6;
        let _1378___mcc_h1 = (_source20).cf7;
        let _1379___mcc_h2 = (_source20).cf8;
        let _1380_cf8 = _1379___mcc_h2;
        let _1381_cf7 = _1378___mcc_h1;
        let _1382_cf6 = _1377___mcc_h0;
        return true;
      } else if (_source20.is_DC5) {
        return (_module.D0.create_DC1(true, new BigNumber(252), true)).dtor_cf1;
      } else {
        let _1383___mcc_h3 = (_source20).cf5;
        let _1384_cf5 = _1383___mcc_h3;
        return ((_this).f12).IsDisjointFrom(_dafny.Set.fromElements((_module.D0.create_DC1(false, (_this).f11, true)).dtor_cf3));
      }
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let _1385_v0;
      let _nw223 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
      _1385_v0 = _nw223;
      let _1386_v1;
      let _nw224 = Array((new BigNumber(13)).toNumber()).fill(false);
      _1386_v1 = _nw224;
      let _1387_v2;
      _1387_v2 = _dafny.Set.fromElements(_1386_v1);
      let _index178 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1385_v0).length));
      (_1385_v0)[_index178] = _1387_v2;
      let _index179 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1385_v0).length));
      (_1385_v0)[_index179] = _1387_v2;
      let _1388_v3;
      _1388_v3 = true;
      let _1389_i0;
      _1389_i0 = _dafny.ZERO;
      L15: {
        while (_1388_v3) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1389_i0)) {
              break L15;
            }
            _1389_i0 = (_1389_i0).plus(_dafny.ONE);
            let _1390_v4;
            _1390_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1388_v3,_1388_v3);
            _1390_v4 = _1390_v4;
            let _1391_v5;
            _1391_v5 = _dafny.MultiSet.fromElements((_this).f11);
            let _1392_v6;
            _1392_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1391_v5,!(_1388_v3) || (_1388_v3));
            (globalState).f1 = new BigNumber((_1392_v6).length);
            _1388_v3 = _module.__default.fm0(globalState);
            (globalState).f6 = _1388_v3;
          }
        }
      }
      (globalState).f1 = (_this).f11;
      if (((_1388_v3) ? (_1388_v3) : (_1388_v3))) {
        let _1393_v7;
        _1393_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,new BigNumber(194));
        let _1394_v8;
        let _nw225 = Array((new BigNumber(6)).toNumber());
        _nw225[(_dafny.ZERO).toNumber()] = new BigNumber(-194);
        _nw225[(_dafny.ONE).toNumber()] = (_this).f11;
        _nw225[(new BigNumber(2)).toNumber()] = (_this).fm3((_this).f11, _1388_v3, _1388_v3, globalState);
        _nw225[(new BigNumber(3)).toNumber()] = (_this).f11;
        _nw225[(new BigNumber(4)).toNumber()] = (_this).f11;
        _nw225[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt((((_1393_v7).contains((_this).f11)) ? ((_1393_v7).get((_this).f11)) : ((_this).f11)), (_this).f11);
        _1394_v8 = _nw225;
        let _index180 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1394_v8).length));
        (_1394_v8)[_index180] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f11, (_this).f11));
        let _1395_v9;
        _1395_v9 = _dafny.Seq.of(!(_1388_v3), _1388_v3);
        let _1396_v10;
        _1396_v10 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1388_v3));
        let _1397_v11;
        _1397_v11 = _dafny.Seq.of((_this).f11, (_this).f11, (_this).f11);
        let _1398_v12;
        _1398_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f12).length),false);
        let _index181 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1394_v8).length));
        let _rhs247 = (_this).f11;
        let _rhs248 = (_1395_v9)[_module.__default.safeIndex(new BigNumber((((_1396_v10)[_module.__default.safeIndex((_this).fm3(new BigNumber((_1397_v11).length), _1388_v3, _1388_v3, globalState), new BigNumber((_1396_v10).length))]).Merge(_1398_v12)).length), new BigNumber((_1395_v9).length))];
        let _rhs249 = _1388_v3;
        let _lhs201 = _1394_v8;
        let _lhs202 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1394_v8).length));
        let _lhs203 = globalState;
        let _lhs204 = globalState;
        _lhs201[_lhs202] = _rhs247;
        _lhs203.f7 = _rhs248;
        _lhs204.f6 = _rhs249;
        let _1399_v13;
        let _nw226 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _1399_v13 = _nw226;
        let _index182 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1399_v13).length));
        (_1399_v13)[_index182] = _1395_v9;
        let _index183 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1399_v13).length));
        (_1399_v13)[_index183] = _1395_v9;
        let _1400_v14;
        _1400_v14 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f11);
        (globalState).f1 = (_module.__default.fm2(new BigNumber(384), (_1394_v8)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_1394_v8).length))], globalState)).plus((_dafny.ZERO).minus((((_1393_v7).contains((_1394_v8)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_1394_v8).length))])) ? ((_1393_v7).get((_1394_v8)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_1394_v8).length))])) : (new BigNumber((_1400_v14).length)))));
        let _1401_v15;
        _1401_v15 = _dafny.Seq.UnicodeFromString("irx");
        _1401_v15 = _dafny.Seq.update(_1401_v15, _module.__default.safeIndex(new BigNumber(382), new BigNumber((_1401_v15).length)), new _dafny.CodePoint('m'.codePointAt(0)));
        let _1402_v16;
        let _nw227 = new _module.C2();
        _nw227.__ctor(p0, ((_this).f11).minus((_this).f11));
        _1402_v16 = _nw227;
      } else {
        let _1403_v17;
        _1403_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1388_v3,_1388_v3);
        let _1404_v18;
        _1404_v18 = _module.D8.create_DC15(_1403_v17);
        let _1405_v19;
        _1405_v19 = _module.D8.create_DC17(_module.D8.create_DC17(_1404_v18));
        let _1406_v20;
        _1406_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1405_v19,(_this).f11);
        let _1407_v21;
        _1407_v21 = _module.D13.create_DC27(_1406_v20);
        let _1408_v22;
        _1408_v22 = _module.D13.create_DC29(_1407_v21);
        let _source21 = _1408_v22;
        if (_source21.is_DC28) {
          let _1409___mcc_h0 = (_source21).cf40;
          let _1410___mcc_h1 = (_source21).cf41;
          let _1411___mcc_h2 = (_source21).cf42;
          let _1412_cf42 = _1411___mcc_h2;
          let _1413_cf41 = _1410___mcc_h1;
          let _1414_cf40 = _1409___mcc_h0;
          let _1415_v23;
          _1415_v23 = _dafny.MultiSet.fromElements(_1412_cf42);
          let _1416_v24;
          _1416_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1415_v23).cardinality()),_1412_cf42);
          let _1417_v25;
          _1417_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1388_v3,_module.D12.create_DC26(_1416_v24));
          let _1418_v26;
          _1418_v26 = _dafny.MultiSet.fromElements(_1417_v25, _1417_v25);
          let _1419_v27;
          _1419_v27 = _dafny.Seq.of(_1413_cf41, new BigNumber(262));
          let _1420_v28;
          _1420_v28 = _dafny.Map.Empty.slice().updateUnsafe(!(_1412_cf42),_1413_cf41);
          let _1421_v29;
          let _nw228 = Array((new BigNumber(20)).toNumber());
          _nw228[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw228[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(_1413_cf41, new BigNumber(328));
          _nw228[(new BigNumber(2)).toNumber()] = (_this).f11;
          _nw228[(new BigNumber(3)).toNumber()] = (_this).f11;
          _nw228[(new BigNumber(4)).toNumber()] = _1413_cf41;
          _nw228[(new BigNumber(5)).toNumber()] = new BigNumber(6);
          _nw228[(new BigNumber(6)).toNumber()] = _1413_cf41;
          _nw228[(new BigNumber(7)).toNumber()] = (_this).f11;
          _nw228[(new BigNumber(8)).toNumber()] = ((_dafny.ZERO).minus((((_1418_v26).contains(_1417_v25)) ? ((_1418_v26).get(_1417_v25)) : ((_this).f11)))).plus((_this).f11);
          _nw228[(new BigNumber(9)).toNumber()] = (_this).f11;
          _nw228[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1419_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(371)), function (_1422_i1) {
            return (_this).f11;
          }))).length);
          _nw228[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt((_this).f11, (((_1420_v28).contains(_1388_v3)) ? ((_1420_v28).get(_1388_v3)) : ((_this).f11)));
          _nw228[(new BigNumber(12)).toNumber()] = _1413_cf41;
          _nw228[(new BigNumber(13)).toNumber()] = _1413_cf41;
          _nw228[(new BigNumber(14)).toNumber()] = (new BigNumber((_1415_v23).cardinality())).minus((_this).fm6(true, (_this).f11, globalState));
          _nw228[(new BigNumber(15)).toNumber()] = _1413_cf41;
          _nw228[(new BigNumber(16)).toNumber()] = (_this).f11;
          _nw228[(new BigNumber(17)).toNumber()] = _1413_cf41;
          _nw228[(new BigNumber(18)).toNumber()] = _1413_cf41;
          _nw228[(new BigNumber(19)).toNumber()] = new BigNumber(798);
          _1421_v29 = _nw228;
          let _index184 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_1421_v29).length));
          (_1421_v29)[_index184] = (_this).f11;
          let _1423_v30;
          _1423_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1414_cf40,new BigNumber((_1415_v23).cardinality()));
          let _1424_v31;
          _1424_v31 = _dafny.Map.Empty.slice().updateUnsafe(!(_1412_cf42),_1414_cf40);
          let _index185 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_1421_v29).length));
          (_1421_v29)[_index185] = (((_1423_v30).contains((((_1424_v31).contains(_1388_v3)) ? ((_1424_v31).get(_1388_v3)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-548)), function (_1426_i2) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }))))) ? ((_1423_v30).get((((_1424_v31).contains(_1388_v3)) ? ((_1424_v31).get(_1388_v3)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-548)), function (_1425_i2) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }))))) : ((_this).fm3((_this).f11, _1412_cf42, _1388_v3, globalState)));
          let _1427_v32;
          _1427_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1421_v29,_1388_v3);
          let _1428_v33;
          _1428_v33 = _module.D10.create_DC20(_1427_v32);
          let _1429_v34;
          _1429_v34 = _dafny.MultiSet.fromElements(_1428_v33, _1428_v33);
          _1429_v34 = _1429_v34;
          let _1430_v35;
          let _nw229 = new _module.C1();
          _nw229.__ctor(false, _1412_cf42);
          _1430_v35 = _nw229;
          let _1431_v36;
          _1431_v36 = _dafny.Seq.of(_1430_v35);
          let _index186 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_1421_v29).length));
          let _rhs250 = (new BigNumber((_1416_v24).length)).multipliedBy(_1413_cf41);
          let _rhs251 = (_dafny.ZERO).minus(((_dafny.ZERO).minus((_dafny.ZERO).minus((_1413_cf41).multipliedBy((_1421_v29)[_module.__default.safeIndex(new BigNumber(221), new BigNumber((_1421_v29).length))])))).minus(_1413_cf41));
          let _rhs252 = (_1430_v35).f19;
          let _rhs253 = _1431_v36;
          let _lhs205 = _1421_v29;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_1421_v29).length));
          let _lhs207 = globalState;
          let _lhs208 = globalState;
          _lhs205[_lhs206] = _rhs250;
          _lhs207.f1 = _rhs251;
          _lhs208.f6 = _rhs252;
          _1431_v36 = _rhs253;
          (globalState).f1 = (_this).f11;
        } else if (_source21.is_DC27) {
          let _1432___mcc_h3 = (_source21).cf39;
          let _1433_cf39 = _1432___mcc_h3;
          let _1434_v37;
          _1434_v37 = _dafny.Seq.of(_1388_v3, _1388_v3, _1388_v3);
          let _1435_v38;
          _1435_v38 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of((_this).f11), _dafny.Seq.of((_this).f11, (_dafny.ZERO).minus((_this).f11), new BigNumber((_1434_v37).length))),(_this).f11);
          _1435_v38 = _1435_v38;
          let _1436_v39;
          _1436_v39 = _dafny.Seq.of((_this).f11);
          let _1437_v40;
          _1437_v40 = _dafny.Seq.UnicodeFromString("bdlkv");
          let _1438_v41;
          _1438_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(new BigNumber(78), (_this).f11, globalState),(_this).f11);
          let _1439_v42;
          _1439_v42 = _dafny.MultiSet.fromElements((_this).f11);
          let _1440_v43;
          _1440_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1439_v42,_1438_v41);
          let _1441_v44;
          let _nw230 = Array((new BigNumber(16)).toNumber());
          _nw230[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw230[(_dafny.ONE).toNumber()] = ((_this).f11).multipliedBy((_this).f11);
          _nw230[(new BigNumber(2)).toNumber()] = (_this).f11;
          _nw230[(new BigNumber(3)).toNumber()] = new BigNumber(-574);
          _nw230[(new BigNumber(4)).toNumber()] = (_this).f11;
          _nw230[(new BigNumber(5)).toNumber()] = (_this).f11;
          _nw230[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_1436_v39).length)).plus((_this).f11));
          _nw230[(new BigNumber(7)).toNumber()] = (_this).f11;
          _nw230[(new BigNumber(8)).toNumber()] = (new BigNumber((_1437_v40).length)).multipliedBy((_this).f11);
          _nw230[(new BigNumber(9)).toNumber()] = new BigNumber((_1403_v17).length);
          _nw230[(new BigNumber(10)).toNumber()] = (_this).f11;
          _nw230[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1438_v41).length), (_this).f11);
          _nw230[(new BigNumber(12)).toNumber()] = (_this).f11;
          _nw230[(new BigNumber(13)).toNumber()] = (_this).f11;
          _nw230[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_1436_v39)[_module.__default.safeIndex((_this).f11, new BigNumber((_1436_v39).length))]);
          _nw230[(new BigNumber(15)).toNumber()] = ((_this).fm3((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f11)), (_this).fm4((_this).f11, _1440_v43, (_this).f11, _1388_v3, globalState), _1388_v3, globalState)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1388_v3)).length));
          _1441_v44 = _nw230;
          let _index187 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_1441_v44).length));
          (_1441_v44)[_index187] = new BigNumber(((_1439_v42).Union(_1439_v42)).cardinality());
          let _index188 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1441_v44).length));
          (_1441_v44)[_index188] = (_this).f11;
          let _1442_v45;
          _1442_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1386_v1,(_this).f11);
          let _index189 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_1441_v44).length));
          let _index190 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1441_v44).length));
          let _rhs254 = _module.__default.safeDivisionInt(new BigNumber((_1434_v37).length), (_this).f11);
          let _rhs255 = ((((_1442_v45).contains(_1386_v1)) ? ((_1442_v45).get(_1386_v1)) : ((_this).f11))).minus((_this).f11);
          let _lhs209 = _1441_v44;
          let _lhs210 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_1441_v44).length));
          let _lhs211 = _1441_v44;
          let _lhs212 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1441_v44).length));
          _lhs209[_lhs210] = _rhs254;
          _lhs211[_lhs212] = _rhs255;
          let _1443_v46;
          _1443_v46 = _module.D10.create_DC21((((_1403_v17).contains(_1388_v3)) ? ((_1403_v17).get(_1388_v3)) : (_1388_v3)), _1388_v3);
          let _1444_v47;
          _1444_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1443_v46,_dafny.Map.Empty.slice().updateUnsafe((_1441_v44)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((_1441_v44).length))],(_this).f11));
          _1444_v47 = (_1444_v47).update(_1443_v46, _1438_v41);
          let _1445_v48;
          let _nw231 = Array((new BigNumber(25)).toNumber()).fill(_module.D12.Default());
          _1445_v48 = _nw231;
          let _1446_v49;
          let _init29 = ((_1447_v3) => function (_1448_i3) {
            return _module.D8.create_DC16(_1447_v3, new BigNumber(-553));
          })(_1388_v3);
          let _nw232 = Array((new BigNumber(10)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw232.length); _i0_29++) {
            _nw232[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1446_v49 = _nw232;
          let _1449_v50;
          _1449_v50 = _module.D12.create_DC25(_1446_v49);
          let _index191 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1445_v48).length));
          (_1445_v48)[_index191] = _1449_v50;
          let _index192 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1445_v48).length));
          (_1445_v48)[_index192] = _module.D12.create_DC25(_1446_v49);
        } else {
          let _1450___mcc_h4 = (_source21).cf43;
          let _1451_cf43 = _1450___mcc_h4;
          let _1452_v51;
          _1452_v51 = _dafny.Seq.UnicodeFromString("x");
          _1452_v51 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("djqq"), _1452_v51);
          let _1453_v52;
          _1453_v52 = new _dafny.CodePoint('k'.codePointAt(0));
          _1453_v52 = p0;
          let _1454_v53;
          _1454_v53 = _module.D15.create_DC36((_this).f11, _1388_v3, _1388_v3, _1388_v3, (_dafny.ZERO).minus((_this).f11));
          (globalState).f1 = (_this).fm6((_1454_v53).dtor_cf58, (_this).f11, globalState);
          _1453_v52 = p0;
        }
        let _1455_v54;
        _1455_v54 = _module.D15.create_DC34();
        let _1456_v55;
        _1456_v55 = _dafny.Set.fromElements(_1455_v54, _1455_v54, _1455_v54);
        let _source22 = _module.__default.fm47(_1456_v55, _1388_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(407)), ((_1457_v3) => function (_1458_i4) {
          return _dafny.Seq.of(!(_1457_v3));
        })(_1388_v3)), _1455_v54, globalState);
        if (_source22.is_DC28) {
          let _1459___mcc_h5 = (_source22).cf40;
          let _1460___mcc_h6 = (_source22).cf41;
          let _1461___mcc_h7 = (_source22).cf42;
          let _1462_cf42 = _1461___mcc_h7;
          let _1463_cf41 = _1460___mcc_h6;
          let _1464_cf40 = _1459___mcc_h5;
          let _1465_v56;
          _1465_v56 = _module.D8.create_DC15(_1403_v17);
          let _1466_v57;
          _1466_v57 = _dafny.MultiSet.fromElements(_1388_v3, _1462_cf42);
          let _1467_v58;
          _1467_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1464_cf40,(((_1466_v57).contains(_1388_v3)) ? ((_1466_v57).get(_1388_v3)) : (_1463_cf41)));
          let _1468_v60;
          _1468_v60 = _dafny.Seq.of(_1464_cf40, _1464_cf40, _1464_cf40, _dafny.Seq.UnicodeFromString("nhsif"), _1464_cf40);
          let _1469_v61;
          let _nw233 = new _module.C4();
          _nw233.__ctor(_dafny.MultiSet.fromElements((_this).f11), _1388_v3, _dafny.Set.fromElements(_1388_v3), (_this).f11);
          _1469_v61 = _nw233;
          let _1470_v62;
          let _nw234 = Array((new BigNumber(17)).toNumber());
          _nw234[(_dafny.ZERO).toNumber()] = _1469_v61;
          _nw234[(_dafny.ONE).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(2)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(3)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(4)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(5)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(6)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(7)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(8)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(9)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(10)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(11)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(12)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(13)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(14)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(15)).toNumber()] = _1469_v61;
          _nw234[(new BigNumber(16)).toNumber()] = _1469_v61;
          _1470_v62 = _nw234;
          let _1471_v63;
          _1471_v63 = _dafny.MultiSet.fromElements(_1470_v62, _1470_v62, _1470_v62);
          let _1472_v64;
          _1472_v64 = _dafny.Seq.of((_this).f11, _1463_cf41);
          let _1473_v65;
          _1473_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1466_v57);
          let _1474_v66;
          let _nw235 = Array((new BigNumber(29)).toNumber());
          _nw235[(_dafny.ZERO).toNumber()] = _1463_cf41;
          _nw235[(_dafny.ONE).toNumber()] = (_this).fm3((_this).f11, _1388_v3, _1388_v3, globalState);
          _nw235[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1463_cf41), (_this).f11);
          _nw235[(new BigNumber(3)).toNumber()] = (_this).f11;
          _nw235[(new BigNumber(4)).toNumber()] = (_1463_cf41).multipliedBy((_this).f11);
          _nw235[(new BigNumber(5)).toNumber()] = (_this).fm6(_1462_cf42, _1463_cf41, globalState);
          _nw235[(new BigNumber(6)).toNumber()] = _1463_cf41;
          _nw235[(new BigNumber(7)).toNumber()] = _1463_cf41;
          _nw235[(new BigNumber(8)).toNumber()] = new BigNumber(((_1467_v58).Merge(function () {
            let _coll64 = new _dafny.Map();
            for (const _compr_64 of (_1468_v60).Elements) {
              let _1475_v59 = _compr_64;
              if (_dafny.Seq.contains(_1468_v60, _1475_v59)) {
                _coll64.push([_1475_v59,new BigNumber(510)]);
              }
            }
            return _coll64;
          }())).length);
          _nw235[(new BigNumber(9)).toNumber()] = (_this).f11;
          _nw235[(new BigNumber(10)).toNumber()] = _1463_cf41;
          _nw235[(new BigNumber(11)).toNumber()] = (_this).f11;
          _nw235[(new BigNumber(12)).toNumber()] = (_this).f11;
          _nw235[(new BigNumber(13)).toNumber()] = (_this).f11;
          _nw235[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_1471_v63).cardinality())).minus((_this).f11));
          _nw235[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-181), (_1469_v61).fm3(new BigNumber((_1464_cf40).length), _1462_cf42, _1462_cf42, globalState));
          _nw235[(new BigNumber(16)).toNumber()] = (_1469_v61).f11;
          _nw235[(new BigNumber(17)).toNumber()] = new BigNumber(438);
          _nw235[(new BigNumber(18)).toNumber()] = (_this).fm3((_dafny.ZERO).minus((_1469_v61).f11), _1388_v3, !(_1388_v3), globalState);
          _nw235[(new BigNumber(19)).toNumber()] = (_1463_cf41).multipliedBy((_this).f11);
          _nw235[(new BigNumber(20)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(_1388_v3,_1462_cf42))).update(_1388_v3, _1403_v17)).length);
          _nw235[(new BigNumber(21)).toNumber()] = new BigNumber((_1403_v17).length);
          _nw235[(new BigNumber(22)).toNumber()] = new BigNumber((_1472_v64).length);
          _nw235[(new BigNumber(23)).toNumber()] = _1463_cf41;
          _nw235[(new BigNumber(24)).toNumber()] = (_this).f11;
          _nw235[(new BigNumber(25)).toNumber()] = new BigNumber(((_this).f12).length);
          _nw235[(new BigNumber(26)).toNumber()] = (_1463_cf41).plus((_this).f11);
          _nw235[(new BigNumber(27)).toNumber()] = new BigNumber((_1473_v65).length);
          _nw235[(new BigNumber(28)).toNumber()] = _1463_cf41;
          _1474_v66 = _nw235;
          let _1476_v67;
          _1476_v67 = _dafny.MultiSet.fromElements((_1463_cf41).plus((_this).f11));
          let _1477_v68;
          _1477_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1463_cf41,_1463_cf41);
          let _1478_v69;
          _1478_v69 = _dafny.Seq.of(!(_1388_v3), _1388_v3);
          let _rhs256 = (((_1476_v67).contains((_dafny.ZERO).minus((((_1477_v68).contains(new BigNumber(-405))) ? ((_1477_v68).get(new BigNumber(-405))) : ((_this).f11))))) ? ((_1476_v67).get((_dafny.ZERO).minus((((_1477_v68).contains(new BigNumber(-405))) ? ((_1477_v68).get(new BigNumber(-405))) : ((_this).f11))))) : ((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(762), new BigNumber((_1478_v69).length)))));
          let _rhs257 = new BigNumber((_dafny.Seq.Concat(_1464_cf40, _dafny.Seq.Create(_module.__default.abs(new BigNumber(443)), ((_1479_p0) => function (_1480_i5) {
            return _1479_p0;
          })(p0)))).length);
          let _rhs258 = _module.__default.fm48(new BigNumber(458), globalState);
          let _rhs259 = (_1462_cf42) === (_1462_cf42);
          let _rhs260 = _1474_v66;
          let _lhs213 = globalState;
          let _lhs214 = globalState;
          let _lhs215 = globalState;
          _lhs213.f1 = _rhs256;
          _lhs214.f1 = _rhs257;
          _1465_v56 = _rhs258;
          _lhs215.f7 = _rhs259;
          _1474_v66 = _rhs260;
          _1388_v3 = true;
          let _1481_v70;
          let _nw236 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1481_v70 = _nw236;
          let _index193 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_1481_v70).length));
          (_1481_v70)[_index193] = p0;
          let _index194 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_1481_v70).length));
          (_1481_v70)[_index194] = p0;
          let _1482_v71;
          _1482_v71 = _module.D6.create_DC10(_1478_v69);
          let _rhs261 = _1462_cf42;
          let _rhs262 = _1482_v71;
          let _lhs216 = globalState;
          _lhs216.f6 = _rhs261;
          _1482_v71 = _rhs262;
        } else if (_source22.is_DC27) {
          let _1483___mcc_h8 = (_source22).cf39;
          let _1484_cf39 = _1483___mcc_h8;
          let _index195 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1386_v1).length));
          (_1386_v1)[_index195] = (_1388_v3) === (_1388_v3);
          let _1485_v72;
          _1485_v72 = _dafny.Seq.UnicodeFromString("ctfp");
          let _1486_v73;
          _1486_v73 = _module.D8.create_DC16(_1388_v3, new BigNumber((_1485_v72).length));
          let _index196 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_1386_v1).length));
          (_1386_v1)[_index196] = (_1486_v73).dtor_cf23;
          let _1487_v74;
          _1487_v74 = _dafny.Seq.of((_this).f11);
          _1487_v74 = _dafny.Seq.of((_this).f11, (_this).f11, (_this).f11);
          let _1488_v75;
          _1488_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1388_v3,new BigNumber(879));
          let _1489_v76;
          _1489_v76 = _dafny.Seq.of(_1488_v75);
          let _1490_v77;
          _1490_v77 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_dafny.ZERO).minus(new BigNumber((_1489_v76).length)), (_this).f11, globalState),_1487_v74);
          _1490_v77 = (_1490_v77).update((((_1386_v1)[_module.__default.safeIndex(new BigNumber(744), new BigNumber((_1386_v1).length))]) ? ((_dafny.ZERO).minus(new BigNumber((_module.__default.fm1(globalState)).length))) : (new BigNumber(649))), _1487_v74);
          let _1491_v78;
          let _nw237 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _1491_v78 = _nw237;
          let _1492_v79;
          _1492_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1491_v78);
          _1492_v79 = (_1492_v79).update((_this).f11, _1491_v78);
        } else {
          let _1493___mcc_h9 = (_source22).cf43;
          let _1494_cf43 = _1493___mcc_h9;
          let _1495_v80;
          _1495_v80 = _dafny.Seq.of(_1388_v3);
          let _1496_v81;
          _1496_v81 = _dafny.Set.fromElements(p0, p0, p0, _module.__default.fm20(_1495_v80, (_this).f11, _1388_v3, globalState));
          _1496_v81 = _1496_v81;
          let _1497_v82;
          _1497_v82 = _dafny.Seq.UnicodeFromString("l");
          let _1498_v83;
          _1498_v83 = _dafny.Seq.of((_this).f11);
          let _1499_v84;
          let _nw238 = Array((new BigNumber(3)).toNumber());
          _nw238[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw238[(_dafny.ONE).toNumber()] = (_this).f11;
          _nw238[(new BigNumber(2)).toNumber()] = (_this).f11;
          _1499_v84 = _nw238;
          let _1500_v85;
          _1500_v85 = _dafny.MultiSet.fromElements(_1388_v3, _1388_v3);
          let _1501_v86;
          _1501_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_1499_v84, _1499_v84)).length),new BigNumber((_1500_v85).cardinality()));
          let _1502_v87;
          let _nw239 = Array((new BigNumber(15)).toNumber());
          _nw239[(_dafny.ZERO).toNumber()] = ((_this).f11).multipliedBy(new BigNumber((_1497_v82).length));
          _nw239[(_dafny.ONE).toNumber()] = (_this).f11;
          _nw239[(new BigNumber(2)).toNumber()] = (_1498_v83)[_module.__default.safeIndex((((_1501_v86).contains((_module.D11.create_DC23(_module.__default.fm0(globalState), new BigNumber(19), (_this).f11)).dtor_cf33)) ? ((_1501_v86).get((_module.D11.create_DC23(_module.__default.fm0(globalState), new BigNumber(19), (_this).f11)).dtor_cf33)) : (new BigNumber((_dafny.Seq.UnicodeFromString("hwt")).length))), new BigNumber((_1498_v83).length))];
          _nw239[(new BigNumber(3)).toNumber()] = (_this).f11;
          _nw239[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt((_this).f11, (_this).f11);
          _nw239[(new BigNumber(5)).toNumber()] = (_this).f11;
          _nw239[(new BigNumber(6)).toNumber()] = (_this).f11;
          _nw239[(new BigNumber(7)).toNumber()] = (_this).fm3((_this).f11, !(_1388_v3), _1388_v3, globalState);
          _nw239[(new BigNumber(8)).toNumber()] = (_this).f11;
          _nw239[(new BigNumber(9)).toNumber()] = (_this).f11;
          _nw239[(new BigNumber(10)).toNumber()] = (_this).f11;
          _nw239[(new BigNumber(11)).toNumber()] = ((_this).f11).plus(new BigNumber((_1495_v80).length));
          _nw239[(new BigNumber(12)).toNumber()] = (_this).f11;
          _nw239[(new BigNumber(13)).toNumber()] = (_this).f11;
          _nw239[(new BigNumber(14)).toNumber()] = (_this).f11;
          _1502_v87 = _nw239;
          let _index197 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_1499_v84).length));
          (_1499_v84)[_index197] = new BigNumber(835);
          let _1503_v88;
          _1503_v88 = _module.D14.create_DC31(_1388_v3, (_this).f11, (_this).f11);
          let _1504_v89;
          _1504_v89 = _dafny.MultiSet.fromElements(_1497_v82);
          let _index198 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_1499_v84).length));
          (_1499_v84)[_index198] = _module.__default.safeModuloInt((_1503_v88).dtor_cf46, _module.__default.safeModuloInt((_this).f11, new BigNumber((_1504_v89).cardinality())));
          let _1505_v90;
          let _nw240 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1505_v90 = _nw240;
          let _1506_v91;
          _1506_v91 = _1505_v90;
          _1505_v90 = (_1506_v91);
          (globalState).f1 = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_1495_v80, _dafny.Seq.Concat(_1495_v80, _1495_v80)), _module.__default.safeIndex(new BigNumber(-764), new BigNumber((_dafny.Seq.Concat(_1495_v80, _dafny.Seq.Concat(_1495_v80, _1495_v80))).length)), _1388_v3), _module.__default.safeIndex(_module.__default.safeModuloInt((_1499_v84)[_module.__default.safeIndex(new BigNumber(585), new BigNumber((_1499_v84).length))], (_1499_v84)[_module.__default.safeIndex(new BigNumber(585), new BigNumber((_1499_v84).length))]), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1495_v80, _dafny.Seq.Concat(_1495_v80, _1495_v80)), _module.__default.safeIndex(new BigNumber(-764), new BigNumber((_dafny.Seq.Concat(_1495_v80, _dafny.Seq.Concat(_1495_v80, _1495_v80))).length)), _1388_v3)).length)), ((_this).f11).isEqualTo((_this).f11)), _module.__default.safeIndex(new BigNumber(-761), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_1495_v80, _dafny.Seq.Concat(_1495_v80, _1495_v80)), _module.__default.safeIndex(new BigNumber(-764), new BigNumber((_dafny.Seq.Concat(_1495_v80, _dafny.Seq.Concat(_1495_v80, _1495_v80))).length)), _1388_v3), _module.__default.safeIndex(_module.__default.safeModuloInt((_1499_v84)[_module.__default.safeIndex(new BigNumber(585), new BigNumber((_1499_v84).length))], (_1499_v84)[_module.__default.safeIndex(new BigNumber(585), new BigNumber((_1499_v84).length))]), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1495_v80, _dafny.Seq.Concat(_1495_v80, _1495_v80)), _module.__default.safeIndex(new BigNumber(-764), new BigNumber((_dafny.Seq.Concat(_1495_v80, _dafny.Seq.Concat(_1495_v80, _1495_v80))).length)), _1388_v3)).length)), ((_this).f11).isEqualTo((_this).f11))).length)), _1388_v3))).cardinality());
        }
        let _1507_v92;
        let _nw241 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _1507_v92 = _nw241;
        let _index199 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1507_v92).length));
        (_1507_v92)[_index199] = (_this).f11;
        let _index200 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_1507_v92).length));
        (_1507_v92)[_index200] = (_this).fm3(((_this).f11).minus((_this).f11), _1388_v3, !(true), globalState);
        (globalState).f7 = ((_1507_v92)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_1507_v92).length))]).isEqualTo((_this).f11);
        let _1508_v93;
        _1508_v93 = _dafny.Seq.of(_1507_v92);
        let _1509_v94;
        _1509_v94 = _dafny.MultiSet.fromElements(_1388_v3);
        let _1510_v95;
        _1510_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(51),_1388_v3);
        let _1511_v96;
        let _nw242 = Array((new BigNumber(13)).toNumber());
        _nw242[(_dafny.ZERO).toNumber()] = _1507_v92;
        _nw242[(_dafny.ONE).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(2)).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(3)).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(4)).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(5)).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(6)).toNumber()] = (_1508_v93)[_module.__default.safeIndex((((_1509_v94).contains(_1388_v3)) ? ((_1509_v94).get(_1388_v3)) : (new BigNumber((_1510_v95).length))), new BigNumber((_1508_v93).length))];
        _nw242[(new BigNumber(7)).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(8)).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(9)).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(10)).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(11)).toNumber()] = _1507_v92;
        _nw242[(new BigNumber(12)).toNumber()] = _1507_v92;
        _1511_v96 = _nw242;
        let _1512_v97;
        let _nw243 = Array((new BigNumber(8)).toNumber());
        _nw243[(_dafny.ZERO).toNumber()] = _1511_v96;
        _nw243[(_dafny.ONE).toNumber()] = _1511_v96;
        _nw243[(new BigNumber(2)).toNumber()] = _1511_v96;
        _nw243[(new BigNumber(3)).toNumber()] = _1511_v96;
        _nw243[(new BigNumber(4)).toNumber()] = (_1511_v96);
        _nw243[(new BigNumber(5)).toNumber()] = _1511_v96;
        _nw243[(new BigNumber(6)).toNumber()] = _1511_v96;
        _nw243[(new BigNumber(7)).toNumber()] = _1511_v96;
        _1512_v97 = _nw243;
        let _1513_v98;
        _1513_v98 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1511_v96);
        let _index201 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_1512_v97).length));
        (_1512_v97)[_index201] = (((_1513_v98).contains((_1507_v92)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_1507_v92).length))])) ? ((_1513_v98).get((_1507_v92)[_module.__default.safeIndex(new BigNumber(440), new BigNumber((_1507_v92).length))])) : (_1511_v96));
        let _index202 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_1512_v97).length));
        let _nw244 = Array((new BigNumber(29)).toNumber()).fill([]);
        (_1512_v97)[_index202] = _nw244;
      }
      let _1514_v99;
      _1514_v99 = _module.D15.create_DC36((_this).f11, _1388_v3, _1388_v3, _1388_v3, (_this).f11);
      (globalState).f1 = (_dafny.ZERO).minus((_1514_v99).dtor_cf55);
      let _1515_v100;
      _1515_v100 = _dafny.Seq.UnicodeFromString("dohvyph");
      let _hi1 = new BigNumber((_1515_v100).length);
      for (let _1516_i6 = ((_this).f11).multipliedBy((_this).f11); _1516_i6.isLessThan(_hi1); _1516_i6 = _1516_i6.plus(_dafny.ONE)) {
        let _1517_v101;
        let _nw245 = Array((new BigNumber(17)).toNumber()).fill([]);
        _1517_v101 = _nw245;
        let _1518_v102;
        _1518_v102 = _1517_v101;
        let _1519_v103;
        let _nw246 = Array((new BigNumber(8)).toNumber());
        _nw246[(_dafny.ZERO).toNumber()] = _1518_v102;
        _nw246[(_dafny.ONE).toNumber()] = _1518_v102;
        _nw246[(new BigNumber(2)).toNumber()] = _1517_v101;
        _nw246[(new BigNumber(3)).toNumber()] = _1518_v102;
        _nw246[(new BigNumber(4)).toNumber()] = _1518_v102;
        _nw246[(new BigNumber(5)).toNumber()] = _1517_v101;
        _nw246[(new BigNumber(6)).toNumber()] = _1518_v102;
        _nw246[(new BigNumber(7)).toNumber()] = _1518_v102;
        _1519_v103 = _nw246;
        let _1520_v104;
        _1520_v104 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f11), _1516_i6);
        let _1521_v105;
        _1521_v105 = _module.D0.create_DC0(new BigNumber((_1520_v104).length));
        let _1522_v106;
        _1522_v106 = _dafny.Seq.of(!(_1388_v3));
        let _1523_v107;
        _1523_v107 = _dafny.Seq.of((_this).fm5(function (_pat_let31_0) {
          return function (_1524_dt__update__tmp_h0) {
            return function (_pat_let32_0) {
              return function (_1525_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_1525_dt__update_hcf0_h0);
              }(_pat_let32_0);
            }((_this).f11);
          }(_pat_let31_0);
        }(_1521_v105), _1522_v106, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(481)), ((_1526_p0) => function (_1527_i7) {
          return _1526_p0;
        })(p0))).length), p0, globalState));
        let _1528_v108;
        _1528_v108 = _dafny.MultiSet.fromElements(new BigNumber((_1522_v106).length));
        let _1529_v109;
        _1529_v109 = _dafny.Set.fromElements((_this).f11, _1516_i6);
        let _1530_v110;
        _1530_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1388_v3,new BigNumber(716));
        let _1531_v111;
        _1531_v111 = _dafny.Seq.of(_1520_v104, _1520_v104);
        let _1532_v112;
        let _nw247 = Array((new BigNumber(21)).toNumber());
        _nw247[(_dafny.ZERO).toNumber()] = new BigNumber(840);
        _nw247[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1516_i6)).length);
        _nw247[(new BigNumber(2)).toNumber()] = _1516_i6;
        _nw247[(new BigNumber(3)).toNumber()] = (_this).f11;
        _nw247[(new BigNumber(4)).toNumber()] = new BigNumber((_1528_v108).cardinality());
        _nw247[(new BigNumber(5)).toNumber()] = new BigNumber((_1529_v109).length);
        _nw247[(new BigNumber(6)).toNumber()] = (_this).f11;
        _nw247[(new BigNumber(7)).toNumber()] = (_this).f11;
        _nw247[(new BigNumber(8)).toNumber()] = _1516_i6;
        _nw247[(new BigNumber(9)).toNumber()] = (((_1530_v110).contains(_1388_v3)) ? ((_1530_v110).get(_1388_v3)) : (_1516_i6));
        _nw247[(new BigNumber(10)).toNumber()] = _1516_i6;
        _nw247[(new BigNumber(11)).toNumber()] = (_this).f11;
        _nw247[(new BigNumber(12)).toNumber()] = _module.__default.fm2((_this).f11, _1516_i6, globalState);
        _nw247[(new BigNumber(13)).toNumber()] = new BigNumber((_1523_v107).length);
        _nw247[(new BigNumber(14)).toNumber()] = new BigNumber((_1531_v111).length);
        _nw247[(new BigNumber(15)).toNumber()] = _1516_i6;
        _nw247[(new BigNumber(16)).toNumber()] = _1516_i6;
        _nw247[(new BigNumber(17)).toNumber()] = _1516_i6;
        _nw247[(new BigNumber(18)).toNumber()] = _1516_i6;
        _nw247[(new BigNumber(19)).toNumber()] = _1516_i6;
        _nw247[(new BigNumber(20)).toNumber()] = _1516_i6;
        _1532_v112 = _nw247;
        let _1533_v113;
        _1533_v113 = _dafny.Map.Empty.slice().updateUnsafe(_1519_v103,_1532_v112);
        let _1534_v114;
        _1534_v114 = _module.D11.create_DC22((_1533_v113).Merge(_1533_v113));
        let _source23 = _1534_v114;
        if (_source23.is_DC23) {
          let _1535___mcc_h10 = (_source23).cf31;
          let _1536___mcc_h11 = (_source23).cf32;
          let _1537___mcc_h12 = (_source23).cf33;
          let _1538_cf33 = _1537___mcc_h12;
          let _1539_cf32 = _1536___mcc_h11;
          let _1540_cf31 = _1535___mcc_h10;
          let _1541_v115;
          _1541_v115 = _dafny.MultiSet.fromElements(_1386_v1, _1386_v1);
          let _1542_v116;
          _1542_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1539_cf32,p0);
          let _rhs263 = (new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1515_v100, _1515_v100), _module.__default.safeIndex(_1538_cf33, new BigNumber((_dafny.Seq.Concat(_1515_v100, _1515_v100)).length)), (((_1542_v116).contains(_1538_cf33)) ? ((_1542_v116).get(_1538_cf33)) : (p0)))).length)).minus(_1539_cf32);
          let _rhs264 = _1541_v115;
          let _lhs217 = globalState;
          _lhs217.f1 = _rhs263;
          _1541_v115 = _rhs264;
          _1521_v105 = _1521_v105;
          _1528_v108 = (_1528_v108).Union(_1528_v108);
          let _1543_v117;
          let _nw248 = new _module.C2();
          _nw248.__ctor(((_1540_cf31) ? (p0) : (new _dafny.CodePoint('n'.codePointAt(0)))), (_1539_cf32).multipliedBy(_1539_cf32));
          _1543_v117 = _nw248;
        } else if (_source23.is_DC24) {
          let _1544___mcc_h13 = (_source23).cf34;
          let _1545___mcc_h14 = (_source23).cf35;
          let _1546___mcc_h15 = (_source23).cf36;
          let _1547_cf36 = _1546___mcc_h15;
          let _1548_cf35 = _1545___mcc_h14;
          let _1549_cf34 = _1544___mcc_h13;
          let _1550_v118;
          _1550_v118 = _dafny.Set.fromElements(p0);
          _1550_v118 = _1550_v118;
          let _init30 = function (_1551_i8) {
            return (_1551_i8).minus(new BigNumber((_dafny.Seq.UnicodeFromString("v")).length));
          };
          let _nw249 = Array((new BigNumber(9)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw249.length); _i0_30++) {
            _nw249[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1532_v112 = _nw249;
          let _index203 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1532_v112).length));
          (_1532_v112)[_index203] = _1516_i6;
          let _index204 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1532_v112).length));
          (_1532_v112)[_index204] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1516_i6));
          _1549_cf34 = new BigNumber(700);
        } else {
          let _1552___mcc_h16 = (_source23).cf30;
          let _1553_cf30 = _1552___mcc_h16;
          (globalState).f1 = (_module.__default.fm2((_this).f11, _1516_i6, globalState)).minus(_1516_i6);
          let _1554_v119;
          _1554_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1516_i6,_1516_i6);
          let _rhs265 = _dafny.Seq.Concat(_1520_v104, _dafny.Seq.of((_this).f11, _1516_i6, new BigNumber((_1515_v100).length)));
          let _rhs266 = (_1516_i6).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_1516_i6, _1516_i6));
          let _rhs267 = _1532_v112;
          let _rhs268 = _module.__default.safeDivisionInt(new BigNumber((_1554_v119).length), (_this).f11);
          let _rhs269 = _1386_v1;
          let _lhs218 = globalState;
          let _lhs219 = globalState;
          _1520_v104 = _rhs265;
          _lhs218.f6 = _rhs266;
          _1532_v112 = _rhs267;
          _lhs219.f1 = _rhs268;
          _1386_v1 = _rhs269;
          let _1555_v120;
          _1555_v120 = _module.D11.create_DC24((_this).f11, _1523_v107, p0);
          r0 = (_dafny.MultiSet.FromArray((_1555_v120).dtor_cf35)).Union(_dafny.MultiSet.fromElements(false, true, _1388_v3, _1388_v3));
          let _1556_v122;
          _1556_v122 = _dafny.Seq.of(function () {
            let _coll65 = new _dafny.Map();
            for (const _compr_65 of (_1520_v104).Elements) {
              let _1557_v121 = _compr_65;
              if (_dafny.Seq.contains(_1520_v104, _1557_v121)) {
                _coll65.push([(_1557_v121).minus(_1516_i6),_1388_v3]);
              }
            }
            return _coll65;
          }());
          let _1558_v123;
          _1558_v123 = _module.D0.create_DC1(_1388_v3, new BigNumber(((_1556_v122)[_module.__default.safeIndex(_1516_i6, new BigNumber((_1556_v122).length))]).length), _1388_v3);
          let _1559_v124;
          let _nw250 = Array((new BigNumber(10)).toNumber());
          _nw250[(_dafny.ZERO).toNumber()] = _1558_v123;
          _nw250[(_dafny.ONE).toNumber()] = _module.__default.fm49(_1388_v3, globalState);
          _nw250[(new BigNumber(2)).toNumber()] = _1558_v123;
          _nw250[(new BigNumber(3)).toNumber()] = _1558_v123;
          _nw250[(new BigNumber(4)).toNumber()] = _1558_v123;
          _nw250[(new BigNumber(5)).toNumber()] = _1558_v123;
          _nw250[(new BigNumber(6)).toNumber()] = _1558_v123;
          _nw250[(new BigNumber(7)).toNumber()] = _1558_v123;
          _nw250[(new BigNumber(8)).toNumber()] = _1558_v123;
          _nw250[(new BigNumber(9)).toNumber()] = _1558_v123;
          _1559_v124 = _nw250;
          let _1560_v125;
          _1560_v125 = _dafny.MultiSet.fromElements(_1559_v124);
          let _1561_v126;
          _1561_v126 = _dafny.Seq.of(_1559_v124, _1559_v124);
          let _1562_v127;
          _1562_v127 = _dafny.Seq.of(_1559_v124, _1559_v124, (_1561_v126)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-797)), ((_1563_p0) => function (_1564_i9) {
            return _1563_p0;
          })(p0))).length), new BigNumber((_1561_v126).length))], _1559_v124);
          _1560_v125 = (_dafny.MultiSet.fromElements(_1559_v124, _1559_v124, _1559_v124, (_1562_v127)[_module.__default.safeIndex((_this).f11, new BigNumber((_1562_v127).length))], _1559_v124)).Difference(_1560_v125);
        }
        (globalState).f1 = _1516_i6;
        let _1565_v128;
        _1565_v128 = _dafny.Map.Empty.slice().updateUnsafe(_1388_v3,p0);
        let _1566_v129;
        _1566_v129 = _module.D11.create_DC23(_1388_v3, _1516_i6, (_this).f11);
        let _1567_v130;
        let _nw251 = Array((new BigNumber(17)).toNumber());
        _nw251[(_dafny.ZERO).toNumber()] = p0;
        _nw251[(_dafny.ONE).toNumber()] = (_1515_v100)[_module.__default.safeIndex(_1516_i6, new BigNumber((_1515_v100).length))];
        _nw251[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
        _nw251[(new BigNumber(3)).toNumber()] = p0;
        _nw251[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
        _nw251[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
        _nw251[(new BigNumber(6)).toNumber()] = (((_1565_v128).contains((_1566_v129).dtor_cf31)) ? ((_1565_v128).get((_1566_v129).dtor_cf31)) : (p0));
        _nw251[(new BigNumber(7)).toNumber()] = p0;
        _nw251[(new BigNumber(8)).toNumber()] = p0;
        _nw251[(new BigNumber(9)).toNumber()] = p0;
        _nw251[(new BigNumber(10)).toNumber()] = p0;
        _nw251[(new BigNumber(11)).toNumber()] = p0;
        _nw251[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
        _nw251[(new BigNumber(13)).toNumber()] = p0;
        _nw251[(new BigNumber(14)).toNumber()] = ((_1388_v3) ? (p0) : (p0));
        _nw251[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
        _nw251[(new BigNumber(16)).toNumber()] = p0;
        _1567_v130 = _nw251;
        let _index205 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_1567_v130).length));
        (_1567_v130)[_index205] = p0;
        let _index206 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_1567_v130).length));
        (_1567_v130)[_index206] = p0;
        _1515_v100 = _1515_v100;
      }
      let _1568_v131;
      _1568_v131 = _dafny.MultiSet.fromElements(_1388_v3, _1388_v3, true);
      r0 = _1568_v131;
      return r0;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.Map.Empty;
      (globalState).f1 = _module.__default.safeDivisionInt((_this).f11, (_dafny.ZERO).minus(p0));
      let _1569_v0;
      let _nw252 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1569_v0 = _nw252;
      let _1570_v1;
      _1570_v1 = _dafny.Seq.UnicodeFromString("sva");
      let _index207 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length));
      (_1569_v0)[_index207] = _1570_v1;
      let _index208 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length));
      (_1569_v0)[_index208] = _dafny.Seq.Concat(_module.__default.fm1(globalState), _dafny.Seq.Concat(_1570_v1, _1570_v1));
      let _1571_v2;
      _1571_v2 = true;
      if (!(_1571_v2)) {
        let _1572_v3;
        _1572_v3 = _dafny.Seq.of(_1571_v2);
        let _1573_v4;
        _1573_v4 = _dafny.MultiSet.fromElements(_1571_v2);
        let _source24 = _module.__default.fm50(new BigNumber((_1572_v3).length), _1571_v2, _1571_v2, (_1572_v3)[_module.__default.safeIndex((((_1573_v4).contains(_1571_v2)) ? ((_1573_v4).get(_1571_v2)) : (p0)), new BigNumber((_1572_v3).length))], globalState);
        if (_source24.is_DC23) {
          let _1574___mcc_h0 = (_source24).cf31;
          let _1575___mcc_h1 = (_source24).cf32;
          let _1576___mcc_h2 = (_source24).cf33;
          let _1577_cf33 = _1576___mcc_h2;
          let _1578_cf32 = _1575___mcc_h1;
          let _1579_cf31 = _1574___mcc_h0;
          let _1580_v5;
          let _init31 = ((_1581_v2, _1582_cf31) => function (_1583_i0) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_1581_v2,_1581_v2)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1582_cf31,false));
          })(_1571_v2, _1579_cf31);
          let _nw253 = Array((new BigNumber(13)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw253.length); _i0_31++) {
            _nw253[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1580_v5 = _nw253;
          _1580_v5 = _1580_v5;
          (globalState).f7 = _1579_cf31;
          let _1584_v6;
          _1584_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,_1579_cf31);
          let _1585_v7;
          _1585_v7 = _dafny.Seq.of(_1584_v6);
          let _1586_v8;
          _1586_v8 = _module.D8.create_DC15((_1585_v7)[_module.__default.safeIndex(new BigNumber(178), new BigNumber((_1585_v7).length))]);
          _1586_v8 = _1586_v8;
          let _1587_v9;
          let _init32 = ((_1588_v3, _1589_v1, _1590_cf31) => function (_1591_i1) {
            return _dafny.Seq.update(_1588_v3, _module.__default.safeIndex(new BigNumber((_1589_v1).length), new BigNumber((_1588_v3).length)), !(_1590_cf31));
          })(_1572_v3, _1570_v1, _1579_cf31);
          let _nw254 = Array((new BigNumber(4)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw254.length); _i0_32++) {
            _nw254[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1587_v9 = _nw254;
          let _index209 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_1587_v9).length));
          (_1587_v9)[_index209] = _1572_v3;
          let _index210 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_1587_v9).length));
          let _rhs270 = _module.__default.fm19(_1579_cf31, !(_1579_cf31), globalState);
          let _rhs271 = _1571_v2;
          let _lhs220 = _1587_v9;
          let _lhs221 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_1587_v9).length));
          _lhs220[_lhs221] = _rhs270;
          _1571_v2 = _rhs271;
        } else if (_source24.is_DC24) {
          let _1592___mcc_h3 = (_source24).cf34;
          let _1593___mcc_h4 = (_source24).cf35;
          let _1594___mcc_h5 = (_source24).cf36;
          let _1595_cf36 = _1594___mcc_h5;
          let _1596_cf35 = _1593___mcc_h4;
          let _1597_cf34 = _1592___mcc_h3;
          let _1598_v10;
          let _nw255 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _1598_v10 = _nw255;
          let _index211 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_1598_v10).length));
          (_1598_v10)[_index211] = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f11), p0);
          let _index212 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_1598_v10).length));
          (_1598_v10)[_index212] = (_this).f11;
          _1598_v10 = _1598_v10;
          let _1599_v11;
          _1599_v11 = _dafny.Set.fromElements((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))], _1570_v1);
          _1599_v11 = _1599_v11;
          let _1600_v12;
          _1600_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_dafny.Seq.Concat(_1570_v1, _dafny.Seq.update(_1570_v1, _module.__default.safeIndex(p0, new BigNumber((_1570_v1).length)), _1595_cf36)));
          _1600_v12 = (_1600_v12).update((_this).f11, (_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))]);
        } else {
          let _1601___mcc_h6 = (_source24).cf30;
          let _1602_cf30 = _1601___mcc_h6;
          let _1603_v13;
          let _nw256 = Array((new BigNumber(6)).toNumber());
          _nw256[(_dafny.ZERO).toNumber()] = p0;
          _nw256[(_dafny.ONE).toNumber()] = (p0).multipliedBy(p0);
          _nw256[(new BigNumber(2)).toNumber()] = p0;
          _nw256[(new BigNumber(3)).toNumber()] = _module.__default.fm2((_this).f11, new BigNumber((_dafny.MultiSet.FromArray(_1572_v3)).cardinality()), globalState);
          _nw256[(new BigNumber(4)).toNumber()] = new BigNumber(((_this).f12).length);
          _nw256[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt((_this).f11, (_dafny.ZERO).minus((_this).f11));
          _1603_v13 = _nw256;
          let _index213 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1603_v13).length));
          (_1603_v13)[_index213] = (_this).f11;
          let _index214 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1603_v13).length));
          (_1603_v13)[_index214] = (_this).f11;
          let _1604_v14;
          let _nw257 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
          _1604_v14 = _nw257;
          let _1605_v15;
          _1605_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1604_v14,new BigNumber(((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))]).length));
          let _1606_v16;
          _1606_v16 = _dafny.MultiSet.fromElements((_this).f11, (_this).f11, (_this).f11, (_dafny.ZERO).minus(p0));
          let _index215 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1603_v13).length));
          let _index216 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1603_v13).length));
          let _rhs272 = (((_1605_v15).contains(_1604_v14)) ? ((_1605_v15).get(_1604_v14)) : (new BigNumber((_dafny.Seq.Concat(_1572_v3, _dafny.Seq.of(_1571_v2))).length)));
          let _rhs273 = false;
          let _rhs274 = (_dafny.ZERO).minus(p0);
          let _rhs275 = (_1606_v16).IsSubsetOf((_1606_v16).update((_this).f11, _module.__default.abs((_this).f11)));
          let _lhs222 = _1603_v13;
          let _lhs223 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1603_v13).length));
          let _lhs224 = globalState;
          let _lhs225 = _1603_v13;
          let _lhs226 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1603_v13).length));
          let _lhs227 = globalState;
          _lhs222[_lhs223] = _rhs272;
          _lhs224.f7 = _rhs273;
          _lhs225[_lhs226] = _rhs274;
          _lhs227.f7 = _rhs275;
          let _1607_v17;
          _1607_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,p0);
          let _1608_v18;
          _1608_v18 = _dafny.Set.fromElements(_1607_v17);
          _1608_v18 = _dafny.Set.fromElements(_1607_v17, _1607_v17);
          let _1609_v19;
          _1609_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_1572_v3, _1572_v3)).length),_1572_v3);
          let _1610_v20;
          let _nw258 = new _module.C1();
          _nw258.__ctor(_1571_v2, _1571_v2);
          _1610_v20 = _nw258;
          let _1611_v21;
          _1611_v21 = _1609_v19;
          let _1612_v22;
          _1612_v22 = _module.D13.create_DC28(_dafny.Seq.UnicodeFromString("bqjjbmfo"), (_1603_v13)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_1603_v13).length))], (_1610_v20).f18);
          let _1613_v23;
          _1613_v23 = _dafny.Map.Empty.slice().updateUnsafe((((_1610_v20).f18) ? (_1571_v2) : ((_1610_v20).f18)),_1610_v20);
          let _rhs276 = ((_1571_v2) ? ((_1611_v21)) : ((_1609_v19).update((_1612_v22).dtor_cf41, _1572_v3)));
          let _rhs277 = (((_1613_v23).contains((_1610_v20).f19)) ? ((_1613_v23).get((_1610_v20).f19)) : (_1610_v20));
          _1609_v19 = _rhs276;
          _1610_v20 = _rhs277;
          let _1614_v24;
          _1614_v24 = _dafny.Seq.of((_1603_v13)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_1603_v13).length))], new BigNumber(-974));
          let _1615_v25;
          _1615_v25 = _dafny.Map.Empty.slice().updateUnsafe((_1603_v13)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_1603_v13).length))],(new BigNumber((_1614_v24).length)).multipliedBy(p0));
          _1615_v25 = (_1615_v25).update(new BigNumber(799), (_1603_v13)[_module.__default.safeIndex(new BigNumber(334), new BigNumber((_1603_v13).length))]);
        }
        let _1616_v26;
        _1616_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))]).length))).length),p0);
        let _1617_v27;
        _1617_v27 = _module.D16.create_DC38(_1616_v26);
        let _1618_v28;
        _1618_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1617_v27,_1571_v2);
        let _pat_let_tv43 = _1616_v26;
        let _1619_v29;
        _1619_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,_dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_1618_v28).update(function (_pat_let33_0) {
          return function (_1620_dt__update__tmp_h0) {
            return function (_pat_let34_0) {
              return function (_1621_dt__update_hcf61_h0) {
                return _module.D16.create_DC38(_1621_dt__update_hcf61_h0);
              }(_pat_let34_0);
            }(_pat_let_tv43);
          }(_pat_let33_0);
        }(_module.D16.create_DC38(_module.__default.fm46(globalState))), false)));
        let _1622_v30;
        _1622_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_1618_v28);
        _1619_v29 = (_1619_v29).update(_1571_v2, _1622_v30);
        let _1623_v31;
        let _nw259 = Array((new BigNumber(13)).toNumber());
        _1623_v31 = _nw259;
        let _1624_v32;
        _1624_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1623_v31,(_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f11)));
        let _1625_v33;
        let _nw260 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _1625_v33 = _nw260;
        let _index217 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1625_v33).length));
        (_1625_v33)[_index217] = (_this).f11;
        let _index218 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1625_v33).length));
        let _rhs278 = _dafny.Map.Empty.slice().updateUnsafe(_1623_v31,(p0).plus(p0));
        let _rhs279 = !(_1571_v2);
        let _rhs280 = _module.__default.safeDivisionInt((_this).f11, (_this).f11);
        let _rhs281 = ((_1571_v2) ? ((_dafny.ZERO).minus(new BigNumber((_1572_v3).length))) : (new BigNumber(183)));
        let _lhs228 = globalState;
        let _lhs229 = globalState;
        let _lhs230 = _1625_v33;
        let _lhs231 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1625_v33).length));
        _1624_v32 = _rhs278;
        _lhs228.f7 = _rhs279;
        _lhs229.f1 = _rhs280;
        _lhs230[_lhs231] = _rhs281;
        let _1626_v34;
        _1626_v34 = _dafny.Seq.of((_this).f11);
        (globalState).f1 = (_this).fm3(new BigNumber((_1626_v34).length), _1571_v2, true, globalState);
        (globalState).f7 = _1571_v2;
      } else {
        let _1627_v35;
        let _nw261 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _1627_v35 = _nw261;
        let _1628_v36;
        _1628_v36 = _dafny.Seq.of((_this).f11, (_this).f11);
        let _1629_v37;
        _1629_v37 = _module.D14.create_DC31(_1571_v2, (_this).f11, (_this).f11);
        let _1630_v38;
        _1630_v38 = _module.D14.create_DC32(_1629_v37);
        let _1631_v39;
        _1631_v39 = _dafny.Seq.of(_1630_v38);
        let _index219 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1627_v35).length));
        (_1627_v35)[_index219] = new BigNumber((_dafny.Seq.Concat(_module.__default.fm51(_1628_v36, p0, p0, p0, globalState), _1631_v39)).length);
        let _index220 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1627_v35).length));
        (_1627_v35)[_index220] = (_dafny.ZERO).minus((_this).f11);
        if (_1571_v2) {
          let _1632_v40;
          let _nw262 = Array((_dafny.ONE).toNumber()).fill(false);
          _1632_v40 = _nw262;
          let _index221 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_1632_v40).length));
          (_1632_v40)[_index221] = _1571_v2;
          let _index222 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_1632_v40).length));
          (_1632_v40)[_index222] = _1571_v2;
          _1628_v36 = _1628_v36;
          (globalState).f1 = (_dafny.ZERO).minus((_this).f11);
          let _1633_v41;
          let _nw263 = Array((new BigNumber(24)).toNumber()).fill([]);
          _1633_v41 = _nw263;
          let _index223 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_1633_v41).length));
          (_1633_v41)[_index223] = _1632_v40;
          let _index224 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_1633_v41).length));
          (_1633_v41)[_index224] = _1632_v40;
          let _1634_v42;
          _1634_v42 = _dafny.Seq.of((_1632_v40)[_module.__default.safeIndex(new BigNumber(225), new BigNumber((_1632_v40).length))]);
          let _1635_v43;
          let _nw264 = new _module.C0();
          _nw264.__ctor(_1634_v42, new BigNumber(225));
          _1635_v43 = _nw264;
        } else {
          _1627_v35 = _1627_v35;
          (globalState).f1 = (_this).f11;
          let _1636_v44;
          _1636_v44 = _module.D19.create_DC42(_1627_v35);
          let _1637_v45;
          let _nw265 = Array((new BigNumber(20)).toNumber());
          _nw265[(_dafny.ZERO).toNumber()] = _1627_v35;
          _nw265[(_dafny.ONE).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(2)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(3)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(4)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(5)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(6)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(7)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(8)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(9)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(10)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(11)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(12)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(13)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(14)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(15)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(16)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(17)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(18)).toNumber()] = _1627_v35;
          _nw265[(new BigNumber(19)).toNumber()] = (_1636_v44).dtor_cf68;
          _1637_v45 = _nw265;
          let _index225 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1637_v45).length));
          (_1637_v45)[_index225] = _1627_v35;
          let _index226 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1637_v45).length));
          (_1637_v45)[_index226] = _1627_v35;
          let _1638_v46;
          _1638_v46 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1639_v47;
          _1639_v47 = _dafny.Seq.of(_1571_v2, _1571_v2, _1571_v2);
          let _rhs282 = ((_1571_v2) ? (_1571_v2) : ((_1639_v47)[_module.__default.safeIndex((_this).f11, new BigNumber((_1639_v47).length))]));
          let _rhs283 = _1638_v46;
          _1571_v2 = _rhs282;
          _1638_v46 = _rhs283;
          (globalState).f1 = p0;
        }
        (globalState).f1 = (_this).f11;
        (globalState).f7 = (_1571_v2) || (_1571_v2);
        (globalState).f7 = _module.__default.fm0(globalState);
      }
      if (((_1571_v2) ? (!(_1571_v2)) : (!(p0).isEqualTo(new BigNumber(522))))) {
        let _1640_v48;
        _1640_v48 = _module.D15.create_DC34();
        _1640_v48 = ((_1571_v2) ? (_1640_v48) : (((_1571_v2) ? (_1640_v48) : (_1640_v48))));
        if (_1571_v2) {
          (globalState).f6 = (((_this).f12).Difference((_this).f12)).IsSubsetOf((_this).f12);
          _1570_v1 = _dafny.Seq.UnicodeFromString("xcrg");
          let _1641_v49;
          _1641_v49 = new _dafny.CodePoint('h'.codePointAt(0));
          _1641_v49 = (_1570_v1)[_module.__default.safeIndex(p0, new BigNumber((_1570_v1).length))];
          let _1642_v50;
          _1642_v50 = _module.D13.create_DC28(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("qaqirwuan"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("qaqirwuan")).length)), _1641_v49), (_this).f11, _module.__default.fm0(globalState));
          (globalState).f6 = (_1642_v50).dtor_cf42;
          let _1643_v51;
          let _nw266 = Array((new BigNumber(3)).toNumber());
          _nw266[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw266[(_dafny.ONE).toNumber()] = p0;
          _nw266[(new BigNumber(2)).toNumber()] = p0;
          _1643_v51 = _nw266;
          let _1644_v52;
          _1644_v52 = _dafny.Seq.of(_1643_v51, _1643_v51);
          let _1645_v53;
          _1645_v53 = _dafny.MultiSet.fromElements(new BigNumber(591));
          let _1646_v54;
          _1646_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,_1571_v2);
          _1644_v52 = _dafny.Seq.update(_1644_v52, _module.__default.safeIndex((((_1645_v53).contains(new BigNumber((_1646_v54).length))) ? ((_1645_v53).get(new BigNumber((_1646_v54).length))) : ((_this).f11)), new BigNumber((_1644_v52).length)), _1643_v51);
        } else {
          let _1647_v55;
          let _nw267 = Array((new BigNumber(7)).toNumber());
          _1647_v55 = _nw267;
          let _1648_v56;
          _1648_v56 = _dafny.Seq.of(_module.__default.fm52((_dafny.ZERO).minus((_this).f11), (_this).f11, globalState));
          let _1649_v57;
          _1649_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1647_v55,_dafny.Seq.Concat(_1648_v56, _1648_v56));
          _1649_v57 = (_1649_v57).update(_1647_v55, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-320)), ((_1650_v2) => function (_1651_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_1650_v2, new BigNumber((_dafny.Seq.of(_1650_v2, _1650_v2)).length), !(_1650_v2)),_1650_v2);
          })(_1571_v2)), _1648_v56));
          (globalState).f1 = ((_this).f11).minus(p0);
          let _1652_v58;
          _1652_v58 = _dafny.Seq.of(p0);
          (globalState).f6 = !((((_this).f11).multipliedBy(p0)).isLessThanOrEqualTo(new BigNumber((_1652_v58).length)));
          (globalState).f1 = (p0).minus((_this).f11);
          (globalState).f1 = _module.__default.safeModuloInt(p0, (_this).f11);
        }
        let _1653_v59;
        _1653_v59 = _dafny.Seq.of(p0);
        let _1654_v60;
        _1654_v60 = _dafny.Set.fromElements(_1653_v59, _dafny.Seq.Concat(_1653_v59, _1653_v59));
        _1654_v60 = (_1654_v60).Difference(_1654_v60);
        let _1655_v61;
        _1655_v61 = _module.D0.create_DC0((_this).f11);
        let _1656_v62;
        _1656_v62 = _dafny.Seq.of(false);
        let _1657_v63;
        _1657_v63 = _dafny.MultiSet.fromElements(p0);
        let _1658_v64;
        _1658_v64 = new _dafny.CodePoint('r'.codePointAt(0));
        if ((_this).fm5(_1655_v61, _1656_v62, new BigNumber((_1657_v63).cardinality()), _1658_v64, globalState)) {
          (globalState).f7 = _1571_v2;
          let _index227 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length));
          (_1569_v0)[_index227] = _dafny.Seq.Concat((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))], _dafny.Seq.update((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))], _module.__default.safeIndex(new BigNumber((_1653_v59).length), new BigNumber(((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))]).length)), _1658_v64));
          let _1659_v65;
          let _nw268 = Array((new BigNumber(29)).toNumber());
          _1659_v65 = _nw268;
          let _1660_v66;
          let _nw269 = new _module.C4();
          _nw269.__ctor(_1657_v63, _1571_v2, (_this).f12, new BigNumber(-299));
          _1660_v66 = _nw269;
          let _index228 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1659_v65).length));
          (_1659_v65)[_index228] = ((_1571_v2) ? (_1660_v66) : (_1660_v66));
          let _1661_v67;
          _1661_v67 = _dafny.MultiSet.fromElements((_1660_v66).f17, true, _1571_v2, (_1660_v66).f17, true);
          let _index229 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1659_v65).length));
          let _nw270 = new _module.C4();
          _nw270.__ctor(((false) ? (_1657_v63) : (_dafny.MultiSet.fromElements((_1653_v59)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f11)).length), new BigNumber((_1653_v59).length))]))), _1571_v2, (_this).f12, (_dafny.ZERO).minus((p0).multipliedBy((((_1661_v67).contains(true)) ? ((_1661_v67).get(true)) : (p0)))));
          (_1659_v65)[_index229] = _nw270;
          let _1662_v68;
          let _nw271 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _1662_v68 = _nw271;
          let _index230 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_1662_v68).length));
          (_1662_v68)[_index230] = (_this).f11;
          let _1663_v69;
          let _init33 = ((_1664_v2) => function (_1665_i3) {
            return _1664_v2;
          })(_1571_v2);
          let _nw272 = Array((_dafny.ONE).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw272.length); _i0_33++) {
            _nw272[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1663_v69 = _nw272;
          let _1666_v70;
          let _nw273 = Array((new BigNumber(3)).toNumber());
          _nw273[(_dafny.ZERO).toNumber()] = _1663_v69;
          _nw273[(_dafny.ONE).toNumber()] = _1663_v69;
          _nw273[(new BigNumber(2)).toNumber()] = _1663_v69;
          _1666_v70 = _nw273;
          let _index231 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_1666_v70).length));
          (_1666_v70)[_index231] = _1663_v69;
          let _1667_v71;
          _1667_v71 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_1656_v62, _module.__default.safeIndex(p0, new BigNumber((_1656_v62).length)), (_1660_v66).f17)).length),p0);
          let _1668_v72;
          _1668_v72 = _dafny.Seq.of(_1667_v71, _1667_v71);
          let _1669_v73;
          _1669_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,new BigNumber(((_1668_v72)[_module.__default.safeIndex(p0, new BigNumber((_1668_v72).length))]).length));
          let _1670_v74;
          _1670_v74 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1660_v66).f17);
          let _1671_v75;
          _1671_v75 = _dafny.Map.Empty.slice().updateUnsafe(!(_1571_v2),_1571_v2);
          let _index232 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_1662_v68).length));
          let _index233 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_1666_v70).length));
          let _rhs284 = (((_1669_v73).contains(!((_1660_v66).f17) || ((((_1670_v74).contains(new BigNumber((_1671_v75).length))) ? ((_1670_v74).get(new BigNumber((_1671_v75).length))) : (_1571_v2))))) ? ((_1669_v73).get(!((_1660_v66).f17) || ((((_1670_v74).contains(new BigNumber((_1671_v75).length))) ? ((_1670_v74).get(new BigNumber((_1671_v75).length))) : (_1571_v2))))) : ((_this).f11));
          let _rhs285 = _1663_v69;
          let _rhs286 = ((_1660_v66).f17) || (_module.__default.fm0(globalState));
          let _rhs287 = (_1656_v62)[_module.__default.safeIndex(p0, new BigNumber((_1656_v62).length))];
          let _rhs288 = (_1660_v66).f17;
          let _lhs232 = _1662_v68;
          let _lhs233 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_1662_v68).length));
          let _lhs234 = _1666_v70;
          let _lhs235 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_1666_v70).length));
          let _lhs236 = globalState;
          let _lhs237 = globalState;
          let _lhs238 = globalState;
          _lhs232[_lhs233] = _rhs284;
          _lhs234[_lhs235] = _rhs285;
          _lhs236.f6 = _rhs286;
          _lhs237.f7 = _rhs287;
          _lhs238.f7 = _rhs288;
          let _1672_v76;
          let _nw274 = Array((new BigNumber(18)).toNumber());
          _nw274[(_dafny.ZERO).toNumber()] = _1662_v68;
          _nw274[(_dafny.ONE).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(2)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(3)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(4)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(5)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(6)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(7)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(8)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(9)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(10)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(11)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(12)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(13)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(14)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(15)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(16)).toNumber()] = _1662_v68;
          _nw274[(new BigNumber(17)).toNumber()] = _1662_v68;
          _1672_v76 = _nw274;
          let _1673_v77;
          _1673_v77 = _1672_v76;
          let _1674_v78;
          let _nw275 = Array((new BigNumber(16)).toNumber());
          _nw275[(_dafny.ZERO).toNumber()] = _1673_v77;
          _nw275[(_dafny.ONE).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(2)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(3)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(4)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(5)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(6)).toNumber()] = _1672_v76;
          _nw275[(new BigNumber(7)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(8)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(9)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(10)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(11)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(12)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(13)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(14)).toNumber()] = _1673_v77;
          _nw275[(new BigNumber(15)).toNumber()] = _1673_v77;
          _1674_v78 = _nw275;
          let _1675_v79;
          _1675_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1674_v78,_1662_v68);
          let _1676_v80;
          _1676_v80 = _module.D11.create_DC22(_1675_v79);
          let _pat_let_tv44 = _1675_v79;
          let _1677_v81;
          _1677_v81 = _module.D11.create_DC22((function (_pat_let35_0) {
  return function (_1678_dt__update__tmp_h1) {
    return function (_pat_let36_0) {
      return function (_1679_dt__update_hcf30_h0) {
        return _module.D11.create_DC22(_1679_dt__update_hcf30_h0);
      }(_pat_let36_0);
    }(_pat_let_tv44);
  }(_pat_let35_0);
}(_1676_v80)).dtor_cf30);
          let _arr2 = (_1666_v70)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_1666_v70).length))];
          let _index234 = _module.__default.safeIndex(new BigNumber(828), new BigNumber(((_1666_v70)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_1666_v70).length))]).length));
          _arr2[_index234] = !((_1660_v66).f17) || ((_1660_v66).f17);
          let _index235 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_1662_v68).length));
          let _arr3 = (_1666_v70)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_1666_v70).length))];
          let _index236 = _module.__default.safeIndex(new BigNumber(828), new BigNumber(((_1666_v70)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_1666_v70).length))]).length));
          let _rhs289 = (_module.__default.safeModuloInt((_1662_v68)[_module.__default.safeIndex(new BigNumber(865), new BigNumber((_1662_v68).length))], (_1660_v66).fm6((_1660_v66).f17, new BigNumber((_1667_v71).length), globalState))).multipliedBy((_1662_v68)[_module.__default.safeIndex(new BigNumber(865), new BigNumber((_1662_v68).length))]);
          let _rhs290 = new BigNumber((_1570_v1).length);
          let _rhs291 = _1676_v80;
          let _rhs292 = (_this).f11;
          let _rhs293 = (_1660_v66).f17;
          let _lhs239 = _1662_v68;
          let _lhs240 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_1662_v68).length));
          let _lhs241 = globalState;
          let _lhs242 = globalState;
          let _lhs243 = (_1666_v70)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_1666_v70).length))];
          let _lhs244 = _module.__default.safeIndex(new BigNumber(828), new BigNumber(((_1666_v70)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_1666_v70).length))]).length));
          _lhs239[_lhs240] = _rhs289;
          _lhs241.f1 = _rhs290;
          _1676_v80 = _rhs291;
          _lhs242.f1 = _rhs292;
          _lhs243[_lhs244] = _rhs293;
        } else {
          _1656_v62 = _dafny.Seq.Concat(_1656_v62, _1656_v62);
          let _1680_v82;
          _1680_v82 = _dafny.MultiSet.fromElements(_1571_v2, _1571_v2, _1571_v2, _1571_v2, _1571_v2);
          let _1681_v83;
          _1681_v83 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1680_v82);
          let _1682_v84;
          _1682_v84 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_1571_v2), (((_1681_v83).contains((_this).f11)) ? ((_1681_v83).get((_this).f11)) : (_1680_v82)));
          _1682_v84 = _dafny.Seq.of(_module.__default.fm37(globalState));
          (globalState).f1 = _module.__default.safeModuloInt((_this).fm6(true, (_this).f11, globalState), p0);
          _1571_v2 = _1571_v2;
          let _rhs294 = (_this).fm6(_1571_v2, new BigNumber((_dafny.Seq.Concat(_1653_v59, _dafny.Seq.Create(_module.__default.abs(new BigNumber(95)), ((_1683_p0) => function (_1684_i4) {
            return _1683_p0;
          })(p0)))).length), globalState);
          let _rhs295 = (_1653_v59)[_module.__default.safeIndex(_module.__default.safeModuloInt(p0, (_this).f11), new BigNumber((_1653_v59).length))];
          let _rhs296 = _1571_v2;
          let _rhs297 = new BigNumber(474);
          let _lhs245 = globalState;
          let _lhs246 = globalState;
          let _lhs247 = globalState;
          let _lhs248 = globalState;
          _lhs245.f1 = _rhs294;
          _lhs246.f1 = _rhs295;
          _lhs247.f7 = _rhs296;
          _lhs248.f1 = _rhs297;
        }
        let _1685_v85;
        _1685_v85 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f11);
        _1685_v85 = _1685_v85;
      } else {
        let _1686_v86;
        _1686_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,_dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), function (_1687_i5) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }));
        let _1688_v87;
        _1688_v87 = _module.D14.create_DC30(_1686_v86);
        let _1689_v88;
        _1689_v88 = _module.D14.create_DC32(_1688_v87);
        let _1690_v89;
        _1690_v89 = _module.D7.create_DC13(((_this).f12).Difference((_this).f12), (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f11)), new BigNumber(-673));
        let _index237 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length));
        let _rhs298 = _1570_v1;
        let _rhs299 = _1689_v88;
        let _rhs300 = _module.__default.safeModuloInt(p0, new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(p0))).length));
        let _rhs301 = _1690_v89;
        let _rhs302 = _module.__default.safeDivisionInt((_this).f11, p0);
        let _lhs249 = _1569_v0;
        let _lhs250 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length));
        let _lhs251 = globalState;
        let _lhs252 = globalState;
        _lhs249[_lhs250] = _rhs298;
        _1689_v88 = _rhs299;
        _lhs251.f1 = _rhs300;
        _1690_v89 = _rhs301;
        _lhs252.f1 = _rhs302;
        (globalState).f1 = (_this).f11;
        let _1691_v90;
        _1691_v90 = _module.D15.create_DC36(p0, _1571_v2, _1571_v2, _1571_v2, (_this).f11);
        if ((_1691_v90).dtor_cf56) {
          (globalState).f6 = (_1571_v2) || (_1571_v2);
          (globalState).f1 = new BigNumber(841);
          let _1692_v91;
          _1692_v91 = _dafny.MultiSet.fromElements(p0);
          let _1693_v92;
          let _nw276 = new _module.C4();
          _nw276.__ctor(_1692_v91, _1571_v2, (_this).f12, (_this).f11);
          _1693_v92 = _nw276;
          let _1694_v93;
          _1694_v93 = _dafny.Map.Empty.slice().updateUnsafe(!(_1571_v2),_1693_v92);
          let _1695_v94;
          _1695_v94 = _dafny.Seq.of(_1693_v92);
          let _1696_v95;
          let _nw277 = Array((new BigNumber(26)).toNumber());
          _nw277[(_dafny.ZERO).toNumber()] = _1693_v92;
          _nw277[(_dafny.ONE).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(2)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(3)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(4)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(5)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(6)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(7)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(8)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(9)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(10)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(11)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(12)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(13)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(14)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(15)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(16)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(17)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(18)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(19)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(20)).toNumber()] = (((_1694_v93).contains(_1571_v2)) ? ((_1694_v93).get(_1571_v2)) : (_1693_v92));
          _nw277[(new BigNumber(21)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(22)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(23)).toNumber()] = _1693_v92;
          _nw277[(new BigNumber(24)).toNumber()] = (_1695_v94)[_module.__default.safeIndex((_module.D0.create_DC1(_1571_v2, (_1693_v92).f11, _1571_v2)).dtor_cf2, new BigNumber((_1695_v94).length))];
          _nw277[(new BigNumber(25)).toNumber()] = _1693_v92;
          _1696_v95 = _nw277;
          _1696_v95 = _1696_v95;
          let _1697_v96;
          _1697_v96 = _dafny.Seq.of(p0, (_1693_v92).f11);
          let _1698_v97;
          _1698_v97 = _module.D6.create_DC11(new BigNumber((_1697_v96).length), (_this).f11, p0);
          _1697_v96 = _dafny.Seq.Concat(_1697_v96, _module.__default.fm34(_1571_v2, _1698_v97, globalState));
          (globalState).f1 = p0;
        } else {
          (globalState).f6 = (_1571_v2) && (_module.__default.fm0(globalState));
          let _1699_v98;
          _1699_v98 = _module.D9.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(_1571_v2,p0));
          let _1700_v99;
          _1700_v99 = _dafny.Seq.of(_1571_v2);
          let _1701_v100;
          _1701_v100 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,new BigNumber((_1700_v99).length));
          let _pat_let_tv45 = _1701_v100;
          (globalState).f1 = _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1571_v2)).length), new BigNumber(((function (_pat_let37_0) {
            return function (_1702_dt__update__tmp_h2) {
              return function (_pat_let38_0) {
                return function (_1703_dt__update_hcf26_h0) {
                  return _module.D9.create_DC18(_1703_dt__update_hcf26_h0);
                }(_pat_let38_0);
              }(_pat_let_tv45);
            }(_pat_let37_0);
          }(_1699_v98)).dtor_cf26).length));
          _1701_v100 = _1701_v100;
          let _1704_v101;
          _1704_v101 = new _dafny.CodePoint('r'.codePointAt(0));
          _1704_v101 = _1704_v101;
          (globalState).f1 = (_this).f11;
        }
        let _1705_v102;
        let _nw278 = new _module.C1();
        _nw278.__ctor(_1571_v2, !(_1571_v2) || (_1571_v2));
        _1705_v102 = _nw278;
        _1705_v102 = _1705_v102;
        let _1706_v103;
        _1706_v103 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1707_v104;
        _1707_v104 = _dafny.Set.fromElements(_1706_v103, _1706_v103, _1706_v103, _1706_v103, _1706_v103);
        let _1708_v105;
        _1708_v105 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1707_v104).length),(_1705_v102).f18);
        let _1709_v106;
        _1709_v106 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))]).length),(((_1708_v105).contains((_this).f11)) ? ((_1708_v105).get((_this).f11)) : (_1571_v2)));
        let _1710_v107;
        _1710_v107 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,new BigNumber((_dafny.MultiSet.fromElements(_1571_v2, _1571_v2)).cardinality()));
        let _1711_v108;
        _1711_v108 = _dafny.MultiSet.fromElements(new BigNumber((_1570_v1).length));
        let _1712_v109;
        let _nw279 = Array((new BigNumber(16)).toNumber());
        _nw279[(_dafny.ZERO).toNumber()] = false;
        _nw279[(_dafny.ONE).toNumber()] = !(_1571_v2) || ((_1705_v102).f18);
        _nw279[(new BigNumber(2)).toNumber()] = _1571_v2;
        _nw279[(new BigNumber(3)).toNumber()] = ((_1705_v102).f18) && ((_1705_v102).f18);
        _nw279[(new BigNumber(4)).toNumber()] = (((_1708_v105).contains(new BigNumber((_dafny.Seq.of(_1706_v103, _1706_v103)).length))) ? ((_1708_v105).get(new BigNumber((_dafny.Seq.of(_1706_v103, _1706_v103)).length))) : ((_1705_v102).f19));
        _nw279[(new BigNumber(5)).toNumber()] = !(!((_1705_v102).f19));
        _nw279[(new BigNumber(6)).toNumber()] = !(new BigNumber((_dafny.MultiSet.fromElements((_1705_v102).f18, false)).cardinality())).isEqualTo((_this).f11);
        _nw279[(new BigNumber(7)).toNumber()] = !((_1705_v102).f19) || ((_1705_v102).f19);
        _nw279[(new BigNumber(8)).toNumber()] = (_1705_v102).f19;
        _nw279[(new BigNumber(9)).toNumber()] = ((_this).f11).isLessThanOrEqualTo(p0);
        _nw279[(new BigNumber(10)).toNumber()] = (_1705_v102).f19;
        _nw279[(new BigNumber(11)).toNumber()] = (_1705_v102).f19;
        _nw279[(new BigNumber(12)).toNumber()] = (p0).isEqualTo(new BigNumber(((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))]).length));
        _nw279[(new BigNumber(13)).toNumber()] = (new BigNumber((_1710_v107).length)).isLessThanOrEqualTo(new BigNumber(-68));
        _nw279[(new BigNumber(14)).toNumber()] = (_1711_v108).IsProperSubsetOf(_1711_v108);
        _nw279[(new BigNumber(15)).toNumber()] = !((_1705_v102).f19) || (_1571_v2);
        _1712_v109 = _nw279;
        let _1713_v110;
        _1713_v110 = _module.D14.create_DC31((_1705_v102).f18, (_this).f11, new BigNumber(341));
        let _index238 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1712_v109).length));
        (_1712_v109)[_index238] = ((_this).f11).isLessThan((_1713_v110).dtor_cf46);
        let _index239 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1712_v109).length));
        (_1712_v109)[_index239] = (_1705_v102).f19;
      }
      let _1714_v111;
      let _init34 = function (_1715_i6) {
        return (_1715_i6).multipliedBy((_this).f11);
      };
      let _nw280 = Array((new BigNumber(22)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw280.length); _i0_34++) {
        _nw280[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _1714_v111 = _nw280;
      let _index240 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length));
      (_1714_v111)[_index240] = p0;
      let _index241 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length));
      (_1714_v111)[_index241] = (_this).f11;
      if (_1571_v2) {
        let _1716_v112;
        _1716_v112 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1717_v113;
        _1717_v113 = _dafny.Seq.of(_1571_v2);
        let _1718_v114;
        _1718_v114 = _dafny.Set.fromElements((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]);
        let _1719_v115;
        _1719_v115 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,_1571_v2);
        let _1720_v116;
        _1720_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1718_v114,_1719_v115);
        let _1721_v117;
        _1721_v117 = _dafny.MultiSet.fromElements(_1716_v112, _module.__default.fm20(_1717_v113, (_dafny.ZERO).minus((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]), _1571_v2, globalState), _module.__default.fm20(_1717_v113, new BigNumber((_1720_v116).length), _1571_v2, globalState));
        (globalState).f1 = ((((_1721_v117).contains(_1716_v112)) ? ((_1721_v117).get(_1716_v112)) : ((_this).f11))).multipliedBy(new BigNumber((_dafny.Seq.of((_1717_v113)[_module.__default.safeIndex((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))], new BigNumber((_1717_v113).length))], _1571_v2)).length));
        let _1722_v118;
        _1722_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v2,(((_this).f15).dtor_cf8).multipliedBy(new BigNumber((_1570_v1).length)));
        _1722_v118 = (_1722_v118).update(((_this).f11).isLessThan(new BigNumber(102)), (_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]);
        _1570_v1 = (_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))];
        (globalState).f1 = (_this).f11;
        let _1723_v119;
        _1723_v119 = _module.D0.create_DC0((p0).multipliedBy((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]));
        let _source25 = _1723_v119;
        if (_source25.is_DC1) {
          let _1724___mcc_h7 = (_source25).cf1;
          let _1725___mcc_h8 = (_source25).cf2;
          let _1726___mcc_h9 = (_source25).cf3;
          let _1727_cf3 = _1726___mcc_h9;
          let _1728_cf2 = _1725___mcc_h8;
          let _1729_cf1 = _1724___mcc_h7;
          let _1730_v120;
          _1730_v120 = _dafny.Map.Empty.slice().updateUnsafe((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))],true);
          (globalState).f1 = (new BigNumber(((_1730_v120).Merge(_1730_v120)).length)).minus((p0).multipliedBy((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]));
          let _1731_v121;
          _1731_v121 = _dafny.MultiSet.fromElements(_1727_cf3);
          let _1732_v122;
          _1732_v122 = _dafny.Map.Empty.slice().updateUnsafe((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))],(_this).f11);
          let _1733_v123;
          _1733_v123 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_1728_cf2),_1732_v122);
          _1571_v2 = ((_dafny.MultiSet.fromElements((_this).fm4(new BigNumber((_1731_v121).cardinality()), _1733_v123, p0, _1729_cf1, globalState), _1729_cf1)).Difference(_dafny.MultiSet.fromElements(_1727_cf3))).IsProperSubsetOf(_1731_v121);
          _1717_v113 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_1717_v113, _1717_v113), _1717_v113), _module.__default.safeIndex(_module.__default.safeModuloInt(_1728_cf2, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(559))).cardinality())), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1717_v113, _1717_v113), _1717_v113)).length)), (new BigNumber(518)).isLessThanOrEqualTo((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]));
          let _1734_v124;
          _1734_v124 = _dafny.Map.Empty.slice().updateUnsafe(_1716_v112,_1571_v2);
          let _1735_v125;
          let _nw281 = Array((new BigNumber(26)).toNumber());
          _nw281[(_dafny.ZERO).toNumber()] = !((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]).isEqualTo(p0);
          _nw281[(_dafny.ONE).toNumber()] = _dafny.areEqual(_module.__default.fm1(globalState), _dafny.Seq.UnicodeFromString("vlp"));
          _nw281[(new BigNumber(2)).toNumber()] = (_1717_v113)[_module.__default.safeIndex(p0, new BigNumber((_1717_v113).length))];
          _nw281[(new BigNumber(3)).toNumber()] = !(!(_1727_cf3) || (_1729_cf1));
          _nw281[(new BigNumber(4)).toNumber()] = (((_1734_v124).contains(_1716_v112)) ? ((_1734_v124).get(_1716_v112)) : (_1729_cf1));
          _nw281[(new BigNumber(5)).toNumber()] = !(true);
          _nw281[(new BigNumber(6)).toNumber()] = (_1717_v113)[_module.__default.safeIndex(new BigNumber(476), new BigNumber((_1717_v113).length))];
          _nw281[(new BigNumber(7)).toNumber()] = _1571_v2;
          _nw281[(new BigNumber(8)).toNumber()] = _1727_cf3;
          _nw281[(new BigNumber(9)).toNumber()] = _1727_cf3;
          _nw281[(new BigNumber(10)).toNumber()] = _1727_cf3;
          _nw281[(new BigNumber(11)).toNumber()] = _1571_v2;
          _nw281[(new BigNumber(12)).toNumber()] = _1729_cf1;
          _nw281[(new BigNumber(13)).toNumber()] = _1727_cf3;
          _nw281[(new BigNumber(14)).toNumber()] = _1727_cf3;
          _nw281[(new BigNumber(15)).toNumber()] = _1571_v2;
          _nw281[(new BigNumber(16)).toNumber()] = !((_1731_v121).IsProperSubsetOf(_dafny.MultiSet.fromElements(!(_1571_v2))));
          _nw281[(new BigNumber(17)).toNumber()] = _1571_v2;
          _nw281[(new BigNumber(18)).toNumber()] = _1729_cf1;
          _nw281[(new BigNumber(19)).toNumber()] = _1727_cf3;
          _nw281[(new BigNumber(20)).toNumber()] = !(_1729_cf1);
          _nw281[(new BigNumber(21)).toNumber()] = _1571_v2;
          _nw281[(new BigNumber(22)).toNumber()] = (!(_1729_cf1)) || (_1727_cf3);
          _nw281[(new BigNumber(23)).toNumber()] = _1727_cf3;
          _nw281[(new BigNumber(24)).toNumber()] = (_1729_cf1) || (_1571_v2);
          _nw281[(new BigNumber(25)).toNumber()] = _1571_v2;
          _1735_v125 = _nw281;
          let _index242 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1735_v125).length));
          (_1735_v125)[_index242] = false;
          let _1736_v127;
          _1736_v127 = _dafny.Seq.of(_1718_v114);
          let _index243 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1735_v125).length));
          let _index244 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length));
          let _rhs303 = (function () {
            let _coll66 = new _dafny.Set();
            for (const _compr_66 of (_1718_v114).Elements) {
              let _1737_v126 = _compr_66;
              if ((_1718_v114).contains(_1737_v126)) {
                _coll66.add(_module.__default.safeModuloInt(_1737_v126, new BigNumber(575)));
              }
            }
            return _coll66;
          }()).IsDisjointFrom((_1736_v127)[_module.__default.safeIndex(p0, new BigNumber((_1736_v127).length))]);
          let _rhs304 = _module.__default.fm1(globalState);
          let _rhs305 = (_dafny.ZERO).minus(_module.__default.fm2(p0, p0, globalState));
          let _lhs253 = _1735_v125;
          let _lhs254 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1735_v125).length));
          let _lhs255 = _1569_v0;
          let _lhs256 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length));
          let _lhs257 = globalState;
          _lhs253[_lhs254] = _rhs303;
          _lhs255[_lhs256] = _rhs304;
          _lhs257.f1 = _rhs305;
        } else {
          let _1738___mcc_h10 = (_source25).cf0;
          let _1739_cf0 = _1738___mcc_h10;
          let _1740_v128;
          let _init35 = ((_1741_v2) => function (_1742_i7) {
            return _dafny.MultiSet.fromElements(!(!(_1741_v2)), (_module.D10.create_DC21(_1741_v2, _1741_v2)).dtor_cf28, !(_1741_v2));
          })(_1571_v2);
          let _nw282 = Array((new BigNumber(4)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw282.length); _i0_35++) {
            _nw282[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1740_v128 = _nw282;
          let _rhs306 = _1740_v128;
          let _rhs307 = new _dafny.CodePoint('j'.codePointAt(0));
          let _rhs308 = _module.__default.safeDivisionInt((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))], ((_1571_v2) ? (new BigNumber(412)) : (p0)));
          let _rhs309 = _dafny.Seq.IsProperPrefixOf(_1570_v1, ((_1571_v2) ? (_dafny.Seq.UnicodeFromString("qqo")) : (_1570_v1)));
          let _lhs258 = globalState;
          let _lhs259 = globalState;
          _1740_v128 = _rhs306;
          _1716_v112 = _rhs307;
          _lhs258.f1 = _rhs308;
          _lhs259.f7 = _rhs309;
          (globalState).f6 = _1571_v2;
          let _1743_v129;
          let _nw283 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
          _1743_v129 = _nw283;
          let _rhs310 = !(_1571_v2);
          let _rhs311 = _1743_v129;
          let _lhs260 = globalState;
          _lhs260.f7 = _rhs310;
          _1743_v129 = _rhs311;
          let _1744_v130;
          let _init36 = ((_1745_v2) => function (_1746_i8) {
            return _1745_v2;
          })(_1571_v2);
          let _nw284 = Array((new BigNumber(2)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw284.length); _i0_36++) {
            _nw284[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1744_v130 = _nw284;
          let _1747_v131;
          _1747_v131 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f11,new BigNumber(812))).length),(_this).f11);
          let _index245 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_1744_v130).length));
          (_1744_v130)[_index245] = (_1747_v131).equals(_1747_v131);
          let _1748_v132;
          _1748_v132 = _dafny.Map.Empty.slice().updateUnsafe(_1714_v111,_1571_v2);
          let _1749_v133;
          _1749_v133 = _module.D10.create_DC20(_1748_v132);
          let _1750_v134;
          _1750_v134 = _dafny.Map.Empty.slice().updateUnsafe(_1749_v133,(p0).minus((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]));
          let _index246 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_1744_v130).length));
          let _rhs312 = _1571_v2;
          let _rhs313 = _1750_v134;
          let _rhs314 = _1714_v111;
          let _lhs261 = _1744_v130;
          let _lhs262 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_1744_v130).length));
          _lhs261[_lhs262] = _rhs312;
          _1750_v134 = _rhs313;
          _1714_v111 = _rhs314;
        }
      } else {
        let _init37 = ((_1751_v111) => function (_1752_i9) {
          return (_1752_i9).multipliedBy((_1751_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1751_v111).length))]);
        })(_1714_v111);
        let _nw285 = Array((new BigNumber(7)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw285.length); _i0_37++) {
          _nw285[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1714_v111 = _nw285;
        let _1753_v135;
        let _nw286 = new _module.C5();
        _nw286.__ctor();
        _1753_v135 = _nw286;
        let _1754_v136;
        let _nw287 = Array((new BigNumber(14)).toNumber());
        _nw287[(_dafny.ZERO).toNumber()] = _1753_v135;
        _nw287[(_dafny.ONE).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(2)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(3)).toNumber()] = (_module.D20.create_DC44(_1753_v135)).dtor_cf74;
        _nw287[(new BigNumber(4)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(5)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(6)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(7)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(8)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(9)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(10)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(11)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(12)).toNumber()] = _1753_v135;
        _nw287[(new BigNumber(13)).toNumber()] = _1753_v135;
        _1754_v136 = _nw287;
        let _1755_v137;
        _1755_v137 = _dafny.Set.fromElements(_1754_v136, _1754_v136, ((_1571_v2) ? (_1754_v136) : (_1754_v136)));
        _1755_v137 = ((_1755_v137).Union(_1755_v137)).Intersect(_dafny.Set.fromElements(_1754_v136, _1754_v136));
        _1570_v1 = _dafny.Seq.Concat(_module.__default.fm1(globalState), _1570_v1);
        let _1756_v138;
        _1756_v138 = _dafny.Seq.of(_1571_v2);
        (globalState).f7 = (_1756_v138)[_module.__default.safeIndex(new BigNumber(287), new BigNumber((_1756_v138).length))];
        let _1757_v139;
        _1757_v139 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1758_v140;
        _1758_v140 = _dafny.Map.Empty.slice().updateUnsafe(false,_1571_v2);
        let _1759_v141;
        _1759_v141 = _dafny.MultiSet.fromElements(_1758_v140, _1758_v140, _1758_v140);
        let _1760_v142;
        let _nw288 = Array((new BigNumber(8)).toNumber());
        _nw288[(_dafny.ZERO).toNumber()] = !(p0).isEqualTo(new BigNumber(883));
        _nw288[(_dafny.ONE).toNumber()] = _1571_v2;
        _nw288[(new BigNumber(2)).toNumber()] = _module.__default.fm0(globalState);
        _nw288[(new BigNumber(3)).toNumber()] = _1571_v2;
        _nw288[(new BigNumber(4)).toNumber()] = _dafny.areEqual(_dafny.Seq.UnicodeFromString("c"), (_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))]);
        _nw288[(new BigNumber(5)).toNumber()] = _1571_v2;
        _nw288[(new BigNumber(6)).toNumber()] = !_dafny.Seq.contains((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))], _1757_v139);
        _nw288[(new BigNumber(7)).toNumber()] = (_1759_v141).IsDisjointFrom(_1759_v141);
        _1760_v142 = _nw288;
        let _1761_v144;
        _1761_v144 = _dafny.MultiSet.fromElements((_this).f11, (_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))], p0);
        let _1762_v145;
        _1762_v145 = _dafny.Seq.of(_1761_v144);
        let _index247 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_1760_v142).length));
        (_1760_v142)[_index247] = (_this).fm4((_this).f11, function () {
          let _coll67 = new _dafny.Map();
          for (const _compr_67 of (_1762_v145).Elements) {
            let _1763_v143 = _compr_67;
            if (_dafny.Seq.contains(_1762_v145, _1763_v143)) {
              _coll67.push([_1763_v143,_dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(p0))]);
            }
          }
          return _coll67;
        }(), (_this).f11, (_1756_v138)[_module.__default.safeIndex(p0, new BigNumber((_1756_v138).length))], globalState);
        let _1764_v146;
        _1764_v146 = _dafny.Set.fromElements(_1757_v139);
        let _index248 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_1760_v142).length));
        (_1760_v142)[_index248] = (_1764_v146).IsSubsetOf(_1764_v146);
      }
      let _1765_v147;
      _1765_v147 = _dafny.MultiSet.fromElements(_1571_v2, _1571_v2, _1571_v2);
      let _1766_v148;
      _1766_v148 = _dafny.Set.fromElements(new BigNumber((_1570_v1).length));
      let _1767_v149;
      _1767_v149 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1766_v148).length),_1571_v2);
      let _1768_v150;
      _1768_v150 = _module.D12.create_DC26(_1767_v149);
      let _1769_v151;
      _1769_v151 = _dafny.Map.Empty.slice().updateUnsafe(_1768_v150,_1765_v147);
      let _1770_v152;
      let _nw289 = Array((new BigNumber(7)).toNumber());
      _nw289[(_dafny.ZERO).toNumber()] = _1765_v147;
      _nw289[(_dafny.ONE).toNumber()] = _1765_v147;
      _nw289[(new BigNumber(2)).toNumber()] = _1765_v147;
      _nw289[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_1571_v2, _1571_v2);
      _nw289[(new BigNumber(4)).toNumber()] = _1765_v147;
      _nw289[(new BigNumber(5)).toNumber()] = _1765_v147;
      _nw289[(new BigNumber(6)).toNumber()] = (((_1769_v151).contains(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]),_1571_v2)))) ? ((_1769_v151).get(_module.D12.create_DC26(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1714_v111)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_1714_v111).length))]),_1571_v2)))) : (_1765_v147));
      _1770_v152 = _nw289;
      r0 = _1770_v152;
      let _1771_v153;
      _1771_v153 = _dafny.Seq.of(true, false);
      let _1772_v154;
      _1772_v154 = _dafny.MultiSet.fromElements(new BigNumber(((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))]).length));
      let _1773_v155;
      _1773_v155 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))], _module.__default.safeIndex((_this).f11, new BigNumber(((_1569_v0)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_1569_v0).length))]).length)), _module.__default.fm20(_1771_v153, new BigNumber((_1570_v1).length), _1571_v2, globalState)),_1772_v154);
      r1 = (_1773_v155).Merge((_1773_v155).Merge(_1773_v155));
      return [r0, r1];
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f12 = _dafny.Set.Empty;
      this._f11 = _dafny.ZERO;
      this._f14 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f14, f12, f11) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f12 = f12;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_module.D0.create_DC1(false, new BigNumber((_dafny.Seq.UnicodeFromString("s")).length), true)).dtor_cf1);
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      let _source26 = _module.D0.create_DC0((_this).f11);
      if (_source26.is_DC1) {
        let _1774___mcc_h0 = (_source26).cf1;
        let _1775___mcc_h1 = (_source26).cf2;
        let _1776___mcc_h2 = (_source26).cf3;
        let _1777_cf3 = _1776___mcc_h2;
        let _1778_cf2 = _1775___mcc_h1;
        let _1779_cf1 = _1774___mcc_h0;
        return _1778_cf2;
      } else {
        let _1780___mcc_h3 = (_source26).cf0;
        let _1781_cf0 = _1780___mcc_h3;
        return (_this).f11;
      }
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(function (_source27) {
        if (_source27.is_DC1) {
          let _1782___mcc_h0 = (_source27).cf1;
          let _1783___mcc_h1 = (_source27).cf2;
          let _1784___mcc_h2 = (_source27).cf3;
          let _1785_cf3 = _1784___mcc_h2;
          let _1786_cf2 = _1783___mcc_h1;
          let _1787_cf1 = _1782___mcc_h0;
          return ((_this).f11).plus((_this).f11);
        } else {
          let _1788___mcc_h3 = (_source27).cf0;
          let _1789_cf0 = _1788___mcc_h3;
          return (_this).f11;
        }
      }(((true) ? (_module.D0.create_DC1(!(false), (_this).f11, true)) : (_module.D0.create_DC1(true, (_this).f11, false)))));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.IsProperPrefixOf((_dafny.Seq.Create(_module.__default.abs(new BigNumber(97)), function (_1790_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("kdsxn"));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let _1791_v0;
      _1791_v0 = true;
      let _1792_v1;
      _1792_v1 = _dafny.MultiSet.fromElements(_1791_v0, _1791_v0, _1791_v0, _1791_v0);
      let _1793_v2;
      _1793_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1791_v0,_dafny.Seq.UnicodeFromString("lupwwwla"));
      let _hi2 = new BigNumber(((((_1793_v2).contains(_1791_v0)) ? ((_1793_v2).get(_1791_v0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-92)), ((_1794_p0) => function (_1795_i1) {
        return _1794_p0;
      })(p0))))).length);
      for (let _1796_i0 = new BigNumber((_1792_v1).cardinality()); _1796_i0.isLessThan(_hi2); _1796_i0 = _1796_i0.plus(_dafny.ONE)) {
        let _1797_v3;
        _1797_v3 = _dafny.Seq.UnicodeFromString("yyslj");
        let _1798_v4;
        _1798_v4 = _dafny.MultiSet.fromElements(_1797_v3);
        let _1799_v5;
        _1799_v5 = _module.D2.create_DC3(_1798_v4);
        let _1800_v6;
        _1800_v6 = _dafny.Seq.of(_1791_v0, _1791_v0, !(((_1799_v5).dtor_cf5).IsSubsetOf(_1798_v4)), (_1791_v0) === (false));
        _1800_v6 = _dafny.Seq.update(_1800_v6, _module.__default.safeIndex((_this).f11, new BigNumber((_1800_v6).length)), _1791_v0);
        (globalState).f1 = _1796_i0;
        (globalState).f6 = (_dafny.Set.fromElements(_1791_v0)).IsSubsetOf((_this).f12);
        let _1801_v7;
        _1801_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1791_v0);
        let _1802_v8;
        _1802_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_1793_v2).contains(_1791_v0)) ? ((_1793_v2).get(_1791_v0)) : (_1797_v3))).length),new BigNumber((_1801_v7).length));
        let _1803_v9;
        let _nw290 = new _module.C2();
        _nw290.__ctor(_module.__default.fm20(_1800_v6, (_this).f11, _1791_v0, globalState), new BigNumber(((_1802_v8).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-721),(_this).f11))).length));
        _1803_v9 = _nw290;
      }
      let _1804_v10;
      _1804_v10 = _dafny.MultiSet.fromElements((_this).f11);
      let _1805_v11;
      let _nw291 = new _module.C4();
      _nw291.__ctor(_1804_v10, _1791_v0, ((_this).f12).Intersect((_this).f12), _module.__default.safeModuloInt(new BigNumber(-754), new BigNumber((_1792_v1).cardinality())));
      _1805_v11 = _nw291;
      let _hi3 = (_this).f11;
      for (let _1806_i2 = (_this).f11; _1806_i2.isLessThan(_hi3); _1806_i2 = _1806_i2.plus(_dafny.ONE)) {
        let _1807_v12;
        _1807_v12 = _module.D9.create_DC19();
        let _1808_v13;
        _1808_v13 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f11), (_this).f11);
        let _1809_v14;
        let _nw292 = Array((new BigNumber(15)).toNumber());
        _nw292[(_dafny.ZERO).toNumber()] = _1807_v12;
        _nw292[(_dafny.ONE).toNumber()] = _1807_v12;
        _nw292[(new BigNumber(2)).toNumber()] = (((_1805_v11).f17) ? (_1807_v12) : (_1807_v12));
        _nw292[(new BigNumber(3)).toNumber()] = _1807_v12;
        _nw292[(new BigNumber(4)).toNumber()] = _module.D9.create_DC19();
        _nw292[(new BigNumber(5)).toNumber()] = _1807_v12;
        _nw292[(new BigNumber(6)).toNumber()] = _module.D9.create_DC19();
        _nw292[(new BigNumber(7)).toNumber()] = _1807_v12;
        _nw292[(new BigNumber(8)).toNumber()] = _1807_v12;
        _nw292[(new BigNumber(9)).toNumber()] = _1807_v12;
        _nw292[(new BigNumber(10)).toNumber()] = _1807_v12;
        _nw292[(new BigNumber(11)).toNumber()] = _module.__default.fm42(_dafny.Set.fromElements(_1808_v13), globalState);
        _nw292[(new BigNumber(12)).toNumber()] = _module.D9.create_DC19();
        _nw292[(new BigNumber(13)).toNumber()] = ((_1791_v0) ? (_1807_v12) : (_1807_v12));
        _nw292[(new BigNumber(14)).toNumber()] = _1807_v12;
        _1809_v14 = _nw292;
        let _index249 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1809_v14).length));
        (_1809_v14)[_index249] = _1807_v12;
        let _1810_v15;
        _1810_v15 = _module.D8.create_DC16((_1805_v11).f17, _1806_i2);
        let _1811_v16;
        _1811_v16 = _module.D14.create_DC31(_1791_v0, (_dafny.ZERO).minus(_1806_i2), (_this).f11);
        let _1812_v17;
        _1812_v17 = _dafny.Seq.of((_1805_v11).f17, _1791_v0, _1791_v0);
        let _1813_v18;
        let _nw293 = Array((new BigNumber(26)).toNumber());
        _nw293[(_dafny.ZERO).toNumber()] = (new BigNumber(537)).isLessThanOrEqualTo(_1806_i2);
        _nw293[(_dafny.ONE).toNumber()] = (((_1805_v11).f17) ? (false) : ((_1805_v11).f17));
        _nw293[(new BigNumber(2)).toNumber()] = true;
        _nw293[(new BigNumber(3)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(4)).toNumber()] = (_1810_v15).dtor_cf23;
        _nw293[(new BigNumber(5)).toNumber()] = ((_this).f11).isLessThan((_this).f11);
        _nw293[(new BigNumber(6)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(7)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(8)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(9)).toNumber()] = _module.__default.fm0(globalState);
        _nw293[(new BigNumber(10)).toNumber()] = (new BigNumber(861)).isLessThan(_1806_i2);
        _nw293[(new BigNumber(11)).toNumber()] = ((_this).f11).isLessThan((_this).f11);
        _nw293[(new BigNumber(12)).toNumber()] = false;
        _nw293[(new BigNumber(13)).toNumber()] = (_1806_i2).isLessThanOrEqualTo(_1806_i2);
        _nw293[(new BigNumber(14)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(15)).toNumber()] = (_1811_v16).dtor_cf45;
        _nw293[(new BigNumber(16)).toNumber()] = !(new BigNumber((_dafny.Set.fromElements((_1805_v11).f17, (_1805_v11).f17)).length)).isEqualTo(_1806_i2);
        _nw293[(new BigNumber(17)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(18)).toNumber()] = _1791_v0;
        _nw293[(new BigNumber(19)).toNumber()] = _1791_v0;
        _nw293[(new BigNumber(20)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(21)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(22)).toNumber()] = (_1812_v17)[_module.__default.safeIndex((_1808_v13)[_module.__default.safeIndex((_dafny.ZERO).minus(_1806_i2), new BigNumber((_1808_v13).length))], new BigNumber((_1812_v17).length))];
        _nw293[(new BigNumber(23)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(24)).toNumber()] = (_1805_v11).f17;
        _nw293[(new BigNumber(25)).toNumber()] = (_1805_v11).f17;
        _1813_v18 = _nw293;
        let _1814_v19;
        _1814_v19 = _dafny.Seq.UnicodeFromString("kdckwdveg");
        let _1815_v20;
        _1815_v20 = _1814_v19;
        let _1816_v21;
        _1816_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1815_v20,(_1805_v11).f17);
        let _index250 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1813_v18).length));
        (_1813_v18)[_index250] = (((((_1816_v21).contains(_1815_v20)) ? ((_1816_v21).get(_1815_v20)) : ((_1805_v11).f17))) ? ((_1805_v11).f17) : (false));
        let _1817_v22;
        _1817_v22 = _dafny.Map.Empty.slice().updateUnsafe((_1805_v11).f17,(_1805_v11).f17);
        let _index251 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1809_v14).length));
        let _index252 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1813_v18).length));
        let _rhs315 = _1807_v12;
        let _rhs316 = !(_1806_i2).isEqualTo((_1805_v11).fm6((_1805_v11).f17, _module.__default.fm2(_1806_i2, _1806_i2, globalState), globalState));
        let _rhs317 = (_1812_v17)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_1817_v22).length), _1806_i2), new BigNumber((_1812_v17).length))];
        let _rhs318 = (_1805_v11).f17;
        let _lhs263 = _1809_v14;
        let _lhs264 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1809_v14).length));
        let _lhs265 = globalState;
        let _lhs266 = _1813_v18;
        let _lhs267 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1813_v18).length));
        let _lhs268 = globalState;
        _lhs263[_lhs264] = _rhs315;
        _lhs265.f6 = _rhs316;
        _lhs266[_lhs267] = _rhs317;
        _lhs268.f6 = _rhs318;
        let _1818_v23;
        _1818_v23 = _module.D16.create_DC39((_this).f11, (_1813_v18)[_module.__default.safeIndex(new BigNumber(665), new BigNumber((_1813_v18).length))], _1817_v22, p0);
        let _1819_v24;
        _1819_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1805_v11,_1812_v17);
        let _1820_v25;
        let _nw294 = new _module.C3();
        _nw294.__ctor(_module.__default.safeDivisionInt((_1818_v23).dtor_cf62, (_this).f11), (_dafny.ZERO).minus((_this).f11), new BigNumber((_1819_v24).length), ((_this).f12).Difference((_this).f12));
        _1820_v25 = _nw294;
        let _index253 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1813_v18).length));
        let _rhs319 = !((_1805_v11).f17);
        let _rhs320 = (_1805_v11).f17;
        let _rhs321 = _1820_v25;
        let _rhs322 = (_1820_v25).f11;
        let _lhs269 = _1813_v18;
        let _lhs270 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1813_v18).length));
        let _lhs271 = globalState;
        _1791_v0 = _rhs319;
        _lhs269[_lhs270] = _rhs320;
        _1820_v25 = _rhs321;
        _lhs271.f1 = _rhs322;
        (globalState).f6 = !_dafny.areEqual(_1814_v19, ((!((_1813_v18)[_module.__default.safeIndex(new BigNumber(665), new BigNumber((_1813_v18).length))])) ? (_dafny.Seq.update(_1814_v19, _module.__default.safeIndex(_1806_i2, new BigNumber((_1814_v19).length)), p0)) : (_1814_v19)));
        let _1821_v26;
        _1821_v26 = _module.D7.create_DC13((_this).f12, _1806_i2, (new BigNumber((_1814_v19).length)).minus((_this).f11));
        let _1822_v27;
        _1822_v27 = _dafny.Set.fromElements((_1820_v25).f11);
        let _1823_v28;
        _1823_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Seq.update(_1814_v19, _module.__default.safeIndex(_1806_i2, new BigNumber((_1814_v19).length)), p0)).length));
        let _pat_let_tv46 = _1823_v28;
        _1821_v26 = (((_1822_v27).IsSubsetOf(_1822_v27)) ? (_module.__default.fm36((_1813_v18)[_module.__default.safeIndex(new BigNumber(665), new BigNumber((_1813_v18).length))], _module.__default.fm0(globalState), (_1820_v25).f11, globalState)) : (function (_pat_let39_0) {
          return function (_1824_dt__update__tmp_h0) {
            return function (_pat_let40_0) {
              return function (_1825_dt__update_hcf18_h0) {
                return function (_pat_let41_0) {
                  return function (_1826_dt__update_hcf20_h0) {
                    return _module.D7.create_DC13(_1825_dt__update_hcf18_h0, (_1824_dt__update__tmp_h0).dtor_cf19, _1826_dt__update_hcf20_h0);
                  }(_pat_let41_0);
                }(new BigNumber((_pat_let_tv46).length));
              }(_pat_let40_0);
            }((_this).f12);
          }(_pat_let39_0);
        }(_1821_v26)));
      }
      (globalState).f1 = (_dafny.ZERO).minus((_this).f11);
      let _1827_v29;
      _1827_v29 = _dafny.Seq.UnicodeFromString("sctxpfjb");
      let _hi4 = (new BigNumber(-812)).minus(new BigNumber((_1827_v29).length));
      for (let _1828_i3 = (_dafny.ZERO).minus((_this).f11); _1828_i3.isLessThan(_hi4); _1828_i3 = _1828_i3.plus(_dafny.ONE)) {
        let _1829_v30;
        let _1830_v31;
        let _out31;
        let _out32;
        let _outcollector10 = (_1805_v11).m1(_1828_i3, globalState);
        _out31 = _outcollector10[0];
        _out32 = _outcollector10[1];
        _1829_v30 = _out31;
        _1830_v31 = _out32;
        let _1831_v32;
        _1831_v32 = _dafny.Seq.of((_1805_v11).f17, (_1805_v11).f17);
        let _1832_v33;
        let _nw295 = new _module.C1();
        _nw295.__ctor((_1805_v11).f17, (_this).fm5(_module.D0.create_DC0((_this).f11), _1831_v32, _1828_i3, p0, globalState));
        _1832_v33 = _nw295;
        if ((_1832_v33).f18) {
          let _1833_v34;
          _1833_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1828_i3,p0);
          let _1834_v35;
          _1834_v35 = _dafny.Seq.of(_1828_i3);
          let _1835_v36;
          _1835_v36 = _dafny.Seq.of(new BigNumber((_1833_v34).length), (_1828_i3).minus((_this).f11), _module.__default.safeDivisionInt(new BigNumber((_1834_v35).length), _1828_i3));
          _1835_v36 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(534)), ((_1836_i3) => function (_1837_i4) {
            return _1836_i3;
          })(_1828_i3));
          let _1838_v37;
          let _init38 = function (_1839_i5) {
            return _module.__default.safeModuloInt(_1839_i5, (_this).f11);
          };
          let _nw296 = Array((new BigNumber(24)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw296.length); _i0_38++) {
            _nw296[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1838_v37 = _nw296;
          let _index254 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_1838_v37).length));
          (_1838_v37)[_index254] = ((_this).f11).multipliedBy((_this).f11);
          let _index255 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_1838_v37).length));
          (_1838_v37)[_index255] = _module.__default.safeDivisionInt((_this).f11, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1828_i3,_dafny.Seq.update(_1827_v29, _module.__default.safeIndex(_1828_i3, new BigNumber((_1827_v29).length)), p0))).length)));
          let _1840_v38;
          _1840_v38 = _dafny.Set.fromElements((_this).fm6((_1832_v33).f19, ((((_1805_v11).f16).contains(new BigNumber((_1831_v32).length))) ? (((_1805_v11).f16).get(new BigNumber((_1831_v32).length))) : ((_1838_v37)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_1838_v37).length))])), globalState), new BigNumber(532), _1828_i3);
          let _1841_v39;
          _1841_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_1840_v38).Union(_1840_v38));
          _1840_v38 = (((_1841_v39).contains((_dafny.ZERO).minus((_1838_v37)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_1838_v37).length))]))) ? ((_1841_v39).get((_dafny.ZERO).minus((_1838_v37)[_module.__default.safeIndex(new BigNumber(935), new BigNumber((_1838_v37).length))]))) : (_dafny.Set.fromElements(new BigNumber((_module.__default.fm1(globalState)).length), (_this).f11)));
          let _pat_let_tv47 = _1832_v33;
          let _pat_let_tv48 = _1832_v33;
          let _pat_let_tv49 = _1805_v11;
          (globalState).f7 = !_dafny.areEqual(_dafny.Seq.of((_1805_v11).f17, (_module.D19.create_DC43((_1832_v33).f18, p0, (_1805_v11).f17, function (_pat_let42_0) {
  return function (_1842_dt__update__tmp_h1) {
    return function (_pat_let43_0) {
      return function (_1843_dt__update_hcf13_h0) {
        return _module.D6.create_DC10(_1843_dt__update_hcf13_h0);
      }(_pat_let43_0);
    }(_dafny.Seq.of((_pat_let_tv47).f18, (_pat_let_tv48).f18, (_pat_let_tv49).f17));
  }(_pat_let42_0);
}(_module.D6.create_DC10(_dafny.Seq.of((_1805_v11).f17))), _1827_v29)).dtor_cf69, _1791_v0, (_1831_v32)[_module.__default.safeIndex((((_1792_v1).contains((_1805_v11).f17)) ? ((_1792_v1).get((_1805_v11).f17)) : ((_this).f11)), new BigNumber((_1831_v32).length))]), _1831_v32);
          let _1844_v41;
          _1844_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1828_i3,new BigNumber((function () {
            let _coll68 = new _dafny.Set();
            for (const _compr_68 of (_dafny.Seq.of(p0, p0)).Elements) {
              let _1845_v40 = _compr_68;
              if (_dafny.Seq.contains(_dafny.Seq.of(p0, p0), _1845_v40)) {
                _coll68.add(_1845_v40);
              }
            }
            return _coll68;
          }()).length));
          let _1846_v42;
          let _nw297 = new _module.C3();
          _nw297.__ctor((_this).f11, _1828_i3, _1828_i3, _dafny.Set.fromElements((_1832_v33).f19));
          _1846_v42 = _nw297;
          let _1847_v43;
          _1847_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1844_v41,_1846_v42);
          let _1848_v44;
          _1848_v44 = _dafny.Set.fromElements((((_1847_v43).contains(_1844_v41)) ? ((_1847_v43).get(_1844_v41)) : (_1846_v42)), _1846_v42, _1846_v42);
          (globalState).f1 = (_1828_i3).minus((_dafny.ZERO).minus(new BigNumber((_1848_v44).length)));
        } else {
          let _1849_v45;
          _1849_v45 = _dafny.Set.fromElements((_1805_v11).f17, !((_1832_v33).fm14((_this).f14, globalState)));
          let _1850_v46;
          _1850_v46 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_1828_i3),_1849_v45);
          let _1851_v47;
          _1851_v47 = _module.D0.create_DC0((_this).f11);
          let _rhs323 = ((_this).f11).minus((_this).f11);
          let _rhs324 = _1791_v0;
          let _rhs325 = _1828_i3;
          let _rhs326 = (((_1832_v33).f19) ? ((((_1850_v46).contains(_1851_v47)) ? ((_1850_v46).get(_1851_v47)) : (_module.__default.fm41(globalState)))) : (_1849_v45));
          let _rhs327 = _module.__default.safeModuloInt((_this).f11, new BigNumber((_dafny.Seq.Concat(_1827_v29, _module.__default.fm1(globalState))).length));
          let _lhs272 = globalState;
          let _lhs273 = globalState;
          let _lhs274 = globalState;
          let _lhs275 = globalState;
          _lhs272.f1 = _rhs323;
          _lhs273.f7 = _rhs324;
          _lhs274.f1 = _rhs325;
          _1849_v45 = _rhs326;
          _lhs275.f1 = _rhs327;
          (globalState).f1 = (_this).f11;
          let _1852_v48;
          _1852_v48 = _dafny.Set.fromElements(p0, _module.__default.fm20(_1831_v32, (_1805_v11).fm3(_1828_i3, (_1805_v11).f17, (_1832_v33).f19, globalState), (_1805_v11).f17, globalState), p0, p0);
          (globalState).f1 = ((_this).f11).plus(((_this).fm6((_1832_v33).f18, (_this).f11, globalState)).multipliedBy(new BigNumber((_1852_v48).length)));
          (globalState).f7 = ((_this).f11).isLessThan((_dafny.ZERO).minus((_this).f11));
          let _rhs328 = _module.__default.safeDivisionInt(((_this).f11).minus((_dafny.ZERO).minus((_this).f11)), new BigNumber(-154));
          let _rhs329 = _module.__default.fm0(globalState);
          let _rhs330 = (_1832_v33).fm14((_this).f14, globalState);
          let _rhs331 = _dafny.Seq.Concat(_1827_v29, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("stifwiti"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(265)), ((_1853_p0) => function (_1854_i6) {
            return _1853_p0;
          })(p0))), _module.__default.safeIndex(_1828_i3, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("stifwiti"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(265)), ((_1855_p0) => function (_1856_i6) {
            return _1855_p0;
          })(p0)))).length)), p0));
          let _lhs276 = globalState;
          let _lhs277 = globalState;
          let _lhs278 = globalState;
          _lhs276.f1 = _rhs328;
          _lhs277.f6 = _rhs329;
          _lhs278.f6 = _rhs330;
          _1827_v29 = _rhs331;
        }
        let _1857_v49;
        let _nw298 = Array((new BigNumber(26)).toNumber()).fill([]);
        _1857_v49 = _nw298;
        let _index256 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_1857_v49).length));
        (_1857_v49)[_index256] = _1829_v30;
        let _index257 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_1857_v49).length));
        (_1857_v49)[_index257] = _1829_v30;
      }
      let _1858_v50;
      let _nw299 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _1858_v50 = _nw299;
      _1858_v50 = _1858_v50;
      r0 = _module.__default.fm37(globalState);
      return r0;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.Map.Empty;
      let _1859_v0;
      _1859_v0 = false;
      let _1860_v1;
      _1860_v1 = new _dafny.CodePoint('c'.codePointAt(0));
      let _1861_v2;
      _1861_v2 = _dafny.Set.fromElements(_1860_v1, _1860_v1, _1860_v1);
      let _1862_v3;
      _1862_v3 = _dafny.MultiSet.fromElements((_this).f12, (_this).f12);
      let _1863_i0;
      _1863_i0 = _dafny.ZERO;
      L16: {
        while ((_module.__default.fm53((_this).f11, _1859_v0, (_dafny.ZERO).minus(new BigNumber((_1861_v2).length)), globalState)).equals(_1862_v3)) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1863_i0)) {
              break L16;
            }
            _1863_i0 = (_1863_i0).plus(_dafny.ONE);
            let _1864_v4;
            let _nw300 = Array((new BigNumber(18)).toNumber()).fill(false);
            _1864_v4 = _nw300;
            let _index258 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_1864_v4).length));
            (_1864_v4)[_index258] = _1859_v0;
            let _index259 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_1864_v4).length));
            (_1864_v4)[_index259] = _1859_v0;
            let _1865_v5;
            _1865_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(619),_1859_v0);
            let _1866_v6;
            _1866_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f11);
            (globalState).f7 = !(_dafny.Map.Empty.slice().updateUnsafe((_this).f11,new BigNumber((_1865_v5).length))).equals(((_1866_v6).update(p0, new BigNumber(-967))).update((_this).f11, p0));
            let _1867_v7;
            _1867_v7 = _module.D0.create_DC0(p0);
            let _1868_v8;
            _1868_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1859_v0,(_1864_v4)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_1864_v4).length))]);
            let _1869_v9;
            _1869_v9 = _module.D16.create_DC39(p0, _1859_v0, _1868_v8, new _dafny.CodePoint('m'.codePointAt(0)));
            let _1870_v10;
            _1870_v10 = _dafny.Seq.of((_1869_v9).dtor_cf63);
            let _index260 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_1864_v4).length));
            (_1864_v4)[_index260] = (_this).fm5(_1867_v7, _dafny.Seq.Concat(_1870_v10, _1870_v10), p0, _1860_v1, globalState);
            if ((_1864_v4)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_1864_v4).length))]) {
              (globalState).f1 = p0;
              let _1871_v11;
              let _nw301 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
              _1871_v11 = _nw301;
              let _nw302 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
              _1871_v11 = _nw302;
              (globalState).f1 = new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("gh"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("gh")).length)), ((_1859_v0) ? (_1860_v1) : (_1860_v1)))).length);
              _1865_v5 = (_1865_v5).update((_this).f11, true);
              let _1872_v12;
              _1872_v12 = _dafny.MultiSet.fromElements(_1859_v0, (_1864_v4)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_1864_v4).length))]);
              (globalState).f7 = (new BigNumber((_1872_v12).cardinality())).isLessThan(new BigNumber(-682));
            } else {
              (globalState).f1 = new BigNumber((_dafny.Seq.UnicodeFromString("psoaf")).length);
              let _1873_v13;
              _1873_v13 = _dafny.MultiSet.fromElements(p0, new BigNumber(331));
              let _1874_v14;
              _1874_v14 = _dafny.Seq.UnicodeFromString("acp");
              let _1875_v15;
              let _nw303 = new _module.C4();
              _nw303.__ctor(_1873_v13, !_dafny.areEqual(_1874_v14, _1874_v14), (_this).f12, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_1860_v1)).length));
              _1875_v15 = _nw303;
              (globalState).f6 = ((new BigNumber(-207)).plus(p0)).isLessThanOrEqualTo(p0);
              (globalState).f1 = (_dafny.ZERO).minus(p0);
              let _1876_v16;
              let _nw304 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
              _1876_v16 = _nw304;
              let _index261 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1876_v16).length));
              (_1876_v16)[_index261] = (_dafny.ZERO).minus(p0);
              let _index262 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1876_v16).length));
              (_1876_v16)[_index262] = (p0).minus(new BigNumber((_dafny.Seq.of((_this).f11)).length));
            }
          }
        }
      }
      let _1877_v17;
      _1877_v17 = _dafny.MultiSet.fromElements(p0);
      let _1878_v18;
      let _nw305 = new _module.C4();
      _nw305.__ctor(_1877_v17, _1859_v0, _dafny.Set.fromElements(_1859_v0), new BigNumber((_dafny.Seq.of(_1859_v0, false, _1859_v0, _1859_v0, _1859_v0)).length));
      _1878_v18 = _nw305;
      let _1879_v19;
      _1879_v19 = _dafny.Set.fromElements(p0, new BigNumber((_dafny.Set.fromElements(_1878_v18)).length));
      let _1880_v20;
      _1880_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1879_v19,new BigNumber((_dafny.Seq.UnicodeFromString("eo")).length));
      let _1881_v21;
      _1881_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1880_v20).length),_module.__default.fm1(globalState));
      let _1882_v22;
      _1882_v22 = _dafny.Seq.of(_1860_v1);
      let _1883_v23;
      _1883_v23 = _module.D16.create_DC39((_this).f11, _1859_v0, _dafny.Map.Empty.slice().updateUnsafe(_1859_v0,false), _1860_v1);
      let _1884_v24;
      _1884_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1859_v0,_1859_v0);
      let _1885_v25;
      _1885_v25 = _dafny.Seq.of(_1884_v24, _1884_v24, _1884_v24, _dafny.Map.Empty.slice().updateUnsafe(true,_1859_v0));
      let _1886_v26;
      _1886_v26 = _dafny.Seq.of(_1859_v0);
      let _1887_v27;
      _1887_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1859_v0,new BigNumber(555));
      let _1888_v28;
      let _nw306 = Array((new BigNumber(19)).toNumber());
      _nw306[(_dafny.ZERO).toNumber()] = !(_1859_v0);
      _nw306[(_dafny.ONE).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(2)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(3)).toNumber()] = ((_this).f11).isLessThanOrEqualTo((_this).f11);
      _nw306[(new BigNumber(4)).toNumber()] = !(_1881_v21).contains((_this).f11);
      _nw306[(new BigNumber(5)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(6)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(7)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1882_v22, _dafny.Seq.of((_1883_v23).dtor_cf65));
      _nw306[(new BigNumber(8)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(9)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(10)).toNumber()] = ((_1878_v18).f11).isLessThan(new BigNumber((_1885_v25).length));
      _nw306[(new BigNumber(11)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(12)).toNumber()] = (((_1884_v24).contains(_1859_v0)) ? ((_1884_v24).get(_1859_v0)) : (_1859_v0));
      _nw306[(new BigNumber(13)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(14)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(15)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(16)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(17)).toNumber()] = _1859_v0;
      _nw306[(new BigNumber(18)).toNumber()] = (_1886_v26)[_module.__default.safeIndex((_module.D11.create_DC23(true, (_1878_v18).f11, new BigNumber((_1887_v27).length))).dtor_cf32, new BigNumber((_1886_v26).length))];
      _1888_v28 = _nw306;
      let _index263 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_1888_v28).length));
      (_1888_v28)[_index263] = _1859_v0;
      let _index264 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_1888_v28).length));
      (_1888_v28)[_index264] = _1859_v0;
      let _1889_v29;
      let _nw307 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _1889_v29 = _nw307;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1889_v29).length))) {
        let _1890_i1 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1890_i1)) && ((_1890_i1).isLessThan(new BigNumber((_1889_v29).length))))) {
          (_1889_v29)[(_1890_i1)] = _module.__default.safeDivisionInt(_1890_i1, (_dafny.ZERO).minus(p0));
        }
      }
      let _1891_v30;
      _1891_v30 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p0, p0),(_1888_v28)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1888_v28).length))]);
      let _1892_v31;
      _1892_v31 = _dafny.Seq.of((_this).f11);
      _1891_v30 = (_1891_v30).update(((_1878_v18).f11).plus(new BigNumber((_1892_v31).length)), _1859_v0);
      let _1893_v32;
      _1893_v32 = _dafny.MultiSet.fromElements(_1859_v0, ((_1878_v18).f11).isLessThan((_1878_v18).f11));
      _1893_v32 = _1893_v32;
      (globalState).f1 = (_this).f11;
      let _1894_v33;
      let _nw308 = Array((new BigNumber(26)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1894_v33 = _nw308;
      r0 = _1894_v33;
      let _1895_v34;
      _1895_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1882_v22,_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-102)), ((_1896_v18) => function (_1897_i2) {
        return (_1896_v18).f11;
      })(_1878_v18))));
      r1 = (_1895_v34).Merge(_1895_v34);
      return [r0, r1];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f12 = _dafny.Set.Empty;
      this._f11 = _dafny.ZERO;
      this._f13 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f13, f12, f11) {
      let _this = this;
      (_this)._f13 = f13;
      (_this)._f12 = f12;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((_this).f13) {
        return (_this).f13;
      } else {
        return true;
      }
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return (_this).f11;
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(((_dafny.MultiSet.fromElements((_this).f11, (_this).f11)).Difference((_dafny.MultiSet.fromElements((_this).f11)).Union(_dafny.MultiSet.fromElements((_this).f11)))).cardinality());
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("ypsi"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(3)), function (_1898_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }));
    };
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f11;
    };
    fm8(p0, globalState) {
      let _this = this;
      return new BigNumber(-6);
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      if (((_this).f13) && ((_this).f13)) {
        let _1899_v0;
        _1899_v0 = _module.D0.create_DC0((_this).f11);
        let _1900_v1;
        let _nw309 = Array((new BigNumber(17)).toNumber());
        _nw309[(_dafny.ZERO).toNumber()] = (_this).f13;
        _nw309[(_dafny.ONE).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(2)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(3)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(4)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(5)).toNumber()] = (_this).fm5(_1899_v0, _dafny.Seq.of((_this).f13, (_this).f13), (_this).f11, p0, globalState);
        _nw309[(new BigNumber(6)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(7)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(8)).toNumber()] = false;
        _nw309[(new BigNumber(9)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(10)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(11)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(12)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(13)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(14)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(15)).toNumber()] = (_this).f13;
        _nw309[(new BigNumber(16)).toNumber()] = (_this).f13;
        _1900_v1 = _nw309;
        let _1901_v2;
        _1901_v2 = _dafny.Seq.of(_1900_v1, _1900_v1);
        let _1902_v3;
        _1902_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_1901_v2);
        let _1903_v4;
        _1903_v4 = _dafny.Seq.of(_1901_v2);
        let _1904_v5;
        _1904_v5 = _dafny.Seq.of((_this).f13, !((_this).f13), (_this).f13);
        let _1905_v6;
        let _nw310 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1905_v6 = _nw310;
        let _1906_v7;
        _1906_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1905_v6,_1901_v2);
        let _1907_v8;
        let _nw311 = Array((new BigNumber(27)).toNumber());
        _nw311[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1901_v2, _1901_v2);
        _nw311[(_dafny.ONE).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(2)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(3)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(4)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(5)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(6)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(7)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(8)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(9)).toNumber()] = _dafny.Seq.update((((_1902_v3).contains((_this).f13)) ? ((_1902_v3).get((_this).f13)) : (_dafny.Seq.update((_1903_v4)[_module.__default.safeIndex(new BigNumber((_1904_v5).length), new BigNumber((_1903_v4).length))], _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_this).f13)).length), new BigNumber(((_1903_v4)[_module.__default.safeIndex(new BigNumber((_1904_v5).length), new BigNumber((_1903_v4).length))]).length)), _1900_v1))), _module.__default.safeIndex((_this).f11, new BigNumber(((((_1902_v3).contains((_this).f13)) ? ((_1902_v3).get((_this).f13)) : (_dafny.Seq.update((_1903_v4)[_module.__default.safeIndex(new BigNumber((_1904_v5).length), new BigNumber((_1903_v4).length))], _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_this).f13)).length), new BigNumber(((_1903_v4)[_module.__default.safeIndex(new BigNumber((_1904_v5).length), new BigNumber((_1903_v4).length))]).length)), _1900_v1)))).length)), (_1901_v2)[_module.__default.safeIndex((_this).f11, new BigNumber((_1901_v2).length))]);
        _nw311[(new BigNumber(10)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(11)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_1901_v2, _module.__default.safeIndex((_this).f11, new BigNumber((_1901_v2).length)), _1900_v1);
        _nw311[(new BigNumber(13)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(14)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(15)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(16)).toNumber()] = (((_1906_v7).contains(_1905_v6)) ? ((_1906_v7).get(_1905_v6)) : (_1901_v2));
        _nw311[(new BigNumber(17)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(18)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(19)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(20)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(21)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_1901_v2, _1901_v2);
        _nw311[(new BigNumber(23)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(24)).toNumber()] = _1901_v2;
        _nw311[(new BigNumber(25)).toNumber()] = _dafny.Seq.of(_1900_v1);
        _nw311[(new BigNumber(26)).toNumber()] = _dafny.Seq.Concat(_1901_v2, _1901_v2);
        _1907_v8 = _nw311;
        let _index265 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1907_v8).length));
        (_1907_v8)[_index265] = _1901_v2;
        let _index266 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1907_v8).length));
        (_1907_v8)[_index266] = _1901_v2;
        let _1908_v9;
        _1908_v9 = _dafny.Seq.UnicodeFromString("doukbfk");
        _1908_v9 = _1908_v9;
        let _1909_v10;
        _1909_v10 = new _dafny.CodePoint('m'.codePointAt(0));
        _1909_v10 = _1909_v10;
        (globalState).f1 = (_this).f11;
        let _1910_v11;
        _1910_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f11);
        let _1911_v12;
        _1911_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1910_v11,_dafny.Set.fromElements(_module.__default.fm0(globalState)));
        let _1912_v13;
        let _nw312 = new _module.C2();
        _nw312.__ctor(_1909_v10, new BigNumber(((_1911_v12).Merge(_1911_v12)).length));
        _1912_v13 = _nw312;
      } else {
        let _1913_v14;
        _1913_v14 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f11));
        let _1914_v15;
        let _nw313 = new _module.C4();
        _nw313.__ctor(_1913_v14, (_this).f13, (_this).f12, (_this).f11);
        _1914_v15 = _nw313;
        let _1915_v16;
        _1915_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_1914_v15);
        let _1916_v17;
        _1916_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1913_v14,_dafny.Map.Empty.slice().updateUnsafe((_this).f11,new BigNumber((_1915_v16).length)));
        let _1917_v18;
        _1917_v18 = _dafny.MultiSet.fromElements((_this).f13, false);
        let _1918_v19;
        _1918_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4((_this).f11, _1916_v17, (((_1917_v18).contains((_this).f13)) ? ((_1917_v18).get((_this).f13)) : ((_this).f11)), (_this).f13, globalState),(_1914_v15).f11);
        let _1919_v21;
        _1919_v21 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("iswspx")).length));
        let _1920_v22;
        let _nw314 = Array((new BigNumber(26)).toNumber());
        _nw314[(_dafny.ZERO).toNumber()] = (_this).f11;
        _nw314[(_dafny.ONE).toNumber()] = (_this).f11;
        _nw314[(new BigNumber(2)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(3)).toNumber()] = ((!((_this).f13)) ? (new BigNumber(912)) : ((_this).f11));
        _nw314[(new BigNumber(4)).toNumber()] = new BigNumber(331);
        _nw314[(new BigNumber(5)).toNumber()] = (_this).f11;
        _nw314[(new BigNumber(6)).toNumber()] = new BigNumber(((_this).f12).length);
        _nw314[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt((_1914_v15).f11, new BigNumber(382));
        _nw314[(new BigNumber(8)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(9)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(10)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(11)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(12)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(13)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(14)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(15)).toNumber()] = (new BigNumber(934)).minus((_this).f11);
        _nw314[(new BigNumber(16)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(17)).toNumber()] = new BigNumber((function () {
          let _coll69 = new _dafny.Map();
          for (const _compr_69 of (_1913_v14).Elements) {
            let _1921_v20 = _compr_69;
            if ((_1913_v14).contains(_1921_v20)) {
              _coll69.push([(_1921_v20).plus((_1914_v15).f11),new BigNumber(244)]);
            }
          }
          return _coll69;
        }()).length);
        _nw314[(new BigNumber(18)).toNumber()] = (_this).f11;
        _nw314[(new BigNumber(19)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(20)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(21)).toNumber()] = (_1914_v15).f11;
        _nw314[(new BigNumber(22)).toNumber()] = (_this).f11;
        _nw314[(new BigNumber(23)).toNumber()] = ((_1914_v15).f11).minus((_this).f11);
        _nw314[(new BigNumber(24)).toNumber()] = (_this).f11;
        _nw314[(new BigNumber(25)).toNumber()] = ((true) ? ((_1919_v21)[_module.__default.safeIndex((_1914_v15).f11, new BigNumber((_1919_v21).length))]) : (new BigNumber(723)));
        _1920_v22 = _nw314;
        let _index267 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_1920_v22).length));
        (_1920_v22)[_index267] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), ((_1922_p0) => function (_1923_i0) {
          return _1922_p0;
        })(p0))).length);
        let _1924_v23;
        _1924_v23 = _module.D2.create_DC4(new BigNumber(894), (_1914_v15).f11, (_this).f11);
        let _1925_v24;
        _1925_v24 = _dafny.Seq.UnicodeFromString("hlasr");
        let _1926_v25;
        let _nw315 = new _module.C6();
        _nw315.__ctor(_1924_v23, _dafny.Set.fromElements(!(false)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1925_v24, _dafny.Seq.update(_module.__default.fm1(globalState), _module.__default.safeIndex((_this).f11, new BigNumber((_module.__default.fm1(globalState)).length)), p0))).length)));
        _1926_v25 = _nw315;
        let _1927_v26;
        _1927_v26 = _dafny.Seq.of(!((_this).f13), (_this).f13);
        let _1928_v27;
        _1928_v27 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("twrkpvv"));
        let _index268 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_1920_v22).length));
        let _rhs332 = _module.__default.safeDivisionInt(new BigNumber((_1927_v26).length), new BigNumber((_1919_v21).length));
        let _rhs333 = (_1918_v19).Merge(_1918_v19);
        let _rhs334 = new BigNumber(((_dafny.MultiSet.fromElements(_1925_v24)).Intersect(_dafny.MultiSet.FromArray(_1928_v27))).cardinality());
        let _rhs335 = _1926_v25;
        let _rhs336 = _module.__default.safeDivisionInt((((_this).fm4((_1914_v15).f11, _1916_v17, (((_1913_v14).contains((_this).f11)) ? ((_1913_v14).get((_this).f11)) : (new BigNumber((_1927_v26).length))), (_this).f13, globalState)) ? ((_this).f11) : ((_1914_v15).f11)), (_this).f11);
        let _lhs279 = globalState;
        let _lhs280 = _1920_v22;
        let _lhs281 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_1920_v22).length));
        let _lhs282 = globalState;
        _lhs279.f1 = _rhs332;
        _1918_v19 = _rhs333;
        _lhs280[_lhs281] = _rhs334;
        _1926_v25 = _rhs335;
        _lhs282.f1 = _rhs336;
        let _1929_v28;
        let _1930_v29;
        let _out33;
        let _out34;
        let _outcollector11 = (_1926_v25).m1((_1919_v21)[_module.__default.safeIndex((_1920_v22)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_1920_v22).length))], new BigNumber((_1919_v21).length))], globalState);
        _out33 = _outcollector11[0];
        _out34 = _outcollector11[1];
        _1929_v28 = _out33;
        _1930_v29 = _out34;
        (globalState).f1 = (_this).f11;
        let _1931_v30;
        _1931_v30 = _dafny.Seq.of(_1919_v21, _1919_v21);
        (globalState).f1 = ((_this).f11).plus(new BigNumber(((((_this).f13) ? (_1919_v21) : ((_1931_v30)[_module.__default.safeIndex((_1914_v15).f11, new BigNumber((_1931_v30).length))]))).length));
        let _index269 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_1920_v22).length));
        (_1920_v22)[_index269] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1920_v22)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_1920_v22).length))]));
      }
      let _1932_v31;
      _1932_v31 = _dafny.Seq.of(new BigNumber(-537));
      let _1933_v32;
      _1933_v32 = _module.D0.create_DC1((_this).f13, (_this).f11, (_this).f13);
      let _1934_v33;
      _1934_v33 = _dafny.Seq.UnicodeFromString("liexbvcq");
      let _1935_v34;
      let _nw316 = Array((new BigNumber(28)).toNumber()).fill(false);
      _1935_v34 = _nw316;
      let _1936_v35;
      _1936_v35 = _dafny.Set.fromElements(_1935_v34, _1935_v34);
      let _1937_v36;
      _1937_v36 = _module.D21.create_DC47(_1936_v35);
      let _1938_v37;
      let _nw317 = Array((new BigNumber(21)).toNumber());
      _nw317[(_dafny.ZERO).toNumber()] = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of((_this).f11), _1932_v31));
      _nw317[(_dafny.ONE).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(2)).toNumber()] = (_1933_v32).dtor_cf3;
      _nw317[(new BigNumber(3)).toNumber()] = _dafny.Seq.contains(_dafny.Seq.update(_1934_v33, _module.__default.safeIndex((_this).f11, new BigNumber((_1934_v33).length)), new _dafny.CodePoint('q'.codePointAt(0))), p0);
      _nw317[(new BigNumber(4)).toNumber()] = !((_1937_v36).dtor_cf77).equals(_dafny.Set.fromElements(_1935_v34));
      _nw317[(new BigNumber(5)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(6)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(7)).toNumber()] = !((((_this).f13) ? ((_this).f13) : ((_this).f13)));
      _nw317[(new BigNumber(8)).toNumber()] = !((_this).f11).isEqualTo((_this).f11);
      _nw317[(new BigNumber(9)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(10)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(11)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(12)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(13)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(14)).toNumber()] = !((_this).f13) || ((_this).f13);
      _nw317[(new BigNumber(15)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(16)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(17)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(18)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(19)).toNumber()] = (_this).f13;
      _nw317[(new BigNumber(20)).toNumber()] = (_this).f13;
      _1938_v37 = _nw317;
      let _index270 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
      (_1938_v37)[_index270] = true;
      let _1939_v38;
      _1939_v38 = _module.D7.create_DC14(_module.D7.create_DC12(_1932_v31));
      let _1940_v39;
      _1940_v39 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(new BigNumber((_1932_v31).length), (_this).f11, globalState),!((_this).f13));
      let _1941_v41;
      _1941_v41 = _dafny.MultiSet.fromElements((_this).f11, (_this).f11, (_this).f11);
      let _1942_v42;
      _1942_v42 = _dafny.Set.fromElements(_1940_v39, (((_this).f13) ? (_1940_v39) : (_1940_v39)), _1940_v39, (function () {
        let _coll70 = new _dafny.Map();
        for (const _compr_70 of ((_1941_v41).update((_this).f11, _module.__default.abs((_dafny.ZERO).minus((_this).f11)))).Elements) {
          let _1943_v40 = _compr_70;
          if (((_1941_v41).update((_this).f11, _module.__default.abs((_dafny.ZERO).minus((_this).f11)))).contains(_1943_v40)) {
            _coll70.push([(_1943_v40).multipliedBy((_this).f11),!((_this).f13)]);
          }
        }
        return _coll70;
      }()).Merge(_1940_v39), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1934_v33).length),(_this).f13));
      let _1944_v44;
      _1944_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1941_v41,function () {
        let _coll71 = new _dafny.Map();
        for (const _compr_71 of _dafny.IntegerRange(new BigNumber(50), new BigNumber(-202))) {
          let _1945_v43 = _compr_71;
          if (((new BigNumber(50)).isLessThanOrEqualTo(_1945_v43)) && ((_1945_v43).isLessThan(new BigNumber(-202)))) {
            _coll71.push([_module.__default.safeDivisionInt(_1945_v43, (_this).f11),(_dafny.ZERO).minus((_this).f11)]);
          }
        }
        return _coll71;
      }());
      let _1946_v45;
      _1946_v45 = _dafny.Seq.of(_1932_v31);
      let _1947_v46;
      let _nw318 = new _module.C7();
      _nw318.__ctor(_1946_v45, (_this).f12, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f11),true)).length));
      _1947_v46 = _nw318;
      let _1948_v47;
      _1948_v47 = _dafny.Set.fromElements(_1947_v46, _1947_v46);
      let _1949_v48;
      _1949_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4((_this).f11, _1944_v44, new BigNumber((_1948_v47).length), (_this).f13, globalState),(_this).f11);
      let _1950_v49;
      _1950_v49 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).fm3(new BigNumber((_1949_v48).length), (_this).f13, (_this).f13, globalState)),new BigNumber(452));
      let _1951_v50;
      _1951_v50 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_1932_v31).length)),_1950_v49);
      let _1952_v52;
      _1952_v52 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll72 = new _dafny.Map();
        for (const _compr_72 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f13)).Keys.Elements) {
          let _1953_v51 = _compr_72;
          if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f13)).contains(_1953_v51)) {
            _coll72.push([_module.__default.safeDivisionInt(_1953_v51, (_this).f11),(_this).f13]);
          }
        }
        return _coll72;
      }(),(_this).f13);
      let _index271 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
      let _rhs337 = (_this).fm4(((_this).f11).multipliedBy((_this).f11), _1944_v44, (_this).f11, true, globalState);
      let _rhs338 = _1939_v38;
      let _rhs339 = function () {
        let _coll73 = new _dafny.Set();
        for (const _compr_73 of (_1952_v52).Keys.Elements) {
          let _1954_v53 = _compr_73;
          if ((_1952_v52).contains(_1954_v53)) {
            _coll73.add(_1954_v53);
          }
        }
        return _coll73;
      }();
      let _rhs340 = (_this).f11;
      let _lhs283 = _1938_v37;
      let _lhs284 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
      let _lhs285 = globalState;
      _lhs283[_lhs284] = _rhs337;
      _1939_v38 = _rhs338;
      _1942_v42 = _rhs339;
      _lhs285.f1 = _rhs340;
      let _1955_v54;
      _1955_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,new _dafny.CodePoint('h'.codePointAt(0)));
      let _1956_v55;
      _1956_v55 = _dafny.Seq.of((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], (_this).f13);
      let _1957_v56;
      let _nw319 = Array((new BigNumber(21)).toNumber());
      _nw319[(_dafny.ZERO).toNumber()] = _module.__default.fm1(globalState);
      _nw319[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("yur");
      _nw319[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(677)), ((_1958_p0) => function (_1959_i1) {
        return _1958_p0;
      })(p0));
      _nw319[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(527)), ((_1960_p0) => function (_1961_i2) {
        return _1960_p0;
      })(p0));
      _nw319[(new BigNumber(4)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(5)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("ivoblt"), _module.__default.safeIndex(new BigNumber(448), new BigNumber((_dafny.Seq.UnicodeFromString("ivoblt")).length)), (((_1955_v54).contains(new BigNumber((_1956_v55).length))) ? ((_1955_v54).get(new BigNumber((_1956_v55).length))) : (p0)));
      _nw319[(new BigNumber(7)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(8)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1934_v33, _dafny.Seq.Create(_module.__default.abs(new BigNumber(716)), ((_1962_p0) => function (_1963_i3) {
        return _1962_p0;
      })(p0)));
      _nw319[(new BigNumber(10)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(11)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(12)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(13)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(14)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_1934_v33, _1934_v33);
      _nw319[(new BigNumber(16)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("ugocbljk");
      _nw319[(new BigNumber(18)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(19)).toNumber()] = _1934_v33;
      _nw319[(new BigNumber(20)).toNumber()] = _1934_v33;
      _1957_v56 = _nw319;
      _1957_v56 = _1957_v56;
      if ((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]) {
        (globalState).f7 = !((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]) || ((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]);
        (globalState).f1 = (_this).f11;
        (globalState).f6 = (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))];
        (globalState).f1 = new BigNumber((_1956_v55).length);
        let _index272 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
        (_1938_v37)[_index272] = (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))];
      } else {
        let _1964_v57;
        _1964_v57 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1965_v58;
        let _nw320 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _1965_v58 = _nw320;
        let _1966_v59;
        _1966_v59 = _dafny.Set.fromElements(_1965_v58);
        let _rhs341 = (new BigNumber(163)).multipliedBy(new BigNumber((_1966_v59).length));
        let _rhs342 = (_1934_v33)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_1956_v55, _1956_v55)).length), new BigNumber((_1934_v33).length))];
        let _rhs343 = (_this).f11;
        let _lhs286 = globalState;
        let _lhs287 = globalState;
        _lhs286.f1 = _rhs341;
        _1964_v57 = _rhs342;
        _lhs287.f1 = _rhs343;
        let _1967_v60;
        _1967_v60 = _module.D7.create_DC12(_dafny.Seq.Concat(_1932_v31, _dafny.Seq.Create(_module.__default.abs(new BigNumber(335)), function (_1968_i4) {
  return (_this).f11;
})));
        let _source28 = _1967_v60;
        if (_source28.is_DC13) {
          let _1969___mcc_h0 = (_source28).cf18;
          let _1970___mcc_h1 = (_source28).cf19;
          let _1971___mcc_h2 = (_source28).cf20;
          let _1972_cf20 = _1971___mcc_h2;
          let _1973_cf19 = _1970___mcc_h1;
          let _1974_cf18 = _1969___mcc_h0;
          let _1975_v61;
          _1975_v61 = _module.D8.create_DC16((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], (_this).f11);
          let _1976_v62;
          let _nw321 = new _module.C3();
          _nw321.__ctor(_1972_cf20, new BigNumber(36), _1972_cf20, (_this).f12);
          _1976_v62 = _nw321;
          let _1977_v63;
          _1977_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1976_v62,_1976_v62.f23);
          let _1978_v64;
          _1978_v64 = _dafny.Set.fromElements(_1977_v63);
          let _1979_v65;
          _1979_v65 = _module.D2.create_DC4((_dafny.ZERO).minus((_1975_v61).dtor_cf24), new BigNumber((_1978_v64).length), _1976_v62.f23);
          let _1980_v66;
          let _nw322 = new _module.C6();
          _nw322.__ctor(_1979_v65, ((_this).f12).Intersect(_1974_cf18), (_this).f11);
          _1980_v66 = _nw322;
          let _1981_v67;
          _1981_v67 = _dafny.Seq.of(_1965_v58, _1965_v58, _1965_v58, _1965_v58);
          let _index273 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_1965_v58).length));
          (_1965_v58)[_index273] = new BigNumber((_1981_v67).length);
          let _index274 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_1965_v58).length));
          (_1965_v58)[_index274] = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_this).f11, _1972_cf20), (_dafny.ZERO).minus((_1976_v62).f22));
          let _index275 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_1965_v58).length));
          (_1965_v58)[_index275] = (_dafny.ZERO).minus((((_1975_v61).dtor_cf24).multipliedBy(new BigNumber(-219))).multipliedBy((_this).f11));
          let _1982_v68;
          _1982_v68 = _dafny.Set.fromElements((_1976_v62).f22);
          let _1983_v70;
          _1983_v70 = _dafny.Seq.of(_1982_v68, _1982_v68, _1982_v68);
          let _1984_v71;
          _1984_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1964_v57,(_1965_v58)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_1965_v58).length))]);
          let _1985_v72;
          _1985_v72 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_module.__default.fm54(_1949_v48, globalState), _1984_v71),_1982_v68);
          let _1986_v75;
          _1986_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1934_v33,_module.__default.fm18(new BigNumber(((_this).f12).length), globalState));
          let _1987_v76;
          let _nw323 = Array((new BigNumber(24)).toNumber());
          _nw323[(_dafny.ZERO).toNumber()] = _1982_v68;
          _nw323[(_dafny.ONE).toNumber()] = (function () {
            let _coll74 = new _dafny.Set();
            for (const _compr_74 of _dafny.IntegerRange(new BigNumber(212), new BigNumber(905))) {
              let _1988_v69 = _compr_74;
              if (((new BigNumber(212)).isLessThanOrEqualTo(_1988_v69)) && ((_1988_v69).isLessThan(new BigNumber(905)))) {
                _coll74.add(_module.__default.safeModuloInt(_1988_v69, (_this).f11));
              }
            }
            return _coll74;
          }()).Union(_1982_v68);
          _nw323[(new BigNumber(2)).toNumber()] = (_1982_v68).Union(_dafny.Set.fromElements(_1972_cf20));
          _nw323[(new BigNumber(3)).toNumber()] = _1982_v68;
          _nw323[(new BigNumber(4)).toNumber()] = (_1983_v70)[_module.__default.safeIndex(_1973_cf19, new BigNumber((_1983_v70).length))];
          _nw323[(new BigNumber(5)).toNumber()] = (_module.D22.create_DC51(_1982_v68)).dtor_cf86;
          _nw323[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(new BigNumber(49), _1972_cf20, (_1976_v62).f22);
          _nw323[(new BigNumber(7)).toNumber()] = _1982_v68;
          _nw323[(new BigNumber(8)).toNumber()] = _1982_v68;
          _nw323[(new BigNumber(9)).toNumber()] = _1982_v68;
          _nw323[(new BigNumber(10)).toNumber()] = (_1982_v68).Union(_1982_v68);
          _nw323[(new BigNumber(11)).toNumber()] = (_1982_v68).Intersect(_1982_v68);
          _nw323[(new BigNumber(12)).toNumber()] = _1982_v68;
          _nw323[(new BigNumber(13)).toNumber()] = (((_1985_v72).contains(_dafny.MultiSet.fromElements(function () {
            let _coll76 = new _dafny.Map();
            for (const _compr_76 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), ((_1993_p0) => function (_1994_i5) {
              return _1993_p0;
            })(p0))).Elements) {
              let _1995_v73 = _compr_76;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), ((_1996_p0) => function (_1994_i5) {
                return _1996_p0;
              })(p0)), _1995_v73)) {
                _coll76.push([_1995_v73,_1976_v62.f23]);
              }
            }
            return _coll76;
          }()))) ? ((_1985_v72).get(_dafny.MultiSet.fromElements(function () {
            let _coll75 = new _dafny.Map();
            for (const _compr_75 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), ((_1989_p0) => function (_1990_i5) {
              return _1989_p0;
            })(p0))).Elements) {
              let _1991_v73 = _compr_75;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), ((_1992_p0) => function (_1990_i5) {
                return _1992_p0;
              })(p0)), _1991_v73)) {
                _coll75.push([_1991_v73,_1976_v62.f23]);
              }
            }
            return _coll75;
          }()))) : (_1982_v68));
          _nw323[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_1934_v33).length), new BigNumber((_dafny.MultiSet.fromElements((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))])).cardinality()));
          _nw323[(new BigNumber(15)).toNumber()] = _module.__default.fm26(globalState);
          _nw323[(new BigNumber(16)).toNumber()] = _1982_v68;
          _nw323[(new BigNumber(17)).toNumber()] = (function () {
            let _coll77 = new _dafny.Set();
            for (const _compr_77 of _dafny.IntegerRange(new BigNumber(-328), new BigNumber(346))) {
              let _1997_v74 = _compr_77;
              if (((new BigNumber(-328)).isLessThanOrEqualTo(_1997_v74)) && ((_1997_v74).isLessThan(new BigNumber(346)))) {
                _coll77.add(_module.__default.safeDivisionInt(_1997_v74, new BigNumber((_dafny.Seq.UnicodeFromString("mak")).length)));
              }
            }
            return _coll77;
          }()).Intersect(_dafny.Set.fromElements(_1976_v62.f23, (_this).f11, (_1932_v31)[_module.__default.safeIndex((_1965_v58)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_1965_v58).length))], new BigNumber((_1932_v31).length))]));
          _nw323[(new BigNumber(18)).toNumber()] = _1982_v68;
          _nw323[(new BigNumber(19)).toNumber()] = (((_1986_v75).contains(_dafny.Seq.UnicodeFromString("hcfn"))) ? ((_1986_v75).get(_dafny.Seq.UnicodeFromString("hcfn"))) : (_1982_v68));
          _nw323[(new BigNumber(20)).toNumber()] = _1982_v68;
          _nw323[(new BigNumber(21)).toNumber()] = _1982_v68;
          _nw323[(new BigNumber(22)).toNumber()] = _module.__default.fm26(globalState);
          _nw323[(new BigNumber(23)).toNumber()] = _1982_v68;
          _1987_v76 = _nw323;
          _1987_v76 = _1987_v76;
        } else if (_source28.is_DC12) {
          let _1998___mcc_h3 = (_source28).cf17;
          let _1999_cf17 = _1998___mcc_h3;
          let _index276 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_1965_v58).length));
          (_1965_v58)[_index276] = ((_this).f11).plus((_this).f11);
          let _2000_v77;
          _2000_v77 = _dafny.Seq.of(_1940_v39, _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f13));
          let _index277 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_1965_v58).length));
          let _index278 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
          let _rhs344 = _dafny.Seq.Concat(_1956_v55, _1956_v55);
          let _rhs345 = (_this).fm3(new BigNumber(-429), ((_2000_v77)[_module.__default.safeIndex((_this).f11, new BigNumber((_2000_v77).length))]).contains((_this).f11), (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], globalState);
          let _rhs346 = _module.__default.safeModuloInt((_this).f11, (_this).f11);
          let _rhs347 = !_dafny.areEqual(_dafny.Seq.Concat(_1934_v33, _1934_v33), _1934_v33);
          let _rhs348 = (_this).f11;
          let _lhs288 = _1965_v58;
          let _lhs289 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_1965_v58).length));
          let _lhs290 = globalState;
          let _lhs291 = _1938_v37;
          let _lhs292 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
          let _lhs293 = globalState;
          _1956_v55 = _rhs344;
          _lhs288[_lhs289] = _rhs345;
          _lhs290.f1 = _rhs346;
          _lhs291[_lhs292] = _rhs347;
          _lhs293.f1 = _rhs348;
          _1934_v33 = _dafny.Seq.Concat(_1934_v33, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("okrvwip"), _1934_v33));
          (globalState).f1 = (_dafny.ZERO).minus((_this).f11);
          let _2001_v78;
          _2001_v78 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1956_v55, _1956_v55),(_1965_v58)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_1965_v58).length))]);
          let _2002_v79;
          _2002_v79 = _module.D7.create_DC13((_this).f12, new BigNumber((_1955_v54).length), new BigNumber(485));
          _2001_v78 = (_2001_v78).update(_module.__default.fm35(new BigNumber((_1934_v33).length), (_1965_v58)[_module.__default.safeIndex(new BigNumber(549), new BigNumber((_1965_v58).length))], _2002_v79, new BigNumber((_1941_v41).cardinality()), globalState), (_this).f11);
        } else {
          let _2003___mcc_h4 = (_source28).cf21;
          let _2004_cf21 = _2003___mcc_h4;
          _1934_v33 = _1934_v33;
          (globalState).f1 = (_this).f11;
          let _2005_v80;
          _2005_v80 = _module.D8.create_DC16((_this).f13, new BigNumber((_dafny.Set.fromElements((_this).f11)).length));
          let _2006_v81;
          _2006_v81 = _module.D8.create_DC17(_2005_v80);
          let _2007_v82;
          _2007_v82 = _module.D8.create_DC17(_2005_v80);
          let _2008_v83;
          _2008_v83 = _module.D8.create_DC17(_2006_v81);
          let _2009_v84;
          _2009_v84 = _dafny.Map.Empty.slice().updateUnsafe((((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]) ? (_2008_v83) : (_2008_v83)),new _dafny.CodePoint('p'.codePointAt(0)));
          _2009_v84 = (_2009_v84).update((((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]) ? (_2008_v83) : (_2008_v83)), p0);
          let _2010_v85;
          _2010_v85 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]));
          let _2011_v86;
          _2011_v86 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_module.__default.fm12((_this).f11, _dafny.Seq.UnicodeFromString("pwgdxagxs"), globalState), _1956_v55),_2010_v85);
          _2011_v86 = (_2011_v86).update(_dafny.Seq.Concat(_1956_v55, _dafny.Seq.of(true, (_this).f13)), _2010_v85);
        }
        let _2012_v87;
        _2012_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1934_v33,(_this).f11);
        let _2013_v88;
        let _nw324 = new _module.C4();
        _nw324.__ctor(_dafny.MultiSet.fromElements((_this).f11), (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], (_this).f12, (_this).f11);
        _2013_v88 = _nw324;
        let _2014_v89;
        _2014_v89 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2012_v87).length),_2013_v88);
        let _2015_v90;
        _2015_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1950_v49).length),_1955_v54);
        let _2016_v91;
        _2016_v91 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,p0);
        let _2017_v92;
        _2017_v92 = _module.D11.create_DC23(true, (_2013_v88).f11, new BigNumber((_2016_v91).length));
        let _2018_v94;
        _2018_v94 = _dafny.Set.fromElements(_1955_v54, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm46(globalState)).length),new _dafny.CodePoint('k'.codePointAt(0))), (_dafny.Map.Empty.slice().updateUnsafe((_2013_v88).f11,new _dafny.CodePoint('p'.codePointAt(0)))).update((_2017_v92).dtor_cf33, _1964_v57), function () {
          let _coll78 = new _dafny.Map();
          for (const _compr_78 of _dafny.IntegerRange(new BigNumber(301), new BigNumber(662))) {
            let _2019_v93 = _compr_78;
            if (((new BigNumber(301)).isLessThanOrEqualTo(_2019_v93)) && ((_2019_v93).isLessThan(new BigNumber(662)))) {
              _coll78.push([_module.__default.safeModuloInt(_2019_v93, (_this).f11),p0]);
            }
          }
          return _coll78;
        }());
        if (!(_2018_v94).contains(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(315),_1964_v57)).update(new BigNumber((_2014_v89).length), p0)).Merge((((_2015_v90).contains((_2013_v88).f11)) ? ((_2015_v90).get((_2013_v88).f11)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1956_v55).length),p0)))))) {
          let _index279 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
          (_1938_v37)[_index279] = false;
          let _2020_v95;
          _2020_v95 = _module.D0.create_DC0((_1932_v31)[_module.__default.safeIndex(new BigNumber((_1949_v48).length), new BigNumber((_1932_v31).length))]);
          let _pat_let_tv50 = _1967_v60;
          let _index280 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
          (_1938_v37)[_index280] = (_1947_v46).fm5(function (_pat_let44_0) {
            return function (_2023_dt__update__tmp_h1) {
              return function (_pat_let47_0) {
                return function (_2024_dt__update_hcf0_h1) {
                  return _module.D0.create_DC0(_2024_dt__update_hcf0_h1);
                }(_pat_let47_0);
              }(new BigNumber(((_pat_let_tv50).dtor_cf17).length));
            }(_pat_let44_0);
          }(function (_pat_let45_0) {
            return function (_2021_dt__update__tmp_h0) {
              return function (_pat_let46_0) {
                return function (_2022_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_2022_dt__update_hcf0_h0);
                }(_pat_let46_0);
              }(new BigNumber(128));
            }(_pat_let45_0);
          }(_2020_v95)), _1956_v55, _module.__default.safeDivisionInt((_this).f11, new BigNumber(830)), _1964_v57, globalState);
          (globalState).f1 = (_module.__default.safeDivisionInt((_this).f11, (_this).f11)).plus(new BigNumber(684));
          let _2025_v96;
          _2025_v96 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
          let _2026_v97;
          _2026_v97 = _module.D16.create_DC39(new BigNumber((_1950_v49).length), (_this).f13, _2025_v96, new _dafny.CodePoint('n'.codePointAt(0)));
          let _2027_v98;
          let _nw325 = new _module.C2();
          _nw325.__ctor((_2026_v97).dtor_cf65, (_2013_v88).f11);
          _2027_v98 = _nw325;
          let _2028_v99;
          _2028_v99 = _dafny.Map.Empty.slice().updateUnsafe(_2027_v98,(_2013_v88).f11);
          _2028_v99 = (_2028_v99).update(_2027_v98, ((_2013_v88).f11).minus((_2013_v88).f11));
          let _2029_v100;
          let _nw326 = Array((new BigNumber(12)).toNumber()).fill([]);
          _2029_v100 = _nw326;
          let _index281 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_2029_v100).length));
          (_2029_v100)[_index281] = _1935_v34;
          let _index282 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_2029_v100).length));
          let _rhs349 = _1935_v34;
          let _rhs350 = (_2013_v88).f11;
          let _lhs294 = _2029_v100;
          let _lhs295 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_2029_v100).length));
          let _lhs296 = globalState;
          _lhs294[_lhs295] = _rhs349;
          _lhs296.f1 = _rhs350;
        } else {
          (globalState).f6 = (_2017_v92).dtor_cf31;
          let _index283 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
          (_1938_v37)[_index283] = (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))];
          (globalState).f1 = (_2013_v88).f11;
          let _2030_v101;
          let _nw327 = new _module.C2();
          _nw327.__ctor(new _dafny.CodePoint('f'.codePointAt(0)), (_2013_v88).f11);
          _2030_v101 = _nw327;
          let _pat_let_tv51 = _1934_v33;
          (globalState).f1 = ((_2013_v88).f11).multipliedBy((function (_pat_let48_0) {
            return function (_2031_dt__update__tmp_h2) {
              return function (_pat_let49_0) {
                return function (_2032_dt__update_hcf19_h0) {
                  return _module.D7.create_DC13((_2031_dt__update__tmp_h2).dtor_cf18, _2032_dt__update_hcf19_h0, (_2031_dt__update__tmp_h2).dtor_cf20);
                }(_pat_let49_0);
              }(new BigNumber((_pat_let_tv51).length));
            }(_pat_let48_0);
          }(_module.D7.create_DC13((_this).f12, (_2013_v88).f11, (_2013_v88).f11))).dtor_cf20);
        }
        let _2033_v102;
        _2033_v102 = _module.D6.create_DC9(_1940_v39);
        let _pat_let_tv52 = _1940_v39;
        _2033_v102 = function (_pat_let50_0) {
          return function (_2034_dt__update__tmp_h3) {
            return function (_pat_let51_0) {
              return function (_2035_dt__update_hcf12_h0) {
                return _module.D6.create_DC9(_2035_dt__update_hcf12_h0);
              }(_pat_let51_0);
            }(_pat_let_tv52);
          }(_pat_let50_0);
        }(_2033_v102);
        (globalState).f6 = (!((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))])) || ((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]);
      }
      if (((((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]) ? ((_this).f11) : ((_this).f11))).isLessThanOrEqualTo((_this).f11)) {
        let _rhs351 = (_this).f13;
        let _rhs352 = (false) && ((_this).f13);
        let _rhs353 = ((((_this).f13) || (!((_this).f13))) ? ((_this).f11) : ((_this).f11));
        let _lhs297 = globalState;
        let _lhs298 = globalState;
        let _lhs299 = globalState;
        _lhs297.f7 = _rhs351;
        _lhs298.f6 = _rhs352;
        _lhs299.f1 = _rhs353;
        _1955_v54 = (_1955_v54).update(new BigNumber(569), p0);
        (globalState).f7 = ((_this).f11).isLessThan(((_this).f11).minus((_this).f11));
        let _index284 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
        (_1938_v37)[_index284] = (_this).f13;
        let _pat_let_tv53 = _1936_v35;
        let _2036_v103;
        _2036_v103 = _dafny.MultiSet.fromElements(_1937_v36, _1937_v36, _1937_v36, _1937_v36, function (_pat_let52_0) {
          return function (_2037_dt__update__tmp_h4) {
            return function (_pat_let53_0) {
              return function (_2038_dt__update_hcf77_h0) {
                return _module.D21.create_DC47(_2038_dt__update_hcf77_h0);
              }(_pat_let53_0);
            }(_pat_let_tv53);
          }(_pat_let52_0);
        }(_1937_v36));
        _2036_v103 = (_2036_v103).update(_1937_v36, _module.__default.abs((_this).f11));
      } else {
        (globalState).f1 = (_this).f11;
        (globalState).f1 = new BigNumber(112);
        let _2039_v104;
        _2039_v104 = _module.D2.create_DC4(new BigNumber(121), (_this).f11, new BigNumber(728));
        let _2040_v105;
        _2040_v105 = _dafny.Seq.of(_2039_v104);
        let _2041_v106;
        _2041_v106 = _dafny.MultiSet.fromElements((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], (_this).f13);
        (globalState).f1 = ((_this).f11).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2040_v105, _2040_v105), _module.__default.safeIndex(new BigNumber((_2041_v106).cardinality()), new BigNumber((_dafny.Seq.Concat(_2040_v105, _2040_v105)).length)), _module.D2.create_DC4(new BigNumber((_1932_v31).length), (_this).f11, (_this).f11))).length));
        let _2042_v107;
        _2042_v107 = _dafny.Seq.of(_1956_v55);
        let _2043_v108;
        let _nw328 = new _module.C0();
        _nw328.__ctor(_dafny.Seq.Concat((_2042_v107)[_module.__default.safeIndex((_this).f11, new BigNumber((_2042_v107).length))], _1956_v55), (_this).fm7((_this).f11, (_this).f13, _dafny.MultiSet.FromArray(_1956_v55), globalState));
        _2043_v108 = _nw328;
        _2043_v108 = _2043_v108;
        (globalState).f1 = new BigNumber(901);
      }
      let _2044_v109;
      _2044_v109 = _module.D16.create_DC38(_1950_v49);
      let _source29 = function (_source30) {
        if (_source30.is_DC39) {
          let _2045___mcc_h17 = (_source30).cf62;
          let _2046___mcc_h18 = (_source30).cf63;
          let _2047___mcc_h19 = (_source30).cf64;
          let _2048___mcc_h20 = (_source30).cf65;
          let _2049_cf65 = _2048___mcc_h20;
          let _2050_cf64 = _2047___mcc_h19;
          let _2051_cf63 = _2046___mcc_h18;
          let _2052_cf62 = _2045___mcc_h17;
          return _module.D15.create_DC34();
        } else {
          let _2053___mcc_h21 = (_source30).cf61;
          let _2054_cf61 = _2053___mcc_h21;
          return _module.D15.create_DC34();
        }
      }(_2044_v109);
      if (_source29.is_DC34) {
        (globalState).f1 = ((false) ? ((_this).f11) : (((_module.__default.fm0(globalState)) ? ((_dafny.ZERO).minus((_this).f11)) : ((_this).f11))));
        let _2055_v110;
        let _nw329 = new _module.C2();
        _nw329.__ctor(_module.__default.fm20(_dafny.Seq.of((_this).f13, !((_this).f13), (_this).f13, (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]), new BigNumber(311), (_this).f13, globalState), new BigNumber(-613));
        _2055_v110 = _nw329;
        let _2056_v111;
        _2056_v111 = _dafny.Map.Empty.slice().updateUnsafe(_2055_v110,(_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]);
        let _2057_v112;
        let _nw330 = new _module.C2();
        _nw330.__ctor(p0, (_this).f11);
        _2057_v112 = _nw330;
        let _2058_v113;
        _2058_v113 = _dafny.Set.fromElements(_2057_v112, _2057_v112);
        let _rhs354 = (_2056_v111).Merge(_2056_v111);
        let _rhs355 = !((_1956_v55)[_module.__default.safeIndex(new BigNumber(((_2058_v113).Union(_2058_v113)).length), new BigNumber((_1956_v55).length))]);
        let _rhs356 = (_this).f11;
        let _rhs357 = _module.__default.safeDivisionInt((((_this).f13) ? ((_dafny.ZERO).minus((_this).f11)) : (new BigNumber(25))), (_this).f11);
        let _lhs300 = globalState;
        let _lhs301 = globalState;
        let _lhs302 = globalState;
        _2056_v111 = _rhs354;
        _lhs300.f6 = _rhs355;
        _lhs301.f1 = _rhs356;
        _lhs302.f1 = _rhs357;
        let _2059_v114;
        _2059_v114 = _dafny.Map.Empty.slice().updateUnsafe(_2044_v109,(_2055_v110).f11);
        _2059_v114 = (_2059_v114).update(_2044_v109, new BigNumber(-321));
        let _2060_v115;
        let _init39 = ((_2061_v110, _2062_v37) => function (_2063_i6) {
          return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_2061_v110).f11),_dafny.Map.Empty.slice().updateUnsafe((_2062_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_2062_v37).length))],(_2062_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_2062_v37).length))]));
        })(_2055_v110, _1938_v37);
        let _nw331 = Array((new BigNumber(28)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw331.length); _i0_39++) {
          _nw331[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _2060_v115 = _nw331;
        let _2064_v116;
        _2064_v116 = _dafny.Set.fromElements((_2055_v110).f11, (_2055_v110).f11);
        let _2065_v117;
        _2065_v117 = _dafny.Map.Empty.slice().updateUnsafe(true,true);
        let _2066_v118;
        _2066_v118 = _dafny.Map.Empty.slice().updateUnsafe(_2064_v116,_2065_v117);
        let _index285 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2060_v115).length));
        (_2060_v115)[_index285] = _2066_v118;
        let _2067_v119;
        _2067_v119 = _dafny.Seq.of(_1956_v55, _1956_v55, _1956_v55);
        let _2068_v120;
        _2068_v120 = _module.D14.create_DC31(false, (_this).f11, new BigNumber(397));
        let _index286 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2060_v115).length));
        let _rhs358 = ((_module.__default.fm55((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], new BigNumber((_1934_v33).length), (_this).f13, globalState)).Merge(_2066_v118)).update(_2064_v116, _2065_v117);
        let _rhs359 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat((_2067_v119)[_module.__default.safeIndex((_this).f11, new BigNumber((_2067_v119).length))], _dafny.Seq.of((_2068_v120).dtor_cf45, (_this).f13)), _1956_v55);
        let _lhs303 = _2060_v115;
        let _lhs304 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_2060_v115).length));
        let _lhs305 = globalState;
        _lhs303[_lhs304] = _rhs358;
        _lhs305.f6 = _rhs359;
      } else if (_source29.is_DC35) {
        let _2069___mcc_h5 = (_source29).cf50;
        let _2070___mcc_h6 = (_source29).cf51;
        let _2071___mcc_h7 = (_source29).cf52;
        let _2072___mcc_h8 = (_source29).cf53;
        let _2073___mcc_h9 = (_source29).cf54;
        let _2074_cf54 = _2073___mcc_h9;
        let _2075_cf53 = _2072___mcc_h8;
        let _2076_cf52 = _2071___mcc_h7;
        let _2077_cf51 = _2070___mcc_h6;
        let _2078_cf50 = _2069___mcc_h5;
        let _2079_v121;
        let _nw332 = new _module.C7();
        _nw332.__ctor((_1947_v46).f14, ((_2077_cf51).f12).Difference(_dafny.Set.fromElements(_2078_cf50, (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], false)), (_2077_cf51).f11);
        _2079_v121 = _nw332;
        let _2080_v122;
        _2080_v122 = _module.D2.create_DC4((_this).f11, (_this).f11, (_this).f11);
        let _source31 = _2080_v122;
        if (_source31.is_DC4) {
          let _2081___mcc_h22 = (_source31).cf6;
          let _2082___mcc_h23 = (_source31).cf7;
          let _2083___mcc_h24 = (_source31).cf8;
          let _2084_cf8 = _2083___mcc_h24;
          let _2085_cf7 = _2082___mcc_h23;
          let _2086_cf6 = _2081___mcc_h22;
          let _2087_v124;
          _2087_v124 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_module.__default.fm27((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], globalState), _1932_v31)).length),function () {
            let _coll79 = new _dafny.Map();
            for (const _compr_79 of (_1940_v39).Keys.Elements) {
              let _2088_v123 = _compr_79;
              if ((_1940_v39).contains(_2088_v123)) {
                _coll79.push([(_2088_v123).plus(_2084_cf8),(_this).f11]);
              }
            }
            return _coll79;
          }());
          _2087_v124 = (_2087_v124).update(_2085_cf7, _1950_v49);
          let _2089_v125;
          let _nw333 = new _module.C0();
          _nw333.__ctor(_1956_v55, _2074_cf54);
          _2089_v125 = _nw333;
          (globalState).f6 = (_2078_cf50) && (_2078_cf50);
          let _2090_v126;
          _2090_v126 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2076_cf52);
          let _2091_v128;
          _2091_v128 = _dafny.Set.fromElements(_2090_v126, function () {
            let _coll80 = new _dafny.Map();
            for (const _compr_80 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(203)), function (_2092_i7) {
              return new _dafny.CodePoint('w'.codePointAt(0));
            })).Elements) {
              let _2093_v127 = _compr_80;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(203)), function (_2092_i7) {
                return new _dafny.CodePoint('w'.codePointAt(0));
              }), _2093_v127)) {
                _coll80.push([_2093_v127,(_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]]);
              }
            }
            return _coll80;
          }(), _2090_v126);
          let _2094_v129;
          let _nw334 = new _module.C6();
          _nw334.__ctor(_2080_v122, (_2077_cf51).f12, _module.__default.safeModuloInt(_2086_cf6, new BigNumber((_2091_v128).length)));
          _2094_v129 = _nw334;
        } else if (_source31.is_DC5) {
          let _2095_v130;
          let _nw335 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _2095_v130 = _nw335;
          let _index287 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_2095_v130).length));
          (_2095_v130)[_index287] = _module.__default.safeDivisionInt(new BigNumber((_1934_v33).length), new BigNumber(483));
          let _index288 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_2095_v130).length));
          (_2095_v130)[_index288] = (_dafny.ZERO).minus((_2077_cf51).f11);
          let _index289 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_2095_v130).length));
          let _index290 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_2095_v130).length));
          let _rhs360 = (_1932_v31)[_module.__default.safeIndex(_2074_cf54, new BigNumber((_1932_v31).length))];
          let _rhs361 = _module.__default.safeDivisionInt(((_this).f11).plus((_this).f11), ((_2077_cf51).f11).multipliedBy(new BigNumber((_1934_v33).length)));
          let _rhs362 = new BigNumber(518);
          let _lhs306 = _2095_v130;
          let _lhs307 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_2095_v130).length));
          let _lhs308 = _2095_v130;
          let _lhs309 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_2095_v130).length));
          _lhs306[_lhs307] = _rhs360;
          _lhs308[_lhs309] = _rhs361;
          _2074_cf54 = _rhs362;
          (globalState).f1 = (_2095_v130)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_2095_v130).length))];
          let _2096_v131;
          _2096_v131 = _dafny.Seq.of(_2095_v130, _2095_v130);
          let _2097_v132;
          let _nw336 = Array((new BigNumber(23)).toNumber());
          _nw336[(_dafny.ZERO).toNumber()] = _2095_v130;
          _nw336[(_dafny.ONE).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(2)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(3)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(4)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(5)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(6)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(7)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(8)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(9)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(10)).toNumber()] = (_2096_v131)[_module.__default.safeIndex((_this).f11, new BigNumber((_2096_v131).length))];
          _nw336[(new BigNumber(11)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(12)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(13)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(14)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(15)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(16)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(17)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(18)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(19)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(20)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(21)).toNumber()] = _2095_v130;
          _nw336[(new BigNumber(22)).toNumber()] = _2095_v130;
          _2097_v132 = _nw336;
          let _2098_v133;
          _2098_v133 = _2097_v132;
          _2098_v133 = _2098_v133;
          let _2099_v134;
          let _nw337 = new _module.C6();
          _nw337.__ctor(_module.D2.create_DC4((_this).fm8(_module.__default.fm39((_2095_v130)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_2095_v130).length))], (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], (_2077_cf51).f11, globalState), globalState), (_2095_v130)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_2095_v130).length))], (_this).f11), (_this).f12, new BigNumber((_1934_v33).length));
          _2099_v134 = _nw337;
        } else {
          let _2100___mcc_h25 = (_source31).cf5;
          let _2101_cf5 = _2100___mcc_h25;
          let _2102_v135;
          let _nw338 = Array((new BigNumber(5)).toNumber());
          _nw338[(_dafny.ZERO).toNumber()] = ((_2077_cf51).f11).plus((_this).f11);
          _nw338[(_dafny.ONE).toNumber()] = new BigNumber(289);
          _nw338[(new BigNumber(2)).toNumber()] = (_2077_cf51).f11;
          _nw338[(new BigNumber(3)).toNumber()] = new BigNumber(-902);
          _nw338[(new BigNumber(4)).toNumber()] = (_this).f11;
          _2102_v135 = _nw338;
          let _index291 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2102_v135).length));
          (_2102_v135)[_index291] = (_this).f11;
          let _index292 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2102_v135).length));
          let _rhs363 = _1938_v37;
          let _rhs364 = new BigNumber(976);
          let _rhs365 = (_2077_cf51).f11;
          let _lhs310 = _2102_v135;
          let _lhs311 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2102_v135).length));
          let _lhs312 = globalState;
          _1938_v37 = _rhs363;
          _lhs310[_lhs311] = _rhs364;
          _lhs312.f1 = _rhs365;
          let _2103_v136;
          _2103_v136 = _dafny.MultiSet.fromElements((_this).f13, _2078_cf50);
          (globalState).f1 = ((!(_2103_v136).contains(_2078_cf50)) ? (_module.__default.safeModuloInt((_this).f11, new BigNumber((_dafny.Seq.UnicodeFromString("vxxic")).length))) : ((_this).f11));
          _2076_cf52 = (_2074_cf54).isLessThanOrEqualTo((_this).f11);
          let _2104_v137;
          _2104_v137 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).fm4((_2077_cf51).f11, _dafny.Map.Empty.slice().updateUnsafe(_1941_v41,_1950_v49), (_dafny.ZERO).minus(_module.__default.fm2(new BigNumber(-392), (_this).f11, globalState)), (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], globalState)),((_2078_cf50) ? (_2102_v135) : (_2102_v135)));
          _2104_v137 = (_2104_v137).update((_this).fm4(new BigNumber((_2101_cf5).cardinality()), function () {
            let _coll81 = new _dafny.Map();
            for (const _compr_81 of (_dafny.Set.fromElements(_1941_v41)).Elements) {
              let _2105_v138 = _compr_81;
              if ((_dafny.Set.fromElements(_1941_v41)).contains(_2105_v138)) {
                _coll81.push([_2105_v138,_1950_v49]);
              }
            }
            return _coll81;
          }(), new BigNumber(977), (_this).f13, globalState), _2102_v135);
        }
        (globalState).f7 = (_this).fm4((_this).f11, (_dafny.Map.Empty.slice().updateUnsafe((_1941_v41).update((_2077_cf51).f11, _module.__default.abs((_2077_cf51).f11)),_dafny.Map.Empty.slice().updateUnsafe((_2077_cf51).f11,(_2077_cf51).f11))).Merge(_1951_v50), (_2080_v122).dtor_cf7, false, globalState);
        let _index293 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
        (_1938_v37)[_index293] = (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))];
      } else if (_source29.is_DC36) {
        let _2106___mcc_h10 = (_source29).cf55;
        let _2107___mcc_h11 = (_source29).cf56;
        let _2108___mcc_h12 = (_source29).cf57;
        let _2109___mcc_h13 = (_source29).cf58;
        let _2110___mcc_h14 = (_source29).cf59;
        let _2111_cf59 = _2110___mcc_h14;
        let _2112_cf58 = _2109___mcc_h13;
        let _2113_cf57 = _2108___mcc_h12;
        let _2114_cf56 = _2107___mcc_h11;
        let _2115_cf55 = _2106___mcc_h10;
        if (_2112_cf58) {
          (globalState).f7 = _module.__default.fm0(globalState);
          _1934_v33 = _dafny.Seq.Concat(_1934_v33, _1934_v33);
          _2112_cf58 = _2114_cf56;
          let _index294 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
          (_1938_v37)[_index294] = ((_2114_cf56) ? (_2112_cf58) : ((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]));
          let _2116_v139;
          _2116_v139 = _module.D8.create_DC16(_2114_cf56, _2115_cf55);
          let _2117_v140;
          _2117_v140 = _module.D8.create_DC17(_2116_v139);
          _2117_v140 = _2117_v140;
        } else {
          let _2118_v141;
          let _nw339 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _2118_v141 = _nw339;
          let _index295 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_2118_v141).length));
          (_2118_v141)[_index295] = (_this).f11;
          let _index296 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_2118_v141).length));
          (_2118_v141)[_index296] = new BigNumber(-571);
          (globalState).f1 = _2111_cf59;
          let _index297 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
          (_1938_v37)[_index297] = _dafny.Seq.IsPrefixOf(_module.__default.fm19((_this).f13, false, globalState), _dafny.Seq.Concat(_1956_v55, _1956_v55));
          let _2119_v142;
          let _nw340 = new _module.C2();
          _nw340.__ctor(p0, (_this).f11);
          _2119_v142 = _nw340;
          _1941_v41 = (_dafny.MultiSet.fromElements((_this).f11, _2111_cf59, _2111_cf59)).Union(_1941_v41);
        }
        let _2120_v143;
        _2120_v143 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_1934_v33);
        let _2121_v144;
        _2121_v144 = _dafny.Seq.of(_module.__default.fm22((((_2120_v143).contains(true)) ? ((_2120_v143).get(true)) : (_dafny.Seq.UnicodeFromString("gau"))), _2112_cf58, globalState), _1940_v39);
        let _2122_v145;
        _2122_v145 = _module.D8.create_DC16(_2113_cf57, (((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]) ? (new BigNumber((_2121_v144).length)) : ((_this).f11)));
        let _2123_v146;
        _2123_v146 = _dafny.Map.Empty.slice().updateUnsafe(!(_1949_v48).contains((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]),_2112_cf58);
        let _2124_v147;
        let _nw341 = new _module.C3();
        _nw341.__ctor(_2115_cf55, _2115_cf55, (_this).f11, (_this).f12);
        _2124_v147 = _nw341;
        let _2125_v148;
        _2125_v148 = _dafny.Set.fromElements(_2124_v147);
        let _2126_v149;
        _2126_v149 = _dafny.Seq.of(_2125_v148, _2125_v148, _2125_v148, _2125_v148);
        let _index298 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
        let _index299 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
        let _rhs366 = (_1956_v55)[_module.__default.safeIndex((_2111_cf59).multipliedBy(_2115_cf55), new BigNumber((_1956_v55).length))];
        let _rhs367 = (_1947_v46).fm3(((_this).f11).plus((_this).f11), _2112_cf58, !((((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]) ? (true) : ((_this).f13))), globalState);
        let _rhs368 = ((_2112_cf58) ? ((_this).f13) : ((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]));
        let _rhs369 = _module.D8.create_DC16((_this).f13, (_this).f11);
        let _rhs370 = (((_2123_v146).contains(_2113_cf57)) ? ((_2123_v146).get(_2113_cf57)) : (_dafny.Seq.IsProperPrefixOf(_2126_v149, _2126_v149)));
        let _lhs313 = _1938_v37;
        let _lhs314 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
        let _lhs315 = globalState;
        let _lhs316 = _1938_v37;
        let _lhs317 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length));
        _lhs313[_lhs314] = _rhs366;
        _lhs315.f1 = _rhs367;
        _lhs316[_lhs317] = _rhs368;
        _2122_v145 = _rhs369;
        _2113_cf57 = _rhs370;
        let _2127_v150;
        _2127_v150 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_2123_v146));
        let _2128_v151;
        _2128_v151 = _dafny.Map.Empty.slice().updateUnsafe(_1934_v33,(_2127_v150)[_module.__default.safeIndex((_this).f11, new BigNumber((_2127_v150).length))]);
        let _2129_v152;
        _2129_v152 = _dafny.MultiSet.fromElements(_2123_v146, _2123_v146);
        _2128_v151 = (_2128_v151).update(_dafny.Seq.UnicodeFromString("fq"), ((_2114_cf56) ? (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_2113_cf57,(_this).f13))) : (_2129_v152)));
        _2115_cf55 = (_2124_v147).fm3(new BigNumber(((_2124_v147).f12).length), _2114_cf56, (_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], globalState);
      } else if (_source29.is_DC33) {
        let _2130___mcc_h15 = (_source29).cf49;
        let _2131_cf49 = _2130___mcc_h15;
        _1949_v48 = (_1949_v48).update(false, (_this).f11);
        if (((_this).f11).isLessThanOrEqualTo((_this).f11)) {
          let _2132_v153;
          _2132_v153 = _module.D21.create_DC48((_this).f13);
          let _2133_v154;
          _2133_v154 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let54_0) {
            return function (_2134_dt__update__tmp_h5) {
              return function (_pat_let55_0) {
                return function (_2135_dt__update_hcf78_h0) {
                  return _module.D21.create_DC48(_2135_dt__update_hcf78_h0);
                }(_pat_let55_0);
              }((_this).f13);
            }(_pat_let54_0);
          }(_2132_v153),(_this).f13);
          (globalState).f1 = _module.__default.safeModuloInt((new BigNumber((_2133_v154).length)).multipliedBy(new BigNumber(-913)), ((_this).f11).plus(new BigNumber(391)));
          _1950_v49 = (_1950_v49).Merge((_1950_v49).Merge(_1950_v49));
          _1957_v56 = _1957_v56;
          (globalState).f1 = (_this).f11;
          let _2136_v155;
          let _2137_v156;
          let _out35;
          let _out36;
          let _outcollector12 = (_this).m2(_module.__default.safeModuloInt(new BigNumber(222), (_this).f11), globalState);
          _out35 = _outcollector12[0];
          _out36 = _outcollector12[1];
          _2136_v155 = _out35;
          _2137_v156 = _out36;
        } else {
          (globalState).f1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uoj"), _1934_v33)).length);
          (globalState).f1 = new BigNumber(328);
          (globalState).f1 = new BigNumber(632);
          let _2138_v157;
          let _nw342 = Array((new BigNumber(5)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2138_v157 = _nw342;
          let _2139_v158;
          _2139_v158 = _dafny.MultiSet.fromElements((_this).f13);
          let _index300 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_2138_v157).length));
          (_2138_v157)[_index300] = _2139_v158;
          let _index301 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_2138_v157).length));
          (_2138_v157)[_index301] = _2139_v158;
          let _2140_v159;
          let _nw343 = Array((new BigNumber(6)).toNumber()).fill([]);
          _2140_v159 = _nw343;
          let _2141_v160;
          let _nw344 = Array((new BigNumber(9)).toNumber());
          _nw344[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw344[(_dafny.ONE).toNumber()] = new BigNumber(283);
          _nw344[(new BigNumber(2)).toNumber()] = (_this).f11;
          _nw344[(new BigNumber(3)).toNumber()] = (_this).f11;
          _nw344[(new BigNumber(4)).toNumber()] = (_this).f11;
          _nw344[(new BigNumber(5)).toNumber()] = (_this).f11;
          _nw344[(new BigNumber(6)).toNumber()] = (_this).f11;
          _nw344[(new BigNumber(7)).toNumber()] = (_this).f11;
          _nw344[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("dlq")).length);
          _2141_v160 = _nw344;
          let _2142_v161;
          _2142_v161 = _dafny.Map.Empty.slice().updateUnsafe(_2140_v159,_2141_v160);
          let _2143_v162;
          _2143_v162 = _module.D11.create_DC22(_2142_v161);
          let _2144_v163;
          _2144_v163 = _dafny.MultiSet.fromElements(_2143_v162);
          let _2145_v164;
          _2145_v164 = _module.D7.create_DC13((_this).f12, new BigNumber((_2144_v163).cardinality()), (_this).f11);
          let _2146_v165;
          let _nw345 = new _module.C7();
          _nw345.__ctor((_1947_v46).f14, (_2145_v164).dtor_cf18, (_this).f11);
          _2146_v165 = _nw345;
        }
        let _2147_v166;
        _2147_v166 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(319),_1940_v39);
        let _2148_v168;
        _2148_v168 = _dafny.Seq.of(_2147_v166, function () {
          let _coll82 = new _dafny.Map();
          for (const _compr_82 of (_1932_v31).Elements) {
            let _2149_v167 = _compr_82;
            if (_dafny.Seq.contains(_1932_v31, _2149_v167)) {
              _coll82.push([(_2149_v167).plus((_this).f11),_1940_v39]);
            }
          }
          return _coll82;
        }());
        _2147_v166 = (_2147_v166).Merge((_2148_v168)[_module.__default.safeIndex((_this).f11, new BigNumber((_2148_v168).length))]);
        let _2150_v169;
        _2150_v169 = _dafny.Seq.of(_1938_v37, _1938_v37, _1935_v34);
        let _rhs371 = _2150_v169;
        let _rhs372 = (_this).f13;
        let _lhs318 = globalState;
        _2150_v169 = _rhs371;
        _lhs318.f7 = _rhs372;
      } else {
        let _2151___mcc_h16 = (_source29).cf60;
        let _2152_cf60 = _2151___mcc_h16;
        (globalState).f1 = (((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]) ? ((_this).f11) : (_module.__default.safeDivisionInt((_this).f11, new BigNumber((_dafny.Seq.update(_1934_v33, _module.__default.safeIndex(_module.__default.fm2((_this).fm6((_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))], new BigNumber(240), globalState), (_this).f11, globalState), new BigNumber((_1934_v33).length)), p0)).length))));
        (globalState).f1 = new BigNumber((_1934_v33).length);
        let _2153_v170;
        let _2154_v171;
        let _out37;
        let _out38;
        let _outcollector13 = (_this).m2(new BigNumber((_dafny.Seq.update(_1934_v33, _module.__default.safeIndex((_this).f11, new BigNumber((_1934_v33).length)), p0)).length), globalState);
        _out37 = _outcollector13[0];
        _out38 = _outcollector13[1];
        _2153_v170 = _out37;
        _2154_v171 = _out38;
        let _2155_v172;
        _2155_v172 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1938_v37)[_module.__default.safeIndex(new BigNumber(945), new BigNumber((_1938_v37).length))]);
        _2155_v172 = (_dafny.Map.Empty.slice().updateUnsafe(p0,_2154_v171)).Merge(_2155_v172);
      }
      r0 = _module.__default.fm37(globalState);
      return r0;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.Map.Empty;
      let _2156_v0;
      let _init40 = function (_2157_i0) {
        return _module.__default.safeDivisionInt(_2157_i0, (_this).f11);
      };
      let _nw346 = Array((new BigNumber(13)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw346.length); _i0_40++) {
        _nw346[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _2156_v0 = _nw346;
      let _2158_v1;
      _2158_v1 = new _dafny.CodePoint('j'.codePointAt(0));
      let _index302 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length));
      (_2156_v0)[_index302] = (_dafny.ZERO).minus((_this).fm8(_2158_v1, globalState));
      let _2159_v2;
      _2159_v2 = _dafny.Seq.UnicodeFromString("h");
      let _2160_v3;
      _2160_v3 = _dafny.Seq.of((_this).f13);
      let _index303 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length));
      (_2156_v0)[_index303] = (new BigNumber((_dafny.Seq.Concat(_2159_v2, _2159_v2)).length)).plus(new BigNumber((_2160_v3).length));
      _2158_v1 = _2158_v1;
      let _2161_v4;
      _2161_v4 = _dafny.MultiSet.fromElements((_2160_v3)[_module.__default.safeIndex((_this).f11, new BigNumber((_2160_v3).length))]);
      let _2162_v5;
      _2162_v5 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),(_this).fm7((_dafny.ZERO).minus(new BigNumber((_2161_v4).cardinality())), (_this).f13, _2161_v4, globalState));
      let _hi5 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-346)), ((_2163_v1) => function (_2164_i2) {
        return _2163_v1;
      })(_2158_v1))).length);
      for (let _2165_i1 = (((_2162_v5).contains((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))])) ? ((_2162_v5).get((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))])) : (p0)); _2165_i1.isLessThan(_hi5); _2165_i1 = _2165_i1.plus(_dafny.ONE)) {
        let _2166_v6;
        _2166_v6 = _dafny.MultiSet.fromElements(_2165_i1, (_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))]);
        (globalState).f1 = ((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))]).plus((_dafny.ZERO).minus(((_this).f11).multipliedBy((_this).fm3(new BigNumber(-774), (_this).fm4(_2165_i1, _dafny.Map.Empty.slice().updateUnsafe(_2166_v6,function () {
          let _coll83 = new _dafny.Map();
          for (const _compr_83 of _dafny.IntegerRange(new BigNumber(163), new BigNumber(568))) {
            let _2167_v7 = _compr_83;
            if (((new BigNumber(163)).isLessThanOrEqualTo(_2167_v7)) && ((_2167_v7).isLessThan(new BigNumber(568)))) {
              _coll83.push([(_2167_v7).multipliedBy(new BigNumber(252)),(_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))]]);
            }
          }
          return _coll83;
        }()), (_this).f11, (_this).f13, globalState), (_this).f13, globalState))));
        (globalState).f1 = new BigNumber(-211);
        let _2168_v8;
        _2168_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_2158_v1);
        if (!((_2168_v8).equals(_2168_v8)) || (((_this).f12).IsSubsetOf((_this).f12))) {
          let _2169_v9;
          let _nw347 = new _module.C1();
          _nw347.__ctor((_this).f13, (_this).f13);
          _2169_v9 = _nw347;
          let _2170_v10;
          _2170_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2169_v9,(_2169_v9).f18);
          _2170_v10 = (_2170_v10).update(_2169_v9, (_this).f13);
          let _2171_v11;
          let _nw348 = Array((new BigNumber(20)).toNumber());
          _2171_v11 = _nw348;
          _2171_v11 = _2171_v11;
          let _2172_v12;
          let _nw349 = new _module.C2();
          _nw349.__ctor(_2158_v1, ((_dafny.ZERO).minus(p0)).minus((_dafny.ZERO).minus(p0)));
          _2172_v12 = _nw349;
          let _2173_v13;
          _2173_v13 = _dafny.Map.Empty.slice().updateUnsafe((_2169_v9).f18,_2166_v6);
          let _2174_v14;
          _2174_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2173_v13).length),!((_2169_v9).f19));
          let _rhs373 = false;
          let _rhs374 = (_2174_v14).Merge(function () {
            let _coll84 = new _dafny.Map();
            for (const _compr_84 of _dafny.IntegerRange(new BigNumber(923), new BigNumber(-68))) {
              let _2175_v15 = _compr_84;
              if (((new BigNumber(923)).isLessThanOrEqualTo(_2175_v15)) && ((_2175_v15).isLessThan(new BigNumber(-68)))) {
                _coll84.push([_module.__default.safeModuloInt(_2175_v15, p0),false]);
              }
            }
            return _coll84;
          }());
          let _lhs319 = globalState;
          _lhs319.f7 = _rhs373;
          _2174_v14 = _rhs374;
          let _2176_v16;
          _2176_v16 = _dafny.Seq.of(_2162_v5, _2162_v5);
          let _2177_v17;
          _2177_v17 = _dafny.Set.fromElements(p0, (_this).f11);
          let _2178_v18;
          _2178_v18 = _dafny.Map.Empty.slice().updateUnsafe((_2176_v16)[_module.__default.safeIndex(new BigNumber((_2177_v17).length), new BigNumber((_2176_v16).length))],(_2172_v12).f24);
          _2178_v18 = _2178_v18;
        } else {
          (globalState).f6 = !_dafny.areEqual(_2159_v2, _dafny.Seq.Concat(_dafny.Seq.update(_2159_v2, _module.__default.safeIndex(_2165_i1, new BigNumber((_2159_v2).length)), _2158_v1), _2159_v2));
          let _2179_v19;
          let _init41 = function (_2180_i3) {
            return ((_this).f13) && (!((_this).f13));
          };
          let _nw350 = Array((new BigNumber(18)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw350.length); _i0_41++) {
            _nw350[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _2179_v19 = _nw350;
          _2179_v19 = _2179_v19;
          (globalState).f6 = ((_this).f12).IsProperSubsetOf(((_this).f12).Intersect(_dafny.Set.fromElements((_this).f13, (_this).f13)));
          _2159_v2 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(175)), ((_2181_v1) => function (_2182_i4) {
            return _2181_v1;
          })(_2158_v1));
          (globalState).f6 = !((_this).f13) || ((_this).f13);
        }
        let _2183_v20;
        _2183_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2158_v1,true);
        let _2184_v21;
        _2184_v21 = _dafny.Seq.of(_2183_v20);
        let _rhs375 = (_this).f13;
        let _rhs376 = !(!((_this).f13) || ((_this).f13));
        let _rhs377 = _2184_v21;
        let _rhs378 = ((_this).f13) === (!((((_this).f13) ? ((_this).f13) : ((_this).f13))));
        let _rhs379 = ((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))]).isLessThanOrEqualTo(((_this).f11).minus((_dafny.ZERO).minus((((_2166_v6).contains(new BigNumber((_2159_v2).length))) ? ((_2166_v6).get(new BigNumber((_2159_v2).length))) : ((_this).f11)))));
        let _lhs320 = globalState;
        let _lhs321 = globalState;
        let _lhs322 = globalState;
        let _lhs323 = globalState;
        _lhs320.f7 = _rhs375;
        _lhs321.f7 = _rhs376;
        _2184_v21 = _rhs377;
        _lhs322.f6 = _rhs378;
        _lhs323.f6 = _rhs379;
      }
      let _hi6 = (_this).f11;
      for (let _2185_i5 = (_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))]; _2185_i5.isLessThan(_hi6); _2185_i5 = _2185_i5.plus(_dafny.ONE)) {
        let _2186_v22;
        _2186_v22 = _dafny.Seq.of(new BigNumber((_2160_v3).length));
        let _rhs380 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(291)), ((_2187_i5) => function (_2188_i6) {
          return (_2187_i5).minus((_this).f11);
        })(_2185_i5));
        let _rhs381 = (_2185_i5).minus(_2185_i5);
        let _lhs324 = globalState;
        _2186_v22 = _rhs380;
        _lhs324.f1 = _rhs381;
        let _2189_v23;
        let _nw351 = Array((new BigNumber(8)).toNumber()).fill([]);
        _2189_v23 = _nw351;
        let _index304 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_2189_v23).length));
        (_2189_v23)[_index304] = _2156_v0;
        let _index305 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_2189_v23).length));
        (_2189_v23)[_index305] = _2156_v0;
        let _2190_v24;
        _2190_v24 = _dafny.MultiSet.fromElements((_this).f11);
        let _2191_v25;
        _2191_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2190_v24,(_dafny.ZERO).minus((_this).f11));
        _2191_v25 = (_2191_v25).update(_2190_v24, _module.__default.fm2((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))], (_this).f11, globalState));
        let _2192_v26;
        _2192_v26 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt(_2185_i5, new BigNumber((_module.__default.fm19((_this).f13, (_this).f13, globalState)).length))),_2159_v2);
        _2159_v2 = _dafny.Seq.update((((_2192_v26).contains((_dafny.ZERO).minus((_this).f11))) ? ((_2192_v26).get((_dafny.ZERO).minus((_this).f11))) : (_2159_v2)), _module.__default.safeIndex(_2185_i5, new BigNumber(((((_2192_v26).contains((_dafny.ZERO).minus((_this).f11))) ? ((_2192_v26).get((_dafny.ZERO).minus((_this).f11))) : (_2159_v2))).length)), _2158_v1);
      }
      let _hi7 = p0;
      for (let _2193_i7 = (_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))]; _2193_i7.isLessThan(_hi7); _2193_i7 = _2193_i7.plus(_dafny.ONE)) {
        let _2194_v27;
        _2194_v27 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f13),(_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))]);
        if (((_2160_v3)[_module.__default.safeIndex(p0, new BigNumber((_2160_v3).length))]) === ((_2194_v27).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_dafny.ZERO).minus(_2193_i7))))) {
          let _2195_v28;
          let _nw352 = new _module.C5();
          _nw352.__ctor();
          _2195_v28 = _nw352;
          (globalState).f7 = (_this).f13;
          let _2196_v29;
          let _nw353 = Array((new BigNumber(25)).toNumber()).fill(false);
          _2196_v29 = _nw353;
          let _2197_v30;
          _2197_v30 = _dafny.Seq.of((_this).f11);
          let _2198_v31;
          _2198_v31 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f13);
          let _2199_v32;
          _2199_v32 = _module.D6.create_DC10(_2160_v3);
          let _2200_v33;
          _2200_v33 = _module.D19.create_DC43(false, new _dafny.CodePoint('r'.codePointAt(0)), (_this).f13, _2199_v32, _2159_v2);
          let _index306 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length));
          let _index307 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length));
          let _rhs382 = _2196_v29;
          let _rhs383 = (_2197_v30)[_module.__default.safeIndex(new BigNumber((_2198_v31).length), new BigNumber((_2197_v30).length))];
          let _rhs384 = (((_2200_v33).dtor_cf69) ? (p0) : ((_this).f11));
          let _rhs385 = !((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))]).isEqualTo((_this).f11);
          let _lhs325 = _2156_v0;
          let _lhs326 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length));
          let _lhs327 = _2156_v0;
          let _lhs328 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length));
          let _lhs329 = globalState;
          _2196_v29 = _rhs382;
          _lhs325[_lhs326] = _rhs383;
          _lhs327[_lhs328] = _rhs384;
          _lhs329.f7 = _rhs385;
          let _2201_v34;
          _2201_v34 = _dafny.Set.fromElements((_this).f11, _2193_i7);
          let _2202_v35;
          _2202_v35 = _dafny.MultiSet.fromElements(_2201_v34, _2201_v34);
          let _2203_v36;
          _2203_v36 = _dafny.Map.Empty.slice().updateUnsafe((((_2202_v35).contains(_2201_v34)) ? ((_2202_v35).get(_2201_v34)) : (p0)),(_this).f13);
          let _2204_v37;
          let _nw354 = new _module.C3();
          _nw354.__ctor(((_dafny.ZERO).minus(p0)).minus(new BigNumber((_2203_v36).length)), _module.__default.safeModuloInt((_this).f11, (_2197_v30)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13))).length), new BigNumber((_2197_v30).length))]), new BigNumber((_2197_v30).length), (_this).f12);
          _2204_v37 = _nw354;
          let _2205_v38;
          _2205_v38 = _dafny.Seq.of(_2197_v30, _2197_v30, _2197_v30, _2197_v30);
          let _2206_v39;
          _2206_v39 = _dafny.MultiSet.fromElements(new BigNumber((_2159_v2).length));
          let _2207_v40;
          _2207_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2206_v39,_2162_v5);
          let _2208_v41;
          let _nw355 = new _module.C7();
          _nw355.__ctor(_2205_v38, (_this).f12, (_dafny.ZERO).minus(((!((_this).fm4(p0, _2207_v40, (_2204_v37).f22, (_this).f13, globalState))) ? (new BigNumber((_2160_v3).length)) : ((_dafny.ZERO).minus((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))])))));
          _2208_v41 = _nw355;
          let _rhs386 = _2208_v41;
          let _rhs387 = _2160_v3;
          _2208_v41 = _rhs386;
          _2160_v3 = _rhs387;
        } else {
          let _2209_v42;
          _2209_v42 = _dafny.Set.fromElements(_2159_v2, _2159_v2);
          _2209_v42 = (((_this).f13) ? (_2209_v42) : (_2209_v42));
          let _2210_v43;
          let _nw356 = new _module.C5();
          _nw356.__ctor();
          _2210_v43 = _nw356;
          let _2211_v44;
          let _nw357 = Array((new BigNumber(15)).toNumber()).fill(false);
          _2211_v44 = _nw357;
          let _2212_v45;
          _2212_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2158_v1,(_this).f13);
          let _index308 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2211_v44).length));
          (_2211_v44)[_index308] = (((_2212_v45).contains(_2158_v1)) ? ((_2212_v45).get(_2158_v1)) : ((_this).f13));
          let _index309 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2211_v44).length));
          (_2211_v44)[_index309] = (_this).f13;
          let _index310 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_2211_v44).length));
          (_2211_v44)[_index310] = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tgekcnj"), _2159_v2)).length)).isLessThan((_this).fm6((_this).f13, (_this).f11, globalState));
          (globalState).f7 = !(((_2211_v44)[_module.__default.safeIndex(new BigNumber(217), new BigNumber((_2211_v44).length))]) || (false));
        }
        let _index311 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length));
        (_2156_v0)[_index311] = ((_this).f11).multipliedBy((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))]);
        let _index312 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length));
        (_2156_v0)[_index312] = (_this).f11;
        let _index313 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length));
        (_2156_v0)[_index313] = p0;
      }
      let _2213_v46;
      _2213_v46 = _dafny.MultiSet.fromElements(new BigNumber(629));
      let _2214_v48;
      _2214_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2213_v46,function () {
        let _coll85 = new _dafny.Map();
        for (const _compr_85 of _dafny.IntegerRange(new BigNumber(737), new BigNumber(-125))) {
          let _2215_v47 = _compr_85;
          if (((new BigNumber(737)).isLessThanOrEqualTo(_2215_v47)) && ((_2215_v47).isLessThan(new BigNumber(-125)))) {
            _coll85.push([_module.__default.safeModuloInt(_2215_v47, new BigNumber((_dafny.MultiSet.fromElements(_2161_v4)).cardinality())),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))],true)).length)]);
          }
        }
        return _coll85;
      }());
      let _2216_v49;
      _2216_v49 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),!((_this).f13));
      let _2217_v50;
      let _nw358 = Array((new BigNumber(12)).toNumber());
      _nw358[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).fm4((_this).f11, _2214_v48, new BigNumber(128), true, globalState));
      _nw358[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13);
      _nw358[(new BigNumber(2)).toNumber()] = _2216_v49;
      _nw358[(new BigNumber(3)).toNumber()] = _2216_v49;
      _nw358[(new BigNumber(4)).toNumber()] = _2216_v49;
      _nw358[(new BigNumber(5)).toNumber()] = (_2216_v49).Merge(_2216_v49);
      _nw358[(new BigNumber(6)).toNumber()] = (_2216_v49).Merge(_module.__default.fm11(globalState));
      _nw358[(new BigNumber(7)).toNumber()] = _module.__default.fm11(globalState);
      _nw358[(new BigNumber(8)).toNumber()] = _2216_v49;
      _nw358[(new BigNumber(9)).toNumber()] = _2216_v49;
      _nw358[(new BigNumber(10)).toNumber()] = _2216_v49;
      _nw358[(new BigNumber(11)).toNumber()] = _module.__default.fm11(globalState);
      _2217_v50 = _nw358;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2217_v50).length))) {
        let _2218_i8 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2218_i8)) && ((_2218_i8).isLessThan(new BigNumber((_2217_v50).length))))) {
          (_2217_v50)[(_2218_i8)] = _2216_v49;
        }
      }
      let _2219_v51;
      let _init42 = ((_2220_v4) => function (_2221_i9) {
        return _2220_v4;
      })(_2161_v4);
      let _nw359 = Array((new BigNumber(10)).toNumber());
      for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw359.length); _i0_42++) {
        _nw359[_i0_42] = _init42(new BigNumber(_i0_42));
      }
      _2219_v51 = _nw359;
      let _2222_v52;
      _2222_v52 = _dafny.Seq.of(_2219_v51, _2219_v51, _2219_v51);
      r0 = (_2222_v52)[_module.__default.safeIndex((_2156_v0)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2156_v0).length))], new BigNumber((_2222_v52).length))];
      let _2223_v53;
      _2223_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2159_v2,_2213_v46);
      r1 = (((_2160_v3)[_module.__default.safeIndex((_this).fm8(_2158_v1, globalState), new BigNumber((_2160_v3).length))]) ? (_2223_v53) : ((_2223_v53).Merge(_2223_v53)));
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _2224_v0;
      _2224_v0 = new _dafny.CodePoint('f'.codePointAt(0));
      let _2225_v1;
      _2225_v1 = _dafny.Seq.of(_2224_v0, _2224_v0);
      r1 = _dafny.Seq.contains(_2225_v1, _2224_v0);
      r1 = (_this).f13;
      r0 = (_this).f11;
      if (true) {
        (globalState).f7 = ((_this).f13) === ((_this).f13);
        let _2226_v2;
        _2226_v2 = _dafny.Seq.of((_this).f13, (_this).f13, false, (_this).f13, (_this).f13);
        let _2227_v3;
        _2227_v3 = _module.D6.create_DC10(_2226_v2);
        let _2228_v4;
        let _nw360 = new _module.C0();
        _nw360.__ctor((_2227_v3).dtor_cf13, (_this).f11);
        _2228_v4 = _nw360;
        let _2229_v5;
        _2229_v5 = _module.D23.create_DC54(_2228_v4);
        let _2230_v6;
        let _nw361 = Array((new BigNumber(13)).toNumber());
        _nw361[(_dafny.ZERO).toNumber()] = _2228_v4;
        _nw361[(_dafny.ONE).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(2)).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(3)).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(4)).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(5)).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(6)).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(7)).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(8)).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(9)).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(10)).toNumber()] = (_2229_v5).dtor_cf91;
        _nw361[(new BigNumber(11)).toNumber()] = _2228_v4;
        _nw361[(new BigNumber(12)).toNumber()] = _2228_v4;
        _2230_v6 = _nw361;
        let _index314 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_2230_v6).length));
        (_2230_v6)[_index314] = _2228_v4;
        let _index315 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_2230_v6).length));
        (_2230_v6)[_index315] = _2228_v4;
        let _2231_v7;
        _2231_v7 = _module.D15.create_DC34();
        let _source32 = _2231_v7;
        if (_source32.is_DC34) {
          let _2232_v8;
          _2232_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2224_v0,(_this).f13);
          _2232_v8 = (((_this).f13) ? ((_2232_v8).Merge(_2232_v8)) : ((_2232_v8).Merge((_dafny.Map.Empty.slice().updateUnsafe(_2224_v0,(_this).f13)).update(_2224_v0, (_this).f13))));
          let _2233_v9;
          _2233_v9 = _dafny.Set.fromElements(_2224_v0, _2224_v0);
          let _2234_v10;
          _2234_v10 = _dafny.Seq.of(_2233_v9);
          let _2235_v11;
          _2235_v11 = _dafny.MultiSet.fromElements((_2234_v10)[_module.__default.safeIndex(new BigNumber(621), new BigNumber((_2234_v10).length))]);
          r1 = (_2235_v11).IsSubsetOf(_dafny.MultiSet.FromArray(_2234_v10));
          _2224_v0 = _2224_v0;
          let _rhs388 = (_this).f13;
          let _rhs389 = (_dafny.ZERO).minus((_2228_v4).f21);
          let _rhs390 = _2224_v0;
          let _rhs391 = _module.__default.fm2((_2228_v4).f21, (_dafny.ZERO).minus((_2228_v4).f21), globalState);
          let _lhs330 = globalState;
          let _lhs331 = globalState;
          let _lhs332 = globalState;
          _lhs330.f6 = _rhs388;
          _lhs331.f1 = _rhs389;
          _2224_v0 = _rhs390;
          _lhs332.f1 = _rhs391;
        } else if (_source32.is_DC35) {
          let _2236___mcc_h0 = (_source32).cf50;
          let _2237___mcc_h1 = (_source32).cf51;
          let _2238___mcc_h2 = (_source32).cf52;
          let _2239___mcc_h3 = (_source32).cf53;
          let _2240___mcc_h4 = (_source32).cf54;
          let _2241_cf54 = _2240___mcc_h4;
          let _2242_cf53 = _2239___mcc_h3;
          let _2243_cf52 = _2238___mcc_h2;
          let _2244_cf51 = _2237___mcc_h1;
          let _2245_cf50 = _2236___mcc_h0;
          let _2246_v12;
          let _nw362 = Array((new BigNumber(11)).toNumber()).fill(false);
          _2246_v12 = _nw362;
          let _index316 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2246_v12).length));
          (_2246_v12)[_index316] = (_this).f13;
          let _index317 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2246_v12).length));
          (_2246_v12)[_index317] = (_this).f13;
          let _rhs392 = (_this).f13;
          let _rhs393 = new BigNumber(845);
          let _rhs394 = _2243_cf52;
          let _lhs333 = globalState;
          let _lhs334 = globalState;
          let _lhs335 = globalState;
          _lhs333.f7 = _rhs392;
          _lhs334.f1 = _rhs393;
          _lhs335.f7 = _rhs394;
          let _2247_v13;
          let _init43 = ((_2248_v4) => function (_2249_i0) {
            return (_2249_i0).minus((_2248_v4).f21);
          })(_2228_v4);
          let _nw363 = Array((new BigNumber(16)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw363.length); _i0_43++) {
            _nw363[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _2247_v13 = _nw363;
          let _2250_v14;
          _2250_v14 = _dafny.Set.fromElements(_2247_v13, _2247_v13, _2247_v13, _2247_v13);
          let _2251_v15;
          _2251_v15 = _module.D2.create_DC4(p0, _2241_cf54, new BigNumber((_2242_cf53).length));
          let _2252_v16;
          let _nw364 = new _module.C6();
          _nw364.__ctor(_2251_v15, _dafny.Set.fromElements(_2245_cf50), (_2228_v4).f21);
          _2252_v16 = _nw364;
          let _2253_v17;
          _2253_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2252_v16,_2245_cf50);
          let _2254_v18;
          _2254_v18 = _dafny.Seq.of((_2228_v4).f21, new BigNumber(((_2253_v17).update(_2252_v16, (_this).f13)).length));
          let _2255_v19;
          _2255_v19 = _dafny.Seq.of(_2254_v18, _dafny.Seq.update(_dafny.Seq.of(p0, _2241_cf54, (_this).f11), _module.__default.safeIndex((_2228_v4).f21, new BigNumber((_dafny.Seq.of(p0, _2241_cf54, (_this).f11)).length)), (_this).f11), _dafny.Seq.of((_2244_cf51).f11), _2254_v18, _2254_v18);
          let _2256_v20;
          let _nw365 = new _module.C7();
          _nw365.__ctor(_2255_v19, (_this).f12, p0);
          _2256_v20 = _nw365;
          let _2257_v21;
          let _init44 = ((_2258_v0) => function (_2259_i1) {
            return _2258_v0;
          })(_2224_v0);
          let _nw366 = Array((new BigNumber(10)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw366.length); _i0_44++) {
            _nw366[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _2257_v21 = _nw366;
          let _index318 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_2257_v21).length));
          (_2257_v21)[_index318] = _2224_v0;
          let _2260_v22;
          _2260_v22 = _dafny.Seq.of(_2256_v20);
          let _index319 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_2257_v21).length));
          let _rhs395 = _2243_cf52;
          let _rhs396 = (_2250_v14).Intersect((_dafny.Set.fromElements(_2247_v13)).Union(_2250_v14));
          let _rhs397 = (_2260_v22)[_module.__default.safeIndex((_2228_v4).f21, new BigNumber((_2260_v22).length))];
          let _rhs398 = (_this).f11;
          let _rhs399 = _2224_v0;
          let _lhs336 = globalState;
          let _lhs337 = globalState;
          let _lhs338 = _2257_v21;
          let _lhs339 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_2257_v21).length));
          _lhs336.f6 = _rhs395;
          _2250_v14 = _rhs396;
          _2256_v20 = _rhs397;
          _lhs337.f1 = _rhs398;
          _lhs338[_lhs339] = _rhs399;
          let _2261_v23;
          _2261_v23 = _dafny.Set.fromElements(_2246_v12);
          let _2262_v24;
          let _nw367 = Array((new BigNumber(10)).toNumber());
          _nw367[(_dafny.ZERO).toNumber()] = _2261_v23;
          _nw367[(_dafny.ONE).toNumber()] = (_2261_v23).Intersect(_2261_v23);
          _nw367[(new BigNumber(2)).toNumber()] = (_2261_v23).Union(_2261_v23);
          _nw367[(new BigNumber(3)).toNumber()] = _2261_v23;
          _nw367[(new BigNumber(4)).toNumber()] = _2261_v23;
          _nw367[(new BigNumber(5)).toNumber()] = (_module.D21.create_DC47(_2261_v23)).dtor_cf77;
          _nw367[(new BigNumber(6)).toNumber()] = (_2261_v23).Difference(_2261_v23);
          _nw367[(new BigNumber(7)).toNumber()] = (_2261_v23).Difference(_2261_v23);
          _nw367[(new BigNumber(8)).toNumber()] = _2261_v23;
          _nw367[(new BigNumber(9)).toNumber()] = _2261_v23;
          _2262_v24 = _nw367;
          let _index320 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_2262_v24).length));
          (_2262_v24)[_index320] = _2261_v23;
          let _index321 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_2262_v24).length));
          (_2262_v24)[_index321] = ((_2261_v23).Union(_2261_v23)).Union(_2261_v23);
        } else if (_source32.is_DC36) {
          let _2263___mcc_h5 = (_source32).cf55;
          let _2264___mcc_h6 = (_source32).cf56;
          let _2265___mcc_h7 = (_source32).cf57;
          let _2266___mcc_h8 = (_source32).cf58;
          let _2267___mcc_h9 = (_source32).cf59;
          let _2268_cf59 = _2267___mcc_h9;
          let _2269_cf58 = _2266___mcc_h8;
          let _2270_cf57 = _2265___mcc_h7;
          let _2271_cf56 = _2264___mcc_h6;
          let _2272_cf55 = _2263___mcc_h5;
          let _2273_v25;
          _2273_v25 = _dafny.MultiSet.fromElements(_module.__default.fm2((_2228_v4).f21, (_2228_v4).f21, globalState));
          let _2274_v26;
          _2274_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2270_cf57,_2268_cf59);
          let _2275_v27;
          _2275_v27 = _dafny.Set.fromElements((_2273_v25).contains(_2268_cf59), (new BigNumber(939)).isLessThanOrEqualTo((((_2274_v26).contains(_2269_cf58)) ? ((_2274_v26).get(_2269_cf58)) : (_2268_cf59))), !(new BigNumber(819)).isEqualTo(new BigNumber(369)), _2270_cf57, _2271_cf56);
          let _2276_v28;
          _2276_v28 = _module.D7.create_DC13(_module.__default.fm41(globalState), (_2228_v4).f21, (_2228_v4).f21);
          let _pat_let_tv54 = _2275_v27;
          _2275_v27 = (function (_pat_let56_0) {
            return function (_2277_dt__update__tmp_h0) {
              return function (_pat_let57_0) {
                return function (_2278_dt__update_hcf18_h0) {
                  return _module.D7.create_DC13(_2278_dt__update_hcf18_h0, (_2277_dt__update__tmp_h0).dtor_cf19, (_2277_dt__update__tmp_h0).dtor_cf20);
                }(_pat_let57_0);
              }(_pat_let_tv54);
            }(_pat_let56_0);
          }(_2276_v28)).dtor_cf18;
          _2224_v0 = _2224_v0;
          let _2279_v29;
          let _init45 = function (_2280_i2) {
            return (_this).f13;
          };
          let _nw368 = Array((new BigNumber(18)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw368.length); _i0_45++) {
            _nw368[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _2279_v29 = _nw368;
          let _index322 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_2279_v29).length));
          (_2279_v29)[_index322] = !(_2270_cf57);
          let _index323 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_2279_v29).length));
          (_2279_v29)[_index323] = _2270_cf57;
          _2225_v1 = _dafny.Seq.Concat(_2225_v1, _2225_v1);
        } else if (_source32.is_DC33) {
          let _2281___mcc_h10 = (_source32).cf49;
          let _2282_cf49 = _2281___mcc_h10;
          let _2283_v30;
          _2283_v30 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_2228_v4).f21),(_2228_v4).f21);
          _2283_v30 = (_2283_v30).update(new BigNumber(574), (_this).f11);
          (globalState).f7 = (_this).f13;
          let _index324 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_2230_v6).length));
          (_2230_v6)[_index324] = _2228_v4;
          let _2284_v31;
          let _nw369 = Array((new BigNumber(12)).toNumber()).fill(false);
          _2284_v31 = _nw369;
          let _2285_v32;
          _2285_v32 = _module.D2.create_DC4((_2228_v4).f21, new BigNumber(-724), (_2228_v4).f21);
          let _2286_v33;
          let _nw370 = new _module.C6();
          _nw370.__ctor(_2285_v32, (_this).f12, (_2228_v4).f21);
          _2286_v33 = _nw370;
          let _2287_v34;
          _2287_v34 = _dafny.MultiSet.fromElements(_2286_v33, _2286_v33, _2286_v33);
          let _2288_v35;
          _2288_v35 = _dafny.Set.fromElements(_2287_v34);
          let _index325 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_2284_v31).length));
          (_2284_v31)[_index325] = (_2288_v35).IsProperSubsetOf(_2288_v35);
          let _index326 = _module.__default.safeIndex(new BigNumber(79), new BigNumber((_2284_v31).length));
          (_2284_v31)[_index326] = false;
        } else {
          let _2289___mcc_h11 = (_source32).cf60;
          let _2290_cf60 = _2289___mcc_h11;
          r1 = (_this).f13;
          (globalState).f7 = true;
          (globalState).f1 = (_2228_v4).f21;
          let _2291_v36;
          _2291_v36 = _module.D11.create_DC24(new BigNumber(994), (_2228_v4).f20, _2224_v0);
          let _2292_v37;
          _2292_v37 = _dafny.MultiSet.fromElements((_this).f13, (_this).f13, false);
          let _2293_v38;
          _2293_v38 = _dafny.Seq.of((_this).fm7(p0, (_this).f13, _2292_v37, globalState));
          let _2294_v39;
          _2294_v39 = _module.D15.create_DC36((_this).f11, (_this).f13, (_this).f13, (_this).f13, new BigNumber(369));
          let _2295_v40;
          _2295_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2224_v0,(_this).f13);
          let _2296_v41;
          _2296_v41 = _dafny.MultiSet.fromElements((_this).f11);
          let _2297_v42;
          _2297_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2296_v41).cardinality()),(_this).fm6(false, (_2228_v4).f21, globalState));
          let _2298_v43;
          _2298_v43 = _module.D7.create_DC13(_module.__default.fm41(globalState), (_this).f11, (_2228_v4).f21);
          let _2299_v44;
          let _nw371 = Array((new BigNumber(29)).toNumber());
          _nw371[(_dafny.ZERO).toNumber()] = (_2291_v36).dtor_cf34;
          _nw371[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Set.fromElements((_this).f13)).length);
          _nw371[(new BigNumber(2)).toNumber()] = (_this).fm8(_2224_v0, globalState);
          _nw371[(new BigNumber(3)).toNumber()] = (_2228_v4).f21;
          _nw371[(new BigNumber(4)).toNumber()] = (_this).f11;
          _nw371[(new BigNumber(5)).toNumber()] = new BigNumber(822);
          _nw371[(new BigNumber(6)).toNumber()] = new BigNumber((_2293_v38).length);
          _nw371[(new BigNumber(7)).toNumber()] = (_this).f11;
          _nw371[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2225_v1, _2225_v1)).length);
          _nw371[(new BigNumber(9)).toNumber()] = ((_dafny.ZERO).minus((_2228_v4).f21)).minus(p0);
          _nw371[(new BigNumber(10)).toNumber()] = ((_2294_v39).dtor_cf55).multipliedBy((_this).f11);
          _nw371[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_2295_v40).length)).plus(p0));
          _nw371[(new BigNumber(12)).toNumber()] = new BigNumber(((_2297_v42).Merge(_dafny.Map.Empty.slice().updateUnsafe((_2228_v4).f21,(_this).f11))).length);
          _nw371[(new BigNumber(13)).toNumber()] = p0;
          _nw371[(new BigNumber(14)).toNumber()] = p0;
          _nw371[(new BigNumber(15)).toNumber()] = new BigNumber((_2292_v37).cardinality());
          _nw371[(new BigNumber(16)).toNumber()] = (((_this).f13) ? ((_this).f11) : ((_this).f11));
          _nw371[(new BigNumber(17)).toNumber()] = p0;
          _nw371[(new BigNumber(18)).toNumber()] = p0;
          _nw371[(new BigNumber(19)).toNumber()] = (((_2228_v4).fm21((_this).f13, new BigNumber((_2293_v38).length), globalState)) ? ((_this).f11) : (new BigNumber(8)));
          _nw371[(new BigNumber(20)).toNumber()] = new BigNumber(-315);
          _nw371[(new BigNumber(21)).toNumber()] = new BigNumber((_2292_v37).cardinality());
          _nw371[(new BigNumber(22)).toNumber()] = (_this).f11;
          _nw371[(new BigNumber(23)).toNumber()] = (_2298_v43).dtor_cf20;
          _nw371[(new BigNumber(24)).toNumber()] = (_this).f11;
          _nw371[(new BigNumber(25)).toNumber()] = (_2228_v4).f21;
          _nw371[(new BigNumber(26)).toNumber()] = (_2228_v4).f21;
          _nw371[(new BigNumber(27)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_2228_v4).f21), (_2228_v4).f21);
          _nw371[(new BigNumber(28)).toNumber()] = (_this).f11;
          _2299_v44 = _nw371;
          let _index327 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2299_v44).length));
          (_2299_v44)[_index327] = p0;
          let _index328 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2299_v44).length));
          (_2299_v44)[_index328] = (_this).fm6((_this).f13, (_2228_v4).f21, globalState);
        }
        let _2300_v45;
        let _init46 = ((_2301_v1) => function (_2302_i3) {
          return _2301_v1;
        })(_2225_v1);
        let _nw372 = Array((new BigNumber(10)).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw372.length); _i0_46++) {
          _nw372[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _2300_v45 = _nw372;
        let _index329 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_2300_v45).length));
        (_2300_v45)[_index329] = _2225_v1;
        let _index330 = _module.__default.safeIndex(new BigNumber(806), new BigNumber((_2300_v45).length));
        (_2300_v45)[_index330] = _dafny.Seq.update(_2225_v1, _module.__default.safeIndex(p0, new BigNumber((_2225_v1).length)), _2224_v0);
        let _2303_v46;
        _2303_v46 = _dafny.Seq.of((_2228_v4).f21);
        let _2304_v47;
        _2304_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2224_v0,_2303_v46);
        let _2305_v48;
        _2305_v48 = _dafny.Seq.of(_2303_v46, (((_2304_v47).contains(new _dafny.CodePoint('f'.codePointAt(0)))) ? ((_2304_v47).get(new _dafny.CodePoint('f'.codePointAt(0)))) : (_2303_v46)));
        let _2306_v49;
        let _nw373 = new _module.C7();
        _nw373.__ctor(_2305_v48, (_this).f12, (_2228_v4).f21);
        _2306_v49 = _nw373;
      } else {
        let _2307_v50;
        _2307_v50 = _dafny.Seq.of((_this).f13, (_this).f13, !((_this).f13));
        let _2308_v51;
        _2308_v51 = _dafny.MultiSet.fromElements(_2307_v50, _2307_v50, _2307_v50, _2307_v50, _2307_v50);
        (globalState).f1 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber((_2225_v1).length), (_this).f11), new BigNumber((_2308_v51).cardinality()));
        let _2309_v52;
        _2309_v52 = _module.D2.create_DC4(p0, p0, p0);
        let _2310_v53;
        let _nw374 = new _module.C6();
        _nw374.__ctor(function (_pat_let58_0) {
          return function (_2311_dt__update__tmp_h1) {
            return function (_pat_let59_0) {
              return function (_2312_dt__update_hcf8_h0) {
                return _module.D2.create_DC4((_2311_dt__update__tmp_h1).dtor_cf6, (_2311_dt__update__tmp_h1).dtor_cf7, _2312_dt__update_hcf8_h0);
              }(_pat_let59_0);
            }((_this).f11);
          }(_pat_let58_0);
        }(_2309_v52), _module.__default.fm41(globalState), p0);
        _2310_v53 = _nw374;
        let _2313_v54;
        _2313_v54 = _module.D13.create_DC28(_2225_v1, p0, (_this).f13);
        r0 = ((_2313_v54).dtor_cf41).minus((_this).f11);
        let _2314_v55;
        let _nw375 = Array((new BigNumber(22)).toNumber()).fill([]);
        _2314_v55 = _nw375;
        let _2315_v56;
        let _nw376 = Array((new BigNumber(12)).toNumber()).fill([]);
        _2315_v56 = _nw376;
        let _index331 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_2314_v55).length));
        (_2314_v55)[_index331] = _2315_v56;
        let _index332 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_2314_v55).length));
        let _nw377 = Array((new BigNumber(8)).toNumber()).fill([]);
        (_2314_v55)[_index332] = _nw377;
        if ((_this).f13) {
          let _2316_v57;
          let _init47 = ((_2317_v0) => function (_2318_i4) {
            return _2317_v0;
          })(_2224_v0);
          let _nw378 = Array((new BigNumber(3)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw378.length); _i0_47++) {
            _nw378[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _2316_v57 = _nw378;
          let _2319_v58;
          _2319_v58 = _dafny.Seq.of(_2316_v57, _2316_v57, _2316_v57);
          let _2320_v59;
          _2320_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_2316_v57);
          let _2321_v60;
          let _nw379 = Array((new BigNumber(20)).toNumber());
          _nw379[(_dafny.ZERO).toNumber()] = _2316_v57;
          _nw379[(_dafny.ONE).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(2)).toNumber()] = (_2319_v58)[_module.__default.safeIndex(p0, new BigNumber((_2319_v58).length))];
          _nw379[(new BigNumber(3)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(4)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(5)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(6)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(7)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(8)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(9)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(10)).toNumber()] = (((_2320_v59).contains((_this).f13)) ? ((_2320_v59).get((_this).f13)) : ((_2319_v58)[_module.__default.safeIndex(p0, new BigNumber((_2319_v58).length))]));
          _nw379[(new BigNumber(11)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(12)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(13)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(14)).toNumber()] = (((_this).f13) ? (_2316_v57) : (_2316_v57));
          _nw379[(new BigNumber(15)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(16)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(17)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(18)).toNumber()] = _2316_v57;
          _nw379[(new BigNumber(19)).toNumber()] = _2316_v57;
          _2321_v60 = _nw379;
          let _index333 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2321_v60).length));
          (_2321_v60)[_index333] = _2316_v57;
          let _index334 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2321_v60).length));
          (_2321_v60)[_index334] = _2316_v57;
          let _2322_v61;
          let _nw380 = new _module.C2();
          _nw380.__ctor(new _dafny.CodePoint('s'.codePointAt(0)), p0);
          _2322_v61 = _nw380;
          let _2323_v62;
          _2323_v62 = _dafny.Seq.of(p0);
          (globalState).f1 = ((p0).minus((_2323_v62)[_module.__default.safeIndex((_this).f11, new BigNumber((_2323_v62).length))])).plus(_module.__default.safeDivisionInt(p0, new BigNumber((_2225_v1).length)));
          let _2324_v63;
          _2324_v63 = _module.D7.create_DC12(_2323_v62);
          _2324_v63 = _2324_v63;
          let _2325_v64;
          let _nw381 = Array((new BigNumber(18)).toNumber()).fill(false);
          _2325_v64 = _nw381;
          let _2326_v65;
          _2326_v65 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f11).multipliedBy(new BigNumber((_2225_v1).length)),_2325_v64);
          let _2327_v66;
          _2327_v66 = _dafny.Set.fromElements((_this).f11, (_this).f11);
          let _2328_v67;
          _2328_v67 = _dafny.MultiSet.fromElements(new BigNumber((_2327_v66).length));
          let _rhs400 = new BigNumber((_2326_v65).length);
          let _rhs401 = !(!((_this).f13));
          let _rhs402 = (_2328_v67).IsSubsetOf((_2328_v67).Intersect(_dafny.MultiSet.FromArray(_2323_v62)));
          let _rhs403 = _module.__default.safeModuloInt(p0, (_this).f11);
          let _rhs404 = (_this).f13;
          let _lhs340 = globalState;
          let _lhs341 = globalState;
          let _lhs342 = globalState;
          let _lhs343 = globalState;
          r0 = _rhs400;
          _lhs340.f6 = _rhs401;
          _lhs341.f7 = _rhs402;
          _lhs342.f1 = _rhs403;
          _lhs343.f6 = _rhs404;
        } else {
          let _2329_v68;
          let _nw382 = Array((new BigNumber(25)).toNumber()).fill(false);
          _2329_v68 = _nw382;
          let _2330_v69;
          _2330_v69 = _dafny.Set.fromElements(_2329_v68);
          let _2331_v70;
          let _nw383 = new _module.C6();
          _nw383.__ctor((_2310_v53).f15, (_this).f12, new BigNumber((_2330_v69).length));
          _2331_v70 = _nw383;
          let _2332_v71;
          _2332_v71 = _dafny.Map.Empty.slice().updateUnsafe(_2331_v70,p0);
          let _2333_v72;
          _2333_v72 = _dafny.Seq.of(new BigNumber((_2332_v71).length), new BigNumber((_2225_v1).length));
          let _2334_v73;
          _2334_v73 = _dafny.Seq.of(_2333_v72);
          let _2335_v74;
          _2335_v74 = _dafny.Map.Empty.slice().updateUnsafe((_2331_v70).f11,(_dafny.ZERO).minus((_this).f11));
          let _2336_v75;
          let _nw384 = new _module.C7();
          _nw384.__ctor(_2334_v73, (_dafny.Set.fromElements((_this).f13)).Intersect((_this).f12), ((((_2335_v74).contains((_dafny.ZERO).minus(new BigNumber((_2225_v1).length)))) ? ((_2335_v74).get((_dafny.ZERO).minus(new BigNumber((_2225_v1).length)))) : (new BigNumber((_2225_v1).length)))).minus((_2331_v70).f11));
          _2336_v75 = _nw384;
          let _index335 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_2329_v68).length));
          (_2329_v68)[_index335] = (((_this).f13) ? ((_this).f13) : ((_this).f13));
          let _index336 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_2329_v68).length));
          (_2329_v68)[_index336] = ((_2331_v70).f12).IsProperSubsetOf(((_this).f12).Union((_this).f12));
          let _2337_v76;
          _2337_v76 = _dafny.MultiSet.fromElements((_this).f13, (_2329_v68)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_2329_v68).length))]);
          let _2338_v77;
          _2338_v77 = _2337_v76;
          let _2339_v78;
          _2339_v78 = _dafny.Map.Empty.slice().updateUnsafe(!((_2329_v68)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_2329_v68).length))]),_2338_v77);
          let _2340_v79;
          _2340_v79 = _module.D0.create_DC1((_this).f13, (_2331_v70).f11, (_this).f13);
          let _2341_v80;
          _2341_v80 = _dafny.Seq.of(_2340_v79);
          let _2342_v81;
          _2342_v81 = _module.D8.create_DC16((_this).f13, new BigNumber((_2341_v80).length));
          let _2343_v82;
          _2343_v82 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2224_v0);
          let _2344_v83;
          _2344_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_2329_v68)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_2329_v68).length))]);
          let _2345_v84;
          _2345_v84 = _module.D7.create_DC13((_2331_v70).f12, (_this).f11, (_this).f11);
          let _2346_v85;
          _2346_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_2345_v84);
          let _2347_v86;
          _2347_v86 = _module.D7.create_DC14((((_2346_v85).contains((_this).f11)) ? ((_2346_v85).get((_this).f11)) : (_2345_v84)));
          let _2348_v87;
          _2348_v87 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f11);
          let _2349_v88;
          let _nw385 = Array((new BigNumber(29)).toNumber());
          _nw385[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw385[(_dafny.ONE).toNumber()] = p0;
          _nw385[(new BigNumber(2)).toNumber()] = p0;
          _nw385[(new BigNumber(3)).toNumber()] = (_this).f11;
          _nw385[(new BigNumber(4)).toNumber()] = new BigNumber((_2339_v78).length);
          _nw385[(new BigNumber(5)).toNumber()] = (_2331_v70).f11;
          _nw385[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(p0, p0);
          _nw385[(new BigNumber(7)).toNumber()] = ((_dafny.ZERO).minus((_2331_v70).f11)).minus((_2342_v81).dtor_cf24);
          _nw385[(new BigNumber(8)).toNumber()] = new BigNumber((_2343_v82).length);
          _nw385[(new BigNumber(9)).toNumber()] = (new BigNumber(787)).multipliedBy(p0);
          _nw385[(new BigNumber(10)).toNumber()] = p0;
          _nw385[(new BigNumber(11)).toNumber()] = (_this).f11;
          _nw385[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_this).f11);
          _nw385[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((((_this).f13) ? ((_dafny.ZERO).minus(new BigNumber((_2344_v83).length))) : (p0)));
          _nw385[(new BigNumber(14)).toNumber()] = ((_module.__default.fm0(globalState)) ? ((_2331_v70).f11) : ((_2336_v75).fm6((_this).f13, (_2331_v70).f11, globalState)));
          _nw385[(new BigNumber(15)).toNumber()] = (_this).f11;
          _nw385[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_2347_v86)).length);
          _nw385[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt((_this).f11, new BigNumber((_2337_v76).cardinality()));
          _nw385[(new BigNumber(18)).toNumber()] = p0;
          _nw385[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt(p0, p0);
          _nw385[(new BigNumber(20)).toNumber()] = new BigNumber(-888);
          _nw385[(new BigNumber(21)).toNumber()] = (_this).f11;
          _nw385[(new BigNumber(22)).toNumber()] = p0;
          _nw385[(new BigNumber(23)).toNumber()] = (_2331_v70).f11;
          _nw385[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus(((_this).f11).plus((_2331_v70).f11));
          _nw385[(new BigNumber(25)).toNumber()] = ((_this).f11).plus((_2331_v70).f11);
          _nw385[(new BigNumber(26)).toNumber()] = (((_2329_v68)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_2329_v68).length))]) ? (new BigNumber((_2344_v83).length)) : (new BigNumber((_2348_v87).length)));
          _nw385[(new BigNumber(27)).toNumber()] = p0;
          _nw385[(new BigNumber(28)).toNumber()] = p0;
          _2349_v88 = _nw385;
          let _init48 = ((_2350_v70) => function (_2351_i5) {
            return _module.__default.safeDivisionInt(_2351_i5, (_2350_v70).f11);
          })(_2331_v70);
          let _nw386 = Array((new BigNumber(7)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw386.length); _i0_48++) {
            _nw386[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _2349_v88 = _nw386;
          r1 = (_2329_v68)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_2329_v68).length))];
          (globalState).f1 = ((p0).multipliedBy(new BigNumber(714))).multipliedBy(new BigNumber((function () {
            let _coll86 = new _dafny.Set();
            for (const _compr_86 of (_2344_v83).Keys.Elements) {
              let _2352_v89 = _compr_86;
              if ((_2344_v83).contains(_2352_v89)) {
                _coll86.add((_2352_v89).minus(new BigNumber(611)));
              }
            }
            return _coll86;
          }()).length));
        }
      }
      let _2353_v90;
      _2353_v90 = _dafny.Seq.of((_this).f11);
      let _2354_v91;
      _2354_v91 = _dafny.MultiSet.fromElements(p0, p0, (_this).f11);
      let _2355_v92;
      let _out39;
      _out39 = (_this).m3((_this).f13, new BigNumber(-814), (_this).f11, (_2354_v91).IsProperSubsetOf(_dafny.MultiSet.FromArray(_2353_v90)), globalState);
      _2355_v92 = _out39;
      let _2356_v93;
      _2356_v93 = _dafny.Seq.of(_module.__default.fm0(globalState));
      let _2357_v94;
      _2357_v94 = _dafny.Seq.of((_2356_v93)[_module.__default.safeIndex(p0, new BigNumber((_2356_v93).length))], (_this).f13);
      let _2358_v95;
      _2358_v95 = _module.D0.create_DC0(new BigNumber((_dafny.Seq.update(_2356_v93, _module.__default.safeIndex(p0, new BigNumber((_2356_v93).length)), !((_this).f13))).length));
      let _2359_v96;
      _2359_v96 = _dafny.MultiSet.fromElements(_2224_v0, new _dafny.CodePoint('u'.codePointAt(0)));
      let _2360_v97;
      let _out40;
      _out40 = (_this).m3((_this).fm5(_2358_v95, _2357_v94, new BigNumber((_2359_v96).cardinality()), _2224_v0, globalState), (_this).f11, (_this).f11, (_this).f13, globalState);
      _2360_v97 = _out40;
      r0 = _module.__default.safeDivisionInt(p0, (_this).f11);
      r1 = (_this).f13;
      return [r0, r1];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _2361_v0;
      _2361_v0 = _dafny.Seq.of(true);
      _2361_v0 = _dafny.Seq.Concat(_2361_v0, _2361_v0);
      let _2362_v1;
      _2362_v1 = new _dafny.CodePoint('v'.codePointAt(0));
      let _2363_v2;
      _2363_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2362_v1,p0);
      _2363_v2 = _2363_v2;
      let _2364_v3;
      _2364_v3 = _dafny.Seq.UnicodeFromString("ifvvwmi");
      if (((!((_this).f13)) ? (false) : (!((new BigNumber((_2364_v3).length)).isLessThanOrEqualTo((_this).f11))))) {
        let _2365_v4;
        let _init49 = ((_2366_v0) => function (_2367_i0) {
          return _module.__default.safeModuloInt(_2367_i0, new BigNumber((_2366_v0).length));
        })(_2361_v0);
        let _nw387 = Array((new BigNumber(16)).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw387.length); _i0_49++) {
          _nw387[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _2365_v4 = _nw387;
        let _2368_v5;
        _2368_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2365_v4,p0);
        let _2369_v6;
        _2369_v6 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(_2365_v4,p0)).update(_2365_v4, (_this).f13));
        let _2370_v7;
        _2370_v7 = _module.D19.create_DC42(_2365_v4);
        let _2371_v8;
        let _nw388 = Array((new BigNumber(15)).toNumber());
        _nw388[(_dafny.ZERO).toNumber()] = _2368_v5;
        _nw388[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2365_v4,_module.__default.fm0(globalState));
        _nw388[(new BigNumber(2)).toNumber()] = ((true) ? (_2368_v5) : (_2368_v5));
        _nw388[(new BigNumber(3)).toNumber()] = (_2368_v5).Merge(_2368_v5);
        _nw388[(new BigNumber(4)).toNumber()] = (_2368_v5).Merge(_2368_v5);
        _nw388[(new BigNumber(5)).toNumber()] = (((_this).f13) ? ((_2369_v6)[_module.__default.safeIndex((_this).f11, new BigNumber((_2369_v6).length))]) : (_2368_v5));
        _nw388[(new BigNumber(6)).toNumber()] = _2368_v5;
        _nw388[(new BigNumber(7)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_2370_v7).dtor_cf68,p3)).Merge(_2368_v5);
        _nw388[(new BigNumber(8)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2365_v4,p0)).Merge(_2368_v5);
        _nw388[(new BigNumber(9)).toNumber()] = _2368_v5;
        _nw388[(new BigNumber(10)).toNumber()] = _2368_v5;
        _nw388[(new BigNumber(11)).toNumber()] = _2368_v5;
        _nw388[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2365_v4,(_this).f13);
        _nw388[(new BigNumber(13)).toNumber()] = _2368_v5;
        _nw388[(new BigNumber(14)).toNumber()] = (_2368_v5).Merge(_2368_v5);
        _2371_v8 = _nw388;
        let _index337 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2371_v8).length));
        (_2371_v8)[_index337] = _2368_v5;
        let _index338 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2371_v8).length));
        (_2371_v8)[_index338] = _2368_v5;
        let _2372_v9;
        _2372_v9 = _dafny.Set.fromElements((_this).f11);
        let _2373_v10;
        _2373_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2372_v9).length),p0);
        let _2374_v11;
        _2374_v11 = _dafny.Seq.of(_2373_v10);
        let _2375_v12;
        _2375_v12 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_2374_v11).length));
        let _2376_v13;
        let _nw389 = Array((new BigNumber(22)).toNumber());
        _nw389[(_dafny.ZERO).toNumber()] = _2362_v1;
        _nw389[(_dafny.ONE).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(2)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(3)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(4)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(5)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(6)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
        _nw389[(new BigNumber(8)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(9)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(10)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw389[(new BigNumber(12)).toNumber()] = _module.__default.fm20(_2361_v0, new BigNumber(663), p0, globalState);
        _nw389[(new BigNumber(13)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(14)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(15)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(16)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(17)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(18)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(19)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(20)).toNumber()] = _2362_v1;
        _nw389[(new BigNumber(21)).toNumber()] = _2362_v1;
        _2376_v13 = _nw389;
        let _2377_v14;
        _2377_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2376_v13,_2362_v1);
        let _2378_v15;
        _2378_v15 = _dafny.MultiSet.fromElements((((_2375_v12).contains((_this).f13)) ? ((_2375_v12).get((_this).f13)) : (new BigNumber((_2377_v14).length))), (_dafny.ZERO).minus((_this).f11));
        let _2379_v16;
        _2379_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2378_v15,_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(673)));
        let _2380_v17;
        _2380_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm1(globalState)).length),p2);
        (globalState).f7 = !((_this).fm4((_dafny.ZERO).minus(p1), (_2379_v16).update(_2378_v15, _dafny.Map.Empty.slice().updateUnsafe((_this).f11,new BigNumber((_dafny.MultiSet.fromElements(p2, p2)).cardinality()))), (((_2380_v17).contains((_this).f11)) ? ((_2380_v17).get((_this).f11)) : (new BigNumber(-482))), p3, globalState));
        _2365_v4 = ((((_this).f11).isLessThan((_this).f11)) ? (((p3) ? (_2365_v4) : (_2365_v4))) : (_2365_v4));
        _2376_v13 = _2376_v13;
        let _2381_v18;
        let _nw390 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
        _2381_v18 = _nw390;
        let _rhs405 = p1;
        let _rhs406 = _2381_v18;
        let _rhs407 = p0;
        let _lhs344 = globalState;
        let _lhs345 = globalState;
        _lhs344.f1 = _rhs405;
        _2381_v18 = _rhs406;
        _lhs345.f7 = _rhs407;
      } else {
        let _2382_v19;
        _2382_v19 = _dafny.MultiSet.fromElements((_this).f11);
        let _2383_v21;
        _2383_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2382_v19,function () {
          let _coll87 = new _dafny.Map();
          for (const _compr_87 of _dafny.IntegerRange(new BigNumber(-821), new BigNumber(806))) {
            let _2384_v20 = _compr_87;
            if (((new BigNumber(-821)).isLessThanOrEqualTo(_2384_v20)) && ((_2384_v20).isLessThan(new BigNumber(806)))) {
              _coll87.push([(_2384_v20).plus((_this).f11),p1]);
            }
          }
          return _coll87;
        }());
        let _2385_v22;
        let _nw391 = Array((new BigNumber(19)).toNumber()).fill(false);
        _2385_v22 = _nw391;
        let _2386_v23;
        _2386_v23 = _dafny.Seq.of(_2385_v22);
        (globalState).f7 = (_this).fm4(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-837)), function (_2387_i1) {
          return (_dafny.ZERO).minus(new BigNumber(((_this).f12).length));
        })).length), _2383_v21, _module.__default.safeModuloInt((_this).f11, (_this).f11), _dafny.Seq.IsProperPrefixOf(_2386_v23, _2386_v23), globalState);
        (globalState).f1 = (_this).fm8(_2362_v1, globalState);
        let _2388_v24;
        _2388_v24 = _dafny.Seq.of(new BigNumber(25));
        _2388_v24 = _2388_v24;
        let _2389_v25;
        _2389_v25 = _module.D6.create_DC11(p2, new BigNumber(647), (_this).f11);
        let _2390_v26;
        _2390_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm34(p3, _2389_v25, globalState)).length),new BigNumber(60));
        let _2391_v27;
        _2391_v27 = _dafny.Seq.of(_2390_v26);
        let _2392_v28;
        _2392_v28 = _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), ((_2393_p1) => function (_2394_i2) {
          return new BigNumber((_dafny.Seq.of(_2393_p1)).length);
        })(p1)), _module.__default.safeIndex(new BigNumber(340), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), ((_2395_p1) => function (_2396_i2) {
          return new BigNumber((_dafny.Seq.of(_2395_p1)).length);
        })(p1))).length)), p2), _2388_v24);
        let _2397_v31;
        _2397_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2362_v1);
        let _2398_v32;
        _2398_v32 = _dafny.Map.Empty.slice().updateUnsafe(p3,((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2364_v3).length),new BigNumber((_2397_v31).length))).update(p2, (_this).f11)).update(p1, (_this).f11));
        let _2399_v33;
        let _nw392 = Array((new BigNumber(19)).toNumber());
        _nw392[(_dafny.ZERO).toNumber()] = (_2391_v27)[_module.__default.safeIndex(new BigNumber((_2361_v0).length), new BigNumber((_2391_v27).length))];
        _nw392[(_dafny.ONE).toNumber()] = _2390_v26;
        _nw392[(new BigNumber(2)).toNumber()] = ((p3) ? ((_2390_v26).update(new BigNumber(-527), new BigNumber((_2392_v28).length))) : (_2390_v26));
        _nw392[(new BigNumber(3)).toNumber()] = (_2390_v26).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,p2));
        _nw392[(new BigNumber(4)).toNumber()] = function () {
          let _coll88 = new _dafny.Map();
          for (const _compr_88 of _dafny.IntegerRange(new BigNumber(760), new BigNumber(727))) {
            let _2400_v29 = _compr_88;
            if (((new BigNumber(760)).isLessThanOrEqualTo(_2400_v29)) && ((_2400_v29).isLessThan(new BigNumber(727)))) {
              _coll88.push([(_2400_v29).plus((_this).f11),new BigNumber(86)]);
            }
          }
          return _coll88;
        }();
        _nw392[(new BigNumber(5)).toNumber()] = (_2390_v26).update((((_2382_v19).contains(p2)) ? ((_2382_v19).get(p2)) : (new BigNumber(103))), p2);
        _nw392[(new BigNumber(6)).toNumber()] = function () {
          let _coll89 = new _dafny.Map();
          for (const _compr_89 of _dafny.IntegerRange(new BigNumber(2), new BigNumber(656))) {
            let _2401_v30 = _compr_89;
            if (((new BigNumber(2)).isLessThanOrEqualTo(_2401_v30)) && ((_2401_v30).isLessThan(new BigNumber(656)))) {
              _coll89.push([_module.__default.safeModuloInt(_2401_v30, (_this).f11),new BigNumber(533)]);
            }
          }
          return _coll89;
        }();
        _nw392[(new BigNumber(7)).toNumber()] = _module.__default.fm46(globalState);
        _nw392[(new BigNumber(8)).toNumber()] = ((!(false)) ? (_dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(254)), function (_2402_i3) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        })).length)))) : (_2390_v26));
        _nw392[(new BigNumber(9)).toNumber()] = _module.__default.fm46(globalState);
        _nw392[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,p2);
        _nw392[(new BigNumber(11)).toNumber()] = (_2390_v26).Merge(_2390_v26);
        _nw392[(new BigNumber(12)).toNumber()] = _2390_v26;
        _nw392[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2398_v32).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(704)), ((_2403_v1) => function (_2404_i4) {
          return _2403_v1;
        })(_2362_v1))).length));
        _nw392[(new BigNumber(14)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(p2, new BigNumber(493), globalState),new BigNumber(256))).Merge(_2390_v26);
        _nw392[(new BigNumber(15)).toNumber()] = (_2390_v26).Merge(_2390_v26);
        _nw392[(new BigNumber(16)).toNumber()] = _2390_v26;
        _nw392[(new BigNumber(17)).toNumber()] = _2390_v26;
        _nw392[(new BigNumber(18)).toNumber()] = _2390_v26;
        _2399_v33 = _nw392;
        let _index339 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_2399_v33).length));
        (_2399_v33)[_index339] = _2390_v26;
        let _index340 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_2399_v33).length));
        let _rhs408 = (_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(157))).Merge((_module.__default.fm46(globalState)).Merge(_2390_v26));
        let _rhs409 = (new BigNumber(138)).minus(_module.__default.safeModuloInt(p2, (_dafny.ZERO).minus((_this).f11)));
        let _lhs346 = _2399_v33;
        let _lhs347 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_2399_v33).length));
        let _lhs348 = globalState;
        _lhs346[_lhs347] = _rhs408;
        _lhs348.f1 = _rhs409;
        let _2405_v34;
        _2405_v34 = _module.D6.create_DC10(_2361_v0);
        let _2406_v35;
        _2406_v35 = _module.D19.create_DC43(p3, _2362_v1, p3, _2405_v34, _2364_v3);
        let _2407_v36;
        _2407_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_2364_v3);
        let _2408_v37;
        let _nw393 = Array((new BigNumber(16)).toNumber());
        _nw393[(_dafny.ZERO).toNumber()] = ((p0) ? (_dafny.Seq.UnicodeFromString("nowe")) : (_2364_v3));
        _nw393[(_dafny.ONE).toNumber()] = _2364_v3;
        _nw393[(new BigNumber(2)).toNumber()] = _2364_v3;
        _nw393[(new BigNumber(3)).toNumber()] = _2364_v3;
        _nw393[(new BigNumber(4)).toNumber()] = _2364_v3;
        _nw393[(new BigNumber(5)).toNumber()] = _2364_v3;
        _nw393[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_2364_v3, _2364_v3);
        _nw393[(new BigNumber(7)).toNumber()] = _2364_v3;
        _nw393[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(115)), ((_2409_v1) => function (_2410_i5) {
          return _2409_v1;
        })(_2362_v1)), (_2406_v35).dtor_cf73);
        _nw393[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_2364_v3, _2364_v3);
        _nw393[(new BigNumber(10)).toNumber()] = (((_2407_v36).contains(p2)) ? ((_2407_v36).get(p2)) : (_2364_v3));
        _nw393[(new BigNumber(11)).toNumber()] = _2364_v3;
        _nw393[(new BigNumber(12)).toNumber()] = _2364_v3;
        _nw393[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("lcgs");
        _nw393[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), ((_2411_v1) => function (_2412_i6) {
          return _2411_v1;
        })(_2362_v1)), _2364_v3);
        _nw393[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_2364_v3, _2364_v3);
        _2408_v37 = _nw393;
        let _index341 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_2408_v37).length));
        (_2408_v37)[_index341] = _dafny.Seq.Concat(_2364_v3, _dafny.Seq.UnicodeFromString("dhbyei"));
        let _index342 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_2408_v37).length));
        (_2408_v37)[_index342] = _2364_v3;
      }
      if (p3) {
        let _2413_v38;
        _2413_v38 = _dafny.MultiSet.fromElements(p0, false);
        let _2414_v39;
        let _nw394 = new _module.C3();
        _nw394.__ctor((_this).f11, (((_2413_v38).contains((_this).f13)) ? ((_2413_v38).get((_this).f13)) : (p2)), (_this).f11, _dafny.Set.fromElements(p3));
        _2414_v39 = _nw394;
        let _2415_v40;
        _2415_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2414_v39,(_this).f11);
        let _2416_v41;
        _2416_v41 = _dafny.MultiSet.fromElements(_2415_v40);
        let _2417_v42;
        _2417_v42 = _dafny.MultiSet.fromElements((_2414_v39).f11);
        let _2418_v43;
        let _nw395 = Array((new BigNumber(22)).toNumber());
        _nw395[(_dafny.ZERO).toNumber()] = (((_2416_v41).contains(_2415_v40)) ? ((_2416_v41).get(_2415_v40)) : ((_2414_v39).f11));
        _nw395[(_dafny.ONE).toNumber()] = (_2414_v39).f11;
        _nw395[(new BigNumber(2)).toNumber()] = p2;
        _nw395[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt((_2414_v39).f11, (_this).f11);
        _nw395[(new BigNumber(4)).toNumber()] = (_this).f11;
        _nw395[(new BigNumber(5)).toNumber()] = (_this).f11;
        _nw395[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_2364_v3, _module.__default.safeIndex(p1, new BigNumber((_2364_v3).length)), _2362_v1), _dafny.Seq.UnicodeFromString("lysmecy"))).length);
        _nw395[(new BigNumber(7)).toNumber()] = p1;
        _nw395[(new BigNumber(8)).toNumber()] = ((_this).fm7(new BigNumber((_2417_v42).cardinality()), (_this).f13, _2413_v38, globalState)).minus((_this).f11);
        _nw395[(new BigNumber(9)).toNumber()] = p2;
        _nw395[(new BigNumber(10)).toNumber()] = (((_this).f13) ? ((_dafny.ZERO).minus(p1)) : (new BigNumber(710)));
        _nw395[(new BigNumber(11)).toNumber()] = (_2414_v39).f11;
        _nw395[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-666), (_this).f11);
        _nw395[(new BigNumber(13)).toNumber()] = p2;
        _nw395[(new BigNumber(14)).toNumber()] = new BigNumber(93);
        _nw395[(new BigNumber(15)).toNumber()] = _module.__default.safeModuloInt(p1, new BigNumber((_2364_v3).length));
        _nw395[(new BigNumber(16)).toNumber()] = new BigNumber(-212);
        _nw395[(new BigNumber(17)).toNumber()] = (_2414_v39).f11;
        _nw395[(new BigNumber(18)).toNumber()] = (_this).f11;
        _nw395[(new BigNumber(19)).toNumber()] = ((_2414_v39).f11).multipliedBy(new BigNumber(-92));
        _nw395[(new BigNumber(20)).toNumber()] = (new BigNumber(-196)).plus(p1);
        _nw395[(new BigNumber(21)).toNumber()] = (_this).f11;
        _2418_v43 = _nw395;
        _2418_v43 = _2418_v43;
        let _2419_v44;
        let _2420_v45;
        let _out41;
        let _out42;
        let _outcollector14 = (_2414_v39).m1(p1, globalState);
        _out41 = _outcollector14[0];
        _out42 = _outcollector14[1];
        _2419_v44 = _out41;
        _2420_v45 = _out42;
        let _rhs410 = new BigNumber(-437);
        let _rhs411 = _2362_v1;
        let _lhs349 = globalState;
        _lhs349.f1 = _rhs410;
        _2362_v1 = _rhs411;
        (globalState).f1 = (new BigNumber(939)).multipliedBy(p2);
        (globalState).f1 = new BigNumber((_dafny.Seq.of(p3)).length);
      } else {
        let _2421_v47;
        _2421_v47 = _dafny.MultiSet.fromElements(_2364_v3);
        let _2422_v48;
        _2422_v48 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((function () {
          let _coll90 = new _dafny.Map();
          for (const _compr_90 of (_2421_v47).Elements) {
            let _2423_v46 = _compr_90;
            if ((_2421_v47).contains(_2423_v46)) {
              _coll90.push([_2423_v46,false]);
            }
          }
          return _coll90;
        }()).length));
        let _2424_v49;
        _2424_v49 = _dafny.Seq.of(new BigNumber((_2422_v48).length));
        (globalState).f1 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).fm8(_module.__default.fm20(_2361_v0, new BigNumber((_2424_v49).length), p0, globalState), globalState), p1)))).multipliedBy(p1);
        (globalState).f1 = (_this).f11;
        let _2425_v50;
        _2425_v50 = _module.D2.create_DC4(_module.__default.fm2((_this).f11, (_dafny.ZERO).minus(p1), globalState), p1, p2);
        let _2426_v51;
        let _nw396 = new _module.C6();
        _nw396.__ctor(_2425_v50, (_module.__default.fm41(globalState)).Union((_this).f12), p1);
        _2426_v51 = _nw396;
        let _2427_v52;
        _2427_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2422_v48);
        let _2428_v53;
        _2428_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p3,(_2422_v48).update(p0, p1))).Merge(_2427_v52)).length),_2426_v51);
        _2426_v51 = (((_2428_v53).contains((_this).f11)) ? ((_2428_v53).get((_this).f11)) : (_2426_v51));
        let _2429_v54;
        let _nw397 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _2429_v54 = _nw397;
        let _2430_v55;
        _2430_v55 = _dafny.Map.Empty.slice().updateUnsafe(_2429_v54,p0);
        let _2431_v56;
        _2431_v56 = _module.D10.create_DC20(_2430_v55);
        let _pat_let_tv55 = _2430_v55;
        let _pat_let_tv56 = _2430_v55;
        let _source33 = function (_pat_let60_0) {
          return function (_2434_dt__update__tmp_h1) {
            return function (_pat_let63_0) {
              return function (_2435_dt__update_hcf27_h1) {
                return _module.D10.create_DC20(_2435_dt__update_hcf27_h1);
              }(_pat_let63_0);
            }(_pat_let_tv56);
          }(_pat_let60_0);
        }(function (_pat_let61_0) {
          return function (_2432_dt__update__tmp_h0) {
            return function (_pat_let62_0) {
              return function (_2433_dt__update_hcf27_h0) {
                return _module.D10.create_DC20(_2433_dt__update_hcf27_h0);
              }(_pat_let62_0);
            }(_pat_let_tv55);
          }(_pat_let61_0);
        }(_2431_v56));
        if (_source33.is_DC21) {
          let _2436___mcc_h0 = (_source33).cf28;
          let _2437___mcc_h1 = (_source33).cf29;
          let _2438_cf29 = _2437___mcc_h1;
          let _2439_cf28 = _2436___mcc_h0;
          (globalState).f1 = new BigNumber(178);
          let _2440_v57;
          let _nw398 = Array((new BigNumber(9)).toNumber()).fill(false);
          _2440_v57 = _nw398;
          let _index343 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_2440_v57).length));
          (_2440_v57)[_index343] = (new BigNumber(532)).isLessThan((_this).f11);
          let _index344 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_2440_v57).length));
          (_2440_v57)[_index344] = !_dafny.areEqual(_2364_v3, _dafny.Seq.Concat(_2364_v3, _2364_v3));
          (globalState).f6 = _2439_cf28;
          let _index345 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_2429_v54).length));
          (_2429_v54)[_index345] = (_this).f11;
          let _2441_v58;
          let _nw399 = Array((new BigNumber(27)).toNumber());
          _nw399[(_dafny.ZERO).toNumber()] = _2440_v57;
          _nw399[(_dafny.ONE).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(2)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(3)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(4)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(5)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(6)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(7)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(8)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(9)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(10)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(11)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(12)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(13)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(14)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(15)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(16)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(17)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(18)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(19)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(20)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(21)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(22)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(23)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(24)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(25)).toNumber()] = _2440_v57;
          _nw399[(new BigNumber(26)).toNumber()] = _2440_v57;
          _2441_v58 = _nw399;
          let _index346 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_2429_v54).length));
          let _rhs412 = p2;
          let _rhs413 = _2441_v58;
          let _rhs414 = _2429_v54;
          let _lhs350 = _2429_v54;
          let _lhs351 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_2429_v54).length));
          _lhs350[_lhs351] = _rhs412;
          _2441_v58 = _rhs413;
          _2429_v54 = _rhs414;
        } else {
          let _2442___mcc_h2 = (_source33).cf27;
          let _2443_cf27 = _2442___mcc_h2;
          let _2444_v59;
          _2444_v59 = _dafny.MultiSet.fromElements((_this).f11, (_this).f11);
          let _2445_v60;
          let _nw400 = new _module.C4();
          _nw400.__ctor(_2444_v59, (_this).f13, (_this).f12, new BigNumber(587));
          _2445_v60 = _nw400;
          let _2446_v61;
          _2446_v61 = _dafny.Seq.of(_2445_v60, _2445_v60);
          let _2447_v62;
          _2447_v62 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_2446_v61).length)).multipliedBy(new BigNumber((_2364_v3).length)),((_2424_v49)[_module.__default.safeIndex(p1, new BigNumber((_2424_v49).length))]).plus(p1));
          _2447_v62 = (_2447_v62).update(p2, p2);
          let _2448_v63;
          _2448_v63 = _dafny.Map.Empty.slice().updateUnsafe((_2445_v60).f16,_2447_v62);
          let _2449_v64;
          let _nw401 = Array((new BigNumber(29)).toNumber());
          _nw401[(_dafny.ZERO).toNumber()] = (_2445_v60).f17;
          _nw401[(_dafny.ONE).toNumber()] = p3;
          _nw401[(new BigNumber(2)).toNumber()] = false;
          _nw401[(new BigNumber(3)).toNumber()] = (p2).isLessThanOrEqualTo(new BigNumber((_2422_v48).length));
          _nw401[(new BigNumber(4)).toNumber()] = true;
          _nw401[(new BigNumber(5)).toNumber()] = p3;
          _nw401[(new BigNumber(6)).toNumber()] = (_2445_v60).f17;
          _nw401[(new BigNumber(7)).toNumber()] = (_2445_v60).f17;
          _nw401[(new BigNumber(8)).toNumber()] = !((p2).isLessThanOrEqualTo((_this).f11));
          _nw401[(new BigNumber(9)).toNumber()] = (p1).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length))));
          _nw401[(new BigNumber(10)).toNumber()] = !((_2445_v60).f17);
          _nw401[(new BigNumber(11)).toNumber()] = _dafny.Seq.IsPrefixOf(_2364_v3, _2364_v3);
          _nw401[(new BigNumber(12)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_2364_v3, _dafny.Seq.update(_2364_v3, _module.__default.safeIndex(p1, new BigNumber((_2364_v3).length)), _2362_v1));
          _nw401[(new BigNumber(13)).toNumber()] = !(p3) || (p0);
          _nw401[(new BigNumber(14)).toNumber()] = p0;
          _nw401[(new BigNumber(15)).toNumber()] = p3;
          _nw401[(new BigNumber(16)).toNumber()] = p0;
          _nw401[(new BigNumber(17)).toNumber()] = (_2445_v60).fm4(p2, _2448_v63, p2, p3, globalState);
          _nw401[(new BigNumber(18)).toNumber()] = !(p1).isEqualTo(new BigNumber(674));
          _nw401[(new BigNumber(19)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("xwylrsvr"), _dafny.Seq.UnicodeFromString("rk"));
          _nw401[(new BigNumber(20)).toNumber()] = false;
          _nw401[(new BigNumber(21)).toNumber()] = p3;
          _nw401[(new BigNumber(22)).toNumber()] = ((p3) ? (p3) : ((_this).f13));
          _nw401[(new BigNumber(23)).toNumber()] = (_2445_v60).f17;
          _nw401[(new BigNumber(24)).toNumber()] = !((_this).f11).isEqualTo((_dafny.ZERO).minus((_this).f11));
          _nw401[(new BigNumber(25)).toNumber()] = (_this).f13;
          _nw401[(new BigNumber(26)).toNumber()] = p3;
          _nw401[(new BigNumber(27)).toNumber()] = ((_this).f13) === ((_2445_v60).f17);
          _nw401[(new BigNumber(28)).toNumber()] = (_this).f13;
          _2449_v64 = _nw401;
          let _index347 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2449_v64).length));
          (_2449_v64)[_index347] = p0;
          let _index348 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2449_v64).length));
          (_2449_v64)[_index348] = (p3) || ((_2445_v60).f17);
          (globalState).f1 = p2;
          let _2450_v65;
          let _nw402 = new _module.C3();
          _nw402.__ctor(p2, (_dafny.ZERO).minus(p1), p2, (_this).f12);
          _2450_v65 = _nw402;
          _2450_v65 = _2450_v65;
        }
        let _2451_v66;
        let _nw403 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Set.Empty);
        _2451_v66 = _nw403;
        let _index349 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_2451_v66).length));
        (_2451_v66)[_index349] = _dafny.Set.fromElements(p0, true, p0);
        let _index350 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_2451_v66).length));
        (_2451_v66)[_index350] = (_this).f12;
      }
      _2362_v1 = _2362_v1;
      (globalState).f6 = !(p0);
      r0 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), ((_2452_v1) => function (_2453_i7) {
        return _2452_v1;
      })(_2362_v1)), _2364_v3);
      return r0;
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
