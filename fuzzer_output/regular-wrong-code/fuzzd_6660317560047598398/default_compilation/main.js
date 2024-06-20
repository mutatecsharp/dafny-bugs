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
      return !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(916)), function (_0_i0) {
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(true));
      }), (_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false, false, false, false))));
    };
    static fm1(p0, p1, globalState) {
      return new _dafny.CodePoint('m'.codePointAt(0));
    };
    static fm2(globalState) {
      let _source0 = ((!(!(!(false)))) ? (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(709)), function (_1_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(-390))).length))).length));
      }), _dafny.Seq.of(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (function () {
          let _coll1 = new _dafny.Set();
          for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Seq.of(false)),true)).Keys.Elements) {
            let _2_v1 = _compr_1;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Seq.of(false)),true)).contains(_2_v1)) {
              _coll1.add(_2_v1);
            }
          }
          return _coll1;
        }()).Elements) {
          let _3_v0 = _compr_0;
          if ((function () {
            let _coll2 = new _dafny.Set();
            for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Seq.of(false)),true)).Keys.Elements) {
              let _4_v1 = _compr_2;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Seq.of(false)),true)).contains(_4_v1)) {
                _coll2.add(_4_v1);
              }
            }
            return _coll2;
          }()).contains(_3_v0)) {
            _coll0.push([_3_v0,true]);
          }
        }
        return _coll0;
      }()).length), new BigNumber(120)), _dafny.Seq.of(new BigNumber(796)), _dafny.Seq.of(new BigNumber(126)), _dafny.Seq.of(new BigNumber(285), new BigNumber((_dafny.Seq.UnicodeFromString("ngfivc")).length)))) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(-535), new BigNumber(-862)), _dafny.Seq.of(new BigNumber(569))))));
      let _5___mcc_h0 = _source0;
      let _6_cf88 = _5___mcc_h0;
      return _module.D0.create_DC1();
    };
    static fm3(p0, p1, globalState) {
      return (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((new BigNumber(-560)).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-401))).length))).length)))).cardinality()));
    };
    static fm4(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(710)), function (_7_i0) {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)))).length), new BigNumber(-29), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,false))).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-691),new BigNumber((_dafny.MultiSet.fromElements(true, !(true))).cardinality()))).length), new BigNumber(-9));
      }), ((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(314)), function (_8_i1) {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true, false, false, true))).length));
      })) : (_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(967),new BigNumber(448))).length),new BigNumber((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(366), new BigNumber(228))) {
          let _9_v0 = _compr_3;
          if (((new BigNumber(366)).isLessThanOrEqualTo(_9_v0)) && ((_9_v0).isLessThan(new BigNumber(228)))) {
            _coll3.push([(_9_v0).minus(new BigNumber(((_module.D16.create_DC33(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(734))).length), _dafny.MultiSet.fromElements(false), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("i")).length)), _dafny.Set.fromElements(true))).dtor_cf54).cardinality())),new _dafny.CodePoint('y'.codePointAt(0))]);
          }
        }
        return _coll3;
      }()).length))).length))))));
    };
    static fm5(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(603))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(157)))).Merge((function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Set.fromElements(_module.D0.create_DC1())).Elements) {
          let _10_v0 = _compr_4;
          if ((_dafny.Set.fromElements(_module.D0.create_DC1())).contains(_10_v0)) {
            _coll4.push([_10_v0,new BigNumber(558)]);
          }
        }
        return _coll4;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)))).cardinality()))));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC4(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-606), new BigNumber((_dafny.Seq.of(false)).length)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(167))).length), new BigNumber(841), new BigNumber(909))).length))).length), (_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(464),(_module.D21.create_DC43(true, false, _dafny.Seq.of(!(false)))).dtor_cf70)).length)).multipliedBy(new BigNumber(159))));
    };
    static fm16(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eh"), _dafny.Seq.UnicodeFromString("rgorcc")), _dafny.Seq.UnicodeFromString("oxshrry"));
    };
    static fm18(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))).length), new BigNumber(28)),(_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(false)));
    };
    static fm19(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false, false, false, false));
    };
    static fm20(p0, p1, globalState) {
      let _source1 = _module.D6.create_DC15(_dafny.MultiSet.fromElements(_dafny.Seq.of(false)));
      if (_source1.is_DC16) {
        let _11___mcc_h0 = (_source1).cf20;
        let _12___mcc_h1 = (_source1).cf21;
        let _13___mcc_h2 = (_source1).cf22;
        let _14_cf22 = _13___mcc_h2;
        let _15_cf21 = _12___mcc_h1;
        let _16_cf20 = _11___mcc_h0;
        return _module.D3.create_DC8(_dafny.Seq.of(false, _16_cf20));
      } else {
        let _17___mcc_h3 = (_source1).cf19;
        let _18_cf19 = _17___mcc_h3;
        if (false) {
          return _module.D3.create_DC8(_dafny.Seq.of(true));
        } else {
          return _module.D3.create_DC8(_dafny.Seq.of(!(false), true));
        }
      }
    };
    static fm21(p0, p1, p2, p3, globalState) {
      let _source2 = _module.D19.create_DC38(false, !(true));
      if (_source2.is_DC38) {
        let _19___mcc_h0 = (_source2).cf63;
        let _20___mcc_h1 = (_source2).cf64;
        let _21_cf64 = _20___mcc_h1;
        let _22_cf63 = _19___mcc_h0;
        return _dafny.MultiSet.fromElements(_21_cf64, _21_cf64, _22_cf63, _21_cf64);
      } else if (_source2.is_DC37) {
        let _23___mcc_h2 = (_source2).cf62;
        let _24_cf62 = _23___mcc_h2;
        return (_module.D16.create_DC33(true, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(250))).length))).length), _dafny.MultiSet.fromElements(true), new BigNumber(527), _dafny.Set.fromElements(!(true)))).dtor_cf54;
      } else {
        let _25___mcc_h3 = (_source2).cf65;
        let _26_cf65 = _25___mcc_h3;
        return (_dafny.MultiSet.fromElements(false, true)).Union(_dafny.MultiSet.fromElements(true, true));
      }
    };
    static fm22(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())),_dafny.Seq.UnicodeFromString("n"))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(855), new BigNumber((_dafny.Seq.UnicodeFromString("xt")).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(290)), function (_27_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))),_dafny.Seq.UnicodeFromString("pmrrdi"))));
    };
    static fm23(p0, p1, globalState) {
      return _dafny.Set.fromElements(false);
    };
    static fm24(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Seq.of(!(!(true))), _dafny.Seq.of(true))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.of(false)))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(true)));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC1(), _module.D0.create_DC1()), _dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC1(), _module.D0.create_DC1(), _module.D0.create_DC1(), _module.D0.create_DC1()), _dafny.Seq.of(_module.D0.create_DC1())));
    };
    static fm27(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.of(false))).Union((_dafny.MultiSet.fromElements(_dafny.Seq.of(false, !(true), !(true), false), _dafny.Seq.of(false, true, true))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(true))));
    };
    static fm28(globalState) {
      if ((new BigNumber(958)).isEqualTo(new BigNumber(397))) {
        return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-268));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(388));
      }
    };
    static fm29(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(452),new BigNumber(863))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Set.fromElements(_dafny.Seq.of(true, false), _dafny.Seq.of(true, true, false, !(true), !(true)), _dafny.Seq.of(true, true, false))).Elements) {
          let _28_v0 = _compr_5;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(true, false), _dafny.Seq.of(true, true, false, !(true), !(true)), _dafny.Seq.of(true, true, false))).contains(_28_v0)) {
            _coll5.push([_28_v0,!(true)]);
          }
        }
        return _coll5;
      }()).length),new BigNumber(-354)));
    };
    static fm30(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(812)), _dafny.Seq.of(new BigNumber(-968), new BigNumber((_dafny.Seq.of(false, true)).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(573)), function (_29_i0) {
        return _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber((_dafny.Seq.of(true, true)).length), new BigNumber((function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of _dafny.IntegerRange(new BigNumber(122), new BigNumber(-15))) {
            let _30_v0 = _compr_6;
            if (((new BigNumber(122)).isLessThanOrEqualTo(_30_v0)) && ((_30_v0).isLessThan(new BigNumber(-15)))) {
              _coll6.push([(_30_v0).multipliedBy(new BigNumber(-977)),true]);
            }
          }
          return _coll6;
        }()).length));
      }));
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(-505), new BigNumber((_dafny.Set.fromElements(new BigNumber(-506))).length)));
    };
    static fm32(p0, p1, globalState) {
      return (_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.of(false)).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.UnicodeFromString("vyojj"))).length))).length))).Intersect(_dafny.Set.fromElements(new BigNumber(864)));
    };
    static fm33(p0, p1, p2, globalState) {
      return _module.D6.create_DC15((_dafny.MultiSet.fromElements(_dafny.Seq.of(true, true))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.of(false, true))));
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return ((function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber(236)))).Elements) {
          let _31_v0 = _compr_7;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber(236))), _31_v0)) {
            _coll7.push([_31_v0,_dafny.Seq.UnicodeFromString("e")]);
          }
        }
        return _coll7;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(299), new BigNumber(935), new BigNumber(-557), new BigNumber((_dafny.Seq.UnicodeFromString("chaced")).length), new BigNumber(934))).length), new BigNumber(375))),_dafny.Seq.UnicodeFromString("jnbkik")))).Merge(((false) ? (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(343)), function (_32_i0) {
        return new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality());
      })),_dafny.Seq.UnicodeFromString("wdpr"))) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(25)), function (_33_i1) {
        return new BigNumber((_dafny.Seq.of(true)).length);
      }),new BigNumber(-972))).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(331)), function (_34_i2) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })))));
    };
    static fm35(p0, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("tghobrq")).length))).Intersect(_dafny.Set.fromElements(new BigNumber(-815)))).Intersect((_dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(false, true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("d")).length))).length))), new BigNumber(186))).Union(_dafny.Set.fromElements(new BigNumber(858), new BigNumber(-526), (_module.D11.create_DC24(!(true), true, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("tfoksstqw")).length)))).dtor_cf38, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(653)), function (_35_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length))));
    };
    static fm36(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(35)), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(!(!(false)), true)).cardinality()))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(918)), function (_36_i0) {
        return new BigNumber(903);
      }));
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.UnicodeFromString("gcdnnimh");
    };
    static fm38(p0, p1, p2, p3, globalState) {
      if ((new BigNumber(472)).isLessThan(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ky")).length))).length))) {
        return _module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ys")).length))).length)));
      } else if (!(true)) {
        return _module.D0.create_DC0(new BigNumber(621));
      } else {
        return _module.D0.create_DC0(new BigNumber(820));
      }
    };
    static fm39(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("tjnq"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(462)), function (_37_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),new _dafny.CodePoint('d'.codePointAt(0)))).length)));
    };
    static fm40(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true)), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false)));
    };
    static fm41(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))).length)))).length))).length), new BigNumber((_dafny.Seq.of(false, true)).length), new BigNumber(-193))).Difference(function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(763), new BigNumber(69))) {
          let _38_v0 = _compr_8;
          if (((new BigNumber(763)).isLessThanOrEqualTo(_38_v0)) && ((_38_v0).isLessThan(new BigNumber(69)))) {
            _coll8.add((_38_v0).minus(new BigNumber(256)));
          }
        }
        return _coll8;
      }())).Intersect((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of _dafny.IntegerRange(new BigNumber(288), new BigNumber(859))) {
            let _39_v1 = _compr_10;
            if (((new BigNumber(288)).isLessThanOrEqualTo(_39_v1)) && ((_39_v1).isLessThan(new BigNumber(859)))) {
              _coll10.push([(_39_v1).plus(new BigNumber(596)),new BigNumber((_dafny.Set.fromElements(false, false, true)).length)]);
            }
          }
          return _coll10;
        }()).Keys.Elements) {
          let _40_v2 = _compr_9;
          if ((function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of _dafny.IntegerRange(new BigNumber(288), new BigNumber(859))) {
              let _39_v1 = _compr_11;
              if (((new BigNumber(288)).isLessThanOrEqualTo(_39_v1)) && ((_39_v1).isLessThan(new BigNumber(859)))) {
                _coll11.push([(_39_v1).plus(new BigNumber(596)),new BigNumber((_dafny.Set.fromElements(false, false, true)).length)]);
              }
            }
            return _coll11;
          }()).contains(_40_v2)) {
            _coll9.add((_40_v2).plus(new BigNumber(-500)));
          }
        }
        return _coll9;
      }()).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true, true, false)).length))));
    };
    static fm42(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(459),(_dafny.ZERO).minus(new BigNumber(-980)))).Merge(function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of _dafny.IntegerRange(new BigNumber(581), new BigNumber(197))) {
          let _41_v0 = _compr_12;
          if (((new BigNumber(581)).isLessThanOrEqualTo(_41_v0)) && ((_41_v0).isLessThan(new BigNumber(197)))) {
            _coll12.push([_module.__default.safeDivisionInt(_41_v0, new BigNumber(501)),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(!(!(!(!(false))))), false))).cardinality())]);
          }
        }
        return _coll12;
      }());
    };
    static fm43(globalState) {
      return _module.D6.create_DC16(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("a"), _dafny.Seq.UnicodeFromString("exmbhmdo")), !(true) || (true), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-346)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("onku")).length))));
    };
    static fm44(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(445));
    };
    static fm45(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(true, false), _dafny.MultiSet.fromElements(true, false, false, true, true), _dafny.MultiSet.fromElements(false)), ((!(true)) ? (_dafny.Seq.of(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(true))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(824)), function (_42_i0) {
        return _dafny.MultiSet.fromElements(!(false));
      }))));
    };
    static fm46(p0, globalState) {
      if (true) {
        return _dafny.Set.fromElements(false);
      } else {
        return (_module.D16.create_DC33(false, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length))).length), _dafny.MultiSet.fromElements(false, true), new BigNumber(873), _dafny.Set.fromElements(!(true), true))).dtor_cf56;
      }
    };
    static fm47(p0, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(true), ((true) ? (_dafny.Seq.of(false, true)) : (_dafny.Seq.of(!(true))))));
    };
    static fm49(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(191)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(737), new BigNumber(-741))).length), new BigNumber(344))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-224)), function (_43_i0) {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(606),(_module.D6.create_DC16(true, false, _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(576),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).length)))).dtor_cf20)).length))).length));
      })), _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, true)).length))));
    };
    static fm50(p0, p1, p2, p3, globalState) {
      let _source3 = _module.D1.create_DC3(!(!(true)));
      if (_source3.is_DC4) {
        let _44___mcc_h0 = (_source3).cf3;
        let _45___mcc_h1 = (_source3).cf4;
        let _46_cf4 = _45___mcc_h1;
        let _47_cf3 = _44___mcc_h0;
        return _dafny.Seq.UnicodeFromString("dyg");
      } else {
        let _48___mcc_h2 = (_source3).cf2;
        let _49_cf2 = _48___mcc_h2;
        return _dafny.Seq.UnicodeFromString("dstoi");
      }
    };
    static fm51(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(372))).length),false));
    };
    static fm52(p0, globalState) {
      return _module.D3.create_DC8(((false) ? (_dafny.Seq.of((_module.D16.create_DC33(false, new BigNumber(-609), _dafny.MultiSet.FromArray(_dafny.Seq.of(true, true, false, false)), new BigNumber(-454), _dafny.Set.fromElements(false))).dtor_cf52)) : (_dafny.Seq.of(false))));
    };
    static fm53(p0, p1, globalState) {
      return (_dafny.Set.fromElements(true, false)).Difference((_dafny.Set.fromElements(false, false)).Intersect(_dafny.Set.fromElements(true)));
    };
    static fm54(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,false)), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,true)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(155)), function (_50_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(true,!(false));
      })));
    };
    static fm55(p0, p1, p2, globalState) {
      let _source4 = _module.D21.create_DC43(false, true, _dafny.Seq.of(!(!(true))));
      if (_source4.is_DC43) {
        let _51___mcc_h0 = (_source4).cf69;
        let _52___mcc_h1 = (_source4).cf70;
        let _53___mcc_h2 = (_source4).cf71;
        let _54_cf71 = _53___mcc_h2;
        let _55_cf70 = _52___mcc_h1;
        let _56_cf69 = _51___mcc_h0;
        return _module.D1.create_DC3(false);
      } else if (_source4.is_DC44) {
        let _57___mcc_h3 = (_source4).cf72;
        let _58_cf72 = _57___mcc_h3;
        return _module.D1.create_DC3(false);
      } else {
        let _59___mcc_h4 = (_source4).cf68;
        let _60_cf68 = _59___mcc_h4;
        return _module.D1.create_DC3(true);
      }
    };
    static fm56(p0, p1, p2, globalState) {
      if ((false) && (!(true))) {
        return _dafny.Seq.of(false);
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(!(false)));
      }
    };
    static fm57(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(189)), function (_61_i0) {
        return new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(681)))).length);
      }), _dafny.Seq.of(new BigNumber(-469))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-661)), function (_62_i1) {
        return new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),false), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),false))).length);
      }));
    };
    static fm58(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D12.create_DC27(_dafny.Seq.of(_dafny.Seq.of(true))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D12.create_DC27(_dafny.Seq.of(_dafny.Seq.of(true, false)))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D12.create_DC27(_dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(true)))));
    };
    static fm59(globalState) {
      let _source5 = _module.D21.create_DC43(true, false, _dafny.Seq.of(false, false));
      if (_source5.is_DC43) {
        let _63___mcc_h0 = (_source5).cf69;
        let _64___mcc_h1 = (_source5).cf70;
        let _65___mcc_h2 = (_source5).cf71;
        let _66_cf71 = _65___mcc_h2;
        let _67_cf70 = _64___mcc_h1;
        let _68_cf69 = _63___mcc_h0;
        return function () {
          let _coll13 = new _dafny.Map();
          for (const _compr_13 of (_dafny.Seq.of(_module.D10.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_68_cf69,new BigNumber(804))))).Elements) {
            let _69_v0 = _compr_13;
            if (_dafny.Seq.contains(_dafny.Seq.of(_module.D10.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_68_cf69,new BigNumber(804)))), _69_v0)) {
              _coll13.push([_69_v0,new BigNumber((_dafny.Seq.UnicodeFromString("sg")).length)]);
            }
          }
          return _coll13;
        }();
      } else if (_source5.is_DC44) {
        let _70___mcc_h3 = (_source5).cf72;
        let _71_cf72 = _70___mcc_h3;
        return _dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(_71_cf72,new BigNumber(18))),(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("c")).length))));
      } else {
        let _72___mcc_h4 = (_source5).cf68;
        let _73_cf68 = _72___mcc_h4;
        return (_dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(2))).length))),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(776))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(335))),new BigNumber((_dafny.Set.fromElements(new BigNumber(377), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(762), new BigNumber(-921))).cardinality()))))).length)));
      }
    };
    static fm60(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(false, false), _dafny.Seq.of(false)),(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(909),new BigNumber(368))).length),!(!(!(false))))).Merge(function () {
        let _coll14 = new _dafny.Map();
        for (const _compr_14 of (_dafny.MultiSet.fromElements(new BigNumber(896))).Elements) {
          let _74_v0 = _compr_14;
          if ((_dafny.MultiSet.fromElements(new BigNumber(896))).contains(_74_v0)) {
            _coll14.push([(_74_v0).minus(new BigNumber(75)),true]);
          }
        }
        return _coll14;
      }()));
    };
    static fm61(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(true, !(true), false, true, false)), _dafny.Seq.of(_dafny.Seq.of(!(!(false)), true, true, false, !(false))));
    };
    static fm62(globalState) {
      return _module.D4.create_DC12(((!(true)) ? (new _dafny.CodePoint('p'.codePointAt(0))) : (new _dafny.CodePoint('a'.codePointAt(0)))));
    };
    static fm63(globalState) {
      return _module.D4.create_DC13(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.UnicodeFromString("hxnhw")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("aog")).length)))).length), (false) === (false), new BigNumber(850), _dafny.Seq.of(!(false)));
    };
    static fm64(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(new _dafny.CodePoint('e'.codePointAt(0)))).Union(_dafny.Set.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))))).Union((function () {
        let _coll15 = new _dafny.Set();
        for (const _compr_15 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0)))).Elements) {
          let _75_v0 = _compr_15;
          if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0)))).contains(_75_v0)) {
            _coll15.add(_75_v0);
          }
        }
        return _coll15;
      }()).Difference(_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)))));
    };
    static fm65(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('f'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('r'.codePointAt(0))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('q'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),new _dafny.CodePoint('j'.codePointAt(0)))));
    };
    static fm66(p0, globalState) {
      return _module.D16.create_DC33(!(true) || (false), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("gk")).length), new BigNumber(925)), ((_module.D16.create_DC33(false, new BigNumber((_dafny.Seq.of(false)).length), _dafny.MultiSet.fromElements(!(false)), new BigNumber(249), _dafny.Set.fromElements(false))).dtor_cf54).Intersect(_dafny.MultiSet.fromElements(true)), new BigNumber(-152), (_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(false)));
    };
    static fm67(p0, globalState) {
      if (true) {
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(876),new BigNumber(861)),!(!(true)));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(720), new BigNumber(548))).length),new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("rss"), _dafny.Seq.UnicodeFromString("s"))).cardinality())),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("udygyb")).length),new BigNumber(-885)),!(true)));
      }
    };
    static fm68(globalState) {
      return _module.D12.create_DC27(((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(223)), function (_76_i0) {
  return _dafny.Seq.of(true);
})) : (_dafny.Seq.of(_dafny.Seq.of(false)))));
    };
    static fm69(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-3)))).Union(_dafny.MultiSet.fromElements(new BigNumber(-982)))).cardinality()),_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("igvytirrh"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(614)), function (_77_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })));
    };
    static fm70(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), function (_78_i0) {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0)));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(866)), function (_79_i1) {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)));
      })), _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))), _dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0))), _dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)))));
    };
    static fm71(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),false));
    };
    static fm72(p0, p1, globalState) {
      let _source6 = _dafny.Set.fromElements(new BigNumber(350), new BigNumber((_dafny.Seq.of(true, !(true), true)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(846)), function (_80_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length));
      let _81___mcc_h0 = _source6;
      let _82_cf23 = _81___mcc_h0;
      return _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(510)), function (_83_i1) {
        return _dafny.Seq.of(new BigNumber(180));
      }));
    };
    static fm73(globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(544),new BigNumber((_dafny.Seq.UnicodeFromString("qkivckp")).length))).Keys.Elements) {
          let _84_v0 = _compr_16;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(544),new BigNumber((_dafny.Seq.UnicodeFromString("qkivckp")).length))).contains(_84_v0)) {
            _coll16.push([_module.__default.safeModuloInt(_84_v0, new BigNumber(688)),false]);
          }
        }
        return _coll16;
      }()).length)), _dafny.Seq.of(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,true),!(false)))).length)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(732)))))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(519)), function (_85_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(602),new BigNumber(-700))).length);
      })));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = false;
      let r1 = _dafny.ZERO;
      if (true) {
        let _86_v0;
        let _nw0 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
        _86_v0 = _nw0;
        let _87_v1;
        _87_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,_86_v0);
        _86_v0 = (((_87_v1).contains(p3)) ? ((_87_v1).get(p3)) : (_86_v0));
        let _88_v3;
        _88_v3 = new BigNumber(50);
        let _89_v4;
        _89_v4 = _dafny.Seq.of(_88_v3, _88_v3, _88_v3);
        let _90_v5;
        _90_v5 = _dafny.Set.fromElements(_88_v3);
        let _91_v6;
        _91_v6 = _dafny.MultiSet.fromElements(p3);
        let _92_v7;
        _92_v7 = _dafny.Set.fromElements(p3, p1, p3, p1, p3);
        let _93_v8;
        _93_v8 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
        let _94_v9;
        _94_v9 = _dafny.Seq.of(_93_v8, _dafny.Map.Empty.slice().updateUnsafe(p1,p3), _dafny.Map.Empty.slice().updateUnsafe(p1,p1), _93_v8);
        let _95_v10;
        let _nw1 = Array((new BigNumber(19)).toNumber());
        _nw1[(_dafny.ZERO).toNumber()] = p3;
        _nw1[(_dafny.ONE).toNumber()] = p3;
        _nw1[(new BigNumber(2)).toNumber()] = p1;
        _nw1[(new BigNumber(3)).toNumber()] = !(new BigNumber((function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of (_89_v4).Elements) {
            let _96_v2 = _compr_17;
            if (_dafny.Seq.contains(_89_v4, _96_v2)) {
              _coll17.push([_module.__default.safeDivisionInt(_96_v2, new BigNumber(88)),(_dafny.ZERO).minus(_88_v3)]);
            }
          }
          return _coll17;
        }()).length)).isEqualTo(new BigNumber(803));
        _nw1[(new BigNumber(4)).toNumber()] = p1;
        _nw1[(new BigNumber(5)).toNumber()] = p1;
        _nw1[(new BigNumber(6)).toNumber()] = (_90_v5).IsProperSubsetOf(_90_v5);
        _nw1[(new BigNumber(7)).toNumber()] = p1;
        _nw1[(new BigNumber(8)).toNumber()] = p1;
        _nw1[(new BigNumber(9)).toNumber()] = (_88_v3).isLessThanOrEqualTo(_88_v3);
        _nw1[(new BigNumber(10)).toNumber()] = (_91_v6).IsDisjointFrom(_91_v6);
        _nw1[(new BigNumber(11)).toNumber()] = (_92_v7).IsProperSubsetOf(_module.__default.fm53(_88_v3, _94_v9, globalState));
        _nw1[(new BigNumber(12)).toNumber()] = p1;
        _nw1[(new BigNumber(13)).toNumber()] = p1;
        _nw1[(new BigNumber(14)).toNumber()] = p1;
        _nw1[(new BigNumber(15)).toNumber()] = p1;
        _nw1[(new BigNumber(16)).toNumber()] = !(p1) || (p3);
        _nw1[(new BigNumber(17)).toNumber()] = p3;
        _nw1[(new BigNumber(18)).toNumber()] = p3;
        _95_v10 = _nw1;
        let _97_v11;
        _97_v11 = _dafny.Seq.of(true);
        let _index0 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_95_v10).length));
        (_95_v10)[_index0] = (_97_v11)[_module.__default.safeIndex(_88_v3, new BigNumber((_97_v11).length))];
        let _98_v12;
        _98_v12 = _dafny.Seq.UnicodeFromString("fku");
        let _index1 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_95_v10).length));
        (_95_v10)[_index1] = (_90_v5).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_98_v12).length), _88_v3));
        let _99_v13;
        let _nw2 = Array((new BigNumber(10)).toNumber()).fill([]);
        _99_v13 = _nw2;
        let _100_v14;
        let _nw3 = new _module.C4();
        _nw3.__ctor(_99_v13, (_dafny.ZERO).minus((_dafny.ZERO).minus(_88_v3)), (_95_v10)[_module.__default.safeIndex(new BigNumber(167), new BigNumber((_95_v10).length))]);
        _100_v14 = _nw3;
        _100_v14 = _100_v14;
        let _101_v15;
        let _nw4 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _101_v15 = _nw4;
        let _index2 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_101_v15).length));
        (_101_v15)[_index2] = (new BigNumber((_91_v6).cardinality())).minus(_88_v3);
        let _index3 = _module.__default.safeIndex(new BigNumber(173), new BigNumber((_101_v15).length));
        (_101_v15)[_index3] = (((((p2).contains(new BigNumber((_92_v7).length))) ? ((p2).get(new BigNumber((_92_v7).length))) : (new BigNumber((_98_v12).length)))).plus(_88_v3)).plus(_88_v3);
        let _102_v16;
        let _nw5 = new _module.C2();
        _nw5.__ctor(_98_v12, (_97_v11)[_module.__default.safeIndex((_dafny.ZERO).minus((_101_v15)[_module.__default.safeIndex(new BigNumber(173), new BigNumber((_101_v15).length))]), new BigNumber((_97_v11).length))]);
        _102_v16 = _nw5;
        let _103_v17;
        _103_v17 = _module.D27.create_DC56(_102_v16);
        _102_v16 = (_103_v17).dtor_cf95;
      } else {
        let _104_v18;
        _104_v18 = new _dafny.CodePoint('r'.codePointAt(0));
        let _105_v19;
        let _nw6 = Array((new BigNumber(8)).toNumber());
        _nw6[(_dafny.ZERO).toNumber()] = p1;
        _nw6[(_dafny.ONE).toNumber()] = _module.__default.fm0(_104_v18, globalState);
        _nw6[(new BigNumber(2)).toNumber()] = p3;
        _nw6[(new BigNumber(3)).toNumber()] = p3;
        _nw6[(new BigNumber(4)).toNumber()] = p1;
        _nw6[(new BigNumber(5)).toNumber()] = false;
        _nw6[(new BigNumber(6)).toNumber()] = p3;
        _nw6[(new BigNumber(7)).toNumber()] = p1;
        _105_v19 = _nw6;
        let _106_v20;
        _106_v20 = _dafny.Map.Empty.slice().updateUnsafe(p3,_105_v19);
        let _107_v21;
        let _nw7 = Array((new BigNumber(22)).toNumber());
        _nw7[(_dafny.ZERO).toNumber()] = _105_v19;
        _nw7[(_dafny.ONE).toNumber()] = _105_v19;
        _nw7[(new BigNumber(2)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(3)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(4)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(5)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(6)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(7)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(8)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(9)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(10)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(11)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(12)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(13)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(14)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(15)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(16)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(17)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(18)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(19)).toNumber()] = _105_v19;
        _nw7[(new BigNumber(20)).toNumber()] = (((_106_v20).contains(p3)) ? ((_106_v20).get(p3)) : (_105_v19));
        _nw7[(new BigNumber(21)).toNumber()] = _105_v19;
        _107_v21 = _nw7;
        let _108_v22;
        _108_v22 = new BigNumber(852);
        let _109_v23;
        let _nw8 = new _module.C4();
        _nw8.__ctor(_107_v21, _108_v22, p1);
        _109_v23 = _nw8;
        let _110_v24;
        _110_v24 = _dafny.Set.fromElements(_109_v23);
        let _111_v25;
        _111_v25 = _module.D12.create_DC28(new BigNumber((_110_v24).length), _108_v22, (_109_v23).f4, (_109_v23).f4, _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("tbi")).length), (_109_v23).f6));
        _111_v25 = _111_v25;
        let _112_v26;
        _112_v26 = _dafny.Seq.UnicodeFromString("filnlnc");
        _112_v26 = _112_v26;
        let _113_v27;
        _113_v27 = _dafny.Seq.of((_109_v23).f6, _108_v22);
        let _114_v28;
        _114_v28 = _dafny.MultiSet.fromElements(p3);
        let _115_v29;
        let _nw9 = Array((new BigNumber(22)).toNumber());
        _nw9[(_dafny.ZERO).toNumber()] = _113_v27;
        _nw9[(_dafny.ONE).toNumber()] = _113_v27;
        _nw9[(new BigNumber(2)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(new BigNumber((_112_v26).length));
        _nw9[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_113_v27, _113_v27);
        _nw9[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_113_v27, _module.__default.safeIndex(new BigNumber((_112_v26).length), new BigNumber((_113_v27).length)), (_109_v23).f6);
        _nw9[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_113_v27, _113_v27);
        _nw9[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_113_v27, _113_v27);
        _nw9[(new BigNumber(8)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(9)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(10)).toNumber()] = _dafny.Seq.of((((_114_v28).contains(p1)) ? ((_114_v28).get(p1)) : (_108_v22)));
        _nw9[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_113_v27, _dafny.Seq.of((_113_v27)[_module.__default.safeIndex((_109_v23).f6, new BigNumber((_113_v27).length))], new BigNumber((_112_v26).length), _108_v22, (((_114_v28).contains(false)) ? ((_114_v28).get(false)) : (new BigNumber((p2).cardinality()))), _108_v22));
        _nw9[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(_108_v22, new BigNumber((_112_v26).length));
        _nw9[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_109_v23).f6), _113_v27);
        _nw9[(new BigNumber(14)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(15)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(16)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(17)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(18)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(19)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(20)).toNumber()] = _113_v27;
        _nw9[(new BigNumber(21)).toNumber()] = _module.__default.fm57((_109_v23).f6, (_109_v23).f6, globalState);
        _115_v29 = _nw9;
        let _116_v30;
        _116_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,_108_v22);
        let _117_v31;
        _117_v31 = _module.D10.create_DC21((_116_v30).update(p3, (_109_v23).f6));
        let _118_v32;
        _118_v32 = _dafny.Seq.of(_117_v31);
        let _119_v33;
        _119_v33 = _dafny.Seq.of(_module.__default.fm0(_104_v18, globalState), p3);
        let _rhs0 = !((_119_v33)[_module.__default.safeIndex(_108_v22, new BigNumber((_119_v33).length))]);
        let _rhs1 = _115_v29;
        let _rhs2 = ((false) ? (_118_v32) : (_dafny.Seq.update(_118_v32, _module.__default.safeIndex((_109_v23).f6, new BigNumber((_118_v32).length)), _117_v31)));
        let _rhs3 = (_109_v23).fm7(globalState);
        r0 = _rhs0;
        _115_v29 = _rhs1;
        _118_v32 = _rhs2;
        r1 = _rhs3;
        let _120_v34;
        _120_v34 = _module.D16.create_DC33((_109_v23).f4, (_109_v23).fm9(p3, p3, (_109_v23).f6, globalState), _114_v28, new BigNumber((_112_v26).length), _dafny.Set.fromElements((_109_v23).f4, p1));
        r0 = !(!(((_120_v34).dtor_cf54).IsSubsetOf(_114_v28)) || ((_120_v34).dtor_cf52));
        let _121_v35;
        _121_v35 = _dafny.Map.Empty.slice().updateUnsafe(_108_v22,p3);
        let _122_v36;
        _122_v36 = _dafny.Map.Empty.slice().updateUnsafe(_104_v18,_119_v33);
        let _123_v37;
        let _nw10 = new _module.C1();
        _nw10.__ctor((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(185), (((_114_v28).contains((((_121_v35).contains(new BigNumber(884))) ? ((_121_v35).get(new BigNumber(884))) : (p1)))) ? ((_114_v28).get((((_121_v35).contains(new BigNumber(884))) ? ((_121_v35).get(new BigNumber(884))) : (p1)))) : (_108_v22)))), _108_v22, ((p1) ? (_107_v21) : (_109_v23.f5)), new BigNumber((_122_v36).length), (p2).IsProperSubsetOf(p2));
        _123_v37 = _nw10;
        let _rhs4 = _123_v37;
        let _rhs5 = _dafny.Seq.Concat(_113_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(476)), ((_124_v23) => function (_125_i0) {
          return (_dafny.ZERO).minus((_124_v23).f6);
        })(_109_v23)));
        let _rhs6 = _105_v19;
        _123_v37 = _rhs4;
        _113_v27 = _rhs5;
        _105_v19 = _rhs6;
      }
      let _126_v38;
      _126_v38 = new BigNumber(197);
      let _hi0 = _module.__default.fm3(_126_v38, true, globalState);
      for (let _127_i1 = _126_v38; _127_i1.isLessThan(_hi0); _127_i1 = _127_i1.plus(_dafny.ONE)) {
        _126_v38 = _module.__default.safeModuloInt(_126_v38, _126_v38);
        let _128_v39;
        _128_v39 = _dafny.Seq.of(_126_v38);
        let _129_v40;
        _129_v40 = _dafny.Seq.UnicodeFromString("etegojdr");
        r0 = ((false) ? (p3) : (!_dafny.areEqual(_128_v39, _dafny.Seq.update(_128_v39, _module.__default.safeIndex(new BigNumber((_129_v40).length), new BigNumber((_128_v39).length)), _127_i1))));
        let _130_v41;
        let _nw11 = new _module.C2();
        _nw11.__ctor(_dafny.Seq.Concat(_129_v40, _129_v40), p3);
        _130_v41 = _nw11;
        _130_v41 = ((p1) ? (_130_v41) : (_130_v41));
        let _131_v42;
        _131_v42 = _dafny.Seq.of(!(p1));
        let _132_v43;
        _132_v43 = _dafny.Set.fromElements(_131_v42, _131_v42, _131_v42, _131_v42);
        let _133_v44;
        _133_v44 = _dafny.Seq.of(_131_v42, _131_v42);
        let _134_v46;
        let _nw12 = Array((new BigNumber(2)).toNumber()).fill([]);
        _134_v46 = _nw12;
        let _135_v47;
        let _nw13 = new _module.C9();
        _nw13.__ctor(_127_i1, (_132_v43).Difference(function () {
          let _coll18 = new _dafny.Set();
          for (const _compr_18 of (_133_v44).Elements) {
            let _136_v45 = _compr_18;
            if (_dafny.Seq.contains(_133_v44, _136_v45)) {
              _coll18.add(_136_v45);
            }
          }
          return _coll18;
        }()), _134_v46, _127_i1, p3);
        _135_v47 = _nw13;
      }
      let _137_v48;
      let _init0 = function (_138_i2) {
        return (_138_i2).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new BigNumber(33))).length));
      };
      let _nw14 = Array((new BigNumber(4)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw14.length); _i0_0++) {
        _nw14[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _137_v48 = _nw14;
      let _139_v49;
      _139_v49 = _dafny.Seq.of(_137_v48, _137_v48, _137_v48, _137_v48);
      _139_v49 = _dafny.Seq.Concat(_139_v49, _139_v49);
      _126_v38 = _module.__default.safeDivisionInt(_module.__default.fm3(_126_v38, p1, globalState), _126_v38);
      if (p1) {
        let _index4 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length));
        (_137_v48)[_index4] = new BigNumber(946);
        let _index5 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length));
        (_137_v48)[_index5] = new BigNumber(-463);
        let _140_v50;
        let _init1 = function (_141_i3) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        };
        let _nw15 = Array((new BigNumber(4)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw15.length); _i0_1++) {
          _nw15[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _140_v50 = _nw15;
        let _142_v51;
        _142_v51 = new _dafny.CodePoint('j'.codePointAt(0));
        let _index6 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_140_v50).length));
        (_140_v50)[_index6] = _142_v51;
        let _index7 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_140_v50).length));
        let _rhs7 = _142_v51;
        let _rhs8 = !(p1) || (_module.__default.fm0(_142_v51, globalState));
        let _rhs9 = (_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))];
        let _rhs10 = !(p3) || (!(!(false)));
        let _lhs0 = _140_v50;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_140_v50).length));
        _lhs0[_lhs1] = _rhs7;
        r0 = _rhs8;
        r1 = _rhs9;
        r0 = _rhs10;
        let _index8 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_140_v50).length));
        (_140_v50)[_index8] = new _dafny.CodePoint('b'.codePointAt(0));
        let _143_v52;
        _143_v52 = _dafny.Seq.of(p1);
        let _144_v53;
        _144_v53 = _dafny.Seq.UnicodeFromString("bjpsw");
        let _145_v54;
        _145_v54 = _dafny.Set.fromElements(_143_v52, _dafny.Seq.update(_dafny.Seq.of(false, p3), _module.__default.safeIndex(new BigNumber((_144_v53).length), new BigNumber((_dafny.Seq.of(false, p3)).length)), p3));
        let _146_v55;
        _146_v55 = _module.D11.create_DC24(p1, p3, (_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))]);
        let _147_v56;
        _147_v56 = _module.D1.create_DC3(p1);
        let _148_v57;
        let _nw16 = Array((new BigNumber(27)).toNumber());
        _nw16[(_dafny.ZERO).toNumber()] = p3;
        _nw16[(_dafny.ONE).toNumber()] = p3;
        _nw16[(new BigNumber(2)).toNumber()] = p3;
        _nw16[(new BigNumber(3)).toNumber()] = p1;
        _nw16[(new BigNumber(4)).toNumber()] = true;
        _nw16[(new BigNumber(5)).toNumber()] = (_146_v55).dtor_cf36;
        _nw16[(new BigNumber(6)).toNumber()] = p1;
        _nw16[(new BigNumber(7)).toNumber()] = p3;
        _nw16[(new BigNumber(8)).toNumber()] = p3;
        _nw16[(new BigNumber(9)).toNumber()] = (_147_v56).dtor_cf2;
        _nw16[(new BigNumber(10)).toNumber()] = !(p3);
        _nw16[(new BigNumber(11)).toNumber()] = p3;
        _nw16[(new BigNumber(12)).toNumber()] = p1;
        _nw16[(new BigNumber(13)).toNumber()] = p3;
        _nw16[(new BigNumber(14)).toNumber()] = p3;
        _nw16[(new BigNumber(15)).toNumber()] = p3;
        _nw16[(new BigNumber(16)).toNumber()] = p3;
        _nw16[(new BigNumber(17)).toNumber()] = p1;
        _nw16[(new BigNumber(18)).toNumber()] = p1;
        _nw16[(new BigNumber(19)).toNumber()] = !(p1);
        _nw16[(new BigNumber(20)).toNumber()] = p3;
        _nw16[(new BigNumber(21)).toNumber()] = false;
        _nw16[(new BigNumber(22)).toNumber()] = p1;
        _nw16[(new BigNumber(23)).toNumber()] = p3;
        _nw16[(new BigNumber(24)).toNumber()] = !(p1);
        _nw16[(new BigNumber(25)).toNumber()] = true;
        _nw16[(new BigNumber(26)).toNumber()] = p1;
        _148_v57 = _nw16;
        let _149_v58;
        let _nw17 = Array((new BigNumber(8)).toNumber());
        _nw17[(_dafny.ZERO).toNumber()] = _148_v57;
        _nw17[(_dafny.ONE).toNumber()] = _148_v57;
        _nw17[(new BigNumber(2)).toNumber()] = _148_v57;
        _nw17[(new BigNumber(3)).toNumber()] = _148_v57;
        _nw17[(new BigNumber(4)).toNumber()] = _148_v57;
        _nw17[(new BigNumber(5)).toNumber()] = _148_v57;
        _nw17[(new BigNumber(6)).toNumber()] = _148_v57;
        _nw17[(new BigNumber(7)).toNumber()] = _148_v57;
        _149_v58 = _nw17;
        let _150_v59;
        _150_v59 = _dafny.Seq.of(_149_v58);
        let _151_v60;
        _151_v60 = _dafny.Set.fromElements(_126_v38);
        let _152_v61;
        _152_v61 = _module.D12.create_DC28(new BigNumber((_151_v60).length), _126_v38, p3, p3, new BigNumber(516));
        let _pat_let_tv0 = _126_v38;
        let _153_v62;
        let _nw18 = new _module.C9();
        _nw18.__ctor((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(544)), ((_154_v50) => function (_155_i4) {
          return (_154_v50)[_module.__default.safeIndex(new BigNumber(838), new BigNumber((_154_v50).length))];
        })(_140_v50))).length)).plus((_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))]), _145_v54, (_150_v59)[_module.__default.safeIndex(_126_v38, new BigNumber((_150_v59).length))], (_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))], (function (_pat_let0_0) {
          return function (_156_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_157_dt__update_hcf44_h0) {
                return function (_pat_let2_0) {
                  return function (_158_dt__update_hcf45_h0) {
                    return _module.D12.create_DC28((_156_dt__update__tmp_h0).dtor_cf43, _157_dt__update_hcf44_h0, _158_dt__update_hcf45_h0, (_156_dt__update__tmp_h0).dtor_cf46, (_156_dt__update__tmp_h0).dtor_cf47);
                  }(_pat_let2_0);
                }(true);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_152_v61)).dtor_cf46);
        _153_v62 = _nw18;
        let _159_v63;
        _159_v63 = _dafny.Seq.of(_153_v62.f7);
        let _160_v64;
        _160_v64 = _dafny.Seq.of((_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))], _153_v62.f7, (_dafny.ZERO).minus(_153_v62.f7), _153_v62.f7, (_159_v63)[_module.__default.safeIndex(new BigNumber(339), new BigNumber((_159_v63).length))]);
        let _161_v65;
        _161_v65 = _module.D9.create_DC20(p3, (_140_v50)[_module.__default.safeIndex(new BigNumber(838), new BigNumber((_140_v50).length))], _153_v62.f7);
        let _162_v66;
        let _nw19 = Array((new BigNumber(13)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = _126_v38;
        _nw19[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(975),false)).length), _153_v62.f7);
        _nw19[(new BigNumber(2)).toNumber()] = _153_v62.f7;
        _nw19[(new BigNumber(3)).toNumber()] = _126_v38;
        _nw19[(new BigNumber(4)).toNumber()] = _153_v62.f7;
        _nw19[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_126_v38);
        _nw19[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_159_v63, _dafny.Seq.Create(_module.__default.abs(new BigNumber(729)), ((_163_v48) => function (_164_i5) {
          return (_163_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_163_v48).length))];
        })(_137_v48)))).length);
        _nw19[(new BigNumber(7)).toNumber()] = (_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))];
        _nw19[(new BigNumber(8)).toNumber()] = _module.__default.fm3((_161_v65).dtor_cf28, p3, globalState);
        _nw19[(new BigNumber(9)).toNumber()] = _153_v62.f7;
        _nw19[(new BigNumber(10)).toNumber()] = (_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))];
        _nw19[(new BigNumber(11)).toNumber()] = ((p1) ? (_126_v38) : (new BigNumber(715)));
        _nw19[(new BigNumber(12)).toNumber()] = _153_v62.f7;
        _162_v66 = _nw19;
        let _index9 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length));
        let _rhs11 = _153_v62;
        let _rhs12 = _module.__default.fm3(_126_v38, p3, globalState);
        let _rhs13 = _162_v66;
        let _rhs14 = p1;
        let _lhs2 = _137_v48;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length));
        _153_v62 = _rhs11;
        _lhs2[_lhs3] = _rhs12;
        _162_v66 = _rhs13;
        r0 = _rhs14;
        let _165_v67;
        let _nw20 = new _module.C6();
        _nw20.__ctor(p3, p3, _149_v58, new BigNumber(-356));
        _165_v67 = _nw20;
        let _166_v68;
        _166_v68 = _dafny.Set.fromElements(_165_v67);
        let _167_v69;
        _167_v69 = _dafny.Map.Empty.slice().updateUnsafe(_166_v68,p1);
        let _index10 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length));
        let _index11 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length));
        let _rhs15 = (_126_v38).multipliedBy(_module.__default.fm3(_126_v38, false, globalState));
        let _rhs16 = (_126_v38).minus((_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))]);
        let _rhs17 = ((_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))]).plus(_module.__default.safeModuloInt(_153_v62.f7, _153_v62.f7));
        let _rhs18 = (((!(p1)) ? (new BigNumber((_167_v69).length)) : ((_dafny.ZERO).minus((_137_v48)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length))])))).plus((_126_v38).plus(_153_v62.f7));
        let _lhs4 = _137_v48;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length));
        let _lhs6 = _137_v48;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_137_v48).length));
        _lhs4[_lhs5] = _rhs15;
        _126_v38 = _rhs16;
        _lhs6[_lhs7] = _rhs17;
        r1 = _rhs18;
      } else {
        let _168_v70;
        _168_v70 = new _dafny.CodePoint('g'.codePointAt(0));
        let _169_v71;
        _169_v71 = _dafny.Map.Empty.slice().updateUnsafe(_168_v70,_126_v38);
        let _170_v72;
        _170_v72 = _dafny.Seq.of(_126_v38, _126_v38, (((_169_v71).contains(_168_v70)) ? ((_169_v71).get(_168_v70)) : (_126_v38)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(540)), ((_171_v38) => function (_172_i6) {
          return _171_v38;
        })(_126_v38))).length));
        let _173_v73;
        _173_v73 = _module.D1.create_DC4(_126_v38, _126_v38);
        let _174_v74;
        _174_v74 = _dafny.Map.Empty.slice().updateUnsafe(p3,_126_v38);
        let _175_v75;
        _175_v75 = _dafny.Set.fromElements(_module.__default.safeDivisionInt(new BigNumber((_170_v72).length), _126_v38), _126_v38, (_173_v73).dtor_cf4, (((_174_v74).contains(p1)) ? ((_174_v74).get(p1)) : (_126_v38)));
        _175_v75 = (((!(p1)) ? (_175_v75) : (_175_v75))).Difference((_175_v75).Intersect(_175_v75));
        let _176_v76;
        _176_v76 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _177_v77;
        _177_v77 = _dafny.MultiSet.fromElements(p3, p3);
        let _178_v78;
        _178_v78 = _module.D11.create_DC24(_module.__default.fm0(_168_v70, globalState), (((_176_v76).contains(false)) ? ((_176_v76).get(false)) : (p3)), new BigNumber((_177_v77).cardinality()));
        let _179_v79;
        _179_v79 = _dafny.MultiSet.fromElements(_178_v78);
        let _180_v80;
        let _nw21 = Array((new BigNumber(26)).toNumber());
        _nw21[(_dafny.ZERO).toNumber()] = p3;
        _nw21[(_dafny.ONE).toNumber()] = p3;
        _nw21[(new BigNumber(2)).toNumber()] = p1;
        _nw21[(new BigNumber(3)).toNumber()] = p1;
        _nw21[(new BigNumber(4)).toNumber()] = p1;
        _nw21[(new BigNumber(5)).toNumber()] = p1;
        _nw21[(new BigNumber(6)).toNumber()] = p3;
        _nw21[(new BigNumber(7)).toNumber()] = _module.__default.fm0(_168_v70, globalState);
        _nw21[(new BigNumber(8)).toNumber()] = !(false);
        _nw21[(new BigNumber(9)).toNumber()] = p3;
        _nw21[(new BigNumber(10)).toNumber()] = p3;
        _nw21[(new BigNumber(11)).toNumber()] = p3;
        _nw21[(new BigNumber(12)).toNumber()] = p3;
        _nw21[(new BigNumber(13)).toNumber()] = !(p1);
        _nw21[(new BigNumber(14)).toNumber()] = p1;
        _nw21[(new BigNumber(15)).toNumber()] = !(p3);
        _nw21[(new BigNumber(16)).toNumber()] = true;
        _nw21[(new BigNumber(17)).toNumber()] = p3;
        _nw21[(new BigNumber(18)).toNumber()] = false;
        _nw21[(new BigNumber(19)).toNumber()] = p1;
        _nw21[(new BigNumber(20)).toNumber()] = p3;
        _nw21[(new BigNumber(21)).toNumber()] = p3;
        _nw21[(new BigNumber(22)).toNumber()] = p1;
        _nw21[(new BigNumber(23)).toNumber()] = p1;
        _nw21[(new BigNumber(24)).toNumber()] = p3;
        _nw21[(new BigNumber(25)).toNumber()] = p1;
        _180_v80 = _nw21;
        let _181_v81;
        let _nw22 = Array((new BigNumber(26)).toNumber());
        _nw22[(_dafny.ZERO).toNumber()] = _180_v80;
        _nw22[(_dafny.ONE).toNumber()] = _180_v80;
        _nw22[(new BigNumber(2)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(3)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(4)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(5)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(6)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(7)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(8)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(9)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(10)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(11)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(12)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(13)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(14)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(15)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(16)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(17)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(18)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(19)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(20)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(21)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(22)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(23)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(24)).toNumber()] = _180_v80;
        _nw22[(new BigNumber(25)).toNumber()] = _180_v80;
        _181_v81 = _nw22;
        let _182_v82;
        let _nw23 = new _module.C5();
        _nw23.__ctor((p3) === (p3), new BigNumber((_179_v79).cardinality()), _181_v81, _126_v38, p3);
        _182_v82 = _nw23;
        _182_v82 = _182_v82;
        let _183_v83;
        let _nw24 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
        _183_v83 = _nw24;
        let _184_v84;
        _184_v84 = _dafny.Map.Empty.slice().updateUnsafe(_183_v83,(_182_v82).f13);
        _184_v84 = (_184_v84).update(_183_v83, p3);
        let _185_v85;
        _185_v85 = _dafny.Seq.of(p1);
        let _186_v86;
        _186_v86 = _dafny.Set.fromElements(_185_v85);
        let _187_v87;
        let _nw25 = new _module.C9();
        _nw25.__ctor((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-613)), ((_188_v70) => function (_189_i7) {
          return _188_v70;
        })(_168_v70))).length)), _186_v86, _181_v81, _182_v82.f14, p1);
        _187_v87 = _nw25;
        let _190_v88;
        _190_v88 = _dafny.Map.Empty.slice().updateUnsafe(_187_v87,p3);
        _190_v88 = (_190_v88).update((((_182_v82).f13) ? (_187_v87) : (_187_v87)), p1);
        let _191_v89;
        _191_v89 = _dafny.Seq.UnicodeFromString("xu");
        let _192_v90;
        let _nw26 = new _module.C2();
        _nw26.__ctor(_191_v89, p1);
        _192_v90 = _nw26;
        let _193_v91;
        _193_v91 = _192_v90;
        _193_v91 = _193_v91;
      }
      if (p3) {
        let _194_v92;
        _194_v92 = new _dafny.CodePoint('k'.codePointAt(0));
        let _195_v93;
        _195_v93 = _dafny.Seq.UnicodeFromString("hvok");
        let _196_v94;
        _196_v94 = _module.D9.create_DC19(_195_v93);
        let _197_v95;
        let _nw27 = new _module.C2();
        _nw27.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-339)), function (_198_i8) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _dafny.Seq.contains((_196_v94).dtor_cf25, _194_v92));
        _197_v95 = _nw27;
        let _rhs19 = _197_v95;
        let _rhs20 = (_dafny.ZERO).minus(_126_v38);
        _197_v95 = _rhs19;
        _126_v38 = _rhs20;
        let _199_v96;
        _199_v96 = _dafny.MultiSet.fromElements((_197_v95).f4, p3);
        let _200_v97;
        _200_v97 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
        let _201_v98;
        _201_v98 = _dafny.Seq.of(_200_v97, _200_v97);
        let _source7 = _module.D16.create_DC33((_197_v95).f4, _126_v38, (_199_v96).Difference(_199_v96), _126_v38, _module.__default.fm53(_126_v38, _201_v98, globalState));
        if (_source7.is_DC33) {
          let _202___mcc_h0 = (_source7).cf52;
          let _203___mcc_h1 = (_source7).cf53;
          let _204___mcc_h2 = (_source7).cf54;
          let _205___mcc_h3 = (_source7).cf55;
          let _206___mcc_h4 = (_source7).cf56;
          let _207_cf56 = _206___mcc_h4;
          let _208_cf55 = _205___mcc_h3;
          let _209_cf54 = _204___mcc_h2;
          let _210_cf53 = _203___mcc_h1;
          let _211_cf52 = _202___mcc_h0;
          let _212_v99;
          let _nw28 = Array((new BigNumber(20)).toNumber());
          _212_v99 = _nw28;
          let _213_v100;
          let _nw29 = Array((new BigNumber(14)).toNumber()).fill([]);
          _213_v100 = _nw29;
          let _214_v101;
          _214_v101 = _dafny.Seq.of(_210_cf53);
          let _215_v102;
          let _nw30 = new _module.C4();
          _nw30.__ctor(_213_v100, (_214_v101)[_module.__default.safeIndex(new BigNumber((_214_v101).length), new BigNumber((_214_v101).length))], p3);
          _215_v102 = _nw30;
          let _index12 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_212_v99).length));
          (_212_v99)[_index12] = _215_v102;
          let _216_v103;
          _216_v103 = _dafny.Seq.of(!(p3));
          let _217_v106;
          _217_v106 = _dafny.MultiSet.fromElements(function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of _dafny.IntegerRange(new BigNumber(438), new BigNumber(391))) {
              let _218_v104 = _compr_19;
              if (((new BigNumber(438)).isLessThanOrEqualTo(_218_v104)) && ((_218_v104).isLessThan(new BigNumber(391)))) {
                _coll19.push([_module.__default.safeDivisionInt(_218_v104, new BigNumber((_200_v97).length)),p1]);
              }
            }
            return _coll19;
          }(), function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of _dafny.IntegerRange(new BigNumber(355), new BigNumber(267))) {
              let _219_v105 = _compr_20;
              if (((new BigNumber(355)).isLessThanOrEqualTo(_219_v105)) && ((_219_v105).isLessThan(new BigNumber(267)))) {
                _coll20.push([(_219_v105).multipliedBy(_126_v38),p1]);
              }
            }
            return _coll20;
          }());
          let _index13 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_212_v99).length));
          let _rhs21 = _210_cf53;
          let _rhs22 = (_216_v103)[_module.__default.safeIndex(new BigNumber((_217_v106).cardinality()), new BigNumber((_216_v103).length))];
          let _rhs23 = (((_215_v102).fm8(_210_cf53, _211_cf52, _dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), ((_220_v92) => function (_221_i9) {
            return _220_v92;
          })(_194_v92)), _214_v101, globalState)) ? (_194_v92) : (new _dafny.CodePoint('p'.codePointAt(0))));
          let _rhs24 = _215_v102;
          let _lhs8 = _212_v99;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_212_v99).length));
          r1 = _rhs21;
          r0 = _rhs22;
          _194_v92 = _rhs23;
          _lhs8[_lhs9] = _rhs24;
          let _222_v107;
          let _nw31 = Array((new BigNumber(3)).toNumber());
          _222_v107 = _nw31;
          let _223_v108;
          _223_v108 = _dafny.Set.fromElements(_216_v103, _216_v103, _dafny.Seq.of((_197_v95).f4), _216_v103, _216_v103);
          let _224_v109;
          let _nw32 = new _module.C9();
          _nw32.__ctor(_210_cf53, _223_v108, _213_v100, _208_cf55, _211_cf52);
          _224_v109 = _nw32;
          let _index14 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_222_v107).length));
          (_222_v107)[_index14] = _224_v109;
          let _index15 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_222_v107).length));
          (_222_v107)[_index15] = _224_v109;
          _211_cf52 = !(false);
          _211_cf52 = ((_dafny.ZERO).minus((_215_v102).fm7(globalState))).isLessThan(_210_cf53);
        } else {
          let _225___mcc_h5 = (_source7).cf51;
          let _226_cf51 = _225___mcc_h5;
          _194_v92 = ((((p1) ? ((_197_v95).f4) : ((_197_v95).f4))) ? (_194_v92) : (_module.__default.fm1(_126_v38, (_197_v95).f4, globalState)));
          let _227_v110;
          let _nw33 = Array((new BigNumber(20)).toNumber()).fill(false);
          _227_v110 = _nw33;
          let _228_v111;
          _228_v111 = _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_227_v110, _227_v110, _227_v110, _227_v110, _227_v110)).length));
          r0 = ((_dafny.ZERO).minus((((_197_v95).f4) ? (_module.__default.fm3(_126_v38, (_197_v95).f4, globalState)) : (_126_v38)))).isLessThan((_228_v111)[_module.__default.safeIndex(_126_v38, new BigNumber((_228_v111).length))]);
          let _229_v112;
          let _nw34 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Set.Empty);
          _229_v112 = _nw34;
          let _230_v113;
          let _nw35 = Array((new BigNumber(9)).toNumber()).fill([]);
          _230_v113 = _nw35;
          let _231_v114;
          _231_v114 = _230_v113;
          let _232_v115;
          _232_v115 = _dafny.Set.fromElements(_231_v114, _231_v114);
          let _233_v116;
          _233_v116 = _dafny.Set.fromElements(_dafny.Set.fromElements(_231_v114), _232_v115);
          let _index16 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_229_v112).length));
          (_229_v112)[_index16] = _233_v116;
          let _index17 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_229_v112).length));
          (_229_v112)[_index17] = _dafny.Set.fromElements(_232_v115);
          r1 = (_dafny.ZERO).minus(_126_v38);
        }
        let _234_v117;
        _234_v117 = _dafny.Map.Empty.slice().updateUnsafe(_126_v38,_137_v48);
        _234_v117 = (_234_v117).update(_126_v38, _137_v48);
        _126_v38 = _126_v38;
        r0 = p1;
      } else {
        r0 = (_126_v38).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), function (_235_i10) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length));
        let _236_v118;
        let _nw36 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
        _236_v118 = _nw36;
        let _237_v119;
        _237_v119 = _module.D22.create_DC45(_236_v118);
        _237_v119 = _module.D22.create_DC45(_236_v118);
        let _238_v120;
        _238_v120 = _dafny.Map.Empty.slice().updateUnsafe(_126_v38,p3);
        let _239_v122;
        let _nw37 = Array((new BigNumber(22)).toNumber()).fill([]);
        _239_v122 = _nw37;
        let _240_v123;
        _240_v123 = _dafny.Seq.of(_126_v38, _126_v38);
        let _241_v124;
        let _nw38 = new _module.C1();
        _nw38.__ctor(new BigNumber((function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of (_238_v120).Keys.Elements) {
            let _242_v121 = _compr_21;
            if ((_238_v120).contains(_242_v121)) {
              _coll21.add((_242_v121).plus(new BigNumber((_dafny.Set.fromElements(true, true)).length)));
            }
          }
          return _coll21;
        }()).length), _126_v38, _239_v122, new BigNumber((_240_v123).length), p1);
        _241_v124 = _nw38;
        let _index18 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_137_v48).length));
        (_137_v48)[_index18] = _126_v38;
        let _index19 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_137_v48).length));
        (_137_v48)[_index19] = _126_v38;
        _126_v38 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(623)), function (_243_i11) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        })).length);
      }
      let _244_v125;
      _244_v125 = _dafny.Seq.UnicodeFromString("rjhatt");
      let _245_v126;
      let _nw39 = Array((new BigNumber(22)).toNumber()).fill(false);
      _245_v126 = _nw39;
      let _246_v127;
      _246_v127 = _dafny.Seq.of(_245_v126, _245_v126);
      let _247_v128;
      _247_v128 = _module.D20.create_DC41(_126_v38);
      let _248_v129;
      _248_v129 = _module.D10.create_DC22(_244_v125, p1, p3, (_246_v127)[_module.__default.safeIndex(_126_v38, new BigNumber((_246_v127).length))], (_247_v128).dtor_cf67);
      let _249_v130;
      _249_v130 = _dafny.Map.Empty.slice().updateUnsafe(_126_v38,p3);
      let _250_v131;
      _250_v131 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _251_v132;
      let _nw40 = Array((new BigNumber(24)).toNumber()).fill([]);
      _251_v132 = _nw40;
      let _252_v133;
      let _nw41 = new _module.C6();
      _nw41.__ctor((((_249_v130).contains(_126_v38)) ? ((_249_v130).get(_126_v38)) : (p1)), !((((_250_v131).contains(p1)) ? ((_250_v131).get(p1)) : (p1))), _251_v132, _module.__default.fm3(_126_v38, false, globalState));
      _252_v133 = _nw41;
      let _253_v134;
      _253_v134 = _dafny.Map.Empty.slice().updateUnsafe(_252_v133,_126_v38);
      r0 = (new BigNumber((_253_v134).length)).isLessThanOrEqualTo(new BigNumber(((_248_v129).dtor_cf30).length));
      let _254_v135;
      _254_v135 = new _dafny.CodePoint('t'.codePointAt(0));
      r1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ewmfdfes"), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(714)), function (_255_i12) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length), new BigNumber((_dafny.Seq.UnicodeFromString("ewmfdfes")).length)), _254_v135), _dafny.Seq.UnicodeFromString("nv"))).length);
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _256_v0;
      _256_v0 = false;
      let _257_v1;
      _257_v1 = _dafny.MultiSet.fromElements(_256_v0, _256_v0, _256_v0);
      let _258_globalState;
      let _nw42 = new _module.GlobalState();
      _nw42.__ctor(_257_v1, new BigNumber(-993), false, new BigNumber(862));
      _258_globalState = _nw42;
      let _259_v2;
      _259_v2 = new BigNumber(598);
      let _260_v3;
      let _nw43 = Array((new BigNumber(14)).toNumber());
      _nw43[(_dafny.ZERO).toNumber()] = _259_v2;
      _nw43[(_dafny.ONE).toNumber()] = _259_v2;
      _nw43[(new BigNumber(2)).toNumber()] = _259_v2;
      _nw43[(new BigNumber(3)).toNumber()] = ((_256_v0) ? (_259_v2) : (_259_v2));
      _nw43[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(532)), function (_261_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length), new BigNumber(334));
      _nw43[(new BigNumber(5)).toNumber()] = new BigNumber(-930);
      _nw43[(new BigNumber(6)).toNumber()] = _259_v2;
      _nw43[(new BigNumber(7)).toNumber()] = _259_v2;
      _nw43[(new BigNumber(8)).toNumber()] = _259_v2;
      _nw43[(new BigNumber(9)).toNumber()] = ((_module.D0.create_DC2(_259_v2)).dtor_cf1).minus((_dafny.ZERO).minus(_259_v2));
      _nw43[(new BigNumber(10)).toNumber()] = _259_v2;
      _nw43[(new BigNumber(11)).toNumber()] = (_259_v2).minus(_259_v2);
      _nw43[(new BigNumber(12)).toNumber()] = ((_dafny.ZERO).minus((_dafny.ZERO).minus(_259_v2))).multipliedBy(new BigNumber(743));
      _nw43[(new BigNumber(13)).toNumber()] = _259_v2;
      _260_v3 = _nw43;
      _260_v3 = _260_v3;
      let _262_v4;
      _262_v4 = new _dafny.CodePoint('j'.codePointAt(0));
      _262_v4 = _262_v4;
      let _263_v5;
      _263_v5 = _dafny.Map.Empty.slice().updateUnsafe((_259_v2).multipliedBy(_259_v2),_256_v0);
      _263_v5 = (_263_v5).update((_259_v2).minus(_259_v2), _256_v0);
      let _264_v6;
      _264_v6 = _dafny.Set.fromElements(_259_v2, _259_v2);
      let _265_v7;
      _265_v7 = _dafny.Set.fromElements(_264_v6);
      let _266_v8;
      _266_v8 = _dafny.Seq.UnicodeFromString("smw");
      let _267_v9;
      _267_v9 = _dafny.Map.Empty.slice().updateUnsafe(_265_v7,!(new BigNumber((_266_v8).length)).isEqualTo(_259_v2));
      _256_v0 = (((_267_v9).contains(_265_v7)) ? ((_267_v9).get(_265_v7)) : (_256_v0));
      let _268_v10;
      let _nw44 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Set.Empty);
      _268_v10 = _nw44;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_268_v10).length))) {
        let _269_i1 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_269_i1)) && ((_269_i1).isLessThan(new BigNumber((_268_v10).length))))) {
          (_268_v10)[(_269_i1)] = ((_dafny.Set.fromElements(_256_v0)).Difference(_dafny.Set.fromElements(_256_v0, false))).Difference((_dafny.Set.fromElements(_256_v0, true, _256_v0, false)).Intersect(_dafny.Set.fromElements(_256_v0, _256_v0, (_module.D1.create_DC3(_256_v0)).dtor_cf2, _256_v0, _256_v0)));
        }
      }
      _256_v0 = !(_256_v0) || (_256_v0);
      let _270_v11;
      let _init2 = ((_271_v0) => function (_272_i2) {
        return _271_v0;
      })(_256_v0);
      let _nw45 = Array((new BigNumber(6)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw45.length); _i0_2++) {
        _nw45[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _270_v11 = _nw45;
      let _273_v12;
      _273_v12 = _dafny.Seq.of(_270_v11, _270_v11);
      let _274_v13;
      _274_v13 = _dafny.Map.Empty.slice().updateUnsafe(!(_256_v0),_262_v4);
      let _275_v14;
      _275_v14 = _dafny.Seq.of(_259_v2);
      let _276_v15;
      let _nw46 = Array((new BigNumber(21)).toNumber());
      _nw46[(_dafny.ZERO).toNumber()] = !(false);
      _nw46[(_dafny.ONE).toNumber()] = _module.__default.fm0(_module.__default.fm1((_dafny.ZERO).minus(_259_v2), _256_v0, _258_globalState), _258_globalState);
      _nw46[(new BigNumber(2)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(3)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(4)).toNumber()] = !(_module.__default.fm0(new _dafny.CodePoint('e'.codePointAt(0)), _258_globalState));
      _nw46[(new BigNumber(5)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(6)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(7)).toNumber()] = !(_256_v0);
      _nw46[(new BigNumber(8)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(9)).toNumber()] = !(_256_v0);
      _nw46[(new BigNumber(10)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(11)).toNumber()] = _dafny.Seq.contains(_273_v12, _270_v11);
      _nw46[(new BigNumber(12)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(13)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(14)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(15)).toNumber()] = !(_259_v2).isEqualTo((((_257_v1).contains(_256_v0)) ? ((_257_v1).get(_256_v0)) : (new BigNumber((_266_v8).length))));
      _nw46[(new BigNumber(16)).toNumber()] = ((_256_v0) ? (_256_v0) : (!(_256_v0)));
      _nw46[(new BigNumber(17)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(18)).toNumber()] = _256_v0;
      _nw46[(new BigNumber(19)).toNumber()] = !(_module.__default.fm0((((_274_v13).contains(_256_v0)) ? ((_274_v13).get(_256_v0)) : (_262_v4)), _258_globalState));
      _nw46[(new BigNumber(20)).toNumber()] = (new BigNumber((_266_v8).length)).isLessThan(new BigNumber((_275_v14).length));
      _276_v15 = _nw46;
      let _index20 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length));
      (_276_v15)[_index20] = false;
      let _277_v16;
      _277_v16 = _module.D0.create_DC1();
      let _278_v17;
      _278_v17 = _dafny.Seq.of(_277_v16, _module.__default.fm2(_258_globalState));
      let _index21 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length));
      let _rhs25 = !(_dafny.areEqual(_dafny.Seq.of(_module.D0.create_DC1()), _dafny.Seq.Concat(_278_v17, _278_v17)));
      let _rhs26 = _module.__default.fm0(_262_v4, _258_globalState);
      let _lhs10 = _276_v15;
      let _lhs11 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length));
      _256_v0 = _rhs25;
      _lhs10[_lhs11] = _rhs26;
      let _279_v18;
      _279_v18 = _dafny.MultiSet.fromElements(_259_v2);
      let _280_v19;
      _280_v19 = _dafny.Map.Empty.slice().updateUnsafe(_279_v18,_266_v8);
      let _281_v20;
      let _282_v21;
      let _out0;
      let _out1;
      let _outcollector0 = _module.__default.m0(_280_v19, (((_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))]) ? (!(_256_v0)) : ((_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))])), (_279_v18).update(_259_v2, _module.__default.abs(_259_v2)), (_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))], _258_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _281_v20 = _out0;
      _282_v21 = _out1;
      let _283_v22;
      let _init3 = ((_284_v14) => function (_285_i3) {
        return _dafny.Seq.Concat(_284_v14, _284_v14);
      })(_275_v14);
      let _nw47 = Array((new BigNumber(24)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw47.length); _i0_3++) {
        _nw47[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _283_v22 = _nw47;
      let _index22 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_283_v22).length));
      (_283_v22)[_index22] = _275_v14;
      let _286_v23;
      _286_v23 = _module.D0.create_DC2(_259_v2);
      let _pat_let_tv1 = _277_v16;
      let _pat_let_tv2 = _277_v16;
      let _pat_let_tv3 = _277_v16;
      let _index23 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length));
      let _index24 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_283_v22).length));
      let _rhs27 = function (_source8) {
        if (_source8.is_DC1) {
          return _pat_let_tv1;
        } else if (_source8.is_DC2) {
          let _287___mcc_h0 = (_source8).cf1;
          let _288_cf1 = _287___mcc_h0;
          return _pat_let_tv2;
        } else {
          let _289___mcc_h1 = (_source8).cf0;
          let _290_cf0 = _289___mcc_h1;
          return _pat_let_tv3;
        }
      }(_286_v23);
      let _rhs28 = true;
      let _rhs29 = _module.__default.fm0(_262_v4, _258_globalState);
      let _rhs30 = _275_v14;
      let _rhs31 = _259_v2;
      let _lhs12 = _276_v15;
      let _lhs13 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length));
      let _lhs14 = _283_v22;
      let _lhs15 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_283_v22).length));
      _277_v16 = _rhs27;
      _lhs12[_lhs13] = _rhs28;
      _281_v20 = _rhs29;
      _lhs14[_lhs15] = _rhs30;
      _259_v2 = _rhs31;
      _286_v23 = _286_v23;
      if ((_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))]) {
        let _291_v24;
        _291_v24 = _dafny.Map.Empty.slice().updateUnsafe(_259_v2,_260_v3);
        _260_v3 = (((_291_v24).contains(_259_v2)) ? ((_291_v24).get(_259_v2)) : (_260_v3));
        let _index25 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_260_v3).length));
        (_260_v3)[_index25] = (_282_v21).multipliedBy(_module.__default.fm3(_259_v2, (_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))], _258_globalState));
        let _292_v25;
        _292_v25 = _dafny.Seq.of(_279_v18, (_279_v18).Union(_dafny.MultiSet.fromElements(_259_v2, new BigNumber(673))));
        let _index26 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_260_v3).length));
        let _rhs32 = !(_256_v0);
        let _rhs33 = _module.__default.safeDivisionInt(_282_v21, _259_v2);
        let _rhs34 = _dafny.Seq.Concat(_292_v25, _module.__default.fm4(false, true, _258_globalState));
        let _rhs35 = _286_v23;
        let _lhs16 = _260_v3;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_260_v3).length));
        _281_v20 = _rhs32;
        _lhs16[_lhs17] = _rhs33;
        _292_v25 = _rhs34;
        _286_v23 = _rhs35;
        let _index27 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_260_v3).length));
        (_260_v3)[_index27] = new BigNumber(902);
        let _index28 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length));
        (_276_v15)[_index28] = _281_v20;
        let _293_v26;
        _293_v26 = _module.D2.create_DC5(_276_v15);
        _270_v11 = (_293_v26).dtor_cf5;
      } else {
        let _index29 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_260_v3).length));
        (_260_v3)[_index29] = new BigNumber((((_281_v20) ? (_266_v8) : (_266_v8))).length);
        let _294_v27;
        let _nw48 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _294_v27 = _nw48;
        let _index30 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_294_v27).length));
        (_294_v27)[_index30] = _dafny.Seq.UnicodeFromString("nwenxedn");
        let _295_v28;
        _295_v28 = _dafny.Seq.of(_266_v8, _266_v8, _dafny.Seq.UnicodeFromString("qystgqh"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(635)), ((_296_v4) => function (_297_i4) {
          return _296_v4;
        })(_262_v4)), _dafny.Seq.Concat(_266_v8, _266_v8));
        let _index31 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_260_v3).length));
        let _index32 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_294_v27).length));
        let _rhs36 = new BigNumber((_295_v28).length);
        let _rhs37 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(912)), function (_298_i5) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        });
        let _lhs18 = _260_v3;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_260_v3).length));
        let _lhs20 = _294_v27;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_294_v27).length));
        _lhs18[_lhs19] = _rhs36;
        _lhs20[_lhs21] = _rhs37;
        let _299_v30;
        _299_v30 = _dafny.Seq.of(_279_v18);
        let _300_v31;
        _300_v31 = _dafny.Map.Empty.slice().updateUnsafe(true,_280_v19);
        let _301_v32;
        let _302_v33;
        let _out2;
        let _out3;
        let _outcollector1 = _module.__default.m0((function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of (_299_v30).Elements) {
            let _303_v29 = _compr_22;
            if (_dafny.Seq.contains(_299_v30, _303_v29)) {
              _coll22.push([_303_v29,_266_v8]);
            }
          }
          return _coll22;
        }()).Merge((((_300_v31).contains(_281_v20)) ? ((_300_v31).get(_281_v20)) : (_280_v19))), _256_v0, _279_v18, !(_256_v0), _258_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _301_v32 = _out2;
        _302_v33 = _out3;
        _256_v0 = (_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))];
        let _304_v34;
        _304_v34 = _dafny.Set.fromElements(_266_v8);
        let _305_v35;
        _305_v35 = _dafny.Map.Empty.slice().updateUnsafe(_256_v0,_301_v32);
        let _index33 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_260_v3).length));
        let _index34 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_260_v3).length));
        let _index35 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_260_v3).length));
        let _rhs38 = _module.__default.fm3((new BigNumber((_304_v34).length)).plus(_302_v33), !((_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))]), _258_globalState);
        let _rhs39 = _279_v18;
        let _rhs40 = _282_v21;
        let _rhs41 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_259_v2, (_dafny.ZERO).minus(((_281_v20) ? (_302_v33) : (new BigNumber((_305_v35).length))))));
        let _rhs42 = _302_v33;
        let _lhs22 = _260_v3;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_260_v3).length));
        let _lhs24 = _260_v3;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_260_v3).length));
        let _lhs26 = _260_v3;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_260_v3).length));
        _282_v21 = _rhs38;
        _279_v18 = _rhs39;
        _lhs22[_lhs23] = _rhs40;
        _lhs24[_lhs25] = _rhs41;
        _lhs26[_lhs27] = _rhs42;
        _256_v0 = !(_301_v32) || (_301_v32);
      }
      let _306_v36;
      _306_v36 = _dafny.Map.Empty.slice().updateUnsafe(_277_v16,_282_v21);
      let _307_v40;
      _307_v40 = _dafny.Set.fromElements(_module.__default.fm0(_262_v4, _258_globalState));
      let _308_v41;
      let _nw49 = Array((new BigNumber(12)).toNumber());
      _nw49[(_dafny.ZERO).toNumber()] = _306_v36;
      _nw49[(_dafny.ONE).toNumber()] = (_306_v36).Merge(_306_v36);
      _nw49[(new BigNumber(2)).toNumber()] = _306_v36;
      _nw49[(new BigNumber(3)).toNumber()] = _306_v36;
      _nw49[(new BigNumber(4)).toNumber()] = function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of (_306_v36).Keys.Elements) {
          let _309_v37 = _compr_23;
          if ((_306_v36).contains(_309_v37)) {
            _coll23.push([_309_v37,(_275_v14)[_module.__default.safeIndex(_259_v2, new BigNumber((_275_v14).length))]]);
          }
        }
        return _coll23;
      }();
      _nw49[(new BigNumber(5)).toNumber()] = _306_v36;
      _nw49[(new BigNumber(6)).toNumber()] = _306_v36;
      _nw49[(new BigNumber(7)).toNumber()] = _306_v36;
      _nw49[(new BigNumber(8)).toNumber()] = (function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of (_dafny.Seq.of(_277_v16, _277_v16)).Elements) {
          let _310_v38 = _compr_24;
          if (_dafny.Seq.contains(_dafny.Seq.of(_277_v16, _277_v16), _310_v38)) {
            _coll24.push([_310_v38,new BigNumber((function () {
              let _coll25 = new _dafny.Map();
              for (const _compr_25 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D2.create_DC7(_259_v2)))).Elements) {
                let _311_v39 = _compr_25;
                if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D2.create_DC7(_259_v2)))).contains(_311_v39)) {
                  _coll25.push([_311_v39,_282_v21]);
                }
              }
              return _coll25;
            }()).length)]);
          }
        }
        return _coll24;
      }()).update(_module.D0.create_DC1(), _module.__default.fm3(new BigNumber((_307_v40).length), _256_v0, _258_globalState));
      _nw49[(new BigNumber(9)).toNumber()] = _306_v36;
      _nw49[(new BigNumber(10)).toNumber()] = _module.__default.fm5(_282_v21, (_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))], _258_globalState);
      _nw49[(new BigNumber(11)).toNumber()] = _306_v36;
      _308_v41 = _nw49;
      _308_v41 = ((!(_256_v0)) ? (_308_v41) : (_308_v41));
      let _312_v42;
      _312_v42 = _dafny.MultiSet.fromElements(_262_v4);
      let _313_v43;
      let _314_v44;
      let _out4;
      let _out5;
      let _outcollector2 = _module.__default.m0(_280_v19, false, _279_v18, (_312_v42).contains(new _dafny.CodePoint('n'.codePointAt(0))), _258_globalState);
      _out4 = _outcollector2[0];
      _out5 = _outcollector2[1];
      _313_v43 = _out4;
      _314_v44 = _out5;
      let _index36 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length));
      (_260_v3)[_index36] = _314_v44;
      let _315_v45;
      _315_v45 = _dafny.Map.Empty.slice().updateUnsafe(_260_v3,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-750)), ((_316_v21) => function (_317_i6) {
        return _316_v21;
      })(_282_v21)));
      let _index37 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length));
      (_260_v3)[_index37] = (_dafny.ZERO).minus((new BigNumber((((_281_v20) ? (_315_v45) : (_315_v45))).length)).multipliedBy(_259_v2));
      let _hi1 = _282_v21;
      for (let _318_i7 = _314_v44; _318_i7.isLessThan(_hi1); _318_i7 = _318_i7.plus(_dafny.ONE)) {
        let _index38 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length));
        (_260_v3)[_index38] = (_259_v2).plus((_260_v3)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length))]);
        let _index39 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length));
        (_260_v3)[_index39] = _259_v2;
        let _319_v46;
        _319_v46 = _module.D19.create_DC38(_281_v20, _313_v43);
        let _320_v47;
        _320_v47 = _dafny.Seq.of(false, false);
        let _321_v48;
        _321_v48 = _dafny.Set.fromElements(_320_v47, _320_v47, _320_v47);
        let _322_v49;
        _322_v49 = _dafny.Map.Empty.slice().updateUnsafe(_256_v0,_276_v15);
        let _323_v50;
        _323_v50 = _module.D10.create_DC22(_266_v8, _256_v0, _281_v20, _276_v15, new BigNumber(867));
        let _324_v51;
        let _nw50 = Array((new BigNumber(18)).toNumber());
        _nw50[(_dafny.ZERO).toNumber()] = _270_v11;
        _nw50[(_dafny.ONE).toNumber()] = _270_v11;
        _nw50[(new BigNumber(2)).toNumber()] = _276_v15;
        _nw50[(new BigNumber(3)).toNumber()] = _276_v15;
        _nw50[(new BigNumber(4)).toNumber()] = _276_v15;
        _nw50[(new BigNumber(5)).toNumber()] = _270_v11;
        _nw50[(new BigNumber(6)).toNumber()] = _270_v11;
        _nw50[(new BigNumber(7)).toNumber()] = _276_v15;
        _nw50[(new BigNumber(8)).toNumber()] = _270_v11;
        _nw50[(new BigNumber(9)).toNumber()] = _270_v11;
        _nw50[(new BigNumber(10)).toNumber()] = _276_v15;
        _nw50[(new BigNumber(11)).toNumber()] = _270_v11;
        _nw50[(new BigNumber(12)).toNumber()] = (((_322_v49).contains(_256_v0)) ? ((_322_v49).get(_256_v0)) : (_270_v11));
        _nw50[(new BigNumber(13)).toNumber()] = _276_v15;
        _nw50[(new BigNumber(14)).toNumber()] = _270_v11;
        _nw50[(new BigNumber(15)).toNumber()] = _276_v15;
        _nw50[(new BigNumber(16)).toNumber()] = _270_v11;
        _nw50[(new BigNumber(17)).toNumber()] = (_323_v50).dtor_cf33;
        _324_v51 = _nw50;
        let _325_v52;
        let _nw51 = new _module.C9();
        _nw51.__ctor(_module.__default.fm3(_318_i7, (_319_v46).dtor_cf64, _258_globalState), _321_v48, _324_v51, _259_v2, _module.__default.fm0(_262_v4, _258_globalState));
        _325_v52 = _nw51;
        _260_v3 = _260_v3;
      }
      if (!(!(_281_v20))) {
        let _326_v53;
        let _nw52 = Array((new BigNumber(10)).toNumber());
        _nw52[(_dafny.ZERO).toNumber()] = _270_v11;
        _nw52[(_dafny.ONE).toNumber()] = _270_v11;
        _nw52[(new BigNumber(2)).toNumber()] = _276_v15;
        _nw52[(new BigNumber(3)).toNumber()] = _276_v15;
        _nw52[(new BigNumber(4)).toNumber()] = _270_v11;
        _nw52[(new BigNumber(5)).toNumber()] = _270_v11;
        _nw52[(new BigNumber(6)).toNumber()] = _270_v11;
        _nw52[(new BigNumber(7)).toNumber()] = _270_v11;
        _nw52[(new BigNumber(8)).toNumber()] = _276_v15;
        _nw52[(new BigNumber(9)).toNumber()] = _276_v15;
        _326_v53 = _nw52;
        let _327_v54;
        let _nw53 = new _module.C5();
        _nw53.__ctor(_281_v20, _259_v2, _326_v53, new BigNumber((_266_v8).length), _313_v43);
        _327_v54 = _nw53;
        let _index40 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length));
        (_260_v3)[_index40] = new BigNumber(((_module.D26.create_DC53(_dafny.Seq.of(_327_v54, _327_v54, _327_v54))).dtor_cf89).length);
        let _328_v55;
        _328_v55 = _dafny.Set.fromElements(_262_v4, _262_v4);
        let _329_v56;
        _329_v56 = _module.D6.create_DC16((_327_v54).f13, (((_263_v5).contains(new BigNumber((_328_v55).length))) ? ((_263_v5).get(new BigNumber((_328_v55).length))) : (_313_v43)), (_283_v22)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_283_v22).length))]);
        let _330_v57;
        _330_v57 = _dafny.Seq.of(true, (_327_v54).fm8(_282_v21, _256_v0, _266_v8, (_329_v56).dtor_cf22, _258_globalState));
        _262_v4 = (_266_v8)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_330_v57).length), _259_v2), new BigNumber((_266_v8).length))];
        let _331_v59;
        _331_v59 = _dafny.Seq.of(_326_v53);
        let _332_v60;
        let _nw54 = new _module.C9();
        _nw54.__ctor(_259_v2, function () {
          let _coll26 = new _dafny.Set();
          for (const _compr_26 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-88)), ((_333_v57) => function (_334_i8) {
            return _333_v57;
          })(_330_v57))).Elements) {
            let _335_v58 = _compr_26;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-88)), ((_336_v57) => function (_334_i8) {
              return _336_v57;
            })(_330_v57)), _335_v58)) {
              _coll26.add(_335_v58);
            }
          }
          return _coll26;
        }(), (_331_v59)[_module.__default.safeIndex(_259_v2, new BigNumber((_331_v59).length))], _314_v44, (_327_v54).f13);
        _332_v60 = _nw54;
        let _337_v61;
        _337_v61 = _dafny.Set.fromElements(_332_v60);
        _337_v61 = _337_v61;
        let _index41 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length));
        (_260_v3)[_index41] = (_332_v60).fm9((_327_v54).f13, (_327_v54).fm8((_260_v3)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length))], _313_v43, _266_v8, _dafny.Seq.of(_327_v54.f14, (_260_v3)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length))]), _258_globalState), _282_v21, _258_globalState);
        _313_v43 = (_327_v54).f13;
      } else {
        let _338_v63;
        _338_v63 = _dafny.MultiSet.fromElements(_279_v18);
        let _339_v64;
        let _340_v65;
        let _out6;
        let _out7;
        let _outcollector3 = _module.__default.m0(function () {
          let _coll27 = new _dafny.Map();
          for (const _compr_27 of (_338_v63).Elements) {
            let _341_v62 = _compr_27;
            if ((_338_v63).contains(_341_v62)) {
              _coll27.push([_341_v62,_266_v8]);
            }
          }
          return _coll27;
        }(), _281_v20, _dafny.MultiSet.FromArray(_275_v14), (_281_v20) || (_256_v0), _258_globalState);
        _out6 = _outcollector3[0];
        _out7 = _outcollector3[1];
        _339_v64 = _out6;
        _340_v65 = _out7;
        let _index42 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length));
        (_260_v3)[_index42] = (new BigNumber((_279_v18).cardinality())).multipliedBy(new BigNumber((_264_v6).length));
        let _342_v66;
        _342_v66 = _module.D16.create_DC33(_313_v43, _259_v2, _257_v1, _314_v44, _307_v40);
        if ((!(new BigNumber(980)).isEqualTo((_260_v3)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length))])) && ((_342_v66).dtor_cf52)) {
          let _343_v67;
          let _344_v68;
          let _out8;
          let _out9;
          let _outcollector4 = _module.__default.m0(_280_v19, _313_v43, _279_v18, _281_v20, _258_globalState);
          _out8 = _outcollector4[0];
          _out9 = _outcollector4[1];
          _343_v67 = _out8;
          _344_v68 = _out9;
          _343_v67 = _339_v64;
          let _345_v69;
          _345_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(320),_279_v18);
          let _346_v70;
          _346_v70 = _dafny.Seq.of(_256_v0, !(_345_v69).equals(_345_v69));
          let _index43 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length));
          (_276_v15)[_index43] = (_346_v70)[_module.__default.safeIndex(_259_v2, new BigNumber((_346_v70).length))];
          let _index44 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length));
          (_276_v15)[_index44] = _281_v20;
          _256_v0 = ((((_339_v64) ? (_256_v0) : (_339_v64))) ? (_256_v0) : ((_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))]));
        } else {
          let _347_v71;
          _347_v71 = _dafny.Seq.of(_339_v64, _313_v43);
          _256_v0 = (_347_v71)[_module.__default.safeIndex((_260_v3)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length))], new BigNumber((_347_v71).length))];
          _266_v8 = _dafny.Seq.UnicodeFromString("becsmvx");
          let _348_v72;
          let _349_v73;
          let _out10;
          let _out11;
          let _outcollector5 = _module.__default.m0(_280_v19, (_259_v2).isLessThanOrEqualTo(_259_v2), _279_v18, _313_v43, _258_globalState);
          _out10 = _outcollector5[0];
          _out11 = _outcollector5[1];
          _348_v72 = _out10;
          _349_v73 = _out11;
          let _350_v74;
          _350_v74 = _module.D21.create_DC44(false);
          let _pat_let_tv4 = _313_v43;
          _350_v74 = function (_pat_let3_0) {
            return function (_351_dt__update__tmp_h0) {
              return function (_pat_let4_0) {
                return function (_352_dt__update_hcf72_h0) {
                  return _module.D21.create_DC44(_352_dt__update_hcf72_h0);
                }(_pat_let4_0);
              }(_pat_let_tv4);
            }(_pat_let3_0);
          }(_350_v74);
          _348_v72 = !((_279_v18).Union(_dafny.MultiSet.fromElements(new BigNumber(-977), new BigNumber(547), (((_257_v1).contains(_281_v20)) ? ((_257_v1).get(_281_v20)) : (_314_v44))))).contains(_module.__default.fm3((_260_v3)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length))], _348_v72, _258_globalState));
        }
        let _353_v75;
        _353_v75 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_259_v2, _340_v65), _275_v14);
        let _354_v76;
        _354_v76 = _353_v75;
        let _355_v77;
        _355_v77 = _dafny.Seq.of(_281_v20);
        let _356_v78;
        _356_v78 = _dafny.Seq.of(_355_v77);
        let _357_v79;
        _357_v79 = _module.D12.create_DC27(_356_v78);
        let _358_v80;
        _358_v80 = _dafny.Map.Empty.slice().updateUnsafe(_266_v8,_357_v79);
        let _359_v81;
        _359_v81 = _dafny.Seq.of(_358_v80, _358_v80);
        let _360_v82;
        let _nw55 = Array((new BigNumber(25)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = _354_v76;
        _nw55[(_dafny.ONE).toNumber()] = _354_v76;
        _nw55[(new BigNumber(2)).toNumber()] = _353_v75;
        _nw55[(new BigNumber(3)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(4)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(5)).toNumber()] = _353_v75;
        _nw55[(new BigNumber(6)).toNumber()] = _module.__default.fm72((_359_v81)[_module.__default.safeIndex(_314_v44, new BigNumber((_359_v81).length))], _279_v18, _258_globalState);
        _nw55[(new BigNumber(7)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(8)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(9)).toNumber()] = _module.__default.fm73(_258_globalState);
        _nw55[(new BigNumber(10)).toNumber()] = _353_v75;
        _nw55[(new BigNumber(11)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(12)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(13)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(14)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(15)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(16)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(17)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(18)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(19)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(20)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(790)), ((_361_v14) => function (_362_i9) {
          return _361_v14;
        })(_275_v14)));
        _nw55[(new BigNumber(21)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(22)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(23)).toNumber()] = _354_v76;
        _nw55[(new BigNumber(24)).toNumber()] = (_353_v75).update((_283_v22)[_module.__default.safeIndex(new BigNumber(766), new BigNumber((_283_v22).length))], _module.__default.abs((_260_v3)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_260_v3).length))]));
        _360_v82 = _nw55;
        let _index45 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_360_v82).length));
        (_360_v82)[_index45] = _354_v76;
        let _index46 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_360_v82).length));
        (_360_v82)[_index46] = _354_v76;
        let _rhs43 = _262_v4;
        let _rhs44 = !((_276_v15)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_276_v15).length))]);
        let _rhs45 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_266_v8, _266_v8)).length), _282_v21);
        _262_v4 = _rhs43;
        _281_v20 = _rhs44;
        _259_v2 = _rhs45;
      }
      process.stdout.write(_dafny.toString(_256_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_257_v1).equals(_dafny.MultiSet.fromElements(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_258_globalState).f0).equals(_dafny.MultiSet.fromElements(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_259_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_260_v3)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_262_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_263_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(357604),false).updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_264_v6).equals(_dafny.Set.fromElements(new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v7).equals(_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber(598))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_266_v8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_267_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber(598))),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v10)[_dafny.ZERO]).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_268_v10)[_dafny.ONE]).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v11)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_273_v12).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_274_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_275_v14, _dafny.Seq.of(new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v15)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_278_v17, _dafny.Seq.of(_module.D0.create_DC1(), _module.D0.create_DC1()))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_279_v18).equals(_dafny.MultiSet.fromElements(new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_280_v19).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(598)),_dafny.Seq.UnicodeFromString("smw")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_281_v20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_282_v21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[_dafny.ZERO], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[_dafny.ONE], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(3)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(4)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(5)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(6)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(7)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(8)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(9)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(10)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(11)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(12)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(13)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(14)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(15)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(16)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(17)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(18)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(19)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(20)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(21)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(22)], _dafny.Seq.of(new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_283_v22)[new BigNumber(23)], _dafny.Seq.of(new BigNumber(598), new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_286_v23).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v36).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_307_v40).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(598)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_308_v41)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(),new BigNumber(10)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_312_v42).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_313_v43));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_314_v44));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_315_v45).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC2(cf1) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1();
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
    static create_DC4(cf3, cf4) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC3(cf2) {
      let $dt = new D1(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC6() {
      let $dt = new D2(0);
      return $dt;
    }
    static create_DC7(cf6) {
      let $dt = new D2(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf5) {
      let $dt = new D2(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf5) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf5 === other.cf5;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6();
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
    static create_DC9(cf8, cf9, cf10) {
      let $dt = new D3(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC10(cf11) {
      let $dt = new D3(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC8(cf7) {
      let $dt = new D3(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC11(cf12) {
      let $dt = new D3(3);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(false, _dafny.ZERO, false);
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
    static create_DC13(cf14, cf15, cf16, cf17) {
      let $dt = new D4(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC12(cf13) {
      let $dt = new D4(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(_dafny.ZERO, false, _dafny.ZERO, _dafny.Seq.of());
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
    static create_DC14(cf18) {
      let $dt = new D5(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
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
    static create_DC16(cf20, cf21, cf22) {
      let $dt = new D6(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC15(cf19) {
      let $dt = new D6(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && this.cf21 === other.cf21 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(false, false, _dafny.Seq.of());
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
    static create_DC17(cf23) {
      let $dt = new D7(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23);
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC18(cf24) {
      let $dt = new D8(0);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf24) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC20(cf26, cf27, cf28) {
      let $dt = new D9(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC19(cf25) {
      let $dt = new D9(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC19" + "(" + this.cf25.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC20(false, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC22(cf30, cf31, cf32, cf33, cf34) {
      let $dt = new D10(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC21(cf29) {
      let $dt = new D10(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC22" + "(" + this.cf30.toVerbatimString(true) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31 && this.cf32 === other.cf32 && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC22(_dafny.Seq.UnicodeFromString(""), false, false, [], _dafny.ZERO);
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
    static create_DC24(cf36, cf37, cf38) {
      let $dt = new D11(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC25(cf39, cf40) {
      let $dt = new D11(1);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC26(cf41) {
      let $dt = new D11(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC23(cf35) {
      let $dt = new D11(3);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC23() { return this.$tag === 3; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC23" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf39 === other.cf39 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf41 === other.cf41;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf35 === other.cf35;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC24(false, false, _dafny.ZERO);
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
    static create_DC28(cf43, cf44, cf45, cf46, cf47) {
      let $dt = new D12(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC27(cf42) {
      let $dt = new D12(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44) && this.cf45 === other.cf45 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC28(_dafny.ZERO, _dafny.ZERO, false, false, _dafny.ZERO);
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
    get is_DC29() { return this.$tag === 0; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf48, other.cf48);
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
    static create_DC30(cf49) {
      let $dt = new D14(0);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC30" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf49, other.cf49);
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
    static create_DC31(cf50) {
      let $dt = new D15(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC31" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf50 === other.cf50;
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
    static create_DC33(cf52, cf53, cf54, cf55, cf56) {
      let $dt = new D16(0);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC32(cf51) {
      let $dt = new D16(1);
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
        return "D16.DC33" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC32" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf52 === other.cf52 && _dafny.areEqual(this.cf53, other.cf53) && _dafny.areEqual(this.cf54, other.cf54) && _dafny.areEqual(this.cf55, other.cf55) && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf51 === other.cf51;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC33(false, _dafny.ZERO, _dafny.MultiSet.Empty, _dafny.ZERO, _dafny.Set.Empty);
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
    static create_DC34(cf57) {
      let $dt = new D17(0);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC34" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf57, other.cf57);
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
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC36(cf59, cf60, cf61) {
      let $dt = new D18(0);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC35(cf58) {
      let $dt = new D18(1);
      $dt.cf58 = cf58;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf58() { return this.cf58; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC36" + "(" + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC35" + "(" + _dafny.toString(this.cf58) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf59, other.cf59) && _dafny.areEqual(this.cf60, other.cf60) && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf58, other.cf58);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC36(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC38(cf63, cf64) {
      let $dt = new D19(0);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC37(cf62) {
      let $dt = new D19(1);
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC39(cf65) {
      let $dt = new D19(2);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC39() { return this.$tag === 2; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC38" + "(" + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC37" + "(" + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC39" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf63 === other.cf63 && this.cf64 === other.cf64;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf65, other.cf65);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC38(false, false);
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
    static create_DC41(cf67) {
      let $dt = new D20(0);
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC40(cf66) {
      let $dt = new D20(1);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC41" + "(" + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC40" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC41(_dafny.ZERO);
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
    static create_DC43(cf69, cf70, cf71) {
      let $dt = new D21(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC44(cf72) {
      let $dt = new D21(1);
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC42(cf68) {
      let $dt = new D21(2);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get is_DC42() { return this.$tag === 2; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC43" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC44" + "(" + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC42" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf69 === other.cf69 && this.cf70 === other.cf70 && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf72 === other.cf72;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf68 === other.cf68;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC43(false, false, _dafny.Seq.of());
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
    static create_DC46(cf74, cf75, cf76) {
      let $dt = new D22(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC47(cf77, cf78, cf79, cf80, cf81) {
      let $dt = new D22(1);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC45(cf73) {
      let $dt = new D22(2);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC46" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC47" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC45" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf74, other.cf74) && _dafny.areEqual(this.cf75, other.cf75) && this.cf76 === other.cf76;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77) && this.cf78 === other.cf78 && this.cf79 === other.cf79 && _dafny.areEqual(this.cf80, other.cf80) && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf73 === other.cf73;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC46(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC49(cf83, cf84, cf85) {
      let $dt = new D23(0);
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC48(cf82) {
      let $dt = new D23(1);
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC50(cf86) {
      let $dt = new D23(2);
      $dt.cf86 = cf86;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get is_DC50() { return this.$tag === 2; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf86() { return this.cf86; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC49" + "(" + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC48" + "(" + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC50" + "(" + _dafny.toString(this.cf86) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf83, other.cf83) && _dafny.areEqual(this.cf84, other.cf84) && _dafny.areEqual(this.cf85, other.cf85);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf82 === other.cf82;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf86, other.cf86);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC49(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC51(cf87) {
      let $dt = new D24(0);
      $dt.cf87 = cf87;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get dtor_cf87() { return this.cf87; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC51" + "(" + _dafny.toString(this.cf87) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf87 === other.cf87;
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
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC52(cf88) {
      let $dt = new D25(0);
      $dt.cf88 = cf88;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get dtor_cf88() { return this.cf88; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC52" + "(" + _dafny.toString(this.cf88) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf88, other.cf88);
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
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC54(cf90, cf91, cf92, cf93) {
      let $dt = new D26(0);
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      return $dt;
    }
    static create_DC53(cf89) {
      let $dt = new D26(1);
      $dt.cf89 = cf89;
      return $dt;
    }
    static create_DC55(cf94) {
      let $dt = new D26(2);
      $dt.cf94 = cf94;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get is_DC53() { return this.$tag === 1; }
    get is_DC55() { return this.$tag === 2; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf94() { return this.cf94; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC54" + "(" + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC53" + "(" + _dafny.toString(this.cf89) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC55" + "(" + _dafny.toString(this.cf94) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf90, other.cf90) && _dafny.areEqual(this.cf91, other.cf91) && this.cf92 === other.cf92 && _dafny.areEqual(this.cf93, other.cf93);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf89, other.cf89);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf94, other.cf94);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC54(_dafny.ZERO, _dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC57(cf96) {
      let $dt = new D27(0);
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC56(cf95) {
      let $dt = new D27(1);
      $dt.cf95 = cf95;
      return $dt;
    }
    get is_DC57() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf95() { return this.cf95; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC57" + "(" + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC56" + "(" + _dafny.toString(this.cf95) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf95 === other.cf95;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC57(_dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
        }
      };
    }
  }

  $module.D28 = class D28 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC58(cf97) {
      let $dt = new D28(0);
      $dt.cf97 = cf97;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get dtor_cf97() { return this.cf97; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC58" + "(" + _dafny.toString(this.cf97) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf97, other.cf97);
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
          return D28.Default();
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
      this._f0 = _dafny.MultiSet.Empty;
      this._f1 = _dafny.ZERO;
      this._f2 = false;
      this._f3 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
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
    get f3() {
      let _this = this;
      return _this._f3;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f12 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f12) {
      let _this = this;
      (_this)._f12 = f12;
      return;
    }
    fm17(p0, globalState) {
      let _this = this;
      return true;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f5 = [];
      this._f6 = _dafny.ZERO;
      this._f4 = false;
      this.f16 = _dafny.ZERO;
      this._f15 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    set f5(value) {
      let _this = this;
      _this._f5 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f15, f16, f5, f6, f4) {
      let _this = this;
      (_this)._f15 = f15;
      (_this).f16 = f16;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f4 = f4;
      return;
    }
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f4;
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_this).f15, (_dafny.ZERO).minus((_this).f6));
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (((_this).f4) || (true)) {
        return (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(43)), function (_363_i0) {
          return (_this).f6;
        }))).Intersect(_dafny.MultiSet.fromElements((_this).f6, new BigNumber(469)));
      } else {
        return (_dafny.MultiSet.fromElements(_this.f16)).Union(_dafny.MultiSet.fromElements((_this).f6));
      }
    };
    fm7(globalState) {
      let _this = this;
      let _source9 = _module.D1.create_DC3(!((_this).f4));
      if (_source9.is_DC4) {
        let _364___mcc_h0 = (_source9).cf3;
        let _365___mcc_h1 = (_source9).cf4;
        let _366_cf4 = _365___mcc_h1;
        let _367_cf3 = _364___mcc_h0;
        return new BigNumber((_dafny.Seq.UnicodeFromString("iyylxfgh")).length);
      } else {
        let _368___mcc_h2 = (_source9).cf2;
        let _369_cf2 = _368___mcc_h2;
        return new BigNumber((((true) ? (_dafny.Seq.of(_369_cf2, !(_369_cf2))) : (_dafny.Seq.of(_369_cf2, _369_cf2)))).length);
      }
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _370_v0;
      _370_v0 = _dafny.Set.fromElements((_this).f15);
      let _371_v1;
      _371_v1 = _dafny.MultiSet.fromElements(_this.f16, new BigNumber((_370_v0).length), _this.f16);
      let _372_v2;
      _372_v2 = _module.D3.create_DC9((_this).f4, (_this).f6, !((_this).f4));
      let _source10 = (((_371_v1).IsProperSubsetOf(_dafny.MultiSet.fromElements((_this).f6))) ? (_372_v2) : (function (_pat_let5_0) {
        return function (_373_dt__update__tmp_h0) {
          return function (_pat_let6_0) {
            return function (_374_dt__update_hcf8_h0) {
              return _module.D3.create_DC9(_374_dt__update_hcf8_h0, (_373_dt__update__tmp_h0).dtor_cf9, (_373_dt__update__tmp_h0).dtor_cf10);
            }(_pat_let6_0);
          }(false);
        }(_pat_let5_0);
      }(_372_v2)));
      if (_source10.is_DC9) {
        let _375___mcc_h0 = (_source10).cf8;
        let _376___mcc_h1 = (_source10).cf9;
        let _377___mcc_h2 = (_source10).cf10;
        let _378_cf10 = _377___mcc_h2;
        let _379_cf9 = _376___mcc_h1;
        let _380_cf8 = _375___mcc_h0;
        let _381_v3;
        _381_v3 = new _dafny.CodePoint('k'.codePointAt(0));
        if ((new BigNumber(-828)).isLessThanOrEqualTo((new BigNumber((_module.__default.fm37(_380_cf8, !(_378_cf10), p0, _module.__default.fm0(_381_v3, globalState), globalState)).length)).minus(_379_cf9))) {
          let _382_v4;
          let _nw56 = Array((new BigNumber(4)).toNumber()).fill([]);
          _382_v4 = _nw56;
          let _383_v5;
          let _init4 = ((_384_p0) => function (_385_i0) {
            return (_385_i0).multipliedBy(_384_p0);
          })(p0);
          let _nw57 = Array((new BigNumber(22)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw57.length); _i0_4++) {
            _nw57[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _383_v5 = _nw57;
          let _index47 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_382_v4).length));
          (_382_v4)[_index47] = _383_v5;
          let _index48 = _module.__default.safeIndex(new BigNumber(500), new BigNumber((_382_v4).length));
          (_382_v4)[_index48] = _383_v5;
          let _386_v6;
          _386_v6 = _dafny.Seq.UnicodeFromString("tqnumrpc");
          _386_v6 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(435)), ((_387_v6, _388_cf9, _389_p0) => function (_390_i1) {
            return (_387_v6)[_module.__default.safeIndex(new BigNumber(((_module.D6.create_DC16((_this).f4, (_this).f4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), ((_391_cf9, _392_p0) => function (_393_i2) {
  return new BigNumber((_dafny.Set.fromElements(new BigNumber(732), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_391_cf9,_392_p0)).length))).length);
})(_388_cf9, _389_p0)))).dtor_cf22).length), new BigNumber((_387_v6).length))];
          })(_386_v6, _379_cf9, p0)), _module.__default.safeIndex(new BigNumber(834), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(435)), ((_394_v6, _395_cf9, _396_p0) => function (_397_i1) {
            return (_394_v6)[_module.__default.safeIndex(new BigNumber(((_module.D6.create_DC16((_this).f4, (_this).f4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), ((_398_cf9, _399_p0) => function (_400_i2) {
  return new BigNumber((_dafny.Set.fromElements(new BigNumber(732), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_398_cf9,_399_p0)).length))).length);
})(_395_cf9, _396_p0)))).dtor_cf22).length), new BigNumber((_394_v6).length))];
          })(_386_v6, _379_cf9, p0))).length)), _381_v3);
          let _401_v7;
          _401_v7 = _dafny.Seq.of(_379_cf9);
          let _402_v8;
          _402_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm8((_this).fm7(globalState), (_this).f4, _386_v6, _401_v7, globalState),new BigNumber((_401_v7).length));
          r0 = _402_v8;
          _380_cf8 = _378_cf10;
          let _403_v9;
          _403_v9 = _dafny.Seq.of((_this).f4);
          _403_v9 = _dafny.Seq.Concat(_dafny.Seq.of(_380_cf8), _403_v9);
        } else {
          let _404_v11;
          _404_v11 = _dafny.Seq.of((_this).f4, _380_cf8);
          let _405_v12;
          _405_v12 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll28 = new _dafny.Map();
            for (const _compr_28 of (_dafny.MultiSet.fromElements((_this).f15)).Elements) {
              let _406_v10 = _compr_28;
              if ((_dafny.MultiSet.fromElements((_this).f15)).contains(_406_v10)) {
                _coll28.push([_module.__default.safeModuloInt(_406_v10, p0),(_this).f4]);
              }
            }
            return _coll28;
          }(),_module.__default.fm3(_379_cf9, (_404_v11)[_module.__default.safeIndex(_379_cf9, new BigNumber((_404_v11).length))], globalState));
          _405_v12 = (_405_v12).Merge(_405_v12);
          let _407_v13;
          let _nw58 = new _module.C0();
          _nw58.__ctor(_379_cf9);
          _407_v13 = _nw58;
          (_this).f16 = _module.__default.safeDivisionInt((_this).f15, (_407_v13).f12);
          let _408_v14;
          _408_v14 = _dafny.Seq.of((_dafny.ZERO).minus(p0), new BigNumber(886));
          (_this).f16 = ((_408_v14)[_module.__default.safeIndex(_379_cf9, new BigNumber((_408_v14).length))]).minus(new BigNumber(760));
          let _409_v15;
          _409_v15 = _dafny.MultiSet.fromElements(!(!(_380_cf8)));
          let _410_v16;
          _410_v16 = _dafny.Seq.UnicodeFromString("tplcf");
          (_this).f16 = (_379_cf9).minus(new BigNumber((_dafny.Seq.Concat(_module.__default.fm37(_378_cf10, (_this).f4, (((_409_v15).contains((_this).f4)) ? ((_409_v15).get((_this).f4)) : (new BigNumber(150))), (_this).f4, globalState), _dafny.Seq.update(_410_v16, _module.__default.safeIndex(_this.f16, new BigNumber((_410_v16).length)), _381_v3))).length));
        }
        _380_cf8 = _380_cf8;
        (_this).f5 = _this.f5;
        if (_378_cf10) {
          _380_cf8 = !(_380_cf8);
          (_this).f16 = _module.__default.safeModuloInt(p0, (_this).f6);
          _379_cf9 = (_379_cf9).minus(new BigNumber(905));
          let _411_v17;
          let _nw59 = new _module.C0();
          _nw59.__ctor((_379_cf9).minus(_379_cf9));
          _411_v17 = _nw59;
          let _nw60 = new _module.C0();
          _nw60.__ctor(new BigNumber(921));
          _411_v17 = _nw60;
          let _412_v18;
          _412_v18 = _dafny.Seq.of((_this).f4, true, !(_378_cf10));
          let _413_v19;
          _413_v19 = _dafny.Seq.of(((_this).f6).multipliedBy(new BigNumber(((_module.D12.create_DC27(_dafny.Seq.of(_412_v18))).dtor_cf42).length)));
          _413_v19 = _413_v19;
        } else {
          (_this).f16 = (_dafny.ZERO).minus(_379_cf9);
          let _414_v20;
          let _init5 = ((_415_cf9) => function (_416_i3) {
            return _dafny.Seq.of(_415_cf9, new BigNumber((_dafny.Seq.UnicodeFromString("lyl")).length));
          })(_379_cf9);
          let _nw61 = Array((new BigNumber(8)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw61.length); _i0_5++) {
            _nw61[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _414_v20 = _nw61;
          let _417_v21;
          _417_v21 = _module.D6.create_DC16(_378_cf10, _378_cf10, _dafny.Seq.of(p0));
          let _index49 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_414_v20).length));
          (_414_v20)[_index49] = (_417_v21).dtor_cf22;
          let _418_v22;
          _418_v22 = _dafny.Seq.of((_this).f6);
          let _index50 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_414_v20).length));
          (_414_v20)[_index50] = _418_v22;
          let _419_v23;
          _419_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f16);
          let _420_v24;
          _420_v24 = _dafny.Seq.UnicodeFromString("tlcjkdars");
          _378_cf10 = (((_this).fm8((((_419_v23).contains(_380_cf8)) ? ((_419_v23).get(_380_cf8)) : (new BigNumber((_420_v24).length))), true, _420_v24, _418_v22, globalState)) ? (_378_cf10) : (!((_this).f4)));
          let _421_v25;
          let _nw62 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _421_v25 = _nw62;
          _421_v25 = _421_v25;
          _381_v3 = (((_this).f4) ? (_381_v3) : (_381_v3));
        }
      } else if (_source10.is_DC10) {
        let _422___mcc_h3 = (_source10).cf11;
        let _423_cf11 = _422___mcc_h3;
        if ((_this).f4) {
          let _424_v26;
          let _nw63 = new _module.C0();
          _nw63.__ctor((_this).fm7(globalState));
          _424_v26 = _nw63;
          let _425_v27;
          _425_v27 = _dafny.Seq.UnicodeFromString("nfgsjgori");
          let _426_v28;
          _426_v28 = _module.D9.create_DC19(_425_v27);
          let _pat_let_tv5 = _425_v27;
          let _427_v29;
          _427_v29 = _dafny.MultiSet.fromElements(function (_pat_let7_0) {
            return function (_428_dt__update__tmp_h1) {
              return function (_pat_let8_0) {
                return function (_429_dt__update_hcf25_h0) {
                  return _module.D9.create_DC19(_429_dt__update_hcf25_h0);
                }(_pat_let8_0);
              }(_pat_let_tv5);
            }(_pat_let7_0);
          }(_426_v28), _426_v28);
          let _430_v30;
          _430_v30 = _dafny.Set.fromElements(true, (_this).f4);
          let _431_v31;
          _431_v31 = _dafny.Seq.of(new BigNumber((_430_v30).length));
          let _432_v32;
          _432_v32 = _dafny.Map.Empty.slice().updateUnsafe(_427_v29,_431_v31);
          let _433_v34;
          _433_v34 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(_426_v28)), _dafny.MultiSet.fromElements(_426_v28), _427_v29, _427_v29);
          let _434_v35;
          _434_v35 = _dafny.Seq.of((_this).f4);
          let _435_v36;
          let _nw64 = Array((new BigNumber(11)).toNumber()).fill(false);
          _435_v36 = _nw64;
          let _436_v37;
          _436_v37 = _dafny.Map.Empty.slice().updateUnsafe(_435_v36,_423_cf11);
          let _437_v38;
          let _nw65 = Array((new BigNumber(14)).toNumber());
          _nw65[(_dafny.ZERO).toNumber()] = _this.f16;
          _nw65[(_dafny.ONE).toNumber()] = new BigNumber(((_432_v32).Merge(function () {
            let _coll29 = new _dafny.Map();
            for (const _compr_29 of (_433_v34).Elements) {
              let _438_v33 = _compr_29;
              if ((_433_v34).contains(_438_v33)) {
                _coll29.push([_438_v33,_dafny.Seq.of((_424_v26).f12, new BigNumber(-615))]);
              }
            }
            return _coll29;
          }())).length);
          _nw65[(new BigNumber(2)).toNumber()] = p0;
          _nw65[(new BigNumber(3)).toNumber()] = _423_cf11;
          _nw65[(new BigNumber(4)).toNumber()] = (_this).f6;
          _nw65[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_434_v35).length), p0);
          _nw65[(new BigNumber(6)).toNumber()] = (_this).f6;
          _nw65[(new BigNumber(7)).toNumber()] = (((_this).fm8((_this).f6, (_434_v35)[_module.__default.safeIndex(_module.__default.fm3(new BigNumber(130), false, globalState), new BigNumber((_434_v35).length))], _425_v27, _431_v31, globalState)) ? ((_this).f6) : (new BigNumber((_431_v31).length)));
          _nw65[(new BigNumber(8)).toNumber()] = _this.f16;
          _nw65[(new BigNumber(9)).toNumber()] = p0;
          _nw65[(new BigNumber(10)).toNumber()] = _this.f16;
          _nw65[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((((_436_v37).contains(_435_v36)) ? ((_436_v37).get(_435_v36)) : (new BigNumber((_434_v35).length))));
          _nw65[(new BigNumber(12)).toNumber()] = (_424_v26).f12;
          _nw65[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(_423_cf11, p0);
          _437_v38 = _nw65;
          let _index51 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_437_v38).length));
          (_437_v38)[_index51] = _module.__default.safeDivisionInt(new BigNumber(-342), _module.__default.fm3((_this).f6, (_this).f4, globalState));
          let _index52 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_437_v38).length));
          (_437_v38)[_index52] = _module.__default.fm3(((_this).f6).minus((_424_v26).f12), (_this).f4, globalState);
          let _439_v39;
          _439_v39 = _dafny.Seq.of(_430_v30);
          let _440_v40;
          _440_v40 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_434_v35, _434_v35),_439_v39);
          let _441_v41;
          _441_v41 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f15),_430_v30);
          _440_v40 = (_440_v40).update(_434_v35, _dafny.Seq.of((((_441_v41).contains(new BigNumber(372))) ? ((_441_v41).get(new BigNumber(372))) : (_430_v30))));
          let _442_v42;
          _442_v42 = true;
          let _443_v43;
          _443_v43 = _dafny.MultiSet.fromElements(_442_v42, _442_v42, (_430_v30).IsProperSubsetOf(_dafny.Set.fromElements((_this).f4, (_this).f4, _442_v42, (_this).fm8((_437_v38)[_module.__default.safeIndex(new BigNumber(159), new BigNumber((_437_v38).length))], (_this).f4, _425_v27, _431_v31, globalState), _442_v42)));
          let _rhs46 = (_424_v26).f12;
          let _rhs47 = false;
          let _rhs48 = function () {
            let _coll30 = new _dafny.Set();
            for (const _compr_30 of _dafny.IntegerRange(new BigNumber(681), new BigNumber(893))) {
              let _444_v44 = _compr_30;
              if (((new BigNumber(681)).isLessThanOrEqualTo(_444_v44)) && ((_444_v44).isLessThan(new BigNumber(893)))) {
                _coll30.add((_444_v44).multipliedBy((_424_v26).f12));
              }
            }
            return _coll30;
          }();
          let _rhs49 = (_dafny.MultiSet.fromElements((_this).f4, false)).update(!((_this).f4) || (false), _module.__default.abs((new BigNumber((_425_v27).length)).plus((_424_v26).f12)));
          let _lhs28 = _this;
          _lhs28.f16 = _rhs46;
          _442_v42 = _rhs47;
          _370_v0 = _rhs48;
          _443_v43 = _rhs49;
          let _445_v45;
          _445_v45 = _370_v0;
          let _446_v46;
          _446_v46 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(710)), function (_447_i4) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }));
          let _448_v47;
          _448_v47 = _dafny.Map.Empty.slice().updateUnsafe(_445_v45,(_446_v46).IsDisjointFrom(_446_v46));
          _448_v47 = (_448_v47).update(_445_v45, (_this).f4);
        } else {
          let _449_v48;
          _449_v48 = _dafny.Seq.UnicodeFromString("qrngovhln");
          _449_v48 = _dafny.Seq.Concat(_449_v48, _449_v48);
          let _450_v49;
          _450_v49 = _dafny.MultiSet.fromElements(_449_v48, _449_v48, _449_v48);
          _423_cf11 = _module.__default.safeModuloInt(_module.__default.fm3((_this).f6, (_this).f4, globalState), (((_450_v49).contains(_449_v48)) ? ((_450_v49).get(_449_v48)) : ((_this).f15)));
          let _451_v50;
          let _nw66 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _451_v50 = _nw66;
          let _index53 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_451_v50).length));
          (_451_v50)[_index53] = (_dafny.ZERO).minus(new BigNumber((_370_v0).length));
          let _452_v51;
          _452_v51 = _dafny.Set.fromElements((_this).f4);
          let _index54 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_451_v50).length));
          (_451_v50)[_index54] = _module.__default.safeModuloInt((_this).fm7(globalState), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_452_v51)).length)));
          _423_cf11 = (_451_v50)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_451_v50).length))];
          let _453_v52;
          let _init6 = function (_454_i5) {
            return (_dafny.MultiSet.fromElements(!((_this).f4))).Intersect(_dafny.MultiSet.fromElements((_this).f4, (_this).f4));
          };
          let _nw67 = Array((new BigNumber(9)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw67.length); _i0_6++) {
            _nw67[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _453_v52 = _nw67;
          let _455_v53;
          _455_v53 = _dafny.MultiSet.fromElements((_this).f4);
          let _456_v54;
          _456_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_455_v53);
          let _index55 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_453_v52).length));
          (_453_v52)[_index55] = ((!((_this).f4)) ? (_dafny.MultiSet.fromElements((_this).f4)) : ((((_456_v54).contains((_this).f4)) ? ((_456_v54).get((_this).f4)) : (_455_v53))));
          let _index56 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_453_v52).length));
          (_453_v52)[_index56] = (_455_v53).Difference((_dafny.MultiSet.fromElements((_this).f4)).Difference(_455_v53));
        }
        let _457_v55;
        _457_v55 = _dafny.Seq.UnicodeFromString("dohmu");
        let _458_v56;
        _458_v56 = _dafny.Map.Empty.slice().updateUnsafe(_457_v55,(_this).f6);
        let _459_v57;
        _459_v57 = _dafny.Seq.of(_458_v56);
        let _460_v58;
        let _nw68 = Array((new BigNumber(26)).toNumber()).fill(false);
        _460_v58 = _nw68;
        let _461_v59;
        let _out12;
        _out12 = (_this).m11((_this).f4, !((_this).f4), ((_459_v57)[_module.__default.safeIndex((_this).f6, new BigNumber((_459_v57).length))]).update(_457_v55, _423_cf11), _460_v58, globalState);
        _461_v59 = _out12;
        let _462_v60;
        let _init7 = ((_463_v55) => function (_464_i6) {
          return _dafny.Seq.Concat(_463_v55, _463_v55);
        })(_457_v55);
        let _nw69 = Array((new BigNumber(25)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw69.length); _i0_7++) {
          _nw69[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _462_v60 = _nw69;
        let _index57 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_462_v60).length));
        (_462_v60)[_index57] = _457_v55;
        let _465_v61;
        _465_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_423_cf11);
        let _466_v62;
        _466_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-87),_465_v61);
        let _index58 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_462_v60).length));
        (_462_v60)[_index58] = _module.__default.fm37((_this).f4, (_this).f4, p0, !((((_466_v62).contains(p0)) ? ((_466_v62).get(p0)) : ((_465_v61).update((_this).fm9((_this).f4, true, _423_cf11, globalState), (_this).f15)))).contains((_this).f15), globalState);
        let _467_v63;
        let _init8 = function (_468_i7) {
          return (_dafny.MultiSet.fromElements(_module.D4.create_DC12(new _dafny.CodePoint('f'.codePointAt(0))))).Intersect(_dafny.MultiSet.fromElements(_module.D4.create_DC12(new _dafny.CodePoint('l'.codePointAt(0)))));
        };
        let _nw70 = Array((new BigNumber(21)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw70.length); _i0_8++) {
          _nw70[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _467_v63 = _nw70;
        _467_v63 = _467_v63;
      } else if (_source10.is_DC8) {
        let _469___mcc_h4 = (_source10).cf7;
        let _470_cf7 = _469___mcc_h4;
        let _471_v64;
        let _nw71 = Array((new BigNumber(28)).toNumber()).fill(false);
        _471_v64 = _nw71;
        let _index59 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_471_v64).length));
        (_471_v64)[_index59] = (new BigNumber(548)).isLessThan(_this.f16);
        let _index60 = _module.__default.safeIndex(new BigNumber(707), new BigNumber((_471_v64).length));
        (_471_v64)[_index60] = (_this).f4;
        let _472_v65;
        _472_v65 = _dafny.Seq.UnicodeFromString("dv");
        _472_v65 = _dafny.Seq.UnicodeFromString("qssuchryd");
        let _473_v66;
        let _nw72 = new _module.C0();
        _nw72.__ctor(p0);
        _473_v66 = _nw72;
        let _474_v67;
        let _nw73 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _474_v67 = _nw73;
        let _475_v68;
        _475_v68 = _dafny.Set.fromElements(_474_v67, _474_v67);
        let _476_v69;
        _476_v69 = _475_v68;
        let _477_v70;
        _477_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_476_v69));
        let _478_v71;
        _478_v71 = _dafny.Seq.of(_475_v68);
        _477_v70 = (_477_v70).update(!((_this).f4), (((_this).f4) ? (_475_v68) : ((_478_v71)[_module.__default.safeIndex((_this).f6, new BigNumber((_478_v71).length))])));
      } else {
        let _479___mcc_h5 = (_source10).cf12;
        let _480_cf12 = _479___mcc_h5;
        let _source11 = _module.D11.create_DC24(!((_this).f4), (_this).f4, _module.__default.safeModuloInt(_this.f16, new BigNumber(-625)));
        if (_source11.is_DC24) {
          let _481___mcc_h6 = (_source11).cf36;
          let _482___mcc_h7 = (_source11).cf37;
          let _483___mcc_h8 = (_source11).cf38;
          let _484_cf38 = _483___mcc_h8;
          let _485_cf37 = _482___mcc_h7;
          let _486_cf36 = _481___mcc_h6;
          (_this).f16 = _484_cf38;
          _486_cf36 = _486_cf36;
          let _487_v72;
          let _nw74 = Array((new BigNumber(11)).toNumber()).fill(false);
          _487_v72 = _nw74;
          let _488_v73;
          _488_v73 = _module.D2.create_DC5(_487_v72);
          _488_v73 = _488_v73;
          let _489_v74;
          let _nw75 = Array((new BigNumber(2)).toNumber());
          _nw75[(_dafny.ZERO).toNumber()] = new BigNumber(445);
          _nw75[(_dafny.ONE).toNumber()] = (_this).f15;
          _489_v74 = _nw75;
          let _490_v75;
          _490_v75 = _dafny.Seq.of(_489_v74);
          _490_v75 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_489_v74), _dafny.Seq.update(_490_v75, _module.__default.safeIndex((_this).f6, new BigNumber((_490_v75).length)), _489_v74)), _dafny.Seq.Concat(_490_v75, _490_v75));
        } else if (_source11.is_DC25) {
          let _491___mcc_h9 = (_source11).cf39;
          let _492___mcc_h10 = (_source11).cf40;
          let _493_cf40 = _492___mcc_h10;
          let _494_cf39 = _491___mcc_h9;
          let _495_v76;
          _495_v76 = new _dafny.CodePoint('b'.codePointAt(0));
          let _496_v77;
          _496_v77 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("odwhwaejb"),new BigNumber((_dafny.MultiSet.fromElements(_494_cf39)).cardinality()));
          let _497_v79;
          _497_v79 = _dafny.Seq.UnicodeFromString("dqtrskur");
          let _498_v80;
          _498_v80 = _dafny.Set.fromElements(_497_v79);
          let _499_v81;
          let _init9 = function (_500_i9) {
            return !((_this).f4);
          };
          let _nw76 = Array((new BigNumber(20)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw76.length); _i0_9++) {
            _nw76[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _499_v81 = _nw76;
          let _501_v82;
          let _out13;
          _out13 = (_this).m11(_dafny.Seq.IsProperPrefixOf(_module.__default.fm37(_494_cf39, (_this).f4, (_this).f6, false, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(906)), function (_502_i8) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          })), _module.__default.fm0(_495_v76, globalState), (_496_v77).Merge(function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of (_498_v80).Elements) {
              let _503_v78 = _compr_31;
              if ((_498_v80).contains(_503_v78)) {
                _coll31.push([_503_v78,new BigNumber((_dafny.Set.fromElements((_this).f4, _494_cf39, _494_cf39)).length)]);
              }
            }
            return _coll31;
          }()), _499_v81, globalState);
          _501_v82 = _out13;
          _497_v79 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(812)), ((_504_v76) => function (_505_i10) {
            return _504_v76;
          })(_495_v76));
          let _506_v83;
          _506_v83 = _dafny.Seq.of(_this.f16);
          let _507_v84;
          _507_v84 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_506_v83),_497_v79);
          let _508_v85;
          let _509_v86;
          let _out14;
          let _out15;
          let _outcollector6 = _module.__default.m0(_507_v84, _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("lcxrsang"), _497_v79), (_371_v1).Intersect(_371_v1), !_dafny.Seq.contains(_dafny.Seq.update(_497_v79, _module.__default.safeIndex(p0, new BigNumber((_497_v79).length)), _495_v76), _495_v76), globalState);
          _out14 = _outcollector6[0];
          _out15 = _outcollector6[1];
          _508_v85 = _out14;
          _509_v86 = _out15;
          let _510_v87;
          _510_v87 = _dafny.Seq.of(_494_cf39);
          let _511_v88;
          _511_v88 = _dafny.Map.Empty.slice().updateUnsafe(_509_v86,_495_v76);
          let _512_v89;
          _512_v89 = _dafny.Map.Empty.slice().updateUnsafe(_509_v86,_493_cf40);
          let _513_v90;
          _513_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_512_v89).length),_493_cf40);
          let _514_v91;
          _514_v91 = _dafny.Seq.of(_module.__default.fm40((_this).f4, _508_v85, _509_v86, globalState), _510_v87, _510_v87);
          _508_v85 = !(((_371_v1).Intersect(_dafny.MultiSet.FromArray(_module.__default.fm39(_497_v79, _module.D12.create_DC27(_514_v91), _494_cf39, globalState)))).IsProperSubsetOf((_371_v1).Difference((_this).fm6(_module.__default.fm38((_this).f15, p0, _508_v85, new BigNumber((_510_v87).length), globalState), _module.D0.create_DC0((_this).f15), _511_v88, new BigNumber((_513_v90).length), globalState))));
        } else if (_source11.is_DC26) {
          let _515___mcc_h11 = (_source11).cf41;
          let _516_cf41 = _515___mcc_h11;
          (_this).f16 = _module.__default.safeDivisionInt((_this).f6, (_this).f15);
          let _517_v92;
          let _nw77 = Array((new BigNumber(18)).toNumber()).fill(false);
          _517_v92 = _nw77;
          let _index61 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_517_v92).length));
          (_517_v92)[_index61] = (_this.f16).isLessThanOrEqualTo(p0);
          let _518_v93;
          _518_v93 = _dafny.Seq.of(new BigNumber(106), (_this).f15);
          let _index62 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_517_v92).length));
          (_517_v92)[_index62] = _dafny.areEqual(_518_v93, _518_v93);
          let _index63 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_517_v92).length));
          (_517_v92)[_index63] = ((_517_v92)[_module.__default.safeIndex(new BigNumber(725), new BigNumber((_517_v92).length))]) === ((_517_v92)[_module.__default.safeIndex(new BigNumber(725), new BigNumber((_517_v92).length))]);
          let _519_v94;
          _519_v94 = _module.D0.create_DC0((_this).f15);
          let _520_v95;
          _520_v95 = _dafny.Seq.UnicodeFromString("r");
          let _521_v96;
          _521_v96 = _dafny.Map.Empty.slice().updateUnsafe(p0,_520_v95);
          let _522_v97;
          _522_v97 = _dafny.Seq.of(_520_v95, (((_521_v96).contains(new BigNumber(-706))) ? ((_521_v96).get(new BigNumber(-706))) : (_520_v95)));
          let _523_v98;
          _523_v98 = new _dafny.CodePoint('y'.codePointAt(0));
          let _pat_let_tv6 = p0;
          let _pat_let_tv7 = p0;
          let _index64 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_517_v92).length));
          let _rhs50 = function (_pat_let9_0) {
            return function (_524_dt__update__tmp_h2) {
              return function (_pat_let10_0) {
                return function (_525_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_525_dt__update_hcf0_h0);
                }(_pat_let10_0);
              }((_pat_let_tv6).plus(_pat_let_tv7));
            }(_pat_let9_0);
          }(_519_v94);
          let _rhs51 = _dafny.Seq.Concat(_522_v97, _522_v97);
          let _rhs52 = (_this).f4;
          let _rhs53 = _523_v98;
          let _lhs29 = _517_v92;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_517_v92).length));
          _519_v94 = _rhs50;
          _522_v97 = _rhs51;
          _lhs29[_lhs30] = _rhs52;
          _523_v98 = _rhs53;
        } else {
          let _526___mcc_h12 = (_source11).cf35;
          let _527_cf35 = _526___mcc_h12;
          let _528_v99;
          let _nw78 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _528_v99 = _nw78;
          let _529_v100;
          _529_v100 = _dafny.Seq.UnicodeFromString("yfnegvsi");
          let _530_v101;
          let _init10 = function (_531_i11) {
            return (_531_i11).minus(_this.f16);
          };
          let _nw79 = Array((new BigNumber(19)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw79.length); _i0_10++) {
            _nw79[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _530_v101 = _nw79;
          let _532_v102;
          _532_v102 = _dafny.Map.Empty.slice().updateUnsafe(_529_v100,_530_v101);
          let _index65 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_528_v99).length));
          (_528_v99)[_index65] = (_532_v102).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(252)), function (_533_i12) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }), _530_v101);
          let _index66 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_528_v99).length));
          (_528_v99)[_index66] = _532_v102;
          let _534_v103;
          _534_v103 = _dafny.Seq.of(_this.f16, new BigNumber(357));
          _529_v100 = _module.__default.fm37((_this).f4, !(p0).isEqualTo(new BigNumber((_534_v103).length)), (_dafny.ZERO).minus((_this).f6), (_this).f4, globalState);
          let _535_v104;
          let _nw80 = Array((new BigNumber(28)).toNumber()).fill(false);
          _535_v104 = _nw80;
          let _index67 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_535_v104).length));
          (_535_v104)[_index67] = (_this).f4;
          let _index68 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_535_v104).length));
          (_535_v104)[_index68] = ((_this).f4) && ((_this).f4);
          let _536_v105;
          _536_v105 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3((_this).f15, (_this).f4, globalState),(_535_v104)[_module.__default.safeIndex(new BigNumber(205), new BigNumber((_535_v104).length))]);
          let _537_v106;
          _537_v106 = _dafny.Map.Empty.slice().updateUnsafe(_535_v104,(_this).f15);
          let _538_v107;
          _538_v107 = new _dafny.CodePoint('f'.codePointAt(0));
          let _539_v108;
          _539_v108 = _dafny.Map.Empty.slice().updateUnsafe(_538_v107,(_this).f4);
          let _540_v109;
          _540_v109 = _module.D10.create_DC22(_529_v100, (_this).f4, (((_539_v108).contains(_538_v107)) ? ((_539_v108).get(_538_v107)) : (!((_this).f4))), _535_v104, p0);
          _536_v105 = (_536_v105).update((((_537_v106).contains(_535_v104)) ? ((_537_v106).get(_535_v104)) : ((((_371_v1).contains(_this.f16)) ? ((_371_v1).get(_this.f16)) : ((_this).f15)))), (_540_v109).dtor_cf32);
        }
        let _541_v110;
        let _nw81 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _541_v110 = _nw81;
        let _index69 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_541_v110).length));
        (_541_v110)[_index69] = _module.__default.safeModuloInt(new BigNumber(739), (_this).f15);
        let _index70 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_541_v110).length));
        (_541_v110)[_index70] = (_this).f6;
        (_this).f16 = ((_this).f15).plus((_dafny.ZERO).minus(p0));
        let _542_v111;
        let _nw82 = new _module.C0();
        _nw82.__ctor((_this).f15);
        _542_v111 = _nw82;
      }
      let _543_v112;
      _543_v112 = _dafny.Seq.UnicodeFromString("jmhajedk");
      let _544_v113;
      _544_v113 = _dafny.Set.fromElements(_543_v112);
      let _545_i13;
      _545_i13 = _dafny.ZERO;
      L0: {
        while (!((_544_v113).IsProperSubsetOf((_544_v113).Intersect(_544_v113)))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_545_i13)) {
              break L0;
            }
            _545_i13 = (_545_i13).plus(_dafny.ONE);
            (_this).f16 = new BigNumber((_dafny.Seq.Concat(_543_v112, _543_v112)).length);
            let _546_v114;
            let _nw83 = new _module.C0();
            _nw83.__ctor(_this.f16);
            _546_v114 = _nw83;
            let _547_v115;
            _547_v115 = new _dafny.CodePoint('n'.codePointAt(0));
            let _548_v116;
            _548_v116 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_547_v115);
            let _549_v117;
            _549_v117 = _dafny.Seq.of(((_this).f15).plus(new BigNumber((_548_v116).length)), (_546_v114).f12, (_this).f15);
            let _550_v118;
            let _nw84 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
            _550_v118 = _nw84;
            let _index71 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_550_v118).length));
            (_550_v118)[_index71] = (_this).f15;
            let _551_v119;
            _551_v119 = _dafny.Set.fromElements((_this).f4, (_this).f4);
            let _552_v120;
            _552_v120 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_551_v119);
            let _553_v121;
            _553_v121 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f4);
            let _index72 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_550_v118).length));
            let _rhs54 = _dafny.Seq.Concat(_549_v117, _dafny.Seq.of(new BigNumber((_552_v120).length), new BigNumber(798)));
            let _rhs55 = (_546_v114).f12;
            let _rhs56 = ((_this).f6).multipliedBy(new BigNumber((_553_v121).length));
            let _lhs31 = _this;
            let _lhs32 = _550_v118;
            let _lhs33 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_550_v118).length));
            _549_v117 = _rhs54;
            _lhs31.f16 = _rhs55;
            _lhs32[_lhs33] = _rhs56;
            let _554_v122;
            let _nw85 = Array((new BigNumber(10)).toNumber()).fill(false);
            _554_v122 = _nw85;
            let _index73 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_554_v122).length));
            (_554_v122)[_index73] = (((_this).f4) ? (true) : ((_this).f4));
            let _index74 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_550_v118).length));
            let _index75 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_554_v122).length));
            let _rhs57 = (_dafny.ZERO).minus(((_this).f6).multipliedBy((_550_v118)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_550_v118).length))]));
            let _rhs58 = (_this).f4;
            let _lhs34 = _550_v118;
            let _lhs35 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_550_v118).length));
            let _lhs36 = _554_v122;
            let _lhs37 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_554_v122).length));
            _lhs34[_lhs35] = _rhs57;
            _lhs36[_lhs37] = _rhs58;
          }
        }
      }
      let _555_v123;
      _555_v123 = new _dafny.CodePoint('a'.codePointAt(0));
      let _556_v124;
      _556_v124 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f4);
      let _557_v125;
      _557_v125 = _dafny.MultiSet.fromElements((_this).f4);
      let _558_v126;
      _558_v126 = _module.D12.create_DC28(new BigNumber((_module.__default.fm40((_this).f4, (((_556_v124).contains((((_557_v125).contains((_this).f4)) ? ((_557_v125).get((_this).f4)) : ((_this).f6)))) ? ((_556_v124).get((((_557_v125).contains((_this).f4)) ? ((_557_v125).get((_this).f4)) : ((_this).f6)))) : ((_this).f4)), _this.f16, globalState)).length), (_dafny.ZERO).minus((_this).f6), (_this).f4, (_this).f4, (_this).f15);
      let _pat_let_tv8 = _555_v123;
      let _pat_let_tv9 = _555_v123;
      _555_v123 = function (_source12) {
        if (_source12.is_DC28) {
          let _559___mcc_h13 = (_source12).cf43;
          let _560___mcc_h14 = (_source12).cf44;
          let _561___mcc_h15 = (_source12).cf45;
          let _562___mcc_h16 = (_source12).cf46;
          let _563___mcc_h17 = (_source12).cf47;
          let _564_cf47 = _563___mcc_h17;
          let _565_cf46 = _562___mcc_h16;
          let _566_cf45 = _561___mcc_h15;
          let _567_cf44 = _560___mcc_h14;
          let _568_cf43 = _559___mcc_h13;
          return _pat_let_tv8;
        } else {
          let _569___mcc_h18 = (_source12).cf42;
          let _570_cf42 = _569___mcc_h18;
          return _pat_let_tv9;
        }
      }(_558_v126);
      let _571_v127;
      _571_v127 = _dafny.Seq.of((_this).f15);
      let _572_v128;
      _572_v128 = _dafny.Set.fromElements(_module.D6.create_DC16((_this).f4, (_this).f4, _571_v127), _module.__default.fm43(globalState));
      let _573_v129;
      let _nw86 = Array((new BigNumber(20)).toNumber());
      _nw86[(_dafny.ZERO).toNumber()] = (_this).f4;
      _nw86[(_dafny.ONE).toNumber()] = !((((_556_v124).contains((_this).f6)) ? ((_556_v124).get((_this).f6)) : ((_this).f4)));
      _nw86[(new BigNumber(2)).toNumber()] = !(false);
      _nw86[(new BigNumber(3)).toNumber()] = (p0).isLessThan((_dafny.ZERO).minus((_this).f15));
      _nw86[(new BigNumber(4)).toNumber()] = (_this).f4;
      _nw86[(new BigNumber(5)).toNumber()] = (_this).f4;
      _nw86[(new BigNumber(6)).toNumber()] = (_this).f4;
      _nw86[(new BigNumber(7)).toNumber()] = !((_this).fm8((_this).f15, (_this).f4, _543_v112, _571_v127, globalState));
      _nw86[(new BigNumber(8)).toNumber()] = (_this).f4;
      _nw86[(new BigNumber(9)).toNumber()] = (p0).isLessThanOrEqualTo((_this).f15);
      _nw86[(new BigNumber(10)).toNumber()] = false;
      _nw86[(new BigNumber(11)).toNumber()] = (_this).f4;
      _nw86[(new BigNumber(12)).toNumber()] = (_module.__default.fm41(new BigNumber((_543_v112).length), _555_v123, (_this).f4, globalState)).IsSubsetOf(_module.__default.fm41(new BigNumber(440), _555_v123, (_this).f4, globalState));
      _nw86[(new BigNumber(13)).toNumber()] = _module.__default.fm0(_555_v123, globalState);
      _nw86[(new BigNumber(14)).toNumber()] = (_this).f4;
      _nw86[(new BigNumber(15)).toNumber()] = false;
      _nw86[(new BigNumber(16)).toNumber()] = !((_this).f4) || ((_this).f4);
      _nw86[(new BigNumber(17)).toNumber()] = true;
      _nw86[(new BigNumber(18)).toNumber()] = (_module.__default.fm42(_555_v123, _572_v128, globalState)).contains(p0);
      _nw86[(new BigNumber(19)).toNumber()] = (_this).f4;
      _573_v129 = _nw86;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_573_v129).length))) {
        let _574_i14 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_574_i14)) && ((_574_i14).isLessThan(new BigNumber((_573_v129).length))))) {
          (_573_v129)[(_574_i14)] = (_this).f4;
        }
      }
      let _575_i15;
      _575_i15 = _dafny.ZERO;
      L1: {
        while ((_this).f4) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_575_i15)) {
              break L1;
            }
            _575_i15 = (_575_i15).plus(_dafny.ONE);
            (_this).f16 = (_dafny.ZERO).minus((_this.f16).multipliedBy(p0));
            let _576_v130;
            _576_v130 = _dafny.Seq.of(false, (_this).f4, (_this).f4);
            let _577_v131;
            _577_v131 = _module.D4.create_DC13(_this.f16, false, p0, (_module.D4.create_DC13(new BigNumber(362), true, _this.f16, _576_v130)).dtor_cf17);
            let _578_v132;
            _578_v132 = _dafny.Map.Empty.slice().updateUnsafe(_577_v131,(_this).f4);
            _578_v132 = (_578_v132).update(_577_v131, (_this).f4);
            let _579_v133;
            let _nw87 = new _module.C0();
            _nw87.__ctor(new BigNumber(753));
            _579_v133 = _nw87;
            _579_v133 = _579_v133;
          }
        }
      }
      let _580_v134;
      _580_v134 = false;
      let _index76 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_573_v129).length));
      (_573_v129)[_index76] = true;
      let _581_v135;
      let _nw88 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _581_v135 = _nw88;
      let _582_v136;
      _582_v136 = _dafny.Set.fromElements(_581_v135);
      let _583_v137;
      let _init11 = function (_584_i16) {
        return _module.D0.create_DC1();
      };
      let _nw89 = Array((new BigNumber(20)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw89.length); _i0_11++) {
        _nw89[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _583_v137 = _nw89;
      let _585_v138;
      _585_v138 = _dafny.Map.Empty.slice().updateUnsafe(_582_v136,_583_v137);
      let _586_v139;
      _586_v139 = _dafny.Seq.of(_585_v138);
      let _index77 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_581_v135).length));
      (_581_v135)[_index77] = new BigNumber(((_586_v139)[_module.__default.safeIndex((_this).f15, new BigNumber((_586_v139).length))]).length);
      let _587_v140;
      _587_v140 = _module.D11.create_DC25(_580_v134, p0);
      let _index78 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_573_v129).length));
      let _index79 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_581_v135).length));
      let _rhs59 = (_587_v140).dtor_cf40;
      let _rhs60 = (_580_v134) === (_580_v134);
      let _rhs61 = _580_v134;
      let _rhs62 = p0;
      let _rhs63 = (_this).f4;
      let _lhs38 = _this;
      let _lhs39 = _573_v129;
      let _lhs40 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_573_v129).length));
      let _lhs41 = _581_v135;
      let _lhs42 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_581_v135).length));
      _lhs38.f16 = _rhs59;
      _580_v134 = _rhs60;
      _lhs39[_lhs40] = _rhs61;
      _lhs41[_lhs42] = _rhs62;
      _580_v134 = _rhs63;
      let _588_v141;
      _588_v141 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_555_v123, globalState),p0);
      r0 = _588_v141;
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = [];
      let _589_v0;
      let _nw90 = Array((new BigNumber(28)).toNumber()).fill(false);
      _589_v0 = _nw90;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_589_v0).length))) {
        let _590_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_590_i0)) && ((_590_i0).isLessThan(new BigNumber((_589_v0).length))))) {
          (_589_v0)[(_590_i0)] = !((_this).f4) || ((_this).f4);
        }
      }
      let _591_i1;
      _591_i1 = _dafny.ZERO;
      L2: {
        while (true) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_591_i1)) {
              break L2;
            }
            _591_i1 = (_591_i1).plus(_dafny.ONE);
            let _592_v1;
            _592_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,true);
            _592_v1 = (_592_v1).Merge(_592_v1);
            let _593_v2;
            _593_v2 = _dafny.MultiSet.fromElements(p0);
            let _594_v3;
            _594_v3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements((_this).f4)).IsDisjointFrom(_593_v2),(_this).f6);
            _594_v3 = (_module.__default.fm44(globalState)).Merge((_594_v3).Merge(_594_v3));
            let _595_v4;
            _595_v4 = _dafny.Seq.of(!((_this).f4));
            (_this).f16 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber((_595_v4).length))).length);
            let _596_v5;
            let _nw91 = new _module.C0();
            _nw91.__ctor((_this).f6);
            _596_v5 = _nw91;
          }
        }
      }
      let _597_v6;
      _597_v6 = _dafny.Seq.of(p0);
      let _598_v7;
      _598_v7 = _module.D0.create_DC2(new BigNumber((_597_v6).length));
      let _source13 = _598_v7;
      if (_source13.is_DC1) {
        let _599_v8;
        _599_v8 = false;
        let _600_v9;
        _600_v9 = new _dafny.CodePoint('n'.codePointAt(0));
        _599_v8 = _module.__default.fm0(_600_v9, globalState);
        let _601_v10;
        _601_v10 = _dafny.Set.fromElements((_this).f4, _599_v8, _599_v8, false);
        if (((_601_v10).Intersect(_601_v10)).equals(_601_v10)) {
          let _602_v11;
          _602_v11 = _dafny.Seq.of(_589_v0);
          let _603_v12;
          let _nw92 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _603_v12 = _nw92;
          let _604_v13;
          _604_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,!((_this).f4));
          let _605_v14;
          _605_v14 = _dafny.MultiSet.fromElements(_600_v9, _600_v9, _600_v9);
          let _606_v15;
          _606_v15 = _dafny.Seq.of((_this).f15);
          let _607_v16;
          let _nw93 = Array((new BigNumber(26)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = (_this).f6;
          _nw93[(_dafny.ONE).toNumber()] = _this.f16;
          _nw93[(new BigNumber(2)).toNumber()] = _module.__default.fm3((_this).f6, p0, globalState);
          _nw93[(new BigNumber(3)).toNumber()] = _this.f16;
          _nw93[(new BigNumber(4)).toNumber()] = _this.f16;
          _nw93[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_603_v12,(_this).fm7(globalState))).length);
          _nw93[(new BigNumber(6)).toNumber()] = (_this).f6;
          _nw93[(new BigNumber(7)).toNumber()] = (_this).f15;
          _nw93[(new BigNumber(8)).toNumber()] = (_this).f6;
          _nw93[(new BigNumber(9)).toNumber()] = new BigNumber((_604_v13).length);
          _nw93[(new BigNumber(10)).toNumber()] = (_this).f6;
          _nw93[(new BigNumber(11)).toNumber()] = (_this).f15;
          _nw93[(new BigNumber(12)).toNumber()] = new BigNumber((_597_v6).length);
          _nw93[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_this).f15);
          _nw93[(new BigNumber(14)).toNumber()] = _this.f16;
          _nw93[(new BigNumber(15)).toNumber()] = (_this).f15;
          _nw93[(new BigNumber(16)).toNumber()] = (_this).f6;
          _nw93[(new BigNumber(17)).toNumber()] = _this.f16;
          _nw93[(new BigNumber(18)).toNumber()] = (_this).f15;
          _nw93[(new BigNumber(19)).toNumber()] = (_this).f15;
          _nw93[(new BigNumber(20)).toNumber()] = new BigNumber((_605_v14).cardinality());
          _nw93[(new BigNumber(21)).toNumber()] = new BigNumber((_606_v15).length);
          _nw93[(new BigNumber(22)).toNumber()] = _this.f16;
          _nw93[(new BigNumber(23)).toNumber()] = _this.f16;
          _nw93[(new BigNumber(24)).toNumber()] = new BigNumber(-935);
          _nw93[(new BigNumber(25)).toNumber()] = new BigNumber(642);
          _607_v16 = _nw93;
          let _608_v17;
          _608_v17 = _dafny.MultiSet.fromElements(_607_v16, _607_v16);
          let _609_v18;
          _609_v18 = _dafny.Map.Empty.slice().updateUnsafe(!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_589_v0), _602_v11)),((((_608_v17).contains(_603_v12)) ? ((_608_v17).get(_603_v12)) : ((_606_v15)[_module.__default.safeIndex(_this.f16, new BigNumber((_606_v15).length))]))).plus(new BigNumber(907)));
          let _610_v19;
          _610_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(155),_this.f16);
          _609_v18 = (_609_v18).update((_this).f4, (((_610_v19).contains(_this.f16)) ? ((_610_v19).get(_this.f16)) : (_this.f16)));
          let _611_v20;
          _611_v20 = _module.D3.create_DC8(_module.__default.fm40(p0, _599_v8, new BigNumber((_module.__default.fm37(_599_v8, (_this).f4, (_this).f6, (_this).f4, globalState)).length), globalState));
          let _612_v21;
          _612_v21 = _module.D3.create_DC11(_611_v20);
          _612_v21 = _612_v21;
          let _index80 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_589_v0).length));
          (_589_v0)[_index80] = p0;
          let _613_v22;
          _613_v22 = _dafny.Seq.UnicodeFromString("shvveca");
          let _index81 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_589_v0).length));
          let _rhs64 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_613_v22, _613_v22)).length)));
          let _rhs65 = _module.__default.fm3((_this).f15, _599_v8, globalState);
          let _rhs66 = _599_v8;
          let _rhs67 = _599_v8;
          let _rhs68 = _599_v8;
          let _lhs43 = _this;
          let _lhs44 = _this;
          let _lhs45 = _589_v0;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(299), new BigNumber((_589_v0).length));
          _lhs43.f16 = _rhs64;
          _lhs44.f16 = _rhs65;
          _599_v8 = _rhs66;
          _lhs45[_lhs46] = _rhs67;
          _599_v8 = _rhs68;
          let _614_v23;
          let _nw94 = Array((new BigNumber(5)).toNumber()).fill(_dafny.MultiSet.Empty);
          _614_v23 = _nw94;
          let _615_v24;
          _615_v24 = _dafny.Map.Empty.slice().updateUnsafe(_600_v9,(_this).f4);
          let _616_v25;
          _616_v25 = _dafny.Seq.of(_614_v23, _614_v23);
          let _617_v26;
          let _nw95 = Array((new BigNumber(20)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = _614_v23;
          _nw95[(_dafny.ONE).toNumber()] = _614_v23;
          _nw95[(new BigNumber(2)).toNumber()] = (((((_615_v24).contains(_600_v9)) ? ((_615_v24).get(_600_v9)) : ((_589_v0)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_589_v0).length))]))) ? (_614_v23) : (_614_v23));
          _nw95[(new BigNumber(3)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(4)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(5)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(6)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(7)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(8)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(9)).toNumber()] = (_616_v25)[_module.__default.safeIndex(_this.f16, new BigNumber((_616_v25).length))];
          _nw95[(new BigNumber(10)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(11)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(12)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(13)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(14)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(15)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(16)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(17)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(18)).toNumber()] = _614_v23;
          _nw95[(new BigNumber(19)).toNumber()] = _614_v23;
          _617_v26 = _nw95;
          let _index82 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_617_v26).length));
          (_617_v26)[_index82] = _614_v23;
          let _index83 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_617_v26).length));
          (_617_v26)[_index83] = _614_v23;
          _613_v22 = _dafny.Seq.Concat(_dafny.Seq.Concat(_613_v22, _613_v22), _613_v22);
        } else {
          let _618_v27;
          _618_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f4);
          let _619_v28;
          _619_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_618_v27);
          _619_v28 = (_619_v28).update(p0, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(871),_599_v8));
          let _620_v29;
          _620_v29 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f15);
          (_this).f16 = (((_620_v29).contains((_this).f15)) ? ((_620_v29).get((_this).f15)) : (_this.f16));
          (_this).f16 = (_this).f6;
          let _621_v30;
          _621_v30 = _dafny.MultiSet.fromElements((_this).f15, new BigNumber(-639));
          let _622_v31;
          let _nw96 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _622_v31 = _nw96;
          let _623_v32;
          _623_v32 = _dafny.Seq.of(_622_v31, _622_v31, _622_v31);
          let _rhs69 = (_this.f16).plus((_this).f6);
          let _rhs70 = _621_v30;
          let _rhs71 = (_this).f4;
          let _rhs72 = _623_v32;
          let _lhs47 = _this;
          _lhs47.f16 = _rhs69;
          _621_v30 = _rhs70;
          _599_v8 = _rhs71;
          _623_v32 = _rhs72;
          (_this).f16 = _this.f16;
        }
        (_this).f16 = (_this.f16).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(591)), ((_624_v9) => function (_625_i2) {
          return _624_v9;
        })(_600_v9))).length));
        _589_v0 = _589_v0;
      } else if (_source13.is_DC2) {
        let _626___mcc_h0 = (_source13).cf1;
        let _627_cf1 = _626___mcc_h0;
        let _628_v33;
        _628_v33 = false;
        let _629_v34;
        _629_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_589_v0);
        let _630_v35;
        _630_v35 = _dafny.Seq.of(_589_v0, (((_629_v34).contains((_this).f4)) ? ((_629_v34).get((_this).f4)) : (_589_v0)));
        let _631_v36;
        _631_v36 = _dafny.Seq.UnicodeFromString("gx");
        let _632_v37;
        _632_v37 = _dafny.Map.Empty.slice().updateUnsafe((_630_v35)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("htbf")).length), new BigNumber((_630_v35).length))],_module.__default.fm3(new BigNumber((_631_v36).length), p0, globalState));
        let _633_v38;
        _633_v38 = _dafny.Map.Empty.slice().updateUnsafe(_627_cf1,(_this).f15);
        let _634_v39;
        _634_v39 = _dafny.Set.fromElements(_633_v38, _633_v38);
        let _635_v40;
        _635_v40 = _dafny.Map.Empty.slice().updateUnsafe(_628_v33,p0);
        let _636_v41;
        _636_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,_627_cf1);
        let _637_v42;
        _637_v42 = new _dafny.CodePoint('r'.codePointAt(0));
        let _638_v43;
        _638_v43 = _dafny.Set.fromElements(_637_v42, _637_v42);
        let _639_v44;
        let _nw97 = Array((new BigNumber(24)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = _this.f16;
        _nw97[(_dafny.ONE).toNumber()] = new BigNumber((_635_v40).length);
        _nw97[(new BigNumber(2)).toNumber()] = _this.f16;
        _nw97[(new BigNumber(3)).toNumber()] = new BigNumber(377);
        _nw97[(new BigNumber(4)).toNumber()] = _627_cf1;
        _nw97[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.of((_this).f6, _627_cf1, (_this).f6, new BigNumber((_629_v34).length), new BigNumber((_631_v36).length))).length);
        _nw97[(new BigNumber(6)).toNumber()] = (_this).f15;
        _nw97[(new BigNumber(7)).toNumber()] = new BigNumber(495);
        _nw97[(new BigNumber(8)).toNumber()] = _this.f16;
        _nw97[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_this.f16, _this.f16)).cardinality());
        _nw97[(new BigNumber(10)).toNumber()] = (((_636_v41).contains((_this).f4)) ? ((_636_v41).get((_this).f4)) : (_this.f16));
        _nw97[(new BigNumber(11)).toNumber()] = (_this).f6;
        _nw97[(new BigNumber(12)).toNumber()] = _627_cf1;
        _nw97[(new BigNumber(13)).toNumber()] = (_this).f6;
        _nw97[(new BigNumber(14)).toNumber()] = _this.f16;
        _nw97[(new BigNumber(15)).toNumber()] = new BigNumber((_638_v43).length);
        _nw97[(new BigNumber(16)).toNumber()] = _this.f16;
        _nw97[(new BigNumber(17)).toNumber()] = _this.f16;
        _nw97[(new BigNumber(18)).toNumber()] = _627_cf1;
        _nw97[(new BigNumber(19)).toNumber()] = _627_cf1;
        _nw97[(new BigNumber(20)).toNumber()] = new BigNumber((_631_v36).length);
        _nw97[(new BigNumber(21)).toNumber()] = new BigNumber((_636_v41).length);
        _nw97[(new BigNumber(22)).toNumber()] = new BigNumber((_631_v36).length);
        _nw97[(new BigNumber(23)).toNumber()] = new BigNumber(180);
        _639_v44 = _nw97;
        let _640_v45;
        _640_v45 = _dafny.Seq.of(_639_v44, _639_v44, _639_v44, _639_v44, _639_v44);
        let _641_v46;
        _641_v46 = _dafny.Map.Empty.slice().updateUnsafe(_631_v36,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(62)), ((_642_v38) => function (_643_i3) {
          return _642_v38;
        })(_633_v38))).length));
        let _644_v47;
        _644_v47 = _module.D1.create_DC3(p0);
        let _645_v48;
        _645_v48 = _dafny.Map.Empty.slice().updateUnsafe(_644_v47,_627_cf1);
        let _646_v49;
        let _nw98 = Array((new BigNumber(26)).toNumber());
        _nw98[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(_this.f16, new BigNumber((_631_v36).length));
        _nw98[(_dafny.ONE).toNumber()] = (_this).f6;
        _nw98[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_this).f15);
        _nw98[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_597_v6).length), _module.__default.fm3(new BigNumber((_634_v39).length), _628_v33, globalState));
        _nw98[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_627_cf1), _this.f16);
        _nw98[(new BigNumber(5)).toNumber()] = new BigNumber(813);
        _nw98[(new BigNumber(6)).toNumber()] = new BigNumber((_640_v45).length);
        _nw98[(new BigNumber(7)).toNumber()] = (_this).f6;
        _nw98[(new BigNumber(8)).toNumber()] = _this.f16;
        _nw98[(new BigNumber(9)).toNumber()] = new BigNumber(206);
        _nw98[(new BigNumber(10)).toNumber()] = ((_this).f15).multipliedBy((_this).f6);
        _nw98[(new BigNumber(11)).toNumber()] = new BigNumber(-867);
        _nw98[(new BigNumber(12)).toNumber()] = _this.f16;
        _nw98[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt((_this).f15, (_dafny.ZERO).minus((_dafny.ZERO).minus(_627_cf1)));
        _nw98[(new BigNumber(14)).toNumber()] = _module.__default.safeDivisionInt(_627_cf1, new BigNumber(915));
        _nw98[(new BigNumber(15)).toNumber()] = _627_cf1;
        _nw98[(new BigNumber(16)).toNumber()] = (_this).fm9(p0, p0, (((_641_v46).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), ((_649_v42) => function (_650_i4) {
          return _649_v42;
        })(_637_v42)))) ? ((_641_v46).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), ((_647_v42) => function (_648_i4) {
          return _647_v42;
        })(_637_v42)))) : (_this.f16)), globalState);
        _nw98[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_597_v6).length), _this.f16);
        _nw98[(new BigNumber(18)).toNumber()] = _627_cf1;
        _nw98[(new BigNumber(19)).toNumber()] = (_this).f15;
        _nw98[(new BigNumber(20)).toNumber()] = new BigNumber((_645_v48).length);
        _nw98[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("eadjey")).length);
        _nw98[(new BigNumber(22)).toNumber()] = _627_cf1;
        _nw98[(new BigNumber(23)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_597_v6).length), new BigNumber((_631_v36).length));
        _nw98[(new BigNumber(24)).toNumber()] = (_this).f15;
        _nw98[(new BigNumber(25)).toNumber()] = (((_636_v41).contains(_628_v33)) ? ((_636_v41).get(_628_v33)) : (new BigNumber(-222)));
        _646_v49 = _nw98;
        let _index84 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_639_v44).length));
        (_639_v44)[_index84] = _627_cf1;
        let _651_v50;
        _651_v50 = _dafny.Seq.of(new BigNumber(676), _627_cf1);
        let _index85 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_639_v44).length));
        let _rhs73 = !(((_this).f4) === (p0));
        let _rhs74 = (_this).f4;
        let _rhs75 = _632_v37;
        let _rhs76 = ((((_636_v41).contains(_628_v33)) ? ((_636_v41).get(_628_v33)) : (new BigNumber(94)))).plus(_module.__default.safeDivisionInt(_627_cf1, new BigNumber((_651_v50).length)));
        let _rhs77 = _597_v6;
        let _lhs48 = _639_v44;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_639_v44).length));
        _628_v33 = _rhs73;
        _628_v33 = _rhs74;
        _632_v37 = _rhs75;
        _lhs48[_lhs49] = _rhs76;
        _597_v6 = _rhs77;
        let _rhs78 = _628_v33;
        let _rhs79 = _631_v36;
        _628_v33 = _rhs78;
        _631_v36 = _rhs79;
        let _652_v51;
        _652_v51 = _dafny.Map.Empty.slice().updateUnsafe(_630_v35,(_this).f4);
        let _653_v52;
        let _out16;
        _out16 = (_this).m11((((_652_v51).contains(_630_v35)) ? ((_652_v51).get(_630_v35)) : (_628_v33)), _dafny.Seq.IsProperPrefixOf(_651_v50, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-844)), function (_654_i5) {
          return (_this).f6;
        })), (_641_v46).Merge(_641_v46), _589_v0, globalState);
        _653_v52 = _out16;
        if (p0) {
          _628_v33 = (_597_v6)[_module.__default.safeIndex(new BigNumber((_631_v36).length), new BigNumber((_597_v6).length))];
          let _index86 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_589_v0).length));
          (_589_v0)[_index86] = _628_v33;
          let _index87 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_589_v0).length));
          (_589_v0)[_index87] = _628_v33;
          _628_v33 = (_this).f4;
          _635_v40 = (_635_v40).update(p0, _module.__default.fm0(_637_v42, globalState));
          _653_v52 = (new BigNumber(628)).multipliedBy((_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_633_v38).length))).multipliedBy((_639_v44)[_module.__default.safeIndex(new BigNumber(987), new BigNumber((_639_v44).length))])));
        } else {
          let _index88 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_639_v44).length));
          (_639_v44)[_index88] = _module.__default.safeModuloInt(_module.__default.safeModuloInt(_627_cf1, (_this).f15), (_651_v50)[_module.__default.safeIndex(_this.f16, new BigNumber((_651_v50).length))]);
          let _index89 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_589_v0).length));
          (_589_v0)[_index89] = (_597_v6)[_module.__default.safeIndex((_this).f6, new BigNumber((_597_v6).length))];
          let _index90 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_589_v0).length));
          (_589_v0)[_index90] = ((_module.__default.fm0(_637_v42, globalState)) ? (!(_628_v33)) : (!(true)));
          let _655_v53;
          _655_v53 = _dafny.MultiSet.fromElements(_628_v33, _628_v33);
          let _656_v54;
          _656_v54 = _module.D11.create_DC25((_589_v0)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_589_v0).length))], _this.f16);
          let _index91 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_589_v0).length));
          let _index92 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_639_v44).length));
          let _rhs80 = _dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of((_this).f4, p0), _dafny.Seq.of(true, (_this).f4)), (_dafny.MultiSet.FromArray(_597_v6)).IsProperSubsetOf(_655_v53));
          let _rhs81 = ((_this).f15).multipliedBy((_656_v54).dtor_cf40);
          let _rhs82 = _637_v42;
          let _lhs50 = _589_v0;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_589_v0).length));
          let _lhs52 = _639_v44;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_639_v44).length));
          _lhs50[_lhs51] = _rhs80;
          _lhs52[_lhs53] = _rhs81;
          r0 = _rhs82;
          let _index93 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_589_v0).length));
          let _rhs83 = (_dafny.ZERO).minus(new BigNumber(((_633_v38).Merge(_dafny.Map.Empty.slice().updateUnsafe((_639_v44)[_module.__default.safeIndex(new BigNumber(987), new BigNumber((_639_v44).length))],_627_cf1))).length));
          let _rhs84 = (_639_v44)[_module.__default.safeIndex(new BigNumber(987), new BigNumber((_639_v44).length))];
          let _rhs85 = _646_v49;
          let _rhs86 = !(!(_628_v33));
          let _lhs54 = _589_v0;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_589_v0).length));
          _653_v52 = _rhs83;
          _653_v52 = _rhs84;
          _646_v49 = _rhs85;
          _lhs54[_lhs55] = _rhs86;
          _628_v33 = (_589_v0)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_589_v0).length))];
        }
      } else {
        let _657___mcc_h1 = (_source13).cf0;
        let _658_cf0 = _657___mcc_h1;
        let _659_v55;
        let _nw99 = new _module.C0();
        _nw99.__ctor((_this.f16).plus(_658_cf0));
        _659_v55 = _nw99;
        let _660_v56;
        _660_v56 = true;
        let _661_v57;
        let _nw100 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _661_v57 = _nw100;
        let _662_v58;
        _662_v58 = new _dafny.CodePoint('c'.codePointAt(0));
        let _663_v59;
        _663_v59 = _dafny.Map.Empty.slice().updateUnsafe(_659_v55,p0);
        let _664_v60;
        _664_v60 = _dafny.Seq.of(_662_v58);
        let _665_v61;
        _665_v61 = _dafny.MultiSet.fromElements(_589_v0, _589_v0);
        let _666_v62;
        _666_v62 = _dafny.Seq.of(new BigNumber((_665_v61).cardinality()));
        let _667_v63;
        _667_v63 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_661_v57,_662_v58)).length), (_this).f6, _module.__default.fm3((_this).f15, (_this).fm8((_dafny.ZERO).minus(_658_cf0), !((((_663_v59).contains(_659_v55)) ? ((_663_v59).get(_659_v55)) : (p0))), _664_v60, _666_v62, globalState), globalState));
        let _668_v64;
        let _nw101 = Array((new BigNumber(4)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_659_v55).f12);
        _nw101[(_dafny.ONE).toNumber()] = (_this).f15;
        _nw101[(new BigNumber(2)).toNumber()] = (_667_v63)[_module.__default.safeIndex(_658_cf0, new BigNumber((_667_v63).length))];
        _nw101[(new BigNumber(3)).toNumber()] = _658_cf0;
        _668_v64 = _nw101;
        let _index94 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_661_v57).length));
        (_661_v57)[_index94] = (_659_v55).f12;
        let _669_v65;
        _669_v65 = _dafny.MultiSet.fromElements(p0, _660_v56);
        let _index95 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_661_v57).length));
        let _rhs87 = ((_660_v56) ? ((_669_v65).IsSubsetOf(_669_v65)) : ((_this).f4));
        let _rhs88 = ((p0) ? (_this.f16) : (_this.f16));
        let _rhs89 = _658_cf0;
        let _lhs56 = _661_v57;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_661_v57).length));
        _660_v56 = _rhs87;
        _658_cf0 = _rhs88;
        _lhs56[_lhs57] = _rhs89;
        let _670_v66;
        _670_v66 = _dafny.Map.Empty.slice().updateUnsafe(_658_cf0,(_661_v57)[_module.__default.safeIndex(new BigNumber(78), new BigNumber((_661_v57).length))]);
        let _671_v67;
        _671_v67 = _dafny.MultiSet.fromElements(_658_cf0, new BigNumber(-806), new BigNumber((_597_v6).length), (((_670_v66).contains(_this.f16)) ? ((_670_v66).get(_this.f16)) : ((_659_v55).f12)));
        let _672_v68;
        _672_v68 = _671_v67;
        let _673_v69;
        _673_v69 = _dafny.Map.Empty.slice().updateUnsafe(_672_v68,((_dafny.Map.Empty.slice().updateUnsafe((_661_v57)[_module.__default.safeIndex(new BigNumber(78), new BigNumber((_661_v57).length))],(_this).f15))).Merge(_670_v66));
        _673_v69 = _673_v69;
        let _index96 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_661_v57).length));
        (_661_v57)[_index96] = ((_661_v57)[_module.__default.safeIndex(new BigNumber(78), new BigNumber((_661_v57).length))]).plus(new BigNumber((_666_v62).length));
      }
      let _index97 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_589_v0).length));
      (_589_v0)[_index97] = (_this).f4;
      let _674_v70;
      _674_v70 = _dafny.Seq.UnicodeFromString("uehfebb");
      let _675_v71;
      _675_v71 = new _dafny.CodePoint('v'.codePointAt(0));
      let _index98 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_589_v0).length));
      (_589_v0)[_index98] = ((_dafny.Seq.IsProperPrefixOf(_674_v70, _dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), function (_676_i6) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }))) ? (!_dafny.Seq.contains(_674_v70, _675_v71)) : ((_this).f4));
      let _677_v73;
      _677_v73 = _dafny.MultiSet.fromElements((_this).f6);
      let _hi2 = (_this.f16).minus(_this.f16);
      for (let _678_i7 = (_dafny.ZERO).minus(new BigNumber((function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of (_677_v73).Elements) {
          let _692_v72 = _compr_33;
          if ((_677_v73).contains(_692_v72)) {
            _coll33.push([(_692_v72).plus((_dafny.ZERO).minus(new BigNumber((_677_v73).cardinality()))),(((_677_v73).contains((_this).f6)) ? ((_677_v73).get((_this).f6)) : ((_this).f6))]);
          }
        }
        return _coll33;
      }()).length)); _678_i7.isLessThan(_hi2); _678_i7 = _678_i7.plus(_dafny.ONE)) {
        let _679_v75;
        let _init12 = function (_680_i8) {
          return (_680_i8).multipliedBy(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(904)), function () {
            let _coll32 = new _dafny.Set();
            for (const _compr_32 of _dafny.IntegerRange(new BigNumber(739), new BigNumber(930))) {
              let _681_v74 = _compr_32;
              if (((new BigNumber(739)).isLessThanOrEqualTo(_681_v74)) && ((_681_v74).isLessThan(new BigNumber(930)))) {
                _coll32.add(_module.__default.safeDivisionInt(_681_v74, _this.f16));
              }
            }
            return _coll32;
          }(), _dafny.Set.fromElements((_this).f6, (_this).f15))).length));
        };
        let _nw102 = Array((new BigNumber(6)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw102.length); _i0_12++) {
          _nw102[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _679_v75 = _nw102;
        let _index99 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_679_v75).length));
        (_679_v75)[_index99] = (_this).fm9(p0, (_this).f4, _678_i7, globalState);
        let _index100 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_679_v75).length));
        (_679_v75)[_index100] = (_this).f15;
        let _682_v76;
        let _nw103 = new _module.C0();
        _nw103.__ctor((_679_v75)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_679_v75).length))]);
        _682_v76 = _nw103;
        let _683_v77;
        _683_v77 = _dafny.MultiSet.fromElements(p0, (_589_v0)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_589_v0).length))]);
        let _684_v78;
        _684_v78 = _dafny.Seq.of(_683_v77, (_dafny.MultiSet.fromElements(true, (_this).f4, (_this).f4, (_589_v0)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_589_v0).length))], (_this).f4)).Intersect(_683_v77));
        _684_v78 = _dafny.Seq.Concat(_dafny.Seq.Concat(_684_v78, _module.__default.fm45(p0, (_this).f15, (_589_v0)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_589_v0).length))], (_682_v76).f12, globalState)), _dafny.Seq.of(_683_v77));
        let _685_v79;
        _685_v79 = _module.D1.create_DC3((_this).fm8(new BigNumber(779), (_589_v0)[_module.__default.safeIndex(new BigNumber(243), new BigNumber((_589_v0).length))], _674_v70, _dafny.Seq.Create(_module.__default.abs(new BigNumber(520)), ((_686_v75) => function (_687_i9) {
  return (_686_v75)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_686_v75).length))];
})(_679_v75)), globalState));
        let _688_v80;
        let _nw104 = Array((new BigNumber(4)).toNumber()).fill([]);
        _688_v80 = _nw104;
        let _pat_let_tv10 = p0;
        let _689_v81;
        _689_v81 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let11_0) {
          return function (_690_dt__update__tmp_h0) {
            return function (_pat_let12_0) {
              return function (_691_dt__update_hcf2_h0) {
                return _module.D1.create_DC3(_691_dt__update_hcf2_h0);
              }(_pat_let12_0);
            }(_pat_let_tv10);
          }(_pat_let11_0);
        }(_685_v79)).dtor_cf2,_688_v80);
        _689_v81 = (_689_v81).update(!(p0), _688_v80);
      }
      let _index101 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_589_v0).length));
      (_589_v0)[_index101] = (_this).f4;
      r0 = _675_v71;
      let _693_v82;
      let _init13 = function (_694_i10) {
        return _module.D1.create_DC3((_this).f4);
      };
      let _nw105 = Array((new BigNumber(3)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw105.length); _i0_13++) {
        _nw105[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _693_v82 = _nw105;
      r1 = _693_v82;
      return [r0, r1];
    }
    m11(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _695_v0;
      _695_v0 = _dafny.Seq.UnicodeFromString("h");
      _695_v0 = _695_v0;
      let _696_v1;
      _696_v1 = false;
      let _arr0 = _this.f5;
      let _index102 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length));
      _arr0[_index102] = p3;
      let _697_v2;
      _697_v2 = _dafny.Seq.of(_696_v1);
      let _698_v3;
      _698_v3 = _module.D1.create_DC4((_this).f15, _this.f16);
      let _pat_let_tv11 = p1;
      let _pat_let_tv12 = p0;
      let _arr1 = _this.f5;
      let _index103 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length));
      let _rhs90 = (_697_v2)[_module.__default.safeIndex(((_this).f15).minus((_this).f6), new BigNumber((_697_v2).length))];
      let _rhs91 = p3;
      let _rhs92 = p0;
      let _rhs93 = function (_source14) {
        if (_source14.is_DC4) {
          let _699___mcc_h0 = (_source14).cf3;
          let _700___mcc_h1 = (_source14).cf4;
          let _701_cf4 = _700___mcc_h1;
          let _702_cf3 = _699___mcc_h0;
          return _pat_let_tv11;
        } else {
          let _703___mcc_h2 = (_source14).cf2;
          let _704_cf2 = _703___mcc_h2;
          return _pat_let_tv12;
        }
      }(function (_pat_let13_0) {
        return function (_705_dt__update__tmp_h0) {
          return function (_pat_let14_0) {
            return function (_706_dt__update_hcf3_h0) {
              return _module.D1.create_DC4(_706_dt__update_hcf3_h0, (_705_dt__update__tmp_h0).dtor_cf4);
            }(_pat_let14_0);
          }(new BigNumber(-948));
        }(_pat_let13_0);
      }(_698_v3));
      let _rhs94 = false;
      let _lhs58 = _this.f5;
      let _lhs59 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length));
      _696_v1 = _rhs90;
      _lhs58[_lhs59] = _rhs91;
      _696_v1 = _rhs92;
      _696_v1 = _rhs93;
      _696_v1 = _rhs94;
      let _707_v4;
      _707_v4 = new _dafny.CodePoint('d'.codePointAt(0));
      let _708_v5;
      _708_v5 = _dafny.Map.Empty.slice().updateUnsafe(_707_v4,p0);
      let _709_i0;
      _709_i0 = _dafny.ZERO;
      L3: {
        while (((!_dafny.areEqual(_695_v0, _695_v0)) ? (((_this).f6).isLessThan(new BigNumber((_708_v5).length))) : (p1))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_709_i0)) {
              break L3;
            }
            _709_i0 = (_709_i0).plus(_dafny.ONE);
            let _710_v6;
            let _nw106 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
            _710_v6 = _nw106;
            _710_v6 = _710_v6;
            let _711_v7;
            _711_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_dafny.ZERO).minus((_this).f6));
            let _712_v8;
            let _nw107 = new _module.C0();
            _nw107.__ctor((((_this).f4) ? (new BigNumber((_711_v7).length)) : ((_this).f6)));
            _712_v8 = _nw107;
            _696_v1 = _dafny.areEqual(_module.__default.fm40((_this).f4, p1, (_712_v8).f12, globalState), _dafny.Seq.of(_696_v1));
            let _713_v9;
            _713_v9 = _dafny.Set.fromElements((_696_v1) || ((_this).f4), p0);
            _713_v9 = _713_v9;
          }
        }
      }
      if (false) {
        let _714_v10;
        _714_v10 = _dafny.Set.fromElements(p3);
        let _715_v11;
        let _nw108 = Array((new BigNumber(23)).toNumber()).fill([]);
        _715_v11 = _nw108;
        let _716_v12;
        let _nw109 = Array((new BigNumber(12)).toNumber());
        _nw109[(_dafny.ZERO).toNumber()] = _715_v11;
        _nw109[(_dafny.ONE).toNumber()] = _715_v11;
        _nw109[(new BigNumber(2)).toNumber()] = _715_v11;
        _nw109[(new BigNumber(3)).toNumber()] = _715_v11;
        _nw109[(new BigNumber(4)).toNumber()] = _715_v11;
        _nw109[(new BigNumber(5)).toNumber()] = _715_v11;
        _nw109[(new BigNumber(6)).toNumber()] = _715_v11;
        _nw109[(new BigNumber(7)).toNumber()] = _715_v11;
        _nw109[(new BigNumber(8)).toNumber()] = _715_v11;
        _nw109[(new BigNumber(9)).toNumber()] = _715_v11;
        _nw109[(new BigNumber(10)).toNumber()] = _715_v11;
        _nw109[(new BigNumber(11)).toNumber()] = _715_v11;
        _716_v12 = _nw109;
        let _index104 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_716_v12).length));
        (_716_v12)[_index104] = _715_v11;
        let _717_v13;
        _717_v13 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(_this.f16)).length), (_this).fm9((_this).f4, p1, (_this).f15, globalState)),(_this).f15);
        let _718_v14;
        _718_v14 = _dafny.Seq.of(_715_v11, _715_v11, _715_v11, _715_v11);
        let _719_v15;
        _719_v15 = _dafny.MultiSet.fromElements(p0);
        let _index105 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_716_v12).length));
        let _rhs95 = _714_v10;
        let _rhs96 = (_718_v14)[_module.__default.safeIndex((((_719_v15).contains(!(false))) ? ((_719_v15).get(!(false))) : ((_this).f15)), new BigNumber((_718_v14).length))];
        let _rhs97 = (_717_v13).Merge(_717_v13);
        let _lhs60 = _716_v12;
        let _lhs61 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_716_v12).length));
        _714_v10 = _rhs95;
        _lhs60[_lhs61] = _rhs96;
        _717_v13 = _rhs97;
        (_this).f16 = (_this).fm7(globalState);
        if ((_this).f4) {
          let _720_v16;
          let _nw110 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _720_v16 = _nw110;
          let _721_v17;
          _721_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_695_v0).length),_720_v16);
          _721_v17 = _721_v17;
          let _722_v18;
          _722_v18 = _dafny.MultiSet.fromElements((_this).f15);
          _722_v18 = ((_722_v18).Intersect(_dafny.MultiSet.fromElements((_this).f15))).update(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_695_v0).length)), (_this).f15), _module.__default.abs((_this).f15));
          let _arr2 = (_this.f5)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length))];
          let _index106 = _module.__default.safeIndex(new BigNumber(253), new BigNumber(((_this.f5)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length))]).length));
          _arr2[_index106] = (_this).f4;
          let _723_v19;
          _723_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,_695_v0);
          let _724_v20;
          _724_v20 = _dafny.Seq.of((_dafny.ZERO).minus(_this.f16));
          let _725_v21;
          _725_v21 = _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_724_v20).length)));
          let _726_v22;
          _726_v22 = _module.D12.create_DC28((_this).f6, _this.f16, _696_v1, !(_696_v1), (((_725_v21).contains((_this).f4)) ? ((_725_v21).get((_this).f4)) : (new BigNumber(146))));
          let _arr3 = (_this.f5)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length))];
          let _index107 = _module.__default.safeIndex(new BigNumber(253), new BigNumber(((_this.f5)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length))]).length));
          let _rhs98 = new BigNumber((_723_v19).length);
          let _rhs99 = (_726_v22).dtor_cf44;
          let _rhs100 = p0;
          let _lhs62 = _this;
          let _lhs63 = (_this.f5)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length))];
          let _lhs64 = _module.__default.safeIndex(new BigNumber(253), new BigNumber(((_this.f5)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length))]).length));
          _lhs62.f16 = _rhs98;
          r0 = _rhs99;
          _lhs63[_lhs64] = _rhs100;
          let _727_v23;
          _727_v23 = _dafny.Seq.of(_695_v0, _695_v0);
          _696_v1 = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_727_v23, _727_v23), _727_v23));
          let _arr4 = (_this.f5)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length))];
          let _index108 = _module.__default.safeIndex(new BigNumber(253), new BigNumber(((_this.f5)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_this.f5).length))]).length));
          _arr4[_index108] = !((_this).f4);
        } else {
          r0 = ((_this).f15).multipliedBy(_this.f16);
          let _728_v24;
          _728_v24 = _module.D0.create_DC2((_dafny.ZERO).minus((_this).f6));
          _728_v24 = _728_v24;
          (_this).f16 = _this.f16;
          let _729_v25;
          _729_v25 = _dafny.Map.Empty.slice().updateUnsafe(_697_v2,p0);
          let _730_v26;
          _730_v26 = _dafny.Set.fromElements((_this).f15);
          _729_v25 = (_729_v25).update(_697_v2, (_730_v26).IsSubsetOf(_730_v26));
          r0 = new BigNumber((_695_v0).length);
        }
        let _731_v27;
        _731_v27 = _module.D0.create_DC0((_this).f6);
        let _732_v28;
        let _nw111 = new _module.C0();
        _nw111.__ctor(new BigNumber((_697_v2).length));
        _732_v28 = _nw111;
        let _733_v29;
        _733_v29 = _dafny.Map.Empty.slice().updateUnsafe(_732_v28,_696_v1);
        let _734_v30;
        _734_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_733_v29);
        let _735_v31;
        _735_v31 = _dafny.Seq.of(new BigNumber(-159), (_dafny.ZERO).minus(_this.f16));
        let _rhs101 = _731_v27;
        let _rhs102 = ((_dafny.Map.Empty.slice().updateUnsafe(_732_v28,(_this).f4)).update(_732_v28, !((_this).f4))).Merge((((_734_v30).contains((_this).f15)) ? ((_734_v30).get((_this).f15)) : (_733_v29)));
        let _rhs103 = (_this).f6;
        let _rhs104 = (_this).fm8(((_this).f15).multipliedBy((_732_v28).f12), ((_735_v31)[_module.__default.safeIndex((_this).f6, new BigNumber((_735_v31).length))]).isLessThanOrEqualTo((_732_v28).f12), _695_v0, _735_v31, globalState);
        let _rhs105 = (p1) && (p0);
        _731_v27 = _rhs101;
        _733_v29 = _rhs102;
        r0 = _rhs103;
        _696_v1 = _rhs104;
        _696_v1 = _rhs105;
        let _736_v32;
        _736_v32 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_732_v28).f12),(((_719_v15).contains(!(_696_v1))) ? ((_719_v15).get(!(_696_v1))) : ((_this).fm7(globalState))));
        let _737_v33;
        _737_v33 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_732_v28);
        let _738_v34;
        _738_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(_737_v33).equals(_737_v33),_736_v32);
        _736_v32 = (((_738_v34).contains(p0)) ? ((_738_v34).get(p0)) : (_736_v32));
      } else {
        (_this).f16 = _this.f16;
        let _739_v35;
        let _nw112 = new _module.C0();
        _nw112.__ctor(new BigNumber(520));
        _739_v35 = _nw112;
        let _740_v36;
        _740_v36 = _dafny.MultiSet.fromElements(p1);
        let _741_v37;
        _741_v37 = _module.D3.create_DC10(new BigNumber((_740_v36).cardinality()));
        let _742_v38;
        _742_v38 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,((_741_v37).dtor_cf11).isEqualTo((_739_v35).f12));
        let _743_v39;
        _743_v39 = _dafny.Seq.of(_this.f16);
        let _744_v40;
        _744_v40 = _dafny.Seq.of(_dafny.Seq.of((_739_v35).f12), _743_v39, _743_v39, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-982)), ((_745_v35) => function (_746_i1) {
          return (_dafny.ZERO).minus((_745_v35).f12);
        })(_739_v35)), _743_v39);
        _742_v38 = (_742_v38).update((_this).f6, !_dafny.areEqual(_744_v40, _744_v40));
        let _747_v41;
        let _init14 = function (_748_i3) {
          return (_748_i3).minus((_this).f6);
        };
        let _nw113 = Array((new BigNumber(17)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw113.length); _i0_14++) {
          _nw113[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _747_v41 = _nw113;
        let _749_v42;
        _749_v42 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wirva"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-993)), function (_750_i2) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })),_dafny.Map.Empty.slice().updateUnsafe(p1,_747_v41));
        let _751_v43;
        _751_v43 = _dafny.Map.Empty.slice().updateUnsafe(p0,_747_v41);
        _749_v42 = (_749_v42).update(_695_v0, (_751_v43).Merge(_751_v43));
        _696_v1 = (_this).f4;
      }
      let _752_v44;
      _752_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f16);
      r0 = (_this.f16).plus(_module.__default.safeDivisionInt(new BigNumber((_752_v44).length), new BigNumber(-690)));
      let _753_v45;
      _753_v45 = _dafny.Map.Empty.slice().updateUnsafe(_707_v4,(_this).f15);
      if ((_753_v45).equals((_753_v45).Merge((_753_v45).update(_707_v4, new BigNumber(220))))) {
        (_this).f16 = _module.__default.safeModuloInt(_this.f16, (_this).f15);
        let _754_v46;
        _754_v46 = _dafny.Seq.of((new BigNumber((_695_v0).length)).multipliedBy(new BigNumber(668)), (_this.f16).multipliedBy((_this).f15));
        let _755_v47;
        _755_v47 = _module.D12.create_DC27(_dafny.Seq.of(_697_v2));
        _754_v46 = _module.__default.fm39(_695_v0, _755_v47, p1, globalState);
        let _756_v48;
        _756_v48 = _dafny.MultiSet.fromElements(!(false), p0);
        let _757_v49;
        _757_v49 = _dafny.MultiSet.fromElements(_this.f16, _this.f16, new BigNumber((_756_v48).cardinality()));
        let _758_v50;
        _758_v50 = _module.D9.create_DC19(_dafny.Seq.UnicodeFromString("imf"));
        let _759_v51;
        _759_v51 = _dafny.Map.Empty.slice().updateUnsafe(_757_v49,(_758_v50).dtor_cf25);
        let _760_v52;
        _760_v52 = _dafny.Set.fromElements(_this.f16);
        let _761_v53;
        let _762_v54;
        let _out17;
        let _out18;
        let _outcollector7 = _module.__default.m0(_759_v51, (_760_v52).IsSubsetOf(_dafny.Set.fromElements((_this).f6)), _757_v49, _696_v1, globalState);
        _out17 = _outcollector7[0];
        _out18 = _outcollector7[1];
        _761_v53 = _out17;
        _762_v54 = _out18;
        let _763_v55;
        _763_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_754_v46)[_module.__default.safeIndex(new BigNumber((_695_v0).length), new BigNumber((_754_v46).length))]);
        let _764_v56;
        let _nw114 = Array((new BigNumber(6)).toNumber());
        _nw114[(_dafny.ZERO).toNumber()] = (new BigNumber((_695_v0).length)).plus((_dafny.ZERO).minus((_this).f6));
        _nw114[(_dafny.ONE).toNumber()] = new BigNumber(398);
        _nw114[(new BigNumber(2)).toNumber()] = new BigNumber(62);
        _nw114[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(_this.f16, (_dafny.ZERO).minus(new BigNumber((_756_v48).cardinality())));
        _nw114[(new BigNumber(4)).toNumber()] = (((_763_v55).contains((_this).f6)) ? ((_763_v55).get((_this).f6)) : (_this.f16));
        _nw114[(new BigNumber(5)).toNumber()] = (_762_v54).multipliedBy(_this.f16);
        _764_v56 = _nw114;
        let _rhs106 = _764_v56;
        let _rhs107 = p1;
        let _rhs108 = (_this.f16).plus((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), function (_765_i4) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("csd")).length);
        })).length)).plus(new BigNumber((_module.__default.fm46(_707_v4, globalState)).length)));
        let _rhs109 = (((_756_v48).contains(_696_v1)) ? ((_756_v48).get(_696_v1)) : ((_this).f6));
        let _lhs65 = _this;
        let _lhs66 = _this;
        _764_v56 = _rhs106;
        _761_v53 = _rhs107;
        _lhs65.f16 = _rhs108;
        _lhs66.f16 = _rhs109;
        if ((_this).f4) {
          let _766_v57;
          let _767_v58;
          let _out19;
          let _out20;
          let _outcollector8 = _module.__default.m0((_759_v51).update(_757_v49, _695_v0), p0, _757_v49, p1, globalState);
          _out19 = _outcollector8[0];
          _out20 = _outcollector8[1];
          _766_v57 = _out19;
          _767_v58 = _out20;
          let _768_v60;
          _768_v60 = _dafny.Set.fromElements(_757_v49);
          let _769_v61;
          let _770_v62;
          let _out21;
          let _out22;
          let _outcollector9 = _module.__default.m0(function () {
            let _coll34 = new _dafny.Map();
            for (const _compr_34 of (_768_v60).Elements) {
              let _771_v59 = _compr_34;
              if ((_768_v60).contains(_771_v59)) {
                _coll34.push([_771_v59,_dafny.Seq.UnicodeFromString("lgs")]);
              }
            }
            return _coll34;
          }(), (new BigNumber(580)).isLessThanOrEqualTo(new BigNumber((_752_v44).length)), (_757_v49).Union(_757_v49), _696_v1, globalState);
          _out21 = _outcollector9[0];
          _out22 = _outcollector9[1];
          _769_v61 = _out21;
          _770_v62 = _out22;
          _696_v1 = p1;
          _769_v61 = p1;
          r0 = _762_v54;
        } else {
          let _772_v63;
          _772_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,p1);
          _752_v44 = (_752_v44).update((_this).f4, ((_this).f15).multipliedBy((_this).fm9(false, p1, new BigNumber((_772_v63).length), globalState)));
          let _773_v64;
          _773_v64 = _dafny.Map.Empty.slice().updateUnsafe(_761_v53,p0);
          let _774_v65;
          _774_v65 = _dafny.Seq.of(_757_v49);
          let _775_v66;
          _775_v66 = _dafny.Map.Empty.slice().updateUnsafe((_773_v64).Merge(_773_v64),new BigNumber(((_774_v65)[_module.__default.safeIndex((_dafny.ZERO).minus(_this.f16), new BigNumber((_774_v65).length))]).cardinality()));
          _775_v66 = _775_v66;
          let _776_v67;
          let _nw115 = new _module.C0();
          _nw115.__ctor((_this.f16).plus(new BigNumber(843)));
          _776_v67 = _nw115;
          _756_v48 = ((_756_v48).Union(_756_v48)).Union((_756_v48).Union(_module.__default.fm47((_this).f6, globalState)));
          _696_v1 = _761_v53;
        }
      } else {
        _696_v1 = (p0) || (p1);
        _707_v4 = _707_v4;
        let _777_v68;
        let _init15 = function (_778_i5) {
          return (((_this).f4) ? (_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_this).f6, new BigNumber(-594), _this.f16)).cardinality()), (_dafny.ZERO).minus((_this).f15))) : (_dafny.Seq.of(_this.f16)));
        };
        let _nw116 = Array((new BigNumber(15)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw116.length); _i0_15++) {
          _nw116[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _777_v68 = _nw116;
        let _779_v69;
        _779_v69 = _dafny.Map.Empty.slice().updateUnsafe(p1,_696_v1);
        let _780_v70;
        _780_v70 = _dafny.Seq.of(new BigNumber((_779_v69).length));
        let _781_v71;
        _781_v71 = _dafny.Seq.of((_this).f15, new BigNumber((_780_v70).length), (_this).f6);
        let _index109 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_777_v68).length));
        (_777_v68)[_index109] = _781_v71;
        let _782_v72;
        _782_v72 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_707_v4, globalState),_707_v4);
        let _index110 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_777_v68).length));
        (_777_v68)[_index110] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of((_this).f15), _781_v71), _module.__default.safeIndex(new BigNumber(876), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f15), _781_v71)).length)), new BigNumber((_782_v72).length)), _dafny.Seq.of(_this.f16, (_this).f15));
        let _783_v73;
        _783_v73 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_695_v0, _695_v0),(_this).f6);
        _783_v73 = (_783_v73).update(_695_v0, (_this.f16).plus(new BigNumber(424)));
        _696_v1 = p0;
      }
      r0 = _this.f16;
      return r0;
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f4 = false;
      this._f17 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f17, f4) {
      let _this = this;
      (_this)._f17 = f17;
      (_this)._f4 = f4;
      return;
    }
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(new BigNumber(486)));
    };
    fm7(globalState) {
      let _this = this;
      return new BigNumber(831);
    };
    fm48(p0, p1, p2, globalState) {
      let _this = this;
      return (_module.D1.create_DC3((_this).f4)).dtor_cf2;
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _hi3 = p0;
      for (let _784_i0 = (_this).fm7(globalState); _784_i0.isLessThan(_hi3); _784_i0 = _784_i0.plus(_dafny.ONE)) {
        let _785_v0;
        let _nw117 = Array((new BigNumber(18)).toNumber()).fill(false);
        _785_v0 = _nw117;
        let _init16 = function (_786_i1) {
          return (_this).f4;
        };
        let _nw118 = Array((new BigNumber(2)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw118.length); _i0_16++) {
          _nw118[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _785_v0 = _nw118;
        let _787_v1;
        _787_v1 = new BigNumber(805);
        _787_v1 = new BigNumber(-78);
        let _788_v2;
        _788_v2 = _dafny.Set.fromElements((_this).f4);
        let _789_v3;
        let _790_v4;
        let _out23;
        let _out24;
        let _outcollector10 = (_this).m12(_787_v1, (_788_v2).Difference(_788_v2), _dafny.Map.Empty.slice().updateUnsafe(_787_v1,new BigNumber((_788_v2).length)), (_this).f4, globalState);
        _out23 = _outcollector10[0];
        _out24 = _outcollector10[1];
        _789_v3 = _out23;
        _790_v4 = _out24;
        if ((_this).f4) {
          let _791_v5;
          _791_v5 = true;
          _791_v5 = (_788_v2).IsProperSubsetOf(_dafny.Set.fromElements((_this).f4, (_this).f4, (_this).f4));
          let _792_v6;
          _792_v6 = _dafny.Set.fromElements(_784_i0, _789_v3);
          let _793_v7;
          _793_v7 = _dafny.MultiSet.fromElements(false, _791_v5, _791_v5, _791_v5);
          let _794_v8;
          _794_v8 = new _dafny.CodePoint('p'.codePointAt(0));
          let _795_v9;
          _795_v9 = _dafny.Set.fromElements(_792_v6, (_792_v6).Intersect(_dafny.Set.fromElements(_784_i0, new BigNumber((_793_v7).cardinality()), new BigNumber((_dafny.Seq.update((_this).f17, _module.__default.safeIndex(_790_v4, new BigNumber(((_this).f17).length)), _794_v8)).length), new BigNumber(((_793_v7).update(_791_v5, _module.__default.abs((_dafny.ZERO).minus(p0)))).cardinality()))));
          let _796_v10;
          _796_v10 = _module.D11.create_DC26(_791_v5);
          let _797_v11;
          _797_v11 = _dafny.Map.Empty.slice().updateUnsafe(_796_v10,_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hjgsgbnw"), (_this).f17));
          let _798_v12;
          let _nw119 = Array((new BigNumber(16)).toNumber()).fill(_module.D2.Default());
          _798_v12 = _nw119;
          let _index111 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_798_v12).length));
          (_798_v12)[_index111] = _module.D2.create_DC6();
          let _799_v14;
          _799_v14 = _dafny.MultiSet.fromElements(_796_v10);
          let _800_v15;
          _800_v15 = _module.D2.create_DC6();
          let _index112 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_798_v12).length));
          let _rhs110 = ((false) ? (_795_v9) : ((_795_v9).Difference(_795_v9)));
          let _rhs111 = ((_797_v11).Merge(_797_v11)).Merge((function () {
            let _coll35 = new _dafny.Map();
            for (const _compr_35 of (_799_v14).Elements) {
              let _801_v13 = _compr_35;
              if ((_799_v14).contains(_801_v13)) {
                _coll35.push([_801_v13,(_this).f17]);
              }
            }
            return _coll35;
          }()).Merge(_797_v11));
          let _rhs112 = _800_v15;
          let _rhs113 = (!((_this).f4)) && ((_787_v1).isLessThanOrEqualTo(_787_v1));
          let _rhs114 = (_this).f4;
          let _lhs67 = _798_v12;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_798_v12).length));
          _795_v9 = _rhs110;
          _797_v11 = _rhs111;
          _lhs67[_lhs68] = _rhs112;
          _791_v5 = _rhs113;
          _791_v5 = _rhs114;
          let _802_v16;
          let _nw120 = new _module.C0();
          _nw120.__ctor(new BigNumber(-980));
          _802_v16 = _nw120;
          _791_v5 = !(((_791_v5) ? ((_this).f4) : (_791_v5)));
          let _803_v17;
          _803_v17 = _module.D2.create_DC5(_785_v0);
          let _804_v18;
          _804_v18 = _dafny.Map.Empty.slice().updateUnsafe((_791_v5) || ((_this).f4),(_803_v17).dtor_cf5);
          _804_v18 = (_804_v18).update(!(_791_v5) || (_791_v5), _785_v0);
        } else {
          let _805_v19;
          _805_v19 = _dafny.MultiSet.fromElements(_784_i0);
          _805_v19 = (_805_v19).Intersect(_805_v19);
          let _806_v20;
          let _init17 = ((_807_v4) => function (_808_i2) {
            return (_808_i2).minus(_807_v4);
          })(_790_v4);
          let _nw121 = Array((new BigNumber(21)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw121.length); _i0_17++) {
            _nw121[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _806_v20 = _nw121;
          let _index113 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_806_v20).length));
          (_806_v20)[_index113] = _784_i0;
          let _809_v21;
          _809_v21 = _dafny.Seq.of(_789_v3);
          let _810_v22;
          _810_v22 = _dafny.Seq.of(_809_v21);
          let _811_v23;
          _811_v23 = false;
          let _812_v24;
          _812_v24 = _dafny.Seq.of((_this).f4, _811_v23);
          let _813_v25;
          _813_v25 = _dafny.Seq.of((_this).f17, (_this).f17);
          let _index114 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_806_v20).length));
          let _rhs115 = (_this).fm7(globalState);
          let _rhs116 = _dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm49((_this).f17, _812_v24, _789_v3, globalState), _module.__default.safeIndex(new BigNumber(((_813_v25)[_module.__default.safeIndex(_789_v3, new BigNumber((_813_v25).length))]).length), new BigNumber((_module.__default.fm49((_this).f17, _812_v24, _789_v3, globalState)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(345)), ((_814_i0) => function (_815_i3) {
            return _814_i0;
          })(_784_i0))), _dafny.Seq.of(_809_v21));
          let _rhs117 = true;
          let _rhs118 = (_dafny.ZERO).minus((_789_v3).plus((_dafny.ZERO).minus(_787_v1)));
          let _rhs119 = (_this).f4;
          let _lhs69 = _806_v20;
          let _lhs70 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_806_v20).length));
          _lhs69[_lhs70] = _rhs115;
          _810_v22 = _rhs116;
          _811_v23 = _rhs117;
          _790_v4 = _rhs118;
          _811_v23 = _rhs119;
          let _816_v26;
          _816_v26 = _dafny.MultiSet.fromElements(_811_v23, _811_v23, false);
          let _index115 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_806_v20).length));
          (_806_v20)[_index115] = (_dafny.ZERO).minus((((_816_v26).contains(_811_v23)) ? ((_816_v26).get(_811_v23)) : (_790_v4)));
          _788_v2 = _dafny.Set.fromElements(!((_this).f4));
          let _817_v27;
          let _nw122 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _817_v27 = _nw122;
          let _index116 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_817_v27).length));
          (_817_v27)[_index116] = new _dafny.CodePoint('l'.codePointAt(0));
          let _818_v28;
          _818_v28 = new _dafny.CodePoint('c'.codePointAt(0));
          let _index117 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_817_v27).length));
          (_817_v27)[_index117] = _818_v28;
        }
      }
      let _819_v29;
      let _nw123 = Array((new BigNumber(25)).toNumber()).fill(false);
      _819_v29 = _nw123;
      let _820_v30;
      _820_v30 = _module.D10.create_DC22(_dafny.Seq.UnicodeFromString("jdyfjcy"), ((_dafny.ZERO).minus(p0)).isEqualTo(new BigNumber(122)), !((_this).f4), _819_v29, (_dafny.ZERO).minus(p0));
      let _source15 = _820_v30;
      if (_source15.is_DC22) {
        let _821___mcc_h0 = (_source15).cf30;
        let _822___mcc_h1 = (_source15).cf31;
        let _823___mcc_h2 = (_source15).cf32;
        let _824___mcc_h3 = (_source15).cf33;
        let _825___mcc_h4 = (_source15).cf34;
        let _826_cf34 = _825___mcc_h4;
        let _827_cf33 = _824___mcc_h3;
        let _828_cf32 = _823___mcc_h2;
        let _829_cf31 = _822___mcc_h1;
        let _830_cf30 = _821___mcc_h0;
        _826_cf34 = ((_829_cf31) ? (_module.__default.safeDivisionInt(_826_cf34, p0)) : (new BigNumber(569)));
        let _831_v31;
        _831_v31 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f4) ? (_828_cf32) : ((_this).f4)),_826_cf34);
        let _832_v32;
        _832_v32 = _dafny.Seq.of(p0);
        let _833_v33;
        _833_v33 = _dafny.Set.fromElements(_828_cf32, true);
        let _834_v34;
        let _nw124 = Array((new BigNumber(9)).toNumber());
        _nw124[(_dafny.ZERO).toNumber()] = p0;
        _nw124[(_dafny.ONE).toNumber()] = _826_cf34;
        _nw124[(new BigNumber(2)).toNumber()] = new BigNumber((_832_v32).length);
        _nw124[(new BigNumber(3)).toNumber()] = new BigNumber(156);
        _nw124[(new BigNumber(4)).toNumber()] = p0;
        _nw124[(new BigNumber(5)).toNumber()] = p0;
        _nw124[(new BigNumber(6)).toNumber()] = new BigNumber((_830_cf30).length);
        _nw124[(new BigNumber(7)).toNumber()] = (_this).fm7(globalState);
        _nw124[(new BigNumber(8)).toNumber()] = new BigNumber((_833_v33).length);
        _834_v34 = _nw124;
        let _835_v35;
        _835_v35 = _dafny.Map.Empty.slice().updateUnsafe(_834_v34,p0);
        _831_v31 = (_831_v31).update((_dafny.Map.Empty.slice().updateUnsafe(_834_v34,(_this).fm7(globalState))).equals((_835_v35).update(_834_v34, new BigNumber(-360))), p0);
        let _836_v36;
        _836_v36 = _module.D2.create_DC7((_826_cf34).plus(p0));
        let _rhs120 = p0;
        let _rhs121 = _828_cf32;
        let _rhs122 = _836_v36;
        _826_cf34 = _rhs120;
        _828_cf32 = _rhs121;
        _836_v36 = _rhs122;
        let _837_v37;
        let _nw125 = new _module.C0();
        _nw125.__ctor(p0);
        _837_v37 = _nw125;
      } else {
        let _838___mcc_h5 = (_source15).cf29;
        let _839_cf29 = _838___mcc_h5;
        let _840_v38;
        _840_v38 = _module.D2.create_DC6();
        let _source16 = _840_v38;
        if (_source16.is_DC6) {
          let _841_v39;
          _841_v39 = new _dafny.CodePoint('w'.codePointAt(0));
          let _842_v40;
          _842_v40 = _module.D9.create_DC20((_this).f4, _841_v39, new BigNumber(((_this).f17).length));
          let _843_v41;
          _843_v41 = _dafny.Seq.of(_842_v40);
          let _844_v42;
          _844_v42 = _843_v41;
          _843_v41 = (_843_v41);
          let _index118 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_819_v29).length));
          (_819_v29)[_index118] = (((_this).f4) ? ((_this).f4) : ((_this).f4));
          let _845_v43;
          let _nw126 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _845_v43 = _nw126;
          let _index119 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_845_v43).length));
          (_845_v43)[_index119] = _module.__default.safeDivisionInt(p0, p0);
          let _846_v44;
          let _nw127 = Array((new BigNumber(27)).toNumber()).fill([]);
          _846_v44 = _nw127;
          let _847_v45;
          _847_v45 = true;
          let _index120 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_819_v29).length));
          let _index121 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_845_v43).length));
          let _rhs123 = (_this).f4;
          let _rhs124 = p0;
          let _rhs125 = _846_v44;
          let _rhs126 = !(_847_v45);
          let _lhs71 = _819_v29;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_819_v29).length));
          let _lhs73 = _845_v43;
          let _lhs74 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_845_v43).length));
          _lhs71[_lhs72] = _rhs123;
          _lhs73[_lhs74] = _rhs124;
          _846_v44 = _rhs125;
          _847_v45 = _rhs126;
          let _848_v46;
          let _nw128 = Array((new BigNumber(14)).toNumber()).fill(false);
          _848_v46 = _nw128;
          _848_v46 = _848_v46;
          let _849_v47;
          _849_v47 = _dafny.Seq.of((_819_v29)[_module.__default.safeIndex(new BigNumber(751), new BigNumber((_819_v29).length))]);
          let _850_v48;
          _850_v48 = _dafny.Seq.of(new BigNumber((_849_v47).length), p0);
          let _851_v49;
          _851_v49 = _dafny.MultiSet.fromElements(false, (_819_v29)[_module.__default.safeIndex(new BigNumber(751), new BigNumber((_819_v29).length))]);
          let _852_v50;
          _852_v50 = _dafny.Map.Empty.slice().updateUnsafe((_850_v48)[_module.__default.safeIndex(new BigNumber(((_this).f17).length), new BigNumber((_850_v48).length))],(((_819_v29)[_module.__default.safeIndex(new BigNumber(751), new BigNumber((_819_v29).length))]) ? ((_dafny.ZERO).minus(p0)) : (new BigNumber((_851_v49).cardinality()))));
          let _853_v51;
          _853_v51 = _dafny.Set.fromElements((_845_v43)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_845_v43).length))]);
          let _854_v52;
          _854_v52 = _dafny.MultiSet.fromElements(new BigNumber((_853_v51).length), p0);
          _852_v50 = (_852_v50).update((_845_v43)[_module.__default.safeIndex(new BigNumber(437), new BigNumber((_845_v43).length))], ((_847_v45) ? (new BigNumber(460)) : (new BigNumber((_854_v52).cardinality()))));
        } else if (_source16.is_DC7) {
          let _855___mcc_h6 = (_source16).cf6;
          let _856_cf6 = _855___mcc_h6;
          let _857_v53;
          _857_v53 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f4, (_this).f4, !(true)));
          _857_v53 = _857_v53;
          let _858_v54;
          _858_v54 = _dafny.Seq.of((_this).f4);
          let _859_v55;
          _859_v55 = _dafny.Seq.of((((_this).f4) ? ((_858_v54)[_module.__default.safeIndex((_dafny.ZERO).minus(_856_cf6), new BigNumber((_858_v54).length))]) : (false)));
          _859_v55 = _858_v54;
          let _860_v56;
          _860_v56 = _dafny.MultiSet.fromElements(p0, p0);
          let _861_v57;
          _861_v57 = _dafny.Map.Empty.slice().updateUnsafe((((_this).fm48((_this).f4, _856_cf6, p0, globalState)) ? ((_this).f17) : (_module.__default.fm50(new BigNumber(892), p0, new BigNumber((_860_v56).cardinality()), (_this).f4, globalState))),p0);
          let _862_v58;
          _862_v58 = _dafny.Seq.of(p0, p0);
          let _863_v59;
          _863_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_862_v58).length),(_this).f17);
          let _864_v60;
          _864_v60 = _dafny.Seq.of(new BigNumber((_863_v59).length));
          let _865_v61;
          _865_v61 = _dafny.Seq.of(new BigNumber(325), _856_cf6, p0, (_864_v60)[_module.__default.safeIndex(_856_cf6, new BigNumber((_864_v60).length))]);
          _861_v57 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(228)), function (_866_i4) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }),((_865_v61)[_module.__default.safeIndex(p0, new BigNumber((_865_v61).length))]).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(761)), function (_867_i5) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          })).length)));
          let _868_v62;
          _868_v62 = false;
          let _869_v63;
          let _init18 = ((_870_cf6) => function (_871_i6) {
            return _module.__default.safeDivisionInt(_871_i6, _870_cf6);
          })(_856_cf6);
          let _nw129 = Array((new BigNumber(22)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw129.length); _i0_18++) {
            _nw129[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _869_v63 = _nw129;
          let _872_v64;
          _872_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _index122 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_869_v63).length));
          (_869_v63)[_index122] = (new BigNumber((_dafny.Seq.of(_856_cf6)).length)).multipliedBy((((_872_v64).contains(p0)) ? ((_872_v64).get(p0)) : (_856_cf6)));
          let _index123 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_869_v63).length));
          let _rhs127 = _868_v62;
          let _rhs128 = (((_872_v64).contains(_856_cf6)) ? ((_872_v64).get(_856_cf6)) : (p0));
          let _rhs129 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("hasb"), (_this).f17);
          let _lhs75 = _869_v63;
          let _lhs76 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_869_v63).length));
          _868_v62 = _rhs127;
          _lhs75[_lhs76] = _rhs128;
          _868_v62 = _rhs129;
        } else {
          let _873___mcc_h7 = (_source16).cf5;
          let _874_cf5 = _873___mcc_h7;
          let _875_v65;
          let _nw130 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
          _875_v65 = _nw130;
          let _876_v66;
          _876_v66 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4);
          let _877_v67;
          _877_v67 = _dafny.Set.fromElements(_876_v66);
          let _index124 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_875_v65).length));
          (_875_v65)[_index124] = (((_this).f4) ? (_877_v67) : (_877_v67));
          let _878_v68;
          _878_v68 = _dafny.MultiSet.fromElements((_this).f4);
          let _index125 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_875_v65).length));
          (_875_v65)[_index125] = function () {
            let _coll36 = new _dafny.Set();
            for (const _compr_36 of ((_dafny.Map.Empty.slice().updateUnsafe(_876_v66,_dafny.MultiSet.fromElements(false))).update(_876_v66, _878_v68)).Keys.Elements) {
              let _879_v69 = _compr_36;
              if (((_dafny.Map.Empty.slice().updateUnsafe(_876_v66,_dafny.MultiSet.fromElements(false))).update(_876_v66, _878_v68)).contains(_879_v69)) {
                _coll36.add(_879_v69);
              }
            }
            return _coll36;
          }();
          let _880_v70;
          _880_v70 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _881_v71;
          _881_v71 = _dafny.Set.fromElements((_this).f4, (_this).f4);
          let _882_v72;
          let _nw131 = Array((new BigNumber(16)).toNumber());
          _nw131[(_dafny.ZERO).toNumber()] = new BigNumber(((_880_v70).update(p0, p0)).length);
          _nw131[(_dafny.ONE).toNumber()] = p0;
          _nw131[(new BigNumber(2)).toNumber()] = new BigNumber(245);
          _nw131[(new BigNumber(3)).toNumber()] = p0;
          _nw131[(new BigNumber(4)).toNumber()] = new BigNumber((_881_v71).length);
          _nw131[(new BigNumber(5)).toNumber()] = p0;
          _nw131[(new BigNumber(6)).toNumber()] = p0;
          _nw131[(new BigNumber(7)).toNumber()] = p0;
          _nw131[(new BigNumber(8)).toNumber()] = p0;
          _nw131[(new BigNumber(9)).toNumber()] = p0;
          _nw131[(new BigNumber(10)).toNumber()] = new BigNumber((_839_cf29).length);
          _nw131[(new BigNumber(11)).toNumber()] = p0;
          _nw131[(new BigNumber(12)).toNumber()] = p0;
          _nw131[(new BigNumber(13)).toNumber()] = p0;
          _nw131[(new BigNumber(14)).toNumber()] = new BigNumber(((_this).f17).length);
          _nw131[(new BigNumber(15)).toNumber()] = new BigNumber(((_880_v70).Merge(_880_v70)).length);
          _882_v72 = _nw131;
          let _index126 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_882_v72).length));
          (_882_v72)[_index126] = p0;
          let _index127 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_882_v72).length));
          (_882_v72)[_index127] = new BigNumber(844);
          let _883_v73;
          let _nw132 = new _module.C0();
          _nw132.__ctor(((_882_v72)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_882_v72).length))]).multipliedBy((_882_v72)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_882_v72).length))]));
          _883_v73 = _nw132;
          let _884_v74;
          _884_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(74),_881_v71);
          let _885_v75;
          _885_v75 = _dafny.Seq.of((_883_v73).f12, new BigNumber((_884_v74).length));
          let _886_v76;
          _886_v76 = _dafny.Seq.of(((_dafny.ZERO).minus(new BigNumber((_885_v75).length))).minus(p0), (_883_v73).f12);
          _886_v76 = _886_v76;
        }
        let _887_v77;
        _887_v77 = _dafny.Seq.of((_this).f4);
        let _888_v78;
        _888_v78 = _dafny.Seq.of(_887_v77, _887_v77, _dafny.Seq.of((_this).f4, (_this).f4));
        _887_v77 = (_module.D4.create_DC13(p0, true, (_dafny.ZERO).minus(new BigNumber(((_this).f17).length)), (_888_v78)[_module.__default.safeIndex(p0, new BigNumber((_888_v78).length))])).dtor_cf17;
        let _889_v79;
        let _nw133 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _889_v79 = _nw133;
        let _890_v80;
        _890_v80 = _dafny.Map.Empty.slice().updateUnsafe(_889_v79,(_this).f4);
        _890_v80 = (_890_v80).update(_889_v79, (p0).isLessThanOrEqualTo(p0));
        if (((p0).isLessThanOrEqualTo(p0)) || ((p0).isLessThanOrEqualTo(p0))) {
          let _891_v81;
          let _nw134 = Array((new BigNumber(4)).toNumber());
          _nw134[(_dafny.ZERO).toNumber()] = _819_v29;
          _nw134[(_dafny.ONE).toNumber()] = _819_v29;
          _nw134[(new BigNumber(2)).toNumber()] = _819_v29;
          _nw134[(new BigNumber(3)).toNumber()] = _819_v29;
          _891_v81 = _nw134;
          let _892_v82;
          let _nw135 = new _module.C1();
          _nw135.__ctor(p0, p0, _891_v81, (new BigNumber((_dafny.Seq.UnicodeFromString("bw")).length)).multipliedBy(new BigNumber(206)), (_this).f4);
          _892_v82 = _nw135;
          let _893_v83;
          _893_v83 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("mxhimf")).length));
          let _894_v84;
          _894_v84 = _dafny.Map.Empty.slice().updateUnsafe(_893_v83,!((_this).f4) || ((_this).f4));
          let _895_v86;
          _895_v86 = _dafny.Seq.of(_893_v83, _893_v83);
          _894_v84 = ((function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of (_895_v86).Elements) {
              let _896_v85 = _compr_37;
              if (_dafny.Seq.contains(_895_v86, _896_v85)) {
                _coll37.push([_896_v85,(_this).f4]);
              }
            }
            return _coll37;
          }()).Merge(_894_v84)).Merge((_894_v84).Merge(_894_v84));
          let _897_v87;
          let _nw136 = new _module.C0();
          _nw136.__ctor(p0);
          _897_v87 = _nw136;
          let _898_v88;
          _898_v88 = false;
          let _899_v89;
          let _nw137 = Array((new BigNumber(4)).toNumber());
          _nw137[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_this).f17, (_this).f17);
          _nw137[(_dafny.ONE).toNumber()] = (_this).f17;
          _nw137[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat((_this).f17, (_this).f17);
          _nw137[(new BigNumber(3)).toNumber()] = (_this).f17;
          _899_v89 = _nw137;
          let _index128 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_899_v89).length));
          (_899_v89)[_index128] = (_this).f17;
          let _900_v90;
          _900_v90 = new _dafny.CodePoint('e'.codePointAt(0));
          let _901_v91;
          _901_v91 = _dafny.Map.Empty.slice().updateUnsafe(_900_v90,_819_v29);
          let _902_v92;
          _902_v92 = _dafny.MultiSet.fromElements(_819_v29, (((_901_v91).contains(new _dafny.CodePoint('o'.codePointAt(0)))) ? ((_901_v91).get(new _dafny.CodePoint('o'.codePointAt(0)))) : (_819_v29)), _819_v29, _819_v29, ((_898_v88) ? (_819_v29) : (_819_v29)));
          let _903_v93;
          _903_v93 = _dafny.Map.Empty.slice().updateUnsafe((_892_v82).f15,new BigNumber(((_this).f17).length));
          let _index129 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_899_v89).length));
          let _rhs130 = (((_902_v92).contains(_819_v29)) ? ((_902_v92).get(_819_v29)) : (((((_903_v93).contains((_dafny.ZERO).minus(_892_v82.f16))) ? ((_903_v93).get((_dafny.ZERO).minus(_892_v82.f16))) : ((_893_v83)[_module.__default.safeIndex(p0, new BigNumber((_893_v83).length))]))).minus((_892_v82).f15)));
          let _rhs131 = !((_897_v87).f12).isEqualTo((_897_v87).f12);
          let _rhs132 = _898_v88;
          let _rhs133 = _898_v88;
          let _rhs134 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-804)), ((_904_v90) => function (_905_i7) {
            return _904_v90;
          })(_900_v90));
          let _lhs77 = _892_v82;
          let _lhs78 = _899_v89;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_899_v89).length));
          _lhs77.f16 = _rhs130;
          _898_v88 = _rhs131;
          _898_v88 = _rhs132;
          _898_v88 = _rhs133;
          _lhs78[_lhs79] = _rhs134;
          _900_v90 = new _dafny.CodePoint('x'.codePointAt(0));
        } else {
          let _906_v94;
          _906_v94 = new BigNumber(297);
          _906_v94 = p0;
          let _907_v95;
          _907_v95 = false;
          let _908_v96;
          _908_v96 = _module.D1.create_DC4(p0, p0);
          let _909_v97;
          let _nw138 = Array((new BigNumber(10)).toNumber()).fill([]);
          _909_v97 = _nw138;
          let _910_v98;
          let _nw139 = new _module.C1();
          _nw139.__ctor(p0, p0, _909_v97, _906_v94, (_this).f4);
          _910_v98 = _nw139;
          let _911_v99;
          _911_v99 = _dafny.Seq.of(_910_v98);
          let _912_v100;
          _912_v100 = _dafny.Set.fromElements((_this).f17, (_this).f17);
          let _913_v101;
          _913_v101 = _dafny.Map.Empty.slice().updateUnsafe(_912_v100,_911_v99);
          let _914_v102;
          _914_v102 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_907_v95);
          let _915_v103;
          _915_v103 = _dafny.Set.fromElements((_this).f4, (((_914_v102).contains(_907_v95)) ? ((_914_v102).get(_907_v95)) : (true)));
          let _916_v104;
          _916_v104 = new _dafny.CodePoint('e'.codePointAt(0));
          let _917_v105;
          _917_v105 = _dafny.Map.Empty.slice().updateUnsafe(_916_v104,_915_v103);
          let _rhs135 = p0;
          let _rhs136 = !_dafny.areEqual(_911_v99, (((_913_v101).contains(_912_v100)) ? ((_913_v101).get(_912_v100)) : (_911_v99)));
          let _rhs137 = _908_v96;
          let _rhs138 = ((_915_v103).Union((((_917_v105).contains(_916_v104)) ? ((_917_v105).get(_916_v104)) : (_dafny.Set.fromElements(!(true)))))).IsSubsetOf(_dafny.Set.fromElements(_907_v95, (_this).f4));
          let _rhs139 = p0;
          _906_v94 = _rhs135;
          _907_v95 = _rhs136;
          _908_v96 = _rhs137;
          _907_v95 = _rhs138;
          _906_v94 = _rhs139;
          _907_v95 = false;
          let _918_v106;
          _918_v106 = _dafny.Seq.of(_889_v79);
          let _919_v107;
          _919_v107 = _module.D2.create_DC5(_819_v29);
          _907_v95 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_918_v106).length), _906_v94))).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_919_v107)).length))));
          let _920_v108;
          _920_v108 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-14),(_910_v98).f15);
          let _921_v109;
          let _922_v110;
          let _out25;
          let _out26;
          let _outcollector11 = (_this).m12(p0, _915_v103, (_920_v108).Merge(_920_v108), (_this).f4, globalState);
          _out25 = _outcollector11[0];
          _out26 = _outcollector11[1];
          _921_v109 = _out25;
          _922_v110 = _out26;
        }
      }
      let _923_v111;
      let _nw140 = new _module.C0();
      _nw140.__ctor(p0);
      _923_v111 = _nw140;
      let _924_i8;
      _924_i8 = _dafny.ZERO;
      L4: {
        while (((_this).f4) && ((_this).f4)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_924_i8)) {
              break L4;
            }
            _924_i8 = (_924_i8).plus(_dafny.ONE);
            let _index130 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length));
            (_819_v29)[_index130] = (_this).f4;
            let _index131 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length));
            (_819_v29)[_index131] = (_this).f4;
            let _925_v112;
            _925_v112 = _dafny.Map.Empty.slice().updateUnsafe(p0,!((_this).f4));
            let _926_v113;
            _926_v113 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_819_v29)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length))]);
            let _927_v114;
            _927_v114 = _dafny.Set.fromElements((_this).f4);
            _925_v112 = _module.__default.fm51(_dafny.Map.Empty.slice().updateUnsafe((((_926_v113).contains(_module.__default.fm0(new _dafny.CodePoint('c'.codePointAt(0)), globalState))) ? ((_926_v113).get(_module.__default.fm0(new _dafny.CodePoint('c'.codePointAt(0)), globalState))) : (false)),_927_v114), true, globalState);
            let _928_v115;
            _928_v115 = _dafny.Seq.of(_dafny.Seq.IsPrefixOf((_this).f17, (_this).f17), (((_926_v113).contains((_this).f4)) ? ((_926_v113).get((_this).f4)) : ((_819_v29)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length))])), (_819_v29)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length))], (_this).f4);
            let _929_v116;
            _929_v116 = _module.D11.create_DC26((_819_v29)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length))]);
            let _930_v117;
            _930_v117 = _dafny.MultiSet.fromElements((_this).f4);
            let _931_v118;
            _931_v118 = _module.D16.create_DC33((_929_v116).dtor_cf41, (_923_v111).f12, _930_v117, p0, _927_v114);
            let _932_v119;
            _932_v119 = _module.D4.create_DC13(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_931_v118), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(_931_v118)).length)), _931_v118)).length), (_819_v29)[_module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length))], p0, _928_v115);
            let _pat_let_tv13 = p0;
            let _pat_let_tv14 = _928_v115;
            _928_v115 = _dafny.Seq.Concat((function (_pat_let15_0) {
              return function (_933_dt__update__tmp_h1) {
                return function (_pat_let16_0) {
                  return function (_934_dt__update_hcf14_h0) {
                    return function (_pat_let17_0) {
                      return function (_935_dt__update_hcf17_h0) {
                        return _module.D4.create_DC13(_934_dt__update_hcf14_h0, (_933_dt__update__tmp_h1).dtor_cf15, (_933_dt__update__tmp_h1).dtor_cf16, _935_dt__update_hcf17_h0);
                      }(_pat_let17_0);
                    }(_pat_let_tv14);
                  }(_pat_let16_0);
                }(_pat_let_tv13);
              }(_pat_let15_0);
            }(_932_v119)).dtor_cf17, _928_v115);
            let _index132 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length));
            let _index133 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length));
            let _rhs140 = ((((_this).f4) ? (_927_v114) : (_927_v114))).IsSubsetOf(_927_v114);
            let _rhs141 = !((_this).f4);
            let _lhs80 = _819_v29;
            let _lhs81 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length));
            let _lhs82 = _819_v29;
            let _lhs83 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_819_v29).length));
            _lhs80[_lhs81] = _rhs140;
            _lhs82[_lhs83] = _rhs141;
          }
        }
      }
      let _936_v120;
      _936_v120 = new BigNumber(906);
      _936_v120 = _module.__default.safeModuloInt(_936_v120, (_923_v111).f12);
      _936_v120 = (new BigNumber(((_this).f17).length)).plus(_936_v120);
      let _937_v121;
      _937_v121 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f4) ? ((_923_v111).fm17((_this).f4, globalState)) : ((_this).f4)),((_923_v111).f12).minus(new BigNumber(518)));
      r0 = _937_v121;
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = [];
      let _938_v0;
      _938_v0 = _dafny.Seq.of((_this).f4, (_this).f4);
      let _939_v1;
      _939_v1 = _module.D3.create_DC8(_938_v0);
      let _940_v2;
      let _nw141 = Array((new BigNumber(26)).toNumber());
      _nw141[(_dafny.ZERO).toNumber()] = _939_v1;
      _nw141[(_dafny.ONE).toNumber()] = _module.D3.create_DC8(_938_v0);
      _nw141[(new BigNumber(2)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(3)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(4)).toNumber()] = _module.D3.create_DC8(_938_v0);
      _nw141[(new BigNumber(5)).toNumber()] = _module.D3.create_DC8(_938_v0);
      _nw141[(new BigNumber(6)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(7)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(8)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(9)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(10)).toNumber()] = _module.D3.create_DC8(_938_v0);
      _nw141[(new BigNumber(11)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(12)).toNumber()] = _module.__default.fm52(new BigNumber(-402), globalState);
      _nw141[(new BigNumber(13)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(14)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(15)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(16)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(17)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(18)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(19)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(20)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(21)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(22)).toNumber()] = _module.D3.create_DC8(_938_v0);
      _nw141[(new BigNumber(23)).toNumber()] = _module.D3.create_DC8(_938_v0);
      _nw141[(new BigNumber(24)).toNumber()] = _939_v1;
      _nw141[(new BigNumber(25)).toNumber()] = _939_v1;
      _940_v2 = _nw141;
      let _index134 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_940_v2).length));
      (_940_v2)[_index134] = _939_v1;
      let _pat_let_tv15 = _938_v0;
      let _index135 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_940_v2).length));
      (_940_v2)[_index135] = function (_pat_let18_0) {
        return function (_941_dt__update__tmp_h0) {
          return function (_pat_let19_0) {
            return function (_942_dt__update_hcf7_h0) {
              return _module.D3.create_DC8(_942_dt__update_hcf7_h0);
            }(_pat_let19_0);
          }(_pat_let_tv15);
        }(_pat_let18_0);
      }(_939_v1);
      let _943_v3;
      _943_v3 = new BigNumber(187);
      let _944_v4;
      let _nw142 = Array((new BigNumber(9)).toNumber());
      _nw142[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(_943_v3, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_943_v3, _943_v3), _module.__default.safeIndex(new BigNumber(408), new BigNumber((_dafny.Seq.of(_943_v3, _943_v3)).length)), _943_v3)).length));
      _nw142[(_dafny.ONE).toNumber()] = _943_v3;
      _nw142[(new BigNumber(2)).toNumber()] = _943_v3;
      _nw142[(new BigNumber(3)).toNumber()] = _943_v3;
      _nw142[(new BigNumber(4)).toNumber()] = _943_v3;
      _nw142[(new BigNumber(5)).toNumber()] = (_943_v3).plus(new BigNumber(577));
      _nw142[(new BigNumber(6)).toNumber()] = _module.__default.fm3(_943_v3, (_this).f4, globalState);
      _nw142[(new BigNumber(7)).toNumber()] = _943_v3;
      _nw142[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_this).f17).length));
      _944_v4 = _nw142;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_944_v4).length))) {
        let _945_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_945_i0)) && ((_945_i0).isLessThan(new BigNumber((_944_v4).length))))) {
          (_944_v4)[(_945_i0)] = _module.__default.safeModuloInt(_945_i0, (_dafny.ZERO).minus((_943_v3).minus(_943_v3)));
        }
      }
      let _946_v5;
      let _nw143 = Array((_dafny.ONE).toNumber());
      _nw143[(_dafny.ZERO).toNumber()] = (_this).f4;
      _946_v5 = _nw143;
      let _947_v6;
      _947_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,_946_v5);
      let _948_v7;
      _948_v7 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((new BigNumber((_947_v6).length)).multipliedBy(_943_v3)),(_dafny.ZERO).minus(_943_v3));
      _948_v7 = (_948_v7).update(_943_v3, _943_v3);
      let _949_v8;
      _949_v8 = _dafny.MultiSet.fromElements(!((_this).f4), p0);
      let _950_v11;
      _950_v11 = new _dafny.CodePoint('n'.codePointAt(0));
      let _951_v12;
      let _nw144 = Array((new BigNumber(13)).toNumber());
      _nw144[(_dafny.ZERO).toNumber()] = _946_v5;
      _nw144[(_dafny.ONE).toNumber()] = _946_v5;
      _nw144[(new BigNumber(2)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(3)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(4)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(5)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(6)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(7)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(8)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(9)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(10)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(11)).toNumber()] = _946_v5;
      _nw144[(new BigNumber(12)).toNumber()] = _946_v5;
      _951_v12 = _nw144;
      let _952_v14;
      let _nw145 = new _module.C1();
      _nw145.__ctor(new BigNumber((_module.__default.fm50((((_949_v8).contains((_this).f4)) ? ((_949_v8).get((_this).f4)) : (_943_v3)), new BigNumber((function () {
        let _coll38 = new _dafny.Map();
        for (const _compr_38 of _dafny.IntegerRange(new BigNumber(321), new BigNumber(377))) {
          let _953_v9 = _compr_38;
          if (((new BigNumber(321)).isLessThanOrEqualTo(_953_v9)) && ((_953_v9).isLessThan(new BigNumber(377)))) {
            _coll38.push([_module.__default.safeDivisionInt(_953_v9, new BigNumber((function () {
              let _coll39 = new _dafny.Map();
              for (const _compr_39 of _dafny.IntegerRange(new BigNumber(72), new BigNumber(693))) {
                let _954_v10 = _compr_39;
                if (((new BigNumber(72)).isLessThanOrEqualTo(_954_v10)) && ((_954_v10).isLessThan(new BigNumber(693)))) {
                  _coll39.push([(_954_v10).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cii")).length)),(_this).f4]);
                }
              }
              return _coll39;
            }()).length)),new BigNumber(((_this).f17).length)]);
          }
        }
        return _coll38;
      }()).length), _943_v3, _module.__default.fm0(_950_v11, globalState), globalState)).length), _943_v3, _951_v12, _module.__default.fm3(new BigNumber((function () {
        let _coll40 = new _dafny.Map();
        for (const _compr_40 of _dafny.IntegerRange(new BigNumber(412), new BigNumber(-676))) {
          let _955_v13 = _compr_40;
          if (((new BigNumber(412)).isLessThanOrEqualTo(_955_v13)) && ((_955_v13).isLessThan(new BigNumber(-676)))) {
            _coll40.push([(_955_v13).multipliedBy(_943_v3),p0]);
          }
        }
        return _coll40;
      }()).length), p0, globalState), p0);
      _952_v14 = _nw145;
      let _956_v15;
      _956_v15 = _dafny.Set.fromElements(p0, p0, p0);
      _956_v15 = _module.__default.fm53(_module.__default.safeModuloInt(_943_v3, _943_v3), _dafny.Seq.Concat(_module.__default.fm54(_952_v14.f16, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), ((_957_p0) => function (_958_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(false,_957_p0);
      })(p0))), globalState);
      let _959_i2;
      _959_i2 = _dafny.ZERO;
      L5: {
        while (!((((p0) ? (_943_v3) : ((_952_v14).f15))).isLessThan((_952_v14.f16).multipliedBy((_952_v14).f15)))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_959_i2)) {
              break L5;
            }
            _959_i2 = (_959_i2).plus(_dafny.ONE);
            let _index136 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_946_v5).length));
            (_946_v5)[_index136] = ((_952_v14).f15).isLessThan(new BigNumber(21));
            let _index137 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_946_v5).length));
            (_946_v5)[_index137] = _module.__default.fm0(_950_v11, globalState);
            (_952_v14).f16 = _module.__default.fm3((_this).fm7(globalState), p0, globalState);
            let _index138 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_946_v5).length));
            (_946_v5)[_index138] = (_this).f4;
            let _index139 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_946_v5).length));
            (_946_v5)[_index139] = (new BigNumber(((_this).f17).length)).isLessThanOrEqualTo(_943_v3);
          }
        }
      }
      r0 = _950_v11;
      let _960_v16;
      _960_v16 = _module.D1.create_DC3(p0);
      let _nw146 = Array((new BigNumber(12)).toNumber());
      _nw146[(_dafny.ZERO).toNumber()] = _960_v16;
      _nw146[(_dafny.ONE).toNumber()] = _960_v16;
      _nw146[(new BigNumber(2)).toNumber()] = _960_v16;
      _nw146[(new BigNumber(3)).toNumber()] = _960_v16;
      _nw146[(new BigNumber(4)).toNumber()] = _960_v16;
      _nw146[(new BigNumber(5)).toNumber()] = _960_v16;
      _nw146[(new BigNumber(6)).toNumber()] = _960_v16;
      _nw146[(new BigNumber(7)).toNumber()] = _960_v16;
      _nw146[(new BigNumber(8)).toNumber()] = _960_v16;
      _nw146[(new BigNumber(9)).toNumber()] = function (_pat_let20_0) {
        return function (_961_dt__update__tmp_h1) {
          return function (_pat_let21_0) {
            return function (_962_dt__update_hcf2_h0) {
              return _module.D1.create_DC3(_962_dt__update_hcf2_h0);
            }(_pat_let21_0);
          }((_this).f4);
        }(_pat_let20_0);
      }(_960_v16);
      _nw146[(new BigNumber(10)).toNumber()] = _module.__default.fm55(!(p0), (_this).f4, new BigNumber(547), globalState);
      _nw146[(new BigNumber(11)).toNumber()] = _module.D1.create_DC3((_this).f4);
      r1 = _nw146;
      return [r0, r1];
    }
    m12(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _963_v0;
      let _nw147 = Array((new BigNumber(16)).toNumber()).fill(false);
      _963_v0 = _nw147;
      let _hi4 = new BigNumber((_dafny.Seq.of(_963_v0, _963_v0, _963_v0, _963_v0)).length);
      for (let _964_i0 = new BigNumber(((p1).Intersect(_dafny.Set.fromElements((_this).f4, (_this).f4))).length); _964_i0.isLessThan(_hi4); _964_i0 = _964_i0.plus(_dafny.ONE)) {
        let _965_v1;
        let _nw148 = new _module.C0();
        _nw148.__ctor(new BigNumber(271));
        _965_v1 = _nw148;
        let _966_v2;
        _966_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0);
        let _967_v3;
        _967_v3 = _dafny.Seq.of(_964_i0);
        let _968_v4;
        _968_v4 = _module.D9.create_DC19(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-645)), function (_969_i1) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}));
        let _970_v5;
        let _nw149 = Array((new BigNumber(21)).toNumber());
        _nw149[(_dafny.ZERO).toNumber()] = p0;
        _nw149[(_dafny.ONE).toNumber()] = _964_i0;
        _nw149[(new BigNumber(2)).toNumber()] = _964_i0;
        _nw149[(new BigNumber(3)).toNumber()] = new BigNumber(896);
        _nw149[(new BigNumber(4)).toNumber()] = (_965_v1).f12;
        _nw149[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_967_v3)).cardinality());
        _nw149[(new BigNumber(6)).toNumber()] = (_965_v1).f12;
        _nw149[(new BigNumber(7)).toNumber()] = _964_i0;
        _nw149[(new BigNumber(8)).toNumber()] = p0;
        _nw149[(new BigNumber(9)).toNumber()] = _964_i0;
        _nw149[(new BigNumber(10)).toNumber()] = _964_i0;
        _nw149[(new BigNumber(11)).toNumber()] = p0;
        _nw149[(new BigNumber(12)).toNumber()] = new BigNumber(((_this).f17).length);
        _nw149[(new BigNumber(13)).toNumber()] = new BigNumber(((_this).f17).length);
        _nw149[(new BigNumber(14)).toNumber()] = (_965_v1).f12;
        _nw149[(new BigNumber(15)).toNumber()] = new BigNumber(((_968_v4).dtor_cf25).length);
        _nw149[(new BigNumber(16)).toNumber()] = new BigNumber(-648);
        _nw149[(new BigNumber(17)).toNumber()] = _964_i0;
        _nw149[(new BigNumber(18)).toNumber()] = p0;
        _nw149[(new BigNumber(19)).toNumber()] = _964_i0;
        _nw149[(new BigNumber(20)).toNumber()] = p0;
        _970_v5 = _nw149;
        let _971_v6;
        _971_v6 = _dafny.Seq.of(_970_v5);
        let _972_v7;
        _972_v7 = _dafny.Map.Empty.slice().updateUnsafe(_966_v2,_971_v6);
        let _973_v8;
        _973_v8 = _dafny.Seq.of(_971_v6);
        _972_v7 = (_972_v7).update(_966_v2, (_973_v8)[_module.__default.safeIndex(p0, new BigNumber((_973_v8).length))]);
        let _974_v9;
        let _nw150 = Array((new BigNumber(7)).toNumber()).fill([]);
        _974_v9 = _nw150;
        let _index140 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_974_v9).length));
        (_974_v9)[_index140] = _963_v0;
        let _index141 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_963_v0).length));
        (_963_v0)[_index141] = false;
        let _975_v10;
        _975_v10 = new _dafny.CodePoint('c'.codePointAt(0));
        let _976_v11;
        let _nw151 = new _module.C1();
        _nw151.__ctor(_964_i0, p0, _974_v9, (_this).fm7(globalState), p3);
        _976_v11 = _nw151;
        let _977_v12;
        _977_v12 = _dafny.Map.Empty.slice().updateUnsafe(_976_v11,(_976_v11).f4);
        let _index142 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_963_v0).length));
        (_963_v0)[_index142] = (((_977_v12).contains(_976_v11)) ? ((_977_v12).get(_976_v11)) : ((_976_v11).f4));
        let _978_v13;
        _978_v13 = _dafny.Seq.of((_976_v11).f4, (_976_v11).f4);
        let _index143 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_974_v9).length));
        let _index144 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_963_v0).length));
        let _index145 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_963_v0).length));
        let _rhs142 = _963_v0;
        let _rhs143 = (_this).fm7(globalState);
        let _rhs144 = !((((_976_v11).f4) ? (!(false)) : ((_976_v11).f4)));
        let _rhs145 = new _dafny.CodePoint('k'.codePointAt(0));
        let _rhs146 = (((_978_v13)[_module.__default.safeIndex(p0, new BigNumber((_978_v13).length))]) ? (!(p3) || ((_this).f4)) : (p3));
        let _lhs84 = _974_v9;
        let _lhs85 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_974_v9).length));
        let _lhs86 = _963_v0;
        let _lhs87 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_963_v0).length));
        let _lhs88 = _963_v0;
        let _lhs89 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_963_v0).length));
        _lhs84[_lhs85] = _rhs142;
        r0 = _rhs143;
        _lhs86[_lhs87] = _rhs144;
        _975_v10 = _rhs145;
        _lhs88[_lhs89] = _rhs146;
        let _979_v14;
        _979_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_976_v11).f4);
        let _980_v15;
        _980_v15 = _dafny.Seq.of(_979_v14, (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_979_v14).length),p3)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_965_v1).f12,(_976_v11).f4)), _979_v14);
        _980_v15 = _980_v15;
      }
      let _981_v16;
      _981_v16 = new _dafny.CodePoint('l'.codePointAt(0));
      let _982_i2;
      _982_i2 = _dafny.ZERO;
      L6: {
        while (_module.__default.fm0(_981_v16, globalState)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_982_i2)) {
              break L6;
            }
            _982_i2 = (_982_i2).plus(_dafny.ONE);
            let _983_v17;
            _983_v17 = _module.D10.create_DC21(_dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0));
            let _984_v18;
            _984_v18 = _dafny.Seq.of(_983_v17, _983_v17);
            let _985_v19;
            _985_v19 = _dafny.MultiSet.fromElements((new BigNumber(390)).minus(p0), p0, ((_dafny.ZERO).minus(new BigNumber((_984_v18).length))).minus(new BigNumber((p2).length)));
            r0 = (((_985_v19).contains((_dafny.ZERO).minus((p0).multipliedBy((_dafny.ZERO).minus(p0))))) ? ((_985_v19).get((_dafny.ZERO).minus((p0).multipliedBy((_dafny.ZERO).minus(p0))))) : (p0));
            let _986_v20;
            _986_v20 = _dafny.Seq.UnicodeFromString("ulsxdir");
            _986_v20 = _dafny.Seq.Concat(_986_v20, ((!((_this).f4)) ? (_986_v20) : ((_this).f17)));
            let _987_v21;
            _987_v21 = _dafny.MultiSet.fromElements(p3);
            let _988_v22;
            let _nw152 = Array((new BigNumber(8)).toNumber()).fill([]);
            _988_v22 = _nw152;
            let _989_v23;
            let _nw153 = new _module.C1();
            _nw153.__ctor(p0, (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.of(_981_v16, _981_v16, _981_v16)).length)).plus((((_987_v21).contains(true)) ? ((_987_v21).get(true)) : (p0)))), ((p3) ? (_988_v22) : (_988_v22)), p0, !(p0).isEqualTo(new BigNumber(493)));
            _989_v23 = _nw153;
            r0 = ((_989_v23).fm9((_this).f4, (_this).f4, (_989_v23).f15, globalState)).multipliedBy((_989_v23).f15);
          }
        }
      }
      r0 = (_dafny.ZERO).minus(((p0).minus(p0)).minus(_module.__default.safeModuloInt(p0, p0)));
      let _990_v24;
      let _nw154 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _990_v24 = _nw154;
      let _991_v25;
      _991_v25 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus((_this).fm7(globalState))).plus(p0),_990_v24);
      _990_v24 = (((_991_v25).contains(new BigNumber(718))) ? ((_991_v25).get(new BigNumber(718))) : (_990_v24));
      let _992_v26;
      _992_v26 = _module.D2.create_DC7((((_this).f4) ? (p0) : (p0)));
      let _source17 = _992_v26;
      if (_source17.is_DC6) {
        let _993_v27;
        let _nw155 = new _module.C0();
        _nw155.__ctor(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(-298)), p0));
        _993_v27 = _nw155;
        let _994_v28;
        _994_v28 = true;
        _994_v28 = _994_v28;
        r0 = p0;
        r1 = (_993_v27).f12;
      } else if (_source17.is_DC7) {
        let _995___mcc_h0 = (_source17).cf6;
        let _996_cf6 = _995___mcc_h0;
        r1 = p0;
        let _997_v29;
        _997_v29 = false;
        _997_v29 = !(p0).isEqualTo(new BigNumber(((_this).f17).length));
        let _998_v30;
        _998_v30 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0));
        let _999_v31;
        _999_v31 = _dafny.Map.Empty.slice().updateUnsafe(_998_v30,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-304)), ((_1000_v16) => function (_1001_i3) {
          return _1000_v16;
        })(_981_v16)));
        let _1002_v32;
        _1002_v32 = _dafny.Seq.of(p3);
        let _1003_v33;
        _1003_v33 = _dafny.Seq.of(_module.__default.fm56((_this).f4, _996_cf6, p3, globalState), _1002_v32, _1002_v32);
        let _1004_v34;
        _1004_v34 = _dafny.MultiSet.fromElements(_1002_v32);
        let _1005_v35;
        let _1006_v36;
        let _out27;
        let _out28;
        let _outcollector12 = _module.__default.m0(_999_v31, (_dafny.MultiSet.FromArray(_1003_v33)).IsSubsetOf(_1004_v34), _998_v30, (new BigNumber((p2).length)).isEqualTo(p0), globalState);
        _out27 = _outcollector12[0];
        _out28 = _outcollector12[1];
        _1005_v35 = _out27;
        _1006_v36 = _out28;
        let _1007_v37;
        let _nw156 = Array((new BigNumber(17)).toNumber()).fill([]);
        _1007_v37 = _nw156;
        let _1008_v38;
        let _nw157 = new _module.C1();
        _nw157.__ctor(p0, new BigNumber(143), _1007_v37, _1006_v36, !(false));
        _1008_v38 = _nw157;
        let _1009_v39;
        _1009_v39 = _dafny.MultiSet.fromElements((_this).f4, p3);
        let _1010_v40;
        _1010_v40 = _module.D6.create_DC16(false, p3, _dafny.Seq.update(_dafny.Seq.of(_996_cf6), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1005_v35,_1008_v38)).length), new BigNumber((_dafny.Seq.of(_996_cf6)).length)), new BigNumber((_1009_v39).cardinality())));
        let _1011_v41;
        _1011_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(_996_cf6, (_this).f4, globalState),_1010_v40);
        let _1012_v42;
        _1012_v42 = _dafny.Set.fromElements(_963_v0);
        _1011_v41 = (_1011_v41).update(new BigNumber(((_1012_v42).Union(_1012_v42)).length), _1010_v40);
      } else {
        let _1013___mcc_h1 = (_source17).cf5;
        let _1014_cf5 = _1013___mcc_h1;
        let _1015_v43;
        _1015_v43 = _dafny.MultiSet.fromElements(true);
        let _1016_v44;
        _1016_v44 = _dafny.Seq.of(p0, (((_1015_v43).contains(p3)) ? ((_1015_v43).get(p3)) : ((_dafny.ZERO).minus(p0))), p0);
        let _rhs147 = _963_v0;
        let _rhs148 = p0;
        let _rhs149 = (_1016_v44)[_module.__default.safeIndex(_module.__default.safeModuloInt(p0, p0), new BigNumber((_1016_v44).length))];
        _1014_cf5 = _rhs147;
        r0 = _rhs148;
        r1 = _rhs149;
        let _1017_v45;
        _1017_v45 = true;
        _1017_v45 = !((_this).f4);
        let _1018_v46;
        _1018_v46 = _dafny.Map.Empty.slice().updateUnsafe(_990_v24,!((_this).f4));
        let _1019_v47;
        _1019_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1017_v45,_1018_v46);
        _1019_v47 = (_1019_v47).update(p3, ((_1017_v45) ? (_1018_v46) : (_1018_v46)));
        let _1020_v48;
        _1020_v48 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat((_this).f17, _dafny.Seq.UnicodeFromString("x")), _module.__default.safeIndex(new BigNumber(298), new BigNumber((_dafny.Seq.Concat((_this).f17, _dafny.Seq.UnicodeFromString("x"))).length)), _981_v16),(_1015_v43).IsSubsetOf(_dafny.MultiSet.fromElements(p3)));
        _1020_v48 = (_1020_v48).update(((false) ? (_dafny.Seq.UnicodeFromString("jpo")) : (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("qlx"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("qlx")).length)), _981_v16))), (_this).f4);
      }
      if (((_this).f4) && (p3)) {
        let _index146 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_963_v0).length));
        (_963_v0)[_index146] = p3;
        let _index147 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_963_v0).length));
        (_963_v0)[_index147] = (_this).f4;
        let _1021_v49;
        _1021_v49 = _dafny.Seq.of(p0, p0);
        let _1022_v50;
        _1022_v50 = _dafny.Seq.of((((p2).contains((_1021_v49)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_963_v0)[_module.__default.safeIndex(new BigNumber(936), new BigNumber((_963_v0).length))], p3)).length)), new BigNumber((_1021_v49).length))])) ? ((p2).get((_1021_v49)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_963_v0)[_module.__default.safeIndex(new BigNumber(936), new BigNumber((_963_v0).length))], p3)).length)), new BigNumber((_1021_v49).length))])) : (p0)), p0);
        _1022_v50 = _dafny.Seq.update(_1021_v49, _module.__default.safeIndex((_this).fm7(globalState), new BigNumber((_1021_v49).length)), p0);
        let _1023_v51;
        let _nw158 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1023_v51 = _nw158;
        let _index148 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1023_v51).length));
        (_1023_v51)[_index148] = (_this).f17;
        let _1024_v52;
        _1024_v52 = _dafny.Seq.of(!((_963_v0)[_module.__default.safeIndex(new BigNumber(936), new BigNumber((_963_v0).length))]));
        let _index149 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1023_v51).length));
        (_1023_v51)[_index149] = _dafny.Seq.Concat((((_1024_v52)[_module.__default.safeIndex(p0, new BigNumber((_1024_v52).length))]) ? ((_this).f17) : (_dafny.Seq.UnicodeFromString("dehlggigm"))), (_this).f17);
        let _1025_v53;
        let _nw159 = Array((new BigNumber(15)).toNumber());
        _nw159[(_dafny.ZERO).toNumber()] = _981_v16;
        _nw159[(_dafny.ONE).toNumber()] = _981_v16;
        _nw159[(new BigNumber(2)).toNumber()] = _981_v16;
        _nw159[(new BigNumber(3)).toNumber()] = _981_v16;
        _nw159[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
        _nw159[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
        _nw159[(new BigNumber(6)).toNumber()] = _981_v16;
        _nw159[(new BigNumber(7)).toNumber()] = _981_v16;
        _nw159[(new BigNumber(8)).toNumber()] = ((_1023_v51)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_1023_v51).length))])[_module.__default.safeIndex(p0, new BigNumber(((_1023_v51)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_1023_v51).length))]).length))];
        _nw159[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
        _nw159[(new BigNumber(10)).toNumber()] = _module.__default.fm1(p0, p3, globalState);
        _nw159[(new BigNumber(11)).toNumber()] = _981_v16;
        _nw159[(new BigNumber(12)).toNumber()] = _981_v16;
        _nw159[(new BigNumber(13)).toNumber()] = _module.__default.fm1(p0, (_963_v0)[_module.__default.safeIndex(new BigNumber(936), new BigNumber((_963_v0).length))], globalState);
        _nw159[(new BigNumber(14)).toNumber()] = _981_v16;
        _1025_v53 = _nw159;
        _1025_v53 = _1025_v53;
        let _1026_v54;
        let _init19 = ((_1027_p3) => function (_1028_i4) {
          return _1027_p3;
        })(p3);
        let _nw160 = Array((new BigNumber(2)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw160.length); _i0_19++) {
          _nw160[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _1026_v54 = _nw160;
        let _1029_v55;
        _1029_v55 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p0, new BigNumber(135)),_1026_v54);
        _1029_v55 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1026_v54);
      } else {
        let _1030_v56;
        _1030_v56 = false;
        let _1031_v57;
        _1031_v57 = _dafny.Seq.of((_this).f4);
        let _1032_v58;
        _1032_v58 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
        let _1033_v59;
        _1033_v59 = _dafny.Seq.of(new BigNumber((_1032_v58).length));
        let _1034_v60;
        _1034_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1030_v56,new BigNumber((_1033_v59).length));
        let _1035_v61;
        _1035_v61 = _dafny.MultiSet.fromElements((_1031_v57)[_module.__default.safeIndex((((_1034_v60).contains(p3)) ? ((_1034_v60).get(p3)) : (p0)), new BigNumber((_1031_v57).length))], p3);
        let _1036_v62;
        let _nw161 = Array((new BigNumber(2)).toNumber());
        _nw161[(_dafny.ZERO).toNumber()] = _1035_v61;
        _nw161[(_dafny.ONE).toNumber()] = _1035_v61;
        _1036_v62 = _nw161;
        let _1037_v63;
        let _nw162 = Array((new BigNumber(7)).toNumber());
        _nw162[(_dafny.ZERO).toNumber()] = _1036_v62;
        _nw162[(_dafny.ONE).toNumber()] = _1036_v62;
        _nw162[(new BigNumber(2)).toNumber()] = _1036_v62;
        _nw162[(new BigNumber(3)).toNumber()] = _1036_v62;
        _nw162[(new BigNumber(4)).toNumber()] = _1036_v62;
        _nw162[(new BigNumber(5)).toNumber()] = (((_this).f4) ? (_1036_v62) : (_1036_v62));
        _nw162[(new BigNumber(6)).toNumber()] = _1036_v62;
        _1037_v63 = _nw162;
        let _index150 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1037_v63).length));
        (_1037_v63)[_index150] = _1036_v62;
        let _index151 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1037_v63).length));
        let _rhs150 = (_dafny.ZERO).minus(p0);
        let _rhs151 = _963_v0;
        let _rhs152 = !((_this).f4);
        let _rhs153 = _1036_v62;
        let _lhs90 = _1037_v63;
        let _lhs91 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1037_v63).length));
        r1 = _rhs150;
        _963_v0 = _rhs151;
        _1030_v56 = _rhs152;
        _lhs90[_lhs91] = _rhs153;
        let _1038_v64;
        let _nw163 = Array((new BigNumber(20)).toNumber()).fill([]);
        _1038_v64 = _nw163;
        let _1039_v65;
        let _nw164 = new _module.C1();
        _nw164.__ctor((_1033_v59)[_module.__default.safeIndex(p0, new BigNumber((_1033_v59).length))], _module.__default.safeModuloInt(p0, p0), _1038_v64, _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_1031_v57).length), new BigNumber(9))).cardinality()), p0), p3);
        _1039_v65 = _nw164;
        _1030_v56 = _1030_v56;
        _1030_v56 = !_dafny.areEqual(_dafny.Seq.update(_1033_v59, _module.__default.safeIndex((_1039_v65).f15, new BigNumber((_1033_v59).length)), p0), _dafny.Seq.update(_dafny.Seq.Concat(_1033_v59, _module.__default.fm57((_1039_v65).f15, (_dafny.ZERO).minus(_1039_v65.f16), globalState)), _module.__default.safeIndex(_1039_v65.f16, new BigNumber((_dafny.Seq.Concat(_1033_v59, _module.__default.fm57((_1039_v65).f15, (_dafny.ZERO).minus(_1039_v65.f16), globalState))).length)), _1039_v65.f16));
        _1030_v56 = (_1039_v65.f16).isLessThanOrEqualTo(_1039_v65.f16);
      }
      r0 = new BigNumber(-259);
      r1 = (_dafny.ZERO).minus(p0);
      return [r0, r1];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f5 = [];
      this._f6 = _dafny.ZERO;
      this._f4 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    set f5(value) {
      let _this = this;
      _this._f5 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f5, f6, f4) {
      let _this = this;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f4 = f4;
      return;
    }
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(!((_this).f6).isEqualTo((_dafny.ZERO).minus((_this).f6))) || ((_this).f4);
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(476);
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements((_this).f6, (_this).f6)).Difference(_dafny.MultiSet.fromElements((_this).f6, (_this).f6));
    };
    fm7(globalState) {
      let _this = this;
      return (_this).f6;
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _1040_v0;
      let _init20 = ((_1041_p0) => function (_1042_i1) {
        return _module.__default.safeDivisionInt(_1042_i1, _1041_p0);
      })(p0);
      let _nw165 = Array((new BigNumber(22)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw165.length); _i0_20++) {
        _nw165[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _1040_v0 = _nw165;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1040_v0).length))) {
        let _1043_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1043_i0)) && ((_1043_i0).isLessThan(new BigNumber((_1040_v0).length))))) {
          (_1040_v0)[(_1043_i0)] = (_1043_i0).minus(p0);
        }
      }
      let _1044_v1;
      _1044_v1 = _dafny.Seq.of(false);
      let _1045_v2;
      _1045_v2 = _dafny.MultiSet.fromElements(_1044_v1);
      let _1046_v3;
      _1046_v3 = _module.D6.create_DC15(_1045_v2);
      _1046_v3 = _module.__default.fm33((_this).f4, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dqdyv")).length)), ((!((_this).f4)) ? (false) : (!((_this).f4))), globalState);
      let _1047_v4;
      _1047_v4 = _dafny.Seq.of(p0);
      let _1048_v5;
      _1048_v5 = _dafny.Set.fromElements(p0, new BigNumber((_1047_v4).length));
      let _hi5 = (p0).minus(new BigNumber((_1048_v5).length));
      for (let _1049_i2 = p0; _1049_i2.isLessThan(_hi5); _1049_i2 = _1049_i2.plus(_dafny.ONE)) {
        let _1050_v6;
        _1050_v6 = _dafny.Seq.UnicodeFromString("sj");
        let _1051_v7;
        _1051_v7 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(871)), ((_1052_p0) => function (_1053_i3) {
          return _1052_p0;
        })(p0))),_1050_v6);
        let _1054_v8;
        _1054_v8 = _dafny.MultiSet.fromElements(new BigNumber((_1050_v6).length));
        let _1055_v9;
        let _1056_v10;
        let _out29;
        let _out30;
        let _outcollector13 = _module.__default.m0((_module.__default.fm34(_1049_i2, p0, !((_this).f4), (_this).f4, globalState)).Merge(_1051_v7), (_this).f4, _1054_v8, (_this).f4, globalState);
        _out29 = _outcollector13[0];
        _out30 = _outcollector13[1];
        _1055_v9 = _out29;
        _1056_v10 = _out30;
        _1056_v10 = ((_this).f6).multipliedBy(((!((_this).f4)) ? (p0) : ((((_1054_v8).contains(_1049_i2)) ? ((_1054_v8).get(_1049_i2)) : (p0)))));
        let _1057_v11;
        let _nw166 = new _module.C0();
        _nw166.__ctor(new BigNumber(887));
        _1057_v11 = _nw166;
        let _1058_v12;
        _1058_v12 = _module.__default.fm35(_1049_i2, globalState);
        let _source18 = _1058_v12;
        let _1059___mcc_h0 = _source18;
        let _1060_cf23 = _1059___mcc_h0;
        let _1061_v13;
        _1061_v13 = _dafny.Map.Empty.slice().updateUnsafe((_1057_v11).f12,_dafny.Seq.Create(_module.__default.abs(new BigNumber(832)), ((_1062_p0) => function (_1063_i5) {
          return _1062_p0;
        })(p0)));
        let _1064_v14;
        let _nw167 = Array((new BigNumber(28)).toNumber());
        _nw167[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), ((_1065_v11) => function (_1066_i4) {
          return (_1065_v11).f12;
        })(_1057_v11));
        _nw167[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1047_v4, _1047_v4);
        _nw167[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1047_v4, _1047_v4);
        _nw167[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1047_v4, _dafny.Seq.update(_1047_v4, _module.__default.safeIndex(new BigNumber(523), new BigNumber((_1047_v4).length)), _1056_v10));
        _nw167[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat((((_1061_v13).contains((_1057_v11).f12)) ? ((_1061_v13).get((_1057_v11).f12)) : (_1047_v4)), _1047_v4);
        _nw167[(new BigNumber(5)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1047_v4, _dafny.Seq.of(_1049_i2, (_this).fm7(globalState)));
        _nw167[(new BigNumber(7)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(8)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(9)).toNumber()] = (((_this).f4) ? (_1047_v4) : (_dafny.Seq.of(_1049_i2)));
        _nw167[(new BigNumber(10)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(11)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(12)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(13)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(14)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(386)), ((_1067_v10) => function (_1068_i6) {
          return _1067_v10;
        })(_1056_v10));
        _nw167[(new BigNumber(16)).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus((_1057_v11).f12));
        _nw167[(new BigNumber(17)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(18)).toNumber()] = _module.__default.fm36(true, globalState);
        _nw167[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(new BigNumber((_1054_v8).cardinality()), (_dafny.ZERO).minus(_1056_v10));
        _nw167[(new BigNumber(20)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(21)).toNumber()] = _dafny.Seq.of(_1049_i2, (_1057_v11).f12, new BigNumber((_1044_v1).length), _1056_v10);
        _nw167[(new BigNumber(22)).toNumber()] = _1047_v4;
        _nw167[(new BigNumber(23)).toNumber()] = _dafny.Seq.update(_1047_v4, _module.__default.safeIndex((_dafny.ZERO).minus(_1049_i2), new BigNumber((_1047_v4).length)), new BigNumber((_1050_v6).length));
        _nw167[(new BigNumber(24)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(633)), function (_1069_i7) {
          return new BigNumber((_dafny.Seq.of((_this).f4, false)).length);
        });
        _nw167[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), ((_1070_p0) => function (_1071_i8) {
          return _1070_p0;
        })(p0)), _dafny.Seq.of(new BigNumber(-466)));
        _nw167[(new BigNumber(26)).toNumber()] = _module.__default.fm36((_this).fm8(new BigNumber((_dafny.MultiSet.fromElements(_1055_v9, (_this).f4, _1055_v9)).cardinality()), false, _1050_v6, _1047_v4, globalState), globalState);
        _nw167[(new BigNumber(27)).toNumber()] = _1047_v4;
        _1064_v14 = _nw167;
        let _index152 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1064_v14).length));
        (_1064_v14)[_index152] = _1047_v4;
        let _index153 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1064_v14).length));
        (_1064_v14)[_index153] = _1047_v4;
        _1056_v10 = (_1057_v11).f12;
        let _1072_v15;
        _1072_v15 = _dafny.Seq.of(_1050_v6);
        let _1073_v16;
        _1073_v16 = _dafny.Seq.of(_1050_v6, (_1072_v15)[_module.__default.safeIndex((_this).f6, new BigNumber((_1072_v15).length))], _1050_v6);
        let _1074_v17;
        _1074_v17 = _module.D3.create_DC10(new BigNumber(((_1073_v16)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_1056_v10)).length), new BigNumber((_1073_v16).length))]).length));
        _1074_v17 = _1074_v17;
        _1055_v9 = _1055_v9;
      }
      let _1075_v18;
      let _nw168 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
      _1075_v18 = _nw168;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1075_v18).length))) {
        let _1076_i9 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1076_i9)) && ((_1076_i9).isLessThan(new BigNumber((_1075_v18).length))))) {
          (_1075_v18)[(_1076_i9)] = _1044_v1;
        }
      }
      let _1077_v19;
      _1077_v19 = true;
      _1077_v19 = (_this).f4;
      let _source19 = _module.D0.create_DC2(new BigNumber(-988));
      if (_source19.is_DC1) {
        let _1078_v20;
        _1078_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(p0, _1077_v19, globalState),p0);
        let _1079_v21;
        _1079_v21 = _dafny.Seq.UnicodeFromString("vtke");
        _1078_v20 = (_1078_v20).update(new BigNumber((_1079_v21).length), _module.__default.safeDivisionInt(new BigNumber((_1048_v5).length), (_dafny.ZERO).minus(p0)));
        _1077_v19 = (_this).f4;
        let _1080_v22;
        _1080_v22 = new BigNumber(-943);
        let _1081_v23;
        _1081_v23 = _dafny.Set.fromElements(_1040_v0, _1040_v0);
        _1080_v22 = new BigNumber((_1081_v23).length);
        _1077_v19 = (_this).f4;
      } else if (_source19.is_DC2) {
        let _1082___mcc_h1 = (_source19).cf1;
        let _1083_cf1 = _1082___mcc_h1;
        let _1084_v24;
        _1084_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1048_v5,p0);
        _1083_cf1 = _module.__default.safeDivisionInt((((_1084_v24).contains(_1048_v5)) ? ((_1084_v24).get(_1048_v5)) : (_1083_cf1)), (_this).f6);
        let _1085_v25;
        _1085_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f6);
        let _1086_v26;
        _1086_v26 = _module.D10.create_DC21(((!((_this).f4)) ? (_1085_v25) : (_1085_v25)));
        let _source20 = _1086_v26;
        if (_source20.is_DC22) {
          let _1087___mcc_h3 = (_source20).cf30;
          let _1088___mcc_h4 = (_source20).cf31;
          let _1089___mcc_h5 = (_source20).cf32;
          let _1090___mcc_h6 = (_source20).cf33;
          let _1091___mcc_h7 = (_source20).cf34;
          let _1092_cf34 = _1091___mcc_h7;
          let _1093_cf33 = _1090___mcc_h6;
          let _1094_cf32 = _1089___mcc_h5;
          let _1095_cf31 = _1088___mcc_h4;
          let _1096_cf30 = _1087___mcc_h3;
          let _1097_v27;
          _1097_v27 = new _dafny.CodePoint('t'.codePointAt(0));
          let _1098_v28;
          let _out31;
          _out31 = (_this).m9(_module.__default.fm0(_1097_v27, globalState), globalState);
          _1098_v28 = _out31;
          _1083_cf1 = ((_1083_cf1).multipliedBy(_1083_cf1)).plus(new BigNumber(-560));
          let _1099_v29;
          let _nw169 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
          _1099_v29 = _nw169;
          let _1100_v30;
          _1100_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1099_v29,_1077_v19);
          let _1101_v31;
          _1101_v31 = _module.D11.create_DC23(_1099_v29);
          _1100_v30 = (_1100_v30).update((_1101_v31).dtor_cf35, (_this).f4);
          let _1102_v32;
          _1102_v32 = _dafny.MultiSet.fromElements(p0);
          _1077_v19 = ((_1077_v19) ? (_dafny.Seq.IsPrefixOf(_dafny.Seq.of((_this).f6), _1047_v4)) : ((_1102_v32).IsSubsetOf(_1102_v32)));
        } else {
          let _1103___mcc_h8 = (_source20).cf29;
          let _1104_cf29 = _1103___mcc_h8;
          _1045_v2 = _1045_v2;
          let _index154 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_1040_v0).length));
          (_1040_v0)[_index154] = (_this).f6;
          let _1105_v33;
          _1105_v33 = _dafny.Set.fromElements(_1077_v19);
          let _1106_v34;
          _1106_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1105_v33);
          let _1107_v35;
          _1107_v35 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1108_v36;
          let _nw170 = new _module.C0();
          _nw170.__ctor((_this).f6);
          _1108_v36 = _nw170;
          let _1109_v37;
          _1109_v37 = _dafny.MultiSet.fromElements(p0, (_dafny.ZERO).minus((p0).multipliedBy((_this).f6)), (_1083_cf1).multipliedBy(new BigNumber(386)), p0, (_this).fm9((_this).f4, (_this).f4, (_1108_v36).f12, globalState));
          let _1110_v38;
          _1110_v38 = _dafny.Seq.UnicodeFromString("wcnfyfu");
          let _1111_v39;
          _1111_v39 = _dafny.MultiSet.fromElements(_1109_v37);
          let _index155 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_1040_v0).length));
          let _rhs154 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_1106_v34).contains((_this).f6)) ? ((_1106_v34).get((_this).f6)) : (_1105_v33)),_dafny.Map.Empty.slice().updateUnsafe(_1107_v35,_1108_v36))).length);
          let _rhs155 = (((_1109_v37).contains(new BigNumber((_1110_v38).length))) ? ((_1109_v37).get(new BigNumber((_1110_v38).length))) : (_1083_cf1));
          let _rhs156 = (((_1111_v39).contains((_1109_v37).Intersect(_1109_v37))) ? ((_1111_v39).get((_1109_v37).Intersect(_1109_v37))) : (p0));
          let _rhs157 = _1083_cf1;
          let _lhs92 = _1040_v0;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_1040_v0).length));
          _1083_cf1 = _rhs154;
          _1083_cf1 = _rhs155;
          _1083_cf1 = _rhs156;
          _lhs92[_lhs93] = _rhs157;
          let _1112_v40;
          _1112_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber(-194), new BigNumber(341)),_dafny.Seq.UnicodeFromString("mou"));
          _1112_v40 = (_1112_v40).update((_1040_v0)[_module.__default.safeIndex(new BigNumber(483), new BigNumber((_1040_v0).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(70)), ((_1113_v35) => function (_1114_i10) {
            return _1113_v35;
          })(_1107_v35)));
          let _1115_v41;
          _1115_v41 = _dafny.Map.Empty.slice().updateUnsafe(true,_1107_v35);
          let _1116_v42;
          _1116_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1077_v19,(_this).f4);
          let _1117_v43;
          _1117_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1083_cf1,new BigNumber((_1116_v42).length));
          let _1118_v44;
          _1118_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,false);
          _1115_v41 = (_1115_v41).update(((_1040_v0)[_module.__default.safeIndex(new BigNumber(483), new BigNumber((_1040_v0).length))]).isLessThanOrEqualTo((((_1117_v43).contains(_1083_cf1)) ? ((_1117_v43).get(_1083_cf1)) : (new BigNumber((_1118_v44).length)))), _1107_v35);
        }
        _1077_v19 = _1077_v19;
        let _source21 = _1046_v3;
        if (_source21.is_DC16) {
          let _1119___mcc_h9 = (_source21).cf20;
          let _1120___mcc_h10 = (_source21).cf21;
          let _1121___mcc_h11 = (_source21).cf22;
          let _1122_cf22 = _1121___mcc_h11;
          let _1123_cf21 = _1120___mcc_h10;
          let _1124_cf20 = _1119___mcc_h9;
          let _1125_v45;
          let _nw171 = new _module.C1();
          _nw171.__ctor((_dafny.ZERO).minus(_1083_cf1), (_this).f6, _this.f5, new BigNumber(490), false);
          _1125_v45 = _nw171;
          let _1126_v46;
          _1126_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1125_v45,(_1125_v45).f4);
          let _1127_v47;
          _1127_v47 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(p0, _1123_cf21, globalState),_1126_v46);
          _1127_v47 = _1127_v47;
          let _1128_v48;
          _1128_v48 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1129_v49;
          _1129_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1128_v48,_1128_v48);
          let _1130_v50;
          _1130_v50 = _dafny.MultiSet.fromElements(_1129_v49);
          let _1131_v51;
          _1131_v51 = _dafny.Set.fromElements(_1077_v19, false);
          let _1132_v52;
          _1132_v52 = _dafny.MultiSet.fromElements((_this).f6);
          _1124_cf20 = ((_1132_v52).Union(_1132_v52)).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_1130_v50).cardinality()), new BigNumber((_1131_v51).length), p0, _1083_cf1));
          let _1133_v53;
          _1133_v53 = _dafny.MultiSet.fromElements(_1077_v19);
          _1044_v1 = (_module.D4.create_DC13((_1125_v45).f6, _1077_v19, new BigNumber((_1133_v53).cardinality()), _1044_v1)).dtor_cf17;
          let _1134_v54;
          _1134_v54 = _dafny.Seq.UnicodeFromString("wfi");
          let _1135_v55;
          _1135_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1132_v52,_1134_v54);
          let _1136_v56;
          let _1137_v57;
          let _out32;
          let _out33;
          let _outcollector14 = _module.__default.m0(_1135_v55, _dafny.Seq.IsPrefixOf(_1122_cf22, _1047_v4), _1132_v52, (_1125_v45).f4, globalState);
          _out32 = _outcollector14[0];
          _out33 = _outcollector14[1];
          _1136_v56 = _out32;
          _1137_v57 = _out33;
        } else {
          let _1138___mcc_h12 = (_source21).cf19;
          let _1139_cf19 = _1138___mcc_h12;
          _1083_cf1 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).fm9(_1077_v19, true, (_this).f6, globalState), (((_1085_v25).contains((_this).f4)) ? ((_1085_v25).get((_this).f4)) : (p0))));
          let _1140_v58;
          let _nw172 = Array((new BigNumber(2)).toNumber());
          _nw172[(_dafny.ZERO).toNumber()] = _1077_v19;
          _nw172[(_dafny.ONE).toNumber()] = (_this).f4;
          _1140_v58 = _nw172;
          let _1141_v59;
          _1141_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1140_v58,_1083_cf1);
          let _1142_v60;
          _1142_v60 = _this.f5;
          let _1143_v61;
          let _nw173 = new _module.C1();
          _nw173.__ctor((_this).f6, (_this).f6, (_1142_v60), (_this).f6, _1077_v19);
          _1143_v61 = _nw173;
          let _1144_v62;
          _1144_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1140_v58,_1143_v61);
          let _1145_v63;
          _1145_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1083_cf1,new BigNumber((_1144_v62).length));
          let _1146_v64;
          let _nw174 = new _module.C1();
          _nw174.__ctor(new BigNumber((_1145_v63).length), (_this).f6, _this.f5, (_dafny.ZERO).minus(new BigNumber((_1047_v4).length)), (_this).f4);
          _1146_v64 = _nw174;
          let _1147_v65;
          _1147_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1146_v64,_1140_v58);
          _1141_v59 = (_1141_v59).update((((_1147_v65).contains(_1146_v64)) ? ((_1147_v65).get(_1146_v64)) : (_1140_v58)), _1143_v61.f16);
          let _1148_v66;
          _1148_v66 = _dafny.MultiSet.fromElements(false, !(_1077_v19));
          let _1149_v67;
          _1149_v67 = _module.D12.create_DC28(p0, new BigNumber((_1148_v66).cardinality()), false, (_1077_v19) || (_1077_v19), (_this).fm9(_1077_v19, _1077_v19, _1143_v61.f16, globalState));
          _1149_v67 = _1149_v67;
          let _1150_v68;
          _1150_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1143_v61.f16,new _dafny.CodePoint('m'.codePointAt(0)));
          (_1143_v61).f16 = _module.__default.safeModuloInt(p0, new BigNumber((_1150_v68).length));
        }
      } else {
        let _1151___mcc_h2 = (_source19).cf0;
        let _1152_cf0 = _1151___mcc_h2;
        let _1153_v69;
        _1153_v69 = _dafny.Seq.UnicodeFromString("qlh");
        let _1154_v70;
        _1154_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_module.__default.fm39(_1153_v69, _module.D12.create_DC27(_dafny.Seq.of(_dafny.Seq.of((_this).f4))), _module.__default.fm0(new _dafny.CodePoint('r'.codePointAt(0)), globalState), globalState)),_1153_v69);
        let _1155_v71;
        let _1156_v72;
        let _out34;
        let _out35;
        let _outcollector15 = _module.__default.m0((_1154_v70).Merge(_1154_v70), (_this).f4, _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1047_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(68)), ((_1157_v69) => function (_1158_i11) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1157_v69,new _dafny.CodePoint('i'.codePointAt(0)))).length);
        })(_1153_v69)))), !(_1077_v19), globalState);
        _out34 = _outcollector15[0];
        _out35 = _outcollector15[1];
        _1155_v71 = _out34;
        _1156_v72 = _out35;
        if (false) {
          _1155_v71 = _1077_v19;
          _1040_v0 = _1040_v0;
          let _1159_v74;
          let _nw175 = new _module.C1();
          _nw175.__ctor((_dafny.ZERO).minus(_1156_v72), (_1047_v4)[_module.__default.safeIndex((_this).f6, new BigNumber((_1047_v4).length))], _this.f5, new BigNumber(((_1048_v5).Union(function () {
            let _coll41 = new _dafny.Set();
            for (const _compr_41 of _dafny.IntegerRange(new BigNumber(779), new BigNumber(-489))) {
              let _1160_v73 = _compr_41;
              if (((new BigNumber(779)).isLessThanOrEqualTo(_1160_v73)) && ((_1160_v73).isLessThan(new BigNumber(-489)))) {
                _coll41.add((_1160_v73).minus(p0));
              }
            }
            return _coll41;
          }())).length), (_1156_v72).isLessThan(_1156_v72));
          _1159_v74 = _nw175;
          _1077_v19 = (_this).f4;
          let _1161_v75;
          let _nw176 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1161_v75 = _nw176;
          _1161_v75 = _1161_v75;
        } else {
          _1152_cf0 = _1156_v72;
          let _1162_v76;
          _1162_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1044_v1).length),(_this).f6);
          let _1163_v77;
          _1163_v77 = _dafny.MultiSet.fromElements(new BigNumber(226));
          let _1164_v78;
          _1164_v78 = _dafny.Seq.of(_1163_v77);
          _1155_v71 = (_1044_v1)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_1162_v76).length), new BigNumber(((_1164_v78)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_1164_v78).length))]).cardinality())), new BigNumber((_1044_v1).length))];
          let _1165_v79;
          _1165_v79 = _dafny.Map.Empty.slice().updateUnsafe((_1152_cf0).isLessThan(_1152_cf0),_1155_v71);
          _1077_v19 = (((_1165_v79).contains((_this).f4)) ? ((_1165_v79).get((_this).f4)) : (_1077_v19));
          let _1166_v80;
          _1166_v80 = new _dafny.CodePoint('v'.codePointAt(0));
          _1166_v80 = _1166_v80;
          let _1167_v81;
          let _nw177 = Array((new BigNumber(24)).toNumber()).fill(false);
          _1167_v81 = _nw177;
          _1167_v81 = (_module.D2.create_DC5(_1167_v81)).dtor_cf5;
        }
        let _1168_v82;
        _1168_v82 = new _dafny.CodePoint('l'.codePointAt(0));
        _1168_v82 = _1168_v82;
        let _1169_v85;
        _1169_v85 = _dafny.Set.fromElements((_this).f4, _1155_v71, _module.__default.fm0(_1168_v82, globalState));
        let _1170_v86;
        let _nw178 = Array((new BigNumber(6)).toNumber());
        _nw178[(_dafny.ZERO).toNumber()] = _1048_v5;
        _nw178[(_dafny.ONE).toNumber()] = function () {
          let _coll42 = new _dafny.Set();
          for (const _compr_42 of (_1047_v4).Elements) {
            let _1171_v83 = _compr_42;
            if (_dafny.Seq.contains(_1047_v4, _1171_v83)) {
              _coll42.add(_module.__default.safeDivisionInt(_1171_v83, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-932)), function (_1172_i12) {
                return new _dafny.CodePoint('v'.codePointAt(0));
              })).length)));
            }
          }
          return _coll42;
        }();
        _nw178[(new BigNumber(2)).toNumber()] = _1048_v5;
        _nw178[(new BigNumber(3)).toNumber()] = ((_1155_v71) ? (_1048_v5) : (function () {
          let _coll43 = new _dafny.Set();
          for (const _compr_43 of (_1047_v4).Elements) {
            let _1173_v84 = _compr_43;
            if (_dafny.Seq.contains(_1047_v4, _1173_v84)) {
              _coll43.add(_module.__default.safeDivisionInt(_1173_v84, (_dafny.ZERO).minus(new BigNumber(-223))));
            }
          }
          return _coll43;
        }()));
        _nw178[(new BigNumber(4)).toNumber()] = _1048_v5;
        _nw178[(new BigNumber(5)).toNumber()] = _module.__default.fm35(new BigNumber((_1169_v85).length), globalState);
        _1170_v86 = _nw178;
        let _index156 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_1170_v86).length));
        (_1170_v86)[_index156] = _1048_v5;
        let _1174_v87;
        _1174_v87 = _dafny.MultiSet.fromElements(p0);
        let _index157 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_1170_v86).length));
        (_1170_v86)[_index157] = (_1048_v5).Intersect(_dafny.Set.fromElements((((_1174_v87).contains((_this).f6)) ? ((_1174_v87).get((_this).f6)) : ((_this).f6)), p0, (_this).f6, _1152_cf0, p0));
      }
      let _1175_v88;
      _1175_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1077_v19,(_this).f6);
      r0 = _1175_v88;
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = [];
      let _1176_v0;
      _1176_v0 = _dafny.MultiSet.fromElements((_this).f6);
      let _1177_v1;
      let _nw179 = Array((new BigNumber(3)).toNumber());
      _nw179[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((((_1176_v0).contains((_this).f6)) ? ((_1176_v0).get((_this).f6)) : ((_this).f6)));
      _nw179[(_dafny.ONE).toNumber()] = (_this).f6;
      _nw179[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(((_this).f6).multipliedBy((_this).fm7(globalState)));
      _1177_v1 = _nw179;
      let _index158 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
      (_1177_v1)[_index158] = (_this).f6;
      let _index159 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
      (_1177_v1)[_index159] = (_this).f6;
      let _1178_v2;
      _1178_v2 = _dafny.Map.Empty.slice().updateUnsafe((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))],_1177_v1);
      let _1179_v3;
      _1179_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(442),(((_1178_v2).contains((_this).f6)) ? ((_1178_v2).get((_this).f6)) : (_1177_v1)));
      let _1180_v4;
      _1180_v4 = _dafny.Seq.UnicodeFromString("ivk");
      let _1181_v5;
      _1181_v5 = _dafny.Seq.of(_1177_v1, _1177_v1, _1177_v1, _1177_v1);
      let _1182_v6;
      let _nw180 = Array((new BigNumber(24)).toNumber());
      _nw180[(_dafny.ZERO).toNumber()] = _1177_v1;
      _nw180[(_dafny.ONE).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(2)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(3)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(4)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(5)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(6)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(7)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(8)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(9)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(10)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(11)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(12)).toNumber()] = ((!((_this).f4)) ? (_1177_v1) : (_1177_v1));
      _nw180[(new BigNumber(13)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(14)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(15)).toNumber()] = (((_this).f4) ? (_1177_v1) : (_1177_v1));
      _nw180[(new BigNumber(16)).toNumber()] = (((_1179_v3).contains(new BigNumber((_1180_v4).length))) ? ((_1179_v3).get(new BigNumber((_1180_v4).length))) : (_1177_v1));
      _nw180[(new BigNumber(17)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(18)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(19)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(20)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(21)).toNumber()] = _1177_v1;
      _nw180[(new BigNumber(22)).toNumber()] = (_1181_v5)[_module.__default.safeIndex((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))], new BigNumber((_1181_v5).length))];
      _nw180[(new BigNumber(23)).toNumber()] = _1177_v1;
      _1182_v6 = _nw180;
      let _index160 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1182_v6).length));
      (_1182_v6)[_index160] = _1177_v1;
      let _1183_v7;
      _1183_v7 = _module.D2.create_DC7(new BigNumber((_module.__default.fm37(p0, (_this).f4, (_this).f6, p0, globalState)).length));
      let _1184_v8;
      _1184_v8 = _dafny.MultiSet.fromElements(_1183_v7, _1183_v7);
      let _1185_v9;
      _1185_v9 = _dafny.Set.fromElements(p0, p0);
      let _1186_v10;
      _1186_v10 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1184_v8).cardinality())).isEqualTo((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))]),!(new BigNumber((_1185_v9).length)).isEqualTo((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))]));
      let _1187_v11;
      _1187_v11 = true;
      let _1188_v12;
      _1188_v12 = _module.D16.create_DC32(_1177_v1);
      let _1189_v13;
      _1189_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1177_v1,(_1188_v12).dtor_cf51);
      let _index161 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1182_v6).length));
      let _index162 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
      let _rhs158 = (((_1189_v13).contains(_1177_v1)) ? ((_1189_v13).get(_1177_v1)) : (_1177_v1));
      let _rhs159 = (_this).f6;
      let _rhs160 = (_1186_v10).Merge((_1186_v10).Merge(_1186_v10));
      let _rhs161 = !((_this).f4) || ((_this).f4);
      let _lhs94 = _1182_v6;
      let _lhs95 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1182_v6).length));
      let _lhs96 = _1177_v1;
      let _lhs97 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
      _lhs94[_lhs95] = _rhs158;
      _lhs96[_lhs97] = _rhs159;
      _1186_v10 = _rhs160;
      _1187_v11 = _rhs161;
      let _1190_v14;
      let _nw181 = Array((new BigNumber(18)).toNumber()).fill([]);
      _1190_v14 = _nw181;
      let _1191_v15;
      _1191_v15 = _dafny.Seq.of((_this).f4);
      let _1192_v16;
      _1192_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))]);
      let _1193_v17;
      _1193_v17 = _dafny.Seq.of(_1192_v16, _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))])), _1192_v16);
      let _index163 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
      let _index164 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
      let _rhs162 = (_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))];
      let _rhs163 = new BigNumber(((((_this).f4) ? (_dafny.Seq.Concat(_1191_v15, _1191_v15)) : (_dafny.Seq.Concat(_1191_v15, _1191_v15)))).length);
      let _rhs164 = !_dafny.Seq.contains(_dafny.Seq.Concat(_1193_v17, _1193_v17), _1192_v16);
      let _rhs165 = _1187_v11;
      let _rhs166 = _1190_v14;
      let _lhs98 = _1177_v1;
      let _lhs99 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
      let _lhs100 = _1177_v1;
      let _lhs101 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
      _lhs98[_lhs99] = _rhs162;
      _lhs100[_lhs101] = _rhs163;
      _1187_v11 = _rhs164;
      _1187_v11 = _rhs165;
      _1190_v14 = _rhs166;
      let _1194_v18;
      let _nw182 = new _module.C2();
      _nw182.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(587)), function (_1195_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), p0);
      _1194_v18 = _nw182;
      let _1196_v19;
      _1196_v19 = _dafny.Map.Empty.slice().updateUnsafe((_1194_v18).f4,_1194_v18);
      let _1197_v20;
      _1197_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1194_v18);
      let _1198_v21;
      _1198_v21 = _dafny.Seq.of(_1194_v18, (((_1196_v19).contains(p0)) ? ((_1196_v19).get(p0)) : ((((_1197_v20).contains(new BigNumber(-228))) ? ((_1197_v20).get(new BigNumber(-228))) : (_1194_v18)))), _1194_v18, _1194_v18, _1194_v18);
      let _index165 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
      (_1177_v1)[_index165] = (new BigNumber((_1198_v21).length)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-701)), function (_1199_i1) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length));
      let _1200_v22;
      let _init21 = function (_1201_i2) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("q"), _dafny.Seq.UnicodeFromString("fag"));
      };
      let _nw183 = Array((new BigNumber(14)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw183.length); _i0_21++) {
        _nw183[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _1200_v22 = _nw183;
      let _1202_v23;
      _1202_v23 = _dafny.Seq.of(_1180_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(819)), function (_1203_i3) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }));
      let _1204_v24;
      _1204_v24 = new _dafny.CodePoint('f'.codePointAt(0));
      let _index166 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_1200_v22).length));
      (_1200_v22)[_index166] = _dafny.Seq.Concat(_dafny.Seq.update((_1202_v23)[_module.__default.safeIndex((_this).f6, new BigNumber((_1202_v23).length))], _module.__default.safeIndex((_this).f6, new BigNumber(((_1202_v23)[_module.__default.safeIndex((_this).f6, new BigNumber((_1202_v23).length))]).length)), _1204_v24), _dafny.Seq.UnicodeFromString("nusipv"));
      let _index167 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_1200_v22).length));
      (_1200_v22)[_index167] = _dafny.Seq.Create(_module.__default.abs(_dafny.ZERO), function (_1205_i4) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      });
      if (_module.__default.fm0(_1204_v24, globalState)) {
        let _1206_v25;
        let _nw184 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1206_v25 = _nw184;
        let _index168 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_1190_v14).length));
        (_1190_v14)[_index168] = _1206_v25;
        let _index169 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_1190_v14).length));
        (_1190_v14)[_index169] = _1206_v25;
        _1187_v11 = !(p0);
        let _1207_v26;
        _1207_v26 = _module.D11.create_DC26((_1194_v18).f4);
        let _1208_v27;
        _1208_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1207_v26,(_1194_v18).f4);
        let _1209_v28;
        _1209_v28 = _dafny.Seq.of(_1208_v27);
        let _1210_v29;
        _1210_v29 = _module.D18.create_DC35(_1209_v28);
        let _1211_v30;
        _1211_v30 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(451)), ((_1212_v24) => function (_1213_i5) {
          return _1212_v24;
        })(_1204_v24))).length), (_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))]);
        let _1214_v31;
        _1214_v31 = _dafny.Seq.of((_1211_v30)[_module.__default.safeIndex((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))], new BigNumber((_1211_v30).length))]);
        let _1215_v32;
        _1215_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1211_v30).length),_1186_v10);
        let _1216_v33;
        _1216_v33 = _dafny.Seq.of(_1215_v32, _1215_v32);
        let _1217_v34;
        let _nw185 = new _module.C1();
        _nw185.__ctor((_dafny.ZERO).minus(new BigNumber(((_1210_v29).dtor_cf58).length)), new BigNumber(((_1216_v33)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_1187_v11)).length), new BigNumber((_1216_v33).length))]).length), _this.f5, ((_1214_v31)[_module.__default.safeIndex((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))], new BigNumber((_1214_v31).length))]).minus((_1194_v18).fm7(globalState)), true);
        _1217_v34 = _nw185;
        let _1218_v35;
        _1218_v35 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray(_1191_v15));
        let _1219_v36;
        let _nw186 = Array((new BigNumber(27)).toNumber());
        _nw186[(_dafny.ZERO).toNumber()] = _module.__default.fm0(_1204_v24, globalState);
        _nw186[(_dafny.ONE).toNumber()] = !(((_this).f6).isLessThanOrEqualTo((_1217_v34).f15));
        _nw186[(new BigNumber(2)).toNumber()] = (_1218_v35).IsSubsetOf(_1218_v35);
        _nw186[(new BigNumber(3)).toNumber()] = (_this).f4;
        _nw186[(new BigNumber(4)).toNumber()] = false;
        _nw186[(new BigNumber(5)).toNumber()] = _1187_v11;
        _nw186[(new BigNumber(6)).toNumber()] = (_this).f4;
        _nw186[(new BigNumber(7)).toNumber()] = (_this).f4;
        _nw186[(new BigNumber(8)).toNumber()] = true;
        _nw186[(new BigNumber(9)).toNumber()] = (_1194_v18).f4;
        _nw186[(new BigNumber(10)).toNumber()] = (_this).f4;
        _nw186[(new BigNumber(11)).toNumber()] = p0;
        _nw186[(new BigNumber(12)).toNumber()] = _1187_v11;
        _nw186[(new BigNumber(13)).toNumber()] = (_1185_v9).IsSubsetOf(_1185_v9);
        _nw186[(new BigNumber(14)).toNumber()] = (_1194_v18).f4;
        _nw186[(new BigNumber(15)).toNumber()] = (_this).f4;
        _nw186[(new BigNumber(16)).toNumber()] = p0;
        _nw186[(new BigNumber(17)).toNumber()] = ((_1194_v18).f4) || (p0);
        _nw186[(new BigNumber(18)).toNumber()] = p0;
        _nw186[(new BigNumber(19)).toNumber()] = (_1194_v18).f4;
        _nw186[(new BigNumber(20)).toNumber()] = (_this).f4;
        _nw186[(new BigNumber(21)).toNumber()] = p0;
        _nw186[(new BigNumber(22)).toNumber()] = p0;
        _nw186[(new BigNumber(23)).toNumber()] = _1187_v11;
        _nw186[(new BigNumber(24)).toNumber()] = true;
        _nw186[(new BigNumber(25)).toNumber()] = (_1194_v18).f4;
        _nw186[(new BigNumber(26)).toNumber()] = p0;
        _1219_v36 = _nw186;
        let _1220_v37;
        _1220_v37 = _dafny.Set.fromElements((_1200_v22)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_1200_v22).length))]);
        let _index170 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_1219_v36).length));
        (_1219_v36)[_index170] = (_1220_v37).IsProperSubsetOf(_1220_v37);
        let _index171 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
        let _index172 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_1219_v36).length));
        let _rhs167 = _module.__default.safeModuloInt((_this).f6, ((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))]).multipliedBy(_1217_v34.f16));
        let _rhs168 = p0;
        let _rhs169 = (_1194_v18).f4;
        let _lhs102 = _1177_v1;
        let _lhs103 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
        let _lhs104 = _1219_v36;
        let _lhs105 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_1219_v36).length));
        _lhs102[_lhs103] = _rhs167;
        _1187_v11 = _rhs168;
        _lhs104[_lhs105] = _rhs169;
        _1187_v11 = !((_module.D12.create_DC28(_module.__default.fm3(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1191_v15, _module.__default.safeIndex((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))], new BigNumber((_1191_v15).length)), !(p0)), _module.__default.safeIndex(_1217_v34.f16, new BigNumber((_dafny.Seq.update(_1191_v15, _module.__default.safeIndex((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))], new BigNumber((_1191_v15).length)), !(p0))).length)), (_1219_v36)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1219_v36).length))])).length), (_this).f4, globalState), _1217_v34.f16, (_1194_v18).f4, (_1219_v36)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_1219_v36).length))], (_this).f6)).dtor_cf46);
      } else {
        let _index173 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
        (_1177_v1)[_index173] = (_this).f6;
        let _index174 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length));
        (_1177_v1)[_index174] = (_this).f6;
        _1187_v11 = true;
        let _1221_v38;
        let _nw187 = new _module.C2();
        _nw187.__ctor((_1200_v22)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_1200_v22).length))], _1187_v11);
        _1221_v38 = _nw187;
        let _1222_v39;
        _1222_v39 = _dafny.MultiSet.fromElements(_1221_v38, _1221_v38, _1221_v38);
        let _1223_v40;
        let _nw188 = Array((new BigNumber(10)).toNumber());
        _nw188[(_dafny.ZERO).toNumber()] = !((_1222_v39).update(_1221_v38, _module.__default.abs((_this).f6))).contains(_1221_v38);
        _nw188[(_dafny.ONE).toNumber()] = true;
        _nw188[(new BigNumber(2)).toNumber()] = _module.__default.fm0(_1204_v24, globalState);
        _nw188[(new BigNumber(3)).toNumber()] = (((_1194_v18).f4) ? (_1187_v11) : (true));
        _nw188[(new BigNumber(4)).toNumber()] = !(_1185_v9).equals(_1185_v9);
        _nw188[(new BigNumber(5)).toNumber()] = (_1176_v0).equals(_1176_v0);
        _nw188[(new BigNumber(6)).toNumber()] = ((_this).f6).isLessThan((_this).f6);
        _nw188[(new BigNumber(7)).toNumber()] = _1187_v11;
        _nw188[(new BigNumber(8)).toNumber()] = (_this).f4;
        _nw188[(new BigNumber(9)).toNumber()] = (_1194_v18).f4;
        _1223_v40 = _nw188;
        let _index175 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_1223_v40).length));
        (_1223_v40)[_index175] = _1187_v11;
        let _index176 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_1223_v40).length));
        (_1223_v40)[_index176] = (_1194_v18).f4;
        let _1224_v41;
        let _nw189 = new _module.C1();
        _nw189.__ctor((_1177_v1)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1177_v1).length))], new BigNumber(806), _this.f5, ((_this).f6).multipliedBy((_this).f6), _1187_v11);
        _1224_v41 = _nw189;
      }
      r0 = _1204_v24;
      let _1225_v42;
      let _nw190 = Array((new BigNumber(23)).toNumber()).fill(_module.D1.Default());
      _1225_v42 = _nw190;
      r1 = _1225_v42;
      return [r0, r1];
    }
    m9(p0, globalState) {
      let _this = this;
      let r0 = false;
      r0 = !(p0);
      let _1226_v0;
      let _init22 = function (_1227_i0) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(773)), function (_1228_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("uqvrar"));
      };
      let _nw191 = Array((new BigNumber(22)).toNumber());
      for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw191.length); _i0_22++) {
        _nw191[_i0_22] = _init22(new BigNumber(_i0_22));
      }
      _1226_v0 = _nw191;
      let _index177 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length));
      (_1226_v0)[_index177] = _dafny.Seq.UnicodeFromString("t");
      let _1229_v1;
      _1229_v1 = _dafny.Seq.UnicodeFromString("w");
      let _index178 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length));
      (_1226_v0)[_index178] = _dafny.Seq.Concat(_module.__default.fm37(!((_this).f4), (_this).f4, new BigNumber((_1229_v1).length), (_this).f4, globalState), _1229_v1);
      let _1230_v2;
      _1230_v2 = new _dafny.CodePoint('v'.codePointAt(0));
      let _index179 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length));
      (_1226_v0)[_index179] = _dafny.Seq.update(_1229_v1, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f6), new BigNumber((_1229_v1).length)), _1230_v2);
      let _1231_v3;
      _1231_v3 = _dafny.Seq.of((_this).f6, (_this).f6, (_this).f6);
      let _1232_v4;
      let _nw192 = new _module.C1();
      _nw192.__ctor((_dafny.ZERO).minus((_this).f6), new BigNumber((_dafny.Seq.Concat(_1231_v3, _1231_v3)).length), _this.f5, (_this).f6, !((_this).f4));
      _1232_v4 = _nw192;
      if (!((_this).f4)) {
        let _1233_v5;
        let _init23 = function (_1234_i2) {
          return (_this).f4;
        };
        let _nw193 = Array((new BigNumber(5)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw193.length); _i0_23++) {
          _nw193[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _1233_v5 = _nw193;
        let _1235_v6;
        _1235_v6 = _dafny.Map.Empty.slice().updateUnsafe(true,_1233_v5);
        let _1236_v7;
        _1236_v7 = _dafny.Seq.of(_1233_v5, _1233_v5, _1233_v5);
        _1233_v5 = (((_1235_v6).contains((_this).f4)) ? ((_1235_v6).get((_this).f4)) : ((_1236_v7)[_module.__default.safeIndex((_1232_v4).f15, new BigNumber((_1236_v7).length))]));
        let _1237_v8;
        let _nw194 = new _module.C2();
        _nw194.__ctor(_1229_v1, p0);
        _1237_v8 = _nw194;
        (_1232_v4).f16 = (_1232_v4).f15;
        let _1238_v9;
        _1238_v9 = _dafny.Map.Empty.slice().updateUnsafe((_1232_v4).f15,(_this).f4);
        _1238_v9 = (_dafny.Map.Empty.slice().updateUnsafe((_1232_v4).f15,(_this).f4)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1232_v4).f15,p0));
        let _1239_v10;
        _1239_v10 = _module.D10.create_DC22(_dafny.Seq.Create(_module.__default.abs(new BigNumber(791)), ((_1240_v2) => function (_1241_i3) {
  return _1240_v2;
})(_1230_v2)), false, (_this).f4, _1233_v5, new BigNumber(184));
        let _1242_v11;
        _1242_v11 = _module.D1.create_DC4((_1239_v10).dtor_cf34, (new BigNumber((_dafny.Seq.UnicodeFromString("lroyicek")).length)).minus(new BigNumber(455)));
        _1242_v11 = _1242_v11;
      } else {
        r0 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("vtossvsa"), _dafny.Seq.Concat((_1226_v0)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length))], (_1226_v0)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length))]));
        r0 = (new BigNumber(((_1226_v0)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length))]).length)).isLessThanOrEqualTo((((_this).f4) ? (_1232_v4.f16) : ((_this).f6)));
        let _1243_v12;
        let _nw195 = Array((new BigNumber(2)).toNumber());
        _nw195[(_dafny.ZERO).toNumber()] = _1230_v2;
        _nw195[(_dafny.ONE).toNumber()] = _1230_v2;
        _1243_v12 = _nw195;
        let _1244_v13;
        _1244_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1243_v12);
        let _1245_v14;
        _1245_v14 = _dafny.Seq.of(_1243_v12);
        let _1246_v15;
        let _nw196 = Array((new BigNumber(8)).toNumber());
        _nw196[(_dafny.ZERO).toNumber()] = _1243_v12;
        _nw196[(_dafny.ONE).toNumber()] = _1243_v12;
        _nw196[(new BigNumber(2)).toNumber()] = _1243_v12;
        _nw196[(new BigNumber(3)).toNumber()] = _1243_v12;
        _nw196[(new BigNumber(4)).toNumber()] = _1243_v12;
        _nw196[(new BigNumber(5)).toNumber()] = (((_1244_v13).contains(p0)) ? ((_1244_v13).get(p0)) : ((_1245_v14)[_module.__default.safeIndex((_dafny.ZERO).minus((_1232_v4).f15), new BigNumber((_1245_v14).length))]));
        _nw196[(new BigNumber(6)).toNumber()] = _1243_v12;
        _nw196[(new BigNumber(7)).toNumber()] = _1243_v12;
        _1246_v15 = _nw196;
        let _index180 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1246_v15).length));
        (_1246_v15)[_index180] = _1243_v12;
        let _1247_v16;
        _1247_v16 = _dafny.MultiSet.fromElements(_1232_v4.f16);
        let _index181 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length));
        let _index182 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1246_v15).length));
        let _rhs170 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bsbsmmac"), _dafny.Seq.UnicodeFromString("ps")), (_1226_v0)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length))]);
        let _rhs171 = (((_1247_v16).contains(_1232_v4.f16)) ? ((_1247_v16).get(_1232_v4.f16)) : ((_1232_v4).f15));
        let _rhs172 = _1243_v12;
        let _lhs106 = _1226_v0;
        let _lhs107 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length));
        let _lhs108 = _1232_v4;
        let _lhs109 = _1246_v15;
        let _lhs110 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1246_v15).length));
        _lhs106[_lhs107] = _rhs170;
        _lhs108.f16 = _rhs171;
        _lhs109[_lhs110] = _rhs172;
        let _1248_v18;
        let _init24 = ((_1249_v3) => function (_1250_i4) {
          return function () {
            let _coll44 = new _dafny.Set();
            for (const _compr_44 of (_1249_v3).Elements) {
              let _1251_v17 = _compr_44;
              if (_dafny.Seq.contains(_1249_v3, _1251_v17)) {
                _coll44.add((_1251_v17).plus(new BigNumber(-142)));
              }
            }
            return _coll44;
          }();
        })(_1231_v3);
        let _nw197 = Array((_dafny.ONE).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw197.length); _i0_24++) {
          _nw197[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _1248_v18 = _nw197;
        let _1252_v19;
        let _nw198 = Array((new BigNumber(4)).toNumber());
        _nw198[(_dafny.ZERO).toNumber()] = _1248_v18;
        _nw198[(_dafny.ONE).toNumber()] = ((p0) ? (_1248_v18) : (_1248_v18));
        _nw198[(new BigNumber(2)).toNumber()] = ((p0) ? (_1248_v18) : (_1248_v18));
        _nw198[(new BigNumber(3)).toNumber()] = _1248_v18;
        _1252_v19 = _nw198;
        let _index183 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_1252_v19).length));
        (_1252_v19)[_index183] = _1248_v18;
        let _index184 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_1252_v19).length));
        (_1252_v19)[_index184] = (((_this).f4) ? (_1248_v18) : (_1248_v18));
        if (true) {
          let _1253_v20;
          let _init25 = ((_1254_v16) => function (_1255_i5) {
            return _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber((_1254_v16).cardinality()));
          })(_1247_v16);
          let _nw199 = Array((new BigNumber(25)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw199.length); _i0_25++) {
            _nw199[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1253_v20 = _nw199;
          let _1256_v21;
          _1256_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1232_v4.f16);
          let _index185 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1253_v20).length));
          (_1253_v20)[_index185] = _1256_v21;
          let _1257_v22;
          _1257_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4);
          let _1258_v23;
          _1258_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(((_1257_v22).contains((_this).f4)) ? ((_1257_v22).get((_this).f4)) : (p0)));
          let _1259_v24;
          _1259_v24 = _dafny.Seq.of(_1257_v22, _1258_v23, _1257_v22);
          let _index186 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1246_v15).length));
          let _index187 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1253_v20).length));
          let _rhs173 = _dafny.areEqual(_1259_v24, _1259_v24);
          let _rhs174 = _1243_v12;
          let _rhs175 = _1256_v21;
          let _rhs176 = _1247_v16;
          let _lhs111 = _1246_v15;
          let _lhs112 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1246_v15).length));
          let _lhs113 = _1253_v20;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1253_v20).length));
          r0 = _rhs173;
          _lhs111[_lhs112] = _rhs174;
          _lhs113[_lhs114] = _rhs175;
          _1247_v16 = _rhs176;
          let _1260_v25;
          let _nw200 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
          _1260_v25 = _nw200;
          let _1261_v26;
          _1261_v26 = _dafny.MultiSet.fromElements(p0);
          let _1262_v27;
          _1262_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1261_v26).cardinality()),new BigNumber(548));
          let _index188 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_1260_v25).length));
          (_1260_v25)[_index188] = _1262_v27;
          let _1263_v28;
          _1263_v28 = _dafny.Set.fromElements(p0);
          let _index189 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_1260_v25).length));
          (_1260_v25)[_index189] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1263_v28).length),new BigNumber(580));
          r0 = (_this).f4;
          (_1232_v4).f16 = _1232_v4.f16;
          (_1232_v4).f16 = (_1232_v4.f16).plus(((_1232_v4).f15).multipliedBy(_1232_v4.f16));
        } else {
          let _1264_v29;
          _1264_v29 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), ((_1265_v2) => function (_1266_i6) {
            return _1265_v2;
          })(_1230_v2)),new BigNumber(307));
          let _1267_v30;
          _1267_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1264_v29);
          let _1268_v31;
          let _init26 = function (_1269_i7) {
            return true;
          };
          let _nw201 = Array((new BigNumber(16)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw201.length); _i0_26++) {
            _nw201[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1268_v31 = _nw201;
          let _1270_v32;
          let _out36;
          _out36 = (_1232_v4).m11(p0, p0, (((_1267_v30).contains(p0)) ? ((_1267_v30).get(p0)) : (_dafny.Map.Empty.slice().updateUnsafe((_1226_v0)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_1226_v0).length))],_1232_v4.f16))), _1268_v31, globalState);
          _1270_v32 = _out36;
          let _1271_v33;
          _1271_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1232_v4).f15,p0);
          let _1272_v34;
          _1272_v34 = _dafny.Seq.of((((_1271_v33).contains((_1232_v4).f15)) ? ((_1271_v33).get((_1232_v4).f15)) : (false)));
          let _1273_v35;
          _1273_v35 = _dafny.Seq.of(_1272_v34, _1272_v34);
          let _1274_v36;
          _1274_v36 = _module.D12.create_DC27(_1273_v35);
          let _pat_let_tv16 = _1273_v35;
          let _pat_let_tv17 = _1229_v1;
          let _pat_let_tv18 = _1273_v35;
          let _pat_let_tv19 = _1273_v35;
          let _pat_let_tv20 = _1232_v4;
          let _pat_let_tv21 = _1273_v35;
          let _pat_let_tv22 = _1273_v35;
          _1274_v36 = function (_pat_let22_0) {
            return function (_1277_dt__update__tmp_h1) {
              return function (_pat_let25_0) {
                return function (_1278_dt__update_hcf42_h1) {
                  return _module.D12.create_DC27(_1278_dt__update_hcf42_h1);
                }(_pat_let25_0);
              }(_pat_let_tv22);
            }(_pat_let22_0);
          }(function (_pat_let23_0) {
            return function (_1275_dt__update__tmp_h0) {
              return function (_pat_let24_0) {
                return function (_1276_dt__update_hcf42_h0) {
                  return _module.D12.create_DC27(_1276_dt__update_hcf42_h0);
                }(_pat_let24_0);
              }(_dafny.Seq.update(_pat_let_tv16, _module.__default.safeIndex(new BigNumber((_pat_let_tv17).length), new BigNumber((_pat_let_tv18).length)), (_pat_let_tv19)[_module.__default.safeIndex(_pat_let_tv20.f16, new BigNumber((_pat_let_tv21).length))]));
            }(_pat_let23_0);
          }(_1274_v36));
          let _arr5 = _this.f5;
          let _index190 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_this.f5).length));
          _arr5[_index190] = _1268_v31;
          let _arr6 = _this.f5;
          let _index191 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_this.f5).length));
          _arr6[_index191] = _1268_v31;
          let _1279_v37;
          _1279_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1268_v31);
          let _1280_v38;
          _1280_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1230_v2,(((_1279_v37).contains((_this).f4)) ? ((_1279_v37).get((_this).f4)) : ((_this.f5)[_module.__default.safeIndex(new BigNumber(533), new BigNumber((_this.f5).length))])));
          let _1281_v39;
          _1281_v39 = _module.D19.create_DC37(_1280_v38);
          let _rhs177 = (_this).f4;
          let _rhs178 = (_1281_v39).dtor_cf62;
          r0 = _rhs177;
          _1280_v38 = _rhs178;
          let _1282_v40;
          let _1283_v41;
          let _out37;
          let _out38;
          let _outcollector16 = (_1232_v4).m2(p0, globalState);
          _out37 = _outcollector16[0];
          _out38 = _outcollector16[1];
          _1282_v40 = _out37;
          _1283_v41 = _out38;
        }
      }
      let _1284_v42;
      _1284_v42 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1285_v43;
      _1285_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1284_v42,_1230_v2);
      let _1286_v44;
      _1286_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_dafny.Seq.UnicodeFromString("mcrpadd"));
      let _1287_v45;
      let _1288_v46;
      let _1289_v47;
      let _1290_v48;
      let _out39;
      let _out40;
      let _out41;
      let _out42;
      let _outcollector17 = (_this).m10(_dafny.MultiSet.fromElements(new BigNumber((_1285_v43).length), (_this).f6, _1232_v4.f16), _dafny.Seq.IsProperPrefixOf(_1229_v1, (((_1286_v44).contains(_1232_v4.f16)) ? ((_1286_v44).get(_1232_v4.f16)) : (_dafny.Seq.UnicodeFromString("h")))), globalState);
      _out39 = _outcollector17[0];
      _out40 = _outcollector17[1];
      _out41 = _outcollector17[2];
      _out42 = _outcollector17[3];
      _1287_v45 = _out39;
      _1288_v46 = _out40;
      _1289_v47 = _out41;
      _1290_v48 = _out42;
      r0 = p0;
      return r0;
    }
    m10(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _1291_v0;
      let _nw202 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _1291_v0 = _nw202;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1291_v0).length))) {
        let _1292_i0 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1292_i0)) && ((_1292_i0).isLessThan(new BigNumber((_1291_v0).length))))) {
          (_1291_v0)[(_1292_i0)] = _module.__default.safeModuloInt(_1292_i0, (_this).f6);
        }
      }
      let _hi6 = new BigNumber(-365);
      for (let _1293_i1 = (_this).f6; _1293_i1.isLessThan(_hi6); _1293_i1 = _1293_i1.plus(_dafny.ONE)) {
        let _1294_v1;
        _1294_v1 = _module.D16.create_DC32(_1291_v0);
        let _source22 = _1294_v1;
        if (_source22.is_DC33) {
          let _1295___mcc_h0 = (_source22).cf52;
          let _1296___mcc_h1 = (_source22).cf53;
          let _1297___mcc_h2 = (_source22).cf54;
          let _1298___mcc_h3 = (_source22).cf55;
          let _1299___mcc_h4 = (_source22).cf56;
          let _1300_cf56 = _1299___mcc_h4;
          let _1301_cf55 = _1298___mcc_h3;
          let _1302_cf54 = _1297___mcc_h2;
          let _1303_cf53 = _1296___mcc_h1;
          let _1304_cf52 = _1295___mcc_h0;
          _1304_cf52 = !(true);
          let _1305_v2;
          _1305_v2 = _dafny.Seq.of(_1304_cf52, false);
          let _1306_v3;
          _1306_v3 = _module.D4.create_DC13((_this).f6, _1304_cf52, _1293_i1, _1305_v2);
          r0 = (((p1) ? (_1306_v3) : (_1306_v3))).dtor_cf15;
          let _1307_v4;
          let _nw203 = Array((new BigNumber(14)).toNumber()).fill(false);
          _1307_v4 = _nw203;
          _1307_v4 = _1307_v4;
          r2 = (_dafny.ZERO).minus(_1293_i1);
        } else {
          let _1308___mcc_h5 = (_source22).cf51;
          let _1309_cf51 = _1308___mcc_h5;
          let _1310_v5;
          let _nw204 = Array((new BigNumber(16)).toNumber()).fill(false);
          _1310_v5 = _nw204;
          let _index192 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1310_v5).length));
          (_1310_v5)[_index192] = (_this).f4;
          let _1311_v6;
          _1311_v6 = _module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_1293_i1)).length));
          let _1312_v7;
          _1312_v7 = _dafny.Set.fromElements(_1311_v6);
          let _1313_v8;
          _1313_v8 = _dafny.MultiSet.fromElements((_this).f4);
          let _1314_v10;
          _1314_v10 = _dafny.Set.fromElements(_1312_v7, _1312_v7, function () {
            let _coll45 = new _dafny.Set();
            for (const _compr_45 of (_dafny.Seq.of(_module.__default.fm38(new BigNumber((((_1313_v8).update((_this).f4, _module.__default.abs(_1293_i1))).update((_this).f4, _module.__default.abs(_1293_i1))).cardinality()), _1293_i1, p1, (_this).f6, globalState))).Elements) {
              let _1315_v9 = _compr_45;
              if (_dafny.Seq.contains(_dafny.Seq.of(_module.__default.fm38(new BigNumber((((_1313_v8).update((_this).f4, _module.__default.abs(_1293_i1))).update((_this).f4, _module.__default.abs(_1293_i1))).cardinality()), _1293_i1, p1, (_this).f6, globalState)), _1315_v9)) {
                _coll45.add(_1315_v9);
              }
            }
            return _coll45;
          }());
          let _1316_v11;
          _1316_v11 = _dafny.MultiSet.fromElements(new BigNumber(((_dafny.Set.fromElements(_1312_v7)).Difference(_1314_v10)).length));
          let _index193 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1310_v5).length));
          let _rhs179 = !((_this).f4) || ((_this).f4);
          let _rhs180 = _1316_v11;
          let _rhs181 = new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("skjggxj")).length))).length);
          let _lhs115 = _1310_v5;
          let _lhs116 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1310_v5).length));
          _lhs115[_lhs116] = _rhs179;
          _1316_v11 = _rhs180;
          r1 = _rhs181;
          let _index194 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1310_v5).length));
          (_1310_v5)[_index194] = ((_1310_v5)[_module.__default.safeIndex(new BigNumber(287), new BigNumber((_1310_v5).length))]) && (((_1310_v5)[_module.__default.safeIndex(new BigNumber(287), new BigNumber((_1310_v5).length))]) === (false));
          let _1317_v12;
          _1317_v12 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f6), (_this).f6);
          r2 = new BigNumber((_1317_v12).length);
          let _index195 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_1310_v5).length));
          (_1310_v5)[_index195] = p1;
        }
        let _1318_v13;
        let _nw205 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1318_v13 = _nw205;
        let _1319_v14;
        _1319_v14 = _module.D11.create_DC26((_this).f4);
        let _rhs182 = _1318_v13;
        let _rhs183 = _1319_v14;
        _1318_v13 = _rhs182;
        _1319_v14 = _rhs183;
        let _1320_v15;
        _1320_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p1);
        let _1321_v16;
        let _nw206 = new _module.C1();
        _nw206.__ctor((new BigNumber((_1320_v15).length)).multipliedBy(_1293_i1), _1293_i1, _this.f5, _1293_i1, !((_this).f4) || (!(false)));
        _1321_v16 = _nw206;
        let _1322_v17;
        _1322_v17 = new _dafny.CodePoint('b'.codePointAt(0));
        r3 = _module.__default.fm0(_1322_v17, globalState);
      }
      if ((_this).f4) {
        let _1323_v18;
        _1323_v18 = _module.D3.create_DC10(((_this).f6).multipliedBy((_this).f6));
        let _source23 = _1323_v18;
        if (_source23.is_DC9) {
          let _1324___mcc_h6 = (_source23).cf8;
          let _1325___mcc_h7 = (_source23).cf9;
          let _1326___mcc_h8 = (_source23).cf10;
          let _1327_cf10 = _1326___mcc_h8;
          let _1328_cf9 = _1325___mcc_h7;
          let _1329_cf8 = _1324___mcc_h6;
          r2 = (_this).f6;
          let _1330_v19;
          _1330_v19 = _dafny.Seq.UnicodeFromString("uig");
          let _1331_v20;
          _1331_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1330_v19);
          let _1332_v21;
          _1332_v21 = _module.D0.create_DC0(new BigNumber(767));
          let _1333_v22;
          _1333_v22 = new _dafny.CodePoint('t'.codePointAt(0));
          let _1334_v23;
          _1334_v23 = _dafny.Set.fromElements(_1330_v19);
          let _1335_v24;
          let _1336_v25;
          let _out43;
          let _out44;
          let _outcollector18 = _module.__default.m0(_1331_v20, _1329_cf8, (_this).fm6(_module.D0.create_DC0(new BigNumber((p0).cardinality())), _1332_v21, _dafny.Map.Empty.slice().updateUnsafe(_1328_cf9,_1333_v22), new BigNumber((_1334_v23).length), globalState), _1327_cf10, globalState);
          _out43 = _outcollector18[0];
          _out44 = _outcollector18[1];
          _1335_v24 = _out43;
          _1336_v25 = _out44;
          let _1337_v26;
          let _nw207 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _1337_v26 = _nw207;
          let _1338_v27;
          _1338_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1323_v18,_module.__default.fm58(globalState));
          let _1339_v28;
          _1339_v28 = _dafny.Seq.of((_this).f6);
          let _1340_v29;
          _1340_v29 = _dafny.Seq.of((_this).fm8(_1336_v25, false, _1330_v19, _1339_v28, globalState));
          let _1341_v30;
          _1341_v30 = _dafny.Seq.of(_1340_v29);
          let _1342_v31;
          _1342_v31 = _module.D12.create_DC27(_dafny.Seq.update(_1341_v30, _module.__default.safeIndex(new BigNumber((_1340_v29).length), new BigNumber((_1341_v30).length)), _1340_v29));
          let _1343_v32;
          _1343_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1342_v31);
          let _index196 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_1337_v26).length));
          (_1337_v26)[_index196] = (((_1338_v27).contains(_1323_v18)) ? ((_1338_v27).get(_1323_v18)) : (_1343_v32));
          let _1344_v33;
          _1344_v33 = _module.D20.create_DC40(_1343_v32);
          let _index197 = _module.__default.safeIndex(new BigNumber(776), new BigNumber((_1337_v26).length));
          (_1337_v26)[_index197] = (_1343_v32).Merge((_1343_v32).Merge((_module.D20.create_DC40((_1344_v33).dtor_cf66)).dtor_cf66));
          let _1345_v34;
          _1345_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1336_v25,_1327_cf10);
          let _1346_v35;
          let _nw208 = new _module.C1();
          _nw208.__ctor((new BigNumber((_1345_v34).length)).plus(_1328_cf9), _module.__default.safeModuloInt(new BigNumber((_1330_v19).length), _1336_v25), _this.f5, (_this).f6, _1335_v24);
          _1346_v35 = _nw208;
        } else if (_source23.is_DC10) {
          let _1347___mcc_h9 = (_source23).cf11;
          let _1348_cf11 = _1347___mcc_h9;
          r1 = new BigNumber(491);
          let _1349_v36;
          _1349_v36 = _dafny.Seq.UnicodeFromString("fvlmlhsge");
          r2 = new BigNumber((_dafny.Seq.Concat(_1349_v36, _1349_v36)).length);
          let _1350_v37;
          let _nw209 = new _module.C1();
          _nw209.__ctor((_this).f6, _module.__default.safeModuloInt(new BigNumber(753), (_this).f6), _this.f5, _module.__default.safeModuloInt((_this).f6, (_this).f6), (_this).f4);
          _1350_v37 = _nw209;
          _1291_v0 = _1291_v0;
        } else if (_source23.is_DC8) {
          let _1351___mcc_h10 = (_source23).cf7;
          let _1352_cf7 = _1351___mcc_h10;
          let _1353_v38;
          _1353_v38 = new _dafny.CodePoint('e'.codePointAt(0));
          let _1354_v39;
          _1354_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1353_v38,false);
          r0 = !(p1) || ((((_1354_v39).contains(_1353_v38)) ? ((_1354_v39).get(_1353_v38)) : (!((_this).f4))));
          let _1355_v40;
          _1355_v40 = _dafny.Seq.UnicodeFromString("honftv");
          r2 = ((_this).f6).minus((_this).fm9((_this).f4, (_this).f4, new BigNumber((_1355_v40).length), globalState));
          let _1356_v41;
          let _init27 = function (_1357_i2) {
            return true;
          };
          let _nw210 = Array((new BigNumber(26)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw210.length); _i0_27++) {
            _nw210[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1356_v41 = _nw210;
          let _1358_v42;
          _1358_v42 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_1356_v41);
          let _1359_v43;
          _1359_v43 = _module.D19.create_DC37(_1358_v42);
          let _1360_v44;
          _1360_v44 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f4) || (true),_1359_v43);
          _1360_v44 = (_1360_v44).update((_this).f4, _1359_v43);
          let _1361_v45;
          let _nw211 = new _module.C2();
          _nw211.__ctor(_1355_v40, (_this).f4);
          _1361_v45 = _nw211;
          let _1362_v46;
          _1362_v46 = _dafny.Seq.of(_1361_v45, _1361_v45, _1361_v45);
          let _1363_v47;
          _1363_v47 = _dafny.Seq.of(_1362_v46, _dafny.Seq.of(_1361_v45));
          r3 = !_dafny.areEqual(_dafny.Seq.update(_1363_v47, _module.__default.safeIndex((_this).f6, new BigNumber((_1363_v47).length)), _1362_v46), _1363_v47);
        } else {
          let _1364___mcc_h11 = (_source23).cf12;
          let _1365_cf12 = _1364___mcc_h11;
          let _1366_v48;
          _1366_v48 = _dafny.Seq.of((_this).f6, (_this).f6);
          _1366_v48 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), ((_1367_p1) => function (_1368_i3) {
            return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_1367_p1), _dafny.Seq.of((_this).f4, (_this).f4))).length);
          })(p1));
          let _1369_v49;
          _1369_v49 = _dafny.Seq.UnicodeFromString("bwcahcko");
          let _1370_v50;
          _1370_v50 = new _dafny.CodePoint('h'.codePointAt(0));
          let _1371_v51;
          _1371_v51 = _module.D9.create_DC19(_1369_v49);
          _1369_v49 = _dafny.Seq.Concat(_dafny.Seq.update((((_this).f4) ? (_dafny.Seq.UnicodeFromString("ylnncpc")) : (_1369_v49)), _module.__default.safeIndex((_this).f6, new BigNumber(((((_this).f4) ? (_dafny.Seq.UnicodeFromString("ylnncpc")) : (_1369_v49))).length)), _1370_v50), (_1371_v51).dtor_cf25);
          let _1372_v52;
          _1372_v52 = _module.D12.create_DC28((_this).f6, (_dafny.ZERO).minus((_this).f6), (_this).f4, p1, new BigNumber(((((_this).f4) ? (_dafny.Seq.UnicodeFromString("iunuhhphg")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(651)), ((_1373_v50) => function (_1374_i4) {
  return _1373_v50;
})(_1370_v50))))).length));
          let _rhs184 = _1372_v52;
          let _rhs185 = (true) || ((_this).f4);
          let _rhs186 = _1369_v49;
          _1372_v52 = _rhs184;
          r0 = _rhs185;
          _1369_v49 = _rhs186;
          r3 = p1;
        }
        let _1375_v53;
        _1375_v53 = _dafny.Seq.UnicodeFromString("ukokag");
        r1 = new BigNumber((_module.__default.fm37(p1, !(p1), new BigNumber((_1375_v53).length), (_this).fm8((_this).f6, p1, _1375_v53, _dafny.Seq.of(new BigNumber((_1375_v53).length), (_this).f6), globalState), globalState)).length);
        r2 = (_dafny.ZERO).minus((_this).f6);
        let _1376_v54;
        _1376_v54 = _dafny.Set.fromElements((_this).f4, p1);
        let _1377_v55;
        _1377_v55 = _dafny.Seq.of((_1376_v54).IsSubsetOf(_1376_v54));
        r0 = (_1377_v55)[_module.__default.safeIndex(((_this).f6).minus(new BigNumber(779)), new BigNumber((_1377_v55).length))];
        let _1378_v56;
        _1378_v56 = new _dafny.CodePoint('i'.codePointAt(0));
        _1378_v56 = _1378_v56;
      } else {
        r1 = (_this).f6;
        let _1379_v57;
        _1379_v57 = _dafny.Seq.UnicodeFromString("kcdy");
        let _1380_v58;
        _1380_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1379_v57);
        let _1381_v59;
        _1381_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1379_v57,_1380_v58);
        let _1382_v60;
        _1382_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1380_v58);
        _1381_v59 = (_1381_v59).update(_1379_v57, ((((_1382_v60).contains(new BigNumber((_1379_v57).length))) ? ((_1382_v60).get(new BigNumber((_1379_v57).length))) : (_1380_v58))).Merge(_1380_v58));
        let _1383_v61;
        _1383_v61 = _dafny.Seq.of((_this).f6);
        r3 = (_this).fm8((_this).f6, (p1) && ((_this).f4), _1379_v57, _1383_v61, globalState);
        let _index198 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_1291_v0).length));
        (_1291_v0)[_index198] = (_this).f6;
        let _index199 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_1291_v0).length));
        (_1291_v0)[_index199] = (_this).f6;
        let _1384_v62;
        let _nw212 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
        _1384_v62 = _nw212;
        let _1385_v63;
        _1385_v63 = _dafny.Set.fromElements((_this).f6, new BigNumber((_1379_v57).length));
        let _index200 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_1384_v62).length));
        (_1384_v62)[_index200] = _1385_v63;
        let _index201 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_1384_v62).length));
        (_1384_v62)[_index201] = _1385_v63;
      }
      let _1386_v64;
      _1386_v64 = _dafny.Seq.UnicodeFromString("iktm");
      let _hi7 = (_this).f6;
      for (let _1387_i5 = ((_this).f6).minus(new BigNumber((_1386_v64).length)); _1387_i5.isLessThan(_hi7); _1387_i5 = _1387_i5.plus(_dafny.ONE)) {
        let _1388_v65;
        _1388_v65 = _dafny.Seq.of((_this).f6);
        r2 = (_1388_v65)[_module.__default.safeIndex((_this).f6, new BigNumber((_1388_v65).length))];
        let _1389_v66;
        _1389_v66 = _dafny.MultiSet.fromElements((_this).f4, !((_this).f4));
        let _pat_let_tv23 = _1386_v64;
        let _source24 = function (_pat_let26_0) {
          return function (_1392_dt__update__tmp_h1) {
            return function (_pat_let29_0) {
              return function (_1393_dt__update_hcf47_h0) {
                return _module.D12.create_DC28((_1392_dt__update__tmp_h1).dtor_cf43, (_1392_dt__update__tmp_h1).dtor_cf44, (_1392_dt__update__tmp_h1).dtor_cf45, (_1392_dt__update__tmp_h1).dtor_cf46, _1393_dt__update_hcf47_h0);
              }(_pat_let29_0);
            }(new BigNumber((_pat_let_tv23).length));
          }(_pat_let26_0);
        }(function (_pat_let27_0) {
          return function (_1390_dt__update__tmp_h0) {
            return function (_pat_let28_0) {
              return function (_1391_dt__update_hcf43_h0) {
                return _module.D12.create_DC28(_1391_dt__update_hcf43_h0, (_1390_dt__update__tmp_h0).dtor_cf44, (_1390_dt__update__tmp_h0).dtor_cf45, (_1390_dt__update__tmp_h0).dtor_cf46, (_1390_dt__update__tmp_h0).dtor_cf47);
              }(_pat_let28_0);
            }((_this).f6);
          }(_pat_let27_0);
        }(_module.D12.create_DC28((_this).f6, new BigNumber(((p0).update(_1387_i5, _module.__default.abs(new BigNumber((_1389_v66).cardinality())))).cardinality()), p1, true, _1387_i5)));
        if (_source24.is_DC28) {
          let _1394___mcc_h12 = (_source24).cf43;
          let _1395___mcc_h13 = (_source24).cf44;
          let _1396___mcc_h14 = (_source24).cf45;
          let _1397___mcc_h15 = (_source24).cf46;
          let _1398___mcc_h16 = (_source24).cf47;
          let _1399_cf47 = _1398___mcc_h16;
          let _1400_cf46 = _1397___mcc_h15;
          let _1401_cf45 = _1396___mcc_h14;
          let _1402_cf44 = _1395___mcc_h13;
          let _1403_cf43 = _1394___mcc_h12;
          let _1404_v67;
          _1404_v67 = _module.D0.create_DC0((new BigNumber(879)).minus(_1399_cf47));
          _1404_v67 = _1404_v67;
          _1386_v64 = _1386_v64;
          let _1405_v68;
          _1405_v68 = _dafny.Set.fromElements((_this).f6, _1399_cf47);
          _1400_cf46 = (_1387_i5).isEqualTo((_dafny.ZERO).minus((new BigNumber((_1405_v68).length)).multipliedBy((_1388_v65)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1405_v68).length)), new BigNumber((_1388_v65).length))])));
          let _1406_v69;
          _1406_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1399_cf47,p1);
          let _1407_v71;
          _1407_v71 = _dafny.MultiSet.fromElements(((_1401_cf45) ? (_1406_v69) : (_1406_v69)), function () {
            let _coll46 = new _dafny.Map();
            for (const _compr_46 of _dafny.IntegerRange(new BigNumber(416), new BigNumber(-448))) {
              let _1408_v70 = _compr_46;
              if (((new BigNumber(416)).isLessThanOrEqualTo(_1408_v70)) && ((_1408_v70).isLessThan(new BigNumber(-448)))) {
                _coll46.push([_module.__default.safeModuloInt(_1408_v70, (_this).f6),_1400_cf46]);
              }
            }
            return _coll46;
          }(), _1406_v69);
          _1407_v71 = _1407_v71;
        } else {
          let _1409___mcc_h17 = (_source24).cf42;
          let _1410_cf42 = _1409___mcc_h17;
          r3 = !((new BigNumber((_1386_v64).length)).isLessThanOrEqualTo((_this).f6)) || ((_this).f4);
          r1 = _1387_i5;
          let _1411_v72;
          let _nw213 = Array((new BigNumber(3)).toNumber());
          _nw213[(_dafny.ZERO).toNumber()] = ((!(!(p1))) ? (!(p1)) : ((_this).f4));
          _nw213[(_dafny.ONE).toNumber()] = (_this).f4;
          _nw213[(new BigNumber(2)).toNumber()] = (_this).f4;
          _1411_v72 = _nw213;
          let _index202 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_1411_v72).length));
          (_1411_v72)[_index202] = p1;
          let _index203 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_1411_v72).length));
          let _rhs187 = p1;
          let _rhs188 = p1;
          let _lhs117 = _1411_v72;
          let _lhs118 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_1411_v72).length));
          r0 = _rhs187;
          _lhs117[_lhs118] = _rhs188;
          let _index204 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_1411_v72).length));
          (_1411_v72)[_index204] = (_1411_v72)[_module.__default.safeIndex(new BigNumber(154), new BigNumber((_1411_v72).length))];
        }
        let _1412_v73;
        _1412_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber((_1386_v64).length));
        let _1413_v74;
        _1413_v74 = _module.D10.create_DC21(_1412_v73);
        let _1414_v75;
        _1414_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1413_v74,new BigNumber((_1386_v64).length));
        _1414_v75 = _module.__default.fm59(globalState);
        r1 = (_dafny.ZERO).minus((_this).f6);
      }
      let _1415_v76;
      _1415_v76 = _dafny.Seq.of((_this).f6);
      _1415_v76 = ((p1) ? (_dafny.Seq.of(new BigNumber(-887), (_this).f6, (_this).f6, (_this).f6, (_this).f6)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(899)), function (_1416_i6) {
        return (_this).f6;
      })));
      let _1417_v77;
      _1417_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,((_this).f6).minus((_this).f6));
      _1417_v77 = (_module.__default.fm44(globalState)).Merge((_1417_v77).Merge(_1417_v77));
      r0 = true;
      r1 = (_this).f6;
      r2 = (_this).f6;
      r3 = !((_this).f4);
      return [r0, r1, r2, r3];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f5 = [];
      this._f6 = _dafny.ZERO;
      this._f4 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    set f5(value) {
      let _this = this;
      _this._f5 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f5, f6, f4) {
      let _this = this;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f4 = f4;
      return;
    }
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_this).f4);
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(new BigNumber(117), new BigNumber(842)), (((_this).f4) ? ((_this).f6) : ((_this).f6)), (_this).f6, _module.__default.safeDivisionInt((_this).f6, (_this).f6), (_this).f6);
    };
    fm7(globalState) {
      let _this = this;
      return new BigNumber((function (_source25) {
        if (_source25.is_DC20) {
          let _1418___mcc_h0 = (_source25).cf26;
          let _1419___mcc_h1 = (_source25).cf27;
          let _1420___mcc_h2 = (_source25).cf28;
          let _1421_cf28 = _1420___mcc_h2;
          let _1422_cf27 = _1419___mcc_h1;
          let _1423_cf26 = _1418___mcc_h0;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uaimtxsem"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(744)), ((_1424_cf27) => function (_1425_i0) {
            return _1424_cf27;
          })(_1422_cf27)));
        } else {
          let _1426___mcc_h3 = (_source25).cf25;
          let _1427_cf25 = _1426___mcc_h3;
          return _1427_cf25;
        }
      }(_module.D9.create_DC19(_dafny.Seq.UnicodeFromString("sxfde")))).length);
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _1428_v0;
      _1428_v0 = _dafny.Seq.UnicodeFromString("tv");
      let _1429_v1;
      _1429_v1 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4)).length), new BigNumber((_dafny.Seq.of(_1428_v0)).length));
      let _1430_v2;
      let _1431_v3;
      let _out45;
      let _out46;
      let _outcollector19 = _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe(_1429_v1,_1428_v0), !((((_this).f4) ? ((_this).f4) : ((_this).f4))), _1429_v1, (_this).f4, globalState);
      _out45 = _outcollector19[0];
      _out46 = _outcollector19[1];
      _1430_v2 = _out45;
      _1431_v3 = _out46;
      let _1432_v4;
      let _init28 = function (_1433_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      };
      let _nw214 = Array((new BigNumber(9)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw214.length); _i0_28++) {
        _nw214[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _1432_v4 = _nw214;
      let _1434_v5;
      _1434_v5 = new _dafny.CodePoint('i'.codePointAt(0));
      let _index205 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_1432_v4).length));
      (_1432_v4)[_index205] = _1434_v5;
      let _1435_v6;
      _1435_v6 = _dafny.Seq.of(_module.__default.safeDivisionInt(_1431_v3, (_this).f6));
      let _1436_v7;
      _1436_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1430_v2,(_dafny.ZERO).minus(new BigNumber((_module.__default.fm30(_1435_v6, true, globalState)).length)));
      let _1437_v8;
      _1437_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1430_v2,(_this).f4);
      let _1438_v9;
      _1438_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1437_v8,new BigNumber((_dafny.Set.fromElements((_this).f4, _1430_v2)).length));
      let _1439_v11;
      let _nw215 = Array((new BigNumber(22)).toNumber());
      _nw215[(_dafny.ZERO).toNumber()] = new BigNumber(-811);
      _nw215[(_dafny.ONE).toNumber()] = p0;
      _nw215[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-476)), ((_1440_v5) => function (_1441_i1) {
        return _1440_v5;
      })(_1434_v5))).length);
      _nw215[(new BigNumber(3)).toNumber()] = (_this).f6;
      _nw215[(new BigNumber(4)).toNumber()] = (_this).f6;
      _nw215[(new BigNumber(5)).toNumber()] = (_this).f6;
      _nw215[(new BigNumber(6)).toNumber()] = new BigNumber((_1436_v7).length);
      _nw215[(new BigNumber(7)).toNumber()] = (_this).f6;
      _nw215[(new BigNumber(8)).toNumber()] = (_this).f6;
      _nw215[(new BigNumber(9)).toNumber()] = _1431_v3;
      _nw215[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.of(p0, p0, (_dafny.ZERO).minus(_1431_v3), (_this).f6)).length);
      _nw215[(new BigNumber(11)).toNumber()] = (_this).f6;
      _nw215[(new BigNumber(12)).toNumber()] = p0;
      _nw215[(new BigNumber(13)).toNumber()] = (_this).f6;
      _nw215[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ejh")).length);
      _nw215[(new BigNumber(15)).toNumber()] = _module.__default.fm3(_1431_v3, (_this).f4, globalState);
      _nw215[(new BigNumber(16)).toNumber()] = _1431_v3;
      _nw215[(new BigNumber(17)).toNumber()] = _1431_v3;
      _nw215[(new BigNumber(18)).toNumber()] = (((_1438_v9).contains(_dafny.Map.Empty.slice().updateUnsafe(_1430_v2,(_this).f4))) ? ((_1438_v9).get(_dafny.Map.Empty.slice().updateUnsafe(_1430_v2,(_this).f4))) : (new BigNumber((function () {
        let _coll47 = new _dafny.Map();
        for (const _compr_47 of _dafny.IntegerRange(new BigNumber(374), new BigNumber(306))) {
          let _1442_v10 = _compr_47;
          if (((new BigNumber(374)).isLessThanOrEqualTo(_1442_v10)) && ((_1442_v10).isLessThan(new BigNumber(306)))) {
            _coll47.push([(_1442_v10).minus((_this).f6),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1435_v6)).cardinality()),(_this).f4)).length)]);
          }
        }
        return _coll47;
      }()).length)));
      _nw215[(new BigNumber(19)).toNumber()] = p0;
      _nw215[(new BigNumber(20)).toNumber()] = _1431_v3;
      _nw215[(new BigNumber(21)).toNumber()] = (_this).f6;
      _1439_v11 = _nw215;
      let _1443_v12;
      _1443_v12 = _dafny.Seq.of(_1439_v11);
      let _1444_v13;
      _1444_v13 = _dafny.MultiSet.fromElements(_1439_v11, (_1443_v12)[_module.__default.safeIndex(_1431_v3, new BigNumber((_1443_v12).length))]);
      let _1445_v14;
      _1445_v14 = _dafny.MultiSet.fromElements(_1444_v13, (_1444_v13).Difference((_1444_v13).update(_1439_v11, _module.__default.abs(p0))));
      let _index206 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_1432_v4).length));
      let _rhs189 = _1434_v5;
      let _rhs190 = (_this).f6;
      let _rhs191 = (((_1445_v14).contains((_1444_v13).Union(_1444_v13))) ? ((_1445_v14).get((_1444_v13).Union(_1444_v13))) : (((_1430_v2) ? (_1431_v3) : (p0))));
      let _rhs192 = _1435_v6;
      let _lhs119 = _1432_v4;
      let _lhs120 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_1432_v4).length));
      _lhs119[_lhs120] = _rhs189;
      _1431_v3 = _rhs190;
      _1431_v3 = _rhs191;
      _1435_v6 = _rhs192;
      let _1446_v15;
      let _nw216 = new _module.C0();
      _nw216.__ctor(_module.__default.safeDivisionInt(_1431_v3, new BigNumber((_1435_v6).length)));
      _1446_v15 = _nw216;
      let _1447_v16;
      _1447_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1434_v5,_1430_v2);
      _1447_v16 = (_1447_v16).update((_1432_v4)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_1432_v4).length))], (_this).f4);
      let _hi8 = (_dafny.ZERO).minus(((_1446_v15).f12).minus(new BigNumber(281)));
      for (let _1448_i2 = new BigNumber(-659); _1448_i2.isLessThan(_hi8); _1448_i2 = _1448_i2.plus(_dafny.ONE)) {
        let _1449_v17;
        let _nw217 = new _module.C0();
        _nw217.__ctor(_1431_v3);
        _1449_v17 = _nw217;
        let _1450_v18;
        _1450_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jl")).length),_module.__default.fm31(_1448_i2, _1431_v3, new BigNumber((_1428_v0).length), (_this).f4, globalState));
        let _1451_v19;
        _1451_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1431_v3,(_this).f4);
        let _1452_v20;
        _1452_v20 = _dafny.Set.fromElements(new BigNumber((_1429_v1).cardinality()), new BigNumber((_1436_v7).length), (_module.D0.create_DC0((_this).f6)).dtor_cf0, new BigNumber((_1451_v19).length), new BigNumber((_1435_v6).length));
        let _1453_v23;
        let _nw218 = Array((new BigNumber(8)).toNumber());
        _nw218[(_dafny.ZERO).toNumber()] = (((_1450_v18).contains((_1446_v15).f12)) ? ((_1450_v18).get((_1446_v15).f12)) : (_1452_v20));
        _nw218[(_dafny.ONE).toNumber()] = function () {
          let _coll48 = new _dafny.Set();
          for (const _compr_48 of _dafny.IntegerRange(new BigNumber(380), new BigNumber(684))) {
            let _1454_v21 = _compr_48;
            if (((new BigNumber(380)).isLessThanOrEqualTo(_1454_v21)) && ((_1454_v21).isLessThan(new BigNumber(684)))) {
              _coll48.add((_1454_v21).minus((_1446_v15).f12));
            }
          }
          return _coll48;
        }();
        _nw218[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements((_this).f6, (_1446_v15).f12, (_1449_v17).f12, new BigNumber((_1429_v1).cardinality()));
        _nw218[(new BigNumber(3)).toNumber()] = (_module.__default.fm32(new BigNumber(558), _1430_v2, globalState));
        _nw218[(new BigNumber(4)).toNumber()] = _1452_v20;
        _nw218[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(_1448_i2);
        _nw218[(new BigNumber(6)).toNumber()] = _1452_v20;
        _nw218[(new BigNumber(7)).toNumber()] = function () {
          let _coll49 = new _dafny.Set();
          for (const _compr_49 of _dafny.IntegerRange(new BigNumber(788), new BigNumber(145))) {
            let _1455_v22 = _compr_49;
            if (((new BigNumber(788)).isLessThanOrEqualTo(_1455_v22)) && ((_1455_v22).isLessThan(new BigNumber(145)))) {
              _coll49.add((_1455_v22).minus((_1449_v17).f12));
            }
          }
          return _coll49;
        }();
        _1453_v23 = _nw218;
        let _1456_v24;
        _1456_v24 = _dafny.Set.fromElements((_1449_v17).fm17((_this).f4, globalState), !(!((_this).f4)), _1430_v2, (_this).f4, !((_this).f4));
        let _1457_v25;
        _1457_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1456_v24,_1430_v2);
        let _1458_v26;
        _1458_v26 = _dafny.Map.Empty.slice().updateUnsafe(((_1430_v2) ? (_1453_v23) : (_1453_v23)),new BigNumber((_1457_v25).length));
        _1458_v26 = (_1458_v26).update(_1453_v23, _module.__default.safeModuloInt(p0, (_1446_v15).f12));
        let _index207 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1439_v11).length));
        (_1439_v11)[_index207] = new BigNumber((_1436_v7).length);
        let _1459_v27;
        _1459_v27 = _dafny.Seq.of((_this).f4, _1430_v2, _1430_v2, _1430_v2);
        let _index208 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1439_v11).length));
        (_1439_v11)[_index208] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1459_v27, _module.__default.safeIndex(new BigNumber((_1459_v27).length), new BigNumber((_1459_v27).length)), (_this).f4), _1459_v27)).length);
        let _1460_v28;
        _1460_v28 = _dafny.MultiSet.fromElements((_1446_v15).fm17((_this).f4, globalState), (_1446_v15).fm17(true, globalState), _1430_v2);
        _1460_v28 = _1460_v28;
      }
      let _index209 = _module.__default.safeIndex(new BigNumber(290), new BigNumber((_1432_v4).length));
      (_1432_v4)[_index209] = (_1432_v4)[_module.__default.safeIndex(new BigNumber(290), new BigNumber((_1432_v4).length))];
      r0 = _1436_v7;
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = [];
      let _1461_v0;
      _1461_v0 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus((_this).f6)).multipliedBy((_this).f6),(_this).fm7(globalState));
      let _1462_v1;
      let _nw219 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _1462_v1 = _nw219;
      let _1463_v2;
      _1463_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1462_v1,_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_dafny.ZERO).minus((_this).f6)));
      _1461_v0 = (((_1463_v2).contains(_1462_v1)) ? ((_1463_v2).get(_1462_v1)) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6)));
      let _1464_v3;
      _1464_v3 = _dafny.Seq.UnicodeFromString("tjatmipct");
      let _1465_v4;
      _1465_v4 = _dafny.Seq.of(new BigNumber((_1464_v3).length), (_this).f6);
      _1465_v4 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(355)), function (_1466_i0) {
        return (_this).f6;
      });
      let _hi9 = new BigNumber((_1464_v3).length);
      for (let _1467_i1 = (_this).f6; _1467_i1.isLessThan(_hi9); _1467_i1 = _1467_i1.plus(_dafny.ONE)) {
        let _1468_v5;
        _1468_v5 = _dafny.Seq.of(p0);
        let _1469_v6;
        _1469_v6 = _module.D4.create_DC13(_module.__default.safeModuloInt(new BigNumber((_1461_v0).length), _1467_i1), (_this).f4, new BigNumber(-262), _1468_v5);
        let _source26 = _1469_v6;
        if (_source26.is_DC13) {
          let _1470___mcc_h0 = (_source26).cf14;
          let _1471___mcc_h1 = (_source26).cf15;
          let _1472___mcc_h2 = (_source26).cf16;
          let _1473___mcc_h3 = (_source26).cf17;
          let _1474_cf17 = _1473___mcc_h3;
          let _1475_cf16 = _1472___mcc_h2;
          let _1476_cf15 = _1471___mcc_h1;
          let _1477_cf14 = _1470___mcc_h0;
          let _1478_v7;
          _1478_v7 = new _dafny.CodePoint('y'.codePointAt(0));
          _1475_cf16 = ((!_dafny.Seq.contains(_1464_v3, _1478_v7)) ? (_1477_cf14) : (_1477_cf14));
          let _1479_v8;
          let _nw220 = new _module.C0();
          _nw220.__ctor((_this).f6);
          _1479_v8 = _nw220;
          let _1480_v9;
          _1480_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1479_v8,_1467_i1);
          let _1481_v10;
          _1481_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1476_cf15,_1464_v3);
          let _1482_v11;
          let _nw221 = Array((new BigNumber(22)).toNumber());
          _nw221[(_dafny.ZERO).toNumber()] = (_this).f4;
          _nw221[(_dafny.ONE).toNumber()] = (_this).f4;
          _nw221[(new BigNumber(2)).toNumber()] = p0;
          _nw221[(new BigNumber(3)).toNumber()] = p0;
          _nw221[(new BigNumber(4)).toNumber()] = (_this).f4;
          _nw221[(new BigNumber(5)).toNumber()] = (_this).f4;
          _nw221[(new BigNumber(6)).toNumber()] = _1476_cf15;
          _nw221[(new BigNumber(7)).toNumber()] = p0;
          _nw221[(new BigNumber(8)).toNumber()] = p0;
          _nw221[(new BigNumber(9)).toNumber()] = (_this).f4;
          _nw221[(new BigNumber(10)).toNumber()] = !(!(_1476_cf15));
          _nw221[(new BigNumber(11)).toNumber()] = (_this).f4;
          _nw221[(new BigNumber(12)).toNumber()] = (_this).f4;
          _nw221[(new BigNumber(13)).toNumber()] = _1476_cf15;
          _nw221[(new BigNumber(14)).toNumber()] = !(p0);
          _nw221[(new BigNumber(15)).toNumber()] = p0;
          _nw221[(new BigNumber(16)).toNumber()] = (_this).f4;
          _nw221[(new BigNumber(17)).toNumber()] = true;
          _nw221[(new BigNumber(18)).toNumber()] = _1476_cf15;
          _nw221[(new BigNumber(19)).toNumber()] = (_this).fm8(_1475_cf16, _1476_cf15, (((_1481_v10).contains(p0)) ? ((_1481_v10).get(p0)) : (_1464_v3)), _1465_v4, globalState);
          _nw221[(new BigNumber(20)).toNumber()] = _1476_cf15;
          _nw221[(new BigNumber(21)).toNumber()] = (_this).f4;
          _1482_v11 = _nw221;
          _1461_v0 = (_1461_v0).update((new BigNumber((_1480_v9).length)).minus(new BigNumber(450)), (_module.D10.create_DC22(_1464_v3, p0, false, _1482_v11, (_1479_v8).f12)).dtor_cf34);
          let _index210 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1462_v1).length));
          (_1462_v1)[_index210] = new BigNumber(784);
          let _index211 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_1462_v1).length));
          (_1462_v1)[_index211] = (_1467_i1).multipliedBy(new BigNumber(925));
          let _1483_v12;
          _1483_v12 = _module.D0.create_DC2(_1475_cf16);
          let _1484_v13;
          _1484_v13 = _dafny.MultiSet.fromElements(_1478_v7);
          let _1485_v14;
          let _nw222 = new _module.C3();
          _nw222.__ctor(_this.f5, new BigNumber((_1484_v13).cardinality()), _module.__default.fm0(new _dafny.CodePoint('f'.codePointAt(0)), globalState));
          _1485_v14 = _nw222;
          let _1486_v15;
          _1486_v15 = _dafny.Seq.of(_1485_v14, _1485_v14, _1485_v14, _1485_v14, _1485_v14);
          let _pat_let_tv24 = _1468_v5;
          let _pat_let_tv25 = _1475_cf16;
          let _pat_let_tv26 = _1475_cf16;
          let _pat_let_tv27 = _1462_v1;
          let _pat_let_tv28 = _1462_v1;
          let _1487_v16;
          let _nw223 = Array((new BigNumber(21)).toNumber());
          _nw223[(_dafny.ZERO).toNumber()] = _1483_v12;
          _nw223[(_dafny.ONE).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(2)).toNumber()] = function (_pat_let30_0) {
            return function (_1488_dt__update__tmp_h0) {
              return function (_pat_let31_0) {
                return function (_1489_dt__update_hcf1_h0) {
                  return _module.D0.create_DC2(_1489_dt__update_hcf1_h0);
                }(_pat_let31_0);
              }(new BigNumber((_pat_let_tv24).length));
            }(_pat_let30_0);
          }(_module.D0.create_DC2(_1475_cf16));
          _nw223[(new BigNumber(3)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(4)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(5)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(6)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(7)).toNumber()] = function (_pat_let32_0) {
            return function (_1490_dt__update__tmp_h1) {
              return function (_pat_let33_0) {
                return function (_1491_dt__update_hcf1_h1) {
                  return _module.D0.create_DC2(_1491_dt__update_hcf1_h1);
                }(_pat_let33_0);
              }(_pat_let_tv25);
            }(_pat_let32_0);
          }(_module.D0.create_DC2(new BigNumber(291)));
          _nw223[(new BigNumber(8)).toNumber()] = _module.D0.create_DC2(new BigNumber(431));
          _nw223[(new BigNumber(9)).toNumber()] = function (_pat_let34_0) {
            return function (_1492_dt__update__tmp_h2) {
              return function (_pat_let35_0) {
                return function (_1493_dt__update_hcf1_h2) {
                  return _module.D0.create_DC2(_1493_dt__update_hcf1_h2);
                }(_pat_let35_0);
              }(_pat_let_tv26);
            }(_pat_let34_0);
          }(_1483_v12);
          _nw223[(new BigNumber(10)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(11)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(12)).toNumber()] = function (_pat_let36_0) {
            return function (_1494_dt__update__tmp_h3) {
              return function (_pat_let37_0) {
                return function (_1495_dt__update_hcf1_h3) {
                  return _module.D0.create_DC2(_1495_dt__update_hcf1_h3);
                }(_pat_let37_0);
              }((_pat_let_tv28)[_module.__default.safeIndex(new BigNumber(991), new BigNumber((_pat_let_tv27).length))]);
            }(_pat_let36_0);
          }(_1483_v12);
          _nw223[(new BigNumber(13)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(14)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(15)).toNumber()] = _module.D0.create_DC2(new BigNumber((_1486_v15).length));
          _nw223[(new BigNumber(16)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(17)).toNumber()] = _module.D0.create_DC2(_1467_i1);
          _nw223[(new BigNumber(18)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(19)).toNumber()] = _1483_v12;
          _nw223[(new BigNumber(20)).toNumber()] = _module.D0.create_DC2((_1479_v8).f12);
          _1487_v16 = _nw223;
          let _init29 = ((_1496_v12) => function (_1497_i2) {
            return _1496_v12;
          })(_1483_v12);
          let _nw224 = Array((new BigNumber(22)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw224.length); _i0_29++) {
            _nw224[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1487_v16 = _nw224;
        } else {
          let _1498___mcc_h4 = (_source26).cf13;
          let _1499_cf13 = _1498___mcc_h4;
          let _1500_v17;
          _1500_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1462_v1);
          _1500_v17 = (_1500_v17).update(_1467_i1, _1462_v1);
          let _1501_v18;
          _1501_v18 = new BigNumber(277);
          let _1502_v20;
          _1502_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1467_i1),_1499_cf13);
          let _1503_v21;
          _1503_v21 = _dafny.Set.fromElements(_1502_v20);
          _1501_v18 = _module.__default.safeDivisionInt(new BigNumber((function () {
            let _coll50 = new _dafny.Map();
            for (const _compr_50 of _dafny.IntegerRange(new BigNumber(447), new BigNumber(-282))) {
              let _1504_v19 = _compr_50;
              if (((new BigNumber(447)).isLessThanOrEqualTo(_1504_v19)) && ((_1504_v19).isLessThan(new BigNumber(-282)))) {
                _coll50.push([_module.__default.safeDivisionInt(_1504_v19, new BigNumber(592)),_1467_i1]);
              }
            }
            return _coll50;
          }()).length), _module.__default.safeDivisionInt(new BigNumber((_1503_v21).length), (_this).f6));
          let _1505_v22;
          let _nw225 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
          _1505_v22 = _nw225;
          let _index212 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1505_v22).length));
          (_1505_v22)[_index212] = _1465_v4;
          let _index213 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1505_v22).length));
          (_1505_v22)[_index213] = _dafny.Seq.of(_1501_v18);
          let _1506_v23;
          let _nw226 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1506_v23 = _nw226;
          let _index214 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1506_v23).length));
          (_1506_v23)[_index214] = _1499_cf13;
          let _index215 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1506_v23).length));
          (_1506_v23)[_index215] = _1499_cf13;
        }
        let _1507_v24;
        _1507_v24 = true;
        let _1508_v25;
        _1508_v25 = _dafny.MultiSet.fromElements(_1467_i1, (_this).f6);
        let _1509_v26;
        _1509_v26 = _dafny.Set.fromElements(p0, (_this).f4);
        let _1510_v27;
        _1510_v27 = _dafny.Seq.of(_1465_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(980)), ((_1511_i1) => function (_1512_i3) {
          return _1511_i1;
        })(_1467_i1)));
        _1507_v24 = (_this).fm8((((_1508_v25).contains(new BigNumber((_1509_v26).length))) ? ((_1508_v25).get(new BigNumber((_1509_v26).length))) : (new BigNumber(622))), _dafny.areEqual(_1510_v27, _dafny.Seq.of(_1465_v4)), _1464_v3, _dafny.Seq.Concat(_1465_v4, _1465_v4), globalState);
        let _1513_v28;
        _1513_v28 = _module.D11.create_DC24((_this).f4, (_this).f4, new BigNumber((_1464_v3).length));
        let _1514_v29;
        _1514_v29 = _module.D20.create_DC41(((_1507_v24) ? (_1467_i1) : ((_1513_v28).dtor_cf38)));
        _1514_v29 = _module.D20.create_DC41(new BigNumber(130));
        let _1515_v30;
        _1515_v30 = new BigNumber(174);
        let _1516_v31;
        _1516_v31 = _dafny.Set.fromElements(_1467_i1);
        let _1517_v32;
        _1517_v32 = _1516_v31;
        let _1518_v33;
        _1518_v33 = _dafny.Set.fromElements(_1517_v32);
        _1515_v30 = ((_dafny.ZERO).minus(((_dafny.ZERO).minus(_1515_v30)).minus((_this).f6))).plus((((_this).f4) ? (_1515_v30) : (new BigNumber((_1518_v33).length))));
      }
      let _1519_v34;
      _1519_v34 = false;
      let _index216 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length));
      (_1462_v1)[_index216] = ((_this).f6).multipliedBy((_this).f6);
      let _1520_v35;
      _1520_v35 = new BigNumber(978);
      let _1521_v36;
      _1521_v36 = _dafny.Map.Empty.slice().updateUnsafe((_1465_v4)[_module.__default.safeIndex(_1520_v35, new BigNumber((_1465_v4).length))],_1519_v34);
      let _1522_v37;
      _1522_v37 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f4),_1521_v36);
      let _1523_v38;
      _1523_v38 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_this).f6));
      let _1524_v39;
      _1524_v39 = _dafny.MultiSet.fromElements(p0);
      let _index217 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length));
      let _rhs193 = !(_1522_v37).equals(_module.__default.fm60(new BigNumber((_1523_v38).length), (_this).f6, _1520_v35, globalState));
      let _rhs194 = _1520_v35;
      let _rhs195 = _module.__default.safeDivisionInt(new BigNumber((_1524_v39).cardinality()), _1520_v35);
      let _rhs196 = !(!(_1519_v34));
      let _rhs197 = (_this).f4;
      let _lhs121 = _1462_v1;
      let _lhs122 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length));
      _1519_v34 = _rhs193;
      _lhs121[_lhs122] = _rhs194;
      _1520_v35 = _rhs195;
      _1519_v34 = _rhs196;
      _1519_v34 = _rhs197;
      let _1525_v40;
      let _init30 = ((_1526_p0) => function (_1527_i4) {
        return _module.D12.create_DC27(_dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(_1526_p0)));
      })(p0);
      let _nw227 = Array((new BigNumber(28)).toNumber());
      for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw227.length); _i0_30++) {
        _nw227[_i0_30] = _init30(new BigNumber(_i0_30));
      }
      _1525_v40 = _nw227;
      let _1528_v41;
      _1528_v41 = _dafny.Set.fromElements((_this).f4, (_this).f4);
      let _1529_v42;
      _1529_v42 = _module.D12.create_DC27(_module.__default.fm61(_1528_v41, globalState));
      let _index218 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_1525_v40).length));
      (_1525_v40)[_index218] = _1529_v42;
      let _index219 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_1525_v40).length));
      let _rhs198 = _1520_v35;
      let _rhs199 = _1529_v42;
      let _lhs123 = _1525_v40;
      let _lhs124 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_1525_v40).length));
      _1520_v35 = _rhs198;
      _lhs123[_lhs124] = _rhs199;
      let _1530_v43;
      let _nw228 = new _module.C1();
      _nw228.__ctor(_1520_v35, _module.__default.safeModuloInt((_1462_v1)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length))], _1520_v35), _this.f5, (_this).f6, (_this).fm8((_this).f6, (_this).f4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), function (_1531_i5) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-532)), ((_1532_v35) => function (_1533_i6) {
        return _1532_v35;
      })(_1520_v35)), globalState));
      _1530_v43 = _nw228;
      let _1534_v44;
      _1534_v44 = new _dafny.CodePoint('f'.codePointAt(0));
      let _1535_v45;
      _1535_v45 = _module.D9.create_DC20((_this).f4, _1534_v44, _1530_v43.f16);
      let _1536_v46;
      _1536_v46 = _dafny.Seq.of(_1535_v45);
      let _index220 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length));
      let _rhs200 = ((((((_1461_v0).contains((_1462_v1)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length))])) ? ((_1461_v0).get((_1462_v1)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length))])) : ((_1462_v1)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length))]))).isLessThanOrEqualTo((_1462_v1)[_module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length))])) ? (_1530_v43) : (_1530_v43));
      let _rhs201 = p0;
      let _rhs202 = p0;
      let _rhs203 = (_dafny.ZERO).minus(function (_source27) {
        let _1537___mcc_h5 = _source27;
        let _1538_cf57 = _1537___mcc_h5;
        return new BigNumber(911);
      }(_1536_v46));
      let _lhs125 = _1462_v1;
      let _lhs126 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1462_v1).length));
      _1530_v43 = _rhs200;
      _1519_v34 = _rhs201;
      _1519_v34 = _rhs202;
      _lhs125[_lhs126] = _rhs203;
      r0 = _1534_v44;
      let _1539_v47;
      let _nw229 = Array((new BigNumber(2)).toNumber()).fill(_module.D1.Default());
      _1539_v47 = _nw229;
      r1 = _1539_v47;
      return [r0, r1];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f5 = [];
      this._f6 = _dafny.ZERO;
      this._f4 = false;
      this.f14 = _dafny.ZERO;
      this._f13 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    set f5(value) {
      let _this = this;
      _this._f5 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f13, f14, f5, f6, f4) {
      let _this = this;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f4 = f4;
      return;
    }
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((_this).f13) {
        return (_dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4)).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13));
      } else {
        return !(_dafny.MultiSet.fromElements(_this.f14)).contains((_dafny.ZERO).minus(_this.f14));
      }
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(30);
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements((_this).f6, new BigNumber(309), (_this).f6)).Difference(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_this.f14)));
    };
    fm7(globalState) {
      let _this = this;
      let _source28 = _module.D3.create_DC9((_this).f13, (_this).f6, (_this).f4);
      if (_source28.is_DC9) {
        let _1540___mcc_h0 = (_source28).cf8;
        let _1541___mcc_h1 = (_source28).cf9;
        let _1542___mcc_h2 = (_source28).cf10;
        let _1543_cf10 = _1542___mcc_h2;
        let _1544_cf9 = _1541___mcc_h1;
        let _1545_cf8 = _1540___mcc_h0;
        return (_this).f6;
      } else if (_source28.is_DC10) {
        let _1546___mcc_h3 = (_source28).cf11;
        let _1547_cf11 = _1546___mcc_h3;
        return (_this.f14).plus(_this.f14);
      } else if (_source28.is_DC8) {
        let _1548___mcc_h4 = (_source28).cf7;
        let _1549_cf7 = _1548___mcc_h4;
        return (new BigNumber(64)).minus(new BigNumber(((_module.D9.create_DC19(_dafny.Seq.UnicodeFromString("fafdlayl"))).dtor_cf25).length));
      } else {
        let _1550___mcc_h5 = (_source28).cf12;
        let _1551_cf12 = _1550___mcc_h5;
        return (_this).f6;
      }
    };
    fm26(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(_this.f14);
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _1552_v0;
      _1552_v0 = _dafny.MultiSet.fromElements((_this).f13, (_this).f4, (_this).f4, (_this).f13, (_this).f13);
      let _1553_v1;
      _1553_v1 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC15(_module.__default.fm27((_this).f13, new BigNumber(35), _1552_v0, (_this).f13, globalState)),new BigNumber(((_1552_v0).Intersect(_1552_v0)).cardinality()));
      _1553_v1 = _1553_v1;
      (_this).f14 = (_this).f6;
      let _1554_v2;
      _1554_v2 = _dafny.Set.fromElements((_this).f4);
      let _1555_v3;
      _1555_v3 = _dafny.Seq.of(_1554_v2);
      let _1556_v4;
      _1556_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber((_1555_v3).length));
      let _1557_v5;
      _1557_v5 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1558_v6;
      _1558_v6 = _module.D9.create_DC20((_this).f13, _1557_v5, (_this).f6);
      let _1559_v7;
      let _nw230 = new _module.C0();
      _nw230.__ctor((((_1556_v4).contains((_1558_v6).dtor_cf26)) ? ((_1556_v4).get((_1558_v6).dtor_cf26)) : (p0)));
      _1559_v7 = _nw230;
      let _1560_v8;
      _1560_v8 = _dafny.Seq.of((_this).f13);
      let _1561_v9;
      _1561_v9 = _dafny.Seq.of(_dafny.Seq.of((_this).f4, (_this).f4, (_this).f13, (_this).f4, (_this).f4));
      let _1562_v11;
      _1562_v11 = _dafny.Seq.UnicodeFromString("trc");
      let _1563_v12;
      let _init31 = function (_1564_i0) {
        return _module.__default.safeModuloInt(_1564_i0, (_this).f6);
      };
      let _nw231 = Array((_dafny.ONE).toNumber());
      for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw231.length); _i0_31++) {
        _nw231[_i0_31] = _init31(new BigNumber(_i0_31));
      }
      _1563_v12 = _nw231;
      let _1565_v13;
      _1565_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1563_v12);
      let _1566_v14;
      let _nw232 = Array((new BigNumber(20)).toNumber());
      _nw232[(_dafny.ZERO).toNumber()] = _this.f14;
      _nw232[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_this.f14);
      _nw232[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber(59), (_1559_v7).f12, (_1559_v7).f12)).length);
      _nw232[(new BigNumber(3)).toNumber()] = (_this.f14).minus(_this.f14);
      _nw232[(new BigNumber(4)).toNumber()] = (_1559_v7).f12;
      _nw232[(new BigNumber(5)).toNumber()] = (_this).f6;
      _nw232[(new BigNumber(6)).toNumber()] = p0;
      _nw232[(new BigNumber(7)).toNumber()] = (_this).f6;
      _nw232[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1560_v8)).cardinality()),(_1559_v7).f12)).length);
      _nw232[(new BigNumber(9)).toNumber()] = (_this).f6;
      _nw232[(new BigNumber(10)).toNumber()] = _this.f14;
      _nw232[(new BigNumber(11)).toNumber()] = (_1559_v7).f12;
      _nw232[(new BigNumber(12)).toNumber()] = new BigNumber((_1561_v9).length);
      _nw232[(new BigNumber(13)).toNumber()] = new BigNumber((function () {
        let _coll51 = new _dafny.Set();
        for (const _compr_51 of _dafny.IntegerRange(new BigNumber(253), new BigNumber(-291))) {
          let _1567_v10 = _compr_51;
          if (((new BigNumber(253)).isLessThanOrEqualTo(_1567_v10)) && ((_1567_v10).isLessThan(new BigNumber(-291)))) {
            _coll51.add((_1567_v10).multipliedBy((_1559_v7).f12));
          }
        }
        return _coll51;
      }()).length);
      _nw232[(new BigNumber(14)).toNumber()] = ((_this).f6).minus(new BigNumber((_1562_v11).length));
      _nw232[(new BigNumber(15)).toNumber()] = new BigNumber((_1565_v13).length);
      _nw232[(new BigNumber(16)).toNumber()] = (_this).f6;
      _nw232[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(438), _this.f14);
      _nw232[(new BigNumber(18)).toNumber()] = (_this).f6;
      _nw232[(new BigNumber(19)).toNumber()] = (_1559_v7).f12;
      _1566_v14 = _nw232;
      let _index221 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_1563_v12).length));
      (_1563_v12)[_index221] = _this.f14;
      let _1568_v15;
      _1568_v15 = _module.D10.create_DC21(_module.__default.fm28(globalState));
      let _1569_v16;
      _1569_v16 = _dafny.Seq.of(_1555_v3, _1555_v3);
      let _index222 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_1563_v12).length));
      let _rhs204 = (((_this).f13) ? ((_1559_v7).f12) : ((_1559_v7).f12));
      let _rhs205 = (_1568_v15).dtor_cf29;
      let _rhs206 = ((!(_this.f14).isEqualTo(new BigNumber(521))) ? (_dafny.Seq.of(_1554_v2)) : ((_1569_v16)[_module.__default.safeIndex((_1559_v7).f12, new BigNumber((_1569_v16).length))]));
      let _lhs127 = _1563_v12;
      let _lhs128 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_1563_v12).length));
      _lhs127[_lhs128] = _rhs204;
      _1556_v4 = _rhs205;
      _1555_v3 = _rhs206;
      let _1570_v17;
      _1570_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1557_v5,_1557_v5);
      _1570_v17 = (_1570_v17).update(new _dafny.CodePoint('w'.codePointAt(0)), _1557_v5);
      let _1571_v18;
      _1571_v18 = _dafny.MultiSet.fromElements(new BigNumber((_1562_v11).length), new BigNumber((_1554_v2).length));
      let _1572_v19;
      _1572_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v18,_dafny.Seq.UnicodeFromString("ydutpac"));
      let _1573_v20;
      let _1574_v21;
      let _out47;
      let _out48;
      let _outcollector20 = _module.__default.m0((_1572_v19).Merge(_1572_v19), false, _1571_v18, (_this).f4, globalState);
      _out47 = _outcollector20[0];
      _out48 = _outcollector20[1];
      _1573_v20 = _out47;
      _1574_v21 = _out48;
      r0 = _1556_v4;
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = [];
      let _1575_v1;
      _1575_v1 = _dafny.Seq.UnicodeFromString("rptj");
      let _1576_v2;
      _1576_v2 = _dafny.MultiSet.fromElements(new BigNumber((_1575_v1).length));
      let _1577_v3;
      _1577_v3 = _dafny.Seq.of(_this.f14, (_this).f6, _this.f14, _this.f14, (_this).f6);
      let _1578_v4;
      _1578_v4 = _dafny.Seq.of(_1576_v2, _dafny.MultiSet.FromArray(_1577_v3));
      let _1579_v5;
      let _1580_v6;
      let _out49;
      let _out50;
      let _outcollector21 = _module.__default.m0(function () {
        let _coll52 = new _dafny.Map();
        for (const _compr_52 of (_1578_v4).Elements) {
          let _1581_v0 = _compr_52;
          if (_dafny.Seq.contains(_1578_v4, _1581_v0)) {
            _coll52.push([_1581_v0,_1575_v1]);
          }
        }
        return _coll52;
      }(), (_this).f4, _1576_v2, !((_this).f6).isEqualTo((_this).f6), globalState);
      _out49 = _outcollector21[0];
      _out50 = _outcollector21[1];
      _1579_v5 = _out49;
      _1580_v6 = _out50;
      _1575_v1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), function (_1582_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      });
      let _1583_v7;
      let _nw233 = new _module.C0();
      _nw233.__ctor(new BigNumber((_dafny.Seq.of((_this).f4, !(_1579_v5), !((_this).f13))).length));
      _1583_v7 = _nw233;
      let _1584_v8;
      _1584_v8 = new _dafny.CodePoint('p'.codePointAt(0));
      let _1585_v9;
      _1585_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_module.__default.fm0(_1584_v8, globalState));
      let _1586_v10;
      _1586_v10 = _dafny.MultiSet.fromElements(_1584_v8);
      _1585_v9 = (_1585_v9).update((_dafny.ZERO).minus(new BigNumber((_1586_v10).cardinality())), (_1580_v6).isEqualTo((_dafny.ZERO).minus(_this.f14)));
      let _hi10 = new BigNumber(496);
      for (let _1587_i1 = ((_this).f6).minus((_this).f6); _1587_i1.isLessThan(_hi10); _1587_i1 = _1587_i1.plus(_dafny.ONE)) {
        if ((false) && (!(p0))) {
          let _1588_v11;
          let _init32 = function (_1589_i2) {
            return (_this).f4;
          };
          let _nw234 = Array((new BigNumber(10)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw234.length); _i0_32++) {
            _nw234[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1588_v11 = _nw234;
          let _1590_v12;
          _1590_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1588_v11,p0);
          let _1591_v13;
          let _nw235 = new _module.C0();
          _nw235.__ctor((_this).fm26((_1583_v7).f12, (_this).fm8((_this).f6, true, _1575_v1, _1577_v3, globalState), new BigNumber((_1590_v12).length), globalState));
          _1591_v13 = _nw235;
          let _1592_v14;
          _1592_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f6);
          let _1593_v15;
          _1593_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1592_v14).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), ((_1594_v13) => function (_1595_i3) {
            return (_1594_v13).f12;
          })(_1591_v13))).length));
          _1593_v15 = (_1593_v15).update(new BigNumber((function () {
            let _coll53 = new _dafny.Map();
            for (const _compr_53 of _dafny.IntegerRange(new BigNumber(444), new BigNumber(515))) {
              let _1596_v16 = _compr_53;
              if (((new BigNumber(444)).isLessThanOrEqualTo(_1596_v16)) && ((_1596_v16).isLessThan(new BigNumber(515)))) {
                _coll53.push([(_1596_v16).multipliedBy((_1583_v7).f12),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(610), _this.f14)).cardinality())]);
              }
            }
            return _coll53;
          }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_1597_v3) => function (_1598_i4) {
            return _1597_v3;
          })(_1577_v3)), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,new BigNumber(541))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_1599_v3) => function (_1600_i4) {
            return _1599_v3;
          })(_1577_v3))).length)), _1577_v3)).length))).length));
          let _1601_v17;
          let _init33 = ((_1602_v7) => function (_1603_i5) {
            return _module.__default.safeDivisionInt(_1603_i5, (_1602_v7).f12);
          })(_1583_v7);
          let _nw236 = Array((new BigNumber(6)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw236.length); _i0_33++) {
            _nw236[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1601_v17 = _nw236;
          let _index223 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_1601_v17).length));
          (_1601_v17)[_index223] = (new BigNumber(40)).multipliedBy(_this.f14);
          let _1604_v18;
          _1604_v18 = _dafny.MultiSet.fromElements((_this).f4, p0, (_this).f4);
          let _index224 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_1601_v17).length));
          (_1601_v17)[_index224] = _module.__default.safeDivisionInt((_this.f14).minus(new BigNumber((_1604_v18).cardinality())), (_this).f6);
          _1579_v5 = (_this).f4;
          _1580_v6 = (_1583_v7).f12;
        } else {
          let _1605_v19;
          _1605_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this.f14).minus((_1583_v7).f12),(new BigNumber(-598)).minus((_1583_v7).f12));
          _1605_v19 = _module.__default.fm29(globalState);
          _1580_v6 = (_1583_v7).f12;
          r0 = _1584_v8;
          _1584_v8 = _1584_v8;
          let _rhs207 = _this.f14;
          let _rhs208 = ((_1583_v7).f12).multipliedBy(new BigNumber((_1575_v1).length));
          let _rhs209 = (_this).f6;
          let _rhs210 = _1584_v8;
          let _rhs211 = (_1583_v7).f12;
          let _lhs129 = _this;
          let _lhs130 = _this;
          let _lhs131 = _this;
          _1580_v6 = _rhs207;
          _lhs129.f14 = _rhs208;
          _lhs130.f14 = _rhs209;
          r0 = _rhs210;
          _lhs131.f14 = _rhs211;
        }
        if (_1579_v5) {
          (_this).f14 = _1587_i1;
          (_this).f14 = (_1587_i1).multipliedBy(_1580_v6);
          _1580_v6 = (_1583_v7).f12;
          _1579_v5 = !((_this).f13);
          let _1606_v20;
          let _nw237 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1606_v20 = _nw237;
          let _index225 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_1606_v20).length));
          (_1606_v20)[_index225] = _1575_v1;
          let _1607_v21;
          _1607_v21 = _dafny.Seq.of(_1575_v1);
          let _index226 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_1606_v20).length));
          (_1606_v20)[_index226] = _dafny.Seq.Concat((_1607_v21)[_module.__default.safeIndex((_1583_v7).f12, new BigNumber((_1607_v21).length))], _dafny.Seq.Concat((_1607_v21)[_module.__default.safeIndex(new BigNumber((_1577_v3).length), new BigNumber((_1607_v21).length))], _1575_v1));
        } else {
          let _1608_v22;
          let _nw238 = new _module.C2();
          _nw238.__ctor(_1575_v1, p0);
          _1608_v22 = _nw238;
          let _1609_v23;
          _1609_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1579_v5,_1608_v22);
          _1609_v23 = (_1609_v23).update(!((_this).f13) || ((_this).f13), _1608_v22);
          _1579_v5 = !(p0);
          let _1610_v24;
          _1610_v24 = _module.D4.create_DC12(_1584_v8);
          let _1611_v25;
          let _init34 = function (_1612_i6) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          };
          let _nw239 = Array((new BigNumber(9)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw239.length); _i0_34++) {
            _nw239[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1611_v25 = _nw239;
          let _index227 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_1611_v25).length));
          (_1611_v25)[_index227] = _1584_v8;
          let _1613_v26;
          _1613_v26 = _dafny.Seq.of(!(_1579_v5));
          let _1614_v27;
          let _nw240 = Array((new BigNumber(11)).toNumber());
          _nw240[(_dafny.ZERO).toNumber()] = !((_this).f4);
          _nw240[(_dafny.ONE).toNumber()] = (_this).f4;
          _nw240[(new BigNumber(2)).toNumber()] = (_this).f4;
          _nw240[(new BigNumber(3)).toNumber()] = (_this).f13;
          _nw240[(new BigNumber(4)).toNumber()] = false;
          _nw240[(new BigNumber(5)).toNumber()] = (_this).f13;
          _nw240[(new BigNumber(6)).toNumber()] = false;
          _nw240[(new BigNumber(7)).toNumber()] = !(p0) || (true);
          _nw240[(new BigNumber(8)).toNumber()] = !((_this).f4);
          _nw240[(new BigNumber(9)).toNumber()] = _dafny.areEqual(_1613_v26, _1613_v26);
          _nw240[(new BigNumber(10)).toNumber()] = _module.__default.fm0(_1584_v8, globalState);
          _1614_v27 = _nw240;
          let _index228 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_1614_v27).length));
          (_1614_v27)[_index228] = (_this).fm8(_this.f14, p0, _1575_v1, _dafny.Seq.of(_this.f14), globalState);
          let _1615_v28;
          _1615_v28 = _module.D10.create_DC22(_1575_v1, p0, (_this).f13, _1614_v27, _1587_i1);
          let _index229 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_1611_v25).length));
          let _index230 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_1614_v27).length));
          let _rhs212 = _module.__default.fm62(globalState);
          let _rhs213 = (((_1615_v28).dtor_cf31) ? (_1584_v8) : (new _dafny.CodePoint('d'.codePointAt(0))));
          let _rhs214 = (p0) === ((_1587_i1).isEqualTo((_1583_v7).f12));
          let _rhs215 = _1614_v27;
          let _lhs132 = _1611_v25;
          let _lhs133 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_1611_v25).length));
          let _lhs134 = _1614_v27;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_1614_v27).length));
          _1610_v24 = _rhs212;
          _lhs132[_lhs133] = _rhs213;
          _lhs134[_lhs135] = _rhs214;
          _1614_v27 = _rhs215;
          let _index231 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_1614_v27).length));
          let _rhs216 = p0;
          let _rhs217 = (((_1579_v5) ? (new BigNumber(909)) : (new BigNumber((_1575_v1).length)))).isLessThanOrEqualTo(_1580_v6);
          let _lhs136 = _1614_v27;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_1614_v27).length));
          _lhs136[_lhs137] = _rhs216;
          _1579_v5 = _rhs217;
          _1579_v5 = p0;
        }
        let _1616_v29;
        let _nw241 = Array((new BigNumber(4)).toNumber());
        _nw241[(_dafny.ZERO).toNumber()] = (_this).f13;
        _nw241[(_dafny.ONE).toNumber()] = _1579_v5;
        _nw241[(new BigNumber(2)).toNumber()] = (new BigNumber(49)).isLessThan(_1580_v6);
        _nw241[(new BigNumber(3)).toNumber()] = (_this).f4;
        _1616_v29 = _nw241;
        _1616_v29 = _1616_v29;
        let _1617_v30;
        _1617_v30 = _dafny.Set.fromElements(p0, false);
        (_this).f14 = new BigNumber(((_module.D16.create_DC33(_1579_v5, new BigNumber((_1575_v1).length), _dafny.MultiSet.fromElements(p0, true), _1580_v6, _1617_v30)).dtor_cf54).cardinality());
      }
      let _1618_v31;
      _1618_v31 = _dafny.Seq.of(p0, _1579_v5, _1579_v5);
      let _1619_v32;
      _1619_v32 = _dafny.Seq.of(_1618_v31);
      let _1620_v33;
      _1620_v33 = _module.D12.create_DC27(_1619_v32);
      let _source29 = _1620_v33;
      if (_source29.is_DC28) {
        let _1621___mcc_h0 = (_source29).cf43;
        let _1622___mcc_h1 = (_source29).cf44;
        let _1623___mcc_h2 = (_source29).cf45;
        let _1624___mcc_h3 = (_source29).cf46;
        let _1625___mcc_h4 = (_source29).cf47;
        let _1626_cf47 = _1625___mcc_h4;
        let _1627_cf46 = _1624___mcc_h3;
        let _1628_cf45 = _1623___mcc_h2;
        let _1629_cf44 = _1622___mcc_h1;
        let _1630_cf43 = _1621___mcc_h0;
        let _1631_v34;
        let _init35 = function (_1632_i7) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f13,false);
        };
        let _nw242 = Array((new BigNumber(29)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw242.length); _i0_35++) {
          _nw242[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1631_v34 = _nw242;
        _1631_v34 = _1631_v34;
        if (_module.__default.fm0(_1584_v8, globalState)) {
          let _1633_v35;
          _1633_v35 = _dafny.Set.fromElements((_this).f13, _1627_cf46);
          let _1634_v36;
          _1634_v36 = _dafny.Map.Empty.slice().updateUnsafe((_1633_v35).Intersect(_1633_v35),(_this).f4);
          _1634_v36 = (_1634_v36).update((_1633_v35).Union(_dafny.Set.fromElements((_this).f4)), !((_this).f4) || (_1579_v5));
          _1579_v5 = (_this).f13;
          let _1635_v37;
          let _nw243 = new _module.C3();
          _nw243.__ctor(_this.f5, (_dafny.ZERO).minus(_this.f14), p0);
          _1635_v37 = _nw243;
          let _1636_v38;
          _1636_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1635_v37,(_this).f13);
          _1636_v38 = (_1636_v38).update(_1635_v37, _1627_cf46);
          _1584_v8 = _1584_v8;
          _1575_v1 = _dafny.Seq.Concat(_1575_v1, _1575_v1);
        } else {
          _1629_cf44 = new BigNumber(483);
          let _1637_v39;
          let _init36 = function (_1638_i8) {
            return (_1638_i8).multipliedBy(new BigNumber(146));
          };
          let _nw244 = Array((new BigNumber(7)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw244.length); _i0_36++) {
            _nw244[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1637_v39 = _nw244;
          let _1639_v40;
          _1639_v40 = _module.D16.create_DC32(_1637_v39);
          _1637_v39 = (_1639_v40).dtor_cf51;
          _1579_v5 = _1628_cf45;
          let _index232 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_1637_v39).length));
          (_1637_v39)[_index232] = (_1583_v7).f12;
          let _index233 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_1637_v39).length));
          (_1637_v39)[_index233] = new BigNumber(342);
          _1626_cf47 = (_1583_v7).f12;
        }
        let _1640_v41;
        _1640_v41 = _dafny.Set.fromElements((_this).fm7(globalState));
        let _1641_v42;
        _1641_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1628_cf45,_1628_cf45);
        let _1642_v43;
        _1642_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1577_v3).length),new BigNumber((_1641_v42).length));
        let _rhs218 = ((_dafny.Set.fromElements((((_1642_v43).contains((_this).f6)) ? ((_1642_v43).get((_this).f6)) : ((_1583_v7).f12)), (_1583_v7).f12)).Difference(function () {
          let _coll54 = new _dafny.Set();
          for (const _compr_54 of (_1576_v2).Elements) {
            let _1643_v44 = _compr_54;
            if ((_1576_v2).contains(_1643_v44)) {
              _coll54.add((_1643_v44).plus(new BigNumber((_dafny.Set.fromElements(true, true, false, false)).length)));
            }
          }
          return _coll54;
        }())).IsSubsetOf(_1640_v41);
        let _rhs219 = ((_1629_cf44).multipliedBy((_this).f6)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("reowmtl")).length));
        _1579_v5 = _rhs218;
        _1630_cf43 = _rhs219;
        let _1644_v45;
        let _nw245 = new _module.C2();
        _nw245.__ctor(_1575_v1, _1628_cf45);
        _1644_v45 = _nw245;
      } else {
        let _1645___mcc_h5 = (_source29).cf42;
        let _1646_cf42 = _1645___mcc_h5;
        let _1647_v46;
        _1647_v46 = _module.D9.create_DC19(_1575_v1);
        let _1648_v47;
        let _nw246 = Array((new BigNumber(3)).toNumber());
        _nw246[(_dafny.ZERO).toNumber()] = _module.D9.create_DC19(_module.__default.fm37((_this).f13, false, new BigNumber((_1575_v1).length), p0, globalState));
        _nw246[(_dafny.ONE).toNumber()] = _1647_v46;
        _nw246[(new BigNumber(2)).toNumber()] = _1647_v46;
        _1648_v47 = _nw246;
        let _1649_v48;
        _1649_v48 = _dafny.Set.fromElements(_1579_v5);
        let _pat_let_tv29 = _1580_v6;
        let _pat_let_tv30 = _1649_v48;
        let _pat_let_tv31 = globalState;
        let _pat_let_tv32 = _1575_v1;
        let _pat_let_tv33 = _1575_v1;
        let _pat_let_tv34 = _1575_v1;
        let _nw247 = Array((new BigNumber(20)).toNumber());
        _nw247[(_dafny.ZERO).toNumber()] = _module.D9.create_DC19(_1575_v1);
        _nw247[(_dafny.ONE).toNumber()] = _1647_v46;
        _nw247[(new BigNumber(2)).toNumber()] = _module.D9.create_DC19(_1575_v1);
        _nw247[(new BigNumber(3)).toNumber()] = function (_pat_let38_0) {
          return function (_1650_dt__update__tmp_h0) {
            return function (_pat_let39_0) {
              return function (_1651_dt__update_hcf25_h0) {
                return _module.D9.create_DC19(_1651_dt__update_hcf25_h0);
              }(_pat_let39_0);
            }(_module.__default.fm50(_this.f14, _pat_let_tv29, new BigNumber((_pat_let_tv30).length), !(!(!(false))), _pat_let_tv31));
          }(_pat_let38_0);
        }(_1647_v46);
        _nw247[(new BigNumber(4)).toNumber()] = _1647_v46;
        _nw247[(new BigNumber(5)).toNumber()] = _module.D9.create_DC19(_dafny.Seq.UnicodeFromString("ruvt"));
        _nw247[(new BigNumber(6)).toNumber()] = _1647_v46;
        _nw247[(new BigNumber(7)).toNumber()] = _module.D9.create_DC19(_1575_v1);
        _nw247[(new BigNumber(8)).toNumber()] = _1647_v46;
        _nw247[(new BigNumber(9)).toNumber()] = ((p0) ? (function (_pat_let40_0) {
          return function (_1654_dt__update__tmp_h2) {
            return function (_pat_let43_0) {
              return function (_1655_dt__update_hcf25_h2) {
                return _module.D9.create_DC19(_1655_dt__update_hcf25_h2);
              }(_pat_let43_0);
            }(_pat_let_tv33);
          }(_pat_let40_0);
        }(function (_pat_let41_0) {
          return function (_1652_dt__update__tmp_h1) {
            return function (_pat_let42_0) {
              return function (_1653_dt__update_hcf25_h1) {
                return _module.D9.create_DC19(_1653_dt__update_hcf25_h1);
              }(_pat_let42_0);
            }(_pat_let_tv32);
          }(_pat_let41_0);
        }(_1647_v46))) : (_1647_v46));
        _nw247[(new BigNumber(10)).toNumber()] = _1647_v46;
        _nw247[(new BigNumber(11)).toNumber()] = _1647_v46;
        _nw247[(new BigNumber(12)).toNumber()] = _1647_v46;
        _nw247[(new BigNumber(13)).toNumber()] = (((_this).f13) ? (_module.D9.create_DC19(_dafny.Seq.UnicodeFromString("khiwbj"))) : (_1647_v46));
        _nw247[(new BigNumber(14)).toNumber()] = _module.D9.create_DC19(_dafny.Seq.UnicodeFromString("qmqtcd"));
        _nw247[(new BigNumber(15)).toNumber()] = _1647_v46;
        _nw247[(new BigNumber(16)).toNumber()] = _1647_v46;
        _nw247[(new BigNumber(17)).toNumber()] = _module.D9.create_DC19(_1575_v1);
        _nw247[(new BigNumber(18)).toNumber()] = function (_pat_let44_0) {
          return function (_1656_dt__update__tmp_h3) {
            return function (_pat_let45_0) {
              return function (_1657_dt__update_hcf25_h3) {
                return _module.D9.create_DC19(_1657_dt__update_hcf25_h3);
              }(_pat_let45_0);
            }(_pat_let_tv34);
          }(_pat_let44_0);
        }(_1647_v46);
        _nw247[(new BigNumber(19)).toNumber()] = _1647_v46;
        _1648_v47 = _nw247;
        let _1658_v49;
        let _nw248 = Array((new BigNumber(9)).toNumber());
        _nw248[(_dafny.ZERO).toNumber()] = !((_this).f13);
        _nw248[(_dafny.ONE).toNumber()] = _1579_v5;
        _nw248[(new BigNumber(2)).toNumber()] = (_module.__default.fm53(_1580_v6, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), ((_1659_v5) => function (_1660_i9) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1659_v5,(_this).f13);
        })(_1579_v5)), _module.__default.safeIndex((_1583_v7).f12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), ((_1661_v5) => function (_1662_i9) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1661_v5,(_this).f13);
        })(_1579_v5))).length)), (_dafny.Map.Empty.slice().updateUnsafe((_this).f13,_1579_v5)).update(p0, (_this).f13)), globalState)).IsSubsetOf(_1649_v48);
        _nw248[(new BigNumber(3)).toNumber()] = ((p0) ? ((_this).f4) : (_module.__default.fm0(_1584_v8, globalState)));
        _nw248[(new BigNumber(4)).toNumber()] = _dafny.Seq.IsPrefixOf(_1618_v31, _1618_v31);
        _nw248[(new BigNumber(5)).toNumber()] = (_this).f13;
        _nw248[(new BigNumber(6)).toNumber()] = (_this).f4;
        _nw248[(new BigNumber(7)).toNumber()] = (_this).f13;
        _nw248[(new BigNumber(8)).toNumber()] = false;
        _1658_v49 = _nw248;
        let _index234 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_1658_v49).length));
        (_1658_v49)[_index234] = true;
        let _1663_v50;
        let _nw249 = Array((_dafny.ONE).toNumber()).fill(_module.D18.Default());
        _1663_v50 = _nw249;
        let _1664_v51;
        _1664_v51 = _module.D11.create_DC26(_1579_v5);
        let _1665_v52;
        _1665_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1664_v51,true);
        let _1666_v53;
        _1666_v53 = _dafny.Seq.of(_1665_v52);
        let _1667_v54;
        _1667_v54 = _module.D18.create_DC35(_1666_v53);
        let _index235 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_1663_v50).length));
        (_1663_v50)[_index235] = _1667_v54;
        let _index236 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_1658_v49).length));
        let _index237 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_1663_v50).length));
        let _rhs220 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_module.__default.fm37(p0, (_this).f13, _1580_v6, (_this).f13, globalState), _1575_v1), _1575_v1);
        let _rhs221 = _1667_v54;
        let _lhs138 = _1658_v49;
        let _lhs139 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_1658_v49).length));
        let _lhs140 = _1663_v50;
        let _lhs141 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_1663_v50).length));
        _lhs138[_lhs139] = _rhs220;
        _lhs140[_lhs141] = _rhs221;
        let _1668_v55;
        _1668_v55 = _module.D12.create_DC28((_this).f6, (_1583_v7).f12, p0, true, _1580_v6);
        let _rhs222 = _module.__default.safeModuloInt((_this).f6, ((!(_1579_v5)) ? (_1580_v6) : (_1580_v6)));
        let _rhs223 = ((_1649_v48).Intersect(_1649_v48)).IsSubsetOf(_dafny.Set.fromElements((_this).f13, true, (_1668_v55).dtor_cf45, (_this).f4, false));
        let _lhs142 = _this;
        _lhs142.f14 = _rhs222;
        _1579_v5 = _rhs223;
        let _1669_v56;
        let _nw250 = new _module.C1();
        _nw250.__ctor(_this.f14, new BigNumber(425), (((_this).f13) ? (_this.f5) : (_this.f5)), _this.f14, !(p0));
        _1669_v56 = _nw250;
      }
      let _pat_let_tv35 = _1584_v8;
      let _pat_let_tv36 = _1584_v8;
      let _pat_let_tv37 = _1584_v8;
      let _pat_let_tv38 = _1575_v1;
      r0 = function (_source30) {
        if (_source30.is_DC24) {
          let _1670___mcc_h6 = (_source30).cf36;
          let _1671___mcc_h7 = (_source30).cf37;
          let _1672___mcc_h8 = (_source30).cf38;
          let _1673_cf38 = _1672___mcc_h8;
          let _1674_cf37 = _1671___mcc_h7;
          let _1675_cf36 = _1670___mcc_h6;
          return new _dafny.CodePoint('w'.codePointAt(0));
        } else if (_source30.is_DC25) {
          let _1676___mcc_h9 = (_source30).cf39;
          let _1677___mcc_h10 = (_source30).cf40;
          let _1678_cf40 = _1677___mcc_h10;
          let _1679_cf39 = _1676___mcc_h9;
          return _pat_let_tv35;
        } else if (_source30.is_DC26) {
          let _1680___mcc_h11 = (_source30).cf41;
          let _1681_cf41 = _1680___mcc_h11;
          return _pat_let_tv36;
        } else {
          let _1682___mcc_h12 = (_source30).cf35;
          let _1683_cf35 = _1682___mcc_h12;
          return _pat_let_tv37;
        }
      }(function (_pat_let46_0) {
        return function (_1684_dt__update__tmp_h4) {
          return function (_pat_let47_0) {
            return function (_1685_dt__update_hcf38_h0) {
              return _module.D11.create_DC24((_1684_dt__update__tmp_h4).dtor_cf36, (_1684_dt__update__tmp_h4).dtor_cf37, _1685_dt__update_hcf38_h0);
            }(_pat_let47_0);
          }(new BigNumber((_pat_let_tv38).length));
        }(_pat_let46_0);
      }(_module.D11.create_DC24(p0, !((_this).f4), (_1583_v7).f12)));
      let _1686_v57;
      let _nw251 = Array((new BigNumber(19)).toNumber()).fill(_module.D1.Default());
      _1686_v57 = _nw251;
      r1 = _1686_v57;
      return [r0, r1];
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f5 = [];
      this._f6 = _dafny.ZERO;
      this._f4 = false;
      this._f11 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    set f5(value) {
      let _this = this;
      _this._f5 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f11, f4, f5, f6) {
      let _this = this;
      (_this)._f11 = f11;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      return;
    }
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f6))).Intersect(_dafny.MultiSet.fromElements((_this).f6));
    };
    fm7(globalState) {
      let _this = this;
      return (((_this).f6).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f6)).length), (_this).f6)).cardinality()))).multipliedBy(_module.__default.safeDivisionInt((_this).f6, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.UnicodeFromString("hkuc"))).length)));
    };
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(((_this).f11) && (((_this).f6).isLessThan(new BigNumber((_dafny.MultiSet.fromElements((_this).f4, (_this).f4)).cardinality()))));
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm15(globalState) {
      let _this = this;
      let _source31 = _module.D0.create_DC1();
      if (_source31.is_DC1) {
        return (_dafny.ZERO).minus((_this).f6);
      } else if (_source31.is_DC2) {
        let _1687___mcc_h0 = (_source31).cf1;
        let _1688_cf1 = _1687___mcc_h0;
        return _module.__default.safeDivisionInt(_1688_cf1, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f11)).length))).length));
      } else {
        let _1689___mcc_h1 = (_source31).cf0;
        let _1690_cf0 = _1689___mcc_h1;
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f4, (_this).f11), (_module.D3.create_DC8(_dafny.Seq.of(false))).dtor_cf7)).length);
      }
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _1691_v0;
      _1691_v0 = _dafny.Seq.UnicodeFromString("gteox");
      let _1692_v1;
      _1692_v1 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f11),_1691_v0);
      let _1693_v2;
      _1693_v2 = _module.D2.create_DC6();
      let _1694_v3;
      _1694_v3 = _module.D1.create_DC3((_this).f11);
      let _1695_v4;
      let _nw252 = Array((new BigNumber(17)).toNumber());
      _nw252[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f11),_1691_v0);
      _nw252[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1691_v0);
      _nw252[(new BigNumber(2)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_dafny.Seq.update(_module.__default.fm16((_this).f4, (_this).f4, (_this).f11, _1693_v2, globalState), _module.__default.safeIndex((_this).f6, new BigNumber((_module.__default.fm16((_this).f4, (_this).f4, (_this).f11, _1693_v2, globalState)).length)), new _dafny.CodePoint('l'.codePointAt(0))));
      _nw252[(new BigNumber(4)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(5)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(6)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(7)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(8)).toNumber()] = (((_1694_v3).dtor_cf2) ? (_1692_v1) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f11,_dafny.Seq.UnicodeFromString("jwkwygjor"))));
      _nw252[(new BigNumber(9)).toNumber()] = (_1692_v1).Merge(_1692_v1);
      _nw252[(new BigNumber(10)).toNumber()] = (_1692_v1).Merge(_1692_v1);
      _nw252[(new BigNumber(11)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(12)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(13)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(14)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(15)).toNumber()] = _1692_v1;
      _nw252[(new BigNumber(16)).toNumber()] = _1692_v1;
      _1695_v4 = _nw252;
      let _index238 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_1695_v4).length));
      (_1695_v4)[_index238] = (_1692_v1).Merge(_1692_v1);
      let _index239 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_1695_v4).length));
      (_1695_v4)[_index239] = _1692_v1;
      let _hi11 = (_this).fm7(globalState);
      for (let _1696_i0 = p0; _1696_i0.isLessThan(_hi11); _1696_i0 = _1696_i0.plus(_dafny.ONE)) {
        let _1697_v5;
        let _init37 = ((_1698_i0) => function (_1699_i1) {
          return (_1699_i1).minus(_1698_i0);
        })(_1696_i0);
        let _nw253 = Array((new BigNumber(23)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw253.length); _i0_37++) {
          _nw253[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1697_v5 = _nw253;
        let _1700_v6;
        _1700_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f4);
        let _1701_v7;
        _1701_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1697_v5,_1700_v6);
        let _1702_v8;
        _1702_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1696_i0);
        _1701_v7 = (_1701_v7).update(_1697_v5, _dafny.Map.Empty.slice().updateUnsafe((((_1702_v8).contains((_this).f11)) ? ((_1702_v8).get((_this).f11)) : ((_this).f6)),(_this).f4));
        let _1703_v9;
        _1703_v9 = _dafny.MultiSet.fromElements(_1696_i0);
        if ((_1703_v9).IsDisjointFrom(_1703_v9)) {
          let _1704_v10;
          let _nw254 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.of());
          _1704_v10 = _nw254;
          let _1705_v11;
          _1705_v11 = _dafny.Seq.of(_1697_v5);
          let _index240 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1704_v10).length));
          (_1704_v10)[_index240] = _1705_v11;
          let _index241 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1704_v10).length));
          (_1704_v10)[_index241] = _1705_v11;
          let _1706_v12;
          _1706_v12 = new BigNumber(85);
          let _1707_v13;
          _1707_v13 = false;
          let _rhs224 = (p0).minus((_dafny.ZERO).minus((_this).f6));
          let _rhs225 = !((_this).f4);
          _1706_v12 = _rhs224;
          _1707_v13 = _rhs225;
          let _1708_v14;
          let _init38 = ((_1709_v12) => function (_1710_i2) {
            return _module.D0.create_DC0(_1709_v12);
          })(_1706_v12);
          let _nw255 = Array((new BigNumber(16)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw255.length); _i0_38++) {
            _nw255[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1708_v14 = _nw255;
          _1708_v14 = _1708_v14;
          let _1711_v15;
          _1711_v15 = _dafny.Set.fromElements(new BigNumber(-882));
          _1706_v12 = (new BigNumber((_1711_v15).length)).plus(p0);
          let _1712_v16;
          _1712_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1697_v5);
          let _1713_v17;
          _1713_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1712_v16,(((_this).f4) ? ((((_1702_v8).contains(false)) ? ((_1702_v8).get(false)) : ((_this).fm9(!(_1707_v13), (_this).f11, p0, globalState)))) : (_1696_i0)));
          _1713_v17 = (_1713_v17).update((_1712_v16).update((_this).f4, _1697_v5), _1696_i0);
        } else {
          let _1714_v18;
          let _nw256 = new _module.C0();
          _nw256.__ctor(_1696_i0);
          _1714_v18 = _nw256;
          let _1715_v19;
          _1715_v19 = new BigNumber(759);
          let _1716_v20;
          _1716_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(410));
          let _1717_v21;
          _1717_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1716_v20);
          _1715_v19 = new BigNumber((_1717_v21).length);
          _1715_v19 = ((_1714_v18).f12).multipliedBy((_dafny.ZERO).minus(_1715_v19));
          let _1718_v22;
          _1718_v22 = true;
          let _index242 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_1697_v5).length));
          (_1697_v5)[_index242] = (_module.__default.fm3(_1715_v19, _1718_v22, globalState)).plus(_1715_v19);
          let _index243 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_1697_v5).length));
          (_1697_v5)[_index243] = (((_1702_v8).contains((_this).f4)) ? ((_1702_v8).get((_this).f4)) : (_1696_i0));
          let _1719_v23;
          _1719_v23 = _module.D2.create_DC7((_1714_v18).f12);
          let _index244 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_1697_v5).length));
          let _index245 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_1697_v5).length));
          let _rhs226 = _1718_v22;
          let _rhs227 = (_dafny.ZERO).minus(((_1714_v18).f12).minus((new BigNumber((_module.__default.fm16((_this).f4, _1718_v22, (_this).f11, _1693_v2, globalState)).length)).plus(p0)));
          let _rhs228 = (p0).multipliedBy((_1719_v23).dtor_cf6);
          let _rhs229 = (_this).f6;
          let _rhs230 = (_1694_v3).dtor_cf2;
          let _lhs143 = _1697_v5;
          let _lhs144 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_1697_v5).length));
          let _lhs145 = _1697_v5;
          let _lhs146 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_1697_v5).length));
          _1718_v22 = _rhs226;
          _lhs143[_lhs144] = _rhs227;
          _lhs145[_lhs146] = _rhs228;
          _1715_v19 = _rhs229;
          _1718_v22 = _rhs230;
          _1715_v19 = new BigNumber((_1703_v9).cardinality());
        }
        let _1720_v24;
        _1720_v24 = _dafny.Set.fromElements((_this).f4);
        let _1721_v25;
        _1721_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1720_v24,(_this).f11);
        let _1722_v26;
        _1722_v26 = _dafny.Set.fromElements(new BigNumber((_1721_v25).length), p0, p0, p0);
        _1722_v26 = _1722_v26;
        let _1723_v27;
        _1723_v27 = new BigNumber(-271);
        _1723_v27 = _1723_v27;
      }
      let _1724_i3;
      _1724_i3 = _dafny.ZERO;
      L7: {
        while ((_this).f11) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1724_i3)) {
              break L7;
            }
            _1724_i3 = (_1724_i3).plus(_dafny.ONE);
            let _1725_v28;
            let _nw257 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
            _1725_v28 = _nw257;
            _1725_v28 = _1725_v28;
            let _1726_v29;
            _1726_v29 = false;
            let _1727_v30;
            _1727_v30 = _dafny.Seq.of(new BigNumber(583));
            _1726_v29 = _dafny.areEqual(_1727_v30, _1727_v30);
            _1726_v29 = (_this).f11;
            let _1728_v31;
            _1728_v31 = new BigNumber(-981);
            let _rhs231 = (_this).f11;
            let _rhs232 = new BigNumber(759);
            _1726_v29 = _rhs231;
            _1728_v31 = _rhs232;
          }
        }
      }
      let _1729_v32;
      _1729_v32 = _dafny.MultiSet.fromElements((_this).f11);
      let _1730_v33;
      _1730_v33 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f6),(_this).f4);
      if ((((_this).f6).plus((_this).f6)).isEqualTo((p0).multipliedBy((((_1729_v32).contains((_this).f4)) ? ((_1729_v32).get((_this).f4)) : (new BigNumber((_1730_v33).length)))))) {
        let _1731_v34;
        _1731_v34 = _dafny.Seq.of(new BigNumber(327), p0);
        let _1732_v35;
        let _nw258 = new _module.C0();
        _nw258.__ctor(new BigNumber((_1731_v34).length));
        _1732_v35 = _nw258;
        let _1733_v36;
        _1733_v36 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("wfhdckno"));
        let _1734_v37;
        let _nw259 = new _module.C0();
        _nw259.__ctor(new BigNumber((_1733_v36).length));
        _1734_v37 = _nw259;
        let _1735_v38;
        let _init39 = ((_1736_v0, _1737_v34) => function (_1738_i4) {
          return _dafny.Seq.update(_dafny.Seq.Concat(_1736_v0, _1736_v0), _module.__default.safeIndex((_1737_v34)[_module.__default.safeIndex(new BigNumber(239), new BigNumber((_1737_v34).length))], new BigNumber((_dafny.Seq.Concat(_1736_v0, _1736_v0)).length)), new _dafny.CodePoint('j'.codePointAt(0)));
        })(_1691_v0, _1731_v34);
        let _nw260 = Array((new BigNumber(28)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw260.length); _i0_39++) {
          _nw260[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1735_v38 = _nw260;
        let _index246 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1735_v38).length));
        (_1735_v38)[_index246] = _1691_v0;
        let _1739_v39;
        _1739_v39 = new _dafny.CodePoint('s'.codePointAt(0));
        let _index247 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1735_v38).length));
        (_1735_v38)[_index247] = _dafny.Seq.Concat(_1691_v0, _dafny.Seq.Concat(_1691_v0, _dafny.Seq.update(_1691_v0, _module.__default.safeIndex((_this).f6, new BigNumber((_1691_v0).length)), _1739_v39)));
        let _1740_v40;
        _1740_v40 = _dafny.Set.fromElements(p0);
        let _1741_v41;
        let _nw261 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _1741_v41 = _nw261;
        let _1742_v42;
        _1742_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1740_v40).Intersect(_1740_v40)).length),_1741_v41);
        _1742_v42 = (_1742_v42).update(new BigNumber(-824), _1741_v41);
        let _1743_v43;
        _1743_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1691_v0,(_1734_v37).f12);
        let _1744_v44;
        let _nw262 = Array((new BigNumber(6)).toNumber());
        _nw262[(_dafny.ZERO).toNumber()] = (_this).f4;
        _nw262[(_dafny.ONE).toNumber()] = (_this).f11;
        _nw262[(new BigNumber(2)).toNumber()] = (_this).f11;
        _nw262[(new BigNumber(3)).toNumber()] = (_this).f4;
        _nw262[(new BigNumber(4)).toNumber()] = (_this).f4;
        _nw262[(new BigNumber(5)).toNumber()] = (_this).f11;
        _1744_v44 = _nw262;
        let _1745_v45;
        _1745_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1744_v44,(_this).f11);
        _1743_v43 = (_1743_v43).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), ((_1746_v39) => function (_1747_i5) {
          return _1746_v39;
        })(_1739_v39)), new BigNumber((_module.__default.fm18(_1691_v0, new BigNumber((_1745_v45).length), globalState)).length));
      } else {
        let _1748_v46;
        _1748_v46 = new BigNumber(602);
        let _1749_v47;
        _1749_v47 = new _dafny.CodePoint('j'.codePointAt(0));
        let _rhs233 = p0;
        let _rhs234 = p0;
        let _rhs235 = _1749_v47;
        _1748_v46 = _rhs233;
        _1748_v46 = _rhs234;
        _1749_v47 = _rhs235;
        if ((_this).f4) {
          _1748_v46 = p0;
          _1748_v46 = (((p0).isLessThan(_1748_v46)) ? (new BigNumber((_module.__default.fm16((_this).f11, (_this).f4, (_module.D3.create_DC9((_this).f4, _1748_v46, (_this).f11)).dtor_cf10, _1693_v2, globalState)).length)) : (_1748_v46));
          let _1750_v48;
          _1750_v48 = _dafny.Set.fromElements(new BigNumber(274), p0, new BigNumber((_1691_v0).length));
          _1750_v48 = (_1750_v48).Union(_dafny.Set.fromElements(p0));
          let _1751_v49;
          let _init40 = ((_1752_p0) => function (_1753_i6) {
            return (_1753_i6).multipliedBy(_1752_p0);
          })(p0);
          let _nw263 = Array((new BigNumber(11)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw263.length); _i0_40++) {
            _nw263[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1751_v49 = _nw263;
          let _index248 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_1751_v49).length));
          (_1751_v49)[_index248] = (((_this).f4) ? ((_this).f6) : (p0));
          let _index249 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_1751_v49).length));
          (_1751_v49)[_index249] = new BigNumber(348);
          let _1754_v50;
          let _nw264 = new _module.C0();
          _nw264.__ctor((((_this).f11) ? (new BigNumber(-268)) : (p0)));
          _1754_v50 = _nw264;
        } else {
          let _1755_v51;
          _1755_v51 = _module.D4.create_DC12(_1749_v47);
          _1749_v47 = (_1755_v51).dtor_cf13;
          let _1756_v52;
          _1756_v52 = _dafny.Seq.of((new BigNumber((_module.__default.fm16((_this).f11, _module.__default.fm0(_1749_v47, globalState), (_this).f11, _1693_v2, globalState)).length)).multipliedBy((_this).fm15(globalState)));
          let _1757_v53;
          _1757_v53 = _dafny.Seq.of((_this).f4);
          let _1758_v54;
          let _nw265 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _1758_v54 = _nw265;
          let _1759_v55;
          let _nw266 = Array((new BigNumber(19)).toNumber());
          _nw266[(_dafny.ZERO).toNumber()] = _1757_v53;
          _nw266[(_dafny.ONE).toNumber()] = _dafny.Seq.of((_this).f4);
          _nw266[(new BigNumber(2)).toNumber()] = _module.__default.fm19((_this).f11, _module.__default.fm3((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1749_v47,_1758_v54)).length)), (_this).f11, globalState), globalState);
          _nw266[(new BigNumber(3)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(4)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(5)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(6)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(7)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(8)).toNumber()] = (((_this).f4) ? (_1757_v53) : (_1757_v53));
          _nw266[(new BigNumber(9)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(10)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_1757_v53, _module.__default.safeIndex(_1748_v46, new BigNumber((_1757_v53).length)), true), _module.__default.safeIndex((_this).f6, new BigNumber((_dafny.Seq.update(_1757_v53, _module.__default.safeIndex(_1748_v46, new BigNumber((_1757_v53).length)), true)).length)), false);
          _nw266[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_1757_v53, _1757_v53);
          _nw266[(new BigNumber(13)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(14)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(15)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(16)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(17)).toNumber()] = _1757_v53;
          _nw266[(new BigNumber(18)).toNumber()] = _1757_v53;
          _1759_v55 = _nw266;
          let _index250 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1759_v55).length));
          (_1759_v55)[_index250] = _dafny.Seq.Concat((_module.D4.create_DC13((_this).f6, (_this).f11, _1748_v46, _1757_v53)).dtor_cf17, _1757_v53);
          let _index251 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1759_v55).length));
          let _rhs236 = _1756_v52;
          let _rhs237 = p0;
          let _rhs238 = new BigNumber(-794);
          let _rhs239 = _1757_v53;
          let _lhs147 = _1759_v55;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_1759_v55).length));
          _1756_v52 = _rhs236;
          _1748_v46 = _rhs237;
          _1748_v46 = _rhs238;
          _lhs147[_lhs148] = _rhs239;
          let _1760_v56;
          _1760_v56 = _module.D3.create_DC10(p0);
          let _1761_v57;
          _1761_v57 = _module.D3.create_DC11(_1760_v56);
          let _1762_v58;
          let _nw267 = Array((new BigNumber(6)).toNumber());
          _nw267[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw267[(_dafny.ONE).toNumber()] = (_this).f4;
          _nw267[(new BigNumber(2)).toNumber()] = (_this).f4;
          _nw267[(new BigNumber(3)).toNumber()] = (_this).f4;
          _nw267[(new BigNumber(4)).toNumber()] = (_this).f4;
          _nw267[(new BigNumber(5)).toNumber()] = (_this).f4;
          _1762_v58 = _nw267;
          let _1763_v59;
          _1763_v59 = _dafny.MultiSet.fromElements(_1762_v58);
          let _1764_v60;
          _1764_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1761_v57,(_1763_v59).Difference(_1763_v59));
          _1764_v60 = (_1764_v60).update(_1761_v57, _1763_v59);
          let _1765_v61;
          _1765_v61 = _dafny.MultiSet.fromElements((((_this).f4) ? (new BigNumber(-133)) : (p0)), _1748_v46, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-336)), function (_1766_i7) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length), p0);
          _1765_v61 = (_dafny.MultiSet.FromArray(_1756_v52));
          let _1767_v62;
          _1767_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1768_v63;
          _1768_v63 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1765_v61).cardinality()),_1691_v0);
          let _1769_v64;
          _1769_v64 = _dafny.Seq.of(_1768_v63, _1768_v63);
          let _index252 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_1758_v54).length));
          (_1758_v54)[_index252] = (((_this).f4) ? (_module.__default.fm3((((_1767_v62).contains(_1748_v46)) ? ((_1767_v62).get(_1748_v46)) : (_1748_v46)), (_this).f4, globalState)) : (new BigNumber(((_1769_v64)[_module.__default.safeIndex(new BigNumber((_1691_v0).length), new BigNumber((_1769_v64).length))]).length)));
          let _index253 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_1758_v54).length));
          (_1758_v54)[_index253] = (((_this).f11) ? (p0) : (new BigNumber(((_module.D3.create_DC8((_module.__default.fm20((_this).f6, (_this).f4, globalState)).dtor_cf7)).dtor_cf7).length)));
        }
        let _1770_v65;
        let _init41 = ((_1771_v46, _1772_v0, _1773_p0) => function (_1774_i8) {
          return (_dafny.MultiSet.fromElements(_1771_v46, new BigNumber((_1772_v0).length))).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1773_p0, _1771_v46, _1771_v46));
        })(_1748_v46, _1691_v0, p0);
        let _nw268 = Array((new BigNumber(8)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw268.length); _i0_41++) {
          _nw268[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1770_v65 = _nw268;
        _1770_v65 = _1770_v65;
        _1730_v33 = (_1730_v33).update(new BigNumber(451), (_this).f4);
        let _1775_v66;
        _1775_v66 = true;
        let _1776_v67;
        _1776_v67 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f6);
        _1775_v66 = (_dafny.MultiSet.fromElements((_this).f6)).contains(new BigNumber((_1776_v67).length));
      }
      let _1777_v68;
      _1777_v68 = true;
      _1777_v68 = (_this).f4;
      let _1778_v69;
      _1778_v69 = new BigNumber(678);
      _1778_v69 = (_dafny.ZERO).minus(((!((_1777_v68) === (_1777_v68))) ? ((_this).f6) : (p0)));
      let _1779_v70;
      _1779_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1777_v68,p0);
      r0 = _1779_v70;
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = [];
      let _1780_v0;
      _1780_v0 = _dafny.Seq.of((_this).f11, p0, (_this).f4);
      let _1781_v1;
      _1781_v1 = _dafny.Seq.of((_1780_v0)[_module.__default.safeIndex((_this).f6, new BigNumber((_1780_v0).length))], p0, (_this).f4);
      let _1782_v2;
      _1782_v2 = _module.D3.create_DC8(_dafny.Seq.update(_1780_v0, _module.__default.safeIndex((_this).f6, new BigNumber((_1780_v0).length)), (_this).f11));
      let _source32 = _1782_v2;
      if (_source32.is_DC9) {
        let _1783___mcc_h0 = (_source32).cf8;
        let _1784___mcc_h1 = (_source32).cf9;
        let _1785___mcc_h2 = (_source32).cf10;
        let _1786_cf10 = _1785___mcc_h2;
        let _1787_cf9 = _1784___mcc_h1;
        let _1788_cf8 = _1783___mcc_h0;
        let _1789_v3;
        let _nw269 = Array((new BigNumber(9)).toNumber()).fill([]);
        _1789_v3 = _nw269;
        let _1790_v4;
        let _init42 = function (_1791_i0) {
          return !((_this).f11);
        };
        let _nw270 = Array((new BigNumber(18)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw270.length); _i0_42++) {
          _nw270[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1790_v4 = _nw270;
        let _index254 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_1790_v4).length));
        (_1790_v4)[_index254] = (_this).f11;
        let _1792_v5;
        _1792_v5 = _dafny.Seq.of(_1789_v3, _1789_v3, _1789_v3, _1789_v3, _1789_v3);
        let _1793_v6;
        _1793_v6 = _dafny.MultiSet.fromElements(_1790_v4, _1790_v4);
        let _1794_v7;
        _1794_v7 = _dafny.Seq.UnicodeFromString("rmjgkq");
        let _index255 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_1790_v4).length));
        let _rhs240 = (_this).fm9(_1786_cf10, (_this).f11, _1787_cf9, globalState);
        let _rhs241 = (_1792_v5)[_module.__default.safeIndex((((_1793_v6).contains(_1790_v4)) ? ((_1793_v6).get(_1790_v4)) : (new BigNumber((_1794_v7).length))), new BigNumber((_1792_v5).length))];
        let _rhs242 = (_dafny.ZERO).minus((new BigNumber((_1794_v7).length)).plus(_module.__default.safeModuloInt((_this).fm9(true, (_this).f11, _1787_cf9, globalState), (_this).f6)));
        let _rhs243 = ((p0) ? (new BigNumber((_1794_v7).length)) : (_1787_cf9));
        let _rhs244 = _1786_cf10;
        let _lhs149 = _1790_v4;
        let _lhs150 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_1790_v4).length));
        _1787_cf9 = _rhs240;
        _1789_v3 = _rhs241;
        _1787_cf9 = _rhs242;
        _1787_cf9 = _rhs243;
        _lhs149[_lhs150] = _rhs244;
        if ((_this).f4) {
          _1786_cf10 = false;
          let _arr7 = _this.f5;
          let _index256 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_this.f5).length));
          _arr7[_index256] = _1790_v4;
          let _arr8 = _this.f5;
          let _index257 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_this.f5).length));
          _arr8[_index257] = (_module.D2.create_DC5(_1790_v4)).dtor_cf5;
          let _1795_v8;
          _1795_v8 = _module.D3.create_DC10(new BigNumber((_1794_v7).length));
          let _1796_v9;
          let _nw271 = Array((new BigNumber(25)).toNumber());
          _nw271[(_dafny.ZERO).toNumber()] = ((_1786_cf10) ? (_1795_v8) : (_1795_v8));
          _nw271[(_dafny.ONE).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(2)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(3)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(4)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(5)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(6)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(7)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(8)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(9)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(10)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(11)).toNumber()] = _module.D3.create_DC10(_1787_cf9);
          _nw271[(new BigNumber(12)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(13)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(14)).toNumber()] = _module.D3.create_DC10((_this).f6);
          _nw271[(new BigNumber(15)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(16)).toNumber()] = _module.D3.create_DC10((_this).f6);
          _nw271[(new BigNumber(17)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(18)).toNumber()] = ((_1786_cf10) ? (_1795_v8) : (_1795_v8));
          _nw271[(new BigNumber(19)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(20)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(21)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(22)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(23)).toNumber()] = _1795_v8;
          _nw271[(new BigNumber(24)).toNumber()] = _1795_v8;
          _1796_v9 = _nw271;
          let _index258 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1796_v9).length));
          (_1796_v9)[_index258] = _1795_v8;
          let _index259 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1796_v9).length));
          (_1796_v9)[_index259] = _1795_v8;
          _1787_cf9 = _module.__default.safeModuloInt(_1787_cf9, (_this).f6);
          let _1797_v10;
          _1797_v10 = _dafny.Map.Empty.slice().updateUnsafe(true,(_1790_v4)[_module.__default.safeIndex(new BigNumber(874), new BigNumber((_1790_v4).length))]);
          let _1798_v11;
          _1798_v11 = _dafny.Seq.of((_1797_v10).update(_1786_cf10, _1788_cf8));
          _1797_v10 = ((_1798_v11)[_module.__default.safeIndex((_this).f6, new BigNumber((_1798_v11).length))]).Merge(_1797_v10);
        } else {
          let _1799_v12;
          _1799_v12 = _dafny.Map.Empty.slice().updateUnsafe(!((_1781_v1)[_module.__default.safeIndex(_1787_cf9, new BigNumber((_1781_v1).length))]),(_this).f6);
          let _1800_v13;
          let _nw272 = new _module.C0();
          _nw272.__ctor((((_1799_v12).contains(_1786_cf10)) ? ((_1799_v12).get(_1786_cf10)) : ((_this).f6)));
          _1800_v13 = _nw272;
          let _1801_v14;
          _1801_v14 = _module.D2.create_DC7(_module.__default.safeModuloInt((_this).f6, new BigNumber(433)));
          let _1802_v15;
          _1802_v15 = _dafny.Set.fromElements((_1800_v13).f12);
          let _1803_v16;
          _1803_v16 = _dafny.MultiSet.fromElements(_1802_v15, _1802_v15, _1802_v15);
          let _1804_v17;
          _1804_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1800_v13,_1790_v4);
          let _1805_v18;
          _1805_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(376),new BigNumber((_1804_v17).length));
          _1801_v14 = ((!((_1803_v16).update(_1802_v15, _module.__default.abs((_1800_v13).f12))).contains(_dafny.Set.fromElements((_this).f6))) ? (_module.D2.create_DC7(new BigNumber((_1805_v18).length))) : (_1801_v14));
          let _1806_v19;
          _1806_v19 = new _dafny.CodePoint('s'.codePointAt(0));
          r0 = _1806_v19;
          _1786_cf10 = (_this).f4;
          _1787_cf9 = _1787_cf9;
        }
        let _1807_v20;
        let _nw273 = new _module.C0();
        _nw273.__ctor(_module.__default.safeDivisionInt(new BigNumber((_1780_v0).length), (_this).f6));
        _1807_v20 = _nw273;
        let _index260 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_1790_v4).length));
        (_1790_v4)[_index260] = (_1807_v20).fm17(p0, globalState);
      } else if (_source32.is_DC10) {
        let _1808___mcc_h3 = (_source32).cf11;
        let _1809_cf11 = _1808___mcc_h3;
        let _1810_v21;
        _1810_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber(137), _1809_cf11),(((_this).f11) ? ((_this).f11) : (p0)));
        _1810_v21 = (_1810_v21).update(_1809_cf11, (_this).f11);
        let _1811_v22;
        _1811_v22 = false;
        let _1812_v23;
        let _nw274 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _1812_v23 = _nw274;
        let _1813_v24;
        _1813_v24 = _dafny.Seq.of(_1809_cf11);
        let _rhs245 = p0;
        let _rhs246 = _1809_cf11;
        let _rhs247 = !_dafny.areEqual(_1813_v24, _1813_v24);
        let _rhs248 = _1812_v23;
        _1811_v22 = _rhs245;
        _1809_cf11 = _rhs246;
        _1811_v22 = _rhs247;
        _1812_v23 = _rhs248;
        let _1814_v25;
        _1814_v25 = _dafny.Seq.UnicodeFromString("gan");
        let _1815_v26;
        _1815_v26 = new _dafny.CodePoint('j'.codePointAt(0));
        _1814_v25 = (((((_this).f4) ? (_module.__default.fm0(_1815_v26, globalState)) : (!((_this).f4)))) ? (_1814_v25) : (_dafny.Seq.Concat(_1814_v25, _1814_v25)));
        let _1816_v27;
        let _nw275 = Array((new BigNumber(24)).toNumber()).fill([]);
        _1816_v27 = _nw275;
        _1816_v27 = _1816_v27;
      } else if (_source32.is_DC8) {
        let _1817___mcc_h4 = (_source32).cf7;
        let _1818_cf7 = _1817___mcc_h4;
        let _1819_v28;
        let _nw276 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _1819_v28 = _nw276;
        let _1820_v29;
        _1820_v29 = _module.D1.create_DC4((_dafny.ZERO).minus((_this).f6), (_this).f6);
        let _1821_v30;
        _1821_v30 = _dafny.Seq.of((_1820_v29).dtor_cf3);
        let _index261 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length));
        (_1819_v28)[_index261] = (_1821_v30)[_module.__default.safeIndex((_this).f6, new BigNumber((_1821_v30).length))];
        let _1822_v31;
        _1822_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber(476));
        let _1823_v32;
        _1823_v32 = _dafny.MultiSet.fromElements(_1822_v31, _1822_v31);
        let _1824_v33;
        _1824_v33 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1825_v34;
        _1825_v34 = _dafny.Seq.of(_1824_v33);
        let _1826_v35;
        _1826_v35 = _dafny.MultiSet.fromElements(p0, (_this).fm8(new BigNumber(45), (_this).f11, _1825_v34, _1821_v30, globalState), p0, (_this).f11, (_this).f4);
        let _index262 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length));
        (_1819_v28)[_index262] = new BigNumber(((_module.__default.fm21((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f6)), _1823_v32, new BigNumber(193), (_this).f6, globalState)).Difference(_1826_v35)).cardinality());
        let _1827_v36;
        let _nw277 = new _module.C0();
        _nw277.__ctor((_dafny.ZERO).minus(((_1819_v28)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length))]).minus((_1819_v28)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length))])));
        _1827_v36 = _nw277;
        let _1828_v37;
        let _nw278 = Array((new BigNumber(27)).toNumber()).fill(false);
        _1828_v37 = _nw278;
        let _index263 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1828_v37).length));
        (_1828_v37)[_index263] = ((_this).f4) || ((_this).f11);
        let _index264 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1828_v37).length));
        (_1828_v37)[_index264] = !((_1780_v0)[_module.__default.safeIndex((new BigNumber(175)).plus((_1819_v28)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length))]), new BigNumber((_1780_v0).length))]);
        let _1829_v38;
        _1829_v38 = _module.D4.create_DC13((_1819_v28)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length))], (_this).f11, new BigNumber((_1825_v34).length), _1781_v1);
        let _1830_v39;
        _1830_v39 = _dafny.Set.fromElements((_1829_v38).dtor_cf15, p0, (_this).f4);
        let _index265 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1828_v37).length));
        let _index266 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length));
        let _index267 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length));
        let _rhs249 = (_1781_v1)[_module.__default.safeIndex(new BigNumber(((_1830_v39).Union(_1830_v39)).length), new BigNumber((_1781_v1).length))];
        let _rhs250 = ((((p0) ? ((_this).f4) : (p0))) ? ((_this).f6) : ((_1827_v36).f12));
        let _rhs251 = (_1821_v30)[_module.__default.safeIndex((_this).f6, new BigNumber((_1821_v30).length))];
        let _lhs151 = _1828_v37;
        let _lhs152 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_1828_v37).length));
        let _lhs153 = _1819_v28;
        let _lhs154 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length));
        let _lhs155 = _1819_v28;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_1819_v28).length));
        _lhs151[_lhs152] = _rhs249;
        _lhs153[_lhs154] = _rhs250;
        _lhs155[_lhs156] = _rhs251;
      } else {
        let _1831___mcc_h5 = (_source32).cf12;
        let _1832_cf12 = _1831___mcc_h5;
        let _1833_v40;
        let _init43 = function (_1834_i1) {
          return _module.__default.safeModuloInt(_1834_i1, (_this).f6);
        };
        let _nw279 = Array((new BigNumber(19)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw279.length); _i0_43++) {
          _nw279[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1833_v40 = _nw279;
        let _index268 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length));
        (_1833_v40)[_index268] = new BigNumber(257);
        let _index269 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length));
        (_1833_v40)[_index269] = (_this).f6;
        let _index270 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length));
        (_1833_v40)[_index270] = (_this).f6;
        let _1835_v41;
        _1835_v41 = false;
        let _1836_v42;
        let _nw280 = Array((new BigNumber(26)).toNumber()).fill(false);
        _1836_v42 = _nw280;
        let _1837_v43;
        _1837_v43 = _dafny.MultiSet.fromElements(p0, !(_1835_v41), (_this).f4);
        let _1838_v44;
        _1838_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_1837_v43);
        let _1839_v45;
        _1839_v45 = _dafny.MultiSet.fromElements(_1837_v43, _dafny.MultiSet.FromArray(_1781_v1), (((_1838_v44).contains(_1835_v41)) ? ((_1838_v44).get(_1835_v41)) : (_1837_v43)), _1837_v43);
        let _index271 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1836_v42).length));
        (_1836_v42)[_index271] = (_1839_v45).IsSubsetOf(_1839_v45);
        let _index272 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length));
        let _index273 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1836_v42).length));
        let _index274 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length));
        let _rhs252 = true;
        let _rhs253 = (_1833_v40)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length))];
        let _rhs254 = (_this).f11;
        let _rhs255 = _module.__default.safeModuloInt((_1833_v40)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length))], (((_this).f11) ? ((_this).f6) : ((_1833_v40)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length))])));
        let _lhs157 = _1833_v40;
        let _lhs158 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length));
        let _lhs159 = _1836_v42;
        let _lhs160 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1836_v42).length));
        let _lhs161 = _1833_v40;
        let _lhs162 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length));
        _1835_v41 = _rhs252;
        _lhs157[_lhs158] = _rhs253;
        _lhs159[_lhs160] = _rhs254;
        _lhs161[_lhs162] = _rhs255;
        if (true) {
          let _1840_v46;
          _1840_v46 = _module.D0.create_DC2(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements((_this).f6, (_this).f6))).length));
          let _1841_v47;
          _1841_v47 = _dafny.MultiSet.fromElements((_1833_v40)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length))], (_1840_v46).dtor_cf1);
          let _1842_v48;
          _1842_v48 = _module.D3.create_DC9((_this).f11, new BigNumber((_1841_v47).cardinality()), (_this).f11);
          let _1843_v49;
          _1843_v49 = _dafny.Seq.of((_1842_v48).dtor_cf9, (_1833_v40)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length))]);
          let _1844_v50;
          let _1845_v51;
          let _out51;
          let _out52;
          let _outcollector22 = _module.__default.m0(_module.__default.fm22(globalState), (p0) === ((_this).f4), (_dafny.MultiSet.FromArray(_1843_v49)).Intersect(_1841_v47), _1835_v41, globalState);
          _out51 = _outcollector22[0];
          _out52 = _outcollector22[1];
          _1844_v50 = _out51;
          _1845_v51 = _out52;
          let _1846_v52;
          _1846_v52 = _dafny.Map.Empty.slice().updateUnsafe(((_1833_v40)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length))]).plus(_1845_v51),new BigNumber(200));
          _1846_v52 = (_1846_v52).update(_module.__default.fm3(_1845_v51, _1835_v41, globalState), (_1843_v49)[_module.__default.safeIndex(new BigNumber(-562), new BigNumber((_1843_v49).length))]);
          let _1847_v53;
          _1847_v53 = _dafny.Seq.UnicodeFromString("owcvyfr");
          let _1848_v54;
          _1848_v54 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1847_v53).length)).multipliedBy(_1845_v51),_1780_v0);
          _1848_v54 = (_1848_v54).update(_1845_v51, _dafny.Seq.of((_1836_v42)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1836_v42).length))], (_this).f11, (_this).f11, p0));
          let _1849_v55;
          _1849_v55 = new _dafny.CodePoint('d'.codePointAt(0));
          let _index275 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1836_v42).length));
          (_1836_v42)[_index275] = _module.__default.fm0(_1849_v55, globalState);
          let _1850_v56;
          _1850_v56 = _dafny.Set.fromElements((_1836_v42)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_1836_v42).length))]);
          let _1851_v57;
          _1851_v57 = _dafny.Seq.of(_1833_v40, _1833_v40);
          let _rhs256 = _module.__default.fm23(_1845_v51, (_this).f6, globalState);
          let _rhs257 = _1851_v57;
          let _rhs258 = _1835_v41;
          _1850_v56 = _rhs256;
          _1851_v57 = _rhs257;
          _1844_v50 = _rhs258;
        } else {
          _1835_v41 = !((_this).f11);
          let _1852_v58;
          _1852_v58 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1837_v43).cardinality()),true);
          let _1853_v59;
          _1853_v59 = _dafny.MultiSet.fromElements(_1852_v58);
          let _1854_v60;
          _1854_v60 = _dafny.Seq.of(_1852_v58, _1852_v58);
          let _1855_v61;
          _1855_v61 = _dafny.Seq.UnicodeFromString("wgpjpb");
          let _1856_v62;
          _1856_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_1853_v59).Union(_dafny.MultiSet.fromElements(_1852_v58, _1852_v58, _1852_v58, (_1854_v60)[_module.__default.safeIndex(new BigNumber((_1855_v61).length), new BigNumber((_1854_v60).length))], _1852_v58)));
          _1856_v62 = (_1856_v62).update((_1833_v40)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length))], _1853_v59);
          let _index276 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_1833_v40).length));
          (_1833_v40)[_index276] = ((_this).f6).minus((_this).f6);
          let _index277 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_1836_v42).length));
          (_1836_v42)[_index277] = (((_this).fm9(p0, p0, (_this).fm15(globalState), globalState)).multipliedBy((((_1837_v43).contains(p0)) ? ((_1837_v43).get(p0)) : ((_this).f6)))).isLessThanOrEqualTo(new BigNumber(803));
          let _1857_v63;
          _1857_v63 = new _dafny.CodePoint('e'.codePointAt(0));
          r0 = _1857_v63;
        }
      }
      let _1858_i2;
      _1858_i2 = _dafny.ZERO;
      L8: {
        while ((_this).f4) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1858_i2)) {
              break L8;
            }
            _1858_i2 = (_1858_i2).plus(_dafny.ONE);
            let _1859_v64;
            let _init44 = ((_1860_p0) => function (_1861_i3) {
              return _1860_p0;
            })(p0);
            let _nw281 = Array((new BigNumber(21)).toNumber());
            for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw281.length); _i0_44++) {
              _nw281[_i0_44] = _init44(new BigNumber(_i0_44));
            }
            _1859_v64 = _nw281;
            _1859_v64 = _1859_v64;
            let _1862_v65;
            _1862_v65 = true;
            _1862_v65 = (_this).f4;
            let _1863_v66;
            _1863_v66 = _dafny.MultiSet.fromElements(new BigNumber(126), (_this).f6);
            let _1864_v67;
            _1864_v67 = _module.D0.create_DC2(new BigNumber((_1863_v66).cardinality()));
            let _1865_v68;
            _1865_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_1864_v67).dtor_cf1);
            _1865_v68 = (_1865_v68).update((_this).f4, ((_this).f6).plus((_this).f6));
            let _index278 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1859_v64).length));
            (_1859_v64)[_index278] = !(p0);
            let _index279 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1859_v64).length));
            (_1859_v64)[_index279] = p0;
          }
        }
      }
      let _1866_v69;
      _1866_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      _1866_v69 = (_1866_v69).Merge((_1866_v69).Merge(_1866_v69));
      _1780_v0 = _dafny.Seq.Concat(_1780_v0, _1781_v1);
      let _hi12 = _module.__default.safeModuloInt((_this).f6, (_this).f6);
      for (let _1867_i4 = (_this).f6; _1867_i4.isLessThan(_hi12); _1867_i4 = _1867_i4.plus(_dafny.ONE)) {
        let _1868_v70;
        _1868_v70 = new BigNumber(577);
        _1868_v70 = (_this).f6;
        let _1869_v71;
        _1869_v71 = _module.D2.create_DC7((_this).f6);
        let _source33 = _1869_v71;
        if (_source33.is_DC6) {
          let _1870_v72;
          let _1871_v73;
          let _1872_v74;
          let _out53;
          let _out54;
          let _out55;
          let _outcollector23 = (_this).m8(_module.__default.safeModuloInt((_this).f6, _1868_v70), _1868_v70, globalState);
          _out53 = _outcollector23[0];
          _out54 = _outcollector23[1];
          _out55 = _outcollector23[2];
          _1870_v72 = _out53;
          _1871_v73 = _out54;
          _1872_v74 = _out55;
          let _1873_v75;
          _1873_v75 = new _dafny.CodePoint('j'.codePointAt(0));
          let _1874_v76;
          _1874_v76 = _dafny.Seq.UnicodeFromString("ftxvovkv");
          let _1875_v77;
          _1875_v77 = _module.D0.create_DC0(new BigNumber((_1874_v76).length));
          let _1876_v78;
          _1876_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1875_v77,_1870_v72);
          let _1877_v79;
          _1877_v79 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_1873_v75, globalState),_1876_v78);
          let _1878_v80;
          let _nw282 = Array((new BigNumber(5)).toNumber());
          _nw282[(_dafny.ZERO).toNumber()] = _1868_v70;
          _nw282[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((((_1877_v79).contains(!(p0))) ? ((_1877_v79).get(!(p0))) : (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0(_1868_v70),(_this).f4)))).length));
          _nw282[(new BigNumber(2)).toNumber()] = (_this).f6;
          _nw282[(new BigNumber(3)).toNumber()] = new BigNumber(-352);
          _nw282[(new BigNumber(4)).toNumber()] = (_this).f6;
          _1878_v80 = _nw282;
          let _1879_v81;
          _1879_v81 = _dafny.Set.fromElements(_1878_v80);
          _1870_v72 = (_1879_v81).IsDisjointFrom(_1879_v81);
          _1866_v69 = (_1866_v69).update((_this).f6, (_this).f6);
          let _1880_v82;
          _1880_v82 = _module.D0.create_DC2(_1868_v70);
          _1880_v82 = _1880_v82;
        } else if (_source33.is_DC7) {
          let _1881___mcc_h6 = (_source33).cf6;
          let _1882_cf6 = _1881___mcc_h6;
          let _1883_v83;
          let _init45 = ((_1884_p0) => function (_1885_i5) {
            return _1884_p0;
          })(p0);
          let _nw283 = Array((new BigNumber(27)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw283.length); _i0_45++) {
            _nw283[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1883_v83 = _nw283;
          let _1886_v84;
          _1886_v84 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1887_v85;
          _1887_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1886_v84,_1867_i4);
          let _1888_v87;
          _1888_v87 = _dafny.Seq.of(_1886_v84, _1886_v84);
          let _1889_v88;
          _1889_v88 = _dafny.Seq.of(function () {
            let _coll55 = new _dafny.Map();
            for (const _compr_55 of (_1888_v87).Elements) {
              let _1890_v86 = _compr_55;
              if (_dafny.Seq.contains(_1888_v87, _1890_v86)) {
                _coll55.push([_1890_v86,_1868_v70]);
              }
            }
            return _coll55;
          }());
          let _1891_v89;
          _1891_v89 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1883_v83),_dafny.Seq.Concat(_dafny.Seq.of(_1887_v85, _1887_v85), _dafny.Seq.update(_1889_v88, _module.__default.safeIndex(_1882_cf6, new BigNumber((_1889_v88).length)), _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_module.__default.fm3(_1882_cf6, (_this).f11, globalState), (_this).f11, globalState),(_dafny.ZERO).minus((_this).f6)))));
          let _1892_v90;
          _1892_v90 = _dafny.Seq.of(_1883_v83, _1883_v83, _1883_v83);
          let _1893_v91;
          _1893_v91 = _dafny.Seq.of(_1883_v83, _1883_v83, (_1892_v90)[_module.__default.safeIndex(new BigNumber(-553), new BigNumber((_1892_v90).length))], (_module.D2.create_DC5(_1883_v83)).dtor_cf5, _1883_v83);
          _1891_v89 = (_1891_v89).update(_1892_v90, _dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), ((_1894_v85) => function (_1895_i6) {
            return _1894_v85;
          })(_1887_v85)));
          let _1896_v92;
          _1896_v92 = _dafny.Map.Empty.slice().updateUnsafe(_1888_v87,(_this).f6);
          _1896_v92 = (_1896_v92).update(_dafny.Seq.UnicodeFromString("qpjli"), _1882_cf6);
          let _1897_v93;
          _1897_v93 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          r0 = _module.__default.fm1(new BigNumber(((_1897_v93).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false))).length), (_this).f4, globalState);
          let _1898_v94;
          let _nw284 = Array((new BigNumber(4)).toNumber()).fill([]);
          _1898_v94 = _nw284;
          let _1899_v95;
          let _init46 = function (_1900_i7) {
            return (_1900_i7).minus(new BigNumber(964));
          };
          let _nw285 = Array((_dafny.ONE).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw285.length); _i0_46++) {
            _nw285[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1899_v95 = _nw285;
          let _index280 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1898_v94).length));
          (_1898_v94)[_index280] = _1899_v95;
          let _index281 = _module.__default.safeIndex(new BigNumber(906), new BigNumber((_1898_v94).length));
          (_1898_v94)[_index281] = _1899_v95;
        } else {
          let _1901___mcc_h7 = (_source33).cf5;
          let _1902_cf5 = _1901___mcc_h7;
          let _1903_v96;
          let _nw286 = new _module.C0();
          _nw286.__ctor((_this).f6);
          _1903_v96 = _nw286;
          let _1904_v97;
          _1904_v97 = _dafny.MultiSet.fromElements((_1903_v96).f12, (_1903_v96).f12, (_this).f6, _1868_v70);
          let _1905_v98;
          _1905_v98 = _dafny.Map.Empty.slice().updateUnsafe(_1904_v97,_dafny.Seq.UnicodeFromString("synxuae"));
          let _1906_v99;
          let _1907_v100;
          let _out56;
          let _out57;
          let _outcollector24 = _module.__default.m0(_1905_v98, (_this).f11, _dafny.MultiSet.fromElements(_1868_v70), (_this).f11, globalState);
          _out56 = _outcollector24[0];
          _out57 = _outcollector24[1];
          _1906_v99 = _out56;
          _1907_v100 = _out57;
          let _1908_v101;
          _1908_v101 = _dafny.Seq.UnicodeFromString("adljl");
          let _1909_v102;
          _1909_v102 = _dafny.Set.fromElements((((_1780_v0)[_module.__default.safeIndex(_1868_v70, new BigNumber((_1780_v0).length))]) ? (_1908_v101) : (_1908_v101)));
          _1909_v102 = _1909_v102;
          let _1910_v103;
          _1910_v103 = _dafny.Map.Empty.slice().updateUnsafe(!(!(p0)),(_dafny.ZERO).minus(_1868_v70));
          let _1911_v104;
          _1911_v104 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1906_v99,(_1903_v96).f12), _1910_v103, _1910_v103, _1910_v103);
          let _1912_v105;
          let _nw287 = new _module.C0();
          _nw287.__ctor(new BigNumber((_1911_v104).length));
          _1912_v105 = _nw287;
        }
        let _1913_v106;
        _1913_v106 = _dafny.Seq.UnicodeFromString("tasx");
        let _1914_v107;
        let _nw288 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _1914_v107 = _nw288;
        let _index282 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1914_v107).length));
        (_1914_v107)[_index282] = new BigNumber(415);
        let _index283 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1914_v107).length));
        let _rhs259 = _dafny.Seq.UnicodeFromString("xohxe");
        let _rhs260 = new BigNumber((_1913_v106).length);
        let _lhs163 = _1914_v107;
        let _lhs164 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1914_v107).length));
        _1913_v106 = _rhs259;
        _lhs163[_lhs164] = _rhs260;
        let _index284 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_1914_v107).length));
        (_1914_v107)[_index284] = _1868_v70;
      }
      let _1915_v108;
      _1915_v108 = new _dafny.CodePoint('h'.codePointAt(0));
      if (_module.__default.fm0(_1915_v108, globalState)) {
        let _1916_v109;
        _1916_v109 = new BigNumber(639);
        let _1917_v110;
        _1917_v110 = _dafny.MultiSet.fromElements((_this).f4);
        _1916_v109 = ((_this).f6).multipliedBy(new BigNumber((_dafny.Seq.of(_1916_v109, new BigNumber((_1917_v110).cardinality()), (_this).f6)).length));
        let _1918_v111;
        _1918_v111 = false;
        let _1919_v112;
        _1919_v112 = _dafny.Seq.UnicodeFromString("s");
        let _1920_v113;
        _1920_v113 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1919_v112).length),_1918_v111);
        _1918_v111 = (((_1920_v113).contains(_1916_v109)) ? ((_1920_v113).get(_1916_v109)) : ((_this).f11));
        let _1921_v114;
        _1921_v114 = _dafny.Map.Empty.slice().updateUnsafe(_1780_v0,_1918_v111);
        _1921_v114 = (_1921_v114).update(_dafny.Seq.of(p0), p0);
        let _1922_v115;
        _1922_v115 = _dafny.MultiSet.fromElements(_module.__default.fm19(p0, _1916_v109, globalState));
        let _1923_v116;
        _1923_v116 = _module.D1.create_DC3(_1918_v111);
        let _1924_v117;
        _1924_v117 = _dafny.Seq.of(_1780_v0, _1780_v0);
        let _1925_v118;
        let _nw289 = Array((new BigNumber(22)).toNumber());
        _nw289[(_dafny.ZERO).toNumber()] = _1922_v115;
        _nw289[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(_1781_v1, _1780_v0, _dafny.Seq.of((_this).f11), _1780_v0);
        _nw289[(new BigNumber(2)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(3)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(_1781_v1);
        _nw289[(new BigNumber(5)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(6)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements(_1781_v1, _dafny.Seq.of((_this).f11, true), _1780_v0, _1780_v0);
        _nw289[(new BigNumber(8)).toNumber()] = (((_1923_v116).dtor_cf2) ? (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(146)), ((_1926_v0) => function (_1927_i8) {
          return _1926_v0;
        })(_1780_v0)))) : (_module.__default.fm24(_1916_v109, _1915_v108, _1916_v109, globalState)));
        _nw289[(new BigNumber(9)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(10)).toNumber()] = (_1922_v115).Intersect(_dafny.MultiSet.FromArray(_1924_v117));
        _nw289[(new BigNumber(11)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(12)).toNumber()] = (_1922_v115).Intersect(_1922_v115);
        _nw289[(new BigNumber(13)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(14)).toNumber()] = (_1922_v115).Intersect(_1922_v115);
        _nw289[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f4))).Intersect(_1922_v115);
        _nw289[(new BigNumber(16)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(17)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(18)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(19)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(20)).toNumber()] = _1922_v115;
        _nw289[(new BigNumber(21)).toNumber()] = _1922_v115;
        _1925_v118 = _nw289;
        let _1928_v119;
        _1928_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1918_v111,_dafny.Seq.of(_1781_v1));
        let _1929_v120;
        _1929_v120 = _module.D6.create_DC15(_1922_v115);
        let _index285 = _module.__default.safeIndex(new BigNumber(270), new BigNumber((_1925_v118).length));
        (_1925_v118)[_index285] = (_dafny.MultiSet.FromArray((((_1928_v119).contains(p0)) ? ((_1928_v119).get(p0)) : (_1924_v117)))).Union((_1929_v120).dtor_cf19);
        let _1930_v121;
        _1930_v121 = _dafny.Seq.of(_1916_v109, (_dafny.ZERO).minus(_1916_v109));
        let _1931_v122;
        _1931_v122 = _dafny.MultiSet.fromElements(_1930_v121);
        let _1932_v123;
        let _nw290 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
        _1932_v123 = _nw290;
        let _index286 = _module.__default.safeIndex(new BigNumber(270), new BigNumber((_1925_v118).length));
        let _rhs261 = (_1922_v115).update(_dafny.Seq.of(false), _module.__default.abs((_this).f6));
        let _rhs262 = _1931_v122;
        let _rhs263 = _1932_v123;
        let _rhs264 = (_this).f11;
        let _lhs165 = _1925_v118;
        let _lhs166 = _module.__default.safeIndex(new BigNumber(270), new BigNumber((_1925_v118).length));
        _lhs165[_lhs166] = _rhs261;
        _1931_v122 = _rhs262;
        _1932_v123 = _rhs263;
        _1918_v111 = _rhs264;
        _1916_v109 = _module.__default.safeDivisionInt(new BigNumber(313), _1916_v109);
      } else {
        let _1933_v124;
        _1933_v124 = _module.D3.create_DC10((_this).f6);
        _1933_v124 = _module.D3.create_DC10((_this).f6);
        let _1934_v125;
        _1934_v125 = new BigNumber(216);
        _1934_v125 = ((p0) ? ((((_1866_v69).contains(new BigNumber(-103))) ? ((_1866_v69).get(new BigNumber(-103))) : (_1934_v125))) : (new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_1780_v0, _module.__default.safeIndex(_1934_v125, new BigNumber((_1780_v0).length)), _module.__default.fm0(_1915_v108, globalState)), _module.__default.safeIndex(new BigNumber(948), new BigNumber((_dafny.Seq.update(_1780_v0, _module.__default.safeIndex(_1934_v125, new BigNumber((_1780_v0).length)), _module.__default.fm0(_1915_v108, globalState))).length)), true)).length)));
        let _source34 = _module.D3.create_DC8(_1780_v0);
        if (_source34.is_DC9) {
          let _1935___mcc_h8 = (_source34).cf8;
          let _1936___mcc_h9 = (_source34).cf9;
          let _1937___mcc_h10 = (_source34).cf10;
          let _1938_cf10 = _1937___mcc_h10;
          let _1939_cf9 = _1936___mcc_h9;
          let _1940_cf8 = _1935___mcc_h8;
          let _1941_v126;
          _1941_v126 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(676)), ((_1942_v108) => function (_1943_i9) {
            return _1942_v108;
          })(_1915_v108))).length), new BigNumber((_dafny.Set.fromElements((_this).f11, !((_this).f4), _1940_cf8)).length), new BigNumber((_module.__default.fm25(_1939_cf9, (_this).f6, (_this).f4, (_this).f11, globalState)).length), _1934_v125);
          let _1944_v127;
          _1944_v127 = _dafny.Seq.of(_1941_v126);
          let _1945_v128;
          _1945_v128 = _dafny.Seq.of(_1939_cf9, (_this).f6, _1934_v125);
          _1941_v126 = (_1944_v127)[_module.__default.safeIndex(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1934_v125,_1945_v128)).update((_this).f6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), ((_1946_cf9) => function (_1947_i10) {
            return _1946_cf9;
          })(_1939_cf9)))).length), new BigNumber((_1944_v127).length))];
          let _1948_v129;
          _1948_v129 = _dafny.Set.fromElements(_1934_v125);
          _1939_cf9 = (_this).fm9((_this).f11, _1940_cf8, new BigNumber(((_1948_v129)).length), globalState);
          let _1949_v130;
          let _init47 = function (_1950_i11) {
            return (_1950_i11).multipliedBy((_this).f6);
          };
          let _nw291 = Array((new BigNumber(8)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw291.length); _i0_47++) {
            _nw291[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1949_v130 = _nw291;
          let _1951_v131;
          let _init48 = function (_1952_i12) {
            return _module.D0.create_DC0((_this).f6);
          };
          let _nw292 = Array((new BigNumber(10)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw292.length); _i0_48++) {
            _nw292[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1951_v131 = _nw292;
          let _1953_v132;
          _1953_v132 = _module.D0.create_DC0(_1939_cf9);
          let _index287 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_1951_v131).length));
          (_1951_v131)[_index287] = _1953_v132;
          let _1954_v133;
          let _nw293 = Array((new BigNumber(27)).toNumber()).fill(false);
          _1954_v133 = _nw293;
          let _index288 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_1951_v131).length));
          let _rhs265 = _1933_v124;
          let _rhs266 = _1949_v130;
          let _rhs267 = (_dafny.ZERO).minus((_dafny.ZERO).minus(((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(533), _1939_cf9))).plus(((_1938_cf10) ? (_1939_cf9) : (new BigNumber(-614))))));
          let _rhs268 = _1953_v132;
          let _rhs269 = _1954_v133;
          let _lhs167 = _1951_v131;
          let _lhs168 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_1951_v131).length));
          _1933_v124 = _rhs265;
          _1949_v130 = _rhs266;
          _1939_cf9 = _rhs267;
          _lhs167[_lhs168] = _rhs268;
          _1954_v133 = _rhs269;
          _1934_v125 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1939_cf9));
        } else if (_source34.is_DC10) {
          let _1955___mcc_h11 = (_source34).cf11;
          let _1956_cf11 = _1955___mcc_h11;
          let _1957_v134;
          let _init49 = function (_1958_i13) {
            return (_1958_i13).plus((_this).f6);
          };
          let _nw294 = Array((new BigNumber(19)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw294.length); _i0_49++) {
            _nw294[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1957_v134 = _nw294;
          let _1959_v135;
          _1959_v135 = _dafny.Map.Empty.slice().updateUnsafe(_1957_v134,p0);
          _1959_v135 = (((_1959_v135).update(_1957_v134, (_this).f4)).update(_1957_v134, (_this).f4));
          let _1960_v136;
          _1960_v136 = true;
          let _rhs270 = (_this).f6;
          let _rhs271 = ((_this).f6).plus(_1956_cf11);
          let _rhs272 = !_dafny.areEqual((_module.D3.create_DC8(_1780_v0)).dtor_cf7, _dafny.Seq.of(p0, p0, (_this).f11, _1960_v136));
          _1956_cf11 = _rhs270;
          _1956_cf11 = _rhs271;
          _1960_v136 = _rhs272;
          let _1961_v137;
          _1961_v137 = _dafny.Seq.UnicodeFromString("fjydewlp");
          let _1962_v138;
          let _nw295 = new _module.C3();
          _nw295.__ctor(_this.f5, new BigNumber((_1961_v137).length), false);
          _1962_v138 = _nw295;
          let _1963_v139;
          _1963_v139 = _dafny.Map.Empty.slice().updateUnsafe(_1960_v136,_1962_v138);
          let _1964_v140;
          _1964_v140 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1962_v138);
          _1963_v139 = (_1963_v139).update(((p0) ? ((_this).f11) : ((_1962_v138).f4)), (((_1964_v140).contains((_1962_v138).f6)) ? ((_1964_v140).get((_1962_v138).f6)) : (_1962_v138)));
          _1934_v125 = _1934_v125;
        } else if (_source34.is_DC8) {
          let _1965___mcc_h12 = (_source34).cf7;
          let _1966_cf7 = _1965___mcc_h12;
          let _1967_v141;
          let _nw296 = Array((new BigNumber(11)).toNumber()).fill(false);
          _1967_v141 = _nw296;
          let _1968_v142;
          _1968_v142 = _dafny.Seq.of(_1967_v141);
          let _1969_v143;
          _1969_v143 = _dafny.Seq.UnicodeFromString("sxv");
          let _1970_v144;
          _1970_v144 = _module.D4.create_DC13((_dafny.ZERO).minus(_1934_v125), !_dafny.Seq.contains(_1968_v142, _1967_v141), _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_1969_v143).length)), (_this).f6), _1780_v0);
          let _1971_v145;
          _1971_v145 = false;
          let _1972_v146;
          _1972_v146 = _dafny.Seq.of((_this).f6, _1934_v125);
          let _rhs273 = ((_this).f6).plus((_this).f6);
          let _rhs274 = _module.__default.fm63(globalState);
          let _rhs275 = (_1972_v146)[_module.__default.safeIndex(_1934_v125, new BigNumber((_1972_v146).length))];
          let _rhs276 = (_this).f4;
          _1934_v125 = _rhs273;
          _1970_v144 = _rhs274;
          _1934_v125 = _rhs275;
          _1971_v145 = _rhs276;
          let _1973_v147;
          _1973_v147 = _dafny.MultiSet.fromElements((_this).f11);
          _1971_v145 = (!((_this).f11)) === ((_1973_v147).IsProperSubsetOf(_1973_v147));
          let _1974_v148;
          _1974_v148 = _dafny.Map.Empty.slice().updateUnsafe(_1915_v108,_dafny.Seq.Concat(_1969_v143, _1969_v143));
          _1974_v148 = _1974_v148;
          let _1975_v150;
          let _init50 = ((_1976_v125) => function (_1977_i14) {
            return _dafny.MultiSet.fromElements(function () {
              let _coll56 = new _dafny.Set();
              for (const _compr_56 of _dafny.IntegerRange(new BigNumber(528), new BigNumber(26))) {
                let _1978_v149 = _compr_56;
                if (((new BigNumber(528)).isLessThanOrEqualTo(_1978_v149)) && ((_1978_v149).isLessThan(new BigNumber(26)))) {
                  _coll56.add(_module.__default.safeDivisionInt(_1978_v149, _1976_v125));
                }
              }
              return _coll56;
            }(), _dafny.Set.fromElements((_this).f6));
          })(_1934_v125);
          let _nw297 = Array((new BigNumber(11)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw297.length); _i0_50++) {
            _nw297[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _1975_v150 = _nw297;
          let _1979_v151;
          _1979_v151 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-561),(_module.D3.create_DC9(_1971_v145, (_this).fm15(globalState), (_this).f11)).dtor_cf8);
          let _1980_v152;
          _1980_v152 = _dafny.Set.fromElements(new BigNumber((_1979_v151).length));
          let _1981_v153;
          _1981_v153 = _dafny.MultiSet.fromElements(_1980_v152, _1980_v152);
          let _index289 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_1975_v150).length));
          (_1975_v150)[_index289] = _1981_v153;
          let _1982_v154;
          _1982_v154 = _dafny.Seq.of(_1980_v152, _1980_v152, _1980_v152);
          let _index290 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_1975_v150).length));
          (_1975_v150)[_index290] = (_1981_v153).update(((_1982_v154)[_module.__default.safeIndex(new BigNumber((_1969_v143).length), new BigNumber((_1982_v154).length))]).Difference(_1980_v152), _module.__default.abs((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_1972_v146).length), _1934_v125))));
        } else {
          let _1983___mcc_h13 = (_source34).cf12;
          let _1984_cf12 = _1983___mcc_h13;
          _1934_v125 = (_this).fm15(globalState);
          let _1985_v155;
          _1985_v155 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,new BigNumber(725));
          let _1986_v156;
          _1986_v156 = _dafny.MultiSet.fromElements(_1985_v155, _module.__default.fm44(globalState), _1985_v155, _1985_v155, _1985_v155);
          _1986_v156 = (_1986_v156).Union((_1986_v156).Union((_dafny.MultiSet.fromElements(_1985_v155)).update(_1985_v155, _module.__default.abs(_module.__default.fm3(_1934_v125, p0, globalState)))));
          _1915_v108 = _1915_v108;
          _1985_v155 = _1985_v155;
        }
        _1934_v125 = ((p0) ? ((_dafny.ZERO).minus((_this).f6)) : (((_this).f6).plus((_this).f6)));
        let _1987_v157;
        _1987_v157 = true;
        _1987_v157 = p0;
      }
      let _1988_v158;
      _1988_v158 = _dafny.Seq.UnicodeFromString("ehhyianjb");
      let _1989_v159;
      let _nw298 = new _module.C2();
      _nw298.__ctor(_1988_v158, p0);
      _1989_v159 = _nw298;
      let _1990_v160;
      _1990_v160 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(692),_1989_v159);
      r0 = _module.__default.fm1(_module.__default.fm3((_this).f6, p0, globalState), ((_this).fm7(globalState)).isLessThanOrEqualTo(new BigNumber((_1990_v160).length)), globalState);
      let _1991_v161;
      _1991_v161 = _module.D1.create_DC3((_this).f4);
      let _1992_v162;
      _1992_v162 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,p0);
      let _1993_v163;
      let _nw299 = Array((new BigNumber(29)).toNumber());
      _nw299[(_dafny.ZERO).toNumber()] = _1991_v161;
      _nw299[(_dafny.ONE).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(2)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(3)).toNumber()] = _module.__default.fm55((((_1992_v162).contains(p0)) ? ((_1992_v162).get(p0)) : ((_this).f11)), (_this).f4, (_this).f6, globalState);
      _nw299[(new BigNumber(4)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(5)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(6)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(7)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(8)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(9)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(10)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(11)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(12)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(13)).toNumber()] = _module.D1.create_DC3(_module.__default.fm0(_1915_v108, globalState));
      _nw299[(new BigNumber(14)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(15)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(16)).toNumber()] = _module.D1.create_DC3((_this).f11);
      _nw299[(new BigNumber(17)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(18)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(19)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(20)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(21)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(22)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(23)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(24)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(25)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(26)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(27)).toNumber()] = _1991_v161;
      _nw299[(new BigNumber(28)).toNumber()] = _1991_v161;
      _1993_v163 = _nw299;
      let _1994_v164;
      _1994_v164 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1993_v163);
      let _1995_v165;
      _1995_v165 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6);
      let _1996_v166;
      _1996_v166 = _dafny.Map.Empty.slice().updateUnsafe(_1995_v165,(_this).f4);
      r1 = (((_1994_v164).contains((((_1996_v166).contains(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6))) ? ((_1996_v166).get(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6))) : ((_this).f4)))) ? ((_1994_v164).get((((_1996_v166).contains(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6))) ? ((_1996_v166).get(_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6))) : ((_this).f4)))) : (_1993_v163));
      return [r0, r1];
    }
    m7(globalState) {
      let _this = this;
      let _1997_v0;
      _1997_v0 = true;
      _1997_v0 = (_this).f11;
      let _1998_v1;
      _1998_v1 = new BigNumber(843);
      let _1999_v2;
      _1999_v2 = _dafny.Seq.of((_this).f6);
      let _2000_v3;
      _2000_v3 = _dafny.Seq.UnicodeFromString("gnnf");
      let _2001_v4;
      _2001_v4 = _dafny.Seq.of(true);
      let _2002_v5;
      _2002_v5 = _dafny.Seq.of(_2001_v4);
      let _2003_v6;
      _2003_v6 = _module.D12.create_DC27(_2002_v5);
      let _2004_v7;
      _2004_v7 = _dafny.Seq.of(_1999_v2, _dafny.Seq.Concat(_module.__default.fm39(_2000_v3, _2003_v6, (_this).f11, globalState), _dafny.Seq.of(_1998_v1, (_this).f6)));
      _1998_v1 = new BigNumber(((_2004_v7)[_module.__default.safeIndex((((_this).f4) ? ((_this).f6) : ((_this).f6)), new BigNumber((_2004_v7).length))]).length);
      let _2005_v8;
      let _nw300 = Array((new BigNumber(14)).toNumber()).fill(false);
      _2005_v8 = _nw300;
      _2005_v8 = _2005_v8;
      let _2006_v9;
      let _nw301 = new _module.C4();
      _nw301.__ctor(_this.f5, _1998_v1, (_this).f4);
      _2006_v9 = _nw301;
      let _rhs277 = _module.__default.safeDivisionInt(_1998_v1, (_this).f6);
      let _rhs278 = !(new BigNumber((((_1997_v0) ? (_2001_v4) : (_dafny.Seq.update(_2001_v4, _module.__default.safeIndex((_this).f6, new BigNumber((_2001_v4).length)), (_this).f11)))).length)).isEqualTo(_1998_v1);
      _1998_v1 = _rhs277;
      _1997_v0 = _rhs278;
      let _2007_v10;
      _2007_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1998_v1);
      let _2008_v11;
      let _nw302 = new _module.C1();
      _nw302.__ctor(_1998_v1, ((((_2007_v10).contains(_1998_v1)) ? ((_2007_v10).get(_1998_v1)) : (_1998_v1))).multipliedBy(new BigNumber((_2000_v3).length)), _this.f5, _1998_v1, !(_1997_v0));
      _2008_v11 = _nw302;
      return;
    }
    m8(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _module.D1.Default();
      let _2009_v0;
      _2009_v0 = new _dafny.CodePoint('a'.codePointAt(0));
      let _2010_v1;
      let _nw303 = Array((new BigNumber(15)).toNumber()).fill(false);
      _2010_v1 = _nw303;
      let _2011_v2;
      _2011_v2 = _dafny.Map.Empty.slice().updateUnsafe(_2009_v0,_2010_v1);
      let _2012_v3;
      _2012_v3 = _module.D19.create_DC39(_module.D19.create_DC37(_2011_v2));
      let _2013_v4;
      _2013_v4 = _module.D19.create_DC37((_2011_v2).update(_2009_v0, _2010_v1));
      _2012_v3 = _module.D19.create_DC39(_2013_v4);
      if ((_this).f11) {
        let _2014_v5;
        _2014_v5 = new BigNumber(594);
        _2014_v5 = (p1).plus(_module.__default.fm3((_this).f6, (_this).f11, globalState));
        let _2015_v6;
        _2015_v6 = _dafny.Set.fromElements(!((_this).f4));
        _2015_v6 = _2015_v6;
        r1 = ((_this).f11) === ((((_this).f4) ? ((_this).f11) : ((_this).f11)));
        _2014_v5 = (_2014_v5).plus((_this).f6);
        r0 = (_this).f4;
      } else {
        let _arr9 = _this.f5;
        let _index291 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_this.f5).length));
        _arr9[_index291] = _2010_v1;
        let _arr10 = _this.f5;
        let _index292 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_this.f5).length));
        _arr10[_index292] = _2010_v1;
        let _2016_v7;
        _2016_v7 = _dafny.Seq.of((_this).f11, (_this).f4, true, (_this).f4);
        let _2017_v8;
        _2017_v8 = _dafny.Seq.of(_dafny.Seq.Concat(_2016_v7, _2016_v7), _dafny.Seq.Concat(_2016_v7, _2016_v7), _dafny.Seq.Concat(_dafny.Seq.of((_this).f11), _2016_v7));
        _2017_v8 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), ((_2018_v7) => function (_2019_i0) {
          return _2018_v7;
        })(_2016_v7));
        let _2020_v9;
        _2020_v9 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("brp"));
        let _2021_v10;
        _2021_v10 = _dafny.MultiSet.fromElements(new BigNumber(255), (((_2020_v9).contains(_dafny.Seq.UnicodeFromString("cufybjlks"))) ? ((_2020_v9).get(_dafny.Seq.UnicodeFromString("cufybjlks"))) : (p0)));
        let _2022_v11;
        _2022_v11 = _module.D0.create_DC0((_this).f6);
        let _2023_v12;
        _2023_v12 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2009_v0);
        let _2024_v13;
        _2024_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4);
        let _2025_v14;
        let _2026_v15;
        let _out58;
        let _out59;
        let _outcollector25 = _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe(_2021_v10,_dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), ((_2027_v0) => function (_2028_i1) {
          return _2027_v0;
        })(_2009_v0))), (_this).f4, ((_this).fm6(_2022_v11, _2022_v11, _2023_v12, p1, globalState)).Union(_2021_v10), !(false) || ((((_2024_v13).contains((_this).f11)) ? ((_2024_v13).get((_this).f11)) : ((_this).f4))), globalState);
        _out58 = _outcollector25[0];
        _out59 = _outcollector25[1];
        _2025_v14 = _out58;
        _2026_v15 = _out59;
        let _2029_v16;
        _2029_v16 = _dafny.Seq.UnicodeFromString("wwdadof");
        let _2030_v17;
        _2030_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2026_v15,_2029_v16);
        let _2031_v18;
        _2031_v18 = _module.D2.create_DC6();
        _2030_v17 = (_2030_v17).update(_module.__default.safeDivisionInt(p1, new BigNumber(442)), _dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm16(_2025_v14, (_this).f11, false, _2031_v18, globalState), _module.__default.safeIndex(p1, new BigNumber((_module.__default.fm16(_2025_v14, (_this).f11, false, _2031_v18, globalState)).length)), _2009_v0), _2029_v16));
        let _2032_v19;
        let _nw304 = Array((new BigNumber(2)).toNumber());
        _nw304[(_dafny.ZERO).toNumber()] = (_this.f5)[_module.__default.safeIndex(new BigNumber(951), new BigNumber((_this.f5).length))];
        _nw304[(_dafny.ONE).toNumber()] = _2010_v1;
        _2032_v19 = _nw304;
        let _2033_v20;
        _2033_v20 = _2032_v19;
        let _2034_v21;
        let _nw305 = new _module.C5();
        _nw305.__ctor((_this).f4, (_dafny.ZERO).minus(p0), (_2033_v20), (new BigNumber((_2029_v16).length)).multipliedBy(p0), _2025_v14);
        _2034_v21 = _nw305;
      }
      let _2035_v22;
      _2035_v22 = new BigNumber(861);
      _2035_v22 = (_2035_v22).multipliedBy((p1).minus(p1));
      if ((_this).f11) {
        let _2036_v23;
        _2036_v23 = _dafny.Seq.UnicodeFromString("senka");
        let _2037_v24;
        _2037_v24 = _dafny.Map.Empty.slice().updateUnsafe(_this.f5,_dafny.areEqual(_2036_v23, _2036_v23));
        _2037_v24 = (_2037_v24).update(_this.f5, (_this).f4);
        let _2038_v25;
        _2038_v25 = _dafny.Set.fromElements((_this).f4, !((_this).f11), (_this).f11);
        _2038_v25 = (_2038_v25).Intersect(_2038_v25);
        r0 = (_this).f4;
        (_this).m7(globalState);
        _2009_v0 = _2009_v0;
      } else {
        let _2039_v26;
        _2039_v26 = _module.D1.create_DC3(!(true));
        r2 = _2039_v26;
        let _2040_v27;
        _2040_v27 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f4),(_this).f4);
        let _2041_v28;
        _2041_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,(((_2040_v27).contains((_this).f4)) ? ((_2040_v27).get((_this).f4)) : (!((_this).f11))));
        _2041_v28 = (_2041_v28).update(_2035_v22, (_this).f11);
        r0 = (_this).f11;
        let _index293 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_2010_v1).length));
        (_2010_v1)[_index293] = (_this).f4;
        let _index294 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_2010_v1).length));
        (_2010_v1)[_index294] = (_this).f11;
        let _2042_v29;
        _2042_v29 = _dafny.Seq.UnicodeFromString("pgj");
        let _2043_v30;
        let _nw306 = new _module.C2();
        _nw306.__ctor(_2042_v29, (_2010_v1)[_module.__default.safeIndex(new BigNumber(250), new BigNumber((_2010_v1).length))]);
        _2043_v30 = _nw306;
      }
      let _2044_v31;
      _2044_v31 = _dafny.Seq.UnicodeFromString("a");
      let _2045_v32;
      let _nw307 = new _module.C2();
      _nw307.__ctor(_2044_v31, (_this).f11);
      _2045_v32 = _nw307;
      _2045_v32 = _2045_v32;
      let _2046_v33;
      _2046_v33 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f11) ? ((_this).f11) : ((_this).f4)),(_this).f11);
      _2046_v33 = _2046_v33;
      r0 = (_this).f4;
      r1 = (((_this).f6).plus(_2035_v22)).isLessThanOrEqualTo((_this).f6);
      let _2047_v34;
      _2047_v34 = _module.D1.create_DC3((_this).f11);
      r2 = _2047_v34;
      return [r0, r1, r2];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f5 = [];
      this._f6 = _dafny.ZERO;
      this._f4 = false;
      this._f10 = [];
      this._f9 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    set f5(value) {
      let _this = this;
      _this._f5 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f9, f10, f5, f6, f4) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f4 = f4;
      return;
    }
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((((_this).f4) ? ((_this).f4) : ((_this).f4))) {
        return _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(799)), function (_2048_i0) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        }), _dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0))));
      } else {
        return ((_this).f6).isEqualTo((_dafny.ZERO).minus((_this).f6));
      }
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f6;
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (((false) ? (_dafny.MultiSet.fromElements((_this).f6)) : (_dafny.MultiSet.fromElements((_this).f6)))).Difference((_dafny.MultiSet.fromElements((_this).f6)).Difference(_dafny.MultiSet.fromElements(new BigNumber(472))));
    };
    fm7(globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f6, (_this).f6))).minus(new BigNumber(677));
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of((_this).f4), _dafny.Seq.of(!((_this).f4), (_this).f4, (_this).f4));
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_this).f6).isEqualTo(new BigNumber(664));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _2049_v0;
      let _nw308 = new _module.C1();
      _nw308.__ctor((_this).f6, p0, _this.f5, p0, (_this).f4);
      _2049_v0 = _nw308;
      let _2050_v1;
      _2050_v1 = _dafny.Seq.of((_this).f4, (_this).f4, true, (_this).f4, !((_this).f4));
      let _source35 = _module.D4.create_DC13(new BigNumber(194), !((_this).f4), ((_this).f6).minus(_2049_v0.f16), _2050_v1);
      if (_source35.is_DC13) {
        let _2051___mcc_h0 = (_source35).cf14;
        let _2052___mcc_h1 = (_source35).cf15;
        let _2053___mcc_h2 = (_source35).cf16;
        let _2054___mcc_h3 = (_source35).cf17;
        let _2055_cf17 = _2054___mcc_h3;
        let _2056_cf16 = _2053___mcc_h2;
        let _2057_cf15 = _2052___mcc_h1;
        let _2058_cf14 = _2051___mcc_h0;
        let _2059_v2;
        _2059_v2 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_module.__default.fm64(p0, (_this).f6, _2057_cf15, globalState)).length)));
        let _2060_v3;
        _2060_v3 = _module.D9.create_DC20(_2057_cf15, new _dafny.CodePoint('n'.codePointAt(0)), _2049_v0.f16);
        let _2061_v4;
        let _nw309 = new _module.C6();
        _nw309.__ctor((new BigNumber((_2059_v2).length)).isLessThan(new BigNumber(272)), (_this).f4, _this.f5, (_2060_v3).dtor_cf28);
        _2061_v4 = _nw309;
        _2057_cf15 = true;
        _2057_cf15 = !(!((_2061_v4).f11));
        (_2049_v0).f16 = ((_2049_v0).f15).minus(new BigNumber(-913));
      } else {
        let _2062___mcc_h4 = (_source35).cf13;
        let _2063_cf13 = _2062___mcc_h4;
        let _2064_v5;
        let _nw310 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _2064_v5 = _nw310;
        let _index295 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_2064_v5).length));
        (_2064_v5)[_index295] = new BigNumber((_2050_v1).length);
        let _index296 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_2064_v5).length));
        (_2064_v5)[_index296] = (_2049_v0).f15;
        let _2065_v6;
        let _init51 = ((_2066_v0, _2067_p0) => function (_2068_i0) {
          return (_dafny.MultiSet.fromElements(_2066_v0.f16)).Union(_dafny.MultiSet.fromElements(_2067_p0));
        })(_2049_v0, p0);
        let _nw311 = Array((new BigNumber(23)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw311.length); _i0_51++) {
          _nw311[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _2065_v6 = _nw311;
        let _2069_v7;
        _2069_v7 = _dafny.MultiSet.fromElements(p0);
        let _index297 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_2065_v6).length));
        (_2065_v6)[_index297] = (((_this).f4) ? (_2069_v7) : (_2069_v7));
        let _2070_v8;
        _2070_v8 = true;
        let _2071_v9;
        _2071_v9 = _dafny.Map.Empty.slice().updateUnsafe((_2049_v0).f15,true);
        let _2072_v10;
        let _nw312 = new _module.C3();
        _nw312.__ctor(_this.f5, (_2049_v0).f15, (((_this).fm14(_2071_v9, true, _2049_v0.f16, false, globalState)) ? ((_this).f4) : (_2070_v8)));
        _2072_v10 = _nw312;
        let _2073_v11;
        _2073_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2070_v8,_2070_v8);
        let _2074_v12;
        _2074_v12 = _dafny.MultiSet.fromElements(_2070_v8, (((_2073_v11).contains((_this).f4)) ? ((_2073_v11).get((_this).f4)) : (_2070_v8)));
        let _2075_v13;
        _2075_v13 = _dafny.Seq.UnicodeFromString("sdatbjvc");
        let _2076_v14;
        let _nw313 = new _module.C2();
        _nw313.__ctor(_dafny.Seq.UnicodeFromString("xyw"), (_this).f4);
        _2076_v14 = _nw313;
        let _2077_v15;
        _2077_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2049_v0.f16,_2076_v14);
        let _2078_v16;
        _2078_v16 = _module.D16.create_DC33(false, (_this).f6, _2074_v12, new BigNumber((_2075_v13).length), _module.__default.fm23(new BigNumber((_2077_v15).length), _2049_v0.f16, globalState));
        let _2079_v17;
        _2079_v17 = _module.D9.create_DC20((_2078_v16).dtor_cf52, _2063_cf13, _2049_v0.f16);
        let _pat_let_tv39 = _2073_v11;
        let _2080_v18;
        let _nw314 = Array((new BigNumber(6)).toNumber());
        _nw314[(_dafny.ZERO).toNumber()] = _2063_cf13;
        _nw314[(_dafny.ONE).toNumber()] = _2063_cf13;
        _nw314[(new BigNumber(2)).toNumber()] = (function (_pat_let48_0) {
          return function (_2081_dt__update__tmp_h0) {
            return function (_pat_let49_0) {
              return function (_2082_dt__update_hcf26_h0) {
                return function (_pat_let50_0) {
                  return function (_2083_dt__update_hcf28_h0) {
                    return _module.D9.create_DC20(_2082_dt__update_hcf26_h0, (_2081_dt__update__tmp_h0).dtor_cf27, _2083_dt__update_hcf28_h0);
                  }(_pat_let50_0);
                }(new BigNumber((_pat_let_tv39).length));
              }(_pat_let49_0);
            }((_this).f4);
          }(_pat_let48_0);
        }(_2079_v17)).dtor_cf27;
        _nw314[(new BigNumber(3)).toNumber()] = (((_this).f4) ? (_2063_cf13) : (_2063_cf13));
        _nw314[(new BigNumber(4)).toNumber()] = _2063_cf13;
        _nw314[(new BigNumber(5)).toNumber()] = _2063_cf13;
        _2080_v18 = _nw314;
        let _index298 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_2065_v6).length));
        let _rhs279 = _2069_v7;
        let _rhs280 = (((_this).f4) || (true)) && (!((_2049_v0).f15).isEqualTo(_2049_v0.f16));
        let _rhs281 = _2072_v10;
        let _rhs282 = _2080_v18;
        let _lhs169 = _2065_v6;
        let _lhs170 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_2065_v6).length));
        _lhs169[_lhs170] = _rhs279;
        _2070_v8 = _rhs280;
        _2072_v10 = _rhs281;
        _2080_v18 = _rhs282;
        (_2049_v0).f16 = _2049_v0.f16;
        let _index299 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_2064_v5).length));
        let _rhs283 = (_module.__default.safeDivisionInt(new BigNumber(989), (_2049_v0).f15)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cfu")).length));
        let _rhs284 = new BigNumber(196);
        let _rhs285 = (_dafny.ZERO).minus(_2049_v0.f16);
        let _lhs171 = _2049_v0;
        let _lhs172 = _2049_v0;
        let _lhs173 = _2064_v5;
        let _lhs174 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_2064_v5).length));
        _lhs171.f16 = _rhs283;
        _lhs172.f16 = _rhs284;
        _lhs173[_lhs174] = _rhs285;
      }
      let _2084_v19;
      _2084_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2049_v0.f16);
      let _2085_v20;
      _2085_v20 = _dafny.Map.Empty.slice().updateUnsafe(true,_2084_v19);
      let _2086_v21;
      _2086_v21 = _dafny.Seq.of(_2084_v19, (((_2085_v20).contains(_module.__default.fm0(new _dafny.CodePoint('a'.codePointAt(0)), globalState))) ? ((_2085_v20).get(_module.__default.fm0(new _dafny.CodePoint('a'.codePointAt(0)), globalState))) : (_2084_v19)));
      let _2087_v22;
      _2087_v22 = _dafny.MultiSet.fromElements(new BigNumber(((_2086_v21)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_2086_v21).length))]).length), _2049_v0.f16, p0);
      let _2088_v23;
      _2088_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2087_v22,_dafny.Seq.UnicodeFromString("rcxfakrum"));
      let _2089_v24;
      let _2090_v25;
      let _out60;
      let _out61;
      let _outcollector26 = _module.__default.m0(_2088_v23, (_this).f4, _2087_v22, (_this).f4, globalState);
      _out60 = _outcollector26[0];
      _out61 = _outcollector26[1];
      _2089_v24 = _out60;
      _2090_v25 = _out61;
      let _index300 = _module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length));
      ((_this).f10)[_index300] = _2089_v24;
      let _2091_v26;
      _2091_v26 = _module.D11.create_DC26(!((_this).f4));
      let _2092_v27;
      _2092_v27 = _dafny.Set.fromElements(_2091_v26, _2091_v26, _2091_v26, _2091_v26);
      let _index301 = _module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length));
      ((_this).f10)[_index301] = !(_2090_v25).isEqualTo(new BigNumber(((_2092_v27).Intersect(_2092_v27)).length));
      let _2093_v28;
      _2093_v28 = _dafny.Seq.UnicodeFromString("lvmi");
      let _2094_v29;
      _2094_v29 = _dafny.Map.Empty.slice().updateUnsafe(false,((_this).f10)[_module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length))]);
      let _2095_v30;
      _2095_v30 = _dafny.Seq.of((_this).f6, new BigNumber((_2050_v1).length));
      let _2096_v31;
      _2096_v31 = _dafny.Seq.of(_2049_v0.f16, new BigNumber((_2095_v30).length));
      let _2097_v32;
      let _nw315 = Array((new BigNumber(14)).toNumber());
      _nw315[(_dafny.ZERO).toNumber()] = (_this).fm7(globalState);
      _nw315[(_dafny.ONE).toNumber()] = (_2049_v0).f15;
      _nw315[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2093_v28).length), (_2049_v0).f15);
      _nw315[(new BigNumber(3)).toNumber()] = new BigNumber((_2094_v29).length);
      _nw315[(new BigNumber(4)).toNumber()] = (_this).f6;
      _nw315[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(_2090_v25, new BigNumber((_2050_v1).length));
      _nw315[(new BigNumber(6)).toNumber()] = ((_this).f6).multipliedBy(p0);
      _nw315[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(34), new BigNumber(583));
      _nw315[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2096_v31).length), new BigNumber(967));
      _nw315[(new BigNumber(9)).toNumber()] = _2049_v0.f16;
      _nw315[(new BigNumber(10)).toNumber()] = (_2049_v0).f15;
      _nw315[(new BigNumber(11)).toNumber()] = _2049_v0.f16;
      _nw315[(new BigNumber(12)).toNumber()] = _2090_v25;
      _nw315[(new BigNumber(13)).toNumber()] = (_2049_v0).f15;
      _2097_v32 = _nw315;
      let _2098_v33;
      _2098_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(30),((_this).f10)[_module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length))]);
      let _2099_v34;
      let _nw316 = Array((new BigNumber(25)).toNumber()).fill(false);
      _2099_v34 = _nw316;
      let _2100_v35;
      _2100_v35 = _module.D10.create_DC22(_dafny.Seq.UnicodeFromString("tsfhsrurn"), _2089_v24, (_this).f4, _2099_v34, (_this).f6);
      let _2101_v36;
      _2101_v36 = _dafny.MultiSet.fromElements((_this).f4, _2089_v24);
      let _2102_v37;
      let _nw317 = new _module.C4();
      _nw317.__ctor(_this.f5, _2090_v25, _2089_v24);
      _2102_v37 = _nw317;
      let _2103_v38;
      _2103_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2102_v37,_2093_v28);
      let _2104_v39;
      _2104_v39 = _dafny.Set.fromElements((_this).f4, _2089_v24);
      let _2105_v40;
      _2105_v40 = _module.D16.create_DC33((((_2098_v33).contains(_2049_v0.f16)) ? ((_2098_v33).get(_2049_v0.f16)) : (((_this).f10)[_module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length))])), (_2102_v37).fm7(globalState), _dafny.MultiSet.fromElements(((_this).f10)[_module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length))]), new BigNumber((_2087_v22).cardinality()), _2104_v39);
      let _2106_v41;
      _2106_v41 = _module.D16.create_DC33((_this).fm14(_2098_v33, ((_this).f10)[_module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length))], (_2049_v0).f15, (_2100_v35).dtor_cf32, globalState), (_dafny.ZERO).minus(((((_2084_v19).contains(_2049_v0.f16)) ? ((_2084_v19).get(_2049_v0.f16)) : (new BigNumber((_2087_v22).cardinality())))).plus(new BigNumber(431))), ((_2101_v36).update(_2089_v24, _module.__default.abs(new BigNumber((_2103_v38).length)))).Intersect(_2101_v36), (_dafny.ZERO).minus((_2049_v0.f16).multipliedBy(_2049_v0.f16)), _dafny.Set.fromElements(_2089_v24, _2089_v24, true, !(_2089_v24), !((_2105_v40).dtor_cf52)));
      let _index302 = _module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length));
      let _rhs286 = new BigNumber(154);
      let _rhs287 = new BigNumber(842);
      let _rhs288 = _2097_v32;
      let _rhs289 = _2105_v40;
      let _rhs290 = ((_this).f10)[_module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length))];
      let _lhs175 = _2049_v0;
      let _lhs176 = (_this).f10;
      let _lhs177 = _module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length));
      _lhs175.f16 = _rhs286;
      _2090_v25 = _rhs287;
      _2097_v32 = _rhs288;
      _2105_v40 = _rhs289;
      _lhs176[_lhs177] = _rhs290;
      let _index303 = _module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length));
      let _rhs291 = p0;
      let _rhs292 = (((_2090_v25).isEqualTo(_2049_v0.f16)) ? (false) : (false));
      let _lhs178 = _2049_v0;
      let _lhs179 = (_this).f10;
      let _lhs180 = _module.__default.safeIndex(new BigNumber(270), new BigNumber(((_this).f10).length));
      _lhs178.f16 = _rhs291;
      _lhs179[_lhs180] = _rhs292;
      let _2107_v42;
      _2107_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0);
      r0 = _2107_v42;
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = [];
      let _2108_v0;
      _2108_v0 = new BigNumber(-830);
      let _2109_v1;
      _2109_v1 = _dafny.Seq.UnicodeFromString("q");
      let _2110_v2;
      let _nw318 = new _module.C3();
      _nw318.__ctor(_this.f5, new BigNumber((_2109_v1).length), p0);
      _2110_v2 = _nw318;
      let _2111_v3;
      _2111_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2110_v2,false);
      _2108_v0 = (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2111_v3,(_this).f6)).length), _2108_v0)).multipliedBy((_this).f6);
      let _2112_i0;
      _2112_i0 = _dafny.ZERO;
      L9: {
        while (p0) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2112_i0)) {
              break L9;
            }
            _2112_i0 = (_2112_i0).plus(_dafny.ONE);
            let _2113_v4;
            let _nw319 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2113_v4 = _nw319;
            let _index304 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_2113_v4).length));
            (_2113_v4)[_index304] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(578)), function (_2114_i1) {
              return new _dafny.CodePoint('o'.codePointAt(0));
            });
            let _2115_v5;
            _2115_v5 = false;
            let _2116_v6;
            let _init52 = function (_2117_i2) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            };
            let _nw320 = Array((new BigNumber(13)).toNumber());
            for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw320.length); _i0_52++) {
              _nw320[_i0_52] = _init52(new BigNumber(_i0_52));
            }
            _2116_v6 = _nw320;
            let _2118_v7;
            _2118_v7 = new _dafny.CodePoint('h'.codePointAt(0));
            let _index305 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_2116_v6).length));
            (_2116_v6)[_index305] = _2118_v7;
            let _2119_v8;
            _2119_v8 = _dafny.Set.fromElements(new BigNumber(474), (_this).f6, (_this).f6, (_dafny.ZERO).minus((_this).fm7(globalState)));
            let _2120_v9;
            _2120_v9 = _dafny.Seq.of(!((_this).f4));
            let _2121_v10;
            _2121_v10 = _dafny.Seq.of(_2109_v1);
            let _2122_v11;
            _2122_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2121_v10);
            let _index306 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_2113_v4).length));
            let _index307 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_2116_v6).length));
            let _rhs293 = _2109_v1;
            let _rhs294 = _dafny.Seq.contains(_2120_v9, true);
            let _rhs295 = new BigNumber(((((_2122_v11).contains((_this).f10)) ? ((_2122_v11).get((_this).f10)) : (_dafny.Seq.Concat(_2121_v10, _2121_v10)))).length);
            let _rhs296 = _2118_v7;
            let _rhs297 = _2119_v8;
            let _lhs181 = _2113_v4;
            let _lhs182 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_2113_v4).length));
            let _lhs183 = _2116_v6;
            let _lhs184 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_2116_v6).length));
            _lhs181[_lhs182] = _rhs293;
            _2115_v5 = _rhs294;
            _2108_v0 = _rhs295;
            _lhs183[_lhs184] = _rhs296;
            _2119_v8 = _rhs297;
            _2108_v0 = _module.__default.safeDivisionInt(new BigNumber(404), new BigNumber((_2109_v1).length));
            let _2123_v12;
            _2123_v12 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2118_v7);
            let _2124_v13;
            _2124_v13 = _dafny.Seq.of((_module.__default.fm65(_2108_v0, new BigNumber(625), new BigNumber(209), globalState)).update(true, (_2116_v6)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_2116_v6).length))]), _2123_v12);
            _2108_v0 = (_dafny.ZERO).minus(new BigNumber((((_2123_v12).update(_2115_v5, (_2116_v6)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_2116_v6).length))])).Merge((_2123_v12).Merge((_2124_v13)[_module.__default.safeIndex(new BigNumber((_2109_v1).length), new BigNumber((_2124_v13).length))]))).length));
            _2115_v5 = p0;
          }
        }
      }
      let _2125_v14;
      let _nw321 = new _module.C6();
      _nw321.__ctor((_this).f4, p0, _this.f5, _2108_v0);
      _2125_v14 = _nw321;
      let _2126_v15;
      _2126_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2125_v14);
      let _2127_v16;
      _2127_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
      let _2128_v17;
      _2128_v17 = _module.D3.create_DC9(false, (new BigNumber((_2126_v15).length)).plus(new BigNumber((_2127_v16).length)), (_2125_v14).f4);
      let _source36 = _2128_v17;
      if (_source36.is_DC9) {
        let _2129___mcc_h0 = (_source36).cf8;
        let _2130___mcc_h1 = (_source36).cf9;
        let _2131___mcc_h2 = (_source36).cf10;
        let _2132_cf10 = _2131___mcc_h2;
        let _2133_cf9 = _2130___mcc_h1;
        let _2134_cf8 = _2129___mcc_h0;
        let _2135_v18;
        _2135_v18 = _dafny.Set.fromElements((_2125_v14).f4, p0, (_this).f4);
        let _2136_v19;
        let _init53 = ((_2137_cf9) => function (_2138_i3) {
          return _module.__default.safeModuloInt(_2138_i3, _2137_cf9);
        })(_2133_cf9);
        let _nw322 = Array((new BigNumber(7)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw322.length); _i0_53++) {
          _nw322[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _2136_v19 = _nw322;
        let _2139_v20;
        _2139_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2136_v19,_2135_v18);
        let _2140_v21;
        _2140_v21 = new _dafny.CodePoint('x'.codePointAt(0));
        let _2141_v22;
        _2141_v22 = _dafny.Map.Empty.slice().updateUnsafe(true,(_2125_v14).f4);
        let _2142_v23;
        let _nw323 = Array((new BigNumber(25)).toNumber());
        _nw323[(_dafny.ZERO).toNumber()] = _2135_v18;
        _nw323[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(true);
        _nw323[(new BigNumber(2)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(3)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(4)).toNumber()] = (_dafny.Set.fromElements(_2132_cf10, p0, (_2125_v14).f4)).Difference(_2135_v18);
        _nw323[(new BigNumber(5)).toNumber()] = ((_2134_cf8) ? (_2135_v18) : (_2135_v18));
        _nw323[(new BigNumber(6)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(7)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(_2134_cf8);
        _nw323[(new BigNumber(9)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(10)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(11)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(p0, (_2125_v14).f4, (_2125_v14).f4, _2132_cf10, (_2125_v14).f4);
        _nw323[(new BigNumber(13)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(14)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(15)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(16)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(17)).toNumber()] = (_2135_v18).Difference(_2135_v18);
        _nw323[(new BigNumber(18)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(19)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(20)).toNumber()] = ((((_2139_v20).contains(_2136_v19)) ? ((_2139_v20).get(_2136_v19)) : (_2135_v18))).Union(_module.__default.fm46(_2140_v21, globalState));
        _nw323[(new BigNumber(21)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(22)).toNumber()] = (_2135_v18).Intersect(_2135_v18);
        _nw323[(new BigNumber(23)).toNumber()] = _2135_v18;
        _nw323[(new BigNumber(24)).toNumber()] = (_2135_v18).Union(_module.__default.fm53((_this).f6, _dafny.Seq.of(_2141_v22, _2141_v22), globalState));
        _2142_v23 = _nw323;
        let _2143_v24;
        _2143_v24 = _dafny.MultiSet.fromElements((_2125_v14).f4);
        let _index308 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_2142_v23).length));
        (_2142_v23)[_index308] = _module.__default.fm53(_2133_cf9, _module.__default.fm54(new BigNumber((_2143_v24).cardinality()), globalState), globalState);
        let _index309 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_2142_v23).length));
        (_2142_v23)[_index309] = _2135_v18;
        let _2144_v25;
        _2144_v25 = _this.f5;
        let _2145_v26;
        let _nw324 = new _module.C5();
        _nw324.__ctor(_2132_cf10, (_this).f6, (_2144_v25), _2133_cf9, (_2125_v14).f4);
        _2145_v26 = _nw324;
        if ((!((_2145_v26).f13)) || ((_2108_v0).isLessThan(_2108_v0))) {
          let _2146_v27;
          let _nw325 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Set.Empty);
          _2146_v27 = _nw325;
          let _2147_v28;
          let _nw326 = Array((new BigNumber(6)).toNumber()).fill(false);
          _2147_v28 = _nw326;
          let _index310 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_2136_v19).length));
          (_2136_v19)[_index310] = _module.__default.fm3(_2108_v0, _2134_cf8, globalState);
          let _index311 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_2136_v19).length));
          let _rhs298 = _2146_v27;
          let _rhs299 = _2147_v28;
          let _rhs300 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_2108_v0), new BigNumber(61)));
          let _lhs185 = _2136_v19;
          let _lhs186 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_2136_v19).length));
          _2146_v27 = _rhs298;
          _2147_v28 = _rhs299;
          _lhs185[_lhs186] = _rhs300;
          let _2148_v29;
          _2148_v29 = _dafny.MultiSet.fromElements((_this).f10);
          _2148_v29 = _dafny.MultiSet.fromElements((_this).f10);
          _2132_cf10 = ((_2108_v0).plus((_this).f6)).isLessThan(((_dafny.ZERO).minus((_dafny.ZERO).minus((_2136_v19)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_2136_v19).length))]))).multipliedBy(_2145_v26.f14));
          let _2149_v30;
          _2149_v30 = _dafny.Seq.of((_dafny.ZERO).minus(_2145_v26.f14));
          _2108_v0 = (_2149_v30)[_module.__default.safeIndex(_2145_v26.f14, new BigNumber((_2149_v30).length))];
          let _2150_v31;
          let _nw327 = new _module.C4();
          _nw327.__ctor(_this.f5, _2145_v26.f14, _2134_cf8);
          _2150_v31 = _nw327;
        } else {
          let _2151_v32;
          _2151_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2108_v0);
          _2127_v16 = (_2127_v16).update(((_this).f6).plus(new BigNumber((_2151_v32).length)), _2133_cf9);
          _2108_v0 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_2145_v26.f14), _2145_v26.f14);
          _2108_v0 = new BigNumber((_2109_v1).length);
          let _2152_v33;
          _2152_v33 = _module.D2.create_DC5((_this).f10);
          _2152_v33 = _2152_v33;
          let _2153_v34;
          _2153_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2145_v26.f14,_2132_cf10);
          _2153_v34 = (_2153_v34).update(((_this).f6).minus((_this).f6), !((_2125_v14).f4));
        }
        _2132_cf10 = !((_2125_v14).f4);
      } else if (_source36.is_DC10) {
        let _2154___mcc_h3 = (_source36).cf11;
        let _2155_cf11 = _2154___mcc_h3;
        let _2156_v35;
        _2156_v35 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f4)).length)),(_this).f10);
        _2156_v35 = (_2156_v35).update((_this).f6, (_this).f10);
        let _2157_v36;
        _2157_v36 = _module.D12.create_DC28(_2108_v0, _2155_cf11, false, (_this).f4, _2155_cf11);
        if ((_2157_v36).dtor_cf46) {
          let _2158_v37;
          _2158_v37 = _dafny.Map.Empty.slice().updateUnsafe((_2125_v14).f4,_2109_v1);
          let _2159_v39;
          let _nw328 = Array((new BigNumber(7)).toNumber());
          _nw328[(_dafny.ZERO).toNumber()] = new BigNumber(154);
          _nw328[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2109_v1).length), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,_2108_v0)).update(p0, _2108_v0)).length));
          _nw328[(new BigNumber(2)).toNumber()] = (_this).fm7(globalState);
          _nw328[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of _dafny.IntegerRange(new BigNumber(587), new BigNumber(174))) {
              let _2160_v38 = _compr_57;
              if (((new BigNumber(587)).isLessThanOrEqualTo(_2160_v38)) && ((_2160_v38).isLessThan(new BigNumber(174)))) {
                _coll57.add((_2160_v38).minus((_this).f6));
              }
            }
            return _coll57;
          }()).length), _2155_cf11);
          _nw328[(new BigNumber(4)).toNumber()] = (_this).f6;
          _nw328[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements((_2125_v14).f4, (_this).f4, p0)).cardinality());
          _nw328[(new BigNumber(6)).toNumber()] = (_this).f6;
          _2159_v39 = _nw328;
          let _2161_v40;
          _2161_v40 = _dafny.Seq.of(p0);
          let _index312 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2159_v39).length));
          (_2159_v39)[_index312] = new BigNumber((_dafny.Seq.Concat(_2161_v40, _2161_v40)).length);
          let _2162_v41;
          let _nw329 = Array((new BigNumber(27)).toNumber()).fill([]);
          _2162_v41 = _nw329;
          let _index313 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_2162_v41).length));
          (_2162_v41)[_index313] = _2159_v39;
          let _2163_v42;
          _2163_v42 = false;
          let _2164_v43;
          _2164_v43 = new _dafny.CodePoint('h'.codePointAt(0));
          let _2165_v44;
          _2165_v44 = _dafny.Map.Empty.slice().updateUnsafe((_2125_v14).f4,_2108_v0);
          let _index314 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2159_v39).length));
          let _index315 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_2162_v41).length));
          let _rhs301 = (_dafny.Map.Empty.slice().updateUnsafe((_this).f4,_2109_v1)).Merge(((_this).f9).Merge(_2158_v37));
          let _rhs302 = _2108_v0;
          let _rhs303 = _2159_v39;
          let _rhs304 = _module.__default.fm0(_2164_v43, globalState);
          let _rhs305 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_2165_v44).Merge(_2165_v44)).length)))).isEqualTo(new BigNumber(-513));
          let _lhs187 = _2159_v39;
          let _lhs188 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2159_v39).length));
          let _lhs189 = _2162_v41;
          let _lhs190 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_2162_v41).length));
          _2158_v37 = _rhs301;
          _lhs187[_lhs188] = _rhs302;
          _lhs189[_lhs190] = _rhs303;
          _2163_v42 = _rhs304;
          _2163_v42 = _rhs305;
          let _2166_v45;
          _2166_v45 = _dafny.Set.fromElements(!((_2125_v14).f4), !((_2125_v14).f4));
          let _2167_v46;
          _2167_v46 = _module.D16.create_DC33(false, _2155_cf11, _dafny.MultiSet.fromElements(false), (_this).f6, _2166_v45);
          _2163_v42 = (_dafny.MultiSet.fromElements((_2125_v14).f4)).IsDisjointFrom((_2167_v46).dtor_cf54);
          let _2168_v47;
          let _nw330 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Map.Empty);
          _2168_v47 = _nw330;
          let _2169_v48;
          _2169_v48 = _dafny.MultiSet.fromElements((_2159_v39)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_2159_v39).length))], new BigNumber((_2165_v44).length), _2155_cf11);
          let _2170_v49;
          let _nw331 = Array((new BigNumber(12)).toNumber());
          _nw331[(_dafny.ZERO).toNumber()] = _2164_v43;
          _nw331[(_dafny.ONE).toNumber()] = _module.__default.fm1((_2159_v39)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_2159_v39).length))], (_2125_v14).f4, globalState);
          _nw331[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
          _nw331[(new BigNumber(3)).toNumber()] = _2164_v43;
          _nw331[(new BigNumber(4)).toNumber()] = _2164_v43;
          _nw331[(new BigNumber(5)).toNumber()] = _2164_v43;
          _nw331[(new BigNumber(6)).toNumber()] = _2164_v43;
          _nw331[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _nw331[(new BigNumber(8)).toNumber()] = _2164_v43;
          _nw331[(new BigNumber(9)).toNumber()] = _2164_v43;
          _nw331[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
          _nw331[(new BigNumber(11)).toNumber()] = _2164_v43;
          _2170_v49 = _nw331;
          let _2171_v50;
          _2171_v50 = _dafny.Map.Empty.slice().updateUnsafe((((_2169_v48).contains(_2155_cf11)) ? ((_2169_v48).get(_2155_cf11)) : (new BigNumber(879))),_2170_v49);
          let _index316 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_2168_v47).length));
          (_2168_v47)[_index316] = _2171_v50;
          let _2172_v51;
          _2172_v51 = _module.D11.create_DC25(true, _2155_cf11);
          let _2173_v52;
          _2173_v52 = _dafny.Set.fromElements(_2108_v0);
          let _2174_v53;
          _2174_v53 = _dafny.Seq.of((_2172_v51).dtor_cf40, new BigNumber((_dafny.Seq.of(_2173_v52)).length));
          let _2175_v54;
          _2175_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2174_v53).length),(_2125_v14).f4);
          let _2176_v55;
          _2176_v55 = _dafny.Set.fromElements(_2175_v54);
          let _2177_v56;
          _2177_v56 = _dafny.MultiSet.fromElements((_2125_v14).f4, !(_2176_v55).equals(_2176_v55), _2163_v42);
          let _index317 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_2168_v47).length));
          let _rhs306 = (p0) || ((_2125_v14).f4);
          let _rhs307 = _2163_v42;
          let _rhs308 = new BigNumber((_2177_v56).cardinality());
          let _rhs309 = ((p0) ? ((_2171_v50).Merge(_2171_v50)) : ((_2171_v50).Merge(_2171_v50)));
          let _lhs191 = _2168_v47;
          let _lhs192 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_2168_v47).length));
          _2163_v42 = _rhs306;
          _2163_v42 = _rhs307;
          _2108_v0 = _rhs308;
          _lhs191[_lhs192] = _rhs309;
          let _2178_v57;
          let _nw332 = new _module.C1();
          _nw332.__ctor((_2159_v39)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_2159_v39).length))], (_2159_v39)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_2159_v39).length))], _this.f5, (((_2165_v44).contains(p0)) ? ((_2165_v44).get(p0)) : ((_this).f6)), (_2125_v14).f4);
          _2178_v57 = _nw332;
          let _2179_v58;
          _2179_v58 = _dafny.Seq.of((((_this).f4) ? (_2178_v57) : (_2178_v57)), _2178_v57, _2178_v57, _2178_v57, _2178_v57);
          let _2180_v59;
          _2180_v59 = _dafny.Seq.of(_dafny.Seq.Concat(_2179_v58, _dafny.Seq.update(_2179_v58, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length), new BigNumber((_2179_v58).length)), _2178_v57)));
          _2179_v58 = (_2180_v59)[_module.__default.safeIndex(new BigNumber(628), new BigNumber((_2180_v59).length))];
          let _2181_v60;
          _2181_v60 = _module.D9.create_DC20((_2125_v14).f4, _2164_v43, _module.__default.safeDivisionInt((_this).f6, (_this).f6));
          let _2182_v61;
          _2182_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2164_v43,(_this).f10);
          let _2183_v62;
          _2183_v62 = _module.D19.create_DC37(_2182_v61);
          let _index318 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2159_v39).length));
          let _rhs310 = (new BigNumber((function () {
            let _coll58 = new _dafny.Set();
            for (const _compr_58 of _dafny.IntegerRange(new BigNumber(375), new BigNumber(806))) {
              let _2184_v63 = _compr_58;
              if (((new BigNumber(375)).isLessThanOrEqualTo(_2184_v63)) && ((_2184_v63).isLessThan(new BigNumber(806)))) {
                _coll58.add((_2184_v63).multipliedBy((_2178_v57).f15));
              }
            }
            return _coll58;
          }()).length)).multipliedBy(_2108_v0);
          let _rhs311 = _2181_v60;
          let _rhs312 = (_2125_v14).f4;
          let _rhs313 = _2183_v62;
          let _lhs193 = _2159_v39;
          let _lhs194 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2159_v39).length));
          _lhs193[_lhs194] = _rhs310;
          _2181_v60 = _rhs311;
          _2163_v42 = _rhs312;
          _2183_v62 = _rhs313;
        } else {
          _2155_cf11 = _2108_v0;
          let _2185_v64;
          _2185_v64 = false;
          let _index319 = _module.__default.safeIndex(new BigNumber(635), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index319] = !((_2125_v14).f4);
          let _2186_v65;
          let _nw333 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _2186_v65 = _nw333;
          let _2187_v66;
          _2187_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2108_v0,_2186_v65);
          let _2188_v67;
          _2188_v67 = _dafny.Set.fromElements(false);
          let _2189_v68;
          _2189_v68 = _dafny.Seq.of(_2188_v67, _2188_v67, _2188_v67, _2188_v67);
          let _2190_v69;
          _2190_v69 = _2127_v16;
          let _index320 = _module.__default.safeIndex(new BigNumber(635), new BigNumber(((_this).f10).length));
          let _rhs314 = (new BigNumber((_2187_v66).length)).isLessThanOrEqualTo(_2108_v0);
          let _rhs315 = !(!(_dafny.MultiSet.fromElements(_2155_cf11)).equals((_dafny.MultiSet.fromElements((_this).f6, new BigNumber(762))).update((_this).f6, _module.__default.abs(new BigNumber((_2189_v68).length)))));
          let _rhs316 = true;
          let _rhs317 = !(!(!(_2127_v16).equals((_2190_v69))) || ((_2125_v14).f4));
          let _lhs195 = (_this).f10;
          let _lhs196 = _module.__default.safeIndex(new BigNumber(635), new BigNumber(((_this).f10).length));
          _2185_v64 = _rhs314;
          _2185_v64 = _rhs315;
          _lhs195[_lhs196] = _rhs316;
          _2185_v64 = _rhs317;
          _2108_v0 = _2108_v0;
          let _index321 = _module.__default.safeIndex(new BigNumber(635), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index321] = !(p0) || (((_2125_v14).f4) || (p0));
          let _2191_v70;
          _2191_v70 = _dafny.Map.Empty.slice().updateUnsafe(p0,((_this).f10)[_module.__default.safeIndex(new BigNumber(635), new BigNumber(((_this).f10).length))]);
          let _2192_v71;
          _2192_v71 = _dafny.Seq.of(_2155_cf11);
          let _2193_v72;
          _2193_v72 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC16(true, (((_2191_v70).contains((_2125_v14).f4)) ? ((_2191_v70).get((_2125_v14).f4)) : ((_2125_v14).f4)), _2192_v71),!(_2185_v64));
          _2193_v72 = _2193_v72;
        }
        let _2194_v73;
        _2194_v73 = _dafny.Seq.of(true);
        let _2195_v74;
        _2195_v74 = _dafny.Set.fromElements(_2108_v0, new BigNumber((_2194_v73).length), _2108_v0);
        if (((!((_2125_v14).f4)) ? ((_2195_v74).IsProperSubsetOf(_2195_v74)) : (p0))) {
          let _2196_v75;
          _2196_v75 = false;
          let _2197_v76;
          _2197_v76 = _dafny.MultiSet.fromElements(new BigNumber(((_module.D10.create_DC22(_module.__default.fm50(_2108_v0, _2108_v0, _2155_cf11, (_2125_v14).f4, globalState), p0, (_this).f4, (_this).f10, (_this).f6)).dtor_cf30).length), new BigNumber(-163), _2155_cf11);
          _2196_v75 = (new BigNumber(((_2197_v76).update(_2155_cf11, _module.__default.abs((_this).f6))).cardinality())).isEqualTo(_2108_v0);
          _2155_cf11 = (_this).f6;
          _2155_cf11 = (((_this).f6).plus((_this).f6)).minus(new BigNumber((_2109_v1).length));
          _2155_cf11 = (((_module.__default.fm66(_2108_v0, globalState)).dtor_cf52) ? ((_2110_v2).fm7(globalState)) : (new BigNumber((_2194_v73).length)));
          let _index322 = _module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index322] = p0;
          let _index323 = _module.__default.safeIndex(new BigNumber(31), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index323] = (_2108_v0).isEqualTo(_2108_v0);
        } else {
          let _2198_v77;
          _2198_v77 = true;
          let _2199_v78;
          _2199_v78 = _module.D11.create_DC25(_2198_v77, (_this).f6);
          _2198_v77 = (_2199_v78).dtor_cf39;
          let _2200_v79;
          _2200_v79 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2155_cf11, (_this).f6),(_2125_v14).f4);
          let _2201_v80;
          _2201_v80 = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_this).f6)).cardinality()), _2155_cf11);
          let _2202_v81;
          let _nw334 = new _module.C1();
          _nw334.__ctor(new BigNumber(((_2200_v79).update(_2201_v80, (_2125_v14).f4)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(400)), ((_2203_cf11) => function (_2204_i4) {
            return _2203_cf11;
          })(_2155_cf11))).length), _this.f5, _2155_cf11, (_2125_v14).f4);
          _2202_v81 = _nw334;
          let _2205_v82;
          _2205_v82 = _dafny.Map.Empty.slice().updateUnsafe(_2198_v77,(_dafny.ZERO).minus((_this).f6));
          (_2202_v81).f16 = (_2202_v81.f16).plus(_module.__default.safeModuloInt(new BigNumber((_2205_v82).length), (_2125_v14).fm7(globalState)));
          let _2206_v83;
          _2206_v83 = _dafny.MultiSet.fromElements(_2202_v81.f16, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2202_v81.f16,_2198_v77)).length), (_2202_v81).f15);
          let _2207_v84;
          _2207_v84 = _dafny.Seq.of(_2194_v73);
          let _2208_v85;
          _2208_v85 = _dafny.MultiSet.fromElements(((_dafny.ZERO).minus((_dafny.ZERO).minus(_2202_v81.f16))).multipliedBy(_2155_cf11), _module.__default.safeDivisionInt((_this).f6, _2108_v0), (((_2206_v83).contains(new BigNumber(220))) ? ((_2206_v83).get(new BigNumber(220))) : ((((_2205_v82).contains((_this).f4)) ? ((_2205_v82).get((_this).f4)) : ((_module.D4.create_DC13(new BigNumber((_2127_v16).length), _2198_v77, _2108_v0, (_2207_v84)[_module.__default.safeIndex(_2202_v81.f16, new BigNumber((_2207_v84).length))])).dtor_cf16)))), (_this).f6);
          let _2209_v86;
          _2209_v86 = _dafny.Map.Empty.slice().updateUnsafe(_2202_v81.f16,(_this).f4);
          _2208_v85 = ((_dafny.MultiSet.FromArray(_2201_v80)).update(new BigNumber((_dafny.Seq.UnicodeFromString("tfiwvdm")).length), _module.__default.abs(new BigNumber((_2209_v86).length)))).update(_2108_v0, _module.__default.abs((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f6))));
          let _2210_v87;
          let _nw335 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
          _2210_v87 = _nw335;
          let _2211_v88;
          _2211_v88 = _module.D10.create_DC21(_2205_v82);
          let _2212_v89;
          _2212_v89 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v88,_2108_v0);
          let _2213_v90;
          _2213_v90 = _dafny.Seq.of(_2212_v89);
          let _index324 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_2210_v87).length));
          (_2210_v87)[_index324] = _2213_v90;
          let _index325 = _module.__default.safeIndex(new BigNumber(172), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index325] = (_2125_v14).f4;
          let _2214_v91;
          _2214_v91 = _dafny.Set.fromElements(_2109_v1, _dafny.Seq.UnicodeFromString("ugmeom"), _2109_v1);
          let _2215_v92;
          _2215_v92 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2109_v1, _dafny.Seq.UnicodeFromString("q")),(new BigNumber((_dafny.MultiSet.fromElements((_this).f6, _2202_v81.f16, _2155_cf11, _2155_cf11, (_dafny.ZERO).minus(new BigNumber((_2214_v91).length)))).cardinality())).multipliedBy(new BigNumber((_2201_v80).length)));
          let _2216_v93;
          _2216_v93 = new _dafny.CodePoint('w'.codePointAt(0));
          let _index326 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_2210_v87).length));
          let _index327 = _module.__default.safeIndex(new BigNumber(172), new BigNumber(((_this).f10).length));
          let _rhs318 = p0;
          let _rhs319 = _2213_v90;
          let _rhs320 = (_2125_v14).f4;
          let _rhs321 = _2216_v93;
          let _rhs322 = (_2215_v92).update(_2109_v1, _2202_v81.f16);
          let _lhs197 = _2210_v87;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_2210_v87).length));
          let _lhs199 = (_this).f10;
          let _lhs200 = _module.__default.safeIndex(new BigNumber(172), new BigNumber(((_this).f10).length));
          _2198_v77 = _rhs318;
          _lhs197[_lhs198] = _rhs319;
          _lhs199[_lhs200] = _rhs320;
          r0 = _rhs321;
          _2215_v92 = _rhs322;
        }
        let _arr11 = _this.f5;
        let _index328 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_this.f5).length));
        _arr11[_index328] = (_this).f10;
        let _2217_v94;
        _2217_v94 = _dafny.Seq.of((_this).f10, (_this).f10, (((_2125_v14).f4) ? ((_this).f10) : ((_this).f10)), (_this).f10, (((_2156_v35).contains(_2108_v0)) ? ((_2156_v35).get(_2108_v0)) : ((_this).f10)));
        let _arr12 = _this.f5;
        let _index329 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_this.f5).length));
        _arr12[_index329] = (_2217_v94)[_module.__default.safeIndex(_module.__default.safeModuloInt(_2108_v0, new BigNumber(122)), new BigNumber((_2217_v94).length))];
      } else if (_source36.is_DC8) {
        let _2218___mcc_h4 = (_source36).cf7;
        let _2219_cf7 = _2218___mcc_h4;
        let _2220_v95;
        _2220_v95 = _module.__default.fm29(globalState);
        let _2221_v96;
        _2221_v96 = _dafny.Seq.of(_2219_cf7);
        let _2222_v97;
        _2222_v97 = _module.D20.create_DC40(_dafny.Map.Empty.slice().updateUnsafe((_this).f4,_module.D12.create_DC27(_2221_v96)));
        let _2223_v98;
        _2223_v98 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2219_cf7).length),_2222_v97);
        let _index330 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index330] = (_this).f4;
        let _2224_v99;
        _2224_v99 = _dafny.MultiSet.fromElements(_2219_cf7, _dafny.Seq.of(p0, (_this).f4, p0));
        let _2225_v100;
        _2225_v100 = _module.D6.create_DC15((_2224_v99).update(_dafny.Seq.update(_2219_cf7, _module.__default.safeIndex(new BigNumber(-524), new BigNumber((_2219_cf7).length)), (_2125_v14).f4), _module.__default.abs(_2108_v0)));
        let _index331 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
        let _rhs323 = (_this).f6;
        let _rhs324 = _2220_v95;
        let _rhs325 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2222_v97);
        let _rhs326 = p0;
        let _rhs327 = _2225_v100;
        let _lhs201 = (_this).f10;
        let _lhs202 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
        _2108_v0 = _rhs323;
        _2220_v95 = _rhs324;
        _2223_v98 = _rhs325;
        _lhs201[_lhs202] = _rhs326;
        _2225_v100 = _rhs327;
        let _2226_v101;
        let _init54 = ((_2227_v1) => function (_2228_i5) {
          return _dafny.Seq.update(_dafny.Seq.Concat(_2227_v1, _2227_v1), _module.__default.safeIndex((_this).f6, new BigNumber((_dafny.Seq.Concat(_2227_v1, _2227_v1)).length)), new _dafny.CodePoint('s'.codePointAt(0)));
        })(_2109_v1);
        let _nw336 = Array((new BigNumber(4)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw336.length); _i0_54++) {
          _nw336[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _2226_v101 = _nw336;
        let _index332 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2226_v101).length));
        (_2226_v101)[_index332] = _2109_v1;
        let _2229_v102;
        _2229_v102 = _dafny.Seq.of((_this).f6);
        let _2230_v103;
        _2230_v103 = new _dafny.CodePoint('o'.codePointAt(0));
        let _index333 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2226_v101).length));
        (_2226_v101)[_index333] = _dafny.Seq.update(_module.__default.fm37((_this).f4, (((_2125_v14).f4) ? ((_2125_v14).f4) : ((_2110_v2).fm8(_2108_v0, true, _2109_v1, _2229_v102, globalState))), (_this).f6, (_2125_v14).f4, globalState), _module.__default.safeIndex(_2108_v0, new BigNumber((_module.__default.fm37((_this).f4, (((_2125_v14).f4) ? ((_2125_v14).f4) : ((_2110_v2).fm8(_2108_v0, true, _2109_v1, _2229_v102, globalState))), (_this).f6, (_2125_v14).f4, globalState)).length)), _2230_v103);
        if (p0) {
          let _index334 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index334] = (!((_2219_cf7)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_2219_cf7).length))]) || ((_2125_v14).f4)) && ((_this).f4);
          let _2231_v104;
          let _init55 = function (_2232_i6) {
            return (_module.D1.create_DC3(((_this).f10)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length))])).dtor_cf2;
          };
          let _nw337 = Array((new BigNumber(6)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw337.length); _i0_55++) {
            _nw337[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _2231_v104 = _nw337;
          _2231_v104 = ((_dafny.Seq.IsPrefixOf(_2229_v102, _2229_v102)) ? (_2231_v104) : (_2231_v104));
          _2108_v0 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_this).f6, new BigNumber((_2219_cf7).length)), new BigNumber(846));
          let _index335 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index335] = !(((_this).f10)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length))]);
          _2108_v0 = _2108_v0;
        } else {
          let _index336 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          let _index337 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2226_v101).length));
          let _index338 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          let _rhs328 = !(!((_2125_v14).f4) || (true));
          let _rhs329 = _dafny.Seq.Concat((_2226_v101)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_2226_v101).length))], _2109_v1);
          let _rhs330 = (((_2125_v14).f4) ? ((_2108_v0).isLessThanOrEqualTo(_2108_v0)) : ((_2108_v0).isLessThanOrEqualTo((_this).f6)));
          let _rhs331 = _2108_v0;
          let _lhs203 = (_this).f10;
          let _lhs204 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          let _lhs205 = _2226_v101;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_2226_v101).length));
          let _lhs207 = (_this).f10;
          let _lhs208 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          _lhs203[_lhs204] = _rhs328;
          _lhs205[_lhs206] = _rhs329;
          _lhs207[_lhs208] = _rhs330;
          _2108_v0 = _rhs331;
          let _index339 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index339] = (_dafny.Map.Empty.slice().updateUnsafe(true,(_2125_v14).f4)).contains((_module.D3.create_DC9(_module.__default.fm0(_2230_v103, globalState), (_this).f6, ((_this).f10)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length))])).dtor_cf8);
          let _2233_v105;
          _2233_v105 = _module.D12.create_DC27(_dafny.Seq.update(_2221_v96, _module.__default.safeIndex((_this).f6, new BigNumber((_2221_v96).length)), _2219_cf7));
          let _index340 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          let _index341 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          let _rhs332 = _dafny.Seq.Concat(_module.__default.fm39((_2226_v101)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_2226_v101).length))], _2233_v105, (_2125_v14).f4, globalState), _dafny.Seq.Concat(_dafny.Seq.of(_2108_v0), _2229_v102));
          let _rhs333 = (_2229_v102)[_module.__default.safeIndex((_this).f6, new BigNumber((_2229_v102).length))];
          let _rhs334 = (_2125_v14).f4;
          let _rhs335 = !(false);
          let _lhs209 = (_this).f10;
          let _lhs210 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          let _lhs211 = (_this).f10;
          let _lhs212 = _module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length));
          _2229_v102 = _rhs332;
          _2108_v0 = _rhs333;
          _lhs209[_lhs210] = _rhs334;
          _lhs211[_lhs212] = _rhs335;
          let _2234_v106;
          let _nw338 = new _module.C6();
          _nw338.__ctor(((_this).f10)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length))], (_this).f4, _this.f5, (_this).f6);
          _2234_v106 = _nw338;
          let _2235_v107;
          _2235_v107 = _dafny.Map.Empty.slice().updateUnsafe(_2234_v106,_2109_v1);
          let _2236_v108;
          _2236_v108 = _dafny.Map.Empty.slice().updateUnsafe((_2125_v14).f4,(_this).f6);
          let _2237_v109;
          _2237_v109 = _dafny.MultiSet.fromElements((_2234_v106).f6, (((_2236_v108).contains(!((_2110_v2).fm8(new BigNumber((_2219_cf7).length), (_2234_v106).f4, _2109_v1, _2229_v102, globalState)))) ? ((_2236_v108).get(!((_2110_v2).fm8(new BigNumber((_2219_cf7).length), (_2234_v106).f4, _2109_v1, _2229_v102, globalState)))) : ((_this).f6)));
          let _2238_v110;
          _2238_v110 = _dafny.Map.Empty.slice().updateUnsafe((_2234_v106).f6,_2109_v1);
          let _2239_v111;
          _2239_v111 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f10)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length))],(_2125_v14).f4);
          let _2240_v112;
          _2240_v112 = _dafny.MultiSet.fromElements(((_this).f10)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length))]);
          let _2241_v113;
          let _nw339 = Array((new BigNumber(26)).toNumber());
          _nw339[(_dafny.ZERO).toNumber()] = (_2110_v2).fm7(globalState);
          _nw339[(_dafny.ONE).toNumber()] = _2108_v0;
          _nw339[(new BigNumber(2)).toNumber()] = (_this).fm9(true, false, (_this).f6, globalState);
          _nw339[(new BigNumber(3)).toNumber()] = new BigNumber(((_2127_v16).Merge(_2127_v16)).length);
          _nw339[(new BigNumber(4)).toNumber()] = _2108_v0;
          _nw339[(new BigNumber(5)).toNumber()] = _2108_v0;
          _nw339[(new BigNumber(6)).toNumber()] = (new BigNumber((_2235_v107).length)).multipliedBy(new BigNumber((_2219_cf7).length));
          _nw339[(new BigNumber(7)).toNumber()] = _2108_v0;
          _nw339[(new BigNumber(8)).toNumber()] = ((_this).f6).plus(_2108_v0);
          _nw339[(new BigNumber(9)).toNumber()] = (_2234_v106).f6;
          _nw339[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2229_v102, _2229_v102)).length);
          _nw339[(new BigNumber(11)).toNumber()] = new BigNumber(((_2237_v109).Difference(_2237_v109)).cardinality());
          _nw339[(new BigNumber(12)).toNumber()] = (_2234_v106).f6;
          _nw339[(new BigNumber(13)).toNumber()] = (_this).f6;
          _nw339[(new BigNumber(14)).toNumber()] = (((_2234_v106).f4) ? (new BigNumber((_2238_v110).length)) : ((_this).f6));
          _nw339[(new BigNumber(15)).toNumber()] = _2108_v0;
          _nw339[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.of(false)).length);
          _nw339[(new BigNumber(17)).toNumber()] = (_this).f6;
          _nw339[(new BigNumber(18)).toNumber()] = (_2234_v106).f6;
          _nw339[(new BigNumber(19)).toNumber()] = (_this).f6;
          _nw339[(new BigNumber(20)).toNumber()] = (new BigNumber((_dafny.Set.fromElements(((_this).f10)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length))])).length)).multipliedBy((_this).f6);
          _nw339[(new BigNumber(21)).toNumber()] = _2108_v0;
          _nw339[(new BigNumber(22)).toNumber()] = new BigNumber((_2109_v1).length);
          _nw339[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.Concat((_2226_v101)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_2226_v101).length))], (_2226_v101)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_2226_v101).length))])).length);
          _nw339[(new BigNumber(24)).toNumber()] = new BigNumber((_2239_v111).length);
          _nw339[(new BigNumber(25)).toNumber()] = new BigNumber((_2240_v112).cardinality());
          _2241_v113 = _nw339;
          let _2242_v114;
          _2242_v114 = _dafny.Map.Empty.slice().updateUnsafe(_2127_v16,(_2125_v14).f4);
          let _index342 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_2241_v113).length));
          (_2241_v113)[_index342] = new BigNumber(((((_2125_v14).f4) ? (_dafny.Seq.of(new BigNumber((_2242_v114).length))) : (_2229_v102))).length);
          let _index343 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_2241_v113).length));
          (_2241_v113)[_index343] = _module.__default.safeDivisionInt(_2108_v0, (_this).f6);
          let _2243_v115;
          let _init56 = ((_2244_v106) => function (_2245_i7) {
            return (_2244_v106).f4;
          })(_2234_v106);
          let _nw340 = Array((_dafny.ONE).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw340.length); _i0_56++) {
            _nw340[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _2243_v115 = _nw340;
          let _2246_v116;
          _2246_v116 = _dafny.Map.Empty.slice().updateUnsafe(_2230_v103,_2243_v115);
          let _2247_v117;
          _2247_v117 = _module.D19.create_DC37(_2246_v116);
          let _2248_v118;
          _2248_v118 = _module.D19.create_DC39(_2247_v117);
          let _2249_v119;
          _2249_v119 = _dafny.Map.Empty.slice().updateUnsafe(_2248_v118,(_dafny.ZERO).minus((_this).f6));
          let _2250_v120;
          _2250_v120 = _dafny.Map.Empty.slice().updateUnsafe(_2230_v103,new BigNumber((_2249_v119).length));
          _2250_v120 = (_2250_v120).update(_2230_v103, new BigNumber(235));
        }
        let _2251_v121;
        let _nw341 = new _module.C1();
        _nw341.__ctor(new BigNumber(669), (_this).f6, _this.f5, new BigNumber((_dafny.Seq.UnicodeFromString("kv")).length), (_2125_v14).f4);
        _2251_v121 = _nw341;
        let _2252_v122;
        _2252_v122 = _dafny.Seq.of(_module.__default.fm57(_2108_v0, _2108_v0, globalState), _2229_v102, _dafny.Seq.Create(_module.__default.abs(new BigNumber(875)), function (_2253_i8) {
          return new BigNumber(246);
        }), _2229_v102);
        let _nw342 = new _module.C5();
        _nw342.__ctor((_2251_v121).f4, _module.__default.fm3(_module.__default.fm3((_this).f6, true, globalState), ((_this).f10)[_module.__default.safeIndex(new BigNumber(211), new BigNumber(((_this).f10).length))], globalState), _2251_v121.f5, new BigNumber((_2252_v122).length), p0);
        _2251_v121 = _nw342;
      } else {
        let _2254___mcc_h5 = (_source36).cf12;
        let _2255_cf12 = _2254___mcc_h5;
        let _2256_v123;
        _2256_v123 = _dafny.Map.Empty.slice().updateUnsafe(_2109_v1,new _dafny.CodePoint('a'.codePointAt(0)));
        let _2257_v124;
        _2257_v124 = _dafny.MultiSet.fromElements((_2125_v14).f4, true, (_2125_v14).f4);
        let _2258_v125;
        _2258_v125 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(),(((_2127_v16).contains(new BigNumber((_2256_v123).length))) ? ((_2127_v16).get(new BigNumber((_2256_v123).length))) : ((((_2257_v124).contains((_2125_v14).f4)) ? ((_2257_v124).get((_2125_v14).f4)) : ((_2110_v2).fm9(true, false, new BigNumber(632), globalState))))));
        let _2259_v126;
        _2259_v126 = _module.D2.create_DC6();
        _2258_v125 = (_2258_v125).update(_2259_v126, _module.__default.safeModuloInt((_this).f6, (_this).f6));
        let _2260_v127;
        let _init57 = ((_2261_v0) => function (_2262_i9) {
          return _module.__default.safeDivisionInt(_2262_i9, _2261_v0);
        })(_2108_v0);
        let _nw343 = Array((new BigNumber(25)).toNumber());
        for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw343.length); _i0_57++) {
          _nw343[_i0_57] = _init57(new BigNumber(_i0_57));
        }
        _2260_v127 = _nw343;
        let _2263_v128;
        _2263_v128 = _module.D16.create_DC32(_2260_v127);
        let _source37 = ((p0) ? (_2263_v128) : (_2263_v128));
        if (_source37.is_DC33) {
          let _2264___mcc_h6 = (_source37).cf52;
          let _2265___mcc_h7 = (_source37).cf53;
          let _2266___mcc_h8 = (_source37).cf54;
          let _2267___mcc_h9 = (_source37).cf55;
          let _2268___mcc_h10 = (_source37).cf56;
          let _2269_cf56 = _2268___mcc_h10;
          let _2270_cf55 = _2267___mcc_h9;
          let _2271_cf54 = _2266___mcc_h8;
          let _2272_cf53 = _2265___mcc_h7;
          let _2273_cf52 = _2264___mcc_h6;
          let _2274_v129;
          _2274_v129 = _dafny.Map.Empty.slice().updateUnsafe((_2125_v14).f4,_2270_cf55);
          _2108_v0 = _module.__default.safeModuloInt((_this).f6, (new BigNumber((_2274_v129).length)).plus(new BigNumber((_2109_v1).length)));
          let _2275_v130;
          _2275_v130 = new _dafny.CodePoint('f'.codePointAt(0));
          r0 = _2275_v130;
          let _2276_v131;
          _2276_v131 = _module.D4.create_DC12(new _dafny.CodePoint('p'.codePointAt(0)));
          _2276_v131 = _2276_v131;
          let _2277_v132;
          let _init58 = ((_2278_v0, _2279_cf55, _2280_v130) => function (_2281_i10) {
            return (_dafny.Set.fromElements(_dafny.Seq.of((_this).f6, _2278_v0, (_this).f6), _dafny.Seq.of(_2279_cf55, new BigNumber((_dafny.Seq.UnicodeFromString("dv")).length), new BigNumber((_dafny.Set.fromElements(_2280_v130, _2280_v130)).length)), _dafny.Seq.of(_2278_v0))).Difference(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(741))));
          })(_2108_v0, _2270_cf55, _2275_v130);
          let _nw344 = Array((new BigNumber(6)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw344.length); _i0_58++) {
            _nw344[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _2277_v132 = _nw344;
          let _2282_v133;
          _2282_v133 = _dafny.Seq.of((_dafny.ZERO).minus(_2270_cf55));
          let _2283_v134;
          _2283_v134 = _dafny.MultiSet.fromElements(_2125_v14);
          let _2284_v135;
          _2284_v135 = _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("tm")).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(283)), ((_2285_v0) => function (_2286_i11) {
            return _2285_v0;
          })(_2108_v0)), _2282_v133, _2282_v133, _dafny.Seq.of(new BigNumber(982), new BigNumber((_2283_v134).cardinality())));
          let _index344 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2277_v132).length));
          (_2277_v132)[_index344] = _2284_v135;
          let _index345 = _module.__default.safeIndex(new BigNumber(963), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index345] = _2273_cf52;
          let _2287_v136;
          _2287_v136 = _module.D6.create_DC16((_2125_v14).f4, (_this).f4, _2282_v133);
          let _2288_v137;
          _2288_v137 = _dafny.Set.fromElements(_2287_v136);
          let _2289_v138;
          _2289_v138 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm42(_2275_v130, _2288_v137, globalState),_2275_v130);
          let _index346 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_2260_v127).length));
          (_2260_v127)[_index346] = _2108_v0;
          let _2290_v139;
          _2290_v139 = _dafny.Seq.of(_2273_cf52, _2273_cf52);
          let _2291_v140;
          _2291_v140 = _dafny.Seq.of(_2290_v139, _2290_v139, _2290_v139);
          let _2292_v141;
          _2292_v141 = _module.D12.create_DC27(_2291_v140);
          let _2293_v142;
          _2293_v142 = _dafny.Map.Empty.slice().updateUnsafe((_2125_v14).f4,(_this).f4);
          let _2294_v143;
          _2294_v143 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(((_2293_v142).contains(p0)) ? ((_2293_v142).get(p0)) : ((_2125_v14).f4)));
          let _2295_v144;
          _2295_v144 = _module.D9.create_DC19(_dafny.Seq.UnicodeFromString("mskr"));
          let _index347 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2277_v132).length));
          let _index348 = _module.__default.safeIndex(new BigNumber(963), new BigNumber(((_this).f10).length));
          let _index349 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_2260_v127).length));
          let _rhs336 = _dafny.Set.fromElements(_dafny.Seq.Concat(_2282_v133, _module.__default.fm39(_2109_v1, _2292_v141, (((_2294_v143).contains(new BigNumber((_2282_v133).length))) ? ((_2294_v143).get(new BigNumber((_2282_v133).length))) : (p0)), globalState)));
          let _rhs337 = (_this).f4;
          let _rhs338 = (_2289_v138).Merge(_2289_v138);
          let _rhs339 = new BigNumber(((_2295_v144).dtor_cf25).length);
          let _lhs213 = _2277_v132;
          let _lhs214 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_2277_v132).length));
          let _lhs215 = (_this).f10;
          let _lhs216 = _module.__default.safeIndex(new BigNumber(963), new BigNumber(((_this).f10).length));
          let _lhs217 = _2260_v127;
          let _lhs218 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_2260_v127).length));
          _lhs213[_lhs214] = _rhs336;
          _lhs215[_lhs216] = _rhs337;
          _2289_v138 = _rhs338;
          _lhs217[_lhs218] = _rhs339;
        } else {
          let _2296___mcc_h11 = (_source37).cf51;
          let _2297_cf51 = _2296___mcc_h11;
          let _2298_v145;
          _2298_v145 = _module.D3.create_DC10((_this).f6);
          _2108_v0 = _module.__default.safeModuloInt((_2298_v145).dtor_cf11, new BigNumber(765));
          let _2299_v146;
          _2299_v146 = _dafny.Seq.of(false, (_2110_v2).fm8(_2108_v0, (_this).f4, _2109_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), function (_2300_i12) {
            return (_this).f6;
          }), globalState));
          let _index350 = _module.__default.safeIndex(new BigNumber(257), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index350] = (_2299_v146)[_module.__default.safeIndex((_this).f6, new BigNumber((_2299_v146).length))];
          let _2301_v147;
          let _init59 = function (_2302_i13) {
            return _module.D0.create_DC1();
          };
          let _nw345 = Array((new BigNumber(12)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw345.length); _i0_59++) {
            _nw345[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _2301_v147 = _nw345;
          let _2303_v148;
          _2303_v148 = true;
          let _2304_v149;
          _2304_v149 = new _dafny.CodePoint('p'.codePointAt(0));
          let _2305_v150;
          let _nw346 = Array((new BigNumber(10)).toNumber());
          _nw346[(_dafny.ZERO).toNumber()] = _2109_v1;
          _nw346[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("xkjwa");
          _nw346[(new BigNumber(2)).toNumber()] = _2109_v1;
          _nw346[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tc"), _2109_v1), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2125_v14).f4,_2304_v149)).length)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tc"), _2109_v1)).length)), new _dafny.CodePoint('c'.codePointAt(0)));
          _nw346[(new BigNumber(4)).toNumber()] = _2109_v1;
          _nw346[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("vtrm");
          _nw346[(new BigNumber(6)).toNumber()] = _2109_v1;
          _nw346[(new BigNumber(7)).toNumber()] = (((_2125_v14).f4) ? (_2109_v1) : (_2109_v1));
          _nw346[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(471)), function (_2306_i14) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          });
          _nw346[(new BigNumber(9)).toNumber()] = _2109_v1;
          _2305_v150 = _nw346;
          let _index351 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_2305_v150).length));
          (_2305_v150)[_index351] = (((_this).f4) ? (_2109_v1) : (_2109_v1));
          let _index352 = _module.__default.safeIndex(new BigNumber(257), new BigNumber(((_this).f10).length));
          let _index353 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_2305_v150).length));
          let _rhs340 = (_2125_v14).f4;
          let _rhs341 = (_module.D21.create_DC42(_2301_v147)).dtor_cf68;
          let _rhs342 = (_2125_v14).f4;
          let _rhs343 = _2108_v0;
          let _rhs344 = _dafny.Seq.Concat(_2109_v1, _2109_v1);
          let _lhs219 = (_this).f10;
          let _lhs220 = _module.__default.safeIndex(new BigNumber(257), new BigNumber(((_this).f10).length));
          let _lhs221 = _2305_v150;
          let _lhs222 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_2305_v150).length));
          _lhs219[_lhs220] = _rhs340;
          _2301_v147 = _rhs341;
          _2303_v148 = _rhs342;
          _2108_v0 = _rhs343;
          _lhs221[_lhs222] = _rhs344;
          let _2307_v151;
          _2307_v151 = _dafny.Map.Empty.slice().updateUnsafe(false,((_this).f10)[_module.__default.safeIndex(new BigNumber(257), new BigNumber(((_this).f10).length))]);
          let _2308_v152;
          _2308_v152 = _dafny.Set.fromElements((_this).f6);
          let _2309_v153;
          _2309_v153 = _dafny.Map.Empty.slice().updateUnsafe(_2307_v151,new BigNumber((_2308_v152).length));
          let _rhs345 = _module.__default.fm0(new _dafny.CodePoint('o'.codePointAt(0)), globalState);
          let _rhs346 = !((_2309_v153).equals(_2309_v153));
          _2303_v148 = _rhs345;
          _2303_v148 = _rhs346;
          _2108_v0 = _2108_v0;
        }
        let _2310_v154;
        _2310_v154 = _dafny.Map.Empty.slice().updateUnsafe(_2109_v1,new BigNumber((_2109_v1).length));
        let _2311_v155;
        _2311_v155 = _dafny.Map.Empty.slice().updateUnsafe(_2310_v154,_2108_v0);
        _2311_v155 = (_2311_v155).update(_dafny.Map.Empty.slice().updateUnsafe(_2109_v1,(_dafny.ZERO).minus((_this).f6)), (_this).f6);
        let _2312_v156;
        _2312_v156 = _dafny.Seq.of(_2108_v0);
        let _2313_v157;
        _2313_v157 = _dafny.Set.fromElements(p0);
        _2108_v0 = (_2312_v156)[_module.__default.safeIndex(new BigNumber((_2313_v157).length), new BigNumber((_2312_v156).length))];
      }
      let _2314_v158;
      _2314_v158 = true;
      _2314_v158 = (_this).f4;
      let _2315_v159;
      _2315_v159 = new _dafny.CodePoint('k'.codePointAt(0));
      r0 = _2315_v159;
      _2108_v0 = _2108_v0;
      r0 = _2315_v159;
      let _2316_v160;
      let _nw347 = Array((new BigNumber(19)).toNumber()).fill(_module.D1.Default());
      _2316_v160 = _nw347;
      r1 = _2316_v160;
      return [r0, r1];
    }
    m6(globalState) {
      let _this = this;
      let r0 = _module.D1.Default();
      let r1 = _dafny.Seq.UnicodeFromString("");
      let _2317_i0;
      _2317_i0 = _dafny.ZERO;
      L10: {
        while ((_this).f4) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2317_i0)) {
              break L10;
            }
            _2317_i0 = (_2317_i0).plus(_dafny.ONE);
            let _2318_v0;
            let _nw348 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
            _2318_v0 = _nw348;
            let _2319_v1;
            _2319_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f4);
            let _index354 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_2318_v0).length));
            (_2318_v0)[_index354] = (_2319_v1).Merge(_2319_v1);
            let _index355 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_2318_v0).length));
            (_2318_v0)[_index355] = (_2319_v1).Merge(((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f4)).update((_this).f10, (_this).f4)).Merge(_2319_v1));
            let _2320_v2;
            let _nw349 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
            _2320_v2 = _nw349;
            let _index356 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_2320_v2).length));
            (_2320_v2)[_index356] = (_this).f6;
            let _2321_v3;
            _2321_v3 = new _dafny.CodePoint('w'.codePointAt(0));
            let _index357 = _module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f10).length));
            ((_this).f10)[_index357] = _module.__default.fm0(_2321_v3, globalState);
            let _2322_v4;
            _2322_v4 = true;
            let _index358 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_2320_v2).length));
            let _index359 = _module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f10).length));
            let _rhs347 = (_this).f6;
            let _rhs348 = (_this).f4;
            let _rhs349 = _2322_v4;
            let _lhs223 = _2320_v2;
            let _lhs224 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_2320_v2).length));
            let _lhs225 = (_this).f10;
            let _lhs226 = _module.__default.safeIndex(new BigNumber(597), new BigNumber(((_this).f10).length));
            _lhs223[_lhs224] = _rhs347;
            _lhs225[_lhs226] = _rhs348;
            _2322_v4 = _rhs349;
            let _2323_v5;
            let _init60 = ((_2324_v2) => function (_2325_i1) {
              return (_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_2324_v2)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_2324_v2).length))])).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6)));
            })(_2320_v2);
            let _nw350 = Array((new BigNumber(13)).toNumber());
            for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw350.length); _i0_60++) {
              _nw350[_i0_60] = _init60(new BigNumber(_i0_60));
            }
            _2323_v5 = _nw350;
            let _2326_v6;
            _2326_v6 = _dafny.Seq.of((_module.D1.create_DC4((_2320_v2)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_2320_v2).length))], (_2320_v2)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_2320_v2).length))])).dtor_cf3);
            let _2327_v7;
            _2327_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2326_v6).length),(_this).f6);
            let _index360 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2323_v5).length));
            (_2323_v5)[_index360] = _2327_v7;
            let _index361 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2323_v5).length));
            (_2323_v5)[_index361] = _2327_v7;
            let _2328_v8;
            let _nw351 = new _module.C6();
            _nw351.__ctor(_module.__default.fm0(_2321_v3, globalState), false, _this.f5, _module.__default.safeModuloInt((_this).f6, (_this).f6));
            _2328_v8 = _nw351;
          }
        }
      }
      let _hi13 = (_this).f6;
      for (let _2329_i2 = (_this).f6; _2329_i2.isLessThan(_hi13); _2329_i2 = _2329_i2.plus(_dafny.ONE)) {
        let _index362 = _module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index362] = (false) || (true);
        let _index363 = _module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index363] = (_2329_i2).isLessThanOrEqualTo((_this).f6);
        let _index364 = _module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index364] = ((_this).f10)[_module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length))];
        let _2330_v9;
        let _init61 = function (_2331_i3) {
          return _dafny.Map.Empty.slice().updateUnsafe(true,((_this).f10)[_module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length))]);
        };
        let _nw352 = Array((new BigNumber(24)).toNumber());
        for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw352.length); _i0_61++) {
          _nw352[_i0_61] = _init61(new BigNumber(_i0_61));
        }
        _2330_v9 = _nw352;
        let _2332_v10;
        _2332_v10 = _module.D11.create_DC23(_2330_v9);
        let _2333_v11;
        _2333_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2332_v10);
        let _2334_v12;
        _2334_v12 = new BigNumber(472);
        let _rhs350 = _2333_v11;
        let _rhs351 = _2329_i2;
        _2333_v11 = _rhs350;
        _2334_v12 = _rhs351;
        let _index365 = _module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length));
        let _index366 = _module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length));
        let _rhs352 = !(((_this).f10)[_module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length))]);
        let _rhs353 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_2329_i2).multipliedBy(new BigNumber(697))));
        let _rhs354 = ((_dafny.ZERO).minus(_2334_v12)).minus((_this).fm7(globalState));
        let _rhs355 = false;
        let _lhs227 = (_this).f10;
        let _lhs228 = _module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length));
        let _lhs229 = (_this).f10;
        let _lhs230 = _module.__default.safeIndex(new BigNumber(324), new BigNumber(((_this).f10).length));
        _lhs227[_lhs228] = _rhs352;
        _2334_v12 = _rhs353;
        _2334_v12 = _rhs354;
        _lhs229[_lhs230] = _rhs355;
      }
      let _2335_v13;
      _2335_v13 = false;
      _2335_v13 = _2335_v13;
      let _index367 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
      ((_this).f10)[_index367] = (_2335_v13) || ((_this).f4);
      let _index368 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
      ((_this).f10)[_index368] = _2335_v13;
      let _2336_v14;
      _2336_v14 = _dafny.Seq.UnicodeFromString("ptmrel");
      let _2337_v15;
      _2337_v15 = _dafny.Seq.of(new BigNumber((((((_this).f9).contains((_this).f4)) ? (((_this).f9).get((_this).f4)) : (_2336_v14))).length), (_this).f6, (_this).f6, (_this).f6);
      let _2338_v16;
      _2338_v16 = _dafny.Seq.of(((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))]);
      let _2339_v17;
      _2339_v17 = _dafny.MultiSet.fromElements(_2338_v16);
      let _2340_v18;
      _2340_v18 = _dafny.Map.Empty.slice().updateUnsafe(true,_module.D6.create_DC15(_2339_v17));
      if ((new BigNumber((_2340_v18).length)).isLessThan((_2337_v15)[_module.__default.safeIndex((_this).f6, new BigNumber((_2337_v15).length))])) {
        let _2341_v19;
        _2341_v19 = _dafny.Set.fromElements((_this).f6);
        let _2342_v20;
        _2342_v20 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))],new BigNumber((_2341_v19).length));
        let _2343_v21;
        _2343_v21 = _module.D10.create_DC21(_2342_v20);
        let _source38 = _2343_v21;
        if (_source38.is_DC22) {
          let _2344___mcc_h0 = (_source38).cf30;
          let _2345___mcc_h1 = (_source38).cf31;
          let _2346___mcc_h2 = (_source38).cf32;
          let _2347___mcc_h3 = (_source38).cf33;
          let _2348___mcc_h4 = (_source38).cf34;
          let _2349_cf34 = _2348___mcc_h4;
          let _2350_cf33 = _2347___mcc_h3;
          let _2351_cf32 = _2346___mcc_h2;
          let _2352_cf31 = _2345___mcc_h1;
          let _2353_cf30 = _2344___mcc_h0;
          _2353_cf30 = _dafny.Seq.Concat(_2353_cf30, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mlwyh"), _2353_cf30));
          let _2354_v22;
          let _init62 = function (_2355_i4) {
            return (_2355_i4).minus(new BigNumber(199));
          };
          let _nw353 = Array((new BigNumber(28)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw353.length); _i0_62++) {
            _nw353[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _2354_v22 = _nw353;
          let _2356_v23;
          _2356_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2337_v15).length),(_this).f6);
          let _index369 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_2354_v22).length));
          (_2354_v22)[_index369] = (((_2356_v23).contains(new BigNumber((_2336_v14).length))) ? ((_2356_v23).get(new BigNumber((_2336_v14).length))) : (_2349_cf34));
          let _index370 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_2354_v22).length));
          (_2354_v22)[_index370] = (_this).f6;
          let _2357_v24;
          _2357_v24 = _module.D11.create_DC25(_2352_cf31, (_this).f6);
          let _pat_let_tv40 = _2335_v13;
          let _2358_v25;
          let _nw354 = new _module.C5();
          _nw354.__ctor((function (_pat_let51_0) {
            return function (_2359_dt__update__tmp_h0) {
              return function (_pat_let52_0) {
                return function (_2360_dt__update_hcf39_h0) {
                  return _module.D11.create_DC25(_2360_dt__update_hcf39_h0, (_2359_dt__update__tmp_h0).dtor_cf40);
                }(_pat_let52_0);
              }(_pat_let_tv40);
            }(_pat_let51_0);
          }(_2357_v24)).dtor_cf39, (_dafny.ZERO).minus((_this).f6), _this.f5, new BigNumber((_dafny.Seq.UnicodeFromString("romg")).length), ((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))]);
          _2358_v25 = _nw354;
          let _2361_v26;
          let _init63 = ((_2362_cf32) => function (_2363_i5) {
            return _dafny.MultiSet.fromElements(_2362_cf32);
          })(_2351_cf32);
          let _nw355 = Array((new BigNumber(18)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw355.length); _i0_63++) {
            _nw355[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _2361_v26 = _nw355;
          _2361_v26 = _2361_v26;
        } else {
          let _2364___mcc_h5 = (_source38).cf29;
          let _2365_cf29 = _2364___mcc_h5;
          let _2366_v27;
          _2366_v27 = new BigNumber(371);
          _2366_v27 = _2366_v27;
          let _2367_v28;
          _2367_v28 = _dafny.MultiSet.fromElements(((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))], (_this).f4, (_this).f4, _2335_v13);
          let _2368_v29;
          _2368_v29 = _dafny.MultiSet.fromElements(_2366_v27, new BigNumber((_2367_v28).cardinality()));
          let _2369_v30;
          _2369_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2368_v29,_2336_v14);
          let _2370_v31;
          let _2371_v32;
          let _out62;
          let _out63;
          let _outcollector27 = _module.__default.m0(_2369_v30, ((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))], _2368_v29, (new BigNumber((_dafny.Seq.update(_2337_v15, _module.__default.safeIndex(_2366_v27, new BigNumber((_2337_v15).length)), _2366_v27)).length)).isLessThan(new BigNumber(253)), globalState);
          _out62 = _outcollector27[0];
          _out63 = _outcollector27[1];
          _2370_v31 = _out62;
          _2371_v32 = _out63;
          _2366_v27 = (_this).f6;
          let _index371 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
          ((_this).f10)[_index371] = (_this).f4;
        }
        if (_2335_v13) {
          let _2372_v33;
          _2372_v33 = _dafny.MultiSet.fromElements((_this).f6);
          let _2373_v34;
          _2373_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2372_v33,_2336_v14);
          let _2374_v35;
          let _2375_v36;
          let _out64;
          let _out65;
          let _outcollector28 = _module.__default.m0(_2373_v34, true, _dafny.MultiSet.FromArray(_2337_v15), _2335_v13, globalState);
          _out64 = _outcollector28[0];
          _out65 = _outcollector28[1];
          _2374_v35 = _out64;
          _2375_v36 = _out65;
          let _2376_v37;
          let _init64 = ((_2377_v14) => function (_2378_i6) {
            return (_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(929)), function (_2379_i7) {
              return new _dafny.CodePoint('a'.codePointAt(0));
            }))).Union(_dafny.Set.fromElements(_2377_v14, _dafny.Seq.UnicodeFromString("mj")));
          })(_2336_v14);
          let _nw356 = Array((new BigNumber(6)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw356.length); _i0_64++) {
            _nw356[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2376_v37 = _nw356;
          let _2380_v38;
          _2380_v38 = _dafny.Set.fromElements(_2336_v14);
          let _index372 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_2376_v37).length));
          (_2376_v37)[_index372] = _2380_v38;
          let _index373 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_2376_v37).length));
          (_2376_v37)[_index373] = _2380_v38;
          _2375_v36 = _2375_v36;
          let _2381_v39;
          let _nw357 = new _module.C4();
          _nw357.__ctor(_this.f5, (_2337_v15)[_module.__default.safeIndex((_this).f6, new BigNumber((_2337_v15).length))], (_this).f4);
          _2381_v39 = _nw357;
          let _2382_v40;
          _2382_v40 = _module.D4.create_DC13(_2375_v36, _2374_v35, new BigNumber((_dafny.Seq.UnicodeFromString("wdwcp")).length), _2338_v16);
          let _2383_v41;
          _2383_v41 = new _dafny.CodePoint('i'.codePointAt(0));
          let _rhs356 = ((_this).f6).minus((_2382_v40).dtor_cf14);
          let _rhs357 = _2335_v13;
          let _rhs358 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_2336_v14, _module.__default.safeIndex(_2375_v36, new BigNumber((_2336_v14).length)), _2383_v41), _dafny.Seq.Concat(_dafny.Seq.update(_2336_v14, _module.__default.safeIndex((_this).f6, new BigNumber((_2336_v14).length)), _2383_v41), _2336_v14))).length);
          let _rhs359 = _2375_v36;
          let _rhs360 = (_2375_v36).plus(new BigNumber(439));
          _2375_v36 = _rhs356;
          _2374_v35 = _rhs357;
          _2375_v36 = _rhs358;
          _2375_v36 = _rhs359;
          _2375_v36 = _rhs360;
        } else {
          let _2384_v42;
          let _init65 = function (_2385_i8) {
            return (_2385_i8).minus((_this).f6);
          };
          let _nw358 = Array((new BigNumber(26)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw358.length); _i0_65++) {
            _nw358[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2384_v42 = _nw358;
          let _index374 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_2384_v42).length));
          (_2384_v42)[_index374] = ((_this).f6).minus((_this).f6);
          let _index375 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_2384_v42).length));
          (_2384_v42)[_index375] = (_this).f6;
          let _index376 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_2384_v42).length));
          (_2384_v42)[_index376] = ((_2384_v42)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_2384_v42).length))]).multipliedBy((_2384_v42)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_2384_v42).length))]);
          let _2386_v44;
          _2386_v44 = _dafny.MultiSet.FromArray(_2337_v15);
          let _2387_v45;
          _2387_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2386_v44,((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))]);
          let _2388_v46;
          _2388_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (_2387_v45).Keys.Elements) {
              let _2389_v43 = _compr_59;
              if ((_2387_v45).contains(_2389_v43)) {
                _coll59.push([_2389_v43,new BigNumber((_2336_v14).length)]);
              }
            }
            return _coll59;
          }()).length),(_this).f4);
          _2388_v46 = _2388_v46;
          _2335_v13 = _2335_v13;
          let _index377 = _module.__default.safeIndex(new BigNumber(291), new BigNumber((_2384_v42).length));
          (_2384_v42)[_index377] = (_this).f6;
        }
        _2337_v15 = _dafny.Seq.of(_module.__default.safeModuloInt((_this).f6, new BigNumber((_2342_v20).length)), (_this).f6);
        let _2390_v47;
        _2390_v47 = new BigNumber(355);
        _2390_v47 = ((_2335_v13) ? ((_this).f6) : (_2390_v47));
        let _index378 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index378] = ((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))];
      } else {
        let _2391_v48;
        let _nw359 = Array((_dafny.ONE).toNumber()).fill(_module.D21.Default());
        _2391_v48 = _nw359;
        let _2392_v49;
        let _init66 = function (_2393_i9) {
          return _module.D0.create_DC1();
        };
        let _nw360 = Array((_dafny.ONE).toNumber());
        for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw360.length); _i0_66++) {
          _nw360[_i0_66] = _init66(new BigNumber(_i0_66));
        }
        _2392_v49 = _nw360;
        let _2394_v50;
        _2394_v50 = _module.D21.create_DC42(_2392_v49);
        let _pat_let_tv41 = _2392_v49;
        let _index379 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_2391_v48).length));
        (_2391_v48)[_index379] = function (_pat_let53_0) {
          return function (_2395_dt__update__tmp_h1) {
            return function (_pat_let54_0) {
              return function (_2396_dt__update_hcf68_h0) {
                return _module.D21.create_DC42(_2396_dt__update_hcf68_h0);
              }(_pat_let54_0);
            }(_pat_let_tv41);
          }(_pat_let53_0);
        }(_2394_v50);
        let _index380 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_2391_v48).length));
        (_2391_v48)[_index380] = _2394_v50;
        let _2397_v51;
        _2397_v51 = new _dafny.CodePoint('a'.codePointAt(0));
        _2397_v51 = _2397_v51;
        let _2398_v52;
        _2398_v52 = _dafny.Set.fromElements(true);
        let _index381 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index381] = (_2398_v52).IsSubsetOf(_2398_v52);
        let _2399_v53;
        _2399_v53 = _dafny.Seq.of(_2337_v15, _2337_v15, _2337_v15);
        let _2400_v54;
        _2400_v54 = _dafny.MultiSet.fromElements((_this).f6, (_this).f6, new BigNumber((_2399_v53).length));
        let _2401_v55;
        _2401_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2400_v54).cardinality()),((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))]);
        let _index382 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index382] = (_2335_v13) === ((_this).fm14(_2401_v55, !((_this).f4), (_this).f6, ((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))], globalState));
        let _2402_v56;
        let _init67 = function (_2403_i10) {
          return _module.__default.safeDivisionInt(_2403_i10, (_this).f6);
        };
        let _nw361 = Array((new BigNumber(13)).toNumber());
        for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw361.length); _i0_67++) {
          _nw361[_i0_67] = _init67(new BigNumber(_i0_67));
        }
        _2402_v56 = _nw361;
        let _index383 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_2402_v56).length));
        (_2402_v56)[_index383] = ((_this).f6).plus(_module.__default.fm3((_this).f6, true, globalState));
        let _index384 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_2402_v56).length));
        (_2402_v56)[_index384] = (_this).f6;
      }
      let _2404_v57;
      _2404_v57 = new _dafny.CodePoint('n'.codePointAt(0));
      if (_module.__default.fm0(_2404_v57, globalState)) {
        let _index385 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index385] = _2335_v13;
        let _2405_v58;
        let _nw362 = new _module.C6();
        _nw362.__ctor((_this).f4, _2335_v13, _this.f5, ((false) ? ((_this).f6) : ((_this).f6)));
        _2405_v58 = _nw362;
        let _2406_v59;
        _2406_v59 = new BigNumber(-792);
        let _index386 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
        let _rhs361 = _2405_v58;
        let _rhs362 = ((_this).fm9((_2405_v58).f11, ((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))], new BigNumber((_2336_v14).length), globalState)).isEqualTo(_2406_v59);
        let _rhs363 = _2406_v59;
        let _lhs231 = (_this).f10;
        let _lhs232 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
        _2405_v58 = _rhs361;
        _lhs231[_lhs232] = _rhs362;
        _2406_v59 = _rhs363;
        _2406_v59 = (_this).f6;
        _2406_v59 = _2406_v59;
        let _2407_v60;
        let _init68 = ((_2408_v59) => function (_2409_i11) {
          return (_2409_i11).multipliedBy(_2408_v59);
        })(_2406_v59);
        let _nw363 = Array((new BigNumber(15)).toNumber());
        for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw363.length); _i0_68++) {
          _nw363[_i0_68] = _init68(new BigNumber(_i0_68));
        }
        _2407_v60 = _nw363;
        let _index387 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_2407_v60).length));
        (_2407_v60)[_index387] = _module.__default.safeDivisionInt(_2406_v59, (_this).f6);
        let _2410_v61;
        _2410_v61 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2406_v59));
        let _2411_v62;
        _2411_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2410_v61,_2335_v13);
        let _2412_v63;
        _2412_v63 = _2410_v61;
        let _2413_v64;
        _2413_v64 = _dafny.Map.Empty.slice().updateUnsafe(_2406_v59,(_2412_v63));
        let _2414_v65;
        _2414_v65 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f10)[_module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length))],(((_2411_v62).contains((((_2413_v64).contains(_2406_v59)) ? ((_2413_v64).get(_2406_v59)) : (_2410_v61)))) ? ((_2411_v62).get((((_2413_v64).contains(_2406_v59)) ? ((_2413_v64).get(_2406_v59)) : (_2410_v61)))) : ((_2405_v58).f11)));
        let _index388 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_2407_v60).length));
        let _rhs364 = (_dafny.ZERO).minus((_2406_v59).minus((new BigNumber((_2337_v15).length)).multipliedBy(_2406_v59)));
        let _rhs365 = (((_2414_v65).contains((_2405_v58).f11)) ? ((_2414_v65).get((_2405_v58).f11)) : ((_2405_v58).f11));
        let _rhs366 = (_this).fm7(globalState);
        let _rhs367 = _2405_v58;
        let _lhs233 = _2407_v60;
        let _lhs234 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_2407_v60).length));
        _lhs233[_lhs234] = _rhs364;
        _2335_v13 = _rhs365;
        _2406_v59 = _rhs366;
        _2405_v58 = _rhs367;
      } else {
        let _2415_v66;
        _2415_v66 = new BigNumber(313);
        _2415_v66 = _module.__default.safeDivisionInt(_2415_v66, ((_this).f6).plus(new BigNumber((_dafny.Seq.UnicodeFromString("gpm")).length)));
        let _2416_v67;
        let _nw364 = Array((new BigNumber(27)).toNumber()).fill(false);
        _2416_v67 = _nw364;
        _2416_v67 = _2416_v67;
        let _2417_v68;
        _2417_v68 = _module.D9.create_DC19(_dafny.Seq.UnicodeFromString("p"));
        let _2418_v69;
        _2418_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2417_v68,(_this).f4);
        let _2419_v70;
        _2419_v70 = _dafny.Map.Empty.slice().updateUnsafe((_2418_v69).Merge(_2418_v69),_2404_v57);
        _2419_v70 = (_2419_v70).update(_2418_v69, _2404_v57);
        _2415_v66 = (_dafny.ZERO).minus((_2415_v66).plus(_2415_v66));
        let _index389 = _module.__default.safeIndex(new BigNumber(304), new BigNumber(((_this).f10).length));
        ((_this).f10)[_index389] = (((_this).f6).isEqualTo((_this).f6)) && (!(!_dafny.areEqual(_2337_v15, _2337_v15)));
      }
      let _2420_v71;
      _2420_v71 = _dafny.Set.fromElements((_this).f6);
      let _2421_v72;
      _2421_v72 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f6),!(!(_2335_v13)));
      let _2422_v73;
      let _nw365 = Array((new BigNumber(28)).toNumber());
      _nw365[(_dafny.ZERO).toNumber()] = (_this).f6;
      _nw365[(_dafny.ONE).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(2)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(3)).toNumber()] = new BigNumber((_2420_v71).length);
      _nw365[(new BigNumber(4)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(5)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(6)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(7)).toNumber()] = new BigNumber(617);
      _nw365[(new BigNumber(8)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(9)).toNumber()] = new BigNumber(844);
      _nw365[(new BigNumber(10)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(11)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(12)).toNumber()] = new BigNumber(-957);
      _nw365[(new BigNumber(13)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(14)).toNumber()] = new BigNumber(860);
      _nw365[(new BigNumber(15)).toNumber()] = new BigNumber((_2336_v14).length);
      _nw365[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus((_this).f6);
      _nw365[(new BigNumber(17)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(18)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(19)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(20)).toNumber()] = new BigNumber((_2336_v14).length);
      _nw365[(new BigNumber(21)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(22)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(23)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(24)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(25)).toNumber()] = (_this).f6;
      _nw365[(new BigNumber(26)).toNumber()] = new BigNumber((_2421_v72).length);
      _nw365[(new BigNumber(27)).toNumber()] = new BigNumber(682);
      _2422_v73 = _nw365;
      let _2423_v74;
      _2423_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2422_v73,(_this).f6);
      let _pat_let_tv42 = _2423_v74;
      r0 = function (_pat_let55_0) {
        return function (_2426_dt__update__tmp_h2) {
          return function (_pat_let56_0) {
            return function (_2427_dt__update_hcf2_h0) {
              return _module.D1.create_DC3(_2427_dt__update_hcf2_h0);
            }(_pat_let56_0);
          }((new BigNumber((_pat_let_tv42).length)).isLessThanOrEqualTo((_this).f6));
        }(_pat_let55_0);
      }(_module.D1.create_DC3((_this).fm8((_this).f6, _2335_v13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-493)), ((_2424_v57) => function (_2425_i12) {
  return _2424_v57;
})(_2404_v57)), _2337_v15, globalState)));
      r1 = _2336_v14;
      return [r0, r1];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
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
    fm11(globalState) {
      let _this = this;
      return _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("prxjdttbg"), _dafny.Seq.UnicodeFromString("or"));
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(157)), function (_2428_i0) {
        return new BigNumber(-632);
      }), _dafny.Seq.of(new BigNumber(-337), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(854),new BigNumber((_dafny.Seq.of(true, false)).length)))).length));
    };
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let _2429_v0;
      _2429_v0 = _module.D0.create_DC0(new BigNumber(986));
      let _2430_i0;
      _2430_i0 = _dafny.ZERO;
      L11: {
        while (!((_this).fm12(p2, _2429_v0, p1, p2, globalState))) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2430_i0)) {
              break L11;
            }
            _2430_i0 = (_2430_i0).plus(_dafny.ONE);
            let _2431_v1;
            _2431_v1 = new _dafny.CodePoint('h'.codePointAt(0));
            _2431_v1 = _2431_v1;
            let _2432_v2;
            _2432_v2 = _dafny.Seq.UnicodeFromString("mhlv");
            _2432_v2 = _dafny.Seq.Concat(_2432_v2, _2432_v2);
            let _2433_v3;
            _2433_v3 = true;
            let _2434_v4;
            let _nw366 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
            _2434_v4 = _nw366;
            let _rhs368 = p2;
            let _rhs369 = _2434_v4;
            _2433_v3 = _rhs368;
            _2434_v4 = _rhs369;
            let _2435_v5;
            _2435_v5 = new BigNumber(337);
            let _rhs370 = (new BigNumber((_2432_v2).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(p0));
            let _rhs371 = new BigNumber((_2432_v2).length);
            _2433_v3 = _rhs370;
            _2435_v5 = _rhs371;
          }
        }
      }
      let _2436_v6;
      _2436_v6 = new BigNumber(631);
      _2436_v6 = (p1).multipliedBy(p0);
      let _2437_v7;
      let _nw367 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _2437_v7 = _nw367;
      let _index390 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_2437_v7).length));
      (_2437_v7)[_index390] = ((true) ? (new BigNumber(194)) : (p1));
      let _2438_v8;
      _2438_v8 = _module.D11.create_DC26(p2);
      let _2439_v9;
      _2439_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      let _2440_v10;
      _2440_v10 = _dafny.Seq.UnicodeFromString("drtvj");
      let _2441_v11;
      _2441_v11 = _module.D12.create_DC28(p0, new BigNumber((_2440_v10).length), p2, !(p2), p0);
      let _2442_v12;
      _2442_v12 = new _dafny.CodePoint('a'.codePointAt(0));
      let _2443_v13;
      _2443_v13 = _dafny.Seq.of(p2);
      let _2444_v14;
      _2444_v14 = _module.D21.create_DC43(p2, false, _2443_v13);
      let _2445_v15;
      _2445_v15 = _dafny.Seq.of(p2, (_2444_v14).dtor_cf70, p2);
      let _2446_v16;
      let _nw368 = Array((new BigNumber(21)).toNumber());
      _nw368[(_dafny.ZERO).toNumber()] = p2;
      _nw368[(_dafny.ONE).toNumber()] = p2;
      _nw368[(new BigNumber(2)).toNumber()] = true;
      _nw368[(new BigNumber(3)).toNumber()] = p2;
      _nw368[(new BigNumber(4)).toNumber()] = p2;
      _nw368[(new BigNumber(5)).toNumber()] = (_2438_v8).dtor_cf41;
      _nw368[(new BigNumber(6)).toNumber()] = p2;
      _nw368[(new BigNumber(7)).toNumber()] = p2;
      _nw368[(new BigNumber(8)).toNumber()] = p2;
      _nw368[(new BigNumber(9)).toNumber()] = p2;
      _nw368[(new BigNumber(10)).toNumber()] = p2;
      _nw368[(new BigNumber(11)).toNumber()] = p2;
      _nw368[(new BigNumber(12)).toNumber()] = p2;
      _nw368[(new BigNumber(13)).toNumber()] = p2;
      _nw368[(new BigNumber(14)).toNumber()] = p2;
      _nw368[(new BigNumber(15)).toNumber()] = p2;
      _nw368[(new BigNumber(16)).toNumber()] = (((_2439_v9).contains(p2)) ? ((_2439_v9).get(p2)) : (true));
      _nw368[(new BigNumber(17)).toNumber()] = p2;
      _nw368[(new BigNumber(18)).toNumber()] = (_module.D9.create_DC20((_2441_v11).dtor_cf46, _2442_v12, _2436_v6)).dtor_cf26;
      _nw368[(new BigNumber(19)).toNumber()] = !((_2443_v13)[_module.__default.safeIndex(p0, new BigNumber((_2443_v13).length))]);
      _nw368[(new BigNumber(20)).toNumber()] = p2;
      _2446_v16 = _nw368;
      let _2447_v17;
      let _nw369 = Array((new BigNumber(28)).toNumber());
      _nw369[(_dafny.ZERO).toNumber()] = _2446_v16;
      _nw369[(_dafny.ONE).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(2)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(3)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(4)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(5)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(6)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(7)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(8)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(9)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(10)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(11)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(12)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(13)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(14)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(15)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(16)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(17)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(18)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(19)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(20)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(21)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(22)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(23)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(24)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(25)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(26)).toNumber()] = _2446_v16;
      _nw369[(new BigNumber(27)).toNumber()] = _2446_v16;
      _2447_v17 = _nw369;
      let _2448_v18;
      _2448_v18 = _dafny.MultiSet.fromElements(new BigNumber(-502));
      let _2449_v19;
      let _nw370 = new _module.C3();
      _nw370.__ctor(_2447_v17, (_2436_v6).multipliedBy(p0), (_2448_v18).IsProperSubsetOf(_2448_v18));
      _2449_v19 = _nw370;
      let _arr13 = _2449_v19.f5;
      let _index391 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_2449_v19.f5).length));
      _arr13[_index391] = _2446_v16;
      let _2450_v20;
      _2450_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2436_v6);
      let _2451_v21;
      _2451_v21 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_2445_v15, _module.__default.safeIndex(new BigNumber((_2450_v20).length), new BigNumber((_2445_v15).length)), (_2449_v19).f4));
      let _2452_v22;
      _2452_v22 = _module.D6.create_DC15(_2451_v21);
      let _pat_let_tv43 = _2442_v12;
      let _pat_let_tv44 = _2440_v10;
      let _pat_let_tv45 = _2436_v6;
      let _pat_let_tv46 = _2440_v10;
      let _index392 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_2437_v7).length));
      let _arr14 = _2449_v19.f5;
      let _index393 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_2449_v19.f5).length));
      let _rhs372 = p0;
      let _rhs373 = _2449_v19;
      let _rhs374 = function (_source39) {
        if (_source39.is_DC16) {
          let _2453___mcc_h0 = (_source39).cf20;
          let _2454___mcc_h1 = (_source39).cf21;
          let _2455___mcc_h2 = (_source39).cf22;
          let _2456_cf22 = _2455___mcc_h2;
          let _2457_cf21 = _2454___mcc_h1;
          let _2458_cf20 = _2453___mcc_h0;
          return _pat_let_tv43;
        } else {
          let _2459___mcc_h3 = (_source39).cf19;
          let _2460_cf19 = _2459___mcc_h3;
          return (_pat_let_tv44)[_module.__default.safeIndex(_pat_let_tv45, new BigNumber((_pat_let_tv46).length))];
        }
      }(_2452_v22);
      let _rhs375 = _2446_v16;
      let _lhs235 = _2437_v7;
      let _lhs236 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_2437_v7).length));
      let _lhs237 = _2449_v19.f5;
      let _lhs238 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_2449_v19.f5).length));
      _lhs235[_lhs236] = _rhs372;
      _2449_v19 = _rhs373;
      _2442_v12 = _rhs374;
      _lhs237[_lhs238] = _rhs375;
      let _2461_i1;
      _2461_i1 = _dafny.ZERO;
      L12: {
        while ((_2449_v19).f4) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2461_i1)) {
              break L12;
            }
            _2461_i1 = (_2461_i1).plus(_dafny.ONE);
            _2437_v7 = _2437_v7;
            let _2462_v23;
            let _init69 = ((_2463_v12) => function (_2464_i2) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(205)), ((_2465_v12) => function (_2466_i3) {
                return _2465_v12;
              })(_2463_v12));
            })(_2442_v12);
            let _nw371 = Array((new BigNumber(23)).toNumber());
            for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw371.length); _i0_69++) {
              _nw371[_i0_69] = _init69(new BigNumber(_i0_69));
            }
            _2462_v23 = _nw371;
            let _2467_v24;
            _2467_v24 = _module.D1.create_DC4((_2437_v7)[_module.__default.safeIndex(new BigNumber(803), new BigNumber((_2437_v7).length))], _2436_v6);
            let _2468_v25;
            _2468_v25 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_module.__default.fm47((_2449_v19).f6, globalState))).length));
            let _2469_v26;
            let _nw372 = new _module.C1();
            _nw372.__ctor((_2437_v7)[_module.__default.safeIndex(new BigNumber(803), new BigNumber((_2437_v7).length))], new BigNumber(-616), _2449_v19.f5, new BigNumber((_2443_v13).length), p2);
            _2469_v26 = _nw372;
            let _index394 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_2462_v23).length));
            (_2462_v23)[_index394] = _module.__default.fm50((_2467_v24).dtor_cf3, new BigNumber((_dafny.Seq.update(_2468_v25, _module.__default.safeIndex(p1, new BigNumber((_2468_v25).length)), new BigNumber((_dafny.Seq.of(p1, (_2449_v19).f6)).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_2442_v12, globalState),_2469_v26)).length), (_2449_v19).f4, globalState);
            let _index395 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_2462_v23).length));
            (_2462_v23)[_index395] = _dafny.Seq.update((((_2449_v19).f4) ? (_dafny.Seq.Concat(_2440_v10, _2440_v10)) : (_2440_v10)), _module.__default.safeIndex((new BigNumber((_2440_v10).length)).plus(new BigNumber((_2440_v10).length)), new BigNumber(((((_2449_v19).f4) ? (_dafny.Seq.Concat(_2440_v10, _2440_v10)) : (_2440_v10))).length)), _2442_v12);
            let _2470_v27;
            _2470_v27 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),p0);
            _2470_v27 = (_2470_v27).update(_2442_v12, p1);
            _2447_v17 = _2449_v19.f5;
          }
        }
      }
      let _2471_v28;
      _2471_v28 = _2448_v18;
      let _2472_v29;
      _2472_v29 = _dafny.Seq.of(_2448_v18, _2471_v28, _2471_v28);
      let _2473_v30;
      _2473_v30 = _dafny.Set.fromElements(_2472_v29);
      _2473_v30 = _2473_v30;
      let _2474_v31;
      _2474_v31 = false;
      _2474_v31 = (_2449_v19).f4;
      return;
    }
    m5(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _2475_v0;
      _2475_v0 = false;
      if (_2475_v0) {
        r0 = !(_2475_v0);
        let _2476_v1;
        _2476_v1 = _dafny.Set.fromElements(_2475_v0, _2475_v0);
        r1 = (_2476_v1).IsSubsetOf(_2476_v1);
        let _2477_v2;
        _2477_v2 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p0).minus(new BigNumber((_dafny.Seq.of(new BigNumber(259))).length)));
        _2477_v2 = (_2477_v2).update(p0, p0);
        r1 = (p0).isLessThanOrEqualTo(p0);
        let _2478_v3;
        let _nw373 = Array((new BigNumber(4)).toNumber()).fill([]);
        _2478_v3 = _nw373;
        let _2479_v4;
        let _nw374 = new _module.C3();
        _nw374.__ctor(_2478_v3, p0, _2475_v0);
        _2479_v4 = _nw374;
      } else {
        let _2480_v5;
        _2480_v5 = _dafny.Seq.UnicodeFromString("h");
        _2475_v0 = ((p0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_2480_v5).length)))).isEqualTo(new BigNumber(793));
        if (!(_2475_v0)) {
          let _2481_v6;
          _2481_v6 = _module.D2.create_DC6();
          _2481_v6 = _2481_v6;
          let _2482_v7;
          _2482_v7 = _dafny.Seq.of(new BigNumber(270), (_dafny.ZERO).minus(p0));
          let _2483_v8;
          _2483_v8 = _dafny.Seq.of(_2475_v0, _2475_v0);
          let _2484_v9;
          _2484_v9 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_2483_v8)).cardinality()));
          let _2485_v10;
          let _nw375 = Array((new BigNumber(23)).toNumber());
          _nw375[(_dafny.ZERO).toNumber()] = p0;
          _nw375[(_dafny.ONE).toNumber()] = new BigNumber((_2480_v5).length);
          _nw375[(new BigNumber(2)).toNumber()] = p0;
          _nw375[(new BigNumber(3)).toNumber()] = new BigNumber((_2482_v7).length);
          _nw375[(new BigNumber(4)).toNumber()] = new BigNumber(-91);
          _nw375[(new BigNumber(5)).toNumber()] = p0;
          _nw375[(new BigNumber(6)).toNumber()] = p0;
          _nw375[(new BigNumber(7)).toNumber()] = p0;
          _nw375[(new BigNumber(8)).toNumber()] = new BigNumber((_2484_v9).cardinality());
          _nw375[(new BigNumber(9)).toNumber()] = p0;
          _nw375[(new BigNumber(10)).toNumber()] = p0;
          _nw375[(new BigNumber(11)).toNumber()] = p0;
          _nw375[(new BigNumber(12)).toNumber()] = new BigNumber(-870);
          _nw375[(new BigNumber(13)).toNumber()] = p0;
          _nw375[(new BigNumber(14)).toNumber()] = p0;
          _nw375[(new BigNumber(15)).toNumber()] = p0;
          _nw375[(new BigNumber(16)).toNumber()] = p0;
          _nw375[(new BigNumber(17)).toNumber()] = p0;
          _nw375[(new BigNumber(18)).toNumber()] = new BigNumber(951);
          _nw375[(new BigNumber(19)).toNumber()] = new BigNumber(424);
          _nw375[(new BigNumber(20)).toNumber()] = p0;
          _nw375[(new BigNumber(21)).toNumber()] = p0;
          _nw375[(new BigNumber(22)).toNumber()] = p0;
          _2485_v10 = _nw375;
          let _2486_v11;
          _2486_v11 = _dafny.Set.fromElements(_2485_v10, _2485_v10, _2485_v10);
          let _2487_v12;
          _2487_v12 = _2486_v11;
          let _2488_v13;
          _2488_v13 = _dafny.Set.fromElements(_2486_v11, _2487_v12);
          _2475_v0 = (_2488_v13).IsProperSubsetOf(_2488_v13);
          let _2489_v14;
          _2489_v14 = _dafny.Map.Empty.slice().updateUnsafe((p0).multipliedBy(p0),_2475_v0);
          _2489_v14 = (_2489_v14).update((p0).minus((_2482_v7)[_module.__default.safeIndex(new BigNumber(-179), new BigNumber((_2482_v7).length))]), _2475_v0);
          let _2490_v15;
          _2490_v15 = new _dafny.CodePoint('d'.codePointAt(0));
          let _2491_v16;
          _2491_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2475_v0,_2490_v15);
          let _2492_v17;
          _2492_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2475_v0,new BigNumber((_2491_v16).length));
          let _2493_v18;
          _2493_v18 = _module.D10.create_DC21(_2492_v17);
          let _2494_v19;
          _2494_v19 = _dafny.Set.fromElements(p0, new BigNumber(85));
          let _2495_v22;
          let _nw376 = Array((new BigNumber(19)).toNumber());
          _nw376[(_dafny.ZERO).toNumber()] = (_2494_v19).Intersect(_2494_v19);
          _nw376[(_dafny.ONE).toNumber()] = (_2494_v19).Difference(_2494_v19);
          _nw376[(new BigNumber(2)).toNumber()] = _module.__default.fm31(new BigNumber((_2489_v14).length), p0, p0, false, globalState);
          _nw376[(new BigNumber(3)).toNumber()] = _2494_v19;
          _nw376[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_2480_v5).length));
          _nw376[(new BigNumber(5)).toNumber()] = _2494_v19;
          _nw376[(new BigNumber(6)).toNumber()] = (_dafny.Set.fromElements(new BigNumber(738), p0)).Difference(_2494_v19);
          _nw376[(new BigNumber(7)).toNumber()] = _2494_v19;
          _nw376[(new BigNumber(8)).toNumber()] = _2494_v19;
          _nw376[(new BigNumber(9)).toNumber()] = _2494_v19;
          _nw376[(new BigNumber(10)).toNumber()] = _2494_v19;
          _nw376[(new BigNumber(11)).toNumber()] = function () {
            let _coll60 = new _dafny.Set();
            for (const _compr_60 of _dafny.IntegerRange(new BigNumber(141), new BigNumber(982))) {
              let _2496_v20 = _compr_60;
              if (((new BigNumber(141)).isLessThanOrEqualTo(_2496_v20)) && ((_2496_v20).isLessThan(new BigNumber(982)))) {
                _coll60.add((_2496_v20).minus(p0));
              }
            }
            return _coll60;
          }();
          _nw376[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(new BigNumber(230));
          _nw376[(new BigNumber(13)).toNumber()] = _2494_v19;
          _nw376[(new BigNumber(14)).toNumber()] = _2494_v19;
          _nw376[(new BigNumber(15)).toNumber()] = (_dafny.Set.fromElements(p0)).Intersect(_2494_v19);
          _nw376[(new BigNumber(16)).toNumber()] = function () {
            let _coll61 = new _dafny.Set();
            for (const _compr_61 of _dafny.IntegerRange(new BigNumber(-129), new BigNumber(732))) {
              let _2497_v21 = _compr_61;
              if (((new BigNumber(-129)).isLessThanOrEqualTo(_2497_v21)) && ((_2497_v21).isLessThan(new BigNumber(732)))) {
                _coll61.add(_module.__default.safeDivisionInt(_2497_v21, p0));
              }
            }
            return _coll61;
          }();
          _nw376[(new BigNumber(17)).toNumber()] = _2494_v19;
          _nw376[(new BigNumber(18)).toNumber()] = _2494_v19;
          _2495_v22 = _nw376;
          let _index396 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_2495_v22).length));
          (_2495_v22)[_index396] = (_2494_v19).Difference(_2494_v19);
          let _2498_v23;
          _2498_v23 = _dafny.Seq.of(_2480_v5);
          let _2499_v24;
          _2499_v24 = _dafny.MultiSet.fromElements(_2475_v0, false, !(_2475_v0));
          let _2500_v25;
          _2500_v25 = _dafny.Set.fromElements(_2475_v0);
          let _2501_v26;
          _2501_v26 = _module.D16.create_DC33(_2475_v0, new BigNumber((_2480_v5).length), _2499_v24, p0, _2500_v25);
          let _2502_v27;
          _2502_v27 = _dafny.Set.fromElements(p0, p0, p0, p0, p0);
          let _index397 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_2495_v22).length));
          let _rhs376 = !(false) || (!_dafny.Seq.contains((_2498_v23)[_module.__default.safeIndex(p0, new BigNumber((_2498_v23).length))], _2490_v15));
          let _rhs377 = _module.D10.create_DC21(_2492_v17);
          let _rhs378 = (_2501_v26).dtor_cf52;
          let _rhs379 = (_2502_v27);
          let _lhs239 = _2495_v22;
          let _lhs240 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_2495_v22).length));
          _2475_v0 = _rhs376;
          _2493_v18 = _rhs377;
          r1 = _rhs378;
          _lhs239[_lhs240] = _rhs379;
          let _2503_v28;
          _2503_v28 = new BigNumber(113);
          _2503_v28 = (_dafny.ZERO).minus(p0);
        } else {
          let _2504_v29;
          _2504_v29 = new BigNumber(868);
          _2504_v29 = (_dafny.ZERO).minus(p0);
          r0 = _2475_v0;
          let _2505_v30;
          _2505_v30 = _dafny.MultiSet.fromElements(_2504_v29);
          let _2506_v31;
          _2506_v31 = _dafny.Seq.of(_2505_v30);
          _2504_v29 = new BigNumber((_2506_v31).length);
          let _2507_v32;
          let _nw377 = Array((new BigNumber(27)).toNumber()).fill([]);
          _2507_v32 = _nw377;
          let _2508_v33;
          _2508_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _2509_v34;
          _2509_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2508_v33,_2475_v0);
          let _2510_v35;
          let _nw378 = new _module.C4();
          _nw378.__ctor(_2507_v32, p0, (_module.__default.fm67(_2475_v0, globalState)).equals(_2509_v34));
          _2510_v35 = _nw378;
          let _2511_v36;
          let _init70 = ((_2512_v0) => function (_2513_i0) {
            return (_dafny.Set.fromElements(_2512_v0, _2512_v0)).IsSubsetOf(_dafny.Set.fromElements(_2512_v0, _2512_v0));
          })(_2475_v0);
          let _nw379 = Array((new BigNumber(28)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw379.length); _i0_70++) {
            _nw379[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2511_v36 = _nw379;
          let _2514_v37;
          _2514_v37 = _module.D21.create_DC44(_2475_v0);
          let _pat_let_tv47 = _2475_v0;
          let _index398 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_2511_v36).length));
          (_2511_v36)[_index398] = (function (_pat_let57_0) {
            return function (_2515_dt__update__tmp_h1) {
              return function (_pat_let58_0) {
                return function (_2516_dt__update_hcf72_h0) {
                  return _module.D21.create_DC44(_2516_dt__update_hcf72_h0);
                }(_pat_let58_0);
              }(_pat_let_tv47);
            }(_pat_let57_0);
          }(_2514_v37)).dtor_cf72;
          let _index399 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_2511_v36).length));
          (_2511_v36)[_index399] = true;
        }
        _2480_v5 = _2480_v5;
        let _2517_v38;
        let _init71 = ((_2518_p0) => function (_2519_i1) {
          return (_2519_i1).plus(_2518_p0);
        })(p0);
        let _nw380 = Array((new BigNumber(4)).toNumber());
        for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw380.length); _i0_71++) {
          _nw380[_i0_71] = _init71(new BigNumber(_i0_71));
        }
        _2517_v38 = _nw380;
        let _index400 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2517_v38).length));
        (_2517_v38)[_index400] = new BigNumber(647);
        let _2520_v39;
        _2520_v39 = _dafny.MultiSet.fromElements(!(_2475_v0));
        let _index401 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2517_v38).length));
        (_2517_v38)[_index401] = (new BigNumber((_2520_v39).cardinality())).multipliedBy(p0);
        let _2521_v40;
        _2521_v40 = _dafny.Seq.of(_2475_v0, _2475_v0, _2475_v0, _2475_v0);
        let _2522_v41;
        _2522_v41 = _module.D4.create_DC13(new BigNumber(137), _2475_v0, new BigNumber((_2480_v5).length), _2521_v40);
        let _2523_v42;
        _2523_v42 = _dafny.MultiSet.fromElements((_2517_v38)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_2517_v38).length))]);
        let _index402 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_2517_v38).length));
        (_2517_v38)[_index402] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update((_2522_v41).dtor_cf17, _module.__default.safeIndex((_dafny.ZERO).minus((((_2523_v42).contains(p0)) ? ((_2523_v42).get(p0)) : (p0))), new BigNumber(((_2522_v41).dtor_cf17).length)), _2475_v0), _2521_v40), _dafny.Seq.Concat(_2521_v40, _2521_v40))).length);
      }
      let _2524_v43;
      _2524_v43 = _dafny.Set.fromElements(false);
      _2475_v0 = ((_2524_v43).Difference(_2524_v43)).IsDisjointFrom((_2524_v43).Difference(_2524_v43));
      if (_2475_v0) {
        let _2525_v44;
        _2525_v44 = new _dafny.CodePoint('p'.codePointAt(0));
        let _2526_v45;
        _2526_v45 = _module.D4.create_DC12(_2525_v44);
        let _pat_let_tv48 = _2525_v44;
        let _pat_let_tv49 = _2525_v44;
        let _pat_let_tv50 = _2475_v0;
        let _source40 = function (_pat_let59_0) {
          return function (_2527_dt__update__tmp_h2) {
            return function (_pat_let60_0) {
              return function (_2528_dt__update_hcf13_h0) {
                return _module.D4.create_DC12(_2528_dt__update_hcf13_h0);
              }(_pat_let60_0);
            }(((_pat_let_tv50) ? (_pat_let_tv48) : (_pat_let_tv49)));
          }(_pat_let59_0);
        }(_2526_v45);
        if (_source40.is_DC13) {
          let _2529___mcc_h0 = (_source40).cf14;
          let _2530___mcc_h1 = (_source40).cf15;
          let _2531___mcc_h2 = (_source40).cf16;
          let _2532___mcc_h3 = (_source40).cf17;
          let _2533_cf17 = _2532___mcc_h3;
          let _2534_cf16 = _2531___mcc_h2;
          let _2535_cf15 = _2530___mcc_h1;
          let _2536_cf14 = _2529___mcc_h0;
          _2534_cf16 = _module.__default.fm3(_2536_cf14, _2535_cf15, globalState);
          let _2537_v46;
          let _nw381 = new _module.C2();
          _nw381.__ctor(_dafny.Seq.UnicodeFromString("bgj"), _2475_v0);
          _2537_v46 = _nw381;
          let _2538_v47;
          _2538_v47 = _dafny.Seq.of((_2534_cf16).plus(p0), new BigNumber(-328));
          _2538_v47 = _dafny.Seq.Concat(_2538_v47, _2538_v47);
          r0 = _2535_cf15;
        } else {
          let _2539___mcc_h4 = (_source40).cf13;
          let _2540_cf13 = _2539___mcc_h4;
          let _2541_v48;
          let _nw382 = Array((new BigNumber(10)).toNumber()).fill([]);
          _2541_v48 = _nw382;
          let _2542_v49;
          let _nw383 = new _module.C1();
          _nw383.__ctor(p0, p0, _2541_v48, p0, _2475_v0);
          _2542_v49 = _nw383;
          _2542_v49 = _2542_v49;
          (_2542_v49).f16 = (_2542_v49).fm7(globalState);
          let _2543_v50;
          _2543_v50 = _dafny.Map.Empty.slice().updateUnsafe(_2475_v0,_2475_v0);
          let _2544_v51;
          _2544_v51 = _dafny.MultiSet.fromElements(_2475_v0);
          _2543_v50 = (_2543_v50).update(true, (_dafny.MultiSet.FromArray(_dafny.Seq.of(_2475_v0))).equals(_2544_v51));
          let _2545_v52;
          _2545_v52 = _dafny.MultiSet.fromElements(_2525_v44, _2525_v44, _2540_cf13, _2525_v44);
          _2545_v52 = ((_2475_v0) ? (_2545_v52) : (_dafny.MultiSet.fromElements(_2540_cf13)));
        }
        let _2546_v53;
        let _nw384 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2546_v53 = _nw384;
        let _index403 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_2546_v53).length));
        (_2546_v53)[_index403] = _2525_v44;
        let _index404 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_2546_v53).length));
        (_2546_v53)[_index404] = _2525_v44;
        let _2547_v54;
        _2547_v54 = new BigNumber(-581);
        _2547_v54 = _2547_v54;
        let _2548_v55;
        _2548_v55 = _dafny.Seq.UnicodeFromString("pbuouf");
        let _2549_v56;
        _2549_v56 = _dafny.MultiSet.fromElements(_2548_v55, _2548_v55);
        let _2550_v57;
        _2550_v57 = _module.D11.create_DC25(((_dafny.ZERO).minus(p0)).isEqualTo(p0), _module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_2549_v56).cardinality()))), (_dafny.ZERO).minus(p0)));
        _2550_v57 = _2550_v57;
        let _2551_v58;
        let _nw385 = Array((new BigNumber(4)).toNumber());
        _nw385[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_2547_v54);
        _nw385[(_dafny.ONE).toNumber()] = p0;
        _nw385[(new BigNumber(2)).toNumber()] = _2547_v54;
        _nw385[(new BigNumber(3)).toNumber()] = _2547_v54;
        _2551_v58 = _nw385;
        let _2552_v59;
        _2552_v59 = _module.D16.create_DC32(_2551_v58);
        let _source41 = _2552_v59;
        if (_source41.is_DC33) {
          let _2553___mcc_h5 = (_source41).cf52;
          let _2554___mcc_h6 = (_source41).cf53;
          let _2555___mcc_h7 = (_source41).cf54;
          let _2556___mcc_h8 = (_source41).cf55;
          let _2557___mcc_h9 = (_source41).cf56;
          let _2558_cf56 = _2557___mcc_h9;
          let _2559_cf55 = _2556___mcc_h8;
          let _2560_cf54 = _2555___mcc_h7;
          let _2561_cf53 = _2554___mcc_h6;
          let _2562_cf52 = _2553___mcc_h5;
          r0 = !(((_2562_cf52) ? (((_2475_v0) ? (_2475_v0) : (true))) : (_2475_v0)));
          let _2563_v60;
          _2563_v60 = _module.D11.create_DC26(_2475_v0);
          let _2564_v61;
          _2564_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2475_v0,_2475_v0);
          let _2565_v62;
          _2565_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2563_v60,(((_2564_v61).contains(_2475_v0)) ? ((_2564_v61).get(_2475_v0)) : (_2562_cf52)));
          let _2566_v63;
          _2566_v63 = _dafny.Seq.of(_2565_v62);
          let _2567_v64;
          _2567_v64 = _module.D18.create_DC35(_2566_v63);
          let _2568_v65;
          let _nw386 = Array((new BigNumber(24)).toNumber()).fill([]);
          _2568_v65 = _nw386;
          let _2569_v66;
          _2569_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2525_v44,_2475_v0);
          let _2570_v67;
          let _nw387 = Array((new BigNumber(10)).toNumber()).fill([]);
          _2570_v67 = _nw387;
          let _2571_v68;
          let _nw388 = new _module.C6();
          _nw388.__ctor(_2475_v0, (((_2569_v66).contains(new _dafny.CodePoint('g'.codePointAt(0)))) ? ((_2569_v66).get(new _dafny.CodePoint('g'.codePointAt(0)))) : (_2475_v0)), _2570_v67, _2559_cf55);
          _2571_v68 = _nw388;
          let _2572_v69;
          let _nw389 = Array((new BigNumber(27)).toNumber());
          _nw389[(_dafny.ZERO).toNumber()] = _2571_v68;
          _nw389[(_dafny.ONE).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(2)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(3)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(4)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(5)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(6)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(7)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(8)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(9)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(10)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(11)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(12)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(13)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(14)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(15)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(16)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(17)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(18)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(19)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(20)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(21)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(22)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(23)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(24)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(25)).toNumber()] = _2571_v68;
          _nw389[(new BigNumber(26)).toNumber()] = _2571_v68;
          _2572_v69 = _nw389;
          let _2573_v70;
          _2573_v70 = _dafny.MultiSet.fromElements(_2572_v69, _2572_v69);
          let _2574_v71;
          _2574_v71 = _dafny.Map.Empty.slice().updateUnsafe(((_2562_cf52) ? (_2573_v70) : (_2573_v70)),_2525_v44);
          let _index405 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_2546_v53).length));
          let _rhs380 = _module.__default.fm1(new BigNumber((_2548_v55).length), _2475_v0, globalState);
          let _rhs381 = _2562_cf52;
          let _rhs382 = _2567_v64;
          let _rhs383 = _2568_v65;
          let _rhs384 = (((_2574_v71).contains(_dafny.MultiSet.fromElements(_2572_v69))) ? ((_2574_v71).get(_dafny.MultiSet.fromElements(_2572_v69))) : (_2525_v44));
          let _lhs241 = _2546_v53;
          let _lhs242 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_2546_v53).length));
          _lhs241[_lhs242] = _rhs380;
          _2562_cf52 = _rhs381;
          _2567_v64 = _rhs382;
          _2568_v65 = _rhs383;
          _2525_v44 = _rhs384;
          let _2575_v72;
          _2575_v72 = _dafny.Seq.of(_2558_cf56);
          let _2576_v73;
          _2576_v73 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2559_cf55),new BigNumber(-139));
          let _2577_v74;
          _2577_v74 = _dafny.Seq.of(_2562_cf52, (_2571_v68).f4);
          let _2578_v75;
          _2578_v75 = _module.D9.create_DC19(_2548_v55);
          let _pat_let_tv51 = _2548_v55;
          let _2579_v76;
          let _nw390 = Array((new BigNumber(29)).toNumber());
          _nw390[(_dafny.ZERO).toNumber()] = _2562_cf52;
          _nw390[(_dafny.ONE).toNumber()] = _2475_v0;
          _nw390[(new BigNumber(2)).toNumber()] = _2562_cf52;
          _nw390[(new BigNumber(3)).toNumber()] = true;
          _nw390[(new BigNumber(4)).toNumber()] = !(_2562_cf52);
          _nw390[(new BigNumber(5)).toNumber()] = (_2571_v68).f4;
          _nw390[(new BigNumber(6)).toNumber()] = (_2562_cf52) && (false);
          _nw390[(new BigNumber(7)).toNumber()] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(614)), ((_2580_v53) => function (_2581_i2) {
            return (_2580_v53)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_2580_v53).length))];
          })(_2546_v53))).length)).isLessThan(_2547_v54);
          _nw390[(new BigNumber(8)).toNumber()] = _2562_cf52;
          _nw390[(new BigNumber(9)).toNumber()] = _2562_cf52;
          _nw390[(new BigNumber(10)).toNumber()] = _2475_v0;
          _nw390[(new BigNumber(11)).toNumber()] = ((_dafny.ZERO).minus(p0)).isLessThan(_2547_v54);
          _nw390[(new BigNumber(12)).toNumber()] = ((_2575_v72)[_module.__default.safeIndex((((_2576_v73).contains((_2571_v68).f6)) ? ((_2576_v73).get((_2571_v68).f6)) : (_2547_v54)), new BigNumber((_2575_v72).length))]).contains(!(_2475_v0));
          _nw390[(new BigNumber(13)).toNumber()] = _2475_v0;
          _nw390[(new BigNumber(14)).toNumber()] = _module.__default.fm0(_2525_v44, globalState);
          _nw390[(new BigNumber(15)).toNumber()] = (_2577_v74)[_module.__default.safeIndex(p0, new BigNumber((_2577_v74).length))];
          _nw390[(new BigNumber(16)).toNumber()] = _2475_v0;
          _nw390[(new BigNumber(17)).toNumber()] = false;
          _nw390[(new BigNumber(18)).toNumber()] = _2562_cf52;
          _nw390[(new BigNumber(19)).toNumber()] = (_2571_v68).f4;
          _nw390[(new BigNumber(20)).toNumber()] = _2475_v0;
          _nw390[(new BigNumber(21)).toNumber()] = _2562_cf52;
          _nw390[(new BigNumber(22)).toNumber()] = _2562_cf52;
          _nw390[(new BigNumber(23)).toNumber()] = _2475_v0;
          _nw390[(new BigNumber(24)).toNumber()] = (_2571_v68).f4;
          _nw390[(new BigNumber(25)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), ((_2582_v53) => function (_2583_i3) {
            return (_2582_v53)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_2582_v53).length))];
          })(_2546_v53)), (function (_pat_let61_0) {
            return function (_2584_dt__update__tmp_h3) {
              return function (_pat_let62_0) {
                return function (_2585_dt__update_hcf25_h0) {
                  return _module.D9.create_DC19(_2585_dt__update_hcf25_h0);
                }(_pat_let62_0);
              }(_pat_let_tv51);
            }(_pat_let61_0);
          }(_2578_v75)).dtor_cf25);
          _nw390[(new BigNumber(26)).toNumber()] = false;
          _nw390[(new BigNumber(27)).toNumber()] = (_2571_v68).f4;
          _nw390[(new BigNumber(28)).toNumber()] = !((_2571_v68).f4) || (_2562_cf52);
          _2579_v76 = _nw390;
          let _index406 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_2579_v76).length));
          (_2579_v76)[_index406] = _2562_cf52;
          let _2586_v77;
          _2586_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2576_v73,(_2571_v68).f4);
          let _2587_v78;
          _2587_v78 = _dafny.Map.Empty.slice().updateUnsafe((_2571_v68).f6,_2551_v58);
          let _2588_v79;
          _2588_v79 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(new BigNumber((_2586_v77).length), _2475_v0, globalState),(((_2587_v78).contains(_2561_cf53)) ? ((_2587_v78).get(_2561_cf53)) : (_2551_v58)));
          let _2589_v80;
          _2589_v80 = _module.D20.create_DC41(_2547_v54);
          let _2590_v81;
          _2590_v81 = _dafny.MultiSet.fromElements((_2571_v68).f6);
          let _2591_v82;
          _2591_v82 = _dafny.Seq.of((_2571_v68).f6, _2559_cf55);
          let _2592_v83;
          _2592_v83 = _dafny.MultiSet.fromElements(new BigNumber((_2590_v81).cardinality()), new BigNumber((_2591_v82).length), _2547_v54, _2561_cf53, new BigNumber(943));
          let _2593_v84;
          let _nw391 = new _module.C1();
          _nw391.__ctor(new BigNumber(925), p0, _2570_v67, _2561_cf53, !(_2475_v0));
          _2593_v84 = _nw391;
          let _2594_v85;
          _2594_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(625),_dafny.Seq.of(new BigNumber(((_2590_v81).update(_2559_cf55, _module.__default.abs(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2593_v84,new BigNumber((_2560_cf54).cardinality()))).length)))).cardinality())));
          let _index407 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_2579_v76).length));
          let _rhs385 = (_2571_v68).fm8((_2589_v80).dtor_cf67, _2562_cf52, _dafny.Seq.Create(_module.__default.abs(new BigNumber(69)), ((_2595_v53) => function (_2596_i4) {
            return (_2595_v53)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_2595_v53).length))];
          })(_2546_v53)), (((_2594_v85).contains(new BigNumber(852))) ? ((_2594_v85).get(new BigNumber(852))) : (_2591_v82)), globalState);
          let _rhs386 = (_2571_v68).f4;
          let _rhs387 = !((_2571_v68).f4);
          let _rhs388 = (_2588_v79).Merge(_2588_v79);
          let _lhs243 = _2579_v76;
          let _lhs244 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_2579_v76).length));
          _lhs243[_lhs244] = _rhs385;
          _2562_cf52 = _rhs386;
          r0 = _rhs387;
          _2588_v79 = _rhs388;
          let _2597_v86;
          let _nw392 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _2597_v86 = _nw392;
          let _2598_v87;
          _2598_v87 = _dafny.Set.fromElements(_2579_v76, _2579_v76);
          let _2599_v88;
          _2599_v88 = _module.D22.create_DC45(_2597_v86);
          let _2600_v89;
          _2600_v89 = _dafny.Map.Empty.slice().updateUnsafe(_2598_v87,(_2599_v88).dtor_cf73);
          _2597_v86 = (((_2600_v89).contains(_2598_v87)) ? ((_2600_v89).get(_2598_v87)) : (_2597_v86));
        } else {
          let _2601___mcc_h10 = (_source41).cf51;
          let _2602_cf51 = _2601___mcc_h10;
          let _2603_v90;
          let _nw393 = Array((new BigNumber(22)).toNumber()).fill(false);
          _2603_v90 = _nw393;
          let _2604_v91;
          _2604_v91 = _dafny.Seq.of(_2547_v54, p0, p0, new BigNumber(13));
          let _2605_v92;
          _2605_v92 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_2603_v90, _2603_v90, _2603_v90, _2603_v90, _2603_v90),new BigNumber((_dafny.MultiSet.FromArray(_2604_v91)).cardinality()));
          let _2606_v93;
          _2606_v93 = _dafny.Map.Empty.slice().updateUnsafe((((_2605_v92).contains(_dafny.Seq.of(_2603_v90, _2603_v90, _2603_v90, _2603_v90))) ? ((_2605_v92).get(_dafny.Seq.of(_2603_v90, _2603_v90, _2603_v90, _2603_v90))) : (_2547_v54)),_module.__default.fm40(false, true, p0, globalState));
          _2606_v93 = (_2606_v93).update(new BigNumber(192), _module.__default.fm19(_2475_v0, p0, globalState));
          let _index408 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_2603_v90).length));
          (_2603_v90)[_index408] = (_this).fm11(globalState);
          let _index409 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_2603_v90).length));
          (_2603_v90)[_index409] = _2475_v0;
          let _2607_v94;
          let _init72 = ((_2608_v44) => function (_2609_i5) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(370)), ((_2610_v44) => function (_2611_i6) {
              return _2610_v44;
            })(_2608_v44));
          })(_2525_v44);
          let _nw394 = Array((new BigNumber(5)).toNumber());
          for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw394.length); _i0_72++) {
            _nw394[_i0_72] = _init72(new BigNumber(_i0_72));
          }
          _2607_v94 = _nw394;
          let _2612_v95;
          let _init73 = ((_2613_p0, _2614_v54, _2615_v90) => function (_2616_i7) {
            return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2613_p0,(_dafny.ZERO).minus(_2614_v54))).length),(_2615_v90)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_2615_v90).length))]);
          })(p0, _2547_v54, _2603_v90);
          let _nw395 = Array((new BigNumber(24)).toNumber());
          for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw395.length); _i0_73++) {
            _nw395[_i0_73] = _init73(new BigNumber(_i0_73));
          }
          _2612_v95 = _nw395;
          let _rhs389 = _2607_v94;
          let _rhs390 = _2612_v95;
          _2607_v94 = _rhs389;
          _2612_v95 = _rhs390;
          let _2617_v96;
          _2617_v96 = _dafny.Seq.of(_2548_v55, _dafny.Seq.Create(_module.__default.abs(new BigNumber(525)), ((_2618_v44) => function (_2619_i8) {
            return _2618_v44;
          })(_2525_v44)), _2548_v55, _2548_v55);
          let _2620_v97;
          _2620_v97 = _module.D12.create_DC27(_dafny.Seq.Create(_module.__default.abs(new BigNumber(195)), ((_2621_v0) => function (_2622_i9) {
  return _dafny.Seq.of(_2621_v0);
})(_2475_v0)));
          let _2623_v98;
          _2623_v98 = _dafny.Set.fromElements(_2547_v54, p0);
          let _2624_v99;
          _2624_v99 = _dafny.Map.Empty.slice().updateUnsafe(!((_2603_v90)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_2603_v90).length))]),_2623_v98);
          let _index410 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_2603_v90).length));
          let _rhs391 = (p0).isLessThanOrEqualTo(p0);
          let _rhs392 = _module.__default.fm39(_dafny.Seq.Concat(_dafny.Seq.update(_2548_v55, _module.__default.safeIndex(_2547_v54, new BigNumber((_2548_v55).length)), (_2546_v53)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_2546_v53).length))]), (_2617_v96)[_module.__default.safeIndex(p0, new BigNumber((_2617_v96).length))]), _2620_v97, _2475_v0, globalState);
          let _rhs393 = (_this).fm11(globalState);
          let _rhs394 = (new BigNumber(((((_2624_v99).contains(false)) ? ((_2624_v99).get(false)) : (_2623_v98))).length)).plus(_2547_v54);
          let _rhs395 = _2547_v54;
          let _lhs245 = _2603_v90;
          let _lhs246 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_2603_v90).length));
          r1 = _rhs391;
          _2604_v91 = _rhs392;
          _lhs245[_lhs246] = _rhs393;
          _2547_v54 = _rhs394;
          _2547_v54 = _rhs395;
        }
      } else {
        let _2625_v100;
        _2625_v100 = _dafny.MultiSet.fromElements(_2475_v0);
        r0 = ((_2625_v100).Difference(_2625_v100)).IsDisjointFrom(_dafny.MultiSet.fromElements(_2475_v0, _2475_v0, _2475_v0));
        let _2626_v101;
        let _init74 = ((_2627_v0) => function (_2628_i10) {
          return (_2627_v0) === (_2627_v0);
        })(_2475_v0);
        let _nw396 = Array((new BigNumber(24)).toNumber());
        for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw396.length); _i0_74++) {
          _nw396[_i0_74] = _init74(new BigNumber(_i0_74));
        }
        _2626_v101 = _nw396;
        let _index411 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_2626_v101).length));
        (_2626_v101)[_index411] = _2475_v0;
        let _2629_v102;
        _2629_v102 = _dafny.Map.Empty.slice().updateUnsafe(_2475_v0,p0);
        let _index412 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_2626_v101).length));
        (_2626_v101)[_index412] = ((((_2629_v102).contains(_2475_v0)) ? ((_2629_v102).get(_2475_v0)) : (p0))).isEqualTo(_module.__default.fm3(new BigNumber((_2524_v43).length), _2475_v0, globalState));
        let _2630_v103;
        let _init75 = ((_2631_p0) => function (_2632_i11) {
          return (_2632_i11).minus(_2631_p0);
        })(p0);
        let _nw397 = Array((new BigNumber(11)).toNumber());
        for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw397.length); _i0_75++) {
          _nw397[_i0_75] = _init75(new BigNumber(_i0_75));
        }
        _2630_v103 = _nw397;
        let _index413 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2630_v103).length));
        (_2630_v103)[_index413] = (_dafny.ZERO).minus(p0);
        let _index414 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2630_v103).length));
        (_2630_v103)[_index414] = p0;
        let _2633_v104;
        let _nw398 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
        _2633_v104 = _nw398;
        let _2634_v105;
        _2634_v105 = _module.D22.create_DC45(_2633_v104);
        let _pat_let_tv52 = _2634_v105;
        let _2635_v106;
        let _nw399 = Array((new BigNumber(4)).toNumber());
        _nw399[(_dafny.ZERO).toNumber()] = _2634_v105;
        _nw399[(_dafny.ONE).toNumber()] = _2634_v105;
        _nw399[(new BigNumber(2)).toNumber()] = function (_pat_let63_0) {
          return function (_2636_dt__update__tmp_h4) {
            return function (_pat_let64_0) {
              return function (_2637_dt__update_hcf73_h0) {
                return _module.D22.create_DC45(_2637_dt__update_hcf73_h0);
              }(_pat_let64_0);
            }((_pat_let_tv52).dtor_cf73);
          }(_pat_let63_0);
        }(_2634_v105);
        _nw399[(new BigNumber(3)).toNumber()] = _2634_v105;
        _2635_v106 = _nw399;
        let _index415 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_2635_v106).length));
        (_2635_v106)[_index415] = _2634_v105;
        let _index416 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_2635_v106).length));
        (_2635_v106)[_index416] = _2634_v105;
        let _2638_v107;
        let _nw400 = new _module.C0();
        _nw400.__ctor((_2630_v103)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_2630_v103).length))]);
        _2638_v107 = _nw400;
        let _2639_v108;
        _2639_v108 = _dafny.Seq.of(_2638_v107, _2638_v107);
        let _2640_v109;
        _2640_v109 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2475_v0,p0)).length),_2638_v107);
        let _2641_v110;
        _2641_v110 = _module.D23.create_DC48(_2638_v107);
        let _2642_v111;
        let _nw401 = Array((new BigNumber(24)).toNumber());
        _nw401[(_dafny.ZERO).toNumber()] = _2638_v107;
        _nw401[(_dafny.ONE).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(2)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(3)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(4)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(5)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(6)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(7)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(8)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(9)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(10)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(11)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(12)).toNumber()] = (_2639_v108)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_2638_v107).f12, p0)).length), new BigNumber((_2639_v108).length))];
        _nw401[(new BigNumber(13)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(14)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(15)).toNumber()] = (((_2640_v109).contains(p0)) ? ((_2640_v109).get(p0)) : (_2638_v107));
        _nw401[(new BigNumber(16)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(17)).toNumber()] = (_2641_v110).dtor_cf82;
        _nw401[(new BigNumber(18)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(19)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(20)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(21)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(22)).toNumber()] = _2638_v107;
        _nw401[(new BigNumber(23)).toNumber()] = (_2641_v110).dtor_cf82;
        _2642_v111 = _nw401;
        _2642_v111 = _2642_v111;
      }
      let _2643_v112;
      let _init76 = ((_2644_p0) => function (_2645_i13) {
        return (_2645_i13).multipliedBy(_2644_p0);
      })(p0);
      let _nw402 = Array((new BigNumber(29)).toNumber());
      for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw402.length); _i0_76++) {
        _nw402[_i0_76] = _init76(new BigNumber(_i0_76));
      }
      _2643_v112 = _nw402;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2643_v112).length))) {
        let _2646_i12 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2646_i12)) && ((_2646_i12).isLessThan(new BigNumber((_2643_v112).length))))) {
          (_2643_v112)[(_2646_i12)] = _module.__default.safeModuloInt(_2646_i12, (p0).plus(p0));
        }
      }
      let _2647_v113;
      let _nw403 = Array((new BigNumber(19)).toNumber());
      _2647_v113 = _nw403;
      let _2648_v114;
      let _init77 = ((_2649_v0) => function (_2650_i14) {
        return _2649_v0;
      })(_2475_v0);
      let _nw404 = Array((new BigNumber(26)).toNumber());
      for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw404.length); _i0_77++) {
        _nw404[_i0_77] = _init77(new BigNumber(_i0_77));
      }
      _2648_v114 = _nw404;
      let _2651_v115;
      let _nw405 = Array((new BigNumber(23)).toNumber());
      _nw405[(_dafny.ZERO).toNumber()] = _2648_v114;
      _nw405[(_dafny.ONE).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(2)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(3)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(4)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(5)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(6)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(7)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(8)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(9)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(10)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(11)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(12)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(13)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(14)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(15)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(16)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(17)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(18)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(19)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(20)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(21)).toNumber()] = _2648_v114;
      _nw405[(new BigNumber(22)).toNumber()] = _2648_v114;
      _2651_v115 = _nw405;
      let _2652_v116;
      _2652_v116 = _dafny.Seq.UnicodeFromString("clpbitvi");
      let _2653_v117;
      _2653_v117 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("vvtigi"), _2652_v116);
      let _2654_v118;
      let _nw406 = new _module.C3();
      _nw406.__ctor(_2651_v115, new BigNumber((_2653_v117).length), _2475_v0);
      _2654_v118 = _nw406;
      let _index417 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_2647_v113).length));
      (_2647_v113)[_index417] = _2654_v118;
      let _2655_v119;
      _2655_v119 = _dafny.Map.Empty.slice().updateUnsafe(_2475_v0,p0);
      let _2656_v120;
      let _nw407 = new _module.C0();
      _nw407.__ctor((_dafny.ZERO).minus(new BigNumber((_2655_v119).length)));
      _2656_v120 = _nw407;
      let _index418 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_2647_v113).length));
      let _rhs396 = _2475_v0;
      let _rhs397 = _2654_v118;
      let _rhs398 = _2656_v120;
      let _lhs247 = _2647_v113;
      let _lhs248 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_2647_v113).length));
      r1 = _rhs396;
      _lhs247[_lhs248] = _rhs397;
      _2656_v120 = _rhs398;
      let _2657_v121;
      _2657_v121 = new BigNumber(-191);
      _2657_v121 = ((_dafny.ZERO).minus((_2656_v120).f12)).minus((_2656_v120).f12);
      let _2658_v122;
      _2658_v122 = _module.D22.create_DC47(new BigNumber((_2652_v116).length), !(!(_2475_v0)), _2475_v0, (_2656_v120).f12, (_dafny.ZERO).minus(p0));
      r0 = (_2658_v122).dtor_cf79;
      r1 = _2475_v0;
      return [r0, r1];
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f5 = [];
      this._f6 = _dafny.ZERO;
      this._f4 = false;
      this.f7 = _dafny.ZERO;
      this._f8 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    set f5(value) {
      let _this = this;
      _this._f5 = value;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f7, f8, f5, f6, f4) {
      let _this = this;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f4 = f4;
      return;
    }
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f4;
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      let _source42 = (((_this).f4) ? (_module.D1.create_DC4(new BigNumber((_dafny.Seq.of(_this.f7)).length), new BigNumber(912))) : (_module.D1.create_DC4((_this).f6, new BigNumber(392))));
      if (_source42.is_DC4) {
        let _2659___mcc_h0 = (_source42).cf3;
        let _2660___mcc_h1 = (_source42).cf4;
        let _2661_cf4 = _2660___mcc_h1;
        let _2662_cf3 = _2659___mcc_h0;
        return (_dafny.ZERO).minus(_2662_cf3);
      } else {
        let _2663___mcc_h2 = (_source42).cf2;
        let _2664_cf2 = _2663___mcc_h2;
        return (_this).f6;
      }
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f4))).cardinality()), (_this).f6, _this.f7, new BigNumber(-280))).Union((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("fku")).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_this).f4)).length))));
    };
    fm7(globalState) {
      let _this = this;
      return _this.f7;
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _2665_v0;
      let _init78 = function (_2666_i1) {
        return (_this).f4;
      };
      let _nw408 = Array((new BigNumber(18)).toNumber());
      for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw408.length); _i0_78++) {
        _nw408[_i0_78] = _init78(new BigNumber(_i0_78));
      }
      _2665_v0 = _nw408;
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2665_v0).length))) {
        let _2667_i0 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2667_i0)) && ((_2667_i0).isLessThan(new BigNumber((_2665_v0).length))))) {
          (_2665_v0)[(_2667_i0)] = (p0).isLessThanOrEqualTo((_this).f6);
        }
      }
      let _hi14 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_this.f7, _this.f7));
      for (let _2668_i2 = _module.__default.safeModuloInt(_this.f7, _this.f7); _2668_i2.isLessThan(_hi14); _2668_i2 = _2668_i2.plus(_dafny.ONE)) {
        let _2669_v1;
        _2669_v1 = _module.D1.create_DC4((_this).f6, p0);
        let _2670_v2;
        _2670_v2 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f4),((false) ? (_2669_v1) : (_2669_v1)));
        let _2671_v3;
        _2671_v3 = _dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)));
        let _2672_v4;
        _2672_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2671_v3);
        let _2673_v5;
        _2673_v5 = new _dafny.CodePoint('h'.codePointAt(0));
        let _2674_v6;
        _2674_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber((_dafny.Seq.update(_2671_v3, _module.__default.safeIndex(_2668_i2, new BigNumber((_2671_v3).length)), _2673_v5)).length));
        _2670_v2 = (_2670_v2).update((_this).f4, _module.__default.fm10((_this).f4, false, _2672_v4, new BigNumber((_2674_v6).length), globalState));
        let _2675_v7;
        _2675_v7 = true;
        _2675_v7 = (_this).f4;
        let _2676_v8;
        let _nw409 = new _module.C0();
        _nw409.__ctor((_2668_i2).plus((_dafny.ZERO).minus((_this).f6)));
        _2676_v8 = _nw409;
        if (_2675_v7) {
          let _2677_v9;
          _2677_v9 = _dafny.MultiSet.fromElements((_2676_v8).f12, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_2674_v6).length), new BigNumber(886))).cardinality()));
          let _2678_v10;
          _2678_v10 = _dafny.Map.Empty.slice().updateUnsafe((_2677_v9).Intersect(_2677_v9),new BigNumber(451));
          let _2679_v11;
          _2679_v11 = _module.D0.create_DC0(_2668_i2);
          let _2680_v12;
          _2680_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_2673_v5);
          _2678_v10 = (_2678_v10).update((_dafny.MultiSet.fromElements(p0, (_dafny.ZERO).minus((_2676_v8).f12), _this.f7)).Union((_this).fm6(_2679_v11, _2679_v11, _2680_v12, new BigNumber(251), globalState)), _this.f7);
          let _rhs399 = (_this).f4;
          let _rhs400 = new BigNumber((_dafny.Seq.of(_dafny.Seq.of(_2668_i2))).length);
          let _lhs249 = _this;
          _2675_v7 = _rhs399;
          _lhs249.f7 = _rhs400;
          _2675_v7 = _2675_v7;
          let _2681_v13;
          let _init79 = ((_2682_v3) => function (_2683_i3) {
            return _2682_v3;
          })(_2671_v3);
          let _nw410 = Array((new BigNumber(14)).toNumber());
          for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw410.length); _i0_79++) {
            _nw410[_i0_79] = _init79(new BigNumber(_i0_79));
          }
          _2681_v13 = _nw410;
          let _index419 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_2681_v13).length));
          (_2681_v13)[_index419] = _module.__default.fm16(_2675_v7, (_this).f4, !(_2675_v7), _module.D2.create_DC6(), globalState);
          let _2684_v14;
          _2684_v14 = _dafny.MultiSet.fromElements(_2679_v11, _2679_v11);
          let _2685_v15;
          let _nw411 = Array((new BigNumber(2)).toNumber());
          _nw411[(_dafny.ZERO).toNumber()] = _2684_v14;
          _nw411[(_dafny.ONE).toNumber()] = _2684_v14;
          _2685_v15 = _nw411;
          let _index420 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_2685_v15).length));
          (_2685_v15)[_index420] = _dafny.MultiSet.fromElements(_2679_v11);
          let _2686_v16;
          _2686_v16 = _dafny.Seq.of(function (_pat_let65_0) {
            return function (_2687_dt__update__tmp_h0) {
              return function (_pat_let66_0) {
                return function (_2688_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_2688_dt__update_hcf0_h0);
                }(_pat_let66_0);
              }((_dafny.ZERO).minus(_2668_i2));
            }(_pat_let65_0);
          }(_2679_v11));
          let _index421 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_2681_v13).length));
          let _index422 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_2685_v15).length));
          let _rhs401 = _2671_v3;
          let _rhs402 = (_2684_v14).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_2686_v16, _2686_v16)));
          let _lhs250 = _2681_v13;
          let _lhs251 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_2681_v13).length));
          let _lhs252 = _2685_v15;
          let _lhs253 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_2685_v15).length));
          _lhs250[_lhs251] = _rhs401;
          _lhs252[_lhs253] = _rhs402;
          let _2689_v17;
          _2689_v17 = _module.D12.create_DC28(p0, _this.f7, (_this).f4, true, p0);
          let _2690_v18;
          _2690_v18 = _dafny.Seq.of((_2689_v17).dtor_cf46, (_this).f4, (_2675_v7) === (_2675_v7));
          let _2691_v19;
          _2691_v19 = _dafny.Seq.of((_this).f6);
          _2690_v18 = _dafny.Seq.update(_dafny.Seq.update(_2690_v18, _module.__default.safeIndex(_2668_i2, new BigNumber((_2690_v18).length)), (_2690_v18)[_module.__default.safeIndex(new BigNumber((_2691_v19).length), new BigNumber((_2690_v18).length))]), _module.__default.safeIndex(_module.__default.safeModuloInt((_this).f6, _this.f7), new BigNumber((_dafny.Seq.update(_2690_v18, _module.__default.safeIndex(_2668_i2, new BigNumber((_2690_v18).length)), (_2690_v18)[_module.__default.safeIndex(new BigNumber((_2691_v19).length), new BigNumber((_2690_v18).length))])).length)), !(true));
        } else {
          let _2692_v20;
          _2692_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2671_v3,_module.__default.fm1(_this.f7, _2675_v7, globalState));
          _2692_v20 = (_2692_v20).update(_dafny.Seq.UnicodeFromString("yrkp"), _2673_v5);
          let _2693_v21;
          let _nw412 = new _module.C6();
          _nw412.__ctor((_this).f4, (_this).f4, ((_2675_v7) ? (_this.f5) : (_this.f5)), (_2676_v8).f12);
          _2693_v21 = _nw412;
          let _2694_v22;
          _2694_v22 = _dafny.Set.fromElements(!(_2675_v7));
          _2675_v7 = !(_2694_v22).equals(_module.__default.fm46(new _dafny.CodePoint('t'.codePointAt(0)), globalState));
          let _2695_v23;
          _2695_v23 = _module.D3.create_DC9(true, p0, (_this).f4);
          let _2696_v24;
          _2696_v24 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f6),(_2695_v23).dtor_cf9);
          let _2697_v25;
          _2697_v25 = _dafny.Seq.of(_2675_v7, (_this).f4);
          let _2698_v26;
          _2698_v26 = _dafny.Seq.of(new BigNumber((_2697_v25).length), (_this).f6);
          let _2699_v27;
          _2699_v27 = _module.D6.create_DC16(false, (_2693_v21).f11, _2698_v26);
          let _2700_v28;
          _2700_v28 = _dafny.Set.fromElements(_2699_v27, _2699_v27, _2699_v27);
          let _rhs403 = (_module.__default.fm42(_2673_v5, _2700_v28, globalState)).Merge(_2696_v24);
          let _rhs404 = (_2668_i2).isLessThan((_2676_v8).f12);
          let _rhs405 = !((_2676_v8).f12).isEqualTo((_this).fm7(globalState));
          let _rhs406 = (p0).minus(new BigNumber(345));
          let _rhs407 = p0;
          let _lhs254 = _this;
          let _lhs255 = _this;
          _2696_v24 = _rhs403;
          _2675_v7 = _rhs404;
          _2675_v7 = _rhs405;
          _lhs254.f7 = _rhs406;
          _lhs255.f7 = _rhs407;
          _2675_v7 = !((_2676_v8).f12).isEqualTo(((_this).f6).minus(_this.f7));
        }
      }
      let _2701_i4;
      _2701_i4 = _dafny.ZERO;
      L13: {
        while ((_this).f4) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2701_i4)) {
              break L13;
            }
            _2701_i4 = (_2701_i4).plus(_dafny.ONE);
            (_this).f7 = (_dafny.ZERO).minus((_this).f6);
            let _2702_v29;
            _2702_v29 = _dafny.Seq.UnicodeFromString("cnnryc");
            let _2703_v30;
            let _nw413 = new _module.C2();
            _nw413.__ctor(_2702_v29, (_this).f4);
            _2703_v30 = _nw413;
            let _2704_v31;
            _2704_v31 = _dafny.Seq.of(_2703_v30, _2703_v30, _2703_v30);
            let _2705_v32;
            _2705_v32 = _2703_v30;
            let _2706_v33;
            let _nw414 = Array((new BigNumber(18)).toNumber());
            _nw414[(_dafny.ZERO).toNumber()] = _2703_v30;
            _nw414[(_dafny.ONE).toNumber()] = (_2704_v31)[_module.__default.safeIndex(p0, new BigNumber((_2704_v31).length))];
            _nw414[(new BigNumber(2)).toNumber()] = (_2704_v31)[_module.__default.safeIndex(new BigNumber((_2702_v29).length), new BigNumber((_2704_v31).length))];
            _nw414[(new BigNumber(3)).toNumber()] = (((_this).f4) ? (_2703_v30) : (_2703_v30));
            _nw414[(new BigNumber(4)).toNumber()] = (_2705_v32);
            _nw414[(new BigNumber(5)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(6)).toNumber()] = (((_this).f4) ? (_2703_v30) : (_2703_v30));
            _nw414[(new BigNumber(7)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(8)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(9)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(10)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(11)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(12)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(13)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(14)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(15)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(16)).toNumber()] = _2703_v30;
            _nw414[(new BigNumber(17)).toNumber()] = _2703_v30;
            _2706_v33 = _nw414;
            let _index423 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_2706_v33).length));
            (_2706_v33)[_index423] = _2703_v30;
            let _index424 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_2706_v33).length));
            (_2706_v33)[_index424] = _2703_v30;
            let _source43 = _module.__default.fm68(globalState);
            if (_source43.is_DC28) {
              let _2707___mcc_h0 = (_source43).cf43;
              let _2708___mcc_h1 = (_source43).cf44;
              let _2709___mcc_h2 = (_source43).cf45;
              let _2710___mcc_h3 = (_source43).cf46;
              let _2711___mcc_h4 = (_source43).cf47;
              let _2712_cf47 = _2711___mcc_h4;
              let _2713_cf46 = _2710___mcc_h3;
              let _2714_cf45 = _2709___mcc_h2;
              let _2715_cf44 = _2708___mcc_h1;
              let _2716_cf43 = _2707___mcc_h0;
              _2716_cf43 = new BigNumber((_dafny.Seq.Concat(_2702_v29, _2702_v29)).length);
              let _2717_v34;
              _2717_v34 = _dafny.Map.Empty.slice().updateUnsafe((_2703_v30).f4,!(_2713_cf46));
              _2717_v34 = _2717_v34;
              let _2718_v35;
              let _init80 = ((_2719_cf46, _2720_v30, _2721_cf45) => function (_2722_i5) {
                return _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_2719_cf46, (_2720_v30).f4, !((_2720_v30).f4), (_this).f4, _2721_cf45)).length));
              })(_2713_cf46, _2703_v30, _2714_cf45);
              let _nw415 = Array((new BigNumber(28)).toNumber());
              for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw415.length); _i0_80++) {
                _nw415[_i0_80] = _init80(new BigNumber(_i0_80));
              }
              _2718_v35 = _nw415;
              let _index425 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_2718_v35).length));
              (_2718_v35)[_index425] = _dafny.Set.fromElements(_2715_cf44);
              let _index426 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_2718_v35).length));
              (_2718_v35)[_index426] = (_dafny.Set.fromElements(p0, p0, new BigNumber((_dafny.Seq.of(_2713_cf46)).length), _2716_cf43, (_this).fm7(globalState))).Intersect(_dafny.Set.fromElements((_this).f6));
              let _rhs408 = _2714_cf45;
              let _rhs409 = ((_2718_v35)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2718_v35).length))]).IsProperSubsetOf(((_2718_v35)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2718_v35).length))]).Intersect((_2718_v35)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2718_v35).length))]));
              _2714_cf45 = _rhs408;
              _2713_cf46 = _rhs409;
            } else {
              let _2723___mcc_h5 = (_source43).cf42;
              let _2724_cf42 = _2723___mcc_h5;
              let _2725_v36;
              _2725_v36 = _dafny.Seq.of((_this).f4, (_2703_v30).f4, (_2703_v30).f4);
              _2725_v36 = _dafny.Seq.update(_2725_v36, _module.__default.safeIndex(_this.f7, new BigNumber((_2725_v36).length)), (_this).f4);
              let _2726_v37;
              _2726_v37 = false;
              _2726_v37 = _2726_v37;
              let _index427 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_2665_v0).length));
              (_2665_v0)[_index427] = _2726_v37;
              let _2727_v38;
              _2727_v38 = _dafny.MultiSet.fromElements((_this).f4, (_2703_v30).f4, true);
              let _index428 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_2665_v0).length));
              (_2665_v0)[_index428] = (_dafny.MultiSet.fromElements(_2726_v37)).IsSubsetOf(_2727_v38);
              (_this).f7 = p0;
            }
            let _2728_v39;
            _2728_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2703_v30,(_2703_v30).f4);
            let _2729_v40;
            _2729_v40 = _module.D12.create_DC28(_this.f7, new BigNumber((_2702_v29).length), (_2703_v30).f4, true, new BigNumber((_2702_v29).length));
            let _2730_v41;
            _2730_v41 = _dafny.Set.fromElements((((_2728_v39).contains((_2706_v33)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_2706_v33).length))])) ? ((_2728_v39).get((_2706_v33)[_module.__default.safeIndex(new BigNumber(704), new BigNumber((_2706_v33).length))])) : ((_2703_v30).f4)), false, (_2729_v40).dtor_cf45, (_this).f4, (_this).f4);
            let _2731_v42;
            _2731_v42 = _dafny.Map.Empty.slice().updateUnsafe((_2703_v30).f4,(_2703_v30).f4);
            let _2732_v43;
            _2732_v43 = _dafny.Seq.of(_2731_v42, _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_2703_v30).f4));
            let _2733_v44;
            _2733_v44 = _dafny.Seq.of(_2730_v41, _module.__default.fm53((_this).f6, _2732_v43, globalState));
            (_this).f7 = _module.__default.safeDivisionInt(new BigNumber(((_2733_v44)[_module.__default.safeIndex(p0, new BigNumber((_2733_v44).length))]).length), (new BigNumber(308)).plus((_this).f6));
          }
        }
      }
      let _2734_v45;
      _2734_v45 = _this.f5;
      let _2735_v46;
      _2735_v46 = false;
      let _2736_v47;
      _2736_v47 = _dafny.Set.fromElements(_this.f7, _this.f7);
      let _2737_v48;
      _2737_v48 = _dafny.Seq.of(false, (_2736_v47).IsProperSubsetOf(_2736_v47), _2735_v46, (_2735_v46) === (_2735_v46));
      let _2738_v49;
      _2738_v49 = _dafny.MultiSet.fromElements(p0, (_this).f6);
      let _rhs410 = _2734_v45;
      let _rhs411 = new BigNumber((_2737_v48).length);
      let _rhs412 = (_2738_v49).equals((_2738_v49).Difference(_2738_v49));
      let _lhs256 = _this;
      _2734_v45 = _rhs410;
      _lhs256.f7 = _rhs411;
      _2735_v46 = _rhs412;
      let _2739_v50;
      let _nw416 = new _module.C0();
      _nw416.__ctor(p0);
      _2739_v50 = _nw416;
      let _2740_v51;
      _2740_v51 = new _dafny.CodePoint('d'.codePointAt(0));
      let _2741_v52;
      _2741_v52 = _dafny.Seq.UnicodeFromString("nom");
      let _2742_v53;
      _2742_v53 = _module.D3.create_DC9((_this).f4, p0, true);
      let _index429 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_2665_v0).length));
      (_2665_v0)[_index429] = (_2742_v53).dtor_cf10;
      let _2743_v54;
      _2743_v54 = _dafny.MultiSet.fromElements(_2738_v49);
      let _2744_v55;
      _2744_v55 = _dafny.Map.Empty.slice().updateUnsafe(_2740_v51,_2743_v54);
      let _index430 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_2665_v0).length));
      let _rhs413 = ((_2735_v46) ? (_this.f7) : (new BigNumber(-737)));
      let _rhs414 = _2740_v51;
      let _rhs415 = _dafny.Seq.UnicodeFromString("dxypm");
      let _rhs416 = ((((_2744_v55).contains(_2740_v51)) ? ((_2744_v55).get(_2740_v51)) : (_2743_v54))).IsProperSubsetOf(_2743_v54);
      let _rhs417 = _2735_v46;
      let _lhs257 = _this;
      let _lhs258 = _2665_v0;
      let _lhs259 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_2665_v0).length));
      _lhs257.f7 = _rhs413;
      _2740_v51 = _rhs414;
      _2741_v52 = _rhs415;
      _2735_v46 = _rhs416;
      _lhs258[_lhs259] = _rhs417;
      let _2745_v56;
      _2745_v56 = _dafny.Map.Empty.slice().updateUnsafe((_2665_v0)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_2665_v0).length))],(_2739_v50).f12);
      r0 = (_2745_v56).Merge(_dafny.Map.Empty.slice().updateUnsafe(!((_2665_v0)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_2665_v0).length))]),(_2739_v50).f12));
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = [];
      let _2746_v0;
      let _nw417 = new _module.C4();
      _nw417.__ctor(_this.f5, _module.__default.safeDivisionInt((_this).fm7(globalState), _this.f7), !(p0));
      _2746_v0 = _nw417;
      let _2747_i0;
      _2747_i0 = _dafny.ZERO;
      L14: {
        while (p0) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2747_i0)) {
              break L14;
            }
            _2747_i0 = (_2747_i0).plus(_dafny.ONE);
            let _2748_v1;
            let _nw418 = new _module.C8();
            _nw418.__ctor();
            _2748_v1 = _nw418;
            let _2749_v2;
            _2749_v2 = _dafny.Seq.UnicodeFromString("bj");
            let _2750_v3;
            _2750_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,_2749_v2);
            _2750_v3 = _module.__default.fm69(globalState);
            let _2751_v4;
            let _init81 = function (_2752_i1) {
              return _module.__default.safeModuloInt(_2752_i1, new BigNumber((_dafny.Seq.of((_this).f4, false)).length));
            };
            let _nw419 = Array((new BigNumber(22)).toNumber());
            for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw419.length); _i0_81++) {
              _nw419[_i0_81] = _init81(new BigNumber(_i0_81));
            }
            _2751_v4 = _nw419;
            _2751_v4 = _2751_v4;
            (_this).f7 = (_this).f6;
          }
        }
      }
      let _2753_v5;
      let _nw420 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _2753_v5 = _nw420;
      _2753_v5 = _2753_v5;
      if (p0) {
        let _index431 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_2753_v5).length));
        (_2753_v5)[_index431] = (_this).f6;
        let _index432 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_2753_v5).length));
        (_2753_v5)[_index432] = _module.__default.safeModuloInt((_this).f6, (_this).f6);
        let _2754_v6;
        _2754_v6 = true;
        let _2755_v7;
        _2755_v7 = _dafny.Set.fromElements(p0);
        let _2756_v8;
        _2756_v8 = new _dafny.CodePoint('w'.codePointAt(0));
        _2754_v6 = ((_module.__default.fm46(_2756_v8, globalState)).Intersect(_2755_v7)).IsProperSubsetOf((_2755_v7).Intersect(_2755_v7));
        let _2757_v9;
        let _nw421 = new _module.C3();
        _nw421.__ctor(_this.f5, (_2753_v5)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_2753_v5).length))], _2754_v6);
        _2757_v9 = _nw421;
        let _2758_v10;
        let _nw422 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
        _2758_v10 = _nw422;
        let _index433 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_2758_v10).length));
        (_2758_v10)[_index433] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(767)), ((_2759_v5) => function (_2760_i2) {
          return (_2759_v5)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_2759_v5).length))];
        })(_2753_v5));
        let _2761_v11;
        _2761_v11 = _dafny.Seq.of(_this.f7);
        let _2762_v12;
        _2762_v12 = _dafny.Seq.UnicodeFromString("xondmryd");
        let _2763_v13;
        _2763_v13 = _dafny.Seq.of(p0, (_this).f4, (_this).f4);
        let _2764_v14;
        _2764_v14 = _module.D12.create_DC27(_dafny.Seq.of(_dafny.Seq.update(_2763_v13, _module.__default.safeIndex((_2753_v5)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_2753_v5).length))], new BigNumber((_2763_v13).length)), _2754_v6)));
        let _2765_v15;
        _2765_v15 = _dafny.Seq.of(_dafny.Seq.Concat(_2761_v11, _dafny.Seq.of((_this).f6)), _dafny.Seq.Concat(_2761_v11, _module.__default.fm39(_2762_v12, _2764_v14, (_this).f4, globalState)));
        let _2766_v16;
        _2766_v16 = _dafny.MultiSet.fromElements((_this).f6, new BigNumber((_dafny.Seq.UnicodeFromString("cprbau")).length));
        let _index434 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_2758_v10).length));
        (_2758_v10)[_index434] = (_2765_v15)[_module.__default.safeIndex(new BigNumber((_2766_v16).cardinality()), new BigNumber((_2765_v15).length))];
        let _2767_v17;
        _2767_v17 = _dafny.MultiSet.fromElements(_2762_v12);
        let _2768_v18;
        _2768_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_2753_v5)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_2753_v5).length))]);
        _2767_v17 = (_2767_v17).Difference((_2767_v17).update(_2762_v12, _module.__default.abs(new BigNumber((_2768_v18).length))));
      } else {
        let _2769_v19;
        _2769_v19 = _dafny.Seq.of((_this).f6, _module.__default.safeModuloInt(_this.f7, _this.f7), _this.f7);
        _2769_v19 = _dafny.Seq.Concat(_2769_v19, _dafny.Seq.Concat(_2769_v19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(484)), function (_2770_i3) {
          return (_this).f6;
        })));
        let _2771_v20;
        _2771_v20 = new _dafny.CodePoint('f'.codePointAt(0));
        let _2772_v21;
        let _nw423 = new _module.C5();
        _nw423.__ctor(!(!(p0)), _this.f7, _this.f5, (_this).f6, _module.__default.fm0(_2771_v20, globalState));
        _2772_v21 = _nw423;
        let _2773_v22;
        _2773_v22 = true;
        let _2774_v23;
        _2774_v23 = _dafny.Set.fromElements(p0);
        let _index435 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_2753_v5).length));
        (_2753_v5)[_index435] = _module.__default.safeDivisionInt(_2772_v21.f14, (_dafny.ZERO).minus(new BigNumber((_2774_v23).length)));
        let _2775_v24;
        _2775_v24 = _dafny.Set.fromElements((_this).f6, new BigNumber(18));
        let _2776_v25;
        _2776_v25 = _dafny.MultiSet.fromElements((_2772_v21).f13, (_2772_v21).f13);
        let _2777_v26;
        _2777_v26 = _dafny.Map.Empty.slice().updateUnsafe((_2772_v21).f13,_dafny.MultiSet.fromElements(p0));
        let _index436 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_2753_v5).length));
        let _rhs418 = (_2775_v24).IsSubsetOf((_2775_v24).Difference(_2775_v24));
        let _rhs419 = _module.__default.safeModuloInt((_this).f6, ((_this).f6).multipliedBy(_this.f7));
        let _rhs420 = (_dafny.MultiSet.fromElements(false, p0)).IsSubsetOf(_2776_v25);
        let _rhs421 = (((_this).f6).multipliedBy(_this.f7)).multipliedBy(((_this).f6).minus((_2772_v21).fm7(globalState)));
        let _rhs422 = !(_2777_v26).contains(!((_2772_v21).f13) || (p0));
        let _lhs260 = _2772_v21;
        let _lhs261 = _2753_v5;
        let _lhs262 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_2753_v5).length));
        _2773_v22 = _rhs418;
        _lhs260.f14 = _rhs419;
        _2773_v22 = _rhs420;
        _lhs261[_lhs262] = _rhs421;
        _2773_v22 = _rhs422;
        _2773_v22 = p0;
        let _index437 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_2753_v5).length));
        (_2753_v5)[_index437] = new BigNumber(45);
      }
      let _2778_v27;
      _2778_v27 = _dafny.Seq.UnicodeFromString("packlnnxk");
      _2778_v27 = _2778_v27;
      let _2779_v28;
      _2779_v28 = _dafny.Seq.of((_this).f4);
      let _2780_v29;
      _2780_v29 = new _dafny.CodePoint('y'.codePointAt(0));
      let _2781_v30;
      _2781_v30 = _dafny.MultiSet.fromElements(_2780_v29);
      let _2782_v31;
      _2782_v31 = _dafny.Seq.of(_2781_v30, _2781_v30);
      if (_dafny.Seq.IsPrefixOf(_module.__default.fm70((_this).f4, new BigNumber(736), _2779_v28, p0, globalState), _dafny.Seq.Concat(_dafny.Seq.of(_2781_v30), _2782_v31))) {
        (_this).f7 = _this.f7;
        let _2783_v32;
        _2783_v32 = _dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_this).f4);
        let _rhs423 = _this.f7;
        let _rhs424 = ((!((((_2783_v32).contains(new BigNumber((_2778_v27).length))) ? ((_2783_v32).get(new BigNumber((_2778_v27).length))) : ((_this).f4))) || (p0)) ? (_this.f7) : ((_this).f6));
        let _rhs425 = _this.f7;
        let _rhs426 = _2779_v28;
        let _lhs263 = _this;
        let _lhs264 = _this;
        let _lhs265 = _this;
        _lhs263.f7 = _rhs423;
        _lhs264.f7 = _rhs424;
        _lhs265.f7 = _rhs425;
        _2779_v28 = _rhs426;
        let _2784_v33;
        _2784_v33 = true;
        _2784_v33 = !(_2784_v33);
        let _2785_v34;
        _2785_v34 = _dafny.Seq.of(_this.f7);
        let _index438 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_2753_v5).length));
        (_2753_v5)[_index438] = (new BigNumber((_2785_v34).length)).multipliedBy((_this).f6);
        let _index439 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2753_v5).length));
        (_2753_v5)[_index439] = (((_this).fm8(_this.f7, _2784_v33, _2778_v27, _dafny.Seq.update(_2785_v34, _module.__default.safeIndex(_this.f7, new BigNumber((_2785_v34).length)), _this.f7), globalState)) ? (new BigNumber((_2779_v28).length)) : (_module.__default.fm3((_this).f6, false, globalState)));
        let _2786_v35;
        _2786_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
        let _2787_v36;
        _2787_v36 = _dafny.MultiSet.fromElements((((_2786_v35).contains(_2784_v33)) ? ((_2786_v35).get(_2784_v33)) : ((_2779_v28)[_module.__default.safeIndex(_this.f7, new BigNumber((_2779_v28).length))])), _2784_v33, !(_2784_v33));
        let _index440 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_2753_v5).length));
        let _index441 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2753_v5).length));
        let _rhs427 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f6, (((_2787_v36).contains((_2746_v0).fm8(_this.f7, false, _2778_v27, _2785_v34, globalState))) ? ((_2787_v36).get((_2746_v0).fm8(_this.f7, false, _2778_v27, _2785_v34, globalState))) : (new BigNumber((_dafny.Seq.of(p0, p0)).length))))));
        let _rhs428 = _this.f7;
        let _rhs429 = _module.__default.fm3((_this).f6, _2784_v33, globalState);
        let _rhs430 = (_2779_v28)[_module.__default.safeIndex((_this).f6, new BigNumber((_2779_v28).length))];
        let _rhs431 = p0;
        let _lhs266 = _2753_v5;
        let _lhs267 = _module.__default.safeIndex(new BigNumber(227), new BigNumber((_2753_v5).length));
        let _lhs268 = _2753_v5;
        let _lhs269 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2753_v5).length));
        let _lhs270 = _this;
        _lhs266[_lhs267] = _rhs427;
        _lhs268[_lhs269] = _rhs428;
        _lhs270.f7 = _rhs429;
        _2784_v33 = _rhs430;
        _2784_v33 = _rhs431;
        let _2788_v37;
        _2788_v37 = _dafny.Seq.of(_2778_v27);
        _2784_v33 = (_2746_v0).fm8((_this).f6, (_this).f4, (_2788_v37)[_module.__default.safeIndex((_this).f6, new BigNumber((_2788_v37).length))], _2785_v34, globalState);
      } else {
        let _2789_v38;
        let _nw424 = Array((_dafny.ONE).toNumber()).fill(false);
        _2789_v38 = _nw424;
        let _2790_v39;
        _2790_v39 = _module.D2.create_DC5(_2789_v38);
        _2790_v39 = _2790_v39;
        _2789_v38 = _2789_v38;
        let _2791_v40;
        let _nw425 = Array((new BigNumber(22)).toNumber()).fill([]);
        _2791_v40 = _nw425;
        let _2792_v41;
        _2792_v41 = _dafny.Seq.of(_2753_v5, _2753_v5);
        let _2793_v42;
        _2793_v42 = _dafny.Map.Empty.slice().updateUnsafe((_2792_v41)[_module.__default.safeIndex(_this.f7, new BigNumber((_2792_v41).length))],_2791_v40);
        _2791_v40 = (((_2793_v42).contains((_2792_v41)[_module.__default.safeIndex((_this).f6, new BigNumber((_2792_v41).length))])) ? ((_2793_v42).get((_2792_v41)[_module.__default.safeIndex((_this).f6, new BigNumber((_2792_v41).length))])) : (_2791_v40));
        let _2794_v43;
        _2794_v43 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f6),_this.f7);
        let _2795_v44;
        _2795_v44 = _dafny.Seq.of((_this).f6);
        (_this).f7 = _module.__default.safeDivisionInt((new BigNumber(480)).plus((((_2794_v43).contains((_2795_v44)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_2795_v44).length))])) ? ((_2794_v43).get((_2795_v44)[_module.__default.safeIndex(new BigNumber(901), new BigNumber((_2795_v44).length))])) : (new BigNumber((_2779_v28).length)))), new BigNumber(444));
        let _2796_v45;
        _2796_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2753_v5,_2789_v38);
        (_this).f7 = new BigNumber((_2796_v45).length);
      }
      r0 = _2780_v29;
      let _nw426 = Array((new BigNumber(9)).toNumber()).fill(_module.D1.Default());
      r1 = _nw426;
      return [r0, r1];
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _2797_v0;
      _2797_v0 = _dafny.Seq.of(true);
      let _2798_v1;
      _2798_v1 = _dafny.Seq.of(_2797_v0);
      let _2799_v2;
      _2799_v2 = _module.D12.create_DC27(_2798_v1);
      let _2800_v3;
      _2800_v3 = _dafny.Seq.of(new BigNumber((_module.__default.fm71(globalState)).length), new BigNumber(953));
      let _2801_v4;
      _2801_v4 = _module.D6.create_DC16((_this).f4, !((_this).f4) || ((_this).f4), _dafny.Seq.Concat(_2800_v3, _2800_v3));
      let _2802_v5;
      _2802_v5 = _dafny.Seq.UnicodeFromString("wweol");
      let _2803_v6;
      _2803_v6 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f4);
      let _2804_v7;
      _2804_v7 = _dafny.Set.fromElements(false);
      let _2805_v8;
      let _nw427 = new _module.C5();
      _nw427.__ctor((_this).f4, new BigNumber((_2804_v7).length), _this.f5, (_this).f6, !(_module.__default.fm0(new _dafny.CodePoint('s'.codePointAt(0)), globalState)));
      _2805_v8 = _nw427;
      let _2806_v9;
      _2806_v9 = _dafny.Seq.of(_2805_v8);
      let _2807_v10;
      _2807_v10 = _dafny.Set.fromElements(_2805_v8, (_2806_v9)[_module.__default.safeIndex(_this.f7, new BigNumber((_2806_v9).length))]);
      let _2808_v11;
      _2808_v11 = _dafny.MultiSet.fromElements(!((_2797_v0)[_module.__default.safeIndex(_2805_v8.f14, new BigNumber((_2797_v0).length))]), (_this).f4, (_this).f4);
      let _2809_v12;
      let _nw428 = Array((new BigNumber(15)).toNumber());
      _nw428[(_dafny.ZERO).toNumber()] = (_this).f6;
      _nw428[(_dafny.ONE).toNumber()] = new BigNumber(((_2803_v6).Merge(_2803_v6)).length);
      _nw428[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_this).f4, (_this).f4, (_this).f4, (_this).f4)).length);
      _nw428[(new BigNumber(3)).toNumber()] = p0;
      _nw428[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(p0, new BigNumber((_2807_v10).length));
      _nw428[(new BigNumber(5)).toNumber()] = (_this).f6;
      _nw428[(new BigNumber(6)).toNumber()] = p0;
      _nw428[(new BigNumber(7)).toNumber()] = (_2805_v8).fm9((_2805_v8).f13, (_2805_v8).f13, p1, globalState);
      _nw428[(new BigNumber(8)).toNumber()] = p1;
      _nw428[(new BigNumber(9)).toNumber()] = (_2805_v8).fm7(globalState);
      _nw428[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((((_2808_v11).contains(false)) ? ((_2808_v11).get(false)) : (p0)));
      _nw428[(new BigNumber(11)).toNumber()] = (p0).minus(_this.f7);
      _nw428[(new BigNumber(12)).toNumber()] = (new BigNumber((_2800_v3).length)).plus(p0);
      _nw428[(new BigNumber(13)).toNumber()] = new BigNumber((_2804_v7).length);
      _nw428[(new BigNumber(14)).toNumber()] = p0;
      _2809_v12 = _nw428;
      let _index442 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
      (_2809_v12)[_index442] = (_module.D4.create_DC13(new BigNumber((_2802_v5).length), (_this).f4, _this.f7, _2797_v0)).dtor_cf14;
      let _2810_v13;
      _2810_v13 = new _dafny.CodePoint('r'.codePointAt(0));
      let _2811_v14;
      _2811_v14 = _module.D9.create_DC19(_dafny.Seq.UnicodeFromString("dks"));
      let _index443 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
      let _rhs432 = _2799_v2;
      let _rhs433 = _module.D6.create_DC16((_2805_v8).f13, _module.__default.fm0(_module.__default.fm1(new BigNumber((_2802_v5).length), (_this).f4, globalState), globalState), _dafny.Seq.update(_2800_v3, _module.__default.safeIndex(p0, new BigNumber((_2800_v3).length)), p1));
      let _rhs434 = ((((_2797_v0)[_module.__default.safeIndex(new BigNumber(-8), new BigNumber((_2797_v0).length))]) ? (_2811_v14) : (_2811_v14))).dtor_cf25;
      let _rhs435 = new BigNumber((_dafny.Seq.UnicodeFromString("lvpib")).length);
      let _rhs436 = _2810_v13;
      let _lhs271 = _2809_v12;
      let _lhs272 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
      _2799_v2 = _rhs432;
      _2801_v4 = _rhs433;
      _2802_v5 = _rhs434;
      _lhs271[_lhs272] = _rhs435;
      _2810_v13 = _rhs436;
      if (false) {
        let _2812_v15;
        _2812_v15 = false;
        let _2813_v16;
        _2813_v16 = _dafny.MultiSet.fromElements(new BigNumber((_2797_v0).length));
        _2812_v15 = (_2813_v16).IsSubsetOf((_2813_v16).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f7))));
        let _index444 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
        (_2809_v12)[_index444] = (_dafny.ZERO).minus(p0);
        let _index445 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
        (_2809_v12)[_index445] = _this.f7;
        let _index446 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
        (_2809_v12)[_index446] = _2805_v8.f14;
        let _2814_v17;
        let _nw429 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
        _2814_v17 = _nw429;
        let _2815_v18;
        _2815_v18 = _dafny.Map.Empty.slice().updateUnsafe(true,_2805_v8.f14);
        let _2816_v19;
        _2816_v19 = _dafny.Set.fromElements(_2815_v18);
        let _index447 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_2814_v17).length));
        (_2814_v17)[_index447] = _2816_v19;
        let _index448 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_2814_v17).length));
        let _rhs437 = (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_2812_v15,_2805_v8.f14), _2815_v18)).Union(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0), _2815_v18, _2815_v18, _2815_v18));
        let _rhs438 = (((_this).f6).minus((_dafny.ZERO).minus(_this.f7))).minus(p0);
        let _rhs439 = (_2805_v8).f13;
        let _lhs273 = _2814_v17;
        let _lhs274 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_2814_v17).length));
        let _lhs275 = _2805_v8;
        _lhs273[_lhs274] = _rhs437;
        _lhs275.f14 = _rhs438;
        _2812_v15 = _rhs439;
      } else {
        let _2817_v20;
        _2817_v20 = true;
        let _2818_v21;
        let _nw430 = Array((new BigNumber(21)).toNumber()).fill(false);
        _2818_v21 = _nw430;
        let _2819_v22;
        _2819_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2818_v21,true);
        _2817_v20 = (((_2819_v22).contains(_2818_v21)) ? ((_2819_v22).get(_2818_v21)) : ((_2805_v8).f13));
        _2817_v20 = _2817_v20;
        let _2820_v23;
        let _nw431 = new _module.C5();
        _nw431.__ctor(((_2805_v8).f13) || (_2817_v20), _this.f7, _this.f5, (_this).f6, ((_2805_v8).f13) && ((_2805_v8).f13));
        _2820_v23 = _nw431;
        _2817_v20 = !((_2820_v23).f13);
        let _2821_v24;
        _2821_v24 = _dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(685)), _2800_v3, _2800_v3, _2800_v3, _module.__default.fm39(_2802_v5, _2799_v2, (_2805_v8).f13, globalState));
        let _2822_v25;
        _2822_v25 = _dafny.MultiSet.fromElements(_2800_v3);
        _2821_v24 = (_2822_v25);
      }
      let _index449 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
      let _rhs440 = (_2809_v12)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length))];
      let _rhs441 = _2809_v12;
      let _lhs276 = _2809_v12;
      let _lhs277 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
      _lhs276[_lhs277] = _rhs440;
      _2809_v12 = _rhs441;
      let _2823_v26;
      _2823_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f4);
      let _2824_v27;
      _2824_v27 = _dafny.MultiSet.fromElements(_2809_v12);
      let _2825_v28;
      let _init82 = ((_2826_v8) => function (_2827_i0) {
        return (_2826_v8).f13;
      })(_2805_v8);
      let _nw432 = Array((new BigNumber(12)).toNumber());
      for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw432.length); _i0_82++) {
        _nw432[_i0_82] = _init82(new BigNumber(_i0_82));
      }
      _2825_v28 = _nw432;
      let _2828_v29;
      _2828_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_2802_v5, _module.__default.safeIndex(new BigNumber((_2824_v27).cardinality()), new BigNumber((_2802_v5).length)), p2)).length),new BigNumber((_dafny.Seq.of(_2825_v28)).length));
      let _2829_v30;
      _2829_v30 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_2828_v29).length));
      let _2830_v31;
      let _nw433 = Array((new BigNumber(16)).toNumber());
      _nw433[(_dafny.ZERO).toNumber()] = (_2805_v8).f13;
      _nw433[(_dafny.ONE).toNumber()] = (_2805_v8).f13;
      _nw433[(new BigNumber(2)).toNumber()] = (_2805_v8).fm8(p1, (_2805_v8).f13, _2802_v5, _2800_v3, globalState);
      _nw433[(new BigNumber(3)).toNumber()] = !((_this).f4);
      _nw433[(new BigNumber(4)).toNumber()] = (_this).f4;
      _nw433[(new BigNumber(5)).toNumber()] = ((_2805_v8).f13) === ((_this).f4);
      _nw433[(new BigNumber(6)).toNumber()] = !(_module.__default.fm0(p2, globalState));
      _nw433[(new BigNumber(7)).toNumber()] = (_this).f4;
      _nw433[(new BigNumber(8)).toNumber()] = (_2805_v8).f13;
      _nw433[(new BigNumber(9)).toNumber()] = ((_this).f4) || ((_module.D4.create_DC13((_2809_v12)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length))], (_2805_v8).f13, p0, _dafny.Seq.update(_2797_v0, _module.__default.safeIndex(new BigNumber((_2797_v0).length), new BigNumber((_2797_v0).length)), (_2805_v8).f13))).dtor_cf15);
      _nw433[(new BigNumber(10)).toNumber()] = (_2805_v8).f13;
      _nw433[(new BigNumber(11)).toNumber()] = !(_2808_v11).equals(_dafny.MultiSet.fromElements((((_2823_v26).contains(new BigNumber((_2829_v30).length))) ? ((_2823_v26).get(new BigNumber((_2829_v30).length))) : ((_2805_v8).f13)), false));
      _nw433[(new BigNumber(12)).toNumber()] = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-420)), ((_2831_v13) => function (_2832_i1) {
        return _2831_v13;
      })(_2810_v13)), _2802_v5));
      _nw433[(new BigNumber(13)).toNumber()] = (_2805_v8).f13;
      _nw433[(new BigNumber(14)).toNumber()] = (_this).f4;
      _nw433[(new BigNumber(15)).toNumber()] = (_this).f4;
      _2830_v31 = _nw433;
      _2825_v28 = _2825_v28;
      let _2833_v32;
      _2833_v32 = true;
      _2833_v32 = false;
      if ((_this).f4) {
        (_2805_v8).f14 = _module.__default.safeDivisionInt(_2805_v8.f14, (_dafny.ZERO).minus((_2809_v12)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length))]));
        _2833_v32 = ((_2805_v8).f13) || ((_dafny.MultiSet.fromElements(new BigNumber((_2797_v0).length), _this.f7)).IsDisjointFrom(_dafny.MultiSet.fromElements(p0)));
        let _index450 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2825_v28).length));
        (_2825_v28)[_index450] = (_2805_v8).f13;
        let _index451 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_2825_v28).length));
        (_2825_v28)[_index451] = (_2805_v8).f13;
        let _2834_v33;
        let _nw434 = new _module.C0();
        _nw434.__ctor(_2805_v8.f14);
        _2834_v33 = _nw434;
        let _index452 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
        (_2809_v12)[_index452] = (_2834_v33).f12;
      } else {
        (_2805_v8).f14 = (_2809_v12)[_module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length))];
        _2803_v6 = (_2803_v6).update((_this).f4, (_2805_v8).f13);
        (_2805_v8).f14 = (_2805_v8.f14).multipliedBy(new BigNumber(248));
        let _index453 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_2809_v12).length));
        (_2809_v12)[_index453] = p0;
        (_2805_v8).f14 = (_this).f6;
      }
      r0 = ((_dafny.Map.Empty.slice().updateUnsafe(_this.f7,(_this).f6)).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2805_v8.f14,new BigNumber((_2802_v5).length))).length), p0)).Merge(_2828_v29);
      return r0;
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
