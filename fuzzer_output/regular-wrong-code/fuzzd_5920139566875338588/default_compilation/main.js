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
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("clrpyqdpw")).length),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(898)), function (_0_i0) {
        return new BigNumber((_dafny.Seq.of(new BigNumber(-165), new BigNumber((_dafny.Set.fromElements(false, false)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),false)).length), new BigNumber(590))).length);
      })).length),true))).Merge(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("qjcdhou")).length))).Elements) {
          let _1_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("qjcdhou")).length)), _1_v0)) {
            _coll0.push([(_1_v0).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),true)).length)),!(true)]);
          }
        }
        return _coll0;
      }());
    };
    static fm1(p0, globalState) {
      if (false) {
        return (_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_module.D16.create_DC45((_dafny.ZERO).minus(new BigNumber(-80)), new _dafny.CodePoint('o'.codePointAt(0)), new BigNumber(808), true)).dtor_cf79, new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)))).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), function (_2_i0) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        })).length))).length)).multipliedBy(new BigNumber((_dafny.Seq.of(false, !(true))).length)));
      } else {
        return (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(781)), function (_3_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        })).length)).minus(new BigNumber(56));
      }
    };
    static fm2(globalState) {
      return !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("grkmmpyo"), _dafny.Seq.UnicodeFromString("holvusflx")), _dafny.Seq.UnicodeFromString("gl")));
    };
    static fm12(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-895)), function (_4_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xvso")).length),new BigNumber((_dafny.Seq.of(true)).length))).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(251)), function (_5_i1) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length), new BigNumber(-624))), _dafny.Seq.of(new BigNumber(789), new BigNumber(253)));
    };
    static fm13(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, !(true)), _dafny.Seq.of(true)), _dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(!(true))));
    };
    static fm14(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(556))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(false)).length))));
    };
    static fm18(p0, p1, globalState) {
      return _module.D0.create_DC0(!(!(true)));
    };
    static fm19(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false)), _dafny.Seq.Concat(_dafny.Seq.of(true, !(true)), _dafny.Seq.of(false, true, true)));
    };
    static fm20(p0, globalState) {
      let _source0 = _module.D13.create_DC36(function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of (_dafny.Seq.of(_module.D7.create_DC16(true, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), function (_6_i0) {
  return new _dafny.CodePoint('u'.codePointAt(0));
})).length)), !(!(false))), _module.D7.create_DC16(true, new BigNumber(-288), false))).Elements) {
    let _7_v0 = _compr_1;
    if (_dafny.Seq.contains(_dafny.Seq.of(_module.D7.create_DC16(true, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), function (_6_i0) {
  return new _dafny.CodePoint('u'.codePointAt(0));
})).length)), !(!(false))), _module.D7.create_DC16(true, new BigNumber(-288), false)), _7_v0)) {
      _coll1.push([_7_v0,_dafny.Seq.UnicodeFromString("hsvfptg")]);
    }
  }
  return _coll1;
}());
      if (_source0.is_DC37) {
        let _8___mcc_h0 = (_source0).cf65;
        let _9___mcc_h1 = (_source0).cf66;
        let _10_cf66 = _9___mcc_h1;
        let _11_cf65 = _8___mcc_h0;
        return _dafny.MultiSet.fromElements(_dafny.Seq.of(_10_cf66, _10_cf66), _dafny.Seq.of(_10_cf66, _11_cf65, _10_cf66), _dafny.Seq.of(_11_cf65, _11_cf65, _11_cf65, _10_cf66), _dafny.Seq.of(_11_cf65));
      } else {
        let _12___mcc_h2 = (_source0).cf64;
        let _13_cf64 = _12___mcc_h2;
        return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-526)), function (_14_i1) {
          return _dafny.Seq.of(false);
        }), _dafny.Seq.of(_dafny.Seq.of(!(true), !(true)))));
      }
    };
    static fm21(globalState) {
      return _dafny.Seq.UnicodeFromString("nqwtxsk");
    };
    static fm22(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-625)), function (_15_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("cqfqwylgv")).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-177), new BigNumber((_dafny.Seq.of(new BigNumber(-425))).length))).cardinality())), _dafny.Seq.of(new BigNumber(-684))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-76)), function (_16_i1) {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)))).length));
      }));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      if ((false) && (true)) {
        return _dafny.Seq.of(new BigNumber(-383), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("bfg")).length),new _dafny.CodePoint('a'.codePointAt(0)))).length));
      } else {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("oi")).length), (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(true, !(false)))).length)))), new BigNumber(-560), new BigNumber(92), new BigNumber(248));
      }
    };
    static fm26(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length),_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("oqcb"), _dafny.Seq.UnicodeFromString("qswfrnph"))).length),new BigNumber(-964))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(515),_dafny.Set.fromElements(new BigNumber(-634))));
    };
    static fm27(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(492),_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length)))).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(361), new BigNumber(313))) {
          let _17_v0 = _compr_2;
          if (((new BigNumber(361)).isLessThanOrEqualTo(_17_v0)) && ((_17_v0).isLessThan(new BigNumber(313)))) {
            _coll2.push([(_17_v0).multipliedBy(new BigNumber(470)),_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-602)), new BigNumber((_dafny.Seq.UnicodeFromString("hbheb")).length))]);
          }
        }
        return _coll2;
      }())).Merge(((true) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(772))).length)))) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(282),function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(699), new BigNumber(68))) {
          let _18_v1 = _compr_3;
          if (((new BigNumber(699)).isLessThanOrEqualTo(_18_v1)) && ((_18_v1).isLessThan(new BigNumber(68)))) {
            _coll3.add((_18_v1).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber(670), new BigNumber(565))).length)));
          }
        }
        return _coll3;
      }()))));
    };
    static fm28(p0, p1, p2, globalState) {
      if (((false) ? (false) : (!(!(false))))) {
        return function () {
          let _coll4 = new _dafny.Set();
          for (const _compr_4 of _dafny.IntegerRange(new BigNumber(768), new BigNumber(29))) {
            let _19_v0 = _compr_4;
            if (((new BigNumber(768)).isLessThanOrEqualTo(_19_v0)) && ((_19_v0).isLessThan(new BigNumber(29)))) {
              _coll4.add(_module.__default.safeDivisionInt(_19_v0, new BigNumber(801)));
            }
          }
          return _coll4;
        }();
      } else {
        return (_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(115)))).cardinality())))).Intersect(_dafny.Set.fromElements(new BigNumber(785), new BigNumber((_dafny.Set.fromElements(new BigNumber(378), new BigNumber((function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(678), new BigNumber(-66))) {
            let _20_v1 = _compr_5;
            if (((new BigNumber(678)).isLessThanOrEqualTo(_20_v1)) && ((_20_v1).isLessThan(new BigNumber(-66)))) {
              _coll5.push([(_20_v1).minus(new BigNumber(325)),true]);
            }
          }
          return _coll5;
        }()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.UnicodeFromString("nb"))).length), new BigNumber((_dafny.Seq.UnicodeFromString("kokftkjn")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(784)), function (_21_i0) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        })).length),_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC34(function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(75), new BigNumber(695))) {
    let _22_v2 = _compr_6;
    if (((new BigNumber(75)).isLessThanOrEqualTo(_22_v2)) && ((_22_v2).isLessThan(new BigNumber(695)))) {
      _coll6.push([(_22_v2).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(654)), function (_23_i1) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)]);
    }
  }
  return _coll6;
}()),true))).length))).length), new BigNumber(709), new BigNumber(124), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("pmilbnjok")).length))).length)));
      }
    };
    static fm29(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D9.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(794),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(304),function () {
  let _coll7 = new _dafny.Set();
  for (const _compr_7 of (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("mphpqtpee")).length), new BigNumber(588))).Elements) {
    let _24_v0 = _compr_7;
    if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("mphpqtpee")).length), new BigNumber(588))).contains(_24_v0)) {
      _coll7.add((_24_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ymib")).length))));
    }
  }
  return _coll7;
}())), _dafny.Seq.of(true));
      if (_source1.is_DC23) {
        let _25___mcc_h0 = (_source1).cf37;
        let _26___mcc_h1 = (_source1).cf38;
        let _27___mcc_h2 = (_source1).cf39;
        let _28___mcc_h3 = (_source1).cf40;
        let _29___mcc_h4 = (_source1).cf41;
        let _30_cf41 = _29___mcc_h4;
        let _31_cf40 = _28___mcc_h3;
        let _32_cf39 = _27___mcc_h2;
        let _33_cf38 = _26___mcc_h1;
        let _34_cf37 = _25___mcc_h0;
        return _30_cf41;
      } else if (_source1.is_DC24) {
        let _35___mcc_h5 = (_source1).cf42;
        let _36___mcc_h6 = (_source1).cf43;
        let _37_cf43 = _36___mcc_h6;
        let _38_cf42 = _35___mcc_h5;
        return new _dafny.CodePoint('t'.codePointAt(0));
      } else if (_source1.is_DC22) {
        let _39___mcc_h7 = (_source1).cf36;
        let _40_cf36 = _39___mcc_h7;
        if (true) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }
      } else {
        let _41___mcc_h8 = (_source1).cf44;
        let _42_cf44 = _41___mcc_h8;
        return new _dafny.CodePoint('a'.codePointAt(0));
      }
    };
    static fm30(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true, false))).Difference(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true, false, false)))).Intersect(_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true))));
    };
    static fm31(p0, p1, globalState) {
      return _dafny.Set.fromElements(!(true), ((true) ? (!(true)) : (true)), !(false) || (true), (new BigNumber(494)).isEqualTo((_dafny.ZERO).minus(new BigNumber(-851))), !(false) || (!(false)));
    };
    static fm32(p0, p1, p2, globalState) {
      return _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(43),false))).length), (new BigNumber(-552)).multipliedBy(new BigNumber((function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(258), new BigNumber(368))) {
          let _43_v0 = _compr_8;
          if (((new BigNumber(258)).isLessThanOrEqualTo(_43_v0)) && ((_43_v0).isLessThan(new BigNumber(368)))) {
            _coll8.push([(_43_v0).minus(new BigNumber(81)),false]);
          }
        }
        return _coll8;
      }()).length)), _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("kchk")).length), new BigNumber(731)));
    };
    static fm33(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("yoy")).length))).length),new BigNumber(229))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(223),new BigNumber(388))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true, false)).length),new BigNumber(642))));
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(!(!(!(false))))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
    };
    static fm35(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC0(!(false))), _dafny.Seq.of(_module.D0.create_DC0(false), _module.D0.create_DC0(false), _module.D0.create_DC0(true)));
    };
    static fm36(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-760))).length))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),new BigNumber(-685))));
    };
    static fm37(p0, p1, globalState) {
      return _module.D10.create_DC28(((true) ? (new BigNumber(-150)) : (new BigNumber(41))));
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(826),_dafny.Seq.UnicodeFromString("fghwqouji"))).Merge(function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_dafny.Seq.of(new BigNumber(757), new BigNumber(853), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("krtxro")).length)))).Elements) {
          let _44_v0 = _compr_9;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(757), new BigNumber(853), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("krtxro")).length))), _44_v0)) {
            _coll9.push([_module.__default.safeModuloInt(_44_v0, new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_module.D13.create_DC36(_dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC16(false, new BigNumber((_dafny.Seq.UnicodeFromString("rsiiwqtrg")).length), false),_dafny.Seq.UnicodeFromString("k"))), _module.D13.create_DC36(function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of (_dafny.Seq.of(_module.D7.create_DC16(false, new BigNumber(-159), true), _module.D7.create_DC16(false, new BigNumber(-950), false), _module.D7.create_DC16(!(false), new BigNumber(953), false), _module.D7.create_DC16(false, new BigNumber((_dafny.Seq.UnicodeFromString("jbwndisey")).length), true), _module.D7.create_DC16(false, new BigNumber(344), true))).Elements) {
    let _45_v1 = _compr_10;
    if (_dafny.Seq.contains(_dafny.Seq.of(_module.D7.create_DC16(false, new BigNumber(-159), true), _module.D7.create_DC16(false, new BigNumber(-950), false), _module.D7.create_DC16(!(false), new BigNumber(953), false), _module.D7.create_DC16(false, new BigNumber((_dafny.Seq.UnicodeFromString("jbwndisey")).length), true), _module.D7.create_DC16(false, new BigNumber(344), true)), _45_v1)) {
      _coll10.push([_45_v1,_dafny.Seq.UnicodeFromString("tmge")]);
    }
  }
  return _coll10;
}()))).length))))).cardinality())),_dafny.Seq.UnicodeFromString("ubeopiw")]);
          }
        }
        return _coll9;
      }());
    };
    static fm39(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!(true), true, false)).length))).cardinality()),true),false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(144),false),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-847)), function (_46_i0) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        })).length))).length))).Elements) {
          let _47_v0 = _compr_11;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-847)), function (_46_i0) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          })).length))).length)), _47_v0)) {
            _coll11.push([(_47_v0).multipliedBy(new BigNumber(954)),false]);
          }
        }
        return _coll11;
      }(),false)));
    };
    static fm40(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D9.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, false)).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(334),_dafny.Set.fromElements(new BigNumber(345)))), _dafny.Seq.of(true)), _module.D9.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(80),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),_dafny.Set.fromElements(new BigNumber(455), new BigNumber(528), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(747)), function (_48_i0) {
  return new _dafny.CodePoint('d'.codePointAt(0));
})).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(563))).length)))), _dafny.Seq.of(false, true))),new BigNumber(488))).Merge(function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(687)), function (_49_i1) {
          return _module.D9.create_DC24(function () {
  let _coll13 = new _dafny.Map();
  for (const _compr_13 of _dafny.IntegerRange(new BigNumber(862), new BigNumber(629))) {
    let _50_v1 = _compr_13;
    if (((new BigNumber(862)).isLessThanOrEqualTo(_50_v1)) && ((_50_v1).isLessThan(new BigNumber(629)))) {
      _coll13.push([_module.__default.safeModuloInt(_50_v1, new BigNumber(-270)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(756),_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-241),new BigNumber(-358))).length)))]);
    }
  }
  return _coll13;
}(), _dafny.Seq.of(true));
        }))).Elements) {
          let _51_v0 = _compr_12;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(687)), function (_49_i1) {
            return _module.D9.create_DC24(function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of _dafny.IntegerRange(new BigNumber(862), new BigNumber(629))) {
    let _50_v1 = _compr_14;
    if (((new BigNumber(862)).isLessThanOrEqualTo(_50_v1)) && ((_50_v1).isLessThan(new BigNumber(629)))) {
      _coll14.push([_module.__default.safeModuloInt(_50_v1, new BigNumber(-270)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(756),_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-241),new BigNumber(-358))).length)))]);
    }
  }
  return _coll14;
}(), _dafny.Seq.of(true));
          })), _51_v0)) {
            _coll12.push([_51_v0,new BigNumber(520)]);
          }
        }
        return _coll12;
      }());
    };
    static fm41(p0, globalState) {
      return _dafny.Seq.of((_dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-901)), new BigNumber(-381), new BigNumber((_dafny.Seq.of(true)).length)))).Difference(_dafny.MultiSet.fromElements(new BigNumber(684), new BigNumber(946))), _dafny.MultiSet.fromElements(new BigNumber(((_module.D8.create_DC21(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-667)), function (_52_i0) {
  return new _dafny.CodePoint('p'.codePointAt(0));
}), _dafny.Set.fromElements(false, true))).dtor_cf34).length), (_module.D11.create_DC32(false, new BigNumber(392), true)).dtor_cf60, new BigNumber((_dafny.Set.fromElements(true, true)).length), new BigNumber(-584), new BigNumber((_dafny.Seq.of(!(true))).length)), ((false) ? (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(112)), function (_53_i1) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length)))) : (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("lkidualy")).length)))));
    };
    static fm42(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(866), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(572))).cardinality()))).length), new BigNumber(556), new BigNumber((_dafny.Seq.UnicodeFromString("ls")).length), new BigNumber(980), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length))))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(-141)));
    };
    static fm43(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(!(true)))).length)))).Union(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(699)), _dafny.MultiSet.fromElements(new BigNumber(546))))).Union(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(-645), new BigNumber(-185)), _dafny.MultiSet.fromElements(new BigNumber(-60), new BigNumber(334), new BigNumber(918)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true, !(true))).length),false)).length), new BigNumber(791))));
    };
    static fm44(globalState) {
      return _module.D13.create_DC37(!(false), false);
    };
    static fm45(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of _dafny.IntegerRange(new BigNumber(809), new BigNumber(705))) {
          let _54_v0 = _compr_15;
          if (((new BigNumber(809)).isLessThanOrEqualTo(_54_v0)) && ((_54_v0).isLessThan(new BigNumber(705)))) {
            _coll15.push([_module.__default.safeDivisionInt(_54_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)),true]);
          }
        }
        return _coll15;
      }()), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-192),!(false)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(745),false))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("afelo")).length),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(661))).length),false)), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(769),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(639),false))));
    };
    static fm46(p0, p1, p2, p3, globalState) {
      let _source2 = _module.D8.create_DC19(_dafny.Set.fromElements(true, false));
      if (_source2.is_DC20) {
        let _55___mcc_h0 = (_source2).cf31;
        let _56___mcc_h1 = (_source2).cf32;
        let _57___mcc_h2 = (_source2).cf33;
        let _58_cf33 = _57___mcc_h2;
        let _59_cf32 = _56___mcc_h1;
        let _60_cf31 = _55___mcc_h0;
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)));
      } else if (_source2.is_DC21) {
        let _61___mcc_h3 = (_source2).cf34;
        let _62___mcc_h4 = (_source2).cf35;
        let _63_cf35 = _62___mcc_h4;
        let _64_cf34 = _61___mcc_h3;
        return (_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0)))).Union(_dafny.MultiSet.fromElements((_64_cf34)[_module.__default.safeIndex(new BigNumber(856), new BigNumber((_64_cf34).length))]));
      } else {
        let _65___mcc_h5 = (_source2).cf30;
        let _66_cf30 = _65___mcc_h5;
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0)));
      }
    };
    static fm47(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(935)), function (_67_i0) {
        return new BigNumber(811);
      })).length),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-766),!(false))),new BigNumber(226));
    };
    static fm48(p0, globalState) {
      return _module.D3.create_DC7(new BigNumber(((function () {
  let _coll16 = new _dafny.Map();
  for (const _compr_16 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality()))).length),!(true))).length),true), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("yma")).length)),true))).Elements) {
    let _68_v0 = _compr_16;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality()))).length),!(true))).length),true), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("yma")).length)),true)), _68_v0)) {
      _coll16.push([_68_v0,!(false)]);
    }
  }
  return _coll16;
}()).Merge(function () {
  let _coll17 = new _dafny.Map();
  for (const _compr_17 of (_dafny.MultiSet.fromElements(function () {
    let _coll18 = new _dafny.Map();
    for (const _compr_18 of _dafny.IntegerRange(new BigNumber(728), new BigNumber(-532))) {
      let _69_v2 = _compr_18;
      if (((new BigNumber(728)).isLessThanOrEqualTo(_69_v2)) && ((_69_v2).isLessThan(new BigNumber(-532)))) {
        _coll18.push([(_69_v2).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)))).length)),false]);
      }
    }
    return _coll18;
  }())).Elements) {
    let _70_v1 = _compr_17;
    if ((_dafny.MultiSet.fromElements(function () {
      let _coll19 = new _dafny.Map();
      for (const _compr_19 of _dafny.IntegerRange(new BigNumber(728), new BigNumber(-532))) {
        let _69_v2 = _compr_19;
        if (((new BigNumber(728)).isLessThanOrEqualTo(_69_v2)) && ((_69_v2).isLessThan(new BigNumber(-532)))) {
          _coll19.push([(_69_v2).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)))).length)),false]);
        }
      }
      return _coll19;
    }())).contains(_70_v1)) {
      _coll17.push([_70_v1,true]);
    }
  }
  return _coll17;
}())).length));
    };
    static fm49(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-586))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(492)));
    };
    static fm50(p0, p1, globalState) {
      if (!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("rqld"), _dafny.Seq.UnicodeFromString("futiieesj")))) {
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(true,!(true)), _dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(false,true)), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
      } else {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), function (_71_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe(true,true);
        });
      }
    };
    static fm51(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(611),_dafny.MultiSet.fromElements(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(true),new _dafny.CodePoint('a'.codePointAt(0))),new BigNumber(187))).length),_dafny.MultiSet.fromElements(false, true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),_dafny.MultiSet.fromElements(!(true))));
    };
    static fm52(p0, p1, globalState) {
      return _module.D3.create_DC8(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qfdjxwbg")).length),new BigNumber((_dafny.Seq.of(!(false), !(false))).length))).length),true),_dafny.Seq.of(new BigNumber(370)))).length), new BigNumber(181)), _dafny.Seq.UnicodeFromString("thgik"), (_dafny.MultiSet.fromElements(true)).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.of(false))), new _dafny.CodePoint('u'.codePointAt(0)));
    };
    static fm53(p0, p1, globalState) {
      return _module.D15.create_DC43(!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("umnb"), _dafny.Seq.UnicodeFromString("sgs"))), (false) && (false), true);
    };
    static fm54(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(369),new _dafny.CodePoint('w'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(918),new _dafny.CodePoint('l'.codePointAt(0))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(598),new _dafny.CodePoint('s'.codePointAt(0))));
    };
    static fm55(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(_module.D11.create_DC31(_dafny.Seq.of(true), new BigNumber(975), new BigNumber((_dafny.MultiSet.fromElements(_module.D8.create_DC21(_dafny.Seq.UnicodeFromString("rmwvibhcx"), _dafny.Set.fromElements(false)))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("dddghnql")).length), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(false))).length)), new BigNumber(760))).length)));
    };
    static fm56(p0, p1, globalState) {
      if (false) {
        return _module.D0.create_DC2(_module.D0.create_DC0(false));
      } else {
        return _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC0(!(true))));
      }
    };
    static fm57(p0, p1, globalState) {
      return _module.D14.create_DC40((new BigNumber(64)).isEqualTo(new BigNumber((_dafny.Seq.of(new BigNumber(849))).length)));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let _72_v0;
      let _init0 = function (_73_i0) {
        return true;
      };
      let _nw0 = Array((new BigNumber(23)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _72_v0 = _nw0;
      let _74_v1;
      _74_v1 = _dafny.Set.fromElements(_72_v0);
      let _75_v2;
      _75_v2 = _module.D3.create_DC6(p2);
      let _76_v3;
      _76_v3 = _dafny.Seq.of(_75_v2);
      let _77_v4;
      _77_v4 = true;
      let _78_v5;
      let _nw1 = new _module.C0();
      _nw1.__ctor(_76_v3, _77_v4);
      _78_v5 = _nw1;
      let _79_v6;
      let _nw2 = Array((new BigNumber(15)).toNumber());
      _nw2[(_dafny.ZERO).toNumber()] = _78_v5;
      _nw2[(_dafny.ONE).toNumber()] = _78_v5;
      _nw2[(new BigNumber(2)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(3)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(4)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(5)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(6)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(7)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(8)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(9)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(10)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(11)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(12)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(13)).toNumber()] = _78_v5;
      _nw2[(new BigNumber(14)).toNumber()] = _78_v5;
      _79_v6 = _nw2;
      let _80_v7;
      _80_v7 = _dafny.Seq.of(_79_v6, _79_v6, _79_v6);
      let _81_v8;
      _81_v8 = _module.D27.create_DC61(_80_v7);
      let _pat_let_tv0 = _80_v7;
      let _source3 = (((_74_v1).IsSubsetOf(_74_v1)) ? (function (_pat_let0_0) {
        return function (_82_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_83_dt__update_hcf101_h0) {
              return _module.D27.create_DC61(_83_dt__update_hcf101_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_module.D27.create_DC61(_80_v7))) : (_81_v8));
      if (_source3.is_DC62) {
        let _84___mcc_h0 = (_source3).cf102;
        let _85___mcc_h1 = (_source3).cf103;
        let _86___mcc_h2 = (_source3).cf104;
        let _87___mcc_h3 = (_source3).cf105;
        let _88_cf105 = _87___mcc_h3;
        let _89_cf104 = _86___mcc_h2;
        let _90_cf103 = _85___mcc_h1;
        let _91_cf102 = _84___mcc_h0;
        let _92_v9;
        _92_v9 = new BigNumber(398);
        let _rhs0 = _92_v9;
        let _rhs1 = _90_cf103;
        let _rhs2 = (_78_v5).f23;
        let _rhs3 = (_dafny.ZERO).minus(_92_v9);
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        _lhs0.f6 = _rhs0;
        _89_cf104 = _rhs1;
        _91_cf102 = _rhs2;
        _lhs1.f6 = _rhs3;
        let _93_v10;
        _93_v10 = _dafny.Seq.of(_92_v9);
        let _94_v11;
        let _nw3 = Array((new BigNumber(22)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = _93_v10;
        _nw3[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_92_v9);
        _nw3[(new BigNumber(2)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(3)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(4)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(5)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_92_v9, _92_v9, _92_v9, (_dafny.ZERO).minus(_92_v9), _92_v9);
        _nw3[(new BigNumber(7)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(8)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(9)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(_92_v9);
        _nw3[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-471)), ((_95_v9) => function (_96_i1) {
          return _95_v9;
        })(_92_v9));
        _nw3[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_93_v10, _module.__default.safeIndex(_92_v9, new BigNumber((_93_v10).length)), _92_v9);
        _nw3[(new BigNumber(13)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(14)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(15)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(16)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(17)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(18)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(19)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(20)).toNumber()] = _93_v10;
        _nw3[(new BigNumber(21)).toNumber()] = _93_v10;
        _94_v11 = _nw3;
        let _97_v12;
        _97_v12 = _module.D14.create_DC38(_94_v11);
        let _98_v13;
        _98_v13 = _dafny.Seq.of(_module.D14.create_DC41(_97_v12));
        _92_v9 = new BigNumber((_98_v13).length);
        let _99_v14;
        _99_v14 = _dafny.Set.fromElements(p3);
        let _100_v15;
        _100_v15 = new _dafny.CodePoint('j'.codePointAt(0));
        let _101_v16;
        _101_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(p3, _module.__default.safeIndex(_92_v9, new BigNumber((p3).length)), _100_v15),_92_v9);
        let _102_v18;
        _102_v18 = _dafny.Set.fromElements(_88_cf105);
        let _103_v19;
        _103_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_102_v18).length),p3);
        let _104_v20;
        _104_v20 = _dafny.Seq.of(_88_cf105);
        let _nw4 = Array((new BigNumber(11)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = (_78_v5).f23;
        _nw4[(_dafny.ONE).toNumber()] = _89_cf104;
        _nw4[(new BigNumber(2)).toNumber()] = (_78_v5).f23;
        _nw4[(new BigNumber(3)).toNumber()] = _88_cf105;
        _nw4[(new BigNumber(4)).toNumber()] = (_78_v5).f23;
        _nw4[(new BigNumber(5)).toNumber()] = _88_cf105;
        _nw4[(new BigNumber(6)).toNumber()] = !(!((function () {
          let _coll20 = new _dafny.Set();
          for (const _compr_20 of (_101_v16).Keys.Elements) {
            let _105_v17 = _compr_20;
            if ((_101_v16).contains(_105_v17)) {
              _coll20.add(_105_v17);
            }
          }
          return _coll20;
        }()).IsSubsetOf(_99_v14)));
        _nw4[(new BigNumber(7)).toNumber()] = _dafny.Seq.IsProperPrefixOf(p3, p3);
        _nw4[(new BigNumber(8)).toNumber()] = (_99_v14).IsSubsetOf(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ctqonb"), (((_103_v19).contains(new BigNumber(831))) ? ((_103_v19).get(new BigNumber(831))) : (p3)), p3, p3, _dafny.Seq.UnicodeFromString("iqkhgippl")));
        _nw4[(new BigNumber(9)).toNumber()] = _77_v4;
        _nw4[(new BigNumber(10)).toNumber()] = (_78_v5).fm17(new BigNumber((p3).length), (_module.D11.create_DC32((_104_v20)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_104_v20).length))], _92_v9, true)).dtor_cf60, _91_cf102, globalState);
        _72_v0 = _nw4;
        let _index0 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((p2).length));
        (p2)[_index0] = _92_v9;
        let _106_v22;
        _106_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_78_v5).f23, globalState),!(_88_cf105));
        let _107_v23;
        _107_v23 = _dafny.Seq.of(_106_v22);
        let _108_v24;
        let _nw5 = new _module.C2();
        _nw5.__ctor(_92_v9, _92_v9, function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of (_107_v23).Elements) {
            let _109_v21 = _compr_21;
            if (_dafny.Seq.contains(_107_v23, _109_v21)) {
              _coll21.push([_109_v21,_92_v9]);
            }
          }
          return _coll21;
        }());
        _108_v24 = _nw5;
        let _110_v25;
        _110_v25 = _dafny.MultiSet.fromElements(_108_v24, _108_v24, _108_v24);
        let _index1 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((p2).length));
        let _rhs4 = _100_v15;
        let _rhs5 = _module.__default.fm1((_110_v25).IsProperSubsetOf(_110_v25), globalState);
        let _lhs2 = p2;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((p2).length));
        _100_v15 = _rhs4;
        _lhs2[_lhs3] = _rhs5;
      } else if (_source3.is_DC63) {
        let _111___mcc_h4 = (_source3).cf106;
        let _112___mcc_h5 = (_source3).cf107;
        let _113_cf107 = _112___mcc_h5;
        let _114_cf106 = _111___mcc_h4;
        _113_cf107 = _113_cf107;
        let _115_v26;
        _115_v26 = _dafny.Set.fromElements((_78_v5).f23);
        let _index2 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_72_v0).length));
        (_72_v0)[_index2] = (_115_v26).IsProperSubsetOf(_115_v26);
        let _116_v27;
        _116_v27 = _module.D8.create_DC20((_dafny.ZERO).minus(_113_cf107), _113_cf107, _113_cf107);
        let _index3 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_72_v0).length));
        (_72_v0)[_index3] = (_module.__default.fm52(_116_v27, _113_cf107, globalState)).dtor_cf10;
        let _index4 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_79_v6).length));
        (_79_v6)[_index4] = _78_v5;
        let _index5 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_79_v6).length));
        (_79_v6)[_index5] = _78_v5;
        let _117_v28;
        let _nw6 = new _module.C0();
        _nw6.__ctor((_78_v5).f22, !((((_78_v5).f23) ? (true) : (!(_77_v4)))));
        _117_v28 = _nw6;
      } else {
        let _118___mcc_h6 = (_source3).cf101;
        let _119_cf101 = _118___mcc_h6;
        (globalState).f1 = _77_v4;
        let _120_v29;
        let _nw7 = new _module.C4();
        _nw7.__ctor();
        _120_v29 = _nw7;
        let _121_v30;
        let _nw8 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _121_v30 = _nw8;
        let _122_v31;
        _122_v31 = new _dafny.CodePoint('t'.codePointAt(0));
        let _index6 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_121_v30).length));
        (_121_v30)[_index6] = _122_v31;
        let _123_v32;
        _123_v32 = _dafny.Seq.UnicodeFromString("nctipsdh");
        let _124_v33;
        _124_v33 = new BigNumber(177);
        let _125_v34;
        let _nw9 = new _module.C5();
        _nw9.__ctor(_124_v33);
        _125_v34 = _nw9;
        let _index7 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_79_v6).length));
        (_79_v6)[_index7] = _78_v5;
        let _index8 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_121_v30).length));
        let _index9 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_79_v6).length));
        let _rhs6 = _122_v31;
        let _rhs7 = p3;
        let _rhs8 = _125_v34;
        let _rhs9 = _78_v5;
        let _lhs4 = _121_v30;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_121_v30).length));
        let _lhs6 = _79_v6;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_79_v6).length));
        _lhs4[_lhs5] = _rhs6;
        _123_v32 = _rhs7;
        _125_v34 = _rhs8;
        _lhs6[_lhs7] = _rhs9;
        let _126_v35;
        _126_v35 = _dafny.Seq.of(new BigNumber(-553));
        let _127_v36;
        _127_v36 = _dafny.Seq.of(_dafny.Seq.of(_126_v35, _126_v35));
        (globalState).f6 = _module.__default.safeModuloInt(new BigNumber(((_127_v36)[_module.__default.safeIndex((_module.D7.create_DC16(true, _124_v33, (_78_v5).f23)).dtor_cf22, new BigNumber((_127_v36).length))]).length), _124_v33);
      }
      let _128_v37;
      let _nw10 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _128_v37 = _nw10;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_128_v37).length))) {
        let _129_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_129_i2)) && ((_129_i2).isLessThan(new BigNumber((_128_v37).length))))) {
          (_128_v37)[(_129_i2)] = (_129_i2).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_77_v4)).length)),new BigNumber(-253))).length));
        }
      }
      let _130_v38;
      _130_v38 = new BigNumber(117);
      (globalState).f1 = (_78_v5).fm17(_130_v38, _130_v38, _77_v4, globalState);
      let _131_i3;
      _131_i3 = _dafny.ZERO;
      L0: {
        while (_dafny.Seq.IsProperPrefixOf(p3, p3)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_131_i3)) {
              break L0;
            }
            _131_i3 = (_131_i3).plus(_dafny.ONE);
            let _132_v39;
            _132_v39 = _dafny.Map.Empty.slice().updateUnsafe((_78_v5).f23,(_78_v5).f23);
            let _133_v40;
            _133_v40 = _module.D1.create_DC3((_132_v39).update(_77_v4, _77_v4));
            let _pat_let_tv1 = _132_v39;
            let _134_v41;
            _134_v41 = _dafny.Set.fromElements(function (_pat_let2_0) {
              return function (_135_dt__update__tmp_h1) {
                return function (_pat_let3_0) {
                  return function (_136_dt__update_hcf3_h0) {
                    return _module.D1.create_DC3(_136_dt__update_hcf3_h0);
                  }(_pat_let3_0);
                }(_pat_let_tv1);
              }(_pat_let2_0);
            }(_133_v40), _133_v40);
            let _137_v42;
            _137_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_134_v41).length),_130_v38);
            _137_v42 = (_137_v42).update(_130_v38, _130_v38);
            let _138_v43;
            _138_v43 = _dafny.Seq.UnicodeFromString("dicognbwi");
            let _139_v44;
            _139_v44 = new _dafny.CodePoint('j'.codePointAt(0));
            _138_v43 = _dafny.Seq.Concat(p3, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), function (_140_i4) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            }), p3), _module.__default.safeIndex(_130_v38, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), function (_141_i4) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            }), p3)).length)), _139_v44));
            if (_module.__default.fm2(globalState)) {
              let _index10 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_128_v37).length));
              (_128_v37)[_index10] = _module.__default.safeDivisionInt(new BigNumber((_137_v42).length), _130_v38);
              let _142_v45;
              _142_v45 = _dafny.Map.Empty.slice().updateUnsafe(_130_v38,p3);
              let _index11 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_128_v37).length));
              (_128_v37)[_index11] = (new BigNumber(((((_142_v45).contains(_130_v38)) ? ((_142_v45).get(_130_v38)) : (_dafny.Seq.update(p3, _module.__default.safeIndex(_130_v38, new BigNumber((p3).length)), new _dafny.CodePoint('m'.codePointAt(0)))))).length)).minus(_130_v38);
              (globalState).f6 = new BigNumber(-970);
              let _143_v46;
              _143_v46 = _module.D20.create_DC50((_78_v5).f23, _130_v38, _130_v38, true, false);
              (globalState).f1 = (((_128_v37)[_module.__default.safeIndex(new BigNumber(534), new BigNumber((_128_v37).length))]).plus(_130_v38)).isLessThan((_143_v46).dtor_cf87);
              let _index12 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_72_v0).length));
              (_72_v0)[_index12] = (_130_v38).isLessThanOrEqualTo(new BigNumber((p3).length));
              let _index13 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_128_v37).length));
              let _index14 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_72_v0).length));
              let _rhs10 = new BigNumber(-76);
              let _rhs11 = (p0).IsProperSubsetOf(p0);
              let _rhs12 = new BigNumber(375);
              let _lhs8 = _128_v37;
              let _lhs9 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_128_v37).length));
              let _lhs10 = _72_v0;
              let _lhs11 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_72_v0).length));
              let _lhs12 = globalState;
              _lhs8[_lhs9] = _rhs10;
              _lhs10[_lhs11] = _rhs11;
              _lhs12.f6 = _rhs12;
              (globalState).f6 = (new BigNumber(443)).minus((new BigNumber(798)).multipliedBy(new BigNumber(54)));
            } else {
              (globalState).f6 = _130_v38;
              let _144_v47;
              _144_v47 = _dafny.Map.Empty.slice().updateUnsafe(_139_v44,!((_78_v5).f23));
              let _index15 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_72_v0).length));
              (_72_v0)[_index15] = (((_144_v47).contains(_139_v44)) ? ((_144_v47).get(_139_v44)) : ((_78_v5).f23));
              let _index16 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_72_v0).length));
              (_72_v0)[_index16] = !_dafny.areEqual(_138_v43, _138_v43);
              let _145_v48;
              let _nw11 = Array((new BigNumber(11)).toNumber()).fill([]);
              _145_v48 = _nw11;
              let _146_v49;
              let _nw12 = Array((new BigNumber(27)).toNumber()).fill(_module.D10.Default());
              _146_v49 = _nw12;
              let _index17 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_145_v48).length));
              (_145_v48)[_index17] = _146_v49;
              let _index18 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_145_v48).length));
              (_145_v48)[_index18] = _146_v49;
              let _index19 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_128_v37).length));
              (_128_v37)[_index19] = (_130_v38).plus(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(_77_v4))).length));
              let _147_v50;
              _147_v50 = _dafny.Map.Empty.slice().updateUnsafe(_130_v38,(_78_v5).f23);
              let _148_v51;
              _148_v51 = _dafny.Seq.of(new BigNumber((_138_v43).length), _130_v38, new BigNumber(-261), _130_v38);
              let _index20 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_128_v37).length));
              let _rhs13 = (_module.__default.safeModuloInt(new BigNumber((_138_v43).length), _130_v38)).isEqualTo(_130_v38);
              let _rhs14 = (new BigNumber((_147_v50).length)).plus(_module.__default.safeDivisionInt(_130_v38, (_dafny.ZERO).minus((_148_v51)[_module.__default.safeIndex(_130_v38, new BigNumber((_148_v51).length))])));
              let _lhs13 = globalState;
              let _lhs14 = _128_v37;
              let _lhs15 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_128_v37).length));
              _lhs13.f1 = _rhs13;
              _lhs14[_lhs15] = _rhs14;
              let _149_v52;
              let _nw13 = Array((new BigNumber(20)).toNumber()).fill(_module.D3.Default());
              _149_v52 = _nw13;
              let _150_v53;
              _150_v53 = _dafny.Set.fromElements(_149_v52, _149_v52);
              let _151_v54;
              _151_v54 = _module.D9.create_DC23(_150_v53, p3, (_78_v5).f23, p2, new _dafny.CodePoint('i'.codePointAt(0)));
              _151_v54 = _module.D9.create_DC23(_dafny.Set.fromElements(_149_v52, _149_v52, _149_v52, _149_v52), _dafny.Seq.update(p3, _module.__default.safeIndex(_130_v38, new BigNumber((p3).length)), new _dafny.CodePoint('a'.codePointAt(0))), (_72_v0)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_72_v0).length))], p2, _139_v44);
            }
            (globalState).f1 = !(_77_v4);
          }
        }
      }
      let _hi0 = _130_v38;
      for (let _152_i5 = _130_v38; _152_i5.isLessThan(_hi0); _152_i5 = _152_i5.plus(_dafny.ONE)) {
        let _153_v55;
        _153_v55 = _dafny.Seq.of(_130_v38, _152_i5, _152_i5);
        (globalState).f6 = (_module.__default.safeModuloInt(new BigNumber((_153_v55).length), _152_i5)).minus(_module.__default.safeModuloInt(_152_i5, new BigNumber((_module.__default.fm21(globalState)).length)));
        (globalState).f1 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(742)), function (_154_i6) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }), p3);
        let _index21 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_72_v0).length));
        (_72_v0)[_index21] = (_78_v5).f23;
        let _index22 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((p2).length));
        (p2)[_index22] = _130_v38;
        let _155_v56;
        _155_v56 = _dafny.Seq.of(_78_v5, _78_v5);
        let _156_v57;
        _156_v57 = _module.D3.create_DC7(_152_i5);
        let _index23 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_72_v0).length));
        let _index24 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((p2).length));
        let _rhs15 = !((_130_v38).plus(_130_v38)).isEqualTo(new BigNumber(520));
        let _rhs16 = (_130_v38).multipliedBy(new BigNumber((_155_v56).length));
        let _rhs17 = (_156_v57).dtor_cf7;
        let _lhs16 = _72_v0;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_72_v0).length));
        let _lhs18 = p2;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((p2).length));
        let _lhs20 = globalState;
        _lhs16[_lhs17] = _rhs15;
        _lhs18[_lhs19] = _rhs16;
        _lhs20.f6 = _rhs17;
        let _index25 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((p2).length));
        (p2)[_index25] = _130_v38;
      }
      _128_v37 = p2;
      return;
    }
    static Main(__noArgsParameter) {
      let _157_v0;
      _157_v0 = true;
      let _158_v1;
      _158_v1 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_157_v0)).length));
      let _159_v2;
      _159_v2 = _dafny.Map.Empty.slice().updateUnsafe(_158_v1,false);
      let _160_v3;
      _160_v3 = new BigNumber(781);
      let _161_v4;
      _161_v4 = _dafny.Map.Empty.slice().updateUnsafe(_160_v3,_157_v0);
      let _162_v5;
      _162_v5 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_160_v3,_157_v0), _161_v4, _161_v4, _161_v4);
      let _163_v6;
      _163_v6 = _dafny.Seq.UnicodeFromString("kajndjn");
      let _164_v7;
      _164_v7 = _dafny.Seq.of(_163_v6);
      let _165_v8;
      let _init1 = ((_166_v6) => function (_167_i0) {
        return (_167_i0).multipliedBy(new BigNumber((_166_v6).length));
      })(_163_v6);
      let _nw14 = Array((new BigNumber(4)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw14.length); _i0_1++) {
        _nw14[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _165_v8 = _nw14;
      let _168_v9;
      _168_v9 = _dafny.Set.fromElements((_dafny.ZERO).minus(_160_v3), _160_v3);
      let _169_v10;
      let _init2 = ((_170_v0) => function (_171_i1) {
        return _170_v0;
      })(_157_v0);
      let _nw15 = Array((new BigNumber(9)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw15.length); _i0_2++) {
        _nw15[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _169_v10 = _nw15;
      let _172_globalState;
      let _nw16 = new _module.GlobalState();
      _nw16.__ctor(new BigNumber(-19), false, true, new BigNumber(989), _159_v2, (_162_v5).Difference(_dafny.MultiSet.fromElements(_161_v4, _161_v4)), new BigNumber(523), false, _dafny.Seq.Concat(_164_v7, _164_v7), _165_v8, false, _168_v9, true, new BigNumber(565), _169_v10);
      _172_globalState = _nw16;
      let _173_v11;
      _173_v11 = new _dafny.CodePoint('w'.codePointAt(0));
      _173_v11 = _173_v11;
      let _174_v12;
      _174_v12 = _dafny.Seq.of(true);
      _160_v3 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_174_v12).length), (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_163_v6).length))).multipliedBy(_160_v3))));
      let _175_v13;
      let _nw17 = Array((new BigNumber(3)).toNumber());
      _nw17[(_dafny.ZERO).toNumber()] = _165_v8;
      _nw17[(_dafny.ONE).toNumber()] = _165_v8;
      _nw17[(new BigNumber(2)).toNumber()] = _165_v8;
      _175_v13 = _nw17;
      let _nw18 = Array((_dafny.ONE).toNumber());
      _nw18[(_dafny.ZERO).toNumber()] = _165_v8;
      _175_v13 = _nw18;
      let _176_v14;
      let _init3 = ((_177_v4) => function (_178_i2) {
        return _177_v4;
      })(_161_v4);
      let _nw19 = Array((new BigNumber(9)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw19.length); _i0_3++) {
        _nw19[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _176_v14 = _nw19;
      let _index26 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_176_v14).length));
      (_176_v14)[_index26] = _dafny.Map.Empty.slice().updateUnsafe(_160_v3,_157_v0);
      let _index27 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_176_v14).length));
      (_176_v14)[_index27] = ((_161_v4).Merge(_dafny.Map.Empty.slice().updateUnsafe(_160_v3,_157_v0))).Merge(_module.__default.fm0(_157_v0, false, _160_v3, _172_globalState));
      let _hi1 = _160_v3;
      for (let _179_i3 = new BigNumber((_174_v12).length); _179_i3.isLessThan(_hi1); _179_i3 = _179_i3.plus(_dafny.ONE)) {
        let _index28 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_169_v10).length));
        (_169_v10)[_index28] = _157_v0;
        let _index29 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_169_v10).length));
        (_169_v10)[_index29] = false;
        let _index30 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_169_v10).length));
        (_169_v10)[_index30] = !((_169_v10)[_module.__default.safeIndex(new BigNumber(390), new BigNumber((_169_v10).length))]);
        let _180_v15;
        _180_v15 = _dafny.MultiSet.fromElements(_165_v8, _165_v8, _165_v8);
        _180_v15 = (_180_v15).Difference((_180_v15).Intersect(_dafny.MultiSet.fromElements(_165_v8)));
        let _181_v16;
        _181_v16 = _dafny.Map.Empty.slice().updateUnsafe(_179_i3,_179_i3);
        let _182_v17;
        _182_v17 = _dafny.Map.Empty.slice().updateUnsafe(_174_v12,((_157_v0) ? (new BigNumber((_181_v16).length)) : (_179_i3)));
        let _183_v18;
        _183_v18 = _module.D0.create_DC0(_157_v0);
        _182_v17 = (_182_v17).update(_174_v12, _module.__default.fm1((_183_v18).dtor_cf0, _172_globalState));
      }
      (_172_globalState).f9 = _165_v8;
      let _184_v19;
      _184_v19 = _dafny.Map.Empty.slice().updateUnsafe(_165_v8,_169_v10);
      _184_v19 = _dafny.Map.Empty.slice().updateUnsafe(_165_v8,_169_v10);
      if ((_174_v12)[_module.__default.safeIndex((new BigNumber((_163_v6).length)).plus(_160_v3), new BigNumber((_174_v12).length))]) {
        let _185_v20;
        _185_v20 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_160_v3);
        let _186_v21;
        _186_v21 = _module.D0.create_DC0(_157_v0);
        let _187_v22;
        _187_v22 = _module.D0.create_DC2(_186_v21);
        let _188_v23;
        _188_v23 = _module.D0.create_DC2(_187_v22);
        let _pat_let_tv2 = _187_v22;
        let _189_v24;
        let _nw20 = Array((new BigNumber(13)).toNumber());
        _nw20[(_dafny.ZERO).toNumber()] = _188_v23;
        _nw20[(_dafny.ONE).toNumber()] = _188_v23;
        _nw20[(new BigNumber(2)).toNumber()] = _module.D0.create_DC2(_187_v22);
        _nw20[(new BigNumber(3)).toNumber()] = _188_v23;
        _nw20[(new BigNumber(4)).toNumber()] = function (_pat_let4_0) {
          return function (_190_dt__update__tmp_h0) {
            return function (_pat_let5_0) {
              return function (_191_dt__update_hcf2_h0) {
                return _module.D0.create_DC2(_191_dt__update_hcf2_h0);
              }(_pat_let5_0);
            }(_pat_let_tv2);
          }(_pat_let4_0);
        }(_188_v23);
        _nw20[(new BigNumber(5)).toNumber()] = _module.D0.create_DC2(_186_v21);
        _nw20[(new BigNumber(6)).toNumber()] = _188_v23;
        _nw20[(new BigNumber(7)).toNumber()] = _188_v23;
        _nw20[(new BigNumber(8)).toNumber()] = _188_v23;
        _nw20[(new BigNumber(9)).toNumber()] = _188_v23;
        _nw20[(new BigNumber(10)).toNumber()] = _188_v23;
        _nw20[(new BigNumber(11)).toNumber()] = _188_v23;
        _nw20[(new BigNumber(12)).toNumber()] = _188_v23;
        _189_v24 = _nw20;
        _module.__default.m0(_dafny.Set.fromElements(new BigNumber((_185_v20).length), _160_v3, _module.__default.fm1(_157_v0, _172_globalState)), _189_v24, _165_v8, _dafny.Seq.UnicodeFromString("hvsasg"), _172_globalState);
        let _192_v25;
        let _nw21 = Array((new BigNumber(16)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _192_v25 = _nw21;
        let _193_v26;
        _193_v26 = _dafny.MultiSet.fromElements(_192_v25);
        _193_v26 = _dafny.MultiSet.fromElements(((_157_v0) ? (_192_v25) : (_192_v25)));
        (_172_globalState).f1 = _157_v0;
        let _index31 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_169_v10).length));
        (_169_v10)[_index31] = _157_v0;
        let _index32 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_169_v10).length));
        (_169_v10)[_index32] = true;
        _module.__default.m0(_168_v9, _189_v24, _165_v8, _dafny.Seq.UnicodeFromString("wwejgvdb"), _172_globalState);
      } else {
        if (_157_v0) {
          (_172_globalState).f1 = _157_v0;
          let _194_v27;
          _194_v27 = _dafny.MultiSet.fromElements(_157_v0, _157_v0);
          let _195_v28;
          _195_v28 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_160_v3);
          let _196_v29;
          _196_v29 = _dafny.Seq.of(_160_v3);
          let _197_v30;
          _197_v30 = _dafny.Seq.of(new BigNumber((_195_v28).length), new BigNumber((_168_v9).length), new BigNumber((_196_v29).length));
          let _198_v31;
          _198_v31 = _dafny.Map.Empty.slice().updateUnsafe(_197_v30,_dafny.Seq.UnicodeFromString("xngsnlvv"));
          _163_v6 = _dafny.Seq.Concat(_dafny.Seq.update(_163_v6, _module.__default.safeIndex((_dafny.ZERO).minus((((_194_v27).contains(_157_v0)) ? ((_194_v27).get(_157_v0)) : (_160_v3))), new BigNumber((_163_v6).length)), _173_v11), (((_198_v31).contains(_196_v29)) ? ((_198_v31).get(_196_v29)) : (_163_v6)));
          let _rhs18 = ((new BigNumber(621)).multipliedBy(_160_v3)).multipliedBy(_160_v3);
          let _rhs19 = (_160_v3).plus(new BigNumber(398));
          let _rhs20 = _157_v0;
          let _rhs21 = _157_v0;
          let _lhs21 = _172_globalState;
          let _lhs22 = _172_globalState;
          let _lhs23 = _172_globalState;
          _lhs21.f6 = _rhs18;
          _lhs22.f6 = _rhs19;
          _lhs23.f1 = _rhs20;
          _157_v0 = _rhs21;
          (_172_globalState).f1 = _157_v0;
          let _199_v32;
          _199_v32 = _module.D0.create_DC1(_module.__default.fm1(_157_v0, _172_globalState));
          let _rhs22 = _dafny.Seq.UnicodeFromString("el");
          let _rhs23 = _module.D0.create_DC1(_160_v3);
          _163_v6 = _rhs22;
          _199_v32 = _rhs23;
        } else {
          _163_v6 = _163_v6;
          _160_v3 = _160_v3;
          let _200_v34;
          _200_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_158_v1).cardinality()),_173_v11);
          let _201_v35;
          _201_v35 = _dafny.Map.Empty.slice().updateUnsafe(false,function () {
            let _coll22 = new _dafny.Map();
            for (const _compr_22 of (_200_v34).Keys.Elements) {
              let _202_v33 = _compr_22;
              if ((_200_v34).contains(_202_v33)) {
                _coll22.push([(_202_v33).minus(_160_v3),_160_v3]);
              }
            }
            return _coll22;
          }());
          let _203_v36;
          _203_v36 = _dafny.Seq.of(_160_v3);
          let _204_v37;
          _204_v37 = _dafny.Seq.of(new BigNumber((_174_v12).length), new BigNumber((_203_v36).length));
          let _205_v39;
          _205_v39 = _dafny.Map.Empty.slice().updateUnsafe(_173_v11,_157_v0);
          let _206_v40;
          _206_v40 = _dafny.Map.Empty.slice().updateUnsafe((_204_v37)[_module.__default.safeIndex((_204_v37)[_module.__default.safeIndex(new BigNumber((function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of (_205_v39).Keys.Elements) {
              let _207_v38 = _compr_23;
              if ((_205_v39).contains(_207_v38)) {
                _coll23.push([_207_v38,_157_v0]);
              }
            }
            return _coll23;
          }()).length), new BigNumber((_204_v37).length))], new BigNumber((_204_v37).length))],_160_v3);
          _201_v35 = (_201_v35).update(_157_v0, (_206_v40).Merge(_206_v40));
          let _index33 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_165_v8).length));
          (_165_v8)[_index33] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_172_globalState),_dafny.Seq.UnicodeFromString("yxdfyxg"))).length);
          let _index34 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_165_v8).length));
          (_165_v8)[_index34] = (_160_v3).plus(_160_v3);
          let _208_v41;
          let _nw22 = new _module.C6();
          _nw22.__ctor();
          _208_v41 = _nw22;
          let _209_v42;
          _209_v42 = _dafny.MultiSet.fromElements(_208_v41);
          let _210_v43;
          let _nw23 = new _module.C9();
          _nw23.__ctor(((_157_v0) ? ((_165_v8)[_module.__default.safeIndex(new BigNumber(503), new BigNumber((_165_v8).length))]) : ((_dafny.ZERO).minus(new BigNumber((_209_v42).cardinality())))));
          _210_v43 = _nw23;
        }
        let _index35 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length));
        (_169_v10)[_index35] = false;
        let _index36 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length));
        (_169_v10)[_index36] = (_157_v0) && (_157_v0);
        let _211_v44;
        _211_v44 = _dafny.Seq.of(_160_v3);
        _211_v44 = (((_160_v3).isLessThan(new BigNumber(-916))) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), ((_212_v3) => function (_213_i4) {
          return (_dafny.ZERO).minus(_212_v3);
        })(_160_v3))) : (_211_v44));
        let _214_v45;
        _214_v45 = _dafny.Map.Empty.slice().updateUnsafe(_160_v3,_160_v3);
        let _215_v47;
        _215_v47 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,function () {
          let _coll24 = new _dafny.Set();
          for (const _compr_24 of (_214_v45).Keys.Elements) {
            let _216_v46 = _compr_24;
            if ((_214_v45).contains(_216_v46)) {
              _coll24.add(_module.__default.safeModuloInt(_216_v46, new BigNumber((_dafny.Seq.UnicodeFromString("co")).length)));
            }
          }
          return _coll24;
        }());
        let _217_v48;
        let _nw24 = Array((new BigNumber(13)).toNumber());
        _nw24[(_dafny.ZERO).toNumber()] = (_215_v47).Merge(_215_v47);
        _nw24[(_dafny.ONE).toNumber()] = _215_v47;
        _nw24[(new BigNumber(2)).toNumber()] = _215_v47;
        _nw24[(new BigNumber(3)).toNumber()] = (_215_v47).Merge(_215_v47);
        _nw24[(new BigNumber(4)).toNumber()] = ((_215_v47).update(_157_v0, _dafny.Set.fromElements(_160_v3, new BigNumber(576)))).Merge(_215_v47);
        _nw24[(new BigNumber(5)).toNumber()] = _215_v47;
        _nw24[(new BigNumber(6)).toNumber()] = _215_v47;
        _nw24[(new BigNumber(7)).toNumber()] = (_215_v47).Merge(_215_v47);
        _nw24[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm2(_172_globalState)),_168_v9);
        _nw24[(new BigNumber(9)).toNumber()] = _215_v47;
        _nw24[(new BigNumber(10)).toNumber()] = _215_v47;
        _nw24[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))],_168_v9);
        _nw24[(new BigNumber(12)).toNumber()] = (_215_v47).Merge(_215_v47);
        _217_v48 = _nw24;
        let _index37 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_217_v48).length));
        (_217_v48)[_index37] = (_215_v47).Merge(_215_v47);
        let _218_v49;
        let _nw25 = Array((new BigNumber(12)).toNumber()).fill(false);
        _218_v49 = _nw25;
        let _219_v50;
        _219_v50 = _dafny.Map.Empty.slice().updateUnsafe((_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))],_218_v49);
        let _220_v51;
        _220_v51 = _dafny.Set.fromElements((_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))]);
        let _index38 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_217_v48).length));
        let _rhs24 = _160_v3;
        let _rhs25 = (_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))];
        let _rhs26 = (((_219_v50).contains((_220_v51).contains((_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))]))) ? ((_219_v50).get((_220_v51).contains((_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))]))) : (_218_v49));
        let _rhs27 = (_dafny.Map.Empty.slice().updateUnsafe((_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))],_168_v9)).Merge(_215_v47);
        let _lhs24 = _172_globalState;
        let _lhs25 = _172_globalState;
        let _lhs26 = _172_globalState;
        let _lhs27 = _217_v48;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_217_v48).length));
        _lhs24.f6 = _rhs24;
        _lhs25.f1 = _rhs25;
        _lhs26.f14 = _rhs26;
        _lhs27[_lhs28] = _rhs27;
        if (_157_v0) {
          let _index39 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_165_v8).length));
          (_165_v8)[_index39] = _160_v3;
          let _index40 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length));
          let _index41 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_165_v8).length));
          let _rhs28 = (_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))];
          let _rhs29 = (_160_v3).plus((_dafny.ZERO).minus((_dafny.ZERO).minus(_160_v3)));
          let _lhs29 = _169_v10;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length));
          let _lhs31 = _165_v8;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_165_v8).length));
          _lhs29[_lhs30] = _rhs28;
          _lhs31[_lhs32] = _rhs29;
          let _index42 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_165_v8).length));
          (_165_v8)[_index42] = _160_v3;
          let _221_v52;
          _221_v52 = _dafny.Map.Empty.slice().updateUnsafe((_176_v14)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_176_v14).length))],_160_v3);
          let _222_v53;
          let _nw26 = new _module.C2();
          _nw26.__ctor((_165_v8)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_165_v8).length))], _160_v3, _221_v52);
          _222_v53 = _nw26;
          _163_v6 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("flwge"), _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_222_v53, _222_v53, _222_v53)).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("flwge")).length)), _173_v11);
          _220_v51 = (_220_v51).Intersect((_220_v51).Union(_220_v51));
          let _223_v54;
          let _nw27 = Array((new BigNumber(17)).toNumber()).fill(_module.D0.Default());
          _223_v54 = _nw27;
          let _224_v55;
          let _nw28 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _224_v55 = _nw28;
          _module.__default.m0(_168_v9, _223_v54, _224_v55, _163_v6, _172_globalState);
        } else {
          let _225_v56;
          _225_v56 = _dafny.Map.Empty.slice().updateUnsafe((_176_v14)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_176_v14).length))],_160_v3);
          let _226_v57;
          let _nw29 = new _module.C2();
          _nw29.__ctor(_160_v3, _160_v3, (_225_v56).Merge(_225_v56));
          _226_v57 = _nw29;
          (_172_globalState).f1 = (_226_v57.f21).isEqualTo(_160_v3);
          let _227_v58;
          let _nw30 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _227_v58 = _nw30;
          _227_v58 = _227_v58;
          let _228_v59;
          let _nw31 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Set.Empty);
          _228_v59 = _nw31;
          let _index43 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_228_v59).length));
          (_228_v59)[_index43] = _dafny.Set.fromElements(_163_v6);
          let _index44 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_228_v59).length));
          (_228_v59)[_index44] = _dafny.Set.fromElements((((_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))]) ? (_163_v6) : (_module.__default.fm21(_172_globalState))), _module.__default.fm21(_172_globalState));
          let _229_v60;
          let _nw32 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
          _229_v60 = _nw32;
          let _230_v61;
          _230_v61 = _module.D11.create_DC31(_174_v12, new BigNumber((_163_v6).length), (_226_v57).f20, (_226_v57).f20, _160_v3);
          let _231_v62;
          _231_v62 = _dafny.Set.fromElements(_230_v61);
          let _index45 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_229_v60).length));
          (_229_v60)[_index45] = (_231_v62).Difference(_dafny.Set.fromElements(_230_v61));
          let _232_v63;
          _232_v63 = _module.D3.create_DC7(new BigNumber((_214_v45).length));
          let _233_v64;
          _233_v64 = _module.D10.create_DC27(_226_v57.f21, _dafny.Seq.UnicodeFromString("upuot"), (_226_v57).f20, _232_v63, _dafny.Map.Empty.slice().updateUnsafe(_218_v49,(_169_v10)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_169_v10).length))]));
          let _index46 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_229_v60).length));
          (_229_v60)[_index46] = _module.__default.fm55((_dafny.ZERO).minus(_160_v3), (_233_v64).dtor_cf48, _160_v3, _172_globalState);
        }
      }
      _173_v11 = _173_v11;
      if (!(_157_v0) || (_157_v0)) {
        let _234_v65;
        _234_v65 = _dafny.Set.fromElements(_173_v11, _173_v11, new _dafny.CodePoint('p'.codePointAt(0)));
        _234_v65 = _234_v65;
        let _235_v66;
        _235_v66 = _dafny.Map.Empty.slice().updateUnsafe((_176_v14)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_176_v14).length))],_160_v3);
        let _236_v67;
        let _nw33 = new _module.C3();
        _nw33.__ctor(_157_v0, new BigNumber((_dafny.Seq.UnicodeFromString("upii")).length), _235_v66);
        _236_v67 = _nw33;
        let _237_v68;
        _237_v68 = _dafny.Set.fromElements(_157_v0);
        let _238_v69;
        _238_v69 = _module.D11.create_DC31(_174_v12, _160_v3, _160_v3, new BigNumber((_237_v68).length), _160_v3);
        let _239_v70;
        let _nw34 = Array((new BigNumber(19)).toNumber());
        _nw34[(_dafny.ZERO).toNumber()] = _174_v12;
        _nw34[(_dafny.ONE).toNumber()] = _174_v12;
        _nw34[(new BigNumber(2)).toNumber()] = _174_v12;
        _nw34[(new BigNumber(3)).toNumber()] = _174_v12;
        _nw34[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_157_v0, !(!(_157_v0)), _157_v0, _157_v0, _157_v0);
        _nw34[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_157_v0);
        _nw34[(new BigNumber(6)).toNumber()] = _174_v12;
        _nw34[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_174_v12, _174_v12);
        _nw34[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_174_v12, _174_v12);
        _nw34[(new BigNumber(9)).toNumber()] = _174_v12;
        _nw34[(new BigNumber(10)).toNumber()] = _174_v12;
        _nw34[(new BigNumber(11)).toNumber()] = _174_v12;
        _nw34[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_174_v12, _module.__default.safeIndex(_module.__default.fm1(_157_v0, _172_globalState), new BigNumber((_174_v12).length)), _157_v0);
        _nw34[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_174_v12, (_238_v69).dtor_cf54), _module.__default.safeIndex(_160_v3, new BigNumber((_dafny.Seq.Concat(_174_v12, (_238_v69).dtor_cf54)).length)), _157_v0);
        _nw34[(new BigNumber(14)).toNumber()] = _module.__default.fm19(_172_globalState);
        _nw34[(new BigNumber(15)).toNumber()] = _174_v12;
        _nw34[(new BigNumber(16)).toNumber()] = _174_v12;
        _nw34[(new BigNumber(17)).toNumber()] = _174_v12;
        _nw34[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_157_v0), _module.__default.safeIndex(new BigNumber(309), new BigNumber((_dafny.Seq.of(_157_v0)).length)), _157_v0);
        _239_v70 = _nw34;
        let _240_v71;
        _240_v71 = _dafny.Seq.of(_174_v12, _174_v12, _dafny.Seq.of(_157_v0));
        let _index47 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_239_v70).length));
        (_239_v70)[_index47] = (_240_v71)[_module.__default.safeIndex(_160_v3, new BigNumber((_240_v71).length))];
        let _241_v72;
        _241_v72 = _dafny.Seq.of(_169_v10, _169_v10);
        let _index48 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_239_v70).length));
        let _rhs30 = _157_v0;
        let _rhs31 = new BigNumber((_241_v72).length);
        let _rhs32 = _module.__default.fm2(_172_globalState);
        let _rhs33 = _236_v67;
        let _rhs34 = _174_v12;
        let _lhs33 = _172_globalState;
        let _lhs34 = _172_globalState;
        let _lhs35 = _172_globalState;
        let _lhs36 = _239_v70;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_239_v70).length));
        _lhs33.f1 = _rhs30;
        _lhs34.f6 = _rhs31;
        _lhs35.f1 = _rhs32;
        _236_v67 = _rhs33;
        _lhs36[_lhs37] = _rhs34;
        let _242_v73;
        _242_v73 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-359)));
        _242_v73 = _242_v73;
        (_172_globalState).f1 = _module.__default.fm2(_172_globalState);
        if (_157_v0) {
          let _243_v74;
          _243_v74 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_157_v0);
          let _244_v75;
          let _out0;
          _out0 = (_236_v67).m11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), ((_245_v11) => function (_246_i5) {
            return _245_v11;
          })(_173_v11)), new BigNumber((_243_v74).length), _172_globalState);
          _244_v75 = _out0;
          let _247_v76;
          let _nw35 = new _module.C8();
          _nw35.__ctor();
          _247_v76 = _nw35;
          _247_v76 = _247_v76;
          let _248_v77;
          _248_v77 = _dafny.Seq.of(_236_v67);
          let _249_v78;
          _249_v78 = _dafny.Map.Empty.slice().updateUnsafe(_248_v77,true);
          let _250_v79;
          let _nw36 = Array((new BigNumber(15)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = (_249_v78).Merge(_249_v78);
          _nw36[(_dafny.ONE).toNumber()] = _249_v78;
          _nw36[(new BigNumber(2)).toNumber()] = _249_v78;
          _nw36[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_248_v77,_157_v0);
          _nw36[(new BigNumber(4)).toNumber()] = _249_v78;
          _nw36[(new BigNumber(5)).toNumber()] = _249_v78;
          _nw36[(new BigNumber(6)).toNumber()] = _249_v78;
          _nw36[(new BigNumber(7)).toNumber()] = (_249_v78).Merge(_249_v78);
          _nw36[(new BigNumber(8)).toNumber()] = _249_v78;
          _nw36[(new BigNumber(9)).toNumber()] = _249_v78;
          _nw36[(new BigNumber(10)).toNumber()] = _249_v78;
          _nw36[(new BigNumber(11)).toNumber()] = (_249_v78).update(_248_v77, _157_v0);
          _nw36[(new BigNumber(12)).toNumber()] = _249_v78;
          _nw36[(new BigNumber(13)).toNumber()] = _249_v78;
          _nw36[(new BigNumber(14)).toNumber()] = _249_v78;
          _250_v79 = _nw36;
          let _index49 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_250_v79).length));
          (_250_v79)[_index49] = _249_v78;
          let _index50 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_250_v79).length));
          (_250_v79)[_index50] = (_249_v78).update(_248_v77, ((_157_v0) ? (_157_v0) : (_157_v0)));
          let _251_v80;
          _251_v80 = _dafny.Map.Empty.slice().updateUnsafe(_165_v8,_157_v0);
          _251_v80 = _251_v80;
          let _252_v81;
          let _nw37 = Array((new BigNumber(6)).toNumber()).fill([]);
          _252_v81 = _nw37;
          let _253_v82;
          _253_v82 = _module.D5.create_DC10(_252_v81);
          let _rhs35 = (_module.__default.safeDivisionInt(_160_v3, new BigNumber((_163_v6).length))).minus((new BigNumber(972)).multipliedBy(new BigNumber((_163_v6).length)));
          let _rhs36 = _160_v3;
          let _rhs37 = _253_v82;
          let _lhs38 = _172_globalState;
          let _lhs39 = _172_globalState;
          _lhs38.f6 = _rhs35;
          _lhs39.f6 = _rhs36;
          _253_v82 = _rhs37;
        } else {
          (_172_globalState).f6 = _160_v3;
          (_172_globalState).f1 = _module.__default.fm2(_172_globalState);
          (_172_globalState).f1 = _157_v0;
          let _index51 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_169_v10).length));
          (_169_v10)[_index51] = _157_v0;
          let _index52 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_169_v10).length));
          (_169_v10)[_index52] = ((_157_v0) ? (_module.__default.fm2(_172_globalState)) : (((_157_v0) ? (_157_v0) : (_157_v0))));
          let _254_v83;
          let _nw38 = new _module.C4();
          _nw38.__ctor();
          _254_v83 = _nw38;
          let _255_v84;
          let _nw39 = Array((new BigNumber(17)).toNumber()).fill(_module.D10.Default());
          _255_v84 = _nw39;
          let _256_v85;
          let _nw40 = Array((new BigNumber(24)).toNumber());
          _256_v85 = _nw40;
          let _257_v86;
          _257_v86 = _dafny.Seq.of(_256_v85, _256_v85);
          let _258_v87;
          _258_v87 = _dafny.Map.Empty.slice().updateUnsafe((_169_v10)[_module.__default.safeIndex(new BigNumber(252), new BigNumber((_169_v10).length))],_module.__default.fm1((_169_v10)[_module.__default.safeIndex(new BigNumber(252), new BigNumber((_169_v10).length))], _172_globalState));
          let _259_v88;
          let _nw41 = Array((new BigNumber(3)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = (_169_v10)[_module.__default.safeIndex(new BigNumber(252), new BigNumber((_169_v10).length))];
          _nw41[(_dafny.ONE).toNumber()] = _157_v0;
          _nw41[(new BigNumber(2)).toNumber()] = false;
          _259_v88 = _nw41;
          let _260_v89;
          _260_v89 = _module.D27.create_DC61(_257_v86);
          let _index53 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_169_v10).length));
          let _rhs38 = !_dafny.areEqual(_dafny.Seq.Concat(_241_v72, _dafny.Seq.update(_241_v72, _module.__default.safeIndex((((_258_v87).contains(_157_v0)) ? ((_258_v87).get(_157_v0)) : ((((_258_v87).contains(_157_v0)) ? ((_258_v87).get(_157_v0)) : (_module.__default.fm1(false, _172_globalState))))), new BigNumber((_241_v72).length)), _259_v88)), _dafny.Seq.Concat(_241_v72, _dafny.Seq.of(_259_v88, _259_v88)));
          let _rhs39 = _254_v83;
          let _rhs40 = _255_v84;
          let _rhs41 = _259_v88;
          let _rhs42 = (_260_v89).dtor_cf101;
          let _lhs40 = _169_v10;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_169_v10).length));
          let _lhs42 = _172_globalState;
          _lhs40[_lhs41] = _rhs38;
          _254_v83 = _rhs39;
          _255_v84 = _rhs40;
          _lhs42.f14 = _rhs41;
          _257_v86 = _rhs42;
        }
      } else {
        let _261_v90;
        let _nw42 = Array((new BigNumber(19)).toNumber()).fill(_module.D0.Default());
        _261_v90 = _nw42;
        _module.__default.m0(_168_v9, _261_v90, _165_v8, _dafny.Seq.Concat(_163_v6, _163_v6), _172_globalState);
        _165_v8 = ((true) ? (_165_v8) : (_165_v8));
        (_172_globalState).f6 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus(_160_v3), _160_v3));
        let _262_v91;
        _262_v91 = _dafny.MultiSet.fromElements(_173_v11);
        _262_v91 = (_262_v91).Intersect(_262_v91);
        let _263_v92;
        _263_v92 = _module.D3.create_DC6(_165_v8);
        let _264_v93;
        _264_v93 = _dafny.Seq.of(_263_v92);
        let _265_v94;
        let _nw43 = new _module.C0();
        _nw43.__ctor(_264_v93, _157_v0);
        _265_v94 = _nw43;
      }
      (_172_globalState).f1 = _module.__default.fm2(_172_globalState);
      let _266_v95;
      _266_v95 = _dafny.Map.Empty.slice().updateUnsafe(true,_160_v3);
      let _267_v96;
      _267_v96 = _dafny.Seq.of(_266_v95, _266_v95);
      let _268_v97;
      let _nw44 = Array((new BigNumber(10)).toNumber());
      _nw44[(_dafny.ZERO).toNumber()] = _266_v95;
      _nw44[(_dafny.ONE).toNumber()] = (_266_v95).Merge(_266_v95);
      _nw44[(new BigNumber(2)).toNumber()] = _266_v95;
      _nw44[(new BigNumber(3)).toNumber()] = (_266_v95).Merge(_266_v95);
      _nw44[(new BigNumber(4)).toNumber()] = _266_v95;
      _nw44[(new BigNumber(5)).toNumber()] = (_267_v96)[_module.__default.safeIndex(_160_v3, new BigNumber((_267_v96).length))];
      _nw44[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_160_v3);
      _nw44[(new BigNumber(7)).toNumber()] = _266_v95;
      _nw44[(new BigNumber(8)).toNumber()] = ((!(_157_v0)) ? (_dafny.Map.Empty.slice().updateUnsafe(true,_160_v3)) : (_266_v95));
      _nw44[(new BigNumber(9)).toNumber()] = _266_v95;
      _268_v97 = _nw44;
      let _index54 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_268_v97).length));
      (_268_v97)[_index54] = (_266_v95).Merge(_266_v95);
      let _269_v98;
      _269_v98 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(647),new BigNumber((_163_v6).length));
      let _270_v99;
      _270_v99 = _module.D12.create_DC34(_269_v98);
      let _index55 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_268_v97).length));
      (_268_v97)[_index55] = _module.__default.fm36(_173_v11, _270_v99, _172_globalState);
      let _271_v100;
      let _nw45 = Array((new BigNumber(29)).toNumber()).fill([]);
      _271_v100 = _nw45;
      let _source4 = _module.D5.create_DC10(_271_v100);
      if (_source4.is_DC11) {
        let _272___mcc_h0 = (_source4).cf14;
        let _273_cf14 = _272___mcc_h0;
        let _274_v101;
        _274_v101 = _dafny.Seq.of(_module.__default.fm1(_157_v0, _172_globalState));
        let _275_v102;
        _275_v102 = _module.D0.create_DC1(new BigNumber((_274_v101).length));
        let _276_v103;
        _276_v103 = _module.D0.create_DC2(_275_v102);
        let _pat_let_tv3 = _275_v102;
        let _277_v104;
        let _nw46 = Array((new BigNumber(17)).toNumber());
        _nw46[(_dafny.ZERO).toNumber()] = _module.D0.create_DC2(_275_v102);
        _nw46[(_dafny.ONE).toNumber()] = _276_v103;
        _nw46[(new BigNumber(2)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(3)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(4)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(5)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(6)).toNumber()] = function (_pat_let6_0) {
          return function (_278_dt__update__tmp_h1) {
            return function (_pat_let7_0) {
              return function (_279_dt__update_hcf2_h1) {
                return _module.D0.create_DC2(_279_dt__update_hcf2_h1);
              }(_pat_let7_0);
            }(_pat_let_tv3);
          }(_pat_let6_0);
        }(_276_v103);
        _nw46[(new BigNumber(7)).toNumber()] = _module.__default.fm56(_273_cf14, _173_v11, _172_globalState);
        _nw46[(new BigNumber(8)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(9)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(10)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(11)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(12)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(13)).toNumber()] = _module.D0.create_DC2(_275_v102);
        _nw46[(new BigNumber(14)).toNumber()] = _276_v103;
        _nw46[(new BigNumber(15)).toNumber()] = _module.D0.create_DC2(_275_v102);
        _nw46[(new BigNumber(16)).toNumber()] = _module.D0.create_DC2(_275_v102);
        _277_v104 = _nw46;
        _module.__default.m0(_168_v9, _277_v104, ((_157_v0) ? (_165_v8) : (_165_v8)), _dafny.Seq.Concat(_163_v6, _163_v6), _172_globalState);
        let _280_v105;
        let _nw47 = new _module.C8();
        _nw47.__ctor();
        _280_v105 = _nw47;
        let _281_v106;
        let _nw48 = new _module.C5();
        _nw48.__ctor(_273_cf14);
        _281_v106 = _nw48;
        (_172_globalState).f6 = _160_v3;
      } else if (_source4.is_DC12) {
        let _282___mcc_h1 = (_source4).cf15;
        let _283___mcc_h2 = (_source4).cf16;
        let _284___mcc_h3 = (_source4).cf17;
        let _285_cf17 = _284___mcc_h3;
        let _286_cf16 = _283___mcc_h2;
        let _287_cf15 = _282___mcc_h1;
        let _288_v107;
        _288_v107 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(932)).minus(_module.__default.fm1(_157_v0, _172_globalState)),_285_cf17);
        _288_v107 = (_288_v107).update(new BigNumber((function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of _dafny.IntegerRange(new BigNumber(-232), new BigNumber(590))) {
            let _289_v108 = _compr_25;
            if (((new BigNumber(-232)).isLessThanOrEqualTo(_289_v108)) && ((_289_v108).isLessThan(new BigNumber(590)))) {
              _coll25.add(_module.__default.safeDivisionInt(_289_v108, _160_v3));
            }
          }
          return _coll25;
        }()).length), _173_v11);
        (_172_globalState).f6 = _160_v3;
        if (true) {
          let _index56 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_165_v8).length));
          (_165_v8)[_index56] = _287_cf15;
          let _290_v109;
          _290_v109 = _dafny.Set.fromElements(_286_cf16, !(_286_cf16));
          let _291_v110;
          _291_v110 = _dafny.Seq.of(_160_v3, new BigNumber((_290_v109).length));
          let _292_v111;
          _292_v111 = _dafny.Seq.of(_291_v110, _291_v110);
          let _index57 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_165_v8).length));
          (_165_v8)[_index57] = new BigNumber(((_292_v111)[_module.__default.safeIndex(_160_v3, new BigNumber((_292_v111).length))]).length);
          _160_v3 = (_160_v3).multipliedBy(_287_cf15);
          let _index58 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_165_v8).length));
          (_165_v8)[_index58] = _287_cf15;
          let _293_v112;
          _293_v112 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jdakhap"), _163_v6), _module.__default.safeIndex(_160_v3, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jdakhap"), _163_v6)).length)), _173_v11),_158_v1);
          _293_v112 = (_293_v112).update(_163_v6, _dafny.MultiSet.FromArray(_291_v110));
          let _294_v113;
          _294_v113 = _module.D0.create_DC1((_165_v8)[_module.__default.safeIndex(new BigNumber(974), new BigNumber((_165_v8).length))]);
          let _295_v114;
          _295_v114 = _module.D0.create_DC2(_294_v113);
          let _pat_let_tv4 = _157_v0;
          let _pat_let_tv5 = _294_v113;
          let _296_v115;
          let _nw49 = Array((new BigNumber(14)).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = _295_v114;
          _nw49[(_dafny.ONE).toNumber()] = _295_v114;
          _nw49[(new BigNumber(2)).toNumber()] = _295_v114;
          _nw49[(new BigNumber(3)).toNumber()] = _295_v114;
          _nw49[(new BigNumber(4)).toNumber()] = _295_v114;
          _nw49[(new BigNumber(5)).toNumber()] = _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_294_v113)));
          _nw49[(new BigNumber(6)).toNumber()] = _295_v114;
          _nw49[(new BigNumber(7)).toNumber()] = function (_pat_let8_0) {
            return function (_297_dt__update__tmp_h2) {
              return function (_pat_let9_0) {
                return function (_298_dt__update_hcf2_h2) {
                  return _module.D0.create_DC2(_298_dt__update_hcf2_h2);
                }(_pat_let9_0);
              }(_module.D0.create_DC0(_pat_let_tv4));
            }(_pat_let8_0);
          }(_295_v114);
          _nw49[(new BigNumber(8)).toNumber()] = _module.D0.create_DC2(_294_v113);
          _nw49[(new BigNumber(9)).toNumber()] = _295_v114;
          _nw49[(new BigNumber(10)).toNumber()] = _295_v114;
          _nw49[(new BigNumber(11)).toNumber()] = _295_v114;
          _nw49[(new BigNumber(12)).toNumber()] = function (_pat_let10_0) {
            return function (_299_dt__update__tmp_h3) {
              return function (_pat_let11_0) {
                return function (_300_dt__update_hcf2_h3) {
                  return _module.D0.create_DC2(_300_dt__update_hcf2_h3);
                }(_pat_let11_0);
              }(_pat_let_tv5);
            }(_pat_let10_0);
          }(_295_v114);
          _nw49[(new BigNumber(13)).toNumber()] = _295_v114;
          _296_v115 = _nw49;
          let _301_v116;
          let _nw50 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _301_v116 = _nw50;
          _module.__default.m0((_168_v9).Union(_168_v9), _296_v115, _301_v116, _163_v6, _172_globalState);
        } else {
          let _302_v117;
          let _init4 = ((_303_v0) => function (_304_i6) {
            return _module.D0.create_DC2(_module.D0.create_DC0(_303_v0));
          })(_157_v0);
          let _nw51 = Array((new BigNumber(27)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw51.length); _i0_4++) {
            _nw51[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _302_v117 = _nw51;
          _module.__default.m0((_168_v9).Union(_168_v9), _302_v117, _165_v8, _dafny.Seq.UnicodeFromString("u"), _172_globalState);
          let _index59 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_165_v8).length));
          (_165_v8)[_index59] = _287_cf15;
          let _index60 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_165_v8).length));
          (_165_v8)[_index60] = _160_v3;
          _286_cf16 = !(_157_v0);
          let _305_v118;
          _305_v118 = _dafny.Seq.of(new BigNumber(-465), (_165_v8)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_165_v8).length))], new BigNumber(233));
          (_172_globalState).f1 = _dafny.Seq.contains(_305_v118, (_165_v8)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_165_v8).length))]);
          let _index61 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_169_v10).length));
          (_169_v10)[_index61] = _157_v0;
          let _306_v119;
          _306_v119 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_157_v0);
          let _307_v120;
          _307_v120 = _module.D20.create_DC50(_157_v0, new BigNumber((_dafny.Seq.UnicodeFromString("korfyx")).length), _287_cf15, _157_v0, _157_v0);
          let _308_v121;
          _308_v121 = _dafny.Set.fromElements((_307_v120).dtor_cf90);
          let _index62 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_169_v10).length));
          (_169_v10)[_index62] = ((_160_v3).isLessThan((_305_v118)[_module.__default.safeIndex(_160_v3, new BigNumber((_305_v118).length))])) || ((new BigNumber((_308_v121).length)).isLessThanOrEqualTo(new BigNumber((_306_v119).length)));
        }
        let _index63 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_165_v8).length));
        (_165_v8)[_index63] = _160_v3;
        let _index64 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_165_v8).length));
        (_165_v8)[_index64] = _287_cf15;
      } else if (_source4.is_DC10) {
        let _309___mcc_h4 = (_source4).cf13;
        let _310_cf13 = _309___mcc_h4;
        _160_v3 = _160_v3;
        (_172_globalState).f1 = false;
        let _311_v123;
        let _nw52 = Array((new BigNumber(20)).toNumber()).fill(_module.D0.Default());
        _311_v123 = _nw52;
        _module.__default.m0(function () {
          let _coll26 = new _dafny.Set();
          for (const _compr_26 of _dafny.IntegerRange(new BigNumber(29), new BigNumber(167))) {
            let _312_v122 = _compr_26;
            if (((new BigNumber(29)).isLessThanOrEqualTo(_312_v122)) && ((_312_v122).isLessThan(new BigNumber(167)))) {
              _coll26.add((_312_v122).multipliedBy(_160_v3));
            }
          }
          return _coll26;
        }(), _311_v123, _165_v8, _163_v6, _172_globalState);
        let _313_v124;
        _313_v124 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_173_v11);
        let _314_v125;
        _314_v125 = _313_v124;
        let _315_v126;
        _315_v126 = _dafny.MultiSet.fromElements(_314_v125, _313_v124);
        (_172_globalState).f6 = new BigNumber(((((true) ? (_315_v126) : (_315_v126))).Difference(_315_v126)).cardinality());
      } else {
        let _316___mcc_h5 = (_source4).cf18;
        let _317_cf18 = _316___mcc_h5;
        let _318_v127;
        let _nw53 = new _module.C4();
        _nw53.__ctor();
        _318_v127 = _nw53;
        _318_v127 = _318_v127;
        let _319_v128;
        let _320_v129;
        let _out1;
        let _out2;
        let _outcollector0 = (_318_v127).m1(_172_globalState);
        _out1 = _outcollector0[0];
        _out2 = _outcollector0[1];
        _319_v128 = _out1;
        _320_v129 = _out2;
        if (true) {
          _160_v3 = _160_v3;
          (_172_globalState).f6 = _319_v128;
          (_172_globalState).f6 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_160_v3, new BigNumber(-943)));
          let _321_v130;
          _321_v130 = _dafny.MultiSet.fromElements(_157_v0);
          let _322_v131;
          _322_v131 = _dafny.Map.Empty.slice().updateUnsafe((_176_v14)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_176_v14).length))],new BigNumber(177));
          let _323_v132;
          let _nw54 = new _module.C2();
          _nw54.__ctor(_160_v3, new BigNumber((_321_v130).cardinality()), _322_v131);
          _323_v132 = _nw54;
          (_172_globalState).f1 = _157_v0;
        } else {
          let _324_v133;
          let _init5 = ((_325_v11) => function (_326_i7) {
            return _325_v11;
          })(_173_v11);
          let _nw55 = Array((new BigNumber(20)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw55.length); _i0_5++) {
            _nw55[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _324_v133 = _nw55;
          _324_v133 = _324_v133;
          let _327_v134;
          let _nw56 = new _module.C8();
          _nw56.__ctor();
          _327_v134 = _nw56;
          let _index65 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_165_v8).length));
          (_165_v8)[_index65] = _module.__default.safeModuloInt(new BigNumber(-570), new BigNumber(-137));
          let _index66 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_165_v8).length));
          (_165_v8)[_index66] = _319_v128;
          _320_v129 = ((_165_v8)[_module.__default.safeIndex(new BigNumber(250), new BigNumber((_165_v8).length))]).isLessThan(_319_v128);
          let _index67 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_169_v10).length));
          (_169_v10)[_index67] = _module.__default.fm2(_172_globalState);
          let _index68 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_169_v10).length));
          (_169_v10)[_index68] = ((_157_v0) ? (!(!(true))) : (_320_v129));
        }
        let _328_v135;
        _328_v135 = _module.D5.create_DC12(_319_v128, _157_v0, _173_v11);
        _320_v129 = (_328_v135).dtor_cf16;
      }
      let _329_v136;
      _329_v136 = _dafny.Seq.of(_160_v3);
      let _330_v138;
      _330_v138 = _module.D0.create_DC0(_157_v0);
      let _331_v139;
      _331_v139 = _module.D0.create_DC2(_330_v138);
      let _332_v140;
      _332_v140 = _module.D0.create_DC2(_330_v138);
      let _333_v141;
      let _nw57 = Array((_dafny.ONE).toNumber());
      _nw57[(_dafny.ZERO).toNumber()] = _332_v140;
      _333_v141 = _nw57;
      _module.__default.m0(function () {
        let _coll27 = new _dafny.Set();
        for (const _compr_27 of (_329_v136).Elements) {
          let _334_v137 = _compr_27;
          if (_dafny.Seq.contains(_329_v136, _334_v137)) {
            _coll27.add(_module.__default.safeModuloInt(_334_v137, new BigNumber(634)));
          }
        }
        return _coll27;
      }(), _333_v141, _165_v8, _dafny.Seq.Concat(_163_v6, _163_v6), _172_globalState);
      if (_module.__default.fm2(_172_globalState)) {
        (_172_globalState).f6 = new BigNumber(286);
        let _335_v142;
        _335_v142 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_157_v0);
        let _336_v143;
        _336_v143 = _dafny.Map.Empty.slice().updateUnsafe(_161_v4,_160_v3);
        let _337_v144;
        let _nw58 = new _module.C3();
        _nw58.__ctor(_dafny.Seq.IsProperPrefixOf((_module.D14.create_DC39(_329_v136, _335_v142, false)).dtor_cf68, _329_v136), new BigNumber((_335_v142).length), (_336_v143).Merge(_336_v143));
        _337_v144 = _nw58;
        let _338_v145;
        _338_v145 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_174_v12));
        let _rhs43 = _337_v144;
        let _rhs44 = true;
        let _rhs45 = (_157_v0) === (_157_v0);
        let _rhs46 = _157_v0;
        let _rhs47 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_163_v6, _163_v6), _163_v6), _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_163_v6, _module.__default.safeIndex(new BigNumber((_168_v9).length), new BigNumber((_163_v6).length)), _173_v11)).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_163_v6, _163_v6), _163_v6)).length)), _173_v11), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_338_v145)[_module.__default.safeIndex(_160_v3, new BigNumber((_338_v145).length))]).cardinality())), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_163_v6, _163_v6), _163_v6), _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_163_v6, _module.__default.safeIndex(new BigNumber((_168_v9).length), new BigNumber((_163_v6).length)), _173_v11)).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_163_v6, _163_v6), _163_v6)).length)), _173_v11)).length)), _173_v11);
        let _lhs43 = _172_globalState;
        _337_v144 = _rhs43;
        _157_v0 = _rhs44;
        _157_v0 = _rhs45;
        _lhs43.f1 = _rhs46;
        _163_v6 = _rhs47;
        let _339_v146;
        _339_v146 = _dafny.MultiSet.fromElements(_157_v0, _157_v0);
        let _340_v147;
        let _nw59 = new _module.C2();
        _nw59.__ctor(_160_v3, new BigNumber((_339_v146).cardinality()), _dafny.Map.Empty.slice().updateUnsafe((_176_v14)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_176_v14).length))],_160_v3));
        _340_v147 = _nw59;
        (_172_globalState).f6 = (((_339_v146).contains(_157_v0)) ? ((_339_v146).get(_157_v0)) : (_340_v147.f21));
        if (true) {
          let _341_v148;
          _341_v148 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_163_v6);
          _341_v148 = (_341_v148).update(_157_v0, _163_v6);
          _338_v145 = _338_v145;
          (_172_globalState).f1 = (_174_v12)[_module.__default.safeIndex(_340_v147.f21, new BigNumber((_174_v12).length))];
          _165_v8 = _165_v8;
          let _342_v149;
          let _out3;
          _out3 = (_337_v144).m11(_163_v6, (_340_v147).f20, _172_globalState);
          _342_v149 = _out3;
        } else {
          (_172_globalState).f6 = (_160_v3).multipliedBy((_340_v147).f20);
          let _343_v150;
          _343_v150 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ep"),_337_v144);
          let _344_v151;
          _344_v151 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_340_v147.f21),_337_v144);
          _343_v150 = (_343_v150).update(_163_v6, (((_344_v151).contains((_340_v147).f20)) ? ((_344_v151).get((_340_v147).f20)) : (_337_v144)));
          let _345_v152;
          _345_v152 = _module.D0.create_DC0(_157_v0);
          (_172_globalState).f1 = (_340_v147).fm16(_345_v152, true, _172_globalState);
          let _346_v153;
          let _347_v154;
          let _out4;
          let _out5;
          let _outcollector1 = (_340_v147).m12((_340_v147).f20, _340_v147.f21, _dafny.Seq.Concat(_163_v6, _163_v6), (_164_v7)[_module.__default.safeIndex(_340_v147.f21, new BigNumber((_164_v7).length))], _172_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _346_v153 = _out4;
          _347_v154 = _out5;
          let _348_v155;
          _348_v155 = _dafny.Set.fromElements(_346_v153, _module.__default.fm2(_172_globalState));
          let _349_v156;
          _349_v156 = _module.D8.create_DC21(_dafny.Seq.UnicodeFromString("ahhjndpbh"), _348_v155);
          let _350_v157;
          let _out6;
          _out6 = (_340_v147).m11(_dafny.Seq.Concat(_dafny.Seq.update((_349_v156).dtor_cf34, _module.__default.safeIndex((_340_v147).f20, new BigNumber(((_349_v156).dtor_cf34).length)), _173_v11), _347_v154), _340_v147.f21, _172_globalState);
          _350_v157 = _out6;
        }
      } else {
        let _351_v158;
        _351_v158 = _dafny.Set.fromElements((_163_v6)[_module.__default.safeIndex(_160_v3, new BigNumber((_163_v6).length))]);
        (_172_globalState).f6 = (_dafny.ZERO).minus(((_160_v3).multipliedBy(_160_v3)).multipliedBy(((_157_v0) ? (_160_v3) : (new BigNumber((_351_v158).length)))));
        _160_v3 = _160_v3;
        let _352_v159;
        _352_v159 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_157_v0);
        let _353_v160;
        _353_v160 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_352_v159).length),_158_v1);
        if (((((_353_v160).contains(_160_v3)) ? ((_353_v160).get(_160_v3)) : (_module.__default.fm42(_module.__default.fm1(!(_157_v0), _172_globalState), _173_v11, _160_v3, true, _172_globalState)))).IsSubsetOf((_158_v1).Union(_158_v1))) {
          _157_v0 = _157_v0;
          let _354_v161;
          let _nw60 = new _module.C8();
          _nw60.__ctor();
          _354_v161 = _nw60;
          let _355_v162;
          _355_v162 = _dafny.Map.Empty.slice().updateUnsafe(_354_v161,new BigNumber(928));
          let _356_v163;
          let _nw61 = new _module.C9();
          _nw61.__ctor(new BigNumber((_163_v6).length));
          _356_v163 = _nw61;
          let _357_v164;
          _357_v164 = _dafny.Seq.of(_356_v163, _356_v163);
          _269_v98 = (_269_v98).update((((_355_v162).contains(_354_v161)) ? ((_355_v162).get(_354_v161)) : (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_357_v164).length),new BigNumber((_dafny.Seq.UnicodeFromString("glsxubw")).length))).update(_160_v3, (_356_v163).f15)).length))), (_dafny.ZERO).minus(new BigNumber((_163_v6).length)));
          let _358_v165;
          let _nw62 = new _module.C4();
          _nw62.__ctor();
          _358_v165 = _nw62;
          let _359_v166;
          let _nw63 = new _module.C5();
          _nw63.__ctor(_160_v3);
          _359_v166 = _nw63;
          let _360_v167;
          _360_v167 = _dafny.Map.Empty.slice().updateUnsafe(_160_v3,_359_v166);
          let _361_v168;
          let _nw64 = new _module.C2();
          _nw64.__ctor(new BigNumber((_360_v167).length), (_356_v163).f15, _dafny.Map.Empty.slice().updateUnsafe((_176_v14)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_176_v14).length))],(_359_v166).f16));
          _361_v168 = _nw64;
          let _362_v169;
          let _nw65 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
          _362_v169 = _nw65;
          let _363_v170;
          _363_v170 = _module.D20.create_DC49(_362_v169);
          let _364_v171;
          _364_v171 = _dafny.Map.Empty.slice().updateUnsafe(_361_v168,_363_v170);
          let _365_v172;
          _365_v172 = _dafny.Map.Empty.slice().updateUnsafe(_364_v171,_module.__default.fm2(_172_globalState));
          let _366_v173;
          let _nw66 = new _module.C5();
          _nw66.__ctor((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_160_v3, new BigNumber((_365_v172).length))));
          _366_v173 = _nw66;
          let _367_v174;
          let _368_v175;
          let _out7;
          let _out8;
          let _outcollector2 = (_361_v168).m12(new BigNumber((_dafny.Seq.Concat(_163_v6, _dafny.Seq.UnicodeFromString("swd"))).length), _160_v3, _dafny.Seq.Concat(_163_v6, _163_v6), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("yeigb"), _module.__default.safeIndex((_356_v163).f15, new BigNumber((_dafny.Seq.UnicodeFromString("yeigb")).length)), _173_v11), _163_v6), _172_globalState);
          _out7 = _outcollector2[0];
          _out8 = _outcollector2[1];
          _367_v174 = _out7;
          _368_v175 = _out8;
        } else {
          let _369_v176;
          _369_v176 = _module.D14.create_DC40(!(_157_v0));
          let _pat_let_tv6 = _157_v0;
          let _pat_let_tv7 = _157_v0;
          let _370_v177;
          let _nw67 = Array((new BigNumber(26)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _369_v176;
          _nw67[(_dafny.ONE).toNumber()] = _369_v176;
          _nw67[(new BigNumber(2)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(3)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(4)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(5)).toNumber()] = _module.__default.fm57(_164_v7, _174_v12, _172_globalState);
          _nw67[(new BigNumber(6)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(7)).toNumber()] = _module.D14.create_DC40(_157_v0);
          _nw67[(new BigNumber(8)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(9)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(10)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(11)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(12)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(13)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(14)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(15)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(16)).toNumber()] = _module.D14.create_DC40(_157_v0);
          _nw67[(new BigNumber(17)).toNumber()] = _module.D14.create_DC40(_157_v0);
          _nw67[(new BigNumber(18)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(19)).toNumber()] = _module.D14.create_DC40(_157_v0);
          _nw67[(new BigNumber(20)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(21)).toNumber()] = function (_pat_let12_0) {
            return function (_373_dt__update__tmp_h5) {
              return function (_pat_let15_0) {
                return function (_374_dt__update_hcf71_h1) {
                  return _module.D14.create_DC40(_374_dt__update_hcf71_h1);
                }(_pat_let15_0);
              }(_pat_let_tv7);
            }(_pat_let12_0);
          }(function (_pat_let13_0) {
            return function (_371_dt__update__tmp_h4) {
              return function (_pat_let14_0) {
                return function (_372_dt__update_hcf71_h0) {
                  return _module.D14.create_DC40(_372_dt__update_hcf71_h0);
                }(_pat_let14_0);
              }(_pat_let_tv6);
            }(_pat_let13_0);
          }(_369_v176));
          _nw67[(new BigNumber(22)).toNumber()] = _module.D14.create_DC40(true);
          _nw67[(new BigNumber(23)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(24)).toNumber()] = _369_v176;
          _nw67[(new BigNumber(25)).toNumber()] = _369_v176;
          _370_v177 = _nw67;
          let _index69 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_370_v177).length));
          (_370_v177)[_index69] = _369_v176;
          let _index70 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_370_v177).length));
          (_370_v177)[_index70] = _module.D14.create_DC40(_157_v0);
          let _375_v178;
          let _init6 = ((_376_v136) => function (_377_i8) {
            return _376_v136;
          })(_329_v136);
          let _nw68 = Array((new BigNumber(5)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw68.length); _i0_6++) {
            _nw68[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _375_v178 = _nw68;
          let _378_v179;
          _378_v179 = _329_v136;
          let _379_v180;
          _379_v180 = _dafny.Set.fromElements(_269_v98, _269_v98);
          let _index71 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_375_v178).length));
          (_375_v178)[_index71] = _dafny.Seq.Concat(_329_v136, _dafny.Seq.update((_378_v179), _module.__default.safeIndex(new BigNumber((_379_v180).length), new BigNumber(((_378_v179)).length)), _160_v3));
          let _380_v181;
          _380_v181 = _dafny.Map.Empty.slice().updateUnsafe(_161_v4,(_dafny.ZERO).minus(_160_v3));
          let _381_v182;
          let _nw69 = new _module.C3();
          _nw69.__ctor(_157_v0, _module.__default.fm1(_157_v0, _172_globalState), _380_v181);
          _381_v182 = _nw69;
          let _index72 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_375_v178).length));
          let _rhs48 = _160_v3;
          let _rhs49 = _173_v11;
          let _rhs50 = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_module.__default.fm28(_160_v3, true, _157_v0, _172_globalState)).length), _160_v3, _160_v3), _329_v136);
          let _rhs51 = _381_v182;
          let _lhs44 = _375_v178;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_375_v178).length));
          _160_v3 = _rhs48;
          _173_v11 = _rhs49;
          _lhs44[_lhs45] = _rhs50;
          _381_v182 = _rhs51;
          let _382_v183;
          _382_v183 = _dafny.Map.Empty.slice().updateUnsafe(_160_v3,_381_v182);
          let _383_v184;
          _383_v184 = _dafny.Map.Empty.slice().updateUnsafe(_157_v0,_382_v183);
          (_172_globalState).f6 = ((_375_v178)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_375_v178).length))])[_module.__default.safeIndex(new BigNumber(((_383_v184).Merge(_383_v184)).length), new BigNumber(((_375_v178)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_375_v178).length))]).length))];
          let _384_v185;
          _384_v185 = _380_v181;
          let _385_v186;
          let _nw70 = new _module.C3();
          _nw70.__ctor(_157_v0, _160_v3, (_384_v185));
          _385_v186 = _nw70;
          let _386_v187;
          _386_v187 = _dafny.Map.Empty.slice().updateUnsafe((_385_v186).f19,_385_v186);
          let _387_v188;
          let _nw71 = Array((new BigNumber(27)).toNumber());
          _nw71[(_dafny.ZERO).toNumber()] = _385_v186;
          _nw71[(_dafny.ONE).toNumber()] = _385_v186;
          _nw71[(new BigNumber(2)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(3)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(4)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(5)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(6)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(7)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(8)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(9)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(10)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(11)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(12)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(13)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(14)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(15)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(16)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(17)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(18)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(19)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(20)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(21)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(22)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(23)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(24)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(25)).toNumber()] = _385_v186;
          _nw71[(new BigNumber(26)).toNumber()] = (((_386_v187).contains((_385_v186).f19)) ? ((_386_v187).get((_385_v186).f19)) : (_385_v186));
          _387_v188 = _nw71;
          let _index73 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_387_v188).length));
          (_387_v188)[_index73] = _385_v186;
          let _index74 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_387_v188).length));
          (_387_v188)[_index74] = _385_v186;
          let _index75 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_175_v13).length));
          (_175_v13)[_index75] = _165_v8;
          let _388_v189;
          _388_v189 = _dafny.Map.Empty.slice().updateUnsafe(_269_v98,(_385_v186).f19);
          let _index76 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_175_v13).length));
          let _nw72 = Array((new BigNumber(20)).toNumber());
          _nw72[(_dafny.ZERO).toNumber()] = _160_v3;
          _nw72[(_dafny.ONE).toNumber()] = ((_385_v186).f19).plus((_385_v186).f19);
          _nw72[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_174_v12).length), new BigNumber((_dafny.Seq.of((_385_v186).f19)).length));
          _nw72[(new BigNumber(3)).toNumber()] = _160_v3;
          _nw72[(new BigNumber(4)).toNumber()] = ((!((_385_v186).f18)) ? (_160_v3) : (new BigNumber(916)));
          _nw72[(new BigNumber(5)).toNumber()] = _160_v3;
          _nw72[(new BigNumber(6)).toNumber()] = _160_v3;
          _nw72[(new BigNumber(7)).toNumber()] = _160_v3;
          _nw72[(new BigNumber(8)).toNumber()] = new BigNumber(106);
          _nw72[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt((_385_v186).f19, new BigNumber((_163_v6).length));
          _nw72[(new BigNumber(10)).toNumber()] = _160_v3;
          _nw72[(new BigNumber(11)).toNumber()] = new BigNumber(-541);
          _nw72[(new BigNumber(12)).toNumber()] = new BigNumber((_163_v6).length);
          _nw72[(new BigNumber(13)).toNumber()] = new BigNumber(-581);
          _nw72[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_385_v186).f19));
          _nw72[(new BigNumber(15)).toNumber()] = _160_v3;
          _nw72[(new BigNumber(16)).toNumber()] = _module.__default.fm1((_385_v186).f18, _172_globalState);
          _nw72[(new BigNumber(17)).toNumber()] = (new BigNumber(942)).multipliedBy(_160_v3);
          _nw72[(new BigNumber(18)).toNumber()] = new BigNumber((_388_v189).length);
          _nw72[(new BigNumber(19)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-559), _160_v3);
          (_175_v13)[_index76] = _nw72;
        }
        _161_v4 = (_161_v4).update(_160_v3, _157_v0);
        let _rhs52 = _165_v8;
        let _rhs53 = _module.__default.fm2(_172_globalState);
        let _rhs54 = (_329_v136)[_module.__default.safeIndex(_160_v3, new BigNumber((_329_v136).length))];
        let _lhs46 = _172_globalState;
        let _lhs47 = _172_globalState;
        let _lhs48 = _172_globalState;
        _lhs46.f9 = _rhs52;
        _lhs47.f1 = _rhs53;
        _lhs48.f6 = _rhs54;
      }
      (_172_globalState).f6 = (_160_v3).minus(new BigNumber(168));
      process.stdout.write(_dafny.toString(_157_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v1).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_v2).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_dafny.ONE),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_160_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_v5).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_163_v6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_164_v7, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("kajndjn")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v8)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v8)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v8)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v9).equals(_dafny.Set.fromElements(new BigNumber(-781), new BigNumber(781)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v10)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v10)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v10)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v10)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v10)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v10)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v10)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v10)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_172_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_172_globalState).f4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_dafny.ONE),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_172_globalState).f5).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_172_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_172_globalState).f8, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("kajndjn"), _dafny.Seq.UnicodeFromString("kajndjn")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_172_globalState).f11).equals(_dafny.Set.fromElements(new BigNumber(-781), new BigNumber(781)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f14)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f14)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f14)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f14)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f14)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f14)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f14)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f14)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_globalState.f14)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_173_v11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_174_v12, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_v13)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_v13)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_v13)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_v13)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_176_v14)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_176_v14)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_176_v14)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_176_v14)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true).updateUnsafe(_dafny.ZERO,true).updateUnsafe(new BigNumber(9),false).updateUnsafe(new BigNumber(898),true).updateUnsafe(new BigNumber(6),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_176_v14)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_176_v14)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_176_v14)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_176_v14)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_176_v14)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(781),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_184_v19).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_266_v95).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_267_v96, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO), _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE).updateUnsafe(false,new BigNumber(-685)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v97)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v98).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(647),new BigNumber(7)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v99).dtor_cf63).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(647),new BigNumber(7)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_329_v136, _dafny.Seq.of(_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_330_v138).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_331_v139).dtor_cf2).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_332_v140).dtor_cf2).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_333_v141)[_dafny.ZERO]).dtor_cf2).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D0(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO);
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
    static create_DC4(cf4) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D1(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC5(cf5) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf7) {
      let $dt = new D3(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC8(cf8, cf9, cf10, cf11) {
      let $dt = new D3(1);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC6(cf6) {
      let $dt = new D3(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf8) + ", " + this.cf9.toVerbatimString(true) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf6 === other.cf6;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7(_dafny.ZERO);
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
    static create_DC9(cf12) {
      let $dt = new D4(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf12) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC11(cf14) {
      let $dt = new D5(0);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC12(cf15, cf16, cf17) {
      let $dt = new D5(1);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC10(cf13) {
      let $dt = new D5(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC13(cf18) {
      let $dt = new D5(3);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get is_DC13() { return this.$tag === 3; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf13 === other.cf13;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC11(_dafny.ZERO);
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
    static create_DC14(cf19) {
      let $dt = new D6(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19);
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf21, cf22, cf23) {
      let $dt = new D7(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC17(cf24, cf25, cf26, cf27, cf28) {
      let $dt = new D7(1);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC15(cf20) {
      let $dt = new D7(2);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC18(cf29) {
      let $dt = new D7(3);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get is_DC18() { return this.$tag === 3; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21 && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24 && this.cf25 === other.cf25 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf20 === other.cf20;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC16(false, _dafny.ZERO, false);
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
    static create_DC20(cf31, cf32, cf33) {
      let $dt = new D8(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC21(cf34, cf35) {
      let $dt = new D8(1);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC19(cf30) {
      let $dt = new D8(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC21" + "(" + this.cf34.toVerbatimString(true) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC20(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC23(cf37, cf38, cf39, cf40, cf41) {
      let $dt = new D9(0);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC24(cf42, cf43) {
      let $dt = new D9(1);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC22(cf36) {
      let $dt = new D9(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC25(cf44) {
      let $dt = new D9(3);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get is_DC25() { return this.$tag === 3; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf37) + ", " + this.cf38.toVerbatimString(true) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf37, other.cf37) && _dafny.areEqual(this.cf38, other.cf38) && this.cf39 === other.cf39 && this.cf40 === other.cf40 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42) && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf36 === other.cf36;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC23(_dafny.Set.Empty, _dafny.Seq.UnicodeFromString(""), false, [], new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC27(cf46, cf47, cf48, cf49, cf50) {
      let $dt = new D10(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC28(cf51) {
      let $dt = new D10(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC26(cf45) {
      let $dt = new D10(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC29(cf52) {
      let $dt = new D10(3);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC29() { return this.$tag === 3; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf46) + ", " + this.cf47.toVerbatimString(true) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 3) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48) && _dafny.areEqual(this.cf49, other.cf49) && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf51, other.cf51);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf45 === other.cf45;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC27(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), _dafny.ZERO, _module.D3.Default(), _dafny.Map.Empty);
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
    static create_DC31(cf54, cf55, cf56, cf57, cf58) {
      let $dt = new D11(0);
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC32(cf59, cf60, cf61) {
      let $dt = new D11(1);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC30(cf53) {
      let $dt = new D11(2);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC33(cf62) {
      let $dt = new D11(3);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf54, other.cf54) && _dafny.areEqual(this.cf55, other.cf55) && _dafny.areEqual(this.cf56, other.cf56) && _dafny.areEqual(this.cf57, other.cf57) && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf59 === other.cf59 && _dafny.areEqual(this.cf60, other.cf60) && this.cf61 === other.cf61;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC31(_dafny.Seq.of(), _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC35() {
      let $dt = new D12(0);
      return $dt;
    }
    static create_DC34(cf63) {
      let $dt = new D12(1);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC35";
      } else if (this.$tag === 1) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf63) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf63, other.cf63);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC35();
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
    static create_DC37(cf65, cf66) {
      let $dt = new D13(0);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC36(cf64) {
      let $dt = new D13(1);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC37" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf64) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf65 === other.cf65 && this.cf66 === other.cf66;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC37(false, false);
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
    static create_DC39(cf68, cf69, cf70) {
      let $dt = new D14(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC40(cf71) {
      let $dt = new D14(1);
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC38(cf67) {
      let $dt = new D14(2);
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC41(cf72) {
      let $dt = new D14(3);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get is_DC41() { return this.$tag === 3; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC39" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC40" + "(" + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC38" + "(" + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 3) {
        return "D14.DC41" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68) && _dafny.areEqual(this.cf69, other.cf69) && this.cf70 === other.cf70;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf71 === other.cf71;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf67 === other.cf67;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf72, other.cf72);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC39(_dafny.Seq.of(), _dafny.Map.Empty, false);
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
    static create_DC43(cf74, cf75, cf76) {
      let $dt = new D15(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC42(cf73) {
      let $dt = new D15(1);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC43" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC42" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf74 === other.cf74 && this.cf75 === other.cf75 && this.cf76 === other.cf76;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC43(false, false, false);
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
    static create_DC45(cf78, cf79, cf80, cf81) {
      let $dt = new D16(0);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC44(cf77) {
      let $dt = new D16(1);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC44" + "(" + _dafny.toString(this.cf77) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79) && _dafny.areEqual(this.cf80, other.cf80) && this.cf81 === other.cf81;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC45(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, false);
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
    static create_DC46(cf82) {
      let $dt = new D17(0);
      $dt.cf82 = cf82;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get dtor_cf82() { return this.cf82; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC46" + "(" + _dafny.toString(this.cf82) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf82, other.cf82);
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
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC47(cf83) {
      let $dt = new D18(0);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC47" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf83, other.cf83);
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
    static create_DC48(cf84) {
      let $dt = new D19(0);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC48" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf84, other.cf84);
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
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC50(cf86, cf87, cf88, cf89, cf90) {
      let $dt = new D20(0);
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC49(cf85) {
      let $dt = new D20(1);
      $dt.cf85 = cf85;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf85() { return this.cf85; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC50" + "(" + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC49" + "(" + _dafny.toString(this.cf85) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf86 === other.cf86 && _dafny.areEqual(this.cf87, other.cf87) && _dafny.areEqual(this.cf88, other.cf88) && this.cf89 === other.cf89 && this.cf90 === other.cf90;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf85 === other.cf85;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC50(false, _dafny.ZERO, _dafny.ZERO, false, false);
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
    static create_DC51(cf91) {
      let $dt = new D21(0);
      $dt.cf91 = cf91;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get dtor_cf91() { return this.cf91; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC51" + "(" + _dafny.toString(this.cf91) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf91 === other.cf91;
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
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC52(cf92) {
      let $dt = new D22(0);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC52" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf92 === other.cf92;
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
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC54() {
      let $dt = new D23(0);
      return $dt;
    }
    static create_DC55(cf94) {
      let $dt = new D23(1);
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC53(cf93) {
      let $dt = new D23(2);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get is_DC55() { return this.$tag === 1; }
    get is_DC53() { return this.$tag === 2; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC54";
      } else if (this.$tag === 1) {
        return "D23.DC55" + "(" + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC53" + "(" + _dafny.toString(this.cf93) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf94, other.cf94);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf93 === other.cf93;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC54();
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
    static create_DC56(cf95) {
      let $dt = new D24(0);
      $dt.cf95 = cf95;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get dtor_cf95() { return this.cf95; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC56" + "(" + _dafny.toString(this.cf95) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf95, other.cf95);
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
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC58(cf97, cf98) {
      let $dt = new D25(0);
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC57(cf96) {
      let $dt = new D25(1);
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC59(cf99) {
      let $dt = new D25(2);
      $dt.cf99 = cf99;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC59() { return this.$tag === 2; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf99() { return this.cf99; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC58" + "(" + _dafny.toString(this.cf97) + ", " + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC57" + "(" + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 2) {
        return "D25.DC59" + "(" + _dafny.toString(this.cf99) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf97, other.cf97) && _dafny.areEqual(this.cf98, other.cf98);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf99, other.cf99);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC58(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC60(cf100) {
      let $dt = new D26(0);
      $dt.cf100 = cf100;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get dtor_cf100() { return this.cf100; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC60" + "(" + _dafny.toString(this.cf100) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf100 === other.cf100;
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
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC62(cf102, cf103, cf104, cf105) {
      let $dt = new D27(0);
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      return $dt;
    }
    static create_DC63(cf106, cf107) {
      let $dt = new D27(1);
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      return $dt;
    }
    static create_DC61(cf101) {
      let $dt = new D27(2);
      $dt.cf101 = cf101;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get is_DC61() { return this.$tag === 2; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf101() { return this.cf101; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC62" + "(" + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ", " + _dafny.toString(this.cf104) + ", " + _dafny.toString(this.cf105) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC63" + "(" + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ")";
      } else if (this.$tag === 2) {
        return "D27.DC61" + "(" + _dafny.toString(this.cf101) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf102 === other.cf102 && this.cf103 === other.cf103 && this.cf104 === other.cf104 && this.cf105 === other.cf105;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf106 === other.cf106 && _dafny.areEqual(this.cf107, other.cf107);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf101, other.cf101);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC62(false, false, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
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
      this.f1 = false;
      this.f6 = _dafny.ZERO;
      this.f9 = [];
      this.f14 = [];
      this._f0 = _dafny.ZERO;
      this._f2 = false;
      this._f3 = _dafny.ZERO;
      this._f4 = _dafny.Map.Empty;
      this._f5 = _dafny.MultiSet.Empty;
      this._f7 = false;
      this._f8 = _dafny.Seq.of();
      this._f10 = false;
      this._f11 = _dafny.Set.Empty;
      this._f12 = false;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
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
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f22 = _dafny.Seq.of();
      this._f23 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f22, f23) {
      let _this = this;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      return;
    }
    fm17(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f23;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f24 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f24) {
      let _this = this;
      (_this)._f24 = f24;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source5 = _module.D1.create_DC4(new _dafny.CodePoint('a'.codePointAt(0)));
      if (_source5.is_DC4) {
        let _389___mcc_h0 = (_source5).cf4;
        let _390_cf4 = _389___mcc_h0;
        return _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC0(false)));
      } else {
        let _391___mcc_h1 = (_source5).cf3;
        let _392_cf3 = _391___mcc_h1;
        return _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1((_this).f24)));
      }
    };
    fm24(p0, p1, globalState) {
      let _this = this;
      return !(((new BigNumber(383)).isLessThan((_this).f24)) === (!(((_this).f24).isEqualTo(new BigNumber((function () {
        let _coll28 = new _dafny.Set();
        for (const _compr_28 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.MultiSet.fromElements(true, true))).length))).Elements) {
          let _393_v0 = _compr_28;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.MultiSet.fromElements(true, true))).length)), _393_v0)) {
            _coll28.add(_module.__default.safeModuloInt(_393_v0, new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-199)))).length)));
          }
        }
        return _coll28;
      }()).length)))));
    };
    fm25(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(920)), function (_394_i0) {
        return (_this).f24;
      }), _dafny.Seq.of((_this).f24)), _dafny.Seq.of((_dafny.ZERO).minus((_this).f24), (_this).f24));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _hi2 = (_this).f24;
      for (let _395_i0 = (_dafny.ZERO).minus((_this).f24); _395_i0.isLessThan(_hi2); _395_i0 = _395_i0.plus(_dafny.ONE)) {
        let _396_v0;
        let _nw73 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _396_v0 = _nw73;
        let _index77 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_396_v0).length));
        (_396_v0)[_index77] = ((_this).f24).multipliedBy((_this).f24);
        let _397_v1;
        _397_v1 = true;
        let _398_v2;
        _398_v2 = _dafny.Set.fromElements(_397_v1);
        let _399_v3;
        _399_v3 = _dafny.Set.fromElements(_398_v2, _398_v2);
        let _400_v4;
        _400_v4 = _dafny.Map.Empty.slice().updateUnsafe(_397_v1,(_this).f24);
        let _401_v5;
        _401_v5 = _dafny.Seq.UnicodeFromString("kffytsi");
        let _402_v6;
        _402_v6 = _dafny.Map.Empty.slice().updateUnsafe(_400_v4,_401_v5);
        let _index78 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_396_v0).length));
        (_396_v0)[_index78] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_399_v3).length), _module.__default.safeModuloInt(new BigNumber((_402_v6).length), (_this).f24)));
        let _403_v7;
        _403_v7 = new _dafny.CodePoint('p'.codePointAt(0));
        _403_v7 = _403_v7;
        let _404_v8;
        let _init7 = ((_405_v1) => function (_406_i1) {
          return _module.D1.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(_405_v1,_405_v1));
        })(_397_v1);
        let _nw74 = Array((new BigNumber(25)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw74.length); _i0_7++) {
          _nw74[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _404_v8 = _nw74;
        let _407_v9;
        _407_v9 = _dafny.Map.Empty.slice().updateUnsafe(_397_v1,_397_v1);
        let _index79 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_404_v8).length));
        (_404_v8)[_index79] = _module.D1.create_DC3(_407_v9);
        let _408_v10;
        _408_v10 = _module.D1.create_DC3(_407_v9);
        let _index80 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_404_v8).length));
        let _rhs55 = _408_v10;
        let _rhs56 = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.update(_dafny.Seq.Concat(_401_v5, _401_v5), _module.__default.safeIndex(new BigNumber(-886), new BigNumber((_dafny.Seq.Concat(_401_v5, _401_v5)).length)), _403_v7), _401_v5));
        let _lhs49 = _404_v8;
        let _lhs50 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_404_v8).length));
        _lhs49[_lhs50] = _rhs55;
        r1 = _rhs56;
        r1 = _397_v1;
      }
      let _409_v11;
      _409_v11 = false;
      let _410_v12;
      _410_v12 = _dafny.Set.fromElements(_409_v11);
      r1 = (_410_v12).IsSubsetOf(_410_v12);
      let _411_v13;
      let _init8 = function (_412_i2) {
        return _module.__default.safeModuloInt(_412_i2, (_this).f24);
      };
      let _nw75 = Array((new BigNumber(3)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw75.length); _i0_8++) {
        _nw75[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _411_v13 = _nw75;
      let _413_v14;
      _413_v14 = _dafny.Seq.of(_411_v13, _411_v13, _411_v13, _411_v13);
      let _414_v15;
      _414_v15 = _dafny.Seq.of(_module.D3.create_DC6((_413_v14)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f24), new BigNumber((_413_v14).length))]));
      let _415_v16;
      let _nw76 = new _module.C0();
      _nw76.__ctor(_414_v15, _409_v11);
      _415_v16 = _nw76;
      let _416_v17;
      _416_v17 = _dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)));
      let _417_v18;
      let _init9 = function (_418_i3) {
        return false;
      };
      let _nw77 = Array((new BigNumber(8)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw77.length); _i0_9++) {
        _nw77[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _417_v18 = _nw77;
      let _419_v19;
      _419_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_416_v17).length),_417_v18);
      let _420_v20;
      _420_v20 = _dafny.Map.Empty.slice().updateUnsafe((((_419_v19).contains((_this).f24)) ? ((_419_v19).get((_this).f24)) : (_417_v18)),(_415_v16).f23);
      let _421_v21;
      _421_v21 = _dafny.Map.Empty.slice().updateUnsafe(true,_420_v20);
      _421_v21 = (_421_v21).update((_415_v16).f23, ((_409_v11) ? (_420_v20) : (_420_v20)));
      let _index81 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length));
      (_411_v13)[_index81] = (_this).f24;
      let _422_v22;
      _422_v22 = _dafny.MultiSet.fromElements(!((_415_v16).f23));
      let _index82 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length));
      (_411_v13)[_index82] = (new BigNumber(192)).plus(new BigNumber((_422_v22).cardinality()));
      let _423_i4;
      _423_i4 = _dafny.ZERO;
      L1: {
        while (_dafny.Seq.IsProperPrefixOf(_416_v17, _416_v17)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_423_i4)) {
              break L1;
            }
            _423_i4 = (_423_i4).plus(_dafny.ONE);
            let _424_v23;
            _424_v23 = _module.D0.create_DC0(!(_409_v11));
            let _425_v24;
            _425_v24 = _dafny.Seq.of(_409_v11, true);
            let _426_v25;
            _426_v25 = _dafny.Map.Empty.slice().updateUnsafe(_424_v23,_425_v24);
            _426_v25 = (_426_v25).update(_424_v23, _425_v24);
            let _427_v26;
            _427_v26 = _dafny.Map.Empty.slice().updateUnsafe((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))],true);
            let _428_v27;
            _428_v27 = _dafny.Map.Empty.slice().updateUnsafe(!((_415_v16).f23),(_this).f24);
            _427_v26 = (_427_v26).update((((_428_v27).contains((_415_v16).f23)) ? ((_428_v27).get((_415_v16).f23)) : ((_this).f24)), _409_v11);
            let _429_v28;
            _429_v28 = new _dafny.CodePoint('s'.codePointAt(0));
            let _index83 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_417_v18).length));
            (_417_v18)[_index83] = (_415_v16).f23;
            let _index84 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_417_v18).length));
            let _rhs57 = (((_415_v16).f23) ? (_429_v28) : (_429_v28));
            let _rhs58 = false;
            let _rhs59 = (_dafny.ZERO).minus((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))]);
            let _lhs51 = _417_v18;
            let _lhs52 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_417_v18).length));
            _429_v28 = _rhs57;
            _lhs51[_lhs52] = _rhs58;
            r0 = _rhs59;
            let _430_v29;
            _430_v29 = _dafny.Set.fromElements((_this).f24, (_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))]);
            let _431_v30;
            let _nw78 = Array((new BigNumber(10)).toNumber());
            _nw78[(_dafny.ZERO).toNumber()] = _409_v11;
            _nw78[(_dafny.ONE).toNumber()] = _409_v11;
            _nw78[(new BigNumber(2)).toNumber()] = (_415_v16).f23;
            _nw78[(new BigNumber(3)).toNumber()] = _409_v11;
            _nw78[(new BigNumber(4)).toNumber()] = !((_415_v16).f23);
            _nw78[(new BigNumber(5)).toNumber()] = (_417_v18)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_417_v18).length))];
            _nw78[(new BigNumber(6)).toNumber()] = _409_v11;
            _nw78[(new BigNumber(7)).toNumber()] = (_415_v16).f23;
            _nw78[(new BigNumber(8)).toNumber()] = (_417_v18)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_417_v18).length))];
            _nw78[(new BigNumber(9)).toNumber()] = (_417_v18)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_417_v18).length))];
            _431_v30 = _nw78;
            let _432_v31;
            _432_v31 = _dafny.Set.fromElements(_431_v30, _431_v30);
            let _433_v32;
            _433_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_432_v31).length),_430_v29);
            let _434_v33;
            _434_v33 = _module.D0.create_DC0(_409_v11);
            let _435_v34;
            _435_v34 = _module.D0.create_DC2(_434_v33);
            let _436_v35;
            let _nw79 = Array((new BigNumber(12)).toNumber());
            _nw79[(_dafny.ZERO).toNumber()] = _416_v17;
            _nw79[(_dafny.ONE).toNumber()] = _416_v17;
            _nw79[(new BigNumber(2)).toNumber()] = _416_v17;
            _nw79[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("ocslm");
            _nw79[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(43)), ((_437_v28) => function (_438_i5) {
              return _437_v28;
            })(_429_v28));
            _nw79[(new BigNumber(5)).toNumber()] = _416_v17;
            _nw79[(new BigNumber(6)).toNumber()] = _416_v17;
            _nw79[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("dmbkie");
            _nw79[(new BigNumber(8)).toNumber()] = _416_v17;
            _nw79[(new BigNumber(9)).toNumber()] = _416_v17;
            _nw79[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-616)), function (_439_i6) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            });
            _nw79[(new BigNumber(11)).toNumber()] = _416_v17;
            _436_v35 = _nw79;
            let _440_v36;
            _440_v36 = _dafny.Map.Empty.slice().updateUnsafe(_436_v35,true);
            let _441_v38;
            _441_v38 = _dafny.Map.Empty.slice().updateUnsafe(_409_v11,_433_v32);
            let _442_v41;
            _442_v41 = _dafny.Seq.of(new BigNumber((_425_v24).length));
            let _443_v42;
            _443_v42 = _dafny.MultiSet.fromElements((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))], (_442_v41)[_module.__default.safeIndex((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))], new BigNumber((_442_v41).length))], (_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))], new BigNumber((_442_v41).length));
            let _444_v43;
            _444_v43 = _dafny.Map.Empty.slice().updateUnsafe((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))],new BigNumber((_443_v42).cardinality()));
            let _445_v44;
            let _nw80 = Array((new BigNumber(18)).toNumber());
            _nw80[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f24),_430_v29);
            _nw80[(_dafny.ONE).toNumber()] = _433_v32;
            _nw80[(new BigNumber(2)).toNumber()] = (_module.__default.fm26(_429_v28, _435_v34, (((_440_v36).contains(_436_v35)) ? ((_440_v36).get(_436_v35)) : (_409_v11)), (_415_v16).f23, globalState));
            _nw80[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_430_v29)).Merge(_433_v32);
            _nw80[(new BigNumber(4)).toNumber()] = (function () {
              let _coll29 = new _dafny.Map();
              for (const _compr_29 of _dafny.IntegerRange(new BigNumber(229), new BigNumber(9))) {
                let _446_v37 = _compr_29;
                if (((new BigNumber(229)).isLessThanOrEqualTo(_446_v37)) && ((_446_v37).isLessThan(new BigNumber(9)))) {
                  _coll29.push([(_446_v37).multipliedBy((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))]),_dafny.Set.fromElements(new BigNumber(511), (_this).f24, (_this).f24)]);
                }
              }
              return _coll29;
            }()).Merge(_433_v32);
            _nw80[(new BigNumber(5)).toNumber()] = _433_v32;
            _nw80[(new BigNumber(6)).toNumber()] = (_433_v32).Merge(_433_v32);
            _nw80[(new BigNumber(7)).toNumber()] = (_433_v32).Merge(_433_v32);
            _nw80[(new BigNumber(8)).toNumber()] = (_433_v32).Merge(_433_v32);
            _nw80[(new BigNumber(9)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Set.fromElements((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))], (_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))]))).Merge(_module.__default.fm27((_this).f24, (_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))], (_this).f24, (_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))], globalState));
            _nw80[(new BigNumber(10)).toNumber()] = _433_v32;
            _nw80[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))],_430_v29);
            _nw80[(new BigNumber(12)).toNumber()] = _433_v32;
            _nw80[(new BigNumber(13)).toNumber()] = (((_415_v16).f23) ? (((((_441_v38).contains((_417_v18)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_417_v18).length))])) ? ((_441_v38).get((_417_v18)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_417_v18).length))])) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,function () {
              let _coll30 = new _dafny.Set();
              for (const _compr_30 of _dafny.IntegerRange(new BigNumber(84), new BigNumber(808))) {
                let _447_v39 = _compr_30;
                if (((new BigNumber(84)).isLessThanOrEqualTo(_447_v39)) && ((_447_v39).isLessThan(new BigNumber(808)))) {
                  _coll30.add((_447_v39).multipliedBy(new BigNumber((_422_v22).cardinality())));
                }
              }
              return _coll30;
            }())))).update((_this).f24, _430_v29)) : (_433_v32));
            _nw80[(new BigNumber(14)).toNumber()] = _433_v32;
            _nw80[(new BigNumber(15)).toNumber()] = _433_v32;
            _nw80[(new BigNumber(16)).toNumber()] = (function () {
              let _coll31 = new _dafny.Map();
              for (const _compr_31 of (_444_v43).Keys.Elements) {
                let _448_v40 = _compr_31;
                if ((_444_v43).contains(_448_v40)) {
                  _coll31.push([(_448_v40).plus(new BigNumber(816)),_430_v29]);
                }
              }
              return _coll31;
            }()).Merge(_433_v32);
            _nw80[(new BigNumber(17)).toNumber()] = (_433_v32).Merge((_433_v32).update(new BigNumber((_416_v17).length), _dafny.Set.fromElements((_this).f24)));
            _445_v44 = _nw80;
            let _index85 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_445_v44).length));
            (_445_v44)[_index85] = _dafny.Map.Empty.slice().updateUnsafe((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))],_430_v29);
            let _449_v45;
            let _init10 = ((_450_v28) => function (_451_i7) {
              return _450_v28;
            })(_429_v28);
            let _nw81 = Array((new BigNumber(29)).toNumber());
            for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw81.length); _i0_10++) {
              _nw81[_i0_10] = _init10(new BigNumber(_i0_10));
            }
            _449_v45 = _nw81;
            let _452_v46;
            _452_v46 = _dafny.Map.Empty.slice().updateUnsafe(_449_v45,_433_v32);
            let _453_v47;
            _453_v47 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_425_v24).length)),_430_v29));
            let _index86 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_445_v44).length));
            (_445_v44)[_index86] = (((_452_v46).contains(_449_v45)) ? ((_452_v46).get(_449_v45)) : ((((_415_v16).f23) ? (_433_v32) : ((_453_v47)[_module.__default.safeIndex((_411_v13)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_411_v13).length))], new BigNumber((_453_v47).length))]))));
          }
        }
      }
      r0 = ((_dafny.ZERO).minus((_this).f24)).multipliedBy((_this).f24);
      r1 = !(!_dafny.areEqual(_413_v14, _413_v14));
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let _454_v0;
      _454_v0 = _dafny.Seq.UnicodeFromString("n");
      let _455_v1;
      _455_v1 = _dafny.Seq.of(_454_v0, _dafny.Seq.UnicodeFromString("wkapm"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(925)), function (_456_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }));
      let _457_v2;
      _457_v2 = _dafny.Seq.of(_455_v1);
      _455_v1 = (_457_v2)[_module.__default.safeIndex(((_this).f24).plus((_this).f24), new BigNumber((_457_v2).length))];
      let _458_v3;
      _458_v3 = _dafny.Seq.of(p0, p0);
      let _hi3 = _module.__default.fm1(p0, globalState);
      for (let _459_i1 = (((_458_v3)[_module.__default.safeIndex((_this).f24, new BigNumber((_458_v3).length))]) ? ((_this).f24) : (new BigNumber(323))); _459_i1.isLessThan(_hi3); _459_i1 = _459_i1.plus(_dafny.ONE)) {
        (globalState).f1 = p0;
        let _460_v4;
        _460_v4 = _module.D5.create_DC11(_459_i1);
        let _461_v5;
        _461_v5 = _module.D5.create_DC13(_460_v4);
        let _462_v6;
        _462_v6 = _dafny.Map.Empty.slice().updateUnsafe(_459_i1,p0);
        let _463_v7;
        _463_v7 = _dafny.Map.Empty.slice().updateUnsafe(_462_v6,(_this).f24);
        let _rhs60 = (_this).f24;
        let _rhs61 = new BigNumber((_dafny.Seq.Concat(_454_v0, _454_v0)).length);
        let _rhs62 = _461_v5;
        let _rhs63 = (_dafny.ZERO).minus((((_463_v7).contains(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,!(false)))) ? ((_463_v7).get(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,!(false)))) : ((_this).f24)));
        let _rhs64 = _module.__default.fm1((p0) || (p0), globalState);
        let _lhs53 = globalState;
        let _lhs54 = globalState;
        let _lhs55 = globalState;
        let _lhs56 = globalState;
        _lhs53.f6 = _rhs60;
        _lhs54.f6 = _rhs61;
        _461_v5 = _rhs62;
        _lhs55.f6 = _rhs63;
        _lhs56.f6 = _rhs64;
        (globalState).f1 = !(true);
        let _464_v8;
        let _init11 = function (_465_i2) {
          return _module.__default.safeDivisionInt(_465_i2, new BigNumber(-187));
        };
        let _nw82 = Array((new BigNumber(19)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw82.length); _i0_11++) {
          _nw82[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _464_v8 = _nw82;
        let _466_v9;
        _466_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_459_i1,_464_v8)).length))).cardinality()));
        let _467_v10;
        _467_v10 = _dafny.Map.Empty.slice().updateUnsafe(_466_v9,(_this).f24);
        (globalState).f6 = ((_this).f24).minus((((_467_v10).contains(_466_v9)) ? ((_467_v10).get(_466_v9)) : ((_this).f24)));
      }
      let _468_v11;
      let _init12 = function (_469_i3) {
        return (_469_i3).plus((_this).f24);
      };
      let _nw83 = Array((new BigNumber(18)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw83.length); _i0_12++) {
        _nw83[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _468_v11 = _nw83;
      let _470_v12;
      _470_v12 = _module.D3.create_DC6(_468_v11);
      let _471_v13;
      _471_v13 = _dafny.Seq.of(_module.D3.create_DC6(_468_v11), _470_v12);
      let _472_v14;
      let _nw84 = new _module.C0();
      _nw84.__ctor(_471_v13, p0);
      _472_v14 = _nw84;
      let _473_v15;
      _473_v15 = _dafny.Map.Empty.slice().updateUnsafe(_472_v14,p0);
      (globalState).f6 = (_module.__default.safeModuloInt(new BigNumber(397), new BigNumber((_473_v15).length))).plus((_this).f24);
      let _474_i4;
      _474_i4 = _dafny.ZERO;
      L2: {
        while ((_472_v14).f23) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_474_i4)) {
              break L2;
            }
            _474_i4 = (_474_i4).plus(_dafny.ONE);
            let _475_v16;
            _475_v16 = new _dafny.CodePoint('y'.codePointAt(0));
            let _476_v17;
            _476_v17 = _dafny.Map.Empty.slice().updateUnsafe(_475_v16,(_this).f24);
            _476_v17 = (_476_v17).update(_475_v16, ((_this).f24).multipliedBy((_this).f24));
            (globalState).f1 = (_dafny.Set.fromElements((_472_v14).f23, (_472_v14).f23)).IsDisjointFrom(_dafny.Set.fromElements(true, p0));
            _475_v16 = _475_v16;
            let _477_v18;
            _477_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
            let _478_v19;
            _478_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_477_v18);
            let _479_v20;
            _479_v20 = _dafny.Seq.of((_this).f24, (_this).f24, new BigNumber((_478_v19).length), (_this).f24, (_this).f24);
            let _480_v21;
            _480_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements(p0, (_472_v14).f23, (_472_v14).f23, (_472_v14).f23, (_472_v14).f23));
            let _481_v22;
            _481_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f24);
            let _482_v23;
            _482_v23 = _dafny.Set.fromElements(true);
            (globalState).f6 = _module.__default.safeModuloInt((_479_v20)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_458_v3, _module.__default.safeIndex(new BigNumber(((((_480_v21).contains((_472_v14).fm17(new BigNumber(864), new BigNumber((_dafny.Seq.of(new BigNumber((_481_v22).length))).length), (_472_v14).f23, globalState))) ? ((_480_v21).get((_472_v14).fm17(new BigNumber(864), new BigNumber((_dafny.Seq.of(new BigNumber((_481_v22).length))).length), (_472_v14).f23, globalState))) : (_482_v23))).length), new BigNumber((_458_v3).length)), (_472_v14).f23)).length), new BigNumber((_479_v20).length))], _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f24), (_this).f24));
          }
        }
      }
      (globalState).f6 = (_dafny.ZERO).minus((_this).f24);
      _454_v0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(403)), function (_483_i5) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      });
      return;
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f17 = _dafny.Map.Empty;
      this.f21 = _dafny.ZERO;
      this._f20 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    __ctor(f20, f21, f17) {
      let _this = this;
      (_this)._f20 = f20;
      (_this).f21 = f21;
      (_this)._f17 = f17;
      return;
    }
    fm15(globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-329)), function (_484_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(542)), function (_485_i1) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }))).length)).minus(new BigNumber(((_dafny.Set.fromElements(_module.D3.create_DC7(new BigNumber((_dafny.Seq.of(false, false)).length)))).Difference(_dafny.Set.fromElements(_module.D3.create_DC7(_this.f21)))).length))));
    };
    fm16(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(559), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f21,!(true))).length), new BigNumber(111), _this.f21, new BigNumber((_dafny.Set.fromElements(_this.f21)).length)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f20)).length), new BigNumber((_dafny.Seq.UnicodeFromString("kkfiqofmd")).length)))).cardinality())).isLessThanOrEqualTo(_this.f21);
    };
    m11(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let _486_v0;
      _486_v0 = true;
      (globalState).f1 = _486_v0;
      let _487_v1;
      _487_v1 = _module.D0.create_DC0(_486_v0);
      let _488_v2;
      _488_v2 = _module.D3.create_DC7((_this).f20);
      let _hi4 = (_dafny.ZERO).minus((_488_v2).dtor_cf7);
      for (let _489_i0 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).fm16(_487_v1, _486_v0, globalState),(_488_v2).dtor_cf7)).length); _489_i0.isLessThan(_hi4); _489_i0 = _489_i0.plus(_dafny.ONE)) {
        let _490_v3;
        let _nw85 = Array((new BigNumber(19)).toNumber()).fill([]);
        _490_v3 = _nw85;
        let _491_v4;
        let _nw86 = Array((new BigNumber(16)).toNumber()).fill([]);
        _491_v4 = _nw86;
        let _492_v5;
        _492_v5 = _module.D5.create_DC10(_491_v4);
        let _index87 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_490_v3).length));
        (_490_v3)[_index87] = (_492_v5).dtor_cf13;
        let _index88 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_490_v3).length));
        (_490_v3)[_index88] = _491_v4;
        let _493_v6;
        _493_v6 = new _dafny.CodePoint('v'.codePointAt(0));
        let _494_v7;
        _494_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_493_v6,_489_i0)).length),new BigNumber(811));
        (globalState).f6 = _module.__default.safeModuloInt(new BigNumber(((_494_v7).Merge(_494_v7)).length), new BigNumber(279));
        let _495_v8;
        _495_v8 = _dafny.Seq.of(p1, _this.f21);
        let _496_v9;
        let _nw87 = Array((new BigNumber(27)).toNumber());
        _nw87[(_dafny.ZERO).toNumber()] = p1;
        _nw87[(_dafny.ONE).toNumber()] = new BigNumber((_494_v7).length);
        _nw87[(new BigNumber(2)).toNumber()] = new BigNumber(659);
        _nw87[(new BigNumber(3)).toNumber()] = (_this).f20;
        _nw87[(new BigNumber(4)).toNumber()] = _489_i0;
        _nw87[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(568)), ((_497_v6) => function (_498_i1) {
          return _497_v6;
        })(_493_v6))).length);
        _nw87[(new BigNumber(6)).toNumber()] = new BigNumber((_module.__default.fm0(_486_v0, _486_v0, _489_i0, globalState)).length);
        _nw87[(new BigNumber(7)).toNumber()] = p1;
        _nw87[(new BigNumber(8)).toNumber()] = (_this).f20;
        _nw87[(new BigNumber(9)).toNumber()] = (_this).f20;
        _nw87[(new BigNumber(10)).toNumber()] = (_this).f20;
        _nw87[(new BigNumber(11)).toNumber()] = new BigNumber(527);
        _nw87[(new BigNumber(12)).toNumber()] = p1;
        _nw87[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), ((_499_i0) => function (_500_i2) {
          return _499_i0;
        })(_489_i0))).length);
        _nw87[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.update(_495_v8, _module.__default.safeIndex((_dafny.ZERO).minus(_489_i0), new BigNumber((_495_v8).length)), (_this).f20)).length);
        _nw87[(new BigNumber(15)).toNumber()] = (_this).f20;
        _nw87[(new BigNumber(16)).toNumber()] = (_this).f20;
        _nw87[(new BigNumber(17)).toNumber()] = (_this).f20;
        _nw87[(new BigNumber(18)).toNumber()] = _this.f21;
        _nw87[(new BigNumber(19)).toNumber()] = _this.f21;
        _nw87[(new BigNumber(20)).toNumber()] = new BigNumber(876);
        _nw87[(new BigNumber(21)).toNumber()] = new BigNumber(693);
        _nw87[(new BigNumber(22)).toNumber()] = new BigNumber((_dafny.Seq.update(p0, _module.__default.safeIndex(_489_i0, new BigNumber((p0).length)), _493_v6)).length);
        _nw87[(new BigNumber(23)).toNumber()] = p1;
        _nw87[(new BigNumber(24)).toNumber()] = new BigNumber(279);
        _nw87[(new BigNumber(25)).toNumber()] = p1;
        _nw87[(new BigNumber(26)).toNumber()] = new BigNumber(-252);
        _496_v9 = _nw87;
        let _501_v10;
        _501_v10 = _module.D3.create_DC6(_496_v9);
        let _502_v11;
        _502_v11 = _dafny.Seq.of(_501_v10);
        let _503_v12;
        let _nw88 = new _module.C0();
        _nw88.__ctor(_dafny.Seq.Concat(_502_v11, _502_v11), _486_v0);
        _503_v12 = _nw88;
        let _504_v13;
        let _init13 = ((_505_i0) => function (_506_i3) {
          return (_dafny.MultiSet.fromElements(_505_i0, _this.f21)).Union(_dafny.MultiSet.fromElements((_this).f20));
        })(_489_i0);
        let _nw89 = Array((new BigNumber(25)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw89.length); _i0_13++) {
          _nw89[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _504_v13 = _nw89;
        let _507_v14;
        _507_v14 = _dafny.MultiSet.fromElements(p1);
        let _index89 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_504_v13).length));
        (_504_v13)[_index89] = _507_v14;
        let _index90 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_504_v13).length));
        (_504_v13)[_index90] = _507_v14;
      }
      let _508_v15;
      _508_v15 = _dafny.Seq.of((_this).f20);
      let _509_v16;
      let _init14 = ((_510_v0) => function (_511_i4) {
        return _510_v0;
      })(_486_v0);
      let _nw90 = Array((new BigNumber(15)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw90.length); _i0_14++) {
        _nw90[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _509_v16 = _nw90;
      let _512_v17;
      _512_v17 = _dafny.Seq.of(_486_v0);
      let _513_v18;
      _513_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_512_v17),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length)));
      let _514_v19;
      let _nw91 = Array((new BigNumber(14)).toNumber());
      _nw91[(_dafny.ZERO).toNumber()] = _this.f21;
      _nw91[(_dafny.ONE).toNumber()] = (_508_v15)[_module.__default.safeIndex(_this.f21, new BigNumber((_508_v15).length))];
      _nw91[(new BigNumber(2)).toNumber()] = p1;
      _nw91[(new BigNumber(3)).toNumber()] = p1;
      _nw91[(new BigNumber(4)).toNumber()] = _this.f21;
      _nw91[(new BigNumber(5)).toNumber()] = p1;
      _nw91[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_509_v16, _509_v16)).length);
      _nw91[(new BigNumber(7)).toNumber()] = p1;
      _nw91[(new BigNumber(8)).toNumber()] = p1;
      _nw91[(new BigNumber(9)).toNumber()] = _this.f21;
      _nw91[(new BigNumber(10)).toNumber()] = _this.f21;
      _nw91[(new BigNumber(11)).toNumber()] = new BigNumber((_513_v18).length);
      _nw91[(new BigNumber(12)).toNumber()] = p1;
      _nw91[(new BigNumber(13)).toNumber()] = (_this).f20;
      _514_v19 = _nw91;
      let _515_v20;
      _515_v20 = _module.D3.create_DC6(_514_v19);
      let _source6 = _515_v20;
      if (_source6.is_DC7) {
        let _516___mcc_h0 = (_source6).cf7;
        let _517_cf7 = _516___mcc_h0;
        let _index91 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_509_v16).length));
        (_509_v16)[_index91] = _486_v0;
        let _index92 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_509_v16).length));
        (_509_v16)[_index92] = _486_v0;
        let _pat_let_tv8 = _514_v19;
        let _518_v21;
        let _nw92 = new _module.C0();
        _nw92.__ctor(_dafny.Seq.of(function (_pat_let16_0) {
          return function (_519_dt__update__tmp_h0) {
            return function (_pat_let17_0) {
              return function (_520_dt__update_hcf6_h0) {
                return _module.D3.create_DC6(_520_dt__update_hcf6_h0);
              }(_pat_let17_0);
            }(_pat_let_tv8);
          }(_pat_let16_0);
        }(_515_v20), _515_v20, _515_v20, _515_v20), !(_486_v0));
        _518_v21 = _nw92;
        let _521_v22;
        let _init15 = ((_522_v21) => function (_523_i5) {
          return _module.D0.create_DC0((_522_v21).f23);
        })(_518_v21);
        let _nw93 = Array((new BigNumber(2)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw93.length); _i0_15++) {
          _nw93[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _521_v22 = _nw93;
        let _524_v23;
        _524_v23 = _dafny.Map.Empty.slice().updateUnsafe((_509_v16)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_509_v16).length))],_486_v0);
        let _525_v24;
        _525_v24 = _dafny.Set.fromElements((_518_v21).f23);
        let _526_v25;
        _526_v25 = _dafny.Set.fromElements(new BigNumber((_525_v24).length));
        let _527_v26;
        _527_v26 = new _dafny.CodePoint('y'.codePointAt(0));
        let _528_v27;
        _528_v27 = _module.D3.create_DC8(_508_v15, p0, (_509_v16)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_509_v16).length))], _527_v26);
        let _pat_let_tv9 = _509_v16;
        let _pat_let_tv10 = _509_v16;
        let _pat_let_tv11 = _486_v0;
        let _nw94 = Array((new BigNumber(24)).toNumber());
        _nw94[(_dafny.ZERO).toNumber()] = _487_v1;
        _nw94[(_dafny.ONE).toNumber()] = _487_v1;
        _nw94[(new BigNumber(2)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(3)).toNumber()] = function (_pat_let18_0) {
          return function (_529_dt__update__tmp_h1) {
            return function (_pat_let19_0) {
              return function (_530_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_530_dt__update_hcf0_h0);
              }(_pat_let19_0);
            }((_pat_let_tv10)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_pat_let_tv9).length))]);
          }(_pat_let18_0);
        }(_487_v1);
        _nw94[(new BigNumber(4)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(5)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(6)).toNumber()] = function (_pat_let20_0) {
          return function (_531_dt__update__tmp_h2) {
            return function (_pat_let21_0) {
              return function (_532_dt__update_hcf0_h1) {
                return _module.D0.create_DC0(_532_dt__update_hcf0_h1);
              }(_pat_let21_0);
            }(_pat_let_tv11);
          }(_pat_let20_0);
        }(_487_v1);
        _nw94[(new BigNumber(7)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(8)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(9)).toNumber()] = _module.__default.fm18((_this).f20, _508_v15, globalState);
        _nw94[(new BigNumber(10)).toNumber()] = _module.__default.fm18(_517_cf7, _508_v15, globalState);
        _nw94[(new BigNumber(11)).toNumber()] = (((_518_v21).fm17(new BigNumber((_dafny.Seq.of((_509_v16)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_509_v16).length))])).length), p1, !((_518_v21).f23), globalState)) ? (_487_v1) : (_487_v1));
        _nw94[(new BigNumber(12)).toNumber()] = _module.__default.fm18(new BigNumber((_524_v23).length), _508_v15, globalState);
        _nw94[(new BigNumber(13)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(14)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(15)).toNumber()] = _module.D0.create_DC0(_486_v0);
        _nw94[(new BigNumber(16)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(17)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(18)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(19)).toNumber()] = _module.__default.fm18((_dafny.ZERO).minus(_this.f21), _508_v15, globalState);
        _nw94[(new BigNumber(20)).toNumber()] = _487_v1;
        _nw94[(new BigNumber(21)).toNumber()] = _module.__default.fm18(new BigNumber((_526_v25).length), (_528_v27).dtor_cf8, globalState);
        _nw94[(new BigNumber(22)).toNumber()] = _module.__default.fm18(_this.f21, _508_v15, globalState);
        _nw94[(new BigNumber(23)).toNumber()] = _487_v1;
        _521_v22 = _nw94;
        let _533_v28;
        _533_v28 = _dafny.MultiSet.fromElements((_518_v21).f23, _486_v0);
        let _rhs65 = (_dafny.MultiSet.FromArray(_512_v17)).IsDisjointFrom(((_dafny.MultiSet.fromElements(_486_v0, (_509_v16)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_509_v16).length))], (_518_v21).f23, (_518_v21).f23)).update((_518_v21).f23, _module.__default.abs(new BigNumber((_dafny.Seq.UnicodeFromString("npqmwqi")).length)))).Union(_533_v28));
        let _rhs66 = new BigNumber(115);
        let _rhs67 = !(!(_486_v0)) || (((_486_v0) ? ((_509_v16)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_509_v16).length))]) : (_486_v0)));
        let _rhs68 = false;
        let _rhs69 = (_dafny.ZERO).minus((new BigNumber((function () {
          let _coll32 = new _dafny.Map();
          for (const _compr_32 of (_dafny.Map.Empty.slice().updateUnsafe(_this.f21,new BigNumber(132))).Keys.Elements) {
            let _534_v29 = _compr_32;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_this.f21,new BigNumber(132))).contains(_534_v29)) {
              _coll32.push([_module.__default.safeDivisionInt(_534_v29, p1),_dafny.Map.Empty.slice().updateUnsafe((_518_v21).f23,_this.f21)]);
            }
          }
          return _coll32;
        }()).length)).minus((_this).f20));
        let _lhs57 = globalState;
        let _lhs58 = globalState;
        let _lhs59 = globalState;
        let _lhs60 = globalState;
        let _lhs61 = globalState;
        _lhs57.f1 = _rhs65;
        _lhs58.f6 = _rhs66;
        _lhs59.f1 = _rhs67;
        _lhs60.f1 = _rhs68;
        _lhs61.f6 = _rhs69;
      } else if (_source6.is_DC8) {
        let _535___mcc_h1 = (_source6).cf8;
        let _536___mcc_h2 = (_source6).cf9;
        let _537___mcc_h3 = (_source6).cf10;
        let _538___mcc_h4 = (_source6).cf11;
        let _539_cf11 = _538___mcc_h4;
        let _540_cf10 = _537___mcc_h3;
        let _541_cf9 = _536___mcc_h2;
        let _542_cf8 = _535___mcc_h1;
        let _index93 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_514_v19).length));
        (_514_v19)[_index93] = new BigNumber((_512_v17).length);
        let _index94 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_514_v19).length));
        (_514_v19)[_index94] = p1;
        _486_v0 = _540_cf10;
        let _543_v30;
        _543_v30 = _dafny.Map.Empty.slice().updateUnsafe(p0,_509_v16);
        let _544_v31;
        _544_v31 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_module.__default.fm19(globalState), _module.__default.safeIndex((_514_v19)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_514_v19).length))], new BigNumber((_module.__default.fm19(globalState)).length)), _540_cf10)).length), p1);
        let _545_v32;
        _545_v32 = _dafny.Map.Empty.slice().updateUnsafe(_544_v31,_509_v16);
        let _rhs70 = (_543_v30).Merge(_543_v30);
        let _rhs71 = (((_545_v32).contains(function () {
          let _coll34 = new _dafny.Set();
          for (const _compr_34 of _dafny.IntegerRange(new BigNumber(593), new BigNumber(214))) {
            let _547_v33 = _compr_34;
            if (((new BigNumber(593)).isLessThanOrEqualTo(_547_v33)) && ((_547_v33).isLessThan(new BigNumber(214)))) {
              _coll34.add((_547_v33).multipliedBy(new BigNumber(699)));
            }
          }
          return _coll34;
        }())) ? ((_545_v32).get(function () {
          let _coll33 = new _dafny.Set();
          for (const _compr_33 of _dafny.IntegerRange(new BigNumber(593), new BigNumber(214))) {
            let _546_v33 = _compr_33;
            if (((new BigNumber(593)).isLessThanOrEqualTo(_546_v33)) && ((_546_v33).isLessThan(new BigNumber(214)))) {
              _coll33.add((_546_v33).multipliedBy(new BigNumber(699)));
            }
          }
          return _coll33;
        }())) : (_509_v16));
        let _rhs72 = ((_486_v0) ? (_540_cf10) : (_540_cf10));
        let _rhs73 = _540_cf10;
        let _lhs62 = globalState;
        let _lhs63 = globalState;
        _543_v30 = _rhs70;
        _509_v16 = _rhs71;
        _lhs62.f1 = _rhs72;
        _lhs63.f1 = _rhs73;
        (globalState).f1 = (_module.D0.create_DC0(_486_v0)).dtor_cf0;
      } else {
        let _548___mcc_h5 = (_source6).cf6;
        let _549_cf6 = _548___mcc_h5;
        let _550_v34;
        _550_v34 = _dafny.Set.fromElements(_486_v0, _486_v0);
        (globalState).f1 = (_this).fm16((((_this).fm16(_module.__default.fm18(new BigNumber((_550_v34).length), _dafny.Seq.Create(_module.__default.abs(new BigNumber(411)), ((_551_p1) => function (_552_i6) {
          return _551_p1;
        })(p1)), globalState), _486_v0, globalState)) ? (_487_v1) : (_487_v1)), _486_v0, globalState);
        let _553_v35;
        _553_v35 = _dafny.Map.Empty.slice().updateUnsafe(_508_v15,_this.f21);
        let _554_v36;
        _554_v36 = _508_v15;
        _553_v35 = (_553_v35).update(_554_v36, (_this).f20);
        _512_v17 = _512_v17;
        if (_486_v0) {
          let _555_v37;
          _555_v37 = _dafny.Seq.of(_515_v20, _515_v20);
          let _556_v38;
          let _nw95 = new _module.C0();
          _nw95.__ctor(_555_v37, _486_v0);
          _556_v38 = _nw95;
          (_this).f21 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_this.f21));
          let _index95 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_509_v16).length));
          (_509_v16)[_index95] = (_556_v38).f23;
          let _557_v39;
          _557_v39 = _dafny.Seq.of(p0);
          let _558_v40;
          _558_v40 = _dafny.Map.Empty.slice().updateUnsafe(_486_v0,p1);
          let _index96 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_509_v16).length));
          (_509_v16)[_index96] = _dafny.Seq.IsProperPrefixOf(p0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ervcnys"), (_557_v39)[_module.__default.safeIndex((((_558_v40).contains((_556_v38).f23)) ? ((_558_v40).get((_556_v38).f23)) : ((_dafny.ZERO).minus((_this).f20))), new BigNumber((_557_v39).length))]));
          let _559_v41;
          let _nw96 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _559_v41 = _nw96;
          let _560_v42;
          _560_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_486_v0);
          let _index97 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_559_v41).length));
          (_559_v41)[_index97] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_560_v42).length),(_556_v38).f23);
          let _561_v43;
          _561_v43 = _dafny.Set.fromElements((_this.f21).minus(p1));
          let _562_v44;
          _562_v44 = _dafny.Seq.UnicodeFromString("iburkss");
          let _563_v45;
          _563_v45 = _dafny.Map.Empty.slice().updateUnsafe(_this.f21,true);
          let _564_v46;
          _564_v46 = _dafny.Seq.of(_563_v45);
          let _565_v47;
          _565_v47 = _dafny.Seq.of(_512_v17, _512_v17, _512_v17);
          let _index98 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_559_v41).length));
          let _rhs74 = (_564_v46)[_module.__default.safeIndex(p1, new BigNumber((_564_v46).length))];
          let _rhs75 = (_561_v43).Union(_dafny.Set.fromElements(_this.f21, _this.f21));
          let _rhs76 = !(_dafny.MultiSet.fromElements((_565_v47)[_module.__default.safeIndex(new BigNumber((_512_v17).length), new BigNumber((_565_v47).length))])).equals(_module.__default.fm20(new BigNumber(313), globalState));
          let _rhs77 = _562_v44;
          let _rhs78 = (_556_v38).f23;
          let _lhs64 = _559_v41;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_559_v41).length));
          let _lhs66 = globalState;
          let _lhs67 = globalState;
          _lhs64[_lhs65] = _rhs74;
          _561_v43 = _rhs75;
          _lhs66.f1 = _rhs76;
          _562_v44 = _rhs77;
          _lhs67.f1 = _rhs78;
          (globalState).f6 = p1;
        } else {
          let _566_v48;
          _566_v48 = _dafny.MultiSet.fromElements(_486_v0, _486_v0, _486_v0, _486_v0, _486_v0);
          (globalState).f1 = !(_566_v48).equals((((_dafny.MultiSet.FromArray(_512_v17)).update(_486_v0, _module.__default.abs(new BigNumber(991)))).update(true, _module.__default.abs(p1))).Intersect(_566_v48));
          let _567_v49;
          let _nw97 = Array((new BigNumber(18)).toNumber()).fill([]);
          _567_v49 = _nw97;
          let _index99 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_567_v49).length));
          (_567_v49)[_index99] = _509_v16;
          let _index100 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_567_v49).length));
          (_567_v49)[_index100] = _509_v16;
          (globalState).f9 = _514_v19;
          _512_v17 = _512_v17;
          let _index101 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_509_v16).length));
          (_509_v16)[_index101] = _486_v0;
          let _568_v50;
          _568_v50 = _dafny.Map.Empty.slice().updateUnsafe(_514_v19,p0);
          let _index102 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_509_v16).length));
          (_509_v16)[_index102] = !(!(_568_v50).contains(_549_cf6));
        }
      }
      let _569_v51;
      let _init16 = function (_570_i7) {
        return _dafny.Seq.UnicodeFromString("cvrhikco");
      };
      let _nw98 = Array((new BigNumber(13)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw98.length); _i0_16++) {
        _nw98[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _569_v51 = _nw98;
      _569_v51 = _569_v51;
      (globalState).f6 = p1;
      let _571_i8;
      _571_i8 = _dafny.ZERO;
      L3: {
        while (_486_v0) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_571_i8)) {
              break L3;
            }
            _571_i8 = (_571_i8).plus(_dafny.ONE);
            let _572_v52;
            _572_v52 = new _dafny.CodePoint('m'.codePointAt(0));
            let _573_v53;
            _573_v53 = _module.D1.create_DC4(_572_v52);
            r0 = (_573_v53).dtor_cf4;
            (_this).f21 = p1;
            (globalState).f1 = false;
            let _index103 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_514_v19).length));
            (_514_v19)[_index103] = (_this).f20;
            let _index104 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_514_v19).length));
            (_514_v19)[_index104] = _this.f21;
          }
        }
      }
      let _574_v54;
      _574_v54 = new _dafny.CodePoint('r'.codePointAt(0));
      r0 = _574_v54;
      return r0;
    }
    m12(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      if (_dafny.Seq.IsPrefixOf(p2, _module.__default.fm21(globalState))) {
        let _575_v0;
        _575_v0 = true;
        let _576_v1;
        _576_v1 = _dafny.MultiSet.fromElements(_this.f21, new BigNumber(585), p0, (_this).f20);
        let _577_v2;
        _577_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_576_v1).cardinality()),_575_v0);
        let _578_v3;
        _578_v3 = _dafny.MultiSet.fromElements((((_577_v2).contains((_this).f20)) ? ((_577_v2).get((_this).f20)) : (_575_v0)), !(_575_v0));
        let _579_v4;
        _579_v4 = _dafny.Map.Empty.slice().updateUnsafe(((_575_v0) ? (p1) : (_this.f21)),new BigNumber((_578_v3).cardinality()));
        _579_v4 = (_579_v4).update((_this).f20, _module.__default.safeDivisionInt(p1, p1));
        let _580_v5;
        let _init17 = ((_581_p1, _582_p2) => function (_583_i0) {
          return (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(124)), function (_584_i1) {
            return _dafny.Seq.of((_this).f20, new BigNumber(829));
          }), _dafny.Seq.of(_dafny.Seq.of(_581_p1, new BigNumber(-382)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), function (_585_i2) {
            return _this.f21;
          })))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), ((_586_p2) => function (_587_i3) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(760)), ((_588_p2) => function (_589_i4) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), ((_590_p2) => function (_591_i5) {
                return new BigNumber((_590_p2).length);
              })(_588_p2));
            })(_586_p2));
          })(_582_p2))));
        })(p1, p2);
        let _nw99 = Array((new BigNumber(22)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw99.length); _i0_17++) {
          _nw99[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _580_v5 = _nw99;
        let _592_v6;
        _592_v6 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-122)), function (_593_i6) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        })).length));
        let _594_v7;
        _594_v7 = _dafny.Seq.of(_592_v6);
        let _595_v8;
        _595_v8 = _dafny.MultiSet.fromElements(_594_v7, _594_v7, _594_v7, _594_v7, _594_v7);
        let _index105 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_580_v5).length));
        (_580_v5)[_index105] = _595_v8;
        let _596_v9;
        _596_v9 = _dafny.Seq.of(_595_v8, (_595_v8).Intersect(_dafny.MultiSet.fromElements(_module.__default.fm22(_575_v0, _575_v0, globalState))), _595_v8, (_595_v8).Intersect(_595_v8));
        let _index106 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_580_v5).length));
        (_580_v5)[_index106] = (_596_v9)[_module.__default.safeIndex((((_578_v3).contains(!(_module.__default.fm2(globalState)))) ? ((_578_v3).get(!(_module.__default.fm2(globalState)))) : (new BigNumber(-794))), new BigNumber((_596_v9).length))];
        let _597_v10;
        _597_v10 = new _dafny.CodePoint('i'.codePointAt(0));
        let _598_v11;
        _598_v11 = _dafny.Seq.of(_575_v0, _575_v0);
        let _599_v12;
        _599_v12 = _module.D5.create_DC12(new BigNumber(15), (_598_v11)[_module.__default.safeIndex(_this.f21, new BigNumber((_598_v11).length))], _597_v10);
        let _600_v13;
        _600_v13 = _module.D5.create_DC11(p0);
        let _601_v14;
        _601_v14 = _dafny.Seq.of(p1, new BigNumber(714), new BigNumber((_579_v4).length));
        let _602_v15;
        let _nw100 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _602_v15 = _nw100;
        let _603_v16;
        _603_v16 = _dafny.Map.Empty.slice().updateUnsafe(true,_602_v15);
        let _604_v17;
        let _nw101 = Array((new BigNumber(29)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = new BigNumber((_576_v1).cardinality());
        _nw101[(_dafny.ONE).toNumber()] = p0;
        _nw101[(new BigNumber(2)).toNumber()] = _this.f21;
        _nw101[(new BigNumber(3)).toNumber()] = p1;
        _nw101[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_575_v0,_597_v10)).length);
        _nw101[(new BigNumber(5)).toNumber()] = p1;
        _nw101[(new BigNumber(6)).toNumber()] = _this.f21;
        _nw101[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(79), p0)).cardinality());
        _nw101[(new BigNumber(8)).toNumber()] = new BigNumber(712);
        _nw101[(new BigNumber(9)).toNumber()] = _this.f21;
        _nw101[(new BigNumber(10)).toNumber()] = p1;
        _nw101[(new BigNumber(11)).toNumber()] = p1;
        _nw101[(new BigNumber(12)).toNumber()] = p1;
        _nw101[(new BigNumber(13)).toNumber()] = (_this).f20;
        _nw101[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.of((_599_v12).dtor_cf15)).length);
        _nw101[(new BigNumber(15)).toNumber()] = p1;
        _nw101[(new BigNumber(16)).toNumber()] = (_this).f20;
        _nw101[(new BigNumber(17)).toNumber()] = (_this).f20;
        _nw101[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("jxavydsx")).length);
        _nw101[(new BigNumber(19)).toNumber()] = (_this).f20;
        _nw101[(new BigNumber(20)).toNumber()] = p1;
        _nw101[(new BigNumber(21)).toNumber()] = new BigNumber(556);
        _nw101[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_600_v13).dtor_cf14);
        _nw101[(new BigNumber(23)).toNumber()] = new BigNumber((_601_v14).length);
        _nw101[(new BigNumber(24)).toNumber()] = p0;
        _nw101[(new BigNumber(25)).toNumber()] = _this.f21;
        _nw101[(new BigNumber(26)).toNumber()] = (_this).f20;
        _nw101[(new BigNumber(27)).toNumber()] = (_dafny.ZERO).minus((((_579_v4).contains(_this.f21)) ? ((_579_v4).get(_this.f21)) : ((_this).f20)));
        _nw101[(new BigNumber(28)).toNumber()] = new BigNumber((_603_v16).length);
        _604_v17 = _nw101;
        let _605_v18;
        _605_v18 = _module.D3.create_DC6(_602_v15);
        let _606_v19;
        _606_v19 = _dafny.Seq.of(_605_v18);
        let _607_v20;
        let _nw102 = new _module.C0();
        _nw102.__ctor(_dafny.Seq.Concat(_606_v19, _606_v19), _575_v0);
        _607_v20 = _nw102;
        (globalState).f6 = (_dafny.ZERO).minus(((_this).f20).plus(p1));
        let _608_v21;
        _608_v21 = _dafny.Map.Empty.slice().updateUnsafe(_578_v3,new BigNumber(998));
        let _609_v22;
        _609_v22 = _dafny.Set.fromElements(_604_v17, (_605_v18).dtor_cf6, _602_v15, _602_v15);
        let _610_v23;
        _610_v23 = _dafny.Map.Empty.slice().updateUnsafe(_575_v0,new BigNumber((_609_v22).length));
        _608_v21 = (_608_v21).update(_578_v3, new BigNumber((_610_v23).length));
      } else {
        let _611_v24;
        _611_v24 = new _dafny.CodePoint('v'.codePointAt(0));
        let _612_v25;
        _612_v25 = _dafny.Set.fromElements(_611_v24);
        _612_v25 = (_612_v25).Intersect((_dafny.Set.fromElements(_611_v24)).Intersect(_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)), _611_v24)));
        let _613_v26;
        let _nw103 = Array((new BigNumber(28)).toNumber()).fill(false);
        _613_v26 = _nw103;
        let _index107 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_613_v26).length));
        (_613_v26)[_index107] = true;
        let _614_v27;
        _614_v27 = false;
        let _index108 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_613_v26).length));
        (_613_v26)[_index108] = _614_v27;
        let _615_v28;
        let _nw104 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _615_v28 = _nw104;
        let _616_v29;
        _616_v29 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_this.f21, (_this).f20),((_614_v27) ? (_615_v28) : (_615_v28)));
        let _617_v30;
        _617_v30 = _dafny.Map.Empty.slice().updateUnsafe((_module.D5.create_DC12(p0, (_613_v26)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_613_v26).length))], _611_v24)).dtor_cf16,_615_v28);
        _616_v29 = ((_616_v29).Merge(_616_v29)).update(p1, (((_617_v30).contains((_613_v26)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_613_v26).length))])) ? ((_617_v30).get((_613_v26)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_613_v26).length))])) : (_615_v28)));
        let _618_v31;
        _618_v31 = _dafny.Seq.of(_614_v27, _614_v27, _614_v27);
        let _619_v32;
        _619_v32 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_618_v31).length)));
        let _620_v33;
        _620_v33 = _619_v32;
        let _621_v34;
        _621_v34 = _dafny.Map.Empty.slice().updateUnsafe((_613_v26)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_613_v26).length))],_611_v24);
        let _622_v35;
        _622_v35 = _module.D3.create_DC8(_619_v32, p2, (_613_v26)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_613_v26).length))], _611_v24);
        _620_v33 = _module.__default.fm23(_621_v34, _dafny.Seq.Concat(p2, (_622_v35).dtor_cf9), (_613_v26)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_613_v26).length))], (_this).f20, globalState);
        r0 = _614_v27;
      }
      let _623_v36;
      _623_v36 = true;
      let _624_v37;
      _624_v37 = _dafny.Seq.of(_623_v36);
      let _hi5 = new BigNumber((_dafny.Seq.update(_624_v37, _module.__default.safeIndex(p0, new BigNumber((_624_v37).length)), _623_v36)).length);
      for (let _625_i7 = _module.__default.safeDivisionInt((_this).fm15(globalState), _this.f21); _625_i7.isLessThan(_hi5); _625_i7 = _625_i7.plus(_dafny.ONE)) {
        let _626_v38;
        let _nw105 = Array((new BigNumber(20)).toNumber()).fill(false);
        _626_v38 = _nw105;
        let _627_v39;
        _627_v39 = _dafny.Set.fromElements(_626_v38);
        let _628_v40;
        let _nw106 = new _module.C1();
        _nw106.__ctor(new BigNumber((_627_v39).length));
        _628_v40 = _nw106;
        let _629_v41;
        _629_v41 = _dafny.Map.Empty.slice().updateUnsafe(_628_v40,_624_v37);
        let _630_v42;
        _630_v42 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(_628_v40,_624_v37));
        let _631_v43;
        let _nw107 = Array((new BigNumber(23)).toNumber());
        _nw107[(_dafny.ZERO).toNumber()] = _629_v41;
        _nw107[(_dafny.ONE).toNumber()] = _629_v41;
        _nw107[(new BigNumber(2)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(3)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(4)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(5)).toNumber()] = (_629_v41).Merge(_629_v41);
        _nw107[(new BigNumber(6)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(7)).toNumber()] = (_629_v41).Merge(_629_v41);
        _nw107[(new BigNumber(8)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(9)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_628_v40,_624_v37)).Merge(_629_v41);
        _nw107[(new BigNumber(10)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(11)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(12)).toNumber()] = (_629_v41).update(_628_v40, _624_v37);
        _nw107[(new BigNumber(13)).toNumber()] = (_629_v41).update(_628_v40, _624_v37);
        _nw107[(new BigNumber(14)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(15)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_628_v40,_624_v37);
        _nw107[(new BigNumber(17)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(18)).toNumber()] = (_629_v41).update(_628_v40, _624_v37);
        _nw107[(new BigNumber(19)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(20)).toNumber()] = _629_v41;
        _nw107[(new BigNumber(21)).toNumber()] = (((_630_v42).contains(_623_v36)) ? ((_630_v42).get(_623_v36)) : (_629_v41));
        _nw107[(new BigNumber(22)).toNumber()] = (_629_v41).Merge((_629_v41).update(_628_v40, _624_v37));
        _631_v43 = _nw107;
        let _index109 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_631_v43).length));
        (_631_v43)[_index109] = _dafny.Map.Empty.slice().updateUnsafe(_628_v40,_624_v37);
        let _index110 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_631_v43).length));
        (_631_v43)[_index110] = _629_v41;
        let _632_v44;
        _632_v44 = _dafny.Map.Empty.slice().updateUnsafe((_625_i7).isLessThan((_dafny.ZERO).minus((_dafny.ZERO).minus(p1))),new BigNumber(-522));
        _632_v44 = (_632_v44).update((_623_v36) && (_623_v36), (_dafny.ZERO).minus(p1));
        let _633_v45;
        let _nw108 = Array((new BigNumber(21)).toNumber()).fill([]);
        _633_v45 = _nw108;
        let _634_v46;
        _634_v46 = _module.D5.create_DC10(_633_v45);
        let _635_v47;
        _635_v47 = _dafny.Seq.of(_634_v46);
        let _636_v48;
        _636_v48 = _module.D7.create_DC15(_628_v40);
        let _637_v49;
        _637_v49 = _dafny.MultiSet.fromElements((_636_v48).dtor_cf20);
        let _638_v50;
        let _nw109 = Array((new BigNumber(13)).toNumber());
        _nw109[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(573)), function (_639_i8) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).length);
        _nw109[(_dafny.ONE).toNumber()] = new BigNumber(680);
        _nw109[(new BigNumber(2)).toNumber()] = p1;
        _nw109[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(((_this).f20).multipliedBy(new BigNumber((_dafny.Seq.update(_635_v47, _module.__default.safeIndex(p0, new BigNumber((_635_v47).length)), _module.D5.create_DC10(_633_v45))).length)));
        _nw109[(new BigNumber(4)).toNumber()] = (_this).f20;
        _nw109[(new BigNumber(5)).toNumber()] = new BigNumber(210);
        _nw109[(new BigNumber(6)).toNumber()] = (((_637_v49).contains(_628_v40)) ? ((_637_v49).get(_628_v40)) : (_this.f21));
        _nw109[(new BigNumber(7)).toNumber()] = new BigNumber((_624_v37).length);
        _nw109[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(_623_v36)).length), new BigNumber(367));
        _nw109[(new BigNumber(9)).toNumber()] = (p0).multipliedBy(new BigNumber(704));
        _nw109[(new BigNumber(10)).toNumber()] = p1;
        _nw109[(new BigNumber(11)).toNumber()] = ((_this).f20).plus(p0);
        _nw109[(new BigNumber(12)).toNumber()] = p0;
        _638_v50 = _nw109;
        let _index111 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_638_v50).length));
        (_638_v50)[_index111] = new BigNumber(-622);
        let _index112 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_638_v50).length));
        let _rhs79 = (_this).f20;
        let _rhs80 = new BigNumber(210);
        let _lhs68 = _638_v50;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_638_v50).length));
        let _lhs70 = globalState;
        _lhs68[_lhs69] = _rhs79;
        _lhs70.f6 = _rhs80;
        let _640_v51;
        _640_v51 = _dafny.Map.Empty.slice().updateUnsafe(_623_v36,_623_v36);
        let _641_v52;
        _641_v52 = _dafny.Set.fromElements(_623_v36);
        if (((new BigNumber((_640_v51).length)).isLessThanOrEqualTo(new BigNumber((_641_v52).length))) === (false)) {
          let _642_v53;
          _642_v53 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("x")).length), ((_638_v50)[_module.__default.safeIndex(new BigNumber(702), new BigNumber((_638_v50).length))]).multipliedBy((_this).f20), _module.__default.safeDivisionInt((_638_v50)[_module.__default.safeIndex(new BigNumber(702), new BigNumber((_638_v50).length))], p0));
          _642_v53 = _642_v53;
          let _643_v54;
          let _nw110 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _643_v54 = _nw110;
          let _index113 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_643_v54).length));
          (_643_v54)[_index113] = _dafny.Map.Empty.slice().updateUnsafe(_638_v50,_625_i7);
          let _644_v55;
          _644_v55 = _dafny.Map.Empty.slice().updateUnsafe(_638_v50,p1);
          let _index114 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_643_v54).length));
          (_643_v54)[_index114] = (_dafny.Map.Empty.slice().updateUnsafe(_638_v50,p0)).Merge(((_623_v36) ? (_644_v55) : (_dafny.Map.Empty.slice().updateUnsafe(_638_v50,new BigNumber((_624_v37).length)))));
          (globalState).f6 = new BigNumber(-256);
          let _645_v56;
          _645_v56 = _dafny.MultiSet.fromElements(_623_v36);
          let _646_v57;
          _646_v57 = _module.D5.create_DC11(p0);
          let _647_v58;
          let _nw111 = Array((new BigNumber(4)).toNumber());
          _nw111[(_dafny.ZERO).toNumber()] = _module.D5.create_DC11(p1);
          _nw111[(_dafny.ONE).toNumber()] = _646_v57;
          _nw111[(new BigNumber(2)).toNumber()] = _646_v57;
          _nw111[(new BigNumber(3)).toNumber()] = _646_v57;
          _647_v58 = _nw111;
          let _648_v59;
          _648_v59 = _dafny.Seq.of(_647_v58, _647_v58, _647_v58);
          let _649_v60;
          let _nw112 = Array((new BigNumber(27)).toNumber());
          _nw112[(_dafny.ZERO).toNumber()] = _647_v58;
          _nw112[(_dafny.ONE).toNumber()] = _647_v58;
          _nw112[(new BigNumber(2)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(3)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(4)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(5)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(6)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(7)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(8)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(9)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(10)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(11)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(12)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(13)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(14)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(15)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(16)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(17)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(18)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(19)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(20)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(21)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(22)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(23)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(24)).toNumber()] = _647_v58;
          _nw112[(new BigNumber(25)).toNumber()] = (_648_v59)[_module.__default.safeIndex(_625_i7, new BigNumber((_648_v59).length))];
          _nw112[(new BigNumber(26)).toNumber()] = _647_v58;
          _649_v60 = _nw112;
          let _650_v61;
          _650_v61 = _dafny.Map.Empty.slice().updateUnsafe((_645_v56).IsDisjointFrom(_645_v56),_649_v60);
          _650_v61 = (_650_v61).update((_624_v37)[_module.__default.safeIndex(p0, new BigNumber((_624_v37).length))], _649_v60);
          let _651_v62;
          _651_v62 = new _dafny.CodePoint('b'.codePointAt(0));
          _651_v62 = _651_v62;
        } else {
          let _652_v63;
          _652_v63 = _dafny.Seq.of(p1, _this.f21, new BigNumber((_dafny.Seq.Concat(p2, p2)).length));
          let _rhs81 = (_652_v63)[_module.__default.safeIndex(new BigNumber((p3).length), new BigNumber((_652_v63).length))];
          let _rhs82 = _623_v36;
          let _lhs71 = globalState;
          _lhs71.f6 = _rhs81;
          _623_v36 = _rhs82;
          let _653_v64;
          _653_v64 = _dafny.Set.fromElements((_638_v50)[_module.__default.safeIndex(new BigNumber(702), new BigNumber((_638_v50).length))], (_this).f20, _625_i7);
          let _654_v65;
          _654_v65 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC6(_638_v50),(_653_v64).Union(_dafny.Set.fromElements(_this.f21)));
          let _655_v66;
          _655_v66 = _module.D3.create_DC6(_638_v50);
          let _pat_let_tv12 = _638_v50;
          let _pat_let_tv13 = _638_v50;
          _653_v64 = (((_654_v65).contains(function (_pat_let24_0) {
            return function (_658_dt__update__tmp_h0) {
              return function (_pat_let25_0) {
                return function (_659_dt__update_hcf6_h0) {
                  return _module.D3.create_DC6(_659_dt__update_hcf6_h0);
                }(_pat_let25_0);
              }(_pat_let_tv13);
            }(_pat_let24_0);
          }(_655_v66))) ? ((_654_v65).get(function (_pat_let22_0) {
            return function (_656_dt__update__tmp_h1) {
              return function (_pat_let23_0) {
                return function (_657_dt__update_hcf6_h1) {
                  return _module.D3.create_DC6(_657_dt__update_hcf6_h1);
                }(_pat_let23_0);
              }(_pat_let_tv12);
            }(_pat_let22_0);
          }(_655_v66))) : (((_623_v36) ? (_module.__default.fm28((_638_v50)[_module.__default.safeIndex(new BigNumber(702), new BigNumber((_638_v50).length))], _623_v36, _623_v36, globalState)) : (_653_v64))));
          let _660_v67;
          _660_v67 = _dafny.MultiSet.fromElements(_625_i7, _module.__default.fm1(_623_v36, globalState));
          _660_v67 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_this.f21));
          (globalState).f1 = !(_623_v36);
          let _661_v68;
          _661_v68 = new _dafny.CodePoint('j'.codePointAt(0));
          let _662_v69;
          _662_v69 = _dafny.MultiSet.fromElements(_623_v36);
          _661_v68 = _module.__default.fm29(_623_v36, ((((_660_v67).contains((_638_v50)[_module.__default.safeIndex(new BigNumber(702), new BigNumber((_638_v50).length))])) ? ((_660_v67).get((_638_v50)[_module.__default.safeIndex(new BigNumber(702), new BigNumber((_638_v50).length))])) : (new BigNumber((_662_v69).cardinality())))).multipliedBy(p0), new BigNumber((_652_v63).length), (_641_v52).IsProperSubsetOf(_641_v52), globalState);
        }
      }
      let _rhs83 = (_module.__default.fm1(!(_623_v36), globalState)).minus((_this).f20);
      let _rhs84 = new BigNumber((p2).length);
      let _rhs85 = _623_v36;
      let _lhs72 = globalState;
      let _lhs73 = globalState;
      let _lhs74 = globalState;
      _lhs72.f6 = _rhs83;
      _lhs73.f6 = _rhs84;
      _lhs74.f1 = _rhs85;
      let _663_v70;
      _663_v70 = _dafny.Map.Empty.slice().updateUnsafe(_623_v36,_623_v36);
      let _664_v71;
      _664_v71 = _dafny.Map.Empty.slice().updateUnsafe(_623_v36,new BigNumber((_663_v70).length));
      let _665_v72;
      _665_v72 = _dafny.Seq.of((((_664_v71).contains(_623_v36)) ? ((_664_v71).get(_623_v36)) : ((_this).f20)), p0, (_this).f20, _this.f21);
      let _666_v73;
      _666_v73 = new _dafny.CodePoint('u'.codePointAt(0));
      let _667_v74;
      _667_v74 = _module.D3.create_DC8(_665_v72, p2, _623_v36, _666_v73);
      let _668_v75;
      _668_v75 = _module.D5.create_DC11(p0);
      (globalState).f6 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_667_v74).dtor_cf9).length))), (_668_v75).dtor_cf14);
      let _669_i9;
      _669_i9 = _dafny.ZERO;
      L4: {
        while (((_this).f20).isLessThanOrEqualTo(p0)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_669_i9)) {
              break L4;
            }
            _669_i9 = (_669_i9).plus(_dafny.ONE);
            let _670_v76;
            let _nw113 = Array((new BigNumber(7)).toNumber()).fill(false);
            _670_v76 = _nw113;
            let _index115 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length));
            (_670_v76)[_index115] = _623_v36;
            let _671_v77;
            _671_v77 = _dafny.MultiSet.fromElements(_623_v36, false);
            let _672_v78;
            _672_v78 = _dafny.MultiSet.fromElements(_665_v72);
            let _index116 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length));
            (_670_v76)[_index116] = !(((p1).isLessThan(new BigNumber((_671_v77).cardinality()))) && ((_672_v78).IsDisjointFrom(_dafny.MultiSet.fromElements(_665_v72))));
            let _673_v79;
            let _init18 = function (_674_i10) {
              return (_674_i10).minus((_this).f20);
            };
            let _nw114 = Array((new BigNumber(17)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw114.length); _i0_18++) {
              _nw114[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _673_v79 = _nw114;
            let _index117 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_673_v79).length));
            (_673_v79)[_index117] = (_dafny.ZERO).minus(p1);
            let _index118 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_673_v79).length));
            (_673_v79)[_index118] = _module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(_this.f21));
            let _675_v80;
            _675_v80 = _module.D3.create_DC6(_673_v79);
            let _676_v81;
            let _nw115 = Array((new BigNumber(10)).toNumber());
            _nw115[(_dafny.ZERO).toNumber()] = _673_v79;
            _nw115[(_dafny.ONE).toNumber()] = _673_v79;
            _nw115[(new BigNumber(2)).toNumber()] = _673_v79;
            _nw115[(new BigNumber(3)).toNumber()] = _673_v79;
            _nw115[(new BigNumber(4)).toNumber()] = _673_v79;
            _nw115[(new BigNumber(5)).toNumber()] = ((!((_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))])) ? (_673_v79) : (_673_v79));
            _nw115[(new BigNumber(6)).toNumber()] = (_675_v80).dtor_cf6;
            _nw115[(new BigNumber(7)).toNumber()] = _673_v79;
            _nw115[(new BigNumber(8)).toNumber()] = _673_v79;
            _nw115[(new BigNumber(9)).toNumber()] = _673_v79;
            _676_v81 = _nw115;
            let _677_v82;
            _677_v82 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, _623_v36)).length),_673_v79);
            let _index119 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_676_v81).length));
            (_676_v81)[_index119] = (((_677_v82).contains(new BigNumber((_665_v72).length))) ? ((_677_v82).get(new BigNumber((_665_v72).length))) : (_673_v79));
            let _index120 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_676_v81).length));
            (_676_v81)[_index120] = _673_v79;
            let _678_v83;
            _678_v83 = _dafny.MultiSet.fromElements(_671_v77, _671_v77, _671_v77);
            let _679_v84;
            _679_v84 = _dafny.MultiSet.fromElements(_624_v37);
            let _680_v85;
            _680_v85 = _dafny.MultiSet.fromElements(p0);
            let _681_v86;
            _681_v86 = _dafny.Seq.of(_671_v77);
            let _682_v87;
            _682_v87 = _dafny.Seq.of(_678_v83);
            let _683_v88;
            _683_v88 = _dafny.Set.fromElements((_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))], (_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))], (_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))]);
            let _684_v89;
            let _nw116 = Array((new BigNumber(22)).toNumber());
            _nw116[(_dafny.ZERO).toNumber()] = _678_v83;
            _nw116[(_dafny.ONE).toNumber()] = _module.__default.fm30(!((_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))]), _dafny.Set.fromElements(new BigNumber((_679_v84).cardinality())), globalState);
            _nw116[(new BigNumber(2)).toNumber()] = _678_v83;
            _nw116[(new BigNumber(3)).toNumber()] = _678_v83;
            _nw116[(new BigNumber(4)).toNumber()] = _678_v83;
            _nw116[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(_671_v77);
            _nw116[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))]))).Difference(_678_v83);
            _nw116[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(562)), ((_685_v36) => function (_686_i11) {
              return _dafny.MultiSet.fromElements(_685_v36);
            })(_623_v36)));
            _nw116[(new BigNumber(8)).toNumber()] = (_678_v83).Difference(_dafny.MultiSet.fromElements(_671_v77, _671_v77));
            _nw116[(new BigNumber(9)).toNumber()] = ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(_623_v36)), _671_v77, _671_v77, _671_v77))).update(_dafny.MultiSet.fromElements((_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))], (_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))], _623_v36, (_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))], (_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))]), _module.__default.abs((((_680_v85).contains(_this.f21)) ? ((_680_v85).get(_this.f21)) : ((_this).f20))))).Difference(_678_v83);
            _nw116[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.FromArray(_681_v86)).Union(_678_v83);
            _nw116[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.FromArray(_681_v86);
            _nw116[(new BigNumber(12)).toNumber()] = _678_v83;
            _nw116[(new BigNumber(13)).toNumber()] = (_682_v87)[_module.__default.safeIndex(new BigNumber((_683_v88).length), new BigNumber((_682_v87).length))];
            _nw116[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(_671_v77);
            _nw116[(new BigNumber(15)).toNumber()] = _678_v83;
            _nw116[(new BigNumber(16)).toNumber()] = _678_v83;
            _nw116[(new BigNumber(17)).toNumber()] = _678_v83;
            _nw116[(new BigNumber(18)).toNumber()] = _678_v83;
            _nw116[(new BigNumber(19)).toNumber()] = (_678_v83).update(_671_v77, _module.__default.abs(new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex((_673_v79)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_673_v79).length))], new BigNumber((p2).length)), _666_v73)).length)));
            _nw116[(new BigNumber(20)).toNumber()] = (((_670_v76)[_module.__default.safeIndex(new BigNumber(386), new BigNumber((_670_v76).length))]) ? (_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(_623_v36)), _dafny.MultiSet.FromArray(_624_v37))) : (_dafny.MultiSet.fromElements(_671_v77, _671_v77)));
            _nw116[(new BigNumber(21)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_681_v86, _681_v86));
            _684_v89 = _nw116;
            let _index121 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_684_v89).length));
            (_684_v89)[_index121] = (_678_v83).update(_671_v77, _module.__default.abs((_673_v79)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_673_v79).length))]));
            let _index122 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_684_v89).length));
            (_684_v89)[_index122] = _678_v83;
          }
        }
      }
      let _687_v90;
      let _nw117 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _687_v90 = _nw117;
      let _688_v91;
      _688_v91 = _module.D3.create_DC6(_687_v90);
      let _689_v92;
      _689_v92 = _module.D3.create_DC6((_688_v91).dtor_cf6);
      let _690_v93;
      _690_v93 = _dafny.Seq.of(_688_v91);
      let _691_v94;
      let _nw118 = new _module.C0();
      _nw118.__ctor(_690_v93, _623_v36);
      _691_v94 = _nw118;
      let _692_v95;
      _692_v95 = _dafny.Map.Empty.slice().updateUnsafe(_691_v94,p0);
      let _hi6 = new BigNumber(945);
      for (let _693_i12 = _module.__default.safeModuloInt(new BigNumber((_663_v70).length), new BigNumber((_692_v95).length)); _693_i12.isLessThan(_hi6); _693_i12 = _693_i12.plus(_dafny.ONE)) {
        _663_v70 = (_663_v70).update(!(_module.__default.fm2(globalState)), (_691_v94).f23);
        let _694_v96;
        _694_v96 = _dafny.Set.fromElements((_691_v94).f23);
        let _695_v97;
        _695_v97 = _module.D0.create_DC1((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p1), _dafny.Seq.of(new BigNumber(218), p0, _this.f21, new BigNumber((_694_v96).length)))).length)));
        let _pat_let_tv14 = _623_v36;
        let _pat_let_tv15 = globalState;
        _695_v97 = function (_pat_let26_0) {
          return function (_696_dt__update__tmp_h2) {
            return function (_pat_let27_0) {
              return function (_697_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_697_dt__update_hcf1_h0);
              }(_pat_let27_0);
            }(_module.__default.fm1(_pat_let_tv14, _pat_let_tv15));
          }(_pat_let26_0);
        }(_module.D0.create_DC1((_this).f20));
        let _698_v98;
        let _nw119 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _698_v98 = _nw119;
        let _index123 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_698_v98).length));
        (_698_v98)[_index123] = p3;
        let _index124 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_698_v98).length));
        (_698_v98)[_index124] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), function (_699_i13) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }), p3), p3);
        if (_623_v36) {
          (globalState).f6 = _this.f21;
          let _700_v99;
          let _nw120 = new _module.C0();
          _nw120.__ctor((_691_v94).f22, _623_v36);
          _700_v99 = _nw120;
          (globalState).f1 = (_691_v94).f23;
          let _701_v100;
          _701_v100 = _dafny.MultiSet.fromElements(((_691_v94).f23) === ((_700_v99).f23));
          let _702_v102;
          _702_v102 = _dafny.Map.Empty.slice().updateUnsafe(_664_v71,(_691_v94).f23);
          let _703_v103;
          let _nw121 = Array((new BigNumber(26)).toNumber()).fill([]);
          _703_v103 = _nw121;
          let _704_v104;
          _704_v104 = _dafny.Seq.of(_703_v103, _703_v103, _703_v103);
          let _705_v105;
          _705_v105 = _module.D5.create_DC10((_704_v104)[_module.__default.safeIndex(p1, new BigNumber((_704_v104).length))]);
          let _index125 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_698_v98).length));
          let _rhs86 = (_dafny.ZERO).minus(((_this).f20).minus((_this.f21).plus(_693_i12)));
          let _rhs87 = (_dafny.ZERO).minus(new BigNumber((function () {
            let _coll35 = new _dafny.Map();
            for (const _compr_35 of (_702_v102).Keys.Elements) {
              let _706_v101 = _compr_35;
              if ((_702_v102).contains(_706_v101)) {
                _coll35.push([_706_v101,_dafny.Seq.IsPrefixOf(p3, _dafny.Seq.UnicodeFromString("fsri"))]);
              }
            }
            return _coll35;
          }()).length));
          let _rhs88 = (_701_v100).update(_623_v36, _module.__default.abs((_dafny.ZERO).minus((_this).fm15(globalState))));
          let _rhs89 = (_698_v98)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_698_v98).length))];
          let _rhs90 = (((_module.D7.create_DC17((_module.__default.fm18(_this.f21, _665_v72, globalState)).dtor_cf0, (_691_v94).fm17(_693_i12, p0, _623_v36, globalState), _700_v99, _705_v105, false)).dtor_cf24) ? ((((_691_v94).f23) ? ((_700_v99).f23) : ((_691_v94).f23))) : ((_700_v99).f23));
          let _lhs75 = globalState;
          let _lhs76 = globalState;
          let _lhs77 = _698_v98;
          let _lhs78 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_698_v98).length));
          let _lhs79 = globalState;
          _lhs75.f6 = _rhs86;
          _lhs76.f6 = _rhs87;
          _701_v100 = _rhs88;
          _lhs77[_lhs78] = _rhs89;
          _lhs79.f1 = _rhs90;
          let _707_v106;
          _707_v106 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(987),p2);
          _707_v106 = (_707_v106).update((_this).f20, (_698_v98)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_698_v98).length))]);
        } else {
          let _708_v107;
          _708_v107 = (_664_v71).Merge(_664_v71);
          _708_v107 = _708_v107;
          let _709_v108;
          let _init19 = ((_710_v94) => function (_711_i14) {
            return (_710_v94).f23;
          })(_691_v94);
          let _nw122 = Array((new BigNumber(4)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw122.length); _i0_19++) {
            _nw122[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _709_v108 = _nw122;
          let _712_v109;
          _712_v109 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_624_v37).length),_709_v108);
          _712_v109 = ((_712_v109).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_709_v108))).Merge(_712_v109);
          let _713_v110;
          _713_v110 = _module.D0.create_DC0(!(!(false)));
          _712_v109 = (_712_v109).update(new BigNumber((_dafny.Set.fromElements((_this).fm16(_713_v110, (_691_v94).f23, globalState))).length), _709_v108);
          (globalState).f1 = (_module.__default.safeDivisionInt(new BigNumber((_665_v72).length), _this.f21)).isEqualTo(_module.__default.fm1((_691_v94).f23, globalState));
          r1 = _dafny.Seq.Concat(_module.__default.fm21(globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-971)), function (_714_i15) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }));
        }
      }
      r0 = !(_623_v36);
      r1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("d"), p2), p2);
      return [r0, r1];
    }
    m14(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _715_v0;
      let _nw123 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _715_v0 = _nw123;
      let _716_v1;
      _716_v1 = _module.D3.create_DC6(_715_v0);
      let _717_v2;
      _717_v2 = _dafny.Set.fromElements(p1);
      let _718_v3;
      let _719_v4;
      let _720_v5;
      let _out9;
      let _out10;
      let _out11;
      let _outcollector3 = (_this).m15(_716_v1, p0, (new BigNumber((_717_v2).length)).multipliedBy(p2), p2, globalState);
      _out9 = _outcollector3[0];
      _out10 = _outcollector3[1];
      _out11 = _outcollector3[2];
      _718_v3 = _out9;
      _719_v4 = _out10;
      _720_v5 = _out11;
      let _721_v6;
      _721_v6 = _dafny.Map.Empty.slice().updateUnsafe(true,p2);
      let _722_v7;
      _722_v7 = _721_v6;
      let _source7 = _722_v7;
      let _723___mcc_h0 = _source7;
      let _724_cf12 = _723___mcc_h0;
      let _725_v8;
      _725_v8 = _dafny.Map.Empty.slice().updateUnsafe(_718_v3,_718_v3);
      let _726_v9;
      let _nw124 = new _module.C1();
      _nw124.__ctor(new BigNumber((_725_v8).length));
      _726_v9 = _nw124;
      (globalState).f6 = p2;
      let _index126 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_715_v0).length));
      (_715_v0)[_index126] = new BigNumber(247);
      let _index127 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_715_v0).length));
      (_715_v0)[_index127] = (_dafny.ZERO).minus((_this).f20);
      if ((_this.f21).isLessThan(new BigNumber(((_725_v8).update(_718_v3, p0)).length))) {
        let _727_v10;
        _727_v10 = _dafny.Seq.of(_716_v1);
        let _728_v11;
        let _nw125 = Array((new BigNumber(20)).toNumber());
        _nw125[(_dafny.ZERO).toNumber()] = _718_v3;
        _nw125[(_dafny.ONE).toNumber()] = false;
        _nw125[(new BigNumber(2)).toNumber()] = p1;
        _nw125[(new BigNumber(3)).toNumber()] = _718_v3;
        _nw125[(new BigNumber(4)).toNumber()] = p0;
        _nw125[(new BigNumber(5)).toNumber()] = true;
        _nw125[(new BigNumber(6)).toNumber()] = true;
        _nw125[(new BigNumber(7)).toNumber()] = _718_v3;
        _nw125[(new BigNumber(8)).toNumber()] = p0;
        _nw125[(new BigNumber(9)).toNumber()] = p1;
        _nw125[(new BigNumber(10)).toNumber()] = false;
        _nw125[(new BigNumber(11)).toNumber()] = _718_v3;
        _nw125[(new BigNumber(12)).toNumber()] = _718_v3;
        _nw125[(new BigNumber(13)).toNumber()] = !(_718_v3);
        _nw125[(new BigNumber(14)).toNumber()] = true;
        _nw125[(new BigNumber(15)).toNumber()] = p1;
        _nw125[(new BigNumber(16)).toNumber()] = !(false);
        _nw125[(new BigNumber(17)).toNumber()] = p1;
        _nw125[(new BigNumber(18)).toNumber()] = p1;
        _nw125[(new BigNumber(19)).toNumber()] = _718_v3;
        _728_v11 = _nw125;
        let _729_v12;
        _729_v12 = _dafny.Map.Empty.slice().updateUnsafe(_728_v11,_728_v11);
        let _730_v13;
        let _nw126 = new _module.C0();
        _nw126.__ctor(_727_v10, (new BigNumber((_729_v12).length)).isLessThan(_this.f21));
        _730_v13 = _nw126;
        let _731_v14;
        _731_v14 = _dafny.Seq.UnicodeFromString("sdistbhg");
        _731_v14 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-230)), function (_732_i0) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        });
        let _733_v15;
        _733_v15 = _dafny.Set.fromElements(_719_v4, _719_v4);
        let _734_v16;
        _734_v16 = _dafny.Seq.of((_dafny.ZERO).minus(((_this).f20).plus(new BigNumber(252))), (_715_v0)[_module.__default.safeIndex(new BigNumber(500), new BigNumber((_715_v0).length))], (_726_v9).f24, (new BigNumber((_733_v15).length)).minus(p2), (_726_v9).f24);
        _734_v16 = _734_v16;
        let _735_v17;
        _735_v17 = _dafny.Seq.of(_717_v2, _717_v2);
        let _736_v18;
        let _nw127 = Array((new BigNumber(26)).toNumber());
        _nw127[(_dafny.ZERO).toNumber()] = (_717_v2).Difference(_717_v2);
        _nw127[(_dafny.ONE).toNumber()] = (_717_v2).Union(_717_v2);
        _nw127[(new BigNumber(2)).toNumber()] = (_717_v2).Difference(_dafny.Set.fromElements((_730_v13).f23, (_730_v13).f23));
        _nw127[(new BigNumber(3)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(4)).toNumber()] = (_717_v2).Intersect(_dafny.Set.fromElements(p0));
        _nw127[(new BigNumber(5)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(6)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(7)).toNumber()] = (_717_v2).Intersect(_717_v2);
        _nw127[(new BigNumber(8)).toNumber()] = _module.__default.fm31(_this.f21, p0, globalState);
        _nw127[(new BigNumber(9)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(10)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(11)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(false, true);
        _nw127[(new BigNumber(13)).toNumber()] = ((_718_v3) ? (_dafny.Set.fromElements(_718_v3)) : (_dafny.Set.fromElements(p0)));
        _nw127[(new BigNumber(14)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(15)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(16)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(17)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(18)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(19)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(20)).toNumber()] = (_735_v17)[_module.__default.safeIndex((_726_v9).f24, new BigNumber((_735_v17).length))];
        _nw127[(new BigNumber(21)).toNumber()] = (_717_v2).Difference(_dafny.Set.fromElements(true, (_730_v13).f23));
        _nw127[(new BigNumber(22)).toNumber()] = _717_v2;
        _nw127[(new BigNumber(23)).toNumber()] = (_717_v2).Difference(_dafny.Set.fromElements(_718_v3, p1, _718_v3));
        _nw127[(new BigNumber(24)).toNumber()] = (_717_v2).Difference(_717_v2);
        _nw127[(new BigNumber(25)).toNumber()] = _717_v2;
        _736_v18 = _nw127;
        let _index128 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_736_v18).length));
        (_736_v18)[_index128] = (((_730_v13).f23) ? (_717_v2) : (_717_v2));
        let _index129 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_736_v18).length));
        (_736_v18)[_index129] = (_717_v2).Intersect(_717_v2);
        _718_v3 = (_730_v13).f23;
      } else {
        let _737_v19;
        let _nw128 = Array((new BigNumber(11)).toNumber());
        _nw128[(_dafny.ZERO).toNumber()] = _718_v3;
        _nw128[(_dafny.ONE).toNumber()] = _718_v3;
        _nw128[(new BigNumber(2)).toNumber()] = p1;
        _nw128[(new BigNumber(3)).toNumber()] = _718_v3;
        _nw128[(new BigNumber(4)).toNumber()] = (_this.f21).isLessThanOrEqualTo((_726_v9).f24);
        _nw128[(new BigNumber(5)).toNumber()] = (_717_v2).IsSubsetOf(_717_v2);
        _nw128[(new BigNumber(6)).toNumber()] = p1;
        _nw128[(new BigNumber(7)).toNumber()] = ((_this).f20).isLessThanOrEqualTo((_715_v0)[_module.__default.safeIndex(new BigNumber(500), new BigNumber((_715_v0).length))]);
        _nw128[(new BigNumber(8)).toNumber()] = (_718_v3) && (p1);
        _nw128[(new BigNumber(9)).toNumber()] = ((_715_v0)[_module.__default.safeIndex(new BigNumber(500), new BigNumber((_715_v0).length))]).isEqualTo((_715_v0)[_module.__default.safeIndex(new BigNumber(500), new BigNumber((_715_v0).length))]);
        _nw128[(new BigNumber(10)).toNumber()] = _718_v3;
        _737_v19 = _nw128;
        let _index130 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_737_v19).length));
        (_737_v19)[_index130] = p0;
        let _index131 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_737_v19).length));
        (_737_v19)[_index131] = p0;
        let _index132 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_715_v0).length));
        (_715_v0)[_index132] = _this.f21;
        let _738_v20;
        _738_v20 = _dafny.Seq.of(_716_v1, _716_v1);
        let _739_v21;
        let _nw129 = new _module.C0();
        _nw129.__ctor(((false) ? (_738_v20) : (_738_v20)), _718_v3);
        _739_v21 = _nw129;
        let _740_v22;
        let _nw130 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
        _740_v22 = _nw130;
        let _741_v23;
        _741_v23 = _dafny.Seq.of(p3, _dafny.Seq.UnicodeFromString("bt"));
        let _index133 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_740_v22).length));
        (_740_v22)[_index133] = _741_v23;
        let _index134 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_740_v22).length));
        (_740_v22)[_index134] = _741_v23;
        _725_v8 = (_725_v8).update((_739_v21).f23, p0);
      }
      let _742_i1;
      _742_i1 = _dafny.ZERO;
      L5: {
        while (_module.__default.fm2(globalState)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_742_i1)) {
              break L5;
            }
            _742_i1 = (_742_i1).plus(_dafny.ONE);
            let _743_v24;
            _743_v24 = _dafny.MultiSet.fromElements(true);
            let _744_v25;
            _744_v25 = _dafny.Map.Empty.slice().updateUnsafe((_743_v24).IsSubsetOf(_dafny.MultiSet.fromElements(!(p1))),p3);
            (_this).f21 = new BigNumber(((((_744_v25).contains(!(p2).isEqualTo(p2))) ? ((_744_v25).get(!(p2).isEqualTo(p2))) : (_dafny.Seq.Concat(p3, p3)))).length);
            let _745_v26;
            _745_v26 = _dafny.Seq.of(_716_v1, _716_v1);
            let _746_v27;
            _746_v27 = _dafny.MultiSet.fromElements(p2);
            let _747_v28;
            let _nw131 = new _module.C0();
            _nw131.__ctor(_745_v26, !(_746_v27).contains(p2));
            _747_v28 = _nw131;
            let _748_v29;
            _748_v29 = _dafny.Seq.of(p1);
            let _749_v30;
            _749_v30 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("hnxddpbbf"), _dafny.Seq.UnicodeFromString("mhsnvto"), p3);
            let _750_v31;
            _750_v31 = _module.D3.create_DC7((_this).f20);
            let _751_v32;
            _751_v32 = _dafny.Set.fromElements((_750_v31).dtor_cf7);
            let _752_v33;
            _752_v33 = _module.D0.create_DC0(false);
            let _753_v34;
            _753_v34 = _module.D8.create_DC21(p3, _717_v2);
            let _754_v35;
            let _nw132 = Array((new BigNumber(28)).toNumber());
            _nw132[(_dafny.ZERO).toNumber()] = !(_dafny.Seq.contains(_module.__default.fm32((_dafny.ZERO).minus(new BigNumber((_748_v29).length)), p2, _718_v3, globalState), _this.f21));
            _nw132[(_dafny.ONE).toNumber()] = false;
            _nw132[(new BigNumber(2)).toNumber()] = (_747_v28).f23;
            _nw132[(new BigNumber(3)).toNumber()] = p1;
            _nw132[(new BigNumber(4)).toNumber()] = _dafny.Seq.contains(_748_v29, p1);
            _nw132[(new BigNumber(5)).toNumber()] = _718_v3;
            _nw132[(new BigNumber(6)).toNumber()] = (_747_v28).f23;
            _nw132[(new BigNumber(7)).toNumber()] = p0;
            _nw132[(new BigNumber(8)).toNumber()] = p0;
            _nw132[(new BigNumber(9)).toNumber()] = true;
            _nw132[(new BigNumber(10)).toNumber()] = true;
            _nw132[(new BigNumber(11)).toNumber()] = (_747_v28).f23;
            _nw132[(new BigNumber(12)).toNumber()] = (new BigNumber((_749_v30).length)).isLessThanOrEqualTo(p2);
            _nw132[(new BigNumber(13)).toNumber()] = p1;
            _nw132[(new BigNumber(14)).toNumber()] = true;
            _nw132[(new BigNumber(15)).toNumber()] = p0;
            _nw132[(new BigNumber(16)).toNumber()] = !(p0);
            _nw132[(new BigNumber(17)).toNumber()] = true;
            _nw132[(new BigNumber(18)).toNumber()] = p0;
            _nw132[(new BigNumber(19)).toNumber()] = (_751_v32).IsProperSubsetOf(_751_v32);
            _nw132[(new BigNumber(20)).toNumber()] = _718_v3;
            _nw132[(new BigNumber(21)).toNumber()] = (_this).fm16(_752_v33, _718_v3, globalState);
            _nw132[(new BigNumber(22)).toNumber()] = (_747_v28).f23;
            _nw132[(new BigNumber(23)).toNumber()] = (_717_v2).IsSubsetOf((_753_v34).dtor_cf35);
            _nw132[(new BigNumber(24)).toNumber()] = ((p1) ? (!(!(p0))) : ((_this).fm16(_module.D0.create_DC0(_718_v3), p0, globalState)));
            _nw132[(new BigNumber(25)).toNumber()] = p0;
            _nw132[(new BigNumber(26)).toNumber()] = (_747_v28).f23;
            _nw132[(new BigNumber(27)).toNumber()] = p0;
            _754_v35 = _nw132;
            let _index135 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_754_v35).length));
            (_754_v35)[_index135] = (p1) && (p1);
            let _index136 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_754_v35).length));
            (_754_v35)[_index136] = _718_v3;
            let _755_v36;
            let _nw133 = new _module.C1();
            _nw133.__ctor(((_dafny.ZERO).minus((_this).f20)).plus(_this.f21));
            _755_v36 = _nw133;
          }
        }
      }
      let _rhs91 = _721_v6;
      let _rhs92 = (_dafny.ZERO).minus(_this.f21);
      _721_v6 = _rhs91;
      r0 = _rhs92;
      if (((_this).f20).isLessThan((p2).plus((_this).f20))) {
        let _756_v37;
        let _nw134 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Set.Empty);
        _756_v37 = _nw134;
        let _index137 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_756_v37).length));
        (_756_v37)[_index137] = _717_v2;
        let _index138 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_756_v37).length));
        (_756_v37)[_index138] = (((p0) ? (_717_v2) : (_717_v2))).Intersect(_717_v2);
        let _757_v38;
        _757_v38 = new _dafny.CodePoint('y'.codePointAt(0));
        _757_v38 = _757_v38;
        _717_v2 = _717_v2;
        if (p0) {
          (globalState).f6 = p2;
          (globalState).f6 = p2;
          let _758_v39;
          _758_v39 = _module.D3.create_DC7(p2);
          let _759_v40;
          _759_v40 = _dafny.Seq.of(p2, p2);
          let _pat_let_tv16 = _759_v40;
          let _760_v41;
          let _nw135 = Array((new BigNumber(19)).toNumber());
          _nw135[(_dafny.ZERO).toNumber()] = _758_v39;
          _nw135[(_dafny.ONE).toNumber()] = _758_v39;
          _nw135[(new BigNumber(2)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(3)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(4)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(5)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(6)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(7)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(8)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(9)).toNumber()] = _module.D3.create_DC7(new BigNumber(236));
          _nw135[(new BigNumber(10)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(11)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(12)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(13)).toNumber()] = _module.D3.create_DC7((_dafny.ZERO).minus(p2));
          _nw135[(new BigNumber(14)).toNumber()] = _module.D3.create_DC7((_this).f20);
          _nw135[(new BigNumber(15)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(16)).toNumber()] = _758_v39;
          _nw135[(new BigNumber(17)).toNumber()] = function (_pat_let28_0) {
            return function (_761_dt__update__tmp_h0) {
              return function (_pat_let29_0) {
                return function (_762_dt__update_hcf7_h0) {
                  return _module.D3.create_DC7(_762_dt__update_hcf7_h0);
                }(_pat_let29_0);
              }((_dafny.ZERO).minus(new BigNumber((_pat_let_tv16).length)));
            }(_pat_let28_0);
          }(_758_v39);
          _nw135[(new BigNumber(18)).toNumber()] = _758_v39;
          _760_v41 = _nw135;
          let _rhs93 = (_dafny.ZERO).minus((p2).minus((p2).minus(new BigNumber(578))));
          let _rhs94 = ((p1) ? (((p1) ? (_760_v41) : (_760_v41))) : (_760_v41));
          let _lhs80 = globalState;
          _lhs80.f6 = _rhs93;
          _760_v41 = _rhs94;
          let _763_v42;
          _763_v42 = _dafny.Seq.of(_716_v1);
          let _764_v43;
          let _nw136 = new _module.C0();
          _nw136.__ctor(_763_v42, _dafny.Seq.contains(p3, _757_v38));
          _764_v43 = _nw136;
          let _765_v44;
          _765_v44 = _dafny.Map.Empty.slice().updateUnsafe(_this.f21,p1);
          let _766_v45;
          let _nw137 = Array((new BigNumber(3)).toNumber());
          _nw137[(_dafny.ZERO).toNumber()] = _765_v44;
          _nw137[(_dafny.ONE).toNumber()] = _765_v44;
          _nw137[(new BigNumber(2)).toNumber()] = _765_v44;
          _766_v45 = _nw137;
          let _767_v46;
          _767_v46 = _module.D0.create_DC0(p1);
          let _768_v47;
          _768_v47 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).fm16(_767_v46, (_764_v43).f23, globalState));
          let _769_v48;
          let _nw138 = Array((new BigNumber(19)).toNumber());
          _nw138[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_759_v40, _module.__default.safeIndex(new BigNumber(309), new BigNumber((_759_v40).length)), (_dafny.ZERO).minus(p2));
          _nw138[(_dafny.ONE).toNumber()] = _759_v40;
          _nw138[(new BigNumber(2)).toNumber()] = _module.__default.fm32((_this).f20, new BigNumber((_768_v47).length), p0, globalState);
          _nw138[(new BigNumber(3)).toNumber()] = _module.__default.fm32(_this.f21, (_this).f20, p1, globalState);
          _nw138[(new BigNumber(4)).toNumber()] = _759_v40;
          _nw138[(new BigNumber(5)).toNumber()] = _759_v40;
          _nw138[(new BigNumber(6)).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus(p2), _this.f21, p2);
          _nw138[(new BigNumber(7)).toNumber()] = _759_v40;
          _nw138[(new BigNumber(8)).toNumber()] = _759_v40;
          _nw138[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_module.__default.fm32(new BigNumber((_module.__default.fm32(_this.f21, _this.f21, p0, globalState)).length), _this.f21, p1, globalState), _module.__default.safeIndex(new BigNumber(485), new BigNumber((_module.__default.fm32(new BigNumber((_module.__default.fm32(_this.f21, _this.f21, p0, globalState)).length), _this.f21, p1, globalState)).length)), _this.f21);
          _nw138[(new BigNumber(10)).toNumber()] = _759_v40;
          _nw138[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), function (_770_i2) {
            return (_this).f20;
          })).length)), _module.__default.safeIndex((_this).f20, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(307)), function (_771_i2) {
            return (_this).f20;
          })).length))).length)), (_758_v39).dtor_cf7);
          _nw138[(new BigNumber(12)).toNumber()] = _759_v40;
          _nw138[(new BigNumber(13)).toNumber()] = _759_v40;
          _nw138[(new BigNumber(14)).toNumber()] = _759_v40;
          _nw138[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_759_v40, _759_v40);
          _nw138[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_759_v40, _dafny.Seq.of(p2, (_this).f20));
          _nw138[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(920)), ((_772_v38, _773_p0) => function (_774_i3) {
            return new BigNumber(((_module.D8.create_DC21(_dafny.Seq.Create(_module.__default.abs(new BigNumber(167)), ((_775_v38) => function (_776_i4) {
  return _775_v38;
})(_772_v38)), _dafny.Set.fromElements(_773_p0))).dtor_cf34).length);
          })(_757_v38, p0));
          _nw138[(new BigNumber(18)).toNumber()] = _759_v40;
          _769_v48 = _nw138;
          let _index139 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_769_v48).length));
          (_769_v48)[_index139] = _dafny.Seq.Concat(_759_v40, _759_v40);
          let _777_v49;
          _777_v49 = _module.D9.create_DC22(_766_v45);
          let _index140 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_756_v37).length));
          let _index141 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_769_v48).length));
          let _rhs95 = (_777_v49).dtor_cf36;
          let _rhs96 = (_756_v37)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_756_v37).length))];
          let _rhs97 = _dafny.Seq.update(_759_v40, _module.__default.safeIndex(_this.f21, new BigNumber((_759_v40).length)), (_this).f20);
          let _rhs98 = p0;
          let _lhs81 = _756_v37;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_756_v37).length));
          let _lhs83 = _769_v48;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_769_v48).length));
          let _lhs85 = globalState;
          _766_v45 = _rhs95;
          _lhs81[_lhs82] = _rhs96;
          _lhs83[_lhs84] = _rhs97;
          _lhs85.f1 = _rhs98;
        } else {
          _718_v3 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(p3, p3), p3);
          let _778_v50;
          _778_v50 = _dafny.Seq.of(_716_v1);
          let _779_v51;
          let _nw139 = new _module.C0();
          _nw139.__ctor(_778_v50, p0);
          _779_v51 = _nw139;
          _719_v4 = (_719_v4).update(new BigNumber(-112), (_dafny.ZERO).minus(new BigNumber(-286)));
          let _780_v52;
          let _init20 = ((_781_p0) => function (_782_i5) {
            return _781_p0;
          })(p0);
          let _nw140 = Array((new BigNumber(7)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw140.length); _i0_20++) {
            _nw140[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _780_v52 = _nw140;
          let _index142 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_780_v52).length));
          (_780_v52)[_index142] = p0;
          let _index143 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_780_v52).length));
          (_780_v52)[_index143] = (_779_v51).f23;
          let _783_v53;
          _783_v53 = _dafny.Set.fromElements((_this).f20);
          let _784_v54;
          _784_v54 = _dafny.Map.Empty.slice().updateUnsafe(p3,_783_v53);
          let _785_v55;
          _785_v55 = _dafny.Seq.of((_779_v51).f23, (_780_v52)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_780_v52).length))], false, (_779_v51).f23, true);
          let _786_v56;
          _786_v56 = _dafny.Seq.of(new BigNumber((_785_v55).length), (_this).f20);
          _784_v54 = (_784_v54).update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tff"), _dafny.Seq.UnicodeFromString("xubovus")), _dafny.Set.fromElements((_786_v56)[_module.__default.safeIndex(p2, new BigNumber((_786_v56).length))], p2));
        }
        let _787_v57;
        _787_v57 = _dafny.Seq.of(_716_v1);
        let _788_v58;
        let _nw141 = new _module.C0();
        _nw141.__ctor(_787_v57, p0);
        _788_v58 = _nw141;
        let _789_v59;
        let _nw142 = Array((new BigNumber(13)).toNumber()).fill([]);
        _789_v59 = _nw142;
        let _790_v60;
        _790_v60 = _module.D5.create_DC10(_789_v59);
        let _791_v61;
        _791_v61 = _module.D7.create_DC17(p0, p0, _788_v58, _790_v60, _718_v3);
        _718_v3 = (_791_v61).dtor_cf28;
      } else {
        if (_718_v3) {
          (globalState).f1 = !(p1);
          let _792_v62;
          let _nw143 = Array((_dafny.ONE).toNumber()).fill([]);
          _792_v62 = _nw143;
          let _793_v63;
          _793_v63 = _dafny.MultiSet.fromElements(_this.f21);
          let _794_v64;
          let _nw144 = new _module.C1();
          _nw144.__ctor(new BigNumber((_793_v63).cardinality()));
          _794_v64 = _nw144;
          let _795_v65;
          _795_v65 = _module.D7.create_DC15(_794_v64);
          let _pat_let_tv17 = _794_v64;
          let _796_v66;
          let _nw145 = Array((new BigNumber(12)).toNumber());
          _nw145[(_dafny.ZERO).toNumber()] = _module.D7.create_DC15(_794_v64);
          _nw145[(_dafny.ONE).toNumber()] = _module.D7.create_DC15(_794_v64);
          _nw145[(new BigNumber(2)).toNumber()] = _795_v65;
          _nw145[(new BigNumber(3)).toNumber()] = _795_v65;
          _nw145[(new BigNumber(4)).toNumber()] = _module.D7.create_DC15(_794_v64);
          _nw145[(new BigNumber(5)).toNumber()] = _795_v65;
          _nw145[(new BigNumber(6)).toNumber()] = _795_v65;
          _nw145[(new BigNumber(7)).toNumber()] = _795_v65;
          _nw145[(new BigNumber(8)).toNumber()] = function (_pat_let30_0) {
            return function (_797_dt__update__tmp_h1) {
              return function (_pat_let31_0) {
                return function (_798_dt__update_hcf20_h0) {
                  return _module.D7.create_DC15(_798_dt__update_hcf20_h0);
                }(_pat_let31_0);
              }(_pat_let_tv17);
            }(_pat_let30_0);
          }(_795_v65);
          _nw145[(new BigNumber(9)).toNumber()] = _795_v65;
          _nw145[(new BigNumber(10)).toNumber()] = _795_v65;
          _nw145[(new BigNumber(11)).toNumber()] = _795_v65;
          _796_v66 = _nw145;
          let _799_v67;
          let _nw146 = Array((new BigNumber(2)).toNumber());
          _nw146[(_dafny.ZERO).toNumber()] = _796_v66;
          _nw146[(_dafny.ONE).toNumber()] = _796_v66;
          _799_v67 = _nw146;
          let _index144 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_792_v62).length));
          (_792_v62)[_index144] = _799_v67;
          let _800_v68;
          let _nw147 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _800_v68 = _nw147;
          let _801_v69;
          _801_v69 = new _dafny.CodePoint('w'.codePointAt(0));
          let _index145 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_800_v68).length));
          (_800_v68)[_index145] = _801_v69;
          let _802_v70;
          _802_v70 = _module.D10.create_DC26(_799_v67);
          let _pat_let_tv18 = _799_v67;
          let _index146 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_792_v62).length));
          let _index147 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_800_v68).length));
          let _rhs99 = (function (_pat_let32_0) {
            return function (_803_dt__update__tmp_h2) {
              return function (_pat_let33_0) {
                return function (_804_dt__update_hcf45_h0) {
                  return _module.D10.create_DC26(_804_dt__update_hcf45_h0);
                }(_pat_let33_0);
              }(_pat_let_tv18);
            }(_pat_let32_0);
          }(_802_v70)).dtor_cf45;
          let _rhs100 = _801_v69;
          let _rhs101 = p0;
          let _lhs86 = _792_v62;
          let _lhs87 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_792_v62).length));
          let _lhs88 = _800_v68;
          let _lhs89 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_800_v68).length));
          let _lhs90 = globalState;
          _lhs86[_lhs87] = _rhs99;
          _lhs88[_lhs89] = _rhs100;
          _lhs90.f1 = _rhs101;
          (globalState).f1 = p0;
          (globalState).f1 = (true) || (true);
          (_this).f21 = _module.__default.safeModuloInt((_module.__default.fm1(p1, globalState)).multipliedBy(_module.__default.fm1(p0, globalState)), _module.__default.safeModuloInt(_this.f21, _this.f21));
        } else {
          let _805_v71;
          _805_v71 = _dafny.Seq.of(_716_v1, _716_v1);
          let _806_v72;
          _806_v72 = _module.D11.create_DC30(_805_v71);
          let _807_v73;
          let _nw148 = new _module.C0();
          _nw148.__ctor((_806_v72).dtor_cf53, p1);
          _807_v73 = _nw148;
          let _808_v74;
          _808_v74 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f21,(_this).f20));
          let _809_v75;
          _809_v75 = _dafny.Seq.of((_this).f20, _this.f21);
          let _810_v76;
          _810_v76 = _dafny.MultiSet.fromElements(_809_v75, _module.__default.fm32((_this).f20, (_this).f20, p0, globalState));
          let _811_v77;
          let _nw149 = Array((new BigNumber(14)).toNumber()).fill(false);
          _811_v77 = _nw149;
          let _index148 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_811_v77).length));
          (_811_v77)[_index148] = !_dafny.Seq.contains(p3, new _dafny.CodePoint('t'.codePointAt(0)));
          let _812_v79;
          _812_v79 = _module.D12.create_DC34(function () {
  let _coll36 = new _dafny.Map();
  for (const _compr_36 of _dafny.IntegerRange(new BigNumber(-882), new BigNumber(664))) {
    let _813_v78 = _compr_36;
    if (((new BigNumber(-882)).isLessThanOrEqualTo(_813_v78)) && ((_813_v78).isLessThan(new BigNumber(664)))) {
      _coll36.push([(_813_v78).multipliedBy((_dafny.ZERO).minus(p2)),(_this).f20]);
    }
  }
  return _coll36;
}());
          let _814_v80;
          _814_v80 = _dafny.Map.Empty.slice().updateUnsafe((_807_v73).f23,(_807_v73).f23);
          let _815_v81;
          _815_v81 = new _dafny.CodePoint('a'.codePointAt(0));
          let _816_v82;
          _816_v82 = _dafny.Map.Empty.slice().updateUnsafe(_815_v81,_810_v76);
          let _817_v83;
          _817_v83 = _module.D11.create_DC32(p0, _this.f21, _718_v3);
          let _818_v84;
          _818_v84 = _module.D7.create_DC16(p1, (_817_v83).dtor_cf60, p0);
          let _pat_let_tv19 = _719_v4;
          let _index149 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_811_v77).length));
          let _rhs102 = (_808_v74).Difference(_dafny.MultiSet.fromElements((function (_pat_let34_0) {
            return function (_819_dt__update__tmp_h3) {
              return function (_pat_let35_0) {
                return function (_820_dt__update_hcf63_h0) {
                  return _module.D12.create_DC34(_820_dt__update_hcf63_h0);
                }(_pat_let35_0);
              }(_pat_let_tv19);
            }(_pat_let34_0);
          }(_812_v79)).dtor_cf63));
          let _rhs103 = ((p1) ? ((_814_v80).contains(p0)) : (((p0) ? (_718_v3) : ((_807_v73).f23))));
          let _rhs104 = ((_810_v76).update(_809_v75, _module.__default.abs((_dafny.ZERO).minus(_this.f21)))).Difference((((_816_v82).contains(_815_v81)) ? ((_816_v82).get(_815_v81)) : (_810_v76)));
          let _rhs105 = (_818_v84).dtor_cf23;
          let _lhs91 = globalState;
          let _lhs92 = _811_v77;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_811_v77).length));
          _808_v74 = _rhs102;
          _lhs91.f1 = _rhs103;
          _810_v76 = _rhs104;
          _lhs92[_lhs93] = _rhs105;
          let _821_v85;
          _821_v85 = _dafny.MultiSet.fromElements(p2, new BigNumber((_module.__default.fm33((_dafny.ZERO).minus(new BigNumber((_717_v2).length)), (_this).f20, true, globalState)).length), p2, _this.f21, p2);
          let _822_v87;
          _822_v87 = _module.D0.create_DC1((_this).f20);
          let _823_v88;
          _823_v88 = _module.D0.create_DC2(_822_v87);
          let _824_v89;
          _824_v89 = _module.D0.create_DC2((_823_v88).dtor_cf2);
          let _pat_let_tv20 = _822_v87;
          let _pat_let_tv21 = _822_v87;
          let _825_v90;
          let _nw150 = Array((new BigNumber(19)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = function (_pat_let36_0) {
            return function (_826_dt__update__tmp_h4) {
              return function (_pat_let37_0) {
                return function (_827_dt__update_hcf2_h0) {
                  return _module.D0.create_DC2(_827_dt__update_hcf2_h0);
                }(_pat_let37_0);
              }(_pat_let_tv20);
            }(_pat_let36_0);
          }(_823_v88);
          _nw150[(_dafny.ONE).toNumber()] = _module.D0.create_DC2(_822_v87);
          _nw150[(new BigNumber(2)).toNumber()] = _823_v88;
          _nw150[(new BigNumber(3)).toNumber()] = _824_v89;
          _nw150[(new BigNumber(4)).toNumber()] = _824_v89;
          _nw150[(new BigNumber(5)).toNumber()] = _823_v88;
          _nw150[(new BigNumber(6)).toNumber()] = _824_v89;
          _nw150[(new BigNumber(7)).toNumber()] = _823_v88;
          _nw150[(new BigNumber(8)).toNumber()] = _823_v88;
          _nw150[(new BigNumber(9)).toNumber()] = _823_v88;
          _nw150[(new BigNumber(10)).toNumber()] = _823_v88;
          _nw150[(new BigNumber(11)).toNumber()] = _823_v88;
          _nw150[(new BigNumber(12)).toNumber()] = _824_v89;
          _nw150[(new BigNumber(13)).toNumber()] = _823_v88;
          _nw150[(new BigNumber(14)).toNumber()] = function (_pat_let38_0) {
            return function (_828_dt__update__tmp_h5) {
              return function (_pat_let39_0) {
                return function (_829_dt__update_hcf2_h1) {
                  return _module.D0.create_DC2(_829_dt__update_hcf2_h1);
                }(_pat_let39_0);
              }(_pat_let_tv21);
            }(_pat_let38_0);
          }(_823_v88);
          _nw150[(new BigNumber(15)).toNumber()] = _823_v88;
          _nw150[(new BigNumber(16)).toNumber()] = _module.D0.create_DC2(_822_v87);
          _nw150[(new BigNumber(17)).toNumber()] = _824_v89;
          _nw150[(new BigNumber(18)).toNumber()] = _824_v89;
          _825_v90 = _nw150;
          _module.__default.m0(function () {
            let _coll37 = new _dafny.Set();
            for (const _compr_37 of (_821_v85).Elements) {
              let _830_v86 = _compr_37;
              if ((_821_v85).contains(_830_v86)) {
                _coll37.add((_830_v86).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-153)), function (_831_i6) {
                  return new _dafny.CodePoint('g'.codePointAt(0));
                })).length))).length)));
              }
            }
            return _coll37;
          }(), _825_v90, _715_v0, p3, globalState);
          let _832_v91;
          _832_v91 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("iug"));
          let _833_v92;
          _833_v92 = _module.D3.create_DC7(_this.f21);
          let _834_v93;
          _834_v93 = _dafny.Map.Empty.slice().updateUnsafe(_811_v77,true);
          let _835_v94;
          _835_v94 = _module.D10.create_DC27(new BigNumber((_814_v80).length), p3, p2, _833_v92, _834_v93);
          let _836_v95;
          _836_v95 = _module.D3.create_DC8(_809_v75, (_835_v94).dtor_cf47, p0, _815_v81);
          let _index150 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_811_v77).length));
          (_811_v77)[_index150] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("wefcbtsh"), _dafny.Seq.Concat(p3, (((_832_v91).contains((_807_v73).f23)) ? ((_832_v91).get((_807_v73).f23)) : ((_836_v95).dtor_cf9))));
          let _837_v96;
          let _nw151 = new _module.C1();
          _nw151.__ctor((_this).f20);
          _837_v96 = _nw151;
        }
        let _pat_let_tv22 = _715_v0;
        let _pat_let_tv23 = _715_v0;
        let _pat_let_tv24 = _715_v0;
        let _838_v97;
        let _nw152 = Array((new BigNumber(11)).toNumber());
        _nw152[(_dafny.ZERO).toNumber()] = _716_v1;
        _nw152[(_dafny.ONE).toNumber()] = _716_v1;
        _nw152[(new BigNumber(2)).toNumber()] = function (_pat_let40_0) {
          return function (_839_dt__update__tmp_h6) {
            return function (_pat_let41_0) {
              return function (_840_dt__update_hcf6_h0) {
                return _module.D3.create_DC6(_840_dt__update_hcf6_h0);
              }(_pat_let41_0);
            }(_pat_let_tv22);
          }(_pat_let40_0);
        }(_716_v1);
        _nw152[(new BigNumber(3)).toNumber()] = _module.D3.create_DC6(_715_v0);
        _nw152[(new BigNumber(4)).toNumber()] = _716_v1;
        _nw152[(new BigNumber(5)).toNumber()] = function (_pat_let42_0) {
          return function (_841_dt__update__tmp_h7) {
            return function (_pat_let43_0) {
              return function (_842_dt__update_hcf6_h1) {
                return _module.D3.create_DC6(_842_dt__update_hcf6_h1);
              }(_pat_let43_0);
            }(_pat_let_tv23);
          }(_pat_let42_0);
        }(_716_v1);
        _nw152[(new BigNumber(6)).toNumber()] = _716_v1;
        _nw152[(new BigNumber(7)).toNumber()] = function (_pat_let44_0) {
          return function (_843_dt__update__tmp_h8) {
            return function (_pat_let45_0) {
              return function (_844_dt__update_hcf6_h2) {
                return _module.D3.create_DC6(_844_dt__update_hcf6_h2);
              }(_pat_let45_0);
            }(_pat_let_tv24);
          }(_pat_let44_0);
        }(_716_v1);
        _nw152[(new BigNumber(8)).toNumber()] = _716_v1;
        _nw152[(new BigNumber(9)).toNumber()] = _716_v1;
        _nw152[(new BigNumber(10)).toNumber()] = _716_v1;
        _838_v97 = _nw152;
        let _845_v98;
        _845_v98 = _dafny.Set.fromElements(_838_v97, _838_v97, _838_v97, _838_v97);
        let _846_v99;
        _846_v99 = _module.D9.create_DC23(_845_v98, p3, _module.__default.fm2(globalState), _715_v0, _module.__default.fm29(_718_v3, (_this).f20, (_this).f20, p0, globalState));
        let _847_v100;
        _847_v100 = _dafny.Seq.of(!(false), _718_v3);
        let _848_v101;
        _848_v101 = _dafny.MultiSet.fromElements((_this).f20, new BigNumber(-776));
        let _849_v102;
        _849_v102 = _dafny.MultiSet.fromElements(p0, _module.__default.fm2(globalState));
        let _850_v103;
        _850_v103 = _module.D3.create_DC7((((_848_v101).contains(p2)) ? ((_848_v101).get(p2)) : (new BigNumber((_849_v102).cardinality()))));
        let _851_v104;
        let _init21 = ((_852_p0) => function (_853_i7) {
          return _852_p0;
        })(p0);
        let _nw153 = Array((new BigNumber(6)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw153.length); _i0_21++) {
          _nw153[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _851_v104 = _nw153;
        let _854_v105;
        _854_v105 = _dafny.Map.Empty.slice().updateUnsafe(_851_v104,p1);
        let _855_v106;
        _855_v106 = _module.D10.create_DC27((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_718_v3,!(p0))).length)).minus((_dafny.ZERO).minus(_this.f21)), _dafny.Seq.Concat(p3, (_846_v99).dtor_cf38), (new BigNumber((_847_v100).length)).plus(_this.f21), _850_v103, _854_v105);
        _855_v106 = _module.D10.create_DC27(((p0) ? (new BigNumber(964)) : ((_this).f20)), p3, (_this.f21).multipliedBy((_this).f20), _850_v103, _dafny.Map.Empty.slice().updateUnsafe(_851_v104,p1));
        r0 = _module.__default.safeModuloInt(new BigNumber(342), (_this.f21).multipliedBy((_this).f20));
        _717_v2 = _717_v2;
        let _index151 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_851_v104).length));
        (_851_v104)[_index151] = p1;
        let _index152 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_851_v104).length));
        (_851_v104)[_index152] = false;
      }
      let _856_v107;
      let _init22 = ((_857_p3) => function (_858_i8) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(370)), function (_859_i9) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }), _857_p3);
      })(p3);
      let _nw154 = Array((new BigNumber(15)).toNumber());
      for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw154.length); _i0_22++) {
        _nw154[_i0_22] = _init22(new BigNumber(_i0_22));
      }
      _856_v107 = _nw154;
      let _index153 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_856_v107).length));
      (_856_v107)[_index153] = p3;
      let _860_v108;
      _860_v108 = _dafny.MultiSet.fromElements(_this.f21);
      let _index154 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_856_v107).length));
      let _rhs106 = p3;
      let _rhs107 = (_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f21))).IsSubsetOf(_860_v108);
      let _rhs108 = _718_v3;
      let _lhs94 = _856_v107;
      let _lhs95 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_856_v107).length));
      let _lhs96 = globalState;
      _lhs94[_lhs95] = _rhs106;
      _718_v3 = _rhs107;
      _lhs96.f1 = _rhs108;
      r0 = (_this).f20;
      return r0;
    }
    m15(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let r2 = [];
      if (p1) {
        (globalState).f6 = (_this).fm15(globalState);
        let _861_v0;
        _861_v0 = _dafny.Seq.of(_module.__default.fm29(p1, p3, _this.f21, p1, globalState));
        let _862_v1;
        _862_v1 = _module.D10.create_DC28(p3);
        let _863_v2;
        _863_v2 = _dafny.Map.Empty.slice().updateUnsafe(true,_862_v1);
        let _864_v3;
        _864_v3 = _dafny.Seq.of(_863_v2, _863_v2);
        let _865_v4;
        let _nw155 = Array((new BigNumber(19)).toNumber());
        _nw155[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,_module.D10.create_DC28(new BigNumber((_861_v0).length)));
        _nw155[(_dafny.ONE).toNumber()] = _863_v2;
        _nw155[(new BigNumber(2)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(3)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(4)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(5)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(6)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(7)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_862_v1);
        _nw155[(new BigNumber(9)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(10)).toNumber()] = (_864_v3)[_module.__default.safeIndex(new BigNumber(837), new BigNumber((_864_v3).length))];
        _nw155[(new BigNumber(11)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(12)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_862_v1);
        _nw155[(new BigNumber(14)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(15)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(16)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(17)).toNumber()] = _863_v2;
        _nw155[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_862_v1);
        _865_v4 = _nw155;
        let _866_v5;
        _866_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f20,p1)).length),_865_v4);
        _866_v5 = (_866_v5).update((_dafny.ZERO).minus(p2), _865_v4);
        let _867_v6;
        let _nw156 = new _module.C1();
        _nw156.__ctor((_this).f20);
        _867_v6 = _nw156;
        (globalState).f6 = (_this).f20;
        let _868_v7;
        _868_v7 = _dafny.MultiSet.fromElements(p3);
        let _869_v8;
        _869_v8 = _dafny.Seq.of((((_868_v7).contains(new BigNumber(857))) ? ((_868_v7).get(new BigNumber(857))) : (_this.f21)));
        (globalState).f6 = (_869_v8)[_module.__default.safeIndex(_this.f21, new BigNumber((_869_v8).length))];
      } else {
        let _870_v9;
        _870_v9 = _dafny.Seq.UnicodeFromString("vkn");
        (globalState).f6 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm21(globalState), _870_v9)).length);
        let _871_v10;
        _871_v10 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f21);
        _871_v10 = (_871_v10).update(_module.__default.fm2(globalState), (p3).minus(p2));
        let _872_v11;
        _872_v11 = _module.D3.create_DC7(new BigNumber(12));
        let _873_v12;
        _873_v12 = new _dafny.CodePoint('y'.codePointAt(0));
        let _874_v13;
        let _nw157 = Array((new BigNumber(22)).toNumber());
        _nw157[(_dafny.ZERO).toNumber()] = _870_v9;
        _nw157[(_dafny.ONE).toNumber()] = ((true) ? (_dafny.Seq.update(_dafny.Seq.update(_870_v9, _module.__default.safeIndex((_872_v11).dtor_cf7, new BigNumber((_870_v9).length)), _873_v12), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.update(_870_v9, _module.__default.safeIndex((_872_v11).dtor_cf7, new BigNumber((_870_v9).length)), _873_v12)).length)), _873_v12)) : (_870_v9));
        _nw157[(new BigNumber(2)).toNumber()] = _870_v9;
        _nw157[(new BigNumber(3)).toNumber()] = _870_v9;
        _nw157[(new BigNumber(4)).toNumber()] = _870_v9;
        _nw157[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_870_v9, _870_v9);
        _nw157[(new BigNumber(6)).toNumber()] = _870_v9;
        _nw157[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("a"), _870_v9);
        _nw157[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("hnastust");
        _nw157[(new BigNumber(9)).toNumber()] = _870_v9;
        _nw157[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_870_v9, _870_v9);
        _nw157[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nsv"), _870_v9);
        _nw157[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm21(globalState), _870_v9);
        _nw157[(new BigNumber(13)).toNumber()] = ((p1) ? (_870_v9) : (_dafny.Seq.UnicodeFromString("jtgqkrddi")));
        _nw157[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("nflhvqi");
        _nw157[(new BigNumber(15)).toNumber()] = _870_v9;
        _nw157[(new BigNumber(16)).toNumber()] = _870_v9;
        _nw157[(new BigNumber(17)).toNumber()] = _870_v9;
        _nw157[(new BigNumber(18)).toNumber()] = _870_v9;
        _nw157[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_870_v9, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(958)), ((_875_v12) => function (_876_i0) {
          return _875_v12;
        })(_873_v12)), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(958)), ((_877_v12) => function (_878_i0) {
          return _877_v12;
        })(_873_v12))).length)), _873_v12));
        _nw157[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_870_v9, _870_v9), _module.__default.safeIndex((_this).f20, new BigNumber((_dafny.Seq.Concat(_870_v9, _870_v9)).length)), _873_v12);
        _nw157[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_870_v9, _module.__default.safeIndex(p3, new BigNumber((_870_v9).length)), new _dafny.CodePoint('t'.codePointAt(0)));
        _874_v13 = _nw157;
        _874_v13 = _874_v13;
        let _879_v14;
        _879_v14 = _dafny.Seq.of(true);
        let _880_v15;
        _880_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p2, new BigNumber((_879_v14).length), (_this).fm15(globalState), new BigNumber(191), new BigNumber(-819)),!(_module.__default.fm2(globalState)));
        let _881_v16;
        _881_v16 = _dafny.Seq.of(p1, (((_880_v15).contains(_dafny.Seq.of(p3))) ? ((_880_v15).get(_dafny.Seq.of(p3))) : (p1)), !(p1));
        (globalState).f1 = (_879_v14)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_879_v14).length))];
        (_this).f21 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber(-411))).plus(p2));
      }
      let _882_v18;
      let _nw158 = new _module.C1();
      _nw158.__ctor(new BigNumber((function () {
        let _coll38 = new _dafny.Set();
        for (const _compr_38 of _dafny.IntegerRange(new BigNumber(-8), new BigNumber(-685))) {
          let _883_v17 = _compr_38;
          if (((new BigNumber(-8)).isLessThanOrEqualTo(_883_v17)) && ((_883_v17).isLessThan(new BigNumber(-685)))) {
            _coll38.add((_883_v17).plus((_this).f20));
          }
        }
        return _coll38;
      }()).length));
      _882_v18 = _nw158;
      (globalState).f1 = p1;
      let _884_v19;
      _884_v19 = _dafny.Set.fromElements(p1);
      let _885_v20;
      _885_v20 = _dafny.Map.Empty.slice().updateUnsafe((_882_v18).f24,(_884_v19).Intersect(_884_v19));
      _885_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p3, (_this).f20),(_884_v19).Intersect(_884_v19));
      let _886_v21;
      _886_v21 = _dafny.Seq.of(p1, p1, false, p1);
      let _887_v22;
      let _nw159 = Array((new BigNumber(25)).toNumber()).fill(false);
      _887_v22 = _nw159;
      let _888_v23;
      _888_v23 = _dafny.MultiSet.fromElements(_887_v22);
      if ((_886_v21)[_module.__default.safeIndex(new BigNumber((_888_v23).cardinality()), new BigNumber((_886_v21).length))]) {
        let _889_v24;
        _889_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,!(p1) || (p1));
        let _890_v25;
        _890_v25 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
        let _891_v26;
        _891_v26 = _module.D7.create_DC16(p1, _this.f21, p1);
        let _892_v27;
        _892_v27 = _dafny.Seq.UnicodeFromString("pcmsmlo");
        let _893_v28;
        _893_v28 = _dafny.Map.Empty.slice().updateUnsafe(_891_v26,_892_v27);
        let _894_v29;
        _894_v29 = _module.D13.create_DC36(_893_v28);
        _889_v24 = (_889_v24).update((((_890_v25).contains((_dafny.ZERO).minus(new BigNumber(((_894_v29).dtor_cf64).length)))) ? ((_890_v25).get((_dafny.ZERO).minus(new BigNumber(((_894_v29).dtor_cf64).length)))) : (p1)), p1);
        let _index155 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length));
        (_887_v22)[_index155] = p1;
        let _index156 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length));
        (_887_v22)[_index156] = p1;
        let _895_v30;
        _895_v30 = _dafny.Seq.of(_this.f21, (_this).f20, p2);
        let _896_v31;
        _896_v31 = _dafny.Seq.of(_889_v24);
        let _897_v32;
        _897_v32 = _dafny.Map.Empty.slice().updateUnsafe((_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))],(_882_v18).f24);
        let _898_v33;
        _898_v33 = _895_v30;
        let _899_v34;
        _899_v34 = _dafny.Map.Empty.slice().updateUnsafe((_896_v31)[_module.__default.safeIndex((((_897_v32).contains((_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))])) ? ((_897_v32).get((_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))])) : (_this.f21)), new BigNumber((_896_v31).length))],(_898_v33));
        let _900_v35;
        _900_v35 = new _dafny.CodePoint('x'.codePointAt(0));
        let _901_v36;
        let _nw160 = Array((new BigNumber(14)).toNumber());
        _nw160[(_dafny.ZERO).toNumber()] = (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))];
        _nw160[(_dafny.ONE).toNumber()] = (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))];
        _nw160[(new BigNumber(2)).toNumber()] = (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))];
        _nw160[(new BigNumber(3)).toNumber()] = (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))];
        _nw160[(new BigNumber(4)).toNumber()] = (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))];
        _nw160[(new BigNumber(5)).toNumber()] = p1;
        _nw160[(new BigNumber(6)).toNumber()] = (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))];
        _nw160[(new BigNumber(7)).toNumber()] = (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))];
        _nw160[(new BigNumber(8)).toNumber()] = (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))];
        _nw160[(new BigNumber(9)).toNumber()] = p1;
        _nw160[(new BigNumber(10)).toNumber()] = p1;
        _nw160[(new BigNumber(11)).toNumber()] = (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))];
        _nw160[(new BigNumber(12)).toNumber()] = p1;
        _nw160[(new BigNumber(13)).toNumber()] = p1;
        _901_v36 = _nw160;
        let _902_v37;
        _902_v37 = _dafny.Set.fromElements(_901_v36, _901_v36);
        let _903_v38;
        _903_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_902_v37).length),(_882_v18).f24);
        _895_v30 = (((_899_v34).contains((_dafny.Map.Empty.slice().updateUnsafe(false,!((_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))]))).Merge(_module.__default.fm34(p2, p1, _module.__default.fm1((_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))], globalState), _900_v35, globalState)))) ? ((_899_v34).get((_dafny.Map.Empty.slice().updateUnsafe(false,!((_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))]))).Merge(_module.__default.fm34(p2, p1, _module.__default.fm1((_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))], globalState), _900_v35, globalState)))) : (_dafny.Seq.of(new BigNumber((_889_v24).length), (_dafny.ZERO).minus(p2), new BigNumber((_903_v38).length), new BigNumber((_884_v19).length))));
        let _904_v39;
        let _nw161 = Array((new BigNumber(28)).toNumber()).fill(_module.D7.Default());
        _904_v39 = _nw161;
        let _905_v40;
        _905_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm29((_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))], _this.f21, p2, (_887_v22)[_module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length))], globalState),_901_v36);
        let _906_v41;
        _906_v41 = _dafny.Map.Empty.slice().updateUnsafe(_905_v40,_904_v39);
        _904_v39 = (((_906_v41).contains(_905_v40)) ? ((_906_v41).get(_905_v40)) : (_904_v39));
        let _index157 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_887_v22).length));
        (_887_v22)[_index157] = p1;
      } else {
        (globalState).f1 = p1;
        let _907_v42;
        _907_v42 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
        let _908_v43;
        _908_v43 = _dafny.Map.Empty.slice().updateUnsafe(_907_v42,p1);
        if (((p1) ? (p1) : ((_882_v18).fm24(p1, _908_v43, globalState)))) {
          let _909_v44;
          let _nw162 = Array((new BigNumber(13)).toNumber()).fill(_module.D3.Default());
          _909_v44 = _nw162;
          let _910_v45;
          _910_v45 = _dafny.Set.fromElements(_909_v44);
          let _911_v46;
          _911_v46 = _dafny.Seq.UnicodeFromString("q");
          let _912_v47;
          _912_v47 = _dafny.MultiSet.fromElements(p1);
          let _913_v48;
          _913_v48 = _dafny.Map.Empty.slice().updateUnsafe(p1,_912_v47);
          let _914_v49;
          _914_v49 = new _dafny.CodePoint('j'.codePointAt(0));
          let _915_v50;
          let _init23 = ((_916_v18) => function (_917_i1) {
            return _module.__default.safeDivisionInt(_917_i1, (_916_v18).f24);
          })(_882_v18);
          let _nw163 = Array((new BigNumber(23)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw163.length); _i0_23++) {
            _nw163[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _915_v50 = _nw163;
          let _rhs109 = p1;
          let _rhs110 = (_module.D9.create_DC23(_910_v45, _dafny.Seq.update(_911_v46, _module.__default.safeIndex(new BigNumber(((((_913_v48).contains(p1)) ? ((_913_v48).get(p1)) : (_912_v47))).cardinality()), new BigNumber((_911_v46).length)), _914_v49), p1, _915_v50, _914_v49)).dtor_cf40;
          let _rhs111 = (_882_v18).f24;
          let _lhs97 = globalState;
          let _lhs98 = globalState;
          r0 = _rhs109;
          _lhs97.f9 = _rhs110;
          _lhs98.f6 = _rhs111;
          let _918_v51;
          _918_v51 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_882_v18).f24), (_882_v18).f24);
          let _919_v52;
          _919_v52 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus((_this).f20));
          let _920_v53;
          _920_v53 = _dafny.Seq.of(_this.f21);
          let _921_v54;
          _921_v54 = _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of((_882_v18).f24, _this.f21)).length), (_882_v18).f24), _918_v51, _918_v51);
          let _922_v55;
          _922_v55 = _dafny.Map.Empty.slice().updateUnsafe(_914_v49,p2);
          let _923_v56;
          let _nw164 = Array((new BigNumber(23)).toNumber());
          _nw164[(_dafny.ZERO).toNumber()] = (_918_v51).update(new BigNumber((_919_v52).length), _module.__default.abs((_module.D0.create_DC1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(738)), ((_924_v49) => function (_925_i2) {
  return _924_v49;
})(_914_v49))).length))).dtor_cf1));
          _nw164[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements((_882_v18).f24, (_this).f20)).update((_882_v18).f24, _module.__default.abs((_882_v18).f24));
          _nw164[(new BigNumber(2)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(3)).toNumber()] = (_918_v51).Intersect(_dafny.MultiSet.FromArray(_920_v53));
          _nw164[(new BigNumber(4)).toNumber()] = (_918_v51).Difference(_dafny.MultiSet.fromElements(new BigNumber(69)));
          _nw164[(new BigNumber(5)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(6)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_this.f21));
          _nw164[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements((_882_v18).f24, new BigNumber((_911_v46).length));
          _nw164[(new BigNumber(9)).toNumber()] = ((_921_v54)[_module.__default.safeIndex((_882_v18).f24, new BigNumber((_921_v54).length))]).Difference(_918_v51);
          _nw164[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber(609), _this.f21, p3);
          _nw164[(new BigNumber(11)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(12)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(13)).toNumber()] = (_918_v51).Difference(_918_v51);
          _nw164[(new BigNumber(14)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(15)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements((_882_v18).f24, p3);
          _nw164[(new BigNumber(17)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(18)).toNumber()] = ((_918_v51).update(new BigNumber((_922_v55).length), _module.__default.abs(p2))).Union(_918_v51);
          _nw164[(new BigNumber(19)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(20)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(21)).toNumber()] = _918_v51;
          _nw164[(new BigNumber(22)).toNumber()] = _918_v51;
          _923_v56 = _nw164;
          let _index158 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_923_v56).length));
          (_923_v56)[_index158] = _dafny.MultiSet.fromElements((_this).f20, (_882_v18).f24, new BigNumber((_884_v19).length), (_882_v18).f24, new BigNumber((_911_v46).length));
          let _926_v57;
          _926_v57 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,p1), _907_v42);
          let _index159 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_887_v22).length));
          (_887_v22)[_index159] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(450)), ((_927_v42) => function (_928_i3) {
            return _927_v42;
          })(_907_v42)), _926_v57);
          let _929_v58;
          let _init24 = ((_930_v49) => function (_931_i4) {
            return _module.D1.create_DC4(_930_v49);
          })(_914_v49);
          let _nw165 = Array((new BigNumber(9)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw165.length); _i0_24++) {
            _nw165[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _929_v58 = _nw165;
          let _932_v59;
          let _init25 = ((_933_v21) => function (_934_i5) {
            return _933_v21;
          })(_886_v21);
          let _nw166 = Array((new BigNumber(12)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw166.length); _i0_25++) {
            _nw166[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _932_v59 = _nw166;
          let _935_v60;
          _935_v60 = _dafny.Map.Empty.slice().updateUnsafe(_932_v59,_929_v58);
          let _index160 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_923_v56).length));
          let _index161 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_887_v22).length));
          let _rhs112 = (p3).minus((_882_v18).f24);
          let _rhs113 = _918_v51;
          let _rhs114 = p1;
          let _rhs115 = (((_935_v60).contains(_932_v59)) ? ((_935_v60).get(_932_v59)) : (_929_v58));
          let _lhs99 = _this;
          let _lhs100 = _923_v56;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_923_v56).length));
          let _lhs102 = _887_v22;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_887_v22).length));
          _lhs99.f21 = _rhs112;
          _lhs100[_lhs101] = _rhs113;
          _lhs102[_lhs103] = _rhs114;
          _929_v58 = _rhs115;
          (globalState).f6 = new BigNumber((_module.__default.fm28(new BigNumber((_884_v19).length), _module.__default.fm2(globalState), ((_882_v18).f24).isLessThan(p2), globalState)).length);
          r0 = (_887_v22)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_887_v22).length))];
          (globalState).f1 = p1;
        } else {
          let _936_v61;
          _936_v61 = _dafny.Seq.UnicodeFromString("ydlpbiue");
          let _937_v62;
          _937_v62 = new _dafny.CodePoint('i'.codePointAt(0));
          let _938_v63;
          let _nw167 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _938_v63 = _nw167;
          let _939_v64;
          _939_v64 = _dafny.Seq.of(_module.D3.create_DC6(_938_v63));
          let _940_v65;
          let _nw168 = new _module.C0();
          _nw168.__ctor(_939_v64, p1);
          _940_v65 = _nw168;
          let _941_v66;
          _941_v66 = _dafny.Map.Empty.slice().updateUnsafe(p2,_887_v22);
          let _942_v67;
          _942_v67 = _dafny.Seq.of(_887_v22);
          let _943_v68;
          _943_v68 = _dafny.Seq.of(p2);
          let _944_v69;
          let _nw169 = Array((new BigNumber(27)).toNumber());
          _nw169[(_dafny.ZERO).toNumber()] = _887_v22;
          _nw169[(_dafny.ONE).toNumber()] = _887_v22;
          _nw169[(new BigNumber(2)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(3)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(4)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(5)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(6)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(7)).toNumber()] = (((_941_v66).contains((_882_v18).f24)) ? ((_941_v66).get((_882_v18).f24)) : (_887_v22));
          _nw169[(new BigNumber(8)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(9)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(10)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(11)).toNumber()] = (_942_v67)[_module.__default.safeIndex(new BigNumber((_943_v68).length), new BigNumber((_942_v67).length))];
          _nw169[(new BigNumber(12)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(13)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(14)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(15)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(16)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(17)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(18)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(19)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(20)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(21)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(22)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(23)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(24)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(25)).toNumber()] = _887_v22;
          _nw169[(new BigNumber(26)).toNumber()] = _887_v22;
          _944_v69 = _nw169;
          let _945_v70;
          _945_v70 = _module.D7.create_DC17(p1, _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(940)), ((_946_v21) => function (_947_i6) {
  return new BigNumber((_946_v21).length);
})(_886_v21)), new BigNumber((_dafny.Seq.update(_936_v61, _module.__default.safeIndex((_882_v18).f24, new BigNumber((_936_v61).length)), _937_v62)).length)), _940_v65, _module.D5.create_DC10(_944_v69), !(p1));
          _945_v70 = _945_v70;
          (globalState).f6 = p2;
          let _948_v71;
          let _init26 = ((_949_v68) => function (_950_i7) {
            return _dafny.Seq.Concat(_949_v68, _949_v68);
          })(_943_v68);
          let _nw170 = Array((new BigNumber(10)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw170.length); _i0_26++) {
            _nw170[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _948_v71 = _nw170;
          let _951_v72;
          _951_v72 = _dafny.Seq.of(_948_v71, _948_v71, (_module.D14.create_DC38(_948_v71)).dtor_cf67);
          _948_v71 = (_951_v72)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_951_v72).length))];
          let _index162 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_887_v22).length));
          (_887_v22)[_index162] = (_940_v65).f23;
          let _index163 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_887_v22).length));
          (_887_v22)[_index163] = (_940_v65).f23;
          let _index164 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_887_v22).length));
          let _index165 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_887_v22).length));
          let _rhs116 = (_940_v65).f23;
          let _rhs117 = (_882_v18).f24;
          let _rhs118 = (_940_v65).f23;
          let _rhs119 = (_940_v65).f23;
          let _lhs104 = _887_v22;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_887_v22).length));
          let _lhs106 = globalState;
          let _lhs107 = _887_v22;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_887_v22).length));
          let _lhs109 = globalState;
          _lhs104[_lhs105] = _rhs116;
          _lhs106.f6 = _rhs117;
          _lhs107[_lhs108] = _rhs118;
          _lhs109.f1 = _rhs119;
          let _952_v73;
          let _nw171 = new _module.C1();
          _nw171.__ctor(_this.f21);
          _952_v73 = _nw171;
        }
        let _953_v74;
        _953_v74 = _dafny.Seq.UnicodeFromString("lsek");
        (globalState).f6 = _module.__default.safeModuloInt(p2, new BigNumber((_953_v74).length));
        let _954_v75;
        _954_v75 = new _dafny.CodePoint('a'.codePointAt(0));
        if (!_dafny.Seq.contains(_953_v74, _954_v75)) {
          (_this).f21 = (_this).f20;
          (globalState).f6 = _module.__default.fm1(((p1) ? (p1) : (p1)), globalState);
          let _955_v76;
          _955_v76 = _dafny.MultiSet.fromElements((_882_v18).f24, (_882_v18).f24, _this.f21);
          let _956_v77;
          _956_v77 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length));
          let _957_v78;
          _957_v78 = _dafny.Seq.of((_956_v77)[_module.__default.safeIndex(_this.f21, new BigNumber((_956_v77).length))], (_882_v18).f24, (_this).f20);
          _955_v76 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_956_v77, _957_v78));
          (globalState).f1 = (new BigNumber(165)).isLessThan((_882_v18).f24);
          let _958_v79;
          let _nw172 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _958_v79 = _nw172;
          let _index166 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_958_v79).length));
          (_958_v79)[_index166] = new BigNumber(-963);
          let _index167 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_958_v79).length));
          (_958_v79)[_index167] = (_dafny.ZERO).minus((_this).fm15(globalState));
        } else {
          let _959_v80;
          _959_v80 = _dafny.Map.Empty.slice().updateUnsafe(!(!(p1) || (p1)),_953_v74);
          _959_v80 = (_959_v80).update(p1, _953_v74);
          _884_v19 = (_884_v19).Union(_884_v19);
          (globalState).f6 = _module.__default.fm1(!(p1), globalState);
          (globalState).f6 = ((_882_v18).f24).multipliedBy(p3);
          (globalState).f6 = ((p1) ? (((_this).f20).minus(_this.f21)) : (new BigNumber(-212)));
        }
        if (p1) {
          let _960_v81;
          _960_v81 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(globalState),_module.__default.safeDivisionInt(p2, p2));
          let _961_v82;
          _961_v82 = _dafny.Seq.of(p3);
          let _962_v83;
          _962_v83 = _dafny.Map.Empty.slice().updateUnsafe(_961_v82,(_dafny.Map.Empty.slice().updateUnsafe(_886_v21,(_882_v18).f24)).Merge(_960_v81));
          _960_v81 = (((_962_v83).contains(_dafny.Seq.Concat(_961_v82, _961_v82))) ? ((_962_v83).get(_dafny.Seq.Concat(_961_v82, _961_v82))) : (_960_v81));
          let _963_v84;
          _963_v84 = _dafny.MultiSet.fromElements((_this).f20, (_882_v18).f24);
          let _964_v85;
          _964_v85 = _dafny.Set.fromElements(_963_v84);
          let _965_v86;
          _965_v86 = _dafny.Map.Empty.slice().updateUnsafe(_964_v85,_module.__default.safeDivisionInt(p2, (_882_v18).f24));
          _965_v86 = (_965_v86).update(_964_v85, (_this).f20);
          (globalState).f6 = new BigNumber(-356);
          let _index168 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_887_v22).length));
          (_887_v22)[_index168] = p1;
          let _index169 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_887_v22).length));
          let _rhs120 = ((p1) ? (_dafny.Seq.IsPrefixOf(_953_v74, _953_v74)) : (p1));
          let _rhs121 = ((p1) ? ((_882_v18).f24) : (new BigNumber(421)));
          let _lhs110 = _887_v22;
          let _lhs111 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_887_v22).length));
          let _lhs112 = _this;
          _lhs110[_lhs111] = _rhs120;
          _lhs112.f21 = _rhs121;
          let _966_v87;
          _966_v87 = _dafny.Map.Empty.slice().updateUnsafe(_954_v75,_953_v74);
          r0 = (_882_v18).fm24(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("eisnpxsp"), (((_966_v87).contains(_954_v75)) ? ((_966_v87).get(_954_v75)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(268)), ((_967_v75) => function (_968_i8) {
            return _967_v75;
          })(_954_v75))))), _908_v43, globalState);
        } else {
          (globalState).f6 = ((_882_v18).f24).minus((_882_v18).f24);
          let _969_v88;
          _969_v88 = _dafny.Set.fromElements(_953_v74);
          let _rhs122 = (_882_v18).f24;
          let _rhs123 = _module.__default.fm2(globalState);
          let _rhs124 = (_969_v88).IsProperSubsetOf(_969_v88);
          let _rhs125 = _887_v22;
          let _lhs113 = globalState;
          let _lhs114 = globalState;
          let _lhs115 = globalState;
          _lhs113.f6 = _rhs122;
          _lhs114.f1 = _rhs123;
          _lhs115.f1 = _rhs124;
          _887_v22 = _rhs125;
          let _970_v89;
          _970_v89 = _module.D5.create_DC12(new BigNumber((_953_v74).length), p1, _954_v75);
          let _pat_let_tv25 = _954_v75;
          let _pat_let_tv26 = _882_v18;
          let _971_v90;
          _971_v90 = _dafny.MultiSet.fromElements(_module.D5.create_DC12((_882_v18).f24, p1, _954_v75), function (_pat_let46_0) {
            return function (_972_dt__update__tmp_h0) {
              return function (_pat_let47_0) {
                return function (_973_dt__update_hcf17_h0) {
                  return function (_pat_let48_0) {
                    return function (_974_dt__update_hcf15_h0) {
                      return _module.D5.create_DC12(_974_dt__update_hcf15_h0, (_972_dt__update__tmp_h0).dtor_cf16, _973_dt__update_hcf17_h0);
                    }(_pat_let48_0);
                  }((_pat_let_tv26).f24);
                }(_pat_let47_0);
              }(_pat_let_tv25);
            }(_pat_let46_0);
          }(_970_v89), _970_v89);
          let _975_v91;
          _975_v91 = _dafny.Seq.of(p0, p0);
          let _976_v92;
          let _nw173 = new _module.C0();
          _nw173.__ctor(_975_v91, true);
          _976_v92 = _nw173;
          let _977_v93;
          _977_v93 = _dafny.Seq.of(_976_v92, _976_v92);
          let _978_v94;
          _978_v94 = _dafny.Set.fromElements((_953_v74)[_module.__default.safeIndex(new BigNumber((_977_v93).length), new BigNumber((_953_v74).length))], _954_v75);
          let _979_v95;
          let _nw174 = Array((new BigNumber(25)).toNumber());
          _nw174[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(676),p1);
          _nw174[(_dafny.ONE).toNumber()] = _907_v42;
          _nw174[(new BigNumber(2)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(3)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-939),p1);
          _nw174[(new BigNumber(5)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(6)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(7)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(8)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(9)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(10)).toNumber()] = _module.__default.fm0(p1, p1, _this.f21, globalState);
          _nw174[(new BigNumber(11)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(12)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(13)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-542)),true);
          _nw174[(new BigNumber(15)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(16)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(17)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(18)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(19)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(20)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(21)).toNumber()] = (_907_v42).update(new BigNumber((_978_v94).length), p1);
          _nw174[(new BigNumber(22)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(23)).toNumber()] = _907_v42;
          _nw174[(new BigNumber(24)).toNumber()] = _907_v42;
          _979_v95 = _nw174;
          let _980_v96;
          _980_v96 = _module.D9.create_DC22(_979_v95);
          let _981_v97;
          let _nw175 = Array((new BigNumber(9)).toNumber());
          _nw175[(_dafny.ZERO).toNumber()] = new BigNumber((_953_v74).length);
          _nw175[(_dafny.ONE).toNumber()] = new BigNumber((_953_v74).length);
          _nw175[(new BigNumber(2)).toNumber()] = new BigNumber(-550);
          _nw175[(new BigNumber(3)).toNumber()] = (_882_v18).f24;
          _nw175[(new BigNumber(4)).toNumber()] = p2;
          _nw175[(new BigNumber(5)).toNumber()] = (_882_v18).f24;
          _nw175[(new BigNumber(6)).toNumber()] = (_this).f20;
          _nw175[(new BigNumber(7)).toNumber()] = p3;
          _nw175[(new BigNumber(8)).toNumber()] = (_882_v18).f24;
          _981_v97 = _nw175;
          let _982_v98;
          let _nw176 = Array((new BigNumber(10)).toNumber());
          _nw176[(_dafny.ZERO).toNumber()] = p0;
          _nw176[(_dafny.ONE).toNumber()] = p0;
          _nw176[(new BigNumber(2)).toNumber()] = p0;
          _nw176[(new BigNumber(3)).toNumber()] = p0;
          _nw176[(new BigNumber(4)).toNumber()] = p0;
          _nw176[(new BigNumber(5)).toNumber()] = p0;
          _nw176[(new BigNumber(6)).toNumber()] = _module.D3.create_DC6(_981_v97);
          _nw176[(new BigNumber(7)).toNumber()] = _module.D3.create_DC6(_981_v97);
          _nw176[(new BigNumber(8)).toNumber()] = p0;
          _nw176[(new BigNumber(9)).toNumber()] = p0;
          _982_v98 = _nw176;
          let _983_v99;
          _983_v99 = _dafny.Set.fromElements(_982_v98);
          let _984_v100;
          _984_v100 = _module.D9.create_DC25(_module.D9.create_DC23(_983_v99, _dafny.Seq.UnicodeFromString("dkxf"), (_976_v92).f23, _981_v97, _954_v75));
          let _985_v101;
          _985_v101 = _dafny.Map.Empty.slice().updateUnsafe(p1,_886_v21);
          let _986_v103;
          let _nw177 = Array((new BigNumber(15)).toNumber());
          _nw177[(_dafny.ZERO).toNumber()] = (_this.f21).minus(new BigNumber((_971_v90).cardinality()));
          _nw177[(_dafny.ONE).toNumber()] = ((p1) ? (p3) : (new BigNumber((_dafny.Seq.UnicodeFromString("hty")).length)));
          _nw177[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_module.D9.create_DC25(_980_v96), _984_v100)).length), p3));
          _nw177[(new BigNumber(3)).toNumber()] = (_882_v18).f24;
          _nw177[(new BigNumber(4)).toNumber()] = new BigNumber(((((_985_v101).contains((_976_v92).f23)) ? ((_985_v101).get((_976_v92).f23)) : (_module.__default.fm19(globalState)))).length);
          _nw177[(new BigNumber(5)).toNumber()] = (new BigNumber((_886_v21).length)).minus(_this.f21);
          _nw177[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_886_v21).length), (_882_v18).f24);
          _nw177[(new BigNumber(7)).toNumber()] = (_882_v18).f24;
          _nw177[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_886_v21).length));
          _nw177[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt((_this).f20, (_this).f20);
          _nw177[(new BigNumber(10)).toNumber()] = p3;
          _nw177[(new BigNumber(11)).toNumber()] = new BigNumber((function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of (_dafny.Map.Empty.slice().updateUnsafe(_this.f21,new BigNumber((_886_v21).length))).Keys.Elements) {
              let _987_v102 = _compr_39;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_this.f21,new BigNumber((_886_v21).length))).contains(_987_v102)) {
                _coll39.push([(_987_v102).multipliedBy((_882_v18).f24),_953_v74]);
              }
            }
            return _coll39;
          }()).length);
          _nw177[(new BigNumber(12)).toNumber()] = (_this).f20;
          _nw177[(new BigNumber(13)).toNumber()] = (_882_v18).f24;
          _nw177[(new BigNumber(14)).toNumber()] = p2;
          _986_v103 = _nw177;
          let _index170 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_986_v103).length));
          (_986_v103)[_index170] = _this.f21;
          let _988_v104;
          _988_v104 = _dafny.Seq.of(_module.__default.fm1(false, globalState), p3);
          let _index171 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_986_v103).length));
          (_986_v103)[_index171] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((new BigNumber(((_module.D3.create_DC8(_988_v104, _953_v74, p1, _954_v75)).dtor_cf9).length)).plus(new BigNumber(391)), _module.__default.fm1((_976_v92).fm17(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_954_v75)).length), new BigNumber(24), (_976_v92).f23, globalState), globalState)));
          let _989_v105;
          let _nw178 = new _module.C0();
          _nw178.__ctor(_dafny.Seq.Concat(_975_v91, (_976_v92).f22), true);
          _989_v105 = _nw178;
          (_882_v18).m2((_989_v105).f23, globalState);
        }
      }
      let _hi7 = (_dafny.ZERO).minus((_882_v18).f24);
      for (let _990_i9 = _module.__default.fm1(p1, globalState); _990_i9.isLessThan(_hi7); _990_i9 = _990_i9.plus(_dafny.ONE)) {
        (globalState).f6 = (_882_v18).f24;
        (globalState).f6 = (_882_v18).f24;
        let _991_v106;
        _991_v106 = _dafny.Set.fromElements(_this.f21);
        let _992_v107;
        _992_v107 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_991_v106).Union(_991_v106));
        _992_v107 = (_992_v107).update(p1, _module.__default.fm28((_882_v18).f24, p1, p1, globalState));
        if (p1) {
          let _index172 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_887_v22).length));
          (_887_v22)[_index172] = p1;
          let _993_v108;
          _993_v108 = _dafny.Seq.UnicodeFromString("kpy");
          let _index173 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_887_v22).length));
          (_887_v22)[_index173] = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ja"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(622)), function (_994_i10) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          })), _993_v108);
          (globalState).f1 = (new BigNumber(-549)).isLessThanOrEqualTo(new BigNumber((_886_v21).length));
          let _995_v109;
          _995_v109 = _dafny.Seq.of(p0, p0);
          let _996_v110;
          let _nw179 = new _module.C0();
          _nw179.__ctor(_995_v109, (_887_v22)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_887_v22).length))]);
          _996_v110 = _nw179;
          let _997_v111;
          _997_v111 = _dafny.MultiSet.fromElements(_996_v110, _996_v110, _996_v110);
          let _998_v112;
          _998_v112 = _dafny.Seq.of(_996_v110);
          let _999_v113;
          _999_v113 = _dafny.Seq.of(_998_v112, _998_v112, _998_v112);
          let _1000_v114;
          let _nw180 = Array((new BigNumber(22)).toNumber());
          _nw180[(_dafny.ZERO).toNumber()] = (_997_v111).update(_996_v110, _module.__default.abs((_882_v18).f24));
          _nw180[(_dafny.ONE).toNumber()] = _997_v111;
          _nw180[(new BigNumber(2)).toNumber()] = (_997_v111).update(_996_v110, _module.__default.abs(p3));
          _nw180[(new BigNumber(3)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(4)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(5)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(6)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(7)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(8)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_998_v112, _998_v112));
          _nw180[(new BigNumber(10)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat((_999_v113)[_module.__default.safeIndex(p2, new BigNumber((_999_v113).length))], _998_v112));
          _nw180[(new BigNumber(12)).toNumber()] = (_997_v111).Union(_997_v111);
          _nw180[(new BigNumber(13)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(14)).toNumber()] = (_dafny.MultiSet.fromElements(_996_v110, _996_v110)).update(_996_v110, _module.__default.abs((_882_v18).f24));
          _nw180[(new BigNumber(15)).toNumber()] = (_997_v111).Intersect(_997_v111);
          _nw180[(new BigNumber(16)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(17)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(18)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(19)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(20)).toNumber()] = _997_v111;
          _nw180[(new BigNumber(21)).toNumber()] = (_dafny.MultiSet.fromElements(_996_v110, _996_v110)).update(_996_v110, _module.__default.abs((_882_v18).f24));
          _1000_v114 = _nw180;
          let _index174 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_1000_v114).length));
          (_1000_v114)[_index174] = _997_v111;
          let _1001_v115;
          _1001_v115 = new _dafny.CodePoint('g'.codePointAt(0));
          let _1002_v116;
          _1002_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1001_v115,(_887_v22)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_887_v22).length))]);
          let _1003_v117;
          _1003_v117 = _dafny.Map.Empty.slice().updateUnsafe(p1,(((_1002_v116).contains(_1001_v115)) ? ((_1002_v116).get(_1001_v115)) : (!((_887_v22)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_887_v22).length))]))));
          let _1004_v118;
          _1004_v118 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1003_v117).length),_dafny.MultiSet.FromArray(_998_v112));
          let _1005_v119;
          let _nw181 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1005_v119 = _nw181;
          let _index175 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_1000_v114).length));
          let _rhs126 = ((((_1004_v118).contains((_882_v18).f24)) ? ((_1004_v118).get((_882_v18).f24)) : (_997_v111))).Difference(_997_v111);
          let _rhs127 = (_887_v22)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_887_v22).length))];
          let _rhs128 = !((_887_v22)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_887_v22).length))]);
          let _rhs129 = _1005_v119;
          let _lhs116 = _1000_v114;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_1000_v114).length));
          let _lhs118 = globalState;
          let _lhs119 = globalState;
          _lhs116[_lhs117] = _rhs126;
          _lhs118.f1 = _rhs127;
          r0 = _rhs128;
          _lhs119.f14 = _rhs129;
          _991_v106 = _module.__default.fm28(_this.f21, (_887_v22)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_887_v22).length))], (_884_v19).IsProperSubsetOf(_884_v19), globalState);
          let _1006_v120;
          _1006_v120 = _dafny.Seq.of(_882_v18, _882_v18);
          _882_v18 = (_1006_v120)[_module.__default.safeIndex(new BigNumber(-25), new BigNumber((_1006_v120).length))];
        } else {
          let _1007_v121;
          let _init27 = function (_1008_i11) {
            return (_1008_i11).minus(new BigNumber(905));
          };
          let _nw182 = Array((new BigNumber(16)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw182.length); _i0_27++) {
            _nw182[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1007_v121 = _nw182;
          let _index176 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1007_v121).length));
          (_1007_v121)[_index176] = new BigNumber((((p1) ? (_884_v19) : (_884_v19))).length);
          let _1009_v122;
          _1009_v122 = _dafny.Seq.UnicodeFromString("kn");
          let _1010_v123;
          _1010_v123 = new _dafny.CodePoint('a'.codePointAt(0));
          let _1011_v124;
          _1011_v124 = _dafny.Map.Empty.slice().updateUnsafe(_1010_v123,_dafny.Map.Empty.slice().updateUnsafe(p1,p1));
          let _1012_v125;
          _1012_v125 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _1013_v126;
          _1013_v126 = _dafny.Map.Empty.slice().updateUnsafe((_882_v18).f24,(((_1011_v124).contains(_1010_v123)) ? ((_1011_v124).get(_1010_v123)) : (_1012_v125)));
          let _index177 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1007_v121).length));
          let _rhs130 = (_882_v18).f24;
          let _rhs131 = (_this.f21).isEqualTo((_dafny.ZERO).minus(new BigNumber((_1013_v126).length)));
          let _rhs132 = _dafny.Seq.update(_1009_v122, _module.__default.safeIndex(_990_i9, new BigNumber((_1009_v122).length)), new _dafny.CodePoint('r'.codePointAt(0)));
          let _rhs133 = p1;
          let _lhs120 = _1007_v121;
          let _lhs121 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1007_v121).length));
          let _lhs122 = globalState;
          let _lhs123 = globalState;
          _lhs120[_lhs121] = _rhs130;
          _lhs122.f1 = _rhs131;
          _1009_v122 = _rhs132;
          _lhs123.f1 = _rhs133;
          let _rhs134 = (_this).fm15(globalState);
          let _rhs135 = p1;
          let _lhs124 = _this;
          _lhs124.f21 = _rhs134;
          r0 = _rhs135;
          let _1014_v127;
          let _nw183 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _1014_v127 = _nw183;
          let _1015_v128;
          _1015_v128 = _dafny.MultiSet.fromElements(_991_v106);
          let _1016_v129;
          _1016_v129 = _dafny.Seq.of(_1009_v122);
          let _1017_v130;
          _1017_v130 = _dafny.Map.Empty.slice().updateUnsafe(_1015_v128,(_1016_v129)[_module.__default.safeIndex((_882_v18).f24, new BigNumber((_1016_v129).length))]);
          let _rhs136 = ((p1) ? (_1014_v127) : (_1014_v127));
          let _rhs137 = !(!(p1));
          let _rhs138 = new BigNumber(((((_1017_v130).contains(_1015_v128)) ? ((_1017_v130).get(_1015_v128)) : (_1009_v122))).length);
          let _rhs139 = new _dafny.CodePoint('v'.codePointAt(0));
          let _lhs125 = globalState;
          let _lhs126 = globalState;
          r2 = _rhs136;
          _lhs125.f1 = _rhs137;
          _lhs126.f6 = _rhs138;
          _1010_v123 = _rhs139;
          let _index178 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_887_v22).length));
          (_887_v22)[_index178] = p1;
          let _index179 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_887_v22).length));
          (_887_v22)[_index179] = p1;
          (globalState).f1 = (_887_v22)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_887_v22).length))];
        }
      }
      r0 = !(!(p1));
      let _1018_v131;
      _1018_v131 = _dafny.Seq.UnicodeFromString("hwtwcswh");
      let _1019_v132;
      _1019_v132 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f20).multipliedBy(p2),((_dafny.ZERO).minus(_this.f21)).minus(new BigNumber((_1018_v131).length)));
      r1 = _1019_v132;
      let _1020_v133;
      let _nw184 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
      _1020_v133 = _nw184;
      r2 = _1020_v133;
      return [r0, r1, r2];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f17 = _dafny.Map.Empty;
      this._f18 = false;
      this._f19 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
    set f17(value) {
      let _this = this;
      _this._f17 = value;
    };
    __ctor(f18, f19, f17) {
      let _this = this;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f17 = f17;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.D0.create_DC2(_module.D0.create_DC0((_this).f18));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      (globalState).f1 = !(((_dafny.ZERO).minus((_this).f19)).isLessThanOrEqualTo((_this).f19));
      let _1021_v0;
      let _nw185 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
      _1021_v0 = _nw185;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1021_v0).length))) {
        let _1022_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1022_i0)) && ((_1022_i0).isLessThan(new BigNumber((_1021_v0).length))))) {
          (_1021_v0)[(_1022_i0)] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), function (_1023_i1) {
            return (_this).f19;
          }), _dafny.Seq.of((_this).f19));
        }
      }
      let _1024_i2;
      _1024_i2 = _dafny.ZERO;
      L6: {
        while ((_this).f18) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1024_i2)) {
              break L6;
            }
            _1024_i2 = (_1024_i2).plus(_dafny.ONE);
            let _1025_v1;
            _1025_v1 = _dafny.Set.fromElements((_this).f18, false);
            if (!(!((_this).f18)) || ((_dafny.Set.fromElements((_this).f18)).IsProperSubsetOf(_1025_v1))) {
              (globalState).f1 = (((_this).f18) ? (!((_this).f18) || (true)) : ((_this).f18));
              let _1026_v2;
              _1026_v2 = _dafny.Seq.of((_this).f18);
              let _1027_v3;
              let _nw186 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
              _1027_v3 = _nw186;
              let _1028_v4;
              _1028_v4 = _module.D3.create_DC6(_1027_v3);
              let _1029_v5;
              _1029_v5 = _dafny.Seq.of(_1028_v4);
              let _1030_v6;
              _1030_v6 = _dafny.MultiSet.fromElements((_this).f18, false, (_this).f18);
              let _1031_v7;
              let _nw187 = new _module.C0();
              _nw187.__ctor((((_1026_v2)[_module.__default.safeIndex((_this).f19, new BigNumber((_1026_v2).length))]) ? (_dafny.Seq.update(_1029_v5, _module.__default.safeIndex(new BigNumber((_1030_v6).cardinality()), new BigNumber((_1029_v5).length)), _1028_v4)) : (_1029_v5)), (_this).f18);
              _1031_v7 = _nw187;
              (globalState).f6 = (_this).f19;
              let _1032_v8;
              let _nw188 = Array((new BigNumber(24)).toNumber()).fill([]);
              _1032_v8 = _nw188;
              let _1033_v9;
              _1033_v9 = _module.D5.create_DC10(_1032_v8);
              let _1034_v10;
              _1034_v10 = _module.D7.create_DC17((_this).f18, (_this).f18, _1031_v7, (((_this).f18) ? (_1033_v9) : (_module.D5.create_DC10(_1032_v8))), (_1031_v7).f23);
              let _1035_v11;
              let _nw189 = Array((_dafny.ONE).toNumber()).fill(false);
              _1035_v11 = _nw189;
              let _index180 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1035_v11).length));
              (_1035_v11)[_index180] = (_1031_v7).f23;
              let _1036_v12;
              _1036_v12 = _dafny.Seq.of((_this).f19);
              let _1037_v13;
              _1037_v13 = new _dafny.CodePoint('m'.codePointAt(0));
              let _1038_v14;
              _1038_v14 = _module.D3.create_DC8(_1036_v12, _dafny.Seq.UnicodeFromString("brgrt"), (_1031_v7).f23, _1037_v13);
              let _1039_v15;
              _1039_v15 = _dafny.Seq.UnicodeFromString("xnu");
              let _1040_v16;
              _1040_v16 = _dafny.MultiSet.fromElements(new BigNumber((_1039_v15).length), (_this).f19);
              let _index181 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1035_v11).length));
              let _rhs140 = _module.D7.create_DC17((_1031_v7).fm17(new BigNumber((_1036_v12).length), new BigNumber(((_1038_v14).dtor_cf8).length), (_1031_v7).f23, globalState), (_1031_v7).f23, _1031_v7, (((_this).f18) ? (_1033_v9) : (_1033_v9)), (_1031_v7).f23);
              let _rhs141 = _module.__default.fm2(globalState);
              let _rhs142 = ((_1040_v16).Difference(_1040_v16)).IsDisjointFrom(_1040_v16);
              let _rhs143 = (_this).f18;
              let _lhs127 = _1035_v11;
              let _lhs128 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1035_v11).length));
              let _lhs129 = globalState;
              let _lhs130 = globalState;
              _1034_v10 = _rhs140;
              _lhs127[_lhs128] = _rhs141;
              _lhs129.f1 = _rhs142;
              _lhs130.f1 = _rhs143;
              (globalState).f1 = _module.__default.fm2(globalState);
            } else {
              r0 = (_dafny.ZERO).minus((_this).f19);
              (globalState).f1 = !((_this).f18);
              let _1041_v17;
              _1041_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_dafny.MultiSet.fromElements((_this).f19));
              let _1042_v18;
              _1042_v18 = _dafny.MultiSet.fromElements(new BigNumber((_1025_v1).length));
              let _1043_v19;
              _1043_v19 = _dafny.Set.fromElements((((_1041_v17).contains((_this).f18)) ? ((_1041_v17).get((_this).f18)) : (_1042_v18)));
              let _1044_v20;
              _1044_v20 = _dafny.Seq.UnicodeFromString("itmiqpxoa");
              let _1045_v21;
              _1045_v21 = _module.D3.create_DC7((_this).f19);
              let _1046_v22;
              let _nw190 = Array((new BigNumber(24)).toNumber()).fill(false);
              _1046_v22 = _nw190;
              let _1047_v23;
              _1047_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1046_v22,(_this).f18);
              let _1048_v24;
              _1048_v24 = _module.D10.create_DC27(new BigNumber((_1043_v19).length), _1044_v20, (_this).f19, _1045_v21, _1047_v23);
              (globalState).f6 = (_1048_v24).dtor_cf46;
              let _1049_v25;
              _1049_v25 = new _dafny.CodePoint('o'.codePointAt(0));
              _1049_v25 = _1049_v25;
              let _1050_v26;
              let _nw191 = new _module.C2();
              _nw191.__ctor(((_this).f19).plus((_this).f19), (_this).f19, _this.f17);
              _1050_v26 = _nw191;
            }
            let _1051_v27;
            _1051_v27 = _dafny.Seq.of((_this).f19, (_this).f19);
            let _1052_v28;
            _1052_v28 = _dafny.Seq.UnicodeFromString("cpvoudn");
            let _1053_v29;
            _1053_v29 = new _dafny.CodePoint('u'.codePointAt(0));
            let _1054_v30;
            _1054_v30 = _module.D3.create_DC8(_dafny.Seq.update(_1051_v27, _module.__default.safeIndex((_this).f19, new BigNumber((_1051_v27).length)), new BigNumber((_dafny.Seq.UnicodeFromString("nlep")).length)), _1052_v28, (_this).f18, _1053_v29);
            let _1055_v31;
            _1055_v31 = _dafny.Seq.of((_1054_v30).dtor_cf10, (_this).f18, false, (_this).f18);
            let _1056_v32;
            _1056_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,!((_this).f18));
            let _1057_v33;
            let _init28 = function (_1058_i3) {
              return (_this).f18;
            };
            let _nw192 = Array((new BigNumber(5)).toNumber());
            for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw192.length); _i0_28++) {
              _nw192[_i0_28] = _init28(new BigNumber(_i0_28));
            }
            _1057_v33 = _nw192;
            let _1059_v34;
            _1059_v34 = _dafny.Seq.of(_1057_v33);
            _1055_v31 = _dafny.Seq.update(_dafny.Seq.Concat((((((_1056_v32).contains((_this).f18)) ? ((_1056_v32).get((_this).f18)) : ((_this).f18))) ? (_1055_v31) : (_1055_v31)), _1055_v31), _module.__default.safeIndex(new BigNumber((_1059_v34).length), new BigNumber((_dafny.Seq.Concat((((((_1056_v32).contains((_this).f18)) ? ((_1056_v32).get((_this).f18)) : ((_this).f18))) ? (_1055_v31) : (_1055_v31)), _1055_v31)).length)), (_this).f18);
            _1056_v32 = (((_dafny.MultiSet.fromElements((_this).f19)).IsProperSubsetOf(_dafny.MultiSet.fromElements((_this).f19, (_this).f19, (_dafny.ZERO).minus((_this).f19)))) ? (_1056_v32) : ((_1056_v32).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18))));
            (globalState).f6 = (_this).f19;
          }
        }
      }
      let _1060_i4;
      _1060_i4 = _dafny.ZERO;
      L7: {
        while (!((_this).f18) || ((_this).f18)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1060_i4)) {
              break L7;
            }
            _1060_i4 = (_1060_i4).plus(_dafny.ONE);
            let _1061_v35;
            _1061_v35 = _dafny.Seq.UnicodeFromString("h");
            _1061_v35 = _module.__default.fm21(globalState);
            let _1062_v36;
            let _nw193 = Array((_dafny.ONE).toNumber());
            _nw193[(_dafny.ZERO).toNumber()] = (_this).f18;
            _1062_v36 = _nw193;
            let _1063_v37;
            let _nw194 = Array((new BigNumber(2)).toNumber());
            _nw194[(_dafny.ZERO).toNumber()] = _1062_v36;
            _nw194[(_dafny.ONE).toNumber()] = _1062_v36;
            _1063_v37 = _nw194;
            let _1064_v38;
            _1064_v38 = _module.D5.create_DC10(_1063_v37);
            let _pat_let_tv27 = _1063_v37;
            let _pat_let_tv28 = _1063_v37;
            let _source8 = function (_pat_let49_0) {
              return function (_1067_dt__update__tmp_h1) {
                return function (_pat_let52_0) {
                  return function (_1068_dt__update_hcf13_h1) {
                    return _module.D5.create_DC10(_1068_dt__update_hcf13_h1);
                  }(_pat_let52_0);
                }(_pat_let_tv28);
              }(_pat_let49_0);
            }(((false) ? (function (_pat_let50_0) {
              return function (_1065_dt__update__tmp_h0) {
                return function (_pat_let51_0) {
                  return function (_1066_dt__update_hcf13_h0) {
                    return _module.D5.create_DC10(_1066_dt__update_hcf13_h0);
                  }(_pat_let51_0);
                }(_pat_let_tv27);
              }(_pat_let50_0);
            }(_1064_v38)) : (_1064_v38)));
            if (_source8.is_DC11) {
              let _1069___mcc_h0 = (_source8).cf14;
              let _1070_cf14 = _1069___mcc_h0;
              let _1071_v39;
              _1071_v39 = new _dafny.CodePoint('x'.codePointAt(0));
              _1071_v39 = (_1061_v35)[_module.__default.safeIndex((_this).f19, new BigNumber((_1061_v35).length))];
              r1 = (_this).f18;
              (globalState).f1 = !(true);
              let _1072_v40;
              let _nw195 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
              _1072_v40 = _nw195;
              let _1073_v42;
              _1073_v42 = _module.D15.create_DC42(function () {
  let _coll40 = new _dafny.Set();
  for (const _compr_40 of _dafny.IntegerRange(new BigNumber(688), new BigNumber(977))) {
    let _1074_v41 = _compr_40;
    if (((new BigNumber(688)).isLessThanOrEqualTo(_1074_v41)) && ((_1074_v41).isLessThan(new BigNumber(977)))) {
      _coll40.add(_module.__default.safeDivisionInt(_1074_v41, _1070_cf14));
    }
  }
  return _coll40;
}());
              let _1075_v43;
              _1075_v43 = _dafny.Map.Empty.slice().updateUnsafe((_1073_v42).dtor_cf73,_1062_v36);
              let _index182 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_1072_v40).length));
              (_1072_v40)[_index182] = _1075_v43;
              let _index183 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_1072_v40).length));
              (_1072_v40)[_index183] = _1075_v43;
            } else if (_source8.is_DC12) {
              let _1076___mcc_h1 = (_source8).cf15;
              let _1077___mcc_h2 = (_source8).cf16;
              let _1078___mcc_h3 = (_source8).cf17;
              let _1079_cf17 = _1078___mcc_h3;
              let _1080_cf16 = _1077___mcc_h2;
              let _1081_cf15 = _1076___mcc_h1;
              _1061_v35 = _dafny.Seq.Concat(_module.__default.fm21(globalState), _dafny.Seq.Concat(_1061_v35, _1061_v35));
              let _1082_v44;
              _1082_v44 = _dafny.Seq.of((_this).f18, _1080_cf16);
              _1081_cf15 = ((new BigNumber((_1082_v44).length)).plus(new BigNumber(853))).multipliedBy(_1081_cf15);
              r0 = _1081_cf15;
              let _1083_v45;
              _1083_v45 = _dafny.Seq.of(new BigNumber(125));
              _1083_v45 = _1083_v45;
            } else if (_source8.is_DC10) {
              let _1084___mcc_h4 = (_source8).cf13;
              let _1085_cf13 = _1084___mcc_h4;
              let _index184 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_1062_v36).length));
              (_1062_v36)[_index184] = (_this).f18;
              let _index185 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_1062_v36).length));
              (_1062_v36)[_index185] = ((_this).f19).isLessThan((_this).f19);
              (globalState).f6 = ((_this).f19).plus((_this).f19);
              let _1086_v46;
              _1086_v46 = _dafny.MultiSet.fromElements((_this).f19);
              let _1087_v47;
              let _nw196 = new _module.C2();
              _nw196.__ctor(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f19), new BigNumber((_1086_v46).cardinality())), (_this).f19, (((_this).f18) ? (_this.f17) : (_this.f17)));
              _1087_v47 = _nw196;
              let _1088_v49;
              _1088_v49 = _dafny.Set.fromElements(false);
              let _1089_v50;
              _1089_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1088_v49).length),(_1087_v47).f20);
              let _1090_v52;
              let _init29 = function (_1091_i5) {
                return (_1091_i5).multipliedBy(new BigNumber(314));
              };
              let _nw197 = Array((new BigNumber(21)).toNumber());
              for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw197.length); _i0_29++) {
                _nw197[_i0_29] = _init29(new BigNumber(_i0_29));
              }
              _1090_v52 = _nw197;
              let _1092_v53;
              _1092_v53 = _dafny.Set.fromElements(_1090_v52);
              let _1093_v54;
              _1093_v54 = _dafny.Seq.of(_1092_v53);
              let _1094_v55;
              _1094_v55 = _dafny.Seq.of(new BigNumber(-736), _module.__default.fm1((_this).f18, globalState), new BigNumber(((_1093_v54)[_module.__default.safeIndex(_1087_v47.f21, new BigNumber((_1093_v54).length))]).length));
              let _1095_v56;
              _1095_v56 = _dafny.Map.Empty.slice().updateUnsafe((_1087_v47).f20,(_1062_v36)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_1062_v36).length))]);
              let _1096_v57;
              _1096_v57 = _dafny.Set.fromElements((_1087_v47).f20, (_this).f19);
              let _1097_v58;
              _1097_v58 = _dafny.Map.Empty.slice().updateUnsafe((_1062_v36)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_1062_v36).length))],new BigNumber((_1096_v57).length));
              let _1098_v59;
              let _nw198 = Array((new BigNumber(10)).toNumber());
              _nw198[(_dafny.ZERO).toNumber()] = _1087_v47.f21;
              _nw198[(_dafny.ONE).toNumber()] = (new BigNumber(-249)).plus((_this).f19);
              _nw198[(new BigNumber(2)).toNumber()] = (new BigNumber((function () {
                let _coll41 = new _dafny.Map();
                for (const _compr_41 of (_1089_v50).Keys.Elements) {
                  let _1099_v48 = _compr_41;
                  if ((_1089_v50).contains(_1099_v48)) {
                    _coll41.push([_module.__default.safeModuloInt(_1099_v48, new BigNumber((function () {
                      let _coll42 = new _dafny.Map();
                      for (const _compr_42 of _dafny.IntegerRange(new BigNumber(378), new BigNumber(24))) {
                        let _1100_v51 = _compr_42;
                        if (((new BigNumber(378)).isLessThanOrEqualTo(_1100_v51)) && ((_1100_v51).isLessThan(new BigNumber(24)))) {
                          _coll42.push([(_1100_v51).plus(new BigNumber(91)),!((_this).f18)]);
                        }
                      }
                      return _coll42;
                    }()).length)),(_this).f18]);
                  }
                }
                return _coll41;
              }()).length)).multipliedBy((_1087_v47).f20);
              _nw198[(new BigNumber(3)).toNumber()] = (_this).f19;
              _nw198[(new BigNumber(4)).toNumber()] = new BigNumber((_1094_v55).length);
              _nw198[(new BigNumber(5)).toNumber()] = _1087_v47.f21;
              _nw198[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1095_v56).length), (_dafny.ZERO).minus(_1087_v47.f21));
              _nw198[(new BigNumber(7)).toNumber()] = (((_1097_v58).contains((_1062_v36)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_1062_v36).length))])) ? ((_1097_v58).get((_1062_v36)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_1062_v36).length))])) : (new BigNumber(-546)));
              _nw198[(new BigNumber(8)).toNumber()] = ((_1087_v47).f20).multipliedBy((_1087_v47).f20);
              _nw198[(new BigNumber(9)).toNumber()] = _1087_v47.f21;
              _1098_v59 = _nw198;
              let _index186 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1090_v52).length));
              (_1090_v52)[_index186] = new BigNumber((_module.__default.fm28((_dafny.ZERO).minus((_1087_v47).f20), (_this).f18, (_1062_v36)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_1062_v36).length))], globalState)).length);
              let _index187 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1090_v52).length));
              (_1090_v52)[_index187] = (_dafny.ZERO).minus((_this).f19);
            } else {
              let _1101___mcc_h5 = (_source8).cf18;
              let _1102_cf18 = _1101___mcc_h5;
              let _1103_v60;
              _1103_v60 = _module.D13.create_DC37((_this).f18, (((_this).f18) ? (false) : ((_this).f18)));
              _1103_v60 = _module.D13.create_DC37((_this).f18, (_this).f18);
              (globalState).f6 = (_this).f19;
              let _1104_v61;
              _1104_v61 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),(_this).f19);
              let _1105_v62;
              let _nw199 = new _module.C2();
              _nw199.__ctor(new BigNumber((_1104_v61).length), (_this).f19, (_this.f17).Merge(_this.f17));
              _1105_v62 = _nw199;
              let _1106_v63;
              let _init30 = function (_1107_i6) {
                return _dafny.Seq.of((_this).f18);
              };
              let _nw200 = Array((new BigNumber(7)).toNumber());
              for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw200.length); _i0_30++) {
                _nw200[_i0_30] = _init30(new BigNumber(_i0_30));
              }
              _1106_v63 = _nw200;
              let _1108_v64;
              _1108_v64 = _dafny.Seq.of((_this).f18);
              let _index188 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_1106_v63).length));
              (_1106_v63)[_index188] = _1108_v64;
              let _index189 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_1106_v63).length));
              (_1106_v63)[_index189] = _dafny.Seq.Concat(_1108_v64, _1108_v64);
            }
            let _1109_v65;
            _1109_v65 = new _dafny.CodePoint('a'.codePointAt(0));
            let _1110_v66;
            _1110_v66 = _module.D5.create_DC12((_this).f19, (_this).f18, _1109_v65);
            let _1111_v67;
            _1111_v67 = _module.D5.create_DC13(_1110_v66);
            let _rhs144 = _1111_v67;
            let _rhs145 = (_this).f18;
            let _rhs146 = !(!(!(_module.__default.fm2(globalState))));
            let _lhs131 = globalState;
            _1111_v67 = _rhs144;
            r1 = _rhs145;
            _lhs131.f1 = _rhs146;
            let _1112_v68;
            let _init31 = function (_1113_i7) {
              return (_1113_i7).plus(new BigNumber((_dafny.Seq.UnicodeFromString("t")).length));
            };
            let _nw201 = Array((new BigNumber(18)).toNumber());
            for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw201.length); _i0_31++) {
              _nw201[_i0_31] = _init31(new BigNumber(_i0_31));
            }
            _1112_v68 = _nw201;
            let _1114_v69;
            _1114_v69 = _module.D3.create_DC6(_1112_v68);
            let _1115_v70;
            let _nw202 = Array((new BigNumber(29)).toNumber());
            _nw202[(_dafny.ZERO).toNumber()] = _module.D3.create_DC6(_1112_v68);
            _nw202[(_dafny.ONE).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(2)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(3)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(4)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(5)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(6)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(7)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(8)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(9)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(10)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(11)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(12)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(13)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(14)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(15)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(16)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(17)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(18)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(19)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(20)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(21)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(22)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(23)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(24)).toNumber()] = _module.D3.create_DC6(_1112_v68);
            _nw202[(new BigNumber(25)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(26)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(27)).toNumber()] = _1114_v69;
            _nw202[(new BigNumber(28)).toNumber()] = _1114_v69;
            _1115_v70 = _nw202;
            let _1116_v71;
            _1116_v71 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18),_1115_v70);
            let _1117_v72;
            _1117_v72 = _dafny.Set.fromElements((((_1116_v71).contains((_this).f18)) ? ((_1116_v71).get((_this).f18)) : (_1115_v70)), _1115_v70);
            let _1118_v73;
            _1118_v73 = _module.D9.create_DC23(_1117_v72, _1061_v35, (_this).f18, _1112_v68, _1109_v65);
            let _1119_v74;
            _1119_v74 = _module.D9.create_DC25(_1118_v73);
            let _1120_v75;
            _1120_v75 = _module.D9.create_DC25(_1119_v74);
            let _1121_v76;
            _1121_v76 = _dafny.MultiSet.fromElements(_1120_v75);
            _1121_v76 = _dafny.MultiSet.fromElements(_1120_v75, _1120_v75, _1120_v75);
          }
        }
      }
      let _1122_v77;
      let _nw203 = Array((new BigNumber(18)).toNumber()).fill([]);
      _1122_v77 = _nw203;
      _1122_v77 = _1122_v77;
      let _1123_v78;
      _1123_v78 = _module.D12.create_DC35();
      let _1124_v79;
      let _init32 = function (_1125_i8) {
        return _module.D8.create_DC20((_this).f19, (_this).f19, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18)).length)));
      };
      let _nw204 = Array((new BigNumber(9)).toNumber());
      for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw204.length); _i0_32++) {
        _nw204[_i0_32] = _init32(new BigNumber(_i0_32));
      }
      _1124_v79 = _nw204;
      let _1126_v80;
      _1126_v80 = _module.D14.create_DC38(_1021_v0);
      let _1127_v81;
      _1127_v81 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1126_v80),(_this).f18);
      let _1128_v82;
      _1128_v82 = _module.D8.create_DC20((_this).f19, new BigNumber((_1127_v81).length), (_this).f19);
      let _index190 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_1124_v79).length));
      (_1124_v79)[_index190] = _1128_v82;
      let _1129_v83;
      _1129_v83 = _dafny.Seq.of(true, true, (((_this).f18) ? ((_this).f18) : ((_this).f18)), ((_this).f18) || ((_this).f18), (_this).f18);
      let _1130_v84;
      _1130_v84 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.of((_this).f18, true), _dafny.Seq.of((_this).f18)), _1129_v83, _1129_v83, _dafny.Seq.Concat(_1129_v83, _1129_v83), _dafny.Seq.of((_this).f18, (_1129_v83)[_module.__default.safeIndex((_this).f19, new BigNumber((_1129_v83).length))], (_this).f18, _module.__default.fm2(globalState), (_this).f18));
      let _1131_v85;
      _1131_v85 = _module.D10.create_DC28(_module.__default.fm1((_this).f18, globalState));
      let _1132_v86;
      _1132_v86 = _dafny.Set.fromElements(_1131_v85);
      let _1133_v87;
      _1133_v87 = _dafny.Set.fromElements(_dafny.Set.fromElements(_module.D10.create_DC28((_module.D5.create_DC11(new BigNumber(65))).dtor_cf14), _1131_v85, _1131_v85, _1131_v85), _1132_v86);
      let _index191 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_1124_v79).length));
      let _rhs147 = _module.D12.create_DC35();
      let _rhs148 = (_this).f18;
      let _rhs149 = _1128_v82;
      let _rhs150 = (_1130_v84)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1133_v87).length)), new BigNumber((_1130_v84).length))];
      let _rhs151 = _module.__default.fm2(globalState);
      let _lhs132 = _1124_v79;
      let _lhs133 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_1124_v79).length));
      let _lhs134 = globalState;
      _1123_v78 = _rhs147;
      r1 = _rhs148;
      _lhs132[_lhs133] = _rhs149;
      _1129_v83 = _rhs150;
      _lhs134.f1 = _rhs151;
      r0 = _module.__default.safeDivisionInt((_this).f19, (_this).f19);
      r1 = !((_this).f18);
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let _1134_v0;
      _1134_v0 = _dafny.Seq.UnicodeFromString("elwpjwmet");
      let _rhs152 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_this).f19, (_this).f19), new BigNumber((_1134_v0).length));
      let _rhs153 = p0;
      let _lhs135 = globalState;
      let _lhs136 = globalState;
      _lhs135.f6 = _rhs152;
      _lhs136.f1 = _rhs153;
      (globalState).f1 = (new BigNumber(333)).isLessThan((((_this).f18) ? (_dafny.ONE) : ((_this).f19)));
      let _1135_v1;
      _1135_v1 = _dafny.MultiSet.fromElements((_this).f19);
      let _1136_v2;
      _1136_v2 = new _dafny.CodePoint('c'.codePointAt(0));
      let _1137_v3;
      _1137_v3 = _dafny.Map.Empty.slice().updateUnsafe((((_1135_v1).contains(new BigNumber((_dafny.Seq.UnicodeFromString("rccywrduq")).length))) ? ((_1135_v1).get(new BigNumber((_dafny.Seq.UnicodeFromString("rccywrduq")).length))) : (new BigNumber(-967))),_1136_v2);
      let _1138_v4;
      _1138_v4 = _dafny.Map.Empty.slice().updateUnsafe((((_1137_v3).contains((_this).f19)) ? ((_1137_v3).get((_this).f19)) : (_1136_v2)),(_dafny.ZERO).minus((_this).f19));
      _1138_v4 = (_1138_v4).update(((p0) ? (_1136_v2) : (_1136_v2)), (_dafny.ZERO).minus(new BigNumber(-416)));
      let _1139_i0;
      _1139_i0 = _dafny.ZERO;
      L8: {
        while (true) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1139_i0)) {
              break L8;
            }
            _1139_i0 = (_1139_i0).plus(_dafny.ONE);
            (globalState).f6 = (_this).f19;
            let _1140_v5;
            _1140_v5 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f18) ? ((_this).f18) : (true)),(_dafny.ZERO).minus((_this).f19));
            _1140_v5 = (_1140_v5).update((_this).f18, (_this).f19);
            (globalState).f6 = new BigNumber(551);
            let _1141_v6;
            _1141_v6 = _dafny.Seq.of(new BigNumber(576), (_this).f19, (_dafny.ZERO).minus((_this).f19), (_this).f19);
            let _1142_v7;
            _1142_v7 = _dafny.Set.fromElements(_1141_v6, _dafny.Seq.update(_1141_v6, _module.__default.safeIndex((_this).f19, new BigNumber((_1141_v6).length)), (_this).f19), _module.__default.fm32((_dafny.ZERO).minus((_this).f19), (_this).f19, (_this).f18, globalState));
            (globalState).f1 = (_1142_v7).IsSubsetOf(_dafny.Set.fromElements(_1141_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), function (_1143_i1) {
              return (_this).f19;
            }), _1141_v6, _1141_v6, _1141_v6));
          }
        }
      }
      if ((_this).f18) {
        let _1144_v8;
        _1144_v8 = _dafny.Set.fromElements(p0);
        _1144_v8 = _1144_v8;
        if (((!(true) || ((_this).f18)) ? (((_this).f19).isLessThan((_this).f19)) : (false))) {
          (_this).m13(globalState);
          _1134_v0 = _dafny.Seq.Concat(_1134_v0, _dafny.Seq.Concat(_1134_v0, _dafny.Seq.update(_1134_v0, _module.__default.safeIndex((_this).f19, new BigNumber((_1134_v0).length)), new _dafny.CodePoint('m'.codePointAt(0)))));
          let _1145_v9;
          _1145_v9 = _module.D3.create_DC7((_this).f19);
          let _1146_v10;
          _1146_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f19);
          let _1147_v11;
          let _nw205 = Array((new BigNumber(11)).toNumber());
          _nw205[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(-104), (_this).f19));
          _nw205[(_dafny.ONE).toNumber()] = (_this).f19;
          _nw205[(new BigNumber(2)).toNumber()] = (_this).f19;
          _nw205[(new BigNumber(3)).toNumber()] = (_this).f19;
          _nw205[(new BigNumber(4)).toNumber()] = (_this).f19;
          _nw205[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.of(p0, p0)).length);
          _nw205[(new BigNumber(6)).toNumber()] = (_1145_v9).dtor_cf7;
          _nw205[(new BigNumber(7)).toNumber()] = ((_this).f19).minus((_this).f19);
          _nw205[(new BigNumber(8)).toNumber()] = new BigNumber((_1134_v0).length);
          _nw205[(new BigNumber(9)).toNumber()] = (_this).f19;
          _nw205[(new BigNumber(10)).toNumber()] = (new BigNumber((_1146_v10).length)).plus(new BigNumber(523));
          _1147_v11 = _nw205;
          let _index192 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1147_v11).length));
          (_1147_v11)[_index192] = (_this).f19;
          let _1148_v12;
          _1148_v12 = _dafny.Seq.of(((_this).f19).minus((_dafny.ZERO).minus((_this).f19)), (_this).f19);
          let _index193 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_1147_v11).length));
          (_1147_v11)[_index193] = (_1148_v12)[_module.__default.safeIndex((((_this).f18) ? (new BigNumber((_1134_v0).length)) : (new BigNumber((_1134_v0).length))), new BigNumber((_1148_v12).length))];
          (globalState).f6 = (_this).f19;
          let _1149_v14;
          _1149_v14 = _dafny.Map.Empty.slice().updateUnsafe((_1147_v11)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_1147_v11).length))],p0);
          _1136_v2 = (((_1137_v3).contains(new BigNumber((function () {
            let _coll44 = new _dafny.Map();
            for (const _compr_44 of (_1149_v14).Keys.Elements) {
              let _1151_v13 = _compr_44;
              if ((_1149_v14).contains(_1151_v13)) {
                _coll44.push([_module.__default.safeModuloInt(_1151_v13, (_this).f19),_1136_v2]);
              }
            }
            return _coll44;
          }()).length))) ? ((_1137_v3).get(new BigNumber((function () {
            let _coll43 = new _dafny.Map();
            for (const _compr_43 of (_1149_v14).Keys.Elements) {
              let _1150_v13 = _compr_43;
              if ((_1149_v14).contains(_1150_v13)) {
                _coll43.push([_module.__default.safeModuloInt(_1150_v13, (_this).f19),_1136_v2]);
              }
            }
            return _coll43;
          }()).length))) : (_1136_v2));
        } else {
          let _1152_v16;
          let _init33 = function (_1153_i2) {
            return _module.D0.create_DC2(_module.D0.create_DC1((_this).f19));
          };
          let _nw206 = Array((new BigNumber(27)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw206.length); _i0_33++) {
            _nw206[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1152_v16 = _nw206;
          let _1154_v17;
          let _init34 = function (_1155_i3) {
            return _module.__default.safeDivisionInt(_1155_i3, (_this).f19);
          };
          let _nw207 = Array((new BigNumber(22)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw207.length); _i0_34++) {
            _nw207[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1154_v17 = _nw207;
          _module.__default.m0(function () {
            let _coll45 = new _dafny.Set();
            for (const _compr_45 of (_dafny.Seq.of(_module.__default.fm1((_this).f18, globalState))).Elements) {
              let _1156_v15 = _compr_45;
              if (_dafny.Seq.contains(_dafny.Seq.of(_module.__default.fm1((_this).f18, globalState)), _1156_v15)) {
                _coll45.add((_1156_v15).minus(new BigNumber(701)));
              }
            }
            return _coll45;
          }(), _1152_v16, _1154_v17, _1134_v0, globalState);
          let _1157_v18;
          _1157_v18 = _module.D3.create_DC6(_1154_v17);
          let _1158_v19;
          _1158_v19 = _dafny.Seq.of(_1157_v18, _module.D3.create_DC6(_1154_v17), _1157_v18, _1157_v18, _module.D3.create_DC6(_1154_v17));
          let _1159_v20;
          let _nw208 = new _module.C0();
          _nw208.__ctor(_dafny.Seq.update(_dafny.Seq.Concat(_1158_v19, _1158_v19), _module.__default.safeIndex(new BigNumber(52), new BigNumber((_dafny.Seq.Concat(_1158_v19, _1158_v19)).length)), _1157_v18), (_this).f18);
          _1159_v20 = _nw208;
          let _index194 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1154_v17).length));
          (_1154_v17)[_index194] = (_this).f19;
          let _index195 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1154_v17).length));
          (_1154_v17)[_index195] = _module.__default.fm1((new BigNumber(253)).isEqualTo((_this).f19), globalState);
          let _1160_v21;
          _1160_v21 = _dafny.Seq.of(p0);
          let _1161_v22;
          _1161_v22 = _module.D11.create_DC31(_1160_v21, (_1154_v17)[_module.__default.safeIndex(new BigNumber(392), new BigNumber((_1154_v17).length))], (_this).f19, _module.__default.fm1(true, globalState), (_this).f19);
          let _1162_v23;
          _1162_v23 = _dafny.Map.Empty.slice().updateUnsafe((_1159_v20).f23,!((_this).f18));
          let _1163_v24;
          let _nw209 = Array((new BigNumber(13)).toNumber());
          _nw209[(_dafny.ZERO).toNumber()] = false;
          _nw209[(_dafny.ONE).toNumber()] = (new BigNumber((_1134_v0).length)).isLessThan((_this).f19);
          _nw209[(new BigNumber(2)).toNumber()] = (_1159_v20).f23;
          _nw209[(new BigNumber(3)).toNumber()] = (_1159_v20).f23;
          _nw209[(new BigNumber(4)).toNumber()] = (_this).f18;
          _nw209[(new BigNumber(5)).toNumber()] = p0;
          _nw209[(new BigNumber(6)).toNumber()] = (_1159_v20).f23;
          _nw209[(new BigNumber(7)).toNumber()] = true;
          _nw209[(new BigNumber(8)).toNumber()] = (_1159_v20).fm17(new BigNumber(((_1161_v22).dtor_cf54).length), (_dafny.ZERO).minus(new BigNumber((_1162_v23).length)), !(!(p0)), globalState);
          _nw209[(new BigNumber(9)).toNumber()] = (_1159_v20).f23;
          _nw209[(new BigNumber(10)).toNumber()] = p0;
          _nw209[(new BigNumber(11)).toNumber()] = !(false);
          _nw209[(new BigNumber(12)).toNumber()] = p0;
          _1163_v24 = _nw209;
          let _index196 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_1163_v24).length));
          (_1163_v24)[_index196] = (_this).f18;
          let _index197 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_1163_v24).length));
          (_1163_v24)[_index197] = (_1159_v20).fm17(_module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.fm1((_this).f18, globalState)), (_1154_v17)[_module.__default.safeIndex(new BigNumber(392), new BigNumber((_1154_v17).length))]), (_this).f19, p0, globalState);
          let _1164_v25;
          let _nw210 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
          _1164_v25 = _nw210;
          _1164_v25 = _1164_v25;
        }
        let _1165_v26;
        let _nw211 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1165_v26 = _nw211;
        let _index198 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1165_v26).length));
        (_1165_v26)[_index198] = !(!(p0));
        let _index199 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1165_v26).length));
        let _rhs154 = (_this).f18;
        let _rhs155 = _module.__default.safeModuloInt((_this).f19, ((_this).f19).multipliedBy((_dafny.ZERO).minus((_this).f19)));
        let _rhs156 = _1134_v0;
        let _rhs157 = _dafny.Seq.Concat(_1134_v0, _1134_v0);
        let _lhs137 = _1165_v26;
        let _lhs138 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1165_v26).length));
        let _lhs139 = globalState;
        _lhs137[_lhs138] = _rhs154;
        _lhs139.f6 = _rhs155;
        _1134_v0 = _rhs156;
        _1134_v0 = _rhs157;
        (globalState).f6 = (_this).f19;
        let _1166_v27;
        let _nw212 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1166_v27 = _nw212;
        let _index200 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_1166_v27).length));
        (_1166_v27)[_index200] = _1136_v2;
        let _index201 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_1166_v27).length));
        (_1166_v27)[_index201] = (_1134_v0)[_module.__default.safeIndex((_this).f19, new BigNumber((_1134_v0).length))];
      } else {
        (globalState).f1 = p0;
        let _1167_v28;
        let _nw213 = Array((new BigNumber(23)).toNumber()).fill([]);
        _1167_v28 = _nw213;
        let _1168_v29;
        _1168_v29 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1167_v28);
        let _1169_v30;
        let _nw214 = Array((new BigNumber(24)).toNumber());
        _nw214[(_dafny.ZERO).toNumber()] = ((p0) ? (_1167_v28) : (_1167_v28));
        _nw214[(_dafny.ONE).toNumber()] = (((_1168_v29).contains(false)) ? ((_1168_v29).get(false)) : (_1167_v28));
        _nw214[(new BigNumber(2)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(3)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(4)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(5)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(6)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(7)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(8)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(9)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(10)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(11)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(12)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(13)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(14)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(15)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(16)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(17)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(18)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(19)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(20)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(21)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(22)).toNumber()] = _1167_v28;
        _nw214[(new BigNumber(23)).toNumber()] = _1167_v28;
        _1169_v30 = _nw214;
        let _index202 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1169_v30).length));
        (_1169_v30)[_index202] = _1167_v28;
        let _index203 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1169_v30).length));
        (_1169_v30)[_index203] = _1167_v28;
        if ((p0) && (p0)) {
          _1167_v28 = (_1169_v30)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1169_v30).length))];
          let _1170_v31;
          _1170_v31 = _dafny.Seq.of((_this).f19);
          let _1171_v32;
          _1171_v32 = _dafny.Map.Empty.slice().updateUnsafe(!(true),(_this).f18);
          let _1172_v33;
          let _nw215 = Array((new BigNumber(7)).toNumber());
          _nw215[(_dafny.ZERO).toNumber()] = (_this).f19;
          _nw215[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber((_1170_v31).length))).length);
          _nw215[(new BigNumber(2)).toNumber()] = (_this).f19;
          _nw215[(new BigNumber(3)).toNumber()] = new BigNumber((_1134_v0).length);
          _nw215[(new BigNumber(4)).toNumber()] = (_this).f19;
          _nw215[(new BigNumber(5)).toNumber()] = new BigNumber((_1170_v31).length);
          _nw215[(new BigNumber(6)).toNumber()] = new BigNumber((_1171_v32).length);
          _1172_v33 = _nw215;
          let _1173_v34;
          _1173_v34 = _module.D3.create_DC6(_1172_v33);
          let _1174_v35;
          _1174_v35 = _dafny.Seq.of(_1173_v34, _1173_v34, _1173_v34);
          let _1175_v36;
          let _nw216 = new _module.C0();
          _nw216.__ctor(_dafny.Seq.Concat(_1174_v35, _1174_v35), p0);
          _1175_v36 = _nw216;
          let _1176_v37;
          _1176_v37 = _dafny.Seq.of((_1175_v36).f23);
          let _1177_v38;
          let _nw217 = Array((new BigNumber(5)).toNumber());
          _nw217[(_dafny.ZERO).toNumber()] = (_1175_v36).f23;
          _nw217[(_dafny.ONE).toNumber()] = p0;
          _nw217[(new BigNumber(2)).toNumber()] = (_1176_v37)[_module.__default.safeIndex((_this).f19, new BigNumber((_1176_v37).length))];
          _nw217[(new BigNumber(3)).toNumber()] = (_this).f18;
          _nw217[(new BigNumber(4)).toNumber()] = p0;
          _1177_v38 = _nw217;
          let _1178_v39;
          _1178_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1137_v3,_1177_v38);
          _1178_v39 = _1178_v39;
          (globalState).f1 = !(false) || ((_1175_v36).f23);
          let _1179_v40;
          let _nw218 = new _module.C1();
          _nw218.__ctor((_this).f19);
          _1179_v40 = _nw218;
          let _1180_v41;
          _1180_v41 = _module.D7.create_DC15(_1179_v40);
          let _pat_let_tv29 = _1180_v41;
          let _1181_v42;
          let _nw219 = Array((new BigNumber(17)).toNumber());
          _nw219[(_dafny.ZERO).toNumber()] = _1180_v41;
          _nw219[(_dafny.ONE).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(2)).toNumber()] = function (_pat_let53_0) {
            return function (_1182_dt__update__tmp_h0) {
              return function (_pat_let54_0) {
                return function (_1183_dt__update_hcf20_h0) {
                  return _module.D7.create_DC15(_1183_dt__update_hcf20_h0);
                }(_pat_let54_0);
              }((_pat_let_tv29).dtor_cf20);
            }(_pat_let53_0);
          }(_1180_v41);
          _nw219[(new BigNumber(3)).toNumber()] = _module.D7.create_DC15(_1179_v40);
          _nw219[(new BigNumber(4)).toNumber()] = _module.D7.create_DC15(_1179_v40);
          _nw219[(new BigNumber(5)).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(6)).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(7)).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(8)).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(9)).toNumber()] = _module.D7.create_DC15(_1179_v40);
          _nw219[(new BigNumber(10)).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(11)).toNumber()] = _module.D7.create_DC15(_1179_v40);
          _nw219[(new BigNumber(12)).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(13)).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(14)).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(15)).toNumber()] = _1180_v41;
          _nw219[(new BigNumber(16)).toNumber()] = _1180_v41;
          _1181_v42 = _nw219;
          let _1184_v43;
          _1184_v43 = _dafny.Seq.of(_1181_v42);
          let _1185_v44;
          let _nw220 = Array((new BigNumber(25)).toNumber());
          _nw220[(_dafny.ZERO).toNumber()] = _1181_v42;
          _nw220[(_dafny.ONE).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(2)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(3)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(4)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(5)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(6)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(7)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(8)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(9)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(10)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(11)).toNumber()] = (_1184_v43)[_module.__default.safeIndex((_this).f19, new BigNumber((_1184_v43).length))];
          _nw220[(new BigNumber(12)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(13)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(14)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(15)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(16)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(17)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(18)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(19)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(20)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(21)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(22)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(23)).toNumber()] = _1181_v42;
          _nw220[(new BigNumber(24)).toNumber()] = _1181_v42;
          _1185_v44 = _nw220;
          let _1186_v45;
          _1186_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1185_v44);
          let _1187_v46;
          _1187_v46 = _module.D5.create_DC12((_this).f19, _module.__default.fm2(globalState), _1136_v2);
          let _1188_v47;
          _1188_v47 = _dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC26((((_1186_v45).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1187_v46).dtor_cf16,(_this).f19)).length))) ? ((_1186_v45).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1187_v46).dtor_cf16,(_this).f19)).length))) : (_1185_v44))),((_this).f19).multipliedBy((_this).f19));
          let _1189_v48;
          _1189_v48 = _module.D10.create_DC26(_1185_v44);
          _1188_v47 = (_1188_v47).update(_1189_v48, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(281)), ((_1190_v2) => function (_1191_i4) {
            return _1190_v2;
          })(_1136_v2))).length));
        } else {
          let _1192_v49;
          _1192_v49 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_dafny.Seq.UnicodeFromString("dpswinn"), _1134_v0),(new BigNumber(118)).isLessThanOrEqualTo((_dafny.ZERO).minus((_this).f19)));
          (globalState).f1 = (((_1192_v49).contains(!((_this).f18))) ? ((_1192_v49).get(!((_this).f18))) : (!(true)));
          let _1193_v50;
          _1193_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1135_v1).cardinality()),p0);
          let _1194_v51;
          let _nw221 = Array((new BigNumber(29)).toNumber());
          _nw221[(_dafny.ZERO).toNumber()] = p0;
          _nw221[(_dafny.ONE).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(2)).toNumber()] = true;
          _nw221[(new BigNumber(3)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(4)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(5)).toNumber()] = !((_this).f18);
          _nw221[(new BigNumber(6)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(7)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(8)).toNumber()] = p0;
          _nw221[(new BigNumber(9)).toNumber()] = p0;
          _nw221[(new BigNumber(10)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(11)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(12)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(13)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(14)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(15)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(16)).toNumber()] = p0;
          _nw221[(new BigNumber(17)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(18)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(19)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(20)).toNumber()] = p0;
          _nw221[(new BigNumber(21)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(22)).toNumber()] = p0;
          _nw221[(new BigNumber(23)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(24)).toNumber()] = (_this).f18;
          _nw221[(new BigNumber(25)).toNumber()] = p0;
          _nw221[(new BigNumber(26)).toNumber()] = p0;
          _nw221[(new BigNumber(27)).toNumber()] = _module.__default.fm2(globalState);
          _nw221[(new BigNumber(28)).toNumber()] = (((_1193_v50).contains((_this).f19)) ? ((_1193_v50).get((_this).f19)) : (p0));
          _1194_v51 = _nw221;
          let _1195_v52;
          let _nw222 = Array((_dafny.ONE).toNumber());
          _nw222[(_dafny.ZERO).toNumber()] = _1194_v51;
          _1195_v52 = _nw222;
          let _index204 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_1195_v52).length));
          (_1195_v52)[_index204] = _1194_v51;
          let _index205 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_1195_v52).length));
          (_1195_v52)[_index205] = _1194_v51;
          (globalState).f1 = p0;
          (globalState).f1 = false;
          _1193_v50 = (_1193_v50).update((_this).f19, (((_this).f18) ? ((_this).f18) : (p0)));
        }
        (globalState).f1 = (_this).f18;
        let _1196_v53;
        _1196_v53 = _dafny.Seq.of(p0);
        let _1197_v54;
        _1197_v54 = _module.D0.create_DC0(!(p0));
        let _1198_v55;
        _1198_v55 = _dafny.Seq.of(_1197_v54);
        let _1199_v56;
        _1199_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_module.__default.fm35(_1196_v53, globalState), _1198_v55),(_this).f19);
        _1199_v56 = (_1199_v56).update((((_this).f18) ? (_1198_v55) : (_1198_v55)), (_this).f19);
      }
      let _1200_v57;
      _1200_v57 = _dafny.MultiSet.fromElements(p0, p0);
      let _1201_v58;
      _1201_v58 = _dafny.MultiSet.fromElements(_1136_v2, _1136_v2, _1136_v2, _1136_v2, _1136_v2);
      let _1202_v59;
      _1202_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1200_v57,_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(115)), function (_1203_i5) {
        return new BigNumber((_dafny.Seq.of((_this).f19)).length);
      }), _dafny.Seq.of(new BigNumber((_1201_v58).cardinality()), (_this).f19)));
      let _1204_v60;
      _1204_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_module.__default.safeModuloInt(new BigNumber(151), (_this).f19));
      let _1205_v61;
      _1205_v61 = _module.D12.create_DC34(_dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19));
      let _rhs158 = _1202_v59;
      let _rhs159 = _module.__default.fm36(_1136_v2, _1205_v61, globalState);
      _1202_v59 = _rhs158;
      _1204_v60 = _rhs159;
      return;
    }
    m11(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      (_this).m13(globalState);
      let _1206_v0;
      _1206_v0 = _dafny.Set.fromElements(p1, _module.__default.fm1((_this).f18, globalState));
      let _1207_v1;
      let _nw223 = Array((new BigNumber(27)).toNumber()).fill(_module.D0.Default());
      _1207_v1 = _nw223;
      let _1208_v2;
      let _nw224 = Array((new BigNumber(2)).toNumber());
      _nw224[(_dafny.ZERO).toNumber()] = p1;
      _nw224[(_dafny.ONE).toNumber()] = p1;
      _1208_v2 = _nw224;
      _module.__default.m0(_1206_v0, _1207_v1, _1208_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(747)), function (_1209_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }), globalState);
      let _1210_v3;
      let _nw225 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
      _1210_v3 = _nw225;
      let _1211_v4;
      _1211_v4 = _dafny.Seq.of((_this).f19, p1);
      let _index206 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_1210_v3).length));
      (_1210_v3)[_index206] = _dafny.Seq.Concat(_1211_v4, _1211_v4);
      let _index207 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_1210_v3).length));
      (_1210_v3)[_index207] = _dafny.Seq.of(p1);
      (globalState).f6 = ((_dafny.ZERO).minus((_this).f19)).plus(p1);
      (globalState).f1 = (_this).f18;
      (globalState).f6 = ((_1210_v3)[_module.__default.safeIndex(new BigNumber(519), new BigNumber((_1210_v3).length))])[_module.__default.safeIndex((_this).f19, new BigNumber(((_1210_v3)[_module.__default.safeIndex(new BigNumber(519), new BigNumber((_1210_v3).length))]).length))];
      let _1212_v5;
      _1212_v5 = new _dafny.CodePoint('r'.codePointAt(0));
      r0 = _1212_v5;
      return r0;
    }
    m12(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      r1 = _dafny.Seq.Concat(_module.__default.fm21(globalState), _dafny.Seq.UnicodeFromString("klxjhxnre"));
      (globalState).f1 = (_this).f18;
      let _hi8 = (_this).f19;
      for (let _1213_i0 = p0; _1213_i0.isLessThan(_hi8); _1213_i0 = _1213_i0.plus(_dafny.ONE)) {
        let _1214_v0;
        let _nw226 = Array((new BigNumber(2)).toNumber()).fill(false);
        _1214_v0 = _nw226;
        let _1215_v1;
        _1215_v1 = _dafny.Seq.of((_this).f18, true, (_this).f18, (_this).f18, _module.__default.fm2(globalState));
        let _index208 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1214_v0).length));
        (_1214_v0)[_index208] = (_1215_v1)[_module.__default.safeIndex(new BigNumber((p3).length), new BigNumber((_1215_v1).length))];
        let _1216_v2;
        _1216_v2 = _dafny.MultiSet.fromElements((_this).f18, (_this).f18);
        let _1217_v3;
        _1217_v3 = _dafny.Map.Empty.slice().updateUnsafe((_1216_v2).Intersect(_dafny.MultiSet.FromArray(_1215_v1)),p2);
        let _1218_v4;
        _1218_v4 = _dafny.MultiSet.fromElements(p0);
        let _1219_v5;
        _1219_v5 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-496));
        let _1220_v6;
        _1220_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1213_i0,(_this).f18);
        let _1221_v7;
        _1221_v7 = _module.D7.create_DC16((_this).f18, new BigNumber(((_1219_v5).update((_this).f18, (_this).f19)).length), (((_1220_v6).contains(p0)) ? ((_1220_v6).get(p0)) : ((_this).f18)));
        let _1222_v8;
        _1222_v8 = _module.D13.create_DC37((_this).f18, (_1221_v7).dtor_cf21);
        let _1223_v10;
        _1223_v10 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_this).f18), _dafny.MultiSet.FromArray(_1215_v1));
        let _1224_v11;
        _1224_v11 = _dafny.Set.fromElements((_this).f18);
        let _index209 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1214_v0).length));
        let _rhs160 = (_1218_v4).IsDisjointFrom(_1218_v4);
        let _rhs161 = (((_1222_v8).dtor_cf66) ? (_1217_v3) : (function () {
          let _coll46 = new _dafny.Map();
          for (const _compr_46 of (_1223_v10).Elements) {
            let _1225_v9 = _compr_46;
            if (_dafny.Seq.contains(_1223_v10, _1225_v9)) {
              _coll46.push([_1225_v9,p3]);
            }
          }
          return _coll46;
        }()));
        let _rhs162 = !((_1224_v11).IsProperSubsetOf(_1224_v11));
        let _rhs163 = ((!(_1224_v11).contains((_this).f18)) ? (p3) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), function (_1226_i1) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })));
        let _lhs140 = _1214_v0;
        let _lhs141 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1214_v0).length));
        _lhs140[_lhs141] = _rhs160;
        _1217_v3 = _rhs161;
        r0 = _rhs162;
        r1 = _rhs163;
        let _1227_v12;
        let _nw227 = new _module.C2();
        _nw227.__ctor(p1, (_this).f19, _this.f17);
        _1227_v12 = _nw227;
        let _1228_v13;
        _1228_v13 = new _dafny.CodePoint('r'.codePointAt(0));
        let _1229_v14;
        _1229_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1219_v5).length),new BigNumber((_dafny.Set.fromElements(_1228_v13, _1228_v13)).length));
        let _1230_v15;
        _1230_v15 = _dafny.Seq.of(new BigNumber((p2).length));
        let _1231_v16;
        let _nw228 = Array((new BigNumber(17)).toNumber());
        _nw228[(_dafny.ZERO).toNumber()] = _1213_i0;
        _nw228[(_dafny.ONE).toNumber()] = (_1227_v12).f20;
        _nw228[(new BigNumber(2)).toNumber()] = (_1227_v12).f20;
        _nw228[(new BigNumber(3)).toNumber()] = (_1227_v12).f20;
        _nw228[(new BigNumber(4)).toNumber()] = (new BigNumber((p3).length)).minus(new BigNumber(859));
        _nw228[(new BigNumber(5)).toNumber()] = (_1227_v12).f20;
        _nw228[(new BigNumber(6)).toNumber()] = (_1227_v12).f20;
        _nw228[(new BigNumber(7)).toNumber()] = (new BigNumber((_dafny.Seq.of(_1214_v0)).length)).plus(new BigNumber((_1224_v11).length));
        _nw228[(new BigNumber(8)).toNumber()] = p1;
        _nw228[(new BigNumber(9)).toNumber()] = (p1).multipliedBy(_module.__default.fm1((_1214_v0)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_1214_v0).length))], globalState));
        _nw228[(new BigNumber(10)).toNumber()] = (new BigNumber((_dafny.Seq.of((_this).f18)).length)).multipliedBy(_1213_i0);
        _nw228[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_1213_i0);
        _nw228[(new BigNumber(12)).toNumber()] = new BigNumber((_1224_v11).length);
        _nw228[(new BigNumber(13)).toNumber()] = p1;
        _nw228[(new BigNumber(14)).toNumber()] = (new BigNumber(-761)).multipliedBy(new BigNumber((_1229_v14).length));
        _nw228[(new BigNumber(15)).toNumber()] = ((_this).f19).multipliedBy((_this).f19);
        _nw228[(new BigNumber(16)).toNumber()] = ((!((_this).f18)) ? (new BigNumber((_1230_v15).length)) : ((_this).f19));
        _1231_v16 = _nw228;
        let _1232_v17;
        _1232_v17 = _module.D11.create_DC31(_1215_v1, _1213_i0, p1, p1, p0);
        let _index210 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1231_v16).length));
        (_1231_v16)[_index210] = (_1232_v17).dtor_cf56;
        let _index211 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1231_v16).length));
        (_1231_v16)[_index211] = _1213_i0;
        let _index212 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1231_v16).length));
        (_1231_v16)[_index212] = new BigNumber(515);
      }
      (globalState).f6 = (_dafny.ZERO).minus(p1);
      (globalState).f6 = p1;
      if ((_this).f18) {
        let _1233_v18;
        let _nw229 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _1233_v18 = _nw229;
        let _1234_v19;
        _1234_v19 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber(-157)).isEqualTo(p0),_1233_v18);
        (globalState).f9 = (((_1234_v19).contains(!((_this).f18))) ? ((_1234_v19).get(!((_this).f18))) : (_1233_v18));
        let _index213 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_1233_v18).length));
        (_1233_v18)[_index213] = p0;
        let _index214 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_1233_v18).length));
        (_1233_v18)[_index214] = (_this).f19;
        let _1235_v20;
        _1235_v20 = _dafny.MultiSet.fromElements((_1233_v18)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_1233_v18).length))], p1);
        let _1236_v22;
        _1236_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f18);
        let _1237_v23;
        _1237_v23 = _dafny.Set.fromElements((function () {
          let _coll47 = new _dafny.Set();
          for (const _compr_47 of (_1235_v20).Elements) {
            let _1238_v21 = _compr_47;
            if ((_1235_v20).contains(_1238_v21)) {
              _coll47.add((_1238_v21).plus(new BigNumber(784)));
            }
          }
          return _coll47;
        }()).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_1236_v22).length))));
        _1237_v23 = _1237_v23;
        let _1239_v24;
        _1239_v24 = _module.D3.create_DC6(_1233_v18);
        let _1240_v25;
        _1240_v25 = _dafny.Seq.of(_module.D3.create_DC6(_1233_v18), _1239_v24, _1239_v24, _1239_v24, _1239_v24);
        let _1241_v26;
        let _nw230 = new _module.C0();
        _nw230.__ctor(_dafny.Seq.Concat(_1240_v25, _1240_v25), (_this).f18);
        _1241_v26 = _nw230;
        let _1242_v27;
        _1242_v27 = _dafny.Seq.of(p2);
        if (_dafny.areEqual(_dafny.Seq.Concat(_1242_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-375)), ((_1243_p3) => function (_1244_i2) {
          return _1243_p3;
        })(p3))), _1242_v27)) {
          let _1245_v28;
          _1245_v28 = _dafny.Seq.of(p0);
          let _1246_v29;
          _1246_v29 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f18),_1245_v28);
          let _1247_v30;
          _1247_v30 = _dafny.Seq.of((((_1246_v29).contains((_this).f18)) ? ((_1246_v29).get((_this).f18)) : (_1245_v28)));
          _1247_v30 = _1247_v30;
          (globalState).f6 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f19));
          let _1248_v31;
          let _nw231 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
          _1248_v31 = _nw231;
          _1248_v31 = _1248_v31;
          let _1249_v32;
          _1249_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1241_v26).f23,(_1233_v18)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_1233_v18).length))]);
          let _1250_v33;
          _1250_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1249_v32,new BigNumber((_module.__default.fm28(p1, false, (_1241_v26).f23, globalState)).length));
          _1250_v33 = (_1250_v33).update(_1249_v32, (_this).f19);
          let _1251_v34;
          let _nw232 = new _module.C1();
          _nw232.__ctor((new BigNumber(721)).plus(new BigNumber((p2).length)));
          _1251_v34 = _nw232;
        } else {
          let _1252_v35;
          _1252_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_1233_v18)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_1233_v18).length))], (_this).f19),!((_1241_v26).f23));
          let _index215 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_1233_v18).length));
          (_1233_v18)[_index215] = new BigNumber((function () {
            let _coll48 = new _dafny.Set();
            for (const _compr_48 of (_1252_v35).Keys.Elements) {
              let _1253_v36 = _compr_48;
              if ((_1252_v35).contains(_1253_v36)) {
                _coll48.add(_1253_v36);
              }
            }
            return _coll48;
          }()).length);
          (globalState).f6 = ((_1233_v18)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_1233_v18).length))]).minus(new BigNumber(23));
          (globalState).f1 = false;
          let _1254_v37;
          let _nw233 = Array((new BigNumber(23)).toNumber()).fill([]);
          _1254_v37 = _nw233;
          let _index216 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1254_v37).length));
          (_1254_v37)[_index216] = _1233_v18;
          let _1255_v38;
          let _nw234 = Array((new BigNumber(10)).toNumber());
          _nw234[(_dafny.ZERO).toNumber()] = _1239_v24;
          _nw234[(_dafny.ONE).toNumber()] = _1239_v24;
          _nw234[(new BigNumber(2)).toNumber()] = _1239_v24;
          _nw234[(new BigNumber(3)).toNumber()] = _module.D3.create_DC6(_1233_v18);
          _nw234[(new BigNumber(4)).toNumber()] = _1239_v24;
          _nw234[(new BigNumber(5)).toNumber()] = _1239_v24;
          _nw234[(new BigNumber(6)).toNumber()] = _1239_v24;
          _nw234[(new BigNumber(7)).toNumber()] = _1239_v24;
          _nw234[(new BigNumber(8)).toNumber()] = _1239_v24;
          _nw234[(new BigNumber(9)).toNumber()] = _1239_v24;
          _1255_v38 = _nw234;
          let _1256_v39;
          _1256_v39 = _dafny.Set.fromElements(_1255_v38, _1255_v38);
          let _1257_v40;
          _1257_v40 = new _dafny.CodePoint('n'.codePointAt(0));
          let _1258_v41;
          _1258_v41 = _module.D9.create_DC23(_1256_v39, p2, !((_1241_v26).f23), _1233_v18, _1257_v40);
          let _index217 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1254_v37).length));
          (_1254_v37)[_index217] = (_1258_v41).dtor_cf40;
          let _1259_v42;
          _1259_v42 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1235_v20);
          let _1260_v43;
          _1260_v43 = _dafny.Seq.of(_this.f17, _this.f17, _this.f17);
          let _1261_v44;
          _1261_v44 = _dafny.Seq.of(new BigNumber(531));
          let _1262_v45;
          _1262_v45 = _1261_v44;
          let _1263_v46;
          let _nw235 = new _module.C2();
          _nw235.__ctor(new BigNumber((_1259_v42).length), (_1233_v18)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_1233_v18).length))], (_1260_v43)[_module.__default.safeIndex((_1261_v44)[_module.__default.safeIndex(new BigNumber(((_1262_v45)).length), new BigNumber((_1261_v44).length))], new BigNumber((_1260_v43).length))]);
          _1263_v46 = _nw235;
        }
      } else {
        let _1264_v47;
        _1264_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,p2);
        let _1265_v48;
        let _nw236 = new _module.C1();
        _nw236.__ctor(p1);
        _1265_v48 = _nw236;
        let _1266_v49;
        _1266_v49 = _module.D7.create_DC15(_1265_v48);
        let _1267_v50;
        _1267_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(193),p0);
        let _1268_v51;
        _1268_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1266_v49,_1267_v50);
        let _1269_v52;
        _1269_v52 = new _dafny.CodePoint('m'.codePointAt(0));
        _1264_v47 = (_1264_v47).update(_module.__default.safeModuloInt(new BigNumber(((((_1268_v51).contains(_1266_v49)) ? ((_1268_v51).get(_1266_v49)) : (_1267_v50))).length), p1), _dafny.Seq.update(p2, _module.__default.safeIndex((_this).f19, new BigNumber((p2).length)), _1269_v52));
        if ((_this).f18) {
          r0 = (_this).f18;
          (globalState).f6 = _module.__default.safeDivisionInt(p0, new BigNumber(20));
          let _1270_v53;
          let _nw237 = new _module.C2();
          _nw237.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f18)).length), p1, _this.f17);
          _1270_v53 = _nw237;
          let _1271_v54;
          _1271_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1270_v53,(_this).f18);
          let _1272_v55;
          _1272_v55 = _dafny.Set.fromElements(new BigNumber((_1271_v54).length), p1);
          let _1273_v56;
          _1273_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1272_v55);
          let _1274_v57;
          _1274_v57 = _dafny.Seq.of(_1273_v56, _1273_v56);
          let _1275_v58;
          _1275_v58 = _dafny.MultiSet.fromElements((_this).f19);
          let _1276_v59;
          _1276_v59 = _dafny.Map.Empty.slice().updateUnsafe((_1274_v57)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(p0, new BigNumber((_1275_v58).cardinality()))).length), new BigNumber((_1274_v57).length))],(_this).f19);
          _1276_v59 = (_1276_v59).update(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1272_v55), (_dafny.ZERO).minus(((_this).f19).minus(p1)));
          let _1277_v60;
          _1277_v60 = _module.D10.create_DC28((_this).f19);
          let _1278_v61;
          _1278_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
          let _1279_v62;
          _1279_v62 = _module.D15.create_DC42(_1272_v55);
          _1277_v60 = _module.__default.fm37(_1278_v61, _1279_v62, globalState);
          (globalState).f1 = (((_this).f18) ? ((_this).f18) : (true));
        } else {
          let _1280_v63;
          let _nw238 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1280_v63 = _nw238;
          let _1281_v64;
          _1281_v64 = _dafny.Set.fromElements(_1269_v52);
          let _1282_v65;
          _1282_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1264_v47,_1281_v64);
          let _1283_v66;
          _1283_v66 = _dafny.Set.fromElements(_1281_v64, (((_1282_v65).contains(_1264_v47)) ? ((_1282_v65).get(_1264_v47)) : (_1281_v64)), _dafny.Set.fromElements(_1269_v52));
          let _index218 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1280_v63).length));
          (_1280_v63)[_index218] = new BigNumber((_1283_v66).length);
          let _1284_v67;
          let _nw239 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1284_v67 = _nw239;
          let _1285_v68;
          let _nw240 = Array((new BigNumber(3)).toNumber()).fill(false);
          _1285_v68 = _nw240;
          let _index219 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1285_v68).length));
          (_1285_v68)[_index219] = (_this).f18;
          let _index220 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1280_v63).length));
          let _index221 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1285_v68).length));
          let _rhs164 = new BigNumber(444);
          let _rhs165 = _1284_v67;
          let _rhs166 = new BigNumber(231);
          let _rhs167 = new BigNumber(49);
          let _rhs168 = (_this).f18;
          let _lhs142 = _1280_v63;
          let _lhs143 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1280_v63).length));
          let _lhs144 = globalState;
          let _lhs145 = globalState;
          let _lhs146 = _1285_v68;
          let _lhs147 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_1285_v68).length));
          _lhs142[_lhs143] = _rhs164;
          _1284_v67 = _rhs165;
          _lhs144.f6 = _rhs166;
          _lhs145.f6 = _rhs167;
          _lhs146[_lhs147] = _rhs168;
          let _1286_v69;
          _1286_v69 = _dafny.MultiSet.fromElements((_1285_v68)[_module.__default.safeIndex(new BigNumber(389), new BigNumber((_1285_v68).length))], (_1285_v68)[_module.__default.safeIndex(new BigNumber(389), new BigNumber((_1285_v68).length))]);
          let _1287_v70;
          _1287_v70 = _dafny.Seq.of((_1285_v68)[_module.__default.safeIndex(new BigNumber(389), new BigNumber((_1285_v68).length))], (_this).f18);
          (globalState).f1 = ((_1286_v69).Union(_1286_v69)).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.update(_1287_v70, _module.__default.safeIndex((_this).f19, new BigNumber((_1287_v70).length)), (_this).f18), _1287_v70)));
          _1267_v50 = (_1267_v50).update(new BigNumber(592), p1);
          (globalState).f6 = new BigNumber((p2).length);
          let _1288_v71;
          _1288_v71 = _module.D11.create_DC31(_1287_v70, (_this).f19, p1, new BigNumber(300), (_1280_v63)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1280_v63).length))]);
          let _1289_v72;
          _1289_v72 = _dafny.Seq.of((_this).f19, (_1280_v63)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1280_v63).length))], new BigNumber((_dafny.Seq.UnicodeFromString("unna")).length), _module.__default.safeModuloInt(p1, (_1288_v71).dtor_cf58));
          let _rhs169 = _dafny.Seq.of(new BigNumber(172), (_this).f19, new BigNumber(334), new BigNumber((p3).length), p0);
          let _rhs170 = (_dafny.ZERO).minus((_this).f19);
          let _lhs148 = globalState;
          _1289_v72 = _rhs169;
          _lhs148.f6 = _rhs170;
        }
        _1269_v52 = (((_this).f18) ? (_1269_v52) : (_1269_v52));
        let _1290_v73;
        let _nw241 = new _module.C1();
        _nw241.__ctor(p0);
        _1290_v73 = _nw241;
        let _1291_v74;
        _1291_v74 = _module.D0.create_DC1(new BigNumber(886));
        let _pat_let_tv30 = p1;
        _1291_v74 = function (_pat_let55_0) {
          return function (_1292_dt__update__tmp_h0) {
            return function (_pat_let56_0) {
              return function (_1293_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_1293_dt__update_hcf1_h0);
              }(_pat_let56_0);
            }(_pat_let_tv30);
          }(_pat_let55_0);
        }(_1291_v74);
      }
      r0 = (_this).f18;
      r1 = _dafny.Seq.Concat(p3, p2);
      return [r0, r1];
    }
    m13(globalState) {
      let _this = this;
      let _1294_v0;
      let _nw242 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
      _1294_v0 = _nw242;
      let _1295_v1;
      let _nw243 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1295_v1 = _nw243;
      let _1296_v2;
      _1296_v2 = _dafny.Seq.of(_1295_v1);
      let _index222 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1294_v0).length));
      (_1294_v0)[_index222] = _dafny.Seq.Concat(_1296_v2, _dafny.Seq.of(_1295_v1, _1295_v1, _1295_v1));
      let _index223 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1294_v0).length));
      (_1294_v0)[_index223] = _dafny.Seq.update(_1296_v2, _module.__default.safeIndex(_module.__default.fm1((_this).f18, globalState), new BigNumber((_1296_v2).length)), _1295_v1);
      let _1297_v3;
      _1297_v3 = _dafny.MultiSet.fromElements((_this).f18);
      (globalState).f6 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_this).f19, new BigNumber((_1297_v3).cardinality())), (((_1297_v3).contains((_this).f18)) ? ((_1297_v3).get((_this).f18)) : ((_this).f19)));
      let _1298_v4;
      _1298_v4 = _dafny.Seq.of((_this).f18);
      let _1299_v5;
      let _nw244 = Array((new BigNumber(10)).toNumber());
      _nw244[(_dafny.ZERO).toNumber()] = (_this).f18;
      _nw244[(_dafny.ONE).toNumber()] = (_this).f18;
      _nw244[(new BigNumber(2)).toNumber()] = (_this).f18;
      _nw244[(new BigNumber(3)).toNumber()] = (_this).f18;
      _nw244[(new BigNumber(4)).toNumber()] = (((_this).f18) ? ((_this).f18) : ((_this).f18));
      _nw244[(new BigNumber(5)).toNumber()] = (_this).f18;
      _nw244[(new BigNumber(6)).toNumber()] = ((_this).f19).isLessThan(new BigNumber(731));
      _nw244[(new BigNumber(7)).toNumber()] = (_1297_v3).IsProperSubsetOf((_dafny.MultiSet.FromArray(_1298_v4)).update((_this).f18, _module.__default.abs((_this).f19)));
      _nw244[(new BigNumber(8)).toNumber()] = (_this).f18;
      _nw244[(new BigNumber(9)).toNumber()] = (_this).f18;
      _1299_v5 = _nw244;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1299_v5).length))) {
        let _1300_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1300_i0)) && ((_1300_i0).isLessThan(new BigNumber((_1299_v5).length))))) {
          (_1299_v5)[(_1300_i0)] = (_this).f18;
        }
      }
      let _index224 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length));
      (_1299_v5)[_index224] = (_this).f18;
      let _1301_v6;
      let _nw245 = new _module.C1();
      _nw245.__ctor((_this).f19);
      _1301_v6 = _nw245;
      let _1302_v7;
      _1302_v7 = _dafny.Set.fromElements(_1301_v6, _1301_v6);
      let _1303_v8;
      let _nw246 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _1303_v8 = _nw246;
      let _1304_v9;
      _1304_v9 = _module.D3.create_DC6(_1303_v8);
      let _1305_v10;
      _1305_v10 = _dafny.Seq.of(_1304_v9);
      let _1306_v11;
      _1306_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1305_v10);
      let _1307_v12;
      _1307_v12 = _module.D11.create_DC30((((_1306_v11).contains((_this).f18)) ? ((_1306_v11).get((_this).f18)) : (_1305_v10)));
      let _1308_v13;
      _1308_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1302_v7,_1307_v12);
      let _1309_v14;
      let _nw247 = Array((new BigNumber(3)).toNumber());
      _nw247[(_dafny.ZERO).toNumber()] = _1303_v8;
      _nw247[(_dafny.ONE).toNumber()] = _1303_v8;
      _nw247[(new BigNumber(2)).toNumber()] = _1303_v8;
      _1309_v14 = _nw247;
      let _index225 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_1309_v14).length));
      (_1309_v14)[_index225] = _1303_v8;
      let _1310_v15;
      _1310_v15 = _dafny.Map.Empty.slice().updateUnsafe((_1301_v6).f24,!((_this).f18));
      let _1311_v16;
      _1311_v16 = _dafny.Seq.of(new BigNumber((_1310_v15).length));
      let _1312_v17;
      _1312_v17 = _dafny.Map.Empty.slice().updateUnsafe(!(!_dafny.Seq.contains(_1311_v16, (_1301_v6).f24)),new BigNumber(891));
      let _index226 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length));
      let _index227 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_1309_v14).length));
      let _rhs171 = ((_this).f18) === (((false) ? ((_this).f18) : ((_this).f18)));
      let _rhs172 = _1308_v13;
      let _rhs173 = (((_1312_v17).contains(true)) ? ((_1312_v17).get(true)) : ((_dafny.ZERO).minus((_this).f19)));
      let _rhs174 = _1303_v8;
      let _lhs149 = _1299_v5;
      let _lhs150 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length));
      let _lhs151 = globalState;
      let _lhs152 = _1309_v14;
      let _lhs153 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_1309_v14).length));
      _lhs149[_lhs150] = _rhs171;
      _1308_v13 = _rhs172;
      _lhs151.f6 = _rhs173;
      _lhs152[_lhs153] = _rhs174;
      (globalState).f1 = (((_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))]) ? (true) : ((_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))]));
      let _1313_v18;
      _1313_v18 = _dafny.Set.fromElements((_this).f18);
      let _1314_i1;
      _1314_i1 = _dafny.ZERO;
      L9: {
        while (!(!((_dafny.Set.fromElements((_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))], (_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))], (_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))], (_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))], (_this).f18)).IsDisjointFrom(_1313_v18))) || ((_this).f18)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1314_i1)) {
              break L9;
            }
            _1314_i1 = (_1314_i1).plus(_dafny.ONE);
            let _1315_v21;
            _1315_v21 = _dafny.Set.fromElements(new BigNumber((function () {
              let _coll49 = new _dafny.Map();
              for (const _compr_49 of (function () {
                let _coll50 = new _dafny.Set();
                for (const _compr_50 of (_1311_v16).Elements) {
                  let _1316_v20 = _compr_50;
                  if (_dafny.Seq.contains(_1311_v16, _1316_v20)) {
                    _coll50.add((_1316_v20).minus(new BigNumber(466)));
                  }
                }
                return _coll50;
              }()).Elements) {
                let _1317_v19 = _compr_49;
                if ((function () {
                  let _coll51 = new _dafny.Set();
                  for (const _compr_51 of (_1311_v16).Elements) {
                    let _1318_v20 = _compr_51;
                    if (_dafny.Seq.contains(_1311_v16, _1318_v20)) {
                      _coll51.add((_1318_v20).minus(new BigNumber(466)));
                    }
                  }
                  return _coll51;
                }()).contains(_1317_v19)) {
                  _coll49.push([_module.__default.safeDivisionInt(_1317_v19, new BigNumber(240)),(_1301_v6).f24]);
                }
              }
              return _coll49;
            }()).length));
            let _1319_v22;
            _1319_v22 = _dafny.Seq.of(_1315_v21, _1315_v21);
            if (!((_1315_v21).IsProperSubsetOf((_1319_v22)[_module.__default.safeIndex((_1301_v6).f24, new BigNumber((_1319_v22).length))]))) {
              let _index228 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length));
              (_1299_v5)[_index228] = true;
              _1313_v18 = _dafny.Set.fromElements((_1313_v18).IsSubsetOf(_dafny.Set.fromElements(!((_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))]))));
              _1310_v15 = (_1310_v15).update((_1301_v6).f24, false);
              let _1320_v24;
              _1320_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f18, (_this).f18),new BigNumber(72));
              (globalState).f1 = !((function () {
                let _coll52 = new _dafny.Map();
                for (const _compr_52 of (_dafny.Map.Empty.slice().updateUnsafe(_1298_v4,(_dafny.ZERO).minus((_1301_v6).f24))).Keys.Elements) {
                  let _1321_v23 = _compr_52;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_1298_v4,(_dafny.ZERO).minus((_1301_v6).f24))).contains(_1321_v23)) {
                    _coll52.push([_1321_v23,(_1301_v6).f24]);
                  }
                }
                return _coll52;
              }()).update(_1298_v4, (_1301_v6).f24)).equals(_1320_v24);
              let _1322_v25;
              _1322_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))]);
              let _1323_v26;
              _1323_v26 = _dafny.MultiSet.fromElements(_1313_v18);
              _1322_v25 = (_1322_v25).update((_this).f18, (_1323_v26).IsProperSubsetOf(_1323_v26));
            } else {
              let _1324_v27;
              _1324_v27 = _dafny.Seq.of(_1299_v5);
              let _1325_v28;
              let _nw248 = Array((new BigNumber(13)).toNumber());
              _nw248[(_dafny.ZERO).toNumber()] = _1299_v5;
              _nw248[(_dafny.ONE).toNumber()] = _1299_v5;
              _nw248[(new BigNumber(2)).toNumber()] = _1299_v5;
              _nw248[(new BigNumber(3)).toNumber()] = (((_this).f18) ? (_1299_v5) : (_1299_v5));
              _nw248[(new BigNumber(4)).toNumber()] = _1299_v5;
              _nw248[(new BigNumber(5)).toNumber()] = _1299_v5;
              _nw248[(new BigNumber(6)).toNumber()] = _1299_v5;
              _nw248[(new BigNumber(7)).toNumber()] = ((false) ? (_1299_v5) : (_1299_v5));
              _nw248[(new BigNumber(8)).toNumber()] = _1299_v5;
              _nw248[(new BigNumber(9)).toNumber()] = (_1324_v27)[_module.__default.safeIndex((_this).f19, new BigNumber((_1324_v27).length))];
              _nw248[(new BigNumber(10)).toNumber()] = _1299_v5;
              _nw248[(new BigNumber(11)).toNumber()] = _1299_v5;
              _nw248[(new BigNumber(12)).toNumber()] = _1299_v5;
              _1325_v28 = _nw248;
              _1325_v28 = _1325_v28;
              let _1326_v29;
              _1326_v29 = _dafny.Map.Empty.slice().updateUnsafe(!(((_dafny.ZERO).minus((_this).f19)).isLessThan((_1301_v6).f24)),true);
              _1326_v29 = (_1326_v29).update((_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))], (_this).f18);
              (globalState).f14 = _1299_v5;
              let _index229 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_1303_v8).length));
              (_1303_v8)[_index229] = _module.__default.fm1(_module.__default.fm2(globalState), globalState);
              let _1327_v30;
              _1327_v30 = _dafny.MultiSet.fromElements(((_this).f19).multipliedBy(new BigNumber(153)), ((_1301_v6).f24).multipliedBy((_1301_v6).f24));
              let _1328_v31;
              _1328_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1297_v3).Intersect(_1297_v3)).cardinality()),_module.__default.fm1((_this).f18, globalState));
              let _1329_v32;
              _1329_v32 = new _dafny.CodePoint('y'.codePointAt(0));
              let _1330_v33;
              _1330_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(310),_dafny.Map.Empty.slice().updateUnsafe((_1301_v6).f24,(_this).f19));
              let _1331_v34;
              _1331_v34 = _dafny.Seq.UnicodeFromString("rboenvi");
              let _index230 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_1303_v8).length));
              let _rhs175 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1329_v32, _1329_v32, _1329_v32)).length));
              let _rhs176 = _1327_v30;
              let _rhs177 = (((false) ? ((((_1330_v33).contains((_1301_v6).f24)) ? ((_1330_v33).get((_1301_v6).f24)) : (_1328_v31))) : (_module.__default.fm33((_1301_v6).f24, (_1301_v6).f24, (_this).f18, globalState)))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1301_v6).f24,new BigNumber((_1331_v34).length)));
              let _rhs178 = _1299_v5;
              let _lhs154 = _1303_v8;
              let _lhs155 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_1303_v8).length));
              let _lhs156 = globalState;
              _lhs154[_lhs155] = _rhs175;
              _1327_v30 = _rhs176;
              _1328_v31 = _rhs177;
              _lhs156.f14 = _rhs178;
              _1326_v29 = (_1326_v29).update((_this).f18, (_this).f18);
            }
            let _1332_v36;
            _1332_v36 = _dafny.Map.Empty.slice().updateUnsafe(function () {
              let _coll53 = new _dafny.Map();
              for (const _compr_53 of _dafny.IntegerRange(new BigNumber(305), new BigNumber(589))) {
                let _1333_v35 = _compr_53;
                if (((new BigNumber(305)).isLessThanOrEqualTo(_1333_v35)) && ((_1333_v35).isLessThan(new BigNumber(589)))) {
                  _coll53.push([(_1333_v35).multipliedBy((_1301_v6).f24),false]);
                }
              }
              return _coll53;
            }(),true);
            let _index231 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length));
            (_1299_v5)[_index231] = (((_1301_v6).fm24((_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))], _1332_v36, globalState)) ? ((_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))]) : ((_this).f18));
            let _1334_v37;
            _1334_v37 = _dafny.Map.Empty.slice().updateUnsafe((_1301_v6).f24,(_1301_v6).f24);
            let _rhs179 = (_1334_v37).Merge(_1334_v37);
            let _rhs180 = (_this).f18;
            let _rhs181 = (_1301_v6).f24;
            let _lhs157 = globalState;
            let _lhs158 = globalState;
            _1334_v37 = _rhs179;
            _lhs157.f1 = _rhs180;
            _lhs158.f6 = _rhs181;
            let _source9 = _module.D13.create_DC37(true, true);
            if (_source9.is_DC37) {
              let _1335___mcc_h0 = (_source9).cf65;
              let _1336___mcc_h1 = (_source9).cf66;
              let _1337_cf66 = _1336___mcc_h1;
              let _1338_cf65 = _1335___mcc_h0;
              let _1339_v38;
              _1339_v38 = new _dafny.CodePoint('y'.codePointAt(0));
              let _1340_v39;
              _1340_v39 = _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)), _1339_v38);
              let _1341_v40;
              _1341_v40 = _dafny.Map.Empty.slice().updateUnsafe((_1301_v6).f24,new _dafny.CodePoint('m'.codePointAt(0)));
              let _1342_v41;
              _1342_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1340_v39).length),_1340_v39);
              let _nw249 = Array((new BigNumber(11)).toNumber());
              _nw249[(_dafny.ZERO).toNumber()] = (_this).f19;
              _nw249[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt((_1301_v6).f24, new BigNumber((_1340_v39).length));
              _nw249[(new BigNumber(2)).toNumber()] = new BigNumber((_1341_v40).length);
              _nw249[(new BigNumber(3)).toNumber()] = (_1301_v6).f24;
              _nw249[(new BigNumber(4)).toNumber()] = (_1301_v6).f24;
              _nw249[(new BigNumber(5)).toNumber()] = (_1301_v6).f24;
              _nw249[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(765)), function (_1343_i2) {
                return (_this).f19;
              }), _1311_v16)).length);
              _nw249[(new BigNumber(7)).toNumber()] = (_1301_v6).f24;
              _nw249[(new BigNumber(8)).toNumber()] = new BigNumber(-42);
              _nw249[(new BigNumber(9)).toNumber()] = (_1301_v6).f24;
              _nw249[(new BigNumber(10)).toNumber()] = new BigNumber(((_1342_v41).Merge(_module.__default.fm38((_this).f19, (_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))], true, (_this).f18, globalState))).length);
              (globalState).f9 = _nw249;
              (globalState).f1 = !((_this).f18);
              (globalState).f6 = (_dafny.ZERO).minus((_1301_v6).f24);
              let _1344_v42;
              _1344_v42 = _module.D5.create_DC12((_this).f19, (_this).f18, _1339_v38);
              let _1345_v43;
              _1345_v43 = _module.D5.create_DC13(_1344_v42);
              _1345_v43 = _1345_v43;
            } else {
              let _1346___mcc_h2 = (_source9).cf64;
              let _1347_cf64 = _1346___mcc_h2;
              let _1348_v44;
              _1348_v44 = _dafny.Seq.of(_1305_v10, _dafny.Seq.update(_1305_v10, _module.__default.safeIndex((_this).f19, new BigNumber((_1305_v10).length)), _1304_v9));
              let _1349_v45;
              let _nw250 = new _module.C0();
              _nw250.__ctor(_dafny.Seq.update((_1348_v44)[_module.__default.safeIndex((_1301_v6).f24, new BigNumber((_1348_v44).length))], _module.__default.safeIndex((_1301_v6).f24, new BigNumber(((_1348_v44)[_module.__default.safeIndex((_1301_v6).f24, new BigNumber((_1348_v44).length))]).length)), _1304_v9), false);
              _1349_v45 = _nw250;
              let _1350_v46;
              _1350_v46 = _dafny.Map.Empty.slice().updateUnsafe((_1349_v45).f23,(_1349_v45).f23);
              _1312_v17 = (_1312_v17).update((_1299_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_1299_v5).length))], (_dafny.ZERO).minus(new BigNumber((_1350_v46).length)));
              let _1351_v47;
              let _nw251 = new _module.C2();
              _nw251.__ctor((_1301_v6).f24, (((_1349_v45).f23) ? (new BigNumber(266)) : ((_1301_v6).f24)), _this.f17);
              _1351_v47 = _nw251;
              _1350_v46 = _1350_v46;
            }
          }
        }
      }
      return;
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

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.D0.create_DC2(_module.D0.create_DC1(new BigNumber((_dafny.Seq.of(new BigNumber(946), new BigNumber(-169), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(true), false, false)).length),true)).length), new BigNumber(586), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(458))).cardinality()))).length)));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _1352_v0;
      _1352_v0 = _dafny.Seq.UnicodeFromString("jyvixh");
      _1352_v0 = _1352_v0;
      let _1353_v1;
      let _init35 = function (_1354_i0) {
        return (_1354_i0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(!(true),true),new BigNumber(783))).length));
      };
      let _nw252 = Array((new BigNumber(13)).toNumber());
      for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw252.length); _i0_35++) {
        _nw252[_i0_35] = _init35(new BigNumber(_i0_35));
      }
      _1353_v1 = _nw252;
      let _1355_v2;
      _1355_v2 = new BigNumber(814);
      let _index232 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_1353_v1).length));
      (_1353_v1)[_index232] = (_dafny.ZERO).minus(_1355_v2);
      let _1356_v3;
      let _nw253 = Array((new BigNumber(2)).toNumber()).fill(false);
      _1356_v3 = _nw253;
      let _1357_v4;
      _1357_v4 = new _dafny.CodePoint('l'.codePointAt(0));
      let _1358_v5;
      _1358_v5 = _dafny.MultiSet.fromElements(_1357_v4);
      let _index233 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_1356_v3).length));
      (_1356_v3)[_index233] = !(_1358_v5).contains(_1357_v4);
      let _1359_v6;
      _1359_v6 = true;
      let _index234 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_1353_v1).length));
      let _index235 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_1356_v3).length));
      let _rhs182 = _module.__default.safeModuloInt(new BigNumber(337), (_1355_v2).minus(_1355_v2));
      let _rhs183 = _module.__default.fm1(_1359_v6, globalState);
      let _rhs184 = _1359_v6;
      let _rhs185 = _module.__default.safeDivisionInt(_1355_v2, (_1355_v2).plus(new BigNumber(-621)));
      let _rhs186 = ((_1359_v6) ? (_1352_v0) : (_1352_v0));
      let _lhs159 = globalState;
      let _lhs160 = _1353_v1;
      let _lhs161 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_1353_v1).length));
      let _lhs162 = _1356_v3;
      let _lhs163 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_1356_v3).length));
      let _lhs164 = globalState;
      _lhs159.f6 = _rhs182;
      _lhs160[_lhs161] = _rhs183;
      _lhs162[_lhs163] = _rhs184;
      _lhs164.f6 = _rhs185;
      _1352_v0 = _rhs186;
      let _1360_v7;
      _1360_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(569),(_1356_v3)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_1356_v3).length))]);
      let _1361_v8;
      _1361_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1360_v7,_1355_v2);
      let _1362_v9;
      let _nw254 = new _module.C3();
      _nw254.__ctor(false, _1355_v2, _1361_v8);
      _1362_v9 = _nw254;
      let _index236 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_1353_v1).length));
      (_1353_v1)[_index236] = (_1362_v9).f19;
      let _1363_v10;
      _1363_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1360_v7).length),_1355_v2);
      let _1364_v11;
      _1364_v11 = _module.D12.create_DC34(_1363_v10);
      let _pat_let_tv31 = _1362_v9;
      let _pat_let_tv32 = _1362_v9;
      let _pat_let_tv33 = _1353_v1;
      let _pat_let_tv34 = _1353_v1;
      let _index237 = _module.__default.safeIndex(new BigNumber(822), new BigNumber((_1353_v1).length));
      (_1353_v1)[_index237] = function (_source10) {
        if (_source10.is_DC35) {
          return _module.__default.safeModuloInt((_pat_let_tv31).f19, (_pat_let_tv32).f19);
        } else {
          let _1365___mcc_h0 = (_source10).cf63;
          let _1366_cf63 = _1365___mcc_h0;
          return (_pat_let_tv34)[_module.__default.safeIndex(new BigNumber(822), new BigNumber((_pat_let_tv33).length))];
        }
      }(_1364_v11);
      _1359_v6 = !(false);
      r0 = (_1353_v1)[_module.__default.safeIndex(new BigNumber(822), new BigNumber((_1353_v1).length))];
      let _1367_v12;
      _1367_v12 = _dafny.Seq.of((_1356_v3)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_1356_v3).length))], _module.__default.fm2(globalState), (_1356_v3)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_1356_v3).length))], _1359_v6, (_1362_v9).f18);
      r1 = !(((_1359_v6) ? (false) : (!((_1355_v2).isEqualTo(new BigNumber((_1367_v12).length))))));
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let _1368_v0;
      let _nw255 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1368_v0 = _nw255;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1368_v0).length))) {
        let _1369_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1369_i0)) && ((_1369_i0).isLessThan(new BigNumber((_1368_v0).length))))) {
          (_1368_v0)[(_1369_i0)] = ((false) ? (new _dafny.CodePoint('u'.codePointAt(0))) : (((false) ? (new _dafny.CodePoint('n'.codePointAt(0))) : ((_module.D5.create_DC12(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(968))).cardinality()), p0, new _dafny.CodePoint('r'.codePointAt(0)))).dtor_cf17))));
        }
      }
      let _1370_v1;
      _1370_v1 = new BigNumber(-497);
      let _1371_v2;
      _1371_v2 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(p0, p0),_1370_v1);
      _1371_v2 = _1371_v2;
      let _1372_v3;
      let _init36 = ((_1373_p0) => function (_1374_i1) {
        return _dafny.MultiSet.fromElements(_1373_p0, _1373_p0, _1373_p0, true);
      })(p0);
      let _nw256 = Array((new BigNumber(17)).toNumber());
      for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw256.length); _i0_36++) {
        _nw256[_i0_36] = _init36(new BigNumber(_i0_36));
      }
      _1372_v3 = _nw256;
      _1372_v3 = _1372_v3;
      if (p0) {
        (globalState).f6 = _module.__default.safeDivisionInt(_1370_v1, _1370_v1);
        let _1375_v4;
        let _nw257 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _1375_v4 = _nw257;
        let _index238 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length));
        (_1375_v4)[_index238] = new BigNumber(-199);
        let _index239 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length));
        (_1375_v4)[_index239] = _1370_v1;
        let _index240 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length));
        (_1375_v4)[_index240] = (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))];
        let _1376_v5;
        _1376_v5 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1377_v6;
        _1377_v6 = _dafny.MultiSet.fromElements(_1376_v5, _1376_v5, _1376_v5, new _dafny.CodePoint('d'.codePointAt(0)));
        let _1378_v7;
        _1378_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))]);
        let _1379_v9;
        _1379_v9 = _dafny.MultiSet.fromElements(_1370_v1, _1370_v1);
        let _1380_v10;
        _1380_v10 = _module.D8.create_DC20(new BigNumber(65), _1370_v1, (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))]);
        let _1381_v11;
        _1381_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1380_v10).dtor_cf31,(_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))]);
        let _1382_v12;
        _1382_v12 = _dafny.Seq.UnicodeFromString("b");
        let _1383_v13;
        _1383_v13 = _module.D5.create_DC12((_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))], p0, _1376_v5);
        let _1384_v14;
        _1384_v14 = _dafny.Set.fromElements(_1383_v13);
        let _1385_v15;
        _1385_v15 = _dafny.Seq.of(p0);
        let _1386_v16;
        _1386_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1385_v15);
        let _1387_v17;
        _1387_v17 = _dafny.MultiSet.fromElements(p0);
        let _1388_v18;
        _1388_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _nw258 = Array((new BigNumber(26)).toNumber());
        _nw258[(_dafny.ZERO).toNumber()] = (((_1377_v6).contains(_1376_v5)) ? ((_1377_v6).get(_1376_v5)) : (_module.__default.fm1(p0, globalState)));
        _nw258[(_dafny.ONE).toNumber()] = _1370_v1;
        _nw258[(new BigNumber(2)).toNumber()] = ((_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))]).multipliedBy(_1370_v1);
        _nw258[(new BigNumber(3)).toNumber()] = new BigNumber((_1378_v7).length);
        _nw258[(new BigNumber(4)).toNumber()] = (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))];
        _nw258[(new BigNumber(5)).toNumber()] = new BigNumber((((function () {
          let _coll54 = new _dafny.Map();
          for (const _compr_54 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(936)), ((_1389_p0) => function (_1390_i2) {
            return new BigNumber((_dafny.Set.fromElements(_1389_p0)).length);
          })(p0))).Elements) {
            let _1391_v8 = _compr_54;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(936)), ((_1392_p0) => function (_1390_i2) {
              return new BigNumber((_dafny.Set.fromElements(_1392_p0)).length);
            })(p0)), _1391_v8)) {
              _coll54.push([_module.__default.safeDivisionInt(_1391_v8, (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))]),(_dafny.ZERO).minus(_1370_v1)]);
            }
          }
          return _coll54;
        }()).update(new BigNumber((_1379_v9).cardinality()), _1370_v1)).Merge(_1381_v11)).length);
        _nw258[(new BigNumber(6)).toNumber()] = (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))];
        _nw258[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(_1370_v1, (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))]);
        _nw258[(new BigNumber(8)).toNumber()] = (((_1381_v11).contains(_module.__default.fm1(p0, globalState))) ? ((_1381_v11).get(_module.__default.fm1(p0, globalState))) : (new BigNumber((_1382_v12).length)));
        _nw258[(new BigNumber(9)).toNumber()] = new BigNumber(((_1378_v7).Merge(_1378_v7)).length);
        _nw258[(new BigNumber(10)).toNumber()] = (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))];
        _nw258[(new BigNumber(11)).toNumber()] = new BigNumber(((_1384_v14).Intersect(_1384_v14)).length);
        _nw258[(new BigNumber(12)).toNumber()] = (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))];
        _nw258[(new BigNumber(13)).toNumber()] = _1370_v1;
        _nw258[(new BigNumber(14)).toNumber()] = _1370_v1;
        _nw258[(new BigNumber(15)).toNumber()] = (_1370_v1).multipliedBy((_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))]);
        _nw258[(new BigNumber(16)).toNumber()] = (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))];
        _nw258[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt(_1370_v1, (_dafny.ZERO).minus(_1370_v1));
        _nw258[(new BigNumber(18)).toNumber()] = ((_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))]).plus(new BigNumber(761));
        _nw258[(new BigNumber(19)).toNumber()] = _module.__default.fm1(p0, globalState);
        _nw258[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1386_v16).length));
        _nw258[(new BigNumber(21)).toNumber()] = (new BigNumber((_1387_v17).cardinality())).multipliedBy((_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))]);
        _nw258[(new BigNumber(22)).toNumber()] = (_1375_v4)[_module.__default.safeIndex(new BigNumber(526), new BigNumber((_1375_v4).length))];
        _nw258[(new BigNumber(23)).toNumber()] = new BigNumber((_1381_v11).length);
        _nw258[(new BigNumber(24)).toNumber()] = _1370_v1;
        _nw258[(new BigNumber(25)).toNumber()] = new BigNumber((_1388_v18).length);
        _1375_v4 = _nw258;
        (globalState).f6 = ((!(p0)) ? (_1370_v1) : (_1370_v1));
      } else {
        _1370_v1 = _1370_v1;
        let _1393_v20;
        _1393_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1370_v1,p0);
        let _1394_v21;
        _1394_v21 = _dafny.MultiSet.fromElements(new BigNumber((function () {
          let _coll55 = new _dafny.Map();
          for (const _compr_55 of (_1393_v20).Keys.Elements) {
            let _1395_v19 = _compr_55;
            if ((_1393_v20).contains(_1395_v19)) {
              _coll55.push([(_1395_v19).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(485)), function (_1396_i3) {
                return new _dafny.CodePoint('e'.codePointAt(0));
              })).length)),_dafny.Seq.UnicodeFromString("fpfcroym")]);
            }
          }
          return _coll55;
        }()).length), _1370_v1);
        let _1397_v22;
        _1397_v22 = _dafny.Seq.of(_1394_v21);
        let _1398_v23;
        _1398_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1397_v22);
        let _1399_v24;
        _1399_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        _1398_v23 = (_1398_v23).update((((_1399_v24).contains(p0)) ? ((_1399_v24).get(p0)) : (p0)), _1397_v22);
        let _1400_v25;
        _1400_v25 = _dafny.MultiSet.fromElements(p0);
        (globalState).f6 = ((((_1400_v25).contains(p0)) ? ((_1400_v25).get(p0)) : (_1370_v1))).plus((_dafny.ZERO).minus(_1370_v1));
        let _1401_v26;
        _1401_v26 = _dafny.Seq.UnicodeFromString("n");
        let _1402_v27;
        _1402_v27 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1403_v28;
        _1403_v28 = _dafny.Seq.of(_1401_v26);
        let _1404_v29;
        let _nw259 = Array((new BigNumber(20)).toNumber());
        _nw259[(_dafny.ZERO).toNumber()] = _1401_v26;
        _nw259[(_dafny.ONE).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_module.__default.fm21(globalState), _module.__default.safeIndex(_1370_v1, new BigNumber((_module.__default.fm21(globalState)).length)), _1402_v27);
        _nw259[(new BigNumber(3)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(258)), ((_1405_v27) => function (_1406_i4) {
          return _1405_v27;
        })(_1402_v27));
        _nw259[(new BigNumber(5)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(6)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("n");
        _nw259[(new BigNumber(8)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(9)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(10)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(11)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("rtuxv");
        _nw259[(new BigNumber(13)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(406)), ((_1407_v27) => function (_1408_i5) {
          return _1407_v27;
        })(_1402_v27));
        _nw259[(new BigNumber(15)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(16)).toNumber()] = (_1403_v28)[_module.__default.safeIndex(_1370_v1, new BigNumber((_1403_v28).length))];
        _nw259[(new BigNumber(17)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(18)).toNumber()] = _1401_v26;
        _nw259[(new BigNumber(19)).toNumber()] = _1401_v26;
        _1404_v29 = _nw259;
        let _1409_v30;
        let _nw260 = Array((new BigNumber(4)).toNumber());
        _nw260[(_dafny.ZERO).toNumber()] = _1404_v29;
        _nw260[(_dafny.ONE).toNumber()] = _1404_v29;
        _nw260[(new BigNumber(2)).toNumber()] = _1404_v29;
        _nw260[(new BigNumber(3)).toNumber()] = _1404_v29;
        _1409_v30 = _nw260;
        let _index241 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_1409_v30).length));
        (_1409_v30)[_index241] = _1404_v29;
        let _index242 = _module.__default.safeIndex(new BigNumber(282), new BigNumber((_1409_v30).length));
        let _init37 = ((_1410_v26) => function (_1411_i6) {
          return _1410_v26;
        })(_1401_v26);
        let _nw261 = Array((new BigNumber(9)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw261.length); _i0_37++) {
          _nw261[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        (_1409_v30)[_index242] = _nw261;
        (globalState).f1 = p0;
      }
      let _1412_i7;
      _1412_i7 = _dafny.ZERO;
      L10: {
        while (p0) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1412_i7)) {
              break L10;
            }
            _1412_i7 = (_1412_i7).plus(_dafny.ONE);
            let _1413_v31;
            _1413_v31 = _dafny.Set.fromElements(_1370_v1, _1370_v1, _1370_v1);
            let _1414_v32;
            let _nw262 = new _module.C1();
            _nw262.__ctor(new BigNumber((_1413_v31).length));
            _1414_v32 = _nw262;
            let _1415_v33;
            let _nw263 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
            _1415_v33 = _nw263;
            let _1416_v34;
            _1416_v34 = _module.D3.create_DC6(_1415_v33);
            let _1417_v35;
            _1417_v35 = _dafny.Seq.of(_1416_v34, _1416_v34, _1416_v34, _1416_v34, _1416_v34);
            let _1418_v36;
            let _nw264 = new _module.C0();
            _nw264.__ctor(_1417_v35, p0);
            _1418_v36 = _nw264;
            if (true) {
              let _1419_v37;
              _1419_v37 = new _dafny.CodePoint('i'.codePointAt(0));
              _1419_v37 = _1419_v37;
              let _1420_v38;
              _1420_v38 = _dafny.Seq.UnicodeFromString("kgkj");
              let _1421_v39;
              _1421_v39 = _dafny.MultiSet.fromElements(_1413_v31);
              let _1422_v40;
              _1422_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1420_v38).length),new BigNumber((_1421_v39).cardinality()));
              let _1423_v41;
              _1423_v41 = _module.D12.create_DC34(_1422_v40);
              let _1424_v43;
              _1424_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1423_v41,function () {
                let _coll56 = new _dafny.Set();
                for (const _compr_56 of _dafny.IntegerRange(new BigNumber(470), new BigNumber(885))) {
                  let _1425_v42 = _compr_56;
                  if (((new BigNumber(470)).isLessThanOrEqualTo(_1425_v42)) && ((_1425_v42).isLessThan(new BigNumber(885)))) {
                    _coll56.add((_1425_v42).multipliedBy((_1414_v32).f24));
                  }
                }
                return _coll56;
              }());
              let _1426_v44;
              let _nw265 = Array((new BigNumber(4)).toNumber());
              _nw265[(_dafny.ZERO).toNumber()] = _1424_v43;
              _nw265[(_dafny.ONE).toNumber()] = _1424_v43;
              _nw265[(new BigNumber(2)).toNumber()] = _1424_v43;
              _nw265[(new BigNumber(3)).toNumber()] = _1424_v43;
              _1426_v44 = _nw265;
              let _index243 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1426_v44).length));
              (_1426_v44)[_index243] = _1424_v43;
              let _1427_v45;
              _1427_v45 = _dafny.MultiSet.fromElements((_1414_v32).f24, _1370_v1);
              let _1428_v46;
              _1428_v46 = _dafny.Seq.of(_1427_v45);
              let _1429_v47;
              _1429_v47 = _dafny.Seq.of(new BigNumber((_1428_v46).length));
              let _1430_v50;
              _1430_v50 = _dafny.Set.fromElements(_1423_v41);
              let _index244 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1426_v44).length));
              let _rhs187 = ((function () {
                let _coll57 = new _dafny.Map();
                for (const _compr_57 of (function () {
                  let _coll58 = new _dafny.Map();
                  for (const _compr_58 of (_1430_v50).Elements) {
                    let _1431_v49 = _compr_58;
                    if ((_1430_v50).contains(_1431_v49)) {
                      _coll58.push([_1431_v49,_1419_v37]);
                    }
                  }
                  return _coll58;
                }()).Keys.Elements) {
                  let _1432_v48 = _compr_57;
                  if ((function () {
                    let _coll59 = new _dafny.Map();
                    for (const _compr_59 of (_1430_v50).Elements) {
                      let _1431_v49 = _compr_59;
                      if ((_1430_v50).contains(_1431_v49)) {
                        _coll59.push([_1431_v49,_1419_v37]);
                      }
                    }
                    return _coll59;
                  }()).contains(_1432_v48)) {
                    _coll57.push([_1432_v48,_1413_v31]);
                  }
                }
                return _coll57;
              }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1423_v41,_1413_v31))).Merge(_1424_v43);
              let _rhs188 = _1429_v47;
              let _lhs165 = _1426_v44;
              let _lhs166 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_1426_v44).length));
              _lhs165[_lhs166] = _rhs187;
              _1429_v47 = _rhs188;
              (globalState).f1 = true;
              _1370_v1 = ((_1414_v32).f24).minus((_1414_v32).f24);
              (globalState).f1 = _module.__default.fm2(globalState);
            } else {
              let _1433_v51;
              _1433_v51 = new _dafny.CodePoint('k'.codePointAt(0));
              let _1434_v52;
              _1434_v52 = _dafny.Seq.UnicodeFromString("okij");
              let _1435_v53;
              _1435_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1433_v51,_1434_v52);
              let _1436_v54;
              _1436_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1433_v51,(_1414_v32).f24);
              let _index245 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1415_v33).length));
              (_1415_v33)[_index245] = _module.__default.safeDivisionInt(new BigNumber((_1435_v53).length), new BigNumber((_module.__default.fm28(new BigNumber((_1436_v54).length), (_1418_v36).f23, true, globalState)).length));
              let _index246 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_1415_v33).length));
              (_1415_v33)[_index246] = new BigNumber((_1434_v52).length);
              (globalState).f1 = (_1418_v36).f23;
              let _rhs189 = (_1418_v36).f23;
              let _rhs190 = _1434_v52;
              let _rhs191 = p0;
              let _lhs167 = globalState;
              let _lhs168 = globalState;
              _lhs167.f1 = _rhs189;
              _1434_v52 = _rhs190;
              _lhs168.f1 = _rhs191;
              let _1437_v55;
              let _init38 = ((_1438_v51) => function (_1439_i8) {
                return _module.D1.create_DC4((_module.D1.create_DC4(_1438_v51)).dtor_cf4);
              })(_1433_v51);
              let _nw266 = Array((new BigNumber(12)).toNumber());
              for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw266.length); _i0_38++) {
                _nw266[_i0_38] = _init38(new BigNumber(_i0_38));
              }
              _1437_v55 = _nw266;
              let _1440_v56;
              _1440_v56 = _dafny.Seq.of((_1418_v36).f23);
              let _1441_v57;
              _1441_v57 = _dafny.MultiSet.fromElements((_1440_v56)[_module.__default.safeIndex((_1415_v33)[_module.__default.safeIndex(new BigNumber(301), new BigNumber((_1415_v33).length))], new BigNumber((_1440_v56).length))], !(p0), p0);
              let _rhs192 = _1437_v55;
              let _rhs193 = _1441_v57;
              let _rhs194 = !(((false) ? (_module.__default.fm2(globalState)) : (p0)));
              let _rhs195 = false;
              let _lhs169 = globalState;
              let _lhs170 = globalState;
              _1437_v55 = _rhs192;
              _1441_v57 = _rhs193;
              _lhs169.f1 = _rhs194;
              _lhs170.f1 = _rhs195;
              _1434_v52 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1434_v52, _dafny.Seq.UnicodeFromString("siy")), _1434_v52);
            }
            let _1442_v58;
            _1442_v58 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1414_v32);
            let _1443_v59;
            let _nw267 = Array((new BigNumber(26)).toNumber());
            _nw267[(_dafny.ZERO).toNumber()] = _1414_v32;
            _nw267[(_dafny.ONE).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(2)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(3)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(4)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(5)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(6)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(7)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(8)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(9)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(10)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(11)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(12)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(13)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(14)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(15)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(16)).toNumber()] = (((_1442_v58).contains((_1418_v36).f23)) ? ((_1442_v58).get((_1418_v36).f23)) : (_1414_v32));
            _nw267[(new BigNumber(17)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(18)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(19)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(20)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(21)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(22)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(23)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(24)).toNumber()] = _1414_v32;
            _nw267[(new BigNumber(25)).toNumber()] = _1414_v32;
            _1443_v59 = _nw267;
            let _1444_v60;
            _1444_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1443_v59,_1370_v1);
            (globalState).f6 = (((_1444_v60).contains(_1443_v59)) ? ((_1444_v60).get(_1443_v59)) : (_1370_v1));
          }
        }
      }
      if (_module.__default.fm2(globalState)) {
        let _1445_v61;
        _1445_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1370_v1,_module.__default.fm21(globalState));
        let _1446_v62;
        _1446_v62 = _dafny.Seq.UnicodeFromString("cxdqb");
        let _1447_v63;
        let _nw268 = Array((new BigNumber(8)).toNumber());
        _nw268[(_dafny.ZERO).toNumber()] = _1370_v1;
        _nw268[(_dafny.ONE).toNumber()] = _1370_v1;
        _nw268[(new BigNumber(2)).toNumber()] = _1370_v1;
        _nw268[(new BigNumber(3)).toNumber()] = _1370_v1;
        _nw268[(new BigNumber(4)).toNumber()] = _1370_v1;
        _nw268[(new BigNumber(5)).toNumber()] = _1370_v1;
        _nw268[(new BigNumber(6)).toNumber()] = _1370_v1;
        _nw268[(new BigNumber(7)).toNumber()] = new BigNumber(((((_1445_v61).contains(_1370_v1)) ? ((_1445_v61).get(_1370_v1)) : (_1446_v62))).length);
        _1447_v63 = _nw268;
        let _1448_v64;
        _1448_v64 = _dafny.Seq.of(_1447_v63);
        _1448_v64 = _dafny.Seq.of(_1447_v63, _1447_v63);
        (globalState).f1 = p0;
        let _1449_v65;
        _1449_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1370_v1,new BigNumber((_module.__default.fm21(globalState)).length));
        let _1450_v66;
        _1450_v66 = _dafny.Seq.of(p0);
        let _1451_v67;
        _1451_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(226),p0);
        let _1452_v68;
        _1452_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1451_v67,_1370_v1);
        let _1453_v69;
        _1453_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe(_1451_v67,new BigNumber((_1450_v66).length)));
        let _1454_v71;
        _1454_v71 = _dafny.MultiSet.fromElements(_1451_v67, _1451_v67);
        let _1455_v72;
        let _nw269 = new _module.C3();
        _nw269.__ctor(p0, (_dafny.ZERO).minus((((_1449_v65).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1450_v66).length))))) ? ((_1449_v65).get((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1450_v66).length))))) : (_1370_v1))), (_1452_v68).Merge((((_1453_v69).contains(p0)) ? ((_1453_v69).get(p0)) : (function () {
          let _coll60 = new _dafny.Map();
          for (const _compr_60 of (_1454_v71).Elements) {
            let _1456_v70 = _compr_60;
            if ((_1454_v71).contains(_1456_v70)) {
              _coll60.push([_1456_v70,_1370_v1]);
            }
          }
          return _coll60;
        }()))));
        _1455_v72 = _nw269;
        (globalState).f1 = (_1455_v72).f18;
        (globalState).f1 = p0;
      } else {
        (globalState).f6 = _1370_v1;
        (globalState).f1 = p0;
        (globalState).f1 = ((p0) ? (_module.__default.fm2(globalState)) : (p0));
        let _1457_v73;
        _1457_v73 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1458_v74;
        _1458_v74 = _dafny.Seq.UnicodeFromString("yehtu");
        let _1459_v75;
        _1459_v75 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(166));
        _1457_v73 = (_1458_v74)[_module.__default.safeIndex((new BigNumber((_1459_v75).length)).minus(_1370_v1), new BigNumber((_1458_v74).length))];
        let _1460_v76;
        _1460_v76 = _dafny.MultiSet.fromElements(p0, p0);
        (globalState).f6 = (_1370_v1).minus((((_1460_v76).contains(p0)) ? ((_1460_v76).get(p0)) : (_1370_v1)));
      }
      return;
    }
    m10(globalState) {
      let _this = this;
      let _1461_v0;
      _1461_v0 = true;
      let _1462_v1;
      _1462_v1 = _dafny.Map.Empty.slice().updateUnsafe((_1461_v0) === (_1461_v0),_module.__default.fm1(_1461_v0, globalState));
      _1462_v1 = _1462_v1;
      let _1463_v2;
      _1463_v2 = _dafny.Seq.UnicodeFromString("as");
      let _1464_v3;
      _1464_v3 = new _dafny.CodePoint('m'.codePointAt(0));
      let _1465_i0;
      _1465_i0 = _dafny.ZERO;
      L11: {
        while (!((_1461_v0) || (_dafny.Seq.IsPrefixOf(_dafny.Seq.update(_1463_v2, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), ((_1497_v0) => function (_1498_i1) {
          return new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(false), _dafny.Set.fromElements(_1497_v0))).length);
        })(_1461_v0))).length), new BigNumber((_1463_v2).length)), _1464_v3), _1463_v2)))) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1465_i0)) {
              break L11;
            }
            _1465_i0 = (_1465_i0).plus(_dafny.ONE);
            let _1466_v4;
            _1466_v4 = new BigNumber(-666);
            let _1467_v5;
            _1467_v5 = _dafny.Seq.of(_1461_v0, _1461_v0, _1461_v0, false);
            let _1468_v6;
            _1468_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(695),new BigNumber(16));
            let _1469_v7;
            _1469_v7 = _module.D12.create_DC34(_1468_v6);
            let _1470_v8;
            let _nw270 = Array((new BigNumber(16)).toNumber());
            _nw270[(_dafny.ZERO).toNumber()] = ((true) ? (_1466_v4) : (_1466_v4));
            _nw270[(_dafny.ONE).toNumber()] = _1466_v4;
            _nw270[(new BigNumber(2)).toNumber()] = _1466_v4;
            _nw270[(new BigNumber(3)).toNumber()] = _module.__default.fm1(_1461_v0, globalState);
            _nw270[(new BigNumber(4)).toNumber()] = _1466_v4;
            _nw270[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), ((_1471_v0) => function (_1472_i2) {
              return new BigNumber((_dafny.Seq.of(_1471_v0)).length);
            })(_1461_v0))).length), _1466_v4, _1466_v4, _1466_v4)).length);
            _nw270[(new BigNumber(6)).toNumber()] = _1466_v4;
            _nw270[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_1466_v4);
            _nw270[(new BigNumber(8)).toNumber()] = new BigNumber((_1467_v5).length);
            _nw270[(new BigNumber(9)).toNumber()] = new BigNumber((_1463_v2).length);
            _nw270[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("btflmnhpm")).length);
            _nw270[(new BigNumber(11)).toNumber()] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1466_v4,_1469_v7)).length)).multipliedBy(_1466_v4);
            _nw270[(new BigNumber(12)).toNumber()] = _module.__default.fm1(true, globalState);
            _nw270[(new BigNumber(13)).toNumber()] = (_module.__default.fm1(_1461_v0, globalState)).minus(_1466_v4);
            _nw270[(new BigNumber(14)).toNumber()] = _1466_v4;
            _nw270[(new BigNumber(15)).toNumber()] = _1466_v4;
            _1470_v8 = _nw270;
            let _index247 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length));
            (_1470_v8)[_index247] = _1466_v4;
            let _1473_v9;
            _1473_v9 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(192)), ((_1474_v3) => function (_1475_i3) {
              return _1474_v3;
            })(_1464_v3)), _1463_v2);
            let _index248 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length));
            (_1470_v8)[_index248] = (_dafny.ZERO).minus((((_1473_v9).contains(_1463_v2)) ? ((_1473_v9).get(_1463_v2)) : (_1466_v4)));
            let _1476_v10;
            _1476_v10 = _dafny.Set.fromElements(!(!(_module.__default.fm2(globalState))));
            _1476_v10 = _1476_v10;
            if (true) {
              let _1477_v11;
              let _nw271 = Array((new BigNumber(24)).toNumber()).fill(false);
              _1477_v11 = _nw271;
              let _index249 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_1477_v11).length));
              (_1477_v11)[_index249] = _1461_v0;
              let _index250 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_1477_v11).length));
              (_1477_v11)[_index250] = _1461_v0;
              let _1478_v12;
              _1478_v12 = _dafny.Set.fromElements((_1470_v8)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length))], (_1470_v8)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length))]);
              let _1479_v13;
              _1479_v13 = _module.D15.create_DC42(_1478_v12);
              let _index251 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length));
              let _rhs196 = _1466_v4;
              let _rhs197 = _1479_v13;
              let _lhs171 = _1470_v8;
              let _lhs172 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length));
              _lhs171[_lhs172] = _rhs196;
              _1479_v13 = _rhs197;
              let _1480_v14;
              _1480_v14 = _dafny.MultiSet.fromElements(_module.__default.fm2(globalState));
              let _1481_v15;
              let _nw272 = Array((new BigNumber(23)).toNumber()).fill([]);
              _1481_v15 = _nw272;
              let _1482_v16;
              _1482_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1480_v14).IsSubsetOf(_1480_v14),_1481_v15);
              _1482_v16 = (_1482_v16).update(_1461_v0, _1481_v15);
              (globalState).f6 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1466_v4, (((_1477_v11)[_module.__default.safeIndex(new BigNumber(818), new BigNumber((_1477_v11).length))]) ? ((_1470_v8)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length))]) : (_module.__default.fm1(_1461_v0, globalState))))));
              let _1483_v17;
              let _nw273 = new _module.C1();
              _nw273.__ctor((_1470_v8)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length))]);
              _1483_v17 = _nw273;
              let _1484_v18;
              _1484_v18 = _module.D7.create_DC15(_1483_v17);
              let _1485_v19;
              _1485_v19 = _dafny.Seq.of(_1483_v17, _1483_v17);
              let _1486_v20;
              let _nw274 = Array((new BigNumber(16)).toNumber());
              _nw274[(_dafny.ZERO).toNumber()] = _1484_v18;
              _nw274[(_dafny.ONE).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(2)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(3)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(4)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(5)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(6)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(7)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(8)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(9)).toNumber()] = _module.D7.create_DC15((_1485_v19)[_module.__default.safeIndex((_1470_v8)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length))], new BigNumber((_1485_v19).length))]);
              _nw274[(new BigNumber(10)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(11)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(12)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(13)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(14)).toNumber()] = _1484_v18;
              _nw274[(new BigNumber(15)).toNumber()] = _1484_v18;
              _1486_v20 = _nw274;
              let _1487_v21;
              _1487_v21 = _dafny.MultiSet.fromElements(_1486_v20, _1486_v20, _1486_v20);
              let _1488_v22;
              let _nw275 = new _module.C2();
              _nw275.__ctor((_1470_v8)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length))], _1466_v4, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_1470_v8)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length))],_1461_v0),_1466_v4));
              _1488_v22 = _nw275;
              let _1489_v23;
              _1489_v23 = _dafny.Seq.of(_1488_v22);
              let _1490_v24;
              let _nw276 = Array((new BigNumber(20)).toNumber());
              _nw276[(_dafny.ZERO).toNumber()] = _1487_v21;
              _nw276[(_dafny.ONE).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(2)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_1486_v20, _1486_v20, _1486_v20);
              _nw276[(new BigNumber(4)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(5)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(6)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(7)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(8)).toNumber()] = ((_1487_v21).update(_1486_v20, _module.__default.abs(_1466_v4))).Difference(_1487_v21);
              _nw276[(new BigNumber(9)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(10)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(11)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(12)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(13)).toNumber()] = (_1487_v21).Intersect((_1487_v21).update(_1486_v20, _module.__default.abs((_1470_v8)[_module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length))])));
              _nw276[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(_1486_v20, _1486_v20, _1486_v20, _1486_v20, _1486_v20);
              _nw276[(new BigNumber(15)).toNumber()] = (_1487_v21).Intersect(_1487_v21);
              _nw276[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements(_1486_v20);
              _nw276[(new BigNumber(17)).toNumber()] = (_1487_v21).update(_1486_v20, _module.__default.abs(new BigNumber((_1489_v23).length)));
              _nw276[(new BigNumber(18)).toNumber()] = _1487_v21;
              _nw276[(new BigNumber(19)).toNumber()] = _1487_v21;
              _1490_v24 = _nw276;
              let _index252 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_1490_v24).length));
              (_1490_v24)[_index252] = _1487_v21;
              let _1491_v25;
              _1491_v25 = _module.D16.create_DC44(_1480_v14);
              let _pat_let_tv35 = _1480_v14;
              let _index253 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_1490_v24).length));
              let _rhs198 = _1488_v22.f21;
              let _rhs199 = ((_1487_v21).Intersect(_dafny.MultiSet.fromElements(_1486_v20, _1486_v20))).Intersect(_1487_v21);
              let _rhs200 = _module.__default.fm2(globalState);
              let _rhs201 = !(((function (_pat_let57_0) {
                return function (_1492_dt__update__tmp_h0) {
                  return function (_pat_let58_0) {
                    return function (_1493_dt__update_hcf77_h0) {
                      return _module.D16.create_DC44(_1493_dt__update_hcf77_h0);
                    }(_pat_let58_0);
                  }(_pat_let_tv35);
                }(_pat_let57_0);
              }(_1491_v25)).dtor_cf77).Intersect(_1480_v14)).contains(false);
              let _lhs173 = globalState;
              let _lhs174 = _1490_v24;
              let _lhs175 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_1490_v24).length));
              let _lhs176 = globalState;
              let _lhs177 = globalState;
              _lhs173.f6 = _rhs198;
              _lhs174[_lhs175] = _rhs199;
              _lhs176.f1 = _rhs200;
              _lhs177.f1 = _rhs201;
            } else {
              _1461_v0 = _1461_v0;
              let _1494_v26;
              let _nw277 = new _module.C1();
              _nw277.__ctor((_1466_v4).multipliedBy(_1466_v4));
              _1494_v26 = _nw277;
              let _index254 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length));
              (_1470_v8)[_index254] = (_1466_v4).plus(_1466_v4);
              let _1495_v27;
              _1495_v27 = _dafny.Map.Empty.slice().updateUnsafe((_1494_v26).fm24(!(_1461_v0), _module.__default.fm39(_1461_v0, globalState), globalState),false);
              _1495_v27 = (_1495_v27).update(!(_1461_v0), !(true) || (_1461_v0));
              let _index255 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1470_v8).length));
              (_1470_v8)[_index255] = new BigNumber((_1467_v5).length);
            }
            let _1496_v28;
            _1496_v28 = _dafny.Set.fromElements(_1470_v8, _1470_v8, _1470_v8);
            _1496_v28 = _1496_v28;
          }
        }
      }
      let _1499_v29;
      _1499_v29 = new BigNumber(538);
      let _1500_v30;
      let _nw278 = Array((new BigNumber(5)).toNumber());
      _1500_v30 = _nw278;
      let _1501_v31;
      _1501_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1464_v3,_1499_v29);
      let _1502_v32;
      _1502_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1500_v30,_1501_v31);
      let _hi9 = _module.__default.safeDivisionInt(new BigNumber((_1502_v32).length), _1499_v29);
      for (let _1503_i4 = _1499_v29; _1503_i4.isLessThan(_hi9); _1503_i4 = _1503_i4.plus(_dafny.ONE)) {
        let _1504_v33;
        let _nw279 = Array((new BigNumber(22)).toNumber()).fill(false);
        _1504_v33 = _nw279;
        (globalState).f14 = _1504_v33;
        let _1505_v34;
        _1505_v34 = _dafny.Map.Empty.slice().updateUnsafe(((_1461_v0) ? (_1461_v0) : (_1461_v0)),_1464_v3);
        let _1506_v35;
        _1506_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1461_v0,_module.__default.fm29(_1461_v0, _1499_v29, _1503_i4, true, globalState));
        _1505_v34 = (_1506_v35);
        let _1507_v36;
        let _nw280 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _1507_v36 = _nw280;
        let _index256 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1507_v36).length));
        (_1507_v36)[_index256] = _1499_v29;
        let _index257 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1507_v36).length));
        (_1507_v36)[_index257] = _1499_v29;
        _1463_v2 = _dafny.Seq.update(_module.__default.fm21(globalState), _module.__default.safeIndex((_1507_v36)[_module.__default.safeIndex(new BigNumber(106), new BigNumber((_1507_v36).length))], new BigNumber((_module.__default.fm21(globalState)).length)), _1464_v3);
      }
      (globalState).f6 = _1499_v29;
      if (_1461_v0) {
        let _1508_v37;
        _1508_v37 = _dafny.Set.fromElements(_1499_v29);
        let _1509_v38;
        _1509_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1508_v37,_1463_v2);
        let _1510_v39;
        _1510_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1499_v29,(((_1462_v1).contains(!(_1461_v0))) ? ((_1462_v1).get(!(_1461_v0))) : (new BigNumber(((((_1509_v38).contains(_1508_v37)) ? ((_1509_v38).get(_1508_v37)) : (_1463_v2))).length))));
        let _1511_v40;
        let _init39 = ((_1512_v0, _1513_v2) => function (_1514_i5) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_1512_v0,_1513_v2)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1512_v0,_1513_v2));
        })(_1461_v0, _1463_v2);
        let _nw281 = Array((new BigNumber(2)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw281.length); _i0_39++) {
          _nw281[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1511_v40 = _nw281;
        let _1515_v41;
        _1515_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1461_v0,_1463_v2);
        let _index258 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_1511_v40).length));
        (_1511_v40)[_index258] = _1515_v41;
        let _1516_v43;
        _1516_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1499_v29,_1461_v0);
        let _1517_v44;
        _1517_v44 = _dafny.Seq.of(_1510_v39, function () {
          let _coll61 = new _dafny.Map();
          for (const _compr_61 of (_1516_v43).Keys.Elements) {
            let _1518_v42 = _compr_61;
            if ((_1516_v43).contains(_1518_v42)) {
              _coll61.push([(_1518_v42).multipliedBy(_1499_v29),_1499_v29]);
            }
          }
          return _coll61;
        }(), _1510_v39, _1510_v39);
        let _index259 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_1511_v40).length));
        let _rhs202 = (_1517_v44)[_module.__default.safeIndex((_dafny.ZERO).minus(_1499_v29), new BigNumber((_1517_v44).length))];
        let _rhs203 = _1515_v41;
        let _rhs204 = _module.__default.fm29(_1461_v0, _1499_v29, _module.__default.fm1(_1461_v0, globalState), false, globalState);
        let _lhs178 = _1511_v40;
        let _lhs179 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_1511_v40).length));
        _1510_v39 = _rhs202;
        _lhs178[_lhs179] = _rhs203;
        _1464_v3 = _rhs204;
        _1461_v0 = _1461_v0;
        (globalState).f6 = new BigNumber(788);
        if (_1461_v0) {
          (globalState).f6 = ((_1499_v29).minus(new BigNumber(-168))).plus(_1499_v29);
          (globalState).f1 = true;
          let _1519_v45;
          let _nw282 = new _module.C1();
          _nw282.__ctor(new BigNumber((_1463_v2).length));
          _1519_v45 = _nw282;
          let _1520_v46;
          _1520_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1461_v0,_1461_v0);
          (globalState).f6 = (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber(((_1520_v46).update(_1461_v0, _1461_v0)).length)).minus(new BigNumber(-717))));
          let _1521_v47;
          let _init40 = ((_1522_v29, _1523_v0, _1524_v2, _1525_v45, _1526_v37) => function (_1527_i6) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D9.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(_1522_v29,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("oorciw")).length),_dafny.Set.fromElements(_1522_v29))), _dafny.Seq.of(_1523_v0)), _module.D9.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1524_v2).length),_dafny.Map.Empty.slice().updateUnsafe((_1525_v45).f24,_1526_v37)), _dafny.Seq.of(_1523_v0, _1523_v0))),new BigNumber(86))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(802)), ((_1528_v29, _1529_v45, _1530_v37, _1531_v0) => function (_1532_i7) {
              return _module.D9.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(_1528_v29,_dafny.Map.Empty.slice().updateUnsafe((_1529_v45).f24,_1530_v37)), _dafny.Seq.of(_1531_v0, _1531_v0));
            })(_1522_v29, _1525_v45, _1526_v37, _1523_v0)),new BigNumber(726)));
          })(_1499_v29, _1461_v0, _1463_v2, _1519_v45, _1508_v37);
          let _nw283 = Array((new BigNumber(28)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw283.length); _i0_40++) {
            _nw283[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1521_v47 = _nw283;
          let _1533_v49;
          _1533_v49 = function () {
            let _coll62 = new _dafny.Map();
            for (const _compr_62 of _dafny.IntegerRange(new BigNumber(-144), new BigNumber(515))) {
              let _1534_v48 = _compr_62;
              if (((new BigNumber(-144)).isLessThanOrEqualTo(_1534_v48)) && ((_1534_v48).isLessThan(new BigNumber(515)))) {
                _coll62.push([_module.__default.safeDivisionInt(_1534_v48, (_1519_v45).f24),_1508_v37]);
              }
            }
            return _coll62;
          }();
          let _1535_v50;
          _1535_v50 = _dafny.Map.Empty.slice().updateUnsafe((_1519_v45).f24,_1533_v49);
          let _1536_v51;
          _1536_v51 = _dafny.Seq.of(_1461_v0, _1461_v0, _1461_v0, _1461_v0);
          let _1537_v52;
          _1537_v52 = _module.D9.create_DC24(_1535_v50, _1536_v51);
          let _1538_v53;
          _1538_v53 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1537_v52),(_dafny.ZERO).minus(_1499_v29));
          let _index260 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1521_v47).length));
          (_1521_v47)[_index260] = _1538_v53;
          let _1539_v54;
          _1539_v54 = _dafny.Seq.of(_1462_v1);
          let _1540_v55;
          _1540_v55 = _dafny.MultiSet.fromElements(_1461_v0, false);
          let _1541_v56;
          _1541_v56 = _dafny.MultiSet.fromElements(new BigNumber(713), (((_1540_v55).contains(_1461_v0)) ? ((_1540_v55).get(_1461_v0)) : (_1499_v29)));
          let _index261 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1521_v47).length));
          let _rhs205 = _1499_v29;
          let _rhs206 = _module.__default.fm40(_1499_v29, globalState);
          let _rhs207 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_1539_v54, _module.__default.safeIndex(new BigNumber(((_1511_v40)[_module.__default.safeIndex(new BigNumber(264), new BigNumber((_1511_v40).length))]).length), new BigNumber((_1539_v54).length)), _1462_v1), _1539_v54), _1539_v54);
          let _rhs208 = new BigNumber(-133);
          let _rhs209 = ((_module.__default.fm1(_1461_v0, globalState)).minus((((_1541_v56).contains((_1519_v45).f24)) ? ((_1541_v56).get((_1519_v45).f24)) : ((_1519_v45).f24)))).minus(_module.__default.safeModuloInt((_1519_v45).f24, _1499_v29));
          let _lhs180 = _1521_v47;
          let _lhs181 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1521_v47).length));
          let _lhs182 = globalState;
          _1499_v29 = _rhs205;
          _lhs180[_lhs181] = _rhs206;
          _1539_v54 = _rhs207;
          _lhs182.f6 = _rhs208;
          _1499_v29 = _rhs209;
        } else {
          let _1542_v57;
          _1542_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length),_1463_v2);
          let _1543_v58;
          let _nw284 = Array((new BigNumber(23)).toNumber());
          _nw284[(_dafny.ZERO).toNumber()] = _1463_v2;
          _nw284[(_dafny.ONE).toNumber()] = _module.__default.fm21(globalState);
          _nw284[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("yfhwoiw");
          _nw284[(new BigNumber(3)).toNumber()] = _1463_v2;
          _nw284[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), function (_1544_i8) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }), _1463_v2);
          _nw284[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("q"), _1463_v2);
          _nw284[(new BigNumber(6)).toNumber()] = _1463_v2;
          _nw284[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("wjvd");
          _nw284[(new BigNumber(8)).toNumber()] = _1463_v2;
          _nw284[(new BigNumber(9)).toNumber()] = _1463_v2;
          _nw284[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1463_v2, _1463_v2);
          _nw284[(new BigNumber(11)).toNumber()] = _1463_v2;
          _nw284[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_1463_v2, _1463_v2);
          _nw284[(new BigNumber(13)).toNumber()] = _1463_v2;
          _nw284[(new BigNumber(14)).toNumber()] = _1463_v2;
          _nw284[(new BigNumber(15)).toNumber()] = (((_1542_v57).contains(_1499_v29)) ? ((_1542_v57).get(_1499_v29)) : (_1463_v2));
          _nw284[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gqj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(882)), ((_1545_v3) => function (_1546_i9) {
            return _1545_v3;
          })(_1464_v3)));
          _nw284[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("gv");
          _nw284[(new BigNumber(18)).toNumber()] = _1463_v2;
          _nw284[(new BigNumber(19)).toNumber()] = _1463_v2;
          _nw284[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("gjbixrfgg");
          _nw284[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("oarbcrux");
          _nw284[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_1463_v2, _1463_v2);
          _1543_v58 = _nw284;
          let _1547_v59;
          _1547_v59 = _dafny.Seq.of(_1463_v2, _1463_v2, _1463_v2, _1463_v2, _1463_v2);
          let _index262 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_1543_v58).length));
          (_1543_v58)[_index262] = (_1547_v59)[_module.__default.safeIndex(_1499_v29, new BigNumber((_1547_v59).length))];
          let _index263 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_1543_v58).length));
          (_1543_v58)[_index263] = _1463_v2;
          let _1548_v60;
          _1548_v60 = _dafny.MultiSet.fromElements(!(_1461_v0) || (_1461_v0), !(!((_1461_v0) || (false))), _1461_v0, (_1499_v29).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("hqtqm")).length)), _1461_v0);
          let _rhs210 = new BigNumber((_1548_v60).cardinality());
          let _rhs211 = !(_module.__default.fm1(_1461_v0, globalState)).isEqualTo(_module.__default.fm1(_1461_v0, globalState));
          let _lhs183 = globalState;
          let _lhs184 = globalState;
          _lhs183.f6 = _rhs210;
          _lhs184.f1 = _rhs211;
          (globalState).f6 = (_1499_v29).multipliedBy(_module.__default.fm1(_1461_v0, globalState));
          let _1549_v61;
          _1549_v61 = _dafny.Map.Empty.slice().updateUnsafe((_1543_v58)[_module.__default.safeIndex(new BigNumber(316), new BigNumber((_1543_v58).length))],_1499_v29);
          let _1550_v62;
          let _init41 = ((_1551_v29) => function (_1552_i10) {
            return _module.__default.safeDivisionInt(_1552_i10, _1551_v29);
          })(_1499_v29);
          let _nw285 = Array((new BigNumber(24)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw285.length); _i0_41++) {
            _nw285[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1550_v62 = _nw285;
          let _1553_v63;
          _1553_v63 = _dafny.Seq.of(_1461_v0, _1461_v0, !(false));
          let _1554_v64;
          _1554_v64 = _module.D12.create_DC34(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1553_v63)).cardinality()),new BigNumber((_1463_v2).length)));
          let _index264 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1550_v62).length));
          (_1550_v62)[_index264] = new BigNumber((_module.__default.fm36(_1464_v3, _1554_v64, globalState)).length);
          let _1555_v65;
          _1555_v65 = _dafny.MultiSet.fromElements(_1499_v29);
          let _1556_v66;
          _1556_v66 = _dafny.Seq.of(_1555_v65);
          let _1557_v67;
          _1557_v67 = _dafny.Seq.of(_1499_v29);
          let _1558_v68;
          _1558_v68 = _dafny.Set.fromElements((_1543_v58)[_module.__default.safeIndex(new BigNumber(316), new BigNumber((_1543_v58).length))], _dafny.Seq.update(_1463_v2, _module.__default.safeIndex(_1499_v29, new BigNumber((_1463_v2).length)), _1464_v3));
          let _index265 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1550_v62).length));
          let _rhs212 = (new BigNumber(957)).minus((_dafny.ZERO).minus(_1499_v29));
          let _rhs213 = (((_1461_v0) ? (_1549_v61) : (_1549_v61))).Merge(_1549_v61);
          let _rhs214 = (_1499_v29).multipliedBy(_1499_v29);
          let _rhs215 = !(!_dafny.areEqual(_dafny.Seq.Concat(_1556_v66, _module.__default.fm41(new BigNumber((_1557_v67).length), globalState)), _dafny.Seq.of(_1555_v65)));
          let _rhs216 = (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("qxbpkigo"))).IsDisjointFrom((_1558_v68).Intersect(_1558_v68));
          let _lhs185 = globalState;
          let _lhs186 = _1550_v62;
          let _lhs187 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1550_v62).length));
          let _lhs188 = globalState;
          let _lhs189 = globalState;
          _lhs185.f6 = _rhs212;
          _1549_v61 = _rhs213;
          _lhs186[_lhs187] = _rhs214;
          _lhs188.f1 = _rhs215;
          _lhs189.f1 = _rhs216;
          let _1559_v69;
          let _nw286 = Array((new BigNumber(6)).toNumber()).fill(false);
          _1559_v69 = _nw286;
          let _index266 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_1559_v69).length));
          (_1559_v69)[_index266] = _1461_v0;
          let _1560_v70;
          _1560_v70 = _dafny.Map.Empty.slice().updateUnsafe((_1550_v62)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1550_v62).length))],_1550_v62);
          let _index267 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_1559_v69).length));
          (_1559_v69)[_index267] = !(new BigNumber((_1560_v70).length)).isEqualTo(_module.__default.safeModuloInt((_dafny.ZERO).minus(_1499_v29), (_1550_v62)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1550_v62).length))]));
        }
        let _1561_v71;
        _1561_v71 = _dafny.Set.fromElements(!(_1461_v0), _1461_v0);
        let _1562_v72;
        _1562_v72 = _module.D8.create_DC21(_1463_v2, _1561_v71);
        let _source11 = _1562_v72;
        if (_source11.is_DC20) {
          let _1563___mcc_h0 = (_source11).cf31;
          let _1564___mcc_h1 = (_source11).cf32;
          let _1565___mcc_h2 = (_source11).cf33;
          let _1566_cf33 = _1565___mcc_h2;
          let _1567_cf32 = _1564___mcc_h1;
          let _1568_cf31 = _1563___mcc_h0;
          (globalState).f1 = !(_1461_v0) || (_1461_v0);
          (globalState).f6 = _module.__default.fm1((_dafny.Set.fromElements(new BigNumber((_1463_v2).length), _1499_v29)).IsSubsetOf(_1508_v37), globalState);
          let _1569_v73;
          _1569_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1461_v0,_1461_v0);
          let _1570_v74;
          _1570_v74 = _dafny.MultiSet.fromElements(new BigNumber((_1569_v73).length), new BigNumber((_dafny.Seq.UnicodeFromString("cnmgv")).length));
          let _1571_v75;
          _1571_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1567_cf32,_module.__default.fm42(_1567_cf32, _1464_v3, _1568_cf31, _1461_v0, globalState));
          let _1572_v76;
          _1572_v76 = _dafny.MultiSet.fromElements(_1570_v74, _1570_v74, (((_1571_v75).contains(_1567_cf32)) ? ((_1571_v75).get(_1567_cf32)) : (_1570_v74)), _1570_v74, _dafny.MultiSet.fromElements(_1566_cf33, _1499_v29));
          let _1573_v77;
          let _nw287 = Array((new BigNumber(15)).toNumber()).fill(false);
          _1573_v77 = _nw287;
          let _1574_v78;
          _1574_v78 = _dafny.MultiSet.fromElements(_1573_v77);
          let _1575_v79;
          _1575_v79 = _dafny.Seq.of(_1499_v29);
          let _1576_v80;
          _1576_v80 = _dafny.Seq.of(_1570_v74, _dafny.MultiSet.fromElements(_1499_v29));
          let _1577_v81;
          let _nw288 = Array((new BigNumber(29)).toNumber());
          _nw288[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_1570_v74);
          _nw288[(_dafny.ONE).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(2)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(3)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_1508_v37).length)), _1570_v74);
          _nw288[(new BigNumber(5)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(6)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(7)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(8)).toNumber()] = (_1572_v76).update((_1570_v74).update((((_1574_v78).contains(_1573_v77)) ? ((_1574_v78).get(_1573_v77)) : (_1499_v29)), _module.__default.abs(_1499_v29)), _module.__default.abs(new BigNumber(-251)));
          _nw288[(new BigNumber(9)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(10)).toNumber()] = _module.__default.fm43(_1461_v0, _module.__default.fm2(globalState), _1461_v0, _1566_cf33, globalState);
          _nw288[(new BigNumber(11)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(12)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(13)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(14)).toNumber()] = ((_1461_v0) ? (_1572_v76) : (_1572_v76));
          _nw288[(new BigNumber(15)).toNumber()] = (_1572_v76).Union(_dafny.MultiSet.fromElements(_1570_v74, _1570_v74));
          _nw288[(new BigNumber(16)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(17)).toNumber()] = _module.__default.fm43(_1461_v0, _1461_v0, !(_1461_v0), new BigNumber(-376), globalState);
          _nw288[(new BigNumber(18)).toNumber()] = (_1572_v76).Intersect(_1572_v76);
          _nw288[(new BigNumber(19)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(20)).toNumber()] = (_1572_v76).Union(_1572_v76);
          _nw288[(new BigNumber(21)).toNumber()] = (_1572_v76).Union(((_1572_v76).update(_dafny.MultiSet.FromArray(_1575_v79), _module.__default.abs(_1499_v29))).update(_1570_v74, _module.__default.abs(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(313)), ((_1578_cf32, _1579_cf33, _1580_v37, _1581_v0) => function (_1582_i11) {
            return _module.D9.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(_1578_cf32,_dafny.Map.Empty.slice().updateUnsafe(_1579_cf33,_1580_v37)), _dafny.Seq.of(_1581_v0, _1581_v0));
          })(_1567_cf32, _1566_cf33, _1508_v37, _1461_v0))).length))));
          _nw288[(new BigNumber(22)).toNumber()] = (_dafny.MultiSet.FromArray(_1576_v80)).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of((_1570_v74).update((((_1462_v1).contains(_1461_v0)) ? ((_1462_v1).get(_1461_v0)) : (_1566_cf33)), _module.__default.abs(_1567_cf32)), _1570_v74)));
          _nw288[(new BigNumber(23)).toNumber()] = (_1572_v76).Difference(_1572_v76);
          _nw288[(new BigNumber(24)).toNumber()] = (_1572_v76).Difference(_1572_v76);
          _nw288[(new BigNumber(25)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(26)).toNumber()] = (_dafny.MultiSet.FromArray(_1576_v80)).Difference(_1572_v76);
          _nw288[(new BigNumber(27)).toNumber()] = _1572_v76;
          _nw288[(new BigNumber(28)).toNumber()] = _dafny.MultiSet.FromArray(_1576_v80);
          _1577_v81 = _nw288;
          let _index268 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_1577_v81).length));
          (_1577_v81)[_index268] = _dafny.MultiSet.fromElements(_1570_v74);
          let _index269 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_1577_v81).length));
          (_1577_v81)[_index269] = _dafny.MultiSet.fromElements((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1567_cf32))).Union(_1570_v74));
          let _1583_v82;
          let _nw289 = new _module.C1();
          _nw289.__ctor(_1568_cf31);
          _1583_v82 = _nw289;
          let _1584_v83;
          _1584_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1566_cf33,_1583_v82);
          _1584_v83 = (_1584_v83).update(_1567_cf32, _1583_v82);
        } else if (_source11.is_DC21) {
          let _1585___mcc_h3 = (_source11).cf34;
          let _1586___mcc_h4 = (_source11).cf35;
          let _1587_cf35 = _1586___mcc_h4;
          let _1588_cf34 = _1585___mcc_h3;
          let _1589_v84;
          let _init42 = ((_1590_v2) => function (_1591_i12) {
            return _1590_v2;
          })(_1463_v2);
          let _nw290 = Array((new BigNumber(13)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw290.length); _i0_42++) {
            _nw290[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1589_v84 = _nw290;
          _1589_v84 = _1589_v84;
          let _1592_v85;
          let _nw291 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _1592_v85 = _nw291;
          let _1593_v86;
          _1593_v86 = _module.D3.create_DC6(_1592_v85);
          let _pat_let_tv36 = _1592_v85;
          let _1594_v87;
          _1594_v87 = _dafny.Seq.of(function (_pat_let59_0) {
            return function (_1595_dt__update__tmp_h1) {
              return function (_pat_let60_0) {
                return function (_1596_dt__update_hcf6_h0) {
                  return _module.D3.create_DC6(_1596_dt__update_hcf6_h0);
                }(_pat_let60_0);
              }(_pat_let_tv36);
            }(_pat_let59_0);
          }(_1593_v86), _1593_v86);
          let _1597_v88;
          let _nw292 = new _module.C0();
          _nw292.__ctor(_1594_v87, _module.__default.fm2(globalState));
          _1597_v88 = _nw292;
          let _1598_v89;
          _1598_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1597_v88,(_1597_v88).f23);
          let _1599_v90;
          _1599_v90 = _dafny.Seq.of(new BigNumber((_1588_cf34).length));
          let _1600_v91;
          _1600_v91 = _dafny.MultiSet.fromElements(new BigNumber((_1599_v90).length));
          let _1601_v92;
          let _nw293 = Array((new BigNumber(7)).toNumber());
          _nw293[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1598_v89).length));
          _nw293[(_dafny.ONE).toNumber()] = new BigNumber((_1600_v91).cardinality());
          _nw293[(new BigNumber(2)).toNumber()] = _1499_v29;
          _nw293[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_1499_v29);
          _nw293[(new BigNumber(4)).toNumber()] = new BigNumber((_1462_v1).length);
          _nw293[(new BigNumber(5)).toNumber()] = new BigNumber((_1588_cf34).length);
          _nw293[(new BigNumber(6)).toNumber()] = _1499_v29;
          _1601_v92 = _nw293;
          let _1602_v93;
          _1602_v93 = _module.D3.create_DC6(_1601_v92);
          let _1603_v94;
          _1603_v94 = _dafny.Seq.of(_1602_v93, _module.D3.create_DC6(_1592_v85), _1602_v93, _1593_v86, _1593_v86);
          let _1604_v95;
          let _nw294 = new _module.C0();
          _nw294.__ctor(_dafny.Seq.Concat(_1594_v87, _dafny.Seq.of(_1593_v86)), _module.__default.fm2(globalState));
          _1604_v95 = _nw294;
          _1604_v95 = _1604_v95;
          (globalState).f6 = _1499_v29;
          (globalState).f6 = (_module.__default.safeModuloInt(_1499_v29, _1499_v29)).plus((_1499_v29).multipliedBy(new BigNumber((_1599_v90).length)));
        } else {
          let _1605___mcc_h5 = (_source11).cf30;
          let _1606_cf30 = _1605___mcc_h5;
          let _1607_v96;
          _1607_v96 = _dafny.Seq.of(_1499_v29, _1499_v29, _1499_v29);
          _1607_v96 = _1607_v96;
          let _1608_v97;
          _1608_v97 = _dafny.Map.Empty.slice().updateUnsafe(_1516_v43,_1499_v29);
          let _1609_v98;
          let _nw295 = new _module.C2();
          _nw295.__ctor((_1499_v29).plus(_1499_v29), _1499_v29, _1608_v97);
          _1609_v98 = _nw295;
          let _1610_v99;
          let _nw296 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1610_v99 = _nw296;
          let _index270 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1610_v99).length));
          (_1610_v99)[_index270] = _1463_v2;
          let _index271 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_1610_v99).length));
          (_1610_v99)[_index271] = _1463_v2;
          let _1611_v100;
          _1611_v100 = _dafny.MultiSet.fromElements(_1461_v0);
          _1611_v100 = _1611_v100;
        }
      } else {
        let _1612_v101;
        let _init43 = ((_1613_v2, _1614_v29, _1615_v3) => function (_1616_i13) {
          return _dafny.areEqual(_dafny.Seq.update(_1613_v2, _module.__default.safeIndex(_1614_v29, new BigNumber((_1613_v2).length)), _1615_v3), _1613_v2);
        })(_1463_v2, _1499_v29, _1464_v3);
        let _nw297 = Array((new BigNumber(8)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw297.length); _i0_43++) {
          _nw297[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1612_v101 = _nw297;
        let _index272 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1612_v101).length));
        (_1612_v101)[_index272] = (_1499_v29).isEqualTo(new BigNumber(103));
        let _1617_v102;
        let _nw298 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _1617_v102 = _nw298;
        let _index273 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1617_v102).length));
        (_1617_v102)[_index273] = _1499_v29;
        let _1618_v103;
        _1618_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1499_v29,_1499_v29);
        let _1619_v104;
        _1619_v104 = _module.D12.create_DC34(_1618_v103);
        let _1620_v105;
        _1620_v105 = _dafny.Seq.of(_1461_v0, _1461_v0, _1461_v0, _1461_v0);
        let _1621_v106;
        _1621_v106 = _dafny.Map.Empty.slice().updateUnsafe(_1461_v0,_1461_v0);
        let _1622_v107;
        let _nw299 = Array((new BigNumber(29)).toNumber());
        _nw299[(_dafny.ZERO).toNumber()] = _1462_v1;
        _nw299[(_dafny.ONE).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(2)).toNumber()] = (_module.__default.fm36(_1464_v3, _1619_v104, globalState)).Merge(_1462_v1);
        _nw299[(new BigNumber(3)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(4)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(5)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(6)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1461_v0,_1499_v29);
        _nw299[(new BigNumber(8)).toNumber()] = (_1462_v1).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1461_v0,_1499_v29));
        _nw299[(new BigNumber(9)).toNumber()] = (_1462_v1).Merge((_1462_v1).update(_1461_v0, new BigNumber((_1463_v2).length)));
        _nw299[(new BigNumber(10)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(11)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(12)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(13)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(14)).toNumber()] = (_1462_v1).Merge(_1462_v1);
        _nw299[(new BigNumber(15)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(16)).toNumber()] = (((_1620_v105)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1620_v105).length))]) ? (_1462_v1) : ((_1462_v1).update(_1461_v0, _1499_v29)));
        _nw299[(new BigNumber(17)).toNumber()] = ((_1462_v1).update(true, _1499_v29)).Merge(_1462_v1);
        _nw299[(new BigNumber(18)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(19)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(20)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(21)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(22)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1461_v0,new BigNumber((_1621_v106).length));
        _nw299[(new BigNumber(24)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(25)).toNumber()] = (_1462_v1).Merge(_module.__default.fm36(_1464_v3, _1619_v104, globalState));
        _nw299[(new BigNumber(26)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(27)).toNumber()] = _1462_v1;
        _nw299[(new BigNumber(28)).toNumber()] = _1462_v1;
        _1622_v107 = _nw299;
        let _1623_v108;
        _1623_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1499_v29,_1462_v1);
        let _index274 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1622_v107).length));
        (_1622_v107)[_index274] = (((_1623_v108).contains(_1499_v29)) ? ((_1623_v108).get(_1499_v29)) : (_1462_v1));
        let _1624_v109;
        _1624_v109 = _dafny.Map.Empty.slice().updateUnsafe(_1612_v101,_1463_v2);
        let _index275 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1612_v101).length));
        let _index276 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1617_v102).length));
        let _index277 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1622_v107).length));
        let _rhs217 = ((_1461_v0) ? (_1461_v0) : ((_module.__default.fm1(_1461_v0, globalState)).isLessThanOrEqualTo(new BigNumber(((((_1624_v109).contains(_1612_v101)) ? ((_1624_v109).get(_1612_v101)) : (_1463_v2))).length))));
        let _rhs218 = _1499_v29;
        let _rhs219 = _1462_v1;
        let _rhs220 = _module.__default.safeDivisionInt(_1499_v29, _1499_v29);
        let _rhs221 = _1461_v0;
        let _lhs190 = _1612_v101;
        let _lhs191 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1612_v101).length));
        let _lhs192 = _1617_v102;
        let _lhs193 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1617_v102).length));
        let _lhs194 = _1622_v107;
        let _lhs195 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_1622_v107).length));
        let _lhs196 = globalState;
        let _lhs197 = globalState;
        _lhs190[_lhs191] = _rhs217;
        _lhs192[_lhs193] = _rhs218;
        _lhs194[_lhs195] = _rhs219;
        _lhs196.f6 = _rhs220;
        _lhs197.f1 = _rhs221;
        let _1625_v110;
        let _nw300 = Array((new BigNumber(19)).toNumber()).fill([]);
        _1625_v110 = _nw300;
        let _1626_v111;
        let _init44 = function (_1627_i14) {
          return _dafny.MultiSet.fromElements(true);
        };
        let _nw301 = Array((new BigNumber(23)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw301.length); _i0_44++) {
          _nw301[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _1626_v111 = _nw301;
        let _1628_v112;
        _1628_v112 = _dafny.Seq.of(_1626_v111);
        let _index278 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1625_v110).length));
        (_1625_v110)[_index278] = (_1628_v112)[_module.__default.safeIndex(_1499_v29, new BigNumber((_1628_v112).length))];
        let _1629_v113;
        _1629_v113 = _dafny.Map.Empty.slice().updateUnsafe((_1617_v102)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_1617_v102).length))],_1461_v0);
        let _index279 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1617_v102).length));
        let _index280 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1625_v110).length));
        let _rhs222 = _module.__default.fm1(false, globalState);
        let _rhs223 = _1461_v0;
        let _rhs224 = ((_1461_v0) ? (new BigNumber(90)) : ((_1617_v102)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((_1617_v102).length))]));
        let _rhs225 = (_1461_v0) && (!(_1629_v113).contains(new BigNumber(432)));
        let _rhs226 = _1626_v111;
        let _lhs198 = globalState;
        let _lhs199 = _1617_v102;
        let _lhs200 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_1617_v102).length));
        let _lhs201 = globalState;
        let _lhs202 = _1625_v110;
        let _lhs203 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_1625_v110).length));
        _lhs198.f6 = _rhs222;
        _1461_v0 = _rhs223;
        _lhs199[_lhs200] = _rhs224;
        _lhs201.f1 = _rhs225;
        _lhs202[_lhs203] = _rhs226;
        _1463_v2 = _1463_v2;
        let _1630_v114;
        _1630_v114 = _dafny.MultiSet.fromElements((_1612_v101)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1612_v101).length))]);
        let _1631_v115;
        _1631_v115 = _dafny.Seq.of(_1629_v113);
        let _1632_v116;
        _1632_v116 = _dafny.Map.Empty.slice().updateUnsafe((((_1631_v115)[_module.__default.safeIndex(_1499_v29, new BigNumber((_1631_v115).length))]).update(_1499_v29, (_1612_v101)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1612_v101).length))])),_1499_v29);
        let _1633_v117;
        _1633_v117 = _1632_v116;
        let _1634_v118;
        let _nw302 = new _module.C2();
        _nw302.__ctor(_1499_v29, (_dafny.ZERO).minus(new BigNumber((((_1630_v114).update((_1612_v101)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1612_v101).length))], _module.__default.abs(new BigNumber((_1629_v113).length)))).Difference(_1630_v114)).cardinality())), ((_1633_v117)).Merge(_1632_v116));
        _1634_v118 = _nw302;
        let _1635_v120;
        let _nw303 = new _module.C3();
        _nw303.__ctor((_1612_v101)[_module.__default.safeIndex(new BigNumber(317), new BigNumber((_1612_v101).length))], (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber(159)).plus((_1634_v118).f20))), function () {
          let _coll63 = new _dafny.Map();
          for (const _compr_63 of (_1632_v116).Keys.Elements) {
            let _1636_v119 = _compr_63;
            if ((_1632_v116).contains(_1636_v119)) {
              _coll63.push([_1636_v119,new BigNumber((_1621_v106).length)]);
            }
          }
          return _coll63;
        }());
        _1635_v120 = _nw303;
      }
      (globalState).f1 = _module.__default.fm2(globalState);
      return;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f16 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f16) {
      let _this = this;
      (_this)._f16 = f16;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.D0.create_DC2(_module.D0.create_DC0(false));
    };
    fm11(p0, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(true)).Union(_dafny.MultiSet.fromElements(true));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _1637_v0;
      _1637_v0 = false;
      let _1638_v1;
      _1638_v1 = _dafny.Seq.of(_1637_v0, _1637_v0);
      let _1639_v2;
      let _nw304 = Array((new BigNumber(12)).toNumber());
      _nw304[(_dafny.ZERO).toNumber()] = _1637_v0;
      _nw304[(_dafny.ONE).toNumber()] = !(_1637_v0);
      _nw304[(new BigNumber(2)).toNumber()] = _1637_v0;
      _nw304[(new BigNumber(3)).toNumber()] = _1637_v0;
      _nw304[(new BigNumber(4)).toNumber()] = false;
      _nw304[(new BigNumber(5)).toNumber()] = _1637_v0;
      _nw304[(new BigNumber(6)).toNumber()] = _1637_v0;
      _nw304[(new BigNumber(7)).toNumber()] = _1637_v0;
      _nw304[(new BigNumber(8)).toNumber()] = _1637_v0;
      _nw304[(new BigNumber(9)).toNumber()] = _1637_v0;
      _nw304[(new BigNumber(10)).toNumber()] = _1637_v0;
      _nw304[(new BigNumber(11)).toNumber()] = _1637_v0;
      _1639_v2 = _nw304;
      let _1640_v3;
      _1640_v3 = _dafny.Seq.of(_1639_v2, _1639_v2);
      let _1641_v4;
      _1641_v4 = _dafny.Seq.of(new BigNumber((_1640_v3).length));
      let _1642_v5;
      let _nw305 = Array((new BigNumber(11)).toNumber());
      _nw305[(_dafny.ZERO).toNumber()] = _1637_v0;
      _nw305[(_dafny.ONE).toNumber()] = _1637_v0;
      _nw305[(new BigNumber(2)).toNumber()] = true;
      _nw305[(new BigNumber(3)).toNumber()] = _1637_v0;
      _nw305[(new BigNumber(4)).toNumber()] = _1637_v0;
      _nw305[(new BigNumber(5)).toNumber()] = (_1638_v1)[_module.__default.safeIndex((_this).f16, new BigNumber((_1638_v1).length))];
      _nw305[(new BigNumber(6)).toNumber()] = ((_dafny.ZERO).minus((_dafny.ZERO).minus((_1641_v4)[_module.__default.safeIndex((_this).f16, new BigNumber((_1641_v4).length))]))).isLessThan((_this).f16);
      _nw305[(new BigNumber(7)).toNumber()] = ((_1637_v0) ? (!(_1637_v0)) : (_1637_v0));
      _nw305[(new BigNumber(8)).toNumber()] = _1637_v0;
      _nw305[(new BigNumber(9)).toNumber()] = _module.__default.fm2(globalState);
      _nw305[(new BigNumber(10)).toNumber()] = _1637_v0;
      _1642_v5 = _nw305;
      let _1643_v6;
      _1643_v6 = new _dafny.CodePoint('n'.codePointAt(0));
      let _1644_v7;
      _1644_v7 = _dafny.Seq.UnicodeFromString("y");
      let _index281 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
      (_1639_v2)[_index281] = _dafny.Seq.contains(_1644_v7, _1643_v6);
      let _index282 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
      (_1639_v2)[_index282] = _1637_v0;
      let _1645_i0;
      _1645_i0 = _dafny.ZERO;
      L12: {
        while (((_1637_v0) ? (false) : ((new BigNumber((_1644_v7).length)).isLessThan((_1641_v4)[_module.__default.safeIndex((_this).f16, new BigNumber((_1641_v4).length))])))) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1645_i0)) {
              break L12;
            }
            _1645_i0 = (_1645_i0).plus(_dafny.ONE);
            let _1646_v8;
            let _nw306 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
            _1646_v8 = _nw306;
            let _1647_v9;
            _1647_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f16, (_this).f16),_1646_v8);
            _1647_v9 = (_1647_v9).update((((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]) ? ((_this).f16) : ((_this).f16)), _1646_v8);
            let _1648_v10;
            _1648_v10 = _dafny.Set.fromElements((_this).f16);
            let _1649_v11;
            let _init45 = ((_1650_v2) => function (_1651_i1) {
              return _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC0((_1650_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1650_v2).length))])));
            })(_1639_v2);
            let _nw307 = Array((new BigNumber(26)).toNumber());
            for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw307.length); _i0_45++) {
              _nw307[_i0_45] = _init45(new BigNumber(_i0_45));
            }
            _1649_v11 = _nw307;
            let _1652_v12;
            let _nw308 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
            _1652_v12 = _nw308;
            _module.__default.m0((((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]) ? (_1648_v10) : (_1648_v10)), _1649_v11, _1652_v12, _1644_v7, globalState);
            let _1653_v13;
            _1653_v13 = _dafny.MultiSet.fromElements(_1637_v0, false);
            let _index283 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_1652_v12).length));
            (_1652_v12)[_index283] = _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm12(_1653_v13, (_this).f16, globalState)).length), (_this).f16);
            let _index284 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_1652_v12).length));
            (_1652_v12)[_index284] = (_this).f16;
            let _index285 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
            (_1639_v2)[_index285] = !(!((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]));
          }
        }
      }
      let _1654_v14;
      let _nw309 = Array((new BigNumber(14)).toNumber());
      _nw309[(_dafny.ZERO).toNumber()] = _1638_v1;
      _nw309[(_dafny.ONE).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(2)).toNumber()] = _module.__default.fm13(_1637_v0, _1643_v6, (_this).f16, globalState);
      _nw309[(new BigNumber(3)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(4)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(5)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(6)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(7)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(8)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(9)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(10)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(11)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(12)).toNumber()] = _1638_v1;
      _nw309[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1638_v1, _1638_v1);
      _1654_v14 = _nw309;
      let _index286 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_1654_v14).length));
      (_1654_v14)[_index286] = _1638_v1;
      let _index287 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_1654_v14).length));
      (_1654_v14)[_index287] = _module.__default.fm13(((_this).f16).isLessThan((_this).f16), _1643_v6, new BigNumber((_1644_v7).length), globalState);
      let _1655_v15;
      _1655_v15 = _dafny.MultiSet.fromElements(false);
      if ((_1655_v15).IsProperSubsetOf((_this).fm11(new BigNumber(103), globalState))) {
        let _1656_v16;
        _1656_v16 = _module.D0.create_DC0((_1638_v1)[_module.__default.safeIndex((_this).f16, new BigNumber((_1638_v1).length))]);
        let _1657_v17;
        _1657_v17 = _dafny.MultiSet.fromElements(_1656_v16, _1656_v16);
        let _1658_v18;
        _1658_v18 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_1637_v0, false));
        let _1659_v19;
        _1659_v19 = _dafny.Map.Empty.slice().updateUnsafe((((_1657_v17).contains(_1656_v16)) ? ((_1657_v17).get(_1656_v16)) : ((_dafny.ZERO).minus((_this).f16))),_1658_v18);
        _1659_v19 = (_1659_v19).update((_this).f16, _1658_v18);
        let _1660_v20;
        _1660_v20 = _1641_v4;
        _1641_v4 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1641_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-332)), function (_1661_i2) {
          return (_this).f16;
        })), (_1660_v20));
        let _1662_v21;
        _1662_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1644_v7).length),(_this).f16);
        (globalState).f1 = !(_1662_v21).equals(_1662_v21);
        let _index288 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
        (_1639_v2)[_index288] = (_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))];
        let _1663_v22;
        let _nw310 = Array((new BigNumber(28)).toNumber());
        _nw310[(_dafny.ZERO).toNumber()] = (_this).f16;
        _nw310[(_dafny.ONE).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(2)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(3)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(4)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(5)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(6)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(7)).toNumber()] = new BigNumber(840);
        _nw310[(new BigNumber(8)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(9)).toNumber()] = new BigNumber(-130);
        _nw310[(new BigNumber(10)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(251),(_this).f16)).length);
        _nw310[(new BigNumber(12)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(13)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("wvp")).length);
        _nw310[(new BigNumber(15)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(16)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(17)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(18)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(19)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(20)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(21)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(22)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(23)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus((_this).f16);
        _nw310[(new BigNumber(25)).toNumber()] = (_dafny.ZERO).minus((_this).f16);
        _nw310[(new BigNumber(26)).toNumber()] = (_this).f16;
        _nw310[(new BigNumber(27)).toNumber()] = new BigNumber((_1644_v7).length);
        _1663_v22 = _nw310;
        let _1664_v23;
        _1664_v23 = _module.D3.create_DC6(_1663_v22);
        let _1665_v24;
        _1665_v24 = _dafny.Set.fromElements((_module.D3.create_DC6(_1663_v22)).dtor_cf6, (_1664_v23).dtor_cf6, _1663_v22, _1663_v22);
        _1665_v24 = (_dafny.Set.fromElements(_1663_v22)).Union((_1665_v24).Difference(_1665_v24));
      } else {
        let _1666_v25;
        _1666_v25 = _module.D0.create_DC0((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]);
        let _source12 = _1666_v25;
        if (_source12.is_DC1) {
          let _1667___mcc_h0 = (_source12).cf1;
          let _1668_cf1 = _1667___mcc_h0;
          let _1669_v26;
          _1669_v26 = _dafny.Set.fromElements(_1668_cf1);
          let _1670_v27;
          _1670_v27 = _dafny.Set.fromElements(new BigNumber((_1669_v26).length), (_this).f16);
          let _1671_v28;
          let _init46 = function (_1672_i3) {
            return _module.D0.create_DC2(_module.D0.create_DC1((_module.D0.create_DC1(new BigNumber(778))).dtor_cf1));
          };
          let _nw311 = Array((new BigNumber(19)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw311.length); _i0_46++) {
            _nw311[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1671_v28 = _nw311;
          let _1673_v29;
          let _nw312 = Array((new BigNumber(6)).toNumber());
          _nw312[(_dafny.ZERO).toNumber()] = (_this).f16;
          _nw312[(_dafny.ONE).toNumber()] = (_this).f16;
          _nw312[(new BigNumber(2)).toNumber()] = new BigNumber(450);
          _nw312[(new BigNumber(3)).toNumber()] = new BigNumber((_1670_v27).length);
          _nw312[(new BigNumber(4)).toNumber()] = (_this).f16;
          _nw312[(new BigNumber(5)).toNumber()] = (_this).f16;
          _1673_v29 = _nw312;
          _module.__default.m0(_1669_v26, _1671_v28, _1673_v29, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), ((_1674_v6) => function (_1675_i4) {
            return _1674_v6;
          })(_1643_v6)), _1644_v7), globalState);
          let _index289 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
          (_1639_v2)[_index289] = _1637_v0;
          let _1676_v30;
          _1676_v30 = _dafny.MultiSet.fromElements(_1668_cf1);
          let _1677_v31;
          _1677_v31 = _module.D1.create_DC4(_1643_v6);
          let _1678_v32;
          let _nw313 = Array((new BigNumber(10)).toNumber());
          _nw313[(_dafny.ZERO).toNumber()] = _1643_v6;
          _nw313[(_dafny.ONE).toNumber()] = _1643_v6;
          _nw313[(new BigNumber(2)).toNumber()] = (_1644_v7)[_module.__default.safeIndex(new BigNumber((_1676_v30).cardinality()), new BigNumber((_1644_v7).length))];
          _nw313[(new BigNumber(3)).toNumber()] = _1643_v6;
          _nw313[(new BigNumber(4)).toNumber()] = (_1677_v31).dtor_cf4;
          _nw313[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
          _nw313[(new BigNumber(6)).toNumber()] = _1643_v6;
          _nw313[(new BigNumber(7)).toNumber()] = _1643_v6;
          _nw313[(new BigNumber(8)).toNumber()] = _1643_v6;
          _nw313[(new BigNumber(9)).toNumber()] = _1643_v6;
          _1678_v32 = _nw313;
          let _index290 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1678_v32).length));
          (_1678_v32)[_index290] = _1643_v6;
          let _index291 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1678_v32).length));
          (_1678_v32)[_index291] = _1643_v6;
          let _1679_v33;
          _1679_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1668_cf1,_1644_v7);
          let _1680_v34;
          _1680_v34 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("mh"), (((_1679_v33).contains(_1668_cf1)) ? ((_1679_v33).get(_1668_cf1)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(221)), ((_1681_v6) => function (_1682_i5) {
            return _1681_v6;
          })(_1643_v6)))));
          let _1683_v35;
          _1683_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1680_v34).cardinality()),(_this).f16);
          (globalState).f6 = _module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(new BigNumber((_1655_v15).cardinality()))).length), (((_1683_v35).contains(_1668_cf1)) ? ((_1683_v35).get(_1668_cf1)) : (_1668_cf1)));
        } else if (_source12.is_DC0) {
          let _1684___mcc_h1 = (_source12).cf0;
          let _1685_cf0 = _1684___mcc_h1;
          let _1686_v36;
          _1686_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1637_v0,new BigNumber((_dafny.MultiSet.fromElements((_this).f16, (_this).f16)).cardinality()));
          let _1687_v37;
          _1687_v37 = _1686_v36;
          let _1688_v38;
          _1688_v38 = _dafny.Map.Empty.slice().updateUnsafe(((_1687_v37)).Merge((_1687_v37)),_1637_v0);
          let _1689_v39;
          _1689_v39 = _1641_v4;
          let _1690_v40;
          _1690_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))], globalState),_1689_v39);
          _1688_v38 = (_1688_v38).update(_module.__default.fm14(_1690_v40, (_this).f16, _1685_cf0, globalState), _1637_v0);
          (globalState).f1 = (_module.__default.fm1(_1685_cf0, globalState)).isEqualTo(new BigNumber((_1644_v7).length));
          let _1691_v42;
          _1691_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(776),(_this).f16);
          let _1692_v43;
          _1692_v43 = _module.D12.create_DC34(_1691_v42);
          let _1693_v44;
          let _nw314 = new _module.C3();
          _nw314.__ctor(_1637_v0, _module.__default.safeModuloInt((_this).f16, new BigNumber(851)), _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll64 = new _dafny.Map();
            for (const _compr_64 of ((_1692_v43).dtor_cf63).Keys.Elements) {
              let _1694_v41 = _compr_64;
              if (((_1692_v43).dtor_cf63).contains(_1694_v41)) {
                _coll64.push([_module.__default.safeDivisionInt(_1694_v41, (_this).f16),_1637_v0]);
              }
            }
            return _coll64;
          }(),(_this).f16));
          _1693_v44 = _nw314;
          let _1695_v45;
          _1695_v45 = _module.D3.create_DC7((_this).f16);
          let _1696_v46;
          _1696_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1639_v2,(_1693_v44).f18);
          let _1697_v47;
          _1697_v47 = _module.D10.create_DC27((_this).f16, _1644_v7, new BigNumber(-202), _1695_v45, _1696_v46);
          let _1698_v48;
          _1698_v48 = _module.D10.create_DC29(_1697_v47);
          let _1699_v49;
          _1699_v49 = _module.D10.create_DC29(_1698_v48);
          let _1700_v50;
          _1700_v50 = _module.D10.create_DC29(_1699_v49);
          let _1701_v51;
          _1701_v51 = _module.D10.create_DC29(_1699_v49);
          let _1702_v52;
          _1702_v52 = _dafny.MultiSet.fromElements(_1701_v51, _1701_v51, _1701_v51);
          _1702_v52 = _1702_v52;
        } else {
          let _1703___mcc_h2 = (_source12).cf2;
          let _1704_cf2 = _1703___mcc_h2;
          (globalState).f6 = (_this).f16;
          r0 = new BigNumber(((_1654_v14)[_module.__default.safeIndex(new BigNumber(43), new BigNumber((_1654_v14).length))]).length);
          let _index292 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
          (_1639_v2)[_index292] = _1637_v0;
          (globalState).f6 = (_this).f16;
        }
        let _index293 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
        let _rhs227 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f16));
        let _rhs228 = (((_1637_v0) ? ((_this).f16) : (new BigNumber(799)))).isLessThan(new BigNumber(705));
        let _rhs229 = _module.__default.fm1((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))], globalState);
        let _rhs230 = _1637_v0;
        let _lhs204 = globalState;
        let _lhs205 = globalState;
        let _lhs206 = globalState;
        let _lhs207 = _1639_v2;
        let _lhs208 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
        _lhs204.f6 = _rhs227;
        _lhs205.f1 = _rhs228;
        _lhs206.f6 = _rhs229;
        _lhs207[_lhs208] = _rhs230;
        let _1705_v53;
        _1705_v53 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)), _1643_v6);
        let _1706_v54;
        _1706_v54 = _dafny.MultiSet.fromElements((_this).f16, new BigNumber((_1705_v53).cardinality()), (_this).f16);
        let _1707_v55;
        _1707_v55 = _dafny.Set.fromElements(new BigNumber((_1706_v54).cardinality()));
        let _1708_v56;
        _1708_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_1707_v55);
        let _1709_v57;
        _1709_v57 = _dafny.Map.Empty.slice().updateUnsafe((_1641_v4)[_module.__default.safeIndex((_this).f16, new BigNumber((_1641_v4).length))],_1708_v56);
        let _1710_v58;
        _1710_v58 = _module.D9.create_DC24(_1709_v57, _1638_v1);
        let _source13 = _1710_v58;
        if (_source13.is_DC23) {
          let _1711___mcc_h3 = (_source13).cf37;
          let _1712___mcc_h4 = (_source13).cf38;
          let _1713___mcc_h5 = (_source13).cf39;
          let _1714___mcc_h6 = (_source13).cf40;
          let _1715___mcc_h7 = (_source13).cf41;
          let _1716_cf41 = _1715___mcc_h7;
          let _1717_cf40 = _1714___mcc_h6;
          let _1718_cf39 = _1713___mcc_h5;
          let _1719_cf38 = _1712___mcc_h4;
          let _1720_cf37 = _1711___mcc_h3;
          let _1721_v59;
          _1721_v59 = _module.D16.create_DC44((_dafny.MultiSet.fromElements(false, _1718_cf39)).Difference(_1655_v15));
          let _rhs231 = (_this).f16;
          let _rhs232 = ((_1637_v0) ? ((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]) : (_1637_v0));
          let _rhs233 = _1721_v59;
          let _lhs209 = globalState;
          _lhs209.f6 = _rhs231;
          _1718_cf39 = _rhs232;
          _1721_v59 = _rhs233;
          let _1722_v60;
          _1722_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(579)), function (_1723_i6) {
            return (_this).f16;
          }),(_this).f16);
          let _1724_v61;
          let _nw315 = Array((new BigNumber(6)).toNumber());
          _nw315[(_dafny.ZERO).toNumber()] = _1722_v60;
          _nw315[(_dafny.ONE).toNumber()] = _1722_v60;
          _nw315[(new BigNumber(2)).toNumber()] = _1722_v60;
          _nw315[(new BigNumber(3)).toNumber()] = _1722_v60;
          _nw315[(new BigNumber(4)).toNumber()] = (_1722_v60).update(_1641_v4, (_this).f16);
          _nw315[(new BigNumber(5)).toNumber()] = (_1722_v60).update(_dafny.Seq.of((_this).f16), (_this).f16);
          _1724_v61 = _nw315;
          let _index294 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_1724_v61).length));
          (_1724_v61)[_index294] = _1722_v60;
          let _1725_v62;
          _1725_v62 = _dafny.Set.fromElements(_1717_cf40);
          let _1726_v63;
          _1726_v63 = _dafny.Set.fromElements(_1637_v0);
          let _index295 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_1724_v61).length));
          let _rhs234 = (_this).f16;
          let _rhs235 = (_1725_v62).IsSubsetOf(_1725_v62);
          let _rhs236 = (_1722_v60).Merge((_1722_v60).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f16),new BigNumber(118))));
          let _rhs237 = new BigNumber((_1641_v4).length);
          let _rhs238 = new BigNumber((_1726_v63).length);
          let _lhs210 = globalState;
          let _lhs211 = _1724_v61;
          let _lhs212 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_1724_v61).length));
          let _lhs213 = globalState;
          let _lhs214 = globalState;
          _lhs210.f6 = _rhs234;
          _1718_cf39 = _rhs235;
          _lhs211[_lhs212] = _rhs236;
          _lhs213.f6 = _rhs237;
          _lhs214.f6 = _rhs238;
          let _1727_v64;
          _1727_v64 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(48)), ((_1728_cf41) => function (_1729_i7) {
            return _1728_cf41;
          })(_1716_cf41)),(((_1654_v14)[_module.__default.safeIndex(new BigNumber(43), new BigNumber((_1654_v14).length))])[_module.__default.safeIndex((_this).f16, new BigNumber(((_1654_v14)[_module.__default.safeIndex(new BigNumber(43), new BigNumber((_1654_v14).length))]).length))]) && (_1718_cf39));
          let _1730_v65;
          _1730_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1719_cf38).length),(_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]);
          let _index296 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
          let _rhs239 = (_this).f16;
          let _rhs240 = (_1727_v64).Merge(_1727_v64);
          let _rhs241 = !(((_1654_v14)[_module.__default.safeIndex(new BigNumber(43), new BigNumber((_1654_v14).length))])[_module.__default.safeIndex((new BigNumber((_1730_v65).length)).plus((_this).f16), new BigNumber(((_1654_v14)[_module.__default.safeIndex(new BigNumber(43), new BigNumber((_1654_v14).length))]).length))]);
          let _lhs215 = _1639_v2;
          let _lhs216 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
          r0 = _rhs239;
          _1727_v64 = _rhs240;
          _lhs215[_lhs216] = _rhs241;
          (globalState).f6 = (_this).f16;
        } else if (_source13.is_DC24) {
          let _1731___mcc_h8 = (_source13).cf42;
          let _1732___mcc_h9 = (_source13).cf43;
          let _1733_cf43 = _1732___mcc_h9;
          let _1734_cf42 = _1731___mcc_h8;
          _1644_v7 = _1644_v7;
          let _1735_v66;
          _1735_v66 = _module.D10.create_DC28((_this).f16);
          _1735_v66 = _module.D10.create_DC28((_this).f16);
          let _index297 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
          (_1639_v2)[_index297] = _module.__default.fm2(globalState);
          let _1736_v67;
          _1736_v67 = _dafny.MultiSet.fromElements(_1638_v1, _1733_cf43, _dafny.Seq.of((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))], _1637_v0), _dafny.Seq.of(_module.__default.fm2(globalState)));
          let _1737_v68;
          _1737_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]);
          let _1738_v69;
          _1738_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1737_v68,(_this).f16);
          let _1739_v70;
          let _nw316 = new _module.C2();
          _nw316.__ctor((((_1736_v67).contains(_dafny.Seq.of((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))], (_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))], _1637_v0))) ? ((_1736_v67).get(_dafny.Seq.of((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))], (_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))], _1637_v0))) : ((_this).f16)), (_this).f16, _1738_v69);
          _1739_v70 = _nw316;
        } else if (_source13.is_DC22) {
          let _1740___mcc_h10 = (_source13).cf36;
          let _1741_cf36 = _1740___mcc_h10;
          let _1742_v71;
          _1742_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1637_v0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(553)), ((_1743_v55) => function (_1744_i8) {
            return _1743_v55;
          })(_1707_v55)));
          let _1745_v72;
          _1745_v72 = _dafny.Seq.of(_1707_v55);
          let _1746_v73;
          _1746_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(480),_1637_v0);
          let _1747_v74;
          _1747_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1746_v73,(_this).f16);
          let _1748_v75;
          let _nw317 = new _module.C2();
          _nw317.__ctor((_this).f16, ((_this).f16).minus(new BigNumber(((((_1742_v71).contains((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))])) ? ((_1742_v71).get((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))])) : (_1745_v72))).length)), (_1747_v74).Merge(_1747_v74));
          _1748_v75 = _nw317;
          let _1749_v76;
          _1749_v76 = _dafny.Set.fromElements(_1637_v0, (_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]);
          let _1750_v77;
          _1750_v77 = _module.D11.create_DC32(!((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]) || (_1637_v0), (_dafny.ZERO).minus((_1748_v75).f20), ((_1706_v54).update(new BigNumber((_1749_v76).length), _module.__default.abs(_1748_v75.f21))).IsProperSubsetOf(_1706_v54));
          _1750_v77 = ((_module.__default.fm2(globalState)) ? ((((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]) ? (_1750_v77) : (_1750_v77))) : (_1750_v77));
          let _1751_v79;
          let _init47 = ((_1752_v6, _1753_v7, _1754_v0) => function (_1755_i9) {
            return function () {
              let _coll65 = new _dafny.Map();
              for (const _compr_65 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_1752_v6),_1753_v7)).Keys.Elements) {
                let _1756_v78 = _compr_65;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(_1752_v6),_1753_v7)).contains(_1756_v78)) {
                  _coll65.push([_1756_v78,_1754_v0]);
                }
              }
              return _coll65;
            }();
          })(_1643_v6, _1644_v7, _1637_v0);
          let _nw318 = Array((new BigNumber(27)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw318.length); _i0_47++) {
            _nw318[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1751_v79 = _nw318;
          let _1757_v80;
          _1757_v80 = _module.D20.create_DC49(_1751_v79);
          let _1758_v81;
          let _nw319 = Array((new BigNumber(26)).toNumber());
          _nw319[(_dafny.ZERO).toNumber()] = _1751_v79;
          _nw319[(_dafny.ONE).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(2)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(3)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(4)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(5)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(6)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(7)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(8)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(9)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(10)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(11)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(12)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(13)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(14)).toNumber()] = (_1757_v80).dtor_cf85;
          _nw319[(new BigNumber(15)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(16)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(17)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(18)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(19)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(20)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(21)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(22)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(23)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(24)).toNumber()] = _1751_v79;
          _nw319[(new BigNumber(25)).toNumber()] = _1751_v79;
          _1758_v81 = _nw319;
          let _index298 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_1758_v81).length));
          (_1758_v81)[_index298] = ((_1637_v0) ? (_1751_v79) : (_1751_v79));
          let _index299 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_1758_v81).length));
          let _rhs242 = ((_1748_v75).f20).isLessThanOrEqualTo(_1748_v75.f21);
          let _rhs243 = ((!(_1637_v0)) ? (_1751_v79) : (_1751_v79));
          let _rhs244 = _1707_v55;
          let _lhs217 = globalState;
          let _lhs218 = _1758_v81;
          let _lhs219 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_1758_v81).length));
          _lhs217.f1 = _rhs242;
          _lhs218[_lhs219] = _rhs243;
          _1707_v55 = _rhs244;
          let _1759_v82;
          let _nw320 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1759_v82 = _nw320;
          _1759_v82 = _1759_v82;
        } else {
          let _1760___mcc_h11 = (_source13).cf44;
          let _1761_cf44 = _1760___mcc_h11;
          let _1762_v83;
          let _init48 = ((_1763_v4) => function (_1764_i10) {
            return (_1764_i10).plus(new BigNumber((_dafny.Seq.update(_1763_v4, _module.__default.safeIndex((_this).f16, new BigNumber((_1763_v4).length)), (_this).f16)).length));
          })(_1641_v4);
          let _nw321 = Array((new BigNumber(20)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw321.length); _i0_48++) {
            _nw321[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1762_v83 = _nw321;
          let _index300 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_1762_v83).length));
          (_1762_v83)[_index300] = (_this).f16;
          let _index301 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_1762_v83).length));
          (_1762_v83)[_index301] = (_this).f16;
          let _1765_v84;
          _1765_v84 = _dafny.Map.Empty.slice().updateUnsafe((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))],_1762_v83);
          (globalState).f6 = ((_1762_v83)[_module.__default.safeIndex(new BigNumber(615), new BigNumber((_1762_v83).length))]).minus(new BigNumber((_1765_v84).length));
          let _1766_v85;
          _1766_v85 = _dafny.MultiSet.fromElements(_1655_v15, _1655_v15);
          _1766_v85 = ((_1637_v0) ? (_1766_v85) : (_1766_v85));
          let _1767_v86;
          _1767_v86 = _dafny.Map.Empty.slice().updateUnsafe(true,_1637_v0);
          let _index302 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
          let _rhs245 = false;
          let _rhs246 = (((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]) ? ((((_1767_v86).contains((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))])) ? ((_1767_v86).get((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))])) : ((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]))) : ((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]));
          let _rhs247 = (_this).f16;
          let _lhs220 = globalState;
          let _lhs221 = _1639_v2;
          let _lhs222 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
          let _lhs223 = globalState;
          _lhs220.f1 = _rhs245;
          _lhs221[_lhs222] = _rhs246;
          _lhs223.f6 = _rhs247;
        }
        let _1768_v87;
        _1768_v87 = _dafny.Set.fromElements(_1637_v0);
        let _1769_v88;
        _1769_v88 = _module.D8.create_DC19(_1768_v87);
        _1769_v88 = _1769_v88;
        if (((_1641_v4)[_module.__default.safeIndex((_this).f16, new BigNumber((_1641_v4).length))]).isLessThan((_this).f16)) {
          r1 = true;
          let _1770_v89;
          _1770_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1643_v6,_1644_v7);
          let _1771_v90;
          _1771_v90 = _module.D12.create_DC35();
          let _1772_v91;
          _1772_v91 = _dafny.Map.Empty.slice().updateUnsafe(_1771_v90,_dafny.Map.Empty.slice().updateUnsafe(_1643_v6,_1644_v7));
          _1770_v89 = (((_1772_v91).contains(_1771_v90)) ? ((_1772_v91).get(_1771_v90)) : (_1770_v89));
          (globalState).f6 = (_dafny.ZERO).minus((_this).f16);
          (globalState).f6 = (_this).f16;
          let _1773_v92;
          let _init49 = function (_1774_i11) {
            return (_1774_i11).plus(new BigNumber(639));
          };
          let _nw322 = Array((new BigNumber(28)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw322.length); _i0_49++) {
            _nw322[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1773_v92 = _nw322;
          (globalState).f9 = _1773_v92;
        } else {
          (globalState).f6 = new BigNumber((_1768_v87).length);
          let _1775_v93;
          let _nw323 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
          _1775_v93 = _nw323;
          let _1776_v94;
          let _nw324 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _1776_v94 = _nw324;
          let _1777_v95;
          _1777_v95 = _dafny.Map.Empty.slice().updateUnsafe(_1775_v93,(((_1638_v1)[_module.__default.safeIndex((_this).f16, new BigNumber((_1638_v1).length))]) ? (_1776_v94) : (_1776_v94)));
          let _1778_v96;
          _1778_v96 = _dafny.Seq.of(_1777_v95);
          let _1779_v97;
          _1779_v97 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1641_v4).length),(_this).f16);
          _1777_v95 = (_1778_v96)[_module.__default.safeIndex(((_this).f16).minus(new BigNumber((_1779_v97).length)), new BigNumber((_1778_v96).length))];
          _1706_v54 = (_1706_v54).Difference((_dafny.MultiSet.FromArray(_dafny.Seq.update(_1641_v4, _module.__default.safeIndex((_this).f16, new BigNumber((_1641_v4).length)), (_this).f16))).update(new BigNumber(52), _module.__default.abs((_this).f16)));
          let _index303 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length));
          (_1639_v2)[_index303] = _1637_v0;
          let _index304 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1776_v94).length));
          (_1776_v94)[_index304] = (_this).f16;
          let _index305 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1776_v94).length));
          let _rhs248 = _1637_v0;
          let _rhs249 = _module.__default.safeDivisionInt((_this).f16, (_this).f16);
          let _rhs250 = (new BigNumber(749)).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("ayqx")).length));
          let _rhs251 = _1637_v0;
          let _rhs252 = (((_1706_v54).contains(_module.__default.safeDivisionInt(new BigNumber(150), (((_1655_v15).contains((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))])) ? ((_1655_v15).get((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))])) : ((_this).f16))))) ? ((_1706_v54).get(_module.__default.safeDivisionInt(new BigNumber(150), (((_1655_v15).contains((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))])) ? ((_1655_v15).get((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))])) : ((_this).f16))))) : (((((_1706_v54).contains(new BigNumber(582))) ? ((_1706_v54).get(new BigNumber(582))) : ((_this).f16))).multipliedBy((_dafny.ZERO).minus((_this).f16))));
          let _lhs224 = globalState;
          let _lhs225 = _1776_v94;
          let _lhs226 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1776_v94).length));
          let _lhs227 = globalState;
          let _lhs228 = globalState;
          _lhs224.f1 = _rhs248;
          _lhs225[_lhs226] = _rhs249;
          _lhs227.f1 = _rhs250;
          _lhs228.f1 = _rhs251;
          r0 = _rhs252;
        }
      }
      _1637_v0 = (_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f16,_1643_v6)).length), (_this).f16)).isLessThanOrEqualTo((_this).f16);
      (globalState).f6 = (_this).f16;
      r0 = (_this).f16;
      r1 = _dafny.Seq.contains(_dafny.Seq.Concat(_module.__default.fm13((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))], new _dafny.CodePoint('f'.codePointAt(0)), (_this).f16, globalState), _1638_v1), (true) || ((_1639_v2)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1639_v2).length))]));
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let _1780_i0;
      _1780_i0 = _dafny.ZERO;
      L13: {
        while (p0) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1780_i0)) {
              break L13;
            }
            _1780_i0 = (_1780_i0).plus(_dafny.ONE);
            let _1781_v0;
            _1781_v0 = _module.D14.create_DC40(false);
            let _1782_v1;
            _1782_v1 = _dafny.Seq.of(p0, (_1781_v0).dtor_cf71);
            _1782_v1 = _dafny.Seq.update(_dafny.Seq.of(p0), _module.__default.safeIndex(((p0) ? ((_this).f16) : ((_this).f16)), new BigNumber((_dafny.Seq.of(p0)).length)), p0);
            (globalState).f6 = new BigNumber(386);
            let _1783_v2;
            let _nw325 = Array((new BigNumber(12)).toNumber()).fill(_module.D3.Default());
            _1783_v2 = _nw325;
            let _1784_v3;
            _1784_v3 = _dafny.Set.fromElements(_1783_v2, _1783_v2);
            let _1785_v4;
            _1785_v4 = _dafny.Seq.UnicodeFromString("eqngik");
            let _1786_v5;
            _1786_v5 = _dafny.Seq.of(_dafny.Seq.of(p0, true), _1782_v1, _1782_v1);
            let _1787_v6;
            let _nw326 = new _module.C1();
            _nw326.__ctor((_this).f16);
            _1787_v6 = _nw326;
            let _1788_v7;
            _1788_v7 = _dafny.Seq.of(_1787_v6, _1787_v6);
            let _1789_v8;
            _1789_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1788_v7);
            let _1790_v9;
            _1790_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,new BigNumber((_1782_v1).length));
            let _1791_v10;
            let _nw327 = Array((new BigNumber(5)).toNumber());
            _nw327[(_dafny.ZERO).toNumber()] = new BigNumber((_1789_v8).length);
            _nw327[(_dafny.ONE).toNumber()] = (((_1790_v9).contains((_this).f16)) ? ((_1790_v9).get((_this).f16)) : ((_this).f16));
            _nw327[(new BigNumber(2)).toNumber()] = (_this).f16;
            _nw327[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_this).f16);
            _nw327[(new BigNumber(4)).toNumber()] = (_this).f16;
            _1791_v10 = _nw327;
            let _1792_v11;
            _1792_v11 = _module.D9.create_DC23(_1784_v3, _1785_v4, _dafny.Seq.IsProperPrefixOf((_1786_v5)[_module.__default.safeIndex((_this).f16, new BigNumber((_1786_v5).length))], _dafny.Seq.of(p0, p0)), _1791_v10, _module.__default.fm29(p0, (_this).f16, (_this).f16, p0, globalState));
            let _source14 = _1792_v11;
            if (_source14.is_DC23) {
              let _1793___mcc_h0 = (_source14).cf37;
              let _1794___mcc_h1 = (_source14).cf38;
              let _1795___mcc_h2 = (_source14).cf39;
              let _1796___mcc_h3 = (_source14).cf40;
              let _1797___mcc_h4 = (_source14).cf41;
              let _1798_cf41 = _1797___mcc_h4;
              let _1799_cf40 = _1796___mcc_h3;
              let _1800_cf39 = _1795___mcc_h2;
              let _1801_cf38 = _1794___mcc_h1;
              let _1802_cf37 = _1793___mcc_h0;
              let _1803_v12;
              _1803_v12 = _dafny.Set.fromElements(p0, p0);
              _1803_v12 = (((_1782_v1)[_module.__default.safeIndex((_this).f16, new BigNumber((_1782_v1).length))]) ? ((_1803_v12).Difference(_1803_v12)) : (_1803_v12));
              _1782_v1 = _1782_v1;
              let _1804_v13;
              let _nw328 = Array((new BigNumber(19)).toNumber()).fill([]);
              _1804_v13 = _nw328;
              let _index306 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_1804_v13).length));
              (_1804_v13)[_index306] = _1791_v10;
              let _index307 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_1804_v13).length));
              (_1804_v13)[_index307] = _1791_v10;
              _1800_cf39 = _module.__default.fm2(globalState);
            } else if (_source14.is_DC24) {
              let _1805___mcc_h5 = (_source14).cf42;
              let _1806___mcc_h6 = (_source14).cf43;
              let _1807_cf43 = _1806___mcc_h6;
              let _1808_cf42 = _1805___mcc_h5;
              (globalState).f1 = _module.__default.fm2(globalState);
              let _1809_v14;
              let _nw329 = new _module.C1();
              _nw329.__ctor((_this).f16);
              _1809_v14 = _nw329;
              let _1810_v15;
              _1810_v15 = _dafny.Seq.of((_1809_v14).f24, new BigNumber(-959), (_this).f16, (_this).f16, new BigNumber(148));
              let _1811_v16;
              _1811_v16 = new _dafny.CodePoint('e'.codePointAt(0));
              let _1812_v17;
              _1812_v17 = _module.D3.create_DC8(_1810_v15, _1785_v4, p0, _1811_v16);
              let _1813_v18;
              _1813_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1812_v17);
              let _1814_v19;
              _1814_v19 = _dafny.MultiSet.fromElements(p0);
              _1813_v18 = (_1813_v18).update((_1814_v19).IsProperSubsetOf(_1814_v19), _1812_v17);
              _1790_v9 = ((((p0) ? (p0) : (p0))) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1807_cf43).length),(_this).f16)) : (_1790_v9));
            } else if (_source14.is_DC22) {
              let _1815___mcc_h7 = (_source14).cf36;
              let _1816_cf36 = _1815___mcc_h7;
              (globalState).f1 = p0;
              let _1817_v21;
              _1817_v21 = _dafny.Map.Empty.slice().updateUnsafe(function () {
                let _coll66 = new _dafny.Map();
                for (const _compr_66 of (_1790_v9).Keys.Elements) {
                  let _1818_v20 = _compr_66;
                  if ((_1790_v9).contains(_1818_v20)) {
                    _coll66.push([_module.__default.safeModuloInt(_1818_v20, (_this).f16),p0]);
                  }
                }
                return _coll66;
              }(),p0);
              let _1819_v22;
              _1819_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_module.__default.fm2(globalState));
              _1817_v21 = (_1817_v21).update(_1819_v22, p0);
              let _1820_v23;
              _1820_v23 = _module.D20.create_DC50(p0, (_this).f16, new BigNumber(719), p0, p0);
              let _1821_v24;
              _1821_v24 = _dafny.Seq.of(_1820_v23, _1820_v23);
              (globalState).f6 = new BigNumber((_1821_v24).length);
              (globalState).f6 = ((_this).f16).plus((_this).f16);
            } else {
              let _1822___mcc_h8 = (_source14).cf44;
              let _1823_cf44 = _1822___mcc_h8;
              let _1824_v25;
              _1824_v25 = _dafny.Seq.of(_1791_v10, _1791_v10, _1791_v10, (_1792_v11).dtor_cf40, _1791_v10);
              let _1825_v26;
              _1825_v26 = new _dafny.CodePoint('a'.codePointAt(0));
              let _1826_v27;
              _1826_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1782_v1).length),_1825_v26);
              let _1827_v28;
              _1827_v28 = _dafny.MultiSet.fromElements(new BigNumber((_1826_v27).length));
              let _1828_v29;
              _1828_v29 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1827_v28).cardinality()));
              _1824_v25 = _dafny.Seq.update(((!(p0) || (p0)) ? (_1824_v25) : (((p0) ? (_1824_v25) : (_1824_v25)))), _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f16, (((_1828_v29).contains(p0)) ? ((_1828_v29).get(p0)) : ((_this).f16)))), new BigNumber((((!(p0) || (p0)) ? (_1824_v25) : (((p0) ? (_1824_v25) : (_1824_v25))))).length)), _1791_v10);
              let _1829_v30;
              let _nw330 = Array((new BigNumber(2)).toNumber());
              _nw330[(_dafny.ZERO).toNumber()] = p0;
              _nw330[(_dafny.ONE).toNumber()] = p0;
              _1829_v30 = _nw330;
              let _index308 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1829_v30).length));
              (_1829_v30)[_index308] = p0;
              let _index309 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1829_v30).length));
              (_1829_v30)[_index309] = p0;
              _1787_v6 = _1787_v6;
              let _1830_v31;
              let _1831_v32;
              let _out12;
              let _out13;
              let _outcollector4 = (_1787_v6).m1(globalState);
              _out12 = _outcollector4[0];
              _out13 = _outcollector4[1];
              _1830_v31 = _out12;
              _1831_v32 = _out13;
            }
            let _1832_v33;
            _1832_v33 = _dafny.Seq.of((_this).f16, new BigNumber(971));
            let _1833_v34;
            _1833_v34 = _dafny.MultiSet.fromElements((_this).f16);
            let _1834_v36;
            _1834_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,p0);
            let _1835_v37;
            _1835_v37 = _dafny.Map.Empty.slice().updateUnsafe(function () {
              let _coll67 = new _dafny.Map();
              for (const _compr_67 of (_1834_v36).Keys.Elements) {
                let _1836_v35 = _compr_67;
                if ((_1834_v36).contains(_1836_v35)) {
                  _coll67.push([(_1836_v35).multipliedBy(new BigNumber(630)),p0]);
                }
              }
              return _coll67;
            }(),(_this).f16);
            let _1837_v38;
            let _nw331 = new _module.C2();
            _nw331.__ctor((_1832_v33)[_module.__default.safeIndex((((_1790_v9).contains(new BigNumber(727))) ? ((_1790_v9).get(new BigNumber(727))) : (new BigNumber((_1833_v34).cardinality()))), new BigNumber((_1832_v33).length))], _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f16), (_this).f16), _1835_v37);
            _1837_v38 = _nw331;
          }
        }
      }
      let _1838_v39;
      _1838_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(693),_dafny.Set.fromElements((_this).f16));
      let _source15 = _1838_v39;
      let _1839___mcc_h9 = _source15;
      let _1840_cf19 = _1839___mcc_h9;
      let _1841_v41;
      let _nw332 = Array((new BigNumber(19)).toNumber()).fill(_module.D0.Default());
      _1841_v41 = _nw332;
      let _1842_v42;
      _1842_v42 = _1841_v41;
      let _1843_v43;
      let _nw333 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _1843_v43 = _nw333;
      let _1844_v44;
      _1844_v44 = _dafny.Seq.UnicodeFromString("bogbeg");
      _module.__default.m0(function () {
        let _coll68 = new _dafny.Set();
        for (const _compr_68 of _dafny.IntegerRange(new BigNumber(972), new BigNumber(980))) {
          let _1845_v40 = _compr_68;
          if (((new BigNumber(972)).isLessThanOrEqualTo(_1845_v40)) && ((_1845_v40).isLessThan(new BigNumber(980)))) {
            _coll68.add((_1845_v40).multipliedBy((_this).f16));
          }
        }
        return _coll68;
      }(), (_1842_v42), _1843_v43, _1844_v44, globalState);
      (globalState).f9 = _1843_v43;
      let _1846_v45;
      _1846_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,p0);
      let _1847_v46;
      _1847_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1843_v43,(_this).f16);
      let _1848_v47;
      _1848_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1846_v45,(_this).f16);
      let _1849_v48;
      let _nw334 = new _module.C3();
      _nw334.__ctor(p0, (_this).f16, _1848_v47);
      _1849_v48 = _nw334;
      let _1850_v49;
      _1850_v49 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1849_v48);
      let _1851_v50;
      let _nw335 = new _module.C3();
      _nw335.__ctor(p0, new BigNumber((_1844_v44).length), _dafny.Map.Empty.slice().updateUnsafe(_1846_v45,(((_1847_v46).contains(_1843_v43)) ? ((_1847_v46).get(_1843_v43)) : (new BigNumber((_1850_v49).length)))));
      _1851_v50 = _nw335;
      if (p0) {
        let _1852_v51;
        _1852_v51 = _module.D20.create_DC50((_1851_v50).f18, (_1851_v50).f19, (_this).f16, !(_module.__default.fm2(globalState)), _module.__default.fm2(globalState));
        let _1853_v52;
        _1853_v52 = _dafny.MultiSet.fromElements((_1851_v50).f18);
        let _1854_v53;
        _1854_v53 = _dafny.MultiSet.fromElements((((_1853_v52).contains(p0)) ? ((_1853_v52).get(p0)) : ((_dafny.ZERO).minus((_1851_v50).f19))));
        let _pat_let_tv37 = _1851_v50;
        let _pat_let_tv38 = _1854_v53;
        let _pat_let_tv39 = _1851_v50;
        (globalState).f1 = !((function (_pat_let61_0) {
          return function (_1855_dt__update__tmp_h0) {
            return function (_pat_let62_0) {
              return function (_1856_dt__update_hcf89_h0) {
                return function (_pat_let63_0) {
                  return function (_1857_dt__update_hcf87_h0) {
                    return function (_pat_let64_0) {
                      return function (_1858_dt__update_hcf88_h0) {
                        return _module.D20.create_DC50((_1855_dt__update__tmp_h0).dtor_cf86, _1857_dt__update_hcf87_h0, _1858_dt__update_hcf88_h0, _1856_dt__update_hcf89_h0, (_1855_dt__update__tmp_h0).dtor_cf90);
                      }(_pat_let64_0);
                    }((_pat_let_tv39).f19);
                  }(_pat_let63_0);
                }(new BigNumber((_pat_let_tv38).cardinality()));
              }(_pat_let62_0);
            }((_pat_let_tv37).f18);
          }(_pat_let61_0);
        }(_1852_v51)).dtor_cf86);
        _1844_v44 = _1844_v44;
        let _1859_v54;
        _1859_v54 = _dafny.Seq.of(_1843_v43);
        let _1860_v55;
        _1860_v55 = _dafny.Set.fromElements((_1851_v50).f19);
        let _1861_v56;
        _1861_v56 = _module.D3.create_DC6((_1859_v54)[_module.__default.safeIndex(new BigNumber((_1860_v55).length), new BigNumber((_1859_v54).length))]);
        let _1862_v57;
        _1862_v57 = _dafny.Seq.of(_1861_v56);
        let _1863_v58;
        let _nw336 = new _module.C0();
        _nw336.__ctor(_1862_v57, (_1851_v50).f18);
        _1863_v58 = _nw336;
        let _1864_v59;
        _1864_v59 = _dafny.Map.Empty.slice().updateUnsafe((_1851_v50).f19,_1863_v58);
        let _1865_v60;
        let _nw337 = Array((new BigNumber(12)).toNumber()).fill(false);
        _1865_v60 = _nw337;
        let _1866_v61;
        _1866_v61 = _dafny.Seq.of(_1865_v60, _1865_v60);
        let _index310 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_1843_v43).length));
        (_1843_v43)[_index310] = new BigNumber((((_1864_v59).update((_1851_v50).f19, _1863_v58)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1866_v61).length),_1863_v58))).length);
        let _index311 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_1843_v43).length));
        (_1843_v43)[_index311] = (_1851_v50).f19;
        (globalState).f6 = (new BigNumber(609)).multipliedBy((_1851_v50).f19);
        let _index312 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_1843_v43).length));
        (_1843_v43)[_index312] = new BigNumber((_1844_v44).length);
      } else {
        (globalState).f6 = _module.__default.fm1(p0, globalState);
        let _1867_v62;
        let _nw338 = new _module.C4();
        _nw338.__ctor();
        _1867_v62 = _nw338;
        let _1868_v63;
        _1868_v63 = new _dafny.CodePoint('j'.codePointAt(0));
        _1868_v63 = _1868_v63;
        let _index313 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_1843_v43).length));
        (_1843_v43)[_index313] = (_this).f16;
        let _index314 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_1843_v43).length));
        (_1843_v43)[_index314] = (_1851_v50).f19;
        let _1869_v64;
        _1869_v64 = _dafny.MultiSet.fromElements((_1851_v50).f18);
        (globalState).f6 = (((_1851_v50).f18) ? ((((_1869_v64).contains(_module.__default.fm2(globalState))) ? ((_1869_v64).get(_module.__default.fm2(globalState))) : (new BigNumber(328)))) : (_module.__default.safeDivisionInt((_this).f16, (_1851_v50).f19)));
      }
      let _1870_v65;
      _1870_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,p0);
      let _1871_v66;
      _1871_v66 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus((_this).f16));
      let _1872_v67;
      _1872_v67 = _1871_v66;
      _1870_v65 = function (_source16) {
        let _1873___mcc_h10 = _source16;
        let _1874_cf12 = _1873___mcc_h10;
        return function () {
          let _coll69 = new _dafny.Map();
          for (const _compr_69 of (_dafny.Set.fromElements((_this).f16, (_this).f16, new BigNumber(715))).Elements) {
            let _1875_v68 = _compr_69;
            if ((_dafny.Set.fromElements((_this).f16, (_this).f16, new BigNumber(715))).contains(_1875_v68)) {
              _coll69.push([_module.__default.safeModuloInt(_1875_v68, (_this).f16),false]);
            }
          }
          return _coll69;
        }();
      }(_1872_v67);
      let _1876_v69;
      let _nw339 = Array((new BigNumber(29)).toNumber()).fill([]);
      _1876_v69 = _nw339;
      let _1877_v70;
      _1877_v70 = _module.D10.create_DC26(_1876_v69);
      _1877_v70 = _module.D10.create_DC26(_1876_v69);
      let _1878_v71;
      let _nw340 = new _module.C4();
      _nw340.__ctor();
      _1878_v71 = _nw340;
      let _1879_v72;
      let _init50 = ((_1880_p0) => function (_1881_i1) {
        return _1880_p0;
      })(p0);
      let _nw341 = Array((new BigNumber(3)).toNumber());
      for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw341.length); _i0_50++) {
        _nw341[_i0_50] = _init50(new BigNumber(_i0_50));
      }
      _1879_v72 = _nw341;
      let _1882_v73;
      let _nw342 = Array((new BigNumber(7)).toNumber());
      _nw342[(_dafny.ZERO).toNumber()] = _1879_v72;
      _nw342[(_dafny.ONE).toNumber()] = _1879_v72;
      _nw342[(new BigNumber(2)).toNumber()] = _1879_v72;
      _nw342[(new BigNumber(3)).toNumber()] = _1879_v72;
      _nw342[(new BigNumber(4)).toNumber()] = _1879_v72;
      _nw342[(new BigNumber(5)).toNumber()] = _1879_v72;
      _nw342[(new BigNumber(6)).toNumber()] = _1879_v72;
      _1882_v73 = _nw342;
      let _1883_v74;
      _1883_v74 = _module.D5.create_DC10(_1882_v73);
      let _source17 = _1883_v74;
      if (_source17.is_DC11) {
        let _1884___mcc_h11 = (_source17).cf14;
        let _1885_cf14 = _1884___mcc_h11;
        let _1886_v75;
        _1886_v75 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1887_v76;
        _1887_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1886_v75);
        let _1888_v77;
        _1888_v77 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
        let _1889_v78;
        _1889_v78 = _module.D14.create_DC39(_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), ((_1890_p0) => function (_1891_i2) {
  return new BigNumber((_dafny.Seq.of(_1890_p0, _1890_p0)).length);
})(p0)), _1888_v77, p0);
        let _1892_v79;
        let _nw343 = Array((new BigNumber(24)).toNumber());
        _nw343[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_1886_v75)).update(p0, _1886_v75);
        _nw343[(_dafny.ONE).toNumber()] = (_1887_v76).Merge(_1887_v76);
        _nw343[(new BigNumber(2)).toNumber()] = (_1887_v76).Merge(_1887_v76);
        _nw343[(new BigNumber(3)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_1889_v78).dtor_cf70,_1886_v75)).Merge(_1887_v76);
        _nw343[(new BigNumber(5)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(6)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(7)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(8)).toNumber()] = (_1887_v76).update(p0, _1886_v75);
        _nw343[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,new _dafny.CodePoint('g'.codePointAt(0)));
        _nw343[(new BigNumber(10)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(11)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(12)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(13)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(14)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_1886_v75)).Merge(_1887_v76);
        _nw343[(new BigNumber(15)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(16)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(17)).toNumber()] = (_1887_v76).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1886_v75));
        _nw343[(new BigNumber(18)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_1886_v75);
        _nw343[(new BigNumber(20)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(21)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(22)).toNumber()] = _1887_v76;
        _nw343[(new BigNumber(23)).toNumber()] = _1887_v76;
        _1892_v79 = _nw343;
        let _1893_v80;
        _1893_v80 = _dafny.Seq.of((_dafny.ZERO).minus(_1885_cf14), _1885_cf14);
        let _1894_v81;
        _1894_v81 = _dafny.Seq.UnicodeFromString("xkxwacsi");
        let _rhs253 = _1892_v79;
        let _rhs254 = _1893_v80;
        let _rhs255 = _1886_v75;
        let _rhs256 = _1894_v81;
        _1892_v79 = _rhs253;
        _1893_v80 = _rhs254;
        _1886_v75 = _rhs255;
        _1894_v81 = _rhs256;
        if (p0) {
          let _1895_v82;
          _1895_v82 = _module.D5.create_DC11(_1885_cf14);
          let _1896_v83;
          let _nw344 = Array((new BigNumber(10)).toNumber());
          _nw344[(_dafny.ZERO).toNumber()] = (_1895_v82).dtor_cf14;
          _nw344[(_dafny.ONE).toNumber()] = (_1893_v80)[_module.__default.safeIndex(_1885_cf14, new BigNumber((_1893_v80).length))];
          _nw344[(new BigNumber(2)).toNumber()] = _1885_cf14;
          _nw344[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_this).f16)).length);
          _nw344[(new BigNumber(4)).toNumber()] = _module.__default.fm1(p0, globalState);
          _nw344[(new BigNumber(5)).toNumber()] = new BigNumber(-536);
          _nw344[(new BigNumber(6)).toNumber()] = _1885_cf14;
          _nw344[(new BigNumber(7)).toNumber()] = new BigNumber(363);
          _nw344[(new BigNumber(8)).toNumber()] = (_this).f16;
          _nw344[(new BigNumber(9)).toNumber()] = (_this).f16;
          _1896_v83 = _nw344;
          let _1897_v84;
          _1897_v84 = _dafny.Set.fromElements(p0, p0);
          let _index315 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1896_v83).length));
          (_1896_v83)[_index315] = new BigNumber((_1897_v84).length);
          let _index316 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1896_v83).length));
          (_1896_v83)[_index316] = (_this).f16;
          (globalState).f6 = new BigNumber(407);
          let _1898_v85;
          let _nw345 = Array((new BigNumber(27)).toNumber()).fill(_module.D14.Default());
          _1898_v85 = _nw345;
          _1898_v85 = _1898_v85;
          _1894_v81 = _1894_v81;
          (globalState).f1 = !(((_dafny.ZERO).minus((_this).f16)).isLessThanOrEqualTo(_1885_cf14)) || (_module.__default.fm2(globalState));
        } else {
          (globalState).f1 = !(p0);
          let _1899_v86;
          let _init51 = ((_1900_cf14) => function (_1901_i3) {
            return (_dafny.MultiSet.fromElements((_this).f16, new BigNumber(860), (_this).f16, (_this).f16)).Union(_dafny.MultiSet.fromElements(_1900_cf14));
          })(_1885_cf14);
          let _nw346 = Array((new BigNumber(2)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw346.length); _i0_51++) {
            _nw346[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1899_v86 = _nw346;
          let _1902_v89;
          _1902_v89 = _dafny.Seq.of(function () {
            let _coll70 = new _dafny.Set();
            for (const _compr_70 of _dafny.IntegerRange(new BigNumber(880), new BigNumber(-220))) {
              let _1903_v88 = _compr_70;
              if (((new BigNumber(880)).isLessThanOrEqualTo(_1903_v88)) && ((_1903_v88).isLessThan(new BigNumber(-220)))) {
                _coll70.add(_module.__default.safeModuloInt(_1903_v88, new BigNumber(239)));
              }
            }
            return _coll70;
          }());
          let _index317 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1899_v86).length));
          (_1899_v86)[_index317] = _dafny.MultiSet.fromElements((((_1871_v66).contains(p0)) ? ((_1871_v66).get(p0)) : (new BigNumber((function () {
            let _coll71 = new _dafny.Map();
            for (const _compr_71 of (_1902_v89).Elements) {
              let _1904_v87 = _compr_71;
              if (_dafny.Seq.contains(_1902_v89, _1904_v87)) {
                _coll71.push([_1904_v87,p0]);
              }
            }
            return _coll71;
          }()).length))));
          let _1905_v90;
          _1905_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1894_v81).length),(_dafny.ZERO).minus(_1885_cf14));
          let _1906_v91;
          _1906_v91 = _dafny.MultiSet.fromElements((_this).f16);
          let _index318 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1899_v86).length));
          (_1899_v86)[_index318] = (_dafny.MultiSet.fromElements(_1885_cf14, (((_1905_v90).contains((_this).f16)) ? ((_1905_v90).get((_this).f16)) : ((_this).f16)))).Difference(_1906_v91);
          (globalState).f1 = ((!(p0) || (p0)) ? (((p0) ? (p0) : (true))) : (p0));
          (globalState).f6 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1894_v81, _dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), ((_1907_v75) => function (_1908_i4) {
            return _1907_v75;
          })(_1886_v75))), _1894_v81)).length);
          let _index319 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1879_v72).length));
          (_1879_v72)[_index319] = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("emichywvs"), _1886_v75);
          let _index320 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1879_v72).length));
          (_1879_v72)[_index320] = (new BigNumber((_module.__default.fm21(globalState)).length)).isLessThanOrEqualTo(_1885_cf14);
        }
        let _1909_v92;
        _1909_v92 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1885_cf14,p0),_1885_cf14);
        let _1910_v93;
        let _nw347 = new _module.C3();
        _nw347.__ctor(p0, new BigNumber((_1888_v77).length), _1909_v92);
        _1910_v93 = _nw347;
        let _1911_v94;
        _1911_v94 = _dafny.Set.fromElements(_1910_v93);
        let _1912_v95;
        _1912_v95 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((((_1871_v66).contains(p0)) ? ((_1871_v66).get(p0)) : (new BigNumber((_1894_v81).length)))),_1911_v94);
        _1912_v95 = _1912_v95;
        let _index321 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1879_v72).length));
        (_1879_v72)[_index321] = p0;
        let _1913_v96;
        _1913_v96 = _dafny.MultiSet.fromElements((_1910_v93).f18);
        let _1914_v97;
        _1914_v97 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), ((_1915_v65) => function (_1916_i5) {
          return new BigNumber((_1915_v65).length);
        })(_1870_v65)), _dafny.Seq.of(new BigNumber(443), (_1910_v93).f19), _module.__default.fm12(_1913_v96, _1885_cf14, globalState));
        let _1917_v98;
        _1917_v98 = _dafny.Set.fromElements((_this).f16);
        let _index322 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1879_v72).length));
        let _rhs257 = !((((_1870_v65).contains(new BigNumber(831))) ? ((_1870_v65).get(new BigNumber(831))) : (!(p0))));
        let _rhs258 = (_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(752)), ((_1918_v96) => function (_1919_i6) {
          return new BigNumber((_1918_v96).cardinality());
        })(_1913_v96)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(266)), ((_1920_v93) => function (_1921_i7) {
          return (_1920_v93).f19;
        })(_1910_v93)))).IsSubsetOf((_1914_v97).Intersect(_1914_v97));
        let _rhs259 = (_1917_v98).IsSubsetOf((_1917_v98).Intersect(_1917_v98));
        let _lhs229 = globalState;
        let _lhs230 = globalState;
        let _lhs231 = _1879_v72;
        let _lhs232 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_1879_v72).length));
        _lhs229.f1 = _rhs257;
        _lhs230.f1 = _rhs258;
        _lhs231[_lhs232] = _rhs259;
      } else if (_source17.is_DC12) {
        let _1922___mcc_h12 = (_source17).cf15;
        let _1923___mcc_h13 = (_source17).cf16;
        let _1924___mcc_h14 = (_source17).cf17;
        let _1925_cf17 = _1924___mcc_h14;
        let _1926_cf16 = _1923___mcc_h13;
        let _1927_cf15 = _1922___mcc_h12;
        (globalState).f1 = !(!(_1926_cf16));
        let _1928_v99;
        let _1929_v100;
        let _out14;
        let _out15;
        let _outcollector5 = (_1878_v71).m1(globalState);
        _out14 = _outcollector5[0];
        _out15 = _outcollector5[1];
        _1928_v99 = _out14;
        _1929_v100 = _out15;
        let _1930_v101;
        _1930_v101 = _dafny.Seq.UnicodeFromString("iohscrqtc");
        let _1931_v102;
        _1931_v102 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,new BigNumber((_1930_v101).length));
        if ((new BigNumber((_1931_v102).length)).isEqualTo(_1927_cf15)) {
          (globalState).f6 = (_this).f16;
          let _1932_v104;
          let _init52 = function (_1933_i8) {
            return _module.D0.create_DC2(_module.D0.create_DC1(new BigNumber((function () {
  let _coll72 = new _dafny.Set();
  for (const _compr_72 of _dafny.IntegerRange(new BigNumber(-897), new BigNumber(-991))) {
    let _1934_v103 = _compr_72;
    if (((new BigNumber(-897)).isLessThanOrEqualTo(_1934_v103)) && ((_1934_v103).isLessThan(new BigNumber(-991)))) {
      _coll72.add((_1934_v103).plus((_this).f16));
    }
  }
  return _coll72;
}()).length)));
          };
          let _nw348 = Array((new BigNumber(4)).toNumber());
          for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw348.length); _i0_52++) {
            _nw348[_i0_52] = _init52(new BigNumber(_i0_52));
          }
          _1932_v104 = _nw348;
          let _1935_v105;
          _1935_v105 = ((p0) ? (_1932_v104) : (_1932_v104));
          _1935_v105 = _1935_v105;
          _1927_cf15 = (_1927_cf15).minus(new BigNumber((_1871_v66).length));
          _1928_v99 = _1928_v99;
          let _1936_v106;
          _1936_v106 = _dafny.MultiSet.fromElements(_1926_cf16);
          let _1937_v107;
          _1937_v107 = _dafny.Set.fromElements(_1926_cf16);
          let _1938_v108;
          _1938_v108 = _dafny.Seq.of(new BigNumber(533));
          let _1939_v109;
          let _nw349 = Array((new BigNumber(20)).toNumber());
          _nw349[(_dafny.ZERO).toNumber()] = _1927_cf15;
          _nw349[(_dafny.ONE).toNumber()] = (_module.__default.fm1(!(p0), globalState)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_1925_cf17, _1925_cf17)).cardinality()));
          _nw349[(new BigNumber(2)).toNumber()] = (_this).f16;
          _nw349[(new BigNumber(3)).toNumber()] = _1928_v99;
          _nw349[(new BigNumber(4)).toNumber()] = (_1927_cf15).minus(_1927_cf15);
          _nw349[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.of(_1927_cf15)).length);
          _nw349[(new BigNumber(6)).toNumber()] = _1927_cf15;
          _nw349[(new BigNumber(7)).toNumber()] = _1928_v99;
          _nw349[(new BigNumber(8)).toNumber()] = (_this).f16;
          _nw349[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(_1927_cf15, new BigNumber(551));
          _nw349[(new BigNumber(10)).toNumber()] = _1928_v99;
          _nw349[(new BigNumber(11)).toNumber()] = (new BigNumber((_1936_v106).cardinality())).plus(new BigNumber(809));
          _nw349[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt(_1927_cf15, (_this).f16);
          _nw349[(new BigNumber(13)).toNumber()] = new BigNumber((_1937_v107).length);
          _nw349[(new BigNumber(14)).toNumber()] = (_this).f16;
          _nw349[(new BigNumber(15)).toNumber()] = _module.__default.fm1(_1929_v100, globalState);
          _nw349[(new BigNumber(16)).toNumber()] = ((p0) ? (_1927_cf15) : (new BigNumber((_1938_v108).length)));
          _nw349[(new BigNumber(17)).toNumber()] = (_1927_cf15).minus(_1927_cf15);
          _nw349[(new BigNumber(18)).toNumber()] = _1928_v99;
          _nw349[(new BigNumber(19)).toNumber()] = _1928_v99;
          _1939_v109 = _nw349;
          let _index323 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_1939_v109).length));
          (_1939_v109)[_index323] = _1928_v99;
          let _1940_v110;
          let _nw350 = Array((new BigNumber(14)).toNumber()).fill(_module.D15.Default());
          _1940_v110 = _nw350;
          let _1941_v111;
          _1941_v111 = _dafny.Map.Empty.slice().updateUnsafe(!(_1929_v100),_1940_v110);
          let _1942_v112;
          _1942_v112 = _dafny.Map.Empty.slice().updateUnsafe(_1941_v111,_1928_v99);
          let _index324 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_1939_v109).length));
          (_1939_v109)[_index324] = ((_1926_cf16) ? ((((_1942_v112).contains(_1941_v111)) ? ((_1942_v112).get(_1941_v111)) : (new BigNumber((_1937_v107).length)))) : (_1927_cf15));
        } else {
          (globalState).f1 = !(p0);
          (globalState).f1 = !(_module.__default.fm2(globalState));
          (globalState).f14 = _1879_v72;
          let _1943_v113;
          _1943_v113 = _dafny.Set.fromElements(_1926_cf16);
          let _1944_v114;
          _1944_v114 = _dafny.Set.fromElements(_1943_v113);
          let _1945_v115;
          _1945_v115 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1928_v99,true),(_dafny.Set.fromElements(_1943_v113, _1943_v113)).IsSubsetOf(_1944_v114));
          _1945_v115 = (_1945_v115).update(_1870_v65, _1926_cf16);
          let _1946_v116;
          let _nw351 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _1946_v116 = _nw351;
          let _1947_v117;
          _1947_v117 = _dafny.MultiSet.fromElements(_1926_cf16, p0);
          let _index325 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1946_v116).length));
          (_1946_v116)[_index325] = new BigNumber((_1947_v117).cardinality());
          let _index326 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_1946_v116).length));
          (_1946_v116)[_index326] = _module.__default.safeDivisionInt((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1928_v99,_1927_cf15)).length)).plus(new BigNumber(154)), _1927_cf15);
        }
        let _1948_v118;
        let _nw352 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _1948_v118 = _nw352;
        let _index327 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_1948_v118).length));
        (_1948_v118)[_index327] = (new BigNumber(788)).plus(_1927_cf15);
        let _index328 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_1948_v118).length));
        (_1948_v118)[_index328] = (((_1931_v102).contains((((_1931_v102).contains((_this).f16)) ? ((_1931_v102).get((_this).f16)) : ((_dafny.ZERO).minus(_1927_cf15))))) ? ((_1931_v102).get((((_1931_v102).contains((_this).f16)) ? ((_1931_v102).get((_this).f16)) : ((_dafny.ZERO).minus(_1927_cf15))))) : (_1928_v99));
      } else if (_source17.is_DC10) {
        let _1949___mcc_h15 = (_source17).cf13;
        let _1950_cf13 = _1949___mcc_h15;
        (globalState).f6 = _module.__default.safeDivisionInt((_this).f16, (_this).f16);
        let _1951_v119;
        let _nw353 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1951_v119 = _nw353;
        let _1952_v120;
        _1952_v120 = _dafny.MultiSet.fromElements(p0, p0);
        let _index329 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1951_v119).length));
        (_1951_v119)[_index329] = (_1952_v120).Intersect(_dafny.MultiSet.fromElements(!(p0), p0, p0, p0));
        let _index330 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1951_v119).length));
        (_1951_v119)[_index330] = _1952_v120;
        let _1953_v121;
        _1953_v121 = _module.D0.create_DC1((_this).f16);
        let _source18 = _1953_v121;
        if (_source18.is_DC1) {
          let _1954___mcc_h17 = (_source18).cf1;
          let _1955_cf1 = _1954___mcc_h17;
          let _1956_v122;
          _1956_v122 = _module.D0.create_DC1(_1955_cf1);
          let _1957_v123;
          _1957_v123 = _module.D0.create_DC2(_1956_v122);
          let _pat_let_tv40 = _1956_v122;
          let _1958_v124;
          let _nw354 = Array((new BigNumber(13)).toNumber());
          _nw354[(_dafny.ZERO).toNumber()] = function (_pat_let65_0) {
            return function (_1959_dt__update__tmp_h1) {
              return function (_pat_let66_0) {
                return function (_1960_dt__update_hcf2_h0) {
                  return _module.D0.create_DC2(_1960_dt__update_hcf2_h0);
                }(_pat_let66_0);
              }(_pat_let_tv40);
            }(_pat_let65_0);
          }(_1957_v123);
          _nw354[(_dafny.ONE).toNumber()] = _1957_v123;
          _nw354[(new BigNumber(2)).toNumber()] = _1957_v123;
          _nw354[(new BigNumber(3)).toNumber()] = _1957_v123;
          _nw354[(new BigNumber(4)).toNumber()] = _1957_v123;
          _nw354[(new BigNumber(5)).toNumber()] = _1957_v123;
          _nw354[(new BigNumber(6)).toNumber()] = _module.D0.create_DC2(_module.D0.create_DC1((_this).f16));
          _nw354[(new BigNumber(7)).toNumber()] = _module.D0.create_DC2(_1956_v122);
          _nw354[(new BigNumber(8)).toNumber()] = _1957_v123;
          _nw354[(new BigNumber(9)).toNumber()] = _1957_v123;
          _nw354[(new BigNumber(10)).toNumber()] = _1957_v123;
          _nw354[(new BigNumber(11)).toNumber()] = _1957_v123;
          _nw354[(new BigNumber(12)).toNumber()] = _1957_v123;
          _1958_v124 = _nw354;
          let _1961_v125;
          let _nw355 = Array((new BigNumber(4)).toNumber());
          _nw355[(_dafny.ZERO).toNumber()] = new BigNumber(220);
          _nw355[(_dafny.ONE).toNumber()] = (_this).f16;
          _nw355[(new BigNumber(2)).toNumber()] = _1955_cf1;
          _nw355[(new BigNumber(3)).toNumber()] = _1955_cf1;
          _1961_v125 = _nw355;
          _module.__default.m0(_dafny.Set.fromElements((_this).f16, _1955_cf1, _1955_cf1, _1955_cf1), _1958_v124, _1961_v125, _dafny.Seq.UnicodeFromString("nruo"), globalState);
          let _1962_v126;
          _1962_v126 = _dafny.Seq.of(p0);
          let _1963_v127;
          _1963_v127 = _dafny.Seq.UnicodeFromString("mypoqw");
          let _1964_v128;
          _1964_v128 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_module.__default.safeDivisionInt(new BigNumber((_1962_v126).length), new BigNumber((_1963_v127).length)));
          _1964_v128 = (_1964_v128).update(_1955_cf1, (_this).f16);
          let _1965_v129;
          _1965_v129 = _dafny.MultiSet.fromElements((new BigNumber((_1963_v127).length)).minus(_1955_cf1));
          let _1966_v130;
          _1966_v130 = _dafny.Seq.of(_1965_v129, _1965_v129);
          let _1967_v131;
          _1967_v131 = _module.D11.create_DC32(p0, (_this).f16, p0);
          _1965_v129 = (_1966_v130)[_module.__default.safeIndex((_1967_v131).dtor_cf60, new BigNumber((_1966_v130).length))];
          _1955_cf1 = (_this).f16;
        } else if (_source18.is_DC0) {
          let _1968___mcc_h18 = (_source18).cf0;
          let _1969_cf0 = _1968___mcc_h18;
          let _1970_v132;
          _1970_v132 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,((p0) ? ((_this).f16) : ((_this).f16)));
          (globalState).f6 = new BigNumber((_1970_v132).length);
          (globalState).f1 = _1969_cf0;
          let _1971_v133;
          _1971_v133 = _dafny.Seq.of(!(_1969_cf0), _1969_cf0, _1969_cf0);
          let _1972_v134;
          _1972_v134 = _module.D13.create_DC37(!(!((_1971_v133)[_module.__default.safeIndex((_this).f16, new BigNumber((_1971_v133).length))])) || (_1969_cf0), (_1952_v120).IsProperSubsetOf((_1951_v119)[_module.__default.safeIndex(new BigNumber(837), new BigNumber((_1951_v119).length))]));
          _1972_v134 = ((p0) ? (_1972_v134) : (_1972_v134));
          (globalState).f1 = p0;
        } else {
          let _1973___mcc_h19 = (_source18).cf2;
          let _1974_cf2 = _1973___mcc_h19;
          let _1975_v135;
          _1975_v135 = _dafny.Seq.of(_module.__default.safeModuloInt((_this).f16, (_this).f16));
          let _1976_v136;
          _1976_v136 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_dafny.Seq.Concat(_1975_v135, _dafny.Seq.of((_dafny.ZERO).minus((_this).f16))));
          _1975_v135 = (((_1976_v136).contains((_this).f16)) ? ((_1976_v136).get((_this).f16)) : (_1975_v135));
          (globalState).f6 = (_this).f16;
          (globalState).f6 = (_dafny.ZERO).minus((_this).f16);
          (globalState).f1 = p0;
        }
        if (p0) {
          (globalState).f1 = false;
          let _1977_v137;
          _1977_v137 = _dafny.Seq.of((_this).f16, (_this).f16, (_this).f16);
          let _1978_v138;
          let _nw356 = Array((new BigNumber(23)).toNumber()).fill(_module.D0.Default());
          _1978_v138 = _nw356;
          let _1979_v139;
          let _init53 = function (_1980_i9) {
            return (_1980_i9).minus((_dafny.ZERO).minus((_this).f16));
          };
          let _nw357 = Array((new BigNumber(19)).toNumber());
          for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw357.length); _i0_53++) {
            _nw357[_i0_53] = _init53(new BigNumber(_i0_53));
          }
          _1979_v139 = _nw357;
          _module.__default.m0(_dafny.Set.fromElements((_this).f16, new BigNumber((_1977_v137).length), (_this).f16, new BigNumber(687)), _1978_v138, _1979_v139, _dafny.Seq.UnicodeFromString("ugbdhjeb"), globalState);
          let _1981_v140;
          _1981_v140 = _dafny.Seq.UnicodeFromString("vgycqsugd");
          let _1982_v141;
          _1982_v141 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f16,p0)).length), new BigNumber((_1981_v140).length));
          (globalState).f1 = ((_this).f16).isLessThan(new BigNumber((_1982_v141).cardinality()));
          let _1983_v142;
          let _nw358 = new _module.C4();
          _nw358.__ctor();
          _1983_v142 = _nw358;
          _1870_v65 = (_1870_v65).update((_this).f16, p0);
        } else {
          let _1984_v143;
          _1984_v143 = new _dafny.CodePoint('x'.codePointAt(0));
          let _1985_v144;
          _1985_v144 = _dafny.Seq.UnicodeFromString("ydawcir");
          let _1986_v145;
          _1986_v145 = _dafny.Seq.of(new BigNumber(361));
          let _1987_v146;
          let _nw359 = Array((new BigNumber(18)).toNumber());
          _nw359[(_dafny.ZERO).toNumber()] = (_this).f16;
          _nw359[(_dafny.ONE).toNumber()] = (new BigNumber((_1985_v144).length)).multipliedBy((_this).f16);
          _nw359[(new BigNumber(2)).toNumber()] = (_this).f16;
          _nw359[(new BigNumber(3)).toNumber()] = new BigNumber((_1986_v145).length);
          _nw359[(new BigNumber(4)).toNumber()] = ((_this).f16).plus((_this).f16);
          _nw359[(new BigNumber(5)).toNumber()] = (_this).f16;
          _nw359[(new BigNumber(6)).toNumber()] = (_this).f16;
          _nw359[(new BigNumber(7)).toNumber()] = (_this).f16;
          _nw359[(new BigNumber(8)).toNumber()] = (_this).f16;
          _nw359[(new BigNumber(9)).toNumber()] = (_this).f16;
          _nw359[(new BigNumber(10)).toNumber()] = (_1986_v145)[_module.__default.safeIndex((_this).f16, new BigNumber((_1986_v145).length))];
          _nw359[(new BigNumber(11)).toNumber()] = (_this).f16;
          _nw359[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_this).f16);
          _nw359[(new BigNumber(13)).toNumber()] = (_this).f16;
          _nw359[(new BigNumber(14)).toNumber()] = new BigNumber(619);
          _nw359[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt((_this).f16, new BigNumber(-314));
          _nw359[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1871_v66).length));
          _nw359[(new BigNumber(17)).toNumber()] = (_this).f16;
          _1987_v146 = _nw359;
          let _1988_v147;
          _1988_v147 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,(_this).f16);
          let _rhs260 = _1984_v143;
          let _rhs261 = new BigNumber(534);
          let _rhs262 = _1987_v146;
          let _rhs263 = (((_1988_v147).contains((_this).f16)) ? ((_1988_v147).get((_this).f16)) : ((new BigNumber((_1985_v144).length)).plus((_this).f16)));
          let _rhs264 = (_1986_v145)[_module.__default.safeIndex((_this).f16, new BigNumber((_1986_v145).length))];
          let _lhs233 = globalState;
          let _lhs234 = globalState;
          let _lhs235 = globalState;
          let _lhs236 = globalState;
          _1984_v143 = _rhs260;
          _lhs233.f6 = _rhs261;
          _lhs234.f9 = _rhs262;
          _lhs235.f6 = _rhs263;
          _lhs236.f6 = _rhs264;
          let _1989_v148;
          let _nw360 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
          _1989_v148 = _nw360;
          let _init54 = ((_1990_v147) => function (_1991_i10) {
            return (_1990_v147).Merge(_1990_v147);
          })(_1988_v147);
          let _nw361 = Array((new BigNumber(21)).toNumber());
          for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw361.length); _i0_54++) {
            _nw361[_i0_54] = _init54(new BigNumber(_i0_54));
          }
          _1989_v148 = _nw361;
          let _index331 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_1879_v72).length));
          (_1879_v72)[_index331] = p0;
          let _index332 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_1879_v72).length));
          (_1879_v72)[_index332] = p0;
          _1985_v144 = _module.__default.fm21(globalState);
          let _1992_v149;
          _1992_v149 = _dafny.MultiSet.fromElements((_module.D16.create_DC45((_this).f16, _1984_v143, (_this).f16, (_1879_v72)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_1879_v72).length))])).dtor_cf78);
          (globalState).f6 = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1985_v144, _module.__default.safeIndex(new BigNumber(((_1992_v149).update(_module.__default.fm1(p0, globalState), _module.__default.abs((_this).f16))).cardinality()), new BigNumber((_1985_v144).length)), _1984_v143), _module.__default.fm21(globalState))).length)).minus(new BigNumber((_1871_v66).length));
        }
      } else {
        let _1993___mcc_h16 = (_source17).cf18;
        let _1994_cf18 = _1993___mcc_h16;
        (globalState).f1 = p0;
        let _index333 = _module.__default.safeIndex(new BigNumber(270), new BigNumber((_1882_v73).length));
        (_1882_v73)[_index333] = _1879_v72;
        let _index334 = _module.__default.safeIndex(new BigNumber(270), new BigNumber((_1882_v73).length));
        (_1882_v73)[_index334] = _1879_v72;
        (globalState).f1 = p0;
        if (p0) {
          let _1995_v150;
          _1995_v150 = new _dafny.CodePoint('t'.codePointAt(0));
          let _1996_v151;
          _1996_v151 = _dafny.MultiSet.fromElements(_1995_v150, _1995_v150);
          let _1997_v152;
          _1997_v152 = _dafny.Seq.of((_this).f16, (_this).f16);
          let _1998_v153;
          _1998_v153 = _dafny.Seq.of((_1997_v152)[_module.__default.safeIndex((_this).f16, new BigNumber((_1997_v152).length))]);
          let _1999_v154;
          let _nw362 = Array((new BigNumber(6)).toNumber());
          _nw362[(_dafny.ZERO).toNumber()] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(565)), function (_2000_i11) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })).length)).multipliedBy((_this).f16);
          _nw362[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_this).f16);
          _nw362[(new BigNumber(2)).toNumber()] = (_this).f16;
          _nw362[(new BigNumber(3)).toNumber()] = (_this).f16;
          _nw362[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(-430), (_this).f16);
          _nw362[(new BigNumber(5)).toNumber()] = (new BigNumber(((_1996_v151).update(_1995_v150, _module.__default.abs(new BigNumber((_1997_v152).length)))).cardinality())).plus(_module.__default.fm1(p0, globalState));
          _1999_v154 = _nw362;
          let _index335 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_1999_v154).length));
          (_1999_v154)[_index335] = (_this).f16;
          let _index336 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_1999_v154).length));
          (_1999_v154)[_index336] = _module.__default.safeDivisionInt((_this).f16, (_this).f16);
          let _2001_v155;
          _2001_v155 = _dafny.Seq.UnicodeFromString("gg");
          (globalState).f6 = _module.__default.safeModuloInt(new BigNumber((_1870_v65).length), ((_1999_v154)[_module.__default.safeIndex(new BigNumber(771), new BigNumber((_1999_v154).length))]).multipliedBy(new BigNumber((_2001_v155).length)));
          (globalState).f6 = (((false) ? ((_1999_v154)[_module.__default.safeIndex(new BigNumber(771), new BigNumber((_1999_v154).length))]) : ((_1999_v154)[_module.__default.safeIndex(new BigNumber(771), new BigNumber((_1999_v154).length))]))).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("gunwfodep")).length));
          let _2002_v156;
          let _2003_v157;
          let _out16;
          let _out17;
          let _outcollector6 = (_1878_v71).m1(globalState);
          _out16 = _outcollector6[0];
          _out17 = _outcollector6[1];
          _2002_v156 = _out16;
          _2003_v157 = _out17;
          let _index337 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_1999_v154).length));
          let _rhs265 = _2003_v157;
          let _rhs266 = (_this).f16;
          let _lhs237 = globalState;
          let _lhs238 = _1999_v154;
          let _lhs239 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_1999_v154).length));
          _lhs237.f1 = _rhs265;
          _lhs238[_lhs239] = _rhs266;
        } else {
          let _2004_v158;
          _2004_v158 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _2005_v159;
          _2005_v159 = _dafny.Map.Empty.slice().updateUnsafe(!((((_2004_v158).contains(false)) ? ((_2004_v158).get(false)) : (p0))),p0);
          _2004_v158 = (_2005_v159).Merge(_2005_v159);
          let _2006_v160;
          _2006_v160 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,new BigNumber((_dafny.Seq.of(p0, p0, true)).length));
          _2006_v160 = (_2006_v160).update((_this).f16, (_this).f16);
          let _2007_v161;
          _2007_v161 = new _dafny.CodePoint('b'.codePointAt(0));
          _2007_v161 = _2007_v161;
          let _2008_v162;
          _2008_v162 = _dafny.Map.Empty.slice().updateUnsafe(_1870_v65,(_dafny.ZERO).minus((_this).f16));
          let _2009_v163;
          let _nw363 = new _module.C3();
          _nw363.__ctor(false, (_this).f16, _2008_v162);
          _2009_v163 = _nw363;
          (globalState).f6 = ((_2009_v163).f19).multipliedBy((_this).f16);
        }
      }
      return;
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.D0.create_DC2(_module.D0.create_DC0(true));
    };
    fm9(p0, globalState) {
      let _this = this;
      return new BigNumber(995);
    };
    fm10(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(!((new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("tbhi")).length)))).cardinality())).isLessThanOrEqualTo(new BigNumber(720))) || ((true) === ((_module.D0.create_DC0(false)).dtor_cf0)));
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _2010_v0;
      _2010_v0 = new BigNumber(277);
      let _2011_v1;
      _2011_v1 = false;
      let _2012_v2;
      _2012_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2011_v1,_2011_v1);
      let _2013_v3;
      let _nw364 = Array((new BigNumber(13)).toNumber());
      _nw364[(_dafny.ZERO).toNumber()] = (_2010_v0).plus((_dafny.ZERO).minus(_2010_v0));
      _nw364[(_dafny.ONE).toNumber()] = _2010_v0;
      _nw364[(new BigNumber(2)).toNumber()] = (new BigNumber(317)).minus(_2010_v0);
      _nw364[(new BigNumber(3)).toNumber()] = _2010_v0;
      _nw364[(new BigNumber(4)).toNumber()] = _2010_v0;
      _nw364[(new BigNumber(5)).toNumber()] = _2010_v0;
      _nw364[(new BigNumber(6)).toNumber()] = _2010_v0;
      _nw364[(new BigNumber(7)).toNumber()] = _module.__default.fm1(!(_2011_v1), globalState);
      _nw364[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_2010_v0);
      _nw364[(new BigNumber(9)).toNumber()] = _2010_v0;
      _nw364[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_2010_v0, _2010_v0));
      _nw364[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(_2010_v0, _2010_v0);
      _nw364[(new BigNumber(12)).toNumber()] = (new BigNumber(((_module.D1.create_DC3(_2012_v2)).dtor_cf3).length)).multipliedBy(new BigNumber(-379));
      _2013_v3 = _nw364;
      let _2014_v4;
      let _nw365 = Array((new BigNumber(11)).toNumber());
      _nw365[(_dafny.ZERO).toNumber()] = _2011_v1;
      _nw365[(_dafny.ONE).toNumber()] = _2011_v1;
      _nw365[(new BigNumber(2)).toNumber()] = _2011_v1;
      _nw365[(new BigNumber(3)).toNumber()] = false;
      _nw365[(new BigNumber(4)).toNumber()] = _2011_v1;
      _nw365[(new BigNumber(5)).toNumber()] = _2011_v1;
      _nw365[(new BigNumber(6)).toNumber()] = _2011_v1;
      _nw365[(new BigNumber(7)).toNumber()] = _2011_v1;
      _nw365[(new BigNumber(8)).toNumber()] = _2011_v1;
      _nw365[(new BigNumber(9)).toNumber()] = _2011_v1;
      _nw365[(new BigNumber(10)).toNumber()] = !(!(_2011_v1));
      _2014_v4 = _nw365;
      let _2015_v5;
      _2015_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2011_v1,_2014_v4);
      let _index338 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_2013_v3).length));
      (_2013_v3)[_index338] = new BigNumber(((_2015_v5).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2011_v1,_2014_v4))).length);
      let _index339 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_2013_v3).length));
      (_2013_v3)[_index339] = (_2010_v0).multipliedBy(_module.__default.safeModuloInt(_2010_v0, new BigNumber(932)));
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2013_v3).length))) {
        let _2016_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2016_i0)) && ((_2016_i0).isLessThan(new BigNumber((_2013_v3).length))))) {
          (_2013_v3)[(_2016_i0)] = (_2016_i0).multipliedBy(new BigNumber(-161));
        }
      }
      let _2017_v6;
      _2017_v6 = _module.D0.create_DC0(_2011_v1);
      let _pat_let_tv41 = _2013_v3;
      let _pat_let_tv42 = _2013_v3;
      _2010_v0 = function (_source19) {
        if (_source19.is_DC1) {
          let _2018___mcc_h0 = (_source19).cf1;
          let _2019_cf1 = _2018___mcc_h0;
          return _2019_cf1;
        } else if (_source19.is_DC0) {
          let _2020___mcc_h1 = (_source19).cf0;
          let _2021_cf0 = _2020___mcc_h1;
          return (_pat_let_tv42)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_pat_let_tv41).length))];
        } else {
          let _2022___mcc_h2 = (_source19).cf2;
          let _2023_cf2 = _2022___mcc_h2;
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("krasuv")).length));
        }
      }(_2017_v6);
      let _2024_v7;
      _2024_v7 = _dafny.Map.Empty.slice().updateUnsafe((_2013_v3)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_2013_v3).length))],true);
      let _2025_v8;
      _2025_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2024_v7,(_2013_v3)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_2013_v3).length))]);
      let _2026_v9;
      let _nw366 = new _module.C3();
      _nw366.__ctor(_2011_v1, _2010_v0, _2025_v8);
      _2026_v9 = _nw366;
      _2010_v0 = _2010_v0;
      let _nw367 = Array((new BigNumber(13)).toNumber()).fill(false);
      (globalState).f14 = _nw367;
      r0 = (_2013_v3)[_module.__default.safeIndex(new BigNumber(8), new BigNumber((_2013_v3).length))];
      r1 = _2011_v1;
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let _2027_v2;
      let _init55 = ((_2028_p0) => function (_2029_i0) {
        return (_dafny.MultiSet.fromElements(new BigNumber(729), new BigNumber(484), (_dafny.ZERO).minus(new BigNumber((function () {
          let _coll73 = new _dafny.Map();
          for (const _compr_73 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll74 = new _dafny.Set();
            for (const _compr_74 of _dafny.IntegerRange(new BigNumber(9), new BigNumber(452))) {
              let _2030_v1 = _compr_74;
              if (((new BigNumber(9)).isLessThanOrEqualTo(_2030_v1)) && ((_2030_v1).isLessThan(new BigNumber(452)))) {
                _coll74.add(_module.__default.safeModuloInt(_2030_v1, new BigNumber(466)));
              }
            }
            return _coll74;
          }()).length),_2028_p0)).Keys.Elements) {
            let _2031_v0 = _compr_73;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
              let _coll75 = new _dafny.Set();
              for (const _compr_75 of _dafny.IntegerRange(new BigNumber(9), new BigNumber(452))) {
                let _2032_v1 = _compr_75;
                if (((new BigNumber(9)).isLessThanOrEqualTo(_2032_v1)) && ((_2032_v1).isLessThan(new BigNumber(452)))) {
                  _coll75.add(_module.__default.safeModuloInt(_2032_v1, new BigNumber(466)));
                }
              }
              return _coll75;
            }()).length),_2028_p0)).contains(_2031_v0)) {
              _coll73.push([(_2031_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cc")).length)),new BigNumber((_dafny.MultiSet.fromElements(_2028_p0)).cardinality())]);
            }
          }
          return _coll73;
        }()).length)))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-753), new BigNumber(-323))));
      })(p0);
      let _nw368 = Array((new BigNumber(15)).toNumber());
      for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw368.length); _i0_55++) {
        _nw368[_i0_55] = _init55(new BigNumber(_i0_55));
      }
      _2027_v2 = _nw368;
      let _2033_v3;
      _2033_v3 = new BigNumber(495);
      let _2034_v4;
      _2034_v4 = _dafny.MultiSet.fromElements(_2033_v3);
      let _index340 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2027_v2).length));
      (_2027_v2)[_index340] = _2034_v4;
      let _2035_v5;
      _2035_v5 = new _dafny.CodePoint('r'.codePointAt(0));
      let _2036_v6;
      let _nw369 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _2036_v6 = _nw369;
      let _2037_v7;
      _2037_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2033_v3,_2033_v3);
      let _2038_v8;
      _2038_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2037_v7).length),_2033_v3);
      let _index341 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_2036_v6).length));
      (_2036_v6)[_index341] = new BigNumber((_2038_v8).length);
      let _index342 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2027_v2).length));
      let _index343 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_2036_v6).length));
      let _rhs267 = (_2034_v4).Difference(_2034_v4);
      let _rhs268 = p0;
      let _rhs269 = _2035_v5;
      let _rhs270 = _2033_v3;
      let _lhs240 = _2027_v2;
      let _lhs241 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2027_v2).length));
      let _lhs242 = globalState;
      let _lhs243 = _2036_v6;
      let _lhs244 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_2036_v6).length));
      _lhs240[_lhs241] = _rhs267;
      _lhs242.f1 = _rhs268;
      _2035_v5 = _rhs269;
      _lhs243[_lhs244] = _rhs270;
      let _2039_v9;
      let _nw370 = Array((new BigNumber(3)).toNumber()).fill(false);
      _2039_v9 = _nw370;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2039_v9).length))) {
        let _2040_i1 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2040_i1)) && ((_2040_i1).isLessThan(new BigNumber((_2039_v9).length))))) {
          (_2039_v9)[(_2040_i1)] = (p0) && (p0);
        }
      }
      let _index344 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_2036_v6).length));
      (_2036_v6)[_index344] = (_2036_v6)[_module.__default.safeIndex(new BigNumber(715), new BigNumber((_2036_v6).length))];
      let _2041_i2;
      _2041_i2 = _dafny.ZERO;
      L14: {
        while (p0) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2041_i2)) {
              break L14;
            }
            _2041_i2 = (_2041_i2).plus(_dafny.ONE);
            let _2042_v10;
            _2042_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2033_v3,p0);
            let _2043_v11;
            _2043_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2042_v10,(_dafny.ZERO).minus((_2036_v6)[_module.__default.safeIndex(new BigNumber(715), new BigNumber((_2036_v6).length))]));
            let _2044_v12;
            let _nw371 = new _module.C2();
            _nw371.__ctor((new BigNumber(837)).plus(new BigNumber(-800)), _2033_v3, _2043_v11);
            _2044_v12 = _nw371;
            if (p0) {
              let _2045_v13;
              _2045_v13 = _2039_v9;
              let _2046_v14;
              _2046_v14 = _dafny.Seq.of(_2039_v9, _2039_v9, (_2045_v13), _2039_v9, _2039_v9);
              let _2047_v15;
              let _nw372 = Array((new BigNumber(11)).toNumber());
              _nw372[(_dafny.ZERO).toNumber()] = (_2046_v14)[_module.__default.safeIndex(_2044_v12.f21, new BigNumber((_2046_v14).length))];
              _nw372[(_dafny.ONE).toNumber()] = _2039_v9;
              _nw372[(new BigNumber(2)).toNumber()] = _2039_v9;
              _nw372[(new BigNumber(3)).toNumber()] = _2039_v9;
              _nw372[(new BigNumber(4)).toNumber()] = _2039_v9;
              _nw372[(new BigNumber(5)).toNumber()] = _2039_v9;
              _nw372[(new BigNumber(6)).toNumber()] = (_2046_v14)[_module.__default.safeIndex(_2033_v3, new BigNumber((_2046_v14).length))];
              _nw372[(new BigNumber(7)).toNumber()] = _2039_v9;
              _nw372[(new BigNumber(8)).toNumber()] = _2039_v9;
              _nw372[(new BigNumber(9)).toNumber()] = _2039_v9;
              _nw372[(new BigNumber(10)).toNumber()] = _2039_v9;
              _2047_v15 = _nw372;
              _2047_v15 = _2047_v15;
              let _2048_v16;
              let _nw373 = new _module.C5();
              _nw373.__ctor(new BigNumber((_dafny.Set.fromElements(_2036_v6)).length));
              _2048_v16 = _nw373;
              (globalState).f6 = _2044_v12.f21;
              let _index345 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_2039_v9).length));
              (_2039_v9)[_index345] = (_2033_v3).isLessThan(new BigNumber(114));
              let _index346 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_2039_v9).length));
              (_2039_v9)[_index346] = p0;
              (globalState).f1 = p0;
            } else {
              let _index347 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_2036_v6).length));
              (_2036_v6)[_index347] = _module.__default.safeModuloInt(_2044_v12.f21, _2033_v3);
              let _2049_v17;
              _2049_v17 = _dafny.Seq.UnicodeFromString("idh");
              _2049_v17 = _2049_v17;
              (globalState).f1 = p0;
              let _2050_v18;
              _2050_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2033_v3,_2049_v17);
              let _rhs271 = _module.__default.fm38(_module.__default.safeModuloInt(_module.__default.fm1(p0, globalState), _module.__default.fm1(p0, globalState)), p0, p0, !(p0), globalState);
              let _rhs272 = new BigNumber(234);
              let _lhs245 = _2044_v12;
              _2050_v18 = _rhs271;
              _lhs245.f21 = _rhs272;
              let _2051_v19;
              let _nw374 = Array((new BigNumber(14)).toNumber());
              _nw374[(_dafny.ZERO).toNumber()] = ((p0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(68)), ((_2052_v5) => function (_2053_i3) {
                return _2052_v5;
              })(_2035_v5))) : (_dafny.Seq.UnicodeFromString("ckh")));
              _nw374[(_dafny.ONE).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(2)).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(3)).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(4)).toNumber()] = ((p0) ? (_dafny.Seq.UnicodeFromString("xyjs")) : (_2049_v17));
              _nw374[(new BigNumber(5)).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(6)).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(7)).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(8)).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(9)).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(10)).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(11)).toNumber()] = _2049_v17;
              _nw374[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(686)), function (_2054_i4) {
                return new _dafny.CodePoint('s'.codePointAt(0));
              });
              _nw374[(new BigNumber(13)).toNumber()] = _2049_v17;
              _2051_v19 = _nw374;
              let _index348 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_2051_v19).length));
              (_2051_v19)[_index348] = (((_2050_v18).contains(_2044_v12.f21)) ? ((_2050_v18).get(_2044_v12.f21)) : (_2049_v17));
              let _index349 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_2051_v19).length));
              (_2051_v19)[_index349] = _2049_v17;
            }
            _2042_v10 = (_2042_v10).update(_2033_v3, p0);
            let _2055_v20;
            let _nw375 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
            _2055_v20 = _nw375;
            let _2056_v21;
            _2056_v21 = _dafny.Seq.UnicodeFromString("pvfjavan");
            let _index350 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_2055_v20).length));
            (_2055_v20)[_index350] = (_dafny.MultiSet.fromElements(_2036_v6, _2036_v6)).update(_2036_v6, _module.__default.abs(new BigNumber((_2056_v21).length)));
            let _2057_v22;
            _2057_v22 = _dafny.Seq.of(p0);
            let _2058_v23;
            _2058_v23 = _dafny.MultiSet.fromElements(_2036_v6);
            let _index351 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_2055_v20).length));
            let _rhs273 = (_module.__default.fm44(globalState)).dtor_cf65;
            let _rhs274 = _2058_v23;
            let _rhs275 = p0;
            let _rhs276 = _2057_v22;
            let _lhs246 = globalState;
            let _lhs247 = _2055_v20;
            let _lhs248 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_2055_v20).length));
            let _lhs249 = globalState;
            _lhs246.f1 = _rhs273;
            _lhs247[_lhs248] = _rhs274;
            _lhs249.f1 = _rhs275;
            _2057_v22 = _rhs276;
          }
        }
      }
      let _index352 = _module.__default.safeIndex(new BigNumber(715), new BigNumber((_2036_v6).length));
      (_2036_v6)[_index352] = (_module.__default.safeDivisionInt(_2033_v3, new BigNumber(880))).plus(_module.__default.fm1(p0, globalState));
      let _2059_v24;
      _2059_v24 = _dafny.Seq.UnicodeFromString("o");
      _2059_v24 = _module.__default.fm21(globalState);
      return;
    }
    m8(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _2060_v0;
      let _nw376 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
      _2060_v0 = _nw376;
      let _2061_v1;
      _2061_v1 = _module.D20.create_DC49(_2060_v0);
      let _pat_let_tv43 = _2060_v0;
      let _pat_let_tv44 = _2060_v0;
      let _2062_v2;
      let _nw377 = Array((new BigNumber(8)).toNumber());
      _nw377[(_dafny.ZERO).toNumber()] = _module.D20.create_DC49(_2060_v0);
      _nw377[(_dafny.ONE).toNumber()] = _2061_v1;
      _nw377[(new BigNumber(2)).toNumber()] = function (_pat_let67_0) {
        return function (_2063_dt__update__tmp_h0) {
          return function (_pat_let68_0) {
            return function (_2064_dt__update_hcf85_h0) {
              return _module.D20.create_DC49(_2064_dt__update_hcf85_h0);
            }(_pat_let68_0);
          }(_pat_let_tv43);
        }(_pat_let67_0);
      }(_2061_v1);
      _nw377[(new BigNumber(3)).toNumber()] = _2061_v1;
      _nw377[(new BigNumber(4)).toNumber()] = function (_pat_let69_0) {
        return function (_2065_dt__update__tmp_h1) {
          return function (_pat_let70_0) {
            return function (_2066_dt__update_hcf85_h1) {
              return _module.D20.create_DC49(_2066_dt__update_hcf85_h1);
            }(_pat_let70_0);
          }(_pat_let_tv44);
        }(_pat_let69_0);
      }(_module.D20.create_DC49(_2060_v0));
      _nw377[(new BigNumber(5)).toNumber()] = _2061_v1;
      _nw377[(new BigNumber(6)).toNumber()] = _2061_v1;
      _nw377[(new BigNumber(7)).toNumber()] = _2061_v1;
      _2062_v2 = _nw377;
      let _index353 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_2062_v2).length));
      (_2062_v2)[_index353] = _2061_v1;
      let _index354 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_2062_v2).length));
      let _rhs277 = _2061_v1;
      let _rhs278 = _module.__default.fm2(globalState);
      let _lhs250 = _2062_v2;
      let _lhs251 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_2062_v2).length));
      let _lhs252 = globalState;
      _lhs250[_lhs251] = _rhs277;
      _lhs252.f1 = _rhs278;
      let _2067_v3;
      _2067_v3 = false;
      (globalState).f1 = _2067_v3;
      let _2068_v4;
      let _init56 = ((_2069_v3) => function (_2070_i0) {
        return _2069_v3;
      })(_2067_v3);
      let _nw378 = Array((new BigNumber(3)).toNumber());
      for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw378.length); _i0_56++) {
        _nw378[_i0_56] = _init56(new BigNumber(_i0_56));
      }
      _2068_v4 = _nw378;
      let _2071_v5;
      _2071_v5 = _2068_v4;
      let _2072_v6;
      _2072_v6 = _dafny.Seq.of(_2067_v3);
      let _2073_v7;
      _2073_v7 = _dafny.MultiSet.fromElements((_2072_v6)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("sqjejkkcx")).length), new BigNumber((_2072_v6).length))], _2067_v3);
      let _2074_v8;
      _2074_v8 = new BigNumber(-321);
      let _rhs279 = _2071_v5;
      let _rhs280 = (_2073_v7).IsSubsetOf(_2073_v7);
      let _rhs281 = _2074_v8;
      let _lhs253 = globalState;
      _2071_v5 = _rhs279;
      r0 = _rhs280;
      _lhs253.f6 = _rhs281;
      if (_2067_v3) {
        r0 = false;
        let _2075_v9;
        let _nw379 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2075_v9 = _nw379;
        let _index355 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_2075_v9).length));
        (_2075_v9)[_index355] = p0;
        let _index356 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_2075_v9).length));
        (_2075_v9)[_index356] = new _dafny.CodePoint('x'.codePointAt(0));
        let _2076_v10;
        let _nw380 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
        _2076_v10 = _nw380;
        let _2077_v11;
        _2077_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2076_v10,!(_2067_v3));
        _2077_v11 = (_2077_v11).update(_2076_v10, !(false) || (_2067_v3));
        let _2078_v12;
        _2078_v12 = _module.D14.create_DC38(_2076_v10);
        let _pat_let_tv45 = _2076_v10;
        _2078_v12 = function (_pat_let71_0) {
          return function (_2079_dt__update__tmp_h2) {
            return function (_pat_let72_0) {
              return function (_2080_dt__update_hcf67_h0) {
                return _module.D14.create_DC38(_2080_dt__update_hcf67_h0);
              }(_pat_let72_0);
            }(_pat_let_tv45);
          }(_pat_let71_0);
        }(_2078_v12);
        let _2081_v13;
        _2081_v13 = _dafny.Seq.of(_2074_v8, _2074_v8, _2074_v8);
        let _2082_v14;
        _2082_v14 = _dafny.Seq.UnicodeFromString("qwmvti");
        let _2083_v15;
        let _init57 = ((_2084_v8) => function (_2085_i1) {
          return _module.__default.safeModuloInt(_2085_i1, _2084_v8);
        })(_2074_v8);
        let _nw381 = Array((new BigNumber(29)).toNumber());
        for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw381.length); _i0_57++) {
          _nw381[_i0_57] = _init57(new BigNumber(_i0_57));
        }
        _2083_v15 = _nw381;
        let _2086_v16;
        _2086_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2083_v15,_2067_v3);
        _2081_v13 = _dafny.Seq.update(_2081_v13, _module.__default.safeIndex((_module.__default.fm1(true, globalState)).plus(new BigNumber((_2082_v14).length)), new BigNumber((_2081_v13).length)), _module.__default.fm1((((_2086_v16).contains(_2083_v15)) ? ((_2086_v16).get(_2083_v15)) : (_2067_v3)), globalState));
      } else {
        (globalState).f1 = _2067_v3;
        let _2087_v17;
        _2087_v17 = new _dafny.CodePoint('t'.codePointAt(0));
        let _2088_v18;
        _2088_v18 = _module.D5.create_DC12(_2074_v8, (_2072_v6)[_module.__default.safeIndex(_2074_v8, new BigNumber((_2072_v6).length))], p0);
        let _rhs282 = new BigNumber((_2072_v6).length);
        let _rhs283 = (_2088_v18).dtor_cf16;
        let _rhs284 = _2067_v3;
        let _rhs285 = p0;
        let _lhs254 = globalState;
        let _lhs255 = globalState;
        _lhs254.f6 = _rhs282;
        _2067_v3 = _rhs283;
        _lhs255.f1 = _rhs284;
        _2087_v17 = _rhs285;
        let _2089_v19;
        let _nw382 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _2089_v19 = _nw382;
        let _2090_v20;
        _2090_v20 = _dafny.Seq.UnicodeFromString("kbvvji");
        let _2091_v21;
        _2091_v21 = _dafny.Set.fromElements(_2090_v20);
        let _2092_v22;
        _2092_v22 = _dafny.Seq.of(_2074_v8);
        let _2093_v23;
        _2093_v23 = _dafny.Seq.of(new BigNumber((_2091_v21).length), new BigNumber((_2092_v22).length));
        let _2094_v24;
        _2094_v24 = _dafny.Map.Empty.slice().updateUnsafe((_2092_v22),_2074_v8);
        let _index357 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_2089_v19).length));
        (_2089_v19)[_index357] = new BigNumber((_2094_v24).length);
        let _index358 = _module.__default.safeIndex(new BigNumber(639), new BigNumber((_2089_v19).length));
        (_2089_v19)[_index358] = _2074_v8;
        let _2095_v25;
        _2095_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2067_v3,_2090_v20);
        _2095_v25 = (_dafny.Map.Empty.slice().updateUnsafe(_2067_v3,_dafny.Seq.Create(_module.__default.abs(new BigNumber(35)), ((_2096_p0) => function (_2097_i2) {
          return _2096_p0;
        })(p0)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("spti"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2067_v3,_2090_v20)));
        let _2098_v26;
        _2098_v26 = _module.D3.create_DC6(_2089_v19);
        let _2099_v27;
        _2099_v27 = _dafny.Seq.of(_2098_v26);
        let _2100_v28;
        let _nw383 = new _module.C0();
        _nw383.__ctor(_2099_v27, false);
        _2100_v28 = _nw383;
      }
      let _2101_v29;
      _2101_v29 = _dafny.Map.Empty.slice().updateUnsafe(_2067_v3,true);
      let _2102_v30;
      let _init58 = ((_2103_v8) => function (_2104_i3) {
        return (_2104_i3).plus(_2103_v8);
      })(_2074_v8);
      let _nw384 = Array((new BigNumber(9)).toNumber());
      for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw384.length); _i0_58++) {
        _nw384[_i0_58] = _init58(new BigNumber(_i0_58));
      }
      _2102_v30 = _nw384;
      let _2105_v31;
      _2105_v31 = _module.D3.create_DC6(_2102_v30);
      let _2106_v32;
      let _nw385 = new _module.C0();
      _nw385.__ctor(_dafny.Seq.of(_2105_v31, _2105_v31, _2105_v31), false);
      _2106_v32 = _nw385;
      let _2107_v33;
      _2107_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2106_v32,_2074_v8);
      let _2108_v34;
      _2108_v34 = _dafny.Seq.of(_2074_v8, (((_2107_v33).contains(_2106_v32)) ? ((_2107_v33).get(_2106_v32)) : (_2074_v8)));
      _2101_v29 = ((_2101_v29).Merge(_2101_v29)).update(_dafny.Seq.contains(_2108_v34, _2074_v8), _2067_v3);
      let _2109_v35;
      _2109_v35 = _dafny.Seq.of(_module.__default.fm19(globalState), _2072_v6);
      _2109_v35 = _2109_v35;
      r0 = (_2106_v32).f23;
      return r0;
    }
    m9(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2110_v0;
      _2110_v0 = _dafny.Seq.UnicodeFromString("twnmp");
      let _2111_v1;
      _2111_v1 = _dafny.Seq.of(new BigNumber((_2110_v0).length));
      let _2112_v2;
      _2112_v2 = new _dafny.CodePoint('q'.codePointAt(0));
      let _2113_v3;
      _2113_v3 = _module.D3.create_DC7(p0);
      let _2114_v4;
      let _nw386 = Array((new BigNumber(7)).toNumber()).fill(false);
      _2114_v4 = _nw386;
      let _2115_v5;
      _2115_v5 = true;
      let _2116_v6;
      _2116_v6 = _dafny.Map.Empty.slice().updateUnsafe(_2114_v4,_2115_v5);
      let _2117_v7;
      _2117_v7 = _module.D10.create_DC27(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2111_v1,_2112_v2)).length), _2110_v0, p0, _2113_v3, _2116_v6);
      let _2118_v8;
      _2118_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2110_v0).length),p0);
      let _2119_v9;
      _2119_v9 = _dafny.MultiSet.fromElements(p0);
      let _2120_v10;
      _2120_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2115_v5,p0);
      let _2121_v11;
      let _nw387 = Array((new BigNumber(21)).toNumber());
      _nw387[(_dafny.ZERO).toNumber()] = p0;
      _nw387[(_dafny.ONE).toNumber()] = new BigNumber(800);
      _nw387[(new BigNumber(2)).toNumber()] = p0;
      _nw387[(new BigNumber(3)).toNumber()] = p0;
      _nw387[(new BigNumber(4)).toNumber()] = p0;
      _nw387[(new BigNumber(5)).toNumber()] = p0;
      _nw387[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length);
      _nw387[(new BigNumber(7)).toNumber()] = p0;
      _nw387[(new BigNumber(8)).toNumber()] = (_2117_v7).dtor_cf48;
      _nw387[(new BigNumber(9)).toNumber()] = new BigNumber((_2118_v8).length);
      _nw387[(new BigNumber(10)).toNumber()] = p0;
      _nw387[(new BigNumber(11)).toNumber()] = p0;
      _nw387[(new BigNumber(12)).toNumber()] = p0;
      _nw387[(new BigNumber(13)).toNumber()] = p0;
      _nw387[(new BigNumber(14)).toNumber()] = p0;
      _nw387[(new BigNumber(15)).toNumber()] = p0;
      _nw387[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), function (_2122_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })).length);
      _nw387[(new BigNumber(17)).toNumber()] = p0;
      _nw387[(new BigNumber(18)).toNumber()] = new BigNumber((_2119_v9).cardinality());
      _nw387[(new BigNumber(19)).toNumber()] = p0;
      _nw387[(new BigNumber(20)).toNumber()] = new BigNumber((_2120_v10).length);
      _2121_v11 = _nw387;
      let _2123_v12;
      _2123_v12 = _module.D3.create_DC6(_2121_v11);
      let _pat_let_tv46 = _2121_v11;
      let _pat_let_tv47 = _2121_v11;
      let _2124_v13;
      let _nw388 = Array((new BigNumber(24)).toNumber());
      _nw388[(_dafny.ZERO).toNumber()] = _2123_v12;
      _nw388[(_dafny.ONE).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(2)).toNumber()] = _module.D3.create_DC6(_2121_v11);
      _nw388[(new BigNumber(3)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(4)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(5)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(6)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(7)).toNumber()] = function (_pat_let73_0) {
        return function (_2125_dt__update__tmp_h0) {
          return function (_pat_let74_0) {
            return function (_2126_dt__update_hcf6_h0) {
              return _module.D3.create_DC6(_2126_dt__update_hcf6_h0);
            }(_pat_let74_0);
          }(_pat_let_tv46);
        }(_pat_let73_0);
      }(_2123_v12);
      _nw388[(new BigNumber(8)).toNumber()] = function (_pat_let75_0) {
        return function (_2127_dt__update__tmp_h1) {
          return function (_pat_let76_0) {
            return function (_2128_dt__update_hcf6_h1) {
              return _module.D3.create_DC6(_2128_dt__update_hcf6_h1);
            }(_pat_let76_0);
          }(_pat_let_tv47);
        }(_pat_let75_0);
      }(_module.D3.create_DC6(_2121_v11));
      _nw388[(new BigNumber(9)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(10)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(11)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(12)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(13)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(14)).toNumber()] = _module.D3.create_DC6(_2121_v11);
      _nw388[(new BigNumber(15)).toNumber()] = _module.D3.create_DC6(_2121_v11);
      _nw388[(new BigNumber(16)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(17)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(18)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(19)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(20)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(21)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(22)).toNumber()] = _2123_v12;
      _nw388[(new BigNumber(23)).toNumber()] = _2123_v12;
      _2124_v13 = _nw388;
      let _2129_v14;
      _2129_v14 = _dafny.Set.fromElements(_2124_v13);
      let _2130_v15;
      _2130_v15 = _module.D9.create_DC23(_2129_v14, _2110_v0, _2115_v5, _2121_v11, _2112_v2);
      let _source20 = _2130_v15;
      if (_source20.is_DC23) {
        let _2131___mcc_h0 = (_source20).cf37;
        let _2132___mcc_h1 = (_source20).cf38;
        let _2133___mcc_h2 = (_source20).cf39;
        let _2134___mcc_h3 = (_source20).cf40;
        let _2135___mcc_h4 = (_source20).cf41;
        let _2136_cf41 = _2135___mcc_h4;
        let _2137_cf40 = _2134___mcc_h3;
        let _2138_cf39 = _2133___mcc_h2;
        let _2139_cf38 = _2132___mcc_h1;
        let _2140_cf37 = _2131___mcc_h0;
        let _2141_v16;
        _2141_v16 = _dafny.MultiSet.fromElements(_2115_v5);
        let _2142_v17;
        _2142_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.__default.fm19(globalState)),_2141_v16);
        let _2143_v18;
        _2143_v18 = _dafny.Seq.of(false, _2138_cf39);
        let _2144_v19;
        _2144_v19 = _dafny.Seq.of(_2143_v18);
        let _2145_v20;
        let _nw389 = Array((_dafny.ONE).toNumber());
        _nw389[(_dafny.ZERO).toNumber()] = (((_2142_v17).contains(_2144_v19)) ? ((_2142_v17).get(_2144_v19)) : (_2141_v16));
        _2145_v20 = _nw389;
        let _index359 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_2145_v20).length));
        (_2145_v20)[_index359] = (_2141_v16).Intersect(_2141_v16);
        let _index360 = _module.__default.safeIndex(new BigNumber(465), new BigNumber((_2145_v20).length));
        (_2145_v20)[_index360] = _dafny.MultiSet.FromArray(_2143_v18);
        _2138_cf39 = _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("esckela"), new _dafny.CodePoint('v'.codePointAt(0)));
        let _index361 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2114_v4).length));
        (_2114_v4)[_index361] = _2115_v5;
        let _index362 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2114_v4).length));
        (_2114_v4)[_index362] = _2115_v5;
        let _2146_v21;
        let _init59 = ((_2147_v5) => function (_2148_i1) {
          return !(_2147_v5);
        })(_2115_v5);
        let _nw390 = Array((new BigNumber(20)).toNumber());
        for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw390.length); _i0_59++) {
          _nw390[_i0_59] = _init59(new BigNumber(_i0_59));
        }
        _2146_v21 = _nw390;
        (globalState).f14 = _2146_v21;
      } else if (_source20.is_DC24) {
        let _2149___mcc_h5 = (_source20).cf42;
        let _2150___mcc_h6 = (_source20).cf43;
        let _2151_cf43 = _2150___mcc_h6;
        let _2152_cf42 = _2149___mcc_h5;
        let _2153_v22;
        _2153_v22 = _dafny.MultiSet.fromElements(_2115_v5, _2115_v5, _2115_v5, _2115_v5, _2115_v5);
        let _index363 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_2121_v11).length));
        (_2121_v11)[_index363] = ((((_2153_v22).contains(_2115_v5)) ? ((_2153_v22).get(_2115_v5)) : (p0))).multipliedBy(p0);
        let _index364 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_2121_v11).length));
        (_2121_v11)[_index364] = _module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("bb")).length), (_dafny.ZERO).minus(p0))));
        (globalState).f6 = (_dafny.ZERO).minus((_2121_v11)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_2121_v11).length))]);
        _2115_v5 = _2115_v5;
        let _2154_v23;
        _2154_v23 = _dafny.Map.Empty.slice().updateUnsafe((_2121_v11)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_2121_v11).length))],false);
        let _2155_v24;
        _2155_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2154_v23,(_2121_v11)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_2121_v11).length))]);
        let _2156_v25;
        let _nw391 = new _module.C2();
        _nw391.__ctor((_2121_v11)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_2121_v11).length))], (((_2118_v8).contains((_2121_v11)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_2121_v11).length))])) ? ((_2118_v8).get((_2121_v11)[_module.__default.safeIndex(new BigNumber(279), new BigNumber((_2121_v11).length))])) : (p0)), (_2155_v24).Merge(_2155_v24));
        _2156_v25 = _nw391;
      } else if (_source20.is_DC22) {
        let _2157___mcc_h7 = (_source20).cf36;
        let _2158_cf36 = _2157___mcc_h7;
        let _2159_v26;
        let _init60 = ((_2160_v5, _2161_p0) => function (_2162_i2) {
          return ((_2160_v5) ? (_dafny.Set.fromElements(_2161_p0)) : (_dafny.Set.fromElements(_2161_p0, _2161_p0, _2161_p0, (_dafny.ZERO).minus(_2161_p0), new BigNumber((_dafny.Seq.of(true)).length))));
        })(_2115_v5, p0);
        let _nw392 = Array((new BigNumber(25)).toNumber());
        for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw392.length); _i0_60++) {
          _nw392[_i0_60] = _init60(new BigNumber(_i0_60));
        }
        _2159_v26 = _nw392;
        let _2163_v28;
        _2163_v28 = _dafny.Set.fromElements(p0, new BigNumber((function () {
          let _coll76 = new _dafny.Set();
          for (const _compr_76 of _dafny.IntegerRange(new BigNumber(271), new BigNumber(842))) {
            let _2164_v27 = _compr_76;
            if (((new BigNumber(271)).isLessThanOrEqualTo(_2164_v27)) && ((_2164_v27).isLessThan(new BigNumber(842)))) {
              _coll76.add((_2164_v27).minus(p0));
            }
          }
          return _coll76;
        }()).length), new BigNumber((_dafny.Set.fromElements(_2115_v5, _2115_v5)).length), _module.__default.fm1(_2115_v5, globalState));
        let _index365 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2159_v26).length));
        (_2159_v26)[_index365] = _2163_v28;
        let _index366 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2159_v26).length));
        (_2159_v26)[_index366] = _dafny.Set.fromElements(p0);
        let _2165_v29;
        _2165_v29 = _dafny.Set.fromElements(_2110_v0, _2110_v0, _2110_v0, _dafny.Seq.UnicodeFromString("vkytdei"), _dafny.Seq.update(_2110_v0, _module.__default.safeIndex(new BigNumber(393), new BigNumber((_2110_v0).length)), _2112_v2));
        let _index367 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_2121_v11).length));
        (_2121_v11)[_index367] = (p0).minus(new BigNumber((_2165_v29).length));
        let _index368 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_2121_v11).length));
        (_2121_v11)[_index368] = (p0).multipliedBy(p0);
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_module.__default.fm21(globalState), _2110_v0), _2110_v0)) {
          (globalState).f1 = _2115_v5;
          (globalState).f6 = (_2121_v11)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_2121_v11).length))];
          let _2166_v30;
          _2166_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2115_v5,_2115_v5);
          let _2167_v31;
          _2167_v31 = _module.D14.create_DC39(_2111_v1, _2166_v30, _2115_v5);
          let _2168_v32;
          _2168_v32 = _module.D15.create_DC43(_2115_v5, (_2167_v31).dtor_cf70, _2115_v5);
          let _2169_v33;
          _2169_v33 = _dafny.MultiSet.fromElements(_2115_v5);
          let _2170_v34;
          _2170_v34 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_2121_v11)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_2121_v11).length))])).length),_2115_v5),p0);
          let _2171_v35;
          _2171_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2115_v5);
          let _2172_v36;
          let _nw393 = new _module.C3();
          _nw393.__ctor((_2115_v5) || ((_2168_v32).dtor_cf74), ((_2115_v5) ? (new BigNumber((_2169_v33).cardinality())) : ((_module.D5.create_DC12(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(855)), ((_2173_v2) => function (_2174_i3) {
  return _2173_v2;
})(_2112_v2))).length), _2115_v5, _2112_v2)).dtor_cf15)), ((_2170_v34).update(_2171_v35, new BigNumber(-637))).Merge((_2170_v34)));
          _2172_v36 = _nw393;
          let _2175_v37;
          _2175_v37 = _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)));
          let _rhs286 = ((p0).multipliedBy((_2121_v11)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_2121_v11).length))])).minus(_module.__default.safeModuloInt((_2172_v36).f19, (_2111_v1)[_module.__default.safeIndex((((_2169_v33).contains(_2115_v5)) ? ((_2169_v33).get(_2115_v5)) : (new BigNumber((_2116_v6).length))), new BigNumber((_2111_v1).length))]));
          let _rhs287 = (new BigNumber(((_2175_v37).Difference(_2175_v37)).length)).isLessThanOrEqualTo(p0);
          let _rhs288 = _2172_v36;
          let _lhs256 = globalState;
          _lhs256.f6 = _rhs286;
          _2115_v5 = _rhs287;
          _2172_v36 = _rhs288;
          let _2176_v38;
          let _nw394 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.of());
          _2176_v38 = _nw394;
          let _2177_v39;
          _2177_v39 = _dafny.Seq.of((_2172_v36).f18);
          let _2178_v40;
          _2178_v40 = _module.D11.create_DC31(_2177_v39, new BigNumber((_2119_v9).cardinality()), (_2172_v36).f19, p0, (_2172_v36).f19);
          let _index369 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_2176_v38).length));
          (_2176_v38)[_index369] = (_2178_v40).dtor_cf54;
          let _index370 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_2176_v38).length));
          let _rhs289 = _2112_v2;
          let _rhs290 = _dafny.Seq.Concat(_2177_v39, _2177_v39);
          let _rhs291 = (_module.__default.safeModuloInt(new BigNumber(258), (_2172_v36).f19)).minus(p0);
          let _lhs257 = _2176_v38;
          let _lhs258 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_2176_v38).length));
          _2112_v2 = _rhs289;
          _lhs257[_lhs258] = _rhs290;
          r0 = _rhs291;
          (globalState).f1 = _2115_v5;
        } else {
          let _index371 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_2121_v11).length));
          (_2121_v11)[_index371] = (_2121_v11)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_2121_v11).length))];
          (globalState).f6 = ((_2115_v5) ? ((_2121_v11)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_2121_v11).length))]) : (p0));
          let _2179_v41;
          let _nw395 = Array((new BigNumber(11)).toNumber()).fill(_module.D16.Default());
          _2179_v41 = _nw395;
          let _2180_v42;
          _2180_v42 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("fyhai"),_2179_v41);
          _2179_v41 = (((_2180_v42).contains(_2110_v0)) ? ((_2180_v42).get(_2110_v0)) : (_2179_v41));
          _2110_v0 = _2110_v0;
          let _2181_v43;
          let _nw396 = new _module.C0();
          _nw396.__ctor(_dafny.Seq.update(_dafny.Seq.of(_2123_v12), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(_2123_v12)).length)), _2123_v12), _2115_v5);
          _2181_v43 = _nw396;
        }
        let _rhs292 = _module.__default.safeModuloInt(new BigNumber(418), p0);
        let _rhs293 = _dafny.Seq.Concat(_2110_v0, _2110_v0);
        let _rhs294 = (new BigNumber(((_2159_v26)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_2159_v26).length))]).length)).isEqualTo(p0);
        let _rhs295 = _2114_v4;
        let _lhs259 = globalState;
        let _lhs260 = globalState;
        let _lhs261 = globalState;
        _lhs259.f6 = _rhs292;
        _2110_v0 = _rhs293;
        _lhs260.f1 = _rhs294;
        _lhs261.f14 = _rhs295;
      } else {
        let _2182___mcc_h8 = (_source20).cf44;
        let _2183_cf44 = _2182___mcc_h8;
        if ((_module.D20.create_DC50(_2115_v5, p0, p0, _2115_v5, _module.__default.fm2(globalState))).dtor_cf89) {
          let _2184_v44;
          let _nw397 = new _module.C1();
          _nw397.__ctor((p0).multipliedBy(p0));
          _2184_v44 = _nw397;
          let _2185_v45;
          _2185_v45 = _dafny.Map.Empty.slice().updateUnsafe((_2184_v44).f24,_module.__default.fm2(globalState));
          _2185_v45 = (_2185_v45).update(p0, _2115_v5);
          let _2186_v46;
          let _nw398 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2186_v46 = _nw398;
          let _index372 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_2186_v46).length));
          (_2186_v46)[_index372] = _2112_v2;
          let _index373 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_2186_v46).length));
          (_2186_v46)[_index373] = _2112_v2;
          let _index374 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_2114_v4).length));
          (_2114_v4)[_index374] = _2115_v5;
          let _2187_v47;
          _2187_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2119_v9).cardinality()),_dafny.Set.fromElements(new BigNumber(619), new BigNumber(-1)));
          let _2188_v48;
          _2188_v48 = _dafny.Set.fromElements(p0);
          let _2189_v49;
          _2189_v49 = _dafny.Map.Empty.slice().updateUnsafe((_2184_v44).f24,_2188_v48);
          let _2190_v50;
          _2190_v50 = _dafny.Set.fromElements(_2115_v5);
          let _2191_v51;
          _2191_v51 = _module.D0.create_DC1(new BigNumber((_dafny.Seq.of(_2115_v5)).length));
          let _2192_v52;
          _2192_v52 = _module.D0.create_DC2(_2191_v51);
          let _2193_v53;
          let _nw399 = Array((new BigNumber(29)).toNumber());
          _nw399[(_dafny.ZERO).toNumber()] = _2187_v47;
          _nw399[(_dafny.ONE).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(2)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(3)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(4)).toNumber()] = _2187_v47;
          _nw399[(new BigNumber(5)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(6)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(7)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2185_v45).length),_dafny.Set.fromElements(p0, new BigNumber((_2190_v50).length)));
          _nw399[(new BigNumber(9)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(10)).toNumber()] = _2187_v47;
          _nw399[(new BigNumber(11)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(12)).toNumber()] = _module.__default.fm26((_2186_v46)[_module.__default.safeIndex(new BigNumber(42), new BigNumber((_2186_v46).length))], _2192_v52, _2115_v5, _2115_v5, globalState);
          _nw399[(new BigNumber(13)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(14)).toNumber()] = _2187_v47;
          _nw399[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).fm9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(601)), ((_2194_v2) => function (_2195_i4) {
            return _2194_v2;
          })(_2112_v2)), globalState),_dafny.Set.fromElements((_2184_v44).f24, p0));
          _nw399[(new BigNumber(16)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(17)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(18)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(19)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(20)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(21)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(22)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(23)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(24)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(25)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(26)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2111_v1).length),_2188_v48);
          _nw399[(new BigNumber(27)).toNumber()] = _2189_v49;
          _nw399[(new BigNumber(28)).toNumber()] = _2189_v49;
          _2193_v53 = _nw399;
          let _2196_v54;
          let _nw400 = Array((new BigNumber(5)).toNumber());
          _nw400[(_dafny.ZERO).toNumber()] = _2193_v53;
          _nw400[(_dafny.ONE).toNumber()] = _2193_v53;
          _nw400[(new BigNumber(2)).toNumber()] = _2193_v53;
          _nw400[(new BigNumber(3)).toNumber()] = _2193_v53;
          _nw400[(new BigNumber(4)).toNumber()] = _2193_v53;
          _2196_v54 = _nw400;
          let _index375 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_2114_v4).length));
          let _rhs296 = !((_2115_v5) || (_2115_v5));
          let _rhs297 = _2196_v54;
          let _lhs262 = _2114_v4;
          let _lhs263 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_2114_v4).length));
          _lhs262[_lhs263] = _rhs296;
          _2196_v54 = _rhs297;
          let _2197_v55;
          let _nw401 = new _module.C5();
          _nw401.__ctor(_module.__default.fm1(_2115_v5, globalState));
          _2197_v55 = _nw401;
        } else {
          (globalState).f6 = p0;
          let _2198_v56;
          _2198_v56 = _module.D0.create_DC0(true);
          let _2199_v57;
          _2199_v57 = _dafny.Map.Empty.slice().updateUnsafe(_2115_v5,_2115_v5);
          let _2200_v58;
          _2200_v58 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2112_v2);
          let _2201_v59;
          _2201_v59 = _module.D3.create_DC8(_2111_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(193)), ((_2202_v2) => function (_2203_i7) {
  return _2202_v2;
})(_2112_v2)), _2115_v5, new _dafny.CodePoint('n'.codePointAt(0)));
          let _2204_v60;
          let _nw402 = Array((new BigNumber(21)).toNumber());
          _nw402[(_dafny.ZERO).toNumber()] = _2110_v0;
          _nw402[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(204)), ((_2205_v2) => function (_2206_i5) {
            return _2205_v2;
          })(_2112_v2)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(374)), ((_2207_v2) => function (_2208_i6) {
            return _2207_v2;
          })(_2112_v2)));
          _nw402[(new BigNumber(2)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("mcf");
          _nw402[(new BigNumber(4)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(5)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(6)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(7)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(8)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(9)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(10)).toNumber()] = (((_this).fm10(_2198_v56, _2199_v57, (((_2200_v58).contains(p0)) ? ((_2200_v58).get(p0)) : (new _dafny.CodePoint('p'.codePointAt(0)))), p0, globalState)) ? (_2110_v0) : (_2110_v0));
          _nw402[(new BigNumber(11)).toNumber()] = (_2201_v59).dtor_cf9;
          _nw402[(new BigNumber(12)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fmsrvacn"), _2110_v0);
          _nw402[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("timqowl");
          _nw402[(new BigNumber(15)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("a"), _2110_v0);
          _nw402[(new BigNumber(17)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(18)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(19)).toNumber()] = _2110_v0;
          _nw402[(new BigNumber(20)).toNumber()] = _2110_v0;
          _2204_v60 = _nw402;
          let _2209_v61;
          _2209_v61 = _module.D8.create_DC21(_2110_v0, _dafny.Set.fromElements(_2115_v5, _2115_v5));
          let _index376 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2204_v60).length));
          (_2204_v60)[_index376] = (_2209_v61).dtor_cf34;
          let _2210_v62;
          _2210_v62 = _dafny.Seq.of(_dafny.Seq.update(_2110_v0, _module.__default.safeIndex(p0, new BigNumber((_2110_v0).length)), _2112_v2), _2110_v0, _dafny.Seq.Concat(_2110_v0, _2110_v0), _2110_v0);
          let _index377 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_2204_v60).length));
          (_2204_v60)[_index377] = (_2210_v62)[_module.__default.safeIndex(p0, new BigNumber((_2210_v62).length))];
          _2115_v5 = (p0).isEqualTo((p0).multipliedBy(p0));
          let _2211_v63;
          _2211_v63 = _dafny.Set.fromElements(_2115_v5);
          (globalState).f6 = new BigNumber(((_2211_v63).Union((_2211_v63).Intersect(_module.__default.fm31((_dafny.ZERO).minus(p0), _2115_v5, globalState)))).length);
          let _index378 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_2114_v4).length));
          (_2114_v4)[_index378] = _2115_v5;
          let _index379 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_2114_v4).length));
          (_2114_v4)[_index379] = _2115_v5;
        }
        let _2212_v64;
        let _nw403 = new _module.C4();
        _nw403.__ctor();
        _2212_v64 = _nw403;
        let _2213_v65;
        let _nw404 = new _module.C5();
        _nw404.__ctor(p0);
        _2213_v65 = _nw404;
        let _2214_v66;
        _2214_v66 = _module.D23.create_DC53(_2213_v65);
        let _2215_v67;
        _2215_v67 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(335)), ((_2216_v2) => function (_2217_i8) {
          return _2216_v2;
        })(_2112_v2))).length)).multipliedBy(_module.__default.fm1(_2115_v5, globalState)),(_2214_v66).dtor_cf93);
        _2215_v67 = (_2215_v67).update(p0, _2213_v65);
        let _2218_v68;
        _2218_v68 = _dafny.MultiSet.fromElements(_2115_v5, _2115_v5, _2115_v5, _2115_v5, !(_2115_v5));
        let _rhs298 = (_2115_v5) && (((_dafny.ZERO).minus(p0)).isLessThan(p0));
        let _rhs299 = _2114_v4;
        let _rhs300 = (_dafny.ZERO).minus((_2213_v65).f16);
        let _rhs301 = _2110_v0;
        let _rhs302 = !(!(!(((_2218_v68).Difference((_2213_v65).fm11(new BigNumber(18), globalState))).IsProperSubsetOf(_2218_v68))));
        let _lhs264 = globalState;
        let _lhs265 = globalState;
        let _lhs266 = globalState;
        _lhs264.f1 = _rhs298;
        _lhs265.f14 = _rhs299;
        _lhs266.f6 = _rhs300;
        _2110_v0 = _rhs301;
        _2115_v5 = _rhs302;
      }
      _2121_v11 = _2121_v11;
      let _2219_v69;
      _2219_v69 = _module.D14.create_DC40(_module.__default.fm2(globalState));
      let _pat_let_tv48 = _2115_v5;
      _2219_v69 = function (_pat_let77_0) {
        return function (_2220_dt__update__tmp_h5) {
          return function (_pat_let78_0) {
            return function (_2221_dt__update_hcf71_h0) {
              return _module.D14.create_DC40(_2221_dt__update_hcf71_h0);
            }(_pat_let78_0);
          }(_pat_let_tv48);
        }(_pat_let77_0);
      }(_2219_v69);
      _2115_v5 = _2115_v5;
      let _2222_v71;
      let _nw405 = new _module.C3();
      _nw405.__ctor((_2119_v9).contains(p0), _module.__default.safeModuloInt(p0, p0), function () {
        let _coll77 = new _dafny.Map();
        for (const _compr_77 of (_module.__default.fm45(p0, globalState)).Elements) {
          let _2223_v70 = _compr_77;
          if (_dafny.Seq.contains(_module.__default.fm45(p0, globalState), _2223_v70)) {
            _coll77.push([_2223_v70,p0]);
          }
        }
        return _coll77;
      }());
      _2222_v71 = _nw405;
      let _hi10 = (_2222_v71).f19;
      for (let _2224_i9 = new BigNumber((_dafny.Seq.Concat(_2110_v0, _2110_v0)).length); _2224_i9.isLessThan(_hi10); _2224_i9 = _2224_i9.plus(_dafny.ONE)) {
        (globalState).f14 = _2114_v4;
        _2115_v5 = (_2222_v71).f18;
        _2115_v5 = _2115_v5;
        (globalState).f1 = ((((_2222_v71).f19).isLessThan((_2222_v71).f19)) ? ((_module.__default.fm1(_2115_v5, globalState)).isLessThan((_2222_v71).f19)) : (((_2222_v71).f18) && (!((_2222_v71).f18))));
      }
      r0 = ((p0).minus((_dafny.ZERO).minus(p0))).minus((_2222_v71).f19);
      return r0;
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (!(_dafny.Set.fromElements(false, !(false), !(true), true)).contains(!(false))) {
        return _module.D0.create_DC2(_module.D0.create_DC0(true));
      } else {
        return _module.D0.create_DC2(_module.D0.create_DC0((_module.D0.create_DC0(true)).dtor_cf0));
      }
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      if (false) {
        return !(true);
      } else {
        return false;
      }
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _2225_v0;
      _2225_v0 = true;
      if (!(((_2225_v0) ? ((true) === (true)) : (_2225_v0)))) {
        let _2226_v1;
        _2226_v1 = _dafny.MultiSet.fromElements(!(_2225_v0), _2225_v0, _2225_v0);
        (globalState).f6 = _module.__default.fm1((_2226_v1).IsProperSubsetOf(_2226_v1), globalState);
        (_this).m7(globalState);
        if (_2225_v0) {
          let _2227_v2;
          _2227_v2 = new BigNumber(834);
          (globalState).f6 = _module.__default.safeModuloInt(_2227_v2, _2227_v2);
          let _2228_v3;
          _2228_v3 = _dafny.Seq.UnicodeFromString("bjdokiwra");
          let _2229_v4;
          let _init61 = ((_2230_v0) => function (_2231_i0) {
            return _2230_v0;
          })(_2225_v0);
          let _nw406 = Array((new BigNumber(20)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw406.length); _i0_61++) {
            _nw406[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _2229_v4 = _nw406;
          let _2232_v5;
          _2232_v5 = _dafny.Seq.of(_2225_v0, _2225_v0);
          let _index380 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_2229_v4).length));
          (_2229_v4)[_index380] = (_2232_v5)[_module.__default.safeIndex(_2227_v2, new BigNumber((_2232_v5).length))];
          let _2233_v6;
          _2233_v6 = _dafny.Seq.of(_2227_v2, _2227_v2);
          let _index381 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_2229_v4).length));
          let _rhs303 = _2228_v3;
          let _rhs304 = _dafny.Seq.IsPrefixOf(_2233_v6, _2233_v6);
          let _lhs267 = _2229_v4;
          let _lhs268 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_2229_v4).length));
          _2228_v3 = _rhs303;
          _lhs267[_lhs268] = _rhs304;
          let _2234_v7;
          let _nw407 = new _module.C4();
          _nw407.__ctor();
          _2234_v7 = _nw407;
          let _2235_v8;
          let _nw408 = new _module.C6();
          _nw408.__ctor();
          _2235_v8 = _nw408;
          let _2236_v9;
          _2236_v9 = new _dafny.CodePoint('i'.codePointAt(0));
          let _2237_v10;
          _2237_v10 = _module.D5.create_DC11(_2227_v2);
          let _2238_v11;
          _2238_v11 = _dafny.Seq.of(_2237_v10);
          let _2239_v12;
          _2239_v12 = _dafny.MultiSet.fromElements(_2227_v2);
          let _2240_v13;
          _2240_v13 = _dafny.Set.fromElements(_2225_v0, _2225_v0);
          let _2241_v14;
          _2241_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2228_v3).length),new BigNumber((_2233_v6).length));
          let _2242_v15;
          let _nw409 = Array((new BigNumber(10)).toNumber());
          _nw409[(_dafny.ZERO).toNumber()] = (new BigNumber((_2238_v11).length)).minus(_2227_v2);
          _nw409[(_dafny.ONE).toNumber()] = _2227_v2;
          _nw409[(new BigNumber(2)).toNumber()] = _2227_v2;
          _nw409[(new BigNumber(3)).toNumber()] = _2227_v2;
          _nw409[(new BigNumber(4)).toNumber()] = _2227_v2;
          _nw409[(new BigNumber(5)).toNumber()] = _2227_v2;
          _nw409[(new BigNumber(6)).toNumber()] = _2227_v2;
          _nw409[(new BigNumber(7)).toNumber()] = ((((_2239_v12).contains(_2227_v2)) ? ((_2239_v12).get(_2227_v2)) : (new BigNumber((_2233_v6).length)))).multipliedBy(_2227_v2);
          _nw409[(new BigNumber(8)).toNumber()] = new BigNumber((_2240_v13).length);
          _nw409[(new BigNumber(9)).toNumber()] = (new BigNumber((_2241_v14).length)).multipliedBy(_2227_v2);
          _2242_v15 = _nw409;
          let _2243_v16;
          _2243_v16 = _dafny.Seq.of(_2228_v3);
          let _2244_v17;
          _2244_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2243_v16).length),false);
          let _index382 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2242_v15).length));
          (_2242_v15)[_index382] = new BigNumber(((_2244_v17).Merge(_2244_v17)).length);
          let _2245_v18;
          _2245_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2225_v0,_2225_v0);
          let _index383 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2242_v15).length));
          let _rhs305 = _module.__default.fm29((_2229_v4)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_2229_v4).length))], _2227_v2, _2227_v2, (((_2245_v18).contains(_2225_v0)) ? ((_2245_v18).get(_2225_v0)) : (_2225_v0)), globalState);
          let _rhs306 = new BigNumber(16);
          let _lhs269 = _2242_v15;
          let _lhs270 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2242_v15).length));
          _2236_v9 = _rhs305;
          _lhs269[_lhs270] = _rhs306;
        } else {
          let _2246_v19;
          _2246_v19 = new BigNumber(990);
          let _2247_v20;
          _2247_v20 = _dafny.Seq.of(_2246_v19);
          let _2248_v21;
          _2248_v21 = new _dafny.CodePoint('d'.codePointAt(0));
          let _2249_v22;
          _2249_v22 = _module.D3.create_DC8(_2247_v20, _dafny.Seq.UnicodeFromString("cgrydx"), _2225_v0, _2248_v21);
          let _2250_v23;
          _2250_v23 = _dafny.Seq.UnicodeFromString("okujifd");
          let _pat_let_tv49 = _2250_v23;
          let _pat_let_tv50 = _2250_v23;
          let _pat_let_tv51 = _2246_v19;
          let _pat_let_tv52 = _2250_v23;
          _2225_v0 = !(!((function (_pat_let79_0) {
            return function (_2251_dt__update__tmp_h0) {
              return function (_pat_let80_0) {
                return function (_2252_dt__update_hcf9_h0) {
                  return function (_pat_let81_0) {
                    return function (_2253_dt__update_hcf11_h0) {
                      return _module.D3.create_DC8((_2251_dt__update__tmp_h0).dtor_cf8, _2252_dt__update_hcf9_h0, (_2251_dt__update__tmp_h0).dtor_cf10, _2253_dt__update_hcf11_h0);
                    }(_pat_let81_0);
                  }((_pat_let_tv50)[_module.__default.safeIndex(_pat_let_tv51, new BigNumber((_pat_let_tv52).length))]);
                }(_pat_let80_0);
              }(_pat_let_tv49);
            }(_pat_let79_0);
          }(_2249_v22)).dtor_cf10));
          let _2254_v24;
          _2254_v24 = _module.D13.create_DC37(!(_2225_v0) || (_2225_v0), _2225_v0);
          _2254_v24 = _2254_v24;
          let _2255_v25;
          let _init62 = ((_2256_v23) => function (_2257_i1) {
            return _2256_v23;
          })(_2250_v23);
          let _nw410 = Array((new BigNumber(5)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw410.length); _i0_62++) {
            _nw410[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _2255_v25 = _nw410;
          let _index384 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_2255_v25).length));
          (_2255_v25)[_index384] = _dafny.Seq.Concat(_2250_v23, _2250_v23);
          let _2258_v26;
          let _nw411 = Array((new BigNumber(26)).toNumber()).fill([]);
          _2258_v26 = _nw411;
          let _2259_v27;
          let _nw412 = Array((new BigNumber(29)).toNumber()).fill(false);
          _2259_v27 = _nw412;
          let _index385 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_2258_v26).length));
          (_2258_v26)[_index385] = _2259_v27;
          let _2260_v28;
          _2260_v28 = _dafny.Seq.of(false);
          let _index386 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_2255_v25).length));
          let _index387 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_2258_v26).length));
          let _rhs307 = _2246_v19;
          let _rhs308 = _dafny.Seq.Concat(_2250_v23, _dafny.Seq.Concat(_2250_v23, _2250_v23));
          let _rhs309 = ((_2225_v0) ? ((_dafny.ZERO).minus((_2246_v19).minus(_2246_v19))) : (new BigNumber((_dafny.Seq.Concat(_2260_v28, _2260_v28)).length)));
          let _rhs310 = _2259_v27;
          let _lhs271 = _2255_v25;
          let _lhs272 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_2255_v25).length));
          let _lhs273 = _2258_v26;
          let _lhs274 = _module.__default.safeIndex(new BigNumber(939), new BigNumber((_2258_v26).length));
          r0 = _rhs307;
          _lhs271[_lhs272] = _rhs308;
          _2246_v19 = _rhs309;
          _lhs273[_lhs274] = _rhs310;
          (globalState).f6 = (new BigNumber(110)).multipliedBy((_2246_v19).plus(_2246_v19));
          let _2261_v29;
          _2261_v29 = _module.D0.create_DC0(_2225_v0);
          (globalState).f1 = (_2261_v29).dtor_cf0;
        }
        let _2262_v30;
        let _init63 = ((_2263_v0) => function (_2264_i2) {
          return _2263_v0;
        })(_2225_v0);
        let _nw413 = Array((new BigNumber(4)).toNumber());
        for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw413.length); _i0_63++) {
          _nw413[_i0_63] = _init63(new BigNumber(_i0_63));
        }
        _2262_v30 = _nw413;
        let _2265_v31;
        _2265_v31 = new BigNumber(605);
        let _2266_v32;
        _2266_v32 = _dafny.Seq.UnicodeFromString("krqgpb");
        let _2267_v34;
        _2267_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2265_v31,_2225_v0);
        let _2268_v35;
        _2268_v35 = _dafny.MultiSet.fromElements(_2267_v34);
        let _2269_v36;
        let _nw414 = new _module.C2();
        _nw414.__ctor(_2265_v31, new BigNumber((_2266_v32).length), function () {
          let _coll78 = new _dafny.Map();
          for (const _compr_78 of (_2268_v35).Elements) {
            let _2270_v33 = _compr_78;
            if ((_2268_v35).contains(_2270_v33)) {
              _coll78.push([_2270_v33,_2265_v31]);
            }
          }
          return _coll78;
        }());
        _2269_v36 = _nw414;
        let _2271_v37;
        _2271_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2269_v36,_2225_v0);
        let _2272_v38;
        _2272_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2225_v0,_2262_v30);
        let _2273_v39;
        _2273_v39 = _module.D0.create_DC0(_2225_v0);
        let _2274_v40;
        _2274_v40 = _2262_v30;
        let _2275_v41;
        let _nw415 = Array((new BigNumber(12)).toNumber());
        _nw415[(_dafny.ZERO).toNumber()] = _2262_v30;
        _nw415[(_dafny.ONE).toNumber()] = (((((_2271_v37).contains(_2269_v36)) ? ((_2271_v37).get(_2269_v36)) : (_2225_v0))) ? (_2262_v30) : (_2262_v30));
        _nw415[(new BigNumber(2)).toNumber()] = (((_2272_v38).contains((_2269_v36).fm16(_2273_v39, _2225_v0, globalState))) ? ((_2272_v38).get((_2269_v36).fm16(_2273_v39, _2225_v0, globalState))) : (_2262_v30));
        _nw415[(new BigNumber(3)).toNumber()] = _2262_v30;
        _nw415[(new BigNumber(4)).toNumber()] = (_2274_v40);
        _nw415[(new BigNumber(5)).toNumber()] = _2262_v30;
        _nw415[(new BigNumber(6)).toNumber()] = _2262_v30;
        _nw415[(new BigNumber(7)).toNumber()] = _2262_v30;
        _nw415[(new BigNumber(8)).toNumber()] = _2262_v30;
        _nw415[(new BigNumber(9)).toNumber()] = _2262_v30;
        _nw415[(new BigNumber(10)).toNumber()] = _2262_v30;
        _nw415[(new BigNumber(11)).toNumber()] = _2262_v30;
        _2275_v41 = _nw415;
        let _2276_v42;
        let _nw416 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2276_v42 = _nw416;
        let _2277_v43;
        _2277_v43 = _dafny.Seq.of(_2266_v32, _dafny.Seq.UnicodeFromString("ddvwnpi"), _2266_v32);
        let _2278_v44;
        _2278_v44 = new _dafny.CodePoint('b'.codePointAt(0));
        let _index388 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_2276_v42).length));
        (_2276_v42)[_index388] = (_2277_v43)[_module.__default.safeIndex(new BigNumber((_module.__default.fm13(_2225_v0, _2278_v44, _2265_v31, globalState)).length), new BigNumber((_2277_v43).length))];
        let _2279_v45;
        _2279_v45 = _dafny.Seq.of(!(_2225_v0) || (_module.__default.fm2(globalState)), _2225_v0, false, _2225_v0, _2225_v0);
        let _index389 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_2276_v42).length));
        let _rhs311 = (_2279_v45)[_module.__default.safeIndex((_2269_v36).f20, new BigNumber((_2279_v45).length))];
        let _rhs312 = _2275_v41;
        let _rhs313 = _dafny.Seq.Concat(_2266_v32, _2266_v32);
        let _rhs314 = _module.__default.fm2(globalState);
        let _lhs275 = _2276_v42;
        let _lhs276 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_2276_v42).length));
        let _lhs277 = globalState;
        r1 = _rhs311;
        _2275_v41 = _rhs312;
        _lhs275[_lhs276] = _rhs313;
        _lhs277.f1 = _rhs314;
        if (((_2269_v36.f21).isLessThanOrEqualTo(_module.__default.fm1(_2225_v0, globalState))) === (((_2225_v0) ? (!(_2225_v0)) : (_2225_v0)))) {
          _2273_v39 = _2273_v39;
          let _2280_v46;
          let _nw417 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _2280_v46 = _nw417;
          let _2281_v47;
          _2281_v47 = _module.D3.create_DC6(_2280_v46);
          let _2282_v48;
          _2282_v48 = _dafny.Seq.of(_2281_v47, _2281_v47);
          let _2283_v49;
          _2283_v49 = _module.D11.create_DC30(_2282_v48);
          let _2284_v50;
          _2284_v50 = _dafny.Map.Empty.slice().updateUnsafe((_2269_v36).f20,(_2269_v36).f20);
          let _2285_v51;
          _2285_v51 = _dafny.Map.Empty.slice().updateUnsafe(_2283_v49,_2284_v50);
          _2285_v51 = (_2285_v51).update(_2283_v49, _2284_v50);
          _2277_v43 = _2277_v43;
          (globalState).f1 = !(_2225_v0);
          (globalState).f14 = _2262_v30;
        } else {
          let _2286_v52;
          _2286_v52 = _module.D14.create_DC40(_2225_v0);
          _2286_v52 = _module.D14.create_DC40(_2225_v0);
          let _2287_v53;
          let _init64 = ((_2288_v34, _2289_v36) => function (_2290_i3) {
            return _dafny.Map.Empty.slice().updateUnsafe(_2288_v34,(_2289_v36).f20);
          })(_2267_v34, _2269_v36);
          let _nw418 = Array((new BigNumber(23)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw418.length); _i0_64++) {
            _nw418[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2287_v53 = _nw418;
          _2287_v53 = _2287_v53;
          let _2291_v54;
          _2291_v54 = _dafny.Map.Empty.slice().updateUnsafe(((_2269_v36).f20).plus(_2265_v31),new BigNumber(((_2276_v42)[_module.__default.safeIndex(new BigNumber(803), new BigNumber((_2276_v42).length))]).length));
          let _2292_v55;
          _2292_v55 = _module.D3.create_DC7(_2265_v31);
          let _rhs315 = (((_2291_v54).contains((_2292_v55).dtor_cf7)) ? ((_2291_v54).get((_2292_v55).dtor_cf7)) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(_2265_v31))));
          let _rhs316 = !(_2225_v0);
          let _rhs317 = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_dafny.ZERO).minus(_2269_v36.f21)).minus(_2269_v36.f21)).plus(new BigNumber(59))));
          let _lhs278 = globalState;
          let _lhs279 = _2269_v36;
          _lhs278.f6 = _rhs315;
          r1 = _rhs316;
          _lhs279.f21 = _rhs317;
          (globalState).f1 = true;
          let _index390 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_2262_v30).length));
          (_2262_v30)[_index390] = false;
          let _index391 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_2262_v30).length));
          let _rhs318 = false;
          let _rhs319 = _module.__default.fm1((_2273_v39).dtor_cf0, globalState);
          let _lhs280 = _2262_v30;
          let _lhs281 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_2262_v30).length));
          let _lhs282 = globalState;
          _lhs280[_lhs281] = _rhs318;
          _lhs282.f6 = _rhs319;
        }
      } else {
        let _2293_v56;
        let _nw419 = Array((new BigNumber(12)).toNumber());
        _nw419[(_dafny.ZERO).toNumber()] = _2225_v0;
        _nw419[(_dafny.ONE).toNumber()] = _2225_v0;
        _nw419[(new BigNumber(2)).toNumber()] = _2225_v0;
        _nw419[(new BigNumber(3)).toNumber()] = _2225_v0;
        _nw419[(new BigNumber(4)).toNumber()] = _2225_v0;
        _nw419[(new BigNumber(5)).toNumber()] = _2225_v0;
        _nw419[(new BigNumber(6)).toNumber()] = _2225_v0;
        _nw419[(new BigNumber(7)).toNumber()] = _module.__default.fm2(globalState);
        _nw419[(new BigNumber(8)).toNumber()] = _2225_v0;
        _nw419[(new BigNumber(9)).toNumber()] = _2225_v0;
        _nw419[(new BigNumber(10)).toNumber()] = _2225_v0;
        _nw419[(new BigNumber(11)).toNumber()] = _2225_v0;
        _2293_v56 = _nw419;
        let _2294_v57;
        _2294_v57 = _dafny.Seq.of(_2293_v56, _2293_v56);
        let _2295_v58;
        _2295_v58 = _dafny.Map.Empty.slice().updateUnsafe(_2225_v0,new BigNumber((_2294_v57).length));
        let _2296_v59;
        let _nw420 = new _module.C1();
        _nw420.__ctor(_module.__default.fm1(_2225_v0, globalState));
        _2296_v59 = _nw420;
        let _2297_v60;
        _2297_v60 = _dafny.Map.Empty.slice().updateUnsafe(_2225_v0,_2296_v59);
        _2295_v58 = (_2295_v58).update(_2225_v0, _module.__default.safeModuloInt(new BigNumber((_2297_v60).length), (_dafny.ZERO).minus((_2296_v59).f24)));
        let _2298_v61;
        _2298_v61 = _dafny.Seq.UnicodeFromString("jqob");
        _2298_v61 = ((_2225_v0) ? (_2298_v61) : (_2298_v61));
        let _2299_v62;
        let _nw421 = Array((new BigNumber(9)).toNumber()).fill(_module.D16.Default());
        _2299_v62 = _nw421;
        let _2300_v63;
        _2300_v63 = new _dafny.CodePoint('s'.codePointAt(0));
        let _index392 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2299_v62).length));
        (_2299_v62)[_index392] = _module.D16.create_DC45((_2296_v59).f24, _2300_v63, new BigNumber((_2298_v61).length), _2225_v0);
        let _2301_v64;
        _2301_v64 = _module.D16.create_DC45(new BigNumber((_2298_v61).length), _2300_v63, (_2296_v59).f24, false);
        let _2302_v65;
        _2302_v65 = _dafny.Map.Empty.slice().updateUnsafe(_2225_v0,_2300_v63);
        let _pat_let_tv53 = _2296_v59;
        let _pat_let_tv54 = _2225_v0;
        let _pat_let_tv55 = _2302_v65;
        let _pat_let_tv56 = _2225_v0;
        let _pat_let_tv57 = globalState;
        let _index393 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_2299_v62).length));
        (_2299_v62)[_index393] = function (_pat_let82_0) {
          return function (_2305_dt__update__tmp_h2) {
            return function (_pat_let85_0) {
              return function (_2306_dt__update_hcf81_h0) {
                return function (_pat_let86_0) {
                  return function (_2307_dt__update_hcf78_h0) {
                    return _module.D16.create_DC45(_2307_dt__update_hcf78_h0, (_2305_dt__update__tmp_h2).dtor_cf79, (_2305_dt__update__tmp_h2).dtor_cf80, _2306_dt__update_hcf81_h0);
                  }(_pat_let86_0);
                }((new BigNumber((_pat_let_tv55).length)).minus(_module.__default.fm1(_pat_let_tv56, _pat_let_tv57)));
              }(_pat_let85_0);
            }(_pat_let_tv54);
          }(_pat_let82_0);
        }(function (_pat_let83_0) {
          return function (_2303_dt__update__tmp_h1) {
            return function (_pat_let84_0) {
              return function (_2304_dt__update_hcf80_h0) {
                return _module.D16.create_DC45((_2303_dt__update__tmp_h1).dtor_cf78, (_2303_dt__update__tmp_h1).dtor_cf79, _2304_dt__update_hcf80_h0, (_2303_dt__update__tmp_h1).dtor_cf81);
              }(_pat_let84_0);
            }((_pat_let_tv53).f24);
          }(_pat_let83_0);
        }(_2301_v64));
        (globalState).f1 = !(_2225_v0);
        let _2308_v66;
        let _nw422 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
        _2308_v66 = _nw422;
        let _2309_v67;
        _2309_v67 = _module.D9.create_DC22(_2308_v66);
        let _2310_v68;
        _2310_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2309_v67,_2293_v56);
        _2310_v68 = _2310_v68;
      }
      r1 = !(_2225_v0);
      let _2311_v69;
      let _nw423 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _2311_v69 = _nw423;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2311_v69).length))) {
        let _2312_i4 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2312_i4)) && ((_2312_i4).isLessThan(new BigNumber((_2311_v69).length))))) {
          (_2311_v69)[(_2312_i4)] = (_2312_i4).minus((new BigNumber(((_module.D14.create_DC39(_dafny.Seq.of(new BigNumber(824)), (_module.D1.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(_2225_v0,_2225_v0))).dtor_cf3, false)).dtor_cf68).length)).minus(new BigNumber(887)));
        }
      }
      let _2313_v70;
      let _nw424 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
      _2313_v70 = _nw424;
      _2313_v70 = _2313_v70;
      r0 = new BigNumber(625);
      let _2314_v71;
      _2314_v71 = _dafny.Set.fromElements(_2225_v0, _2225_v0, !(_2225_v0));
      _2314_v71 = _2314_v71;
      let _2315_v72;
      _2315_v72 = new BigNumber(374);
      let _2316_v73;
      _2316_v73 = _module.D7.create_DC16(_2225_v0, _2315_v72, _2225_v0);
      let _2317_v74;
      _2317_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2316_v73,_dafny.Seq.UnicodeFromString("ty"));
      let _2318_v75;
      _2318_v75 = _module.D13.create_DC36(_2317_v74);
      let _pat_let_tv58 = _2315_v72;
      let _pat_let_tv59 = _2315_v72;
      let _pat_let_tv60 = _2315_v72;
      r0 = function (_source21) {
        if (_source21.is_DC37) {
          let _2319___mcc_h0 = (_source21).cf65;
          let _2320___mcc_h1 = (_source21).cf66;
          let _2321_cf66 = _2320___mcc_h1;
          let _2322_cf65 = _2319___mcc_h0;
          return (new BigNumber((_dafny.Seq.UnicodeFromString("vuc")).length)).plus(_pat_let_tv58);
        } else {
          let _2323___mcc_h2 = (_source21).cf64;
          let _2324_cf64 = _2323___mcc_h2;
          return _module.__default.safeModuloInt(_pat_let_tv59, _pat_let_tv60);
        }
      }(_2318_v75);
      let _2325_v76;
      let _nw425 = new _module.C4();
      _nw425.__ctor();
      _2325_v76 = _nw425;
      let _2326_v77;
      _2326_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2315_v72,_2325_v76);
      r1 = !((new BigNumber((_2326_v77).length)).isLessThan(new BigNumber(-718)));
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      if (p0) {
        let _2327_v0;
        let _nw426 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Set.Empty);
        _2327_v0 = _nw426;
        let _2328_v1;
        _2328_v1 = new BigNumber(693);
        let _2329_v2;
        _2329_v2 = _dafny.Set.fromElements(_2328_v1);
        let _index394 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_2327_v0).length));
        (_2327_v0)[_index394] = (_2329_v2).Union(_2329_v2);
        let _index395 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_2327_v0).length));
        (_2327_v0)[_index395] = (_2329_v2).Intersect(_2329_v2);
        let _2330_v3;
        let _init65 = ((_2331_p0) => function (_2332_i0) {
          return _2331_p0;
        })(p0);
        let _nw427 = Array((new BigNumber(8)).toNumber());
        for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw427.length); _i0_65++) {
          _nw427[_i0_65] = _init65(new BigNumber(_i0_65));
        }
        _2330_v3 = _nw427;
        let _2333_v4;
        let _nw428 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _2333_v4 = _nw428;
        let _2334_v5;
        _2334_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2330_v3,_2333_v4);
        _2334_v5 = (_2334_v5).update((_2330_v3), _2333_v4);
        _2328_v1 = _2328_v1;
        (globalState).f1 = p0;
        let _2335_v6;
        _2335_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2328_v1);
        (globalState).f1 = (_2335_v6).contains(p0);
      } else {
        if (p0) {
          let _2336_v7;
          _2336_v7 = new _dafny.CodePoint('q'.codePointAt(0));
          _2336_v7 = _2336_v7;
          let _2337_v8;
          let _nw429 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _2337_v8 = _nw429;
          let _2338_v9;
          _2338_v9 = _module.D3.create_DC6(_2337_v8);
          let _2339_v10;
          _2339_v10 = _dafny.Seq.of(_module.D3.create_DC6(_2337_v8), _2338_v9, _2338_v9);
          let _2340_v11;
          let _nw430 = new _module.C0();
          _nw430.__ctor(_dafny.Seq.Concat(_dafny.Seq.of(_2338_v9), _2339_v10), p0);
          _2340_v11 = _nw430;
          let _2341_v12;
          let _init66 = ((_2342_p0) => function (_2343_i1) {
            return _2342_p0;
          })(p0);
          let _nw431 = Array((new BigNumber(2)).toNumber());
          for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw431.length); _i0_66++) {
            _nw431[_i0_66] = _init66(new BigNumber(_i0_66));
          }
          _2341_v12 = _nw431;
          let _index396 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_2341_v12).length));
          (_2341_v12)[_index396] = (p0) === (p0);
          let _index397 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_2341_v12).length));
          (_2341_v12)[_index397] = (_2340_v11).f23;
          let _2344_v13;
          _2344_v13 = new BigNumber(946);
          let _2345_v14;
          _2345_v14 = _dafny.Seq.of(false);
          let _index398 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_2341_v12).length));
          (_2341_v12)[_index398] = _dafny.Seq.IsPrefixOf(_module.__default.fm13(p0, new _dafny.CodePoint('q'.codePointAt(0)), _2344_v13, globalState), _2345_v14);
          (globalState).f1 = (new BigNumber(95)).isLessThan((_dafny.ZERO).minus(_2344_v13));
        } else {
          let _2346_v15;
          let _nw432 = Array((new BigNumber(5)).toNumber()).fill(false);
          _2346_v15 = _nw432;
          (globalState).f14 = ((p0) ? (_2346_v15) : (_2346_v15));
          let _2347_v16;
          _2347_v16 = _dafny.Seq.UnicodeFromString("asuvdd");
          _2347_v16 = _2347_v16;
          let _2348_v17;
          _2348_v17 = new _dafny.CodePoint('s'.codePointAt(0));
          let _2349_v18;
          _2349_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(356)), ((_2350_v17) => function (_2351_i2) {
            return _2350_v17;
          })(_2348_v17))).length),p0);
          let _2352_v19;
          _2352_v19 = new BigNumber(335);
          _2348_v17 = (_2347_v16)[_module.__default.safeIndex(((p0) ? (new BigNumber((_2349_v18).length)) : (_2352_v19)), new BigNumber((_2347_v16).length))];
          let _2353_v20;
          let _nw433 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _2353_v20 = _nw433;
          let _2354_v21;
          _2354_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("nmetnnr"),(_dafny.ZERO).minus(_2352_v19));
          let _index399 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_2353_v20).length));
          (_2353_v20)[_index399] = _2354_v21;
          let _index400 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_2353_v20).length));
          (_2353_v20)[_index400] = _2354_v21;
          let _2355_v22;
          _2355_v22 = _dafny.Seq.of(p0, p0, p0, (_module.D14.create_DC40(p0)).dtor_cf71, p0);
          let _2356_v23;
          _2356_v23 = _dafny.MultiSet.fromElements(p0);
          let _2357_v24;
          let _nw434 = new _module.C6();
          _nw434.__ctor();
          _2357_v24 = _nw434;
          let _2358_v25;
          _2358_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2357_v24);
          let _2359_v26;
          let _nw435 = Array((new BigNumber(15)).toNumber());
          _nw435[(_dafny.ZERO).toNumber()] = _module.__default.fm29(p0, _2352_v19, _2352_v19, (_2355_v22)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_module.__default.fm12(_2356_v23, _2352_v19, globalState), _module.__default.safeIndex(_2352_v19, new BigNumber((_module.__default.fm12(_2356_v23, _2352_v19, globalState)).length)), _2352_v19)).length), new BigNumber((_2355_v22).length))], globalState);
          _nw435[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
          _nw435[(new BigNumber(2)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(3)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(4)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(5)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(6)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
          _nw435[(new BigNumber(8)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(9)).toNumber()] = _module.__default.fm29(p0, new BigNumber((_2358_v25).length), _2352_v19, p0, globalState);
          _nw435[(new BigNumber(10)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(11)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(12)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(13)).toNumber()] = _2348_v17;
          _nw435[(new BigNumber(14)).toNumber()] = _2348_v17;
          _2359_v26 = _nw435;
          let _rhs320 = _2359_v26;
          let _rhs321 = _2347_v16;
          _2359_v26 = _rhs320;
          _2347_v16 = _rhs321;
        }
        let _2360_v27;
        let _nw436 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
        _2360_v27 = _nw436;
        let _source22 = _module.D9.create_DC22(_2360_v27);
        if (_source22.is_DC23) {
          let _2361___mcc_h0 = (_source22).cf37;
          let _2362___mcc_h1 = (_source22).cf38;
          let _2363___mcc_h2 = (_source22).cf39;
          let _2364___mcc_h3 = (_source22).cf40;
          let _2365___mcc_h4 = (_source22).cf41;
          let _2366_cf41 = _2365___mcc_h4;
          let _2367_cf40 = _2364___mcc_h3;
          let _2368_cf39 = _2363___mcc_h2;
          let _2369_cf38 = _2362___mcc_h1;
          let _2370_cf37 = _2361___mcc_h0;
          let _2371_v28;
          let _nw437 = Array((new BigNumber(6)).toNumber()).fill([]);
          _2371_v28 = _nw437;
          let _2372_v29;
          _2372_v29 = _dafny.Seq.of(_2371_v28, _2371_v28, _2371_v28);
          _2371_v28 = (_2372_v29)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_2372_v29).length))];
          (globalState).f1 = p0;
          let _2373_v30;
          _2373_v30 = _dafny.MultiSet.fromElements(_2369_cf38);
          let _index401 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_2367_cf40).length));
          (_2367_cf40)[_index401] = new BigNumber(((_2373_v30).update(_2369_cf38, _module.__default.abs(_module.__default.fm1(_2368_cf39, globalState)))).cardinality());
          let _2374_v31;
          _2374_v31 = new BigNumber(133);
          let _2375_v32;
          _2375_v32 = _dafny.Seq.of(_2374_v31, _2374_v31);
          let _index402 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_2367_cf40).length));
          (_2367_cf40)[_index402] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_2375_v32)[_module.__default.safeIndex(_2374_v31, new BigNumber((_2375_v32).length))], _2374_v31));
          let _2376_v33;
          _2376_v33 = _dafny.MultiSet.fromElements(p0);
          let _2377_v34;
          _2377_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2376_v33,false);
          _2377_v34 = (_2377_v34).update(_2376_v33, p0);
        } else if (_source22.is_DC24) {
          let _2378___mcc_h5 = (_source22).cf42;
          let _2379___mcc_h6 = (_source22).cf43;
          let _2380_cf43 = _2379___mcc_h6;
          let _2381_cf42 = _2378___mcc_h5;
          let _2382_v36;
          let _init67 = function (_2383_i3) {
            return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("k")).length),new BigNumber(-215))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(50),new BigNumber((function () {
              let _coll79 = new _dafny.Map();
              for (const _compr_79 of _dafny.IntegerRange(new BigNumber(819), new BigNumber(-438))) {
                let _2384_v35 = _compr_79;
                if (((new BigNumber(819)).isLessThanOrEqualTo(_2384_v35)) && ((_2384_v35).isLessThan(new BigNumber(-438)))) {
                  _coll79.push([(_2384_v35).minus(new BigNumber(-571)),new BigNumber(890)]);
                }
              }
              return _coll79;
            }()).length)));
          };
          let _nw438 = Array((new BigNumber(26)).toNumber());
          for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw438.length); _i0_67++) {
            _nw438[_i0_67] = _init67(new BigNumber(_i0_67));
          }
          _2382_v36 = _nw438;
          let _2385_v37;
          _2385_v37 = _dafny.MultiSet.fromElements(_module.__default.fm1(!(p0), globalState));
          let _2386_v38;
          _2386_v38 = new BigNumber(-416);
          let _2387_v39;
          _2387_v39 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.MultiSet.fromElements(_2386_v38)).cardinality())).isLessThan((((_2385_v37).contains(_2386_v38)) ? ((_2385_v37).get(_2386_v38)) : (_2386_v38))),_2386_v38);
          let _rhs322 = new BigNumber((_2387_v39).length);
          let _rhs323 = p0;
          let _rhs324 = _2382_v36;
          let _rhs325 = _2386_v38;
          let _lhs283 = globalState;
          let _lhs284 = globalState;
          let _lhs285 = globalState;
          _lhs283.f6 = _rhs322;
          _lhs284.f1 = _rhs323;
          _2382_v36 = _rhs324;
          _lhs285.f6 = _rhs325;
          let _2388_v40;
          _2388_v40 = _dafny.Seq.UnicodeFromString("ced");
          let _2389_v41;
          _2389_v41 = _module.D8.create_DC21(_2388_v40, _dafny.Set.fromElements(p0));
          let _2390_v42;
          _2390_v42 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat((_2389_v41).dtor_cf34, _2388_v40),(_dafny.ZERO).minus(_2386_v38));
          let _2391_v43;
          _2391_v43 = _dafny.Set.fromElements(_2388_v40);
          _2390_v42 = (_2390_v42).update(_2388_v40, new BigNumber((_2391_v43).length));
          let _2392_v44;
          let _init68 = ((_2393_p0) => function (_2394_i4) {
            return !(!(_2393_p0));
          })(p0);
          let _nw439 = Array((new BigNumber(22)).toNumber());
          for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw439.length); _i0_68++) {
            _nw439[_i0_68] = _init68(new BigNumber(_i0_68));
          }
          _2392_v44 = _nw439;
          let _index403 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2392_v44).length));
          (_2392_v44)[_index403] = p0;
          let _index404 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2392_v44).length));
          (_2392_v44)[_index404] = (_2386_v38).isLessThan(_2386_v38);
          let _index405 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2392_v44).length));
          (_2392_v44)[_index405] = !((_2387_v39).Merge(_2387_v39)).equals(_2387_v39);
        } else if (_source22.is_DC22) {
          let _2395___mcc_h7 = (_source22).cf36;
          let _2396_cf36 = _2395___mcc_h7;
          let _2397_v45;
          let _nw440 = Array((new BigNumber(3)).toNumber()).fill(false);
          _2397_v45 = _nw440;
          let _index406 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2397_v45).length));
          (_2397_v45)[_index406] = _module.__default.fm2(globalState);
          let _2398_v46;
          _2398_v46 = new BigNumber(865);
          let _index407 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_2397_v45).length));
          (_2397_v45)[_index407] = (_module.__default.fm1(true, globalState)).isLessThan(_2398_v46);
          let _2399_v47;
          let _nw441 = new _module.C4();
          _nw441.__ctor();
          _2399_v47 = _nw441;
          (globalState).f6 = (_dafny.ZERO).minus((_module.D23.create_DC55(_2398_v46)).dtor_cf94);
          let _2400_v48;
          _2400_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2397_v45,_2397_v45);
          let _2401_v49;
          _2401_v49 = _dafny.Seq.of((((_2400_v48).contains(_2397_v45)) ? ((_2400_v48).get(_2397_v45)) : (_2397_v45)));
          _2397_v45 = ((p0) ? (((p0) ? (_2397_v45) : (_2397_v45))) : ((_2401_v49)[_module.__default.safeIndex(_2398_v46, new BigNumber((_2401_v49).length))]));
        } else {
          let _2402___mcc_h8 = (_source22).cf44;
          let _2403_cf44 = _2402___mcc_h8;
          let _2404_v50;
          _2404_v50 = new _dafny.CodePoint('h'.codePointAt(0));
          _2404_v50 = _2404_v50;
          let _2405_v51;
          let _nw442 = new _module.C6();
          _nw442.__ctor();
          _2405_v51 = _nw442;
          let _2406_v52;
          let _nw443 = Array((new BigNumber(22)).toNumber());
          _2406_v52 = _nw443;
          let _2407_v53;
          _2407_v53 = _dafny.Seq.of(p0);
          let _2408_v54;
          _2408_v54 = new BigNumber(981);
          let _2409_v55;
          _2409_v55 = _module.D11.create_DC31(_2407_v53, _2408_v54, new BigNumber((_2407_v53).length), _2408_v54, new BigNumber(982));
          let _2410_v56;
          _2410_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_2409_v55).dtor_cf56,p0),_2408_v54);
          let _2411_v57;
          let _nw444 = new _module.C3();
          _nw444.__ctor(p0, new BigNumber(345), _2410_v56);
          _2411_v57 = _nw444;
          let _index408 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_2406_v52).length));
          (_2406_v52)[_index408] = _2411_v57;
          let _index409 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_2406_v52).length));
          (_2406_v52)[_index409] = _2411_v57;
          let _2412_v58;
          _2412_v58 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(p0),_2408_v54);
          (globalState).f6 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_2408_v54, new BigNumber((_2412_v58).length)), _2408_v54);
        }
        let _2413_v59;
        _2413_v59 = new BigNumber(-464);
        (globalState).f6 = _2413_v59;
        _2413_v59 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_2413_v59));
        let _2414_v60;
        _2414_v60 = _dafny.Set.fromElements(_2413_v59);
        let _2415_v61;
        _2415_v61 = _dafny.MultiSet.fromElements(_2414_v60, (_2414_v60).Intersect(_2414_v60));
        _2415_v61 = _2415_v61;
      }
      let _2416_v62;
      _2416_v62 = new BigNumber(994);
      let _2417_v63;
      _2417_v63 = _dafny.Seq.of(_2416_v62, _2416_v62);
      let _hi11 = (_2416_v62).minus(new BigNumber(-67));
      for (let _2418_i5 = new BigNumber((_dafny.Seq.Concat(_2417_v63, _2417_v63)).length); _2418_i5.isLessThan(_hi11); _2418_i5 = _2418_i5.plus(_dafny.ONE)) {
        if (!(p0)) {
          let _2419_v64;
          _2419_v64 = new _dafny.CodePoint('n'.codePointAt(0));
          let _2420_v65;
          _2420_v65 = _dafny.MultiSet.fromElements(_2416_v62);
          (globalState).f1 = ((p0) ? ((_2420_v65).IsProperSubsetOf(_module.__default.fm42(_2416_v62, _2419_v64, _2416_v62, p0, globalState))) : (p0));
          let _2421_v66;
          let _init69 = function (_2422_i6) {
            return _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC0(true))))));
          };
          let _nw445 = Array((new BigNumber(22)).toNumber());
          for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw445.length); _i0_69++) {
            _nw445[_i0_69] = _init69(new BigNumber(_i0_69));
          }
          _2421_v66 = _nw445;
          let _2423_v67;
          let _init70 = ((_2424_i5) => function (_2425_i7) {
            return (_2425_i7).plus(_2424_i5);
          })(_2418_i5);
          let _nw446 = Array((new BigNumber(2)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw446.length); _i0_70++) {
            _nw446[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2423_v67 = _nw446;
          _module.__default.m0(_dafny.Set.fromElements(_2418_i5, _2418_i5), _2421_v66, _2423_v67, _dafny.Seq.UnicodeFromString("unfypjx"), globalState);
          let _2426_v68;
          _2426_v68 = _dafny.Set.fromElements((_dafny.ZERO).minus(_2416_v62), _2418_i5);
          let _2427_v69;
          _2427_v69 = _dafny.Seq.UnicodeFromString("nxjyeoe");
          _module.__default.m0(_2426_v68, _2421_v66, _2423_v67, _2427_v69, globalState);
          (globalState).f1 = p0;
          (globalState).f6 = (new BigNumber(542)).plus((_dafny.ZERO).minus(new BigNumber((function () {
            let _coll80 = new _dafny.Map();
            for (const _compr_80 of _dafny.IntegerRange(new BigNumber(490), new BigNumber(853))) {
              let _2428_v70 = _compr_80;
              if (((new BigNumber(490)).isLessThanOrEqualTo(_2428_v70)) && ((_2428_v70).isLessThan(new BigNumber(853)))) {
                _coll80.push([_module.__default.safeDivisionInt(_2428_v70, _2416_v62),p0]);
              }
            }
            return _coll80;
          }()).length)));
        } else {
          let _2429_v71;
          _2429_v71 = new _dafny.CodePoint('j'.codePointAt(0));
          _2429_v71 = _2429_v71;
          let _2430_v72;
          _2430_v72 = _dafny.Set.fromElements((_dafny.ZERO).minus(_2418_i5));
          _2430_v72 = _2430_v72;
          (globalState).f1 = p0;
          let _2431_v73;
          _2431_v73 = _dafny.MultiSet.fromElements(_2429_v71, _2429_v71, _2429_v71);
          _2417_v63 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-148)), ((_2432_v63) => function (_2433_i8) {
            return new BigNumber((_2432_v63).length);
          })(_2417_v63)), _2417_v63), _module.__default.safeIndex(new BigNumber((_2431_v73).cardinality()), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-148)), ((_2434_v63) => function (_2435_i8) {
            return new BigNumber((_2434_v63).length);
          })(_2417_v63)), _2417_v63)).length)), (_2416_v62).plus(_2418_i5));
          _2416_v62 = _2418_i5;
        }
        let _2436_v74;
        _2436_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2418_i5,p0);
        let _2437_v75;
        _2437_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2436_v74,new BigNumber((_dafny.Seq.of(!(true), p0)).length));
        let _2438_v76;
        let _nw447 = new _module.C3();
        _nw447.__ctor(false, new BigNumber(-476), _2437_v75);
        _2438_v76 = _nw447;
        let _2439_v77;
        _2439_v77 = _dafny.Seq.of(_2438_v76);
        let _2440_v78;
        _2440_v78 = _dafny.Seq.of(_2439_v77);
        let _2441_v79;
        _2441_v79 = _dafny.Map.Empty.slice().updateUnsafe(p0,(new BigNumber((_2440_v78).length)).isLessThan(_2418_i5));
        _2441_v79 = (_2441_v79).update(true, false);
        let _2442_v80;
        _2442_v80 = new _dafny.CodePoint('k'.codePointAt(0));
        let _2443_v81;
        _2443_v81 = _dafny.Seq.of(_2442_v80, _2442_v80, _2442_v80);
        let _2444_v82;
        _2444_v82 = _dafny.Seq.of(p0, p0, p0, (_module.D5.create_DC12((_2438_v76).f19, p0, _2442_v80)).dtor_cf16, (_2438_v76).f18);
        let _2445_v83;
        _2445_v83 = _dafny.Seq.of(_2443_v81, _2443_v81);
        let _2446_v84;
        let _nw448 = Array((new BigNumber(29)).toNumber());
        _nw448[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_2416_v62, _2418_i5);
        _nw448[(_dafny.ONE).toNumber()] = _2416_v62;
        _nw448[(new BigNumber(2)).toNumber()] = ((p0) ? (new BigNumber((_module.__default.fm31(_2416_v62, true, globalState)).length)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2438_v76).f18,_2416_v62)).length)));
        _nw448[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_2418_i5);
        _nw448[(new BigNumber(4)).toNumber()] = _2416_v62;
        _nw448[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_2418_i5);
        _nw448[(new BigNumber(6)).toNumber()] = new BigNumber(-829);
        _nw448[(new BigNumber(7)).toNumber()] = _2416_v62;
        _nw448[(new BigNumber(8)).toNumber()] = _2416_v62;
        _nw448[(new BigNumber(9)).toNumber()] = _2418_i5;
        _nw448[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(_2418_i5, _2416_v62);
        _nw448[(new BigNumber(11)).toNumber()] = (_2418_i5).multipliedBy(_2416_v62);
        _nw448[(new BigNumber(12)).toNumber()] = _2418_i5;
        _nw448[(new BigNumber(13)).toNumber()] = new BigNumber(461);
        _nw448[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_2443_v81)).cardinality());
        _nw448[(new BigNumber(15)).toNumber()] = (_2438_v76).f19;
        _nw448[(new BigNumber(16)).toNumber()] = new BigNumber((_module.__default.fm28(_2416_v62, (_2438_v76).f18, true, globalState)).length);
        _nw448[(new BigNumber(17)).toNumber()] = new BigNumber(-899);
        _nw448[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt(_2418_i5, _2418_i5);
        _nw448[(new BigNumber(19)).toNumber()] = new BigNumber((_2444_v82).length);
        _nw448[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt((_2438_v76).f19, _2418_i5);
        _nw448[(new BigNumber(21)).toNumber()] = _2418_i5;
        _nw448[(new BigNumber(22)).toNumber()] = (_2438_v76).f19;
        _nw448[(new BigNumber(23)).toNumber()] = _2418_i5;
        _nw448[(new BigNumber(24)).toNumber()] = new BigNumber(((_2445_v83)[_module.__default.safeIndex(_2418_i5, new BigNumber((_2445_v83).length))]).length);
        _nw448[(new BigNumber(25)).toNumber()] = _2418_i5;
        _nw448[(new BigNumber(26)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_2416_v62), new BigNumber(673));
        _nw448[(new BigNumber(27)).toNumber()] = _2416_v62;
        _nw448[(new BigNumber(28)).toNumber()] = _2418_i5;
        _2446_v84 = _nw448;
        (globalState).f9 = _2446_v84;
        (globalState).f1 = p0;
      }
      let _2447_v85;
      let _nw449 = Array((new BigNumber(17)).toNumber());
      _nw449[(_dafny.ZERO).toNumber()] = true;
      _nw449[(_dafny.ONE).toNumber()] = (_this).fm8(_2416_v62, _2416_v62, (_dafny.ZERO).minus(_2416_v62), globalState);
      _nw449[(new BigNumber(2)).toNumber()] = p0;
      _nw449[(new BigNumber(3)).toNumber()] = p0;
      _nw449[(new BigNumber(4)).toNumber()] = p0;
      _nw449[(new BigNumber(5)).toNumber()] = p0;
      _nw449[(new BigNumber(6)).toNumber()] = p0;
      _nw449[(new BigNumber(7)).toNumber()] = p0;
      _nw449[(new BigNumber(8)).toNumber()] = p0;
      _nw449[(new BigNumber(9)).toNumber()] = true;
      _nw449[(new BigNumber(10)).toNumber()] = p0;
      _nw449[(new BigNumber(11)).toNumber()] = _module.__default.fm2(globalState);
      _nw449[(new BigNumber(12)).toNumber()] = p0;
      _nw449[(new BigNumber(13)).toNumber()] = false;
      _nw449[(new BigNumber(14)).toNumber()] = p0;
      _nw449[(new BigNumber(15)).toNumber()] = p0;
      _nw449[(new BigNumber(16)).toNumber()] = true;
      _2447_v85 = _nw449;
      let _2448_v86;
      _2448_v86 = _dafny.MultiSet.fromElements(_2447_v85);
      let _hi12 = new BigNumber((_2448_v86).cardinality());
      for (let _2449_i9 = _2416_v62; _2449_i9.isLessThan(_hi12); _2449_i9 = _2449_i9.plus(_dafny.ONE)) {
        let _2450_v87;
        let _init71 = ((_2451_v62) => function (_2452_i10) {
          return (_2452_i10).minus(_2451_v62);
        })(_2416_v62);
        let _nw450 = Array((new BigNumber(25)).toNumber());
        for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw450.length); _i0_71++) {
          _nw450[_i0_71] = _init71(new BigNumber(_i0_71));
        }
        _2450_v87 = _nw450;
        let _index410 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_2450_v87).length));
        (_2450_v87)[_index410] = ((p0) ? (_2416_v62) : (_2449_i9));
        let _index411 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_2450_v87).length));
        (_2450_v87)[_index411] = _2449_i9;
        let _2453_v88;
        _2453_v88 = _dafny.Seq.of(_module.__default.fm2(globalState));
        _2416_v62 = _module.__default.safeModuloInt(new BigNumber((_2453_v88).length), _2449_i9);
        let _2454_v90;
        _2454_v90 = new _dafny.CodePoint('o'.codePointAt(0));
        let _2455_v91;
        _2455_v91 = _dafny.Set.fromElements(_module.__default.fm2(globalState), p0, p0);
        let _2456_v92;
        _2456_v92 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll81 = new _dafny.Map();
          for (const _compr_81 of (_dafny.Map.Empty.slice().updateUnsafe(_2454_v90,p0)).Keys.Elements) {
            let _2457_v89 = _compr_81;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_2454_v90,p0)).contains(_2457_v89)) {
              _coll81.push([_2457_v89,new BigNumber((_dafny.Set.fromElements(p0)).length)]);
            }
          }
          return _coll81;
        }(),(_2455_v91).IsProperSubsetOf(_2455_v91));
        let _2458_v93;
        _2458_v93 = _dafny.Map.Empty.slice().updateUnsafe(_2454_v90,(_2450_v87)[_module.__default.safeIndex(new BigNumber(357), new BigNumber((_2450_v87).length))]);
        _2456_v92 = (_2456_v92).update(_2458_v93, ((p0) ? (p0) : (p0)));
        let _2459_v94;
        let _nw451 = new _module.C5();
        _nw451.__ctor(_2416_v62);
        _2459_v94 = _nw451;
        let _2460_v95;
        _2460_v95 = _dafny.Map.Empty.slice().updateUnsafe(_2459_v94,_2449_i9);
        let _2461_v96;
        _2461_v96 = _dafny.Set.fromElements(_2417_v63, _2417_v63);
        let _2462_v97;
        _2462_v97 = _dafny.Map.Empty.slice().updateUnsafe(p0,((_2459_v94).f16).isLessThan((((_2460_v95).contains(_2459_v94)) ? ((_2460_v95).get(_2459_v94)) : (new BigNumber((_2461_v96).length)))));
        _2462_v97 = (_2462_v97).update(!(p0), p0);
      }
      let _2463_v98;
      let _nw452 = new _module.C1();
      _nw452.__ctor((_dafny.ZERO).minus(_2416_v62));
      _2463_v98 = _nw452;
      let _2464_v99;
      _2464_v99 = _dafny.Map.Empty.slice().updateUnsafe(_2463_v98,p0);
      _2464_v99 = (_2464_v99).Merge(_2464_v99);
      let _2465_v100;
      let _nw453 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _2465_v100 = _nw453;
      let _index412 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_2465_v100).length));
      (_2465_v100)[_index412] = _2416_v62;
      let _2466_v101;
      let _nw454 = new _module.C4();
      _nw454.__ctor();
      _2466_v101 = _nw454;
      let _2467_v102;
      _2467_v102 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2417_v63).length),_2466_v101);
      let _index413 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_2465_v100).length));
      (_2465_v100)[_index413] = new BigNumber((_dafny.Seq.of((((_2467_v102).contains(_2416_v62)) ? ((_2467_v102).get(_2416_v62)) : (_2466_v101)))).length);
      let _2468_v103;
      _2468_v103 = _dafny.Seq.of(p0, p0, p0);
      let _2469_v104;
      _2469_v104 = new _dafny.CodePoint('v'.codePointAt(0));
      let _2470_v105;
      _2470_v105 = _dafny.MultiSet.fromElements(true, p0, p0, p0, true);
      let _2471_v106;
      _2471_v106 = _dafny.Map.Empty.slice().updateUnsafe(_2470_v105,_2447_v85);
      let _rhs326 = ((!((p0) && (p0))) ? (_2468_v103) : (_module.__default.fm13(_module.__default.fm2(globalState), _2469_v104, _2416_v62, globalState)));
      let _rhs327 = _module.__default.safeDivisionInt((_2417_v63)[_module.__default.safeIndex(new BigNumber((_2471_v106).length), new BigNumber((_2417_v63).length))], (_2465_v100)[_module.__default.safeIndex(new BigNumber(802), new BigNumber((_2465_v100).length))]);
      _2468_v103 = _rhs326;
      _2416_v62 = _rhs327;
      return;
    }
    m7(globalState) {
      let _this = this;
      let _2472_v0;
      _2472_v0 = true;
      let _2473_i0;
      _2473_i0 = _dafny.ZERO;
      L15: {
        while (((_2472_v0) ? (_2472_v0) : (_2472_v0))) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2473_i0)) {
              break L15;
            }
            _2473_i0 = (_2473_i0).plus(_dafny.ONE);
            let _2474_v1;
            _2474_v1 = new BigNumber(-499);
            (globalState).f1 = (_this).fm8(_2474_v1, _2474_v1, new BigNumber(710), globalState);
            let _2475_v2;
            _2475_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2474_v1,_2474_v1);
            let _2476_v3;
            _2476_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2472_v0,_2472_v0);
            let _2477_v4;
            _2477_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2476_v3).length),_2472_v0);
            _2475_v2 = (_2475_v2).update(_2474_v1, new BigNumber((_2477_v4).length));
            let _2478_v5;
            _2478_v5 = _dafny.Set.fromElements(true);
            let _2479_v6;
            _2479_v6 = _dafny.MultiSet.fromElements((_2478_v5).IsSubsetOf(_dafny.Set.fromElements(_2472_v0)), _2472_v0);
            _2479_v6 = _2479_v6;
            (globalState).f1 = !((((_2476_v3).contains(_2472_v0)) ? ((_2476_v3).get(_2472_v0)) : (_2472_v0)));
          }
        }
      }
      let _2480_v7;
      _2480_v7 = _dafny.Set.fromElements(new BigNumber(554));
      let _2481_v8;
      _2481_v8 = new BigNumber(867);
      let _2482_v9;
      _2482_v9 = _dafny.Seq.of(_2481_v8, _2481_v8);
      let _2483_v11;
      _2483_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2481_v8,_2472_v0);
      (globalState).f6 = (((_2480_v7).IsDisjointFrom(function () {
        let _coll82 = new _dafny.Set();
        for (const _compr_82 of (_2482_v9).Elements) {
          let _2484_v10 = _compr_82;
          if (_dafny.Seq.contains(_2482_v9, _2484_v10)) {
            _coll82.add((_2484_v10).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(104))).length)));
          }
        }
        return _coll82;
      }())) ? (new BigNumber(-388)) : (new BigNumber((_2483_v11).length)));
      let _rhs328 = _module.__default.safeDivisionInt(new BigNumber(13), (_dafny.ZERO).minus(_2481_v8));
      let _rhs329 = (_2472_v0) && (_2472_v0);
      _2481_v8 = _rhs328;
      _2472_v0 = _rhs329;
      (globalState).f6 = new BigNumber(-335);
      let _2485_v12;
      _2485_v12 = _dafny.MultiSet.fromElements(_module.__default.fm2(globalState));
      _2485_v12 = _2485_v12;
      if (false) {
        let _2486_v13;
        _2486_v13 = _dafny.MultiSet.fromElements(new BigNumber(678));
        let _2487_v14;
        _2487_v14 = _dafny.Seq.of(_2472_v0);
        (globalState).f1 = (new BigNumber((_2480_v7).length)).isEqualTo(new BigNumber(((_2486_v13).Union(_dafny.MultiSet.fromElements(_2481_v8, _2481_v8, _2481_v8, new BigNumber((_2487_v14).length), _2481_v8))).cardinality()));
        let _2488_v15;
        _2488_v15 = new _dafny.CodePoint('d'.codePointAt(0));
        let _2489_v16;
        _2489_v16 = _dafny.MultiSet.fromElements(_2488_v15);
        let _2490_v17;
        _2490_v17 = _dafny.Map.Empty.slice().updateUnsafe(!(_2472_v0),(_2489_v16).Difference(_2489_v16));
        _2490_v17 = _2490_v17;
        (globalState).f6 = _2481_v8;
        let _2491_v18;
        _2491_v18 = _dafny.Seq.UnicodeFromString("dla");
        let _2492_v19;
        _2492_v19 = _dafny.Map.Empty.slice().updateUnsafe(((((_2485_v12).contains(_2472_v0)) ? ((_2485_v12).get(_2472_v0)) : (_2481_v8))).multipliedBy(new BigNumber((_2491_v18).length)),_dafny.Seq.UnicodeFromString("ug"));
        let _2493_v20;
        _2493_v20 = _dafny.Seq.of(_2491_v18);
        _2491_v18 = (((_2492_v19).contains((_2481_v8).minus(_2481_v8))) ? ((_2492_v19).get((_2481_v8).minus(_2481_v8))) : ((_2493_v20)[_module.__default.safeIndex(_2481_v8, new BigNumber((_2493_v20).length))]));
        let _2494_v21;
        let _nw455 = new _module.C5();
        _nw455.__ctor(_2481_v8);
        _2494_v21 = _nw455;
      } else {
        let _rhs330 = (_2480_v7).Union(_2480_v7);
        let _rhs331 = _module.__default.fm1(_2472_v0, globalState);
        let _lhs286 = globalState;
        _2480_v7 = _rhs330;
        _lhs286.f6 = _rhs331;
        let _2495_v22;
        _2495_v22 = _dafny.Seq.UnicodeFromString("mserh");
        _2495_v22 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2495_v22, _2495_v22), _2495_v22);
        let _2496_v23;
        _2496_v23 = new _dafny.CodePoint('k'.codePointAt(0));
        let _2497_v24;
        _2497_v24 = _dafny.MultiSet.fromElements(_2481_v8, new BigNumber((_dafny.Seq.of(_2472_v0)).length));
        let _2498_v25;
        _2498_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2481_v8,_2496_v23);
        let _2499_v26;
        _2499_v26 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_2496_v23,_2481_v8),new BigNumber((_2498_v25).length));
        let _2500_v27;
        _2500_v27 = _dafny.Map.Empty.slice().updateUnsafe(_2472_v0,_2496_v23);
        let _2501_v28;
        _2501_v28 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(125)), ((_2502_v24) => function (_2503_i1) {
          return new BigNumber((_2502_v24).cardinality());
        })(_2497_v24)),new _dafny.CodePoint('o'.codePointAt(0)));
        let _2504_v29;
        let _nw456 = Array((new BigNumber(20)).toNumber());
        _nw456[(_dafny.ZERO).toNumber()] = _2496_v23;
        _nw456[(_dafny.ONE).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(2)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(3)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(4)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(5)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(6)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(7)).toNumber()] = _module.__default.fm29(!(_2472_v0), _2481_v8, (((_2497_v24).contains(_2481_v8)) ? ((_2497_v24).get(_2481_v8)) : (new BigNumber((_2499_v26).length))), _2472_v0, globalState);
        _nw456[(new BigNumber(8)).toNumber()] = (((_2500_v27).contains(_2472_v0)) ? ((_2500_v27).get(_2472_v0)) : (_2496_v23));
        _nw456[(new BigNumber(9)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(10)).toNumber()] = (((_2501_v28).contains(_2482_v9)) ? ((_2501_v28).get(_2482_v9)) : (_2496_v23));
        _nw456[(new BigNumber(11)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
        _nw456[(new BigNumber(13)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(14)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw456[(new BigNumber(16)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(17)).toNumber()] = _2496_v23;
        _nw456[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _nw456[(new BigNumber(19)).toNumber()] = _2496_v23;
        _2504_v29 = _nw456;
        let _index414 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_2504_v29).length));
        (_2504_v29)[_index414] = _2496_v23;
        let _index415 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_2504_v29).length));
        (_2504_v29)[_index415] = _2496_v23;
        let _2505_v30;
        _2505_v30 = _module.D12.create_DC35();
        _2505_v30 = _2505_v30;
        let _2506_v31;
        let _nw457 = Array((new BigNumber(24)).toNumber()).fill(false);
        _2506_v31 = _nw457;
        let _index416 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_2506_v31).length));
        (_2506_v31)[_index416] = _2472_v0;
        let _index417 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_2506_v31).length));
        (_2506_v31)[_index417] = _2472_v0;
      }
      return;
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
    fm6(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(-572), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(871),new BigNumber(529))).length));
    };
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((true) === (false),!(((false) ? (!(false)) : (true))));
    };
    m5(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = false;
      let _2507_v0;
      _2507_v0 = false;
      let _2508_v1;
      _2508_v1 = _module.D0.create_DC0(_2507_v0);
      let _2509_v2;
      _2509_v2 = _module.D0.create_DC2(_2508_v1);
      let _2510_v3;
      _2510_v3 = _module.D0.create_DC2(_2508_v1);
      let _2511_v4;
      _2511_v4 = _module.D0.create_DC2(_module.D0.create_DC2(_2508_v1));
      _2511_v4 = _2511_v4;
      let _2512_v5;
      _2512_v5 = new BigNumber(844);
      let _2513_i0;
      _2513_i0 = _dafny.ZERO;
      L16: {
        while (((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(_2507_v0))).length)))).isLessThan(_2512_v5)) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2513_i0)) {
              break L16;
            }
            _2513_i0 = (_2513_i0).plus(_dafny.ONE);
            let _2514_v6;
            _2514_v6 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-939)),_2507_v0);
            let _2515_v7;
            _2515_v7 = _dafny.Seq.of((new BigNumber((_2514_v6).length)).isLessThanOrEqualTo(_2512_v5));
            let _2516_v8;
            _2516_v8 = _dafny.Seq.of(_2512_v5);
            _2507_v0 = !(!(!((_2515_v7)[_module.__default.safeIndex((_2516_v8)[_module.__default.safeIndex(_2512_v5, new BigNumber((_2516_v8).length))], new BigNumber((_2515_v7).length))])));
            let _2517_v9;
            let _nw458 = Array((new BigNumber(14)).toNumber());
            _nw458[(_dafny.ZERO).toNumber()] = _2507_v0;
            _nw458[(_dafny.ONE).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(2)).toNumber()] = !(_2507_v0);
            _nw458[(new BigNumber(3)).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(4)).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(5)).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(6)).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(7)).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(8)).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(9)).toNumber()] = false;
            _nw458[(new BigNumber(10)).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(11)).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(12)).toNumber()] = _2507_v0;
            _nw458[(new BigNumber(13)).toNumber()] = _2507_v0;
            _2517_v9 = _nw458;
            let _2518_v10;
            _2518_v10 = _dafny.Seq.UnicodeFromString("umolis");
            let _2519_v11;
            _2519_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2517_v9,new BigNumber((_dafny.MultiSet.fromElements(_2512_v5, new BigNumber((_2518_v10).length), _2512_v5)).cardinality()));
            let _2520_v12;
            _2520_v12 = _dafny.Map.Empty.slice().updateUnsafe(!(_2507_v0),new BigNumber(868));
            let _2521_v13;
            _2521_v13 = _dafny.Map.Empty.slice().updateUnsafe((((_2519_v11).contains(_2517_v9)) ? ((_2519_v11).get(_2517_v9)) : (_2512_v5)),_2520_v12);
            _2521_v13 = (_2521_v13).update(_2512_v5, (_dafny.Map.Empty.slice().updateUnsafe(_2507_v0,new BigNumber(845))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2507_v0,_2512_v5)));
            _2512_v5 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("lxvm")).length), new BigNumber((_2515_v7).length));
            let _2522_v14;
            _2522_v14 = _dafny.MultiSet.fromElements(_2507_v0, _2507_v0);
            let _2523_v15;
            _2523_v15 = _dafny.MultiSet.fromElements(_2522_v14, _2522_v14, (_2522_v14).update(_2507_v0, _module.__default.abs(_2512_v5)), _2522_v14);
            let _2524_v16;
            let _nw459 = new _module.C5();
            _nw459.__ctor(new BigNumber((_2523_v15).cardinality()));
            _2524_v16 = _nw459;
          }
        }
      }
      let _2525_v17;
      _2525_v17 = _dafny.Seq.UnicodeFromString("jkeuhyxcb");
      r1 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(334)), function (_2526_i1) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }), _2525_v17);
      let _2527_v18;
      let _init72 = ((_2528_v0) => function (_2529_i3) {
        return _2528_v0;
      })(_2507_v0);
      let _nw460 = Array((new BigNumber(28)).toNumber());
      for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw460.length); _i0_72++) {
        _nw460[_i0_72] = _init72(new BigNumber(_i0_72));
      }
      _2527_v18 = _nw460;
      let _2530_v19;
      _2530_v19 = _dafny.MultiSet.fromElements(_2512_v5);
      let _hi13 = (new BigNumber((_2530_v19).cardinality())).plus(_2512_v5);
      for (let _2531_i2 = (_2512_v5).minus(new BigNumber((_dafny.MultiSet.fromElements(_2527_v18)).cardinality())); _2531_i2.isLessThan(_hi13); _2531_i2 = _2531_i2.plus(_dafny.ONE)) {
        if (_2507_v0) {
          (globalState).f6 = (_dafny.ZERO).minus(_2512_v5);
          let _2532_v20;
          _2532_v20 = _dafny.Seq.of(_2512_v5);
          let _2533_v21;
          _2533_v21 = new _dafny.CodePoint('c'.codePointAt(0));
          let _2534_v22;
          _2534_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2533_v21,new BigNumber((_module.__default.fm21(globalState)).length));
          let _2535_v23;
          _2535_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2534_v22).length),_2507_v0);
          let _2536_v24;
          let _nw461 = new _module.C2();
          _nw461.__ctor((new BigNumber((_2532_v20).length)).minus(new BigNumber((_2532_v20).length)), _module.__default.safeDivisionInt(new BigNumber(280), _2512_v5), _dafny.Map.Empty.slice().updateUnsafe(_2535_v23,_2512_v5));
          _2536_v24 = _nw461;
          (globalState).f6 = (_this).fm6((_2507_v0) === (!(_2507_v0)), new BigNumber((_2525_v17).length), globalState);
          let _2537_v25;
          _2537_v25 = _dafny.Set.fromElements(_2512_v5);
          _2533_v21 = _module.__default.fm29(((_2532_v20)[_module.__default.safeIndex((_2532_v20)[_module.__default.safeIndex(_2536_v24.f21, new BigNumber((_2532_v20).length))], new BigNumber((_2532_v20).length))]).isLessThanOrEqualTo(_2512_v5), new BigNumber((_2532_v20).length), _2536_v24.f21, (_2537_v25).IsSubsetOf(_2537_v25), globalState);
          let _rhs332 = ((_2507_v0) ? ((_2536_v24).f20) : (_2512_v5));
          let _rhs333 = _2507_v0;
          r0 = _rhs332;
          _2507_v0 = _rhs333;
        } else {
          (globalState).f6 = _2531_i2;
          _2507_v0 = _2507_v0;
          let _2538_v26;
          let _nw462 = Array((new BigNumber(13)).toNumber());
          _nw462[(_dafny.ZERO).toNumber()] = _2527_v18;
          _nw462[(_dafny.ONE).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(2)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(3)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(4)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(5)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(6)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(7)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(8)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(9)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(10)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(11)).toNumber()] = _2527_v18;
          _nw462[(new BigNumber(12)).toNumber()] = _2527_v18;
          _2538_v26 = _nw462;
          let _index418 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_2538_v26).length));
          (_2538_v26)[_index418] = _2527_v18;
          let _2539_v27;
          _2539_v27 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_2527_v18);
          let _index419 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_2538_v26).length));
          (_2538_v26)[_index419] = (((_2539_v27).contains((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2507_v0,_2507_v0)).length)).isLessThanOrEqualTo(_2512_v5))) ? ((_2539_v27).get((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2507_v0,_2507_v0)).length)).isLessThanOrEqualTo(_2512_v5))) : (_2527_v18));
          let _2540_v28;
          _2540_v28 = _dafny.Map.Empty.slice().updateUnsafe(((_2507_v0) ? (!(_2507_v0)) : (_2507_v0)),_2507_v0);
          _2540_v28 = (_2540_v28).update(_2507_v0, _2507_v0);
          let _2541_v29;
          let _nw463 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _2541_v29 = _nw463;
          let _index420 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_2541_v29).length));
          (_2541_v29)[_index420] = _2512_v5;
          let _index421 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2541_v29).length));
          (_2541_v29)[_index421] = (_dafny.ZERO).minus((_this).fm6(!(_2507_v0), _2531_i2, globalState));
          let _index422 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_2527_v18).length));
          (_2527_v18)[_index422] = (_2512_v5).isLessThanOrEqualTo((_dafny.ZERO).minus(_2512_v5));
          let _index423 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_2541_v29).length));
          let _index424 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2541_v29).length));
          let _index425 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_2527_v18).length));
          let _rhs334 = _2531_i2;
          let _rhs335 = _2531_i2;
          let _rhs336 = (_2512_v5).minus(_2512_v5);
          let _rhs337 = !(_2507_v0);
          let _rhs338 = _2531_i2;
          let _lhs287 = _2541_v29;
          let _lhs288 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_2541_v29).length));
          let _lhs289 = _2541_v29;
          let _lhs290 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2541_v29).length));
          let _lhs291 = globalState;
          let _lhs292 = _2527_v18;
          let _lhs293 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_2527_v18).length));
          _lhs287[_lhs288] = _rhs334;
          _lhs289[_lhs290] = _rhs335;
          _lhs291.f6 = _rhs336;
          _lhs292[_lhs293] = _rhs337;
          _2512_v5 = _rhs338;
        }
        r0 = _2512_v5;
        let _2542_v30;
        _2542_v30 = _dafny.Seq.of(_2512_v5);
        let _2543_v31;
        _2543_v31 = _dafny.Set.fromElements(_2531_i2, new BigNumber((_2542_v30).length));
        let _2544_v32;
        _2544_v32 = _module.D15.create_DC42(_2543_v31);
        _2544_v32 = _module.D15.create_DC42((_2543_v31).Difference(_2543_v31));
        _2543_v31 = _module.__default.fm28(new BigNumber((_dafny.Seq.UnicodeFromString("wdyvu")).length), _2507_v0, _2507_v0, globalState);
      }
      let _2545_v33;
      let _nw464 = new _module.C6();
      _nw464.__ctor();
      _2545_v33 = _nw464;
      let _2546_v34;
      let _nw465 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Set.Empty);
      _2546_v34 = _nw465;
      let _2547_v35;
      _2547_v35 = _dafny.MultiSet.fromElements(true);
      let _2548_v36;
      _2548_v36 = _dafny.Set.fromElements((((_2547_v35).contains(_2507_v0)) ? ((_2547_v35).get(_2507_v0)) : (_2512_v5)));
      let _2549_v37;
      _2549_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2546_v34,_2548_v36);
      _2549_v37 = (_2549_v37).update(_2546_v34, _2548_v36);
      r0 = (_2512_v5).multipliedBy(_module.__default.safeModuloInt(_2512_v5, _2512_v5));
      let _2550_v38;
      _2550_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_2525_v17).length)),_2548_v36);
      let _pat_let_tv61 = _2525_v17;
      r1 = function (_source23) {
        let _2551___mcc_h0 = _source23;
        let _2552_cf19 = _2551___mcc_h0;
        return _pat_let_tv61;
      }(_2550_v38);
      let _2553_v39;
      _2553_v39 = _dafny.Seq.of(_2507_v0);
      r2 = !(!((_2553_v39)[_module.__default.safeIndex(_2512_v5, new BigNumber((_2553_v39).length))]) || ((_2507_v0) || (_2507_v0)));
      return [r0, r1, r2];
    }
    m6(p0, p1, globalState) {
      let _this = this;
      let _2554_v0;
      _2554_v0 = new BigNumber(150);
      let _2555_v1;
      _2555_v1 = _dafny.Seq.of(_2554_v0);
      let _2556_v2;
      _2556_v2 = _module.D10.create_DC28(new BigNumber((_2555_v1).length));
      let _hi14 = new BigNumber(-270);
      for (let _2557_i0 = _module.__default.safeDivisionInt(_2554_v0, (_2556_v2).dtor_cf51); _2557_i0.isLessThan(_hi14); _2557_i0 = _2557_i0.plus(_dafny.ONE)) {
        let _2558_v3;
        let _nw466 = new _module.C5();
        _nw466.__ctor(_2554_v0);
        _2558_v3 = _nw466;
        let _2559_v4;
        _2559_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.of(_2557_i0));
        _2555_v1 = _dafny.Seq.Concat((((_2559_v4).contains(p0)) ? ((_2559_v4).get(p0)) : (_2555_v1)), _2555_v1);
        let _2560_v5;
        let _nw467 = new _module.C1();
        _nw467.__ctor(_2554_v0);
        _2560_v5 = _nw467;
        if (p0) {
          let _2561_v6;
          _2561_v6 = new _dafny.CodePoint('c'.codePointAt(0));
          let _2562_v7;
          let _nw468 = Array((new BigNumber(8)).toNumber());
          _nw468[(_dafny.ZERO).toNumber()] = _2561_v6;
          _nw468[(_dafny.ONE).toNumber()] = _module.__default.fm29(p0, (_2558_v3).f16, (_2560_v5).f24, p0, globalState);
          _nw468[(new BigNumber(2)).toNumber()] = _2561_v6;
          _nw468[(new BigNumber(3)).toNumber()] = _2561_v6;
          _nw468[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _nw468[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('t'.codePointAt(0));
          _nw468[(new BigNumber(6)).toNumber()] = _2561_v6;
          _nw468[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
          _2562_v7 = _nw468;
          let _index426 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_2562_v7).length));
          (_2562_v7)[_index426] = _2561_v6;
          let _index427 = _module.__default.safeIndex(new BigNumber(702), new BigNumber((_2562_v7).length));
          (_2562_v7)[_index427] = new _dafny.CodePoint('a'.codePointAt(0));
          let _2563_v8;
          let _init73 = ((_2564_v0) => function (_2565_i1) {
            return _module.__default.safeDivisionInt(_2565_i1, _2564_v0);
          })(_2554_v0);
          let _nw469 = Array((new BigNumber(29)).toNumber());
          for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw469.length); _i0_73++) {
            _nw469[_i0_73] = _init73(new BigNumber(_i0_73));
          }
          _2563_v8 = _nw469;
          let _2566_v9;
          _2566_v9 = _module.D3.create_DC6(_2563_v8);
          let _2567_v10;
          _2567_v10 = _dafny.Seq.of(_2566_v9, _2566_v9, _2566_v9, _2566_v9, _module.D3.create_DC6(_2563_v8));
          let _2568_v11;
          _2568_v11 = _dafny.Seq.UnicodeFromString("iivhhe");
          let _2569_v12;
          let _nw470 = new _module.C0();
          _nw470.__ctor(_dafny.Seq.Concat(_2567_v10, _2567_v10), (_2557_i0).isLessThanOrEqualTo(new BigNumber((_2568_v11).length)));
          _2569_v12 = _nw470;
          let _2570_v13;
          _2570_v13 = _dafny.MultiSet.fromElements((_this).fm6(p0, new BigNumber((_2555_v1).length), globalState));
          let _2571_v14;
          _2571_v14 = _2570_v13;
          let _2572_v15;
          let _nw471 = Array((new BigNumber(7)).toNumber());
          _nw471[(_dafny.ZERO).toNumber()] = _2570_v13;
          _nw471[(_dafny.ONE).toNumber()] = _2570_v13;
          _nw471[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.FromArray(_2555_v1)).Union(_2570_v13);
          _nw471[(new BigNumber(3)).toNumber()] = _2570_v13;
          _nw471[(new BigNumber(4)).toNumber()] = (_2571_v14);
          _nw471[(new BigNumber(5)).toNumber()] = _2570_v13;
          _nw471[(new BigNumber(6)).toNumber()] = _2570_v13;
          _2572_v15 = _nw471;
          _2572_v15 = _2572_v15;
          let _2573_v16;
          _2573_v16 = _dafny.Seq.of(_2570_v13, _2570_v13);
          let _2574_v17;
          _2574_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2557_i0,(_2573_v16)[_module.__default.safeIndex((_2560_v5).f24, new BigNumber((_2573_v16).length))]);
          (globalState).f6 = (((_this).fm6((_2569_v12).f23, new BigNumber((_2555_v1).length), globalState)).plus(new BigNumber((_2574_v17).length))).multipliedBy(((_2558_v3).f16).multipliedBy(_2557_i0));
          let _2575_v18;
          let _nw472 = new _module.C6();
          _nw472.__ctor();
          _2575_v18 = _nw472;
        } else {
          let _2576_v19;
          _2576_v19 = _dafny.Seq.UnicodeFromString("flk");
          _2576_v19 = _dafny.Seq.UnicodeFromString("olvqfpuoo");
          let _2577_v20;
          let _nw473 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _2577_v20 = _nw473;
          let _2578_v21;
          _2578_v21 = _module.D3.create_DC6(_2577_v20);
          let _2579_v22;
          _2579_v22 = _dafny.Seq.of(_module.D3.create_DC6(_2577_v20), _2578_v21);
          let _2580_v23;
          let _nw474 = new _module.C0();
          _nw474.__ctor(_2579_v22, p0);
          _2580_v23 = _nw474;
          let _2581_v24;
          let _nw475 = Array((new BigNumber(3)).toNumber()).fill([]);
          _2581_v24 = _nw475;
          let _2582_v25;
          _2582_v25 = _module.D5.create_DC10(_2581_v24);
          let _2583_v26;
          _2583_v26 = _module.D7.create_DC17(p0, ((p0) ? (p1) : (p1)), _2580_v23, _2582_v25, (p1) && (p0));
          let _2584_v27;
          _2584_v27 = _dafny.Map.Empty.slice().updateUnsafe(!((_2580_v23).f23),_2576_v19);
          _2583_v26 = _module.D7.create_DC17(_dafny.Seq.IsProperPrefixOf(_2576_v19, (((_2584_v27).contains(true)) ? ((_2584_v27).get(true)) : (_2576_v19))), ((_2580_v23).f23) && (p1), _2580_v23, _2582_v25, true);
          _2554_v0 = (_2558_v3).f16;
          let _2585_v28;
          _2585_v28 = _dafny.Set.fromElements((_2580_v23).f23, p0);
          let _2586_v30;
          _2586_v30 = _dafny.Map.Empty.slice().updateUnsafe(true,(_2558_v3).f16);
          let _2587_v31;
          _2587_v31 = _dafny.MultiSet.fromElements(new BigNumber((_2586_v30).length));
          let _2588_v32;
          _2588_v32 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ttqx")).length), (_2558_v3).f16, (((_2587_v31).contains(_2554_v0)) ? ((_2587_v31).get(_2554_v0)) : ((_2558_v3).f16)), _module.__default.fm1(p1, globalState));
          let _2589_v33;
          let _nw476 = Array((new BigNumber(15)).toNumber());
          _nw476[(_dafny.ZERO).toNumber()] = !_dafny.areEqual(_2576_v19, _2576_v19);
          _nw476[(_dafny.ONE).toNumber()] = !(p1);
          _nw476[(new BigNumber(2)).toNumber()] = p0;
          _nw476[(new BigNumber(3)).toNumber()] = (_2580_v23).f23;
          _nw476[(new BigNumber(4)).toNumber()] = (_2580_v23).fm17((_2560_v5).f24, _2554_v0, p0, globalState);
          _nw476[(new BigNumber(5)).toNumber()] = true;
          _nw476[(new BigNumber(6)).toNumber()] = ((_2558_v3).f16).isLessThan((_2558_v3).f16);
          _nw476[(new BigNumber(7)).toNumber()] = p0;
          _nw476[(new BigNumber(8)).toNumber()] = !((_2580_v23).f23);
          _nw476[(new BigNumber(9)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(326)), function (_2590_i2) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }), _2576_v19);
          _nw476[(new BigNumber(10)).toNumber()] = (_2580_v23).f23;
          _nw476[(new BigNumber(11)).toNumber()] = _module.__default.fm2(globalState);
          _nw476[(new BigNumber(12)).toNumber()] = (_2580_v23).f23;
          _nw476[(new BigNumber(13)).toNumber()] = !(!(_2585_v28).contains(false));
          _nw476[(new BigNumber(14)).toNumber()] = !((function () {
            let _coll83 = new _dafny.Set();
            for (const _compr_83 of _dafny.IntegerRange(new BigNumber(136), new BigNumber(970))) {
              let _2591_v29 = _compr_83;
              if (((new BigNumber(136)).isLessThanOrEqualTo(_2591_v29)) && ((_2591_v29).isLessThan(new BigNumber(970)))) {
                _coll83.add(_module.__default.safeModuloInt(_2591_v29, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(408))).cardinality())));
              }
            }
            return _coll83;
          }()).IsDisjointFrom(_2588_v32));
          _2589_v33 = _nw476;
          let _2592_v34;
          _2592_v34 = _dafny.Seq.of(p1);
          let _index428 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_2589_v33).length));
          (_2589_v33)[_index428] = _dafny.areEqual(_2592_v34, _2592_v34);
          let _index429 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_2589_v33).length));
          (_2589_v33)[_index429] = (_2557_i0).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_2558_v3).f16), _2554_v0));
          (globalState).f6 = (_2560_v5).f24;
        }
      }
      _2554_v0 = new BigNumber(-736);
      let _2593_v35;
      _2593_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2554_v0,new BigNumber(-10));
      let _2594_v36;
      let _nw477 = new _module.C7();
      _nw477.__ctor();
      _2594_v36 = _nw477;
      let _2595_v37;
      _2595_v37 = _dafny.MultiSet.fromElements(_2594_v36, _2594_v36, _2594_v36, _2594_v36, _2594_v36);
      let _2596_v38;
      _2596_v38 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2554_v0);
      let _2597_v39;
      _2597_v39 = _dafny.Seq.of(!(p1), false);
      let _2598_v40;
      _2598_v40 = _dafny.Seq.UnicodeFromString("n");
      let _2599_v41;
      _2599_v41 = _dafny.MultiSet.fromElements(_2554_v0, _2554_v0, new BigNumber(21));
      let _2600_v42;
      _2600_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2554_v0,_2599_v41);
      let _2601_v43;
      _2601_v43 = _dafny.Set.fromElements(_2554_v0, _2554_v0);
      let _2602_v44;
      let _nw478 = Array((new BigNumber(19)).toNumber());
      _nw478[(_dafny.ZERO).toNumber()] = (_module.D0.create_DC1(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("qpdpge")).length))).length))).dtor_cf1;
      _nw478[(_dafny.ONE).toNumber()] = _2554_v0;
      _nw478[(new BigNumber(2)).toNumber()] = (((_2593_v35).contains((_dafny.ZERO).minus(_2554_v0))) ? ((_2593_v35).get((_dafny.ZERO).minus(_2554_v0))) : (_2554_v0));
      _nw478[(new BigNumber(3)).toNumber()] = _2554_v0;
      _nw478[(new BigNumber(4)).toNumber()] = _2554_v0;
      _nw478[(new BigNumber(5)).toNumber()] = _2554_v0;
      _nw478[(new BigNumber(6)).toNumber()] = (((_2595_v37).contains(_2594_v36)) ? ((_2595_v37).get(_2594_v36)) : (_2554_v0));
      _nw478[(new BigNumber(7)).toNumber()] = (((_2596_v38).contains(false)) ? ((_2596_v38).get(false)) : (new BigNumber(646)));
      _nw478[(new BigNumber(8)).toNumber()] = _2554_v0;
      _nw478[(new BigNumber(9)).toNumber()] = (_2554_v0).multipliedBy(new BigNumber((_2597_v39).length));
      _nw478[(new BigNumber(10)).toNumber()] = new BigNumber(-468);
      _nw478[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_2554_v0);
      _nw478[(new BigNumber(12)).toNumber()] = (_2554_v0).plus(_2554_v0);
      _nw478[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_2598_v40, _2598_v40)).length));
      _nw478[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((((_2599_v41).contains(_2554_v0)) ? ((_2599_v41).get(_2554_v0)) : (_2554_v0)));
      _nw478[(new BigNumber(15)).toNumber()] = _2554_v0;
      _nw478[(new BigNumber(16)).toNumber()] = (new BigNumber((_2600_v42).length)).multipliedBy(_2554_v0);
      _nw478[(new BigNumber(17)).toNumber()] = new BigNumber((_2601_v43).length);
      _nw478[(new BigNumber(18)).toNumber()] = _2554_v0;
      _2602_v44 = _nw478;
      let _index430 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length));
      (_2602_v44)[_index430] = _2554_v0;
      let _2603_v45;
      _2603_v45 = new _dafny.CodePoint('m'.codePointAt(0));
      let _index431 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length));
      (_2602_v44)[_index431] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2598_v40, _2598_v40), _module.__default.safeIndex(_2554_v0, new BigNumber((_dafny.Seq.Concat(_2598_v40, _2598_v40)).length)), _2603_v45)).length);
      let _2604_v46;
      let _nw479 = Array((new BigNumber(5)).toNumber());
      _2604_v46 = _nw479;
      let _2605_v47;
      _2605_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2554_v0,p0);
      let _2606_v48;
      _2606_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2605_v47,(_2602_v44)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length))]);
      let _2607_v49;
      let _nw480 = new _module.C2();
      _nw480.__ctor(_2554_v0, new BigNumber((_2555_v1).length), _2606_v48);
      _2607_v49 = _nw480;
      let _index432 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_2604_v46).length));
      (_2604_v46)[_index432] = _2607_v49;
      let _index433 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_2604_v46).length));
      (_2604_v46)[_index433] = _2607_v49;
      let _hi15 = (_2602_v44)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length))];
      for (let _2608_i3 = _module.__default.safeModuloInt(new BigNumber((function () {
        let _coll85 = new _dafny.Set();
        for (const _compr_85 of _dafny.IntegerRange(new BigNumber(426), new BigNumber(668))) {
          let _2615_v50 = _compr_85;
          if (((new BigNumber(426)).isLessThanOrEqualTo(_2615_v50)) && ((_2615_v50).isLessThan(new BigNumber(668)))) {
            _coll85.add(_module.__default.safeModuloInt(_2615_v50, (_2602_v44)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length))]));
          }
        }
        return _coll85;
      }()).length), _2554_v0); _2608_i3.isLessThan(_hi15); _2608_i3 = _2608_i3.plus(_dafny.ONE)) {
        let _2609_v51;
        let _nw481 = Array((new BigNumber(7)).toNumber());
        _nw481[(_dafny.ZERO).toNumber()] = _2598_v40;
        _nw481[(_dafny.ONE).toNumber()] = _2598_v40;
        _nw481[(new BigNumber(2)).toNumber()] = ((true) ? (_dafny.Seq.update(_2598_v40, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("io")).length), new BigNumber((_2598_v40).length)), new _dafny.CodePoint('g'.codePointAt(0)))) : (_2598_v40));
        _nw481[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(504)), function (_2610_i4) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        });
        _nw481[(new BigNumber(4)).toNumber()] = _2598_v40;
        _nw481[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_2598_v40, _module.__default.fm21(globalState));
        _nw481[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("n");
        _2609_v51 = _nw481;
        _2609_v51 = _2609_v51;
        (globalState).f6 = _2554_v0;
        let _2611_v52;
        let _nw482 = new _module.C7();
        _nw482.__ctor();
        _2611_v52 = _nw482;
        let _2612_v53;
        _2612_v53 = _dafny.Map.Empty.slice().updateUnsafe((_2602_v44)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length))],_2596_v38);
        let _2613_v55;
        _2613_v55 = _module.D25.create_DC57(function () {
  let _coll84 = new _dafny.Map();
  for (const _compr_84 of _dafny.IntegerRange(new BigNumber(-418), new BigNumber(-301))) {
    let _2614_v54 = _compr_84;
    if (((new BigNumber(-418)).isLessThanOrEqualTo(_2614_v54)) && ((_2614_v54).isLessThan(new BigNumber(-301)))) {
      _coll84.push([_module.__default.safeDivisionInt(_2614_v54, (_2602_v44)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length))]),_2555_v1]);
    }
  }
  return _coll84;
}());
        _2612_v53 = (_2612_v53).update((_dafny.ZERO).minus(new BigNumber((_module.__default.fm21(globalState)).length)), (_2596_v38).Merge(_module.__default.fm14((_2613_v55).dtor_cf96, _2554_v0, true, globalState)));
      }
      let _2616_i5;
      _2616_i5 = _dafny.ZERO;
      L17: {
        while (p0) {
          C17: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2616_i5)) {
              break L17;
            }
            _2616_i5 = (_2616_i5).plus(_dafny.ONE);
            (globalState).f6 = ((_2602_v44)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length))]).plus((_2602_v44)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length))]);
            let _2617_v56;
            let _nw483 = new _module.C2();
            _nw483.__ctor(_module.__default.fm1((_2597_v39)[_module.__default.safeIndex((_2602_v44)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length))], new BigNumber((_2597_v39).length))], globalState), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_2554_v0,new BigNumber((_2555_v1).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2554_v0,_2554_v0))).length), (_2606_v48).Merge(_2607_v49.f17));
            _2617_v56 = _nw483;
            let _2618_v58;
            _2618_v58 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p0);
            (globalState).f1 = !((function () {
              let _coll86 = new _dafny.Set();
              for (const _compr_86 of _dafny.IntegerRange(new BigNumber(78), new BigNumber(315))) {
                let _2619_v57 = _compr_86;
                if (((new BigNumber(78)).isLessThanOrEqualTo(_2619_v57)) && ((_2619_v57).isLessThan(new BigNumber(315)))) {
                  _coll86.add((_2619_v57).plus(new BigNumber((_2598_v40).length)));
                }
              }
              return _coll86;
            }()).Difference(_dafny.Set.fromElements(new BigNumber((_2618_v58).length)))).contains(_2617_v56.f21);
            let _2620_v59;
            _2620_v59 = _dafny.Seq.of(_2617_v56.f21);
            let _index434 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length));
            let _rhs339 = (_2602_v44)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length))];
            let _rhs340 = ((_2617_v56).fm15(globalState)).minus((_dafny.ZERO).minus(new BigNumber(((_2620_v59)).length)));
            let _lhs294 = globalState;
            let _lhs295 = _2602_v44;
            let _lhs296 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_2602_v44).length));
            _lhs294.f6 = _rhs339;
            _lhs295[_lhs296] = _rhs340;
          }
        }
      }
      return;
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f15 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f15) {
      let _this = this;
      (_this)._f15 = f15;
      return;
    }
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source24 = _module.D0.create_DC0(!(true));
      if (_source24.is_DC1) {
        let _2621___mcc_h0 = (_source24).cf1;
        let _2622_cf1 = _2621___mcc_h0;
        return _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1((_this).f15)));
      } else if (_source24.is_DC0) {
        let _2623___mcc_h1 = (_source24).cf0;
        let _2624_cf0 = _2623___mcc_h1;
        if (_2624_cf0) {
          return _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1((_dafny.ZERO).minus((_this).f15))))));
        } else {
          return _module.D0.create_DC2(_module.D0.create_DC1((_this).f15));
        }
      } else {
        let _2625___mcc_h2 = (_source24).cf2;
        let _2626_cf2 = _2625___mcc_h2;
        return _module.D0.create_DC2(_module.D0.create_DC1((_this).f15));
      }
    };
    fm4(p0, globalState) {
      let _this = this;
      return (((_module.D0.create_DC1((_this).f15)).dtor_cf1).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), function (_2627_i0) {
        return (_this).f15;
      })).length))).isLessThanOrEqualTo((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f15),true)).length)).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f15))));
    };
    fm5(p0, globalState) {
      let _this = this;
      return ((_this).f15).plus((_this).f15);
    };
    m1(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _2628_v0;
      _2628_v0 = false;
      let _2629_i0;
      _2629_i0 = _dafny.ZERO;
      L18: {
        while (_2628_v0) {
          C18: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2629_i0)) {
              break L18;
            }
            _2629_i0 = (_2629_i0).plus(_dafny.ONE);
            let _2630_v1;
            _2630_v1 = _dafny.Seq.of(_2628_v0, _2628_v0);
            let _2631_v2;
            let _2632_v3;
            let _2633_v4;
            let _out18;
            let _out19;
            let _out20;
            let _outcollector7 = (_this).m3(_dafny.Seq.Concat(_2630_v1, _2630_v1), globalState);
            _out18 = _outcollector7[0];
            _out19 = _outcollector7[1];
            _out20 = _outcollector7[2];
            _2631_v2 = _out18;
            _2632_v3 = _out19;
            _2633_v4 = _out20;
            let _2634_v5;
            _2634_v5 = _dafny.Set.fromElements(_2632_v3, true, _2631_v2, _2628_v0);
            let _2635_v6;
            _2635_v6 = _dafny.Seq.of(new BigNumber((_2634_v5).length), _module.__default.safeModuloInt(new BigNumber(781), (_dafny.ZERO).minus(new BigNumber((_2630_v1).length))), ((_this).f15).plus(_2633_v4), (_this).f15);
            _2635_v6 = _dafny.Seq.of(new BigNumber(((_dafny.MultiSet.FromArray(_2630_v1)).update(_2631_v2, _module.__default.abs((_dafny.ZERO).minus(_2633_v4)))).cardinality()), (_dafny.ZERO).minus((_dafny.ZERO).minus(_2633_v4)), ((_2631_v2) ? ((_this).f15) : ((_this).f15)));
            let _2636_v7;
            let _nw484 = new _module.C5();
            _nw484.__ctor((_this).f15);
            _2636_v7 = _nw484;
            let _2637_v8;
            _2637_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(561),(_this).f15);
            _2637_v8 = (_2637_v8).update((_dafny.ZERO).minus((_this).f15), (_this).f15);
          }
        }
      }
      let _2638_v9;
      _2638_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2628_v0,_2628_v0);
      let _2639_v10;
      _2639_v10 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_2628_v0,_2628_v0), _2638_v9);
      let _hi16 = (_this).f15;
      for (let _2640_i1 = (_dafny.ZERO).minus(new BigNumber(((_2639_v10).Union(_2639_v10)).cardinality())); _2640_i1.isLessThan(_hi16); _2640_i1 = _2640_i1.plus(_dafny.ONE)) {
        let _2641_v11;
        _2641_v11 = _dafny.Seq.UnicodeFromString("epptbqs");
        let _2642_v12;
        _2642_v12 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("egwsccse"), _dafny.Seq.UnicodeFromString("scuxx")), _2641_v11);
        let _2643_v13;
        _2643_v13 = _module.D13.create_DC37(_2628_v0, _2628_v0);
        let _2644_v14;
        _2644_v14 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2640_i1),_2628_v0);
        let _rhs341 = ((!(_2628_v0)) ? (_2628_v0) : (false));
        let _rhs342 = (_2628_v0) || (!(!(_2628_v0) || (_2628_v0)));
        let _rhs343 = _2642_v12;
        let _rhs344 = new BigNumber(815);
        let _rhs345 = !((new BigNumber((_2644_v14).length)).isLessThanOrEqualTo(new BigNumber((_module.__default.fm0(!((_2643_v13).dtor_cf66), true, (_this).f15, globalState)).length)));
        let _lhs297 = globalState;
        _2628_v0 = _rhs341;
        _2628_v0 = _rhs342;
        _2642_v12 = _rhs343;
        _lhs297.f6 = _rhs344;
        r1 = _rhs345;
        (globalState).f6 = (_this).f15;
        let _2645_v15;
        _2645_v15 = _dafny.Set.fromElements((_this).f15, new BigNumber((_2641_v11).length), (_this).f15, _2640_i1, new BigNumber(21));
        let _2646_v16;
        _2646_v16 = _module.D0.create_DC2(_module.D0.create_DC2((_module.D0.create_DC2(_module.D0.create_DC1((_this).f15))).dtor_cf2));
        let _2647_v17;
        _2647_v17 = _module.D0.create_DC0(_2628_v0);
        let _2648_v18;
        let _nw485 = Array((new BigNumber(27)).toNumber());
        _nw485[(_dafny.ZERO).toNumber()] = _2646_v16;
        _nw485[(_dafny.ONE).toNumber()] = _module.D0.create_DC2(_2647_v17);
        _nw485[(new BigNumber(2)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(3)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(4)).toNumber()] = _module.D0.create_DC2(_2647_v17);
        _nw485[(new BigNumber(5)).toNumber()] = _module.D0.create_DC2(_2647_v17);
        _nw485[(new BigNumber(6)).toNumber()] = _module.D0.create_DC2(_2647_v17);
        _nw485[(new BigNumber(7)).toNumber()] = _module.D0.create_DC2(_2647_v17);
        _nw485[(new BigNumber(8)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(9)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(10)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(11)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(12)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(13)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(14)).toNumber()] = (_this).fm3(_2628_v0, (_this).f15, _2628_v0, _2628_v0, globalState);
        _nw485[(new BigNumber(15)).toNumber()] = (_this).fm3(!(_2628_v0), _2640_i1, _2628_v0, !(_2628_v0), globalState);
        _nw485[(new BigNumber(16)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(17)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(18)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(19)).toNumber()] = _module.D0.create_DC2(_2647_v17);
        _nw485[(new BigNumber(20)).toNumber()] = (_this).fm3(false, (_this).f15, _2628_v0, _2628_v0, globalState);
        _nw485[(new BigNumber(21)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(22)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(23)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(24)).toNumber()] = _2646_v16;
        _nw485[(new BigNumber(25)).toNumber()] = (_this).fm3(_2628_v0, (_this).f15, _2628_v0, !(_2628_v0), globalState);
        _nw485[(new BigNumber(26)).toNumber()] = _2646_v16;
        _2648_v18 = _nw485;
        let _2649_v19;
        let _init74 = ((_2650_v14) => function (_2651_i2) {
          return (_2651_i2).plus(new BigNumber((_2650_v14).length));
        })(_2644_v14);
        let _nw486 = Array((new BigNumber(6)).toNumber());
        for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw486.length); _i0_74++) {
          _nw486[_i0_74] = _init74(new BigNumber(_i0_74));
        }
        _2649_v19 = _nw486;
        _module.__default.m0(_2645_v15, _2648_v18, _2649_v19, _2641_v11, globalState);
        if (_2628_v0) {
          let _index435 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_2649_v19).length));
          (_2649_v19)[_index435] = _2640_i1;
          let _index436 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_2649_v19).length));
          (_2649_v19)[_index436] = _module.__default.safeModuloInt(_2640_i1, (_2640_i1).minus((_this).f15));
          r1 = _2628_v0;
          let _2652_v20;
          _2652_v20 = new _dafny.CodePoint('y'.codePointAt(0));
          let _2653_v21;
          let _2654_v22;
          let _out21;
          let _out22;
          let _outcollector8 = (_this).m4(_dafny.Seq.contains(_2641_v11, _2652_v20), globalState);
          _out21 = _outcollector8[0];
          _out22 = _outcollector8[1];
          _2653_v21 = _out21;
          _2654_v22 = _out22;
          let _2655_v23;
          let _nw487 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2655_v23 = _nw487;
          let _2656_v24;
          _2656_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2652_v20,_2655_v23);
          _2656_v24 = (_2656_v24).update(_2652_v20, _2655_v23);
          let _index437 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_2649_v19).length));
          (_2649_v19)[_index437] = (_2649_v19)[_module.__default.safeIndex(new BigNumber(476), new BigNumber((_2649_v19).length))];
        } else {
          r1 = (_2645_v15).IsProperSubsetOf(_dafny.Set.fromElements(_2640_i1, _2640_i1, _module.__default.fm1(_2628_v0, globalState)));
          let _2657_v25;
          let _nw488 = new _module.C5();
          _nw488.__ctor(((_this).f15).multipliedBy(_2640_i1));
          _2657_v25 = _nw488;
          let _2658_v26;
          _2658_v26 = _dafny.Seq.of(_2640_i1, _2640_i1, (_2657_v25).f16);
          (globalState).f6 = (_2640_i1).minus((_2658_v26)[_module.__default.safeIndex(_2640_i1, new BigNumber((_2658_v26).length))]);
          let _2659_v27;
          _2659_v27 = new _dafny.CodePoint('p'.codePointAt(0));
          let _rhs346 = (_module.D3.create_DC8(_2658_v26, _2641_v11, _2628_v0, _2659_v27)).dtor_cf9;
          let _rhs347 = _dafny.Seq.Concat(_2658_v26, _2658_v26);
          _2641_v11 = _rhs346;
          _2658_v26 = _rhs347;
          _module.__default.m0(_2645_v15, _2648_v18, _2649_v19, ((!(_2628_v0)) ? (_2641_v11) : (_2641_v11)), globalState);
        }
      }
      let _2660_v28;
      _2660_v28 = _dafny.Seq.of((_this).f15);
      let _2661_v29;
      _2661_v29 = _dafny.Seq.of((new BigNumber((_2660_v28).length)).isLessThan((_this).f15));
      if ((_2661_v29)[_module.__default.safeIndex((_this).f15, new BigNumber((_2661_v29).length))]) {
        let _2662_v30;
        let _init75 = function (_2663_i3) {
          return _module.__default.safeDivisionInt(_2663_i3, (_this).f15);
        };
        let _nw489 = Array((new BigNumber(12)).toNumber());
        for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw489.length); _i0_75++) {
          _nw489[_i0_75] = _init75(new BigNumber(_i0_75));
        }
        _2662_v30 = _nw489;
        let _2664_v31;
        let _nw490 = Array((new BigNumber(24)).toNumber()).fill(false);
        _2664_v31 = _nw490;
        let _2665_v32;
        _2665_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2662_v30,_2664_v31);
        let _pat_let_tv62 = _2662_v30;
        _2665_v32 = (_2665_v32).Merge(_dafny.Map.Empty.slice().updateUnsafe((function (_pat_let87_0) {
          return function (_2666_dt__update__tmp_h0) {
            return function (_pat_let88_0) {
              return function (_2667_dt__update_hcf6_h0) {
                return _module.D3.create_DC6(_2667_dt__update_hcf6_h0);
              }(_pat_let88_0);
            }(_pat_let_tv62);
          }(_pat_let87_0);
        }(_module.D3.create_DC6(_2662_v30))).dtor_cf6,_2664_v31));
        _2638_v9 = (_2638_v9).update(_2628_v0, _2628_v0);
        (globalState).f1 = _2628_v0;
        if (_2628_v0) {
          (globalState).f1 = _2628_v0;
          let _2668_v33;
          let _nw491 = Array((new BigNumber(7)).toNumber());
          _2668_v33 = _nw491;
          let _2669_v34;
          let _nw492 = new _module.C4();
          _nw492.__ctor();
          _2669_v34 = _nw492;
          let _index438 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_2668_v33).length));
          (_2668_v33)[_index438] = _2669_v34;
          let _2670_v35;
          _2670_v35 = _dafny.Seq.UnicodeFromString("wlaqsdmhf");
          let _2671_v36;
          _2671_v36 = _2669_v34;
          let _index439 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_2668_v33).length));
          let _rhs348 = (_2671_v36);
          let _rhs349 = _2628_v0;
          let _rhs350 = _2670_v35;
          let _lhs298 = _2668_v33;
          let _lhs299 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_2668_v33).length));
          let _lhs300 = globalState;
          _lhs298[_lhs299] = _rhs348;
          _lhs300.f1 = _rhs349;
          _2670_v35 = _rhs350;
          let _2672_v37;
          _2672_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
          let _2673_v38;
          _2673_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2672_v37,_2628_v0);
          let _2674_v40;
          _2674_v40 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15),new BigNumber((_2672_v37).length));
          _2628_v0 = !((_2673_v38).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2672_v37,_2628_v0))).equals((function () {
            let _coll87 = new _dafny.Map();
            for (const _compr_87 of (_2674_v40).Keys.Elements) {
              let _2675_v39 = _compr_87;
              if ((_2674_v40).contains(_2675_v39)) {
                _coll87.push([_2675_v39,_2628_v0]);
              }
            }
            return _coll87;
          }()).Merge(_2673_v38));
          _2628_v0 = !(true);
          (globalState).f1 = _2628_v0;
        } else {
          _2660_v28 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(415)), function (_2676_i4) {
            return (_this).f15;
          });
          let _2677_v41;
          _2677_v41 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f15), (_dafny.ZERO).minus((_this).f15));
          (globalState).f6 = (new BigNumber((_2677_v41).length)).multipliedBy((_this).f15);
          let _2678_v42;
          let _nw493 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2678_v42 = _nw493;
          let _2679_v43;
          _2679_v43 = _dafny.Seq.UnicodeFromString("eioyj");
          let _index440 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2678_v42).length));
          (_2678_v42)[_index440] = _2679_v43;
          let _index441 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_2678_v42).length));
          (_2678_v42)[_index441] = _dafny.Seq.UnicodeFromString("sewysnmpc");
          _2679_v43 = _dafny.Seq.UnicodeFromString("udky");
          let _2680_v44;
          let _2681_v45;
          let _2682_v46;
          let _out23;
          let _out24;
          let _out25;
          let _outcollector9 = (_this).m3(_dafny.Seq.of(_2628_v0), globalState);
          _out23 = _outcollector9[0];
          _out24 = _outcollector9[1];
          _out25 = _outcollector9[2];
          _2680_v44 = _out23;
          _2681_v45 = _out24;
          _2682_v46 = _out25;
        }
        r0 = (_2660_v28)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f15, new BigNumber(-890)), new BigNumber((_2660_v28).length))];
      } else {
        let _2683_v47;
        let _nw494 = new _module.C4();
        _nw494.__ctor();
        _2683_v47 = _nw494;
        let _2684_v48;
        _2684_v48 = _dafny.Seq.of(_2683_v47);
        _2638_v9 = (_2638_v9).update(_2628_v0, !((_this).f15).isEqualTo(new BigNumber((_dafny.Seq.update(_2684_v48, _module.__default.safeIndex((_2660_v28)[_module.__default.safeIndex((_this).f15, new BigNumber((_2660_v28).length))], new BigNumber((_2684_v48).length)), _2683_v47)).length)));
        let _2685_v49;
        _2685_v49 = _dafny.Seq.UnicodeFromString("hutlw");
        let _2686_v50;
        _2686_v50 = _dafny.Seq.of(_2661_v29);
        (globalState).f6 = (((_this).f15).plus(new BigNumber((_2685_v49).length))).minus(new BigNumber(((_2686_v50)[_module.__default.safeIndex((_this).f15, new BigNumber((_2686_v50).length))]).length));
        let _2687_v51;
        let _nw495 = Array((new BigNumber(28)).toNumber()).fill([]);
        _2687_v51 = _nw495;
        let _2688_v52;
        let _nw496 = Array((new BigNumber(16)).toNumber()).fill(false);
        _2688_v52 = _nw496;
        let _index442 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_2687_v51).length));
        (_2687_v51)[_index442] = _2688_v52;
        let _index443 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_2687_v51).length));
        (_2687_v51)[_index443] = _2688_v52;
        let _2689_v53;
        let _nw497 = new _module.C4();
        _nw497.__ctor();
        _2689_v53 = _nw497;
        let _2690_v54;
        _2690_v54 = _module.D14.create_DC40(_2628_v0);
        let _2691_v55;
        _2691_v55 = _module.D14.create_DC41(_2690_v54);
        let _2692_v56;
        _2692_v56 = _module.D14.create_DC41(_2691_v55);
        let _2693_v57;
        _2693_v57 = _module.D14.create_DC41(_2691_v55);
        let _2694_v58;
        _2694_v58 = _dafny.MultiSet.fromElements(new BigNumber((_2685_v49).length));
        let _rhs351 = _2689_v53;
        let _rhs352 = _2693_v57;
        let _rhs353 = (new BigNumber((_2694_v58).cardinality())).multipliedBy((_this).f15);
        let _lhs301 = globalState;
        _2689_v53 = _rhs351;
        _2693_v57 = _rhs352;
        _lhs301.f6 = _rhs353;
        let _arr0 = (_2687_v51)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2687_v51).length))];
        let _index444 = _module.__default.safeIndex(new BigNumber(124), new BigNumber(((_2687_v51)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2687_v51).length))]).length));
        _arr0[_index444] = _2628_v0;
        let _2695_v59;
        _2695_v59 = new _dafny.CodePoint('b'.codePointAt(0));
        let _2696_v60;
        _2696_v60 = _dafny.MultiSet.fromElements(_2695_v59);
        let _arr1 = (_2687_v51)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2687_v51).length))];
        let _index445 = _module.__default.safeIndex(new BigNumber(124), new BigNumber(((_2687_v51)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2687_v51).length))]).length));
        _arr1[_index445] = ((_module.__default.fm46(!(_2628_v0), true, _2628_v0, _2628_v0, globalState)).Difference(_2696_v60)).IsProperSubsetOf(_2696_v60);
      }
      let _2697_v61;
      let _init76 = ((_2698_v0) => function (_2699_i5) {
        return _2698_v0;
      })(_2628_v0);
      let _nw498 = Array((new BigNumber(25)).toNumber());
      for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw498.length); _i0_76++) {
        _nw498[_i0_76] = _init76(new BigNumber(_i0_76));
      }
      _2697_v61 = _nw498;
      let _index446 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length));
      (_2697_v61)[_index446] = true;
      let _2700_v62;
      _2700_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2628_v0);
      let _2701_v63;
      _2701_v63 = _module.D16.create_DC44(_dafny.MultiSet.FromArray(_2661_v29));
      let _2702_v64;
      _2702_v64 = _dafny.Seq.of(_2701_v63, _2701_v63);
      let _index447 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length));
      let _rhs354 = _2628_v0;
      let _rhs355 = (_2700_v62).Merge(_2700_v62);
      let _rhs356 = !_dafny.areEqual(_dafny.Seq.of(_2701_v63, _2701_v63), _dafny.Seq.Concat(_2702_v64, _dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), ((_2703_v0) => function (_2704_i6) {
        return _module.D16.create_DC44(_dafny.MultiSet.fromElements(_2703_v0));
      })(_2628_v0))));
      let _lhs302 = _2697_v61;
      let _lhs303 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length));
      _lhs302[_lhs303] = _rhs354;
      _2700_v62 = _rhs355;
      r1 = _rhs356;
      let _2705_v65;
      let _nw499 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _2705_v65 = _nw499;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2705_v65).length))) {
        let _2706_i7 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2706_i7)) && ((_2706_i7).isLessThan(new BigNumber((_2705_v65).length))))) {
          (_2705_v65)[(_2706_i7)] = (_2706_i7).minus((_this).f15);
        }
      }
      let _2707_v66;
      let _nw500 = new _module.C7();
      _nw500.__ctor();
      _2707_v66 = _nw500;
      let _2708_v67;
      _2708_v67 = _dafny.MultiSet.fromElements(_2707_v66);
      let _2709_v68;
      _2709_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2628_v0,_2708_v67);
      let _2710_v69;
      _2710_v69 = _dafny.Seq.of(_2707_v66, _2707_v66);
      let _2711_i8;
      _2711_i8 = _dafny.ZERO;
      L19: {
        while ((((_2697_v61)[_module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length))]) ? (_2628_v0) : ((_dafny.MultiSet.FromArray(_2710_v69)).IsProperSubsetOf((((_2709_v68).contains(_2628_v0)) ? ((_2709_v68).get(_2628_v0)) : (_dafny.MultiSet.fromElements(_2707_v66, _2707_v66))))))) {
          C19: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2711_i8)) {
              break L19;
            }
            _2711_i8 = (_2711_i8).plus(_dafny.ONE);
            let _2712_v70;
            _2712_v70 = _dafny.MultiSet.fromElements(false, (_2697_v61)[_module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length))], (_2697_v61)[_module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length))]);
            let _2713_v71;
            _2713_v71 = _dafny.Seq.of((_dafny.MultiSet.FromArray(_2661_v29)).Union(_2712_v70), _2712_v70, _2712_v70, _2712_v70);
            let _rhs357 = ((false) ? (new BigNumber(456)) : ((_this).f15));
            let _rhs358 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements((_2697_v61)[_module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length))])), _2713_v71), _module.__default.safeIndex(((_this).f15).plus((_this).f15), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements((_2697_v61)[_module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length))])), _2713_v71)).length)), (_2712_v70).Intersect(_2712_v70));
            let _lhs304 = globalState;
            _lhs304.f6 = _rhs357;
            _2713_v71 = _rhs358;
            let _index448 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_2705_v65).length));
            (_2705_v65)[_index448] = (_this).f15;
            let _index449 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_2705_v65).length));
            (_2705_v65)[_index449] = (_module.__default.safeDivisionInt((_this).f15, (_dafny.ZERO).minus((_this).f15))).minus((_this).f15);
            let _2714_v72;
            _2714_v72 = _dafny.Seq.of(_2697_v61);
            r0 = new BigNumber((_dafny.Seq.Concat(_2714_v72, _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_2697_v61, _2697_v61, _2697_v61), _module.__default.safeIndex((_this).f15, new BigNumber((_dafny.Seq.of(_2697_v61, _2697_v61, _2697_v61)).length)), _2697_v61), _module.__default.safeIndex((_this).f15, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_2697_v61, _2697_v61, _2697_v61), _module.__default.safeIndex((_this).f15, new BigNumber((_dafny.Seq.of(_2697_v61, _2697_v61, _2697_v61)).length)), _2697_v61)).length)), _2697_v61))).length);
            let _index450 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length));
            (_2697_v61)[_index450] = !_dafny.Seq.contains(_dafny.Seq.of(true), _2628_v0);
          }
        }
      }
      r0 = (_this).f15;
      r1 = (_2697_v61)[_module.__default.safeIndex(new BigNumber(287), new BigNumber((_2697_v61).length))];
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let _2715_v0;
      let _nw501 = Array((new BigNumber(22)).toNumber()).fill(false);
      _2715_v0 = _nw501;
      let _2716_v1;
      _2716_v1 = _dafny.Set.fromElements(_2715_v0);
      _2716_v1 = _2716_v1;
      let _2717_v2;
      let _nw502 = Array((new BigNumber(18)).toNumber()).fill(_module.D15.Default());
      _2717_v2 = _nw502;
      let _2718_v3;
      _2718_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2717_v2,(_this).f15);
      _2718_v3 = (_2718_v3).update(((p0) ? (_2717_v2) : (_2717_v2)), (_this).f15);
      if (p0) {
        let _2719_v4;
        _2719_v4 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(718)).multipliedBy((_this).f15),!(p0));
        let _2720_v5;
        _2720_v5 = _dafny.Seq.of(p0);
        let _2721_v6;
        _2721_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f15);
        let _2722_v7;
        _2722_v7 = _2721_v6;
        let _2723_v8;
        _2723_v8 = _dafny.Set.fromElements(_2722_v7);
        let _2724_v9;
        _2724_v9 = _dafny.MultiSet.fromElements(p0, p0, (_2720_v5)[_module.__default.safeIndex(new BigNumber((_2723_v8).length), new BigNumber((_2720_v5).length))], p0, p0);
        (globalState).f1 = (((_2719_v4).contains(((_this).f15).multipliedBy((_this).f15))) ? ((_2719_v4).get(((_this).f15).multipliedBy((_this).f15))) : ((_dafny.MultiSet.fromElements(p0, p0)).IsSubsetOf(_2724_v9)));
        let _2725_v10;
        let _nw503 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
        _2725_v10 = _nw503;
        let _2726_v11;
        _2726_v11 = _module.D20.create_DC49(_2725_v10);
        let _2727_v12;
        _2727_v12 = _dafny.Set.fromElements(_2726_v11);
        (globalState).f6 = (_dafny.ZERO).minus(((((_2721_v6).contains(p0)) ? ((_2721_v6).get(p0)) : ((_this).f15))).plus(new BigNumber(((_2727_v12).Union(_2727_v12)).length)));
        if ((_2724_v9).IsProperSubsetOf(_2724_v9)) {
          let _2728_v13;
          _2728_v13 = new _dafny.CodePoint('v'.codePointAt(0));
          _2728_v13 = _2728_v13;
          let _2729_v14;
          _2729_v14 = _dafny.Seq.UnicodeFromString("cevl");
          _2729_v14 = _module.__default.fm21(globalState);
          (globalState).f1 = true;
          let _2730_v15;
          _2730_v15 = _dafny.Seq.of(_2715_v0);
          let _2731_v16;
          _2731_v16 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(p0, globalState),_dafny.Seq.of(_2715_v0));
          let _2732_v17;
          let _nw504 = new _module.C7();
          _nw504.__ctor();
          _2732_v17 = _nw504;
          let _2733_v18;
          _2733_v18 = _dafny.Seq.of(_2732_v17);
          let _2734_v19;
          _2734_v19 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(191)), ((_2735_v13) => function (_2736_i0) {
            return _2735_v13;
          })(_2728_v13))).length), (_this).f15);
          let _2737_v20;
          _2737_v20 = _dafny.MultiSet.fromElements(_2732_v17, (_2733_v18)[_module.__default.safeIndex(new BigNumber((_2734_v19).cardinality()), new BigNumber((_2733_v18).length))]);
          let _2738_v21;
          let _nw505 = Array((new BigNumber(27)).toNumber());
          _nw505[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_2715_v0);
          _nw505[(_dafny.ONE).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(2)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(3)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(4)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(5)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(6)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_2715_v0, _2715_v0);
          _nw505[(new BigNumber(8)).toNumber()] = (((_2731_v16).contains((((_2737_v20).contains(_2732_v17)) ? ((_2737_v20).get(_2732_v17)) : ((_this).f15)))) ? ((_2731_v16).get((((_2737_v20).contains(_2732_v17)) ? ((_2737_v20).get(_2732_v17)) : ((_this).f15)))) : (_2730_v15));
          _nw505[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_2730_v15, _2730_v15);
          _nw505[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_2715_v0), _2730_v15);
          _nw505[(new BigNumber(11)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(12)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(13)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(_2715_v0, _2715_v0, _2715_v0, _2715_v0);
          _nw505[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_2730_v15, _2730_v15);
          _nw505[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_2730_v15, _dafny.Seq.update(_2730_v15, _module.__default.safeIndex((_this).f15, new BigNumber((_2730_v15).length)), _2715_v0));
          _nw505[(new BigNumber(17)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(18)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(19)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(20)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(21)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(22)).toNumber()] = ((p0) ? (_2730_v15) : (_2730_v15));
          _nw505[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_2730_v15, _2730_v15);
          _nw505[(new BigNumber(24)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_2715_v0, _2715_v0, _2715_v0, _2715_v0, _2715_v0), _2730_v15), _module.__default.safeIndex(new BigNumber(380), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_2715_v0, _2715_v0, _2715_v0, _2715_v0, _2715_v0), _2730_v15)).length)), _2715_v0);
          _nw505[(new BigNumber(25)).toNumber()] = _2730_v15;
          _nw505[(new BigNumber(26)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_2730_v15, _module.__default.safeIndex((_this).f15, new BigNumber((_2730_v15).length)), _2715_v0), _2730_v15);
          _2738_v21 = _nw505;
          let _index451 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_2738_v21).length));
          (_2738_v21)[_index451] = ((p0) ? (_2730_v15) : (_2730_v15));
          let _2739_v22;
          let _nw506 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
          _2739_v22 = _nw506;
          let _2740_v23;
          _2740_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2719_v4).length),_2728_v13);
          let _index452 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_2739_v22).length));
          (_2739_v22)[_index452] = (_2740_v23).Merge(function () {
            let _coll88 = new _dafny.Map();
            for (const _compr_88 of _dafny.IntegerRange(new BigNumber(163), new BigNumber(781))) {
              let _2741_v24 = _compr_88;
              if (((new BigNumber(163)).isLessThanOrEqualTo(_2741_v24)) && ((_2741_v24).isLessThan(new BigNumber(781)))) {
                _coll88.push([_module.__default.safeModuloInt(_2741_v24, new BigNumber((_2734_v19).cardinality())),_2728_v13]);
              }
            }
            return _coll88;
          }());
          let _2742_v25;
          let _nw507 = Array((new BigNumber(15)).toNumber()).fill([]);
          _2742_v25 = _nw507;
          let _2743_v26;
          _2743_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2728_v13,(_this).f15);
          let _2744_v27;
          let _nw508 = new _module.C2();
          _nw508.__ctor((_this).f15, (_this).f15, _module.__default.fm47((_this).f15, _module.__default.fm1(p0, globalState), (_this).f15, new BigNumber((_2743_v26).length), globalState));
          _2744_v27 = _nw508;
          let _2745_v28;
          _2745_v28 = _dafny.Seq.of(_2744_v27);
          let _2746_v29;
          let _nw509 = Array((new BigNumber(17)).toNumber()).fill(_module.D3.Default());
          _2746_v29 = _nw509;
          let _2747_v30;
          _2747_v30 = _dafny.Set.fromElements(_2746_v29);
          let _2748_v31;
          _2748_v31 = _dafny.Map.Empty.slice().updateUnsafe(_2720_v5,_2729_v14);
          let _2749_v32;
          _2749_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _2750_v33;
          _2750_v33 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f15));
          let _2751_v34;
          let _nw510 = Array((new BigNumber(11)).toNumber());
          _nw510[(_dafny.ZERO).toNumber()] = (_this).f15;
          _nw510[(_dafny.ONE).toNumber()] = new BigNumber((_2749_v32).length);
          _nw510[(new BigNumber(2)).toNumber()] = new BigNumber((_2724_v9).cardinality());
          _nw510[(new BigNumber(3)).toNumber()] = (_this).f15;
          _nw510[(new BigNumber(4)).toNumber()] = (_this).f15;
          _nw510[(new BigNumber(5)).toNumber()] = (_this).f15;
          _nw510[(new BigNumber(6)).toNumber()] = (_this).f15;
          _nw510[(new BigNumber(7)).toNumber()] = (_this).f15;
          _nw510[(new BigNumber(8)).toNumber()] = (_this).f15;
          _nw510[(new BigNumber(9)).toNumber()] = (_this).f15;
          _nw510[(new BigNumber(10)).toNumber()] = new BigNumber((_2750_v33).cardinality());
          _2751_v34 = _nw510;
          let _2752_v35;
          _2752_v35 = _module.D9.create_DC23(_2747_v30, (((_2748_v31).contains(_2720_v5)) ? ((_2748_v31).get(_2720_v5)) : (_2729_v14)), p0, _2751_v34, new _dafny.CodePoint('l'.codePointAt(0)));
          let _2753_v36;
          let _nw511 = Array((new BigNumber(25)).toNumber());
          _nw511[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _nw511[(_dafny.ONE).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(2)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(3)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(4)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(5)).toNumber()] = (_2729_v14)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(551)), ((_2754_v13) => function (_2755_i1) {
            return _2754_v13;
          })(_2728_v13))).length), new BigNumber((_2729_v14).length))];
          _nw511[(new BigNumber(6)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(7)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(8)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
          _nw511[(new BigNumber(10)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(11)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(12)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(13)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
          _nw511[(new BigNumber(15)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(16)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(17)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(18)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(19)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(20)).toNumber()] = _module.__default.fm29(false, new BigNumber((_dafny.Seq.update(_2745_v28, _module.__default.safeIndex((_this).f15, new BigNumber((_2745_v28).length)), _2744_v27)).length), (((_2743_v26).contains(_2728_v13)) ? ((_2743_v26).get(_2728_v13)) : ((_this).f15)), p0, globalState);
          _nw511[(new BigNumber(21)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(22)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(23)).toNumber()] = _2728_v13;
          _nw511[(new BigNumber(24)).toNumber()] = (_2752_v35).dtor_cf41;
          _2753_v36 = _nw511;
          let _index453 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2742_v25).length));
          (_2742_v25)[_index453] = _2753_v36;
          let _index454 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_2738_v21).length));
          let _index455 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_2739_v22).length));
          let _index456 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2742_v25).length));
          let _rhs359 = _2730_v15;
          let _rhs360 = _2722_v7;
          let _rhs361 = _2728_v13;
          let _rhs362 = (_2740_v23).Merge((_2740_v23).update((_this).f15, _2728_v13));
          let _rhs363 = _2753_v36;
          let _lhs305 = _2738_v21;
          let _lhs306 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_2738_v21).length));
          let _lhs307 = _2739_v22;
          let _lhs308 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_2739_v22).length));
          let _lhs309 = _2742_v25;
          let _lhs310 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_2742_v25).length));
          _lhs305[_lhs306] = _rhs359;
          _2722_v7 = _rhs360;
          _2728_v13 = _rhs361;
          _lhs307[_lhs308] = _rhs362;
          _lhs309[_lhs310] = _rhs363;
          (globalState).f1 = (_2720_v5)[_module.__default.safeIndex((_this).f15, new BigNumber((_2720_v5).length))];
        } else {
          (globalState).f1 = p0;
          let _2756_v37;
          let _nw512 = new _module.C5();
          _nw512.__ctor((_dafny.ZERO).minus((_this).f15));
          _2756_v37 = _nw512;
          let _2757_v38;
          _2757_v38 = _dafny.Seq.UnicodeFromString("u");
          let _2758_v39;
          _2758_v39 = _dafny.Map.Empty.slice().updateUnsafe(_module.D23.create_DC53(_2756_v37),_2757_v38);
          _2758_v39 = _2758_v39;
          (globalState).f14 = _2715_v0;
          let _2759_v40;
          _2759_v40 = new _dafny.CodePoint('u'.codePointAt(0));
          _2759_v40 = _2759_v40;
          (globalState).f6 = new BigNumber(-41);
        }
        let _2760_v41;
        _2760_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2719_v4,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15)).length));
        let _2761_v42;
        let _nw513 = new _module.C3();
        _nw513.__ctor(p0, (_this).f15, _2760_v41);
        _2761_v42 = _nw513;
        _2761_v42 = ((false) ? (((p0) ? (_2761_v42) : (_2761_v42))) : (_2761_v42));
        let _2762_v43;
        let _init77 = function (_2763_i2) {
          return (_2763_i2).plus((_this).f15);
        };
        let _nw514 = Array((new BigNumber(10)).toNumber());
        for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw514.length); _i0_77++) {
          _nw514[_i0_77] = _init77(new BigNumber(_i0_77));
        }
        _2762_v43 = _nw514;
        (globalState).f9 = _2762_v43;
      } else {
        (globalState).f6 = (_this).f15;
        let _2764_v44;
        _2764_v44 = new _dafny.CodePoint('s'.codePointAt(0));
        let _2765_v45;
        _2765_v45 = _dafny.MultiSet.fromElements((_this).f15, (_this).f15, (_this).f15);
        let _pat_let_tv63 = _2765_v45;
        let _pat_let_tv64 = _2765_v45;
        let _pat_let_tv65 = _2764_v44;
        let _source25 = function (_pat_let89_0) {
          return function (_2766_dt__update__tmp_h0) {
            return function (_pat_let90_0) {
              return function (_2767_dt__update_hcf16_h0) {
                return function (_pat_let91_0) {
                  return function (_2768_dt__update_hcf17_h0) {
                    return _module.D5.create_DC12((_2766_dt__update__tmp_h0).dtor_cf15, _2767_dt__update_hcf16_h0, _2768_dt__update_hcf17_h0);
                  }(_pat_let91_0);
                }(_pat_let_tv65);
              }(_pat_let90_0);
            }((_pat_let_tv63).IsSubsetOf(_pat_let_tv64));
          }(_pat_let89_0);
        }(_module.D5.create_DC12((_this).f15, p0, _2764_v44));
        if (_source25.is_DC11) {
          let _2769___mcc_h0 = (_source25).cf14;
          let _2770_cf14 = _2769___mcc_h0;
          let _2771_v46;
          _2771_v46 = _dafny.Map.Empty.slice().updateUnsafe(!(p0) || (true),p0);
          _2771_v46 = (_2771_v46).update(p0, p0);
          let _2772_v47;
          _2772_v47 = _dafny.Set.fromElements(_2770_cf14);
          let _2773_v48;
          _2773_v48 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(p0, globalState),_2772_v47);
          let _2774_v49;
          _2774_v49 = _2773_v48;
          _2774_v49 = _2773_v48;
          _2770_cf14 = ((!(p0) || ((_this).fm4(_2770_cf14, globalState))) ? (((_this).f15).multipliedBy(_2770_cf14)) : ((_this).f15));
          let _2775_v50;
          let _nw515 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _2775_v50 = _nw515;
          let _2776_v51;
          _2776_v51 = _dafny.Seq.of(p0);
          let _index457 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_2775_v50).length));
          (_2775_v50)[_index457] = _module.__default.safeModuloInt(new BigNumber((_2776_v51).length), _2770_cf14);
          let _2777_v52;
          _2777_v52 = _dafny.Set.fromElements(_module.__default.fm2(globalState), p0);
          let _2778_v53;
          _2778_v53 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_2777_v52);
          let _2779_v54;
          _2779_v54 = _module.D0.create_DC0(p0);
          let _2780_v55;
          _2780_v55 = _dafny.Seq.of(_2772_v47, _2772_v47);
          let _2781_v56;
          _2781_v56 = _dafny.Map.Empty.slice().updateUnsafe((_2780_v55)[_module.__default.safeIndex(_2770_cf14, new BigNumber((_2780_v55).length))],(_dafny.ZERO).minus(new BigNumber((_2776_v51).length)));
          let _2782_v57;
          _2782_v57 = _dafny.MultiSet.fromElements(_2777_v52);
          let _index458 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_2775_v50).length));
          let _rhs364 = ((((_2781_v56).contains(_2772_v47)) ? ((_2781_v56).get(_2772_v47)) : (_2770_cf14))).minus((((_2782_v57).contains(_2777_v52)) ? ((_2782_v57).get(_2777_v52)) : (new BigNumber(318))));
          let _rhs365 = p0;
          let _rhs366 = _2775_v50;
          let _rhs367 = _2778_v53;
          let _rhs368 = _2779_v54;
          let _lhs311 = _2775_v50;
          let _lhs312 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_2775_v50).length));
          let _lhs313 = globalState;
          let _lhs314 = globalState;
          _lhs311[_lhs312] = _rhs364;
          _lhs313.f1 = _rhs365;
          _lhs314.f9 = _rhs366;
          _2778_v53 = _rhs367;
          _2779_v54 = _rhs368;
        } else if (_source25.is_DC12) {
          let _2783___mcc_h1 = (_source25).cf15;
          let _2784___mcc_h2 = (_source25).cf16;
          let _2785___mcc_h3 = (_source25).cf17;
          let _2786_cf17 = _2785___mcc_h3;
          let _2787_cf16 = _2784___mcc_h2;
          let _2788_cf15 = _2783___mcc_h1;
          let _2789_v58;
          _2789_v58 = _dafny.Seq.of(_2786_cf17, _2764_v44, _2786_cf17);
          _2789_v58 = _module.__default.fm21(globalState);
          let _2790_v59;
          let _nw516 = Array((new BigNumber(3)).toNumber()).fill(_module.D13.Default());
          _2790_v59 = _nw516;
          let _2791_v60;
          _2791_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2790_v59);
          _2788_cf15 = new BigNumber((_2791_v60).length);
          let _2792_v61;
          _2792_v61 = _dafny.MultiSet.fromElements(true, _2787_cf16, _2787_cf16);
          let _2793_v62;
          _2793_v62 = _dafny.Seq.of(_2787_cf16);
          let _2794_v63;
          _2794_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2792_v61,_2793_v62);
          _2794_v63 = _2794_v63;
          let _2795_v64;
          _2795_v64 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_2787_cf16)).length));
          let _2796_v65;
          _2796_v65 = _dafny.Map.Empty.slice().updateUnsafe(_2795_v64,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2788_cf15,_2787_cf16)).length));
          let _2797_v66;
          _2797_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-974),p0);
          let _2798_v67;
          _2798_v67 = _dafny.Seq.of(_2797_v66);
          let _2799_v68;
          _2799_v68 = _dafny.Map.Empty.slice().updateUnsafe((_2798_v67)[_module.__default.safeIndex(_2788_cf15, new BigNumber((_2798_v67).length))],new BigNumber(425));
          let _2800_v69;
          let _nw517 = new _module.C3();
          _nw517.__ctor(p0, new BigNumber((_2796_v65).length), _2799_v68);
          _2800_v69 = _nw517;
          let _2801_v70;
          _2801_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2800_v69);
          (globalState).f1 = ((_module.__default.fm48(_dafny.Map.Empty.slice().updateUnsafe(_2764_v44,_2764_v44), globalState)).dtor_cf7).isEqualTo(new BigNumber((_2801_v70).length));
        } else if (_source25.is_DC10) {
          let _2802___mcc_h4 = (_source25).cf13;
          let _2803_cf13 = _2802___mcc_h4;
          let _2804_v71;
          let _nw518 = Array((new BigNumber(13)).toNumber()).fill(_module.D0.Default());
          _2804_v71 = _nw518;
          let _2805_v72;
          _2805_v72 = _2804_v71;
          let _2806_v73;
          _2806_v73 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p0, p0, (_this).f15, globalState),(_this).f15);
          let _2807_v74;
          _2807_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2805_v72,_2806_v73);
          let _2808_v75;
          _2808_v75 = (((_2807_v74).contains(_2805_v72)) ? ((_2807_v74).get(_2805_v72)) : (_2806_v73));
          _2808_v75 = _2808_v75;
          let _2809_v76;
          let _nw519 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _2809_v76 = _nw519;
          let _index459 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_2809_v76).length));
          (_2809_v76)[_index459] = (_this).f15;
          let _index460 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_2809_v76).length));
          (_2809_v76)[_index460] = (_dafny.ZERO).minus((_this).f15);
          (globalState).f1 = p0;
          let _2810_v77;
          _2810_v77 = _dafny.Seq.UnicodeFromString("do");
          let _index461 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_2809_v76).length));
          (_2809_v76)[_index461] = new BigNumber((((p0) ? (_2810_v77) : (_dafny.Seq.UnicodeFromString("ecflprnxd")))).length);
        } else {
          let _2811___mcc_h5 = (_source25).cf18;
          let _2812_cf18 = _2811___mcc_h5;
          let _2813_v78;
          _2813_v78 = _dafny.Set.fromElements((_this).f15, (_this).f15);
          let _2814_v79;
          _2814_v79 = _module.D0.create_DC0(p0);
          let _2815_v80;
          _2815_v80 = _module.D0.create_DC2(_2814_v79);
          let _2816_v81;
          _2816_v81 = _dafny.Seq.UnicodeFromString("vydknp");
          let _2817_v82;
          _2817_v82 = _module.D16.create_DC45(new BigNumber((_2765_v45).cardinality()), _2764_v44, new BigNumber((_dafny.Seq.of(p0)).length), true);
          let _pat_let_tv66 = _2814_v79;
          let _pat_let_tv67 = _2814_v79;
          let _pat_let_tv68 = _2814_v79;
          let _pat_let_tv69 = _2814_v79;
          let _pat_let_tv70 = p0;
          let _2818_v83;
          let _nw520 = Array((new BigNumber(27)).toNumber());
          _nw520[(_dafny.ZERO).toNumber()] = _2815_v80;
          _nw520[(_dafny.ONE).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(2)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(3)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(4)).toNumber()] = function (_pat_let92_0) {
            return function (_2821_dt__update__tmp_h2) {
              return function (_pat_let95_0) {
                return function (_2822_dt__update_hcf2_h1) {
                  return _module.D0.create_DC2(_2822_dt__update_hcf2_h1);
                }(_pat_let95_0);
              }(_pat_let_tv67);
            }(_pat_let92_0);
          }(function (_pat_let93_0) {
            return function (_2819_dt__update__tmp_h1) {
              return function (_pat_let94_0) {
                return function (_2820_dt__update_hcf2_h0) {
                  return _module.D0.create_DC2(_2820_dt__update_hcf2_h0);
                }(_pat_let94_0);
              }(_pat_let_tv66);
            }(_pat_let93_0);
          }(_2815_v80));
          _nw520[(new BigNumber(5)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(6)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(7)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(8)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(9)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(10)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(11)).toNumber()] = function (_pat_let96_0) {
            return function (_2823_dt__update__tmp_h3) {
              return function (_pat_let97_0) {
                return function (_2824_dt__update_hcf2_h2) {
                  return _module.D0.create_DC2(_2824_dt__update_hcf2_h2);
                }(_pat_let97_0);
              }(_pat_let_tv68);
            }(_pat_let96_0);
          }(_2815_v80);
          _nw520[(new BigNumber(12)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(13)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(14)).toNumber()] = function (_pat_let98_0) {
            return function (_2825_dt__update__tmp_h4) {
              return function (_pat_let99_0) {
                return function (_2826_dt__update_hcf2_h3) {
                  return _module.D0.create_DC2(_2826_dt__update_hcf2_h3);
                }(_pat_let99_0);
              }(_pat_let_tv69);
            }(_pat_let98_0);
          }(_module.D0.create_DC2(_2814_v79));
          _nw520[(new BigNumber(15)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(16)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(17)).toNumber()] = _module.D0.create_DC2(_2814_v79);
          _nw520[(new BigNumber(18)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(19)).toNumber()] = _module.D0.create_DC2(_2814_v79);
          _nw520[(new BigNumber(20)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(21)).toNumber()] = function (_pat_let100_0) {
            return function (_2827_dt__update__tmp_h5) {
              return function (_pat_let101_0) {
                return function (_2828_dt__update_hcf2_h4) {
                  return _module.D0.create_DC2(_2828_dt__update_hcf2_h4);
                }(_pat_let101_0);
              }(_module.D0.create_DC0(_pat_let_tv70));
            }(_pat_let100_0);
          }(_2815_v80);
          _nw520[(new BigNumber(22)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(23)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(24)).toNumber()] = _2815_v80;
          _nw520[(new BigNumber(25)).toNumber()] = (_this).fm3(p0, new BigNumber((_2816_v81).length), p0, (_2817_v82).dtor_cf81, globalState);
          _nw520[(new BigNumber(26)).toNumber()] = _2815_v80;
          _2818_v83 = _nw520;
          let _2829_v84;
          _2829_v84 = _dafny.Seq.of(_2818_v83, _2818_v83);
          let _2830_v85;
          let _init78 = function (_2831_i3) {
            return (_2831_i3).multipliedBy((_this).f15);
          };
          let _nw521 = Array((new BigNumber(24)).toNumber());
          for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw521.length); _i0_78++) {
            _nw521[_i0_78] = _init78(new BigNumber(_i0_78));
          }
          _2830_v85 = _nw521;
          _module.__default.m0(((!(false)) ? (_2813_v78) : (_2813_v78)), (_2829_v84)[_module.__default.safeIndex((_this).f15, new BigNumber((_2829_v84).length))], _2830_v85, _2816_v81, globalState);
          let _2832_v86;
          _2832_v86 = _module.D3.create_DC6(_2830_v85);
          let _2833_v87;
          let _nw522 = Array((new BigNumber(20)).toNumber());
          _nw522[(_dafny.ZERO).toNumber()] = _module.D3.create_DC6(_2830_v85);
          _nw522[(_dafny.ONE).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(2)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(3)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(4)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(5)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(6)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(7)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(8)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(9)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(10)).toNumber()] = _module.D3.create_DC6(_2830_v85);
          _nw522[(new BigNumber(11)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(12)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(13)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(14)).toNumber()] = _module.D3.create_DC6(_2830_v85);
          _nw522[(new BigNumber(15)).toNumber()] = _module.D3.create_DC6(_2830_v85);
          _nw522[(new BigNumber(16)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(17)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(18)).toNumber()] = _2832_v86;
          _nw522[(new BigNumber(19)).toNumber()] = _module.D3.create_DC6(_2830_v85);
          _2833_v87 = _nw522;
          let _2834_v88;
          _2834_v88 = _dafny.Set.fromElements(_2833_v87);
          (globalState).f1 = _dafny.Seq.IsPrefixOf(_2816_v81, (_module.D9.create_DC23(_2834_v88, _dafny.Seq.UnicodeFromString("luiufanv"), p0, _2830_v85, _2764_v44)).dtor_cf38);
          let _index462 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_2830_v85).length));
          (_2830_v85)[_index462] = (_this).f15;
          let _index463 = _module.__default.safeIndex(new BigNumber(660), new BigNumber((_2830_v85).length));
          (_2830_v85)[_index463] = (_this).f15;
          let _2835_v89;
          let _nw523 = new _module.C6();
          _nw523.__ctor();
          _2835_v89 = _nw523;
        }
        let _2836_v90;
        _2836_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(145),(_this).f15);
        let _2837_v91;
        _2837_v91 = _dafny.Set.fromElements(p0);
        let _2838_v92;
        _2838_v92 = _dafny.Map.Empty.slice().updateUnsafe((_2836_v90).Merge(_2836_v90),_2837_v91);
        let _2839_v93;
        _2839_v93 = _dafny.Seq.of(((_this).f15).multipliedBy((_this).f15), ((_this).f15).plus(new BigNumber(576)), (_this).f15, (_this).f15, (_this).f15);
        let _2840_v94;
        _2840_v94 = _dafny.Seq.of(p0);
        let _2841_v95;
        _2841_v95 = _module.D11.create_DC31(_2840_v94, (_this).f15, _module.__default.fm1(p0, globalState), new BigNumber((_2836_v90).length), (_this).f15);
        let _2842_v96;
        _2842_v96 = _dafny.Map.Empty.slice().updateUnsafe((_module.D3.create_DC8(_2839_v93, _dafny.Seq.UnicodeFromString("nsnffddq"), p0, new _dafny.CodePoint('o'.codePointAt(0)))).dtor_cf10,_2764_v44);
        let _2843_v97;
        let _nw524 = Array((new BigNumber(9)).toNumber());
        _nw524[(_dafny.ZERO).toNumber()] = (_this).f15;
        _nw524[(_dafny.ONE).toNumber()] = (_this).f15;
        _nw524[(new BigNumber(2)).toNumber()] = (_this).f15;
        _nw524[(new BigNumber(3)).toNumber()] = (_2841_v95).dtor_cf56;
        _nw524[(new BigNumber(4)).toNumber()] = (_this).f15;
        _nw524[(new BigNumber(5)).toNumber()] = (_this).f15;
        _nw524[(new BigNumber(6)).toNumber()] = (_this).f15;
        _nw524[(new BigNumber(7)).toNumber()] = (_this).f15;
        _nw524[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Set.fromElements(new BigNumber((_2842_v96).length), (_this).f15)).length);
        _2843_v97 = _nw524;
        let _2844_v98;
        _2844_v98 = _dafny.Set.fromElements(_2843_v97);
        let _2845_v99;
        _2845_v99 = _dafny.Map.Empty.slice().updateUnsafe(_2844_v98,_2839_v93);
        let _rhs369 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_dafny.ZERO).minus((_this).f15)),_module.__default.fm31((_this).f15, _module.__default.fm2(globalState), globalState));
        let _rhs370 = (((_2845_v99).contains(_2844_v98)) ? ((_2845_v99).get(_2844_v98)) : (_2839_v93));
        let _rhs371 = _2715_v0;
        let _lhs315 = globalState;
        _2838_v92 = _rhs369;
        _2839_v93 = _rhs370;
        _lhs315.f14 = _rhs371;
        let _index464 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2715_v0).length));
        (_2715_v0)[_index464] = !(p0);
        let _index465 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2715_v0).length));
        (_2715_v0)[_index465] = _module.__default.fm2(globalState);
        let _2846_v100;
        let _nw525 = new _module.C7();
        _nw525.__ctor();
        _2846_v100 = _nw525;
      }
      (globalState).f6 = new BigNumber(309);
      let _2847_i4;
      _2847_i4 = _dafny.ZERO;
      L20: {
        while (p0) {
          C20: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2847_i4)) {
              break L20;
            }
            _2847_i4 = (_2847_i4).plus(_dafny.ONE);
            let _2848_v101;
            let _nw526 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
            _2848_v101 = _nw526;
            let _index466 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_2848_v101).length));
            (_2848_v101)[_index466] = _module.__default.fm49(!(p0), globalState);
            let _2849_v102;
            _2849_v102 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f15);
            let _2850_v103;
            _2850_v103 = _2849_v102;
            let _index467 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_2848_v101).length));
            (_2848_v101)[_index467] = _2849_v102;
            let _2851_v104;
            _2851_v104 = new _dafny.CodePoint('a'.codePointAt(0));
            let _2852_v105;
            _2852_v105 = _dafny.MultiSet.fromElements(_2851_v104);
            let _2853_v106;
            _2853_v106 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2852_v105).cardinality()),(_dafny.ZERO).minus((_this).f15));
            let _2854_v107;
            _2854_v107 = _dafny.Seq.of(false);
            let _2855_v108;
            _2855_v108 = _dafny.Map.Empty.slice().updateUnsafe(_2715_v0,_2853_v106);
            let _2856_v109;
            _2856_v109 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15));
            let _2857_v111;
            let _nw527 = Array((new BigNumber(14)).toNumber());
            _nw527[(_dafny.ZERO).toNumber()] = _2853_v106;
            _nw527[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(863),(_this).f15);
            _nw527[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2854_v107).length),(_this).f15);
            _nw527[(new BigNumber(3)).toNumber()] = (_2853_v106).Merge(_2853_v106);
            _nw527[(new BigNumber(4)).toNumber()] = (_2853_v106).update((_this).f15, (_this).f15);
            _nw527[(new BigNumber(5)).toNumber()] = _2853_v106;
            _nw527[(new BigNumber(6)).toNumber()] = _2853_v106;
            _nw527[(new BigNumber(7)).toNumber()] = _2853_v106;
            _nw527[(new BigNumber(8)).toNumber()] = _2853_v106;
            _nw527[(new BigNumber(9)).toNumber()] = (_2853_v106).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber(313)));
            _nw527[(new BigNumber(10)).toNumber()] = ((((_2855_v108).contains(_2715_v0)) ? ((_2855_v108).get(_2715_v0)) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15)))).Merge((((_2856_v109).contains(_module.__default.fm2(globalState))) ? ((_2856_v109).get(_module.__default.fm2(globalState))) : (_2853_v106)));
            _nw527[(new BigNumber(11)).toNumber()] = function () {
              let _coll89 = new _dafny.Map();
              for (const _compr_89 of _dafny.IntegerRange(new BigNumber(16), new BigNumber(622))) {
                let _2858_v110 = _compr_89;
                if (((new BigNumber(16)).isLessThanOrEqualTo(_2858_v110)) && ((_2858_v110).isLessThan(new BigNumber(622)))) {
                  _coll89.push([(_2858_v110).minus((_this).f15),(_this).f15]);
                }
              }
              return _coll89;
            }();
            _nw527[(new BigNumber(12)).toNumber()] = ((_2853_v106).update((_this).f15, (_this).f15)).Merge(_2853_v106);
            _nw527[(new BigNumber(13)).toNumber()] = _2853_v106;
            _2857_v111 = _nw527;
            let _index468 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_2857_v111).length));
            (_2857_v111)[_index468] = _2853_v106;
            let _index469 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_2857_v111).length));
            (_2857_v111)[_index469] = _2853_v106;
            (globalState).f6 = (_this).f15;
            (globalState).f1 = _module.__default.fm2(globalState);
          }
        }
      }
      (globalState).f1 = ((p0) ? (p0) : (p0));
      return;
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _2859_v0;
      _2859_v0 = _dafny.Seq.of((_this).f15);
      let _2860_v1;
      _2860_v1 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((p0).length)).multipliedBy(new BigNumber((_2859_v0).length)),_module.__default.safeModuloInt((_this).f15, (_this).f15));
      let _2861_v2;
      _2861_v2 = _dafny.Seq.UnicodeFromString("d");
      _2860_v1 = (_2860_v1).update((_this).f15, new BigNumber((_2861_v2).length));
      let _2862_v3;
      _2862_v3 = true;
      if (_2862_v3) {
        _2862_v3 = ((_this).fm5(_2862_v3, globalState)).isLessThan((_this).f15);
        _2859_v0 = _2859_v0;
        let _2863_v4;
        let _init79 = function (_2864_i0) {
          return _module.__default.safeModuloInt(_2864_i0, (_this).f15);
        };
        let _nw528 = Array((new BigNumber(7)).toNumber());
        for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw528.length); _i0_79++) {
          _nw528[_i0_79] = _init79(new BigNumber(_i0_79));
        }
        _2863_v4 = _nw528;
        let _2865_v5;
        _2865_v5 = _module.D3.create_DC6(_2863_v4);
        let _2866_v6;
        _2866_v6 = _dafny.Seq.of(_2865_v5, _2865_v5, _2865_v5);
        let _2867_v7;
        _2867_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,_2866_v6);
        let _2868_v8;
        _2868_v8 = _dafny.MultiSet.fromElements(_2863_v4, _2863_v4);
        let _2869_v9;
        _2869_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2868_v8,!(_module.__default.fm2(globalState)));
        let _2870_v10;
        let _nw529 = new _module.C0();
        _nw529.__ctor((((_2867_v7).contains(_2862_v3)) ? ((_2867_v7).get(_2862_v3)) : (_2866_v6)), (((_2869_v9).contains(_2868_v8)) ? ((_2869_v9).get(_2868_v8)) : (_2862_v3)));
        _2870_v10 = _nw529;
        let _2871_v11;
        _2871_v11 = _module.D11.create_DC32(_2862_v3, (_this).f15, _2862_v3);
        let _2872_v12;
        _2872_v12 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f15);
        let _2873_v13;
        let _nw530 = Array((new BigNumber(27)).toNumber());
        _nw530[(_dafny.ZERO).toNumber()] = (_2870_v10).f23;
        _nw530[(_dafny.ONE).toNumber()] = (_2870_v10).f23;
        _nw530[(new BigNumber(2)).toNumber()] = !(((_2870_v10).f23) || (_2862_v3));
        _nw530[(new BigNumber(3)).toNumber()] = (_2862_v3) || (!(false));
        _nw530[(new BigNumber(4)).toNumber()] = _2862_v3;
        _nw530[(new BigNumber(5)).toNumber()] = (_2870_v10).f23;
        _nw530[(new BigNumber(6)).toNumber()] = !((((_2870_v10).f23) ? (_2862_v3) : (_2862_v3)));
        _nw530[(new BigNumber(7)).toNumber()] = !((_2870_v10).f23);
        _nw530[(new BigNumber(8)).toNumber()] = (_2862_v3) || (_2862_v3);
        _nw530[(new BigNumber(9)).toNumber()] = true;
        _nw530[(new BigNumber(10)).toNumber()] = !((_2870_v10).f23);
        _nw530[(new BigNumber(11)).toNumber()] = !((_dafny.MultiSet.FromArray(_2859_v0)).update(new BigNumber((p0).length), _module.__default.abs((((_2860_v1).contains((_this).f15)) ? ((_2860_v1).get((_this).f15)) : ((_this).f15))))).equals(_dafny.MultiSet.fromElements((_this).f15));
        _nw530[(new BigNumber(12)).toNumber()] = _2862_v3;
        _nw530[(new BigNumber(13)).toNumber()] = !(_2862_v3) || (_2862_v3);
        _nw530[(new BigNumber(14)).toNumber()] = (_2870_v10).fm17(new BigNumber(-577), new BigNumber(982), (_2870_v10).f23, globalState);
        _nw530[(new BigNumber(15)).toNumber()] = !((_this).f15).isEqualTo((_this).f15);
        _nw530[(new BigNumber(16)).toNumber()] = _2862_v3;
        _nw530[(new BigNumber(17)).toNumber()] = !(_dafny.Set.fromElements((_this).f15)).contains((_2871_v11).dtor_cf60);
        _nw530[(new BigNumber(18)).toNumber()] = (((_2870_v10).f23) ? ((_2870_v10).f23) : (_2862_v3));
        _nw530[(new BigNumber(19)).toNumber()] = (new BigNumber((p0).length)).isLessThan((_this).f15);
        _nw530[(new BigNumber(20)).toNumber()] = _module.__default.fm2(globalState);
        _nw530[(new BigNumber(21)).toNumber()] = (_this).fm4(new BigNumber((_dafny.Seq.of((_this).f15, new BigNumber((_2872_v12).length))).length), globalState);
        _nw530[(new BigNumber(22)).toNumber()] = ((_this).f15).isLessThanOrEqualTo(new BigNumber((_2861_v2).length));
        _nw530[(new BigNumber(23)).toNumber()] = !(((_this).f15).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2859_v0).length),(_2870_v10).f23)).length)));
        _nw530[(new BigNumber(24)).toNumber()] = (_2862_v3) && (_2862_v3);
        _nw530[(new BigNumber(25)).toNumber()] = (_2870_v10).f23;
        _nw530[(new BigNumber(26)).toNumber()] = false;
        _2873_v13 = _nw530;
        let _index470 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_2873_v13).length));
        (_2873_v13)[_index470] = !((_2870_v10).f23) || (_2862_v3);
        let _2874_v14;
        _2874_v14 = _module.D20.create_DC50(_2862_v3, (_this).f15, (_this).f15, _2862_v3, (_2870_v10).f23);
        let _index471 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_2873_v13).length));
        (_2873_v13)[_index471] = (_2874_v14).dtor_cf90;
        let _2875_v15;
        let _nw531 = Array((new BigNumber(10)).toNumber());
        _nw531[(_dafny.ZERO).toNumber()] = _2861_v2;
        _nw531[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_2861_v2, _module.__default.fm21(globalState));
        _nw531[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(255)), function (_2876_i1) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        });
        _nw531[(new BigNumber(3)).toNumber()] = _module.__default.fm21(globalState);
        _nw531[(new BigNumber(4)).toNumber()] = (((_2870_v10).f23) ? (_2861_v2) : (_dafny.Seq.UnicodeFromString("togmll")));
        _nw531[(new BigNumber(5)).toNumber()] = _2861_v2;
        _nw531[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("xvnvbtp");
        _nw531[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_2861_v2, _2861_v2);
        _nw531[(new BigNumber(8)).toNumber()] = _2861_v2;
        _nw531[(new BigNumber(9)).toNumber()] = _2861_v2;
        _2875_v15 = _nw531;
        let _2877_v16;
        _2877_v16 = new _dafny.CodePoint('k'.codePointAt(0));
        let _2878_v17;
        _2878_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,true);
        let _2879_v18;
        _2879_v18 = _dafny.Set.fromElements(_2862_v3, false);
        let _2880_v19;
        _2880_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2873_v13);
        let _2881_v20;
        _2881_v20 = _dafny.Set.fromElements((((_2880_v19).contains((_this).f15)) ? ((_2880_v19).get((_this).f15)) : (_2873_v13)), _2873_v13, _2873_v13, _2873_v13);
        let _2882_v21;
        _2882_v21 = _dafny.Map.Empty.slice().updateUnsafe((_2870_v10).f23,(_2870_v10).f23);
        let _2883_v22;
        _2883_v22 = _module.D14.create_DC40((_2873_v13)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_2873_v13).length))]);
        let _2884_v23;
        _2884_v23 = _dafny.MultiSet.fromElements((_2870_v10).f23, (_2873_v13)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((_2873_v13).length))], _2862_v3, (((_2882_v21).contains(_2862_v3)) ? ((_2882_v21).get(_2862_v3)) : ((_2883_v22).dtor_cf71)), (_2870_v10).f23);
        let _rhs372 = _2875_v15;
        let _rhs373 = (_this).f15;
        let _rhs374 = (((new BigNumber((_2879_v18).length)).isLessThanOrEqualTo(new BigNumber((_2878_v17).length))) ? (_module.__default.safeModuloInt(new BigNumber(256), new BigNumber((_dafny.Seq.update(_2861_v2, _module.__default.safeIndex((_this).f15, new BigNumber((_2861_v2).length)), _2877_v16)).length))) : (_module.__default.safeModuloInt((_this).f15, new BigNumber((_2881_v20).length))));
        let _rhs375 = (((_2884_v23).contains(!_dafny.areEqual(_2861_v2, _dafny.Seq.UnicodeFromString("pbpuyrmd")))) ? ((_2884_v23).get(!_dafny.areEqual(_2861_v2, _dafny.Seq.UnicodeFromString("pbpuyrmd")))) : ((_this).f15));
        let _rhs376 = _2877_v16;
        let _lhs316 = globalState;
        let _lhs317 = globalState;
        let _lhs318 = globalState;
        _2875_v15 = _rhs372;
        _lhs316.f6 = _rhs373;
        _lhs317.f6 = _rhs374;
        _lhs318.f6 = _rhs375;
        _2877_v16 = _rhs376;
      } else {
        (globalState).f1 = _2862_v3;
        r1 = !(_2862_v3);
        let _2885_v24;
        let _nw532 = Array((new BigNumber(29)).toNumber()).fill(_module.D16.Default());
        _2885_v24 = _nw532;
        let _index472 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_2885_v24).length));
        (_2885_v24)[_index472] = _module.D16.create_DC44(_dafny.MultiSet.fromElements(false, _2862_v3, _2862_v3));
        let _2886_v25;
        _2886_v25 = _dafny.MultiSet.fromElements(_2862_v3, _2862_v3);
        let _2887_v26;
        _2887_v26 = _module.D16.create_DC44(_2886_v25);
        let _index473 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_2885_v24).length));
        let _rhs377 = ((_this).f15).multipliedBy(((_dafny.ZERO).minus(_module.__default.fm1(_2862_v3, globalState))).minus(new BigNumber((_2859_v0).length)));
        let _rhs378 = (_this).f15;
        let _rhs379 = _2887_v26;
        let _rhs380 = (_dafny.ZERO).minus((_this).f15);
        let _rhs381 = ((_2862_v3) ? (_2861_v2) : (_dafny.Seq.UnicodeFromString("orpgjwi")));
        let _lhs319 = globalState;
        let _lhs320 = _2885_v24;
        let _lhs321 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_2885_v24).length));
        r2 = _rhs377;
        _lhs319.f6 = _rhs378;
        _lhs320[_lhs321] = _rhs379;
        r2 = _rhs380;
        _2861_v2 = _rhs381;
        if ((((_this).f15).minus((_this).f15)).isLessThan((_this).f15)) {
          let _2888_v27;
          let _nw533 = new _module.C6();
          _nw533.__ctor();
          _2888_v27 = _nw533;
          let _2889_v28;
          let _nw534 = Array((new BigNumber(26)).toNumber()).fill(false);
          _2889_v28 = _nw534;
          let _2890_v29;
          _2890_v29 = _dafny.Seq.of(_2889_v28, _2889_v28, _2889_v28);
          let _2891_v30;
          _2891_v30 = _dafny.Set.fromElements(new BigNumber(362), new BigNumber(187), new BigNumber((_2861_v2).length), (_this).f15, (_this).f15);
          let _rhs382 = new BigNumber(((_2891_v30).Union(_2891_v30)).length);
          let _rhs383 = (_this).f15;
          let _rhs384 = (_this).f15;
          let _rhs385 = _dafny.Seq.Concat(_2890_v29, _2890_v29);
          let _lhs322 = globalState;
          let _lhs323 = globalState;
          let _lhs324 = globalState;
          _lhs322.f6 = _rhs382;
          _lhs323.f6 = _rhs383;
          _lhs324.f6 = _rhs384;
          _2890_v29 = _rhs385;
          let _2892_v31;
          _2892_v31 = _dafny.MultiSet.fromElements((_this).f15, new BigNumber(486), new BigNumber((_2861_v2).length));
          let _2893_v32;
          _2893_v32 = new _dafny.CodePoint('j'.codePointAt(0));
          let _2894_v33;
          let _nw535 = Array((new BigNumber(6)).toNumber());
          _nw535[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_this).f15, (_this).f15);
          _nw535[(_dafny.ONE).toNumber()] = ((!(!(_2862_v3))) ? (new BigNumber((_2892_v31).cardinality())) : ((_this).f15));
          _nw535[(new BigNumber(2)).toNumber()] = ((_this).f15).plus((_2888_v27).fm9(_dafny.Seq.of(_2893_v32, _2893_v32), globalState));
          _nw535[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_2859_v0)[_module.__default.safeIndex((_this).f15, new BigNumber((_2859_v0).length))]);
          _nw535[(new BigNumber(4)).toNumber()] = (_this).f15;
          _nw535[(new BigNumber(5)).toNumber()] = new BigNumber((_2886_v25).cardinality());
          _2894_v33 = _nw535;
          let _index474 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_2894_v33).length));
          (_2894_v33)[_index474] = new BigNumber((_2861_v2).length);
          let _index475 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_2894_v33).length));
          (_2894_v33)[_index475] = new BigNumber(133);
          _2893_v32 = _2893_v32;
          let _2895_v34;
          _2895_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,new _dafny.CodePoint('u'.codePointAt(0)));
          let _2896_v35;
          _2896_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,_2862_v3);
          let _index476 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_2894_v33).length));
          let _rhs386 = _2895_v34;
          let _rhs387 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rsimrwma")).length))).multipliedBy((_2894_v33)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_2894_v33).length))]);
          let _rhs388 = (((_2896_v35).contains(!(!(_2862_v3)))) ? ((_2896_v35).get(!(!(_2862_v3)))) : (_2862_v3));
          let _rhs389 = _2861_v2;
          let _lhs325 = _2894_v33;
          let _lhs326 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_2894_v33).length));
          let _lhs327 = globalState;
          _2895_v34 = _rhs386;
          _lhs325[_lhs326] = _rhs387;
          _lhs327.f1 = _rhs388;
          _2861_v2 = _rhs389;
        } else {
          (globalState).f1 = !(_2886_v25).contains(_2862_v3);
          let _2897_v36;
          let _nw536 = Array((new BigNumber(3)).toNumber()).fill(false);
          _2897_v36 = _nw536;
          let _index477 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_2897_v36).length));
          (_2897_v36)[_index477] = _2862_v3;
          let _index478 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_2897_v36).length));
          (_2897_v36)[_index478] = _2862_v3;
          let _2898_v37;
          let _nw537 = new _module.C4();
          _nw537.__ctor();
          _2898_v37 = _nw537;
          let _2899_v38;
          let _nw538 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _2899_v38 = _nw538;
          let _index479 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_2899_v38).length));
          (_2899_v38)[_index479] = (_this).f15;
          let _2900_v39;
          let _nw539 = new _module.C5();
          _nw539.__ctor((_this).f15);
          _2900_v39 = _nw539;
          let _2901_v40;
          _2901_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2886_v25,_2900_v39);
          let _index480 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_2899_v38).length));
          (_2899_v38)[_index480] = new BigNumber(((_2901_v40).Merge(_2901_v40)).length);
          (globalState).f6 = (_dafny.ZERO).minus((new BigNumber(690)).plus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-429)), ((_2902_v36, _2903_v3) => function (_2904_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe((_2902_v36)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_2902_v36).length))],_2903_v3);
          })(_2897_v36, _2862_v3)), _module.__default.fm50((_2900_v39).f16, new BigNumber(-868), globalState))).length)));
        }
        let _2905_v41;
        let _init80 = ((_2906_v3) => function (_2907_i3) {
          return _2906_v3;
        })(_2862_v3);
        let _nw540 = Array((new BigNumber(9)).toNumber());
        for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw540.length); _i0_80++) {
          _nw540[_i0_80] = _init80(new BigNumber(_i0_80));
        }
        _2905_v41 = _nw540;
        let _index481 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2905_v41).length));
        (_2905_v41)[_index481] = _2862_v3;
        let _index482 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2905_v41).length));
        let _rhs390 = !(_2862_v3);
        let _rhs391 = !(_2862_v3);
        let _lhs328 = _2905_v41;
        let _lhs329 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2905_v41).length));
        r1 = _rhs390;
        _lhs328[_lhs329] = _rhs391;
      }
      let _hi17 = (_dafny.ZERO).minus((_this).f15);
      for (let _2908_i4 = (_this).f15; _2908_i4.isLessThan(_hi17); _2908_i4 = _2908_i4.plus(_dafny.ONE)) {
        let _2909_v42;
        let _init81 = function (_2910_i5) {
          return (_2910_i5).minus((_dafny.ZERO).minus((_this).f15));
        };
        let _nw541 = Array((new BigNumber(18)).toNumber());
        for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw541.length); _i0_81++) {
          _nw541[_i0_81] = _init81(new BigNumber(_i0_81));
        }
        _2909_v42 = _nw541;
        let _2911_v43;
        _2911_v43 = new _dafny.CodePoint('j'.codePointAt(0));
        let _index483 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length));
        (_2909_v42)[_index483] = _module.__default.safeDivisionInt((_this).f15, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_2861_v2, _module.__default.safeIndex(_2908_i4, new BigNumber((_2861_v2).length)), _2911_v43)).length)));
        let _2912_v44;
        _2912_v44 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,_2908_i4);
        let _index484 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length));
        (_2909_v42)[_index484] = (new BigNumber((_2912_v44).length)).minus(new BigNumber((_2861_v2).length));
        let _rhs392 = _dafny.Seq.update(_dafny.Seq.Concat(((_2862_v3) ? (_2861_v2) : (_2861_v2)), _2861_v2), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_2861_v2, _2861_v2)).length), new BigNumber((_dafny.Seq.Concat(((_2862_v3) ? (_2861_v2) : (_2861_v2)), _2861_v2)).length)), _2911_v43);
        let _rhs393 = _2862_v3;
        let _rhs394 = _2861_v2;
        let _rhs395 = _2862_v3;
        let _rhs396 = (_this).f15;
        let _lhs330 = globalState;
        _2861_v2 = _rhs392;
        _lhs330.f1 = _rhs393;
        _2861_v2 = _rhs394;
        r1 = _rhs395;
        r2 = _rhs396;
        let _2913_v45;
        _2913_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,_2862_v3);
        let _2914_v46;
        _2914_v46 = _module.D1.create_DC3(_2913_v45);
        let _source26 = _2914_v46;
        if (_source26.is_DC4) {
          let _2915___mcc_h0 = (_source26).cf4;
          let _2916_cf4 = _2915___mcc_h0;
          let _2917_v47;
          _2917_v47 = _dafny.Map.Empty.slice().updateUnsafe(false,_2861_v2);
          _2861_v2 = ((_2862_v3) ? (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nhk"), _dafny.Seq.UnicodeFromString("fiwsofh"))) : (_dafny.Seq.update((((_2917_v47).contains(_2862_v3)) ? ((_2917_v47).get(_2862_v3)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(720)), ((_2918_cf4) => function (_2919_i6) {
            return _2918_cf4;
          })(_2916_cf4)))), _module.__default.safeIndex((_this).f15, new BigNumber(((((_2917_v47).contains(_2862_v3)) ? ((_2917_v47).get(_2862_v3)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(720)), ((_2920_cf4) => function (_2921_i6) {
            return _2920_cf4;
          })(_2916_cf4))))).length)), _2916_cf4)));
          let _2922_v48;
          let _nw542 = new _module.C4();
          _nw542.__ctor();
          _2922_v48 = _nw542;
          let _2923_v49;
          _2923_v49 = _dafny.Seq.of(_2909_v42);
          (globalState).f6 = (_dafny.ZERO).minus((((_2909_v42)[_module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length))]).multipliedBy(new BigNumber((_2923_v49).length))).multipliedBy(new BigNumber(651)));
          let _2924_v50;
          let _nw543 = Array((new BigNumber(12)).toNumber());
          _2924_v50 = _nw543;
          let _nw544 = Array((new BigNumber(22)).toNumber());
          _2924_v50 = _nw544;
        } else {
          let _2925___mcc_h1 = (_source26).cf3;
          let _2926_cf3 = _2925___mcc_h1;
          let _2927_v51;
          _2927_v51 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(661), _2908_i4)).length));
          (globalState).f6 = ((_2862_v3) ? (new BigNumber((_2859_v0).length)) : (new BigNumber((_2927_v51).length)));
          let _2928_v52;
          _2928_v52 = _dafny.Map.Empty.slice().updateUnsafe(_2909_v42,_2861_v2);
          let _2929_v53;
          let _init82 = ((_2930_v3) => function (_2931_i7) {
            return !(_2930_v3);
          })(_2862_v3);
          let _nw545 = Array((new BigNumber(7)).toNumber());
          for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw545.length); _i0_82++) {
            _nw545[_i0_82] = _init82(new BigNumber(_i0_82));
          }
          _2929_v53 = _nw545;
          let _index485 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_2929_v53).length));
          (_2929_v53)[_index485] = _2862_v3;
          let _2932_v54;
          let _nw546 = Array((new BigNumber(5)).toNumber());
          _nw546[(_dafny.ZERO).toNumber()] = _2929_v53;
          _nw546[(_dafny.ONE).toNumber()] = _2929_v53;
          _nw546[(new BigNumber(2)).toNumber()] = _2929_v53;
          _nw546[(new BigNumber(3)).toNumber()] = _2929_v53;
          _nw546[(new BigNumber(4)).toNumber()] = _2929_v53;
          _2932_v54 = _nw546;
          let _2933_v55;
          _2933_v55 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_2929_v53);
          let _index486 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_2932_v54).length));
          (_2932_v54)[_index486] = (((_2933_v55).contains(_2862_v3)) ? ((_2933_v55).get(_2862_v3)) : (_2929_v53));
          let _index487 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_2929_v53).length));
          let _index488 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_2932_v54).length));
          let _index489 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length));
          let _rhs397 = _2928_v52;
          let _rhs398 = (new BigNumber((_2927_v51).length)).isEqualTo(((_this).f15).minus(_2908_i4));
          let _rhs399 = _2929_v53;
          let _rhs400 = (new BigNumber((_2861_v2).length)).multipliedBy((_dafny.ZERO).minus(_module.__default.fm1(true, globalState)));
          let _lhs331 = _2929_v53;
          let _lhs332 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_2929_v53).length));
          let _lhs333 = _2932_v54;
          let _lhs334 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_2932_v54).length));
          let _lhs335 = _2909_v42;
          let _lhs336 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length));
          _2928_v52 = _rhs397;
          _lhs331[_lhs332] = _rhs398;
          _lhs333[_lhs334] = _rhs399;
          _lhs335[_lhs336] = _rhs400;
          let _2934_v56;
          let _nw547 = Array((new BigNumber(19)).toNumber());
          _nw547[(_dafny.ZERO).toNumber()] = _2909_v42;
          _nw547[(_dafny.ONE).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(2)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(3)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(4)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(5)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(6)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(7)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(8)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(9)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(10)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(11)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(12)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(13)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(14)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(15)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(16)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(17)).toNumber()] = _2909_v42;
          _nw547[(new BigNumber(18)).toNumber()] = _2909_v42;
          _2934_v56 = _nw547;
          _2934_v56 = _2934_v56;
          let _2935_v57;
          _2935_v57 = _dafny.MultiSet.fromElements(_2929_v53, (_2932_v54)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_2932_v54).length))], _2929_v53, (_2932_v54)[_module.__default.safeIndex(new BigNumber(196), new BigNumber((_2932_v54).length))]);
          let _rhs401 = (_2909_v42)[_module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length))];
          let _rhs402 = (_2935_v57).IsSubsetOf(_2935_v57);
          r2 = _rhs401;
          r1 = _rhs402;
        }
        let _2936_v58;
        _2936_v58 = _dafny.MultiSet.fromElements(_2862_v3, _2862_v3, _2862_v3, _2862_v3, _2862_v3);
        let _2937_v59;
        _2937_v59 = _dafny.Map.Empty.slice().updateUnsafe((_2909_v42)[_module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length))],_2936_v58);
        let _2938_v60;
        _2938_v60 = _module.D20.create_DC50(!(_2862_v3), (_this).f15, (_2909_v42)[_module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length))], _2862_v3, !(!_dafny.Seq.contains(_2861_v2, _2911_v43)));
        let _2939_v61;
        _2939_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,_2937_v59);
        let _2940_v62;
        let _nw548 = Array((new BigNumber(29)).toNumber());
        _nw548[(_dafny.ZERO).toNumber()] = _2862_v3;
        _nw548[(_dafny.ONE).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(2)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(3)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(4)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(5)).toNumber()] = (p0)[_module.__default.safeIndex((_this).f15, new BigNumber((p0).length))];
        _nw548[(new BigNumber(6)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(7)).toNumber()] = true;
        _nw548[(new BigNumber(8)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(9)).toNumber()] = false;
        _nw548[(new BigNumber(10)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(11)).toNumber()] = !(_2862_v3);
        _nw548[(new BigNumber(12)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(13)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(14)).toNumber()] = true;
        _nw548[(new BigNumber(15)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(16)).toNumber()] = !(_2862_v3);
        _nw548[(new BigNumber(17)).toNumber()] = true;
        _nw548[(new BigNumber(18)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(19)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(20)).toNumber()] = false;
        _nw548[(new BigNumber(21)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(22)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(23)).toNumber()] = !(_2862_v3);
        _nw548[(new BigNumber(24)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(25)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(26)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(27)).toNumber()] = _2862_v3;
        _nw548[(new BigNumber(28)).toNumber()] = _2862_v3;
        _2940_v62 = _nw548;
        let _2941_v63;
        let _nw549 = Array((new BigNumber(27)).toNumber()).fill([]);
        _2941_v63 = _nw549;
        let _2942_v64;
        _2942_v64 = _dafny.Seq.of(_2941_v63, _2941_v63);
        let _2943_v65;
        _2943_v65 = _module.D5.create_DC10((_2942_v64)[_module.__default.safeIndex((_dafny.ZERO).minus((_2909_v42)[_module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length))]), new BigNumber((_2942_v64).length))]);
        let _2944_v66;
        _2944_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2940_v62,_2943_v65);
        let _2945_v67;
        _2945_v67 = _dafny.Map.Empty.slice().updateUnsafe((_2909_v42)[_module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length))],_2862_v3);
        let _2946_v68;
        _2946_v68 = _module.D0.create_DC1((_this).f15);
        let _rhs403 = ((((_2939_v61).contains(!(_2862_v3))) ? ((_2939_v61).get(!(_2862_v3))) : (_module.__default.fm51(new BigNumber(-345), globalState)))).update(new BigNumber(((_2944_v66).Merge(_2944_v66)).length), _2936_v58);
        let _rhs404 = _module.D20.create_DC50(!(_2862_v3), _2908_i4, (_this).f15, !((_2945_v67).update((_2909_v42)[_module.__default.safeIndex(new BigNumber(710), new BigNumber((_2909_v42).length))], _2862_v3)).contains(_2908_i4), !(!((false) === (false))));
        let _rhs405 = (_2912_v44).Merge(_2912_v44);
        let _rhs406 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,(_2946_v68).dtor_cf1);
        _2937_v59 = _rhs403;
        _2938_v60 = _rhs404;
        _2912_v44 = _rhs405;
        _2912_v44 = _rhs406;
      }
      let _2947_v69;
      _2947_v69 = _dafny.MultiSet.fromElements(_2862_v3, _2862_v3, _2862_v3, _2862_v3);
      let _2948_v70;
      _2948_v70 = _module.D8.create_DC20(new BigNumber((_2947_v69).cardinality()), new BigNumber(-620), (_this).f15);
      let _source27 = _module.__default.fm52(_2948_v70, (_this).f15, globalState);
      if (_source27.is_DC7) {
        let _2949___mcc_h2 = (_source27).cf7;
        let _2950_cf7 = _2949___mcc_h2;
        let _2951_v71;
        let _init83 = ((_2952_v3) => function (_2953_i8) {
          return _2952_v3;
        })(_2862_v3);
        let _nw550 = Array((new BigNumber(8)).toNumber());
        for (let _i0_83 = 0; _i0_83 < new BigNumber(_nw550.length); _i0_83++) {
          _nw550[_i0_83] = _init83(new BigNumber(_i0_83));
        }
        _2951_v71 = _nw550;
        let _index490 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_2951_v71).length));
        (_2951_v71)[_index490] = _2862_v3;
        let _2954_v72;
        _2954_v72 = _dafny.MultiSet.fromElements(_2861_v2, _2861_v2, _2861_v2, _2861_v2);
        let _index491 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_2951_v71).length));
        (_2951_v71)[_index491] = (_2954_v72).IsDisjointFrom(((_2862_v3) ? (_2954_v72) : (_2954_v72)));
        _2862_v3 = (_2951_v71)[_module.__default.safeIndex(new BigNumber(37), new BigNumber((_2951_v71).length))];
        let _2955_v73;
        _2955_v73 = _dafny.Set.fromElements(_module.__default.fm1(true, globalState), _2950_cf7);
        let _2956_v75;
        _2956_v75 = _dafny.Seq.of(_2955_v73, function () {
          let _coll90 = new _dafny.Set();
          for (const _compr_90 of _dafny.IntegerRange(new BigNumber(616), new BigNumber(728))) {
            let _2957_v74 = _compr_90;
            if (((new BigNumber(616)).isLessThanOrEqualTo(_2957_v74)) && ((_2957_v74).isLessThan(new BigNumber(728)))) {
              _coll90.add((_2957_v74).minus(_2950_cf7));
            }
          }
          return _coll90;
        }(), (_2955_v73).Difference(_2955_v73), _2955_v73);
        _2956_v75 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2956_v75, _dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), ((_2958_v73) => function (_2959_i9) {
          return _2958_v73;
        })(_2955_v73))), _2956_v75);
        r0 = !(_2862_v3);
      } else if (_source27.is_DC8) {
        let _2960___mcc_h3 = (_source27).cf8;
        let _2961___mcc_h4 = (_source27).cf9;
        let _2962___mcc_h5 = (_source27).cf10;
        let _2963___mcc_h6 = (_source27).cf11;
        let _2964_cf11 = _2963___mcc_h6;
        let _2965_cf10 = _2962___mcc_h5;
        let _2966_cf9 = _2961___mcc_h4;
        let _2967_cf8 = _2960___mcc_h3;
        let _2968_v76;
        _2968_v76 = _dafny.Set.fromElements((_this).f15, (_this).f15);
        let _2969_v77;
        _2969_v77 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f15).isEqualTo(new BigNumber(-976)),_2968_v76);
        _2969_v77 = (_2969_v77).update(_module.__default.fm2(globalState), function () {
          let _coll91 = new _dafny.Set();
          for (const _compr_91 of _dafny.IntegerRange(new BigNumber(42), new BigNumber(375))) {
            let _2970_v78 = _compr_91;
            if (((new BigNumber(42)).isLessThanOrEqualTo(_2970_v78)) && ((_2970_v78).isLessThan(new BigNumber(375)))) {
              _coll91.add(_module.__default.safeModuloInt(_2970_v78, new BigNumber((p0).length)));
            }
          }
          return _coll91;
        }());
        r2 = (_this).f15;
        _2966_cf9 = _dafny.Seq.UnicodeFromString("wxctkag");
        let _2971_v79;
        let _nw551 = new _module.C1();
        _nw551.__ctor(new BigNumber(345));
        _2971_v79 = _nw551;
      } else {
        let _2972___mcc_h7 = (_source27).cf6;
        let _2973_cf6 = _2972___mcc_h7;
        let _2974_v80;
        _2974_v80 = _module.D15.create_DC43(_2862_v3, (_2862_v3) && (_2862_v3), ((_2862_v3) ? (_2862_v3) : (_2862_v3)));
        let _2975_v81;
        let _nw552 = Array((new BigNumber(6)).toNumber());
        _nw552[(_dafny.ZERO).toNumber()] = _2862_v3;
        _nw552[(_dafny.ONE).toNumber()] = _2862_v3;
        _nw552[(new BigNumber(2)).toNumber()] = _2862_v3;
        _nw552[(new BigNumber(3)).toNumber()] = _2862_v3;
        _nw552[(new BigNumber(4)).toNumber()] = _2862_v3;
        _nw552[(new BigNumber(5)).toNumber()] = _2862_v3;
        _2975_v81 = _nw552;
        let _2976_v82;
        _2976_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2975_v81);
        let _rhs407 = (_module.D15.create_DC43(_2862_v3, _2862_v3, (_this).fm4((_this).f15, globalState))).dtor_cf74;
        let _rhs408 = false;
        let _rhs409 = _module.__default.fm53((_2976_v82).contains((_this).f15), (_this).f15, globalState);
        r1 = _rhs407;
        r1 = _rhs408;
        _2974_v80 = _rhs409;
        let _2977_v83;
        _2977_v83 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,_2862_v3);
        let _2978_v84;
        _2978_v84 = _dafny.Set.fromElements(_2862_v3, _2862_v3, _2862_v3);
        let _2979_v85;
        _2979_v85 = _dafny.Seq.of(_2862_v3, _2862_v3, _2862_v3, (_dafny.Map.Empty.slice().updateUnsafe(true,_2862_v3)).equals(_2977_v83), !((_2978_v84).IsDisjointFrom(_module.__default.fm31((_this).f15, _2862_v3, globalState))));
        let _2980_v86;
        _2980_v86 = _dafny.Seq.of(_dafny.Seq.of(_module.__default.fm2(globalState), !(_2862_v3), _2862_v3, _2862_v3), p0);
        _2979_v85 = (_2980_v86)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-413)), new BigNumber((_2980_v86).length))];
        let _index492 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_2973_cf6).length));
        (_2973_cf6)[_index492] = new BigNumber((_2861_v2).length);
        let _index493 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_2973_cf6).length));
        (_2973_cf6)[_index493] = (_this).f15;
        let _2981_v87;
        let _nw553 = new _module.C6();
        _nw553.__ctor();
        _2981_v87 = _nw553;
        _2981_v87 = _2981_v87;
      }
      let _2982_v88;
      _2982_v88 = _dafny.Map.Empty.slice().updateUnsafe(_2861_v2,!(_2862_v3));
      let _2983_v89;
      let _2984_v90;
      let _out26;
      let _out27;
      let _outcollector10 = (_this).m4(!(new BigNumber((_2947_v69).cardinality())).isEqualTo(new BigNumber((_2982_v88).length)), globalState);
      _out26 = _outcollector10[0];
      _out27 = _outcollector10[1];
      _2983_v89 = _out26;
      _2984_v90 = _out27;
      let _2985_v91;
      _2985_v91 = _dafny.Set.fromElements((_this).f15);
      let _2986_v92;
      _2986_v92 = new _dafny.CodePoint('y'.codePointAt(0));
      let _2987_v93;
      _2987_v93 = _module.D5.create_DC12(new BigNumber((_2985_v91).length), _2862_v3, _2986_v92);
      let _source28 = _2987_v93;
      if (_source28.is_DC11) {
        let _2988___mcc_h8 = (_source28).cf14;
        let _2989_cf14 = _2988___mcc_h8;
        if (_2983_v89) {
          let _2990_v94;
          let _nw554 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
          _2990_v94 = _nw554;
          let _2991_v95;
          _2991_v95 = _dafny.Map.Empty.slice().updateUnsafe(false,_2984_v90);
          let _2992_v96;
          _2992_v96 = _dafny.Seq.of(_2991_v95);
          let _index494 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_2990_v94).length));
          (_2990_v94)[_index494] = _dafny.Seq.Concat(_2992_v96, _2992_v96);
          let _index495 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_2990_v94).length));
          (_2990_v94)[_index495] = _dafny.Seq.of(_2991_v95, _2991_v95, (_2991_v95).Merge(_2991_v95));
          let _2993_v97;
          let _nw555 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2993_v97 = _nw555;
          let _index496 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_2993_v97).length));
          (_2993_v97)[_index496] = _2861_v2;
          let _index497 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_2993_v97).length));
          (_2993_v97)[_index497] = ((true) ? (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mltwv"), _2861_v2)) : (_dafny.Seq.update(_2861_v2, _module.__default.safeIndex(_2989_cf14, new BigNumber((_2861_v2).length)), _2986_v92)));
          let _2994_v98;
          let _nw556 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
          _2994_v98 = _nw556;
          let _2995_v99;
          _2995_v99 = _dafny.Map.Empty.slice().updateUnsafe(((_2862_v3) ? (_2994_v98) : (_2994_v98)),_module.__default.fm1(_2983_v89, globalState));
          _2995_v99 = (_2995_v99).update(_2994_v98, (_this).f15);
          _2989_cf14 = (_dafny.ZERO).minus((_this).f15);
          let _2996_v100;
          _2996_v100 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f15).multipliedBy(_2989_cf14),new _dafny.CodePoint('b'.codePointAt(0)));
          _2996_v100 = ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2993_v97)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_2993_v97).length))]).length),_2986_v92)).Merge(_module.__default.fm54(_2991_v95, globalState))).Merge(_2996_v100);
        } else {
          let _2997_v101;
          _2997_v101 = _dafny.Set.fromElements(_2984_v90, !(((_2984_v90) ? (_2983_v89) : (false))));
          let _2998_v102;
          let _nw557 = Array((_dafny.ONE).toNumber());
          _nw557[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),(_this).f15);
          _2998_v102 = _nw557;
          let _2999_v103;
          _2999_v103 = _dafny.MultiSet.fromElements(new BigNumber((_2861_v2).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("qyropc"), _module.__default.safeIndex((_this).f15, new BigNumber((_dafny.Seq.UnicodeFromString("qyropc")).length)), _2986_v92)).length), _2989_cf14);
          let _rhs410 = _2997_v101;
          let _rhs411 = _2998_v102;
          let _rhs412 = !(_2999_v103).contains((_this).f15);
          let _rhs413 = _2986_v92;
          let _rhs414 = _2986_v92;
          let _lhs337 = globalState;
          _2997_v101 = _rhs410;
          _2998_v102 = _rhs411;
          _lhs337.f1 = _rhs412;
          _2986_v92 = _rhs413;
          _2986_v92 = _rhs414;
          (globalState).f6 = new BigNumber((_module.__default.fm13((_2983_v89) === (_2862_v3), _2986_v92, new BigNumber(218), globalState)).length);
          let _3000_v104;
          let _nw558 = Array((new BigNumber(7)).toNumber());
          _nw558[(_dafny.ZERO).toNumber()] = _2947_v69;
          _nw558[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements((_module.D20.create_DC50(_2984_v90, (_dafny.ZERO).minus(_2989_cf14), (_this).f15, !(_2984_v90), _2862_v3)).dtor_cf90);
          _nw558[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(true);
          _nw558[(new BigNumber(3)).toNumber()] = _2947_v69;
          _nw558[(new BigNumber(4)).toNumber()] = _2947_v69;
          _nw558[(new BigNumber(5)).toNumber()] = (_2947_v69).Intersect(_2947_v69);
          _nw558[(new BigNumber(6)).toNumber()] = _2947_v69;
          _3000_v104 = _nw558;
          _3000_v104 = _3000_v104;
          (globalState).f6 = (_this).f15;
          _2984_v90 = _2984_v90;
        }
        let _3001_v105;
        _3001_v105 = _dafny.Set.fromElements(_2983_v89);
        if ((_2862_v3) || ((_3001_v105).IsProperSubsetOf(_3001_v105))) {
          let _3002_v106;
          _3002_v106 = _dafny.MultiSet.fromElements(_2989_cf14, (_this).f15);
          _3002_v106 = _3002_v106;
          _2862_v3 = _2862_v3;
          let _3003_v107;
          let _nw559 = new _module.C6();
          _nw559.__ctor();
          _3003_v107 = _nw559;
          _3003_v107 = _3003_v107;
          _2862_v3 = _module.__default.fm2(globalState);
          let _3004_v108;
          let _nw560 = new _module.C8();
          _nw560.__ctor();
          _3004_v108 = _nw560;
        } else {
          (globalState).f1 = ((_2947_v69).IsProperSubsetOf(_2947_v69)) && (_2862_v3);
          let _3005_v109;
          let _nw561 = new _module.C6();
          _nw561.__ctor();
          _3005_v109 = _nw561;
          let _3006_v110;
          _3006_v110 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_2989_cf14, _2989_cf14),_3005_v109);
          _3006_v110 = _3006_v110;
          let _3007_v111;
          _3007_v111 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(_2989_cf14, new BigNumber(-338)), _module.__default.safeDivisionInt(_2989_cf14, new BigNumber((p0).length)), _module.__default.safeDivisionInt((_this).f15, (_this).f15));
          _3007_v111 = _3007_v111;
          let _3008_v112;
          _3008_v112 = _dafny.Map.Empty.slice().updateUnsafe(_2983_v89,_3001_v105);
          (globalState).f1 = !(((((_3008_v112).contains(_2984_v90)) ? ((_3008_v112).get(_2984_v90)) : (_3001_v105))).IsProperSubsetOf((_3001_v105).Intersect(_3001_v105)));
          (globalState).f6 = ((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f15))).multipliedBy((_this).f15);
        }
        let _3009_v113;
        let _nw562 = new _module.C1();
        _nw562.__ctor(new BigNumber((_2861_v2).length));
        _3009_v113 = _nw562;
        let _3010_v114;
        _3010_v114 = _dafny.Seq.of(_3009_v113, _3009_v113);
        let _3011_v115;
        _3011_v115 = _dafny.Seq.of(_module.__default.fm32((_this).f15, (_this).f15, _2862_v3, globalState), _2859_v0);
        let _3012_v116;
        _3012_v116 = _dafny.Map.Empty.slice().updateUnsafe(_2983_v89,_2984_v90);
        let _3013_v117;
        _3013_v117 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),(_this).f15);
        let _3014_v118;
        let _nw563 = Array((new BigNumber(24)).toNumber());
        _nw563[(_dafny.ZERO).toNumber()] = (new BigNumber((_3010_v114).length)).multipliedBy((_this).f15);
        _nw563[(_dafny.ONE).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(2)).toNumber()] = new BigNumber((_2859_v0).length);
        _nw563[(new BigNumber(3)).toNumber()] = (_this).f15;
        _nw563[(new BigNumber(4)).toNumber()] = (new BigNumber((p0).length)).minus((_this).f15);
        _nw563[(new BigNumber(5)).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(6)).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(7)).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(8)).toNumber()] = (_this).f15;
        _nw563[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-849), (_this).f15);
        _nw563[(new BigNumber(10)).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(11)).toNumber()] = new BigNumber(-277);
        _nw563[(new BigNumber(12)).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat((_3011_v115)[_module.__default.safeIndex((_this).f15, new BigNumber((_3011_v115).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(265)), function (_3015_i10) {
          return (_this).f15;
        }))).length);
        _nw563[(new BigNumber(14)).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(15)).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(16)).toNumber()] = new BigNumber((_3012_v116).length);
        _nw563[(new BigNumber(17)).toNumber()] = (((_3013_v117).contains(false)) ? ((_3013_v117).get(false)) : ((_dafny.ZERO).minus((_this).f15)));
        _nw563[(new BigNumber(18)).toNumber()] = new BigNumber(((_3001_v105).Difference(_3001_v105)).length);
        _nw563[(new BigNumber(19)).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(20)).toNumber()] = _2989_cf14;
        _nw563[(new BigNumber(21)).toNumber()] = (_2989_cf14).minus(_2989_cf14);
        _nw563[(new BigNumber(22)).toNumber()] = new BigNumber(355);
        _nw563[(new BigNumber(23)).toNumber()] = (_this).f15;
        _3014_v118 = _nw563;
        let _index498 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_3014_v118).length));
        (_3014_v118)[_index498] = (_this).f15;
        let _index499 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_3014_v118).length));
        (_3014_v118)[_index499] = _2989_cf14;
        _2947_v69 = _2947_v69;
      } else if (_source28.is_DC12) {
        let _3016___mcc_h9 = (_source28).cf15;
        let _3017___mcc_h10 = (_source28).cf16;
        let _3018___mcc_h11 = (_source28).cf17;
        let _3019_cf17 = _3018___mcc_h11;
        let _3020_cf16 = _3017___mcc_h10;
        let _3021_cf15 = _3016___mcc_h9;
        let _3022_v119;
        let _nw564 = new _module.C8();
        _nw564.__ctor();
        _3022_v119 = _nw564;
        let _3023_v120;
        _3023_v120 = _dafny.MultiSet.fromElements(_3019_cf17, _3019_cf17);
        let _3024_v121;
        _3024_v121 = _module.D20.create_DC50(_3020_cf16, _3021_cf15, (_this).f15, _3020_cf16, !(_2862_v3));
        let _3025_v122;
        _3025_v122 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,_3021_cf15);
        let _3026_v123;
        let _nw565 = Array((new BigNumber(24)).toNumber());
        _nw565[(_dafny.ZERO).toNumber()] = true;
        _nw565[(_dafny.ONE).toNumber()] = _2983_v89;
        _nw565[(new BigNumber(2)).toNumber()] = (true) || (!(!(_2862_v3)));
        _nw565[(new BigNumber(3)).toNumber()] = _2984_v90;
        _nw565[(new BigNumber(4)).toNumber()] = true;
        _nw565[(new BigNumber(5)).toNumber()] = (_2983_v89) || (_2984_v90);
        _nw565[(new BigNumber(6)).toNumber()] = (_2983_v89) && (_2984_v90);
        _nw565[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), _module.__default.fm29(_2983_v89, new BigNumber((_2947_v69).cardinality()), (_this).f15, _2862_v3, globalState), _3019_cf17, _module.__default.fm29(_module.__default.fm2(globalState), (_2987_v93).dtor_cf15, (_this).f15, _2984_v90, globalState), _3019_cf17)).IsProperSubsetOf(_3023_v120);
        _nw565[(new BigNumber(8)).toNumber()] = _2862_v3;
        _nw565[(new BigNumber(9)).toNumber()] = _3020_cf16;
        _nw565[(new BigNumber(10)).toNumber()] = (p0)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_3020_cf16,(_this).f15)).length), new BigNumber((p0).length))];
        _nw565[(new BigNumber(11)).toNumber()] = _3020_cf16;
        _nw565[(new BigNumber(12)).toNumber()] = !((_module.__default.fm2(globalState)) === (_2983_v89));
        _nw565[(new BigNumber(13)).toNumber()] = (_3024_v121).dtor_cf86;
        _nw565[(new BigNumber(14)).toNumber()] = _3020_cf16;
        _nw565[(new BigNumber(15)).toNumber()] = _3020_cf16;
        _nw565[(new BigNumber(16)).toNumber()] = _3020_cf16;
        _nw565[(new BigNumber(17)).toNumber()] = ((_this).f15).isLessThanOrEqualTo((((_3025_v122).contains(_3020_cf16)) ? ((_3025_v122).get(_3020_cf16)) : (_3021_cf15)));
        _nw565[(new BigNumber(18)).toNumber()] = _2984_v90;
        _nw565[(new BigNumber(19)).toNumber()] = _2983_v89;
        _nw565[(new BigNumber(20)).toNumber()] = (_2862_v3) && (_3020_cf16);
        _nw565[(new BigNumber(21)).toNumber()] = !(_2984_v90) || (_3020_cf16);
        _nw565[(new BigNumber(22)).toNumber()] = _2983_v89;
        _nw565[(new BigNumber(23)).toNumber()] = _2862_v3;
        _3026_v123 = _nw565;
        let _index500 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_3026_v123).length));
        (_3026_v123)[_index500] = _2862_v3;
        let _index501 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_3026_v123).length));
        (_3026_v123)[_index501] = _3020_cf16;
        let _3027_v124;
        let _nw566 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _3027_v124 = _nw566;
        let _3028_v125;
        _3028_v125 = _module.D3.create_DC6(_3027_v124);
        let _source29 = _3028_v125;
        if (_source29.is_DC7) {
          let _3029___mcc_h14 = (_source29).cf7;
          let _3030_cf7 = _3029___mcc_h14;
          let _3031_v126;
          _3031_v126 = _module.D16.create_DC45(new BigNumber(478), _2986_v92, _3021_cf15, _2984_v90);
          let _3032_v127;
          _3032_v127 = _dafny.Map.Empty.slice().updateUnsafe(_3021_cf15,_3031_v126);
          let _3033_v128;
          _3033_v128 = _dafny.Set.fromElements(_module.__default.fm2(globalState));
          let _pat_let_tv71 = _3021_cf15;
          _3032_v127 = (_3032_v127).update(new BigNumber((_3033_v128).length), function (_pat_let102_0) {
            return function (_3034_dt__update__tmp_h0) {
              return function (_pat_let103_0) {
                return function (_3035_dt__update_hcf80_h0) {
                  return _module.D16.create_DC45((_3034_dt__update__tmp_h0).dtor_cf78, (_3034_dt__update__tmp_h0).dtor_cf79, _3035_dt__update_hcf80_h0, (_3034_dt__update__tmp_h0).dtor_cf81);
                }(_pat_let103_0);
              }(_pat_let_tv71);
            }(_pat_let102_0);
          }(_3031_v126));
          let _3036_v129;
          _3036_v129 = _module.D1.create_DC4(_2986_v92);
          _3036_v129 = _3036_v129;
          let _3037_v130;
          _3037_v130 = _module.D3.create_DC7((_2859_v0)[_module.__default.safeIndex(new BigNumber(-905), new BigNumber((_2859_v0).length))]);
          (globalState).f6 = (_3037_v130).dtor_cf7;
          let _3038_v131;
          let _nw567 = new _module.C5();
          _nw567.__ctor((_this).f15);
          _3038_v131 = _nw567;
        } else if (_source29.is_DC8) {
          let _3039___mcc_h15 = (_source29).cf8;
          let _3040___mcc_h16 = (_source29).cf9;
          let _3041___mcc_h17 = (_source29).cf10;
          let _3042___mcc_h18 = (_source29).cf11;
          let _3043_cf11 = _3042___mcc_h18;
          let _3044_cf10 = _3041___mcc_h17;
          let _3045_cf9 = _3040___mcc_h16;
          let _3046_cf8 = _3039___mcc_h15;
          let _3047_v132;
          _3047_v132 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? ((_dafny.ZERO).minus(_module.__default.fm1(_2984_v90, globalState))) : (_3021_cf15)),!(new BigNumber((_3025_v122).length)).isEqualTo(_3021_cf15));
          _3047_v132 = (_3047_v132).update(_3021_cf15, _2984_v90);
          let _3048_v133;
          let _nw568 = new _module.C4();
          _nw568.__ctor();
          _3048_v133 = _nw568;
          (globalState).f6 = _3021_cf15;
          (globalState).f14 = _3026_v123;
        } else {
          let _3049___mcc_h19 = (_source29).cf6;
          let _3050_cf6 = _3049___mcc_h19;
          let _rhs415 = (_2983_v89) === (_2983_v89);
          let _rhs416 = _2983_v89;
          r1 = _rhs415;
          _2984_v90 = _rhs416;
          let _3051_v134;
          let _nw569 = new _module.C1();
          _nw569.__ctor((_this).f15);
          _3051_v134 = _nw569;
          let _index502 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_3026_v123).length));
          (_3026_v123)[_index502] = (_3051_v134).fm24(_3020_cf16, _module.__default.fm39(_2984_v90, globalState), globalState);
          let _3052_v135;
          let _nw570 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _3052_v135 = _nw570;
          _3052_v135 = _3052_v135;
        }
        let _3053_v138;
        _3053_v138 = _dafny.Seq.of(_2859_v0);
        let _3054_v139;
        _3054_v139 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_3021_cf15,(_3026_v123)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_3026_v123).length))]),(_this).f15);
        let _3055_v140;
        let _nw571 = new _module.C3();
        _nw571.__ctor((_3026_v123)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_3026_v123).length))], (_this).f15, _3054_v139);
        _3055_v140 = _nw571;
        let _3056_v141;
        _3056_v141 = _dafny.Set.fromElements(_3055_v140, _3055_v140);
        let _3057_v142;
        _3057_v142 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(_3021_cf15));
        let _3058_v143;
        _3058_v143 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((function () {
          let _coll92 = new _dafny.Map();
          for (const _compr_92 of (function () {
            let _coll93 = new _dafny.Map();
            for (const _compr_93 of ((_3053_v138)[_module.__default.safeIndex(new BigNumber((_3056_v141).length), new BigNumber((_3053_v138).length))]).Elements) {
              let _3059_v137 = _compr_93;
              if (_dafny.Seq.contains((_3053_v138)[_module.__default.safeIndex(new BigNumber((_3056_v141).length), new BigNumber((_3053_v138).length))], _3059_v137)) {
                _coll93.push([(_3059_v137).plus(_3021_cf15),(_3026_v123)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_3026_v123).length))]]);
              }
            }
            return _coll93;
          }()).Keys.Elements) {
            let _3060_v136 = _compr_92;
            if ((function () {
              let _coll94 = new _dafny.Map();
              for (const _compr_94 of ((_3053_v138)[_module.__default.safeIndex(new BigNumber((_3056_v141).length), new BigNumber((_3053_v138).length))]).Elements) {
                let _3059_v137 = _compr_94;
                if (_dafny.Seq.contains((_3053_v138)[_module.__default.safeIndex(new BigNumber((_3056_v141).length), new BigNumber((_3053_v138).length))], _3059_v137)) {
                  _coll94.push([(_3059_v137).plus(_3021_cf15),(_3026_v123)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_3026_v123).length))]]);
                }
              }
              return _coll94;
            }()).contains(_3060_v136)) {
              _coll92.push([(_3060_v136).multipliedBy((_this).f15),new BigNumber((_2861_v2).length)]);
            }
          }
          return _coll92;
        }()).length)).isLessThanOrEqualTo(new BigNumber((_2985_v91).length)),(_2985_v91).IsSubsetOf((((_3057_v142).contains(_2984_v90)) ? ((_3057_v142).get(_2984_v90)) : (_dafny.Set.fromElements(new BigNumber((_2861_v2).length))))));
        let _index503 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_3026_v123).length));
        (_3026_v123)[_index503] = (((_3058_v143).contains(_2862_v3)) ? ((_3058_v143).get(_2862_v3)) : (_2862_v3));
      } else if (_source28.is_DC10) {
        let _3061___mcc_h12 = (_source28).cf13;
        let _3062_cf13 = _3061___mcc_h12;
        (globalState).f6 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2859_v0, _dafny.Seq.update(_2859_v0, _module.__default.safeIndex((_this).f15, new BigNumber((_2859_v0).length)), (_this).f15)), _dafny.Seq.Concat(_2859_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-720)), function (_3063_i11) {
          return (_this).f15;
        })))).length);
        _2985_v91 = _2985_v91;
        if (_2984_v90) {
          let _3064_v144;
          let _3065_v145;
          let _out28;
          let _out29;
          let _outcollector11 = (_this).m4(true, globalState);
          _out28 = _outcollector11[0];
          _out29 = _outcollector11[1];
          _3064_v144 = _out28;
          _3065_v145 = _out29;
          _3064_v144 = ((!(_3064_v144)) ? (!(_2983_v89)) : (_3064_v144));
          (globalState).f6 = (new BigNumber(415)).minus((_this).f15);
          let _3066_v146;
          _3066_v146 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2984_v90);
          let _3067_v147;
          _3067_v147 = _dafny.Map.Empty.slice().updateUnsafe(_3066_v146,(_this).f15);
          let _3068_v148;
          let _nw572 = new _module.C2();
          _nw572.__ctor(new BigNumber(71), _module.__default.safeModuloInt(new BigNumber((_3066_v146).length), (_this).f15), _3067_v147);
          _3068_v148 = _nw572;
          let _3069_v149;
          let _nw573 = Array((new BigNumber(5)).toNumber()).fill(false);
          _3069_v149 = _nw573;
          let _3070_v150;
          _3070_v150 = _dafny.Map.Empty.slice().updateUnsafe(_3069_v149,_3067_v147);
          let _3071_v151;
          _3071_v151 = (((_3070_v150).contains(_3069_v149)) ? ((_3070_v150).get(_3069_v149)) : (_3067_v147));
          _3071_v151 = _3071_v151;
        } else {
          let _3072_v152;
          let _init84 = ((_3073_v89) => function (_3074_i12) {
            return !(_3073_v89);
          })(_2983_v89);
          let _nw574 = Array((new BigNumber(19)).toNumber());
          for (let _i0_84 = 0; _i0_84 < new BigNumber(_nw574.length); _i0_84++) {
            _nw574[_i0_84] = _init84(new BigNumber(_i0_84));
          }
          _3072_v152 = _nw574;
          let _index504 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_3072_v152).length));
          (_3072_v152)[_index504] = _2983_v89;
          let _3075_v153;
          let _nw575 = Array((new BigNumber(26)).toNumber()).fill([]);
          _3075_v153 = _nw575;
          let _3076_v154;
          _3076_v154 = _module.D10.create_DC26(_3075_v153);
          let _3077_v155;
          _3077_v155 = _dafny.Set.fromElements(_3076_v154, _3076_v154, _module.D10.create_DC26(_3075_v153), _3076_v154);
          let _index505 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_3072_v152).length));
          (_3072_v152)[_index505] = (_3077_v155).IsSubsetOf(_3077_v155);
          let _3078_v156;
          let _nw576 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _3078_v156 = _nw576;
          let _index506 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_3078_v156).length));
          (_3078_v156)[_index506] = (_this).f15;
          let _index507 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_3078_v156).length));
          (_3078_v156)[_index507] = (_this).f15;
          let _index508 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_3078_v156).length));
          (_3078_v156)[_index508] = (_this).f15;
          r2 = _module.__default.safeDivisionInt((_3078_v156)[_module.__default.safeIndex(new BigNumber(127), new BigNumber((_3078_v156).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-962)), ((_3079_v92) => function (_3080_i13) {
            return _3079_v92;
          })(_2986_v92))).length));
          (globalState).f6 = _module.__default.safeModuloInt((_3078_v156)[_module.__default.safeIndex(new BigNumber(127), new BigNumber((_3078_v156).length))], new BigNumber((_2861_v2).length));
        }
        (globalState).f6 = ((_this).f15).minus((_this).f15);
      } else {
        let _3081___mcc_h13 = (_source28).cf18;
        let _3082_cf18 = _3081___mcc_h13;
        if (_2862_v3) {
          (globalState).f1 = _2862_v3;
          let _3083_v157;
          _3083_v157 = _dafny.Set.fromElements(_2862_v3);
          let _3084_v158;
          let _nw577 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _3084_v158 = _nw577;
          let _3085_v159;
          _3085_v159 = _dafny.Set.fromElements(_3084_v158, _3084_v158, _3084_v158, _3084_v158, _3084_v158);
          let _rhs417 = (_this).f15;
          let _rhs418 = (_3083_v157).Difference((_dafny.Set.fromElements(_2862_v3, !(_2983_v89), false)).Difference(_dafny.Set.fromElements(_2862_v3, _2983_v89)));
          let _rhs419 = ((_3085_v159).Intersect(_3085_v159)).IsProperSubsetOf(_3085_v159);
          let _lhs338 = globalState;
          let _lhs339 = globalState;
          _lhs338.f6 = _rhs417;
          _3083_v157 = _rhs418;
          _lhs339.f1 = _rhs419;
          let _3086_v160;
          _3086_v160 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("dd"), _2861_v2, _dafny.Seq.UnicodeFromString("kvbevapf"), _2861_v2);
          let _3087_v161;
          _3087_v161 = _dafny.Seq.of((_3086_v160)[_module.__default.safeIndex((_this).f15, new BigNumber((_3086_v160).length))], _2861_v2, _module.__default.fm21(globalState), _2861_v2);
          let _3088_v162;
          _3088_v162 = _module.D8.create_DC21((_3086_v160)[_module.__default.safeIndex((_this).f15, new BigNumber((_3086_v160).length))], _3083_v157);
          _3088_v162 = _3088_v162;
          let _index509 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_3084_v158).length));
          (_3084_v158)[_index509] = (_this).f15;
          let _3089_v163;
          _3089_v163 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,!(_2984_v90));
          let _index510 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_3084_v158).length));
          (_3084_v158)[_index510] = ((_this).f15).multipliedBy(new BigNumber((_3089_v163).length));
          let _3090_v164;
          let _nw578 = new _module.C7();
          _nw578.__ctor();
          _3090_v164 = _nw578;
        } else {
          let _3091_v165;
          _3091_v165 = _dafny.Set.fromElements(_2986_v92);
          _3091_v165 = _3091_v165;
          let _3092_v166;
          _3092_v166 = _module.D1.create_DC4(_2986_v92);
          let _3093_v167;
          let _nw579 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _3093_v167 = _nw579;
          let _rhs420 = _3093_v167;
          let _rhs421 = _3092_v166;
          let _rhs422 = (p0)[_module.__default.safeIndex((_2859_v0)[_module.__default.safeIndex(new BigNumber(-491), new BigNumber((_2859_v0).length))], new BigNumber((p0).length))];
          let _rhs423 = !((new BigNumber(727)).isLessThan((_this).f15));
          let _lhs340 = globalState;
          let _lhs341 = globalState;
          _lhs340.f9 = _rhs420;
          _3092_v166 = _rhs421;
          _lhs341.f1 = _rhs422;
          _2984_v90 = _rhs423;
          let _3094_v168;
          _3094_v168 = _dafny.Set.fromElements(_2862_v3, _2862_v3);
          _2862_v3 = !(_3094_v168).equals((_3094_v168).Union(_3094_v168));
          let _3095_v169;
          _3095_v169 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,_2983_v89);
          let _3096_v170;
          _3096_v170 = _dafny.MultiSet.fromElements(_module.__default.fm1(_2983_v89, globalState), (new BigNumber((((_module.D1.create_DC3(_3095_v169)).dtor_cf3).update(_2983_v89, false)).length)).minus(new BigNumber(255)), (_this).f15, (_this).f15);
          _3096_v170 = _3096_v170;
          r2 = (new BigNumber((p0).length)).minus((_dafny.ZERO).minus((_this).f15));
        }
        _2986_v92 = _2986_v92;
        let _3097_v171;
        _3097_v171 = _dafny.Seq.of(_2983_v89, _2983_v89, _2983_v89);
        let _3098_v172;
        _3098_v172 = _dafny.MultiSet.fromElements(new BigNumber(940));
        let _3099_v173;
        _3099_v173 = _module.D11.create_DC31(_dafny.Seq.of(_2983_v89), (_dafny.ZERO).minus((_this).f15), new BigNumber((_3098_v172).cardinality()), (_dafny.ZERO).minus((_this).f15), (_this).f15);
        _3097_v171 = (_3099_v173).dtor_cf54;
        (globalState).f6 = (_this).f15;
      }
      r0 = _2862_v3;
      r1 = ((_this).f15).isLessThan((_this).f15);
      let _3100_v174;
      _3100_v174 = _dafny.Map.Empty.slice().updateUnsafe(_2859_v0,new BigNumber((_2860_v1).length));
      let _3101_v175;
      _3101_v175 = _dafny.Map.Empty.slice().updateUnsafe(_2862_v3,(_dafny.ZERO).minus(new BigNumber((_3100_v174).length)));
      r2 = (_dafny.ZERO).minus(((!(_2984_v90) || (_2862_v3)) ? ((_this).f15) : (new BigNumber((_3101_v175).length))));
      return [r0, r1, r2];
    }
    m4(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _hi18 = (_this).f15;
      for (let _3102_i0 = (_this).f15; _3102_i0.isLessThan(_hi18); _3102_i0 = _3102_i0.plus(_dafny.ONE)) {
        if (((p0) ? (_module.__default.fm2(globalState)) : (p0))) {
          (globalState).f6 = _3102_i0;
          let _3103_v0;
          let _init85 = function (_3104_i1) {
            return true;
          };
          let _nw580 = Array((new BigNumber(19)).toNumber());
          for (let _i0_85 = 0; _i0_85 < new BigNumber(_nw580.length); _i0_85++) {
            _nw580[_i0_85] = _init85(new BigNumber(_i0_85));
          }
          _3103_v0 = _nw580;
          let _index511 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_3103_v0).length));
          (_3103_v0)[_index511] = !(p0);
          let _index512 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_3103_v0).length));
          (_3103_v0)[_index512] = false;
          let _3105_v1;
          _3105_v1 = _dafny.Seq.of(p0, (_3103_v0)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_3103_v0).length))], (_3103_v0)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_3103_v0).length))], !(p0), (_3103_v0)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_3103_v0).length))]);
          _3105_v1 = _3105_v1;
          (globalState).f1 = (_3103_v0)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_3103_v0).length))];
          r1 = p0;
        } else {
          let _3106_v2;
          _3106_v2 = _dafny.Map.Empty.slice().updateUnsafe(_3102_i0,p0);
          let _3107_v3;
          _3107_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_3106_v2);
          let _3108_v5;
          _3108_v5 = _dafny.Map.Empty.slice().updateUnsafe((((_3107_v3).contains((_dafny.ZERO).minus(_3102_i0))) ? ((_3107_v3).get((_dafny.ZERO).minus(_3102_i0))) : (function () {
            let _coll95 = new _dafny.Map();
            for (const _compr_95 of _dafny.IntegerRange(new BigNumber(728), new BigNumber(190))) {
              let _3109_v4 = _compr_95;
              if (((new BigNumber(728)).isLessThanOrEqualTo(_3109_v4)) && ((_3109_v4).isLessThan(new BigNumber(190)))) {
                _coll95.push([_module.__default.safeModuloInt(_3109_v4, (_this).f15),p0]);
              }
            }
            return _coll95;
          }())),_3102_i0);
          let _3110_v6;
          let _nw581 = new _module.C3();
          _nw581.__ctor(!(p0) || (p0), new BigNumber(233), _3108_v5);
          _3110_v6 = _nw581;
          let _3111_v7;
          _3111_v7 = _dafny.Seq.of(p0);
          let _3112_v8;
          _3112_v8 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.of(false), _3111_v7), _dafny.Seq.of((_3110_v6).f18, (_3110_v6).f18, !(p0), p0), _3111_v7, _dafny.Seq.of(false), _dafny.Seq.Concat(_dafny.Seq.of(p0), _3111_v7));
          _3112_v8 = (_3112_v8).Difference(_3112_v8);
          let _3113_v9;
          _3113_v9 = _dafny.MultiSet.fromElements(p0);
          let _3114_v10;
          let _nw582 = Array((new BigNumber(9)).toNumber());
          _nw582[(_dafny.ZERO).toNumber()] = new BigNumber(512);
          _nw582[(_dafny.ONE).toNumber()] = (_3110_v6).f19;
          _nw582[(new BigNumber(2)).toNumber()] = ((_this).f15).plus(_3102_i0);
          _nw582[(new BigNumber(3)).toNumber()] = new BigNumber(729);
          _nw582[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(_3102_i0, _3102_i0);
          _nw582[(new BigNumber(5)).toNumber()] = new BigNumber(154);
          _nw582[(new BigNumber(6)).toNumber()] = _3102_i0;
          _nw582[(new BigNumber(7)).toNumber()] = new BigNumber((_module.__default.fm12(_3113_v9, new BigNumber(759), globalState)).length);
          _nw582[(new BigNumber(8)).toNumber()] = (_3110_v6).f19;
          _3114_v10 = _nw582;
          let _3115_v11;
          _3115_v11 = _dafny.Seq.UnicodeFromString("tviyhhlfk");
          let _index513 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_3114_v10).length));
          (_3114_v10)[_index513] = new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("mpbabn"), _3115_v11)).length);
          let _index514 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_3114_v10).length));
          (_3114_v10)[_index514] = _module.__default.safeModuloInt((_this).f15, (_this).fm5(!(p0), globalState));
          let _3116_v12;
          _3116_v12 = _dafny.Map.Empty.slice().updateUnsafe((_3114_v10)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_3114_v10).length))],(_this).f15);
          let _3117_v13;
          _3117_v13 = _dafny.Seq.of((((_3116_v12).contains(_3102_i0)) ? ((_3116_v12).get(_3102_i0)) : (new BigNumber((_3113_v9).cardinality()))));
          let _3118_v14;
          _3118_v14 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21(globalState),_3117_v13);
          _3118_v14 = (_3118_v14).update(_module.__default.fm21(globalState), _3117_v13);
          (globalState).f6 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(981)), function (_3119_i2) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          })).length);
        }
        let _3120_v15;
        _3120_v15 = _dafny.Seq.UnicodeFromString("tpvilna");
        _3120_v15 = _3120_v15;
        let _3121_v16;
        let _nw583 = Array((new BigNumber(6)).toNumber()).fill([]);
        _3121_v16 = _nw583;
        let _3122_v17;
        let _init86 = ((_3123_p0) => function (_3124_i3) {
          return _dafny.Seq.of(_3123_p0, _3123_p0, _3123_p0);
        })(p0);
        let _nw584 = Array((new BigNumber(5)).toNumber());
        for (let _i0_86 = 0; _i0_86 < new BigNumber(_nw584.length); _i0_86++) {
          _nw584[_i0_86] = _init86(new BigNumber(_i0_86));
        }
        _3122_v17 = _nw584;
        let _index515 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_3121_v16).length));
        (_3121_v16)[_index515] = _3122_v17;
        let _index516 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_3121_v16).length));
        (_3121_v16)[_index516] = _3122_v17;
        (globalState).f1 = p0;
      }
      (globalState).f6 = (_this).f15;
      let _3125_v18;
      _3125_v18 = _dafny.Seq.of((_this).f15);
      let _3126_v19;
      _3126_v19 = _dafny.Seq.of(_3125_v18);
      let _3127_v20;
      _3127_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-546),_3125_v18);
      let _3128_v21;
      let _nw585 = Array((new BigNumber(18)).toNumber());
      _nw585[(_dafny.ZERO).toNumber()] = _3125_v18;
      _nw585[(_dafny.ONE).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(2)).toNumber()] = _dafny.Seq.of((_this).f15, (_this).f15, (_this).f15, new BigNumber(994));
      _nw585[(new BigNumber(3)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(4)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(5)).toNumber()] = (_3126_v19)[_module.__default.safeIndex(new BigNumber(-473), new BigNumber((_3126_v19).length))];
      _nw585[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(340)), function (_3129_i4) {
        return new BigNumber(670);
      });
      _nw585[(new BigNumber(7)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(8)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(9)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(10)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(11)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(12)).toNumber()] = (((_3127_v20).contains((_this).f15)) ? ((_3127_v20).get((_this).f15)) : (_3125_v18));
      _nw585[(new BigNumber(13)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(14)).toNumber()] = _dafny.Seq.of((_this).f15, new BigNumber((_dafny.Seq.UnicodeFromString("doh")).length));
      _nw585[(new BigNumber(15)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(16)).toNumber()] = _3125_v18;
      _nw585[(new BigNumber(17)).toNumber()] = _3125_v18;
      _3128_v21 = _nw585;
      let _3130_v22;
      _3130_v22 = _dafny.Seq.UnicodeFromString("osjud");
      let _3131_v23;
      _3131_v23 = _dafny.Map.Empty.slice().updateUnsafe(_3128_v21,!_dafny.areEqual(_3130_v22, _3130_v22));
      let _3132_v24;
      _3132_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _3133_v25;
      _3133_v25 = _dafny.MultiSet.fromElements(_3132_v24, _3132_v24);
      _3131_v23 = (_3131_v23).update(_3128_v21, ((_this).f15).isEqualTo(new BigNumber((_3133_v25).cardinality())));
      let _3134_v26;
      _3134_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber(961));
      let _3135_v27;
      let _nw586 = new _module.C1();
      _nw586.__ctor(new BigNumber(774));
      _3135_v27 = _nw586;
      let _3136_v28;
      _3136_v28 = _dafny.Seq.of(_3135_v27);
      let _3137_v29;
      _3137_v29 = new _dafny.CodePoint('j'.codePointAt(0));
      let _3138_v30;
      _3138_v30 = _module.D16.create_DC45(new BigNumber((_3136_v28).length), _3137_v29, (_this).f15, p0);
      _3134_v26 = (_3134_v26).update((_this).f15, _module.__default.fm1((_3138_v30).dtor_cf81, globalState));
      _3137_v29 = _3137_v29;
      let _3139_v31;
      let _nw587 = new _module.C5();
      _nw587.__ctor((_3135_v27).f24);
      _3139_v31 = _nw587;
      _3139_v31 = _3139_v31;
      r0 = true;
      r1 = !(_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("jayoxhlrs"), _3137_v29));
      return [r0, r1];
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
